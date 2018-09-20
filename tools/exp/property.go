package exp

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

// TODO: Break out into multiple files.
// TODO: Deserialize & generated structs preserve "unknown" elements.
// TODO: Satisfy the rest of the sort.Interface.
// TODO: Document and comment this code.

const (
	getMethod                 = "Get"
	setMethod                 = "Set"
	hasMethod                 = "Has"
	clearMethod               = "Clear"
	iteratorClearMethod       = "clear"
	isMethod                  = "Is"
	appendMethod              = "Append"
	prependMethod             = "Prepend"
	removeMethod              = "Remove"
	lenMethod                 = "Len"
	serializeMethod           = "Serialize"
	deserializeMethod         = "Deserialize"
	nameMethod                = "Name"
	serializeIteratorMethod   = "serialize"
	deserializeIteratorMethod = "deserialize"
)

func join(s []*jen.Statement) *jen.Statement {
	return joinLines(s, true)
}

func joinSingleSpace(s []*jen.Statement) *jen.Statement {
	return joinLines(s, false)
}

func joinLines(s []*jen.Statement, double bool) *jen.Statement {
	r := &jen.Statement{}
	for i, stmt := range s {
		if i > 0 {
			r.Line()
			if double {
				r.Line()
			}
		}
		r.Add(stmt.Clone())
	}
	return r
}

type Identifier struct {
	LowerName string
	CamelName string
}

type Kind struct {
	Name                  Identifier
	ConcreteKind          string
	Nilable               bool
	HasNaturalLanguageMap bool
	SerializeFnName       string
	DeserializeFnName     string
}

type PropertyGenerator struct {
	Name       Identifier
	Kinds      []Kind
	asIterator bool
}

func (p *PropertyGenerator) structName() string {
	if p.asIterator {
		return p.Name.CamelName
	}
	return fmt.Sprintf("%sProperty", p.Name.CamelName)
}

func (p *PropertyGenerator) propertyName() string {
	return p.Name.LowerName
}

func (p *PropertyGenerator) deserializeFnName() string {
	if p.asIterator {
		return fmt.Sprintf("%s%s", deserializeIteratorMethod, p.Name.CamelName)
	}
	return fmt.Sprintf("%s%sProperty", deserializeMethod, p.Name.CamelName)
}

func (p *PropertyGenerator) getFnName(i int) string {
	if len(p.Kinds) == 1 {
		return getMethod
	}
	return fmt.Sprintf("%s%s", getMethod, p.kindCamelName(i))
}

func (p *PropertyGenerator) serializeFnName() string {
	if p.asIterator {
		return serializeIteratorMethod
	}
	return serializeMethod
}

func (p *PropertyGenerator) kindCamelName(i int) string {
	return p.Kinds[i].Name.CamelName
}

func (p *PropertyGenerator) memberName(i int) string {
	return fmt.Sprintf("%sMember", p.Kinds[i].Name.LowerName)
}

func (p *PropertyGenerator) hasMemberName(i int) string {
	if len(p.Kinds) == 1 && p.Kinds[0].Nilable {
		panic("PropertyGenerator.hasMemberName called for nilable single value")
	}
	return fmt.Sprintf("has%sMember", p.Kinds[i].Name.CamelName)
}

func (p *PropertyGenerator) clearMethodName() string {
	if p.asIterator {
		return iteratorClearMethod
	}
	return clearMethod
}

func (p *PropertyGenerator) commonFuncs() jen.Code {
	return jen.Commentf(
		"%s returns the name of this property: %q.", nameMethod, p.propertyName(),
	).Line().Func().Params(
		jen.Id("t").Id(p.structName()),
	).Id(nameMethod).Params().Params(
		jen.String(),
	).Block(
		jen.Return(
			jen.Lit(p.propertyName()),
		),
	)
}

type FunctionalPropertyGenerator struct {
	PropertyGenerator
}

func (p *FunctionalPropertyGenerator) Generate() *jen.Statement {
	return p.def().Line().Line().Add(p.funcs()).Line().Line().Add(p.commonFuncs()).Line().Line()
}

func (p *FunctionalPropertyGenerator) def() *jen.Statement {
	if len(p.Kinds) == 1 {
		return p.singleTypeDef()
	} else {
		return p.multiTypeDef()
	}
}

func (p *FunctionalPropertyGenerator) funcs() *jen.Statement {
	if len(p.Kinds) == 1 {
		return p.singleTypeFuncs().Line().Line().Add(p.serializationFuncs())
	} else {
		return p.multiTypeFuncs().Line().Line().Add(p.serializationFuncs())
	}
}

func (p *FunctionalPropertyGenerator) serializationFuncs() *jen.Statement {
	serializeFns := jen.Empty()
	for i, kind := range p.Kinds {
		if i > 0 {
			serializeFns = serializeFns.Else()
		}
		if len(p.Kinds) == 1 {
			serializeFns = serializeFns.If(
				jen.Id("t").Dot(hasMethod).Call(),
			)
		} else {
			serializeFns = serializeFns.If(
				jen.Id("t").Dot(fmt.Sprintf("%s%s", isMethod, p.kindCamelName(i))).Call(),
			)
		}
		serializeFns = serializeFns.Block(
			jen.Return(
				jen.Id(kind.SerializeFnName).Call(
					jen.Id("t").Dot(p.getFnName(i)).Call(),
				),
			),
		)
	}
	serialize := jen.Commentf(
		"%s converts this into an interface representation suitable for marshalling into a text or binary format.", p.serializeFnName(),
	).Line().Func().Params(
		jen.Id("t").Id(p.structName()),
	).Id(p.serializeFnName()).Params().Params(
		jen.Interface(),
		jen.Error(),
	).Block(
		serializeFns,
		jen.Return(
			jen.Nil(),
			jen.Nil(),
		),
	)
	deserializeFns := jen.Empty()
	for i, kind := range p.Kinds {
		if i > 0 {
			deserializeFns = deserializeFns.Else()
		}
		values := jen.Dict{
			jen.Id(p.memberName(i)): jen.Id("v"),
		}
		if !kind.Nilable {
			values[jen.Id(p.hasMemberName(i))] = jen.True()
		}
		deserializeFns = deserializeFns.If(
			jen.List(
				jen.Id("v"),
				jen.Id("handled"),
				jen.Err(),
			).Op(":=").Id(kind.DeserializeFnName).Call(
				jen.Id("i"),
			),
			jen.Id("handled"),
		).Block(
			jen.Id("t").Op(":=").Op("&").Id(p.structName()).Values(
				values,
			),
			jen.Return(
				jen.Id("t"),
				jen.Err(),
			),
		)
	}
	var deserialize jen.Code
	if p.asIterator {
		deserialize = jen.Commentf(
			"%s creates an iterator from an element that has been unmarshalled from a text or binary format.", p.deserializeFnName(),
		).Line().Func().Id(
			p.deserializeFnName(),
		).Params(
			jen.Id("i").Interface(),
		).Params(
			jen.Op("*").Id(p.structName()),
			jen.Error(),
		).Block(
			deserializeFns,
			jen.Return(
				jen.Nil(),
				jen.Nil(),
			),
		)
	} else {
		deserialize = jen.Commentf(
			"%s creates a %q property from an interface representation that has been unmarshalled from a text or binary format.", p.deserializeFnName(), p.propertyName(),
		).Line().Func().Id(
			p.deserializeFnName(),
		).Params(
			jen.Id("m").Map(jen.String()).Interface(),
		).Params(
			jen.Op("*").Id(p.structName()),
			jen.Error(),
		).Block(
			jen.If(
				jen.List(
					jen.Id("i"),
					jen.Id("ok"),
				).Op(":=").Id("m").Index(
					jen.Lit(p.propertyName()),
				),
				jen.Id("ok"),
			).Block(
				deserializeFns,
			),
			jen.Return(
				jen.Nil(),
				jen.Nil(),
			),
		)
	}
	return serialize.Line().Line().Add(deserialize)
}

func (p *FunctionalPropertyGenerator) singleTypeDef() *jen.Statement {
	if p.Kinds[0].Nilable {
		comment := jen.Commentf("%s is the functional property %q. It is permitted to be a single nilable value type.", p.structName(), p.propertyName()).Line()
		if p.asIterator {
			comment = jen.Commentf("%s is an iterator for a property. It is permitted to be a single nilable value type.", p.structName()).Line()
		}
		return comment.Type().Id(p.structName()).Struct(
			jen.Id(p.memberName(0)).Id(p.Kinds[0].ConcreteKind),
		)
	} else {
		comment := jen.Commentf("%s is the functional property %q. It is permitted to be a single default-valued value type.", p.structName(), p.propertyName()).Line()
		if p.asIterator {
			comment = jen.Commentf("%s is an iterator for a property. It is permitted to be a single default-valued value type.", p.structName()).Line()
		}
		return comment.Type().Id(p.structName()).Struct(
			jen.Id(p.memberName(0)).Id(p.Kinds[0].ConcreteKind),
			jen.Id(p.hasMemberName(0)).Bool(),
		)
	}
}

func (p *FunctionalPropertyGenerator) singleTypeFuncs() *jen.Statement {
	memberFunc := jen.Func().Params(
		jen.Id("t").Id(p.structName()),
	)
	memberPtrFunc := jen.Func().Params(
		jen.Id("t").Op("*").Id(p.structName()),
	)
	has := jen.Commentf("%s returns true if this property is set.", hasMethod).Line()
	get := jen.Commentf("%s returns the value of this property. When %s returns false, %s will return any arbitrary value.", getMethod, hasMethod, getMethod).Line()
	set := jen.Commentf("%s sets the value of this property. Calling %s afterwards will return true.", setMethod, hasMethod).Line()
	clear := jen.Commentf("%s ensures no value of this property is set. Calling %s afterwards will return false.", p.clearMethodName(), hasMethod).Line()
	if p.Kinds[0].Nilable {
		has = has.Add(memberFunc.Clone().Id(hasMethod).Params().Params(
			jen.Bool(),
		).Block(
			jen.Return(jen.Id("t").Dot(p.memberName(0)).Op("!=").Nil()),
		))
		get = get.Add(memberFunc.Clone().Id(p.getFnName(0)).Params().Params(
			jen.Id(p.Kinds[0].ConcreteKind),
		).Block(
			jen.Return(jen.Id("t").Dot(p.memberName(0))),
		))
		set = set.Add(memberPtrFunc.Clone().Id(setMethod).Params(
			jen.Id("v").Id(p.Kinds[0].ConcreteKind),
		).Block(
			jen.Id("t").Dot(p.memberName(0)).Op("=").Id("v"),
		))
		clear = clear.Add(memberPtrFunc.Clone().Id(p.clearMethodName()).Params().Block(
			jen.Id("t").Dot(p.memberName(0)).Op("=").Nil(),
		))
	} else {
		has = has.Add(memberFunc.Clone().Id(hasMethod).Params().Params(
			jen.Bool(),
		).Block(
			jen.Return(jen.Id("t").Dot(p.hasMemberName(0))),
		))
		get = get.Add(memberFunc.Clone().Id(p.getFnName(0)).Params().Params(
			jen.Id(p.Kinds[0].ConcreteKind),
		).Block(
			jen.Return(jen.Id("t").Dot(p.memberName(0))),
		))
		set = set.Add(memberPtrFunc.Clone().Id(setMethod).Params(
			jen.Id("v").Id(p.Kinds[0].ConcreteKind),
		).Block(
			jen.Id("t").Dot(p.memberName(0)).Op("=").Id("v"),
			jen.Id("t").Dot(p.hasMemberName(0)).Op("=").True(),
		))
		clear = clear.Add(memberPtrFunc.Clone().Id(p.clearMethodName()).Params().Block(
			jen.Id("t").Dot(p.hasMemberName(0)).Op("=").False(),
		))
	}
	return has.Line().Line().Add(get).Line().Line().Add(set).Line().Line().Add(clear)
}

func (p *FunctionalPropertyGenerator) multiTypeDef() *jen.Statement {
	kindMembers := make([]*jen.Statement, 0, len(p.Kinds))
	for i, kind := range p.Kinds {
		if kind.Nilable {
			kindMembers = append(kindMembers, jen.Id(p.memberName(i)).Id(p.Kinds[i].ConcreteKind))
		} else {
			kindMembers = append(kindMembers, jen.Id(p.memberName(i)).Id(p.Kinds[i].ConcreteKind).Line().Add(
				jen.Id(p.hasMemberName(i)).Bool(),
			))
		}
	}
	explanation := jen.Commentf(
		"At most, one type of value can be present, or none at all. Setting a value will",
	).Line().Commentf(
		"clear the other types of values so that only one of the 'Is' methods will return",
	).Line().Commentf(
		"true.",
	).Line().Comment("").Line().Commentf(
		"It is possible to clear all values, so that this property is empty.",
	).Line()
	comment := jen.Commentf(
		"%s is the functional property %q. It is permitted to be one of multiple value types.", p.structName(), p.propertyName(),
	).Line().Comment("").Line().Add(explanation)
	if p.asIterator {
		comment = jen.Commentf(
			"%s is an iterator for a property. It is permitted to be one of multiple value types.", p.structName(),
		).Line().Comment("").Line().Add(explanation)
	}
	return comment.Type().Id(p.structName()).Struct(joinSingleSpace(kindMembers))
}

func (p *FunctionalPropertyGenerator) multiTypeFuncs() *jen.Statement {
	funcs := &jen.Statement{}
	memberFunc := jen.Func().Params(
		jen.Id("t").Id(p.structName()),
	)
	memberPtrFunc := jen.Func().Params(
		jen.Id("t").Op("*").Id(p.structName()),
	)
	clearLine := make([]*jen.Statement, len(p.Kinds))
	isLine := make([]*jen.Statement, len(p.Kinds))
	for i, kind := range p.Kinds {
		if kind.Nilable {
			clearLine[i] = jen.Id("t").Dot(p.memberName(i)).Op("=").Nil()
		} else {
			clearLine[i] = jen.Id("t").Dot(p.hasMemberName(i)).Op("=").False()
		}
		isLine[i] = jen.Id("t").Dot(fmt.Sprintf("%s%s", isMethod, p.kindCamelName(i))).Parens(nil)
		if i != len(p.Kinds)-1 {
			isLine[i].Op("||")
		}
	}
	for i, kind := range p.Kinds {
		isMethodName := fmt.Sprintf("%s%s", isMethod, p.kindCamelName(i))
		setMethodName := fmt.Sprintf("%s%s", setMethod, p.kindCamelName(i))
		is := jen.Commentf("%s returns true if this property has a type of value of %q.", isMethodName, kind.ConcreteKind).Line()
		get := jen.Commentf("%s returns the value of this property. When %s returns false, %s will return an arbitrary value.", p.getFnName(i), isMethodName, p.getFnName(i)).Line()
		set := jen.Commentf("%s sets the value of this property. Calling %s afterwards returns true.", setMethodName, isMethodName).Line()
		if kind.Nilable {
			is = is.Add(memberFunc.Clone().Id(
				isMethodName,
			).Params().Params(
				jen.Bool(),
			).Block(
				jen.Return(jen.Id("t").Dot(p.memberName(i)).Op("!=").Nil()),
			))
			set = set.Add(memberPtrFunc.Clone().Id(
				setMethodName,
			).Params(
				jen.Id("v").Id(kind.ConcreteKind),
			).Block(
				jen.Id("t").Dot(p.clearMethodName()).Call(),
				jen.Id("t").Dot(p.memberName(i)).Op("=").Id("v"),
			))
		} else {
			is = is.Add(memberFunc.Clone().Id(
				isMethodName,
			).Params().Params(
				jen.Bool(),
			).Block(
				jen.Return(jen.Id("t").Dot(p.hasMemberName(i))),
			))
			set = set.Add(memberPtrFunc.Clone().Id(
				setMethodName,
			).Params(
				jen.Id("v").Id(kind.ConcreteKind),
			).Block(
				jen.Id("t").Dot(p.clearMethodName()).Call(),
				jen.Id("t").Dot(p.memberName(i)).Op("=").Id("v"),
				jen.Id("t").Dot(p.hasMemberName(i)).Op("=").True(),
			))
		}
		get = get.Add(memberFunc.Clone().Id(
			p.getFnName(i),
		).Params().Params(
			jen.Id(kind.ConcreteKind),
		).Block(
			jen.Return(jen.Id("t").Dot(p.memberName(i))),
		))
		funcs.Add(is).Line().Line()
		funcs.Add(get).Line().Line()
		funcs.Add(set).Line().Line()
	}
	hasAnyMethodName := fmt.Sprintf("%sAny", hasMethod)
	clear := jen.Commentf(
		"%s ensures no value of this property is set. Calling %s or any of the 'Is' methods afterwards will return false.", p.clearMethodName(), hasAnyMethodName,
	).Line().Add(memberPtrFunc.Clone().Id(p.clearMethodName()).Params().Block(
		joinSingleSpace(clearLine),
	))
	hasAny := jen.Commentf(
		"%s returns true if any of the different values is set.", hasAnyMethodName,
	).Line().Add(memberFunc.Clone().Id(
		hasAnyMethodName,
	).Params().Params(
		jen.Bool(),
	).Block(
		jen.Return(joinSingleSpace(isLine)),
	))
	funcs.Add(clear).Line().Line().Add(hasAny)
	return funcs
}

type NonFunctionalPropertyGenerator struct {
	PropertyGenerator
}

func (p *NonFunctionalPropertyGenerator) Generate() *jen.Statement {
	return p.def().Line().Line().Add(p.funcs()).Line().Line().Add(p.commonFuncs())
}

func (p *NonFunctionalPropertyGenerator) iteratorTypeName() Identifier {
	return Identifier{
		LowerName: fmt.Sprintf("%sPropertyIterator", p.Name.LowerName),
		CamelName: fmt.Sprintf("%sPropertyIterator", p.Name.CamelName),
	}
}

func (p *NonFunctionalPropertyGenerator) elementTypeGenerator() *FunctionalPropertyGenerator {
	return &FunctionalPropertyGenerator{
		PropertyGenerator{
			Name:       p.iteratorTypeName(),
			Kinds:      p.Kinds,
			asIterator: true,
		},
	}
}

func (p *NonFunctionalPropertyGenerator) def() *jen.Statement {
	comment := jen.Commentf("%s is the non-functional property %q. It is permitted to have one or more values, and of different value types.", p.structName(), p.propertyName()).Line()
	return p.elementTypeGenerator().def().Line().Line().Add(
		comment.Type().Id(p.structName()).Index().Id(p.iteratorTypeName().CamelName),
	)
}

func (p *NonFunctionalPropertyGenerator) funcs() *jen.Statement {
	funcs := p.elementTypeGenerator().funcs()
	memberFunc := jen.Func().Params(
		jen.Id("t").Id(p.structName()),
	)
	memberPtrFunc := jen.Func().Params(
		jen.Id("t").Op("*").Id(p.structName()),
	)
	prepend := make([]*jen.Statement, 0, len(p.Kinds))
	appendFns := make([]*jen.Statement, 0, len(p.Kinds))
	for i, kind := range p.Kinds {
		dict := jen.Dict{
			jen.Id(p.memberName(i)): jen.Id("v"),
		}
		if !kind.Nilable {
			dict[jen.Id(p.hasMemberName(i))] = jen.True()
		}
		prependMethodName := fmt.Sprintf("%s%s", prependMethod, p.kindCamelName(i))
		prepend = append(prepend, jen.Commentf(
			"%s prepends a %s value to the front of a list of the property %q.", prependMethodName, kind.ConcreteKind, p.propertyName(),
		).Line().Add(memberPtrFunc.Clone().Id(
			prependMethodName,
		).Params(
			jen.Id("v").Id(kind.ConcreteKind),
		).Block(
			jen.Op("*").Id("t").Op("=").Append(
				jen.Index().Id(p.iteratorTypeName().CamelName).Values(
					jen.Values(dict),
				),
				jen.Op("*").Id("t").Op("..."),
			),
		)))
		appendMethodName := fmt.Sprintf("%s%s", appendMethod, p.kindCamelName(i))
		appendFns = append(appendFns, jen.Commentf(
			"%s appends a %s value to the back of a list of the property %q", appendMethodName, kind.ConcreteKind, p.propertyName(),
		).Line().Add(memberPtrFunc.Clone().Id(
			appendMethodName,
		).Params(
			jen.Id("v").Id(kind.ConcreteKind),
		).Block(
			jen.Op("*").Id("t").Op("=").Append(
				jen.Op("*").Id("t"),
				jen.Id(p.iteratorTypeName().CamelName).Values(
					dict,
				),
			),
		)))
	}
	remove := jen.Commentf(
		"%s deletes an element at the specified index from a list of the property %q, regardless of its type.", removeMethod, p.propertyName(),
	).Line().Add(memberPtrFunc.Clone().Id(removeMethod).Params(
		jen.Id("idx").Int(),
	).Block(
		jen.Copy(
			jen.Parens(
				jen.Op("*").Id("t"),
			).Index(
				jen.Id("idx"),
				jen.Empty(),
			),
			jen.Parens(
				jen.Op("*").Id("t"),
			).Index(
				jen.Id("idx").Op("+").Lit(1),
				jen.Empty(),
			),
		),
		jen.Parens(
			jen.Op("*").Id("t"),
		).Index(
			jen.Len(jen.Op("*").Id("t")).Op("-").Lit(1),
		).Op("=").Id(p.iteratorTypeName().CamelName).Values(),
		jen.Op("*").Id("t").Op("=").Parens(
			jen.Op("*").Id("t"),
		).Index(
			jen.Empty(),
			jen.Len(jen.Op("*").Id("t")).Op("-").Lit(1),
		),
	))
	lenFn := jen.Commentf(
		"%s returns the number of values that exist for the %s property.", lenMethod, p.propertyName(),
	).Line().Add(memberFunc.Clone().Id(lenMethod).Params().Params(
		jen.Id("length").Int(),
	).Block(
		jen.Return(
			jen.Len(
				jen.Id("t"),
			),
		),
	))
	return funcs.Line().Line().Add(
		join(appendFns),
	).Line().Line().Add(
		join(prepend),
	).Line().Line().Add(
		remove,
	).Line().Line().Add(
		lenFn,
	).Line().Line().Add(
		p.serializationFuncs(),
	)
}

func (p *NonFunctionalPropertyGenerator) serializationFuncs() *jen.Statement {
	serialize := jen.Commentf(
		"%s converts this into an interface representation suitable for marshalling into a text or binary format.", p.serializeFnName(),
	).Line().Add(jen.Func().Params(
		jen.Id("t").Id(p.structName()),
	).Id(p.serializeFnName()).Params().Params(
		jen.Interface(),
		jen.Error(),
	).Block(
		jen.Id("s").Op(":=").Make(
			jen.Index().Interface(),
			jen.Lit(0),
			jen.Len(jen.Id("t")),
		),
		jen.For(
			jen.List(
				jen.Id("_"),
				jen.Id("iterator"),
			).Op(":=").Range().Id("t"),
		).Block(
			jen.If(
				jen.List(
					jen.Id("b"),
					jen.Err(),
				).Op(":=").Id("iterator").Dot(serializeIteratorMethod).Call(),
				jen.Err().Op("!=").Nil(),
			).Block(
				jen.Return(
					jen.Id("s"),
					jen.Err(),
				),
			).Else().Block(
				jen.Id("s").Op("=").Append(
					jen.Id("s"),
					jen.Id("b"),
				),
			),
		),
		jen.Return(
			jen.Id("s"),
			jen.Nil(),
		),
	))
	deserializeFn := func(variable string) jen.Code {
		return jen.If(
			jen.List(
				jen.Id("p"),
				jen.Err(),
			).Op(":=").Id(p.elementTypeGenerator().deserializeFnName()).Call(
				jen.Id(variable),
			),
			jen.Err().Op("!=").Nil(),
		).Block(
			jen.Return(
				jen.Id("t"),
				jen.Err(),
			),
		).Else().If(
			jen.Id("p").Op("!=").Nil(),
		).Block(
			jen.Id("t").Op("=").Append(
				jen.Id("t"),
				jen.Op("*").Id("p"),
			),
		)
	}
	deserialize := jen.Commentf(
		"%s creates a %q property from an interface representation that has been unmarshalled from a text or binary format.", p.deserializeFnName(), p.propertyName(),
	).Line().Add(jen.Func().Id(
		p.deserializeFnName(),
	).Params(
		jen.Id("m").Map(jen.String()).Interface(),
	).Params(
		jen.Id(p.structName()),
		jen.Error(),
	).Block(
		jen.Var().Id("t").Index().Id(p.iteratorTypeName().CamelName),
		jen.If(
			jen.List(
				jen.Id("i"),
				jen.Id("ok"),
			).Op(":=").Id("m").Index(
				jen.Lit(p.propertyName()),
			),
			jen.Id("ok"),
		).Block(
			jen.If(
				jen.List(
					jen.Id("list"),
					jen.Id("ok"),
				).Op(":=").Id("i").Assert(
					jen.Index().Interface(),
				),
				jen.Id("ok"),
			).Block(
				jen.For(
					jen.List(
						jen.Id("_"),
						jen.Id("iterator"),
					).Op(":=").Range().Id("list"),
				).Block(
					deserializeFn("iterator"),
				),
			).Else().Block(
				deserializeFn("i"),
			),
		),
		jen.Return(
			jen.Id("t"),
			jen.Nil(),
		),
	))
	return serialize.Line().Line().Add(deserialize)
}

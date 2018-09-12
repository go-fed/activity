package exp

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

// TODO: Auto-generate documentation and comments.
// TODO: Break out into multiple files.
// TODO: Generate Serialize and Deserialize functions for multi-type properties.

const (
	getMethod           = "Get"
	setMethod           = "Set"
	hasMethod           = "Has"
	clearMethod         = "Clear"
	iteratorClearMethod = "clear"
	isMethod            = "Is"
	appendMethod        = "Append"
	prependMethod       = "Prepend"
	removeMethod        = "Remove"
	lenMethod           = "Len"
	serializeMethod     = "Serialize"
	deserializeMethod   = "Deserialize"
)

func join(s []*jen.Statement) *jen.Statement {
	r := &jen.Statement{}
	for i, stmt := range s {
		if i > 0 {
			r.Line()
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
	return fmt.Sprintf("%s%s", deserializeMethod, p.Name.CamelName)
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

type FunctionalPropertyGenerator struct {
	PropertyGenerator
}

func (p *FunctionalPropertyGenerator) Generate() *jen.Statement {
	return p.def().Line().Add(p.funcs())
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
		return p.singleTypeFuncs().Line().Add(p.serializationFuncs())
	} else {
		return p.multiTypeFuncs().Line().Add(p.serializationFuncs())
	}
}

func (p *FunctionalPropertyGenerator) serializationFuncs() *jen.Statement {
	serializeFns := jen.Empty()
	for i, kind := range p.Kinds {
		if i > 0 {
			serializeFns = serializeFns.Else()
		}
		serializeFns = serializeFns.If(
			jen.Id("t").Dot(hasMethod).Call(),
		).Block(
			jen.List(
				jen.Id("b"),
				jen.Err(),
			).Op(":=").Id(kind.SerializeFnName).Call(
				jen.Id("t").Dot(getMethod).Call(),
			),
			jen.Return(
				jen.Lit(p.propertyName()),
				jen.Id("b"),
				jen.Err(),
			),
		)
	}
	serialize := jen.Func().Params(
		jen.Id("t").Id(p.structName()),
	).Id(serializeMethod).Params().Params(
		jen.String(),
		jen.Interface(),
		jen.Error(),
	).Block(
		serializeFns,
		jen.Return(
			jen.Lit(""),
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
	deserialize := jen.Func().Id(
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
	return serialize.Line().Add(deserialize)
}

func (p *FunctionalPropertyGenerator) singleTypeDef() *jen.Statement {
	if p.Kinds[0].Nilable {
		return jen.Type().Id(p.structName()).Struct(
			jen.Id(p.memberName(0)).Id(p.Kinds[0].ConcreteKind),
		)
	} else {
		return jen.Type().Id(p.structName()).Struct(
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
	var has *jen.Statement
	var get *jen.Statement
	var set *jen.Statement
	var clear *jen.Statement
	if p.Kinds[0].Nilable {
		has = memberFunc.Clone().Id(hasMethod).Params().Params(
			jen.Bool(),
		).Block(
			jen.Return(jen.Id("t").Dot(p.memberName(0)).Op("!=").Nil()),
		)
		get = memberFunc.Clone().Id(getMethod).Params().Params(
			jen.Id(p.Kinds[0].ConcreteKind),
		).Block(
			jen.Return(jen.Id("t").Dot(p.memberName(0))),
		)
		set = memberPtrFunc.Clone().Id(setMethod).Params(
			jen.Id("v").Id(p.Kinds[0].ConcreteKind),
		).Block(
			jen.Id("t").Dot(p.memberName(0)).Op("=").Id("v"),
		)
		clear = memberPtrFunc.Clone().Id(p.clearMethodName()).Params().Block(
			jen.Id("t").Dot(p.memberName(0)).Op("=").Nil(),
		)
	} else {
		has = memberFunc.Clone().Id(hasMethod).Params().Params(
			jen.Bool(),
		).Block(
			jen.Return(jen.Id("t").Dot(p.hasMemberName(0))),
		)
		get = memberFunc.Clone().Id(getMethod).Params().Params(
			jen.Id(p.Kinds[0].ConcreteKind),
		).Block(
			jen.Return(jen.Id("t").Dot(p.memberName(0))),
		)
		set = memberPtrFunc.Clone().Id(setMethod).Params(
			jen.Id("v").Id(p.Kinds[0].ConcreteKind),
		).Block(
			jen.Id("t").Dot(p.memberName(0)).Op("=").Id("v"),
			jen.Id("t").Dot(p.hasMemberName(0)).Op("=").True(),
		)
		clear = memberPtrFunc.Clone().Id(p.clearMethodName()).Params().Block(
			jen.Id("t").Dot(p.hasMemberName(0)).Op("=").False(),
		)
	}
	return has.Line().Add(get).Line().Add(set).Line().Add(clear)
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
	return jen.Type().Id(p.structName()).Struct(join(kindMembers))
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
		var is *jen.Statement
		var get *jen.Statement
		var set *jen.Statement
		if kind.Nilable {
			is = memberFunc.Clone().Id(
				fmt.Sprintf("%s%s", isMethod, p.kindCamelName(i)),
			).Params().Params(
				jen.Bool(),
			).Block(
				jen.Return(jen.Id("t").Dot(p.memberName(i)).Op("!=").Nil()),
			)
			set = memberPtrFunc.Clone().Id(
				fmt.Sprintf("%s%s", setMethod, p.kindCamelName(i)),
			).Params(
				jen.Id("v").Id(kind.ConcreteKind),
			).Block(
				jen.Id("t").Dot(p.clearMethodName()).Call(),
				jen.Id("t").Dot(p.memberName(i)).Op("=").Id("v"),
			)
		} else {
			is = memberFunc.Clone().Id(
				fmt.Sprintf("%s%s", isMethod, p.kindCamelName(i)),
			).Params().Params(
				jen.Bool(),
			).Block(
				jen.Return(jen.Id("t").Dot(p.hasMemberName(i))),
			)
			set = memberPtrFunc.Clone().Id(
				fmt.Sprintf("%s%s", setMethod, p.kindCamelName(i)),
			).Params(
				jen.Id("v").Id(kind.ConcreteKind),
			).Block(
				jen.Id("t").Dot(p.clearMethodName()).Call(),
				jen.Id("t").Dot(p.memberName(i)).Op("=").Id("v"),
				jen.Id("t").Dot(p.hasMemberName(i)).Op("=").True(),
			)
		}
		get = memberFunc.Clone().Id(
			fmt.Sprintf("%s%s", getMethod, p.kindCamelName(i)),
		).Params().Params(
			jen.Id(kind.ConcreteKind),
		).Block(
			jen.Return(jen.Id("t").Dot(p.memberName(i))),
		)
		funcs.Add(is).Line()
		funcs.Add(get).Line()
		funcs.Add(set).Line()
	}
	clear := memberPtrFunc.Clone().Id(p.clearMethodName()).Params().Block(
		join(clearLine),
	)
	hasAny := memberFunc.Clone().Id(
		fmt.Sprintf("%sAny", hasMethod),
	).Params().Params(
		jen.Bool(),
	).Block(
		jen.Return(join(isLine)),
	)
	funcs.Add(clear).Line().Add(hasAny).Line()
	return funcs
}

type NonFunctionalPropertyGenerator struct {
	PropertyGenerator
}

func (p *NonFunctionalPropertyGenerator) Generate() *jen.Statement {
	return p.def().Line().Add(p.funcs())
}

func (p *PropertyGenerator) typeName() string {
	return fmt.Sprintf("%sProperty", p.Name.CamelName)
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
	return p.elementTypeGenerator().def().Line().Add(
		jen.Type().Id(p.typeName()).Index().Id(p.iteratorTypeName().CamelName),
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
		prepend = append(prepend, memberPtrFunc.Clone().Id(
			fmt.Sprintf("%s%s", prependMethod, p.kindCamelName(i)),
		).Params(
			jen.Id("v").Id(kind.ConcreteKind),
		).Block(
			jen.Op("*").Id("t").Op("=").Append(
				jen.Index().Id(p.iteratorTypeName().CamelName).Values(
					jen.Values(dict),
				),
				jen.Op("*").Id("t").Op("..."),
			),
		))
		appendFns = append(appendFns, memberPtrFunc.Clone().Id(
			fmt.Sprintf("%s%s", appendMethod, p.kindCamelName(i)),
		).Params(
			jen.Id("v").Id(kind.ConcreteKind),
		).Block(
			jen.Op("*").Id("t").Op("=").Append(
				jen.Op("*").Id("t"),
				jen.Id(p.iteratorTypeName().CamelName).Values(
					dict,
				),
			),
		))
	}
	remove := memberPtrFunc.Clone().Id(removeMethod).Params(
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
	)
	lenFn := memberFunc.Clone().Id(lenMethod).Params().Params(
		jen.Id("length").Int(),
	).Block(
		jen.Return(
			jen.Len(
				jen.Id("t"),
			),
		),
	)
	return funcs.Line().Add(join(appendFns)).Line().Add(join(prepend)).Line().Add(remove).Line().Add(lenFn)
}

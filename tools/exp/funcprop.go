package exp

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

type FunctionalPropertyGenerator struct {
	PropertyGenerator
}

func (p *FunctionalPropertyGenerator) Definition() jen.Code {
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
	memberFunc := jen.Func().Params(
		jen.Id("t").Id(p.structName()),
	)
	kindIndexFns := make([]*jen.Statement, 0, len(p.Kinds))
	for i, _ := range p.Kinds {
		if len(p.Kinds) > 1 {
			kindIndexFns = append(kindIndexFns, jen.If(
				jen.Id("t").Dot(p.isMethodName(i)).Call(),
			).Block(
				jen.Return(jen.Lit(i)),
			))
		} else {
			kindIndexFns = append(kindIndexFns, jen.If(
				jen.Id("t").Dot(hasMethod).Call(),
			).Block(
				jen.Return(jen.Lit(i)),
			))
		}
	}
	kindIndexFn := jen.Commentf(
		"%s computes an arbitrary value for indexing this kind of value.", kindIndexMethod,
	).Line().Add(memberFunc.Clone().Id(kindIndexMethod).Params().Int().Block(
		joinSingleLine(kindIndexFns),
		jen.Return(jen.Lit(-1)),
	))
	funcs := jen.Empty()
	funcs.Add(
		p.serializationFuncs(),
	).Line().Line().Add(
		kindIndexFn,
	).Line().Line()
	if len(p.Kinds) == 1 {
		return funcs.Add(p.singleTypeFuncs())
	} else {
		return funcs.Add(p.multiTypeFuncs())
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
	return comment.Type().Id(p.structName()).Struct(joinSingleLine(kindMembers))
}

func (p *FunctionalPropertyGenerator) multiTypeFuncs() *jen.Statement {
	funcs := jen.Empty()
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
		setMethodName := fmt.Sprintf("%s%s", setMethod, p.kindCamelName(i))
		is := jen.Commentf("%s returns true if this property has a type of value of %q.", p.isMethodName(i), kind.ConcreteKind).Line()
		get := jen.Commentf("%s returns the value of this property. When %s returns false, %s will return an arbitrary value.", p.getFnName(i), p.isMethodName(i), p.getFnName(i)).Line()
		set := jen.Commentf("%s sets the value of this property. Calling %s afterwards returns true.", setMethodName, p.isMethodName(i)).Line()
		if kind.Nilable {
			is = is.Add(memberFunc.Clone().Id(
				p.isMethodName(i),
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
				p.isMethodName(i),
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
		joinSingleLine(clearLine),
	))
	hasAny := jen.Commentf(
		"%s returns true if any of the different values is set.", hasAnyMethodName,
	).Line().Add(memberFunc.Clone().Id(
		hasAnyMethodName,
	).Params().Params(
		jen.Bool(),
	).Block(
		jen.Return(joinSingleLine(isLine)),
	))
	funcs.Add(clear).Line().Line().Add(hasAny)
	return funcs
}

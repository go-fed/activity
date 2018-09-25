package exp

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

type NonFunctionalPropertyGenerator struct {
	PropertyGenerator
}

func (p *NonFunctionalPropertyGenerator) Definition() jen.Code {
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
	less := jen.Empty()
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
		if i > 0 {
			less.Else()
		}
		less.If(
			jen.Id("idx1").Op("==").Lit(i),
		).Block(
			jen.Id("lhs").Op(":=").Id("t").Index(jen.Id("i")).Dot(p.getFnName(i)).Call(),
			jen.Id("rhs").Op(":=").Id("t").Index(jen.Id("j")).Dot(p.getFnName(i)).Call(),
			jen.Return(jen.Id(kind.LessFnName).Call(
				jen.Id("lhs"),
				jen.Id("rhs"),
			)),
		)
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
		"%s returns the number of values that exist for the %q property.", lenMethod, p.propertyName(),
	).Line().Add(memberFunc.Clone().Id(lenMethod).Params().Params(
		jen.Id("length").Int(),
	).Block(
		jen.Return(
			jen.Len(
				jen.Id("t"),
			),
		),
	))
	swapFn := jen.Commentf(
		"%s swaps the location of values at two indices for the %q property.", swapMethod, p.propertyName(),
	).Line().Add(memberFunc.Clone().Id(swapMethod).Params(
		jen.Id("i"),
		jen.Id("j").Int(),
	).Params().Block(
		jen.List(
			jen.Id("t").Index(jen.Id("i")),
			jen.Id("t").Index(jen.Id("j")),
		).Op("=").List(
			jen.Id("t").Index(jen.Id("j")),
			jen.Id("t").Index(jen.Id("i")),
		),
	))
	lessFn := jen.Commentf(
		"%s computes whether another property is less than this one. Mixing types results in a consistent but arbitrary ordering", lessMethod,
	).Line().Add(memberFunc.Clone().Id(lessMethod).Params(
		jen.Id("i"),
		jen.Id("j").Int(),
	).Bool().Block(
		jen.Id("idx1").Op(":=").Id("t").Dot(kindIndexMethod).Call(jen.Id("i")),
		jen.Id("idx2").Op(":=").Id("t").Dot(kindIndexMethod).Call(jen.Id("j")),
		jen.If(jen.Id("idx1").Op("<").Id("idx2")).Block(
			jen.Return(jen.True()),
		).Else().If(jen.Id("idx1").Op("==").Id("idx2")).Block(
			less,
		),
		jen.Return(jen.False()),
	))
	kindIndexFn := jen.Commentf(
		"%s computes an arbitrary value for indexing this kind of value.", kindIndexMethod,
	).Line().Add(memberFunc.Clone().Id(kindIndexMethod).Params(jen.Id("idx").Int()).Int().Block(
		jen.Return(
			jen.Id("t").Index(jen.Id("idx")).Dot(kindIndexMethod).Call(),
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
		swapFn,
	).Line().Line().Add(
		lessFn,
	).Line().Line().Add(
		kindIndexFn,
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

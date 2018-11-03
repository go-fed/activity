package props

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/go-fed/activity/tools/exp/codegen"
)

// NonFunctionalPropertyGenerator produces Go code for properties that can have
// more than one value. The resulting property is a type that is a list of
// iterators; each iterator is a concrete struct type. The property can be
// sorted and iterated over so individual elements can be inspected.
type NonFunctionalPropertyGenerator struct {
	PropertyGenerator
}

// Definitions produces the Go code definitions, which can generate their Go
// implementations. The struct is the iterator for various values of the
// property, which is defined by the type definition.
func (p *NonFunctionalPropertyGenerator) Definitions() (*codegen.Struct, *codegen.Typedef) {
	methods, funcs := p.serializationFuncs()
	methods = append(methods, p.funcs()...)
	property := codegen.NewTypedef(
		jen.Commentf("%s is the non-functional property %q. It is permitted to have one or more values, and of different value types.", p.StructName(), p.PropertyName()),
		p.StructName(),
		jen.Index().Id(p.iteratorTypeName().CamelName),
		methods,
		funcs)
	iterator := p.elementTypeGenerator().Definition()
	return iterator, property
}

// iteratorTypeName determines the identifier to use for the iterator type.
func (p *NonFunctionalPropertyGenerator) iteratorTypeName() Identifier {
	return Identifier{
		LowerName: fmt.Sprintf("%sPropertyIterator", p.Name.LowerName),
		CamelName: fmt.Sprintf("%sPropertyIterator", p.Name.CamelName),
	}
}

// elementTypeGenerator produces a FunctionalPropertyGenerator for the iterator
// type.
func (p *NonFunctionalPropertyGenerator) elementTypeGenerator() *FunctionalPropertyGenerator {
	return &FunctionalPropertyGenerator{
		PropertyGenerator{
			Package: p.PropertyGenerator.Package,
			Name:    p.iteratorTypeName(),
			Kinds:   p.Kinds,
			HasNaturalLanguageMap: p.PropertyGenerator.HasNaturalLanguageMap,
			asIterator:            true,
		},
	}
}

// funcs produces the methods needed for the NonFunctional property.
func (p *NonFunctionalPropertyGenerator) funcs() []*codegen.Method {
	var methods []*codegen.Method
	less := jen.Empty()
	for i, kind := range p.Kinds {
		dict := jen.Dict{
			jen.Id(p.memberName(i)): jen.Id("v"),
		}
		if !kind.Nilable {
			dict[jen.Id(p.hasMemberName(i))] = jen.True()
		}
		// Prepend Method
		prependMethodName := fmt.Sprintf("%s%s", prependMethod, p.kindCamelName(i))
		methods = append(methods,
			codegen.NewCommentedPointerMethod(
				p.packageName(),
				prependMethodName,
				p.StructName(),
				[]jen.Code{jen.Id("v").Id(kind.ConcreteKind)},
				/*ret=*/ nil,
				[]jen.Code{
					jen.Op("*").Id(codegen.This()).Op("=").Append(
						jen.Index().Id(p.iteratorTypeName().CamelName).Values(
							jen.Values(dict),
						),
						jen.Op("*").Id(codegen.This()).Op("..."),
					),
				},
				jen.Commentf("%s prepends a %s value to the front of a list of the property %q.", prependMethodName, kind.ConcreteKind, p.PropertyName())))
		// Append Method
		appendMethodName := fmt.Sprintf("%s%s", appendMethod, p.kindCamelName(i))
		methods = append(methods,
			codegen.NewCommentedPointerMethod(
				p.packageName(),
				appendMethodName,
				p.StructName(),
				[]jen.Code{jen.Id("v").Id(kind.ConcreteKind)},
				/*ret=*/ nil,
				[]jen.Code{
					jen.Op("*").Id(codegen.This()).Op("=").Append(
						jen.Op("*").Id(codegen.This()),
						jen.Id(p.iteratorTypeName().CamelName).Values(
							dict,
						),
					),
				},
				jen.Commentf("%s appends a %s value to the back of a list of the property %q", appendMethodName, kind.ConcreteKind, p.PropertyName())))
		// Less logic
		if i > 0 {
			less.Else()
		}
		less.If(
			jen.Id("idx1").Op("==").Lit(i),
		).Block(
			jen.Id("lhs").Op(":=").Id(codegen.This()).Index(jen.Id("i")).Dot(p.getFnName(i)).Call(),
			jen.Id("rhs").Op(":=").Id(codegen.This()).Index(jen.Id("j")).Dot(p.getFnName(i)).Call(),
			jen.Return(kind.LessFn.Call(
				jen.Id("lhs"),
				jen.Id("rhs"),
			)),
		)
	}
	// Remove Method
	methods = append(methods,
		codegen.NewCommentedPointerMethod(
			p.packageName(),
			removeMethod,
			p.StructName(),
			[]jen.Code{jen.Id("idx").Int()},
			/*ret=*/ nil,
			[]jen.Code{
				jen.Copy(
					jen.Parens(
						jen.Op("*").Id(codegen.This()),
					).Index(
						jen.Id("idx"),
						jen.Empty(),
					),
					jen.Parens(
						jen.Op("*").Id(codegen.This()),
					).Index(
						jen.Id("idx").Op("+").Lit(1),
						jen.Empty(),
					),
				),
				jen.Parens(
					jen.Op("*").Id(codegen.This()),
				).Index(
					jen.Len(jen.Op("*").Id(codegen.This())).Op("-").Lit(1),
				).Op("=").Id(p.iteratorTypeName().CamelName).Values(),
				jen.Op("*").Id(codegen.This()).Op("=").Parens(
					jen.Op("*").Id(codegen.This()),
				).Index(
					jen.Empty(),
					jen.Len(jen.Op("*").Id(codegen.This())).Op("-").Lit(1),
				),
			},
			jen.Commentf("%s deletes an element at the specified index from a list of the property %q, regardless of its type.", removeMethod, p.PropertyName())))
	// Len Method
	methods = append(methods,
		codegen.NewCommentedValueMethod(
			p.packageName(),
			lenMethod,
			p.StructName(),
			/*params=*/ nil,
			[]jen.Code{jen.Id("length").Int()},
			[]jen.Code{
				jen.Return(
					jen.Len(
						jen.Id(codegen.This()),
					),
				),
			},
			jen.Commentf("%s returns the number of values that exist for the %q property.", lenMethod, p.PropertyName())))
	// Swap Method
	methods = append(methods,
		codegen.NewCommentedValueMethod(
			p.packageName(),
			swapMethod,
			p.StructName(),
			[]jen.Code{
				jen.Id("i"),
				jen.Id("j").Int(),
			},
			/*ret=*/ nil,
			[]jen.Code{
				jen.List(
					jen.Id(codegen.This()).Index(jen.Id("i")),
					jen.Id(codegen.This()).Index(jen.Id("j")),
				).Op("=").List(
					jen.Id(codegen.This()).Index(jen.Id("j")),
					jen.Id(codegen.This()).Index(jen.Id("i")),
				),
			},
			jen.Commentf("%s swaps the location of values at two indices for the %q property.", swapMethod, p.PropertyName())))
	// Less Method
	methods = append(methods,
		codegen.NewCommentedValueMethod(
			p.packageName(),
			lessMethod,
			p.StructName(),
			[]jen.Code{
				jen.Id("i"),
				jen.Id("j").Int(),
			},
			[]jen.Code{jen.Bool()},
			[]jen.Code{
				jen.Id("idx1").Op(":=").Id(codegen.This()).Dot(kindIndexMethod).Call(jen.Id("i")),
				jen.Id("idx2").Op(":=").Id(codegen.This()).Dot(kindIndexMethod).Call(jen.Id("j")),
				jen.If(jen.Id("idx1").Op("<").Id("idx2")).Block(
					jen.Return(jen.True()),
				).Else().If(jen.Id("idx1").Op("==").Id("idx2")).Block(
					less,
				),
				jen.Return(jen.False()),
			},
			jen.Commentf("%s computes whether another property is less than this one. Mixing types results in a consistent but arbitrary ordering", lessMethod)))
	// Kind Method
	methods = append(methods,
		codegen.NewCommentedValueMethod(
			p.packageName(),
			kindIndexMethod,
			p.StructName(),
			[]jen.Code{jen.Id("idx").Int()},
			[]jen.Code{jen.Int()},
			[]jen.Code{
				jen.Return(
					jen.Id(codegen.This()).Index(jen.Id("idx")).Dot(kindIndexMethod).Call(),
				),
			},
			jen.Commentf("%s computes an arbitrary value for indexing this kind of value.", kindIndexMethod)))
	return methods
}

// serializationFuncs produces the Methods and Functions needed for a
// NonFunctional property to be serialized and deserialized to and from an
// encoding.
func (p *NonFunctionalPropertyGenerator) serializationFuncs() ([]*codegen.Method, []*codegen.Function) {
	serialize := []*codegen.Method{
		codegen.NewCommentedValueMethod(
			p.packageName(),
			p.serializeFnName(),
			p.StructName(),
			/*params=*/ nil,
			[]jen.Code{jen.Interface(), jen.Error()},
			[]jen.Code{
				jen.Id("s").Op(":=").Make(
					jen.Index().Interface(),
					jen.Lit(0),
					jen.Len(jen.Id(codegen.This())),
				),
				jen.For(
					jen.List(
						jen.Id("_"),
						jen.Id("iterator"),
					).Op(":=").Range().Id(codegen.This()),
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
			},
			jen.Commentf("%s converts this into an interface representation suitable for marshalling into a text or binary format.", p.serializeFnName()),
		),
	}
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
				jen.Id(codegen.This()),
				jen.Err(),
			),
		).Else().If(
			jen.Id("p").Op("!=").Nil(),
		).Block(
			jen.Id(codegen.This()).Op("=").Append(
				jen.Id(codegen.This()),
				jen.Op("*").Id("p"),
			),
		)
	}
	deserialize := []*codegen.Function{
		codegen.NewCommentedFunction(
			p.packageName(),
			p.deserializeFnName(),
			[]jen.Code{jen.Id("m").Map(jen.String()).Interface()},
			[]jen.Code{jen.Id(p.StructName()), jen.Error()},
			[]jen.Code{
				jen.Var().Id(codegen.This()).Index().Id(p.iteratorTypeName().CamelName),
				jen.If(
					jen.List(
						jen.Id("i"),
						jen.Id("ok"),
					).Op(":=").Id("m").Index(
						jen.Lit(p.PropertyName()),
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
					jen.Id(codegen.This()),
					jen.Nil(),
				),
			},
			jen.Commentf("%s creates a %q property from an interface representation that has been unmarshalled from a text or binary format.", p.deserializeFnName(), p.PropertyName()),
		),
	}
	return serialize, deserialize
}

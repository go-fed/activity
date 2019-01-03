package props

import (
	"fmt"
	"github.com/cjslep/activity/tools/exp/codegen"
	"github.com/dave/jennifer/jen"
	"sync"
)

// NonFunctionalPropertyGenerator produces Go code for properties that can have
// more than one value. The resulting property is a type that is a list of
// iterators; each iterator is a concrete struct type. The property can be
// sorted and iterated over so individual elements can be inspected.
type NonFunctionalPropertyGenerator struct {
	PropertyGenerator
	cacheOnce     sync.Once
	cachedStruct  *codegen.Struct
	cachedTypedef *codegen.Typedef
}

// NewNonFunctionalPropertyGenerator is a convenience constructor to create
// NonFunctionalPropertyGenerators.
//
// PropertyGenerators shoulf be in the first pass to construct, before types and
// other generators are constructed.
func NewNonFunctionalPropertyGenerator(pm *PackageManager,
	name Identifier,
	kinds []Kind,
	hasNaturalLanguageMap bool) *NonFunctionalPropertyGenerator {
	return &NonFunctionalPropertyGenerator{
		PropertyGenerator: PropertyGenerator{
			PackageManager:        pm,
			HasNaturalLanguageMap: hasNaturalLanguageMap,
			Name:                  name,
			Kinds:                 kinds,
		},
	}
}

// toInterfaces creates the interface versions of the definitions generated.
//
// TODO: Remove -- redundant?
func (p *NonFunctionalPropertyGenerator) toInterfaces(pkg Package) []*codegen.Interface {
	s, t := p.Definitions()
	return []*codegen.Interface{
		s.ToInterface(pkg.Path(), p.elementTypeGenerator().InterfaceName(), ""),
		t.ToInterface(pkg.Path(), p.InterfaceName(), ""),
	}
}

// InterfaceDefinitions creates interface definitions in the provided package.
func (p *NonFunctionalPropertyGenerator) InterfaceDefinitions(pkg Package) []*codegen.Interface {
	s, t := p.Definitions()
	return []*codegen.Interface{
		// TODO: Comments
		s.ToInterface(pkg.Path(), p.elementTypeGenerator().InterfaceName(), ""),
		t.ToInterface(pkg.Path(), p.InterfaceName(), p.Comments()),
	}
}

// Definitions produces the Go code definitions, which can generate their Go
// implementations. The struct is the iterator for various values of the
// property, which is defined by the type definition.
func (p *NonFunctionalPropertyGenerator) Definitions() (*codegen.Struct, *codegen.Typedef) {
	p.cacheOnce.Do(func() {
		var methods []*codegen.Method
		var funcs []*codegen.Function
		ser, deser := p.serializationFuncs()
		methods = append(methods, ser)
		funcs = append(funcs, deser)
		methods = append(methods, p.funcs()...)
		property := codegen.NewTypedef(
			jen.Commentf("%s is the non-functional property %q. It is permitted to have one or more values, and of different value types.", p.StructName(), p.PropertyName()),
			p.StructName(),
			jen.Index().Id(p.iteratorTypeName().CamelName),
			methods,
			funcs)
		iterator := p.elementTypeGenerator().Definition()
		p.cachedStruct, p.cachedTypedef = iterator, property
	})
	return p.cachedStruct, p.cachedTypedef
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
		PropertyGenerator: PropertyGenerator{
			PackageManager:        p.PropertyGenerator.PackageManager,
			Name:                  p.iteratorTypeName(),
			Kinds:                 p.Kinds,
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
				p.GetPrivatePackage().Path(),
				prependMethodName,
				p.StructName(),
				[]jen.Code{jen.Id("v").Add(kind.ConcreteKind)},
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
				p.GetPrivatePackage().Path(),
				appendMethodName,
				p.StructName(),
				[]jen.Code{jen.Id("v").Add(kind.ConcreteKind)},
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
		// LessFn is nil case -- call comparison Less method directly on the LHS.
		// TODO: Move this logic to a Kind method (see funcprop.go too)
		lessCall := jen.Id("lhs").Dot(compareLessMethod).Call(jen.Id("rhs"))
		if kind.LessFn != nil {
			// LessFn is indeed a function -- call this function
			lessCall = kind.LessFn.Clone().Call(
				jen.Id("lhs"),
				jen.Id("rhs"),
			)
		}
		less.If(
			jen.Id("idx1").Op("==").Lit(i),
		).Block(
			jen.Id("lhs").Op(":=").Id(codegen.This()).Index(jen.Id("i")).Dot(p.getFnName(i)).Call(),
			jen.Id("rhs").Op(":=").Id(codegen.This()).Index(jen.Id("j")).Dot(p.getFnName(i)).Call(),
			jen.Return(lessCall),
		)
	}
	// Remove Method
	methods = append(methods,
		codegen.NewCommentedPointerMethod(
			p.GetPrivatePackage().Path(),
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
			p.GetPrivatePackage().Path(),
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
			p.GetPrivatePackage().Path(),
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
			p.GetPrivatePackage().Path(),
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
			p.GetPrivatePackage().Path(),
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
	// LessThan Method
	lessCode := jen.Empty().Add(
		jen.Id("l1").Op(":=").Id(codegen.This()).Dot(lenMethod).Call().Line(),
		jen.Id("l2").Op(":=").Id("o").Dot(lenMethod).Call().Line(),
		jen.Id("l").Op(":=").Id("l1").Line(),
		jen.If(
			jen.Id("l2").Op("<").Id("l1"),
		).Block(
			jen.Id("l").Op("=").Id("l2"),
		))
	methods = append(methods, codegen.NewCommentedValueMethod(
		p.GetPrivatePackage().Path(),
		compareLessMethod,
		p.StructName(),
		[]jen.Code{jen.Id("o").Id(p.InterfaceName())},
		[]jen.Code{jen.Bool()},
		[]jen.Code{
			lessCode,
			jen.For(
				jen.Id("i").Op(":=").Lit(0),
				jen.Id("i").Op("<").Id("l"),
				jen.Id("i").Op("++"),
			).Block(
				jen.If(
					jen.Id(codegen.This()).Index(jen.Id("i")).Dot(compareLessMethod).Call(jen.Id("o").Index(jen.Id("i"))),
				).Block(
					jen.Return(jen.True()),
				).Else().If(
					jen.Id("o").Index(jen.Id("i")).Dot(compareLessMethod).Call(jen.Id(codegen.This()).Index(jen.Id("i"))),
				).Block(
					jen.Return(jen.False()),
				),
			),
			jen.Return(jen.Id("l1").Op("<").Id("l2")),
		},
		jen.Commentf("%s compares two instances of this property with an arbitrary but stable comparison.", compareLessMethod),
	))
	return methods
}

// serializationFuncs produces the Methods and Functions needed for a
// NonFunctional property to be serialized and deserialized to and from an
// encoding.
func (p *NonFunctionalPropertyGenerator) serializationFuncs() (*codegen.Method, *codegen.Function) {
	serialize := codegen.NewCommentedValueMethod(
		p.GetPrivatePackage().Path(),
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
		jen.Commentf("%s converts this into an interface representation suitable for marshalling into a text or binary format.", p.serializeFnName()))
	deserializeFn := func(variable string) jen.Code {
		return jen.If(
			jen.List(
				jen.Id("p"),
				jen.Err(),
			).Op(":=").Id(p.elementTypeGenerator().DeserializeFnName()).Call(
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
	deserialize := codegen.NewCommentedFunction(
		p.GetPrivatePackage().Path(),
		p.DeserializeFnName(),
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
		jen.Commentf("%s creates a %q property from an interface representation that has been unmarshalled from a text or binary format.", p.DeserializeFnName(), p.PropertyName()))
	return serialize, deserialize
}

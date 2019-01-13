package gen

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
func NewNonFunctionalPropertyGenerator(vocabName string,
	pm *PackageManager,
	name Identifier,
	comment string,
	kinds []Kind,
	hasNaturalLanguageMap bool) *NonFunctionalPropertyGenerator {
	return &NonFunctionalPropertyGenerator{
		PropertyGenerator: PropertyGenerator{
			vocabName:             vocabName,
			packageManager:        pm,
			hasNaturalLanguageMap: hasNaturalLanguageMap,
			name:                  name,
			comment:               comment,
			kinds:                 kinds,
		},
	}
}

// InterfaceDefinitions creates interface definitions in the provided package.
func (p *NonFunctionalPropertyGenerator) InterfaceDefinitions(pkg Package) []*codegen.Interface {
	s, t := p.Definitions()
	return []*codegen.Interface{
		s.ToInterface(pkg.Path(), p.elementTypeGenerator().InterfaceName(), fmt.Sprintf("%s represents a single value for the %q property.", p.elementTypeGenerator().InterfaceName(), p.PropertyName())),
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
			fmt.Sprintf("%s is the non-functional property %q. It is permitted to have one or more values, and of different value types.", p.StructName(), p.PropertyName()),
			p.StructName(),
			jen.Index().Op("*").Id(p.iteratorTypeName().CamelName),
			methods,
			funcs)
		iterator := p.elementTypeGenerator().Definition()
		p.cachedStruct, p.cachedTypedef = iterator, property
	})
	return p.cachedStruct, p.cachedTypedef
}

// iteratorInterfaceName gets the interface name for the iterator.
func (p *NonFunctionalPropertyGenerator) iteratorInterfaceName() string {
	return fmt.Sprintf("%sInterface", p.iteratorTypeName().CamelName)
}

// elementTypeGenerator produces a FunctionalPropertyGenerator for the iterator
// type.
func (p *NonFunctionalPropertyGenerator) elementTypeGenerator() *FunctionalPropertyGenerator {
	return &FunctionalPropertyGenerator{
		PropertyGenerator: PropertyGenerator{
			packageManager:        p.PropertyGenerator.packageManager,
			name:                  p.iteratorTypeName(),
			kinds:                 p.kinds,
			hasNaturalLanguageMap: p.PropertyGenerator.hasNaturalLanguageMap,
			asIterator:            true,
		},
	}
}

// funcs produces the methods needed for the NonFunctional property.
func (p *NonFunctionalPropertyGenerator) funcs() []*codegen.Method {
	var methods []*codegen.Method
	less := jen.Empty()
	for i, kind := range p.kinds {
		dict := jen.Dict{
			jen.Id(parentMemberName): jen.Id(codegen.This()),
			jen.Id(p.memberName(i)):  jen.Id("v"),
		}
		if !kind.Nilable {
			dict[jen.Id(p.hasMemberName(i))] = jen.True()
		}
		// Prepend Method
		prependDict := jen.Dict{}
		for k, v := range dict {
			prependDict[k] = v
		}
		prependDict[jen.Id(myIndexMemberName)] = jen.Lit(0)
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
						jen.Index().Op("*").Id(p.iteratorTypeName().CamelName).Values(
							jen.Values(prependDict),
						),
						jen.Op("*").Id(codegen.This()).Op("..."),
					),
					jen.For(
						jen.Id("i").Op(":=").Lit(1),
						jen.Id("i").Op("<").Id(codegen.This()).Dot(lenMethod).Call(),
						jen.Id("i").Op("++"),
					).Block(
						jen.Parens(
							jen.Op("*").Id(codegen.This()),
						).Index(jen.Id("i")).Dot(myIndexMemberName).Op("=").Id("i"),
					),
				},
				fmt.Sprintf("%s prepends a %s value to the front of a list of the property %q. Invalidates all iterators.", prependMethodName, kind.Name.LowerName, p.PropertyName())))
		// Append Method
		appendDict := jen.Dict{}
		for k, v := range dict {
			appendDict[k] = v
		}
		appendDict[jen.Id(myIndexMemberName)] = jen.Id(codegen.This()).Dot(lenMethod).Call()
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
						jen.Op("&").Id(p.iteratorTypeName().CamelName).Values(
							appendDict,
						),
					),
				},
				fmt.Sprintf("%s appends a %s value to the back of a list of the property %q. Invalidates iterators that are traversing using %s.", appendMethodName, kind.Name.LowerName, p.PropertyName(), prevMethod)))
		// Set Method
		setDict := jen.Dict{}
		for k, v := range dict {
			setDict[k] = v
		}
		setDict[jen.Id(myIndexMemberName)] = jen.Id("idx")
		setMethodName := p.setFnName(i)
		methods = append(methods,
			codegen.NewCommentedPointerMethod(
				p.GetPrivatePackage().Path(),
				setMethodName,
				p.StructName(),
				[]jen.Code{jen.Id("idx").Int(), jen.Id("v").Add(kind.ConcreteKind)},
				/*ret=*/ nil,
				[]jen.Code{
					jen.Parens(jen.Op("*").Id(codegen.This())).Index(jen.Id("idx")).Dot(parentMemberName).Op("=").Nil(),
					jen.Parens(jen.Op("*").Id(codegen.This())).Index(jen.Id("idx")).Op("=").Op("&").Id(p.iteratorTypeName().CamelName).Values(
						setDict,
					),
				},
				fmt.Sprintf("%s sets a %s value to be at the specified index for the property %q. Panics if the index is out of bounds. Invalidates all iterators.", setMethodName, kind.Name.LowerName, p.PropertyName())))
		// Less logic
		if i > 0 {
			less.Else()
		}
		lessCall := kind.lessFnCode(jen.Id("lhs"), jen.Id("rhs"))
		less.If(
			jen.Id("idx1").Op("==").Lit(i),
		).Block(
			jen.Id("lhs").Op(":=").Id(codegen.This()).Index(jen.Id("i")).Dot(p.getFnName(i)).Call(),
			jen.Id("rhs").Op(":=").Id(codegen.This()).Index(jen.Id("j")).Dot(p.getFnName(i)).Call(),
			jen.Return(lessCall),
		)
	}
	// IRI Prepend, Append, Set, and Less logic
	methods = append(methods,
		codegen.NewCommentedPointerMethod(
			p.GetPrivatePackage().Path(),
			fmt.Sprintf("%sIRI", prependMethod),
			p.StructName(),
			[]jen.Code{jen.Id("v").Op("*").Qual("net/url", "URL")},
			/*ret=*/ nil,
			[]jen.Code{
				jen.Op("*").Id(codegen.This()).Op("=").Append(
					jen.Index().Op("*").Id(p.iteratorTypeName().CamelName).Values(
						jen.Values(jen.Dict{
							jen.Id(iriMember):         jen.Id("v"),
							jen.Id(parentMemberName):  jen.Id(codegen.This()),
							jen.Id(myIndexMemberName): jen.Lit(0),
						}),
					),
					jen.Op("*").Id(codegen.This()).Op("..."),
				),
				jen.For(
					jen.Id("i").Op(":=").Lit(1),
					jen.Id("i").Op("<").Id(codegen.This()).Dot(lenMethod).Call(),
					jen.Id("i").Op("++"),
				).Block(
					jen.Parens(
						jen.Op("*").Id(codegen.This()),
					).Index(jen.Id("i")).Dot(myIndexMemberName).Op("=").Id("i"),
				),
			},
			fmt.Sprintf("%sIRI prepends an IRI value to the front of a list of the property %q.", prependMethod, p.PropertyName())))
	methods = append(methods,
		codegen.NewCommentedPointerMethod(
			p.GetPrivatePackage().Path(),
			fmt.Sprintf("%sIRI", appendMethod),
			p.StructName(),
			[]jen.Code{jen.Id("v").Op("*").Qual("net/url", "URL")},
			/*ret=*/ nil,
			[]jen.Code{
				jen.Op("*").Id(codegen.This()).Op("=").Append(
					jen.Op("*").Id(codegen.This()),
					jen.Op("&").Id(p.iteratorTypeName().CamelName).Values(
						jen.Dict{
							jen.Id(iriMember):         jen.Id("v"),
							jen.Id(parentMemberName):  jen.Id(codegen.This()),
							jen.Id(myIndexMemberName): jen.Id(codegen.This()).Dot(lenMethod).Call(),
						},
					),
				),
			},
			fmt.Sprintf("%sIRI appends an IRI value to the back of a list of the property %q", appendMethod, p.PropertyName())))
	methods = append(methods,
		codegen.NewCommentedPointerMethod(
			p.GetPrivatePackage().Path(),
			fmt.Sprintf("%sIRI", setMethod),
			p.StructName(),
			[]jen.Code{jen.Id("idx").Int(), jen.Id("v").Op("*").Qual("net/url", "URL")},
			/*ret=*/ nil,
			[]jen.Code{
				jen.Parens(jen.Op("*").Id(codegen.This())).Index(jen.Id("idx")).Dot(parentMemberName).Op("=").Nil(),
				jen.Parens(jen.Op("*").Id(codegen.This())).Index(jen.Id("idx")).Op("=").Op("&").Id(p.iteratorTypeName().CamelName).Values(
					jen.Dict{
						jen.Id(iriMember):         jen.Id("v"),
						jen.Id(parentMemberName):  jen.Id(codegen.This()),
						jen.Id(myIndexMemberName): jen.Id("idx"),
					},
				),
			},
			fmt.Sprintf("%sIRI sets an IRI value to be at the specified index for the property %q. Panics if the index is out of bounds.", setMethod, p.PropertyName())))
	less = less.Else().If(
		jen.Id("idx1").Op("==").Lit(iriKindIndex),
	).Block(
		jen.Id("lhs").Op(":=").Id(codegen.This()).Index(jen.Id("i")).Dot(getIRIMethod).Call(),
		jen.Id("rhs").Op(":=").Id(codegen.This()).Index(jen.Id("j")).Dot(getIRIMethod).Call(),
		jen.Return(
			jen.Id("lhs").Dot("String").Call().Op("<").Id("rhs").Dot("String").Call(),
		),
	)
	// Remove Method
	methods = append(methods,
		codegen.NewCommentedPointerMethod(
			p.GetPrivatePackage().Path(),
			removeMethod,
			p.StructName(),
			[]jen.Code{jen.Id("idx").Int()},
			/*ret=*/ nil,
			[]jen.Code{
				jen.Parens(jen.Op("*").Id(codegen.This())).Index(jen.Id("idx")).Dot(parentMemberName).Op("=").Nil(),
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
				).Op("=").Op("&").Id(p.iteratorTypeName().CamelName).Values(),
				jen.Op("*").Id(codegen.This()).Op("=").Parens(
					jen.Op("*").Id(codegen.This()),
				).Index(
					jen.Empty(),
					jen.Len(jen.Op("*").Id(codegen.This())).Op("-").Lit(1),
				),
				jen.For(
					jen.Id("i").Op(":=").Id("idx"),
					jen.Id("i").Op("<").Id(codegen.This()).Dot(lenMethod).Call(),
					jen.Id("i").Op("++"),
				).Block(
					jen.Parens(
						jen.Op("*").Id(codegen.This()),
					).Index(jen.Id("i")).Dot(myIndexMemberName).Op("=").Id("i"),
				),
			},
			fmt.Sprintf("%s deletes an element at the specified index from a list of the property %q, regardless of its type. Panics if the index is out of bounds. Invalidates all iterators.", removeMethod, p.PropertyName())))
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
			fmt.Sprintf("%s returns the number of values that exist for the %q property.", lenMethod, p.PropertyName())))
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
			fmt.Sprintf("%s swaps the location of values at two indices for the %q property.", swapMethod, p.PropertyName())))
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
			fmt.Sprintf("%s computes whether another property is less than this one. Mixing types results in a consistent but arbitrary ordering", lessMethod)))
	// KindIndex Method
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
			fmt.Sprintf("%s computes an arbitrary value for indexing this kind of value. This is a leaky API method specifically needed only for alternate implementations for go-fed. Applications should not use this method. Panics if the index is out of bounds.", kindIndexMethod)))
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
		[]jen.Code{jen.Id("o").Qual(p.GetPublicPackage().Path(), p.InterfaceName())},
		[]jen.Code{jen.Bool()},
		[]jen.Code{
			lessCode,
			jen.For(
				jen.Id("i").Op(":=").Lit(0),
				jen.Id("i").Op("<").Id("l"),
				jen.Id("i").Op("++"),
			).Block(
				jen.If(
					jen.Id(codegen.This()).Index(jen.Id("i")).Dot(compareLessMethod).Call(jen.Id("o").Dot(atMethodName).Call(jen.Id("i"))),
				).Block(
					jen.Return(jen.True()),
				).Else().If(
					jen.Id("o").Dot(atMethodName).Call(jen.Id("i")).Dot(compareLessMethod).Call(jen.Id(codegen.This()).Index(jen.Id("i"))),
				).Block(
					jen.Return(jen.False()),
				),
			),
			jen.Return(jen.Id("l1").Op("<").Id("l2")),
		},
		fmt.Sprintf("%s compares two instances of this property with an arbitrary but stable comparison. Applications should not use this because it is only meant to help alternative implementations to go-fed to be able to normalize nonfunctional properties.", compareLessMethod),
	))
	// At Method
	methods = append(methods, codegen.NewCommentedValueMethod(
		p.GetPrivatePackage().Path(),
		atMethodName,
		p.StructName(),
		[]jen.Code{jen.Id("index").Int()},
		[]jen.Code{jen.Qual(p.GetPublicPackage().Path(), p.iteratorInterfaceName())},
		[]jen.Code{
			jen.Return(
				jen.Id(codegen.This()).Index(jen.Id("index")),
			),
		},
		fmt.Sprintf("%s returns the property value for the specified index. Panics if the index is out of bounds.", atMethodName)))
	// Empty Method
	methods = append(methods, codegen.NewCommentedValueMethod(
		p.GetPrivatePackage().Path(),
		emptyMethod,
		p.StructName(),
		/*params=*/ nil,
		[]jen.Code{jen.Bool()},
		[]jen.Code{
			jen.Return(
				jen.Id(codegen.This()).Dot(lenMethod).Call().Op("==").Lit(0),
			),
		},
		fmt.Sprintf("%s returns returns true if there are no elements.", emptyMethod)))
	// Begin Method
	methods = append(methods, codegen.NewCommentedValueMethod(
		p.GetPrivatePackage().Path(),
		beginMethod,
		p.StructName(),
		/*params=*/ nil,
		[]jen.Code{jen.Qual(p.GetPublicPackage().Path(), p.iteratorInterfaceName())},
		[]jen.Code{
			jen.If(
				jen.Id(codegen.This()).Dot(emptyMethod).Call(),
			).Block(
				jen.Return(jen.Nil()),
			).Else().Block(
				jen.Return(
					jen.Id(codegen.This()).Index(jen.Lit(0)),
				),
			),
		},
		fmt.Sprintf("%s returns the first iterator, or nil if empty. Can be used with the iterator's %s method and this property's %s method to iterate from front to back through all values.", beginMethod, nextMethod, endMethod)))
	// End Method
	methods = append(methods, codegen.NewCommentedValueMethod(
		p.GetPrivatePackage().Path(),
		endMethod,
		p.StructName(),
		/*params=*/ nil,
		[]jen.Code{jen.Qual(p.GetPublicPackage().Path(), p.iteratorInterfaceName())},
		[]jen.Code{
			jen.Return(jen.Nil()),
		},
		fmt.Sprintf("%s returns beyond-the-last iterator, which is nil. Can be used with the iterator's %s method and this property's %s method to iterate from front to back through all values.", endMethod, nextMethod, beginMethod)))
	methods = append(methods, p.commonMethods()...)
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
		fmt.Sprintf("%s converts this into an interface representation suitable for marshalling into a text or binary format. Applications should not need this function as most typical use cases serialize types instead of individual properties. It is exposed for alternatives to go-fed implementations to use.", p.serializeFnName()))
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
			jen.Id("ret").Op(":=").Id(p.StructName()).Call(jen.Id(codegen.This())),
			jen.Return(
				jen.Op("&").Id("ret"),
				jen.Err(),
			),
		).Else().If(
			jen.Id("p").Op("!=").Nil(),
		).Block(
			jen.Id(codegen.This()).Op("=").Append(
				jen.Id(codegen.This()),
				jen.Id("p"),
			),
		)
	}
	deserialize := codegen.NewCommentedFunction(
		p.GetPrivatePackage().Path(),
		p.DeserializeFnName(),
		[]jen.Code{jen.Id("m").Map(jen.String()).Interface()},
		[]jen.Code{jen.Qual(p.GetPublicPackage().Path(), p.InterfaceName()), jen.Error()},
		[]jen.Code{
			jen.Var().Id(codegen.This()).Index().Op("*").Id(p.iteratorTypeName().CamelName),
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
			jen.Id("ret").Op(":=").Id(p.StructName()).Call(jen.Id(codegen.This())),
			jen.Id("pRet").Op(":=").Op("&").Id("ret"),
			jen.Commentf("Set up the properties for iteration."),
			jen.For(
				jen.List(jen.Id("idx"), jen.Id("ele")).Op(":=").Range().Id("ret"),
			).Block(
				jen.Id("ele").Dot(parentMemberName).Op("=").Id("pRet"),
				jen.Id("ele").Dot(myIndexMemberName).Op("=").Id("idx"),
			),
			jen.Return(
				jen.Id("pRet"),
				jen.Nil(),
			),
		},
		fmt.Sprintf("%s creates a %q property from an interface representation that has been unmarshalled from a text or binary format.", p.DeserializeFnName(), p.PropertyName()))
	return serialize, deserialize
}

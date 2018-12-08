package types

import (
	"fmt"
	"github.com/cjslep/activity/tools/exp/codegen"
	"github.com/dave/jennifer/jen"
	"sync"
)

const (
	typeInterfaceName  = "Type"
	extendedByMethod   = "IsExtendedBy"
	extendsMethod      = "Extends"
	disjointWithMethod = "IsDisjointWith"
	nameMethod         = "Name"
)

// TypeInterface returns the Type Interface that is needed for ActivityStream
// types to compile for methods dealing with extending, in the inheritance
// sense.
func TypeInterface(pkg string) *codegen.Interface {
	comment := fmt.Sprintf("%s represents an ActivityStreams type.", typeInterfaceName)
	funcs := []codegen.FunctionSignature{
		{
			Name:    nameMethod,
			Params:  nil,
			Ret:     []jen.Code{jen.String()},
			Comment: fmt.Sprintf("%s returns the ActivityStreams type name.", nameMethod),
		},
	}
	return codegen.NewInterface(pkg, typeInterfaceName, funcs, comment)
}

// Property represents a property of an ActivityStreams type.
type Property interface {
	PropertyName() string
	StructName() string
}

// TypeGenerator represents an ActivityStream type definition to generate in Go.
type TypeGenerator struct {
	packageName  string
	typeName     string
	comment      string
	properties   map[string]Property
	extends      []*TypeGenerator
	disjoint     []*TypeGenerator
	extendedBy   []*TypeGenerator
	cacheOnce    sync.Once
	cachedStruct *codegen.Struct
}

// NewTypeGenerator creates a new generator for a specific ActivityStreams Core
// or extension type. It will return an error if there are multiple properties
// have the same Name.
//
// The extends and disjoint parameters are allowed to be nil. These lists must
// also have unique (non-duplicated) elements.
//
// All TypeGenerators must be created before the Definition method is called, to
// ensure that type extension, in the inheritence sense, is properly set up.
func NewTypeGenerator(packageName, typeName, comment string,
	properties []Property,
	extends, disjoint []*TypeGenerator) (*TypeGenerator, error) {
	t := &TypeGenerator{
		packageName: packageName,
		typeName:    typeName,
		comment:     comment,
		properties:  make(map[string]Property, len(properties)),
		extends:     extends,
		disjoint:    disjoint,
	}
	for _, property := range properties {
		if _, has := t.properties[property.PropertyName()]; has {
			return nil, fmt.Errorf("type already has property with name %q", property.PropertyName())
		}
		t.properties[property.PropertyName()] = property
	}
	// Complete doubly-linked extends/extendedBy lists.
	for _, ext := range extends {
		ext.extendedBy = append(ext.extendedBy, t)
	}
	return t, nil
}

// Comment returns the comment for this type.
func (t *TypeGenerator) Comment() string {
	return t.comment
}

// TypeName returns the ActivityStreams name for this type.
func (t *TypeGenerator) TypeName() string {
	return t.typeName
}

// Extends returns the generators of types that this ActivityStreams type
// extends from.
func (t *TypeGenerator) Extends() []*TypeGenerator {
	return t.extends
}

// ExtendedBy returns the generators of types that extend from this
// ActivityStreams type.
func (t *TypeGenerator) ExtendedBy() []*TypeGenerator {
	return t.extendedBy
}

// Disjoint returns the generators of types that this ActivityStreams type is
// disjoint to.
func (t *TypeGenerator) Disjoint() []*TypeGenerator {
	return t.disjoint
}

// extendsFnName determines the name of the Extends function, which
// determines if this ActivityStreams type extends another one.
func (t *TypeGenerator) extendsFnName() string {
	return fmt.Sprintf("%s%s", t.TypeName(), extendsMethod)
}

// extendedByFnName determines the name of the ExtendedBy function, which
// determines if another ActivityStreams type extends this one.
func (t *TypeGenerator) extendedByFnName() string {
	return fmt.Sprintf("%s%s", t.TypeName(), extendedByMethod)
}

// disjointWithFnName determines the name of the DisjointWith function, which
// determines if another ActivityStreams type is disjoint with this one.
func (t *TypeGenerator) disjointWithFnName() string {
	return fmt.Sprintf("%s%s", t.TypeName(), disjointWithMethod)
}

// Definition generates the golang code for this ActivityStreams type.
func (t *TypeGenerator) Definition() *codegen.Struct {
	t.cacheOnce.Do(func() {
		members := make([]jen.Code, 0, len(t.properties))
		for name, property := range t.properties {
			members = append(members, jen.Id(name).Id(property.StructName()))
		}
		t.cachedStruct = codegen.NewStruct(
			jen.Commentf(t.Comment()),
			t.TypeName(),
			[]*codegen.Method{
				t.nameDefinition(),
				t.extendsDefinition(),
			},
			[]*codegen.Function{
				t.extendedByDefinition(),
				t.disjointWithDefinition(),
			},
			members)
	})
	return t.cachedStruct
}

// nameDefinition generates the golang method for returning the ActivityStreams
// type name.
func (t *TypeGenerator) nameDefinition() *codegen.Method {
	return codegen.NewCommentedValueMethod(
		t.packageName,
		nameMethod,
		t.TypeName(),
		/*params=*/ nil,
		[]jen.Code{jen.String()},
		[]jen.Code{
			jen.Return(jen.Lit(t.TypeName())),
		},
		jen.Commentf("%s returns the name of this type.", nameMethod))
}

// getAllParentExtends recursivley determines all the parent types that this
// type extends from.
func (t *TypeGenerator) getAllParentExtends(s []string, tg *TypeGenerator) {
	for _, e := range tg.Extends() {
		s = append(s, e.TypeName())
		t.getAllParentExtends(s, e)
	}
}

// extendsDefinition generates the golang method for determining if this
// ActivityStreams type extends another type. It requires the Type interface.
func (t *TypeGenerator) extendsDefinition() *codegen.Method {
	var extendNames []string
	t.getAllParentExtends(extendNames, t)
	extensions := make([]jen.Code, len(extendNames))
	for i, e := range extendNames {
		extensions[i] = jen.Lit(e)
	}
	impl := []jen.Code{jen.Comment("Shortcut implementation: this does not extend anything."), jen.Return(jen.False())}
	if len(extensions) > 0 {
		impl = []jen.Code{jen.Id("extensions").Op(":=").Index().String().Values(extensions...),
			jen.For(jen.List(
				jen.Id("_"),
				jen.Id("ext"),
			).Op(":=").Range().Id("extensions")).Block(
				jen.If(
					jen.Id("ext").Op("==").Id("other").Dot(nameMethod).Call(),
				).Block(
					jen.Return(jen.True()),
				),
			),
			jen.Return(jen.False())}
	}
	return codegen.NewCommentedValueMethod(
		t.packageName,
		t.extendsFnName(),
		t.TypeName(),
		[]jen.Code{jen.Id("other").Id(typeInterfaceName)},
		[]jen.Code{jen.Bool()},
		impl,
		jen.Commentf("%s returns true if the %s type extends from the other type.", t.extendsFnName(), t.TypeName()))
}

// getAllChildrenExtendBy recursivley determines all the child types that this
// type is extended by.
func (t *TypeGenerator) getAllChildrenExtendedBy(s []string, tg *TypeGenerator) {
	for _, e := range tg.ExtendedBy() {
		s = append(s, e.TypeName())
		t.getAllChildrenExtendedBy(s, e)
	}
}

// extendedByDefinition generates the golang function for determining if
// another ActivityStreams type extends this type. It requires the Type
// interface.
func (t *TypeGenerator) extendedByDefinition() *codegen.Function {
	var extendNames []string
	t.getAllChildrenExtendedBy(extendNames, t)
	extensions := make([]jen.Code, len(extendNames))
	for i, e := range extendNames {
		extensions[i] = jen.Lit(e)
	}
	impl := []jen.Code{jen.Comment("Shortcut implementation: is not extended by anything."), jen.Return(jen.False())}
	if len(extensions) > 0 {
		impl = []jen.Code{jen.Id("extensions").Op(":=").Index().String().Values(extensions...),
			jen.For(jen.List(
				jen.Id("_"),
				jen.Id("ext"),
			).Op(":=").Range().Id("extensions")).Block(
				jen.If(
					jen.Id("ext").Op("==").Id("other").Dot(nameMethod).Call(),
				).Block(
					jen.Return(jen.True()),
				),
			),
			jen.Return(jen.False())}
	}
	return codegen.NewCommentedFunction(
		t.packageName,
		t.extendedByFnName(),
		[]jen.Code{jen.Id("other").Id(typeInterfaceName)},
		[]jen.Code{jen.Bool()},
		impl,
		jen.Commentf("%s returns true if the other provided type extends from the %s type.", t.extendedByFnName(), t.TypeName()))
}

// getAllChildrenDisjointWith recursivley determines all the child types that this
// type is disjoint with.
func (t *TypeGenerator) getAllDisjointWith(s []string) {
	for _, e := range t.Disjoint() {
		s = append(s, e.TypeName())
		// Get all the disjoint type's children.
		t.getAllChildrenExtendedBy(s, e)
	}
}

// disjointWithDefinition generates the golang function for determining if
// another ActivityStreams type is disjoint with this type. It requires the Type
// interface.
func (t *TypeGenerator) disjointWithDefinition() *codegen.Function {
	var disjointNames []string
	t.getAllDisjointWith(disjointNames)
	disjointWith := make([]jen.Code, len(disjointNames))
	for i, d := range disjointNames {
		disjointWith[i] = jen.Lit(d)
	}
	impl := []jen.Code{jen.Comment("Shortcut implementation: is not disjoint with anything."), jen.Return(jen.False())}
	if len(disjointWith) > 0 {
		impl = []jen.Code{jen.Id("disjointWith").Op(":=").Index().String().Values(disjointWith...),
			jen.For(jen.List(
				jen.Id("_"),
				jen.Id("disjoint"),
			).Op(":=").Range().Id("disjointWith")).Block(
				jen.If(
					jen.Id("disjoint").Op("==").Id("other").Dot(nameMethod).Call(),
				).Block(
					jen.Return(jen.True()),
				),
			),
			jen.Return(jen.False())}
	}
	return codegen.NewCommentedFunction(
		t.packageName,
		t.disjointWithFnName(),
		[]jen.Code{jen.Id("other").Id(typeInterfaceName)},
		[]jen.Code{jen.Bool()},
		impl,
		jen.Commentf("%s returns true if the other provided type is disjoint with the %s type.", t.disjointWithFnName(), t.TypeName()))
}

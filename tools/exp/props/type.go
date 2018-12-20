package props

import (
	"fmt"
	"github.com/cjslep/activity/tools/exp/codegen"
	"github.com/dave/jennifer/jen"
	"sort"
	"strings"
	"sync"
)

const (
	typeInterfaceName   = "Type"
	extendedByMethod    = "IsExtendedBy"
	extendsMethod       = "Extends"
	disjointWithMethod  = "IsDisjointWith"
	typeNameMethod      = "Name"
	serializeMethodName = "Serialize"
	deserializeFnName   = "Deserialize"
	lessFnName          = "Less"
)

// TypeInterface returns the Type Interface that is needed for ActivityStream
// types to compile for methods dealing with extending, in the inheritance
// sense.
func TypeInterface(pkg string) *codegen.Interface {
	comment := fmt.Sprintf("%s represents an ActivityStreams type.", typeInterfaceName)
	funcs := []codegen.FunctionSignature{
		{
			Name:    typeNameMethod,
			Params:  nil,
			Ret:     []jen.Code{jen.String()},
			Comment: fmt.Sprintf("%s returns the ActivityStreams type name.", typeNameMethod),
		},
	}
	return codegen.NewInterface(pkg, typeInterfaceName, funcs, comment)
}

// Property represents a property of an ActivityStreams type.
type Property interface {
	PackageName() string
	PropertyName() string
	StructName() string
	SetKindFns(name string, ser, deser, less *codegen.Function) error
	DeserializeFnName() string
}

// TypeGenerator represents an ActivityStream type definition to generate in Go.
type TypeGenerator struct {
	packageName       string
	typeName          string
	comment           string
	properties        map[string]Property
	withoutProperties map[string]Property
	extends           []*TypeGenerator
	disjoint          []*TypeGenerator
	extendedBy        []*TypeGenerator
	cacheOnce         sync.Once
	cachedStruct      *codegen.Struct
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
// Additionally, all properties whose range is this type should have their
// SetKindFns method called with this TypeGenerator's KindSerializationFuncs for
// all code generation to correctly reference each other.
func NewTypeGenerator(packageName, typeName, comment string,
	properties, withoutProperties []Property,
	extends, disjoint []*TypeGenerator) (*TypeGenerator, error) {
	t := &TypeGenerator{
		packageName:       packageName,
		typeName:          typeName,
		comment:           comment,
		properties:        make(map[string]Property, len(properties)),
		withoutProperties: make(map[string]Property, len(withoutProperties)),
		extends:           extends,
		disjoint:          disjoint,
	}
	for _, property := range properties {
		if _, has := t.properties[property.PropertyName()]; has {
			return nil, fmt.Errorf("type already has property with name %q", property.PropertyName())
		}
		t.properties[property.PropertyName()] = property
	}
	for _, wop := range withoutProperties {
		if _, has := t.withoutProperties[wop.PropertyName()]; has {
			return nil, fmt.Errorf("type already has withoutproperty with name %q", wop.PropertyName())
		}
		t.withoutProperties[wop.PropertyName()] = wop
	}
	// Complete doubly-linked extends/extendedBy lists.
	for _, ext := range extends {
		ext.extendedBy = append(ext.extendedBy, t)
	}
	return t, nil
}

// AddDisjoint adds another TypeGenerator that is disjoint to this one.
func (t *TypeGenerator) AddDisjoint(o *TypeGenerator) {
	t.disjoint = append(t.disjoint, o)
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

// Properties returns the Properties of this type, mapped by their property
// name.
func (t *TypeGenerator) Properties() map[string]Property {
	return t.properties
}

// WithoutProperties returns the properties that do not apply to this type,
// mapped by their property name.
func (t *TypeGenerator) WithoutProperties() map[string]Property {
	return t.withoutProperties
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
		members := t.members()
		m := t.serializationMethod()
		ser, deser, less := t.KindSerializationFuncs()
		t.cachedStruct = codegen.NewStruct(
			jen.Commentf(t.Comment()),
			t.TypeName(),
			[]*codegen.Method{
				t.nameDefinition(),
				m,
			},
			[]*codegen.Function{
				t.extendsDefinition(),
				t.extendedByDefinition(),
				t.disjointWithDefinition(),
				ser,
				deser,
				less,
			},
			members)
	})
	return t.cachedStruct
}

func (t *TypeGenerator) allProperties() map[string]Property {
	p := t.properties
	// Properties of parents that are extended, minus DoesNotApplyTo
	var extends []*TypeGenerator
	extends = t.getAllParentExtends(extends, t)
	for _, ext := range t.extends {
		for k, v := range ext.Properties() {
			p[k] = v
		}
	}
	for _, ext := range t.extends {
		for k, _ := range ext.WithoutProperties() {
			delete(p, k)
		}
	}
	return p
}

// sortedProperty is a slice of Properties that implements the Sort interface.
type sortedProperty []Property

func (s sortedProperty) Less(i, j int) bool {
	return s[i].PropertyName() < s[j].PropertyName()
}

func (s sortedProperty) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortedProperty) Len() int {
	return len(s)
}

func (t *TypeGenerator) members() (members []jen.Code) {
	p := t.allProperties()
	// Sort the properties for readability
	sortedMembers := make(sortedProperty, 0, len(p))
	for _, property := range p {
		sortedMembers = append(sortedMembers, property)
	}
	sort.Sort(sortedMembers)
	// Convert to jen.Code
	members = make([]jen.Code, 0, len(p))
	for _, property := range sortedMembers {
		members = append(members, jen.Id(strings.Title(property.PropertyName())).Qual(property.PackageName(), property.StructName()))
	}
	return
}

// nameDefinition generates the golang method for returning the ActivityStreams
// type name.
func (t *TypeGenerator) nameDefinition() *codegen.Method {
	return codegen.NewCommentedValueMethod(
		t.packageName,
		typeNameMethod,
		t.TypeName(),
		/*params=*/ nil,
		[]jen.Code{jen.String()},
		[]jen.Code{
			jen.Return(jen.Lit(t.TypeName())),
		},
		jen.Commentf("%s returns the name of this type.", typeNameMethod))
}

// getAllParentExtends recursively determines all the parent types that this
// type extends from.
func (t *TypeGenerator) getAllParentExtends(s []*TypeGenerator, tg *TypeGenerator) []*TypeGenerator {
	for _, e := range tg.Extends() {
		s = append(s, e)
		s = append(s, t.getAllParentExtends(s, e)...)
	}
	return s
}

// extendsDefinition generates the golang function for determining if this
// ActivityStreams type extends another type. It requires the Type interface.
func (t *TypeGenerator) extendsDefinition() *codegen.Function {
	var extends []*TypeGenerator
	extends = t.getAllParentExtends(extends, t)
	extendNames := make(map[string]struct{}, len(extends))
	for _, ext := range extends {
		extendNames[ext.TypeName()] = struct{}{}
	}
	extensions := make([]jen.Code, len(extendNames))
	for e := range extendNames {
		extensions = append(extensions, jen.Lit(e))
	}
	impl := []jen.Code{jen.Comment("Shortcut implementation: this does not extend anything."), jen.Return(jen.False())}
	if len(extensions) > 0 {
		impl = []jen.Code{jen.Id("extensions").Op(":=").Index().String().Values(extensions...),
			jen.For(jen.List(
				jen.Id("_"),
				jen.Id("ext"),
			).Op(":=").Range().Id("extensions")).Block(
				jen.If(
					jen.Id("ext").Op("==").Id("other").Dot(typeNameMethod).Call(),
				).Block(
					jen.Return(jen.True()),
				),
			),
			jen.Return(jen.False())}
	}
	return codegen.NewCommentedFunction(
		t.packageName,
		t.extendsFnName(),
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
					jen.Id("ext").Op("==").Id("other").Dot(typeNameMethod).Call(),
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
					jen.Id("disjoint").Op("==").Id("other").Dot(typeNameMethod).Call(),
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

// serializationMethod returns the method needed to serialize a TypeGenerator as
// a property.
func (t *TypeGenerator) serializationMethod() (ser *codegen.Method) {
	ser = codegen.NewCommentedValueMethod(
		t.packageName,
		serializeMethodName,
		t.TypeName(),
		/*params=*/ nil,
		[]jen.Code{jen.Interface(), jen.Error()},
		[]jen.Code{
			// TODO
			jen.Commentf("TODO: Serialization code for %s", t.TypeName()),
		},
		jen.Commentf("%s converts this into an interface representation suitable for marshalling into a text or binary format.", serializeMethodName))
	return
}

// KindSerializationFuncs returns free function references that can be used to
// treat a TypeGenerator as another property's Kind.
func (t *TypeGenerator) KindSerializationFuncs() (ser, deser, less *codegen.Function) {
	serName := fmt.Sprintf("%s%s", serializeMethodName, t.TypeName())
	ser = codegen.NewCommentedFunction(
		t.packageName,
		serName,
		[]jen.Code{jen.Id("s").Id(t.TypeName())},
		[]jen.Code{jen.Interface(), jen.Error()},
		[]jen.Code{
			jen.Return(
				jen.Id("s").Dot(serializeMethodName).Call(),
			),
		},
		jen.Commentf("%s calls %s on the %s type.", serName, serializeMethodName, t.TypeName()))
	deserName := fmt.Sprintf("%s%s", deserializeFnName, t.TypeName())
	deserCode := jen.Empty()
	for name, prop := range t.allProperties() {
		deserCode = deserCode.Add(
			jen.If(
				jen.List(
					jen.Id("p"),
					jen.Err(),
				).Op(":=").Qual(prop.PackageName(), prop.DeserializeFnName()).Call(jen.Id("m")),
				jen.Err().Op("!=").Nil(),
			).Block(
				jen.Return(jen.Nil(), jen.Err()),
			).Else().Block(
				jen.Id(codegen.This()).Dot(strings.Title(name)).Op("=").Op("*").Id("p"),
			).Line())
	}
	deser = codegen.NewCommentedFunction(
		t.packageName,
		deserName,
		[]jen.Code{jen.Id("m").Map(jen.String()).Interface()},
		[]jen.Code{jen.Op("*").Id(t.TypeName()), jen.Error()},
		[]jen.Code{
			jen.Id(codegen.This()).Op(":=").Op("&").Id(t.TypeName()).Values(),
			deserCode,
			jen.Return(jen.Id(codegen.This()), jen.Nil()),
		},
		jen.Commentf("%s creates a %s from a map representation that has been unmarshalled from a text or binary format.", deserName, t.TypeName()))
	lessName := fmt.Sprintf("%s%s", lessFnName, t.TypeName())
	less = codegen.NewCommentedFunction(
		t.packageName,
		lessName,
		[]jen.Code{
			jen.Id("i"),
			jen.Id("j").Op("*").Id(t.TypeName()),
		},
		[]jen.Code{jen.Bool()},
		[]jen.Code{
			// TODO
			jen.Commentf("TODO: Less code for %s", t.TypeName()),
		},
		jen.Commentf("%s computes which %s is lesser, with an arbitrary but stable determination", lessName, t.TypeName()))
	return
}

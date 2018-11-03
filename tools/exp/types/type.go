package types

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/go-fed/activity/tools/exp/codegen"
)

const (
	typeInterfaceName  = "Type"
	extendsOtherMethod = "IsExtendedBy"
	nameMethod         = "Name"
)

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
	packageName string
	typeName    string
	comment     string
	properties  map[string]Property
	extends     []*TypeGenerator
	disjoint    []*TypeGenerator
}

// NewTypeGenerator creates a new generator for a specific ActivityStreams Core
// or extension type. It will return an error if there are multiple properties
// have the same Name.
func NewTypeGenerator(packageName, typeName, comment string,
	properties []Property,
	extends, disjoint []*TypeGenerator) (*TypeGenerator, error) {
	t := &TypeGenerator{
		packageName: packageName,
		typeName:    typeName,
		comment:     comment,
		properties:  make(map[string]Property, len(properties)),
		extends: extends,
		disjoint: disjoint,
	}
	for _, property := range properties {
		if _, has := t.properties[property.PropertyName()]; has {
			return nil, fmt.Errorf("type already has property with name %q", property.PropertyName())
		}
		t.properties[property.PropertyName()] = property
	}
	return t, nil
}

func (t *TypeGenerator) Comment() string {
	return t.comment
}

func (t *TypeGenerator) TypeName() string {
	return t.typeName
}

func (t *TypeGenerator) extendsOtherFnName() string {
	return fmt.Sprintf("%s%s", t.TypeName(), extendsOtherMethod)
}

func (t *TypeGenerator) Definition() *codegen.Struct {
	// TODO: Cache this generation.
	members := make([]jen.Code, 0, len(t.properties))
	for name, property := range t.properties {
		members = append(members, jen.Id(name).Id(property.StructName()))
	}
	return codegen.NewStruct(
		jen.Commentf(t.Comment()),
		t.TypeName(),
		[]*codegen.Method{
			t.nameDefinition(),
		},
		[]*codegen.Function{
			t.extendsOtherDefinition(),
		},
		members)
}

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

func (t *TypeGenerator) extendsOtherDefinition() *codegen.Function {
	// TODO: Chain up higher than 1 level
	extensions := make([]jen.Code, len(t.extends))
	for i, e := range t.extends {
		extensions[i] = jen.Lit(e.TypeName())
	}
	impl := []jen.Code{jen.Comment("Shortcut implementation: is not extended by anything."), jen.Return(jen.False())}
	if len(extensions) > 0 {
		impl = []jen.Code{jen.Id("extensions").Op(":=").Index().String().Values(extensions...),
			jen.For(jen.List(
				jen.Id("_"),
				jen.Id("ext"),
			).Op(":=").Range().Id("extensions")).Block(
				jen.If(
					jen.Id("ext").Dot(nameMethod).Call().Op("==").Id("other").Dot(nameMethod).Call(),
				).Block(
					jen.Return(jen.True()),
				),
			),
			jen.Return(jen.False())}
	}
	return codegen.NewCommentedFunction(
		t.packageName,
		t.extendsOtherFnName(),
		[]jen.Code{jen.Id("other").Id(typeInterfaceName)},
		[]jen.Code{jen.Bool()},
		impl,
		jen.Commentf("%s returns true if the other provided type extends the %s type", t.extendsOtherFnName(), t.TypeName()))
}

// TODO: Disjoint function, similar to extends function.

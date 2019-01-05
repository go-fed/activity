package rdf

import (
	"fmt"
	"github.com/cjslep/activity/tools/exp/codegen"
	"github.com/dave/jennifer/jen"
	"net/url"
	"strings"
)

const (
	rdfSpec        = "http://www.w3.org/1999/02/22-rdf-syntax-ns#"
	langstringSpec = "langString"
	propertySpec   = "Property"
)

func SerializeValueFunction(pkg, valueName string,
	concreteType jen.Code,
	impl []jen.Code) *codegen.Function {
	name := fmt.Sprintf("Serialize%s", strings.Title(valueName))
	return codegen.NewCommentedFunction(
		pkg,
		name,
		[]jen.Code{jen.Id(codegen.This()).Add(concreteType)},
		[]jen.Code{jen.Interface(), jen.Error()},
		impl,
		jen.Commentf("%s converts a %s value to an interface representation suitable for marshalling into a text or binary format.", name, valueName))
}

func DeserializeValueFunction(pkg, valueName string,
	concreteType jen.Code,
	impl []jen.Code) *codegen.Function {
	name := fmt.Sprintf("Deserialize%s", strings.Title(valueName))
	return codegen.NewCommentedFunction(
		pkg,
		name,
		[]jen.Code{jen.Id(codegen.This()).Interface()},
		[]jen.Code{concreteType, jen.Error()},
		impl,
		jen.Commentf("%s creates %s value from an interface representation that has been unmarshalled from a text or binary format.", name, valueName))
}

func LessFunction(pkg, valueName string,
	concreteType jen.Code,
	impl []jen.Code) *codegen.Function {
	name := fmt.Sprintf("Less%s", strings.Title(valueName))
	return codegen.NewCommentedFunction(
		pkg,
		name,
		[]jen.Code{jen.List(jen.Id("lhs"), jen.Id("rhs")).Add(concreteType)},
		[]jen.Code{jen.Bool()},
		impl,
		jen.Commentf("%s returns true if the left %s value is less than the right value.", name, valueName))
}

type RDFOntology struct {
	Package string
	alias   string
}

func (o *RDFOntology) SpecURI() string {
	return rdfSpec
}

func (o *RDFOntology) Load() ([]RDFNode, error) {
	return o.LoadAsAlias("")
}

func (o *RDFOntology) LoadAsAlias(s string) ([]RDFNode, error) {
	o.alias = s
	return []RDFNode{
		&AliasedDelegate{
			Spec:     rdfSpec,
			Alias:    s,
			Name:     langstringSpec,
			Delegate: &langstring{pkg: o.Package, alias: o.alias},
		},
		&AliasedDelegate{
			Spec:     rdfSpec,
			Alias:    s,
			Name:     propertySpec,
			Delegate: &property{},
		},
	}, nil
}

func (o *RDFOntology) LoadSpecificAsAlias(alias, name string) ([]RDFNode, error) {
	switch name {
	case langstringSpec:
		return []RDFNode{
			&AliasedDelegate{
				Spec:     "",
				Alias:    "",
				Name:     alias,
				Delegate: &langstring{pkg: o.Package, alias: o.alias},
			},
		}, nil
	case propertySpec:
		return []RDFNode{
			&AliasedDelegate{
				Spec:     "",
				Alias:    "",
				Name:     alias,
				Delegate: &property{},
			},
		}, nil
	}
	return nil, fmt.Errorf("rdf ontology cannot find %q to make alias %q", name, alias)
}

func (o *RDFOntology) LoadElement(name string, payload map[string]interface{}) ([]RDFNode, error) {
	return nil, nil
}

func (o *RDFOntology) GetByName(name string) (RDFNode, error) {
	name = strings.TrimPrefix(name, o.SpecURI())
	switch name {
	case langstringSpec:
		return &langstring{pkg: o.Package, alias: o.alias}, nil
	case propertySpec:
		return &property{}, nil
	}
	return nil, fmt.Errorf("rdf ontology could not find node for name %s", name)
}

var _ RDFNode = &langstring{}

type langstring struct {
	alias string
	pkg   string
}

func (l *langstring) Enter(key string, ctx *ParsingContext) (bool, error) {
	return true, fmt.Errorf("rdf langstring cannot be entered")
}

func (l *langstring) Exit(key string, ctx *ParsingContext) (bool, error) {
	return true, fmt.Errorf("rdf langstring cannot be exited")
}

func (l *langstring) Apply(key string, value interface{}, ctx *ParsingContext) (bool, error) {
	for k, p := range ctx.Result.Vocab.Properties {
		for _, ref := range p.Range {
			if ref.Name == langstringSpec && ref.Vocab == l.alias {
				p.NaturalLanguageMap = true
				ctx.Result.Vocab.Properties[k] = p
				break
			}
		}
	}
	u, e := url.Parse(rdfSpec + langstringSpec)
	if e != nil {
		return true, e
	}
	e = ctx.Result.GetReference(rdfSpec).SetValue(langstringSpec, &VocabularyValue{
		Name:           langstringSpec,
		URI:            u,
		DefinitionType: "map[string]string",
		Zero:           "nil",
		SerializeFn: SerializeValueFunction(
			l.pkg,
			langstringSpec,
			jen.Map(jen.String()).String(),
			[]jen.Code{
				jen.Return(
					jen.Id(codegen.This()),
					jen.Nil(),
				),
			}),
		DeserializeFn: DeserializeValueFunction(
			l.pkg,
			langstringSpec,
			jen.Map(jen.String()).String(),
			[]jen.Code{
				jen.If(
					jen.List(
						jen.Id("m"),
						jen.Id("ok"),
					).Op(":=").Id(codegen.This()).Assert(jen.Map(jen.String()).Interface()),
					jen.Id("ok"),
				).Block(
					jen.Id("r").Op(":=").Make(jen.Map(jen.String()).String()),
					jen.For(
						jen.List(
							jen.Id("k"),
							jen.Id("v"),
						).Op(":=").Range().Id("m"),
					).Block(
						jen.If(
							jen.List(
								jen.Id("s"),
								jen.Id("ok"),
							).Op(":=").Id("v").Assert(jen.String()),
							jen.Id("ok"),
						).Block(
							jen.Id("r").Index(jen.Id("k")).Op("=").Id("s"),
						).Else().Block(
							jen.Return(
								jen.Nil(),
								jen.Qual("fmt", "Errorf").Call(
									jen.Lit("value %v cannot be interpreted as a string for rdf:langString"),
									jen.Id("v"),
								),
							),
						),
					),
					jen.Return(
						jen.Id("r"),
						jen.Nil(),
					),
				).Else().Block(
					jen.Return(
						jen.Nil(),
						jen.Qual("fmt", "Errorf").Call(
							jen.Lit("%v cannot be interpreted as a map[string]interface{} for rdf:langString"),
							jen.Id(codegen.This()),
						),
					),
				),
			}),
		LessFn: LessFunction(
			l.pkg,
			langstringSpec,
			jen.Map(jen.String()).String(),
			[]jen.Code{
				jen.Var().Id("lk").Index().String(),
				jen.Var().Id("rk").Index().String(),
				jen.For(
					jen.List(
						jen.Id("k"),
					).Op(":=").Range().Id("lhs"),
				).Block(
					jen.Id("lk").Op("=").Append(
						jen.Id("lk"),
						jen.Id("k"),
					),
				),
				jen.For(
					jen.List(
						jen.Id("k"),
					).Op(":=").Range().Id("rhs"),
				).Block(
					jen.Id("rk").Op("=").Append(
						jen.Id("rk"),
						jen.Id("k"),
					),
				),
				jen.Qual("sort", "Sort").Call(
					jen.Qual("sort", "StringSlice").Call(jen.Id("lk")),
				),
				jen.Qual("sort", "Sort").Call(
					jen.Qual("sort", "StringSlice").Call(jen.Id("rk")),
				),
				jen.For(
					jen.Id("i").Op(":=").Lit(0),
					jen.Id("i").Op("<").Len(jen.Id("lk")).Op("&&").Id("i").Op("<").Len(jen.Id("rk")),
					jen.Id("i").Op("++"),
				).Block(
					jen.If(
						jen.Id("lk").Index(jen.Id("i")).Op("<").Id("rk").Index(jen.Id("i")),
					).Block(
						jen.Return(jen.True()),
					).Else().If(
						jen.Id("rk").Index(jen.Id("i")).Op("<").Id("lk").Index(jen.Id("i")),
					).Block(
						jen.Return(jen.False()),
					).Else().If(
						jen.Id("lhs").Index(jen.Id("lk").Index(jen.Id("i"))).Op("<").Id("rhs").Index(jen.Id("rk").Index(jen.Id("i"))),
					).Block(
						jen.Return(jen.True()),
					).Else().If(
						jen.Id("rhs").Index(jen.Id("rk").Index(jen.Id("i"))).Op("<").Id("lhs").Index(jen.Id("lk").Index(jen.Id("i"))),
					).Block(
						jen.Return(jen.False()),
					),
				),
				jen.If(
					jen.Len(jen.Id("lk")).Op("<").Len(jen.Id("rk")),
				).Block(
					jen.Return(jen.True()),
				).Else().Block(
					jen.Return(jen.False()),
				),
			}),
	})
	return true, e
}

var _ RDFNode = &property{}

type property struct{}

func (p *property) Enter(key string, ctx *ParsingContext) (bool, error) {
	return true, fmt.Errorf("rdf property cannot be entered")
}

func (p *property) Exit(key string, ctx *ParsingContext) (bool, error) {
	return true, fmt.Errorf("rdf property cannot be exited")
}

func (p *property) Apply(key string, value interface{}, ctx *ParsingContext) (bool, error) {
	// Prepare a new VocabularyProperty in the context. If one already
	// exists, skip.
	if _, ok := ctx.Current.(*VocabularyProperty); ok {
		return true, nil
	} else if !ctx.IsReset() {
		return true, fmt.Errorf("rdf property applied with non-reset ParsingContext")
	}
	ctx.Current = &VocabularyProperty{}
	return true, nil
}

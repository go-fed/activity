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
				// TODO
			}),
		DeserializeFn: DeserializeValueFunction(
			l.pkg,
			langstringSpec,
			jen.Map(jen.String()).String(),
			[]jen.Code{
				// TODO
			}),
		LessFn: LessFunction(
			l.pkg,
			langstringSpec,
			jen.Map(jen.String()).String(),
			[]jen.Code{
				// TODO
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

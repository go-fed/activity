package rdf

import (
	"fmt"
	"strings"
)

const (
	rdfSpec        = "http://www.w3.org/1999/02/22-rdf-syntax-ns#"
	langstringSpec = "langString"
	propertySpec   = "Property"
	valueSpec      = "value"
)

type RDFOntology struct {
	alias string
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
			Delegate: &langstring{alias: o.alias},
		},
		&AliasedDelegate{
			Spec:     rdfSpec,
			Alias:    s,
			Name:     propertySpec,
			Delegate: &property{},
		},
		&AliasedDelegate{
			Spec:     rdfSpec,
			Alias:    s,
			Name:     valueSpec,
			Delegate: &value{},
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
				Delegate: &langstring{alias: o.alias},
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
	case valueSpec:
		return []RDFNode{
			&AliasedDelegate{
				Spec:     "",
				Alias:    "",
				Name:     alias,
				Delegate: &value{},
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
		return &langstring{alias: o.alias}, nil
	case propertySpec:
		return &property{}, nil
	case valueSpec:
		return &value{}, nil
	}
	return nil, fmt.Errorf("rdf ontology could not find node for name %s", name)
}

var _ RDFNode = &langstring{}

type langstring struct {
	alias string
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
	return true, nil
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

var _ RDFNode = &value{}

type value struct{}

func (v *value) Enter(key string, ctx *ParsingContext) (bool, error) {
	return true, fmt.Errorf("rdf value cannot be entered")
}

func (v *value) Exit(key string, ctx *ParsingContext) (bool, error) {
	return true, fmt.Errorf("rdf value cannot be exited")
}

func (v *value) Apply(key string, value interface{}, ctx *ParsingContext) (bool, error) {
	// TODO: Act as value
	return true, fmt.Errorf("rdf value cannot be applied")
}

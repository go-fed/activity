package rdf

import (
	"fmt"
)

const (
	rdfSpec        = "http://www.w3.org/1999/02/22-rdf-syntax-ns#"
	langstringSpec = "langstring"
	propertySpec   = "Property"
	valueSpec      = "value"
)

type RDFOntology struct{}

func (o *RDFOntology) SpecURI() string {
	return rdfSpec
}

func (o *RDFOntology) Load() ([]RDFNode, error) {
	return o.LoadAsAlias("")
}

func (o *RDFOntology) LoadAsAlias(s string) ([]RDFNode, error) {
	return []RDFNode{
		&AliasedDelegate{
			Spec:     rdfSpec,
			Alias:    s,
			Name:     langstringSpec,
			Delegate: &langstring{},
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

func (o *RDFOntology) LoadElement(name string, payload map[string]interface{}) ([]RDFNode, error) {
	return nil, nil
}

var _ RDFNode = &langstring{}

type langstring struct{}

func (l *langstring) Enter(key string, ctx *ParsingContext) (bool, error) {
	return true, fmt.Errorf("rdf langstring cannot be entered")
}

func (l *langstring) Exit(key string, ctx *ParsingContext) (bool, error) {
	return true, fmt.Errorf("rdf langstring cannot be exited")
}

func (l *langstring) Apply(key string, value interface{}, ctx *ParsingContext) (bool, error) {
	return true, fmt.Errorf("rdf langstring cannot be applied")
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
	return true, fmt.Errorf("rdf property cannot be applied")
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
	return true, fmt.Errorf("rdf value cannot be applied")
}

package schema

import (
	"fmt"
	"github.com/cjslep/activity/tools/exp/rdf"
)

const (
	schemaSpec     = "http://schema.org/"
	exampleSpec    = "workExample"
	mainEntitySpec = "mainEntity"
	urlSpec        = "URL"
)

type SchemaOntology struct{}

func (o *SchemaOntology) SpecURI() string {
	return schemaSpec
}

func (o *SchemaOntology) Load() ([]rdf.RDFNode, error) {
	return o.LoadAsAlias("")
}

func (o *SchemaOntology) LoadAsAlias(s string) ([]rdf.RDFNode, error) {
	return []rdf.RDFNode{
		&rdf.AliasedDelegate{
			Spec:     schemaSpec,
			Alias:    s,
			Name:     exampleSpec,
			Delegate: &example{},
		},
		&rdf.AliasedDelegate{
			Spec:     schemaSpec,
			Alias:    s,
			Name:     mainEntitySpec,
			Delegate: &mainEntity{},
		},
		&rdf.AliasedDelegate{
			Spec:     schemaSpec,
			Alias:    s,
			Name:     urlSpec,
			Delegate: &URL{},
		},
	}, nil
}

func (o *SchemaOntology) LoadSpecificAsAlias(alias, name string) ([]rdf.RDFNode, error) {
	switch name {
	case exampleSpec:
		return []rdf.RDFNode{
			&rdf.AliasedDelegate{
				Spec:     "",
				Alias:    "",
				Name:     alias,
				Delegate: &example{},
			},
		}, nil
	case mainEntitySpec:
		return []rdf.RDFNode{
			&rdf.AliasedDelegate{
				Spec:     "",
				Alias:    "",
				Name:     alias,
				Delegate: &mainEntity{},
			},
		}, nil
	case urlSpec:
		return []rdf.RDFNode{
			&rdf.AliasedDelegate{
				Spec:     "",
				Alias:    "",
				Name:     alias,
				Delegate: &URL{},
			},
		}, nil
	}
	return nil, fmt.Errorf("schema ontology cannot find %q to alias to %q", name, alias)
}

func (o *SchemaOntology) LoadElement(name string, payload map[string]interface{}) ([]rdf.RDFNode, error) {
	return nil, nil
}

var _ rdf.RDFNode = &example{}

type example struct{}

func (e *example) Enter(key string, ctx *rdf.ParsingContext) (bool, error) {
	// TODO
	return true, nil
}

func (e *example) Exit(key string, ctx *rdf.ParsingContext) (bool, error) {
	// TODO
	return true, nil
}

func (e *example) Apply(key string, value interface{}, ctx *rdf.ParsingContext) (bool, error) {
	// TODO
	return true, nil
}

var _ rdf.RDFNode = &mainEntity{}

type mainEntity struct{}

func (m *mainEntity) Enter(key string, ctx *rdf.ParsingContext) (bool, error) {
	// TODO
	return true, nil
}

func (m *mainEntity) Exit(key string, ctx *rdf.ParsingContext) (bool, error) {
	// TODO
	return true, nil
}

func (m *mainEntity) Apply(key string, value interface{}, ctx *rdf.ParsingContext) (bool, error) {
	// TODO
	return true, nil
}

var _ rdf.RDFNode = &URL{}

type URL struct{}

func (u *URL) Enter(key string, ctx *rdf.ParsingContext) (bool, error) {
	// TODO
	return true, nil
}

func (u *URL) Exit(key string, ctx *rdf.ParsingContext) (bool, error) {
	// TODO
	return true, nil
}

func (u *URL) Apply(key string, value interface{}, ctx *rdf.ParsingContext) (bool, error) {
	// TODO
	return true, nil
}

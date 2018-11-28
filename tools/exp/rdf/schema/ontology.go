package schema

import (
	"github.com/go-fed/activity/tools/exp/rdf"
)

type SchemaOntology struct{}

func (o *SchemaOntology) SpecURI() string {
	return "http://schema.org/"
}

func (o *SchemaOntology) Load() ([]rdf.RDFNode, error) {
	return nil, nil
}

func (o *SchemaOntology) LoadAsAlias(s string) ([]rdf.RDFNode, error) {
	return nil, nil
}

func (o *SchemaOntology) LoadElement(name string, payload map[string]interface{}) ([]rdf.RDFNode, error) {
	return nil, nil
}

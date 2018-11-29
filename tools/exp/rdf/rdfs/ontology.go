package rdfs

import (
	"github.com/cjslep/activity/tools/exp/rdf"
)

type RDFSchemaOntology struct{}

func (o *RDFSchemaOntology) SpecURI() string {
	return "http://www.w3.org/2000/01/rdf-schema#"
}

func (o *RDFSchemaOntology) Load() ([]rdf.RDFNode, error) {
	return nil, nil
}

func (o *RDFSchemaOntology) LoadAsAlias(s string) ([]rdf.RDFNode, error) {
	return nil, nil
}

func (o *RDFSchemaOntology) LoadElement(name string, payload map[string]interface{}) ([]rdf.RDFNode, error) {
	return nil, nil
}

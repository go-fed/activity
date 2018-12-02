package rdfs

import (
	"fmt"
	"github.com/cjslep/activity/tools/exp/rdf"
	"strings"
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

func (o *RDFSchemaOntology) LoadSpecificAsAlias(alias, name string) ([]rdf.RDFNode, error) {
	return nil, nil
}

func (o *RDFSchemaOntology) LoadElement(name string, payload map[string]interface{}) ([]rdf.RDFNode, error) {
	return nil, nil
}

func (o *RDFSchemaOntology) GetByName(name string) (rdf.RDFNode, error) {
	name = strings.TrimPrefix(name, o.SpecURI())
	switch name {
	// TODO
	}
	return nil, fmt.Errorf("rdfs ontology could not find node for name %s", name)
}

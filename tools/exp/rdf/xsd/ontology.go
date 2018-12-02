package xsd

import (
	"github.com/cjslep/activity/tools/exp/rdf"
)

type XMLOntology struct{}

func (o *XMLOntology) SpecURI() string {
	return "http://www.w3.org/2001/XMLSchema#"
}

func (o *XMLOntology) Load() ([]rdf.RDFNode, error) {
	return nil, nil
}

func (o *XMLOntology) LoadAsAlias(s string) ([]rdf.RDFNode, error) {
	return nil, nil
}

func (o *XMLOntology) LoadSpecificAsAlias(alias, name string) ([]rdf.RDFNode, error) {
	return nil, nil
}

func (o *XMLOntology) LoadElement(name string, payload map[string]interface{}) ([]rdf.RDFNode, error) {
	return nil, nil
}

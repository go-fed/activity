package xsd

import (
	"fmt"
	"github.com/cjslep/activity/tools/exp/rdf"
	"strings"
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

func (o *XMLOntology) GetByName(name string) (rdf.RDFNode, error) {
	name = strings.TrimPrefix(name, o.SpecURI())
	switch name {
	// TODO
	}
	return nil, fmt.Errorf("xsd ontology could not find node for name %s", name)
}

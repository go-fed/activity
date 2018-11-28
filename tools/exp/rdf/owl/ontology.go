package owl

import (
	"github.com/go-fed/activity/tools/exp/rdf"
)

type OWLOntology struct{}

func (o *OWLOntology) SpecURI() string {
	return "http://www.w3.org/2002/07/owl#"
}

func (o *OWLOntology) Load() ([]rdf.RDFNode, error) {
	return nil, nil
}

func (o *OWLOntology) LoadAsAlias(s string) ([]rdf.RDFNode, error) {
	return nil, nil
}

func (o *OWLOntology) LoadElement(name string, payload map[string]interface{}) ([]rdf.RDFNode, error) {
	return nil, nil
}

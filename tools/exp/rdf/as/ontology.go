package as

import (
	"github.com/go-fed/activity/tools/exp/rdf"
)

type ActivityStreamsOntology struct{}

func (o *ActivityStreamsOntology) SpecURI() string {
	return "https://www.w3.org/ns/activitystreams"
}

func (o *ActivityStreamsOntology) Load() ([]rdf.RDFNode, error) {
	return nil, nil
}

func (o *ActivityStreamsOntology) LoadAsAlias(s string) ([]rdf.RDFNode, error) {
	return nil, nil
}

func (o *ActivityStreamsOntology) LoadElement(name string, payload map[string]interface{}) ([]rdf.RDFNode, error) {
	return nil, nil
}

package xsd

import (
	"fmt"
	"github.com/cjslep/activity/tools/exp/rdf"
	"net/url"
	"strings"
)

const (
	xmlSpec    = "http://www.w3.org/2001/XMLSchema#"
	anyURISpec = "anyURI"
)

type XMLOntology struct{}

func (o *XMLOntology) SpecURI() string {
	return xmlSpec
}

func (o *XMLOntology) Load() ([]rdf.RDFNode, error) {
	return o.LoadAsAlias("")
}

func (o *XMLOntology) LoadAsAlias(s string) ([]rdf.RDFNode, error) {
	return []rdf.RDFNode{
		&rdf.AliasedDelegate{
			Spec:     xmlSpec,
			Alias:    s,
			Name:     anyURISpec,
			Delegate: &anyURI{},
		},
	}, nil
}

func (o *XMLOntology) LoadSpecificAsAlias(alias, name string) ([]rdf.RDFNode, error) {
	switch name {
	case anyURISpec:
		return []rdf.RDFNode{
			&rdf.AliasedDelegate{
				Spec:     "",
				Alias:    "",
				Name:     alias,
				Delegate: &anyURI{},
			},
		}, nil
	}
	return nil, fmt.Errorf("xml ontology cannot find %q to alias to %q", name, alias)
}

func (o *XMLOntology) LoadElement(name string, payload map[string]interface{}) ([]rdf.RDFNode, error) {
	return nil, nil
}

func (o *XMLOntology) GetByName(name string) (rdf.RDFNode, error) {
	name = strings.TrimPrefix(name, o.SpecURI())
	switch name {
	case anyURISpec:
		return &anyURI{}, nil
	}
	return nil, fmt.Errorf("xsd ontology could not find node for name %s", name)
}

var _ rdf.RDFNode = &anyURI{}

type anyURI struct{}

func (a *anyURI) Enter(key string, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("xsd anyURI cannot be entered")
}

func (a *anyURI) Exit(key string, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("xsd anyURI cannot be exited")
}

func (a *anyURI) Apply(key string, value interface{}, ctx *rdf.ParsingContext) (bool, error) {
	v := ctx.Result.GetReference(xmlSpec)
	if len(v.Values[anyURISpec].Name) == 0 {
		u, err := url.Parse(xmlSpec + anyURISpec)
		if err != nil {
			return true, err
		}
		val := &rdf.VocabularyValue{
			Name:           "anyURI",
			URI:            u,
			DefinitionType: "*url.URL",
			Zero:           "&url.URL{}",
		}
		if err = v.SetValue(anyURISpec, val); err != nil {
			return true, err
		}
	}
	return true, nil
}

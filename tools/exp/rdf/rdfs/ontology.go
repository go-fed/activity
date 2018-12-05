package rdfs

import (
	"fmt"
	"github.com/cjslep/activity/tools/exp/rdf"
	"strings"
)

const (
	rdfsSpecURI       = "http://www.w3.org/2000/01/rdf-schema#"
	commentSpec       = "comment"
	domainSpec        = "domain"
	isDefinedBySpec   = "isDefinedBy"
	rangeSpec         = "range"
	subClassOfSpec    = "subClassOf"
	subPropertyOfSpec = "subPropertyOf"
)

type RDFSchemaOntology struct{}

func (o *RDFSchemaOntology) SpecURI() string {
	return rdfsSpecURI
}

func (o *RDFSchemaOntology) Load() ([]rdf.RDFNode, error) {
	return o.LoadAsAlias("")
}

func (o *RDFSchemaOntology) LoadAsAlias(s string) ([]rdf.RDFNode, error) {
	return []rdf.RDFNode{
		&rdf.AliasedDelegate{
			Spec:     rdfsSpecURI,
			Alias:    s,
			Name:     commentSpec,
			Delegate: &comment{},
		},
	}, nil
}

func (o *RDFSchemaOntology) LoadSpecificAsAlias(alias, name string) ([]rdf.RDFNode, error) {
	switch name {
	case commentSpec:
		return []rdf.RDFNode{
			&rdf.AliasedDelegate{
				Spec:     "",
				Alias:    "",
				Name:     alias,
				Delegate: &comment{},
			},
		}, nil
	case domainSpec:
		return []rdf.RDFNode{
			&rdf.AliasedDelegate{
				Spec:     "",
				Alias:    "",
				Name:     alias,
				Delegate: &domain{},
			},
		}, nil
	case isDefinedBySpec:
		return []rdf.RDFNode{
			&rdf.AliasedDelegate{
				Spec:     "",
				Alias:    "",
				Name:     alias,
				Delegate: &isDefinedBy{},
			},
		}, nil
	case rangeSpec:
		return []rdf.RDFNode{
			&rdf.AliasedDelegate{
				Spec:     "",
				Alias:    "",
				Name:     alias,
				Delegate: &ranges{},
			},
		}, nil
	case subClassOfSpec:
		return []rdf.RDFNode{
			&rdf.AliasedDelegate{
				Spec:     "",
				Alias:    "",
				Name:     alias,
				Delegate: &subClassOf{},
			},
		}, nil
	case subPropertyOfSpec:
		return []rdf.RDFNode{
			&rdf.AliasedDelegate{
				Spec:     "",
				Alias:    "",
				Name:     alias,
				Delegate: &comment{},
			},
		}, nil
	}
	return nil, fmt.Errorf("rdfs ontology cannot find %q to alias to %q", name, alias)
}

func (o *RDFSchemaOntology) LoadElement(name string, payload map[string]interface{}) ([]rdf.RDFNode, error) {
	return nil, nil
}

func (o *RDFSchemaOntology) GetByName(name string) (rdf.RDFNode, error) {
	name = strings.TrimPrefix(name, o.SpecURI())
	switch name {
	case commentSpec:
		return &comment{}, nil
	case domainSpec:
		return &domain{}, nil
	case isDefinedBySpec:
		return &isDefinedBy{}, nil
	case rangeSpec:
		return &ranges{}, nil
	case subClassOfSpec:
		return &subClassOf{}, nil
	case subPropertyOfSpec:
		return &subPropertyOf{}, nil
	}
	return nil, fmt.Errorf("rdfs ontology could not find node for name %s", name)
}

var _ rdf.RDFNode = &comment{}

type comment struct{}

func (n *comment) Enter(key string, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("rdfs comment cannot be entered")
}

func (n *comment) Exit(key string, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("rdfs comment cannot be exited")
}

func (n *comment) Apply(key string, value interface{}, ctx *rdf.ParsingContext) (bool, error) {
	note, ok := value.(string)
	if !ok {
		return true, fmt.Errorf("rdf comment not given string value")
	}
	if ctx.Current == nil {
		return true, fmt.Errorf("rdf comment given nil Current")
	}
	noteSetter, ok := ctx.Current.(rdf.NotesSetter)
	if !ok {
		return true, fmt.Errorf("rdf comment not given NotesSetter")
	}
	noteSetter.SetNotes(note)
	return true, nil
}

var _ rdf.RDFNode = &domain{}

type domain struct{}

func (d *domain) Enter(key string, ctx *rdf.ParsingContext) (bool, error) {
	// TODO
	return true, nil
}

func (d *domain) Exit(key string, ctx *rdf.ParsingContext) (bool, error) {
	// TODO
	return true, nil
}

func (d *domain) Apply(key string, value interface{}, ctx *rdf.ParsingContext) (bool, error) {
	// TODO
	return true, nil
}

var _ rdf.RDFNode = &isDefinedBy{}

type isDefinedBy struct{}

func (i *isDefinedBy) Enter(key string, ctx *rdf.ParsingContext) (bool, error) {
	// TODO
	return true, nil
}

func (i *isDefinedBy) Exit(key string, ctx *rdf.ParsingContext) (bool, error) {
	// TODO
	return true, nil
}

func (i *isDefinedBy) Apply(key string, value interface{}, ctx *rdf.ParsingContext) (bool, error) {
	// TODO
	return true, nil
}

var _ rdf.RDFNode = &ranges{}

type ranges struct{}

func (r *ranges) Enter(key string, ctx *rdf.ParsingContext) (bool, error) {
	// TODO
	return true, nil
}

func (r *ranges) Exit(key string, ctx *rdf.ParsingContext) (bool, error) {
	// TODO
	return true, nil
}

func (r *ranges) Apply(key string, value interface{}, ctx *rdf.ParsingContext) (bool, error) {
	// TODO
	return true, nil
}

var _ rdf.RDFNode = &subClassOf{}

type subClassOf struct{}

func (s *subClassOf) Enter(key string, ctx *rdf.ParsingContext) (bool, error) {
	ctx.Push()
	ctx.Current = &rdf.VocabularyReference{}
	return true, nil
}

func (s *subClassOf) Exit(key string, ctx *rdf.ParsingContext) (bool, error) {
	i := ctx.Current
	ctx.Pop()
	vr, ok := i.(*rdf.VocabularyReference)
	if !ok {
		return true, fmt.Errorf("rdfs subclassof exit did not get *rdf.VocabularyReference")
	}
	vt, ok := ctx.Current.(*rdf.VocabularyType)
	if !ok {
		return true, fmt.Errorf("rdf subclassof exit Current is not *rdf.VocabularyType")
	}
	vt.Extends = append(vt.Extends, *vr)
	return true, nil
}

func (s *subClassOf) Apply(key string, value interface{}, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("rdfs subclassof cannot be applied")
}

var _ rdf.RDFNode = &subPropertyOf{}

type subPropertyOf struct{}

func (s *subPropertyOf) Enter(key string, ctx *rdf.ParsingContext) (bool, error) {
	// TODO
	return true, nil
}

func (s *subPropertyOf) Exit(key string, ctx *rdf.ParsingContext) (bool, error) {
	// TODO
	return true, nil
}

func (s *subPropertyOf) Apply(key string, value interface{}, ctx *rdf.ParsingContext) (bool, error) {
	// TODO
	return true, nil
}

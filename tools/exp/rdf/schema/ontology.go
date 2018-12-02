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
	nameSpec       = "name"
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
			Delegate: &url{},
		},
		&rdf.AliasedDelegate{
			Spec:     schemaSpec,
			Alias:    s,
			Name:     nameSpec,
			Delegate: &name{},
		},
	}, nil
}

func (o *SchemaOntology) LoadSpecificAsAlias(alias, n string) ([]rdf.RDFNode, error) {
	switch n {
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
				Delegate: &url{},
			},
		}, nil
	case nameSpec:
		return []rdf.RDFNode{
			&rdf.AliasedDelegate{
				Spec:     "",
				Alias:    "",
				Name:     alias,
				Delegate: &name{},
			},
		}, nil
	}
	return nil, fmt.Errorf("schema ontology cannot find %q to alias to %q", n, alias)
}

func (o *SchemaOntology) LoadElement(name string, payload map[string]interface{}) ([]rdf.RDFNode, error) {
	return nil, nil
}

var _ rdf.RDFNode = &example{}

type example struct{}

func (e *example) Enter(key string, ctx *rdf.ParsingContext) (bool, error) {
	ctx.Push()
	ctx.Current = &rdf.VocabularyExample{}
	return true, nil
}

func (e *example) Exit(key string, ctx *rdf.ParsingContext) (bool, error) {
	ei := ctx.Current
	ctx.Pop()
	if ve, ok := ei.(*rdf.VocabularyExample); !ok {
		return true, fmt.Errorf("schema example did not pop a *VocabularyExample")
	} else if ea, ok := ctx.Current.(rdf.ExampleAdder); !ok {
		return true, fmt.Errorf("schema example not given an ExampleAdder")
	} else {
		ea.AddExample(ve)
	}
	return true, nil
}

func (e *example) Apply(key string, value interface{}, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("schema example cannot be applied")
}

var _ rdf.RDFNode = &mainEntity{}

type mainEntity struct{}

func (m *mainEntity) Enter(key string, ctx *rdf.ParsingContext) (bool, error) {
	ctx.Push()
	ctx.SetOnlyApplyThisNode(m)
	return true, nil
}

func (m *mainEntity) Exit(key string, ctx *rdf.ParsingContext) (bool, error) {
	// Save the example
	example := ctx.Current
	// Undo the Enter operations
	ctx.ResetOnlyApplyThisNode()
	ctx.Pop()
	// Set the example data
	if vEx, ok := ctx.Current.(*rdf.VocabularyExample); !ok {
		return true, fmt.Errorf("mainEntity exit not given a *VocabularyExample")
	} else {
		vEx.Example = example
	}
	return true, nil
}

func (m *mainEntity) Apply(key string, value interface{}, ctx *rdf.ParsingContext) (bool, error) {
	ctx.Current = value
	return true, nil
}

var _ rdf.RDFNode = &url{}

type url struct{}

func (u *url) Enter(key string, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("schema url cannot be entered")
}

func (u *url) Exit(key string, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("schema url cannot be exited")
}

func (u *url) Apply(key string, value interface{}, ctx *rdf.ParsingContext) (bool, error) {
	if urlString, ok := value.(string); !ok {
		return true, fmt.Errorf("schema url not given a string")
	} else if uriSetter, ok := ctx.Current.(rdf.URISetter); !ok {
		return true, fmt.Errorf("schema url not given a URISetter in context")
	} else {
		return true, uriSetter.SetURI(urlString)
	}
}

var _ rdf.RDFNode = &name{}

type name struct{}

func (n *name) Enter(key string, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("schema name cannot be entered")
}

func (n *name) Exit(key string, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("schema name cannot be exited")
}

func (n *name) Apply(key string, value interface{}, ctx *rdf.ParsingContext) (bool, error) {
	if s, ok := value.(string); !ok {
		return true, fmt.Errorf("schema name not given string")
	} else if ns, ok := ctx.Current.(rdf.NameSetter); !ok {
		return true, fmt.Errorf("schema name not given NameSetter in context")
	} else {
		ns.SetName(s)
		return true, nil
	}
}

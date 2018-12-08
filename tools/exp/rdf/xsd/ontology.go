package xsd

import (
	"fmt"
	"github.com/cjslep/activity/tools/exp/rdf"
	"net/url"
	"strings"
)

const (
	xmlSpec                = "http://www.w3.org/2001/XMLSchema#"
	anyURISpec             = "anyURI"
	dateTimeSpec           = "dateTime"
	floatSpec              = "float"
	stringSpec             = "string"
	booleanSpec            = "boolean"
	nonNegativeIntegerSpec = "nonNegativeInteger"
	durationSpec           = "duration"
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
		&rdf.AliasedDelegate{
			Spec:     xmlSpec,
			Alias:    s,
			Name:     dateTimeSpec,
			Delegate: &dateTime{},
		},
		&rdf.AliasedDelegate{
			Spec:     xmlSpec,
			Alias:    s,
			Name:     floatSpec,
			Delegate: &float{},
		},
		&rdf.AliasedDelegate{
			Spec:     xmlSpec,
			Alias:    s,
			Name:     stringSpec,
			Delegate: &xmlString{},
		},
		&rdf.AliasedDelegate{
			Spec:     xmlSpec,
			Alias:    s,
			Name:     booleanSpec,
			Delegate: &boolean{},
		},
		&rdf.AliasedDelegate{
			Spec:     xmlSpec,
			Alias:    s,
			Name:     nonNegativeIntegerSpec,
			Delegate: &nonNegativeInteger{},
		},
		&rdf.AliasedDelegate{
			Spec:     xmlSpec,
			Alias:    s,
			Name:     durationSpec,
			Delegate: &duration{},
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
	case dateTimeSpec:
		return []rdf.RDFNode{
			&rdf.AliasedDelegate{
				Spec:     "",
				Alias:    "",
				Name:     alias,
				Delegate: &dateTime{},
			},
		}, nil
	case floatSpec:
		return []rdf.RDFNode{
			&rdf.AliasedDelegate{
				Spec:     "",
				Alias:    "",
				Name:     alias,
				Delegate: &float{},
			},
		}, nil
	case stringSpec:
		return []rdf.RDFNode{
			&rdf.AliasedDelegate{
				Spec:     "",
				Alias:    "",
				Name:     alias,
				Delegate: &xmlString{},
			},
		}, nil
	case booleanSpec:
		return []rdf.RDFNode{
			&rdf.AliasedDelegate{
				Spec:     "",
				Alias:    "",
				Name:     alias,
				Delegate: &boolean{},
			},
		}, nil
	case nonNegativeIntegerSpec:
		return []rdf.RDFNode{
			&rdf.AliasedDelegate{
				Spec:     "",
				Alias:    "",
				Name:     alias,
				Delegate: &nonNegativeInteger{},
			},
		}, nil
	case durationSpec:
		return []rdf.RDFNode{
			&rdf.AliasedDelegate{
				Spec:     "",
				Alias:    "",
				Name:     alias,
				Delegate: &duration{},
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
	case dateTimeSpec:
		return &dateTime{}, nil
	case floatSpec:
		return &float{}, nil
	case stringSpec:
		return &xmlString{}, nil
	case booleanSpec:
		return &boolean{}, nil
	case nonNegativeIntegerSpec:
		return &nonNegativeInteger{}, nil
	case durationSpec:
		return &duration{}, nil
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
			Name:           anyURISpec,
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

var _ rdf.RDFNode = &dateTime{}

type dateTime struct{}

func (d *dateTime) Enter(key string, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("xsd dateTime cannot be entered")
}

func (d *dateTime) Exit(key string, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("xsd dateTime cannot be exited")
}

func (d *dateTime) Apply(key string, value interface{}, ctx *rdf.ParsingContext) (bool, error) {
	v := ctx.Result.GetReference(xmlSpec)
	if len(v.Values[dateTimeSpec].Name) == 0 {
		u, err := url.Parse(xmlSpec + dateTimeSpec)
		if err != nil {
			return true, err
		}
		val := &rdf.VocabularyValue{
			Name:           dateTimeSpec,
			URI:            u,
			DefinitionType: "time.Time",
			Zero:           "&time.Time{}",
		}
		if err = v.SetValue(dateTimeSpec, val); err != nil {
			return true, err
		}
	}
	return true, nil
}

var _ rdf.RDFNode = &float{}

type float struct{}

func (f *float) Enter(key string, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("xsd float cannot be entered")
}

func (f *float) Exit(key string, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("xsd float cannot be exited")
}

func (f *float) Apply(key string, value interface{}, ctx *rdf.ParsingContext) (bool, error) {
	v := ctx.Result.GetReference(xmlSpec)
	if len(v.Values[floatSpec].Name) == 0 {
		u, err := url.Parse(xmlSpec + floatSpec)
		if err != nil {
			return true, err
		}
		val := &rdf.VocabularyValue{
			Name:           floatSpec,
			URI:            u,
			DefinitionType: "float32",
			Zero:           "0.0",
		}
		if err = v.SetValue(floatSpec, val); err != nil {
			return true, err
		}
	}
	return true, nil
}

var _ rdf.RDFNode = &xmlString{}

type xmlString struct{}

func (*xmlString) Enter(key string, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("xsd string cannot be entered")
}

func (*xmlString) Exit(key string, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("xsd string cannot be exited")
}

func (*xmlString) Apply(key string, value interface{}, ctx *rdf.ParsingContext) (bool, error) {
	v := ctx.Result.GetReference(xmlSpec)
	if len(v.Values[stringSpec].Name) == 0 {
		u, err := url.Parse(xmlSpec + stringSpec)
		if err != nil {
			return true, err
		}
		val := &rdf.VocabularyValue{
			Name:           stringSpec,
			URI:            u,
			DefinitionType: "string",
			Zero:           "\"\"",
		}
		if err = v.SetValue(stringSpec, val); err != nil {
			return true, err
		}
	}
	return true, nil
}

var _ rdf.RDFNode = &boolean{}

type boolean struct{}

func (*boolean) Enter(key string, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("xsd boolean cannot be entered")
}

func (*boolean) Exit(key string, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("xsd boolean cannot be exited")
}

func (*boolean) Apply(key string, value interface{}, ctx *rdf.ParsingContext) (bool, error) {
	v := ctx.Result.GetReference(xmlSpec)
	if len(v.Values[booleanSpec].Name) == 0 {
		u, err := url.Parse(xmlSpec + booleanSpec)
		if err != nil {
			return true, err
		}
		val := &rdf.VocabularyValue{
			Name:           booleanSpec,
			URI:            u,
			DefinitionType: "string",
			Zero:           "\"\"",
		}
		if err = v.SetValue(booleanSpec, val); err != nil {
			return true, err
		}
	}
	return true, nil
}

var _ rdf.RDFNode = &nonNegativeInteger{}

type nonNegativeInteger struct{}

func (*nonNegativeInteger) Enter(key string, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("xsd nonNegativeInteger cannot be entered")
}

func (*nonNegativeInteger) Exit(key string, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("xsd nonNegativeInteger cannot be exited")
}

func (*nonNegativeInteger) Apply(key string, value interface{}, ctx *rdf.ParsingContext) (bool, error) {
	v := ctx.Result.GetReference(xmlSpec)
	if len(v.Values[nonNegativeIntegerSpec].Name) == 0 {
		u, err := url.Parse(xmlSpec + nonNegativeIntegerSpec)
		if err != nil {
			return true, err
		}
		val := &rdf.VocabularyValue{
			Name:           nonNegativeIntegerSpec,
			URI:            u,
			DefinitionType: "int",
			Zero:           "0",
		}
		if err = v.SetValue(nonNegativeIntegerSpec, val); err != nil {
			return true, err
		}
	}
	return true, nil
}

var _ rdf.RDFNode = &duration{}

type duration struct{}

func (*duration) Enter(key string, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("xsd duration cannot be entered")
}

func (*duration) Exit(key string, ctx *rdf.ParsingContext) (bool, error) {
	return true, fmt.Errorf("xsd duration cannot be exited")
}

func (*duration) Apply(key string, value interface{}, ctx *rdf.ParsingContext) (bool, error) {
	v := ctx.Result.GetReference(xmlSpec)
	if len(v.Values[durationSpec].Name) == 0 {
		u, err := url.Parse(xmlSpec + durationSpec)
		if err != nil {
			return true, err
		}
		val := &rdf.VocabularyValue{
			Name:           durationSpec,
			URI:            u,
			DefinitionType: "time.Duration",
			Zero:           "time.Duration(0)",
		}
		if err = v.SetValue(durationSpec, val); err != nil {
			return true, err
		}
	}
	return true, nil
}

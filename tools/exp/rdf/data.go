package rdf

import (
	"bytes"
	"fmt"
	"github.com/cjslep/activity/tools/exp/codegen"
	"net/url"
)

// ParsedVocabulary is the internal data structure produced after parsing the
// definition of an ActivityStream vocabulary. It is the intermediate
// understanding of the specification in the context of certain ontologies. It
// also contains additional scratch space for use by the parser.
//
// At the end of parsing, the ParsedVocabulary is not guaranteed to be
// semantically valid, just that the parser resolved all important ontological
// details.
type ParsedVocabulary struct {
	Vocab      Vocabulary
	References map[string]*Vocabulary
}

func (p *ParsedVocabulary) GetReference(uri string) *Vocabulary {
	if p.References == nil {
		p.References = make(map[string]*Vocabulary, 0)
	}
	if _, ok := p.References[uri]; !ok {
		p.References[uri] = &Vocabulary{}
	}
	return p.References[uri]
}

func (p ParsedVocabulary) String() string {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("Vocab:\n%s", p.Vocab))
	for k, v := range p.References {
		b.WriteString(fmt.Sprintf("Reference %s:\n\t%s\n", k, v))
	}
	return b.String()
}

// Vocabulary contains the type, property, and value definitions for a single
// ActivityStreams or extension vocabulary.
type Vocabulary struct {
	Name       string
	Types      map[string]VocabularyType
	Properties map[string]VocabularyProperty
	Values     map[string]VocabularyValue
}

func (v Vocabulary) GetName() string {
	return v.Name
}

func (v Vocabulary) String() string {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("Vocabulary %q\n", v.Name))
	for k, v := range v.Types {
		b.WriteString(fmt.Sprintf("Type %s:\n\t%s\n", k, v))
	}
	for k, v := range v.Properties {
		b.WriteString(fmt.Sprintf("Property %s:\n\t%s\n", k, v))
	}
	for k, v := range v.Values {
		b.WriteString(fmt.Sprintf("Value %s:\n\t%s\n", k, v))
	}
	return b.String()
}

func (v *Vocabulary) SetType(name string, a *VocabularyType) error {
	if v.Types == nil {
		v.Types = make(map[string]VocabularyType, 1)
	}
	if _, has := v.Types[name]; has {
		return fmt.Errorf("name %q already exists for vocabulary Types", name)
	}
	v.Types[name] = *a
	return nil
}

func (v *Vocabulary) SetProperty(name string, a *VocabularyProperty) error {
	if v.Properties == nil {
		v.Properties = make(map[string]VocabularyProperty, 1)
	}
	if _, has := v.Properties[name]; has {
		return fmt.Errorf("name already exists for vocabulary Properties")
	}
	v.Properties[name] = *a
	return nil
}

func (v *Vocabulary) SetValue(name string, a *VocabularyValue) error {
	if v.Values == nil {
		v.Values = make(map[string]VocabularyValue, 1)
	}
	if _, has := v.Values[name]; has {
		return fmt.Errorf("name already exists for vocabulary Values")
	}
	v.Values[name] = *a
	return nil
}

// VocabularyValue represents a value type that properties can take on.
type VocabularyValue struct {
	Name           string
	URI            *url.URL
	DefinitionType string
	Zero           string
	SerializeFn    *codegen.Function
	DeserializeFn  *codegen.Function
	LessFn         *codegen.Function
}

func (v VocabularyValue) String() string {
	return fmt.Sprintf("Value=%s,%s,%s,%s", v.Name, v.URI, v.DefinitionType, v.Zero)
}

func (v *VocabularyValue) SetName(s string) {
	v.Name = s
}

func (v VocabularyValue) GetName() string {
	return v.Name
}

func (v *VocabularyValue) SetURI(s string) error {
	var e error
	v.URI, e = url.Parse(s)
	return e
}

var (
	_ NameSetter = &VocabularyValue{}
	_ NameGetter = &VocabularyValue{}
	_ URISetter  = &VocabularyValue{}
)

// VocabularyType represents a single ActivityStream type in a vocabulary.
type VocabularyType struct {
	Name              string
	URI               *url.URL
	Notes             string
	DisjointWith      []VocabularyReference
	Extends           []VocabularyReference
	Examples          []VocabularyExample
	Properties        []VocabularyReference
	WithoutProperties []VocabularyReference
}

func (v VocabularyType) String() string {
	return fmt.Sprintf("Type=%s,%s,%s\n\tDJW=%s\n\tExt=%s\n\tEx=%s", v.Name, v.URI, v.Notes, v.DisjointWith, v.Extends, v.Examples)
}

func (v *VocabularyType) SetName(s string) {
	v.Name = s
}

func (v VocabularyType) GetName() string {
	return v.Name
}

func (v VocabularyType) TypeName() string {
	return v.Name
}

func (v *VocabularyType) SetURI(s string) error {
	var e error
	v.URI, e = url.Parse(s)
	return e
}

func (v *VocabularyType) SetNotes(s string) {
	v.Notes = s
}

func (v *VocabularyType) AddExample(e *VocabularyExample) {
	v.Examples = append(v.Examples, *e)
}

var (
	_ NameSetter   = &VocabularyType{}
	_ NameGetter   = &VocabularyType{}
	_ URISetter    = &VocabularyType{}
	_ NotesSetter  = &VocabularyType{}
	_ ExampleAdder = &VocabularyType{}
)

// VocabularyProperty represents a single ActivityStream property type in a
// vocabulary.
type VocabularyProperty struct {
	Name           string
	URI            *url.URL
	Notes          string
	Domain         []VocabularyReference
	Range          []VocabularyReference
	DoesNotApplyTo []VocabularyReference
	Examples       []VocabularyExample
	// SubpropertyOf is ignorable as long as data is set up correctly TODO: Is this still correct?
	SubpropertyOf      VocabularyReference // Must be a VocabularyProperty
	Functional         bool
	NaturalLanguageMap bool
}

func (v VocabularyProperty) String() string {
	return fmt.Sprintf("Property=%s,%s,%s\n\tD=%s\n\tR=%s\n\tEx=%s\n\tSub=%s\n\tDNApply=%s\n\tfunc=%t,natLangMap=%t", v.Name, v.URI, v.Notes, v.Domain, v.Range, v.Examples, v.SubpropertyOf, v.DoesNotApplyTo, v.Functional, v.NaturalLanguageMap)
}

func (v *VocabularyProperty) SetName(s string) {
	v.Name = s
}

func (v VocabularyProperty) GetName() string {
	return v.Name
}

func (v VocabularyProperty) PropertyName() string {
	return v.Name
}

func (v *VocabularyProperty) SetURI(s string) error {
	var e error
	v.URI, e = url.Parse(s)
	return e
}

func (v *VocabularyProperty) SetNotes(s string) {
	v.Notes = s
}

func (v *VocabularyProperty) AddExample(e *VocabularyExample) {
	v.Examples = append(v.Examples, *e)
}

var (
	_ NameSetter   = &VocabularyProperty{}
	_ NameGetter   = &VocabularyProperty{}
	_ URISetter    = &VocabularyProperty{}
	_ NotesSetter  = &VocabularyProperty{}
	_ ExampleAdder = &VocabularyProperty{}
)

// VocabularyExample documents an Example for an ActivityStream type or property
// in the vocabulary.
type VocabularyExample struct {
	Name    string
	URI     *url.URL
	Example interface{}
}

func (v VocabularyExample) String() string {
	return fmt.Sprintf("VocabularyExample: %s,%s,%s", v.Name, v.URI, v.Example)
}

func (v *VocabularyExample) SetName(s string) {
	v.Name = s
}

func (v VocabularyExample) GetName() string {
	return v.Name
}

func (v *VocabularyExample) SetURI(s string) error {
	var e error
	v.URI, e = url.Parse(s)
	return e
}

var (
	_ NameSetter = &VocabularyExample{}
	_ NameGetter = &VocabularyExample{}
	_ URISetter  = &VocabularyExample{}
)

// VocabularyReference refers to another Vocabulary reference, either a
// VocabularyType, VocabularyValue, or a VocabularyProperty. It may refer to
// another Vocabulary's type or property entirely.
type VocabularyReference struct {
	Name  string
	URI   *url.URL
	Vocab string // If present, must match key in ParsedVocabulary.References
}

func (v VocabularyReference) String() string {
	return fmt.Sprintf("VocabularyReference: %s,%s,%s", v.Name, v.URI, v.Vocab)
}

func (v *VocabularyReference) SetName(s string) {
	v.Name = s
}

func (v VocabularyReference) GetName() string {
	return v.Name
}

func (v *VocabularyReference) SetURI(s string) error {
	var e error
	v.URI, e = url.Parse(s)
	return e
}

var (
	_ NameSetter = &VocabularyReference{}
	_ NameGetter = &VocabularyReference{}
	_ URISetter  = &VocabularyReference{}
)

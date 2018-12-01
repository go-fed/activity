package rdf

import (
	"fmt"
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
	References map[string]Vocabulary
}

// Vocabulary contains the type, property, and value definitions for a single
// ActivityStreams or extension vocabulary.
type Vocabulary struct {
	Types      map[string]VocabularyType
	Properties map[string]VocabularyProperty
	Values     map[string]VocabularyValue
}

func (v *Vocabulary) SetType(name string, a *VocabularyType) error {
	if v.Types == nil {
		v.Types = make(map[string]VocabularyType, 1)
	}
	if _, has := v.Types[name]; has {
		return fmt.Errorf("name already exists for vocabulary Types")
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
}

func (v *VocabularyValue) SetName(s string) {
	v.Name = s
}

func (v *VocabularyValue) SetURI(s string) error {
	var e error
	v.URI, e = url.Parse(s)
	return e
}

var (
	_ NameSetter = &VocabularyValue{}
	_ URISetter  = &VocabularyValue{}
)

// VocabularyType represents a single ActivityStream type in a vocabulary.
type VocabularyType struct {
	Name              string
	URI               *url.URL
	Notes             string
	DisjointWith      []VocabularyReference
	Extends           []VocabularyReference // TODO: Object improperly extends Link
	Properties        []VocabularyReference // TODO: Check for duplication
	WithoutProperties []VocabularyReference // TODO: Missing for IntransitiveActivity
	Examples          []VocabularyExample
}

func (v *VocabularyType) SetName(s string) {
	v.Name = s
}

func (v *VocabularyType) SetURI(s string) error {
	var e error
	v.URI, e = url.Parse(s)
	return e
}

func (v *VocabularyType) SetNotes(s string) {
	v.Notes = s
}

var (
	_ NameSetter  = &VocabularyType{}
	_ URISetter   = &VocabularyType{}
	_ NotesSetter = &VocabularyType{}
)

// VocabularyProperty represents a single ActivityStream property type in a
// vocabulary.
type VocabularyProperty struct {
	Name   string
	URI    *url.URL
	Notes  string
	Domain []VocabularyReference
	Range  []VocabularyReference
	// SubpropertyOf is ignorable as long as data is set up correctly TODO: Is this still correct?
	SubpropertyOf      VocabularyReference // Must be a VocabularyProperty
	Functional         bool
	NaturalLanguageMap bool
}

func (v *VocabularyProperty) SetName(s string) {
	v.Name = s
}

func (v *VocabularyProperty) SetURI(s string) error {
	var e error
	v.URI, e = url.Parse(s)
	return e
}

func (v *VocabularyProperty) SetNotes(s string) {
	v.Notes = s
}

var (
	_ NameSetter  = &VocabularyProperty{}
	_ URISetter   = &VocabularyProperty{}
	_ NotesSetter = &VocabularyProperty{}
)

// VocabularyExample documents an Example for an ActivityStream type or property
// in the vocabulary.
type VocabularyExample struct {
	Name    string
	URI     *url.URL
	Example map[string]interface{}
}

func (v *VocabularyExample) SetName(s string) {
	v.Name = s
}

func (v *VocabularyExample) SetURI(s string) error {
	var e error
	v.URI, e = url.Parse(s)
	return e
}

var (
	_ NameSetter = &VocabularyExample{}
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

func (v *VocabularyReference) SetName(s string) {
	v.Name = s
}

func (v *VocabularyReference) SetURI(s string) error {
	var e error
	v.URI, e = url.Parse(s)
	return e
}

var (
	_ NameSetter = &VocabularyReference{}
	_ URISetter  = &VocabularyReference{}
)
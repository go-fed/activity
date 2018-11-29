package rdf

import (
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

// VocabularyValue represents a value type that properties can take on.
type VocabularyValue struct {
	Name           string
	URI            *url.URL
	DefinitionType string
	Zero           string
}

// VocabularyType represents a single ActivityStream type in a vocabulary.
type VocabularyType struct {
	Name              string
	URI               *url.URL
	Notes             string
	DisjointWith      []VocabularyReference
	Extends           []VocabularyReference
	Properties        []VocabularyReference // TODO: Check for duplication
	WithoutProperties []VocabularyReference // TODO: Missing for IntransitiveActivity
	Examples          []VocabularyExample
}

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

// VocabularyExample documents an Example for an ActivityStream type or property
// in the vocabulary.
type VocabularyExample struct {
	Name    string
	URI     *url.URL
	Example map[string]interface{}
}

// VocabularyReference refers to another Vocabulary reference, either a
// VocabularyType, VocabularyValue, or a VocabularyProperty. It may refer to
// another Vocabulary's type or property entirely.
type VocabularyReference struct {
	Name  string
	URI   *url.URL
	Vocab string // If present, must match key in ParsedVocabulary.References
}

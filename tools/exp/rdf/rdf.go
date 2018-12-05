package rdf

import (
	"fmt"
	"strings"
)

const (
	ALIAS_DELIMITER = ":"
	HTTP            = "http"
	HTTPS           = "https"
	ID              = "@id"
)

// IsKeyApplicable returns true if the key has a spec or alias prefix and the
// property is equal to the desired name.
//
// If 'alias' is an empty string, it is ignored.
func IsKeyApplicable(key, spec, alias, name string) bool {
	if key == spec+name {
		return true
	} else if len(alias) > 0 {
		strs := strings.Split(key, ALIAS_DELIMITER)
		if len(strs) > 1 && strs[0] != HTTP && strs[0] != HTTPS {
			return strs[0] == alias && strs[1] == name
		}
	}
	return false
}

// SplitAlias splits a possibly-aliased string, without splitting on the colon
// if it is part of the http or https spec.
func SplitAlias(s string) []string {
	strs := strings.Split(s, ALIAS_DELIMITER)
	if len(strs) == 1 {
		return strs
	} else if strs[0] == HTTP || strs[0] == HTTPS {
		return []string{s}
	} else {
		return strs
	}
}

// Ontology returns different RDF "actions" or "handlers" that are able to
// interpret the schema definitions as actions upon a set of data, specific
// for this ontology.
type Ontology interface {
	// SpecURI refers to the URI location of this ontology.
	// TODO: Handle both http and https.
	SpecURI() string

	// The Load methods deal with determining how best to apply an ontology
	// based on the context specified by the data. This is before the data
	// is actually processed.

	// Load loads the entire ontology.
	Load() ([]RDFNode, error)
	// LoadAsAlias loads the entire ontology with a specific alias.
	LoadAsAlias(s string) ([]RDFNode, error)
	// LoadSpecificAsAlias loads a specific element of the ontology by
	// being able to handle the specific alias as its name instead.
	LoadSpecificAsAlias(alias, name string) ([]RDFNode, error)
	// LoadElement loads a specific element of the ontology based on the
	// object definition.
	LoadElement(name string, payload map[string]interface{}) ([]RDFNode, error)

	// The Get methods deal with determining how best to apply an ontology
	// during processing. This is a result of certain nodes having highly
	// contextual effects.

	// GetByName returns an RDFNode associated with the given name. Note
	// that the name may either be fully-qualified (in the case it was not
	// aliased) or it may be just the element name (in the case it was
	// aliased).
	GetByName(name string) (RDFNode, error)
}

// aliasedNode represents a context element that has a special reserved alias.
type aliasedNode struct {
	Alias string
	Nodes []RDFNode
}

// RDFRegistry manages the different ontologies needed to determine the
// generated Go code.
type RDFRegistry struct {
	ontologies   map[string]Ontology
	aliases      map[string]string
	aliasedNodes map[string]aliasedNode
}

// NewRDFRegistry returns a new RDFRegistry.
func NewRDFRegistry() *RDFRegistry {
	return &RDFRegistry{
		ontologies:   make(map[string]Ontology),
		aliases:      make(map[string]string),
		aliasedNodes: make(map[string]aliasedNode),
	}
}

// setAlias sets an alias for a string.
func (r *RDFRegistry) setAlias(alias, s string) error {
	if _, ok := r.aliases[alias]; ok {
		return fmt.Errorf("already have alias for %s", alias)
	}
	r.aliases[alias] = s
	return nil
}

// setAliasedNode sets an alias for a node.
func (r *RDFRegistry) setAliasedNode(alias string, nodes []RDFNode) error {
	if _, ok := r.aliasedNodes[alias]; ok {
		return fmt.Errorf("already have aliased node for %s", alias)
	}
	r.aliasedNodes[alias] = aliasedNode{
		Alias: alias,
		Nodes: nodes,
	}
	return nil
}

// getOngology resolves an alias to a particular Ontology.
func (r *RDFRegistry) getOntology(alias string) (Ontology, error) {
	if ontologyName, ok := r.aliases[alias]; !ok {
		return nil, fmt.Errorf("missing alias %q", alias)
	} else if ontology, ok := r.ontologies[ontologyName]; !ok {
		return nil, fmt.Errorf("alias %q resolved but missing ontology with name %q", alias, ontologyName)
	} else {
		return ontology, nil
	}
}

// AddOntology adds an RDF ontology to the registry.
func (r *RDFRegistry) AddOntology(o Ontology) error {
	if r.ontologies == nil {
		r.ontologies = make(map[string]Ontology, 1)
	}
	s := o.SpecURI()
	if _, ok := r.ontologies[s]; ok {
		return fmt.Errorf("ontology already registered for %q", s)
	}
	r.ontologies[s] = o
	return nil
}

// getFor gets RDFKeyers based on a context's string.
//
// Package public.
func (r *RDFRegistry) getFor(s string) (n []RDFNode, e error) {
	ontology, ok := r.ontologies[s]
	if !ok {
		e = fmt.Errorf("no ontology for %s", s)
		return
	}
	return ontology.Load()
}

// getForAliased gets RDFKeyers based on a context's string.
//
// Private to this file.
func (r *RDFRegistry) getForAliased(alias, s string) (n []RDFNode, e error) {
	ontology, ok := r.ontologies[s]
	if !ok {
		e = fmt.Errorf("no ontology for %s", s)
		return
	}
	return ontology.LoadAsAlias(alias)
}

// getAliased gets RDFKeyers based on a context string and its
// alias.
//
// Package public.
func (r *RDFRegistry) getAliased(alias, s string) (n []RDFNode, e error) {
	strs := SplitAlias(s)
	if len(strs) == 1 {
		if e = r.setAlias(alias, s); e != nil {
			return
		}
		return r.getForAliased(alias, s)
	} else if len(strs) == 2 {
		var o Ontology
		o, e = r.getOntology(strs[0])
		if e != nil {
			return
		}
		n, e = o.LoadSpecificAsAlias(alias, strs[1])
		return
	} else {
		e = fmt.Errorf("too many delimiters in %s", s)
		return
	}
}

// getAliasedObject gets RDFKeyers based on a context object and
// its alias and definition.
//
// Package public.
func (r *RDFRegistry) getAliasedObject(alias string, object map[string]interface{}) (n []RDFNode, e error) {
	raw, ok := object[ID]
	if !ok {
		e = fmt.Errorf("aliased object does not have %s value", ID)
		return
	}
	if element, ok := raw.(string); !ok {
		e = fmt.Errorf("element in getAliasedObject must be a string")
		return
	} else {
		strs := SplitAlias(element)
		if len(strs) == 1 {
			n, e = r.getFor(strs[0])
		} else if len(strs) == 2 {
			var o Ontology
			o, e = r.getOntology(strs[0])
			if e != nil {
				return
			}
			n, e = o.LoadElement(alias, object)
		}
		if e != nil {
			return
		}
		e = r.setAliasedNode(alias, n)
		return
	}
}

// getNode fetches a node based on a string. It may be aliased or not.
//
// Package public.
func (r *RDFRegistry) getNode(s string) (n RDFNode, e error) {
	strs := SplitAlias(s)
	if len(strs) == 2 {
		if ontName, ok := r.aliases[strs[0]]; !ok {
			e = fmt.Errorf("no alias to ontology for %s", strs[0])
			return
		} else if ontology, ok := r.ontologies[ontName]; !ok {
			e = fmt.Errorf("no ontology named %s for alias %s", ontName, strs[0])
			return
		} else {
			n, e = ontology.GetByName(strs[1])
			return
		}
	} else if len(strs) == 1 {
		for _, ontology := range r.ontologies {
			if strings.HasPrefix(s, ontology.SpecURI()) {
				n, e = ontology.GetByName(s)
				return
			}
		}
		e = fmt.Errorf("getNode could not find ontology for %s", s)
		return
	} else {
		e = fmt.Errorf("getNode given unhandled node name: %s", s)
		return
	}
}

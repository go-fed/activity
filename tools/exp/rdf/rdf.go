package rdf

import (
	"fmt"
	"strings"
	"sync"
)

const (
	ALIAS_DELIMITER = ":"
	ID              = "@id"
)

// Ontology returns different RDF "actions" or "handlers" that are able to
// interpret the schema definitions as actions upon a set of data, specific
// for this ontology.
type Ontology interface {
	// String representation of this ontology.
	String() string
	// Load loads the entire ontology.
	Load() ([]RDFNode, error)
	// Load loads the entire ontology with a specific alias.
	LoadAsAlias(s string) ([]RDFNode, error)
	// LoadElement loads a specific element of the ontology by name. The
	// payload may be nil.
	LoadElement(name string, payload map[string]interface{}) ([]RDFNode, error)
}

// aliasedNode represents a context element that has a special reserved alias.
type aliasedNode struct {
	Alias string
	Nodes []RDFNode
}

// RDFRegistry implements RDFGetter and manages the different ontologies needed
// to determine the generated Go code.
type RDFRegistry struct {
	ontologies   map[string]Ontology
	aliases      map[string]string
	aliasedNodes map[string]aliasedNode
	mu           sync.Mutex
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

// loadElement will handle the aliasing of an ontology and retrieve the nodes
// required for a specific element within that ontology.
func (r *RDFRegistry) loadElement(alias, element string, payload map[string]interface{}) (n []RDFNode, e error) {
	if ontName, ok := r.aliases[alias]; !ok {
		e = fmt.Errorf("no alias to ontology for %s", alias)
		return
	} else if ontology, ok := r.ontologies[ontName]; !ok {
		e = fmt.Errorf("no ontology named %s for alias %s", ontName, alias)
		return
	} else {
		n, e = ontology.LoadElement(element, payload)
		return
	}
}

// AddOntology adds an RDF ontology to the registry.
func (r *RDFRegistry) AddOntology(s string, o Ontology) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.ontologies == nil {
		r.ontologies = make(map[string]Ontology, 1)
	}
	if _, ok := r.ontologies[s]; ok {
		return fmt.Errorf("ontology already registered for %q", s)
	}
	r.ontologies[s] = o
	return nil
}

// GetFor gets RDFKeyers and RDFValuers based on a context's string.
//
// Implements RDFGetter.
func (r *RDFRegistry) GetFor(s string) (n []RDFNode, e error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	ontology, ok := r.ontologies[s]
	if !ok {
		e = fmt.Errorf("no ontology for %s", s)
		return
	}
	return ontology.Load()
}

// GetAliased gets RDFKeyers and RDFValuers based on a context string and its
// alias.
//
// Implements RDFGetter.
func (r *RDFRegistry) GetAliased(alias, s string) (n []RDFNode, e error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	strs := strings.Split(s, ALIAS_DELIMITER)
	if len(strs) == 1 {
		if e = r.setAlias(alias, s); e != nil {
			return
		}
		return r.GetFor(s)
	} else if len(strs) == 2 {
		var o Ontology
		o, e = r.getOntology(strs[0])
		if e != nil {
			return
		}
		n, e = o.LoadElement(strs[1], nil)
		return
	} else {
		e = fmt.Errorf("too many delimiters in %s", s)
		return
	}
}

// GetAliasedObject gets RDFKeyers and RDFValuers based on a context object and
// its alias and definition.
//
// Implements RDFGetter.
func (r *RDFRegistry) GetAliasedObject(alias string, object map[string]interface{}) (n []RDFNode, e error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var iObj interface{} = object
	if obj, ok := iObj.(map[string]string); !ok {
		e = fmt.Errorf("object in GetAliasedObject must be a map[string]string")
		return
	} else {
		element, ok := obj[ID]
		if !ok {
			e = fmt.Errorf("aliased object does not have %s value", ID)
			return
		}
		var nodes []RDFNode
		strs := strings.Split(element, ALIAS_DELIMITER)
		if len(strs) == 1 {
			n, e = r.GetFor(strs[0])
		} else if len(strs) == 2 {
			var o Ontology
			o, e = r.getOntology(strs[0])
			if e != nil {
				return
			}
			n, e = o.LoadElement(strs[1], object)
			return
		}
		if e != nil {
			return
		}
		if e = r.setAliasedNode(alias, nodes); e != nil {
			return
		}
		return
	}
}

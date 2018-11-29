package rdf

import (
	"fmt"
)

const (
	JSON_LD_CONTEXT = "@context"
)

// JSONLD is an alias for the generic map of keys to interfaces, presumably
// parsed from a JSON-encoded context definition file.
type JSONLD map[string]interface{}

// ParsingContext contains the results of the parsing as well as scratch space
// required for RDFNodes to be able to statefully apply changes.
type ParsingContext struct {
	Result  *ParsedVocabulary
	Current interface{}
}

type NameSetter interface {
	SetName(string)
}

type URISetter interface {
	SetURI(string) error
}

type NotesSetter interface {
	SetNotes(string)
}

// RDFNode is able to operate on a specific key if it applies towards its
// ontology (determined at creation time). It applies the value in its own
// specific implementation on the context.
type RDFNode interface {
	Enter(key string, ctx *ParsingContext) (bool, error)
	Exit(key string, ctx *ParsingContext) (bool, error)
	Apply(key string, value interface{}, ctx *ParsingContext) (bool, error)
}

// ParseVocabulary parses the specified input as an ActivityStreams context that
// specifies a Core, Extended, or Extension vocabulary.
func ParseVocabulary(registry *RDFRegistry, input JSONLD) (vocabulary *ParsedVocabulary, err error) {
	var nodes []RDFNode
	nodes, err = parseJSONLDContext(registry, input)
	if err != nil {
		return
	}
	vocabulary = &ParsedVocabulary{}
	ctx := &ParsingContext{
		Result: vocabulary,
	}
	err = apply(nodes, input, ctx)
	return
}

// apply takes a specification input to populate the ParsingContext, based on
// the capabilities of the RDFNodes created from ontologies.
func apply(nodes []RDFNode, input JSONLD, ctx *ParsingContext) error {
	for k, v := range input {
		// Skip the context as it has already been parsed to create the
		// nodes.
		if k == JSON_LD_CONTEXT {
			continue
		}
		if mapValue, ok := v.(map[string]interface{}); ok {
			if err := enterFirstNode(nodes, k, ctx); err != nil {
				return err
			} else if err = apply(nodes, mapValue, ctx); err != nil {
				return err
			} else if err = exitFirstNode(nodes, k, ctx); err != nil {
				return err
			}
		} else if arrValue, ok := v.([]interface{}); ok {
			for _, val := range arrValue {
				// First, enter for this key
				if err := enterFirstNode(nodes, k, ctx); err != nil {
					return err
				}
				// Recur or handle the value as necessary.
				if mapValue, ok := val.(map[string]interface{}); ok {
					if err := apply(nodes, mapValue, ctx); err != nil {
						return err
					}
				} else if err := applyFirstNode(nodes, k, val, ctx); err != nil {
					return err
				}
				// Finally, exit for this key
				if err := exitFirstNode(nodes, k, ctx); err != nil {
					return err
				}
			}
		} else if err := applyFirstNode(nodes, k, v, ctx); err != nil {
			return err
		}
	}
	return nil
}

// enterFirstNode will Enter the first RDFNode that returns true or an error.
func enterFirstNode(nodes []RDFNode, key string, ctx *ParsingContext) error {
	for _, node := range nodes {
		if applied, err := node.Enter(key, ctx); applied {
			return nil
		} else if err != nil {
			return err
		}
	}
	return nil
}

// exitFirstNode will Exit the first RDFNode that returns true or an error.
func exitFirstNode(nodes []RDFNode, key string, ctx *ParsingContext) error {
	for _, node := range nodes {
		if applied, err := node.Exit(key, ctx); applied {
			return nil
		} else if err != nil {
			return err
		}
	}
	return nil
}

// applyFirstNode will Apply the first RDFNode that returns true or an error.
func applyFirstNode(nodes []RDFNode, key string, value interface{}, ctx *ParsingContext) error {
	for _, node := range nodes {
		if applied, err := node.Apply(key, value, ctx); applied {
			return nil
		} else if err != nil {
			return err
		}
	}
	return nil
}

// parseJSONLDContext implements a super basic JSON-LD @context parsing
// algorithm in order to build a set of nodes which will be able to parse the
// rest of the document.
func parseJSONLDContext(registry *RDFRegistry, input JSONLD) (nodes []RDFNode, err error) {
	i, ok := input[JSON_LD_CONTEXT]
	if !ok {
		err = fmt.Errorf("no @context in input")
		return
	}
	if inArray, ok := i.([]interface{}); ok {
		// @context is an array
		for _, iVal := range inArray {
			if valMap, ok := iVal.(map[string]interface{}); ok {
				// Element is a JSON Object (dictionary)
				for alias, val := range valMap {
					if s, ok := val.(string); ok {
						var n []RDFNode
						n, err = registry.getAliased(alias, s)
						if err != nil {
							return
						}
						nodes = append(nodes, n...)
					} else if aliasedMap, ok := val.(map[string]interface{}); ok {
						var n []RDFNode
						n, err = registry.getAliasedObject(alias, aliasedMap)
						if err != nil {
							return
						}
						nodes = append(nodes, n...)
					} else {
						err = fmt.Errorf("@context value in dict in array is neither a dict nor a string")
						return
					}
				}
			} else if s, ok := iVal.(string); ok {
				// Element is a single value
				var n []RDFNode
				n, err = registry.getFor(s)
				if err != nil {
					return
				}
				nodes = append(nodes, n...)
			} else {
				err = fmt.Errorf("@context value in array is neither a dict nor a string")
				return
			}
		}
	} else if inMap, ok := i.(map[string]interface{}); ok {
		// @context is a JSON object (dictionary)
		for alias, iVal := range inMap {
			if s, ok := iVal.(string); ok {
				var n []RDFNode
				n, err = registry.getAliased(alias, s)
				if err != nil {
					return
				}
				nodes = append(nodes, n...)
			} else if aliasedMap, ok := iVal.(map[string]interface{}); ok {
				var n []RDFNode
				n, err = registry.getAliasedObject(alias, aliasedMap)
				if err != nil {
					return
				}
				nodes = append(nodes, n...)
			} else {
				err = fmt.Errorf("@context value in dict is neither a dict nor a string")
				return
			}
		}
	} else {
		// @context is a single value
		s, ok := i.(string)
		if !ok {
			err = fmt.Errorf("single @context value is not a string")
		}
		return registry.getFor(s)
	}
	return
}

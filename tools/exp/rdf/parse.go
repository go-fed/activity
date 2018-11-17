package rdf

import (
	"fmt"
)

const (
	JSON_LD_CONTEXT = "@context"
)

type JSONLD map[string]interface{}

type RDFGetter interface {
	// GetFor gets RDFNodes based on a context's string.
	GetFor(s string) ([]RDFNode, error)
	// GetAliased gets based on a context string and its alias.
	GetAliased(alias, s string) ([]RDFNode, error)
	// GetAliasedObject gets based on a context object and its alias and
	// definition.
	GetAliasedObject(alias string, object map[string]interface{}) ([]RDFNode, error)
}

type ParseContext interface{}

type RDFNode interface {
	Apply(key string, value interface{}, ctx ParseContext) (bool, error)
}

// ParseJSONLDContext implements a super basic JSON-LD @context parsing
// algorithm in order to build a tree that can parse the rest of the document.
func ParseJSONLDContext(rdfGetter RDFGetter, input JSONLD) (nodes []RDFNode, err error) {
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
						n, err = rdfGetter.GetAliased(alias, s)
						if err != nil {
							return
						}
						nodes = append(nodes, n...)
					} else if aliasedMap, ok := val.(map[string]interface{}); ok {
						var n []RDFNode
						n, err = rdfGetter.GetAliasedObject(alias, aliasedMap)
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
				n, err = rdfGetter.GetFor(s)
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
				n, err = rdfGetter.GetAliased(alias, s)
				if err != nil {
					return
				}
				nodes = append(nodes, n...)
			} else if aliasedMap, ok := iVal.(map[string]interface{}); ok {
				var n []RDFNode
				n, err = rdfGetter.GetAliasedObject(alias, aliasedMap)
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
		return rdfGetter.GetFor(s)
	}
	return
}

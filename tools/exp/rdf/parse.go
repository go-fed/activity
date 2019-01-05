package rdf

import (
	"fmt"
)

const (
	JSON_LD_CONTEXT = "@context"
	JSON_LD_TYPE    = "@type"
	JSON_LD_TYPE_AS = "type"
)

// JSONLD is an alias for the generic map of keys to interfaces, presumably
// parsed from a JSON-encoded context definition file.
type JSONLD map[string]interface{}

// ParsingContext contains the results of the parsing as well as scratch space
// required for RDFNodes to be able to statefully apply changes.
type ParsingContext struct {
	Result  *ParsedVocabulary
	Current interface{}
	Name    string
	Stack   []interface{}
	// Applies the node only for the next level of processing.
	//
	// Do not touch, instead use the accessor methods.
	OnlyApplyThisNodeNextLevel RDFNode
	OnlyApplied                bool
	// Applies the node once, for the rest of the data. This skips the
	// recursive parsing, and the node's Apply is given an empty string
	// for a key.
	//
	// Do not touch, instead use the accessor methods.
	OnlyApplyThisNode RDFNode
}

func (p *ParsingContext) SetOnlyApplyThisNode(n RDFNode) {
	p.OnlyApplyThisNode = n
}

func (p *ParsingContext) ResetOnlyApplyThisNode() {
	p.OnlyApplyThisNode = nil
}

func (p *ParsingContext) SetOnlyApplyThisNodeNextLevel(n RDFNode) {
	p.OnlyApplyThisNodeNextLevel = n
	p.OnlyApplied = false
}

func (p *ParsingContext) GetNextNodes(n []RDFNode) (r []RDFNode, clearFn func()) {
	if p.OnlyApplyThisNodeNextLevel == nil {
		return n, func() {}
	} else if p.OnlyApplied {
		return n, func() {}
	} else {
		p.OnlyApplied = true
		return []RDFNode{p.OnlyApplyThisNodeNextLevel}, func() {
			p.OnlyApplied = false
		}
	}
}

func (p *ParsingContext) ResetOnlyAppliedThisNodeNextLevel() {
	p.OnlyApplyThisNodeNextLevel = nil
	p.OnlyApplied = false
}

func (p *ParsingContext) Push() {
	p.Stack = append([]interface{}{p.Current}, p.Stack...)
	p.Current = nil
}

func (p *ParsingContext) Pop() {
	p.Current = p.Stack[0]
	p.Stack = p.Stack[1:]
	if ng, ok := p.Current.(NameGetter); ok {
		p.Name = ng.GetName()
	}
}

func (p *ParsingContext) IsReset() bool {
	return p.Current == nil &&
		p.Name == ""
}

func (p *ParsingContext) Reset() {
	p.Current = nil
	p.Name = ""
}

type NameSetter interface {
	SetName(string)
}

type NameGetter interface {
	GetName() string
}

type URISetter interface {
	SetURI(string) error
}

type NotesSetter interface {
	SetNotes(string)
}

type ExampleAdder interface {
	AddExample(*VocabularyExample)
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
func ParseVocabulary(name string, registry *RDFRegistry, input JSONLD) (vocabulary *ParsedVocabulary, err error) {
	var nodes []RDFNode
	nodes, err = parseJSONLDContext(registry, input)
	if err != nil {
		return
	}
	vocabulary = &ParsedVocabulary{}
	ctx := &ParsingContext{
		Result: vocabulary,
	}
	// Prepend well-known JSON LD parsing nodes. Order matters, so that the
	// parser can understand things like types so that other nodes do not
	// hijack processing.
	nodes = append(jsonLDNodes(registry), nodes...)
	// Step 1: Parse all core data, excluding:
	//   - Value types
	//   - Referenced types
	//   - VocabularyType's 'Properties' and 'WithoutProperties' fields
	//
	// This is all horrible code but it works, so....
	err = apply(nodes, input, ctx)
	if err != nil {
		return
	}
	// Step 2: Populate value and referenced types.
	err = resolveReferences(registry, ctx)
	if err != nil {
		return
	}
	// Step 3: Populate VocabularyType's 'Properties' and
	// 'WithoutProperties' fields
	err = populatePropertiesOnTypes(ctx)
	// Populate this parsed vocabulary's name
	vocabulary.Vocab.Name = name
	return
}

// populatePropertiesOnTypes populates the 'Properties' and 'WithoutProperties'
// entries on a VocabularyType.
func populatePropertiesOnTypes(ctx *ParsingContext) error {
	for _, p := range ctx.Result.Vocab.Properties {
		if err := populatePropertyOnTypes(p, "", ctx); err != nil {
			return err
		}
	}
	for vName, ref := range ctx.Result.References {
		for _, p := range ref.Properties {
			if err := populatePropertyOnTypes(p, vName, ctx); err != nil {
				return err
			}
		}
	}
	return nil
}

// populatePropertyOnTypes populates the VocabularyType's 'Properties' and
// 'WithoutProperties' fields based on the 'Domain' and 'DoesNotApplyTo'.
func populatePropertyOnTypes(p VocabularyProperty, vocabName string, ctx *ParsingContext) error {
	ref := VocabularyReference{
		Name:  p.Name,
		URI:   p.URI,
		Vocab: vocabName,
	}
	for _, d := range p.Domain {
		if len(d.Vocab) == 0 {
			t, ok := ctx.Result.Vocab.Types[d.Name]
			if !ok {
				return fmt.Errorf("cannot populate property on type %q for desired vocab", d.Name)
			}
			t.Properties = append(t.Properties, ref)
			ctx.Result.Vocab.Types[d.Name] = t
		} else {
			v, ok := ctx.Result.References[d.Vocab]
			if !ok {
				return fmt.Errorf("cannot populate property on type for vocab %q", d.Vocab)
			}
			t, ok := v.Types[d.Name]
			if !ok {
				return fmt.Errorf("cannot populate property on type %q for vocab %q", d.Name, d.Vocab)
			}
			t.Properties = append(t.Properties, ref)
			v.Types[d.Name] = t
		}
	}
	for _, dna := range p.DoesNotApplyTo {
		if len(dna.Vocab) == 0 {
			t, ok := ctx.Result.Vocab.Types[dna.Name]
			if !ok {
				return fmt.Errorf("cannot populate withoutproperty on type %q for desired vocab", dna.Name)
			}
			t.WithoutProperties = append(t.WithoutProperties, ref)
			ctx.Result.Vocab.Types[dna.Name] = t
		} else {
			v, ok := ctx.Result.References[dna.Vocab]
			if !ok {
				return fmt.Errorf("cannot populate withoutproperty on type for vocab %q", dna.Vocab)
			}
			t, ok := v.Types[dna.Name]
			if !ok {
				return fmt.Errorf("cannot populate withoutproperty on type %q for vocab %q", dna.Name, dna.Vocab)
			}
			t.WithoutProperties = append(t.WithoutProperties, ref)
			v.Types[dna.Name] = t
		}
	}
	return nil
}

// resolveReferences ensures that all references mentioned have been
// successfully parsed, and if not attempts to search the ontologies for any
// values, types, and properties that need to be referenced.
//
// Currently, this is the only way that values are added to the
// ParsedVocabulary.
func resolveReferences(registry *RDFRegistry, ctx *ParsingContext) error {
	vocabulary := ctx.Result
	for _, t := range vocabulary.Vocab.Types {
		for _, ref := range t.DisjointWith {
			if err := resolveReference(ref, registry, ctx); err != nil {
				return err
			}
		}
		for _, ref := range t.Extends {
			if err := resolveReference(ref, registry, ctx); err != nil {
				return err
			}
		}
	}
	for _, p := range vocabulary.Vocab.Properties {
		for _, ref := range p.Domain {
			if err := resolveReference(ref, registry, ctx); err != nil {
				return err
			}
		}
		for _, ref := range p.Range {
			if err := resolveReference(ref, registry, ctx); err != nil {
				return err
			}
		}
		for _, ref := range p.DoesNotApplyTo {
			if err := resolveReference(ref, registry, ctx); err != nil {
				return err
			}
		}
		if len(p.SubpropertyOf.Name) > 0 {
			if err := resolveReference(p.SubpropertyOf, registry, ctx); err != nil {
				return err
			}
		}
	}
	return nil
}

// resolveReference will attempt to resolve the reference by either finding it
// in the known References of the vocabulary, or load it from the registry. Will
// fail if a reference is not found.
func resolveReference(reference VocabularyReference, registry *RDFRegistry, ctx *ParsingContext) error {
	name := reference.Name
	vocab := &ctx.Result.Vocab
	if len(reference.Vocab) > 0 {
		name = joinAlias(reference.Vocab, reference.Name)
		url, e := registry.ResolveAlias(reference.Vocab)
		if e != nil {
			return e
		}
		vocab = ctx.Result.GetReference(url)
	}
	if _, ok := vocab.Types[reference.Name]; ok {
		return nil
	} else if _, ok := vocab.Properties[reference.Name]; ok {
		return nil
	} else if _, ok := vocab.Values[reference.Name]; ok {
		return nil
	} else if n, e := registry.getNode(name); e != nil {
		return e
	} else {
		applicable, e := n.Apply("", nil, ctx)
		if !applicable {
			return fmt.Errorf("cannot resolve reference with unapplicable node for %s", reference)
		} else if e != nil {
			return e
		}
		return nil
	}
}

// apply takes a specification input to populate the ParsingContext, based on
// the capabilities of the RDFNodes created from ontologies.
//
// This function will populate all non-value data in the Vocabulary. It does not
// populate the 'Properties' nor the 'WithoutProperties' fields on any
// VocabularyType.
func apply(nodes []RDFNode, input JSONLD, ctx *ParsingContext) error {
	// Hijacked processing: Process the rest of the data in this single
	// node.
	if ctx.OnlyApplyThisNode != nil {
		if applied, err := ctx.OnlyApplyThisNode.Apply("", input, ctx); !applied {
			return fmt.Errorf("applying requested node failed")
		} else {
			return err
		}
	}
	// Special processing: '@type' or 'type' if they are present
	if v, ok := input[JSON_LD_TYPE]; ok {
		if err := doApply(nodes, JSON_LD_TYPE, v, ctx); err != nil {
			return err
		}
	} else if v, ok := input[JSON_LD_TYPE_AS]; ok {
		if err := doApply(nodes, JSON_LD_TYPE_AS, v, ctx); err != nil {
			return err
		}
	}
	// Normal recursive processing
	for k, v := range input {
		// Skip things we have already processed: context and type
		if k == JSON_LD_CONTEXT {
			continue
		} else if k == JSON_LD_TYPE {
			continue
		} else if k == JSON_LD_TYPE_AS {
			continue
		}
		if err := doApply(nodes, k, v, ctx); err != nil {
			return err
		}
	}
	return nil
}

// doApply actually does the application logic for the apply function.
func doApply(nodes []RDFNode,
	k string, v interface{},
	ctx *ParsingContext) error {
	// Hijacked processing: Only use the ParsingContext's node to
	// handle all elements.
	recurNodes := nodes
	enterApplyExitNodes, clearFn := ctx.GetNextNodes(nodes)
	defer clearFn()
	// Normal recursive processing
	if mapValue, ok := v.(map[string]interface{}); ok {
		if err := enterFirstNode(enterApplyExitNodes, k, ctx); err != nil {
			return err
		} else if err = apply(recurNodes, mapValue, ctx); err != nil {
			return err
		} else if err = exitFirstNode(enterApplyExitNodes, k, ctx); err != nil {
			return err
		}
	} else if arrValue, ok := v.([]interface{}); ok {
		for _, val := range arrValue {
			// First, enter for this key
			if err := enterFirstNode(enterApplyExitNodes, k, ctx); err != nil {
				return err
			}
			// Recur or handle the value as necessary.
			if mapValue, ok := val.(map[string]interface{}); ok {
				if err := apply(recurNodes, mapValue, ctx); err != nil {
					return err
				}
			} else if err := applyFirstNode(enterApplyExitNodes, k, val, ctx); err != nil {
				return err
			}
			// Finally, exit for this key
			if err := exitFirstNode(enterApplyExitNodes, k, ctx); err != nil {
				return err
			}
		}
	} else if err := applyFirstNode(enterApplyExitNodes, k, v, ctx); err != nil {
		return err
	}
	return nil
}

// enterFirstNode will Enter the first RDFNode that returns true or an error.
func enterFirstNode(nodes []RDFNode, key string, ctx *ParsingContext) error {
	for _, node := range nodes {
		if applied, err := node.Enter(key, ctx); applied {
			return err
		} else if err != nil {
			return err
		}
	}
	return fmt.Errorf("no RDFNode applicable for entering %q", key)
}

// exitFirstNode will Exit the first RDFNode that returns true or an error.
func exitFirstNode(nodes []RDFNode, key string, ctx *ParsingContext) error {
	for _, node := range nodes {
		if applied, err := node.Exit(key, ctx); applied {
			return err
		} else if err != nil {
			return err
		}
	}
	return fmt.Errorf("no RDFNode applicable for exiting %q", key)
}

// applyFirstNode will Apply the first RDFNode that returns true or an error.
func applyFirstNode(nodes []RDFNode, key string, value interface{}, ctx *ParsingContext) error {
	for _, node := range nodes {
		if applied, err := node.Apply(key, value, ctx); applied {
			return err
		} else if err != nil {
			return err
		}
	}
	return fmt.Errorf("no RDFNode applicable for applying %q with value %v", key, value)
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

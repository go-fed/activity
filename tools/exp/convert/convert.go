package convert

import (
	"fmt"
	"github.com/cjslep/activity/tools/exp/props"
	"github.com/cjslep/activity/tools/exp/rdf"
	"github.com/cjslep/activity/tools/exp/types"
	"strings"
)

type vocabulary struct {
	Kinds   map[string]*props.Kind
	FProps  map[string]*props.FunctionalPropertyGenerator
	NFProps map[string]*props.NonFunctionalPropertyGenerator
	Types   map[string]*types.TypeGenerator
}

func newVocabulary() vocabulary {
	return vocabulary{
		Kinds:   make(map[string]*props.Kind, 0),
		FProps:  make(map[string]*props.FunctionalPropertyGenerator, 0),
		NFProps: make(map[string]*props.NonFunctionalPropertyGenerator, 0),
		Types:   make(map[string]*types.TypeGenerator, 0),
	}
}

type PropertyPackagePolicy int

const (
	FlatUnderRoot PropertyPackagePolicy = iota
	IndividualUnderRoot
)

type Converter struct {
	Registry              *rdf.RDFRegistry
	VocabularyRoot        string
	PropertyPackagePolicy PropertyPackagePolicy
	PropertyPackageRoot   string
}

func (c Converter) convert(p *rdf.ParsedVocabulary) (v vocabulary, e error) {
	v = newVocabulary()
	for k, val := range p.Vocab.Values {
		v.Kinds[k] = c.convertValue(val)
	}
	for k, prop := range p.Vocab.Properties {
		if prop.Functional {
			v.FProps[k], e = c.convertFunctionalProperty(prop, v.Kinds, p.Vocab, p.References)
		} else {
			v.NFProps[k], e = c.convertNonFunctionalProperty(prop, v.Kinds, p.Vocab, p.References)
		}
		if e != nil {
			return
		}
	}
	// TODO: Types in dependency order.
	return
}

func (c Converter) convertFunctionalProperty(p rdf.VocabularyProperty,
	kinds map[string]*props.Kind,
	v rdf.Vocabulary,
	refs map[string]*rdf.Vocabulary) (fp *props.FunctionalPropertyGenerator, e error) {
	var k []props.Kind
	k, e = c.propertyKinds(p, kinds, v, refs)
	if e != nil {
		return
	}
	var pkg string
	pkg, e = c.propertyPackageName(p)
	if e != nil {
		return
	}
	fp = props.NewFunctionalPropertyGenerator(
		pkg,
		c.toIdentifier(p),
		k,
		p.NaturalLanguageMap)
	return
}

func (c Converter) convertNonFunctionalProperty(p rdf.VocabularyProperty,
	kinds map[string]*props.Kind,
	v rdf.Vocabulary,
	refs map[string]*rdf.Vocabulary) (nfp *props.NonFunctionalPropertyGenerator, e error) {
	var k []props.Kind
	k, e = c.propertyKinds(p, kinds, v, refs)
	if e != nil {
		return
	}
	var pkg string
	pkg, e = c.propertyPackageName(p)
	if e != nil {
		return
	}
	nfp = props.NewNonFunctionalPropertyGenerator(
		pkg,
		c.toIdentifier(p),
		k,
		p.NaturalLanguageMap)
	return
}

func (c Converter) convertValue(v rdf.VocabularyValue) (k *props.Kind) {
	k = &props.Kind{
		Name:          c.toIdentifier(v),
		ConcreteKind:  v.DefinitionType,
		Nilable:       c.isNilable(v.DefinitionType),
		SerializeFn:   v.SerializeFn,
		DeserializeFn: v.DeserializeFn,
		LessFn:        v.LessFn,
	}
	return
}

func (c Converter) convertTypeToKind(v rdf.VocabularyType) (k *props.Kind) {
	k = &props.Kind{
		Name:         c.toIdentifier(v),
		ConcreteKind: strings.Title(v.Name),
		Nilable:      true,
		// TODO
		SerializeFn:   types.SerializeFn,
		DeserializeFn: types.DeserializeFn,
		LessFn:        types.LessFn,
	}
	return
}

func (c Converter) propertyKinds(v rdf.VocabularyProperty,
	kinds map[string]*props.Kind,
	vocab rdf.Vocabulary,
	refs map[string]*rdf.Vocabulary) (k []props.Kind, e error) {
	for _, r := range v.Range {
		if len(r.Vocab) == 0 {
			if kind, ok := kinds[r.Name]; !ok {
				// It is a Type of the vocabulary
				if t, ok := vocab.Types[r.Name]; !ok {
					e = fmt.Errorf("cannot find own kind with name %q", r.Name)
					return
				} else {
					k = append(k, *c.convertTypeToKind(t))
				}
			} else {
				// It is a Value of the vocabulary
				k = append(k, *kind)
			}
		} else {
			var url string
			url, e = c.Registry.ResolveAlias(r.Vocab)
			if e != nil {
				return
			}
			refVocab, ok := refs[url]
			if !ok {
				e = fmt.Errorf("references do not contain %s", url)
				return
			}
			if val, ok := refVocab.Values[r.Name]; !ok {
				// It is a Type of the vocabulary instead
				if t, ok := refVocab.Types[r.Name]; !ok {
					e = fmt.Errorf("cannot find kind with name %q in %s", r.Name, url)
					return
				} else {
					k = append(k, *c.convertTypeToKind(t))
				}
			} else {
				// It is a Value of the vocabulary
				k = append(k, *c.convertValue(val))
			}
		}
	}
	return
}

func (c Converter) propertyPackageName(v rdf.VocabularyProperty) (pkg string, e error) {
	switch c.PropertyPackagePolicy {
	case FlatUnderRoot:
		pkg = c.PropertyPackageRoot
	case IndividualUnderRoot:
		pkg = v.Name
	default:
		e = fmt.Errorf("unrecognized PropertyPackagePolicy: %v", c.PropertyPackagePolicy)
	}
	return
}

// TODO: Use this?
func (c Converter) propertyPackageFile(v rdf.VocabularyProperty) (pkg string, e error) {
	switch c.PropertyPackagePolicy {
	case FlatUnderRoot:
		pkg = fmt.Sprintf("%s/%s", c.VocabularyRoot, c.PropertyPackageRoot)
	case IndividualUnderRoot:
		pkg = fmt.Sprintf("%s/%s/%s", c.VocabularyRoot, c.PropertyPackageRoot, v.Name)
	default:
		e = fmt.Errorf("unrecognized PropertyPackagePolicy: %v", c.PropertyPackagePolicy)
	}
	return
}

func (c Converter) toIdentifier(n rdf.NameGetter) props.Identifier {
	return props.Identifier{
		LowerName: n.GetName(),
		CamelName: strings.Title(n.GetName()),
	}
}

func (c Converter) isNilable(goType string) bool {
	return goType[0] == '*'
}

package convert

import (
	"fmt"
	"github.com/cjslep/activity/tools/exp/props"
	"github.com/cjslep/activity/tools/exp/rdf"
	"github.com/cjslep/activity/tools/exp/types"
	"github.com/dave/jennifer/jen"
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
	PropertyFlatUnderRoot PropertyPackagePolicy = iota
	PropertyIndividualUnderRoot
	PropertyFlatUnderVocabularyRoot
)

type TypePackagePolicy int

const (
	TypeFlatUnderRoot TypePackagePolicy = iota
	TypeIndividualUnderRoot
	TypeFlatUnderVocabularyRoot
)

type Converter struct {
	Registry              *rdf.RDFRegistry
	VocabularyRoot        string
	PropertyPackagePolicy PropertyPackagePolicy
	PropertyPackageRoot   string
	TypePackageRoot       string
	TypePackagePolicy     TypePackagePolicy
}

func (c Converter) Convert(p *rdf.ParsedVocabulary) (f []*jen.File, e error) {
	var v vocabulary
	v, e = c.convertVocabulary(p)
	if e != nil {
		return
	}
	// TODO
	return
}

func (c Converter) convertVocabulary(p *rdf.ParsedVocabulary) (v vocabulary, e error) {
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
	// Instead of building a dependency tree, naively keep iterating through
	// 'allTypes' until it is empty (good) or we get stuck (return error).
	allTypes := make([]rdf.VocabularyType, 0, len(p.Vocab.Types))
	for _, t := range p.Vocab.Types {
		allTypes = append(allTypes, t)
	}
	for {
		if len(allTypes) == 0 {
			break
		}
		stuck := true
		for _, t := range allTypes {
			var tg *types.TypeGenerator
			tg, e = c.convertType(t, p.Vocab, v.FProps, v.NFProps, v.Types)
			if e != nil {
				return
			}
			v.Types[t.Name] = tg
		}
		if stuck {
			e = fmt.Errorf("converting types got stuck in dependency cycle")
			return
		}
	}
	return
}

func (c Converter) convertType(t rdf.VocabularyType,
	v rdf.Vocabulary,
	existingFProps map[string]*props.FunctionalPropertyGenerator,
	existingNFProps map[string]*props.NonFunctionalPropertyGenerator,
	existingTypes map[string]*types.TypeGenerator) (tg *types.TypeGenerator, e error) {
	// Determine the types package name
	var pkg string
	pkg, e = c.typePackageName(t)
	if e != nil {
		return
	}
	// Determine the properties for this type
	var p []types.Property
	for k, prop := range v.Properties {
		for _, ref := range prop.Domain {
			if len(ref.Vocab) != 0 {
				e = fmt.Errorf("unhandled use case: property domain outside its vocabulary")
				return
			} else if ref.Name == t.Name {
				if prop.Functional {
					p = append(p, existingFProps[prop.Name])
				} else {
					p = append(p, existingNFProps[prop.Name])
				}
				break
			}
		}
	}
	// Determine what this type extends
	var ext []*types.TypeGenerator
	for _, ex := range t.Extends {
		if len(ex.Vocab) != 0 {
			// TODO: This should be fixed
			e = fmt.Errorf("unhandled use case: type extends another type outside its vocabulary")
			return
		} else {
			ext = append(ext, existingTypes[ex.Name])
		}
	}
	tg, e = types.NewTypeGenerator(
		pkg,
		strings.Title(t.Name), // Must be same in convertTypeToKind
		t.Notes,
		p,
		ext,
		nil)
	if e != nil {
		return
	}
	// Apply disjoint if both sides are available.
	for _, disj := range t.DisjointWith {
		if len(disj.Vocab) != 0 {
			// TODO: This should be fixed
			e = fmt.Errorf("unhandled use case: type is disjoint with another type outside its vocabulary")
			return
		} else if disjointType, ok := existingTypes[disj.Name]; ok {
			disjointType.AddDisjoint(tg)
			tg.AddDisjoint(disjointType)
		}
	}
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
		ConcreteKind: strings.Title(v.Name), // Must be same in convertType
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

func (c Converter) typePackageName(v rdf.VocabularyType) (pkg string, e error) {
	switch c.TypePackagePolicy {
	case TypeFlatUnderRoot:
		pkg = c.TypePackageRoot
	case TypeIndividualUnderRoot:
		pkg = v.Name
	case TypeFlatUnderVocabularyRoot:
		pkg = c.VocabularyRoot
	default:
		e = fmt.Errorf("unrecognized TypePackagePolicy: %v", c.TypePackagePolicy)
	}
	return
}

func (c Converter) propertyPackageName(v rdf.VocabularyProperty) (pkg string, e error) {
	switch c.PropertyPackagePolicy {
	case PropertyFlatUnderRoot:
		pkg = c.PropertyPackageRoot
	case PropertyIndividualUnderRoot:
		pkg = v.Name
	case PropertyFlatUnderVocabularyRoot:
		pkg = c.VocabularyRoot
	default:
		e = fmt.Errorf("unrecognized PropertyPackagePolicy: %v", c.PropertyPackagePolicy)
	}
	return
}

// TODO: Use this?
func (c Converter) typePackageFile(v rdf.VocabularyType) (pkg string, e error) {
	switch c.TypePackagePolicy {
	case TypeFlatUnderRoot:
		pkg = fmt.Sprintf("%s/%s", c.VocabularyRoot, c.TypePackageRoot)
	case TypeIndividualUnderRoot:
		pkg = fmt.Sprintf("%s/%s/%s", c.VocabularyRoot, c.TypePackageRoot, v.Name)
	case TypeFlatUnderVocabularyRoot:
		pkg = c.VocabularyRoot
	default:
		e = fmt.Errorf("unrecognized TypePackagePolicy: %v", c.TypePackagePolicy)
	}
	return
}

// TODO: Use this?
func (c Converter) propertyPackageFile(v rdf.VocabularyProperty) (pkg string, e error) {
	switch c.PropertyPackagePolicy {
	case PropertyFlatUnderRoot:
		pkg = fmt.Sprintf("%s/%s", c.VocabularyRoot, c.PropertyPackageRoot)
	case PropertyIndividualUnderRoot:
		pkg = fmt.Sprintf("%s/%s/%s", c.VocabularyRoot, c.PropertyPackageRoot, v.Name)
	case PropertyFlatUnderVocabularyRoot:
		pkg = c.VocabularyRoot
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

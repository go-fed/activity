package convert

import (
	"fmt"
	"github.com/cjslep/activity/tools/exp/props"
	"github.com/cjslep/activity/tools/exp/rdf"
	"github.com/dave/jennifer/jen"
	"strings"
)

type File struct {
	F         *jen.File
	FileName  string
	Directory string
}

type vocabulary struct {
	Kinds   map[string]*props.Kind
	FProps  map[string]*props.FunctionalPropertyGenerator
	NFProps map[string]*props.NonFunctionalPropertyGenerator
	Types   map[string]*props.TypeGenerator
}

func newVocabulary() vocabulary {
	return vocabulary{
		Kinds:   make(map[string]*props.Kind, 0),
		FProps:  make(map[string]*props.FunctionalPropertyGenerator, 0),
		NFProps: make(map[string]*props.NonFunctionalPropertyGenerator, 0),
		Types:   make(map[string]*props.TypeGenerator, 0),
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
	TypePackagePolicy     TypePackagePolicy
	TypePackageRoot       string
}

func (c Converter) Convert(p *rdf.ParsedVocabulary) (f []*File, e error) {
	var v vocabulary
	v, e = c.convertVocabulary(p)
	if e != nil {
		return
	}
	f, e = c.convertToFiles(v)
	return
}

func (c Converter) convertToFiles(v vocabulary) (f []*File, e error) {
	for _, _ = range v.Kinds {
		// TODO: Implement
	}
	for _, i := range v.FProps {
		var pkg string
		pkg, e = c.propertyPackageFile(i)
		if e != nil {
			return
		}
		var dir string
		dir, e = c.propertyPackageDirectory(i)
		if e != nil {
			return
		}
		file := jen.NewFilePath(pkg)
		file.Add(i.Definition().Definition())
		f = append(f, &File{
			F:         file,
			FileName:  fmt.Sprintf("gen_%s.go", i.PropertyName()),
			Directory: dir,
		})
	}
	for _, i := range v.NFProps {
		var pkg string
		pkg, e = c.propertyPackageFile(i)
		if e != nil {
			return
		}
		var dir string
		dir, e = c.propertyPackageDirectory(i)
		if e != nil {
			return
		}
		file := jen.NewFilePath(pkg)
		s, t := i.Definitions()
		file.Add(s.Definition()).Line().Add(t.Definition())
		f = append(f, &File{
			F:         file,
			FileName:  fmt.Sprintf("gen_%s.go", i.PropertyName()),
			Directory: dir,
		})
	}
	for _, i := range v.Types {
		var pkg string
		pkg, e = c.typePackageFile(i)
		if e != nil {
			return
		}
		var dir string
		dir, e = c.typePackageDirectory(i)
		if e != nil {
			return
		}
		file := jen.NewFilePath(pkg)
		file.Add(i.Definition().Definition())
		f = append(f, &File{
			F:         file,
			FileName:  fmt.Sprintf("gen_%s.go", i.TypeName()),
			Directory: dir,
		})
	}
	return
}

// convertVocabulary works in a two-pass system: first converting all known
// properties, and then the types.
//
// Due to the fact that properties rely on the Kind abstraction, and both
// properties and types can be Kinds, this introduces tight coupling between
// the two so that callbacks can fill in missing links in data that isn't known
// beforehand (ex: how to serialize, deserialize, and compare types).
//
// This feels very hacky and could be decoupled using standard design patterns,
// but since there is no need, it isn't addressed now.
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
		for i, t := range allTypes {
			if allExtendsAreIn(t, v.Types) {
				var tg *props.TypeGenerator
				tg, e = c.convertType(t, p.Vocab, v.FProps, v.NFProps, v.Types)
				if e != nil {
					return
				}
				v.Types[t.Name] = tg
				stuck = false
				// Delete the one we just did.
				allTypes[i] = allTypes[len(allTypes)-1]
				allTypes = allTypes[:len(allTypes)-1]
				break
			}
		}
		if stuck {
			e = fmt.Errorf("converting props got stuck in dependency cycle")
			return
		}
	}
	return
}

func (c Converter) convertType(t rdf.VocabularyType,
	v rdf.Vocabulary,
	existingFProps map[string]*props.FunctionalPropertyGenerator,
	existingNFProps map[string]*props.NonFunctionalPropertyGenerator,
	existingTypes map[string]*props.TypeGenerator) (tg *props.TypeGenerator, e error) {
	// Determine the props package name
	var pkg string
	pkg, e = c.typePackageName(t)
	if e != nil {
		return
	}
	// Determine the properties for this type
	var p []props.Property
	for _, prop := range t.Properties {
		if len(prop.Vocab) != 0 {
			e = fmt.Errorf("unhandled use case: property domain outside its vocabulary")
			return
		} else {
			var property props.Property
			var ok bool
			property, ok = existingFProps[prop.Name]
			if !ok {
				property, ok = existingNFProps[prop.Name]
				if !ok {
					e = fmt.Errorf("cannot find property with name: %s", prop.Name)
					return
				}
			}
			p = append(p, property)
		}
	}
	// Determine WithoutProperties for this type
	var wop []props.Property
	for _, prop := range t.WithoutProperties {
		if len(prop.Vocab) != 0 {
			e = fmt.Errorf("unhandled use case: withoutproperty domain outside its vocabulary")
			return
		} else {
			var property props.Property
			var ok bool
			property, ok = existingFProps[prop.Name]
			if !ok {
				property, ok = existingNFProps[prop.Name]
				if !ok {
					e = fmt.Errorf("cannot find property with name: %s", prop.Name)
					return
				}
			}
			wop = append(wop, property)
		}
	}
	// Determine what this type extends
	var ext []*props.TypeGenerator
	for _, ex := range t.Extends {
		if len(ex.Vocab) != 0 {
			// TODO: This should be fixed to handle references
			e = fmt.Errorf("unhandled use case: type extends another type outside its vocabulary")
			return
		} else {
			ext = append(ext, existingTypes[ex.Name])
		}
	}
	tg, e = props.NewTypeGenerator(
		pkg,
		c.convertTypeToName(t),
		t.Notes,
		p,
		wop,
		ext,
		nil)
	if e != nil {
		return
	}
	// Apply disjoint if both sides are available because the TypeGenerator
	// does not know the entire vocabulary, so cannot do this lookup and
	// create this connection for us.
	//
	// TODO: Pass in the disjoint and have the TypeGenerator complete the
	// doubly-linked connection for us.
	for _, disj := range t.DisjointWith {
		if len(disj.Vocab) != 0 {
			// TODO: This should be fixed to handle references
			e = fmt.Errorf("unhandled use case: type is disjoint with another type outside its vocabulary")
			return
		} else if disjointType, ok := existingTypes[disj.Name]; ok {
			disjointType.AddDisjoint(tg)
			tg.AddDisjoint(disjointType)
		}
	}
	// Apply the type's KindSerializationFuncs to the property because there
	// is no way for the TypeGenerator to know all properties who have a
	// range of this type.
	//
	// TODO: Pass in these properties to the TypeGenerator constructor so it
	// can build these double-links properly. Note this would also need to
	// apply to referenced properties, possibly.
	for _, prop := range existingFProps {
		for _, kind := range prop.Kinds {
			if kind.Name.LowerName == tg.TypeName() {
				ser, deser, less := tg.KindSerializationFuncs()
				if e = prop.SetKindFns(tg.TypeName(), ser, deser, less); e != nil {
					return
				}
			}
		}
	}
	for _, prop := range existingNFProps {
		for _, kind := range prop.Kinds {
			if kind.Name.LowerName == tg.TypeName() {
				ser, deser, less := tg.KindSerializationFuncs()
				if e = prop.SetKindFns(tg.TypeName(), ser, deser, less); e != nil {
					return
				}
			}
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

func (c Converter) convertTypeToKind(v rdf.VocabularyType) (k *props.Kind, e error) {
	k = &props.Kind{
		Name:         c.toIdentifier(v),
		ConcreteKind: c.convertTypeToConcreteKind(v),
		Nilable:      true,
		// Instead of populating:
		//   - SerializeFn
		//   - DeserializeFn
		//   - LessFn
		//
		// The TypeGenerator is responsible for calling setKindFns on
		// the properties, to property wire a Property's Kind back to
		// the Type's implementation.
	}
	return
}

func (c Converter) convertTypeToName(v rdf.VocabularyType) string {
	return strings.Title(v.Name)
}

func (c Converter) convertTypeToConcreteKind(v rdf.VocabularyType) string {
	return "*" + c.convertTypeToName(v)
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
					var kt *props.Kind
					kt, e = c.convertTypeToKind(t)
					if e != nil {
						return
					}
					k = append(k, *kt)
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
					var kt *props.Kind
					kt, e = c.convertTypeToKind(t)
					if e != nil {
						return
					}
					k = append(k, *kt)
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

func (c Converter) typePackageDirectory(v *props.TypeGenerator) (dir string, e error) {
	switch c.TypePackagePolicy {
	case TypeFlatUnderRoot:
		dir = fmt.Sprintf("%s/%s/", c.VocabularyRoot, c.TypePackageRoot)
	case TypeIndividualUnderRoot:
		dir = fmt.Sprintf("%s/%s/%s/", c.VocabularyRoot, c.TypePackageRoot, v.TypeName())
	case TypeFlatUnderVocabularyRoot:
		dir = c.VocabularyRoot + "/"
	default:
		e = fmt.Errorf("unrecognized TypePackagePolicy: %v", c.TypePackagePolicy)
	}
	return
}

func (c Converter) typePackageFile(v *props.TypeGenerator) (pkg string, e error) {
	switch c.TypePackagePolicy {
	case TypeFlatUnderRoot:
		pkg = fmt.Sprintf("%s/%s", c.VocabularyRoot, c.TypePackageRoot)
	case TypeIndividualUnderRoot:
		pkg = fmt.Sprintf("%s/%s/%s", c.VocabularyRoot, c.TypePackageRoot, v.TypeName())
	case TypeFlatUnderVocabularyRoot:
		pkg = c.VocabularyRoot
	default:
		e = fmt.Errorf("unrecognized TypePackagePolicy: %v", c.TypePackagePolicy)
	}
	return
}

type propertyNamer interface {
	PropertyName() string
}

var (
	_ propertyNamer = &props.FunctionalPropertyGenerator{}
	_ propertyNamer = &props.NonFunctionalPropertyGenerator{}
)

func (c Converter) propertyPackageDirectory(v propertyNamer) (dir string, e error) {
	switch c.PropertyPackagePolicy {
	case PropertyFlatUnderRoot:
		dir = fmt.Sprintf("%s/%s/", c.VocabularyRoot, c.PropertyPackageRoot)
	case PropertyIndividualUnderRoot:
		dir = fmt.Sprintf("%s/%s/%s/", c.VocabularyRoot, c.PropertyPackageRoot, v.PropertyName())
	case PropertyFlatUnderVocabularyRoot:
		dir = c.VocabularyRoot + "/"
	default:
		e = fmt.Errorf("unrecognized PropertyPackagePolicy: %v", c.PropertyPackagePolicy)
	}
	return
}

func (c Converter) propertyPackageFile(v propertyNamer) (pkg string, e error) {
	switch c.PropertyPackagePolicy {
	case PropertyFlatUnderRoot:
		pkg = fmt.Sprintf("%s/%s", c.VocabularyRoot, c.PropertyPackageRoot)
	case PropertyIndividualUnderRoot:
		pkg = fmt.Sprintf("%s/%s/%s", c.VocabularyRoot, c.PropertyPackageRoot, v.PropertyName())
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

func allExtendsAreIn(t rdf.VocabularyType, v map[string]*props.TypeGenerator) bool {
	for _, e := range t.Extends {
		if len(e.Vocab) != 0 {
			// TODO: This should be fixed to handle references
			return false
		} else if _, ok := v[e.Name]; !ok {
			return false
		}
	}
	return true
}

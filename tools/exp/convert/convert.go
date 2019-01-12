package convert

import (
	"fmt"
	"github.com/cjslep/activity/tools/exp/codegen"
	"github.com/cjslep/activity/tools/exp/gen"
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
	Values     map[string]*gen.Kind
	FProps     map[string]*gen.FunctionalPropertyGenerator
	NFProps    map[string]*gen.NonFunctionalPropertyGenerator
	Types      map[string]*gen.TypeGenerator
	Manager    *gen.ManagerGenerator
	References map[string]*vocabulary
}

func newVocabulary() vocabulary {
	return vocabulary{
		Values:     make(map[string]*gen.Kind, 0),
		FProps:     make(map[string]*gen.FunctionalPropertyGenerator, 0),
		NFProps:    make(map[string]*gen.NonFunctionalPropertyGenerator, 0),
		Types:      make(map[string]*gen.TypeGenerator, 0),
		References: make(map[string]*vocabulary, 0),
	}
}

func (v vocabulary) typeArray() []*gen.TypeGenerator {
	tg := make([]*gen.TypeGenerator, 0, len(v.Types))
	for _, t := range v.Types {
		tg = append(tg, t)
	}
	return tg
}

func (v vocabulary) propArray() []*gen.PropertyGenerator {
	fp := make([]*gen.PropertyGenerator, 0, len(v.FProps)+len(v.NFProps))
	for _, f := range v.FProps {
		fp = append(fp, &f.PropertyGenerator)
	}
	for _, f := range v.NFProps {
		fp = append(fp, &f.PropertyGenerator)
	}
	return fp
}

func (v vocabulary) funcPropArray() []*gen.FunctionalPropertyGenerator {
	fp := make([]*gen.FunctionalPropertyGenerator, 0, len(v.FProps))
	for _, f := range v.FProps {
		fp = append(fp, f)
	}
	return fp
}

func (v vocabulary) nonFuncPropArray() []*gen.NonFunctionalPropertyGenerator {
	nfp := make([]*gen.NonFunctionalPropertyGenerator, 0, len(v.NFProps))
	for _, nf := range v.NFProps {
		nfp = append(nfp, nf)
	}
	return nfp
}

type PackagePolicy int

const (
	FlatUnderRoot PackagePolicy = iota
	IndividualUnderRoot
)

type Converter struct {
	Registry       *rdf.RDFRegistry
	GenRoot        *gen.PackageManager
	VocabularyName string
	ValueRoot      *gen.PackageManager
	PackagePolicy  PackagePolicy
}

func (c Converter) Convert(p *rdf.ParsedVocabulary) (f []*File, e error) {
	var v vocabulary
	v, e = c.convertVocabulary(p)
	if e != nil {
		return
	}
	for k, refVocab := range p.References {
		// Create a copy, but with the Reference moved to Vocab.
		refP := p
		refP.Vocab = *refVocab
		delete(refP.References, k)

		var refV vocabulary
		refV, e = c.convertVocabulary(refP)
		if e != nil {
			return
		}
		v.References[k] = &refV
	}
	f, e = c.convertToFiles(v)
	return
}

func (c Converter) convertToFiles(v vocabulary) (f []*File, e error) {
	// Values -- include all referenced values too.
	for _, v := range v.Values {
		pkg := c.valuePackage(v)
		f = append(f, convertValue(pkg, v))
	}
	for _, ref := range v.References {
		for _, v := range ref.Values {
			pkg := c.valuePackage(v)
			f = append(f, convertValue(pkg, v))
		}
	}
	// Functional Properties
	for _, i := range v.FProps {
		var pm *gen.PackageManager
		pm, e = c.propertyPackageManager(i)
		if e != nil {
			return
		}
		// Implementation
		priv := pm.PrivatePackage()
		file := jen.NewFilePath(priv.Path())
		file.Add(i.Definition().Definition())
		f = append(f, &File{
			F:         file,
			FileName:  fmt.Sprintf("gen_property_%s.go", i.PropertyName()),
			Directory: priv.WriteDir(),
		})
		// Interface
		pub := pm.PublicPackage()
		file = jen.NewFilePath(pub.Path())
		file.Add(i.InterfaceDefinition(pm.PublicPackage()).Definition())
		f = append(f, &File{
			F:         file,
			FileName:  fmt.Sprintf("gen_property_%s_interface.go", i.PropertyName()),
			Directory: pub.WriteDir(),
		})
	}
	// Non-Functional Properties
	for _, i := range v.NFProps {
		var pm *gen.PackageManager
		pm, e = c.propertyPackageManager(i)
		if e != nil {
			return
		}
		// Implementation
		priv := pm.PrivatePackage()
		file := jen.NewFilePath(priv.Path())
		s, t := i.Definitions()
		file.Add(s.Definition()).Line().Add(t.Definition())
		f = append(f, &File{
			F:         file,
			FileName:  fmt.Sprintf("gen_property_%s.go", i.PropertyName()),
			Directory: priv.WriteDir(),
		})
		// Interface
		pub := pm.PublicPackage()
		file = jen.NewFilePath(pub.Path())
		for _, intf := range i.InterfaceDefinitions(pm.PublicPackage()) {
			file.Add(intf.Definition()).Line()
		}
		f = append(f, &File{
			F:         file,
			FileName:  fmt.Sprintf("gen_property_%s_interface.go", i.PropertyName()),
			Directory: pub.WriteDir(),
		})
	}
	// Types
	for _, i := range v.Types {
		var pm *gen.PackageManager
		pm, e = c.typePackageManager(i)
		if e != nil {
			return
		}
		// Implementation
		priv := pm.PrivatePackage()
		file := jen.NewFilePath(priv.Path())
		file.Add(i.Definition().Definition())
		f = append(f, &File{
			F:         file,
			FileName:  fmt.Sprintf("gen_type_%s.go", i.TypeName()),
			Directory: priv.WriteDir(),
		})
		// Interface
		pub := pm.PublicPackage()
		file = jen.NewFilePath(pub.Path())
		file.Add(i.InterfaceDefinition(pm.PublicPackage()).Definition())
		f = append(f, &File{
			F:         file,
			FileName:  fmt.Sprintf("gen_type_%s_interface.go", i.TypeName()),
			Directory: pub.WriteDir(),
		})
	}
	pkgFiles, err := c.packageFiles(v)
	if err != nil {
		e = err
		return
	}
	f = append(f, pkgFiles...)
	// Manager
	pub := c.GenRoot.PublicPackage()
	file := jen.NewFilePath(pub.Path())
	file.Add(v.Manager.Definition().Definition())
	f = append(f, &File{
		F:         file,
		FileName:  "gen_manager.go",
		Directory: pub.WriteDir(),
	})
	var files []*File
	files, e = c.rootFiles(pub, c.VocabularyName, v)
	if e != nil {
		return
	}
	f = append(f, files...)
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
		v.Values[k] = c.convertValue(val)
	}
	for k, prop := range p.Vocab.Properties {
		if prop.Functional {
			v.FProps[k], e = c.convertFunctionalProperty(prop, v.Values, p.Vocab, p.References)
		} else {
			v.NFProps[k], e = c.convertNonFunctionalProperty(prop, v.Values, p.Vocab, p.References)
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
				var tg *gen.TypeGenerator
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
			e = fmt.Errorf("converting gen got stuck in dependency cycle")
			return
		}
	}
	v.Manager, e = gen.NewManagerGenerator(
		c.GenRoot.PublicPackage(),
		v.typeArray(),
		v.funcPropArray(),
		v.nonFuncPropArray())
	return
}

func (c Converter) convertType(t rdf.VocabularyType,
	v rdf.Vocabulary,
	existingFProps map[string]*gen.FunctionalPropertyGenerator,
	existingNFProps map[string]*gen.NonFunctionalPropertyGenerator,
	existingTypes map[string]*gen.TypeGenerator) (tg *gen.TypeGenerator, e error) {
	// Determine the gen package name
	var pm *gen.PackageManager
	pm, e = c.typePackageManager(t)
	if e != nil {
		return
	}
	// Determine the properties for this type
	var p []gen.Property
	for _, prop := range t.Properties {
		if len(prop.Vocab) != 0 {
			e = fmt.Errorf("unhandled use case: property domain outside its vocabulary")
			return
		} else {
			var property gen.Property
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
	var wop []gen.Property
	for _, prop := range t.WithoutProperties {
		if len(prop.Vocab) != 0 {
			e = fmt.Errorf("unhandled use case: withoutproperty domain outside its vocabulary")
			return
		} else {
			var property gen.Property
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
	var ext []*gen.TypeGenerator
	for _, ex := range t.Extends {
		if len(ex.Vocab) != 0 {
			// TODO: This should be fixed to handle references
			e = fmt.Errorf("unhandled use case: type extends another type outside its vocabulary")
			return
		} else {
			ext = append(ext, existingTypes[ex.Name])
		}
	}
	// Apply disjoint if both sides are available because the TypeGenerator
	// does not know the entire vocabulary, so cannot do this lookup and
	// create this connection for us.
	var disjoint []*gen.TypeGenerator
	for _, disj := range t.DisjointWith {
		if len(disj.Vocab) != 0 {
			// TODO: This should be fixed to handle references
			e = fmt.Errorf("unhandled use case: type is disjoint with another type outside its vocabulary")
			return
		} else if disjointType, ok := existingTypes[disj.Name]; ok {
			disjoint = append(disjoint, disjointType)
		}
	}
	// Pass in properties whose range is this type so it can build
	// references properly.
	//
	// Note that the Kinds container on properties contains both types and
	// values.
	//
	// TODO: Enable this for referenced properties.
	name := c.convertTypeToName(t)
	var rangeProps []gen.Property
	for _, prop := range existingFProps {
		for _, kind := range prop.GetKinds() {
			if kind.Name.LowerName == name {
				rangeProps = append(rangeProps, prop)
			}
		}
	}
	for _, prop := range existingNFProps {
		for _, kind := range prop.GetKinds() {
			if kind.Name.LowerName == name {
				rangeProps = append(rangeProps, prop)
			}
		}
	}
	tg, e = gen.NewTypeGenerator(
		v.GetName(),
		pm,
		name,
		t.Notes,
		p,
		wop,
		rangeProps,
		ext,
		disjoint)
	return
}

func (c Converter) convertFunctionalProperty(p rdf.VocabularyProperty,
	kinds map[string]*gen.Kind,
	v rdf.Vocabulary,
	refs map[string]*rdf.Vocabulary) (fp *gen.FunctionalPropertyGenerator, e error) {
	var k []gen.Kind
	k, e = c.propertyKinds(p, kinds, v, refs)
	if e != nil {
		return
	}
	var pm *gen.PackageManager
	pm, e = c.propertyPackageManager(p)
	if e != nil {
		return
	}
	fp = gen.NewFunctionalPropertyGenerator(
		v.GetName(),
		pm,
		c.toIdentifier(p),
		p.Notes,
		k,
		p.NaturalLanguageMap)
	return
}

func (c Converter) convertNonFunctionalProperty(p rdf.VocabularyProperty,
	kinds map[string]*gen.Kind,
	v rdf.Vocabulary,
	refs map[string]*rdf.Vocabulary) (nfp *gen.NonFunctionalPropertyGenerator, e error) {
	var k []gen.Kind
	k, e = c.propertyKinds(p, kinds, v, refs)
	if e != nil {
		return
	}
	var pm *gen.PackageManager
	pm, e = c.propertyPackageManager(p)
	if e != nil {
		return
	}
	nfp = gen.NewNonFunctionalPropertyGenerator(
		v.GetName(),
		pm,
		c.toIdentifier(p),
		p.Notes,
		k,
		p.NaturalLanguageMap)
	return
}

func (c Converter) convertValue(v rdf.VocabularyValue) (k *gen.Kind) {
	s := v.SerializeFn.CloneToPackage(c.vocabValuePackage(v).Path())
	d := v.DeserializeFn.CloneToPackage(c.vocabValuePackage(v).Path())
	l := v.LessFn.CloneToPackage(c.vocabValuePackage(v).Path())
	k = &gen.Kind{
		Name:           c.toIdentifier(v),
		ConcreteKind:   v.DefinitionType,
		Nilable:        v.IsNilable,
		IsURI:          v.IsURI,
		SerializeFn:    s.QualifiedName(),
		DeserializeFn:  d.QualifiedName(),
		LessFn:         l.QualifiedName(),
		SerializeDef:   s,
		DeserializeDef: d,
		LessDef:        l,
	}
	return
}

func (c Converter) convertTypeToKind(v rdf.VocabularyType) (k *gen.Kind, e error) {
	k = &gen.Kind{
		Name:    c.toIdentifier(v),
		Nilable: true,
		IsURI:   false,
		// Instead of populating:
		//   - ConcreteKind
		//   - SerializeFn
		//   - DeserializeFn
		//   - LessFn
		//
		// The TypeGenerator is responsible for calling SetKindFns on
		// the properties, to property wire a Property's Kind back to
		// the Type's implementation.
	}
	return
}

func (c Converter) convertTypeToName(v rdf.VocabularyType) string {
	return strings.Title(v.Name)
}

func (c Converter) propertyKinds(v rdf.VocabularyProperty,
	kinds map[string]*gen.Kind,
	vocab rdf.Vocabulary,
	refs map[string]*rdf.Vocabulary) (k []gen.Kind, e error) {
	for _, r := range v.Range {
		if len(r.Vocab) == 0 {
			if kind, ok := kinds[r.Name]; !ok {
				// It is a Type of the vocabulary
				if t, ok := vocab.Types[r.Name]; !ok {
					e = fmt.Errorf("cannot find own kind with name %q", r.Name)
					return
				} else {
					var kt *gen.Kind
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
					var kt *gen.Kind
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

func (c Converter) valuePackage(v *gen.Kind) gen.Package {
	return c.ValueRoot.Sub(v.Name.LowerName).PublicPackage()
}

func (c Converter) vocabValuePackage(v rdf.VocabularyValue) gen.Package {
	return c.ValueRoot.Sub(c.toIdentifier(v).LowerName).PublicPackage()
}

func (c Converter) typePackageManager(v typeNamer) (pkg *gen.PackageManager, e error) {
	return c.packageManager("type_" + v.TypeName())
}

func (c Converter) propertyPackageManager(v propertyNamer) (pkg *gen.PackageManager, e error) {
	return c.packageManager("property_" + v.PropertyName())
}

func (c Converter) packageManager(s string) (pkg *gen.PackageManager, e error) {
	s = strings.ToLower(s)
	switch c.PackagePolicy {
	case FlatUnderRoot:
		pkg = c.GenRoot.Sub(strings.ToLower(c.VocabularyName))
	case IndividualUnderRoot:
		pkg = c.GenRoot.Sub(strings.ToLower(c.VocabularyName)).SubPrivate(s)
	default:
		e = fmt.Errorf("unrecognized PackagePolicy: %v", c.PackagePolicy)
	}
	return
}

func (c Converter) rootFiles(pkg gen.Package, vocabName string, v vocabulary) (f []*File, e error) {
	pg := gen.NewPackageGenerator()
	ctors, ext, disj, extBy, globalVar, initFn := pg.RootDefinitions(vocabName, v.Manager, v.typeArray(), v.propArray())
	initFile := jen.NewFilePath(pkg.Path())
	initFile.Add(globalVar).Line().Add(initFn.Definition()).Line()
	f = append(f, &File{
		F:         initFile,
		FileName:  "gen_init.go",
		Directory: pkg.WriteDir(),
	})
	f = append(f, funcsToFile(pkg, ctors, fmt.Sprintf("gen_pkg_%s_constructors.go", vocabName)))
	f = append(f, funcsToFile(pkg, ext, fmt.Sprintf("gen_pkg_%s_extends.go", vocabName)))
	f = append(f, funcsToFile(pkg, disj, fmt.Sprintf("gen_pkg_%s_disjoint.go", vocabName)))
	f = append(f, funcsToFile(pkg, extBy, fmt.Sprintf("gen_pkg_%s_extendedby.go", vocabName)))
	return
}

func (c Converter) packageFiles(v vocabulary) (f []*File, e error) {
	switch c.PackagePolicy {
	case FlatUnderRoot:
		// Only need one for all types.
		pg := gen.NewPackageGenerator()
		pubI := pg.PublicDefinitions(v.typeArray())
		// Public
		pub := v.typeArray()[0].PublicPackage()
		file := jen.NewFilePath(pub.Path())
		file.Add(pubI.Definition())
		f = append(f, &File{
			F:         file,
			FileName:  "gen_pkg.go",
			Directory: pub.WriteDir(),
		})
		// Private
		s, i, fn := pg.PrivateDefinitions(v.typeArray(), v.propArray())
		priv := v.typeArray()[0].PrivatePackage()
		file = jen.NewFilePath(priv.Path())
		file.Add(
			s,
		).Line().Add(
			i.Definition(),
		).Line().Add(
			fn.Definition(),
		).Line()
		f = append(f, &File{
			F:         file,
			FileName:  "gen_pkg.go",
			Directory: priv.WriteDir(),
		})
	case IndividualUnderRoot:
		for _, tg := range v.Types {
			var file []*File
			file, e = c.typePackageFiles(tg)
			if e != nil {
				return
			}
			f = append(f, file...)
		}
		for _, pg := range v.FProps {
			var file []*File
			file, e = c.propertyPackageFiles(&pg.PropertyGenerator)
			if e != nil {
				return
			}
			f = append(f, file...)
		}
		for _, pg := range v.NFProps {
			var file []*File
			file, e = c.propertyPackageFiles(&pg.PropertyGenerator)
			if e != nil {
				return
			}
			f = append(f, file...)
		}
	default:
		e = fmt.Errorf("unrecognized PackagePolicy: %v", c.PackagePolicy)
	}
	return
}

func (c Converter) typePackageFiles(tg *gen.TypeGenerator) (f []*File, e error) {
	// Only need one for all types.
	tpg := gen.NewTypePackageGenerator()
	pubI := tpg.PublicDefinitions([]*gen.TypeGenerator{tg})
	// Public
	pub := tg.PublicPackage()
	file := jen.NewFilePath(pub.Path())
	file.Add(pubI.Definition())
	f = append(f, &File{
		F:         file,
		FileName:  "gen_pkg.go",
		Directory: pub.WriteDir(),
	})
	// Private
	s, i, fn := tpg.PrivateDefinitions([]*gen.TypeGenerator{tg})
	priv := tg.PrivatePackage()
	file = jen.NewFilePath(priv.Path())
	file.Add(
		s,
	).Line().Add(
		i.Definition(),
	).Line().Add(
		fn.Definition(),
	).Line()
	f = append(f, &File{
		F:         file,
		FileName:  "gen_pkg.go",
		Directory: priv.WriteDir(),
	})
	return
}

func (c Converter) propertyPackageFiles(pg *gen.PropertyGenerator) (f []*File, e error) {
	// Only need one for all types.
	ppg := gen.NewPropertyPackageGenerator()
	// Private
	s, i, fn := ppg.PrivateDefinitions([]*gen.PropertyGenerator{pg})
	priv := pg.GetPrivatePackage()
	file := jen.NewFilePath(priv.Path())
	file.Add(
		s,
	).Line().Add(
		i.Definition(),
	).Line().Add(
		fn.Definition(),
	).Line()
	f = append(f, &File{
		F:         file,
		FileName:  "gen_pkg.go",
		Directory: priv.WriteDir(),
	})
	return
}

type typeNamer interface {
	TypeName() string
}

var (
	_ typeNamer = &gen.TypeGenerator{}
	_ typeNamer = &rdf.VocabularyType{}
)

type propertyNamer interface {
	PropertyName() string
}

var (
	_ propertyNamer = &gen.FunctionalPropertyGenerator{}
	_ propertyNamer = &gen.NonFunctionalPropertyGenerator{}
	_ propertyNamer = &rdf.VocabularyProperty{}
)

func (c Converter) toIdentifier(n rdf.NameGetter) gen.Identifier {
	return gen.Identifier{
		LowerName: n.GetName(),
		CamelName: strings.Title(n.GetName()),
	}
}

func allExtendsAreIn(t rdf.VocabularyType, v map[string]*gen.TypeGenerator) bool {
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

func convertValue(pkg gen.Package, v *gen.Kind) *File {
	file := jen.NewFilePath(pkg.Path())
	file.Add(
		v.SerializeDef.Definition(),
	).Line().Add(
		v.DeserializeDef.Definition(),
	).Line().Add(
		v.LessDef.Definition())
	return &File{
		F:         file,
		FileName:  fmt.Sprintf("gen_%s.go", v.Name.LowerName),
		Directory: pkg.WriteDir(),
	}
}

func funcsToFile(pkg gen.Package, fns []*codegen.Function, filename string) *File {
	file := jen.NewFilePath(pkg.Path())
	for _, fn := range fns {
		file.Add(fn.Definition()).Line()
	}
	return &File{
		F:         file,
		FileName:  filename,
		Directory: pkg.WriteDir(),
	}
}

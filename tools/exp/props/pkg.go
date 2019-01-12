package props

import (
	"fmt"
	"github.com/cjslep/activity/tools/exp/codegen"
	"github.com/dave/jennifer/jen"
	"strings"
)

// PackageManager manages the path and names of a package consisting of a public
// and a private portion.
type PackageManager struct {
	prefix  string
	root    string
	public  string
	private string
}

func NewPackageManager(prefix, root string) *PackageManager {
	return &PackageManager{
		prefix:  prefix,
		root:    root,
		public:  "",
		private: "impl",
	}
}

// PublicPackage returns the public package.
func (p *PackageManager) PublicPackage() Package {
	return p.toPackage(p.public, true)
}

// PrivatePackage returns the private package.
func (p *PackageManager) PrivatePackage() Package {
	return p.toPackage(p.private, false)
}

// Sub creates a PackageManager clone that manages a subdirectory.
func (p *PackageManager) Sub(name string) *PackageManager {
	return &PackageManager{
		prefix:  p.prefix,
		root:    fmt.Sprintf("%s/%s", p.root, name),
		public:  p.public,
		private: p.private,
	}
}

// SubPrivate creates a PackageManager clone where the private package is one
// subdirectory further.
func (p *PackageManager) SubPrivate(name string) *PackageManager {
	return &PackageManager{
		prefix:  p.prefix,
		root:    p.root,
		public:  p.public,
		private: fmt.Sprintf("%s/%s", p.private, name),
	}
}

func (p *PackageManager) toPackage(suffix string, public bool) Package {
	path := p.root
	if len(suffix) > 0 {
		path = strings.Join([]string{p.root, suffix}, "/")
	}
	s := strings.Split(path, "/")
	name := s[len(s)-1]
	return Package{
		prefix:   p.prefix,
		path:     path,
		name:     name,
		isPublic: public,
		parent:   p,
	}
}

type Package struct {
	prefix   string
	path     string
	name     string
	isPublic bool
	parent   *PackageManager
}

func (p Package) Path() string {
	return p.prefix + "/" + p.path
}

func (p Package) WriteDir() string {
	return p.path
}

func (p Package) Name() string {
	return p.name
}

func (p Package) IsPublic() bool {
	return p.isPublic
}

func (p Package) Parent() *PackageManager {
	return p.parent
}

const (
	managerInterfaceName   = "privateManager"
	setManagerFunctionName = "SetManager"
)

// TypePackageGenerator manages generating one-time files needed for types.
type TypePackageGenerator struct{}

// NewTypePackageGenerator creates a new TypePackageGenerator.
func NewTypePackageGenerator() *TypePackageGenerator {
	return &TypePackageGenerator{}
}

// RootDefinitions creates functions needed at the root level of the package declarations.
func (t *TypePackageGenerator) RootDefinitions(vocabName string, m *ManagerGenerator, tgs []*TypeGenerator) (ctors, ext, disj, extBy []*codegen.Function, globalManager *jen.Statement, init *codegen.Function) {
	return rootDefinitions(vocabName, m, tgs)
}

// PublicDefinitions creates the public-facing code generated definitions needed
// once per package.
//
// Precondition: The passed-in generators are the complete set of type
// generators within a package.
func (t *TypePackageGenerator) PublicDefinitions(tgs []*TypeGenerator) (typeI *codegen.Interface) {
	return publicTypeDefinitions(tgs)
}

// PrivateDefinitions creates the private code generated definitions needed once
// per package.
//
// Precondition: The passed-in generators are the complete set of type
// generators within a package.
func (t *TypePackageGenerator) PrivateDefinitions(tgs []*TypeGenerator) (mgrVar *jen.Statement, mgrI *codegen.Interface, setMgrFn *codegen.Function) {
	return privateManagerHookDefinitions(tgs, nil)
}

// PropertyPackageGenerator manages generating one-time files needed for
// properties.
type PropertyPackageGenerator struct{}

// NewPropertyPackageGenerator creates a new PropertyPackageGenerator.
func NewPropertyPackageGenerator() *PropertyPackageGenerator {
	return &PropertyPackageGenerator{}
}

// PrivateDefinitions creates the private code generated definitions needed once
// per package.
//
// Precondition: The passed-in generators are the complete set of type
// generators within a package.
func (p *PropertyPackageGenerator) PrivateDefinitions(pgs []*PropertyGenerator) (*jen.Statement, *codegen.Interface, *codegen.Function) {
	return privateManagerHookDefinitions(nil, pgs)
}

// PackageGenerator maanges generating one-time files needed for both type and
// property implementations.
type PackageGenerator struct{}

// NewPackageGenerator creates a new PackageGenerator.
func NewPackageGenerator() *PackageGenerator {
	return &PackageGenerator{}
}

// RootDefinitions creates functions needed at the root level of the package declarations.
func (t *PackageGenerator) RootDefinitions(vocabName string, m *ManagerGenerator, tgs []*TypeGenerator) (ctors, ext, disj, extBy []*codegen.Function, globalManager *jen.Statement, init *codegen.Function) {
	return rootDefinitions(vocabName, m, tgs)
}

// PublicDefinitions creates the public-facing code generated definitions needed
// once per package.
//
// Precondition: The passed-in generators are the complete set of type
// generators within a package.
func (t *PackageGenerator) PublicDefinitions(tgs []*TypeGenerator) *codegen.Interface {
	return publicTypeDefinitions(tgs)
}

// PrivateDefinitions creates the private code generated definitions needed once
// per package.
//
// Precondition: The passed-in generators are the complete set of type
// generators within a package.
func (t *PackageGenerator) PrivateDefinitions(tgs []*TypeGenerator, pgs []*PropertyGenerator) (*jen.Statement, *codegen.Interface, *codegen.Function) {
	return privateManagerHookDefinitions(tgs, pgs)
}

// privateManagerHookDefinitions creates common code needed by types and
// properties to properly hook in the manager at initialization time.
func privateManagerHookDefinitions(tgs []*TypeGenerator, pgs []*PropertyGenerator) (mgrVar *jen.Statement, mgrI *codegen.Interface, setMgrFn *codegen.Function) {
	fnsMap := make(map[string]codegen.FunctionSignature)
	for _, tg := range tgs {
		for _, m := range tg.getAllManagerMethods() {
			v := m.ToFunctionSignature()
			fnsMap[v.Name] = v
		}
	}
	for _, pg := range pgs {
		for _, m := range pg.getAllManagerMethods() {
			v := m.ToFunctionSignature()
			fnsMap[v.Name] = v
		}
	}
	var fns []codegen.FunctionSignature
	for _, v := range fnsMap {
		fns = append(fns, v)
	}
	var path string
	if tgs != nil {
		path = tgs[0].PrivatePackage().Path()
	} else {
		path = pgs[0].GetPrivatePackage().Path()
	}
	return jen.Var().Id(managerInitName()).Id(managerInterfaceName),
		codegen.NewInterface(path,
			managerInterfaceName,
			fns,
			fmt.Sprintf("%s abstracts the code-generated manager that provides access to concrete implementations.", managerInterfaceName)),
		codegen.NewCommentedFunction(path,
			setManagerFunctionName,
			[]jen.Code{
				jen.Id("m").Id(managerInterfaceName),
			},
			/*ret=*/ nil,
			[]jen.Code{
				jen.Id(managerInitName()).Op("=").Id("m"),
			},
			fmt.Sprintf("%s sets the manager package-global variable. For internal use only, do not use as part of Application behavior. Must be called at golang init time.", setManagerFunctionName))
}

// publicTypeDefinitions creates common types needed by types for their public
// package.
func publicTypeDefinitions(tgs []*TypeGenerator) (typeI *codegen.Interface) {
	return TypeInterface(tgs[0].PublicPackage())
}

// rootDefinitions creates common functions needed at the root level of the
// package declarations.
func rootDefinitions(vocabName string, m *ManagerGenerator, tgs []*TypeGenerator) (ctors, ext, disj, extBy []*codegen.Function, globalManager *jen.Statement, init *codegen.Function) {
	// Type constructors
	for _, tg := range tgs {
		ctors = append(ctors, codegen.NewCommentedFunction(
			m.pkg.Path(),
			fmt.Sprintf("New%s%s", tg.PublicPackage().Name(), tg.TypeName()),
			/*params=*/ nil,
			[]jen.Code{jen.Qual(tg.PublicPackage().Path(), tg.InterfaceName())},
			[]jen.Code{
				jen.Return(
					tg.constructorFn().Call(),
				),
			},
			fmt.Sprintf("New%s%s creates a new %s", tg.PublicPackage().Name(), tg.TypeName(), tg.InterfaceName())))
	}
	// Extends
	for _, tg := range tgs {
		f, _ := tg.extendsDefinition()
		name := fmt.Sprintf("%s%s", vocabName, f.Name())
		ext = append(ext, codegen.NewCommentedFunction(
			m.pkg.Path(),
			name,
			[]jen.Code{jen.Id("other").Qual(tg.PublicPackage().Path(), typeInterfaceName)},
			[]jen.Code{jen.Bool()},
			[]jen.Code{
				jen.Return(
					f.Call(jen.Id("other")),
				),
			},
			fmt.Sprintf("%s returns true if %s extends from the other's type.", name, tg.TypeName())))
	}
	// DisjointWith
	for _, tg := range tgs {
		f := tg.disjointWithDefinition()
		name := fmt.Sprintf("%s%s", vocabName, f.Name())
		disj = append(disj, codegen.NewCommentedFunction(
			m.pkg.Path(),
			name,
			[]jen.Code{jen.Id("other").Qual(tg.PublicPackage().Path(), typeInterfaceName)},
			[]jen.Code{jen.Bool()},
			[]jen.Code{
				jen.Return(
					f.Call(jen.Id("other")),
				),
			},
			fmt.Sprintf("%s returns true if %s is disjoint with the other's type.", name, tg.TypeName())))
	}
	// ExtendedBy
	for _, tg := range tgs {
		f := tg.extendedByDefinition()
		name := fmt.Sprintf("%s%s", vocabName, f.Name())
		extBy = append(extBy, codegen.NewCommentedFunction(
			m.pkg.Path(),
			name,
			[]jen.Code{jen.Id("other").Qual(tg.PublicPackage().Path(), typeInterfaceName)},
			[]jen.Code{jen.Bool()},
			[]jen.Code{
				jen.Return(
					f.Call(jen.Id("other")),
				),
			},
			fmt.Sprintf("%s returns true if the other's type extends from %s.", name, tg.TypeName())))
	}
	// init
	globalManager = jen.Var().Id(managerInitName()).Op("*").Qual(m.pkg.Path(), managerName)
	callInitsMap := make(map[string]jen.Code, len(tgs))
	for _, tg := range tgs {
		callInitsMap[tg.PrivatePackage().Path()] = jen.Qual(tg.PrivatePackage().Path(), setManagerFunctionName).Call(
			jen.Qual(m.pkg.Path(), managerInitName()),
		)
	}
	callInits := make([]jen.Code, 0, len(callInitsMap))
	for _, c := range callInitsMap {
		callInits = append(callInits, c)
	}
	init = codegen.NewCommentedFunction(
		m.pkg.Path(),
		"init",
		/*params=*/ nil,
		/*ret=*/ nil,
		append([]jen.Code{
			jen.Qual(m.pkg.Path(), managerInitName()).Op("=").Op("&").Qual(m.pkg.Path(), managerName).Values(),
		}, callInits...),
		fmt.Sprintf("init handles the 'magic' of creating a %s and dependency-injecting it into every other code-generated package. This gives the implementations access to create any type needed to deserialize, without relying on the other specific concrete implementations. In order to replace a go-fed created type with your own, be sure to have the manager call your own implementation's deserialize functions instead of the built-in type. Finally, each implementation views the %s as an interface with only a subset of funcitons available. This means this %s implements the union of those interfaces.", managerName, managerName, managerName))
	return
}

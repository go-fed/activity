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
func (t *TypePackageGenerator) RootDefinitions(tgs []*TypeGenerator) (ctors []*codegen.Function) {
	// Type constructors
	for _, tg := range tgs {
		ctors = append(ctors, codegen.NewCommentedFunction(
			tg.PublicPackage().Path(),
			fmt.Sprintf("New%s%s", tg.PublicPackage().Name(), tg.TypeName()),
			/*params=*/ nil,
			[]jen.Code{jen.Qual(tg.PublicPackage().Path(), tg.InterfaceName())},
			[]jen.Code{
				jen.Return(
					jen.Op("&").Qual(tg.PrivatePackage().Path(), tg.TypeName()).Values(
						jen.Dict{
							jen.Id(unknownMember): jen.Make(jen.Map(jen.String()).Interface(), jen.Lit(0)),
						},
					),
				),
			},
			fmt.Sprintf("New%s%s creates a new %s", tg.PublicPackage().Name(), tg.TypeName(), tg.InterfaceName())))
	}
	return
}

// PublicDefinitions creates the public-facing code generated definitions needed
// once per package.
//
// Precondition: The passed-in generators are the complete set of type
// generators within a package.
func (t *TypePackageGenerator) PublicDefinitions(tgs []*TypeGenerator) (typeI *codegen.Interface) {
	return TypeInterface(tgs[0].PublicPackage())
}

// PrivateDefinitions creates the private code generated definitions needed once
// per package.
//
// Precondition: The passed-in generators are the complete set of type
// generators within a package.
func (t *TypePackageGenerator) PrivateDefinitions(tgs []*TypeGenerator) (mgrVar *jen.Statement, mgrI *codegen.Interface, setMgrFn *codegen.Function) {
	fnsMap := make(map[string]codegen.FunctionSignature)
	for _, tg := range tgs {
		for _, m := range tg.getAllManagerMethods() {
			v := m.ToFunctionSignature()
			fnsMap[v.Name] = v
		}
	}
	var fns []codegen.FunctionSignature
	for _, v := range fnsMap {
		fns = append(fns, v)
	}
	return jen.Var().Id(managerInitName()).Id(managerInterfaceName),
		codegen.NewInterface(tgs[0].PrivatePackage().Path(),
			managerInterfaceName,
			fns,
			fmt.Sprintf("%s abstracts the code-generated manager that provides access to concrete implementations.", managerInterfaceName)),
		codegen.NewCommentedFunction(tgs[0].PrivatePackage().Path(),
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
	fnsMap := make(map[string]codegen.FunctionSignature)
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
	return jen.Var().Id(managerInitName()).Id(managerInterfaceName),
		codegen.NewInterface(pgs[0].GetPrivatePackage().Path(),
			managerInterfaceName,
			fns,
			fmt.Sprintf("%s abstracts the code-generated manager that provides access to concrete implementations.", managerInterfaceName)),
		codegen.NewCommentedFunction(pgs[0].GetPrivatePackage().Path(),
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

// PackageGenerator maanges generating one-time files needed for both type and
// property implementations.
type PackageGenerator struct{}

// NewPackageGenerator creates a new PackageGenerator.
func NewPackageGenerator() *PackageGenerator {
	return &PackageGenerator{}
}

// RootDefinitions creates functions needed at the root level of the package declarations.
func (t *PackageGenerator) RootDefinitions(tgs []*TypeGenerator) (ctors []*codegen.Function) {
	// Type constructors
	for _, tg := range tgs {
		ctors = append(ctors, codegen.NewCommentedFunction(
			tg.PublicPackage().Path(),
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
	return
}

// PublicDefinitions creates the public-facing code generated definitions needed
// once per package.
//
// Precondition: The passed-in generators are the complete set of type
// generators within a package.
func (t *PackageGenerator) PublicDefinitions(tgs []*TypeGenerator) *codegen.Interface {
	return TypeInterface(tgs[0].PublicPackage())
}

// PrivateDefinitions creates the private code generated definitions needed once
// per package.
//
// Precondition: The passed-in generators are the complete set of type
// generators within a package.
func (t *PackageGenerator) PrivateDefinitions(tgs []*TypeGenerator, pgs []*PropertyGenerator) (*jen.Statement, *codegen.Interface, *codegen.Function) {
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
	return jen.Var().Id(managerInitName()).Id(managerInterfaceName),
		codegen.NewInterface(tgs[0].PrivatePackage().Path(),
			managerInterfaceName,
			fns,
			fmt.Sprintf("%s abstracts the code-generated manager that provides access to concrete implementations.", managerInterfaceName)),
		codegen.NewCommentedFunction(tgs[0].PrivatePackage().Path(),
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

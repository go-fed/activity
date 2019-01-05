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

// PublicDefinitions creates the public-facing code generated definitions needed
// once per package.
//
// Precondition: The passed-in generators are the complete set of type
// generators within a package.
func (t *TypePackageGenerator) PublicDefinitions(tgs []*TypeGenerator) *codegen.Interface {
	return TypeInterface(tgs[0].PublicPackage())
}

// PrivateDefinitions creates the private code generated definitions needed once
// per package.
//
// Precondition: The passed-in generators are the complete set of type
// generators within a package.
func (t *TypePackageGenerator) PrivateDefinitions(tgs []*TypeGenerator) (*jen.Statement, *codegen.Interface, *codegen.Function) {
	var fns []codegen.FunctionSignature
	for _, tg := range tgs {
		for _, m := range tg.getAllManagerMethods() {
			fns = append(fns, m.ToFunctionSignature())
		}
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
			jen.Commentf("%s sets the manager package-global variable. For internal use only, do not use as part of Application behavior. Must be called at golang init time.", setManagerFunctionName))
}

// PropertyPackageGenerator manages generating one-time files needed for
// properties.
type PropertyPackageGenerator struct{}

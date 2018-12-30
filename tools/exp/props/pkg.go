package props

import (
	"fmt"
	"strings"
)

// PackageManager manages the path and names of a package consisting of a public
// and a private portion.
type PackageManager struct {
	root    string
	public  string
	private string
}

func NewPackageManager(root string) *PackageManager {
	return &PackageManager{
		root:    root,
		public:  "",
		private: "internal",
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
		path:     path,
		name:     name,
		isPublic: public,
		parent:   p,
	}
}

type Package struct {
	path     string
	name     string
	isPublic bool
	parent   *PackageManager
}

func (p Package) Path() string {
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

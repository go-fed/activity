package props

import (
	"fmt"
	"github.com/cjslep/activity/tools/exp/codegen"
	"github.com/dave/jennifer/jen"
	"strings"
	"sync"
)

const (
	managerName        = "ASManager"
	managerInitVarName = "mgr"
)

func managerInitName() string {
	return managerInitVarName
}

// Generates the ActivityStreamManager that handles the static and/or dynamic
// loading of ActivityStream Core, Extended, and any extension types.
//
// This also provides interfaces to break the recursive/cyclic dependencies
// between properties and types. The previous version of this tool did not
// attempt to solve this problem, and instead just created one big and bloated
// library in order to avoid having to break the dependence. This version of
// the tool instead will generate interfaces for all of the required types.
//
// This means that developers will only ever need to interact with these
// interfaces, and could switch out using this implementation for another one of
// their own choosing.
//
// Also note that the manager links against all the implementations to generate
// a comprehensive registry. So while individual properties and types are able
// to be compiled separately, this generated output will link against all of
// these libraries.
//
// TODO: Improve the code generation to examine specific Golang code to
// determine which types to actually generate, and prune the unneeded types.
// This would cut down on the bloat on a per-program basis.
type ManagerGenerator struct {
	pm  PackageManager
	tg  []*TypeGenerator
	fp  []*FunctionalPropertyGenerator
	nfp []*NonFunctionalPropertyGenerator
	// Constructed at creation time. These rely on pointer stability,
	// which should happen as none of these generators are treated as
	// values.
	tgManagedMethods  map[*TypeGenerator]*managedMethods
	fpManagedMethods  map[*FunctionalPropertyGenerator]*managedMethods
	nfpManagedMethods map[*NonFunctionalPropertyGenerator]*managedMethods
	// Cached during manager code generation.
	cacheOnce    sync.Once
	cachedStruct *codegen.Struct
}

// managedMethods caches the specific methods and interfaces mapped to specific
// properties and types.
type managedMethods struct {
	deserializor *codegen.Method
	// TODO: Delete these?
	publicDeserializor *codegen.Method
	// Interface for a manager
	mIface *codegen.Interface
	// Interface(s) for the property/type
	ifaces []*codegen.Interface
}

// NewManagerGenerator creates a new manager system.
//
// This generator should be constructed in the third pass, after types and
// property generators are all constructed.
func NewManagerGenerator(pm PackageManager,
	tg []*TypeGenerator,
	fp []*FunctionalPropertyGenerator,
	nfp []*NonFunctionalPropertyGenerator) (*ManagerGenerator, error) {
	mg := &ManagerGenerator{
		pm:                pm,
		tg:                tg,
		fp:                fp,
		nfp:               nfp,
		tgManagedMethods:  make(map[*TypeGenerator]*managedMethods, len(tg)),
		fpManagedMethods:  make(map[*FunctionalPropertyGenerator]*managedMethods, len(fp)),
		nfpManagedMethods: make(map[*NonFunctionalPropertyGenerator]*managedMethods, len(nfp)),
	}
	// Pass 1: Get all deserializor-like methods created. Further passes may
	// rely on already having this data available in the manager.
	for _, t := range tg {
		mg.tgManagedMethods[t] = &managedMethods{
			// TODO: Figure out how to use this instead of the Kind abstraction
			deserializor:       mg.createPrivateDeserializationMethodForType(t),
			publicDeserializor: nil,
			mIface:             nil,
		}
	}
	for _, p := range fp {
		mg.fpManagedMethods[p] = &managedMethods{
			deserializor:       mg.createPrivateDeserializationMethodForFuncProperty(p),
			publicDeserializor: nil,
			mIface:             nil,
		}
	}
	for _, p := range nfp {
		mg.nfpManagedMethods[p] = &managedMethods{
			deserializor:       mg.createPrivateDeserializationMethodForNonFuncProperty(p),
			publicDeserializor: nil,
			mIface:             nil,
		}
	}
	// Pass 2: Inform the type of this ManagerGenerator so that it can keep
	// all of its bookkeeping straight.
	for _, t := range tg {
		if e := t.apply(mg); e != nil {
			return nil, e
		}
	}
	// Pass 3: Populate interfaces of the types and properties, which relies
	// on the first pass's data population.
	for _, t := range tg {
		publicPkg := t.PublicPackage()
		mg.tgManagedMethods[t].ifaces = []*codegen.Interface{t.toInterface(publicPkg)}
	}
	// TODO: Move these back to pass 1
	for _, p := range fp {
		publicPkg := p.GetPublicPackage()
		mg.fpManagedMethods[p].ifaces = []*codegen.Interface{p.toInterface(publicPkg)}
	}
	for _, p := range nfp {
		publicPkg := p.GetPublicPackage()
		mg.nfpManagedMethods[p].ifaces = p.toInterfaces(publicPkg)
	}
	return mg, nil
}

func (m *ManagerGenerator) publicManager() *codegen.Struct {
	// TODO
	return nil
}

func (m *ManagerGenerator) privatePackage() Package {
	return m.pm.PrivatePackage()
}

func (m *ManagerGenerator) getPrivateDeserializationMethodForType(t *TypeGenerator) *codegen.Method {
	return m.tgManagedMethods[t].deserializor
}

func (m *ManagerGenerator) getPrivateDeserializationMethodForProperty(p Property) *codegen.Method {
	switch v := p.(type) {
	case *FunctionalPropertyGenerator:
		return m.fpManagedMethods[v].deserializor
	case *NonFunctionalPropertyGenerator:
		return m.nfpManagedMethods[v].deserializor
	default:
		panic("unknown property type")
	}
}

// PrivateManager creates a manager implementation that works with the concrete
// types required by the other PropertyGenerators and TypeGenerators for
// serializing and deserializing.
//
// Applications should NOT use this private manager as it will force them to
// rely on the concrete type implementations. The public version uses interfaces
// which isolates Application code from this specific go-fed implementation. If
// another alternative to go-fed were to be created, it could target those
// interfaces and be a drop-in replacement for everyone's applications.
//
// It is necessary to acheive isolation without cyclic dependencies: types and
// properties can each belong in their own package (if desired) to minimize
// binary bloat. This is theoretical until the below TODO is tackled because
// this concrete implementation relies on linking against all of the types and
// properties to behave correctly. There is no simple way around this. But I am
// preparing for the better solution here.
//
// TODO: Analyze a given program, and only generate the types and properties in
// use by the go program.
func (m *ManagerGenerator) PrivateManager() *codegen.Struct {
	var methods []*codegen.Method
	// TODO: Interface versions of these methods should also be present for
	// these specific types.
	for _, tg := range m.tgManagedMethods {
		methods = append(methods, tg.deserializor)
	}
	for _, fp := range m.fpManagedMethods {
		methods = append(methods, fp.deserializor)
	}
	for _, nfp := range m.nfpManagedMethods {
		methods = append(methods, nfp.deserializor)
	}
	s := codegen.NewStruct(
		jen.Commentf(fmt.Sprintf("%s privately manages concrete types and deserializations for internal use by generated code. Application code should use the public version instead, which uses interfaces to abstract away the generated code and allow apps to not entirely rely on go-fed should they choose not to.", managerName)),
		managerName,
		methods,
		/*functions=*/ nil,
		/*members=*/ nil)
	return s
}

// createPrivateDeserializationMethodForType creates a new method for the
// private manager.
//
// TODO: Unify these three methods behind some kind of interface.
func (m *ManagerGenerator) createPrivateDeserializationMethodForType(tg *TypeGenerator) *codegen.Method {
	dn := tg.deserializationFnName()
	pkg := m.pm.PrivatePackage()
	// TODO: Better naming scheme from package name
	name := fmt.Sprintf("%s%s%s", dn, tg.PrivatePackage().Name(), strings.Title(pkg.Name()))
	return codegen.NewCommentedValueMethod(
		pkg.Path(),
		name,
		managerName,
		/*param=*/ nil,
		[]jen.Code{
			jen.Func().Params(
				jen.Map(jen.String()).Interface(),
			).Params(
				// TODO: Qualify this.
				jen.Op("*").Id(tg.TypeName()),
				jen.Error(),
			),
		},
		[]jen.Code{
			jen.Return(
				jen.Qual(tg.PrivatePackage().Path(), dn),
			),
		},
		jen.Commentf("%s returns the deserialization method for the %q type in package %q", name, tg.TypeName(), tg.PrivatePackage().Name()))
}

// createPrivateDeserializationMethodForFuncProperty creates a new method for the
// private manager.
//
// TODO: Unify these three methods behind some kind of interface.
func (m *ManagerGenerator) createPrivateDeserializationMethodForFuncProperty(fp *FunctionalPropertyGenerator) *codegen.Method {
	dn := fp.DeserializeFnName()
	pkg := m.pm.PrivatePackage()
	// TODO: Better naming scheme from package name
	name := fmt.Sprintf("%s%s%s", dn, fp.GetPrivatePackage().Name(), strings.Title(pkg.Name()))
	return codegen.NewCommentedValueMethod(
		pkg.Path(),
		name,
		managerName,
		/*param=*/ nil,
		[]jen.Code{
			jen.Func().Params(
				jen.Map(jen.String()).Interface(),
			).Params(
				// TODO: Qualify this.
				jen.Op("*").Id(fp.StructName()),
				jen.Error(),
			),
		},
		[]jen.Code{
			jen.Return(
				jen.Qual(fp.GetPrivatePackage().Path(), dn),
			),
		},
		jen.Commentf("%s returns the deserialization method for the %q functional property in package %q", name, fp.StructName(), fp.GetPrivatePackage().Name()))
}

// createPrivateDeserializationMethodForNonFuncProperty creates a new method for the
// private manager.
//
// TODO: Unify these three methods behind some kind of interface.
func (m *ManagerGenerator) createPrivateDeserializationMethodForNonFuncProperty(nfp *NonFunctionalPropertyGenerator) *codegen.Method {
	dn := nfp.DeserializeFnName()
	pkg := m.pm.PrivatePackage()
	// TODO: Better naming scheme from package name
	name := fmt.Sprintf("%s%s%s", dn, nfp.GetPrivatePackage().Name(), strings.Title(pkg.Name()))
	return codegen.NewCommentedValueMethod(
		pkg.Path(),
		name,
		managerName,
		/*param=*/ nil,
		[]jen.Code{
			jen.Func().Params(
				jen.Map(jen.String()).Interface(),
			).Params(
				// TODO: Qualify this.
				jen.Op("*").Id(nfp.StructName()),
				jen.Error(),
			),
		},
		[]jen.Code{
			jen.Return(
				jen.Qual(nfp.GetPrivatePackage().Path(), dn),
			),
		},
		jen.Commentf("%s returns the deserialization method for the %q non-functional property in package %q", name, nfp.StructName(), nfp.GetPrivatePackage().Name()))
}

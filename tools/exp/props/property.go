package props

import (
	"fmt"
	"github.com/cjslep/activity/tools/exp/codegen"
	"github.com/dave/jennifer/jen"
)

const (
	// Method names for generated code
	getMethod                 = "Get"
	setMethod                 = "Set"
	hasMethod                 = "Has"
	clearMethod               = "Clear"
	iteratorClearMethod       = "clear"
	isMethod                  = "Is"
	appendMethod              = "Append"
	prependMethod             = "Prepend"
	removeMethod              = "Remove"
	lenMethod                 = "Len"
	swapMethod                = "Swap"
	lessMethod                = "Less"
	kindIndexMethod           = "KindIndex"
	serializeMethod           = "Serialize"
	deserializeMethod         = "Deserialize"
	nameMethod                = "Name"
	serializeIteratorMethod   = "serialize"
	deserializeIteratorMethod = "deserialize"
	isLanguageMapMethod       = "IsLanguageMap"
	hasLanguageMethod         = "HasLanguage"
	getLanguageMethod         = "GetLanguage"
	setLanguageMethod         = "SetLanguage"
	// Member names for generated code
	unknownMemberName = "unknown"
	langMapMember     = "langMap"
)

// join appends a bunch of Go Code together, each on their own line.
func join(s []jen.Code) *jen.Statement {
	r := jen.Empty()
	for i, stmt := range s {
		if i > 0 {
			r.Line()
		}
		r.Add(stmt)
	}
	return r
}

// Identifier determines how a name will appear in documentation and Go code.
type Identifier struct {
	// LowerName is the typical name used in documentation.
	LowerName string
	// CamelName is the typical name used in identifiers in code.
	CamelName string
}

// Kind is data that describes a concrete Go type, how to serialize and
// deserialize such types, compare the types, and other meta-information to use
// during Go code generation.
//
// Only represents values and other types.
type Kind struct {
	Name Identifier
	// ConcreteKind is expected to be properly qualified.
	ConcreteKind *jen.Statement
	Nilable      bool

	// TODO: Untangle the package management mess so that the below do not
	// need to be duplicated.

	// These <FuncName>Fn types are for qualified names of the functions.
	// Expected to always be non-nil: a function is needed to deserialize.
	DeserializeFn *jen.Statement
	// If any of these are nil at generation time, assume to call the method
	// on the object directly (instead of a qualified function).
	SerializeFn *jen.Statement
	LessFn      *jen.Statement

	// The following are only used for values, not types, as actual implementations
	SerializeDef   *codegen.Function
	DeserializeDef *codegen.Function
	LessDef        *codegen.Function
}

func (k Kind) lessFnCode(this, other *jen.Statement) *jen.Statement {
	// LessFn is nil case -- call comparison Less method directly on the LHS
	lessCall := this.Clone().Dot(compareLessMethod).Call(other.Clone())
	if k.LessFn != nil {
		// LessFn is indeed a function -- call this function
		lessCall = k.LessFn.Clone().Call(
			this.Clone(),
			other.Clone(),
		)
	}
	return lessCall
}

func (k Kind) deserializeFnCode(this *jen.Statement) *jen.Statement {
	// LessFn is not nil, this means it is a value.
	if k.LessFn != nil {
		return k.DeserializeFn.Clone().Call(this)
	} else {
		// If LessFn is nil, this means it is a type. Which requires an
		// additional Call.
		return k.DeserializeFn.Clone().Call().Call(this)
	}
}

// PropertyGenerator is a common base struct used in both Functional and
// NonFunctional ActivityStreams properties. It provides common naming patterns,
// logic, and common Go code to be generated.
//
// It also properly handles the concept of generating Go code for property
// iterators, which are needed for NonFunctional properties.
//
// TODO: Make this type private
type PropertyGenerator struct {
	vocabName      string
	managerMethods []*codegen.Method
	// TODO: Make these private
	PackageManager        *PackageManager
	Name                  Identifier
	Comment               string
	Kinds                 []Kind
	HasNaturalLanguageMap bool
	asIterator            bool
}

// VocabName returns this property's vocabulary name.
func (p *PropertyGenerator) VocabName() string {
	return p.vocabName
}

// GetPrivatePackage gets this property's private Package.
func (p *PropertyGenerator) GetPrivatePackage() Package {
	return p.PackageManager.PrivatePackage()
}

// GetPublicPackage gets this property's public Package.
func (p *PropertyGenerator) GetPublicPackage() Package {
	return p.PackageManager.PublicPackage()
}

// SetKindFns allows TypeGenerators to later notify this Property what functions
// to use when generating the serialization code.
//
// The name parameter must match the LowerName of an Identifier.
//
// This feels very hacky.
func (p *PropertyGenerator) SetKindFns(name string, qualKind *jen.Statement, deser *codegen.Method) error {
	for i, kind := range p.Kinds {
		if kind.Name.LowerName == name {
			if kind.SerializeFn != nil || kind.DeserializeFn != nil || kind.LessFn != nil {
				return fmt.Errorf("property kind already has serialization functions set for %q", name)
			}
			kind.ConcreteKind = qualKind
			kind.DeserializeFn = deser.On(managerInitName())
			p.managerMethods = append(p.managerMethods, deser)
			p.Kinds[i] = kind
			return nil
		}
	}
	return fmt.Errorf("cannot find property kind %q", name)
}

// getAllManagerMethods returns the list of manager methods used by this
// property.
func (p *PropertyGenerator) getAllManagerMethods() []*codegen.Method {
	return p.managerMethods
}

// StructName returns the name of the type, which may or may not be a struct,
// to generate.
func (p *PropertyGenerator) StructName() string {
	if p.asIterator {
		return p.Name.CamelName
	}
	return fmt.Sprintf("%sProperty", p.Name.CamelName)
}

// InterfaceName returns the interface name of the property type.
func (p *PropertyGenerator) InterfaceName() string {
	return fmt.Sprintf("%sInterface", p.StructName())
}

// PropertyName returns the name of this property, as defined in
// specifications. It is not suitable for use in generated code function
// identifiers.
func (p *PropertyGenerator) PropertyName() string {
	return p.Name.LowerName
}

// Comments returns the comment for this property.
func (p *PropertyGenerator) Comments() string {
	return p.Comment
}

// DeserializeFnName returns the identifier of the function that deserializes
// raw JSON into the generated Go type.
func (p *PropertyGenerator) DeserializeFnName() string {
	if p.asIterator {
		return fmt.Sprintf("%s%s", deserializeIteratorMethod, p.Name.CamelName)
	}
	return fmt.Sprintf("%s%sProperty", deserializeMethod, p.Name.CamelName)
}

// getFnName returns the identifier of the function that fetches concrete types
// of the property.
func (p *PropertyGenerator) getFnName(i int) string {
	if len(p.Kinds) == 1 {
		return getMethod
	}
	return fmt.Sprintf("%s%s", getMethod, p.kindCamelName(i))
}

// setFnName returns the identifier of the function that sets concrete types
// of the property.
func (p *PropertyGenerator) setFnName(i int) string {
	if len(p.Kinds) == 1 {
		return setMethod
	}
	return fmt.Sprintf("%s%s", setMethod, p.kindCamelName(i))
}

// serializeFnName returns the identifier of the function that serializes the
// generated Go type into raw JSON.
func (p *PropertyGenerator) serializeFnName() string {
	if p.asIterator {
		return serializeIteratorMethod
	}
	return serializeMethod
}

// kindCamelName returns an identifier-friendly name for the kind at the
// specified index.
//
// It will panic if 'i' is out of range.
func (p *PropertyGenerator) kindCamelName(i int) string {
	return p.Kinds[i].Name.CamelName
}

// memberName returns the identifier to use for the kind at the specified index.
//
// It will panic if 'i' is out of range.
func (p *PropertyGenerator) memberName(i int) string {
	return fmt.Sprintf("%sMember", p.Kinds[i].Name.LowerName)
}

// hasMemberName returns the identifier to use for struct members that determine
// whether non-nilable types have been set. Panics if called for a Kind that is
// nilable.
func (p *PropertyGenerator) hasMemberName(i int) string {
	if len(p.Kinds) == 1 && p.Kinds[0].Nilable {
		panic("PropertyGenerator.hasMemberName called for nilable single value")
	}
	return fmt.Sprintf("has%sMember", p.Kinds[i].Name.CamelName)
}

// clearMethodName returns the identifier to use for methods that clear all
// values from the property.
func (p *PropertyGenerator) clearMethodName() string {
	if p.asIterator {
		return iteratorClearMethod
	}
	return clearMethod
}

// commonMethods returns methods common to every property.
func (p *PropertyGenerator) commonMethods() []*codegen.Method {
	return []*codegen.Method{
		codegen.NewCommentedValueMethod(
			p.GetPrivatePackage().Path(),
			nameMethod,
			p.StructName(),
			/*params=*/ nil,
			[]jen.Code{jen.String()},
			[]jen.Code{
				jen.Return(
					jen.Lit(p.PropertyName()),
				),
			},
			jen.Commentf("%s returns the name of this property: %q.", nameMethod, p.PropertyName()),
		),
	}
}

// isMethodName returns the identifier to use for methods that determine if a
// property holds a specific Kind of value.
func (p *PropertyGenerator) isMethodName(i int) string {
	return fmt.Sprintf("%s%s", isMethod, p.kindCamelName(i))
}

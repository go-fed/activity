// Package gen contains the libraries and algorithms used to generate the
// code for the activitystreams package.
package gen

import (
	"bytes"
	"fmt"
	"github.com/go-fed/activity/tools/defs"
	"github.com/go-fed/activity/tools/vocab/gen"
	"go/format"
	"strings"
)

const (
	resolverName       = "Resolver"
	dispatchFnName     = "dispatch"
	rawMemberName      = "raw"
	rawFunctionName    = "Raw"
	convenienceComment = " This is a convenience wrapper of a type with the same name in the vocab package. Accessing it with the " + rawFunctionName + " function allows direct manipulaton of the object, and does not provide the same integrity guarantees as this package."
)

const (
	UnknownLanguage = "und"
)

func GenerateConvenienceTypes(types []*defs.Type) ([]byte, error) {
	p := generatePackageDefinition()
	p.Defs = append(p.Defs, generateResolver(types))
	for _, t := range types {
		funcs, defs := generateDefinitions(t)
		p.F = append(p.F, funcs...)
		p.Defs = append(p.Defs, defs...)
	}
	return format.Source([]byte(p.Generate()))
}

func generatePackageDefinition() *defs.PackageDef {
	return &defs.PackageDef{
		Name:    "streams",
		Comment: "Package streams is a convenience wrapper around the raw ActivityStream vocabulary. This package is code-generated to permit more powerful expressions and manipulations of the ActivityStreams Vocabulary types. This package also does not permit use of 'unknown' properties, or those that are outside of the ActivityStream Vocabulary specification. However, it still correctly propagates them when repeatedly re-and-de-serialized. Custom extensions of the vocabulary are supported by modifying the data definitions in the generation tool and rerunning it. Do not modify this package directly.",
		Imports: []string{"fmt", "github.com/go-fed/activity/vocab", "net/url", "time"},
		Raw: `type Resolution int

const (
	Resolved Resolution = iota
	RawResolutionNeeded
	Unresolved
)

type Presence int

const (
	NoPresence Presence = iota
	ConvenientPresence
	RawPresence
)`,
	}
}

func generateResolver(types []*defs.Type) *defs.StructDef {
	this := &defs.StructDef{
		Typename: resolverName,
		Comment:  fmt.Sprintf("%s contains callback functions to execute when it Deserializes a raw map[string]interface{} into a concrete type. Clients can set only the callbacks they care about and handle the resulting concrete type.", resolverName),
	}
	for _, t := range types {
		name := fmt.Sprintf("%sCallback", t.Name)
		sig := fmt.Sprintf("func(*%s) error", t.Name)
		c := fmt.Sprintf("Callback function for the %s type", t.Name)
		this.M = append(this.M, &defs.StructMember{name, sig, c})
	}
	this.M = append(this.M, &defs.StructMember{"AnyObjectCallback", "func(vocab.ObjectType) error", "Callback function for any type that satisfies the vocab.ObjectType interface. Note that this will be called in addition to the specific type callbacks."})
	this.M = append(this.M, &defs.StructMember{"AnyLinkCallback", "func(vocab.LinkType) error", "Callback function for any type that satisfies the vocab.LinkType interface. Note that this will be called in addition to the specific type callbacks."})
	this.M = append(this.M, &defs.StructMember{"AnyActivityCallback", "func(vocab.ActivityType) error", "Callback function for any type that satisfies the vocab.ActivityType interface. Note that this will be called in addition to the specific type callbacks."})
	this.F = []*defs.MemberFunctionDef{
		{
			Name:    dispatchFnName,
			Comment: fmt.Sprintf("%s routes the given type to the appropriate %s callback.", dispatchFnName, resolverName),
			P:       this,
			Args:    []*defs.FunctionVarDef{{"i", "interface{}"}},
			Return:  []*defs.FunctionVarDef{{"handled", "bool"}, {"err", "error"}},
			Body: func() string {
				var b bytes.Buffer
				for _, t := range types {
					name := fmt.Sprintf("%sCallback", t.Name)
					b.WriteString(fmt.Sprintf("// Begin generateResolver for type '%s'\n", t.Name))
					b.WriteString(fmt.Sprintf("if rawV, ok := i.(*vocab.%s); ok {\n", t.Name))
					b.WriteString(fmt.Sprintf("if t.%s != nil {\n", name))
					b.WriteString(fmt.Sprintf("v := &%s{%s: rawV}\n", t.Name, rawMemberName))
					b.WriteString(fmt.Sprintf("return true, t.%s(v)\n", name))
					b.WriteString("} else {\n")
					b.WriteString("return false, nil\n")
					b.WriteString("}\n")
					b.WriteString("}\n")
					b.WriteString(fmt.Sprintf("// End generateResolver for type '%s'\n", t.Name))
				}
				b.WriteString("if obj, ok := i.(vocab.ObjectType); ok {\n")
				b.WriteString("if t.AnyObjectCallback != nil {\n")
				b.WriteString("return true, t.AnyObjectCallback(obj)\n")
				b.WriteString("}\n")
				b.WriteString("}\n")
				b.WriteString("if link, ok := i.(vocab.LinkType); ok {\n")
				b.WriteString("if t.AnyLinkCallback != nil {\n")
				b.WriteString("return true, t.AnyLinkCallback(link)\n")
				b.WriteString("}\n")
				b.WriteString("}\n")
				b.WriteString("if activity, ok := i.(vocab.ActivityType); ok {\n")
				b.WriteString("if t.AnyActivityCallback != nil {\n")
				b.WriteString("return true, t.AnyActivityCallback(activity)\n")
				b.WriteString("}\n")
				b.WriteString("}\n")
				b.WriteString("return false, fmt.Errorf(\"The interface did not match any known types: %T\", i)\n")
				return b.String()
			},
		},
		{
			Name:    "Deserialize",
			Comment: "Determines which concrete type to deserialize this json-unmarshalled item into, returning an error if it cannot determine which type to deserialize into. The appropriate callback, if present, will then be invoked with the concrete deserialized type. If the callback function returns an error, it is passed back through Deserialize.",
			P:       this,
			Args:    []*defs.FunctionVarDef{{"m", "map[string]interface{}"}},
			Return:  []*defs.FunctionVarDef{{"err", "error"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("var typeStringVals []string\n")
				b.WriteString("typeInterface, ok := m[\"type\"]\n")
				b.WriteString("if !ok {\n")
				b.WriteString("return fmt.Errorf(\"Cannot determine type: missing 'type' property\")\n")
				b.WriteString("}\n")
				b.WriteString("if typeStr, ok := typeInterface.(string); ok {\n")
				b.WriteString("typeStringVals = append(typeStringVals, typeStr)\n")
				b.WriteString("} else if typeSlice, ok := typeInterface.([]interface{}); ok {\n")
				b.WriteString("for _, elem := range typeSlice {")
				b.WriteString("if typeStr, ok := elem.(string); ok {\n")
				b.WriteString("typeStringVals = append(typeStringVals, typeStr)\n")
				b.WriteString("}\n")
				b.WriteString("}\n")
				b.WriteString("if len(typeStringVals) == 0 {\n")
				b.WriteString("return fmt.Errorf(\"Cannot determine type: 'type' property is []interface{} with no string elements: %+v\", typeInterface)\n")
				b.WriteString("}\n")
				b.WriteString("} else {\n")
				b.WriteString("return fmt.Errorf(\"Cannot determine type: 'type' property is not string nor []interface{}: %T\", typeInterface)\n")
				b.WriteString("}\n")
				for _, t := range types {
					name := fmt.Sprintf("%sCallback", t.Name)
					b.WriteString(fmt.Sprintf("// Begin generateResolver for type '%s'\n", t.Name))
					b.WriteString("for _, typeName := range typeStringVals {\n")
					b.WriteString(fmt.Sprintf("if typeName == \"%s\" {\n", t.Name))
					b.WriteString(fmt.Sprintf("if t.%s != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {\n", name))
					b.WriteString(fmt.Sprintf("v := &vocab.%s{}\n", t.Name))
					b.WriteString("if err := v.Deserialize(m); err != nil {\n")
					b.WriteString("return err\n")
					b.WriteString("}\n")
					b.WriteString(fmt.Sprintf("as := &%s{v}\n", t.Name))
					b.WriteString(fmt.Sprintf("if t.%s != nil {\n", name))
					b.WriteString(fmt.Sprintf("if err := t.%s(as); err != nil {\n", name))
					b.WriteString("return err\n")
					b.WriteString("}\n")
					b.WriteString("}\n")
					b.WriteString("var i interface{} = v\n")
					b.WriteString("if obj, ok := i.(vocab.ObjectType); ok {\n")
					b.WriteString("if t.AnyObjectCallback != nil {\n")
					b.WriteString("if err := t.AnyObjectCallback(obj); err != nil {\n")
					b.WriteString("return err\n")
					b.WriteString("}\n")
					b.WriteString("}\n")
					b.WriteString("}\n")
					b.WriteString("if link, ok := i.(vocab.LinkType); ok {\n")
					b.WriteString("if t.AnyLinkCallback != nil {\n")
					b.WriteString("if err := t.AnyLinkCallback(link); err != nil {\n")
					b.WriteString("return err\n")
					b.WriteString("}\n")
					b.WriteString("}\n")
					b.WriteString("}\n")
					b.WriteString("if activity, ok := i.(vocab.ActivityType); ok {\n")
					b.WriteString("if t.AnyActivityCallback != nil {\n")
					b.WriteString("if err := t.AnyActivityCallback(activity); err != nil {\n")
					b.WriteString("return err\n")
					b.WriteString("}\n")
					b.WriteString("}\n")
					b.WriteString("}\n")
					b.WriteString("return nil\n")
					b.WriteString("} else {\n")
					b.WriteString("return nil\n")
					b.WriteString("}\n")
					b.WriteString("}\n")
					b.WriteString("}\n")
					b.WriteString(fmt.Sprintf("// End generateResolver for type '%s'\n", t.Name))
				}
				b.WriteString("return fmt.Errorf(\"The 'type' property did not match any known types: %+v\", typeStringVals)\n")
				return b.String()
			},
		},
	}
	return this
}

func generateDefinitions(t *defs.Type) (fd []*defs.FunctionDef, sd []*defs.StructDef) {
	this := &defs.StructDef{
		Typename: t.Name,
		Comment:  t.Notes + convenienceComment,
		M:        []*defs.StructMember{{rawMemberName, "*vocab." + t.Name, "The raw type from the vocab package"}},
	}
	sd = append(sd, this)
	this.F = append(this.F, []*defs.MemberFunctionDef{
		{
			Name:    rawFunctionName,
			Comment: "Raw returns the vocab type for manaual manipulation. Note that manipulating the underlying type to be in an inconsistent state may cause this convenience type's methods to later fail.",
			P:       this,
			Return:  []*defs.FunctionVarDef{{"n", "*vocab." + t.Name}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString(fmt.Sprintf("return t.%s\n", rawMemberName))
				return b.String()
			},
		},
		{
			Name:    "Serialize",
			Comment: "Serialize turns this object into a map[string]interface{}.",
			P:       this,
			Return:  []*defs.FunctionVarDef{{"m", "map[string]interface{}"}, {"err", "error"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString(fmt.Sprintf("return t.%s.Serialize()\n", rawMemberName))
				return b.String()
			},
		},
	}...)
	for _, p := range t.GetProperties() {
		generatePropertyHelpers(p, this)
	}
	return
}

func generatePropertyHelpers(p *defs.PropertyType, this *defs.StructDef) {
	if p.Functional {
		generateFunctionalPropertyHelper(p, this)
	} else {
		generateNonFunctionalPropertyHelper(p, this)
	}
	if p.NaturalLanguageMap {
		generateNaturalLanguageConvenience(p, this)
	}
}

func generateNaturalLanguageConvenience(p *defs.PropertyType, this *defs.StructDef) {
	titleName := strings.Title(p.Name)
	this.F = append(this.F, []*defs.MemberFunctionDef{
		{
			Name:    fmt.Sprintf("%sLanguages", titleName),
			Comment: fmt.Sprintf("%sLanguages returns all languages for this property's language mapping, or nil if there are none.", titleName),
			P:       this,
			Return:  []*defs.FunctionVarDef{{"l", "[]string"}},
			Body: func() string {
				return fmt.Sprintf("return t.%s.%sMapLanguages()\n", rawMemberName, titleName)
			},
		},
		{
			Name:    fmt.Sprintf("Get%sForLanguage", titleName),
			Comment: fmt.Sprintf("Get%sMap retrieves the value of '%s' for the specified language, or an empty string if it does not exist", titleName, p.Name),
			P:       this,
			Args:    []*defs.FunctionVarDef{{"l", "string"}},
			Return:  []*defs.FunctionVarDef{{"v", "string"}},
			Body: func() string {
				return fmt.Sprintf("return t.%s.Get%sMap(l)\n", rawMemberName, titleName)
			},
		},
		{
			Name:    fmt.Sprintf("Set%sForLanguage", titleName),
			Comment: fmt.Sprintf("Set%sForLanguage sets the value of '%s' for the specified language", titleName, p.Name),
			P:       this,
			Args:    []*defs.FunctionVarDef{{"l", "string"}, {"v", "string"}},
			Body: func() string {
				return fmt.Sprintf("t.%s.Set%sMap(l, v)\n", rawMemberName, titleName)
			},
		},
	}...)
}

func generateFunctionalPropertyHelper(p *defs.PropertyType, this *defs.StructDef) {
	ref := p.Range[0]
	if p.PreferIRIConvenience {
		generateFunctionalIRI(p, this)
	} else if ref.T != nil {
		if rangeHasObjectAndLink(p.Range) {
			generateFunctionalObjectLink(p, this, rangeHasMoreThanJustObjectAndLink(p.Range))
		} else {
			generateFunctionalPropertyType(p, this, ref.T, len(p.Range) == 1)
		}
	} else if ref.V != nil {
		generateFunctionalPropertyValue(p, this, ref.V, len(p.Range) == 1)
	} else if ref.Any {
		generateFunctionalPropertyAny(p, this)
	} else {
		panic("Unhandled RangeReference: T == V == nil, Any == false")
	}
}

func generateFunctionalIRI(p *defs.PropertyType, this *defs.StructDef) {
	kind := deref(defs.IriValueType.DefinitionType)
	titleName := strings.Title(p.Name)
	onlyType := len(p.Range) == 1
	iri := "IRI"
	if defs.HasAnyURI(p.Range) {
		iri = strings.Title(defs.AnyURIValueTypeName())
	}
	this.F = append(this.F, []*defs.MemberFunctionDef{
		{
			Name:    fmt.Sprintf("Get%s", titleName),
			Comment: fmt.Sprintf("Get%s attempts to get this '%s' property as a %s. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.", titleName, p.Name, kind),
			P:       this,
			Return:  []*defs.FunctionVarDef{{"r", "Resolution"}, {"k", kind}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("r = Unresolved\n")
				b.WriteString("handled := false\n")
				if onlyType {
					b.WriteString(fmt.Sprintf("if t.%s.Has%s%s() {\n", rawMemberName, titleName, iri))
					b.WriteString(fmt.Sprintf("k = t.%s.Get%s%s()\n", rawMemberName, titleName, iri))
				} else {
					b.WriteString(fmt.Sprintf("if t.%s.Is%s%s() {\n", rawMemberName, titleName, iri))
					b.WriteString(fmt.Sprintf("k = t.%s.Get%s%s()\n", rawMemberName, titleName, iri))
				}
				b.WriteString("if handled {\n")
				b.WriteString("r = Resolved\n")
				b.WriteString("}\n")
				b.WriteString("} ")
				for _, elem := range p.Range {
					if elem.V != nil && (defs.IsIRIValueType(elem.V) || defs.IsAnyURIValueType(elem.V)) {
						continue
					}
					b.WriteString(fmt.Sprintf("else if t.%s.Is%s%s() {\n", rawMemberName, titleName, strings.Title(gen.Name(elem))))
					b.WriteString("r = RawResolutionNeeded\n")
					b.WriteString("}")
				}
				b.WriteString("\nreturn\n")
				return b.String()
			},
		},
		{
			Name:    fmt.Sprintf("Has%s", titleName),
			Comment: fmt.Sprintf("Has%s returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.", titleName),
			P:       this,
			Return:  []*defs.FunctionVarDef{{"p", "Presence"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("p = NoPresence\n")
				if onlyType {
					b.WriteString(fmt.Sprintf("if t.%s.Has%s%s() {\n", rawMemberName, titleName, iri))
				} else {
					b.WriteString(fmt.Sprintf("if t.%s.Is%s%s() {\n", rawMemberName, titleName, iri))
				}
				b.WriteString("p = ConvenientPresence\n")
				b.WriteString("} ")
				for _, elem := range p.Range {
					if elem.V != nil && (defs.IsIRIValueType(elem.V) || defs.IsAnyURIValueType(elem.V)) {
						continue
					}
					b.WriteString(fmt.Sprintf("else if t.%s.Is%s%s() {\n", rawMemberName, titleName, strings.Title(gen.Name(elem))))
					b.WriteString("p = RawPresence\n")
					b.WriteString("}")
				}
				b.WriteString("\nreturn\n")
				return b.String()
			},
		},
		{
			Name:    fmt.Sprintf("Set%s", titleName),
			Comment: fmt.Sprintf("Set%s sets the value for property '%s'.", titleName, p.Name),
			P:       this,
			Args:    []*defs.FunctionVarDef{{"k", kind}},
			Body: func() string {
				var b bytes.Buffer
				if onlyType {
					b.WriteString(fmt.Sprintf("t.%s.Set%s%s(k)\n", rawMemberName, titleName, iri))
				} else {
					b.WriteString(fmt.Sprintf("t.%s.Set%s%s(k)\n", rawMemberName, titleName, iri))
				}
				return b.String()
			},
		},
	}...)
}

func generateFunctionalObjectLink(p *defs.PropertyType, this *defs.StructDef, hasMoreTypes bool) {
	titleName := strings.Title(p.Name)
	this.F = append(this.F, []*defs.MemberFunctionDef{
		{
			Name:    fmt.Sprintf("Resolve%s", titleName),
			Comment: fmt.Sprintf("Resolve%s passes the actual concrete type to the resolver for handing property %s. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.", titleName, p.Name),
			P:       this,
			Args:    []*defs.FunctionVarDef{{"r", "*" + resolverName}},
			Return:  []*defs.FunctionVarDef{{"s", "Resolution"}, {"err", "error"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("s = Unresolved\n")
				b.WriteString("handled := false\n")
				b.WriteString(fmt.Sprintf("if t.%s.Is%sObject() {\n", rawMemberName, titleName))
				b.WriteString(fmt.Sprintf("handled, err = r.%s(t.%s.Get%sObject())\n", dispatchFnName, rawMemberName, titleName))
				b.WriteString("if handled {\n")
				b.WriteString("s = Resolved\n")
				b.WriteString("}\n")
				b.WriteString(fmt.Sprintf("} else if t.%s.Is%sLink() {\n", rawMemberName, titleName))
				b.WriteString(fmt.Sprintf("handled, err = r.%s(t.%s.Get%sLink())\n", dispatchFnName, rawMemberName, titleName))
				b.WriteString("if handled {\n")
				b.WriteString("s = Resolved\n")
				b.WriteString("}\n")
				if !hasMoreTypes {
					b.WriteString("} else {\n")
					b.WriteString("s = Resolved\n")
					b.WriteString("}\n")
					b.WriteString("return\n")
				} else {
					b.WriteString("} ")
					for _, elem := range p.Range {
						if elem.T != nil && (elem.T.Name == "Object" || elem.T.Name == "Link") {
							continue
						}
						b.WriteString(fmt.Sprintf("else if t.%s.Is%s%s() {\n", rawMemberName, titleName, strings.Title(gen.Name(elem))))
						b.WriteString("s = RawResolutionNeeded\n")
						b.WriteString("}")
					}
					b.WriteString("\nreturn\n")
				}
				return b.String()
			},
		},
		{
			Name:    fmt.Sprintf("Has%s", titleName),
			Comment: fmt.Sprintf("Has%s returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.", titleName),
			P:       this,
			Return:  []*defs.FunctionVarDef{{"p", "Presence"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("p = NoPresence\n")
				b.WriteString(fmt.Sprintf("if t.%s.Is%sObject() {\n", rawMemberName, titleName))
				b.WriteString("p = ConvenientPresence\n")
				b.WriteString(fmt.Sprintf("} else if t.%s.Is%sLink() {\n", rawMemberName, titleName))
				b.WriteString("p = ConvenientPresence\n")
				if !hasMoreTypes {
					b.WriteString("}\n")
					b.WriteString("return\n")
				} else {
					b.WriteString("} ")
					for _, elem := range p.Range {
						if elem.T != nil && (elem.T.Name == "Object" || elem.T.Name == "Link") {
							continue
						}
						b.WriteString(fmt.Sprintf("else if t.%s.Is%s%s() {\n", rawMemberName, titleName, strings.Title(gen.Name(elem))))
						b.WriteString("p = RawPresence\n")
						b.WriteString("}")
					}
					b.WriteString("\nreturn\n")
				}
				return b.String()
			},
		},
		{
			Name:    fmt.Sprintf("Set%s", titleName),
			Comment: fmt.Sprintf("Set%s sets this value to be an 'Object' type.", titleName),
			P:       this,
			Args:    []*defs.FunctionVarDef{{"i", "vocab.ObjectType"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString(fmt.Sprintf("t.%s.Set%sObject(i)\n", rawMemberName, titleName))
				return b.String()
			},
		},
		{
			Name:    fmt.Sprintf("Set%sLink", titleName),
			Comment: fmt.Sprintf("Set%sLink sets this value to be an 'Link' type.", titleName),
			P:       this,
			Args:    []*defs.FunctionVarDef{{"i", "vocab.LinkType"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString(fmt.Sprintf("t.%s.Set%sLink(i)\n", rawMemberName, titleName))
				return b.String()
			},
		},
	}...)
}

func generateFunctionalPropertyType(p *defs.PropertyType, this *defs.StructDef, ref *defs.Type, onlyType bool) {
	additionalTitleName := strings.Title(gen.Name(p.Range[0]))
	if !onlyType && defs.IsOnlyOtherPropertyBesidesIRI(0, p.Range) {
		additionalTitleName = ""
	}
	titleName := strings.Title(p.Name)
	this.F = append(this.F, []*defs.MemberFunctionDef{
		{
			Name:    fmt.Sprintf("Resolve%s", titleName),
			Comment: fmt.Sprintf("Resolve%s passes the actual concrete type to the resolver for handing property %s. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.", titleName, p.Name),
			P:       this,
			Args:    []*defs.FunctionVarDef{{"r", "*" + resolverName}},
			Return:  []*defs.FunctionVarDef{{"s", "Resolution"}, {"err", "error"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("s = Unresolved\n")
				b.WriteString("handled := false\n")
				if onlyType {
					b.WriteString(fmt.Sprintf("if t.%s.Has%s() {\n", rawMemberName, titleName))
					b.WriteString(fmt.Sprintf("handled, err = r.%s(t.%s.Get%s())\n", dispatchFnName, rawMemberName, titleName))
				} else {
					b.WriteString(fmt.Sprintf("if t.%s.Is%s%s() {\n", rawMemberName, titleName, additionalTitleName))
					b.WriteString(fmt.Sprintf("handled, err = r.%s(t.%s.Get%s%s())\n", dispatchFnName, rawMemberName, titleName, additionalTitleName))
				}
				b.WriteString("if handled {\n")
				b.WriteString("s = Resolved\n")
				b.WriteString("}\n")
				b.WriteString("} ")
				for idx, elem := range p.Range {
					if idx == 0 {
						continue
					}
					b.WriteString(fmt.Sprintf("else if t.%s.Is%s%s() {\n", rawMemberName, titleName, strings.Title(gen.Name(elem))))
					b.WriteString("s = RawResolutionNeeded\n")
					b.WriteString("}")
				}
				b.WriteString("\nreturn\n")
				return b.String()
			},
		},
		{
			Name:    fmt.Sprintf("Has%s", titleName),
			Comment: fmt.Sprintf("Has%s returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.", titleName),
			P:       this,
			Return:  []*defs.FunctionVarDef{{"p", "Presence"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("p = NoPresence\n")
				if onlyType {
					b.WriteString(fmt.Sprintf("if t.%s.Has%s() {\n", rawMemberName, titleName))
				} else {
					b.WriteString(fmt.Sprintf("if t.%s.Is%s%s() {\n", rawMemberName, titleName, additionalTitleName))
				}
				b.WriteString("p = ConvenientPresence\n")
				b.WriteString("} ")
				for idx, elem := range p.Range {
					if idx == 0 {
						continue
					}
					b.WriteString(fmt.Sprintf("else if t.%s.Is%s%s() {\n", rawMemberName, titleName, strings.Title(gen.Name(elem))))
					b.WriteString("p = RawPresence\n")
					b.WriteString("}")
				}
				b.WriteString("\nreturn\n")
				return b.String()
			},
		},
		{
			Name:    fmt.Sprintf("Set%s", titleName),
			Comment: fmt.Sprintf("Set%s sets this value to be a '%s' type.", titleName, ref.Name),
			P:       this,
			Args:    []*defs.FunctionVarDef{{"i", "vocab." + gen.InterfaceName(ref)}},
			Body: func() string {
				var b bytes.Buffer
				if onlyType {
					b.WriteString(fmt.Sprintf("t.%s.Set%s(i)\n", rawMemberName, titleName))
				} else {
					b.WriteString(fmt.Sprintf("t.%s.Set%s%s(i)\n", rawMemberName, titleName, additionalTitleName))
				}
				return b.String()
			},
		},
	}...)
}

func generateFunctionalPropertyValue(p *defs.PropertyType, this *defs.StructDef, value *defs.ValueType, onlyType bool) {
	additionalTitleName := strings.Title(gen.Name(p.Range[0]))
	if !onlyType && defs.IsOnlyOtherPropertyBesidesIRI(0, p.Range) {
		additionalTitleName = ""
	}
	kind := deref(value.DefinitionType)
	titleName := strings.Title(p.Name)
	this.F = append(this.F, []*defs.MemberFunctionDef{
		{
			Name:    fmt.Sprintf("Get%s", titleName),
			Comment: fmt.Sprintf("Get%s attempts to get this '%s' property as a %s. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.", titleName, p.Name, kind),
			P:       this,
			Return:  []*defs.FunctionVarDef{{"r", "Resolution"}, {"k", kind}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("r = Unresolved\n")
				b.WriteString("handled := false\n")
				if onlyType {
					b.WriteString(fmt.Sprintf("if t.%s.Has%s() {\n", rawMemberName, titleName))
					b.WriteString(fmt.Sprintf("k = t.%s.Get%s()\n", rawMemberName, titleName))
				} else {
					b.WriteString(fmt.Sprintf("if t.%s.Is%s%s() {\n", rawMemberName, titleName, additionalTitleName))
					b.WriteString(fmt.Sprintf("k = t.%s.Get%s%s()\n", rawMemberName, titleName, additionalTitleName))
				}
				b.WriteString("if handled {\n")
				b.WriteString("r = Resolved\n")
				b.WriteString("}\n")
				b.WriteString("} ")
				for idx, elem := range p.Range {
					if idx == 0 {
						continue
					}
					b.WriteString(fmt.Sprintf("else if t.%s.Is%s%s() {\n", rawMemberName, titleName, strings.Title(gen.Name(elem))))
					b.WriteString("r = RawResolutionNeeded\n")
					b.WriteString("}")
				}
				b.WriteString("\nreturn\n")
				return b.String()
			},
		},
		{
			Name:    fmt.Sprintf("Has%s", titleName),
			Comment: fmt.Sprintf("Has%s returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.", titleName),
			P:       this,
			Return:  []*defs.FunctionVarDef{{"p", "Presence"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("p = NoPresence\n")
				if onlyType {
					b.WriteString(fmt.Sprintf("if t.%s.Has%s() {\n", rawMemberName, titleName))
				} else {
					b.WriteString(fmt.Sprintf("if t.%s.Is%s%s() {\n", rawMemberName, titleName, additionalTitleName))
				}
				b.WriteString("p = ConvenientPresence\n")
				b.WriteString("} ")
				for idx, elem := range p.Range {
					if idx == 0 {
						continue
					}
					b.WriteString(fmt.Sprintf("else if t.%s.Is%s%s() {\n", rawMemberName, titleName, strings.Title(gen.Name(elem))))
					b.WriteString("p = RawPresence\n")
					b.WriteString("}")
				}
				b.WriteString("\nreturn\n")
				return b.String()
			},
		},
		{
			Name:    fmt.Sprintf("Set%s", titleName),
			Comment: fmt.Sprintf("Set%s sets the value for property '%s'.", titleName, p.Name),
			P:       this,
			Args:    []*defs.FunctionVarDef{{"k", kind}},
			Body: func() string {
				var b bytes.Buffer
				if onlyType {
					b.WriteString(fmt.Sprintf("t.%s.Set%s(k)\n", rawMemberName, titleName))
				} else {
					b.WriteString(fmt.Sprintf("t.%s.Set%s%s(k)\n", rawMemberName, titleName, additionalTitleName))
				}
				return b.String()
			},
		},
	}...)
}

func generateFunctionalPropertyAny(p *defs.PropertyType, this *defs.StructDef) {
	titleName := strings.Title(p.Name)
	this.F = append(this.F, []*defs.MemberFunctionDef{
		{
			Name:    fmt.Sprintf("Get%s", titleName),
			Comment: fmt.Sprintf("Get%s attempts to get this '%s' property as a string. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.", titleName, p.Name),
			P:       this,
			Return:  []*defs.FunctionVarDef{{"r", "Resolution"}, {"s", "string"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("r = Unresolved\n")
				b.WriteString(fmt.Sprintf("if tmp := t.%s.Get%s(); tmp != nil {\n", rawMemberName, titleName))
				b.WriteString("ok := false\n")
				b.WriteString("if s, ok = tmp.(string); ok {\n")
				b.WriteString("r = Resolved\n")
				b.WriteString("} else {\n")
				b.WriteString("r = RawResolutionNeeded\n")
				b.WriteString("}\n")
				b.WriteString("}\n")
				b.WriteString("return\n")
				return b.String()
			},
		},
		{
			Name:    fmt.Sprintf("Set%s", titleName),
			Comment: fmt.Sprintf("Set%s sets the value for property '%s'.", titleName, p.Name),
			P:       this,
			Args:    []*defs.FunctionVarDef{{"i", "interface{}"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString(fmt.Sprintf("t.%s.Set%s(i)\n", rawMemberName, titleName))
				return b.String()
			},
		},
	}...)
}

func generateNonFunctionalPropertyHelper(p *defs.PropertyType, this *defs.StructDef) {
	ref := p.Range[0]
	if p.PreferIRIConvenience {
		generateNonFunctionalIRI(p, this)
	} else if ref.T != nil {
		if rangeHasObjectAndLink(p.Range) {
			generateNonFunctionalObjectLink(p, this, rangeHasMoreThanJustObjectAndLink(p.Range))
		} else {
			generateNonFunctionalPropertyType(p, this, ref.T, len(p.Range) == 1)
		}
	} else if ref.V != nil {
		generateNonFunctionalPropertyValue(p, this, ref.V, len(p.Range) == 1)
	} else if ref.Any {
		generateNonFunctionalPropertyAny(p, this)
	} else {
		panic("Unhandled RangeReference: T == V == nil, Any == false")
	}
}

func generateNonFunctionalIRI(p *defs.PropertyType, this *defs.StructDef) {
	kind := deref(defs.IriValueType.DefinitionType)
	titleName := strings.Title(p.Name)
	onlyType := len(p.Range) == 1
	iri := "IRI"
	if defs.HasAnyURI(p.Range) {
		iri = strings.Title(defs.AnyURIValueTypeName())
	}
	this.F = append(this.F, []*defs.MemberFunctionDef{
		{
			Name:    fmt.Sprintf("Len%s", titleName),
			Comment: fmt.Sprintf("Len%s returns the number of values this property contains. Each index be used with Has%s to determine if Get%s is safe to call or if raw handling would be needed.", titleName, titleName, titleName, p.Name),
			P:       this,
			Return:  []*defs.FunctionVarDef{{"idx", "int"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString(fmt.Sprintf("return t.%s.%sLen()\n", rawMemberName, titleName))
				return b.String()
			},
		},
		{
			Name:    fmt.Sprintf("Get%s", titleName),
			Comment: fmt.Sprintf("Get%s attempts to get this '%s' property as a %s. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.", titleName, p.Name, kind),
			P:       this,
			Args:    []*defs.FunctionVarDef{{"idx", "int"}},
			Return:  []*defs.FunctionVarDef{{"r", "Resolution"}, {"k", kind}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("r = Unresolved\n")
				b.WriteString("handled := false\n")
				if onlyType {
					b.WriteString(fmt.Sprintf("if /*t.%s.Has%s%s(idx)*/ true {\n", rawMemberName, titleName, iri))
					b.WriteString(fmt.Sprintf("k = t.%s.Get%s%s(idx)\n", rawMemberName, titleName, iri))
				} else {
					b.WriteString(fmt.Sprintf("if t.%s.Is%s%s(idx) {\n", rawMemberName, titleName, iri))
					b.WriteString(fmt.Sprintf("k = t.%s.Get%s%s(idx)\n", rawMemberName, titleName, iri))
				}
				b.WriteString("if handled {\n")
				b.WriteString("r = Resolved\n")
				b.WriteString("}\n")
				b.WriteString("} ")
				for _, elem := range p.Range {
					if elem.V != nil && (defs.IsIRIValueType(elem.V) || defs.IsAnyURIValueType(elem.V)) {
						continue
					}
					b.WriteString(fmt.Sprintf("else if t.%s.Is%s%s(idx) {\n", rawMemberName, titleName, strings.Title(gen.Name(elem))))
					b.WriteString("r = RawResolutionNeeded\n")
					b.WriteString("}")
				}
				b.WriteString("\nreturn\n")
				return b.String()
			},
		},
		{
			Name:    fmt.Sprintf("Add%s", titleName),
			Comment: fmt.Sprintf("Add%s appends the value for property '%s'.", titleName, p.Name),
			P:       this,
			Args:    []*defs.FunctionVarDef{{"k", kind}},
			Body: func() string {
				var b bytes.Buffer
				if onlyType {
					b.WriteString(fmt.Sprintf("t.%s.Add%s%s(k)\n", rawMemberName, titleName, iri))
				} else {
					b.WriteString(fmt.Sprintf("t.%s.Add%s%s(k)\n", rawMemberName, titleName, iri))
				}
				return b.String()
			},
		},
		{
			Name:    fmt.Sprintf("Remove%s", titleName),
			Comment: fmt.Sprintf("Remove%s deletes the value from the specified index for property '%s'.", titleName, p.Name),
			P:       this,
			Args:    []*defs.FunctionVarDef{{"idx", "int"}},
			Body: func() string {
				var b bytes.Buffer
				if onlyType {
					b.WriteString(fmt.Sprintf("t.%s.Remove%s%s(idx)\n", rawMemberName, titleName, iri))
				} else {
					b.WriteString(fmt.Sprintf("t.%s.Remove%s%s(idx)\n", rawMemberName, titleName, iri))
				}
				return b.String()
			},
		},
	}...)
	if !onlyType {
		this.F = append(this.F, []*defs.MemberFunctionDef{
			{
				Name:    fmt.Sprintf("Has%s", titleName),
				Comment: fmt.Sprintf("Has%s returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.", titleName),
				P:       this,
				Args:    []*defs.FunctionVarDef{{"idx", "int"}},
				Return:  []*defs.FunctionVarDef{{"p", "Presence"}},
				Body: func() string {
					var b bytes.Buffer
					b.WriteString("p = NoPresence\n")
					b.WriteString(fmt.Sprintf("if t.%s.Is%s%s(idx) {\n", rawMemberName, titleName, iri))
					b.WriteString("p = ConvenientPresence\n")
					b.WriteString("} ")
					for idx, elem := range p.Range {
						if idx == 0 {
							continue
						}
						b.WriteString(fmt.Sprintf("else if t.%s.Is%s%s(idx) {\n", rawMemberName, titleName, strings.Title(gen.Name(elem))))
						b.WriteString("p = RawPresence\n")
						b.WriteString("}")
					}
					b.WriteString("\nreturn\n")
					return b.String()
				},
			},
		}...)
	}
}

func generateNonFunctionalObjectLink(p *defs.PropertyType, this *defs.StructDef, hasMoreTypes bool) {
	titleName := strings.Title(p.Name)
	this.F = append(this.F, []*defs.MemberFunctionDef{
		{
			Name:    fmt.Sprintf("Len%s", titleName),
			Comment: fmt.Sprintf("Len%s returns the number of values this property contains. Each index be used with Has%s to determine if Resolve%s is safe to call or if raw handling would be needed.", titleName, titleName, titleName, p.Name),
			P:       this,
			Return:  []*defs.FunctionVarDef{{"idx", "int"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString(fmt.Sprintf("return t.%s.%sLen()\n", rawMemberName, titleName))
				return b.String()
			},
		},
		{
			Name:    fmt.Sprintf("Resolve%s", titleName),
			Comment: fmt.Sprintf("Resolve%s passes the actual concrete type to the resolver for handing property %s. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.", titleName, p.Name),
			P:       this,
			Args:    []*defs.FunctionVarDef{{"r", "*" + resolverName}, {"idx", "int"}},
			Return:  []*defs.FunctionVarDef{{"s", "Resolution"}, {"err", "error"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("s = Unresolved\n")
				b.WriteString("handled := false\n")
				b.WriteString(fmt.Sprintf("if t.%s.Is%sObject(idx) {\n", rawMemberName, titleName))
				b.WriteString(fmt.Sprintf("handled, err = r.%s(t.%s.Get%sObject(idx))\n", dispatchFnName, rawMemberName, titleName))
				b.WriteString("if handled {\n")
				b.WriteString("s = Resolved\n")
				b.WriteString("}\n")
				b.WriteString(fmt.Sprintf("} else if t.%s.Is%sLink(idx) {\n", rawMemberName, titleName))
				b.WriteString(fmt.Sprintf("handled, err = r.%s(t.%s.Get%sLink(idx))\n", dispatchFnName, rawMemberName, titleName))
				b.WriteString("if handled {\n")
				b.WriteString("s = Resolved\n")
				b.WriteString("}\n")
				if !hasMoreTypes {
					b.WriteString("} else {\n")
					b.WriteString("s = Resolved\n")
					b.WriteString("}\n")
					b.WriteString("return\n")
				} else {
					b.WriteString("} ")
					for _, elem := range p.Range {
						if elem.T != nil && (elem.T.Name == "Object" || elem.T.Name == "Link") {
							continue
						}
						b.WriteString(fmt.Sprintf("else if t.%s.Is%s%s(idx) {\n", rawMemberName, titleName, strings.Title(gen.Name(elem))))
						b.WriteString("s = RawResolutionNeeded\n")
						b.WriteString("}")
					}
					b.WriteString("\nreturn\n")
				}
				return b.String()
			},
		},
		{
			Name:    fmt.Sprintf("Has%s", titleName),
			Comment: fmt.Sprintf("Has%s returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.", titleName),
			P:       this,
			Args:    []*defs.FunctionVarDef{{"idx", "int"}},
			Return:  []*defs.FunctionVarDef{{"p", "Presence"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("p = NoPresence\n")
				b.WriteString(fmt.Sprintf("if t.%s.Is%sObject(idx) {\n", rawMemberName, titleName))
				b.WriteString("p = ConvenientPresence\n")
				b.WriteString(fmt.Sprintf("} else if t.%s.Is%sLink(idx) {\n", rawMemberName, titleName))
				b.WriteString("p = ConvenientPresence\n")
				if !hasMoreTypes {
					b.WriteString("}\n")
					b.WriteString("return\n")
				} else {
					b.WriteString("} ")
					for _, elem := range p.Range {
						if elem.T != nil && (elem.T.Name == "Object" || elem.T.Name == "Link") {
							continue
						}
						b.WriteString(fmt.Sprintf("else if t.%s.Is%s%s(idx) {\n", rawMemberName, titleName, strings.Title(gen.Name(elem))))
						b.WriteString("p = RawPresence\n")
						b.WriteString("}")
					}
					b.WriteString("\nreturn\n")
				}
				return b.String()
			},
		},
		{
			Name:    fmt.Sprintf("Add%s", titleName),
			Comment: fmt.Sprintf("Add%s adds an 'Object' typed value.", titleName),
			P:       this,
			Args:    []*defs.FunctionVarDef{{"i", "vocab.ObjectType"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString(fmt.Sprintf("t.%s.Add%sObject(i)\n", rawMemberName, titleName))
				return b.String()
			},
		},
		{
			Name:    fmt.Sprintf("Set%sLink", titleName),
			Comment: fmt.Sprintf("Set%sLink adds a 'Link' typed value.", titleName),
			P:       this,
			Args:    []*defs.FunctionVarDef{{"i", "vocab.LinkType"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString(fmt.Sprintf("t.%s.Add%sLink(i)\n", rawMemberName, titleName))
				return b.String()
			},
		},
	}...)
}

func generateNonFunctionalPropertyType(p *defs.PropertyType, this *defs.StructDef, ref *defs.Type, onlyType bool) {
	additionalTitleName := strings.Title(gen.Name(p.Range[0]))
	if !onlyType && defs.IsOnlyOtherPropertyBesidesIRI(0, p.Range) {
		additionalTitleName = ""
	}
	titleName := strings.Title(p.Name)
	this.F = append(this.F, []*defs.MemberFunctionDef{
		{
			Name:    fmt.Sprintf("Len%s", titleName),
			Comment: fmt.Sprintf("Len%s returns the number of values this property contains. Each index be used with Has%s to determine if Resolve%s is safe to call or if raw handling would be needed.", titleName, titleName, titleName, p.Name),
			P:       this,
			Return:  []*defs.FunctionVarDef{{"idx", "int"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString(fmt.Sprintf("return t.%s.%sLen()\n", rawMemberName, titleName))
				return b.String()
			},
		},
		{
			Name:    fmt.Sprintf("Resolve%s", titleName),
			Comment: fmt.Sprintf("Resolve%s passes the actual concrete type to the resolver for handing property %s. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.", titleName, p.Name),
			P:       this,
			Args:    []*defs.FunctionVarDef{{"r", "*" + resolverName}, {"idx", "int"}},
			Return:  []*defs.FunctionVarDef{{"s", "Resolution"}, {"err", "error"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("s = Unresolved\n")
				b.WriteString("handled := false\n")
				if onlyType {
					b.WriteString(fmt.Sprintf("if /*t.%s.Has%s(idx)*/ true {\n", rawMemberName, titleName))
					b.WriteString(fmt.Sprintf("handled, err = r.%s(t.%s.Get%s(idx))\n", dispatchFnName, rawMemberName, titleName))
				} else {
					b.WriteString(fmt.Sprintf("if t.%s.Is%s%s(idx) {\n", rawMemberName, titleName, additionalTitleName))
					b.WriteString(fmt.Sprintf("handled, err = r.%s(t.%s.Get%s%s(idx))\n", dispatchFnName, rawMemberName, titleName, additionalTitleName))
				}
				b.WriteString("if handled {\n")
				b.WriteString("s = Resolved\n")
				b.WriteString("}\n")
				b.WriteString("} ")
				for idx, elem := range p.Range {
					if idx == 0 {
						continue
					}
					b.WriteString(fmt.Sprintf("else if t.%s.Is%s%s(idx) {\n", rawMemberName, titleName, strings.Title(gen.Name(elem))))
					b.WriteString("s = RawResolutionNeeded\n")
					b.WriteString("}")
				}
				b.WriteString("\nreturn\n")
				return b.String()
			},
		},
		{
			Name:    fmt.Sprintf("Add%s", titleName),
			Comment: fmt.Sprintf("Add%s appends the value for property '%s'.", titleName, p.Name),
			P:       this,
			Args:    []*defs.FunctionVarDef{{"i", "vocab." + gen.InterfaceName(ref)}},
			Body: func() string {
				var b bytes.Buffer
				if onlyType {
					b.WriteString(fmt.Sprintf("t.%s.Add%s(i)\n", rawMemberName, titleName))
				} else {
					b.WriteString(fmt.Sprintf("t.%s.Add%s%s(i)\n", rawMemberName, titleName, additionalTitleName))
				}
				return b.String()
			},
		},
		{
			Name:    fmt.Sprintf("Remove%s", titleName),
			Comment: fmt.Sprintf("Remove%s deletes the value from the specified index for property '%s'.", titleName, p.Name),
			P:       this,
			Args:    []*defs.FunctionVarDef{{"idx", "int"}},
			Body: func() string {
				var b bytes.Buffer
				if onlyType {
					b.WriteString(fmt.Sprintf("t.%s.Remove%s(idx)\n", rawMemberName, titleName))
				} else {
					b.WriteString(fmt.Sprintf("t.%s.Remove%s%s(idx)\n", rawMemberName, titleName, additionalTitleName))
				}
				return b.String()
			},
		},
	}...)
	if !onlyType {
		this.F = append(this.F, []*defs.MemberFunctionDef{
			{
				Name:    fmt.Sprintf("Has%s", titleName),
				Comment: fmt.Sprintf("Has%s returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.", titleName),
				P:       this,
				Args:    []*defs.FunctionVarDef{{"idx", "int"}},
				Return:  []*defs.FunctionVarDef{{"p", "Presence"}},
				Body: func() string {
					var b bytes.Buffer
					b.WriteString("p = NoPresence\n")
					b.WriteString(fmt.Sprintf("if t.%s.Is%s%s(idx) {\n", rawMemberName, titleName, additionalTitleName))
					b.WriteString("p = ConvenientPresence\n")
					b.WriteString("} ")
					for idx, elem := range p.Range {
						if idx == 0 {
							continue
						}
						b.WriteString(fmt.Sprintf("else if t.%s.Is%s%s(idx) {\n", rawMemberName, titleName, strings.Title(gen.Name(elem))))
						b.WriteString("p = RawPresence\n")
						b.WriteString("}")
					}
					b.WriteString("\nreturn\n")
					return b.String()
				},
			},
		}...)
	}
}

func generateNonFunctionalPropertyValue(p *defs.PropertyType, this *defs.StructDef, value *defs.ValueType, onlyType bool) {
	additionalTitleName := strings.Title(gen.Name(p.Range[0]))
	if !onlyType && defs.IsOnlyOtherPropertyBesidesIRI(0, p.Range) {
		additionalTitleName = ""
	}
	kind := deref(value.DefinitionType)
	titleName := strings.Title(p.Name)
	this.F = append(this.F, []*defs.MemberFunctionDef{
		{
			Name:    fmt.Sprintf("Len%s", titleName),
			Comment: fmt.Sprintf("Len%s returns the number of values this property contains. Each index be used with Has%s to determine if Get%s is safe to call or if raw handling would be needed.", titleName, titleName, titleName, p.Name),
			P:       this,
			Return:  []*defs.FunctionVarDef{{"idx", "int"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString(fmt.Sprintf("return t.%s.%sLen()\n", rawMemberName, titleName))
				return b.String()
			},
		},
		{
			Name:    fmt.Sprintf("Get%s", titleName),
			Comment: fmt.Sprintf("Get%s attempts to get this '%s' property as a %s. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.", titleName, p.Name, kind),
			P:       this,
			Args:    []*defs.FunctionVarDef{{"idx", "int"}},
			Return:  []*defs.FunctionVarDef{{"r", "Resolution"}, {"k", kind}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("r = Unresolved\n")
				b.WriteString("handled := false\n")
				if onlyType {
					b.WriteString(fmt.Sprintf("if /*t.%s.Has%s(idx)*/ true {\n", rawMemberName, titleName))
					b.WriteString(fmt.Sprintf("k = t.%s.Get%s(idx)\n", rawMemberName, titleName))
				} else {
					b.WriteString(fmt.Sprintf("if t.%s.Is%s%s(idx) {\n", rawMemberName, titleName, additionalTitleName))
					b.WriteString(fmt.Sprintf("k = t.%s.Get%s%s(idx)\n", rawMemberName, titleName, additionalTitleName))
				}
				b.WriteString("if handled {\n")
				b.WriteString("r = Resolved\n")
				b.WriteString("}\n")
				b.WriteString("} ")
				for idx, elem := range p.Range {
					if idx == 0 {
						continue
					}
					b.WriteString(fmt.Sprintf("else if t.%s.Is%s%s(idx) {\n", rawMemberName, titleName, strings.Title(gen.Name(elem))))
					b.WriteString("r = RawResolutionNeeded\n")
					b.WriteString("}")
				}
				b.WriteString("\nreturn\n")
				return b.String()
			},
		},
		{
			Name:    fmt.Sprintf("Add%s", titleName),
			Comment: fmt.Sprintf("Add%s appends the value for property '%s'.", titleName, p.Name),
			P:       this,
			Args:    []*defs.FunctionVarDef{{"k", kind}},
			Body: func() string {
				var b bytes.Buffer
				if onlyType {
					b.WriteString(fmt.Sprintf("t.%s.Add%s(k)\n", rawMemberName, titleName))
				} else {
					b.WriteString(fmt.Sprintf("t.%s.Add%s%s(k)\n", rawMemberName, titleName, additionalTitleName))
				}
				return b.String()
			},
		},
		{
			Name:    fmt.Sprintf("Remove%s", titleName),
			Comment: fmt.Sprintf("Remove%s deletes the value from the specified index for property '%s'.", titleName, p.Name),
			P:       this,
			Args:    []*defs.FunctionVarDef{{"idx", "int"}},
			Body: func() string {
				var b bytes.Buffer
				if onlyType {
					b.WriteString(fmt.Sprintf("t.%s.Remove%s(idx)\n", rawMemberName, titleName))
				} else {
					b.WriteString(fmt.Sprintf("t.%s.Remove%s%s(idx)\n", rawMemberName, titleName, additionalTitleName))
				}
				return b.String()
			},
		},
	}...)
	if !onlyType {
		this.F = append(this.F, []*defs.MemberFunctionDef{
			{
				Name:    fmt.Sprintf("Has%s", titleName),
				Comment: fmt.Sprintf("Has%s returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.", titleName),
				P:       this,
				Args:    []*defs.FunctionVarDef{{"idx", "int"}},
				Return:  []*defs.FunctionVarDef{{"p", "Presence"}},
				Body: func() string {
					var b bytes.Buffer
					b.WriteString("p = NoPresence\n")
					b.WriteString(fmt.Sprintf("if t.%s.Is%s%s(idx) {\n", rawMemberName, titleName, additionalTitleName))
					b.WriteString("p = ConvenientPresence\n")
					b.WriteString("} ")
					for idx, elem := range p.Range {
						if idx == 0 {
							continue
						}
						b.WriteString(fmt.Sprintf("else if t.%s.Is%s%s(idx) {\n", rawMemberName, titleName, strings.Title(gen.Name(elem))))
						b.WriteString("p = RawPresence\n")
						b.WriteString("}")
					}
					b.WriteString("\nreturn\n")
					return b.String()
				},
			},
		}...)
	}
}

func generateNonFunctionalPropertyAny(p *defs.PropertyType, this *defs.StructDef) {
	titleName := strings.Title(p.Name)
	this.F = append(this.F, []*defs.MemberFunctionDef{
		{
			Name:    fmt.Sprintf("Len%s", titleName),
			Comment: fmt.Sprintf("Len%s returns the number of values this property contains. Each index be used with Has%s to determine if Get%s is safe to call or if raw handling would be needed.", titleName, titleName, titleName, p.Name),
			P:       this,
			Return:  []*defs.FunctionVarDef{{"idx", "int"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString(fmt.Sprintf("return t.%s.%sLen()\n", rawMemberName, titleName))
				return b.String()
			},
		},
		{
			Name:    fmt.Sprintf("Get%s", titleName),
			Comment: fmt.Sprintf("Get%s attempts to get this '%s' property as a string. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.", titleName, p.Name),
			P:       this,
			Args:    []*defs.FunctionVarDef{{"idx", "int"}},
			Return:  []*defs.FunctionVarDef{{"r", "Resolution"}, {"s", "string"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("r = Unresolved\n")
				b.WriteString(fmt.Sprintf("if tmp := t.%s.Get%s(idx); tmp != nil {\n", rawMemberName, titleName))
				b.WriteString("ok := false\n")
				b.WriteString("if s, ok = tmp.(string); ok {\n")
				b.WriteString("r = Resolved\n")
				b.WriteString("} else {\n")
				b.WriteString("r = RawResolutionNeeded\n")
				b.WriteString("}\n")
				b.WriteString("}\n")
				b.WriteString("return\n")
				return b.String()
			},
		},
		{
			Name:    fmt.Sprintf("Add%s", titleName),
			Comment: fmt.Sprintf("Add%s appends the value for property '%s'.", titleName, p.Name),
			P:       this,
			Args:    []*defs.FunctionVarDef{{"i", "interface{}"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString(fmt.Sprintf("t.%s.Add%s(i)\n", rawMemberName, titleName))
				return b.String()
			},
		},
		{
			Name:    fmt.Sprintf("Remove%s", titleName),
			Comment: fmt.Sprintf("Remove%s deletes the value from the specified index for property '%s'.", titleName, p.Name),
			P:       this,
			Args:    []*defs.FunctionVarDef{{"idx", "int"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString(fmt.Sprintf("t.%s.Remove%s(idx)\n", rawMemberName, titleName))
				return b.String()
			},
		},
	}...)
}

func rangeHasObjectAndLink(r []defs.RangeReference) bool {
	hasObj := false
	hasLink := false
	for _, ref := range r {
		if ref.T != nil {
			if ref.T.Name == "Object" {
				hasObj = true
			} else if ref.T.Name == "Link" {
				hasLink = true
			}
		}
	}
	return hasObj && hasLink
}

func rangeHasMoreThanJustObjectAndLink(r []defs.RangeReference) bool {
	for _, ref := range r {
		if ref.Any {
			panic("Cannot have Any and types in range")
		}
		if ref.V != nil {
			return true
		} else if ref.T != nil && ref.T.Name != "Object" && ref.T.Name != "Link" {
			return true
		}
	}
	return false
}

func deref(k string) string {
	if k[0] == '*' {
		return k[1:]
	}
	return k
}

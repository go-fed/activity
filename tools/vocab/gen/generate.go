// Package gen contains the libraries and algorithms used to generate the
// code for the vocab package.
package gen

import (
	"bytes"
	"fmt"
	"github.com/go-fed/activity/tools/defs"
	"go/format"
	"strings"
)

const (
	objectName                    = "Object"
	linkName                      = "Link"
	ObjectTypeName                = "ObjectType"
	LinkTypeName                  = "LinkType"
	typePropertyName              = "type"
	resolveLinkName               = "resolveLink"
	resolveObjectName             = "resolveObject"
	unknownValueDeserializeFnName = "unknownValueDeserialize"
	unknownValueSerializeFnName   = "unknownValueSerialize"
)

func GenerateImplementations(types []*defs.Type, properties []*defs.PropertyType, values []*defs.ValueType) ([]byte, error) {
	// Validate inputs
	err := validateDomains(properties)
	if err != nil {
		return nil, err
	}
	err = validateProperties(types)
	if err != nil {
		return nil, err
	}
	err = validateValues(values, properties)
	if err != nil {
		return nil, err
	}

	p := generatePackageDefinition()
	p.Defs = append(p.Defs, generateUnknownType())

	// Add ValueType serialize & deserialize functions
	for _, v := range values {
		p.F = append(p.F, v.DeserializeFn, v.SerializeFn)
	}
	p.F = append(p.F, defs.IRIFuncs()...)
	p.F = append(p.F, generateHasTypeFuncs(types)...)
	p.I = append(p.I, generateTyperInterface())

	// Add functions to resolve string 'name' into concrete types
	p.F = append(p.F, generateResolveObjectFunction(types))
	p.F = append(p.F, generateResolveLinkFunction(types))
	unknown := generateUnknownValueType()
	p.F = append(p.F, unknown.DeserializeFn, unknown.SerializeFn)

	for _, t := range types {
		funcs, defs, interfaces := generateDefinitions(t)
		p.F = append(p.F, funcs...)
		p.Defs = append(p.Defs, defs...)
		p.I = append(p.I, interfaces...)
	}
	return format.Source([]byte(p.Generate()))
}

func generateTyperInterface() *defs.InterfaceDef {
	return &defs.InterfaceDef{
		Typename: "Typer",
		Comment:  "Typer supports common functions for determining an ActivityStream type",
		F: []*defs.FunctionDef{
			{
				Name:   "TypeLen",
				Return: []*defs.FunctionVarDef{{Name: "l", Type: "int"}},
			},
			{
				Name:   "GetType",
				Args:   []*defs.FunctionVarDef{{Name: "index", Type: "int"}},
				Return: []*defs.FunctionVarDef{{Name: "v", Type: "interface{}"}},
			},
		},
	}
}

func generateHasTypeFuncs(types []*defs.Type) (f []*defs.FunctionDef) {
	var activityTypes []string
	for _, t := range types {
		t := t
		if defs.IsActivity(t) {
			activityTypes = append(activityTypes, t.Name)
		}
		f = append(f, &defs.FunctionDef{
			Name:    fmt.Sprintf("HasType%s", t.Name),
			Comment: fmt.Sprintf("HasType%s returns true if the Typer has a type of %s.", t.Name, t.Name),
			Args:    []*defs.FunctionVarDef{{Name: "t", Type: "Typer"}},
			Return:  []*defs.FunctionVarDef{{Name: "b", Type: "bool"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("for i := 0; i < t.TypeLen(); i++ {\n")
				b.WriteString("v := t.GetType(i)\n")
				b.WriteString("if s, ok := v.(string); ok {\n")
				b.WriteString(fmt.Sprintf("if s == \"%s\" {\n", t.Name))
				b.WriteString("return true\n")
				b.WriteString("}\n")
				b.WriteString("}\n")
				b.WriteString("}\n")
				b.WriteString("return false\n")
				return b.String()
			},
		})
	}
	f = append(f, &defs.FunctionDef{
		Name:    "IsActivityType",
		Comment: "Returns true if the provided Typer is an Activity.",
		Args:    []*defs.FunctionVarDef{{Name: "t", Type: "Typer"}},
		Return:  []*defs.FunctionVarDef{{Name: "b", Type: "bool"}},
		Body: func() string {
			var b bytes.Buffer
			b.WriteString(fmt.Sprintf("var activityTypes = []string{\"%s\"}\n", strings.Join(activityTypes, "\", \"")))
			b.WriteString("hasType := make(map[string]bool, 1)\n")
			b.WriteString("for i := 0; i < t.TypeLen(); i++ {\n")
			b.WriteString("v := t.GetType(i)\n")
			b.WriteString("if s, ok := v.(string); ok {\n")
			b.WriteString("hasType[s] = true\n")
			b.WriteString("}\n")
			b.WriteString("}\n")
			b.WriteString("for _, t := range activityTypes {\n")
			b.WriteString("if hasType[t] {\n")
			b.WriteString("return true\n")
			b.WriteString("}\n")
			b.WriteString("}\n")
			b.WriteString("return false\n")
			return b.String()
		},
	})
	return
}

func generatePackageDefinition() *defs.PackageDef {
	return &defs.PackageDef{
		Name:    "vocab",
		Comment: "Package vocab provides an implementation of serializing and deserializing activity streams into native golang structs without relying on reflection. This package is code-generated from the vocabulary specification available at https://www.w3.org/TR/activitystreams-vocabulary and by design forgoes full resolution of raw JSON-LD data. However, custom extensions of the vocabulary are supported by modifying the data definitions in the generation tool and rerunning it. Do not modify this package directly.",
		Imports: []string{"fmt", "time", "net/url", "regexp", "strconv", "math"},
		I: []*defs.InterfaceDef{
			{
				Typename: "Serializer",
				Comment:  "Serializer implementations can serialize themselves to a generic map form.",
				F: []*defs.FunctionDef{{
					Name:   "Serialize",
					Return: []*defs.FunctionVarDef{{"m", "map[string]interface{}"}, {"e", "error"}},
				}},
			},
			{
				Typename: "Deserializer",
				Comment:  "Deserializer implementations can deserialize themselves from a generic map form.",
				F: []*defs.FunctionDef{{
					Name:   "Deserialize",
					Args:   []*defs.FunctionVarDef{{"m", "map[string]interface{}"}},
					Return: []*defs.FunctionVarDef{{"e", "error"}},
				}},
			},
		},
	}
}

func generateUnknownValueType() *defs.ValueType {
	return &defs.ValueType{
		Name:           "Unknown Value",
		DefinitionType: "interface{}",
		Zero:           "nil",
		DeserializeFn: &defs.FunctionDef{
			Name:    unknownValueDeserializeFnName,
			Comment: "unknownValueDeserialize transparently stores the object.",
			Args:    []*defs.FunctionVarDef{{"v", "interface{}"}},
			Return:  []*defs.FunctionVarDef{{"o", "interface{}"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("o = v\n")
				b.WriteString("return\n")
				return b.String()
			},
		},
		SerializeFn: &defs.FunctionDef{
			Name:    unknownValueSerializeFnName,
			Comment: "unknownValueSerialize transparently returns the object.",
			Args:    []*defs.FunctionVarDef{{"v", "interface{}"}},
			Return:  []*defs.FunctionVarDef{{"o", "interface{}"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("o = v\n")
				b.WriteString("return\n")
				return b.String()
			},
		},
	}
}

func generateUnknownType() *defs.StructDef {
	u := &defs.StructDef{
		Typename: "Unknown",
		Comment:  "Unknown is an entry whose root type is unknown.",
		M: []*defs.StructMember{{
			Name:    "u",
			Type:    "map[string]interface{}",
			Comment: "Raw unknown, untyped values",
		}},
	}
	u.F = append(u.F, []*defs.MemberFunctionDef{
		{
			Name:    "Serialize",
			Comment: "Serialize turns this object into a map[string]interface{}. Note that for the Unknown type, the \"type\" property is NOT populated with anything special during this process.",
			P:       u,
			Return:  []*defs.FunctionVarDef{{"m", "map[string]interface{}"}, {"err", "error"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("m = t.u\n")
				b.WriteString("return\n")
				return b.String()
			},
		},
		{
			Name:    "Deserialize",
			Comment: "Deserialize populates this object from a map[string]interface{}",
			P:       u,
			Args:    []*defs.FunctionVarDef{{"m", "map[string]interface{}"}},
			Return:  []*defs.FunctionVarDef{{"err", "error"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("t.u = m\n")
				b.WriteString("return\n")
				return b.String()
			},
		},
		{
			Name:    "HasField",
			Comment: "HasField determines whether the call to GetField is safe with the specified field",
			P:       u,
			Args:    []*defs.FunctionVarDef{{"f", "string"}},
			Return:  []*defs.FunctionVarDef{{"ok", "bool"}},
			Body: func() string {
				return "return t.u != nil && t.u[f] != nil\n"
			},
		},
		{
			Name:    "GetField",
			Comment: "GetField returns the unknown field value",
			P:       u,
			Args:    []*defs.FunctionVarDef{{"f", "string"}},
			Return:  []*defs.FunctionVarDef{{"v", "interface{}"}},
			Body: func() string {
				return "return t.u[f]"
			},
		},
		{
			Name:    "SetField",
			Comment: "SetField sets the unknown field value",
			P:       u,
			Args:    []*defs.FunctionVarDef{{"f", "string"}, {"i", "interface{}"}},
			Return:  []*defs.FunctionVarDef{{"this", "*" + u.Typename}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("if t.u == nil {\n")
				b.WriteString("t.u = make(map[string]interface{})\n")
				b.WriteString("}\n")
				b.WriteString("t.u[f] = i\n")
				b.WriteString("return t\n")
				return b.String()
			},
		},
	}...)
	return u
}

func generateResolveObjectFunction(types []*defs.Type) *defs.FunctionDef {
	return &defs.FunctionDef{
		Name:    resolveObjectName,
		Comment: fmt.Sprintf("%s turns a string type that extends Object into a concrete type.", resolveObjectName),
		Args:    []*defs.FunctionVarDef{{"s", "string"}},
		Return:  []*defs.FunctionVarDef{{"i", "interface{}"}},
		Body: func() string {
			var b bytes.Buffer
			for _, r := range types {
				if isAnObjectType(r) {
					b.WriteString(fmt.Sprintf("if s == \"%s\" {\n", r.Name))
					b.WriteString(fmt.Sprintf("return &%s{}\n", r.Name))
					b.WriteString("}\n")
				}
			}
			b.WriteString("return nil\n")
			return b.String()
		},
	}
}

func generateResolveLinkFunction(types []*defs.Type) *defs.FunctionDef {
	return &defs.FunctionDef{
		Name:    resolveLinkName,
		Comment: fmt.Sprintf("%s turns a string type that extends Link into a concrete type.", resolveLinkName),
		Args:    []*defs.FunctionVarDef{{"s", "string"}},
		Return:  []*defs.FunctionVarDef{{"i", "interface{}"}},
		Body: func() string {
			var b bytes.Buffer
			for _, r := range types {
				if isALinkType(r) {
					b.WriteString(fmt.Sprintf("if s == \"%s\" {\n", r.Name))
					b.WriteString(fmt.Sprintf("return &%s{}\n", r.Name))
					b.WriteString("}\n")
				}
			}
			b.WriteString("return nil\n")
			return b.String()
		},
	}
}

func isAnObjectType(t *defs.Type) bool {
	obj, _ := getIsAType(t)
	return obj
}

func isALinkType(t *defs.Type) bool {
	_, link := getIsAType(t)
	return link
}

func getIsAType(t *defs.Type) (obj, link bool) {
	if t.Name == objectName {
		return true, false
	} else if t.Name == linkName {
		return false, true
	}
	for _, e := range t.Extends {
		if obj, link = getIsAType(e); obj || link {
			return
		}
	}
	panic("Unknown root is-a-type")
}

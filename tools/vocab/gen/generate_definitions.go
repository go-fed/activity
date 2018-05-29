package gen

import (
	"bytes"
	"fmt"
	"github.com/go-fed/activity/tools/defs"
)

func generateDefinitions(t *defs.Type) (fd []*defs.FunctionDef, sd []*defs.StructDef, x []*defs.InterfaceDef) {
	this := &defs.StructDef{
		Typename: t.Name,
		Comment:  t.Notes,
		M:        []*defs.StructMember{{"unknown_", "map[string]interface{}", "An unknown value."}},
	}
	sd = append(sd, this)
	thisInterface := &defs.InterfaceDef{
		Typename: InterfaceName(t),
		Comment:  fmt.Sprintf("%s is an interface for accepting types that extend from '%s'.", InterfaceName(t), t.Name),
		O:        []string{"Serializer", "Deserializer"},
	}
	x = append(x, thisInterface)
	var serializeFragments []string
	var deserializeFragments []string
	for _, p := range t.GetProperties() {
		pf, pd, s, d := generatePropertyDefinitions(p, this, thisInterface)
		serializeFragments = append(serializeFragments, s)
		deserializeFragments = append(deserializeFragments, d)
		fd = append(fd, pf...)
		sd = append(sd, pd...)
	}
	generateWithoutProperties(t, this, thisInterface)
	generateAddUnknownFunction(t, this)
	generateHasUnknownFunction(t, this)
	generateRemoveUnknownFunction(t, this)
	generateSerializeFunction(t, this, serializeFragments)
	generateDeserializeFunction(t, this, deserializeFragments)
	return
}

func Name(r defs.RangeReference) string {
	if r.T != nil {
		return r.T.Name
	} else if r.V != nil {
		return r.V.Name
	} else if r.Any {
		return "any"
	} else {
		panic("RangeReference must have exactly one member set")
	}
}

func Type(r defs.RangeReference) string {
	if r.T != nil {
		return InterfaceName(r.T)
	} else if r.V != nil {
		return r.V.DefinitionType
	} else if r.Any {
		return "interface{}"
	} else {
		panic("RangeReference must have exactly one member set")
	}
}

func Deserialize(r defs.RangeReference, mapName, field string, slice bool) string {
	if r.T != nil {
		var parseCode bytes.Buffer
		parseCode.WriteString(fmt.Sprintf("tmp := &%s{}\n", r.T.Name))
		parseCode.WriteString("err = tmp.Deserialize(v)\n")
		parseCode.WriteString("if err != nil {\nreturn err\n}\n")
		var b bytes.Buffer
		b.WriteString("// Begin generation by RangeReference.Deserialize for Type\n")
		deserializeCode(&b, parseCode.String(), mapName, "map[string]interface{}", field, slice, false)
		b.WriteString("// End generation by RangeReference.Deserialize for Type\n")
		return b.String()
	} else if r.V != nil {
		var b bytes.Buffer
		b.WriteString("// Begin generation by RangeReference.Deserialize for Value\n")
		deserializeCode(&b, fmt.Sprintf("tmp, err := %s(v)\n", r.V.DeserializeFn.Name), mapName, "interface{}", field, slice, false)
		b.WriteString("// End generation by RangeReference.Deserialize for Value\n")
		return b.String()
	} else {
		panic("Bad call to Deserialize for Any")
	}
}

func Serialize(r defs.RangeReference, mapName, field string, slice bool) string {
	if r.T != nil {
		var sCode bytes.Buffer
		sCode.WriteString("tmp, err := v.Serialize()\n")
		sCode.WriteString("if err != nil {\nreturn m, err\n}\n")
		var b bytes.Buffer
		b.WriteString("// Begin generation by RangeReference.Serialize for Type\n")
		serializeCode(&b, sCode.String(), mapName, field, slice)
		b.WriteString("// End generation by RangeReference.Serialize for Type\n")
		return b.String()
	} else if r.V != nil {
		deref := "*"
		if slice || isIRIType(Type(r)) {
			deref = ""
		}
		var b bytes.Buffer
		b.WriteString("// Begin generation by RangeReference.Serialize for Value\n")
		serializeCode(&b, fmt.Sprintf("tmp := %s(%sv)\n", r.V.SerializeFn.Name, deref), mapName, field, slice)
		b.WriteString("// End generation by RangeReference.Serialize for Value\n")
		return b.String()
	} else {
		panic("Bad call to Serialize for Any")
	}
}

func generateRemoveUnknownFunction(t *defs.Type, this *defs.StructDef) {
	m := &defs.MemberFunctionDef{
		Name:    "RemoveUnknown",
		Comment: "RemoveUnknown removes a raw extension from this object with the specified key",
		P:       this,
		Args:    []*defs.FunctionVarDef{{"k", "string"}},
		Return:  []*defs.FunctionVarDef{{"this", "*" + t.Name}},
		Body: func() string {
			var b bytes.Buffer
			b.WriteString("delete(t.unknown_, k)\n")
			b.WriteString("return t\n")
			return b.String()
		},
	}
	this.F = append(this.F, m)
}

func generateHasUnknownFunction(t *defs.Type, this *defs.StructDef) {
	m := &defs.MemberFunctionDef{
		Name:    "HasUnknown",
		Comment: "HasUnknown returns true if there is an unknown object with the specified key",
		P:       this,
		Args:    []*defs.FunctionVarDef{{"k", "string"}},
		Return:  []*defs.FunctionVarDef{{"b", "bool"}},
		Body: func() string {
			var b bytes.Buffer
			b.WriteString("if t.unknown_ == nil {\n")
			b.WriteString("return false")
			b.WriteString("}\n")
			b.WriteString("_, ok := t.unknown_[k]\n")
			b.WriteString("return ok\n")
			return b.String()
		},
	}
	this.F = append(this.F, m)
}

func generateAddUnknownFunction(t *defs.Type, this *defs.StructDef) {
	m := &defs.MemberFunctionDef{
		Name:    "AddUnknown",
		Comment: "AddUnknown adds a raw extension to this object with the specified key",
		P:       this,
		Args:    []*defs.FunctionVarDef{{"k", "string"}, {"i", "interface{}"}},
		Return:  []*defs.FunctionVarDef{{"this", "*" + t.Name}},
		Body: func() string {
			var b bytes.Buffer
			b.WriteString("if t.unknown_ == nil {\n")
			b.WriteString("t.unknown_ = make(map[string]interface{})\n")
			b.WriteString("}\n")
			b.WriteString("t.unknown_[k] = i\n")
			b.WriteString("return t\n")
			return b.String()
		},
	}
	this.F = append(this.F, m)
}

func generateDeserializeFunction(t *defs.Type, this *defs.StructDef, fragments []string) {
	d := &defs.MemberFunctionDef{
		Name:    "Deserialize",
		Comment: "Deserialize populates this object from a map[string]interface{}",
		P:       this,
		Args:    []*defs.FunctionVarDef{{"m", "map[string]interface{}"}},
		Return:  []*defs.FunctionVarDef{{"err", "error"}},
		Body: func() string {
			var b bytes.Buffer
			b.WriteString("for k, v := range m {\n")
			b.WriteString("handled := false\n")
			for _, s := range fragments {
				b.WriteString("if !handled {\n")
				b.WriteString(s)
				b.WriteString("}\n")
			}
			b.WriteString("if !handled && k != \"@context\"{\n")
			b.WriteString("if t.unknown_ == nil {\n")
			b.WriteString("t.unknown_ = make(map[string]interface{})\n")
			b.WriteString("}\n")
			b.WriteString(fmt.Sprintf("t.unknown_[k] = %s(v)\n", unknownValueDeserializeFnName))
			b.WriteString("}\n")
			b.WriteString("}\n")
			b.WriteString("return\n")
			return b.String()
		},
	}
	this.F = append(this.F, d)
}

func generateSerializeFunction(t *defs.Type, this *defs.StructDef, fragments []string) {
	d := &defs.MemberFunctionDef{
		Name:    "Serialize",
		Comment: fmt.Sprintf("Serialize turns this object into a map[string]interface{}. Note that the \"type\" property will automatically be populated with \"%s\" if not manually set by the caller", t.Name),
		P:       this,
		Return:  []*defs.FunctionVarDef{{"m", "map[string]interface{}"}, {"err", "error"}},
		Body: func() string {
			var b bytes.Buffer
			b.WriteString("m = make(map[string]interface{})\n")
			b.WriteString("for k, v := range t.unknown_ {\n")
			b.WriteString(fmt.Sprintf("m[k] = %s(v)\n", unknownValueSerializeFnName))
			b.WriteString("}\n")
			b.WriteString("var typeAlreadySet bool\n")
			b.WriteString(fmt.Sprintf("for _, k := range t.%s {\n", cleanName(typePropertyName)))
			b.WriteString("if ks, ok := k.(string); ok {\n")
			b.WriteString(fmt.Sprintf("if ks == \"%s\" {\n", t.Name))
			b.WriteString("typeAlreadySet = true\n")
			b.WriteString("break\n")
			b.WriteString("}\n")
			b.WriteString("}\n")
			b.WriteString("}\n")
			b.WriteString("if !typeAlreadySet {\n")
			b.WriteString(fmt.Sprintf("t.%s = append(t.%s, \"%s\")\n", cleanName(typePropertyName), cleanName(typePropertyName), t.Name))
			b.WriteString("}\n")
			for _, s := range fragments {
				b.WriteString(s)
			}
			b.WriteString("return\n")
			return b.String()
		},
	}
	this.F = append(this.F, d)
}

func generateWithoutProperties(d *defs.Type, this *defs.StructDef, it *defs.InterfaceDef) {
	hasNamed := make(map[string]bool, 0)
	for _, p := range d.GetProperties() {
		hasNamed[p.Name] = true
	}
	for _, t := range d.WithoutProperties {
		if hasNamed[t.Name] {
			// No need to stub -- already has a replacement.
			continue
		}
		clonedThis := &defs.StructDef{
			Typename: this.Typename,
			Comment:  this.Comment,
		}
		_, _, _, _ = generatePropertyDefinitions(t, clonedThis, it)
		for _, f := range clonedThis.F {
			if len(f.Return) > 1 {
				panic("generateWithoutProperties can only handle replacing functions with zero or one return types")
			}
			f.P = this
			f.Comment = fmt.Sprintf("%s is NOT a valid property for this type; calling its associated methods will always yield an empty-equivalent value such as false, nil, or empty string. This includes instances where it should return itself. This ugliness is a symptom of the fundamental design of the ActivityStream vocabulary as instead of 'W-is-a-X' relationships it contains the notion of 'W-is-a-X-except-for-Y'.", f.Name)
			if len(f.Return) > 0 {
				returnType := f.Return[0].Type
				f.Body = func() string {
					var b bytes.Buffer
					if returnType == "url.URL" {
						b.WriteString("return url.URL{}\n")
					} else if isSlice(returnType) || isPtrType(returnType) {
						b.WriteString("return nil\n")
					} else if returnType == "bool" {
						b.WriteString("return false\n")
					} else if returnType == "int" {
						b.WriteString("return 0\n")
					} else if returnType == "interface{}" {
						b.WriteString("return nil\n")
					} else {
						found := false
						for _, v := range t.Range {
							if v.V == nil {
								continue
							}
							if returnType == deref(Type(v)) {
								found = true
								b.WriteString(fmt.Sprintf("return %s\n", v.V.Zero))
							}
						}
						if !found {
							b.WriteString("return nil\n")
						}
					}
					return b.String()
				}
				this.F = append(this.F, f)
			} else {
				f.Body = func() string { return "\n" }
				this.F = append(this.F, f)
			}
		}
	}
}

func InterfaceName(t *defs.Type) string {
	return fmt.Sprintf("%sType", t.Name)
}

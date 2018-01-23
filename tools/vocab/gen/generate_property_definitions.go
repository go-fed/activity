package gen

import (
	"bytes"
	"fmt"
	"github.com/go-fed/activity/tools/defs"
	"strings"
)

func generatePropertyDefinitions(t *defs.PropertyType, this *defs.StructDef, i *defs.InterfaceDef) (fd []*defs.FunctionDef, sd []*defs.StructDef, s, d string) {
	if isAny(t) {
		fd, sd, s, d = generateAnyDefinitions(t, this, i)
	} else if isSingleType(t) {
		fd, sd, s, d = generateSingleTypeDefinitions(t, this, i)
	} else {
		fd, sd, s, d = generateMultiTypeDefinitions(t, this, i)
	}
	if t.NaturalLanguageMap {
		fdMap, sdMap, sMap, dMap := generateNaturalLanguageMap(t, this, i)
		fd = append(fd, fdMap...)
		sd = append(sd, sdMap...)
		s += "\n" + sMap
		d += "\n" + dMap
	}
	return
}

func generateNaturalLanguageMap(t *defs.PropertyType, this *defs.StructDef, i *defs.InterfaceDef) (fd []*defs.FunctionDef, sd []*defs.StructDef, s, d string) {
	mapMember := &defs.StructMember{
		Name:    t.Name + "Map",
		Type:    "map[string]string",
		Comment: fmt.Sprintf("The '%sMap' value holds language-specific values for property '%s'"),
	}
	this.M = append(this.M, mapMember)
	this.F = append(this.F, []*defs.MemberFunctionDef{
		{
			Name:    fmt.Sprintf("%sMapLanguages", strings.Title(t.Name)),
			Comment: fmt.Sprintf("%sMapLanguages returns all languages for this property's language mapping, or nil if there are none.", strings.Title(t.Name)),
			P:       this,
			Return:  []*defs.FunctionVarDef{{"l", "[]string"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString(fmt.Sprintf("if t.%sMap == nil || len(t.%sMap) == 0 {\n", t.Name, t.Name))
				b.WriteString("return nil\n")
				b.WriteString("}\n")
				b.WriteString(fmt.Sprintf("for k := range t.%sMap {\n", t.Name))
				b.WriteString("l = append(l, k)\n")
				b.WriteString("}\n")
				b.WriteString("return\n")
				return b.String()
			},
		},
		{
			Name:    fmt.Sprintf("Get%sMap", strings.Title(t.Name)),
			Comment: fmt.Sprintf("Get%sMap retrieves the value of the property for the specified language, or an empty string if it does not exist", strings.Title(t.Name)),
			P:       this,
			Args:    []*defs.FunctionVarDef{{"l", "string"}},
			Return:  []*defs.FunctionVarDef{{"v", "string"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString(fmt.Sprintf("if t.%sMap == nil {\n", t.Name))
				b.WriteString("return \"\"\n")
				b.WriteString("}\n")
				b.WriteString("ok := false\n")
				b.WriteString(fmt.Sprintf("v, ok = t.%sMap[l]\n", t.Name))
				b.WriteString("if !ok {\n")
				b.WriteString("return \"\"\n")
				b.WriteString("}\n")
				b.WriteString("return v\n")
				return b.String()
			},
		},
		{
			Name:    fmt.Sprintf("Set%sMap", strings.Title(t.Name)),
			Comment: fmt.Sprintf("Set$sMap sets the value of the property for the specified language", strings.Title(t.Name)),
			P:       this,
			Args:    []*defs.FunctionVarDef{{"l", "string"}, {"v", "string"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString(fmt.Sprintf("if t.%sMap == nil {\n", t.Name))
				b.WriteString(fmt.Sprintf("t.%sMap = make(map[string]string)\n", t.Name))
				b.WriteString("}\n")
				b.WriteString(fmt.Sprintf("t.%sMap[l] = v\n", t.Name))
				return b.String()
			},
		},
	}...)
	i.F = append(i.F, []*defs.FunctionDef{
		{
			Name:    fmt.Sprintf("%sMapLanguages", strings.Title(t.Name)),
			Comment: fmt.Sprintf("%sMapLanguages returns all languages for this property's language mapping, or nil if there are none.", strings.Title(t.Name)),
			Return:  []*defs.FunctionVarDef{{"l", "[]string"}},
		},
		{
			Name:    fmt.Sprintf("Get%sMap", strings.Title(t.Name)),
			Comment: fmt.Sprintf("Get%sMap retrieves the value of the property for the specified language, or an empty string if it does not exist", strings.Title(t.Name)),
			Args:    []*defs.FunctionVarDef{{"l", "string"}},
			Return:  []*defs.FunctionVarDef{{"v", "string"}},
		},
		{
			Name:    fmt.Sprintf("Set%sMap", strings.Title(t.Name)),
			Comment: fmt.Sprintf("Set$sMap sets the value of the property for the specified language", strings.Title(t.Name)),
			Args:    []*defs.FunctionVarDef{{"l", "string"}, {"v", "string"}},
		},
	}...)
	var b bytes.Buffer
	b.WriteString("// Begin generation by generateNaturalLanguageMap\n")
	b.WriteString(fmt.Sprintf("if k == \"%sMap\" {\n", t.Name))
	b.WriteString("if vMap, ok := v.(map[string]interface{}); ok {\n")
	b.WriteString("val := make(map[string]string)\n")
	b.WriteString("for k, iVal := range vMap {\n")
	b.WriteString("if sVal, ok := iVal.(string); ok {\n")
	b.WriteString("val[k] = sVal\n")
	b.WriteString("}\n")
	b.WriteString("}\n")
	b.WriteString(fmt.Sprintf("t.%sMap = val\n", t.Name))
	b.WriteString("handled = true\n")
	b.WriteString("}\n")
	b.WriteString("}\n")
	b.WriteString("// End generation by generateNaturalLanguageMap\n")
	d = b.String()
	var bs bytes.Buffer
	bs.WriteString("// Begin generation by generateNaturalLanguageMap\n")
	bs.WriteString(fmt.Sprintf("if t.%sMap != nil && len(t.%sMap) >= 0 {\n", t.Name, t.Name))
	bs.WriteString(fmt.Sprintf("m[\"%sMap\"] = t.%sMap\n", t.Name, t.Name))
	bs.WriteString("}\n")
	bs.WriteString("// End generation by generateNaturalLanguageMap\n")
	s = bs.String()
	return
}

func generateAnyDefinitions(t *defs.PropertyType, this *defs.StructDef, i *defs.InterfaceDef) (fd []*defs.FunctionDef, sd []*defs.StructDef, s, d string) {
	if t.Functional {
		return generateFunctionalAnyDefinition(t, this, i)
	} else {
		return generateNonFunctionalAnyDefinition(t, this, i)
	}
}

func generateFunctionalAnyDefinition(t *defs.PropertyType, this *defs.StructDef, i *defs.InterfaceDef) (fd []*defs.FunctionDef, sd []*defs.StructDef, s, d string) {
	titleName := strings.Title(t.Name)
	anyMember := &defs.StructMember{
		Name:    cleanName(t.Name),
		Type:    "interface{}",
		Comment: fmt.Sprintf("The functional '%s' value can hold any type, but only a single value", t.Name),
	}
	this.M = append(this.M, anyMember)
	this.F = append(this.F, &defs.MemberFunctionDef{
		Name:    fmt.Sprintf("Get%s", titleName),
		Comment: fmt.Sprintf("Get%s returns the value for %s", titleName, t.Name),
		P:       this,
		Return:  []*defs.FunctionVarDef{{"v", anyMember.Type}},
		Body: func() string {
			var b bytes.Buffer
			b.WriteString(fmt.Sprintf("return t.%s\n", anyMember.Name))
			return b.String()
		},
	}, &defs.MemberFunctionDef{
		Name:    fmt.Sprintf("Set%s", titleName),
		Comment: fmt.Sprintf("Set%s sets the value of %s", titleName, t.Name),
		P:       this,
		Args:    []*defs.FunctionVarDef{{"v", anyMember.Type}},
		Body: func() string {
			var b bytes.Buffer
			b.WriteString(fmt.Sprintf("t.%s = v\n", anyMember.Name))
			return b.String()
		},
	})
	i.F = append(i.F, []*defs.FunctionDef{
		{
			Name:    fmt.Sprintf("Get%s", titleName),
			Comment: fmt.Sprintf("Get%s returns the value for %s", titleName, t.Name),
			Return:  []*defs.FunctionVarDef{{"v", anyMember.Type}},
		},
		{
			Name:    fmt.Sprintf("Set%s", titleName),
			Comment: fmt.Sprintf("Set%s sets the value of %s", titleName, t.Name),
			Args:    []*defs.FunctionVarDef{{"v", anyMember.Type}},
		},
	}...)
	var b bytes.Buffer
	b.WriteString("// Begin generation by generateFunctionalAnyDefinition\n")
	b.WriteString(fmt.Sprintf("if k == \"%s\" {\n", t.Name))
	b.WriteString(fmt.Sprintf("t.%s = v\n", anyMember.Name))
	b.WriteString("handled = true\n")
	b.WriteString("}\n")
	b.WriteString("// End generation by generateFunctionalAnyDefinition\n")
	d = b.String()
	var bs bytes.Buffer
	bs.WriteString("// Begin generation by generateFunctionalAnyDefinition\n")
	bs.WriteString(fmt.Sprintf("if t.%s != nil {\n", anyMember.Name))
	bs.WriteString(fmt.Sprintf("m[\"%s\"] = t.%s\n", t.Name, anyMember.Name))
	bs.WriteString("}\n")
	bs.WriteString("// End generation by generateFunctionalAnyDefinition\n")
	s = bs.String()
	return
}

func generateNonFunctionalAnyDefinition(t *defs.PropertyType, this *defs.StructDef, i *defs.InterfaceDef) (fd []*defs.FunctionDef, sd []*defs.StructDef, s, d string) {
	titleName := strings.Title(t.Name)
	anyMember := &defs.StructMember{
		Name:    cleanName(t.Name),
		Type:    "[]interface{}",
		Comment: fmt.Sprintf("The '%s' value can hold any type and any number of values", t.Name),
	}
	this.M = append(this.M, anyMember)
	this.F = append(this.F, &defs.MemberFunctionDef{
		Name:    fmt.Sprintf("%sLen", titleName),
		Comment: fmt.Sprintf("%sLen determines the number of elements able to be used for the Get%s and Remove%s functions", titleName, titleName, titleName),
		P:       this,
		Return:  []*defs.FunctionVarDef{{"l", "int"}},
		Body: func() string {
			return fmt.Sprintf("return len(t.%s)\n", anyMember.Name)
		},
	}, &defs.MemberFunctionDef{
		Name:    fmt.Sprintf("Get%s", titleName),
		Comment: fmt.Sprintf("Get%s returns the value for the specified index", titleName),
		P:       this,
		Args:    []*defs.FunctionVarDef{{"index", "int"}},
		Return:  []*defs.FunctionVarDef{{"v", deSlice(anyMember.Type)}},
		Body: func() string {
			var b bytes.Buffer
			b.WriteString(fmt.Sprintf("return t.%s[index]\n", anyMember.Name))
			return b.String()
		},
	}, &defs.MemberFunctionDef{
		Name:    fmt.Sprintf("Add%s", titleName),
		Comment: fmt.Sprintf("Add%s adds another value of %s", titleName, t.Name),
		P:       this,
		Args:    []*defs.FunctionVarDef{{"v", deSlice(anyMember.Type)}},
		Body: func() string {
			var b bytes.Buffer
			b.WriteString(fmt.Sprintf("t.%s = append(t.%s, v)\n", anyMember.Name, anyMember.Name))
			return b.String()
		},
	}, &defs.MemberFunctionDef{
		Name:    fmt.Sprintf("Remove%s", titleName),
		Comment: fmt.Sprintf("Remove%s deletes the value from the specified index", titleName),
		P:       this,
		Args:    []*defs.FunctionVarDef{{"index", "int"}},
		Body: func() string {
			sliceVar := fmt.Sprintf("t.%s", anyMember.Name)
			var b bytes.Buffer
			b.WriteString(fmt.Sprintf("copy(%s[index:], %s[index+1:])\n", sliceVar, sliceVar))
			b.WriteString(fmt.Sprintf("%s[len(%s)-1] = nil\n", sliceVar, sliceVar))
			b.WriteString(fmt.Sprintf("%s = %s[:len(%s)-1]\n", sliceVar, sliceVar, sliceVar))
			return b.String()
		},
	})
	i.F = append(i.F, []*defs.FunctionDef{
		{
			Name:    fmt.Sprintf("%sLen", titleName),
			Comment: fmt.Sprintf("%sLen determines the number of elements able to be used for the Get%s and Remove%s functions", titleName, titleName, titleName),
			Return:  []*defs.FunctionVarDef{{"l", "int"}},
		},
		{
			Name:    fmt.Sprintf("Get%s", titleName),
			Comment: fmt.Sprintf("Get%s returns the value for the specified index", titleName),
			Args:    []*defs.FunctionVarDef{{"index", "int"}},
			Return:  []*defs.FunctionVarDef{{"v", deSlice(anyMember.Type)}},
		},
		{
			Name:    fmt.Sprintf("Add%s", titleName),
			Comment: fmt.Sprintf("Add%s adds another value of %s", titleName, t.Name),
			Args:    []*defs.FunctionVarDef{{"v", deSlice(anyMember.Type)}},
		},
		{
			Name:    fmt.Sprintf("Remove%s", titleName),
			Comment: fmt.Sprintf("Remove%s deletes the value from the specified index", titleName),
			Args:    []*defs.FunctionVarDef{{"index", "int"}},
		},
	}...)
	tmpVar := fmt.Sprintf("tmp%sSlice", titleName)
	var b bytes.Buffer
	b.WriteString("// Begin generation by generateNonFunctionalAnyDefinition\n")
	b.WriteString(fmt.Sprintf("if k == \"%s\" {\n", t.Name))
	b.WriteString(fmt.Sprintf("if %s, ok := v.([]interface{}); ok {\n", tmpVar))
	b.WriteString(fmt.Sprintf("t.%s = %s\n", anyMember.Name, tmpVar))
	b.WriteString("handled = true\n")
	b.WriteString("} else {\n")
	b.WriteString(fmt.Sprintf("t.%s = []interface{}{v}\n", anyMember.Name))
	b.WriteString("handled = true\n")
	b.WriteString("}\n")
	b.WriteString("}\n")
	b.WriteString("// End generation by generateNonFunctionalAnyDefinition\n")
	d = b.String()
	var bs bytes.Buffer
	bs.WriteString("// Begin generation by generateNonFunctionalAnyDefinition\n")
	bs.WriteString(fmt.Sprintf("if t.%s != nil {\n", anyMember.Name))
	bs.WriteString(fmt.Sprintf("if len(t.%s) == 1 {\n", anyMember.Name))
	bs.WriteString(fmt.Sprintf("m[\"%s\"] = t.%s[0]\n", t.Name, anyMember.Name))
	bs.WriteString("} else {\n")
	bs.WriteString(fmt.Sprintf("m[\"%s\"] = t.%s\n", t.Name, anyMember.Name))
	bs.WriteString("}\n")
	bs.WriteString("}\n")
	bs.WriteString("// End generation by generateNonFunctionalAnyDefinition\n")
	s = bs.String()
	return
}

func generateSingleTypeDefinitions(t *defs.PropertyType, this *defs.StructDef, i *defs.InterfaceDef) (fd []*defs.FunctionDef, sd []*defs.StructDef, s, d string) {
	if t.Functional {
		return generateFunctionalSingleTypeDefinition(t, this, i)
	} else {
		return generateNonFunctionalSingleTypeDefinition(t, this, i)
	}
}

func generateFunctionalSingleTypeDefinition(t *defs.PropertyType, this *defs.StructDef, i *defs.InterfaceDef) (fd []*defs.FunctionDef, sd []*defs.StructDef, s, d string) {
	member := &defs.StructMember{
		Name:    cleanName(t.Name),
		Type:    Type(t.Range[0]),
		Comment: fmt.Sprintf("The functional '%s' value holds a single type and a single value", t.Name),
	}
	titleName := strings.Title(t.Name)
	isPtrType := isPtrType(member.Type)
	returnType := deref(member.Type)
	this.M = append(this.M, member)
	this.F = append(this.F, &defs.MemberFunctionDef{
		Name:    fmt.Sprintf("Has%s", titleName),
		Comment: fmt.Sprintf("Has%s determines whether the call to Get%s is safe", titleName, titleName),
		P:       this,
		Return:  []*defs.FunctionVarDef{{"ok", "bool"}},
		Body: func() string {
			return fmt.Sprintf("return t.%s != nil\n", member.Name)
		},
	}, &defs.MemberFunctionDef{
		Name:    fmt.Sprintf("Get%s", titleName),
		Comment: fmt.Sprintf("Get%s returns the value for %s", titleName, t.Name),
		P:       this,
		Return:  []*defs.FunctionVarDef{{"v", returnType}},
		Body: func() string {
			var b bytes.Buffer
			b.WriteString(fmt.Sprintf("return "))
			if isPtrType {
				b.WriteString("*")
			}
			b.WriteString(fmt.Sprintf("t.%s\n", member.Name))
			return b.String()
		},
	}, &defs.MemberFunctionDef{
		Name:    fmt.Sprintf("Set%s", titleName),
		Comment: fmt.Sprintf("Set%s sets the value of %s", titleName, t.Name),
		P:       this,
		Args:    []*defs.FunctionVarDef{{"v", returnType}},
		Body: func() string {
			var b bytes.Buffer
			b.WriteString(fmt.Sprintf("t.%s = ", member.Name))
			if isPtrType {
				b.WriteString("&")
			}
			b.WriteString("v\n")
			return b.String()
		},
	}, &defs.MemberFunctionDef{
		Name:    fmt.Sprintf("HasUnknown%s", titleName),
		Comment: fmt.Sprintf("HasUnknown%s determines whether the call to GetUnknown%s is safe", titleName, titleName),
		P:       this,
		Return:  []*defs.FunctionVarDef{{"ok", "bool"}},
		Body: func() string {
			return fmt.Sprintf("return t.unknown_ != nil && t.unknown_[\"%s\"] != nil\n", t.Name)
		},
	}, &defs.MemberFunctionDef{
		Name:    fmt.Sprintf("GetUnknown%s", titleName),
		Comment: fmt.Sprintf("GetUnknown%s returns the unknown value for %s", titleName, t.Name),
		P:       this,
		Return:  []*defs.FunctionVarDef{{"v", "interface{}"}},
		Body: func() string {
			var b bytes.Buffer
			b.WriteString(fmt.Sprintf("return t.unknown_[\"%s\"]\n", t.Name))
			return b.String()
		},
	}, &defs.MemberFunctionDef{
		Name:    fmt.Sprintf("SetUnknown%s", titleName),
		Comment: fmt.Sprintf("SetUnknown%s sets the unknown value of %s", titleName, t.Name),
		P:       this,
		Args:    []*defs.FunctionVarDef{{"i", "interface{}"}},
		Body: func() string {
			var b bytes.Buffer
			b.WriteString("if t.unknown_ == nil {\n")
			b.WriteString("t.unknown_ = make(map[string]interface{})")
			b.WriteString("}\n")
			b.WriteString(fmt.Sprintf("t.unknown_[\"%s\"] = i\n", t.Name))
			return b.String()
		},
	})
	i.F = append(i.F, []*defs.FunctionDef{
		{
			Name:    fmt.Sprintf("Has%s", titleName),
			Comment: fmt.Sprintf("Has%s determines whether the call to Get%s is safe", titleName, titleName),
			Return:  []*defs.FunctionVarDef{{"ok", "bool"}},
		},
		{
			Name:    fmt.Sprintf("Get%s", titleName),
			Comment: fmt.Sprintf("Get%s returns the value for %s", titleName, t.Name),
			Return:  []*defs.FunctionVarDef{{"v", returnType}},
		},
		{
			Name:    fmt.Sprintf("Set%s", titleName),
			Comment: fmt.Sprintf("Set%s sets the value of %s", titleName, t.Name),
			Args:    []*defs.FunctionVarDef{{"v", returnType}},
		},
		{
			Name:    fmt.Sprintf("HasUnknown%s", titleName),
			Comment: fmt.Sprintf("HasUnknown%s determines whether the call to GetUnknown%s is safe", titleName, titleName),
			Return:  []*defs.FunctionVarDef{{"ok", "bool"}},
		},
		{
			Name:    fmt.Sprintf("GetUnknown%s", titleName),
			Comment: fmt.Sprintf("GetUnknown%s returns the unknown value for %s", titleName, t.Name),
			Return:  []*defs.FunctionVarDef{{"v", "interface{}"}},
		},
		{
			Name:    fmt.Sprintf("SetUnknown%s", titleName),
			Comment: fmt.Sprintf("SetUnknown%s sets the unknown value of %s", titleName, t.Name),
			Args:    []*defs.FunctionVarDef{{"i", "interface{}"}},
		},
	}...)
	d = Deserialize(t.Range[0], t.Name, member.Name, false)
	s = Serialize(t.Range[0], t.Name, member.Name, false)
	return
}

func generateNonFunctionalSingleTypeDefinition(t *defs.PropertyType, this *defs.StructDef, i *defs.InterfaceDef) (fd []*defs.FunctionDef, sd []*defs.StructDef, s, d string) {
	member := &defs.StructMember{
		Name:    cleanName(t.Name),
		Type:    fmt.Sprintf("[]%s", deref(Type(t.Range[0]))),
		Comment: fmt.Sprintf("The '%s' value holds a single type and any number of values", t.Name),
	}
	titleName := strings.Title(t.Name)
	this.M = append(this.M, member)
	this.F = append(this.F, &defs.MemberFunctionDef{
		Name:    fmt.Sprintf("%sLen", titleName),
		Comment: fmt.Sprintf("%sLen determines the number of elements able to be used for the Get%s and Remove%s functions", titleName, titleName, titleName),
		P:       this,
		Return:  []*defs.FunctionVarDef{{"l", "int"}},
		Body: func() string {
			return fmt.Sprintf("return len(t.%s)\n", member.Name)
		},
	}, &defs.MemberFunctionDef{
		Name:    fmt.Sprintf("Get%s", titleName),
		Comment: fmt.Sprintf("Get%s returns the value for the specified index", titleName),
		P:       this,
		Args:    []*defs.FunctionVarDef{{"index", "int"}},
		Return:  []*defs.FunctionVarDef{{"v", deSlice(member.Type)}},
		Body: func() string {
			var b bytes.Buffer
			b.WriteString(fmt.Sprintf("return t.%s[index]\n", member.Name))
			return b.String()
		},
	}, &defs.MemberFunctionDef{
		Name:    fmt.Sprintf("Add%s", titleName),
		Comment: fmt.Sprintf("Add%s adds another value of %s", titleName, t.Name),
		P:       this,
		Args:    []*defs.FunctionVarDef{{"v", deSlice(member.Type)}},
		Body: func() string {
			var b bytes.Buffer
			b.WriteString(fmt.Sprintf("t.%s = append(t.%s, v)\n", member.Name, member.Name))
			return b.String()
		},
	}, &defs.MemberFunctionDef{
		Name:    fmt.Sprintf("Remove%s", titleName),
		Comment: fmt.Sprintf("Remove%s deletes the value from the specified index", titleName),
		P:       this,
		Args:    []*defs.FunctionVarDef{{"index", "int"}},
		Body: func() string {
			sliceVar := fmt.Sprintf("t.%s", member.Name)
			var b bytes.Buffer
			b.WriteString(fmt.Sprintf("copy(%s[index:], %s[index+1:])\n", sliceVar, sliceVar))
			if isPtrType(deSlice(member.Type)) {
				b.WriteString(fmt.Sprintf("%s[len(%s)-1] = nil\n", sliceVar, sliceVar))
			}
			b.WriteString(fmt.Sprintf("%s = %s[:len(%s)-1]\n", sliceVar, sliceVar, sliceVar))
			return b.String()
		},
	}, &defs.MemberFunctionDef{
		Name:    fmt.Sprintf("HasUnknown%s", titleName),
		Comment: fmt.Sprintf("HasUnknown%s determines whether the call to GetUnknown%s is safe", titleName, titleName),
		P:       this,
		Return:  []*defs.FunctionVarDef{{"ok", "bool"}},
		Body: func() string {
			return fmt.Sprintf("return t.unknown_ != nil && t.unknown_[\"%s\"] != nil\n", t.Name)
		},
	}, &defs.MemberFunctionDef{
		Name:    fmt.Sprintf("GetUnknown%s", titleName),
		Comment: fmt.Sprintf("GetUnknown%s returns the unknown value for %s", titleName, t.Name),
		P:       this,
		Return:  []*defs.FunctionVarDef{{"v", "interface{}"}},
		Body: func() string {
			var b bytes.Buffer
			b.WriteString(fmt.Sprintf("return t.unknown_[\"%s\"]\n", t.Name))
			return b.String()
		},
	}, &defs.MemberFunctionDef{
		Name:    fmt.Sprintf("SetUnknown%s", titleName),
		Comment: fmt.Sprintf("SetUnknown%s sets the unknown value of %s", titleName, t.Name),
		P:       this,
		Args:    []*defs.FunctionVarDef{{"i", "interface{}"}},
		Body: func() string {
			var b bytes.Buffer
			b.WriteString("if t.unknown_ == nil {\n")
			b.WriteString("t.unknown_ = make(map[string]interface{})")
			b.WriteString("}\n")
			b.WriteString(fmt.Sprintf("t.unknown_[\"%s\"] = i\n", t.Name))
			return b.String()
		},
	})
	i.F = append(i.F, []*defs.FunctionDef{
		{
			Name:    fmt.Sprintf("%sLen", titleName),
			Comment: fmt.Sprintf("%sLen determines the number of elements able to be used for the Get%s and Remove%s functions", titleName, titleName, titleName),
			Return:  []*defs.FunctionVarDef{{"l", "int"}},
		},
		{
			Name:    fmt.Sprintf("Get%s", titleName),
			Comment: fmt.Sprintf("Get%s returns the value for the specified index", titleName),
			Args:    []*defs.FunctionVarDef{{"index", "int"}},
			Return:  []*defs.FunctionVarDef{{"v", deSlice(member.Type)}},
		},
		{
			Name:    fmt.Sprintf("Add%s", titleName),
			Comment: fmt.Sprintf("Add%s adds another value of %s", titleName, t.Name),
			Args:    []*defs.FunctionVarDef{{"v", deSlice(member.Type)}},
		},
		{
			Name:    fmt.Sprintf("Remove%s", titleName),
			Comment: fmt.Sprintf("Remove%s deletes the value from the specified index", titleName),
			Args:    []*defs.FunctionVarDef{{"index", "int"}},
		},
		{
			Name:    fmt.Sprintf("HasUnknown%s", titleName),
			Comment: fmt.Sprintf("HasUnknown%s determines whether the call to GetUnknown%s is safe", titleName, titleName),
			Return:  []*defs.FunctionVarDef{{"ok", "bool"}},
		},
		{
			Name:    fmt.Sprintf("GetUnknown%s", titleName),
			Comment: fmt.Sprintf("GetUnknown%s returns the unknown value for %s", titleName, t.Name),
			Return:  []*defs.FunctionVarDef{{"v", "interface{}"}},
		},
		{
			Name:    fmt.Sprintf("SetUnknown%s", titleName),
			Comment: fmt.Sprintf("SetUnknown%s sets the unknown value of %s", titleName, t.Name),
			Args:    []*defs.FunctionVarDef{{"i", "interface{}"}},
		},
	}...)
	d = Deserialize(t.Range[0], t.Name, member.Name, true)
	s = Serialize(t.Range[0], t.Name, member.Name, true)
	return
}

func generateMultiTypeDefinitions(t *defs.PropertyType, this *defs.StructDef, i *defs.InterfaceDef) (fd []*defs.FunctionDef, sd []*defs.StructDef, s, d string) {
	if t.Functional {
		return generateFunctionalMultiTypeDefinition(t, this, i)
	} else {
		return generateNonFunctionalMultiTypeDefinition(t, this, i)
	}
}

func generateFunctionalMultiTypeDefinition(t *defs.PropertyType, this *defs.StructDef, i *defs.InterfaceDef) (fd []*defs.FunctionDef, sd []*defs.StructDef, s, d string) {
	titleName := strings.Title(t.Name)
	intermed, f := generateIntermediateTypeDefinition(t, strings.Title(this.Typename))
	fd = append(fd, f...)
	sd = append(sd, intermed)
	thisIntermed := &defs.StructMember{
		Name:    cleanName(t.Name),
		Type:    "*" + intermed.Typename,
		Comment: fmt.Sprintf("The functional '%s' value could have multiple types, but only a single value", t.Name),
	}
	this.M = append(this.M, thisIntermed)
	for idx, r := range t.Range {
		r := r
		member := &defs.StructMember{
			Name:    cleanName(Name(r)),
			Type:    Type(r),
			Comment: fmt.Sprintf("Stores possible %s type for %s property", Type(r), t.Name),
		}
		intermed.M = append(intermed.M, member)
		retKind := deref(member.Type)
		typeExtensionName := strings.Title(Name(r))
		if defs.IsOnlyOtherPropertyBesidesIRI(idx, t.Range) {
			typeExtensionName = ""
		}
		this.F = append(this.F, []*defs.MemberFunctionDef{
			{
				Name:    fmt.Sprintf("Is%s%s", titleName, typeExtensionName),
				Comment: fmt.Sprintf("Is%s%s determines whether the call to Get%s%s is safe", titleName, typeExtensionName, titleName, typeExtensionName),
				P:       this,
				Return:  []*defs.FunctionVarDef{{"ok", "bool"}},
				Body: func() string {
					return fmt.Sprintf("return t.%s != nil && t.%s.%s != nil\n", thisIntermed.Name, thisIntermed.Name, member.Name)
				},
			},
			{
				Name:    fmt.Sprintf("Get%s%s", titleName, typeExtensionName),
				Comment: fmt.Sprintf("Get%s%s returns the value safely if Is%s%s returned true", titleName, typeExtensionName, titleName, typeExtensionName),
				P:       this,
				Return:  []*defs.FunctionVarDef{{"v", retKind}},
				Body: func() string {
					var b bytes.Buffer
					b.WriteString("return ")
					if isPtrType(Type(r)) {
						b.WriteString("*")
					}
					b.WriteString(fmt.Sprintf("t.%s.%s\n", thisIntermed.Name, member.Name))
					return b.String()
				},
			},
			{
				Name:    fmt.Sprintf("Set%s%s", titleName, typeExtensionName),
				Comment: fmt.Sprintf("Set%s%s sets the value of %s to be of %s type", titleName, typeExtensionName, cleanName(t.Name), retKind),
				P:       this,
				Args:    []*defs.FunctionVarDef{{"v", retKind}},
				Body: func() string {
					var b bytes.Buffer
					b.WriteString(fmt.Sprintf("t.%s = &%s{%s:", thisIntermed.Name, intermed.Typename, member.Name))
					if isPtrType(Type(r)) {
						b.WriteString("&")
					}
					b.WriteString("v}\n")
					return b.String()
				},
			},
		}...)
		i.F = append(i.F, []*defs.FunctionDef{
			{
				Name:    fmt.Sprintf("Is%s%s", titleName, typeExtensionName),
				Comment: fmt.Sprintf("Is%s%s determines whether the call to Get%s%s is safe", titleName, typeExtensionName, titleName, typeExtensionName),
				Return:  []*defs.FunctionVarDef{{"ok", "bool"}},
			},
			{
				Name:    fmt.Sprintf("Get%s%s", titleName, typeExtensionName),
				Comment: fmt.Sprintf("Get%s%s returns the value safely if Is%s%s returned true", titleName, typeExtensionName, titleName, typeExtensionName),
				Return:  []*defs.FunctionVarDef{{"v", retKind}},
			},
			{
				Name:    fmt.Sprintf("Set%s%s", titleName, typeExtensionName),
				Comment: fmt.Sprintf("Set%s%s sets the value of %s to be of %s type", titleName, typeExtensionName, cleanName(t.Name), retKind),
				Args:    []*defs.FunctionVarDef{{"v", retKind}},
			},
		}...)
	}
	var b bytes.Buffer
	b.WriteString("// Begin generation by generateFunctionalMultiTypeDefinition\n")
	b.WriteString(fmt.Sprintf("if k == \"%s\" {\n", t.Name))
	b.WriteString(fmt.Sprintf("t.%s, err = deserialize%s(v)\n", thisIntermed.Name, strings.Title(intermed.Typename)))
	b.WriteString("if err != nil {\nreturn err\n}\n")
	b.WriteString("handled = true\n")
	b.WriteString("}\n")
	b.WriteString("// End generation by generateFunctionalMultiTypeDefinition\n")
	d = b.String()
	var bs bytes.Buffer
	bs.WriteString("// Begin generation by generateFunctionalMultiTypeDefinition\n")
	bs.WriteString(fmt.Sprintf("if t.%s != nil {\n", thisIntermed.Name))
	bs.WriteString(fmt.Sprintf("if v, err := serialize%s(t.%s); err == nil {\n", strings.Title(intermed.Typename), thisIntermed.Name))
	bs.WriteString(fmt.Sprintf("m[\"%s\"] = v\n", t.Name))
	bs.WriteString("} else {\n")
	bs.WriteString("return m, err\n")
	bs.WriteString("}\n")
	bs.WriteString("}\n")
	bs.WriteString("// End generation by generateFunctionalMultiTypeDefinition\n")
	s = bs.String()
	generateUnknownFunctions(t, this, thisIntermed, intermed)
	return
}

func generateNonFunctionalMultiTypeDefinition(t *defs.PropertyType, this *defs.StructDef, i *defs.InterfaceDef) (fd []*defs.FunctionDef, sd []*defs.StructDef, s, d string) {
	titleName := strings.Title(t.Name)
	intermed, f := generateIntermediateTypeDefinition(t, strings.Title(this.Typename))
	fd = append(fd, f...)
	sd = append(sd, intermed)
	thisIntermed := &defs.StructMember{
		Name:    cleanName(t.Name),
		Type:    fmt.Sprintf("[]*%s", intermed.Typename),
		Comment: fmt.Sprintf("The '%s' value could have multiple types and values", t.Name),
	}
	this.M = append(this.M, thisIntermed)
	for idx, r := range t.Range {
		r := r
		member := &defs.StructMember{
			Name:    cleanName(Name(r)),
			Type:    Type(r),
			Comment: fmt.Sprintf("Stores possible %s type for %s property", Type(r), t.Name),
		}
		intermed.M = append(intermed.M, member)
		retKind := deref(member.Type)
		typeExtensionName := strings.Title(Name(r))
		if defs.IsOnlyOtherPropertyBesidesIRI(idx, t.Range) {
			typeExtensionName = ""
		}
		if idx == 0 {
			this.F = append(this.F, &defs.MemberFunctionDef{
				Name:    fmt.Sprintf("%sLen", titleName),
				Comment: fmt.Sprintf("%sLen determines the number of elements able to be used for the Is%s%s, Get%s%s, and Remove%s%s functions", titleName, titleName, typeExtensionName, titleName, typeExtensionName, titleName, typeExtensionName),
				P:       this,
				Return:  []*defs.FunctionVarDef{{"l", "int"}},
				Body: func() string {
					return fmt.Sprintf("return len(t.%s)\n", thisIntermed.Name)
				},
			})
			i.F = append(i.F, &defs.FunctionDef{
				Name:    fmt.Sprintf("%sLen", titleName),
				Comment: fmt.Sprintf("%sLen determines the number of elements able to be used for the Is%s%s, Get%s%s, and Remove%s%s functions", titleName, titleName, typeExtensionName, titleName, typeExtensionName, titleName, typeExtensionName),
				Return:  []*defs.FunctionVarDef{{"l", "int"}},
			})
		}
		this.F = append(this.F, []*defs.MemberFunctionDef{
			{
				Name:    fmt.Sprintf("Is%s%s", titleName, typeExtensionName),
				Comment: fmt.Sprintf("Is%s%s determines whether the call to Get%s%s is safe for the specified index", titleName, typeExtensionName, titleName, typeExtensionName),
				P:       this,
				Args:    []*defs.FunctionVarDef{{"index", "int"}},
				Return:  []*defs.FunctionVarDef{{"ok", "bool"}},
				Body: func() string {
					return fmt.Sprintf("return t.%s[index].%s != nil\n", thisIntermed.Name, member.Name)
				},
			},
			{
				Name:    fmt.Sprintf("Get%s%s", titleName, typeExtensionName),
				Comment: fmt.Sprintf("Get%s%s returns the value safely if Is%s%s returned true for the specified index", titleName, typeExtensionName, titleName, typeExtensionName),
				P:       this,
				Args:    []*defs.FunctionVarDef{{"index", "int"}},
				Return:  []*defs.FunctionVarDef{{"v", retKind}},
				Body: func() string {
					var b bytes.Buffer
					b.WriteString("return ")
					if isPtrType(Type(r)) {
						b.WriteString("*")
					}
					b.WriteString(fmt.Sprintf("t.%s[index].%s\n", thisIntermed.Name, member.Name))
					return b.String()
				},
			},
			{
				Name:    fmt.Sprintf("Add%s%s", titleName, typeExtensionName),
				Comment: fmt.Sprintf("Add%s%s adds another value of %s to be of %s type", titleName, typeExtensionName, cleanName(t.Name), retKind),
				P:       this,
				Args:    []*defs.FunctionVarDef{{"v", retKind}},
				Body: func() string {
					var b bytes.Buffer
					b.WriteString(fmt.Sprintf("t.%s = append(t.%s, &%s{%s:", thisIntermed.Name, thisIntermed.Name, intermed.Typename, member.Name))
					if isPtrType(Type(r)) {
						b.WriteString("&")
					}
					b.WriteString("v})\n")
					return b.String()
				},
			},
			{
				Name:    fmt.Sprintf("Remove%s%s", titleName, typeExtensionName),
				Comment: fmt.Sprintf("Remove%s%s deletes the value from the specified index", titleName, typeExtensionName),
				P:       this,
				Args:    []*defs.FunctionVarDef{{"index", "int"}},
				Body: func() string {
					sliceVar := fmt.Sprintf("t.%s", thisIntermed.Name)
					var b bytes.Buffer
					b.WriteString(fmt.Sprintf("copy(%s[index:], %s[index+1:])\n", sliceVar, sliceVar))
					b.WriteString(fmt.Sprintf("%s[len(%s)-1] = %s\n", sliceVar, sliceVar, "nil"))
					b.WriteString(fmt.Sprintf("%s = %s[:len(%s)-1]\n", sliceVar, sliceVar, sliceVar))
					return b.String()
				},
			},
		}...)
		i.F = append(i.F, []*defs.FunctionDef{
			{
				Name:    fmt.Sprintf("Is%s%s", titleName, typeExtensionName),
				Comment: fmt.Sprintf("Is%s%s determines whether the call to Get%s%s is safe for the specified index", titleName, typeExtensionName, titleName, typeExtensionName),
				Args:    []*defs.FunctionVarDef{{"index", "int"}},
				Return:  []*defs.FunctionVarDef{{"ok", "bool"}},
			},
			{
				Name:    fmt.Sprintf("Get%s%s", titleName, typeExtensionName),
				Comment: fmt.Sprintf("Get%s%s returns the value safely if Is%s%s returned true for the specified index", titleName, typeExtensionName, titleName, typeExtensionName),
				Args:    []*defs.FunctionVarDef{{"index", "int"}},
				Return:  []*defs.FunctionVarDef{{"v", retKind}},
			},
			{
				Name:    fmt.Sprintf("Add%s%s", titleName, typeExtensionName),
				Comment: fmt.Sprintf("Add%s%s adds another value of %s to be of %s type", titleName, typeExtensionName, cleanName(t.Name), retKind),
				Args:    []*defs.FunctionVarDef{{"v", retKind}},
			},
			{
				Name:    fmt.Sprintf("Remove%s%s", titleName, typeExtensionName),
				Comment: fmt.Sprintf("Remove%s%s deletes the value from the specified index", titleName, typeExtensionName),
				Args:    []*defs.FunctionVarDef{{"index", "int"}},
			},
		}...)
	}
	var b bytes.Buffer
	b.WriteString("// Begin generation by generateNonFunctionalMultiTypeDefinition\n")
	b.WriteString(fmt.Sprintf("if k == \"%s\" {\n", t.Name))
	b.WriteString("if tmpMap, ok := v.(map[string]interface{}); ok {\n")
	b.WriteString(fmt.Sprintf("tmp, err := deserialize%s(tmpMap)\n", strings.Title(intermed.Typename)))
	b.WriteString("if err != nil {\nreturn err\n}\n")
	b.WriteString(fmt.Sprintf("t.%s = []*%s{tmp}\n", thisIntermed.Name, intermed.Typename))
	b.WriteString("handled = true\n")
	b.WriteString("} else if tmpSlice, ok := v.([]interface{}); ok {\n")
	b.WriteString(fmt.Sprintf("t.%s, err = deserializeSlice%s(tmpSlice)\n", thisIntermed.Name, strings.Title(intermed.Typename)))
	b.WriteString("if err != nil {\nreturn err\n}\n")
	b.WriteString("handled = true\n")
	b.WriteString("} else {\n")
	b.WriteString(fmt.Sprintf("tmp := &%s{}\n", intermed.Typename))
	b.WriteString("err = tmp.Deserialize(v)\n")
	b.WriteString("if err != nil {\nreturn err\n}\n")
	b.WriteString(fmt.Sprintf("t.%s = []*%s{tmp}\n", thisIntermed.Name, intermed.Typename))
	b.WriteString("handled = true\n")
	b.WriteString("}\n")
	b.WriteString("}\n")
	b.WriteString("// End generation by generateNonFunctionalMultiTypeDefinition\n")
	d = b.String()
	var bs bytes.Buffer
	bs.WriteString("// Begin generation by generateNonFunctionalMultiTypeDefinition\n")
	bs.WriteString(fmt.Sprintf("if v, err := serializeSlice%s(t.%s); err == nil && v != nil {\n", strings.Title(intermed.Typename), thisIntermed.Name))
	bs.WriteString("if len(v) == 1 {\n")
	bs.WriteString(fmt.Sprintf("m[\"%s\"] = v[0]\n", t.Name))
	bs.WriteString("} else {\n")
	bs.WriteString(fmt.Sprintf("m[\"%s\"] = v\n", t.Name))
	bs.WriteString("}\n")
	bs.WriteString("} else if err != nil {\n")
	bs.WriteString("return m, err\n")
	bs.WriteString("}\n")
	bs.WriteString("// End generation by generateNonFunctionalMultiTypeDefinition\n")
	s = bs.String()
	generateUnknownFunctions(t, this, thisIntermed, intermed)
	return
}

func generateIntermediateTypeDefinition(t *defs.PropertyType, parentTitle string) (intermed *defs.StructDef, f []*defs.FunctionDef) {
	intermed = &defs.StructDef{
		Typename: fmt.Sprintf("%s%sIntermediateType", cleanName(t.Name), parentTitle),
		Comment:  fmt.Sprintf("%s%sIntermediateType will only have one of its values set at most", cleanName(t.Name), parentTitle),
		M:        []*defs.StructMember{{"unknown_", "interface{}", "An unknown value."}},
	}
	intermed.F = []*defs.MemberFunctionDef{
		{
			Name:    "Deserialize",
			Comment: "Deserialize takes an interface{} and attempts to create a valid intermediate type.",
			P:       intermed,
			Args:    []*defs.FunctionVarDef{{"i", "interface{}"}},
			Return:  []*defs.FunctionVarDef{{"err", "error"}},
			Body: func() string {
				bHasType := false
				for _, r := range t.Range {
					if r.T != nil {
						bHasType = true
						break
					}
				}
				var b bytes.Buffer
				b.WriteString("matched := false\n")
				b.WriteString("if m, ok := i.(map[string]interface{}); ok {\n")
				if bHasType {
					b.WriteString("if tv, ok := m[\"type\"]; ok {\n")
					b.WriteString("var types []string\n")
					b.WriteString("if tvs, ok := tv.([]interface{}); ok {\n")
					b.WriteString("for _, tvi := range tvs {\n")
					b.WriteString("if typeString, ok := tvi.(string); ok {\n")
					b.WriteString("types = append(types, typeString)\n")
					b.WriteString("}\n")
					b.WriteString("}\n")
					b.WriteString("} else if typeString, ok := tv.(string); ok {\n")
					b.WriteString("types = append(types, typeString)\n")
					b.WriteString("}\n")
					for _, r := range t.Range {
						if r.T != nil {
							var resolve string
							if isALinkType(r.T) {
								resolve = resolveLinkName
							} else if isAnObjectType(r.T) {
								resolve = resolveObjectName
							} else {
								panic("Unknown resolution function: not object and not link")
							}
							b.WriteString("if !matched {\n")
							b.WriteString("for _, kind := range types {\n")
							b.WriteString(fmt.Sprintf("if t.%s, ok = %s(kind).(%sType); t.%s != nil && ok {\n", cleanName(Name(r)), resolve, Name(r), cleanName(Name(r))))
							b.WriteString(fmt.Sprintf("err = t.%s.Deserialize(m)\n", cleanName(Name(r))))
							b.WriteString("matched = true\n")
							b.WriteString("break\n")
							b.WriteString("}\n")
							b.WriteString("}\n")
							b.WriteString("}\n")
						}
					}
					b.WriteString("} else {\n")
					b.WriteString("t.unknown_ = m\n")
					b.WriteString("}\n")
				} else {
					b.WriteString("err = fmt.Errorf(\"Given map but nothing to do with it for this type: %v\", m)")
				}
				b.WriteString("} else {\n")
				for _, r := range t.Range {
					if r.V != nil {
						b.WriteString("if !matched {\n")
						b.WriteString(fmt.Sprintf("t.%s, err = %s(i)\n", cleanName(Name(r)), r.V.DeserializeFn.Name))
						b.WriteString("if err != nil {\n")
						b.WriteString(fmt.Sprintf("t.%s = nil\n", cleanName(Name(r))))
						b.WriteString("} else {\n")
						b.WriteString("matched = true\n")
						b.WriteString("}\n")
						b.WriteString("}\n")
					}
				}
				b.WriteString("}\n")
				b.WriteString("if !matched {\n")
				b.WriteString(fmt.Sprintf("t.unknown_ = %s(i)\n", unknownValueDeserializeFnName))
				b.WriteString("}\n")
				b.WriteString("return\n")
				return b.String()
			},
		},
		{
			Name:    "Serialize",
			Comment: "Serialize turns this object into an interface{}.",
			P:       intermed,
			Return:  []*defs.FunctionVarDef{{"i", "interface{}"}, {"err", "error"}},
			Body: func() string {
				var b bytes.Buffer
				for _, r := range t.Range {
					b.WriteString(fmt.Sprintf("if t.%s != nil {\n", cleanName(Name(r))))
					if r.T != nil {
						b.WriteString(fmt.Sprintf("i, err = t.%s.Serialize()\n", cleanName(Name(r))))
						b.WriteString("return\n")
					} else if r.V != nil {
						b.WriteString(fmt.Sprintf("i = %s(*t.%s)\n", r.V.SerializeFn.Name, cleanName(Name(r))))
						b.WriteString("return\n")
					} else {
						b.WriteString(fmt.Sprintf("i = t.%s\n", cleanName(Name(r))))
						b.WriteString("return\n")
					}
					b.WriteString("}\n")
				}
				b.WriteString("if t.unknown_ != nil {\n")
				b.WriteString(fmt.Sprintf("i = %s(t.unknown_)\n", unknownValueSerializeFnName))
				b.WriteString("return\n")
				b.WriteString("}\n")
				b.WriteString(fmt.Sprintf("err = fmt.Errorf(\"Could not serialize all-nil for type %s: %%v\", t)\n", intermed.Typename))
				b.WriteString("return")
				return b.String()
			},
		},
	}
	f = append(f, []*defs.FunctionDef{
		{
			Name:    fmt.Sprintf("deserialize%s", strings.Title(intermed.Typename)),
			Comment: fmt.Sprintf("deserialize%s will accept a map to create a %s", intermed.Typename, intermed.Typename),
			Args:    []*defs.FunctionVarDef{{"in", "interface{}"}},
			Return:  []*defs.FunctionVarDef{{"t", "*" + intermed.Typename}, {"err", "error"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString(fmt.Sprintf("tmp := &%s{}\n", intermed.Typename))
				b.WriteString("err = tmp.Deserialize(in)\n")
				b.WriteString("return tmp, err\n")
				return b.String()
			},
		},
		{
			Name:    fmt.Sprintf("deserializeSlice%s", strings.Title(intermed.Typename)),
			Comment: fmt.Sprintf("deserializeSlice %s will accept a slice to create a slice of %s", intermed.Typename, intermed.Typename),
			Args:    []*defs.FunctionVarDef{{"in", "[]interface{}"}},
			Return:  []*defs.FunctionVarDef{{"t", "[]*" + intermed.Typename}, {"err", "error"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("for _, i := range in {\n")
				b.WriteString(fmt.Sprintf("tmp := &%s{}\n", intermed.Typename))
				b.WriteString("err = tmp.Deserialize(i)\n")
				b.WriteString("if err != nil {\nreturn\n}\n")
				b.WriteString("t = append(t, tmp)\n")
				b.WriteString("}\n")
				b.WriteString("return\n")
				return b.String()
			},
		},
		{
			Name:    fmt.Sprintf("serialize%s", strings.Title(intermed.Typename)),
			Comment: fmt.Sprintf("serialize%s will accept a %s to create a map", intermed.Typename, intermed.Typename),
			Args:    []*defs.FunctionVarDef{{"t", "*" + intermed.Typename}},
			Return:  []*defs.FunctionVarDef{{"i", "interface{}"}, {"err", "error"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("i, err = t.Serialize()\n")
				b.WriteString("return\n")
				return b.String()
			},
		},
		{
			Name:    fmt.Sprintf("serializeSlice%s", strings.Title(intermed.Typename)),
			Comment: fmt.Sprintf("serializeSlice%s will accept a slice of %s to create a slice result", intermed.Typename, intermed.Typename),
			Args:    []*defs.FunctionVarDef{{"s", "[]*" + intermed.Typename}},
			Return:  []*defs.FunctionVarDef{{"out", "[]interface{}"}, {"err", "error"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("for _, t := range s {\n")
				b.WriteString("v, err := t.Serialize()\n")
				b.WriteString("if err != nil {\nreturn nil, err\n}\n")
				b.WriteString("out = append(out, v)\n")
				b.WriteString("}\n")
				b.WriteString("return\n")
				return b.String()
			},
		},
	}...)
	return
}

func generateUnknownFunctions(t *defs.PropertyType, this *defs.StructDef, thisIntermed *defs.StructMember, intermed *defs.StructDef) {
	titleName := strings.Title(t.Name)
	this.F = append(this.F, []*defs.MemberFunctionDef{
		{
			Name:    fmt.Sprintf("HasUnknown%s", titleName),
			Comment: fmt.Sprintf("HasUnknown%s determines whether the call to GetUnknown%s is safe", titleName, titleName),
			P:       this,
			Return:  []*defs.FunctionVarDef{{"ok", "bool"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString(fmt.Sprintf("return t.%s != nil && t.%s", thisIntermed.Name, thisIntermed.Name))
				if !t.Functional {
					b.WriteString("[0]")
				}
				b.WriteString(".unknown_ != nil\n")
				return b.String()
			},
		}, &defs.MemberFunctionDef{
			Name:    fmt.Sprintf("GetUnknown%s", titleName),
			Comment: fmt.Sprintf("GetUnknown%s returns the unknown value for %s", titleName, t.Name),
			P:       this,
			Return:  []*defs.FunctionVarDef{{"v", "interface{}"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString(fmt.Sprintf("return t.%s", thisIntermed.Name))
				if !t.Functional {
					b.WriteString("[0]")
				}
				b.WriteString(".unknown_\n")
				return b.String()
			},
		}, &defs.MemberFunctionDef{
			Name:    fmt.Sprintf("SetUnknown%s", titleName),
			Comment: fmt.Sprintf("SetUnknown%s sets the unknown value of %s", titleName, t.Name),
			P:       this,
			Args:    []*defs.FunctionVarDef{{"i", "interface{}"}},
			Body: func() string {
				var b bytes.Buffer
				b.WriteString("if t.unknown_ == nil {\n")
				b.WriteString("t.unknown_ = make(map[string]interface{})")
				b.WriteString("}\n")
				b.WriteString(fmt.Sprintf("tmp := &%s{}\n", intermed.Typename))
				b.WriteString("tmp.unknown_ = i\n")
				b.WriteString(fmt.Sprintf("t.%s = ", thisIntermed.Name))
				if t.Functional {
					b.WriteString("tmp\n")
				} else {
					b.WriteString(fmt.Sprintf("append(t.%s, tmp)\n", thisIntermed.Name))
				}
				return b.String()
			},
		},
	}...)
}

func isSingleType(t *defs.PropertyType) bool {
	return len(t.Range) == 1
}

func isAny(t *defs.PropertyType) bool {
	for _, r := range t.Range {
		if r.Any {
			return true
		}
	}
	return false
}

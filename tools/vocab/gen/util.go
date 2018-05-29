package gen

import (
	"bytes"
	"fmt"
)

func cleanName(s string) string {
	if s == "type" {
		s = "typeName"
	} else if s == "string" {
		s = "stringName"
	}
	return s
}

func isSlice(k string) bool {
	return len(k) > 1 && k[0] == '[' && k[1] == ']'
}

func deSlice(k string) string {
	if k[0] == '[' && k[1] == ']' {
		return k[2:]
	}
	return k
}

func isPtrType(k string) bool {
	return k[0] == '*'
}

func deref(k string) string {
	if k[0] == '*' {
		return k[1:]
	}
	return k
}

func isIRIType(k string) bool {
	return k == "*url.URL"
}

func deserializeCode(b *bytes.Buffer, parseCode, varName, typeName, field string, slice, derefAppend bool) {
	if slice {
		sliceDeserializeCode(b, parseCode, varName, typeName, field, derefAppend)
	} else {
		singleDeserializeCode(b, parseCode, varName, typeName, field)
	}
}

func sliceDeserializeCode(b *bytes.Buffer, parseCode, mapName, typeName, field string, deref bool) {
	b.WriteString(fmt.Sprintf("if k == \"%s\" {\n", mapName))
	b.WriteString("if tmpSlice, ok := v.([]interface{}); ok {\n")
	b.WriteString(fmt.Sprintf("for _, tmpElem := range tmpSlice {\n"))
	b.WriteString(fmt.Sprintf("if v, ok := tmpElem.(%s); ok {\n", typeName))
	b.WriteString(parseCode)
	b.WriteString("if err != nil {\nreturn err\n}\n")
	if deref {
		b.WriteString(fmt.Sprintf("t.%s = append(t.%s, *tmp)\n", field, field))
	} else {
		b.WriteString(fmt.Sprintf("t.%s = append(t.%s, tmp)\n", field, field))
	}
	b.WriteString("handled = true\n")
	b.WriteString("}\n")
	b.WriteString("}\n")
	b.WriteString(fmt.Sprintf("} else if v, ok := v.(%s); ok {\n", typeName))
	b.WriteString(parseCode)
	b.WriteString("if err != nil {\nreturn err\n}\n")
	if deref {
		b.WriteString(fmt.Sprintf("t.%s = append(t.%s, *tmp)\n", field, field))
	} else {
		b.WriteString(fmt.Sprintf("t.%s = append(t.%s, tmp)\n", field, field))
	}
	b.WriteString("handled = true\n")
	b.WriteString("}\n")
	b.WriteString("}\n")
}

func singleDeserializeCode(b *bytes.Buffer, parseCode, mapName, typeName, field string) {
	b.WriteString(fmt.Sprintf("if k == \"%s\" {\n", mapName))
	b.WriteString(fmt.Sprintf("if v, ok := v.(%s); ok {\n", typeName))
	b.WriteString(parseCode)
	b.WriteString("if err != nil {\nreturn err\n}\n")
	b.WriteString(fmt.Sprintf("t.%s = tmp\n", field))
	b.WriteString("handled = true\n")
	b.WriteString("}\n")
	b.WriteString("}\n")
}

func serializeCode(b *bytes.Buffer, sCode, mapName, field string, slice bool) {
	if slice {
		sliceSerializeCode(b, sCode, mapName, field)
	} else {
		singleSerializeCode(b, sCode, mapName, field)
	}
}

func sliceSerializeCode(b *bytes.Buffer, serializeCode, mapName, field string) {
	b.WriteString(fmt.Sprintf("var %sTemp []interface{}\n", mapName))
	b.WriteString(fmt.Sprintf("for _, v := range t.%s {\n", field))
	b.WriteString(serializeCode)
	b.WriteString(fmt.Sprintf("%sTemp = append(%sTemp, tmp)\n", mapName, mapName))
	b.WriteString("}\n")
	b.WriteString(fmt.Sprintf("if %sTemp != nil {\n", mapName))
	b.WriteString(fmt.Sprintf("if len(%sTemp) == 1 {\n", mapName))
	b.WriteString(fmt.Sprintf("m[\"%s\"] = %sTemp[0]\n", mapName, mapName))
	b.WriteString("} else {\n")
	b.WriteString(fmt.Sprintf("m[\"%s\"] = %sTemp\n", mapName, mapName))
	b.WriteString("}\n")
	b.WriteString("}\n")
}

func singleSerializeCode(b *bytes.Buffer, serializeCode, mapName, field string) {
	b.WriteString(fmt.Sprintf("if t.%s != nil {\n", field))
	b.WriteString(fmt.Sprintf("%sSerializeFunc := func() (interface{}, error) {\n", mapName))
	b.WriteString(fmt.Sprintf("v := t.%s\n", field))
	b.WriteString(serializeCode)
	b.WriteString("return tmp, nil\n")
	b.WriteString("}\n")
	b.WriteString(fmt.Sprintf("%sResult, err := %sSerializeFunc()\n", mapName, mapName))
	b.WriteString("if err == nil {\n")
	b.WriteString(fmt.Sprintf("m[\"%s\"] = %sResult\n", mapName, mapName))
	b.WriteString("} else {\n")
	b.WriteString("return m, err\n")
	b.WriteString("}\n")
	b.WriteString("}\n")
}

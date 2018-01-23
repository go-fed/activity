package defs

import (
	"bytes"
	"fmt"
	"strings"
)

type StructMember struct {
	Name    string
	Type    string
	Comment string
}

func (s *StructMember) Generate() string {
	return fmt.Sprintf("// %s\n%s %s", s.Comment, s.Name, s.Type)
}

type StructDef struct {
	Typename string
	M        []*StructMember
	F        []*MemberFunctionDef
	Compose  []*StructDef
	Comment  string
}

func (s *StructDef) Generate() string {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("// %s\ntype %s struct {\n", s.Comment, s.Typename))
	for _, c := range s.Compose {
		b.WriteString(fmt.Sprintf("*%s\n", c.Typename))
	}
	for _, m := range s.M {
		b.WriteString(m.Generate())
		b.WriteString("\n")
	}
	b.WriteString("}\n\n")
	for _, f := range s.F {
		b.WriteString(f.Generate())
		b.WriteString("\n\n")
	}
	return b.String()
}

type InterfaceDef struct {
	Typename string
	Comment  string
	// Body is ignored
	O []string // Other interfaceDef Typenames
	F []*FunctionDef
}

func (i *InterfaceDef) Generate() string {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("// %s\n", i.Comment))
	b.WriteString(fmt.Sprintf("type %s interface {\n", i.Typename))
	for _, o := range i.O {
		b.WriteString(fmt.Sprintf("%s\n", o))
	}
	for _, f := range i.F {
		b.WriteString(fmt.Sprintf("%s(", f.Name))
		var args []string
		for _, a := range f.Args {
			args = append(args, a.Generate())
		}
		b.WriteString(strings.Join(args, ", "))
		b.WriteString(")")
		if len(f.Return) > 0 {
			b.WriteString(" (")
			var ret []string
			for _, r := range f.Return {
				ret = append(ret, r.Generate())
			}
			b.WriteString(strings.Join(ret, ", "))
			b.WriteString(")")
		}
		b.WriteString("\n")
	}
	b.WriteString("}\n")
	return b.String()
}

type PackageDef struct {
	Name    string
	Comment string
	Imports []string
	Defs    []*StructDef
	F       []*FunctionDef
	I       []*InterfaceDef
	Raw     string
}

func (p *PackageDef) Generate() string {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("// %s\npackage %s\n\n", p.Comment, p.Name))
	if len(p.Imports) > 0 {
		b.WriteString("import (\n")
		for _, i := range p.Imports {
			b.WriteString(fmt.Sprintf("\"%s\"\n", i))
		}
		b.WriteString(")\n\n")
	}
	if len(p.Raw) > 0 {
		b.WriteString(p.Raw)
		b.WriteString("\n\n")
	}
	for _, i := range p.I {
		b.WriteString(i.Generate())
		b.WriteString("\n\n")
	}
	for _, d := range p.Defs {
		b.WriteString(d.Generate())
		b.WriteString("\n\n")
	}
	for _, f := range p.F {
		b.WriteString(f.Generate())
		b.WriteString("\n\n")
	}
	return b.String()
}

type MemberFunctionDef struct {
	Name    string
	Comment string
	P       *StructDef
	Args    []*FunctionVarDef
	Return  []*FunctionVarDef
	Body    func() string
}

func (f *MemberFunctionDef) Generate() string {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("// %s\nfunc (t *%s) %s(", f.Comment, f.P.Typename, f.Name))
	var args []string
	for _, a := range f.Args {
		args = append(args, a.Generate())
	}
	b.WriteString(strings.Join(args, ", "))
	b.WriteString(")")
	if len(f.Return) > 0 {
		b.WriteString(" (")
		var ret []string
		for _, r := range f.Return {
			ret = append(ret, r.Generate())
		}
		b.WriteString(strings.Join(ret, ", "))
		b.WriteString(")")
	}
	b.WriteString(" {\n")
	b.WriteString(f.Body())
	b.WriteString("\n}\n\n")
	return b.String()
}

type FunctionDef struct {
	Name    string
	Comment string
	Args    []*FunctionVarDef
	Return  []*FunctionVarDef
	Body    func() string
}

func (f *FunctionDef) Generate() string {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("// %s\nfunc %s(", f.Comment, f.Name))
	var args []string
	for _, a := range f.Args {
		args = append(args, a.Generate())
	}
	b.WriteString(strings.Join(args, ", "))
	b.WriteString(")")
	if len(f.Return) > 0 {
		b.WriteString(" (")
		var ret []string
		for _, r := range f.Return {
			ret = append(ret, r.Generate())
		}
		b.WriteString(strings.Join(ret, ", "))
		b.WriteString(")")
	}
	b.WriteString(" {\n")
	b.WriteString(f.Body())
	b.WriteString("\n}\n\n")
	return b.String()
}

type FunctionVarDef struct {
	Name string
	Type string
}

func (f *FunctionVarDef) Generate() string {
	return fmt.Sprintf("%s %s", f.Name, f.Type)
}

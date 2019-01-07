package codegen

import (
	"github.com/dave/jennifer/jen"
	"sort"
)

type sortedFunctionSignature []FunctionSignature

func (s sortedFunctionSignature) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

func (s sortedFunctionSignature) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortedFunctionSignature) Len() int {
	return len(s)
}

type FunctionSignature struct {
	Name    string
	Params  []jen.Code
	Ret     []jen.Code
	Comment string
}

type Interface struct {
	qual      *jen.Statement
	name      string
	functions []FunctionSignature
	comment   string
}

func NewInterface(pkg, name string,
	funcs []FunctionSignature,
	comment string) *Interface {
	i := &Interface{
		qual:      jen.Qual(pkg, name),
		name:      name,
		functions: funcs,
		comment:   comment,
	}
	sort.Sort(sortedFunctionSignature(i.functions))
	return i
}

func (i Interface) Definition() jen.Code {
	stmts := jen.Empty()
	if len(i.comment) > 0 {
		stmts = jen.Comment(i.comment).Line()
	}
	defs := make([]jen.Code, 0, len(i.functions))
	for _, fn := range i.functions {
		def := jen.Empty()
		if len(fn.Comment) > 0 {
			def.Comment(fn.Comment).Line()
		}
		def.Id(fn.Name).Params(fn.Params...)
		if len(fn.Ret) > 0 {
			def.Params(fn.Ret...)
		}
		defs = append(defs, def)
	}
	return stmts.Type().Id(i.name).Interface(defs...)
}

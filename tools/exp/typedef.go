package exp

import (
	"github.com/dave/jennifer/jen"
)

type Typedef struct {
	comment      jen.Code
	name         string
	concreteType jen.Code
	methods      map[string]*Method
	constructors map[string]*Function
}

func NewTypedef(comment jen.Code,
	name string,
	concreteType jen.Code,
	methods []*Method,
	constructors []*Function) *Typedef {
	t := &Typedef{
		comment:      comment,
		name:         name,
		concreteType: concreteType,
		methods:      make(map[string]*Method, len(methods)),
		constructors: make(map[string]*Function, len(constructors)),
	}
	for _, m := range methods {
		t.methods[m.Name()] = m
	}
	for _, c := range constructors {
		t.constructors[c.Name()] = c
	}
	return t
}

func (t *Typedef) Definition() jen.Code {
	def := jen.Empty().Add(
		t.comment,
	).Line().Type().Id(
		t.name,
	).Add(
		t.concreteType,
	)
	for _, c := range t.constructors {
		def = def.Line().Line().Add(c.Definition())
	}
	for _, m := range t.methods {
		def = def.Line().Line().Add(m.Definition())
	}
	return def
}

func (t *Typedef) Method(name string) *Method {
	return t.methods[name]
}

func (t *Typedef) Constructors(name string) *Function {
	return t.constructors[name]
}

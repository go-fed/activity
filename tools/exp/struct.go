package exp

import (
	"github.com/dave/jennifer/jen"
)

type Struct struct {
	comment      jen.Code
	name         string
	methods      map[string]*Method
	constructors map[string]*Function
	members      []jen.Code
}

func NewStruct(comment jen.Code,
	name string,
	methods []*Method,
	constructors []*Function,
	members []jen.Code) *Struct {
	s := &Struct{
		comment:      comment,
		name:         name,
		methods:      make(map[string]*Method, len(methods)),
		constructors: make(map[string]*Function, len(constructors)),
		members:      members,
	}
	for _, m := range methods {
		s.methods[m.Name()] = m
	}
	for _, c := range constructors {
		s.constructors[c.Name()] = c
	}
	return s
}

func (s *Struct) Definition() jen.Code {
	comment := jen.Empty()
	if s.comment != nil {
		comment = jen.Empty().Add(s.comment).Line()
	}
	def := comment.Type().Id(s.name).Struct(
		joinSingleLine(s.members),
	)
	for _, c := range s.constructors {
		def = def.Line().Line().Add(c.Definition())
	}
	for _, m := range s.methods {
		def = def.Line().Line().Add(m.Definition())
	}
	return def
}

func (s *Struct) Method(name string) *Method {
	return s.methods[name]
}

func (s *Struct) Constructors(name string) *Function {
	return s.constructors[name]
}

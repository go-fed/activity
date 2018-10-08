package exp

import (
	"github.com/dave/jennifer/jen"
)

type memberType int

const (
	this = "this"
)

const (
	valueMember memberType = iota
	pointerMember
)

func This() string {
	return this
}

type Function struct {
	qual    *jen.Statement
	name    string
	params  []jen.Code
	ret     []jen.Code
	block   []jen.Code
	comment jen.Code
}

func NewCommentedFunction(pkg, name string,
	params, ret, block []jen.Code,
	comment jen.Code) *Function {
	return &Function{
		qual:    jen.Qual(pkg, name),
		name:    name,
		params:  params,
		ret:     ret,
		block:   block,
		comment: comment,
	}
}

func NewFunction(pkg, name string,
	params, ret, block []jen.Code) *Function {
	return &Function{
		qual:    jen.Qual(pkg, name),
		name:    name,
		params:  params,
		ret:     ret,
		block:   block,
		comment: nil,
	}
}

func (m Function) Definition() jen.Code {
	stmts := jen.Empty()
	if m.comment != nil {
		stmts = jen.Empty().Add(m.comment).Line()
	}
	return stmts.Add(jen.Func().Id(m.name).Params(
		m.params...,
	).Params(
		m.ret...,
	).Block(
		m.block...,
	))
}

func (m Function) Call(params ...jen.Code) jen.Code {
	return m.qual.Call(params...)
}

func (m Function) Name() string {
	return m.name
}

type Method struct {
	member     memberType
	structName string
	function   *Function
}

func NewCommentedValueMethod(pkg, name, structName string,
	params, ret, block []jen.Code,
	comment jen.Code) *Method {
	return &Method{
		member:     valueMember,
		structName: structName,
		function: &Function{
			qual:    jen.Qual(pkg, name),
			name:    name,
			params:  params,
			ret:     ret,
			block:   block,
			comment: comment,
		},
	}
}

func NewValueMethod(pkg, name, structName string,
	params, ret, block []jen.Code) *Method {
	return &Method{
		member:     valueMember,
		structName: structName,
		function: &Function{
			qual:    jen.Qual(pkg, name),
			name:    name,
			params:  params,
			ret:     ret,
			block:   block,
			comment: nil,
		},
	}
}

func NewCommentedPointerMethod(pkg, name, structName string,
	params, ret, block []jen.Code,
	comment jen.Code) *Method {
	return &Method{
		member:     pointerMember,
		structName: structName,
		function: &Function{
			qual:    jen.Qual(pkg, name),
			name:    name,
			params:  params,
			ret:     ret,
			block:   block,
			comment: comment,
		},
	}
}

func NewPointerMethod(pkg, name, structName string,
	params, ret, block []jen.Code) *Method {
	return &Method{
		member:     pointerMember,
		structName: structName,
		function: &Function{
			qual:    jen.Qual(pkg, name),
			name:    name,
			params:  params,
			ret:     ret,
			block:   block,
			comment: nil,
		},
	}
}

func (m Method) Definition() jen.Code {
	comment := jen.Empty()
	if m.function.comment != nil {
		comment = jen.Empty().Add(m.function.comment).Line()
	}
	funcDef := jen.Empty()
	switch m.member {
	case pointerMember:
		funcDef = jen.Func().Params(
			jen.Id(This()).Op("*").Id(m.structName),
		)
	case valueMember:
		funcDef = jen.Func().Params(
			jen.Id(This()).Id(m.structName),
		)
	default:
		panic("unhandled method memberType")
	}
	return comment.Add(funcDef.Id(
		m.function.name,
	).Params(
		m.function.params...,
	).Params(
		m.function.ret...,
	).Block(
		m.function.block...,
	))
}

func (m Method) Call(on string, params ...jen.Code) jen.Code {
	return jen.Id(on).Call(params...)
}

func (m Method) Name() string {
	return m.function.name
}

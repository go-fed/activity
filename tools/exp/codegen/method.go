package codegen

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

// This returns the string variable used by members to refer to themselves.
func This() string {
	return this
}

// Function represents a free function, not a method, for Go code to be
// generated.
type Function struct {
	qual    *jen.Statement
	name    string
	params  []jen.Code
	ret     []jen.Code
	block   []jen.Code
	comment jen.Code
}

// NewCommentedFunction creates a new function with a comment.
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

// NewFunction creates a new function without any comments.
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

// Definition generates the Go code required to define and implement this
// function.
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

// Call generates the Go code required to call this function, with qualifier if
// required.
func (m Function) Call(params ...jen.Code) jen.Code {
	return m.qual.Clone().Call(params...)
}

// Name returns the identifier of this function.
func (m Function) Name() string {
	return m.name
}

// Method represents a method on a type, not a free function, for Go code to be
// generated.
type Method struct {
	member     memberType
	structName string
	function   *Function
}

// NewCommentedValueMethod defines a commented method for the value of a type.
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

// NewValueMethod defines a method for the value of a type. It is not commented.
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

// NewCommentedPointerMethod defines a commented method for the pointer to a
// type.
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

// NewPointerMethod defines a method for the pointer to a type. It is not
// commented.
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

// Definition generates the Go code required to define and implement this
// method.
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

// Call generates the Go code required to call this method, with qualifier if
// required.
func (m Method) Call(on string, params ...jen.Code) jen.Code {
	return jen.Id(on).Call(params...)
}

// Name returns the identifier of this function.
func (m Method) Name() string {
	return m.function.name
}

// ToFunctionSignature obtains this method's FunctionSignature.
func (m Method) ToFunctionSignature() FunctionSignature {
	return FunctionSignature{
		Name:   m.Name(),
		Params: m.function.params,
		Ret:    m.function.ret,
	}
}

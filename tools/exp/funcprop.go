package exp

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

// FunctionalPropertyGenerator produces Go code for properties that can have
// only one value. The resulting property is a struct type that can have one
// value that could be from multiple Kinds of values. If there is only one
// allowed Kind, then a smaller API is generated as a special case.
type FunctionalPropertyGenerator struct {
	PropertyGenerator
}

// Definition produces the Go Struct code definition, which can generate its Go
// implementations.
func (p *FunctionalPropertyGenerator) Definition() *Struct {
	if len(p.Kinds) == 1 {
		return p.singleTypeDef()
	} else {
		return p.multiTypeDef()
	}
}

// funcs produces the methods needed for the functional property.
func (p *FunctionalPropertyGenerator) funcs() []*Method {
	kindIndexFns := make([]jen.Code, 0, len(p.Kinds))
	for i, _ := range p.Kinds {
		if len(p.Kinds) > 1 {
			kindIndexFns = append(kindIndexFns, jen.If(
				jen.Id(This()).Dot(p.isMethodName(i)).Call(),
			).Block(
				jen.Return(jen.Lit(i)),
			))
		} else {
			kindIndexFns = append(kindIndexFns, jen.If(
				jen.Id(This()).Dot(hasMethod).Call(),
			).Block(
				jen.Return(jen.Lit(i)),
			))
		}
	}
	return []*Method{
		NewCommentedValueMethod(
			p.packageName(),
			kindIndexMethod,
			p.structName(),
			/*params=*/ nil,
			[]jen.Code{jen.Int()},
			[]jen.Code{
				join(kindIndexFns),
				jen.Return(jen.Lit(-1)),
			},
			jen.Commentf("%s computes an arbitrary value for indexing this kind of value.", kindIndexMethod),
		),
	}
}

// serializationFuncs produces the Methods and Functions needed for a
// functional property to be serialized and deserialized to and from an
// encoding.
func (p *FunctionalPropertyGenerator) serializationFuncs() ([]*Method, []*Function) {
	serializeFns := jen.Empty()
	for i, kind := range p.Kinds {
		if i > 0 {
			serializeFns = serializeFns.Else()
		}
		if len(p.Kinds) == 1 {
			serializeFns = serializeFns.If(
				jen.Id(This()).Dot(hasMethod).Call(),
			)
		} else {
			serializeFns = serializeFns.If(
				jen.Id(This()).Dot(p.isMethodName(i)).Call(),
			)
		}
		serializeFns = serializeFns.Block(
			jen.Return(
				jen.Id(kind.SerializeFnName).Call(
					jen.Id(This()).Dot(p.getFnName(i)).Call(),
				),
			),
		)
	}
	serialize := []*Method{
		NewCommentedValueMethod(
			p.packageName(),
			p.serializeFnName(),
			p.structName(),
			/*params=*/ nil,
			[]jen.Code{jen.Interface(), jen.Error()},
			[]jen.Code{serializeFns, jen.Return(
				jen.Nil(),
				jen.Nil(),
			)},
			jen.Commentf("%s converts this into an interface representation suitable for marshalling into a text or binary format.", p.serializeFnName()),
		),
	}
	deserializeFns := jen.Empty()
	for i, kind := range p.Kinds {
		if i > 0 {
			deserializeFns = deserializeFns.Else()
		}
		values := jen.Dict{
			jen.Id(p.memberName(i)): jen.Id("v"),
		}
		if !kind.Nilable {
			values[jen.Id(p.hasMemberName(i))] = jen.True()
		}
		deserializeFns = deserializeFns.If(
			jen.List(
				jen.Id("v"),
				jen.Id("handled"),
				jen.Err(),
			).Op(":=").Id(kind.DeserializeFnName).Call(
				jen.Id("i"),
			),
			jen.Id("handled"),
		).Block(
			jen.Id(This()).Op(":=").Op("&").Id(p.structName()).Values(
				values,
			),
			jen.Return(
				jen.Id(This()),
				jen.Err(),
			),
		)
	}
	var deserialize []*Function
	if p.asIterator {
		deserialize = append(deserialize,
			NewCommentedFunction(
				p.packageName(),
				p.deserializeFnName(),
				[]jen.Code{jen.Id("i").Interface()},
				[]jen.Code{jen.Op("*").Id(p.structName()), jen.Error()},
				[]jen.Code{
					deserializeFns,
					jen.Return(
						jen.Nil(),
						jen.Nil(),
					),
				},
				jen.Commentf("%s creates an iterator from an element that has been unmarshalled from a text or binary format.", p.deserializeFnName()),
			))
	} else {
		deserialize = append(deserialize,
			NewCommentedFunction(
				p.packageName(),
				p.deserializeFnName(),
				[]jen.Code{jen.Id("m").Map(jen.String()).Interface()},
				[]jen.Code{jen.Op("*").Id(p.structName()), jen.Error()},
				[]jen.Code{
					jen.If(
						jen.List(
							jen.Id("i"),
							jen.Id("ok"),
						).Op(":=").Id("m").Index(
							jen.Lit(p.propertyName()),
						),
						jen.Id("ok"),
					).Block(
						deserializeFns,
					),
					jen.Return(
						jen.Nil(),
						jen.Nil(),
					),
				},
				jen.Commentf("%s creates a %q property from an interface representation that has been unmarshalled from a text or binary format.", p.deserializeFnName(), p.propertyName()),
			))
	}
	return serialize, deserialize
}

// singleTypeDef generates a special-case simplified API for a functional
// property that can only be a single Kind of value.
func (p *FunctionalPropertyGenerator) singleTypeDef() *Struct {
	var comment jen.Code
	var kindMembers []jen.Code
	if p.Kinds[0].Nilable {
		comment = jen.Commentf("%s is the functional property %q. It is permitted to be a single nilable value type.", p.structName(), p.propertyName())
		if p.asIterator {
			comment = jen.Commentf("%s is an iterator for a property. It is permitted to be a single nilable value type.", p.structName())
		}
		kindMembers = []jen.Code{jen.Id(p.memberName(0)).Id(p.Kinds[0].ConcreteKind)}
	} else {
		comment = jen.Commentf("%s is the functional property %q. It is permitted to be a single default-valued value type.", p.structName(), p.propertyName())
		if p.asIterator {
			comment = jen.Commentf("%s is an iterator for a property. It is permitted to be a single default-valued value type.", p.structName())
		}
		kindMembers = []jen.Code{
			jen.Id(p.memberName(0)).Id(p.Kinds[0].ConcreteKind),
			jen.Id(p.hasMemberName(0)).Bool(),
		}
	}
	methods, funcs := p.serializationFuncs()
	methods = append(methods, p.singleTypeFuncs()...)
	methods = append(methods, p.funcs()...)
	methods = append(methods, p.commonMethods()...)
	return NewStruct(comment,
		p.structName(),
		methods,
		funcs,
		kindMembers)
}

// singleTypeFuncs generates the special-case simplified methods for a
// functional property with exactly one Kind of value.
func (p *FunctionalPropertyGenerator) singleTypeFuncs() []*Method {
	var methods []*Method
	// Has Method
	hasComment := jen.Commentf("%s returns true if this property is set.", hasMethod)
	if p.Kinds[0].Nilable {
		methods = append(methods, NewCommentedValueMethod(
			p.packageName(),
			hasMethod,
			p.structName(),
			/*params=*/ nil,
			[]jen.Code{jen.Bool()},
			[]jen.Code{jen.Return(jen.Id(This()).Dot(p.memberName(0)).Op("!=").Nil())},
			hasComment,
		))
	} else {
		methods = append(methods, NewCommentedValueMethod(
			p.packageName(),
			hasMethod,
			p.structName(),
			/*params=*/ nil,
			[]jen.Code{jen.Bool()},
			[]jen.Code{jen.Return(jen.Id(This()).Dot(p.hasMemberName(0)))},
			hasComment,
		))
	}
	// Get Method
	getComment := jen.Commentf("%s returns the value of this property. When %s returns false, %s will return any arbitrary value.", getMethod, hasMethod, getMethod)
	methods = append(methods, NewCommentedValueMethod(
		p.packageName(),
		p.getFnName(0),
		p.structName(),
		/*params=*/ nil,
		[]jen.Code{jen.Id(p.Kinds[0].ConcreteKind)},
		[]jen.Code{jen.Return(jen.Id(This()).Dot(p.memberName(0)))},
		getComment,
	))
	// Set Method
	setComment := jen.Commentf("%s sets the value of this property. Calling %s afterwards will return true.", setMethod, hasMethod)
	if p.Kinds[0].Nilable {
		methods = append(methods, NewCommentedPointerMethod(
			p.packageName(),
			setMethod,
			p.structName(),
			[]jen.Code{jen.Id("v").Id(p.Kinds[0].ConcreteKind)},
			/*ret=*/ nil,
			[]jen.Code{jen.Id(This()).Dot(p.memberName(0)).Op("=").Id("v")},
			setComment,
		))
	} else {
		methods = append(methods, NewCommentedPointerMethod(
			p.packageName(),
			setMethod,
			p.structName(),
			[]jen.Code{jen.Id("v").Id(p.Kinds[0].ConcreteKind)},
			/*ret=*/ nil,
			[]jen.Code{
				jen.Id(This()).Dot(p.memberName(0)).Op("=").Id("v"),
				jen.Id(This()).Dot(p.hasMemberName(0)).Op("=").True(),
			},
			setComment,
		))
	}
	// Clear Method
	clearComment := jen.Commentf("%s ensures no value of this property is set. Calling %s afterwards will return false.", p.clearMethodName(), hasMethod).Line()
	if p.Kinds[0].Nilable {
		methods = append(methods, NewCommentedPointerMethod(
			p.packageName(),
			p.clearMethodName(),
			p.structName(),
			/*params=*/ nil,
			/*ret=*/ nil,
			[]jen.Code{jen.Id(This()).Dot(p.memberName(0)).Op("=").Nil()},
			clearComment,
		))
	} else {
		methods = append(methods, NewCommentedPointerMethod(
			p.packageName(),
			p.clearMethodName(),
			p.structName(),
			/*params=*/ nil,
			/*ret=*/ nil,
			[]jen.Code{jen.Id(This()).Dot(p.hasMemberName(0)).Op("=").False()},
			clearComment,
		))
	}
	return methods
}

// multiTypeDef generates an API for a functional property that can be multiple
// Kinds of value.
func (p *FunctionalPropertyGenerator) multiTypeDef() *Struct {
	kindMembers := make([]jen.Code, 0, len(p.Kinds))
	for i, kind := range p.Kinds {
		if kind.Nilable {
			kindMembers = append(kindMembers, jen.Id(p.memberName(i)).Id(p.Kinds[i].ConcreteKind))
		} else {
			kindMembers = append(kindMembers, jen.Id(p.memberName(i)).Id(p.Kinds[i].ConcreteKind))
			kindMembers = append(kindMembers, jen.Id(p.hasMemberName(i)).Bool())
		}
	}
	explanation := jen.Commentf(
		"At most, one type of value can be present, or none at all. Setting a value will",
	).Line().Commentf(
		"clear the other types of values so that only one of the 'Is' methods will return",
	).Line().Commentf(
		"true.",
	).Line().Comment("").Line().Commentf(
		"It is possible to clear all values, so that this property is empty.",
	)
	comment := jen.Commentf(
		"%s is the functional property %q. It is permitted to be one of multiple value types.", p.structName(), p.propertyName(),
	).Line().Comment("").Line().Add(explanation)
	if p.asIterator {
		comment = jen.Commentf(
			"%s is an iterator for a property. It is permitted to be one of multiple value types.", p.structName(),
		).Line().Comment("").Line().Add(explanation)
	}
	methods, funcs := p.serializationFuncs()
	methods = append(methods, p.multiTypeFuncs()...)
	methods = append(methods, p.funcs()...)
	methods = append(methods, p.commonMethods()...)
	return NewStruct(comment,
		p.structName(),
		methods,
		funcs,
		kindMembers)
}

// multiTypeFuncs generates the methods for a functional property with more than
// one Kind of value.
func (p *FunctionalPropertyGenerator) multiTypeFuncs() []*Method {
	var methods []*Method
	// HasAny Method
	hasAnyMethodName := fmt.Sprintf("%sAny", hasMethod)
	isLine := make([]jen.Code, len(p.Kinds))
	for i := range p.Kinds {
		or := jen.Empty()
		if i != len(p.Kinds)-1 {
			or = jen.Op("||")
		}
		isLine[i] = jen.Id(This()).Dot(p.isMethodName(i)).Parens(nil).Add(or)
	}
	methods = append(methods, NewCommentedPointerMethod(
		p.packageName(),
		hasAnyMethodName,
		p.structName(),
		/*params=*/ nil,
		[]jen.Code{jen.Bool()},
		[]jen.Code{jen.Return(join(isLine))},
		jen.Commentf(
			"%s returns true if any of the different values is set.", hasAnyMethodName,
		),
	))
	// Clear Method
	clearLine := make([]jen.Code, len(p.Kinds))
	for i, kind := range p.Kinds {
		if kind.Nilable {
			clearLine[i] = jen.Id(This()).Dot(p.memberName(i)).Op("=").Nil()
		} else {
			clearLine[i] = jen.Id(This()).Dot(p.hasMemberName(i)).Op("=").False()
		}
	}
	methods = append(methods, NewCommentedPointerMethod(
		p.packageName(),
		p.clearMethodName(),
		p.structName(),
		/*params=*/ nil,
		/*ret=*/ nil,
		clearLine,
		jen.Commentf(
			"%s ensures no value of this property is set. Calling %s or any of the 'Is' methods afterwards will return false.", p.clearMethodName(), hasAnyMethodName,
		),
	))
	// Is Method
	for i, kind := range p.Kinds {
		isComment := jen.Commentf("%s returns true if this property has a type of value of %q.", p.isMethodName(i), kind.ConcreteKind)
		if kind.Nilable {
			methods = append(methods, NewCommentedValueMethod(
				p.packageName(),
				p.isMethodName(i),
				p.structName(),
				/*params=*/ nil,
				[]jen.Code{jen.Bool()},
				[]jen.Code{jen.Return(jen.Id(This()).Dot(p.memberName(i)).Op("!=").Nil())},
				isComment,
			))
		} else {
			methods = append(methods, NewCommentedValueMethod(
				p.packageName(),
				p.isMethodName(i),
				p.structName(),
				/*params=*/ nil,
				[]jen.Code{jen.Bool()},
				[]jen.Code{jen.Return(jen.Id(This()).Dot(p.hasMemberName(i)))},
				isComment,
			))
		}
	}
	// Set Method
	for i, kind := range p.Kinds {
		setMethodName := fmt.Sprintf("%s%s", setMethod, p.kindCamelName(i))
		setComment := jen.Commentf("%s sets the value of this property. Calling %s afterwards returns true.", setMethodName, p.isMethodName(i))
		if kind.Nilable {
			methods = append(methods, NewCommentedPointerMethod(
				p.packageName(),
				setMethodName,
				p.structName(),
				[]jen.Code{jen.Id("v").Id(kind.ConcreteKind)},
				/*ret=*/ nil,
				[]jen.Code{
					jen.Id(This()).Dot(p.clearMethodName()).Call(),
					jen.Id(This()).Dot(p.memberName(i)).Op("=").Id("v"),
				},
				setComment,
			))
		} else {
			methods = append(methods, NewCommentedPointerMethod(
				p.packageName(),
				setMethodName,
				p.structName(),
				[]jen.Code{jen.Id("v").Id(kind.ConcreteKind)},
				/*ret=*/ nil,
				[]jen.Code{
					jen.Id(This()).Dot(p.clearMethodName()).Call(),
					jen.Id(This()).Dot(p.memberName(i)).Op("=").Id("v"),
					jen.Id(This()).Dot(p.hasMemberName(i)).Op("=").True(),
				},
				setComment,
			))
		}
	}
	// Get Method
	for i, kind := range p.Kinds {
		getComment := jen.Commentf("%s returns the value of this property. When %s returns false, %s will return an arbitrary value.", p.getFnName(i), p.isMethodName(i), p.getFnName(i))
		methods = append(methods, NewCommentedValueMethod(
			p.packageName(),
			p.getFnName(i),
			p.structName(),
			/*params=*/ nil,
			[]jen.Code{jen.Id(kind.ConcreteKind)},
			[]jen.Code{jen.Return(jen.Id(This()).Dot(p.memberName(i)))},
			getComment,
		))
	}
	return methods
}

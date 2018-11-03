package props

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/go-fed/activity/tools/exp/codegen"
)

// FunctionalPropertyGenerator produces Go code for properties that can have
// only one value. The resulting property is a struct type that can have one
// value that could be from multiple Kinds of values. If there is only one
// allowed Kind, then a smaller API is generated as a special case.
type FunctionalPropertyGenerator struct {
	PropertyGenerator
}

// isSingleTypeDef determines whether a special-case API can be generated for
// one allowed Kind.
func (p *FunctionalPropertyGenerator) isSingleTypeDef() bool {
	return len(p.Kinds) == 1
}

// Definition produces the Go Struct code definition, which can generate its Go
// implementations.
func (p *FunctionalPropertyGenerator) Definition() *codegen.Struct {
	if p.isSingleTypeDef() {
		return p.singleTypeDef()
	} else {
		return p.multiTypeDef()
	}
}

// clearNonLanguageMapMembers generates the code required to clear all values,
// including unknown values, from this property except for the natural language
// map. If this property can handle a natural language map, then it is up to the
// calling code to determine whether to set the 'langMapMember' to nil.
func (p *FunctionalPropertyGenerator) clearNonLanguageMapMembers() []jen.Code {
	if p.isSingleTypeDef() {
		return p.singleTypeClearNonLanguageMapMembers()
	} else {
		return p.multiTypeClearNonLanguageMapMembers()
	}
}

// singleTypeClearNonLanguageMapMembers generates code to clear all members for
// the special case single-Kind property.
func (p *FunctionalPropertyGenerator) singleTypeClearNonLanguageMapMembers() []jen.Code {
	clearCode := []jen.Code{
		jen.Id(codegen.This()).Dot(unknownMemberName).Op("=").Nil(),
	}
	if p.Kinds[0].Nilable {
		clearCode = append(clearCode, jen.Id(codegen.This()).Dot(p.memberName(0)).Op("=").Nil())
	} else {
		clearCode = append(clearCode, jen.Id(codegen.This()).Dot(p.hasMemberName(0)).Op("=").False())
	}
	return clearCode
}

// multiTypeClearNonLanguageMapMembers generates code to clear all members for
// a property with multiple Kinds.
func (p *FunctionalPropertyGenerator) multiTypeClearNonLanguageMapMembers() []jen.Code {
	clearLine := make([]jen.Code, len(p.Kinds)+2) // +2 for the unknown, and maybe language map
	for i, kind := range p.Kinds {
		if kind.Nilable {
			clearLine[i] = jen.Id(codegen.This()).Dot(p.memberName(i)).Op("=").Nil()
		} else {
			clearLine[i] = jen.Id(codegen.This()).Dot(p.hasMemberName(i)).Op("=").False()
		}
	}
	clearLine = append(clearLine, jen.Id(codegen.This()).Dot(unknownMemberName).Op("=").Nil())
	return clearLine
}

// funcs produces the methods needed for the functional property.
func (p *FunctionalPropertyGenerator) funcs() []*codegen.Method {
	kindIndexFns := make([]jen.Code, 0, len(p.Kinds))
	for i, _ := range p.Kinds {
		if p.isSingleTypeDef() {
			kindIndexFns = append(kindIndexFns, jen.If(
				jen.Id(codegen.This()).Dot(hasMethod).Call(),
			).Block(
				jen.Return(jen.Lit(i)),
			))
		} else {
			kindIndexFns = append(kindIndexFns, jen.If(
				jen.Id(codegen.This()).Dot(p.isMethodName(i)).Call(),
			).Block(
				jen.Return(jen.Lit(i)),
			))
		}
	}
	methods := []*codegen.Method{
		codegen.NewCommentedValueMethod(
			p.packageName(),
			kindIndexMethod,
			p.StructName(),
			/*params=*/ nil,
			[]jen.Code{jen.Int()},
			[]jen.Code{
				join(kindIndexFns),
				jen.Return(jen.Lit(-1)),
			},
			jen.Commentf("%s computes an arbitrary value for indexing this kind of value.", kindIndexMethod),
		),
	}
	if p.HasNaturalLanguageMap {
		// IsLanguageMap Method
		methods = append(methods,
			codegen.NewCommentedValueMethod(
				p.packageName(),
				isLanguageMapMethod,
				p.StructName(),
				/*params=*/ nil,
				[]jen.Code{jen.Bool()},
				[]jen.Code{
					jen.Return(jen.Id(codegen.This()).Dot(langMapMember).Op("!=").Nil()),
				},
				jen.Commentf(
					"%s determines if this property is represented by a natural language map.",
					isLanguageMapMethod,
				).Line().Commentf("").Line().Commentf(
					"When true, use %s, %s, and %s methods to access and mutate the natural language map.",
					hasLanguageMethod,
					getLanguageMethod,
					setLanguageMethod,
				).Line().Commentf(
					"The %s method can be used to clear the natural language map.",
					p.clearMethodName(),
				).Line().Commentf("").Line().Commentf(
					"Note that this method is only used for natural language representations, and does not determine the presence nor absence of other values for this property.",
				)))
		// HasLanguage Method
		methods = append(methods,
			codegen.NewCommentedValueMethod(
				p.packageName(),
				hasLanguageMethod,
				p.StructName(),
				[]jen.Code{jen.Id("bcp47").String()},
				[]jen.Code{jen.Bool()},
				[]jen.Code{
					jen.If(
						jen.Id(codegen.This()).Dot(langMapMember).Op("==").Nil(),
					).Block(
						jen.Return(jen.False()),
					).Else().Block(
						jen.List(
							jen.Id("_"),
							jen.Id("ok"),
						).Op(":=").Id(codegen.This()).Dot(langMapMember).Index(
							jen.Id("bcp47"),
						),
						jen.Return(jen.Id("ok")),
					),
				},
				jen.Commentf(
					"%s returns true if the natural language map has an entry for the specified BCP47 language code.",
					hasLanguageMethod,
				),
			))
		// GetLanguage Method
		methods = append(methods,
			codegen.NewCommentedValueMethod(
				p.packageName(),
				getLanguageMethod,
				p.StructName(),
				[]jen.Code{jen.Id("bcp47").String()},
				[]jen.Code{jen.String()},
				[]jen.Code{
					jen.If(
						jen.Id(codegen.This()).Dot(langMapMember).Op("==").Nil(),
					).Block(
						jen.Return(jen.Lit("")),
					).Else().If(
						jen.List(
							jen.Id("v"),
							jen.Id("ok"),
						).Op(":=").Id(codegen.This()).Dot(langMapMember).Index(
							jen.Id("bcp47"),
						),
						jen.Id("ok"),
					).Block(
						jen.Return(jen.Id("v")),
					).Else().Block(
						jen.Return(jen.Lit("")),
					),
				},
				jen.Commentf(
					"%s returns the value for the specified BCP47 language code, or an empty string if it is either not a language map or no value is present.",
					getLanguageMethod,
				),
			))
		// SetLanguage Method
		methods = append(methods,
			codegen.NewCommentedPointerMethod(
				p.packageName(),
				setLanguageMethod,
				p.StructName(),
				[]jen.Code{
					jen.Id("bcp47"),
					jen.Id("value").String(),
				},
				/*ret=*/ nil,
				append(p.clearNonLanguageMapMembers(),
					[]jen.Code{
						jen.If(
							jen.Id(codegen.This()).Dot(langMapMember).Op("==").Nil(),
						).Block(
							jen.Id(codegen.This()).Dot(langMapMember).Op("=").Make(
								jen.Map(jen.String()).String(),
							),
						),
						jen.Id(codegen.This()).Dot(langMapMember).Index(
							jen.Id("bcp47"),
						).Op("=").Id("value"),
					}...,
				),
				jen.Commentf(
					"%s sets the value for the specified BCP47 language code.",
					setLanguageMethod,
				),
			))
	}
	return methods
}

// serializationFuncs produces the Methods and Functions needed for a
// functional property to be serialized and deserialized to and from an
// encoding.
func (p *FunctionalPropertyGenerator) serializationFuncs() ([]*codegen.Method, []*codegen.Function) {
	serializeFns := jen.Empty()
	for i, kind := range p.Kinds {
		if i > 0 {
			serializeFns = serializeFns.Else()
		}
		if p.isSingleTypeDef() {
			serializeFns = serializeFns.If(
				jen.Id(codegen.This()).Dot(hasMethod).Call(),
			)
		} else {
			serializeFns = serializeFns.If(
				jen.Id(codegen.This()).Dot(p.isMethodName(i)).Call(),
			)
		}
		serializeFns = serializeFns.Block(
			jen.Return(
				kind.SerializeFn.Call(
					jen.Id(codegen.This()).Dot(p.getFnName(i)).Call(),
				),
			),
		)
	}
	serialize := []*codegen.Method{
		codegen.NewCommentedValueMethod(
			p.packageName(),
			p.serializeFnName(),
			p.StructName(),
			/*params=*/ nil,
			[]jen.Code{jen.Interface(), jen.Error()},
			[]jen.Code{serializeFns, jen.Return(
				jen.Id(codegen.This()).Dot(unknownMemberName),
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
			).Op(":=").Add(kind.DeserializeFn.Call(
				jen.Id("i"),
			)),
			jen.Id("handled"),
		).Block(
			jen.Id(codegen.This()).Op(":=").Op("&").Id(p.StructName()).Values(
				values,
			),
			jen.Return(
				jen.Id(codegen.This()),
				jen.Err(),
			),
		)
	}
	var deserialize []*codegen.Function
	if p.asIterator {
		deserialize = append(deserialize,
			codegen.NewCommentedFunction(
				p.packageName(),
				p.deserializeFnName(),
				[]jen.Code{jen.Id("i").Interface()},
				[]jen.Code{jen.Op("*").Id(p.StructName()), jen.Error()},
				[]jen.Code{
					deserializeFns.Add(p.unknownDeserializeCode()),
					jen.Return(
						jen.Nil(),
						jen.Nil(),
					),
				},
				jen.Commentf("%s creates an iterator from an element that has been unmarshalled from a text or binary format.", p.deserializeFnName()),
			))
	} else {
		deserialize = append(deserialize,
			codegen.NewCommentedFunction(
				p.packageName(),
				p.deserializeFnName(),
				[]jen.Code{jen.Id("m").Map(jen.String()).Interface()},
				[]jen.Code{jen.Op("*").Id(p.StructName()), jen.Error()},
				[]jen.Code{
					jen.If(
						jen.List(
							jen.Id("i"),
							jen.Id("ok"),
						).Op(":=").Id("m").Index(
							jen.Lit(p.PropertyName()),
						),
						jen.Id("ok"),
					).Block(
						deserializeFns.Add(p.unknownDeserializeCode()),
					),
					jen.Return(
						jen.Nil(),
						jen.Nil(),
					),
				},
				jen.Commentf("%s creates a %q property from an interface representation that has been unmarshalled from a text or binary format.", p.deserializeFnName(), p.PropertyName()),
			))
	}
	return serialize, deserialize
}

// singleTypeDef generates a special-case simplified API for a functional
// property that can only be a single Kind of value.
func (p *FunctionalPropertyGenerator) singleTypeDef() *codegen.Struct {
	var comment jen.Code
	var kindMembers []jen.Code
	if p.Kinds[0].Nilable {
		comment = jen.Commentf("%s is the functional property %q. It is permitted to be a single nilable value type.", p.StructName(), p.PropertyName())
		if p.asIterator {
			comment = jen.Commentf("%s is an iterator for a property. It is permitted to be a single nilable value type.", p.StructName())
		}
		kindMembers = []jen.Code{
			jen.Id(p.memberName(0)).Id(p.Kinds[0].ConcreteKind),
		}
	} else {
		comment = jen.Commentf("%s is the functional property %q. It is permitted to be a single default-valued value type.", p.StructName(), p.PropertyName())
		if p.asIterator {
			comment = jen.Commentf("%s is an iterator for a property. It is permitted to be a single default-valued value type.", p.StructName())
		}
		kindMembers = []jen.Code{
			jen.Id(p.memberName(0)).Id(p.Kinds[0].ConcreteKind),
			jen.Id(p.hasMemberName(0)).Bool(),
		}
	}
	kindMembers = append(kindMembers, p.unknownMemberDef())
	if p.HasNaturalLanguageMap {
		kindMembers = append(kindMembers, jen.Id(langMapMember).Map(jen.String()).String())
	}
	methods, funcs := p.serializationFuncs()
	methods = append(methods, p.singleTypeFuncs()...)
	methods = append(methods, p.funcs()...)
	methods = append(methods, p.commonMethods()...)
	return codegen.NewStruct(comment,
		p.StructName(),
		methods,
		funcs,
		kindMembers)
}

// singleTypeFuncs generates the special-case simplified methods for a
// functional property with exactly one Kind of value.
func (p *FunctionalPropertyGenerator) singleTypeFuncs() []*codegen.Method {
	var methods []*codegen.Method
	// Has Method
	hasComment := jen.Commentf("%s returns true if this property is set.", hasMethod)
	if p.HasNaturalLanguageMap {
		hasComment = jen.Commentf(
			"%s returns true if this property is set and is not a natural language map.",
			hasMethod,
		).Line().Commentf("").Line().Commentf(
			"When true, the %s and %s methods may be used to access and set this property.",
			getMethod,
			p.setFnName(0),
		).Line().Commentf("").Line().Commentf(
			"To determine if the property was set as a natural language map, use the %s method instead.",
			isLanguageMapMethod,
		)
	}
	if p.Kinds[0].Nilable {
		methods = append(methods, codegen.NewCommentedValueMethod(
			p.packageName(),
			hasMethod,
			p.StructName(),
			/*params=*/ nil,
			[]jen.Code{jen.Bool()},
			[]jen.Code{jen.Return(jen.Id(codegen.This()).Dot(p.memberName(0)).Op("!=").Nil())},
			hasComment,
		))
	} else {
		methods = append(methods, codegen.NewCommentedValueMethod(
			p.packageName(),
			hasMethod,
			p.StructName(),
			/*params=*/ nil,
			[]jen.Code{jen.Bool()},
			[]jen.Code{jen.Return(jen.Id(codegen.This()).Dot(p.hasMemberName(0)))},
			hasComment,
		))
	}
	// Get Method
	getComment := jen.Commentf("%s returns the value of this property. When %s returns false, %s will return any arbitrary value.", getMethod, hasMethod, getMethod)
	methods = append(methods, codegen.NewCommentedValueMethod(
		p.packageName(),
		p.getFnName(0),
		p.StructName(),
		/*params=*/ nil,
		[]jen.Code{jen.Id(p.Kinds[0].ConcreteKind)},
		[]jen.Code{jen.Return(jen.Id(codegen.This()).Dot(p.memberName(0)))},
		getComment,
	))
	// Set Method
	setComment := jen.Commentf("%s sets the value of this property. Calling %s afterwards will return true.", p.setFnName(0), hasMethod)
	if p.HasNaturalLanguageMap {
		setComment = jen.Commentf(
			"%s sets the value of this property and clears the natural language map.",
			p.setFnName(0),
		).Line().Commentf("").Line().Commentf(
			"Calling %s afterwards will return true. Calling %s afterwards returns false.",
			hasMethod,
			isLanguageMapMethod,
		)
	}
	if p.Kinds[0].Nilable {
		methods = append(methods, codegen.NewCommentedPointerMethod(
			p.packageName(),
			p.setFnName(0),
			p.StructName(),
			[]jen.Code{jen.Id("v").Id(p.Kinds[0].ConcreteKind)},
			/*ret=*/ nil,
			[]jen.Code{
				jen.Id(codegen.This()).Dot(p.clearMethodName()).Call(),
				jen.Id(codegen.This()).Dot(p.memberName(0)).Op("=").Id("v"),
			},
			setComment,
		))
	} else {
		methods = append(methods, codegen.NewCommentedPointerMethod(
			p.packageName(),
			p.setFnName(0),
			p.StructName(),
			[]jen.Code{jen.Id("v").Id(p.Kinds[0].ConcreteKind)},
			/*ret=*/ nil,
			[]jen.Code{
				jen.Id(codegen.This()).Dot(p.clearMethodName()).Call(),
				jen.Id(codegen.This()).Dot(p.memberName(0)).Op("=").Id("v"),
				jen.Id(codegen.This()).Dot(p.hasMemberName(0)).Op("=").True(),
			},
			setComment,
		))
	}
	// Clear Method
	clearComment := jen.Commentf("%s ensures no value of this property is set. Calling %s afterwards will return false.", p.clearMethodName(), hasMethod)
	clearCode := p.singleTypeClearNonLanguageMapMembers()
	if p.HasNaturalLanguageMap {
		clearComment = jen.Commentf(
			"%s ensures no value and no language map for this property is set.",
			p.clearMethodName(),
		).Line().Commentf("").Line().Commentf(
			"Calling %s or %s afterwards will return false.",
			hasMethod,
			isLanguageMapMethod,
		)
		clearCode = append(clearCode, jen.Id(codegen.This()).Dot(langMapMember).Op("=").Nil())
	}
	if p.Kinds[0].Nilable {
		methods = append(methods, codegen.NewCommentedPointerMethod(
			p.packageName(),
			p.clearMethodName(),
			p.StructName(),
			/*params=*/ nil,
			/*ret=*/ nil,
			clearCode,
			clearComment,
		))
	} else {
		methods = append(methods, codegen.NewCommentedPointerMethod(
			p.packageName(),
			p.clearMethodName(),
			p.StructName(),
			/*params=*/ nil,
			/*ret=*/ nil,
			clearCode,
			clearComment,
		))
	}
	return methods
}

// multiTypeDef generates an API for a functional property that can be multiple
// Kinds of value.
func (p *FunctionalPropertyGenerator) multiTypeDef() *codegen.Struct {
	kindMembers := make([]jen.Code, 0, len(p.Kinds))
	for i, kind := range p.Kinds {
		if kind.Nilable {
			kindMembers = append(kindMembers, jen.Id(p.memberName(i)).Id(p.Kinds[i].ConcreteKind))
		} else {
			kindMembers = append(kindMembers, jen.Id(p.memberName(i)).Id(p.Kinds[i].ConcreteKind))
			kindMembers = append(kindMembers, jen.Id(p.hasMemberName(i)).Bool())
		}
	}
	kindMembers = append(kindMembers, p.unknownMemberDef())
	if p.HasNaturalLanguageMap {
		kindMembers = append(kindMembers, jen.Id(langMapMember).Map(jen.String()).String())
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
		"%s is the functional property %q. It is permitted to be one of multiple value types.", p.StructName(), p.PropertyName(),
	).Line().Comment("").Line().Add(explanation)
	if p.asIterator {
		comment = jen.Commentf(
			"%s is an iterator for a property. It is permitted to be one of multiple value types.", p.StructName(),
		).Line().Comment("").Line().Add(explanation)
	}
	methods, funcs := p.serializationFuncs()
	methods = append(methods, p.multiTypeFuncs()...)
	methods = append(methods, p.funcs()...)
	methods = append(methods, p.commonMethods()...)
	return codegen.NewStruct(comment,
		p.StructName(),
		methods,
		funcs,
		kindMembers)
}

// multiTypeFuncs generates the methods for a functional property with more than
// one Kind of value.
func (p *FunctionalPropertyGenerator) multiTypeFuncs() []*codegen.Method {
	var methods []*codegen.Method
	// HasAny Method
	hasAnyMethodName := fmt.Sprintf("%sAny", hasMethod)
	isLine := make([]jen.Code, len(p.Kinds))
	for i := range p.Kinds {
		or := jen.Empty()
		if i != len(p.Kinds)-1 {
			or = jen.Op("||")
		}
		isLine[i] = jen.Id(codegen.This()).Dot(p.isMethodName(i)).Parens(nil).Add(or)
	}
	hasAnyComment := jen.Commentf(
		"%s returns true if any of the different values is set.", hasAnyMethodName,
	)
	if p.HasNaturalLanguageMap {
		hasAnyComment = jen.Commentf(
			"%s returns true if any of the values are set, except for the natural language map",
			hasAnyMethodName,
		).Line().Commentf("").Line().Commentf(
			"When true, the specific has, getter, and setter methods may be used to determine what kind of value there is to access and set this property.",
		).Line().Commentf("").Line().Commentf(
			"To determine if the property was set as a natural language map, use the %s method instead.",
			isLanguageMapMethod,
		)
	}
	methods = append(methods, codegen.NewCommentedPointerMethod(
		p.packageName(),
		hasAnyMethodName,
		p.StructName(),
		/*params=*/ nil,
		[]jen.Code{jen.Bool()},
		[]jen.Code{jen.Return(join(isLine))},
		hasAnyComment,
	))
	// Clear Method
	clearComment := jen.Commentf(
		"%s ensures no value of this property is set. Calling %s or any of the 'Is' methods afterwards will return false.", p.clearMethodName(), hasAnyMethodName,
	)
	clearLine := p.multiTypeClearNonLanguageMapMembers()
	if p.HasNaturalLanguageMap {
		clearComment = jen.Commentf(
			"%s ensures no value and no language map for this property is set.",
			p.clearMethodName(),
		).Line().Commentf("").Line().Commentf(
			"Calling %s or any of the 'Is' methods afterwards will return false.",
			hasAnyMethodName,
		)
		clearLine = append(clearLine, jen.Id(codegen.This()).Dot(langMapMember).Op("=").Nil())
	}
	methods = append(methods, codegen.NewCommentedPointerMethod(
		p.packageName(),
		p.clearMethodName(),
		p.StructName(),
		/*params=*/ nil,
		/*ret=*/ nil,
		clearLine,
		clearComment,
	))
	// Is Method
	for i, kind := range p.Kinds {
		isComment := jen.Commentf("%s returns true if this property has a type of value of %q.", p.isMethodName(i), kind.ConcreteKind).Line().Commentf("").Line().Commentf(
			"When true, use the %s and %s methods to access and set this property.",
			p.getFnName(i),
			p.setFnName(i),
		)
		if p.HasNaturalLanguageMap {
			isComment = isComment.Line().Commentf("").Line().Commentf(
				"To determine if the property was set as a natural language map, use the %s method instead.",
				isLanguageMapMethod,
			)
		}
		if kind.Nilable {
			methods = append(methods, codegen.NewCommentedValueMethod(
				p.packageName(),
				p.isMethodName(i),
				p.StructName(),
				/*params=*/ nil,
				[]jen.Code{jen.Bool()},
				[]jen.Code{jen.Return(jen.Id(codegen.This()).Dot(p.memberName(i)).Op("!=").Nil())},
				isComment,
			))
		} else {
			methods = append(methods, codegen.NewCommentedValueMethod(
				p.packageName(),
				p.isMethodName(i),
				p.StructName(),
				/*params=*/ nil,
				[]jen.Code{jen.Bool()},
				[]jen.Code{jen.Return(jen.Id(codegen.This()).Dot(p.hasMemberName(i)))},
				isComment,
			))
		}
	}
	// Set Method
	for i, kind := range p.Kinds {
		setComment := jen.Commentf("%s sets the value of this property. Calling %s afterwards returns true.", p.setFnName(i), p.isMethodName(i))
		if p.HasNaturalLanguageMap {
			setComment = jen.Commentf(
				"%s sets the value of this property and clears the natural language map.",
				p.setFnName(i),
			).Line().Commentf("").Line().Commentf(
				"Calling %s afterwards will return true. Calling %s afterwards returns false.",
				p.isMethodName(i),
				isLanguageMapMethod,
			)
		}
		if kind.Nilable {
			methods = append(methods, codegen.NewCommentedPointerMethod(
				p.packageName(),
				p.setFnName(i),
				p.StructName(),
				[]jen.Code{jen.Id("v").Id(kind.ConcreteKind)},
				/*ret=*/ nil,
				[]jen.Code{
					jen.Id(codegen.This()).Dot(p.clearMethodName()).Call(),
					jen.Id(codegen.This()).Dot(p.memberName(i)).Op("=").Id("v"),
				},
				setComment,
			))
		} else {
			methods = append(methods, codegen.NewCommentedPointerMethod(
				p.packageName(),
				p.setFnName(i),
				p.StructName(),
				[]jen.Code{jen.Id("v").Id(kind.ConcreteKind)},
				/*ret=*/ nil,
				[]jen.Code{
					jen.Id(codegen.This()).Dot(p.clearMethodName()).Call(),
					jen.Id(codegen.This()).Dot(p.memberName(i)).Op("=").Id("v"),
					jen.Id(codegen.This()).Dot(p.hasMemberName(i)).Op("=").True(),
				},
				setComment,
			))
		}
	}
	// Get Method
	for i, kind := range p.Kinds {
		getComment := jen.Commentf("%s returns the value of this property. When %s returns false, %s will return an arbitrary value.", p.getFnName(i), p.isMethodName(i), p.getFnName(i))
		methods = append(methods, codegen.NewCommentedValueMethod(
			p.packageName(),
			p.getFnName(i),
			p.StructName(),
			/*params=*/ nil,
			[]jen.Code{jen.Id(kind.ConcreteKind)},
			[]jen.Code{jen.Return(jen.Id(codegen.This()).Dot(p.memberName(i)))},
			getComment,
		))
	}
	return methods
}

// unknownMemberDef returns the definition of a struct member that handles
// a property whose type is unknown.
func (p *FunctionalPropertyGenerator) unknownMemberDef() jen.Code {
	return jen.Id(unknownMemberName).Index().Byte()
}

// unknownDeserializeCode generates the "else if it's a []byte" code used for
// deserializing unknown values.
func (p *FunctionalPropertyGenerator) unknownDeserializeCode() jen.Code {
	return jen.Else().If(
		jen.List(
			jen.Id("v"),
			jen.Id("ok"),
		).Op(":=").Id("i").Assert(
			jen.Index().Byte(),
		),
		jen.Id("ok"),
	).Block(
		jen.Id(codegen.This()).Op(":=").Op("&").Id(p.StructName()).Values(
			jen.Dict{
				jen.Id(unknownMemberName): jen.Id("v"),
			},
		),
		jen.Return(
			jen.Id(codegen.This()),
			jen.Err(),
		),
	)
}

package gen

import (
	"fmt"
	"github.com/cjslep/activity/tools/exp/codegen"
	"github.com/dave/jennifer/jen"
	"sync"
)

const (
	typeResolverStructName         = "TypeResolver"
	predicateResolverStructName    = "PredicateResolver"
	resolveMethod                  = "Resolve"
	resolveMethodPrivate           = "resolve"
	applyMethod                    = "Apply"
	activityStreamInterface        = "ActivityStreamsInterface"
	callbackMember                 = "callbacks"
	predicateMember                = "predicate"
	errorNoMatch                   = "ErrNoCallbackMatch"
	errorUnhandled                 = "ErrUnknownType"
	errorPredicateUnmatched        = "ErrPredicateUnmatched"
	errorCannotTypeAssert          = "errCannotTypeAssertType"
	errorCannotTypeAssertPredicate = "errCannotTypeAssertPredicate"
	isUnFnName                     = "IsUnknownOrUnmatchedErr"
)

// TODO: Interface-driven resolvers. For hierarchy.

// ResolverGenerator generates the code required for the TypeResolver and the
// PredicateResolver.
type ResolverGenerator struct {
	pkg                                Package
	types                              []*TypeGenerator
	cacheOnce                          sync.Once
	cachedPredicate                    *codegen.Struct
	cachedType                         *codegen.Struct
	cachedErrNoMatch                   jen.Code
	cachedErrUnhandled                 jen.Code
	cachedErrPredicateUnmatched        jen.Code
	cachedErrCannotTypeAssert          jen.Code
	cachedErrCannotTypeAssertPredicate jen.Code
	cachedIsUnFn                       *codegen.Function
	cachedASInterface                  *codegen.Interface
}

// Creates a new ResolverGenerator for generating all the methods, functions,
// errors, interface, and struct types needed for them.
//
// Must be constructed after all TypeGenerators.
func NewResolverGenerator(
	tgs []*TypeGenerator,
	pkg Package) *ResolverGenerator {
	return &ResolverGenerator{
		pkg:   pkg,
		types: tgs,
	}
}

// Definition returns the TypeResolver and PredicateResolver.
func (r *ResolverGenerator) Definition() (types, predicates *codegen.Struct, errs []jen.Code, isUnFn *codegen.Function, asInterface *codegen.Interface) {
	r.cacheOnce.Do(func() {
		r.cachedType = codegen.NewStruct(
			// TODO: Comment
			fmt.Sprintf("", typeResolverStructName),
			typeResolverStructName,
			r.typeResolverMethods(),
			r.typeResolverFunctions(),
			r.typeResolverMembers())
		r.cachedPredicate = codegen.NewStruct(
			// TODO: Comment
			fmt.Sprintf("", predicateResolverStructName),
			predicateResolverStructName,
			r.predicateResolverMethods(),
			r.predicateResolverFunctions(),
			r.predicateResolverMembers())
		r.cachedErrNoMatch = r.errorNoMatch()
		r.cachedErrUnhandled = r.errorUnhandled()
		r.cachedErrPredicateUnmatched = r.errorPredicateUnmatched()
		r.cachedErrCannotTypeAssert = r.errorCannotTypeAssert()
		r.cachedErrCannotTypeAssertPredicate = r.errorCannotTypeAssertPredicate()
		r.cachedIsUnFn = r.isUnFn()
		r.cachedASInterface = r.asInterface()
	})
	return r.cachedType, r.cachedPredicate, []jen.Code{
		r.cachedErrNoMatch,
		r.cachedErrUnhandled,
		r.cachedErrPredicateUnmatched,
		r.cachedErrCannotTypeAssert,
		r.cachedErrCannotTypeAssertPredicate,
	}, r.cachedIsUnFn, r.cachedASInterface
}

// errorNoMatch returns the declaration for the ErrNoMatch global value.
func (r *ResolverGenerator) errorNoMatch() jen.Code {
	// TODO: Comment
	return jen.Var().Id(errorNoMatch).Error().Op("=").Qual("errors", "New").Call(jen.Lit("activity stream did not match the callback function"))
}

// errorUnhandled returns the declaration for the ErrUnhandled global value.
func (r *ResolverGenerator) errorUnhandled() jen.Code {
	// TODO: Comment
	return jen.Var().Id(errorUnhandled).Error().Op("=").Qual("errors", "New").Call(jen.Lit("activity stream did not match any known types"))
}

// errorCannotTypeAssert returns the declaration for the errCannotTypeAssert
// global value.
func (r *ResolverGenerator) errorCannotTypeAssert() jen.Code {
	// TODO: Comment
	return jen.Var().Id(errorCannotTypeAssert).Error().Op("=").Qual("errors", "New").Call(jen.Lit("activity stream type cannot be asserted to its interface"))
}

// errorCannotTypeAssertPredicate returns the declaration for the
// errCannotTypeAssert global value.
func (r *ResolverGenerator) errorCannotTypeAssertPredicate() jen.Code {
	// TODO: Comment
	return jen.Var().Id(errorCannotTypeAssertPredicate).Error().Op("=").Qual("errors", "New").Call(jen.Lit("predicate cannot be type asserted to a known function type"))
}

// errorPredicateUnmatched returns the declaration for the ErrPredicateUnmatched
// global value.
func (r *ResolverGenerator) errorPredicateUnmatched() jen.Code {
	// TODO: Comment
	return jen.Var().Id(errorPredicateUnmatched).Error().Op("=").Qual("errors", "New").Call(jen.Lit("activity stream did not match type demanded by predicate"))
}

// isUnFn returns a function that returns true if an error is one dealing with
// Unmatched or Unhandled errors.
func (r *ResolverGenerator) isUnFn() *codegen.Function {
	return codegen.NewCommentedFunction(
		r.pkg.Path(),
		isUnFnName,
		[]jen.Code{
			jen.Err().Error(),
		},
		[]jen.Code{
			jen.Bool(),
		},
		[]jen.Code{
			jen.Return(
				jen.Err().Op("==").Id(errorPredicateUnmatched).Op(
					"||",
				).Err().Op("==").Id(errorUnhandled).Op(
					"||",
				).Err().Op("==").Id(errorNoMatch),
			),
		},
		// TODO: Comment
		fmt.Sprintf("%s", isUnFnName))

}

// typeResolverMethods returns the methods for the TypeResolver.
func (r *ResolverGenerator) typeResolverMethods() (m []*codegen.Method) {
	m = append(m, r.resolveMethod(resolveMethod, typeResolverStructName))
	return
}

// predicateResolverMethods returns the methods for the PredicateResolver.
func (r *ResolverGenerator) predicateResolverMethods() (m []*codegen.Method) {
	m = append(m, r.applyMethod())
	m = append(m, r.resolveMethod(resolveMethodPrivate, predicateResolverStructName))
	return
}

// typeResolverFunctions returns the functions for the TypeResolver.
func (r *ResolverGenerator) typeResolverFunctions() (f []*codegen.Function) {
	f = append(f, codegen.NewCommentedFunction(
		r.pkg.Path(),
		fmt.Sprintf("%s%s", constructorName, typeResolverStructName),
		[]jen.Code{
			jen.Id("callbacks").Index().Interface(),
		},
		[]jen.Code{
			jen.Op("*").Id(typeResolverStructName),
			jen.Error(),
		},
		[]jen.Code{
			jen.For(
				jen.List(
					jen.Id("_"),
					jen.Id("cb"),
				).Op(":=").Range().Id("callbacks"),
			).Block(
				jen.Commentf("Each callback function must satisfy one known function signature, or else we will generate a runtime error instead of silently fail."),
				jen.Switch(
					jen.Id("cb").Assert(jen.Type()),
				).Block(
					r.mustAssertToKnownTypes("cb"),
				),
			),
			jen.Return(
				jen.Op("&").Id(typeResolverStructName).Values(
					jen.Dict{
						jen.Id(callbackMember): jen.Id("callbacks"),
					},
				),
				jen.Nil(),
			),
		},
		// TODO: Comment
		fmt.Sprintf("%s%s", constructorName, typeResolverStructName)))
	return
}

// predicateResolverFunctions returns the functions for the PredicateResolver.
func (r *ResolverGenerator) predicateResolverFunctions() (f []*codegen.Function) {
	f = append(f, codegen.NewCommentedFunction(
		r.pkg.Path(),
		fmt.Sprintf("%s%s", constructorName, predicateResolverStructName),
		[]jen.Code{
			jen.Id("callbacks").Index().Interface(),
			jen.Id("predicate").Interface(),
		},
		[]jen.Code{
			jen.Op("*").Id(predicateResolverStructName),
			jen.Error(),
		},
		[]jen.Code{
			jen.For(
				jen.List(
					jen.Id("_"),
					jen.Id("cb"),
				).Op(":=").Range().Id("callbacks"),
			).Block(
				jen.Commentf("Each callback function must satisfy one known function signature, or else we will generate a runtime error instead of silently fail."),
				jen.Switch(
					jen.Id("cb").Assert(jen.Type()),
				).Block(
					r.mustAssertToKnownTypes("cb"),
				),
			),
			jen.Commentf("The predicate must satisfy one known predicate function signature, or else we will generate a runtime error instead of silently fail."),
			jen.Switch(
				jen.Id("predicate").Assert(jen.Type()),
			).Block(
				r.mustAssertToKnownPredicate("predicate"),
			),
			jen.Return(
				jen.Op("&").Id(typeResolverStructName).Values(
					jen.Dict{
						jen.Id(callbackMember):  jen.Id("callbacks"),
						jen.Id(predicateMember): jen.Id("predicate"),
					},
				),
				jen.Nil(),
			),
		},
		// TODO: Comment
		fmt.Sprintf("%s%s", constructorName, predicateResolverStructName)))
	return
}

// typeResolverMembers returns the members for the TypeResolver.
func (r *ResolverGenerator) typeResolverMembers() (m []jen.Code) {
	m = append(m, jen.Id(callbackMember).Index().Interface())
	return
}

// predicateResolverMembers returns the members for the PredicateResolver.
func (r *ResolverGenerator) predicateResolverMembers() (m []jen.Code) {
	m = append(m, jen.Id(callbackMember).Index().Interface())
	m = append(m, jen.Id(predicateMember).Interface())
	return
}

// resolveMethod returns the Resolve method.
func (r *ResolverGenerator) resolveMethod(name, structName string) *codegen.Method {
	impl := jen.Empty()
	for _, t := range r.types {
		impl = impl.Case(
			jen.Lit(t.TypeName()),
		).Block(
			jen.If(
				jen.List(
					jen.Id("fn"),
					jen.Id("ok"),
				).Op(":=").Id("i").Assert(
					jen.Func().Parens(
						jen.List(
							jen.Qual("context", "Context"),
							jen.Qual(t.PublicPackage().Path(), t.InterfaceName()),
						),
					).Error(),
				),
				jen.Id("ok"),
			).Block(
				jen.If(
					jen.List(
						jen.Id("v"),
						jen.Id("ok"),
					).Op(":=").Id("o").Assert(
						jen.Qual(t.PublicPackage().Path(), t.InterfaceName()),
					),
					jen.Id("ok"),
				).Block(
					jen.Return(
						jen.Id("fn").Call(jen.Id("ctx"), jen.Id("v")),
					),
				).Else().Block(
					jen.Commentf("This occurs when the implementation is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code."),
					jen.Return(
						jen.Id(errorCannotTypeAssert),
					),
				),
			),
		).Line()
	}
	return codegen.NewCommentedValueMethod(
		r.pkg.Path(),
		name,
		structName,
		[]jen.Code{
			jen.Id("ctx").Qual("context", "Context"),
			jen.Id("o").Id(activityStreamInterface),
		},
		[]jen.Code{
			jen.Error(),
		},
		[]jen.Code{
			jen.For(
				jen.List(
					jen.Id("_"),
					jen.Id("i"),
				).Op(":=").Range().Id(codegen.This()).Id(callbackMember),
			).Block(
				jen.Switch(jen.Id("o").Dot(typeNameMethod).Call()).Block(
					impl.Default().Block(
						jen.Return(
							jen.Id(errorUnhandled),
						),
					),
				),
			),
			jen.Return(
				jen.Id(errorNoMatch),
			),
		},
		// TODO: Comment
		fmt.Sprintf("", resolveMethod))
}

// mustAssertToKnownTypes creates the type assertion switch statement that will
// return an error if the parameter named does not match any of the expected
// function signatures.
func (r *ResolverGenerator) mustAssertToKnownTypes(paramName string) jen.Code {
	c := jen.Empty()
	for _, t := range r.types {
		c = c.Case(
			jen.Func().Parens(
				jen.List(
					jen.Qual("context", "Context"),
					jen.Qual(t.PublicPackage().Path(), t.InterfaceName()),
				),
			).Error(),
		).Block(
			jen.Commentf("Do nothing, this callback has a correct signature."),
		).Line()
	}
	c = c.Default().Block(
		jen.Return(
			jen.Nil(),
			jen.Qual("errors", "New").Call(jen.Lit("a callback function is of the wrong signature and would never be called")),
		),
	)
	return c
}

// applyMethod generates the code required for the Apply method.
func (r *ResolverGenerator) applyMethod() *codegen.Method {
	impl := jen.Empty()
	for _, t := range r.types {
		impl = impl.Case(
			jen.Func().Parens(
				jen.List(
					jen.Qual("context", "Context"),
					jen.Qual(t.PublicPackage().Path(), t.InterfaceName()),
				),
			).Parens(
				jen.List(
					jen.Bool(),
					jen.Error(),
				),
			),
		).Block(
			jen.If(
				jen.List(
					jen.Id("v"),
					jen.Id("ok"),
				).Op(":=").Id("o").Assert(
					jen.Qual(t.PublicPackage().Path(), t.InterfaceName()),
				),
				jen.Id("ok"),
			).Block(
				jen.List(
					jen.Id("predicatePasses"),
					jen.Err(),
				).Op("=").Id("fn").Call(jen.Id("ctx"), jen.Id("v")),
			).Else().Block(
				jen.Return(
					jen.False(),
					jen.Id(errorPredicateUnmatched),
				),
			),
		).Line()
	}
	return codegen.NewCommentedValueMethod(
		r.pkg.Path(),
		applyMethod,
		predicateResolverStructName,
		[]jen.Code{
			jen.Id("ctx").Qual("context", "Context"),
			jen.Id("o").Id(activityStreamInterface),
		},
		[]jen.Code{
			jen.Bool(),
			jen.Error(),
		},
		[]jen.Code{
			jen.Var().Id("predicatePasses").Bool(),
			jen.Var().Err().Error(),
			jen.Switch(jen.Id("fn").Op(":=").Id(codegen.This()).Dot(predicateMember).Assert(jen.Type())).Block(
				impl.Default().Block(
					jen.Commentf("The constructor should guard against this error. If it is encountered, then there is a bug in the code generator."),
					jen.Return(
						jen.False(),
						jen.Id(errorCannotTypeAssertPredicate),
					),
				),
			),
			jen.If(
				jen.Id("predicatePasses"),
			).Block(
				jen.Return(
					jen.True(),
					jen.Id(codegen.This()).Dot(resolveMethodPrivate).Call(
						jen.Id("ctx"),
						jen.Id("o"),
					),
				),
			).Else().Block(
				jen.Return(
					jen.False(),
					jen.Nil(),
				),
			),
		},
		// TODO: Comment
		fmt.Sprintf("%s", applyMethod))
}

// mustAssertToKnownPredicate ensures the parameter name types-asserts to a
// known signature, or returns an error.
func (r *ResolverGenerator) mustAssertToKnownPredicate(paramName string) jen.Code {
	c := jen.Empty()
	for _, t := range r.types {
		c = c.Case(
			jen.Func().Parens(
				jen.List(
					jen.Qual("context", "Context"),
					jen.Qual(t.PublicPackage().Path(), t.InterfaceName()),
				),
			).Parens(
				jen.List(
					jen.Bool(),
					jen.Error(),
				),
			),
		).Block(
			jen.Commentf("Do nothing, this predicate has a correct signature."),
		).Line()
	}
	c = c.Default().Block(
		jen.Return(
			jen.Nil(),
			jen.Qual("errors", "New").Call(jen.Lit("the predicate function is of the wrong signature and would never be called")),
		),
	)
	return c
}

// asInterface returns the ActivityStreamsInterface.
func (r *ResolverGenerator) asInterface() *codegen.Interface {
	return codegen.NewInterface(
		r.pkg.Path(),
		activityStreamInterface,
		[]codegen.FunctionSignature{
			{
				Name:   typeNameMethod,
				Params: nil,
				Ret:    []jen.Code{jen.String()},
				// TODO: Comment
				Comment: fmt.Sprintf("%s", typeNameMethod),
			},
		},
		// TODO: Comment
		fmt.Sprintf("%s", activityStreamInterface))
}

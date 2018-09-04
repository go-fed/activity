package exp

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

// TODO: Set should clear out other types not being set
// TODO: Clear should be non-type-specific.
// TODO: Has for single type values should be Is.
// TODO: Has should be non-type-specific and be HasAnyValue.
// TODO: Generate Prepend for non-functional type.
// TODO: Generate Remove for non-functional type.
// TODO: Do not generate Clear for iterator
// TODO: Break out into multiple files.
// TODO: Generate Serialize and Deserialize functions

const (
	getMethod   = "Get"
	setMethod   = "Set"
	hasMethod   = "Has"
	clearMethod = "Clear"
	isMethod    = "Is"
)

type Identifier struct {
	LowerName string
	CamelName string
}

type Kind struct {
	Name                  Identifier
	ConcreteKind          string
	Nilable               bool
	HasNaturalLanguageMap bool
}

type PropertyGenerator struct {
	Name       Identifier
	Kinds      []Kind
	asIterator bool
}

func (p *PropertyGenerator) structName() string {
	if p.asIterator {
		return p.Name.CamelName
	}
	return fmt.Sprintf("%sProperty", p.Name.CamelName)
}

func (p *PropertyGenerator) kindCamelName(i int) string {
	return p.Kinds[i].Name.CamelName
}

func (p *PropertyGenerator) memberName(i int) string {
	return fmt.Sprintf("%sMember", p.Kinds[i].Name.LowerName)
}

func (p *PropertyGenerator) hasMemberName(i int) string {
	if len(p.Kinds) == 1 && p.Kinds[0].Nilable {
		panic("PropertyGenerator.hasMemberName called for nilable single value")
	}
	return fmt.Sprintf("has%sMember", p.Kinds[i].Name.CamelName)
}

type FunctionalPropertyGenerator struct {
	PropertyGenerator
}

func (p *FunctionalPropertyGenerator) Generate() *jen.Statement {
	return p.def().Line().Add(p.funcs())
}

func (p *FunctionalPropertyGenerator) def() *jen.Statement {
	if len(p.Kinds) == 1 {
		return p.singleTypeDef()
	} else {
		return p.multiTypeDef()
	}
}

func (p *FunctionalPropertyGenerator) funcs() *jen.Statement {
	if len(p.Kinds) == 1 {
		return p.singleTypeFuncs()
	} else {
		return p.multiTypeFuncs()
	}
}

func (p *FunctionalPropertyGenerator) singleTypeDef() *jen.Statement {
	if p.Kinds[0].Nilable {
		return jen.Type().Id(p.structName()).Id(p.Kinds[0].ConcreteKind)
	} else {
		return jen.Type().Id(p.structName()).Struct(
			jen.Id(p.memberName(0)).Id(p.Kinds[0].ConcreteKind),
			jen.Id(p.hasMemberName(0)).Bool(),
		)
	}
}

func (p *FunctionalPropertyGenerator) singleTypeFuncs() *jen.Statement {
	memberFunc := jen.Func().Params(
		jen.Id("t").Id(p.structName()),
	)
	memberPtrFunc := jen.Func().Params(
		jen.Id("t").Op("*").Id(p.structName()),
	)
	var has *jen.Statement
	var get *jen.Statement
	var set *jen.Statement
	var clear *jen.Statement
	if p.Kinds[0].Nilable {
		has = memberFunc.Clone().Id(hasMethod).Params().Params(
			jen.Bool(),
		).Block(
			jen.Return(jen.Id("t").Op("!=").Nil()),
		)
		get = memberFunc.Clone().Id(getMethod).Params().Params(
			jen.Id(p.Kinds[0].ConcreteKind),
		).Block(
			jen.Return(jen.Id("t")),
		)
		set = memberPtrFunc.Clone().Id(setMethod).Params(
			jen.Id("v").Id(p.Kinds[0].ConcreteKind),
		).Block(
			jen.Op("*").Id("t").Op("=").Id("v"),
		)
		clear = memberPtrFunc.Clone().Id(clearMethod).Params().Block(
			jen.Op("*").Id("t").Op("=").Nil(),
		)
	} else {
		has = memberFunc.Clone().Id(hasMethod).Params().Params(
			jen.Bool(),
		).Block(
			jen.Return(jen.Id("t").Dot(p.hasMemberName(0))),
		)
		get = memberFunc.Clone().Id(getMethod).Params().Params(
			jen.Id(p.Kinds[0].ConcreteKind),
		).Block(
			jen.Return(jen.Id("t").Dot(p.memberName(0))),
		)
		set = memberPtrFunc.Clone().Id(setMethod).Params(
			jen.Id("v").Id(p.Kinds[0].ConcreteKind),
		).Block(
			jen.Id("t").Dot(p.memberName(0)).Op("=").Id("v"),
			jen.Id("t").Dot(p.hasMemberName(0)).Op("=").True(),
		)
		clear = memberPtrFunc.Clone().Id(clearMethod).Params().Block(
			jen.Id("t").Dot(p.hasMemberName(0)).Op("=").False(),
		)
	}
	return has.Line().Add(get).Line().Add(set).Line().Add(clear)
}

func (p *FunctionalPropertyGenerator) multiTypeDef() *jen.Statement {
	members := &jen.Statement{}
	for i, kind := range p.Kinds {
		var kindMembers *jen.Statement
		if kind.Nilable {
			kindMembers = jen.Id(p.memberName(i)).Id(p.Kinds[i].ConcreteKind)
		} else {
			kindMembers = jen.Id(p.memberName(i)).Id(p.Kinds[i].ConcreteKind)
			kindMembers.Line().Add(
				jen.Id(p.hasMemberName(i)).Bool(),
			)
		}
		if i > 0 {
			members.Line()
		}
		members.Add(kindMembers)
	}
	return jen.Type().Id(p.structName()).Struct(members)
}

func (p *FunctionalPropertyGenerator) multiTypeFuncs() *jen.Statement {
	funcs := &jen.Statement{}
	memberFunc := jen.Func().Params(
		jen.Id("t").Id(p.structName()),
	)
	memberPtrFunc := jen.Func().Params(
		jen.Id("t").Op("*").Id(p.structName()),
	)
	for i, kind := range p.Kinds {
		var is *jen.Statement
		var get *jen.Statement
		var set *jen.Statement
		var clear *jen.Statement
		if kind.Nilable {
			is = memberFunc.Clone().Id(
				fmt.Sprintf("%s%s", isMethod, p.kindCamelName(i)),
			).Params().Params(
				jen.Bool(),
			).Block(
				jen.Return(jen.Id("t").Dot(p.memberName(i)).Op("!=").Nil()),
			)
			set = memberPtrFunc.Clone().Id(
				fmt.Sprintf("%s%s", setMethod, p.kindCamelName(i)),
			).Params(
				jen.Id("v").Id(kind.ConcreteKind),
			).Block(
				jen.Id("t").Dot(p.memberName(i)).Op("=").Id("v"),
			)
			clear = memberPtrFunc.Clone().Id(clearMethod).Params().Block(
				jen.Id("t").Dot(p.memberName(i)).Op("=").Nil(),
			)
		} else {
			is = memberFunc.Clone().Id(
				fmt.Sprintf("%s%s", isMethod, p.kindCamelName(i)),
			).Params().Params(
				jen.Bool(),
			).Block(
				jen.Return(jen.Id("t").Dot(p.hasMemberName(i))),
			)
			set = memberPtrFunc.Clone().Id(
				fmt.Sprintf("%s%s", setMethod, p.kindCamelName(i)),
			).Params(
				jen.Id("v").Id(kind.ConcreteKind),
			).Block(
				jen.Id("t").Dot(p.memberName(i)).Op("=").Id("v"),
				jen.Id("t").Dot(p.hasMemberName(i)).Op("=").True(),
			)
			clear = memberPtrFunc.Clone().Id(clearMethod).Params().Block(
				jen.Id("t").Dot(p.hasMemberName(i)).Op("=").False(),
			)
		}
		get = memberFunc.Clone().Id(
			fmt.Sprintf("%s%s", getMethod, p.kindCamelName(i)),
		).Params().Params(
			jen.Id(kind.ConcreteKind),
		).Block(
			jen.Return(jen.Id("t").Dot(p.memberName(i))),
		)
		funcs.Add(is).Line()
		funcs.Add(get).Line()
		funcs.Add(set).Line()
		funcs.Add(clear).Line()
	}
	return funcs
}

type NonFunctionalPropertyGenerator struct {
	PropertyGenerator
}

func (p *NonFunctionalPropertyGenerator) Generate() *jen.Statement {
	return p.def().Line().Add(p.funcs())
}

func (p *PropertyGenerator) typeName() string {
	return fmt.Sprintf("%sProperty", p.Name.CamelName)
}

func (p *NonFunctionalPropertyGenerator) iteratorTypeName() Identifier {
	return Identifier{
		LowerName: fmt.Sprintf("%sPropertyIterator", p.Name.LowerName),
		CamelName: fmt.Sprintf("%sPropertyIterator", p.Name.CamelName),
	}
}

func (p *NonFunctionalPropertyGenerator) elementTypeGenerator() *FunctionalPropertyGenerator {
	return &FunctionalPropertyGenerator{
		PropertyGenerator{
			Name:       p.iteratorTypeName(),
			Kinds:      p.Kinds,
			asIterator: true,
		},
	}
}

func (p *NonFunctionalPropertyGenerator) def() *jen.Statement {
	return p.elementTypeGenerator().def().Line().Add(
		jen.Type().Id(p.typeName()).Index().Id(p.iteratorTypeName().CamelName),
	)
}

func (p *NonFunctionalPropertyGenerator) funcs() *jen.Statement {
	return p.elementTypeGenerator().funcs()
}

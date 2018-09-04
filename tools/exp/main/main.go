package main

import (
	"fmt"
	"github.com/go-fed/activity/tools/exp"
)

func main() {
	x := &exp.FunctionalPropertyGenerator{
		exp.PropertyGenerator{
			Name: exp.Identifier{
				LowerName: "testFunctional",
				CamelName: "TestFunctional",
			},
			Kinds: []exp.Kind{
				{
					Name: exp.Identifier{
						LowerName: "iri",
						CamelName: "IRI",
					},
					ConcreteKind:          "*url.URL",
					Nilable:               true,
					HasNaturalLanguageMap: false,
				},
			},
		},
	}
	y := &exp.FunctionalPropertyGenerator{
		exp.PropertyGenerator{
			Name: exp.Identifier{
				LowerName: "testFunctionalNonnil",
				CamelName: "TestFunctionalNonil",
			},
			Kinds: []exp.Kind{
				{
					Name: exp.Identifier{
						LowerName: "number",
						CamelName: "Number",
					},
					ConcreteKind:          "int",
					Nilable:               false,
					HasNaturalLanguageMap: false,
				},
			},
		},
	}
	z := &exp.FunctionalPropertyGenerator{
		exp.PropertyGenerator{
			Name: exp.Identifier{
				LowerName: "testFunctionalMultiType",
				CamelName: "TestFunctionalMultiType",
			},
			Kinds: []exp.Kind{
				{
					Name: exp.Identifier{
						LowerName: "iri",
						CamelName: "IRI",
					},
					ConcreteKind:          "*url.URL",
					Nilable:               true,
					HasNaturalLanguageMap: false,
				},
				{
					Name: exp.Identifier{
						LowerName: "number",
						CamelName: "Number",
					},
					ConcreteKind:          "int",
					Nilable:               false,
					HasNaturalLanguageMap: false,
				},
			},
		},
	}
	zz := &exp.NonFunctionalPropertyGenerator{
		exp.PropertyGenerator{
			Name: exp.Identifier{
				LowerName: "testNonFunctionalMultiType",
				CamelName: "TestNonFunctionalMultiType",
			},
			Kinds: []exp.Kind{
				{
					Name: exp.Identifier{
						LowerName: "iri",
						CamelName: "IRI",
					},
					ConcreteKind:          "*url.URL",
					Nilable:               true,
					HasNaturalLanguageMap: false,
				},
				{
					Name: exp.Identifier{
						LowerName: "number",
						CamelName: "Number",
					},
					ConcreteKind:          "int",
					Nilable:               false,
					HasNaturalLanguageMap: false,
				},
			},
		},
	}
	fmt.Printf("%#v\n\n", x.Generate())
	fmt.Printf("%#v\n\n", y.Generate())
	fmt.Printf("%#v\n\n", z.Generate())
	fmt.Printf("%#v\n\n", zz.Generate())
}

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
					SerializeFnName:       "SerializeIRI",
					DeserializeFnName:     "DeserializeIRI",
					LessFnName:            "LessIRI",
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
					SerializeFnName:       "SerializeNumber",
					DeserializeFnName:     "DeserializeNumber",
					LessFnName:            "LessNumber",
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
					SerializeFnName:       "SerializeIRI",
					DeserializeFnName:     "DeserializeIRI",
					LessFnName:            "LessIRI",
				},
				{
					Name: exp.Identifier{
						LowerName: "number",
						CamelName: "Number",
					},
					ConcreteKind:          "int",
					Nilable:               false,
					HasNaturalLanguageMap: false,
					SerializeFnName:       "SerializeNumber",
					DeserializeFnName:     "DeserializeNumber",
					LessFnName:            "LessNumber",
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
					SerializeFnName:       "SerializeIRI",
					DeserializeFnName:     "DeserializeIRI",
					LessFnName:            "LessIRI",
				},
				{
					Name: exp.Identifier{
						LowerName: "number",
						CamelName: "Number",
					},
					ConcreteKind:          "int",
					Nilable:               false,
					HasNaturalLanguageMap: false,
					SerializeFnName:       "SerializeNumber",
					DeserializeFnName:     "DeserializeNumber",
					LessFnName:            "LessNumber",
				},
			},
		},
	}
	fmt.Printf("%#v", x.Definition())
	fmt.Printf("%#v", y.Definition())
	fmt.Printf("%#v", z.Definition())
	fmt.Printf("%#v", zz.Definition())
}

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/cjslep/activity/tools/exp/rdf"
	"github.com/cjslep/activity/tools/exp/rdf/owl"
	"github.com/cjslep/activity/tools/exp/rdf/rdfs"
	"github.com/cjslep/activity/tools/exp/rdf/schema"
	"github.com/cjslep/activity/tools/exp/rdf/xsd"
	"io/ioutil"
)

var registry *rdf.RDFRegistry

func mustAddOntology(o rdf.Ontology) {
	if registry == nil {
		registry = rdf.NewRDFRegistry()
	}
	if err := registry.AddOntology(o); err != nil {
		panic(err)
	}
}

func init() {
	mustAddOntology(&xsd.XMLOntology{})
	mustAddOntology(&owl.OWLOntology{})
	mustAddOntology(&rdf.RDFOntology{})
	mustAddOntology(&rdfs.RDFSchemaOntology{})
	mustAddOntology(&schema.SchemaOntology{})
}

var (
	input = flag.String("input", "spec.json", "Input JSON-LD specification used to generate Go code.")
)

func main() {
	flag.Parse()

	b, err := ioutil.ReadFile(*input)
	if err != nil {
		panic(err)
	}
	var inputJSON map[string]interface{}
	err = json.Unmarshal(b, &inputJSON)
	if err != nil {
		panic(err)
	}
	_, err = rdf.ParseVocabulary(registry, inputJSON)
	if err != nil {
		panic(err)
	}
	fmt.Printf("done\n")
}

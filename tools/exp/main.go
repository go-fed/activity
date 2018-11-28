package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-fed/activity/tools/exp/rdf"
	"github.com/go-fed/activity/tools/exp/rdf/as"
	"github.com/go-fed/activity/tools/exp/rdf/owl"
	"github.com/go-fed/activity/tools/exp/rdf/rdfs"
	"github.com/go-fed/activity/tools/exp/rdf/schema"
	"github.com/go-fed/activity/tools/exp/rdf/xsd"
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
	mustAddOntology(&as.ActivityStreamsOntology{})
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
	nodes, err := rdf.ParseJSONLDContext(registry, inputJSON)
	if err != nil {
		panic(err)
	}
	fmt.Printf("len(nodes) = %d\n", len(nodes))
}

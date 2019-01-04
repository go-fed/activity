package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/cjslep/activity/tools/exp/convert"
	"github.com/cjslep/activity/tools/exp/props"
	"github.com/cjslep/activity/tools/exp/rdf"
	"github.com/cjslep/activity/tools/exp/rdf/owl"
	"github.com/cjslep/activity/tools/exp/rdf/rdfs"
	"github.com/cjslep/activity/tools/exp/rdf/schema"
	"github.com/cjslep/activity/tools/exp/rdf/xsd"
	"io/ioutil"
	"os"
	"strings"
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
	mustAddOntology(&xsd.XMLOntology{Package: "xml"})
	mustAddOntology(&owl.OWLOntology{})
	mustAddOntology(&rdf.RDFOntology{Package: "rdf"})
	mustAddOntology(&rdfs.RDFSchemaOntology{})
	mustAddOntology(&schema.SchemaOntology{})
}

var (
	input = flag.String("input", "spec.json", "Input JSON-LD specification used to generate Go code.")
	// TODO: Be more rigorous when applying this. Also, clear the default value I am using for convenience.
	prefix = flag.String("prefix", "github.com/cjslep/activity/tools/exp/tmp", "Package prefix to use for all generated package paths. This should be the prefix in the GOPATH directory if generating in a subdirectory.")
	// TODO: Use this flag
	root   = flag.String("root", "github.com/go-fed/activity/", "Go import path prefix for generated packages")
	xmlpkg = flag.String("xmlpkg", "github.com/go-fed/activity/tools/exp/ref/xml", "Go package location for known XML references")
	rdfpkg = flag.String("rdfpkg", "github.com/go-fed/activity/tools/exp/ref/rdf", "Go package location for known RDF references")
	// TODO: Use this flag
	writeWellKnown = flag.Bool("write-well-known", false, "When true, also outputs well-known specifications to './ref' subdirectories (ex: XML, RDF)")
)

type list []string

func (l *list) String() string {
	return strings.Join(*l, ",")
}

func (l *list) Set(v string) error {
	vals := strings.Split(v, ",")
	*l = append(*l, vals...)
	return nil
}

func main() {
	// TODO: Use only one kind of flag style.
	var ref list
	var refspec list
	var refpkg list
	flag.Var(&ref, "ref", "Input JSON-LD specification of other referenced specifications. Must be the same size and in the same order as the 'refspec' and 'refpkg' flags")
	flag.Var(&refspec, "refspec", "Base URL for other referenced specifications. Must be the same size and in the same order as the 'ref' and 'refpkg' flags")
	flag.Var(&refpkg, "refpkg", "Golang package location for other referenced specifications. Must be the same size and in the same order as the 'ref' and 'refspec' flags")
	flag.Parse()
	// TODO: Flag validation

	b, err := ioutil.ReadFile(*input)
	if err != nil {
		panic(err)
	}
	var inputJSON map[string]interface{}
	err = json.Unmarshal(b, &inputJSON)
	if err != nil {
		panic(err)
	}
	p, err := rdf.ParseVocabulary(registry, inputJSON)
	if err != nil {
		panic(err)
	}
	c := &convert.Converter{
		Registry:              registry,
		VocabularyRoot:        props.NewPackageManager(*prefix, "gen/as"),
		ValueRoot:             props.NewPackageManager(*prefix, "gen/vals"),
		PropertyPackagePolicy: convert.PropertyFlatUnderRoot,
		PropertyPackageRoot:   props.NewPackageManager(*prefix, "gen/as/props"),
		TypePackagePolicy:     convert.TypeFlatUnderRoot,
		TypePackageRoot:       props.NewPackageManager(*prefix, "gen/as/types"),
	}
	f, err := c.Convert(p)
	if err != nil {
		panic(err)
	}
	for _, file := range f {
		file.F.ImportName(*xmlpkg, "xml")
		file.F.ImportName(*rdfpkg, "rdf")
		if e := os.MkdirAll("./"+file.Directory, 0777); e != nil {
			panic(e)
		}
		if e := file.F.Save("./" + file.Directory + "/" + file.FileName); e != nil {
			panic(e)
		}
	}
	fmt.Printf("done")
	// fmt.Printf("done\n%s\n", p)
}

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/cjslep/activity/tools/exp/convert"
	"github.com/cjslep/activity/tools/exp/gen"
	"github.com/cjslep/activity/tools/exp/rdf"
	"github.com/cjslep/activity/tools/exp/rdf/owl"
	"github.com/cjslep/activity/tools/exp/rdf/rdfs"
	"github.com/cjslep/activity/tools/exp/rdf/rfc"
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
	mustAddOntology(&rfc.RFCOntology{Package: "rfc"})
}

type list []string

func (l *list) String() string {
	return strings.Join(*l, ",")
}

func (l *list) Set(v string) error {
	vals := strings.Split(v, ",")
	*l = append(*l, vals...)
	return nil
}

type CommandLineFlags struct {
	specs      list
	prefix     *string
	individual *bool
}

func NewCommandLineFlags() *CommandLineFlags {
	c := &CommandLineFlags{
		// TODO: Be more rigorous when applying this. Also, clear the default value I am using for convenience.
		prefix:     flag.String("prefix", "github.com/cjslep/activity/tools/exp/tmp", "Package prefix to use for all generated package paths. This should be the prefix in the GOPATH directory if generating in a subdirectory."),
		individual: flag.Bool("individual", false, "Whether to generate types and properties in individual packages."),
	}
	flag.Var(&(c.specs), "spec", "Input JSON-LD specification used to generate Go code.")
	flag.Parse()
	if err := c.validate(); err != nil {
		panic(err)
	}
	return c
}

func (c *CommandLineFlags) validate() error {
	if len(c.specs) == 0 {
		return fmt.Errorf("specs must not be empty")
	}
	return nil
}

func main() {
	cmd := NewCommandLineFlags()

	inputJSONs := make([]rdf.JSONLD, 0, len(cmd.specs))
	for _, spec := range cmd.specs {
		b, err := ioutil.ReadFile(spec)
		if err != nil {
			panic(err)
		}
		var inputJSON map[string]interface{}
		err = json.Unmarshal(b, &inputJSON)
		if err != nil {
			panic(err)
		}
		inputJSONs = append(inputJSONs, inputJSON)
	}
	p, err := rdf.ParseVocabularies(registry, inputJSONs)
	if err != nil {
		panic(err)
	}
	policy := convert.FlatUnderRoot
	if *cmd.individual {
		policy = convert.IndividualUnderRoot
	}
	c := &convert.Converter{
		Registry:      registry,
		GenRoot:       gen.NewPackageManager(*cmd.prefix, "gen"),
		PackagePolicy: policy,
	}
	f, err := c.Convert(p)
	if err != nil {
		panic(err)
	}
	for _, file := range f {
		if e := os.MkdirAll("./"+file.Directory, 0777); e != nil {
			panic(e)
		}
		if e := file.F.Save("./" + file.Directory + "/" + file.FileName); e != nil {
			panic(e)
		}
	}
	fmt.Printf("Done\n")
}

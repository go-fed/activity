package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-fed/activity/tools/exp/convert"
	"github.com/go-fed/activity/tools/exp/gen"
	"github.com/go-fed/activity/tools/exp/rdf"
	"github.com/go-fed/activity/tools/exp/rdf/owl"
	"github.com/go-fed/activity/tools/exp/rdf/rdfs"
	"github.com/go-fed/activity/tools/exp/rdf/rfc"
	"github.com/go-fed/activity/tools/exp/rdf/schema"
	"github.com/go-fed/activity/tools/exp/rdf/xsd"
	"io/ioutil"
	"os"
	"strings"
)

// Global registry of "known" RDF ontologies. This manages the built-in
// knowledge of how to parse specific linked data documents. It may be cloned
// in the course of processing a JSON-LD document, due to "@context" dictating
// certain ontologies being aliased in some specifications and not others.
var registry *rdf.RDFRegistry

// mustAddOntology ensures that the registry global variable is not nil, and
// then adds the specific ontology or panics if it cannot.
func mustAddOntology(o rdf.Ontology) {
	if registry == nil {
		registry = rdf.NewRDFRegistry()
	}
	if err := registry.AddOntology(o); err != nil {
		panic(err)
	}
}

// At init time, get our built-in knowledge of OWL and other RDF ontologies
// into the registry, before main executes.
func init() {
	mustAddOntology(&xsd.XMLOntology{Package: "xml"})
	mustAddOntology(&owl.OWLOntology{})
	mustAddOntology(&rdf.RDFOntology{Package: "rdf"})
	mustAddOntology(&rdfs.RDFSchemaOntology{})
	mustAddOntology(&schema.SchemaOntology{})
	mustAddOntology(&rfc.RFCOntology{Package: "rfc"})
}

// list is a flag-friendly comma-separated list of strings. Also allows multiple
// definitions of the flag to not overwrite each other and instead result in a
// list of strings.
//
// The values of the flag cannot contain commas within them because the value
// will be split into two.
type list []string

// String turns this list into a single comma-separated string.
func (l *list) String() string {
	return strings.Join(*l, ",")
}

// Set adds a string value to the list, after splitting on the comma separator.
func (l *list) Set(v string) error {
	vals := strings.Split(v, ",")
	*l = append(*l, vals...)
	return nil
}

// CommandLineFlags manages the flags defined by this tool.
type CommandLineFlags struct {
	specs  list
	prefix string
}

// NewCommandLineFlags defines the flags expected to be used by this tool. Calls
// flag.Parse on behalf of the main program, and validates the flags. Returns an
// error if validation fails.
func NewCommandLineFlags() (*CommandLineFlags, error) {
	c := &CommandLineFlags{}
	// TODO: Be more rigorous when applying this. Also, clear the default value I am using for convenience.
	flag.StringVar(
		&c.prefix,
		"prefix",
		"github.com/go-fed/activity/tools/exp/tmp",
		"Package prefix to use for all generated package paths. This should be the prefix in the GOPATH directory if generating in a subdirectory.")
	flag.Var(&(c.specs), "spec", "Input JSON-LD specification used to generate Go code.")
	flag.Parse()
	return c, c.Validate()
}

// Validate applies custom validation logic to flags and returns an error if any
// flags violate these rules.
func (c *CommandLineFlags) Validate() error {
	if len(c.specs) == 0 {
		return fmt.Errorf("spec flag must not be empty")
	}
	return nil
}

// ReadSpecs returns the JSONLD contents of files specified in the 'spec' flag.
func (c *CommandLineFlags) ReadSpecs() (j []rdf.JSONLD, err error) {
	j = make([]rdf.JSONLD, 0, len(c.specs))
	for _, spec := range c.specs {
		var b []byte
		b, err = ioutil.ReadFile(spec)
		if err != nil {
			return
		}
		var inputJSON map[string]interface{}
		err = json.Unmarshal(b, &inputJSON)
		if err != nil {
			return
		}
		j = append(j, inputJSON)
	}
	return
}

func main() {
	// Read, Parse, and Validate command line flags
	cmd, err := NewCommandLineFlags()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Read input specification files
	fmt.Printf("Reading input specifications...\n")
	inputJSONs, err := cmd.ReadSpecs()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Parse specifications
	fmt.Printf("Parsing %d vocabularies...\n", len(inputJSONs))
	p, err := rdf.ParseVocabularies(registry, inputJSONs)
	if err != nil {
		panic(err)
	}

	// Convert to generated code
	fmt.Printf("Converting %d types, properties, and values...\n", p.Size())
	c := &convert.Converter{
		GenRoot:       gen.NewPackageManager(cmd.prefix, "gen"),
		PackagePolicy: convert.IndividualUnderRoot,
	}
	f, err := c.Convert(p)
	if err != nil {
		panic(err)
	}

	// Write generated code
	fmt.Printf("Writing %d files...\n", len(f))
	for _, file := range f {
		if e := os.MkdirAll("./"+file.Directory, 0777); e != nil {
			panic(e)
		}
		if e := file.F.Save("./" + file.Directory + "/" + file.FileName); e != nil {
			panic(e)
		}
	}
	fmt.Printf("Done!\n")
}

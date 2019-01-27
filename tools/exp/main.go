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

const (
	pathFlag = "path"
	specFlag = "spec"
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

// settableString is a flag-friendly string that distinguishes an empty string
// due to not being set and explicitly being set as empty at the command line.
type settableString struct {
	set bool
	str string
}

// String simply returns the string value of this settableString.
func (s *settableString) String() string {
	return s.str
}

// Set will mark this settableString's set as true and store the value.
func (s *settableString) Set(v string) error {
	s.set = true
	s.str = v
	return nil
}

// IsSet returns true if this value was explicitly set as a flag value.
func (s settableString) IsSet() bool {
	return s.set
}

// CommandLineFlags manages the flags defined by this tool.
type CommandLineFlags struct {
	// Flags
	specs list
	path  settableString
	// Additional data
	pathAutoDetected bool
}

// NewCommandLineFlags defines the flags expected to be used by this tool. Calls
// flag.Parse on behalf of the main program, and validates the flags. Returns an
// error if validation fails.
func NewCommandLineFlags() (*CommandLineFlags, error) {
	c := &CommandLineFlags{}
	flag.Var(
		&c.path,
		pathFlag,
		"Package path to use for all generated package paths. If using GOPATH, this is automatically detected as $GOPATH/<path>/ when generating in a subdirectory. Cannot be explicitly set to be empty.")
	flag.Var(&(c.specs), specFlag, "Input JSON-LD specification used to generate Go code.")
	flag.Parse()
	return c, c.Validate()
}

// detectPath attempts to detect the path to use when generating the code. The
// path is only detected if the tool is running in a subdirectory of GOPATH,
// and will be set to $GOPATH/<path>/. After this method runs without errors,
// c.path.IsSet will always return true.
//
// When auto-detecting, if GOPATH is not set then will return an error.
//
// If the path has already been set at the command line, does nothing.
func (c *CommandLineFlags) detectPath() error {
	if c.path.IsSet() {
		return nil
	}
	gopath, isSet := os.LookupEnv("GOPATH")
	if !isSet {
		return fmt.Errorf("cannot detect %q because GOPATH environmental variable is not set and %q flag was not explicitly set", pathFlag, pathFlag)
	}
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	if !strings.HasPrefix(pwd, gopath) {
		return fmt.Errorf("cannot detect %q because current working directory is not under GOPATH and %q flag was not explicitly set", pathFlag, pathFlag)
	}
	c.pathAutoDetected = true
	gopath = strings.Join([]string{gopath, "src", ""}, "/")
	c.path.Set(strings.TrimPrefix(pwd, gopath))
	return nil
}

// Validate applies custom validation logic to flags and returns an error if any
// flags violate these rules.
func (c *CommandLineFlags) Validate() error {
	if len(c.specs) == 0 {
		return fmt.Errorf("%q flag must not be empty", specFlag)
	}
	if err := c.detectPath(); err != nil {
		return err
	}
	if len(c.path.String()) == 0 {
		return fmt.Errorf("%q flag must not be empty", pathFlag)
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

// AutoDetectedPath returns true if the path flag was auto-detected.
func (c *CommandLineFlags) AutoDetectedPath() bool {
	return c.pathAutoDetected
}

// Path returns the path flag.
func (c *CommandLineFlags) Path() string {
	return c.path.String()
}

func main() {
	// Read, Parse, and Validate command line flags
	cmd, err := NewCommandLineFlags()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print auto-determined values
	if cmd.AutoDetectedPath() {
		fmt.Printf("Auto-detected path: %s\n", cmd.Path())
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
		GenRoot:       gen.NewPackageManager(cmd.Path(), ""),
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

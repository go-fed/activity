package rdf

type RDFOntology struct{}

func (o *RDFOntology) SpecURI() string {
	return "http://www.w3.org/1999/02/22-rdf-syntax-ns#"
}

func (o *RDFOntology) Load() ([]RDFNode, error) {
	return nil, nil
}

func (o *RDFOntology) LoadAsAlias(s string) ([]RDFNode, error) {
	return nil, nil
}

func (o *RDFOntology) LoadElement(name string, payload map[string]interface{}) ([]RDFNode, error) {
	return nil, nil
}

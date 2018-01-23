package defs

type Type struct {
	Name              string
	URI               string
	Notes             string
	DisjointWith      []*Type
	Extends           []*Type
	Properties        []*PropertyType
	WithoutProperties []*PropertyType
}

func (t *Type) GetProperties() []*PropertyType {
	omit := make(map[string]bool)
	for _, v := range t.WithoutProperties {
		omit[v.Name] = true
	}
	var parentProps []*PropertyType
	for _, p := range t.Extends {
		parentProps = append(parentProps, p.GetProperties()...)
	}
	properties := make([]*PropertyType, 0, len(t.Properties)+len(parentProps))
	set := make(map[string]bool)
	for _, v := range t.Properties {
		if !omit[v.Name] {
			if !set[v.Name] {
				properties = append(properties, v)
				set[v.Name] = true
			}
		}
	}
	for _, v := range parentProps {
		if !omit[v.Name] {
			if !set[v.Name] {
				properties = append(properties, v)
				set[v.Name] = true
			}
		}
	}
	return properties
}

func (t *Type) AllExtendsNames() []string {
	v := []string{t.Name}
	for _, e := range t.Extends {
		v = append(v, e.AllExtendsNames()...)
	}
	return v
}

type PropertyType struct {
	Name   string
	URI    string
	Notes  string
	Domain []DomainReference
	Range  []RangeReference
	// SubpropertyOf is ignorable as long as data is set up correctly
	SubpropertyOf        *PropertyType
	Functional           bool
	NaturalLanguageMap   bool
	PreferIRIConvenience bool
}

type ValueType struct {
	Name           string
	URI            string
	DefinitionType string
	Zero           string
	ZeroValue      string
	DeserializeFn  *FunctionDef
	SerializeFn    *FunctionDef
}

type DomainReference struct {
	// One of:
	T   *Type
	Any bool
}

type RangeReference struct {
	// One of:
	T   *Type
	V   *ValueType
	Any bool
}

func IsOnlyOtherPropertyBesidesIRI(nonIriIndex int, r []RangeReference) bool {
	if len(r) != 2 {
		return false
	}
	o := (nonIriIndex + 1) % 2
	return r[o].V != nil && IsIRIValueType(r[o].V)
}

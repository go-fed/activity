package propertyunits

import (
	"fmt"
	anyuri "github.com/go-fed/activity/streams/values/anyURI"
	string1 "github.com/go-fed/activity/streams/values/string"
	vocab "github.com/go-fed/activity/streams/vocab"
	"net/url"
)

// UnitsProperty is the functional property "units". It is permitted to be one of
// multiple value types. At most, one type of value can be present, or none at
// all. Setting a value will clear the other types of values so that only one
// of the 'Is' methods will return true. It is possible to clear all values,
// so that this property is empty.
type UnitsProperty struct {
	stringMember    string
	hasStringMember bool
	anyURIMember    *url.URL
	unknown         []byte
	alias           string
}

// DeserializeUnitsProperty creates a "units" property from an interface
// representation that has been unmarshalled from a text or binary format.
func DeserializeUnitsProperty(m map[string]interface{}, aliasMap map[string]string) (*UnitsProperty, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	propName := "units"
	if len(alias) > 0 {
		// Use alias both to find the property, and set within the property.
		propName = fmt.Sprintf("%s:%s", alias, "units")
	}
	if i, ok := m[propName]; ok {
		if v, err := string1.DeserializeString(i); err == nil {
			this := &UnitsProperty{
				alias:           alias,
				hasStringMember: true,
				stringMember:    v,
			}
			return this, nil
		} else if v, err := anyuri.DeserializeAnyURI(i); err == nil {
			this := &UnitsProperty{
				alias:        alias,
				anyURIMember: v,
			}
			return this, nil
		} else if str, ok := i.(string); ok {
			this := &UnitsProperty{
				alias:   alias,
				unknown: []byte(str),
			}
			return this, nil
		} else {
			return nil, fmt.Errorf("could not deserialize %q property", "units")
		}
	}
	return nil, nil
}

// NewUnitsProperty creates a new units property.
func NewUnitsProperty() *UnitsProperty {
	return &UnitsProperty{alias: ""}
}

// Clear ensures no value of this property is set. Calling HasAny or any of the
// 'Is' methods afterwards will return false.
func (this *UnitsProperty) Clear() {
	this.hasStringMember = false
	this.anyURIMember = nil
	this.unknown = nil
}

// GetAnyURI returns the value of this property. When IsAnyURI returns false,
// GetAnyURI will return an arbitrary value.
func (this UnitsProperty) GetAnyURI() *url.URL {
	return this.anyURIMember
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return an arbitrary value.
func (this UnitsProperty) GetIRI() *url.URL {
	return this.anyURIMember
}

// GetString returns the value of this property. When IsString returns false,
// GetString will return an arbitrary value.
func (this UnitsProperty) GetString() string {
	return this.stringMember
}

// HasAny returns true if any of the different values is set.
func (this UnitsProperty) HasAny() bool {
	return this.IsString() ||
		this.IsAnyURI()
}

// IsAnyURI returns true if this property has a type of "anyURI". When true, use
// the GetAnyURI and SetAnyURI methods to access and set this property.
func (this UnitsProperty) IsAnyURI() bool {
	return this.anyURIMember != nil
}

// IsIRI returns true if this property is an IRI. When true, use GetIRI and SetIRI
// to access and set this property
func (this UnitsProperty) IsIRI() bool {
	return this.anyURIMember != nil
}

// IsString returns true if this property has a type of "string". When true, use
// the GetString and SetString methods to access and set this property.
func (this UnitsProperty) IsString() bool {
	return this.hasStringMember
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this UnitsProperty) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/TR/activitystreams-vocabulary": this.alias}
	var child map[string]string

	/*
	   Since the literal maps in this function are determined at
	   code-generation time, this loop should not overwrite an existing key with a
	   new value.
	*/
	for k, v := range child {
		m[k] = v
	}
	return m
}

// KindIndex computes an arbitrary value for indexing this kind of value. This is
// a leaky API detail only for folks looking to replace the go-fed
// implementation. Applications should not use this method.
func (this UnitsProperty) KindIndex() int {
	if this.IsString() {
		return 0
	}
	if this.IsAnyURI() {
		return 1
	}
	if this.IsIRI() {
		return -2
	}
	return -1
}

// LessThan compares two instances of this property with an arbitrary but stable
// comparison. Applications should not use this because it is only meant to
// help alternative implementations to go-fed to be able to normalize
// nonfunctional properties.
func (this UnitsProperty) LessThan(o vocab.UnitsPropertyInterface) bool {
	idx1 := this.KindIndex()
	idx2 := o.KindIndex()
	if idx1 < idx2 {
		return true
	} else if idx1 > idx2 {
		return false
	} else if this.IsString() {
		return string1.LessString(this.GetString(), o.GetString())
	} else if this.IsAnyURI() {
		return anyuri.LessAnyURI(this.GetAnyURI(), o.GetAnyURI())
	}
	return false
}

// Name returns the name of this property: "units".
func (this UnitsProperty) Name() string {
	return "units"
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this UnitsProperty) Serialize() (interface{}, error) {
	if this.IsString() {
		return string1.SerializeString(this.GetString())
	} else if this.IsAnyURI() {
		return anyuri.SerializeAnyURI(this.GetAnyURI())
	}
	return string(this.unknown), nil
}

// SetAnyURI sets the value of this property. Calling IsAnyURI afterwards returns
// true.
func (this *UnitsProperty) SetAnyURI(v *url.URL) {
	this.Clear()
	this.anyURIMember = v
}

// SetIRI sets the value of this property. Calling IsIRI afterwards returns true.
func (this *UnitsProperty) SetIRI(v *url.URL) {
	this.Clear()
	this.SetAnyURI(v)
}

// SetString sets the value of this property. Calling IsString afterwards returns
// true.
func (this *UnitsProperty) SetString(v string) {
	this.Clear()
	this.stringMember = v
	this.hasStringMember = true
}

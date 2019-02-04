package propertyid

import (
	"fmt"
	anyuri "github.com/go-fed/activity/streams/values/anyURI"
	vocab "github.com/go-fed/activity/streams/vocab"
	"net/url"
)

// IdProperty is the functional property "id". It is permitted to be a single
// nilable value type.
type IdProperty struct {
	anyURIMember *url.URL
	unknown      []byte
	alias        string
}

// DeserializeIdProperty creates a "id" property from an interface representation
// that has been unmarshalled from a text or binary format.
func DeserializeIdProperty(m map[string]interface{}, aliasMap map[string]string) (*IdProperty, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	propName := "id"
	if len(alias) > 0 {
		// Use alias both to find the property, and set within the property.
		propName = fmt.Sprintf("%s:%s", alias, "id")
	}
	if i, ok := m[propName]; ok {
		if v, err := anyuri.DeserializeAnyURI(i); err == nil {
			this := &IdProperty{
				alias:        alias,
				anyURIMember: v,
			}
			return this, nil
		} else if str, ok := i.(string); ok {
			this := &IdProperty{
				alias:   alias,
				unknown: []byte(str),
			}
			return this, nil
		} else {
			return nil, fmt.Errorf("could not deserialize %q property", "id")
		}
	}
	return nil, nil
}

// NewIdProperty creates a new id property.
func NewIdProperty() *IdProperty {
	return &IdProperty{alias: ""}
}

// Clear ensures no value of this property is set. Calling IsAnyURI afterwards
// will return false.
func (this *IdProperty) Clear() {
	this.unknown = nil
	this.anyURIMember = nil
}

// Get returns the value of this property. When IsAnyURI returns false, Get will
// return any arbitrary value.
func (this IdProperty) Get() *url.URL {
	return this.anyURIMember
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return any arbitrary value.
func (this IdProperty) GetIRI() *url.URL {
	return this.anyURIMember
}

// HasAny returns true if the value or IRI is set.
func (this IdProperty) HasAny() bool {
	return this.IsAnyURI()
}

// IsAnyURI returns true if this property is set and not an IRI.
func (this IdProperty) IsAnyURI() bool {
	return this.anyURIMember != nil
}

// IsIRI returns true if this property is an IRI.
func (this IdProperty) IsIRI() bool {
	return this.anyURIMember != nil
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this IdProperty) JSONLDContext() map[string]string {
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
func (this IdProperty) KindIndex() int {
	if this.IsAnyURI() {
		return 0
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
func (this IdProperty) LessThan(o vocab.IdPropertyInterface) bool {
	if this.IsIRI() {
		// IRIs are always less than other values, none, or unknowns
		return true
	} else if o.IsIRI() {
		// This other, none, or unknown value is always greater than IRIs
		return false
	}
	// LessThan comparison for the single value or unknown value.
	if !this.IsAnyURI() && !o.IsAnyURI() {
		// Both are unknowns.
		return false
	} else if this.IsAnyURI() && !o.IsAnyURI() {
		// Values are always greater than unknown values.
		return false
	} else if !this.IsAnyURI() && o.IsAnyURI() {
		// Unknowns are always less than known values.
		return true
	} else {
		// Actual comparison.
		return anyuri.LessAnyURI(this.Get(), o.Get())
	}
}

// Name returns the name of this property: "id".
func (this IdProperty) Name() string {
	return "id"
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this IdProperty) Serialize() (interface{}, error) {
	if this.IsAnyURI() {
		return anyuri.SerializeAnyURI(this.Get())
	}
	return string(this.unknown), nil
}

// Set sets the value of this property. Calling IsAnyURI afterwards will return
// true.
func (this *IdProperty) Set(v *url.URL) {
	this.Clear()
	this.anyURIMember = v
}

// SetIRI sets the value of this property. Calling IsIRI afterwards will return
// true.
func (this *IdProperty) SetIRI(v *url.URL) {
	this.Clear()
	this.Set(v)
}

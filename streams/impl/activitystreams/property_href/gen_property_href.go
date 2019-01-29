package propertyhref

import (
	"fmt"
	anyuri "github.com/go-fed/activity/streams/values/anyURI"
	vocab "github.com/go-fed/activity/streams/vocab"
	"net/url"
)

// HrefProperty is the functional property "href". It is permitted to be a single
// nilable value type.
type HrefProperty struct {
	anyURIMember *url.URL
	unknown      []byte
	iri          *url.URL
	alias        string
}

// DeserializeHrefProperty creates a "href" property from an interface
// representation that has been unmarshalled from a text or binary format.
func DeserializeHrefProperty(m map[string]interface{}, aliasMap map[string]string) (*HrefProperty, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	propName := "href"
	if len(alias) > 0 {
		// Use alias both to find the property, and set within the property.
		propName = fmt.Sprintf("%s:%s", alias, "href")
	}
	if i, ok := m[propName]; ok {
		if v, err := anyuri.DeserializeAnyURI(i); err != nil {
			this := &HrefProperty{
				alias:        alias,
				anyURIMember: v,
			}
			return this, nil
		} else if v, ok := i.([]byte); ok {
			this := &HrefProperty{
				alias:   alias,
				unknown: v,
			}
			return this, nil
		} else {
			return nil, fmt.Errorf("could not deserialize %q property", "href")
		}
	}
	return nil, nil
}

// NewHrefProperty creates a new href property.
func NewHrefProperty() *HrefProperty {
	return &HrefProperty{alias: ""}
}

// Clear ensures no value of this property is set. Calling IsAnyURI afterwards
// will return false.
func (this *HrefProperty) Clear() {
	this.unknown = nil
	this.iri = nil
	this.anyURIMember = nil
}

// Get returns the value of this property. When IsAnyURI returns false, Get will
// return any arbitrary value.
func (this HrefProperty) Get() *url.URL {
	return this.anyURIMember
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return any arbitrary value.
func (this HrefProperty) GetIRI() *url.URL {
	return this.iri
}

// HasAny returns true if the value or IRI is set.
func (this HrefProperty) HasAny() bool {
	return this.IsAnyURI() || this.iri != nil
}

// IsAnyURI returns true if this property is set and not an IRI.
func (this HrefProperty) IsAnyURI() bool {
	return this.anyURIMember != nil
}

// IsIRI returns true if this property is an IRI.
func (this HrefProperty) IsIRI() bool {
	return this.iri != nil
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this HrefProperty) JSONLDContext() map[string]string {
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
func (this HrefProperty) KindIndex() int {
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
func (this HrefProperty) LessThan(o vocab.HrefPropertyInterface) bool {
	// LessThan comparison for if either or both are IRIs.
	if this.IsIRI() && o.IsIRI() {
		return this.iri.String() < o.GetIRI().String()
	} else if this.IsIRI() {
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

// Name returns the name of this property: "href".
func (this HrefProperty) Name() string {
	return "href"
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this HrefProperty) Serialize() (interface{}, error) {
	if this.IsAnyURI() {
		return anyuri.SerializeAnyURI(this.Get())
	} else if this.IsIRI() {
		return this.iri.String(), nil
	}
	return this.unknown, nil
}

// Set sets the value of this property. Calling IsAnyURI afterwards will return
// true.
func (this *HrefProperty) Set(v *url.URL) {
	this.Clear()
	this.anyURIMember = v
}

// SetIRI sets the value of this property. Calling IsIRI afterwards will return
// true.
func (this *HrefProperty) SetIRI(v *url.URL) {
	this.Clear()
	this.iri = v
}

package propertydescribes

import (
	"fmt"
	vocab "github.com/go-fed/activity/streams/vocab"
	"net/url"
)

// DescribesProperty is the functional property "describes". It is permitted to be
// a single nilable value type.
type DescribesProperty struct {
	ObjectMember vocab.ObjectInterface
	unknown      []byte
	iri          *url.URL
	alias        string
}

// DeserializeDescribesProperty creates a "describes" property from an interface
// representation that has been unmarshalled from a text or binary format.
func DeserializeDescribesProperty(m map[string]interface{}, aliasMap map[string]string) (*DescribesProperty, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	propName := "describes"
	if len(alias) > 0 {
		// Use alias both to find the property, and set within the property.
		propName = fmt.Sprintf("%s:%s", alias, "describes")
	}
	if i, ok := m[propName]; ok {
		if s, ok := i.(string); ok {
			u, err := url.Parse(s)
			// If error exists, don't error out -- skip this and treat as unknown string ([]byte) at worst
			if err == nil {
				this := &DescribesProperty{
					alias: alias,
					iri:   u,
				}
				return this, nil
			}
		}
		if m, ok := i.(map[string]interface{}); ok {
			if v, err := mgr.DeserializeObjectActivityStreams()(m, aliasMap); err != nil {
				this := &DescribesProperty{
					ObjectMember: v,
					alias:        alias,
				}
				return this, nil
			}
		} else if v, ok := i.([]byte); ok {
			this := &DescribesProperty{
				alias:   alias,
				unknown: v,
			}
			return this, nil
		} else {
			return nil, fmt.Errorf("could not deserialize %q property", "describes")
		}
	}
	return nil, nil
}

// NewDescribesProperty creates a new describes property.
func NewDescribesProperty() *DescribesProperty {
	return &DescribesProperty{alias: ""}
}

// Clear ensures no value of this property is set. Calling IsObject afterwards
// will return false.
func (this *DescribesProperty) Clear() {
	this.unknown = nil
	this.iri = nil
	this.ObjectMember = nil
}

// Get returns the value of this property. When IsObject returns false, Get will
// return any arbitrary value.
func (this DescribesProperty) Get() vocab.ObjectInterface {
	return this.ObjectMember
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return any arbitrary value.
func (this DescribesProperty) GetIRI() *url.URL {
	return this.iri
}

// HasAny returns true if the value or IRI is set.
func (this DescribesProperty) HasAny() bool {
	return this.IsObject() || this.iri != nil
}

// IsIRI returns true if this property is an IRI.
func (this DescribesProperty) IsIRI() bool {
	return this.iri != nil
}

// IsObject returns true if this property is set and not an IRI.
func (this DescribesProperty) IsObject() bool {
	return this.ObjectMember != nil
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this DescribesProperty) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/TR/activitystreams-vocabulary": this.alias}
	var child map[string]string
	if this.IsObject() {
		child = this.Get().JSONLDContext()
	}
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
func (this DescribesProperty) KindIndex() int {
	if this.IsObject() {
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
func (this DescribesProperty) LessThan(o vocab.DescribesPropertyInterface) bool {
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
	if !this.IsObject() && !o.IsObject() {
		// Both are unknowns.
		return false
	} else if this.IsObject() && !o.IsObject() {
		// Values are always greater than unknown values.
		return false
	} else if !this.IsObject() && o.IsObject() {
		// Unknowns are always less than known values.
		return true
	} else {
		// Actual comparison.
		return this.Get().LessThan(o.Get())
	}
}

// Name returns the name of this property: "describes".
func (this DescribesProperty) Name() string {
	return "describes"
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this DescribesProperty) Serialize() (interface{}, error) {
	if this.IsObject() {
		return this.Get().Serialize()
	} else if this.IsIRI() {
		return this.iri.String(), nil
	}
	return this.unknown, nil
}

// Set sets the value of this property. Calling IsObject afterwards will return
// true.
func (this *DescribesProperty) Set(v vocab.ObjectInterface) {
	this.Clear()
	this.ObjectMember = v
}

// SetIRI sets the value of this property. Calling IsIRI afterwards will return
// true.
func (this *DescribesProperty) SetIRI(v *url.URL) {
	this.Clear()
	this.iri = v
}

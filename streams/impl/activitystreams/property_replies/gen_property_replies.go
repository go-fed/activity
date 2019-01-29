package propertyreplies

import (
	"fmt"
	vocab "github.com/go-fed/activity/streams/vocab"
	"net/url"
)

// RepliesProperty is the functional property "replies". It is permitted to be a
// single nilable value type.
type RepliesProperty struct {
	CollectionMember vocab.CollectionInterface
	unknown          []byte
	iri              *url.URL
	alias            string
}

// DeserializeRepliesProperty creates a "replies" property from an interface
// representation that has been unmarshalled from a text or binary format.
func DeserializeRepliesProperty(m map[string]interface{}, aliasMap map[string]string) (*RepliesProperty, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	propName := "replies"
	if len(alias) > 0 {
		// Use alias both to find the property, and set within the property.
		propName = fmt.Sprintf("%s:%s", alias, "replies")
	}
	if i, ok := m[propName]; ok {
		if s, ok := i.(string); ok {
			u, err := url.Parse(s)
			// If error exists, don't error out -- skip this and treat as unknown string ([]byte) at worst
			if err == nil {
				this := &RepliesProperty{
					alias: alias,
					iri:   u,
				}
				return this, nil
			}
		}
		if m, ok := i.(map[string]interface{}); ok {
			if v, err := mgr.DeserializeCollectionActivityStreams()(m, aliasMap); err != nil {
				this := &RepliesProperty{
					CollectionMember: v,
					alias:            alias,
				}
				return this, nil
			}
		} else if v, ok := i.([]byte); ok {
			this := &RepliesProperty{
				alias:   alias,
				unknown: v,
			}
			return this, nil
		} else {
			return nil, fmt.Errorf("could not deserialize %q property", "replies")
		}
	}
	return nil, nil
}

// NewRepliesProperty creates a new replies property.
func NewRepliesProperty() *RepliesProperty {
	return &RepliesProperty{alias: ""}
}

// Clear ensures no value of this property is set. Calling IsCollection afterwards
// will return false.
func (this *RepliesProperty) Clear() {
	this.unknown = nil
	this.iri = nil
	this.CollectionMember = nil
}

// Get returns the value of this property. When IsCollection returns false, Get
// will return any arbitrary value.
func (this RepliesProperty) Get() vocab.CollectionInterface {
	return this.CollectionMember
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return any arbitrary value.
func (this RepliesProperty) GetIRI() *url.URL {
	return this.iri
}

// HasAny returns true if the value or IRI is set.
func (this RepliesProperty) HasAny() bool {
	return this.IsCollection() || this.iri != nil
}

// IsCollection returns true if this property is set and not an IRI.
func (this RepliesProperty) IsCollection() bool {
	return this.CollectionMember != nil
}

// IsIRI returns true if this property is an IRI.
func (this RepliesProperty) IsIRI() bool {
	return this.iri != nil
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this RepliesProperty) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/TR/activitystreams-vocabulary": this.alias}
	var child map[string]string
	if this.IsCollection() {
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
func (this RepliesProperty) KindIndex() int {
	if this.IsCollection() {
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
func (this RepliesProperty) LessThan(o vocab.RepliesPropertyInterface) bool {
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
	if !this.IsCollection() && !o.IsCollection() {
		// Both are unknowns.
		return false
	} else if this.IsCollection() && !o.IsCollection() {
		// Values are always greater than unknown values.
		return false
	} else if !this.IsCollection() && o.IsCollection() {
		// Unknowns are always less than known values.
		return true
	} else {
		// Actual comparison.
		return this.Get().LessThan(o.Get())
	}
}

// Name returns the name of this property: "replies".
func (this RepliesProperty) Name() string {
	return "replies"
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this RepliesProperty) Serialize() (interface{}, error) {
	if this.IsCollection() {
		return this.Get().Serialize()
	} else if this.IsIRI() {
		return this.iri.String(), nil
	}
	return this.unknown, nil
}

// Set sets the value of this property. Calling IsCollection afterwards will
// return true.
func (this *RepliesProperty) Set(v vocab.CollectionInterface) {
	this.Clear()
	this.CollectionMember = v
}

// SetIRI sets the value of this property. Calling IsIRI afterwards will return
// true.
func (this *RepliesProperty) SetIRI(v *url.URL) {
	this.Clear()
	this.iri = v
}

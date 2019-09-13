package propertypublickey

import (
	"fmt"
	vocab "github.com/go-fed/activity/streams/vocab"
	"net/url"
)

// ActivityStreamsPublicKeyProperty is the functional property "publicKey". It is
// permitted to be a single nilable value type.
type ActivityStreamsPublicKeyProperty struct {
	activitystreamsPublicKeyMember vocab.ActivityStreamsPublicKey
	unknown                        interface{}
	iri                            *url.URL
	alias                          string
}

// DeserializePublicKeyProperty creates a "publicKey" property from an interface
// representation that has been unmarshalled from a text or binary format.
func DeserializePublicKeyProperty(m map[string]interface{}, aliasMap map[string]string) (*ActivityStreamsPublicKeyProperty, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/ns/activitystreams"]; ok {
		alias = a
	}
	propName := "publicKey"
	if len(alias) > 0 {
		// Use alias both to find the property, and set within the property.
		propName = fmt.Sprintf("%s:%s", alias, "publicKey")
	}
	i, ok := m[propName]

	if ok {
		if s, ok := i.(string); ok {
			u, err := url.Parse(s)
			// If error exists, don't error out -- skip this and treat as unknown string ([]byte) at worst
			// Also, if no scheme exists, don't treat it as a URL -- net/url is greedy
			if err == nil && len(u.Scheme) > 0 {
				this := &ActivityStreamsPublicKeyProperty{
					alias: alias,
					iri:   u,
				}
				return this, nil
			}
		}
		if m, ok := i.(map[string]interface{}); ok {
			if v, err := mgr.DeserializePublicKeyActivityStreams()(m, aliasMap); err == nil {
				this := &ActivityStreamsPublicKeyProperty{
					activitystreamsPublicKeyMember: v,
					alias:                          alias,
				}
				return this, nil
			}
		}
		this := &ActivityStreamsPublicKeyProperty{
			alias:   alias,
			unknown: i,
		}
		return this, nil
	}
	return nil, nil
}

// NewActivityStreamsPublicKeyProperty creates a new publicKey property.
func NewActivityStreamsPublicKeyProperty() *ActivityStreamsPublicKeyProperty {
	return &ActivityStreamsPublicKeyProperty{alias: ""}
}

// Clear ensures no value of this property is set. Calling
// IsActivityStreamsPublicKey afterwards will return false.
func (this *ActivityStreamsPublicKeyProperty) Clear() {
	this.unknown = nil
	this.iri = nil
	this.activitystreamsPublicKeyMember = nil
}

// Get returns the value of this property. When IsActivityStreamsPublicKey returns
// false, Get will return any arbitrary value.
func (this ActivityStreamsPublicKeyProperty) Get() vocab.ActivityStreamsPublicKey {
	return this.activitystreamsPublicKeyMember
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return any arbitrary value.
func (this ActivityStreamsPublicKeyProperty) GetIRI() *url.URL {
	return this.iri
}

// GetType returns the value in this property as a Type. Returns nil if the value
// is not an ActivityStreams type, such as an IRI or another value.
func (this ActivityStreamsPublicKeyProperty) GetType() vocab.Type {
	if this.IsActivityStreamsPublicKey() {
		return this.Get()
	}

	return nil
}

// HasAny returns true if the value or IRI is set.
func (this ActivityStreamsPublicKeyProperty) HasAny() bool {
	return this.IsActivityStreamsPublicKey() || this.iri != nil
}

// IsActivityStreamsPublicKey returns true if this property is set and not an IRI.
func (this ActivityStreamsPublicKeyProperty) IsActivityStreamsPublicKey() bool {
	return this.activitystreamsPublicKeyMember != nil
}

// IsIRI returns true if this property is an IRI.
func (this ActivityStreamsPublicKeyProperty) IsIRI() bool {
	return this.iri != nil
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this ActivityStreamsPublicKeyProperty) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/ns/activitystreams": this.alias}
	var child map[string]string
	if this.IsActivityStreamsPublicKey() {
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
func (this ActivityStreamsPublicKeyProperty) KindIndex() int {
	if this.IsActivityStreamsPublicKey() {
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
func (this ActivityStreamsPublicKeyProperty) LessThan(o vocab.ActivityStreamsPublicKeyProperty) bool {
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
	if !this.IsActivityStreamsPublicKey() && !o.IsActivityStreamsPublicKey() {
		// Both are unknowns.
		return false
	} else if this.IsActivityStreamsPublicKey() && !o.IsActivityStreamsPublicKey() {
		// Values are always greater than unknown values.
		return false
	} else if !this.IsActivityStreamsPublicKey() && o.IsActivityStreamsPublicKey() {
		// Unknowns are always less than known values.
		return true
	} else {
		// Actual comparison.
		return this.Get().LessThan(o.Get())
	}
}

// Name returns the name of this property: "publicKey".
func (this ActivityStreamsPublicKeyProperty) Name() string {
	return "publicKey"
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this ActivityStreamsPublicKeyProperty) Serialize() (interface{}, error) {
	if this.IsActivityStreamsPublicKey() {
		return this.Get().Serialize()
	} else if this.IsIRI() {
		return this.iri.String(), nil
	}
	return this.unknown, nil
}

// Set sets the value of this property. Calling IsActivityStreamsPublicKey
// afterwards will return true.
func (this *ActivityStreamsPublicKeyProperty) Set(v vocab.ActivityStreamsPublicKey) {
	this.Clear()
	this.activitystreamsPublicKeyMember = v
}

// SetIRI sets the value of this property. Calling IsIRI afterwards will return
// true.
func (this *ActivityStreamsPublicKeyProperty) SetIRI(v *url.URL) {
	this.Clear()
	this.iri = v
}

// SetType attempts to set the property for the arbitrary type. Returns an error
// if it is not a valid type to set on this property.
func (this *ActivityStreamsPublicKeyProperty) SetType(t vocab.Type) error {
	if v, ok := t.(vocab.ActivityStreamsPublicKey); ok {
		this.Set(v)
		return nil
	}

	return fmt.Errorf("illegal type to set on publicKey property: %T", t)
}

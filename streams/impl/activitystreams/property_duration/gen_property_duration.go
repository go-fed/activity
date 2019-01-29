package propertyduration

import (
	"fmt"
	duration "github.com/go-fed/activity/streams/values/duration"
	vocab "github.com/go-fed/activity/streams/vocab"
	"net/url"
	"time"
)

// DurationProperty is the functional property "duration". It is permitted to be a
// single default-valued value type.
type DurationProperty struct {
	durationMember    time.Duration
	hasDurationMember bool
	unknown           []byte
	iri               *url.URL
	alias             string
}

// DeserializeDurationProperty creates a "duration" property from an interface
// representation that has been unmarshalled from a text or binary format.
func DeserializeDurationProperty(m map[string]interface{}, aliasMap map[string]string) (*DurationProperty, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	propName := "duration"
	if len(alias) > 0 {
		// Use alias both to find the property, and set within the property.
		propName = fmt.Sprintf("%s:%s", alias, "duration")
	}
	if i, ok := m[propName]; ok {
		if s, ok := i.(string); ok {
			u, err := url.Parse(s)
			// If error exists, don't error out -- skip this and treat as unknown string ([]byte) at worst
			if err == nil {
				this := &DurationProperty{
					alias: alias,
					iri:   u,
				}
				return this, nil
			}
		}
		if v, err := duration.DeserializeDuration(i); err != nil {
			this := &DurationProperty{
				alias:             alias,
				durationMember:    v,
				hasDurationMember: true,
			}
			return this, nil
		} else if v, ok := i.([]byte); ok {
			this := &DurationProperty{
				alias:   alias,
				unknown: v,
			}
			return this, nil
		} else {
			return nil, fmt.Errorf("could not deserialize %q property", "duration")
		}
	}
	return nil, nil
}

// NewDurationProperty creates a new duration property.
func NewDurationProperty() *DurationProperty {
	return &DurationProperty{alias: ""}
}

// Clear ensures no value of this property is set. Calling IsDuration afterwards
// will return false.
func (this *DurationProperty) Clear() {
	this.unknown = nil
	this.iri = nil
	this.hasDurationMember = false
}

// Get returns the value of this property. When IsDuration returns false, Get will
// return any arbitrary value.
func (this DurationProperty) Get() time.Duration {
	return this.durationMember
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return any arbitrary value.
func (this DurationProperty) GetIRI() *url.URL {
	return this.iri
}

// HasAny returns true if the value or IRI is set.
func (this DurationProperty) HasAny() bool {
	return this.IsDuration() || this.iri != nil
}

// IsDuration returns true if this property is set and not an IRI.
func (this DurationProperty) IsDuration() bool {
	return this.hasDurationMember
}

// IsIRI returns true if this property is an IRI.
func (this DurationProperty) IsIRI() bool {
	return this.iri != nil
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this DurationProperty) JSONLDContext() map[string]string {
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
func (this DurationProperty) KindIndex() int {
	if this.IsDuration() {
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
func (this DurationProperty) LessThan(o vocab.DurationPropertyInterface) bool {
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
	if !this.IsDuration() && !o.IsDuration() {
		// Both are unknowns.
		return false
	} else if this.IsDuration() && !o.IsDuration() {
		// Values are always greater than unknown values.
		return false
	} else if !this.IsDuration() && o.IsDuration() {
		// Unknowns are always less than known values.
		return true
	} else {
		// Actual comparison.
		return duration.LessDuration(this.Get(), o.Get())
	}
}

// Name returns the name of this property: "duration".
func (this DurationProperty) Name() string {
	return "duration"
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this DurationProperty) Serialize() (interface{}, error) {
	if this.IsDuration() {
		return duration.SerializeDuration(this.Get())
	} else if this.IsIRI() {
		return this.iri.String(), nil
	}
	return this.unknown, nil
}

// Set sets the value of this property. Calling IsDuration afterwards will return
// true.
func (this *DurationProperty) Set(v time.Duration) {
	this.Clear()
	this.durationMember = v
	this.hasDurationMember = true
}

// SetIRI sets the value of this property. Calling IsIRI afterwards will return
// true.
func (this *DurationProperty) SetIRI(v *url.URL) {
	this.Clear()
	this.iri = v
}

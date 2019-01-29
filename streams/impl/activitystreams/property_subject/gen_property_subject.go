package propertysubject

import (
	"fmt"
	vocab "github.com/go-fed/activity/streams/vocab"
	"net/url"
)

// SubjectProperty is the functional property "subject". It is permitted to be one
// of multiple value types. At most, one type of value can be present, or none
// at all. Setting a value will clear the other types of values so that only
// one of the 'Is' methods will return true. It is possible to clear all
// values, so that this property is empty.
type SubjectProperty struct {
	LinkMember   vocab.LinkInterface
	ObjectMember vocab.ObjectInterface
	unknown      []byte
	iri          *url.URL
	alias        string
}

// DeserializeSubjectProperty creates a "subject" property from an interface
// representation that has been unmarshalled from a text or binary format.
func DeserializeSubjectProperty(m map[string]interface{}, aliasMap map[string]string) (*SubjectProperty, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	propName := "subject"
	if len(alias) > 0 {
		// Use alias both to find the property, and set within the property.
		propName = fmt.Sprintf("%s:%s", alias, "subject")
	}
	if i, ok := m[propName]; ok {
		if s, ok := i.(string); ok {
			u, err := url.Parse(s)
			// If error exists, don't error out -- skip this and treat as unknown string ([]byte) at worst
			if err == nil {
				this := &SubjectProperty{
					alias: alias,
					iri:   u,
				}
				return this, nil
			}
		}
		if m, ok := i.(map[string]interface{}); ok {
			if v, err := mgr.DeserializeLinkActivityStreams()(m, aliasMap); err != nil {
				this := &SubjectProperty{
					LinkMember: v,
					alias:      alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeObjectActivityStreams()(m, aliasMap); err != nil {
				this := &SubjectProperty{
					ObjectMember: v,
					alias:        alias,
				}
				return this, nil
			}
		} else if v, ok := i.([]byte); ok {
			this := &SubjectProperty{
				alias:   alias,
				unknown: v,
			}
			return this, nil
		} else {
			return nil, fmt.Errorf("could not deserialize %q property", "subject")
		}
	}
	return nil, nil
}

// NewSubjectProperty creates a new subject property.
func NewSubjectProperty() *SubjectProperty {
	return &SubjectProperty{alias: ""}
}

// Clear ensures no value of this property is set. Calling HasAny or any of the
// 'Is' methods afterwards will return false.
func (this *SubjectProperty) Clear() {
	this.LinkMember = nil
	this.ObjectMember = nil
	this.unknown = nil
	this.iri = nil
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return an arbitrary value.
func (this SubjectProperty) GetIRI() *url.URL {
	return this.iri
}

// GetLink returns the value of this property. When IsLink returns false, GetLink
// will return an arbitrary value.
func (this SubjectProperty) GetLink() vocab.LinkInterface {
	return this.LinkMember
}

// GetObject returns the value of this property. When IsObject returns false,
// GetObject will return an arbitrary value.
func (this SubjectProperty) GetObject() vocab.ObjectInterface {
	return this.ObjectMember
}

// HasAny returns true if any of the different values is set.
func (this SubjectProperty) HasAny() bool {
	return this.IsLink() ||
		this.IsObject() ||
		this.iri != nil
}

// IsIRI returns true if this property is an IRI. When true, use GetIRI and SetIRI
// to access and set this property
func (this SubjectProperty) IsIRI() bool {
	return this.iri != nil
}

// IsLink returns true if this property has a type of "Link". When true, use the
// GetLink and SetLink methods to access and set this property.
func (this SubjectProperty) IsLink() bool {
	return this.LinkMember != nil
}

// IsObject returns true if this property has a type of "Object". When true, use
// the GetObject and SetObject methods to access and set this property.
func (this SubjectProperty) IsObject() bool {
	return this.ObjectMember != nil
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this SubjectProperty) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/TR/activitystreams-vocabulary": this.alias}
	var child map[string]string
	if this.IsLink() {
		child = this.GetLink().JSONLDContext()
	} else if this.IsObject() {
		child = this.GetObject().JSONLDContext()
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
func (this SubjectProperty) KindIndex() int {
	if this.IsLink() {
		return 0
	}
	if this.IsObject() {
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
func (this SubjectProperty) LessThan(o vocab.SubjectPropertyInterface) bool {
	idx1 := this.KindIndex()
	idx2 := o.KindIndex()
	if idx1 < idx2 {
		return true
	} else if idx1 > idx2 {
		return false
	} else if this.IsLink() {
		return this.GetLink().LessThan(o.GetLink())
	} else if this.IsObject() {
		return this.GetObject().LessThan(o.GetObject())
	} else if this.IsIRI() {
		return this.iri.String() < o.GetIRI().String()
	}
	return false
}

// Name returns the name of this property: "subject".
func (this SubjectProperty) Name() string {
	return "subject"
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this SubjectProperty) Serialize() (interface{}, error) {
	if this.IsLink() {
		return this.GetLink().Serialize()
	} else if this.IsObject() {
		return this.GetObject().Serialize()
	} else if this.IsIRI() {
		return this.iri.String(), nil
	}
	return this.unknown, nil
}

// SetIRI sets the value of this property. Calling IsIRI afterwards returns true.
func (this *SubjectProperty) SetIRI(v *url.URL) {
	this.Clear()
	this.iri = v
}

// SetLink sets the value of this property. Calling IsLink afterwards returns true.
func (this *SubjectProperty) SetLink(v vocab.LinkInterface) {
	this.Clear()
	this.LinkMember = v
}

// SetObject sets the value of this property. Calling IsObject afterwards returns
// true.
func (this *SubjectProperty) SetObject(v vocab.ObjectInterface) {
	this.Clear()
	this.ObjectMember = v
}

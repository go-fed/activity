package vocab

import (
	"net/url"
	"time"
)

// ClosedPropertyIteratorInterface represents a single value for the "closed"
// property.
type ClosedPropertyIteratorInterface interface {
	// GetBoolean returns the value of this property. When IsBoolean returns
	// false, GetBoolean will return an arbitrary value.
	GetBoolean() bool
	// GetDateTime returns the value of this property. When IsDateTime returns
	// false, GetDateTime will return an arbitrary value.
	GetDateTime() time.Time
	// GetIRI returns the IRI of this property. When IsIRI returns false,
	// GetIRI will return an arbitrary value.
	GetIRI() *url.URL
	// GetLink returns the value of this property. When IsLink returns false,
	// GetLink will return an arbitrary value.
	GetLink() LinkInterface
	// GetObject returns the value of this property. When IsObject returns
	// false, GetObject will return an arbitrary value.
	GetObject() ObjectInterface
	// HasAny returns true if any of the different values is set.
	HasAny() bool
	// IsBoolean returns true if this property has a type of "boolean". When
	// true, use the GetBoolean and SetBoolean methods to access and set
	// this property.
	IsBoolean() bool
	// IsDateTime returns true if this property has a type of "dateTime". When
	// true, use the GetDateTime and SetDateTime methods to access and set
	// this property.
	IsDateTime() bool
	// IsIRI returns true if this property is an IRI. When true, use GetIRI
	// and SetIRI to access and set this property
	IsIRI() bool
	// IsLink returns true if this property has a type of "Link". When true,
	// use the GetLink and SetLink methods to access and set this property.
	IsLink() bool
	// IsObject returns true if this property has a type of "Object". When
	// true, use the GetObject and SetObject methods to access and set
	// this property.
	IsObject() bool
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this property and the specific values that are set. The value
	// in the map is the alias used to import the property's value or
	// values.
	JSONLDContext() map[string]string
	// KindIndex computes an arbitrary value for indexing this kind of value.
	// This is a leaky API detail only for folks looking to replace the
	// go-fed implementation. Applications should not use this method.
	KindIndex() int
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o ClosedPropertyIteratorInterface) bool
	// Name returns the name of this property: "closed".
	Name() string
	// Next returns the next iterator, or nil if there is no next iterator.
	Next() ClosedPropertyIteratorInterface
	// Prev returns the previous iterator, or nil if there is no previous
	// iterator.
	Prev() ClosedPropertyIteratorInterface
	// SetBoolean sets the value of this property. Calling IsBoolean
	// afterwards returns true.
	SetBoolean(v bool)
	// SetDateTime sets the value of this property. Calling IsDateTime
	// afterwards returns true.
	SetDateTime(v time.Time)
	// SetIRI sets the value of this property. Calling IsIRI afterwards
	// returns true.
	SetIRI(v *url.URL)
	// SetLink sets the value of this property. Calling IsLink afterwards
	// returns true.
	SetLink(v LinkInterface)
	// SetObject sets the value of this property. Calling IsObject afterwards
	// returns true.
	SetObject(v ObjectInterface)
}

// Indicates that a question has been closed, and answers are no longer accepted.
//
// Example 93 (https://www.w3.org/TR/activitystreams-vocabulary/#ex94b-jsonld):
//   {
//     "closed": "2016-05-10T00:00:00Z",
//     "name": "What is the answer?",
//     "type": "Question"
//   }
type ClosedPropertyInterface interface {
	// AppendBoolean appends a boolean value to the back of a list of the
	// property "closed". Invalidates iterators that are traversing using
	// Prev.
	AppendBoolean(v bool)
	// AppendDateTime appends a dateTime value to the back of a list of the
	// property "closed". Invalidates iterators that are traversing using
	// Prev.
	AppendDateTime(v time.Time)
	// AppendIRI appends an IRI value to the back of a list of the property
	// "closed"
	AppendIRI(v *url.URL)
	// AppendLink appends a Link value to the back of a list of the property
	// "closed". Invalidates iterators that are traversing using Prev.
	AppendLink(v LinkInterface)
	// AppendObject appends a Object value to the back of a list of the
	// property "closed". Invalidates iterators that are traversing using
	// Prev.
	AppendObject(v ObjectInterface)
	// At returns the property value for the specified index. Panics if the
	// index is out of bounds.
	At(index int) ClosedPropertyIteratorInterface
	// Begin returns the first iterator, or nil if empty. Can be used with the
	// iterator's Next method and this property's End method to iterate
	// from front to back through all values.
	Begin() ClosedPropertyIteratorInterface
	// Empty returns returns true if there are no elements.
	Empty() bool
	// End returns beyond-the-last iterator, which is nil. Can be used with
	// the iterator's Next method and this property's Begin method to
	// iterate from front to back through all values.
	End() ClosedPropertyIteratorInterface
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this property and the specific values that are set. The value
	// in the map is the alias used to import the property's value or
	// values.
	JSONLDContext() map[string]string
	// KindIndex computes an arbitrary value for indexing this kind of value.
	// This is a leaky API method specifically needed only for alternate
	// implementations for go-fed. Applications should not use this
	// method. Panics if the index is out of bounds.
	KindIndex(idx int) int
	// Len returns the number of values that exist for the "closed" property.
	Len() (length int)
	// Less computes whether another property is less than this one. Mixing
	// types results in a consistent but arbitrary ordering
	Less(i, j int) bool
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o ClosedPropertyInterface) bool
	// Name returns the name of this property: "closed".
	Name() string
	// PrependBoolean prepends a boolean value to the front of a list of the
	// property "closed". Invalidates all iterators.
	PrependBoolean(v bool)
	// PrependDateTime prepends a dateTime value to the front of a list of the
	// property "closed". Invalidates all iterators.
	PrependDateTime(v time.Time)
	// PrependIRI prepends an IRI value to the front of a list of the property
	// "closed".
	PrependIRI(v *url.URL)
	// PrependLink prepends a Link value to the front of a list of the
	// property "closed". Invalidates all iterators.
	PrependLink(v LinkInterface)
	// PrependObject prepends a Object value to the front of a list of the
	// property "closed". Invalidates all iterators.
	PrependObject(v ObjectInterface)
	// Remove deletes an element at the specified index from a list of the
	// property "closed", regardless of its type. Panics if the index is
	// out of bounds. Invalidates all iterators.
	Remove(idx int)
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format. Applications should not
	// need this function as most typical use cases serialize types
	// instead of individual properties. It is exposed for alternatives to
	// go-fed implementations to use.
	Serialize() (interface{}, error)
	// SetBoolean sets a boolean value to be at the specified index for the
	// property "closed". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetBoolean(idx int, v bool)
	// SetDateTime sets a dateTime value to be at the specified index for the
	// property "closed". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetDateTime(idx int, v time.Time)
	// SetIRI sets an IRI value to be at the specified index for the property
	// "closed". Panics if the index is out of bounds.
	SetIRI(idx int, v *url.URL)
	// SetLink sets a Link value to be at the specified index for the property
	// "closed". Panics if the index is out of bounds. Invalidates all
	// iterators.
	SetLink(idx int, v LinkInterface)
	// SetObject sets a Object value to be at the specified index for the
	// property "closed". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetObject(idx int, v ObjectInterface)
	// Swap swaps the location of values at two indices for the "closed"
	// property.
	Swap(i, j int)
}

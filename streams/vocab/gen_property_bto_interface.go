package vocab

import "net/url"

// BtoPropertyIteratorInterface represents a single value for the "bto" property.
type BtoPropertyIteratorInterface interface {
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
	LessThan(o BtoPropertyIteratorInterface) bool
	// Name returns the name of this property: "bto".
	Name() string
	// Next returns the next iterator, or nil if there is no next iterator.
	Next() BtoPropertyIteratorInterface
	// Prev returns the previous iterator, or nil if there is no previous
	// iterator.
	Prev() BtoPropertyIteratorInterface
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

// Identifies an Object that is part of the private primary audience of this
// Object.
//
// Example 71 (https://www.w3.org/TR/activitystreams-vocabulary/#ex69-jsonld):
//   {
//     "actor": "http://sally.example.org",
//     "bto": "http://joe.example.org",
//     "object": "http://example.org/posts/1",
//     "summary": "Sally offered a post to John",
//     "target": "http://john.example.org",
//     "type": "Offer"
//   }
type BtoPropertyInterface interface {
	// AppendIRI appends an IRI value to the back of a list of the property
	// "bto"
	AppendIRI(v *url.URL)
	// AppendLink appends a Link value to the back of a list of the property
	// "bto". Invalidates iterators that are traversing using Prev.
	AppendLink(v LinkInterface)
	// AppendObject appends a Object value to the back of a list of the
	// property "bto". Invalidates iterators that are traversing using
	// Prev.
	AppendObject(v ObjectInterface)
	// At returns the property value for the specified index. Panics if the
	// index is out of bounds.
	At(index int) BtoPropertyIteratorInterface
	// Begin returns the first iterator, or nil if empty. Can be used with the
	// iterator's Next method and this property's End method to iterate
	// from front to back through all values.
	Begin() BtoPropertyIteratorInterface
	// Empty returns returns true if there are no elements.
	Empty() bool
	// End returns beyond-the-last iterator, which is nil. Can be used with
	// the iterator's Next method and this property's Begin method to
	// iterate from front to back through all values.
	End() BtoPropertyIteratorInterface
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
	// Len returns the number of values that exist for the "bto" property.
	Len() (length int)
	// Less computes whether another property is less than this one. Mixing
	// types results in a consistent but arbitrary ordering
	Less(i, j int) bool
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o BtoPropertyInterface) bool
	// Name returns the name of this property: "bto".
	Name() string
	// PrependIRI prepends an IRI value to the front of a list of the property
	// "bto".
	PrependIRI(v *url.URL)
	// PrependLink prepends a Link value to the front of a list of the
	// property "bto". Invalidates all iterators.
	PrependLink(v LinkInterface)
	// PrependObject prepends a Object value to the front of a list of the
	// property "bto". Invalidates all iterators.
	PrependObject(v ObjectInterface)
	// Remove deletes an element at the specified index from a list of the
	// property "bto", regardless of its type. Panics if the index is out
	// of bounds. Invalidates all iterators.
	Remove(idx int)
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format. Applications should not
	// need this function as most typical use cases serialize types
	// instead of individual properties. It is exposed for alternatives to
	// go-fed implementations to use.
	Serialize() (interface{}, error)
	// SetIRI sets an IRI value to be at the specified index for the property
	// "bto". Panics if the index is out of bounds.
	SetIRI(idx int, v *url.URL)
	// SetLink sets a Link value to be at the specified index for the property
	// "bto". Panics if the index is out of bounds. Invalidates all
	// iterators.
	SetLink(idx int, v LinkInterface)
	// SetObject sets a Object value to be at the specified index for the
	// property "bto". Panics if the index is out of bounds. Invalidates
	// all iterators.
	SetObject(idx int, v ObjectInterface)
	// Swap swaps the location of values at two indices for the "bto" property.
	Swap(i, j int)
}

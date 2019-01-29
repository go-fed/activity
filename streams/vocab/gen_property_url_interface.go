package vocab

import "net/url"

// UrlPropertyIteratorInterface represents a single value for the "url" property.
type UrlPropertyIteratorInterface interface {
	// GetAnyURI returns the value of this property. When IsAnyURI returns
	// false, GetAnyURI will return an arbitrary value.
	GetAnyURI() *url.URL
	// GetIRI returns the IRI of this property. When IsIRI returns false,
	// GetIRI will return an arbitrary value.
	GetIRI() *url.URL
	// GetLink returns the value of this property. When IsLink returns false,
	// GetLink will return an arbitrary value.
	GetLink() LinkInterface
	// HasAny returns true if any of the different values is set.
	HasAny() bool
	// IsAnyURI returns true if this property has a type of "anyURI". When
	// true, use the GetAnyURI and SetAnyURI methods to access and set
	// this property.
	IsAnyURI() bool
	// IsIRI returns true if this property is an IRI. When true, use GetIRI
	// and SetIRI to access and set this property
	IsIRI() bool
	// IsLink returns true if this property has a type of "Link". When true,
	// use the GetLink and SetLink methods to access and set this property.
	IsLink() bool
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
	LessThan(o UrlPropertyIteratorInterface) bool
	// Name returns the name of this property: "url".
	Name() string
	// Next returns the next iterator, or nil if there is no next iterator.
	Next() UrlPropertyIteratorInterface
	// Prev returns the previous iterator, or nil if there is no previous
	// iterator.
	Prev() UrlPropertyIteratorInterface
	// SetAnyURI sets the value of this property. Calling IsAnyURI afterwards
	// returns true.
	SetAnyURI(v *url.URL)
	// SetIRI sets the value of this property. Calling IsIRI afterwards
	// returns true.
	SetIRI(v *url.URL)
	// SetLink sets the value of this property. Calling IsLink afterwards
	// returns true.
	SetLink(v LinkInterface)
}

// Identifies one or more links to representations of the object
//
// Example 109 (https://www.w3.org/TR/activitystreams-vocabulary/#ex124-jsonld):
//   {
//     "name": "4Q Sales Forecast",
//     "type": "Document",
//     "url": "http://example.org/4q-sales-forecast.pdf"
//   }
//
// Example 110 (https://www.w3.org/TR/activitystreams-vocabulary/#ex125-jsonld):
//   {
//     "name": "4Q Sales Forecast",
//     "type": "Document",
//     "url": {
//       "type": "owl:Class",
//       "url": "http://example.org/4q-sales-forecast.pdf"
//     }
//   }
//
// Example 111 (https://www.w3.org/TR/activitystreams-vocabulary/#ex126-jsonld):
//   {
//     "name": "4Q Sales Forecast",
//     "type": "Document",
//     "url": [
//       {
//         "mediaType": "application/pdf",
//         "type": "owl:Class",
//         "url": "http://example.org/4q-sales-forecast.pdf"
//       },
//       {
//         "mediaType": "text/html",
//         "type": "owl:Class",
//         "url": "http://example.org/4q-sales-forecast.html"
//       }
//     ]
//   }
type UrlPropertyInterface interface {
	// AppendAnyURI appends a anyURI value to the back of a list of the
	// property "url". Invalidates iterators that are traversing using
	// Prev.
	AppendAnyURI(v *url.URL)
	// AppendIRI appends an IRI value to the back of a list of the property
	// "url"
	AppendIRI(v *url.URL)
	// AppendLink appends a Link value to the back of a list of the property
	// "url". Invalidates iterators that are traversing using Prev.
	AppendLink(v LinkInterface)
	// At returns the property value for the specified index. Panics if the
	// index is out of bounds.
	At(index int) UrlPropertyIteratorInterface
	// Begin returns the first iterator, or nil if empty. Can be used with the
	// iterator's Next method and this property's End method to iterate
	// from front to back through all values.
	Begin() UrlPropertyIteratorInterface
	// Empty returns returns true if there are no elements.
	Empty() bool
	// End returns beyond-the-last iterator, which is nil. Can be used with
	// the iterator's Next method and this property's Begin method to
	// iterate from front to back through all values.
	End() UrlPropertyIteratorInterface
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
	// Len returns the number of values that exist for the "url" property.
	Len() (length int)
	// Less computes whether another property is less than this one. Mixing
	// types results in a consistent but arbitrary ordering
	Less(i, j int) bool
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o UrlPropertyInterface) bool
	// Name returns the name of this property: "url".
	Name() string
	// PrependAnyURI prepends a anyURI value to the front of a list of the
	// property "url". Invalidates all iterators.
	PrependAnyURI(v *url.URL)
	// PrependIRI prepends an IRI value to the front of a list of the property
	// "url".
	PrependIRI(v *url.URL)
	// PrependLink prepends a Link value to the front of a list of the
	// property "url". Invalidates all iterators.
	PrependLink(v LinkInterface)
	// Remove deletes an element at the specified index from a list of the
	// property "url", regardless of its type. Panics if the index is out
	// of bounds. Invalidates all iterators.
	Remove(idx int)
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format. Applications should not
	// need this function as most typical use cases serialize types
	// instead of individual properties. It is exposed for alternatives to
	// go-fed implementations to use.
	Serialize() (interface{}, error)
	// SetAnyURI sets a anyURI value to be at the specified index for the
	// property "url". Panics if the index is out of bounds. Invalidates
	// all iterators.
	SetAnyURI(idx int, v *url.URL)
	// SetIRI sets an IRI value to be at the specified index for the property
	// "url". Panics if the index is out of bounds.
	SetIRI(idx int, v *url.URL)
	// SetLink sets a Link value to be at the specified index for the property
	// "url". Panics if the index is out of bounds. Invalidates all
	// iterators.
	SetLink(idx int, v LinkInterface)
	// Swap swaps the location of values at two indices for the "url" property.
	Swap(i, j int)
}

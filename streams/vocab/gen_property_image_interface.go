package vocab

import "net/url"

// ImagePropertyIteratorInterface represents a single value for the "image"
// property.
type ImagePropertyIteratorInterface interface {
	// GetIRI returns the IRI of this property. When IsIRI returns false,
	// GetIRI will return an arbitrary value.
	GetIRI() *url.URL
	// GetImage returns the value of this property. When IsImage returns
	// false, GetImage will return an arbitrary value.
	GetImage() ImageInterface
	// GetLink returns the value of this property. When IsLink returns false,
	// GetLink will return an arbitrary value.
	GetLink() LinkInterface
	// HasAny returns true if any of the different values is set.
	HasAny() bool
	// IsIRI returns true if this property is an IRI. When true, use GetIRI
	// and SetIRI to access and set this property
	IsIRI() bool
	// IsImage returns true if this property has a type of "Image". When true,
	// use the GetImage and SetImage methods to access and set this
	// property.
	IsImage() bool
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
	LessThan(o ImagePropertyIteratorInterface) bool
	// Name returns the name of this property: "image".
	Name() string
	// Next returns the next iterator, or nil if there is no next iterator.
	Next() ImagePropertyIteratorInterface
	// Prev returns the previous iterator, or nil if there is no previous
	// iterator.
	Prev() ImagePropertyIteratorInterface
	// SetIRI sets the value of this property. Calling IsIRI afterwards
	// returns true.
	SetIRI(v *url.URL)
	// SetImage sets the value of this property. Calling IsImage afterwards
	// returns true.
	SetImage(v ImageInterface)
	// SetLink sets the value of this property. Calling IsLink afterwards
	// returns true.
	SetLink(v LinkInterface)
}

// Indicates an entity that describes an image for this object. Unlike the icon
// property, there are no aspect ratio or display size limitations assumed.
//
// Example 81 (https://www.w3.org/TR/activitystreams-vocabulary/#ex80-jsonld):
//   {
//     "content": "This is all there is.",
//     "image": {
//       "name": "A Cat",
//       "type": "Image",
//       "url": "http://example.org/cat.png"
//     },
//     "name": "A simple note",
//     "type": "Note"
//   }
//
// Example 82 (https://www.w3.org/TR/activitystreams-vocabulary/#ex81-jsonld):
//   {
//     "content": "This is all there is.",
//     "image": [
//       {
//         "name": "Cat 1",
//         "type": "Image",
//         "url": "http://example.org/cat1.png"
//       },
//       {
//         "name": "Cat 2",
//         "type": "Image",
//         "url": "http://example.org/cat2.png"
//       }
//     ],
//     "name": "A simple note",
//     "type": "Note"
//   }
type ImagePropertyInterface interface {
	// AppendIRI appends an IRI value to the back of a list of the property
	// "image"
	AppendIRI(v *url.URL)
	// AppendImage appends a Image value to the back of a list of the property
	// "image". Invalidates iterators that are traversing using Prev.
	AppendImage(v ImageInterface)
	// AppendLink appends a Link value to the back of a list of the property
	// "image". Invalidates iterators that are traversing using Prev.
	AppendLink(v LinkInterface)
	// At returns the property value for the specified index. Panics if the
	// index is out of bounds.
	At(index int) ImagePropertyIteratorInterface
	// Begin returns the first iterator, or nil if empty. Can be used with the
	// iterator's Next method and this property's End method to iterate
	// from front to back through all values.
	Begin() ImagePropertyIteratorInterface
	// Empty returns returns true if there are no elements.
	Empty() bool
	// End returns beyond-the-last iterator, which is nil. Can be used with
	// the iterator's Next method and this property's Begin method to
	// iterate from front to back through all values.
	End() ImagePropertyIteratorInterface
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
	// Len returns the number of values that exist for the "image" property.
	Len() (length int)
	// Less computes whether another property is less than this one. Mixing
	// types results in a consistent but arbitrary ordering
	Less(i, j int) bool
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o ImagePropertyInterface) bool
	// Name returns the name of this property: "image".
	Name() string
	// PrependIRI prepends an IRI value to the front of a list of the property
	// "image".
	PrependIRI(v *url.URL)
	// PrependImage prepends a Image value to the front of a list of the
	// property "image". Invalidates all iterators.
	PrependImage(v ImageInterface)
	// PrependLink prepends a Link value to the front of a list of the
	// property "image". Invalidates all iterators.
	PrependLink(v LinkInterface)
	// Remove deletes an element at the specified index from a list of the
	// property "image", regardless of its type. Panics if the index is
	// out of bounds. Invalidates all iterators.
	Remove(idx int)
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format. Applications should not
	// need this function as most typical use cases serialize types
	// instead of individual properties. It is exposed for alternatives to
	// go-fed implementations to use.
	Serialize() (interface{}, error)
	// SetIRI sets an IRI value to be at the specified index for the property
	// "image". Panics if the index is out of bounds.
	SetIRI(idx int, v *url.URL)
	// SetImage sets a Image value to be at the specified index for the
	// property "image". Panics if the index is out of bounds. Invalidates
	// all iterators.
	SetImage(idx int, v ImageInterface)
	// SetLink sets a Link value to be at the specified index for the property
	// "image". Panics if the index is out of bounds. Invalidates all
	// iterators.
	SetLink(idx int, v LinkInterface)
	// Swap swaps the location of values at two indices for the "image"
	// property.
	Swap(i, j int)
}

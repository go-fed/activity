package vocab

import "net/url"

// IconPropertyIteratorInterface represents a single value for the "icon" property.
type IconPropertyIteratorInterface interface {
	// GetIRI returns the IRI of this property. When IsIRI returns false,
	// GetIRI will return an arbitrary value.
	GetIRI() *url.URL
	// GetImage returns the value of this property. When IsImage returns
	// false, GetImage will return an arbitrary value.
	GetImage() ImageInterface
	// GetLink returns the value of this property. When IsLink returns false,
	// GetLink will return an arbitrary value.
	GetLink() LinkInterface
	// GetMention returns the value of this property. When IsMention returns
	// false, GetMention will return an arbitrary value.
	GetMention() MentionInterface
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
	// IsMention returns true if this property has a type of "Mention". When
	// true, use the GetMention and SetMention methods to access and set
	// this property.
	IsMention() bool
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
	LessThan(o IconPropertyIteratorInterface) bool
	// Name returns the name of this property: "icon".
	Name() string
	// Next returns the next iterator, or nil if there is no next iterator.
	Next() IconPropertyIteratorInterface
	// Prev returns the previous iterator, or nil if there is no previous
	// iterator.
	Prev() IconPropertyIteratorInterface
	// SetIRI sets the value of this property. Calling IsIRI afterwards
	// returns true.
	SetIRI(v *url.URL)
	// SetImage sets the value of this property. Calling IsImage afterwards
	// returns true.
	SetImage(v ImageInterface)
	// SetLink sets the value of this property. Calling IsLink afterwards
	// returns true.
	SetLink(v LinkInterface)
	// SetMention sets the value of this property. Calling IsMention
	// afterwards returns true.
	SetMention(v MentionInterface)
}

// Indicates an entity that describes an icon for this object. The image should
// have an aspect ratio of one (horizontal) to one (vertical) and should be
// suitable for presentation at a small size.
//
// Example 79 (https://www.w3.org/TR/activitystreams-vocabulary/#ex77-jsonld):
//   {
//     "content": "This is all there is.",
//     "icon": {
//       "height": 16,
//       "name": "Note icon",
//       "type": "Image",
//       "url": "http://example.org/note.png",
//       "width": 16
//     },
//     "summary": "A simple note",
//     "type": "Note"
//   }
//
// Example 80 (https://www.w3.org/TR/activitystreams-vocabulary/#ex78-jsonld):
//   {
//     "content": "A simple note",
//     "icon": [
//       {
//         "height": 16,
//         "summary": "Note (16x16)",
//         "type": "Image",
//         "url": "http://example.org/note1.png",
//         "width": 16
//       },
//       {
//         "height": 32,
//         "summary": "Note (32x32)",
//         "type": "Image",
//         "url": "http://example.org/note2.png",
//         "width": 32
//       }
//     ],
//     "summary": "A simple note",
//     "type": "Note"
//   }
type IconPropertyInterface interface {
	// AppendIRI appends an IRI value to the back of a list of the property
	// "icon"
	AppendIRI(v *url.URL)
	// AppendImage appends a Image value to the back of a list of the property
	// "icon". Invalidates iterators that are traversing using Prev.
	AppendImage(v ImageInterface)
	// AppendLink appends a Link value to the back of a list of the property
	// "icon". Invalidates iterators that are traversing using Prev.
	AppendLink(v LinkInterface)
	// AppendMention appends a Mention value to the back of a list of the
	// property "icon". Invalidates iterators that are traversing using
	// Prev.
	AppendMention(v MentionInterface)
	// At returns the property value for the specified index. Panics if the
	// index is out of bounds.
	At(index int) IconPropertyIteratorInterface
	// Begin returns the first iterator, or nil if empty. Can be used with the
	// iterator's Next method and this property's End method to iterate
	// from front to back through all values.
	Begin() IconPropertyIteratorInterface
	// Empty returns returns true if there are no elements.
	Empty() bool
	// End returns beyond-the-last iterator, which is nil. Can be used with
	// the iterator's Next method and this property's Begin method to
	// iterate from front to back through all values.
	End() IconPropertyIteratorInterface
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
	// Len returns the number of values that exist for the "icon" property.
	Len() (length int)
	// Less computes whether another property is less than this one. Mixing
	// types results in a consistent but arbitrary ordering
	Less(i, j int) bool
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o IconPropertyInterface) bool
	// Name returns the name of this property: "icon".
	Name() string
	// PrependIRI prepends an IRI value to the front of a list of the property
	// "icon".
	PrependIRI(v *url.URL)
	// PrependImage prepends a Image value to the front of a list of the
	// property "icon". Invalidates all iterators.
	PrependImage(v ImageInterface)
	// PrependLink prepends a Link value to the front of a list of the
	// property "icon". Invalidates all iterators.
	PrependLink(v LinkInterface)
	// PrependMention prepends a Mention value to the front of a list of the
	// property "icon". Invalidates all iterators.
	PrependMention(v MentionInterface)
	// Remove deletes an element at the specified index from a list of the
	// property "icon", regardless of its type. Panics if the index is out
	// of bounds. Invalidates all iterators.
	Remove(idx int)
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format. Applications should not
	// need this function as most typical use cases serialize types
	// instead of individual properties. It is exposed for alternatives to
	// go-fed implementations to use.
	Serialize() (interface{}, error)
	// SetIRI sets an IRI value to be at the specified index for the property
	// "icon". Panics if the index is out of bounds.
	SetIRI(idx int, v *url.URL)
	// SetImage sets a Image value to be at the specified index for the
	// property "icon". Panics if the index is out of bounds. Invalidates
	// all iterators.
	SetImage(idx int, v ImageInterface)
	// SetLink sets a Link value to be at the specified index for the property
	// "icon". Panics if the index is out of bounds. Invalidates all
	// iterators.
	SetLink(idx int, v LinkInterface)
	// SetMention sets a Mention value to be at the specified index for the
	// property "icon". Panics if the index is out of bounds. Invalidates
	// all iterators.
	SetMention(idx int, v MentionInterface)
	// Swap swaps the location of values at two indices for the "icon"
	// property.
	Swap(i, j int)
}

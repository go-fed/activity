package vocab

import "net/url"

// ContentPropertyIteratorInterface represents a single value for the "content"
// property.
type ContentPropertyIteratorInterface interface {
	// GetIRI returns the IRI of this property. When IsIRI returns false,
	// GetIRI will return an arbitrary value.
	GetIRI() *url.URL
	// GetLangString returns the value of this property. When IsLangString
	// returns false, GetLangString will return an arbitrary value.
	GetLangString() map[string]string
	// GetLanguage returns the value for the specified BCP47 language code, or
	// an empty string if it is either not a language map or no value is
	// present.
	GetLanguage(bcp47 string) string
	// GetString returns the value of this property. When IsString returns
	// false, GetString will return an arbitrary value.
	GetString() string
	// HasAny returns true if any of the values are set, except for the
	// natural language map. When true, the specific has, getter, and
	// setter methods may be used to determine what kind of value there is
	// to access and set this property. To determine if the property was
	// set as a natural language map, use the IsLangString method instead.
	HasAny() bool
	// HasLanguage returns true if the natural language map has an entry for
	// the specified BCP47 language code.
	HasLanguage(bcp47 string) bool
	// IsIRI returns true if this property is an IRI. When true, use GetIRI
	// and SetIRI to access and set this property
	IsIRI() bool
	// IsLangString returns true if this property has a type of "langString".
	// When true, use the GetLangString and SetLangString methods to
	// access and set this property.. To determine if the property was set
	// as a natural language map, use the IsLangString method instead.
	IsLangString() bool
	// IsString returns true if this property has a type of "string". When
	// true, use the GetString and SetString methods to access and set
	// this property.. To determine if the property was set as a natural
	// language map, use the IsLangString method instead.
	IsString() bool
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
	LessThan(o ContentPropertyIteratorInterface) bool
	// Name returns the name of this property: "content".
	Name() string
	// Next returns the next iterator, or nil if there is no next iterator.
	Next() ContentPropertyIteratorInterface
	// Prev returns the previous iterator, or nil if there is no previous
	// iterator.
	Prev() ContentPropertyIteratorInterface
	// SetIRI sets the value of this property. Calling IsIRI afterwards
	// returns true.
	SetIRI(v *url.URL)
	// SetLangString sets the value of this property and clears the natural
	// language map. Calling IsLangString afterwards will return true.
	// Calling IsLangString afterwards returns false.
	SetLangString(v map[string]string)
	// SetLanguage sets the value for the specified BCP47 language code.
	SetLanguage(bcp47, value string)
	// SetString sets the value of this property and clears the natural
	// language map. Calling IsString afterwards will return true. Calling
	// IsLangString afterwards returns false.
	SetString(v string)
}

// The content or textual representation of the Object encoded as a JSON string.
// By default, the value of content is HTML. The mediaType property can be
// used in the object to indicate a different content type. The content MAY be
// expressed using multiple language-tagged values.
//
// Example 114 (https://www.w3.org/TR/activitystreams-vocabulary/#ex130-jsonld):
//   {
//     "content": "A \u003cem\u003esimple\u003c/em\u003e note",
//     "summary": "A simple note",
//     "type": "Note"
//   }
//
// Example 115 (https://www.w3.org/TR/activitystreams-vocabulary/#ex131-jsonld):
//   {
//     "contentMap": {
//       "en": "A \u003cem\u003esimple\u003c/em\u003e note",
//       "es": "Una nota \u003cem\u003esencilla\u003c/em\u003e",
//       "zh-hans": "一段\u003cem\u003e简单的\u003c/em\u003e笔记"
//     },
//     "summary": "A simple note",
//     "type": "Note"
//   }
//
// Example 116 (https://www.w3.org/TR/activitystreams-vocabulary/#ex130b-jsonld):
//   {
//     "content": "## A simple note\nA simple markdown `note`",
//     "mediaType": "text/markdown",
//     "summary": "A simple note",
//     "type": "Note"
//   }
type ContentPropertyInterface interface {
	// AppendIRI appends an IRI value to the back of a list of the property
	// "content"
	AppendIRI(v *url.URL)
	// AppendLangString appends a langString value to the back of a list of
	// the property "content". Invalidates iterators that are traversing
	// using Prev.
	AppendLangString(v map[string]string)
	// AppendString appends a string value to the back of a list of the
	// property "content". Invalidates iterators that are traversing using
	// Prev.
	AppendString(v string)
	// At returns the property value for the specified index. Panics if the
	// index is out of bounds.
	At(index int) ContentPropertyIteratorInterface
	// Begin returns the first iterator, or nil if empty. Can be used with the
	// iterator's Next method and this property's End method to iterate
	// from front to back through all values.
	Begin() ContentPropertyIteratorInterface
	// Empty returns returns true if there are no elements.
	Empty() bool
	// End returns beyond-the-last iterator, which is nil. Can be used with
	// the iterator's Next method and this property's Begin method to
	// iterate from front to back through all values.
	End() ContentPropertyIteratorInterface
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
	// Len returns the number of values that exist for the "content" property.
	Len() (length int)
	// Less computes whether another property is less than this one. Mixing
	// types results in a consistent but arbitrary ordering
	Less(i, j int) bool
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o ContentPropertyInterface) bool
	// Name returns the name of this property: "content".
	Name() string
	// PrependIRI prepends an IRI value to the front of a list of the property
	// "content".
	PrependIRI(v *url.URL)
	// PrependLangString prepends a langString value to the front of a list of
	// the property "content". Invalidates all iterators.
	PrependLangString(v map[string]string)
	// PrependString prepends a string value to the front of a list of the
	// property "content". Invalidates all iterators.
	PrependString(v string)
	// Remove deletes an element at the specified index from a list of the
	// property "content", regardless of its type. Panics if the index is
	// out of bounds. Invalidates all iterators.
	Remove(idx int)
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format. Applications should not
	// need this function as most typical use cases serialize types
	// instead of individual properties. It is exposed for alternatives to
	// go-fed implementations to use.
	Serialize() (interface{}, error)
	// SetIRI sets an IRI value to be at the specified index for the property
	// "content". Panics if the index is out of bounds.
	SetIRI(idx int, v *url.URL)
	// SetLangString sets a langString value to be at the specified index for
	// the property "content". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetLangString(idx int, v map[string]string)
	// SetString sets a string value to be at the specified index for the
	// property "content". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetString(idx int, v string)
	// Swap swaps the location of values at two indices for the "content"
	// property.
	Swap(i, j int)
}

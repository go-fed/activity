package vocab

import "net/url"

// Identifies a Collection containing objects considered to be responses to this
// object.
//
// Example 104 (https://www.w3.org/TR/activitystreams-vocabulary/#ex112-jsonld):
//   {
//     "content": "I am fine.",
//     "id": "http://www.test.example/notes/1",
//     "replies": {
//       "items": {
//         "content": "I am glad to hear it.",
//         "inReplyTo": "http://www.test.example/notes/1",
//         "summary": "A response to the note",
//         "type": "Note"
//       },
//       "totalItems": 1,
//       "type": "Collection"
//     },
//     "summary": "A simple note",
//     "type": "Note"
//   }
type RepliesPropertyInterface interface {
	// Clear ensures no value of this property is set. Calling IsCollection
	// afterwards will return false.
	Clear()
	// Get returns the value of this property. When IsCollection returns
	// false, Get will return any arbitrary value.
	Get() CollectionInterface
	// GetIRI returns the IRI of this property. When IsIRI returns false,
	// GetIRI will return any arbitrary value.
	GetIRI() *url.URL
	// HasAny returns true if the value or IRI is set.
	HasAny() bool
	// IsCollection returns true if this property is set and not an IRI.
	IsCollection() bool
	// IsIRI returns true if this property is an IRI.
	IsIRI() bool
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
	LessThan(o RepliesPropertyInterface) bool
	// Name returns the name of this property: "replies".
	Name() string
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format. Applications should not
	// need this function as most typical use cases serialize types
	// instead of individual properties. It is exposed for alternatives to
	// go-fed implementations to use.
	Serialize() (interface{}, error)
	// Set sets the value of this property. Calling IsCollection afterwards
	// will return true.
	Set(v CollectionInterface)
	// SetIRI sets the value of this property. Calling IsIRI afterwards will
	// return true.
	SetIRI(v *url.URL)
}

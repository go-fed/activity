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
	// Clear ensures no value of this property is set. Calling HasAny or any
	// of the 'Is' methods afterwards will return false.
	Clear()
	// GetCollection returns the value of this property. When IsCollection
	// returns false, GetCollection will return an arbitrary value.
	GetCollection() CollectionInterface
	// GetCollectionPage returns the value of this property. When
	// IsCollectionPage returns false, GetCollectionPage will return an
	// arbitrary value.
	GetCollectionPage() CollectionPageInterface
	// GetIRI returns the IRI of this property. When IsIRI returns false,
	// GetIRI will return an arbitrary value.
	GetIRI() *url.URL
	// GetOrderedCollection returns the value of this property. When
	// IsOrderedCollection returns false, GetOrderedCollection will return
	// an arbitrary value.
	GetOrderedCollection() OrderedCollectionInterface
	// GetOrderedCollectionPage returns the value of this property. When
	// IsOrderedCollectionPage returns false, GetOrderedCollectionPage
	// will return an arbitrary value.
	GetOrderedCollectionPage() OrderedCollectionPageInterface
	// HasAny returns true if any of the different values is set.
	HasAny() bool
	// IsCollection returns true if this property has a type of "Collection".
	// When true, use the GetCollection and SetCollection methods to
	// access and set this property.
	IsCollection() bool
	// IsCollectionPage returns true if this property has a type of
	// "CollectionPage". When true, use the GetCollectionPage and
	// SetCollectionPage methods to access and set this property.
	IsCollectionPage() bool
	// IsIRI returns true if this property is an IRI. When true, use GetIRI
	// and SetIRI to access and set this property
	IsIRI() bool
	// IsOrderedCollection returns true if this property has a type of
	// "OrderedCollection". When true, use the GetOrderedCollection and
	// SetOrderedCollection methods to access and set this property.
	IsOrderedCollection() bool
	// IsOrderedCollectionPage returns true if this property has a type of
	// "OrderedCollectionPage". When true, use the
	// GetOrderedCollectionPage and SetOrderedCollectionPage methods to
	// access and set this property.
	IsOrderedCollectionPage() bool
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
	// SetCollection sets the value of this property. Calling IsCollection
	// afterwards returns true.
	SetCollection(v CollectionInterface)
	// SetCollectionPage sets the value of this property. Calling
	// IsCollectionPage afterwards returns true.
	SetCollectionPage(v CollectionPageInterface)
	// SetIRI sets the value of this property. Calling IsIRI afterwards
	// returns true.
	SetIRI(v *url.URL)
	// SetOrderedCollection sets the value of this property. Calling
	// IsOrderedCollection afterwards returns true.
	SetOrderedCollection(v OrderedCollectionInterface)
	// SetOrderedCollectionPage sets the value of this property. Calling
	// IsOrderedCollectionPage afterwards returns true.
	SetOrderedCollectionPage(v OrderedCollectionPageInterface)
}

package vocab

import "net/url"

// An ActivityStreams OrderedCollection comprised of all the messages produced by
// the actor
//
// Example 9 (https://www.w3.org/TR/activitypub/#example9):
//   {
//     "@context": [
//       "https://www.w3.org/ns/activitystreams",
//       {
//         "@language": "ja"
//       }
//     ],
//     "followers": "https://kenzoishii.example.com/followers.json",
//     "following": "https://kenzoishii.example.com/following.json",
//     "icon": [
//       "https://kenzoishii.example.com/image/165987aklre4"
//     ],
//     "id": "https://kenzoishii.example.com/",
//     "inbox": "https://kenzoishii.example.com/inbox.json",
//     "liked": "https://kenzoishii.example.com/liked.json",
//     "name": "石井健蔵",
//     "outbox": "https://kenzoishii.example.com/feed.json",
//     "preferredUsername": "kenzoishii",
//     "summary": "この方はただの例です",
//     "type": "Person"
//   }
type OutboxPropertyInterface interface {
	// Clear ensures no value of this property is set. Calling
	// IsOrderedCollection afterwards will return false.
	Clear()
	// Get returns the value of this property. When IsOrderedCollection
	// returns false, Get will return any arbitrary value.
	Get() OrderedCollectionInterface
	// GetIRI returns the IRI of this property. When IsIRI returns false,
	// GetIRI will return any arbitrary value.
	GetIRI() *url.URL
	// HasAny returns true if the value or IRI is set.
	HasAny() bool
	// IsIRI returns true if this property is an IRI.
	IsIRI() bool
	// IsOrderedCollection returns true if this property is set and not an IRI.
	IsOrderedCollection() bool
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
	LessThan(o OutboxPropertyInterface) bool
	// Name returns the name of this property: "outbox".
	Name() string
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format. Applications should not
	// need this function as most typical use cases serialize types
	// instead of individual properties. It is exposed for alternatives to
	// go-fed implementations to use.
	Serialize() (interface{}, error)
	// Set sets the value of this property. Calling IsOrderedCollection
	// afterwards will return true.
	Set(v OrderedCollectionInterface)
	// SetIRI sets the value of this property. Calling IsIRI afterwards will
	// return true.
	SetIRI(v *url.URL)
}

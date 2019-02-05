package vocab

import "net/url"

// StreamsPropertyIteratorInterface represents a single value for the "streams"
// property.
type StreamsPropertyIteratorInterface interface {
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
	LessThan(o StreamsPropertyIteratorInterface) bool
	// Name returns the name of this property: "streams".
	Name() string
	// Next returns the next iterator, or nil if there is no next iterator.
	Next() StreamsPropertyIteratorInterface
	// Prev returns the previous iterator, or nil if there is no previous
	// iterator.
	Prev() StreamsPropertyIteratorInterface
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

// A list of supplementary Collections which may be of interest
type StreamsPropertyInterface interface {
	// AppendCollection appends a Collection value to the back of a list of
	// the property "streams". Invalidates iterators that are traversing
	// using Prev.
	AppendCollection(v CollectionInterface)
	// AppendCollectionPage appends a CollectionPage value to the back of a
	// list of the property "streams". Invalidates iterators that are
	// traversing using Prev.
	AppendCollectionPage(v CollectionPageInterface)
	// AppendIRI appends an IRI value to the back of a list of the property
	// "streams"
	AppendIRI(v *url.URL)
	// AppendOrderedCollection appends a OrderedCollection value to the back
	// of a list of the property "streams". Invalidates iterators that are
	// traversing using Prev.
	AppendOrderedCollection(v OrderedCollectionInterface)
	// AppendOrderedCollectionPage appends a OrderedCollectionPage value to
	// the back of a list of the property "streams". Invalidates iterators
	// that are traversing using Prev.
	AppendOrderedCollectionPage(v OrderedCollectionPageInterface)
	// At returns the property value for the specified index. Panics if the
	// index is out of bounds.
	At(index int) StreamsPropertyIteratorInterface
	// Begin returns the first iterator, or nil if empty. Can be used with the
	// iterator's Next method and this property's End method to iterate
	// from front to back through all values.
	Begin() StreamsPropertyIteratorInterface
	// Empty returns returns true if there are no elements.
	Empty() bool
	// End returns beyond-the-last iterator, which is nil. Can be used with
	// the iterator's Next method and this property's Begin method to
	// iterate from front to back through all values.
	End() StreamsPropertyIteratorInterface
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
	// Len returns the number of values that exist for the "streams" property.
	Len() (length int)
	// Less computes whether another property is less than this one. Mixing
	// types results in a consistent but arbitrary ordering
	Less(i, j int) bool
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o StreamsPropertyInterface) bool
	// Name returns the name of this property: "streams".
	Name() string
	// PrependCollection prepends a Collection value to the front of a list of
	// the property "streams". Invalidates all iterators.
	PrependCollection(v CollectionInterface)
	// PrependCollectionPage prepends a CollectionPage value to the front of a
	// list of the property "streams". Invalidates all iterators.
	PrependCollectionPage(v CollectionPageInterface)
	// PrependIRI prepends an IRI value to the front of a list of the property
	// "streams".
	PrependIRI(v *url.URL)
	// PrependOrderedCollection prepends a OrderedCollection value to the
	// front of a list of the property "streams". Invalidates all
	// iterators.
	PrependOrderedCollection(v OrderedCollectionInterface)
	// PrependOrderedCollectionPage prepends a OrderedCollectionPage value to
	// the front of a list of the property "streams". Invalidates all
	// iterators.
	PrependOrderedCollectionPage(v OrderedCollectionPageInterface)
	// Remove deletes an element at the specified index from a list of the
	// property "streams", regardless of its type. Panics if the index is
	// out of bounds. Invalidates all iterators.
	Remove(idx int)
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format. Applications should not
	// need this function as most typical use cases serialize types
	// instead of individual properties. It is exposed for alternatives to
	// go-fed implementations to use.
	Serialize() (interface{}, error)
	// SetCollection sets a Collection value to be at the specified index for
	// the property "streams". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetCollection(idx int, v CollectionInterface)
	// SetCollectionPage sets a CollectionPage value to be at the specified
	// index for the property "streams". Panics if the index is out of
	// bounds. Invalidates all iterators.
	SetCollectionPage(idx int, v CollectionPageInterface)
	// SetIRI sets an IRI value to be at the specified index for the property
	// "streams". Panics if the index is out of bounds.
	SetIRI(idx int, v *url.URL)
	// SetOrderedCollection sets a OrderedCollection value to be at the
	// specified index for the property "streams". Panics if the index is
	// out of bounds. Invalidates all iterators.
	SetOrderedCollection(idx int, v OrderedCollectionInterface)
	// SetOrderedCollectionPage sets a OrderedCollectionPage value to be at
	// the specified index for the property "streams". Panics if the index
	// is out of bounds. Invalidates all iterators.
	SetOrderedCollectionPage(idx int, v OrderedCollectionPageInterface)
	// Swap swaps the location of values at two indices for the "streams"
	// property.
	Swap(i, j int)
}

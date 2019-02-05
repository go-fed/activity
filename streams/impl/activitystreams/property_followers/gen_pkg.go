package propertyfollowers

import vocab "github.com/go-fed/activity/streams/vocab"

var mgr privateManager

// privateManager abstracts the code-generated manager that provides access to
// concrete implementations.
type privateManager interface {
	// DeserializeCollectionActivityStreams returns the deserialization method
	// for the "CollectionInterface" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeCollectionActivityStreams() func(map[string]interface{}, map[string]string) (vocab.CollectionInterface, error)
	// DeserializeCollectionPageActivityStreams returns the deserialization
	// method for the "CollectionPageInterface" non-functional property in
	// the vocabulary "ActivityStreams"
	DeserializeCollectionPageActivityStreams() func(map[string]interface{}, map[string]string) (vocab.CollectionPageInterface, error)
	// DeserializeOrderedCollectionActivityStreams returns the deserialization
	// method for the "OrderedCollectionInterface" non-functional property
	// in the vocabulary "ActivityStreams"
	DeserializeOrderedCollectionActivityStreams() func(map[string]interface{}, map[string]string) (vocab.OrderedCollectionInterface, error)
	// DeserializeOrderedCollectionPageActivityStreams returns the
	// deserialization method for the "OrderedCollectionPageInterface"
	// non-functional property in the vocabulary "ActivityStreams"
	DeserializeOrderedCollectionPageActivityStreams() func(map[string]interface{}, map[string]string) (vocab.OrderedCollectionPageInterface, error)
}

// SetManager sets the manager package-global variable. For internal use only, do
// not use as part of Application behavior. Must be called at golang init time.
func SetManager(m privateManager) {
	mgr = m
}

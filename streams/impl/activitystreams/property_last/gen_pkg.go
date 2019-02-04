package propertylast

import vocab "github.com/go-fed/activity/streams/vocab"

var mgr privateManager

// privateManager abstracts the code-generated manager that provides access to
// concrete implementations.
type privateManager interface {
	// DeserializeCollectionPageActivityStreams returns the deserialization
	// method for the "CollectionPageInterface" non-functional property in
	// the vocabulary "ActivityStreams"
	DeserializeCollectionPageActivityStreams() func(map[string]interface{}, map[string]string) (vocab.CollectionPageInterface, error)
	// DeserializeLinkActivityStreams returns the deserialization method for
	// the "LinkInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeLinkActivityStreams() func(map[string]interface{}, map[string]string) (vocab.LinkInterface, error)
	// DeserializeMentionActivityStreams returns the deserialization method
	// for the "MentionInterface" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeMentionActivityStreams() func(map[string]interface{}, map[string]string) (vocab.MentionInterface, error)
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

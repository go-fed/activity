package propertyimage

import vocab "github.com/go-fed/activity/streams/vocab"

var mgr privateManager

// privateManager abstracts the code-generated manager that provides access to
// concrete implementations.
type privateManager interface {
	// DeserializeImageActivityStreams returns the deserialization method for
	// the "ImageInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeImageActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ImageInterface, error)
	// DeserializeLinkActivityStreams returns the deserialization method for
	// the "LinkInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeLinkActivityStreams() func(map[string]interface{}, map[string]string) (vocab.LinkInterface, error)
	// DeserializeMentionActivityStreams returns the deserialization method
	// for the "MentionInterface" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeMentionActivityStreams() func(map[string]interface{}, map[string]string) (vocab.MentionInterface, error)
}

// SetManager sets the manager package-global variable. For internal use only, do
// not use as part of Application behavior. Must be called at golang init time.
func SetManager(m privateManager) {
	mgr = m
}

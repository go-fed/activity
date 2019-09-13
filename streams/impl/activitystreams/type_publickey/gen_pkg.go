package typepublickey

import vocab "github.com/go-fed/activity/streams/vocab"

var mgr privateManager

var typePropertyConstructor func() vocab.ActivityStreamsTypeProperty

// privateManager abstracts the code-generated manager that provides access to
// concrete implementations.
type privateManager interface {
	// DeserializeIdPropertyActivityStreams returns the deserialization method
	// for the "ActivityStreamsIdProperty" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeIdPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsIdProperty, error)
	// DeserializeOwnerPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsOwnerProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeOwnerPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsOwnerProperty, error)
	// DeserializePublicKeyPemPropertyActivityStreams returns the
	// deserialization method for the
	// "ActivityStreamsPublicKeyPemProperty" non-functional property in
	// the vocabulary "ActivityStreams"
	DeserializePublicKeyPemPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsPublicKeyPemProperty, error)
	// DeserializeTypePropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsTypeProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeTypePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsTypeProperty, error)
}

// jsonldContexter is a private interface to determine the JSON-LD contexts and
// aliases needed for functional and non-functional properties. It is a helper
// interface for this implementation.
type jsonldContexter interface {
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this property and the specific values that are set. The value
	// in the map is the alias used to import the property's value or
	// values.
	JSONLDContext() map[string]string
}

// SetManager sets the manager package-global variable. For internal use only, do
// not use as part of Application behavior. Must be called at golang init time.
func SetManager(m privateManager) {
	mgr = m
}

// SetTypePropertyConstructor sets the "type" property's constructor in the
// package-global variable. For internal use only, do not use as part of
// Application behavior. Must be called at golang init time. Permits
// ActivityStreams types to correctly set their "type" property at
// construction time, so users don't have to remember to do so each time. It
// is dependency injected so other go-fed compatible implementations could
// inject their own type.
func SetTypePropertyConstructor(f func() vocab.ActivityStreamsTypeProperty) {
	typePropertyConstructor = f
}

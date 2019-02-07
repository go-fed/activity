package typemention

import vocab "github.com/go-fed/activity/streams/vocab"

var mgr privateManager

var typePropertyConstructor func() vocab.TypePropertyInterface

// privateManager abstracts the code-generated manager that provides access to
// concrete implementations.
type privateManager interface {
	// DeserializeAttributedToPropertyActivityStreams returns the
	// deserialization method for the "AttributedToPropertyInterface"
	// non-functional property in the vocabulary "ActivityStreams"
	DeserializeAttributedToPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.AttributedToPropertyInterface, error)
	// DeserializeHeightPropertyActivityStreams returns the deserialization
	// method for the "HeightPropertyInterface" non-functional property in
	// the vocabulary "ActivityStreams"
	DeserializeHeightPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.HeightPropertyInterface, error)
	// DeserializeHrefPropertyActivityStreams returns the deserialization
	// method for the "HrefPropertyInterface" non-functional property in
	// the vocabulary "ActivityStreams"
	DeserializeHrefPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.HrefPropertyInterface, error)
	// DeserializeHreflangPropertyActivityStreams returns the deserialization
	// method for the "HreflangPropertyInterface" non-functional property
	// in the vocabulary "ActivityStreams"
	DeserializeHreflangPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.HreflangPropertyInterface, error)
	// DeserializeIdPropertyActivityStreams returns the deserialization method
	// for the "IdPropertyInterface" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeIdPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.IdPropertyInterface, error)
	// DeserializeMediaTypePropertyActivityStreams returns the deserialization
	// method for the "MediaTypePropertyInterface" non-functional property
	// in the vocabulary "ActivityStreams"
	DeserializeMediaTypePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.MediaTypePropertyInterface, error)
	// DeserializeNamePropertyActivityStreams returns the deserialization
	// method for the "NamePropertyInterface" non-functional property in
	// the vocabulary "ActivityStreams"
	DeserializeNamePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.NamePropertyInterface, error)
	// DeserializePreviewPropertyActivityStreams returns the deserialization
	// method for the "PreviewPropertyInterface" non-functional property
	// in the vocabulary "ActivityStreams"
	DeserializePreviewPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.PreviewPropertyInterface, error)
	// DeserializeRelPropertyActivityStreams returns the deserialization
	// method for the "RelPropertyInterface" non-functional property in
	// the vocabulary "ActivityStreams"
	DeserializeRelPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.RelPropertyInterface, error)
	// DeserializeSummaryPropertyActivityStreams returns the deserialization
	// method for the "SummaryPropertyInterface" non-functional property
	// in the vocabulary "ActivityStreams"
	DeserializeSummaryPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.SummaryPropertyInterface, error)
	// DeserializeTypePropertyActivityStreams returns the deserialization
	// method for the "TypePropertyInterface" non-functional property in
	// the vocabulary "ActivityStreams"
	DeserializeTypePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.TypePropertyInterface, error)
	// DeserializeWidthPropertyActivityStreams returns the deserialization
	// method for the "WidthPropertyInterface" non-functional property in
	// the vocabulary "ActivityStreams"
	DeserializeWidthPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.WidthPropertyInterface, error)
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
func SetTypePropertyConstructor(f func() vocab.TypePropertyInterface) {
	typePropertyConstructor = f
}

package vocab

// Type represents an ActivityStreams type.
type Type interface {
	// GetActivityStreamsId returns the "id" property if it exists, and nil
	// otherwise.
	GetActivityStreamsId() ActivityStreamsIdProperty
	// GetName returns the ActivityStreams type name.
	GetName() string
	// SetActivityStreamsId sets the "id" property.
	SetActivityStreamsId(ActivityStreamsIdProperty)
	// VocabularyURI returns the vocabulary's URI as a string.
	VocabularyURI() string
}

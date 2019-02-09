package vocab

// Type represents an ActivityStreams type.
type Type interface {
	// GetName returns the ActivityStreams type name.
	GetName() string
	// VocabularyURI returns the vocabulary's URI as a string.
	VocabularyURI() string
}

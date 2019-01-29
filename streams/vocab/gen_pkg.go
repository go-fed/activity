package vocab

// Type represents an ActivityStreams type.
type Type interface {
	// GetName returns the ActivityStreams type name.
	GetName() string
}

package pub

import (
	"github.com/go-fed/activity/streams/vocab"
)

// Activity represents any ActivityStreams Activity type.
type Activity interface {
	// Activity is also a vocab.Type
	vocab.Type
	// GetActivityStreamsActor returns the "actor" property if it exists, and
	// nil otherwise.
	GetActivityStreamsActor() vocab.ActivityStreamsActorProperty
	// GetActivityStreamsAudience returns the "audience" property if it
	// exists, and nil otherwise.
	GetActivityStreamsAudience() vocab.ActivityStreamsAudienceProperty
	// GetActivityStreamsBcc returns the "bcc" property if it exists, and nil
	// otherwise.
	GetActivityStreamsBcc() vocab.ActivityStreamsBccProperty
	// GetActivityStreamsBto returns the "bto" property if it exists, and nil
	// otherwise.
	GetActivityStreamsBto() vocab.ActivityStreamsBtoProperty
	// GetActivityStreamsCc returns the "cc" property if it exists, and nil
	// otherwise.
	GetActivityStreamsCc() vocab.ActivityStreamsCcProperty
	// GetActivityStreamsTo returns the "to" property if it exists, and nil
	// otherwise.
	GetActivityStreamsTo() vocab.ActivityStreamsToProperty
	// GetActivityStreamsAttributedTo returns the "attributedTo" property if
	// it exists, and nil otherwise.
	GetActivityStreamsAttributedTo() vocab.ActivityStreamsAttributedToProperty
}

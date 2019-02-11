package pub

import (
	"github.com/go-fed/activity/streams/vocab"
)

// inReplyToer is an ActivityStreams type with a 'inReplyTo' property
type inReplyToer interface {
	GetActivityStreamsInReplyTo() vocab.ActivityStreamsInReplyToProperty
}

// objecter is an ActivityStreams type with a 'object' property
type objecter interface {
	GetActivityStreamsObject() vocab.ActivityStreamsObjectProperty
}

// targeter is an ActivityStreams type with a 'target' property
type targeter interface {
	GetActivityStreamsTarget() vocab.ActivityStreamsTargetProperty
}

// tagger is an ActivityStreams type with a 'tag' property
type tagger interface {
	GetActivityStreamsTag() vocab.ActivityStreamsTagProperty
}

// hrefer is an ActivityStreams type with a 'href' property
type hrefer interface {
	GetActivityStreamsHref() vocab.ActivityStreamsHrefProperty
}

// itemser is an ActivityStreams type with a 'items' property
type itemser interface {
	GetActivityStreamsItems() vocab.ActivityStreamsItemsProperty
}

// orderedItemser is an ActivityStreams type with a 'orderedItems' property
type orderedItemser interface {
	GetActivityStreamsOrderedItems() vocab.ActivityStreamsOrderedItemsProperty
}

// publisheder is an ActivityStreams type with a 'published' property
type publisheder interface {
	GetActivityStreamsPublished() vocab.ActivityStreamsPublishedProperty
}

// toer is an ActivityStreams type with a 'to' property
type toer interface {
	GetActivityStreamsTo() vocab.ActivityStreamsToProperty
}

// btoer is an ActivityStreams type with a 'bto' property
type btoer interface {
	GetActivityStreamsBto() vocab.ActivityStreamsBtoProperty
}

// ccer is an ActivityStreams type with a 'cc' property
type ccer interface {
	GetActivityStreamsCc() vocab.ActivityStreamsCcProperty
}

// bccer is an ActivityStreams type with a 'bcc' property
type bccer interface {
	GetActivityStreamsBcc() vocab.ActivityStreamsBccProperty
}

// audiencer is an ActivityStreams type with a 'audience' property
type audiencer interface {
	GetActivityStreamsAudience() vocab.ActivityStreamsAudienceProperty
}

// inboxer is an ActivityStreams type with a 'inbox' property
type inboxer interface {
	GetActivityStreamsInbox() vocab.ActivityStreamsInboxProperty
}

// attributedToer is an ActivityStreams type with a 'attributedTo' property
type attributedToer interface {
	GetActivityStreamsAttributedTo() vocab.ActivityStreamsAttributedToProperty
}

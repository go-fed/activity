package vocab

// Represents a formal or informal collective of Actors.
//
// Example 43 (https://www.w3.org/TR/activitystreams-vocabulary/#ex37-jsonld):
//   {
//     "name": "Big Beards of Austin",
//     "type": "Group"
//   }
type ActivityStreamsGroup interface {
	// GetActivityStreamsAltitude returns the "altitude" property if it
	// exists, and nil otherwise.
	GetActivityStreamsAltitude() ActivityStreamsAltitudeProperty
	// GetActivityStreamsAttachment returns the "attachment" property if it
	// exists, and nil otherwise.
	GetActivityStreamsAttachment() ActivityStreamsAttachmentProperty
	// GetActivityStreamsAttributedTo returns the "attributedTo" property if
	// it exists, and nil otherwise.
	GetActivityStreamsAttributedTo() ActivityStreamsAttributedToProperty
	// GetActivityStreamsAudience returns the "audience" property if it
	// exists, and nil otherwise.
	GetActivityStreamsAudience() ActivityStreamsAudienceProperty
	// GetActivityStreamsBcc returns the "bcc" property if it exists, and nil
	// otherwise.
	GetActivityStreamsBcc() ActivityStreamsBccProperty
	// GetActivityStreamsBto returns the "bto" property if it exists, and nil
	// otherwise.
	GetActivityStreamsBto() ActivityStreamsBtoProperty
	// GetActivityStreamsCc returns the "cc" property if it exists, and nil
	// otherwise.
	GetActivityStreamsCc() ActivityStreamsCcProperty
	// GetActivityStreamsContent returns the "content" property if it exists,
	// and nil otherwise.
	GetActivityStreamsContent() ActivityStreamsContentProperty
	// GetActivityStreamsContext returns the "context" property if it exists,
	// and nil otherwise.
	GetActivityStreamsContext() ActivityStreamsContextProperty
	// GetActivityStreamsDuration returns the "duration" property if it
	// exists, and nil otherwise.
	GetActivityStreamsDuration() ActivityStreamsDurationProperty
	// GetActivityStreamsEndTime returns the "endTime" property if it exists,
	// and nil otherwise.
	GetActivityStreamsEndTime() ActivityStreamsEndTimeProperty
	// GetActivityStreamsFollowers returns the "followers" property if it
	// exists, and nil otherwise.
	GetActivityStreamsFollowers() ActivityStreamsFollowersProperty
	// GetActivityStreamsFollowing returns the "following" property if it
	// exists, and nil otherwise.
	GetActivityStreamsFollowing() ActivityStreamsFollowingProperty
	// GetActivityStreamsGenerator returns the "generator" property if it
	// exists, and nil otherwise.
	GetActivityStreamsGenerator() ActivityStreamsGeneratorProperty
	// GetActivityStreamsIcon returns the "icon" property if it exists, and
	// nil otherwise.
	GetActivityStreamsIcon() ActivityStreamsIconProperty
	// GetActivityStreamsId returns the "id" property if it exists, and nil
	// otherwise.
	GetActivityStreamsId() ActivityStreamsIdProperty
	// GetActivityStreamsImage returns the "image" property if it exists, and
	// nil otherwise.
	GetActivityStreamsImage() ActivityStreamsImageProperty
	// GetActivityStreamsInReplyTo returns the "inReplyTo" property if it
	// exists, and nil otherwise.
	GetActivityStreamsInReplyTo() ActivityStreamsInReplyToProperty
	// GetActivityStreamsInbox returns the "inbox" property if it exists, and
	// nil otherwise.
	GetActivityStreamsInbox() ActivityStreamsInboxProperty
	// GetActivityStreamsLiked returns the "liked" property if it exists, and
	// nil otherwise.
	GetActivityStreamsLiked() ActivityStreamsLikedProperty
	// GetActivityStreamsLikes returns the "likes" property if it exists, and
	// nil otherwise.
	GetActivityStreamsLikes() ActivityStreamsLikesProperty
	// GetActivityStreamsLocation returns the "location" property if it
	// exists, and nil otherwise.
	GetActivityStreamsLocation() ActivityStreamsLocationProperty
	// GetActivityStreamsMediaType returns the "mediaType" property if it
	// exists, and nil otherwise.
	GetActivityStreamsMediaType() ActivityStreamsMediaTypeProperty
	// GetActivityStreamsName returns the "name" property if it exists, and
	// nil otherwise.
	GetActivityStreamsName() ActivityStreamsNameProperty
	// GetActivityStreamsObject returns the "object" property if it exists,
	// and nil otherwise.
	GetActivityStreamsObject() ActivityStreamsObjectProperty
	// GetActivityStreamsOutbox returns the "outbox" property if it exists,
	// and nil otherwise.
	GetActivityStreamsOutbox() ActivityStreamsOutboxProperty
	// GetActivityStreamsPreferredUsername returns the "preferredUsername"
	// property if it exists, and nil otherwise.
	GetActivityStreamsPreferredUsername() ActivityStreamsPreferredUsernameProperty
	// GetActivityStreamsPreview returns the "preview" property if it exists,
	// and nil otherwise.
	GetActivityStreamsPreview() ActivityStreamsPreviewProperty
	// GetActivityStreamsPublished returns the "published" property if it
	// exists, and nil otherwise.
	GetActivityStreamsPublished() ActivityStreamsPublishedProperty
	// GetActivityStreamsReplies returns the "replies" property if it exists,
	// and nil otherwise.
	GetActivityStreamsReplies() ActivityStreamsRepliesProperty
	// GetActivityStreamsShares returns the "shares" property if it exists,
	// and nil otherwise.
	GetActivityStreamsShares() ActivityStreamsSharesProperty
	// GetActivityStreamsStartTime returns the "startTime" property if it
	// exists, and nil otherwise.
	GetActivityStreamsStartTime() ActivityStreamsStartTimeProperty
	// GetActivityStreamsStreams returns the "streams" property if it exists,
	// and nil otherwise.
	GetActivityStreamsStreams() ActivityStreamsStreamsProperty
	// GetActivityStreamsSummary returns the "summary" property if it exists,
	// and nil otherwise.
	GetActivityStreamsSummary() ActivityStreamsSummaryProperty
	// GetActivityStreamsTag returns the "tag" property if it exists, and nil
	// otherwise.
	GetActivityStreamsTag() ActivityStreamsTagProperty
	// GetActivityStreamsTo returns the "to" property if it exists, and nil
	// otherwise.
	GetActivityStreamsTo() ActivityStreamsToProperty
	// GetActivityStreamsType returns the "type" property if it exists, and
	// nil otherwise.
	GetActivityStreamsType() ActivityStreamsTypeProperty
	// GetActivityStreamsUpdated returns the "updated" property if it exists,
	// and nil otherwise.
	GetActivityStreamsUpdated() ActivityStreamsUpdatedProperty
	// GetActivityStreamsUrl returns the "url" property if it exists, and nil
	// otherwise.
	GetActivityStreamsUrl() ActivityStreamsUrlProperty
	// GetName returns the name of this type.
	GetName() string
	// GetUnknownProperties returns the unknown properties for the Group type.
	// Note that this should not be used by app developers. It is only
	// used to help determine which implementation is LessThan the other.
	// Developers who are creating a different implementation of this
	// type's interface can use this method in their LessThan
	// implementation, but routine ActivityPub applications should not use
	// this to bypass the code generation tool.
	GetUnknownProperties() map[string]interface{}
	// IsExtending returns true if the Group type extends from the other type.
	IsExtending(other Type) bool
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this type and the specific properties that are set. The value
	// in the map is the alias used to import the type and its properties.
	JSONLDContext() map[string]string
	// LessThan computes if this Group is lesser, with an arbitrary but stable
	// determination.
	LessThan(o ActivityStreamsGroup) bool
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format.
	Serialize() (map[string]interface{}, error)
	// SetActivityStreamsAltitude returns the "altitude" property if it
	// exists, and nil otherwise.
	SetActivityStreamsAltitude(i ActivityStreamsAltitudeProperty)
	// SetActivityStreamsAttachment returns the "attachment" property if it
	// exists, and nil otherwise.
	SetActivityStreamsAttachment(i ActivityStreamsAttachmentProperty)
	// SetActivityStreamsAttributedTo returns the "attributedTo" property if
	// it exists, and nil otherwise.
	SetActivityStreamsAttributedTo(i ActivityStreamsAttributedToProperty)
	// SetActivityStreamsAudience returns the "audience" property if it
	// exists, and nil otherwise.
	SetActivityStreamsAudience(i ActivityStreamsAudienceProperty)
	// SetActivityStreamsBcc returns the "bcc" property if it exists, and nil
	// otherwise.
	SetActivityStreamsBcc(i ActivityStreamsBccProperty)
	// SetActivityStreamsBto returns the "bto" property if it exists, and nil
	// otherwise.
	SetActivityStreamsBto(i ActivityStreamsBtoProperty)
	// SetActivityStreamsCc returns the "cc" property if it exists, and nil
	// otherwise.
	SetActivityStreamsCc(i ActivityStreamsCcProperty)
	// SetActivityStreamsContent returns the "content" property if it exists,
	// and nil otherwise.
	SetActivityStreamsContent(i ActivityStreamsContentProperty)
	// SetActivityStreamsContext returns the "context" property if it exists,
	// and nil otherwise.
	SetActivityStreamsContext(i ActivityStreamsContextProperty)
	// SetActivityStreamsDuration returns the "duration" property if it
	// exists, and nil otherwise.
	SetActivityStreamsDuration(i ActivityStreamsDurationProperty)
	// SetActivityStreamsEndTime returns the "endTime" property if it exists,
	// and nil otherwise.
	SetActivityStreamsEndTime(i ActivityStreamsEndTimeProperty)
	// SetActivityStreamsFollowers returns the "followers" property if it
	// exists, and nil otherwise.
	SetActivityStreamsFollowers(i ActivityStreamsFollowersProperty)
	// SetActivityStreamsFollowing returns the "following" property if it
	// exists, and nil otherwise.
	SetActivityStreamsFollowing(i ActivityStreamsFollowingProperty)
	// SetActivityStreamsGenerator returns the "generator" property if it
	// exists, and nil otherwise.
	SetActivityStreamsGenerator(i ActivityStreamsGeneratorProperty)
	// SetActivityStreamsIcon returns the "icon" property if it exists, and
	// nil otherwise.
	SetActivityStreamsIcon(i ActivityStreamsIconProperty)
	// SetActivityStreamsId returns the "id" property if it exists, and nil
	// otherwise.
	SetActivityStreamsId(i ActivityStreamsIdProperty)
	// SetActivityStreamsImage returns the "image" property if it exists, and
	// nil otherwise.
	SetActivityStreamsImage(i ActivityStreamsImageProperty)
	// SetActivityStreamsInReplyTo returns the "inReplyTo" property if it
	// exists, and nil otherwise.
	SetActivityStreamsInReplyTo(i ActivityStreamsInReplyToProperty)
	// SetActivityStreamsInbox returns the "inbox" property if it exists, and
	// nil otherwise.
	SetActivityStreamsInbox(i ActivityStreamsInboxProperty)
	// SetActivityStreamsLiked returns the "liked" property if it exists, and
	// nil otherwise.
	SetActivityStreamsLiked(i ActivityStreamsLikedProperty)
	// SetActivityStreamsLikes returns the "likes" property if it exists, and
	// nil otherwise.
	SetActivityStreamsLikes(i ActivityStreamsLikesProperty)
	// SetActivityStreamsLocation returns the "location" property if it
	// exists, and nil otherwise.
	SetActivityStreamsLocation(i ActivityStreamsLocationProperty)
	// SetActivityStreamsMediaType returns the "mediaType" property if it
	// exists, and nil otherwise.
	SetActivityStreamsMediaType(i ActivityStreamsMediaTypeProperty)
	// SetActivityStreamsName returns the "name" property if it exists, and
	// nil otherwise.
	SetActivityStreamsName(i ActivityStreamsNameProperty)
	// SetActivityStreamsObject returns the "object" property if it exists,
	// and nil otherwise.
	SetActivityStreamsObject(i ActivityStreamsObjectProperty)
	// SetActivityStreamsOutbox returns the "outbox" property if it exists,
	// and nil otherwise.
	SetActivityStreamsOutbox(i ActivityStreamsOutboxProperty)
	// SetActivityStreamsPreferredUsername returns the "preferredUsername"
	// property if it exists, and nil otherwise.
	SetActivityStreamsPreferredUsername(i ActivityStreamsPreferredUsernameProperty)
	// SetActivityStreamsPreview returns the "preview" property if it exists,
	// and nil otherwise.
	SetActivityStreamsPreview(i ActivityStreamsPreviewProperty)
	// SetActivityStreamsPublished returns the "published" property if it
	// exists, and nil otherwise.
	SetActivityStreamsPublished(i ActivityStreamsPublishedProperty)
	// SetActivityStreamsReplies returns the "replies" property if it exists,
	// and nil otherwise.
	SetActivityStreamsReplies(i ActivityStreamsRepliesProperty)
	// SetActivityStreamsShares returns the "shares" property if it exists,
	// and nil otherwise.
	SetActivityStreamsShares(i ActivityStreamsSharesProperty)
	// SetActivityStreamsStartTime returns the "startTime" property if it
	// exists, and nil otherwise.
	SetActivityStreamsStartTime(i ActivityStreamsStartTimeProperty)
	// SetActivityStreamsStreams returns the "streams" property if it exists,
	// and nil otherwise.
	SetActivityStreamsStreams(i ActivityStreamsStreamsProperty)
	// SetActivityStreamsSummary returns the "summary" property if it exists,
	// and nil otherwise.
	SetActivityStreamsSummary(i ActivityStreamsSummaryProperty)
	// SetActivityStreamsTag returns the "tag" property if it exists, and nil
	// otherwise.
	SetActivityStreamsTag(i ActivityStreamsTagProperty)
	// SetActivityStreamsTo returns the "to" property if it exists, and nil
	// otherwise.
	SetActivityStreamsTo(i ActivityStreamsToProperty)
	// SetActivityStreamsType returns the "type" property if it exists, and
	// nil otherwise.
	SetActivityStreamsType(i ActivityStreamsTypeProperty)
	// SetActivityStreamsUpdated returns the "updated" property if it exists,
	// and nil otherwise.
	SetActivityStreamsUpdated(i ActivityStreamsUpdatedProperty)
	// SetActivityStreamsUrl returns the "url" property if it exists, and nil
	// otherwise.
	SetActivityStreamsUrl(i ActivityStreamsUrlProperty)
	// VocabularyURI returns the vocabulary's URI as a string.
	VocabularyURI() string
}

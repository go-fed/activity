package vocab

// A Collection is a subtype of Object that represents ordered or unordered sets
// of Object or Link instances. Refer to the Activity Streams 2.0 Core
// specification for a complete description of the Collection type.
//
// Example 5 (https://www.w3.org/TR/activitystreams-vocabulary/#ex5-jsonld):
//   {
//     "items": [
//       {
//         "name": "A Simple Note",
//         "type": "Note"
//       },
//       {
//         "name": "Another Simple Note",
//         "type": "Note"
//       }
//     ],
//     "summary": "Sally's notes",
//     "totalItems": 2,
//     "type": "Collection"
//   }
type ActivityStreamsCollection interface {
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
	// GetActivityStreamsCurrent returns the "current" property if it exists,
	// and nil otherwise.
	GetActivityStreamsCurrent() ActivityStreamsCurrentProperty
	// GetActivityStreamsDuration returns the "duration" property if it
	// exists, and nil otherwise.
	GetActivityStreamsDuration() ActivityStreamsDurationProperty
	// GetActivityStreamsEndTime returns the "endTime" property if it exists,
	// and nil otherwise.
	GetActivityStreamsEndTime() ActivityStreamsEndTimeProperty
	// GetActivityStreamsFirst returns the "first" property if it exists, and
	// nil otherwise.
	GetActivityStreamsFirst() ActivityStreamsFirstProperty
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
	// GetActivityStreamsItems returns the "items" property if it exists, and
	// nil otherwise.
	GetActivityStreamsItems() ActivityStreamsItemsProperty
	// GetActivityStreamsLast returns the "last" property if it exists, and
	// nil otherwise.
	GetActivityStreamsLast() ActivityStreamsLastProperty
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
	// GetActivityStreamsSummary returns the "summary" property if it exists,
	// and nil otherwise.
	GetActivityStreamsSummary() ActivityStreamsSummaryProperty
	// GetActivityStreamsTag returns the "tag" property if it exists, and nil
	// otherwise.
	GetActivityStreamsTag() ActivityStreamsTagProperty
	// GetActivityStreamsTo returns the "to" property if it exists, and nil
	// otherwise.
	GetActivityStreamsTo() ActivityStreamsToProperty
	// GetActivityStreamsTotalItems returns the "totalItems" property if it
	// exists, and nil otherwise.
	GetActivityStreamsTotalItems() ActivityStreamsTotalItemsProperty
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
	// GetUnknownProperties returns the unknown properties for the Collection
	// type. Note that this should not be used by app developers. It is
	// only used to help determine which implementation is LessThan the
	// other. Developers who are creating a different implementation of
	// this type's interface can use this method in their LessThan
	// implementation, but routine ActivityPub applications should not use
	// this to bypass the code generation tool.
	GetUnknownProperties() map[string]interface{}
	// IsExtending returns true if the Collection type extends from the other
	// type.
	IsExtending(other Type) bool
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this type and the specific properties that are set. The value
	// in the map is the alias used to import the type and its properties.
	JSONLDContext() map[string]string
	// LessThan computes if this Collection is lesser, with an arbitrary but
	// stable determination.
	LessThan(o ActivityStreamsCollection) bool
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
	// SetActivityStreamsCurrent returns the "current" property if it exists,
	// and nil otherwise.
	SetActivityStreamsCurrent(i ActivityStreamsCurrentProperty)
	// SetActivityStreamsDuration returns the "duration" property if it
	// exists, and nil otherwise.
	SetActivityStreamsDuration(i ActivityStreamsDurationProperty)
	// SetActivityStreamsEndTime returns the "endTime" property if it exists,
	// and nil otherwise.
	SetActivityStreamsEndTime(i ActivityStreamsEndTimeProperty)
	// SetActivityStreamsFirst returns the "first" property if it exists, and
	// nil otherwise.
	SetActivityStreamsFirst(i ActivityStreamsFirstProperty)
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
	// SetActivityStreamsItems returns the "items" property if it exists, and
	// nil otherwise.
	SetActivityStreamsItems(i ActivityStreamsItemsProperty)
	// SetActivityStreamsLast returns the "last" property if it exists, and
	// nil otherwise.
	SetActivityStreamsLast(i ActivityStreamsLastProperty)
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
	// SetActivityStreamsSummary returns the "summary" property if it exists,
	// and nil otherwise.
	SetActivityStreamsSummary(i ActivityStreamsSummaryProperty)
	// SetActivityStreamsTag returns the "tag" property if it exists, and nil
	// otherwise.
	SetActivityStreamsTag(i ActivityStreamsTagProperty)
	// SetActivityStreamsTo returns the "to" property if it exists, and nil
	// otherwise.
	SetActivityStreamsTo(i ActivityStreamsToProperty)
	// SetActivityStreamsTotalItems returns the "totalItems" property if it
	// exists, and nil otherwise.
	SetActivityStreamsTotalItems(i ActivityStreamsTotalItemsProperty)
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

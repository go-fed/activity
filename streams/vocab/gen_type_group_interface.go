package vocab

// Represents a formal or informal collective of Actors.
//
// Example 43 (https://www.w3.org/TR/activitystreams-vocabulary/#ex37-jsonld):
//   {
//     "name": "Big Beards of Austin",
//     "type": "Group"
//   }
type GroupInterface interface {
	// GetAltitude returns the "altitude" property if it exists, and nil
	// otherwise.
	GetAltitude() AltitudePropertyInterface
	// GetAttachment returns the "attachment" property if it exists, and nil
	// otherwise.
	GetAttachment() AttachmentPropertyInterface
	// GetAttributedTo returns the "attributedTo" property if it exists, and
	// nil otherwise.
	GetAttributedTo() AttributedToPropertyInterface
	// GetAudience returns the "audience" property if it exists, and nil
	// otherwise.
	GetAudience() AudiencePropertyInterface
	// GetBcc returns the "bcc" property if it exists, and nil otherwise.
	GetBcc() BccPropertyInterface
	// GetBto returns the "bto" property if it exists, and nil otherwise.
	GetBto() BtoPropertyInterface
	// GetCc returns the "cc" property if it exists, and nil otherwise.
	GetCc() CcPropertyInterface
	// GetContent returns the "content" property if it exists, and nil
	// otherwise.
	GetContent() ContentPropertyInterface
	// GetContext returns the "context" property if it exists, and nil
	// otherwise.
	GetContext() ContextPropertyInterface
	// GetDuration returns the "duration" property if it exists, and nil
	// otherwise.
	GetDuration() DurationPropertyInterface
	// GetEndTime returns the "endTime" property if it exists, and nil
	// otherwise.
	GetEndTime() EndTimePropertyInterface
	// GetFollowers returns the "followers" property if it exists, and nil
	// otherwise.
	GetFollowers() FollowersPropertyInterface
	// GetFollowing returns the "following" property if it exists, and nil
	// otherwise.
	GetFollowing() FollowingPropertyInterface
	// GetGenerator returns the "generator" property if it exists, and nil
	// otherwise.
	GetGenerator() GeneratorPropertyInterface
	// GetIcon returns the "icon" property if it exists, and nil otherwise.
	GetIcon() IconPropertyInterface
	// GetId returns the "id" property if it exists, and nil otherwise.
	GetId() IdPropertyInterface
	// GetImage returns the "image" property if it exists, and nil otherwise.
	GetImage() ImagePropertyInterface
	// GetInReplyTo returns the "inReplyTo" property if it exists, and nil
	// otherwise.
	GetInReplyTo() InReplyToPropertyInterface
	// GetInbox returns the "inbox" property if it exists, and nil otherwise.
	GetInbox() InboxPropertyInterface
	// GetLiked returns the "liked" property if it exists, and nil otherwise.
	GetLiked() LikedPropertyInterface
	// GetLikes returns the "likes" property if it exists, and nil otherwise.
	GetLikes() LikesPropertyInterface
	// GetLocation returns the "location" property if it exists, and nil
	// otherwise.
	GetLocation() LocationPropertyInterface
	// GetMediaType returns the "mediaType" property if it exists, and nil
	// otherwise.
	GetMediaType() MediaTypePropertyInterface
	// GetName returns the "name" property if it exists, and nil otherwise.
	GetName() NamePropertyInterface
	// GetObject returns the "object" property if it exists, and nil otherwise.
	GetObject() ObjectPropertyInterface
	// GetOutbox returns the "outbox" property if it exists, and nil otherwise.
	GetOutbox() OutboxPropertyInterface
	// GetPreferredUsername returns the "preferredUsername" property if it
	// exists, and nil otherwise.
	GetPreferredUsername() PreferredUsernamePropertyInterface
	// GetPreview returns the "preview" property if it exists, and nil
	// otherwise.
	GetPreview() PreviewPropertyInterface
	// GetPublished returns the "published" property if it exists, and nil
	// otherwise.
	GetPublished() PublishedPropertyInterface
	// GetReplies returns the "replies" property if it exists, and nil
	// otherwise.
	GetReplies() RepliesPropertyInterface
	// GetShares returns the "shares" property if it exists, and nil otherwise.
	GetShares() SharesPropertyInterface
	// GetStartTime returns the "startTime" property if it exists, and nil
	// otherwise.
	GetStartTime() StartTimePropertyInterface
	// GetStreams returns the "streams" property if it exists, and nil
	// otherwise.
	GetStreams() StreamsPropertyInterface
	// GetSummary returns the "summary" property if it exists, and nil
	// otherwise.
	GetSummary() SummaryPropertyInterface
	// GetTag returns the "tag" property if it exists, and nil otherwise.
	GetTag() TagPropertyInterface
	// GetTo returns the "to" property if it exists, and nil otherwise.
	GetTo() ToPropertyInterface
	// GetType returns the "type" property if it exists, and nil otherwise.
	GetType() TypePropertyInterface
	// GetUnknownProperties returns the unknown properties for the Group type.
	// Note that this should not be used by app developers. It is only
	// used to help determine which implementation is LessThan the other.
	// Developers who are creating a different implementation of this
	// type's interface can use this method in their LessThan
	// implementation, but routine ActivityPub applications should not use
	// this to bypass the code generation tool.
	GetUnknownProperties() map[string]interface{}
	// GetUpdated returns the "updated" property if it exists, and nil
	// otherwise.
	GetUpdated() UpdatedPropertyInterface
	// GetUrl returns the "url" property if it exists, and nil otherwise.
	GetUrl() UrlPropertyInterface
	// IsExtending returns true if the Group type extends from the other type.
	IsExtending(other Type) bool
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this type and the specific properties that are set. The value
	// in the map is the alias used to import the type and its properties.
	JSONLDContext() map[string]string
	// LessThan computes if this Group is lesser, with an arbitrary but stable
	// determination.
	LessThan(o GroupInterface) bool
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format.
	Serialize() (map[string]interface{}, error)
	// SetAltitude returns the "altitude" property if it exists, and nil
	// otherwise.
	SetAltitude(i AltitudePropertyInterface)
	// SetAttachment returns the "attachment" property if it exists, and nil
	// otherwise.
	SetAttachment(i AttachmentPropertyInterface)
	// SetAttributedTo returns the "attributedTo" property if it exists, and
	// nil otherwise.
	SetAttributedTo(i AttributedToPropertyInterface)
	// SetAudience returns the "audience" property if it exists, and nil
	// otherwise.
	SetAudience(i AudiencePropertyInterface)
	// SetBcc returns the "bcc" property if it exists, and nil otherwise.
	SetBcc(i BccPropertyInterface)
	// SetBto returns the "bto" property if it exists, and nil otherwise.
	SetBto(i BtoPropertyInterface)
	// SetCc returns the "cc" property if it exists, and nil otherwise.
	SetCc(i CcPropertyInterface)
	// SetContent returns the "content" property if it exists, and nil
	// otherwise.
	SetContent(i ContentPropertyInterface)
	// SetContext returns the "context" property if it exists, and nil
	// otherwise.
	SetContext(i ContextPropertyInterface)
	// SetDuration returns the "duration" property if it exists, and nil
	// otherwise.
	SetDuration(i DurationPropertyInterface)
	// SetEndTime returns the "endTime" property if it exists, and nil
	// otherwise.
	SetEndTime(i EndTimePropertyInterface)
	// SetFollowers returns the "followers" property if it exists, and nil
	// otherwise.
	SetFollowers(i FollowersPropertyInterface)
	// SetFollowing returns the "following" property if it exists, and nil
	// otherwise.
	SetFollowing(i FollowingPropertyInterface)
	// SetGenerator returns the "generator" property if it exists, and nil
	// otherwise.
	SetGenerator(i GeneratorPropertyInterface)
	// SetIcon returns the "icon" property if it exists, and nil otherwise.
	SetIcon(i IconPropertyInterface)
	// SetId returns the "id" property if it exists, and nil otherwise.
	SetId(i IdPropertyInterface)
	// SetImage returns the "image" property if it exists, and nil otherwise.
	SetImage(i ImagePropertyInterface)
	// SetInReplyTo returns the "inReplyTo" property if it exists, and nil
	// otherwise.
	SetInReplyTo(i InReplyToPropertyInterface)
	// SetInbox returns the "inbox" property if it exists, and nil otherwise.
	SetInbox(i InboxPropertyInterface)
	// SetLiked returns the "liked" property if it exists, and nil otherwise.
	SetLiked(i LikedPropertyInterface)
	// SetLikes returns the "likes" property if it exists, and nil otherwise.
	SetLikes(i LikesPropertyInterface)
	// SetLocation returns the "location" property if it exists, and nil
	// otherwise.
	SetLocation(i LocationPropertyInterface)
	// SetMediaType returns the "mediaType" property if it exists, and nil
	// otherwise.
	SetMediaType(i MediaTypePropertyInterface)
	// SetName returns the "name" property if it exists, and nil otherwise.
	SetName(i NamePropertyInterface)
	// SetObject returns the "object" property if it exists, and nil otherwise.
	SetObject(i ObjectPropertyInterface)
	// SetOutbox returns the "outbox" property if it exists, and nil otherwise.
	SetOutbox(i OutboxPropertyInterface)
	// SetPreferredUsername returns the "preferredUsername" property if it
	// exists, and nil otherwise.
	SetPreferredUsername(i PreferredUsernamePropertyInterface)
	// SetPreview returns the "preview" property if it exists, and nil
	// otherwise.
	SetPreview(i PreviewPropertyInterface)
	// SetPublished returns the "published" property if it exists, and nil
	// otherwise.
	SetPublished(i PublishedPropertyInterface)
	// SetReplies returns the "replies" property if it exists, and nil
	// otherwise.
	SetReplies(i RepliesPropertyInterface)
	// SetShares returns the "shares" property if it exists, and nil otherwise.
	SetShares(i SharesPropertyInterface)
	// SetStartTime returns the "startTime" property if it exists, and nil
	// otherwise.
	SetStartTime(i StartTimePropertyInterface)
	// SetStreams returns the "streams" property if it exists, and nil
	// otherwise.
	SetStreams(i StreamsPropertyInterface)
	// SetSummary returns the "summary" property if it exists, and nil
	// otherwise.
	SetSummary(i SummaryPropertyInterface)
	// SetTag returns the "tag" property if it exists, and nil otherwise.
	SetTag(i TagPropertyInterface)
	// SetTo returns the "to" property if it exists, and nil otherwise.
	SetTo(i ToPropertyInterface)
	// SetType returns the "type" property if it exists, and nil otherwise.
	SetType(i TypePropertyInterface)
	// SetUpdated returns the "updated" property if it exists, and nil
	// otherwise.
	SetUpdated(i UpdatedPropertyInterface)
	// SetUrl returns the "url" property if it exists, and nil otherwise.
	SetUrl(i UrlPropertyInterface)
}

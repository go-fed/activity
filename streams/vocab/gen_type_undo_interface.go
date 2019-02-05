package vocab

// Indicates that the actor is undoing the object. In most cases, the object will
// be an Activity describing some previously performed action (for instance, a
// person may have previously "liked" an article but, for whatever reason,
// might choose to undo that like at some later point in time). The target and
// origin typically have no defined meaning.
//
// Example 29 (https://www.w3.org/TR/activitystreams-vocabulary/#ex32-jsonld):
//   {
//     "actor": "http://sally.example.org",
//     "object": {
//       "actor": "http://sally.example.org",
//       "object": "http://example.org/posts/1",
//       "target": "http://john.example.org",
//       "type": "Offer"
//     },
//     "summary": "Sally retracted her offer to John",
//     "type": "Undo"
//   }
type UndoInterface interface {
	// GetActor returns the "actor" property if it exists, and nil otherwise.
	GetActor() ActorPropertyInterface
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
	// GetInstrument returns the "instrument" property if it exists, and nil
	// otherwise.
	GetInstrument() InstrumentPropertyInterface
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
	// GetOrigin returns the "origin" property if it exists, and nil otherwise.
	GetOrigin() OriginPropertyInterface
	// GetPreview returns the "preview" property if it exists, and nil
	// otherwise.
	GetPreview() PreviewPropertyInterface
	// GetPublished returns the "published" property if it exists, and nil
	// otherwise.
	GetPublished() PublishedPropertyInterface
	// GetReplies returns the "replies" property if it exists, and nil
	// otherwise.
	GetReplies() RepliesPropertyInterface
	// GetResult returns the "result" property if it exists, and nil otherwise.
	GetResult() ResultPropertyInterface
	// GetShares returns the "shares" property if it exists, and nil otherwise.
	GetShares() SharesPropertyInterface
	// GetStartTime returns the "startTime" property if it exists, and nil
	// otherwise.
	GetStartTime() StartTimePropertyInterface
	// GetSummary returns the "summary" property if it exists, and nil
	// otherwise.
	GetSummary() SummaryPropertyInterface
	// GetTag returns the "tag" property if it exists, and nil otherwise.
	GetTag() TagPropertyInterface
	// GetTarget returns the "target" property if it exists, and nil otherwise.
	GetTarget() TargetPropertyInterface
	// GetTo returns the "to" property if it exists, and nil otherwise.
	GetTo() ToPropertyInterface
	// GetType returns the "type" property if it exists, and nil otherwise.
	GetType() TypePropertyInterface
	// GetUnknownProperties returns the unknown properties for the Undo type.
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
	// IsExtending returns true if the Undo type extends from the other type.
	IsExtending(other Type) bool
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this type and the specific properties that are set. The value
	// in the map is the alias used to import the type and its properties.
	JSONLDContext() map[string]string
	// LessThan computes if this Undo is lesser, with an arbitrary but stable
	// determination.
	LessThan(o UndoInterface) bool
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format.
	Serialize() (map[string]interface{}, error)
	// SetActor returns the "actor" property if it exists, and nil otherwise.
	SetActor(i ActorPropertyInterface)
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
	// SetInstrument returns the "instrument" property if it exists, and nil
	// otherwise.
	SetInstrument(i InstrumentPropertyInterface)
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
	// SetOrigin returns the "origin" property if it exists, and nil otherwise.
	SetOrigin(i OriginPropertyInterface)
	// SetPreview returns the "preview" property if it exists, and nil
	// otherwise.
	SetPreview(i PreviewPropertyInterface)
	// SetPublished returns the "published" property if it exists, and nil
	// otherwise.
	SetPublished(i PublishedPropertyInterface)
	// SetReplies returns the "replies" property if it exists, and nil
	// otherwise.
	SetReplies(i RepliesPropertyInterface)
	// SetResult returns the "result" property if it exists, and nil otherwise.
	SetResult(i ResultPropertyInterface)
	// SetShares returns the "shares" property if it exists, and nil otherwise.
	SetShares(i SharesPropertyInterface)
	// SetStartTime returns the "startTime" property if it exists, and nil
	// otherwise.
	SetStartTime(i StartTimePropertyInterface)
	// SetSummary returns the "summary" property if it exists, and nil
	// otherwise.
	SetSummary(i SummaryPropertyInterface)
	// SetTag returns the "tag" property if it exists, and nil otherwise.
	SetTag(i TagPropertyInterface)
	// SetTarget returns the "target" property if it exists, and nil otherwise.
	SetTarget(i TargetPropertyInterface)
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

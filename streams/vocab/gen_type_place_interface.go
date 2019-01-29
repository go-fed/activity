package vocab

// Represents a logical or physical location. See 5.3 Representing Places for
// additional information.
//
// Example 56 (https://www.w3.org/TR/activitystreams-vocabulary/#ex57-jsonld):
//   {
//     "name": "Work",
//     "type": "Place"
//   }
//
// Example 57 (https://www.w3.org/TR/activitystreams-vocabulary/#ex58-jsonld):
//   {
//     "latitude": 36.75,
//     "longitude": 119.7667,
//     "name": "Fresno Area",
//     "radius": 15,
//     "type": "Place",
//     "units": "miles"
//   }
type PlaceInterface interface {
	// GetAccuracy returns the "accuracy" property if it exists, and nil
	// otherwise.
	GetAccuracy() AccuracyPropertyInterface
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
	// GetLatitude returns the "latitude" property if it exists, and nil
	// otherwise.
	GetLatitude() LatitudePropertyInterface
	// GetLocation returns the "location" property if it exists, and nil
	// otherwise.
	GetLocation() LocationPropertyInterface
	// GetLongitude returns the "longitude" property if it exists, and nil
	// otherwise.
	GetLongitude() LongitudePropertyInterface
	// GetMediaType returns the "mediaType" property if it exists, and nil
	// otherwise.
	GetMediaType() MediaTypePropertyInterface
	// GetName returns the "name" property if it exists, and nil otherwise.
	GetName() NamePropertyInterface
	// GetObject returns the "object" property if it exists, and nil otherwise.
	GetObject() ObjectPropertyInterface
	// GetPreview returns the "preview" property if it exists, and nil
	// otherwise.
	GetPreview() PreviewPropertyInterface
	// GetPublished returns the "published" property if it exists, and nil
	// otherwise.
	GetPublished() PublishedPropertyInterface
	// GetRadius returns the "radius" property if it exists, and nil otherwise.
	GetRadius() RadiusPropertyInterface
	// GetReplies returns the "replies" property if it exists, and nil
	// otherwise.
	GetReplies() RepliesPropertyInterface
	// GetStartTime returns the "startTime" property if it exists, and nil
	// otherwise.
	GetStartTime() StartTimePropertyInterface
	// GetSummary returns the "summary" property if it exists, and nil
	// otherwise.
	GetSummary() SummaryPropertyInterface
	// GetTag returns the "tag" property if it exists, and nil otherwise.
	GetTag() TagPropertyInterface
	// GetTo returns the "to" property if it exists, and nil otherwise.
	GetTo() ToPropertyInterface
	// GetType returns the "type" property if it exists, and nil otherwise.
	GetType() TypePropertyInterface
	// GetUnits returns the "units" property if it exists, and nil otherwise.
	GetUnits() UnitsPropertyInterface
	// GetUnknownProperties returns the unknown properties for the Place type.
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
	// IsExtending returns true if the Place type extends from the other type.
	IsExtending(other Type) bool
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this type and the specific properties that are set. The value
	// in the map is the alias used to import the type and its properties.
	JSONLDContext() map[string]string
	// LessThan computes if this Place is lesser, with an arbitrary but stable
	// determination.
	LessThan(o PlaceInterface) bool
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format.
	Serialize() (interface{}, error)
	// SetAccuracy returns the "accuracy" property if it exists, and nil
	// otherwise.
	SetAccuracy(i AccuracyPropertyInterface)
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
	// SetLatitude returns the "latitude" property if it exists, and nil
	// otherwise.
	SetLatitude(i LatitudePropertyInterface)
	// SetLocation returns the "location" property if it exists, and nil
	// otherwise.
	SetLocation(i LocationPropertyInterface)
	// SetLongitude returns the "longitude" property if it exists, and nil
	// otherwise.
	SetLongitude(i LongitudePropertyInterface)
	// SetMediaType returns the "mediaType" property if it exists, and nil
	// otherwise.
	SetMediaType(i MediaTypePropertyInterface)
	// SetName returns the "name" property if it exists, and nil otherwise.
	SetName(i NamePropertyInterface)
	// SetObject returns the "object" property if it exists, and nil otherwise.
	SetObject(i ObjectPropertyInterface)
	// SetPreview returns the "preview" property if it exists, and nil
	// otherwise.
	SetPreview(i PreviewPropertyInterface)
	// SetPublished returns the "published" property if it exists, and nil
	// otherwise.
	SetPublished(i PublishedPropertyInterface)
	// SetRadius returns the "radius" property if it exists, and nil otherwise.
	SetRadius(i RadiusPropertyInterface)
	// SetReplies returns the "replies" property if it exists, and nil
	// otherwise.
	SetReplies(i RepliesPropertyInterface)
	// SetStartTime returns the "startTime" property if it exists, and nil
	// otherwise.
	SetStartTime(i StartTimePropertyInterface)
	// SetSummary returns the "summary" property if it exists, and nil
	// otherwise.
	SetSummary(i SummaryPropertyInterface)
	// SetTag returns the "tag" property if it exists, and nil otherwise.
	SetTag(i TagPropertyInterface)
	// SetTo returns the "to" property if it exists, and nil otherwise.
	SetTo(i ToPropertyInterface)
	// SetType returns the "type" property if it exists, and nil otherwise.
	SetType(i TypePropertyInterface)
	// SetUnits returns the "units" property if it exists, and nil otherwise.
	SetUnits(i UnitsPropertyInterface)
	// SetUpdated returns the "updated" property if it exists, and nil
	// otherwise.
	SetUpdated(i UpdatedPropertyInterface)
	// SetUrl returns the "url" property if it exists, and nil otherwise.
	SetUrl(i UrlPropertyInterface)
}

package vocab

// Used to represent ordered subsets of items from an OrderedCollection. Refer to
// the Activity Streams 2.0 Core for a complete description of the
// OrderedCollectionPage object.
//
// Example 8 (https://www.w3.org/TR/activitystreams-vocabulary/#ex6c-jsonld):
//   {
//     "id": "http://example.org/foo?page=1",
//     "orderedItems": [
//       {
//         "name": "A Simple Note",
//         "type": "Note"
//       },
//       {
//         "name": "Another Simple Note",
//         "type": "Note"
//       }
//     ],
//     "partOf": "http://example.org/foo",
//     "summary": "Page 1 of Sally's notes",
//     "type": "OrderedCollectionPage"
//   }
type OrderedCollectionPageInterface interface {
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
	// GetCurrent returns the "current" property if it exists, and nil
	// otherwise.
	GetCurrent() CurrentPropertyInterface
	// GetDuration returns the "duration" property if it exists, and nil
	// otherwise.
	GetDuration() DurationPropertyInterface
	// GetEndTime returns the "endTime" property if it exists, and nil
	// otherwise.
	GetEndTime() EndTimePropertyInterface
	// GetFirst returns the "first" property if it exists, and nil otherwise.
	GetFirst() FirstPropertyInterface
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
	// GetItems returns the "items" property if it exists, and nil otherwise.
	GetItems() ItemsPropertyInterface
	// GetLast returns the "last" property if it exists, and nil otherwise.
	GetLast() LastPropertyInterface
	// GetLocation returns the "location" property if it exists, and nil
	// otherwise.
	GetLocation() LocationPropertyInterface
	// GetMediaType returns the "mediaType" property if it exists, and nil
	// otherwise.
	GetMediaType() MediaTypePropertyInterface
	// GetName returns the "name" property if it exists, and nil otherwise.
	GetName() NamePropertyInterface
	// GetNext returns the "next" property if it exists, and nil otherwise.
	GetNext() NextPropertyInterface
	// GetObject returns the "object" property if it exists, and nil otherwise.
	GetObject() ObjectPropertyInterface
	// GetPartOf returns the "partOf" property if it exists, and nil otherwise.
	GetPartOf() PartOfPropertyInterface
	// GetPrev returns the "prev" property if it exists, and nil otherwise.
	GetPrev() PrevPropertyInterface
	// GetPreview returns the "preview" property if it exists, and nil
	// otherwise.
	GetPreview() PreviewPropertyInterface
	// GetPublished returns the "published" property if it exists, and nil
	// otherwise.
	GetPublished() PublishedPropertyInterface
	// GetReplies returns the "replies" property if it exists, and nil
	// otherwise.
	GetReplies() RepliesPropertyInterface
	// GetStartIndex returns the "startIndex" property if it exists, and nil
	// otherwise.
	GetStartIndex() StartIndexPropertyInterface
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
	// GetTotalItems returns the "totalItems" property if it exists, and nil
	// otherwise.
	GetTotalItems() TotalItemsPropertyInterface
	// GetType returns the "type" property if it exists, and nil otherwise.
	GetType() TypePropertyInterface
	// GetUnknownProperties returns the unknown properties for the
	// OrderedCollectionPage type. Note that this should not be used by
	// app developers. It is only used to help determine which
	// implementation is LessThan the other. Developers who are creating a
	// different implementation of this type's interface can use this
	// method in their LessThan implementation, but routine ActivityPub
	// applications should not use this to bypass the code generation tool.
	GetUnknownProperties() map[string]interface{}
	// GetUpdated returns the "updated" property if it exists, and nil
	// otherwise.
	GetUpdated() UpdatedPropertyInterface
	// GetUrl returns the "url" property if it exists, and nil otherwise.
	GetUrl() UrlPropertyInterface
	// IsExtending returns true if the OrderedCollectionPage type extends from
	// the other type.
	IsExtending(other Type) bool
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this type and the specific properties that are set. The value
	// in the map is the alias used to import the type and its properties.
	JSONLDContext() map[string]string
	// LessThan computes if this OrderedCollectionPage is lesser, with an
	// arbitrary but stable determination.
	LessThan(o OrderedCollectionPageInterface) bool
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format.
	Serialize() (interface{}, error)
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
	// SetCurrent returns the "current" property if it exists, and nil
	// otherwise.
	SetCurrent(i CurrentPropertyInterface)
	// SetDuration returns the "duration" property if it exists, and nil
	// otherwise.
	SetDuration(i DurationPropertyInterface)
	// SetEndTime returns the "endTime" property if it exists, and nil
	// otherwise.
	SetEndTime(i EndTimePropertyInterface)
	// SetFirst returns the "first" property if it exists, and nil otherwise.
	SetFirst(i FirstPropertyInterface)
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
	// SetItems returns the "items" property if it exists, and nil otherwise.
	SetItems(i ItemsPropertyInterface)
	// SetLast returns the "last" property if it exists, and nil otherwise.
	SetLast(i LastPropertyInterface)
	// SetLocation returns the "location" property if it exists, and nil
	// otherwise.
	SetLocation(i LocationPropertyInterface)
	// SetMediaType returns the "mediaType" property if it exists, and nil
	// otherwise.
	SetMediaType(i MediaTypePropertyInterface)
	// SetName returns the "name" property if it exists, and nil otherwise.
	SetName(i NamePropertyInterface)
	// SetNext returns the "next" property if it exists, and nil otherwise.
	SetNext(i NextPropertyInterface)
	// SetObject returns the "object" property if it exists, and nil otherwise.
	SetObject(i ObjectPropertyInterface)
	// SetPartOf returns the "partOf" property if it exists, and nil otherwise.
	SetPartOf(i PartOfPropertyInterface)
	// SetPrev returns the "prev" property if it exists, and nil otherwise.
	SetPrev(i PrevPropertyInterface)
	// SetPreview returns the "preview" property if it exists, and nil
	// otherwise.
	SetPreview(i PreviewPropertyInterface)
	// SetPublished returns the "published" property if it exists, and nil
	// otherwise.
	SetPublished(i PublishedPropertyInterface)
	// SetReplies returns the "replies" property if it exists, and nil
	// otherwise.
	SetReplies(i RepliesPropertyInterface)
	// SetStartIndex returns the "startIndex" property if it exists, and nil
	// otherwise.
	SetStartIndex(i StartIndexPropertyInterface)
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
	// SetTotalItems returns the "totalItems" property if it exists, and nil
	// otherwise.
	SetTotalItems(i TotalItemsPropertyInterface)
	// SetType returns the "type" property if it exists, and nil otherwise.
	SetType(i TypePropertyInterface)
	// SetUpdated returns the "updated" property if it exists, and nil
	// otherwise.
	SetUpdated(i UpdatedPropertyInterface)
	// SetUrl returns the "url" property if it exists, and nil otherwise.
	SetUrl(i UrlPropertyInterface)
}

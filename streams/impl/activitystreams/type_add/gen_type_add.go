package typeadd

import vocab "github.com/go-fed/activity/streams/vocab"

// Indicates that the actor has added the object to the target. If the target
// property is not explicitly specified, the target would need to be
// determined implicitly by context. The origin can be used to identify the
// context from which the object originated.
//
// Example 12 (https://www.w3.org/TR/activitystreams-vocabulary/#ex9-jsonld):
//   {
//     "actor": {
//       "name": "Sally",
//       "type": "Person"
//     },
//     "object": "http://example.org/abc",
//     "summary": "Sally added an object",
//     "type": "Add"
//   }
//
// Example 13 (https://www.w3.org/TR/activitystreams-vocabulary/#ex10-jsonld):
//   {
//     "actor": {
//       "name": "Sally",
//       "type": "Person"
//     },
//     "object": {
//       "name": "A picture of my cat",
//       "type": "Image",
//       "url": "http://example.org/img/cat.png"
//     },
//     "origin": {
//       "name": "Camera Roll",
//       "type": "Collection"
//     },
//     "summary": "Sally added a picture of her cat to her cat picture
// collection",
//     "target": {
//       "name": "My Cat Pictures",
//       "type": "Collection"
//     },
//     "type": "Add"
//   }
type Add struct {
	Actor        vocab.ActorPropertyInterface
	Altitude     vocab.AltitudePropertyInterface
	Attachment   vocab.AttachmentPropertyInterface
	AttributedTo vocab.AttributedToPropertyInterface
	Audience     vocab.AudiencePropertyInterface
	Bcc          vocab.BccPropertyInterface
	Bto          vocab.BtoPropertyInterface
	Cc           vocab.CcPropertyInterface
	Content      vocab.ContentPropertyInterface
	Context      vocab.ContextPropertyInterface
	Duration     vocab.DurationPropertyInterface
	EndTime      vocab.EndTimePropertyInterface
	Generator    vocab.GeneratorPropertyInterface
	Icon         vocab.IconPropertyInterface
	Id           vocab.IdPropertyInterface
	Image        vocab.ImagePropertyInterface
	InReplyTo    vocab.InReplyToPropertyInterface
	Instrument   vocab.InstrumentPropertyInterface
	Location     vocab.LocationPropertyInterface
	MediaType    vocab.MediaTypePropertyInterface
	Name         vocab.NamePropertyInterface
	Object       vocab.ObjectPropertyInterface
	Origin       vocab.OriginPropertyInterface
	Preview      vocab.PreviewPropertyInterface
	Published    vocab.PublishedPropertyInterface
	Replies      vocab.RepliesPropertyInterface
	Result       vocab.ResultPropertyInterface
	StartTime    vocab.StartTimePropertyInterface
	Summary      vocab.SummaryPropertyInterface
	Tag          vocab.TagPropertyInterface
	Target       vocab.TargetPropertyInterface
	To           vocab.ToPropertyInterface
	Type         vocab.TypePropertyInterface
	Updated      vocab.UpdatedPropertyInterface
	Url          vocab.UrlPropertyInterface
	alias        string
	unknown      map[string]interface{}
}

// AddExtends returns true if the Add type extends from the other type.
func AddExtends(other vocab.Type) bool {
	extensions := []string{"Activity", "Object"}
	for _, ext := range extensions {
		if ext == other.GetName() {
			return true
		}
	}
	return false
}

// AddIsDisjointWith returns true if the other provided type is disjoint with the
// Add type.
func AddIsDisjointWith(other vocab.Type) bool {
	disjointWith := []string{"Link", "Mention"}
	for _, disjoint := range disjointWith {
		if disjoint == other.GetName() {
			return true
		}
	}
	return false
}

// AddIsExtendedBy returns true if the other provided type extends from the Add
// type.
func AddIsExtendedBy(other vocab.Type) bool {
	// Shortcut implementation: is not extended by anything.
	return false
}

// DeserializeAdd creates a Add from a map representation that has been
// unmarshalled from a text or binary format.
func DeserializeAdd(m map[string]interface{}, aliasMap map[string]string) (*Add, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	this := &Add{
		alias:   alias,
		unknown: make(map[string]interface{}),
	}
	// Begin: Known property deserialization
	if p, err := mgr.DeserializeActorPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Actor = p
	}
	if p, err := mgr.DeserializeAltitudePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Altitude = p
	}
	if p, err := mgr.DeserializeAttachmentPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Attachment = p
	}
	if p, err := mgr.DeserializeAttributedToPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.AttributedTo = p
	}
	if p, err := mgr.DeserializeAudiencePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Audience = p
	}
	if p, err := mgr.DeserializeBccPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Bcc = p
	}
	if p, err := mgr.DeserializeBtoPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Bto = p
	}
	if p, err := mgr.DeserializeCcPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Cc = p
	}
	if p, err := mgr.DeserializeContentPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Content = p
	}
	if p, err := mgr.DeserializeContextPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Context = p
	}
	if p, err := mgr.DeserializeDurationPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Duration = p
	}
	if p, err := mgr.DeserializeEndTimePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.EndTime = p
	}
	if p, err := mgr.DeserializeGeneratorPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Generator = p
	}
	if p, err := mgr.DeserializeIconPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Icon = p
	}
	if p, err := mgr.DeserializeIdPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Id = p
	}
	if p, err := mgr.DeserializeImagePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Image = p
	}
	if p, err := mgr.DeserializeInReplyToPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.InReplyTo = p
	}
	if p, err := mgr.DeserializeInstrumentPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Instrument = p
	}
	if p, err := mgr.DeserializeLocationPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Location = p
	}
	if p, err := mgr.DeserializeMediaTypePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.MediaType = p
	}
	if p, err := mgr.DeserializeNamePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Name = p
	}
	if p, err := mgr.DeserializeObjectPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Object = p
	}
	if p, err := mgr.DeserializeOriginPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Origin = p
	}
	if p, err := mgr.DeserializePreviewPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Preview = p
	}
	if p, err := mgr.DeserializePublishedPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Published = p
	}
	if p, err := mgr.DeserializeRepliesPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Replies = p
	}
	if p, err := mgr.DeserializeResultPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Result = p
	}
	if p, err := mgr.DeserializeStartTimePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.StartTime = p
	}
	if p, err := mgr.DeserializeSummaryPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Summary = p
	}
	if p, err := mgr.DeserializeTagPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Tag = p
	}
	if p, err := mgr.DeserializeTargetPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Target = p
	}
	if p, err := mgr.DeserializeToPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.To = p
	}
	if p, err := mgr.DeserializeTypePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Type = p
	}
	if p, err := mgr.DeserializeUpdatedPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Updated = p
	}
	if p, err := mgr.DeserializeUrlPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Url = p
	}
	// End: Known property deserialization

	// Begin: Unknown deserialization
	for k, v := range m {
		// Begin: Code that ensures a property name is unknown
		if k == "actor" {
			continue
		} else if k == "altitude" {
			continue
		} else if k == "attachment" {
			continue
		} else if k == "attributedTo" {
			continue
		} else if k == "audience" {
			continue
		} else if k == "bcc" {
			continue
		} else if k == "bto" {
			continue
		} else if k == "cc" {
			continue
		} else if k == "content" {
			continue
		} else if k == "context" {
			continue
		} else if k == "duration" {
			continue
		} else if k == "endTime" {
			continue
		} else if k == "generator" {
			continue
		} else if k == "icon" {
			continue
		} else if k == "id" {
			continue
		} else if k == "image" {
			continue
		} else if k == "inReplyTo" {
			continue
		} else if k == "instrument" {
			continue
		} else if k == "location" {
			continue
		} else if k == "mediaType" {
			continue
		} else if k == "name" {
			continue
		} else if k == "object" {
			continue
		} else if k == "origin" {
			continue
		} else if k == "preview" {
			continue
		} else if k == "published" {
			continue
		} else if k == "replies" {
			continue
		} else if k == "result" {
			continue
		} else if k == "startTime" {
			continue
		} else if k == "summary" {
			continue
		} else if k == "tag" {
			continue
		} else if k == "target" {
			continue
		} else if k == "to" {
			continue
		} else if k == "type" {
			continue
		} else if k == "updated" {
			continue
		} else if k == "url" {
			continue
		} // End: Code that ensures a property name is unknown

		this.unknown[k] = v
	}
	// End: Unknown deserialization

	return this, nil
}

// NewAdd creates a new Add type
func NewAdd() *Add {
	return &Add{
		alias:   "",
		unknown: make(map[string]interface{}, 0),
	}
}

// GetActor returns the "actor" property if it exists, and nil otherwise.
func (this Add) GetActor() vocab.ActorPropertyInterface {
	return this.Actor
}

// GetAltitude returns the "altitude" property if it exists, and nil otherwise.
func (this Add) GetAltitude() vocab.AltitudePropertyInterface {
	return this.Altitude
}

// GetAttachment returns the "attachment" property if it exists, and nil otherwise.
func (this Add) GetAttachment() vocab.AttachmentPropertyInterface {
	return this.Attachment
}

// GetAttributedTo returns the "attributedTo" property if it exists, and nil
// otherwise.
func (this Add) GetAttributedTo() vocab.AttributedToPropertyInterface {
	return this.AttributedTo
}

// GetAudience returns the "audience" property if it exists, and nil otherwise.
func (this Add) GetAudience() vocab.AudiencePropertyInterface {
	return this.Audience
}

// GetBcc returns the "bcc" property if it exists, and nil otherwise.
func (this Add) GetBcc() vocab.BccPropertyInterface {
	return this.Bcc
}

// GetBto returns the "bto" property if it exists, and nil otherwise.
func (this Add) GetBto() vocab.BtoPropertyInterface {
	return this.Bto
}

// GetCc returns the "cc" property if it exists, and nil otherwise.
func (this Add) GetCc() vocab.CcPropertyInterface {
	return this.Cc
}

// GetContent returns the "content" property if it exists, and nil otherwise.
func (this Add) GetContent() vocab.ContentPropertyInterface {
	return this.Content
}

// GetContext returns the "context" property if it exists, and nil otherwise.
func (this Add) GetContext() vocab.ContextPropertyInterface {
	return this.Context
}

// GetDuration returns the "duration" property if it exists, and nil otherwise.
func (this Add) GetDuration() vocab.DurationPropertyInterface {
	return this.Duration
}

// GetEndTime returns the "endTime" property if it exists, and nil otherwise.
func (this Add) GetEndTime() vocab.EndTimePropertyInterface {
	return this.EndTime
}

// GetGenerator returns the "generator" property if it exists, and nil otherwise.
func (this Add) GetGenerator() vocab.GeneratorPropertyInterface {
	return this.Generator
}

// GetIcon returns the "icon" property if it exists, and nil otherwise.
func (this Add) GetIcon() vocab.IconPropertyInterface {
	return this.Icon
}

// GetId returns the "id" property if it exists, and nil otherwise.
func (this Add) GetId() vocab.IdPropertyInterface {
	return this.Id
}

// GetImage returns the "image" property if it exists, and nil otherwise.
func (this Add) GetImage() vocab.ImagePropertyInterface {
	return this.Image
}

// GetInReplyTo returns the "inReplyTo" property if it exists, and nil otherwise.
func (this Add) GetInReplyTo() vocab.InReplyToPropertyInterface {
	return this.InReplyTo
}

// GetInstrument returns the "instrument" property if it exists, and nil otherwise.
func (this Add) GetInstrument() vocab.InstrumentPropertyInterface {
	return this.Instrument
}

// GetLocation returns the "location" property if it exists, and nil otherwise.
func (this Add) GetLocation() vocab.LocationPropertyInterface {
	return this.Location
}

// GetMediaType returns the "mediaType" property if it exists, and nil otherwise.
func (this Add) GetMediaType() vocab.MediaTypePropertyInterface {
	return this.MediaType
}

// GetName returns the "name" property if it exists, and nil otherwise.
func (this Add) GetName() vocab.NamePropertyInterface {
	return this.Name
}

// GetObject returns the "object" property if it exists, and nil otherwise.
func (this Add) GetObject() vocab.ObjectPropertyInterface {
	return this.Object
}

// GetOrigin returns the "origin" property if it exists, and nil otherwise.
func (this Add) GetOrigin() vocab.OriginPropertyInterface {
	return this.Origin
}

// GetPreview returns the "preview" property if it exists, and nil otherwise.
func (this Add) GetPreview() vocab.PreviewPropertyInterface {
	return this.Preview
}

// GetPublished returns the "published" property if it exists, and nil otherwise.
func (this Add) GetPublished() vocab.PublishedPropertyInterface {
	return this.Published
}

// GetReplies returns the "replies" property if it exists, and nil otherwise.
func (this Add) GetReplies() vocab.RepliesPropertyInterface {
	return this.Replies
}

// GetResult returns the "result" property if it exists, and nil otherwise.
func (this Add) GetResult() vocab.ResultPropertyInterface {
	return this.Result
}

// GetStartTime returns the "startTime" property if it exists, and nil otherwise.
func (this Add) GetStartTime() vocab.StartTimePropertyInterface {
	return this.StartTime
}

// GetSummary returns the "summary" property if it exists, and nil otherwise.
func (this Add) GetSummary() vocab.SummaryPropertyInterface {
	return this.Summary
}

// GetTag returns the "tag" property if it exists, and nil otherwise.
func (this Add) GetTag() vocab.TagPropertyInterface {
	return this.Tag
}

// GetTarget returns the "target" property if it exists, and nil otherwise.
func (this Add) GetTarget() vocab.TargetPropertyInterface {
	return this.Target
}

// GetTo returns the "to" property if it exists, and nil otherwise.
func (this Add) GetTo() vocab.ToPropertyInterface {
	return this.To
}

// GetType returns the "type" property if it exists, and nil otherwise.
func (this Add) GetType() vocab.TypePropertyInterface {
	return this.Type
}

// GetUnknownProperties returns the unknown properties for the Add type. Note that
// this should not be used by app developers. It is only used to help
// determine which implementation is LessThan the other. Developers who are
// creating a different implementation of this type's interface can use this
// method in their LessThan implementation, but routine ActivityPub
// applications should not use this to bypass the code generation tool.
func (this Add) GetUnknownProperties() map[string]interface{} {
	return this.unknown
}

// GetUpdated returns the "updated" property if it exists, and nil otherwise.
func (this Add) GetUpdated() vocab.UpdatedPropertyInterface {
	return this.Updated
}

// GetUrl returns the "url" property if it exists, and nil otherwise.
func (this Add) GetUrl() vocab.UrlPropertyInterface {
	return this.Url
}

// IsExtending returns true if the Add type extends from the other type.
func (this Add) IsExtending(other vocab.Type) bool {
	return AddExtends(other)
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// type and the specific properties that are set. The value in the map is the
// alias used to import the type and its properties.
func (this Add) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/TR/activitystreams-vocabulary": this.alias}
	m = this.helperJSONLDContext(this.Actor, m)
	m = this.helperJSONLDContext(this.Altitude, m)
	m = this.helperJSONLDContext(this.Attachment, m)
	m = this.helperJSONLDContext(this.AttributedTo, m)
	m = this.helperJSONLDContext(this.Audience, m)
	m = this.helperJSONLDContext(this.Bcc, m)
	m = this.helperJSONLDContext(this.Bto, m)
	m = this.helperJSONLDContext(this.Cc, m)
	m = this.helperJSONLDContext(this.Content, m)
	m = this.helperJSONLDContext(this.Context, m)
	m = this.helperJSONLDContext(this.Duration, m)
	m = this.helperJSONLDContext(this.EndTime, m)
	m = this.helperJSONLDContext(this.Generator, m)
	m = this.helperJSONLDContext(this.Icon, m)
	m = this.helperJSONLDContext(this.Id, m)
	m = this.helperJSONLDContext(this.Image, m)
	m = this.helperJSONLDContext(this.InReplyTo, m)
	m = this.helperJSONLDContext(this.Instrument, m)
	m = this.helperJSONLDContext(this.Location, m)
	m = this.helperJSONLDContext(this.MediaType, m)
	m = this.helperJSONLDContext(this.Name, m)
	m = this.helperJSONLDContext(this.Object, m)
	m = this.helperJSONLDContext(this.Origin, m)
	m = this.helperJSONLDContext(this.Preview, m)
	m = this.helperJSONLDContext(this.Published, m)
	m = this.helperJSONLDContext(this.Replies, m)
	m = this.helperJSONLDContext(this.Result, m)
	m = this.helperJSONLDContext(this.StartTime, m)
	m = this.helperJSONLDContext(this.Summary, m)
	m = this.helperJSONLDContext(this.Tag, m)
	m = this.helperJSONLDContext(this.Target, m)
	m = this.helperJSONLDContext(this.To, m)
	m = this.helperJSONLDContext(this.Type, m)
	m = this.helperJSONLDContext(this.Updated, m)
	m = this.helperJSONLDContext(this.Url, m)

	return m
}

// LessThan computes if this Add is lesser, with an arbitrary but stable
// determination.
func (this Add) LessThan(o vocab.AddInterface) bool {
	// Begin: Compare known properties
	// Compare property "actor"
	if this.Actor.LessThan(o.GetActor()) {
		return true
	} else if o.GetActor().LessThan(this.Actor) {
		return false
	}
	// Compare property "altitude"
	if this.Altitude.LessThan(o.GetAltitude()) {
		return true
	} else if o.GetAltitude().LessThan(this.Altitude) {
		return false
	}
	// Compare property "attachment"
	if this.Attachment.LessThan(o.GetAttachment()) {
		return true
	} else if o.GetAttachment().LessThan(this.Attachment) {
		return false
	}
	// Compare property "attributedTo"
	if this.AttributedTo.LessThan(o.GetAttributedTo()) {
		return true
	} else if o.GetAttributedTo().LessThan(this.AttributedTo) {
		return false
	}
	// Compare property "audience"
	if this.Audience.LessThan(o.GetAudience()) {
		return true
	} else if o.GetAudience().LessThan(this.Audience) {
		return false
	}
	// Compare property "bcc"
	if this.Bcc.LessThan(o.GetBcc()) {
		return true
	} else if o.GetBcc().LessThan(this.Bcc) {
		return false
	}
	// Compare property "bto"
	if this.Bto.LessThan(o.GetBto()) {
		return true
	} else if o.GetBto().LessThan(this.Bto) {
		return false
	}
	// Compare property "cc"
	if this.Cc.LessThan(o.GetCc()) {
		return true
	} else if o.GetCc().LessThan(this.Cc) {
		return false
	}
	// Compare property "content"
	if this.Content.LessThan(o.GetContent()) {
		return true
	} else if o.GetContent().LessThan(this.Content) {
		return false
	}
	// Compare property "context"
	if this.Context.LessThan(o.GetContext()) {
		return true
	} else if o.GetContext().LessThan(this.Context) {
		return false
	}
	// Compare property "duration"
	if this.Duration.LessThan(o.GetDuration()) {
		return true
	} else if o.GetDuration().LessThan(this.Duration) {
		return false
	}
	// Compare property "endTime"
	if this.EndTime.LessThan(o.GetEndTime()) {
		return true
	} else if o.GetEndTime().LessThan(this.EndTime) {
		return false
	}
	// Compare property "generator"
	if this.Generator.LessThan(o.GetGenerator()) {
		return true
	} else if o.GetGenerator().LessThan(this.Generator) {
		return false
	}
	// Compare property "icon"
	if this.Icon.LessThan(o.GetIcon()) {
		return true
	} else if o.GetIcon().LessThan(this.Icon) {
		return false
	}
	// Compare property "id"
	if this.Id.LessThan(o.GetId()) {
		return true
	} else if o.GetId().LessThan(this.Id) {
		return false
	}
	// Compare property "image"
	if this.Image.LessThan(o.GetImage()) {
		return true
	} else if o.GetImage().LessThan(this.Image) {
		return false
	}
	// Compare property "inReplyTo"
	if this.InReplyTo.LessThan(o.GetInReplyTo()) {
		return true
	} else if o.GetInReplyTo().LessThan(this.InReplyTo) {
		return false
	}
	// Compare property "instrument"
	if this.Instrument.LessThan(o.GetInstrument()) {
		return true
	} else if o.GetInstrument().LessThan(this.Instrument) {
		return false
	}
	// Compare property "location"
	if this.Location.LessThan(o.GetLocation()) {
		return true
	} else if o.GetLocation().LessThan(this.Location) {
		return false
	}
	// Compare property "mediaType"
	if this.MediaType.LessThan(o.GetMediaType()) {
		return true
	} else if o.GetMediaType().LessThan(this.MediaType) {
		return false
	}
	// Compare property "name"
	if this.Name.LessThan(o.GetName()) {
		return true
	} else if o.GetName().LessThan(this.Name) {
		return false
	}
	// Compare property "object"
	if this.Object.LessThan(o.GetObject()) {
		return true
	} else if o.GetObject().LessThan(this.Object) {
		return false
	}
	// Compare property "origin"
	if this.Origin.LessThan(o.GetOrigin()) {
		return true
	} else if o.GetOrigin().LessThan(this.Origin) {
		return false
	}
	// Compare property "preview"
	if this.Preview.LessThan(o.GetPreview()) {
		return true
	} else if o.GetPreview().LessThan(this.Preview) {
		return false
	}
	// Compare property "published"
	if this.Published.LessThan(o.GetPublished()) {
		return true
	} else if o.GetPublished().LessThan(this.Published) {
		return false
	}
	// Compare property "replies"
	if this.Replies.LessThan(o.GetReplies()) {
		return true
	} else if o.GetReplies().LessThan(this.Replies) {
		return false
	}
	// Compare property "result"
	if this.Result.LessThan(o.GetResult()) {
		return true
	} else if o.GetResult().LessThan(this.Result) {
		return false
	}
	// Compare property "startTime"
	if this.StartTime.LessThan(o.GetStartTime()) {
		return true
	} else if o.GetStartTime().LessThan(this.StartTime) {
		return false
	}
	// Compare property "summary"
	if this.Summary.LessThan(o.GetSummary()) {
		return true
	} else if o.GetSummary().LessThan(this.Summary) {
		return false
	}
	// Compare property "tag"
	if this.Tag.LessThan(o.GetTag()) {
		return true
	} else if o.GetTag().LessThan(this.Tag) {
		return false
	}
	// Compare property "target"
	if this.Target.LessThan(o.GetTarget()) {
		return true
	} else if o.GetTarget().LessThan(this.Target) {
		return false
	}
	// Compare property "to"
	if this.To.LessThan(o.GetTo()) {
		return true
	} else if o.GetTo().LessThan(this.To) {
		return false
	}
	// Compare property "type"
	if this.Type.LessThan(o.GetType()) {
		return true
	} else if o.GetType().LessThan(this.Type) {
		return false
	}
	// Compare property "updated"
	if this.Updated.LessThan(o.GetUpdated()) {
		return true
	} else if o.GetUpdated().LessThan(this.Updated) {
		return false
	}
	// Compare property "url"
	if this.Url.LessThan(o.GetUrl()) {
		return true
	} else if o.GetUrl().LessThan(this.Url) {
		return false
	}
	// End: Compare known properties

	// Begin: Compare unknown properties (only by number of them)
	if len(this.unknown) < len(o.GetUnknownProperties()) {
		return true
	} else if len(o.GetUnknownProperties()) < len(this.unknown) {
		return false
	} // End: Compare unknown properties (only by number of them)

	// All properties are the same.
	return false
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format.
func (this Add) Serialize() (interface{}, error) {
	m := make(map[string]interface{})
	// Begin: Serialize known properties
	// Maybe serialize property "actor"
	if i, err := this.Actor.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Actor.Name()] = i
	}
	// Maybe serialize property "altitude"
	if i, err := this.Altitude.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Altitude.Name()] = i
	}
	// Maybe serialize property "attachment"
	if i, err := this.Attachment.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Attachment.Name()] = i
	}
	// Maybe serialize property "attributedTo"
	if i, err := this.AttributedTo.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.AttributedTo.Name()] = i
	}
	// Maybe serialize property "audience"
	if i, err := this.Audience.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Audience.Name()] = i
	}
	// Maybe serialize property "bcc"
	if i, err := this.Bcc.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Bcc.Name()] = i
	}
	// Maybe serialize property "bto"
	if i, err := this.Bto.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Bto.Name()] = i
	}
	// Maybe serialize property "cc"
	if i, err := this.Cc.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Cc.Name()] = i
	}
	// Maybe serialize property "content"
	if i, err := this.Content.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Content.Name()] = i
	}
	// Maybe serialize property "context"
	if i, err := this.Context.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Context.Name()] = i
	}
	// Maybe serialize property "duration"
	if i, err := this.Duration.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Duration.Name()] = i
	}
	// Maybe serialize property "endTime"
	if i, err := this.EndTime.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.EndTime.Name()] = i
	}
	// Maybe serialize property "generator"
	if i, err := this.Generator.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Generator.Name()] = i
	}
	// Maybe serialize property "icon"
	if i, err := this.Icon.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Icon.Name()] = i
	}
	// Maybe serialize property "id"
	if i, err := this.Id.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Id.Name()] = i
	}
	// Maybe serialize property "image"
	if i, err := this.Image.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Image.Name()] = i
	}
	// Maybe serialize property "inReplyTo"
	if i, err := this.InReplyTo.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.InReplyTo.Name()] = i
	}
	// Maybe serialize property "instrument"
	if i, err := this.Instrument.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Instrument.Name()] = i
	}
	// Maybe serialize property "location"
	if i, err := this.Location.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Location.Name()] = i
	}
	// Maybe serialize property "mediaType"
	if i, err := this.MediaType.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.MediaType.Name()] = i
	}
	// Maybe serialize property "name"
	if i, err := this.Name.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Name.Name()] = i
	}
	// Maybe serialize property "object"
	if i, err := this.Object.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Object.Name()] = i
	}
	// Maybe serialize property "origin"
	if i, err := this.Origin.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Origin.Name()] = i
	}
	// Maybe serialize property "preview"
	if i, err := this.Preview.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Preview.Name()] = i
	}
	// Maybe serialize property "published"
	if i, err := this.Published.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Published.Name()] = i
	}
	// Maybe serialize property "replies"
	if i, err := this.Replies.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Replies.Name()] = i
	}
	// Maybe serialize property "result"
	if i, err := this.Result.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Result.Name()] = i
	}
	// Maybe serialize property "startTime"
	if i, err := this.StartTime.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.StartTime.Name()] = i
	}
	// Maybe serialize property "summary"
	if i, err := this.Summary.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Summary.Name()] = i
	}
	// Maybe serialize property "tag"
	if i, err := this.Tag.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Tag.Name()] = i
	}
	// Maybe serialize property "target"
	if i, err := this.Target.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Target.Name()] = i
	}
	// Maybe serialize property "to"
	if i, err := this.To.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.To.Name()] = i
	}
	// Maybe serialize property "type"
	if i, err := this.Type.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Type.Name()] = i
	}
	// Maybe serialize property "updated"
	if i, err := this.Updated.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Updated.Name()] = i
	}
	// Maybe serialize property "url"
	if i, err := this.Url.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Url.Name()] = i
	}
	// End: Serialize known properties

	// Begin: Serialize unknown properties
	for k, v := range this.unknown {
		// To be safe, ensure we aren't overwriting a known property
		if _, has := m[k]; !has {
			m[k] = v
		}
	}
	// End: Serialize unknown properties

	return m, nil
}

// SetActor returns the "actor" property if it exists, and nil otherwise.
func (this Add) SetActor(i vocab.ActorPropertyInterface) {
	this.Actor = i
}

// SetAltitude returns the "altitude" property if it exists, and nil otherwise.
func (this Add) SetAltitude(i vocab.AltitudePropertyInterface) {
	this.Altitude = i
}

// SetAttachment returns the "attachment" property if it exists, and nil otherwise.
func (this Add) SetAttachment(i vocab.AttachmentPropertyInterface) {
	this.Attachment = i
}

// SetAttributedTo returns the "attributedTo" property if it exists, and nil
// otherwise.
func (this Add) SetAttributedTo(i vocab.AttributedToPropertyInterface) {
	this.AttributedTo = i
}

// SetAudience returns the "audience" property if it exists, and nil otherwise.
func (this Add) SetAudience(i vocab.AudiencePropertyInterface) {
	this.Audience = i
}

// SetBcc returns the "bcc" property if it exists, and nil otherwise.
func (this Add) SetBcc(i vocab.BccPropertyInterface) {
	this.Bcc = i
}

// SetBto returns the "bto" property if it exists, and nil otherwise.
func (this Add) SetBto(i vocab.BtoPropertyInterface) {
	this.Bto = i
}

// SetCc returns the "cc" property if it exists, and nil otherwise.
func (this Add) SetCc(i vocab.CcPropertyInterface) {
	this.Cc = i
}

// SetContent returns the "content" property if it exists, and nil otherwise.
func (this Add) SetContent(i vocab.ContentPropertyInterface) {
	this.Content = i
}

// SetContext returns the "context" property if it exists, and nil otherwise.
func (this Add) SetContext(i vocab.ContextPropertyInterface) {
	this.Context = i
}

// SetDuration returns the "duration" property if it exists, and nil otherwise.
func (this Add) SetDuration(i vocab.DurationPropertyInterface) {
	this.Duration = i
}

// SetEndTime returns the "endTime" property if it exists, and nil otherwise.
func (this Add) SetEndTime(i vocab.EndTimePropertyInterface) {
	this.EndTime = i
}

// SetGenerator returns the "generator" property if it exists, and nil otherwise.
func (this Add) SetGenerator(i vocab.GeneratorPropertyInterface) {
	this.Generator = i
}

// SetIcon returns the "icon" property if it exists, and nil otherwise.
func (this Add) SetIcon(i vocab.IconPropertyInterface) {
	this.Icon = i
}

// SetId returns the "id" property if it exists, and nil otherwise.
func (this Add) SetId(i vocab.IdPropertyInterface) {
	this.Id = i
}

// SetImage returns the "image" property if it exists, and nil otherwise.
func (this Add) SetImage(i vocab.ImagePropertyInterface) {
	this.Image = i
}

// SetInReplyTo returns the "inReplyTo" property if it exists, and nil otherwise.
func (this Add) SetInReplyTo(i vocab.InReplyToPropertyInterface) {
	this.InReplyTo = i
}

// SetInstrument returns the "instrument" property if it exists, and nil otherwise.
func (this Add) SetInstrument(i vocab.InstrumentPropertyInterface) {
	this.Instrument = i
}

// SetLocation returns the "location" property if it exists, and nil otherwise.
func (this Add) SetLocation(i vocab.LocationPropertyInterface) {
	this.Location = i
}

// SetMediaType returns the "mediaType" property if it exists, and nil otherwise.
func (this Add) SetMediaType(i vocab.MediaTypePropertyInterface) {
	this.MediaType = i
}

// SetName returns the "name" property if it exists, and nil otherwise.
func (this Add) SetName(i vocab.NamePropertyInterface) {
	this.Name = i
}

// SetObject returns the "object" property if it exists, and nil otherwise.
func (this Add) SetObject(i vocab.ObjectPropertyInterface) {
	this.Object = i
}

// SetOrigin returns the "origin" property if it exists, and nil otherwise.
func (this Add) SetOrigin(i vocab.OriginPropertyInterface) {
	this.Origin = i
}

// SetPreview returns the "preview" property if it exists, and nil otherwise.
func (this Add) SetPreview(i vocab.PreviewPropertyInterface) {
	this.Preview = i
}

// SetPublished returns the "published" property if it exists, and nil otherwise.
func (this Add) SetPublished(i vocab.PublishedPropertyInterface) {
	this.Published = i
}

// SetReplies returns the "replies" property if it exists, and nil otherwise.
func (this Add) SetReplies(i vocab.RepliesPropertyInterface) {
	this.Replies = i
}

// SetResult returns the "result" property if it exists, and nil otherwise.
func (this Add) SetResult(i vocab.ResultPropertyInterface) {
	this.Result = i
}

// SetStartTime returns the "startTime" property if it exists, and nil otherwise.
func (this Add) SetStartTime(i vocab.StartTimePropertyInterface) {
	this.StartTime = i
}

// SetSummary returns the "summary" property if it exists, and nil otherwise.
func (this Add) SetSummary(i vocab.SummaryPropertyInterface) {
	this.Summary = i
}

// SetTag returns the "tag" property if it exists, and nil otherwise.
func (this Add) SetTag(i vocab.TagPropertyInterface) {
	this.Tag = i
}

// SetTarget returns the "target" property if it exists, and nil otherwise.
func (this Add) SetTarget(i vocab.TargetPropertyInterface) {
	this.Target = i
}

// SetTo returns the "to" property if it exists, and nil otherwise.
func (this Add) SetTo(i vocab.ToPropertyInterface) {
	this.To = i
}

// SetType returns the "type" property if it exists, and nil otherwise.
func (this Add) SetType(i vocab.TypePropertyInterface) {
	this.Type = i
}

// SetUpdated returns the "updated" property if it exists, and nil otherwise.
func (this Add) SetUpdated(i vocab.UpdatedPropertyInterface) {
	this.Updated = i
}

// SetUrl returns the "url" property if it exists, and nil otherwise.
func (this Add) SetUrl(i vocab.UrlPropertyInterface) {
	this.Url = i
}

// helperJSONLDContext obtains the context uris and their aliases from a property,
// if it is not nil.
func (this Add) helperJSONLDContext(i jsonldContexter, toMerge map[string]string) map[string]string {
	if i == nil {
		return toMerge
	}
	for k, v := range i.JSONLDContext() {
		/*
		   Since the literal maps in this function are determined at
		   code-generation time, this loop should not overwrite an existing key with a
		   new value.
		*/
		toMerge[k] = v
	}
	return toMerge
}

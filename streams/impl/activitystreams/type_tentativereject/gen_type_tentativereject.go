package typetentativereject

import (
	"fmt"
	propertytype "github.com/go-fed/activity/streams/impl/activitystreams/property_type"
	vocab "github.com/go-fed/activity/streams/vocab"
	"strings"
)

// A specialization of Reject in which the rejection is considered tentative.
//
// Example 26 (https://www.w3.org/TR/activitystreams-vocabulary/#ex27-jsonld):
//   {
//     "actor": {
//       "name": "Sally",
//       "type": "Person"
//     },
//     "object": {
//       "actor": "http://john.example.org",
//       "object": {
//         "name": "Going-Away Party for Jim",
//         "type": "Event"
//       },
//       "type": "Invite"
//     },
//     "summary": "Sally tentatively rejected an invitation to a party",
//     "type": "TentativeReject"
//   }
type TentativeReject struct {
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
	Likes        vocab.LikesPropertyInterface
	Location     vocab.LocationPropertyInterface
	MediaType    vocab.MediaTypePropertyInterface
	Name         vocab.NamePropertyInterface
	Object       vocab.ObjectPropertyInterface
	Origin       vocab.OriginPropertyInterface
	Preview      vocab.PreviewPropertyInterface
	Published    vocab.PublishedPropertyInterface
	Replies      vocab.RepliesPropertyInterface
	Result       vocab.ResultPropertyInterface
	Shares       vocab.SharesPropertyInterface
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

// DeserializeTentativeReject creates a TentativeReject from a map representation
// that has been unmarshalled from a text or binary format.
func DeserializeTentativeReject(m map[string]interface{}, aliasMap map[string]string) (*TentativeReject, error) {
	alias := ""
	aliasPrefix := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
		aliasPrefix = a + ":"
	}
	this := &TentativeReject{
		alias:   alias,
		unknown: make(map[string]interface{}),
	}
	if typeValue, ok := m["type"]; !ok {
		return nil, fmt.Errorf("no \"type\" property in map")
	} else if typeString, ok := typeValue.(string); ok {
		typeName := strings.TrimPrefix(typeString, aliasPrefix)
		if typeName != "TentativeReject" {
			return nil, fmt.Errorf("\"type\" property is not of %q type: %s", "TentativeReject", typeName)
		}
		// Fall through, success in finding a proper Type
	} else if arrType, ok := typeValue.([]interface{}); ok {
		found := false
		for _, elemVal := range arrType {
			if typeString, ok := elemVal.(string); ok && strings.TrimPrefix(typeString, aliasPrefix) == "TentativeReject" {
				found = true
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("could not find a \"type\" property of value %q", "TentativeReject")
		}
		// Fall through, success in finding a proper Type
	} else {
		return nil, fmt.Errorf("\"type\" property is unrecognized type: %T", typeValue)
	}
	// Begin: Known property deserialization
	if p, err := mgr.DeserializeActorPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.Actor = p
	}
	if p, err := mgr.DeserializeAltitudePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.Altitude = p
	}
	if p, err := mgr.DeserializeAttachmentPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.Attachment = p
	}
	if p, err := mgr.DeserializeAttributedToPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.AttributedTo = p
	}
	if p, err := mgr.DeserializeAudiencePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.Audience = p
	}
	if p, err := mgr.DeserializeBccPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.Bcc = p
	}
	if p, err := mgr.DeserializeBtoPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.Bto = p
	}
	if p, err := mgr.DeserializeCcPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.Cc = p
	}
	if p, err := mgr.DeserializeContentPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.Content = p
	}
	if p, err := mgr.DeserializeContextPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.Context = p
	}
	if p, err := mgr.DeserializeDurationPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.Duration = p
	}
	if p, err := mgr.DeserializeEndTimePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.EndTime = p
	}
	if p, err := mgr.DeserializeGeneratorPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.Generator = p
	}
	if p, err := mgr.DeserializeIconPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.Icon = p
	}
	if p, err := mgr.DeserializeIdPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.Id = p
	}
	if p, err := mgr.DeserializeImagePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.Image = p
	}
	if p, err := mgr.DeserializeInReplyToPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.InReplyTo = p
	}
	if p, err := mgr.DeserializeInstrumentPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.Instrument = p
	}
	if p, err := mgr.DeserializeLikesPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.Likes = p
	}
	if p, err := mgr.DeserializeLocationPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.Location = p
	}
	if p, err := mgr.DeserializeMediaTypePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.MediaType = p
	}
	if p, err := mgr.DeserializeNamePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.Name = p
	}
	if p, err := mgr.DeserializeObjectPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.Object = p
	}
	if p, err := mgr.DeserializeOriginPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.Origin = p
	}
	if p, err := mgr.DeserializePreviewPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.Preview = p
	}
	if p, err := mgr.DeserializePublishedPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.Published = p
	}
	if p, err := mgr.DeserializeRepliesPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.Replies = p
	}
	if p, err := mgr.DeserializeResultPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.Result = p
	}
	if p, err := mgr.DeserializeSharesPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.Shares = p
	}
	if p, err := mgr.DeserializeStartTimePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.StartTime = p
	}
	if p, err := mgr.DeserializeSummaryPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.Summary = p
	}
	if p, err := mgr.DeserializeTagPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.Tag = p
	}
	if p, err := mgr.DeserializeTargetPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.Target = p
	}
	if p, err := mgr.DeserializeToPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.To = p
	}
	if p, err := mgr.DeserializeTypePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.Type = p
	}
	if p, err := mgr.DeserializeUpdatedPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.Updated = p
	}
	if p, err := mgr.DeserializeUrlPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
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
		} else if k == "likes" {
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
		} else if k == "shares" {
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

// NewTentativeReject creates a new TentativeReject type
func NewTentativeReject() *TentativeReject {
	typeProp := propertytype.NewTypeProperty()
	typeProp.AppendString("TentativeReject")
	return &TentativeReject{
		Type:    typeProp,
		alias:   "",
		unknown: make(map[string]interface{}, 0),
	}
}

// TentativeRejectExtends returns true if the TentativeReject type extends from
// the other type.
func TentativeRejectExtends(other vocab.Type) bool {
	extensions := []string{"Activity", "Object", "Reject"}
	for _, ext := range extensions {
		if ext == other.GetName() {
			return true
		}
	}
	return false
}

// TentativeRejectIsDisjointWith returns true if the other provided type is
// disjoint with the TentativeReject type.
func TentativeRejectIsDisjointWith(other vocab.Type) bool {
	disjointWith := []string{"Link", "Mention"}
	for _, disjoint := range disjointWith {
		if disjoint == other.GetName() {
			return true
		}
	}
	return false
}

// TentativeRejectIsExtendedBy returns true if the other provided type extends
// from the TentativeReject type.
func TentativeRejectIsExtendedBy(other vocab.Type) bool {
	// Shortcut implementation: is not extended by anything.
	return false
}

// GetActor returns the "actor" property if it exists, and nil otherwise.
func (this TentativeReject) GetActor() vocab.ActorPropertyInterface {
	return this.Actor
}

// GetAltitude returns the "altitude" property if it exists, and nil otherwise.
func (this TentativeReject) GetAltitude() vocab.AltitudePropertyInterface {
	return this.Altitude
}

// GetAttachment returns the "attachment" property if it exists, and nil otherwise.
func (this TentativeReject) GetAttachment() vocab.AttachmentPropertyInterface {
	return this.Attachment
}

// GetAttributedTo returns the "attributedTo" property if it exists, and nil
// otherwise.
func (this TentativeReject) GetAttributedTo() vocab.AttributedToPropertyInterface {
	return this.AttributedTo
}

// GetAudience returns the "audience" property if it exists, and nil otherwise.
func (this TentativeReject) GetAudience() vocab.AudiencePropertyInterface {
	return this.Audience
}

// GetBcc returns the "bcc" property if it exists, and nil otherwise.
func (this TentativeReject) GetBcc() vocab.BccPropertyInterface {
	return this.Bcc
}

// GetBto returns the "bto" property if it exists, and nil otherwise.
func (this TentativeReject) GetBto() vocab.BtoPropertyInterface {
	return this.Bto
}

// GetCc returns the "cc" property if it exists, and nil otherwise.
func (this TentativeReject) GetCc() vocab.CcPropertyInterface {
	return this.Cc
}

// GetContent returns the "content" property if it exists, and nil otherwise.
func (this TentativeReject) GetContent() vocab.ContentPropertyInterface {
	return this.Content
}

// GetContext returns the "context" property if it exists, and nil otherwise.
func (this TentativeReject) GetContext() vocab.ContextPropertyInterface {
	return this.Context
}

// GetDuration returns the "duration" property if it exists, and nil otherwise.
func (this TentativeReject) GetDuration() vocab.DurationPropertyInterface {
	return this.Duration
}

// GetEndTime returns the "endTime" property if it exists, and nil otherwise.
func (this TentativeReject) GetEndTime() vocab.EndTimePropertyInterface {
	return this.EndTime
}

// GetGenerator returns the "generator" property if it exists, and nil otherwise.
func (this TentativeReject) GetGenerator() vocab.GeneratorPropertyInterface {
	return this.Generator
}

// GetIcon returns the "icon" property if it exists, and nil otherwise.
func (this TentativeReject) GetIcon() vocab.IconPropertyInterface {
	return this.Icon
}

// GetId returns the "id" property if it exists, and nil otherwise.
func (this TentativeReject) GetId() vocab.IdPropertyInterface {
	return this.Id
}

// GetImage returns the "image" property if it exists, and nil otherwise.
func (this TentativeReject) GetImage() vocab.ImagePropertyInterface {
	return this.Image
}

// GetInReplyTo returns the "inReplyTo" property if it exists, and nil otherwise.
func (this TentativeReject) GetInReplyTo() vocab.InReplyToPropertyInterface {
	return this.InReplyTo
}

// GetInstrument returns the "instrument" property if it exists, and nil otherwise.
func (this TentativeReject) GetInstrument() vocab.InstrumentPropertyInterface {
	return this.Instrument
}

// GetLikes returns the "likes" property if it exists, and nil otherwise.
func (this TentativeReject) GetLikes() vocab.LikesPropertyInterface {
	return this.Likes
}

// GetLocation returns the "location" property if it exists, and nil otherwise.
func (this TentativeReject) GetLocation() vocab.LocationPropertyInterface {
	return this.Location
}

// GetMediaType returns the "mediaType" property if it exists, and nil otherwise.
func (this TentativeReject) GetMediaType() vocab.MediaTypePropertyInterface {
	return this.MediaType
}

// GetName returns the "name" property if it exists, and nil otherwise.
func (this TentativeReject) GetName() vocab.NamePropertyInterface {
	return this.Name
}

// GetObject returns the "object" property if it exists, and nil otherwise.
func (this TentativeReject) GetObject() vocab.ObjectPropertyInterface {
	return this.Object
}

// GetOrigin returns the "origin" property if it exists, and nil otherwise.
func (this TentativeReject) GetOrigin() vocab.OriginPropertyInterface {
	return this.Origin
}

// GetPreview returns the "preview" property if it exists, and nil otherwise.
func (this TentativeReject) GetPreview() vocab.PreviewPropertyInterface {
	return this.Preview
}

// GetPublished returns the "published" property if it exists, and nil otherwise.
func (this TentativeReject) GetPublished() vocab.PublishedPropertyInterface {
	return this.Published
}

// GetReplies returns the "replies" property if it exists, and nil otherwise.
func (this TentativeReject) GetReplies() vocab.RepliesPropertyInterface {
	return this.Replies
}

// GetResult returns the "result" property if it exists, and nil otherwise.
func (this TentativeReject) GetResult() vocab.ResultPropertyInterface {
	return this.Result
}

// GetShares returns the "shares" property if it exists, and nil otherwise.
func (this TentativeReject) GetShares() vocab.SharesPropertyInterface {
	return this.Shares
}

// GetStartTime returns the "startTime" property if it exists, and nil otherwise.
func (this TentativeReject) GetStartTime() vocab.StartTimePropertyInterface {
	return this.StartTime
}

// GetSummary returns the "summary" property if it exists, and nil otherwise.
func (this TentativeReject) GetSummary() vocab.SummaryPropertyInterface {
	return this.Summary
}

// GetTag returns the "tag" property if it exists, and nil otherwise.
func (this TentativeReject) GetTag() vocab.TagPropertyInterface {
	return this.Tag
}

// GetTarget returns the "target" property if it exists, and nil otherwise.
func (this TentativeReject) GetTarget() vocab.TargetPropertyInterface {
	return this.Target
}

// GetTo returns the "to" property if it exists, and nil otherwise.
func (this TentativeReject) GetTo() vocab.ToPropertyInterface {
	return this.To
}

// GetType returns the "type" property if it exists, and nil otherwise.
func (this TentativeReject) GetType() vocab.TypePropertyInterface {
	return this.Type
}

// GetUnknownProperties returns the unknown properties for the TentativeReject
// type. Note that this should not be used by app developers. It is only used
// to help determine which implementation is LessThan the other. Developers
// who are creating a different implementation of this type's interface can
// use this method in their LessThan implementation, but routine ActivityPub
// applications should not use this to bypass the code generation tool.
func (this TentativeReject) GetUnknownProperties() map[string]interface{} {
	return this.unknown
}

// GetUpdated returns the "updated" property if it exists, and nil otherwise.
func (this TentativeReject) GetUpdated() vocab.UpdatedPropertyInterface {
	return this.Updated
}

// GetUrl returns the "url" property if it exists, and nil otherwise.
func (this TentativeReject) GetUrl() vocab.UrlPropertyInterface {
	return this.Url
}

// IsExtending returns true if the TentativeReject type extends from the other
// type.
func (this TentativeReject) IsExtending(other vocab.Type) bool {
	return TentativeRejectExtends(other)
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// type and the specific properties that are set. The value in the map is the
// alias used to import the type and its properties.
func (this TentativeReject) JSONLDContext() map[string]string {
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
	m = this.helperJSONLDContext(this.Likes, m)
	m = this.helperJSONLDContext(this.Location, m)
	m = this.helperJSONLDContext(this.MediaType, m)
	m = this.helperJSONLDContext(this.Name, m)
	m = this.helperJSONLDContext(this.Object, m)
	m = this.helperJSONLDContext(this.Origin, m)
	m = this.helperJSONLDContext(this.Preview, m)
	m = this.helperJSONLDContext(this.Published, m)
	m = this.helperJSONLDContext(this.Replies, m)
	m = this.helperJSONLDContext(this.Result, m)
	m = this.helperJSONLDContext(this.Shares, m)
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

// LessThan computes if this TentativeReject is lesser, with an arbitrary but
// stable determination.
func (this TentativeReject) LessThan(o vocab.TentativeRejectInterface) bool {
	// Begin: Compare known properties
	// Compare property "actor"
	if lhs, rhs := this.Actor, o.GetActor(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "altitude"
	if lhs, rhs := this.Altitude, o.GetAltitude(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "attachment"
	if lhs, rhs := this.Attachment, o.GetAttachment(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "attributedTo"
	if lhs, rhs := this.AttributedTo, o.GetAttributedTo(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "audience"
	if lhs, rhs := this.Audience, o.GetAudience(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "bcc"
	if lhs, rhs := this.Bcc, o.GetBcc(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "bto"
	if lhs, rhs := this.Bto, o.GetBto(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "cc"
	if lhs, rhs := this.Cc, o.GetCc(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "content"
	if lhs, rhs := this.Content, o.GetContent(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "context"
	if lhs, rhs := this.Context, o.GetContext(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "duration"
	if lhs, rhs := this.Duration, o.GetDuration(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "endTime"
	if lhs, rhs := this.EndTime, o.GetEndTime(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "generator"
	if lhs, rhs := this.Generator, o.GetGenerator(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "icon"
	if lhs, rhs := this.Icon, o.GetIcon(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "id"
	if lhs, rhs := this.Id, o.GetId(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "image"
	if lhs, rhs := this.Image, o.GetImage(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "inReplyTo"
	if lhs, rhs := this.InReplyTo, o.GetInReplyTo(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "instrument"
	if lhs, rhs := this.Instrument, o.GetInstrument(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "likes"
	if lhs, rhs := this.Likes, o.GetLikes(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "location"
	if lhs, rhs := this.Location, o.GetLocation(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "mediaType"
	if lhs, rhs := this.MediaType, o.GetMediaType(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "name"
	if lhs, rhs := this.Name, o.GetName(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "object"
	if lhs, rhs := this.Object, o.GetObject(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "origin"
	if lhs, rhs := this.Origin, o.GetOrigin(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "preview"
	if lhs, rhs := this.Preview, o.GetPreview(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "published"
	if lhs, rhs := this.Published, o.GetPublished(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "replies"
	if lhs, rhs := this.Replies, o.GetReplies(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "result"
	if lhs, rhs := this.Result, o.GetResult(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "shares"
	if lhs, rhs := this.Shares, o.GetShares(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "startTime"
	if lhs, rhs := this.StartTime, o.GetStartTime(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "summary"
	if lhs, rhs := this.Summary, o.GetSummary(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "tag"
	if lhs, rhs := this.Tag, o.GetTag(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "target"
	if lhs, rhs := this.Target, o.GetTarget(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "to"
	if lhs, rhs := this.To, o.GetTo(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "type"
	if lhs, rhs := this.Type, o.GetType(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "updated"
	if lhs, rhs := this.Updated, o.GetUpdated(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "url"
	if lhs, rhs := this.Url, o.GetUrl(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
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
func (this TentativeReject) Serialize() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	typeName := "TentativeReject"
	if len(this.alias) > 0 {
		typeName = this.alias + ":" + "TentativeReject"
	}
	m["type"] = typeName
	// Begin: Serialize known properties
	// Maybe serialize property "actor"
	if this.Actor != nil {
		if i, err := this.Actor.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.Actor.Name()] = i
		}
	}
	// Maybe serialize property "altitude"
	if this.Altitude != nil {
		if i, err := this.Altitude.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.Altitude.Name()] = i
		}
	}
	// Maybe serialize property "attachment"
	if this.Attachment != nil {
		if i, err := this.Attachment.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.Attachment.Name()] = i
		}
	}
	// Maybe serialize property "attributedTo"
	if this.AttributedTo != nil {
		if i, err := this.AttributedTo.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.AttributedTo.Name()] = i
		}
	}
	// Maybe serialize property "audience"
	if this.Audience != nil {
		if i, err := this.Audience.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.Audience.Name()] = i
		}
	}
	// Maybe serialize property "bcc"
	if this.Bcc != nil {
		if i, err := this.Bcc.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.Bcc.Name()] = i
		}
	}
	// Maybe serialize property "bto"
	if this.Bto != nil {
		if i, err := this.Bto.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.Bto.Name()] = i
		}
	}
	// Maybe serialize property "cc"
	if this.Cc != nil {
		if i, err := this.Cc.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.Cc.Name()] = i
		}
	}
	// Maybe serialize property "content"
	if this.Content != nil {
		if i, err := this.Content.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.Content.Name()] = i
		}
	}
	// Maybe serialize property "context"
	if this.Context != nil {
		if i, err := this.Context.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.Context.Name()] = i
		}
	}
	// Maybe serialize property "duration"
	if this.Duration != nil {
		if i, err := this.Duration.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.Duration.Name()] = i
		}
	}
	// Maybe serialize property "endTime"
	if this.EndTime != nil {
		if i, err := this.EndTime.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.EndTime.Name()] = i
		}
	}
	// Maybe serialize property "generator"
	if this.Generator != nil {
		if i, err := this.Generator.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.Generator.Name()] = i
		}
	}
	// Maybe serialize property "icon"
	if this.Icon != nil {
		if i, err := this.Icon.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.Icon.Name()] = i
		}
	}
	// Maybe serialize property "id"
	if this.Id != nil {
		if i, err := this.Id.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.Id.Name()] = i
		}
	}
	// Maybe serialize property "image"
	if this.Image != nil {
		if i, err := this.Image.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.Image.Name()] = i
		}
	}
	// Maybe serialize property "inReplyTo"
	if this.InReplyTo != nil {
		if i, err := this.InReplyTo.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.InReplyTo.Name()] = i
		}
	}
	// Maybe serialize property "instrument"
	if this.Instrument != nil {
		if i, err := this.Instrument.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.Instrument.Name()] = i
		}
	}
	// Maybe serialize property "likes"
	if this.Likes != nil {
		if i, err := this.Likes.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.Likes.Name()] = i
		}
	}
	// Maybe serialize property "location"
	if this.Location != nil {
		if i, err := this.Location.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.Location.Name()] = i
		}
	}
	// Maybe serialize property "mediaType"
	if this.MediaType != nil {
		if i, err := this.MediaType.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.MediaType.Name()] = i
		}
	}
	// Maybe serialize property "name"
	if this.Name != nil {
		if i, err := this.Name.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.Name.Name()] = i
		}
	}
	// Maybe serialize property "object"
	if this.Object != nil {
		if i, err := this.Object.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.Object.Name()] = i
		}
	}
	// Maybe serialize property "origin"
	if this.Origin != nil {
		if i, err := this.Origin.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.Origin.Name()] = i
		}
	}
	// Maybe serialize property "preview"
	if this.Preview != nil {
		if i, err := this.Preview.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.Preview.Name()] = i
		}
	}
	// Maybe serialize property "published"
	if this.Published != nil {
		if i, err := this.Published.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.Published.Name()] = i
		}
	}
	// Maybe serialize property "replies"
	if this.Replies != nil {
		if i, err := this.Replies.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.Replies.Name()] = i
		}
	}
	// Maybe serialize property "result"
	if this.Result != nil {
		if i, err := this.Result.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.Result.Name()] = i
		}
	}
	// Maybe serialize property "shares"
	if this.Shares != nil {
		if i, err := this.Shares.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.Shares.Name()] = i
		}
	}
	// Maybe serialize property "startTime"
	if this.StartTime != nil {
		if i, err := this.StartTime.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.StartTime.Name()] = i
		}
	}
	// Maybe serialize property "summary"
	if this.Summary != nil {
		if i, err := this.Summary.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.Summary.Name()] = i
		}
	}
	// Maybe serialize property "tag"
	if this.Tag != nil {
		if i, err := this.Tag.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.Tag.Name()] = i
		}
	}
	// Maybe serialize property "target"
	if this.Target != nil {
		if i, err := this.Target.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.Target.Name()] = i
		}
	}
	// Maybe serialize property "to"
	if this.To != nil {
		if i, err := this.To.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.To.Name()] = i
		}
	}
	// Maybe serialize property "type"
	if this.Type != nil {
		if i, err := this.Type.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.Type.Name()] = i
		}
	}
	// Maybe serialize property "updated"
	if this.Updated != nil {
		if i, err := this.Updated.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.Updated.Name()] = i
		}
	}
	// Maybe serialize property "url"
	if this.Url != nil {
		if i, err := this.Url.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.Url.Name()] = i
		}
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
func (this *TentativeReject) SetActor(i vocab.ActorPropertyInterface) {
	this.Actor = i
}

// SetAltitude returns the "altitude" property if it exists, and nil otherwise.
func (this *TentativeReject) SetAltitude(i vocab.AltitudePropertyInterface) {
	this.Altitude = i
}

// SetAttachment returns the "attachment" property if it exists, and nil otherwise.
func (this *TentativeReject) SetAttachment(i vocab.AttachmentPropertyInterface) {
	this.Attachment = i
}

// SetAttributedTo returns the "attributedTo" property if it exists, and nil
// otherwise.
func (this *TentativeReject) SetAttributedTo(i vocab.AttributedToPropertyInterface) {
	this.AttributedTo = i
}

// SetAudience returns the "audience" property if it exists, and nil otherwise.
func (this *TentativeReject) SetAudience(i vocab.AudiencePropertyInterface) {
	this.Audience = i
}

// SetBcc returns the "bcc" property if it exists, and nil otherwise.
func (this *TentativeReject) SetBcc(i vocab.BccPropertyInterface) {
	this.Bcc = i
}

// SetBto returns the "bto" property if it exists, and nil otherwise.
func (this *TentativeReject) SetBto(i vocab.BtoPropertyInterface) {
	this.Bto = i
}

// SetCc returns the "cc" property if it exists, and nil otherwise.
func (this *TentativeReject) SetCc(i vocab.CcPropertyInterface) {
	this.Cc = i
}

// SetContent returns the "content" property if it exists, and nil otherwise.
func (this *TentativeReject) SetContent(i vocab.ContentPropertyInterface) {
	this.Content = i
}

// SetContext returns the "context" property if it exists, and nil otherwise.
func (this *TentativeReject) SetContext(i vocab.ContextPropertyInterface) {
	this.Context = i
}

// SetDuration returns the "duration" property if it exists, and nil otherwise.
func (this *TentativeReject) SetDuration(i vocab.DurationPropertyInterface) {
	this.Duration = i
}

// SetEndTime returns the "endTime" property if it exists, and nil otherwise.
func (this *TentativeReject) SetEndTime(i vocab.EndTimePropertyInterface) {
	this.EndTime = i
}

// SetGenerator returns the "generator" property if it exists, and nil otherwise.
func (this *TentativeReject) SetGenerator(i vocab.GeneratorPropertyInterface) {
	this.Generator = i
}

// SetIcon returns the "icon" property if it exists, and nil otherwise.
func (this *TentativeReject) SetIcon(i vocab.IconPropertyInterface) {
	this.Icon = i
}

// SetId returns the "id" property if it exists, and nil otherwise.
func (this *TentativeReject) SetId(i vocab.IdPropertyInterface) {
	this.Id = i
}

// SetImage returns the "image" property if it exists, and nil otherwise.
func (this *TentativeReject) SetImage(i vocab.ImagePropertyInterface) {
	this.Image = i
}

// SetInReplyTo returns the "inReplyTo" property if it exists, and nil otherwise.
func (this *TentativeReject) SetInReplyTo(i vocab.InReplyToPropertyInterface) {
	this.InReplyTo = i
}

// SetInstrument returns the "instrument" property if it exists, and nil otherwise.
func (this *TentativeReject) SetInstrument(i vocab.InstrumentPropertyInterface) {
	this.Instrument = i
}

// SetLikes returns the "likes" property if it exists, and nil otherwise.
func (this *TentativeReject) SetLikes(i vocab.LikesPropertyInterface) {
	this.Likes = i
}

// SetLocation returns the "location" property if it exists, and nil otherwise.
func (this *TentativeReject) SetLocation(i vocab.LocationPropertyInterface) {
	this.Location = i
}

// SetMediaType returns the "mediaType" property if it exists, and nil otherwise.
func (this *TentativeReject) SetMediaType(i vocab.MediaTypePropertyInterface) {
	this.MediaType = i
}

// SetName returns the "name" property if it exists, and nil otherwise.
func (this *TentativeReject) SetName(i vocab.NamePropertyInterface) {
	this.Name = i
}

// SetObject returns the "object" property if it exists, and nil otherwise.
func (this *TentativeReject) SetObject(i vocab.ObjectPropertyInterface) {
	this.Object = i
}

// SetOrigin returns the "origin" property if it exists, and nil otherwise.
func (this *TentativeReject) SetOrigin(i vocab.OriginPropertyInterface) {
	this.Origin = i
}

// SetPreview returns the "preview" property if it exists, and nil otherwise.
func (this *TentativeReject) SetPreview(i vocab.PreviewPropertyInterface) {
	this.Preview = i
}

// SetPublished returns the "published" property if it exists, and nil otherwise.
func (this *TentativeReject) SetPublished(i vocab.PublishedPropertyInterface) {
	this.Published = i
}

// SetReplies returns the "replies" property if it exists, and nil otherwise.
func (this *TentativeReject) SetReplies(i vocab.RepliesPropertyInterface) {
	this.Replies = i
}

// SetResult returns the "result" property if it exists, and nil otherwise.
func (this *TentativeReject) SetResult(i vocab.ResultPropertyInterface) {
	this.Result = i
}

// SetShares returns the "shares" property if it exists, and nil otherwise.
func (this *TentativeReject) SetShares(i vocab.SharesPropertyInterface) {
	this.Shares = i
}

// SetStartTime returns the "startTime" property if it exists, and nil otherwise.
func (this *TentativeReject) SetStartTime(i vocab.StartTimePropertyInterface) {
	this.StartTime = i
}

// SetSummary returns the "summary" property if it exists, and nil otherwise.
func (this *TentativeReject) SetSummary(i vocab.SummaryPropertyInterface) {
	this.Summary = i
}

// SetTag returns the "tag" property if it exists, and nil otherwise.
func (this *TentativeReject) SetTag(i vocab.TagPropertyInterface) {
	this.Tag = i
}

// SetTarget returns the "target" property if it exists, and nil otherwise.
func (this *TentativeReject) SetTarget(i vocab.TargetPropertyInterface) {
	this.Target = i
}

// SetTo returns the "to" property if it exists, and nil otherwise.
func (this *TentativeReject) SetTo(i vocab.ToPropertyInterface) {
	this.To = i
}

// SetType returns the "type" property if it exists, and nil otherwise.
func (this *TentativeReject) SetType(i vocab.TypePropertyInterface) {
	this.Type = i
}

// SetUpdated returns the "updated" property if it exists, and nil otherwise.
func (this *TentativeReject) SetUpdated(i vocab.UpdatedPropertyInterface) {
	this.Updated = i
}

// SetUrl returns the "url" property if it exists, and nil otherwise.
func (this *TentativeReject) SetUrl(i vocab.UrlPropertyInterface) {
	this.Url = i
}

// helperJSONLDContext obtains the context uris and their aliases from a property,
// if it is not nil.
func (this TentativeReject) helperJSONLDContext(i jsonldContexter, toMerge map[string]string) map[string]string {
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

package typecollectionpage

import (
	"fmt"
	vocab "github.com/go-fed/activity/streams/vocab"
	"strings"
)

// Used to represent distinct subsets of items from a Collection. Refer to the
// Activity Streams 2.0 Core for a complete description of the CollectionPage
// object.
//
// Example 7 (https://www.w3.org/TR/activitystreams-vocabulary/#ex6b-jsonld):
//   {
//     "id": "http://example.org/foo?page=1",
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
//     "partOf": "http://example.org/foo",
//     "summary": "Page 1 of Sally's notes",
//     "type": "CollectionPage"
//   }
type CollectionPage struct {
	Altitude     vocab.AltitudePropertyInterface
	Attachment   vocab.AttachmentPropertyInterface
	AttributedTo vocab.AttributedToPropertyInterface
	Audience     vocab.AudiencePropertyInterface
	Bcc          vocab.BccPropertyInterface
	Bto          vocab.BtoPropertyInterface
	Cc           vocab.CcPropertyInterface
	Content      vocab.ContentPropertyInterface
	Context      vocab.ContextPropertyInterface
	Current      vocab.CurrentPropertyInterface
	Duration     vocab.DurationPropertyInterface
	EndTime      vocab.EndTimePropertyInterface
	First        vocab.FirstPropertyInterface
	Generator    vocab.GeneratorPropertyInterface
	Icon         vocab.IconPropertyInterface
	Id           vocab.IdPropertyInterface
	Image        vocab.ImagePropertyInterface
	InReplyTo    vocab.InReplyToPropertyInterface
	Items        vocab.ItemsPropertyInterface
	Last         vocab.LastPropertyInterface
	Likes        vocab.LikesPropertyInterface
	Location     vocab.LocationPropertyInterface
	MediaType    vocab.MediaTypePropertyInterface
	Name         vocab.NamePropertyInterface
	Next         vocab.NextPropertyInterface
	Object       vocab.ObjectPropertyInterface
	PartOf       vocab.PartOfPropertyInterface
	Prev         vocab.PrevPropertyInterface
	Preview      vocab.PreviewPropertyInterface
	Published    vocab.PublishedPropertyInterface
	Replies      vocab.RepliesPropertyInterface
	Shares       vocab.SharesPropertyInterface
	StartTime    vocab.StartTimePropertyInterface
	Summary      vocab.SummaryPropertyInterface
	Tag          vocab.TagPropertyInterface
	To           vocab.ToPropertyInterface
	TotalItems   vocab.TotalItemsPropertyInterface
	Type         vocab.TypePropertyInterface
	Updated      vocab.UpdatedPropertyInterface
	Url          vocab.UrlPropertyInterface
	alias        string
	unknown      map[string]interface{}
}

// CollectionPageExtends returns true if the CollectionPage type extends from the
// other type.
func CollectionPageExtends(other vocab.Type) bool {
	extensions := []string{"Collection", "Object"}
	for _, ext := range extensions {
		if ext == other.GetName() {
			return true
		}
	}
	return false
}

// CollectionPageIsDisjointWith returns true if the other provided type is
// disjoint with the CollectionPage type.
func CollectionPageIsDisjointWith(other vocab.Type) bool {
	disjointWith := []string{"Link", "Mention"}
	for _, disjoint := range disjointWith {
		if disjoint == other.GetName() {
			return true
		}
	}
	return false
}

// CollectionPageIsExtendedBy returns true if the other provided type extends from
// the CollectionPage type.
func CollectionPageIsExtendedBy(other vocab.Type) bool {
	extensions := []string{"OrderedCollectionPage"}
	for _, ext := range extensions {
		if ext == other.GetName() {
			return true
		}
	}
	return false
}

// DeserializeCollectionPage creates a CollectionPage from a map representation
// that has been unmarshalled from a text or binary format.
func DeserializeCollectionPage(m map[string]interface{}, aliasMap map[string]string) (*CollectionPage, error) {
	alias := ""
	aliasPrefix := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
		aliasPrefix = a + ":"
	}
	this := &CollectionPage{
		alias:   alias,
		unknown: make(map[string]interface{}),
	}
	if typeValue, ok := m["type"]; !ok {
		return nil, fmt.Errorf("no \"type\" property in map")
	} else if typeString, ok := typeValue.(string); ok {
		typeName := strings.TrimPrefix(typeString, aliasPrefix)
		if typeName != "CollectionPage" {
			return nil, fmt.Errorf("\"type\" property is not of %q type: %s", "CollectionPage", typeName)
		}
		// Fall through, success in finding a proper Type
	} else if arrType, ok := typeValue.([]interface{}); ok {
		found := false
		for _, elemVal := range arrType {
			if typeString, ok := elemVal.(string); ok && strings.TrimPrefix(typeString, aliasPrefix) == "CollectionPage" {
				found = true
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("could not find a \"type\" property of value %q", "CollectionPage")
		}
		// Fall through, success in finding a proper Type
	} else {
		return nil, fmt.Errorf("\"type\" property is unrecognized type: %T", typeValue)
	}
	// Begin: Known property deserialization
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
	if p, err := mgr.DeserializeCurrentPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.Current = p
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
	if p, err := mgr.DeserializeFirstPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.First = p
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
	if p, err := mgr.DeserializeItemsPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.Items = p
	}
	if p, err := mgr.DeserializeLastPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.Last = p
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
	if p, err := mgr.DeserializeNextPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.Next = p
	}
	if p, err := mgr.DeserializeObjectPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.Object = p
	}
	if p, err := mgr.DeserializePartOfPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.PartOf = p
	}
	if p, err := mgr.DeserializePrevPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.Prev = p
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
	if p, err := mgr.DeserializeToPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.To = p
	}
	if p, err := mgr.DeserializeTotalItemsPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.TotalItems = p
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
		if k == "altitude" {
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
		} else if k == "current" {
			continue
		} else if k == "duration" {
			continue
		} else if k == "endTime" {
			continue
		} else if k == "first" {
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
		} else if k == "items" {
			continue
		} else if k == "last" {
			continue
		} else if k == "likes" {
			continue
		} else if k == "location" {
			continue
		} else if k == "mediaType" {
			continue
		} else if k == "name" {
			continue
		} else if k == "next" {
			continue
		} else if k == "object" {
			continue
		} else if k == "partOf" {
			continue
		} else if k == "prev" {
			continue
		} else if k == "preview" {
			continue
		} else if k == "published" {
			continue
		} else if k == "replies" {
			continue
		} else if k == "shares" {
			continue
		} else if k == "startTime" {
			continue
		} else if k == "summary" {
			continue
		} else if k == "tag" {
			continue
		} else if k == "to" {
			continue
		} else if k == "totalItems" {
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

// NewCollectionPage creates a new CollectionPage type
func NewCollectionPage() *CollectionPage {
	return &CollectionPage{
		alias:   "",
		unknown: make(map[string]interface{}, 0),
	}
}

// GetAltitude returns the "altitude" property if it exists, and nil otherwise.
func (this CollectionPage) GetAltitude() vocab.AltitudePropertyInterface {
	return this.Altitude
}

// GetAttachment returns the "attachment" property if it exists, and nil otherwise.
func (this CollectionPage) GetAttachment() vocab.AttachmentPropertyInterface {
	return this.Attachment
}

// GetAttributedTo returns the "attributedTo" property if it exists, and nil
// otherwise.
func (this CollectionPage) GetAttributedTo() vocab.AttributedToPropertyInterface {
	return this.AttributedTo
}

// GetAudience returns the "audience" property if it exists, and nil otherwise.
func (this CollectionPage) GetAudience() vocab.AudiencePropertyInterface {
	return this.Audience
}

// GetBcc returns the "bcc" property if it exists, and nil otherwise.
func (this CollectionPage) GetBcc() vocab.BccPropertyInterface {
	return this.Bcc
}

// GetBto returns the "bto" property if it exists, and nil otherwise.
func (this CollectionPage) GetBto() vocab.BtoPropertyInterface {
	return this.Bto
}

// GetCc returns the "cc" property if it exists, and nil otherwise.
func (this CollectionPage) GetCc() vocab.CcPropertyInterface {
	return this.Cc
}

// GetContent returns the "content" property if it exists, and nil otherwise.
func (this CollectionPage) GetContent() vocab.ContentPropertyInterface {
	return this.Content
}

// GetContext returns the "context" property if it exists, and nil otherwise.
func (this CollectionPage) GetContext() vocab.ContextPropertyInterface {
	return this.Context
}

// GetCurrent returns the "current" property if it exists, and nil otherwise.
func (this CollectionPage) GetCurrent() vocab.CurrentPropertyInterface {
	return this.Current
}

// GetDuration returns the "duration" property if it exists, and nil otherwise.
func (this CollectionPage) GetDuration() vocab.DurationPropertyInterface {
	return this.Duration
}

// GetEndTime returns the "endTime" property if it exists, and nil otherwise.
func (this CollectionPage) GetEndTime() vocab.EndTimePropertyInterface {
	return this.EndTime
}

// GetFirst returns the "first" property if it exists, and nil otherwise.
func (this CollectionPage) GetFirst() vocab.FirstPropertyInterface {
	return this.First
}

// GetGenerator returns the "generator" property if it exists, and nil otherwise.
func (this CollectionPage) GetGenerator() vocab.GeneratorPropertyInterface {
	return this.Generator
}

// GetIcon returns the "icon" property if it exists, and nil otherwise.
func (this CollectionPage) GetIcon() vocab.IconPropertyInterface {
	return this.Icon
}

// GetId returns the "id" property if it exists, and nil otherwise.
func (this CollectionPage) GetId() vocab.IdPropertyInterface {
	return this.Id
}

// GetImage returns the "image" property if it exists, and nil otherwise.
func (this CollectionPage) GetImage() vocab.ImagePropertyInterface {
	return this.Image
}

// GetInReplyTo returns the "inReplyTo" property if it exists, and nil otherwise.
func (this CollectionPage) GetInReplyTo() vocab.InReplyToPropertyInterface {
	return this.InReplyTo
}

// GetItems returns the "items" property if it exists, and nil otherwise.
func (this CollectionPage) GetItems() vocab.ItemsPropertyInterface {
	return this.Items
}

// GetLast returns the "last" property if it exists, and nil otherwise.
func (this CollectionPage) GetLast() vocab.LastPropertyInterface {
	return this.Last
}

// GetLikes returns the "likes" property if it exists, and nil otherwise.
func (this CollectionPage) GetLikes() vocab.LikesPropertyInterface {
	return this.Likes
}

// GetLocation returns the "location" property if it exists, and nil otherwise.
func (this CollectionPage) GetLocation() vocab.LocationPropertyInterface {
	return this.Location
}

// GetMediaType returns the "mediaType" property if it exists, and nil otherwise.
func (this CollectionPage) GetMediaType() vocab.MediaTypePropertyInterface {
	return this.MediaType
}

// GetName returns the "name" property if it exists, and nil otherwise.
func (this CollectionPage) GetName() vocab.NamePropertyInterface {
	return this.Name
}

// GetNext returns the "next" property if it exists, and nil otherwise.
func (this CollectionPage) GetNext() vocab.NextPropertyInterface {
	return this.Next
}

// GetObject returns the "object" property if it exists, and nil otherwise.
func (this CollectionPage) GetObject() vocab.ObjectPropertyInterface {
	return this.Object
}

// GetPartOf returns the "partOf" property if it exists, and nil otherwise.
func (this CollectionPage) GetPartOf() vocab.PartOfPropertyInterface {
	return this.PartOf
}

// GetPrev returns the "prev" property if it exists, and nil otherwise.
func (this CollectionPage) GetPrev() vocab.PrevPropertyInterface {
	return this.Prev
}

// GetPreview returns the "preview" property if it exists, and nil otherwise.
func (this CollectionPage) GetPreview() vocab.PreviewPropertyInterface {
	return this.Preview
}

// GetPublished returns the "published" property if it exists, and nil otherwise.
func (this CollectionPage) GetPublished() vocab.PublishedPropertyInterface {
	return this.Published
}

// GetReplies returns the "replies" property if it exists, and nil otherwise.
func (this CollectionPage) GetReplies() vocab.RepliesPropertyInterface {
	return this.Replies
}

// GetShares returns the "shares" property if it exists, and nil otherwise.
func (this CollectionPage) GetShares() vocab.SharesPropertyInterface {
	return this.Shares
}

// GetStartTime returns the "startTime" property if it exists, and nil otherwise.
func (this CollectionPage) GetStartTime() vocab.StartTimePropertyInterface {
	return this.StartTime
}

// GetSummary returns the "summary" property if it exists, and nil otherwise.
func (this CollectionPage) GetSummary() vocab.SummaryPropertyInterface {
	return this.Summary
}

// GetTag returns the "tag" property if it exists, and nil otherwise.
func (this CollectionPage) GetTag() vocab.TagPropertyInterface {
	return this.Tag
}

// GetTo returns the "to" property if it exists, and nil otherwise.
func (this CollectionPage) GetTo() vocab.ToPropertyInterface {
	return this.To
}

// GetTotalItems returns the "totalItems" property if it exists, and nil otherwise.
func (this CollectionPage) GetTotalItems() vocab.TotalItemsPropertyInterface {
	return this.TotalItems
}

// GetType returns the "type" property if it exists, and nil otherwise.
func (this CollectionPage) GetType() vocab.TypePropertyInterface {
	return this.Type
}

// GetUnknownProperties returns the unknown properties for the CollectionPage
// type. Note that this should not be used by app developers. It is only used
// to help determine which implementation is LessThan the other. Developers
// who are creating a different implementation of this type's interface can
// use this method in their LessThan implementation, but routine ActivityPub
// applications should not use this to bypass the code generation tool.
func (this CollectionPage) GetUnknownProperties() map[string]interface{} {
	return this.unknown
}

// GetUpdated returns the "updated" property if it exists, and nil otherwise.
func (this CollectionPage) GetUpdated() vocab.UpdatedPropertyInterface {
	return this.Updated
}

// GetUrl returns the "url" property if it exists, and nil otherwise.
func (this CollectionPage) GetUrl() vocab.UrlPropertyInterface {
	return this.Url
}

// IsExtending returns true if the CollectionPage type extends from the other type.
func (this CollectionPage) IsExtending(other vocab.Type) bool {
	return CollectionPageExtends(other)
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// type and the specific properties that are set. The value in the map is the
// alias used to import the type and its properties.
func (this CollectionPage) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/TR/activitystreams-vocabulary": this.alias}
	m = this.helperJSONLDContext(this.Altitude, m)
	m = this.helperJSONLDContext(this.Attachment, m)
	m = this.helperJSONLDContext(this.AttributedTo, m)
	m = this.helperJSONLDContext(this.Audience, m)
	m = this.helperJSONLDContext(this.Bcc, m)
	m = this.helperJSONLDContext(this.Bto, m)
	m = this.helperJSONLDContext(this.Cc, m)
	m = this.helperJSONLDContext(this.Content, m)
	m = this.helperJSONLDContext(this.Context, m)
	m = this.helperJSONLDContext(this.Current, m)
	m = this.helperJSONLDContext(this.Duration, m)
	m = this.helperJSONLDContext(this.EndTime, m)
	m = this.helperJSONLDContext(this.First, m)
	m = this.helperJSONLDContext(this.Generator, m)
	m = this.helperJSONLDContext(this.Icon, m)
	m = this.helperJSONLDContext(this.Id, m)
	m = this.helperJSONLDContext(this.Image, m)
	m = this.helperJSONLDContext(this.InReplyTo, m)
	m = this.helperJSONLDContext(this.Items, m)
	m = this.helperJSONLDContext(this.Last, m)
	m = this.helperJSONLDContext(this.Likes, m)
	m = this.helperJSONLDContext(this.Location, m)
	m = this.helperJSONLDContext(this.MediaType, m)
	m = this.helperJSONLDContext(this.Name, m)
	m = this.helperJSONLDContext(this.Next, m)
	m = this.helperJSONLDContext(this.Object, m)
	m = this.helperJSONLDContext(this.PartOf, m)
	m = this.helperJSONLDContext(this.Prev, m)
	m = this.helperJSONLDContext(this.Preview, m)
	m = this.helperJSONLDContext(this.Published, m)
	m = this.helperJSONLDContext(this.Replies, m)
	m = this.helperJSONLDContext(this.Shares, m)
	m = this.helperJSONLDContext(this.StartTime, m)
	m = this.helperJSONLDContext(this.Summary, m)
	m = this.helperJSONLDContext(this.Tag, m)
	m = this.helperJSONLDContext(this.To, m)
	m = this.helperJSONLDContext(this.TotalItems, m)
	m = this.helperJSONLDContext(this.Type, m)
	m = this.helperJSONLDContext(this.Updated, m)
	m = this.helperJSONLDContext(this.Url, m)

	return m
}

// LessThan computes if this CollectionPage is lesser, with an arbitrary but
// stable determination.
func (this CollectionPage) LessThan(o vocab.CollectionPageInterface) bool {
	// Begin: Compare known properties
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
	// Compare property "current"
	if lhs, rhs := this.Current, o.GetCurrent(); lhs != nil && rhs != nil {
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
	// Compare property "first"
	if lhs, rhs := this.First, o.GetFirst(); lhs != nil && rhs != nil {
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
	// Compare property "items"
	if lhs, rhs := this.Items, o.GetItems(); lhs != nil && rhs != nil {
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
	// Compare property "last"
	if lhs, rhs := this.Last, o.GetLast(); lhs != nil && rhs != nil {
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
	// Compare property "next"
	if lhs, rhs := this.Next, o.GetNext(); lhs != nil && rhs != nil {
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
	// Compare property "partOf"
	if lhs, rhs := this.PartOf, o.GetPartOf(); lhs != nil && rhs != nil {
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
	// Compare property "prev"
	if lhs, rhs := this.Prev, o.GetPrev(); lhs != nil && rhs != nil {
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
	// Compare property "totalItems"
	if lhs, rhs := this.TotalItems, o.GetTotalItems(); lhs != nil && rhs != nil {
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
func (this CollectionPage) Serialize() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	typeName := "CollectionPage"
	if len(this.alias) > 0 {
		typeName = this.alias + ":" + "CollectionPage"
	}
	m["type"] = typeName
	// Begin: Serialize known properties
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
	// Maybe serialize property "current"
	if this.Current != nil {
		if i, err := this.Current.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.Current.Name()] = i
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
	// Maybe serialize property "first"
	if this.First != nil {
		if i, err := this.First.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.First.Name()] = i
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
	// Maybe serialize property "items"
	if this.Items != nil {
		if i, err := this.Items.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.Items.Name()] = i
		}
	}
	// Maybe serialize property "last"
	if this.Last != nil {
		if i, err := this.Last.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.Last.Name()] = i
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
	// Maybe serialize property "next"
	if this.Next != nil {
		if i, err := this.Next.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.Next.Name()] = i
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
	// Maybe serialize property "partOf"
	if this.PartOf != nil {
		if i, err := this.PartOf.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.PartOf.Name()] = i
		}
	}
	// Maybe serialize property "prev"
	if this.Prev != nil {
		if i, err := this.Prev.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.Prev.Name()] = i
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
	// Maybe serialize property "to"
	if this.To != nil {
		if i, err := this.To.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.To.Name()] = i
		}
	}
	// Maybe serialize property "totalItems"
	if this.TotalItems != nil {
		if i, err := this.TotalItems.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.TotalItems.Name()] = i
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

// SetAltitude returns the "altitude" property if it exists, and nil otherwise.
func (this *CollectionPage) SetAltitude(i vocab.AltitudePropertyInterface) {
	this.Altitude = i
}

// SetAttachment returns the "attachment" property if it exists, and nil otherwise.
func (this *CollectionPage) SetAttachment(i vocab.AttachmentPropertyInterface) {
	this.Attachment = i
}

// SetAttributedTo returns the "attributedTo" property if it exists, and nil
// otherwise.
func (this *CollectionPage) SetAttributedTo(i vocab.AttributedToPropertyInterface) {
	this.AttributedTo = i
}

// SetAudience returns the "audience" property if it exists, and nil otherwise.
func (this *CollectionPage) SetAudience(i vocab.AudiencePropertyInterface) {
	this.Audience = i
}

// SetBcc returns the "bcc" property if it exists, and nil otherwise.
func (this *CollectionPage) SetBcc(i vocab.BccPropertyInterface) {
	this.Bcc = i
}

// SetBto returns the "bto" property if it exists, and nil otherwise.
func (this *CollectionPage) SetBto(i vocab.BtoPropertyInterface) {
	this.Bto = i
}

// SetCc returns the "cc" property if it exists, and nil otherwise.
func (this *CollectionPage) SetCc(i vocab.CcPropertyInterface) {
	this.Cc = i
}

// SetContent returns the "content" property if it exists, and nil otherwise.
func (this *CollectionPage) SetContent(i vocab.ContentPropertyInterface) {
	this.Content = i
}

// SetContext returns the "context" property if it exists, and nil otherwise.
func (this *CollectionPage) SetContext(i vocab.ContextPropertyInterface) {
	this.Context = i
}

// SetCurrent returns the "current" property if it exists, and nil otherwise.
func (this *CollectionPage) SetCurrent(i vocab.CurrentPropertyInterface) {
	this.Current = i
}

// SetDuration returns the "duration" property if it exists, and nil otherwise.
func (this *CollectionPage) SetDuration(i vocab.DurationPropertyInterface) {
	this.Duration = i
}

// SetEndTime returns the "endTime" property if it exists, and nil otherwise.
func (this *CollectionPage) SetEndTime(i vocab.EndTimePropertyInterface) {
	this.EndTime = i
}

// SetFirst returns the "first" property if it exists, and nil otherwise.
func (this *CollectionPage) SetFirst(i vocab.FirstPropertyInterface) {
	this.First = i
}

// SetGenerator returns the "generator" property if it exists, and nil otherwise.
func (this *CollectionPage) SetGenerator(i vocab.GeneratorPropertyInterface) {
	this.Generator = i
}

// SetIcon returns the "icon" property if it exists, and nil otherwise.
func (this *CollectionPage) SetIcon(i vocab.IconPropertyInterface) {
	this.Icon = i
}

// SetId returns the "id" property if it exists, and nil otherwise.
func (this *CollectionPage) SetId(i vocab.IdPropertyInterface) {
	this.Id = i
}

// SetImage returns the "image" property if it exists, and nil otherwise.
func (this *CollectionPage) SetImage(i vocab.ImagePropertyInterface) {
	this.Image = i
}

// SetInReplyTo returns the "inReplyTo" property if it exists, and nil otherwise.
func (this *CollectionPage) SetInReplyTo(i vocab.InReplyToPropertyInterface) {
	this.InReplyTo = i
}

// SetItems returns the "items" property if it exists, and nil otherwise.
func (this *CollectionPage) SetItems(i vocab.ItemsPropertyInterface) {
	this.Items = i
}

// SetLast returns the "last" property if it exists, and nil otherwise.
func (this *CollectionPage) SetLast(i vocab.LastPropertyInterface) {
	this.Last = i
}

// SetLikes returns the "likes" property if it exists, and nil otherwise.
func (this *CollectionPage) SetLikes(i vocab.LikesPropertyInterface) {
	this.Likes = i
}

// SetLocation returns the "location" property if it exists, and nil otherwise.
func (this *CollectionPage) SetLocation(i vocab.LocationPropertyInterface) {
	this.Location = i
}

// SetMediaType returns the "mediaType" property if it exists, and nil otherwise.
func (this *CollectionPage) SetMediaType(i vocab.MediaTypePropertyInterface) {
	this.MediaType = i
}

// SetName returns the "name" property if it exists, and nil otherwise.
func (this *CollectionPage) SetName(i vocab.NamePropertyInterface) {
	this.Name = i
}

// SetNext returns the "next" property if it exists, and nil otherwise.
func (this *CollectionPage) SetNext(i vocab.NextPropertyInterface) {
	this.Next = i
}

// SetObject returns the "object" property if it exists, and nil otherwise.
func (this *CollectionPage) SetObject(i vocab.ObjectPropertyInterface) {
	this.Object = i
}

// SetPartOf returns the "partOf" property if it exists, and nil otherwise.
func (this *CollectionPage) SetPartOf(i vocab.PartOfPropertyInterface) {
	this.PartOf = i
}

// SetPrev returns the "prev" property if it exists, and nil otherwise.
func (this *CollectionPage) SetPrev(i vocab.PrevPropertyInterface) {
	this.Prev = i
}

// SetPreview returns the "preview" property if it exists, and nil otherwise.
func (this *CollectionPage) SetPreview(i vocab.PreviewPropertyInterface) {
	this.Preview = i
}

// SetPublished returns the "published" property if it exists, and nil otherwise.
func (this *CollectionPage) SetPublished(i vocab.PublishedPropertyInterface) {
	this.Published = i
}

// SetReplies returns the "replies" property if it exists, and nil otherwise.
func (this *CollectionPage) SetReplies(i vocab.RepliesPropertyInterface) {
	this.Replies = i
}

// SetShares returns the "shares" property if it exists, and nil otherwise.
func (this *CollectionPage) SetShares(i vocab.SharesPropertyInterface) {
	this.Shares = i
}

// SetStartTime returns the "startTime" property if it exists, and nil otherwise.
func (this *CollectionPage) SetStartTime(i vocab.StartTimePropertyInterface) {
	this.StartTime = i
}

// SetSummary returns the "summary" property if it exists, and nil otherwise.
func (this *CollectionPage) SetSummary(i vocab.SummaryPropertyInterface) {
	this.Summary = i
}

// SetTag returns the "tag" property if it exists, and nil otherwise.
func (this *CollectionPage) SetTag(i vocab.TagPropertyInterface) {
	this.Tag = i
}

// SetTo returns the "to" property if it exists, and nil otherwise.
func (this *CollectionPage) SetTo(i vocab.ToPropertyInterface) {
	this.To = i
}

// SetTotalItems returns the "totalItems" property if it exists, and nil otherwise.
func (this *CollectionPage) SetTotalItems(i vocab.TotalItemsPropertyInterface) {
	this.TotalItems = i
}

// SetType returns the "type" property if it exists, and nil otherwise.
func (this *CollectionPage) SetType(i vocab.TypePropertyInterface) {
	this.Type = i
}

// SetUpdated returns the "updated" property if it exists, and nil otherwise.
func (this *CollectionPage) SetUpdated(i vocab.UpdatedPropertyInterface) {
	this.Updated = i
}

// SetUrl returns the "url" property if it exists, and nil otherwise.
func (this *CollectionPage) SetUrl(i vocab.UrlPropertyInterface) {
	this.Url = i
}

// helperJSONLDContext obtains the context uris and their aliases from a property,
// if it is not nil.
func (this CollectionPage) helperJSONLDContext(i jsonldContexter, toMerge map[string]string) map[string]string {
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

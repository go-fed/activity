//
package vocab

import (
	"net/url"
	"time"
)

// AcceptType is an interface for accepting types that extend from 'Accept'.
type AcceptType interface {
	Serializer
	Deserializer
	ActorLen() (l int)
	IsActorObject(index int) (ok bool)
	GetActorObject(index int) (v ObjectType)
	AppendActorObject(v ObjectType)
	PrependActorObject(v ObjectType)
	RemoveActorObject(index int)
	IsActorLink(index int) (ok bool)
	GetActorLink(index int) (v LinkType)
	AppendActorLink(v LinkType)
	PrependActorLink(v LinkType)
	RemoveActorLink(index int)
	IsActorIRI(index int) (ok bool)
	GetActorIRI(index int) (v *url.URL)
	AppendActorIRI(v *url.URL)
	PrependActorIRI(v *url.URL)
	RemoveActorIRI(index int)
	HasUnknownActor() (ok bool)
	GetUnknownActor() (v interface{})
	SetUnknownActor(i interface{})
	ObjectLen() (l int)
	IsObject(index int) (ok bool)
	GetObject(index int) (v ObjectType)
	AppendObject(v ObjectType)
	PrependObject(v ObjectType)
	RemoveObject(index int)
	IsObjectIRI(index int) (ok bool)
	GetObjectIRI(index int) (v *url.URL)
	AppendObjectIRI(v *url.URL)
	PrependObjectIRI(v *url.URL)
	RemoveObjectIRI(index int)
	HasUnknownObject() (ok bool)
	GetUnknownObject() (v interface{})
	SetUnknownObject(i interface{})
	TargetLen() (l int)
	IsTargetObject(index int) (ok bool)
	GetTargetObject(index int) (v ObjectType)
	AppendTargetObject(v ObjectType)
	PrependTargetObject(v ObjectType)
	RemoveTargetObject(index int)
	IsTargetLink(index int) (ok bool)
	GetTargetLink(index int) (v LinkType)
	AppendTargetLink(v LinkType)
	PrependTargetLink(v LinkType)
	RemoveTargetLink(index int)
	IsTargetIRI(index int) (ok bool)
	GetTargetIRI(index int) (v *url.URL)
	AppendTargetIRI(v *url.URL)
	PrependTargetIRI(v *url.URL)
	RemoveTargetIRI(index int)
	HasUnknownTarget() (ok bool)
	GetUnknownTarget() (v interface{})
	SetUnknownTarget(i interface{})
	ResultLen() (l int)
	IsResultObject(index int) (ok bool)
	GetResultObject(index int) (v ObjectType)
	AppendResultObject(v ObjectType)
	PrependResultObject(v ObjectType)
	RemoveResultObject(index int)
	IsResultLink(index int) (ok bool)
	GetResultLink(index int) (v LinkType)
	AppendResultLink(v LinkType)
	PrependResultLink(v LinkType)
	RemoveResultLink(index int)
	IsResultIRI(index int) (ok bool)
	GetResultIRI(index int) (v *url.URL)
	AppendResultIRI(v *url.URL)
	PrependResultIRI(v *url.URL)
	RemoveResultIRI(index int)
	HasUnknownResult() (ok bool)
	GetUnknownResult() (v interface{})
	SetUnknownResult(i interface{})
	OriginLen() (l int)
	IsOriginObject(index int) (ok bool)
	GetOriginObject(index int) (v ObjectType)
	AppendOriginObject(v ObjectType)
	PrependOriginObject(v ObjectType)
	RemoveOriginObject(index int)
	IsOriginLink(index int) (ok bool)
	GetOriginLink(index int) (v LinkType)
	AppendOriginLink(v LinkType)
	PrependOriginLink(v LinkType)
	RemoveOriginLink(index int)
	IsOriginIRI(index int) (ok bool)
	GetOriginIRI(index int) (v *url.URL)
	AppendOriginIRI(v *url.URL)
	PrependOriginIRI(v *url.URL)
	RemoveOriginIRI(index int)
	HasUnknownOrigin() (ok bool)
	GetUnknownOrigin() (v interface{})
	SetUnknownOrigin(i interface{})
	InstrumentLen() (l int)
	IsInstrumentObject(index int) (ok bool)
	GetInstrumentObject(index int) (v ObjectType)
	AppendInstrumentObject(v ObjectType)
	PrependInstrumentObject(v ObjectType)
	RemoveInstrumentObject(index int)
	IsInstrumentLink(index int) (ok bool)
	GetInstrumentLink(index int) (v LinkType)
	AppendInstrumentLink(v LinkType)
	PrependInstrumentLink(v LinkType)
	RemoveInstrumentLink(index int)
	IsInstrumentIRI(index int) (ok bool)
	GetInstrumentIRI(index int) (v *url.URL)
	AppendInstrumentIRI(v *url.URL)
	PrependInstrumentIRI(v *url.URL)
	RemoveInstrumentIRI(index int)
	HasUnknownInstrument() (ok bool)
	GetUnknownInstrument() (v interface{})
	SetUnknownInstrument(i interface{})
	IsAltitude() (ok bool)
	GetAltitude() (v float64)
	SetAltitude(v float64)
	IsAltitudeIRI() (ok bool)
	GetAltitudeIRI() (v *url.URL)
	SetAltitudeIRI(v *url.URL)
	HasUnknownAltitude() (ok bool)
	GetUnknownAltitude() (v interface{})
	SetUnknownAltitude(i interface{})
	AttachmentLen() (l int)
	IsAttachmentObject(index int) (ok bool)
	GetAttachmentObject(index int) (v ObjectType)
	AppendAttachmentObject(v ObjectType)
	PrependAttachmentObject(v ObjectType)
	RemoveAttachmentObject(index int)
	IsAttachmentLink(index int) (ok bool)
	GetAttachmentLink(index int) (v LinkType)
	AppendAttachmentLink(v LinkType)
	PrependAttachmentLink(v LinkType)
	RemoveAttachmentLink(index int)
	IsAttachmentIRI(index int) (ok bool)
	GetAttachmentIRI(index int) (v *url.URL)
	AppendAttachmentIRI(v *url.URL)
	PrependAttachmentIRI(v *url.URL)
	RemoveAttachmentIRI(index int)
	HasUnknownAttachment() (ok bool)
	GetUnknownAttachment() (v interface{})
	SetUnknownAttachment(i interface{})
	AttributedToLen() (l int)
	IsAttributedToObject(index int) (ok bool)
	GetAttributedToObject(index int) (v ObjectType)
	AppendAttributedToObject(v ObjectType)
	PrependAttributedToObject(v ObjectType)
	RemoveAttributedToObject(index int)
	IsAttributedToLink(index int) (ok bool)
	GetAttributedToLink(index int) (v LinkType)
	AppendAttributedToLink(v LinkType)
	PrependAttributedToLink(v LinkType)
	RemoveAttributedToLink(index int)
	IsAttributedToIRI(index int) (ok bool)
	GetAttributedToIRI(index int) (v *url.URL)
	AppendAttributedToIRI(v *url.URL)
	PrependAttributedToIRI(v *url.URL)
	RemoveAttributedToIRI(index int)
	HasUnknownAttributedTo() (ok bool)
	GetUnknownAttributedTo() (v interface{})
	SetUnknownAttributedTo(i interface{})
	AudienceLen() (l int)
	IsAudienceObject(index int) (ok bool)
	GetAudienceObject(index int) (v ObjectType)
	AppendAudienceObject(v ObjectType)
	PrependAudienceObject(v ObjectType)
	RemoveAudienceObject(index int)
	IsAudienceLink(index int) (ok bool)
	GetAudienceLink(index int) (v LinkType)
	AppendAudienceLink(v LinkType)
	PrependAudienceLink(v LinkType)
	RemoveAudienceLink(index int)
	IsAudienceIRI(index int) (ok bool)
	GetAudienceIRI(index int) (v *url.URL)
	AppendAudienceIRI(v *url.URL)
	PrependAudienceIRI(v *url.URL)
	RemoveAudienceIRI(index int)
	HasUnknownAudience() (ok bool)
	GetUnknownAudience() (v interface{})
	SetUnknownAudience(i interface{})
	ContentLen() (l int)
	IsContentString(index int) (ok bool)
	GetContentString(index int) (v string)
	AppendContentString(v string)
	PrependContentString(v string)
	RemoveContentString(index int)
	IsContentLangString(index int) (ok bool)
	GetContentLangString(index int) (v string)
	AppendContentLangString(v string)
	PrependContentLangString(v string)
	RemoveContentLangString(index int)
	IsContentIRI(index int) (ok bool)
	GetContentIRI(index int) (v *url.URL)
	AppendContentIRI(v *url.URL)
	PrependContentIRI(v *url.URL)
	RemoveContentIRI(index int)
	HasUnknownContent() (ok bool)
	GetUnknownContent() (v interface{})
	SetUnknownContent(i interface{})
	ContentMapLanguages() (l []string)
	GetContentMap(l string) (v string)
	SetContentMap(l string, v string)
	ContextLen() (l int)
	IsContextObject(index int) (ok bool)
	GetContextObject(index int) (v ObjectType)
	AppendContextObject(v ObjectType)
	PrependContextObject(v ObjectType)
	RemoveContextObject(index int)
	IsContextLink(index int) (ok bool)
	GetContextLink(index int) (v LinkType)
	AppendContextLink(v LinkType)
	PrependContextLink(v LinkType)
	RemoveContextLink(index int)
	IsContextIRI(index int) (ok bool)
	GetContextIRI(index int) (v *url.URL)
	AppendContextIRI(v *url.URL)
	PrependContextIRI(v *url.URL)
	RemoveContextIRI(index int)
	HasUnknownContext() (ok bool)
	GetUnknownContext() (v interface{})
	SetUnknownContext(i interface{})
	NameLen() (l int)
	IsNameString(index int) (ok bool)
	GetNameString(index int) (v string)
	AppendNameString(v string)
	PrependNameString(v string)
	RemoveNameString(index int)
	IsNameLangString(index int) (ok bool)
	GetNameLangString(index int) (v string)
	AppendNameLangString(v string)
	PrependNameLangString(v string)
	RemoveNameLangString(index int)
	IsNameIRI(index int) (ok bool)
	GetNameIRI(index int) (v *url.URL)
	AppendNameIRI(v *url.URL)
	PrependNameIRI(v *url.URL)
	RemoveNameIRI(index int)
	HasUnknownName() (ok bool)
	GetUnknownName() (v interface{})
	SetUnknownName(i interface{})
	NameMapLanguages() (l []string)
	GetNameMap(l string) (v string)
	SetNameMap(l string, v string)
	IsEndTime() (ok bool)
	GetEndTime() (v time.Time)
	SetEndTime(v time.Time)
	IsEndTimeIRI() (ok bool)
	GetEndTimeIRI() (v *url.URL)
	SetEndTimeIRI(v *url.URL)
	HasUnknownEndTime() (ok bool)
	GetUnknownEndTime() (v interface{})
	SetUnknownEndTime(i interface{})
	GeneratorLen() (l int)
	IsGeneratorObject(index int) (ok bool)
	GetGeneratorObject(index int) (v ObjectType)
	AppendGeneratorObject(v ObjectType)
	PrependGeneratorObject(v ObjectType)
	RemoveGeneratorObject(index int)
	IsGeneratorLink(index int) (ok bool)
	GetGeneratorLink(index int) (v LinkType)
	AppendGeneratorLink(v LinkType)
	PrependGeneratorLink(v LinkType)
	RemoveGeneratorLink(index int)
	IsGeneratorIRI(index int) (ok bool)
	GetGeneratorIRI(index int) (v *url.URL)
	AppendGeneratorIRI(v *url.URL)
	PrependGeneratorIRI(v *url.URL)
	RemoveGeneratorIRI(index int)
	HasUnknownGenerator() (ok bool)
	GetUnknownGenerator() (v interface{})
	SetUnknownGenerator(i interface{})
	IconLen() (l int)
	IsIconImage(index int) (ok bool)
	GetIconImage(index int) (v ImageType)
	AppendIconImage(v ImageType)
	PrependIconImage(v ImageType)
	RemoveIconImage(index int)
	IsIconLink(index int) (ok bool)
	GetIconLink(index int) (v LinkType)
	AppendIconLink(v LinkType)
	PrependIconLink(v LinkType)
	RemoveIconLink(index int)
	IsIconIRI(index int) (ok bool)
	GetIconIRI(index int) (v *url.URL)
	AppendIconIRI(v *url.URL)
	PrependIconIRI(v *url.URL)
	RemoveIconIRI(index int)
	HasUnknownIcon() (ok bool)
	GetUnknownIcon() (v interface{})
	SetUnknownIcon(i interface{})
	HasId() (ok bool)
	GetId() (v *url.URL)
	SetId(v *url.URL)
	HasUnknownId() (ok bool)
	GetUnknownId() (v interface{})
	SetUnknownId(i interface{})
	ImageLen() (l int)
	IsImageImage(index int) (ok bool)
	GetImageImage(index int) (v ImageType)
	AppendImageImage(v ImageType)
	PrependImageImage(v ImageType)
	RemoveImageImage(index int)
	IsImageLink(index int) (ok bool)
	GetImageLink(index int) (v LinkType)
	AppendImageLink(v LinkType)
	PrependImageLink(v LinkType)
	RemoveImageLink(index int)
	IsImageIRI(index int) (ok bool)
	GetImageIRI(index int) (v *url.URL)
	AppendImageIRI(v *url.URL)
	PrependImageIRI(v *url.URL)
	RemoveImageIRI(index int)
	HasUnknownImage() (ok bool)
	GetUnknownImage() (v interface{})
	SetUnknownImage(i interface{})
	InReplyToLen() (l int)
	IsInReplyToObject(index int) (ok bool)
	GetInReplyToObject(index int) (v ObjectType)
	AppendInReplyToObject(v ObjectType)
	PrependInReplyToObject(v ObjectType)
	RemoveInReplyToObject(index int)
	IsInReplyToLink(index int) (ok bool)
	GetInReplyToLink(index int) (v LinkType)
	AppendInReplyToLink(v LinkType)
	PrependInReplyToLink(v LinkType)
	RemoveInReplyToLink(index int)
	IsInReplyToIRI(index int) (ok bool)
	GetInReplyToIRI(index int) (v *url.URL)
	AppendInReplyToIRI(v *url.URL)
	PrependInReplyToIRI(v *url.URL)
	RemoveInReplyToIRI(index int)
	HasUnknownInReplyTo() (ok bool)
	GetUnknownInReplyTo() (v interface{})
	SetUnknownInReplyTo(i interface{})
	LocationLen() (l int)
	IsLocationObject(index int) (ok bool)
	GetLocationObject(index int) (v ObjectType)
	AppendLocationObject(v ObjectType)
	PrependLocationObject(v ObjectType)
	RemoveLocationObject(index int)
	IsLocationLink(index int) (ok bool)
	GetLocationLink(index int) (v LinkType)
	AppendLocationLink(v LinkType)
	PrependLocationLink(v LinkType)
	RemoveLocationLink(index int)
	IsLocationIRI(index int) (ok bool)
	GetLocationIRI(index int) (v *url.URL)
	AppendLocationIRI(v *url.URL)
	PrependLocationIRI(v *url.URL)
	RemoveLocationIRI(index int)
	HasUnknownLocation() (ok bool)
	GetUnknownLocation() (v interface{})
	SetUnknownLocation(i interface{})
	PreviewLen() (l int)
	IsPreviewObject(index int) (ok bool)
	GetPreviewObject(index int) (v ObjectType)
	AppendPreviewObject(v ObjectType)
	PrependPreviewObject(v ObjectType)
	RemovePreviewObject(index int)
	IsPreviewLink(index int) (ok bool)
	GetPreviewLink(index int) (v LinkType)
	AppendPreviewLink(v LinkType)
	PrependPreviewLink(v LinkType)
	RemovePreviewLink(index int)
	IsPreviewIRI(index int) (ok bool)
	GetPreviewIRI(index int) (v *url.URL)
	AppendPreviewIRI(v *url.URL)
	PrependPreviewIRI(v *url.URL)
	RemovePreviewIRI(index int)
	HasUnknownPreview() (ok bool)
	GetUnknownPreview() (v interface{})
	SetUnknownPreview(i interface{})
	IsPublished() (ok bool)
	GetPublished() (v time.Time)
	SetPublished(v time.Time)
	IsPublishedIRI() (ok bool)
	GetPublishedIRI() (v *url.URL)
	SetPublishedIRI(v *url.URL)
	HasUnknownPublished() (ok bool)
	GetUnknownPublished() (v interface{})
	SetUnknownPublished(i interface{})
	IsReplies() (ok bool)
	GetReplies() (v CollectionType)
	SetReplies(v CollectionType)
	IsRepliesIRI() (ok bool)
	GetRepliesIRI() (v *url.URL)
	SetRepliesIRI(v *url.URL)
	HasUnknownReplies() (ok bool)
	GetUnknownReplies() (v interface{})
	SetUnknownReplies(i interface{})
	IsStartTime() (ok bool)
	GetStartTime() (v time.Time)
	SetStartTime(v time.Time)
	IsStartTimeIRI() (ok bool)
	GetStartTimeIRI() (v *url.URL)
	SetStartTimeIRI(v *url.URL)
	HasUnknownStartTime() (ok bool)
	GetUnknownStartTime() (v interface{})
	SetUnknownStartTime(i interface{})
	SummaryLen() (l int)
	IsSummaryString(index int) (ok bool)
	GetSummaryString(index int) (v string)
	AppendSummaryString(v string)
	PrependSummaryString(v string)
	RemoveSummaryString(index int)
	IsSummaryLangString(index int) (ok bool)
	GetSummaryLangString(index int) (v string)
	AppendSummaryLangString(v string)
	PrependSummaryLangString(v string)
	RemoveSummaryLangString(index int)
	IsSummaryIRI(index int) (ok bool)
	GetSummaryIRI(index int) (v *url.URL)
	AppendSummaryIRI(v *url.URL)
	PrependSummaryIRI(v *url.URL)
	RemoveSummaryIRI(index int)
	HasUnknownSummary() (ok bool)
	GetUnknownSummary() (v interface{})
	SetUnknownSummary(i interface{})
	SummaryMapLanguages() (l []string)
	GetSummaryMap(l string) (v string)
	SetSummaryMap(l string, v string)
	TagLen() (l int)
	IsTagObject(index int) (ok bool)
	GetTagObject(index int) (v ObjectType)
	AppendTagObject(v ObjectType)
	PrependTagObject(v ObjectType)
	RemoveTagObject(index int)
	IsTagLink(index int) (ok bool)
	GetTagLink(index int) (v LinkType)
	AppendTagLink(v LinkType)
	PrependTagLink(v LinkType)
	RemoveTagLink(index int)
	IsTagIRI(index int) (ok bool)
	GetTagIRI(index int) (v *url.URL)
	AppendTagIRI(v *url.URL)
	PrependTagIRI(v *url.URL)
	RemoveTagIRI(index int)
	HasUnknownTag() (ok bool)
	GetUnknownTag() (v interface{})
	SetUnknownTag(i interface{})
	TypeLen() (l int)
	GetType(index int) (v interface{})
	AppendType(v interface{})
	PrependType(v interface{})
	RemoveType(index int)
	IsUpdated() (ok bool)
	GetUpdated() (v time.Time)
	SetUpdated(v time.Time)
	IsUpdatedIRI() (ok bool)
	GetUpdatedIRI() (v *url.URL)
	SetUpdatedIRI(v *url.URL)
	HasUnknownUpdated() (ok bool)
	GetUnknownUpdated() (v interface{})
	SetUnknownUpdated(i interface{})
	UrlLen() (l int)
	IsUrlAnyURI(index int) (ok bool)
	GetUrlAnyURI(index int) (v *url.URL)
	AppendUrlAnyURI(v *url.URL)
	PrependUrlAnyURI(v *url.URL)
	RemoveUrlAnyURI(index int)
	IsUrlLink(index int) (ok bool)
	GetUrlLink(index int) (v LinkType)
	AppendUrlLink(v LinkType)
	PrependUrlLink(v LinkType)
	RemoveUrlLink(index int)
	HasUnknownUrl() (ok bool)
	GetUnknownUrl() (v interface{})
	SetUnknownUrl(i interface{})
	ToLen() (l int)
	IsToObject(index int) (ok bool)
	GetToObject(index int) (v ObjectType)
	AppendToObject(v ObjectType)
	PrependToObject(v ObjectType)
	RemoveToObject(index int)
	IsToLink(index int) (ok bool)
	GetToLink(index int) (v LinkType)
	AppendToLink(v LinkType)
	PrependToLink(v LinkType)
	RemoveToLink(index int)
	IsToIRI(index int) (ok bool)
	GetToIRI(index int) (v *url.URL)
	AppendToIRI(v *url.URL)
	PrependToIRI(v *url.URL)
	RemoveToIRI(index int)
	HasUnknownTo() (ok bool)
	GetUnknownTo() (v interface{})
	SetUnknownTo(i interface{})
	BtoLen() (l int)
	IsBtoObject(index int) (ok bool)
	GetBtoObject(index int) (v ObjectType)
	AppendBtoObject(v ObjectType)
	PrependBtoObject(v ObjectType)
	RemoveBtoObject(index int)
	IsBtoLink(index int) (ok bool)
	GetBtoLink(index int) (v LinkType)
	AppendBtoLink(v LinkType)
	PrependBtoLink(v LinkType)
	RemoveBtoLink(index int)
	IsBtoIRI(index int) (ok bool)
	GetBtoIRI(index int) (v *url.URL)
	AppendBtoIRI(v *url.URL)
	PrependBtoIRI(v *url.URL)
	RemoveBtoIRI(index int)
	HasUnknownBto() (ok bool)
	GetUnknownBto() (v interface{})
	SetUnknownBto(i interface{})
	CcLen() (l int)
	IsCcObject(index int) (ok bool)
	GetCcObject(index int) (v ObjectType)
	AppendCcObject(v ObjectType)
	PrependCcObject(v ObjectType)
	RemoveCcObject(index int)
	IsCcLink(index int) (ok bool)
	GetCcLink(index int) (v LinkType)
	AppendCcLink(v LinkType)
	PrependCcLink(v LinkType)
	RemoveCcLink(index int)
	IsCcIRI(index int) (ok bool)
	GetCcIRI(index int) (v *url.URL)
	AppendCcIRI(v *url.URL)
	PrependCcIRI(v *url.URL)
	RemoveCcIRI(index int)
	HasUnknownCc() (ok bool)
	GetUnknownCc() (v interface{})
	SetUnknownCc(i interface{})
	BccLen() (l int)
	IsBccObject(index int) (ok bool)
	GetBccObject(index int) (v ObjectType)
	AppendBccObject(v ObjectType)
	PrependBccObject(v ObjectType)
	RemoveBccObject(index int)
	IsBccLink(index int) (ok bool)
	GetBccLink(index int) (v LinkType)
	AppendBccLink(v LinkType)
	PrependBccLink(v LinkType)
	RemoveBccLink(index int)
	IsBccIRI(index int) (ok bool)
	GetBccIRI(index int) (v *url.URL)
	AppendBccIRI(v *url.URL)
	PrependBccIRI(v *url.URL)
	RemoveBccIRI(index int)
	HasUnknownBcc() (ok bool)
	GetUnknownBcc() (v interface{})
	SetUnknownBcc(i interface{})
	IsMediaType() (ok bool)
	GetMediaType() (v string)
	SetMediaType(v string)
	IsMediaTypeIRI() (ok bool)
	GetMediaTypeIRI() (v *url.URL)
	SetMediaTypeIRI(v *url.URL)
	HasUnknownMediaType() (ok bool)
	GetUnknownMediaType() (v interface{})
	SetUnknownMediaType(i interface{})
	IsDuration() (ok bool)
	GetDuration() (v time.Duration)
	SetDuration(v time.Duration)
	IsDurationIRI() (ok bool)
	GetDurationIRI() (v *url.URL)
	SetDurationIRI(v *url.URL)
	HasUnknownDuration() (ok bool)
	GetUnknownDuration() (v interface{})
	SetUnknownDuration(i interface{})
	IsSource() (ok bool)
	GetSource() (v ObjectType)
	SetSource(v ObjectType)
	IsSourceIRI() (ok bool)
	GetSourceIRI() (v *url.URL)
	SetSourceIRI(v *url.URL)
	HasUnknownSource() (ok bool)
	GetUnknownSource() (v interface{})
	SetUnknownSource(i interface{})
	IsInboxOrderedCollection() (ok bool)
	GetInboxOrderedCollection() (v OrderedCollectionType)
	SetInboxOrderedCollection(v OrderedCollectionType)
	IsInboxAnyURI() (ok bool)
	GetInboxAnyURI() (v *url.URL)
	SetInboxAnyURI(v *url.URL)
	HasUnknownInbox() (ok bool)
	GetUnknownInbox() (v interface{})
	SetUnknownInbox(i interface{})
	IsOutboxOrderedCollection() (ok bool)
	GetOutboxOrderedCollection() (v OrderedCollectionType)
	SetOutboxOrderedCollection(v OrderedCollectionType)
	IsOutboxAnyURI() (ok bool)
	GetOutboxAnyURI() (v *url.URL)
	SetOutboxAnyURI(v *url.URL)
	HasUnknownOutbox() (ok bool)
	GetUnknownOutbox() (v interface{})
	SetUnknownOutbox(i interface{})
	IsFollowingCollection() (ok bool)
	GetFollowingCollection() (v CollectionType)
	SetFollowingCollection(v CollectionType)
	IsFollowingOrderedCollection() (ok bool)
	GetFollowingOrderedCollection() (v OrderedCollectionType)
	SetFollowingOrderedCollection(v OrderedCollectionType)
	IsFollowingAnyURI() (ok bool)
	GetFollowingAnyURI() (v *url.URL)
	SetFollowingAnyURI(v *url.URL)
	HasUnknownFollowing() (ok bool)
	GetUnknownFollowing() (v interface{})
	SetUnknownFollowing(i interface{})
	IsFollowersCollection() (ok bool)
	GetFollowersCollection() (v CollectionType)
	SetFollowersCollection(v CollectionType)
	IsFollowersOrderedCollection() (ok bool)
	GetFollowersOrderedCollection() (v OrderedCollectionType)
	SetFollowersOrderedCollection(v OrderedCollectionType)
	IsFollowersAnyURI() (ok bool)
	GetFollowersAnyURI() (v *url.URL)
	SetFollowersAnyURI(v *url.URL)
	HasUnknownFollowers() (ok bool)
	GetUnknownFollowers() (v interface{})
	SetUnknownFollowers(i interface{})
	IsLikedCollection() (ok bool)
	GetLikedCollection() (v CollectionType)
	SetLikedCollection(v CollectionType)
	IsLikedOrderedCollection() (ok bool)
	GetLikedOrderedCollection() (v OrderedCollectionType)
	SetLikedOrderedCollection(v OrderedCollectionType)
	IsLikedAnyURI() (ok bool)
	GetLikedAnyURI() (v *url.URL)
	SetLikedAnyURI(v *url.URL)
	HasUnknownLiked() (ok bool)
	GetUnknownLiked() (v interface{})
	SetUnknownLiked(i interface{})
	IsLikesCollection() (ok bool)
	GetLikesCollection() (v CollectionType)
	SetLikesCollection(v CollectionType)
	IsLikesOrderedCollection() (ok bool)
	GetLikesOrderedCollection() (v OrderedCollectionType)
	SetLikesOrderedCollection(v OrderedCollectionType)
	IsLikesAnyURI() (ok bool)
	GetLikesAnyURI() (v *url.URL)
	SetLikesAnyURI(v *url.URL)
	HasUnknownLikes() (ok bool)
	GetUnknownLikes() (v interface{})
	SetUnknownLikes(i interface{})
	StreamsLen() (l int)
	GetStreams(index int) (v *url.URL)
	AppendStreams(v *url.URL)
	PrependStreams(v *url.URL)
	RemoveStreams(index int)
	HasUnknownStreams() (ok bool)
	GetUnknownStreams() (v interface{})
	SetUnknownStreams(i interface{})
	IsPreferredUsername() (ok bool)
	GetPreferredUsername() (v string)
	SetPreferredUsername(v string)
	IsPreferredUsernameIRI() (ok bool)
	GetPreferredUsernameIRI() (v *url.URL)
	SetPreferredUsernameIRI(v *url.URL)
	HasUnknownPreferredUsername() (ok bool)
	GetUnknownPreferredUsername() (v interface{})
	SetUnknownPreferredUsername(i interface{})
	PreferredUsernameMapLanguages() (l []string)
	GetPreferredUsernameMap(l string) (v string)
	SetPreferredUsernameMap(l string, v string)
	IsEndpoints() (ok bool)
	GetEndpoints() (v ObjectType)
	SetEndpoints(v ObjectType)
	IsEndpointsIRI() (ok bool)
	GetEndpointsIRI() (v *url.URL)
	SetEndpointsIRI(v *url.URL)
	HasUnknownEndpoints() (ok bool)
	GetUnknownEndpoints() (v interface{})
	SetUnknownEndpoints(i interface{})
	HasProxyUrl() (ok bool)
	GetProxyUrl() (v *url.URL)
	SetProxyUrl(v *url.URL)
	HasUnknownProxyUrl() (ok bool)
	GetUnknownProxyUrl() (v interface{})
	SetUnknownProxyUrl(i interface{})
	HasOauthAuthorizationEndpoint() (ok bool)
	GetOauthAuthorizationEndpoint() (v *url.URL)
	SetOauthAuthorizationEndpoint(v *url.URL)
	HasUnknownOauthAuthorizationEndpoint() (ok bool)
	GetUnknownOauthAuthorizationEndpoint() (v interface{})
	SetUnknownOauthAuthorizationEndpoint(i interface{})
	HasOauthTokenEndpoint() (ok bool)
	GetOauthTokenEndpoint() (v *url.URL)
	SetOauthTokenEndpoint(v *url.URL)
	HasUnknownOauthTokenEndpoint() (ok bool)
	GetUnknownOauthTokenEndpoint() (v interface{})
	SetUnknownOauthTokenEndpoint(i interface{})
	HasProvideClientKey() (ok bool)
	GetProvideClientKey() (v *url.URL)
	SetProvideClientKey(v *url.URL)
	HasUnknownProvideClientKey() (ok bool)
	GetUnknownProvideClientKey() (v interface{})
	SetUnknownProvideClientKey(i interface{})
	HasSignClientKey() (ok bool)
	GetSignClientKey() (v *url.URL)
	SetSignClientKey(v *url.URL)
	HasUnknownSignClientKey() (ok bool)
	GetUnknownSignClientKey() (v interface{})
	SetUnknownSignClientKey(i interface{})
	HasSharedInbox() (ok bool)
	GetSharedInbox() (v *url.URL)
	SetSharedInbox(v *url.URL)
	HasUnknownSharedInbox() (ok bool)
	GetUnknownSharedInbox() (v interface{})
	SetUnknownSharedInbox(i interface{})
	IsSharesCollection() (ok bool)
	GetSharesCollection() (v CollectionType)
	SetSharesCollection(v CollectionType)
	IsSharesOrderedCollection() (ok bool)
	GetSharesOrderedCollection() (v OrderedCollectionType)
	SetSharesOrderedCollection(v OrderedCollectionType)
	IsSharesAnyURI() (ok bool)
	GetSharesAnyURI() (v *url.URL)
	SetSharesAnyURI(v *url.URL)
	HasUnknownShares() (ok bool)
	GetUnknownShares() (v interface{})
	SetUnknownShares(i interface{})
	IsPublic() (b bool)
}

// Indicates that the actor accepts the object. The target property can be used in certain circumstances to indicate the context into which the object has been accepted.
type Accept struct {
	// An unknown value.
	unknown_ map[string]interface{}
	// The 'actor' value could have multiple types and values
	actor []*actorIntermediateType
	// The 'object' value could have multiple types and values
	object []*objectIntermediateType
	// The 'target' value could have multiple types and values
	target []*targetIntermediateType
	// The 'result' value could have multiple types and values
	result []*resultIntermediateType
	// The 'origin' value could have multiple types and values
	origin []*originIntermediateType
	// The 'instrument' value could have multiple types and values
	instrument []*instrumentIntermediateType
	// The functional 'altitude' value could have multiple types, but only a single value
	altitude *altitudeIntermediateType
	// The 'attachment' value could have multiple types and values
	attachment []*attachmentIntermediateType
	// The 'attributedTo' value could have multiple types and values
	attributedTo []*attributedToIntermediateType
	// The 'audience' value could have multiple types and values
	audience []*audienceIntermediateType
	// The 'content' value could have multiple types and values
	content []*contentIntermediateType
	// The 'contentMap' value holds language-specific values for property 'content'
	contentMap map[string]string
	// The 'context' value could have multiple types and values
	context []*contextIntermediateType
	// The 'name' value could have multiple types and values
	name []*nameIntermediateType
	// The 'nameMap' value holds language-specific values for property 'name'
	nameMap map[string]string
	// The functional 'endTime' value could have multiple types, but only a single value
	endTime *endTimeIntermediateType
	// The 'generator' value could have multiple types and values
	generator []*generatorIntermediateType
	// The 'icon' value could have multiple types and values
	icon []*iconIntermediateType
	// The functional 'id' value holds a single type and a single value
	id *url.URL
	// The 'image' value could have multiple types and values
	image []*imageIntermediateType
	// The 'inReplyTo' value could have multiple types and values
	inReplyTo []*inReplyToIntermediateType
	// The 'location' value could have multiple types and values
	location []*locationIntermediateType
	// The 'preview' value could have multiple types and values
	preview []*previewIntermediateType
	// The functional 'published' value could have multiple types, but only a single value
	published *publishedIntermediateType
	// The functional 'replies' value could have multiple types, but only a single value
	replies *repliesIntermediateType
	// The functional 'startTime' value could have multiple types, but only a single value
	startTime *startTimeIntermediateType
	// The 'summary' value could have multiple types and values
	summary []*summaryIntermediateType
	// The 'summaryMap' value holds language-specific values for property 'summary'
	summaryMap map[string]string
	// The 'tag' value could have multiple types and values
	tag []*tagIntermediateType
	// The 'type' value can hold any type and any number of values
	typeName []interface{}
	// The functional 'updated' value could have multiple types, but only a single value
	updated *updatedIntermediateType
	// The 'url' value could have multiple types and values
	url []*urlIntermediateType
	// The 'to' value could have multiple types and values
	to []*toIntermediateType
	// The 'bto' value could have multiple types and values
	bto []*btoIntermediateType
	// The 'cc' value could have multiple types and values
	cc []*ccIntermediateType
	// The 'bcc' value could have multiple types and values
	bcc []*bccIntermediateType
	// The functional 'mediaType' value could have multiple types, but only a single value
	mediaType *mediaTypeIntermediateType
	// The functional 'duration' value could have multiple types, but only a single value
	duration *durationIntermediateType
	// The functional 'source' value could have multiple types, but only a single value
	source *sourceIntermediateType
	// The functional 'inbox' value could have multiple types, but only a single value
	inbox *inboxIntermediateType
	// The functional 'outbox' value could have multiple types, but only a single value
	outbox *outboxIntermediateType
	// The functional 'following' value could have multiple types, but only a single value
	following *followingIntermediateType
	// The functional 'followers' value could have multiple types, but only a single value
	followers *followersIntermediateType
	// The functional 'liked' value could have multiple types, but only a single value
	liked *likedIntermediateType
	// The functional 'likes' value could have multiple types, but only a single value
	likes *likesIntermediateType
	// The 'streams' value holds a single type and any number of values
	streams []*url.URL
	// The functional 'preferredUsername' value could have multiple types, but only a single value
	preferredUsername *preferredUsernameIntermediateType
	// The 'preferredUsernameMap' value holds language-specific values for property 'preferredUsername'
	preferredUsernameMap map[string]string
	// The functional 'endpoints' value could have multiple types, but only a single value
	endpoints *endpointsIntermediateType
	// The functional 'proxyUrl' value holds a single type and a single value
	proxyUrl *url.URL
	// The functional 'oauthAuthorizationEndpoint' value holds a single type and a single value
	oauthAuthorizationEndpoint *url.URL
	// The functional 'oauthTokenEndpoint' value holds a single type and a single value
	oauthTokenEndpoint *url.URL
	// The functional 'provideClientKey' value holds a single type and a single value
	provideClientKey *url.URL
	// The functional 'signClientKey' value holds a single type and a single value
	signClientKey *url.URL
	// The functional 'sharedInbox' value holds a single type and a single value
	sharedInbox *url.URL
	// The functional 'shares' value could have multiple types, but only a single value
	shares *sharesIntermediateType
}

// ActorLen determines the number of elements able to be used for the IsActorObject, GetActorObject, and RemoveActorObject functions
func (t *Accept) ActorLen() (l int) {
	return len(t.actor)

}

// IsActorObject determines whether the call to GetActorObject is safe for the specified index
func (t *Accept) IsActorObject(index int) (ok bool) {
	return t.actor[index].Object != nil

}

// GetActorObject returns the value safely if IsActorObject returned true for the specified index
func (t *Accept) GetActorObject(index int) (v ObjectType) {
	return t.actor[index].Object

}

// AppendActorObject adds to the back of actor a ObjectType type
func (t *Accept) AppendActorObject(v ObjectType) {
	t.actor = append(t.actor, &actorIntermediateType{Object: v})

}

// PrependActorObject adds to the front of actor a ObjectType type
func (t *Accept) PrependActorObject(v ObjectType) {
	t.actor = append([]*actorIntermediateType{&actorIntermediateType{Object: v}}, t.actor...)

}

// RemoveActorObject deletes the value from the specified index
func (t *Accept) RemoveActorObject(index int) {
	copy(t.actor[index:], t.actor[index+1:])
	t.actor[len(t.actor)-1] = nil
	t.actor = t.actor[:len(t.actor)-1]

}

// IsActorLink determines whether the call to GetActorLink is safe for the specified index
func (t *Accept) IsActorLink(index int) (ok bool) {
	return t.actor[index].Link != nil

}

// GetActorLink returns the value safely if IsActorLink returned true for the specified index
func (t *Accept) GetActorLink(index int) (v LinkType) {
	return t.actor[index].Link

}

// AppendActorLink adds to the back of actor a LinkType type
func (t *Accept) AppendActorLink(v LinkType) {
	t.actor = append(t.actor, &actorIntermediateType{Link: v})

}

// PrependActorLink adds to the front of actor a LinkType type
func (t *Accept) PrependActorLink(v LinkType) {
	t.actor = append([]*actorIntermediateType{&actorIntermediateType{Link: v}}, t.actor...)

}

// RemoveActorLink deletes the value from the specified index
func (t *Accept) RemoveActorLink(index int) {
	copy(t.actor[index:], t.actor[index+1:])
	t.actor[len(t.actor)-1] = nil
	t.actor = t.actor[:len(t.actor)-1]

}

// IsActorIRI determines whether the call to GetActorIRI is safe for the specified index
func (t *Accept) IsActorIRI(index int) (ok bool) {
	return t.actor[index].IRI != nil

}

// GetActorIRI returns the value safely if IsActorIRI returned true for the specified index
func (t *Accept) GetActorIRI(index int) (v *url.URL) {
	return t.actor[index].IRI

}

// AppendActorIRI adds to the back of actor a *url.URL type
func (t *Accept) AppendActorIRI(v *url.URL) {
	t.actor = append(t.actor, &actorIntermediateType{IRI: v})

}

// PrependActorIRI adds to the front of actor a *url.URL type
func (t *Accept) PrependActorIRI(v *url.URL) {
	t.actor = append([]*actorIntermediateType{&actorIntermediateType{IRI: v}}, t.actor...)

}

// RemoveActorIRI deletes the value from the specified index
func (t *Accept) RemoveActorIRI(index int) {
	copy(t.actor[index:], t.actor[index+1:])
	t.actor[len(t.actor)-1] = nil
	t.actor = t.actor[:len(t.actor)-1]

}

// HasUnknownActor determines whether the call to GetUnknownActor is safe
func (t *Accept) HasUnknownActor() (ok bool) {
	return t.actor != nil && t.actor[0].unknown_ != nil

}

// GetUnknownActor returns the unknown value for actor
func (t *Accept) GetUnknownActor() (v interface{}) {
	return t.actor[0].unknown_

}

// SetUnknownActor sets the unknown value of actor
func (t *Accept) SetUnknownActor(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &actorIntermediateType{}
	tmp.unknown_ = i
	t.actor = append(t.actor, tmp)

}

// ObjectLen determines the number of elements able to be used for the IsObject, GetObject, and RemoveObject functions
func (t *Accept) ObjectLen() (l int) {
	return len(t.object)

}

// IsObject determines whether the call to GetObject is safe for the specified index
func (t *Accept) IsObject(index int) (ok bool) {
	return t.object[index].Object != nil

}

// GetObject returns the value safely if IsObject returned true for the specified index
func (t *Accept) GetObject(index int) (v ObjectType) {
	return t.object[index].Object

}

// AppendObject adds to the back of object a ObjectType type
func (t *Accept) AppendObject(v ObjectType) {
	t.object = append(t.object, &objectIntermediateType{Object: v})

}

// PrependObject adds to the front of object a ObjectType type
func (t *Accept) PrependObject(v ObjectType) {
	t.object = append([]*objectIntermediateType{&objectIntermediateType{Object: v}}, t.object...)

}

// RemoveObject deletes the value from the specified index
func (t *Accept) RemoveObject(index int) {
	copy(t.object[index:], t.object[index+1:])
	t.object[len(t.object)-1] = nil
	t.object = t.object[:len(t.object)-1]

}

// IsObjectIRI determines whether the call to GetObjectIRI is safe for the specified index
func (t *Accept) IsObjectIRI(index int) (ok bool) {
	return t.object[index].IRI != nil

}

// GetObjectIRI returns the value safely if IsObjectIRI returned true for the specified index
func (t *Accept) GetObjectIRI(index int) (v *url.URL) {
	return t.object[index].IRI

}

// AppendObjectIRI adds to the back of object a *url.URL type
func (t *Accept) AppendObjectIRI(v *url.URL) {
	t.object = append(t.object, &objectIntermediateType{IRI: v})

}

// PrependObjectIRI adds to the front of object a *url.URL type
func (t *Accept) PrependObjectIRI(v *url.URL) {
	t.object = append([]*objectIntermediateType{&objectIntermediateType{IRI: v}}, t.object...)

}

// RemoveObjectIRI deletes the value from the specified index
func (t *Accept) RemoveObjectIRI(index int) {
	copy(t.object[index:], t.object[index+1:])
	t.object[len(t.object)-1] = nil
	t.object = t.object[:len(t.object)-1]

}

// HasUnknownObject determines whether the call to GetUnknownObject is safe
func (t *Accept) HasUnknownObject() (ok bool) {
	return t.object != nil && t.object[0].unknown_ != nil

}

// GetUnknownObject returns the unknown value for object
func (t *Accept) GetUnknownObject() (v interface{}) {
	return t.object[0].unknown_

}

// SetUnknownObject sets the unknown value of object
func (t *Accept) SetUnknownObject(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &objectIntermediateType{}
	tmp.unknown_ = i
	t.object = append(t.object, tmp)

}

// TargetLen determines the number of elements able to be used for the IsTargetObject, GetTargetObject, and RemoveTargetObject functions
func (t *Accept) TargetLen() (l int) {
	return len(t.target)

}

// IsTargetObject determines whether the call to GetTargetObject is safe for the specified index
func (t *Accept) IsTargetObject(index int) (ok bool) {
	return t.target[index].Object != nil

}

// GetTargetObject returns the value safely if IsTargetObject returned true for the specified index
func (t *Accept) GetTargetObject(index int) (v ObjectType) {
	return t.target[index].Object

}

// AppendTargetObject adds to the back of target a ObjectType type
func (t *Accept) AppendTargetObject(v ObjectType) {
	t.target = append(t.target, &targetIntermediateType{Object: v})

}

// PrependTargetObject adds to the front of target a ObjectType type
func (t *Accept) PrependTargetObject(v ObjectType) {
	t.target = append([]*targetIntermediateType{&targetIntermediateType{Object: v}}, t.target...)

}

// RemoveTargetObject deletes the value from the specified index
func (t *Accept) RemoveTargetObject(index int) {
	copy(t.target[index:], t.target[index+1:])
	t.target[len(t.target)-1] = nil
	t.target = t.target[:len(t.target)-1]

}

// IsTargetLink determines whether the call to GetTargetLink is safe for the specified index
func (t *Accept) IsTargetLink(index int) (ok bool) {
	return t.target[index].Link != nil

}

// GetTargetLink returns the value safely if IsTargetLink returned true for the specified index
func (t *Accept) GetTargetLink(index int) (v LinkType) {
	return t.target[index].Link

}

// AppendTargetLink adds to the back of target a LinkType type
func (t *Accept) AppendTargetLink(v LinkType) {
	t.target = append(t.target, &targetIntermediateType{Link: v})

}

// PrependTargetLink adds to the front of target a LinkType type
func (t *Accept) PrependTargetLink(v LinkType) {
	t.target = append([]*targetIntermediateType{&targetIntermediateType{Link: v}}, t.target...)

}

// RemoveTargetLink deletes the value from the specified index
func (t *Accept) RemoveTargetLink(index int) {
	copy(t.target[index:], t.target[index+1:])
	t.target[len(t.target)-1] = nil
	t.target = t.target[:len(t.target)-1]

}

// IsTargetIRI determines whether the call to GetTargetIRI is safe for the specified index
func (t *Accept) IsTargetIRI(index int) (ok bool) {
	return t.target[index].IRI != nil

}

// GetTargetIRI returns the value safely if IsTargetIRI returned true for the specified index
func (t *Accept) GetTargetIRI(index int) (v *url.URL) {
	return t.target[index].IRI

}

// AppendTargetIRI adds to the back of target a *url.URL type
func (t *Accept) AppendTargetIRI(v *url.URL) {
	t.target = append(t.target, &targetIntermediateType{IRI: v})

}

// PrependTargetIRI adds to the front of target a *url.URL type
func (t *Accept) PrependTargetIRI(v *url.URL) {
	t.target = append([]*targetIntermediateType{&targetIntermediateType{IRI: v}}, t.target...)

}

// RemoveTargetIRI deletes the value from the specified index
func (t *Accept) RemoveTargetIRI(index int) {
	copy(t.target[index:], t.target[index+1:])
	t.target[len(t.target)-1] = nil
	t.target = t.target[:len(t.target)-1]

}

// HasUnknownTarget determines whether the call to GetUnknownTarget is safe
func (t *Accept) HasUnknownTarget() (ok bool) {
	return t.target != nil && t.target[0].unknown_ != nil

}

// GetUnknownTarget returns the unknown value for target
func (t *Accept) GetUnknownTarget() (v interface{}) {
	return t.target[0].unknown_

}

// SetUnknownTarget sets the unknown value of target
func (t *Accept) SetUnknownTarget(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &targetIntermediateType{}
	tmp.unknown_ = i
	t.target = append(t.target, tmp)

}

// ResultLen determines the number of elements able to be used for the IsResultObject, GetResultObject, and RemoveResultObject functions
func (t *Accept) ResultLen() (l int) {
	return len(t.result)

}

// IsResultObject determines whether the call to GetResultObject is safe for the specified index
func (t *Accept) IsResultObject(index int) (ok bool) {
	return t.result[index].Object != nil

}

// GetResultObject returns the value safely if IsResultObject returned true for the specified index
func (t *Accept) GetResultObject(index int) (v ObjectType) {
	return t.result[index].Object

}

// AppendResultObject adds to the back of result a ObjectType type
func (t *Accept) AppendResultObject(v ObjectType) {
	t.result = append(t.result, &resultIntermediateType{Object: v})

}

// PrependResultObject adds to the front of result a ObjectType type
func (t *Accept) PrependResultObject(v ObjectType) {
	t.result = append([]*resultIntermediateType{&resultIntermediateType{Object: v}}, t.result...)

}

// RemoveResultObject deletes the value from the specified index
func (t *Accept) RemoveResultObject(index int) {
	copy(t.result[index:], t.result[index+1:])
	t.result[len(t.result)-1] = nil
	t.result = t.result[:len(t.result)-1]

}

// IsResultLink determines whether the call to GetResultLink is safe for the specified index
func (t *Accept) IsResultLink(index int) (ok bool) {
	return t.result[index].Link != nil

}

// GetResultLink returns the value safely if IsResultLink returned true for the specified index
func (t *Accept) GetResultLink(index int) (v LinkType) {
	return t.result[index].Link

}

// AppendResultLink adds to the back of result a LinkType type
func (t *Accept) AppendResultLink(v LinkType) {
	t.result = append(t.result, &resultIntermediateType{Link: v})

}

// PrependResultLink adds to the front of result a LinkType type
func (t *Accept) PrependResultLink(v LinkType) {
	t.result = append([]*resultIntermediateType{&resultIntermediateType{Link: v}}, t.result...)

}

// RemoveResultLink deletes the value from the specified index
func (t *Accept) RemoveResultLink(index int) {
	copy(t.result[index:], t.result[index+1:])
	t.result[len(t.result)-1] = nil
	t.result = t.result[:len(t.result)-1]

}

// IsResultIRI determines whether the call to GetResultIRI is safe for the specified index
func (t *Accept) IsResultIRI(index int) (ok bool) {
	return t.result[index].IRI != nil

}

// GetResultIRI returns the value safely if IsResultIRI returned true for the specified index
func (t *Accept) GetResultIRI(index int) (v *url.URL) {
	return t.result[index].IRI

}

// AppendResultIRI adds to the back of result a *url.URL type
func (t *Accept) AppendResultIRI(v *url.URL) {
	t.result = append(t.result, &resultIntermediateType{IRI: v})

}

// PrependResultIRI adds to the front of result a *url.URL type
func (t *Accept) PrependResultIRI(v *url.URL) {
	t.result = append([]*resultIntermediateType{&resultIntermediateType{IRI: v}}, t.result...)

}

// RemoveResultIRI deletes the value from the specified index
func (t *Accept) RemoveResultIRI(index int) {
	copy(t.result[index:], t.result[index+1:])
	t.result[len(t.result)-1] = nil
	t.result = t.result[:len(t.result)-1]

}

// HasUnknownResult determines whether the call to GetUnknownResult is safe
func (t *Accept) HasUnknownResult() (ok bool) {
	return t.result != nil && t.result[0].unknown_ != nil

}

// GetUnknownResult returns the unknown value for result
func (t *Accept) GetUnknownResult() (v interface{}) {
	return t.result[0].unknown_

}

// SetUnknownResult sets the unknown value of result
func (t *Accept) SetUnknownResult(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &resultIntermediateType{}
	tmp.unknown_ = i
	t.result = append(t.result, tmp)

}

// OriginLen determines the number of elements able to be used for the IsOriginObject, GetOriginObject, and RemoveOriginObject functions
func (t *Accept) OriginLen() (l int) {
	return len(t.origin)

}

// IsOriginObject determines whether the call to GetOriginObject is safe for the specified index
func (t *Accept) IsOriginObject(index int) (ok bool) {
	return t.origin[index].Object != nil

}

// GetOriginObject returns the value safely if IsOriginObject returned true for the specified index
func (t *Accept) GetOriginObject(index int) (v ObjectType) {
	return t.origin[index].Object

}

// AppendOriginObject adds to the back of origin a ObjectType type
func (t *Accept) AppendOriginObject(v ObjectType) {
	t.origin = append(t.origin, &originIntermediateType{Object: v})

}

// PrependOriginObject adds to the front of origin a ObjectType type
func (t *Accept) PrependOriginObject(v ObjectType) {
	t.origin = append([]*originIntermediateType{&originIntermediateType{Object: v}}, t.origin...)

}

// RemoveOriginObject deletes the value from the specified index
func (t *Accept) RemoveOriginObject(index int) {
	copy(t.origin[index:], t.origin[index+1:])
	t.origin[len(t.origin)-1] = nil
	t.origin = t.origin[:len(t.origin)-1]

}

// IsOriginLink determines whether the call to GetOriginLink is safe for the specified index
func (t *Accept) IsOriginLink(index int) (ok bool) {
	return t.origin[index].Link != nil

}

// GetOriginLink returns the value safely if IsOriginLink returned true for the specified index
func (t *Accept) GetOriginLink(index int) (v LinkType) {
	return t.origin[index].Link

}

// AppendOriginLink adds to the back of origin a LinkType type
func (t *Accept) AppendOriginLink(v LinkType) {
	t.origin = append(t.origin, &originIntermediateType{Link: v})

}

// PrependOriginLink adds to the front of origin a LinkType type
func (t *Accept) PrependOriginLink(v LinkType) {
	t.origin = append([]*originIntermediateType{&originIntermediateType{Link: v}}, t.origin...)

}

// RemoveOriginLink deletes the value from the specified index
func (t *Accept) RemoveOriginLink(index int) {
	copy(t.origin[index:], t.origin[index+1:])
	t.origin[len(t.origin)-1] = nil
	t.origin = t.origin[:len(t.origin)-1]

}

// IsOriginIRI determines whether the call to GetOriginIRI is safe for the specified index
func (t *Accept) IsOriginIRI(index int) (ok bool) {
	return t.origin[index].IRI != nil

}

// GetOriginIRI returns the value safely if IsOriginIRI returned true for the specified index
func (t *Accept) GetOriginIRI(index int) (v *url.URL) {
	return t.origin[index].IRI

}

// AppendOriginIRI adds to the back of origin a *url.URL type
func (t *Accept) AppendOriginIRI(v *url.URL) {
	t.origin = append(t.origin, &originIntermediateType{IRI: v})

}

// PrependOriginIRI adds to the front of origin a *url.URL type
func (t *Accept) PrependOriginIRI(v *url.URL) {
	t.origin = append([]*originIntermediateType{&originIntermediateType{IRI: v}}, t.origin...)

}

// RemoveOriginIRI deletes the value from the specified index
func (t *Accept) RemoveOriginIRI(index int) {
	copy(t.origin[index:], t.origin[index+1:])
	t.origin[len(t.origin)-1] = nil
	t.origin = t.origin[:len(t.origin)-1]

}

// HasUnknownOrigin determines whether the call to GetUnknownOrigin is safe
func (t *Accept) HasUnknownOrigin() (ok bool) {
	return t.origin != nil && t.origin[0].unknown_ != nil

}

// GetUnknownOrigin returns the unknown value for origin
func (t *Accept) GetUnknownOrigin() (v interface{}) {
	return t.origin[0].unknown_

}

// SetUnknownOrigin sets the unknown value of origin
func (t *Accept) SetUnknownOrigin(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &originIntermediateType{}
	tmp.unknown_ = i
	t.origin = append(t.origin, tmp)

}

// InstrumentLen determines the number of elements able to be used for the IsInstrumentObject, GetInstrumentObject, and RemoveInstrumentObject functions
func (t *Accept) InstrumentLen() (l int) {
	return len(t.instrument)

}

// IsInstrumentObject determines whether the call to GetInstrumentObject is safe for the specified index
func (t *Accept) IsInstrumentObject(index int) (ok bool) {
	return t.instrument[index].Object != nil

}

// GetInstrumentObject returns the value safely if IsInstrumentObject returned true for the specified index
func (t *Accept) GetInstrumentObject(index int) (v ObjectType) {
	return t.instrument[index].Object

}

// AppendInstrumentObject adds to the back of instrument a ObjectType type
func (t *Accept) AppendInstrumentObject(v ObjectType) {
	t.instrument = append(t.instrument, &instrumentIntermediateType{Object: v})

}

// PrependInstrumentObject adds to the front of instrument a ObjectType type
func (t *Accept) PrependInstrumentObject(v ObjectType) {
	t.instrument = append([]*instrumentIntermediateType{&instrumentIntermediateType{Object: v}}, t.instrument...)

}

// RemoveInstrumentObject deletes the value from the specified index
func (t *Accept) RemoveInstrumentObject(index int) {
	copy(t.instrument[index:], t.instrument[index+1:])
	t.instrument[len(t.instrument)-1] = nil
	t.instrument = t.instrument[:len(t.instrument)-1]

}

// IsInstrumentLink determines whether the call to GetInstrumentLink is safe for the specified index
func (t *Accept) IsInstrumentLink(index int) (ok bool) {
	return t.instrument[index].Link != nil

}

// GetInstrumentLink returns the value safely if IsInstrumentLink returned true for the specified index
func (t *Accept) GetInstrumentLink(index int) (v LinkType) {
	return t.instrument[index].Link

}

// AppendInstrumentLink adds to the back of instrument a LinkType type
func (t *Accept) AppendInstrumentLink(v LinkType) {
	t.instrument = append(t.instrument, &instrumentIntermediateType{Link: v})

}

// PrependInstrumentLink adds to the front of instrument a LinkType type
func (t *Accept) PrependInstrumentLink(v LinkType) {
	t.instrument = append([]*instrumentIntermediateType{&instrumentIntermediateType{Link: v}}, t.instrument...)

}

// RemoveInstrumentLink deletes the value from the specified index
func (t *Accept) RemoveInstrumentLink(index int) {
	copy(t.instrument[index:], t.instrument[index+1:])
	t.instrument[len(t.instrument)-1] = nil
	t.instrument = t.instrument[:len(t.instrument)-1]

}

// IsInstrumentIRI determines whether the call to GetInstrumentIRI is safe for the specified index
func (t *Accept) IsInstrumentIRI(index int) (ok bool) {
	return t.instrument[index].IRI != nil

}

// GetInstrumentIRI returns the value safely if IsInstrumentIRI returned true for the specified index
func (t *Accept) GetInstrumentIRI(index int) (v *url.URL) {
	return t.instrument[index].IRI

}

// AppendInstrumentIRI adds to the back of instrument a *url.URL type
func (t *Accept) AppendInstrumentIRI(v *url.URL) {
	t.instrument = append(t.instrument, &instrumentIntermediateType{IRI: v})

}

// PrependInstrumentIRI adds to the front of instrument a *url.URL type
func (t *Accept) PrependInstrumentIRI(v *url.URL) {
	t.instrument = append([]*instrumentIntermediateType{&instrumentIntermediateType{IRI: v}}, t.instrument...)

}

// RemoveInstrumentIRI deletes the value from the specified index
func (t *Accept) RemoveInstrumentIRI(index int) {
	copy(t.instrument[index:], t.instrument[index+1:])
	t.instrument[len(t.instrument)-1] = nil
	t.instrument = t.instrument[:len(t.instrument)-1]

}

// HasUnknownInstrument determines whether the call to GetUnknownInstrument is safe
func (t *Accept) HasUnknownInstrument() (ok bool) {
	return t.instrument != nil && t.instrument[0].unknown_ != nil

}

// GetUnknownInstrument returns the unknown value for instrument
func (t *Accept) GetUnknownInstrument() (v interface{}) {
	return t.instrument[0].unknown_

}

// SetUnknownInstrument sets the unknown value of instrument
func (t *Accept) SetUnknownInstrument(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &instrumentIntermediateType{}
	tmp.unknown_ = i
	t.instrument = append(t.instrument, tmp)

}

// IsAltitude determines whether the call to GetAltitude is safe
func (t *Accept) IsAltitude() (ok bool) {
	return t.altitude != nil && t.altitude.float != nil

}

// GetAltitude returns the value safely if IsAltitude returned true
func (t *Accept) GetAltitude() (v float64) {
	return *t.altitude.float

}

// SetAltitude sets the value of altitude to be of float64 type
func (t *Accept) SetAltitude(v float64) {
	t.altitude = &altitudeIntermediateType{float: &v}

}

// IsAltitudeIRI determines whether the call to GetAltitudeIRI is safe
func (t *Accept) IsAltitudeIRI() (ok bool) {
	return t.altitude != nil && t.altitude.IRI != nil

}

// GetAltitudeIRI returns the value safely if IsAltitudeIRI returned true
func (t *Accept) GetAltitudeIRI() (v *url.URL) {
	return t.altitude.IRI

}

// SetAltitudeIRI sets the value of altitude to be of *url.URL type
func (t *Accept) SetAltitudeIRI(v *url.URL) {
	t.altitude = &altitudeIntermediateType{IRI: v}

}

// HasUnknownAltitude determines whether the call to GetUnknownAltitude is safe
func (t *Accept) HasUnknownAltitude() (ok bool) {
	return t.altitude != nil && t.altitude.unknown_ != nil

}

// GetUnknownAltitude returns the unknown value for altitude
func (t *Accept) GetUnknownAltitude() (v interface{}) {
	return t.altitude.unknown_

}

// SetUnknownAltitude sets the unknown value of altitude
func (t *Accept) SetUnknownAltitude(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &altitudeIntermediateType{}
	tmp.unknown_ = i
	t.altitude = tmp

}

// AttachmentLen determines the number of elements able to be used for the IsAttachmentObject, GetAttachmentObject, and RemoveAttachmentObject functions
func (t *Accept) AttachmentLen() (l int) {
	return len(t.attachment)

}

// IsAttachmentObject determines whether the call to GetAttachmentObject is safe for the specified index
func (t *Accept) IsAttachmentObject(index int) (ok bool) {
	return t.attachment[index].Object != nil

}

// GetAttachmentObject returns the value safely if IsAttachmentObject returned true for the specified index
func (t *Accept) GetAttachmentObject(index int) (v ObjectType) {
	return t.attachment[index].Object

}

// AppendAttachmentObject adds to the back of attachment a ObjectType type
func (t *Accept) AppendAttachmentObject(v ObjectType) {
	t.attachment = append(t.attachment, &attachmentIntermediateType{Object: v})

}

// PrependAttachmentObject adds to the front of attachment a ObjectType type
func (t *Accept) PrependAttachmentObject(v ObjectType) {
	t.attachment = append([]*attachmentIntermediateType{&attachmentIntermediateType{Object: v}}, t.attachment...)

}

// RemoveAttachmentObject deletes the value from the specified index
func (t *Accept) RemoveAttachmentObject(index int) {
	copy(t.attachment[index:], t.attachment[index+1:])
	t.attachment[len(t.attachment)-1] = nil
	t.attachment = t.attachment[:len(t.attachment)-1]

}

// IsAttachmentLink determines whether the call to GetAttachmentLink is safe for the specified index
func (t *Accept) IsAttachmentLink(index int) (ok bool) {
	return t.attachment[index].Link != nil

}

// GetAttachmentLink returns the value safely if IsAttachmentLink returned true for the specified index
func (t *Accept) GetAttachmentLink(index int) (v LinkType) {
	return t.attachment[index].Link

}

// AppendAttachmentLink adds to the back of attachment a LinkType type
func (t *Accept) AppendAttachmentLink(v LinkType) {
	t.attachment = append(t.attachment, &attachmentIntermediateType{Link: v})

}

// PrependAttachmentLink adds to the front of attachment a LinkType type
func (t *Accept) PrependAttachmentLink(v LinkType) {
	t.attachment = append([]*attachmentIntermediateType{&attachmentIntermediateType{Link: v}}, t.attachment...)

}

// RemoveAttachmentLink deletes the value from the specified index
func (t *Accept) RemoveAttachmentLink(index int) {
	copy(t.attachment[index:], t.attachment[index+1:])
	t.attachment[len(t.attachment)-1] = nil
	t.attachment = t.attachment[:len(t.attachment)-1]

}

// IsAttachmentIRI determines whether the call to GetAttachmentIRI is safe for the specified index
func (t *Accept) IsAttachmentIRI(index int) (ok bool) {
	return t.attachment[index].IRI != nil

}

// GetAttachmentIRI returns the value safely if IsAttachmentIRI returned true for the specified index
func (t *Accept) GetAttachmentIRI(index int) (v *url.URL) {
	return t.attachment[index].IRI

}

// AppendAttachmentIRI adds to the back of attachment a *url.URL type
func (t *Accept) AppendAttachmentIRI(v *url.URL) {
	t.attachment = append(t.attachment, &attachmentIntermediateType{IRI: v})

}

// PrependAttachmentIRI adds to the front of attachment a *url.URL type
func (t *Accept) PrependAttachmentIRI(v *url.URL) {
	t.attachment = append([]*attachmentIntermediateType{&attachmentIntermediateType{IRI: v}}, t.attachment...)

}

// RemoveAttachmentIRI deletes the value from the specified index
func (t *Accept) RemoveAttachmentIRI(index int) {
	copy(t.attachment[index:], t.attachment[index+1:])
	t.attachment[len(t.attachment)-1] = nil
	t.attachment = t.attachment[:len(t.attachment)-1]

}

// HasUnknownAttachment determines whether the call to GetUnknownAttachment is safe
func (t *Accept) HasUnknownAttachment() (ok bool) {
	return t.attachment != nil && t.attachment[0].unknown_ != nil

}

// GetUnknownAttachment returns the unknown value for attachment
func (t *Accept) GetUnknownAttachment() (v interface{}) {
	return t.attachment[0].unknown_

}

// SetUnknownAttachment sets the unknown value of attachment
func (t *Accept) SetUnknownAttachment(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &attachmentIntermediateType{}
	tmp.unknown_ = i
	t.attachment = append(t.attachment, tmp)

}

// AttributedToLen determines the number of elements able to be used for the IsAttributedToObject, GetAttributedToObject, and RemoveAttributedToObject functions
func (t *Accept) AttributedToLen() (l int) {
	return len(t.attributedTo)

}

// IsAttributedToObject determines whether the call to GetAttributedToObject is safe for the specified index
func (t *Accept) IsAttributedToObject(index int) (ok bool) {
	return t.attributedTo[index].Object != nil

}

// GetAttributedToObject returns the value safely if IsAttributedToObject returned true for the specified index
func (t *Accept) GetAttributedToObject(index int) (v ObjectType) {
	return t.attributedTo[index].Object

}

// AppendAttributedToObject adds to the back of attributedTo a ObjectType type
func (t *Accept) AppendAttributedToObject(v ObjectType) {
	t.attributedTo = append(t.attributedTo, &attributedToIntermediateType{Object: v})

}

// PrependAttributedToObject adds to the front of attributedTo a ObjectType type
func (t *Accept) PrependAttributedToObject(v ObjectType) {
	t.attributedTo = append([]*attributedToIntermediateType{&attributedToIntermediateType{Object: v}}, t.attributedTo...)

}

// RemoveAttributedToObject deletes the value from the specified index
func (t *Accept) RemoveAttributedToObject(index int) {
	copy(t.attributedTo[index:], t.attributedTo[index+1:])
	t.attributedTo[len(t.attributedTo)-1] = nil
	t.attributedTo = t.attributedTo[:len(t.attributedTo)-1]

}

// IsAttributedToLink determines whether the call to GetAttributedToLink is safe for the specified index
func (t *Accept) IsAttributedToLink(index int) (ok bool) {
	return t.attributedTo[index].Link != nil

}

// GetAttributedToLink returns the value safely if IsAttributedToLink returned true for the specified index
func (t *Accept) GetAttributedToLink(index int) (v LinkType) {
	return t.attributedTo[index].Link

}

// AppendAttributedToLink adds to the back of attributedTo a LinkType type
func (t *Accept) AppendAttributedToLink(v LinkType) {
	t.attributedTo = append(t.attributedTo, &attributedToIntermediateType{Link: v})

}

// PrependAttributedToLink adds to the front of attributedTo a LinkType type
func (t *Accept) PrependAttributedToLink(v LinkType) {
	t.attributedTo = append([]*attributedToIntermediateType{&attributedToIntermediateType{Link: v}}, t.attributedTo...)

}

// RemoveAttributedToLink deletes the value from the specified index
func (t *Accept) RemoveAttributedToLink(index int) {
	copy(t.attributedTo[index:], t.attributedTo[index+1:])
	t.attributedTo[len(t.attributedTo)-1] = nil
	t.attributedTo = t.attributedTo[:len(t.attributedTo)-1]

}

// IsAttributedToIRI determines whether the call to GetAttributedToIRI is safe for the specified index
func (t *Accept) IsAttributedToIRI(index int) (ok bool) {
	return t.attributedTo[index].IRI != nil

}

// GetAttributedToIRI returns the value safely if IsAttributedToIRI returned true for the specified index
func (t *Accept) GetAttributedToIRI(index int) (v *url.URL) {
	return t.attributedTo[index].IRI

}

// AppendAttributedToIRI adds to the back of attributedTo a *url.URL type
func (t *Accept) AppendAttributedToIRI(v *url.URL) {
	t.attributedTo = append(t.attributedTo, &attributedToIntermediateType{IRI: v})

}

// PrependAttributedToIRI adds to the front of attributedTo a *url.URL type
func (t *Accept) PrependAttributedToIRI(v *url.URL) {
	t.attributedTo = append([]*attributedToIntermediateType{&attributedToIntermediateType{IRI: v}}, t.attributedTo...)

}

// RemoveAttributedToIRI deletes the value from the specified index
func (t *Accept) RemoveAttributedToIRI(index int) {
	copy(t.attributedTo[index:], t.attributedTo[index+1:])
	t.attributedTo[len(t.attributedTo)-1] = nil
	t.attributedTo = t.attributedTo[:len(t.attributedTo)-1]

}

// HasUnknownAttributedTo determines whether the call to GetUnknownAttributedTo is safe
func (t *Accept) HasUnknownAttributedTo() (ok bool) {
	return t.attributedTo != nil && t.attributedTo[0].unknown_ != nil

}

// GetUnknownAttributedTo returns the unknown value for attributedTo
func (t *Accept) GetUnknownAttributedTo() (v interface{}) {
	return t.attributedTo[0].unknown_

}

// SetUnknownAttributedTo sets the unknown value of attributedTo
func (t *Accept) SetUnknownAttributedTo(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &attributedToIntermediateType{}
	tmp.unknown_ = i
	t.attributedTo = append(t.attributedTo, tmp)

}

// AudienceLen determines the number of elements able to be used for the IsAudienceObject, GetAudienceObject, and RemoveAudienceObject functions
func (t *Accept) AudienceLen() (l int) {
	return len(t.audience)

}

// IsAudienceObject determines whether the call to GetAudienceObject is safe for the specified index
func (t *Accept) IsAudienceObject(index int) (ok bool) {
	return t.audience[index].Object != nil

}

// GetAudienceObject returns the value safely if IsAudienceObject returned true for the specified index
func (t *Accept) GetAudienceObject(index int) (v ObjectType) {
	return t.audience[index].Object

}

// AppendAudienceObject adds to the back of audience a ObjectType type
func (t *Accept) AppendAudienceObject(v ObjectType) {
	t.audience = append(t.audience, &audienceIntermediateType{Object: v})

}

// PrependAudienceObject adds to the front of audience a ObjectType type
func (t *Accept) PrependAudienceObject(v ObjectType) {
	t.audience = append([]*audienceIntermediateType{&audienceIntermediateType{Object: v}}, t.audience...)

}

// RemoveAudienceObject deletes the value from the specified index
func (t *Accept) RemoveAudienceObject(index int) {
	copy(t.audience[index:], t.audience[index+1:])
	t.audience[len(t.audience)-1] = nil
	t.audience = t.audience[:len(t.audience)-1]

}

// IsAudienceLink determines whether the call to GetAudienceLink is safe for the specified index
func (t *Accept) IsAudienceLink(index int) (ok bool) {
	return t.audience[index].Link != nil

}

// GetAudienceLink returns the value safely if IsAudienceLink returned true for the specified index
func (t *Accept) GetAudienceLink(index int) (v LinkType) {
	return t.audience[index].Link

}

// AppendAudienceLink adds to the back of audience a LinkType type
func (t *Accept) AppendAudienceLink(v LinkType) {
	t.audience = append(t.audience, &audienceIntermediateType{Link: v})

}

// PrependAudienceLink adds to the front of audience a LinkType type
func (t *Accept) PrependAudienceLink(v LinkType) {
	t.audience = append([]*audienceIntermediateType{&audienceIntermediateType{Link: v}}, t.audience...)

}

// RemoveAudienceLink deletes the value from the specified index
func (t *Accept) RemoveAudienceLink(index int) {
	copy(t.audience[index:], t.audience[index+1:])
	t.audience[len(t.audience)-1] = nil
	t.audience = t.audience[:len(t.audience)-1]

}

// IsAudienceIRI determines whether the call to GetAudienceIRI is safe for the specified index
func (t *Accept) IsAudienceIRI(index int) (ok bool) {
	return t.audience[index].IRI != nil

}

// GetAudienceIRI returns the value safely if IsAudienceIRI returned true for the specified index
func (t *Accept) GetAudienceIRI(index int) (v *url.URL) {
	return t.audience[index].IRI

}

// AppendAudienceIRI adds to the back of audience a *url.URL type
func (t *Accept) AppendAudienceIRI(v *url.URL) {
	t.audience = append(t.audience, &audienceIntermediateType{IRI: v})

}

// PrependAudienceIRI adds to the front of audience a *url.URL type
func (t *Accept) PrependAudienceIRI(v *url.URL) {
	t.audience = append([]*audienceIntermediateType{&audienceIntermediateType{IRI: v}}, t.audience...)

}

// RemoveAudienceIRI deletes the value from the specified index
func (t *Accept) RemoveAudienceIRI(index int) {
	copy(t.audience[index:], t.audience[index+1:])
	t.audience[len(t.audience)-1] = nil
	t.audience = t.audience[:len(t.audience)-1]

}

// HasUnknownAudience determines whether the call to GetUnknownAudience is safe
func (t *Accept) HasUnknownAudience() (ok bool) {
	return t.audience != nil && t.audience[0].unknown_ != nil

}

// GetUnknownAudience returns the unknown value for audience
func (t *Accept) GetUnknownAudience() (v interface{}) {
	return t.audience[0].unknown_

}

// SetUnknownAudience sets the unknown value of audience
func (t *Accept) SetUnknownAudience(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &audienceIntermediateType{}
	tmp.unknown_ = i
	t.audience = append(t.audience, tmp)

}

// ContentLen determines the number of elements able to be used for the IsContentString, GetContentString, and RemoveContentString functions
func (t *Accept) ContentLen() (l int) {
	return len(t.content)

}

// IsContentString determines whether the call to GetContentString is safe for the specified index
func (t *Accept) IsContentString(index int) (ok bool) {
	return t.content[index].stringName != nil

}

// GetContentString returns the value safely if IsContentString returned true for the specified index
func (t *Accept) GetContentString(index int) (v string) {
	return *t.content[index].stringName

}

// AppendContentString adds to the back of content a string type
func (t *Accept) AppendContentString(v string) {
	t.content = append(t.content, &contentIntermediateType{stringName: &v})

}

// PrependContentString adds to the front of content a string type
func (t *Accept) PrependContentString(v string) {
	t.content = append([]*contentIntermediateType{&contentIntermediateType{stringName: &v}}, t.content...)

}

// RemoveContentString deletes the value from the specified index
func (t *Accept) RemoveContentString(index int) {
	copy(t.content[index:], t.content[index+1:])
	t.content[len(t.content)-1] = nil
	t.content = t.content[:len(t.content)-1]

}

// IsContentLangString determines whether the call to GetContentLangString is safe for the specified index
func (t *Accept) IsContentLangString(index int) (ok bool) {
	return t.content[index].langString != nil

}

// GetContentLangString returns the value safely if IsContentLangString returned true for the specified index
func (t *Accept) GetContentLangString(index int) (v string) {
	return *t.content[index].langString

}

// AppendContentLangString adds to the back of content a string type
func (t *Accept) AppendContentLangString(v string) {
	t.content = append(t.content, &contentIntermediateType{langString: &v})

}

// PrependContentLangString adds to the front of content a string type
func (t *Accept) PrependContentLangString(v string) {
	t.content = append([]*contentIntermediateType{&contentIntermediateType{langString: &v}}, t.content...)

}

// RemoveContentLangString deletes the value from the specified index
func (t *Accept) RemoveContentLangString(index int) {
	copy(t.content[index:], t.content[index+1:])
	t.content[len(t.content)-1] = nil
	t.content = t.content[:len(t.content)-1]

}

// IsContentIRI determines whether the call to GetContentIRI is safe for the specified index
func (t *Accept) IsContentIRI(index int) (ok bool) {
	return t.content[index].IRI != nil

}

// GetContentIRI returns the value safely if IsContentIRI returned true for the specified index
func (t *Accept) GetContentIRI(index int) (v *url.URL) {
	return t.content[index].IRI

}

// AppendContentIRI adds to the back of content a *url.URL type
func (t *Accept) AppendContentIRI(v *url.URL) {
	t.content = append(t.content, &contentIntermediateType{IRI: v})

}

// PrependContentIRI adds to the front of content a *url.URL type
func (t *Accept) PrependContentIRI(v *url.URL) {
	t.content = append([]*contentIntermediateType{&contentIntermediateType{IRI: v}}, t.content...)

}

// RemoveContentIRI deletes the value from the specified index
func (t *Accept) RemoveContentIRI(index int) {
	copy(t.content[index:], t.content[index+1:])
	t.content[len(t.content)-1] = nil
	t.content = t.content[:len(t.content)-1]

}

// HasUnknownContent determines whether the call to GetUnknownContent is safe
func (t *Accept) HasUnknownContent() (ok bool) {
	return t.content != nil && t.content[0].unknown_ != nil

}

// GetUnknownContent returns the unknown value for content
func (t *Accept) GetUnknownContent() (v interface{}) {
	return t.content[0].unknown_

}

// SetUnknownContent sets the unknown value of content
func (t *Accept) SetUnknownContent(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &contentIntermediateType{}
	tmp.unknown_ = i
	t.content = append(t.content, tmp)

}

// ContentMapLanguages returns all languages for this property's language mapping, or nil if there are none.
func (t *Accept) ContentMapLanguages() (l []string) {
	if t.contentMap == nil || len(t.contentMap) == 0 {
		return nil
	}
	for k := range t.contentMap {
		l = append(l, k)
	}
	return

}

// GetContentMap retrieves the value of the property for the specified language, or an empty string if it does not exist
func (t *Accept) GetContentMap(l string) (v string) {
	if t.contentMap == nil {
		return ""
	}
	ok := false
	v, ok = t.contentMap[l]
	if !ok {
		return ""
	}
	return v

}

// SetContentMap sets the value of the property for the specified language
func (t *Accept) SetContentMap(l string, v string) {
	if t.contentMap == nil {
		t.contentMap = make(map[string]string)
	}
	t.contentMap[l] = v

}

// ContextLen determines the number of elements able to be used for the IsContextObject, GetContextObject, and RemoveContextObject functions
func (t *Accept) ContextLen() (l int) {
	return len(t.context)

}

// IsContextObject determines whether the call to GetContextObject is safe for the specified index
func (t *Accept) IsContextObject(index int) (ok bool) {
	return t.context[index].Object != nil

}

// GetContextObject returns the value safely if IsContextObject returned true for the specified index
func (t *Accept) GetContextObject(index int) (v ObjectType) {
	return t.context[index].Object

}

// AppendContextObject adds to the back of context a ObjectType type
func (t *Accept) AppendContextObject(v ObjectType) {
	t.context = append(t.context, &contextIntermediateType{Object: v})

}

// PrependContextObject adds to the front of context a ObjectType type
func (t *Accept) PrependContextObject(v ObjectType) {
	t.context = append([]*contextIntermediateType{&contextIntermediateType{Object: v}}, t.context...)

}

// RemoveContextObject deletes the value from the specified index
func (t *Accept) RemoveContextObject(index int) {
	copy(t.context[index:], t.context[index+1:])
	t.context[len(t.context)-1] = nil
	t.context = t.context[:len(t.context)-1]

}

// IsContextLink determines whether the call to GetContextLink is safe for the specified index
func (t *Accept) IsContextLink(index int) (ok bool) {
	return t.context[index].Link != nil

}

// GetContextLink returns the value safely if IsContextLink returned true for the specified index
func (t *Accept) GetContextLink(index int) (v LinkType) {
	return t.context[index].Link

}

// AppendContextLink adds to the back of context a LinkType type
func (t *Accept) AppendContextLink(v LinkType) {
	t.context = append(t.context, &contextIntermediateType{Link: v})

}

// PrependContextLink adds to the front of context a LinkType type
func (t *Accept) PrependContextLink(v LinkType) {
	t.context = append([]*contextIntermediateType{&contextIntermediateType{Link: v}}, t.context...)

}

// RemoveContextLink deletes the value from the specified index
func (t *Accept) RemoveContextLink(index int) {
	copy(t.context[index:], t.context[index+1:])
	t.context[len(t.context)-1] = nil
	t.context = t.context[:len(t.context)-1]

}

// IsContextIRI determines whether the call to GetContextIRI is safe for the specified index
func (t *Accept) IsContextIRI(index int) (ok bool) {
	return t.context[index].IRI != nil

}

// GetContextIRI returns the value safely if IsContextIRI returned true for the specified index
func (t *Accept) GetContextIRI(index int) (v *url.URL) {
	return t.context[index].IRI

}

// AppendContextIRI adds to the back of context a *url.URL type
func (t *Accept) AppendContextIRI(v *url.URL) {
	t.context = append(t.context, &contextIntermediateType{IRI: v})

}

// PrependContextIRI adds to the front of context a *url.URL type
func (t *Accept) PrependContextIRI(v *url.URL) {
	t.context = append([]*contextIntermediateType{&contextIntermediateType{IRI: v}}, t.context...)

}

// RemoveContextIRI deletes the value from the specified index
func (t *Accept) RemoveContextIRI(index int) {
	copy(t.context[index:], t.context[index+1:])
	t.context[len(t.context)-1] = nil
	t.context = t.context[:len(t.context)-1]

}

// HasUnknownContext determines whether the call to GetUnknownContext is safe
func (t *Accept) HasUnknownContext() (ok bool) {
	return t.context != nil && t.context[0].unknown_ != nil

}

// GetUnknownContext returns the unknown value for context
func (t *Accept) GetUnknownContext() (v interface{}) {
	return t.context[0].unknown_

}

// SetUnknownContext sets the unknown value of context
func (t *Accept) SetUnknownContext(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &contextIntermediateType{}
	tmp.unknown_ = i
	t.context = append(t.context, tmp)

}

// NameLen determines the number of elements able to be used for the IsNameString, GetNameString, and RemoveNameString functions
func (t *Accept) NameLen() (l int) {
	return len(t.name)

}

// IsNameString determines whether the call to GetNameString is safe for the specified index
func (t *Accept) IsNameString(index int) (ok bool) {
	return t.name[index].stringName != nil

}

// GetNameString returns the value safely if IsNameString returned true for the specified index
func (t *Accept) GetNameString(index int) (v string) {
	return *t.name[index].stringName

}

// AppendNameString adds to the back of name a string type
func (t *Accept) AppendNameString(v string) {
	t.name = append(t.name, &nameIntermediateType{stringName: &v})

}

// PrependNameString adds to the front of name a string type
func (t *Accept) PrependNameString(v string) {
	t.name = append([]*nameIntermediateType{&nameIntermediateType{stringName: &v}}, t.name...)

}

// RemoveNameString deletes the value from the specified index
func (t *Accept) RemoveNameString(index int) {
	copy(t.name[index:], t.name[index+1:])
	t.name[len(t.name)-1] = nil
	t.name = t.name[:len(t.name)-1]

}

// IsNameLangString determines whether the call to GetNameLangString is safe for the specified index
func (t *Accept) IsNameLangString(index int) (ok bool) {
	return t.name[index].langString != nil

}

// GetNameLangString returns the value safely if IsNameLangString returned true for the specified index
func (t *Accept) GetNameLangString(index int) (v string) {
	return *t.name[index].langString

}

// AppendNameLangString adds to the back of name a string type
func (t *Accept) AppendNameLangString(v string) {
	t.name = append(t.name, &nameIntermediateType{langString: &v})

}

// PrependNameLangString adds to the front of name a string type
func (t *Accept) PrependNameLangString(v string) {
	t.name = append([]*nameIntermediateType{&nameIntermediateType{langString: &v}}, t.name...)

}

// RemoveNameLangString deletes the value from the specified index
func (t *Accept) RemoveNameLangString(index int) {
	copy(t.name[index:], t.name[index+1:])
	t.name[len(t.name)-1] = nil
	t.name = t.name[:len(t.name)-1]

}

// IsNameIRI determines whether the call to GetNameIRI is safe for the specified index
func (t *Accept) IsNameIRI(index int) (ok bool) {
	return t.name[index].IRI != nil

}

// GetNameIRI returns the value safely if IsNameIRI returned true for the specified index
func (t *Accept) GetNameIRI(index int) (v *url.URL) {
	return t.name[index].IRI

}

// AppendNameIRI adds to the back of name a *url.URL type
func (t *Accept) AppendNameIRI(v *url.URL) {
	t.name = append(t.name, &nameIntermediateType{IRI: v})

}

// PrependNameIRI adds to the front of name a *url.URL type
func (t *Accept) PrependNameIRI(v *url.URL) {
	t.name = append([]*nameIntermediateType{&nameIntermediateType{IRI: v}}, t.name...)

}

// RemoveNameIRI deletes the value from the specified index
func (t *Accept) RemoveNameIRI(index int) {
	copy(t.name[index:], t.name[index+1:])
	t.name[len(t.name)-1] = nil
	t.name = t.name[:len(t.name)-1]

}

// HasUnknownName determines whether the call to GetUnknownName is safe
func (t *Accept) HasUnknownName() (ok bool) {
	return t.name != nil && t.name[0].unknown_ != nil

}

// GetUnknownName returns the unknown value for name
func (t *Accept) GetUnknownName() (v interface{}) {
	return t.name[0].unknown_

}

// SetUnknownName sets the unknown value of name
func (t *Accept) SetUnknownName(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &nameIntermediateType{}
	tmp.unknown_ = i
	t.name = append(t.name, tmp)

}

// NameMapLanguages returns all languages for this property's language mapping, or nil if there are none.
func (t *Accept) NameMapLanguages() (l []string) {
	if t.nameMap == nil || len(t.nameMap) == 0 {
		return nil
	}
	for k := range t.nameMap {
		l = append(l, k)
	}
	return

}

// GetNameMap retrieves the value of the property for the specified language, or an empty string if it does not exist
func (t *Accept) GetNameMap(l string) (v string) {
	if t.nameMap == nil {
		return ""
	}
	ok := false
	v, ok = t.nameMap[l]
	if !ok {
		return ""
	}
	return v

}

// SetNameMap sets the value of the property for the specified language
func (t *Accept) SetNameMap(l string, v string) {
	if t.nameMap == nil {
		t.nameMap = make(map[string]string)
	}
	t.nameMap[l] = v

}

// IsEndTime determines whether the call to GetEndTime is safe
func (t *Accept) IsEndTime() (ok bool) {
	return t.endTime != nil && t.endTime.dateTime != nil

}

// GetEndTime returns the value safely if IsEndTime returned true
func (t *Accept) GetEndTime() (v time.Time) {
	return *t.endTime.dateTime

}

// SetEndTime sets the value of endTime to be of time.Time type
func (t *Accept) SetEndTime(v time.Time) {
	t.endTime = &endTimeIntermediateType{dateTime: &v}

}

// IsEndTimeIRI determines whether the call to GetEndTimeIRI is safe
func (t *Accept) IsEndTimeIRI() (ok bool) {
	return t.endTime != nil && t.endTime.IRI != nil

}

// GetEndTimeIRI returns the value safely if IsEndTimeIRI returned true
func (t *Accept) GetEndTimeIRI() (v *url.URL) {
	return t.endTime.IRI

}

// SetEndTimeIRI sets the value of endTime to be of *url.URL type
func (t *Accept) SetEndTimeIRI(v *url.URL) {
	t.endTime = &endTimeIntermediateType{IRI: v}

}

// HasUnknownEndTime determines whether the call to GetUnknownEndTime is safe
func (t *Accept) HasUnknownEndTime() (ok bool) {
	return t.endTime != nil && t.endTime.unknown_ != nil

}

// GetUnknownEndTime returns the unknown value for endTime
func (t *Accept) GetUnknownEndTime() (v interface{}) {
	return t.endTime.unknown_

}

// SetUnknownEndTime sets the unknown value of endTime
func (t *Accept) SetUnknownEndTime(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &endTimeIntermediateType{}
	tmp.unknown_ = i
	t.endTime = tmp

}

// GeneratorLen determines the number of elements able to be used for the IsGeneratorObject, GetGeneratorObject, and RemoveGeneratorObject functions
func (t *Accept) GeneratorLen() (l int) {
	return len(t.generator)

}

// IsGeneratorObject determines whether the call to GetGeneratorObject is safe for the specified index
func (t *Accept) IsGeneratorObject(index int) (ok bool) {
	return t.generator[index].Object != nil

}

// GetGeneratorObject returns the value safely if IsGeneratorObject returned true for the specified index
func (t *Accept) GetGeneratorObject(index int) (v ObjectType) {
	return t.generator[index].Object

}

// AppendGeneratorObject adds to the back of generator a ObjectType type
func (t *Accept) AppendGeneratorObject(v ObjectType) {
	t.generator = append(t.generator, &generatorIntermediateType{Object: v})

}

// PrependGeneratorObject adds to the front of generator a ObjectType type
func (t *Accept) PrependGeneratorObject(v ObjectType) {
	t.generator = append([]*generatorIntermediateType{&generatorIntermediateType{Object: v}}, t.generator...)

}

// RemoveGeneratorObject deletes the value from the specified index
func (t *Accept) RemoveGeneratorObject(index int) {
	copy(t.generator[index:], t.generator[index+1:])
	t.generator[len(t.generator)-1] = nil
	t.generator = t.generator[:len(t.generator)-1]

}

// IsGeneratorLink determines whether the call to GetGeneratorLink is safe for the specified index
func (t *Accept) IsGeneratorLink(index int) (ok bool) {
	return t.generator[index].Link != nil

}

// GetGeneratorLink returns the value safely if IsGeneratorLink returned true for the specified index
func (t *Accept) GetGeneratorLink(index int) (v LinkType) {
	return t.generator[index].Link

}

// AppendGeneratorLink adds to the back of generator a LinkType type
func (t *Accept) AppendGeneratorLink(v LinkType) {
	t.generator = append(t.generator, &generatorIntermediateType{Link: v})

}

// PrependGeneratorLink adds to the front of generator a LinkType type
func (t *Accept) PrependGeneratorLink(v LinkType) {
	t.generator = append([]*generatorIntermediateType{&generatorIntermediateType{Link: v}}, t.generator...)

}

// RemoveGeneratorLink deletes the value from the specified index
func (t *Accept) RemoveGeneratorLink(index int) {
	copy(t.generator[index:], t.generator[index+1:])
	t.generator[len(t.generator)-1] = nil
	t.generator = t.generator[:len(t.generator)-1]

}

// IsGeneratorIRI determines whether the call to GetGeneratorIRI is safe for the specified index
func (t *Accept) IsGeneratorIRI(index int) (ok bool) {
	return t.generator[index].IRI != nil

}

// GetGeneratorIRI returns the value safely if IsGeneratorIRI returned true for the specified index
func (t *Accept) GetGeneratorIRI(index int) (v *url.URL) {
	return t.generator[index].IRI

}

// AppendGeneratorIRI adds to the back of generator a *url.URL type
func (t *Accept) AppendGeneratorIRI(v *url.URL) {
	t.generator = append(t.generator, &generatorIntermediateType{IRI: v})

}

// PrependGeneratorIRI adds to the front of generator a *url.URL type
func (t *Accept) PrependGeneratorIRI(v *url.URL) {
	t.generator = append([]*generatorIntermediateType{&generatorIntermediateType{IRI: v}}, t.generator...)

}

// RemoveGeneratorIRI deletes the value from the specified index
func (t *Accept) RemoveGeneratorIRI(index int) {
	copy(t.generator[index:], t.generator[index+1:])
	t.generator[len(t.generator)-1] = nil
	t.generator = t.generator[:len(t.generator)-1]

}

// HasUnknownGenerator determines whether the call to GetUnknownGenerator is safe
func (t *Accept) HasUnknownGenerator() (ok bool) {
	return t.generator != nil && t.generator[0].unknown_ != nil

}

// GetUnknownGenerator returns the unknown value for generator
func (t *Accept) GetUnknownGenerator() (v interface{}) {
	return t.generator[0].unknown_

}

// SetUnknownGenerator sets the unknown value of generator
func (t *Accept) SetUnknownGenerator(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &generatorIntermediateType{}
	tmp.unknown_ = i
	t.generator = append(t.generator, tmp)

}

// IconLen determines the number of elements able to be used for the IsIconImage, GetIconImage, and RemoveIconImage functions
func (t *Accept) IconLen() (l int) {
	return len(t.icon)

}

// IsIconImage determines whether the call to GetIconImage is safe for the specified index
func (t *Accept) IsIconImage(index int) (ok bool) {
	return t.icon[index].Image != nil

}

// GetIconImage returns the value safely if IsIconImage returned true for the specified index
func (t *Accept) GetIconImage(index int) (v ImageType) {
	return t.icon[index].Image

}

// AppendIconImage adds to the back of icon a ImageType type
func (t *Accept) AppendIconImage(v ImageType) {
	t.icon = append(t.icon, &iconIntermediateType{Image: v})

}

// PrependIconImage adds to the front of icon a ImageType type
func (t *Accept) PrependIconImage(v ImageType) {
	t.icon = append([]*iconIntermediateType{&iconIntermediateType{Image: v}}, t.icon...)

}

// RemoveIconImage deletes the value from the specified index
func (t *Accept) RemoveIconImage(index int) {
	copy(t.icon[index:], t.icon[index+1:])
	t.icon[len(t.icon)-1] = nil
	t.icon = t.icon[:len(t.icon)-1]

}

// IsIconLink determines whether the call to GetIconLink is safe for the specified index
func (t *Accept) IsIconLink(index int) (ok bool) {
	return t.icon[index].Link != nil

}

// GetIconLink returns the value safely if IsIconLink returned true for the specified index
func (t *Accept) GetIconLink(index int) (v LinkType) {
	return t.icon[index].Link

}

// AppendIconLink adds to the back of icon a LinkType type
func (t *Accept) AppendIconLink(v LinkType) {
	t.icon = append(t.icon, &iconIntermediateType{Link: v})

}

// PrependIconLink adds to the front of icon a LinkType type
func (t *Accept) PrependIconLink(v LinkType) {
	t.icon = append([]*iconIntermediateType{&iconIntermediateType{Link: v}}, t.icon...)

}

// RemoveIconLink deletes the value from the specified index
func (t *Accept) RemoveIconLink(index int) {
	copy(t.icon[index:], t.icon[index+1:])
	t.icon[len(t.icon)-1] = nil
	t.icon = t.icon[:len(t.icon)-1]

}

// IsIconIRI determines whether the call to GetIconIRI is safe for the specified index
func (t *Accept) IsIconIRI(index int) (ok bool) {
	return t.icon[index].IRI != nil

}

// GetIconIRI returns the value safely if IsIconIRI returned true for the specified index
func (t *Accept) GetIconIRI(index int) (v *url.URL) {
	return t.icon[index].IRI

}

// AppendIconIRI adds to the back of icon a *url.URL type
func (t *Accept) AppendIconIRI(v *url.URL) {
	t.icon = append(t.icon, &iconIntermediateType{IRI: v})

}

// PrependIconIRI adds to the front of icon a *url.URL type
func (t *Accept) PrependIconIRI(v *url.URL) {
	t.icon = append([]*iconIntermediateType{&iconIntermediateType{IRI: v}}, t.icon...)

}

// RemoveIconIRI deletes the value from the specified index
func (t *Accept) RemoveIconIRI(index int) {
	copy(t.icon[index:], t.icon[index+1:])
	t.icon[len(t.icon)-1] = nil
	t.icon = t.icon[:len(t.icon)-1]

}

// HasUnknownIcon determines whether the call to GetUnknownIcon is safe
func (t *Accept) HasUnknownIcon() (ok bool) {
	return t.icon != nil && t.icon[0].unknown_ != nil

}

// GetUnknownIcon returns the unknown value for icon
func (t *Accept) GetUnknownIcon() (v interface{}) {
	return t.icon[0].unknown_

}

// SetUnknownIcon sets the unknown value of icon
func (t *Accept) SetUnknownIcon(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &iconIntermediateType{}
	tmp.unknown_ = i
	t.icon = append(t.icon, tmp)

}

// HasId determines whether the call to GetId is safe
func (t *Accept) HasId() (ok bool) {
	return t.id != nil

}

// GetId returns the value for id
func (t *Accept) GetId() (v *url.URL) {
	return t.id

}

// SetId sets the value of id
func (t *Accept) SetId(v *url.URL) {
	t.id = v

}

// HasUnknownId determines whether the call to GetUnknownId is safe
func (t *Accept) HasUnknownId() (ok bool) {
	return t.unknown_ != nil && t.unknown_["id"] != nil

}

// GetUnknownId returns the unknown value for id
func (t *Accept) GetUnknownId() (v interface{}) {
	return t.unknown_["id"]

}

// SetUnknownId sets the unknown value of id
func (t *Accept) SetUnknownId(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["id"] = i

}

// ImageLen determines the number of elements able to be used for the IsImageImage, GetImageImage, and RemoveImageImage functions
func (t *Accept) ImageLen() (l int) {
	return len(t.image)

}

// IsImageImage determines whether the call to GetImageImage is safe for the specified index
func (t *Accept) IsImageImage(index int) (ok bool) {
	return t.image[index].Image != nil

}

// GetImageImage returns the value safely if IsImageImage returned true for the specified index
func (t *Accept) GetImageImage(index int) (v ImageType) {
	return t.image[index].Image

}

// AppendImageImage adds to the back of image a ImageType type
func (t *Accept) AppendImageImage(v ImageType) {
	t.image = append(t.image, &imageIntermediateType{Image: v})

}

// PrependImageImage adds to the front of image a ImageType type
func (t *Accept) PrependImageImage(v ImageType) {
	t.image = append([]*imageIntermediateType{&imageIntermediateType{Image: v}}, t.image...)

}

// RemoveImageImage deletes the value from the specified index
func (t *Accept) RemoveImageImage(index int) {
	copy(t.image[index:], t.image[index+1:])
	t.image[len(t.image)-1] = nil
	t.image = t.image[:len(t.image)-1]

}

// IsImageLink determines whether the call to GetImageLink is safe for the specified index
func (t *Accept) IsImageLink(index int) (ok bool) {
	return t.image[index].Link != nil

}

// GetImageLink returns the value safely if IsImageLink returned true for the specified index
func (t *Accept) GetImageLink(index int) (v LinkType) {
	return t.image[index].Link

}

// AppendImageLink adds to the back of image a LinkType type
func (t *Accept) AppendImageLink(v LinkType) {
	t.image = append(t.image, &imageIntermediateType{Link: v})

}

// PrependImageLink adds to the front of image a LinkType type
func (t *Accept) PrependImageLink(v LinkType) {
	t.image = append([]*imageIntermediateType{&imageIntermediateType{Link: v}}, t.image...)

}

// RemoveImageLink deletes the value from the specified index
func (t *Accept) RemoveImageLink(index int) {
	copy(t.image[index:], t.image[index+1:])
	t.image[len(t.image)-1] = nil
	t.image = t.image[:len(t.image)-1]

}

// IsImageIRI determines whether the call to GetImageIRI is safe for the specified index
func (t *Accept) IsImageIRI(index int) (ok bool) {
	return t.image[index].IRI != nil

}

// GetImageIRI returns the value safely if IsImageIRI returned true for the specified index
func (t *Accept) GetImageIRI(index int) (v *url.URL) {
	return t.image[index].IRI

}

// AppendImageIRI adds to the back of image a *url.URL type
func (t *Accept) AppendImageIRI(v *url.URL) {
	t.image = append(t.image, &imageIntermediateType{IRI: v})

}

// PrependImageIRI adds to the front of image a *url.URL type
func (t *Accept) PrependImageIRI(v *url.URL) {
	t.image = append([]*imageIntermediateType{&imageIntermediateType{IRI: v}}, t.image...)

}

// RemoveImageIRI deletes the value from the specified index
func (t *Accept) RemoveImageIRI(index int) {
	copy(t.image[index:], t.image[index+1:])
	t.image[len(t.image)-1] = nil
	t.image = t.image[:len(t.image)-1]

}

// HasUnknownImage determines whether the call to GetUnknownImage is safe
func (t *Accept) HasUnknownImage() (ok bool) {
	return t.image != nil && t.image[0].unknown_ != nil

}

// GetUnknownImage returns the unknown value for image
func (t *Accept) GetUnknownImage() (v interface{}) {
	return t.image[0].unknown_

}

// SetUnknownImage sets the unknown value of image
func (t *Accept) SetUnknownImage(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &imageIntermediateType{}
	tmp.unknown_ = i
	t.image = append(t.image, tmp)

}

// InReplyToLen determines the number of elements able to be used for the IsInReplyToObject, GetInReplyToObject, and RemoveInReplyToObject functions
func (t *Accept) InReplyToLen() (l int) {
	return len(t.inReplyTo)

}

// IsInReplyToObject determines whether the call to GetInReplyToObject is safe for the specified index
func (t *Accept) IsInReplyToObject(index int) (ok bool) {
	return t.inReplyTo[index].Object != nil

}

// GetInReplyToObject returns the value safely if IsInReplyToObject returned true for the specified index
func (t *Accept) GetInReplyToObject(index int) (v ObjectType) {
	return t.inReplyTo[index].Object

}

// AppendInReplyToObject adds to the back of inReplyTo a ObjectType type
func (t *Accept) AppendInReplyToObject(v ObjectType) {
	t.inReplyTo = append(t.inReplyTo, &inReplyToIntermediateType{Object: v})

}

// PrependInReplyToObject adds to the front of inReplyTo a ObjectType type
func (t *Accept) PrependInReplyToObject(v ObjectType) {
	t.inReplyTo = append([]*inReplyToIntermediateType{&inReplyToIntermediateType{Object: v}}, t.inReplyTo...)

}

// RemoveInReplyToObject deletes the value from the specified index
func (t *Accept) RemoveInReplyToObject(index int) {
	copy(t.inReplyTo[index:], t.inReplyTo[index+1:])
	t.inReplyTo[len(t.inReplyTo)-1] = nil
	t.inReplyTo = t.inReplyTo[:len(t.inReplyTo)-1]

}

// IsInReplyToLink determines whether the call to GetInReplyToLink is safe for the specified index
func (t *Accept) IsInReplyToLink(index int) (ok bool) {
	return t.inReplyTo[index].Link != nil

}

// GetInReplyToLink returns the value safely if IsInReplyToLink returned true for the specified index
func (t *Accept) GetInReplyToLink(index int) (v LinkType) {
	return t.inReplyTo[index].Link

}

// AppendInReplyToLink adds to the back of inReplyTo a LinkType type
func (t *Accept) AppendInReplyToLink(v LinkType) {
	t.inReplyTo = append(t.inReplyTo, &inReplyToIntermediateType{Link: v})

}

// PrependInReplyToLink adds to the front of inReplyTo a LinkType type
func (t *Accept) PrependInReplyToLink(v LinkType) {
	t.inReplyTo = append([]*inReplyToIntermediateType{&inReplyToIntermediateType{Link: v}}, t.inReplyTo...)

}

// RemoveInReplyToLink deletes the value from the specified index
func (t *Accept) RemoveInReplyToLink(index int) {
	copy(t.inReplyTo[index:], t.inReplyTo[index+1:])
	t.inReplyTo[len(t.inReplyTo)-1] = nil
	t.inReplyTo = t.inReplyTo[:len(t.inReplyTo)-1]

}

// IsInReplyToIRI determines whether the call to GetInReplyToIRI is safe for the specified index
func (t *Accept) IsInReplyToIRI(index int) (ok bool) {
	return t.inReplyTo[index].IRI != nil

}

// GetInReplyToIRI returns the value safely if IsInReplyToIRI returned true for the specified index
func (t *Accept) GetInReplyToIRI(index int) (v *url.URL) {
	return t.inReplyTo[index].IRI

}

// AppendInReplyToIRI adds to the back of inReplyTo a *url.URL type
func (t *Accept) AppendInReplyToIRI(v *url.URL) {
	t.inReplyTo = append(t.inReplyTo, &inReplyToIntermediateType{IRI: v})

}

// PrependInReplyToIRI adds to the front of inReplyTo a *url.URL type
func (t *Accept) PrependInReplyToIRI(v *url.URL) {
	t.inReplyTo = append([]*inReplyToIntermediateType{&inReplyToIntermediateType{IRI: v}}, t.inReplyTo...)

}

// RemoveInReplyToIRI deletes the value from the specified index
func (t *Accept) RemoveInReplyToIRI(index int) {
	copy(t.inReplyTo[index:], t.inReplyTo[index+1:])
	t.inReplyTo[len(t.inReplyTo)-1] = nil
	t.inReplyTo = t.inReplyTo[:len(t.inReplyTo)-1]

}

// HasUnknownInReplyTo determines whether the call to GetUnknownInReplyTo is safe
func (t *Accept) HasUnknownInReplyTo() (ok bool) {
	return t.inReplyTo != nil && t.inReplyTo[0].unknown_ != nil

}

// GetUnknownInReplyTo returns the unknown value for inReplyTo
func (t *Accept) GetUnknownInReplyTo() (v interface{}) {
	return t.inReplyTo[0].unknown_

}

// SetUnknownInReplyTo sets the unknown value of inReplyTo
func (t *Accept) SetUnknownInReplyTo(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &inReplyToIntermediateType{}
	tmp.unknown_ = i
	t.inReplyTo = append(t.inReplyTo, tmp)

}

// LocationLen determines the number of elements able to be used for the IsLocationObject, GetLocationObject, and RemoveLocationObject functions
func (t *Accept) LocationLen() (l int) {
	return len(t.location)

}

// IsLocationObject determines whether the call to GetLocationObject is safe for the specified index
func (t *Accept) IsLocationObject(index int) (ok bool) {
	return t.location[index].Object != nil

}

// GetLocationObject returns the value safely if IsLocationObject returned true for the specified index
func (t *Accept) GetLocationObject(index int) (v ObjectType) {
	return t.location[index].Object

}

// AppendLocationObject adds to the back of location a ObjectType type
func (t *Accept) AppendLocationObject(v ObjectType) {
	t.location = append(t.location, &locationIntermediateType{Object: v})

}

// PrependLocationObject adds to the front of location a ObjectType type
func (t *Accept) PrependLocationObject(v ObjectType) {
	t.location = append([]*locationIntermediateType{&locationIntermediateType{Object: v}}, t.location...)

}

// RemoveLocationObject deletes the value from the specified index
func (t *Accept) RemoveLocationObject(index int) {
	copy(t.location[index:], t.location[index+1:])
	t.location[len(t.location)-1] = nil
	t.location = t.location[:len(t.location)-1]

}

// IsLocationLink determines whether the call to GetLocationLink is safe for the specified index
func (t *Accept) IsLocationLink(index int) (ok bool) {
	return t.location[index].Link != nil

}

// GetLocationLink returns the value safely if IsLocationLink returned true for the specified index
func (t *Accept) GetLocationLink(index int) (v LinkType) {
	return t.location[index].Link

}

// AppendLocationLink adds to the back of location a LinkType type
func (t *Accept) AppendLocationLink(v LinkType) {
	t.location = append(t.location, &locationIntermediateType{Link: v})

}

// PrependLocationLink adds to the front of location a LinkType type
func (t *Accept) PrependLocationLink(v LinkType) {
	t.location = append([]*locationIntermediateType{&locationIntermediateType{Link: v}}, t.location...)

}

// RemoveLocationLink deletes the value from the specified index
func (t *Accept) RemoveLocationLink(index int) {
	copy(t.location[index:], t.location[index+1:])
	t.location[len(t.location)-1] = nil
	t.location = t.location[:len(t.location)-1]

}

// IsLocationIRI determines whether the call to GetLocationIRI is safe for the specified index
func (t *Accept) IsLocationIRI(index int) (ok bool) {
	return t.location[index].IRI != nil

}

// GetLocationIRI returns the value safely if IsLocationIRI returned true for the specified index
func (t *Accept) GetLocationIRI(index int) (v *url.URL) {
	return t.location[index].IRI

}

// AppendLocationIRI adds to the back of location a *url.URL type
func (t *Accept) AppendLocationIRI(v *url.URL) {
	t.location = append(t.location, &locationIntermediateType{IRI: v})

}

// PrependLocationIRI adds to the front of location a *url.URL type
func (t *Accept) PrependLocationIRI(v *url.URL) {
	t.location = append([]*locationIntermediateType{&locationIntermediateType{IRI: v}}, t.location...)

}

// RemoveLocationIRI deletes the value from the specified index
func (t *Accept) RemoveLocationIRI(index int) {
	copy(t.location[index:], t.location[index+1:])
	t.location[len(t.location)-1] = nil
	t.location = t.location[:len(t.location)-1]

}

// HasUnknownLocation determines whether the call to GetUnknownLocation is safe
func (t *Accept) HasUnknownLocation() (ok bool) {
	return t.location != nil && t.location[0].unknown_ != nil

}

// GetUnknownLocation returns the unknown value for location
func (t *Accept) GetUnknownLocation() (v interface{}) {
	return t.location[0].unknown_

}

// SetUnknownLocation sets the unknown value of location
func (t *Accept) SetUnknownLocation(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &locationIntermediateType{}
	tmp.unknown_ = i
	t.location = append(t.location, tmp)

}

// PreviewLen determines the number of elements able to be used for the IsPreviewObject, GetPreviewObject, and RemovePreviewObject functions
func (t *Accept) PreviewLen() (l int) {
	return len(t.preview)

}

// IsPreviewObject determines whether the call to GetPreviewObject is safe for the specified index
func (t *Accept) IsPreviewObject(index int) (ok bool) {
	return t.preview[index].Object != nil

}

// GetPreviewObject returns the value safely if IsPreviewObject returned true for the specified index
func (t *Accept) GetPreviewObject(index int) (v ObjectType) {
	return t.preview[index].Object

}

// AppendPreviewObject adds to the back of preview a ObjectType type
func (t *Accept) AppendPreviewObject(v ObjectType) {
	t.preview = append(t.preview, &previewIntermediateType{Object: v})

}

// PrependPreviewObject adds to the front of preview a ObjectType type
func (t *Accept) PrependPreviewObject(v ObjectType) {
	t.preview = append([]*previewIntermediateType{&previewIntermediateType{Object: v}}, t.preview...)

}

// RemovePreviewObject deletes the value from the specified index
func (t *Accept) RemovePreviewObject(index int) {
	copy(t.preview[index:], t.preview[index+1:])
	t.preview[len(t.preview)-1] = nil
	t.preview = t.preview[:len(t.preview)-1]

}

// IsPreviewLink determines whether the call to GetPreviewLink is safe for the specified index
func (t *Accept) IsPreviewLink(index int) (ok bool) {
	return t.preview[index].Link != nil

}

// GetPreviewLink returns the value safely if IsPreviewLink returned true for the specified index
func (t *Accept) GetPreviewLink(index int) (v LinkType) {
	return t.preview[index].Link

}

// AppendPreviewLink adds to the back of preview a LinkType type
func (t *Accept) AppendPreviewLink(v LinkType) {
	t.preview = append(t.preview, &previewIntermediateType{Link: v})

}

// PrependPreviewLink adds to the front of preview a LinkType type
func (t *Accept) PrependPreviewLink(v LinkType) {
	t.preview = append([]*previewIntermediateType{&previewIntermediateType{Link: v}}, t.preview...)

}

// RemovePreviewLink deletes the value from the specified index
func (t *Accept) RemovePreviewLink(index int) {
	copy(t.preview[index:], t.preview[index+1:])
	t.preview[len(t.preview)-1] = nil
	t.preview = t.preview[:len(t.preview)-1]

}

// IsPreviewIRI determines whether the call to GetPreviewIRI is safe for the specified index
func (t *Accept) IsPreviewIRI(index int) (ok bool) {
	return t.preview[index].IRI != nil

}

// GetPreviewIRI returns the value safely if IsPreviewIRI returned true for the specified index
func (t *Accept) GetPreviewIRI(index int) (v *url.URL) {
	return t.preview[index].IRI

}

// AppendPreviewIRI adds to the back of preview a *url.URL type
func (t *Accept) AppendPreviewIRI(v *url.URL) {
	t.preview = append(t.preview, &previewIntermediateType{IRI: v})

}

// PrependPreviewIRI adds to the front of preview a *url.URL type
func (t *Accept) PrependPreviewIRI(v *url.URL) {
	t.preview = append([]*previewIntermediateType{&previewIntermediateType{IRI: v}}, t.preview...)

}

// RemovePreviewIRI deletes the value from the specified index
func (t *Accept) RemovePreviewIRI(index int) {
	copy(t.preview[index:], t.preview[index+1:])
	t.preview[len(t.preview)-1] = nil
	t.preview = t.preview[:len(t.preview)-1]

}

// HasUnknownPreview determines whether the call to GetUnknownPreview is safe
func (t *Accept) HasUnknownPreview() (ok bool) {
	return t.preview != nil && t.preview[0].unknown_ != nil

}

// GetUnknownPreview returns the unknown value for preview
func (t *Accept) GetUnknownPreview() (v interface{}) {
	return t.preview[0].unknown_

}

// SetUnknownPreview sets the unknown value of preview
func (t *Accept) SetUnknownPreview(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &previewIntermediateType{}
	tmp.unknown_ = i
	t.preview = append(t.preview, tmp)

}

// IsPublished determines whether the call to GetPublished is safe
func (t *Accept) IsPublished() (ok bool) {
	return t.published != nil && t.published.dateTime != nil

}

// GetPublished returns the value safely if IsPublished returned true
func (t *Accept) GetPublished() (v time.Time) {
	return *t.published.dateTime

}

// SetPublished sets the value of published to be of time.Time type
func (t *Accept) SetPublished(v time.Time) {
	t.published = &publishedIntermediateType{dateTime: &v}

}

// IsPublishedIRI determines whether the call to GetPublishedIRI is safe
func (t *Accept) IsPublishedIRI() (ok bool) {
	return t.published != nil && t.published.IRI != nil

}

// GetPublishedIRI returns the value safely if IsPublishedIRI returned true
func (t *Accept) GetPublishedIRI() (v *url.URL) {
	return t.published.IRI

}

// SetPublishedIRI sets the value of published to be of *url.URL type
func (t *Accept) SetPublishedIRI(v *url.URL) {
	t.published = &publishedIntermediateType{IRI: v}

}

// HasUnknownPublished determines whether the call to GetUnknownPublished is safe
func (t *Accept) HasUnknownPublished() (ok bool) {
	return t.published != nil && t.published.unknown_ != nil

}

// GetUnknownPublished returns the unknown value for published
func (t *Accept) GetUnknownPublished() (v interface{}) {
	return t.published.unknown_

}

// SetUnknownPublished sets the unknown value of published
func (t *Accept) SetUnknownPublished(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &publishedIntermediateType{}
	tmp.unknown_ = i
	t.published = tmp

}

// IsReplies determines whether the call to GetReplies is safe
func (t *Accept) IsReplies() (ok bool) {
	return t.replies != nil && t.replies.Collection != nil

}

// GetReplies returns the value safely if IsReplies returned true
func (t *Accept) GetReplies() (v CollectionType) {
	return t.replies.Collection

}

// SetReplies sets the value of replies to be of CollectionType type
func (t *Accept) SetReplies(v CollectionType) {
	t.replies = &repliesIntermediateType{Collection: v}

}

// IsRepliesIRI determines whether the call to GetRepliesIRI is safe
func (t *Accept) IsRepliesIRI() (ok bool) {
	return t.replies != nil && t.replies.IRI != nil

}

// GetRepliesIRI returns the value safely if IsRepliesIRI returned true
func (t *Accept) GetRepliesIRI() (v *url.URL) {
	return t.replies.IRI

}

// SetRepliesIRI sets the value of replies to be of *url.URL type
func (t *Accept) SetRepliesIRI(v *url.URL) {
	t.replies = &repliesIntermediateType{IRI: v}

}

// HasUnknownReplies determines whether the call to GetUnknownReplies is safe
func (t *Accept) HasUnknownReplies() (ok bool) {
	return t.replies != nil && t.replies.unknown_ != nil

}

// GetUnknownReplies returns the unknown value for replies
func (t *Accept) GetUnknownReplies() (v interface{}) {
	return t.replies.unknown_

}

// SetUnknownReplies sets the unknown value of replies
func (t *Accept) SetUnknownReplies(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &repliesIntermediateType{}
	tmp.unknown_ = i
	t.replies = tmp

}

// IsStartTime determines whether the call to GetStartTime is safe
func (t *Accept) IsStartTime() (ok bool) {
	return t.startTime != nil && t.startTime.dateTime != nil

}

// GetStartTime returns the value safely if IsStartTime returned true
func (t *Accept) GetStartTime() (v time.Time) {
	return *t.startTime.dateTime

}

// SetStartTime sets the value of startTime to be of time.Time type
func (t *Accept) SetStartTime(v time.Time) {
	t.startTime = &startTimeIntermediateType{dateTime: &v}

}

// IsStartTimeIRI determines whether the call to GetStartTimeIRI is safe
func (t *Accept) IsStartTimeIRI() (ok bool) {
	return t.startTime != nil && t.startTime.IRI != nil

}

// GetStartTimeIRI returns the value safely if IsStartTimeIRI returned true
func (t *Accept) GetStartTimeIRI() (v *url.URL) {
	return t.startTime.IRI

}

// SetStartTimeIRI sets the value of startTime to be of *url.URL type
func (t *Accept) SetStartTimeIRI(v *url.URL) {
	t.startTime = &startTimeIntermediateType{IRI: v}

}

// HasUnknownStartTime determines whether the call to GetUnknownStartTime is safe
func (t *Accept) HasUnknownStartTime() (ok bool) {
	return t.startTime != nil && t.startTime.unknown_ != nil

}

// GetUnknownStartTime returns the unknown value for startTime
func (t *Accept) GetUnknownStartTime() (v interface{}) {
	return t.startTime.unknown_

}

// SetUnknownStartTime sets the unknown value of startTime
func (t *Accept) SetUnknownStartTime(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &startTimeIntermediateType{}
	tmp.unknown_ = i
	t.startTime = tmp

}

// SummaryLen determines the number of elements able to be used for the IsSummaryString, GetSummaryString, and RemoveSummaryString functions
func (t *Accept) SummaryLen() (l int) {
	return len(t.summary)

}

// IsSummaryString determines whether the call to GetSummaryString is safe for the specified index
func (t *Accept) IsSummaryString(index int) (ok bool) {
	return t.summary[index].stringName != nil

}

// GetSummaryString returns the value safely if IsSummaryString returned true for the specified index
func (t *Accept) GetSummaryString(index int) (v string) {
	return *t.summary[index].stringName

}

// AppendSummaryString adds to the back of summary a string type
func (t *Accept) AppendSummaryString(v string) {
	t.summary = append(t.summary, &summaryIntermediateType{stringName: &v})

}

// PrependSummaryString adds to the front of summary a string type
func (t *Accept) PrependSummaryString(v string) {
	t.summary = append([]*summaryIntermediateType{&summaryIntermediateType{stringName: &v}}, t.summary...)

}

// RemoveSummaryString deletes the value from the specified index
func (t *Accept) RemoveSummaryString(index int) {
	copy(t.summary[index:], t.summary[index+1:])
	t.summary[len(t.summary)-1] = nil
	t.summary = t.summary[:len(t.summary)-1]

}

// IsSummaryLangString determines whether the call to GetSummaryLangString is safe for the specified index
func (t *Accept) IsSummaryLangString(index int) (ok bool) {
	return t.summary[index].langString != nil

}

// GetSummaryLangString returns the value safely if IsSummaryLangString returned true for the specified index
func (t *Accept) GetSummaryLangString(index int) (v string) {
	return *t.summary[index].langString

}

// AppendSummaryLangString adds to the back of summary a string type
func (t *Accept) AppendSummaryLangString(v string) {
	t.summary = append(t.summary, &summaryIntermediateType{langString: &v})

}

// PrependSummaryLangString adds to the front of summary a string type
func (t *Accept) PrependSummaryLangString(v string) {
	t.summary = append([]*summaryIntermediateType{&summaryIntermediateType{langString: &v}}, t.summary...)

}

// RemoveSummaryLangString deletes the value from the specified index
func (t *Accept) RemoveSummaryLangString(index int) {
	copy(t.summary[index:], t.summary[index+1:])
	t.summary[len(t.summary)-1] = nil
	t.summary = t.summary[:len(t.summary)-1]

}

// IsSummaryIRI determines whether the call to GetSummaryIRI is safe for the specified index
func (t *Accept) IsSummaryIRI(index int) (ok bool) {
	return t.summary[index].IRI != nil

}

// GetSummaryIRI returns the value safely if IsSummaryIRI returned true for the specified index
func (t *Accept) GetSummaryIRI(index int) (v *url.URL) {
	return t.summary[index].IRI

}

// AppendSummaryIRI adds to the back of summary a *url.URL type
func (t *Accept) AppendSummaryIRI(v *url.URL) {
	t.summary = append(t.summary, &summaryIntermediateType{IRI: v})

}

// PrependSummaryIRI adds to the front of summary a *url.URL type
func (t *Accept) PrependSummaryIRI(v *url.URL) {
	t.summary = append([]*summaryIntermediateType{&summaryIntermediateType{IRI: v}}, t.summary...)

}

// RemoveSummaryIRI deletes the value from the specified index
func (t *Accept) RemoveSummaryIRI(index int) {
	copy(t.summary[index:], t.summary[index+1:])
	t.summary[len(t.summary)-1] = nil
	t.summary = t.summary[:len(t.summary)-1]

}

// HasUnknownSummary determines whether the call to GetUnknownSummary is safe
func (t *Accept) HasUnknownSummary() (ok bool) {
	return t.summary != nil && t.summary[0].unknown_ != nil

}

// GetUnknownSummary returns the unknown value for summary
func (t *Accept) GetUnknownSummary() (v interface{}) {
	return t.summary[0].unknown_

}

// SetUnknownSummary sets the unknown value of summary
func (t *Accept) SetUnknownSummary(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &summaryIntermediateType{}
	tmp.unknown_ = i
	t.summary = append(t.summary, tmp)

}

// SummaryMapLanguages returns all languages for this property's language mapping, or nil if there are none.
func (t *Accept) SummaryMapLanguages() (l []string) {
	if t.summaryMap == nil || len(t.summaryMap) == 0 {
		return nil
	}
	for k := range t.summaryMap {
		l = append(l, k)
	}
	return

}

// GetSummaryMap retrieves the value of the property for the specified language, or an empty string if it does not exist
func (t *Accept) GetSummaryMap(l string) (v string) {
	if t.summaryMap == nil {
		return ""
	}
	ok := false
	v, ok = t.summaryMap[l]
	if !ok {
		return ""
	}
	return v

}

// SetSummaryMap sets the value of the property for the specified language
func (t *Accept) SetSummaryMap(l string, v string) {
	if t.summaryMap == nil {
		t.summaryMap = make(map[string]string)
	}
	t.summaryMap[l] = v

}

// TagLen determines the number of elements able to be used for the IsTagObject, GetTagObject, and RemoveTagObject functions
func (t *Accept) TagLen() (l int) {
	return len(t.tag)

}

// IsTagObject determines whether the call to GetTagObject is safe for the specified index
func (t *Accept) IsTagObject(index int) (ok bool) {
	return t.tag[index].Object != nil

}

// GetTagObject returns the value safely if IsTagObject returned true for the specified index
func (t *Accept) GetTagObject(index int) (v ObjectType) {
	return t.tag[index].Object

}

// AppendTagObject adds to the back of tag a ObjectType type
func (t *Accept) AppendTagObject(v ObjectType) {
	t.tag = append(t.tag, &tagIntermediateType{Object: v})

}

// PrependTagObject adds to the front of tag a ObjectType type
func (t *Accept) PrependTagObject(v ObjectType) {
	t.tag = append([]*tagIntermediateType{&tagIntermediateType{Object: v}}, t.tag...)

}

// RemoveTagObject deletes the value from the specified index
func (t *Accept) RemoveTagObject(index int) {
	copy(t.tag[index:], t.tag[index+1:])
	t.tag[len(t.tag)-1] = nil
	t.tag = t.tag[:len(t.tag)-1]

}

// IsTagLink determines whether the call to GetTagLink is safe for the specified index
func (t *Accept) IsTagLink(index int) (ok bool) {
	return t.tag[index].Link != nil

}

// GetTagLink returns the value safely if IsTagLink returned true for the specified index
func (t *Accept) GetTagLink(index int) (v LinkType) {
	return t.tag[index].Link

}

// AppendTagLink adds to the back of tag a LinkType type
func (t *Accept) AppendTagLink(v LinkType) {
	t.tag = append(t.tag, &tagIntermediateType{Link: v})

}

// PrependTagLink adds to the front of tag a LinkType type
func (t *Accept) PrependTagLink(v LinkType) {
	t.tag = append([]*tagIntermediateType{&tagIntermediateType{Link: v}}, t.tag...)

}

// RemoveTagLink deletes the value from the specified index
func (t *Accept) RemoveTagLink(index int) {
	copy(t.tag[index:], t.tag[index+1:])
	t.tag[len(t.tag)-1] = nil
	t.tag = t.tag[:len(t.tag)-1]

}

// IsTagIRI determines whether the call to GetTagIRI is safe for the specified index
func (t *Accept) IsTagIRI(index int) (ok bool) {
	return t.tag[index].IRI != nil

}

// GetTagIRI returns the value safely if IsTagIRI returned true for the specified index
func (t *Accept) GetTagIRI(index int) (v *url.URL) {
	return t.tag[index].IRI

}

// AppendTagIRI adds to the back of tag a *url.URL type
func (t *Accept) AppendTagIRI(v *url.URL) {
	t.tag = append(t.tag, &tagIntermediateType{IRI: v})

}

// PrependTagIRI adds to the front of tag a *url.URL type
func (t *Accept) PrependTagIRI(v *url.URL) {
	t.tag = append([]*tagIntermediateType{&tagIntermediateType{IRI: v}}, t.tag...)

}

// RemoveTagIRI deletes the value from the specified index
func (t *Accept) RemoveTagIRI(index int) {
	copy(t.tag[index:], t.tag[index+1:])
	t.tag[len(t.tag)-1] = nil
	t.tag = t.tag[:len(t.tag)-1]

}

// HasUnknownTag determines whether the call to GetUnknownTag is safe
func (t *Accept) HasUnknownTag() (ok bool) {
	return t.tag != nil && t.tag[0].unknown_ != nil

}

// GetUnknownTag returns the unknown value for tag
func (t *Accept) GetUnknownTag() (v interface{}) {
	return t.tag[0].unknown_

}

// SetUnknownTag sets the unknown value of tag
func (t *Accept) SetUnknownTag(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &tagIntermediateType{}
	tmp.unknown_ = i
	t.tag = append(t.tag, tmp)

}

// TypeLen determines the number of elements able to be used for the GetType and RemoveType functions
func (t *Accept) TypeLen() (l int) {
	return len(t.typeName)

}

// GetType returns the value for the specified index
func (t *Accept) GetType(index int) (v interface{}) {
	return t.typeName[index]

}

// AppendType adds a value to the back of type
func (t *Accept) AppendType(v interface{}) {
	t.typeName = append(t.typeName, v)

}

// PrependType adds a value to the front of type
func (t *Accept) PrependType(v interface{}) {
	t.typeName = append([]interface{}{v}, t.typeName...)

}

// RemoveType deletes the value from the specified index
func (t *Accept) RemoveType(index int) {
	copy(t.typeName[index:], t.typeName[index+1:])
	t.typeName[len(t.typeName)-1] = nil
	t.typeName = t.typeName[:len(t.typeName)-1]

}

// IsUpdated determines whether the call to GetUpdated is safe
func (t *Accept) IsUpdated() (ok bool) {
	return t.updated != nil && t.updated.dateTime != nil

}

// GetUpdated returns the value safely if IsUpdated returned true
func (t *Accept) GetUpdated() (v time.Time) {
	return *t.updated.dateTime

}

// SetUpdated sets the value of updated to be of time.Time type
func (t *Accept) SetUpdated(v time.Time) {
	t.updated = &updatedIntermediateType{dateTime: &v}

}

// IsUpdatedIRI determines whether the call to GetUpdatedIRI is safe
func (t *Accept) IsUpdatedIRI() (ok bool) {
	return t.updated != nil && t.updated.IRI != nil

}

// GetUpdatedIRI returns the value safely if IsUpdatedIRI returned true
func (t *Accept) GetUpdatedIRI() (v *url.URL) {
	return t.updated.IRI

}

// SetUpdatedIRI sets the value of updated to be of *url.URL type
func (t *Accept) SetUpdatedIRI(v *url.URL) {
	t.updated = &updatedIntermediateType{IRI: v}

}

// HasUnknownUpdated determines whether the call to GetUnknownUpdated is safe
func (t *Accept) HasUnknownUpdated() (ok bool) {
	return t.updated != nil && t.updated.unknown_ != nil

}

// GetUnknownUpdated returns the unknown value for updated
func (t *Accept) GetUnknownUpdated() (v interface{}) {
	return t.updated.unknown_

}

// SetUnknownUpdated sets the unknown value of updated
func (t *Accept) SetUnknownUpdated(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &updatedIntermediateType{}
	tmp.unknown_ = i
	t.updated = tmp

}

// UrlLen determines the number of elements able to be used for the IsUrlAnyURI, GetUrlAnyURI, and RemoveUrlAnyURI functions
func (t *Accept) UrlLen() (l int) {
	return len(t.url)

}

// IsUrlAnyURI determines whether the call to GetUrlAnyURI is safe for the specified index
func (t *Accept) IsUrlAnyURI(index int) (ok bool) {
	return t.url[index].anyURI != nil

}

// GetUrlAnyURI returns the value safely if IsUrlAnyURI returned true for the specified index
func (t *Accept) GetUrlAnyURI(index int) (v *url.URL) {
	return t.url[index].anyURI

}

// AppendUrlAnyURI adds to the back of url a *url.URL type
func (t *Accept) AppendUrlAnyURI(v *url.URL) {
	t.url = append(t.url, &urlIntermediateType{anyURI: v})

}

// PrependUrlAnyURI adds to the front of url a *url.URL type
func (t *Accept) PrependUrlAnyURI(v *url.URL) {
	t.url = append([]*urlIntermediateType{&urlIntermediateType{anyURI: v}}, t.url...)

}

// RemoveUrlAnyURI deletes the value from the specified index
func (t *Accept) RemoveUrlAnyURI(index int) {
	copy(t.url[index:], t.url[index+1:])
	t.url[len(t.url)-1] = nil
	t.url = t.url[:len(t.url)-1]

}

// IsUrlLink determines whether the call to GetUrlLink is safe for the specified index
func (t *Accept) IsUrlLink(index int) (ok bool) {
	return t.url[index].Link != nil

}

// GetUrlLink returns the value safely if IsUrlLink returned true for the specified index
func (t *Accept) GetUrlLink(index int) (v LinkType) {
	return t.url[index].Link

}

// AppendUrlLink adds to the back of url a LinkType type
func (t *Accept) AppendUrlLink(v LinkType) {
	t.url = append(t.url, &urlIntermediateType{Link: v})

}

// PrependUrlLink adds to the front of url a LinkType type
func (t *Accept) PrependUrlLink(v LinkType) {
	t.url = append([]*urlIntermediateType{&urlIntermediateType{Link: v}}, t.url...)

}

// RemoveUrlLink deletes the value from the specified index
func (t *Accept) RemoveUrlLink(index int) {
	copy(t.url[index:], t.url[index+1:])
	t.url[len(t.url)-1] = nil
	t.url = t.url[:len(t.url)-1]

}

// HasUnknownUrl determines whether the call to GetUnknownUrl is safe
func (t *Accept) HasUnknownUrl() (ok bool) {
	return t.url != nil && t.url[0].unknown_ != nil

}

// GetUnknownUrl returns the unknown value for url
func (t *Accept) GetUnknownUrl() (v interface{}) {
	return t.url[0].unknown_

}

// SetUnknownUrl sets the unknown value of url
func (t *Accept) SetUnknownUrl(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &urlIntermediateType{}
	tmp.unknown_ = i
	t.url = append(t.url, tmp)

}

// ToLen determines the number of elements able to be used for the IsToObject, GetToObject, and RemoveToObject functions
func (t *Accept) ToLen() (l int) {
	return len(t.to)

}

// IsToObject determines whether the call to GetToObject is safe for the specified index
func (t *Accept) IsToObject(index int) (ok bool) {
	return t.to[index].Object != nil

}

// GetToObject returns the value safely if IsToObject returned true for the specified index
func (t *Accept) GetToObject(index int) (v ObjectType) {
	return t.to[index].Object

}

// AppendToObject adds to the back of to a ObjectType type
func (t *Accept) AppendToObject(v ObjectType) {
	t.to = append(t.to, &toIntermediateType{Object: v})

}

// PrependToObject adds to the front of to a ObjectType type
func (t *Accept) PrependToObject(v ObjectType) {
	t.to = append([]*toIntermediateType{&toIntermediateType{Object: v}}, t.to...)

}

// RemoveToObject deletes the value from the specified index
func (t *Accept) RemoveToObject(index int) {
	copy(t.to[index:], t.to[index+1:])
	t.to[len(t.to)-1] = nil
	t.to = t.to[:len(t.to)-1]

}

// IsToLink determines whether the call to GetToLink is safe for the specified index
func (t *Accept) IsToLink(index int) (ok bool) {
	return t.to[index].Link != nil

}

// GetToLink returns the value safely if IsToLink returned true for the specified index
func (t *Accept) GetToLink(index int) (v LinkType) {
	return t.to[index].Link

}

// AppendToLink adds to the back of to a LinkType type
func (t *Accept) AppendToLink(v LinkType) {
	t.to = append(t.to, &toIntermediateType{Link: v})

}

// PrependToLink adds to the front of to a LinkType type
func (t *Accept) PrependToLink(v LinkType) {
	t.to = append([]*toIntermediateType{&toIntermediateType{Link: v}}, t.to...)

}

// RemoveToLink deletes the value from the specified index
func (t *Accept) RemoveToLink(index int) {
	copy(t.to[index:], t.to[index+1:])
	t.to[len(t.to)-1] = nil
	t.to = t.to[:len(t.to)-1]

}

// IsToIRI determines whether the call to GetToIRI is safe for the specified index
func (t *Accept) IsToIRI(index int) (ok bool) {
	return t.to[index].IRI != nil

}

// GetToIRI returns the value safely if IsToIRI returned true for the specified index
func (t *Accept) GetToIRI(index int) (v *url.URL) {
	return t.to[index].IRI

}

// AppendToIRI adds to the back of to a *url.URL type
func (t *Accept) AppendToIRI(v *url.URL) {
	t.to = append(t.to, &toIntermediateType{IRI: v})

}

// PrependToIRI adds to the front of to a *url.URL type
func (t *Accept) PrependToIRI(v *url.URL) {
	t.to = append([]*toIntermediateType{&toIntermediateType{IRI: v}}, t.to...)

}

// RemoveToIRI deletes the value from the specified index
func (t *Accept) RemoveToIRI(index int) {
	copy(t.to[index:], t.to[index+1:])
	t.to[len(t.to)-1] = nil
	t.to = t.to[:len(t.to)-1]

}

// HasUnknownTo determines whether the call to GetUnknownTo is safe
func (t *Accept) HasUnknownTo() (ok bool) {
	return t.to != nil && t.to[0].unknown_ != nil

}

// GetUnknownTo returns the unknown value for to
func (t *Accept) GetUnknownTo() (v interface{}) {
	return t.to[0].unknown_

}

// SetUnknownTo sets the unknown value of to
func (t *Accept) SetUnknownTo(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &toIntermediateType{}
	tmp.unknown_ = i
	t.to = append(t.to, tmp)

}

// BtoLen determines the number of elements able to be used for the IsBtoObject, GetBtoObject, and RemoveBtoObject functions
func (t *Accept) BtoLen() (l int) {
	return len(t.bto)

}

// IsBtoObject determines whether the call to GetBtoObject is safe for the specified index
func (t *Accept) IsBtoObject(index int) (ok bool) {
	return t.bto[index].Object != nil

}

// GetBtoObject returns the value safely if IsBtoObject returned true for the specified index
func (t *Accept) GetBtoObject(index int) (v ObjectType) {
	return t.bto[index].Object

}

// AppendBtoObject adds to the back of bto a ObjectType type
func (t *Accept) AppendBtoObject(v ObjectType) {
	t.bto = append(t.bto, &btoIntermediateType{Object: v})

}

// PrependBtoObject adds to the front of bto a ObjectType type
func (t *Accept) PrependBtoObject(v ObjectType) {
	t.bto = append([]*btoIntermediateType{&btoIntermediateType{Object: v}}, t.bto...)

}

// RemoveBtoObject deletes the value from the specified index
func (t *Accept) RemoveBtoObject(index int) {
	copy(t.bto[index:], t.bto[index+1:])
	t.bto[len(t.bto)-1] = nil
	t.bto = t.bto[:len(t.bto)-1]

}

// IsBtoLink determines whether the call to GetBtoLink is safe for the specified index
func (t *Accept) IsBtoLink(index int) (ok bool) {
	return t.bto[index].Link != nil

}

// GetBtoLink returns the value safely if IsBtoLink returned true for the specified index
func (t *Accept) GetBtoLink(index int) (v LinkType) {
	return t.bto[index].Link

}

// AppendBtoLink adds to the back of bto a LinkType type
func (t *Accept) AppendBtoLink(v LinkType) {
	t.bto = append(t.bto, &btoIntermediateType{Link: v})

}

// PrependBtoLink adds to the front of bto a LinkType type
func (t *Accept) PrependBtoLink(v LinkType) {
	t.bto = append([]*btoIntermediateType{&btoIntermediateType{Link: v}}, t.bto...)

}

// RemoveBtoLink deletes the value from the specified index
func (t *Accept) RemoveBtoLink(index int) {
	copy(t.bto[index:], t.bto[index+1:])
	t.bto[len(t.bto)-1] = nil
	t.bto = t.bto[:len(t.bto)-1]

}

// IsBtoIRI determines whether the call to GetBtoIRI is safe for the specified index
func (t *Accept) IsBtoIRI(index int) (ok bool) {
	return t.bto[index].IRI != nil

}

// GetBtoIRI returns the value safely if IsBtoIRI returned true for the specified index
func (t *Accept) GetBtoIRI(index int) (v *url.URL) {
	return t.bto[index].IRI

}

// AppendBtoIRI adds to the back of bto a *url.URL type
func (t *Accept) AppendBtoIRI(v *url.URL) {
	t.bto = append(t.bto, &btoIntermediateType{IRI: v})

}

// PrependBtoIRI adds to the front of bto a *url.URL type
func (t *Accept) PrependBtoIRI(v *url.URL) {
	t.bto = append([]*btoIntermediateType{&btoIntermediateType{IRI: v}}, t.bto...)

}

// RemoveBtoIRI deletes the value from the specified index
func (t *Accept) RemoveBtoIRI(index int) {
	copy(t.bto[index:], t.bto[index+1:])
	t.bto[len(t.bto)-1] = nil
	t.bto = t.bto[:len(t.bto)-1]

}

// HasUnknownBto determines whether the call to GetUnknownBto is safe
func (t *Accept) HasUnknownBto() (ok bool) {
	return t.bto != nil && t.bto[0].unknown_ != nil

}

// GetUnknownBto returns the unknown value for bto
func (t *Accept) GetUnknownBto() (v interface{}) {
	return t.bto[0].unknown_

}

// SetUnknownBto sets the unknown value of bto
func (t *Accept) SetUnknownBto(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &btoIntermediateType{}
	tmp.unknown_ = i
	t.bto = append(t.bto, tmp)

}

// CcLen determines the number of elements able to be used for the IsCcObject, GetCcObject, and RemoveCcObject functions
func (t *Accept) CcLen() (l int) {
	return len(t.cc)

}

// IsCcObject determines whether the call to GetCcObject is safe for the specified index
func (t *Accept) IsCcObject(index int) (ok bool) {
	return t.cc[index].Object != nil

}

// GetCcObject returns the value safely if IsCcObject returned true for the specified index
func (t *Accept) GetCcObject(index int) (v ObjectType) {
	return t.cc[index].Object

}

// AppendCcObject adds to the back of cc a ObjectType type
func (t *Accept) AppendCcObject(v ObjectType) {
	t.cc = append(t.cc, &ccIntermediateType{Object: v})

}

// PrependCcObject adds to the front of cc a ObjectType type
func (t *Accept) PrependCcObject(v ObjectType) {
	t.cc = append([]*ccIntermediateType{&ccIntermediateType{Object: v}}, t.cc...)

}

// RemoveCcObject deletes the value from the specified index
func (t *Accept) RemoveCcObject(index int) {
	copy(t.cc[index:], t.cc[index+1:])
	t.cc[len(t.cc)-1] = nil
	t.cc = t.cc[:len(t.cc)-1]

}

// IsCcLink determines whether the call to GetCcLink is safe for the specified index
func (t *Accept) IsCcLink(index int) (ok bool) {
	return t.cc[index].Link != nil

}

// GetCcLink returns the value safely if IsCcLink returned true for the specified index
func (t *Accept) GetCcLink(index int) (v LinkType) {
	return t.cc[index].Link

}

// AppendCcLink adds to the back of cc a LinkType type
func (t *Accept) AppendCcLink(v LinkType) {
	t.cc = append(t.cc, &ccIntermediateType{Link: v})

}

// PrependCcLink adds to the front of cc a LinkType type
func (t *Accept) PrependCcLink(v LinkType) {
	t.cc = append([]*ccIntermediateType{&ccIntermediateType{Link: v}}, t.cc...)

}

// RemoveCcLink deletes the value from the specified index
func (t *Accept) RemoveCcLink(index int) {
	copy(t.cc[index:], t.cc[index+1:])
	t.cc[len(t.cc)-1] = nil
	t.cc = t.cc[:len(t.cc)-1]

}

// IsCcIRI determines whether the call to GetCcIRI is safe for the specified index
func (t *Accept) IsCcIRI(index int) (ok bool) {
	return t.cc[index].IRI != nil

}

// GetCcIRI returns the value safely if IsCcIRI returned true for the specified index
func (t *Accept) GetCcIRI(index int) (v *url.URL) {
	return t.cc[index].IRI

}

// AppendCcIRI adds to the back of cc a *url.URL type
func (t *Accept) AppendCcIRI(v *url.URL) {
	t.cc = append(t.cc, &ccIntermediateType{IRI: v})

}

// PrependCcIRI adds to the front of cc a *url.URL type
func (t *Accept) PrependCcIRI(v *url.URL) {
	t.cc = append([]*ccIntermediateType{&ccIntermediateType{IRI: v}}, t.cc...)

}

// RemoveCcIRI deletes the value from the specified index
func (t *Accept) RemoveCcIRI(index int) {
	copy(t.cc[index:], t.cc[index+1:])
	t.cc[len(t.cc)-1] = nil
	t.cc = t.cc[:len(t.cc)-1]

}

// HasUnknownCc determines whether the call to GetUnknownCc is safe
func (t *Accept) HasUnknownCc() (ok bool) {
	return t.cc != nil && t.cc[0].unknown_ != nil

}

// GetUnknownCc returns the unknown value for cc
func (t *Accept) GetUnknownCc() (v interface{}) {
	return t.cc[0].unknown_

}

// SetUnknownCc sets the unknown value of cc
func (t *Accept) SetUnknownCc(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &ccIntermediateType{}
	tmp.unknown_ = i
	t.cc = append(t.cc, tmp)

}

// BccLen determines the number of elements able to be used for the IsBccObject, GetBccObject, and RemoveBccObject functions
func (t *Accept) BccLen() (l int) {
	return len(t.bcc)

}

// IsBccObject determines whether the call to GetBccObject is safe for the specified index
func (t *Accept) IsBccObject(index int) (ok bool) {
	return t.bcc[index].Object != nil

}

// GetBccObject returns the value safely if IsBccObject returned true for the specified index
func (t *Accept) GetBccObject(index int) (v ObjectType) {
	return t.bcc[index].Object

}

// AppendBccObject adds to the back of bcc a ObjectType type
func (t *Accept) AppendBccObject(v ObjectType) {
	t.bcc = append(t.bcc, &bccIntermediateType{Object: v})

}

// PrependBccObject adds to the front of bcc a ObjectType type
func (t *Accept) PrependBccObject(v ObjectType) {
	t.bcc = append([]*bccIntermediateType{&bccIntermediateType{Object: v}}, t.bcc...)

}

// RemoveBccObject deletes the value from the specified index
func (t *Accept) RemoveBccObject(index int) {
	copy(t.bcc[index:], t.bcc[index+1:])
	t.bcc[len(t.bcc)-1] = nil
	t.bcc = t.bcc[:len(t.bcc)-1]

}

// IsBccLink determines whether the call to GetBccLink is safe for the specified index
func (t *Accept) IsBccLink(index int) (ok bool) {
	return t.bcc[index].Link != nil

}

// GetBccLink returns the value safely if IsBccLink returned true for the specified index
func (t *Accept) GetBccLink(index int) (v LinkType) {
	return t.bcc[index].Link

}

// AppendBccLink adds to the back of bcc a LinkType type
func (t *Accept) AppendBccLink(v LinkType) {
	t.bcc = append(t.bcc, &bccIntermediateType{Link: v})

}

// PrependBccLink adds to the front of bcc a LinkType type
func (t *Accept) PrependBccLink(v LinkType) {
	t.bcc = append([]*bccIntermediateType{&bccIntermediateType{Link: v}}, t.bcc...)

}

// RemoveBccLink deletes the value from the specified index
func (t *Accept) RemoveBccLink(index int) {
	copy(t.bcc[index:], t.bcc[index+1:])
	t.bcc[len(t.bcc)-1] = nil
	t.bcc = t.bcc[:len(t.bcc)-1]

}

// IsBccIRI determines whether the call to GetBccIRI is safe for the specified index
func (t *Accept) IsBccIRI(index int) (ok bool) {
	return t.bcc[index].IRI != nil

}

// GetBccIRI returns the value safely if IsBccIRI returned true for the specified index
func (t *Accept) GetBccIRI(index int) (v *url.URL) {
	return t.bcc[index].IRI

}

// AppendBccIRI adds to the back of bcc a *url.URL type
func (t *Accept) AppendBccIRI(v *url.URL) {
	t.bcc = append(t.bcc, &bccIntermediateType{IRI: v})

}

// PrependBccIRI adds to the front of bcc a *url.URL type
func (t *Accept) PrependBccIRI(v *url.URL) {
	t.bcc = append([]*bccIntermediateType{&bccIntermediateType{IRI: v}}, t.bcc...)

}

// RemoveBccIRI deletes the value from the specified index
func (t *Accept) RemoveBccIRI(index int) {
	copy(t.bcc[index:], t.bcc[index+1:])
	t.bcc[len(t.bcc)-1] = nil
	t.bcc = t.bcc[:len(t.bcc)-1]

}

// HasUnknownBcc determines whether the call to GetUnknownBcc is safe
func (t *Accept) HasUnknownBcc() (ok bool) {
	return t.bcc != nil && t.bcc[0].unknown_ != nil

}

// GetUnknownBcc returns the unknown value for bcc
func (t *Accept) GetUnknownBcc() (v interface{}) {
	return t.bcc[0].unknown_

}

// SetUnknownBcc sets the unknown value of bcc
func (t *Accept) SetUnknownBcc(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &bccIntermediateType{}
	tmp.unknown_ = i
	t.bcc = append(t.bcc, tmp)

}

// IsMediaType determines whether the call to GetMediaType is safe
func (t *Accept) IsMediaType() (ok bool) {
	return t.mediaType != nil && t.mediaType.mimeMediaTypeValue != nil

}

// GetMediaType returns the value safely if IsMediaType returned true
func (t *Accept) GetMediaType() (v string) {
	return *t.mediaType.mimeMediaTypeValue

}

// SetMediaType sets the value of mediaType to be of string type
func (t *Accept) SetMediaType(v string) {
	t.mediaType = &mediaTypeIntermediateType{mimeMediaTypeValue: &v}

}

// IsMediaTypeIRI determines whether the call to GetMediaTypeIRI is safe
func (t *Accept) IsMediaTypeIRI() (ok bool) {
	return t.mediaType != nil && t.mediaType.IRI != nil

}

// GetMediaTypeIRI returns the value safely if IsMediaTypeIRI returned true
func (t *Accept) GetMediaTypeIRI() (v *url.URL) {
	return t.mediaType.IRI

}

// SetMediaTypeIRI sets the value of mediaType to be of *url.URL type
func (t *Accept) SetMediaTypeIRI(v *url.URL) {
	t.mediaType = &mediaTypeIntermediateType{IRI: v}

}

// HasUnknownMediaType determines whether the call to GetUnknownMediaType is safe
func (t *Accept) HasUnknownMediaType() (ok bool) {
	return t.mediaType != nil && t.mediaType.unknown_ != nil

}

// GetUnknownMediaType returns the unknown value for mediaType
func (t *Accept) GetUnknownMediaType() (v interface{}) {
	return t.mediaType.unknown_

}

// SetUnknownMediaType sets the unknown value of mediaType
func (t *Accept) SetUnknownMediaType(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &mediaTypeIntermediateType{}
	tmp.unknown_ = i
	t.mediaType = tmp

}

// IsDuration determines whether the call to GetDuration is safe
func (t *Accept) IsDuration() (ok bool) {
	return t.duration != nil && t.duration.duration != nil

}

// GetDuration returns the value safely if IsDuration returned true
func (t *Accept) GetDuration() (v time.Duration) {
	return *t.duration.duration

}

// SetDuration sets the value of duration to be of time.Duration type
func (t *Accept) SetDuration(v time.Duration) {
	t.duration = &durationIntermediateType{duration: &v}

}

// IsDurationIRI determines whether the call to GetDurationIRI is safe
func (t *Accept) IsDurationIRI() (ok bool) {
	return t.duration != nil && t.duration.IRI != nil

}

// GetDurationIRI returns the value safely if IsDurationIRI returned true
func (t *Accept) GetDurationIRI() (v *url.URL) {
	return t.duration.IRI

}

// SetDurationIRI sets the value of duration to be of *url.URL type
func (t *Accept) SetDurationIRI(v *url.URL) {
	t.duration = &durationIntermediateType{IRI: v}

}

// HasUnknownDuration determines whether the call to GetUnknownDuration is safe
func (t *Accept) HasUnknownDuration() (ok bool) {
	return t.duration != nil && t.duration.unknown_ != nil

}

// GetUnknownDuration returns the unknown value for duration
func (t *Accept) GetUnknownDuration() (v interface{}) {
	return t.duration.unknown_

}

// SetUnknownDuration sets the unknown value of duration
func (t *Accept) SetUnknownDuration(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &durationIntermediateType{}
	tmp.unknown_ = i
	t.duration = tmp

}

// IsSource determines whether the call to GetSource is safe
func (t *Accept) IsSource() (ok bool) {
	return t.source != nil && t.source.Object != nil

}

// GetSource returns the value safely if IsSource returned true
func (t *Accept) GetSource() (v ObjectType) {
	return t.source.Object

}

// SetSource sets the value of source to be of ObjectType type
func (t *Accept) SetSource(v ObjectType) {
	t.source = &sourceIntermediateType{Object: v}

}

// IsSourceIRI determines whether the call to GetSourceIRI is safe
func (t *Accept) IsSourceIRI() (ok bool) {
	return t.source != nil && t.source.IRI != nil

}

// GetSourceIRI returns the value safely if IsSourceIRI returned true
func (t *Accept) GetSourceIRI() (v *url.URL) {
	return t.source.IRI

}

// SetSourceIRI sets the value of source to be of *url.URL type
func (t *Accept) SetSourceIRI(v *url.URL) {
	t.source = &sourceIntermediateType{IRI: v}

}

// HasUnknownSource determines whether the call to GetUnknownSource is safe
func (t *Accept) HasUnknownSource() (ok bool) {
	return t.source != nil && t.source.unknown_ != nil

}

// GetUnknownSource returns the unknown value for source
func (t *Accept) GetUnknownSource() (v interface{}) {
	return t.source.unknown_

}

// SetUnknownSource sets the unknown value of source
func (t *Accept) SetUnknownSource(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &sourceIntermediateType{}
	tmp.unknown_ = i
	t.source = tmp

}

// IsInboxOrderedCollection determines whether the call to GetInboxOrderedCollection is safe
func (t *Accept) IsInboxOrderedCollection() (ok bool) {
	return t.inbox != nil && t.inbox.OrderedCollection != nil

}

// GetInboxOrderedCollection returns the value safely if IsInboxOrderedCollection returned true
func (t *Accept) GetInboxOrderedCollection() (v OrderedCollectionType) {
	return t.inbox.OrderedCollection

}

// SetInboxOrderedCollection sets the value of inbox to be of OrderedCollectionType type
func (t *Accept) SetInboxOrderedCollection(v OrderedCollectionType) {
	t.inbox = &inboxIntermediateType{OrderedCollection: v}

}

// IsInboxAnyURI determines whether the call to GetInboxAnyURI is safe
func (t *Accept) IsInboxAnyURI() (ok bool) {
	return t.inbox != nil && t.inbox.anyURI != nil

}

// GetInboxAnyURI returns the value safely if IsInboxAnyURI returned true
func (t *Accept) GetInboxAnyURI() (v *url.URL) {
	return t.inbox.anyURI

}

// SetInboxAnyURI sets the value of inbox to be of *url.URL type
func (t *Accept) SetInboxAnyURI(v *url.URL) {
	t.inbox = &inboxIntermediateType{anyURI: v}

}

// HasUnknownInbox determines whether the call to GetUnknownInbox is safe
func (t *Accept) HasUnknownInbox() (ok bool) {
	return t.inbox != nil && t.inbox.unknown_ != nil

}

// GetUnknownInbox returns the unknown value for inbox
func (t *Accept) GetUnknownInbox() (v interface{}) {
	return t.inbox.unknown_

}

// SetUnknownInbox sets the unknown value of inbox
func (t *Accept) SetUnknownInbox(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &inboxIntermediateType{}
	tmp.unknown_ = i
	t.inbox = tmp

}

// IsOutboxOrderedCollection determines whether the call to GetOutboxOrderedCollection is safe
func (t *Accept) IsOutboxOrderedCollection() (ok bool) {
	return t.outbox != nil && t.outbox.OrderedCollection != nil

}

// GetOutboxOrderedCollection returns the value safely if IsOutboxOrderedCollection returned true
func (t *Accept) GetOutboxOrderedCollection() (v OrderedCollectionType) {
	return t.outbox.OrderedCollection

}

// SetOutboxOrderedCollection sets the value of outbox to be of OrderedCollectionType type
func (t *Accept) SetOutboxOrderedCollection(v OrderedCollectionType) {
	t.outbox = &outboxIntermediateType{OrderedCollection: v}

}

// IsOutboxAnyURI determines whether the call to GetOutboxAnyURI is safe
func (t *Accept) IsOutboxAnyURI() (ok bool) {
	return t.outbox != nil && t.outbox.anyURI != nil

}

// GetOutboxAnyURI returns the value safely if IsOutboxAnyURI returned true
func (t *Accept) GetOutboxAnyURI() (v *url.URL) {
	return t.outbox.anyURI

}

// SetOutboxAnyURI sets the value of outbox to be of *url.URL type
func (t *Accept) SetOutboxAnyURI(v *url.URL) {
	t.outbox = &outboxIntermediateType{anyURI: v}

}

// HasUnknownOutbox determines whether the call to GetUnknownOutbox is safe
func (t *Accept) HasUnknownOutbox() (ok bool) {
	return t.outbox != nil && t.outbox.unknown_ != nil

}

// GetUnknownOutbox returns the unknown value for outbox
func (t *Accept) GetUnknownOutbox() (v interface{}) {
	return t.outbox.unknown_

}

// SetUnknownOutbox sets the unknown value of outbox
func (t *Accept) SetUnknownOutbox(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &outboxIntermediateType{}
	tmp.unknown_ = i
	t.outbox = tmp

}

// IsFollowingCollection determines whether the call to GetFollowingCollection is safe
func (t *Accept) IsFollowingCollection() (ok bool) {
	return t.following != nil && t.following.Collection != nil

}

// GetFollowingCollection returns the value safely if IsFollowingCollection returned true
func (t *Accept) GetFollowingCollection() (v CollectionType) {
	return t.following.Collection

}

// SetFollowingCollection sets the value of following to be of CollectionType type
func (t *Accept) SetFollowingCollection(v CollectionType) {
	t.following = &followingIntermediateType{Collection: v}

}

// IsFollowingOrderedCollection determines whether the call to GetFollowingOrderedCollection is safe
func (t *Accept) IsFollowingOrderedCollection() (ok bool) {
	return t.following != nil && t.following.OrderedCollection != nil

}

// GetFollowingOrderedCollection returns the value safely if IsFollowingOrderedCollection returned true
func (t *Accept) GetFollowingOrderedCollection() (v OrderedCollectionType) {
	return t.following.OrderedCollection

}

// SetFollowingOrderedCollection sets the value of following to be of OrderedCollectionType type
func (t *Accept) SetFollowingOrderedCollection(v OrderedCollectionType) {
	t.following = &followingIntermediateType{OrderedCollection: v}

}

// IsFollowingAnyURI determines whether the call to GetFollowingAnyURI is safe
func (t *Accept) IsFollowingAnyURI() (ok bool) {
	return t.following != nil && t.following.anyURI != nil

}

// GetFollowingAnyURI returns the value safely if IsFollowingAnyURI returned true
func (t *Accept) GetFollowingAnyURI() (v *url.URL) {
	return t.following.anyURI

}

// SetFollowingAnyURI sets the value of following to be of *url.URL type
func (t *Accept) SetFollowingAnyURI(v *url.URL) {
	t.following = &followingIntermediateType{anyURI: v}

}

// HasUnknownFollowing determines whether the call to GetUnknownFollowing is safe
func (t *Accept) HasUnknownFollowing() (ok bool) {
	return t.following != nil && t.following.unknown_ != nil

}

// GetUnknownFollowing returns the unknown value for following
func (t *Accept) GetUnknownFollowing() (v interface{}) {
	return t.following.unknown_

}

// SetUnknownFollowing sets the unknown value of following
func (t *Accept) SetUnknownFollowing(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &followingIntermediateType{}
	tmp.unknown_ = i
	t.following = tmp

}

// IsFollowersCollection determines whether the call to GetFollowersCollection is safe
func (t *Accept) IsFollowersCollection() (ok bool) {
	return t.followers != nil && t.followers.Collection != nil

}

// GetFollowersCollection returns the value safely if IsFollowersCollection returned true
func (t *Accept) GetFollowersCollection() (v CollectionType) {
	return t.followers.Collection

}

// SetFollowersCollection sets the value of followers to be of CollectionType type
func (t *Accept) SetFollowersCollection(v CollectionType) {
	t.followers = &followersIntermediateType{Collection: v}

}

// IsFollowersOrderedCollection determines whether the call to GetFollowersOrderedCollection is safe
func (t *Accept) IsFollowersOrderedCollection() (ok bool) {
	return t.followers != nil && t.followers.OrderedCollection != nil

}

// GetFollowersOrderedCollection returns the value safely if IsFollowersOrderedCollection returned true
func (t *Accept) GetFollowersOrderedCollection() (v OrderedCollectionType) {
	return t.followers.OrderedCollection

}

// SetFollowersOrderedCollection sets the value of followers to be of OrderedCollectionType type
func (t *Accept) SetFollowersOrderedCollection(v OrderedCollectionType) {
	t.followers = &followersIntermediateType{OrderedCollection: v}

}

// IsFollowersAnyURI determines whether the call to GetFollowersAnyURI is safe
func (t *Accept) IsFollowersAnyURI() (ok bool) {
	return t.followers != nil && t.followers.anyURI != nil

}

// GetFollowersAnyURI returns the value safely if IsFollowersAnyURI returned true
func (t *Accept) GetFollowersAnyURI() (v *url.URL) {
	return t.followers.anyURI

}

// SetFollowersAnyURI sets the value of followers to be of *url.URL type
func (t *Accept) SetFollowersAnyURI(v *url.URL) {
	t.followers = &followersIntermediateType{anyURI: v}

}

// HasUnknownFollowers determines whether the call to GetUnknownFollowers is safe
func (t *Accept) HasUnknownFollowers() (ok bool) {
	return t.followers != nil && t.followers.unknown_ != nil

}

// GetUnknownFollowers returns the unknown value for followers
func (t *Accept) GetUnknownFollowers() (v interface{}) {
	return t.followers.unknown_

}

// SetUnknownFollowers sets the unknown value of followers
func (t *Accept) SetUnknownFollowers(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &followersIntermediateType{}
	tmp.unknown_ = i
	t.followers = tmp

}

// IsLikedCollection determines whether the call to GetLikedCollection is safe
func (t *Accept) IsLikedCollection() (ok bool) {
	return t.liked != nil && t.liked.Collection != nil

}

// GetLikedCollection returns the value safely if IsLikedCollection returned true
func (t *Accept) GetLikedCollection() (v CollectionType) {
	return t.liked.Collection

}

// SetLikedCollection sets the value of liked to be of CollectionType type
func (t *Accept) SetLikedCollection(v CollectionType) {
	t.liked = &likedIntermediateType{Collection: v}

}

// IsLikedOrderedCollection determines whether the call to GetLikedOrderedCollection is safe
func (t *Accept) IsLikedOrderedCollection() (ok bool) {
	return t.liked != nil && t.liked.OrderedCollection != nil

}

// GetLikedOrderedCollection returns the value safely if IsLikedOrderedCollection returned true
func (t *Accept) GetLikedOrderedCollection() (v OrderedCollectionType) {
	return t.liked.OrderedCollection

}

// SetLikedOrderedCollection sets the value of liked to be of OrderedCollectionType type
func (t *Accept) SetLikedOrderedCollection(v OrderedCollectionType) {
	t.liked = &likedIntermediateType{OrderedCollection: v}

}

// IsLikedAnyURI determines whether the call to GetLikedAnyURI is safe
func (t *Accept) IsLikedAnyURI() (ok bool) {
	return t.liked != nil && t.liked.anyURI != nil

}

// GetLikedAnyURI returns the value safely if IsLikedAnyURI returned true
func (t *Accept) GetLikedAnyURI() (v *url.URL) {
	return t.liked.anyURI

}

// SetLikedAnyURI sets the value of liked to be of *url.URL type
func (t *Accept) SetLikedAnyURI(v *url.URL) {
	t.liked = &likedIntermediateType{anyURI: v}

}

// HasUnknownLiked determines whether the call to GetUnknownLiked is safe
func (t *Accept) HasUnknownLiked() (ok bool) {
	return t.liked != nil && t.liked.unknown_ != nil

}

// GetUnknownLiked returns the unknown value for liked
func (t *Accept) GetUnknownLiked() (v interface{}) {
	return t.liked.unknown_

}

// SetUnknownLiked sets the unknown value of liked
func (t *Accept) SetUnknownLiked(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &likedIntermediateType{}
	tmp.unknown_ = i
	t.liked = tmp

}

// IsLikesCollection determines whether the call to GetLikesCollection is safe
func (t *Accept) IsLikesCollection() (ok bool) {
	return t.likes != nil && t.likes.Collection != nil

}

// GetLikesCollection returns the value safely if IsLikesCollection returned true
func (t *Accept) GetLikesCollection() (v CollectionType) {
	return t.likes.Collection

}

// SetLikesCollection sets the value of likes to be of CollectionType type
func (t *Accept) SetLikesCollection(v CollectionType) {
	t.likes = &likesIntermediateType{Collection: v}

}

// IsLikesOrderedCollection determines whether the call to GetLikesOrderedCollection is safe
func (t *Accept) IsLikesOrderedCollection() (ok bool) {
	return t.likes != nil && t.likes.OrderedCollection != nil

}

// GetLikesOrderedCollection returns the value safely if IsLikesOrderedCollection returned true
func (t *Accept) GetLikesOrderedCollection() (v OrderedCollectionType) {
	return t.likes.OrderedCollection

}

// SetLikesOrderedCollection sets the value of likes to be of OrderedCollectionType type
func (t *Accept) SetLikesOrderedCollection(v OrderedCollectionType) {
	t.likes = &likesIntermediateType{OrderedCollection: v}

}

// IsLikesAnyURI determines whether the call to GetLikesAnyURI is safe
func (t *Accept) IsLikesAnyURI() (ok bool) {
	return t.likes != nil && t.likes.anyURI != nil

}

// GetLikesAnyURI returns the value safely if IsLikesAnyURI returned true
func (t *Accept) GetLikesAnyURI() (v *url.URL) {
	return t.likes.anyURI

}

// SetLikesAnyURI sets the value of likes to be of *url.URL type
func (t *Accept) SetLikesAnyURI(v *url.URL) {
	t.likes = &likesIntermediateType{anyURI: v}

}

// HasUnknownLikes determines whether the call to GetUnknownLikes is safe
func (t *Accept) HasUnknownLikes() (ok bool) {
	return t.likes != nil && t.likes.unknown_ != nil

}

// GetUnknownLikes returns the unknown value for likes
func (t *Accept) GetUnknownLikes() (v interface{}) {
	return t.likes.unknown_

}

// SetUnknownLikes sets the unknown value of likes
func (t *Accept) SetUnknownLikes(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &likesIntermediateType{}
	tmp.unknown_ = i
	t.likes = tmp

}

// StreamsLen determines the number of elements able to be used for the GetStreams and RemoveStreams functions
func (t *Accept) StreamsLen() (l int) {
	return len(t.streams)

}

// GetStreams returns the value for the specified index
func (t *Accept) GetStreams(index int) (v *url.URL) {
	return t.streams[index]

}

// AppendStreams adds a value to the back of streams
func (t *Accept) AppendStreams(v *url.URL) {
	t.streams = append(t.streams, v)

}

// PrependStreams adds a value to the front of streams
func (t *Accept) PrependStreams(v *url.URL) {
	t.streams = append([]*url.URL{v}, t.streams...)

}

// RemoveStreams deletes the value from the specified index
func (t *Accept) RemoveStreams(index int) {
	copy(t.streams[index:], t.streams[index+1:])
	t.streams[len(t.streams)-1] = nil
	t.streams = t.streams[:len(t.streams)-1]

}

// HasUnknownStreams determines whether the call to GetUnknownStreams is safe
func (t *Accept) HasUnknownStreams() (ok bool) {
	return t.unknown_ != nil && t.unknown_["streams"] != nil

}

// GetUnknownStreams returns the unknown value for streams
func (t *Accept) GetUnknownStreams() (v interface{}) {
	return t.unknown_["streams"]

}

// SetUnknownStreams sets the unknown value of streams
func (t *Accept) SetUnknownStreams(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["streams"] = i

}

// IsPreferredUsername determines whether the call to GetPreferredUsername is safe
func (t *Accept) IsPreferredUsername() (ok bool) {
	return t.preferredUsername != nil && t.preferredUsername.stringName != nil

}

// GetPreferredUsername returns the value safely if IsPreferredUsername returned true
func (t *Accept) GetPreferredUsername() (v string) {
	return *t.preferredUsername.stringName

}

// SetPreferredUsername sets the value of preferredUsername to be of string type
func (t *Accept) SetPreferredUsername(v string) {
	t.preferredUsername = &preferredUsernameIntermediateType{stringName: &v}

}

// IsPreferredUsernameIRI determines whether the call to GetPreferredUsernameIRI is safe
func (t *Accept) IsPreferredUsernameIRI() (ok bool) {
	return t.preferredUsername != nil && t.preferredUsername.IRI != nil

}

// GetPreferredUsernameIRI returns the value safely if IsPreferredUsernameIRI returned true
func (t *Accept) GetPreferredUsernameIRI() (v *url.URL) {
	return t.preferredUsername.IRI

}

// SetPreferredUsernameIRI sets the value of preferredUsername to be of *url.URL type
func (t *Accept) SetPreferredUsernameIRI(v *url.URL) {
	t.preferredUsername = &preferredUsernameIntermediateType{IRI: v}

}

// HasUnknownPreferredUsername determines whether the call to GetUnknownPreferredUsername is safe
func (t *Accept) HasUnknownPreferredUsername() (ok bool) {
	return t.preferredUsername != nil && t.preferredUsername.unknown_ != nil

}

// GetUnknownPreferredUsername returns the unknown value for preferredUsername
func (t *Accept) GetUnknownPreferredUsername() (v interface{}) {
	return t.preferredUsername.unknown_

}

// SetUnknownPreferredUsername sets the unknown value of preferredUsername
func (t *Accept) SetUnknownPreferredUsername(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &preferredUsernameIntermediateType{}
	tmp.unknown_ = i
	t.preferredUsername = tmp

}

// PreferredUsernameMapLanguages returns all languages for this property's language mapping, or nil if there are none.
func (t *Accept) PreferredUsernameMapLanguages() (l []string) {
	if t.preferredUsernameMap == nil || len(t.preferredUsernameMap) == 0 {
		return nil
	}
	for k := range t.preferredUsernameMap {
		l = append(l, k)
	}
	return

}

// GetPreferredUsernameMap retrieves the value of the property for the specified language, or an empty string if it does not exist
func (t *Accept) GetPreferredUsernameMap(l string) (v string) {
	if t.preferredUsernameMap == nil {
		return ""
	}
	ok := false
	v, ok = t.preferredUsernameMap[l]
	if !ok {
		return ""
	}
	return v

}

// SetPreferredUsernameMap sets the value of the property for the specified language
func (t *Accept) SetPreferredUsernameMap(l string, v string) {
	if t.preferredUsernameMap == nil {
		t.preferredUsernameMap = make(map[string]string)
	}
	t.preferredUsernameMap[l] = v

}

// IsEndpoints determines whether the call to GetEndpoints is safe
func (t *Accept) IsEndpoints() (ok bool) {
	return t.endpoints != nil && t.endpoints.Object != nil

}

// GetEndpoints returns the value safely if IsEndpoints returned true
func (t *Accept) GetEndpoints() (v ObjectType) {
	return t.endpoints.Object

}

// SetEndpoints sets the value of endpoints to be of ObjectType type
func (t *Accept) SetEndpoints(v ObjectType) {
	t.endpoints = &endpointsIntermediateType{Object: v}

}

// IsEndpointsIRI determines whether the call to GetEndpointsIRI is safe
func (t *Accept) IsEndpointsIRI() (ok bool) {
	return t.endpoints != nil && t.endpoints.IRI != nil

}

// GetEndpointsIRI returns the value safely if IsEndpointsIRI returned true
func (t *Accept) GetEndpointsIRI() (v *url.URL) {
	return t.endpoints.IRI

}

// SetEndpointsIRI sets the value of endpoints to be of *url.URL type
func (t *Accept) SetEndpointsIRI(v *url.URL) {
	t.endpoints = &endpointsIntermediateType{IRI: v}

}

// HasUnknownEndpoints determines whether the call to GetUnknownEndpoints is safe
func (t *Accept) HasUnknownEndpoints() (ok bool) {
	return t.endpoints != nil && t.endpoints.unknown_ != nil

}

// GetUnknownEndpoints returns the unknown value for endpoints
func (t *Accept) GetUnknownEndpoints() (v interface{}) {
	return t.endpoints.unknown_

}

// SetUnknownEndpoints sets the unknown value of endpoints
func (t *Accept) SetUnknownEndpoints(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &endpointsIntermediateType{}
	tmp.unknown_ = i
	t.endpoints = tmp

}

// HasProxyUrl determines whether the call to GetProxyUrl is safe
func (t *Accept) HasProxyUrl() (ok bool) {
	return t.proxyUrl != nil

}

// GetProxyUrl returns the value for proxyUrl
func (t *Accept) GetProxyUrl() (v *url.URL) {
	return t.proxyUrl

}

// SetProxyUrl sets the value of proxyUrl
func (t *Accept) SetProxyUrl(v *url.URL) {
	t.proxyUrl = v

}

// HasUnknownProxyUrl determines whether the call to GetUnknownProxyUrl is safe
func (t *Accept) HasUnknownProxyUrl() (ok bool) {
	return t.unknown_ != nil && t.unknown_["proxyUrl"] != nil

}

// GetUnknownProxyUrl returns the unknown value for proxyUrl
func (t *Accept) GetUnknownProxyUrl() (v interface{}) {
	return t.unknown_["proxyUrl"]

}

// SetUnknownProxyUrl sets the unknown value of proxyUrl
func (t *Accept) SetUnknownProxyUrl(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["proxyUrl"] = i

}

// HasOauthAuthorizationEndpoint determines whether the call to GetOauthAuthorizationEndpoint is safe
func (t *Accept) HasOauthAuthorizationEndpoint() (ok bool) {
	return t.oauthAuthorizationEndpoint != nil

}

// GetOauthAuthorizationEndpoint returns the value for oauthAuthorizationEndpoint
func (t *Accept) GetOauthAuthorizationEndpoint() (v *url.URL) {
	return t.oauthAuthorizationEndpoint

}

// SetOauthAuthorizationEndpoint sets the value of oauthAuthorizationEndpoint
func (t *Accept) SetOauthAuthorizationEndpoint(v *url.URL) {
	t.oauthAuthorizationEndpoint = v

}

// HasUnknownOauthAuthorizationEndpoint determines whether the call to GetUnknownOauthAuthorizationEndpoint is safe
func (t *Accept) HasUnknownOauthAuthorizationEndpoint() (ok bool) {
	return t.unknown_ != nil && t.unknown_["oauthAuthorizationEndpoint"] != nil

}

// GetUnknownOauthAuthorizationEndpoint returns the unknown value for oauthAuthorizationEndpoint
func (t *Accept) GetUnknownOauthAuthorizationEndpoint() (v interface{}) {
	return t.unknown_["oauthAuthorizationEndpoint"]

}

// SetUnknownOauthAuthorizationEndpoint sets the unknown value of oauthAuthorizationEndpoint
func (t *Accept) SetUnknownOauthAuthorizationEndpoint(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["oauthAuthorizationEndpoint"] = i

}

// HasOauthTokenEndpoint determines whether the call to GetOauthTokenEndpoint is safe
func (t *Accept) HasOauthTokenEndpoint() (ok bool) {
	return t.oauthTokenEndpoint != nil

}

// GetOauthTokenEndpoint returns the value for oauthTokenEndpoint
func (t *Accept) GetOauthTokenEndpoint() (v *url.URL) {
	return t.oauthTokenEndpoint

}

// SetOauthTokenEndpoint sets the value of oauthTokenEndpoint
func (t *Accept) SetOauthTokenEndpoint(v *url.URL) {
	t.oauthTokenEndpoint = v

}

// HasUnknownOauthTokenEndpoint determines whether the call to GetUnknownOauthTokenEndpoint is safe
func (t *Accept) HasUnknownOauthTokenEndpoint() (ok bool) {
	return t.unknown_ != nil && t.unknown_["oauthTokenEndpoint"] != nil

}

// GetUnknownOauthTokenEndpoint returns the unknown value for oauthTokenEndpoint
func (t *Accept) GetUnknownOauthTokenEndpoint() (v interface{}) {
	return t.unknown_["oauthTokenEndpoint"]

}

// SetUnknownOauthTokenEndpoint sets the unknown value of oauthTokenEndpoint
func (t *Accept) SetUnknownOauthTokenEndpoint(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["oauthTokenEndpoint"] = i

}

// HasProvideClientKey determines whether the call to GetProvideClientKey is safe
func (t *Accept) HasProvideClientKey() (ok bool) {
	return t.provideClientKey != nil

}

// GetProvideClientKey returns the value for provideClientKey
func (t *Accept) GetProvideClientKey() (v *url.URL) {
	return t.provideClientKey

}

// SetProvideClientKey sets the value of provideClientKey
func (t *Accept) SetProvideClientKey(v *url.URL) {
	t.provideClientKey = v

}

// HasUnknownProvideClientKey determines whether the call to GetUnknownProvideClientKey is safe
func (t *Accept) HasUnknownProvideClientKey() (ok bool) {
	return t.unknown_ != nil && t.unknown_["provideClientKey"] != nil

}

// GetUnknownProvideClientKey returns the unknown value for provideClientKey
func (t *Accept) GetUnknownProvideClientKey() (v interface{}) {
	return t.unknown_["provideClientKey"]

}

// SetUnknownProvideClientKey sets the unknown value of provideClientKey
func (t *Accept) SetUnknownProvideClientKey(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["provideClientKey"] = i

}

// HasSignClientKey determines whether the call to GetSignClientKey is safe
func (t *Accept) HasSignClientKey() (ok bool) {
	return t.signClientKey != nil

}

// GetSignClientKey returns the value for signClientKey
func (t *Accept) GetSignClientKey() (v *url.URL) {
	return t.signClientKey

}

// SetSignClientKey sets the value of signClientKey
func (t *Accept) SetSignClientKey(v *url.URL) {
	t.signClientKey = v

}

// HasUnknownSignClientKey determines whether the call to GetUnknownSignClientKey is safe
func (t *Accept) HasUnknownSignClientKey() (ok bool) {
	return t.unknown_ != nil && t.unknown_["signClientKey"] != nil

}

// GetUnknownSignClientKey returns the unknown value for signClientKey
func (t *Accept) GetUnknownSignClientKey() (v interface{}) {
	return t.unknown_["signClientKey"]

}

// SetUnknownSignClientKey sets the unknown value of signClientKey
func (t *Accept) SetUnknownSignClientKey(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["signClientKey"] = i

}

// HasSharedInbox determines whether the call to GetSharedInbox is safe
func (t *Accept) HasSharedInbox() (ok bool) {
	return t.sharedInbox != nil

}

// GetSharedInbox returns the value for sharedInbox
func (t *Accept) GetSharedInbox() (v *url.URL) {
	return t.sharedInbox

}

// SetSharedInbox sets the value of sharedInbox
func (t *Accept) SetSharedInbox(v *url.URL) {
	t.sharedInbox = v

}

// HasUnknownSharedInbox determines whether the call to GetUnknownSharedInbox is safe
func (t *Accept) HasUnknownSharedInbox() (ok bool) {
	return t.unknown_ != nil && t.unknown_["sharedInbox"] != nil

}

// GetUnknownSharedInbox returns the unknown value for sharedInbox
func (t *Accept) GetUnknownSharedInbox() (v interface{}) {
	return t.unknown_["sharedInbox"]

}

// SetUnknownSharedInbox sets the unknown value of sharedInbox
func (t *Accept) SetUnknownSharedInbox(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["sharedInbox"] = i

}

// IsSharesCollection determines whether the call to GetSharesCollection is safe
func (t *Accept) IsSharesCollection() (ok bool) {
	return t.shares != nil && t.shares.Collection != nil

}

// GetSharesCollection returns the value safely if IsSharesCollection returned true
func (t *Accept) GetSharesCollection() (v CollectionType) {
	return t.shares.Collection

}

// SetSharesCollection sets the value of shares to be of CollectionType type
func (t *Accept) SetSharesCollection(v CollectionType) {
	t.shares = &sharesIntermediateType{Collection: v}

}

// IsSharesOrderedCollection determines whether the call to GetSharesOrderedCollection is safe
func (t *Accept) IsSharesOrderedCollection() (ok bool) {
	return t.shares != nil && t.shares.OrderedCollection != nil

}

// GetSharesOrderedCollection returns the value safely if IsSharesOrderedCollection returned true
func (t *Accept) GetSharesOrderedCollection() (v OrderedCollectionType) {
	return t.shares.OrderedCollection

}

// SetSharesOrderedCollection sets the value of shares to be of OrderedCollectionType type
func (t *Accept) SetSharesOrderedCollection(v OrderedCollectionType) {
	t.shares = &sharesIntermediateType{OrderedCollection: v}

}

// IsSharesAnyURI determines whether the call to GetSharesAnyURI is safe
func (t *Accept) IsSharesAnyURI() (ok bool) {
	return t.shares != nil && t.shares.anyURI != nil

}

// GetSharesAnyURI returns the value safely if IsSharesAnyURI returned true
func (t *Accept) GetSharesAnyURI() (v *url.URL) {
	return t.shares.anyURI

}

// SetSharesAnyURI sets the value of shares to be of *url.URL type
func (t *Accept) SetSharesAnyURI(v *url.URL) {
	t.shares = &sharesIntermediateType{anyURI: v}

}

// HasUnknownShares determines whether the call to GetUnknownShares is safe
func (t *Accept) HasUnknownShares() (ok bool) {
	return t.shares != nil && t.shares.unknown_ != nil

}

// GetUnknownShares returns the unknown value for shares
func (t *Accept) GetUnknownShares() (v interface{}) {
	return t.shares.unknown_

}

// SetUnknownShares sets the unknown value of shares
func (t *Accept) SetUnknownShares(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &sharesIntermediateType{}
	tmp.unknown_ = i
	t.shares = tmp

}

// AddUnknown adds an unknown property to this object with the specified key
func (t *Accept) AddUnknown(k string, i interface{}) (this *Accept) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_[k] = i
	return t

}

// HasUnknown returns true if there is an unknown property with the specified key
func (t *Accept) HasUnknown(k string) (b bool) {
	if t.unknown_ == nil {
		return false
	}
	_, ok := t.unknown_[k]
	return ok

}

// RemoveUnknown removes a raw extension from this object with the specified key
func (t *Accept) RemoveUnknown(k string) (this *Accept) {
	delete(t.unknown_, k)
	return t

}

// GetUnknown fetches an unknown property from this object with the specified key. Note that this will panic if HasUnknown would return false.
func (t *Accept) GetUnknown(k string) (i interface{}) {
	return t.unknown_[k]

}

// Serialize turns this object into a map[string]interface{}. Note that the "type" property will automatically be populated with "Accept" if not manually set by the caller
func (t *Accept) Serialize() (m map[string]interface{}, err error) {
	m = make(map[string]interface{})
	for k, v := range t.unknown_ {
		m[k] = unknownValueSerialize(v)
	}
	var typeAlreadySet bool
	for _, k := range t.typeName {
		if ks, ok := k.(string); ok {
			if ks == "Accept" {
				typeAlreadySet = true
				break
			}
		}
	}
	if !typeAlreadySet {
		t.typeName = append(t.typeName, "Accept")
	}
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceActorIntermediateType(t.actor); err == nil && v != nil {
		if len(v) == 1 {
			m["actor"] = v[0]
		} else {
			m["actor"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceObjectIntermediateType(t.object); err == nil && v != nil {
		if len(v) == 1 {
			m["object"] = v[0]
		} else {
			m["object"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceTargetIntermediateType(t.target); err == nil && v != nil {
		if len(v) == 1 {
			m["target"] = v[0]
		} else {
			m["target"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceResultIntermediateType(t.result); err == nil && v != nil {
		if len(v) == 1 {
			m["result"] = v[0]
		} else {
			m["result"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceOriginIntermediateType(t.origin); err == nil && v != nil {
		if len(v) == 1 {
			m["origin"] = v[0]
		} else {
			m["origin"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceInstrumentIntermediateType(t.instrument); err == nil && v != nil {
		if len(v) == 1 {
			m["instrument"] = v[0]
		} else {
			m["instrument"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.altitude != nil {
		if v, err := serializeAltitudeIntermediateType(t.altitude); err == nil {
			m["altitude"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceAttachmentIntermediateType(t.attachment); err == nil && v != nil {
		if len(v) == 1 {
			m["attachment"] = v[0]
		} else {
			m["attachment"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceAttributedToIntermediateType(t.attributedTo); err == nil && v != nil {
		if len(v) == 1 {
			m["attributedTo"] = v[0]
		} else {
			m["attributedTo"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceAudienceIntermediateType(t.audience); err == nil && v != nil {
		if len(v) == 1 {
			m["audience"] = v[0]
		} else {
			m["audience"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceContentIntermediateType(t.content); err == nil && v != nil {
		if len(v) == 1 {
			m["content"] = v[0]
		} else {
			m["content"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition

	// Begin generation by generateNaturalLanguageMap
	if t.contentMap != nil && len(t.contentMap) >= 0 {
		m["contentMap"] = t.contentMap
	}
	// End generation by generateNaturalLanguageMap
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceContextIntermediateType(t.context); err == nil && v != nil {
		if len(v) == 1 {
			m["context"] = v[0]
		} else {
			m["context"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceNameIntermediateType(t.name); err == nil && v != nil {
		if len(v) == 1 {
			m["name"] = v[0]
		} else {
			m["name"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition

	// Begin generation by generateNaturalLanguageMap
	if t.nameMap != nil && len(t.nameMap) >= 0 {
		m["nameMap"] = t.nameMap
	}
	// End generation by generateNaturalLanguageMap
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.endTime != nil {
		if v, err := serializeEndTimeIntermediateType(t.endTime); err == nil {
			m["endTime"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceGeneratorIntermediateType(t.generator); err == nil && v != nil {
		if len(v) == 1 {
			m["generator"] = v[0]
		} else {
			m["generator"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceIconIntermediateType(t.icon); err == nil && v != nil {
		if len(v) == 1 {
			m["icon"] = v[0]
		} else {
			m["icon"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by RangeReference.Serialize for Value
	if t.id != nil {
		idSerializeFunc := func() (interface{}, error) {
			v := t.id
			tmp := anyURISerialize(v)
			return tmp, nil
		}
		idResult, err := idSerializeFunc()
		if err == nil {
			m["id"] = idResult
		} else {
			return m, err
		}
	}
	// End generation by RangeReference.Serialize for Value
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceImageIntermediateType(t.image); err == nil && v != nil {
		if len(v) == 1 {
			m["image"] = v[0]
		} else {
			m["image"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceInReplyToIntermediateType(t.inReplyTo); err == nil && v != nil {
		if len(v) == 1 {
			m["inReplyTo"] = v[0]
		} else {
			m["inReplyTo"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceLocationIntermediateType(t.location); err == nil && v != nil {
		if len(v) == 1 {
			m["location"] = v[0]
		} else {
			m["location"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSlicePreviewIntermediateType(t.preview); err == nil && v != nil {
		if len(v) == 1 {
			m["preview"] = v[0]
		} else {
			m["preview"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.published != nil {
		if v, err := serializePublishedIntermediateType(t.published); err == nil {
			m["published"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.replies != nil {
		if v, err := serializeRepliesIntermediateType(t.replies); err == nil {
			m["replies"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.startTime != nil {
		if v, err := serializeStartTimeIntermediateType(t.startTime); err == nil {
			m["startTime"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceSummaryIntermediateType(t.summary); err == nil && v != nil {
		if len(v) == 1 {
			m["summary"] = v[0]
		} else {
			m["summary"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition

	// Begin generation by generateNaturalLanguageMap
	if t.summaryMap != nil && len(t.summaryMap) >= 0 {
		m["summaryMap"] = t.summaryMap
	}
	// End generation by generateNaturalLanguageMap
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceTagIntermediateType(t.tag); err == nil && v != nil {
		if len(v) == 1 {
			m["tag"] = v[0]
		} else {
			m["tag"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalAnyDefinition
	if t.typeName != nil {
		if len(t.typeName) == 1 {
			m["type"] = t.typeName[0]
		} else {
			m["type"] = t.typeName
		}
	}
	// End generation by generateNonFunctionalAnyDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.updated != nil {
		if v, err := serializeUpdatedIntermediateType(t.updated); err == nil {
			m["updated"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceUrlIntermediateType(t.url); err == nil && v != nil {
		if len(v) == 1 {
			m["url"] = v[0]
		} else {
			m["url"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceToIntermediateType(t.to); err == nil && v != nil {
		if len(v) == 1 {
			m["to"] = v[0]
		} else {
			m["to"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceBtoIntermediateType(t.bto); err == nil && v != nil {
		if len(v) == 1 {
			m["bto"] = v[0]
		} else {
			m["bto"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceCcIntermediateType(t.cc); err == nil && v != nil {
		if len(v) == 1 {
			m["cc"] = v[0]
		} else {
			m["cc"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceBccIntermediateType(t.bcc); err == nil && v != nil {
		if len(v) == 1 {
			m["bcc"] = v[0]
		} else {
			m["bcc"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.mediaType != nil {
		if v, err := serializeMediaTypeIntermediateType(t.mediaType); err == nil {
			m["mediaType"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.duration != nil {
		if v, err := serializeDurationIntermediateType(t.duration); err == nil {
			m["duration"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.source != nil {
		if v, err := serializeSourceIntermediateType(t.source); err == nil {
			m["source"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.inbox != nil {
		if v, err := serializeInboxIntermediateType(t.inbox); err == nil {
			m["inbox"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.outbox != nil {
		if v, err := serializeOutboxIntermediateType(t.outbox); err == nil {
			m["outbox"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.following != nil {
		if v, err := serializeFollowingIntermediateType(t.following); err == nil {
			m["following"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.followers != nil {
		if v, err := serializeFollowersIntermediateType(t.followers); err == nil {
			m["followers"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.liked != nil {
		if v, err := serializeLikedIntermediateType(t.liked); err == nil {
			m["liked"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.likes != nil {
		if v, err := serializeLikesIntermediateType(t.likes); err == nil {
			m["likes"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by RangeReference.Serialize for Value
	var streamsTemp []interface{}
	for _, v := range t.streams {
		tmp := anyURISerialize(v)
		streamsTemp = append(streamsTemp, tmp)
	}
	if streamsTemp != nil {
		if len(streamsTemp) == 1 {
			m["streams"] = streamsTemp[0]
		} else {
			m["streams"] = streamsTemp
		}
	}
	// End generation by RangeReference.Serialize for Value
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.preferredUsername != nil {
		if v, err := serializePreferredUsernameIntermediateType(t.preferredUsername); err == nil {
			m["preferredUsername"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition

	// Begin generation by generateNaturalLanguageMap
	if t.preferredUsernameMap != nil && len(t.preferredUsernameMap) >= 0 {
		m["preferredUsernameMap"] = t.preferredUsernameMap
	}
	// End generation by generateNaturalLanguageMap
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.endpoints != nil {
		if v, err := serializeEndpointsIntermediateType(t.endpoints); err == nil {
			m["endpoints"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by RangeReference.Serialize for Value
	if t.proxyUrl != nil {
		proxyUrlSerializeFunc := func() (interface{}, error) {
			v := t.proxyUrl
			tmp := anyURISerialize(v)
			return tmp, nil
		}
		proxyUrlResult, err := proxyUrlSerializeFunc()
		if err == nil {
			m["proxyUrl"] = proxyUrlResult
		} else {
			return m, err
		}
	}
	// End generation by RangeReference.Serialize for Value
	// Begin generation by RangeReference.Serialize for Value
	if t.oauthAuthorizationEndpoint != nil {
		oauthAuthorizationEndpointSerializeFunc := func() (interface{}, error) {
			v := t.oauthAuthorizationEndpoint
			tmp := anyURISerialize(v)
			return tmp, nil
		}
		oauthAuthorizationEndpointResult, err := oauthAuthorizationEndpointSerializeFunc()
		if err == nil {
			m["oauthAuthorizationEndpoint"] = oauthAuthorizationEndpointResult
		} else {
			return m, err
		}
	}
	// End generation by RangeReference.Serialize for Value
	// Begin generation by RangeReference.Serialize for Value
	if t.oauthTokenEndpoint != nil {
		oauthTokenEndpointSerializeFunc := func() (interface{}, error) {
			v := t.oauthTokenEndpoint
			tmp := anyURISerialize(v)
			return tmp, nil
		}
		oauthTokenEndpointResult, err := oauthTokenEndpointSerializeFunc()
		if err == nil {
			m["oauthTokenEndpoint"] = oauthTokenEndpointResult
		} else {
			return m, err
		}
	}
	// End generation by RangeReference.Serialize for Value
	// Begin generation by RangeReference.Serialize for Value
	if t.provideClientKey != nil {
		provideClientKeySerializeFunc := func() (interface{}, error) {
			v := t.provideClientKey
			tmp := anyURISerialize(v)
			return tmp, nil
		}
		provideClientKeyResult, err := provideClientKeySerializeFunc()
		if err == nil {
			m["provideClientKey"] = provideClientKeyResult
		} else {
			return m, err
		}
	}
	// End generation by RangeReference.Serialize for Value
	// Begin generation by RangeReference.Serialize for Value
	if t.signClientKey != nil {
		signClientKeySerializeFunc := func() (interface{}, error) {
			v := t.signClientKey
			tmp := anyURISerialize(v)
			return tmp, nil
		}
		signClientKeyResult, err := signClientKeySerializeFunc()
		if err == nil {
			m["signClientKey"] = signClientKeyResult
		} else {
			return m, err
		}
	}
	// End generation by RangeReference.Serialize for Value
	// Begin generation by RangeReference.Serialize for Value
	if t.sharedInbox != nil {
		sharedInboxSerializeFunc := func() (interface{}, error) {
			v := t.sharedInbox
			tmp := anyURISerialize(v)
			return tmp, nil
		}
		sharedInboxResult, err := sharedInboxSerializeFunc()
		if err == nil {
			m["sharedInbox"] = sharedInboxResult
		} else {
			return m, err
		}
	}
	// End generation by RangeReference.Serialize for Value
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.shares != nil {
		if v, err := serializeSharesIntermediateType(t.shares); err == nil {
			m["shares"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	return

}

// Deserialize populates this object from a map[string]interface{}
func (t *Accept) Deserialize(m map[string]interface{}) (err error) {
	for k, v := range m {
		handled := false
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "actor" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeActorIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.actor = []*actorIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.actor, err = deserializeSliceActorIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &actorIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.actor = []*actorIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "object" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeObjectIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.object = []*objectIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.object, err = deserializeSliceObjectIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &objectIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.object = []*objectIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "target" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeTargetIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.target = []*targetIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.target, err = deserializeSliceTargetIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &targetIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.target = []*targetIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "result" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeResultIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.result = []*resultIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.result, err = deserializeSliceResultIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &resultIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.result = []*resultIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "origin" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeOriginIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.origin = []*originIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.origin, err = deserializeSliceOriginIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &originIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.origin = []*originIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "instrument" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeInstrumentIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.instrument = []*instrumentIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.instrument, err = deserializeSliceInstrumentIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &instrumentIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.instrument = []*instrumentIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "altitude" {
				t.altitude, err = deserializeAltitudeIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "attachment" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeAttachmentIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.attachment = []*attachmentIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.attachment, err = deserializeSliceAttachmentIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &attachmentIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.attachment = []*attachmentIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "attributedTo" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeAttributedToIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.attributedTo = []*attributedToIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.attributedTo, err = deserializeSliceAttributedToIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &attributedToIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.attributedTo = []*attributedToIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "audience" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeAudienceIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.audience = []*audienceIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.audience, err = deserializeSliceAudienceIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &audienceIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.audience = []*audienceIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "content" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeContentIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.content = []*contentIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.content, err = deserializeSliceContentIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &contentIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.content = []*contentIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition

			// Begin generation by generateNaturalLanguageMap
			if k == "contentMap" {
				if vMap, ok := v.(map[string]interface{}); ok {
					val := make(map[string]string)
					for k, iVal := range vMap {
						if sVal, ok := iVal.(string); ok {
							val[k] = sVal
						}
					}
					t.contentMap = val
					handled = true
				}
			}
			// End generation by generateNaturalLanguageMap
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "context" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeContextIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.context = []*contextIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.context, err = deserializeSliceContextIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &contextIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.context = []*contextIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "name" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeNameIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.name = []*nameIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.name, err = deserializeSliceNameIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &nameIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.name = []*nameIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition

			// Begin generation by generateNaturalLanguageMap
			if k == "nameMap" {
				if vMap, ok := v.(map[string]interface{}); ok {
					val := make(map[string]string)
					for k, iVal := range vMap {
						if sVal, ok := iVal.(string); ok {
							val[k] = sVal
						}
					}
					t.nameMap = val
					handled = true
				}
			}
			// End generation by generateNaturalLanguageMap
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "endTime" {
				t.endTime, err = deserializeEndTimeIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "generator" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeGeneratorIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.generator = []*generatorIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.generator, err = deserializeSliceGeneratorIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &generatorIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.generator = []*generatorIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "icon" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeIconIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.icon = []*iconIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.icon, err = deserializeSliceIconIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &iconIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.icon = []*iconIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by RangeReference.Deserialize for Value
			if k == "id" {
				if v, ok := v.(interface{}); ok {
					tmp, err := anyURIDeserialize(v)
					if err != nil {
						return err
					}
					t.id = tmp
					handled = true
				}
			}
			// End generation by RangeReference.Deserialize for Value
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "image" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeImageIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.image = []*imageIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.image, err = deserializeSliceImageIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &imageIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.image = []*imageIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "inReplyTo" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeInReplyToIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.inReplyTo = []*inReplyToIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.inReplyTo, err = deserializeSliceInReplyToIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &inReplyToIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.inReplyTo = []*inReplyToIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "location" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeLocationIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.location = []*locationIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.location, err = deserializeSliceLocationIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &locationIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.location = []*locationIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "preview" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializePreviewIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.preview = []*previewIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.preview, err = deserializeSlicePreviewIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &previewIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.preview = []*previewIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "published" {
				t.published, err = deserializePublishedIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "replies" {
				t.replies, err = deserializeRepliesIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "startTime" {
				t.startTime, err = deserializeStartTimeIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "summary" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeSummaryIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.summary = []*summaryIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.summary, err = deserializeSliceSummaryIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &summaryIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.summary = []*summaryIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition

			// Begin generation by generateNaturalLanguageMap
			if k == "summaryMap" {
				if vMap, ok := v.(map[string]interface{}); ok {
					val := make(map[string]string)
					for k, iVal := range vMap {
						if sVal, ok := iVal.(string); ok {
							val[k] = sVal
						}
					}
					t.summaryMap = val
					handled = true
				}
			}
			// End generation by generateNaturalLanguageMap
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "tag" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeTagIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.tag = []*tagIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.tag, err = deserializeSliceTagIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &tagIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.tag = []*tagIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalAnyDefinition
			if k == "type" {
				if tmpTypeSlice, ok := v.([]interface{}); ok {
					t.typeName = tmpTypeSlice
					handled = true
				} else {
					t.typeName = []interface{}{v}
					handled = true
				}
			}
			// End generation by generateNonFunctionalAnyDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "updated" {
				t.updated, err = deserializeUpdatedIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "url" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeUrlIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.url = []*urlIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.url, err = deserializeSliceUrlIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &urlIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.url = []*urlIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "to" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeToIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.to = []*toIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.to, err = deserializeSliceToIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &toIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.to = []*toIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "bto" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeBtoIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.bto = []*btoIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.bto, err = deserializeSliceBtoIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &btoIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.bto = []*btoIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "cc" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeCcIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.cc = []*ccIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.cc, err = deserializeSliceCcIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &ccIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.cc = []*ccIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "bcc" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeBccIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.bcc = []*bccIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.bcc, err = deserializeSliceBccIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &bccIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.bcc = []*bccIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "mediaType" {
				t.mediaType, err = deserializeMediaTypeIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "duration" {
				t.duration, err = deserializeDurationIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "source" {
				t.source, err = deserializeSourceIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "inbox" {
				t.inbox, err = deserializeInboxIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "outbox" {
				t.outbox, err = deserializeOutboxIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "following" {
				t.following, err = deserializeFollowingIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "followers" {
				t.followers, err = deserializeFollowersIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "liked" {
				t.liked, err = deserializeLikedIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "likes" {
				t.likes, err = deserializeLikesIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by RangeReference.Deserialize for Value
			if k == "streams" {
				if tmpSlice, ok := v.([]interface{}); ok {
					for _, tmpElem := range tmpSlice {
						if v, ok := tmpElem.(interface{}); ok {
							tmp, err := anyURIDeserialize(v)
							if err != nil {
								return err
							}
							t.streams = append(t.streams, tmp)
							handled = true
						}
					}
				} else if v, ok := v.(interface{}); ok {
					tmp, err := anyURIDeserialize(v)
					if err != nil {
						return err
					}
					t.streams = append(t.streams, tmp)
					handled = true
				}
			}
			// End generation by RangeReference.Deserialize for Value
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "preferredUsername" {
				t.preferredUsername, err = deserializePreferredUsernameIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition

			// Begin generation by generateNaturalLanguageMap
			if k == "preferredUsernameMap" {
				if vMap, ok := v.(map[string]interface{}); ok {
					val := make(map[string]string)
					for k, iVal := range vMap {
						if sVal, ok := iVal.(string); ok {
							val[k] = sVal
						}
					}
					t.preferredUsernameMap = val
					handled = true
				}
			}
			// End generation by generateNaturalLanguageMap
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "endpoints" {
				t.endpoints, err = deserializeEndpointsIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by RangeReference.Deserialize for Value
			if k == "proxyUrl" {
				if v, ok := v.(interface{}); ok {
					tmp, err := anyURIDeserialize(v)
					if err != nil {
						return err
					}
					t.proxyUrl = tmp
					handled = true
				}
			}
			// End generation by RangeReference.Deserialize for Value
		}
		if !handled {
			// Begin generation by RangeReference.Deserialize for Value
			if k == "oauthAuthorizationEndpoint" {
				if v, ok := v.(interface{}); ok {
					tmp, err := anyURIDeserialize(v)
					if err != nil {
						return err
					}
					t.oauthAuthorizationEndpoint = tmp
					handled = true
				}
			}
			// End generation by RangeReference.Deserialize for Value
		}
		if !handled {
			// Begin generation by RangeReference.Deserialize for Value
			if k == "oauthTokenEndpoint" {
				if v, ok := v.(interface{}); ok {
					tmp, err := anyURIDeserialize(v)
					if err != nil {
						return err
					}
					t.oauthTokenEndpoint = tmp
					handled = true
				}
			}
			// End generation by RangeReference.Deserialize for Value
		}
		if !handled {
			// Begin generation by RangeReference.Deserialize for Value
			if k == "provideClientKey" {
				if v, ok := v.(interface{}); ok {
					tmp, err := anyURIDeserialize(v)
					if err != nil {
						return err
					}
					t.provideClientKey = tmp
					handled = true
				}
			}
			// End generation by RangeReference.Deserialize for Value
		}
		if !handled {
			// Begin generation by RangeReference.Deserialize for Value
			if k == "signClientKey" {
				if v, ok := v.(interface{}); ok {
					tmp, err := anyURIDeserialize(v)
					if err != nil {
						return err
					}
					t.signClientKey = tmp
					handled = true
				}
			}
			// End generation by RangeReference.Deserialize for Value
		}
		if !handled {
			// Begin generation by RangeReference.Deserialize for Value
			if k == "sharedInbox" {
				if v, ok := v.(interface{}); ok {
					tmp, err := anyURIDeserialize(v)
					if err != nil {
						return err
					}
					t.sharedInbox = tmp
					handled = true
				}
			}
			// End generation by RangeReference.Deserialize for Value
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "shares" {
				t.shares, err = deserializeSharesIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled && k != "@context" {
			if t.unknown_ == nil {
				t.unknown_ = make(map[string]interface{})
			}
			t.unknown_[k] = unknownValueDeserialize(v)
		}
	}
	return

}

// IsPublic returns true if the 'to', 'bto', 'cc', or 'bcc' properties address the special Public ActivityPub collection
func (t *Accept) IsPublic() (b bool) {
	for i := 0; i < t.ToLen(); i++ {
		if t.IsToIRI(i) && t.GetToIRI(i).String() == "https://www.w3.org/ns/activitystreams#Public" {
			return true
		}
	}
	for i := 0; i < t.BtoLen(); i++ {
		if t.IsBtoIRI(i) && t.GetBtoIRI(i).String() == "https://www.w3.org/ns/activitystreams#Public" {
			return true
		}
	}
	for i := 0; i < t.CcLen(); i++ {
		if t.IsCcIRI(i) && t.GetCcIRI(i).String() == "https://www.w3.org/ns/activitystreams#Public" {
			return true
		}
	}
	for i := 0; i < t.BccLen(); i++ {
		if t.IsBccIRI(i) && t.GetBccIRI(i).String() == "https://www.w3.org/ns/activitystreams#Public" {
			return true
		}
	}
	return false

}

//
package vocab

import (
	"fmt"
	"net/url"
	"time"
)

// TentativeAcceptType is an interface for accepting types that extend from 'TentativeAccept'.
type TentativeAcceptType interface {
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
	IsAltitude() (ok bool)
	GetAltitude() (v float64)
	SetAltitude(v float64)
	IsAltitudeIRI() (ok bool)
	GetAltitudeIRI() (v *url.URL)
	SetAltitudeIRI(v *url.URL)
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
	NameMapLanguages() (l []string)
	GetNameMap(l string) (v string)
	SetNameMap(l string, v string)
	IsEndTime() (ok bool)
	GetEndTime() (v time.Time)
	SetEndTime(v time.Time)
	IsEndTimeIRI() (ok bool)
	GetEndTimeIRI() (v *url.URL)
	SetEndTimeIRI(v *url.URL)
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
	IsPublished() (ok bool)
	GetPublished() (v time.Time)
	SetPublished(v time.Time)
	IsPublishedIRI() (ok bool)
	GetPublishedIRI() (v *url.URL)
	SetPublishedIRI(v *url.URL)
	IsReplies() (ok bool)
	GetReplies() (v CollectionType)
	SetReplies(v CollectionType)
	IsRepliesIRI() (ok bool)
	GetRepliesIRI() (v *url.URL)
	SetRepliesIRI(v *url.URL)
	IsStartTime() (ok bool)
	GetStartTime() (v time.Time)
	SetStartTime(v time.Time)
	IsStartTimeIRI() (ok bool)
	GetStartTimeIRI() (v *url.URL)
	SetStartTimeIRI(v *url.URL)
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
	IsMediaType() (ok bool)
	GetMediaType() (v string)
	SetMediaType(v string)
	IsMediaTypeIRI() (ok bool)
	GetMediaTypeIRI() (v *url.URL)
	SetMediaTypeIRI(v *url.URL)
	IsDuration() (ok bool)
	GetDuration() (v time.Duration)
	SetDuration(v time.Duration)
	IsDurationIRI() (ok bool)
	GetDurationIRI() (v *url.URL)
	SetDurationIRI(v *url.URL)
	IsSource() (ok bool)
	GetSource() (v ObjectType)
	SetSource(v ObjectType)
	IsSourceIRI() (ok bool)
	GetSourceIRI() (v *url.URL)
	SetSourceIRI(v *url.URL)
	IsInboxOrderedCollection() (ok bool)
	GetInboxOrderedCollection() (v OrderedCollectionType)
	SetInboxOrderedCollection(v OrderedCollectionType)
	IsInboxAnyURI() (ok bool)
	GetInboxAnyURI() (v *url.URL)
	SetInboxAnyURI(v *url.URL)
	IsOutboxOrderedCollection() (ok bool)
	GetOutboxOrderedCollection() (v OrderedCollectionType)
	SetOutboxOrderedCollection(v OrderedCollectionType)
	IsOutboxAnyURI() (ok bool)
	GetOutboxAnyURI() (v *url.URL)
	SetOutboxAnyURI(v *url.URL)
	IsFollowingCollection() (ok bool)
	GetFollowingCollection() (v CollectionType)
	SetFollowingCollection(v CollectionType)
	IsFollowingOrderedCollection() (ok bool)
	GetFollowingOrderedCollection() (v OrderedCollectionType)
	SetFollowingOrderedCollection(v OrderedCollectionType)
	IsFollowingAnyURI() (ok bool)
	GetFollowingAnyURI() (v *url.URL)
	SetFollowingAnyURI(v *url.URL)
	IsFollowersCollection() (ok bool)
	GetFollowersCollection() (v CollectionType)
	SetFollowersCollection(v CollectionType)
	IsFollowersOrderedCollection() (ok bool)
	GetFollowersOrderedCollection() (v OrderedCollectionType)
	SetFollowersOrderedCollection(v OrderedCollectionType)
	IsFollowersAnyURI() (ok bool)
	GetFollowersAnyURI() (v *url.URL)
	SetFollowersAnyURI(v *url.URL)
	IsLikedCollection() (ok bool)
	GetLikedCollection() (v CollectionType)
	SetLikedCollection(v CollectionType)
	IsLikedOrderedCollection() (ok bool)
	GetLikedOrderedCollection() (v OrderedCollectionType)
	SetLikedOrderedCollection(v OrderedCollectionType)
	IsLikedAnyURI() (ok bool)
	GetLikedAnyURI() (v *url.URL)
	SetLikedAnyURI(v *url.URL)
	IsLikesCollection() (ok bool)
	GetLikesCollection() (v CollectionType)
	SetLikesCollection(v CollectionType)
	IsLikesOrderedCollection() (ok bool)
	GetLikesOrderedCollection() (v OrderedCollectionType)
	SetLikesOrderedCollection(v OrderedCollectionType)
	IsLikesAnyURI() (ok bool)
	GetLikesAnyURI() (v *url.URL)
	SetLikesAnyURI(v *url.URL)
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
	PreferredUsernameMapLanguages() (l []string)
	GetPreferredUsernameMap(l string) (v string)
	SetPreferredUsernameMap(l string, v string)
	IsEndpoints() (ok bool)
	GetEndpoints() (v ObjectType)
	SetEndpoints(v ObjectType)
	IsEndpointsIRI() (ok bool)
	GetEndpointsIRI() (v *url.URL)
	SetEndpointsIRI(v *url.URL)
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
}

// A specialization of Accept indicating that the acceptance is tentative.
type TentativeAccept struct {
	// An unknown value.
	unknown_ map[string]interface{}
	// The 'actor' value could have multiple types and values
	actor []*actorTentativeAcceptIntermediateType
	// The 'object' value could have multiple types and values
	object []*objectTentativeAcceptIntermediateType
	// The 'target' value could have multiple types and values
	target []*targetTentativeAcceptIntermediateType
	// The 'result' value could have multiple types and values
	result []*resultTentativeAcceptIntermediateType
	// The 'origin' value could have multiple types and values
	origin []*originTentativeAcceptIntermediateType
	// The 'instrument' value could have multiple types and values
	instrument []*instrumentTentativeAcceptIntermediateType
	// The functional 'altitude' value could have multiple types, but only a single value
	altitude *altitudeTentativeAcceptIntermediateType
	// The 'attachment' value could have multiple types and values
	attachment []*attachmentTentativeAcceptIntermediateType
	// The 'attributedTo' value could have multiple types and values
	attributedTo []*attributedToTentativeAcceptIntermediateType
	// The 'audience' value could have multiple types and values
	audience []*audienceTentativeAcceptIntermediateType
	// The 'content' value could have multiple types and values
	content []*contentTentativeAcceptIntermediateType
	// The 'contentMap' value holds language-specific values for property 'content'
	contentMap map[string]string
	// The 'context' value could have multiple types and values
	context []*contextTentativeAcceptIntermediateType
	// The 'name' value could have multiple types and values
	name []*nameTentativeAcceptIntermediateType
	// The 'nameMap' value holds language-specific values for property 'name'
	nameMap map[string]string
	// The functional 'endTime' value could have multiple types, but only a single value
	endTime *endTimeTentativeAcceptIntermediateType
	// The 'generator' value could have multiple types and values
	generator []*generatorTentativeAcceptIntermediateType
	// The 'icon' value could have multiple types and values
	icon []*iconTentativeAcceptIntermediateType
	// The functional 'id' value holds a single type and a single value
	id *url.URL
	// The 'image' value could have multiple types and values
	image []*imageTentativeAcceptIntermediateType
	// The 'inReplyTo' value could have multiple types and values
	inReplyTo []*inReplyToTentativeAcceptIntermediateType
	// The 'location' value could have multiple types and values
	location []*locationTentativeAcceptIntermediateType
	// The 'preview' value could have multiple types and values
	preview []*previewTentativeAcceptIntermediateType
	// The functional 'published' value could have multiple types, but only a single value
	published *publishedTentativeAcceptIntermediateType
	// The functional 'replies' value could have multiple types, but only a single value
	replies *repliesTentativeAcceptIntermediateType
	// The functional 'startTime' value could have multiple types, but only a single value
	startTime *startTimeTentativeAcceptIntermediateType
	// The 'summary' value could have multiple types and values
	summary []*summaryTentativeAcceptIntermediateType
	// The 'summaryMap' value holds language-specific values for property 'summary'
	summaryMap map[string]string
	// The 'tag' value could have multiple types and values
	tag []*tagTentativeAcceptIntermediateType
	// The 'type' value can hold any type and any number of values
	typeName []interface{}
	// The functional 'updated' value could have multiple types, but only a single value
	updated *updatedTentativeAcceptIntermediateType
	// The 'url' value could have multiple types and values
	url []*urlTentativeAcceptIntermediateType
	// The 'to' value could have multiple types and values
	to []*toTentativeAcceptIntermediateType
	// The 'bto' value could have multiple types and values
	bto []*btoTentativeAcceptIntermediateType
	// The 'cc' value could have multiple types and values
	cc []*ccTentativeAcceptIntermediateType
	// The 'bcc' value could have multiple types and values
	bcc []*bccTentativeAcceptIntermediateType
	// The functional 'mediaType' value could have multiple types, but only a single value
	mediaType *mediaTypeTentativeAcceptIntermediateType
	// The functional 'duration' value could have multiple types, but only a single value
	duration *durationTentativeAcceptIntermediateType
	// The functional 'source' value could have multiple types, but only a single value
	source *sourceTentativeAcceptIntermediateType
	// The functional 'inbox' value could have multiple types, but only a single value
	inbox *inboxTentativeAcceptIntermediateType
	// The functional 'outbox' value could have multiple types, but only a single value
	outbox *outboxTentativeAcceptIntermediateType
	// The functional 'following' value could have multiple types, but only a single value
	following *followingTentativeAcceptIntermediateType
	// The functional 'followers' value could have multiple types, but only a single value
	followers *followersTentativeAcceptIntermediateType
	// The functional 'liked' value could have multiple types, but only a single value
	liked *likedTentativeAcceptIntermediateType
	// The functional 'likes' value could have multiple types, but only a single value
	likes *likesTentativeAcceptIntermediateType
	// The 'streams' value holds a single type and any number of values
	streams []*url.URL
	// The functional 'preferredUsername' value could have multiple types, but only a single value
	preferredUsername *preferredUsernameTentativeAcceptIntermediateType
	// The 'preferredUsernameMap' value holds language-specific values for property 'preferredUsername'
	preferredUsernameMap map[string]string
	// The functional 'endpoints' value could have multiple types, but only a single value
	endpoints *endpointsTentativeAcceptIntermediateType
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
}

// ActorLen determines the number of elements able to be used for the IsActorObject, GetActorObject, and RemoveActorObject functions
func (t *TentativeAccept) ActorLen() (l int) {
	return len(t.actor)

}

// IsActorObject determines whether the call to GetActorObject is safe for the specified index
func (t *TentativeAccept) IsActorObject(index int) (ok bool) {
	return t.actor[index].Object != nil

}

// GetActorObject returns the value safely if IsActorObject returned true for the specified index
func (t *TentativeAccept) GetActorObject(index int) (v ObjectType) {
	return t.actor[index].Object

}

// AppendActorObject adds to the back of actor a ObjectType type
func (t *TentativeAccept) AppendActorObject(v ObjectType) {
	t.actor = append(t.actor, &actorTentativeAcceptIntermediateType{Object: v})

}

// PrependActorObject adds to the front of actor a ObjectType type
func (t *TentativeAccept) PrependActorObject(v ObjectType) {
	t.actor = append([]*actorTentativeAcceptIntermediateType{&actorTentativeAcceptIntermediateType{Object: v}}, t.actor...)

}

// RemoveActorObject deletes the value from the specified index
func (t *TentativeAccept) RemoveActorObject(index int) {
	copy(t.actor[index:], t.actor[index+1:])
	t.actor[len(t.actor)-1] = nil
	t.actor = t.actor[:len(t.actor)-1]

}

// IsActorLink determines whether the call to GetActorLink is safe for the specified index
func (t *TentativeAccept) IsActorLink(index int) (ok bool) {
	return t.actor[index].Link != nil

}

// GetActorLink returns the value safely if IsActorLink returned true for the specified index
func (t *TentativeAccept) GetActorLink(index int) (v LinkType) {
	return t.actor[index].Link

}

// AppendActorLink adds to the back of actor a LinkType type
func (t *TentativeAccept) AppendActorLink(v LinkType) {
	t.actor = append(t.actor, &actorTentativeAcceptIntermediateType{Link: v})

}

// PrependActorLink adds to the front of actor a LinkType type
func (t *TentativeAccept) PrependActorLink(v LinkType) {
	t.actor = append([]*actorTentativeAcceptIntermediateType{&actorTentativeAcceptIntermediateType{Link: v}}, t.actor...)

}

// RemoveActorLink deletes the value from the specified index
func (t *TentativeAccept) RemoveActorLink(index int) {
	copy(t.actor[index:], t.actor[index+1:])
	t.actor[len(t.actor)-1] = nil
	t.actor = t.actor[:len(t.actor)-1]

}

// IsActorIRI determines whether the call to GetActorIRI is safe for the specified index
func (t *TentativeAccept) IsActorIRI(index int) (ok bool) {
	return t.actor[index].IRI != nil

}

// GetActorIRI returns the value safely if IsActorIRI returned true for the specified index
func (t *TentativeAccept) GetActorIRI(index int) (v *url.URL) {
	return t.actor[index].IRI

}

// AppendActorIRI adds to the back of actor a *url.URL type
func (t *TentativeAccept) AppendActorIRI(v *url.URL) {
	t.actor = append(t.actor, &actorTentativeAcceptIntermediateType{IRI: v})

}

// PrependActorIRI adds to the front of actor a *url.URL type
func (t *TentativeAccept) PrependActorIRI(v *url.URL) {
	t.actor = append([]*actorTentativeAcceptIntermediateType{&actorTentativeAcceptIntermediateType{IRI: v}}, t.actor...)

}

// RemoveActorIRI deletes the value from the specified index
func (t *TentativeAccept) RemoveActorIRI(index int) {
	copy(t.actor[index:], t.actor[index+1:])
	t.actor[len(t.actor)-1] = nil
	t.actor = t.actor[:len(t.actor)-1]

}

// HasUnknownActor determines whether the call to GetUnknownActor is safe
func (t *TentativeAccept) HasUnknownActor() (ok bool) {
	return t.actor != nil && t.actor[0].unknown_ != nil

}

// GetUnknownActor returns the unknown value for actor
func (t *TentativeAccept) GetUnknownActor() (v interface{}) {
	return t.actor[0].unknown_

}

// SetUnknownActor sets the unknown value of actor
func (t *TentativeAccept) SetUnknownActor(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &actorTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.actor = append(t.actor, tmp)

}

// ObjectLen determines the number of elements able to be used for the IsObject, GetObject, and RemoveObject functions
func (t *TentativeAccept) ObjectLen() (l int) {
	return len(t.object)

}

// IsObject determines whether the call to GetObject is safe for the specified index
func (t *TentativeAccept) IsObject(index int) (ok bool) {
	return t.object[index].Object != nil

}

// GetObject returns the value safely if IsObject returned true for the specified index
func (t *TentativeAccept) GetObject(index int) (v ObjectType) {
	return t.object[index].Object

}

// AppendObject adds to the back of object a ObjectType type
func (t *TentativeAccept) AppendObject(v ObjectType) {
	t.object = append(t.object, &objectTentativeAcceptIntermediateType{Object: v})

}

// PrependObject adds to the front of object a ObjectType type
func (t *TentativeAccept) PrependObject(v ObjectType) {
	t.object = append([]*objectTentativeAcceptIntermediateType{&objectTentativeAcceptIntermediateType{Object: v}}, t.object...)

}

// RemoveObject deletes the value from the specified index
func (t *TentativeAccept) RemoveObject(index int) {
	copy(t.object[index:], t.object[index+1:])
	t.object[len(t.object)-1] = nil
	t.object = t.object[:len(t.object)-1]

}

// IsObjectIRI determines whether the call to GetObjectIRI is safe for the specified index
func (t *TentativeAccept) IsObjectIRI(index int) (ok bool) {
	return t.object[index].IRI != nil

}

// GetObjectIRI returns the value safely if IsObjectIRI returned true for the specified index
func (t *TentativeAccept) GetObjectIRI(index int) (v *url.URL) {
	return t.object[index].IRI

}

// AppendObjectIRI adds to the back of object a *url.URL type
func (t *TentativeAccept) AppendObjectIRI(v *url.URL) {
	t.object = append(t.object, &objectTentativeAcceptIntermediateType{IRI: v})

}

// PrependObjectIRI adds to the front of object a *url.URL type
func (t *TentativeAccept) PrependObjectIRI(v *url.URL) {
	t.object = append([]*objectTentativeAcceptIntermediateType{&objectTentativeAcceptIntermediateType{IRI: v}}, t.object...)

}

// RemoveObjectIRI deletes the value from the specified index
func (t *TentativeAccept) RemoveObjectIRI(index int) {
	copy(t.object[index:], t.object[index+1:])
	t.object[len(t.object)-1] = nil
	t.object = t.object[:len(t.object)-1]

}

// HasUnknownObject determines whether the call to GetUnknownObject is safe
func (t *TentativeAccept) HasUnknownObject() (ok bool) {
	return t.object != nil && t.object[0].unknown_ != nil

}

// GetUnknownObject returns the unknown value for object
func (t *TentativeAccept) GetUnknownObject() (v interface{}) {
	return t.object[0].unknown_

}

// SetUnknownObject sets the unknown value of object
func (t *TentativeAccept) SetUnknownObject(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &objectTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.object = append(t.object, tmp)

}

// TargetLen determines the number of elements able to be used for the IsTargetObject, GetTargetObject, and RemoveTargetObject functions
func (t *TentativeAccept) TargetLen() (l int) {
	return len(t.target)

}

// IsTargetObject determines whether the call to GetTargetObject is safe for the specified index
func (t *TentativeAccept) IsTargetObject(index int) (ok bool) {
	return t.target[index].Object != nil

}

// GetTargetObject returns the value safely if IsTargetObject returned true for the specified index
func (t *TentativeAccept) GetTargetObject(index int) (v ObjectType) {
	return t.target[index].Object

}

// AppendTargetObject adds to the back of target a ObjectType type
func (t *TentativeAccept) AppendTargetObject(v ObjectType) {
	t.target = append(t.target, &targetTentativeAcceptIntermediateType{Object: v})

}

// PrependTargetObject adds to the front of target a ObjectType type
func (t *TentativeAccept) PrependTargetObject(v ObjectType) {
	t.target = append([]*targetTentativeAcceptIntermediateType{&targetTentativeAcceptIntermediateType{Object: v}}, t.target...)

}

// RemoveTargetObject deletes the value from the specified index
func (t *TentativeAccept) RemoveTargetObject(index int) {
	copy(t.target[index:], t.target[index+1:])
	t.target[len(t.target)-1] = nil
	t.target = t.target[:len(t.target)-1]

}

// IsTargetLink determines whether the call to GetTargetLink is safe for the specified index
func (t *TentativeAccept) IsTargetLink(index int) (ok bool) {
	return t.target[index].Link != nil

}

// GetTargetLink returns the value safely if IsTargetLink returned true for the specified index
func (t *TentativeAccept) GetTargetLink(index int) (v LinkType) {
	return t.target[index].Link

}

// AppendTargetLink adds to the back of target a LinkType type
func (t *TentativeAccept) AppendTargetLink(v LinkType) {
	t.target = append(t.target, &targetTentativeAcceptIntermediateType{Link: v})

}

// PrependTargetLink adds to the front of target a LinkType type
func (t *TentativeAccept) PrependTargetLink(v LinkType) {
	t.target = append([]*targetTentativeAcceptIntermediateType{&targetTentativeAcceptIntermediateType{Link: v}}, t.target...)

}

// RemoveTargetLink deletes the value from the specified index
func (t *TentativeAccept) RemoveTargetLink(index int) {
	copy(t.target[index:], t.target[index+1:])
	t.target[len(t.target)-1] = nil
	t.target = t.target[:len(t.target)-1]

}

// IsTargetIRI determines whether the call to GetTargetIRI is safe for the specified index
func (t *TentativeAccept) IsTargetIRI(index int) (ok bool) {
	return t.target[index].IRI != nil

}

// GetTargetIRI returns the value safely if IsTargetIRI returned true for the specified index
func (t *TentativeAccept) GetTargetIRI(index int) (v *url.URL) {
	return t.target[index].IRI

}

// AppendTargetIRI adds to the back of target a *url.URL type
func (t *TentativeAccept) AppendTargetIRI(v *url.URL) {
	t.target = append(t.target, &targetTentativeAcceptIntermediateType{IRI: v})

}

// PrependTargetIRI adds to the front of target a *url.URL type
func (t *TentativeAccept) PrependTargetIRI(v *url.URL) {
	t.target = append([]*targetTentativeAcceptIntermediateType{&targetTentativeAcceptIntermediateType{IRI: v}}, t.target...)

}

// RemoveTargetIRI deletes the value from the specified index
func (t *TentativeAccept) RemoveTargetIRI(index int) {
	copy(t.target[index:], t.target[index+1:])
	t.target[len(t.target)-1] = nil
	t.target = t.target[:len(t.target)-1]

}

// HasUnknownTarget determines whether the call to GetUnknownTarget is safe
func (t *TentativeAccept) HasUnknownTarget() (ok bool) {
	return t.target != nil && t.target[0].unknown_ != nil

}

// GetUnknownTarget returns the unknown value for target
func (t *TentativeAccept) GetUnknownTarget() (v interface{}) {
	return t.target[0].unknown_

}

// SetUnknownTarget sets the unknown value of target
func (t *TentativeAccept) SetUnknownTarget(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &targetTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.target = append(t.target, tmp)

}

// ResultLen determines the number of elements able to be used for the IsResultObject, GetResultObject, and RemoveResultObject functions
func (t *TentativeAccept) ResultLen() (l int) {
	return len(t.result)

}

// IsResultObject determines whether the call to GetResultObject is safe for the specified index
func (t *TentativeAccept) IsResultObject(index int) (ok bool) {
	return t.result[index].Object != nil

}

// GetResultObject returns the value safely if IsResultObject returned true for the specified index
func (t *TentativeAccept) GetResultObject(index int) (v ObjectType) {
	return t.result[index].Object

}

// AppendResultObject adds to the back of result a ObjectType type
func (t *TentativeAccept) AppendResultObject(v ObjectType) {
	t.result = append(t.result, &resultTentativeAcceptIntermediateType{Object: v})

}

// PrependResultObject adds to the front of result a ObjectType type
func (t *TentativeAccept) PrependResultObject(v ObjectType) {
	t.result = append([]*resultTentativeAcceptIntermediateType{&resultTentativeAcceptIntermediateType{Object: v}}, t.result...)

}

// RemoveResultObject deletes the value from the specified index
func (t *TentativeAccept) RemoveResultObject(index int) {
	copy(t.result[index:], t.result[index+1:])
	t.result[len(t.result)-1] = nil
	t.result = t.result[:len(t.result)-1]

}

// IsResultLink determines whether the call to GetResultLink is safe for the specified index
func (t *TentativeAccept) IsResultLink(index int) (ok bool) {
	return t.result[index].Link != nil

}

// GetResultLink returns the value safely if IsResultLink returned true for the specified index
func (t *TentativeAccept) GetResultLink(index int) (v LinkType) {
	return t.result[index].Link

}

// AppendResultLink adds to the back of result a LinkType type
func (t *TentativeAccept) AppendResultLink(v LinkType) {
	t.result = append(t.result, &resultTentativeAcceptIntermediateType{Link: v})

}

// PrependResultLink adds to the front of result a LinkType type
func (t *TentativeAccept) PrependResultLink(v LinkType) {
	t.result = append([]*resultTentativeAcceptIntermediateType{&resultTentativeAcceptIntermediateType{Link: v}}, t.result...)

}

// RemoveResultLink deletes the value from the specified index
func (t *TentativeAccept) RemoveResultLink(index int) {
	copy(t.result[index:], t.result[index+1:])
	t.result[len(t.result)-1] = nil
	t.result = t.result[:len(t.result)-1]

}

// IsResultIRI determines whether the call to GetResultIRI is safe for the specified index
func (t *TentativeAccept) IsResultIRI(index int) (ok bool) {
	return t.result[index].IRI != nil

}

// GetResultIRI returns the value safely if IsResultIRI returned true for the specified index
func (t *TentativeAccept) GetResultIRI(index int) (v *url.URL) {
	return t.result[index].IRI

}

// AppendResultIRI adds to the back of result a *url.URL type
func (t *TentativeAccept) AppendResultIRI(v *url.URL) {
	t.result = append(t.result, &resultTentativeAcceptIntermediateType{IRI: v})

}

// PrependResultIRI adds to the front of result a *url.URL type
func (t *TentativeAccept) PrependResultIRI(v *url.URL) {
	t.result = append([]*resultTentativeAcceptIntermediateType{&resultTentativeAcceptIntermediateType{IRI: v}}, t.result...)

}

// RemoveResultIRI deletes the value from the specified index
func (t *TentativeAccept) RemoveResultIRI(index int) {
	copy(t.result[index:], t.result[index+1:])
	t.result[len(t.result)-1] = nil
	t.result = t.result[:len(t.result)-1]

}

// HasUnknownResult determines whether the call to GetUnknownResult is safe
func (t *TentativeAccept) HasUnknownResult() (ok bool) {
	return t.result != nil && t.result[0].unknown_ != nil

}

// GetUnknownResult returns the unknown value for result
func (t *TentativeAccept) GetUnknownResult() (v interface{}) {
	return t.result[0].unknown_

}

// SetUnknownResult sets the unknown value of result
func (t *TentativeAccept) SetUnknownResult(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &resultTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.result = append(t.result, tmp)

}

// OriginLen determines the number of elements able to be used for the IsOriginObject, GetOriginObject, and RemoveOriginObject functions
func (t *TentativeAccept) OriginLen() (l int) {
	return len(t.origin)

}

// IsOriginObject determines whether the call to GetOriginObject is safe for the specified index
func (t *TentativeAccept) IsOriginObject(index int) (ok bool) {
	return t.origin[index].Object != nil

}

// GetOriginObject returns the value safely if IsOriginObject returned true for the specified index
func (t *TentativeAccept) GetOriginObject(index int) (v ObjectType) {
	return t.origin[index].Object

}

// AppendOriginObject adds to the back of origin a ObjectType type
func (t *TentativeAccept) AppendOriginObject(v ObjectType) {
	t.origin = append(t.origin, &originTentativeAcceptIntermediateType{Object: v})

}

// PrependOriginObject adds to the front of origin a ObjectType type
func (t *TentativeAccept) PrependOriginObject(v ObjectType) {
	t.origin = append([]*originTentativeAcceptIntermediateType{&originTentativeAcceptIntermediateType{Object: v}}, t.origin...)

}

// RemoveOriginObject deletes the value from the specified index
func (t *TentativeAccept) RemoveOriginObject(index int) {
	copy(t.origin[index:], t.origin[index+1:])
	t.origin[len(t.origin)-1] = nil
	t.origin = t.origin[:len(t.origin)-1]

}

// IsOriginLink determines whether the call to GetOriginLink is safe for the specified index
func (t *TentativeAccept) IsOriginLink(index int) (ok bool) {
	return t.origin[index].Link != nil

}

// GetOriginLink returns the value safely if IsOriginLink returned true for the specified index
func (t *TentativeAccept) GetOriginLink(index int) (v LinkType) {
	return t.origin[index].Link

}

// AppendOriginLink adds to the back of origin a LinkType type
func (t *TentativeAccept) AppendOriginLink(v LinkType) {
	t.origin = append(t.origin, &originTentativeAcceptIntermediateType{Link: v})

}

// PrependOriginLink adds to the front of origin a LinkType type
func (t *TentativeAccept) PrependOriginLink(v LinkType) {
	t.origin = append([]*originTentativeAcceptIntermediateType{&originTentativeAcceptIntermediateType{Link: v}}, t.origin...)

}

// RemoveOriginLink deletes the value from the specified index
func (t *TentativeAccept) RemoveOriginLink(index int) {
	copy(t.origin[index:], t.origin[index+1:])
	t.origin[len(t.origin)-1] = nil
	t.origin = t.origin[:len(t.origin)-1]

}

// IsOriginIRI determines whether the call to GetOriginIRI is safe for the specified index
func (t *TentativeAccept) IsOriginIRI(index int) (ok bool) {
	return t.origin[index].IRI != nil

}

// GetOriginIRI returns the value safely if IsOriginIRI returned true for the specified index
func (t *TentativeAccept) GetOriginIRI(index int) (v *url.URL) {
	return t.origin[index].IRI

}

// AppendOriginIRI adds to the back of origin a *url.URL type
func (t *TentativeAccept) AppendOriginIRI(v *url.URL) {
	t.origin = append(t.origin, &originTentativeAcceptIntermediateType{IRI: v})

}

// PrependOriginIRI adds to the front of origin a *url.URL type
func (t *TentativeAccept) PrependOriginIRI(v *url.URL) {
	t.origin = append([]*originTentativeAcceptIntermediateType{&originTentativeAcceptIntermediateType{IRI: v}}, t.origin...)

}

// RemoveOriginIRI deletes the value from the specified index
func (t *TentativeAccept) RemoveOriginIRI(index int) {
	copy(t.origin[index:], t.origin[index+1:])
	t.origin[len(t.origin)-1] = nil
	t.origin = t.origin[:len(t.origin)-1]

}

// HasUnknownOrigin determines whether the call to GetUnknownOrigin is safe
func (t *TentativeAccept) HasUnknownOrigin() (ok bool) {
	return t.origin != nil && t.origin[0].unknown_ != nil

}

// GetUnknownOrigin returns the unknown value for origin
func (t *TentativeAccept) GetUnknownOrigin() (v interface{}) {
	return t.origin[0].unknown_

}

// SetUnknownOrigin sets the unknown value of origin
func (t *TentativeAccept) SetUnknownOrigin(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &originTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.origin = append(t.origin, tmp)

}

// InstrumentLen determines the number of elements able to be used for the IsInstrumentObject, GetInstrumentObject, and RemoveInstrumentObject functions
func (t *TentativeAccept) InstrumentLen() (l int) {
	return len(t.instrument)

}

// IsInstrumentObject determines whether the call to GetInstrumentObject is safe for the specified index
func (t *TentativeAccept) IsInstrumentObject(index int) (ok bool) {
	return t.instrument[index].Object != nil

}

// GetInstrumentObject returns the value safely if IsInstrumentObject returned true for the specified index
func (t *TentativeAccept) GetInstrumentObject(index int) (v ObjectType) {
	return t.instrument[index].Object

}

// AppendInstrumentObject adds to the back of instrument a ObjectType type
func (t *TentativeAccept) AppendInstrumentObject(v ObjectType) {
	t.instrument = append(t.instrument, &instrumentTentativeAcceptIntermediateType{Object: v})

}

// PrependInstrumentObject adds to the front of instrument a ObjectType type
func (t *TentativeAccept) PrependInstrumentObject(v ObjectType) {
	t.instrument = append([]*instrumentTentativeAcceptIntermediateType{&instrumentTentativeAcceptIntermediateType{Object: v}}, t.instrument...)

}

// RemoveInstrumentObject deletes the value from the specified index
func (t *TentativeAccept) RemoveInstrumentObject(index int) {
	copy(t.instrument[index:], t.instrument[index+1:])
	t.instrument[len(t.instrument)-1] = nil
	t.instrument = t.instrument[:len(t.instrument)-1]

}

// IsInstrumentLink determines whether the call to GetInstrumentLink is safe for the specified index
func (t *TentativeAccept) IsInstrumentLink(index int) (ok bool) {
	return t.instrument[index].Link != nil

}

// GetInstrumentLink returns the value safely if IsInstrumentLink returned true for the specified index
func (t *TentativeAccept) GetInstrumentLink(index int) (v LinkType) {
	return t.instrument[index].Link

}

// AppendInstrumentLink adds to the back of instrument a LinkType type
func (t *TentativeAccept) AppendInstrumentLink(v LinkType) {
	t.instrument = append(t.instrument, &instrumentTentativeAcceptIntermediateType{Link: v})

}

// PrependInstrumentLink adds to the front of instrument a LinkType type
func (t *TentativeAccept) PrependInstrumentLink(v LinkType) {
	t.instrument = append([]*instrumentTentativeAcceptIntermediateType{&instrumentTentativeAcceptIntermediateType{Link: v}}, t.instrument...)

}

// RemoveInstrumentLink deletes the value from the specified index
func (t *TentativeAccept) RemoveInstrumentLink(index int) {
	copy(t.instrument[index:], t.instrument[index+1:])
	t.instrument[len(t.instrument)-1] = nil
	t.instrument = t.instrument[:len(t.instrument)-1]

}

// IsInstrumentIRI determines whether the call to GetInstrumentIRI is safe for the specified index
func (t *TentativeAccept) IsInstrumentIRI(index int) (ok bool) {
	return t.instrument[index].IRI != nil

}

// GetInstrumentIRI returns the value safely if IsInstrumentIRI returned true for the specified index
func (t *TentativeAccept) GetInstrumentIRI(index int) (v *url.URL) {
	return t.instrument[index].IRI

}

// AppendInstrumentIRI adds to the back of instrument a *url.URL type
func (t *TentativeAccept) AppendInstrumentIRI(v *url.URL) {
	t.instrument = append(t.instrument, &instrumentTentativeAcceptIntermediateType{IRI: v})

}

// PrependInstrumentIRI adds to the front of instrument a *url.URL type
func (t *TentativeAccept) PrependInstrumentIRI(v *url.URL) {
	t.instrument = append([]*instrumentTentativeAcceptIntermediateType{&instrumentTentativeAcceptIntermediateType{IRI: v}}, t.instrument...)

}

// RemoveInstrumentIRI deletes the value from the specified index
func (t *TentativeAccept) RemoveInstrumentIRI(index int) {
	copy(t.instrument[index:], t.instrument[index+1:])
	t.instrument[len(t.instrument)-1] = nil
	t.instrument = t.instrument[:len(t.instrument)-1]

}

// HasUnknownInstrument determines whether the call to GetUnknownInstrument is safe
func (t *TentativeAccept) HasUnknownInstrument() (ok bool) {
	return t.instrument != nil && t.instrument[0].unknown_ != nil

}

// GetUnknownInstrument returns the unknown value for instrument
func (t *TentativeAccept) GetUnknownInstrument() (v interface{}) {
	return t.instrument[0].unknown_

}

// SetUnknownInstrument sets the unknown value of instrument
func (t *TentativeAccept) SetUnknownInstrument(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &instrumentTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.instrument = append(t.instrument, tmp)

}

// IsAltitude determines whether the call to GetAltitude is safe
func (t *TentativeAccept) IsAltitude() (ok bool) {
	return t.altitude != nil && t.altitude.float != nil

}

// GetAltitude returns the value safely if IsAltitude returned true
func (t *TentativeAccept) GetAltitude() (v float64) {
	return *t.altitude.float

}

// SetAltitude sets the value of altitude to be of float64 type
func (t *TentativeAccept) SetAltitude(v float64) {
	t.altitude = &altitudeTentativeAcceptIntermediateType{float: &v}

}

// IsAltitudeIRI determines whether the call to GetAltitudeIRI is safe
func (t *TentativeAccept) IsAltitudeIRI() (ok bool) {
	return t.altitude != nil && t.altitude.IRI != nil

}

// GetAltitudeIRI returns the value safely if IsAltitudeIRI returned true
func (t *TentativeAccept) GetAltitudeIRI() (v *url.URL) {
	return t.altitude.IRI

}

// SetAltitudeIRI sets the value of altitude to be of *url.URL type
func (t *TentativeAccept) SetAltitudeIRI(v *url.URL) {
	t.altitude = &altitudeTentativeAcceptIntermediateType{IRI: v}

}

// HasUnknownAltitude determines whether the call to GetUnknownAltitude is safe
func (t *TentativeAccept) HasUnknownAltitude() (ok bool) {
	return t.altitude != nil && t.altitude.unknown_ != nil

}

// GetUnknownAltitude returns the unknown value for altitude
func (t *TentativeAccept) GetUnknownAltitude() (v interface{}) {
	return t.altitude.unknown_

}

// SetUnknownAltitude sets the unknown value of altitude
func (t *TentativeAccept) SetUnknownAltitude(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &altitudeTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.altitude = tmp

}

// AttachmentLen determines the number of elements able to be used for the IsAttachmentObject, GetAttachmentObject, and RemoveAttachmentObject functions
func (t *TentativeAccept) AttachmentLen() (l int) {
	return len(t.attachment)

}

// IsAttachmentObject determines whether the call to GetAttachmentObject is safe for the specified index
func (t *TentativeAccept) IsAttachmentObject(index int) (ok bool) {
	return t.attachment[index].Object != nil

}

// GetAttachmentObject returns the value safely if IsAttachmentObject returned true for the specified index
func (t *TentativeAccept) GetAttachmentObject(index int) (v ObjectType) {
	return t.attachment[index].Object

}

// AppendAttachmentObject adds to the back of attachment a ObjectType type
func (t *TentativeAccept) AppendAttachmentObject(v ObjectType) {
	t.attachment = append(t.attachment, &attachmentTentativeAcceptIntermediateType{Object: v})

}

// PrependAttachmentObject adds to the front of attachment a ObjectType type
func (t *TentativeAccept) PrependAttachmentObject(v ObjectType) {
	t.attachment = append([]*attachmentTentativeAcceptIntermediateType{&attachmentTentativeAcceptIntermediateType{Object: v}}, t.attachment...)

}

// RemoveAttachmentObject deletes the value from the specified index
func (t *TentativeAccept) RemoveAttachmentObject(index int) {
	copy(t.attachment[index:], t.attachment[index+1:])
	t.attachment[len(t.attachment)-1] = nil
	t.attachment = t.attachment[:len(t.attachment)-1]

}

// IsAttachmentLink determines whether the call to GetAttachmentLink is safe for the specified index
func (t *TentativeAccept) IsAttachmentLink(index int) (ok bool) {
	return t.attachment[index].Link != nil

}

// GetAttachmentLink returns the value safely if IsAttachmentLink returned true for the specified index
func (t *TentativeAccept) GetAttachmentLink(index int) (v LinkType) {
	return t.attachment[index].Link

}

// AppendAttachmentLink adds to the back of attachment a LinkType type
func (t *TentativeAccept) AppendAttachmentLink(v LinkType) {
	t.attachment = append(t.attachment, &attachmentTentativeAcceptIntermediateType{Link: v})

}

// PrependAttachmentLink adds to the front of attachment a LinkType type
func (t *TentativeAccept) PrependAttachmentLink(v LinkType) {
	t.attachment = append([]*attachmentTentativeAcceptIntermediateType{&attachmentTentativeAcceptIntermediateType{Link: v}}, t.attachment...)

}

// RemoveAttachmentLink deletes the value from the specified index
func (t *TentativeAccept) RemoveAttachmentLink(index int) {
	copy(t.attachment[index:], t.attachment[index+1:])
	t.attachment[len(t.attachment)-1] = nil
	t.attachment = t.attachment[:len(t.attachment)-1]

}

// IsAttachmentIRI determines whether the call to GetAttachmentIRI is safe for the specified index
func (t *TentativeAccept) IsAttachmentIRI(index int) (ok bool) {
	return t.attachment[index].IRI != nil

}

// GetAttachmentIRI returns the value safely if IsAttachmentIRI returned true for the specified index
func (t *TentativeAccept) GetAttachmentIRI(index int) (v *url.URL) {
	return t.attachment[index].IRI

}

// AppendAttachmentIRI adds to the back of attachment a *url.URL type
func (t *TentativeAccept) AppendAttachmentIRI(v *url.URL) {
	t.attachment = append(t.attachment, &attachmentTentativeAcceptIntermediateType{IRI: v})

}

// PrependAttachmentIRI adds to the front of attachment a *url.URL type
func (t *TentativeAccept) PrependAttachmentIRI(v *url.URL) {
	t.attachment = append([]*attachmentTentativeAcceptIntermediateType{&attachmentTentativeAcceptIntermediateType{IRI: v}}, t.attachment...)

}

// RemoveAttachmentIRI deletes the value from the specified index
func (t *TentativeAccept) RemoveAttachmentIRI(index int) {
	copy(t.attachment[index:], t.attachment[index+1:])
	t.attachment[len(t.attachment)-1] = nil
	t.attachment = t.attachment[:len(t.attachment)-1]

}

// HasUnknownAttachment determines whether the call to GetUnknownAttachment is safe
func (t *TentativeAccept) HasUnknownAttachment() (ok bool) {
	return t.attachment != nil && t.attachment[0].unknown_ != nil

}

// GetUnknownAttachment returns the unknown value for attachment
func (t *TentativeAccept) GetUnknownAttachment() (v interface{}) {
	return t.attachment[0].unknown_

}

// SetUnknownAttachment sets the unknown value of attachment
func (t *TentativeAccept) SetUnknownAttachment(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &attachmentTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.attachment = append(t.attachment, tmp)

}

// AttributedToLen determines the number of elements able to be used for the IsAttributedToObject, GetAttributedToObject, and RemoveAttributedToObject functions
func (t *TentativeAccept) AttributedToLen() (l int) {
	return len(t.attributedTo)

}

// IsAttributedToObject determines whether the call to GetAttributedToObject is safe for the specified index
func (t *TentativeAccept) IsAttributedToObject(index int) (ok bool) {
	return t.attributedTo[index].Object != nil

}

// GetAttributedToObject returns the value safely if IsAttributedToObject returned true for the specified index
func (t *TentativeAccept) GetAttributedToObject(index int) (v ObjectType) {
	return t.attributedTo[index].Object

}

// AppendAttributedToObject adds to the back of attributedTo a ObjectType type
func (t *TentativeAccept) AppendAttributedToObject(v ObjectType) {
	t.attributedTo = append(t.attributedTo, &attributedToTentativeAcceptIntermediateType{Object: v})

}

// PrependAttributedToObject adds to the front of attributedTo a ObjectType type
func (t *TentativeAccept) PrependAttributedToObject(v ObjectType) {
	t.attributedTo = append([]*attributedToTentativeAcceptIntermediateType{&attributedToTentativeAcceptIntermediateType{Object: v}}, t.attributedTo...)

}

// RemoveAttributedToObject deletes the value from the specified index
func (t *TentativeAccept) RemoveAttributedToObject(index int) {
	copy(t.attributedTo[index:], t.attributedTo[index+1:])
	t.attributedTo[len(t.attributedTo)-1] = nil
	t.attributedTo = t.attributedTo[:len(t.attributedTo)-1]

}

// IsAttributedToLink determines whether the call to GetAttributedToLink is safe for the specified index
func (t *TentativeAccept) IsAttributedToLink(index int) (ok bool) {
	return t.attributedTo[index].Link != nil

}

// GetAttributedToLink returns the value safely if IsAttributedToLink returned true for the specified index
func (t *TentativeAccept) GetAttributedToLink(index int) (v LinkType) {
	return t.attributedTo[index].Link

}

// AppendAttributedToLink adds to the back of attributedTo a LinkType type
func (t *TentativeAccept) AppendAttributedToLink(v LinkType) {
	t.attributedTo = append(t.attributedTo, &attributedToTentativeAcceptIntermediateType{Link: v})

}

// PrependAttributedToLink adds to the front of attributedTo a LinkType type
func (t *TentativeAccept) PrependAttributedToLink(v LinkType) {
	t.attributedTo = append([]*attributedToTentativeAcceptIntermediateType{&attributedToTentativeAcceptIntermediateType{Link: v}}, t.attributedTo...)

}

// RemoveAttributedToLink deletes the value from the specified index
func (t *TentativeAccept) RemoveAttributedToLink(index int) {
	copy(t.attributedTo[index:], t.attributedTo[index+1:])
	t.attributedTo[len(t.attributedTo)-1] = nil
	t.attributedTo = t.attributedTo[:len(t.attributedTo)-1]

}

// IsAttributedToIRI determines whether the call to GetAttributedToIRI is safe for the specified index
func (t *TentativeAccept) IsAttributedToIRI(index int) (ok bool) {
	return t.attributedTo[index].IRI != nil

}

// GetAttributedToIRI returns the value safely if IsAttributedToIRI returned true for the specified index
func (t *TentativeAccept) GetAttributedToIRI(index int) (v *url.URL) {
	return t.attributedTo[index].IRI

}

// AppendAttributedToIRI adds to the back of attributedTo a *url.URL type
func (t *TentativeAccept) AppendAttributedToIRI(v *url.URL) {
	t.attributedTo = append(t.attributedTo, &attributedToTentativeAcceptIntermediateType{IRI: v})

}

// PrependAttributedToIRI adds to the front of attributedTo a *url.URL type
func (t *TentativeAccept) PrependAttributedToIRI(v *url.URL) {
	t.attributedTo = append([]*attributedToTentativeAcceptIntermediateType{&attributedToTentativeAcceptIntermediateType{IRI: v}}, t.attributedTo...)

}

// RemoveAttributedToIRI deletes the value from the specified index
func (t *TentativeAccept) RemoveAttributedToIRI(index int) {
	copy(t.attributedTo[index:], t.attributedTo[index+1:])
	t.attributedTo[len(t.attributedTo)-1] = nil
	t.attributedTo = t.attributedTo[:len(t.attributedTo)-1]

}

// HasUnknownAttributedTo determines whether the call to GetUnknownAttributedTo is safe
func (t *TentativeAccept) HasUnknownAttributedTo() (ok bool) {
	return t.attributedTo != nil && t.attributedTo[0].unknown_ != nil

}

// GetUnknownAttributedTo returns the unknown value for attributedTo
func (t *TentativeAccept) GetUnknownAttributedTo() (v interface{}) {
	return t.attributedTo[0].unknown_

}

// SetUnknownAttributedTo sets the unknown value of attributedTo
func (t *TentativeAccept) SetUnknownAttributedTo(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &attributedToTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.attributedTo = append(t.attributedTo, tmp)

}

// AudienceLen determines the number of elements able to be used for the IsAudienceObject, GetAudienceObject, and RemoveAudienceObject functions
func (t *TentativeAccept) AudienceLen() (l int) {
	return len(t.audience)

}

// IsAudienceObject determines whether the call to GetAudienceObject is safe for the specified index
func (t *TentativeAccept) IsAudienceObject(index int) (ok bool) {
	return t.audience[index].Object != nil

}

// GetAudienceObject returns the value safely if IsAudienceObject returned true for the specified index
func (t *TentativeAccept) GetAudienceObject(index int) (v ObjectType) {
	return t.audience[index].Object

}

// AppendAudienceObject adds to the back of audience a ObjectType type
func (t *TentativeAccept) AppendAudienceObject(v ObjectType) {
	t.audience = append(t.audience, &audienceTentativeAcceptIntermediateType{Object: v})

}

// PrependAudienceObject adds to the front of audience a ObjectType type
func (t *TentativeAccept) PrependAudienceObject(v ObjectType) {
	t.audience = append([]*audienceTentativeAcceptIntermediateType{&audienceTentativeAcceptIntermediateType{Object: v}}, t.audience...)

}

// RemoveAudienceObject deletes the value from the specified index
func (t *TentativeAccept) RemoveAudienceObject(index int) {
	copy(t.audience[index:], t.audience[index+1:])
	t.audience[len(t.audience)-1] = nil
	t.audience = t.audience[:len(t.audience)-1]

}

// IsAudienceLink determines whether the call to GetAudienceLink is safe for the specified index
func (t *TentativeAccept) IsAudienceLink(index int) (ok bool) {
	return t.audience[index].Link != nil

}

// GetAudienceLink returns the value safely if IsAudienceLink returned true for the specified index
func (t *TentativeAccept) GetAudienceLink(index int) (v LinkType) {
	return t.audience[index].Link

}

// AppendAudienceLink adds to the back of audience a LinkType type
func (t *TentativeAccept) AppendAudienceLink(v LinkType) {
	t.audience = append(t.audience, &audienceTentativeAcceptIntermediateType{Link: v})

}

// PrependAudienceLink adds to the front of audience a LinkType type
func (t *TentativeAccept) PrependAudienceLink(v LinkType) {
	t.audience = append([]*audienceTentativeAcceptIntermediateType{&audienceTentativeAcceptIntermediateType{Link: v}}, t.audience...)

}

// RemoveAudienceLink deletes the value from the specified index
func (t *TentativeAccept) RemoveAudienceLink(index int) {
	copy(t.audience[index:], t.audience[index+1:])
	t.audience[len(t.audience)-1] = nil
	t.audience = t.audience[:len(t.audience)-1]

}

// IsAudienceIRI determines whether the call to GetAudienceIRI is safe for the specified index
func (t *TentativeAccept) IsAudienceIRI(index int) (ok bool) {
	return t.audience[index].IRI != nil

}

// GetAudienceIRI returns the value safely if IsAudienceIRI returned true for the specified index
func (t *TentativeAccept) GetAudienceIRI(index int) (v *url.URL) {
	return t.audience[index].IRI

}

// AppendAudienceIRI adds to the back of audience a *url.URL type
func (t *TentativeAccept) AppendAudienceIRI(v *url.URL) {
	t.audience = append(t.audience, &audienceTentativeAcceptIntermediateType{IRI: v})

}

// PrependAudienceIRI adds to the front of audience a *url.URL type
func (t *TentativeAccept) PrependAudienceIRI(v *url.URL) {
	t.audience = append([]*audienceTentativeAcceptIntermediateType{&audienceTentativeAcceptIntermediateType{IRI: v}}, t.audience...)

}

// RemoveAudienceIRI deletes the value from the specified index
func (t *TentativeAccept) RemoveAudienceIRI(index int) {
	copy(t.audience[index:], t.audience[index+1:])
	t.audience[len(t.audience)-1] = nil
	t.audience = t.audience[:len(t.audience)-1]

}

// HasUnknownAudience determines whether the call to GetUnknownAudience is safe
func (t *TentativeAccept) HasUnknownAudience() (ok bool) {
	return t.audience != nil && t.audience[0].unknown_ != nil

}

// GetUnknownAudience returns the unknown value for audience
func (t *TentativeAccept) GetUnknownAudience() (v interface{}) {
	return t.audience[0].unknown_

}

// SetUnknownAudience sets the unknown value of audience
func (t *TentativeAccept) SetUnknownAudience(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &audienceTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.audience = append(t.audience, tmp)

}

// ContentLen determines the number of elements able to be used for the IsContentString, GetContentString, and RemoveContentString functions
func (t *TentativeAccept) ContentLen() (l int) {
	return len(t.content)

}

// IsContentString determines whether the call to GetContentString is safe for the specified index
func (t *TentativeAccept) IsContentString(index int) (ok bool) {
	return t.content[index].stringName != nil

}

// GetContentString returns the value safely if IsContentString returned true for the specified index
func (t *TentativeAccept) GetContentString(index int) (v string) {
	return *t.content[index].stringName

}

// AppendContentString adds to the back of content a string type
func (t *TentativeAccept) AppendContentString(v string) {
	t.content = append(t.content, &contentTentativeAcceptIntermediateType{stringName: &v})

}

// PrependContentString adds to the front of content a string type
func (t *TentativeAccept) PrependContentString(v string) {
	t.content = append([]*contentTentativeAcceptIntermediateType{&contentTentativeAcceptIntermediateType{stringName: &v}}, t.content...)

}

// RemoveContentString deletes the value from the specified index
func (t *TentativeAccept) RemoveContentString(index int) {
	copy(t.content[index:], t.content[index+1:])
	t.content[len(t.content)-1] = nil
	t.content = t.content[:len(t.content)-1]

}

// IsContentLangString determines whether the call to GetContentLangString is safe for the specified index
func (t *TentativeAccept) IsContentLangString(index int) (ok bool) {
	return t.content[index].langString != nil

}

// GetContentLangString returns the value safely if IsContentLangString returned true for the specified index
func (t *TentativeAccept) GetContentLangString(index int) (v string) {
	return *t.content[index].langString

}

// AppendContentLangString adds to the back of content a string type
func (t *TentativeAccept) AppendContentLangString(v string) {
	t.content = append(t.content, &contentTentativeAcceptIntermediateType{langString: &v})

}

// PrependContentLangString adds to the front of content a string type
func (t *TentativeAccept) PrependContentLangString(v string) {
	t.content = append([]*contentTentativeAcceptIntermediateType{&contentTentativeAcceptIntermediateType{langString: &v}}, t.content...)

}

// RemoveContentLangString deletes the value from the specified index
func (t *TentativeAccept) RemoveContentLangString(index int) {
	copy(t.content[index:], t.content[index+1:])
	t.content[len(t.content)-1] = nil
	t.content = t.content[:len(t.content)-1]

}

// IsContentIRI determines whether the call to GetContentIRI is safe for the specified index
func (t *TentativeAccept) IsContentIRI(index int) (ok bool) {
	return t.content[index].IRI != nil

}

// GetContentIRI returns the value safely if IsContentIRI returned true for the specified index
func (t *TentativeAccept) GetContentIRI(index int) (v *url.URL) {
	return t.content[index].IRI

}

// AppendContentIRI adds to the back of content a *url.URL type
func (t *TentativeAccept) AppendContentIRI(v *url.URL) {
	t.content = append(t.content, &contentTentativeAcceptIntermediateType{IRI: v})

}

// PrependContentIRI adds to the front of content a *url.URL type
func (t *TentativeAccept) PrependContentIRI(v *url.URL) {
	t.content = append([]*contentTentativeAcceptIntermediateType{&contentTentativeAcceptIntermediateType{IRI: v}}, t.content...)

}

// RemoveContentIRI deletes the value from the specified index
func (t *TentativeAccept) RemoveContentIRI(index int) {
	copy(t.content[index:], t.content[index+1:])
	t.content[len(t.content)-1] = nil
	t.content = t.content[:len(t.content)-1]

}

// HasUnknownContent determines whether the call to GetUnknownContent is safe
func (t *TentativeAccept) HasUnknownContent() (ok bool) {
	return t.content != nil && t.content[0].unknown_ != nil

}

// GetUnknownContent returns the unknown value for content
func (t *TentativeAccept) GetUnknownContent() (v interface{}) {
	return t.content[0].unknown_

}

// SetUnknownContent sets the unknown value of content
func (t *TentativeAccept) SetUnknownContent(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &contentTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.content = append(t.content, tmp)

}

// ContentMapLanguages returns all languages for this property's language mapping, or nil if there are none.
func (t *TentativeAccept) ContentMapLanguages() (l []string) {
	if t.contentMap == nil || len(t.contentMap) == 0 {
		return nil
	}
	for k := range t.contentMap {
		l = append(l, k)
	}
	return

}

// GetContentMap retrieves the value of the property for the specified language, or an empty string if it does not exist
func (t *TentativeAccept) GetContentMap(l string) (v string) {
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
func (t *TentativeAccept) SetContentMap(l string, v string) {
	if t.contentMap == nil {
		t.contentMap = make(map[string]string)
	}
	t.contentMap[l] = v

}

// ContextLen determines the number of elements able to be used for the IsContextObject, GetContextObject, and RemoveContextObject functions
func (t *TentativeAccept) ContextLen() (l int) {
	return len(t.context)

}

// IsContextObject determines whether the call to GetContextObject is safe for the specified index
func (t *TentativeAccept) IsContextObject(index int) (ok bool) {
	return t.context[index].Object != nil

}

// GetContextObject returns the value safely if IsContextObject returned true for the specified index
func (t *TentativeAccept) GetContextObject(index int) (v ObjectType) {
	return t.context[index].Object

}

// AppendContextObject adds to the back of context a ObjectType type
func (t *TentativeAccept) AppendContextObject(v ObjectType) {
	t.context = append(t.context, &contextTentativeAcceptIntermediateType{Object: v})

}

// PrependContextObject adds to the front of context a ObjectType type
func (t *TentativeAccept) PrependContextObject(v ObjectType) {
	t.context = append([]*contextTentativeAcceptIntermediateType{&contextTentativeAcceptIntermediateType{Object: v}}, t.context...)

}

// RemoveContextObject deletes the value from the specified index
func (t *TentativeAccept) RemoveContextObject(index int) {
	copy(t.context[index:], t.context[index+1:])
	t.context[len(t.context)-1] = nil
	t.context = t.context[:len(t.context)-1]

}

// IsContextLink determines whether the call to GetContextLink is safe for the specified index
func (t *TentativeAccept) IsContextLink(index int) (ok bool) {
	return t.context[index].Link != nil

}

// GetContextLink returns the value safely if IsContextLink returned true for the specified index
func (t *TentativeAccept) GetContextLink(index int) (v LinkType) {
	return t.context[index].Link

}

// AppendContextLink adds to the back of context a LinkType type
func (t *TentativeAccept) AppendContextLink(v LinkType) {
	t.context = append(t.context, &contextTentativeAcceptIntermediateType{Link: v})

}

// PrependContextLink adds to the front of context a LinkType type
func (t *TentativeAccept) PrependContextLink(v LinkType) {
	t.context = append([]*contextTentativeAcceptIntermediateType{&contextTentativeAcceptIntermediateType{Link: v}}, t.context...)

}

// RemoveContextLink deletes the value from the specified index
func (t *TentativeAccept) RemoveContextLink(index int) {
	copy(t.context[index:], t.context[index+1:])
	t.context[len(t.context)-1] = nil
	t.context = t.context[:len(t.context)-1]

}

// IsContextIRI determines whether the call to GetContextIRI is safe for the specified index
func (t *TentativeAccept) IsContextIRI(index int) (ok bool) {
	return t.context[index].IRI != nil

}

// GetContextIRI returns the value safely if IsContextIRI returned true for the specified index
func (t *TentativeAccept) GetContextIRI(index int) (v *url.URL) {
	return t.context[index].IRI

}

// AppendContextIRI adds to the back of context a *url.URL type
func (t *TentativeAccept) AppendContextIRI(v *url.URL) {
	t.context = append(t.context, &contextTentativeAcceptIntermediateType{IRI: v})

}

// PrependContextIRI adds to the front of context a *url.URL type
func (t *TentativeAccept) PrependContextIRI(v *url.URL) {
	t.context = append([]*contextTentativeAcceptIntermediateType{&contextTentativeAcceptIntermediateType{IRI: v}}, t.context...)

}

// RemoveContextIRI deletes the value from the specified index
func (t *TentativeAccept) RemoveContextIRI(index int) {
	copy(t.context[index:], t.context[index+1:])
	t.context[len(t.context)-1] = nil
	t.context = t.context[:len(t.context)-1]

}

// HasUnknownContext determines whether the call to GetUnknownContext is safe
func (t *TentativeAccept) HasUnknownContext() (ok bool) {
	return t.context != nil && t.context[0].unknown_ != nil

}

// GetUnknownContext returns the unknown value for context
func (t *TentativeAccept) GetUnknownContext() (v interface{}) {
	return t.context[0].unknown_

}

// SetUnknownContext sets the unknown value of context
func (t *TentativeAccept) SetUnknownContext(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &contextTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.context = append(t.context, tmp)

}

// NameLen determines the number of elements able to be used for the IsNameString, GetNameString, and RemoveNameString functions
func (t *TentativeAccept) NameLen() (l int) {
	return len(t.name)

}

// IsNameString determines whether the call to GetNameString is safe for the specified index
func (t *TentativeAccept) IsNameString(index int) (ok bool) {
	return t.name[index].stringName != nil

}

// GetNameString returns the value safely if IsNameString returned true for the specified index
func (t *TentativeAccept) GetNameString(index int) (v string) {
	return *t.name[index].stringName

}

// AppendNameString adds to the back of name a string type
func (t *TentativeAccept) AppendNameString(v string) {
	t.name = append(t.name, &nameTentativeAcceptIntermediateType{stringName: &v})

}

// PrependNameString adds to the front of name a string type
func (t *TentativeAccept) PrependNameString(v string) {
	t.name = append([]*nameTentativeAcceptIntermediateType{&nameTentativeAcceptIntermediateType{stringName: &v}}, t.name...)

}

// RemoveNameString deletes the value from the specified index
func (t *TentativeAccept) RemoveNameString(index int) {
	copy(t.name[index:], t.name[index+1:])
	t.name[len(t.name)-1] = nil
	t.name = t.name[:len(t.name)-1]

}

// IsNameLangString determines whether the call to GetNameLangString is safe for the specified index
func (t *TentativeAccept) IsNameLangString(index int) (ok bool) {
	return t.name[index].langString != nil

}

// GetNameLangString returns the value safely if IsNameLangString returned true for the specified index
func (t *TentativeAccept) GetNameLangString(index int) (v string) {
	return *t.name[index].langString

}

// AppendNameLangString adds to the back of name a string type
func (t *TentativeAccept) AppendNameLangString(v string) {
	t.name = append(t.name, &nameTentativeAcceptIntermediateType{langString: &v})

}

// PrependNameLangString adds to the front of name a string type
func (t *TentativeAccept) PrependNameLangString(v string) {
	t.name = append([]*nameTentativeAcceptIntermediateType{&nameTentativeAcceptIntermediateType{langString: &v}}, t.name...)

}

// RemoveNameLangString deletes the value from the specified index
func (t *TentativeAccept) RemoveNameLangString(index int) {
	copy(t.name[index:], t.name[index+1:])
	t.name[len(t.name)-1] = nil
	t.name = t.name[:len(t.name)-1]

}

// IsNameIRI determines whether the call to GetNameIRI is safe for the specified index
func (t *TentativeAccept) IsNameIRI(index int) (ok bool) {
	return t.name[index].IRI != nil

}

// GetNameIRI returns the value safely if IsNameIRI returned true for the specified index
func (t *TentativeAccept) GetNameIRI(index int) (v *url.URL) {
	return t.name[index].IRI

}

// AppendNameIRI adds to the back of name a *url.URL type
func (t *TentativeAccept) AppendNameIRI(v *url.URL) {
	t.name = append(t.name, &nameTentativeAcceptIntermediateType{IRI: v})

}

// PrependNameIRI adds to the front of name a *url.URL type
func (t *TentativeAccept) PrependNameIRI(v *url.URL) {
	t.name = append([]*nameTentativeAcceptIntermediateType{&nameTentativeAcceptIntermediateType{IRI: v}}, t.name...)

}

// RemoveNameIRI deletes the value from the specified index
func (t *TentativeAccept) RemoveNameIRI(index int) {
	copy(t.name[index:], t.name[index+1:])
	t.name[len(t.name)-1] = nil
	t.name = t.name[:len(t.name)-1]

}

// HasUnknownName determines whether the call to GetUnknownName is safe
func (t *TentativeAccept) HasUnknownName() (ok bool) {
	return t.name != nil && t.name[0].unknown_ != nil

}

// GetUnknownName returns the unknown value for name
func (t *TentativeAccept) GetUnknownName() (v interface{}) {
	return t.name[0].unknown_

}

// SetUnknownName sets the unknown value of name
func (t *TentativeAccept) SetUnknownName(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &nameTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.name = append(t.name, tmp)

}

// NameMapLanguages returns all languages for this property's language mapping, or nil if there are none.
func (t *TentativeAccept) NameMapLanguages() (l []string) {
	if t.nameMap == nil || len(t.nameMap) == 0 {
		return nil
	}
	for k := range t.nameMap {
		l = append(l, k)
	}
	return

}

// GetNameMap retrieves the value of the property for the specified language, or an empty string if it does not exist
func (t *TentativeAccept) GetNameMap(l string) (v string) {
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
func (t *TentativeAccept) SetNameMap(l string, v string) {
	if t.nameMap == nil {
		t.nameMap = make(map[string]string)
	}
	t.nameMap[l] = v

}

// IsEndTime determines whether the call to GetEndTime is safe
func (t *TentativeAccept) IsEndTime() (ok bool) {
	return t.endTime != nil && t.endTime.dateTime != nil

}

// GetEndTime returns the value safely if IsEndTime returned true
func (t *TentativeAccept) GetEndTime() (v time.Time) {
	return *t.endTime.dateTime

}

// SetEndTime sets the value of endTime to be of time.Time type
func (t *TentativeAccept) SetEndTime(v time.Time) {
	t.endTime = &endTimeTentativeAcceptIntermediateType{dateTime: &v}

}

// IsEndTimeIRI determines whether the call to GetEndTimeIRI is safe
func (t *TentativeAccept) IsEndTimeIRI() (ok bool) {
	return t.endTime != nil && t.endTime.IRI != nil

}

// GetEndTimeIRI returns the value safely if IsEndTimeIRI returned true
func (t *TentativeAccept) GetEndTimeIRI() (v *url.URL) {
	return t.endTime.IRI

}

// SetEndTimeIRI sets the value of endTime to be of *url.URL type
func (t *TentativeAccept) SetEndTimeIRI(v *url.URL) {
	t.endTime = &endTimeTentativeAcceptIntermediateType{IRI: v}

}

// HasUnknownEndTime determines whether the call to GetUnknownEndTime is safe
func (t *TentativeAccept) HasUnknownEndTime() (ok bool) {
	return t.endTime != nil && t.endTime.unknown_ != nil

}

// GetUnknownEndTime returns the unknown value for endTime
func (t *TentativeAccept) GetUnknownEndTime() (v interface{}) {
	return t.endTime.unknown_

}

// SetUnknownEndTime sets the unknown value of endTime
func (t *TentativeAccept) SetUnknownEndTime(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &endTimeTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.endTime = tmp

}

// GeneratorLen determines the number of elements able to be used for the IsGeneratorObject, GetGeneratorObject, and RemoveGeneratorObject functions
func (t *TentativeAccept) GeneratorLen() (l int) {
	return len(t.generator)

}

// IsGeneratorObject determines whether the call to GetGeneratorObject is safe for the specified index
func (t *TentativeAccept) IsGeneratorObject(index int) (ok bool) {
	return t.generator[index].Object != nil

}

// GetGeneratorObject returns the value safely if IsGeneratorObject returned true for the specified index
func (t *TentativeAccept) GetGeneratorObject(index int) (v ObjectType) {
	return t.generator[index].Object

}

// AppendGeneratorObject adds to the back of generator a ObjectType type
func (t *TentativeAccept) AppendGeneratorObject(v ObjectType) {
	t.generator = append(t.generator, &generatorTentativeAcceptIntermediateType{Object: v})

}

// PrependGeneratorObject adds to the front of generator a ObjectType type
func (t *TentativeAccept) PrependGeneratorObject(v ObjectType) {
	t.generator = append([]*generatorTentativeAcceptIntermediateType{&generatorTentativeAcceptIntermediateType{Object: v}}, t.generator...)

}

// RemoveGeneratorObject deletes the value from the specified index
func (t *TentativeAccept) RemoveGeneratorObject(index int) {
	copy(t.generator[index:], t.generator[index+1:])
	t.generator[len(t.generator)-1] = nil
	t.generator = t.generator[:len(t.generator)-1]

}

// IsGeneratorLink determines whether the call to GetGeneratorLink is safe for the specified index
func (t *TentativeAccept) IsGeneratorLink(index int) (ok bool) {
	return t.generator[index].Link != nil

}

// GetGeneratorLink returns the value safely if IsGeneratorLink returned true for the specified index
func (t *TentativeAccept) GetGeneratorLink(index int) (v LinkType) {
	return t.generator[index].Link

}

// AppendGeneratorLink adds to the back of generator a LinkType type
func (t *TentativeAccept) AppendGeneratorLink(v LinkType) {
	t.generator = append(t.generator, &generatorTentativeAcceptIntermediateType{Link: v})

}

// PrependGeneratorLink adds to the front of generator a LinkType type
func (t *TentativeAccept) PrependGeneratorLink(v LinkType) {
	t.generator = append([]*generatorTentativeAcceptIntermediateType{&generatorTentativeAcceptIntermediateType{Link: v}}, t.generator...)

}

// RemoveGeneratorLink deletes the value from the specified index
func (t *TentativeAccept) RemoveGeneratorLink(index int) {
	copy(t.generator[index:], t.generator[index+1:])
	t.generator[len(t.generator)-1] = nil
	t.generator = t.generator[:len(t.generator)-1]

}

// IsGeneratorIRI determines whether the call to GetGeneratorIRI is safe for the specified index
func (t *TentativeAccept) IsGeneratorIRI(index int) (ok bool) {
	return t.generator[index].IRI != nil

}

// GetGeneratorIRI returns the value safely if IsGeneratorIRI returned true for the specified index
func (t *TentativeAccept) GetGeneratorIRI(index int) (v *url.URL) {
	return t.generator[index].IRI

}

// AppendGeneratorIRI adds to the back of generator a *url.URL type
func (t *TentativeAccept) AppendGeneratorIRI(v *url.URL) {
	t.generator = append(t.generator, &generatorTentativeAcceptIntermediateType{IRI: v})

}

// PrependGeneratorIRI adds to the front of generator a *url.URL type
func (t *TentativeAccept) PrependGeneratorIRI(v *url.URL) {
	t.generator = append([]*generatorTentativeAcceptIntermediateType{&generatorTentativeAcceptIntermediateType{IRI: v}}, t.generator...)

}

// RemoveGeneratorIRI deletes the value from the specified index
func (t *TentativeAccept) RemoveGeneratorIRI(index int) {
	copy(t.generator[index:], t.generator[index+1:])
	t.generator[len(t.generator)-1] = nil
	t.generator = t.generator[:len(t.generator)-1]

}

// HasUnknownGenerator determines whether the call to GetUnknownGenerator is safe
func (t *TentativeAccept) HasUnknownGenerator() (ok bool) {
	return t.generator != nil && t.generator[0].unknown_ != nil

}

// GetUnknownGenerator returns the unknown value for generator
func (t *TentativeAccept) GetUnknownGenerator() (v interface{}) {
	return t.generator[0].unknown_

}

// SetUnknownGenerator sets the unknown value of generator
func (t *TentativeAccept) SetUnknownGenerator(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &generatorTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.generator = append(t.generator, tmp)

}

// IconLen determines the number of elements able to be used for the IsIconImage, GetIconImage, and RemoveIconImage functions
func (t *TentativeAccept) IconLen() (l int) {
	return len(t.icon)

}

// IsIconImage determines whether the call to GetIconImage is safe for the specified index
func (t *TentativeAccept) IsIconImage(index int) (ok bool) {
	return t.icon[index].Image != nil

}

// GetIconImage returns the value safely if IsIconImage returned true for the specified index
func (t *TentativeAccept) GetIconImage(index int) (v ImageType) {
	return t.icon[index].Image

}

// AppendIconImage adds to the back of icon a ImageType type
func (t *TentativeAccept) AppendIconImage(v ImageType) {
	t.icon = append(t.icon, &iconTentativeAcceptIntermediateType{Image: v})

}

// PrependIconImage adds to the front of icon a ImageType type
func (t *TentativeAccept) PrependIconImage(v ImageType) {
	t.icon = append([]*iconTentativeAcceptIntermediateType{&iconTentativeAcceptIntermediateType{Image: v}}, t.icon...)

}

// RemoveIconImage deletes the value from the specified index
func (t *TentativeAccept) RemoveIconImage(index int) {
	copy(t.icon[index:], t.icon[index+1:])
	t.icon[len(t.icon)-1] = nil
	t.icon = t.icon[:len(t.icon)-1]

}

// IsIconLink determines whether the call to GetIconLink is safe for the specified index
func (t *TentativeAccept) IsIconLink(index int) (ok bool) {
	return t.icon[index].Link != nil

}

// GetIconLink returns the value safely if IsIconLink returned true for the specified index
func (t *TentativeAccept) GetIconLink(index int) (v LinkType) {
	return t.icon[index].Link

}

// AppendIconLink adds to the back of icon a LinkType type
func (t *TentativeAccept) AppendIconLink(v LinkType) {
	t.icon = append(t.icon, &iconTentativeAcceptIntermediateType{Link: v})

}

// PrependIconLink adds to the front of icon a LinkType type
func (t *TentativeAccept) PrependIconLink(v LinkType) {
	t.icon = append([]*iconTentativeAcceptIntermediateType{&iconTentativeAcceptIntermediateType{Link: v}}, t.icon...)

}

// RemoveIconLink deletes the value from the specified index
func (t *TentativeAccept) RemoveIconLink(index int) {
	copy(t.icon[index:], t.icon[index+1:])
	t.icon[len(t.icon)-1] = nil
	t.icon = t.icon[:len(t.icon)-1]

}

// IsIconIRI determines whether the call to GetIconIRI is safe for the specified index
func (t *TentativeAccept) IsIconIRI(index int) (ok bool) {
	return t.icon[index].IRI != nil

}

// GetIconIRI returns the value safely if IsIconIRI returned true for the specified index
func (t *TentativeAccept) GetIconIRI(index int) (v *url.URL) {
	return t.icon[index].IRI

}

// AppendIconIRI adds to the back of icon a *url.URL type
func (t *TentativeAccept) AppendIconIRI(v *url.URL) {
	t.icon = append(t.icon, &iconTentativeAcceptIntermediateType{IRI: v})

}

// PrependIconIRI adds to the front of icon a *url.URL type
func (t *TentativeAccept) PrependIconIRI(v *url.URL) {
	t.icon = append([]*iconTentativeAcceptIntermediateType{&iconTentativeAcceptIntermediateType{IRI: v}}, t.icon...)

}

// RemoveIconIRI deletes the value from the specified index
func (t *TentativeAccept) RemoveIconIRI(index int) {
	copy(t.icon[index:], t.icon[index+1:])
	t.icon[len(t.icon)-1] = nil
	t.icon = t.icon[:len(t.icon)-1]

}

// HasUnknownIcon determines whether the call to GetUnknownIcon is safe
func (t *TentativeAccept) HasUnknownIcon() (ok bool) {
	return t.icon != nil && t.icon[0].unknown_ != nil

}

// GetUnknownIcon returns the unknown value for icon
func (t *TentativeAccept) GetUnknownIcon() (v interface{}) {
	return t.icon[0].unknown_

}

// SetUnknownIcon sets the unknown value of icon
func (t *TentativeAccept) SetUnknownIcon(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &iconTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.icon = append(t.icon, tmp)

}

// HasId determines whether the call to GetId is safe
func (t *TentativeAccept) HasId() (ok bool) {
	return t.id != nil

}

// GetId returns the value for id
func (t *TentativeAccept) GetId() (v *url.URL) {
	return t.id

}

// SetId sets the value of id
func (t *TentativeAccept) SetId(v *url.URL) {
	t.id = v

}

// HasUnknownId determines whether the call to GetUnknownId is safe
func (t *TentativeAccept) HasUnknownId() (ok bool) {
	return t.unknown_ != nil && t.unknown_["id"] != nil

}

// GetUnknownId returns the unknown value for id
func (t *TentativeAccept) GetUnknownId() (v interface{}) {
	return t.unknown_["id"]

}

// SetUnknownId sets the unknown value of id
func (t *TentativeAccept) SetUnknownId(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["id"] = i

}

// ImageLen determines the number of elements able to be used for the IsImageImage, GetImageImage, and RemoveImageImage functions
func (t *TentativeAccept) ImageLen() (l int) {
	return len(t.image)

}

// IsImageImage determines whether the call to GetImageImage is safe for the specified index
func (t *TentativeAccept) IsImageImage(index int) (ok bool) {
	return t.image[index].Image != nil

}

// GetImageImage returns the value safely if IsImageImage returned true for the specified index
func (t *TentativeAccept) GetImageImage(index int) (v ImageType) {
	return t.image[index].Image

}

// AppendImageImage adds to the back of image a ImageType type
func (t *TentativeAccept) AppendImageImage(v ImageType) {
	t.image = append(t.image, &imageTentativeAcceptIntermediateType{Image: v})

}

// PrependImageImage adds to the front of image a ImageType type
func (t *TentativeAccept) PrependImageImage(v ImageType) {
	t.image = append([]*imageTentativeAcceptIntermediateType{&imageTentativeAcceptIntermediateType{Image: v}}, t.image...)

}

// RemoveImageImage deletes the value from the specified index
func (t *TentativeAccept) RemoveImageImage(index int) {
	copy(t.image[index:], t.image[index+1:])
	t.image[len(t.image)-1] = nil
	t.image = t.image[:len(t.image)-1]

}

// IsImageLink determines whether the call to GetImageLink is safe for the specified index
func (t *TentativeAccept) IsImageLink(index int) (ok bool) {
	return t.image[index].Link != nil

}

// GetImageLink returns the value safely if IsImageLink returned true for the specified index
func (t *TentativeAccept) GetImageLink(index int) (v LinkType) {
	return t.image[index].Link

}

// AppendImageLink adds to the back of image a LinkType type
func (t *TentativeAccept) AppendImageLink(v LinkType) {
	t.image = append(t.image, &imageTentativeAcceptIntermediateType{Link: v})

}

// PrependImageLink adds to the front of image a LinkType type
func (t *TentativeAccept) PrependImageLink(v LinkType) {
	t.image = append([]*imageTentativeAcceptIntermediateType{&imageTentativeAcceptIntermediateType{Link: v}}, t.image...)

}

// RemoveImageLink deletes the value from the specified index
func (t *TentativeAccept) RemoveImageLink(index int) {
	copy(t.image[index:], t.image[index+1:])
	t.image[len(t.image)-1] = nil
	t.image = t.image[:len(t.image)-1]

}

// IsImageIRI determines whether the call to GetImageIRI is safe for the specified index
func (t *TentativeAccept) IsImageIRI(index int) (ok bool) {
	return t.image[index].IRI != nil

}

// GetImageIRI returns the value safely if IsImageIRI returned true for the specified index
func (t *TentativeAccept) GetImageIRI(index int) (v *url.URL) {
	return t.image[index].IRI

}

// AppendImageIRI adds to the back of image a *url.URL type
func (t *TentativeAccept) AppendImageIRI(v *url.URL) {
	t.image = append(t.image, &imageTentativeAcceptIntermediateType{IRI: v})

}

// PrependImageIRI adds to the front of image a *url.URL type
func (t *TentativeAccept) PrependImageIRI(v *url.URL) {
	t.image = append([]*imageTentativeAcceptIntermediateType{&imageTentativeAcceptIntermediateType{IRI: v}}, t.image...)

}

// RemoveImageIRI deletes the value from the specified index
func (t *TentativeAccept) RemoveImageIRI(index int) {
	copy(t.image[index:], t.image[index+1:])
	t.image[len(t.image)-1] = nil
	t.image = t.image[:len(t.image)-1]

}

// HasUnknownImage determines whether the call to GetUnknownImage is safe
func (t *TentativeAccept) HasUnknownImage() (ok bool) {
	return t.image != nil && t.image[0].unknown_ != nil

}

// GetUnknownImage returns the unknown value for image
func (t *TentativeAccept) GetUnknownImage() (v interface{}) {
	return t.image[0].unknown_

}

// SetUnknownImage sets the unknown value of image
func (t *TentativeAccept) SetUnknownImage(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &imageTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.image = append(t.image, tmp)

}

// InReplyToLen determines the number of elements able to be used for the IsInReplyToObject, GetInReplyToObject, and RemoveInReplyToObject functions
func (t *TentativeAccept) InReplyToLen() (l int) {
	return len(t.inReplyTo)

}

// IsInReplyToObject determines whether the call to GetInReplyToObject is safe for the specified index
func (t *TentativeAccept) IsInReplyToObject(index int) (ok bool) {
	return t.inReplyTo[index].Object != nil

}

// GetInReplyToObject returns the value safely if IsInReplyToObject returned true for the specified index
func (t *TentativeAccept) GetInReplyToObject(index int) (v ObjectType) {
	return t.inReplyTo[index].Object

}

// AppendInReplyToObject adds to the back of inReplyTo a ObjectType type
func (t *TentativeAccept) AppendInReplyToObject(v ObjectType) {
	t.inReplyTo = append(t.inReplyTo, &inReplyToTentativeAcceptIntermediateType{Object: v})

}

// PrependInReplyToObject adds to the front of inReplyTo a ObjectType type
func (t *TentativeAccept) PrependInReplyToObject(v ObjectType) {
	t.inReplyTo = append([]*inReplyToTentativeAcceptIntermediateType{&inReplyToTentativeAcceptIntermediateType{Object: v}}, t.inReplyTo...)

}

// RemoveInReplyToObject deletes the value from the specified index
func (t *TentativeAccept) RemoveInReplyToObject(index int) {
	copy(t.inReplyTo[index:], t.inReplyTo[index+1:])
	t.inReplyTo[len(t.inReplyTo)-1] = nil
	t.inReplyTo = t.inReplyTo[:len(t.inReplyTo)-1]

}

// IsInReplyToLink determines whether the call to GetInReplyToLink is safe for the specified index
func (t *TentativeAccept) IsInReplyToLink(index int) (ok bool) {
	return t.inReplyTo[index].Link != nil

}

// GetInReplyToLink returns the value safely if IsInReplyToLink returned true for the specified index
func (t *TentativeAccept) GetInReplyToLink(index int) (v LinkType) {
	return t.inReplyTo[index].Link

}

// AppendInReplyToLink adds to the back of inReplyTo a LinkType type
func (t *TentativeAccept) AppendInReplyToLink(v LinkType) {
	t.inReplyTo = append(t.inReplyTo, &inReplyToTentativeAcceptIntermediateType{Link: v})

}

// PrependInReplyToLink adds to the front of inReplyTo a LinkType type
func (t *TentativeAccept) PrependInReplyToLink(v LinkType) {
	t.inReplyTo = append([]*inReplyToTentativeAcceptIntermediateType{&inReplyToTentativeAcceptIntermediateType{Link: v}}, t.inReplyTo...)

}

// RemoveInReplyToLink deletes the value from the specified index
func (t *TentativeAccept) RemoveInReplyToLink(index int) {
	copy(t.inReplyTo[index:], t.inReplyTo[index+1:])
	t.inReplyTo[len(t.inReplyTo)-1] = nil
	t.inReplyTo = t.inReplyTo[:len(t.inReplyTo)-1]

}

// IsInReplyToIRI determines whether the call to GetInReplyToIRI is safe for the specified index
func (t *TentativeAccept) IsInReplyToIRI(index int) (ok bool) {
	return t.inReplyTo[index].IRI != nil

}

// GetInReplyToIRI returns the value safely if IsInReplyToIRI returned true for the specified index
func (t *TentativeAccept) GetInReplyToIRI(index int) (v *url.URL) {
	return t.inReplyTo[index].IRI

}

// AppendInReplyToIRI adds to the back of inReplyTo a *url.URL type
func (t *TentativeAccept) AppendInReplyToIRI(v *url.URL) {
	t.inReplyTo = append(t.inReplyTo, &inReplyToTentativeAcceptIntermediateType{IRI: v})

}

// PrependInReplyToIRI adds to the front of inReplyTo a *url.URL type
func (t *TentativeAccept) PrependInReplyToIRI(v *url.URL) {
	t.inReplyTo = append([]*inReplyToTentativeAcceptIntermediateType{&inReplyToTentativeAcceptIntermediateType{IRI: v}}, t.inReplyTo...)

}

// RemoveInReplyToIRI deletes the value from the specified index
func (t *TentativeAccept) RemoveInReplyToIRI(index int) {
	copy(t.inReplyTo[index:], t.inReplyTo[index+1:])
	t.inReplyTo[len(t.inReplyTo)-1] = nil
	t.inReplyTo = t.inReplyTo[:len(t.inReplyTo)-1]

}

// HasUnknownInReplyTo determines whether the call to GetUnknownInReplyTo is safe
func (t *TentativeAccept) HasUnknownInReplyTo() (ok bool) {
	return t.inReplyTo != nil && t.inReplyTo[0].unknown_ != nil

}

// GetUnknownInReplyTo returns the unknown value for inReplyTo
func (t *TentativeAccept) GetUnknownInReplyTo() (v interface{}) {
	return t.inReplyTo[0].unknown_

}

// SetUnknownInReplyTo sets the unknown value of inReplyTo
func (t *TentativeAccept) SetUnknownInReplyTo(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &inReplyToTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.inReplyTo = append(t.inReplyTo, tmp)

}

// LocationLen determines the number of elements able to be used for the IsLocationObject, GetLocationObject, and RemoveLocationObject functions
func (t *TentativeAccept) LocationLen() (l int) {
	return len(t.location)

}

// IsLocationObject determines whether the call to GetLocationObject is safe for the specified index
func (t *TentativeAccept) IsLocationObject(index int) (ok bool) {
	return t.location[index].Object != nil

}

// GetLocationObject returns the value safely if IsLocationObject returned true for the specified index
func (t *TentativeAccept) GetLocationObject(index int) (v ObjectType) {
	return t.location[index].Object

}

// AppendLocationObject adds to the back of location a ObjectType type
func (t *TentativeAccept) AppendLocationObject(v ObjectType) {
	t.location = append(t.location, &locationTentativeAcceptIntermediateType{Object: v})

}

// PrependLocationObject adds to the front of location a ObjectType type
func (t *TentativeAccept) PrependLocationObject(v ObjectType) {
	t.location = append([]*locationTentativeAcceptIntermediateType{&locationTentativeAcceptIntermediateType{Object: v}}, t.location...)

}

// RemoveLocationObject deletes the value from the specified index
func (t *TentativeAccept) RemoveLocationObject(index int) {
	copy(t.location[index:], t.location[index+1:])
	t.location[len(t.location)-1] = nil
	t.location = t.location[:len(t.location)-1]

}

// IsLocationLink determines whether the call to GetLocationLink is safe for the specified index
func (t *TentativeAccept) IsLocationLink(index int) (ok bool) {
	return t.location[index].Link != nil

}

// GetLocationLink returns the value safely if IsLocationLink returned true for the specified index
func (t *TentativeAccept) GetLocationLink(index int) (v LinkType) {
	return t.location[index].Link

}

// AppendLocationLink adds to the back of location a LinkType type
func (t *TentativeAccept) AppendLocationLink(v LinkType) {
	t.location = append(t.location, &locationTentativeAcceptIntermediateType{Link: v})

}

// PrependLocationLink adds to the front of location a LinkType type
func (t *TentativeAccept) PrependLocationLink(v LinkType) {
	t.location = append([]*locationTentativeAcceptIntermediateType{&locationTentativeAcceptIntermediateType{Link: v}}, t.location...)

}

// RemoveLocationLink deletes the value from the specified index
func (t *TentativeAccept) RemoveLocationLink(index int) {
	copy(t.location[index:], t.location[index+1:])
	t.location[len(t.location)-1] = nil
	t.location = t.location[:len(t.location)-1]

}

// IsLocationIRI determines whether the call to GetLocationIRI is safe for the specified index
func (t *TentativeAccept) IsLocationIRI(index int) (ok bool) {
	return t.location[index].IRI != nil

}

// GetLocationIRI returns the value safely if IsLocationIRI returned true for the specified index
func (t *TentativeAccept) GetLocationIRI(index int) (v *url.URL) {
	return t.location[index].IRI

}

// AppendLocationIRI adds to the back of location a *url.URL type
func (t *TentativeAccept) AppendLocationIRI(v *url.URL) {
	t.location = append(t.location, &locationTentativeAcceptIntermediateType{IRI: v})

}

// PrependLocationIRI adds to the front of location a *url.URL type
func (t *TentativeAccept) PrependLocationIRI(v *url.URL) {
	t.location = append([]*locationTentativeAcceptIntermediateType{&locationTentativeAcceptIntermediateType{IRI: v}}, t.location...)

}

// RemoveLocationIRI deletes the value from the specified index
func (t *TentativeAccept) RemoveLocationIRI(index int) {
	copy(t.location[index:], t.location[index+1:])
	t.location[len(t.location)-1] = nil
	t.location = t.location[:len(t.location)-1]

}

// HasUnknownLocation determines whether the call to GetUnknownLocation is safe
func (t *TentativeAccept) HasUnknownLocation() (ok bool) {
	return t.location != nil && t.location[0].unknown_ != nil

}

// GetUnknownLocation returns the unknown value for location
func (t *TentativeAccept) GetUnknownLocation() (v interface{}) {
	return t.location[0].unknown_

}

// SetUnknownLocation sets the unknown value of location
func (t *TentativeAccept) SetUnknownLocation(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &locationTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.location = append(t.location, tmp)

}

// PreviewLen determines the number of elements able to be used for the IsPreviewObject, GetPreviewObject, and RemovePreviewObject functions
func (t *TentativeAccept) PreviewLen() (l int) {
	return len(t.preview)

}

// IsPreviewObject determines whether the call to GetPreviewObject is safe for the specified index
func (t *TentativeAccept) IsPreviewObject(index int) (ok bool) {
	return t.preview[index].Object != nil

}

// GetPreviewObject returns the value safely if IsPreviewObject returned true for the specified index
func (t *TentativeAccept) GetPreviewObject(index int) (v ObjectType) {
	return t.preview[index].Object

}

// AppendPreviewObject adds to the back of preview a ObjectType type
func (t *TentativeAccept) AppendPreviewObject(v ObjectType) {
	t.preview = append(t.preview, &previewTentativeAcceptIntermediateType{Object: v})

}

// PrependPreviewObject adds to the front of preview a ObjectType type
func (t *TentativeAccept) PrependPreviewObject(v ObjectType) {
	t.preview = append([]*previewTentativeAcceptIntermediateType{&previewTentativeAcceptIntermediateType{Object: v}}, t.preview...)

}

// RemovePreviewObject deletes the value from the specified index
func (t *TentativeAccept) RemovePreviewObject(index int) {
	copy(t.preview[index:], t.preview[index+1:])
	t.preview[len(t.preview)-1] = nil
	t.preview = t.preview[:len(t.preview)-1]

}

// IsPreviewLink determines whether the call to GetPreviewLink is safe for the specified index
func (t *TentativeAccept) IsPreviewLink(index int) (ok bool) {
	return t.preview[index].Link != nil

}

// GetPreviewLink returns the value safely if IsPreviewLink returned true for the specified index
func (t *TentativeAccept) GetPreviewLink(index int) (v LinkType) {
	return t.preview[index].Link

}

// AppendPreviewLink adds to the back of preview a LinkType type
func (t *TentativeAccept) AppendPreviewLink(v LinkType) {
	t.preview = append(t.preview, &previewTentativeAcceptIntermediateType{Link: v})

}

// PrependPreviewLink adds to the front of preview a LinkType type
func (t *TentativeAccept) PrependPreviewLink(v LinkType) {
	t.preview = append([]*previewTentativeAcceptIntermediateType{&previewTentativeAcceptIntermediateType{Link: v}}, t.preview...)

}

// RemovePreviewLink deletes the value from the specified index
func (t *TentativeAccept) RemovePreviewLink(index int) {
	copy(t.preview[index:], t.preview[index+1:])
	t.preview[len(t.preview)-1] = nil
	t.preview = t.preview[:len(t.preview)-1]

}

// IsPreviewIRI determines whether the call to GetPreviewIRI is safe for the specified index
func (t *TentativeAccept) IsPreviewIRI(index int) (ok bool) {
	return t.preview[index].IRI != nil

}

// GetPreviewIRI returns the value safely if IsPreviewIRI returned true for the specified index
func (t *TentativeAccept) GetPreviewIRI(index int) (v *url.URL) {
	return t.preview[index].IRI

}

// AppendPreviewIRI adds to the back of preview a *url.URL type
func (t *TentativeAccept) AppendPreviewIRI(v *url.URL) {
	t.preview = append(t.preview, &previewTentativeAcceptIntermediateType{IRI: v})

}

// PrependPreviewIRI adds to the front of preview a *url.URL type
func (t *TentativeAccept) PrependPreviewIRI(v *url.URL) {
	t.preview = append([]*previewTentativeAcceptIntermediateType{&previewTentativeAcceptIntermediateType{IRI: v}}, t.preview...)

}

// RemovePreviewIRI deletes the value from the specified index
func (t *TentativeAccept) RemovePreviewIRI(index int) {
	copy(t.preview[index:], t.preview[index+1:])
	t.preview[len(t.preview)-1] = nil
	t.preview = t.preview[:len(t.preview)-1]

}

// HasUnknownPreview determines whether the call to GetUnknownPreview is safe
func (t *TentativeAccept) HasUnknownPreview() (ok bool) {
	return t.preview != nil && t.preview[0].unknown_ != nil

}

// GetUnknownPreview returns the unknown value for preview
func (t *TentativeAccept) GetUnknownPreview() (v interface{}) {
	return t.preview[0].unknown_

}

// SetUnknownPreview sets the unknown value of preview
func (t *TentativeAccept) SetUnknownPreview(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &previewTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.preview = append(t.preview, tmp)

}

// IsPublished determines whether the call to GetPublished is safe
func (t *TentativeAccept) IsPublished() (ok bool) {
	return t.published != nil && t.published.dateTime != nil

}

// GetPublished returns the value safely if IsPublished returned true
func (t *TentativeAccept) GetPublished() (v time.Time) {
	return *t.published.dateTime

}

// SetPublished sets the value of published to be of time.Time type
func (t *TentativeAccept) SetPublished(v time.Time) {
	t.published = &publishedTentativeAcceptIntermediateType{dateTime: &v}

}

// IsPublishedIRI determines whether the call to GetPublishedIRI is safe
func (t *TentativeAccept) IsPublishedIRI() (ok bool) {
	return t.published != nil && t.published.IRI != nil

}

// GetPublishedIRI returns the value safely if IsPublishedIRI returned true
func (t *TentativeAccept) GetPublishedIRI() (v *url.URL) {
	return t.published.IRI

}

// SetPublishedIRI sets the value of published to be of *url.URL type
func (t *TentativeAccept) SetPublishedIRI(v *url.URL) {
	t.published = &publishedTentativeAcceptIntermediateType{IRI: v}

}

// HasUnknownPublished determines whether the call to GetUnknownPublished is safe
func (t *TentativeAccept) HasUnknownPublished() (ok bool) {
	return t.published != nil && t.published.unknown_ != nil

}

// GetUnknownPublished returns the unknown value for published
func (t *TentativeAccept) GetUnknownPublished() (v interface{}) {
	return t.published.unknown_

}

// SetUnknownPublished sets the unknown value of published
func (t *TentativeAccept) SetUnknownPublished(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &publishedTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.published = tmp

}

// IsReplies determines whether the call to GetReplies is safe
func (t *TentativeAccept) IsReplies() (ok bool) {
	return t.replies != nil && t.replies.Collection != nil

}

// GetReplies returns the value safely if IsReplies returned true
func (t *TentativeAccept) GetReplies() (v CollectionType) {
	return t.replies.Collection

}

// SetReplies sets the value of replies to be of CollectionType type
func (t *TentativeAccept) SetReplies(v CollectionType) {
	t.replies = &repliesTentativeAcceptIntermediateType{Collection: v}

}

// IsRepliesIRI determines whether the call to GetRepliesIRI is safe
func (t *TentativeAccept) IsRepliesIRI() (ok bool) {
	return t.replies != nil && t.replies.IRI != nil

}

// GetRepliesIRI returns the value safely if IsRepliesIRI returned true
func (t *TentativeAccept) GetRepliesIRI() (v *url.URL) {
	return t.replies.IRI

}

// SetRepliesIRI sets the value of replies to be of *url.URL type
func (t *TentativeAccept) SetRepliesIRI(v *url.URL) {
	t.replies = &repliesTentativeAcceptIntermediateType{IRI: v}

}

// HasUnknownReplies determines whether the call to GetUnknownReplies is safe
func (t *TentativeAccept) HasUnknownReplies() (ok bool) {
	return t.replies != nil && t.replies.unknown_ != nil

}

// GetUnknownReplies returns the unknown value for replies
func (t *TentativeAccept) GetUnknownReplies() (v interface{}) {
	return t.replies.unknown_

}

// SetUnknownReplies sets the unknown value of replies
func (t *TentativeAccept) SetUnknownReplies(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &repliesTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.replies = tmp

}

// IsStartTime determines whether the call to GetStartTime is safe
func (t *TentativeAccept) IsStartTime() (ok bool) {
	return t.startTime != nil && t.startTime.dateTime != nil

}

// GetStartTime returns the value safely if IsStartTime returned true
func (t *TentativeAccept) GetStartTime() (v time.Time) {
	return *t.startTime.dateTime

}

// SetStartTime sets the value of startTime to be of time.Time type
func (t *TentativeAccept) SetStartTime(v time.Time) {
	t.startTime = &startTimeTentativeAcceptIntermediateType{dateTime: &v}

}

// IsStartTimeIRI determines whether the call to GetStartTimeIRI is safe
func (t *TentativeAccept) IsStartTimeIRI() (ok bool) {
	return t.startTime != nil && t.startTime.IRI != nil

}

// GetStartTimeIRI returns the value safely if IsStartTimeIRI returned true
func (t *TentativeAccept) GetStartTimeIRI() (v *url.URL) {
	return t.startTime.IRI

}

// SetStartTimeIRI sets the value of startTime to be of *url.URL type
func (t *TentativeAccept) SetStartTimeIRI(v *url.URL) {
	t.startTime = &startTimeTentativeAcceptIntermediateType{IRI: v}

}

// HasUnknownStartTime determines whether the call to GetUnknownStartTime is safe
func (t *TentativeAccept) HasUnknownStartTime() (ok bool) {
	return t.startTime != nil && t.startTime.unknown_ != nil

}

// GetUnknownStartTime returns the unknown value for startTime
func (t *TentativeAccept) GetUnknownStartTime() (v interface{}) {
	return t.startTime.unknown_

}

// SetUnknownStartTime sets the unknown value of startTime
func (t *TentativeAccept) SetUnknownStartTime(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &startTimeTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.startTime = tmp

}

// SummaryLen determines the number of elements able to be used for the IsSummaryString, GetSummaryString, and RemoveSummaryString functions
func (t *TentativeAccept) SummaryLen() (l int) {
	return len(t.summary)

}

// IsSummaryString determines whether the call to GetSummaryString is safe for the specified index
func (t *TentativeAccept) IsSummaryString(index int) (ok bool) {
	return t.summary[index].stringName != nil

}

// GetSummaryString returns the value safely if IsSummaryString returned true for the specified index
func (t *TentativeAccept) GetSummaryString(index int) (v string) {
	return *t.summary[index].stringName

}

// AppendSummaryString adds to the back of summary a string type
func (t *TentativeAccept) AppendSummaryString(v string) {
	t.summary = append(t.summary, &summaryTentativeAcceptIntermediateType{stringName: &v})

}

// PrependSummaryString adds to the front of summary a string type
func (t *TentativeAccept) PrependSummaryString(v string) {
	t.summary = append([]*summaryTentativeAcceptIntermediateType{&summaryTentativeAcceptIntermediateType{stringName: &v}}, t.summary...)

}

// RemoveSummaryString deletes the value from the specified index
func (t *TentativeAccept) RemoveSummaryString(index int) {
	copy(t.summary[index:], t.summary[index+1:])
	t.summary[len(t.summary)-1] = nil
	t.summary = t.summary[:len(t.summary)-1]

}

// IsSummaryLangString determines whether the call to GetSummaryLangString is safe for the specified index
func (t *TentativeAccept) IsSummaryLangString(index int) (ok bool) {
	return t.summary[index].langString != nil

}

// GetSummaryLangString returns the value safely if IsSummaryLangString returned true for the specified index
func (t *TentativeAccept) GetSummaryLangString(index int) (v string) {
	return *t.summary[index].langString

}

// AppendSummaryLangString adds to the back of summary a string type
func (t *TentativeAccept) AppendSummaryLangString(v string) {
	t.summary = append(t.summary, &summaryTentativeAcceptIntermediateType{langString: &v})

}

// PrependSummaryLangString adds to the front of summary a string type
func (t *TentativeAccept) PrependSummaryLangString(v string) {
	t.summary = append([]*summaryTentativeAcceptIntermediateType{&summaryTentativeAcceptIntermediateType{langString: &v}}, t.summary...)

}

// RemoveSummaryLangString deletes the value from the specified index
func (t *TentativeAccept) RemoveSummaryLangString(index int) {
	copy(t.summary[index:], t.summary[index+1:])
	t.summary[len(t.summary)-1] = nil
	t.summary = t.summary[:len(t.summary)-1]

}

// IsSummaryIRI determines whether the call to GetSummaryIRI is safe for the specified index
func (t *TentativeAccept) IsSummaryIRI(index int) (ok bool) {
	return t.summary[index].IRI != nil

}

// GetSummaryIRI returns the value safely if IsSummaryIRI returned true for the specified index
func (t *TentativeAccept) GetSummaryIRI(index int) (v *url.URL) {
	return t.summary[index].IRI

}

// AppendSummaryIRI adds to the back of summary a *url.URL type
func (t *TentativeAccept) AppendSummaryIRI(v *url.URL) {
	t.summary = append(t.summary, &summaryTentativeAcceptIntermediateType{IRI: v})

}

// PrependSummaryIRI adds to the front of summary a *url.URL type
func (t *TentativeAccept) PrependSummaryIRI(v *url.URL) {
	t.summary = append([]*summaryTentativeAcceptIntermediateType{&summaryTentativeAcceptIntermediateType{IRI: v}}, t.summary...)

}

// RemoveSummaryIRI deletes the value from the specified index
func (t *TentativeAccept) RemoveSummaryIRI(index int) {
	copy(t.summary[index:], t.summary[index+1:])
	t.summary[len(t.summary)-1] = nil
	t.summary = t.summary[:len(t.summary)-1]

}

// HasUnknownSummary determines whether the call to GetUnknownSummary is safe
func (t *TentativeAccept) HasUnknownSummary() (ok bool) {
	return t.summary != nil && t.summary[0].unknown_ != nil

}

// GetUnknownSummary returns the unknown value for summary
func (t *TentativeAccept) GetUnknownSummary() (v interface{}) {
	return t.summary[0].unknown_

}

// SetUnknownSummary sets the unknown value of summary
func (t *TentativeAccept) SetUnknownSummary(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &summaryTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.summary = append(t.summary, tmp)

}

// SummaryMapLanguages returns all languages for this property's language mapping, or nil if there are none.
func (t *TentativeAccept) SummaryMapLanguages() (l []string) {
	if t.summaryMap == nil || len(t.summaryMap) == 0 {
		return nil
	}
	for k := range t.summaryMap {
		l = append(l, k)
	}
	return

}

// GetSummaryMap retrieves the value of the property for the specified language, or an empty string if it does not exist
func (t *TentativeAccept) GetSummaryMap(l string) (v string) {
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
func (t *TentativeAccept) SetSummaryMap(l string, v string) {
	if t.summaryMap == nil {
		t.summaryMap = make(map[string]string)
	}
	t.summaryMap[l] = v

}

// TagLen determines the number of elements able to be used for the IsTagObject, GetTagObject, and RemoveTagObject functions
func (t *TentativeAccept) TagLen() (l int) {
	return len(t.tag)

}

// IsTagObject determines whether the call to GetTagObject is safe for the specified index
func (t *TentativeAccept) IsTagObject(index int) (ok bool) {
	return t.tag[index].Object != nil

}

// GetTagObject returns the value safely if IsTagObject returned true for the specified index
func (t *TentativeAccept) GetTagObject(index int) (v ObjectType) {
	return t.tag[index].Object

}

// AppendTagObject adds to the back of tag a ObjectType type
func (t *TentativeAccept) AppendTagObject(v ObjectType) {
	t.tag = append(t.tag, &tagTentativeAcceptIntermediateType{Object: v})

}

// PrependTagObject adds to the front of tag a ObjectType type
func (t *TentativeAccept) PrependTagObject(v ObjectType) {
	t.tag = append([]*tagTentativeAcceptIntermediateType{&tagTentativeAcceptIntermediateType{Object: v}}, t.tag...)

}

// RemoveTagObject deletes the value from the specified index
func (t *TentativeAccept) RemoveTagObject(index int) {
	copy(t.tag[index:], t.tag[index+1:])
	t.tag[len(t.tag)-1] = nil
	t.tag = t.tag[:len(t.tag)-1]

}

// IsTagLink determines whether the call to GetTagLink is safe for the specified index
func (t *TentativeAccept) IsTagLink(index int) (ok bool) {
	return t.tag[index].Link != nil

}

// GetTagLink returns the value safely if IsTagLink returned true for the specified index
func (t *TentativeAccept) GetTagLink(index int) (v LinkType) {
	return t.tag[index].Link

}

// AppendTagLink adds to the back of tag a LinkType type
func (t *TentativeAccept) AppendTagLink(v LinkType) {
	t.tag = append(t.tag, &tagTentativeAcceptIntermediateType{Link: v})

}

// PrependTagLink adds to the front of tag a LinkType type
func (t *TentativeAccept) PrependTagLink(v LinkType) {
	t.tag = append([]*tagTentativeAcceptIntermediateType{&tagTentativeAcceptIntermediateType{Link: v}}, t.tag...)

}

// RemoveTagLink deletes the value from the specified index
func (t *TentativeAccept) RemoveTagLink(index int) {
	copy(t.tag[index:], t.tag[index+1:])
	t.tag[len(t.tag)-1] = nil
	t.tag = t.tag[:len(t.tag)-1]

}

// IsTagIRI determines whether the call to GetTagIRI is safe for the specified index
func (t *TentativeAccept) IsTagIRI(index int) (ok bool) {
	return t.tag[index].IRI != nil

}

// GetTagIRI returns the value safely if IsTagIRI returned true for the specified index
func (t *TentativeAccept) GetTagIRI(index int) (v *url.URL) {
	return t.tag[index].IRI

}

// AppendTagIRI adds to the back of tag a *url.URL type
func (t *TentativeAccept) AppendTagIRI(v *url.URL) {
	t.tag = append(t.tag, &tagTentativeAcceptIntermediateType{IRI: v})

}

// PrependTagIRI adds to the front of tag a *url.URL type
func (t *TentativeAccept) PrependTagIRI(v *url.URL) {
	t.tag = append([]*tagTentativeAcceptIntermediateType{&tagTentativeAcceptIntermediateType{IRI: v}}, t.tag...)

}

// RemoveTagIRI deletes the value from the specified index
func (t *TentativeAccept) RemoveTagIRI(index int) {
	copy(t.tag[index:], t.tag[index+1:])
	t.tag[len(t.tag)-1] = nil
	t.tag = t.tag[:len(t.tag)-1]

}

// HasUnknownTag determines whether the call to GetUnknownTag is safe
func (t *TentativeAccept) HasUnknownTag() (ok bool) {
	return t.tag != nil && t.tag[0].unknown_ != nil

}

// GetUnknownTag returns the unknown value for tag
func (t *TentativeAccept) GetUnknownTag() (v interface{}) {
	return t.tag[0].unknown_

}

// SetUnknownTag sets the unknown value of tag
func (t *TentativeAccept) SetUnknownTag(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &tagTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.tag = append(t.tag, tmp)

}

// TypeLen determines the number of elements able to be used for the GetType and RemoveType functions
func (t *TentativeAccept) TypeLen() (l int) {
	return len(t.typeName)

}

// GetType returns the value for the specified index
func (t *TentativeAccept) GetType(index int) (v interface{}) {
	return t.typeName[index]

}

// AppendType adds a value to the back of type
func (t *TentativeAccept) AppendType(v interface{}) {
	t.typeName = append(t.typeName, v)

}

// PrependType adds a value to the front of type
func (t *TentativeAccept) PrependType(v interface{}) {
	t.typeName = append([]interface{}{v}, t.typeName...)

}

// RemoveType deletes the value from the specified index
func (t *TentativeAccept) RemoveType(index int) {
	copy(t.typeName[index:], t.typeName[index+1:])
	t.typeName[len(t.typeName)-1] = nil
	t.typeName = t.typeName[:len(t.typeName)-1]

}

// IsUpdated determines whether the call to GetUpdated is safe
func (t *TentativeAccept) IsUpdated() (ok bool) {
	return t.updated != nil && t.updated.dateTime != nil

}

// GetUpdated returns the value safely if IsUpdated returned true
func (t *TentativeAccept) GetUpdated() (v time.Time) {
	return *t.updated.dateTime

}

// SetUpdated sets the value of updated to be of time.Time type
func (t *TentativeAccept) SetUpdated(v time.Time) {
	t.updated = &updatedTentativeAcceptIntermediateType{dateTime: &v}

}

// IsUpdatedIRI determines whether the call to GetUpdatedIRI is safe
func (t *TentativeAccept) IsUpdatedIRI() (ok bool) {
	return t.updated != nil && t.updated.IRI != nil

}

// GetUpdatedIRI returns the value safely if IsUpdatedIRI returned true
func (t *TentativeAccept) GetUpdatedIRI() (v *url.URL) {
	return t.updated.IRI

}

// SetUpdatedIRI sets the value of updated to be of *url.URL type
func (t *TentativeAccept) SetUpdatedIRI(v *url.URL) {
	t.updated = &updatedTentativeAcceptIntermediateType{IRI: v}

}

// HasUnknownUpdated determines whether the call to GetUnknownUpdated is safe
func (t *TentativeAccept) HasUnknownUpdated() (ok bool) {
	return t.updated != nil && t.updated.unknown_ != nil

}

// GetUnknownUpdated returns the unknown value for updated
func (t *TentativeAccept) GetUnknownUpdated() (v interface{}) {
	return t.updated.unknown_

}

// SetUnknownUpdated sets the unknown value of updated
func (t *TentativeAccept) SetUnknownUpdated(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &updatedTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.updated = tmp

}

// UrlLen determines the number of elements able to be used for the IsUrlAnyURI, GetUrlAnyURI, and RemoveUrlAnyURI functions
func (t *TentativeAccept) UrlLen() (l int) {
	return len(t.url)

}

// IsUrlAnyURI determines whether the call to GetUrlAnyURI is safe for the specified index
func (t *TentativeAccept) IsUrlAnyURI(index int) (ok bool) {
	return t.url[index].anyURI != nil

}

// GetUrlAnyURI returns the value safely if IsUrlAnyURI returned true for the specified index
func (t *TentativeAccept) GetUrlAnyURI(index int) (v *url.URL) {
	return t.url[index].anyURI

}

// AppendUrlAnyURI adds to the back of url a *url.URL type
func (t *TentativeAccept) AppendUrlAnyURI(v *url.URL) {
	t.url = append(t.url, &urlTentativeAcceptIntermediateType{anyURI: v})

}

// PrependUrlAnyURI adds to the front of url a *url.URL type
func (t *TentativeAccept) PrependUrlAnyURI(v *url.URL) {
	t.url = append([]*urlTentativeAcceptIntermediateType{&urlTentativeAcceptIntermediateType{anyURI: v}}, t.url...)

}

// RemoveUrlAnyURI deletes the value from the specified index
func (t *TentativeAccept) RemoveUrlAnyURI(index int) {
	copy(t.url[index:], t.url[index+1:])
	t.url[len(t.url)-1] = nil
	t.url = t.url[:len(t.url)-1]

}

// IsUrlLink determines whether the call to GetUrlLink is safe for the specified index
func (t *TentativeAccept) IsUrlLink(index int) (ok bool) {
	return t.url[index].Link != nil

}

// GetUrlLink returns the value safely if IsUrlLink returned true for the specified index
func (t *TentativeAccept) GetUrlLink(index int) (v LinkType) {
	return t.url[index].Link

}

// AppendUrlLink adds to the back of url a LinkType type
func (t *TentativeAccept) AppendUrlLink(v LinkType) {
	t.url = append(t.url, &urlTentativeAcceptIntermediateType{Link: v})

}

// PrependUrlLink adds to the front of url a LinkType type
func (t *TentativeAccept) PrependUrlLink(v LinkType) {
	t.url = append([]*urlTentativeAcceptIntermediateType{&urlTentativeAcceptIntermediateType{Link: v}}, t.url...)

}

// RemoveUrlLink deletes the value from the specified index
func (t *TentativeAccept) RemoveUrlLink(index int) {
	copy(t.url[index:], t.url[index+1:])
	t.url[len(t.url)-1] = nil
	t.url = t.url[:len(t.url)-1]

}

// HasUnknownUrl determines whether the call to GetUnknownUrl is safe
func (t *TentativeAccept) HasUnknownUrl() (ok bool) {
	return t.url != nil && t.url[0].unknown_ != nil

}

// GetUnknownUrl returns the unknown value for url
func (t *TentativeAccept) GetUnknownUrl() (v interface{}) {
	return t.url[0].unknown_

}

// SetUnknownUrl sets the unknown value of url
func (t *TentativeAccept) SetUnknownUrl(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &urlTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.url = append(t.url, tmp)

}

// ToLen determines the number of elements able to be used for the IsToObject, GetToObject, and RemoveToObject functions
func (t *TentativeAccept) ToLen() (l int) {
	return len(t.to)

}

// IsToObject determines whether the call to GetToObject is safe for the specified index
func (t *TentativeAccept) IsToObject(index int) (ok bool) {
	return t.to[index].Object != nil

}

// GetToObject returns the value safely if IsToObject returned true for the specified index
func (t *TentativeAccept) GetToObject(index int) (v ObjectType) {
	return t.to[index].Object

}

// AppendToObject adds to the back of to a ObjectType type
func (t *TentativeAccept) AppendToObject(v ObjectType) {
	t.to = append(t.to, &toTentativeAcceptIntermediateType{Object: v})

}

// PrependToObject adds to the front of to a ObjectType type
func (t *TentativeAccept) PrependToObject(v ObjectType) {
	t.to = append([]*toTentativeAcceptIntermediateType{&toTentativeAcceptIntermediateType{Object: v}}, t.to...)

}

// RemoveToObject deletes the value from the specified index
func (t *TentativeAccept) RemoveToObject(index int) {
	copy(t.to[index:], t.to[index+1:])
	t.to[len(t.to)-1] = nil
	t.to = t.to[:len(t.to)-1]

}

// IsToLink determines whether the call to GetToLink is safe for the specified index
func (t *TentativeAccept) IsToLink(index int) (ok bool) {
	return t.to[index].Link != nil

}

// GetToLink returns the value safely if IsToLink returned true for the specified index
func (t *TentativeAccept) GetToLink(index int) (v LinkType) {
	return t.to[index].Link

}

// AppendToLink adds to the back of to a LinkType type
func (t *TentativeAccept) AppendToLink(v LinkType) {
	t.to = append(t.to, &toTentativeAcceptIntermediateType{Link: v})

}

// PrependToLink adds to the front of to a LinkType type
func (t *TentativeAccept) PrependToLink(v LinkType) {
	t.to = append([]*toTentativeAcceptIntermediateType{&toTentativeAcceptIntermediateType{Link: v}}, t.to...)

}

// RemoveToLink deletes the value from the specified index
func (t *TentativeAccept) RemoveToLink(index int) {
	copy(t.to[index:], t.to[index+1:])
	t.to[len(t.to)-1] = nil
	t.to = t.to[:len(t.to)-1]

}

// IsToIRI determines whether the call to GetToIRI is safe for the specified index
func (t *TentativeAccept) IsToIRI(index int) (ok bool) {
	return t.to[index].IRI != nil

}

// GetToIRI returns the value safely if IsToIRI returned true for the specified index
func (t *TentativeAccept) GetToIRI(index int) (v *url.URL) {
	return t.to[index].IRI

}

// AppendToIRI adds to the back of to a *url.URL type
func (t *TentativeAccept) AppendToIRI(v *url.URL) {
	t.to = append(t.to, &toTentativeAcceptIntermediateType{IRI: v})

}

// PrependToIRI adds to the front of to a *url.URL type
func (t *TentativeAccept) PrependToIRI(v *url.URL) {
	t.to = append([]*toTentativeAcceptIntermediateType{&toTentativeAcceptIntermediateType{IRI: v}}, t.to...)

}

// RemoveToIRI deletes the value from the specified index
func (t *TentativeAccept) RemoveToIRI(index int) {
	copy(t.to[index:], t.to[index+1:])
	t.to[len(t.to)-1] = nil
	t.to = t.to[:len(t.to)-1]

}

// HasUnknownTo determines whether the call to GetUnknownTo is safe
func (t *TentativeAccept) HasUnknownTo() (ok bool) {
	return t.to != nil && t.to[0].unknown_ != nil

}

// GetUnknownTo returns the unknown value for to
func (t *TentativeAccept) GetUnknownTo() (v interface{}) {
	return t.to[0].unknown_

}

// SetUnknownTo sets the unknown value of to
func (t *TentativeAccept) SetUnknownTo(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &toTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.to = append(t.to, tmp)

}

// BtoLen determines the number of elements able to be used for the IsBtoObject, GetBtoObject, and RemoveBtoObject functions
func (t *TentativeAccept) BtoLen() (l int) {
	return len(t.bto)

}

// IsBtoObject determines whether the call to GetBtoObject is safe for the specified index
func (t *TentativeAccept) IsBtoObject(index int) (ok bool) {
	return t.bto[index].Object != nil

}

// GetBtoObject returns the value safely if IsBtoObject returned true for the specified index
func (t *TentativeAccept) GetBtoObject(index int) (v ObjectType) {
	return t.bto[index].Object

}

// AppendBtoObject adds to the back of bto a ObjectType type
func (t *TentativeAccept) AppendBtoObject(v ObjectType) {
	t.bto = append(t.bto, &btoTentativeAcceptIntermediateType{Object: v})

}

// PrependBtoObject adds to the front of bto a ObjectType type
func (t *TentativeAccept) PrependBtoObject(v ObjectType) {
	t.bto = append([]*btoTentativeAcceptIntermediateType{&btoTentativeAcceptIntermediateType{Object: v}}, t.bto...)

}

// RemoveBtoObject deletes the value from the specified index
func (t *TentativeAccept) RemoveBtoObject(index int) {
	copy(t.bto[index:], t.bto[index+1:])
	t.bto[len(t.bto)-1] = nil
	t.bto = t.bto[:len(t.bto)-1]

}

// IsBtoLink determines whether the call to GetBtoLink is safe for the specified index
func (t *TentativeAccept) IsBtoLink(index int) (ok bool) {
	return t.bto[index].Link != nil

}

// GetBtoLink returns the value safely if IsBtoLink returned true for the specified index
func (t *TentativeAccept) GetBtoLink(index int) (v LinkType) {
	return t.bto[index].Link

}

// AppendBtoLink adds to the back of bto a LinkType type
func (t *TentativeAccept) AppendBtoLink(v LinkType) {
	t.bto = append(t.bto, &btoTentativeAcceptIntermediateType{Link: v})

}

// PrependBtoLink adds to the front of bto a LinkType type
func (t *TentativeAccept) PrependBtoLink(v LinkType) {
	t.bto = append([]*btoTentativeAcceptIntermediateType{&btoTentativeAcceptIntermediateType{Link: v}}, t.bto...)

}

// RemoveBtoLink deletes the value from the specified index
func (t *TentativeAccept) RemoveBtoLink(index int) {
	copy(t.bto[index:], t.bto[index+1:])
	t.bto[len(t.bto)-1] = nil
	t.bto = t.bto[:len(t.bto)-1]

}

// IsBtoIRI determines whether the call to GetBtoIRI is safe for the specified index
func (t *TentativeAccept) IsBtoIRI(index int) (ok bool) {
	return t.bto[index].IRI != nil

}

// GetBtoIRI returns the value safely if IsBtoIRI returned true for the specified index
func (t *TentativeAccept) GetBtoIRI(index int) (v *url.URL) {
	return t.bto[index].IRI

}

// AppendBtoIRI adds to the back of bto a *url.URL type
func (t *TentativeAccept) AppendBtoIRI(v *url.URL) {
	t.bto = append(t.bto, &btoTentativeAcceptIntermediateType{IRI: v})

}

// PrependBtoIRI adds to the front of bto a *url.URL type
func (t *TentativeAccept) PrependBtoIRI(v *url.URL) {
	t.bto = append([]*btoTentativeAcceptIntermediateType{&btoTentativeAcceptIntermediateType{IRI: v}}, t.bto...)

}

// RemoveBtoIRI deletes the value from the specified index
func (t *TentativeAccept) RemoveBtoIRI(index int) {
	copy(t.bto[index:], t.bto[index+1:])
	t.bto[len(t.bto)-1] = nil
	t.bto = t.bto[:len(t.bto)-1]

}

// HasUnknownBto determines whether the call to GetUnknownBto is safe
func (t *TentativeAccept) HasUnknownBto() (ok bool) {
	return t.bto != nil && t.bto[0].unknown_ != nil

}

// GetUnknownBto returns the unknown value for bto
func (t *TentativeAccept) GetUnknownBto() (v interface{}) {
	return t.bto[0].unknown_

}

// SetUnknownBto sets the unknown value of bto
func (t *TentativeAccept) SetUnknownBto(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &btoTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.bto = append(t.bto, tmp)

}

// CcLen determines the number of elements able to be used for the IsCcObject, GetCcObject, and RemoveCcObject functions
func (t *TentativeAccept) CcLen() (l int) {
	return len(t.cc)

}

// IsCcObject determines whether the call to GetCcObject is safe for the specified index
func (t *TentativeAccept) IsCcObject(index int) (ok bool) {
	return t.cc[index].Object != nil

}

// GetCcObject returns the value safely if IsCcObject returned true for the specified index
func (t *TentativeAccept) GetCcObject(index int) (v ObjectType) {
	return t.cc[index].Object

}

// AppendCcObject adds to the back of cc a ObjectType type
func (t *TentativeAccept) AppendCcObject(v ObjectType) {
	t.cc = append(t.cc, &ccTentativeAcceptIntermediateType{Object: v})

}

// PrependCcObject adds to the front of cc a ObjectType type
func (t *TentativeAccept) PrependCcObject(v ObjectType) {
	t.cc = append([]*ccTentativeAcceptIntermediateType{&ccTentativeAcceptIntermediateType{Object: v}}, t.cc...)

}

// RemoveCcObject deletes the value from the specified index
func (t *TentativeAccept) RemoveCcObject(index int) {
	copy(t.cc[index:], t.cc[index+1:])
	t.cc[len(t.cc)-1] = nil
	t.cc = t.cc[:len(t.cc)-1]

}

// IsCcLink determines whether the call to GetCcLink is safe for the specified index
func (t *TentativeAccept) IsCcLink(index int) (ok bool) {
	return t.cc[index].Link != nil

}

// GetCcLink returns the value safely if IsCcLink returned true for the specified index
func (t *TentativeAccept) GetCcLink(index int) (v LinkType) {
	return t.cc[index].Link

}

// AppendCcLink adds to the back of cc a LinkType type
func (t *TentativeAccept) AppendCcLink(v LinkType) {
	t.cc = append(t.cc, &ccTentativeAcceptIntermediateType{Link: v})

}

// PrependCcLink adds to the front of cc a LinkType type
func (t *TentativeAccept) PrependCcLink(v LinkType) {
	t.cc = append([]*ccTentativeAcceptIntermediateType{&ccTentativeAcceptIntermediateType{Link: v}}, t.cc...)

}

// RemoveCcLink deletes the value from the specified index
func (t *TentativeAccept) RemoveCcLink(index int) {
	copy(t.cc[index:], t.cc[index+1:])
	t.cc[len(t.cc)-1] = nil
	t.cc = t.cc[:len(t.cc)-1]

}

// IsCcIRI determines whether the call to GetCcIRI is safe for the specified index
func (t *TentativeAccept) IsCcIRI(index int) (ok bool) {
	return t.cc[index].IRI != nil

}

// GetCcIRI returns the value safely if IsCcIRI returned true for the specified index
func (t *TentativeAccept) GetCcIRI(index int) (v *url.URL) {
	return t.cc[index].IRI

}

// AppendCcIRI adds to the back of cc a *url.URL type
func (t *TentativeAccept) AppendCcIRI(v *url.URL) {
	t.cc = append(t.cc, &ccTentativeAcceptIntermediateType{IRI: v})

}

// PrependCcIRI adds to the front of cc a *url.URL type
func (t *TentativeAccept) PrependCcIRI(v *url.URL) {
	t.cc = append([]*ccTentativeAcceptIntermediateType{&ccTentativeAcceptIntermediateType{IRI: v}}, t.cc...)

}

// RemoveCcIRI deletes the value from the specified index
func (t *TentativeAccept) RemoveCcIRI(index int) {
	copy(t.cc[index:], t.cc[index+1:])
	t.cc[len(t.cc)-1] = nil
	t.cc = t.cc[:len(t.cc)-1]

}

// HasUnknownCc determines whether the call to GetUnknownCc is safe
func (t *TentativeAccept) HasUnknownCc() (ok bool) {
	return t.cc != nil && t.cc[0].unknown_ != nil

}

// GetUnknownCc returns the unknown value for cc
func (t *TentativeAccept) GetUnknownCc() (v interface{}) {
	return t.cc[0].unknown_

}

// SetUnknownCc sets the unknown value of cc
func (t *TentativeAccept) SetUnknownCc(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &ccTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.cc = append(t.cc, tmp)

}

// BccLen determines the number of elements able to be used for the IsBccObject, GetBccObject, and RemoveBccObject functions
func (t *TentativeAccept) BccLen() (l int) {
	return len(t.bcc)

}

// IsBccObject determines whether the call to GetBccObject is safe for the specified index
func (t *TentativeAccept) IsBccObject(index int) (ok bool) {
	return t.bcc[index].Object != nil

}

// GetBccObject returns the value safely if IsBccObject returned true for the specified index
func (t *TentativeAccept) GetBccObject(index int) (v ObjectType) {
	return t.bcc[index].Object

}

// AppendBccObject adds to the back of bcc a ObjectType type
func (t *TentativeAccept) AppendBccObject(v ObjectType) {
	t.bcc = append(t.bcc, &bccTentativeAcceptIntermediateType{Object: v})

}

// PrependBccObject adds to the front of bcc a ObjectType type
func (t *TentativeAccept) PrependBccObject(v ObjectType) {
	t.bcc = append([]*bccTentativeAcceptIntermediateType{&bccTentativeAcceptIntermediateType{Object: v}}, t.bcc...)

}

// RemoveBccObject deletes the value from the specified index
func (t *TentativeAccept) RemoveBccObject(index int) {
	copy(t.bcc[index:], t.bcc[index+1:])
	t.bcc[len(t.bcc)-1] = nil
	t.bcc = t.bcc[:len(t.bcc)-1]

}

// IsBccLink determines whether the call to GetBccLink is safe for the specified index
func (t *TentativeAccept) IsBccLink(index int) (ok bool) {
	return t.bcc[index].Link != nil

}

// GetBccLink returns the value safely if IsBccLink returned true for the specified index
func (t *TentativeAccept) GetBccLink(index int) (v LinkType) {
	return t.bcc[index].Link

}

// AppendBccLink adds to the back of bcc a LinkType type
func (t *TentativeAccept) AppendBccLink(v LinkType) {
	t.bcc = append(t.bcc, &bccTentativeAcceptIntermediateType{Link: v})

}

// PrependBccLink adds to the front of bcc a LinkType type
func (t *TentativeAccept) PrependBccLink(v LinkType) {
	t.bcc = append([]*bccTentativeAcceptIntermediateType{&bccTentativeAcceptIntermediateType{Link: v}}, t.bcc...)

}

// RemoveBccLink deletes the value from the specified index
func (t *TentativeAccept) RemoveBccLink(index int) {
	copy(t.bcc[index:], t.bcc[index+1:])
	t.bcc[len(t.bcc)-1] = nil
	t.bcc = t.bcc[:len(t.bcc)-1]

}

// IsBccIRI determines whether the call to GetBccIRI is safe for the specified index
func (t *TentativeAccept) IsBccIRI(index int) (ok bool) {
	return t.bcc[index].IRI != nil

}

// GetBccIRI returns the value safely if IsBccIRI returned true for the specified index
func (t *TentativeAccept) GetBccIRI(index int) (v *url.URL) {
	return t.bcc[index].IRI

}

// AppendBccIRI adds to the back of bcc a *url.URL type
func (t *TentativeAccept) AppendBccIRI(v *url.URL) {
	t.bcc = append(t.bcc, &bccTentativeAcceptIntermediateType{IRI: v})

}

// PrependBccIRI adds to the front of bcc a *url.URL type
func (t *TentativeAccept) PrependBccIRI(v *url.URL) {
	t.bcc = append([]*bccTentativeAcceptIntermediateType{&bccTentativeAcceptIntermediateType{IRI: v}}, t.bcc...)

}

// RemoveBccIRI deletes the value from the specified index
func (t *TentativeAccept) RemoveBccIRI(index int) {
	copy(t.bcc[index:], t.bcc[index+1:])
	t.bcc[len(t.bcc)-1] = nil
	t.bcc = t.bcc[:len(t.bcc)-1]

}

// HasUnknownBcc determines whether the call to GetUnknownBcc is safe
func (t *TentativeAccept) HasUnknownBcc() (ok bool) {
	return t.bcc != nil && t.bcc[0].unknown_ != nil

}

// GetUnknownBcc returns the unknown value for bcc
func (t *TentativeAccept) GetUnknownBcc() (v interface{}) {
	return t.bcc[0].unknown_

}

// SetUnknownBcc sets the unknown value of bcc
func (t *TentativeAccept) SetUnknownBcc(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &bccTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.bcc = append(t.bcc, tmp)

}

// IsMediaType determines whether the call to GetMediaType is safe
func (t *TentativeAccept) IsMediaType() (ok bool) {
	return t.mediaType != nil && t.mediaType.mimeMediaTypeValue != nil

}

// GetMediaType returns the value safely if IsMediaType returned true
func (t *TentativeAccept) GetMediaType() (v string) {
	return *t.mediaType.mimeMediaTypeValue

}

// SetMediaType sets the value of mediaType to be of string type
func (t *TentativeAccept) SetMediaType(v string) {
	t.mediaType = &mediaTypeTentativeAcceptIntermediateType{mimeMediaTypeValue: &v}

}

// IsMediaTypeIRI determines whether the call to GetMediaTypeIRI is safe
func (t *TentativeAccept) IsMediaTypeIRI() (ok bool) {
	return t.mediaType != nil && t.mediaType.IRI != nil

}

// GetMediaTypeIRI returns the value safely if IsMediaTypeIRI returned true
func (t *TentativeAccept) GetMediaTypeIRI() (v *url.URL) {
	return t.mediaType.IRI

}

// SetMediaTypeIRI sets the value of mediaType to be of *url.URL type
func (t *TentativeAccept) SetMediaTypeIRI(v *url.URL) {
	t.mediaType = &mediaTypeTentativeAcceptIntermediateType{IRI: v}

}

// HasUnknownMediaType determines whether the call to GetUnknownMediaType is safe
func (t *TentativeAccept) HasUnknownMediaType() (ok bool) {
	return t.mediaType != nil && t.mediaType.unknown_ != nil

}

// GetUnknownMediaType returns the unknown value for mediaType
func (t *TentativeAccept) GetUnknownMediaType() (v interface{}) {
	return t.mediaType.unknown_

}

// SetUnknownMediaType sets the unknown value of mediaType
func (t *TentativeAccept) SetUnknownMediaType(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &mediaTypeTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.mediaType = tmp

}

// IsDuration determines whether the call to GetDuration is safe
func (t *TentativeAccept) IsDuration() (ok bool) {
	return t.duration != nil && t.duration.duration != nil

}

// GetDuration returns the value safely if IsDuration returned true
func (t *TentativeAccept) GetDuration() (v time.Duration) {
	return *t.duration.duration

}

// SetDuration sets the value of duration to be of time.Duration type
func (t *TentativeAccept) SetDuration(v time.Duration) {
	t.duration = &durationTentativeAcceptIntermediateType{duration: &v}

}

// IsDurationIRI determines whether the call to GetDurationIRI is safe
func (t *TentativeAccept) IsDurationIRI() (ok bool) {
	return t.duration != nil && t.duration.IRI != nil

}

// GetDurationIRI returns the value safely if IsDurationIRI returned true
func (t *TentativeAccept) GetDurationIRI() (v *url.URL) {
	return t.duration.IRI

}

// SetDurationIRI sets the value of duration to be of *url.URL type
func (t *TentativeAccept) SetDurationIRI(v *url.URL) {
	t.duration = &durationTentativeAcceptIntermediateType{IRI: v}

}

// HasUnknownDuration determines whether the call to GetUnknownDuration is safe
func (t *TentativeAccept) HasUnknownDuration() (ok bool) {
	return t.duration != nil && t.duration.unknown_ != nil

}

// GetUnknownDuration returns the unknown value for duration
func (t *TentativeAccept) GetUnknownDuration() (v interface{}) {
	return t.duration.unknown_

}

// SetUnknownDuration sets the unknown value of duration
func (t *TentativeAccept) SetUnknownDuration(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &durationTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.duration = tmp

}

// IsSource determines whether the call to GetSource is safe
func (t *TentativeAccept) IsSource() (ok bool) {
	return t.source != nil && t.source.Object != nil

}

// GetSource returns the value safely if IsSource returned true
func (t *TentativeAccept) GetSource() (v ObjectType) {
	return t.source.Object

}

// SetSource sets the value of source to be of ObjectType type
func (t *TentativeAccept) SetSource(v ObjectType) {
	t.source = &sourceTentativeAcceptIntermediateType{Object: v}

}

// IsSourceIRI determines whether the call to GetSourceIRI is safe
func (t *TentativeAccept) IsSourceIRI() (ok bool) {
	return t.source != nil && t.source.IRI != nil

}

// GetSourceIRI returns the value safely if IsSourceIRI returned true
func (t *TentativeAccept) GetSourceIRI() (v *url.URL) {
	return t.source.IRI

}

// SetSourceIRI sets the value of source to be of *url.URL type
func (t *TentativeAccept) SetSourceIRI(v *url.URL) {
	t.source = &sourceTentativeAcceptIntermediateType{IRI: v}

}

// HasUnknownSource determines whether the call to GetUnknownSource is safe
func (t *TentativeAccept) HasUnknownSource() (ok bool) {
	return t.source != nil && t.source.unknown_ != nil

}

// GetUnknownSource returns the unknown value for source
func (t *TentativeAccept) GetUnknownSource() (v interface{}) {
	return t.source.unknown_

}

// SetUnknownSource sets the unknown value of source
func (t *TentativeAccept) SetUnknownSource(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &sourceTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.source = tmp

}

// IsInboxOrderedCollection determines whether the call to GetInboxOrderedCollection is safe
func (t *TentativeAccept) IsInboxOrderedCollection() (ok bool) {
	return t.inbox != nil && t.inbox.OrderedCollection != nil

}

// GetInboxOrderedCollection returns the value safely if IsInboxOrderedCollection returned true
func (t *TentativeAccept) GetInboxOrderedCollection() (v OrderedCollectionType) {
	return t.inbox.OrderedCollection

}

// SetInboxOrderedCollection sets the value of inbox to be of OrderedCollectionType type
func (t *TentativeAccept) SetInboxOrderedCollection(v OrderedCollectionType) {
	t.inbox = &inboxTentativeAcceptIntermediateType{OrderedCollection: v}

}

// IsInboxAnyURI determines whether the call to GetInboxAnyURI is safe
func (t *TentativeAccept) IsInboxAnyURI() (ok bool) {
	return t.inbox != nil && t.inbox.anyURI != nil

}

// GetInboxAnyURI returns the value safely if IsInboxAnyURI returned true
func (t *TentativeAccept) GetInboxAnyURI() (v *url.URL) {
	return t.inbox.anyURI

}

// SetInboxAnyURI sets the value of inbox to be of *url.URL type
func (t *TentativeAccept) SetInboxAnyURI(v *url.URL) {
	t.inbox = &inboxTentativeAcceptIntermediateType{anyURI: v}

}

// HasUnknownInbox determines whether the call to GetUnknownInbox is safe
func (t *TentativeAccept) HasUnknownInbox() (ok bool) {
	return t.inbox != nil && t.inbox.unknown_ != nil

}

// GetUnknownInbox returns the unknown value for inbox
func (t *TentativeAccept) GetUnknownInbox() (v interface{}) {
	return t.inbox.unknown_

}

// SetUnknownInbox sets the unknown value of inbox
func (t *TentativeAccept) SetUnknownInbox(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &inboxTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.inbox = tmp

}

// IsOutboxOrderedCollection determines whether the call to GetOutboxOrderedCollection is safe
func (t *TentativeAccept) IsOutboxOrderedCollection() (ok bool) {
	return t.outbox != nil && t.outbox.OrderedCollection != nil

}

// GetOutboxOrderedCollection returns the value safely if IsOutboxOrderedCollection returned true
func (t *TentativeAccept) GetOutboxOrderedCollection() (v OrderedCollectionType) {
	return t.outbox.OrderedCollection

}

// SetOutboxOrderedCollection sets the value of outbox to be of OrderedCollectionType type
func (t *TentativeAccept) SetOutboxOrderedCollection(v OrderedCollectionType) {
	t.outbox = &outboxTentativeAcceptIntermediateType{OrderedCollection: v}

}

// IsOutboxAnyURI determines whether the call to GetOutboxAnyURI is safe
func (t *TentativeAccept) IsOutboxAnyURI() (ok bool) {
	return t.outbox != nil && t.outbox.anyURI != nil

}

// GetOutboxAnyURI returns the value safely if IsOutboxAnyURI returned true
func (t *TentativeAccept) GetOutboxAnyURI() (v *url.URL) {
	return t.outbox.anyURI

}

// SetOutboxAnyURI sets the value of outbox to be of *url.URL type
func (t *TentativeAccept) SetOutboxAnyURI(v *url.URL) {
	t.outbox = &outboxTentativeAcceptIntermediateType{anyURI: v}

}

// HasUnknownOutbox determines whether the call to GetUnknownOutbox is safe
func (t *TentativeAccept) HasUnknownOutbox() (ok bool) {
	return t.outbox != nil && t.outbox.unknown_ != nil

}

// GetUnknownOutbox returns the unknown value for outbox
func (t *TentativeAccept) GetUnknownOutbox() (v interface{}) {
	return t.outbox.unknown_

}

// SetUnknownOutbox sets the unknown value of outbox
func (t *TentativeAccept) SetUnknownOutbox(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &outboxTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.outbox = tmp

}

// IsFollowingCollection determines whether the call to GetFollowingCollection is safe
func (t *TentativeAccept) IsFollowingCollection() (ok bool) {
	return t.following != nil && t.following.Collection != nil

}

// GetFollowingCollection returns the value safely if IsFollowingCollection returned true
func (t *TentativeAccept) GetFollowingCollection() (v CollectionType) {
	return t.following.Collection

}

// SetFollowingCollection sets the value of following to be of CollectionType type
func (t *TentativeAccept) SetFollowingCollection(v CollectionType) {
	t.following = &followingTentativeAcceptIntermediateType{Collection: v}

}

// IsFollowingOrderedCollection determines whether the call to GetFollowingOrderedCollection is safe
func (t *TentativeAccept) IsFollowingOrderedCollection() (ok bool) {
	return t.following != nil && t.following.OrderedCollection != nil

}

// GetFollowingOrderedCollection returns the value safely if IsFollowingOrderedCollection returned true
func (t *TentativeAccept) GetFollowingOrderedCollection() (v OrderedCollectionType) {
	return t.following.OrderedCollection

}

// SetFollowingOrderedCollection sets the value of following to be of OrderedCollectionType type
func (t *TentativeAccept) SetFollowingOrderedCollection(v OrderedCollectionType) {
	t.following = &followingTentativeAcceptIntermediateType{OrderedCollection: v}

}

// IsFollowingAnyURI determines whether the call to GetFollowingAnyURI is safe
func (t *TentativeAccept) IsFollowingAnyURI() (ok bool) {
	return t.following != nil && t.following.anyURI != nil

}

// GetFollowingAnyURI returns the value safely if IsFollowingAnyURI returned true
func (t *TentativeAccept) GetFollowingAnyURI() (v *url.URL) {
	return t.following.anyURI

}

// SetFollowingAnyURI sets the value of following to be of *url.URL type
func (t *TentativeAccept) SetFollowingAnyURI(v *url.URL) {
	t.following = &followingTentativeAcceptIntermediateType{anyURI: v}

}

// HasUnknownFollowing determines whether the call to GetUnknownFollowing is safe
func (t *TentativeAccept) HasUnknownFollowing() (ok bool) {
	return t.following != nil && t.following.unknown_ != nil

}

// GetUnknownFollowing returns the unknown value for following
func (t *TentativeAccept) GetUnknownFollowing() (v interface{}) {
	return t.following.unknown_

}

// SetUnknownFollowing sets the unknown value of following
func (t *TentativeAccept) SetUnknownFollowing(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &followingTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.following = tmp

}

// IsFollowersCollection determines whether the call to GetFollowersCollection is safe
func (t *TentativeAccept) IsFollowersCollection() (ok bool) {
	return t.followers != nil && t.followers.Collection != nil

}

// GetFollowersCollection returns the value safely if IsFollowersCollection returned true
func (t *TentativeAccept) GetFollowersCollection() (v CollectionType) {
	return t.followers.Collection

}

// SetFollowersCollection sets the value of followers to be of CollectionType type
func (t *TentativeAccept) SetFollowersCollection(v CollectionType) {
	t.followers = &followersTentativeAcceptIntermediateType{Collection: v}

}

// IsFollowersOrderedCollection determines whether the call to GetFollowersOrderedCollection is safe
func (t *TentativeAccept) IsFollowersOrderedCollection() (ok bool) {
	return t.followers != nil && t.followers.OrderedCollection != nil

}

// GetFollowersOrderedCollection returns the value safely if IsFollowersOrderedCollection returned true
func (t *TentativeAccept) GetFollowersOrderedCollection() (v OrderedCollectionType) {
	return t.followers.OrderedCollection

}

// SetFollowersOrderedCollection sets the value of followers to be of OrderedCollectionType type
func (t *TentativeAccept) SetFollowersOrderedCollection(v OrderedCollectionType) {
	t.followers = &followersTentativeAcceptIntermediateType{OrderedCollection: v}

}

// IsFollowersAnyURI determines whether the call to GetFollowersAnyURI is safe
func (t *TentativeAccept) IsFollowersAnyURI() (ok bool) {
	return t.followers != nil && t.followers.anyURI != nil

}

// GetFollowersAnyURI returns the value safely if IsFollowersAnyURI returned true
func (t *TentativeAccept) GetFollowersAnyURI() (v *url.URL) {
	return t.followers.anyURI

}

// SetFollowersAnyURI sets the value of followers to be of *url.URL type
func (t *TentativeAccept) SetFollowersAnyURI(v *url.URL) {
	t.followers = &followersTentativeAcceptIntermediateType{anyURI: v}

}

// HasUnknownFollowers determines whether the call to GetUnknownFollowers is safe
func (t *TentativeAccept) HasUnknownFollowers() (ok bool) {
	return t.followers != nil && t.followers.unknown_ != nil

}

// GetUnknownFollowers returns the unknown value for followers
func (t *TentativeAccept) GetUnknownFollowers() (v interface{}) {
	return t.followers.unknown_

}

// SetUnknownFollowers sets the unknown value of followers
func (t *TentativeAccept) SetUnknownFollowers(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &followersTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.followers = tmp

}

// IsLikedCollection determines whether the call to GetLikedCollection is safe
func (t *TentativeAccept) IsLikedCollection() (ok bool) {
	return t.liked != nil && t.liked.Collection != nil

}

// GetLikedCollection returns the value safely if IsLikedCollection returned true
func (t *TentativeAccept) GetLikedCollection() (v CollectionType) {
	return t.liked.Collection

}

// SetLikedCollection sets the value of liked to be of CollectionType type
func (t *TentativeAccept) SetLikedCollection(v CollectionType) {
	t.liked = &likedTentativeAcceptIntermediateType{Collection: v}

}

// IsLikedOrderedCollection determines whether the call to GetLikedOrderedCollection is safe
func (t *TentativeAccept) IsLikedOrderedCollection() (ok bool) {
	return t.liked != nil && t.liked.OrderedCollection != nil

}

// GetLikedOrderedCollection returns the value safely if IsLikedOrderedCollection returned true
func (t *TentativeAccept) GetLikedOrderedCollection() (v OrderedCollectionType) {
	return t.liked.OrderedCollection

}

// SetLikedOrderedCollection sets the value of liked to be of OrderedCollectionType type
func (t *TentativeAccept) SetLikedOrderedCollection(v OrderedCollectionType) {
	t.liked = &likedTentativeAcceptIntermediateType{OrderedCollection: v}

}

// IsLikedAnyURI determines whether the call to GetLikedAnyURI is safe
func (t *TentativeAccept) IsLikedAnyURI() (ok bool) {
	return t.liked != nil && t.liked.anyURI != nil

}

// GetLikedAnyURI returns the value safely if IsLikedAnyURI returned true
func (t *TentativeAccept) GetLikedAnyURI() (v *url.URL) {
	return t.liked.anyURI

}

// SetLikedAnyURI sets the value of liked to be of *url.URL type
func (t *TentativeAccept) SetLikedAnyURI(v *url.URL) {
	t.liked = &likedTentativeAcceptIntermediateType{anyURI: v}

}

// HasUnknownLiked determines whether the call to GetUnknownLiked is safe
func (t *TentativeAccept) HasUnknownLiked() (ok bool) {
	return t.liked != nil && t.liked.unknown_ != nil

}

// GetUnknownLiked returns the unknown value for liked
func (t *TentativeAccept) GetUnknownLiked() (v interface{}) {
	return t.liked.unknown_

}

// SetUnknownLiked sets the unknown value of liked
func (t *TentativeAccept) SetUnknownLiked(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &likedTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.liked = tmp

}

// IsLikesCollection determines whether the call to GetLikesCollection is safe
func (t *TentativeAccept) IsLikesCollection() (ok bool) {
	return t.likes != nil && t.likes.Collection != nil

}

// GetLikesCollection returns the value safely if IsLikesCollection returned true
func (t *TentativeAccept) GetLikesCollection() (v CollectionType) {
	return t.likes.Collection

}

// SetLikesCollection sets the value of likes to be of CollectionType type
func (t *TentativeAccept) SetLikesCollection(v CollectionType) {
	t.likes = &likesTentativeAcceptIntermediateType{Collection: v}

}

// IsLikesOrderedCollection determines whether the call to GetLikesOrderedCollection is safe
func (t *TentativeAccept) IsLikesOrderedCollection() (ok bool) {
	return t.likes != nil && t.likes.OrderedCollection != nil

}

// GetLikesOrderedCollection returns the value safely if IsLikesOrderedCollection returned true
func (t *TentativeAccept) GetLikesOrderedCollection() (v OrderedCollectionType) {
	return t.likes.OrderedCollection

}

// SetLikesOrderedCollection sets the value of likes to be of OrderedCollectionType type
func (t *TentativeAccept) SetLikesOrderedCollection(v OrderedCollectionType) {
	t.likes = &likesTentativeAcceptIntermediateType{OrderedCollection: v}

}

// IsLikesAnyURI determines whether the call to GetLikesAnyURI is safe
func (t *TentativeAccept) IsLikesAnyURI() (ok bool) {
	return t.likes != nil && t.likes.anyURI != nil

}

// GetLikesAnyURI returns the value safely if IsLikesAnyURI returned true
func (t *TentativeAccept) GetLikesAnyURI() (v *url.URL) {
	return t.likes.anyURI

}

// SetLikesAnyURI sets the value of likes to be of *url.URL type
func (t *TentativeAccept) SetLikesAnyURI(v *url.URL) {
	t.likes = &likesTentativeAcceptIntermediateType{anyURI: v}

}

// HasUnknownLikes determines whether the call to GetUnknownLikes is safe
func (t *TentativeAccept) HasUnknownLikes() (ok bool) {
	return t.likes != nil && t.likes.unknown_ != nil

}

// GetUnknownLikes returns the unknown value for likes
func (t *TentativeAccept) GetUnknownLikes() (v interface{}) {
	return t.likes.unknown_

}

// SetUnknownLikes sets the unknown value of likes
func (t *TentativeAccept) SetUnknownLikes(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &likesTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.likes = tmp

}

// StreamsLen determines the number of elements able to be used for the GetStreams and RemoveStreams functions
func (t *TentativeAccept) StreamsLen() (l int) {
	return len(t.streams)

}

// GetStreams returns the value for the specified index
func (t *TentativeAccept) GetStreams(index int) (v *url.URL) {
	return t.streams[index]

}

// AppendStreams adds a value to the back of streams
func (t *TentativeAccept) AppendStreams(v *url.URL) {
	t.streams = append(t.streams, v)

}

// PrependStreams adds a value to the front of streams
func (t *TentativeAccept) PrependStreams(v *url.URL) {
	t.streams = append([]*url.URL{v}, t.streams...)

}

// RemoveStreams deletes the value from the specified index
func (t *TentativeAccept) RemoveStreams(index int) {
	copy(t.streams[index:], t.streams[index+1:])
	t.streams[len(t.streams)-1] = nil
	t.streams = t.streams[:len(t.streams)-1]

}

// HasUnknownStreams determines whether the call to GetUnknownStreams is safe
func (t *TentativeAccept) HasUnknownStreams() (ok bool) {
	return t.unknown_ != nil && t.unknown_["streams"] != nil

}

// GetUnknownStreams returns the unknown value for streams
func (t *TentativeAccept) GetUnknownStreams() (v interface{}) {
	return t.unknown_["streams"]

}

// SetUnknownStreams sets the unknown value of streams
func (t *TentativeAccept) SetUnknownStreams(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["streams"] = i

}

// IsPreferredUsername determines whether the call to GetPreferredUsername is safe
func (t *TentativeAccept) IsPreferredUsername() (ok bool) {
	return t.preferredUsername != nil && t.preferredUsername.stringName != nil

}

// GetPreferredUsername returns the value safely if IsPreferredUsername returned true
func (t *TentativeAccept) GetPreferredUsername() (v string) {
	return *t.preferredUsername.stringName

}

// SetPreferredUsername sets the value of preferredUsername to be of string type
func (t *TentativeAccept) SetPreferredUsername(v string) {
	t.preferredUsername = &preferredUsernameTentativeAcceptIntermediateType{stringName: &v}

}

// IsPreferredUsernameIRI determines whether the call to GetPreferredUsernameIRI is safe
func (t *TentativeAccept) IsPreferredUsernameIRI() (ok bool) {
	return t.preferredUsername != nil && t.preferredUsername.IRI != nil

}

// GetPreferredUsernameIRI returns the value safely if IsPreferredUsernameIRI returned true
func (t *TentativeAccept) GetPreferredUsernameIRI() (v *url.URL) {
	return t.preferredUsername.IRI

}

// SetPreferredUsernameIRI sets the value of preferredUsername to be of *url.URL type
func (t *TentativeAccept) SetPreferredUsernameIRI(v *url.URL) {
	t.preferredUsername = &preferredUsernameTentativeAcceptIntermediateType{IRI: v}

}

// HasUnknownPreferredUsername determines whether the call to GetUnknownPreferredUsername is safe
func (t *TentativeAccept) HasUnknownPreferredUsername() (ok bool) {
	return t.preferredUsername != nil && t.preferredUsername.unknown_ != nil

}

// GetUnknownPreferredUsername returns the unknown value for preferredUsername
func (t *TentativeAccept) GetUnknownPreferredUsername() (v interface{}) {
	return t.preferredUsername.unknown_

}

// SetUnknownPreferredUsername sets the unknown value of preferredUsername
func (t *TentativeAccept) SetUnknownPreferredUsername(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &preferredUsernameTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.preferredUsername = tmp

}

// PreferredUsernameMapLanguages returns all languages for this property's language mapping, or nil if there are none.
func (t *TentativeAccept) PreferredUsernameMapLanguages() (l []string) {
	if t.preferredUsernameMap == nil || len(t.preferredUsernameMap) == 0 {
		return nil
	}
	for k := range t.preferredUsernameMap {
		l = append(l, k)
	}
	return

}

// GetPreferredUsernameMap retrieves the value of the property for the specified language, or an empty string if it does not exist
func (t *TentativeAccept) GetPreferredUsernameMap(l string) (v string) {
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
func (t *TentativeAccept) SetPreferredUsernameMap(l string, v string) {
	if t.preferredUsernameMap == nil {
		t.preferredUsernameMap = make(map[string]string)
	}
	t.preferredUsernameMap[l] = v

}

// IsEndpoints determines whether the call to GetEndpoints is safe
func (t *TentativeAccept) IsEndpoints() (ok bool) {
	return t.endpoints != nil && t.endpoints.Object != nil

}

// GetEndpoints returns the value safely if IsEndpoints returned true
func (t *TentativeAccept) GetEndpoints() (v ObjectType) {
	return t.endpoints.Object

}

// SetEndpoints sets the value of endpoints to be of ObjectType type
func (t *TentativeAccept) SetEndpoints(v ObjectType) {
	t.endpoints = &endpointsTentativeAcceptIntermediateType{Object: v}

}

// IsEndpointsIRI determines whether the call to GetEndpointsIRI is safe
func (t *TentativeAccept) IsEndpointsIRI() (ok bool) {
	return t.endpoints != nil && t.endpoints.IRI != nil

}

// GetEndpointsIRI returns the value safely if IsEndpointsIRI returned true
func (t *TentativeAccept) GetEndpointsIRI() (v *url.URL) {
	return t.endpoints.IRI

}

// SetEndpointsIRI sets the value of endpoints to be of *url.URL type
func (t *TentativeAccept) SetEndpointsIRI(v *url.URL) {
	t.endpoints = &endpointsTentativeAcceptIntermediateType{IRI: v}

}

// HasUnknownEndpoints determines whether the call to GetUnknownEndpoints is safe
func (t *TentativeAccept) HasUnknownEndpoints() (ok bool) {
	return t.endpoints != nil && t.endpoints.unknown_ != nil

}

// GetUnknownEndpoints returns the unknown value for endpoints
func (t *TentativeAccept) GetUnknownEndpoints() (v interface{}) {
	return t.endpoints.unknown_

}

// SetUnknownEndpoints sets the unknown value of endpoints
func (t *TentativeAccept) SetUnknownEndpoints(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &endpointsTentativeAcceptIntermediateType{}
	tmp.unknown_ = i
	t.endpoints = tmp

}

// HasProxyUrl determines whether the call to GetProxyUrl is safe
func (t *TentativeAccept) HasProxyUrl() (ok bool) {
	return t.proxyUrl != nil

}

// GetProxyUrl returns the value for proxyUrl
func (t *TentativeAccept) GetProxyUrl() (v *url.URL) {
	return t.proxyUrl

}

// SetProxyUrl sets the value of proxyUrl
func (t *TentativeAccept) SetProxyUrl(v *url.URL) {
	t.proxyUrl = v

}

// HasUnknownProxyUrl determines whether the call to GetUnknownProxyUrl is safe
func (t *TentativeAccept) HasUnknownProxyUrl() (ok bool) {
	return t.unknown_ != nil && t.unknown_["proxyUrl"] != nil

}

// GetUnknownProxyUrl returns the unknown value for proxyUrl
func (t *TentativeAccept) GetUnknownProxyUrl() (v interface{}) {
	return t.unknown_["proxyUrl"]

}

// SetUnknownProxyUrl sets the unknown value of proxyUrl
func (t *TentativeAccept) SetUnknownProxyUrl(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["proxyUrl"] = i

}

// HasOauthAuthorizationEndpoint determines whether the call to GetOauthAuthorizationEndpoint is safe
func (t *TentativeAccept) HasOauthAuthorizationEndpoint() (ok bool) {
	return t.oauthAuthorizationEndpoint != nil

}

// GetOauthAuthorizationEndpoint returns the value for oauthAuthorizationEndpoint
func (t *TentativeAccept) GetOauthAuthorizationEndpoint() (v *url.URL) {
	return t.oauthAuthorizationEndpoint

}

// SetOauthAuthorizationEndpoint sets the value of oauthAuthorizationEndpoint
func (t *TentativeAccept) SetOauthAuthorizationEndpoint(v *url.URL) {
	t.oauthAuthorizationEndpoint = v

}

// HasUnknownOauthAuthorizationEndpoint determines whether the call to GetUnknownOauthAuthorizationEndpoint is safe
func (t *TentativeAccept) HasUnknownOauthAuthorizationEndpoint() (ok bool) {
	return t.unknown_ != nil && t.unknown_["oauthAuthorizationEndpoint"] != nil

}

// GetUnknownOauthAuthorizationEndpoint returns the unknown value for oauthAuthorizationEndpoint
func (t *TentativeAccept) GetUnknownOauthAuthorizationEndpoint() (v interface{}) {
	return t.unknown_["oauthAuthorizationEndpoint"]

}

// SetUnknownOauthAuthorizationEndpoint sets the unknown value of oauthAuthorizationEndpoint
func (t *TentativeAccept) SetUnknownOauthAuthorizationEndpoint(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["oauthAuthorizationEndpoint"] = i

}

// HasOauthTokenEndpoint determines whether the call to GetOauthTokenEndpoint is safe
func (t *TentativeAccept) HasOauthTokenEndpoint() (ok bool) {
	return t.oauthTokenEndpoint != nil

}

// GetOauthTokenEndpoint returns the value for oauthTokenEndpoint
func (t *TentativeAccept) GetOauthTokenEndpoint() (v *url.URL) {
	return t.oauthTokenEndpoint

}

// SetOauthTokenEndpoint sets the value of oauthTokenEndpoint
func (t *TentativeAccept) SetOauthTokenEndpoint(v *url.URL) {
	t.oauthTokenEndpoint = v

}

// HasUnknownOauthTokenEndpoint determines whether the call to GetUnknownOauthTokenEndpoint is safe
func (t *TentativeAccept) HasUnknownOauthTokenEndpoint() (ok bool) {
	return t.unknown_ != nil && t.unknown_["oauthTokenEndpoint"] != nil

}

// GetUnknownOauthTokenEndpoint returns the unknown value for oauthTokenEndpoint
func (t *TentativeAccept) GetUnknownOauthTokenEndpoint() (v interface{}) {
	return t.unknown_["oauthTokenEndpoint"]

}

// SetUnknownOauthTokenEndpoint sets the unknown value of oauthTokenEndpoint
func (t *TentativeAccept) SetUnknownOauthTokenEndpoint(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["oauthTokenEndpoint"] = i

}

// HasProvideClientKey determines whether the call to GetProvideClientKey is safe
func (t *TentativeAccept) HasProvideClientKey() (ok bool) {
	return t.provideClientKey != nil

}

// GetProvideClientKey returns the value for provideClientKey
func (t *TentativeAccept) GetProvideClientKey() (v *url.URL) {
	return t.provideClientKey

}

// SetProvideClientKey sets the value of provideClientKey
func (t *TentativeAccept) SetProvideClientKey(v *url.URL) {
	t.provideClientKey = v

}

// HasUnknownProvideClientKey determines whether the call to GetUnknownProvideClientKey is safe
func (t *TentativeAccept) HasUnknownProvideClientKey() (ok bool) {
	return t.unknown_ != nil && t.unknown_["provideClientKey"] != nil

}

// GetUnknownProvideClientKey returns the unknown value for provideClientKey
func (t *TentativeAccept) GetUnknownProvideClientKey() (v interface{}) {
	return t.unknown_["provideClientKey"]

}

// SetUnknownProvideClientKey sets the unknown value of provideClientKey
func (t *TentativeAccept) SetUnknownProvideClientKey(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["provideClientKey"] = i

}

// HasSignClientKey determines whether the call to GetSignClientKey is safe
func (t *TentativeAccept) HasSignClientKey() (ok bool) {
	return t.signClientKey != nil

}

// GetSignClientKey returns the value for signClientKey
func (t *TentativeAccept) GetSignClientKey() (v *url.URL) {
	return t.signClientKey

}

// SetSignClientKey sets the value of signClientKey
func (t *TentativeAccept) SetSignClientKey(v *url.URL) {
	t.signClientKey = v

}

// HasUnknownSignClientKey determines whether the call to GetUnknownSignClientKey is safe
func (t *TentativeAccept) HasUnknownSignClientKey() (ok bool) {
	return t.unknown_ != nil && t.unknown_["signClientKey"] != nil

}

// GetUnknownSignClientKey returns the unknown value for signClientKey
func (t *TentativeAccept) GetUnknownSignClientKey() (v interface{}) {
	return t.unknown_["signClientKey"]

}

// SetUnknownSignClientKey sets the unknown value of signClientKey
func (t *TentativeAccept) SetUnknownSignClientKey(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["signClientKey"] = i

}

// HasSharedInbox determines whether the call to GetSharedInbox is safe
func (t *TentativeAccept) HasSharedInbox() (ok bool) {
	return t.sharedInbox != nil

}

// GetSharedInbox returns the value for sharedInbox
func (t *TentativeAccept) GetSharedInbox() (v *url.URL) {
	return t.sharedInbox

}

// SetSharedInbox sets the value of sharedInbox
func (t *TentativeAccept) SetSharedInbox(v *url.URL) {
	t.sharedInbox = v

}

// HasUnknownSharedInbox determines whether the call to GetUnknownSharedInbox is safe
func (t *TentativeAccept) HasUnknownSharedInbox() (ok bool) {
	return t.unknown_ != nil && t.unknown_["sharedInbox"] != nil

}

// GetUnknownSharedInbox returns the unknown value for sharedInbox
func (t *TentativeAccept) GetUnknownSharedInbox() (v interface{}) {
	return t.unknown_["sharedInbox"]

}

// SetUnknownSharedInbox sets the unknown value of sharedInbox
func (t *TentativeAccept) SetUnknownSharedInbox(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["sharedInbox"] = i

}

// AddUnknown adds a raw extension to this object with the specified key
func (t *TentativeAccept) AddUnknown(k string, i interface{}) (this *TentativeAccept) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_[k] = i
	return t

}

// HasUnknown returns true if there is an unknown object with the specified key
func (t *TentativeAccept) HasUnknown(k string) (b bool) {
	if t.unknown_ == nil {
		return false
	}
	_, ok := t.unknown_[k]
	return ok

}

// RemoveUnknown removes a raw extension from this object with the specified key
func (t *TentativeAccept) RemoveUnknown(k string) (this *TentativeAccept) {
	delete(t.unknown_, k)
	return t

}

// Serialize turns this object into a map[string]interface{}. Note that the "type" property will automatically be populated with "TentativeAccept" if not manually set by the caller
func (t *TentativeAccept) Serialize() (m map[string]interface{}, err error) {
	m = make(map[string]interface{})
	for k, v := range t.unknown_ {
		m[k] = unknownValueSerialize(v)
	}
	var typeAlreadySet bool
	for _, k := range t.typeName {
		if ks, ok := k.(string); ok {
			if ks == "TentativeAccept" {
				typeAlreadySet = true
				break
			}
		}
	}
	if !typeAlreadySet {
		t.typeName = append(t.typeName, "TentativeAccept")
	}
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceActorTentativeAcceptIntermediateType(t.actor); err == nil && v != nil {
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
	if v, err := serializeSliceObjectTentativeAcceptIntermediateType(t.object); err == nil && v != nil {
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
	if v, err := serializeSliceTargetTentativeAcceptIntermediateType(t.target); err == nil && v != nil {
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
	if v, err := serializeSliceResultTentativeAcceptIntermediateType(t.result); err == nil && v != nil {
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
	if v, err := serializeSliceOriginTentativeAcceptIntermediateType(t.origin); err == nil && v != nil {
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
	if v, err := serializeSliceInstrumentTentativeAcceptIntermediateType(t.instrument); err == nil && v != nil {
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
		if v, err := serializeAltitudeTentativeAcceptIntermediateType(t.altitude); err == nil {
			m["altitude"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceAttachmentTentativeAcceptIntermediateType(t.attachment); err == nil && v != nil {
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
	if v, err := serializeSliceAttributedToTentativeAcceptIntermediateType(t.attributedTo); err == nil && v != nil {
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
	if v, err := serializeSliceAudienceTentativeAcceptIntermediateType(t.audience); err == nil && v != nil {
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
	if v, err := serializeSliceContentTentativeAcceptIntermediateType(t.content); err == nil && v != nil {
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
	if v, err := serializeSliceContextTentativeAcceptIntermediateType(t.context); err == nil && v != nil {
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
	if v, err := serializeSliceNameTentativeAcceptIntermediateType(t.name); err == nil && v != nil {
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
		if v, err := serializeEndTimeTentativeAcceptIntermediateType(t.endTime); err == nil {
			m["endTime"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceGeneratorTentativeAcceptIntermediateType(t.generator); err == nil && v != nil {
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
	if v, err := serializeSliceIconTentativeAcceptIntermediateType(t.icon); err == nil && v != nil {
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
	if v, err := serializeSliceImageTentativeAcceptIntermediateType(t.image); err == nil && v != nil {
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
	if v, err := serializeSliceInReplyToTentativeAcceptIntermediateType(t.inReplyTo); err == nil && v != nil {
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
	if v, err := serializeSliceLocationTentativeAcceptIntermediateType(t.location); err == nil && v != nil {
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
	if v, err := serializeSlicePreviewTentativeAcceptIntermediateType(t.preview); err == nil && v != nil {
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
		if v, err := serializePublishedTentativeAcceptIntermediateType(t.published); err == nil {
			m["published"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.replies != nil {
		if v, err := serializeRepliesTentativeAcceptIntermediateType(t.replies); err == nil {
			m["replies"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.startTime != nil {
		if v, err := serializeStartTimeTentativeAcceptIntermediateType(t.startTime); err == nil {
			m["startTime"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceSummaryTentativeAcceptIntermediateType(t.summary); err == nil && v != nil {
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
	if v, err := serializeSliceTagTentativeAcceptIntermediateType(t.tag); err == nil && v != nil {
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
		if v, err := serializeUpdatedTentativeAcceptIntermediateType(t.updated); err == nil {
			m["updated"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceUrlTentativeAcceptIntermediateType(t.url); err == nil && v != nil {
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
	if v, err := serializeSliceToTentativeAcceptIntermediateType(t.to); err == nil && v != nil {
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
	if v, err := serializeSliceBtoTentativeAcceptIntermediateType(t.bto); err == nil && v != nil {
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
	if v, err := serializeSliceCcTentativeAcceptIntermediateType(t.cc); err == nil && v != nil {
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
	if v, err := serializeSliceBccTentativeAcceptIntermediateType(t.bcc); err == nil && v != nil {
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
		if v, err := serializeMediaTypeTentativeAcceptIntermediateType(t.mediaType); err == nil {
			m["mediaType"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.duration != nil {
		if v, err := serializeDurationTentativeAcceptIntermediateType(t.duration); err == nil {
			m["duration"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.source != nil {
		if v, err := serializeSourceTentativeAcceptIntermediateType(t.source); err == nil {
			m["source"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.inbox != nil {
		if v, err := serializeInboxTentativeAcceptIntermediateType(t.inbox); err == nil {
			m["inbox"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.outbox != nil {
		if v, err := serializeOutboxTentativeAcceptIntermediateType(t.outbox); err == nil {
			m["outbox"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.following != nil {
		if v, err := serializeFollowingTentativeAcceptIntermediateType(t.following); err == nil {
			m["following"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.followers != nil {
		if v, err := serializeFollowersTentativeAcceptIntermediateType(t.followers); err == nil {
			m["followers"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.liked != nil {
		if v, err := serializeLikedTentativeAcceptIntermediateType(t.liked); err == nil {
			m["liked"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.likes != nil {
		if v, err := serializeLikesTentativeAcceptIntermediateType(t.likes); err == nil {
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
		if v, err := serializePreferredUsernameTentativeAcceptIntermediateType(t.preferredUsername); err == nil {
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
		if v, err := serializeEndpointsTentativeAcceptIntermediateType(t.endpoints); err == nil {
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
	return

}

// Deserialize populates this object from a map[string]interface{}
func (t *TentativeAccept) Deserialize(m map[string]interface{}) (err error) {
	for k, v := range m {
		handled := false
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "actor" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeActorTentativeAcceptIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.actor = []*actorTentativeAcceptIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.actor, err = deserializeSliceActorTentativeAcceptIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &actorTentativeAcceptIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.actor = []*actorTentativeAcceptIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "object" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeObjectTentativeAcceptIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.object = []*objectTentativeAcceptIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.object, err = deserializeSliceObjectTentativeAcceptIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &objectTentativeAcceptIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.object = []*objectTentativeAcceptIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "target" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeTargetTentativeAcceptIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.target = []*targetTentativeAcceptIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.target, err = deserializeSliceTargetTentativeAcceptIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &targetTentativeAcceptIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.target = []*targetTentativeAcceptIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "result" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeResultTentativeAcceptIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.result = []*resultTentativeAcceptIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.result, err = deserializeSliceResultTentativeAcceptIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &resultTentativeAcceptIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.result = []*resultTentativeAcceptIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "origin" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeOriginTentativeAcceptIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.origin = []*originTentativeAcceptIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.origin, err = deserializeSliceOriginTentativeAcceptIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &originTentativeAcceptIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.origin = []*originTentativeAcceptIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "instrument" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeInstrumentTentativeAcceptIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.instrument = []*instrumentTentativeAcceptIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.instrument, err = deserializeSliceInstrumentTentativeAcceptIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &instrumentTentativeAcceptIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.instrument = []*instrumentTentativeAcceptIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "altitude" {
				t.altitude, err = deserializeAltitudeTentativeAcceptIntermediateType(v)
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
					tmp, err := deserializeAttachmentTentativeAcceptIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.attachment = []*attachmentTentativeAcceptIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.attachment, err = deserializeSliceAttachmentTentativeAcceptIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &attachmentTentativeAcceptIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.attachment = []*attachmentTentativeAcceptIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "attributedTo" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeAttributedToTentativeAcceptIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.attributedTo = []*attributedToTentativeAcceptIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.attributedTo, err = deserializeSliceAttributedToTentativeAcceptIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &attributedToTentativeAcceptIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.attributedTo = []*attributedToTentativeAcceptIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "audience" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeAudienceTentativeAcceptIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.audience = []*audienceTentativeAcceptIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.audience, err = deserializeSliceAudienceTentativeAcceptIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &audienceTentativeAcceptIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.audience = []*audienceTentativeAcceptIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "content" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeContentTentativeAcceptIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.content = []*contentTentativeAcceptIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.content, err = deserializeSliceContentTentativeAcceptIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &contentTentativeAcceptIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.content = []*contentTentativeAcceptIntermediateType{tmp}
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
					tmp, err := deserializeContextTentativeAcceptIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.context = []*contextTentativeAcceptIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.context, err = deserializeSliceContextTentativeAcceptIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &contextTentativeAcceptIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.context = []*contextTentativeAcceptIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "name" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeNameTentativeAcceptIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.name = []*nameTentativeAcceptIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.name, err = deserializeSliceNameTentativeAcceptIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &nameTentativeAcceptIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.name = []*nameTentativeAcceptIntermediateType{tmp}
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
				t.endTime, err = deserializeEndTimeTentativeAcceptIntermediateType(v)
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
					tmp, err := deserializeGeneratorTentativeAcceptIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.generator = []*generatorTentativeAcceptIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.generator, err = deserializeSliceGeneratorTentativeAcceptIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &generatorTentativeAcceptIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.generator = []*generatorTentativeAcceptIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "icon" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeIconTentativeAcceptIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.icon = []*iconTentativeAcceptIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.icon, err = deserializeSliceIconTentativeAcceptIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &iconTentativeAcceptIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.icon = []*iconTentativeAcceptIntermediateType{tmp}
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
					tmp, err := deserializeImageTentativeAcceptIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.image = []*imageTentativeAcceptIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.image, err = deserializeSliceImageTentativeAcceptIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &imageTentativeAcceptIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.image = []*imageTentativeAcceptIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "inReplyTo" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeInReplyToTentativeAcceptIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.inReplyTo = []*inReplyToTentativeAcceptIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.inReplyTo, err = deserializeSliceInReplyToTentativeAcceptIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &inReplyToTentativeAcceptIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.inReplyTo = []*inReplyToTentativeAcceptIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "location" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeLocationTentativeAcceptIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.location = []*locationTentativeAcceptIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.location, err = deserializeSliceLocationTentativeAcceptIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &locationTentativeAcceptIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.location = []*locationTentativeAcceptIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "preview" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializePreviewTentativeAcceptIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.preview = []*previewTentativeAcceptIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.preview, err = deserializeSlicePreviewTentativeAcceptIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &previewTentativeAcceptIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.preview = []*previewTentativeAcceptIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "published" {
				t.published, err = deserializePublishedTentativeAcceptIntermediateType(v)
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
				t.replies, err = deserializeRepliesTentativeAcceptIntermediateType(v)
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
				t.startTime, err = deserializeStartTimeTentativeAcceptIntermediateType(v)
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
					tmp, err := deserializeSummaryTentativeAcceptIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.summary = []*summaryTentativeAcceptIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.summary, err = deserializeSliceSummaryTentativeAcceptIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &summaryTentativeAcceptIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.summary = []*summaryTentativeAcceptIntermediateType{tmp}
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
					tmp, err := deserializeTagTentativeAcceptIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.tag = []*tagTentativeAcceptIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.tag, err = deserializeSliceTagTentativeAcceptIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &tagTentativeAcceptIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.tag = []*tagTentativeAcceptIntermediateType{tmp}
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
				t.updated, err = deserializeUpdatedTentativeAcceptIntermediateType(v)
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
					tmp, err := deserializeUrlTentativeAcceptIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.url = []*urlTentativeAcceptIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.url, err = deserializeSliceUrlTentativeAcceptIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &urlTentativeAcceptIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.url = []*urlTentativeAcceptIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "to" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeToTentativeAcceptIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.to = []*toTentativeAcceptIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.to, err = deserializeSliceToTentativeAcceptIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &toTentativeAcceptIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.to = []*toTentativeAcceptIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "bto" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeBtoTentativeAcceptIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.bto = []*btoTentativeAcceptIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.bto, err = deserializeSliceBtoTentativeAcceptIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &btoTentativeAcceptIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.bto = []*btoTentativeAcceptIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "cc" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeCcTentativeAcceptIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.cc = []*ccTentativeAcceptIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.cc, err = deserializeSliceCcTentativeAcceptIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &ccTentativeAcceptIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.cc = []*ccTentativeAcceptIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "bcc" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeBccTentativeAcceptIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.bcc = []*bccTentativeAcceptIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.bcc, err = deserializeSliceBccTentativeAcceptIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &bccTentativeAcceptIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.bcc = []*bccTentativeAcceptIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "mediaType" {
				t.mediaType, err = deserializeMediaTypeTentativeAcceptIntermediateType(v)
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
				t.duration, err = deserializeDurationTentativeAcceptIntermediateType(v)
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
				t.source, err = deserializeSourceTentativeAcceptIntermediateType(v)
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
				t.inbox, err = deserializeInboxTentativeAcceptIntermediateType(v)
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
				t.outbox, err = deserializeOutboxTentativeAcceptIntermediateType(v)
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
				t.following, err = deserializeFollowingTentativeAcceptIntermediateType(v)
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
				t.followers, err = deserializeFollowersTentativeAcceptIntermediateType(v)
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
				t.liked, err = deserializeLikedTentativeAcceptIntermediateType(v)
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
				t.likes, err = deserializeLikesTentativeAcceptIntermediateType(v)
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
				t.preferredUsername, err = deserializePreferredUsernameTentativeAcceptIntermediateType(v)
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
				t.endpoints, err = deserializeEndpointsTentativeAcceptIntermediateType(v)
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
		if !handled && k != "@context" {
			if t.unknown_ == nil {
				t.unknown_ = make(map[string]interface{})
			}
			t.unknown_[k] = unknownValueDeserialize(v)
		}
	}
	return

}

// actorTentativeAcceptIntermediateType will only have one of its values set at most
type actorTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for actor property
	Object ObjectType
	// Stores possible LinkType type for actor property
	Link LinkType
	// Stores possible *url.URL type for actor property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *actorTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *actorTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// objectTentativeAcceptIntermediateType will only have one of its values set at most
type objectTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for object property
	Object ObjectType
	// Stores possible *url.URL type for object property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *objectTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *objectTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// targetTentativeAcceptIntermediateType will only have one of its values set at most
type targetTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for target property
	Object ObjectType
	// Stores possible LinkType type for target property
	Link LinkType
	// Stores possible *url.URL type for target property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *targetTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *targetTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// resultTentativeAcceptIntermediateType will only have one of its values set at most
type resultTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for result property
	Object ObjectType
	// Stores possible LinkType type for result property
	Link LinkType
	// Stores possible *url.URL type for result property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *resultTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *resultTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// originTentativeAcceptIntermediateType will only have one of its values set at most
type originTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for origin property
	Object ObjectType
	// Stores possible LinkType type for origin property
	Link LinkType
	// Stores possible *url.URL type for origin property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *originTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *originTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// instrumentTentativeAcceptIntermediateType will only have one of its values set at most
type instrumentTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for instrument property
	Object ObjectType
	// Stores possible LinkType type for instrument property
	Link LinkType
	// Stores possible *url.URL type for instrument property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *instrumentTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *instrumentTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// altitudeTentativeAcceptIntermediateType will only have one of its values set at most
type altitudeTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *float64 type for altitude property
	float *float64
	// Stores possible *url.URL type for altitude property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *altitudeTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.float, err = floatDeserialize(i)
			if err != nil {
				t.float = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *altitudeTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.float != nil {
		i = floatSerialize(*t.float)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// attachmentTentativeAcceptIntermediateType will only have one of its values set at most
type attachmentTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for attachment property
	Object ObjectType
	// Stores possible LinkType type for attachment property
	Link LinkType
	// Stores possible *url.URL type for attachment property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *attachmentTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *attachmentTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// attributedToTentativeAcceptIntermediateType will only have one of its values set at most
type attributedToTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for attributedTo property
	Object ObjectType
	// Stores possible LinkType type for attributedTo property
	Link LinkType
	// Stores possible *url.URL type for attributedTo property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *attributedToTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *attributedToTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// audienceTentativeAcceptIntermediateType will only have one of its values set at most
type audienceTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for audience property
	Object ObjectType
	// Stores possible LinkType type for audience property
	Link LinkType
	// Stores possible *url.URL type for audience property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *audienceTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *audienceTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// contentTentativeAcceptIntermediateType will only have one of its values set at most
type contentTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *string type for content property
	stringName *string
	// Stores possible *string type for content property
	langString *string
	// Stores possible *url.URL type for content property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *contentTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.stringName, err = stringDeserialize(i)
			if err != nil {
				t.stringName = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.langString, err = langStringDeserialize(i)
			if err != nil {
				t.langString = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *contentTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.stringName != nil {
		i = stringSerialize(*t.stringName)
		return
	}
	if t.langString != nil {
		i = langStringSerialize(*t.langString)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// contextTentativeAcceptIntermediateType will only have one of its values set at most
type contextTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for context property
	Object ObjectType
	// Stores possible LinkType type for context property
	Link LinkType
	// Stores possible *url.URL type for context property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *contextTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *contextTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// nameTentativeAcceptIntermediateType will only have one of its values set at most
type nameTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *string type for name property
	stringName *string
	// Stores possible *string type for name property
	langString *string
	// Stores possible *url.URL type for name property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *nameTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.stringName, err = stringDeserialize(i)
			if err != nil {
				t.stringName = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.langString, err = langStringDeserialize(i)
			if err != nil {
				t.langString = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *nameTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.stringName != nil {
		i = stringSerialize(*t.stringName)
		return
	}
	if t.langString != nil {
		i = langStringSerialize(*t.langString)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// endTimeTentativeAcceptIntermediateType will only have one of its values set at most
type endTimeTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *time.Time type for endTime property
	dateTime *time.Time
	// Stores possible *url.URL type for endTime property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *endTimeTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.dateTime, err = dateTimeDeserialize(i)
			if err != nil {
				t.dateTime = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *endTimeTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.dateTime != nil {
		i = dateTimeSerialize(*t.dateTime)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// generatorTentativeAcceptIntermediateType will only have one of its values set at most
type generatorTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for generator property
	Object ObjectType
	// Stores possible LinkType type for generator property
	Link LinkType
	// Stores possible *url.URL type for generator property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *generatorTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *generatorTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// iconTentativeAcceptIntermediateType will only have one of its values set at most
type iconTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ImageType type for icon property
	Image ImageType
	// Stores possible LinkType type for icon property
	Link LinkType
	// Stores possible *url.URL type for icon property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *iconTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Image, ok = resolveObject(kind).(ImageType); t.Image != nil && ok {
						err = t.Image.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *iconTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.Image != nil {
		i, err = t.Image.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// imageTentativeAcceptIntermediateType will only have one of its values set at most
type imageTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ImageType type for image property
	Image ImageType
	// Stores possible LinkType type for image property
	Link LinkType
	// Stores possible *url.URL type for image property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *imageTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Image, ok = resolveObject(kind).(ImageType); t.Image != nil && ok {
						err = t.Image.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *imageTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.Image != nil {
		i, err = t.Image.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// inReplyToTentativeAcceptIntermediateType will only have one of its values set at most
type inReplyToTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for inReplyTo property
	Object ObjectType
	// Stores possible LinkType type for inReplyTo property
	Link LinkType
	// Stores possible *url.URL type for inReplyTo property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *inReplyToTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *inReplyToTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// locationTentativeAcceptIntermediateType will only have one of its values set at most
type locationTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for location property
	Object ObjectType
	// Stores possible LinkType type for location property
	Link LinkType
	// Stores possible *url.URL type for location property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *locationTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *locationTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// previewTentativeAcceptIntermediateType will only have one of its values set at most
type previewTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for preview property
	Object ObjectType
	// Stores possible LinkType type for preview property
	Link LinkType
	// Stores possible *url.URL type for preview property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *previewTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *previewTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// publishedTentativeAcceptIntermediateType will only have one of its values set at most
type publishedTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *time.Time type for published property
	dateTime *time.Time
	// Stores possible *url.URL type for published property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *publishedTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.dateTime, err = dateTimeDeserialize(i)
			if err != nil {
				t.dateTime = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *publishedTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.dateTime != nil {
		i = dateTimeSerialize(*t.dateTime)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// repliesTentativeAcceptIntermediateType will only have one of its values set at most
type repliesTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible CollectionType type for replies property
	Collection CollectionType
	// Stores possible *url.URL type for replies property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *repliesTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Collection, ok = resolveObject(kind).(CollectionType); t.Collection != nil && ok {
						err = t.Collection.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *repliesTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.Collection != nil {
		i, err = t.Collection.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// startTimeTentativeAcceptIntermediateType will only have one of its values set at most
type startTimeTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *time.Time type for startTime property
	dateTime *time.Time
	// Stores possible *url.URL type for startTime property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *startTimeTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.dateTime, err = dateTimeDeserialize(i)
			if err != nil {
				t.dateTime = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *startTimeTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.dateTime != nil {
		i = dateTimeSerialize(*t.dateTime)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// summaryTentativeAcceptIntermediateType will only have one of its values set at most
type summaryTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *string type for summary property
	stringName *string
	// Stores possible *string type for summary property
	langString *string
	// Stores possible *url.URL type for summary property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *summaryTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.stringName, err = stringDeserialize(i)
			if err != nil {
				t.stringName = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.langString, err = langStringDeserialize(i)
			if err != nil {
				t.langString = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *summaryTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.stringName != nil {
		i = stringSerialize(*t.stringName)
		return
	}
	if t.langString != nil {
		i = langStringSerialize(*t.langString)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// tagTentativeAcceptIntermediateType will only have one of its values set at most
type tagTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for tag property
	Object ObjectType
	// Stores possible LinkType type for tag property
	Link LinkType
	// Stores possible *url.URL type for tag property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *tagTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *tagTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// updatedTentativeAcceptIntermediateType will only have one of its values set at most
type updatedTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *time.Time type for updated property
	dateTime *time.Time
	// Stores possible *url.URL type for updated property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *updatedTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.dateTime, err = dateTimeDeserialize(i)
			if err != nil {
				t.dateTime = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *updatedTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.dateTime != nil {
		i = dateTimeSerialize(*t.dateTime)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// urlTentativeAcceptIntermediateType will only have one of its values set at most
type urlTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *url.URL type for url property
	anyURI *url.URL
	// Stores possible LinkType type for url property
	Link LinkType
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *urlTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.anyURI, err = anyURIDeserialize(i)
			if err != nil {
				t.anyURI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *urlTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.anyURI != nil {
		i = anyURISerialize(t.anyURI)
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// toTentativeAcceptIntermediateType will only have one of its values set at most
type toTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for to property
	Object ObjectType
	// Stores possible LinkType type for to property
	Link LinkType
	// Stores possible *url.URL type for to property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *toTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *toTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// btoTentativeAcceptIntermediateType will only have one of its values set at most
type btoTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for bto property
	Object ObjectType
	// Stores possible LinkType type for bto property
	Link LinkType
	// Stores possible *url.URL type for bto property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *btoTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *btoTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// ccTentativeAcceptIntermediateType will only have one of its values set at most
type ccTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for cc property
	Object ObjectType
	// Stores possible LinkType type for cc property
	Link LinkType
	// Stores possible *url.URL type for cc property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *ccTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *ccTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// bccTentativeAcceptIntermediateType will only have one of its values set at most
type bccTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for bcc property
	Object ObjectType
	// Stores possible LinkType type for bcc property
	Link LinkType
	// Stores possible *url.URL type for bcc property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *bccTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.Link, ok = resolveLink(kind).(LinkType); t.Link != nil && ok {
						err = t.Link.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *bccTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// mediaTypeTentativeAcceptIntermediateType will only have one of its values set at most
type mediaTypeTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *string type for mediaType property
	mimeMediaTypeValue *string
	// Stores possible *url.URL type for mediaType property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *mediaTypeTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.mimeMediaTypeValue, err = mimeMediaTypeValueDeserialize(i)
			if err != nil {
				t.mimeMediaTypeValue = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *mediaTypeTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.mimeMediaTypeValue != nil {
		i = mimeMediaTypeValueSerialize(*t.mimeMediaTypeValue)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// durationTentativeAcceptIntermediateType will only have one of its values set at most
type durationTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *time.Duration type for duration property
	duration *time.Duration
	// Stores possible *url.URL type for duration property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *durationTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.duration, err = durationDeserialize(i)
			if err != nil {
				t.duration = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *durationTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.duration != nil {
		i = durationSerialize(*t.duration)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// sourceTentativeAcceptIntermediateType will only have one of its values set at most
type sourceTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for source property
	Object ObjectType
	// Stores possible *url.URL type for source property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *sourceTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *sourceTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// inboxTentativeAcceptIntermediateType will only have one of its values set at most
type inboxTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible OrderedCollectionType type for inbox property
	OrderedCollection OrderedCollectionType
	// Stores possible *url.URL type for inbox property
	anyURI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *inboxTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.OrderedCollection, ok = resolveObject(kind).(OrderedCollectionType); t.OrderedCollection != nil && ok {
						err = t.OrderedCollection.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.anyURI, err = anyURIDeserialize(i)
			if err != nil {
				t.anyURI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *inboxTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.OrderedCollection != nil {
		i, err = t.OrderedCollection.Serialize()
		return
	}
	if t.anyURI != nil {
		i = anyURISerialize(t.anyURI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// outboxTentativeAcceptIntermediateType will only have one of its values set at most
type outboxTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible OrderedCollectionType type for outbox property
	OrderedCollection OrderedCollectionType
	// Stores possible *url.URL type for outbox property
	anyURI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *outboxTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.OrderedCollection, ok = resolveObject(kind).(OrderedCollectionType); t.OrderedCollection != nil && ok {
						err = t.OrderedCollection.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.anyURI, err = anyURIDeserialize(i)
			if err != nil {
				t.anyURI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *outboxTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.OrderedCollection != nil {
		i, err = t.OrderedCollection.Serialize()
		return
	}
	if t.anyURI != nil {
		i = anyURISerialize(t.anyURI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// followingTentativeAcceptIntermediateType will only have one of its values set at most
type followingTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible CollectionType type for following property
	Collection CollectionType
	// Stores possible OrderedCollectionType type for following property
	OrderedCollection OrderedCollectionType
	// Stores possible *url.URL type for following property
	anyURI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *followingTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Collection, ok = resolveObject(kind).(CollectionType); t.Collection != nil && ok {
						err = t.Collection.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.OrderedCollection, ok = resolveObject(kind).(OrderedCollectionType); t.OrderedCollection != nil && ok {
						err = t.OrderedCollection.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.anyURI, err = anyURIDeserialize(i)
			if err != nil {
				t.anyURI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *followingTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.Collection != nil {
		i, err = t.Collection.Serialize()
		return
	}
	if t.OrderedCollection != nil {
		i, err = t.OrderedCollection.Serialize()
		return
	}
	if t.anyURI != nil {
		i = anyURISerialize(t.anyURI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// followersTentativeAcceptIntermediateType will only have one of its values set at most
type followersTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible CollectionType type for followers property
	Collection CollectionType
	// Stores possible OrderedCollectionType type for followers property
	OrderedCollection OrderedCollectionType
	// Stores possible *url.URL type for followers property
	anyURI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *followersTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Collection, ok = resolveObject(kind).(CollectionType); t.Collection != nil && ok {
						err = t.Collection.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.OrderedCollection, ok = resolveObject(kind).(OrderedCollectionType); t.OrderedCollection != nil && ok {
						err = t.OrderedCollection.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.anyURI, err = anyURIDeserialize(i)
			if err != nil {
				t.anyURI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *followersTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.Collection != nil {
		i, err = t.Collection.Serialize()
		return
	}
	if t.OrderedCollection != nil {
		i, err = t.OrderedCollection.Serialize()
		return
	}
	if t.anyURI != nil {
		i = anyURISerialize(t.anyURI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// likedTentativeAcceptIntermediateType will only have one of its values set at most
type likedTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible CollectionType type for liked property
	Collection CollectionType
	// Stores possible OrderedCollectionType type for liked property
	OrderedCollection OrderedCollectionType
	// Stores possible *url.URL type for liked property
	anyURI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *likedTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Collection, ok = resolveObject(kind).(CollectionType); t.Collection != nil && ok {
						err = t.Collection.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.OrderedCollection, ok = resolveObject(kind).(OrderedCollectionType); t.OrderedCollection != nil && ok {
						err = t.OrderedCollection.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.anyURI, err = anyURIDeserialize(i)
			if err != nil {
				t.anyURI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *likedTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.Collection != nil {
		i, err = t.Collection.Serialize()
		return
	}
	if t.OrderedCollection != nil {
		i, err = t.OrderedCollection.Serialize()
		return
	}
	if t.anyURI != nil {
		i = anyURISerialize(t.anyURI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// likesTentativeAcceptIntermediateType will only have one of its values set at most
type likesTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible CollectionType type for likes property
	Collection CollectionType
	// Stores possible OrderedCollectionType type for likes property
	OrderedCollection OrderedCollectionType
	// Stores possible *url.URL type for likes property
	anyURI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *likesTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Collection, ok = resolveObject(kind).(CollectionType); t.Collection != nil && ok {
						err = t.Collection.Deserialize(m)
						matched = true
						break
					}
				}
			}
			if !matched {
				for _, kind := range types {
					if t.OrderedCollection, ok = resolveObject(kind).(OrderedCollectionType); t.OrderedCollection != nil && ok {
						err = t.OrderedCollection.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.anyURI, err = anyURIDeserialize(i)
			if err != nil {
				t.anyURI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *likesTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.Collection != nil {
		i, err = t.Collection.Serialize()
		return
	}
	if t.OrderedCollection != nil {
		i, err = t.OrderedCollection.Serialize()
		return
	}
	if t.anyURI != nil {
		i = anyURISerialize(t.anyURI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// preferredUsernameTentativeAcceptIntermediateType will only have one of its values set at most
type preferredUsernameTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *string type for preferredUsername property
	stringName *string
	// Stores possible *url.URL type for preferredUsername property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *preferredUsernameTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.stringName, err = stringDeserialize(i)
			if err != nil {
				t.stringName = nil
			} else {
				matched = true
			}
		}
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *preferredUsernameTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.stringName != nil {
		i = stringSerialize(*t.stringName)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// endpointsTentativeAcceptIntermediateType will only have one of its values set at most
type endpointsTentativeAcceptIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for endpoints property
	Object ObjectType
	// Stores possible *url.URL type for endpoints property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *endpointsTentativeAcceptIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		if tv, ok := m["type"]; ok {
			var types []string
			if tvs, ok := tv.([]interface{}); ok {
				for _, tvi := range tvs {
					if typeString, ok := tvi.(string); ok {
						types = append(types, typeString)
					}
				}
			} else if typeString, ok := tv.(string); ok {
				types = append(types, typeString)
			}
			if !matched {
				for _, kind := range types {
					if t.Object, ok = resolveObject(kind).(ObjectType); t.Object != nil && ok {
						err = t.Object.Deserialize(m)
						matched = true
						break
					}
				}
			}
		} else {
			t.unknown_ = m
		}
	} else if i != nil {
		if !matched {
			t.IRI, err = IRIDeserialize(i)
			if err != nil {
				t.IRI = nil
			} else {
				matched = true
			}
		}
	}
	if !matched {
		t.unknown_ = unknownValueDeserialize(i)
	}
	return

}

// Serialize turns this object into an interface{}.
func (t *endpointsTentativeAcceptIntermediateType) Serialize() (i interface{}, err error) {
	if t.Object != nil {
		i, err = t.Object.Serialize()
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// deserializeactorTentativeAcceptIntermediateType will accept a map to create a actorTentativeAcceptIntermediateType
func deserializeActorTentativeAcceptIntermediateType(in interface{}) (t *actorTentativeAcceptIntermediateType, err error) {
	tmp := &actorTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice actorTentativeAcceptIntermediateType will accept a slice to create a slice of actorTentativeAcceptIntermediateType
func deserializeSliceActorTentativeAcceptIntermediateType(in []interface{}) (t []*actorTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &actorTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeactorTentativeAcceptIntermediateType will accept a actorTentativeAcceptIntermediateType to create a map
func serializeActorTentativeAcceptIntermediateType(t *actorTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceactorTentativeAcceptIntermediateType will accept a slice of actorTentativeAcceptIntermediateType to create a slice result
func serializeSliceActorTentativeAcceptIntermediateType(s []*actorTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeobjectTentativeAcceptIntermediateType will accept a map to create a objectTentativeAcceptIntermediateType
func deserializeObjectTentativeAcceptIntermediateType(in interface{}) (t *objectTentativeAcceptIntermediateType, err error) {
	tmp := &objectTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice objectTentativeAcceptIntermediateType will accept a slice to create a slice of objectTentativeAcceptIntermediateType
func deserializeSliceObjectTentativeAcceptIntermediateType(in []interface{}) (t []*objectTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &objectTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeobjectTentativeAcceptIntermediateType will accept a objectTentativeAcceptIntermediateType to create a map
func serializeObjectTentativeAcceptIntermediateType(t *objectTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceobjectTentativeAcceptIntermediateType will accept a slice of objectTentativeAcceptIntermediateType to create a slice result
func serializeSliceObjectTentativeAcceptIntermediateType(s []*objectTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializetargetTentativeAcceptIntermediateType will accept a map to create a targetTentativeAcceptIntermediateType
func deserializeTargetTentativeAcceptIntermediateType(in interface{}) (t *targetTentativeAcceptIntermediateType, err error) {
	tmp := &targetTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice targetTentativeAcceptIntermediateType will accept a slice to create a slice of targetTentativeAcceptIntermediateType
func deserializeSliceTargetTentativeAcceptIntermediateType(in []interface{}) (t []*targetTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &targetTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializetargetTentativeAcceptIntermediateType will accept a targetTentativeAcceptIntermediateType to create a map
func serializeTargetTentativeAcceptIntermediateType(t *targetTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicetargetTentativeAcceptIntermediateType will accept a slice of targetTentativeAcceptIntermediateType to create a slice result
func serializeSliceTargetTentativeAcceptIntermediateType(s []*targetTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeresultTentativeAcceptIntermediateType will accept a map to create a resultTentativeAcceptIntermediateType
func deserializeResultTentativeAcceptIntermediateType(in interface{}) (t *resultTentativeAcceptIntermediateType, err error) {
	tmp := &resultTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice resultTentativeAcceptIntermediateType will accept a slice to create a slice of resultTentativeAcceptIntermediateType
func deserializeSliceResultTentativeAcceptIntermediateType(in []interface{}) (t []*resultTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &resultTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeresultTentativeAcceptIntermediateType will accept a resultTentativeAcceptIntermediateType to create a map
func serializeResultTentativeAcceptIntermediateType(t *resultTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceresultTentativeAcceptIntermediateType will accept a slice of resultTentativeAcceptIntermediateType to create a slice result
func serializeSliceResultTentativeAcceptIntermediateType(s []*resultTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeoriginTentativeAcceptIntermediateType will accept a map to create a originTentativeAcceptIntermediateType
func deserializeOriginTentativeAcceptIntermediateType(in interface{}) (t *originTentativeAcceptIntermediateType, err error) {
	tmp := &originTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice originTentativeAcceptIntermediateType will accept a slice to create a slice of originTentativeAcceptIntermediateType
func deserializeSliceOriginTentativeAcceptIntermediateType(in []interface{}) (t []*originTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &originTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeoriginTentativeAcceptIntermediateType will accept a originTentativeAcceptIntermediateType to create a map
func serializeOriginTentativeAcceptIntermediateType(t *originTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceoriginTentativeAcceptIntermediateType will accept a slice of originTentativeAcceptIntermediateType to create a slice result
func serializeSliceOriginTentativeAcceptIntermediateType(s []*originTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeinstrumentTentativeAcceptIntermediateType will accept a map to create a instrumentTentativeAcceptIntermediateType
func deserializeInstrumentTentativeAcceptIntermediateType(in interface{}) (t *instrumentTentativeAcceptIntermediateType, err error) {
	tmp := &instrumentTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice instrumentTentativeAcceptIntermediateType will accept a slice to create a slice of instrumentTentativeAcceptIntermediateType
func deserializeSliceInstrumentTentativeAcceptIntermediateType(in []interface{}) (t []*instrumentTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &instrumentTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeinstrumentTentativeAcceptIntermediateType will accept a instrumentTentativeAcceptIntermediateType to create a map
func serializeInstrumentTentativeAcceptIntermediateType(t *instrumentTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceinstrumentTentativeAcceptIntermediateType will accept a slice of instrumentTentativeAcceptIntermediateType to create a slice result
func serializeSliceInstrumentTentativeAcceptIntermediateType(s []*instrumentTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializealtitudeTentativeAcceptIntermediateType will accept a map to create a altitudeTentativeAcceptIntermediateType
func deserializeAltitudeTentativeAcceptIntermediateType(in interface{}) (t *altitudeTentativeAcceptIntermediateType, err error) {
	tmp := &altitudeTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice altitudeTentativeAcceptIntermediateType will accept a slice to create a slice of altitudeTentativeAcceptIntermediateType
func deserializeSliceAltitudeTentativeAcceptIntermediateType(in []interface{}) (t []*altitudeTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &altitudeTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializealtitudeTentativeAcceptIntermediateType will accept a altitudeTentativeAcceptIntermediateType to create a map
func serializeAltitudeTentativeAcceptIntermediateType(t *altitudeTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicealtitudeTentativeAcceptIntermediateType will accept a slice of altitudeTentativeAcceptIntermediateType to create a slice result
func serializeSliceAltitudeTentativeAcceptIntermediateType(s []*altitudeTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeattachmentTentativeAcceptIntermediateType will accept a map to create a attachmentTentativeAcceptIntermediateType
func deserializeAttachmentTentativeAcceptIntermediateType(in interface{}) (t *attachmentTentativeAcceptIntermediateType, err error) {
	tmp := &attachmentTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice attachmentTentativeAcceptIntermediateType will accept a slice to create a slice of attachmentTentativeAcceptIntermediateType
func deserializeSliceAttachmentTentativeAcceptIntermediateType(in []interface{}) (t []*attachmentTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &attachmentTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeattachmentTentativeAcceptIntermediateType will accept a attachmentTentativeAcceptIntermediateType to create a map
func serializeAttachmentTentativeAcceptIntermediateType(t *attachmentTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceattachmentTentativeAcceptIntermediateType will accept a slice of attachmentTentativeAcceptIntermediateType to create a slice result
func serializeSliceAttachmentTentativeAcceptIntermediateType(s []*attachmentTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeattributedToTentativeAcceptIntermediateType will accept a map to create a attributedToTentativeAcceptIntermediateType
func deserializeAttributedToTentativeAcceptIntermediateType(in interface{}) (t *attributedToTentativeAcceptIntermediateType, err error) {
	tmp := &attributedToTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice attributedToTentativeAcceptIntermediateType will accept a slice to create a slice of attributedToTentativeAcceptIntermediateType
func deserializeSliceAttributedToTentativeAcceptIntermediateType(in []interface{}) (t []*attributedToTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &attributedToTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeattributedToTentativeAcceptIntermediateType will accept a attributedToTentativeAcceptIntermediateType to create a map
func serializeAttributedToTentativeAcceptIntermediateType(t *attributedToTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceattributedToTentativeAcceptIntermediateType will accept a slice of attributedToTentativeAcceptIntermediateType to create a slice result
func serializeSliceAttributedToTentativeAcceptIntermediateType(s []*attributedToTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeaudienceTentativeAcceptIntermediateType will accept a map to create a audienceTentativeAcceptIntermediateType
func deserializeAudienceTentativeAcceptIntermediateType(in interface{}) (t *audienceTentativeAcceptIntermediateType, err error) {
	tmp := &audienceTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice audienceTentativeAcceptIntermediateType will accept a slice to create a slice of audienceTentativeAcceptIntermediateType
func deserializeSliceAudienceTentativeAcceptIntermediateType(in []interface{}) (t []*audienceTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &audienceTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeaudienceTentativeAcceptIntermediateType will accept a audienceTentativeAcceptIntermediateType to create a map
func serializeAudienceTentativeAcceptIntermediateType(t *audienceTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceaudienceTentativeAcceptIntermediateType will accept a slice of audienceTentativeAcceptIntermediateType to create a slice result
func serializeSliceAudienceTentativeAcceptIntermediateType(s []*audienceTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializecontentTentativeAcceptIntermediateType will accept a map to create a contentTentativeAcceptIntermediateType
func deserializeContentTentativeAcceptIntermediateType(in interface{}) (t *contentTentativeAcceptIntermediateType, err error) {
	tmp := &contentTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice contentTentativeAcceptIntermediateType will accept a slice to create a slice of contentTentativeAcceptIntermediateType
func deserializeSliceContentTentativeAcceptIntermediateType(in []interface{}) (t []*contentTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &contentTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializecontentTentativeAcceptIntermediateType will accept a contentTentativeAcceptIntermediateType to create a map
func serializeContentTentativeAcceptIntermediateType(t *contentTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicecontentTentativeAcceptIntermediateType will accept a slice of contentTentativeAcceptIntermediateType to create a slice result
func serializeSliceContentTentativeAcceptIntermediateType(s []*contentTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializecontextTentativeAcceptIntermediateType will accept a map to create a contextTentativeAcceptIntermediateType
func deserializeContextTentativeAcceptIntermediateType(in interface{}) (t *contextTentativeAcceptIntermediateType, err error) {
	tmp := &contextTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice contextTentativeAcceptIntermediateType will accept a slice to create a slice of contextTentativeAcceptIntermediateType
func deserializeSliceContextTentativeAcceptIntermediateType(in []interface{}) (t []*contextTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &contextTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializecontextTentativeAcceptIntermediateType will accept a contextTentativeAcceptIntermediateType to create a map
func serializeContextTentativeAcceptIntermediateType(t *contextTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicecontextTentativeAcceptIntermediateType will accept a slice of contextTentativeAcceptIntermediateType to create a slice result
func serializeSliceContextTentativeAcceptIntermediateType(s []*contextTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializenameTentativeAcceptIntermediateType will accept a map to create a nameTentativeAcceptIntermediateType
func deserializeNameTentativeAcceptIntermediateType(in interface{}) (t *nameTentativeAcceptIntermediateType, err error) {
	tmp := &nameTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice nameTentativeAcceptIntermediateType will accept a slice to create a slice of nameTentativeAcceptIntermediateType
func deserializeSliceNameTentativeAcceptIntermediateType(in []interface{}) (t []*nameTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &nameTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializenameTentativeAcceptIntermediateType will accept a nameTentativeAcceptIntermediateType to create a map
func serializeNameTentativeAcceptIntermediateType(t *nameTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicenameTentativeAcceptIntermediateType will accept a slice of nameTentativeAcceptIntermediateType to create a slice result
func serializeSliceNameTentativeAcceptIntermediateType(s []*nameTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeendTimeTentativeAcceptIntermediateType will accept a map to create a endTimeTentativeAcceptIntermediateType
func deserializeEndTimeTentativeAcceptIntermediateType(in interface{}) (t *endTimeTentativeAcceptIntermediateType, err error) {
	tmp := &endTimeTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice endTimeTentativeAcceptIntermediateType will accept a slice to create a slice of endTimeTentativeAcceptIntermediateType
func deserializeSliceEndTimeTentativeAcceptIntermediateType(in []interface{}) (t []*endTimeTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &endTimeTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeendTimeTentativeAcceptIntermediateType will accept a endTimeTentativeAcceptIntermediateType to create a map
func serializeEndTimeTentativeAcceptIntermediateType(t *endTimeTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceendTimeTentativeAcceptIntermediateType will accept a slice of endTimeTentativeAcceptIntermediateType to create a slice result
func serializeSliceEndTimeTentativeAcceptIntermediateType(s []*endTimeTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializegeneratorTentativeAcceptIntermediateType will accept a map to create a generatorTentativeAcceptIntermediateType
func deserializeGeneratorTentativeAcceptIntermediateType(in interface{}) (t *generatorTentativeAcceptIntermediateType, err error) {
	tmp := &generatorTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice generatorTentativeAcceptIntermediateType will accept a slice to create a slice of generatorTentativeAcceptIntermediateType
func deserializeSliceGeneratorTentativeAcceptIntermediateType(in []interface{}) (t []*generatorTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &generatorTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializegeneratorTentativeAcceptIntermediateType will accept a generatorTentativeAcceptIntermediateType to create a map
func serializeGeneratorTentativeAcceptIntermediateType(t *generatorTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicegeneratorTentativeAcceptIntermediateType will accept a slice of generatorTentativeAcceptIntermediateType to create a slice result
func serializeSliceGeneratorTentativeAcceptIntermediateType(s []*generatorTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeiconTentativeAcceptIntermediateType will accept a map to create a iconTentativeAcceptIntermediateType
func deserializeIconTentativeAcceptIntermediateType(in interface{}) (t *iconTentativeAcceptIntermediateType, err error) {
	tmp := &iconTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice iconTentativeAcceptIntermediateType will accept a slice to create a slice of iconTentativeAcceptIntermediateType
func deserializeSliceIconTentativeAcceptIntermediateType(in []interface{}) (t []*iconTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &iconTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeiconTentativeAcceptIntermediateType will accept a iconTentativeAcceptIntermediateType to create a map
func serializeIconTentativeAcceptIntermediateType(t *iconTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceiconTentativeAcceptIntermediateType will accept a slice of iconTentativeAcceptIntermediateType to create a slice result
func serializeSliceIconTentativeAcceptIntermediateType(s []*iconTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeimageTentativeAcceptIntermediateType will accept a map to create a imageTentativeAcceptIntermediateType
func deserializeImageTentativeAcceptIntermediateType(in interface{}) (t *imageTentativeAcceptIntermediateType, err error) {
	tmp := &imageTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice imageTentativeAcceptIntermediateType will accept a slice to create a slice of imageTentativeAcceptIntermediateType
func deserializeSliceImageTentativeAcceptIntermediateType(in []interface{}) (t []*imageTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &imageTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeimageTentativeAcceptIntermediateType will accept a imageTentativeAcceptIntermediateType to create a map
func serializeImageTentativeAcceptIntermediateType(t *imageTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceimageTentativeAcceptIntermediateType will accept a slice of imageTentativeAcceptIntermediateType to create a slice result
func serializeSliceImageTentativeAcceptIntermediateType(s []*imageTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeinReplyToTentativeAcceptIntermediateType will accept a map to create a inReplyToTentativeAcceptIntermediateType
func deserializeInReplyToTentativeAcceptIntermediateType(in interface{}) (t *inReplyToTentativeAcceptIntermediateType, err error) {
	tmp := &inReplyToTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice inReplyToTentativeAcceptIntermediateType will accept a slice to create a slice of inReplyToTentativeAcceptIntermediateType
func deserializeSliceInReplyToTentativeAcceptIntermediateType(in []interface{}) (t []*inReplyToTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &inReplyToTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeinReplyToTentativeAcceptIntermediateType will accept a inReplyToTentativeAcceptIntermediateType to create a map
func serializeInReplyToTentativeAcceptIntermediateType(t *inReplyToTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceinReplyToTentativeAcceptIntermediateType will accept a slice of inReplyToTentativeAcceptIntermediateType to create a slice result
func serializeSliceInReplyToTentativeAcceptIntermediateType(s []*inReplyToTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializelocationTentativeAcceptIntermediateType will accept a map to create a locationTentativeAcceptIntermediateType
func deserializeLocationTentativeAcceptIntermediateType(in interface{}) (t *locationTentativeAcceptIntermediateType, err error) {
	tmp := &locationTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice locationTentativeAcceptIntermediateType will accept a slice to create a slice of locationTentativeAcceptIntermediateType
func deserializeSliceLocationTentativeAcceptIntermediateType(in []interface{}) (t []*locationTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &locationTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializelocationTentativeAcceptIntermediateType will accept a locationTentativeAcceptIntermediateType to create a map
func serializeLocationTentativeAcceptIntermediateType(t *locationTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicelocationTentativeAcceptIntermediateType will accept a slice of locationTentativeAcceptIntermediateType to create a slice result
func serializeSliceLocationTentativeAcceptIntermediateType(s []*locationTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializepreviewTentativeAcceptIntermediateType will accept a map to create a previewTentativeAcceptIntermediateType
func deserializePreviewTentativeAcceptIntermediateType(in interface{}) (t *previewTentativeAcceptIntermediateType, err error) {
	tmp := &previewTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice previewTentativeAcceptIntermediateType will accept a slice to create a slice of previewTentativeAcceptIntermediateType
func deserializeSlicePreviewTentativeAcceptIntermediateType(in []interface{}) (t []*previewTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &previewTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializepreviewTentativeAcceptIntermediateType will accept a previewTentativeAcceptIntermediateType to create a map
func serializePreviewTentativeAcceptIntermediateType(t *previewTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicepreviewTentativeAcceptIntermediateType will accept a slice of previewTentativeAcceptIntermediateType to create a slice result
func serializeSlicePreviewTentativeAcceptIntermediateType(s []*previewTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializepublishedTentativeAcceptIntermediateType will accept a map to create a publishedTentativeAcceptIntermediateType
func deserializePublishedTentativeAcceptIntermediateType(in interface{}) (t *publishedTentativeAcceptIntermediateType, err error) {
	tmp := &publishedTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice publishedTentativeAcceptIntermediateType will accept a slice to create a slice of publishedTentativeAcceptIntermediateType
func deserializeSlicePublishedTentativeAcceptIntermediateType(in []interface{}) (t []*publishedTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &publishedTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializepublishedTentativeAcceptIntermediateType will accept a publishedTentativeAcceptIntermediateType to create a map
func serializePublishedTentativeAcceptIntermediateType(t *publishedTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicepublishedTentativeAcceptIntermediateType will accept a slice of publishedTentativeAcceptIntermediateType to create a slice result
func serializeSlicePublishedTentativeAcceptIntermediateType(s []*publishedTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializerepliesTentativeAcceptIntermediateType will accept a map to create a repliesTentativeAcceptIntermediateType
func deserializeRepliesTentativeAcceptIntermediateType(in interface{}) (t *repliesTentativeAcceptIntermediateType, err error) {
	tmp := &repliesTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice repliesTentativeAcceptIntermediateType will accept a slice to create a slice of repliesTentativeAcceptIntermediateType
func deserializeSliceRepliesTentativeAcceptIntermediateType(in []interface{}) (t []*repliesTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &repliesTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializerepliesTentativeAcceptIntermediateType will accept a repliesTentativeAcceptIntermediateType to create a map
func serializeRepliesTentativeAcceptIntermediateType(t *repliesTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicerepliesTentativeAcceptIntermediateType will accept a slice of repliesTentativeAcceptIntermediateType to create a slice result
func serializeSliceRepliesTentativeAcceptIntermediateType(s []*repliesTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializestartTimeTentativeAcceptIntermediateType will accept a map to create a startTimeTentativeAcceptIntermediateType
func deserializeStartTimeTentativeAcceptIntermediateType(in interface{}) (t *startTimeTentativeAcceptIntermediateType, err error) {
	tmp := &startTimeTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice startTimeTentativeAcceptIntermediateType will accept a slice to create a slice of startTimeTentativeAcceptIntermediateType
func deserializeSliceStartTimeTentativeAcceptIntermediateType(in []interface{}) (t []*startTimeTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &startTimeTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializestartTimeTentativeAcceptIntermediateType will accept a startTimeTentativeAcceptIntermediateType to create a map
func serializeStartTimeTentativeAcceptIntermediateType(t *startTimeTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicestartTimeTentativeAcceptIntermediateType will accept a slice of startTimeTentativeAcceptIntermediateType to create a slice result
func serializeSliceStartTimeTentativeAcceptIntermediateType(s []*startTimeTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializesummaryTentativeAcceptIntermediateType will accept a map to create a summaryTentativeAcceptIntermediateType
func deserializeSummaryTentativeAcceptIntermediateType(in interface{}) (t *summaryTentativeAcceptIntermediateType, err error) {
	tmp := &summaryTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice summaryTentativeAcceptIntermediateType will accept a slice to create a slice of summaryTentativeAcceptIntermediateType
func deserializeSliceSummaryTentativeAcceptIntermediateType(in []interface{}) (t []*summaryTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &summaryTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializesummaryTentativeAcceptIntermediateType will accept a summaryTentativeAcceptIntermediateType to create a map
func serializeSummaryTentativeAcceptIntermediateType(t *summaryTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicesummaryTentativeAcceptIntermediateType will accept a slice of summaryTentativeAcceptIntermediateType to create a slice result
func serializeSliceSummaryTentativeAcceptIntermediateType(s []*summaryTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializetagTentativeAcceptIntermediateType will accept a map to create a tagTentativeAcceptIntermediateType
func deserializeTagTentativeAcceptIntermediateType(in interface{}) (t *tagTentativeAcceptIntermediateType, err error) {
	tmp := &tagTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice tagTentativeAcceptIntermediateType will accept a slice to create a slice of tagTentativeAcceptIntermediateType
func deserializeSliceTagTentativeAcceptIntermediateType(in []interface{}) (t []*tagTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &tagTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializetagTentativeAcceptIntermediateType will accept a tagTentativeAcceptIntermediateType to create a map
func serializeTagTentativeAcceptIntermediateType(t *tagTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicetagTentativeAcceptIntermediateType will accept a slice of tagTentativeAcceptIntermediateType to create a slice result
func serializeSliceTagTentativeAcceptIntermediateType(s []*tagTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeupdatedTentativeAcceptIntermediateType will accept a map to create a updatedTentativeAcceptIntermediateType
func deserializeUpdatedTentativeAcceptIntermediateType(in interface{}) (t *updatedTentativeAcceptIntermediateType, err error) {
	tmp := &updatedTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice updatedTentativeAcceptIntermediateType will accept a slice to create a slice of updatedTentativeAcceptIntermediateType
func deserializeSliceUpdatedTentativeAcceptIntermediateType(in []interface{}) (t []*updatedTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &updatedTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeupdatedTentativeAcceptIntermediateType will accept a updatedTentativeAcceptIntermediateType to create a map
func serializeUpdatedTentativeAcceptIntermediateType(t *updatedTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceupdatedTentativeAcceptIntermediateType will accept a slice of updatedTentativeAcceptIntermediateType to create a slice result
func serializeSliceUpdatedTentativeAcceptIntermediateType(s []*updatedTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeurlTentativeAcceptIntermediateType will accept a map to create a urlTentativeAcceptIntermediateType
func deserializeUrlTentativeAcceptIntermediateType(in interface{}) (t *urlTentativeAcceptIntermediateType, err error) {
	tmp := &urlTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice urlTentativeAcceptIntermediateType will accept a slice to create a slice of urlTentativeAcceptIntermediateType
func deserializeSliceUrlTentativeAcceptIntermediateType(in []interface{}) (t []*urlTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &urlTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeurlTentativeAcceptIntermediateType will accept a urlTentativeAcceptIntermediateType to create a map
func serializeUrlTentativeAcceptIntermediateType(t *urlTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceurlTentativeAcceptIntermediateType will accept a slice of urlTentativeAcceptIntermediateType to create a slice result
func serializeSliceUrlTentativeAcceptIntermediateType(s []*urlTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializetoTentativeAcceptIntermediateType will accept a map to create a toTentativeAcceptIntermediateType
func deserializeToTentativeAcceptIntermediateType(in interface{}) (t *toTentativeAcceptIntermediateType, err error) {
	tmp := &toTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice toTentativeAcceptIntermediateType will accept a slice to create a slice of toTentativeAcceptIntermediateType
func deserializeSliceToTentativeAcceptIntermediateType(in []interface{}) (t []*toTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &toTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializetoTentativeAcceptIntermediateType will accept a toTentativeAcceptIntermediateType to create a map
func serializeToTentativeAcceptIntermediateType(t *toTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicetoTentativeAcceptIntermediateType will accept a slice of toTentativeAcceptIntermediateType to create a slice result
func serializeSliceToTentativeAcceptIntermediateType(s []*toTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializebtoTentativeAcceptIntermediateType will accept a map to create a btoTentativeAcceptIntermediateType
func deserializeBtoTentativeAcceptIntermediateType(in interface{}) (t *btoTentativeAcceptIntermediateType, err error) {
	tmp := &btoTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice btoTentativeAcceptIntermediateType will accept a slice to create a slice of btoTentativeAcceptIntermediateType
func deserializeSliceBtoTentativeAcceptIntermediateType(in []interface{}) (t []*btoTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &btoTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializebtoTentativeAcceptIntermediateType will accept a btoTentativeAcceptIntermediateType to create a map
func serializeBtoTentativeAcceptIntermediateType(t *btoTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicebtoTentativeAcceptIntermediateType will accept a slice of btoTentativeAcceptIntermediateType to create a slice result
func serializeSliceBtoTentativeAcceptIntermediateType(s []*btoTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeccTentativeAcceptIntermediateType will accept a map to create a ccTentativeAcceptIntermediateType
func deserializeCcTentativeAcceptIntermediateType(in interface{}) (t *ccTentativeAcceptIntermediateType, err error) {
	tmp := &ccTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice ccTentativeAcceptIntermediateType will accept a slice to create a slice of ccTentativeAcceptIntermediateType
func deserializeSliceCcTentativeAcceptIntermediateType(in []interface{}) (t []*ccTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &ccTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeccTentativeAcceptIntermediateType will accept a ccTentativeAcceptIntermediateType to create a map
func serializeCcTentativeAcceptIntermediateType(t *ccTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceccTentativeAcceptIntermediateType will accept a slice of ccTentativeAcceptIntermediateType to create a slice result
func serializeSliceCcTentativeAcceptIntermediateType(s []*ccTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializebccTentativeAcceptIntermediateType will accept a map to create a bccTentativeAcceptIntermediateType
func deserializeBccTentativeAcceptIntermediateType(in interface{}) (t *bccTentativeAcceptIntermediateType, err error) {
	tmp := &bccTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice bccTentativeAcceptIntermediateType will accept a slice to create a slice of bccTentativeAcceptIntermediateType
func deserializeSliceBccTentativeAcceptIntermediateType(in []interface{}) (t []*bccTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &bccTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializebccTentativeAcceptIntermediateType will accept a bccTentativeAcceptIntermediateType to create a map
func serializeBccTentativeAcceptIntermediateType(t *bccTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicebccTentativeAcceptIntermediateType will accept a slice of bccTentativeAcceptIntermediateType to create a slice result
func serializeSliceBccTentativeAcceptIntermediateType(s []*bccTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializemediaTypeTentativeAcceptIntermediateType will accept a map to create a mediaTypeTentativeAcceptIntermediateType
func deserializeMediaTypeTentativeAcceptIntermediateType(in interface{}) (t *mediaTypeTentativeAcceptIntermediateType, err error) {
	tmp := &mediaTypeTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice mediaTypeTentativeAcceptIntermediateType will accept a slice to create a slice of mediaTypeTentativeAcceptIntermediateType
func deserializeSliceMediaTypeTentativeAcceptIntermediateType(in []interface{}) (t []*mediaTypeTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &mediaTypeTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializemediaTypeTentativeAcceptIntermediateType will accept a mediaTypeTentativeAcceptIntermediateType to create a map
func serializeMediaTypeTentativeAcceptIntermediateType(t *mediaTypeTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicemediaTypeTentativeAcceptIntermediateType will accept a slice of mediaTypeTentativeAcceptIntermediateType to create a slice result
func serializeSliceMediaTypeTentativeAcceptIntermediateType(s []*mediaTypeTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializedurationTentativeAcceptIntermediateType will accept a map to create a durationTentativeAcceptIntermediateType
func deserializeDurationTentativeAcceptIntermediateType(in interface{}) (t *durationTentativeAcceptIntermediateType, err error) {
	tmp := &durationTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice durationTentativeAcceptIntermediateType will accept a slice to create a slice of durationTentativeAcceptIntermediateType
func deserializeSliceDurationTentativeAcceptIntermediateType(in []interface{}) (t []*durationTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &durationTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializedurationTentativeAcceptIntermediateType will accept a durationTentativeAcceptIntermediateType to create a map
func serializeDurationTentativeAcceptIntermediateType(t *durationTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicedurationTentativeAcceptIntermediateType will accept a slice of durationTentativeAcceptIntermediateType to create a slice result
func serializeSliceDurationTentativeAcceptIntermediateType(s []*durationTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializesourceTentativeAcceptIntermediateType will accept a map to create a sourceTentativeAcceptIntermediateType
func deserializeSourceTentativeAcceptIntermediateType(in interface{}) (t *sourceTentativeAcceptIntermediateType, err error) {
	tmp := &sourceTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice sourceTentativeAcceptIntermediateType will accept a slice to create a slice of sourceTentativeAcceptIntermediateType
func deserializeSliceSourceTentativeAcceptIntermediateType(in []interface{}) (t []*sourceTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &sourceTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializesourceTentativeAcceptIntermediateType will accept a sourceTentativeAcceptIntermediateType to create a map
func serializeSourceTentativeAcceptIntermediateType(t *sourceTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicesourceTentativeAcceptIntermediateType will accept a slice of sourceTentativeAcceptIntermediateType to create a slice result
func serializeSliceSourceTentativeAcceptIntermediateType(s []*sourceTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeinboxTentativeAcceptIntermediateType will accept a map to create a inboxTentativeAcceptIntermediateType
func deserializeInboxTentativeAcceptIntermediateType(in interface{}) (t *inboxTentativeAcceptIntermediateType, err error) {
	tmp := &inboxTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice inboxTentativeAcceptIntermediateType will accept a slice to create a slice of inboxTentativeAcceptIntermediateType
func deserializeSliceInboxTentativeAcceptIntermediateType(in []interface{}) (t []*inboxTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &inboxTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeinboxTentativeAcceptIntermediateType will accept a inboxTentativeAcceptIntermediateType to create a map
func serializeInboxTentativeAcceptIntermediateType(t *inboxTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceinboxTentativeAcceptIntermediateType will accept a slice of inboxTentativeAcceptIntermediateType to create a slice result
func serializeSliceInboxTentativeAcceptIntermediateType(s []*inboxTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeoutboxTentativeAcceptIntermediateType will accept a map to create a outboxTentativeAcceptIntermediateType
func deserializeOutboxTentativeAcceptIntermediateType(in interface{}) (t *outboxTentativeAcceptIntermediateType, err error) {
	tmp := &outboxTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice outboxTentativeAcceptIntermediateType will accept a slice to create a slice of outboxTentativeAcceptIntermediateType
func deserializeSliceOutboxTentativeAcceptIntermediateType(in []interface{}) (t []*outboxTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &outboxTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeoutboxTentativeAcceptIntermediateType will accept a outboxTentativeAcceptIntermediateType to create a map
func serializeOutboxTentativeAcceptIntermediateType(t *outboxTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceoutboxTentativeAcceptIntermediateType will accept a slice of outboxTentativeAcceptIntermediateType to create a slice result
func serializeSliceOutboxTentativeAcceptIntermediateType(s []*outboxTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializefollowingTentativeAcceptIntermediateType will accept a map to create a followingTentativeAcceptIntermediateType
func deserializeFollowingTentativeAcceptIntermediateType(in interface{}) (t *followingTentativeAcceptIntermediateType, err error) {
	tmp := &followingTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice followingTentativeAcceptIntermediateType will accept a slice to create a slice of followingTentativeAcceptIntermediateType
func deserializeSliceFollowingTentativeAcceptIntermediateType(in []interface{}) (t []*followingTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &followingTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializefollowingTentativeAcceptIntermediateType will accept a followingTentativeAcceptIntermediateType to create a map
func serializeFollowingTentativeAcceptIntermediateType(t *followingTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicefollowingTentativeAcceptIntermediateType will accept a slice of followingTentativeAcceptIntermediateType to create a slice result
func serializeSliceFollowingTentativeAcceptIntermediateType(s []*followingTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializefollowersTentativeAcceptIntermediateType will accept a map to create a followersTentativeAcceptIntermediateType
func deserializeFollowersTentativeAcceptIntermediateType(in interface{}) (t *followersTentativeAcceptIntermediateType, err error) {
	tmp := &followersTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice followersTentativeAcceptIntermediateType will accept a slice to create a slice of followersTentativeAcceptIntermediateType
func deserializeSliceFollowersTentativeAcceptIntermediateType(in []interface{}) (t []*followersTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &followersTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializefollowersTentativeAcceptIntermediateType will accept a followersTentativeAcceptIntermediateType to create a map
func serializeFollowersTentativeAcceptIntermediateType(t *followersTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicefollowersTentativeAcceptIntermediateType will accept a slice of followersTentativeAcceptIntermediateType to create a slice result
func serializeSliceFollowersTentativeAcceptIntermediateType(s []*followersTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializelikedTentativeAcceptIntermediateType will accept a map to create a likedTentativeAcceptIntermediateType
func deserializeLikedTentativeAcceptIntermediateType(in interface{}) (t *likedTentativeAcceptIntermediateType, err error) {
	tmp := &likedTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice likedTentativeAcceptIntermediateType will accept a slice to create a slice of likedTentativeAcceptIntermediateType
func deserializeSliceLikedTentativeAcceptIntermediateType(in []interface{}) (t []*likedTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &likedTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializelikedTentativeAcceptIntermediateType will accept a likedTentativeAcceptIntermediateType to create a map
func serializeLikedTentativeAcceptIntermediateType(t *likedTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicelikedTentativeAcceptIntermediateType will accept a slice of likedTentativeAcceptIntermediateType to create a slice result
func serializeSliceLikedTentativeAcceptIntermediateType(s []*likedTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializelikesTentativeAcceptIntermediateType will accept a map to create a likesTentativeAcceptIntermediateType
func deserializeLikesTentativeAcceptIntermediateType(in interface{}) (t *likesTentativeAcceptIntermediateType, err error) {
	tmp := &likesTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice likesTentativeAcceptIntermediateType will accept a slice to create a slice of likesTentativeAcceptIntermediateType
func deserializeSliceLikesTentativeAcceptIntermediateType(in []interface{}) (t []*likesTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &likesTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializelikesTentativeAcceptIntermediateType will accept a likesTentativeAcceptIntermediateType to create a map
func serializeLikesTentativeAcceptIntermediateType(t *likesTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicelikesTentativeAcceptIntermediateType will accept a slice of likesTentativeAcceptIntermediateType to create a slice result
func serializeSliceLikesTentativeAcceptIntermediateType(s []*likesTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializepreferredUsernameTentativeAcceptIntermediateType will accept a map to create a preferredUsernameTentativeAcceptIntermediateType
func deserializePreferredUsernameTentativeAcceptIntermediateType(in interface{}) (t *preferredUsernameTentativeAcceptIntermediateType, err error) {
	tmp := &preferredUsernameTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice preferredUsernameTentativeAcceptIntermediateType will accept a slice to create a slice of preferredUsernameTentativeAcceptIntermediateType
func deserializeSlicePreferredUsernameTentativeAcceptIntermediateType(in []interface{}) (t []*preferredUsernameTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &preferredUsernameTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializepreferredUsernameTentativeAcceptIntermediateType will accept a preferredUsernameTentativeAcceptIntermediateType to create a map
func serializePreferredUsernameTentativeAcceptIntermediateType(t *preferredUsernameTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicepreferredUsernameTentativeAcceptIntermediateType will accept a slice of preferredUsernameTentativeAcceptIntermediateType to create a slice result
func serializeSlicePreferredUsernameTentativeAcceptIntermediateType(s []*preferredUsernameTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeendpointsTentativeAcceptIntermediateType will accept a map to create a endpointsTentativeAcceptIntermediateType
func deserializeEndpointsTentativeAcceptIntermediateType(in interface{}) (t *endpointsTentativeAcceptIntermediateType, err error) {
	tmp := &endpointsTentativeAcceptIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice endpointsTentativeAcceptIntermediateType will accept a slice to create a slice of endpointsTentativeAcceptIntermediateType
func deserializeSliceEndpointsTentativeAcceptIntermediateType(in []interface{}) (t []*endpointsTentativeAcceptIntermediateType, err error) {
	for _, i := range in {
		tmp := &endpointsTentativeAcceptIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeendpointsTentativeAcceptIntermediateType will accept a endpointsTentativeAcceptIntermediateType to create a map
func serializeEndpointsTentativeAcceptIntermediateType(t *endpointsTentativeAcceptIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceendpointsTentativeAcceptIntermediateType will accept a slice of endpointsTentativeAcceptIntermediateType to create a slice result
func serializeSliceEndpointsTentativeAcceptIntermediateType(s []*endpointsTentativeAcceptIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

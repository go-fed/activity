//
package vocab

import (
	"fmt"
	"net/url"
	"time"
)

// IntransitiveActivityType is an interface for accepting types that extend from 'IntransitiveActivity'.
type IntransitiveActivityType interface {
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
}

// Instances of IntransitiveActivity are a subtype of Activity representing intransitive actions. The object property is therefore inappropriate for these activities.
type IntransitiveActivity struct {
	// An unknown value.
	unknown_ map[string]interface{}
	// The 'actor' value could have multiple types and values
	actor []*actorIntransitiveActivityIntermediateType
	// The 'target' value could have multiple types and values
	target []*targetIntransitiveActivityIntermediateType
	// The 'result' value could have multiple types and values
	result []*resultIntransitiveActivityIntermediateType
	// The 'origin' value could have multiple types and values
	origin []*originIntransitiveActivityIntermediateType
	// The 'instrument' value could have multiple types and values
	instrument []*instrumentIntransitiveActivityIntermediateType
	// The functional 'altitude' value could have multiple types, but only a single value
	altitude *altitudeIntransitiveActivityIntermediateType
	// The 'attachment' value could have multiple types and values
	attachment []*attachmentIntransitiveActivityIntermediateType
	// The 'attributedTo' value could have multiple types and values
	attributedTo []*attributedToIntransitiveActivityIntermediateType
	// The 'audience' value could have multiple types and values
	audience []*audienceIntransitiveActivityIntermediateType
	// The 'content' value could have multiple types and values
	content []*contentIntransitiveActivityIntermediateType
	// The 'contentMap' value holds language-specific values for property 'content'
	contentMap map[string]string
	// The 'context' value could have multiple types and values
	context []*contextIntransitiveActivityIntermediateType
	// The 'name' value could have multiple types and values
	name []*nameIntransitiveActivityIntermediateType
	// The 'nameMap' value holds language-specific values for property 'name'
	nameMap map[string]string
	// The functional 'endTime' value could have multiple types, but only a single value
	endTime *endTimeIntransitiveActivityIntermediateType
	// The 'generator' value could have multiple types and values
	generator []*generatorIntransitiveActivityIntermediateType
	// The 'icon' value could have multiple types and values
	icon []*iconIntransitiveActivityIntermediateType
	// The functional 'id' value holds a single type and a single value
	id *url.URL
	// The 'image' value could have multiple types and values
	image []*imageIntransitiveActivityIntermediateType
	// The 'inReplyTo' value could have multiple types and values
	inReplyTo []*inReplyToIntransitiveActivityIntermediateType
	// The 'location' value could have multiple types and values
	location []*locationIntransitiveActivityIntermediateType
	// The 'preview' value could have multiple types and values
	preview []*previewIntransitiveActivityIntermediateType
	// The functional 'published' value could have multiple types, but only a single value
	published *publishedIntransitiveActivityIntermediateType
	// The functional 'replies' value could have multiple types, but only a single value
	replies *repliesIntransitiveActivityIntermediateType
	// The functional 'startTime' value could have multiple types, but only a single value
	startTime *startTimeIntransitiveActivityIntermediateType
	// The 'summary' value could have multiple types and values
	summary []*summaryIntransitiveActivityIntermediateType
	// The 'summaryMap' value holds language-specific values for property 'summary'
	summaryMap map[string]string
	// The 'tag' value could have multiple types and values
	tag []*tagIntransitiveActivityIntermediateType
	// The 'type' value can hold any type and any number of values
	typeName []interface{}
	// The functional 'updated' value could have multiple types, but only a single value
	updated *updatedIntransitiveActivityIntermediateType
	// The 'url' value could have multiple types and values
	url []*urlIntransitiveActivityIntermediateType
	// The 'to' value could have multiple types and values
	to []*toIntransitiveActivityIntermediateType
	// The 'bto' value could have multiple types and values
	bto []*btoIntransitiveActivityIntermediateType
	// The 'cc' value could have multiple types and values
	cc []*ccIntransitiveActivityIntermediateType
	// The 'bcc' value could have multiple types and values
	bcc []*bccIntransitiveActivityIntermediateType
	// The functional 'mediaType' value could have multiple types, but only a single value
	mediaType *mediaTypeIntransitiveActivityIntermediateType
	// The functional 'duration' value could have multiple types, but only a single value
	duration *durationIntransitiveActivityIntermediateType
	// The functional 'source' value could have multiple types, but only a single value
	source *sourceIntransitiveActivityIntermediateType
	// The functional 'inbox' value could have multiple types, but only a single value
	inbox *inboxIntransitiveActivityIntermediateType
	// The functional 'outbox' value could have multiple types, but only a single value
	outbox *outboxIntransitiveActivityIntermediateType
	// The functional 'following' value could have multiple types, but only a single value
	following *followingIntransitiveActivityIntermediateType
	// The functional 'followers' value could have multiple types, but only a single value
	followers *followersIntransitiveActivityIntermediateType
	// The functional 'liked' value could have multiple types, but only a single value
	liked *likedIntransitiveActivityIntermediateType
	// The functional 'likes' value could have multiple types, but only a single value
	likes *likesIntransitiveActivityIntermediateType
	// The 'streams' value holds a single type and any number of values
	streams []*url.URL
	// The functional 'preferredUsername' value could have multiple types, but only a single value
	preferredUsername *preferredUsernameIntransitiveActivityIntermediateType
	// The 'preferredUsernameMap' value holds language-specific values for property 'preferredUsername'
	preferredUsernameMap map[string]string
	// The functional 'endpoints' value could have multiple types, but only a single value
	endpoints *endpointsIntransitiveActivityIntermediateType
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
func (t *IntransitiveActivity) ActorLen() (l int) {
	return len(t.actor)

}

// IsActorObject determines whether the call to GetActorObject is safe for the specified index
func (t *IntransitiveActivity) IsActorObject(index int) (ok bool) {
	return t.actor[index].Object != nil

}

// GetActorObject returns the value safely if IsActorObject returned true for the specified index
func (t *IntransitiveActivity) GetActorObject(index int) (v ObjectType) {
	return t.actor[index].Object

}

// AppendActorObject adds to the back of actor a ObjectType type
func (t *IntransitiveActivity) AppendActorObject(v ObjectType) {
	t.actor = append(t.actor, &actorIntransitiveActivityIntermediateType{Object: v})

}

// PrependActorObject adds to the front of actor a ObjectType type
func (t *IntransitiveActivity) PrependActorObject(v ObjectType) {
	t.actor = append([]*actorIntransitiveActivityIntermediateType{&actorIntransitiveActivityIntermediateType{Object: v}}, t.actor...)

}

// RemoveActorObject deletes the value from the specified index
func (t *IntransitiveActivity) RemoveActorObject(index int) {
	copy(t.actor[index:], t.actor[index+1:])
	t.actor[len(t.actor)-1] = nil
	t.actor = t.actor[:len(t.actor)-1]

}

// IsActorLink determines whether the call to GetActorLink is safe for the specified index
func (t *IntransitiveActivity) IsActorLink(index int) (ok bool) {
	return t.actor[index].Link != nil

}

// GetActorLink returns the value safely if IsActorLink returned true for the specified index
func (t *IntransitiveActivity) GetActorLink(index int) (v LinkType) {
	return t.actor[index].Link

}

// AppendActorLink adds to the back of actor a LinkType type
func (t *IntransitiveActivity) AppendActorLink(v LinkType) {
	t.actor = append(t.actor, &actorIntransitiveActivityIntermediateType{Link: v})

}

// PrependActorLink adds to the front of actor a LinkType type
func (t *IntransitiveActivity) PrependActorLink(v LinkType) {
	t.actor = append([]*actorIntransitiveActivityIntermediateType{&actorIntransitiveActivityIntermediateType{Link: v}}, t.actor...)

}

// RemoveActorLink deletes the value from the specified index
func (t *IntransitiveActivity) RemoveActorLink(index int) {
	copy(t.actor[index:], t.actor[index+1:])
	t.actor[len(t.actor)-1] = nil
	t.actor = t.actor[:len(t.actor)-1]

}

// IsActorIRI determines whether the call to GetActorIRI is safe for the specified index
func (t *IntransitiveActivity) IsActorIRI(index int) (ok bool) {
	return t.actor[index].IRI != nil

}

// GetActorIRI returns the value safely if IsActorIRI returned true for the specified index
func (t *IntransitiveActivity) GetActorIRI(index int) (v *url.URL) {
	return t.actor[index].IRI

}

// AppendActorIRI adds to the back of actor a *url.URL type
func (t *IntransitiveActivity) AppendActorIRI(v *url.URL) {
	t.actor = append(t.actor, &actorIntransitiveActivityIntermediateType{IRI: v})

}

// PrependActorIRI adds to the front of actor a *url.URL type
func (t *IntransitiveActivity) PrependActorIRI(v *url.URL) {
	t.actor = append([]*actorIntransitiveActivityIntermediateType{&actorIntransitiveActivityIntermediateType{IRI: v}}, t.actor...)

}

// RemoveActorIRI deletes the value from the specified index
func (t *IntransitiveActivity) RemoveActorIRI(index int) {
	copy(t.actor[index:], t.actor[index+1:])
	t.actor[len(t.actor)-1] = nil
	t.actor = t.actor[:len(t.actor)-1]

}

// HasUnknownActor determines whether the call to GetUnknownActor is safe
func (t *IntransitiveActivity) HasUnknownActor() (ok bool) {
	return t.actor != nil && t.actor[0].unknown_ != nil

}

// GetUnknownActor returns the unknown value for actor
func (t *IntransitiveActivity) GetUnknownActor() (v interface{}) {
	return t.actor[0].unknown_

}

// SetUnknownActor sets the unknown value of actor
func (t *IntransitiveActivity) SetUnknownActor(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &actorIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.actor = append(t.actor, tmp)

}

// TargetLen determines the number of elements able to be used for the IsTargetObject, GetTargetObject, and RemoveTargetObject functions
func (t *IntransitiveActivity) TargetLen() (l int) {
	return len(t.target)

}

// IsTargetObject determines whether the call to GetTargetObject is safe for the specified index
func (t *IntransitiveActivity) IsTargetObject(index int) (ok bool) {
	return t.target[index].Object != nil

}

// GetTargetObject returns the value safely if IsTargetObject returned true for the specified index
func (t *IntransitiveActivity) GetTargetObject(index int) (v ObjectType) {
	return t.target[index].Object

}

// AppendTargetObject adds to the back of target a ObjectType type
func (t *IntransitiveActivity) AppendTargetObject(v ObjectType) {
	t.target = append(t.target, &targetIntransitiveActivityIntermediateType{Object: v})

}

// PrependTargetObject adds to the front of target a ObjectType type
func (t *IntransitiveActivity) PrependTargetObject(v ObjectType) {
	t.target = append([]*targetIntransitiveActivityIntermediateType{&targetIntransitiveActivityIntermediateType{Object: v}}, t.target...)

}

// RemoveTargetObject deletes the value from the specified index
func (t *IntransitiveActivity) RemoveTargetObject(index int) {
	copy(t.target[index:], t.target[index+1:])
	t.target[len(t.target)-1] = nil
	t.target = t.target[:len(t.target)-1]

}

// IsTargetLink determines whether the call to GetTargetLink is safe for the specified index
func (t *IntransitiveActivity) IsTargetLink(index int) (ok bool) {
	return t.target[index].Link != nil

}

// GetTargetLink returns the value safely if IsTargetLink returned true for the specified index
func (t *IntransitiveActivity) GetTargetLink(index int) (v LinkType) {
	return t.target[index].Link

}

// AppendTargetLink adds to the back of target a LinkType type
func (t *IntransitiveActivity) AppendTargetLink(v LinkType) {
	t.target = append(t.target, &targetIntransitiveActivityIntermediateType{Link: v})

}

// PrependTargetLink adds to the front of target a LinkType type
func (t *IntransitiveActivity) PrependTargetLink(v LinkType) {
	t.target = append([]*targetIntransitiveActivityIntermediateType{&targetIntransitiveActivityIntermediateType{Link: v}}, t.target...)

}

// RemoveTargetLink deletes the value from the specified index
func (t *IntransitiveActivity) RemoveTargetLink(index int) {
	copy(t.target[index:], t.target[index+1:])
	t.target[len(t.target)-1] = nil
	t.target = t.target[:len(t.target)-1]

}

// IsTargetIRI determines whether the call to GetTargetIRI is safe for the specified index
func (t *IntransitiveActivity) IsTargetIRI(index int) (ok bool) {
	return t.target[index].IRI != nil

}

// GetTargetIRI returns the value safely if IsTargetIRI returned true for the specified index
func (t *IntransitiveActivity) GetTargetIRI(index int) (v *url.URL) {
	return t.target[index].IRI

}

// AppendTargetIRI adds to the back of target a *url.URL type
func (t *IntransitiveActivity) AppendTargetIRI(v *url.URL) {
	t.target = append(t.target, &targetIntransitiveActivityIntermediateType{IRI: v})

}

// PrependTargetIRI adds to the front of target a *url.URL type
func (t *IntransitiveActivity) PrependTargetIRI(v *url.URL) {
	t.target = append([]*targetIntransitiveActivityIntermediateType{&targetIntransitiveActivityIntermediateType{IRI: v}}, t.target...)

}

// RemoveTargetIRI deletes the value from the specified index
func (t *IntransitiveActivity) RemoveTargetIRI(index int) {
	copy(t.target[index:], t.target[index+1:])
	t.target[len(t.target)-1] = nil
	t.target = t.target[:len(t.target)-1]

}

// HasUnknownTarget determines whether the call to GetUnknownTarget is safe
func (t *IntransitiveActivity) HasUnknownTarget() (ok bool) {
	return t.target != nil && t.target[0].unknown_ != nil

}

// GetUnknownTarget returns the unknown value for target
func (t *IntransitiveActivity) GetUnknownTarget() (v interface{}) {
	return t.target[0].unknown_

}

// SetUnknownTarget sets the unknown value of target
func (t *IntransitiveActivity) SetUnknownTarget(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &targetIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.target = append(t.target, tmp)

}

// ResultLen determines the number of elements able to be used for the IsResultObject, GetResultObject, and RemoveResultObject functions
func (t *IntransitiveActivity) ResultLen() (l int) {
	return len(t.result)

}

// IsResultObject determines whether the call to GetResultObject is safe for the specified index
func (t *IntransitiveActivity) IsResultObject(index int) (ok bool) {
	return t.result[index].Object != nil

}

// GetResultObject returns the value safely if IsResultObject returned true for the specified index
func (t *IntransitiveActivity) GetResultObject(index int) (v ObjectType) {
	return t.result[index].Object

}

// AppendResultObject adds to the back of result a ObjectType type
func (t *IntransitiveActivity) AppendResultObject(v ObjectType) {
	t.result = append(t.result, &resultIntransitiveActivityIntermediateType{Object: v})

}

// PrependResultObject adds to the front of result a ObjectType type
func (t *IntransitiveActivity) PrependResultObject(v ObjectType) {
	t.result = append([]*resultIntransitiveActivityIntermediateType{&resultIntransitiveActivityIntermediateType{Object: v}}, t.result...)

}

// RemoveResultObject deletes the value from the specified index
func (t *IntransitiveActivity) RemoveResultObject(index int) {
	copy(t.result[index:], t.result[index+1:])
	t.result[len(t.result)-1] = nil
	t.result = t.result[:len(t.result)-1]

}

// IsResultLink determines whether the call to GetResultLink is safe for the specified index
func (t *IntransitiveActivity) IsResultLink(index int) (ok bool) {
	return t.result[index].Link != nil

}

// GetResultLink returns the value safely if IsResultLink returned true for the specified index
func (t *IntransitiveActivity) GetResultLink(index int) (v LinkType) {
	return t.result[index].Link

}

// AppendResultLink adds to the back of result a LinkType type
func (t *IntransitiveActivity) AppendResultLink(v LinkType) {
	t.result = append(t.result, &resultIntransitiveActivityIntermediateType{Link: v})

}

// PrependResultLink adds to the front of result a LinkType type
func (t *IntransitiveActivity) PrependResultLink(v LinkType) {
	t.result = append([]*resultIntransitiveActivityIntermediateType{&resultIntransitiveActivityIntermediateType{Link: v}}, t.result...)

}

// RemoveResultLink deletes the value from the specified index
func (t *IntransitiveActivity) RemoveResultLink(index int) {
	copy(t.result[index:], t.result[index+1:])
	t.result[len(t.result)-1] = nil
	t.result = t.result[:len(t.result)-1]

}

// IsResultIRI determines whether the call to GetResultIRI is safe for the specified index
func (t *IntransitiveActivity) IsResultIRI(index int) (ok bool) {
	return t.result[index].IRI != nil

}

// GetResultIRI returns the value safely if IsResultIRI returned true for the specified index
func (t *IntransitiveActivity) GetResultIRI(index int) (v *url.URL) {
	return t.result[index].IRI

}

// AppendResultIRI adds to the back of result a *url.URL type
func (t *IntransitiveActivity) AppendResultIRI(v *url.URL) {
	t.result = append(t.result, &resultIntransitiveActivityIntermediateType{IRI: v})

}

// PrependResultIRI adds to the front of result a *url.URL type
func (t *IntransitiveActivity) PrependResultIRI(v *url.URL) {
	t.result = append([]*resultIntransitiveActivityIntermediateType{&resultIntransitiveActivityIntermediateType{IRI: v}}, t.result...)

}

// RemoveResultIRI deletes the value from the specified index
func (t *IntransitiveActivity) RemoveResultIRI(index int) {
	copy(t.result[index:], t.result[index+1:])
	t.result[len(t.result)-1] = nil
	t.result = t.result[:len(t.result)-1]

}

// HasUnknownResult determines whether the call to GetUnknownResult is safe
func (t *IntransitiveActivity) HasUnknownResult() (ok bool) {
	return t.result != nil && t.result[0].unknown_ != nil

}

// GetUnknownResult returns the unknown value for result
func (t *IntransitiveActivity) GetUnknownResult() (v interface{}) {
	return t.result[0].unknown_

}

// SetUnknownResult sets the unknown value of result
func (t *IntransitiveActivity) SetUnknownResult(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &resultIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.result = append(t.result, tmp)

}

// OriginLen determines the number of elements able to be used for the IsOriginObject, GetOriginObject, and RemoveOriginObject functions
func (t *IntransitiveActivity) OriginLen() (l int) {
	return len(t.origin)

}

// IsOriginObject determines whether the call to GetOriginObject is safe for the specified index
func (t *IntransitiveActivity) IsOriginObject(index int) (ok bool) {
	return t.origin[index].Object != nil

}

// GetOriginObject returns the value safely if IsOriginObject returned true for the specified index
func (t *IntransitiveActivity) GetOriginObject(index int) (v ObjectType) {
	return t.origin[index].Object

}

// AppendOriginObject adds to the back of origin a ObjectType type
func (t *IntransitiveActivity) AppendOriginObject(v ObjectType) {
	t.origin = append(t.origin, &originIntransitiveActivityIntermediateType{Object: v})

}

// PrependOriginObject adds to the front of origin a ObjectType type
func (t *IntransitiveActivity) PrependOriginObject(v ObjectType) {
	t.origin = append([]*originIntransitiveActivityIntermediateType{&originIntransitiveActivityIntermediateType{Object: v}}, t.origin...)

}

// RemoveOriginObject deletes the value from the specified index
func (t *IntransitiveActivity) RemoveOriginObject(index int) {
	copy(t.origin[index:], t.origin[index+1:])
	t.origin[len(t.origin)-1] = nil
	t.origin = t.origin[:len(t.origin)-1]

}

// IsOriginLink determines whether the call to GetOriginLink is safe for the specified index
func (t *IntransitiveActivity) IsOriginLink(index int) (ok bool) {
	return t.origin[index].Link != nil

}

// GetOriginLink returns the value safely if IsOriginLink returned true for the specified index
func (t *IntransitiveActivity) GetOriginLink(index int) (v LinkType) {
	return t.origin[index].Link

}

// AppendOriginLink adds to the back of origin a LinkType type
func (t *IntransitiveActivity) AppendOriginLink(v LinkType) {
	t.origin = append(t.origin, &originIntransitiveActivityIntermediateType{Link: v})

}

// PrependOriginLink adds to the front of origin a LinkType type
func (t *IntransitiveActivity) PrependOriginLink(v LinkType) {
	t.origin = append([]*originIntransitiveActivityIntermediateType{&originIntransitiveActivityIntermediateType{Link: v}}, t.origin...)

}

// RemoveOriginLink deletes the value from the specified index
func (t *IntransitiveActivity) RemoveOriginLink(index int) {
	copy(t.origin[index:], t.origin[index+1:])
	t.origin[len(t.origin)-1] = nil
	t.origin = t.origin[:len(t.origin)-1]

}

// IsOriginIRI determines whether the call to GetOriginIRI is safe for the specified index
func (t *IntransitiveActivity) IsOriginIRI(index int) (ok bool) {
	return t.origin[index].IRI != nil

}

// GetOriginIRI returns the value safely if IsOriginIRI returned true for the specified index
func (t *IntransitiveActivity) GetOriginIRI(index int) (v *url.URL) {
	return t.origin[index].IRI

}

// AppendOriginIRI adds to the back of origin a *url.URL type
func (t *IntransitiveActivity) AppendOriginIRI(v *url.URL) {
	t.origin = append(t.origin, &originIntransitiveActivityIntermediateType{IRI: v})

}

// PrependOriginIRI adds to the front of origin a *url.URL type
func (t *IntransitiveActivity) PrependOriginIRI(v *url.URL) {
	t.origin = append([]*originIntransitiveActivityIntermediateType{&originIntransitiveActivityIntermediateType{IRI: v}}, t.origin...)

}

// RemoveOriginIRI deletes the value from the specified index
func (t *IntransitiveActivity) RemoveOriginIRI(index int) {
	copy(t.origin[index:], t.origin[index+1:])
	t.origin[len(t.origin)-1] = nil
	t.origin = t.origin[:len(t.origin)-1]

}

// HasUnknownOrigin determines whether the call to GetUnknownOrigin is safe
func (t *IntransitiveActivity) HasUnknownOrigin() (ok bool) {
	return t.origin != nil && t.origin[0].unknown_ != nil

}

// GetUnknownOrigin returns the unknown value for origin
func (t *IntransitiveActivity) GetUnknownOrigin() (v interface{}) {
	return t.origin[0].unknown_

}

// SetUnknownOrigin sets the unknown value of origin
func (t *IntransitiveActivity) SetUnknownOrigin(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &originIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.origin = append(t.origin, tmp)

}

// InstrumentLen determines the number of elements able to be used for the IsInstrumentObject, GetInstrumentObject, and RemoveInstrumentObject functions
func (t *IntransitiveActivity) InstrumentLen() (l int) {
	return len(t.instrument)

}

// IsInstrumentObject determines whether the call to GetInstrumentObject is safe for the specified index
func (t *IntransitiveActivity) IsInstrumentObject(index int) (ok bool) {
	return t.instrument[index].Object != nil

}

// GetInstrumentObject returns the value safely if IsInstrumentObject returned true for the specified index
func (t *IntransitiveActivity) GetInstrumentObject(index int) (v ObjectType) {
	return t.instrument[index].Object

}

// AppendInstrumentObject adds to the back of instrument a ObjectType type
func (t *IntransitiveActivity) AppendInstrumentObject(v ObjectType) {
	t.instrument = append(t.instrument, &instrumentIntransitiveActivityIntermediateType{Object: v})

}

// PrependInstrumentObject adds to the front of instrument a ObjectType type
func (t *IntransitiveActivity) PrependInstrumentObject(v ObjectType) {
	t.instrument = append([]*instrumentIntransitiveActivityIntermediateType{&instrumentIntransitiveActivityIntermediateType{Object: v}}, t.instrument...)

}

// RemoveInstrumentObject deletes the value from the specified index
func (t *IntransitiveActivity) RemoveInstrumentObject(index int) {
	copy(t.instrument[index:], t.instrument[index+1:])
	t.instrument[len(t.instrument)-1] = nil
	t.instrument = t.instrument[:len(t.instrument)-1]

}

// IsInstrumentLink determines whether the call to GetInstrumentLink is safe for the specified index
func (t *IntransitiveActivity) IsInstrumentLink(index int) (ok bool) {
	return t.instrument[index].Link != nil

}

// GetInstrumentLink returns the value safely if IsInstrumentLink returned true for the specified index
func (t *IntransitiveActivity) GetInstrumentLink(index int) (v LinkType) {
	return t.instrument[index].Link

}

// AppendInstrumentLink adds to the back of instrument a LinkType type
func (t *IntransitiveActivity) AppendInstrumentLink(v LinkType) {
	t.instrument = append(t.instrument, &instrumentIntransitiveActivityIntermediateType{Link: v})

}

// PrependInstrumentLink adds to the front of instrument a LinkType type
func (t *IntransitiveActivity) PrependInstrumentLink(v LinkType) {
	t.instrument = append([]*instrumentIntransitiveActivityIntermediateType{&instrumentIntransitiveActivityIntermediateType{Link: v}}, t.instrument...)

}

// RemoveInstrumentLink deletes the value from the specified index
func (t *IntransitiveActivity) RemoveInstrumentLink(index int) {
	copy(t.instrument[index:], t.instrument[index+1:])
	t.instrument[len(t.instrument)-1] = nil
	t.instrument = t.instrument[:len(t.instrument)-1]

}

// IsInstrumentIRI determines whether the call to GetInstrumentIRI is safe for the specified index
func (t *IntransitiveActivity) IsInstrumentIRI(index int) (ok bool) {
	return t.instrument[index].IRI != nil

}

// GetInstrumentIRI returns the value safely if IsInstrumentIRI returned true for the specified index
func (t *IntransitiveActivity) GetInstrumentIRI(index int) (v *url.URL) {
	return t.instrument[index].IRI

}

// AppendInstrumentIRI adds to the back of instrument a *url.URL type
func (t *IntransitiveActivity) AppendInstrumentIRI(v *url.URL) {
	t.instrument = append(t.instrument, &instrumentIntransitiveActivityIntermediateType{IRI: v})

}

// PrependInstrumentIRI adds to the front of instrument a *url.URL type
func (t *IntransitiveActivity) PrependInstrumentIRI(v *url.URL) {
	t.instrument = append([]*instrumentIntransitiveActivityIntermediateType{&instrumentIntransitiveActivityIntermediateType{IRI: v}}, t.instrument...)

}

// RemoveInstrumentIRI deletes the value from the specified index
func (t *IntransitiveActivity) RemoveInstrumentIRI(index int) {
	copy(t.instrument[index:], t.instrument[index+1:])
	t.instrument[len(t.instrument)-1] = nil
	t.instrument = t.instrument[:len(t.instrument)-1]

}

// HasUnknownInstrument determines whether the call to GetUnknownInstrument is safe
func (t *IntransitiveActivity) HasUnknownInstrument() (ok bool) {
	return t.instrument != nil && t.instrument[0].unknown_ != nil

}

// GetUnknownInstrument returns the unknown value for instrument
func (t *IntransitiveActivity) GetUnknownInstrument() (v interface{}) {
	return t.instrument[0].unknown_

}

// SetUnknownInstrument sets the unknown value of instrument
func (t *IntransitiveActivity) SetUnknownInstrument(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &instrumentIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.instrument = append(t.instrument, tmp)

}

// IsAltitude determines whether the call to GetAltitude is safe
func (t *IntransitiveActivity) IsAltitude() (ok bool) {
	return t.altitude != nil && t.altitude.float != nil

}

// GetAltitude returns the value safely if IsAltitude returned true
func (t *IntransitiveActivity) GetAltitude() (v float64) {
	return *t.altitude.float

}

// SetAltitude sets the value of altitude to be of float64 type
func (t *IntransitiveActivity) SetAltitude(v float64) {
	t.altitude = &altitudeIntransitiveActivityIntermediateType{float: &v}

}

// IsAltitudeIRI determines whether the call to GetAltitudeIRI is safe
func (t *IntransitiveActivity) IsAltitudeIRI() (ok bool) {
	return t.altitude != nil && t.altitude.IRI != nil

}

// GetAltitudeIRI returns the value safely if IsAltitudeIRI returned true
func (t *IntransitiveActivity) GetAltitudeIRI() (v *url.URL) {
	return t.altitude.IRI

}

// SetAltitudeIRI sets the value of altitude to be of *url.URL type
func (t *IntransitiveActivity) SetAltitudeIRI(v *url.URL) {
	t.altitude = &altitudeIntransitiveActivityIntermediateType{IRI: v}

}

// HasUnknownAltitude determines whether the call to GetUnknownAltitude is safe
func (t *IntransitiveActivity) HasUnknownAltitude() (ok bool) {
	return t.altitude != nil && t.altitude.unknown_ != nil

}

// GetUnknownAltitude returns the unknown value for altitude
func (t *IntransitiveActivity) GetUnknownAltitude() (v interface{}) {
	return t.altitude.unknown_

}

// SetUnknownAltitude sets the unknown value of altitude
func (t *IntransitiveActivity) SetUnknownAltitude(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &altitudeIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.altitude = tmp

}

// AttachmentLen determines the number of elements able to be used for the IsAttachmentObject, GetAttachmentObject, and RemoveAttachmentObject functions
func (t *IntransitiveActivity) AttachmentLen() (l int) {
	return len(t.attachment)

}

// IsAttachmentObject determines whether the call to GetAttachmentObject is safe for the specified index
func (t *IntransitiveActivity) IsAttachmentObject(index int) (ok bool) {
	return t.attachment[index].Object != nil

}

// GetAttachmentObject returns the value safely if IsAttachmentObject returned true for the specified index
func (t *IntransitiveActivity) GetAttachmentObject(index int) (v ObjectType) {
	return t.attachment[index].Object

}

// AppendAttachmentObject adds to the back of attachment a ObjectType type
func (t *IntransitiveActivity) AppendAttachmentObject(v ObjectType) {
	t.attachment = append(t.attachment, &attachmentIntransitiveActivityIntermediateType{Object: v})

}

// PrependAttachmentObject adds to the front of attachment a ObjectType type
func (t *IntransitiveActivity) PrependAttachmentObject(v ObjectType) {
	t.attachment = append([]*attachmentIntransitiveActivityIntermediateType{&attachmentIntransitiveActivityIntermediateType{Object: v}}, t.attachment...)

}

// RemoveAttachmentObject deletes the value from the specified index
func (t *IntransitiveActivity) RemoveAttachmentObject(index int) {
	copy(t.attachment[index:], t.attachment[index+1:])
	t.attachment[len(t.attachment)-1] = nil
	t.attachment = t.attachment[:len(t.attachment)-1]

}

// IsAttachmentLink determines whether the call to GetAttachmentLink is safe for the specified index
func (t *IntransitiveActivity) IsAttachmentLink(index int) (ok bool) {
	return t.attachment[index].Link != nil

}

// GetAttachmentLink returns the value safely if IsAttachmentLink returned true for the specified index
func (t *IntransitiveActivity) GetAttachmentLink(index int) (v LinkType) {
	return t.attachment[index].Link

}

// AppendAttachmentLink adds to the back of attachment a LinkType type
func (t *IntransitiveActivity) AppendAttachmentLink(v LinkType) {
	t.attachment = append(t.attachment, &attachmentIntransitiveActivityIntermediateType{Link: v})

}

// PrependAttachmentLink adds to the front of attachment a LinkType type
func (t *IntransitiveActivity) PrependAttachmentLink(v LinkType) {
	t.attachment = append([]*attachmentIntransitiveActivityIntermediateType{&attachmentIntransitiveActivityIntermediateType{Link: v}}, t.attachment...)

}

// RemoveAttachmentLink deletes the value from the specified index
func (t *IntransitiveActivity) RemoveAttachmentLink(index int) {
	copy(t.attachment[index:], t.attachment[index+1:])
	t.attachment[len(t.attachment)-1] = nil
	t.attachment = t.attachment[:len(t.attachment)-1]

}

// IsAttachmentIRI determines whether the call to GetAttachmentIRI is safe for the specified index
func (t *IntransitiveActivity) IsAttachmentIRI(index int) (ok bool) {
	return t.attachment[index].IRI != nil

}

// GetAttachmentIRI returns the value safely if IsAttachmentIRI returned true for the specified index
func (t *IntransitiveActivity) GetAttachmentIRI(index int) (v *url.URL) {
	return t.attachment[index].IRI

}

// AppendAttachmentIRI adds to the back of attachment a *url.URL type
func (t *IntransitiveActivity) AppendAttachmentIRI(v *url.URL) {
	t.attachment = append(t.attachment, &attachmentIntransitiveActivityIntermediateType{IRI: v})

}

// PrependAttachmentIRI adds to the front of attachment a *url.URL type
func (t *IntransitiveActivity) PrependAttachmentIRI(v *url.URL) {
	t.attachment = append([]*attachmentIntransitiveActivityIntermediateType{&attachmentIntransitiveActivityIntermediateType{IRI: v}}, t.attachment...)

}

// RemoveAttachmentIRI deletes the value from the specified index
func (t *IntransitiveActivity) RemoveAttachmentIRI(index int) {
	copy(t.attachment[index:], t.attachment[index+1:])
	t.attachment[len(t.attachment)-1] = nil
	t.attachment = t.attachment[:len(t.attachment)-1]

}

// HasUnknownAttachment determines whether the call to GetUnknownAttachment is safe
func (t *IntransitiveActivity) HasUnknownAttachment() (ok bool) {
	return t.attachment != nil && t.attachment[0].unknown_ != nil

}

// GetUnknownAttachment returns the unknown value for attachment
func (t *IntransitiveActivity) GetUnknownAttachment() (v interface{}) {
	return t.attachment[0].unknown_

}

// SetUnknownAttachment sets the unknown value of attachment
func (t *IntransitiveActivity) SetUnknownAttachment(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &attachmentIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.attachment = append(t.attachment, tmp)

}

// AttributedToLen determines the number of elements able to be used for the IsAttributedToObject, GetAttributedToObject, and RemoveAttributedToObject functions
func (t *IntransitiveActivity) AttributedToLen() (l int) {
	return len(t.attributedTo)

}

// IsAttributedToObject determines whether the call to GetAttributedToObject is safe for the specified index
func (t *IntransitiveActivity) IsAttributedToObject(index int) (ok bool) {
	return t.attributedTo[index].Object != nil

}

// GetAttributedToObject returns the value safely if IsAttributedToObject returned true for the specified index
func (t *IntransitiveActivity) GetAttributedToObject(index int) (v ObjectType) {
	return t.attributedTo[index].Object

}

// AppendAttributedToObject adds to the back of attributedTo a ObjectType type
func (t *IntransitiveActivity) AppendAttributedToObject(v ObjectType) {
	t.attributedTo = append(t.attributedTo, &attributedToIntransitiveActivityIntermediateType{Object: v})

}

// PrependAttributedToObject adds to the front of attributedTo a ObjectType type
func (t *IntransitiveActivity) PrependAttributedToObject(v ObjectType) {
	t.attributedTo = append([]*attributedToIntransitiveActivityIntermediateType{&attributedToIntransitiveActivityIntermediateType{Object: v}}, t.attributedTo...)

}

// RemoveAttributedToObject deletes the value from the specified index
func (t *IntransitiveActivity) RemoveAttributedToObject(index int) {
	copy(t.attributedTo[index:], t.attributedTo[index+1:])
	t.attributedTo[len(t.attributedTo)-1] = nil
	t.attributedTo = t.attributedTo[:len(t.attributedTo)-1]

}

// IsAttributedToLink determines whether the call to GetAttributedToLink is safe for the specified index
func (t *IntransitiveActivity) IsAttributedToLink(index int) (ok bool) {
	return t.attributedTo[index].Link != nil

}

// GetAttributedToLink returns the value safely if IsAttributedToLink returned true for the specified index
func (t *IntransitiveActivity) GetAttributedToLink(index int) (v LinkType) {
	return t.attributedTo[index].Link

}

// AppendAttributedToLink adds to the back of attributedTo a LinkType type
func (t *IntransitiveActivity) AppendAttributedToLink(v LinkType) {
	t.attributedTo = append(t.attributedTo, &attributedToIntransitiveActivityIntermediateType{Link: v})

}

// PrependAttributedToLink adds to the front of attributedTo a LinkType type
func (t *IntransitiveActivity) PrependAttributedToLink(v LinkType) {
	t.attributedTo = append([]*attributedToIntransitiveActivityIntermediateType{&attributedToIntransitiveActivityIntermediateType{Link: v}}, t.attributedTo...)

}

// RemoveAttributedToLink deletes the value from the specified index
func (t *IntransitiveActivity) RemoveAttributedToLink(index int) {
	copy(t.attributedTo[index:], t.attributedTo[index+1:])
	t.attributedTo[len(t.attributedTo)-1] = nil
	t.attributedTo = t.attributedTo[:len(t.attributedTo)-1]

}

// IsAttributedToIRI determines whether the call to GetAttributedToIRI is safe for the specified index
func (t *IntransitiveActivity) IsAttributedToIRI(index int) (ok bool) {
	return t.attributedTo[index].IRI != nil

}

// GetAttributedToIRI returns the value safely if IsAttributedToIRI returned true for the specified index
func (t *IntransitiveActivity) GetAttributedToIRI(index int) (v *url.URL) {
	return t.attributedTo[index].IRI

}

// AppendAttributedToIRI adds to the back of attributedTo a *url.URL type
func (t *IntransitiveActivity) AppendAttributedToIRI(v *url.URL) {
	t.attributedTo = append(t.attributedTo, &attributedToIntransitiveActivityIntermediateType{IRI: v})

}

// PrependAttributedToIRI adds to the front of attributedTo a *url.URL type
func (t *IntransitiveActivity) PrependAttributedToIRI(v *url.URL) {
	t.attributedTo = append([]*attributedToIntransitiveActivityIntermediateType{&attributedToIntransitiveActivityIntermediateType{IRI: v}}, t.attributedTo...)

}

// RemoveAttributedToIRI deletes the value from the specified index
func (t *IntransitiveActivity) RemoveAttributedToIRI(index int) {
	copy(t.attributedTo[index:], t.attributedTo[index+1:])
	t.attributedTo[len(t.attributedTo)-1] = nil
	t.attributedTo = t.attributedTo[:len(t.attributedTo)-1]

}

// HasUnknownAttributedTo determines whether the call to GetUnknownAttributedTo is safe
func (t *IntransitiveActivity) HasUnknownAttributedTo() (ok bool) {
	return t.attributedTo != nil && t.attributedTo[0].unknown_ != nil

}

// GetUnknownAttributedTo returns the unknown value for attributedTo
func (t *IntransitiveActivity) GetUnknownAttributedTo() (v interface{}) {
	return t.attributedTo[0].unknown_

}

// SetUnknownAttributedTo sets the unknown value of attributedTo
func (t *IntransitiveActivity) SetUnknownAttributedTo(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &attributedToIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.attributedTo = append(t.attributedTo, tmp)

}

// AudienceLen determines the number of elements able to be used for the IsAudienceObject, GetAudienceObject, and RemoveAudienceObject functions
func (t *IntransitiveActivity) AudienceLen() (l int) {
	return len(t.audience)

}

// IsAudienceObject determines whether the call to GetAudienceObject is safe for the specified index
func (t *IntransitiveActivity) IsAudienceObject(index int) (ok bool) {
	return t.audience[index].Object != nil

}

// GetAudienceObject returns the value safely if IsAudienceObject returned true for the specified index
func (t *IntransitiveActivity) GetAudienceObject(index int) (v ObjectType) {
	return t.audience[index].Object

}

// AppendAudienceObject adds to the back of audience a ObjectType type
func (t *IntransitiveActivity) AppendAudienceObject(v ObjectType) {
	t.audience = append(t.audience, &audienceIntransitiveActivityIntermediateType{Object: v})

}

// PrependAudienceObject adds to the front of audience a ObjectType type
func (t *IntransitiveActivity) PrependAudienceObject(v ObjectType) {
	t.audience = append([]*audienceIntransitiveActivityIntermediateType{&audienceIntransitiveActivityIntermediateType{Object: v}}, t.audience...)

}

// RemoveAudienceObject deletes the value from the specified index
func (t *IntransitiveActivity) RemoveAudienceObject(index int) {
	copy(t.audience[index:], t.audience[index+1:])
	t.audience[len(t.audience)-1] = nil
	t.audience = t.audience[:len(t.audience)-1]

}

// IsAudienceLink determines whether the call to GetAudienceLink is safe for the specified index
func (t *IntransitiveActivity) IsAudienceLink(index int) (ok bool) {
	return t.audience[index].Link != nil

}

// GetAudienceLink returns the value safely if IsAudienceLink returned true for the specified index
func (t *IntransitiveActivity) GetAudienceLink(index int) (v LinkType) {
	return t.audience[index].Link

}

// AppendAudienceLink adds to the back of audience a LinkType type
func (t *IntransitiveActivity) AppendAudienceLink(v LinkType) {
	t.audience = append(t.audience, &audienceIntransitiveActivityIntermediateType{Link: v})

}

// PrependAudienceLink adds to the front of audience a LinkType type
func (t *IntransitiveActivity) PrependAudienceLink(v LinkType) {
	t.audience = append([]*audienceIntransitiveActivityIntermediateType{&audienceIntransitiveActivityIntermediateType{Link: v}}, t.audience...)

}

// RemoveAudienceLink deletes the value from the specified index
func (t *IntransitiveActivity) RemoveAudienceLink(index int) {
	copy(t.audience[index:], t.audience[index+1:])
	t.audience[len(t.audience)-1] = nil
	t.audience = t.audience[:len(t.audience)-1]

}

// IsAudienceIRI determines whether the call to GetAudienceIRI is safe for the specified index
func (t *IntransitiveActivity) IsAudienceIRI(index int) (ok bool) {
	return t.audience[index].IRI != nil

}

// GetAudienceIRI returns the value safely if IsAudienceIRI returned true for the specified index
func (t *IntransitiveActivity) GetAudienceIRI(index int) (v *url.URL) {
	return t.audience[index].IRI

}

// AppendAudienceIRI adds to the back of audience a *url.URL type
func (t *IntransitiveActivity) AppendAudienceIRI(v *url.URL) {
	t.audience = append(t.audience, &audienceIntransitiveActivityIntermediateType{IRI: v})

}

// PrependAudienceIRI adds to the front of audience a *url.URL type
func (t *IntransitiveActivity) PrependAudienceIRI(v *url.URL) {
	t.audience = append([]*audienceIntransitiveActivityIntermediateType{&audienceIntransitiveActivityIntermediateType{IRI: v}}, t.audience...)

}

// RemoveAudienceIRI deletes the value from the specified index
func (t *IntransitiveActivity) RemoveAudienceIRI(index int) {
	copy(t.audience[index:], t.audience[index+1:])
	t.audience[len(t.audience)-1] = nil
	t.audience = t.audience[:len(t.audience)-1]

}

// HasUnknownAudience determines whether the call to GetUnknownAudience is safe
func (t *IntransitiveActivity) HasUnknownAudience() (ok bool) {
	return t.audience != nil && t.audience[0].unknown_ != nil

}

// GetUnknownAudience returns the unknown value for audience
func (t *IntransitiveActivity) GetUnknownAudience() (v interface{}) {
	return t.audience[0].unknown_

}

// SetUnknownAudience sets the unknown value of audience
func (t *IntransitiveActivity) SetUnknownAudience(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &audienceIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.audience = append(t.audience, tmp)

}

// ContentLen determines the number of elements able to be used for the IsContentString, GetContentString, and RemoveContentString functions
func (t *IntransitiveActivity) ContentLen() (l int) {
	return len(t.content)

}

// IsContentString determines whether the call to GetContentString is safe for the specified index
func (t *IntransitiveActivity) IsContentString(index int) (ok bool) {
	return t.content[index].stringName != nil

}

// GetContentString returns the value safely if IsContentString returned true for the specified index
func (t *IntransitiveActivity) GetContentString(index int) (v string) {
	return *t.content[index].stringName

}

// AppendContentString adds to the back of content a string type
func (t *IntransitiveActivity) AppendContentString(v string) {
	t.content = append(t.content, &contentIntransitiveActivityIntermediateType{stringName: &v})

}

// PrependContentString adds to the front of content a string type
func (t *IntransitiveActivity) PrependContentString(v string) {
	t.content = append([]*contentIntransitiveActivityIntermediateType{&contentIntransitiveActivityIntermediateType{stringName: &v}}, t.content...)

}

// RemoveContentString deletes the value from the specified index
func (t *IntransitiveActivity) RemoveContentString(index int) {
	copy(t.content[index:], t.content[index+1:])
	t.content[len(t.content)-1] = nil
	t.content = t.content[:len(t.content)-1]

}

// IsContentLangString determines whether the call to GetContentLangString is safe for the specified index
func (t *IntransitiveActivity) IsContentLangString(index int) (ok bool) {
	return t.content[index].langString != nil

}

// GetContentLangString returns the value safely if IsContentLangString returned true for the specified index
func (t *IntransitiveActivity) GetContentLangString(index int) (v string) {
	return *t.content[index].langString

}

// AppendContentLangString adds to the back of content a string type
func (t *IntransitiveActivity) AppendContentLangString(v string) {
	t.content = append(t.content, &contentIntransitiveActivityIntermediateType{langString: &v})

}

// PrependContentLangString adds to the front of content a string type
func (t *IntransitiveActivity) PrependContentLangString(v string) {
	t.content = append([]*contentIntransitiveActivityIntermediateType{&contentIntransitiveActivityIntermediateType{langString: &v}}, t.content...)

}

// RemoveContentLangString deletes the value from the specified index
func (t *IntransitiveActivity) RemoveContentLangString(index int) {
	copy(t.content[index:], t.content[index+1:])
	t.content[len(t.content)-1] = nil
	t.content = t.content[:len(t.content)-1]

}

// IsContentIRI determines whether the call to GetContentIRI is safe for the specified index
func (t *IntransitiveActivity) IsContentIRI(index int) (ok bool) {
	return t.content[index].IRI != nil

}

// GetContentIRI returns the value safely if IsContentIRI returned true for the specified index
func (t *IntransitiveActivity) GetContentIRI(index int) (v *url.URL) {
	return t.content[index].IRI

}

// AppendContentIRI adds to the back of content a *url.URL type
func (t *IntransitiveActivity) AppendContentIRI(v *url.URL) {
	t.content = append(t.content, &contentIntransitiveActivityIntermediateType{IRI: v})

}

// PrependContentIRI adds to the front of content a *url.URL type
func (t *IntransitiveActivity) PrependContentIRI(v *url.URL) {
	t.content = append([]*contentIntransitiveActivityIntermediateType{&contentIntransitiveActivityIntermediateType{IRI: v}}, t.content...)

}

// RemoveContentIRI deletes the value from the specified index
func (t *IntransitiveActivity) RemoveContentIRI(index int) {
	copy(t.content[index:], t.content[index+1:])
	t.content[len(t.content)-1] = nil
	t.content = t.content[:len(t.content)-1]

}

// HasUnknownContent determines whether the call to GetUnknownContent is safe
func (t *IntransitiveActivity) HasUnknownContent() (ok bool) {
	return t.content != nil && t.content[0].unknown_ != nil

}

// GetUnknownContent returns the unknown value for content
func (t *IntransitiveActivity) GetUnknownContent() (v interface{}) {
	return t.content[0].unknown_

}

// SetUnknownContent sets the unknown value of content
func (t *IntransitiveActivity) SetUnknownContent(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &contentIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.content = append(t.content, tmp)

}

// ContentMapLanguages returns all languages for this property's language mapping, or nil if there are none.
func (t *IntransitiveActivity) ContentMapLanguages() (l []string) {
	if t.contentMap == nil || len(t.contentMap) == 0 {
		return nil
	}
	for k := range t.contentMap {
		l = append(l, k)
	}
	return

}

// GetContentMap retrieves the value of the property for the specified language, or an empty string if it does not exist
func (t *IntransitiveActivity) GetContentMap(l string) (v string) {
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
func (t *IntransitiveActivity) SetContentMap(l string, v string) {
	if t.contentMap == nil {
		t.contentMap = make(map[string]string)
	}
	t.contentMap[l] = v

}

// ContextLen determines the number of elements able to be used for the IsContextObject, GetContextObject, and RemoveContextObject functions
func (t *IntransitiveActivity) ContextLen() (l int) {
	return len(t.context)

}

// IsContextObject determines whether the call to GetContextObject is safe for the specified index
func (t *IntransitiveActivity) IsContextObject(index int) (ok bool) {
	return t.context[index].Object != nil

}

// GetContextObject returns the value safely if IsContextObject returned true for the specified index
func (t *IntransitiveActivity) GetContextObject(index int) (v ObjectType) {
	return t.context[index].Object

}

// AppendContextObject adds to the back of context a ObjectType type
func (t *IntransitiveActivity) AppendContextObject(v ObjectType) {
	t.context = append(t.context, &contextIntransitiveActivityIntermediateType{Object: v})

}

// PrependContextObject adds to the front of context a ObjectType type
func (t *IntransitiveActivity) PrependContextObject(v ObjectType) {
	t.context = append([]*contextIntransitiveActivityIntermediateType{&contextIntransitiveActivityIntermediateType{Object: v}}, t.context...)

}

// RemoveContextObject deletes the value from the specified index
func (t *IntransitiveActivity) RemoveContextObject(index int) {
	copy(t.context[index:], t.context[index+1:])
	t.context[len(t.context)-1] = nil
	t.context = t.context[:len(t.context)-1]

}

// IsContextLink determines whether the call to GetContextLink is safe for the specified index
func (t *IntransitiveActivity) IsContextLink(index int) (ok bool) {
	return t.context[index].Link != nil

}

// GetContextLink returns the value safely if IsContextLink returned true for the specified index
func (t *IntransitiveActivity) GetContextLink(index int) (v LinkType) {
	return t.context[index].Link

}

// AppendContextLink adds to the back of context a LinkType type
func (t *IntransitiveActivity) AppendContextLink(v LinkType) {
	t.context = append(t.context, &contextIntransitiveActivityIntermediateType{Link: v})

}

// PrependContextLink adds to the front of context a LinkType type
func (t *IntransitiveActivity) PrependContextLink(v LinkType) {
	t.context = append([]*contextIntransitiveActivityIntermediateType{&contextIntransitiveActivityIntermediateType{Link: v}}, t.context...)

}

// RemoveContextLink deletes the value from the specified index
func (t *IntransitiveActivity) RemoveContextLink(index int) {
	copy(t.context[index:], t.context[index+1:])
	t.context[len(t.context)-1] = nil
	t.context = t.context[:len(t.context)-1]

}

// IsContextIRI determines whether the call to GetContextIRI is safe for the specified index
func (t *IntransitiveActivity) IsContextIRI(index int) (ok bool) {
	return t.context[index].IRI != nil

}

// GetContextIRI returns the value safely if IsContextIRI returned true for the specified index
func (t *IntransitiveActivity) GetContextIRI(index int) (v *url.URL) {
	return t.context[index].IRI

}

// AppendContextIRI adds to the back of context a *url.URL type
func (t *IntransitiveActivity) AppendContextIRI(v *url.URL) {
	t.context = append(t.context, &contextIntransitiveActivityIntermediateType{IRI: v})

}

// PrependContextIRI adds to the front of context a *url.URL type
func (t *IntransitiveActivity) PrependContextIRI(v *url.URL) {
	t.context = append([]*contextIntransitiveActivityIntermediateType{&contextIntransitiveActivityIntermediateType{IRI: v}}, t.context...)

}

// RemoveContextIRI deletes the value from the specified index
func (t *IntransitiveActivity) RemoveContextIRI(index int) {
	copy(t.context[index:], t.context[index+1:])
	t.context[len(t.context)-1] = nil
	t.context = t.context[:len(t.context)-1]

}

// HasUnknownContext determines whether the call to GetUnknownContext is safe
func (t *IntransitiveActivity) HasUnknownContext() (ok bool) {
	return t.context != nil && t.context[0].unknown_ != nil

}

// GetUnknownContext returns the unknown value for context
func (t *IntransitiveActivity) GetUnknownContext() (v interface{}) {
	return t.context[0].unknown_

}

// SetUnknownContext sets the unknown value of context
func (t *IntransitiveActivity) SetUnknownContext(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &contextIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.context = append(t.context, tmp)

}

// NameLen determines the number of elements able to be used for the IsNameString, GetNameString, and RemoveNameString functions
func (t *IntransitiveActivity) NameLen() (l int) {
	return len(t.name)

}

// IsNameString determines whether the call to GetNameString is safe for the specified index
func (t *IntransitiveActivity) IsNameString(index int) (ok bool) {
	return t.name[index].stringName != nil

}

// GetNameString returns the value safely if IsNameString returned true for the specified index
func (t *IntransitiveActivity) GetNameString(index int) (v string) {
	return *t.name[index].stringName

}

// AppendNameString adds to the back of name a string type
func (t *IntransitiveActivity) AppendNameString(v string) {
	t.name = append(t.name, &nameIntransitiveActivityIntermediateType{stringName: &v})

}

// PrependNameString adds to the front of name a string type
func (t *IntransitiveActivity) PrependNameString(v string) {
	t.name = append([]*nameIntransitiveActivityIntermediateType{&nameIntransitiveActivityIntermediateType{stringName: &v}}, t.name...)

}

// RemoveNameString deletes the value from the specified index
func (t *IntransitiveActivity) RemoveNameString(index int) {
	copy(t.name[index:], t.name[index+1:])
	t.name[len(t.name)-1] = nil
	t.name = t.name[:len(t.name)-1]

}

// IsNameLangString determines whether the call to GetNameLangString is safe for the specified index
func (t *IntransitiveActivity) IsNameLangString(index int) (ok bool) {
	return t.name[index].langString != nil

}

// GetNameLangString returns the value safely if IsNameLangString returned true for the specified index
func (t *IntransitiveActivity) GetNameLangString(index int) (v string) {
	return *t.name[index].langString

}

// AppendNameLangString adds to the back of name a string type
func (t *IntransitiveActivity) AppendNameLangString(v string) {
	t.name = append(t.name, &nameIntransitiveActivityIntermediateType{langString: &v})

}

// PrependNameLangString adds to the front of name a string type
func (t *IntransitiveActivity) PrependNameLangString(v string) {
	t.name = append([]*nameIntransitiveActivityIntermediateType{&nameIntransitiveActivityIntermediateType{langString: &v}}, t.name...)

}

// RemoveNameLangString deletes the value from the specified index
func (t *IntransitiveActivity) RemoveNameLangString(index int) {
	copy(t.name[index:], t.name[index+1:])
	t.name[len(t.name)-1] = nil
	t.name = t.name[:len(t.name)-1]

}

// IsNameIRI determines whether the call to GetNameIRI is safe for the specified index
func (t *IntransitiveActivity) IsNameIRI(index int) (ok bool) {
	return t.name[index].IRI != nil

}

// GetNameIRI returns the value safely if IsNameIRI returned true for the specified index
func (t *IntransitiveActivity) GetNameIRI(index int) (v *url.URL) {
	return t.name[index].IRI

}

// AppendNameIRI adds to the back of name a *url.URL type
func (t *IntransitiveActivity) AppendNameIRI(v *url.URL) {
	t.name = append(t.name, &nameIntransitiveActivityIntermediateType{IRI: v})

}

// PrependNameIRI adds to the front of name a *url.URL type
func (t *IntransitiveActivity) PrependNameIRI(v *url.URL) {
	t.name = append([]*nameIntransitiveActivityIntermediateType{&nameIntransitiveActivityIntermediateType{IRI: v}}, t.name...)

}

// RemoveNameIRI deletes the value from the specified index
func (t *IntransitiveActivity) RemoveNameIRI(index int) {
	copy(t.name[index:], t.name[index+1:])
	t.name[len(t.name)-1] = nil
	t.name = t.name[:len(t.name)-1]

}

// HasUnknownName determines whether the call to GetUnknownName is safe
func (t *IntransitiveActivity) HasUnknownName() (ok bool) {
	return t.name != nil && t.name[0].unknown_ != nil

}

// GetUnknownName returns the unknown value for name
func (t *IntransitiveActivity) GetUnknownName() (v interface{}) {
	return t.name[0].unknown_

}

// SetUnknownName sets the unknown value of name
func (t *IntransitiveActivity) SetUnknownName(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &nameIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.name = append(t.name, tmp)

}

// NameMapLanguages returns all languages for this property's language mapping, or nil if there are none.
func (t *IntransitiveActivity) NameMapLanguages() (l []string) {
	if t.nameMap == nil || len(t.nameMap) == 0 {
		return nil
	}
	for k := range t.nameMap {
		l = append(l, k)
	}
	return

}

// GetNameMap retrieves the value of the property for the specified language, or an empty string if it does not exist
func (t *IntransitiveActivity) GetNameMap(l string) (v string) {
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
func (t *IntransitiveActivity) SetNameMap(l string, v string) {
	if t.nameMap == nil {
		t.nameMap = make(map[string]string)
	}
	t.nameMap[l] = v

}

// IsEndTime determines whether the call to GetEndTime is safe
func (t *IntransitiveActivity) IsEndTime() (ok bool) {
	return t.endTime != nil && t.endTime.dateTime != nil

}

// GetEndTime returns the value safely if IsEndTime returned true
func (t *IntransitiveActivity) GetEndTime() (v time.Time) {
	return *t.endTime.dateTime

}

// SetEndTime sets the value of endTime to be of time.Time type
func (t *IntransitiveActivity) SetEndTime(v time.Time) {
	t.endTime = &endTimeIntransitiveActivityIntermediateType{dateTime: &v}

}

// IsEndTimeIRI determines whether the call to GetEndTimeIRI is safe
func (t *IntransitiveActivity) IsEndTimeIRI() (ok bool) {
	return t.endTime != nil && t.endTime.IRI != nil

}

// GetEndTimeIRI returns the value safely if IsEndTimeIRI returned true
func (t *IntransitiveActivity) GetEndTimeIRI() (v *url.URL) {
	return t.endTime.IRI

}

// SetEndTimeIRI sets the value of endTime to be of *url.URL type
func (t *IntransitiveActivity) SetEndTimeIRI(v *url.URL) {
	t.endTime = &endTimeIntransitiveActivityIntermediateType{IRI: v}

}

// HasUnknownEndTime determines whether the call to GetUnknownEndTime is safe
func (t *IntransitiveActivity) HasUnknownEndTime() (ok bool) {
	return t.endTime != nil && t.endTime.unknown_ != nil

}

// GetUnknownEndTime returns the unknown value for endTime
func (t *IntransitiveActivity) GetUnknownEndTime() (v interface{}) {
	return t.endTime.unknown_

}

// SetUnknownEndTime sets the unknown value of endTime
func (t *IntransitiveActivity) SetUnknownEndTime(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &endTimeIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.endTime = tmp

}

// GeneratorLen determines the number of elements able to be used for the IsGeneratorObject, GetGeneratorObject, and RemoveGeneratorObject functions
func (t *IntransitiveActivity) GeneratorLen() (l int) {
	return len(t.generator)

}

// IsGeneratorObject determines whether the call to GetGeneratorObject is safe for the specified index
func (t *IntransitiveActivity) IsGeneratorObject(index int) (ok bool) {
	return t.generator[index].Object != nil

}

// GetGeneratorObject returns the value safely if IsGeneratorObject returned true for the specified index
func (t *IntransitiveActivity) GetGeneratorObject(index int) (v ObjectType) {
	return t.generator[index].Object

}

// AppendGeneratorObject adds to the back of generator a ObjectType type
func (t *IntransitiveActivity) AppendGeneratorObject(v ObjectType) {
	t.generator = append(t.generator, &generatorIntransitiveActivityIntermediateType{Object: v})

}

// PrependGeneratorObject adds to the front of generator a ObjectType type
func (t *IntransitiveActivity) PrependGeneratorObject(v ObjectType) {
	t.generator = append([]*generatorIntransitiveActivityIntermediateType{&generatorIntransitiveActivityIntermediateType{Object: v}}, t.generator...)

}

// RemoveGeneratorObject deletes the value from the specified index
func (t *IntransitiveActivity) RemoveGeneratorObject(index int) {
	copy(t.generator[index:], t.generator[index+1:])
	t.generator[len(t.generator)-1] = nil
	t.generator = t.generator[:len(t.generator)-1]

}

// IsGeneratorLink determines whether the call to GetGeneratorLink is safe for the specified index
func (t *IntransitiveActivity) IsGeneratorLink(index int) (ok bool) {
	return t.generator[index].Link != nil

}

// GetGeneratorLink returns the value safely if IsGeneratorLink returned true for the specified index
func (t *IntransitiveActivity) GetGeneratorLink(index int) (v LinkType) {
	return t.generator[index].Link

}

// AppendGeneratorLink adds to the back of generator a LinkType type
func (t *IntransitiveActivity) AppendGeneratorLink(v LinkType) {
	t.generator = append(t.generator, &generatorIntransitiveActivityIntermediateType{Link: v})

}

// PrependGeneratorLink adds to the front of generator a LinkType type
func (t *IntransitiveActivity) PrependGeneratorLink(v LinkType) {
	t.generator = append([]*generatorIntransitiveActivityIntermediateType{&generatorIntransitiveActivityIntermediateType{Link: v}}, t.generator...)

}

// RemoveGeneratorLink deletes the value from the specified index
func (t *IntransitiveActivity) RemoveGeneratorLink(index int) {
	copy(t.generator[index:], t.generator[index+1:])
	t.generator[len(t.generator)-1] = nil
	t.generator = t.generator[:len(t.generator)-1]

}

// IsGeneratorIRI determines whether the call to GetGeneratorIRI is safe for the specified index
func (t *IntransitiveActivity) IsGeneratorIRI(index int) (ok bool) {
	return t.generator[index].IRI != nil

}

// GetGeneratorIRI returns the value safely if IsGeneratorIRI returned true for the specified index
func (t *IntransitiveActivity) GetGeneratorIRI(index int) (v *url.URL) {
	return t.generator[index].IRI

}

// AppendGeneratorIRI adds to the back of generator a *url.URL type
func (t *IntransitiveActivity) AppendGeneratorIRI(v *url.URL) {
	t.generator = append(t.generator, &generatorIntransitiveActivityIntermediateType{IRI: v})

}

// PrependGeneratorIRI adds to the front of generator a *url.URL type
func (t *IntransitiveActivity) PrependGeneratorIRI(v *url.URL) {
	t.generator = append([]*generatorIntransitiveActivityIntermediateType{&generatorIntransitiveActivityIntermediateType{IRI: v}}, t.generator...)

}

// RemoveGeneratorIRI deletes the value from the specified index
func (t *IntransitiveActivity) RemoveGeneratorIRI(index int) {
	copy(t.generator[index:], t.generator[index+1:])
	t.generator[len(t.generator)-1] = nil
	t.generator = t.generator[:len(t.generator)-1]

}

// HasUnknownGenerator determines whether the call to GetUnknownGenerator is safe
func (t *IntransitiveActivity) HasUnknownGenerator() (ok bool) {
	return t.generator != nil && t.generator[0].unknown_ != nil

}

// GetUnknownGenerator returns the unknown value for generator
func (t *IntransitiveActivity) GetUnknownGenerator() (v interface{}) {
	return t.generator[0].unknown_

}

// SetUnknownGenerator sets the unknown value of generator
func (t *IntransitiveActivity) SetUnknownGenerator(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &generatorIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.generator = append(t.generator, tmp)

}

// IconLen determines the number of elements able to be used for the IsIconImage, GetIconImage, and RemoveIconImage functions
func (t *IntransitiveActivity) IconLen() (l int) {
	return len(t.icon)

}

// IsIconImage determines whether the call to GetIconImage is safe for the specified index
func (t *IntransitiveActivity) IsIconImage(index int) (ok bool) {
	return t.icon[index].Image != nil

}

// GetIconImage returns the value safely if IsIconImage returned true for the specified index
func (t *IntransitiveActivity) GetIconImage(index int) (v ImageType) {
	return t.icon[index].Image

}

// AppendIconImage adds to the back of icon a ImageType type
func (t *IntransitiveActivity) AppendIconImage(v ImageType) {
	t.icon = append(t.icon, &iconIntransitiveActivityIntermediateType{Image: v})

}

// PrependIconImage adds to the front of icon a ImageType type
func (t *IntransitiveActivity) PrependIconImage(v ImageType) {
	t.icon = append([]*iconIntransitiveActivityIntermediateType{&iconIntransitiveActivityIntermediateType{Image: v}}, t.icon...)

}

// RemoveIconImage deletes the value from the specified index
func (t *IntransitiveActivity) RemoveIconImage(index int) {
	copy(t.icon[index:], t.icon[index+1:])
	t.icon[len(t.icon)-1] = nil
	t.icon = t.icon[:len(t.icon)-1]

}

// IsIconLink determines whether the call to GetIconLink is safe for the specified index
func (t *IntransitiveActivity) IsIconLink(index int) (ok bool) {
	return t.icon[index].Link != nil

}

// GetIconLink returns the value safely if IsIconLink returned true for the specified index
func (t *IntransitiveActivity) GetIconLink(index int) (v LinkType) {
	return t.icon[index].Link

}

// AppendIconLink adds to the back of icon a LinkType type
func (t *IntransitiveActivity) AppendIconLink(v LinkType) {
	t.icon = append(t.icon, &iconIntransitiveActivityIntermediateType{Link: v})

}

// PrependIconLink adds to the front of icon a LinkType type
func (t *IntransitiveActivity) PrependIconLink(v LinkType) {
	t.icon = append([]*iconIntransitiveActivityIntermediateType{&iconIntransitiveActivityIntermediateType{Link: v}}, t.icon...)

}

// RemoveIconLink deletes the value from the specified index
func (t *IntransitiveActivity) RemoveIconLink(index int) {
	copy(t.icon[index:], t.icon[index+1:])
	t.icon[len(t.icon)-1] = nil
	t.icon = t.icon[:len(t.icon)-1]

}

// IsIconIRI determines whether the call to GetIconIRI is safe for the specified index
func (t *IntransitiveActivity) IsIconIRI(index int) (ok bool) {
	return t.icon[index].IRI != nil

}

// GetIconIRI returns the value safely if IsIconIRI returned true for the specified index
func (t *IntransitiveActivity) GetIconIRI(index int) (v *url.URL) {
	return t.icon[index].IRI

}

// AppendIconIRI adds to the back of icon a *url.URL type
func (t *IntransitiveActivity) AppendIconIRI(v *url.URL) {
	t.icon = append(t.icon, &iconIntransitiveActivityIntermediateType{IRI: v})

}

// PrependIconIRI adds to the front of icon a *url.URL type
func (t *IntransitiveActivity) PrependIconIRI(v *url.URL) {
	t.icon = append([]*iconIntransitiveActivityIntermediateType{&iconIntransitiveActivityIntermediateType{IRI: v}}, t.icon...)

}

// RemoveIconIRI deletes the value from the specified index
func (t *IntransitiveActivity) RemoveIconIRI(index int) {
	copy(t.icon[index:], t.icon[index+1:])
	t.icon[len(t.icon)-1] = nil
	t.icon = t.icon[:len(t.icon)-1]

}

// HasUnknownIcon determines whether the call to GetUnknownIcon is safe
func (t *IntransitiveActivity) HasUnknownIcon() (ok bool) {
	return t.icon != nil && t.icon[0].unknown_ != nil

}

// GetUnknownIcon returns the unknown value for icon
func (t *IntransitiveActivity) GetUnknownIcon() (v interface{}) {
	return t.icon[0].unknown_

}

// SetUnknownIcon sets the unknown value of icon
func (t *IntransitiveActivity) SetUnknownIcon(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &iconIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.icon = append(t.icon, tmp)

}

// HasId determines whether the call to GetId is safe
func (t *IntransitiveActivity) HasId() (ok bool) {
	return t.id != nil

}

// GetId returns the value for id
func (t *IntransitiveActivity) GetId() (v *url.URL) {
	return t.id

}

// SetId sets the value of id
func (t *IntransitiveActivity) SetId(v *url.URL) {
	t.id = v

}

// HasUnknownId determines whether the call to GetUnknownId is safe
func (t *IntransitiveActivity) HasUnknownId() (ok bool) {
	return t.unknown_ != nil && t.unknown_["id"] != nil

}

// GetUnknownId returns the unknown value for id
func (t *IntransitiveActivity) GetUnknownId() (v interface{}) {
	return t.unknown_["id"]

}

// SetUnknownId sets the unknown value of id
func (t *IntransitiveActivity) SetUnknownId(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["id"] = i

}

// ImageLen determines the number of elements able to be used for the IsImageImage, GetImageImage, and RemoveImageImage functions
func (t *IntransitiveActivity) ImageLen() (l int) {
	return len(t.image)

}

// IsImageImage determines whether the call to GetImageImage is safe for the specified index
func (t *IntransitiveActivity) IsImageImage(index int) (ok bool) {
	return t.image[index].Image != nil

}

// GetImageImage returns the value safely if IsImageImage returned true for the specified index
func (t *IntransitiveActivity) GetImageImage(index int) (v ImageType) {
	return t.image[index].Image

}

// AppendImageImage adds to the back of image a ImageType type
func (t *IntransitiveActivity) AppendImageImage(v ImageType) {
	t.image = append(t.image, &imageIntransitiveActivityIntermediateType{Image: v})

}

// PrependImageImage adds to the front of image a ImageType type
func (t *IntransitiveActivity) PrependImageImage(v ImageType) {
	t.image = append([]*imageIntransitiveActivityIntermediateType{&imageIntransitiveActivityIntermediateType{Image: v}}, t.image...)

}

// RemoveImageImage deletes the value from the specified index
func (t *IntransitiveActivity) RemoveImageImage(index int) {
	copy(t.image[index:], t.image[index+1:])
	t.image[len(t.image)-1] = nil
	t.image = t.image[:len(t.image)-1]

}

// IsImageLink determines whether the call to GetImageLink is safe for the specified index
func (t *IntransitiveActivity) IsImageLink(index int) (ok bool) {
	return t.image[index].Link != nil

}

// GetImageLink returns the value safely if IsImageLink returned true for the specified index
func (t *IntransitiveActivity) GetImageLink(index int) (v LinkType) {
	return t.image[index].Link

}

// AppendImageLink adds to the back of image a LinkType type
func (t *IntransitiveActivity) AppendImageLink(v LinkType) {
	t.image = append(t.image, &imageIntransitiveActivityIntermediateType{Link: v})

}

// PrependImageLink adds to the front of image a LinkType type
func (t *IntransitiveActivity) PrependImageLink(v LinkType) {
	t.image = append([]*imageIntransitiveActivityIntermediateType{&imageIntransitiveActivityIntermediateType{Link: v}}, t.image...)

}

// RemoveImageLink deletes the value from the specified index
func (t *IntransitiveActivity) RemoveImageLink(index int) {
	copy(t.image[index:], t.image[index+1:])
	t.image[len(t.image)-1] = nil
	t.image = t.image[:len(t.image)-1]

}

// IsImageIRI determines whether the call to GetImageIRI is safe for the specified index
func (t *IntransitiveActivity) IsImageIRI(index int) (ok bool) {
	return t.image[index].IRI != nil

}

// GetImageIRI returns the value safely if IsImageIRI returned true for the specified index
func (t *IntransitiveActivity) GetImageIRI(index int) (v *url.URL) {
	return t.image[index].IRI

}

// AppendImageIRI adds to the back of image a *url.URL type
func (t *IntransitiveActivity) AppendImageIRI(v *url.URL) {
	t.image = append(t.image, &imageIntransitiveActivityIntermediateType{IRI: v})

}

// PrependImageIRI adds to the front of image a *url.URL type
func (t *IntransitiveActivity) PrependImageIRI(v *url.URL) {
	t.image = append([]*imageIntransitiveActivityIntermediateType{&imageIntransitiveActivityIntermediateType{IRI: v}}, t.image...)

}

// RemoveImageIRI deletes the value from the specified index
func (t *IntransitiveActivity) RemoveImageIRI(index int) {
	copy(t.image[index:], t.image[index+1:])
	t.image[len(t.image)-1] = nil
	t.image = t.image[:len(t.image)-1]

}

// HasUnknownImage determines whether the call to GetUnknownImage is safe
func (t *IntransitiveActivity) HasUnknownImage() (ok bool) {
	return t.image != nil && t.image[0].unknown_ != nil

}

// GetUnknownImage returns the unknown value for image
func (t *IntransitiveActivity) GetUnknownImage() (v interface{}) {
	return t.image[0].unknown_

}

// SetUnknownImage sets the unknown value of image
func (t *IntransitiveActivity) SetUnknownImage(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &imageIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.image = append(t.image, tmp)

}

// InReplyToLen determines the number of elements able to be used for the IsInReplyToObject, GetInReplyToObject, and RemoveInReplyToObject functions
func (t *IntransitiveActivity) InReplyToLen() (l int) {
	return len(t.inReplyTo)

}

// IsInReplyToObject determines whether the call to GetInReplyToObject is safe for the specified index
func (t *IntransitiveActivity) IsInReplyToObject(index int) (ok bool) {
	return t.inReplyTo[index].Object != nil

}

// GetInReplyToObject returns the value safely if IsInReplyToObject returned true for the specified index
func (t *IntransitiveActivity) GetInReplyToObject(index int) (v ObjectType) {
	return t.inReplyTo[index].Object

}

// AppendInReplyToObject adds to the back of inReplyTo a ObjectType type
func (t *IntransitiveActivity) AppendInReplyToObject(v ObjectType) {
	t.inReplyTo = append(t.inReplyTo, &inReplyToIntransitiveActivityIntermediateType{Object: v})

}

// PrependInReplyToObject adds to the front of inReplyTo a ObjectType type
func (t *IntransitiveActivity) PrependInReplyToObject(v ObjectType) {
	t.inReplyTo = append([]*inReplyToIntransitiveActivityIntermediateType{&inReplyToIntransitiveActivityIntermediateType{Object: v}}, t.inReplyTo...)

}

// RemoveInReplyToObject deletes the value from the specified index
func (t *IntransitiveActivity) RemoveInReplyToObject(index int) {
	copy(t.inReplyTo[index:], t.inReplyTo[index+1:])
	t.inReplyTo[len(t.inReplyTo)-1] = nil
	t.inReplyTo = t.inReplyTo[:len(t.inReplyTo)-1]

}

// IsInReplyToLink determines whether the call to GetInReplyToLink is safe for the specified index
func (t *IntransitiveActivity) IsInReplyToLink(index int) (ok bool) {
	return t.inReplyTo[index].Link != nil

}

// GetInReplyToLink returns the value safely if IsInReplyToLink returned true for the specified index
func (t *IntransitiveActivity) GetInReplyToLink(index int) (v LinkType) {
	return t.inReplyTo[index].Link

}

// AppendInReplyToLink adds to the back of inReplyTo a LinkType type
func (t *IntransitiveActivity) AppendInReplyToLink(v LinkType) {
	t.inReplyTo = append(t.inReplyTo, &inReplyToIntransitiveActivityIntermediateType{Link: v})

}

// PrependInReplyToLink adds to the front of inReplyTo a LinkType type
func (t *IntransitiveActivity) PrependInReplyToLink(v LinkType) {
	t.inReplyTo = append([]*inReplyToIntransitiveActivityIntermediateType{&inReplyToIntransitiveActivityIntermediateType{Link: v}}, t.inReplyTo...)

}

// RemoveInReplyToLink deletes the value from the specified index
func (t *IntransitiveActivity) RemoveInReplyToLink(index int) {
	copy(t.inReplyTo[index:], t.inReplyTo[index+1:])
	t.inReplyTo[len(t.inReplyTo)-1] = nil
	t.inReplyTo = t.inReplyTo[:len(t.inReplyTo)-1]

}

// IsInReplyToIRI determines whether the call to GetInReplyToIRI is safe for the specified index
func (t *IntransitiveActivity) IsInReplyToIRI(index int) (ok bool) {
	return t.inReplyTo[index].IRI != nil

}

// GetInReplyToIRI returns the value safely if IsInReplyToIRI returned true for the specified index
func (t *IntransitiveActivity) GetInReplyToIRI(index int) (v *url.URL) {
	return t.inReplyTo[index].IRI

}

// AppendInReplyToIRI adds to the back of inReplyTo a *url.URL type
func (t *IntransitiveActivity) AppendInReplyToIRI(v *url.URL) {
	t.inReplyTo = append(t.inReplyTo, &inReplyToIntransitiveActivityIntermediateType{IRI: v})

}

// PrependInReplyToIRI adds to the front of inReplyTo a *url.URL type
func (t *IntransitiveActivity) PrependInReplyToIRI(v *url.URL) {
	t.inReplyTo = append([]*inReplyToIntransitiveActivityIntermediateType{&inReplyToIntransitiveActivityIntermediateType{IRI: v}}, t.inReplyTo...)

}

// RemoveInReplyToIRI deletes the value from the specified index
func (t *IntransitiveActivity) RemoveInReplyToIRI(index int) {
	copy(t.inReplyTo[index:], t.inReplyTo[index+1:])
	t.inReplyTo[len(t.inReplyTo)-1] = nil
	t.inReplyTo = t.inReplyTo[:len(t.inReplyTo)-1]

}

// HasUnknownInReplyTo determines whether the call to GetUnknownInReplyTo is safe
func (t *IntransitiveActivity) HasUnknownInReplyTo() (ok bool) {
	return t.inReplyTo != nil && t.inReplyTo[0].unknown_ != nil

}

// GetUnknownInReplyTo returns the unknown value for inReplyTo
func (t *IntransitiveActivity) GetUnknownInReplyTo() (v interface{}) {
	return t.inReplyTo[0].unknown_

}

// SetUnknownInReplyTo sets the unknown value of inReplyTo
func (t *IntransitiveActivity) SetUnknownInReplyTo(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &inReplyToIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.inReplyTo = append(t.inReplyTo, tmp)

}

// LocationLen determines the number of elements able to be used for the IsLocationObject, GetLocationObject, and RemoveLocationObject functions
func (t *IntransitiveActivity) LocationLen() (l int) {
	return len(t.location)

}

// IsLocationObject determines whether the call to GetLocationObject is safe for the specified index
func (t *IntransitiveActivity) IsLocationObject(index int) (ok bool) {
	return t.location[index].Object != nil

}

// GetLocationObject returns the value safely if IsLocationObject returned true for the specified index
func (t *IntransitiveActivity) GetLocationObject(index int) (v ObjectType) {
	return t.location[index].Object

}

// AppendLocationObject adds to the back of location a ObjectType type
func (t *IntransitiveActivity) AppendLocationObject(v ObjectType) {
	t.location = append(t.location, &locationIntransitiveActivityIntermediateType{Object: v})

}

// PrependLocationObject adds to the front of location a ObjectType type
func (t *IntransitiveActivity) PrependLocationObject(v ObjectType) {
	t.location = append([]*locationIntransitiveActivityIntermediateType{&locationIntransitiveActivityIntermediateType{Object: v}}, t.location...)

}

// RemoveLocationObject deletes the value from the specified index
func (t *IntransitiveActivity) RemoveLocationObject(index int) {
	copy(t.location[index:], t.location[index+1:])
	t.location[len(t.location)-1] = nil
	t.location = t.location[:len(t.location)-1]

}

// IsLocationLink determines whether the call to GetLocationLink is safe for the specified index
func (t *IntransitiveActivity) IsLocationLink(index int) (ok bool) {
	return t.location[index].Link != nil

}

// GetLocationLink returns the value safely if IsLocationLink returned true for the specified index
func (t *IntransitiveActivity) GetLocationLink(index int) (v LinkType) {
	return t.location[index].Link

}

// AppendLocationLink adds to the back of location a LinkType type
func (t *IntransitiveActivity) AppendLocationLink(v LinkType) {
	t.location = append(t.location, &locationIntransitiveActivityIntermediateType{Link: v})

}

// PrependLocationLink adds to the front of location a LinkType type
func (t *IntransitiveActivity) PrependLocationLink(v LinkType) {
	t.location = append([]*locationIntransitiveActivityIntermediateType{&locationIntransitiveActivityIntermediateType{Link: v}}, t.location...)

}

// RemoveLocationLink deletes the value from the specified index
func (t *IntransitiveActivity) RemoveLocationLink(index int) {
	copy(t.location[index:], t.location[index+1:])
	t.location[len(t.location)-1] = nil
	t.location = t.location[:len(t.location)-1]

}

// IsLocationIRI determines whether the call to GetLocationIRI is safe for the specified index
func (t *IntransitiveActivity) IsLocationIRI(index int) (ok bool) {
	return t.location[index].IRI != nil

}

// GetLocationIRI returns the value safely if IsLocationIRI returned true for the specified index
func (t *IntransitiveActivity) GetLocationIRI(index int) (v *url.URL) {
	return t.location[index].IRI

}

// AppendLocationIRI adds to the back of location a *url.URL type
func (t *IntransitiveActivity) AppendLocationIRI(v *url.URL) {
	t.location = append(t.location, &locationIntransitiveActivityIntermediateType{IRI: v})

}

// PrependLocationIRI adds to the front of location a *url.URL type
func (t *IntransitiveActivity) PrependLocationIRI(v *url.URL) {
	t.location = append([]*locationIntransitiveActivityIntermediateType{&locationIntransitiveActivityIntermediateType{IRI: v}}, t.location...)

}

// RemoveLocationIRI deletes the value from the specified index
func (t *IntransitiveActivity) RemoveLocationIRI(index int) {
	copy(t.location[index:], t.location[index+1:])
	t.location[len(t.location)-1] = nil
	t.location = t.location[:len(t.location)-1]

}

// HasUnknownLocation determines whether the call to GetUnknownLocation is safe
func (t *IntransitiveActivity) HasUnknownLocation() (ok bool) {
	return t.location != nil && t.location[0].unknown_ != nil

}

// GetUnknownLocation returns the unknown value for location
func (t *IntransitiveActivity) GetUnknownLocation() (v interface{}) {
	return t.location[0].unknown_

}

// SetUnknownLocation sets the unknown value of location
func (t *IntransitiveActivity) SetUnknownLocation(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &locationIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.location = append(t.location, tmp)

}

// PreviewLen determines the number of elements able to be used for the IsPreviewObject, GetPreviewObject, and RemovePreviewObject functions
func (t *IntransitiveActivity) PreviewLen() (l int) {
	return len(t.preview)

}

// IsPreviewObject determines whether the call to GetPreviewObject is safe for the specified index
func (t *IntransitiveActivity) IsPreviewObject(index int) (ok bool) {
	return t.preview[index].Object != nil

}

// GetPreviewObject returns the value safely if IsPreviewObject returned true for the specified index
func (t *IntransitiveActivity) GetPreviewObject(index int) (v ObjectType) {
	return t.preview[index].Object

}

// AppendPreviewObject adds to the back of preview a ObjectType type
func (t *IntransitiveActivity) AppendPreviewObject(v ObjectType) {
	t.preview = append(t.preview, &previewIntransitiveActivityIntermediateType{Object: v})

}

// PrependPreviewObject adds to the front of preview a ObjectType type
func (t *IntransitiveActivity) PrependPreviewObject(v ObjectType) {
	t.preview = append([]*previewIntransitiveActivityIntermediateType{&previewIntransitiveActivityIntermediateType{Object: v}}, t.preview...)

}

// RemovePreviewObject deletes the value from the specified index
func (t *IntransitiveActivity) RemovePreviewObject(index int) {
	copy(t.preview[index:], t.preview[index+1:])
	t.preview[len(t.preview)-1] = nil
	t.preview = t.preview[:len(t.preview)-1]

}

// IsPreviewLink determines whether the call to GetPreviewLink is safe for the specified index
func (t *IntransitiveActivity) IsPreviewLink(index int) (ok bool) {
	return t.preview[index].Link != nil

}

// GetPreviewLink returns the value safely if IsPreviewLink returned true for the specified index
func (t *IntransitiveActivity) GetPreviewLink(index int) (v LinkType) {
	return t.preview[index].Link

}

// AppendPreviewLink adds to the back of preview a LinkType type
func (t *IntransitiveActivity) AppendPreviewLink(v LinkType) {
	t.preview = append(t.preview, &previewIntransitiveActivityIntermediateType{Link: v})

}

// PrependPreviewLink adds to the front of preview a LinkType type
func (t *IntransitiveActivity) PrependPreviewLink(v LinkType) {
	t.preview = append([]*previewIntransitiveActivityIntermediateType{&previewIntransitiveActivityIntermediateType{Link: v}}, t.preview...)

}

// RemovePreviewLink deletes the value from the specified index
func (t *IntransitiveActivity) RemovePreviewLink(index int) {
	copy(t.preview[index:], t.preview[index+1:])
	t.preview[len(t.preview)-1] = nil
	t.preview = t.preview[:len(t.preview)-1]

}

// IsPreviewIRI determines whether the call to GetPreviewIRI is safe for the specified index
func (t *IntransitiveActivity) IsPreviewIRI(index int) (ok bool) {
	return t.preview[index].IRI != nil

}

// GetPreviewIRI returns the value safely if IsPreviewIRI returned true for the specified index
func (t *IntransitiveActivity) GetPreviewIRI(index int) (v *url.URL) {
	return t.preview[index].IRI

}

// AppendPreviewIRI adds to the back of preview a *url.URL type
func (t *IntransitiveActivity) AppendPreviewIRI(v *url.URL) {
	t.preview = append(t.preview, &previewIntransitiveActivityIntermediateType{IRI: v})

}

// PrependPreviewIRI adds to the front of preview a *url.URL type
func (t *IntransitiveActivity) PrependPreviewIRI(v *url.URL) {
	t.preview = append([]*previewIntransitiveActivityIntermediateType{&previewIntransitiveActivityIntermediateType{IRI: v}}, t.preview...)

}

// RemovePreviewIRI deletes the value from the specified index
func (t *IntransitiveActivity) RemovePreviewIRI(index int) {
	copy(t.preview[index:], t.preview[index+1:])
	t.preview[len(t.preview)-1] = nil
	t.preview = t.preview[:len(t.preview)-1]

}

// HasUnknownPreview determines whether the call to GetUnknownPreview is safe
func (t *IntransitiveActivity) HasUnknownPreview() (ok bool) {
	return t.preview != nil && t.preview[0].unknown_ != nil

}

// GetUnknownPreview returns the unknown value for preview
func (t *IntransitiveActivity) GetUnknownPreview() (v interface{}) {
	return t.preview[0].unknown_

}

// SetUnknownPreview sets the unknown value of preview
func (t *IntransitiveActivity) SetUnknownPreview(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &previewIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.preview = append(t.preview, tmp)

}

// IsPublished determines whether the call to GetPublished is safe
func (t *IntransitiveActivity) IsPublished() (ok bool) {
	return t.published != nil && t.published.dateTime != nil

}

// GetPublished returns the value safely if IsPublished returned true
func (t *IntransitiveActivity) GetPublished() (v time.Time) {
	return *t.published.dateTime

}

// SetPublished sets the value of published to be of time.Time type
func (t *IntransitiveActivity) SetPublished(v time.Time) {
	t.published = &publishedIntransitiveActivityIntermediateType{dateTime: &v}

}

// IsPublishedIRI determines whether the call to GetPublishedIRI is safe
func (t *IntransitiveActivity) IsPublishedIRI() (ok bool) {
	return t.published != nil && t.published.IRI != nil

}

// GetPublishedIRI returns the value safely if IsPublishedIRI returned true
func (t *IntransitiveActivity) GetPublishedIRI() (v *url.URL) {
	return t.published.IRI

}

// SetPublishedIRI sets the value of published to be of *url.URL type
func (t *IntransitiveActivity) SetPublishedIRI(v *url.URL) {
	t.published = &publishedIntransitiveActivityIntermediateType{IRI: v}

}

// HasUnknownPublished determines whether the call to GetUnknownPublished is safe
func (t *IntransitiveActivity) HasUnknownPublished() (ok bool) {
	return t.published != nil && t.published.unknown_ != nil

}

// GetUnknownPublished returns the unknown value for published
func (t *IntransitiveActivity) GetUnknownPublished() (v interface{}) {
	return t.published.unknown_

}

// SetUnknownPublished sets the unknown value of published
func (t *IntransitiveActivity) SetUnknownPublished(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &publishedIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.published = tmp

}

// IsReplies determines whether the call to GetReplies is safe
func (t *IntransitiveActivity) IsReplies() (ok bool) {
	return t.replies != nil && t.replies.Collection != nil

}

// GetReplies returns the value safely if IsReplies returned true
func (t *IntransitiveActivity) GetReplies() (v CollectionType) {
	return t.replies.Collection

}

// SetReplies sets the value of replies to be of CollectionType type
func (t *IntransitiveActivity) SetReplies(v CollectionType) {
	t.replies = &repliesIntransitiveActivityIntermediateType{Collection: v}

}

// IsRepliesIRI determines whether the call to GetRepliesIRI is safe
func (t *IntransitiveActivity) IsRepliesIRI() (ok bool) {
	return t.replies != nil && t.replies.IRI != nil

}

// GetRepliesIRI returns the value safely if IsRepliesIRI returned true
func (t *IntransitiveActivity) GetRepliesIRI() (v *url.URL) {
	return t.replies.IRI

}

// SetRepliesIRI sets the value of replies to be of *url.URL type
func (t *IntransitiveActivity) SetRepliesIRI(v *url.URL) {
	t.replies = &repliesIntransitiveActivityIntermediateType{IRI: v}

}

// HasUnknownReplies determines whether the call to GetUnknownReplies is safe
func (t *IntransitiveActivity) HasUnknownReplies() (ok bool) {
	return t.replies != nil && t.replies.unknown_ != nil

}

// GetUnknownReplies returns the unknown value for replies
func (t *IntransitiveActivity) GetUnknownReplies() (v interface{}) {
	return t.replies.unknown_

}

// SetUnknownReplies sets the unknown value of replies
func (t *IntransitiveActivity) SetUnknownReplies(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &repliesIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.replies = tmp

}

// IsStartTime determines whether the call to GetStartTime is safe
func (t *IntransitiveActivity) IsStartTime() (ok bool) {
	return t.startTime != nil && t.startTime.dateTime != nil

}

// GetStartTime returns the value safely if IsStartTime returned true
func (t *IntransitiveActivity) GetStartTime() (v time.Time) {
	return *t.startTime.dateTime

}

// SetStartTime sets the value of startTime to be of time.Time type
func (t *IntransitiveActivity) SetStartTime(v time.Time) {
	t.startTime = &startTimeIntransitiveActivityIntermediateType{dateTime: &v}

}

// IsStartTimeIRI determines whether the call to GetStartTimeIRI is safe
func (t *IntransitiveActivity) IsStartTimeIRI() (ok bool) {
	return t.startTime != nil && t.startTime.IRI != nil

}

// GetStartTimeIRI returns the value safely if IsStartTimeIRI returned true
func (t *IntransitiveActivity) GetStartTimeIRI() (v *url.URL) {
	return t.startTime.IRI

}

// SetStartTimeIRI sets the value of startTime to be of *url.URL type
func (t *IntransitiveActivity) SetStartTimeIRI(v *url.URL) {
	t.startTime = &startTimeIntransitiveActivityIntermediateType{IRI: v}

}

// HasUnknownStartTime determines whether the call to GetUnknownStartTime is safe
func (t *IntransitiveActivity) HasUnknownStartTime() (ok bool) {
	return t.startTime != nil && t.startTime.unknown_ != nil

}

// GetUnknownStartTime returns the unknown value for startTime
func (t *IntransitiveActivity) GetUnknownStartTime() (v interface{}) {
	return t.startTime.unknown_

}

// SetUnknownStartTime sets the unknown value of startTime
func (t *IntransitiveActivity) SetUnknownStartTime(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &startTimeIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.startTime = tmp

}

// SummaryLen determines the number of elements able to be used for the IsSummaryString, GetSummaryString, and RemoveSummaryString functions
func (t *IntransitiveActivity) SummaryLen() (l int) {
	return len(t.summary)

}

// IsSummaryString determines whether the call to GetSummaryString is safe for the specified index
func (t *IntransitiveActivity) IsSummaryString(index int) (ok bool) {
	return t.summary[index].stringName != nil

}

// GetSummaryString returns the value safely if IsSummaryString returned true for the specified index
func (t *IntransitiveActivity) GetSummaryString(index int) (v string) {
	return *t.summary[index].stringName

}

// AppendSummaryString adds to the back of summary a string type
func (t *IntransitiveActivity) AppendSummaryString(v string) {
	t.summary = append(t.summary, &summaryIntransitiveActivityIntermediateType{stringName: &v})

}

// PrependSummaryString adds to the front of summary a string type
func (t *IntransitiveActivity) PrependSummaryString(v string) {
	t.summary = append([]*summaryIntransitiveActivityIntermediateType{&summaryIntransitiveActivityIntermediateType{stringName: &v}}, t.summary...)

}

// RemoveSummaryString deletes the value from the specified index
func (t *IntransitiveActivity) RemoveSummaryString(index int) {
	copy(t.summary[index:], t.summary[index+1:])
	t.summary[len(t.summary)-1] = nil
	t.summary = t.summary[:len(t.summary)-1]

}

// IsSummaryLangString determines whether the call to GetSummaryLangString is safe for the specified index
func (t *IntransitiveActivity) IsSummaryLangString(index int) (ok bool) {
	return t.summary[index].langString != nil

}

// GetSummaryLangString returns the value safely if IsSummaryLangString returned true for the specified index
func (t *IntransitiveActivity) GetSummaryLangString(index int) (v string) {
	return *t.summary[index].langString

}

// AppendSummaryLangString adds to the back of summary a string type
func (t *IntransitiveActivity) AppendSummaryLangString(v string) {
	t.summary = append(t.summary, &summaryIntransitiveActivityIntermediateType{langString: &v})

}

// PrependSummaryLangString adds to the front of summary a string type
func (t *IntransitiveActivity) PrependSummaryLangString(v string) {
	t.summary = append([]*summaryIntransitiveActivityIntermediateType{&summaryIntransitiveActivityIntermediateType{langString: &v}}, t.summary...)

}

// RemoveSummaryLangString deletes the value from the specified index
func (t *IntransitiveActivity) RemoveSummaryLangString(index int) {
	copy(t.summary[index:], t.summary[index+1:])
	t.summary[len(t.summary)-1] = nil
	t.summary = t.summary[:len(t.summary)-1]

}

// IsSummaryIRI determines whether the call to GetSummaryIRI is safe for the specified index
func (t *IntransitiveActivity) IsSummaryIRI(index int) (ok bool) {
	return t.summary[index].IRI != nil

}

// GetSummaryIRI returns the value safely if IsSummaryIRI returned true for the specified index
func (t *IntransitiveActivity) GetSummaryIRI(index int) (v *url.URL) {
	return t.summary[index].IRI

}

// AppendSummaryIRI adds to the back of summary a *url.URL type
func (t *IntransitiveActivity) AppendSummaryIRI(v *url.URL) {
	t.summary = append(t.summary, &summaryIntransitiveActivityIntermediateType{IRI: v})

}

// PrependSummaryIRI adds to the front of summary a *url.URL type
func (t *IntransitiveActivity) PrependSummaryIRI(v *url.URL) {
	t.summary = append([]*summaryIntransitiveActivityIntermediateType{&summaryIntransitiveActivityIntermediateType{IRI: v}}, t.summary...)

}

// RemoveSummaryIRI deletes the value from the specified index
func (t *IntransitiveActivity) RemoveSummaryIRI(index int) {
	copy(t.summary[index:], t.summary[index+1:])
	t.summary[len(t.summary)-1] = nil
	t.summary = t.summary[:len(t.summary)-1]

}

// HasUnknownSummary determines whether the call to GetUnknownSummary is safe
func (t *IntransitiveActivity) HasUnknownSummary() (ok bool) {
	return t.summary != nil && t.summary[0].unknown_ != nil

}

// GetUnknownSummary returns the unknown value for summary
func (t *IntransitiveActivity) GetUnknownSummary() (v interface{}) {
	return t.summary[0].unknown_

}

// SetUnknownSummary sets the unknown value of summary
func (t *IntransitiveActivity) SetUnknownSummary(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &summaryIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.summary = append(t.summary, tmp)

}

// SummaryMapLanguages returns all languages for this property's language mapping, or nil if there are none.
func (t *IntransitiveActivity) SummaryMapLanguages() (l []string) {
	if t.summaryMap == nil || len(t.summaryMap) == 0 {
		return nil
	}
	for k := range t.summaryMap {
		l = append(l, k)
	}
	return

}

// GetSummaryMap retrieves the value of the property for the specified language, or an empty string if it does not exist
func (t *IntransitiveActivity) GetSummaryMap(l string) (v string) {
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
func (t *IntransitiveActivity) SetSummaryMap(l string, v string) {
	if t.summaryMap == nil {
		t.summaryMap = make(map[string]string)
	}
	t.summaryMap[l] = v

}

// TagLen determines the number of elements able to be used for the IsTagObject, GetTagObject, and RemoveTagObject functions
func (t *IntransitiveActivity) TagLen() (l int) {
	return len(t.tag)

}

// IsTagObject determines whether the call to GetTagObject is safe for the specified index
func (t *IntransitiveActivity) IsTagObject(index int) (ok bool) {
	return t.tag[index].Object != nil

}

// GetTagObject returns the value safely if IsTagObject returned true for the specified index
func (t *IntransitiveActivity) GetTagObject(index int) (v ObjectType) {
	return t.tag[index].Object

}

// AppendTagObject adds to the back of tag a ObjectType type
func (t *IntransitiveActivity) AppendTagObject(v ObjectType) {
	t.tag = append(t.tag, &tagIntransitiveActivityIntermediateType{Object: v})

}

// PrependTagObject adds to the front of tag a ObjectType type
func (t *IntransitiveActivity) PrependTagObject(v ObjectType) {
	t.tag = append([]*tagIntransitiveActivityIntermediateType{&tagIntransitiveActivityIntermediateType{Object: v}}, t.tag...)

}

// RemoveTagObject deletes the value from the specified index
func (t *IntransitiveActivity) RemoveTagObject(index int) {
	copy(t.tag[index:], t.tag[index+1:])
	t.tag[len(t.tag)-1] = nil
	t.tag = t.tag[:len(t.tag)-1]

}

// IsTagLink determines whether the call to GetTagLink is safe for the specified index
func (t *IntransitiveActivity) IsTagLink(index int) (ok bool) {
	return t.tag[index].Link != nil

}

// GetTagLink returns the value safely if IsTagLink returned true for the specified index
func (t *IntransitiveActivity) GetTagLink(index int) (v LinkType) {
	return t.tag[index].Link

}

// AppendTagLink adds to the back of tag a LinkType type
func (t *IntransitiveActivity) AppendTagLink(v LinkType) {
	t.tag = append(t.tag, &tagIntransitiveActivityIntermediateType{Link: v})

}

// PrependTagLink adds to the front of tag a LinkType type
func (t *IntransitiveActivity) PrependTagLink(v LinkType) {
	t.tag = append([]*tagIntransitiveActivityIntermediateType{&tagIntransitiveActivityIntermediateType{Link: v}}, t.tag...)

}

// RemoveTagLink deletes the value from the specified index
func (t *IntransitiveActivity) RemoveTagLink(index int) {
	copy(t.tag[index:], t.tag[index+1:])
	t.tag[len(t.tag)-1] = nil
	t.tag = t.tag[:len(t.tag)-1]

}

// IsTagIRI determines whether the call to GetTagIRI is safe for the specified index
func (t *IntransitiveActivity) IsTagIRI(index int) (ok bool) {
	return t.tag[index].IRI != nil

}

// GetTagIRI returns the value safely if IsTagIRI returned true for the specified index
func (t *IntransitiveActivity) GetTagIRI(index int) (v *url.URL) {
	return t.tag[index].IRI

}

// AppendTagIRI adds to the back of tag a *url.URL type
func (t *IntransitiveActivity) AppendTagIRI(v *url.URL) {
	t.tag = append(t.tag, &tagIntransitiveActivityIntermediateType{IRI: v})

}

// PrependTagIRI adds to the front of tag a *url.URL type
func (t *IntransitiveActivity) PrependTagIRI(v *url.URL) {
	t.tag = append([]*tagIntransitiveActivityIntermediateType{&tagIntransitiveActivityIntermediateType{IRI: v}}, t.tag...)

}

// RemoveTagIRI deletes the value from the specified index
func (t *IntransitiveActivity) RemoveTagIRI(index int) {
	copy(t.tag[index:], t.tag[index+1:])
	t.tag[len(t.tag)-1] = nil
	t.tag = t.tag[:len(t.tag)-1]

}

// HasUnknownTag determines whether the call to GetUnknownTag is safe
func (t *IntransitiveActivity) HasUnknownTag() (ok bool) {
	return t.tag != nil && t.tag[0].unknown_ != nil

}

// GetUnknownTag returns the unknown value for tag
func (t *IntransitiveActivity) GetUnknownTag() (v interface{}) {
	return t.tag[0].unknown_

}

// SetUnknownTag sets the unknown value of tag
func (t *IntransitiveActivity) SetUnknownTag(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &tagIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.tag = append(t.tag, tmp)

}

// TypeLen determines the number of elements able to be used for the GetType and RemoveType functions
func (t *IntransitiveActivity) TypeLen() (l int) {
	return len(t.typeName)

}

// GetType returns the value for the specified index
func (t *IntransitiveActivity) GetType(index int) (v interface{}) {
	return t.typeName[index]

}

// AppendType adds a value to the back of type
func (t *IntransitiveActivity) AppendType(v interface{}) {
	t.typeName = append(t.typeName, v)

}

// PrependType adds a value to the front of type
func (t *IntransitiveActivity) PrependType(v interface{}) {
	t.typeName = append([]interface{}{v}, t.typeName...)

}

// RemoveType deletes the value from the specified index
func (t *IntransitiveActivity) RemoveType(index int) {
	copy(t.typeName[index:], t.typeName[index+1:])
	t.typeName[len(t.typeName)-1] = nil
	t.typeName = t.typeName[:len(t.typeName)-1]

}

// IsUpdated determines whether the call to GetUpdated is safe
func (t *IntransitiveActivity) IsUpdated() (ok bool) {
	return t.updated != nil && t.updated.dateTime != nil

}

// GetUpdated returns the value safely if IsUpdated returned true
func (t *IntransitiveActivity) GetUpdated() (v time.Time) {
	return *t.updated.dateTime

}

// SetUpdated sets the value of updated to be of time.Time type
func (t *IntransitiveActivity) SetUpdated(v time.Time) {
	t.updated = &updatedIntransitiveActivityIntermediateType{dateTime: &v}

}

// IsUpdatedIRI determines whether the call to GetUpdatedIRI is safe
func (t *IntransitiveActivity) IsUpdatedIRI() (ok bool) {
	return t.updated != nil && t.updated.IRI != nil

}

// GetUpdatedIRI returns the value safely if IsUpdatedIRI returned true
func (t *IntransitiveActivity) GetUpdatedIRI() (v *url.URL) {
	return t.updated.IRI

}

// SetUpdatedIRI sets the value of updated to be of *url.URL type
func (t *IntransitiveActivity) SetUpdatedIRI(v *url.URL) {
	t.updated = &updatedIntransitiveActivityIntermediateType{IRI: v}

}

// HasUnknownUpdated determines whether the call to GetUnknownUpdated is safe
func (t *IntransitiveActivity) HasUnknownUpdated() (ok bool) {
	return t.updated != nil && t.updated.unknown_ != nil

}

// GetUnknownUpdated returns the unknown value for updated
func (t *IntransitiveActivity) GetUnknownUpdated() (v interface{}) {
	return t.updated.unknown_

}

// SetUnknownUpdated sets the unknown value of updated
func (t *IntransitiveActivity) SetUnknownUpdated(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &updatedIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.updated = tmp

}

// UrlLen determines the number of elements able to be used for the IsUrlAnyURI, GetUrlAnyURI, and RemoveUrlAnyURI functions
func (t *IntransitiveActivity) UrlLen() (l int) {
	return len(t.url)

}

// IsUrlAnyURI determines whether the call to GetUrlAnyURI is safe for the specified index
func (t *IntransitiveActivity) IsUrlAnyURI(index int) (ok bool) {
	return t.url[index].anyURI != nil

}

// GetUrlAnyURI returns the value safely if IsUrlAnyURI returned true for the specified index
func (t *IntransitiveActivity) GetUrlAnyURI(index int) (v *url.URL) {
	return t.url[index].anyURI

}

// AppendUrlAnyURI adds to the back of url a *url.URL type
func (t *IntransitiveActivity) AppendUrlAnyURI(v *url.URL) {
	t.url = append(t.url, &urlIntransitiveActivityIntermediateType{anyURI: v})

}

// PrependUrlAnyURI adds to the front of url a *url.URL type
func (t *IntransitiveActivity) PrependUrlAnyURI(v *url.URL) {
	t.url = append([]*urlIntransitiveActivityIntermediateType{&urlIntransitiveActivityIntermediateType{anyURI: v}}, t.url...)

}

// RemoveUrlAnyURI deletes the value from the specified index
func (t *IntransitiveActivity) RemoveUrlAnyURI(index int) {
	copy(t.url[index:], t.url[index+1:])
	t.url[len(t.url)-1] = nil
	t.url = t.url[:len(t.url)-1]

}

// IsUrlLink determines whether the call to GetUrlLink is safe for the specified index
func (t *IntransitiveActivity) IsUrlLink(index int) (ok bool) {
	return t.url[index].Link != nil

}

// GetUrlLink returns the value safely if IsUrlLink returned true for the specified index
func (t *IntransitiveActivity) GetUrlLink(index int) (v LinkType) {
	return t.url[index].Link

}

// AppendUrlLink adds to the back of url a LinkType type
func (t *IntransitiveActivity) AppendUrlLink(v LinkType) {
	t.url = append(t.url, &urlIntransitiveActivityIntermediateType{Link: v})

}

// PrependUrlLink adds to the front of url a LinkType type
func (t *IntransitiveActivity) PrependUrlLink(v LinkType) {
	t.url = append([]*urlIntransitiveActivityIntermediateType{&urlIntransitiveActivityIntermediateType{Link: v}}, t.url...)

}

// RemoveUrlLink deletes the value from the specified index
func (t *IntransitiveActivity) RemoveUrlLink(index int) {
	copy(t.url[index:], t.url[index+1:])
	t.url[len(t.url)-1] = nil
	t.url = t.url[:len(t.url)-1]

}

// HasUnknownUrl determines whether the call to GetUnknownUrl is safe
func (t *IntransitiveActivity) HasUnknownUrl() (ok bool) {
	return t.url != nil && t.url[0].unknown_ != nil

}

// GetUnknownUrl returns the unknown value for url
func (t *IntransitiveActivity) GetUnknownUrl() (v interface{}) {
	return t.url[0].unknown_

}

// SetUnknownUrl sets the unknown value of url
func (t *IntransitiveActivity) SetUnknownUrl(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &urlIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.url = append(t.url, tmp)

}

// ToLen determines the number of elements able to be used for the IsToObject, GetToObject, and RemoveToObject functions
func (t *IntransitiveActivity) ToLen() (l int) {
	return len(t.to)

}

// IsToObject determines whether the call to GetToObject is safe for the specified index
func (t *IntransitiveActivity) IsToObject(index int) (ok bool) {
	return t.to[index].Object != nil

}

// GetToObject returns the value safely if IsToObject returned true for the specified index
func (t *IntransitiveActivity) GetToObject(index int) (v ObjectType) {
	return t.to[index].Object

}

// AppendToObject adds to the back of to a ObjectType type
func (t *IntransitiveActivity) AppendToObject(v ObjectType) {
	t.to = append(t.to, &toIntransitiveActivityIntermediateType{Object: v})

}

// PrependToObject adds to the front of to a ObjectType type
func (t *IntransitiveActivity) PrependToObject(v ObjectType) {
	t.to = append([]*toIntransitiveActivityIntermediateType{&toIntransitiveActivityIntermediateType{Object: v}}, t.to...)

}

// RemoveToObject deletes the value from the specified index
func (t *IntransitiveActivity) RemoveToObject(index int) {
	copy(t.to[index:], t.to[index+1:])
	t.to[len(t.to)-1] = nil
	t.to = t.to[:len(t.to)-1]

}

// IsToLink determines whether the call to GetToLink is safe for the specified index
func (t *IntransitiveActivity) IsToLink(index int) (ok bool) {
	return t.to[index].Link != nil

}

// GetToLink returns the value safely if IsToLink returned true for the specified index
func (t *IntransitiveActivity) GetToLink(index int) (v LinkType) {
	return t.to[index].Link

}

// AppendToLink adds to the back of to a LinkType type
func (t *IntransitiveActivity) AppendToLink(v LinkType) {
	t.to = append(t.to, &toIntransitiveActivityIntermediateType{Link: v})

}

// PrependToLink adds to the front of to a LinkType type
func (t *IntransitiveActivity) PrependToLink(v LinkType) {
	t.to = append([]*toIntransitiveActivityIntermediateType{&toIntransitiveActivityIntermediateType{Link: v}}, t.to...)

}

// RemoveToLink deletes the value from the specified index
func (t *IntransitiveActivity) RemoveToLink(index int) {
	copy(t.to[index:], t.to[index+1:])
	t.to[len(t.to)-1] = nil
	t.to = t.to[:len(t.to)-1]

}

// IsToIRI determines whether the call to GetToIRI is safe for the specified index
func (t *IntransitiveActivity) IsToIRI(index int) (ok bool) {
	return t.to[index].IRI != nil

}

// GetToIRI returns the value safely if IsToIRI returned true for the specified index
func (t *IntransitiveActivity) GetToIRI(index int) (v *url.URL) {
	return t.to[index].IRI

}

// AppendToIRI adds to the back of to a *url.URL type
func (t *IntransitiveActivity) AppendToIRI(v *url.URL) {
	t.to = append(t.to, &toIntransitiveActivityIntermediateType{IRI: v})

}

// PrependToIRI adds to the front of to a *url.URL type
func (t *IntransitiveActivity) PrependToIRI(v *url.URL) {
	t.to = append([]*toIntransitiveActivityIntermediateType{&toIntransitiveActivityIntermediateType{IRI: v}}, t.to...)

}

// RemoveToIRI deletes the value from the specified index
func (t *IntransitiveActivity) RemoveToIRI(index int) {
	copy(t.to[index:], t.to[index+1:])
	t.to[len(t.to)-1] = nil
	t.to = t.to[:len(t.to)-1]

}

// HasUnknownTo determines whether the call to GetUnknownTo is safe
func (t *IntransitiveActivity) HasUnknownTo() (ok bool) {
	return t.to != nil && t.to[0].unknown_ != nil

}

// GetUnknownTo returns the unknown value for to
func (t *IntransitiveActivity) GetUnknownTo() (v interface{}) {
	return t.to[0].unknown_

}

// SetUnknownTo sets the unknown value of to
func (t *IntransitiveActivity) SetUnknownTo(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &toIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.to = append(t.to, tmp)

}

// BtoLen determines the number of elements able to be used for the IsBtoObject, GetBtoObject, and RemoveBtoObject functions
func (t *IntransitiveActivity) BtoLen() (l int) {
	return len(t.bto)

}

// IsBtoObject determines whether the call to GetBtoObject is safe for the specified index
func (t *IntransitiveActivity) IsBtoObject(index int) (ok bool) {
	return t.bto[index].Object != nil

}

// GetBtoObject returns the value safely if IsBtoObject returned true for the specified index
func (t *IntransitiveActivity) GetBtoObject(index int) (v ObjectType) {
	return t.bto[index].Object

}

// AppendBtoObject adds to the back of bto a ObjectType type
func (t *IntransitiveActivity) AppendBtoObject(v ObjectType) {
	t.bto = append(t.bto, &btoIntransitiveActivityIntermediateType{Object: v})

}

// PrependBtoObject adds to the front of bto a ObjectType type
func (t *IntransitiveActivity) PrependBtoObject(v ObjectType) {
	t.bto = append([]*btoIntransitiveActivityIntermediateType{&btoIntransitiveActivityIntermediateType{Object: v}}, t.bto...)

}

// RemoveBtoObject deletes the value from the specified index
func (t *IntransitiveActivity) RemoveBtoObject(index int) {
	copy(t.bto[index:], t.bto[index+1:])
	t.bto[len(t.bto)-1] = nil
	t.bto = t.bto[:len(t.bto)-1]

}

// IsBtoLink determines whether the call to GetBtoLink is safe for the specified index
func (t *IntransitiveActivity) IsBtoLink(index int) (ok bool) {
	return t.bto[index].Link != nil

}

// GetBtoLink returns the value safely if IsBtoLink returned true for the specified index
func (t *IntransitiveActivity) GetBtoLink(index int) (v LinkType) {
	return t.bto[index].Link

}

// AppendBtoLink adds to the back of bto a LinkType type
func (t *IntransitiveActivity) AppendBtoLink(v LinkType) {
	t.bto = append(t.bto, &btoIntransitiveActivityIntermediateType{Link: v})

}

// PrependBtoLink adds to the front of bto a LinkType type
func (t *IntransitiveActivity) PrependBtoLink(v LinkType) {
	t.bto = append([]*btoIntransitiveActivityIntermediateType{&btoIntransitiveActivityIntermediateType{Link: v}}, t.bto...)

}

// RemoveBtoLink deletes the value from the specified index
func (t *IntransitiveActivity) RemoveBtoLink(index int) {
	copy(t.bto[index:], t.bto[index+1:])
	t.bto[len(t.bto)-1] = nil
	t.bto = t.bto[:len(t.bto)-1]

}

// IsBtoIRI determines whether the call to GetBtoIRI is safe for the specified index
func (t *IntransitiveActivity) IsBtoIRI(index int) (ok bool) {
	return t.bto[index].IRI != nil

}

// GetBtoIRI returns the value safely if IsBtoIRI returned true for the specified index
func (t *IntransitiveActivity) GetBtoIRI(index int) (v *url.URL) {
	return t.bto[index].IRI

}

// AppendBtoIRI adds to the back of bto a *url.URL type
func (t *IntransitiveActivity) AppendBtoIRI(v *url.URL) {
	t.bto = append(t.bto, &btoIntransitiveActivityIntermediateType{IRI: v})

}

// PrependBtoIRI adds to the front of bto a *url.URL type
func (t *IntransitiveActivity) PrependBtoIRI(v *url.URL) {
	t.bto = append([]*btoIntransitiveActivityIntermediateType{&btoIntransitiveActivityIntermediateType{IRI: v}}, t.bto...)

}

// RemoveBtoIRI deletes the value from the specified index
func (t *IntransitiveActivity) RemoveBtoIRI(index int) {
	copy(t.bto[index:], t.bto[index+1:])
	t.bto[len(t.bto)-1] = nil
	t.bto = t.bto[:len(t.bto)-1]

}

// HasUnknownBto determines whether the call to GetUnknownBto is safe
func (t *IntransitiveActivity) HasUnknownBto() (ok bool) {
	return t.bto != nil && t.bto[0].unknown_ != nil

}

// GetUnknownBto returns the unknown value for bto
func (t *IntransitiveActivity) GetUnknownBto() (v interface{}) {
	return t.bto[0].unknown_

}

// SetUnknownBto sets the unknown value of bto
func (t *IntransitiveActivity) SetUnknownBto(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &btoIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.bto = append(t.bto, tmp)

}

// CcLen determines the number of elements able to be used for the IsCcObject, GetCcObject, and RemoveCcObject functions
func (t *IntransitiveActivity) CcLen() (l int) {
	return len(t.cc)

}

// IsCcObject determines whether the call to GetCcObject is safe for the specified index
func (t *IntransitiveActivity) IsCcObject(index int) (ok bool) {
	return t.cc[index].Object != nil

}

// GetCcObject returns the value safely if IsCcObject returned true for the specified index
func (t *IntransitiveActivity) GetCcObject(index int) (v ObjectType) {
	return t.cc[index].Object

}

// AppendCcObject adds to the back of cc a ObjectType type
func (t *IntransitiveActivity) AppendCcObject(v ObjectType) {
	t.cc = append(t.cc, &ccIntransitiveActivityIntermediateType{Object: v})

}

// PrependCcObject adds to the front of cc a ObjectType type
func (t *IntransitiveActivity) PrependCcObject(v ObjectType) {
	t.cc = append([]*ccIntransitiveActivityIntermediateType{&ccIntransitiveActivityIntermediateType{Object: v}}, t.cc...)

}

// RemoveCcObject deletes the value from the specified index
func (t *IntransitiveActivity) RemoveCcObject(index int) {
	copy(t.cc[index:], t.cc[index+1:])
	t.cc[len(t.cc)-1] = nil
	t.cc = t.cc[:len(t.cc)-1]

}

// IsCcLink determines whether the call to GetCcLink is safe for the specified index
func (t *IntransitiveActivity) IsCcLink(index int) (ok bool) {
	return t.cc[index].Link != nil

}

// GetCcLink returns the value safely if IsCcLink returned true for the specified index
func (t *IntransitiveActivity) GetCcLink(index int) (v LinkType) {
	return t.cc[index].Link

}

// AppendCcLink adds to the back of cc a LinkType type
func (t *IntransitiveActivity) AppendCcLink(v LinkType) {
	t.cc = append(t.cc, &ccIntransitiveActivityIntermediateType{Link: v})

}

// PrependCcLink adds to the front of cc a LinkType type
func (t *IntransitiveActivity) PrependCcLink(v LinkType) {
	t.cc = append([]*ccIntransitiveActivityIntermediateType{&ccIntransitiveActivityIntermediateType{Link: v}}, t.cc...)

}

// RemoveCcLink deletes the value from the specified index
func (t *IntransitiveActivity) RemoveCcLink(index int) {
	copy(t.cc[index:], t.cc[index+1:])
	t.cc[len(t.cc)-1] = nil
	t.cc = t.cc[:len(t.cc)-1]

}

// IsCcIRI determines whether the call to GetCcIRI is safe for the specified index
func (t *IntransitiveActivity) IsCcIRI(index int) (ok bool) {
	return t.cc[index].IRI != nil

}

// GetCcIRI returns the value safely if IsCcIRI returned true for the specified index
func (t *IntransitiveActivity) GetCcIRI(index int) (v *url.URL) {
	return t.cc[index].IRI

}

// AppendCcIRI adds to the back of cc a *url.URL type
func (t *IntransitiveActivity) AppendCcIRI(v *url.URL) {
	t.cc = append(t.cc, &ccIntransitiveActivityIntermediateType{IRI: v})

}

// PrependCcIRI adds to the front of cc a *url.URL type
func (t *IntransitiveActivity) PrependCcIRI(v *url.URL) {
	t.cc = append([]*ccIntransitiveActivityIntermediateType{&ccIntransitiveActivityIntermediateType{IRI: v}}, t.cc...)

}

// RemoveCcIRI deletes the value from the specified index
func (t *IntransitiveActivity) RemoveCcIRI(index int) {
	copy(t.cc[index:], t.cc[index+1:])
	t.cc[len(t.cc)-1] = nil
	t.cc = t.cc[:len(t.cc)-1]

}

// HasUnknownCc determines whether the call to GetUnknownCc is safe
func (t *IntransitiveActivity) HasUnknownCc() (ok bool) {
	return t.cc != nil && t.cc[0].unknown_ != nil

}

// GetUnknownCc returns the unknown value for cc
func (t *IntransitiveActivity) GetUnknownCc() (v interface{}) {
	return t.cc[0].unknown_

}

// SetUnknownCc sets the unknown value of cc
func (t *IntransitiveActivity) SetUnknownCc(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &ccIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.cc = append(t.cc, tmp)

}

// BccLen determines the number of elements able to be used for the IsBccObject, GetBccObject, and RemoveBccObject functions
func (t *IntransitiveActivity) BccLen() (l int) {
	return len(t.bcc)

}

// IsBccObject determines whether the call to GetBccObject is safe for the specified index
func (t *IntransitiveActivity) IsBccObject(index int) (ok bool) {
	return t.bcc[index].Object != nil

}

// GetBccObject returns the value safely if IsBccObject returned true for the specified index
func (t *IntransitiveActivity) GetBccObject(index int) (v ObjectType) {
	return t.bcc[index].Object

}

// AppendBccObject adds to the back of bcc a ObjectType type
func (t *IntransitiveActivity) AppendBccObject(v ObjectType) {
	t.bcc = append(t.bcc, &bccIntransitiveActivityIntermediateType{Object: v})

}

// PrependBccObject adds to the front of bcc a ObjectType type
func (t *IntransitiveActivity) PrependBccObject(v ObjectType) {
	t.bcc = append([]*bccIntransitiveActivityIntermediateType{&bccIntransitiveActivityIntermediateType{Object: v}}, t.bcc...)

}

// RemoveBccObject deletes the value from the specified index
func (t *IntransitiveActivity) RemoveBccObject(index int) {
	copy(t.bcc[index:], t.bcc[index+1:])
	t.bcc[len(t.bcc)-1] = nil
	t.bcc = t.bcc[:len(t.bcc)-1]

}

// IsBccLink determines whether the call to GetBccLink is safe for the specified index
func (t *IntransitiveActivity) IsBccLink(index int) (ok bool) {
	return t.bcc[index].Link != nil

}

// GetBccLink returns the value safely if IsBccLink returned true for the specified index
func (t *IntransitiveActivity) GetBccLink(index int) (v LinkType) {
	return t.bcc[index].Link

}

// AppendBccLink adds to the back of bcc a LinkType type
func (t *IntransitiveActivity) AppendBccLink(v LinkType) {
	t.bcc = append(t.bcc, &bccIntransitiveActivityIntermediateType{Link: v})

}

// PrependBccLink adds to the front of bcc a LinkType type
func (t *IntransitiveActivity) PrependBccLink(v LinkType) {
	t.bcc = append([]*bccIntransitiveActivityIntermediateType{&bccIntransitiveActivityIntermediateType{Link: v}}, t.bcc...)

}

// RemoveBccLink deletes the value from the specified index
func (t *IntransitiveActivity) RemoveBccLink(index int) {
	copy(t.bcc[index:], t.bcc[index+1:])
	t.bcc[len(t.bcc)-1] = nil
	t.bcc = t.bcc[:len(t.bcc)-1]

}

// IsBccIRI determines whether the call to GetBccIRI is safe for the specified index
func (t *IntransitiveActivity) IsBccIRI(index int) (ok bool) {
	return t.bcc[index].IRI != nil

}

// GetBccIRI returns the value safely if IsBccIRI returned true for the specified index
func (t *IntransitiveActivity) GetBccIRI(index int) (v *url.URL) {
	return t.bcc[index].IRI

}

// AppendBccIRI adds to the back of bcc a *url.URL type
func (t *IntransitiveActivity) AppendBccIRI(v *url.URL) {
	t.bcc = append(t.bcc, &bccIntransitiveActivityIntermediateType{IRI: v})

}

// PrependBccIRI adds to the front of bcc a *url.URL type
func (t *IntransitiveActivity) PrependBccIRI(v *url.URL) {
	t.bcc = append([]*bccIntransitiveActivityIntermediateType{&bccIntransitiveActivityIntermediateType{IRI: v}}, t.bcc...)

}

// RemoveBccIRI deletes the value from the specified index
func (t *IntransitiveActivity) RemoveBccIRI(index int) {
	copy(t.bcc[index:], t.bcc[index+1:])
	t.bcc[len(t.bcc)-1] = nil
	t.bcc = t.bcc[:len(t.bcc)-1]

}

// HasUnknownBcc determines whether the call to GetUnknownBcc is safe
func (t *IntransitiveActivity) HasUnknownBcc() (ok bool) {
	return t.bcc != nil && t.bcc[0].unknown_ != nil

}

// GetUnknownBcc returns the unknown value for bcc
func (t *IntransitiveActivity) GetUnknownBcc() (v interface{}) {
	return t.bcc[0].unknown_

}

// SetUnknownBcc sets the unknown value of bcc
func (t *IntransitiveActivity) SetUnknownBcc(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &bccIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.bcc = append(t.bcc, tmp)

}

// IsMediaType determines whether the call to GetMediaType is safe
func (t *IntransitiveActivity) IsMediaType() (ok bool) {
	return t.mediaType != nil && t.mediaType.mimeMediaTypeValue != nil

}

// GetMediaType returns the value safely if IsMediaType returned true
func (t *IntransitiveActivity) GetMediaType() (v string) {
	return *t.mediaType.mimeMediaTypeValue

}

// SetMediaType sets the value of mediaType to be of string type
func (t *IntransitiveActivity) SetMediaType(v string) {
	t.mediaType = &mediaTypeIntransitiveActivityIntermediateType{mimeMediaTypeValue: &v}

}

// IsMediaTypeIRI determines whether the call to GetMediaTypeIRI is safe
func (t *IntransitiveActivity) IsMediaTypeIRI() (ok bool) {
	return t.mediaType != nil && t.mediaType.IRI != nil

}

// GetMediaTypeIRI returns the value safely if IsMediaTypeIRI returned true
func (t *IntransitiveActivity) GetMediaTypeIRI() (v *url.URL) {
	return t.mediaType.IRI

}

// SetMediaTypeIRI sets the value of mediaType to be of *url.URL type
func (t *IntransitiveActivity) SetMediaTypeIRI(v *url.URL) {
	t.mediaType = &mediaTypeIntransitiveActivityIntermediateType{IRI: v}

}

// HasUnknownMediaType determines whether the call to GetUnknownMediaType is safe
func (t *IntransitiveActivity) HasUnknownMediaType() (ok bool) {
	return t.mediaType != nil && t.mediaType.unknown_ != nil

}

// GetUnknownMediaType returns the unknown value for mediaType
func (t *IntransitiveActivity) GetUnknownMediaType() (v interface{}) {
	return t.mediaType.unknown_

}

// SetUnknownMediaType sets the unknown value of mediaType
func (t *IntransitiveActivity) SetUnknownMediaType(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &mediaTypeIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.mediaType = tmp

}

// IsDuration determines whether the call to GetDuration is safe
func (t *IntransitiveActivity) IsDuration() (ok bool) {
	return t.duration != nil && t.duration.duration != nil

}

// GetDuration returns the value safely if IsDuration returned true
func (t *IntransitiveActivity) GetDuration() (v time.Duration) {
	return *t.duration.duration

}

// SetDuration sets the value of duration to be of time.Duration type
func (t *IntransitiveActivity) SetDuration(v time.Duration) {
	t.duration = &durationIntransitiveActivityIntermediateType{duration: &v}

}

// IsDurationIRI determines whether the call to GetDurationIRI is safe
func (t *IntransitiveActivity) IsDurationIRI() (ok bool) {
	return t.duration != nil && t.duration.IRI != nil

}

// GetDurationIRI returns the value safely if IsDurationIRI returned true
func (t *IntransitiveActivity) GetDurationIRI() (v *url.URL) {
	return t.duration.IRI

}

// SetDurationIRI sets the value of duration to be of *url.URL type
func (t *IntransitiveActivity) SetDurationIRI(v *url.URL) {
	t.duration = &durationIntransitiveActivityIntermediateType{IRI: v}

}

// HasUnknownDuration determines whether the call to GetUnknownDuration is safe
func (t *IntransitiveActivity) HasUnknownDuration() (ok bool) {
	return t.duration != nil && t.duration.unknown_ != nil

}

// GetUnknownDuration returns the unknown value for duration
func (t *IntransitiveActivity) GetUnknownDuration() (v interface{}) {
	return t.duration.unknown_

}

// SetUnknownDuration sets the unknown value of duration
func (t *IntransitiveActivity) SetUnknownDuration(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &durationIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.duration = tmp

}

// IsSource determines whether the call to GetSource is safe
func (t *IntransitiveActivity) IsSource() (ok bool) {
	return t.source != nil && t.source.Object != nil

}

// GetSource returns the value safely if IsSource returned true
func (t *IntransitiveActivity) GetSource() (v ObjectType) {
	return t.source.Object

}

// SetSource sets the value of source to be of ObjectType type
func (t *IntransitiveActivity) SetSource(v ObjectType) {
	t.source = &sourceIntransitiveActivityIntermediateType{Object: v}

}

// IsSourceIRI determines whether the call to GetSourceIRI is safe
func (t *IntransitiveActivity) IsSourceIRI() (ok bool) {
	return t.source != nil && t.source.IRI != nil

}

// GetSourceIRI returns the value safely if IsSourceIRI returned true
func (t *IntransitiveActivity) GetSourceIRI() (v *url.URL) {
	return t.source.IRI

}

// SetSourceIRI sets the value of source to be of *url.URL type
func (t *IntransitiveActivity) SetSourceIRI(v *url.URL) {
	t.source = &sourceIntransitiveActivityIntermediateType{IRI: v}

}

// HasUnknownSource determines whether the call to GetUnknownSource is safe
func (t *IntransitiveActivity) HasUnknownSource() (ok bool) {
	return t.source != nil && t.source.unknown_ != nil

}

// GetUnknownSource returns the unknown value for source
func (t *IntransitiveActivity) GetUnknownSource() (v interface{}) {
	return t.source.unknown_

}

// SetUnknownSource sets the unknown value of source
func (t *IntransitiveActivity) SetUnknownSource(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &sourceIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.source = tmp

}

// IsInboxOrderedCollection determines whether the call to GetInboxOrderedCollection is safe
func (t *IntransitiveActivity) IsInboxOrderedCollection() (ok bool) {
	return t.inbox != nil && t.inbox.OrderedCollection != nil

}

// GetInboxOrderedCollection returns the value safely if IsInboxOrderedCollection returned true
func (t *IntransitiveActivity) GetInboxOrderedCollection() (v OrderedCollectionType) {
	return t.inbox.OrderedCollection

}

// SetInboxOrderedCollection sets the value of inbox to be of OrderedCollectionType type
func (t *IntransitiveActivity) SetInboxOrderedCollection(v OrderedCollectionType) {
	t.inbox = &inboxIntransitiveActivityIntermediateType{OrderedCollection: v}

}

// IsInboxAnyURI determines whether the call to GetInboxAnyURI is safe
func (t *IntransitiveActivity) IsInboxAnyURI() (ok bool) {
	return t.inbox != nil && t.inbox.anyURI != nil

}

// GetInboxAnyURI returns the value safely if IsInboxAnyURI returned true
func (t *IntransitiveActivity) GetInboxAnyURI() (v *url.URL) {
	return t.inbox.anyURI

}

// SetInboxAnyURI sets the value of inbox to be of *url.URL type
func (t *IntransitiveActivity) SetInboxAnyURI(v *url.URL) {
	t.inbox = &inboxIntransitiveActivityIntermediateType{anyURI: v}

}

// HasUnknownInbox determines whether the call to GetUnknownInbox is safe
func (t *IntransitiveActivity) HasUnknownInbox() (ok bool) {
	return t.inbox != nil && t.inbox.unknown_ != nil

}

// GetUnknownInbox returns the unknown value for inbox
func (t *IntransitiveActivity) GetUnknownInbox() (v interface{}) {
	return t.inbox.unknown_

}

// SetUnknownInbox sets the unknown value of inbox
func (t *IntransitiveActivity) SetUnknownInbox(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &inboxIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.inbox = tmp

}

// IsOutboxOrderedCollection determines whether the call to GetOutboxOrderedCollection is safe
func (t *IntransitiveActivity) IsOutboxOrderedCollection() (ok bool) {
	return t.outbox != nil && t.outbox.OrderedCollection != nil

}

// GetOutboxOrderedCollection returns the value safely if IsOutboxOrderedCollection returned true
func (t *IntransitiveActivity) GetOutboxOrderedCollection() (v OrderedCollectionType) {
	return t.outbox.OrderedCollection

}

// SetOutboxOrderedCollection sets the value of outbox to be of OrderedCollectionType type
func (t *IntransitiveActivity) SetOutboxOrderedCollection(v OrderedCollectionType) {
	t.outbox = &outboxIntransitiveActivityIntermediateType{OrderedCollection: v}

}

// IsOutboxAnyURI determines whether the call to GetOutboxAnyURI is safe
func (t *IntransitiveActivity) IsOutboxAnyURI() (ok bool) {
	return t.outbox != nil && t.outbox.anyURI != nil

}

// GetOutboxAnyURI returns the value safely if IsOutboxAnyURI returned true
func (t *IntransitiveActivity) GetOutboxAnyURI() (v *url.URL) {
	return t.outbox.anyURI

}

// SetOutboxAnyURI sets the value of outbox to be of *url.URL type
func (t *IntransitiveActivity) SetOutboxAnyURI(v *url.URL) {
	t.outbox = &outboxIntransitiveActivityIntermediateType{anyURI: v}

}

// HasUnknownOutbox determines whether the call to GetUnknownOutbox is safe
func (t *IntransitiveActivity) HasUnknownOutbox() (ok bool) {
	return t.outbox != nil && t.outbox.unknown_ != nil

}

// GetUnknownOutbox returns the unknown value for outbox
func (t *IntransitiveActivity) GetUnknownOutbox() (v interface{}) {
	return t.outbox.unknown_

}

// SetUnknownOutbox sets the unknown value of outbox
func (t *IntransitiveActivity) SetUnknownOutbox(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &outboxIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.outbox = tmp

}

// IsFollowingCollection determines whether the call to GetFollowingCollection is safe
func (t *IntransitiveActivity) IsFollowingCollection() (ok bool) {
	return t.following != nil && t.following.Collection != nil

}

// GetFollowingCollection returns the value safely if IsFollowingCollection returned true
func (t *IntransitiveActivity) GetFollowingCollection() (v CollectionType) {
	return t.following.Collection

}

// SetFollowingCollection sets the value of following to be of CollectionType type
func (t *IntransitiveActivity) SetFollowingCollection(v CollectionType) {
	t.following = &followingIntransitiveActivityIntermediateType{Collection: v}

}

// IsFollowingOrderedCollection determines whether the call to GetFollowingOrderedCollection is safe
func (t *IntransitiveActivity) IsFollowingOrderedCollection() (ok bool) {
	return t.following != nil && t.following.OrderedCollection != nil

}

// GetFollowingOrderedCollection returns the value safely if IsFollowingOrderedCollection returned true
func (t *IntransitiveActivity) GetFollowingOrderedCollection() (v OrderedCollectionType) {
	return t.following.OrderedCollection

}

// SetFollowingOrderedCollection sets the value of following to be of OrderedCollectionType type
func (t *IntransitiveActivity) SetFollowingOrderedCollection(v OrderedCollectionType) {
	t.following = &followingIntransitiveActivityIntermediateType{OrderedCollection: v}

}

// IsFollowingAnyURI determines whether the call to GetFollowingAnyURI is safe
func (t *IntransitiveActivity) IsFollowingAnyURI() (ok bool) {
	return t.following != nil && t.following.anyURI != nil

}

// GetFollowingAnyURI returns the value safely if IsFollowingAnyURI returned true
func (t *IntransitiveActivity) GetFollowingAnyURI() (v *url.URL) {
	return t.following.anyURI

}

// SetFollowingAnyURI sets the value of following to be of *url.URL type
func (t *IntransitiveActivity) SetFollowingAnyURI(v *url.URL) {
	t.following = &followingIntransitiveActivityIntermediateType{anyURI: v}

}

// HasUnknownFollowing determines whether the call to GetUnknownFollowing is safe
func (t *IntransitiveActivity) HasUnknownFollowing() (ok bool) {
	return t.following != nil && t.following.unknown_ != nil

}

// GetUnknownFollowing returns the unknown value for following
func (t *IntransitiveActivity) GetUnknownFollowing() (v interface{}) {
	return t.following.unknown_

}

// SetUnknownFollowing sets the unknown value of following
func (t *IntransitiveActivity) SetUnknownFollowing(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &followingIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.following = tmp

}

// IsFollowersCollection determines whether the call to GetFollowersCollection is safe
func (t *IntransitiveActivity) IsFollowersCollection() (ok bool) {
	return t.followers != nil && t.followers.Collection != nil

}

// GetFollowersCollection returns the value safely if IsFollowersCollection returned true
func (t *IntransitiveActivity) GetFollowersCollection() (v CollectionType) {
	return t.followers.Collection

}

// SetFollowersCollection sets the value of followers to be of CollectionType type
func (t *IntransitiveActivity) SetFollowersCollection(v CollectionType) {
	t.followers = &followersIntransitiveActivityIntermediateType{Collection: v}

}

// IsFollowersOrderedCollection determines whether the call to GetFollowersOrderedCollection is safe
func (t *IntransitiveActivity) IsFollowersOrderedCollection() (ok bool) {
	return t.followers != nil && t.followers.OrderedCollection != nil

}

// GetFollowersOrderedCollection returns the value safely if IsFollowersOrderedCollection returned true
func (t *IntransitiveActivity) GetFollowersOrderedCollection() (v OrderedCollectionType) {
	return t.followers.OrderedCollection

}

// SetFollowersOrderedCollection sets the value of followers to be of OrderedCollectionType type
func (t *IntransitiveActivity) SetFollowersOrderedCollection(v OrderedCollectionType) {
	t.followers = &followersIntransitiveActivityIntermediateType{OrderedCollection: v}

}

// IsFollowersAnyURI determines whether the call to GetFollowersAnyURI is safe
func (t *IntransitiveActivity) IsFollowersAnyURI() (ok bool) {
	return t.followers != nil && t.followers.anyURI != nil

}

// GetFollowersAnyURI returns the value safely if IsFollowersAnyURI returned true
func (t *IntransitiveActivity) GetFollowersAnyURI() (v *url.URL) {
	return t.followers.anyURI

}

// SetFollowersAnyURI sets the value of followers to be of *url.URL type
func (t *IntransitiveActivity) SetFollowersAnyURI(v *url.URL) {
	t.followers = &followersIntransitiveActivityIntermediateType{anyURI: v}

}

// HasUnknownFollowers determines whether the call to GetUnknownFollowers is safe
func (t *IntransitiveActivity) HasUnknownFollowers() (ok bool) {
	return t.followers != nil && t.followers.unknown_ != nil

}

// GetUnknownFollowers returns the unknown value for followers
func (t *IntransitiveActivity) GetUnknownFollowers() (v interface{}) {
	return t.followers.unknown_

}

// SetUnknownFollowers sets the unknown value of followers
func (t *IntransitiveActivity) SetUnknownFollowers(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &followersIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.followers = tmp

}

// IsLikedCollection determines whether the call to GetLikedCollection is safe
func (t *IntransitiveActivity) IsLikedCollection() (ok bool) {
	return t.liked != nil && t.liked.Collection != nil

}

// GetLikedCollection returns the value safely if IsLikedCollection returned true
func (t *IntransitiveActivity) GetLikedCollection() (v CollectionType) {
	return t.liked.Collection

}

// SetLikedCollection sets the value of liked to be of CollectionType type
func (t *IntransitiveActivity) SetLikedCollection(v CollectionType) {
	t.liked = &likedIntransitiveActivityIntermediateType{Collection: v}

}

// IsLikedOrderedCollection determines whether the call to GetLikedOrderedCollection is safe
func (t *IntransitiveActivity) IsLikedOrderedCollection() (ok bool) {
	return t.liked != nil && t.liked.OrderedCollection != nil

}

// GetLikedOrderedCollection returns the value safely if IsLikedOrderedCollection returned true
func (t *IntransitiveActivity) GetLikedOrderedCollection() (v OrderedCollectionType) {
	return t.liked.OrderedCollection

}

// SetLikedOrderedCollection sets the value of liked to be of OrderedCollectionType type
func (t *IntransitiveActivity) SetLikedOrderedCollection(v OrderedCollectionType) {
	t.liked = &likedIntransitiveActivityIntermediateType{OrderedCollection: v}

}

// IsLikedAnyURI determines whether the call to GetLikedAnyURI is safe
func (t *IntransitiveActivity) IsLikedAnyURI() (ok bool) {
	return t.liked != nil && t.liked.anyURI != nil

}

// GetLikedAnyURI returns the value safely if IsLikedAnyURI returned true
func (t *IntransitiveActivity) GetLikedAnyURI() (v *url.URL) {
	return t.liked.anyURI

}

// SetLikedAnyURI sets the value of liked to be of *url.URL type
func (t *IntransitiveActivity) SetLikedAnyURI(v *url.URL) {
	t.liked = &likedIntransitiveActivityIntermediateType{anyURI: v}

}

// HasUnknownLiked determines whether the call to GetUnknownLiked is safe
func (t *IntransitiveActivity) HasUnknownLiked() (ok bool) {
	return t.liked != nil && t.liked.unknown_ != nil

}

// GetUnknownLiked returns the unknown value for liked
func (t *IntransitiveActivity) GetUnknownLiked() (v interface{}) {
	return t.liked.unknown_

}

// SetUnknownLiked sets the unknown value of liked
func (t *IntransitiveActivity) SetUnknownLiked(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &likedIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.liked = tmp

}

// IsLikesCollection determines whether the call to GetLikesCollection is safe
func (t *IntransitiveActivity) IsLikesCollection() (ok bool) {
	return t.likes != nil && t.likes.Collection != nil

}

// GetLikesCollection returns the value safely if IsLikesCollection returned true
func (t *IntransitiveActivity) GetLikesCollection() (v CollectionType) {
	return t.likes.Collection

}

// SetLikesCollection sets the value of likes to be of CollectionType type
func (t *IntransitiveActivity) SetLikesCollection(v CollectionType) {
	t.likes = &likesIntransitiveActivityIntermediateType{Collection: v}

}

// IsLikesOrderedCollection determines whether the call to GetLikesOrderedCollection is safe
func (t *IntransitiveActivity) IsLikesOrderedCollection() (ok bool) {
	return t.likes != nil && t.likes.OrderedCollection != nil

}

// GetLikesOrderedCollection returns the value safely if IsLikesOrderedCollection returned true
func (t *IntransitiveActivity) GetLikesOrderedCollection() (v OrderedCollectionType) {
	return t.likes.OrderedCollection

}

// SetLikesOrderedCollection sets the value of likes to be of OrderedCollectionType type
func (t *IntransitiveActivity) SetLikesOrderedCollection(v OrderedCollectionType) {
	t.likes = &likesIntransitiveActivityIntermediateType{OrderedCollection: v}

}

// IsLikesAnyURI determines whether the call to GetLikesAnyURI is safe
func (t *IntransitiveActivity) IsLikesAnyURI() (ok bool) {
	return t.likes != nil && t.likes.anyURI != nil

}

// GetLikesAnyURI returns the value safely if IsLikesAnyURI returned true
func (t *IntransitiveActivity) GetLikesAnyURI() (v *url.URL) {
	return t.likes.anyURI

}

// SetLikesAnyURI sets the value of likes to be of *url.URL type
func (t *IntransitiveActivity) SetLikesAnyURI(v *url.URL) {
	t.likes = &likesIntransitiveActivityIntermediateType{anyURI: v}

}

// HasUnknownLikes determines whether the call to GetUnknownLikes is safe
func (t *IntransitiveActivity) HasUnknownLikes() (ok bool) {
	return t.likes != nil && t.likes.unknown_ != nil

}

// GetUnknownLikes returns the unknown value for likes
func (t *IntransitiveActivity) GetUnknownLikes() (v interface{}) {
	return t.likes.unknown_

}

// SetUnknownLikes sets the unknown value of likes
func (t *IntransitiveActivity) SetUnknownLikes(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &likesIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.likes = tmp

}

// StreamsLen determines the number of elements able to be used for the GetStreams and RemoveStreams functions
func (t *IntransitiveActivity) StreamsLen() (l int) {
	return len(t.streams)

}

// GetStreams returns the value for the specified index
func (t *IntransitiveActivity) GetStreams(index int) (v *url.URL) {
	return t.streams[index]

}

// AppendStreams adds a value to the back of streams
func (t *IntransitiveActivity) AppendStreams(v *url.URL) {
	t.streams = append(t.streams, v)

}

// PrependStreams adds a value to the front of streams
func (t *IntransitiveActivity) PrependStreams(v *url.URL) {
	t.streams = append([]*url.URL{v}, t.streams...)

}

// RemoveStreams deletes the value from the specified index
func (t *IntransitiveActivity) RemoveStreams(index int) {
	copy(t.streams[index:], t.streams[index+1:])
	t.streams[len(t.streams)-1] = nil
	t.streams = t.streams[:len(t.streams)-1]

}

// HasUnknownStreams determines whether the call to GetUnknownStreams is safe
func (t *IntransitiveActivity) HasUnknownStreams() (ok bool) {
	return t.unknown_ != nil && t.unknown_["streams"] != nil

}

// GetUnknownStreams returns the unknown value for streams
func (t *IntransitiveActivity) GetUnknownStreams() (v interface{}) {
	return t.unknown_["streams"]

}

// SetUnknownStreams sets the unknown value of streams
func (t *IntransitiveActivity) SetUnknownStreams(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["streams"] = i

}

// IsPreferredUsername determines whether the call to GetPreferredUsername is safe
func (t *IntransitiveActivity) IsPreferredUsername() (ok bool) {
	return t.preferredUsername != nil && t.preferredUsername.stringName != nil

}

// GetPreferredUsername returns the value safely if IsPreferredUsername returned true
func (t *IntransitiveActivity) GetPreferredUsername() (v string) {
	return *t.preferredUsername.stringName

}

// SetPreferredUsername sets the value of preferredUsername to be of string type
func (t *IntransitiveActivity) SetPreferredUsername(v string) {
	t.preferredUsername = &preferredUsernameIntransitiveActivityIntermediateType{stringName: &v}

}

// IsPreferredUsernameIRI determines whether the call to GetPreferredUsernameIRI is safe
func (t *IntransitiveActivity) IsPreferredUsernameIRI() (ok bool) {
	return t.preferredUsername != nil && t.preferredUsername.IRI != nil

}

// GetPreferredUsernameIRI returns the value safely if IsPreferredUsernameIRI returned true
func (t *IntransitiveActivity) GetPreferredUsernameIRI() (v *url.URL) {
	return t.preferredUsername.IRI

}

// SetPreferredUsernameIRI sets the value of preferredUsername to be of *url.URL type
func (t *IntransitiveActivity) SetPreferredUsernameIRI(v *url.URL) {
	t.preferredUsername = &preferredUsernameIntransitiveActivityIntermediateType{IRI: v}

}

// HasUnknownPreferredUsername determines whether the call to GetUnknownPreferredUsername is safe
func (t *IntransitiveActivity) HasUnknownPreferredUsername() (ok bool) {
	return t.preferredUsername != nil && t.preferredUsername.unknown_ != nil

}

// GetUnknownPreferredUsername returns the unknown value for preferredUsername
func (t *IntransitiveActivity) GetUnknownPreferredUsername() (v interface{}) {
	return t.preferredUsername.unknown_

}

// SetUnknownPreferredUsername sets the unknown value of preferredUsername
func (t *IntransitiveActivity) SetUnknownPreferredUsername(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &preferredUsernameIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.preferredUsername = tmp

}

// PreferredUsernameMapLanguages returns all languages for this property's language mapping, or nil if there are none.
func (t *IntransitiveActivity) PreferredUsernameMapLanguages() (l []string) {
	if t.preferredUsernameMap == nil || len(t.preferredUsernameMap) == 0 {
		return nil
	}
	for k := range t.preferredUsernameMap {
		l = append(l, k)
	}
	return

}

// GetPreferredUsernameMap retrieves the value of the property for the specified language, or an empty string if it does not exist
func (t *IntransitiveActivity) GetPreferredUsernameMap(l string) (v string) {
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
func (t *IntransitiveActivity) SetPreferredUsernameMap(l string, v string) {
	if t.preferredUsernameMap == nil {
		t.preferredUsernameMap = make(map[string]string)
	}
	t.preferredUsernameMap[l] = v

}

// IsEndpoints determines whether the call to GetEndpoints is safe
func (t *IntransitiveActivity) IsEndpoints() (ok bool) {
	return t.endpoints != nil && t.endpoints.Object != nil

}

// GetEndpoints returns the value safely if IsEndpoints returned true
func (t *IntransitiveActivity) GetEndpoints() (v ObjectType) {
	return t.endpoints.Object

}

// SetEndpoints sets the value of endpoints to be of ObjectType type
func (t *IntransitiveActivity) SetEndpoints(v ObjectType) {
	t.endpoints = &endpointsIntransitiveActivityIntermediateType{Object: v}

}

// IsEndpointsIRI determines whether the call to GetEndpointsIRI is safe
func (t *IntransitiveActivity) IsEndpointsIRI() (ok bool) {
	return t.endpoints != nil && t.endpoints.IRI != nil

}

// GetEndpointsIRI returns the value safely if IsEndpointsIRI returned true
func (t *IntransitiveActivity) GetEndpointsIRI() (v *url.URL) {
	return t.endpoints.IRI

}

// SetEndpointsIRI sets the value of endpoints to be of *url.URL type
func (t *IntransitiveActivity) SetEndpointsIRI(v *url.URL) {
	t.endpoints = &endpointsIntransitiveActivityIntermediateType{IRI: v}

}

// HasUnknownEndpoints determines whether the call to GetUnknownEndpoints is safe
func (t *IntransitiveActivity) HasUnknownEndpoints() (ok bool) {
	return t.endpoints != nil && t.endpoints.unknown_ != nil

}

// GetUnknownEndpoints returns the unknown value for endpoints
func (t *IntransitiveActivity) GetUnknownEndpoints() (v interface{}) {
	return t.endpoints.unknown_

}

// SetUnknownEndpoints sets the unknown value of endpoints
func (t *IntransitiveActivity) SetUnknownEndpoints(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &endpointsIntransitiveActivityIntermediateType{}
	tmp.unknown_ = i
	t.endpoints = tmp

}

// HasProxyUrl determines whether the call to GetProxyUrl is safe
func (t *IntransitiveActivity) HasProxyUrl() (ok bool) {
	return t.proxyUrl != nil

}

// GetProxyUrl returns the value for proxyUrl
func (t *IntransitiveActivity) GetProxyUrl() (v *url.URL) {
	return t.proxyUrl

}

// SetProxyUrl sets the value of proxyUrl
func (t *IntransitiveActivity) SetProxyUrl(v *url.URL) {
	t.proxyUrl = v

}

// HasUnknownProxyUrl determines whether the call to GetUnknownProxyUrl is safe
func (t *IntransitiveActivity) HasUnknownProxyUrl() (ok bool) {
	return t.unknown_ != nil && t.unknown_["proxyUrl"] != nil

}

// GetUnknownProxyUrl returns the unknown value for proxyUrl
func (t *IntransitiveActivity) GetUnknownProxyUrl() (v interface{}) {
	return t.unknown_["proxyUrl"]

}

// SetUnknownProxyUrl sets the unknown value of proxyUrl
func (t *IntransitiveActivity) SetUnknownProxyUrl(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["proxyUrl"] = i

}

// HasOauthAuthorizationEndpoint determines whether the call to GetOauthAuthorizationEndpoint is safe
func (t *IntransitiveActivity) HasOauthAuthorizationEndpoint() (ok bool) {
	return t.oauthAuthorizationEndpoint != nil

}

// GetOauthAuthorizationEndpoint returns the value for oauthAuthorizationEndpoint
func (t *IntransitiveActivity) GetOauthAuthorizationEndpoint() (v *url.URL) {
	return t.oauthAuthorizationEndpoint

}

// SetOauthAuthorizationEndpoint sets the value of oauthAuthorizationEndpoint
func (t *IntransitiveActivity) SetOauthAuthorizationEndpoint(v *url.URL) {
	t.oauthAuthorizationEndpoint = v

}

// HasUnknownOauthAuthorizationEndpoint determines whether the call to GetUnknownOauthAuthorizationEndpoint is safe
func (t *IntransitiveActivity) HasUnknownOauthAuthorizationEndpoint() (ok bool) {
	return t.unknown_ != nil && t.unknown_["oauthAuthorizationEndpoint"] != nil

}

// GetUnknownOauthAuthorizationEndpoint returns the unknown value for oauthAuthorizationEndpoint
func (t *IntransitiveActivity) GetUnknownOauthAuthorizationEndpoint() (v interface{}) {
	return t.unknown_["oauthAuthorizationEndpoint"]

}

// SetUnknownOauthAuthorizationEndpoint sets the unknown value of oauthAuthorizationEndpoint
func (t *IntransitiveActivity) SetUnknownOauthAuthorizationEndpoint(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["oauthAuthorizationEndpoint"] = i

}

// HasOauthTokenEndpoint determines whether the call to GetOauthTokenEndpoint is safe
func (t *IntransitiveActivity) HasOauthTokenEndpoint() (ok bool) {
	return t.oauthTokenEndpoint != nil

}

// GetOauthTokenEndpoint returns the value for oauthTokenEndpoint
func (t *IntransitiveActivity) GetOauthTokenEndpoint() (v *url.URL) {
	return t.oauthTokenEndpoint

}

// SetOauthTokenEndpoint sets the value of oauthTokenEndpoint
func (t *IntransitiveActivity) SetOauthTokenEndpoint(v *url.URL) {
	t.oauthTokenEndpoint = v

}

// HasUnknownOauthTokenEndpoint determines whether the call to GetUnknownOauthTokenEndpoint is safe
func (t *IntransitiveActivity) HasUnknownOauthTokenEndpoint() (ok bool) {
	return t.unknown_ != nil && t.unknown_["oauthTokenEndpoint"] != nil

}

// GetUnknownOauthTokenEndpoint returns the unknown value for oauthTokenEndpoint
func (t *IntransitiveActivity) GetUnknownOauthTokenEndpoint() (v interface{}) {
	return t.unknown_["oauthTokenEndpoint"]

}

// SetUnknownOauthTokenEndpoint sets the unknown value of oauthTokenEndpoint
func (t *IntransitiveActivity) SetUnknownOauthTokenEndpoint(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["oauthTokenEndpoint"] = i

}

// HasProvideClientKey determines whether the call to GetProvideClientKey is safe
func (t *IntransitiveActivity) HasProvideClientKey() (ok bool) {
	return t.provideClientKey != nil

}

// GetProvideClientKey returns the value for provideClientKey
func (t *IntransitiveActivity) GetProvideClientKey() (v *url.URL) {
	return t.provideClientKey

}

// SetProvideClientKey sets the value of provideClientKey
func (t *IntransitiveActivity) SetProvideClientKey(v *url.URL) {
	t.provideClientKey = v

}

// HasUnknownProvideClientKey determines whether the call to GetUnknownProvideClientKey is safe
func (t *IntransitiveActivity) HasUnknownProvideClientKey() (ok bool) {
	return t.unknown_ != nil && t.unknown_["provideClientKey"] != nil

}

// GetUnknownProvideClientKey returns the unknown value for provideClientKey
func (t *IntransitiveActivity) GetUnknownProvideClientKey() (v interface{}) {
	return t.unknown_["provideClientKey"]

}

// SetUnknownProvideClientKey sets the unknown value of provideClientKey
func (t *IntransitiveActivity) SetUnknownProvideClientKey(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["provideClientKey"] = i

}

// HasSignClientKey determines whether the call to GetSignClientKey is safe
func (t *IntransitiveActivity) HasSignClientKey() (ok bool) {
	return t.signClientKey != nil

}

// GetSignClientKey returns the value for signClientKey
func (t *IntransitiveActivity) GetSignClientKey() (v *url.URL) {
	return t.signClientKey

}

// SetSignClientKey sets the value of signClientKey
func (t *IntransitiveActivity) SetSignClientKey(v *url.URL) {
	t.signClientKey = v

}

// HasUnknownSignClientKey determines whether the call to GetUnknownSignClientKey is safe
func (t *IntransitiveActivity) HasUnknownSignClientKey() (ok bool) {
	return t.unknown_ != nil && t.unknown_["signClientKey"] != nil

}

// GetUnknownSignClientKey returns the unknown value for signClientKey
func (t *IntransitiveActivity) GetUnknownSignClientKey() (v interface{}) {
	return t.unknown_["signClientKey"]

}

// SetUnknownSignClientKey sets the unknown value of signClientKey
func (t *IntransitiveActivity) SetUnknownSignClientKey(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["signClientKey"] = i

}

// HasSharedInbox determines whether the call to GetSharedInbox is safe
func (t *IntransitiveActivity) HasSharedInbox() (ok bool) {
	return t.sharedInbox != nil

}

// GetSharedInbox returns the value for sharedInbox
func (t *IntransitiveActivity) GetSharedInbox() (v *url.URL) {
	return t.sharedInbox

}

// SetSharedInbox sets the value of sharedInbox
func (t *IntransitiveActivity) SetSharedInbox(v *url.URL) {
	t.sharedInbox = v

}

// HasUnknownSharedInbox determines whether the call to GetUnknownSharedInbox is safe
func (t *IntransitiveActivity) HasUnknownSharedInbox() (ok bool) {
	return t.unknown_ != nil && t.unknown_["sharedInbox"] != nil

}

// GetUnknownSharedInbox returns the unknown value for sharedInbox
func (t *IntransitiveActivity) GetUnknownSharedInbox() (v interface{}) {
	return t.unknown_["sharedInbox"]

}

// SetUnknownSharedInbox sets the unknown value of sharedInbox
func (t *IntransitiveActivity) SetUnknownSharedInbox(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["sharedInbox"] = i

}

// ObjectLen is NOT a valid property for this type; calling its associated methods will always yield an empty-equivalent value such as false, nil, or empty string. This includes instances where it should return itself. This ugliness is a symptom of the fundamental design of the ActivityStream vocabulary as instead of 'W-is-a-X' relationships it contains the notion of 'W-is-a-X-except-for-Y'.
func (t *IntransitiveActivity) ObjectLen() (l int) {
	return 0

}

// IsObject is NOT a valid property for this type; calling its associated methods will always yield an empty-equivalent value such as false, nil, or empty string. This includes instances where it should return itself. This ugliness is a symptom of the fundamental design of the ActivityStream vocabulary as instead of 'W-is-a-X' relationships it contains the notion of 'W-is-a-X-except-for-Y'.
func (t *IntransitiveActivity) IsObject(index int) (ok bool) {
	return false

}

// GetObject is NOT a valid property for this type; calling its associated methods will always yield an empty-equivalent value such as false, nil, or empty string. This includes instances where it should return itself. This ugliness is a symptom of the fundamental design of the ActivityStream vocabulary as instead of 'W-is-a-X' relationships it contains the notion of 'W-is-a-X-except-for-Y'.
func (t *IntransitiveActivity) GetObject(index int) (v ObjectType) {
	return nil

}

// AppendObject is NOT a valid property for this type; calling its associated methods will always yield an empty-equivalent value such as false, nil, or empty string. This includes instances where it should return itself. This ugliness is a symptom of the fundamental design of the ActivityStream vocabulary as instead of 'W-is-a-X' relationships it contains the notion of 'W-is-a-X-except-for-Y'.
func (t *IntransitiveActivity) AppendObject(v ObjectType) {

}

// PrependObject is NOT a valid property for this type; calling its associated methods will always yield an empty-equivalent value such as false, nil, or empty string. This includes instances where it should return itself. This ugliness is a symptom of the fundamental design of the ActivityStream vocabulary as instead of 'W-is-a-X' relationships it contains the notion of 'W-is-a-X-except-for-Y'.
func (t *IntransitiveActivity) PrependObject(v ObjectType) {

}

// RemoveObject is NOT a valid property for this type; calling its associated methods will always yield an empty-equivalent value such as false, nil, or empty string. This includes instances where it should return itself. This ugliness is a symptom of the fundamental design of the ActivityStream vocabulary as instead of 'W-is-a-X' relationships it contains the notion of 'W-is-a-X-except-for-Y'.
func (t *IntransitiveActivity) RemoveObject(index int) {

}

// IsObjectIRI is NOT a valid property for this type; calling its associated methods will always yield an empty-equivalent value such as false, nil, or empty string. This includes instances where it should return itself. This ugliness is a symptom of the fundamental design of the ActivityStream vocabulary as instead of 'W-is-a-X' relationships it contains the notion of 'W-is-a-X-except-for-Y'.
func (t *IntransitiveActivity) IsObjectIRI(index int) (ok bool) {
	return false

}

// GetObjectIRI is NOT a valid property for this type; calling its associated methods will always yield an empty-equivalent value such as false, nil, or empty string. This includes instances where it should return itself. This ugliness is a symptom of the fundamental design of the ActivityStream vocabulary as instead of 'W-is-a-X' relationships it contains the notion of 'W-is-a-X-except-for-Y'.
func (t *IntransitiveActivity) GetObjectIRI(index int) (v *url.URL) {
	return nil

}

// AppendObjectIRI is NOT a valid property for this type; calling its associated methods will always yield an empty-equivalent value such as false, nil, or empty string. This includes instances where it should return itself. This ugliness is a symptom of the fundamental design of the ActivityStream vocabulary as instead of 'W-is-a-X' relationships it contains the notion of 'W-is-a-X-except-for-Y'.
func (t *IntransitiveActivity) AppendObjectIRI(v *url.URL) {

}

// PrependObjectIRI is NOT a valid property for this type; calling its associated methods will always yield an empty-equivalent value such as false, nil, or empty string. This includes instances where it should return itself. This ugliness is a symptom of the fundamental design of the ActivityStream vocabulary as instead of 'W-is-a-X' relationships it contains the notion of 'W-is-a-X-except-for-Y'.
func (t *IntransitiveActivity) PrependObjectIRI(v *url.URL) {

}

// RemoveObjectIRI is NOT a valid property for this type; calling its associated methods will always yield an empty-equivalent value such as false, nil, or empty string. This includes instances where it should return itself. This ugliness is a symptom of the fundamental design of the ActivityStream vocabulary as instead of 'W-is-a-X' relationships it contains the notion of 'W-is-a-X-except-for-Y'.
func (t *IntransitiveActivity) RemoveObjectIRI(index int) {

}

// HasUnknownObject is NOT a valid property for this type; calling its associated methods will always yield an empty-equivalent value such as false, nil, or empty string. This includes instances where it should return itself. This ugliness is a symptom of the fundamental design of the ActivityStream vocabulary as instead of 'W-is-a-X' relationships it contains the notion of 'W-is-a-X-except-for-Y'.
func (t *IntransitiveActivity) HasUnknownObject() (ok bool) {
	return false

}

// GetUnknownObject is NOT a valid property for this type; calling its associated methods will always yield an empty-equivalent value such as false, nil, or empty string. This includes instances where it should return itself. This ugliness is a symptom of the fundamental design of the ActivityStream vocabulary as instead of 'W-is-a-X' relationships it contains the notion of 'W-is-a-X-except-for-Y'.
func (t *IntransitiveActivity) GetUnknownObject() (v interface{}) {
	return nil

}

// SetUnknownObject is NOT a valid property for this type; calling its associated methods will always yield an empty-equivalent value such as false, nil, or empty string. This includes instances where it should return itself. This ugliness is a symptom of the fundamental design of the ActivityStream vocabulary as instead of 'W-is-a-X' relationships it contains the notion of 'W-is-a-X-except-for-Y'.
func (t *IntransitiveActivity) SetUnknownObject(i interface{}) {

}

// AddUnknown adds a raw extension to this object with the specified key
func (t *IntransitiveActivity) AddUnknown(k string, i interface{}) (this *IntransitiveActivity) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_[k] = i
	return t

}

// HasUnknown returns true if there is an unknown object with the specified key
func (t *IntransitiveActivity) HasUnknown(k string) (b bool) {
	if t.unknown_ == nil {
		return false
	}
	_, ok := t.unknown_[k]
	return ok

}

// RemoveUnknown removes a raw extension from this object with the specified key
func (t *IntransitiveActivity) RemoveUnknown(k string) (this *IntransitiveActivity) {
	delete(t.unknown_, k)
	return t

}

// Serialize turns this object into a map[string]interface{}. Note that the "type" property will automatically be populated with "IntransitiveActivity" if not manually set by the caller
func (t *IntransitiveActivity) Serialize() (m map[string]interface{}, err error) {
	m = make(map[string]interface{})
	for k, v := range t.unknown_ {
		m[k] = unknownValueSerialize(v)
	}
	var typeAlreadySet bool
	for _, k := range t.typeName {
		if ks, ok := k.(string); ok {
			if ks == "IntransitiveActivity" {
				typeAlreadySet = true
				break
			}
		}
	}
	if !typeAlreadySet {
		t.typeName = append(t.typeName, "IntransitiveActivity")
	}
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceActorIntransitiveActivityIntermediateType(t.actor); err == nil && v != nil {
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
	if v, err := serializeSliceTargetIntransitiveActivityIntermediateType(t.target); err == nil && v != nil {
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
	if v, err := serializeSliceResultIntransitiveActivityIntermediateType(t.result); err == nil && v != nil {
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
	if v, err := serializeSliceOriginIntransitiveActivityIntermediateType(t.origin); err == nil && v != nil {
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
	if v, err := serializeSliceInstrumentIntransitiveActivityIntermediateType(t.instrument); err == nil && v != nil {
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
		if v, err := serializeAltitudeIntransitiveActivityIntermediateType(t.altitude); err == nil {
			m["altitude"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceAttachmentIntransitiveActivityIntermediateType(t.attachment); err == nil && v != nil {
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
	if v, err := serializeSliceAttributedToIntransitiveActivityIntermediateType(t.attributedTo); err == nil && v != nil {
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
	if v, err := serializeSliceAudienceIntransitiveActivityIntermediateType(t.audience); err == nil && v != nil {
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
	if v, err := serializeSliceContentIntransitiveActivityIntermediateType(t.content); err == nil && v != nil {
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
	if v, err := serializeSliceContextIntransitiveActivityIntermediateType(t.context); err == nil && v != nil {
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
	if v, err := serializeSliceNameIntransitiveActivityIntermediateType(t.name); err == nil && v != nil {
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
		if v, err := serializeEndTimeIntransitiveActivityIntermediateType(t.endTime); err == nil {
			m["endTime"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceGeneratorIntransitiveActivityIntermediateType(t.generator); err == nil && v != nil {
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
	if v, err := serializeSliceIconIntransitiveActivityIntermediateType(t.icon); err == nil && v != nil {
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
	if v, err := serializeSliceImageIntransitiveActivityIntermediateType(t.image); err == nil && v != nil {
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
	if v, err := serializeSliceInReplyToIntransitiveActivityIntermediateType(t.inReplyTo); err == nil && v != nil {
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
	if v, err := serializeSliceLocationIntransitiveActivityIntermediateType(t.location); err == nil && v != nil {
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
	if v, err := serializeSlicePreviewIntransitiveActivityIntermediateType(t.preview); err == nil && v != nil {
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
		if v, err := serializePublishedIntransitiveActivityIntermediateType(t.published); err == nil {
			m["published"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.replies != nil {
		if v, err := serializeRepliesIntransitiveActivityIntermediateType(t.replies); err == nil {
			m["replies"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.startTime != nil {
		if v, err := serializeStartTimeIntransitiveActivityIntermediateType(t.startTime); err == nil {
			m["startTime"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceSummaryIntransitiveActivityIntermediateType(t.summary); err == nil && v != nil {
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
	if v, err := serializeSliceTagIntransitiveActivityIntermediateType(t.tag); err == nil && v != nil {
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
		if v, err := serializeUpdatedIntransitiveActivityIntermediateType(t.updated); err == nil {
			m["updated"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceUrlIntransitiveActivityIntermediateType(t.url); err == nil && v != nil {
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
	if v, err := serializeSliceToIntransitiveActivityIntermediateType(t.to); err == nil && v != nil {
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
	if v, err := serializeSliceBtoIntransitiveActivityIntermediateType(t.bto); err == nil && v != nil {
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
	if v, err := serializeSliceCcIntransitiveActivityIntermediateType(t.cc); err == nil && v != nil {
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
	if v, err := serializeSliceBccIntransitiveActivityIntermediateType(t.bcc); err == nil && v != nil {
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
		if v, err := serializeMediaTypeIntransitiveActivityIntermediateType(t.mediaType); err == nil {
			m["mediaType"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.duration != nil {
		if v, err := serializeDurationIntransitiveActivityIntermediateType(t.duration); err == nil {
			m["duration"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.source != nil {
		if v, err := serializeSourceIntransitiveActivityIntermediateType(t.source); err == nil {
			m["source"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.inbox != nil {
		if v, err := serializeInboxIntransitiveActivityIntermediateType(t.inbox); err == nil {
			m["inbox"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.outbox != nil {
		if v, err := serializeOutboxIntransitiveActivityIntermediateType(t.outbox); err == nil {
			m["outbox"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.following != nil {
		if v, err := serializeFollowingIntransitiveActivityIntermediateType(t.following); err == nil {
			m["following"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.followers != nil {
		if v, err := serializeFollowersIntransitiveActivityIntermediateType(t.followers); err == nil {
			m["followers"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.liked != nil {
		if v, err := serializeLikedIntransitiveActivityIntermediateType(t.liked); err == nil {
			m["liked"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.likes != nil {
		if v, err := serializeLikesIntransitiveActivityIntermediateType(t.likes); err == nil {
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
		if v, err := serializePreferredUsernameIntransitiveActivityIntermediateType(t.preferredUsername); err == nil {
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
		if v, err := serializeEndpointsIntransitiveActivityIntermediateType(t.endpoints); err == nil {
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
func (t *IntransitiveActivity) Deserialize(m map[string]interface{}) (err error) {
	for k, v := range m {
		handled := false
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "actor" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeActorIntransitiveActivityIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.actor = []*actorIntransitiveActivityIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.actor, err = deserializeSliceActorIntransitiveActivityIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &actorIntransitiveActivityIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.actor = []*actorIntransitiveActivityIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "target" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeTargetIntransitiveActivityIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.target = []*targetIntransitiveActivityIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.target, err = deserializeSliceTargetIntransitiveActivityIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &targetIntransitiveActivityIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.target = []*targetIntransitiveActivityIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "result" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeResultIntransitiveActivityIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.result = []*resultIntransitiveActivityIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.result, err = deserializeSliceResultIntransitiveActivityIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &resultIntransitiveActivityIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.result = []*resultIntransitiveActivityIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "origin" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeOriginIntransitiveActivityIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.origin = []*originIntransitiveActivityIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.origin, err = deserializeSliceOriginIntransitiveActivityIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &originIntransitiveActivityIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.origin = []*originIntransitiveActivityIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "instrument" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeInstrumentIntransitiveActivityIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.instrument = []*instrumentIntransitiveActivityIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.instrument, err = deserializeSliceInstrumentIntransitiveActivityIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &instrumentIntransitiveActivityIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.instrument = []*instrumentIntransitiveActivityIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "altitude" {
				t.altitude, err = deserializeAltitudeIntransitiveActivityIntermediateType(v)
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
					tmp, err := deserializeAttachmentIntransitiveActivityIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.attachment = []*attachmentIntransitiveActivityIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.attachment, err = deserializeSliceAttachmentIntransitiveActivityIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &attachmentIntransitiveActivityIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.attachment = []*attachmentIntransitiveActivityIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "attributedTo" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeAttributedToIntransitiveActivityIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.attributedTo = []*attributedToIntransitiveActivityIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.attributedTo, err = deserializeSliceAttributedToIntransitiveActivityIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &attributedToIntransitiveActivityIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.attributedTo = []*attributedToIntransitiveActivityIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "audience" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeAudienceIntransitiveActivityIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.audience = []*audienceIntransitiveActivityIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.audience, err = deserializeSliceAudienceIntransitiveActivityIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &audienceIntransitiveActivityIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.audience = []*audienceIntransitiveActivityIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "content" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeContentIntransitiveActivityIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.content = []*contentIntransitiveActivityIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.content, err = deserializeSliceContentIntransitiveActivityIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &contentIntransitiveActivityIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.content = []*contentIntransitiveActivityIntermediateType{tmp}
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
					tmp, err := deserializeContextIntransitiveActivityIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.context = []*contextIntransitiveActivityIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.context, err = deserializeSliceContextIntransitiveActivityIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &contextIntransitiveActivityIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.context = []*contextIntransitiveActivityIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "name" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeNameIntransitiveActivityIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.name = []*nameIntransitiveActivityIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.name, err = deserializeSliceNameIntransitiveActivityIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &nameIntransitiveActivityIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.name = []*nameIntransitiveActivityIntermediateType{tmp}
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
				t.endTime, err = deserializeEndTimeIntransitiveActivityIntermediateType(v)
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
					tmp, err := deserializeGeneratorIntransitiveActivityIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.generator = []*generatorIntransitiveActivityIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.generator, err = deserializeSliceGeneratorIntransitiveActivityIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &generatorIntransitiveActivityIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.generator = []*generatorIntransitiveActivityIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "icon" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeIconIntransitiveActivityIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.icon = []*iconIntransitiveActivityIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.icon, err = deserializeSliceIconIntransitiveActivityIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &iconIntransitiveActivityIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.icon = []*iconIntransitiveActivityIntermediateType{tmp}
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
					tmp, err := deserializeImageIntransitiveActivityIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.image = []*imageIntransitiveActivityIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.image, err = deserializeSliceImageIntransitiveActivityIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &imageIntransitiveActivityIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.image = []*imageIntransitiveActivityIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "inReplyTo" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeInReplyToIntransitiveActivityIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.inReplyTo = []*inReplyToIntransitiveActivityIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.inReplyTo, err = deserializeSliceInReplyToIntransitiveActivityIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &inReplyToIntransitiveActivityIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.inReplyTo = []*inReplyToIntransitiveActivityIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "location" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeLocationIntransitiveActivityIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.location = []*locationIntransitiveActivityIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.location, err = deserializeSliceLocationIntransitiveActivityIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &locationIntransitiveActivityIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.location = []*locationIntransitiveActivityIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "preview" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializePreviewIntransitiveActivityIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.preview = []*previewIntransitiveActivityIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.preview, err = deserializeSlicePreviewIntransitiveActivityIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &previewIntransitiveActivityIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.preview = []*previewIntransitiveActivityIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "published" {
				t.published, err = deserializePublishedIntransitiveActivityIntermediateType(v)
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
				t.replies, err = deserializeRepliesIntransitiveActivityIntermediateType(v)
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
				t.startTime, err = deserializeStartTimeIntransitiveActivityIntermediateType(v)
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
					tmp, err := deserializeSummaryIntransitiveActivityIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.summary = []*summaryIntransitiveActivityIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.summary, err = deserializeSliceSummaryIntransitiveActivityIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &summaryIntransitiveActivityIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.summary = []*summaryIntransitiveActivityIntermediateType{tmp}
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
					tmp, err := deserializeTagIntransitiveActivityIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.tag = []*tagIntransitiveActivityIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.tag, err = deserializeSliceTagIntransitiveActivityIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &tagIntransitiveActivityIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.tag = []*tagIntransitiveActivityIntermediateType{tmp}
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
				t.updated, err = deserializeUpdatedIntransitiveActivityIntermediateType(v)
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
					tmp, err := deserializeUrlIntransitiveActivityIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.url = []*urlIntransitiveActivityIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.url, err = deserializeSliceUrlIntransitiveActivityIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &urlIntransitiveActivityIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.url = []*urlIntransitiveActivityIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "to" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeToIntransitiveActivityIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.to = []*toIntransitiveActivityIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.to, err = deserializeSliceToIntransitiveActivityIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &toIntransitiveActivityIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.to = []*toIntransitiveActivityIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "bto" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeBtoIntransitiveActivityIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.bto = []*btoIntransitiveActivityIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.bto, err = deserializeSliceBtoIntransitiveActivityIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &btoIntransitiveActivityIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.bto = []*btoIntransitiveActivityIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "cc" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeCcIntransitiveActivityIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.cc = []*ccIntransitiveActivityIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.cc, err = deserializeSliceCcIntransitiveActivityIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &ccIntransitiveActivityIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.cc = []*ccIntransitiveActivityIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "bcc" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeBccIntransitiveActivityIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.bcc = []*bccIntransitiveActivityIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.bcc, err = deserializeSliceBccIntransitiveActivityIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &bccIntransitiveActivityIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.bcc = []*bccIntransitiveActivityIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "mediaType" {
				t.mediaType, err = deserializeMediaTypeIntransitiveActivityIntermediateType(v)
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
				t.duration, err = deserializeDurationIntransitiveActivityIntermediateType(v)
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
				t.source, err = deserializeSourceIntransitiveActivityIntermediateType(v)
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
				t.inbox, err = deserializeInboxIntransitiveActivityIntermediateType(v)
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
				t.outbox, err = deserializeOutboxIntransitiveActivityIntermediateType(v)
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
				t.following, err = deserializeFollowingIntransitiveActivityIntermediateType(v)
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
				t.followers, err = deserializeFollowersIntransitiveActivityIntermediateType(v)
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
				t.liked, err = deserializeLikedIntransitiveActivityIntermediateType(v)
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
				t.likes, err = deserializeLikesIntransitiveActivityIntermediateType(v)
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
				t.preferredUsername, err = deserializePreferredUsernameIntransitiveActivityIntermediateType(v)
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
				t.endpoints, err = deserializeEndpointsIntransitiveActivityIntermediateType(v)
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

// actorIntransitiveActivityIntermediateType will only have one of its values set at most
type actorIntransitiveActivityIntermediateType struct {
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
func (t *actorIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *actorIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// targetIntransitiveActivityIntermediateType will only have one of its values set at most
type targetIntransitiveActivityIntermediateType struct {
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
func (t *targetIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *targetIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// resultIntransitiveActivityIntermediateType will only have one of its values set at most
type resultIntransitiveActivityIntermediateType struct {
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
func (t *resultIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *resultIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// originIntransitiveActivityIntermediateType will only have one of its values set at most
type originIntransitiveActivityIntermediateType struct {
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
func (t *originIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *originIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// instrumentIntransitiveActivityIntermediateType will only have one of its values set at most
type instrumentIntransitiveActivityIntermediateType struct {
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
func (t *instrumentIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *instrumentIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// altitudeIntransitiveActivityIntermediateType will only have one of its values set at most
type altitudeIntransitiveActivityIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *float64 type for altitude property
	float *float64
	// Stores possible *url.URL type for altitude property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *altitudeIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *altitudeIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// attachmentIntransitiveActivityIntermediateType will only have one of its values set at most
type attachmentIntransitiveActivityIntermediateType struct {
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
func (t *attachmentIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *attachmentIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// attributedToIntransitiveActivityIntermediateType will only have one of its values set at most
type attributedToIntransitiveActivityIntermediateType struct {
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
func (t *attributedToIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *attributedToIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// audienceIntransitiveActivityIntermediateType will only have one of its values set at most
type audienceIntransitiveActivityIntermediateType struct {
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
func (t *audienceIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *audienceIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// contentIntransitiveActivityIntermediateType will only have one of its values set at most
type contentIntransitiveActivityIntermediateType struct {
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
func (t *contentIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *contentIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// contextIntransitiveActivityIntermediateType will only have one of its values set at most
type contextIntransitiveActivityIntermediateType struct {
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
func (t *contextIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *contextIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// nameIntransitiveActivityIntermediateType will only have one of its values set at most
type nameIntransitiveActivityIntermediateType struct {
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
func (t *nameIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *nameIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// endTimeIntransitiveActivityIntermediateType will only have one of its values set at most
type endTimeIntransitiveActivityIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *time.Time type for endTime property
	dateTime *time.Time
	// Stores possible *url.URL type for endTime property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *endTimeIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *endTimeIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// generatorIntransitiveActivityIntermediateType will only have one of its values set at most
type generatorIntransitiveActivityIntermediateType struct {
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
func (t *generatorIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *generatorIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// iconIntransitiveActivityIntermediateType will only have one of its values set at most
type iconIntransitiveActivityIntermediateType struct {
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
func (t *iconIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *iconIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// imageIntransitiveActivityIntermediateType will only have one of its values set at most
type imageIntransitiveActivityIntermediateType struct {
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
func (t *imageIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *imageIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// inReplyToIntransitiveActivityIntermediateType will only have one of its values set at most
type inReplyToIntransitiveActivityIntermediateType struct {
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
func (t *inReplyToIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *inReplyToIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// locationIntransitiveActivityIntermediateType will only have one of its values set at most
type locationIntransitiveActivityIntermediateType struct {
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
func (t *locationIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *locationIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// previewIntransitiveActivityIntermediateType will only have one of its values set at most
type previewIntransitiveActivityIntermediateType struct {
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
func (t *previewIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *previewIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// publishedIntransitiveActivityIntermediateType will only have one of its values set at most
type publishedIntransitiveActivityIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *time.Time type for published property
	dateTime *time.Time
	// Stores possible *url.URL type for published property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *publishedIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *publishedIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// repliesIntransitiveActivityIntermediateType will only have one of its values set at most
type repliesIntransitiveActivityIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible CollectionType type for replies property
	Collection CollectionType
	// Stores possible *url.URL type for replies property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *repliesIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *repliesIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// startTimeIntransitiveActivityIntermediateType will only have one of its values set at most
type startTimeIntransitiveActivityIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *time.Time type for startTime property
	dateTime *time.Time
	// Stores possible *url.URL type for startTime property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *startTimeIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *startTimeIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// summaryIntransitiveActivityIntermediateType will only have one of its values set at most
type summaryIntransitiveActivityIntermediateType struct {
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
func (t *summaryIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *summaryIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// tagIntransitiveActivityIntermediateType will only have one of its values set at most
type tagIntransitiveActivityIntermediateType struct {
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
func (t *tagIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *tagIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// updatedIntransitiveActivityIntermediateType will only have one of its values set at most
type updatedIntransitiveActivityIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *time.Time type for updated property
	dateTime *time.Time
	// Stores possible *url.URL type for updated property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *updatedIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *updatedIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// urlIntransitiveActivityIntermediateType will only have one of its values set at most
type urlIntransitiveActivityIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *url.URL type for url property
	anyURI *url.URL
	// Stores possible LinkType type for url property
	Link LinkType
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *urlIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *urlIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// toIntransitiveActivityIntermediateType will only have one of its values set at most
type toIntransitiveActivityIntermediateType struct {
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
func (t *toIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *toIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// btoIntransitiveActivityIntermediateType will only have one of its values set at most
type btoIntransitiveActivityIntermediateType struct {
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
func (t *btoIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *btoIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// ccIntransitiveActivityIntermediateType will only have one of its values set at most
type ccIntransitiveActivityIntermediateType struct {
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
func (t *ccIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *ccIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// bccIntransitiveActivityIntermediateType will only have one of its values set at most
type bccIntransitiveActivityIntermediateType struct {
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
func (t *bccIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *bccIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// mediaTypeIntransitiveActivityIntermediateType will only have one of its values set at most
type mediaTypeIntransitiveActivityIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *string type for mediaType property
	mimeMediaTypeValue *string
	// Stores possible *url.URL type for mediaType property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *mediaTypeIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *mediaTypeIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// durationIntransitiveActivityIntermediateType will only have one of its values set at most
type durationIntransitiveActivityIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *time.Duration type for duration property
	duration *time.Duration
	// Stores possible *url.URL type for duration property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *durationIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *durationIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// sourceIntransitiveActivityIntermediateType will only have one of its values set at most
type sourceIntransitiveActivityIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for source property
	Object ObjectType
	// Stores possible *url.URL type for source property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *sourceIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *sourceIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// inboxIntransitiveActivityIntermediateType will only have one of its values set at most
type inboxIntransitiveActivityIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible OrderedCollectionType type for inbox property
	OrderedCollection OrderedCollectionType
	// Stores possible *url.URL type for inbox property
	anyURI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *inboxIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *inboxIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// outboxIntransitiveActivityIntermediateType will only have one of its values set at most
type outboxIntransitiveActivityIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible OrderedCollectionType type for outbox property
	OrderedCollection OrderedCollectionType
	// Stores possible *url.URL type for outbox property
	anyURI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *outboxIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *outboxIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// followingIntransitiveActivityIntermediateType will only have one of its values set at most
type followingIntransitiveActivityIntermediateType struct {
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
func (t *followingIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *followingIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// followersIntransitiveActivityIntermediateType will only have one of its values set at most
type followersIntransitiveActivityIntermediateType struct {
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
func (t *followersIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *followersIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// likedIntransitiveActivityIntermediateType will only have one of its values set at most
type likedIntransitiveActivityIntermediateType struct {
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
func (t *likedIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *likedIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// likesIntransitiveActivityIntermediateType will only have one of its values set at most
type likesIntransitiveActivityIntermediateType struct {
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
func (t *likesIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *likesIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// preferredUsernameIntransitiveActivityIntermediateType will only have one of its values set at most
type preferredUsernameIntransitiveActivityIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *string type for preferredUsername property
	stringName *string
	// Stores possible *url.URL type for preferredUsername property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *preferredUsernameIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *preferredUsernameIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// endpointsIntransitiveActivityIntermediateType will only have one of its values set at most
type endpointsIntransitiveActivityIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for endpoints property
	Object ObjectType
	// Stores possible *url.URL type for endpoints property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *endpointsIntransitiveActivityIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *endpointsIntransitiveActivityIntermediateType) Serialize() (i interface{}, err error) {
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

// deserializeactorIntransitiveActivityIntermediateType will accept a map to create a actorIntransitiveActivityIntermediateType
func deserializeActorIntransitiveActivityIntermediateType(in interface{}) (t *actorIntransitiveActivityIntermediateType, err error) {
	tmp := &actorIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice actorIntransitiveActivityIntermediateType will accept a slice to create a slice of actorIntransitiveActivityIntermediateType
func deserializeSliceActorIntransitiveActivityIntermediateType(in []interface{}) (t []*actorIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &actorIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeactorIntransitiveActivityIntermediateType will accept a actorIntransitiveActivityIntermediateType to create a map
func serializeActorIntransitiveActivityIntermediateType(t *actorIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceactorIntransitiveActivityIntermediateType will accept a slice of actorIntransitiveActivityIntermediateType to create a slice result
func serializeSliceActorIntransitiveActivityIntermediateType(s []*actorIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializetargetIntransitiveActivityIntermediateType will accept a map to create a targetIntransitiveActivityIntermediateType
func deserializeTargetIntransitiveActivityIntermediateType(in interface{}) (t *targetIntransitiveActivityIntermediateType, err error) {
	tmp := &targetIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice targetIntransitiveActivityIntermediateType will accept a slice to create a slice of targetIntransitiveActivityIntermediateType
func deserializeSliceTargetIntransitiveActivityIntermediateType(in []interface{}) (t []*targetIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &targetIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializetargetIntransitiveActivityIntermediateType will accept a targetIntransitiveActivityIntermediateType to create a map
func serializeTargetIntransitiveActivityIntermediateType(t *targetIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicetargetIntransitiveActivityIntermediateType will accept a slice of targetIntransitiveActivityIntermediateType to create a slice result
func serializeSliceTargetIntransitiveActivityIntermediateType(s []*targetIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeresultIntransitiveActivityIntermediateType will accept a map to create a resultIntransitiveActivityIntermediateType
func deserializeResultIntransitiveActivityIntermediateType(in interface{}) (t *resultIntransitiveActivityIntermediateType, err error) {
	tmp := &resultIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice resultIntransitiveActivityIntermediateType will accept a slice to create a slice of resultIntransitiveActivityIntermediateType
func deserializeSliceResultIntransitiveActivityIntermediateType(in []interface{}) (t []*resultIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &resultIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeresultIntransitiveActivityIntermediateType will accept a resultIntransitiveActivityIntermediateType to create a map
func serializeResultIntransitiveActivityIntermediateType(t *resultIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceresultIntransitiveActivityIntermediateType will accept a slice of resultIntransitiveActivityIntermediateType to create a slice result
func serializeSliceResultIntransitiveActivityIntermediateType(s []*resultIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeoriginIntransitiveActivityIntermediateType will accept a map to create a originIntransitiveActivityIntermediateType
func deserializeOriginIntransitiveActivityIntermediateType(in interface{}) (t *originIntransitiveActivityIntermediateType, err error) {
	tmp := &originIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice originIntransitiveActivityIntermediateType will accept a slice to create a slice of originIntransitiveActivityIntermediateType
func deserializeSliceOriginIntransitiveActivityIntermediateType(in []interface{}) (t []*originIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &originIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeoriginIntransitiveActivityIntermediateType will accept a originIntransitiveActivityIntermediateType to create a map
func serializeOriginIntransitiveActivityIntermediateType(t *originIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceoriginIntransitiveActivityIntermediateType will accept a slice of originIntransitiveActivityIntermediateType to create a slice result
func serializeSliceOriginIntransitiveActivityIntermediateType(s []*originIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeinstrumentIntransitiveActivityIntermediateType will accept a map to create a instrumentIntransitiveActivityIntermediateType
func deserializeInstrumentIntransitiveActivityIntermediateType(in interface{}) (t *instrumentIntransitiveActivityIntermediateType, err error) {
	tmp := &instrumentIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice instrumentIntransitiveActivityIntermediateType will accept a slice to create a slice of instrumentIntransitiveActivityIntermediateType
func deserializeSliceInstrumentIntransitiveActivityIntermediateType(in []interface{}) (t []*instrumentIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &instrumentIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeinstrumentIntransitiveActivityIntermediateType will accept a instrumentIntransitiveActivityIntermediateType to create a map
func serializeInstrumentIntransitiveActivityIntermediateType(t *instrumentIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceinstrumentIntransitiveActivityIntermediateType will accept a slice of instrumentIntransitiveActivityIntermediateType to create a slice result
func serializeSliceInstrumentIntransitiveActivityIntermediateType(s []*instrumentIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializealtitudeIntransitiveActivityIntermediateType will accept a map to create a altitudeIntransitiveActivityIntermediateType
func deserializeAltitudeIntransitiveActivityIntermediateType(in interface{}) (t *altitudeIntransitiveActivityIntermediateType, err error) {
	tmp := &altitudeIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice altitudeIntransitiveActivityIntermediateType will accept a slice to create a slice of altitudeIntransitiveActivityIntermediateType
func deserializeSliceAltitudeIntransitiveActivityIntermediateType(in []interface{}) (t []*altitudeIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &altitudeIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializealtitudeIntransitiveActivityIntermediateType will accept a altitudeIntransitiveActivityIntermediateType to create a map
func serializeAltitudeIntransitiveActivityIntermediateType(t *altitudeIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicealtitudeIntransitiveActivityIntermediateType will accept a slice of altitudeIntransitiveActivityIntermediateType to create a slice result
func serializeSliceAltitudeIntransitiveActivityIntermediateType(s []*altitudeIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeattachmentIntransitiveActivityIntermediateType will accept a map to create a attachmentIntransitiveActivityIntermediateType
func deserializeAttachmentIntransitiveActivityIntermediateType(in interface{}) (t *attachmentIntransitiveActivityIntermediateType, err error) {
	tmp := &attachmentIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice attachmentIntransitiveActivityIntermediateType will accept a slice to create a slice of attachmentIntransitiveActivityIntermediateType
func deserializeSliceAttachmentIntransitiveActivityIntermediateType(in []interface{}) (t []*attachmentIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &attachmentIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeattachmentIntransitiveActivityIntermediateType will accept a attachmentIntransitiveActivityIntermediateType to create a map
func serializeAttachmentIntransitiveActivityIntermediateType(t *attachmentIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceattachmentIntransitiveActivityIntermediateType will accept a slice of attachmentIntransitiveActivityIntermediateType to create a slice result
func serializeSliceAttachmentIntransitiveActivityIntermediateType(s []*attachmentIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeattributedToIntransitiveActivityIntermediateType will accept a map to create a attributedToIntransitiveActivityIntermediateType
func deserializeAttributedToIntransitiveActivityIntermediateType(in interface{}) (t *attributedToIntransitiveActivityIntermediateType, err error) {
	tmp := &attributedToIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice attributedToIntransitiveActivityIntermediateType will accept a slice to create a slice of attributedToIntransitiveActivityIntermediateType
func deserializeSliceAttributedToIntransitiveActivityIntermediateType(in []interface{}) (t []*attributedToIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &attributedToIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeattributedToIntransitiveActivityIntermediateType will accept a attributedToIntransitiveActivityIntermediateType to create a map
func serializeAttributedToIntransitiveActivityIntermediateType(t *attributedToIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceattributedToIntransitiveActivityIntermediateType will accept a slice of attributedToIntransitiveActivityIntermediateType to create a slice result
func serializeSliceAttributedToIntransitiveActivityIntermediateType(s []*attributedToIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeaudienceIntransitiveActivityIntermediateType will accept a map to create a audienceIntransitiveActivityIntermediateType
func deserializeAudienceIntransitiveActivityIntermediateType(in interface{}) (t *audienceIntransitiveActivityIntermediateType, err error) {
	tmp := &audienceIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice audienceIntransitiveActivityIntermediateType will accept a slice to create a slice of audienceIntransitiveActivityIntermediateType
func deserializeSliceAudienceIntransitiveActivityIntermediateType(in []interface{}) (t []*audienceIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &audienceIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeaudienceIntransitiveActivityIntermediateType will accept a audienceIntransitiveActivityIntermediateType to create a map
func serializeAudienceIntransitiveActivityIntermediateType(t *audienceIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceaudienceIntransitiveActivityIntermediateType will accept a slice of audienceIntransitiveActivityIntermediateType to create a slice result
func serializeSliceAudienceIntransitiveActivityIntermediateType(s []*audienceIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializecontentIntransitiveActivityIntermediateType will accept a map to create a contentIntransitiveActivityIntermediateType
func deserializeContentIntransitiveActivityIntermediateType(in interface{}) (t *contentIntransitiveActivityIntermediateType, err error) {
	tmp := &contentIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice contentIntransitiveActivityIntermediateType will accept a slice to create a slice of contentIntransitiveActivityIntermediateType
func deserializeSliceContentIntransitiveActivityIntermediateType(in []interface{}) (t []*contentIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &contentIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializecontentIntransitiveActivityIntermediateType will accept a contentIntransitiveActivityIntermediateType to create a map
func serializeContentIntransitiveActivityIntermediateType(t *contentIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicecontentIntransitiveActivityIntermediateType will accept a slice of contentIntransitiveActivityIntermediateType to create a slice result
func serializeSliceContentIntransitiveActivityIntermediateType(s []*contentIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializecontextIntransitiveActivityIntermediateType will accept a map to create a contextIntransitiveActivityIntermediateType
func deserializeContextIntransitiveActivityIntermediateType(in interface{}) (t *contextIntransitiveActivityIntermediateType, err error) {
	tmp := &contextIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice contextIntransitiveActivityIntermediateType will accept a slice to create a slice of contextIntransitiveActivityIntermediateType
func deserializeSliceContextIntransitiveActivityIntermediateType(in []interface{}) (t []*contextIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &contextIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializecontextIntransitiveActivityIntermediateType will accept a contextIntransitiveActivityIntermediateType to create a map
func serializeContextIntransitiveActivityIntermediateType(t *contextIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicecontextIntransitiveActivityIntermediateType will accept a slice of contextIntransitiveActivityIntermediateType to create a slice result
func serializeSliceContextIntransitiveActivityIntermediateType(s []*contextIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializenameIntransitiveActivityIntermediateType will accept a map to create a nameIntransitiveActivityIntermediateType
func deserializeNameIntransitiveActivityIntermediateType(in interface{}) (t *nameIntransitiveActivityIntermediateType, err error) {
	tmp := &nameIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice nameIntransitiveActivityIntermediateType will accept a slice to create a slice of nameIntransitiveActivityIntermediateType
func deserializeSliceNameIntransitiveActivityIntermediateType(in []interface{}) (t []*nameIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &nameIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializenameIntransitiveActivityIntermediateType will accept a nameIntransitiveActivityIntermediateType to create a map
func serializeNameIntransitiveActivityIntermediateType(t *nameIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicenameIntransitiveActivityIntermediateType will accept a slice of nameIntransitiveActivityIntermediateType to create a slice result
func serializeSliceNameIntransitiveActivityIntermediateType(s []*nameIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeendTimeIntransitiveActivityIntermediateType will accept a map to create a endTimeIntransitiveActivityIntermediateType
func deserializeEndTimeIntransitiveActivityIntermediateType(in interface{}) (t *endTimeIntransitiveActivityIntermediateType, err error) {
	tmp := &endTimeIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice endTimeIntransitiveActivityIntermediateType will accept a slice to create a slice of endTimeIntransitiveActivityIntermediateType
func deserializeSliceEndTimeIntransitiveActivityIntermediateType(in []interface{}) (t []*endTimeIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &endTimeIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeendTimeIntransitiveActivityIntermediateType will accept a endTimeIntransitiveActivityIntermediateType to create a map
func serializeEndTimeIntransitiveActivityIntermediateType(t *endTimeIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceendTimeIntransitiveActivityIntermediateType will accept a slice of endTimeIntransitiveActivityIntermediateType to create a slice result
func serializeSliceEndTimeIntransitiveActivityIntermediateType(s []*endTimeIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializegeneratorIntransitiveActivityIntermediateType will accept a map to create a generatorIntransitiveActivityIntermediateType
func deserializeGeneratorIntransitiveActivityIntermediateType(in interface{}) (t *generatorIntransitiveActivityIntermediateType, err error) {
	tmp := &generatorIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice generatorIntransitiveActivityIntermediateType will accept a slice to create a slice of generatorIntransitiveActivityIntermediateType
func deserializeSliceGeneratorIntransitiveActivityIntermediateType(in []interface{}) (t []*generatorIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &generatorIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializegeneratorIntransitiveActivityIntermediateType will accept a generatorIntransitiveActivityIntermediateType to create a map
func serializeGeneratorIntransitiveActivityIntermediateType(t *generatorIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicegeneratorIntransitiveActivityIntermediateType will accept a slice of generatorIntransitiveActivityIntermediateType to create a slice result
func serializeSliceGeneratorIntransitiveActivityIntermediateType(s []*generatorIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeiconIntransitiveActivityIntermediateType will accept a map to create a iconIntransitiveActivityIntermediateType
func deserializeIconIntransitiveActivityIntermediateType(in interface{}) (t *iconIntransitiveActivityIntermediateType, err error) {
	tmp := &iconIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice iconIntransitiveActivityIntermediateType will accept a slice to create a slice of iconIntransitiveActivityIntermediateType
func deserializeSliceIconIntransitiveActivityIntermediateType(in []interface{}) (t []*iconIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &iconIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeiconIntransitiveActivityIntermediateType will accept a iconIntransitiveActivityIntermediateType to create a map
func serializeIconIntransitiveActivityIntermediateType(t *iconIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceiconIntransitiveActivityIntermediateType will accept a slice of iconIntransitiveActivityIntermediateType to create a slice result
func serializeSliceIconIntransitiveActivityIntermediateType(s []*iconIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeimageIntransitiveActivityIntermediateType will accept a map to create a imageIntransitiveActivityIntermediateType
func deserializeImageIntransitiveActivityIntermediateType(in interface{}) (t *imageIntransitiveActivityIntermediateType, err error) {
	tmp := &imageIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice imageIntransitiveActivityIntermediateType will accept a slice to create a slice of imageIntransitiveActivityIntermediateType
func deserializeSliceImageIntransitiveActivityIntermediateType(in []interface{}) (t []*imageIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &imageIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeimageIntransitiveActivityIntermediateType will accept a imageIntransitiveActivityIntermediateType to create a map
func serializeImageIntransitiveActivityIntermediateType(t *imageIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceimageIntransitiveActivityIntermediateType will accept a slice of imageIntransitiveActivityIntermediateType to create a slice result
func serializeSliceImageIntransitiveActivityIntermediateType(s []*imageIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeinReplyToIntransitiveActivityIntermediateType will accept a map to create a inReplyToIntransitiveActivityIntermediateType
func deserializeInReplyToIntransitiveActivityIntermediateType(in interface{}) (t *inReplyToIntransitiveActivityIntermediateType, err error) {
	tmp := &inReplyToIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice inReplyToIntransitiveActivityIntermediateType will accept a slice to create a slice of inReplyToIntransitiveActivityIntermediateType
func deserializeSliceInReplyToIntransitiveActivityIntermediateType(in []interface{}) (t []*inReplyToIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &inReplyToIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeinReplyToIntransitiveActivityIntermediateType will accept a inReplyToIntransitiveActivityIntermediateType to create a map
func serializeInReplyToIntransitiveActivityIntermediateType(t *inReplyToIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceinReplyToIntransitiveActivityIntermediateType will accept a slice of inReplyToIntransitiveActivityIntermediateType to create a slice result
func serializeSliceInReplyToIntransitiveActivityIntermediateType(s []*inReplyToIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializelocationIntransitiveActivityIntermediateType will accept a map to create a locationIntransitiveActivityIntermediateType
func deserializeLocationIntransitiveActivityIntermediateType(in interface{}) (t *locationIntransitiveActivityIntermediateType, err error) {
	tmp := &locationIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice locationIntransitiveActivityIntermediateType will accept a slice to create a slice of locationIntransitiveActivityIntermediateType
func deserializeSliceLocationIntransitiveActivityIntermediateType(in []interface{}) (t []*locationIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &locationIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializelocationIntransitiveActivityIntermediateType will accept a locationIntransitiveActivityIntermediateType to create a map
func serializeLocationIntransitiveActivityIntermediateType(t *locationIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicelocationIntransitiveActivityIntermediateType will accept a slice of locationIntransitiveActivityIntermediateType to create a slice result
func serializeSliceLocationIntransitiveActivityIntermediateType(s []*locationIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializepreviewIntransitiveActivityIntermediateType will accept a map to create a previewIntransitiveActivityIntermediateType
func deserializePreviewIntransitiveActivityIntermediateType(in interface{}) (t *previewIntransitiveActivityIntermediateType, err error) {
	tmp := &previewIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice previewIntransitiveActivityIntermediateType will accept a slice to create a slice of previewIntransitiveActivityIntermediateType
func deserializeSlicePreviewIntransitiveActivityIntermediateType(in []interface{}) (t []*previewIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &previewIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializepreviewIntransitiveActivityIntermediateType will accept a previewIntransitiveActivityIntermediateType to create a map
func serializePreviewIntransitiveActivityIntermediateType(t *previewIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicepreviewIntransitiveActivityIntermediateType will accept a slice of previewIntransitiveActivityIntermediateType to create a slice result
func serializeSlicePreviewIntransitiveActivityIntermediateType(s []*previewIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializepublishedIntransitiveActivityIntermediateType will accept a map to create a publishedIntransitiveActivityIntermediateType
func deserializePublishedIntransitiveActivityIntermediateType(in interface{}) (t *publishedIntransitiveActivityIntermediateType, err error) {
	tmp := &publishedIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice publishedIntransitiveActivityIntermediateType will accept a slice to create a slice of publishedIntransitiveActivityIntermediateType
func deserializeSlicePublishedIntransitiveActivityIntermediateType(in []interface{}) (t []*publishedIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &publishedIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializepublishedIntransitiveActivityIntermediateType will accept a publishedIntransitiveActivityIntermediateType to create a map
func serializePublishedIntransitiveActivityIntermediateType(t *publishedIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicepublishedIntransitiveActivityIntermediateType will accept a slice of publishedIntransitiveActivityIntermediateType to create a slice result
func serializeSlicePublishedIntransitiveActivityIntermediateType(s []*publishedIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializerepliesIntransitiveActivityIntermediateType will accept a map to create a repliesIntransitiveActivityIntermediateType
func deserializeRepliesIntransitiveActivityIntermediateType(in interface{}) (t *repliesIntransitiveActivityIntermediateType, err error) {
	tmp := &repliesIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice repliesIntransitiveActivityIntermediateType will accept a slice to create a slice of repliesIntransitiveActivityIntermediateType
func deserializeSliceRepliesIntransitiveActivityIntermediateType(in []interface{}) (t []*repliesIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &repliesIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializerepliesIntransitiveActivityIntermediateType will accept a repliesIntransitiveActivityIntermediateType to create a map
func serializeRepliesIntransitiveActivityIntermediateType(t *repliesIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicerepliesIntransitiveActivityIntermediateType will accept a slice of repliesIntransitiveActivityIntermediateType to create a slice result
func serializeSliceRepliesIntransitiveActivityIntermediateType(s []*repliesIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializestartTimeIntransitiveActivityIntermediateType will accept a map to create a startTimeIntransitiveActivityIntermediateType
func deserializeStartTimeIntransitiveActivityIntermediateType(in interface{}) (t *startTimeIntransitiveActivityIntermediateType, err error) {
	tmp := &startTimeIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice startTimeIntransitiveActivityIntermediateType will accept a slice to create a slice of startTimeIntransitiveActivityIntermediateType
func deserializeSliceStartTimeIntransitiveActivityIntermediateType(in []interface{}) (t []*startTimeIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &startTimeIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializestartTimeIntransitiveActivityIntermediateType will accept a startTimeIntransitiveActivityIntermediateType to create a map
func serializeStartTimeIntransitiveActivityIntermediateType(t *startTimeIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicestartTimeIntransitiveActivityIntermediateType will accept a slice of startTimeIntransitiveActivityIntermediateType to create a slice result
func serializeSliceStartTimeIntransitiveActivityIntermediateType(s []*startTimeIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializesummaryIntransitiveActivityIntermediateType will accept a map to create a summaryIntransitiveActivityIntermediateType
func deserializeSummaryIntransitiveActivityIntermediateType(in interface{}) (t *summaryIntransitiveActivityIntermediateType, err error) {
	tmp := &summaryIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice summaryIntransitiveActivityIntermediateType will accept a slice to create a slice of summaryIntransitiveActivityIntermediateType
func deserializeSliceSummaryIntransitiveActivityIntermediateType(in []interface{}) (t []*summaryIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &summaryIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializesummaryIntransitiveActivityIntermediateType will accept a summaryIntransitiveActivityIntermediateType to create a map
func serializeSummaryIntransitiveActivityIntermediateType(t *summaryIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicesummaryIntransitiveActivityIntermediateType will accept a slice of summaryIntransitiveActivityIntermediateType to create a slice result
func serializeSliceSummaryIntransitiveActivityIntermediateType(s []*summaryIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializetagIntransitiveActivityIntermediateType will accept a map to create a tagIntransitiveActivityIntermediateType
func deserializeTagIntransitiveActivityIntermediateType(in interface{}) (t *tagIntransitiveActivityIntermediateType, err error) {
	tmp := &tagIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice tagIntransitiveActivityIntermediateType will accept a slice to create a slice of tagIntransitiveActivityIntermediateType
func deserializeSliceTagIntransitiveActivityIntermediateType(in []interface{}) (t []*tagIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &tagIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializetagIntransitiveActivityIntermediateType will accept a tagIntransitiveActivityIntermediateType to create a map
func serializeTagIntransitiveActivityIntermediateType(t *tagIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicetagIntransitiveActivityIntermediateType will accept a slice of tagIntransitiveActivityIntermediateType to create a slice result
func serializeSliceTagIntransitiveActivityIntermediateType(s []*tagIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeupdatedIntransitiveActivityIntermediateType will accept a map to create a updatedIntransitiveActivityIntermediateType
func deserializeUpdatedIntransitiveActivityIntermediateType(in interface{}) (t *updatedIntransitiveActivityIntermediateType, err error) {
	tmp := &updatedIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice updatedIntransitiveActivityIntermediateType will accept a slice to create a slice of updatedIntransitiveActivityIntermediateType
func deserializeSliceUpdatedIntransitiveActivityIntermediateType(in []interface{}) (t []*updatedIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &updatedIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeupdatedIntransitiveActivityIntermediateType will accept a updatedIntransitiveActivityIntermediateType to create a map
func serializeUpdatedIntransitiveActivityIntermediateType(t *updatedIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceupdatedIntransitiveActivityIntermediateType will accept a slice of updatedIntransitiveActivityIntermediateType to create a slice result
func serializeSliceUpdatedIntransitiveActivityIntermediateType(s []*updatedIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeurlIntransitiveActivityIntermediateType will accept a map to create a urlIntransitiveActivityIntermediateType
func deserializeUrlIntransitiveActivityIntermediateType(in interface{}) (t *urlIntransitiveActivityIntermediateType, err error) {
	tmp := &urlIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice urlIntransitiveActivityIntermediateType will accept a slice to create a slice of urlIntransitiveActivityIntermediateType
func deserializeSliceUrlIntransitiveActivityIntermediateType(in []interface{}) (t []*urlIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &urlIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeurlIntransitiveActivityIntermediateType will accept a urlIntransitiveActivityIntermediateType to create a map
func serializeUrlIntransitiveActivityIntermediateType(t *urlIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceurlIntransitiveActivityIntermediateType will accept a slice of urlIntransitiveActivityIntermediateType to create a slice result
func serializeSliceUrlIntransitiveActivityIntermediateType(s []*urlIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializetoIntransitiveActivityIntermediateType will accept a map to create a toIntransitiveActivityIntermediateType
func deserializeToIntransitiveActivityIntermediateType(in interface{}) (t *toIntransitiveActivityIntermediateType, err error) {
	tmp := &toIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice toIntransitiveActivityIntermediateType will accept a slice to create a slice of toIntransitiveActivityIntermediateType
func deserializeSliceToIntransitiveActivityIntermediateType(in []interface{}) (t []*toIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &toIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializetoIntransitiveActivityIntermediateType will accept a toIntransitiveActivityIntermediateType to create a map
func serializeToIntransitiveActivityIntermediateType(t *toIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicetoIntransitiveActivityIntermediateType will accept a slice of toIntransitiveActivityIntermediateType to create a slice result
func serializeSliceToIntransitiveActivityIntermediateType(s []*toIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializebtoIntransitiveActivityIntermediateType will accept a map to create a btoIntransitiveActivityIntermediateType
func deserializeBtoIntransitiveActivityIntermediateType(in interface{}) (t *btoIntransitiveActivityIntermediateType, err error) {
	tmp := &btoIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice btoIntransitiveActivityIntermediateType will accept a slice to create a slice of btoIntransitiveActivityIntermediateType
func deserializeSliceBtoIntransitiveActivityIntermediateType(in []interface{}) (t []*btoIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &btoIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializebtoIntransitiveActivityIntermediateType will accept a btoIntransitiveActivityIntermediateType to create a map
func serializeBtoIntransitiveActivityIntermediateType(t *btoIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicebtoIntransitiveActivityIntermediateType will accept a slice of btoIntransitiveActivityIntermediateType to create a slice result
func serializeSliceBtoIntransitiveActivityIntermediateType(s []*btoIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeccIntransitiveActivityIntermediateType will accept a map to create a ccIntransitiveActivityIntermediateType
func deserializeCcIntransitiveActivityIntermediateType(in interface{}) (t *ccIntransitiveActivityIntermediateType, err error) {
	tmp := &ccIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice ccIntransitiveActivityIntermediateType will accept a slice to create a slice of ccIntransitiveActivityIntermediateType
func deserializeSliceCcIntransitiveActivityIntermediateType(in []interface{}) (t []*ccIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &ccIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeccIntransitiveActivityIntermediateType will accept a ccIntransitiveActivityIntermediateType to create a map
func serializeCcIntransitiveActivityIntermediateType(t *ccIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceccIntransitiveActivityIntermediateType will accept a slice of ccIntransitiveActivityIntermediateType to create a slice result
func serializeSliceCcIntransitiveActivityIntermediateType(s []*ccIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializebccIntransitiveActivityIntermediateType will accept a map to create a bccIntransitiveActivityIntermediateType
func deserializeBccIntransitiveActivityIntermediateType(in interface{}) (t *bccIntransitiveActivityIntermediateType, err error) {
	tmp := &bccIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice bccIntransitiveActivityIntermediateType will accept a slice to create a slice of bccIntransitiveActivityIntermediateType
func deserializeSliceBccIntransitiveActivityIntermediateType(in []interface{}) (t []*bccIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &bccIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializebccIntransitiveActivityIntermediateType will accept a bccIntransitiveActivityIntermediateType to create a map
func serializeBccIntransitiveActivityIntermediateType(t *bccIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicebccIntransitiveActivityIntermediateType will accept a slice of bccIntransitiveActivityIntermediateType to create a slice result
func serializeSliceBccIntransitiveActivityIntermediateType(s []*bccIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializemediaTypeIntransitiveActivityIntermediateType will accept a map to create a mediaTypeIntransitiveActivityIntermediateType
func deserializeMediaTypeIntransitiveActivityIntermediateType(in interface{}) (t *mediaTypeIntransitiveActivityIntermediateType, err error) {
	tmp := &mediaTypeIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice mediaTypeIntransitiveActivityIntermediateType will accept a slice to create a slice of mediaTypeIntransitiveActivityIntermediateType
func deserializeSliceMediaTypeIntransitiveActivityIntermediateType(in []interface{}) (t []*mediaTypeIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &mediaTypeIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializemediaTypeIntransitiveActivityIntermediateType will accept a mediaTypeIntransitiveActivityIntermediateType to create a map
func serializeMediaTypeIntransitiveActivityIntermediateType(t *mediaTypeIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicemediaTypeIntransitiveActivityIntermediateType will accept a slice of mediaTypeIntransitiveActivityIntermediateType to create a slice result
func serializeSliceMediaTypeIntransitiveActivityIntermediateType(s []*mediaTypeIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializedurationIntransitiveActivityIntermediateType will accept a map to create a durationIntransitiveActivityIntermediateType
func deserializeDurationIntransitiveActivityIntermediateType(in interface{}) (t *durationIntransitiveActivityIntermediateType, err error) {
	tmp := &durationIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice durationIntransitiveActivityIntermediateType will accept a slice to create a slice of durationIntransitiveActivityIntermediateType
func deserializeSliceDurationIntransitiveActivityIntermediateType(in []interface{}) (t []*durationIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &durationIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializedurationIntransitiveActivityIntermediateType will accept a durationIntransitiveActivityIntermediateType to create a map
func serializeDurationIntransitiveActivityIntermediateType(t *durationIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicedurationIntransitiveActivityIntermediateType will accept a slice of durationIntransitiveActivityIntermediateType to create a slice result
func serializeSliceDurationIntransitiveActivityIntermediateType(s []*durationIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializesourceIntransitiveActivityIntermediateType will accept a map to create a sourceIntransitiveActivityIntermediateType
func deserializeSourceIntransitiveActivityIntermediateType(in interface{}) (t *sourceIntransitiveActivityIntermediateType, err error) {
	tmp := &sourceIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice sourceIntransitiveActivityIntermediateType will accept a slice to create a slice of sourceIntransitiveActivityIntermediateType
func deserializeSliceSourceIntransitiveActivityIntermediateType(in []interface{}) (t []*sourceIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &sourceIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializesourceIntransitiveActivityIntermediateType will accept a sourceIntransitiveActivityIntermediateType to create a map
func serializeSourceIntransitiveActivityIntermediateType(t *sourceIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicesourceIntransitiveActivityIntermediateType will accept a slice of sourceIntransitiveActivityIntermediateType to create a slice result
func serializeSliceSourceIntransitiveActivityIntermediateType(s []*sourceIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeinboxIntransitiveActivityIntermediateType will accept a map to create a inboxIntransitiveActivityIntermediateType
func deserializeInboxIntransitiveActivityIntermediateType(in interface{}) (t *inboxIntransitiveActivityIntermediateType, err error) {
	tmp := &inboxIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice inboxIntransitiveActivityIntermediateType will accept a slice to create a slice of inboxIntransitiveActivityIntermediateType
func deserializeSliceInboxIntransitiveActivityIntermediateType(in []interface{}) (t []*inboxIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &inboxIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeinboxIntransitiveActivityIntermediateType will accept a inboxIntransitiveActivityIntermediateType to create a map
func serializeInboxIntransitiveActivityIntermediateType(t *inboxIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceinboxIntransitiveActivityIntermediateType will accept a slice of inboxIntransitiveActivityIntermediateType to create a slice result
func serializeSliceInboxIntransitiveActivityIntermediateType(s []*inboxIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeoutboxIntransitiveActivityIntermediateType will accept a map to create a outboxIntransitiveActivityIntermediateType
func deserializeOutboxIntransitiveActivityIntermediateType(in interface{}) (t *outboxIntransitiveActivityIntermediateType, err error) {
	tmp := &outboxIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice outboxIntransitiveActivityIntermediateType will accept a slice to create a slice of outboxIntransitiveActivityIntermediateType
func deserializeSliceOutboxIntransitiveActivityIntermediateType(in []interface{}) (t []*outboxIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &outboxIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeoutboxIntransitiveActivityIntermediateType will accept a outboxIntransitiveActivityIntermediateType to create a map
func serializeOutboxIntransitiveActivityIntermediateType(t *outboxIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceoutboxIntransitiveActivityIntermediateType will accept a slice of outboxIntransitiveActivityIntermediateType to create a slice result
func serializeSliceOutboxIntransitiveActivityIntermediateType(s []*outboxIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializefollowingIntransitiveActivityIntermediateType will accept a map to create a followingIntransitiveActivityIntermediateType
func deserializeFollowingIntransitiveActivityIntermediateType(in interface{}) (t *followingIntransitiveActivityIntermediateType, err error) {
	tmp := &followingIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice followingIntransitiveActivityIntermediateType will accept a slice to create a slice of followingIntransitiveActivityIntermediateType
func deserializeSliceFollowingIntransitiveActivityIntermediateType(in []interface{}) (t []*followingIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &followingIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializefollowingIntransitiveActivityIntermediateType will accept a followingIntransitiveActivityIntermediateType to create a map
func serializeFollowingIntransitiveActivityIntermediateType(t *followingIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicefollowingIntransitiveActivityIntermediateType will accept a slice of followingIntransitiveActivityIntermediateType to create a slice result
func serializeSliceFollowingIntransitiveActivityIntermediateType(s []*followingIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializefollowersIntransitiveActivityIntermediateType will accept a map to create a followersIntransitiveActivityIntermediateType
func deserializeFollowersIntransitiveActivityIntermediateType(in interface{}) (t *followersIntransitiveActivityIntermediateType, err error) {
	tmp := &followersIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice followersIntransitiveActivityIntermediateType will accept a slice to create a slice of followersIntransitiveActivityIntermediateType
func deserializeSliceFollowersIntransitiveActivityIntermediateType(in []interface{}) (t []*followersIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &followersIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializefollowersIntransitiveActivityIntermediateType will accept a followersIntransitiveActivityIntermediateType to create a map
func serializeFollowersIntransitiveActivityIntermediateType(t *followersIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicefollowersIntransitiveActivityIntermediateType will accept a slice of followersIntransitiveActivityIntermediateType to create a slice result
func serializeSliceFollowersIntransitiveActivityIntermediateType(s []*followersIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializelikedIntransitiveActivityIntermediateType will accept a map to create a likedIntransitiveActivityIntermediateType
func deserializeLikedIntransitiveActivityIntermediateType(in interface{}) (t *likedIntransitiveActivityIntermediateType, err error) {
	tmp := &likedIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice likedIntransitiveActivityIntermediateType will accept a slice to create a slice of likedIntransitiveActivityIntermediateType
func deserializeSliceLikedIntransitiveActivityIntermediateType(in []interface{}) (t []*likedIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &likedIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializelikedIntransitiveActivityIntermediateType will accept a likedIntransitiveActivityIntermediateType to create a map
func serializeLikedIntransitiveActivityIntermediateType(t *likedIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicelikedIntransitiveActivityIntermediateType will accept a slice of likedIntransitiveActivityIntermediateType to create a slice result
func serializeSliceLikedIntransitiveActivityIntermediateType(s []*likedIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializelikesIntransitiveActivityIntermediateType will accept a map to create a likesIntransitiveActivityIntermediateType
func deserializeLikesIntransitiveActivityIntermediateType(in interface{}) (t *likesIntransitiveActivityIntermediateType, err error) {
	tmp := &likesIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice likesIntransitiveActivityIntermediateType will accept a slice to create a slice of likesIntransitiveActivityIntermediateType
func deserializeSliceLikesIntransitiveActivityIntermediateType(in []interface{}) (t []*likesIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &likesIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializelikesIntransitiveActivityIntermediateType will accept a likesIntransitiveActivityIntermediateType to create a map
func serializeLikesIntransitiveActivityIntermediateType(t *likesIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicelikesIntransitiveActivityIntermediateType will accept a slice of likesIntransitiveActivityIntermediateType to create a slice result
func serializeSliceLikesIntransitiveActivityIntermediateType(s []*likesIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializepreferredUsernameIntransitiveActivityIntermediateType will accept a map to create a preferredUsernameIntransitiveActivityIntermediateType
func deserializePreferredUsernameIntransitiveActivityIntermediateType(in interface{}) (t *preferredUsernameIntransitiveActivityIntermediateType, err error) {
	tmp := &preferredUsernameIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice preferredUsernameIntransitiveActivityIntermediateType will accept a slice to create a slice of preferredUsernameIntransitiveActivityIntermediateType
func deserializeSlicePreferredUsernameIntransitiveActivityIntermediateType(in []interface{}) (t []*preferredUsernameIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &preferredUsernameIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializepreferredUsernameIntransitiveActivityIntermediateType will accept a preferredUsernameIntransitiveActivityIntermediateType to create a map
func serializePreferredUsernameIntransitiveActivityIntermediateType(t *preferredUsernameIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicepreferredUsernameIntransitiveActivityIntermediateType will accept a slice of preferredUsernameIntransitiveActivityIntermediateType to create a slice result
func serializeSlicePreferredUsernameIntransitiveActivityIntermediateType(s []*preferredUsernameIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeendpointsIntransitiveActivityIntermediateType will accept a map to create a endpointsIntransitiveActivityIntermediateType
func deserializeEndpointsIntransitiveActivityIntermediateType(in interface{}) (t *endpointsIntransitiveActivityIntermediateType, err error) {
	tmp := &endpointsIntransitiveActivityIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice endpointsIntransitiveActivityIntermediateType will accept a slice to create a slice of endpointsIntransitiveActivityIntermediateType
func deserializeSliceEndpointsIntransitiveActivityIntermediateType(in []interface{}) (t []*endpointsIntransitiveActivityIntermediateType, err error) {
	for _, i := range in {
		tmp := &endpointsIntransitiveActivityIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeendpointsIntransitiveActivityIntermediateType will accept a endpointsIntransitiveActivityIntermediateType to create a map
func serializeEndpointsIntransitiveActivityIntermediateType(t *endpointsIntransitiveActivityIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceendpointsIntransitiveActivityIntermediateType will accept a slice of endpointsIntransitiveActivityIntermediateType to create a slice result
func serializeSliceEndpointsIntransitiveActivityIntermediateType(s []*endpointsIntransitiveActivityIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

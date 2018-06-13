//
package vocab

import (
	"fmt"
	"net/url"
	"time"
)

// CollectionType is an interface for accepting types that extend from 'Collection'.
type CollectionType interface {
	Serializer
	Deserializer
	IsTotalItems() (ok bool)
	GetTotalItems() (v int64)
	SetTotalItems(v int64)
	IsTotalItemsIRI() (ok bool)
	GetTotalItemsIRI() (v *url.URL)
	SetTotalItemsIRI(v *url.URL)
	IsCurrentCollectionPage() (ok bool)
	GetCurrentCollectionPage() (v CollectionPageType)
	SetCurrentCollectionPage(v CollectionPageType)
	IsCurrentLink() (ok bool)
	GetCurrentLink() (v LinkType)
	SetCurrentLink(v LinkType)
	IsCurrentIRI() (ok bool)
	GetCurrentIRI() (v *url.URL)
	SetCurrentIRI(v *url.URL)
	IsFirstCollectionPage() (ok bool)
	GetFirstCollectionPage() (v CollectionPageType)
	SetFirstCollectionPage(v CollectionPageType)
	IsFirstLink() (ok bool)
	GetFirstLink() (v LinkType)
	SetFirstLink(v LinkType)
	IsFirstIRI() (ok bool)
	GetFirstIRI() (v *url.URL)
	SetFirstIRI(v *url.URL)
	IsLastCollectionPage() (ok bool)
	GetLastCollectionPage() (v CollectionPageType)
	SetLastCollectionPage(v CollectionPageType)
	IsLastLink() (ok bool)
	GetLastLink() (v LinkType)
	SetLastLink(v LinkType)
	IsLastIRI() (ok bool)
	GetLastIRI() (v *url.URL)
	SetLastIRI(v *url.URL)
	ItemsLen() (l int)
	IsItemsObject(index int) (ok bool)
	GetItemsObject(index int) (v ObjectType)
	AppendItemsObject(v ObjectType)
	PrependItemsObject(v ObjectType)
	RemoveItemsObject(index int)
	IsItemsLink(index int) (ok bool)
	GetItemsLink(index int) (v LinkType)
	AppendItemsLink(v LinkType)
	PrependItemsLink(v LinkType)
	RemoveItemsLink(index int)
	IsItemsIRI(index int) (ok bool)
	GetItemsIRI(index int) (v *url.URL)
	AppendItemsIRI(v *url.URL)
	PrependItemsIRI(v *url.URL)
	RemoveItemsIRI(index int)
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

// A Collection is a subtype of Object that represents ordered or unordered sets of Object or Link instances. Refer to the Activity Streams 2.0 Core specification for a complete description of the Collection type.
type Collection struct {
	// An unknown value.
	unknown_ map[string]interface{}
	// The functional 'totalItems' value could have multiple types, but only a single value
	totalItems *totalItemsCollectionIntermediateType
	// The functional 'current' value could have multiple types, but only a single value
	current *currentCollectionIntermediateType
	// The functional 'first' value could have multiple types, but only a single value
	first *firstCollectionIntermediateType
	// The functional 'last' value could have multiple types, but only a single value
	last *lastCollectionIntermediateType
	// The 'items' value could have multiple types and values
	items []*itemsCollectionIntermediateType
	// The functional 'altitude' value could have multiple types, but only a single value
	altitude *altitudeCollectionIntermediateType
	// The 'attachment' value could have multiple types and values
	attachment []*attachmentCollectionIntermediateType
	// The 'attributedTo' value could have multiple types and values
	attributedTo []*attributedToCollectionIntermediateType
	// The 'audience' value could have multiple types and values
	audience []*audienceCollectionIntermediateType
	// The 'content' value could have multiple types and values
	content []*contentCollectionIntermediateType
	// The 'contentMap' value holds language-specific values for property 'content'
	contentMap map[string]string
	// The 'context' value could have multiple types and values
	context []*contextCollectionIntermediateType
	// The 'name' value could have multiple types and values
	name []*nameCollectionIntermediateType
	// The 'nameMap' value holds language-specific values for property 'name'
	nameMap map[string]string
	// The functional 'endTime' value could have multiple types, but only a single value
	endTime *endTimeCollectionIntermediateType
	// The 'generator' value could have multiple types and values
	generator []*generatorCollectionIntermediateType
	// The 'icon' value could have multiple types and values
	icon []*iconCollectionIntermediateType
	// The functional 'id' value holds a single type and a single value
	id *url.URL
	// The 'image' value could have multiple types and values
	image []*imageCollectionIntermediateType
	// The 'inReplyTo' value could have multiple types and values
	inReplyTo []*inReplyToCollectionIntermediateType
	// The 'location' value could have multiple types and values
	location []*locationCollectionIntermediateType
	// The 'preview' value could have multiple types and values
	preview []*previewCollectionIntermediateType
	// The functional 'published' value could have multiple types, but only a single value
	published *publishedCollectionIntermediateType
	// The functional 'replies' value could have multiple types, but only a single value
	replies *repliesCollectionIntermediateType
	// The functional 'startTime' value could have multiple types, but only a single value
	startTime *startTimeCollectionIntermediateType
	// The 'summary' value could have multiple types and values
	summary []*summaryCollectionIntermediateType
	// The 'summaryMap' value holds language-specific values for property 'summary'
	summaryMap map[string]string
	// The 'tag' value could have multiple types and values
	tag []*tagCollectionIntermediateType
	// The 'type' value can hold any type and any number of values
	typeName []interface{}
	// The functional 'updated' value could have multiple types, but only a single value
	updated *updatedCollectionIntermediateType
	// The 'url' value could have multiple types and values
	url []*urlCollectionIntermediateType
	// The 'to' value could have multiple types and values
	to []*toCollectionIntermediateType
	// The 'bto' value could have multiple types and values
	bto []*btoCollectionIntermediateType
	// The 'cc' value could have multiple types and values
	cc []*ccCollectionIntermediateType
	// The 'bcc' value could have multiple types and values
	bcc []*bccCollectionIntermediateType
	// The functional 'mediaType' value could have multiple types, but only a single value
	mediaType *mediaTypeCollectionIntermediateType
	// The functional 'duration' value could have multiple types, but only a single value
	duration *durationCollectionIntermediateType
	// The functional 'source' value could have multiple types, but only a single value
	source *sourceCollectionIntermediateType
	// The functional 'inbox' value could have multiple types, but only a single value
	inbox *inboxCollectionIntermediateType
	// The functional 'outbox' value could have multiple types, but only a single value
	outbox *outboxCollectionIntermediateType
	// The functional 'following' value could have multiple types, but only a single value
	following *followingCollectionIntermediateType
	// The functional 'followers' value could have multiple types, but only a single value
	followers *followersCollectionIntermediateType
	// The functional 'liked' value could have multiple types, but only a single value
	liked *likedCollectionIntermediateType
	// The functional 'likes' value could have multiple types, but only a single value
	likes *likesCollectionIntermediateType
	// The 'streams' value holds a single type and any number of values
	streams []*url.URL
	// The functional 'preferredUsername' value could have multiple types, but only a single value
	preferredUsername *preferredUsernameCollectionIntermediateType
	// The 'preferredUsernameMap' value holds language-specific values for property 'preferredUsername'
	preferredUsernameMap map[string]string
	// The functional 'endpoints' value could have multiple types, but only a single value
	endpoints *endpointsCollectionIntermediateType
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

// IsTotalItems determines whether the call to GetTotalItems is safe
func (t *Collection) IsTotalItems() (ok bool) {
	return t.totalItems != nil && t.totalItems.nonNegativeInteger != nil

}

// GetTotalItems returns the value safely if IsTotalItems returned true
func (t *Collection) GetTotalItems() (v int64) {
	return *t.totalItems.nonNegativeInteger

}

// SetTotalItems sets the value of totalItems to be of int64 type
func (t *Collection) SetTotalItems(v int64) {
	t.totalItems = &totalItemsCollectionIntermediateType{nonNegativeInteger: &v}

}

// IsTotalItemsIRI determines whether the call to GetTotalItemsIRI is safe
func (t *Collection) IsTotalItemsIRI() (ok bool) {
	return t.totalItems != nil && t.totalItems.IRI != nil

}

// GetTotalItemsIRI returns the value safely if IsTotalItemsIRI returned true
func (t *Collection) GetTotalItemsIRI() (v *url.URL) {
	return t.totalItems.IRI

}

// SetTotalItemsIRI sets the value of totalItems to be of *url.URL type
func (t *Collection) SetTotalItemsIRI(v *url.URL) {
	t.totalItems = &totalItemsCollectionIntermediateType{IRI: v}

}

// HasUnknownTotalItems determines whether the call to GetUnknownTotalItems is safe
func (t *Collection) HasUnknownTotalItems() (ok bool) {
	return t.totalItems != nil && t.totalItems.unknown_ != nil

}

// GetUnknownTotalItems returns the unknown value for totalItems
func (t *Collection) GetUnknownTotalItems() (v interface{}) {
	return t.totalItems.unknown_

}

// SetUnknownTotalItems sets the unknown value of totalItems
func (t *Collection) SetUnknownTotalItems(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &totalItemsCollectionIntermediateType{}
	tmp.unknown_ = i
	t.totalItems = tmp

}

// IsCurrentCollectionPage determines whether the call to GetCurrentCollectionPage is safe
func (t *Collection) IsCurrentCollectionPage() (ok bool) {
	return t.current != nil && t.current.CollectionPage != nil

}

// GetCurrentCollectionPage returns the value safely if IsCurrentCollectionPage returned true
func (t *Collection) GetCurrentCollectionPage() (v CollectionPageType) {
	return t.current.CollectionPage

}

// SetCurrentCollectionPage sets the value of current to be of CollectionPageType type
func (t *Collection) SetCurrentCollectionPage(v CollectionPageType) {
	t.current = &currentCollectionIntermediateType{CollectionPage: v}

}

// IsCurrentLink determines whether the call to GetCurrentLink is safe
func (t *Collection) IsCurrentLink() (ok bool) {
	return t.current != nil && t.current.Link != nil

}

// GetCurrentLink returns the value safely if IsCurrentLink returned true
func (t *Collection) GetCurrentLink() (v LinkType) {
	return t.current.Link

}

// SetCurrentLink sets the value of current to be of LinkType type
func (t *Collection) SetCurrentLink(v LinkType) {
	t.current = &currentCollectionIntermediateType{Link: v}

}

// IsCurrentIRI determines whether the call to GetCurrentIRI is safe
func (t *Collection) IsCurrentIRI() (ok bool) {
	return t.current != nil && t.current.IRI != nil

}

// GetCurrentIRI returns the value safely if IsCurrentIRI returned true
func (t *Collection) GetCurrentIRI() (v *url.URL) {
	return t.current.IRI

}

// SetCurrentIRI sets the value of current to be of *url.URL type
func (t *Collection) SetCurrentIRI(v *url.URL) {
	t.current = &currentCollectionIntermediateType{IRI: v}

}

// HasUnknownCurrent determines whether the call to GetUnknownCurrent is safe
func (t *Collection) HasUnknownCurrent() (ok bool) {
	return t.current != nil && t.current.unknown_ != nil

}

// GetUnknownCurrent returns the unknown value for current
func (t *Collection) GetUnknownCurrent() (v interface{}) {
	return t.current.unknown_

}

// SetUnknownCurrent sets the unknown value of current
func (t *Collection) SetUnknownCurrent(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &currentCollectionIntermediateType{}
	tmp.unknown_ = i
	t.current = tmp

}

// IsFirstCollectionPage determines whether the call to GetFirstCollectionPage is safe
func (t *Collection) IsFirstCollectionPage() (ok bool) {
	return t.first != nil && t.first.CollectionPage != nil

}

// GetFirstCollectionPage returns the value safely if IsFirstCollectionPage returned true
func (t *Collection) GetFirstCollectionPage() (v CollectionPageType) {
	return t.first.CollectionPage

}

// SetFirstCollectionPage sets the value of first to be of CollectionPageType type
func (t *Collection) SetFirstCollectionPage(v CollectionPageType) {
	t.first = &firstCollectionIntermediateType{CollectionPage: v}

}

// IsFirstLink determines whether the call to GetFirstLink is safe
func (t *Collection) IsFirstLink() (ok bool) {
	return t.first != nil && t.first.Link != nil

}

// GetFirstLink returns the value safely if IsFirstLink returned true
func (t *Collection) GetFirstLink() (v LinkType) {
	return t.first.Link

}

// SetFirstLink sets the value of first to be of LinkType type
func (t *Collection) SetFirstLink(v LinkType) {
	t.first = &firstCollectionIntermediateType{Link: v}

}

// IsFirstIRI determines whether the call to GetFirstIRI is safe
func (t *Collection) IsFirstIRI() (ok bool) {
	return t.first != nil && t.first.IRI != nil

}

// GetFirstIRI returns the value safely if IsFirstIRI returned true
func (t *Collection) GetFirstIRI() (v *url.URL) {
	return t.first.IRI

}

// SetFirstIRI sets the value of first to be of *url.URL type
func (t *Collection) SetFirstIRI(v *url.URL) {
	t.first = &firstCollectionIntermediateType{IRI: v}

}

// HasUnknownFirst determines whether the call to GetUnknownFirst is safe
func (t *Collection) HasUnknownFirst() (ok bool) {
	return t.first != nil && t.first.unknown_ != nil

}

// GetUnknownFirst returns the unknown value for first
func (t *Collection) GetUnknownFirst() (v interface{}) {
	return t.first.unknown_

}

// SetUnknownFirst sets the unknown value of first
func (t *Collection) SetUnknownFirst(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &firstCollectionIntermediateType{}
	tmp.unknown_ = i
	t.first = tmp

}

// IsLastCollectionPage determines whether the call to GetLastCollectionPage is safe
func (t *Collection) IsLastCollectionPage() (ok bool) {
	return t.last != nil && t.last.CollectionPage != nil

}

// GetLastCollectionPage returns the value safely if IsLastCollectionPage returned true
func (t *Collection) GetLastCollectionPage() (v CollectionPageType) {
	return t.last.CollectionPage

}

// SetLastCollectionPage sets the value of last to be of CollectionPageType type
func (t *Collection) SetLastCollectionPage(v CollectionPageType) {
	t.last = &lastCollectionIntermediateType{CollectionPage: v}

}

// IsLastLink determines whether the call to GetLastLink is safe
func (t *Collection) IsLastLink() (ok bool) {
	return t.last != nil && t.last.Link != nil

}

// GetLastLink returns the value safely if IsLastLink returned true
func (t *Collection) GetLastLink() (v LinkType) {
	return t.last.Link

}

// SetLastLink sets the value of last to be of LinkType type
func (t *Collection) SetLastLink(v LinkType) {
	t.last = &lastCollectionIntermediateType{Link: v}

}

// IsLastIRI determines whether the call to GetLastIRI is safe
func (t *Collection) IsLastIRI() (ok bool) {
	return t.last != nil && t.last.IRI != nil

}

// GetLastIRI returns the value safely if IsLastIRI returned true
func (t *Collection) GetLastIRI() (v *url.URL) {
	return t.last.IRI

}

// SetLastIRI sets the value of last to be of *url.URL type
func (t *Collection) SetLastIRI(v *url.URL) {
	t.last = &lastCollectionIntermediateType{IRI: v}

}

// HasUnknownLast determines whether the call to GetUnknownLast is safe
func (t *Collection) HasUnknownLast() (ok bool) {
	return t.last != nil && t.last.unknown_ != nil

}

// GetUnknownLast returns the unknown value for last
func (t *Collection) GetUnknownLast() (v interface{}) {
	return t.last.unknown_

}

// SetUnknownLast sets the unknown value of last
func (t *Collection) SetUnknownLast(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &lastCollectionIntermediateType{}
	tmp.unknown_ = i
	t.last = tmp

}

// ItemsLen determines the number of elements able to be used for the IsItemsObject, GetItemsObject, and RemoveItemsObject functions
func (t *Collection) ItemsLen() (l int) {
	return len(t.items)

}

// IsItemsObject determines whether the call to GetItemsObject is safe for the specified index
func (t *Collection) IsItemsObject(index int) (ok bool) {
	return t.items[index].Object != nil

}

// GetItemsObject returns the value safely if IsItemsObject returned true for the specified index
func (t *Collection) GetItemsObject(index int) (v ObjectType) {
	return t.items[index].Object

}

// AppendItemsObject adds to the back of items a ObjectType type
func (t *Collection) AppendItemsObject(v ObjectType) {
	t.items = append(t.items, &itemsCollectionIntermediateType{Object: v})

}

// PrependItemsObject adds to the front of items a ObjectType type
func (t *Collection) PrependItemsObject(v ObjectType) {
	t.items = append([]*itemsCollectionIntermediateType{&itemsCollectionIntermediateType{Object: v}}, t.items...)

}

// RemoveItemsObject deletes the value from the specified index
func (t *Collection) RemoveItemsObject(index int) {
	copy(t.items[index:], t.items[index+1:])
	t.items[len(t.items)-1] = nil
	t.items = t.items[:len(t.items)-1]

}

// IsItemsLink determines whether the call to GetItemsLink is safe for the specified index
func (t *Collection) IsItemsLink(index int) (ok bool) {
	return t.items[index].Link != nil

}

// GetItemsLink returns the value safely if IsItemsLink returned true for the specified index
func (t *Collection) GetItemsLink(index int) (v LinkType) {
	return t.items[index].Link

}

// AppendItemsLink adds to the back of items a LinkType type
func (t *Collection) AppendItemsLink(v LinkType) {
	t.items = append(t.items, &itemsCollectionIntermediateType{Link: v})

}

// PrependItemsLink adds to the front of items a LinkType type
func (t *Collection) PrependItemsLink(v LinkType) {
	t.items = append([]*itemsCollectionIntermediateType{&itemsCollectionIntermediateType{Link: v}}, t.items...)

}

// RemoveItemsLink deletes the value from the specified index
func (t *Collection) RemoveItemsLink(index int) {
	copy(t.items[index:], t.items[index+1:])
	t.items[len(t.items)-1] = nil
	t.items = t.items[:len(t.items)-1]

}

// IsItemsIRI determines whether the call to GetItemsIRI is safe for the specified index
func (t *Collection) IsItemsIRI(index int) (ok bool) {
	return t.items[index].IRI != nil

}

// GetItemsIRI returns the value safely if IsItemsIRI returned true for the specified index
func (t *Collection) GetItemsIRI(index int) (v *url.URL) {
	return t.items[index].IRI

}

// AppendItemsIRI adds to the back of items a *url.URL type
func (t *Collection) AppendItemsIRI(v *url.URL) {
	t.items = append(t.items, &itemsCollectionIntermediateType{IRI: v})

}

// PrependItemsIRI adds to the front of items a *url.URL type
func (t *Collection) PrependItemsIRI(v *url.URL) {
	t.items = append([]*itemsCollectionIntermediateType{&itemsCollectionIntermediateType{IRI: v}}, t.items...)

}

// RemoveItemsIRI deletes the value from the specified index
func (t *Collection) RemoveItemsIRI(index int) {
	copy(t.items[index:], t.items[index+1:])
	t.items[len(t.items)-1] = nil
	t.items = t.items[:len(t.items)-1]

}

// HasUnknownItems determines whether the call to GetUnknownItems is safe
func (t *Collection) HasUnknownItems() (ok bool) {
	return t.items != nil && t.items[0].unknown_ != nil

}

// GetUnknownItems returns the unknown value for items
func (t *Collection) GetUnknownItems() (v interface{}) {
	return t.items[0].unknown_

}

// SetUnknownItems sets the unknown value of items
func (t *Collection) SetUnknownItems(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &itemsCollectionIntermediateType{}
	tmp.unknown_ = i
	t.items = append(t.items, tmp)

}

// IsAltitude determines whether the call to GetAltitude is safe
func (t *Collection) IsAltitude() (ok bool) {
	return t.altitude != nil && t.altitude.float != nil

}

// GetAltitude returns the value safely if IsAltitude returned true
func (t *Collection) GetAltitude() (v float64) {
	return *t.altitude.float

}

// SetAltitude sets the value of altitude to be of float64 type
func (t *Collection) SetAltitude(v float64) {
	t.altitude = &altitudeCollectionIntermediateType{float: &v}

}

// IsAltitudeIRI determines whether the call to GetAltitudeIRI is safe
func (t *Collection) IsAltitudeIRI() (ok bool) {
	return t.altitude != nil && t.altitude.IRI != nil

}

// GetAltitudeIRI returns the value safely if IsAltitudeIRI returned true
func (t *Collection) GetAltitudeIRI() (v *url.URL) {
	return t.altitude.IRI

}

// SetAltitudeIRI sets the value of altitude to be of *url.URL type
func (t *Collection) SetAltitudeIRI(v *url.URL) {
	t.altitude = &altitudeCollectionIntermediateType{IRI: v}

}

// HasUnknownAltitude determines whether the call to GetUnknownAltitude is safe
func (t *Collection) HasUnknownAltitude() (ok bool) {
	return t.altitude != nil && t.altitude.unknown_ != nil

}

// GetUnknownAltitude returns the unknown value for altitude
func (t *Collection) GetUnknownAltitude() (v interface{}) {
	return t.altitude.unknown_

}

// SetUnknownAltitude sets the unknown value of altitude
func (t *Collection) SetUnknownAltitude(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &altitudeCollectionIntermediateType{}
	tmp.unknown_ = i
	t.altitude = tmp

}

// AttachmentLen determines the number of elements able to be used for the IsAttachmentObject, GetAttachmentObject, and RemoveAttachmentObject functions
func (t *Collection) AttachmentLen() (l int) {
	return len(t.attachment)

}

// IsAttachmentObject determines whether the call to GetAttachmentObject is safe for the specified index
func (t *Collection) IsAttachmentObject(index int) (ok bool) {
	return t.attachment[index].Object != nil

}

// GetAttachmentObject returns the value safely if IsAttachmentObject returned true for the specified index
func (t *Collection) GetAttachmentObject(index int) (v ObjectType) {
	return t.attachment[index].Object

}

// AppendAttachmentObject adds to the back of attachment a ObjectType type
func (t *Collection) AppendAttachmentObject(v ObjectType) {
	t.attachment = append(t.attachment, &attachmentCollectionIntermediateType{Object: v})

}

// PrependAttachmentObject adds to the front of attachment a ObjectType type
func (t *Collection) PrependAttachmentObject(v ObjectType) {
	t.attachment = append([]*attachmentCollectionIntermediateType{&attachmentCollectionIntermediateType{Object: v}}, t.attachment...)

}

// RemoveAttachmentObject deletes the value from the specified index
func (t *Collection) RemoveAttachmentObject(index int) {
	copy(t.attachment[index:], t.attachment[index+1:])
	t.attachment[len(t.attachment)-1] = nil
	t.attachment = t.attachment[:len(t.attachment)-1]

}

// IsAttachmentLink determines whether the call to GetAttachmentLink is safe for the specified index
func (t *Collection) IsAttachmentLink(index int) (ok bool) {
	return t.attachment[index].Link != nil

}

// GetAttachmentLink returns the value safely if IsAttachmentLink returned true for the specified index
func (t *Collection) GetAttachmentLink(index int) (v LinkType) {
	return t.attachment[index].Link

}

// AppendAttachmentLink adds to the back of attachment a LinkType type
func (t *Collection) AppendAttachmentLink(v LinkType) {
	t.attachment = append(t.attachment, &attachmentCollectionIntermediateType{Link: v})

}

// PrependAttachmentLink adds to the front of attachment a LinkType type
func (t *Collection) PrependAttachmentLink(v LinkType) {
	t.attachment = append([]*attachmentCollectionIntermediateType{&attachmentCollectionIntermediateType{Link: v}}, t.attachment...)

}

// RemoveAttachmentLink deletes the value from the specified index
func (t *Collection) RemoveAttachmentLink(index int) {
	copy(t.attachment[index:], t.attachment[index+1:])
	t.attachment[len(t.attachment)-1] = nil
	t.attachment = t.attachment[:len(t.attachment)-1]

}

// IsAttachmentIRI determines whether the call to GetAttachmentIRI is safe for the specified index
func (t *Collection) IsAttachmentIRI(index int) (ok bool) {
	return t.attachment[index].IRI != nil

}

// GetAttachmentIRI returns the value safely if IsAttachmentIRI returned true for the specified index
func (t *Collection) GetAttachmentIRI(index int) (v *url.URL) {
	return t.attachment[index].IRI

}

// AppendAttachmentIRI adds to the back of attachment a *url.URL type
func (t *Collection) AppendAttachmentIRI(v *url.URL) {
	t.attachment = append(t.attachment, &attachmentCollectionIntermediateType{IRI: v})

}

// PrependAttachmentIRI adds to the front of attachment a *url.URL type
func (t *Collection) PrependAttachmentIRI(v *url.URL) {
	t.attachment = append([]*attachmentCollectionIntermediateType{&attachmentCollectionIntermediateType{IRI: v}}, t.attachment...)

}

// RemoveAttachmentIRI deletes the value from the specified index
func (t *Collection) RemoveAttachmentIRI(index int) {
	copy(t.attachment[index:], t.attachment[index+1:])
	t.attachment[len(t.attachment)-1] = nil
	t.attachment = t.attachment[:len(t.attachment)-1]

}

// HasUnknownAttachment determines whether the call to GetUnknownAttachment is safe
func (t *Collection) HasUnknownAttachment() (ok bool) {
	return t.attachment != nil && t.attachment[0].unknown_ != nil

}

// GetUnknownAttachment returns the unknown value for attachment
func (t *Collection) GetUnknownAttachment() (v interface{}) {
	return t.attachment[0].unknown_

}

// SetUnknownAttachment sets the unknown value of attachment
func (t *Collection) SetUnknownAttachment(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &attachmentCollectionIntermediateType{}
	tmp.unknown_ = i
	t.attachment = append(t.attachment, tmp)

}

// AttributedToLen determines the number of elements able to be used for the IsAttributedToObject, GetAttributedToObject, and RemoveAttributedToObject functions
func (t *Collection) AttributedToLen() (l int) {
	return len(t.attributedTo)

}

// IsAttributedToObject determines whether the call to GetAttributedToObject is safe for the specified index
func (t *Collection) IsAttributedToObject(index int) (ok bool) {
	return t.attributedTo[index].Object != nil

}

// GetAttributedToObject returns the value safely if IsAttributedToObject returned true for the specified index
func (t *Collection) GetAttributedToObject(index int) (v ObjectType) {
	return t.attributedTo[index].Object

}

// AppendAttributedToObject adds to the back of attributedTo a ObjectType type
func (t *Collection) AppendAttributedToObject(v ObjectType) {
	t.attributedTo = append(t.attributedTo, &attributedToCollectionIntermediateType{Object: v})

}

// PrependAttributedToObject adds to the front of attributedTo a ObjectType type
func (t *Collection) PrependAttributedToObject(v ObjectType) {
	t.attributedTo = append([]*attributedToCollectionIntermediateType{&attributedToCollectionIntermediateType{Object: v}}, t.attributedTo...)

}

// RemoveAttributedToObject deletes the value from the specified index
func (t *Collection) RemoveAttributedToObject(index int) {
	copy(t.attributedTo[index:], t.attributedTo[index+1:])
	t.attributedTo[len(t.attributedTo)-1] = nil
	t.attributedTo = t.attributedTo[:len(t.attributedTo)-1]

}

// IsAttributedToLink determines whether the call to GetAttributedToLink is safe for the specified index
func (t *Collection) IsAttributedToLink(index int) (ok bool) {
	return t.attributedTo[index].Link != nil

}

// GetAttributedToLink returns the value safely if IsAttributedToLink returned true for the specified index
func (t *Collection) GetAttributedToLink(index int) (v LinkType) {
	return t.attributedTo[index].Link

}

// AppendAttributedToLink adds to the back of attributedTo a LinkType type
func (t *Collection) AppendAttributedToLink(v LinkType) {
	t.attributedTo = append(t.attributedTo, &attributedToCollectionIntermediateType{Link: v})

}

// PrependAttributedToLink adds to the front of attributedTo a LinkType type
func (t *Collection) PrependAttributedToLink(v LinkType) {
	t.attributedTo = append([]*attributedToCollectionIntermediateType{&attributedToCollectionIntermediateType{Link: v}}, t.attributedTo...)

}

// RemoveAttributedToLink deletes the value from the specified index
func (t *Collection) RemoveAttributedToLink(index int) {
	copy(t.attributedTo[index:], t.attributedTo[index+1:])
	t.attributedTo[len(t.attributedTo)-1] = nil
	t.attributedTo = t.attributedTo[:len(t.attributedTo)-1]

}

// IsAttributedToIRI determines whether the call to GetAttributedToIRI is safe for the specified index
func (t *Collection) IsAttributedToIRI(index int) (ok bool) {
	return t.attributedTo[index].IRI != nil

}

// GetAttributedToIRI returns the value safely if IsAttributedToIRI returned true for the specified index
func (t *Collection) GetAttributedToIRI(index int) (v *url.URL) {
	return t.attributedTo[index].IRI

}

// AppendAttributedToIRI adds to the back of attributedTo a *url.URL type
func (t *Collection) AppendAttributedToIRI(v *url.URL) {
	t.attributedTo = append(t.attributedTo, &attributedToCollectionIntermediateType{IRI: v})

}

// PrependAttributedToIRI adds to the front of attributedTo a *url.URL type
func (t *Collection) PrependAttributedToIRI(v *url.URL) {
	t.attributedTo = append([]*attributedToCollectionIntermediateType{&attributedToCollectionIntermediateType{IRI: v}}, t.attributedTo...)

}

// RemoveAttributedToIRI deletes the value from the specified index
func (t *Collection) RemoveAttributedToIRI(index int) {
	copy(t.attributedTo[index:], t.attributedTo[index+1:])
	t.attributedTo[len(t.attributedTo)-1] = nil
	t.attributedTo = t.attributedTo[:len(t.attributedTo)-1]

}

// HasUnknownAttributedTo determines whether the call to GetUnknownAttributedTo is safe
func (t *Collection) HasUnknownAttributedTo() (ok bool) {
	return t.attributedTo != nil && t.attributedTo[0].unknown_ != nil

}

// GetUnknownAttributedTo returns the unknown value for attributedTo
func (t *Collection) GetUnknownAttributedTo() (v interface{}) {
	return t.attributedTo[0].unknown_

}

// SetUnknownAttributedTo sets the unknown value of attributedTo
func (t *Collection) SetUnknownAttributedTo(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &attributedToCollectionIntermediateType{}
	tmp.unknown_ = i
	t.attributedTo = append(t.attributedTo, tmp)

}

// AudienceLen determines the number of elements able to be used for the IsAudienceObject, GetAudienceObject, and RemoveAudienceObject functions
func (t *Collection) AudienceLen() (l int) {
	return len(t.audience)

}

// IsAudienceObject determines whether the call to GetAudienceObject is safe for the specified index
func (t *Collection) IsAudienceObject(index int) (ok bool) {
	return t.audience[index].Object != nil

}

// GetAudienceObject returns the value safely if IsAudienceObject returned true for the specified index
func (t *Collection) GetAudienceObject(index int) (v ObjectType) {
	return t.audience[index].Object

}

// AppendAudienceObject adds to the back of audience a ObjectType type
func (t *Collection) AppendAudienceObject(v ObjectType) {
	t.audience = append(t.audience, &audienceCollectionIntermediateType{Object: v})

}

// PrependAudienceObject adds to the front of audience a ObjectType type
func (t *Collection) PrependAudienceObject(v ObjectType) {
	t.audience = append([]*audienceCollectionIntermediateType{&audienceCollectionIntermediateType{Object: v}}, t.audience...)

}

// RemoveAudienceObject deletes the value from the specified index
func (t *Collection) RemoveAudienceObject(index int) {
	copy(t.audience[index:], t.audience[index+1:])
	t.audience[len(t.audience)-1] = nil
	t.audience = t.audience[:len(t.audience)-1]

}

// IsAudienceLink determines whether the call to GetAudienceLink is safe for the specified index
func (t *Collection) IsAudienceLink(index int) (ok bool) {
	return t.audience[index].Link != nil

}

// GetAudienceLink returns the value safely if IsAudienceLink returned true for the specified index
func (t *Collection) GetAudienceLink(index int) (v LinkType) {
	return t.audience[index].Link

}

// AppendAudienceLink adds to the back of audience a LinkType type
func (t *Collection) AppendAudienceLink(v LinkType) {
	t.audience = append(t.audience, &audienceCollectionIntermediateType{Link: v})

}

// PrependAudienceLink adds to the front of audience a LinkType type
func (t *Collection) PrependAudienceLink(v LinkType) {
	t.audience = append([]*audienceCollectionIntermediateType{&audienceCollectionIntermediateType{Link: v}}, t.audience...)

}

// RemoveAudienceLink deletes the value from the specified index
func (t *Collection) RemoveAudienceLink(index int) {
	copy(t.audience[index:], t.audience[index+1:])
	t.audience[len(t.audience)-1] = nil
	t.audience = t.audience[:len(t.audience)-1]

}

// IsAudienceIRI determines whether the call to GetAudienceIRI is safe for the specified index
func (t *Collection) IsAudienceIRI(index int) (ok bool) {
	return t.audience[index].IRI != nil

}

// GetAudienceIRI returns the value safely if IsAudienceIRI returned true for the specified index
func (t *Collection) GetAudienceIRI(index int) (v *url.URL) {
	return t.audience[index].IRI

}

// AppendAudienceIRI adds to the back of audience a *url.URL type
func (t *Collection) AppendAudienceIRI(v *url.URL) {
	t.audience = append(t.audience, &audienceCollectionIntermediateType{IRI: v})

}

// PrependAudienceIRI adds to the front of audience a *url.URL type
func (t *Collection) PrependAudienceIRI(v *url.URL) {
	t.audience = append([]*audienceCollectionIntermediateType{&audienceCollectionIntermediateType{IRI: v}}, t.audience...)

}

// RemoveAudienceIRI deletes the value from the specified index
func (t *Collection) RemoveAudienceIRI(index int) {
	copy(t.audience[index:], t.audience[index+1:])
	t.audience[len(t.audience)-1] = nil
	t.audience = t.audience[:len(t.audience)-1]

}

// HasUnknownAudience determines whether the call to GetUnknownAudience is safe
func (t *Collection) HasUnknownAudience() (ok bool) {
	return t.audience != nil && t.audience[0].unknown_ != nil

}

// GetUnknownAudience returns the unknown value for audience
func (t *Collection) GetUnknownAudience() (v interface{}) {
	return t.audience[0].unknown_

}

// SetUnknownAudience sets the unknown value of audience
func (t *Collection) SetUnknownAudience(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &audienceCollectionIntermediateType{}
	tmp.unknown_ = i
	t.audience = append(t.audience, tmp)

}

// ContentLen determines the number of elements able to be used for the IsContentString, GetContentString, and RemoveContentString functions
func (t *Collection) ContentLen() (l int) {
	return len(t.content)

}

// IsContentString determines whether the call to GetContentString is safe for the specified index
func (t *Collection) IsContentString(index int) (ok bool) {
	return t.content[index].stringName != nil

}

// GetContentString returns the value safely if IsContentString returned true for the specified index
func (t *Collection) GetContentString(index int) (v string) {
	return *t.content[index].stringName

}

// AppendContentString adds to the back of content a string type
func (t *Collection) AppendContentString(v string) {
	t.content = append(t.content, &contentCollectionIntermediateType{stringName: &v})

}

// PrependContentString adds to the front of content a string type
func (t *Collection) PrependContentString(v string) {
	t.content = append([]*contentCollectionIntermediateType{&contentCollectionIntermediateType{stringName: &v}}, t.content...)

}

// RemoveContentString deletes the value from the specified index
func (t *Collection) RemoveContentString(index int) {
	copy(t.content[index:], t.content[index+1:])
	t.content[len(t.content)-1] = nil
	t.content = t.content[:len(t.content)-1]

}

// IsContentLangString determines whether the call to GetContentLangString is safe for the specified index
func (t *Collection) IsContentLangString(index int) (ok bool) {
	return t.content[index].langString != nil

}

// GetContentLangString returns the value safely if IsContentLangString returned true for the specified index
func (t *Collection) GetContentLangString(index int) (v string) {
	return *t.content[index].langString

}

// AppendContentLangString adds to the back of content a string type
func (t *Collection) AppendContentLangString(v string) {
	t.content = append(t.content, &contentCollectionIntermediateType{langString: &v})

}

// PrependContentLangString adds to the front of content a string type
func (t *Collection) PrependContentLangString(v string) {
	t.content = append([]*contentCollectionIntermediateType{&contentCollectionIntermediateType{langString: &v}}, t.content...)

}

// RemoveContentLangString deletes the value from the specified index
func (t *Collection) RemoveContentLangString(index int) {
	copy(t.content[index:], t.content[index+1:])
	t.content[len(t.content)-1] = nil
	t.content = t.content[:len(t.content)-1]

}

// IsContentIRI determines whether the call to GetContentIRI is safe for the specified index
func (t *Collection) IsContentIRI(index int) (ok bool) {
	return t.content[index].IRI != nil

}

// GetContentIRI returns the value safely if IsContentIRI returned true for the specified index
func (t *Collection) GetContentIRI(index int) (v *url.URL) {
	return t.content[index].IRI

}

// AppendContentIRI adds to the back of content a *url.URL type
func (t *Collection) AppendContentIRI(v *url.URL) {
	t.content = append(t.content, &contentCollectionIntermediateType{IRI: v})

}

// PrependContentIRI adds to the front of content a *url.URL type
func (t *Collection) PrependContentIRI(v *url.URL) {
	t.content = append([]*contentCollectionIntermediateType{&contentCollectionIntermediateType{IRI: v}}, t.content...)

}

// RemoveContentIRI deletes the value from the specified index
func (t *Collection) RemoveContentIRI(index int) {
	copy(t.content[index:], t.content[index+1:])
	t.content[len(t.content)-1] = nil
	t.content = t.content[:len(t.content)-1]

}

// HasUnknownContent determines whether the call to GetUnknownContent is safe
func (t *Collection) HasUnknownContent() (ok bool) {
	return t.content != nil && t.content[0].unknown_ != nil

}

// GetUnknownContent returns the unknown value for content
func (t *Collection) GetUnknownContent() (v interface{}) {
	return t.content[0].unknown_

}

// SetUnknownContent sets the unknown value of content
func (t *Collection) SetUnknownContent(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &contentCollectionIntermediateType{}
	tmp.unknown_ = i
	t.content = append(t.content, tmp)

}

// ContentMapLanguages returns all languages for this property's language mapping, or nil if there are none.
func (t *Collection) ContentMapLanguages() (l []string) {
	if t.contentMap == nil || len(t.contentMap) == 0 {
		return nil
	}
	for k := range t.contentMap {
		l = append(l, k)
	}
	return

}

// GetContentMap retrieves the value of the property for the specified language, or an empty string if it does not exist
func (t *Collection) GetContentMap(l string) (v string) {
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
func (t *Collection) SetContentMap(l string, v string) {
	if t.contentMap == nil {
		t.contentMap = make(map[string]string)
	}
	t.contentMap[l] = v

}

// ContextLen determines the number of elements able to be used for the IsContextObject, GetContextObject, and RemoveContextObject functions
func (t *Collection) ContextLen() (l int) {
	return len(t.context)

}

// IsContextObject determines whether the call to GetContextObject is safe for the specified index
func (t *Collection) IsContextObject(index int) (ok bool) {
	return t.context[index].Object != nil

}

// GetContextObject returns the value safely if IsContextObject returned true for the specified index
func (t *Collection) GetContextObject(index int) (v ObjectType) {
	return t.context[index].Object

}

// AppendContextObject adds to the back of context a ObjectType type
func (t *Collection) AppendContextObject(v ObjectType) {
	t.context = append(t.context, &contextCollectionIntermediateType{Object: v})

}

// PrependContextObject adds to the front of context a ObjectType type
func (t *Collection) PrependContextObject(v ObjectType) {
	t.context = append([]*contextCollectionIntermediateType{&contextCollectionIntermediateType{Object: v}}, t.context...)

}

// RemoveContextObject deletes the value from the specified index
func (t *Collection) RemoveContextObject(index int) {
	copy(t.context[index:], t.context[index+1:])
	t.context[len(t.context)-1] = nil
	t.context = t.context[:len(t.context)-1]

}

// IsContextLink determines whether the call to GetContextLink is safe for the specified index
func (t *Collection) IsContextLink(index int) (ok bool) {
	return t.context[index].Link != nil

}

// GetContextLink returns the value safely if IsContextLink returned true for the specified index
func (t *Collection) GetContextLink(index int) (v LinkType) {
	return t.context[index].Link

}

// AppendContextLink adds to the back of context a LinkType type
func (t *Collection) AppendContextLink(v LinkType) {
	t.context = append(t.context, &contextCollectionIntermediateType{Link: v})

}

// PrependContextLink adds to the front of context a LinkType type
func (t *Collection) PrependContextLink(v LinkType) {
	t.context = append([]*contextCollectionIntermediateType{&contextCollectionIntermediateType{Link: v}}, t.context...)

}

// RemoveContextLink deletes the value from the specified index
func (t *Collection) RemoveContextLink(index int) {
	copy(t.context[index:], t.context[index+1:])
	t.context[len(t.context)-1] = nil
	t.context = t.context[:len(t.context)-1]

}

// IsContextIRI determines whether the call to GetContextIRI is safe for the specified index
func (t *Collection) IsContextIRI(index int) (ok bool) {
	return t.context[index].IRI != nil

}

// GetContextIRI returns the value safely if IsContextIRI returned true for the specified index
func (t *Collection) GetContextIRI(index int) (v *url.URL) {
	return t.context[index].IRI

}

// AppendContextIRI adds to the back of context a *url.URL type
func (t *Collection) AppendContextIRI(v *url.URL) {
	t.context = append(t.context, &contextCollectionIntermediateType{IRI: v})

}

// PrependContextIRI adds to the front of context a *url.URL type
func (t *Collection) PrependContextIRI(v *url.URL) {
	t.context = append([]*contextCollectionIntermediateType{&contextCollectionIntermediateType{IRI: v}}, t.context...)

}

// RemoveContextIRI deletes the value from the specified index
func (t *Collection) RemoveContextIRI(index int) {
	copy(t.context[index:], t.context[index+1:])
	t.context[len(t.context)-1] = nil
	t.context = t.context[:len(t.context)-1]

}

// HasUnknownContext determines whether the call to GetUnknownContext is safe
func (t *Collection) HasUnknownContext() (ok bool) {
	return t.context != nil && t.context[0].unknown_ != nil

}

// GetUnknownContext returns the unknown value for context
func (t *Collection) GetUnknownContext() (v interface{}) {
	return t.context[0].unknown_

}

// SetUnknownContext sets the unknown value of context
func (t *Collection) SetUnknownContext(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &contextCollectionIntermediateType{}
	tmp.unknown_ = i
	t.context = append(t.context, tmp)

}

// NameLen determines the number of elements able to be used for the IsNameString, GetNameString, and RemoveNameString functions
func (t *Collection) NameLen() (l int) {
	return len(t.name)

}

// IsNameString determines whether the call to GetNameString is safe for the specified index
func (t *Collection) IsNameString(index int) (ok bool) {
	return t.name[index].stringName != nil

}

// GetNameString returns the value safely if IsNameString returned true for the specified index
func (t *Collection) GetNameString(index int) (v string) {
	return *t.name[index].stringName

}

// AppendNameString adds to the back of name a string type
func (t *Collection) AppendNameString(v string) {
	t.name = append(t.name, &nameCollectionIntermediateType{stringName: &v})

}

// PrependNameString adds to the front of name a string type
func (t *Collection) PrependNameString(v string) {
	t.name = append([]*nameCollectionIntermediateType{&nameCollectionIntermediateType{stringName: &v}}, t.name...)

}

// RemoveNameString deletes the value from the specified index
func (t *Collection) RemoveNameString(index int) {
	copy(t.name[index:], t.name[index+1:])
	t.name[len(t.name)-1] = nil
	t.name = t.name[:len(t.name)-1]

}

// IsNameLangString determines whether the call to GetNameLangString is safe for the specified index
func (t *Collection) IsNameLangString(index int) (ok bool) {
	return t.name[index].langString != nil

}

// GetNameLangString returns the value safely if IsNameLangString returned true for the specified index
func (t *Collection) GetNameLangString(index int) (v string) {
	return *t.name[index].langString

}

// AppendNameLangString adds to the back of name a string type
func (t *Collection) AppendNameLangString(v string) {
	t.name = append(t.name, &nameCollectionIntermediateType{langString: &v})

}

// PrependNameLangString adds to the front of name a string type
func (t *Collection) PrependNameLangString(v string) {
	t.name = append([]*nameCollectionIntermediateType{&nameCollectionIntermediateType{langString: &v}}, t.name...)

}

// RemoveNameLangString deletes the value from the specified index
func (t *Collection) RemoveNameLangString(index int) {
	copy(t.name[index:], t.name[index+1:])
	t.name[len(t.name)-1] = nil
	t.name = t.name[:len(t.name)-1]

}

// IsNameIRI determines whether the call to GetNameIRI is safe for the specified index
func (t *Collection) IsNameIRI(index int) (ok bool) {
	return t.name[index].IRI != nil

}

// GetNameIRI returns the value safely if IsNameIRI returned true for the specified index
func (t *Collection) GetNameIRI(index int) (v *url.URL) {
	return t.name[index].IRI

}

// AppendNameIRI adds to the back of name a *url.URL type
func (t *Collection) AppendNameIRI(v *url.URL) {
	t.name = append(t.name, &nameCollectionIntermediateType{IRI: v})

}

// PrependNameIRI adds to the front of name a *url.URL type
func (t *Collection) PrependNameIRI(v *url.URL) {
	t.name = append([]*nameCollectionIntermediateType{&nameCollectionIntermediateType{IRI: v}}, t.name...)

}

// RemoveNameIRI deletes the value from the specified index
func (t *Collection) RemoveNameIRI(index int) {
	copy(t.name[index:], t.name[index+1:])
	t.name[len(t.name)-1] = nil
	t.name = t.name[:len(t.name)-1]

}

// HasUnknownName determines whether the call to GetUnknownName is safe
func (t *Collection) HasUnknownName() (ok bool) {
	return t.name != nil && t.name[0].unknown_ != nil

}

// GetUnknownName returns the unknown value for name
func (t *Collection) GetUnknownName() (v interface{}) {
	return t.name[0].unknown_

}

// SetUnknownName sets the unknown value of name
func (t *Collection) SetUnknownName(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &nameCollectionIntermediateType{}
	tmp.unknown_ = i
	t.name = append(t.name, tmp)

}

// NameMapLanguages returns all languages for this property's language mapping, or nil if there are none.
func (t *Collection) NameMapLanguages() (l []string) {
	if t.nameMap == nil || len(t.nameMap) == 0 {
		return nil
	}
	for k := range t.nameMap {
		l = append(l, k)
	}
	return

}

// GetNameMap retrieves the value of the property for the specified language, or an empty string if it does not exist
func (t *Collection) GetNameMap(l string) (v string) {
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
func (t *Collection) SetNameMap(l string, v string) {
	if t.nameMap == nil {
		t.nameMap = make(map[string]string)
	}
	t.nameMap[l] = v

}

// IsEndTime determines whether the call to GetEndTime is safe
func (t *Collection) IsEndTime() (ok bool) {
	return t.endTime != nil && t.endTime.dateTime != nil

}

// GetEndTime returns the value safely if IsEndTime returned true
func (t *Collection) GetEndTime() (v time.Time) {
	return *t.endTime.dateTime

}

// SetEndTime sets the value of endTime to be of time.Time type
func (t *Collection) SetEndTime(v time.Time) {
	t.endTime = &endTimeCollectionIntermediateType{dateTime: &v}

}

// IsEndTimeIRI determines whether the call to GetEndTimeIRI is safe
func (t *Collection) IsEndTimeIRI() (ok bool) {
	return t.endTime != nil && t.endTime.IRI != nil

}

// GetEndTimeIRI returns the value safely if IsEndTimeIRI returned true
func (t *Collection) GetEndTimeIRI() (v *url.URL) {
	return t.endTime.IRI

}

// SetEndTimeIRI sets the value of endTime to be of *url.URL type
func (t *Collection) SetEndTimeIRI(v *url.URL) {
	t.endTime = &endTimeCollectionIntermediateType{IRI: v}

}

// HasUnknownEndTime determines whether the call to GetUnknownEndTime is safe
func (t *Collection) HasUnknownEndTime() (ok bool) {
	return t.endTime != nil && t.endTime.unknown_ != nil

}

// GetUnknownEndTime returns the unknown value for endTime
func (t *Collection) GetUnknownEndTime() (v interface{}) {
	return t.endTime.unknown_

}

// SetUnknownEndTime sets the unknown value of endTime
func (t *Collection) SetUnknownEndTime(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &endTimeCollectionIntermediateType{}
	tmp.unknown_ = i
	t.endTime = tmp

}

// GeneratorLen determines the number of elements able to be used for the IsGeneratorObject, GetGeneratorObject, and RemoveGeneratorObject functions
func (t *Collection) GeneratorLen() (l int) {
	return len(t.generator)

}

// IsGeneratorObject determines whether the call to GetGeneratorObject is safe for the specified index
func (t *Collection) IsGeneratorObject(index int) (ok bool) {
	return t.generator[index].Object != nil

}

// GetGeneratorObject returns the value safely if IsGeneratorObject returned true for the specified index
func (t *Collection) GetGeneratorObject(index int) (v ObjectType) {
	return t.generator[index].Object

}

// AppendGeneratorObject adds to the back of generator a ObjectType type
func (t *Collection) AppendGeneratorObject(v ObjectType) {
	t.generator = append(t.generator, &generatorCollectionIntermediateType{Object: v})

}

// PrependGeneratorObject adds to the front of generator a ObjectType type
func (t *Collection) PrependGeneratorObject(v ObjectType) {
	t.generator = append([]*generatorCollectionIntermediateType{&generatorCollectionIntermediateType{Object: v}}, t.generator...)

}

// RemoveGeneratorObject deletes the value from the specified index
func (t *Collection) RemoveGeneratorObject(index int) {
	copy(t.generator[index:], t.generator[index+1:])
	t.generator[len(t.generator)-1] = nil
	t.generator = t.generator[:len(t.generator)-1]

}

// IsGeneratorLink determines whether the call to GetGeneratorLink is safe for the specified index
func (t *Collection) IsGeneratorLink(index int) (ok bool) {
	return t.generator[index].Link != nil

}

// GetGeneratorLink returns the value safely if IsGeneratorLink returned true for the specified index
func (t *Collection) GetGeneratorLink(index int) (v LinkType) {
	return t.generator[index].Link

}

// AppendGeneratorLink adds to the back of generator a LinkType type
func (t *Collection) AppendGeneratorLink(v LinkType) {
	t.generator = append(t.generator, &generatorCollectionIntermediateType{Link: v})

}

// PrependGeneratorLink adds to the front of generator a LinkType type
func (t *Collection) PrependGeneratorLink(v LinkType) {
	t.generator = append([]*generatorCollectionIntermediateType{&generatorCollectionIntermediateType{Link: v}}, t.generator...)

}

// RemoveGeneratorLink deletes the value from the specified index
func (t *Collection) RemoveGeneratorLink(index int) {
	copy(t.generator[index:], t.generator[index+1:])
	t.generator[len(t.generator)-1] = nil
	t.generator = t.generator[:len(t.generator)-1]

}

// IsGeneratorIRI determines whether the call to GetGeneratorIRI is safe for the specified index
func (t *Collection) IsGeneratorIRI(index int) (ok bool) {
	return t.generator[index].IRI != nil

}

// GetGeneratorIRI returns the value safely if IsGeneratorIRI returned true for the specified index
func (t *Collection) GetGeneratorIRI(index int) (v *url.URL) {
	return t.generator[index].IRI

}

// AppendGeneratorIRI adds to the back of generator a *url.URL type
func (t *Collection) AppendGeneratorIRI(v *url.URL) {
	t.generator = append(t.generator, &generatorCollectionIntermediateType{IRI: v})

}

// PrependGeneratorIRI adds to the front of generator a *url.URL type
func (t *Collection) PrependGeneratorIRI(v *url.URL) {
	t.generator = append([]*generatorCollectionIntermediateType{&generatorCollectionIntermediateType{IRI: v}}, t.generator...)

}

// RemoveGeneratorIRI deletes the value from the specified index
func (t *Collection) RemoveGeneratorIRI(index int) {
	copy(t.generator[index:], t.generator[index+1:])
	t.generator[len(t.generator)-1] = nil
	t.generator = t.generator[:len(t.generator)-1]

}

// HasUnknownGenerator determines whether the call to GetUnknownGenerator is safe
func (t *Collection) HasUnknownGenerator() (ok bool) {
	return t.generator != nil && t.generator[0].unknown_ != nil

}

// GetUnknownGenerator returns the unknown value for generator
func (t *Collection) GetUnknownGenerator() (v interface{}) {
	return t.generator[0].unknown_

}

// SetUnknownGenerator sets the unknown value of generator
func (t *Collection) SetUnknownGenerator(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &generatorCollectionIntermediateType{}
	tmp.unknown_ = i
	t.generator = append(t.generator, tmp)

}

// IconLen determines the number of elements able to be used for the IsIconImage, GetIconImage, and RemoveIconImage functions
func (t *Collection) IconLen() (l int) {
	return len(t.icon)

}

// IsIconImage determines whether the call to GetIconImage is safe for the specified index
func (t *Collection) IsIconImage(index int) (ok bool) {
	return t.icon[index].Image != nil

}

// GetIconImage returns the value safely if IsIconImage returned true for the specified index
func (t *Collection) GetIconImage(index int) (v ImageType) {
	return t.icon[index].Image

}

// AppendIconImage adds to the back of icon a ImageType type
func (t *Collection) AppendIconImage(v ImageType) {
	t.icon = append(t.icon, &iconCollectionIntermediateType{Image: v})

}

// PrependIconImage adds to the front of icon a ImageType type
func (t *Collection) PrependIconImage(v ImageType) {
	t.icon = append([]*iconCollectionIntermediateType{&iconCollectionIntermediateType{Image: v}}, t.icon...)

}

// RemoveIconImage deletes the value from the specified index
func (t *Collection) RemoveIconImage(index int) {
	copy(t.icon[index:], t.icon[index+1:])
	t.icon[len(t.icon)-1] = nil
	t.icon = t.icon[:len(t.icon)-1]

}

// IsIconLink determines whether the call to GetIconLink is safe for the specified index
func (t *Collection) IsIconLink(index int) (ok bool) {
	return t.icon[index].Link != nil

}

// GetIconLink returns the value safely if IsIconLink returned true for the specified index
func (t *Collection) GetIconLink(index int) (v LinkType) {
	return t.icon[index].Link

}

// AppendIconLink adds to the back of icon a LinkType type
func (t *Collection) AppendIconLink(v LinkType) {
	t.icon = append(t.icon, &iconCollectionIntermediateType{Link: v})

}

// PrependIconLink adds to the front of icon a LinkType type
func (t *Collection) PrependIconLink(v LinkType) {
	t.icon = append([]*iconCollectionIntermediateType{&iconCollectionIntermediateType{Link: v}}, t.icon...)

}

// RemoveIconLink deletes the value from the specified index
func (t *Collection) RemoveIconLink(index int) {
	copy(t.icon[index:], t.icon[index+1:])
	t.icon[len(t.icon)-1] = nil
	t.icon = t.icon[:len(t.icon)-1]

}

// IsIconIRI determines whether the call to GetIconIRI is safe for the specified index
func (t *Collection) IsIconIRI(index int) (ok bool) {
	return t.icon[index].IRI != nil

}

// GetIconIRI returns the value safely if IsIconIRI returned true for the specified index
func (t *Collection) GetIconIRI(index int) (v *url.URL) {
	return t.icon[index].IRI

}

// AppendIconIRI adds to the back of icon a *url.URL type
func (t *Collection) AppendIconIRI(v *url.URL) {
	t.icon = append(t.icon, &iconCollectionIntermediateType{IRI: v})

}

// PrependIconIRI adds to the front of icon a *url.URL type
func (t *Collection) PrependIconIRI(v *url.URL) {
	t.icon = append([]*iconCollectionIntermediateType{&iconCollectionIntermediateType{IRI: v}}, t.icon...)

}

// RemoveIconIRI deletes the value from the specified index
func (t *Collection) RemoveIconIRI(index int) {
	copy(t.icon[index:], t.icon[index+1:])
	t.icon[len(t.icon)-1] = nil
	t.icon = t.icon[:len(t.icon)-1]

}

// HasUnknownIcon determines whether the call to GetUnknownIcon is safe
func (t *Collection) HasUnknownIcon() (ok bool) {
	return t.icon != nil && t.icon[0].unknown_ != nil

}

// GetUnknownIcon returns the unknown value for icon
func (t *Collection) GetUnknownIcon() (v interface{}) {
	return t.icon[0].unknown_

}

// SetUnknownIcon sets the unknown value of icon
func (t *Collection) SetUnknownIcon(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &iconCollectionIntermediateType{}
	tmp.unknown_ = i
	t.icon = append(t.icon, tmp)

}

// HasId determines whether the call to GetId is safe
func (t *Collection) HasId() (ok bool) {
	return t.id != nil

}

// GetId returns the value for id
func (t *Collection) GetId() (v *url.URL) {
	return t.id

}

// SetId sets the value of id
func (t *Collection) SetId(v *url.URL) {
	t.id = v

}

// HasUnknownId determines whether the call to GetUnknownId is safe
func (t *Collection) HasUnknownId() (ok bool) {
	return t.unknown_ != nil && t.unknown_["id"] != nil

}

// GetUnknownId returns the unknown value for id
func (t *Collection) GetUnknownId() (v interface{}) {
	return t.unknown_["id"]

}

// SetUnknownId sets the unknown value of id
func (t *Collection) SetUnknownId(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["id"] = i

}

// ImageLen determines the number of elements able to be used for the IsImageImage, GetImageImage, and RemoveImageImage functions
func (t *Collection) ImageLen() (l int) {
	return len(t.image)

}

// IsImageImage determines whether the call to GetImageImage is safe for the specified index
func (t *Collection) IsImageImage(index int) (ok bool) {
	return t.image[index].Image != nil

}

// GetImageImage returns the value safely if IsImageImage returned true for the specified index
func (t *Collection) GetImageImage(index int) (v ImageType) {
	return t.image[index].Image

}

// AppendImageImage adds to the back of image a ImageType type
func (t *Collection) AppendImageImage(v ImageType) {
	t.image = append(t.image, &imageCollectionIntermediateType{Image: v})

}

// PrependImageImage adds to the front of image a ImageType type
func (t *Collection) PrependImageImage(v ImageType) {
	t.image = append([]*imageCollectionIntermediateType{&imageCollectionIntermediateType{Image: v}}, t.image...)

}

// RemoveImageImage deletes the value from the specified index
func (t *Collection) RemoveImageImage(index int) {
	copy(t.image[index:], t.image[index+1:])
	t.image[len(t.image)-1] = nil
	t.image = t.image[:len(t.image)-1]

}

// IsImageLink determines whether the call to GetImageLink is safe for the specified index
func (t *Collection) IsImageLink(index int) (ok bool) {
	return t.image[index].Link != nil

}

// GetImageLink returns the value safely if IsImageLink returned true for the specified index
func (t *Collection) GetImageLink(index int) (v LinkType) {
	return t.image[index].Link

}

// AppendImageLink adds to the back of image a LinkType type
func (t *Collection) AppendImageLink(v LinkType) {
	t.image = append(t.image, &imageCollectionIntermediateType{Link: v})

}

// PrependImageLink adds to the front of image a LinkType type
func (t *Collection) PrependImageLink(v LinkType) {
	t.image = append([]*imageCollectionIntermediateType{&imageCollectionIntermediateType{Link: v}}, t.image...)

}

// RemoveImageLink deletes the value from the specified index
func (t *Collection) RemoveImageLink(index int) {
	copy(t.image[index:], t.image[index+1:])
	t.image[len(t.image)-1] = nil
	t.image = t.image[:len(t.image)-1]

}

// IsImageIRI determines whether the call to GetImageIRI is safe for the specified index
func (t *Collection) IsImageIRI(index int) (ok bool) {
	return t.image[index].IRI != nil

}

// GetImageIRI returns the value safely if IsImageIRI returned true for the specified index
func (t *Collection) GetImageIRI(index int) (v *url.URL) {
	return t.image[index].IRI

}

// AppendImageIRI adds to the back of image a *url.URL type
func (t *Collection) AppendImageIRI(v *url.URL) {
	t.image = append(t.image, &imageCollectionIntermediateType{IRI: v})

}

// PrependImageIRI adds to the front of image a *url.URL type
func (t *Collection) PrependImageIRI(v *url.URL) {
	t.image = append([]*imageCollectionIntermediateType{&imageCollectionIntermediateType{IRI: v}}, t.image...)

}

// RemoveImageIRI deletes the value from the specified index
func (t *Collection) RemoveImageIRI(index int) {
	copy(t.image[index:], t.image[index+1:])
	t.image[len(t.image)-1] = nil
	t.image = t.image[:len(t.image)-1]

}

// HasUnknownImage determines whether the call to GetUnknownImage is safe
func (t *Collection) HasUnknownImage() (ok bool) {
	return t.image != nil && t.image[0].unknown_ != nil

}

// GetUnknownImage returns the unknown value for image
func (t *Collection) GetUnknownImage() (v interface{}) {
	return t.image[0].unknown_

}

// SetUnknownImage sets the unknown value of image
func (t *Collection) SetUnknownImage(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &imageCollectionIntermediateType{}
	tmp.unknown_ = i
	t.image = append(t.image, tmp)

}

// InReplyToLen determines the number of elements able to be used for the IsInReplyToObject, GetInReplyToObject, and RemoveInReplyToObject functions
func (t *Collection) InReplyToLen() (l int) {
	return len(t.inReplyTo)

}

// IsInReplyToObject determines whether the call to GetInReplyToObject is safe for the specified index
func (t *Collection) IsInReplyToObject(index int) (ok bool) {
	return t.inReplyTo[index].Object != nil

}

// GetInReplyToObject returns the value safely if IsInReplyToObject returned true for the specified index
func (t *Collection) GetInReplyToObject(index int) (v ObjectType) {
	return t.inReplyTo[index].Object

}

// AppendInReplyToObject adds to the back of inReplyTo a ObjectType type
func (t *Collection) AppendInReplyToObject(v ObjectType) {
	t.inReplyTo = append(t.inReplyTo, &inReplyToCollectionIntermediateType{Object: v})

}

// PrependInReplyToObject adds to the front of inReplyTo a ObjectType type
func (t *Collection) PrependInReplyToObject(v ObjectType) {
	t.inReplyTo = append([]*inReplyToCollectionIntermediateType{&inReplyToCollectionIntermediateType{Object: v}}, t.inReplyTo...)

}

// RemoveInReplyToObject deletes the value from the specified index
func (t *Collection) RemoveInReplyToObject(index int) {
	copy(t.inReplyTo[index:], t.inReplyTo[index+1:])
	t.inReplyTo[len(t.inReplyTo)-1] = nil
	t.inReplyTo = t.inReplyTo[:len(t.inReplyTo)-1]

}

// IsInReplyToLink determines whether the call to GetInReplyToLink is safe for the specified index
func (t *Collection) IsInReplyToLink(index int) (ok bool) {
	return t.inReplyTo[index].Link != nil

}

// GetInReplyToLink returns the value safely if IsInReplyToLink returned true for the specified index
func (t *Collection) GetInReplyToLink(index int) (v LinkType) {
	return t.inReplyTo[index].Link

}

// AppendInReplyToLink adds to the back of inReplyTo a LinkType type
func (t *Collection) AppendInReplyToLink(v LinkType) {
	t.inReplyTo = append(t.inReplyTo, &inReplyToCollectionIntermediateType{Link: v})

}

// PrependInReplyToLink adds to the front of inReplyTo a LinkType type
func (t *Collection) PrependInReplyToLink(v LinkType) {
	t.inReplyTo = append([]*inReplyToCollectionIntermediateType{&inReplyToCollectionIntermediateType{Link: v}}, t.inReplyTo...)

}

// RemoveInReplyToLink deletes the value from the specified index
func (t *Collection) RemoveInReplyToLink(index int) {
	copy(t.inReplyTo[index:], t.inReplyTo[index+1:])
	t.inReplyTo[len(t.inReplyTo)-1] = nil
	t.inReplyTo = t.inReplyTo[:len(t.inReplyTo)-1]

}

// IsInReplyToIRI determines whether the call to GetInReplyToIRI is safe for the specified index
func (t *Collection) IsInReplyToIRI(index int) (ok bool) {
	return t.inReplyTo[index].IRI != nil

}

// GetInReplyToIRI returns the value safely if IsInReplyToIRI returned true for the specified index
func (t *Collection) GetInReplyToIRI(index int) (v *url.URL) {
	return t.inReplyTo[index].IRI

}

// AppendInReplyToIRI adds to the back of inReplyTo a *url.URL type
func (t *Collection) AppendInReplyToIRI(v *url.URL) {
	t.inReplyTo = append(t.inReplyTo, &inReplyToCollectionIntermediateType{IRI: v})

}

// PrependInReplyToIRI adds to the front of inReplyTo a *url.URL type
func (t *Collection) PrependInReplyToIRI(v *url.URL) {
	t.inReplyTo = append([]*inReplyToCollectionIntermediateType{&inReplyToCollectionIntermediateType{IRI: v}}, t.inReplyTo...)

}

// RemoveInReplyToIRI deletes the value from the specified index
func (t *Collection) RemoveInReplyToIRI(index int) {
	copy(t.inReplyTo[index:], t.inReplyTo[index+1:])
	t.inReplyTo[len(t.inReplyTo)-1] = nil
	t.inReplyTo = t.inReplyTo[:len(t.inReplyTo)-1]

}

// HasUnknownInReplyTo determines whether the call to GetUnknownInReplyTo is safe
func (t *Collection) HasUnknownInReplyTo() (ok bool) {
	return t.inReplyTo != nil && t.inReplyTo[0].unknown_ != nil

}

// GetUnknownInReplyTo returns the unknown value for inReplyTo
func (t *Collection) GetUnknownInReplyTo() (v interface{}) {
	return t.inReplyTo[0].unknown_

}

// SetUnknownInReplyTo sets the unknown value of inReplyTo
func (t *Collection) SetUnknownInReplyTo(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &inReplyToCollectionIntermediateType{}
	tmp.unknown_ = i
	t.inReplyTo = append(t.inReplyTo, tmp)

}

// LocationLen determines the number of elements able to be used for the IsLocationObject, GetLocationObject, and RemoveLocationObject functions
func (t *Collection) LocationLen() (l int) {
	return len(t.location)

}

// IsLocationObject determines whether the call to GetLocationObject is safe for the specified index
func (t *Collection) IsLocationObject(index int) (ok bool) {
	return t.location[index].Object != nil

}

// GetLocationObject returns the value safely if IsLocationObject returned true for the specified index
func (t *Collection) GetLocationObject(index int) (v ObjectType) {
	return t.location[index].Object

}

// AppendLocationObject adds to the back of location a ObjectType type
func (t *Collection) AppendLocationObject(v ObjectType) {
	t.location = append(t.location, &locationCollectionIntermediateType{Object: v})

}

// PrependLocationObject adds to the front of location a ObjectType type
func (t *Collection) PrependLocationObject(v ObjectType) {
	t.location = append([]*locationCollectionIntermediateType{&locationCollectionIntermediateType{Object: v}}, t.location...)

}

// RemoveLocationObject deletes the value from the specified index
func (t *Collection) RemoveLocationObject(index int) {
	copy(t.location[index:], t.location[index+1:])
	t.location[len(t.location)-1] = nil
	t.location = t.location[:len(t.location)-1]

}

// IsLocationLink determines whether the call to GetLocationLink is safe for the specified index
func (t *Collection) IsLocationLink(index int) (ok bool) {
	return t.location[index].Link != nil

}

// GetLocationLink returns the value safely if IsLocationLink returned true for the specified index
func (t *Collection) GetLocationLink(index int) (v LinkType) {
	return t.location[index].Link

}

// AppendLocationLink adds to the back of location a LinkType type
func (t *Collection) AppendLocationLink(v LinkType) {
	t.location = append(t.location, &locationCollectionIntermediateType{Link: v})

}

// PrependLocationLink adds to the front of location a LinkType type
func (t *Collection) PrependLocationLink(v LinkType) {
	t.location = append([]*locationCollectionIntermediateType{&locationCollectionIntermediateType{Link: v}}, t.location...)

}

// RemoveLocationLink deletes the value from the specified index
func (t *Collection) RemoveLocationLink(index int) {
	copy(t.location[index:], t.location[index+1:])
	t.location[len(t.location)-1] = nil
	t.location = t.location[:len(t.location)-1]

}

// IsLocationIRI determines whether the call to GetLocationIRI is safe for the specified index
func (t *Collection) IsLocationIRI(index int) (ok bool) {
	return t.location[index].IRI != nil

}

// GetLocationIRI returns the value safely if IsLocationIRI returned true for the specified index
func (t *Collection) GetLocationIRI(index int) (v *url.URL) {
	return t.location[index].IRI

}

// AppendLocationIRI adds to the back of location a *url.URL type
func (t *Collection) AppendLocationIRI(v *url.URL) {
	t.location = append(t.location, &locationCollectionIntermediateType{IRI: v})

}

// PrependLocationIRI adds to the front of location a *url.URL type
func (t *Collection) PrependLocationIRI(v *url.URL) {
	t.location = append([]*locationCollectionIntermediateType{&locationCollectionIntermediateType{IRI: v}}, t.location...)

}

// RemoveLocationIRI deletes the value from the specified index
func (t *Collection) RemoveLocationIRI(index int) {
	copy(t.location[index:], t.location[index+1:])
	t.location[len(t.location)-1] = nil
	t.location = t.location[:len(t.location)-1]

}

// HasUnknownLocation determines whether the call to GetUnknownLocation is safe
func (t *Collection) HasUnknownLocation() (ok bool) {
	return t.location != nil && t.location[0].unknown_ != nil

}

// GetUnknownLocation returns the unknown value for location
func (t *Collection) GetUnknownLocation() (v interface{}) {
	return t.location[0].unknown_

}

// SetUnknownLocation sets the unknown value of location
func (t *Collection) SetUnknownLocation(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &locationCollectionIntermediateType{}
	tmp.unknown_ = i
	t.location = append(t.location, tmp)

}

// PreviewLen determines the number of elements able to be used for the IsPreviewObject, GetPreviewObject, and RemovePreviewObject functions
func (t *Collection) PreviewLen() (l int) {
	return len(t.preview)

}

// IsPreviewObject determines whether the call to GetPreviewObject is safe for the specified index
func (t *Collection) IsPreviewObject(index int) (ok bool) {
	return t.preview[index].Object != nil

}

// GetPreviewObject returns the value safely if IsPreviewObject returned true for the specified index
func (t *Collection) GetPreviewObject(index int) (v ObjectType) {
	return t.preview[index].Object

}

// AppendPreviewObject adds to the back of preview a ObjectType type
func (t *Collection) AppendPreviewObject(v ObjectType) {
	t.preview = append(t.preview, &previewCollectionIntermediateType{Object: v})

}

// PrependPreviewObject adds to the front of preview a ObjectType type
func (t *Collection) PrependPreviewObject(v ObjectType) {
	t.preview = append([]*previewCollectionIntermediateType{&previewCollectionIntermediateType{Object: v}}, t.preview...)

}

// RemovePreviewObject deletes the value from the specified index
func (t *Collection) RemovePreviewObject(index int) {
	copy(t.preview[index:], t.preview[index+1:])
	t.preview[len(t.preview)-1] = nil
	t.preview = t.preview[:len(t.preview)-1]

}

// IsPreviewLink determines whether the call to GetPreviewLink is safe for the specified index
func (t *Collection) IsPreviewLink(index int) (ok bool) {
	return t.preview[index].Link != nil

}

// GetPreviewLink returns the value safely if IsPreviewLink returned true for the specified index
func (t *Collection) GetPreviewLink(index int) (v LinkType) {
	return t.preview[index].Link

}

// AppendPreviewLink adds to the back of preview a LinkType type
func (t *Collection) AppendPreviewLink(v LinkType) {
	t.preview = append(t.preview, &previewCollectionIntermediateType{Link: v})

}

// PrependPreviewLink adds to the front of preview a LinkType type
func (t *Collection) PrependPreviewLink(v LinkType) {
	t.preview = append([]*previewCollectionIntermediateType{&previewCollectionIntermediateType{Link: v}}, t.preview...)

}

// RemovePreviewLink deletes the value from the specified index
func (t *Collection) RemovePreviewLink(index int) {
	copy(t.preview[index:], t.preview[index+1:])
	t.preview[len(t.preview)-1] = nil
	t.preview = t.preview[:len(t.preview)-1]

}

// IsPreviewIRI determines whether the call to GetPreviewIRI is safe for the specified index
func (t *Collection) IsPreviewIRI(index int) (ok bool) {
	return t.preview[index].IRI != nil

}

// GetPreviewIRI returns the value safely if IsPreviewIRI returned true for the specified index
func (t *Collection) GetPreviewIRI(index int) (v *url.URL) {
	return t.preview[index].IRI

}

// AppendPreviewIRI adds to the back of preview a *url.URL type
func (t *Collection) AppendPreviewIRI(v *url.URL) {
	t.preview = append(t.preview, &previewCollectionIntermediateType{IRI: v})

}

// PrependPreviewIRI adds to the front of preview a *url.URL type
func (t *Collection) PrependPreviewIRI(v *url.URL) {
	t.preview = append([]*previewCollectionIntermediateType{&previewCollectionIntermediateType{IRI: v}}, t.preview...)

}

// RemovePreviewIRI deletes the value from the specified index
func (t *Collection) RemovePreviewIRI(index int) {
	copy(t.preview[index:], t.preview[index+1:])
	t.preview[len(t.preview)-1] = nil
	t.preview = t.preview[:len(t.preview)-1]

}

// HasUnknownPreview determines whether the call to GetUnknownPreview is safe
func (t *Collection) HasUnknownPreview() (ok bool) {
	return t.preview != nil && t.preview[0].unknown_ != nil

}

// GetUnknownPreview returns the unknown value for preview
func (t *Collection) GetUnknownPreview() (v interface{}) {
	return t.preview[0].unknown_

}

// SetUnknownPreview sets the unknown value of preview
func (t *Collection) SetUnknownPreview(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &previewCollectionIntermediateType{}
	tmp.unknown_ = i
	t.preview = append(t.preview, tmp)

}

// IsPublished determines whether the call to GetPublished is safe
func (t *Collection) IsPublished() (ok bool) {
	return t.published != nil && t.published.dateTime != nil

}

// GetPublished returns the value safely if IsPublished returned true
func (t *Collection) GetPublished() (v time.Time) {
	return *t.published.dateTime

}

// SetPublished sets the value of published to be of time.Time type
func (t *Collection) SetPublished(v time.Time) {
	t.published = &publishedCollectionIntermediateType{dateTime: &v}

}

// IsPublishedIRI determines whether the call to GetPublishedIRI is safe
func (t *Collection) IsPublishedIRI() (ok bool) {
	return t.published != nil && t.published.IRI != nil

}

// GetPublishedIRI returns the value safely if IsPublishedIRI returned true
func (t *Collection) GetPublishedIRI() (v *url.URL) {
	return t.published.IRI

}

// SetPublishedIRI sets the value of published to be of *url.URL type
func (t *Collection) SetPublishedIRI(v *url.URL) {
	t.published = &publishedCollectionIntermediateType{IRI: v}

}

// HasUnknownPublished determines whether the call to GetUnknownPublished is safe
func (t *Collection) HasUnknownPublished() (ok bool) {
	return t.published != nil && t.published.unknown_ != nil

}

// GetUnknownPublished returns the unknown value for published
func (t *Collection) GetUnknownPublished() (v interface{}) {
	return t.published.unknown_

}

// SetUnknownPublished sets the unknown value of published
func (t *Collection) SetUnknownPublished(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &publishedCollectionIntermediateType{}
	tmp.unknown_ = i
	t.published = tmp

}

// IsReplies determines whether the call to GetReplies is safe
func (t *Collection) IsReplies() (ok bool) {
	return t.replies != nil && t.replies.Collection != nil

}

// GetReplies returns the value safely if IsReplies returned true
func (t *Collection) GetReplies() (v CollectionType) {
	return t.replies.Collection

}

// SetReplies sets the value of replies to be of CollectionType type
func (t *Collection) SetReplies(v CollectionType) {
	t.replies = &repliesCollectionIntermediateType{Collection: v}

}

// IsRepliesIRI determines whether the call to GetRepliesIRI is safe
func (t *Collection) IsRepliesIRI() (ok bool) {
	return t.replies != nil && t.replies.IRI != nil

}

// GetRepliesIRI returns the value safely if IsRepliesIRI returned true
func (t *Collection) GetRepliesIRI() (v *url.URL) {
	return t.replies.IRI

}

// SetRepliesIRI sets the value of replies to be of *url.URL type
func (t *Collection) SetRepliesIRI(v *url.URL) {
	t.replies = &repliesCollectionIntermediateType{IRI: v}

}

// HasUnknownReplies determines whether the call to GetUnknownReplies is safe
func (t *Collection) HasUnknownReplies() (ok bool) {
	return t.replies != nil && t.replies.unknown_ != nil

}

// GetUnknownReplies returns the unknown value for replies
func (t *Collection) GetUnknownReplies() (v interface{}) {
	return t.replies.unknown_

}

// SetUnknownReplies sets the unknown value of replies
func (t *Collection) SetUnknownReplies(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &repliesCollectionIntermediateType{}
	tmp.unknown_ = i
	t.replies = tmp

}

// IsStartTime determines whether the call to GetStartTime is safe
func (t *Collection) IsStartTime() (ok bool) {
	return t.startTime != nil && t.startTime.dateTime != nil

}

// GetStartTime returns the value safely if IsStartTime returned true
func (t *Collection) GetStartTime() (v time.Time) {
	return *t.startTime.dateTime

}

// SetStartTime sets the value of startTime to be of time.Time type
func (t *Collection) SetStartTime(v time.Time) {
	t.startTime = &startTimeCollectionIntermediateType{dateTime: &v}

}

// IsStartTimeIRI determines whether the call to GetStartTimeIRI is safe
func (t *Collection) IsStartTimeIRI() (ok bool) {
	return t.startTime != nil && t.startTime.IRI != nil

}

// GetStartTimeIRI returns the value safely if IsStartTimeIRI returned true
func (t *Collection) GetStartTimeIRI() (v *url.URL) {
	return t.startTime.IRI

}

// SetStartTimeIRI sets the value of startTime to be of *url.URL type
func (t *Collection) SetStartTimeIRI(v *url.URL) {
	t.startTime = &startTimeCollectionIntermediateType{IRI: v}

}

// HasUnknownStartTime determines whether the call to GetUnknownStartTime is safe
func (t *Collection) HasUnknownStartTime() (ok bool) {
	return t.startTime != nil && t.startTime.unknown_ != nil

}

// GetUnknownStartTime returns the unknown value for startTime
func (t *Collection) GetUnknownStartTime() (v interface{}) {
	return t.startTime.unknown_

}

// SetUnknownStartTime sets the unknown value of startTime
func (t *Collection) SetUnknownStartTime(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &startTimeCollectionIntermediateType{}
	tmp.unknown_ = i
	t.startTime = tmp

}

// SummaryLen determines the number of elements able to be used for the IsSummaryString, GetSummaryString, and RemoveSummaryString functions
func (t *Collection) SummaryLen() (l int) {
	return len(t.summary)

}

// IsSummaryString determines whether the call to GetSummaryString is safe for the specified index
func (t *Collection) IsSummaryString(index int) (ok bool) {
	return t.summary[index].stringName != nil

}

// GetSummaryString returns the value safely if IsSummaryString returned true for the specified index
func (t *Collection) GetSummaryString(index int) (v string) {
	return *t.summary[index].stringName

}

// AppendSummaryString adds to the back of summary a string type
func (t *Collection) AppendSummaryString(v string) {
	t.summary = append(t.summary, &summaryCollectionIntermediateType{stringName: &v})

}

// PrependSummaryString adds to the front of summary a string type
func (t *Collection) PrependSummaryString(v string) {
	t.summary = append([]*summaryCollectionIntermediateType{&summaryCollectionIntermediateType{stringName: &v}}, t.summary...)

}

// RemoveSummaryString deletes the value from the specified index
func (t *Collection) RemoveSummaryString(index int) {
	copy(t.summary[index:], t.summary[index+1:])
	t.summary[len(t.summary)-1] = nil
	t.summary = t.summary[:len(t.summary)-1]

}

// IsSummaryLangString determines whether the call to GetSummaryLangString is safe for the specified index
func (t *Collection) IsSummaryLangString(index int) (ok bool) {
	return t.summary[index].langString != nil

}

// GetSummaryLangString returns the value safely if IsSummaryLangString returned true for the specified index
func (t *Collection) GetSummaryLangString(index int) (v string) {
	return *t.summary[index].langString

}

// AppendSummaryLangString adds to the back of summary a string type
func (t *Collection) AppendSummaryLangString(v string) {
	t.summary = append(t.summary, &summaryCollectionIntermediateType{langString: &v})

}

// PrependSummaryLangString adds to the front of summary a string type
func (t *Collection) PrependSummaryLangString(v string) {
	t.summary = append([]*summaryCollectionIntermediateType{&summaryCollectionIntermediateType{langString: &v}}, t.summary...)

}

// RemoveSummaryLangString deletes the value from the specified index
func (t *Collection) RemoveSummaryLangString(index int) {
	copy(t.summary[index:], t.summary[index+1:])
	t.summary[len(t.summary)-1] = nil
	t.summary = t.summary[:len(t.summary)-1]

}

// IsSummaryIRI determines whether the call to GetSummaryIRI is safe for the specified index
func (t *Collection) IsSummaryIRI(index int) (ok bool) {
	return t.summary[index].IRI != nil

}

// GetSummaryIRI returns the value safely if IsSummaryIRI returned true for the specified index
func (t *Collection) GetSummaryIRI(index int) (v *url.URL) {
	return t.summary[index].IRI

}

// AppendSummaryIRI adds to the back of summary a *url.URL type
func (t *Collection) AppendSummaryIRI(v *url.URL) {
	t.summary = append(t.summary, &summaryCollectionIntermediateType{IRI: v})

}

// PrependSummaryIRI adds to the front of summary a *url.URL type
func (t *Collection) PrependSummaryIRI(v *url.URL) {
	t.summary = append([]*summaryCollectionIntermediateType{&summaryCollectionIntermediateType{IRI: v}}, t.summary...)

}

// RemoveSummaryIRI deletes the value from the specified index
func (t *Collection) RemoveSummaryIRI(index int) {
	copy(t.summary[index:], t.summary[index+1:])
	t.summary[len(t.summary)-1] = nil
	t.summary = t.summary[:len(t.summary)-1]

}

// HasUnknownSummary determines whether the call to GetUnknownSummary is safe
func (t *Collection) HasUnknownSummary() (ok bool) {
	return t.summary != nil && t.summary[0].unknown_ != nil

}

// GetUnknownSummary returns the unknown value for summary
func (t *Collection) GetUnknownSummary() (v interface{}) {
	return t.summary[0].unknown_

}

// SetUnknownSummary sets the unknown value of summary
func (t *Collection) SetUnknownSummary(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &summaryCollectionIntermediateType{}
	tmp.unknown_ = i
	t.summary = append(t.summary, tmp)

}

// SummaryMapLanguages returns all languages for this property's language mapping, or nil if there are none.
func (t *Collection) SummaryMapLanguages() (l []string) {
	if t.summaryMap == nil || len(t.summaryMap) == 0 {
		return nil
	}
	for k := range t.summaryMap {
		l = append(l, k)
	}
	return

}

// GetSummaryMap retrieves the value of the property for the specified language, or an empty string if it does not exist
func (t *Collection) GetSummaryMap(l string) (v string) {
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
func (t *Collection) SetSummaryMap(l string, v string) {
	if t.summaryMap == nil {
		t.summaryMap = make(map[string]string)
	}
	t.summaryMap[l] = v

}

// TagLen determines the number of elements able to be used for the IsTagObject, GetTagObject, and RemoveTagObject functions
func (t *Collection) TagLen() (l int) {
	return len(t.tag)

}

// IsTagObject determines whether the call to GetTagObject is safe for the specified index
func (t *Collection) IsTagObject(index int) (ok bool) {
	return t.tag[index].Object != nil

}

// GetTagObject returns the value safely if IsTagObject returned true for the specified index
func (t *Collection) GetTagObject(index int) (v ObjectType) {
	return t.tag[index].Object

}

// AppendTagObject adds to the back of tag a ObjectType type
func (t *Collection) AppendTagObject(v ObjectType) {
	t.tag = append(t.tag, &tagCollectionIntermediateType{Object: v})

}

// PrependTagObject adds to the front of tag a ObjectType type
func (t *Collection) PrependTagObject(v ObjectType) {
	t.tag = append([]*tagCollectionIntermediateType{&tagCollectionIntermediateType{Object: v}}, t.tag...)

}

// RemoveTagObject deletes the value from the specified index
func (t *Collection) RemoveTagObject(index int) {
	copy(t.tag[index:], t.tag[index+1:])
	t.tag[len(t.tag)-1] = nil
	t.tag = t.tag[:len(t.tag)-1]

}

// IsTagLink determines whether the call to GetTagLink is safe for the specified index
func (t *Collection) IsTagLink(index int) (ok bool) {
	return t.tag[index].Link != nil

}

// GetTagLink returns the value safely if IsTagLink returned true for the specified index
func (t *Collection) GetTagLink(index int) (v LinkType) {
	return t.tag[index].Link

}

// AppendTagLink adds to the back of tag a LinkType type
func (t *Collection) AppendTagLink(v LinkType) {
	t.tag = append(t.tag, &tagCollectionIntermediateType{Link: v})

}

// PrependTagLink adds to the front of tag a LinkType type
func (t *Collection) PrependTagLink(v LinkType) {
	t.tag = append([]*tagCollectionIntermediateType{&tagCollectionIntermediateType{Link: v}}, t.tag...)

}

// RemoveTagLink deletes the value from the specified index
func (t *Collection) RemoveTagLink(index int) {
	copy(t.tag[index:], t.tag[index+1:])
	t.tag[len(t.tag)-1] = nil
	t.tag = t.tag[:len(t.tag)-1]

}

// IsTagIRI determines whether the call to GetTagIRI is safe for the specified index
func (t *Collection) IsTagIRI(index int) (ok bool) {
	return t.tag[index].IRI != nil

}

// GetTagIRI returns the value safely if IsTagIRI returned true for the specified index
func (t *Collection) GetTagIRI(index int) (v *url.URL) {
	return t.tag[index].IRI

}

// AppendTagIRI adds to the back of tag a *url.URL type
func (t *Collection) AppendTagIRI(v *url.URL) {
	t.tag = append(t.tag, &tagCollectionIntermediateType{IRI: v})

}

// PrependTagIRI adds to the front of tag a *url.URL type
func (t *Collection) PrependTagIRI(v *url.URL) {
	t.tag = append([]*tagCollectionIntermediateType{&tagCollectionIntermediateType{IRI: v}}, t.tag...)

}

// RemoveTagIRI deletes the value from the specified index
func (t *Collection) RemoveTagIRI(index int) {
	copy(t.tag[index:], t.tag[index+1:])
	t.tag[len(t.tag)-1] = nil
	t.tag = t.tag[:len(t.tag)-1]

}

// HasUnknownTag determines whether the call to GetUnknownTag is safe
func (t *Collection) HasUnknownTag() (ok bool) {
	return t.tag != nil && t.tag[0].unknown_ != nil

}

// GetUnknownTag returns the unknown value for tag
func (t *Collection) GetUnknownTag() (v interface{}) {
	return t.tag[0].unknown_

}

// SetUnknownTag sets the unknown value of tag
func (t *Collection) SetUnknownTag(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &tagCollectionIntermediateType{}
	tmp.unknown_ = i
	t.tag = append(t.tag, tmp)

}

// TypeLen determines the number of elements able to be used for the GetType and RemoveType functions
func (t *Collection) TypeLen() (l int) {
	return len(t.typeName)

}

// GetType returns the value for the specified index
func (t *Collection) GetType(index int) (v interface{}) {
	return t.typeName[index]

}

// AppendType adds a value to the back of type
func (t *Collection) AppendType(v interface{}) {
	t.typeName = append(t.typeName, v)

}

// PrependType adds a value to the front of type
func (t *Collection) PrependType(v interface{}) {
	t.typeName = append([]interface{}{v}, t.typeName...)

}

// RemoveType deletes the value from the specified index
func (t *Collection) RemoveType(index int) {
	copy(t.typeName[index:], t.typeName[index+1:])
	t.typeName[len(t.typeName)-1] = nil
	t.typeName = t.typeName[:len(t.typeName)-1]

}

// IsUpdated determines whether the call to GetUpdated is safe
func (t *Collection) IsUpdated() (ok bool) {
	return t.updated != nil && t.updated.dateTime != nil

}

// GetUpdated returns the value safely if IsUpdated returned true
func (t *Collection) GetUpdated() (v time.Time) {
	return *t.updated.dateTime

}

// SetUpdated sets the value of updated to be of time.Time type
func (t *Collection) SetUpdated(v time.Time) {
	t.updated = &updatedCollectionIntermediateType{dateTime: &v}

}

// IsUpdatedIRI determines whether the call to GetUpdatedIRI is safe
func (t *Collection) IsUpdatedIRI() (ok bool) {
	return t.updated != nil && t.updated.IRI != nil

}

// GetUpdatedIRI returns the value safely if IsUpdatedIRI returned true
func (t *Collection) GetUpdatedIRI() (v *url.URL) {
	return t.updated.IRI

}

// SetUpdatedIRI sets the value of updated to be of *url.URL type
func (t *Collection) SetUpdatedIRI(v *url.URL) {
	t.updated = &updatedCollectionIntermediateType{IRI: v}

}

// HasUnknownUpdated determines whether the call to GetUnknownUpdated is safe
func (t *Collection) HasUnknownUpdated() (ok bool) {
	return t.updated != nil && t.updated.unknown_ != nil

}

// GetUnknownUpdated returns the unknown value for updated
func (t *Collection) GetUnknownUpdated() (v interface{}) {
	return t.updated.unknown_

}

// SetUnknownUpdated sets the unknown value of updated
func (t *Collection) SetUnknownUpdated(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &updatedCollectionIntermediateType{}
	tmp.unknown_ = i
	t.updated = tmp

}

// UrlLen determines the number of elements able to be used for the IsUrlAnyURI, GetUrlAnyURI, and RemoveUrlAnyURI functions
func (t *Collection) UrlLen() (l int) {
	return len(t.url)

}

// IsUrlAnyURI determines whether the call to GetUrlAnyURI is safe for the specified index
func (t *Collection) IsUrlAnyURI(index int) (ok bool) {
	return t.url[index].anyURI != nil

}

// GetUrlAnyURI returns the value safely if IsUrlAnyURI returned true for the specified index
func (t *Collection) GetUrlAnyURI(index int) (v *url.URL) {
	return t.url[index].anyURI

}

// AppendUrlAnyURI adds to the back of url a *url.URL type
func (t *Collection) AppendUrlAnyURI(v *url.URL) {
	t.url = append(t.url, &urlCollectionIntermediateType{anyURI: v})

}

// PrependUrlAnyURI adds to the front of url a *url.URL type
func (t *Collection) PrependUrlAnyURI(v *url.URL) {
	t.url = append([]*urlCollectionIntermediateType{&urlCollectionIntermediateType{anyURI: v}}, t.url...)

}

// RemoveUrlAnyURI deletes the value from the specified index
func (t *Collection) RemoveUrlAnyURI(index int) {
	copy(t.url[index:], t.url[index+1:])
	t.url[len(t.url)-1] = nil
	t.url = t.url[:len(t.url)-1]

}

// IsUrlLink determines whether the call to GetUrlLink is safe for the specified index
func (t *Collection) IsUrlLink(index int) (ok bool) {
	return t.url[index].Link != nil

}

// GetUrlLink returns the value safely if IsUrlLink returned true for the specified index
func (t *Collection) GetUrlLink(index int) (v LinkType) {
	return t.url[index].Link

}

// AppendUrlLink adds to the back of url a LinkType type
func (t *Collection) AppendUrlLink(v LinkType) {
	t.url = append(t.url, &urlCollectionIntermediateType{Link: v})

}

// PrependUrlLink adds to the front of url a LinkType type
func (t *Collection) PrependUrlLink(v LinkType) {
	t.url = append([]*urlCollectionIntermediateType{&urlCollectionIntermediateType{Link: v}}, t.url...)

}

// RemoveUrlLink deletes the value from the specified index
func (t *Collection) RemoveUrlLink(index int) {
	copy(t.url[index:], t.url[index+1:])
	t.url[len(t.url)-1] = nil
	t.url = t.url[:len(t.url)-1]

}

// HasUnknownUrl determines whether the call to GetUnknownUrl is safe
func (t *Collection) HasUnknownUrl() (ok bool) {
	return t.url != nil && t.url[0].unknown_ != nil

}

// GetUnknownUrl returns the unknown value for url
func (t *Collection) GetUnknownUrl() (v interface{}) {
	return t.url[0].unknown_

}

// SetUnknownUrl sets the unknown value of url
func (t *Collection) SetUnknownUrl(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &urlCollectionIntermediateType{}
	tmp.unknown_ = i
	t.url = append(t.url, tmp)

}

// ToLen determines the number of elements able to be used for the IsToObject, GetToObject, and RemoveToObject functions
func (t *Collection) ToLen() (l int) {
	return len(t.to)

}

// IsToObject determines whether the call to GetToObject is safe for the specified index
func (t *Collection) IsToObject(index int) (ok bool) {
	return t.to[index].Object != nil

}

// GetToObject returns the value safely if IsToObject returned true for the specified index
func (t *Collection) GetToObject(index int) (v ObjectType) {
	return t.to[index].Object

}

// AppendToObject adds to the back of to a ObjectType type
func (t *Collection) AppendToObject(v ObjectType) {
	t.to = append(t.to, &toCollectionIntermediateType{Object: v})

}

// PrependToObject adds to the front of to a ObjectType type
func (t *Collection) PrependToObject(v ObjectType) {
	t.to = append([]*toCollectionIntermediateType{&toCollectionIntermediateType{Object: v}}, t.to...)

}

// RemoveToObject deletes the value from the specified index
func (t *Collection) RemoveToObject(index int) {
	copy(t.to[index:], t.to[index+1:])
	t.to[len(t.to)-1] = nil
	t.to = t.to[:len(t.to)-1]

}

// IsToLink determines whether the call to GetToLink is safe for the specified index
func (t *Collection) IsToLink(index int) (ok bool) {
	return t.to[index].Link != nil

}

// GetToLink returns the value safely if IsToLink returned true for the specified index
func (t *Collection) GetToLink(index int) (v LinkType) {
	return t.to[index].Link

}

// AppendToLink adds to the back of to a LinkType type
func (t *Collection) AppendToLink(v LinkType) {
	t.to = append(t.to, &toCollectionIntermediateType{Link: v})

}

// PrependToLink adds to the front of to a LinkType type
func (t *Collection) PrependToLink(v LinkType) {
	t.to = append([]*toCollectionIntermediateType{&toCollectionIntermediateType{Link: v}}, t.to...)

}

// RemoveToLink deletes the value from the specified index
func (t *Collection) RemoveToLink(index int) {
	copy(t.to[index:], t.to[index+1:])
	t.to[len(t.to)-1] = nil
	t.to = t.to[:len(t.to)-1]

}

// IsToIRI determines whether the call to GetToIRI is safe for the specified index
func (t *Collection) IsToIRI(index int) (ok bool) {
	return t.to[index].IRI != nil

}

// GetToIRI returns the value safely if IsToIRI returned true for the specified index
func (t *Collection) GetToIRI(index int) (v *url.URL) {
	return t.to[index].IRI

}

// AppendToIRI adds to the back of to a *url.URL type
func (t *Collection) AppendToIRI(v *url.URL) {
	t.to = append(t.to, &toCollectionIntermediateType{IRI: v})

}

// PrependToIRI adds to the front of to a *url.URL type
func (t *Collection) PrependToIRI(v *url.URL) {
	t.to = append([]*toCollectionIntermediateType{&toCollectionIntermediateType{IRI: v}}, t.to...)

}

// RemoveToIRI deletes the value from the specified index
func (t *Collection) RemoveToIRI(index int) {
	copy(t.to[index:], t.to[index+1:])
	t.to[len(t.to)-1] = nil
	t.to = t.to[:len(t.to)-1]

}

// HasUnknownTo determines whether the call to GetUnknownTo is safe
func (t *Collection) HasUnknownTo() (ok bool) {
	return t.to != nil && t.to[0].unknown_ != nil

}

// GetUnknownTo returns the unknown value for to
func (t *Collection) GetUnknownTo() (v interface{}) {
	return t.to[0].unknown_

}

// SetUnknownTo sets the unknown value of to
func (t *Collection) SetUnknownTo(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &toCollectionIntermediateType{}
	tmp.unknown_ = i
	t.to = append(t.to, tmp)

}

// BtoLen determines the number of elements able to be used for the IsBtoObject, GetBtoObject, and RemoveBtoObject functions
func (t *Collection) BtoLen() (l int) {
	return len(t.bto)

}

// IsBtoObject determines whether the call to GetBtoObject is safe for the specified index
func (t *Collection) IsBtoObject(index int) (ok bool) {
	return t.bto[index].Object != nil

}

// GetBtoObject returns the value safely if IsBtoObject returned true for the specified index
func (t *Collection) GetBtoObject(index int) (v ObjectType) {
	return t.bto[index].Object

}

// AppendBtoObject adds to the back of bto a ObjectType type
func (t *Collection) AppendBtoObject(v ObjectType) {
	t.bto = append(t.bto, &btoCollectionIntermediateType{Object: v})

}

// PrependBtoObject adds to the front of bto a ObjectType type
func (t *Collection) PrependBtoObject(v ObjectType) {
	t.bto = append([]*btoCollectionIntermediateType{&btoCollectionIntermediateType{Object: v}}, t.bto...)

}

// RemoveBtoObject deletes the value from the specified index
func (t *Collection) RemoveBtoObject(index int) {
	copy(t.bto[index:], t.bto[index+1:])
	t.bto[len(t.bto)-1] = nil
	t.bto = t.bto[:len(t.bto)-1]

}

// IsBtoLink determines whether the call to GetBtoLink is safe for the specified index
func (t *Collection) IsBtoLink(index int) (ok bool) {
	return t.bto[index].Link != nil

}

// GetBtoLink returns the value safely if IsBtoLink returned true for the specified index
func (t *Collection) GetBtoLink(index int) (v LinkType) {
	return t.bto[index].Link

}

// AppendBtoLink adds to the back of bto a LinkType type
func (t *Collection) AppendBtoLink(v LinkType) {
	t.bto = append(t.bto, &btoCollectionIntermediateType{Link: v})

}

// PrependBtoLink adds to the front of bto a LinkType type
func (t *Collection) PrependBtoLink(v LinkType) {
	t.bto = append([]*btoCollectionIntermediateType{&btoCollectionIntermediateType{Link: v}}, t.bto...)

}

// RemoveBtoLink deletes the value from the specified index
func (t *Collection) RemoveBtoLink(index int) {
	copy(t.bto[index:], t.bto[index+1:])
	t.bto[len(t.bto)-1] = nil
	t.bto = t.bto[:len(t.bto)-1]

}

// IsBtoIRI determines whether the call to GetBtoIRI is safe for the specified index
func (t *Collection) IsBtoIRI(index int) (ok bool) {
	return t.bto[index].IRI != nil

}

// GetBtoIRI returns the value safely if IsBtoIRI returned true for the specified index
func (t *Collection) GetBtoIRI(index int) (v *url.URL) {
	return t.bto[index].IRI

}

// AppendBtoIRI adds to the back of bto a *url.URL type
func (t *Collection) AppendBtoIRI(v *url.URL) {
	t.bto = append(t.bto, &btoCollectionIntermediateType{IRI: v})

}

// PrependBtoIRI adds to the front of bto a *url.URL type
func (t *Collection) PrependBtoIRI(v *url.URL) {
	t.bto = append([]*btoCollectionIntermediateType{&btoCollectionIntermediateType{IRI: v}}, t.bto...)

}

// RemoveBtoIRI deletes the value from the specified index
func (t *Collection) RemoveBtoIRI(index int) {
	copy(t.bto[index:], t.bto[index+1:])
	t.bto[len(t.bto)-1] = nil
	t.bto = t.bto[:len(t.bto)-1]

}

// HasUnknownBto determines whether the call to GetUnknownBto is safe
func (t *Collection) HasUnknownBto() (ok bool) {
	return t.bto != nil && t.bto[0].unknown_ != nil

}

// GetUnknownBto returns the unknown value for bto
func (t *Collection) GetUnknownBto() (v interface{}) {
	return t.bto[0].unknown_

}

// SetUnknownBto sets the unknown value of bto
func (t *Collection) SetUnknownBto(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &btoCollectionIntermediateType{}
	tmp.unknown_ = i
	t.bto = append(t.bto, tmp)

}

// CcLen determines the number of elements able to be used for the IsCcObject, GetCcObject, and RemoveCcObject functions
func (t *Collection) CcLen() (l int) {
	return len(t.cc)

}

// IsCcObject determines whether the call to GetCcObject is safe for the specified index
func (t *Collection) IsCcObject(index int) (ok bool) {
	return t.cc[index].Object != nil

}

// GetCcObject returns the value safely if IsCcObject returned true for the specified index
func (t *Collection) GetCcObject(index int) (v ObjectType) {
	return t.cc[index].Object

}

// AppendCcObject adds to the back of cc a ObjectType type
func (t *Collection) AppendCcObject(v ObjectType) {
	t.cc = append(t.cc, &ccCollectionIntermediateType{Object: v})

}

// PrependCcObject adds to the front of cc a ObjectType type
func (t *Collection) PrependCcObject(v ObjectType) {
	t.cc = append([]*ccCollectionIntermediateType{&ccCollectionIntermediateType{Object: v}}, t.cc...)

}

// RemoveCcObject deletes the value from the specified index
func (t *Collection) RemoveCcObject(index int) {
	copy(t.cc[index:], t.cc[index+1:])
	t.cc[len(t.cc)-1] = nil
	t.cc = t.cc[:len(t.cc)-1]

}

// IsCcLink determines whether the call to GetCcLink is safe for the specified index
func (t *Collection) IsCcLink(index int) (ok bool) {
	return t.cc[index].Link != nil

}

// GetCcLink returns the value safely if IsCcLink returned true for the specified index
func (t *Collection) GetCcLink(index int) (v LinkType) {
	return t.cc[index].Link

}

// AppendCcLink adds to the back of cc a LinkType type
func (t *Collection) AppendCcLink(v LinkType) {
	t.cc = append(t.cc, &ccCollectionIntermediateType{Link: v})

}

// PrependCcLink adds to the front of cc a LinkType type
func (t *Collection) PrependCcLink(v LinkType) {
	t.cc = append([]*ccCollectionIntermediateType{&ccCollectionIntermediateType{Link: v}}, t.cc...)

}

// RemoveCcLink deletes the value from the specified index
func (t *Collection) RemoveCcLink(index int) {
	copy(t.cc[index:], t.cc[index+1:])
	t.cc[len(t.cc)-1] = nil
	t.cc = t.cc[:len(t.cc)-1]

}

// IsCcIRI determines whether the call to GetCcIRI is safe for the specified index
func (t *Collection) IsCcIRI(index int) (ok bool) {
	return t.cc[index].IRI != nil

}

// GetCcIRI returns the value safely if IsCcIRI returned true for the specified index
func (t *Collection) GetCcIRI(index int) (v *url.URL) {
	return t.cc[index].IRI

}

// AppendCcIRI adds to the back of cc a *url.URL type
func (t *Collection) AppendCcIRI(v *url.URL) {
	t.cc = append(t.cc, &ccCollectionIntermediateType{IRI: v})

}

// PrependCcIRI adds to the front of cc a *url.URL type
func (t *Collection) PrependCcIRI(v *url.URL) {
	t.cc = append([]*ccCollectionIntermediateType{&ccCollectionIntermediateType{IRI: v}}, t.cc...)

}

// RemoveCcIRI deletes the value from the specified index
func (t *Collection) RemoveCcIRI(index int) {
	copy(t.cc[index:], t.cc[index+1:])
	t.cc[len(t.cc)-1] = nil
	t.cc = t.cc[:len(t.cc)-1]

}

// HasUnknownCc determines whether the call to GetUnknownCc is safe
func (t *Collection) HasUnknownCc() (ok bool) {
	return t.cc != nil && t.cc[0].unknown_ != nil

}

// GetUnknownCc returns the unknown value for cc
func (t *Collection) GetUnknownCc() (v interface{}) {
	return t.cc[0].unknown_

}

// SetUnknownCc sets the unknown value of cc
func (t *Collection) SetUnknownCc(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &ccCollectionIntermediateType{}
	tmp.unknown_ = i
	t.cc = append(t.cc, tmp)

}

// BccLen determines the number of elements able to be used for the IsBccObject, GetBccObject, and RemoveBccObject functions
func (t *Collection) BccLen() (l int) {
	return len(t.bcc)

}

// IsBccObject determines whether the call to GetBccObject is safe for the specified index
func (t *Collection) IsBccObject(index int) (ok bool) {
	return t.bcc[index].Object != nil

}

// GetBccObject returns the value safely if IsBccObject returned true for the specified index
func (t *Collection) GetBccObject(index int) (v ObjectType) {
	return t.bcc[index].Object

}

// AppendBccObject adds to the back of bcc a ObjectType type
func (t *Collection) AppendBccObject(v ObjectType) {
	t.bcc = append(t.bcc, &bccCollectionIntermediateType{Object: v})

}

// PrependBccObject adds to the front of bcc a ObjectType type
func (t *Collection) PrependBccObject(v ObjectType) {
	t.bcc = append([]*bccCollectionIntermediateType{&bccCollectionIntermediateType{Object: v}}, t.bcc...)

}

// RemoveBccObject deletes the value from the specified index
func (t *Collection) RemoveBccObject(index int) {
	copy(t.bcc[index:], t.bcc[index+1:])
	t.bcc[len(t.bcc)-1] = nil
	t.bcc = t.bcc[:len(t.bcc)-1]

}

// IsBccLink determines whether the call to GetBccLink is safe for the specified index
func (t *Collection) IsBccLink(index int) (ok bool) {
	return t.bcc[index].Link != nil

}

// GetBccLink returns the value safely if IsBccLink returned true for the specified index
func (t *Collection) GetBccLink(index int) (v LinkType) {
	return t.bcc[index].Link

}

// AppendBccLink adds to the back of bcc a LinkType type
func (t *Collection) AppendBccLink(v LinkType) {
	t.bcc = append(t.bcc, &bccCollectionIntermediateType{Link: v})

}

// PrependBccLink adds to the front of bcc a LinkType type
func (t *Collection) PrependBccLink(v LinkType) {
	t.bcc = append([]*bccCollectionIntermediateType{&bccCollectionIntermediateType{Link: v}}, t.bcc...)

}

// RemoveBccLink deletes the value from the specified index
func (t *Collection) RemoveBccLink(index int) {
	copy(t.bcc[index:], t.bcc[index+1:])
	t.bcc[len(t.bcc)-1] = nil
	t.bcc = t.bcc[:len(t.bcc)-1]

}

// IsBccIRI determines whether the call to GetBccIRI is safe for the specified index
func (t *Collection) IsBccIRI(index int) (ok bool) {
	return t.bcc[index].IRI != nil

}

// GetBccIRI returns the value safely if IsBccIRI returned true for the specified index
func (t *Collection) GetBccIRI(index int) (v *url.URL) {
	return t.bcc[index].IRI

}

// AppendBccIRI adds to the back of bcc a *url.URL type
func (t *Collection) AppendBccIRI(v *url.URL) {
	t.bcc = append(t.bcc, &bccCollectionIntermediateType{IRI: v})

}

// PrependBccIRI adds to the front of bcc a *url.URL type
func (t *Collection) PrependBccIRI(v *url.URL) {
	t.bcc = append([]*bccCollectionIntermediateType{&bccCollectionIntermediateType{IRI: v}}, t.bcc...)

}

// RemoveBccIRI deletes the value from the specified index
func (t *Collection) RemoveBccIRI(index int) {
	copy(t.bcc[index:], t.bcc[index+1:])
	t.bcc[len(t.bcc)-1] = nil
	t.bcc = t.bcc[:len(t.bcc)-1]

}

// HasUnknownBcc determines whether the call to GetUnknownBcc is safe
func (t *Collection) HasUnknownBcc() (ok bool) {
	return t.bcc != nil && t.bcc[0].unknown_ != nil

}

// GetUnknownBcc returns the unknown value for bcc
func (t *Collection) GetUnknownBcc() (v interface{}) {
	return t.bcc[0].unknown_

}

// SetUnknownBcc sets the unknown value of bcc
func (t *Collection) SetUnknownBcc(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &bccCollectionIntermediateType{}
	tmp.unknown_ = i
	t.bcc = append(t.bcc, tmp)

}

// IsMediaType determines whether the call to GetMediaType is safe
func (t *Collection) IsMediaType() (ok bool) {
	return t.mediaType != nil && t.mediaType.mimeMediaTypeValue != nil

}

// GetMediaType returns the value safely if IsMediaType returned true
func (t *Collection) GetMediaType() (v string) {
	return *t.mediaType.mimeMediaTypeValue

}

// SetMediaType sets the value of mediaType to be of string type
func (t *Collection) SetMediaType(v string) {
	t.mediaType = &mediaTypeCollectionIntermediateType{mimeMediaTypeValue: &v}

}

// IsMediaTypeIRI determines whether the call to GetMediaTypeIRI is safe
func (t *Collection) IsMediaTypeIRI() (ok bool) {
	return t.mediaType != nil && t.mediaType.IRI != nil

}

// GetMediaTypeIRI returns the value safely if IsMediaTypeIRI returned true
func (t *Collection) GetMediaTypeIRI() (v *url.URL) {
	return t.mediaType.IRI

}

// SetMediaTypeIRI sets the value of mediaType to be of *url.URL type
func (t *Collection) SetMediaTypeIRI(v *url.URL) {
	t.mediaType = &mediaTypeCollectionIntermediateType{IRI: v}

}

// HasUnknownMediaType determines whether the call to GetUnknownMediaType is safe
func (t *Collection) HasUnknownMediaType() (ok bool) {
	return t.mediaType != nil && t.mediaType.unknown_ != nil

}

// GetUnknownMediaType returns the unknown value for mediaType
func (t *Collection) GetUnknownMediaType() (v interface{}) {
	return t.mediaType.unknown_

}

// SetUnknownMediaType sets the unknown value of mediaType
func (t *Collection) SetUnknownMediaType(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &mediaTypeCollectionIntermediateType{}
	tmp.unknown_ = i
	t.mediaType = tmp

}

// IsDuration determines whether the call to GetDuration is safe
func (t *Collection) IsDuration() (ok bool) {
	return t.duration != nil && t.duration.duration != nil

}

// GetDuration returns the value safely if IsDuration returned true
func (t *Collection) GetDuration() (v time.Duration) {
	return *t.duration.duration

}

// SetDuration sets the value of duration to be of time.Duration type
func (t *Collection) SetDuration(v time.Duration) {
	t.duration = &durationCollectionIntermediateType{duration: &v}

}

// IsDurationIRI determines whether the call to GetDurationIRI is safe
func (t *Collection) IsDurationIRI() (ok bool) {
	return t.duration != nil && t.duration.IRI != nil

}

// GetDurationIRI returns the value safely if IsDurationIRI returned true
func (t *Collection) GetDurationIRI() (v *url.URL) {
	return t.duration.IRI

}

// SetDurationIRI sets the value of duration to be of *url.URL type
func (t *Collection) SetDurationIRI(v *url.URL) {
	t.duration = &durationCollectionIntermediateType{IRI: v}

}

// HasUnknownDuration determines whether the call to GetUnknownDuration is safe
func (t *Collection) HasUnknownDuration() (ok bool) {
	return t.duration != nil && t.duration.unknown_ != nil

}

// GetUnknownDuration returns the unknown value for duration
func (t *Collection) GetUnknownDuration() (v interface{}) {
	return t.duration.unknown_

}

// SetUnknownDuration sets the unknown value of duration
func (t *Collection) SetUnknownDuration(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &durationCollectionIntermediateType{}
	tmp.unknown_ = i
	t.duration = tmp

}

// IsSource determines whether the call to GetSource is safe
func (t *Collection) IsSource() (ok bool) {
	return t.source != nil && t.source.Object != nil

}

// GetSource returns the value safely if IsSource returned true
func (t *Collection) GetSource() (v ObjectType) {
	return t.source.Object

}

// SetSource sets the value of source to be of ObjectType type
func (t *Collection) SetSource(v ObjectType) {
	t.source = &sourceCollectionIntermediateType{Object: v}

}

// IsSourceIRI determines whether the call to GetSourceIRI is safe
func (t *Collection) IsSourceIRI() (ok bool) {
	return t.source != nil && t.source.IRI != nil

}

// GetSourceIRI returns the value safely if IsSourceIRI returned true
func (t *Collection) GetSourceIRI() (v *url.URL) {
	return t.source.IRI

}

// SetSourceIRI sets the value of source to be of *url.URL type
func (t *Collection) SetSourceIRI(v *url.URL) {
	t.source = &sourceCollectionIntermediateType{IRI: v}

}

// HasUnknownSource determines whether the call to GetUnknownSource is safe
func (t *Collection) HasUnknownSource() (ok bool) {
	return t.source != nil && t.source.unknown_ != nil

}

// GetUnknownSource returns the unknown value for source
func (t *Collection) GetUnknownSource() (v interface{}) {
	return t.source.unknown_

}

// SetUnknownSource sets the unknown value of source
func (t *Collection) SetUnknownSource(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &sourceCollectionIntermediateType{}
	tmp.unknown_ = i
	t.source = tmp

}

// IsInboxOrderedCollection determines whether the call to GetInboxOrderedCollection is safe
func (t *Collection) IsInboxOrderedCollection() (ok bool) {
	return t.inbox != nil && t.inbox.OrderedCollection != nil

}

// GetInboxOrderedCollection returns the value safely if IsInboxOrderedCollection returned true
func (t *Collection) GetInboxOrderedCollection() (v OrderedCollectionType) {
	return t.inbox.OrderedCollection

}

// SetInboxOrderedCollection sets the value of inbox to be of OrderedCollectionType type
func (t *Collection) SetInboxOrderedCollection(v OrderedCollectionType) {
	t.inbox = &inboxCollectionIntermediateType{OrderedCollection: v}

}

// IsInboxAnyURI determines whether the call to GetInboxAnyURI is safe
func (t *Collection) IsInboxAnyURI() (ok bool) {
	return t.inbox != nil && t.inbox.anyURI != nil

}

// GetInboxAnyURI returns the value safely if IsInboxAnyURI returned true
func (t *Collection) GetInboxAnyURI() (v *url.URL) {
	return t.inbox.anyURI

}

// SetInboxAnyURI sets the value of inbox to be of *url.URL type
func (t *Collection) SetInboxAnyURI(v *url.URL) {
	t.inbox = &inboxCollectionIntermediateType{anyURI: v}

}

// HasUnknownInbox determines whether the call to GetUnknownInbox is safe
func (t *Collection) HasUnknownInbox() (ok bool) {
	return t.inbox != nil && t.inbox.unknown_ != nil

}

// GetUnknownInbox returns the unknown value for inbox
func (t *Collection) GetUnknownInbox() (v interface{}) {
	return t.inbox.unknown_

}

// SetUnknownInbox sets the unknown value of inbox
func (t *Collection) SetUnknownInbox(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &inboxCollectionIntermediateType{}
	tmp.unknown_ = i
	t.inbox = tmp

}

// IsOutboxOrderedCollection determines whether the call to GetOutboxOrderedCollection is safe
func (t *Collection) IsOutboxOrderedCollection() (ok bool) {
	return t.outbox != nil && t.outbox.OrderedCollection != nil

}

// GetOutboxOrderedCollection returns the value safely if IsOutboxOrderedCollection returned true
func (t *Collection) GetOutboxOrderedCollection() (v OrderedCollectionType) {
	return t.outbox.OrderedCollection

}

// SetOutboxOrderedCollection sets the value of outbox to be of OrderedCollectionType type
func (t *Collection) SetOutboxOrderedCollection(v OrderedCollectionType) {
	t.outbox = &outboxCollectionIntermediateType{OrderedCollection: v}

}

// IsOutboxAnyURI determines whether the call to GetOutboxAnyURI is safe
func (t *Collection) IsOutboxAnyURI() (ok bool) {
	return t.outbox != nil && t.outbox.anyURI != nil

}

// GetOutboxAnyURI returns the value safely if IsOutboxAnyURI returned true
func (t *Collection) GetOutboxAnyURI() (v *url.URL) {
	return t.outbox.anyURI

}

// SetOutboxAnyURI sets the value of outbox to be of *url.URL type
func (t *Collection) SetOutboxAnyURI(v *url.URL) {
	t.outbox = &outboxCollectionIntermediateType{anyURI: v}

}

// HasUnknownOutbox determines whether the call to GetUnknownOutbox is safe
func (t *Collection) HasUnknownOutbox() (ok bool) {
	return t.outbox != nil && t.outbox.unknown_ != nil

}

// GetUnknownOutbox returns the unknown value for outbox
func (t *Collection) GetUnknownOutbox() (v interface{}) {
	return t.outbox.unknown_

}

// SetUnknownOutbox sets the unknown value of outbox
func (t *Collection) SetUnknownOutbox(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &outboxCollectionIntermediateType{}
	tmp.unknown_ = i
	t.outbox = tmp

}

// IsFollowingCollection determines whether the call to GetFollowingCollection is safe
func (t *Collection) IsFollowingCollection() (ok bool) {
	return t.following != nil && t.following.Collection != nil

}

// GetFollowingCollection returns the value safely if IsFollowingCollection returned true
func (t *Collection) GetFollowingCollection() (v CollectionType) {
	return t.following.Collection

}

// SetFollowingCollection sets the value of following to be of CollectionType type
func (t *Collection) SetFollowingCollection(v CollectionType) {
	t.following = &followingCollectionIntermediateType{Collection: v}

}

// IsFollowingOrderedCollection determines whether the call to GetFollowingOrderedCollection is safe
func (t *Collection) IsFollowingOrderedCollection() (ok bool) {
	return t.following != nil && t.following.OrderedCollection != nil

}

// GetFollowingOrderedCollection returns the value safely if IsFollowingOrderedCollection returned true
func (t *Collection) GetFollowingOrderedCollection() (v OrderedCollectionType) {
	return t.following.OrderedCollection

}

// SetFollowingOrderedCollection sets the value of following to be of OrderedCollectionType type
func (t *Collection) SetFollowingOrderedCollection(v OrderedCollectionType) {
	t.following = &followingCollectionIntermediateType{OrderedCollection: v}

}

// IsFollowingAnyURI determines whether the call to GetFollowingAnyURI is safe
func (t *Collection) IsFollowingAnyURI() (ok bool) {
	return t.following != nil && t.following.anyURI != nil

}

// GetFollowingAnyURI returns the value safely if IsFollowingAnyURI returned true
func (t *Collection) GetFollowingAnyURI() (v *url.URL) {
	return t.following.anyURI

}

// SetFollowingAnyURI sets the value of following to be of *url.URL type
func (t *Collection) SetFollowingAnyURI(v *url.URL) {
	t.following = &followingCollectionIntermediateType{anyURI: v}

}

// HasUnknownFollowing determines whether the call to GetUnknownFollowing is safe
func (t *Collection) HasUnknownFollowing() (ok bool) {
	return t.following != nil && t.following.unknown_ != nil

}

// GetUnknownFollowing returns the unknown value for following
func (t *Collection) GetUnknownFollowing() (v interface{}) {
	return t.following.unknown_

}

// SetUnknownFollowing sets the unknown value of following
func (t *Collection) SetUnknownFollowing(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &followingCollectionIntermediateType{}
	tmp.unknown_ = i
	t.following = tmp

}

// IsFollowersCollection determines whether the call to GetFollowersCollection is safe
func (t *Collection) IsFollowersCollection() (ok bool) {
	return t.followers != nil && t.followers.Collection != nil

}

// GetFollowersCollection returns the value safely if IsFollowersCollection returned true
func (t *Collection) GetFollowersCollection() (v CollectionType) {
	return t.followers.Collection

}

// SetFollowersCollection sets the value of followers to be of CollectionType type
func (t *Collection) SetFollowersCollection(v CollectionType) {
	t.followers = &followersCollectionIntermediateType{Collection: v}

}

// IsFollowersOrderedCollection determines whether the call to GetFollowersOrderedCollection is safe
func (t *Collection) IsFollowersOrderedCollection() (ok bool) {
	return t.followers != nil && t.followers.OrderedCollection != nil

}

// GetFollowersOrderedCollection returns the value safely if IsFollowersOrderedCollection returned true
func (t *Collection) GetFollowersOrderedCollection() (v OrderedCollectionType) {
	return t.followers.OrderedCollection

}

// SetFollowersOrderedCollection sets the value of followers to be of OrderedCollectionType type
func (t *Collection) SetFollowersOrderedCollection(v OrderedCollectionType) {
	t.followers = &followersCollectionIntermediateType{OrderedCollection: v}

}

// IsFollowersAnyURI determines whether the call to GetFollowersAnyURI is safe
func (t *Collection) IsFollowersAnyURI() (ok bool) {
	return t.followers != nil && t.followers.anyURI != nil

}

// GetFollowersAnyURI returns the value safely if IsFollowersAnyURI returned true
func (t *Collection) GetFollowersAnyURI() (v *url.URL) {
	return t.followers.anyURI

}

// SetFollowersAnyURI sets the value of followers to be of *url.URL type
func (t *Collection) SetFollowersAnyURI(v *url.URL) {
	t.followers = &followersCollectionIntermediateType{anyURI: v}

}

// HasUnknownFollowers determines whether the call to GetUnknownFollowers is safe
func (t *Collection) HasUnknownFollowers() (ok bool) {
	return t.followers != nil && t.followers.unknown_ != nil

}

// GetUnknownFollowers returns the unknown value for followers
func (t *Collection) GetUnknownFollowers() (v interface{}) {
	return t.followers.unknown_

}

// SetUnknownFollowers sets the unknown value of followers
func (t *Collection) SetUnknownFollowers(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &followersCollectionIntermediateType{}
	tmp.unknown_ = i
	t.followers = tmp

}

// IsLikedCollection determines whether the call to GetLikedCollection is safe
func (t *Collection) IsLikedCollection() (ok bool) {
	return t.liked != nil && t.liked.Collection != nil

}

// GetLikedCollection returns the value safely if IsLikedCollection returned true
func (t *Collection) GetLikedCollection() (v CollectionType) {
	return t.liked.Collection

}

// SetLikedCollection sets the value of liked to be of CollectionType type
func (t *Collection) SetLikedCollection(v CollectionType) {
	t.liked = &likedCollectionIntermediateType{Collection: v}

}

// IsLikedOrderedCollection determines whether the call to GetLikedOrderedCollection is safe
func (t *Collection) IsLikedOrderedCollection() (ok bool) {
	return t.liked != nil && t.liked.OrderedCollection != nil

}

// GetLikedOrderedCollection returns the value safely if IsLikedOrderedCollection returned true
func (t *Collection) GetLikedOrderedCollection() (v OrderedCollectionType) {
	return t.liked.OrderedCollection

}

// SetLikedOrderedCollection sets the value of liked to be of OrderedCollectionType type
func (t *Collection) SetLikedOrderedCollection(v OrderedCollectionType) {
	t.liked = &likedCollectionIntermediateType{OrderedCollection: v}

}

// IsLikedAnyURI determines whether the call to GetLikedAnyURI is safe
func (t *Collection) IsLikedAnyURI() (ok bool) {
	return t.liked != nil && t.liked.anyURI != nil

}

// GetLikedAnyURI returns the value safely if IsLikedAnyURI returned true
func (t *Collection) GetLikedAnyURI() (v *url.URL) {
	return t.liked.anyURI

}

// SetLikedAnyURI sets the value of liked to be of *url.URL type
func (t *Collection) SetLikedAnyURI(v *url.URL) {
	t.liked = &likedCollectionIntermediateType{anyURI: v}

}

// HasUnknownLiked determines whether the call to GetUnknownLiked is safe
func (t *Collection) HasUnknownLiked() (ok bool) {
	return t.liked != nil && t.liked.unknown_ != nil

}

// GetUnknownLiked returns the unknown value for liked
func (t *Collection) GetUnknownLiked() (v interface{}) {
	return t.liked.unknown_

}

// SetUnknownLiked sets the unknown value of liked
func (t *Collection) SetUnknownLiked(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &likedCollectionIntermediateType{}
	tmp.unknown_ = i
	t.liked = tmp

}

// IsLikesCollection determines whether the call to GetLikesCollection is safe
func (t *Collection) IsLikesCollection() (ok bool) {
	return t.likes != nil && t.likes.Collection != nil

}

// GetLikesCollection returns the value safely if IsLikesCollection returned true
func (t *Collection) GetLikesCollection() (v CollectionType) {
	return t.likes.Collection

}

// SetLikesCollection sets the value of likes to be of CollectionType type
func (t *Collection) SetLikesCollection(v CollectionType) {
	t.likes = &likesCollectionIntermediateType{Collection: v}

}

// IsLikesOrderedCollection determines whether the call to GetLikesOrderedCollection is safe
func (t *Collection) IsLikesOrderedCollection() (ok bool) {
	return t.likes != nil && t.likes.OrderedCollection != nil

}

// GetLikesOrderedCollection returns the value safely if IsLikesOrderedCollection returned true
func (t *Collection) GetLikesOrderedCollection() (v OrderedCollectionType) {
	return t.likes.OrderedCollection

}

// SetLikesOrderedCollection sets the value of likes to be of OrderedCollectionType type
func (t *Collection) SetLikesOrderedCollection(v OrderedCollectionType) {
	t.likes = &likesCollectionIntermediateType{OrderedCollection: v}

}

// IsLikesAnyURI determines whether the call to GetLikesAnyURI is safe
func (t *Collection) IsLikesAnyURI() (ok bool) {
	return t.likes != nil && t.likes.anyURI != nil

}

// GetLikesAnyURI returns the value safely if IsLikesAnyURI returned true
func (t *Collection) GetLikesAnyURI() (v *url.URL) {
	return t.likes.anyURI

}

// SetLikesAnyURI sets the value of likes to be of *url.URL type
func (t *Collection) SetLikesAnyURI(v *url.URL) {
	t.likes = &likesCollectionIntermediateType{anyURI: v}

}

// HasUnknownLikes determines whether the call to GetUnknownLikes is safe
func (t *Collection) HasUnknownLikes() (ok bool) {
	return t.likes != nil && t.likes.unknown_ != nil

}

// GetUnknownLikes returns the unknown value for likes
func (t *Collection) GetUnknownLikes() (v interface{}) {
	return t.likes.unknown_

}

// SetUnknownLikes sets the unknown value of likes
func (t *Collection) SetUnknownLikes(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &likesCollectionIntermediateType{}
	tmp.unknown_ = i
	t.likes = tmp

}

// StreamsLen determines the number of elements able to be used for the GetStreams and RemoveStreams functions
func (t *Collection) StreamsLen() (l int) {
	return len(t.streams)

}

// GetStreams returns the value for the specified index
func (t *Collection) GetStreams(index int) (v *url.URL) {
	return t.streams[index]

}

// AppendStreams adds a value to the back of streams
func (t *Collection) AppendStreams(v *url.URL) {
	t.streams = append(t.streams, v)

}

// PrependStreams adds a value to the front of streams
func (t *Collection) PrependStreams(v *url.URL) {
	t.streams = append([]*url.URL{v}, t.streams...)

}

// RemoveStreams deletes the value from the specified index
func (t *Collection) RemoveStreams(index int) {
	copy(t.streams[index:], t.streams[index+1:])
	t.streams[len(t.streams)-1] = nil
	t.streams = t.streams[:len(t.streams)-1]

}

// HasUnknownStreams determines whether the call to GetUnknownStreams is safe
func (t *Collection) HasUnknownStreams() (ok bool) {
	return t.unknown_ != nil && t.unknown_["streams"] != nil

}

// GetUnknownStreams returns the unknown value for streams
func (t *Collection) GetUnknownStreams() (v interface{}) {
	return t.unknown_["streams"]

}

// SetUnknownStreams sets the unknown value of streams
func (t *Collection) SetUnknownStreams(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["streams"] = i

}

// IsPreferredUsername determines whether the call to GetPreferredUsername is safe
func (t *Collection) IsPreferredUsername() (ok bool) {
	return t.preferredUsername != nil && t.preferredUsername.stringName != nil

}

// GetPreferredUsername returns the value safely if IsPreferredUsername returned true
func (t *Collection) GetPreferredUsername() (v string) {
	return *t.preferredUsername.stringName

}

// SetPreferredUsername sets the value of preferredUsername to be of string type
func (t *Collection) SetPreferredUsername(v string) {
	t.preferredUsername = &preferredUsernameCollectionIntermediateType{stringName: &v}

}

// IsPreferredUsernameIRI determines whether the call to GetPreferredUsernameIRI is safe
func (t *Collection) IsPreferredUsernameIRI() (ok bool) {
	return t.preferredUsername != nil && t.preferredUsername.IRI != nil

}

// GetPreferredUsernameIRI returns the value safely if IsPreferredUsernameIRI returned true
func (t *Collection) GetPreferredUsernameIRI() (v *url.URL) {
	return t.preferredUsername.IRI

}

// SetPreferredUsernameIRI sets the value of preferredUsername to be of *url.URL type
func (t *Collection) SetPreferredUsernameIRI(v *url.URL) {
	t.preferredUsername = &preferredUsernameCollectionIntermediateType{IRI: v}

}

// HasUnknownPreferredUsername determines whether the call to GetUnknownPreferredUsername is safe
func (t *Collection) HasUnknownPreferredUsername() (ok bool) {
	return t.preferredUsername != nil && t.preferredUsername.unknown_ != nil

}

// GetUnknownPreferredUsername returns the unknown value for preferredUsername
func (t *Collection) GetUnknownPreferredUsername() (v interface{}) {
	return t.preferredUsername.unknown_

}

// SetUnknownPreferredUsername sets the unknown value of preferredUsername
func (t *Collection) SetUnknownPreferredUsername(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &preferredUsernameCollectionIntermediateType{}
	tmp.unknown_ = i
	t.preferredUsername = tmp

}

// PreferredUsernameMapLanguages returns all languages for this property's language mapping, or nil if there are none.
func (t *Collection) PreferredUsernameMapLanguages() (l []string) {
	if t.preferredUsernameMap == nil || len(t.preferredUsernameMap) == 0 {
		return nil
	}
	for k := range t.preferredUsernameMap {
		l = append(l, k)
	}
	return

}

// GetPreferredUsernameMap retrieves the value of the property for the specified language, or an empty string if it does not exist
func (t *Collection) GetPreferredUsernameMap(l string) (v string) {
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
func (t *Collection) SetPreferredUsernameMap(l string, v string) {
	if t.preferredUsernameMap == nil {
		t.preferredUsernameMap = make(map[string]string)
	}
	t.preferredUsernameMap[l] = v

}

// IsEndpoints determines whether the call to GetEndpoints is safe
func (t *Collection) IsEndpoints() (ok bool) {
	return t.endpoints != nil && t.endpoints.Object != nil

}

// GetEndpoints returns the value safely if IsEndpoints returned true
func (t *Collection) GetEndpoints() (v ObjectType) {
	return t.endpoints.Object

}

// SetEndpoints sets the value of endpoints to be of ObjectType type
func (t *Collection) SetEndpoints(v ObjectType) {
	t.endpoints = &endpointsCollectionIntermediateType{Object: v}

}

// IsEndpointsIRI determines whether the call to GetEndpointsIRI is safe
func (t *Collection) IsEndpointsIRI() (ok bool) {
	return t.endpoints != nil && t.endpoints.IRI != nil

}

// GetEndpointsIRI returns the value safely if IsEndpointsIRI returned true
func (t *Collection) GetEndpointsIRI() (v *url.URL) {
	return t.endpoints.IRI

}

// SetEndpointsIRI sets the value of endpoints to be of *url.URL type
func (t *Collection) SetEndpointsIRI(v *url.URL) {
	t.endpoints = &endpointsCollectionIntermediateType{IRI: v}

}

// HasUnknownEndpoints determines whether the call to GetUnknownEndpoints is safe
func (t *Collection) HasUnknownEndpoints() (ok bool) {
	return t.endpoints != nil && t.endpoints.unknown_ != nil

}

// GetUnknownEndpoints returns the unknown value for endpoints
func (t *Collection) GetUnknownEndpoints() (v interface{}) {
	return t.endpoints.unknown_

}

// SetUnknownEndpoints sets the unknown value of endpoints
func (t *Collection) SetUnknownEndpoints(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &endpointsCollectionIntermediateType{}
	tmp.unknown_ = i
	t.endpoints = tmp

}

// HasProxyUrl determines whether the call to GetProxyUrl is safe
func (t *Collection) HasProxyUrl() (ok bool) {
	return t.proxyUrl != nil

}

// GetProxyUrl returns the value for proxyUrl
func (t *Collection) GetProxyUrl() (v *url.URL) {
	return t.proxyUrl

}

// SetProxyUrl sets the value of proxyUrl
func (t *Collection) SetProxyUrl(v *url.URL) {
	t.proxyUrl = v

}

// HasUnknownProxyUrl determines whether the call to GetUnknownProxyUrl is safe
func (t *Collection) HasUnknownProxyUrl() (ok bool) {
	return t.unknown_ != nil && t.unknown_["proxyUrl"] != nil

}

// GetUnknownProxyUrl returns the unknown value for proxyUrl
func (t *Collection) GetUnknownProxyUrl() (v interface{}) {
	return t.unknown_["proxyUrl"]

}

// SetUnknownProxyUrl sets the unknown value of proxyUrl
func (t *Collection) SetUnknownProxyUrl(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["proxyUrl"] = i

}

// HasOauthAuthorizationEndpoint determines whether the call to GetOauthAuthorizationEndpoint is safe
func (t *Collection) HasOauthAuthorizationEndpoint() (ok bool) {
	return t.oauthAuthorizationEndpoint != nil

}

// GetOauthAuthorizationEndpoint returns the value for oauthAuthorizationEndpoint
func (t *Collection) GetOauthAuthorizationEndpoint() (v *url.URL) {
	return t.oauthAuthorizationEndpoint

}

// SetOauthAuthorizationEndpoint sets the value of oauthAuthorizationEndpoint
func (t *Collection) SetOauthAuthorizationEndpoint(v *url.URL) {
	t.oauthAuthorizationEndpoint = v

}

// HasUnknownOauthAuthorizationEndpoint determines whether the call to GetUnknownOauthAuthorizationEndpoint is safe
func (t *Collection) HasUnknownOauthAuthorizationEndpoint() (ok bool) {
	return t.unknown_ != nil && t.unknown_["oauthAuthorizationEndpoint"] != nil

}

// GetUnknownOauthAuthorizationEndpoint returns the unknown value for oauthAuthorizationEndpoint
func (t *Collection) GetUnknownOauthAuthorizationEndpoint() (v interface{}) {
	return t.unknown_["oauthAuthorizationEndpoint"]

}

// SetUnknownOauthAuthorizationEndpoint sets the unknown value of oauthAuthorizationEndpoint
func (t *Collection) SetUnknownOauthAuthorizationEndpoint(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["oauthAuthorizationEndpoint"] = i

}

// HasOauthTokenEndpoint determines whether the call to GetOauthTokenEndpoint is safe
func (t *Collection) HasOauthTokenEndpoint() (ok bool) {
	return t.oauthTokenEndpoint != nil

}

// GetOauthTokenEndpoint returns the value for oauthTokenEndpoint
func (t *Collection) GetOauthTokenEndpoint() (v *url.URL) {
	return t.oauthTokenEndpoint

}

// SetOauthTokenEndpoint sets the value of oauthTokenEndpoint
func (t *Collection) SetOauthTokenEndpoint(v *url.URL) {
	t.oauthTokenEndpoint = v

}

// HasUnknownOauthTokenEndpoint determines whether the call to GetUnknownOauthTokenEndpoint is safe
func (t *Collection) HasUnknownOauthTokenEndpoint() (ok bool) {
	return t.unknown_ != nil && t.unknown_["oauthTokenEndpoint"] != nil

}

// GetUnknownOauthTokenEndpoint returns the unknown value for oauthTokenEndpoint
func (t *Collection) GetUnknownOauthTokenEndpoint() (v interface{}) {
	return t.unknown_["oauthTokenEndpoint"]

}

// SetUnknownOauthTokenEndpoint sets the unknown value of oauthTokenEndpoint
func (t *Collection) SetUnknownOauthTokenEndpoint(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["oauthTokenEndpoint"] = i

}

// HasProvideClientKey determines whether the call to GetProvideClientKey is safe
func (t *Collection) HasProvideClientKey() (ok bool) {
	return t.provideClientKey != nil

}

// GetProvideClientKey returns the value for provideClientKey
func (t *Collection) GetProvideClientKey() (v *url.URL) {
	return t.provideClientKey

}

// SetProvideClientKey sets the value of provideClientKey
func (t *Collection) SetProvideClientKey(v *url.URL) {
	t.provideClientKey = v

}

// HasUnknownProvideClientKey determines whether the call to GetUnknownProvideClientKey is safe
func (t *Collection) HasUnknownProvideClientKey() (ok bool) {
	return t.unknown_ != nil && t.unknown_["provideClientKey"] != nil

}

// GetUnknownProvideClientKey returns the unknown value for provideClientKey
func (t *Collection) GetUnknownProvideClientKey() (v interface{}) {
	return t.unknown_["provideClientKey"]

}

// SetUnknownProvideClientKey sets the unknown value of provideClientKey
func (t *Collection) SetUnknownProvideClientKey(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["provideClientKey"] = i

}

// HasSignClientKey determines whether the call to GetSignClientKey is safe
func (t *Collection) HasSignClientKey() (ok bool) {
	return t.signClientKey != nil

}

// GetSignClientKey returns the value for signClientKey
func (t *Collection) GetSignClientKey() (v *url.URL) {
	return t.signClientKey

}

// SetSignClientKey sets the value of signClientKey
func (t *Collection) SetSignClientKey(v *url.URL) {
	t.signClientKey = v

}

// HasUnknownSignClientKey determines whether the call to GetUnknownSignClientKey is safe
func (t *Collection) HasUnknownSignClientKey() (ok bool) {
	return t.unknown_ != nil && t.unknown_["signClientKey"] != nil

}

// GetUnknownSignClientKey returns the unknown value for signClientKey
func (t *Collection) GetUnknownSignClientKey() (v interface{}) {
	return t.unknown_["signClientKey"]

}

// SetUnknownSignClientKey sets the unknown value of signClientKey
func (t *Collection) SetUnknownSignClientKey(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["signClientKey"] = i

}

// HasSharedInbox determines whether the call to GetSharedInbox is safe
func (t *Collection) HasSharedInbox() (ok bool) {
	return t.sharedInbox != nil

}

// GetSharedInbox returns the value for sharedInbox
func (t *Collection) GetSharedInbox() (v *url.URL) {
	return t.sharedInbox

}

// SetSharedInbox sets the value of sharedInbox
func (t *Collection) SetSharedInbox(v *url.URL) {
	t.sharedInbox = v

}

// HasUnknownSharedInbox determines whether the call to GetUnknownSharedInbox is safe
func (t *Collection) HasUnknownSharedInbox() (ok bool) {
	return t.unknown_ != nil && t.unknown_["sharedInbox"] != nil

}

// GetUnknownSharedInbox returns the unknown value for sharedInbox
func (t *Collection) GetUnknownSharedInbox() (v interface{}) {
	return t.unknown_["sharedInbox"]

}

// SetUnknownSharedInbox sets the unknown value of sharedInbox
func (t *Collection) SetUnknownSharedInbox(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["sharedInbox"] = i

}

// AddUnknown adds a raw extension to this object with the specified key
func (t *Collection) AddUnknown(k string, i interface{}) (this *Collection) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_[k] = i
	return t

}

// HasUnknown returns true if there is an unknown object with the specified key
func (t *Collection) HasUnknown(k string) (b bool) {
	if t.unknown_ == nil {
		return false
	}
	_, ok := t.unknown_[k]
	return ok

}

// RemoveUnknown removes a raw extension from this object with the specified key
func (t *Collection) RemoveUnknown(k string) (this *Collection) {
	delete(t.unknown_, k)
	return t

}

// Serialize turns this object into a map[string]interface{}. Note that the "type" property will automatically be populated with "Collection" if not manually set by the caller
func (t *Collection) Serialize() (m map[string]interface{}, err error) {
	m = make(map[string]interface{})
	for k, v := range t.unknown_ {
		m[k] = unknownValueSerialize(v)
	}
	var typeAlreadySet bool
	for _, k := range t.typeName {
		if ks, ok := k.(string); ok {
			if ks == "Collection" {
				typeAlreadySet = true
				break
			}
		}
	}
	if !typeAlreadySet {
		t.typeName = append(t.typeName, "Collection")
	}
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.totalItems != nil {
		if v, err := serializeTotalItemsCollectionIntermediateType(t.totalItems); err == nil {
			m["totalItems"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.current != nil {
		if v, err := serializeCurrentCollectionIntermediateType(t.current); err == nil {
			m["current"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.first != nil {
		if v, err := serializeFirstCollectionIntermediateType(t.first); err == nil {
			m["first"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.last != nil {
		if v, err := serializeLastCollectionIntermediateType(t.last); err == nil {
			m["last"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceItemsCollectionIntermediateType(t.items); err == nil && v != nil {
		if len(v) == 1 {
			m["items"] = v[0]
		} else {
			m["items"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.altitude != nil {
		if v, err := serializeAltitudeCollectionIntermediateType(t.altitude); err == nil {
			m["altitude"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceAttachmentCollectionIntermediateType(t.attachment); err == nil && v != nil {
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
	if v, err := serializeSliceAttributedToCollectionIntermediateType(t.attributedTo); err == nil && v != nil {
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
	if v, err := serializeSliceAudienceCollectionIntermediateType(t.audience); err == nil && v != nil {
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
	if v, err := serializeSliceContentCollectionIntermediateType(t.content); err == nil && v != nil {
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
	if v, err := serializeSliceContextCollectionIntermediateType(t.context); err == nil && v != nil {
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
	if v, err := serializeSliceNameCollectionIntermediateType(t.name); err == nil && v != nil {
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
		if v, err := serializeEndTimeCollectionIntermediateType(t.endTime); err == nil {
			m["endTime"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceGeneratorCollectionIntermediateType(t.generator); err == nil && v != nil {
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
	if v, err := serializeSliceIconCollectionIntermediateType(t.icon); err == nil && v != nil {
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
	if v, err := serializeSliceImageCollectionIntermediateType(t.image); err == nil && v != nil {
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
	if v, err := serializeSliceInReplyToCollectionIntermediateType(t.inReplyTo); err == nil && v != nil {
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
	if v, err := serializeSliceLocationCollectionIntermediateType(t.location); err == nil && v != nil {
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
	if v, err := serializeSlicePreviewCollectionIntermediateType(t.preview); err == nil && v != nil {
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
		if v, err := serializePublishedCollectionIntermediateType(t.published); err == nil {
			m["published"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.replies != nil {
		if v, err := serializeRepliesCollectionIntermediateType(t.replies); err == nil {
			m["replies"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.startTime != nil {
		if v, err := serializeStartTimeCollectionIntermediateType(t.startTime); err == nil {
			m["startTime"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceSummaryCollectionIntermediateType(t.summary); err == nil && v != nil {
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
	if v, err := serializeSliceTagCollectionIntermediateType(t.tag); err == nil && v != nil {
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
		if v, err := serializeUpdatedCollectionIntermediateType(t.updated); err == nil {
			m["updated"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceUrlCollectionIntermediateType(t.url); err == nil && v != nil {
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
	if v, err := serializeSliceToCollectionIntermediateType(t.to); err == nil && v != nil {
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
	if v, err := serializeSliceBtoCollectionIntermediateType(t.bto); err == nil && v != nil {
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
	if v, err := serializeSliceCcCollectionIntermediateType(t.cc); err == nil && v != nil {
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
	if v, err := serializeSliceBccCollectionIntermediateType(t.bcc); err == nil && v != nil {
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
		if v, err := serializeMediaTypeCollectionIntermediateType(t.mediaType); err == nil {
			m["mediaType"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.duration != nil {
		if v, err := serializeDurationCollectionIntermediateType(t.duration); err == nil {
			m["duration"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.source != nil {
		if v, err := serializeSourceCollectionIntermediateType(t.source); err == nil {
			m["source"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.inbox != nil {
		if v, err := serializeInboxCollectionIntermediateType(t.inbox); err == nil {
			m["inbox"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.outbox != nil {
		if v, err := serializeOutboxCollectionIntermediateType(t.outbox); err == nil {
			m["outbox"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.following != nil {
		if v, err := serializeFollowingCollectionIntermediateType(t.following); err == nil {
			m["following"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.followers != nil {
		if v, err := serializeFollowersCollectionIntermediateType(t.followers); err == nil {
			m["followers"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.liked != nil {
		if v, err := serializeLikedCollectionIntermediateType(t.liked); err == nil {
			m["liked"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.likes != nil {
		if v, err := serializeLikesCollectionIntermediateType(t.likes); err == nil {
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
		if v, err := serializePreferredUsernameCollectionIntermediateType(t.preferredUsername); err == nil {
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
		if v, err := serializeEndpointsCollectionIntermediateType(t.endpoints); err == nil {
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
func (t *Collection) Deserialize(m map[string]interface{}) (err error) {
	for k, v := range m {
		handled := false
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "totalItems" {
				t.totalItems, err = deserializeTotalItemsCollectionIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "current" {
				t.current, err = deserializeCurrentCollectionIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "first" {
				t.first, err = deserializeFirstCollectionIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "last" {
				t.last, err = deserializeLastCollectionIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "items" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeItemsCollectionIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.items = []*itemsCollectionIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.items, err = deserializeSliceItemsCollectionIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &itemsCollectionIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.items = []*itemsCollectionIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "altitude" {
				t.altitude, err = deserializeAltitudeCollectionIntermediateType(v)
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
					tmp, err := deserializeAttachmentCollectionIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.attachment = []*attachmentCollectionIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.attachment, err = deserializeSliceAttachmentCollectionIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &attachmentCollectionIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.attachment = []*attachmentCollectionIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "attributedTo" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeAttributedToCollectionIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.attributedTo = []*attributedToCollectionIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.attributedTo, err = deserializeSliceAttributedToCollectionIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &attributedToCollectionIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.attributedTo = []*attributedToCollectionIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "audience" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeAudienceCollectionIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.audience = []*audienceCollectionIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.audience, err = deserializeSliceAudienceCollectionIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &audienceCollectionIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.audience = []*audienceCollectionIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "content" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeContentCollectionIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.content = []*contentCollectionIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.content, err = deserializeSliceContentCollectionIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &contentCollectionIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.content = []*contentCollectionIntermediateType{tmp}
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
					tmp, err := deserializeContextCollectionIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.context = []*contextCollectionIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.context, err = deserializeSliceContextCollectionIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &contextCollectionIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.context = []*contextCollectionIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "name" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeNameCollectionIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.name = []*nameCollectionIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.name, err = deserializeSliceNameCollectionIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &nameCollectionIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.name = []*nameCollectionIntermediateType{tmp}
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
				t.endTime, err = deserializeEndTimeCollectionIntermediateType(v)
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
					tmp, err := deserializeGeneratorCollectionIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.generator = []*generatorCollectionIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.generator, err = deserializeSliceGeneratorCollectionIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &generatorCollectionIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.generator = []*generatorCollectionIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "icon" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeIconCollectionIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.icon = []*iconCollectionIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.icon, err = deserializeSliceIconCollectionIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &iconCollectionIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.icon = []*iconCollectionIntermediateType{tmp}
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
					tmp, err := deserializeImageCollectionIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.image = []*imageCollectionIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.image, err = deserializeSliceImageCollectionIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &imageCollectionIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.image = []*imageCollectionIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "inReplyTo" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeInReplyToCollectionIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.inReplyTo = []*inReplyToCollectionIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.inReplyTo, err = deserializeSliceInReplyToCollectionIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &inReplyToCollectionIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.inReplyTo = []*inReplyToCollectionIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "location" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeLocationCollectionIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.location = []*locationCollectionIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.location, err = deserializeSliceLocationCollectionIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &locationCollectionIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.location = []*locationCollectionIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "preview" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializePreviewCollectionIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.preview = []*previewCollectionIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.preview, err = deserializeSlicePreviewCollectionIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &previewCollectionIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.preview = []*previewCollectionIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "published" {
				t.published, err = deserializePublishedCollectionIntermediateType(v)
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
				t.replies, err = deserializeRepliesCollectionIntermediateType(v)
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
				t.startTime, err = deserializeStartTimeCollectionIntermediateType(v)
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
					tmp, err := deserializeSummaryCollectionIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.summary = []*summaryCollectionIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.summary, err = deserializeSliceSummaryCollectionIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &summaryCollectionIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.summary = []*summaryCollectionIntermediateType{tmp}
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
					tmp, err := deserializeTagCollectionIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.tag = []*tagCollectionIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.tag, err = deserializeSliceTagCollectionIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &tagCollectionIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.tag = []*tagCollectionIntermediateType{tmp}
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
				t.updated, err = deserializeUpdatedCollectionIntermediateType(v)
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
					tmp, err := deserializeUrlCollectionIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.url = []*urlCollectionIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.url, err = deserializeSliceUrlCollectionIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &urlCollectionIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.url = []*urlCollectionIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "to" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeToCollectionIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.to = []*toCollectionIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.to, err = deserializeSliceToCollectionIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &toCollectionIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.to = []*toCollectionIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "bto" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeBtoCollectionIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.bto = []*btoCollectionIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.bto, err = deserializeSliceBtoCollectionIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &btoCollectionIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.bto = []*btoCollectionIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "cc" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeCcCollectionIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.cc = []*ccCollectionIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.cc, err = deserializeSliceCcCollectionIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &ccCollectionIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.cc = []*ccCollectionIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "bcc" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeBccCollectionIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.bcc = []*bccCollectionIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.bcc, err = deserializeSliceBccCollectionIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &bccCollectionIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.bcc = []*bccCollectionIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "mediaType" {
				t.mediaType, err = deserializeMediaTypeCollectionIntermediateType(v)
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
				t.duration, err = deserializeDurationCollectionIntermediateType(v)
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
				t.source, err = deserializeSourceCollectionIntermediateType(v)
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
				t.inbox, err = deserializeInboxCollectionIntermediateType(v)
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
				t.outbox, err = deserializeOutboxCollectionIntermediateType(v)
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
				t.following, err = deserializeFollowingCollectionIntermediateType(v)
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
				t.followers, err = deserializeFollowersCollectionIntermediateType(v)
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
				t.liked, err = deserializeLikedCollectionIntermediateType(v)
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
				t.likes, err = deserializeLikesCollectionIntermediateType(v)
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
				t.preferredUsername, err = deserializePreferredUsernameCollectionIntermediateType(v)
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
				t.endpoints, err = deserializeEndpointsCollectionIntermediateType(v)
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

// totalItemsCollectionIntermediateType will only have one of its values set at most
type totalItemsCollectionIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *int64 type for totalItems property
	nonNegativeInteger *int64
	// Stores possible *url.URL type for totalItems property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *totalItemsCollectionIntermediateType) Deserialize(i interface{}) (err error) {
	matched := false
	if m, ok := i.(map[string]interface{}); ok {
		err = fmt.Errorf("Given map but nothing to do with it for this type: %v", m)
	} else if i != nil {
		if !matched {
			t.nonNegativeInteger, err = nonNegativeIntegerDeserialize(i)
			if err != nil {
				t.nonNegativeInteger = nil
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
func (t *totalItemsCollectionIntermediateType) Serialize() (i interface{}, err error) {
	if t.nonNegativeInteger != nil {
		i = nonNegativeIntegerSerialize(*t.nonNegativeInteger)
		return
	}
	if t.IRI != nil {
		i = IRISerialize(t.IRI)
		return
	}
	i = unknownValueSerialize(t.unknown_)
	return
}

// currentCollectionIntermediateType will only have one of its values set at most
type currentCollectionIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible CollectionPageType type for current property
	CollectionPage CollectionPageType
	// Stores possible LinkType type for current property
	Link LinkType
	// Stores possible *url.URL type for current property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *currentCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
					if t.CollectionPage, ok = resolveObject(kind).(CollectionPageType); t.CollectionPage != nil && ok {
						err = t.CollectionPage.Deserialize(m)
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
func (t *currentCollectionIntermediateType) Serialize() (i interface{}, err error) {
	if t.CollectionPage != nil {
		i, err = t.CollectionPage.Serialize()
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

// firstCollectionIntermediateType will only have one of its values set at most
type firstCollectionIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible CollectionPageType type for first property
	CollectionPage CollectionPageType
	// Stores possible LinkType type for first property
	Link LinkType
	// Stores possible *url.URL type for first property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *firstCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
					if t.CollectionPage, ok = resolveObject(kind).(CollectionPageType); t.CollectionPage != nil && ok {
						err = t.CollectionPage.Deserialize(m)
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
func (t *firstCollectionIntermediateType) Serialize() (i interface{}, err error) {
	if t.CollectionPage != nil {
		i, err = t.CollectionPage.Serialize()
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

// lastCollectionIntermediateType will only have one of its values set at most
type lastCollectionIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible CollectionPageType type for last property
	CollectionPage CollectionPageType
	// Stores possible LinkType type for last property
	Link LinkType
	// Stores possible *url.URL type for last property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *lastCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
					if t.CollectionPage, ok = resolveObject(kind).(CollectionPageType); t.CollectionPage != nil && ok {
						err = t.CollectionPage.Deserialize(m)
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
func (t *lastCollectionIntermediateType) Serialize() (i interface{}, err error) {
	if t.CollectionPage != nil {
		i, err = t.CollectionPage.Serialize()
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

// itemsCollectionIntermediateType will only have one of its values set at most
type itemsCollectionIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for items property
	Object ObjectType
	// Stores possible LinkType type for items property
	Link LinkType
	// Stores possible *url.URL type for items property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *itemsCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *itemsCollectionIntermediateType) Serialize() (i interface{}, err error) {
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

// altitudeCollectionIntermediateType will only have one of its values set at most
type altitudeCollectionIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *float64 type for altitude property
	float *float64
	// Stores possible *url.URL type for altitude property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *altitudeCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *altitudeCollectionIntermediateType) Serialize() (i interface{}, err error) {
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

// attachmentCollectionIntermediateType will only have one of its values set at most
type attachmentCollectionIntermediateType struct {
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
func (t *attachmentCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *attachmentCollectionIntermediateType) Serialize() (i interface{}, err error) {
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

// attributedToCollectionIntermediateType will only have one of its values set at most
type attributedToCollectionIntermediateType struct {
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
func (t *attributedToCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *attributedToCollectionIntermediateType) Serialize() (i interface{}, err error) {
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

// audienceCollectionIntermediateType will only have one of its values set at most
type audienceCollectionIntermediateType struct {
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
func (t *audienceCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *audienceCollectionIntermediateType) Serialize() (i interface{}, err error) {
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

// contentCollectionIntermediateType will only have one of its values set at most
type contentCollectionIntermediateType struct {
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
func (t *contentCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *contentCollectionIntermediateType) Serialize() (i interface{}, err error) {
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

// contextCollectionIntermediateType will only have one of its values set at most
type contextCollectionIntermediateType struct {
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
func (t *contextCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *contextCollectionIntermediateType) Serialize() (i interface{}, err error) {
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

// nameCollectionIntermediateType will only have one of its values set at most
type nameCollectionIntermediateType struct {
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
func (t *nameCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *nameCollectionIntermediateType) Serialize() (i interface{}, err error) {
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

// endTimeCollectionIntermediateType will only have one of its values set at most
type endTimeCollectionIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *time.Time type for endTime property
	dateTime *time.Time
	// Stores possible *url.URL type for endTime property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *endTimeCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *endTimeCollectionIntermediateType) Serialize() (i interface{}, err error) {
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

// generatorCollectionIntermediateType will only have one of its values set at most
type generatorCollectionIntermediateType struct {
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
func (t *generatorCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *generatorCollectionIntermediateType) Serialize() (i interface{}, err error) {
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

// iconCollectionIntermediateType will only have one of its values set at most
type iconCollectionIntermediateType struct {
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
func (t *iconCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *iconCollectionIntermediateType) Serialize() (i interface{}, err error) {
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

// imageCollectionIntermediateType will only have one of its values set at most
type imageCollectionIntermediateType struct {
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
func (t *imageCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *imageCollectionIntermediateType) Serialize() (i interface{}, err error) {
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

// inReplyToCollectionIntermediateType will only have one of its values set at most
type inReplyToCollectionIntermediateType struct {
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
func (t *inReplyToCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *inReplyToCollectionIntermediateType) Serialize() (i interface{}, err error) {
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

// locationCollectionIntermediateType will only have one of its values set at most
type locationCollectionIntermediateType struct {
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
func (t *locationCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *locationCollectionIntermediateType) Serialize() (i interface{}, err error) {
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

// previewCollectionIntermediateType will only have one of its values set at most
type previewCollectionIntermediateType struct {
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
func (t *previewCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *previewCollectionIntermediateType) Serialize() (i interface{}, err error) {
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

// publishedCollectionIntermediateType will only have one of its values set at most
type publishedCollectionIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *time.Time type for published property
	dateTime *time.Time
	// Stores possible *url.URL type for published property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *publishedCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *publishedCollectionIntermediateType) Serialize() (i interface{}, err error) {
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

// repliesCollectionIntermediateType will only have one of its values set at most
type repliesCollectionIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible CollectionType type for replies property
	Collection CollectionType
	// Stores possible *url.URL type for replies property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *repliesCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *repliesCollectionIntermediateType) Serialize() (i interface{}, err error) {
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

// startTimeCollectionIntermediateType will only have one of its values set at most
type startTimeCollectionIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *time.Time type for startTime property
	dateTime *time.Time
	// Stores possible *url.URL type for startTime property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *startTimeCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *startTimeCollectionIntermediateType) Serialize() (i interface{}, err error) {
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

// summaryCollectionIntermediateType will only have one of its values set at most
type summaryCollectionIntermediateType struct {
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
func (t *summaryCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *summaryCollectionIntermediateType) Serialize() (i interface{}, err error) {
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

// tagCollectionIntermediateType will only have one of its values set at most
type tagCollectionIntermediateType struct {
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
func (t *tagCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *tagCollectionIntermediateType) Serialize() (i interface{}, err error) {
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

// updatedCollectionIntermediateType will only have one of its values set at most
type updatedCollectionIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *time.Time type for updated property
	dateTime *time.Time
	// Stores possible *url.URL type for updated property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *updatedCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *updatedCollectionIntermediateType) Serialize() (i interface{}, err error) {
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

// urlCollectionIntermediateType will only have one of its values set at most
type urlCollectionIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *url.URL type for url property
	anyURI *url.URL
	// Stores possible LinkType type for url property
	Link LinkType
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *urlCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *urlCollectionIntermediateType) Serialize() (i interface{}, err error) {
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

// toCollectionIntermediateType will only have one of its values set at most
type toCollectionIntermediateType struct {
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
func (t *toCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *toCollectionIntermediateType) Serialize() (i interface{}, err error) {
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

// btoCollectionIntermediateType will only have one of its values set at most
type btoCollectionIntermediateType struct {
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
func (t *btoCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *btoCollectionIntermediateType) Serialize() (i interface{}, err error) {
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

// ccCollectionIntermediateType will only have one of its values set at most
type ccCollectionIntermediateType struct {
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
func (t *ccCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *ccCollectionIntermediateType) Serialize() (i interface{}, err error) {
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

// bccCollectionIntermediateType will only have one of its values set at most
type bccCollectionIntermediateType struct {
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
func (t *bccCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *bccCollectionIntermediateType) Serialize() (i interface{}, err error) {
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

// mediaTypeCollectionIntermediateType will only have one of its values set at most
type mediaTypeCollectionIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *string type for mediaType property
	mimeMediaTypeValue *string
	// Stores possible *url.URL type for mediaType property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *mediaTypeCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *mediaTypeCollectionIntermediateType) Serialize() (i interface{}, err error) {
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

// durationCollectionIntermediateType will only have one of its values set at most
type durationCollectionIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *time.Duration type for duration property
	duration *time.Duration
	// Stores possible *url.URL type for duration property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *durationCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *durationCollectionIntermediateType) Serialize() (i interface{}, err error) {
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

// sourceCollectionIntermediateType will only have one of its values set at most
type sourceCollectionIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for source property
	Object ObjectType
	// Stores possible *url.URL type for source property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *sourceCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *sourceCollectionIntermediateType) Serialize() (i interface{}, err error) {
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

// inboxCollectionIntermediateType will only have one of its values set at most
type inboxCollectionIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible OrderedCollectionType type for inbox property
	OrderedCollection OrderedCollectionType
	// Stores possible *url.URL type for inbox property
	anyURI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *inboxCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *inboxCollectionIntermediateType) Serialize() (i interface{}, err error) {
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

// outboxCollectionIntermediateType will only have one of its values set at most
type outboxCollectionIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible OrderedCollectionType type for outbox property
	OrderedCollection OrderedCollectionType
	// Stores possible *url.URL type for outbox property
	anyURI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *outboxCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *outboxCollectionIntermediateType) Serialize() (i interface{}, err error) {
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

// followingCollectionIntermediateType will only have one of its values set at most
type followingCollectionIntermediateType struct {
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
func (t *followingCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *followingCollectionIntermediateType) Serialize() (i interface{}, err error) {
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

// followersCollectionIntermediateType will only have one of its values set at most
type followersCollectionIntermediateType struct {
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
func (t *followersCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *followersCollectionIntermediateType) Serialize() (i interface{}, err error) {
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

// likedCollectionIntermediateType will only have one of its values set at most
type likedCollectionIntermediateType struct {
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
func (t *likedCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *likedCollectionIntermediateType) Serialize() (i interface{}, err error) {
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

// likesCollectionIntermediateType will only have one of its values set at most
type likesCollectionIntermediateType struct {
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
func (t *likesCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *likesCollectionIntermediateType) Serialize() (i interface{}, err error) {
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

// preferredUsernameCollectionIntermediateType will only have one of its values set at most
type preferredUsernameCollectionIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *string type for preferredUsername property
	stringName *string
	// Stores possible *url.URL type for preferredUsername property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *preferredUsernameCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *preferredUsernameCollectionIntermediateType) Serialize() (i interface{}, err error) {
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

// endpointsCollectionIntermediateType will only have one of its values set at most
type endpointsCollectionIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for endpoints property
	Object ObjectType
	// Stores possible *url.URL type for endpoints property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *endpointsCollectionIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *endpointsCollectionIntermediateType) Serialize() (i interface{}, err error) {
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

// deserializetotalItemsCollectionIntermediateType will accept a map to create a totalItemsCollectionIntermediateType
func deserializeTotalItemsCollectionIntermediateType(in interface{}) (t *totalItemsCollectionIntermediateType, err error) {
	tmp := &totalItemsCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice totalItemsCollectionIntermediateType will accept a slice to create a slice of totalItemsCollectionIntermediateType
func deserializeSliceTotalItemsCollectionIntermediateType(in []interface{}) (t []*totalItemsCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &totalItemsCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializetotalItemsCollectionIntermediateType will accept a totalItemsCollectionIntermediateType to create a map
func serializeTotalItemsCollectionIntermediateType(t *totalItemsCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicetotalItemsCollectionIntermediateType will accept a slice of totalItemsCollectionIntermediateType to create a slice result
func serializeSliceTotalItemsCollectionIntermediateType(s []*totalItemsCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializecurrentCollectionIntermediateType will accept a map to create a currentCollectionIntermediateType
func deserializeCurrentCollectionIntermediateType(in interface{}) (t *currentCollectionIntermediateType, err error) {
	tmp := &currentCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice currentCollectionIntermediateType will accept a slice to create a slice of currentCollectionIntermediateType
func deserializeSliceCurrentCollectionIntermediateType(in []interface{}) (t []*currentCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &currentCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializecurrentCollectionIntermediateType will accept a currentCollectionIntermediateType to create a map
func serializeCurrentCollectionIntermediateType(t *currentCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicecurrentCollectionIntermediateType will accept a slice of currentCollectionIntermediateType to create a slice result
func serializeSliceCurrentCollectionIntermediateType(s []*currentCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializefirstCollectionIntermediateType will accept a map to create a firstCollectionIntermediateType
func deserializeFirstCollectionIntermediateType(in interface{}) (t *firstCollectionIntermediateType, err error) {
	tmp := &firstCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice firstCollectionIntermediateType will accept a slice to create a slice of firstCollectionIntermediateType
func deserializeSliceFirstCollectionIntermediateType(in []interface{}) (t []*firstCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &firstCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializefirstCollectionIntermediateType will accept a firstCollectionIntermediateType to create a map
func serializeFirstCollectionIntermediateType(t *firstCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicefirstCollectionIntermediateType will accept a slice of firstCollectionIntermediateType to create a slice result
func serializeSliceFirstCollectionIntermediateType(s []*firstCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializelastCollectionIntermediateType will accept a map to create a lastCollectionIntermediateType
func deserializeLastCollectionIntermediateType(in interface{}) (t *lastCollectionIntermediateType, err error) {
	tmp := &lastCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice lastCollectionIntermediateType will accept a slice to create a slice of lastCollectionIntermediateType
func deserializeSliceLastCollectionIntermediateType(in []interface{}) (t []*lastCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &lastCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializelastCollectionIntermediateType will accept a lastCollectionIntermediateType to create a map
func serializeLastCollectionIntermediateType(t *lastCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicelastCollectionIntermediateType will accept a slice of lastCollectionIntermediateType to create a slice result
func serializeSliceLastCollectionIntermediateType(s []*lastCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeitemsCollectionIntermediateType will accept a map to create a itemsCollectionIntermediateType
func deserializeItemsCollectionIntermediateType(in interface{}) (t *itemsCollectionIntermediateType, err error) {
	tmp := &itemsCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice itemsCollectionIntermediateType will accept a slice to create a slice of itemsCollectionIntermediateType
func deserializeSliceItemsCollectionIntermediateType(in []interface{}) (t []*itemsCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &itemsCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeitemsCollectionIntermediateType will accept a itemsCollectionIntermediateType to create a map
func serializeItemsCollectionIntermediateType(t *itemsCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceitemsCollectionIntermediateType will accept a slice of itemsCollectionIntermediateType to create a slice result
func serializeSliceItemsCollectionIntermediateType(s []*itemsCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializealtitudeCollectionIntermediateType will accept a map to create a altitudeCollectionIntermediateType
func deserializeAltitudeCollectionIntermediateType(in interface{}) (t *altitudeCollectionIntermediateType, err error) {
	tmp := &altitudeCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice altitudeCollectionIntermediateType will accept a slice to create a slice of altitudeCollectionIntermediateType
func deserializeSliceAltitudeCollectionIntermediateType(in []interface{}) (t []*altitudeCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &altitudeCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializealtitudeCollectionIntermediateType will accept a altitudeCollectionIntermediateType to create a map
func serializeAltitudeCollectionIntermediateType(t *altitudeCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicealtitudeCollectionIntermediateType will accept a slice of altitudeCollectionIntermediateType to create a slice result
func serializeSliceAltitudeCollectionIntermediateType(s []*altitudeCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeattachmentCollectionIntermediateType will accept a map to create a attachmentCollectionIntermediateType
func deserializeAttachmentCollectionIntermediateType(in interface{}) (t *attachmentCollectionIntermediateType, err error) {
	tmp := &attachmentCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice attachmentCollectionIntermediateType will accept a slice to create a slice of attachmentCollectionIntermediateType
func deserializeSliceAttachmentCollectionIntermediateType(in []interface{}) (t []*attachmentCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &attachmentCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeattachmentCollectionIntermediateType will accept a attachmentCollectionIntermediateType to create a map
func serializeAttachmentCollectionIntermediateType(t *attachmentCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceattachmentCollectionIntermediateType will accept a slice of attachmentCollectionIntermediateType to create a slice result
func serializeSliceAttachmentCollectionIntermediateType(s []*attachmentCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeattributedToCollectionIntermediateType will accept a map to create a attributedToCollectionIntermediateType
func deserializeAttributedToCollectionIntermediateType(in interface{}) (t *attributedToCollectionIntermediateType, err error) {
	tmp := &attributedToCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice attributedToCollectionIntermediateType will accept a slice to create a slice of attributedToCollectionIntermediateType
func deserializeSliceAttributedToCollectionIntermediateType(in []interface{}) (t []*attributedToCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &attributedToCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeattributedToCollectionIntermediateType will accept a attributedToCollectionIntermediateType to create a map
func serializeAttributedToCollectionIntermediateType(t *attributedToCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceattributedToCollectionIntermediateType will accept a slice of attributedToCollectionIntermediateType to create a slice result
func serializeSliceAttributedToCollectionIntermediateType(s []*attributedToCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeaudienceCollectionIntermediateType will accept a map to create a audienceCollectionIntermediateType
func deserializeAudienceCollectionIntermediateType(in interface{}) (t *audienceCollectionIntermediateType, err error) {
	tmp := &audienceCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice audienceCollectionIntermediateType will accept a slice to create a slice of audienceCollectionIntermediateType
func deserializeSliceAudienceCollectionIntermediateType(in []interface{}) (t []*audienceCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &audienceCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeaudienceCollectionIntermediateType will accept a audienceCollectionIntermediateType to create a map
func serializeAudienceCollectionIntermediateType(t *audienceCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceaudienceCollectionIntermediateType will accept a slice of audienceCollectionIntermediateType to create a slice result
func serializeSliceAudienceCollectionIntermediateType(s []*audienceCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializecontentCollectionIntermediateType will accept a map to create a contentCollectionIntermediateType
func deserializeContentCollectionIntermediateType(in interface{}) (t *contentCollectionIntermediateType, err error) {
	tmp := &contentCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice contentCollectionIntermediateType will accept a slice to create a slice of contentCollectionIntermediateType
func deserializeSliceContentCollectionIntermediateType(in []interface{}) (t []*contentCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &contentCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializecontentCollectionIntermediateType will accept a contentCollectionIntermediateType to create a map
func serializeContentCollectionIntermediateType(t *contentCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicecontentCollectionIntermediateType will accept a slice of contentCollectionIntermediateType to create a slice result
func serializeSliceContentCollectionIntermediateType(s []*contentCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializecontextCollectionIntermediateType will accept a map to create a contextCollectionIntermediateType
func deserializeContextCollectionIntermediateType(in interface{}) (t *contextCollectionIntermediateType, err error) {
	tmp := &contextCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice contextCollectionIntermediateType will accept a slice to create a slice of contextCollectionIntermediateType
func deserializeSliceContextCollectionIntermediateType(in []interface{}) (t []*contextCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &contextCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializecontextCollectionIntermediateType will accept a contextCollectionIntermediateType to create a map
func serializeContextCollectionIntermediateType(t *contextCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicecontextCollectionIntermediateType will accept a slice of contextCollectionIntermediateType to create a slice result
func serializeSliceContextCollectionIntermediateType(s []*contextCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializenameCollectionIntermediateType will accept a map to create a nameCollectionIntermediateType
func deserializeNameCollectionIntermediateType(in interface{}) (t *nameCollectionIntermediateType, err error) {
	tmp := &nameCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice nameCollectionIntermediateType will accept a slice to create a slice of nameCollectionIntermediateType
func deserializeSliceNameCollectionIntermediateType(in []interface{}) (t []*nameCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &nameCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializenameCollectionIntermediateType will accept a nameCollectionIntermediateType to create a map
func serializeNameCollectionIntermediateType(t *nameCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicenameCollectionIntermediateType will accept a slice of nameCollectionIntermediateType to create a slice result
func serializeSliceNameCollectionIntermediateType(s []*nameCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeendTimeCollectionIntermediateType will accept a map to create a endTimeCollectionIntermediateType
func deserializeEndTimeCollectionIntermediateType(in interface{}) (t *endTimeCollectionIntermediateType, err error) {
	tmp := &endTimeCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice endTimeCollectionIntermediateType will accept a slice to create a slice of endTimeCollectionIntermediateType
func deserializeSliceEndTimeCollectionIntermediateType(in []interface{}) (t []*endTimeCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &endTimeCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeendTimeCollectionIntermediateType will accept a endTimeCollectionIntermediateType to create a map
func serializeEndTimeCollectionIntermediateType(t *endTimeCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceendTimeCollectionIntermediateType will accept a slice of endTimeCollectionIntermediateType to create a slice result
func serializeSliceEndTimeCollectionIntermediateType(s []*endTimeCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializegeneratorCollectionIntermediateType will accept a map to create a generatorCollectionIntermediateType
func deserializeGeneratorCollectionIntermediateType(in interface{}) (t *generatorCollectionIntermediateType, err error) {
	tmp := &generatorCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice generatorCollectionIntermediateType will accept a slice to create a slice of generatorCollectionIntermediateType
func deserializeSliceGeneratorCollectionIntermediateType(in []interface{}) (t []*generatorCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &generatorCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializegeneratorCollectionIntermediateType will accept a generatorCollectionIntermediateType to create a map
func serializeGeneratorCollectionIntermediateType(t *generatorCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicegeneratorCollectionIntermediateType will accept a slice of generatorCollectionIntermediateType to create a slice result
func serializeSliceGeneratorCollectionIntermediateType(s []*generatorCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeiconCollectionIntermediateType will accept a map to create a iconCollectionIntermediateType
func deserializeIconCollectionIntermediateType(in interface{}) (t *iconCollectionIntermediateType, err error) {
	tmp := &iconCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice iconCollectionIntermediateType will accept a slice to create a slice of iconCollectionIntermediateType
func deserializeSliceIconCollectionIntermediateType(in []interface{}) (t []*iconCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &iconCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeiconCollectionIntermediateType will accept a iconCollectionIntermediateType to create a map
func serializeIconCollectionIntermediateType(t *iconCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceiconCollectionIntermediateType will accept a slice of iconCollectionIntermediateType to create a slice result
func serializeSliceIconCollectionIntermediateType(s []*iconCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeimageCollectionIntermediateType will accept a map to create a imageCollectionIntermediateType
func deserializeImageCollectionIntermediateType(in interface{}) (t *imageCollectionIntermediateType, err error) {
	tmp := &imageCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice imageCollectionIntermediateType will accept a slice to create a slice of imageCollectionIntermediateType
func deserializeSliceImageCollectionIntermediateType(in []interface{}) (t []*imageCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &imageCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeimageCollectionIntermediateType will accept a imageCollectionIntermediateType to create a map
func serializeImageCollectionIntermediateType(t *imageCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceimageCollectionIntermediateType will accept a slice of imageCollectionIntermediateType to create a slice result
func serializeSliceImageCollectionIntermediateType(s []*imageCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeinReplyToCollectionIntermediateType will accept a map to create a inReplyToCollectionIntermediateType
func deserializeInReplyToCollectionIntermediateType(in interface{}) (t *inReplyToCollectionIntermediateType, err error) {
	tmp := &inReplyToCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice inReplyToCollectionIntermediateType will accept a slice to create a slice of inReplyToCollectionIntermediateType
func deserializeSliceInReplyToCollectionIntermediateType(in []interface{}) (t []*inReplyToCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &inReplyToCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeinReplyToCollectionIntermediateType will accept a inReplyToCollectionIntermediateType to create a map
func serializeInReplyToCollectionIntermediateType(t *inReplyToCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceinReplyToCollectionIntermediateType will accept a slice of inReplyToCollectionIntermediateType to create a slice result
func serializeSliceInReplyToCollectionIntermediateType(s []*inReplyToCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializelocationCollectionIntermediateType will accept a map to create a locationCollectionIntermediateType
func deserializeLocationCollectionIntermediateType(in interface{}) (t *locationCollectionIntermediateType, err error) {
	tmp := &locationCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice locationCollectionIntermediateType will accept a slice to create a slice of locationCollectionIntermediateType
func deserializeSliceLocationCollectionIntermediateType(in []interface{}) (t []*locationCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &locationCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializelocationCollectionIntermediateType will accept a locationCollectionIntermediateType to create a map
func serializeLocationCollectionIntermediateType(t *locationCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicelocationCollectionIntermediateType will accept a slice of locationCollectionIntermediateType to create a slice result
func serializeSliceLocationCollectionIntermediateType(s []*locationCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializepreviewCollectionIntermediateType will accept a map to create a previewCollectionIntermediateType
func deserializePreviewCollectionIntermediateType(in interface{}) (t *previewCollectionIntermediateType, err error) {
	tmp := &previewCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice previewCollectionIntermediateType will accept a slice to create a slice of previewCollectionIntermediateType
func deserializeSlicePreviewCollectionIntermediateType(in []interface{}) (t []*previewCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &previewCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializepreviewCollectionIntermediateType will accept a previewCollectionIntermediateType to create a map
func serializePreviewCollectionIntermediateType(t *previewCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicepreviewCollectionIntermediateType will accept a slice of previewCollectionIntermediateType to create a slice result
func serializeSlicePreviewCollectionIntermediateType(s []*previewCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializepublishedCollectionIntermediateType will accept a map to create a publishedCollectionIntermediateType
func deserializePublishedCollectionIntermediateType(in interface{}) (t *publishedCollectionIntermediateType, err error) {
	tmp := &publishedCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice publishedCollectionIntermediateType will accept a slice to create a slice of publishedCollectionIntermediateType
func deserializeSlicePublishedCollectionIntermediateType(in []interface{}) (t []*publishedCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &publishedCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializepublishedCollectionIntermediateType will accept a publishedCollectionIntermediateType to create a map
func serializePublishedCollectionIntermediateType(t *publishedCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicepublishedCollectionIntermediateType will accept a slice of publishedCollectionIntermediateType to create a slice result
func serializeSlicePublishedCollectionIntermediateType(s []*publishedCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializerepliesCollectionIntermediateType will accept a map to create a repliesCollectionIntermediateType
func deserializeRepliesCollectionIntermediateType(in interface{}) (t *repliesCollectionIntermediateType, err error) {
	tmp := &repliesCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice repliesCollectionIntermediateType will accept a slice to create a slice of repliesCollectionIntermediateType
func deserializeSliceRepliesCollectionIntermediateType(in []interface{}) (t []*repliesCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &repliesCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializerepliesCollectionIntermediateType will accept a repliesCollectionIntermediateType to create a map
func serializeRepliesCollectionIntermediateType(t *repliesCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicerepliesCollectionIntermediateType will accept a slice of repliesCollectionIntermediateType to create a slice result
func serializeSliceRepliesCollectionIntermediateType(s []*repliesCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializestartTimeCollectionIntermediateType will accept a map to create a startTimeCollectionIntermediateType
func deserializeStartTimeCollectionIntermediateType(in interface{}) (t *startTimeCollectionIntermediateType, err error) {
	tmp := &startTimeCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice startTimeCollectionIntermediateType will accept a slice to create a slice of startTimeCollectionIntermediateType
func deserializeSliceStartTimeCollectionIntermediateType(in []interface{}) (t []*startTimeCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &startTimeCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializestartTimeCollectionIntermediateType will accept a startTimeCollectionIntermediateType to create a map
func serializeStartTimeCollectionIntermediateType(t *startTimeCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicestartTimeCollectionIntermediateType will accept a slice of startTimeCollectionIntermediateType to create a slice result
func serializeSliceStartTimeCollectionIntermediateType(s []*startTimeCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializesummaryCollectionIntermediateType will accept a map to create a summaryCollectionIntermediateType
func deserializeSummaryCollectionIntermediateType(in interface{}) (t *summaryCollectionIntermediateType, err error) {
	tmp := &summaryCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice summaryCollectionIntermediateType will accept a slice to create a slice of summaryCollectionIntermediateType
func deserializeSliceSummaryCollectionIntermediateType(in []interface{}) (t []*summaryCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &summaryCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializesummaryCollectionIntermediateType will accept a summaryCollectionIntermediateType to create a map
func serializeSummaryCollectionIntermediateType(t *summaryCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicesummaryCollectionIntermediateType will accept a slice of summaryCollectionIntermediateType to create a slice result
func serializeSliceSummaryCollectionIntermediateType(s []*summaryCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializetagCollectionIntermediateType will accept a map to create a tagCollectionIntermediateType
func deserializeTagCollectionIntermediateType(in interface{}) (t *tagCollectionIntermediateType, err error) {
	tmp := &tagCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice tagCollectionIntermediateType will accept a slice to create a slice of tagCollectionIntermediateType
func deserializeSliceTagCollectionIntermediateType(in []interface{}) (t []*tagCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &tagCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializetagCollectionIntermediateType will accept a tagCollectionIntermediateType to create a map
func serializeTagCollectionIntermediateType(t *tagCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicetagCollectionIntermediateType will accept a slice of tagCollectionIntermediateType to create a slice result
func serializeSliceTagCollectionIntermediateType(s []*tagCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeupdatedCollectionIntermediateType will accept a map to create a updatedCollectionIntermediateType
func deserializeUpdatedCollectionIntermediateType(in interface{}) (t *updatedCollectionIntermediateType, err error) {
	tmp := &updatedCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice updatedCollectionIntermediateType will accept a slice to create a slice of updatedCollectionIntermediateType
func deserializeSliceUpdatedCollectionIntermediateType(in []interface{}) (t []*updatedCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &updatedCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeupdatedCollectionIntermediateType will accept a updatedCollectionIntermediateType to create a map
func serializeUpdatedCollectionIntermediateType(t *updatedCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceupdatedCollectionIntermediateType will accept a slice of updatedCollectionIntermediateType to create a slice result
func serializeSliceUpdatedCollectionIntermediateType(s []*updatedCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeurlCollectionIntermediateType will accept a map to create a urlCollectionIntermediateType
func deserializeUrlCollectionIntermediateType(in interface{}) (t *urlCollectionIntermediateType, err error) {
	tmp := &urlCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice urlCollectionIntermediateType will accept a slice to create a slice of urlCollectionIntermediateType
func deserializeSliceUrlCollectionIntermediateType(in []interface{}) (t []*urlCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &urlCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeurlCollectionIntermediateType will accept a urlCollectionIntermediateType to create a map
func serializeUrlCollectionIntermediateType(t *urlCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceurlCollectionIntermediateType will accept a slice of urlCollectionIntermediateType to create a slice result
func serializeSliceUrlCollectionIntermediateType(s []*urlCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializetoCollectionIntermediateType will accept a map to create a toCollectionIntermediateType
func deserializeToCollectionIntermediateType(in interface{}) (t *toCollectionIntermediateType, err error) {
	tmp := &toCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice toCollectionIntermediateType will accept a slice to create a slice of toCollectionIntermediateType
func deserializeSliceToCollectionIntermediateType(in []interface{}) (t []*toCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &toCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializetoCollectionIntermediateType will accept a toCollectionIntermediateType to create a map
func serializeToCollectionIntermediateType(t *toCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicetoCollectionIntermediateType will accept a slice of toCollectionIntermediateType to create a slice result
func serializeSliceToCollectionIntermediateType(s []*toCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializebtoCollectionIntermediateType will accept a map to create a btoCollectionIntermediateType
func deserializeBtoCollectionIntermediateType(in interface{}) (t *btoCollectionIntermediateType, err error) {
	tmp := &btoCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice btoCollectionIntermediateType will accept a slice to create a slice of btoCollectionIntermediateType
func deserializeSliceBtoCollectionIntermediateType(in []interface{}) (t []*btoCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &btoCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializebtoCollectionIntermediateType will accept a btoCollectionIntermediateType to create a map
func serializeBtoCollectionIntermediateType(t *btoCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicebtoCollectionIntermediateType will accept a slice of btoCollectionIntermediateType to create a slice result
func serializeSliceBtoCollectionIntermediateType(s []*btoCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeccCollectionIntermediateType will accept a map to create a ccCollectionIntermediateType
func deserializeCcCollectionIntermediateType(in interface{}) (t *ccCollectionIntermediateType, err error) {
	tmp := &ccCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice ccCollectionIntermediateType will accept a slice to create a slice of ccCollectionIntermediateType
func deserializeSliceCcCollectionIntermediateType(in []interface{}) (t []*ccCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &ccCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeccCollectionIntermediateType will accept a ccCollectionIntermediateType to create a map
func serializeCcCollectionIntermediateType(t *ccCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceccCollectionIntermediateType will accept a slice of ccCollectionIntermediateType to create a slice result
func serializeSliceCcCollectionIntermediateType(s []*ccCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializebccCollectionIntermediateType will accept a map to create a bccCollectionIntermediateType
func deserializeBccCollectionIntermediateType(in interface{}) (t *bccCollectionIntermediateType, err error) {
	tmp := &bccCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice bccCollectionIntermediateType will accept a slice to create a slice of bccCollectionIntermediateType
func deserializeSliceBccCollectionIntermediateType(in []interface{}) (t []*bccCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &bccCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializebccCollectionIntermediateType will accept a bccCollectionIntermediateType to create a map
func serializeBccCollectionIntermediateType(t *bccCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicebccCollectionIntermediateType will accept a slice of bccCollectionIntermediateType to create a slice result
func serializeSliceBccCollectionIntermediateType(s []*bccCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializemediaTypeCollectionIntermediateType will accept a map to create a mediaTypeCollectionIntermediateType
func deserializeMediaTypeCollectionIntermediateType(in interface{}) (t *mediaTypeCollectionIntermediateType, err error) {
	tmp := &mediaTypeCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice mediaTypeCollectionIntermediateType will accept a slice to create a slice of mediaTypeCollectionIntermediateType
func deserializeSliceMediaTypeCollectionIntermediateType(in []interface{}) (t []*mediaTypeCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &mediaTypeCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializemediaTypeCollectionIntermediateType will accept a mediaTypeCollectionIntermediateType to create a map
func serializeMediaTypeCollectionIntermediateType(t *mediaTypeCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicemediaTypeCollectionIntermediateType will accept a slice of mediaTypeCollectionIntermediateType to create a slice result
func serializeSliceMediaTypeCollectionIntermediateType(s []*mediaTypeCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializedurationCollectionIntermediateType will accept a map to create a durationCollectionIntermediateType
func deserializeDurationCollectionIntermediateType(in interface{}) (t *durationCollectionIntermediateType, err error) {
	tmp := &durationCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice durationCollectionIntermediateType will accept a slice to create a slice of durationCollectionIntermediateType
func deserializeSliceDurationCollectionIntermediateType(in []interface{}) (t []*durationCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &durationCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializedurationCollectionIntermediateType will accept a durationCollectionIntermediateType to create a map
func serializeDurationCollectionIntermediateType(t *durationCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicedurationCollectionIntermediateType will accept a slice of durationCollectionIntermediateType to create a slice result
func serializeSliceDurationCollectionIntermediateType(s []*durationCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializesourceCollectionIntermediateType will accept a map to create a sourceCollectionIntermediateType
func deserializeSourceCollectionIntermediateType(in interface{}) (t *sourceCollectionIntermediateType, err error) {
	tmp := &sourceCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice sourceCollectionIntermediateType will accept a slice to create a slice of sourceCollectionIntermediateType
func deserializeSliceSourceCollectionIntermediateType(in []interface{}) (t []*sourceCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &sourceCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializesourceCollectionIntermediateType will accept a sourceCollectionIntermediateType to create a map
func serializeSourceCollectionIntermediateType(t *sourceCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicesourceCollectionIntermediateType will accept a slice of sourceCollectionIntermediateType to create a slice result
func serializeSliceSourceCollectionIntermediateType(s []*sourceCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeinboxCollectionIntermediateType will accept a map to create a inboxCollectionIntermediateType
func deserializeInboxCollectionIntermediateType(in interface{}) (t *inboxCollectionIntermediateType, err error) {
	tmp := &inboxCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice inboxCollectionIntermediateType will accept a slice to create a slice of inboxCollectionIntermediateType
func deserializeSliceInboxCollectionIntermediateType(in []interface{}) (t []*inboxCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &inboxCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeinboxCollectionIntermediateType will accept a inboxCollectionIntermediateType to create a map
func serializeInboxCollectionIntermediateType(t *inboxCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceinboxCollectionIntermediateType will accept a slice of inboxCollectionIntermediateType to create a slice result
func serializeSliceInboxCollectionIntermediateType(s []*inboxCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeoutboxCollectionIntermediateType will accept a map to create a outboxCollectionIntermediateType
func deserializeOutboxCollectionIntermediateType(in interface{}) (t *outboxCollectionIntermediateType, err error) {
	tmp := &outboxCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice outboxCollectionIntermediateType will accept a slice to create a slice of outboxCollectionIntermediateType
func deserializeSliceOutboxCollectionIntermediateType(in []interface{}) (t []*outboxCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &outboxCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeoutboxCollectionIntermediateType will accept a outboxCollectionIntermediateType to create a map
func serializeOutboxCollectionIntermediateType(t *outboxCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceoutboxCollectionIntermediateType will accept a slice of outboxCollectionIntermediateType to create a slice result
func serializeSliceOutboxCollectionIntermediateType(s []*outboxCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializefollowingCollectionIntermediateType will accept a map to create a followingCollectionIntermediateType
func deserializeFollowingCollectionIntermediateType(in interface{}) (t *followingCollectionIntermediateType, err error) {
	tmp := &followingCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice followingCollectionIntermediateType will accept a slice to create a slice of followingCollectionIntermediateType
func deserializeSliceFollowingCollectionIntermediateType(in []interface{}) (t []*followingCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &followingCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializefollowingCollectionIntermediateType will accept a followingCollectionIntermediateType to create a map
func serializeFollowingCollectionIntermediateType(t *followingCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicefollowingCollectionIntermediateType will accept a slice of followingCollectionIntermediateType to create a slice result
func serializeSliceFollowingCollectionIntermediateType(s []*followingCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializefollowersCollectionIntermediateType will accept a map to create a followersCollectionIntermediateType
func deserializeFollowersCollectionIntermediateType(in interface{}) (t *followersCollectionIntermediateType, err error) {
	tmp := &followersCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice followersCollectionIntermediateType will accept a slice to create a slice of followersCollectionIntermediateType
func deserializeSliceFollowersCollectionIntermediateType(in []interface{}) (t []*followersCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &followersCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializefollowersCollectionIntermediateType will accept a followersCollectionIntermediateType to create a map
func serializeFollowersCollectionIntermediateType(t *followersCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicefollowersCollectionIntermediateType will accept a slice of followersCollectionIntermediateType to create a slice result
func serializeSliceFollowersCollectionIntermediateType(s []*followersCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializelikedCollectionIntermediateType will accept a map to create a likedCollectionIntermediateType
func deserializeLikedCollectionIntermediateType(in interface{}) (t *likedCollectionIntermediateType, err error) {
	tmp := &likedCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice likedCollectionIntermediateType will accept a slice to create a slice of likedCollectionIntermediateType
func deserializeSliceLikedCollectionIntermediateType(in []interface{}) (t []*likedCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &likedCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializelikedCollectionIntermediateType will accept a likedCollectionIntermediateType to create a map
func serializeLikedCollectionIntermediateType(t *likedCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicelikedCollectionIntermediateType will accept a slice of likedCollectionIntermediateType to create a slice result
func serializeSliceLikedCollectionIntermediateType(s []*likedCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializelikesCollectionIntermediateType will accept a map to create a likesCollectionIntermediateType
func deserializeLikesCollectionIntermediateType(in interface{}) (t *likesCollectionIntermediateType, err error) {
	tmp := &likesCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice likesCollectionIntermediateType will accept a slice to create a slice of likesCollectionIntermediateType
func deserializeSliceLikesCollectionIntermediateType(in []interface{}) (t []*likesCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &likesCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializelikesCollectionIntermediateType will accept a likesCollectionIntermediateType to create a map
func serializeLikesCollectionIntermediateType(t *likesCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicelikesCollectionIntermediateType will accept a slice of likesCollectionIntermediateType to create a slice result
func serializeSliceLikesCollectionIntermediateType(s []*likesCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializepreferredUsernameCollectionIntermediateType will accept a map to create a preferredUsernameCollectionIntermediateType
func deserializePreferredUsernameCollectionIntermediateType(in interface{}) (t *preferredUsernameCollectionIntermediateType, err error) {
	tmp := &preferredUsernameCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice preferredUsernameCollectionIntermediateType will accept a slice to create a slice of preferredUsernameCollectionIntermediateType
func deserializeSlicePreferredUsernameCollectionIntermediateType(in []interface{}) (t []*preferredUsernameCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &preferredUsernameCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializepreferredUsernameCollectionIntermediateType will accept a preferredUsernameCollectionIntermediateType to create a map
func serializePreferredUsernameCollectionIntermediateType(t *preferredUsernameCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicepreferredUsernameCollectionIntermediateType will accept a slice of preferredUsernameCollectionIntermediateType to create a slice result
func serializeSlicePreferredUsernameCollectionIntermediateType(s []*preferredUsernameCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeendpointsCollectionIntermediateType will accept a map to create a endpointsCollectionIntermediateType
func deserializeEndpointsCollectionIntermediateType(in interface{}) (t *endpointsCollectionIntermediateType, err error) {
	tmp := &endpointsCollectionIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice endpointsCollectionIntermediateType will accept a slice to create a slice of endpointsCollectionIntermediateType
func deserializeSliceEndpointsCollectionIntermediateType(in []interface{}) (t []*endpointsCollectionIntermediateType, err error) {
	for _, i := range in {
		tmp := &endpointsCollectionIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeendpointsCollectionIntermediateType will accept a endpointsCollectionIntermediateType to create a map
func serializeEndpointsCollectionIntermediateType(t *endpointsCollectionIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceendpointsCollectionIntermediateType will accept a slice of endpointsCollectionIntermediateType to create a slice result
func serializeSliceEndpointsCollectionIntermediateType(s []*endpointsCollectionIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

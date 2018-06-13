//
package vocab

import (
	"fmt"
	"net/url"
	"time"
)

// OrderedCollectionPageType is an interface for accepting types that extend from 'OrderedCollectionPage'.
type OrderedCollectionPageType interface {
	Serializer
	Deserializer
	IsStartIndex() (ok bool)
	GetStartIndex() (v int64)
	SetStartIndex(v int64)
	IsStartIndexIRI() (ok bool)
	GetStartIndexIRI() (v *url.URL)
	SetStartIndexIRI(v *url.URL)
	IsNextOrderedCollectionPage() (ok bool)
	GetNextOrderedCollectionPage() (v OrderedCollectionPageType)
	SetNextOrderedCollectionPage(v OrderedCollectionPageType)
	IsNextLink() (ok bool)
	GetNextLink() (v LinkType)
	SetNextLink(v LinkType)
	IsNextIRI() (ok bool)
	GetNextIRI() (v *url.URL)
	SetNextIRI(v *url.URL)
	IsPrevOrderedCollectionPage() (ok bool)
	GetPrevOrderedCollectionPage() (v OrderedCollectionPageType)
	SetPrevOrderedCollectionPage(v OrderedCollectionPageType)
	IsPrevLink() (ok bool)
	GetPrevLink() (v LinkType)
	SetPrevLink(v LinkType)
	IsPrevIRI() (ok bool)
	GetPrevIRI() (v *url.URL)
	SetPrevIRI(v *url.URL)
	OrderedItemsLen() (l int)
	IsOrderedItemsObject(index int) (ok bool)
	GetOrderedItemsObject(index int) (v ObjectType)
	AppendOrderedItemsObject(v ObjectType)
	PrependOrderedItemsObject(v ObjectType)
	RemoveOrderedItemsObject(index int)
	IsOrderedItemsLink(index int) (ok bool)
	GetOrderedItemsLink(index int) (v LinkType)
	AppendOrderedItemsLink(v LinkType)
	PrependOrderedItemsLink(v LinkType)
	RemoveOrderedItemsLink(index int)
	IsOrderedItemsIRI(index int) (ok bool)
	GetOrderedItemsIRI(index int) (v *url.URL)
	AppendOrderedItemsIRI(v *url.URL)
	PrependOrderedItemsIRI(v *url.URL)
	RemoveOrderedItemsIRI(index int)
	IsCurrentOrderedCollectionPage() (ok bool)
	GetCurrentOrderedCollectionPage() (v OrderedCollectionPageType)
	SetCurrentOrderedCollectionPage(v OrderedCollectionPageType)
	IsCurrentLink() (ok bool)
	GetCurrentLink() (v LinkType)
	SetCurrentLink(v LinkType)
	IsCurrentIRI() (ok bool)
	GetCurrentIRI() (v *url.URL)
	SetCurrentIRI(v *url.URL)
	IsFirstOrderedCollectionPage() (ok bool)
	GetFirstOrderedCollectionPage() (v OrderedCollectionPageType)
	SetFirstOrderedCollectionPage(v OrderedCollectionPageType)
	IsFirstLink() (ok bool)
	GetFirstLink() (v LinkType)
	SetFirstLink(v LinkType)
	IsFirstIRI() (ok bool)
	GetFirstIRI() (v *url.URL)
	SetFirstIRI(v *url.URL)
	IsLastOrderedCollectionPage() (ok bool)
	GetLastOrderedCollectionPage() (v OrderedCollectionPageType)
	SetLastOrderedCollectionPage(v OrderedCollectionPageType)
	IsLastLink() (ok bool)
	GetLastLink() (v LinkType)
	SetLastLink(v LinkType)
	IsLastIRI() (ok bool)
	GetLastIRI() (v *url.URL)
	SetLastIRI(v *url.URL)
	IsTotalItems() (ok bool)
	GetTotalItems() (v int64)
	SetTotalItems(v int64)
	IsTotalItemsIRI() (ok bool)
	GetTotalItemsIRI() (v *url.URL)
	SetTotalItemsIRI(v *url.URL)
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
	IsPartOfLink() (ok bool)
	GetPartOfLink() (v LinkType)
	SetPartOfLink(v LinkType)
	IsPartOfCollection() (ok bool)
	GetPartOfCollection() (v CollectionType)
	SetPartOfCollection(v CollectionType)
	IsPartOfIRI() (ok bool)
	GetPartOfIRI() (v *url.URL)
	SetPartOfIRI(v *url.URL)
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
}

// Used to represent ordered subsets of items from an OrderedCollection. Refer to the Activity Streams 2.0 Core for a complete description of the OrderedCollectionPage object.
type OrderedCollectionPage struct {
	// An unknown value.
	unknown_ map[string]interface{}
	// The functional 'startIndex' value could have multiple types, but only a single value
	startIndex *startIndexOrderedCollectionPageIntermediateType
	// The functional 'next' value could have multiple types, but only a single value
	next *nextOrderedCollectionPageIntermediateType
	// The functional 'prev' value could have multiple types, but only a single value
	prev *prevOrderedCollectionPageIntermediateType
	// The 'orderedItems' value could have multiple types and values
	orderedItems []*orderedItemsOrderedCollectionPageIntermediateType
	// The functional 'current' value could have multiple types, but only a single value
	current *currentOrderedCollectionPageIntermediateType
	// The functional 'first' value could have multiple types, but only a single value
	first *firstOrderedCollectionPageIntermediateType
	// The functional 'last' value could have multiple types, but only a single value
	last *lastOrderedCollectionPageIntermediateType
	// The functional 'totalItems' value could have multiple types, but only a single value
	totalItems *totalItemsOrderedCollectionPageIntermediateType
	// The functional 'altitude' value could have multiple types, but only a single value
	altitude *altitudeOrderedCollectionPageIntermediateType
	// The 'attachment' value could have multiple types and values
	attachment []*attachmentOrderedCollectionPageIntermediateType
	// The 'attributedTo' value could have multiple types and values
	attributedTo []*attributedToOrderedCollectionPageIntermediateType
	// The 'audience' value could have multiple types and values
	audience []*audienceOrderedCollectionPageIntermediateType
	// The 'content' value could have multiple types and values
	content []*contentOrderedCollectionPageIntermediateType
	// The 'contentMap' value holds language-specific values for property 'content'
	contentMap map[string]string
	// The 'context' value could have multiple types and values
	context []*contextOrderedCollectionPageIntermediateType
	// The 'name' value could have multiple types and values
	name []*nameOrderedCollectionPageIntermediateType
	// The 'nameMap' value holds language-specific values for property 'name'
	nameMap map[string]string
	// The functional 'endTime' value could have multiple types, but only a single value
	endTime *endTimeOrderedCollectionPageIntermediateType
	// The 'generator' value could have multiple types and values
	generator []*generatorOrderedCollectionPageIntermediateType
	// The 'icon' value could have multiple types and values
	icon []*iconOrderedCollectionPageIntermediateType
	// The functional 'id' value holds a single type and a single value
	id *url.URL
	// The 'image' value could have multiple types and values
	image []*imageOrderedCollectionPageIntermediateType
	// The 'inReplyTo' value could have multiple types and values
	inReplyTo []*inReplyToOrderedCollectionPageIntermediateType
	// The 'location' value could have multiple types and values
	location []*locationOrderedCollectionPageIntermediateType
	// The 'preview' value could have multiple types and values
	preview []*previewOrderedCollectionPageIntermediateType
	// The functional 'published' value could have multiple types, but only a single value
	published *publishedOrderedCollectionPageIntermediateType
	// The functional 'replies' value could have multiple types, but only a single value
	replies *repliesOrderedCollectionPageIntermediateType
	// The functional 'startTime' value could have multiple types, but only a single value
	startTime *startTimeOrderedCollectionPageIntermediateType
	// The 'summary' value could have multiple types and values
	summary []*summaryOrderedCollectionPageIntermediateType
	// The 'summaryMap' value holds language-specific values for property 'summary'
	summaryMap map[string]string
	// The 'tag' value could have multiple types and values
	tag []*tagOrderedCollectionPageIntermediateType
	// The 'type' value can hold any type and any number of values
	typeName []interface{}
	// The functional 'updated' value could have multiple types, but only a single value
	updated *updatedOrderedCollectionPageIntermediateType
	// The 'url' value could have multiple types and values
	url []*urlOrderedCollectionPageIntermediateType
	// The 'to' value could have multiple types and values
	to []*toOrderedCollectionPageIntermediateType
	// The 'bto' value could have multiple types and values
	bto []*btoOrderedCollectionPageIntermediateType
	// The 'cc' value could have multiple types and values
	cc []*ccOrderedCollectionPageIntermediateType
	// The 'bcc' value could have multiple types and values
	bcc []*bccOrderedCollectionPageIntermediateType
	// The functional 'mediaType' value could have multiple types, but only a single value
	mediaType *mediaTypeOrderedCollectionPageIntermediateType
	// The functional 'duration' value could have multiple types, but only a single value
	duration *durationOrderedCollectionPageIntermediateType
	// The functional 'source' value could have multiple types, but only a single value
	source *sourceOrderedCollectionPageIntermediateType
	// The functional 'inbox' value could have multiple types, but only a single value
	inbox *inboxOrderedCollectionPageIntermediateType
	// The functional 'outbox' value could have multiple types, but only a single value
	outbox *outboxOrderedCollectionPageIntermediateType
	// The functional 'following' value could have multiple types, but only a single value
	following *followingOrderedCollectionPageIntermediateType
	// The functional 'followers' value could have multiple types, but only a single value
	followers *followersOrderedCollectionPageIntermediateType
	// The functional 'liked' value could have multiple types, but only a single value
	liked *likedOrderedCollectionPageIntermediateType
	// The functional 'likes' value could have multiple types, but only a single value
	likes *likesOrderedCollectionPageIntermediateType
	// The 'streams' value holds a single type and any number of values
	streams []*url.URL
	// The functional 'preferredUsername' value could have multiple types, but only a single value
	preferredUsername *preferredUsernameOrderedCollectionPageIntermediateType
	// The 'preferredUsernameMap' value holds language-specific values for property 'preferredUsername'
	preferredUsernameMap map[string]string
	// The functional 'endpoints' value could have multiple types, but only a single value
	endpoints *endpointsOrderedCollectionPageIntermediateType
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
	// The functional 'partOf' value could have multiple types, but only a single value
	partOf *partOfOrderedCollectionPageIntermediateType
}

// IsStartIndex determines whether the call to GetStartIndex is safe
func (t *OrderedCollectionPage) IsStartIndex() (ok bool) {
	return t.startIndex != nil && t.startIndex.nonNegativeInteger != nil

}

// GetStartIndex returns the value safely if IsStartIndex returned true
func (t *OrderedCollectionPage) GetStartIndex() (v int64) {
	return *t.startIndex.nonNegativeInteger

}

// SetStartIndex sets the value of startIndex to be of int64 type
func (t *OrderedCollectionPage) SetStartIndex(v int64) {
	t.startIndex = &startIndexOrderedCollectionPageIntermediateType{nonNegativeInteger: &v}

}

// IsStartIndexIRI determines whether the call to GetStartIndexIRI is safe
func (t *OrderedCollectionPage) IsStartIndexIRI() (ok bool) {
	return t.startIndex != nil && t.startIndex.IRI != nil

}

// GetStartIndexIRI returns the value safely if IsStartIndexIRI returned true
func (t *OrderedCollectionPage) GetStartIndexIRI() (v *url.URL) {
	return t.startIndex.IRI

}

// SetStartIndexIRI sets the value of startIndex to be of *url.URL type
func (t *OrderedCollectionPage) SetStartIndexIRI(v *url.URL) {
	t.startIndex = &startIndexOrderedCollectionPageIntermediateType{IRI: v}

}

// HasUnknownStartIndex determines whether the call to GetUnknownStartIndex is safe
func (t *OrderedCollectionPage) HasUnknownStartIndex() (ok bool) {
	return t.startIndex != nil && t.startIndex.unknown_ != nil

}

// GetUnknownStartIndex returns the unknown value for startIndex
func (t *OrderedCollectionPage) GetUnknownStartIndex() (v interface{}) {
	return t.startIndex.unknown_

}

// SetUnknownStartIndex sets the unknown value of startIndex
func (t *OrderedCollectionPage) SetUnknownStartIndex(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &startIndexOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.startIndex = tmp

}

// IsNextOrderedCollectionPage determines whether the call to GetNextOrderedCollectionPage is safe
func (t *OrderedCollectionPage) IsNextOrderedCollectionPage() (ok bool) {
	return t.next != nil && t.next.OrderedCollectionPage != nil

}

// GetNextOrderedCollectionPage returns the value safely if IsNextOrderedCollectionPage returned true
func (t *OrderedCollectionPage) GetNextOrderedCollectionPage() (v OrderedCollectionPageType) {
	return t.next.OrderedCollectionPage

}

// SetNextOrderedCollectionPage sets the value of next to be of OrderedCollectionPageType type
func (t *OrderedCollectionPage) SetNextOrderedCollectionPage(v OrderedCollectionPageType) {
	t.next = &nextOrderedCollectionPageIntermediateType{OrderedCollectionPage: v}

}

// IsNextLink determines whether the call to GetNextLink is safe
func (t *OrderedCollectionPage) IsNextLink() (ok bool) {
	return t.next != nil && t.next.Link != nil

}

// GetNextLink returns the value safely if IsNextLink returned true
func (t *OrderedCollectionPage) GetNextLink() (v LinkType) {
	return t.next.Link

}

// SetNextLink sets the value of next to be of LinkType type
func (t *OrderedCollectionPage) SetNextLink(v LinkType) {
	t.next = &nextOrderedCollectionPageIntermediateType{Link: v}

}

// IsNextIRI determines whether the call to GetNextIRI is safe
func (t *OrderedCollectionPage) IsNextIRI() (ok bool) {
	return t.next != nil && t.next.IRI != nil

}

// GetNextIRI returns the value safely if IsNextIRI returned true
func (t *OrderedCollectionPage) GetNextIRI() (v *url.URL) {
	return t.next.IRI

}

// SetNextIRI sets the value of next to be of *url.URL type
func (t *OrderedCollectionPage) SetNextIRI(v *url.URL) {
	t.next = &nextOrderedCollectionPageIntermediateType{IRI: v}

}

// HasUnknownNext determines whether the call to GetUnknownNext is safe
func (t *OrderedCollectionPage) HasUnknownNext() (ok bool) {
	return t.next != nil && t.next.unknown_ != nil

}

// GetUnknownNext returns the unknown value for next
func (t *OrderedCollectionPage) GetUnknownNext() (v interface{}) {
	return t.next.unknown_

}

// SetUnknownNext sets the unknown value of next
func (t *OrderedCollectionPage) SetUnknownNext(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &nextOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.next = tmp

}

// IsPrevOrderedCollectionPage determines whether the call to GetPrevOrderedCollectionPage is safe
func (t *OrderedCollectionPage) IsPrevOrderedCollectionPage() (ok bool) {
	return t.prev != nil && t.prev.OrderedCollectionPage != nil

}

// GetPrevOrderedCollectionPage returns the value safely if IsPrevOrderedCollectionPage returned true
func (t *OrderedCollectionPage) GetPrevOrderedCollectionPage() (v OrderedCollectionPageType) {
	return t.prev.OrderedCollectionPage

}

// SetPrevOrderedCollectionPage sets the value of prev to be of OrderedCollectionPageType type
func (t *OrderedCollectionPage) SetPrevOrderedCollectionPage(v OrderedCollectionPageType) {
	t.prev = &prevOrderedCollectionPageIntermediateType{OrderedCollectionPage: v}

}

// IsPrevLink determines whether the call to GetPrevLink is safe
func (t *OrderedCollectionPage) IsPrevLink() (ok bool) {
	return t.prev != nil && t.prev.Link != nil

}

// GetPrevLink returns the value safely if IsPrevLink returned true
func (t *OrderedCollectionPage) GetPrevLink() (v LinkType) {
	return t.prev.Link

}

// SetPrevLink sets the value of prev to be of LinkType type
func (t *OrderedCollectionPage) SetPrevLink(v LinkType) {
	t.prev = &prevOrderedCollectionPageIntermediateType{Link: v}

}

// IsPrevIRI determines whether the call to GetPrevIRI is safe
func (t *OrderedCollectionPage) IsPrevIRI() (ok bool) {
	return t.prev != nil && t.prev.IRI != nil

}

// GetPrevIRI returns the value safely if IsPrevIRI returned true
func (t *OrderedCollectionPage) GetPrevIRI() (v *url.URL) {
	return t.prev.IRI

}

// SetPrevIRI sets the value of prev to be of *url.URL type
func (t *OrderedCollectionPage) SetPrevIRI(v *url.URL) {
	t.prev = &prevOrderedCollectionPageIntermediateType{IRI: v}

}

// HasUnknownPrev determines whether the call to GetUnknownPrev is safe
func (t *OrderedCollectionPage) HasUnknownPrev() (ok bool) {
	return t.prev != nil && t.prev.unknown_ != nil

}

// GetUnknownPrev returns the unknown value for prev
func (t *OrderedCollectionPage) GetUnknownPrev() (v interface{}) {
	return t.prev.unknown_

}

// SetUnknownPrev sets the unknown value of prev
func (t *OrderedCollectionPage) SetUnknownPrev(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &prevOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.prev = tmp

}

// OrderedItemsLen determines the number of elements able to be used for the IsOrderedItemsObject, GetOrderedItemsObject, and RemoveOrderedItemsObject functions
func (t *OrderedCollectionPage) OrderedItemsLen() (l int) {
	return len(t.orderedItems)

}

// IsOrderedItemsObject determines whether the call to GetOrderedItemsObject is safe for the specified index
func (t *OrderedCollectionPage) IsOrderedItemsObject(index int) (ok bool) {
	return t.orderedItems[index].Object != nil

}

// GetOrderedItemsObject returns the value safely if IsOrderedItemsObject returned true for the specified index
func (t *OrderedCollectionPage) GetOrderedItemsObject(index int) (v ObjectType) {
	return t.orderedItems[index].Object

}

// AppendOrderedItemsObject adds to the back of orderedItems a ObjectType type
func (t *OrderedCollectionPage) AppendOrderedItemsObject(v ObjectType) {
	t.orderedItems = append(t.orderedItems, &orderedItemsOrderedCollectionPageIntermediateType{Object: v})

}

// PrependOrderedItemsObject adds to the front of orderedItems a ObjectType type
func (t *OrderedCollectionPage) PrependOrderedItemsObject(v ObjectType) {
	t.orderedItems = append([]*orderedItemsOrderedCollectionPageIntermediateType{&orderedItemsOrderedCollectionPageIntermediateType{Object: v}}, t.orderedItems...)

}

// RemoveOrderedItemsObject deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveOrderedItemsObject(index int) {
	copy(t.orderedItems[index:], t.orderedItems[index+1:])
	t.orderedItems[len(t.orderedItems)-1] = nil
	t.orderedItems = t.orderedItems[:len(t.orderedItems)-1]

}

// IsOrderedItemsLink determines whether the call to GetOrderedItemsLink is safe for the specified index
func (t *OrderedCollectionPage) IsOrderedItemsLink(index int) (ok bool) {
	return t.orderedItems[index].Link != nil

}

// GetOrderedItemsLink returns the value safely if IsOrderedItemsLink returned true for the specified index
func (t *OrderedCollectionPage) GetOrderedItemsLink(index int) (v LinkType) {
	return t.orderedItems[index].Link

}

// AppendOrderedItemsLink adds to the back of orderedItems a LinkType type
func (t *OrderedCollectionPage) AppendOrderedItemsLink(v LinkType) {
	t.orderedItems = append(t.orderedItems, &orderedItemsOrderedCollectionPageIntermediateType{Link: v})

}

// PrependOrderedItemsLink adds to the front of orderedItems a LinkType type
func (t *OrderedCollectionPage) PrependOrderedItemsLink(v LinkType) {
	t.orderedItems = append([]*orderedItemsOrderedCollectionPageIntermediateType{&orderedItemsOrderedCollectionPageIntermediateType{Link: v}}, t.orderedItems...)

}

// RemoveOrderedItemsLink deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveOrderedItemsLink(index int) {
	copy(t.orderedItems[index:], t.orderedItems[index+1:])
	t.orderedItems[len(t.orderedItems)-1] = nil
	t.orderedItems = t.orderedItems[:len(t.orderedItems)-1]

}

// IsOrderedItemsIRI determines whether the call to GetOrderedItemsIRI is safe for the specified index
func (t *OrderedCollectionPage) IsOrderedItemsIRI(index int) (ok bool) {
	return t.orderedItems[index].IRI != nil

}

// GetOrderedItemsIRI returns the value safely if IsOrderedItemsIRI returned true for the specified index
func (t *OrderedCollectionPage) GetOrderedItemsIRI(index int) (v *url.URL) {
	return t.orderedItems[index].IRI

}

// AppendOrderedItemsIRI adds to the back of orderedItems a *url.URL type
func (t *OrderedCollectionPage) AppendOrderedItemsIRI(v *url.URL) {
	t.orderedItems = append(t.orderedItems, &orderedItemsOrderedCollectionPageIntermediateType{IRI: v})

}

// PrependOrderedItemsIRI adds to the front of orderedItems a *url.URL type
func (t *OrderedCollectionPage) PrependOrderedItemsIRI(v *url.URL) {
	t.orderedItems = append([]*orderedItemsOrderedCollectionPageIntermediateType{&orderedItemsOrderedCollectionPageIntermediateType{IRI: v}}, t.orderedItems...)

}

// RemoveOrderedItemsIRI deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveOrderedItemsIRI(index int) {
	copy(t.orderedItems[index:], t.orderedItems[index+1:])
	t.orderedItems[len(t.orderedItems)-1] = nil
	t.orderedItems = t.orderedItems[:len(t.orderedItems)-1]

}

// HasUnknownOrderedItems determines whether the call to GetUnknownOrderedItems is safe
func (t *OrderedCollectionPage) HasUnknownOrderedItems() (ok bool) {
	return t.orderedItems != nil && t.orderedItems[0].unknown_ != nil

}

// GetUnknownOrderedItems returns the unknown value for orderedItems
func (t *OrderedCollectionPage) GetUnknownOrderedItems() (v interface{}) {
	return t.orderedItems[0].unknown_

}

// SetUnknownOrderedItems sets the unknown value of orderedItems
func (t *OrderedCollectionPage) SetUnknownOrderedItems(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &orderedItemsOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.orderedItems = append(t.orderedItems, tmp)

}

// IsCurrentOrderedCollectionPage determines whether the call to GetCurrentOrderedCollectionPage is safe
func (t *OrderedCollectionPage) IsCurrentOrderedCollectionPage() (ok bool) {
	return t.current != nil && t.current.OrderedCollectionPage != nil

}

// GetCurrentOrderedCollectionPage returns the value safely if IsCurrentOrderedCollectionPage returned true
func (t *OrderedCollectionPage) GetCurrentOrderedCollectionPage() (v OrderedCollectionPageType) {
	return t.current.OrderedCollectionPage

}

// SetCurrentOrderedCollectionPage sets the value of current to be of OrderedCollectionPageType type
func (t *OrderedCollectionPage) SetCurrentOrderedCollectionPage(v OrderedCollectionPageType) {
	t.current = &currentOrderedCollectionPageIntermediateType{OrderedCollectionPage: v}

}

// IsCurrentLink determines whether the call to GetCurrentLink is safe
func (t *OrderedCollectionPage) IsCurrentLink() (ok bool) {
	return t.current != nil && t.current.Link != nil

}

// GetCurrentLink returns the value safely if IsCurrentLink returned true
func (t *OrderedCollectionPage) GetCurrentLink() (v LinkType) {
	return t.current.Link

}

// SetCurrentLink sets the value of current to be of LinkType type
func (t *OrderedCollectionPage) SetCurrentLink(v LinkType) {
	t.current = &currentOrderedCollectionPageIntermediateType{Link: v}

}

// IsCurrentIRI determines whether the call to GetCurrentIRI is safe
func (t *OrderedCollectionPage) IsCurrentIRI() (ok bool) {
	return t.current != nil && t.current.IRI != nil

}

// GetCurrentIRI returns the value safely if IsCurrentIRI returned true
func (t *OrderedCollectionPage) GetCurrentIRI() (v *url.URL) {
	return t.current.IRI

}

// SetCurrentIRI sets the value of current to be of *url.URL type
func (t *OrderedCollectionPage) SetCurrentIRI(v *url.URL) {
	t.current = &currentOrderedCollectionPageIntermediateType{IRI: v}

}

// HasUnknownCurrent determines whether the call to GetUnknownCurrent is safe
func (t *OrderedCollectionPage) HasUnknownCurrent() (ok bool) {
	return t.current != nil && t.current.unknown_ != nil

}

// GetUnknownCurrent returns the unknown value for current
func (t *OrderedCollectionPage) GetUnknownCurrent() (v interface{}) {
	return t.current.unknown_

}

// SetUnknownCurrent sets the unknown value of current
func (t *OrderedCollectionPage) SetUnknownCurrent(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &currentOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.current = tmp

}

// IsFirstOrderedCollectionPage determines whether the call to GetFirstOrderedCollectionPage is safe
func (t *OrderedCollectionPage) IsFirstOrderedCollectionPage() (ok bool) {
	return t.first != nil && t.first.OrderedCollectionPage != nil

}

// GetFirstOrderedCollectionPage returns the value safely if IsFirstOrderedCollectionPage returned true
func (t *OrderedCollectionPage) GetFirstOrderedCollectionPage() (v OrderedCollectionPageType) {
	return t.first.OrderedCollectionPage

}

// SetFirstOrderedCollectionPage sets the value of first to be of OrderedCollectionPageType type
func (t *OrderedCollectionPage) SetFirstOrderedCollectionPage(v OrderedCollectionPageType) {
	t.first = &firstOrderedCollectionPageIntermediateType{OrderedCollectionPage: v}

}

// IsFirstLink determines whether the call to GetFirstLink is safe
func (t *OrderedCollectionPage) IsFirstLink() (ok bool) {
	return t.first != nil && t.first.Link != nil

}

// GetFirstLink returns the value safely if IsFirstLink returned true
func (t *OrderedCollectionPage) GetFirstLink() (v LinkType) {
	return t.first.Link

}

// SetFirstLink sets the value of first to be of LinkType type
func (t *OrderedCollectionPage) SetFirstLink(v LinkType) {
	t.first = &firstOrderedCollectionPageIntermediateType{Link: v}

}

// IsFirstIRI determines whether the call to GetFirstIRI is safe
func (t *OrderedCollectionPage) IsFirstIRI() (ok bool) {
	return t.first != nil && t.first.IRI != nil

}

// GetFirstIRI returns the value safely if IsFirstIRI returned true
func (t *OrderedCollectionPage) GetFirstIRI() (v *url.URL) {
	return t.first.IRI

}

// SetFirstIRI sets the value of first to be of *url.URL type
func (t *OrderedCollectionPage) SetFirstIRI(v *url.URL) {
	t.first = &firstOrderedCollectionPageIntermediateType{IRI: v}

}

// HasUnknownFirst determines whether the call to GetUnknownFirst is safe
func (t *OrderedCollectionPage) HasUnknownFirst() (ok bool) {
	return t.first != nil && t.first.unknown_ != nil

}

// GetUnknownFirst returns the unknown value for first
func (t *OrderedCollectionPage) GetUnknownFirst() (v interface{}) {
	return t.first.unknown_

}

// SetUnknownFirst sets the unknown value of first
func (t *OrderedCollectionPage) SetUnknownFirst(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &firstOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.first = tmp

}

// IsLastOrderedCollectionPage determines whether the call to GetLastOrderedCollectionPage is safe
func (t *OrderedCollectionPage) IsLastOrderedCollectionPage() (ok bool) {
	return t.last != nil && t.last.OrderedCollectionPage != nil

}

// GetLastOrderedCollectionPage returns the value safely if IsLastOrderedCollectionPage returned true
func (t *OrderedCollectionPage) GetLastOrderedCollectionPage() (v OrderedCollectionPageType) {
	return t.last.OrderedCollectionPage

}

// SetLastOrderedCollectionPage sets the value of last to be of OrderedCollectionPageType type
func (t *OrderedCollectionPage) SetLastOrderedCollectionPage(v OrderedCollectionPageType) {
	t.last = &lastOrderedCollectionPageIntermediateType{OrderedCollectionPage: v}

}

// IsLastLink determines whether the call to GetLastLink is safe
func (t *OrderedCollectionPage) IsLastLink() (ok bool) {
	return t.last != nil && t.last.Link != nil

}

// GetLastLink returns the value safely if IsLastLink returned true
func (t *OrderedCollectionPage) GetLastLink() (v LinkType) {
	return t.last.Link

}

// SetLastLink sets the value of last to be of LinkType type
func (t *OrderedCollectionPage) SetLastLink(v LinkType) {
	t.last = &lastOrderedCollectionPageIntermediateType{Link: v}

}

// IsLastIRI determines whether the call to GetLastIRI is safe
func (t *OrderedCollectionPage) IsLastIRI() (ok bool) {
	return t.last != nil && t.last.IRI != nil

}

// GetLastIRI returns the value safely if IsLastIRI returned true
func (t *OrderedCollectionPage) GetLastIRI() (v *url.URL) {
	return t.last.IRI

}

// SetLastIRI sets the value of last to be of *url.URL type
func (t *OrderedCollectionPage) SetLastIRI(v *url.URL) {
	t.last = &lastOrderedCollectionPageIntermediateType{IRI: v}

}

// HasUnknownLast determines whether the call to GetUnknownLast is safe
func (t *OrderedCollectionPage) HasUnknownLast() (ok bool) {
	return t.last != nil && t.last.unknown_ != nil

}

// GetUnknownLast returns the unknown value for last
func (t *OrderedCollectionPage) GetUnknownLast() (v interface{}) {
	return t.last.unknown_

}

// SetUnknownLast sets the unknown value of last
func (t *OrderedCollectionPage) SetUnknownLast(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &lastOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.last = tmp

}

// IsTotalItems determines whether the call to GetTotalItems is safe
func (t *OrderedCollectionPage) IsTotalItems() (ok bool) {
	return t.totalItems != nil && t.totalItems.nonNegativeInteger != nil

}

// GetTotalItems returns the value safely if IsTotalItems returned true
func (t *OrderedCollectionPage) GetTotalItems() (v int64) {
	return *t.totalItems.nonNegativeInteger

}

// SetTotalItems sets the value of totalItems to be of int64 type
func (t *OrderedCollectionPage) SetTotalItems(v int64) {
	t.totalItems = &totalItemsOrderedCollectionPageIntermediateType{nonNegativeInteger: &v}

}

// IsTotalItemsIRI determines whether the call to GetTotalItemsIRI is safe
func (t *OrderedCollectionPage) IsTotalItemsIRI() (ok bool) {
	return t.totalItems != nil && t.totalItems.IRI != nil

}

// GetTotalItemsIRI returns the value safely if IsTotalItemsIRI returned true
func (t *OrderedCollectionPage) GetTotalItemsIRI() (v *url.URL) {
	return t.totalItems.IRI

}

// SetTotalItemsIRI sets the value of totalItems to be of *url.URL type
func (t *OrderedCollectionPage) SetTotalItemsIRI(v *url.URL) {
	t.totalItems = &totalItemsOrderedCollectionPageIntermediateType{IRI: v}

}

// HasUnknownTotalItems determines whether the call to GetUnknownTotalItems is safe
func (t *OrderedCollectionPage) HasUnknownTotalItems() (ok bool) {
	return t.totalItems != nil && t.totalItems.unknown_ != nil

}

// GetUnknownTotalItems returns the unknown value for totalItems
func (t *OrderedCollectionPage) GetUnknownTotalItems() (v interface{}) {
	return t.totalItems.unknown_

}

// SetUnknownTotalItems sets the unknown value of totalItems
func (t *OrderedCollectionPage) SetUnknownTotalItems(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &totalItemsOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.totalItems = tmp

}

// IsAltitude determines whether the call to GetAltitude is safe
func (t *OrderedCollectionPage) IsAltitude() (ok bool) {
	return t.altitude != nil && t.altitude.float != nil

}

// GetAltitude returns the value safely if IsAltitude returned true
func (t *OrderedCollectionPage) GetAltitude() (v float64) {
	return *t.altitude.float

}

// SetAltitude sets the value of altitude to be of float64 type
func (t *OrderedCollectionPage) SetAltitude(v float64) {
	t.altitude = &altitudeOrderedCollectionPageIntermediateType{float: &v}

}

// IsAltitudeIRI determines whether the call to GetAltitudeIRI is safe
func (t *OrderedCollectionPage) IsAltitudeIRI() (ok bool) {
	return t.altitude != nil && t.altitude.IRI != nil

}

// GetAltitudeIRI returns the value safely if IsAltitudeIRI returned true
func (t *OrderedCollectionPage) GetAltitudeIRI() (v *url.URL) {
	return t.altitude.IRI

}

// SetAltitudeIRI sets the value of altitude to be of *url.URL type
func (t *OrderedCollectionPage) SetAltitudeIRI(v *url.URL) {
	t.altitude = &altitudeOrderedCollectionPageIntermediateType{IRI: v}

}

// HasUnknownAltitude determines whether the call to GetUnknownAltitude is safe
func (t *OrderedCollectionPage) HasUnknownAltitude() (ok bool) {
	return t.altitude != nil && t.altitude.unknown_ != nil

}

// GetUnknownAltitude returns the unknown value for altitude
func (t *OrderedCollectionPage) GetUnknownAltitude() (v interface{}) {
	return t.altitude.unknown_

}

// SetUnknownAltitude sets the unknown value of altitude
func (t *OrderedCollectionPage) SetUnknownAltitude(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &altitudeOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.altitude = tmp

}

// AttachmentLen determines the number of elements able to be used for the IsAttachmentObject, GetAttachmentObject, and RemoveAttachmentObject functions
func (t *OrderedCollectionPage) AttachmentLen() (l int) {
	return len(t.attachment)

}

// IsAttachmentObject determines whether the call to GetAttachmentObject is safe for the specified index
func (t *OrderedCollectionPage) IsAttachmentObject(index int) (ok bool) {
	return t.attachment[index].Object != nil

}

// GetAttachmentObject returns the value safely if IsAttachmentObject returned true for the specified index
func (t *OrderedCollectionPage) GetAttachmentObject(index int) (v ObjectType) {
	return t.attachment[index].Object

}

// AppendAttachmentObject adds to the back of attachment a ObjectType type
func (t *OrderedCollectionPage) AppendAttachmentObject(v ObjectType) {
	t.attachment = append(t.attachment, &attachmentOrderedCollectionPageIntermediateType{Object: v})

}

// PrependAttachmentObject adds to the front of attachment a ObjectType type
func (t *OrderedCollectionPage) PrependAttachmentObject(v ObjectType) {
	t.attachment = append([]*attachmentOrderedCollectionPageIntermediateType{&attachmentOrderedCollectionPageIntermediateType{Object: v}}, t.attachment...)

}

// RemoveAttachmentObject deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveAttachmentObject(index int) {
	copy(t.attachment[index:], t.attachment[index+1:])
	t.attachment[len(t.attachment)-1] = nil
	t.attachment = t.attachment[:len(t.attachment)-1]

}

// IsAttachmentLink determines whether the call to GetAttachmentLink is safe for the specified index
func (t *OrderedCollectionPage) IsAttachmentLink(index int) (ok bool) {
	return t.attachment[index].Link != nil

}

// GetAttachmentLink returns the value safely if IsAttachmentLink returned true for the specified index
func (t *OrderedCollectionPage) GetAttachmentLink(index int) (v LinkType) {
	return t.attachment[index].Link

}

// AppendAttachmentLink adds to the back of attachment a LinkType type
func (t *OrderedCollectionPage) AppendAttachmentLink(v LinkType) {
	t.attachment = append(t.attachment, &attachmentOrderedCollectionPageIntermediateType{Link: v})

}

// PrependAttachmentLink adds to the front of attachment a LinkType type
func (t *OrderedCollectionPage) PrependAttachmentLink(v LinkType) {
	t.attachment = append([]*attachmentOrderedCollectionPageIntermediateType{&attachmentOrderedCollectionPageIntermediateType{Link: v}}, t.attachment...)

}

// RemoveAttachmentLink deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveAttachmentLink(index int) {
	copy(t.attachment[index:], t.attachment[index+1:])
	t.attachment[len(t.attachment)-1] = nil
	t.attachment = t.attachment[:len(t.attachment)-1]

}

// IsAttachmentIRI determines whether the call to GetAttachmentIRI is safe for the specified index
func (t *OrderedCollectionPage) IsAttachmentIRI(index int) (ok bool) {
	return t.attachment[index].IRI != nil

}

// GetAttachmentIRI returns the value safely if IsAttachmentIRI returned true for the specified index
func (t *OrderedCollectionPage) GetAttachmentIRI(index int) (v *url.URL) {
	return t.attachment[index].IRI

}

// AppendAttachmentIRI adds to the back of attachment a *url.URL type
func (t *OrderedCollectionPage) AppendAttachmentIRI(v *url.URL) {
	t.attachment = append(t.attachment, &attachmentOrderedCollectionPageIntermediateType{IRI: v})

}

// PrependAttachmentIRI adds to the front of attachment a *url.URL type
func (t *OrderedCollectionPage) PrependAttachmentIRI(v *url.URL) {
	t.attachment = append([]*attachmentOrderedCollectionPageIntermediateType{&attachmentOrderedCollectionPageIntermediateType{IRI: v}}, t.attachment...)

}

// RemoveAttachmentIRI deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveAttachmentIRI(index int) {
	copy(t.attachment[index:], t.attachment[index+1:])
	t.attachment[len(t.attachment)-1] = nil
	t.attachment = t.attachment[:len(t.attachment)-1]

}

// HasUnknownAttachment determines whether the call to GetUnknownAttachment is safe
func (t *OrderedCollectionPage) HasUnknownAttachment() (ok bool) {
	return t.attachment != nil && t.attachment[0].unknown_ != nil

}

// GetUnknownAttachment returns the unknown value for attachment
func (t *OrderedCollectionPage) GetUnknownAttachment() (v interface{}) {
	return t.attachment[0].unknown_

}

// SetUnknownAttachment sets the unknown value of attachment
func (t *OrderedCollectionPage) SetUnknownAttachment(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &attachmentOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.attachment = append(t.attachment, tmp)

}

// AttributedToLen determines the number of elements able to be used for the IsAttributedToObject, GetAttributedToObject, and RemoveAttributedToObject functions
func (t *OrderedCollectionPage) AttributedToLen() (l int) {
	return len(t.attributedTo)

}

// IsAttributedToObject determines whether the call to GetAttributedToObject is safe for the specified index
func (t *OrderedCollectionPage) IsAttributedToObject(index int) (ok bool) {
	return t.attributedTo[index].Object != nil

}

// GetAttributedToObject returns the value safely if IsAttributedToObject returned true for the specified index
func (t *OrderedCollectionPage) GetAttributedToObject(index int) (v ObjectType) {
	return t.attributedTo[index].Object

}

// AppendAttributedToObject adds to the back of attributedTo a ObjectType type
func (t *OrderedCollectionPage) AppendAttributedToObject(v ObjectType) {
	t.attributedTo = append(t.attributedTo, &attributedToOrderedCollectionPageIntermediateType{Object: v})

}

// PrependAttributedToObject adds to the front of attributedTo a ObjectType type
func (t *OrderedCollectionPage) PrependAttributedToObject(v ObjectType) {
	t.attributedTo = append([]*attributedToOrderedCollectionPageIntermediateType{&attributedToOrderedCollectionPageIntermediateType{Object: v}}, t.attributedTo...)

}

// RemoveAttributedToObject deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveAttributedToObject(index int) {
	copy(t.attributedTo[index:], t.attributedTo[index+1:])
	t.attributedTo[len(t.attributedTo)-1] = nil
	t.attributedTo = t.attributedTo[:len(t.attributedTo)-1]

}

// IsAttributedToLink determines whether the call to GetAttributedToLink is safe for the specified index
func (t *OrderedCollectionPage) IsAttributedToLink(index int) (ok bool) {
	return t.attributedTo[index].Link != nil

}

// GetAttributedToLink returns the value safely if IsAttributedToLink returned true for the specified index
func (t *OrderedCollectionPage) GetAttributedToLink(index int) (v LinkType) {
	return t.attributedTo[index].Link

}

// AppendAttributedToLink adds to the back of attributedTo a LinkType type
func (t *OrderedCollectionPage) AppendAttributedToLink(v LinkType) {
	t.attributedTo = append(t.attributedTo, &attributedToOrderedCollectionPageIntermediateType{Link: v})

}

// PrependAttributedToLink adds to the front of attributedTo a LinkType type
func (t *OrderedCollectionPage) PrependAttributedToLink(v LinkType) {
	t.attributedTo = append([]*attributedToOrderedCollectionPageIntermediateType{&attributedToOrderedCollectionPageIntermediateType{Link: v}}, t.attributedTo...)

}

// RemoveAttributedToLink deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveAttributedToLink(index int) {
	copy(t.attributedTo[index:], t.attributedTo[index+1:])
	t.attributedTo[len(t.attributedTo)-1] = nil
	t.attributedTo = t.attributedTo[:len(t.attributedTo)-1]

}

// IsAttributedToIRI determines whether the call to GetAttributedToIRI is safe for the specified index
func (t *OrderedCollectionPage) IsAttributedToIRI(index int) (ok bool) {
	return t.attributedTo[index].IRI != nil

}

// GetAttributedToIRI returns the value safely if IsAttributedToIRI returned true for the specified index
func (t *OrderedCollectionPage) GetAttributedToIRI(index int) (v *url.URL) {
	return t.attributedTo[index].IRI

}

// AppendAttributedToIRI adds to the back of attributedTo a *url.URL type
func (t *OrderedCollectionPage) AppendAttributedToIRI(v *url.URL) {
	t.attributedTo = append(t.attributedTo, &attributedToOrderedCollectionPageIntermediateType{IRI: v})

}

// PrependAttributedToIRI adds to the front of attributedTo a *url.URL type
func (t *OrderedCollectionPage) PrependAttributedToIRI(v *url.URL) {
	t.attributedTo = append([]*attributedToOrderedCollectionPageIntermediateType{&attributedToOrderedCollectionPageIntermediateType{IRI: v}}, t.attributedTo...)

}

// RemoveAttributedToIRI deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveAttributedToIRI(index int) {
	copy(t.attributedTo[index:], t.attributedTo[index+1:])
	t.attributedTo[len(t.attributedTo)-1] = nil
	t.attributedTo = t.attributedTo[:len(t.attributedTo)-1]

}

// HasUnknownAttributedTo determines whether the call to GetUnknownAttributedTo is safe
func (t *OrderedCollectionPage) HasUnknownAttributedTo() (ok bool) {
	return t.attributedTo != nil && t.attributedTo[0].unknown_ != nil

}

// GetUnknownAttributedTo returns the unknown value for attributedTo
func (t *OrderedCollectionPage) GetUnknownAttributedTo() (v interface{}) {
	return t.attributedTo[0].unknown_

}

// SetUnknownAttributedTo sets the unknown value of attributedTo
func (t *OrderedCollectionPage) SetUnknownAttributedTo(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &attributedToOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.attributedTo = append(t.attributedTo, tmp)

}

// AudienceLen determines the number of elements able to be used for the IsAudienceObject, GetAudienceObject, and RemoveAudienceObject functions
func (t *OrderedCollectionPage) AudienceLen() (l int) {
	return len(t.audience)

}

// IsAudienceObject determines whether the call to GetAudienceObject is safe for the specified index
func (t *OrderedCollectionPage) IsAudienceObject(index int) (ok bool) {
	return t.audience[index].Object != nil

}

// GetAudienceObject returns the value safely if IsAudienceObject returned true for the specified index
func (t *OrderedCollectionPage) GetAudienceObject(index int) (v ObjectType) {
	return t.audience[index].Object

}

// AppendAudienceObject adds to the back of audience a ObjectType type
func (t *OrderedCollectionPage) AppendAudienceObject(v ObjectType) {
	t.audience = append(t.audience, &audienceOrderedCollectionPageIntermediateType{Object: v})

}

// PrependAudienceObject adds to the front of audience a ObjectType type
func (t *OrderedCollectionPage) PrependAudienceObject(v ObjectType) {
	t.audience = append([]*audienceOrderedCollectionPageIntermediateType{&audienceOrderedCollectionPageIntermediateType{Object: v}}, t.audience...)

}

// RemoveAudienceObject deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveAudienceObject(index int) {
	copy(t.audience[index:], t.audience[index+1:])
	t.audience[len(t.audience)-1] = nil
	t.audience = t.audience[:len(t.audience)-1]

}

// IsAudienceLink determines whether the call to GetAudienceLink is safe for the specified index
func (t *OrderedCollectionPage) IsAudienceLink(index int) (ok bool) {
	return t.audience[index].Link != nil

}

// GetAudienceLink returns the value safely if IsAudienceLink returned true for the specified index
func (t *OrderedCollectionPage) GetAudienceLink(index int) (v LinkType) {
	return t.audience[index].Link

}

// AppendAudienceLink adds to the back of audience a LinkType type
func (t *OrderedCollectionPage) AppendAudienceLink(v LinkType) {
	t.audience = append(t.audience, &audienceOrderedCollectionPageIntermediateType{Link: v})

}

// PrependAudienceLink adds to the front of audience a LinkType type
func (t *OrderedCollectionPage) PrependAudienceLink(v LinkType) {
	t.audience = append([]*audienceOrderedCollectionPageIntermediateType{&audienceOrderedCollectionPageIntermediateType{Link: v}}, t.audience...)

}

// RemoveAudienceLink deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveAudienceLink(index int) {
	copy(t.audience[index:], t.audience[index+1:])
	t.audience[len(t.audience)-1] = nil
	t.audience = t.audience[:len(t.audience)-1]

}

// IsAudienceIRI determines whether the call to GetAudienceIRI is safe for the specified index
func (t *OrderedCollectionPage) IsAudienceIRI(index int) (ok bool) {
	return t.audience[index].IRI != nil

}

// GetAudienceIRI returns the value safely if IsAudienceIRI returned true for the specified index
func (t *OrderedCollectionPage) GetAudienceIRI(index int) (v *url.URL) {
	return t.audience[index].IRI

}

// AppendAudienceIRI adds to the back of audience a *url.URL type
func (t *OrderedCollectionPage) AppendAudienceIRI(v *url.URL) {
	t.audience = append(t.audience, &audienceOrderedCollectionPageIntermediateType{IRI: v})

}

// PrependAudienceIRI adds to the front of audience a *url.URL type
func (t *OrderedCollectionPage) PrependAudienceIRI(v *url.URL) {
	t.audience = append([]*audienceOrderedCollectionPageIntermediateType{&audienceOrderedCollectionPageIntermediateType{IRI: v}}, t.audience...)

}

// RemoveAudienceIRI deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveAudienceIRI(index int) {
	copy(t.audience[index:], t.audience[index+1:])
	t.audience[len(t.audience)-1] = nil
	t.audience = t.audience[:len(t.audience)-1]

}

// HasUnknownAudience determines whether the call to GetUnknownAudience is safe
func (t *OrderedCollectionPage) HasUnknownAudience() (ok bool) {
	return t.audience != nil && t.audience[0].unknown_ != nil

}

// GetUnknownAudience returns the unknown value for audience
func (t *OrderedCollectionPage) GetUnknownAudience() (v interface{}) {
	return t.audience[0].unknown_

}

// SetUnknownAudience sets the unknown value of audience
func (t *OrderedCollectionPage) SetUnknownAudience(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &audienceOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.audience = append(t.audience, tmp)

}

// ContentLen determines the number of elements able to be used for the IsContentString, GetContentString, and RemoveContentString functions
func (t *OrderedCollectionPage) ContentLen() (l int) {
	return len(t.content)

}

// IsContentString determines whether the call to GetContentString is safe for the specified index
func (t *OrderedCollectionPage) IsContentString(index int) (ok bool) {
	return t.content[index].stringName != nil

}

// GetContentString returns the value safely if IsContentString returned true for the specified index
func (t *OrderedCollectionPage) GetContentString(index int) (v string) {
	return *t.content[index].stringName

}

// AppendContentString adds to the back of content a string type
func (t *OrderedCollectionPage) AppendContentString(v string) {
	t.content = append(t.content, &contentOrderedCollectionPageIntermediateType{stringName: &v})

}

// PrependContentString adds to the front of content a string type
func (t *OrderedCollectionPage) PrependContentString(v string) {
	t.content = append([]*contentOrderedCollectionPageIntermediateType{&contentOrderedCollectionPageIntermediateType{stringName: &v}}, t.content...)

}

// RemoveContentString deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveContentString(index int) {
	copy(t.content[index:], t.content[index+1:])
	t.content[len(t.content)-1] = nil
	t.content = t.content[:len(t.content)-1]

}

// IsContentLangString determines whether the call to GetContentLangString is safe for the specified index
func (t *OrderedCollectionPage) IsContentLangString(index int) (ok bool) {
	return t.content[index].langString != nil

}

// GetContentLangString returns the value safely if IsContentLangString returned true for the specified index
func (t *OrderedCollectionPage) GetContentLangString(index int) (v string) {
	return *t.content[index].langString

}

// AppendContentLangString adds to the back of content a string type
func (t *OrderedCollectionPage) AppendContentLangString(v string) {
	t.content = append(t.content, &contentOrderedCollectionPageIntermediateType{langString: &v})

}

// PrependContentLangString adds to the front of content a string type
func (t *OrderedCollectionPage) PrependContentLangString(v string) {
	t.content = append([]*contentOrderedCollectionPageIntermediateType{&contentOrderedCollectionPageIntermediateType{langString: &v}}, t.content...)

}

// RemoveContentLangString deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveContentLangString(index int) {
	copy(t.content[index:], t.content[index+1:])
	t.content[len(t.content)-1] = nil
	t.content = t.content[:len(t.content)-1]

}

// IsContentIRI determines whether the call to GetContentIRI is safe for the specified index
func (t *OrderedCollectionPage) IsContentIRI(index int) (ok bool) {
	return t.content[index].IRI != nil

}

// GetContentIRI returns the value safely if IsContentIRI returned true for the specified index
func (t *OrderedCollectionPage) GetContentIRI(index int) (v *url.URL) {
	return t.content[index].IRI

}

// AppendContentIRI adds to the back of content a *url.URL type
func (t *OrderedCollectionPage) AppendContentIRI(v *url.URL) {
	t.content = append(t.content, &contentOrderedCollectionPageIntermediateType{IRI: v})

}

// PrependContentIRI adds to the front of content a *url.URL type
func (t *OrderedCollectionPage) PrependContentIRI(v *url.URL) {
	t.content = append([]*contentOrderedCollectionPageIntermediateType{&contentOrderedCollectionPageIntermediateType{IRI: v}}, t.content...)

}

// RemoveContentIRI deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveContentIRI(index int) {
	copy(t.content[index:], t.content[index+1:])
	t.content[len(t.content)-1] = nil
	t.content = t.content[:len(t.content)-1]

}

// HasUnknownContent determines whether the call to GetUnknownContent is safe
func (t *OrderedCollectionPage) HasUnknownContent() (ok bool) {
	return t.content != nil && t.content[0].unknown_ != nil

}

// GetUnknownContent returns the unknown value for content
func (t *OrderedCollectionPage) GetUnknownContent() (v interface{}) {
	return t.content[0].unknown_

}

// SetUnknownContent sets the unknown value of content
func (t *OrderedCollectionPage) SetUnknownContent(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &contentOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.content = append(t.content, tmp)

}

// ContentMapLanguages returns all languages for this property's language mapping, or nil if there are none.
func (t *OrderedCollectionPage) ContentMapLanguages() (l []string) {
	if t.contentMap == nil || len(t.contentMap) == 0 {
		return nil
	}
	for k := range t.contentMap {
		l = append(l, k)
	}
	return

}

// GetContentMap retrieves the value of the property for the specified language, or an empty string if it does not exist
func (t *OrderedCollectionPage) GetContentMap(l string) (v string) {
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
func (t *OrderedCollectionPage) SetContentMap(l string, v string) {
	if t.contentMap == nil {
		t.contentMap = make(map[string]string)
	}
	t.contentMap[l] = v

}

// ContextLen determines the number of elements able to be used for the IsContextObject, GetContextObject, and RemoveContextObject functions
func (t *OrderedCollectionPage) ContextLen() (l int) {
	return len(t.context)

}

// IsContextObject determines whether the call to GetContextObject is safe for the specified index
func (t *OrderedCollectionPage) IsContextObject(index int) (ok bool) {
	return t.context[index].Object != nil

}

// GetContextObject returns the value safely if IsContextObject returned true for the specified index
func (t *OrderedCollectionPage) GetContextObject(index int) (v ObjectType) {
	return t.context[index].Object

}

// AppendContextObject adds to the back of context a ObjectType type
func (t *OrderedCollectionPage) AppendContextObject(v ObjectType) {
	t.context = append(t.context, &contextOrderedCollectionPageIntermediateType{Object: v})

}

// PrependContextObject adds to the front of context a ObjectType type
func (t *OrderedCollectionPage) PrependContextObject(v ObjectType) {
	t.context = append([]*contextOrderedCollectionPageIntermediateType{&contextOrderedCollectionPageIntermediateType{Object: v}}, t.context...)

}

// RemoveContextObject deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveContextObject(index int) {
	copy(t.context[index:], t.context[index+1:])
	t.context[len(t.context)-1] = nil
	t.context = t.context[:len(t.context)-1]

}

// IsContextLink determines whether the call to GetContextLink is safe for the specified index
func (t *OrderedCollectionPage) IsContextLink(index int) (ok bool) {
	return t.context[index].Link != nil

}

// GetContextLink returns the value safely if IsContextLink returned true for the specified index
func (t *OrderedCollectionPage) GetContextLink(index int) (v LinkType) {
	return t.context[index].Link

}

// AppendContextLink adds to the back of context a LinkType type
func (t *OrderedCollectionPage) AppendContextLink(v LinkType) {
	t.context = append(t.context, &contextOrderedCollectionPageIntermediateType{Link: v})

}

// PrependContextLink adds to the front of context a LinkType type
func (t *OrderedCollectionPage) PrependContextLink(v LinkType) {
	t.context = append([]*contextOrderedCollectionPageIntermediateType{&contextOrderedCollectionPageIntermediateType{Link: v}}, t.context...)

}

// RemoveContextLink deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveContextLink(index int) {
	copy(t.context[index:], t.context[index+1:])
	t.context[len(t.context)-1] = nil
	t.context = t.context[:len(t.context)-1]

}

// IsContextIRI determines whether the call to GetContextIRI is safe for the specified index
func (t *OrderedCollectionPage) IsContextIRI(index int) (ok bool) {
	return t.context[index].IRI != nil

}

// GetContextIRI returns the value safely if IsContextIRI returned true for the specified index
func (t *OrderedCollectionPage) GetContextIRI(index int) (v *url.URL) {
	return t.context[index].IRI

}

// AppendContextIRI adds to the back of context a *url.URL type
func (t *OrderedCollectionPage) AppendContextIRI(v *url.URL) {
	t.context = append(t.context, &contextOrderedCollectionPageIntermediateType{IRI: v})

}

// PrependContextIRI adds to the front of context a *url.URL type
func (t *OrderedCollectionPage) PrependContextIRI(v *url.URL) {
	t.context = append([]*contextOrderedCollectionPageIntermediateType{&contextOrderedCollectionPageIntermediateType{IRI: v}}, t.context...)

}

// RemoveContextIRI deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveContextIRI(index int) {
	copy(t.context[index:], t.context[index+1:])
	t.context[len(t.context)-1] = nil
	t.context = t.context[:len(t.context)-1]

}

// HasUnknownContext determines whether the call to GetUnknownContext is safe
func (t *OrderedCollectionPage) HasUnknownContext() (ok bool) {
	return t.context != nil && t.context[0].unknown_ != nil

}

// GetUnknownContext returns the unknown value for context
func (t *OrderedCollectionPage) GetUnknownContext() (v interface{}) {
	return t.context[0].unknown_

}

// SetUnknownContext sets the unknown value of context
func (t *OrderedCollectionPage) SetUnknownContext(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &contextOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.context = append(t.context, tmp)

}

// NameLen determines the number of elements able to be used for the IsNameString, GetNameString, and RemoveNameString functions
func (t *OrderedCollectionPage) NameLen() (l int) {
	return len(t.name)

}

// IsNameString determines whether the call to GetNameString is safe for the specified index
func (t *OrderedCollectionPage) IsNameString(index int) (ok bool) {
	return t.name[index].stringName != nil

}

// GetNameString returns the value safely if IsNameString returned true for the specified index
func (t *OrderedCollectionPage) GetNameString(index int) (v string) {
	return *t.name[index].stringName

}

// AppendNameString adds to the back of name a string type
func (t *OrderedCollectionPage) AppendNameString(v string) {
	t.name = append(t.name, &nameOrderedCollectionPageIntermediateType{stringName: &v})

}

// PrependNameString adds to the front of name a string type
func (t *OrderedCollectionPage) PrependNameString(v string) {
	t.name = append([]*nameOrderedCollectionPageIntermediateType{&nameOrderedCollectionPageIntermediateType{stringName: &v}}, t.name...)

}

// RemoveNameString deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveNameString(index int) {
	copy(t.name[index:], t.name[index+1:])
	t.name[len(t.name)-1] = nil
	t.name = t.name[:len(t.name)-1]

}

// IsNameLangString determines whether the call to GetNameLangString is safe for the specified index
func (t *OrderedCollectionPage) IsNameLangString(index int) (ok bool) {
	return t.name[index].langString != nil

}

// GetNameLangString returns the value safely if IsNameLangString returned true for the specified index
func (t *OrderedCollectionPage) GetNameLangString(index int) (v string) {
	return *t.name[index].langString

}

// AppendNameLangString adds to the back of name a string type
func (t *OrderedCollectionPage) AppendNameLangString(v string) {
	t.name = append(t.name, &nameOrderedCollectionPageIntermediateType{langString: &v})

}

// PrependNameLangString adds to the front of name a string type
func (t *OrderedCollectionPage) PrependNameLangString(v string) {
	t.name = append([]*nameOrderedCollectionPageIntermediateType{&nameOrderedCollectionPageIntermediateType{langString: &v}}, t.name...)

}

// RemoveNameLangString deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveNameLangString(index int) {
	copy(t.name[index:], t.name[index+1:])
	t.name[len(t.name)-1] = nil
	t.name = t.name[:len(t.name)-1]

}

// IsNameIRI determines whether the call to GetNameIRI is safe for the specified index
func (t *OrderedCollectionPage) IsNameIRI(index int) (ok bool) {
	return t.name[index].IRI != nil

}

// GetNameIRI returns the value safely if IsNameIRI returned true for the specified index
func (t *OrderedCollectionPage) GetNameIRI(index int) (v *url.URL) {
	return t.name[index].IRI

}

// AppendNameIRI adds to the back of name a *url.URL type
func (t *OrderedCollectionPage) AppendNameIRI(v *url.URL) {
	t.name = append(t.name, &nameOrderedCollectionPageIntermediateType{IRI: v})

}

// PrependNameIRI adds to the front of name a *url.URL type
func (t *OrderedCollectionPage) PrependNameIRI(v *url.URL) {
	t.name = append([]*nameOrderedCollectionPageIntermediateType{&nameOrderedCollectionPageIntermediateType{IRI: v}}, t.name...)

}

// RemoveNameIRI deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveNameIRI(index int) {
	copy(t.name[index:], t.name[index+1:])
	t.name[len(t.name)-1] = nil
	t.name = t.name[:len(t.name)-1]

}

// HasUnknownName determines whether the call to GetUnknownName is safe
func (t *OrderedCollectionPage) HasUnknownName() (ok bool) {
	return t.name != nil && t.name[0].unknown_ != nil

}

// GetUnknownName returns the unknown value for name
func (t *OrderedCollectionPage) GetUnknownName() (v interface{}) {
	return t.name[0].unknown_

}

// SetUnknownName sets the unknown value of name
func (t *OrderedCollectionPage) SetUnknownName(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &nameOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.name = append(t.name, tmp)

}

// NameMapLanguages returns all languages for this property's language mapping, or nil if there are none.
func (t *OrderedCollectionPage) NameMapLanguages() (l []string) {
	if t.nameMap == nil || len(t.nameMap) == 0 {
		return nil
	}
	for k := range t.nameMap {
		l = append(l, k)
	}
	return

}

// GetNameMap retrieves the value of the property for the specified language, or an empty string if it does not exist
func (t *OrderedCollectionPage) GetNameMap(l string) (v string) {
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
func (t *OrderedCollectionPage) SetNameMap(l string, v string) {
	if t.nameMap == nil {
		t.nameMap = make(map[string]string)
	}
	t.nameMap[l] = v

}

// IsEndTime determines whether the call to GetEndTime is safe
func (t *OrderedCollectionPage) IsEndTime() (ok bool) {
	return t.endTime != nil && t.endTime.dateTime != nil

}

// GetEndTime returns the value safely if IsEndTime returned true
func (t *OrderedCollectionPage) GetEndTime() (v time.Time) {
	return *t.endTime.dateTime

}

// SetEndTime sets the value of endTime to be of time.Time type
func (t *OrderedCollectionPage) SetEndTime(v time.Time) {
	t.endTime = &endTimeOrderedCollectionPageIntermediateType{dateTime: &v}

}

// IsEndTimeIRI determines whether the call to GetEndTimeIRI is safe
func (t *OrderedCollectionPage) IsEndTimeIRI() (ok bool) {
	return t.endTime != nil && t.endTime.IRI != nil

}

// GetEndTimeIRI returns the value safely if IsEndTimeIRI returned true
func (t *OrderedCollectionPage) GetEndTimeIRI() (v *url.URL) {
	return t.endTime.IRI

}

// SetEndTimeIRI sets the value of endTime to be of *url.URL type
func (t *OrderedCollectionPage) SetEndTimeIRI(v *url.URL) {
	t.endTime = &endTimeOrderedCollectionPageIntermediateType{IRI: v}

}

// HasUnknownEndTime determines whether the call to GetUnknownEndTime is safe
func (t *OrderedCollectionPage) HasUnknownEndTime() (ok bool) {
	return t.endTime != nil && t.endTime.unknown_ != nil

}

// GetUnknownEndTime returns the unknown value for endTime
func (t *OrderedCollectionPage) GetUnknownEndTime() (v interface{}) {
	return t.endTime.unknown_

}

// SetUnknownEndTime sets the unknown value of endTime
func (t *OrderedCollectionPage) SetUnknownEndTime(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &endTimeOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.endTime = tmp

}

// GeneratorLen determines the number of elements able to be used for the IsGeneratorObject, GetGeneratorObject, and RemoveGeneratorObject functions
func (t *OrderedCollectionPage) GeneratorLen() (l int) {
	return len(t.generator)

}

// IsGeneratorObject determines whether the call to GetGeneratorObject is safe for the specified index
func (t *OrderedCollectionPage) IsGeneratorObject(index int) (ok bool) {
	return t.generator[index].Object != nil

}

// GetGeneratorObject returns the value safely if IsGeneratorObject returned true for the specified index
func (t *OrderedCollectionPage) GetGeneratorObject(index int) (v ObjectType) {
	return t.generator[index].Object

}

// AppendGeneratorObject adds to the back of generator a ObjectType type
func (t *OrderedCollectionPage) AppendGeneratorObject(v ObjectType) {
	t.generator = append(t.generator, &generatorOrderedCollectionPageIntermediateType{Object: v})

}

// PrependGeneratorObject adds to the front of generator a ObjectType type
func (t *OrderedCollectionPage) PrependGeneratorObject(v ObjectType) {
	t.generator = append([]*generatorOrderedCollectionPageIntermediateType{&generatorOrderedCollectionPageIntermediateType{Object: v}}, t.generator...)

}

// RemoveGeneratorObject deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveGeneratorObject(index int) {
	copy(t.generator[index:], t.generator[index+1:])
	t.generator[len(t.generator)-1] = nil
	t.generator = t.generator[:len(t.generator)-1]

}

// IsGeneratorLink determines whether the call to GetGeneratorLink is safe for the specified index
func (t *OrderedCollectionPage) IsGeneratorLink(index int) (ok bool) {
	return t.generator[index].Link != nil

}

// GetGeneratorLink returns the value safely if IsGeneratorLink returned true for the specified index
func (t *OrderedCollectionPage) GetGeneratorLink(index int) (v LinkType) {
	return t.generator[index].Link

}

// AppendGeneratorLink adds to the back of generator a LinkType type
func (t *OrderedCollectionPage) AppendGeneratorLink(v LinkType) {
	t.generator = append(t.generator, &generatorOrderedCollectionPageIntermediateType{Link: v})

}

// PrependGeneratorLink adds to the front of generator a LinkType type
func (t *OrderedCollectionPage) PrependGeneratorLink(v LinkType) {
	t.generator = append([]*generatorOrderedCollectionPageIntermediateType{&generatorOrderedCollectionPageIntermediateType{Link: v}}, t.generator...)

}

// RemoveGeneratorLink deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveGeneratorLink(index int) {
	copy(t.generator[index:], t.generator[index+1:])
	t.generator[len(t.generator)-1] = nil
	t.generator = t.generator[:len(t.generator)-1]

}

// IsGeneratorIRI determines whether the call to GetGeneratorIRI is safe for the specified index
func (t *OrderedCollectionPage) IsGeneratorIRI(index int) (ok bool) {
	return t.generator[index].IRI != nil

}

// GetGeneratorIRI returns the value safely if IsGeneratorIRI returned true for the specified index
func (t *OrderedCollectionPage) GetGeneratorIRI(index int) (v *url.URL) {
	return t.generator[index].IRI

}

// AppendGeneratorIRI adds to the back of generator a *url.URL type
func (t *OrderedCollectionPage) AppendGeneratorIRI(v *url.URL) {
	t.generator = append(t.generator, &generatorOrderedCollectionPageIntermediateType{IRI: v})

}

// PrependGeneratorIRI adds to the front of generator a *url.URL type
func (t *OrderedCollectionPage) PrependGeneratorIRI(v *url.URL) {
	t.generator = append([]*generatorOrderedCollectionPageIntermediateType{&generatorOrderedCollectionPageIntermediateType{IRI: v}}, t.generator...)

}

// RemoveGeneratorIRI deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveGeneratorIRI(index int) {
	copy(t.generator[index:], t.generator[index+1:])
	t.generator[len(t.generator)-1] = nil
	t.generator = t.generator[:len(t.generator)-1]

}

// HasUnknownGenerator determines whether the call to GetUnknownGenerator is safe
func (t *OrderedCollectionPage) HasUnknownGenerator() (ok bool) {
	return t.generator != nil && t.generator[0].unknown_ != nil

}

// GetUnknownGenerator returns the unknown value for generator
func (t *OrderedCollectionPage) GetUnknownGenerator() (v interface{}) {
	return t.generator[0].unknown_

}

// SetUnknownGenerator sets the unknown value of generator
func (t *OrderedCollectionPage) SetUnknownGenerator(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &generatorOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.generator = append(t.generator, tmp)

}

// IconLen determines the number of elements able to be used for the IsIconImage, GetIconImage, and RemoveIconImage functions
func (t *OrderedCollectionPage) IconLen() (l int) {
	return len(t.icon)

}

// IsIconImage determines whether the call to GetIconImage is safe for the specified index
func (t *OrderedCollectionPage) IsIconImage(index int) (ok bool) {
	return t.icon[index].Image != nil

}

// GetIconImage returns the value safely if IsIconImage returned true for the specified index
func (t *OrderedCollectionPage) GetIconImage(index int) (v ImageType) {
	return t.icon[index].Image

}

// AppendIconImage adds to the back of icon a ImageType type
func (t *OrderedCollectionPage) AppendIconImage(v ImageType) {
	t.icon = append(t.icon, &iconOrderedCollectionPageIntermediateType{Image: v})

}

// PrependIconImage adds to the front of icon a ImageType type
func (t *OrderedCollectionPage) PrependIconImage(v ImageType) {
	t.icon = append([]*iconOrderedCollectionPageIntermediateType{&iconOrderedCollectionPageIntermediateType{Image: v}}, t.icon...)

}

// RemoveIconImage deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveIconImage(index int) {
	copy(t.icon[index:], t.icon[index+1:])
	t.icon[len(t.icon)-1] = nil
	t.icon = t.icon[:len(t.icon)-1]

}

// IsIconLink determines whether the call to GetIconLink is safe for the specified index
func (t *OrderedCollectionPage) IsIconLink(index int) (ok bool) {
	return t.icon[index].Link != nil

}

// GetIconLink returns the value safely if IsIconLink returned true for the specified index
func (t *OrderedCollectionPage) GetIconLink(index int) (v LinkType) {
	return t.icon[index].Link

}

// AppendIconLink adds to the back of icon a LinkType type
func (t *OrderedCollectionPage) AppendIconLink(v LinkType) {
	t.icon = append(t.icon, &iconOrderedCollectionPageIntermediateType{Link: v})

}

// PrependIconLink adds to the front of icon a LinkType type
func (t *OrderedCollectionPage) PrependIconLink(v LinkType) {
	t.icon = append([]*iconOrderedCollectionPageIntermediateType{&iconOrderedCollectionPageIntermediateType{Link: v}}, t.icon...)

}

// RemoveIconLink deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveIconLink(index int) {
	copy(t.icon[index:], t.icon[index+1:])
	t.icon[len(t.icon)-1] = nil
	t.icon = t.icon[:len(t.icon)-1]

}

// IsIconIRI determines whether the call to GetIconIRI is safe for the specified index
func (t *OrderedCollectionPage) IsIconIRI(index int) (ok bool) {
	return t.icon[index].IRI != nil

}

// GetIconIRI returns the value safely if IsIconIRI returned true for the specified index
func (t *OrderedCollectionPage) GetIconIRI(index int) (v *url.URL) {
	return t.icon[index].IRI

}

// AppendIconIRI adds to the back of icon a *url.URL type
func (t *OrderedCollectionPage) AppendIconIRI(v *url.URL) {
	t.icon = append(t.icon, &iconOrderedCollectionPageIntermediateType{IRI: v})

}

// PrependIconIRI adds to the front of icon a *url.URL type
func (t *OrderedCollectionPage) PrependIconIRI(v *url.URL) {
	t.icon = append([]*iconOrderedCollectionPageIntermediateType{&iconOrderedCollectionPageIntermediateType{IRI: v}}, t.icon...)

}

// RemoveIconIRI deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveIconIRI(index int) {
	copy(t.icon[index:], t.icon[index+1:])
	t.icon[len(t.icon)-1] = nil
	t.icon = t.icon[:len(t.icon)-1]

}

// HasUnknownIcon determines whether the call to GetUnknownIcon is safe
func (t *OrderedCollectionPage) HasUnknownIcon() (ok bool) {
	return t.icon != nil && t.icon[0].unknown_ != nil

}

// GetUnknownIcon returns the unknown value for icon
func (t *OrderedCollectionPage) GetUnknownIcon() (v interface{}) {
	return t.icon[0].unknown_

}

// SetUnknownIcon sets the unknown value of icon
func (t *OrderedCollectionPage) SetUnknownIcon(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &iconOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.icon = append(t.icon, tmp)

}

// HasId determines whether the call to GetId is safe
func (t *OrderedCollectionPage) HasId() (ok bool) {
	return t.id != nil

}

// GetId returns the value for id
func (t *OrderedCollectionPage) GetId() (v *url.URL) {
	return t.id

}

// SetId sets the value of id
func (t *OrderedCollectionPage) SetId(v *url.URL) {
	t.id = v

}

// HasUnknownId determines whether the call to GetUnknownId is safe
func (t *OrderedCollectionPage) HasUnknownId() (ok bool) {
	return t.unknown_ != nil && t.unknown_["id"] != nil

}

// GetUnknownId returns the unknown value for id
func (t *OrderedCollectionPage) GetUnknownId() (v interface{}) {
	return t.unknown_["id"]

}

// SetUnknownId sets the unknown value of id
func (t *OrderedCollectionPage) SetUnknownId(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["id"] = i

}

// ImageLen determines the number of elements able to be used for the IsImageImage, GetImageImage, and RemoveImageImage functions
func (t *OrderedCollectionPage) ImageLen() (l int) {
	return len(t.image)

}

// IsImageImage determines whether the call to GetImageImage is safe for the specified index
func (t *OrderedCollectionPage) IsImageImage(index int) (ok bool) {
	return t.image[index].Image != nil

}

// GetImageImage returns the value safely if IsImageImage returned true for the specified index
func (t *OrderedCollectionPage) GetImageImage(index int) (v ImageType) {
	return t.image[index].Image

}

// AppendImageImage adds to the back of image a ImageType type
func (t *OrderedCollectionPage) AppendImageImage(v ImageType) {
	t.image = append(t.image, &imageOrderedCollectionPageIntermediateType{Image: v})

}

// PrependImageImage adds to the front of image a ImageType type
func (t *OrderedCollectionPage) PrependImageImage(v ImageType) {
	t.image = append([]*imageOrderedCollectionPageIntermediateType{&imageOrderedCollectionPageIntermediateType{Image: v}}, t.image...)

}

// RemoveImageImage deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveImageImage(index int) {
	copy(t.image[index:], t.image[index+1:])
	t.image[len(t.image)-1] = nil
	t.image = t.image[:len(t.image)-1]

}

// IsImageLink determines whether the call to GetImageLink is safe for the specified index
func (t *OrderedCollectionPage) IsImageLink(index int) (ok bool) {
	return t.image[index].Link != nil

}

// GetImageLink returns the value safely if IsImageLink returned true for the specified index
func (t *OrderedCollectionPage) GetImageLink(index int) (v LinkType) {
	return t.image[index].Link

}

// AppendImageLink adds to the back of image a LinkType type
func (t *OrderedCollectionPage) AppendImageLink(v LinkType) {
	t.image = append(t.image, &imageOrderedCollectionPageIntermediateType{Link: v})

}

// PrependImageLink adds to the front of image a LinkType type
func (t *OrderedCollectionPage) PrependImageLink(v LinkType) {
	t.image = append([]*imageOrderedCollectionPageIntermediateType{&imageOrderedCollectionPageIntermediateType{Link: v}}, t.image...)

}

// RemoveImageLink deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveImageLink(index int) {
	copy(t.image[index:], t.image[index+1:])
	t.image[len(t.image)-1] = nil
	t.image = t.image[:len(t.image)-1]

}

// IsImageIRI determines whether the call to GetImageIRI is safe for the specified index
func (t *OrderedCollectionPage) IsImageIRI(index int) (ok bool) {
	return t.image[index].IRI != nil

}

// GetImageIRI returns the value safely if IsImageIRI returned true for the specified index
func (t *OrderedCollectionPage) GetImageIRI(index int) (v *url.URL) {
	return t.image[index].IRI

}

// AppendImageIRI adds to the back of image a *url.URL type
func (t *OrderedCollectionPage) AppendImageIRI(v *url.URL) {
	t.image = append(t.image, &imageOrderedCollectionPageIntermediateType{IRI: v})

}

// PrependImageIRI adds to the front of image a *url.URL type
func (t *OrderedCollectionPage) PrependImageIRI(v *url.URL) {
	t.image = append([]*imageOrderedCollectionPageIntermediateType{&imageOrderedCollectionPageIntermediateType{IRI: v}}, t.image...)

}

// RemoveImageIRI deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveImageIRI(index int) {
	copy(t.image[index:], t.image[index+1:])
	t.image[len(t.image)-1] = nil
	t.image = t.image[:len(t.image)-1]

}

// HasUnknownImage determines whether the call to GetUnknownImage is safe
func (t *OrderedCollectionPage) HasUnknownImage() (ok bool) {
	return t.image != nil && t.image[0].unknown_ != nil

}

// GetUnknownImage returns the unknown value for image
func (t *OrderedCollectionPage) GetUnknownImage() (v interface{}) {
	return t.image[0].unknown_

}

// SetUnknownImage sets the unknown value of image
func (t *OrderedCollectionPage) SetUnknownImage(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &imageOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.image = append(t.image, tmp)

}

// InReplyToLen determines the number of elements able to be used for the IsInReplyToObject, GetInReplyToObject, and RemoveInReplyToObject functions
func (t *OrderedCollectionPage) InReplyToLen() (l int) {
	return len(t.inReplyTo)

}

// IsInReplyToObject determines whether the call to GetInReplyToObject is safe for the specified index
func (t *OrderedCollectionPage) IsInReplyToObject(index int) (ok bool) {
	return t.inReplyTo[index].Object != nil

}

// GetInReplyToObject returns the value safely if IsInReplyToObject returned true for the specified index
func (t *OrderedCollectionPage) GetInReplyToObject(index int) (v ObjectType) {
	return t.inReplyTo[index].Object

}

// AppendInReplyToObject adds to the back of inReplyTo a ObjectType type
func (t *OrderedCollectionPage) AppendInReplyToObject(v ObjectType) {
	t.inReplyTo = append(t.inReplyTo, &inReplyToOrderedCollectionPageIntermediateType{Object: v})

}

// PrependInReplyToObject adds to the front of inReplyTo a ObjectType type
func (t *OrderedCollectionPage) PrependInReplyToObject(v ObjectType) {
	t.inReplyTo = append([]*inReplyToOrderedCollectionPageIntermediateType{&inReplyToOrderedCollectionPageIntermediateType{Object: v}}, t.inReplyTo...)

}

// RemoveInReplyToObject deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveInReplyToObject(index int) {
	copy(t.inReplyTo[index:], t.inReplyTo[index+1:])
	t.inReplyTo[len(t.inReplyTo)-1] = nil
	t.inReplyTo = t.inReplyTo[:len(t.inReplyTo)-1]

}

// IsInReplyToLink determines whether the call to GetInReplyToLink is safe for the specified index
func (t *OrderedCollectionPage) IsInReplyToLink(index int) (ok bool) {
	return t.inReplyTo[index].Link != nil

}

// GetInReplyToLink returns the value safely if IsInReplyToLink returned true for the specified index
func (t *OrderedCollectionPage) GetInReplyToLink(index int) (v LinkType) {
	return t.inReplyTo[index].Link

}

// AppendInReplyToLink adds to the back of inReplyTo a LinkType type
func (t *OrderedCollectionPage) AppendInReplyToLink(v LinkType) {
	t.inReplyTo = append(t.inReplyTo, &inReplyToOrderedCollectionPageIntermediateType{Link: v})

}

// PrependInReplyToLink adds to the front of inReplyTo a LinkType type
func (t *OrderedCollectionPage) PrependInReplyToLink(v LinkType) {
	t.inReplyTo = append([]*inReplyToOrderedCollectionPageIntermediateType{&inReplyToOrderedCollectionPageIntermediateType{Link: v}}, t.inReplyTo...)

}

// RemoveInReplyToLink deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveInReplyToLink(index int) {
	copy(t.inReplyTo[index:], t.inReplyTo[index+1:])
	t.inReplyTo[len(t.inReplyTo)-1] = nil
	t.inReplyTo = t.inReplyTo[:len(t.inReplyTo)-1]

}

// IsInReplyToIRI determines whether the call to GetInReplyToIRI is safe for the specified index
func (t *OrderedCollectionPage) IsInReplyToIRI(index int) (ok bool) {
	return t.inReplyTo[index].IRI != nil

}

// GetInReplyToIRI returns the value safely if IsInReplyToIRI returned true for the specified index
func (t *OrderedCollectionPage) GetInReplyToIRI(index int) (v *url.URL) {
	return t.inReplyTo[index].IRI

}

// AppendInReplyToIRI adds to the back of inReplyTo a *url.URL type
func (t *OrderedCollectionPage) AppendInReplyToIRI(v *url.URL) {
	t.inReplyTo = append(t.inReplyTo, &inReplyToOrderedCollectionPageIntermediateType{IRI: v})

}

// PrependInReplyToIRI adds to the front of inReplyTo a *url.URL type
func (t *OrderedCollectionPage) PrependInReplyToIRI(v *url.URL) {
	t.inReplyTo = append([]*inReplyToOrderedCollectionPageIntermediateType{&inReplyToOrderedCollectionPageIntermediateType{IRI: v}}, t.inReplyTo...)

}

// RemoveInReplyToIRI deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveInReplyToIRI(index int) {
	copy(t.inReplyTo[index:], t.inReplyTo[index+1:])
	t.inReplyTo[len(t.inReplyTo)-1] = nil
	t.inReplyTo = t.inReplyTo[:len(t.inReplyTo)-1]

}

// HasUnknownInReplyTo determines whether the call to GetUnknownInReplyTo is safe
func (t *OrderedCollectionPage) HasUnknownInReplyTo() (ok bool) {
	return t.inReplyTo != nil && t.inReplyTo[0].unknown_ != nil

}

// GetUnknownInReplyTo returns the unknown value for inReplyTo
func (t *OrderedCollectionPage) GetUnknownInReplyTo() (v interface{}) {
	return t.inReplyTo[0].unknown_

}

// SetUnknownInReplyTo sets the unknown value of inReplyTo
func (t *OrderedCollectionPage) SetUnknownInReplyTo(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &inReplyToOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.inReplyTo = append(t.inReplyTo, tmp)

}

// LocationLen determines the number of elements able to be used for the IsLocationObject, GetLocationObject, and RemoveLocationObject functions
func (t *OrderedCollectionPage) LocationLen() (l int) {
	return len(t.location)

}

// IsLocationObject determines whether the call to GetLocationObject is safe for the specified index
func (t *OrderedCollectionPage) IsLocationObject(index int) (ok bool) {
	return t.location[index].Object != nil

}

// GetLocationObject returns the value safely if IsLocationObject returned true for the specified index
func (t *OrderedCollectionPage) GetLocationObject(index int) (v ObjectType) {
	return t.location[index].Object

}

// AppendLocationObject adds to the back of location a ObjectType type
func (t *OrderedCollectionPage) AppendLocationObject(v ObjectType) {
	t.location = append(t.location, &locationOrderedCollectionPageIntermediateType{Object: v})

}

// PrependLocationObject adds to the front of location a ObjectType type
func (t *OrderedCollectionPage) PrependLocationObject(v ObjectType) {
	t.location = append([]*locationOrderedCollectionPageIntermediateType{&locationOrderedCollectionPageIntermediateType{Object: v}}, t.location...)

}

// RemoveLocationObject deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveLocationObject(index int) {
	copy(t.location[index:], t.location[index+1:])
	t.location[len(t.location)-1] = nil
	t.location = t.location[:len(t.location)-1]

}

// IsLocationLink determines whether the call to GetLocationLink is safe for the specified index
func (t *OrderedCollectionPage) IsLocationLink(index int) (ok bool) {
	return t.location[index].Link != nil

}

// GetLocationLink returns the value safely if IsLocationLink returned true for the specified index
func (t *OrderedCollectionPage) GetLocationLink(index int) (v LinkType) {
	return t.location[index].Link

}

// AppendLocationLink adds to the back of location a LinkType type
func (t *OrderedCollectionPage) AppendLocationLink(v LinkType) {
	t.location = append(t.location, &locationOrderedCollectionPageIntermediateType{Link: v})

}

// PrependLocationLink adds to the front of location a LinkType type
func (t *OrderedCollectionPage) PrependLocationLink(v LinkType) {
	t.location = append([]*locationOrderedCollectionPageIntermediateType{&locationOrderedCollectionPageIntermediateType{Link: v}}, t.location...)

}

// RemoveLocationLink deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveLocationLink(index int) {
	copy(t.location[index:], t.location[index+1:])
	t.location[len(t.location)-1] = nil
	t.location = t.location[:len(t.location)-1]

}

// IsLocationIRI determines whether the call to GetLocationIRI is safe for the specified index
func (t *OrderedCollectionPage) IsLocationIRI(index int) (ok bool) {
	return t.location[index].IRI != nil

}

// GetLocationIRI returns the value safely if IsLocationIRI returned true for the specified index
func (t *OrderedCollectionPage) GetLocationIRI(index int) (v *url.URL) {
	return t.location[index].IRI

}

// AppendLocationIRI adds to the back of location a *url.URL type
func (t *OrderedCollectionPage) AppendLocationIRI(v *url.URL) {
	t.location = append(t.location, &locationOrderedCollectionPageIntermediateType{IRI: v})

}

// PrependLocationIRI adds to the front of location a *url.URL type
func (t *OrderedCollectionPage) PrependLocationIRI(v *url.URL) {
	t.location = append([]*locationOrderedCollectionPageIntermediateType{&locationOrderedCollectionPageIntermediateType{IRI: v}}, t.location...)

}

// RemoveLocationIRI deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveLocationIRI(index int) {
	copy(t.location[index:], t.location[index+1:])
	t.location[len(t.location)-1] = nil
	t.location = t.location[:len(t.location)-1]

}

// HasUnknownLocation determines whether the call to GetUnknownLocation is safe
func (t *OrderedCollectionPage) HasUnknownLocation() (ok bool) {
	return t.location != nil && t.location[0].unknown_ != nil

}

// GetUnknownLocation returns the unknown value for location
func (t *OrderedCollectionPage) GetUnknownLocation() (v interface{}) {
	return t.location[0].unknown_

}

// SetUnknownLocation sets the unknown value of location
func (t *OrderedCollectionPage) SetUnknownLocation(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &locationOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.location = append(t.location, tmp)

}

// PreviewLen determines the number of elements able to be used for the IsPreviewObject, GetPreviewObject, and RemovePreviewObject functions
func (t *OrderedCollectionPage) PreviewLen() (l int) {
	return len(t.preview)

}

// IsPreviewObject determines whether the call to GetPreviewObject is safe for the specified index
func (t *OrderedCollectionPage) IsPreviewObject(index int) (ok bool) {
	return t.preview[index].Object != nil

}

// GetPreviewObject returns the value safely if IsPreviewObject returned true for the specified index
func (t *OrderedCollectionPage) GetPreviewObject(index int) (v ObjectType) {
	return t.preview[index].Object

}

// AppendPreviewObject adds to the back of preview a ObjectType type
func (t *OrderedCollectionPage) AppendPreviewObject(v ObjectType) {
	t.preview = append(t.preview, &previewOrderedCollectionPageIntermediateType{Object: v})

}

// PrependPreviewObject adds to the front of preview a ObjectType type
func (t *OrderedCollectionPage) PrependPreviewObject(v ObjectType) {
	t.preview = append([]*previewOrderedCollectionPageIntermediateType{&previewOrderedCollectionPageIntermediateType{Object: v}}, t.preview...)

}

// RemovePreviewObject deletes the value from the specified index
func (t *OrderedCollectionPage) RemovePreviewObject(index int) {
	copy(t.preview[index:], t.preview[index+1:])
	t.preview[len(t.preview)-1] = nil
	t.preview = t.preview[:len(t.preview)-1]

}

// IsPreviewLink determines whether the call to GetPreviewLink is safe for the specified index
func (t *OrderedCollectionPage) IsPreviewLink(index int) (ok bool) {
	return t.preview[index].Link != nil

}

// GetPreviewLink returns the value safely if IsPreviewLink returned true for the specified index
func (t *OrderedCollectionPage) GetPreviewLink(index int) (v LinkType) {
	return t.preview[index].Link

}

// AppendPreviewLink adds to the back of preview a LinkType type
func (t *OrderedCollectionPage) AppendPreviewLink(v LinkType) {
	t.preview = append(t.preview, &previewOrderedCollectionPageIntermediateType{Link: v})

}

// PrependPreviewLink adds to the front of preview a LinkType type
func (t *OrderedCollectionPage) PrependPreviewLink(v LinkType) {
	t.preview = append([]*previewOrderedCollectionPageIntermediateType{&previewOrderedCollectionPageIntermediateType{Link: v}}, t.preview...)

}

// RemovePreviewLink deletes the value from the specified index
func (t *OrderedCollectionPage) RemovePreviewLink(index int) {
	copy(t.preview[index:], t.preview[index+1:])
	t.preview[len(t.preview)-1] = nil
	t.preview = t.preview[:len(t.preview)-1]

}

// IsPreviewIRI determines whether the call to GetPreviewIRI is safe for the specified index
func (t *OrderedCollectionPage) IsPreviewIRI(index int) (ok bool) {
	return t.preview[index].IRI != nil

}

// GetPreviewIRI returns the value safely if IsPreviewIRI returned true for the specified index
func (t *OrderedCollectionPage) GetPreviewIRI(index int) (v *url.URL) {
	return t.preview[index].IRI

}

// AppendPreviewIRI adds to the back of preview a *url.URL type
func (t *OrderedCollectionPage) AppendPreviewIRI(v *url.URL) {
	t.preview = append(t.preview, &previewOrderedCollectionPageIntermediateType{IRI: v})

}

// PrependPreviewIRI adds to the front of preview a *url.URL type
func (t *OrderedCollectionPage) PrependPreviewIRI(v *url.URL) {
	t.preview = append([]*previewOrderedCollectionPageIntermediateType{&previewOrderedCollectionPageIntermediateType{IRI: v}}, t.preview...)

}

// RemovePreviewIRI deletes the value from the specified index
func (t *OrderedCollectionPage) RemovePreviewIRI(index int) {
	copy(t.preview[index:], t.preview[index+1:])
	t.preview[len(t.preview)-1] = nil
	t.preview = t.preview[:len(t.preview)-1]

}

// HasUnknownPreview determines whether the call to GetUnknownPreview is safe
func (t *OrderedCollectionPage) HasUnknownPreview() (ok bool) {
	return t.preview != nil && t.preview[0].unknown_ != nil

}

// GetUnknownPreview returns the unknown value for preview
func (t *OrderedCollectionPage) GetUnknownPreview() (v interface{}) {
	return t.preview[0].unknown_

}

// SetUnknownPreview sets the unknown value of preview
func (t *OrderedCollectionPage) SetUnknownPreview(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &previewOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.preview = append(t.preview, tmp)

}

// IsPublished determines whether the call to GetPublished is safe
func (t *OrderedCollectionPage) IsPublished() (ok bool) {
	return t.published != nil && t.published.dateTime != nil

}

// GetPublished returns the value safely if IsPublished returned true
func (t *OrderedCollectionPage) GetPublished() (v time.Time) {
	return *t.published.dateTime

}

// SetPublished sets the value of published to be of time.Time type
func (t *OrderedCollectionPage) SetPublished(v time.Time) {
	t.published = &publishedOrderedCollectionPageIntermediateType{dateTime: &v}

}

// IsPublishedIRI determines whether the call to GetPublishedIRI is safe
func (t *OrderedCollectionPage) IsPublishedIRI() (ok bool) {
	return t.published != nil && t.published.IRI != nil

}

// GetPublishedIRI returns the value safely if IsPublishedIRI returned true
func (t *OrderedCollectionPage) GetPublishedIRI() (v *url.URL) {
	return t.published.IRI

}

// SetPublishedIRI sets the value of published to be of *url.URL type
func (t *OrderedCollectionPage) SetPublishedIRI(v *url.URL) {
	t.published = &publishedOrderedCollectionPageIntermediateType{IRI: v}

}

// HasUnknownPublished determines whether the call to GetUnknownPublished is safe
func (t *OrderedCollectionPage) HasUnknownPublished() (ok bool) {
	return t.published != nil && t.published.unknown_ != nil

}

// GetUnknownPublished returns the unknown value for published
func (t *OrderedCollectionPage) GetUnknownPublished() (v interface{}) {
	return t.published.unknown_

}

// SetUnknownPublished sets the unknown value of published
func (t *OrderedCollectionPage) SetUnknownPublished(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &publishedOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.published = tmp

}

// IsReplies determines whether the call to GetReplies is safe
func (t *OrderedCollectionPage) IsReplies() (ok bool) {
	return t.replies != nil && t.replies.Collection != nil

}

// GetReplies returns the value safely if IsReplies returned true
func (t *OrderedCollectionPage) GetReplies() (v CollectionType) {
	return t.replies.Collection

}

// SetReplies sets the value of replies to be of CollectionType type
func (t *OrderedCollectionPage) SetReplies(v CollectionType) {
	t.replies = &repliesOrderedCollectionPageIntermediateType{Collection: v}

}

// IsRepliesIRI determines whether the call to GetRepliesIRI is safe
func (t *OrderedCollectionPage) IsRepliesIRI() (ok bool) {
	return t.replies != nil && t.replies.IRI != nil

}

// GetRepliesIRI returns the value safely if IsRepliesIRI returned true
func (t *OrderedCollectionPage) GetRepliesIRI() (v *url.URL) {
	return t.replies.IRI

}

// SetRepliesIRI sets the value of replies to be of *url.URL type
func (t *OrderedCollectionPage) SetRepliesIRI(v *url.URL) {
	t.replies = &repliesOrderedCollectionPageIntermediateType{IRI: v}

}

// HasUnknownReplies determines whether the call to GetUnknownReplies is safe
func (t *OrderedCollectionPage) HasUnknownReplies() (ok bool) {
	return t.replies != nil && t.replies.unknown_ != nil

}

// GetUnknownReplies returns the unknown value for replies
func (t *OrderedCollectionPage) GetUnknownReplies() (v interface{}) {
	return t.replies.unknown_

}

// SetUnknownReplies sets the unknown value of replies
func (t *OrderedCollectionPage) SetUnknownReplies(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &repliesOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.replies = tmp

}

// IsStartTime determines whether the call to GetStartTime is safe
func (t *OrderedCollectionPage) IsStartTime() (ok bool) {
	return t.startTime != nil && t.startTime.dateTime != nil

}

// GetStartTime returns the value safely if IsStartTime returned true
func (t *OrderedCollectionPage) GetStartTime() (v time.Time) {
	return *t.startTime.dateTime

}

// SetStartTime sets the value of startTime to be of time.Time type
func (t *OrderedCollectionPage) SetStartTime(v time.Time) {
	t.startTime = &startTimeOrderedCollectionPageIntermediateType{dateTime: &v}

}

// IsStartTimeIRI determines whether the call to GetStartTimeIRI is safe
func (t *OrderedCollectionPage) IsStartTimeIRI() (ok bool) {
	return t.startTime != nil && t.startTime.IRI != nil

}

// GetStartTimeIRI returns the value safely if IsStartTimeIRI returned true
func (t *OrderedCollectionPage) GetStartTimeIRI() (v *url.URL) {
	return t.startTime.IRI

}

// SetStartTimeIRI sets the value of startTime to be of *url.URL type
func (t *OrderedCollectionPage) SetStartTimeIRI(v *url.URL) {
	t.startTime = &startTimeOrderedCollectionPageIntermediateType{IRI: v}

}

// HasUnknownStartTime determines whether the call to GetUnknownStartTime is safe
func (t *OrderedCollectionPage) HasUnknownStartTime() (ok bool) {
	return t.startTime != nil && t.startTime.unknown_ != nil

}

// GetUnknownStartTime returns the unknown value for startTime
func (t *OrderedCollectionPage) GetUnknownStartTime() (v interface{}) {
	return t.startTime.unknown_

}

// SetUnknownStartTime sets the unknown value of startTime
func (t *OrderedCollectionPage) SetUnknownStartTime(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &startTimeOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.startTime = tmp

}

// SummaryLen determines the number of elements able to be used for the IsSummaryString, GetSummaryString, and RemoveSummaryString functions
func (t *OrderedCollectionPage) SummaryLen() (l int) {
	return len(t.summary)

}

// IsSummaryString determines whether the call to GetSummaryString is safe for the specified index
func (t *OrderedCollectionPage) IsSummaryString(index int) (ok bool) {
	return t.summary[index].stringName != nil

}

// GetSummaryString returns the value safely if IsSummaryString returned true for the specified index
func (t *OrderedCollectionPage) GetSummaryString(index int) (v string) {
	return *t.summary[index].stringName

}

// AppendSummaryString adds to the back of summary a string type
func (t *OrderedCollectionPage) AppendSummaryString(v string) {
	t.summary = append(t.summary, &summaryOrderedCollectionPageIntermediateType{stringName: &v})

}

// PrependSummaryString adds to the front of summary a string type
func (t *OrderedCollectionPage) PrependSummaryString(v string) {
	t.summary = append([]*summaryOrderedCollectionPageIntermediateType{&summaryOrderedCollectionPageIntermediateType{stringName: &v}}, t.summary...)

}

// RemoveSummaryString deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveSummaryString(index int) {
	copy(t.summary[index:], t.summary[index+1:])
	t.summary[len(t.summary)-1] = nil
	t.summary = t.summary[:len(t.summary)-1]

}

// IsSummaryLangString determines whether the call to GetSummaryLangString is safe for the specified index
func (t *OrderedCollectionPage) IsSummaryLangString(index int) (ok bool) {
	return t.summary[index].langString != nil

}

// GetSummaryLangString returns the value safely if IsSummaryLangString returned true for the specified index
func (t *OrderedCollectionPage) GetSummaryLangString(index int) (v string) {
	return *t.summary[index].langString

}

// AppendSummaryLangString adds to the back of summary a string type
func (t *OrderedCollectionPage) AppendSummaryLangString(v string) {
	t.summary = append(t.summary, &summaryOrderedCollectionPageIntermediateType{langString: &v})

}

// PrependSummaryLangString adds to the front of summary a string type
func (t *OrderedCollectionPage) PrependSummaryLangString(v string) {
	t.summary = append([]*summaryOrderedCollectionPageIntermediateType{&summaryOrderedCollectionPageIntermediateType{langString: &v}}, t.summary...)

}

// RemoveSummaryLangString deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveSummaryLangString(index int) {
	copy(t.summary[index:], t.summary[index+1:])
	t.summary[len(t.summary)-1] = nil
	t.summary = t.summary[:len(t.summary)-1]

}

// IsSummaryIRI determines whether the call to GetSummaryIRI is safe for the specified index
func (t *OrderedCollectionPage) IsSummaryIRI(index int) (ok bool) {
	return t.summary[index].IRI != nil

}

// GetSummaryIRI returns the value safely if IsSummaryIRI returned true for the specified index
func (t *OrderedCollectionPage) GetSummaryIRI(index int) (v *url.URL) {
	return t.summary[index].IRI

}

// AppendSummaryIRI adds to the back of summary a *url.URL type
func (t *OrderedCollectionPage) AppendSummaryIRI(v *url.URL) {
	t.summary = append(t.summary, &summaryOrderedCollectionPageIntermediateType{IRI: v})

}

// PrependSummaryIRI adds to the front of summary a *url.URL type
func (t *OrderedCollectionPage) PrependSummaryIRI(v *url.URL) {
	t.summary = append([]*summaryOrderedCollectionPageIntermediateType{&summaryOrderedCollectionPageIntermediateType{IRI: v}}, t.summary...)

}

// RemoveSummaryIRI deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveSummaryIRI(index int) {
	copy(t.summary[index:], t.summary[index+1:])
	t.summary[len(t.summary)-1] = nil
	t.summary = t.summary[:len(t.summary)-1]

}

// HasUnknownSummary determines whether the call to GetUnknownSummary is safe
func (t *OrderedCollectionPage) HasUnknownSummary() (ok bool) {
	return t.summary != nil && t.summary[0].unknown_ != nil

}

// GetUnknownSummary returns the unknown value for summary
func (t *OrderedCollectionPage) GetUnknownSummary() (v interface{}) {
	return t.summary[0].unknown_

}

// SetUnknownSummary sets the unknown value of summary
func (t *OrderedCollectionPage) SetUnknownSummary(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &summaryOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.summary = append(t.summary, tmp)

}

// SummaryMapLanguages returns all languages for this property's language mapping, or nil if there are none.
func (t *OrderedCollectionPage) SummaryMapLanguages() (l []string) {
	if t.summaryMap == nil || len(t.summaryMap) == 0 {
		return nil
	}
	for k := range t.summaryMap {
		l = append(l, k)
	}
	return

}

// GetSummaryMap retrieves the value of the property for the specified language, or an empty string if it does not exist
func (t *OrderedCollectionPage) GetSummaryMap(l string) (v string) {
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
func (t *OrderedCollectionPage) SetSummaryMap(l string, v string) {
	if t.summaryMap == nil {
		t.summaryMap = make(map[string]string)
	}
	t.summaryMap[l] = v

}

// TagLen determines the number of elements able to be used for the IsTagObject, GetTagObject, and RemoveTagObject functions
func (t *OrderedCollectionPage) TagLen() (l int) {
	return len(t.tag)

}

// IsTagObject determines whether the call to GetTagObject is safe for the specified index
func (t *OrderedCollectionPage) IsTagObject(index int) (ok bool) {
	return t.tag[index].Object != nil

}

// GetTagObject returns the value safely if IsTagObject returned true for the specified index
func (t *OrderedCollectionPage) GetTagObject(index int) (v ObjectType) {
	return t.tag[index].Object

}

// AppendTagObject adds to the back of tag a ObjectType type
func (t *OrderedCollectionPage) AppendTagObject(v ObjectType) {
	t.tag = append(t.tag, &tagOrderedCollectionPageIntermediateType{Object: v})

}

// PrependTagObject adds to the front of tag a ObjectType type
func (t *OrderedCollectionPage) PrependTagObject(v ObjectType) {
	t.tag = append([]*tagOrderedCollectionPageIntermediateType{&tagOrderedCollectionPageIntermediateType{Object: v}}, t.tag...)

}

// RemoveTagObject deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveTagObject(index int) {
	copy(t.tag[index:], t.tag[index+1:])
	t.tag[len(t.tag)-1] = nil
	t.tag = t.tag[:len(t.tag)-1]

}

// IsTagLink determines whether the call to GetTagLink is safe for the specified index
func (t *OrderedCollectionPage) IsTagLink(index int) (ok bool) {
	return t.tag[index].Link != nil

}

// GetTagLink returns the value safely if IsTagLink returned true for the specified index
func (t *OrderedCollectionPage) GetTagLink(index int) (v LinkType) {
	return t.tag[index].Link

}

// AppendTagLink adds to the back of tag a LinkType type
func (t *OrderedCollectionPage) AppendTagLink(v LinkType) {
	t.tag = append(t.tag, &tagOrderedCollectionPageIntermediateType{Link: v})

}

// PrependTagLink adds to the front of tag a LinkType type
func (t *OrderedCollectionPage) PrependTagLink(v LinkType) {
	t.tag = append([]*tagOrderedCollectionPageIntermediateType{&tagOrderedCollectionPageIntermediateType{Link: v}}, t.tag...)

}

// RemoveTagLink deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveTagLink(index int) {
	copy(t.tag[index:], t.tag[index+1:])
	t.tag[len(t.tag)-1] = nil
	t.tag = t.tag[:len(t.tag)-1]

}

// IsTagIRI determines whether the call to GetTagIRI is safe for the specified index
func (t *OrderedCollectionPage) IsTagIRI(index int) (ok bool) {
	return t.tag[index].IRI != nil

}

// GetTagIRI returns the value safely if IsTagIRI returned true for the specified index
func (t *OrderedCollectionPage) GetTagIRI(index int) (v *url.URL) {
	return t.tag[index].IRI

}

// AppendTagIRI adds to the back of tag a *url.URL type
func (t *OrderedCollectionPage) AppendTagIRI(v *url.URL) {
	t.tag = append(t.tag, &tagOrderedCollectionPageIntermediateType{IRI: v})

}

// PrependTagIRI adds to the front of tag a *url.URL type
func (t *OrderedCollectionPage) PrependTagIRI(v *url.URL) {
	t.tag = append([]*tagOrderedCollectionPageIntermediateType{&tagOrderedCollectionPageIntermediateType{IRI: v}}, t.tag...)

}

// RemoveTagIRI deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveTagIRI(index int) {
	copy(t.tag[index:], t.tag[index+1:])
	t.tag[len(t.tag)-1] = nil
	t.tag = t.tag[:len(t.tag)-1]

}

// HasUnknownTag determines whether the call to GetUnknownTag is safe
func (t *OrderedCollectionPage) HasUnknownTag() (ok bool) {
	return t.tag != nil && t.tag[0].unknown_ != nil

}

// GetUnknownTag returns the unknown value for tag
func (t *OrderedCollectionPage) GetUnknownTag() (v interface{}) {
	return t.tag[0].unknown_

}

// SetUnknownTag sets the unknown value of tag
func (t *OrderedCollectionPage) SetUnknownTag(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &tagOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.tag = append(t.tag, tmp)

}

// TypeLen determines the number of elements able to be used for the GetType and RemoveType functions
func (t *OrderedCollectionPage) TypeLen() (l int) {
	return len(t.typeName)

}

// GetType returns the value for the specified index
func (t *OrderedCollectionPage) GetType(index int) (v interface{}) {
	return t.typeName[index]

}

// AppendType adds a value to the back of type
func (t *OrderedCollectionPage) AppendType(v interface{}) {
	t.typeName = append(t.typeName, v)

}

// PrependType adds a value to the front of type
func (t *OrderedCollectionPage) PrependType(v interface{}) {
	t.typeName = append([]interface{}{v}, t.typeName...)

}

// RemoveType deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveType(index int) {
	copy(t.typeName[index:], t.typeName[index+1:])
	t.typeName[len(t.typeName)-1] = nil
	t.typeName = t.typeName[:len(t.typeName)-1]

}

// IsUpdated determines whether the call to GetUpdated is safe
func (t *OrderedCollectionPage) IsUpdated() (ok bool) {
	return t.updated != nil && t.updated.dateTime != nil

}

// GetUpdated returns the value safely if IsUpdated returned true
func (t *OrderedCollectionPage) GetUpdated() (v time.Time) {
	return *t.updated.dateTime

}

// SetUpdated sets the value of updated to be of time.Time type
func (t *OrderedCollectionPage) SetUpdated(v time.Time) {
	t.updated = &updatedOrderedCollectionPageIntermediateType{dateTime: &v}

}

// IsUpdatedIRI determines whether the call to GetUpdatedIRI is safe
func (t *OrderedCollectionPage) IsUpdatedIRI() (ok bool) {
	return t.updated != nil && t.updated.IRI != nil

}

// GetUpdatedIRI returns the value safely if IsUpdatedIRI returned true
func (t *OrderedCollectionPage) GetUpdatedIRI() (v *url.URL) {
	return t.updated.IRI

}

// SetUpdatedIRI sets the value of updated to be of *url.URL type
func (t *OrderedCollectionPage) SetUpdatedIRI(v *url.URL) {
	t.updated = &updatedOrderedCollectionPageIntermediateType{IRI: v}

}

// HasUnknownUpdated determines whether the call to GetUnknownUpdated is safe
func (t *OrderedCollectionPage) HasUnknownUpdated() (ok bool) {
	return t.updated != nil && t.updated.unknown_ != nil

}

// GetUnknownUpdated returns the unknown value for updated
func (t *OrderedCollectionPage) GetUnknownUpdated() (v interface{}) {
	return t.updated.unknown_

}

// SetUnknownUpdated sets the unknown value of updated
func (t *OrderedCollectionPage) SetUnknownUpdated(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &updatedOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.updated = tmp

}

// UrlLen determines the number of elements able to be used for the IsUrlAnyURI, GetUrlAnyURI, and RemoveUrlAnyURI functions
func (t *OrderedCollectionPage) UrlLen() (l int) {
	return len(t.url)

}

// IsUrlAnyURI determines whether the call to GetUrlAnyURI is safe for the specified index
func (t *OrderedCollectionPage) IsUrlAnyURI(index int) (ok bool) {
	return t.url[index].anyURI != nil

}

// GetUrlAnyURI returns the value safely if IsUrlAnyURI returned true for the specified index
func (t *OrderedCollectionPage) GetUrlAnyURI(index int) (v *url.URL) {
	return t.url[index].anyURI

}

// AppendUrlAnyURI adds to the back of url a *url.URL type
func (t *OrderedCollectionPage) AppendUrlAnyURI(v *url.URL) {
	t.url = append(t.url, &urlOrderedCollectionPageIntermediateType{anyURI: v})

}

// PrependUrlAnyURI adds to the front of url a *url.URL type
func (t *OrderedCollectionPage) PrependUrlAnyURI(v *url.URL) {
	t.url = append([]*urlOrderedCollectionPageIntermediateType{&urlOrderedCollectionPageIntermediateType{anyURI: v}}, t.url...)

}

// RemoveUrlAnyURI deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveUrlAnyURI(index int) {
	copy(t.url[index:], t.url[index+1:])
	t.url[len(t.url)-1] = nil
	t.url = t.url[:len(t.url)-1]

}

// IsUrlLink determines whether the call to GetUrlLink is safe for the specified index
func (t *OrderedCollectionPage) IsUrlLink(index int) (ok bool) {
	return t.url[index].Link != nil

}

// GetUrlLink returns the value safely if IsUrlLink returned true for the specified index
func (t *OrderedCollectionPage) GetUrlLink(index int) (v LinkType) {
	return t.url[index].Link

}

// AppendUrlLink adds to the back of url a LinkType type
func (t *OrderedCollectionPage) AppendUrlLink(v LinkType) {
	t.url = append(t.url, &urlOrderedCollectionPageIntermediateType{Link: v})

}

// PrependUrlLink adds to the front of url a LinkType type
func (t *OrderedCollectionPage) PrependUrlLink(v LinkType) {
	t.url = append([]*urlOrderedCollectionPageIntermediateType{&urlOrderedCollectionPageIntermediateType{Link: v}}, t.url...)

}

// RemoveUrlLink deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveUrlLink(index int) {
	copy(t.url[index:], t.url[index+1:])
	t.url[len(t.url)-1] = nil
	t.url = t.url[:len(t.url)-1]

}

// HasUnknownUrl determines whether the call to GetUnknownUrl is safe
func (t *OrderedCollectionPage) HasUnknownUrl() (ok bool) {
	return t.url != nil && t.url[0].unknown_ != nil

}

// GetUnknownUrl returns the unknown value for url
func (t *OrderedCollectionPage) GetUnknownUrl() (v interface{}) {
	return t.url[0].unknown_

}

// SetUnknownUrl sets the unknown value of url
func (t *OrderedCollectionPage) SetUnknownUrl(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &urlOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.url = append(t.url, tmp)

}

// ToLen determines the number of elements able to be used for the IsToObject, GetToObject, and RemoveToObject functions
func (t *OrderedCollectionPage) ToLen() (l int) {
	return len(t.to)

}

// IsToObject determines whether the call to GetToObject is safe for the specified index
func (t *OrderedCollectionPage) IsToObject(index int) (ok bool) {
	return t.to[index].Object != nil

}

// GetToObject returns the value safely if IsToObject returned true for the specified index
func (t *OrderedCollectionPage) GetToObject(index int) (v ObjectType) {
	return t.to[index].Object

}

// AppendToObject adds to the back of to a ObjectType type
func (t *OrderedCollectionPage) AppendToObject(v ObjectType) {
	t.to = append(t.to, &toOrderedCollectionPageIntermediateType{Object: v})

}

// PrependToObject adds to the front of to a ObjectType type
func (t *OrderedCollectionPage) PrependToObject(v ObjectType) {
	t.to = append([]*toOrderedCollectionPageIntermediateType{&toOrderedCollectionPageIntermediateType{Object: v}}, t.to...)

}

// RemoveToObject deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveToObject(index int) {
	copy(t.to[index:], t.to[index+1:])
	t.to[len(t.to)-1] = nil
	t.to = t.to[:len(t.to)-1]

}

// IsToLink determines whether the call to GetToLink is safe for the specified index
func (t *OrderedCollectionPage) IsToLink(index int) (ok bool) {
	return t.to[index].Link != nil

}

// GetToLink returns the value safely if IsToLink returned true for the specified index
func (t *OrderedCollectionPage) GetToLink(index int) (v LinkType) {
	return t.to[index].Link

}

// AppendToLink adds to the back of to a LinkType type
func (t *OrderedCollectionPage) AppendToLink(v LinkType) {
	t.to = append(t.to, &toOrderedCollectionPageIntermediateType{Link: v})

}

// PrependToLink adds to the front of to a LinkType type
func (t *OrderedCollectionPage) PrependToLink(v LinkType) {
	t.to = append([]*toOrderedCollectionPageIntermediateType{&toOrderedCollectionPageIntermediateType{Link: v}}, t.to...)

}

// RemoveToLink deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveToLink(index int) {
	copy(t.to[index:], t.to[index+1:])
	t.to[len(t.to)-1] = nil
	t.to = t.to[:len(t.to)-1]

}

// IsToIRI determines whether the call to GetToIRI is safe for the specified index
func (t *OrderedCollectionPage) IsToIRI(index int) (ok bool) {
	return t.to[index].IRI != nil

}

// GetToIRI returns the value safely if IsToIRI returned true for the specified index
func (t *OrderedCollectionPage) GetToIRI(index int) (v *url.URL) {
	return t.to[index].IRI

}

// AppendToIRI adds to the back of to a *url.URL type
func (t *OrderedCollectionPage) AppendToIRI(v *url.URL) {
	t.to = append(t.to, &toOrderedCollectionPageIntermediateType{IRI: v})

}

// PrependToIRI adds to the front of to a *url.URL type
func (t *OrderedCollectionPage) PrependToIRI(v *url.URL) {
	t.to = append([]*toOrderedCollectionPageIntermediateType{&toOrderedCollectionPageIntermediateType{IRI: v}}, t.to...)

}

// RemoveToIRI deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveToIRI(index int) {
	copy(t.to[index:], t.to[index+1:])
	t.to[len(t.to)-1] = nil
	t.to = t.to[:len(t.to)-1]

}

// HasUnknownTo determines whether the call to GetUnknownTo is safe
func (t *OrderedCollectionPage) HasUnknownTo() (ok bool) {
	return t.to != nil && t.to[0].unknown_ != nil

}

// GetUnknownTo returns the unknown value for to
func (t *OrderedCollectionPage) GetUnknownTo() (v interface{}) {
	return t.to[0].unknown_

}

// SetUnknownTo sets the unknown value of to
func (t *OrderedCollectionPage) SetUnknownTo(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &toOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.to = append(t.to, tmp)

}

// BtoLen determines the number of elements able to be used for the IsBtoObject, GetBtoObject, and RemoveBtoObject functions
func (t *OrderedCollectionPage) BtoLen() (l int) {
	return len(t.bto)

}

// IsBtoObject determines whether the call to GetBtoObject is safe for the specified index
func (t *OrderedCollectionPage) IsBtoObject(index int) (ok bool) {
	return t.bto[index].Object != nil

}

// GetBtoObject returns the value safely if IsBtoObject returned true for the specified index
func (t *OrderedCollectionPage) GetBtoObject(index int) (v ObjectType) {
	return t.bto[index].Object

}

// AppendBtoObject adds to the back of bto a ObjectType type
func (t *OrderedCollectionPage) AppendBtoObject(v ObjectType) {
	t.bto = append(t.bto, &btoOrderedCollectionPageIntermediateType{Object: v})

}

// PrependBtoObject adds to the front of bto a ObjectType type
func (t *OrderedCollectionPage) PrependBtoObject(v ObjectType) {
	t.bto = append([]*btoOrderedCollectionPageIntermediateType{&btoOrderedCollectionPageIntermediateType{Object: v}}, t.bto...)

}

// RemoveBtoObject deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveBtoObject(index int) {
	copy(t.bto[index:], t.bto[index+1:])
	t.bto[len(t.bto)-1] = nil
	t.bto = t.bto[:len(t.bto)-1]

}

// IsBtoLink determines whether the call to GetBtoLink is safe for the specified index
func (t *OrderedCollectionPage) IsBtoLink(index int) (ok bool) {
	return t.bto[index].Link != nil

}

// GetBtoLink returns the value safely if IsBtoLink returned true for the specified index
func (t *OrderedCollectionPage) GetBtoLink(index int) (v LinkType) {
	return t.bto[index].Link

}

// AppendBtoLink adds to the back of bto a LinkType type
func (t *OrderedCollectionPage) AppendBtoLink(v LinkType) {
	t.bto = append(t.bto, &btoOrderedCollectionPageIntermediateType{Link: v})

}

// PrependBtoLink adds to the front of bto a LinkType type
func (t *OrderedCollectionPage) PrependBtoLink(v LinkType) {
	t.bto = append([]*btoOrderedCollectionPageIntermediateType{&btoOrderedCollectionPageIntermediateType{Link: v}}, t.bto...)

}

// RemoveBtoLink deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveBtoLink(index int) {
	copy(t.bto[index:], t.bto[index+1:])
	t.bto[len(t.bto)-1] = nil
	t.bto = t.bto[:len(t.bto)-1]

}

// IsBtoIRI determines whether the call to GetBtoIRI is safe for the specified index
func (t *OrderedCollectionPage) IsBtoIRI(index int) (ok bool) {
	return t.bto[index].IRI != nil

}

// GetBtoIRI returns the value safely if IsBtoIRI returned true for the specified index
func (t *OrderedCollectionPage) GetBtoIRI(index int) (v *url.URL) {
	return t.bto[index].IRI

}

// AppendBtoIRI adds to the back of bto a *url.URL type
func (t *OrderedCollectionPage) AppendBtoIRI(v *url.URL) {
	t.bto = append(t.bto, &btoOrderedCollectionPageIntermediateType{IRI: v})

}

// PrependBtoIRI adds to the front of bto a *url.URL type
func (t *OrderedCollectionPage) PrependBtoIRI(v *url.URL) {
	t.bto = append([]*btoOrderedCollectionPageIntermediateType{&btoOrderedCollectionPageIntermediateType{IRI: v}}, t.bto...)

}

// RemoveBtoIRI deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveBtoIRI(index int) {
	copy(t.bto[index:], t.bto[index+1:])
	t.bto[len(t.bto)-1] = nil
	t.bto = t.bto[:len(t.bto)-1]

}

// HasUnknownBto determines whether the call to GetUnknownBto is safe
func (t *OrderedCollectionPage) HasUnknownBto() (ok bool) {
	return t.bto != nil && t.bto[0].unknown_ != nil

}

// GetUnknownBto returns the unknown value for bto
func (t *OrderedCollectionPage) GetUnknownBto() (v interface{}) {
	return t.bto[0].unknown_

}

// SetUnknownBto sets the unknown value of bto
func (t *OrderedCollectionPage) SetUnknownBto(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &btoOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.bto = append(t.bto, tmp)

}

// CcLen determines the number of elements able to be used for the IsCcObject, GetCcObject, and RemoveCcObject functions
func (t *OrderedCollectionPage) CcLen() (l int) {
	return len(t.cc)

}

// IsCcObject determines whether the call to GetCcObject is safe for the specified index
func (t *OrderedCollectionPage) IsCcObject(index int) (ok bool) {
	return t.cc[index].Object != nil

}

// GetCcObject returns the value safely if IsCcObject returned true for the specified index
func (t *OrderedCollectionPage) GetCcObject(index int) (v ObjectType) {
	return t.cc[index].Object

}

// AppendCcObject adds to the back of cc a ObjectType type
func (t *OrderedCollectionPage) AppendCcObject(v ObjectType) {
	t.cc = append(t.cc, &ccOrderedCollectionPageIntermediateType{Object: v})

}

// PrependCcObject adds to the front of cc a ObjectType type
func (t *OrderedCollectionPage) PrependCcObject(v ObjectType) {
	t.cc = append([]*ccOrderedCollectionPageIntermediateType{&ccOrderedCollectionPageIntermediateType{Object: v}}, t.cc...)

}

// RemoveCcObject deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveCcObject(index int) {
	copy(t.cc[index:], t.cc[index+1:])
	t.cc[len(t.cc)-1] = nil
	t.cc = t.cc[:len(t.cc)-1]

}

// IsCcLink determines whether the call to GetCcLink is safe for the specified index
func (t *OrderedCollectionPage) IsCcLink(index int) (ok bool) {
	return t.cc[index].Link != nil

}

// GetCcLink returns the value safely if IsCcLink returned true for the specified index
func (t *OrderedCollectionPage) GetCcLink(index int) (v LinkType) {
	return t.cc[index].Link

}

// AppendCcLink adds to the back of cc a LinkType type
func (t *OrderedCollectionPage) AppendCcLink(v LinkType) {
	t.cc = append(t.cc, &ccOrderedCollectionPageIntermediateType{Link: v})

}

// PrependCcLink adds to the front of cc a LinkType type
func (t *OrderedCollectionPage) PrependCcLink(v LinkType) {
	t.cc = append([]*ccOrderedCollectionPageIntermediateType{&ccOrderedCollectionPageIntermediateType{Link: v}}, t.cc...)

}

// RemoveCcLink deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveCcLink(index int) {
	copy(t.cc[index:], t.cc[index+1:])
	t.cc[len(t.cc)-1] = nil
	t.cc = t.cc[:len(t.cc)-1]

}

// IsCcIRI determines whether the call to GetCcIRI is safe for the specified index
func (t *OrderedCollectionPage) IsCcIRI(index int) (ok bool) {
	return t.cc[index].IRI != nil

}

// GetCcIRI returns the value safely if IsCcIRI returned true for the specified index
func (t *OrderedCollectionPage) GetCcIRI(index int) (v *url.URL) {
	return t.cc[index].IRI

}

// AppendCcIRI adds to the back of cc a *url.URL type
func (t *OrderedCollectionPage) AppendCcIRI(v *url.URL) {
	t.cc = append(t.cc, &ccOrderedCollectionPageIntermediateType{IRI: v})

}

// PrependCcIRI adds to the front of cc a *url.URL type
func (t *OrderedCollectionPage) PrependCcIRI(v *url.URL) {
	t.cc = append([]*ccOrderedCollectionPageIntermediateType{&ccOrderedCollectionPageIntermediateType{IRI: v}}, t.cc...)

}

// RemoveCcIRI deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveCcIRI(index int) {
	copy(t.cc[index:], t.cc[index+1:])
	t.cc[len(t.cc)-1] = nil
	t.cc = t.cc[:len(t.cc)-1]

}

// HasUnknownCc determines whether the call to GetUnknownCc is safe
func (t *OrderedCollectionPage) HasUnknownCc() (ok bool) {
	return t.cc != nil && t.cc[0].unknown_ != nil

}

// GetUnknownCc returns the unknown value for cc
func (t *OrderedCollectionPage) GetUnknownCc() (v interface{}) {
	return t.cc[0].unknown_

}

// SetUnknownCc sets the unknown value of cc
func (t *OrderedCollectionPage) SetUnknownCc(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &ccOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.cc = append(t.cc, tmp)

}

// BccLen determines the number of elements able to be used for the IsBccObject, GetBccObject, and RemoveBccObject functions
func (t *OrderedCollectionPage) BccLen() (l int) {
	return len(t.bcc)

}

// IsBccObject determines whether the call to GetBccObject is safe for the specified index
func (t *OrderedCollectionPage) IsBccObject(index int) (ok bool) {
	return t.bcc[index].Object != nil

}

// GetBccObject returns the value safely if IsBccObject returned true for the specified index
func (t *OrderedCollectionPage) GetBccObject(index int) (v ObjectType) {
	return t.bcc[index].Object

}

// AppendBccObject adds to the back of bcc a ObjectType type
func (t *OrderedCollectionPage) AppendBccObject(v ObjectType) {
	t.bcc = append(t.bcc, &bccOrderedCollectionPageIntermediateType{Object: v})

}

// PrependBccObject adds to the front of bcc a ObjectType type
func (t *OrderedCollectionPage) PrependBccObject(v ObjectType) {
	t.bcc = append([]*bccOrderedCollectionPageIntermediateType{&bccOrderedCollectionPageIntermediateType{Object: v}}, t.bcc...)

}

// RemoveBccObject deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveBccObject(index int) {
	copy(t.bcc[index:], t.bcc[index+1:])
	t.bcc[len(t.bcc)-1] = nil
	t.bcc = t.bcc[:len(t.bcc)-1]

}

// IsBccLink determines whether the call to GetBccLink is safe for the specified index
func (t *OrderedCollectionPage) IsBccLink(index int) (ok bool) {
	return t.bcc[index].Link != nil

}

// GetBccLink returns the value safely if IsBccLink returned true for the specified index
func (t *OrderedCollectionPage) GetBccLink(index int) (v LinkType) {
	return t.bcc[index].Link

}

// AppendBccLink adds to the back of bcc a LinkType type
func (t *OrderedCollectionPage) AppendBccLink(v LinkType) {
	t.bcc = append(t.bcc, &bccOrderedCollectionPageIntermediateType{Link: v})

}

// PrependBccLink adds to the front of bcc a LinkType type
func (t *OrderedCollectionPage) PrependBccLink(v LinkType) {
	t.bcc = append([]*bccOrderedCollectionPageIntermediateType{&bccOrderedCollectionPageIntermediateType{Link: v}}, t.bcc...)

}

// RemoveBccLink deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveBccLink(index int) {
	copy(t.bcc[index:], t.bcc[index+1:])
	t.bcc[len(t.bcc)-1] = nil
	t.bcc = t.bcc[:len(t.bcc)-1]

}

// IsBccIRI determines whether the call to GetBccIRI is safe for the specified index
func (t *OrderedCollectionPage) IsBccIRI(index int) (ok bool) {
	return t.bcc[index].IRI != nil

}

// GetBccIRI returns the value safely if IsBccIRI returned true for the specified index
func (t *OrderedCollectionPage) GetBccIRI(index int) (v *url.URL) {
	return t.bcc[index].IRI

}

// AppendBccIRI adds to the back of bcc a *url.URL type
func (t *OrderedCollectionPage) AppendBccIRI(v *url.URL) {
	t.bcc = append(t.bcc, &bccOrderedCollectionPageIntermediateType{IRI: v})

}

// PrependBccIRI adds to the front of bcc a *url.URL type
func (t *OrderedCollectionPage) PrependBccIRI(v *url.URL) {
	t.bcc = append([]*bccOrderedCollectionPageIntermediateType{&bccOrderedCollectionPageIntermediateType{IRI: v}}, t.bcc...)

}

// RemoveBccIRI deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveBccIRI(index int) {
	copy(t.bcc[index:], t.bcc[index+1:])
	t.bcc[len(t.bcc)-1] = nil
	t.bcc = t.bcc[:len(t.bcc)-1]

}

// HasUnknownBcc determines whether the call to GetUnknownBcc is safe
func (t *OrderedCollectionPage) HasUnknownBcc() (ok bool) {
	return t.bcc != nil && t.bcc[0].unknown_ != nil

}

// GetUnknownBcc returns the unknown value for bcc
func (t *OrderedCollectionPage) GetUnknownBcc() (v interface{}) {
	return t.bcc[0].unknown_

}

// SetUnknownBcc sets the unknown value of bcc
func (t *OrderedCollectionPage) SetUnknownBcc(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &bccOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.bcc = append(t.bcc, tmp)

}

// IsMediaType determines whether the call to GetMediaType is safe
func (t *OrderedCollectionPage) IsMediaType() (ok bool) {
	return t.mediaType != nil && t.mediaType.mimeMediaTypeValue != nil

}

// GetMediaType returns the value safely if IsMediaType returned true
func (t *OrderedCollectionPage) GetMediaType() (v string) {
	return *t.mediaType.mimeMediaTypeValue

}

// SetMediaType sets the value of mediaType to be of string type
func (t *OrderedCollectionPage) SetMediaType(v string) {
	t.mediaType = &mediaTypeOrderedCollectionPageIntermediateType{mimeMediaTypeValue: &v}

}

// IsMediaTypeIRI determines whether the call to GetMediaTypeIRI is safe
func (t *OrderedCollectionPage) IsMediaTypeIRI() (ok bool) {
	return t.mediaType != nil && t.mediaType.IRI != nil

}

// GetMediaTypeIRI returns the value safely if IsMediaTypeIRI returned true
func (t *OrderedCollectionPage) GetMediaTypeIRI() (v *url.URL) {
	return t.mediaType.IRI

}

// SetMediaTypeIRI sets the value of mediaType to be of *url.URL type
func (t *OrderedCollectionPage) SetMediaTypeIRI(v *url.URL) {
	t.mediaType = &mediaTypeOrderedCollectionPageIntermediateType{IRI: v}

}

// HasUnknownMediaType determines whether the call to GetUnknownMediaType is safe
func (t *OrderedCollectionPage) HasUnknownMediaType() (ok bool) {
	return t.mediaType != nil && t.mediaType.unknown_ != nil

}

// GetUnknownMediaType returns the unknown value for mediaType
func (t *OrderedCollectionPage) GetUnknownMediaType() (v interface{}) {
	return t.mediaType.unknown_

}

// SetUnknownMediaType sets the unknown value of mediaType
func (t *OrderedCollectionPage) SetUnknownMediaType(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &mediaTypeOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.mediaType = tmp

}

// IsDuration determines whether the call to GetDuration is safe
func (t *OrderedCollectionPage) IsDuration() (ok bool) {
	return t.duration != nil && t.duration.duration != nil

}

// GetDuration returns the value safely if IsDuration returned true
func (t *OrderedCollectionPage) GetDuration() (v time.Duration) {
	return *t.duration.duration

}

// SetDuration sets the value of duration to be of time.Duration type
func (t *OrderedCollectionPage) SetDuration(v time.Duration) {
	t.duration = &durationOrderedCollectionPageIntermediateType{duration: &v}

}

// IsDurationIRI determines whether the call to GetDurationIRI is safe
func (t *OrderedCollectionPage) IsDurationIRI() (ok bool) {
	return t.duration != nil && t.duration.IRI != nil

}

// GetDurationIRI returns the value safely if IsDurationIRI returned true
func (t *OrderedCollectionPage) GetDurationIRI() (v *url.URL) {
	return t.duration.IRI

}

// SetDurationIRI sets the value of duration to be of *url.URL type
func (t *OrderedCollectionPage) SetDurationIRI(v *url.URL) {
	t.duration = &durationOrderedCollectionPageIntermediateType{IRI: v}

}

// HasUnknownDuration determines whether the call to GetUnknownDuration is safe
func (t *OrderedCollectionPage) HasUnknownDuration() (ok bool) {
	return t.duration != nil && t.duration.unknown_ != nil

}

// GetUnknownDuration returns the unknown value for duration
func (t *OrderedCollectionPage) GetUnknownDuration() (v interface{}) {
	return t.duration.unknown_

}

// SetUnknownDuration sets the unknown value of duration
func (t *OrderedCollectionPage) SetUnknownDuration(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &durationOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.duration = tmp

}

// IsSource determines whether the call to GetSource is safe
func (t *OrderedCollectionPage) IsSource() (ok bool) {
	return t.source != nil && t.source.Object != nil

}

// GetSource returns the value safely if IsSource returned true
func (t *OrderedCollectionPage) GetSource() (v ObjectType) {
	return t.source.Object

}

// SetSource sets the value of source to be of ObjectType type
func (t *OrderedCollectionPage) SetSource(v ObjectType) {
	t.source = &sourceOrderedCollectionPageIntermediateType{Object: v}

}

// IsSourceIRI determines whether the call to GetSourceIRI is safe
func (t *OrderedCollectionPage) IsSourceIRI() (ok bool) {
	return t.source != nil && t.source.IRI != nil

}

// GetSourceIRI returns the value safely if IsSourceIRI returned true
func (t *OrderedCollectionPage) GetSourceIRI() (v *url.URL) {
	return t.source.IRI

}

// SetSourceIRI sets the value of source to be of *url.URL type
func (t *OrderedCollectionPage) SetSourceIRI(v *url.URL) {
	t.source = &sourceOrderedCollectionPageIntermediateType{IRI: v}

}

// HasUnknownSource determines whether the call to GetUnknownSource is safe
func (t *OrderedCollectionPage) HasUnknownSource() (ok bool) {
	return t.source != nil && t.source.unknown_ != nil

}

// GetUnknownSource returns the unknown value for source
func (t *OrderedCollectionPage) GetUnknownSource() (v interface{}) {
	return t.source.unknown_

}

// SetUnknownSource sets the unknown value of source
func (t *OrderedCollectionPage) SetUnknownSource(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &sourceOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.source = tmp

}

// IsInboxOrderedCollection determines whether the call to GetInboxOrderedCollection is safe
func (t *OrderedCollectionPage) IsInboxOrderedCollection() (ok bool) {
	return t.inbox != nil && t.inbox.OrderedCollection != nil

}

// GetInboxOrderedCollection returns the value safely if IsInboxOrderedCollection returned true
func (t *OrderedCollectionPage) GetInboxOrderedCollection() (v OrderedCollectionType) {
	return t.inbox.OrderedCollection

}

// SetInboxOrderedCollection sets the value of inbox to be of OrderedCollectionType type
func (t *OrderedCollectionPage) SetInboxOrderedCollection(v OrderedCollectionType) {
	t.inbox = &inboxOrderedCollectionPageIntermediateType{OrderedCollection: v}

}

// IsInboxAnyURI determines whether the call to GetInboxAnyURI is safe
func (t *OrderedCollectionPage) IsInboxAnyURI() (ok bool) {
	return t.inbox != nil && t.inbox.anyURI != nil

}

// GetInboxAnyURI returns the value safely if IsInboxAnyURI returned true
func (t *OrderedCollectionPage) GetInboxAnyURI() (v *url.URL) {
	return t.inbox.anyURI

}

// SetInboxAnyURI sets the value of inbox to be of *url.URL type
func (t *OrderedCollectionPage) SetInboxAnyURI(v *url.URL) {
	t.inbox = &inboxOrderedCollectionPageIntermediateType{anyURI: v}

}

// HasUnknownInbox determines whether the call to GetUnknownInbox is safe
func (t *OrderedCollectionPage) HasUnknownInbox() (ok bool) {
	return t.inbox != nil && t.inbox.unknown_ != nil

}

// GetUnknownInbox returns the unknown value for inbox
func (t *OrderedCollectionPage) GetUnknownInbox() (v interface{}) {
	return t.inbox.unknown_

}

// SetUnknownInbox sets the unknown value of inbox
func (t *OrderedCollectionPage) SetUnknownInbox(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &inboxOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.inbox = tmp

}

// IsOutboxOrderedCollection determines whether the call to GetOutboxOrderedCollection is safe
func (t *OrderedCollectionPage) IsOutboxOrderedCollection() (ok bool) {
	return t.outbox != nil && t.outbox.OrderedCollection != nil

}

// GetOutboxOrderedCollection returns the value safely if IsOutboxOrderedCollection returned true
func (t *OrderedCollectionPage) GetOutboxOrderedCollection() (v OrderedCollectionType) {
	return t.outbox.OrderedCollection

}

// SetOutboxOrderedCollection sets the value of outbox to be of OrderedCollectionType type
func (t *OrderedCollectionPage) SetOutboxOrderedCollection(v OrderedCollectionType) {
	t.outbox = &outboxOrderedCollectionPageIntermediateType{OrderedCollection: v}

}

// IsOutboxAnyURI determines whether the call to GetOutboxAnyURI is safe
func (t *OrderedCollectionPage) IsOutboxAnyURI() (ok bool) {
	return t.outbox != nil && t.outbox.anyURI != nil

}

// GetOutboxAnyURI returns the value safely if IsOutboxAnyURI returned true
func (t *OrderedCollectionPage) GetOutboxAnyURI() (v *url.URL) {
	return t.outbox.anyURI

}

// SetOutboxAnyURI sets the value of outbox to be of *url.URL type
func (t *OrderedCollectionPage) SetOutboxAnyURI(v *url.URL) {
	t.outbox = &outboxOrderedCollectionPageIntermediateType{anyURI: v}

}

// HasUnknownOutbox determines whether the call to GetUnknownOutbox is safe
func (t *OrderedCollectionPage) HasUnknownOutbox() (ok bool) {
	return t.outbox != nil && t.outbox.unknown_ != nil

}

// GetUnknownOutbox returns the unknown value for outbox
func (t *OrderedCollectionPage) GetUnknownOutbox() (v interface{}) {
	return t.outbox.unknown_

}

// SetUnknownOutbox sets the unknown value of outbox
func (t *OrderedCollectionPage) SetUnknownOutbox(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &outboxOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.outbox = tmp

}

// IsFollowingCollection determines whether the call to GetFollowingCollection is safe
func (t *OrderedCollectionPage) IsFollowingCollection() (ok bool) {
	return t.following != nil && t.following.Collection != nil

}

// GetFollowingCollection returns the value safely if IsFollowingCollection returned true
func (t *OrderedCollectionPage) GetFollowingCollection() (v CollectionType) {
	return t.following.Collection

}

// SetFollowingCollection sets the value of following to be of CollectionType type
func (t *OrderedCollectionPage) SetFollowingCollection(v CollectionType) {
	t.following = &followingOrderedCollectionPageIntermediateType{Collection: v}

}

// IsFollowingOrderedCollection determines whether the call to GetFollowingOrderedCollection is safe
func (t *OrderedCollectionPage) IsFollowingOrderedCollection() (ok bool) {
	return t.following != nil && t.following.OrderedCollection != nil

}

// GetFollowingOrderedCollection returns the value safely if IsFollowingOrderedCollection returned true
func (t *OrderedCollectionPage) GetFollowingOrderedCollection() (v OrderedCollectionType) {
	return t.following.OrderedCollection

}

// SetFollowingOrderedCollection sets the value of following to be of OrderedCollectionType type
func (t *OrderedCollectionPage) SetFollowingOrderedCollection(v OrderedCollectionType) {
	t.following = &followingOrderedCollectionPageIntermediateType{OrderedCollection: v}

}

// IsFollowingAnyURI determines whether the call to GetFollowingAnyURI is safe
func (t *OrderedCollectionPage) IsFollowingAnyURI() (ok bool) {
	return t.following != nil && t.following.anyURI != nil

}

// GetFollowingAnyURI returns the value safely if IsFollowingAnyURI returned true
func (t *OrderedCollectionPage) GetFollowingAnyURI() (v *url.URL) {
	return t.following.anyURI

}

// SetFollowingAnyURI sets the value of following to be of *url.URL type
func (t *OrderedCollectionPage) SetFollowingAnyURI(v *url.URL) {
	t.following = &followingOrderedCollectionPageIntermediateType{anyURI: v}

}

// HasUnknownFollowing determines whether the call to GetUnknownFollowing is safe
func (t *OrderedCollectionPage) HasUnknownFollowing() (ok bool) {
	return t.following != nil && t.following.unknown_ != nil

}

// GetUnknownFollowing returns the unknown value for following
func (t *OrderedCollectionPage) GetUnknownFollowing() (v interface{}) {
	return t.following.unknown_

}

// SetUnknownFollowing sets the unknown value of following
func (t *OrderedCollectionPage) SetUnknownFollowing(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &followingOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.following = tmp

}

// IsFollowersCollection determines whether the call to GetFollowersCollection is safe
func (t *OrderedCollectionPage) IsFollowersCollection() (ok bool) {
	return t.followers != nil && t.followers.Collection != nil

}

// GetFollowersCollection returns the value safely if IsFollowersCollection returned true
func (t *OrderedCollectionPage) GetFollowersCollection() (v CollectionType) {
	return t.followers.Collection

}

// SetFollowersCollection sets the value of followers to be of CollectionType type
func (t *OrderedCollectionPage) SetFollowersCollection(v CollectionType) {
	t.followers = &followersOrderedCollectionPageIntermediateType{Collection: v}

}

// IsFollowersOrderedCollection determines whether the call to GetFollowersOrderedCollection is safe
func (t *OrderedCollectionPage) IsFollowersOrderedCollection() (ok bool) {
	return t.followers != nil && t.followers.OrderedCollection != nil

}

// GetFollowersOrderedCollection returns the value safely if IsFollowersOrderedCollection returned true
func (t *OrderedCollectionPage) GetFollowersOrderedCollection() (v OrderedCollectionType) {
	return t.followers.OrderedCollection

}

// SetFollowersOrderedCollection sets the value of followers to be of OrderedCollectionType type
func (t *OrderedCollectionPage) SetFollowersOrderedCollection(v OrderedCollectionType) {
	t.followers = &followersOrderedCollectionPageIntermediateType{OrderedCollection: v}

}

// IsFollowersAnyURI determines whether the call to GetFollowersAnyURI is safe
func (t *OrderedCollectionPage) IsFollowersAnyURI() (ok bool) {
	return t.followers != nil && t.followers.anyURI != nil

}

// GetFollowersAnyURI returns the value safely if IsFollowersAnyURI returned true
func (t *OrderedCollectionPage) GetFollowersAnyURI() (v *url.URL) {
	return t.followers.anyURI

}

// SetFollowersAnyURI sets the value of followers to be of *url.URL type
func (t *OrderedCollectionPage) SetFollowersAnyURI(v *url.URL) {
	t.followers = &followersOrderedCollectionPageIntermediateType{anyURI: v}

}

// HasUnknownFollowers determines whether the call to GetUnknownFollowers is safe
func (t *OrderedCollectionPage) HasUnknownFollowers() (ok bool) {
	return t.followers != nil && t.followers.unknown_ != nil

}

// GetUnknownFollowers returns the unknown value for followers
func (t *OrderedCollectionPage) GetUnknownFollowers() (v interface{}) {
	return t.followers.unknown_

}

// SetUnknownFollowers sets the unknown value of followers
func (t *OrderedCollectionPage) SetUnknownFollowers(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &followersOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.followers = tmp

}

// IsLikedCollection determines whether the call to GetLikedCollection is safe
func (t *OrderedCollectionPage) IsLikedCollection() (ok bool) {
	return t.liked != nil && t.liked.Collection != nil

}

// GetLikedCollection returns the value safely if IsLikedCollection returned true
func (t *OrderedCollectionPage) GetLikedCollection() (v CollectionType) {
	return t.liked.Collection

}

// SetLikedCollection sets the value of liked to be of CollectionType type
func (t *OrderedCollectionPage) SetLikedCollection(v CollectionType) {
	t.liked = &likedOrderedCollectionPageIntermediateType{Collection: v}

}

// IsLikedOrderedCollection determines whether the call to GetLikedOrderedCollection is safe
func (t *OrderedCollectionPage) IsLikedOrderedCollection() (ok bool) {
	return t.liked != nil && t.liked.OrderedCollection != nil

}

// GetLikedOrderedCollection returns the value safely if IsLikedOrderedCollection returned true
func (t *OrderedCollectionPage) GetLikedOrderedCollection() (v OrderedCollectionType) {
	return t.liked.OrderedCollection

}

// SetLikedOrderedCollection sets the value of liked to be of OrderedCollectionType type
func (t *OrderedCollectionPage) SetLikedOrderedCollection(v OrderedCollectionType) {
	t.liked = &likedOrderedCollectionPageIntermediateType{OrderedCollection: v}

}

// IsLikedAnyURI determines whether the call to GetLikedAnyURI is safe
func (t *OrderedCollectionPage) IsLikedAnyURI() (ok bool) {
	return t.liked != nil && t.liked.anyURI != nil

}

// GetLikedAnyURI returns the value safely if IsLikedAnyURI returned true
func (t *OrderedCollectionPage) GetLikedAnyURI() (v *url.URL) {
	return t.liked.anyURI

}

// SetLikedAnyURI sets the value of liked to be of *url.URL type
func (t *OrderedCollectionPage) SetLikedAnyURI(v *url.URL) {
	t.liked = &likedOrderedCollectionPageIntermediateType{anyURI: v}

}

// HasUnknownLiked determines whether the call to GetUnknownLiked is safe
func (t *OrderedCollectionPage) HasUnknownLiked() (ok bool) {
	return t.liked != nil && t.liked.unknown_ != nil

}

// GetUnknownLiked returns the unknown value for liked
func (t *OrderedCollectionPage) GetUnknownLiked() (v interface{}) {
	return t.liked.unknown_

}

// SetUnknownLiked sets the unknown value of liked
func (t *OrderedCollectionPage) SetUnknownLiked(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &likedOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.liked = tmp

}

// IsLikesCollection determines whether the call to GetLikesCollection is safe
func (t *OrderedCollectionPage) IsLikesCollection() (ok bool) {
	return t.likes != nil && t.likes.Collection != nil

}

// GetLikesCollection returns the value safely if IsLikesCollection returned true
func (t *OrderedCollectionPage) GetLikesCollection() (v CollectionType) {
	return t.likes.Collection

}

// SetLikesCollection sets the value of likes to be of CollectionType type
func (t *OrderedCollectionPage) SetLikesCollection(v CollectionType) {
	t.likes = &likesOrderedCollectionPageIntermediateType{Collection: v}

}

// IsLikesOrderedCollection determines whether the call to GetLikesOrderedCollection is safe
func (t *OrderedCollectionPage) IsLikesOrderedCollection() (ok bool) {
	return t.likes != nil && t.likes.OrderedCollection != nil

}

// GetLikesOrderedCollection returns the value safely if IsLikesOrderedCollection returned true
func (t *OrderedCollectionPage) GetLikesOrderedCollection() (v OrderedCollectionType) {
	return t.likes.OrderedCollection

}

// SetLikesOrderedCollection sets the value of likes to be of OrderedCollectionType type
func (t *OrderedCollectionPage) SetLikesOrderedCollection(v OrderedCollectionType) {
	t.likes = &likesOrderedCollectionPageIntermediateType{OrderedCollection: v}

}

// IsLikesAnyURI determines whether the call to GetLikesAnyURI is safe
func (t *OrderedCollectionPage) IsLikesAnyURI() (ok bool) {
	return t.likes != nil && t.likes.anyURI != nil

}

// GetLikesAnyURI returns the value safely if IsLikesAnyURI returned true
func (t *OrderedCollectionPage) GetLikesAnyURI() (v *url.URL) {
	return t.likes.anyURI

}

// SetLikesAnyURI sets the value of likes to be of *url.URL type
func (t *OrderedCollectionPage) SetLikesAnyURI(v *url.URL) {
	t.likes = &likesOrderedCollectionPageIntermediateType{anyURI: v}

}

// HasUnknownLikes determines whether the call to GetUnknownLikes is safe
func (t *OrderedCollectionPage) HasUnknownLikes() (ok bool) {
	return t.likes != nil && t.likes.unknown_ != nil

}

// GetUnknownLikes returns the unknown value for likes
func (t *OrderedCollectionPage) GetUnknownLikes() (v interface{}) {
	return t.likes.unknown_

}

// SetUnknownLikes sets the unknown value of likes
func (t *OrderedCollectionPage) SetUnknownLikes(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &likesOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.likes = tmp

}

// StreamsLen determines the number of elements able to be used for the GetStreams and RemoveStreams functions
func (t *OrderedCollectionPage) StreamsLen() (l int) {
	return len(t.streams)

}

// GetStreams returns the value for the specified index
func (t *OrderedCollectionPage) GetStreams(index int) (v *url.URL) {
	return t.streams[index]

}

// AppendStreams adds a value to the back of streams
func (t *OrderedCollectionPage) AppendStreams(v *url.URL) {
	t.streams = append(t.streams, v)

}

// PrependStreams adds a value to the front of streams
func (t *OrderedCollectionPage) PrependStreams(v *url.URL) {
	t.streams = append([]*url.URL{v}, t.streams...)

}

// RemoveStreams deletes the value from the specified index
func (t *OrderedCollectionPage) RemoveStreams(index int) {
	copy(t.streams[index:], t.streams[index+1:])
	t.streams[len(t.streams)-1] = nil
	t.streams = t.streams[:len(t.streams)-1]

}

// HasUnknownStreams determines whether the call to GetUnknownStreams is safe
func (t *OrderedCollectionPage) HasUnknownStreams() (ok bool) {
	return t.unknown_ != nil && t.unknown_["streams"] != nil

}

// GetUnknownStreams returns the unknown value for streams
func (t *OrderedCollectionPage) GetUnknownStreams() (v interface{}) {
	return t.unknown_["streams"]

}

// SetUnknownStreams sets the unknown value of streams
func (t *OrderedCollectionPage) SetUnknownStreams(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["streams"] = i

}

// IsPreferredUsername determines whether the call to GetPreferredUsername is safe
func (t *OrderedCollectionPage) IsPreferredUsername() (ok bool) {
	return t.preferredUsername != nil && t.preferredUsername.stringName != nil

}

// GetPreferredUsername returns the value safely if IsPreferredUsername returned true
func (t *OrderedCollectionPage) GetPreferredUsername() (v string) {
	return *t.preferredUsername.stringName

}

// SetPreferredUsername sets the value of preferredUsername to be of string type
func (t *OrderedCollectionPage) SetPreferredUsername(v string) {
	t.preferredUsername = &preferredUsernameOrderedCollectionPageIntermediateType{stringName: &v}

}

// IsPreferredUsernameIRI determines whether the call to GetPreferredUsernameIRI is safe
func (t *OrderedCollectionPage) IsPreferredUsernameIRI() (ok bool) {
	return t.preferredUsername != nil && t.preferredUsername.IRI != nil

}

// GetPreferredUsernameIRI returns the value safely if IsPreferredUsernameIRI returned true
func (t *OrderedCollectionPage) GetPreferredUsernameIRI() (v *url.URL) {
	return t.preferredUsername.IRI

}

// SetPreferredUsernameIRI sets the value of preferredUsername to be of *url.URL type
func (t *OrderedCollectionPage) SetPreferredUsernameIRI(v *url.URL) {
	t.preferredUsername = &preferredUsernameOrderedCollectionPageIntermediateType{IRI: v}

}

// HasUnknownPreferredUsername determines whether the call to GetUnknownPreferredUsername is safe
func (t *OrderedCollectionPage) HasUnknownPreferredUsername() (ok bool) {
	return t.preferredUsername != nil && t.preferredUsername.unknown_ != nil

}

// GetUnknownPreferredUsername returns the unknown value for preferredUsername
func (t *OrderedCollectionPage) GetUnknownPreferredUsername() (v interface{}) {
	return t.preferredUsername.unknown_

}

// SetUnknownPreferredUsername sets the unknown value of preferredUsername
func (t *OrderedCollectionPage) SetUnknownPreferredUsername(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &preferredUsernameOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.preferredUsername = tmp

}

// PreferredUsernameMapLanguages returns all languages for this property's language mapping, or nil if there are none.
func (t *OrderedCollectionPage) PreferredUsernameMapLanguages() (l []string) {
	if t.preferredUsernameMap == nil || len(t.preferredUsernameMap) == 0 {
		return nil
	}
	for k := range t.preferredUsernameMap {
		l = append(l, k)
	}
	return

}

// GetPreferredUsernameMap retrieves the value of the property for the specified language, or an empty string if it does not exist
func (t *OrderedCollectionPage) GetPreferredUsernameMap(l string) (v string) {
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
func (t *OrderedCollectionPage) SetPreferredUsernameMap(l string, v string) {
	if t.preferredUsernameMap == nil {
		t.preferredUsernameMap = make(map[string]string)
	}
	t.preferredUsernameMap[l] = v

}

// IsEndpoints determines whether the call to GetEndpoints is safe
func (t *OrderedCollectionPage) IsEndpoints() (ok bool) {
	return t.endpoints != nil && t.endpoints.Object != nil

}

// GetEndpoints returns the value safely if IsEndpoints returned true
func (t *OrderedCollectionPage) GetEndpoints() (v ObjectType) {
	return t.endpoints.Object

}

// SetEndpoints sets the value of endpoints to be of ObjectType type
func (t *OrderedCollectionPage) SetEndpoints(v ObjectType) {
	t.endpoints = &endpointsOrderedCollectionPageIntermediateType{Object: v}

}

// IsEndpointsIRI determines whether the call to GetEndpointsIRI is safe
func (t *OrderedCollectionPage) IsEndpointsIRI() (ok bool) {
	return t.endpoints != nil && t.endpoints.IRI != nil

}

// GetEndpointsIRI returns the value safely if IsEndpointsIRI returned true
func (t *OrderedCollectionPage) GetEndpointsIRI() (v *url.URL) {
	return t.endpoints.IRI

}

// SetEndpointsIRI sets the value of endpoints to be of *url.URL type
func (t *OrderedCollectionPage) SetEndpointsIRI(v *url.URL) {
	t.endpoints = &endpointsOrderedCollectionPageIntermediateType{IRI: v}

}

// HasUnknownEndpoints determines whether the call to GetUnknownEndpoints is safe
func (t *OrderedCollectionPage) HasUnknownEndpoints() (ok bool) {
	return t.endpoints != nil && t.endpoints.unknown_ != nil

}

// GetUnknownEndpoints returns the unknown value for endpoints
func (t *OrderedCollectionPage) GetUnknownEndpoints() (v interface{}) {
	return t.endpoints.unknown_

}

// SetUnknownEndpoints sets the unknown value of endpoints
func (t *OrderedCollectionPage) SetUnknownEndpoints(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &endpointsOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.endpoints = tmp

}

// HasProxyUrl determines whether the call to GetProxyUrl is safe
func (t *OrderedCollectionPage) HasProxyUrl() (ok bool) {
	return t.proxyUrl != nil

}

// GetProxyUrl returns the value for proxyUrl
func (t *OrderedCollectionPage) GetProxyUrl() (v *url.URL) {
	return t.proxyUrl

}

// SetProxyUrl sets the value of proxyUrl
func (t *OrderedCollectionPage) SetProxyUrl(v *url.URL) {
	t.proxyUrl = v

}

// HasUnknownProxyUrl determines whether the call to GetUnknownProxyUrl is safe
func (t *OrderedCollectionPage) HasUnknownProxyUrl() (ok bool) {
	return t.unknown_ != nil && t.unknown_["proxyUrl"] != nil

}

// GetUnknownProxyUrl returns the unknown value for proxyUrl
func (t *OrderedCollectionPage) GetUnknownProxyUrl() (v interface{}) {
	return t.unknown_["proxyUrl"]

}

// SetUnknownProxyUrl sets the unknown value of proxyUrl
func (t *OrderedCollectionPage) SetUnknownProxyUrl(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["proxyUrl"] = i

}

// HasOauthAuthorizationEndpoint determines whether the call to GetOauthAuthorizationEndpoint is safe
func (t *OrderedCollectionPage) HasOauthAuthorizationEndpoint() (ok bool) {
	return t.oauthAuthorizationEndpoint != nil

}

// GetOauthAuthorizationEndpoint returns the value for oauthAuthorizationEndpoint
func (t *OrderedCollectionPage) GetOauthAuthorizationEndpoint() (v *url.URL) {
	return t.oauthAuthorizationEndpoint

}

// SetOauthAuthorizationEndpoint sets the value of oauthAuthorizationEndpoint
func (t *OrderedCollectionPage) SetOauthAuthorizationEndpoint(v *url.URL) {
	t.oauthAuthorizationEndpoint = v

}

// HasUnknownOauthAuthorizationEndpoint determines whether the call to GetUnknownOauthAuthorizationEndpoint is safe
func (t *OrderedCollectionPage) HasUnknownOauthAuthorizationEndpoint() (ok bool) {
	return t.unknown_ != nil && t.unknown_["oauthAuthorizationEndpoint"] != nil

}

// GetUnknownOauthAuthorizationEndpoint returns the unknown value for oauthAuthorizationEndpoint
func (t *OrderedCollectionPage) GetUnknownOauthAuthorizationEndpoint() (v interface{}) {
	return t.unknown_["oauthAuthorizationEndpoint"]

}

// SetUnknownOauthAuthorizationEndpoint sets the unknown value of oauthAuthorizationEndpoint
func (t *OrderedCollectionPage) SetUnknownOauthAuthorizationEndpoint(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["oauthAuthorizationEndpoint"] = i

}

// HasOauthTokenEndpoint determines whether the call to GetOauthTokenEndpoint is safe
func (t *OrderedCollectionPage) HasOauthTokenEndpoint() (ok bool) {
	return t.oauthTokenEndpoint != nil

}

// GetOauthTokenEndpoint returns the value for oauthTokenEndpoint
func (t *OrderedCollectionPage) GetOauthTokenEndpoint() (v *url.URL) {
	return t.oauthTokenEndpoint

}

// SetOauthTokenEndpoint sets the value of oauthTokenEndpoint
func (t *OrderedCollectionPage) SetOauthTokenEndpoint(v *url.URL) {
	t.oauthTokenEndpoint = v

}

// HasUnknownOauthTokenEndpoint determines whether the call to GetUnknownOauthTokenEndpoint is safe
func (t *OrderedCollectionPage) HasUnknownOauthTokenEndpoint() (ok bool) {
	return t.unknown_ != nil && t.unknown_["oauthTokenEndpoint"] != nil

}

// GetUnknownOauthTokenEndpoint returns the unknown value for oauthTokenEndpoint
func (t *OrderedCollectionPage) GetUnknownOauthTokenEndpoint() (v interface{}) {
	return t.unknown_["oauthTokenEndpoint"]

}

// SetUnknownOauthTokenEndpoint sets the unknown value of oauthTokenEndpoint
func (t *OrderedCollectionPage) SetUnknownOauthTokenEndpoint(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["oauthTokenEndpoint"] = i

}

// HasProvideClientKey determines whether the call to GetProvideClientKey is safe
func (t *OrderedCollectionPage) HasProvideClientKey() (ok bool) {
	return t.provideClientKey != nil

}

// GetProvideClientKey returns the value for provideClientKey
func (t *OrderedCollectionPage) GetProvideClientKey() (v *url.URL) {
	return t.provideClientKey

}

// SetProvideClientKey sets the value of provideClientKey
func (t *OrderedCollectionPage) SetProvideClientKey(v *url.URL) {
	t.provideClientKey = v

}

// HasUnknownProvideClientKey determines whether the call to GetUnknownProvideClientKey is safe
func (t *OrderedCollectionPage) HasUnknownProvideClientKey() (ok bool) {
	return t.unknown_ != nil && t.unknown_["provideClientKey"] != nil

}

// GetUnknownProvideClientKey returns the unknown value for provideClientKey
func (t *OrderedCollectionPage) GetUnknownProvideClientKey() (v interface{}) {
	return t.unknown_["provideClientKey"]

}

// SetUnknownProvideClientKey sets the unknown value of provideClientKey
func (t *OrderedCollectionPage) SetUnknownProvideClientKey(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["provideClientKey"] = i

}

// HasSignClientKey determines whether the call to GetSignClientKey is safe
func (t *OrderedCollectionPage) HasSignClientKey() (ok bool) {
	return t.signClientKey != nil

}

// GetSignClientKey returns the value for signClientKey
func (t *OrderedCollectionPage) GetSignClientKey() (v *url.URL) {
	return t.signClientKey

}

// SetSignClientKey sets the value of signClientKey
func (t *OrderedCollectionPage) SetSignClientKey(v *url.URL) {
	t.signClientKey = v

}

// HasUnknownSignClientKey determines whether the call to GetUnknownSignClientKey is safe
func (t *OrderedCollectionPage) HasUnknownSignClientKey() (ok bool) {
	return t.unknown_ != nil && t.unknown_["signClientKey"] != nil

}

// GetUnknownSignClientKey returns the unknown value for signClientKey
func (t *OrderedCollectionPage) GetUnknownSignClientKey() (v interface{}) {
	return t.unknown_["signClientKey"]

}

// SetUnknownSignClientKey sets the unknown value of signClientKey
func (t *OrderedCollectionPage) SetUnknownSignClientKey(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["signClientKey"] = i

}

// HasSharedInbox determines whether the call to GetSharedInbox is safe
func (t *OrderedCollectionPage) HasSharedInbox() (ok bool) {
	return t.sharedInbox != nil

}

// GetSharedInbox returns the value for sharedInbox
func (t *OrderedCollectionPage) GetSharedInbox() (v *url.URL) {
	return t.sharedInbox

}

// SetSharedInbox sets the value of sharedInbox
func (t *OrderedCollectionPage) SetSharedInbox(v *url.URL) {
	t.sharedInbox = v

}

// HasUnknownSharedInbox determines whether the call to GetUnknownSharedInbox is safe
func (t *OrderedCollectionPage) HasUnknownSharedInbox() (ok bool) {
	return t.unknown_ != nil && t.unknown_["sharedInbox"] != nil

}

// GetUnknownSharedInbox returns the unknown value for sharedInbox
func (t *OrderedCollectionPage) GetUnknownSharedInbox() (v interface{}) {
	return t.unknown_["sharedInbox"]

}

// SetUnknownSharedInbox sets the unknown value of sharedInbox
func (t *OrderedCollectionPage) SetUnknownSharedInbox(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_["sharedInbox"] = i

}

// IsPartOfLink determines whether the call to GetPartOfLink is safe
func (t *OrderedCollectionPage) IsPartOfLink() (ok bool) {
	return t.partOf != nil && t.partOf.Link != nil

}

// GetPartOfLink returns the value safely if IsPartOfLink returned true
func (t *OrderedCollectionPage) GetPartOfLink() (v LinkType) {
	return t.partOf.Link

}

// SetPartOfLink sets the value of partOf to be of LinkType type
func (t *OrderedCollectionPage) SetPartOfLink(v LinkType) {
	t.partOf = &partOfOrderedCollectionPageIntermediateType{Link: v}

}

// IsPartOfCollection determines whether the call to GetPartOfCollection is safe
func (t *OrderedCollectionPage) IsPartOfCollection() (ok bool) {
	return t.partOf != nil && t.partOf.Collection != nil

}

// GetPartOfCollection returns the value safely if IsPartOfCollection returned true
func (t *OrderedCollectionPage) GetPartOfCollection() (v CollectionType) {
	return t.partOf.Collection

}

// SetPartOfCollection sets the value of partOf to be of CollectionType type
func (t *OrderedCollectionPage) SetPartOfCollection(v CollectionType) {
	t.partOf = &partOfOrderedCollectionPageIntermediateType{Collection: v}

}

// IsPartOfIRI determines whether the call to GetPartOfIRI is safe
func (t *OrderedCollectionPage) IsPartOfIRI() (ok bool) {
	return t.partOf != nil && t.partOf.IRI != nil

}

// GetPartOfIRI returns the value safely if IsPartOfIRI returned true
func (t *OrderedCollectionPage) GetPartOfIRI() (v *url.URL) {
	return t.partOf.IRI

}

// SetPartOfIRI sets the value of partOf to be of *url.URL type
func (t *OrderedCollectionPage) SetPartOfIRI(v *url.URL) {
	t.partOf = &partOfOrderedCollectionPageIntermediateType{IRI: v}

}

// HasUnknownPartOf determines whether the call to GetUnknownPartOf is safe
func (t *OrderedCollectionPage) HasUnknownPartOf() (ok bool) {
	return t.partOf != nil && t.partOf.unknown_ != nil

}

// GetUnknownPartOf returns the unknown value for partOf
func (t *OrderedCollectionPage) GetUnknownPartOf() (v interface{}) {
	return t.partOf.unknown_

}

// SetUnknownPartOf sets the unknown value of partOf
func (t *OrderedCollectionPage) SetUnknownPartOf(i interface{}) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	tmp := &partOfOrderedCollectionPageIntermediateType{}
	tmp.unknown_ = i
	t.partOf = tmp

}

// ItemsLen is NOT a valid property for this type; calling its associated methods will always yield an empty-equivalent value such as false, nil, or empty string. This includes instances where it should return itself. This ugliness is a symptom of the fundamental design of the ActivityStream vocabulary as instead of 'W-is-a-X' relationships it contains the notion of 'W-is-a-X-except-for-Y'.
func (t *OrderedCollectionPage) ItemsLen() (l int) {
	return 0

}

// IsItemsObject is NOT a valid property for this type; calling its associated methods will always yield an empty-equivalent value such as false, nil, or empty string. This includes instances where it should return itself. This ugliness is a symptom of the fundamental design of the ActivityStream vocabulary as instead of 'W-is-a-X' relationships it contains the notion of 'W-is-a-X-except-for-Y'.
func (t *OrderedCollectionPage) IsItemsObject(index int) (ok bool) {
	return false

}

// GetItemsObject is NOT a valid property for this type; calling its associated methods will always yield an empty-equivalent value such as false, nil, or empty string. This includes instances where it should return itself. This ugliness is a symptom of the fundamental design of the ActivityStream vocabulary as instead of 'W-is-a-X' relationships it contains the notion of 'W-is-a-X-except-for-Y'.
func (t *OrderedCollectionPage) GetItemsObject(index int) (v ObjectType) {
	return nil

}

// AppendItemsObject is NOT a valid property for this type; calling its associated methods will always yield an empty-equivalent value such as false, nil, or empty string. This includes instances where it should return itself. This ugliness is a symptom of the fundamental design of the ActivityStream vocabulary as instead of 'W-is-a-X' relationships it contains the notion of 'W-is-a-X-except-for-Y'.
func (t *OrderedCollectionPage) AppendItemsObject(v ObjectType) {

}

// PrependItemsObject is NOT a valid property for this type; calling its associated methods will always yield an empty-equivalent value such as false, nil, or empty string. This includes instances where it should return itself. This ugliness is a symptom of the fundamental design of the ActivityStream vocabulary as instead of 'W-is-a-X' relationships it contains the notion of 'W-is-a-X-except-for-Y'.
func (t *OrderedCollectionPage) PrependItemsObject(v ObjectType) {

}

// RemoveItemsObject is NOT a valid property for this type; calling its associated methods will always yield an empty-equivalent value such as false, nil, or empty string. This includes instances where it should return itself. This ugliness is a symptom of the fundamental design of the ActivityStream vocabulary as instead of 'W-is-a-X' relationships it contains the notion of 'W-is-a-X-except-for-Y'.
func (t *OrderedCollectionPage) RemoveItemsObject(index int) {

}

// IsItemsLink is NOT a valid property for this type; calling its associated methods will always yield an empty-equivalent value such as false, nil, or empty string. This includes instances where it should return itself. This ugliness is a symptom of the fundamental design of the ActivityStream vocabulary as instead of 'W-is-a-X' relationships it contains the notion of 'W-is-a-X-except-for-Y'.
func (t *OrderedCollectionPage) IsItemsLink(index int) (ok bool) {
	return false

}

// GetItemsLink is NOT a valid property for this type; calling its associated methods will always yield an empty-equivalent value such as false, nil, or empty string. This includes instances where it should return itself. This ugliness is a symptom of the fundamental design of the ActivityStream vocabulary as instead of 'W-is-a-X' relationships it contains the notion of 'W-is-a-X-except-for-Y'.
func (t *OrderedCollectionPage) GetItemsLink(index int) (v LinkType) {
	return nil

}

// AppendItemsLink is NOT a valid property for this type; calling its associated methods will always yield an empty-equivalent value such as false, nil, or empty string. This includes instances where it should return itself. This ugliness is a symptom of the fundamental design of the ActivityStream vocabulary as instead of 'W-is-a-X' relationships it contains the notion of 'W-is-a-X-except-for-Y'.
func (t *OrderedCollectionPage) AppendItemsLink(v LinkType) {

}

// PrependItemsLink is NOT a valid property for this type; calling its associated methods will always yield an empty-equivalent value such as false, nil, or empty string. This includes instances where it should return itself. This ugliness is a symptom of the fundamental design of the ActivityStream vocabulary as instead of 'W-is-a-X' relationships it contains the notion of 'W-is-a-X-except-for-Y'.
func (t *OrderedCollectionPage) PrependItemsLink(v LinkType) {

}

// RemoveItemsLink is NOT a valid property for this type; calling its associated methods will always yield an empty-equivalent value such as false, nil, or empty string. This includes instances where it should return itself. This ugliness is a symptom of the fundamental design of the ActivityStream vocabulary as instead of 'W-is-a-X' relationships it contains the notion of 'W-is-a-X-except-for-Y'.
func (t *OrderedCollectionPage) RemoveItemsLink(index int) {

}

// IsItemsIRI is NOT a valid property for this type; calling its associated methods will always yield an empty-equivalent value such as false, nil, or empty string. This includes instances where it should return itself. This ugliness is a symptom of the fundamental design of the ActivityStream vocabulary as instead of 'W-is-a-X' relationships it contains the notion of 'W-is-a-X-except-for-Y'.
func (t *OrderedCollectionPage) IsItemsIRI(index int) (ok bool) {
	return false

}

// GetItemsIRI is NOT a valid property for this type; calling its associated methods will always yield an empty-equivalent value such as false, nil, or empty string. This includes instances where it should return itself. This ugliness is a symptom of the fundamental design of the ActivityStream vocabulary as instead of 'W-is-a-X' relationships it contains the notion of 'W-is-a-X-except-for-Y'.
func (t *OrderedCollectionPage) GetItemsIRI(index int) (v *url.URL) {
	return nil

}

// AppendItemsIRI is NOT a valid property for this type; calling its associated methods will always yield an empty-equivalent value such as false, nil, or empty string. This includes instances where it should return itself. This ugliness is a symptom of the fundamental design of the ActivityStream vocabulary as instead of 'W-is-a-X' relationships it contains the notion of 'W-is-a-X-except-for-Y'.
func (t *OrderedCollectionPage) AppendItemsIRI(v *url.URL) {

}

// PrependItemsIRI is NOT a valid property for this type; calling its associated methods will always yield an empty-equivalent value such as false, nil, or empty string. This includes instances where it should return itself. This ugliness is a symptom of the fundamental design of the ActivityStream vocabulary as instead of 'W-is-a-X' relationships it contains the notion of 'W-is-a-X-except-for-Y'.
func (t *OrderedCollectionPage) PrependItemsIRI(v *url.URL) {

}

// RemoveItemsIRI is NOT a valid property for this type; calling its associated methods will always yield an empty-equivalent value such as false, nil, or empty string. This includes instances where it should return itself. This ugliness is a symptom of the fundamental design of the ActivityStream vocabulary as instead of 'W-is-a-X' relationships it contains the notion of 'W-is-a-X-except-for-Y'.
func (t *OrderedCollectionPage) RemoveItemsIRI(index int) {

}

// HasUnknownItems is NOT a valid property for this type; calling its associated methods will always yield an empty-equivalent value such as false, nil, or empty string. This includes instances where it should return itself. This ugliness is a symptom of the fundamental design of the ActivityStream vocabulary as instead of 'W-is-a-X' relationships it contains the notion of 'W-is-a-X-except-for-Y'.
func (t *OrderedCollectionPage) HasUnknownItems() (ok bool) {
	return false

}

// GetUnknownItems is NOT a valid property for this type; calling its associated methods will always yield an empty-equivalent value such as false, nil, or empty string. This includes instances where it should return itself. This ugliness is a symptom of the fundamental design of the ActivityStream vocabulary as instead of 'W-is-a-X' relationships it contains the notion of 'W-is-a-X-except-for-Y'.
func (t *OrderedCollectionPage) GetUnknownItems() (v interface{}) {
	return nil

}

// SetUnknownItems is NOT a valid property for this type; calling its associated methods will always yield an empty-equivalent value such as false, nil, or empty string. This includes instances where it should return itself. This ugliness is a symptom of the fundamental design of the ActivityStream vocabulary as instead of 'W-is-a-X' relationships it contains the notion of 'W-is-a-X-except-for-Y'.
func (t *OrderedCollectionPage) SetUnknownItems(i interface{}) {

}

// AddUnknown adds a raw extension to this object with the specified key
func (t *OrderedCollectionPage) AddUnknown(k string, i interface{}) (this *OrderedCollectionPage) {
	if t.unknown_ == nil {
		t.unknown_ = make(map[string]interface{})
	}
	t.unknown_[k] = i
	return t

}

// HasUnknown returns true if there is an unknown object with the specified key
func (t *OrderedCollectionPage) HasUnknown(k string) (b bool) {
	if t.unknown_ == nil {
		return false
	}
	_, ok := t.unknown_[k]
	return ok

}

// RemoveUnknown removes a raw extension from this object with the specified key
func (t *OrderedCollectionPage) RemoveUnknown(k string) (this *OrderedCollectionPage) {
	delete(t.unknown_, k)
	return t

}

// Serialize turns this object into a map[string]interface{}. Note that the "type" property will automatically be populated with "OrderedCollectionPage" if not manually set by the caller
func (t *OrderedCollectionPage) Serialize() (m map[string]interface{}, err error) {
	m = make(map[string]interface{})
	for k, v := range t.unknown_ {
		m[k] = unknownValueSerialize(v)
	}
	var typeAlreadySet bool
	for _, k := range t.typeName {
		if ks, ok := k.(string); ok {
			if ks == "OrderedCollectionPage" {
				typeAlreadySet = true
				break
			}
		}
	}
	if !typeAlreadySet {
		t.typeName = append(t.typeName, "OrderedCollectionPage")
	}
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.startIndex != nil {
		if v, err := serializeStartIndexOrderedCollectionPageIntermediateType(t.startIndex); err == nil {
			m["startIndex"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.next != nil {
		if v, err := serializeNextOrderedCollectionPageIntermediateType(t.next); err == nil {
			m["next"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.prev != nil {
		if v, err := serializePrevOrderedCollectionPageIntermediateType(t.prev); err == nil {
			m["prev"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceOrderedItemsOrderedCollectionPageIntermediateType(t.orderedItems); err == nil && v != nil {
		if len(v) == 1 {
			m["orderedItems"] = v[0]
		} else {
			m["orderedItems"] = v
		}
	} else if err != nil {
		return m, err
	}
	// End generation by generateNonFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.current != nil {
		if v, err := serializeCurrentOrderedCollectionPageIntermediateType(t.current); err == nil {
			m["current"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.first != nil {
		if v, err := serializeFirstOrderedCollectionPageIntermediateType(t.first); err == nil {
			m["first"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.last != nil {
		if v, err := serializeLastOrderedCollectionPageIntermediateType(t.last); err == nil {
			m["last"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.totalItems != nil {
		if v, err := serializeTotalItemsOrderedCollectionPageIntermediateType(t.totalItems); err == nil {
			m["totalItems"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.altitude != nil {
		if v, err := serializeAltitudeOrderedCollectionPageIntermediateType(t.altitude); err == nil {
			m["altitude"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceAttachmentOrderedCollectionPageIntermediateType(t.attachment); err == nil && v != nil {
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
	if v, err := serializeSliceAttributedToOrderedCollectionPageIntermediateType(t.attributedTo); err == nil && v != nil {
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
	if v, err := serializeSliceAudienceOrderedCollectionPageIntermediateType(t.audience); err == nil && v != nil {
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
	if v, err := serializeSliceContentOrderedCollectionPageIntermediateType(t.content); err == nil && v != nil {
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
	if v, err := serializeSliceContextOrderedCollectionPageIntermediateType(t.context); err == nil && v != nil {
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
	if v, err := serializeSliceNameOrderedCollectionPageIntermediateType(t.name); err == nil && v != nil {
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
		if v, err := serializeEndTimeOrderedCollectionPageIntermediateType(t.endTime); err == nil {
			m["endTime"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceGeneratorOrderedCollectionPageIntermediateType(t.generator); err == nil && v != nil {
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
	if v, err := serializeSliceIconOrderedCollectionPageIntermediateType(t.icon); err == nil && v != nil {
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
	if v, err := serializeSliceImageOrderedCollectionPageIntermediateType(t.image); err == nil && v != nil {
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
	if v, err := serializeSliceInReplyToOrderedCollectionPageIntermediateType(t.inReplyTo); err == nil && v != nil {
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
	if v, err := serializeSliceLocationOrderedCollectionPageIntermediateType(t.location); err == nil && v != nil {
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
	if v, err := serializeSlicePreviewOrderedCollectionPageIntermediateType(t.preview); err == nil && v != nil {
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
		if v, err := serializePublishedOrderedCollectionPageIntermediateType(t.published); err == nil {
			m["published"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.replies != nil {
		if v, err := serializeRepliesOrderedCollectionPageIntermediateType(t.replies); err == nil {
			m["replies"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.startTime != nil {
		if v, err := serializeStartTimeOrderedCollectionPageIntermediateType(t.startTime); err == nil {
			m["startTime"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceSummaryOrderedCollectionPageIntermediateType(t.summary); err == nil && v != nil {
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
	if v, err := serializeSliceTagOrderedCollectionPageIntermediateType(t.tag); err == nil && v != nil {
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
		if v, err := serializeUpdatedOrderedCollectionPageIntermediateType(t.updated); err == nil {
			m["updated"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateNonFunctionalMultiTypeDefinition
	if v, err := serializeSliceUrlOrderedCollectionPageIntermediateType(t.url); err == nil && v != nil {
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
	if v, err := serializeSliceToOrderedCollectionPageIntermediateType(t.to); err == nil && v != nil {
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
	if v, err := serializeSliceBtoOrderedCollectionPageIntermediateType(t.bto); err == nil && v != nil {
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
	if v, err := serializeSliceCcOrderedCollectionPageIntermediateType(t.cc); err == nil && v != nil {
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
	if v, err := serializeSliceBccOrderedCollectionPageIntermediateType(t.bcc); err == nil && v != nil {
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
		if v, err := serializeMediaTypeOrderedCollectionPageIntermediateType(t.mediaType); err == nil {
			m["mediaType"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.duration != nil {
		if v, err := serializeDurationOrderedCollectionPageIntermediateType(t.duration); err == nil {
			m["duration"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.source != nil {
		if v, err := serializeSourceOrderedCollectionPageIntermediateType(t.source); err == nil {
			m["source"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.inbox != nil {
		if v, err := serializeInboxOrderedCollectionPageIntermediateType(t.inbox); err == nil {
			m["inbox"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.outbox != nil {
		if v, err := serializeOutboxOrderedCollectionPageIntermediateType(t.outbox); err == nil {
			m["outbox"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.following != nil {
		if v, err := serializeFollowingOrderedCollectionPageIntermediateType(t.following); err == nil {
			m["following"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.followers != nil {
		if v, err := serializeFollowersOrderedCollectionPageIntermediateType(t.followers); err == nil {
			m["followers"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.liked != nil {
		if v, err := serializeLikedOrderedCollectionPageIntermediateType(t.liked); err == nil {
			m["liked"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	// Begin generation by generateFunctionalMultiTypeDefinition
	if t.likes != nil {
		if v, err := serializeLikesOrderedCollectionPageIntermediateType(t.likes); err == nil {
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
		if v, err := serializePreferredUsernameOrderedCollectionPageIntermediateType(t.preferredUsername); err == nil {
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
		if v, err := serializeEndpointsOrderedCollectionPageIntermediateType(t.endpoints); err == nil {
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
	if t.partOf != nil {
		if v, err := serializePartOfOrderedCollectionPageIntermediateType(t.partOf); err == nil {
			m["partOf"] = v
		} else {
			return m, err
		}
	}
	// End generation by generateFunctionalMultiTypeDefinition
	return

}

// Deserialize populates this object from a map[string]interface{}
func (t *OrderedCollectionPage) Deserialize(m map[string]interface{}) (err error) {
	for k, v := range m {
		handled := false
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "startIndex" {
				t.startIndex, err = deserializeStartIndexOrderedCollectionPageIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "next" {
				t.next, err = deserializeNextOrderedCollectionPageIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "prev" {
				t.prev, err = deserializePrevOrderedCollectionPageIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "orderedItems" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeOrderedItemsOrderedCollectionPageIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.orderedItems = []*orderedItemsOrderedCollectionPageIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.orderedItems, err = deserializeSliceOrderedItemsOrderedCollectionPageIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &orderedItemsOrderedCollectionPageIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.orderedItems = []*orderedItemsOrderedCollectionPageIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "current" {
				t.current, err = deserializeCurrentOrderedCollectionPageIntermediateType(v)
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
				t.first, err = deserializeFirstOrderedCollectionPageIntermediateType(v)
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
				t.last, err = deserializeLastOrderedCollectionPageIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "totalItems" {
				t.totalItems, err = deserializeTotalItemsOrderedCollectionPageIntermediateType(v)
				if err != nil {
					return err
				}
				handled = true
			}
			// End generation by generateFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "altitude" {
				t.altitude, err = deserializeAltitudeOrderedCollectionPageIntermediateType(v)
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
					tmp, err := deserializeAttachmentOrderedCollectionPageIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.attachment = []*attachmentOrderedCollectionPageIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.attachment, err = deserializeSliceAttachmentOrderedCollectionPageIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &attachmentOrderedCollectionPageIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.attachment = []*attachmentOrderedCollectionPageIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "attributedTo" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeAttributedToOrderedCollectionPageIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.attributedTo = []*attributedToOrderedCollectionPageIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.attributedTo, err = deserializeSliceAttributedToOrderedCollectionPageIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &attributedToOrderedCollectionPageIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.attributedTo = []*attributedToOrderedCollectionPageIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "audience" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeAudienceOrderedCollectionPageIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.audience = []*audienceOrderedCollectionPageIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.audience, err = deserializeSliceAudienceOrderedCollectionPageIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &audienceOrderedCollectionPageIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.audience = []*audienceOrderedCollectionPageIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "content" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeContentOrderedCollectionPageIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.content = []*contentOrderedCollectionPageIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.content, err = deserializeSliceContentOrderedCollectionPageIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &contentOrderedCollectionPageIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.content = []*contentOrderedCollectionPageIntermediateType{tmp}
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
					tmp, err := deserializeContextOrderedCollectionPageIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.context = []*contextOrderedCollectionPageIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.context, err = deserializeSliceContextOrderedCollectionPageIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &contextOrderedCollectionPageIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.context = []*contextOrderedCollectionPageIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "name" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeNameOrderedCollectionPageIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.name = []*nameOrderedCollectionPageIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.name, err = deserializeSliceNameOrderedCollectionPageIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &nameOrderedCollectionPageIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.name = []*nameOrderedCollectionPageIntermediateType{tmp}
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
				t.endTime, err = deserializeEndTimeOrderedCollectionPageIntermediateType(v)
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
					tmp, err := deserializeGeneratorOrderedCollectionPageIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.generator = []*generatorOrderedCollectionPageIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.generator, err = deserializeSliceGeneratorOrderedCollectionPageIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &generatorOrderedCollectionPageIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.generator = []*generatorOrderedCollectionPageIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "icon" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeIconOrderedCollectionPageIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.icon = []*iconOrderedCollectionPageIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.icon, err = deserializeSliceIconOrderedCollectionPageIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &iconOrderedCollectionPageIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.icon = []*iconOrderedCollectionPageIntermediateType{tmp}
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
					tmp, err := deserializeImageOrderedCollectionPageIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.image = []*imageOrderedCollectionPageIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.image, err = deserializeSliceImageOrderedCollectionPageIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &imageOrderedCollectionPageIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.image = []*imageOrderedCollectionPageIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "inReplyTo" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeInReplyToOrderedCollectionPageIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.inReplyTo = []*inReplyToOrderedCollectionPageIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.inReplyTo, err = deserializeSliceInReplyToOrderedCollectionPageIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &inReplyToOrderedCollectionPageIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.inReplyTo = []*inReplyToOrderedCollectionPageIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "location" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeLocationOrderedCollectionPageIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.location = []*locationOrderedCollectionPageIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.location, err = deserializeSliceLocationOrderedCollectionPageIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &locationOrderedCollectionPageIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.location = []*locationOrderedCollectionPageIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "preview" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializePreviewOrderedCollectionPageIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.preview = []*previewOrderedCollectionPageIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.preview, err = deserializeSlicePreviewOrderedCollectionPageIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &previewOrderedCollectionPageIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.preview = []*previewOrderedCollectionPageIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "published" {
				t.published, err = deserializePublishedOrderedCollectionPageIntermediateType(v)
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
				t.replies, err = deserializeRepliesOrderedCollectionPageIntermediateType(v)
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
				t.startTime, err = deserializeStartTimeOrderedCollectionPageIntermediateType(v)
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
					tmp, err := deserializeSummaryOrderedCollectionPageIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.summary = []*summaryOrderedCollectionPageIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.summary, err = deserializeSliceSummaryOrderedCollectionPageIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &summaryOrderedCollectionPageIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.summary = []*summaryOrderedCollectionPageIntermediateType{tmp}
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
					tmp, err := deserializeTagOrderedCollectionPageIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.tag = []*tagOrderedCollectionPageIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.tag, err = deserializeSliceTagOrderedCollectionPageIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &tagOrderedCollectionPageIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.tag = []*tagOrderedCollectionPageIntermediateType{tmp}
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
				t.updated, err = deserializeUpdatedOrderedCollectionPageIntermediateType(v)
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
					tmp, err := deserializeUrlOrderedCollectionPageIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.url = []*urlOrderedCollectionPageIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.url, err = deserializeSliceUrlOrderedCollectionPageIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &urlOrderedCollectionPageIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.url = []*urlOrderedCollectionPageIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "to" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeToOrderedCollectionPageIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.to = []*toOrderedCollectionPageIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.to, err = deserializeSliceToOrderedCollectionPageIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &toOrderedCollectionPageIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.to = []*toOrderedCollectionPageIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "bto" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeBtoOrderedCollectionPageIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.bto = []*btoOrderedCollectionPageIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.bto, err = deserializeSliceBtoOrderedCollectionPageIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &btoOrderedCollectionPageIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.bto = []*btoOrderedCollectionPageIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "cc" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeCcOrderedCollectionPageIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.cc = []*ccOrderedCollectionPageIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.cc, err = deserializeSliceCcOrderedCollectionPageIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &ccOrderedCollectionPageIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.cc = []*ccOrderedCollectionPageIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateNonFunctionalMultiTypeDefinition
			if k == "bcc" {
				if tmpMap, ok := v.(map[string]interface{}); ok {
					tmp, err := deserializeBccOrderedCollectionPageIntermediateType(tmpMap)
					if err != nil {
						return err
					}
					t.bcc = []*bccOrderedCollectionPageIntermediateType{tmp}
					handled = true
				} else if tmpSlice, ok := v.([]interface{}); ok {
					t.bcc, err = deserializeSliceBccOrderedCollectionPageIntermediateType(tmpSlice)
					if err != nil {
						return err
					}
					handled = true
				} else {
					tmp := &bccOrderedCollectionPageIntermediateType{}
					err = tmp.Deserialize(v)
					if err != nil {
						return err
					}
					t.bcc = []*bccOrderedCollectionPageIntermediateType{tmp}
					handled = true
				}
			}
			// End generation by generateNonFunctionalMultiTypeDefinition
		}
		if !handled {
			// Begin generation by generateFunctionalMultiTypeDefinition
			if k == "mediaType" {
				t.mediaType, err = deserializeMediaTypeOrderedCollectionPageIntermediateType(v)
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
				t.duration, err = deserializeDurationOrderedCollectionPageIntermediateType(v)
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
				t.source, err = deserializeSourceOrderedCollectionPageIntermediateType(v)
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
				t.inbox, err = deserializeInboxOrderedCollectionPageIntermediateType(v)
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
				t.outbox, err = deserializeOutboxOrderedCollectionPageIntermediateType(v)
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
				t.following, err = deserializeFollowingOrderedCollectionPageIntermediateType(v)
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
				t.followers, err = deserializeFollowersOrderedCollectionPageIntermediateType(v)
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
				t.liked, err = deserializeLikedOrderedCollectionPageIntermediateType(v)
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
				t.likes, err = deserializeLikesOrderedCollectionPageIntermediateType(v)
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
				t.preferredUsername, err = deserializePreferredUsernameOrderedCollectionPageIntermediateType(v)
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
				t.endpoints, err = deserializeEndpointsOrderedCollectionPageIntermediateType(v)
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
			if k == "partOf" {
				t.partOf, err = deserializePartOfOrderedCollectionPageIntermediateType(v)
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

// startIndexOrderedCollectionPageIntermediateType will only have one of its values set at most
type startIndexOrderedCollectionPageIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *int64 type for startIndex property
	nonNegativeInteger *int64
	// Stores possible *url.URL type for startIndex property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *startIndexOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *startIndexOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// nextOrderedCollectionPageIntermediateType will only have one of its values set at most
type nextOrderedCollectionPageIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible OrderedCollectionPageType type for next property
	OrderedCollectionPage OrderedCollectionPageType
	// Stores possible LinkType type for next property
	Link LinkType
	// Stores possible *url.URL type for next property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *nextOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
					if t.OrderedCollectionPage, ok = resolveObject(kind).(OrderedCollectionPageType); t.OrderedCollectionPage != nil && ok {
						err = t.OrderedCollectionPage.Deserialize(m)
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
func (t *nextOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
	if t.OrderedCollectionPage != nil {
		i, err = t.OrderedCollectionPage.Serialize()
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

// prevOrderedCollectionPageIntermediateType will only have one of its values set at most
type prevOrderedCollectionPageIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible OrderedCollectionPageType type for prev property
	OrderedCollectionPage OrderedCollectionPageType
	// Stores possible LinkType type for prev property
	Link LinkType
	// Stores possible *url.URL type for prev property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *prevOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
					if t.OrderedCollectionPage, ok = resolveObject(kind).(OrderedCollectionPageType); t.OrderedCollectionPage != nil && ok {
						err = t.OrderedCollectionPage.Deserialize(m)
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
func (t *prevOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
	if t.OrderedCollectionPage != nil {
		i, err = t.OrderedCollectionPage.Serialize()
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

// orderedItemsOrderedCollectionPageIntermediateType will only have one of its values set at most
type orderedItemsOrderedCollectionPageIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for orderedItems property
	Object ObjectType
	// Stores possible LinkType type for orderedItems property
	Link LinkType
	// Stores possible *url.URL type for orderedItems property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *orderedItemsOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *orderedItemsOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// currentOrderedCollectionPageIntermediateType will only have one of its values set at most
type currentOrderedCollectionPageIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible OrderedCollectionPageType type for current property
	OrderedCollectionPage OrderedCollectionPageType
	// Stores possible LinkType type for current property
	Link LinkType
	// Stores possible *url.URL type for current property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *currentOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
					if t.OrderedCollectionPage, ok = resolveObject(kind).(OrderedCollectionPageType); t.OrderedCollectionPage != nil && ok {
						err = t.OrderedCollectionPage.Deserialize(m)
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
func (t *currentOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
	if t.OrderedCollectionPage != nil {
		i, err = t.OrderedCollectionPage.Serialize()
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

// firstOrderedCollectionPageIntermediateType will only have one of its values set at most
type firstOrderedCollectionPageIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible OrderedCollectionPageType type for first property
	OrderedCollectionPage OrderedCollectionPageType
	// Stores possible LinkType type for first property
	Link LinkType
	// Stores possible *url.URL type for first property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *firstOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
					if t.OrderedCollectionPage, ok = resolveObject(kind).(OrderedCollectionPageType); t.OrderedCollectionPage != nil && ok {
						err = t.OrderedCollectionPage.Deserialize(m)
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
func (t *firstOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
	if t.OrderedCollectionPage != nil {
		i, err = t.OrderedCollectionPage.Serialize()
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

// lastOrderedCollectionPageIntermediateType will only have one of its values set at most
type lastOrderedCollectionPageIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible OrderedCollectionPageType type for last property
	OrderedCollectionPage OrderedCollectionPageType
	// Stores possible LinkType type for last property
	Link LinkType
	// Stores possible *url.URL type for last property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *lastOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
					if t.OrderedCollectionPage, ok = resolveObject(kind).(OrderedCollectionPageType); t.OrderedCollectionPage != nil && ok {
						err = t.OrderedCollectionPage.Deserialize(m)
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
func (t *lastOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
	if t.OrderedCollectionPage != nil {
		i, err = t.OrderedCollectionPage.Serialize()
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

// totalItemsOrderedCollectionPageIntermediateType will only have one of its values set at most
type totalItemsOrderedCollectionPageIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *int64 type for totalItems property
	nonNegativeInteger *int64
	// Stores possible *url.URL type for totalItems property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *totalItemsOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *totalItemsOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// altitudeOrderedCollectionPageIntermediateType will only have one of its values set at most
type altitudeOrderedCollectionPageIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *float64 type for altitude property
	float *float64
	// Stores possible *url.URL type for altitude property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *altitudeOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *altitudeOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// attachmentOrderedCollectionPageIntermediateType will only have one of its values set at most
type attachmentOrderedCollectionPageIntermediateType struct {
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
func (t *attachmentOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *attachmentOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// attributedToOrderedCollectionPageIntermediateType will only have one of its values set at most
type attributedToOrderedCollectionPageIntermediateType struct {
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
func (t *attributedToOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *attributedToOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// audienceOrderedCollectionPageIntermediateType will only have one of its values set at most
type audienceOrderedCollectionPageIntermediateType struct {
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
func (t *audienceOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *audienceOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// contentOrderedCollectionPageIntermediateType will only have one of its values set at most
type contentOrderedCollectionPageIntermediateType struct {
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
func (t *contentOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *contentOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// contextOrderedCollectionPageIntermediateType will only have one of its values set at most
type contextOrderedCollectionPageIntermediateType struct {
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
func (t *contextOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *contextOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// nameOrderedCollectionPageIntermediateType will only have one of its values set at most
type nameOrderedCollectionPageIntermediateType struct {
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
func (t *nameOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *nameOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// endTimeOrderedCollectionPageIntermediateType will only have one of its values set at most
type endTimeOrderedCollectionPageIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *time.Time type for endTime property
	dateTime *time.Time
	// Stores possible *url.URL type for endTime property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *endTimeOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *endTimeOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// generatorOrderedCollectionPageIntermediateType will only have one of its values set at most
type generatorOrderedCollectionPageIntermediateType struct {
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
func (t *generatorOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *generatorOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// iconOrderedCollectionPageIntermediateType will only have one of its values set at most
type iconOrderedCollectionPageIntermediateType struct {
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
func (t *iconOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *iconOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// imageOrderedCollectionPageIntermediateType will only have one of its values set at most
type imageOrderedCollectionPageIntermediateType struct {
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
func (t *imageOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *imageOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// inReplyToOrderedCollectionPageIntermediateType will only have one of its values set at most
type inReplyToOrderedCollectionPageIntermediateType struct {
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
func (t *inReplyToOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *inReplyToOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// locationOrderedCollectionPageIntermediateType will only have one of its values set at most
type locationOrderedCollectionPageIntermediateType struct {
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
func (t *locationOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *locationOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// previewOrderedCollectionPageIntermediateType will only have one of its values set at most
type previewOrderedCollectionPageIntermediateType struct {
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
func (t *previewOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *previewOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// publishedOrderedCollectionPageIntermediateType will only have one of its values set at most
type publishedOrderedCollectionPageIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *time.Time type for published property
	dateTime *time.Time
	// Stores possible *url.URL type for published property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *publishedOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *publishedOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// repliesOrderedCollectionPageIntermediateType will only have one of its values set at most
type repliesOrderedCollectionPageIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible CollectionType type for replies property
	Collection CollectionType
	// Stores possible *url.URL type for replies property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *repliesOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *repliesOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// startTimeOrderedCollectionPageIntermediateType will only have one of its values set at most
type startTimeOrderedCollectionPageIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *time.Time type for startTime property
	dateTime *time.Time
	// Stores possible *url.URL type for startTime property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *startTimeOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *startTimeOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// summaryOrderedCollectionPageIntermediateType will only have one of its values set at most
type summaryOrderedCollectionPageIntermediateType struct {
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
func (t *summaryOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *summaryOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// tagOrderedCollectionPageIntermediateType will only have one of its values set at most
type tagOrderedCollectionPageIntermediateType struct {
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
func (t *tagOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *tagOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// updatedOrderedCollectionPageIntermediateType will only have one of its values set at most
type updatedOrderedCollectionPageIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *time.Time type for updated property
	dateTime *time.Time
	// Stores possible *url.URL type for updated property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *updatedOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *updatedOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// urlOrderedCollectionPageIntermediateType will only have one of its values set at most
type urlOrderedCollectionPageIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *url.URL type for url property
	anyURI *url.URL
	// Stores possible LinkType type for url property
	Link LinkType
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *urlOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *urlOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// toOrderedCollectionPageIntermediateType will only have one of its values set at most
type toOrderedCollectionPageIntermediateType struct {
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
func (t *toOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *toOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// btoOrderedCollectionPageIntermediateType will only have one of its values set at most
type btoOrderedCollectionPageIntermediateType struct {
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
func (t *btoOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *btoOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// ccOrderedCollectionPageIntermediateType will only have one of its values set at most
type ccOrderedCollectionPageIntermediateType struct {
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
func (t *ccOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *ccOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// bccOrderedCollectionPageIntermediateType will only have one of its values set at most
type bccOrderedCollectionPageIntermediateType struct {
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
func (t *bccOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *bccOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// mediaTypeOrderedCollectionPageIntermediateType will only have one of its values set at most
type mediaTypeOrderedCollectionPageIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *string type for mediaType property
	mimeMediaTypeValue *string
	// Stores possible *url.URL type for mediaType property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *mediaTypeOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *mediaTypeOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// durationOrderedCollectionPageIntermediateType will only have one of its values set at most
type durationOrderedCollectionPageIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *time.Duration type for duration property
	duration *time.Duration
	// Stores possible *url.URL type for duration property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *durationOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *durationOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// sourceOrderedCollectionPageIntermediateType will only have one of its values set at most
type sourceOrderedCollectionPageIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for source property
	Object ObjectType
	// Stores possible *url.URL type for source property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *sourceOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *sourceOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// inboxOrderedCollectionPageIntermediateType will only have one of its values set at most
type inboxOrderedCollectionPageIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible OrderedCollectionType type for inbox property
	OrderedCollection OrderedCollectionType
	// Stores possible *url.URL type for inbox property
	anyURI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *inboxOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *inboxOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// outboxOrderedCollectionPageIntermediateType will only have one of its values set at most
type outboxOrderedCollectionPageIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible OrderedCollectionType type for outbox property
	OrderedCollection OrderedCollectionType
	// Stores possible *url.URL type for outbox property
	anyURI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *outboxOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *outboxOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// followingOrderedCollectionPageIntermediateType will only have one of its values set at most
type followingOrderedCollectionPageIntermediateType struct {
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
func (t *followingOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *followingOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// followersOrderedCollectionPageIntermediateType will only have one of its values set at most
type followersOrderedCollectionPageIntermediateType struct {
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
func (t *followersOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *followersOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// likedOrderedCollectionPageIntermediateType will only have one of its values set at most
type likedOrderedCollectionPageIntermediateType struct {
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
func (t *likedOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *likedOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// likesOrderedCollectionPageIntermediateType will only have one of its values set at most
type likesOrderedCollectionPageIntermediateType struct {
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
func (t *likesOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *likesOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// preferredUsernameOrderedCollectionPageIntermediateType will only have one of its values set at most
type preferredUsernameOrderedCollectionPageIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible *string type for preferredUsername property
	stringName *string
	// Stores possible *url.URL type for preferredUsername property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *preferredUsernameOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *preferredUsernameOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// endpointsOrderedCollectionPageIntermediateType will only have one of its values set at most
type endpointsOrderedCollectionPageIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible ObjectType type for endpoints property
	Object ObjectType
	// Stores possible *url.URL type for endpoints property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *endpointsOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *endpointsOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
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

// partOfOrderedCollectionPageIntermediateType will only have one of its values set at most
type partOfOrderedCollectionPageIntermediateType struct {
	// An unknown value.
	unknown_ interface{}
	// Stores possible LinkType type for partOf property
	Link LinkType
	// Stores possible CollectionType type for partOf property
	Collection CollectionType
	// Stores possible *url.URL type for partOf property
	IRI *url.URL
}

// Deserialize takes an interface{} and attempts to create a valid intermediate type.
func (t *partOfOrderedCollectionPageIntermediateType) Deserialize(i interface{}) (err error) {
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
func (t *partOfOrderedCollectionPageIntermediateType) Serialize() (i interface{}, err error) {
	if t.Link != nil {
		i, err = t.Link.Serialize()
		return
	}
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

// deserializestartIndexOrderedCollectionPageIntermediateType will accept a map to create a startIndexOrderedCollectionPageIntermediateType
func deserializeStartIndexOrderedCollectionPageIntermediateType(in interface{}) (t *startIndexOrderedCollectionPageIntermediateType, err error) {
	tmp := &startIndexOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice startIndexOrderedCollectionPageIntermediateType will accept a slice to create a slice of startIndexOrderedCollectionPageIntermediateType
func deserializeSliceStartIndexOrderedCollectionPageIntermediateType(in []interface{}) (t []*startIndexOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &startIndexOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializestartIndexOrderedCollectionPageIntermediateType will accept a startIndexOrderedCollectionPageIntermediateType to create a map
func serializeStartIndexOrderedCollectionPageIntermediateType(t *startIndexOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicestartIndexOrderedCollectionPageIntermediateType will accept a slice of startIndexOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceStartIndexOrderedCollectionPageIntermediateType(s []*startIndexOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializenextOrderedCollectionPageIntermediateType will accept a map to create a nextOrderedCollectionPageIntermediateType
func deserializeNextOrderedCollectionPageIntermediateType(in interface{}) (t *nextOrderedCollectionPageIntermediateType, err error) {
	tmp := &nextOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice nextOrderedCollectionPageIntermediateType will accept a slice to create a slice of nextOrderedCollectionPageIntermediateType
func deserializeSliceNextOrderedCollectionPageIntermediateType(in []interface{}) (t []*nextOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &nextOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializenextOrderedCollectionPageIntermediateType will accept a nextOrderedCollectionPageIntermediateType to create a map
func serializeNextOrderedCollectionPageIntermediateType(t *nextOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicenextOrderedCollectionPageIntermediateType will accept a slice of nextOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceNextOrderedCollectionPageIntermediateType(s []*nextOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeprevOrderedCollectionPageIntermediateType will accept a map to create a prevOrderedCollectionPageIntermediateType
func deserializePrevOrderedCollectionPageIntermediateType(in interface{}) (t *prevOrderedCollectionPageIntermediateType, err error) {
	tmp := &prevOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice prevOrderedCollectionPageIntermediateType will accept a slice to create a slice of prevOrderedCollectionPageIntermediateType
func deserializeSlicePrevOrderedCollectionPageIntermediateType(in []interface{}) (t []*prevOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &prevOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeprevOrderedCollectionPageIntermediateType will accept a prevOrderedCollectionPageIntermediateType to create a map
func serializePrevOrderedCollectionPageIntermediateType(t *prevOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceprevOrderedCollectionPageIntermediateType will accept a slice of prevOrderedCollectionPageIntermediateType to create a slice result
func serializeSlicePrevOrderedCollectionPageIntermediateType(s []*prevOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeorderedItemsOrderedCollectionPageIntermediateType will accept a map to create a orderedItemsOrderedCollectionPageIntermediateType
func deserializeOrderedItemsOrderedCollectionPageIntermediateType(in interface{}) (t *orderedItemsOrderedCollectionPageIntermediateType, err error) {
	tmp := &orderedItemsOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice orderedItemsOrderedCollectionPageIntermediateType will accept a slice to create a slice of orderedItemsOrderedCollectionPageIntermediateType
func deserializeSliceOrderedItemsOrderedCollectionPageIntermediateType(in []interface{}) (t []*orderedItemsOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &orderedItemsOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeorderedItemsOrderedCollectionPageIntermediateType will accept a orderedItemsOrderedCollectionPageIntermediateType to create a map
func serializeOrderedItemsOrderedCollectionPageIntermediateType(t *orderedItemsOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceorderedItemsOrderedCollectionPageIntermediateType will accept a slice of orderedItemsOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceOrderedItemsOrderedCollectionPageIntermediateType(s []*orderedItemsOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializecurrentOrderedCollectionPageIntermediateType will accept a map to create a currentOrderedCollectionPageIntermediateType
func deserializeCurrentOrderedCollectionPageIntermediateType(in interface{}) (t *currentOrderedCollectionPageIntermediateType, err error) {
	tmp := &currentOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice currentOrderedCollectionPageIntermediateType will accept a slice to create a slice of currentOrderedCollectionPageIntermediateType
func deserializeSliceCurrentOrderedCollectionPageIntermediateType(in []interface{}) (t []*currentOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &currentOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializecurrentOrderedCollectionPageIntermediateType will accept a currentOrderedCollectionPageIntermediateType to create a map
func serializeCurrentOrderedCollectionPageIntermediateType(t *currentOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicecurrentOrderedCollectionPageIntermediateType will accept a slice of currentOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceCurrentOrderedCollectionPageIntermediateType(s []*currentOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializefirstOrderedCollectionPageIntermediateType will accept a map to create a firstOrderedCollectionPageIntermediateType
func deserializeFirstOrderedCollectionPageIntermediateType(in interface{}) (t *firstOrderedCollectionPageIntermediateType, err error) {
	tmp := &firstOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice firstOrderedCollectionPageIntermediateType will accept a slice to create a slice of firstOrderedCollectionPageIntermediateType
func deserializeSliceFirstOrderedCollectionPageIntermediateType(in []interface{}) (t []*firstOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &firstOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializefirstOrderedCollectionPageIntermediateType will accept a firstOrderedCollectionPageIntermediateType to create a map
func serializeFirstOrderedCollectionPageIntermediateType(t *firstOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicefirstOrderedCollectionPageIntermediateType will accept a slice of firstOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceFirstOrderedCollectionPageIntermediateType(s []*firstOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializelastOrderedCollectionPageIntermediateType will accept a map to create a lastOrderedCollectionPageIntermediateType
func deserializeLastOrderedCollectionPageIntermediateType(in interface{}) (t *lastOrderedCollectionPageIntermediateType, err error) {
	tmp := &lastOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice lastOrderedCollectionPageIntermediateType will accept a slice to create a slice of lastOrderedCollectionPageIntermediateType
func deserializeSliceLastOrderedCollectionPageIntermediateType(in []interface{}) (t []*lastOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &lastOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializelastOrderedCollectionPageIntermediateType will accept a lastOrderedCollectionPageIntermediateType to create a map
func serializeLastOrderedCollectionPageIntermediateType(t *lastOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicelastOrderedCollectionPageIntermediateType will accept a slice of lastOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceLastOrderedCollectionPageIntermediateType(s []*lastOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializetotalItemsOrderedCollectionPageIntermediateType will accept a map to create a totalItemsOrderedCollectionPageIntermediateType
func deserializeTotalItemsOrderedCollectionPageIntermediateType(in interface{}) (t *totalItemsOrderedCollectionPageIntermediateType, err error) {
	tmp := &totalItemsOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice totalItemsOrderedCollectionPageIntermediateType will accept a slice to create a slice of totalItemsOrderedCollectionPageIntermediateType
func deserializeSliceTotalItemsOrderedCollectionPageIntermediateType(in []interface{}) (t []*totalItemsOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &totalItemsOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializetotalItemsOrderedCollectionPageIntermediateType will accept a totalItemsOrderedCollectionPageIntermediateType to create a map
func serializeTotalItemsOrderedCollectionPageIntermediateType(t *totalItemsOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicetotalItemsOrderedCollectionPageIntermediateType will accept a slice of totalItemsOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceTotalItemsOrderedCollectionPageIntermediateType(s []*totalItemsOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializealtitudeOrderedCollectionPageIntermediateType will accept a map to create a altitudeOrderedCollectionPageIntermediateType
func deserializeAltitudeOrderedCollectionPageIntermediateType(in interface{}) (t *altitudeOrderedCollectionPageIntermediateType, err error) {
	tmp := &altitudeOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice altitudeOrderedCollectionPageIntermediateType will accept a slice to create a slice of altitudeOrderedCollectionPageIntermediateType
func deserializeSliceAltitudeOrderedCollectionPageIntermediateType(in []interface{}) (t []*altitudeOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &altitudeOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializealtitudeOrderedCollectionPageIntermediateType will accept a altitudeOrderedCollectionPageIntermediateType to create a map
func serializeAltitudeOrderedCollectionPageIntermediateType(t *altitudeOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicealtitudeOrderedCollectionPageIntermediateType will accept a slice of altitudeOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceAltitudeOrderedCollectionPageIntermediateType(s []*altitudeOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeattachmentOrderedCollectionPageIntermediateType will accept a map to create a attachmentOrderedCollectionPageIntermediateType
func deserializeAttachmentOrderedCollectionPageIntermediateType(in interface{}) (t *attachmentOrderedCollectionPageIntermediateType, err error) {
	tmp := &attachmentOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice attachmentOrderedCollectionPageIntermediateType will accept a slice to create a slice of attachmentOrderedCollectionPageIntermediateType
func deserializeSliceAttachmentOrderedCollectionPageIntermediateType(in []interface{}) (t []*attachmentOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &attachmentOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeattachmentOrderedCollectionPageIntermediateType will accept a attachmentOrderedCollectionPageIntermediateType to create a map
func serializeAttachmentOrderedCollectionPageIntermediateType(t *attachmentOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceattachmentOrderedCollectionPageIntermediateType will accept a slice of attachmentOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceAttachmentOrderedCollectionPageIntermediateType(s []*attachmentOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeattributedToOrderedCollectionPageIntermediateType will accept a map to create a attributedToOrderedCollectionPageIntermediateType
func deserializeAttributedToOrderedCollectionPageIntermediateType(in interface{}) (t *attributedToOrderedCollectionPageIntermediateType, err error) {
	tmp := &attributedToOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice attributedToOrderedCollectionPageIntermediateType will accept a slice to create a slice of attributedToOrderedCollectionPageIntermediateType
func deserializeSliceAttributedToOrderedCollectionPageIntermediateType(in []interface{}) (t []*attributedToOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &attributedToOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeattributedToOrderedCollectionPageIntermediateType will accept a attributedToOrderedCollectionPageIntermediateType to create a map
func serializeAttributedToOrderedCollectionPageIntermediateType(t *attributedToOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceattributedToOrderedCollectionPageIntermediateType will accept a slice of attributedToOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceAttributedToOrderedCollectionPageIntermediateType(s []*attributedToOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeaudienceOrderedCollectionPageIntermediateType will accept a map to create a audienceOrderedCollectionPageIntermediateType
func deserializeAudienceOrderedCollectionPageIntermediateType(in interface{}) (t *audienceOrderedCollectionPageIntermediateType, err error) {
	tmp := &audienceOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice audienceOrderedCollectionPageIntermediateType will accept a slice to create a slice of audienceOrderedCollectionPageIntermediateType
func deserializeSliceAudienceOrderedCollectionPageIntermediateType(in []interface{}) (t []*audienceOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &audienceOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeaudienceOrderedCollectionPageIntermediateType will accept a audienceOrderedCollectionPageIntermediateType to create a map
func serializeAudienceOrderedCollectionPageIntermediateType(t *audienceOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceaudienceOrderedCollectionPageIntermediateType will accept a slice of audienceOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceAudienceOrderedCollectionPageIntermediateType(s []*audienceOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializecontentOrderedCollectionPageIntermediateType will accept a map to create a contentOrderedCollectionPageIntermediateType
func deserializeContentOrderedCollectionPageIntermediateType(in interface{}) (t *contentOrderedCollectionPageIntermediateType, err error) {
	tmp := &contentOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice contentOrderedCollectionPageIntermediateType will accept a slice to create a slice of contentOrderedCollectionPageIntermediateType
func deserializeSliceContentOrderedCollectionPageIntermediateType(in []interface{}) (t []*contentOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &contentOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializecontentOrderedCollectionPageIntermediateType will accept a contentOrderedCollectionPageIntermediateType to create a map
func serializeContentOrderedCollectionPageIntermediateType(t *contentOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicecontentOrderedCollectionPageIntermediateType will accept a slice of contentOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceContentOrderedCollectionPageIntermediateType(s []*contentOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializecontextOrderedCollectionPageIntermediateType will accept a map to create a contextOrderedCollectionPageIntermediateType
func deserializeContextOrderedCollectionPageIntermediateType(in interface{}) (t *contextOrderedCollectionPageIntermediateType, err error) {
	tmp := &contextOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice contextOrderedCollectionPageIntermediateType will accept a slice to create a slice of contextOrderedCollectionPageIntermediateType
func deserializeSliceContextOrderedCollectionPageIntermediateType(in []interface{}) (t []*contextOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &contextOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializecontextOrderedCollectionPageIntermediateType will accept a contextOrderedCollectionPageIntermediateType to create a map
func serializeContextOrderedCollectionPageIntermediateType(t *contextOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicecontextOrderedCollectionPageIntermediateType will accept a slice of contextOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceContextOrderedCollectionPageIntermediateType(s []*contextOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializenameOrderedCollectionPageIntermediateType will accept a map to create a nameOrderedCollectionPageIntermediateType
func deserializeNameOrderedCollectionPageIntermediateType(in interface{}) (t *nameOrderedCollectionPageIntermediateType, err error) {
	tmp := &nameOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice nameOrderedCollectionPageIntermediateType will accept a slice to create a slice of nameOrderedCollectionPageIntermediateType
func deserializeSliceNameOrderedCollectionPageIntermediateType(in []interface{}) (t []*nameOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &nameOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializenameOrderedCollectionPageIntermediateType will accept a nameOrderedCollectionPageIntermediateType to create a map
func serializeNameOrderedCollectionPageIntermediateType(t *nameOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicenameOrderedCollectionPageIntermediateType will accept a slice of nameOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceNameOrderedCollectionPageIntermediateType(s []*nameOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeendTimeOrderedCollectionPageIntermediateType will accept a map to create a endTimeOrderedCollectionPageIntermediateType
func deserializeEndTimeOrderedCollectionPageIntermediateType(in interface{}) (t *endTimeOrderedCollectionPageIntermediateType, err error) {
	tmp := &endTimeOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice endTimeOrderedCollectionPageIntermediateType will accept a slice to create a slice of endTimeOrderedCollectionPageIntermediateType
func deserializeSliceEndTimeOrderedCollectionPageIntermediateType(in []interface{}) (t []*endTimeOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &endTimeOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeendTimeOrderedCollectionPageIntermediateType will accept a endTimeOrderedCollectionPageIntermediateType to create a map
func serializeEndTimeOrderedCollectionPageIntermediateType(t *endTimeOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceendTimeOrderedCollectionPageIntermediateType will accept a slice of endTimeOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceEndTimeOrderedCollectionPageIntermediateType(s []*endTimeOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializegeneratorOrderedCollectionPageIntermediateType will accept a map to create a generatorOrderedCollectionPageIntermediateType
func deserializeGeneratorOrderedCollectionPageIntermediateType(in interface{}) (t *generatorOrderedCollectionPageIntermediateType, err error) {
	tmp := &generatorOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice generatorOrderedCollectionPageIntermediateType will accept a slice to create a slice of generatorOrderedCollectionPageIntermediateType
func deserializeSliceGeneratorOrderedCollectionPageIntermediateType(in []interface{}) (t []*generatorOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &generatorOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializegeneratorOrderedCollectionPageIntermediateType will accept a generatorOrderedCollectionPageIntermediateType to create a map
func serializeGeneratorOrderedCollectionPageIntermediateType(t *generatorOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicegeneratorOrderedCollectionPageIntermediateType will accept a slice of generatorOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceGeneratorOrderedCollectionPageIntermediateType(s []*generatorOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeiconOrderedCollectionPageIntermediateType will accept a map to create a iconOrderedCollectionPageIntermediateType
func deserializeIconOrderedCollectionPageIntermediateType(in interface{}) (t *iconOrderedCollectionPageIntermediateType, err error) {
	tmp := &iconOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice iconOrderedCollectionPageIntermediateType will accept a slice to create a slice of iconOrderedCollectionPageIntermediateType
func deserializeSliceIconOrderedCollectionPageIntermediateType(in []interface{}) (t []*iconOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &iconOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeiconOrderedCollectionPageIntermediateType will accept a iconOrderedCollectionPageIntermediateType to create a map
func serializeIconOrderedCollectionPageIntermediateType(t *iconOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceiconOrderedCollectionPageIntermediateType will accept a slice of iconOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceIconOrderedCollectionPageIntermediateType(s []*iconOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeimageOrderedCollectionPageIntermediateType will accept a map to create a imageOrderedCollectionPageIntermediateType
func deserializeImageOrderedCollectionPageIntermediateType(in interface{}) (t *imageOrderedCollectionPageIntermediateType, err error) {
	tmp := &imageOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice imageOrderedCollectionPageIntermediateType will accept a slice to create a slice of imageOrderedCollectionPageIntermediateType
func deserializeSliceImageOrderedCollectionPageIntermediateType(in []interface{}) (t []*imageOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &imageOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeimageOrderedCollectionPageIntermediateType will accept a imageOrderedCollectionPageIntermediateType to create a map
func serializeImageOrderedCollectionPageIntermediateType(t *imageOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceimageOrderedCollectionPageIntermediateType will accept a slice of imageOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceImageOrderedCollectionPageIntermediateType(s []*imageOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeinReplyToOrderedCollectionPageIntermediateType will accept a map to create a inReplyToOrderedCollectionPageIntermediateType
func deserializeInReplyToOrderedCollectionPageIntermediateType(in interface{}) (t *inReplyToOrderedCollectionPageIntermediateType, err error) {
	tmp := &inReplyToOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice inReplyToOrderedCollectionPageIntermediateType will accept a slice to create a slice of inReplyToOrderedCollectionPageIntermediateType
func deserializeSliceInReplyToOrderedCollectionPageIntermediateType(in []interface{}) (t []*inReplyToOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &inReplyToOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeinReplyToOrderedCollectionPageIntermediateType will accept a inReplyToOrderedCollectionPageIntermediateType to create a map
func serializeInReplyToOrderedCollectionPageIntermediateType(t *inReplyToOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceinReplyToOrderedCollectionPageIntermediateType will accept a slice of inReplyToOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceInReplyToOrderedCollectionPageIntermediateType(s []*inReplyToOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializelocationOrderedCollectionPageIntermediateType will accept a map to create a locationOrderedCollectionPageIntermediateType
func deserializeLocationOrderedCollectionPageIntermediateType(in interface{}) (t *locationOrderedCollectionPageIntermediateType, err error) {
	tmp := &locationOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice locationOrderedCollectionPageIntermediateType will accept a slice to create a slice of locationOrderedCollectionPageIntermediateType
func deserializeSliceLocationOrderedCollectionPageIntermediateType(in []interface{}) (t []*locationOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &locationOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializelocationOrderedCollectionPageIntermediateType will accept a locationOrderedCollectionPageIntermediateType to create a map
func serializeLocationOrderedCollectionPageIntermediateType(t *locationOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicelocationOrderedCollectionPageIntermediateType will accept a slice of locationOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceLocationOrderedCollectionPageIntermediateType(s []*locationOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializepreviewOrderedCollectionPageIntermediateType will accept a map to create a previewOrderedCollectionPageIntermediateType
func deserializePreviewOrderedCollectionPageIntermediateType(in interface{}) (t *previewOrderedCollectionPageIntermediateType, err error) {
	tmp := &previewOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice previewOrderedCollectionPageIntermediateType will accept a slice to create a slice of previewOrderedCollectionPageIntermediateType
func deserializeSlicePreviewOrderedCollectionPageIntermediateType(in []interface{}) (t []*previewOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &previewOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializepreviewOrderedCollectionPageIntermediateType will accept a previewOrderedCollectionPageIntermediateType to create a map
func serializePreviewOrderedCollectionPageIntermediateType(t *previewOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicepreviewOrderedCollectionPageIntermediateType will accept a slice of previewOrderedCollectionPageIntermediateType to create a slice result
func serializeSlicePreviewOrderedCollectionPageIntermediateType(s []*previewOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializepublishedOrderedCollectionPageIntermediateType will accept a map to create a publishedOrderedCollectionPageIntermediateType
func deserializePublishedOrderedCollectionPageIntermediateType(in interface{}) (t *publishedOrderedCollectionPageIntermediateType, err error) {
	tmp := &publishedOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice publishedOrderedCollectionPageIntermediateType will accept a slice to create a slice of publishedOrderedCollectionPageIntermediateType
func deserializeSlicePublishedOrderedCollectionPageIntermediateType(in []interface{}) (t []*publishedOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &publishedOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializepublishedOrderedCollectionPageIntermediateType will accept a publishedOrderedCollectionPageIntermediateType to create a map
func serializePublishedOrderedCollectionPageIntermediateType(t *publishedOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicepublishedOrderedCollectionPageIntermediateType will accept a slice of publishedOrderedCollectionPageIntermediateType to create a slice result
func serializeSlicePublishedOrderedCollectionPageIntermediateType(s []*publishedOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializerepliesOrderedCollectionPageIntermediateType will accept a map to create a repliesOrderedCollectionPageIntermediateType
func deserializeRepliesOrderedCollectionPageIntermediateType(in interface{}) (t *repliesOrderedCollectionPageIntermediateType, err error) {
	tmp := &repliesOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice repliesOrderedCollectionPageIntermediateType will accept a slice to create a slice of repliesOrderedCollectionPageIntermediateType
func deserializeSliceRepliesOrderedCollectionPageIntermediateType(in []interface{}) (t []*repliesOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &repliesOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializerepliesOrderedCollectionPageIntermediateType will accept a repliesOrderedCollectionPageIntermediateType to create a map
func serializeRepliesOrderedCollectionPageIntermediateType(t *repliesOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicerepliesOrderedCollectionPageIntermediateType will accept a slice of repliesOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceRepliesOrderedCollectionPageIntermediateType(s []*repliesOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializestartTimeOrderedCollectionPageIntermediateType will accept a map to create a startTimeOrderedCollectionPageIntermediateType
func deserializeStartTimeOrderedCollectionPageIntermediateType(in interface{}) (t *startTimeOrderedCollectionPageIntermediateType, err error) {
	tmp := &startTimeOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice startTimeOrderedCollectionPageIntermediateType will accept a slice to create a slice of startTimeOrderedCollectionPageIntermediateType
func deserializeSliceStartTimeOrderedCollectionPageIntermediateType(in []interface{}) (t []*startTimeOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &startTimeOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializestartTimeOrderedCollectionPageIntermediateType will accept a startTimeOrderedCollectionPageIntermediateType to create a map
func serializeStartTimeOrderedCollectionPageIntermediateType(t *startTimeOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicestartTimeOrderedCollectionPageIntermediateType will accept a slice of startTimeOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceStartTimeOrderedCollectionPageIntermediateType(s []*startTimeOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializesummaryOrderedCollectionPageIntermediateType will accept a map to create a summaryOrderedCollectionPageIntermediateType
func deserializeSummaryOrderedCollectionPageIntermediateType(in interface{}) (t *summaryOrderedCollectionPageIntermediateType, err error) {
	tmp := &summaryOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice summaryOrderedCollectionPageIntermediateType will accept a slice to create a slice of summaryOrderedCollectionPageIntermediateType
func deserializeSliceSummaryOrderedCollectionPageIntermediateType(in []interface{}) (t []*summaryOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &summaryOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializesummaryOrderedCollectionPageIntermediateType will accept a summaryOrderedCollectionPageIntermediateType to create a map
func serializeSummaryOrderedCollectionPageIntermediateType(t *summaryOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicesummaryOrderedCollectionPageIntermediateType will accept a slice of summaryOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceSummaryOrderedCollectionPageIntermediateType(s []*summaryOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializetagOrderedCollectionPageIntermediateType will accept a map to create a tagOrderedCollectionPageIntermediateType
func deserializeTagOrderedCollectionPageIntermediateType(in interface{}) (t *tagOrderedCollectionPageIntermediateType, err error) {
	tmp := &tagOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice tagOrderedCollectionPageIntermediateType will accept a slice to create a slice of tagOrderedCollectionPageIntermediateType
func deserializeSliceTagOrderedCollectionPageIntermediateType(in []interface{}) (t []*tagOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &tagOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializetagOrderedCollectionPageIntermediateType will accept a tagOrderedCollectionPageIntermediateType to create a map
func serializeTagOrderedCollectionPageIntermediateType(t *tagOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicetagOrderedCollectionPageIntermediateType will accept a slice of tagOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceTagOrderedCollectionPageIntermediateType(s []*tagOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeupdatedOrderedCollectionPageIntermediateType will accept a map to create a updatedOrderedCollectionPageIntermediateType
func deserializeUpdatedOrderedCollectionPageIntermediateType(in interface{}) (t *updatedOrderedCollectionPageIntermediateType, err error) {
	tmp := &updatedOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice updatedOrderedCollectionPageIntermediateType will accept a slice to create a slice of updatedOrderedCollectionPageIntermediateType
func deserializeSliceUpdatedOrderedCollectionPageIntermediateType(in []interface{}) (t []*updatedOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &updatedOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeupdatedOrderedCollectionPageIntermediateType will accept a updatedOrderedCollectionPageIntermediateType to create a map
func serializeUpdatedOrderedCollectionPageIntermediateType(t *updatedOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceupdatedOrderedCollectionPageIntermediateType will accept a slice of updatedOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceUpdatedOrderedCollectionPageIntermediateType(s []*updatedOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeurlOrderedCollectionPageIntermediateType will accept a map to create a urlOrderedCollectionPageIntermediateType
func deserializeUrlOrderedCollectionPageIntermediateType(in interface{}) (t *urlOrderedCollectionPageIntermediateType, err error) {
	tmp := &urlOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice urlOrderedCollectionPageIntermediateType will accept a slice to create a slice of urlOrderedCollectionPageIntermediateType
func deserializeSliceUrlOrderedCollectionPageIntermediateType(in []interface{}) (t []*urlOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &urlOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeurlOrderedCollectionPageIntermediateType will accept a urlOrderedCollectionPageIntermediateType to create a map
func serializeUrlOrderedCollectionPageIntermediateType(t *urlOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceurlOrderedCollectionPageIntermediateType will accept a slice of urlOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceUrlOrderedCollectionPageIntermediateType(s []*urlOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializetoOrderedCollectionPageIntermediateType will accept a map to create a toOrderedCollectionPageIntermediateType
func deserializeToOrderedCollectionPageIntermediateType(in interface{}) (t *toOrderedCollectionPageIntermediateType, err error) {
	tmp := &toOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice toOrderedCollectionPageIntermediateType will accept a slice to create a slice of toOrderedCollectionPageIntermediateType
func deserializeSliceToOrderedCollectionPageIntermediateType(in []interface{}) (t []*toOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &toOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializetoOrderedCollectionPageIntermediateType will accept a toOrderedCollectionPageIntermediateType to create a map
func serializeToOrderedCollectionPageIntermediateType(t *toOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicetoOrderedCollectionPageIntermediateType will accept a slice of toOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceToOrderedCollectionPageIntermediateType(s []*toOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializebtoOrderedCollectionPageIntermediateType will accept a map to create a btoOrderedCollectionPageIntermediateType
func deserializeBtoOrderedCollectionPageIntermediateType(in interface{}) (t *btoOrderedCollectionPageIntermediateType, err error) {
	tmp := &btoOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice btoOrderedCollectionPageIntermediateType will accept a slice to create a slice of btoOrderedCollectionPageIntermediateType
func deserializeSliceBtoOrderedCollectionPageIntermediateType(in []interface{}) (t []*btoOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &btoOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializebtoOrderedCollectionPageIntermediateType will accept a btoOrderedCollectionPageIntermediateType to create a map
func serializeBtoOrderedCollectionPageIntermediateType(t *btoOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicebtoOrderedCollectionPageIntermediateType will accept a slice of btoOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceBtoOrderedCollectionPageIntermediateType(s []*btoOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeccOrderedCollectionPageIntermediateType will accept a map to create a ccOrderedCollectionPageIntermediateType
func deserializeCcOrderedCollectionPageIntermediateType(in interface{}) (t *ccOrderedCollectionPageIntermediateType, err error) {
	tmp := &ccOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice ccOrderedCollectionPageIntermediateType will accept a slice to create a slice of ccOrderedCollectionPageIntermediateType
func deserializeSliceCcOrderedCollectionPageIntermediateType(in []interface{}) (t []*ccOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &ccOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeccOrderedCollectionPageIntermediateType will accept a ccOrderedCollectionPageIntermediateType to create a map
func serializeCcOrderedCollectionPageIntermediateType(t *ccOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceccOrderedCollectionPageIntermediateType will accept a slice of ccOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceCcOrderedCollectionPageIntermediateType(s []*ccOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializebccOrderedCollectionPageIntermediateType will accept a map to create a bccOrderedCollectionPageIntermediateType
func deserializeBccOrderedCollectionPageIntermediateType(in interface{}) (t *bccOrderedCollectionPageIntermediateType, err error) {
	tmp := &bccOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice bccOrderedCollectionPageIntermediateType will accept a slice to create a slice of bccOrderedCollectionPageIntermediateType
func deserializeSliceBccOrderedCollectionPageIntermediateType(in []interface{}) (t []*bccOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &bccOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializebccOrderedCollectionPageIntermediateType will accept a bccOrderedCollectionPageIntermediateType to create a map
func serializeBccOrderedCollectionPageIntermediateType(t *bccOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicebccOrderedCollectionPageIntermediateType will accept a slice of bccOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceBccOrderedCollectionPageIntermediateType(s []*bccOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializemediaTypeOrderedCollectionPageIntermediateType will accept a map to create a mediaTypeOrderedCollectionPageIntermediateType
func deserializeMediaTypeOrderedCollectionPageIntermediateType(in interface{}) (t *mediaTypeOrderedCollectionPageIntermediateType, err error) {
	tmp := &mediaTypeOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice mediaTypeOrderedCollectionPageIntermediateType will accept a slice to create a slice of mediaTypeOrderedCollectionPageIntermediateType
func deserializeSliceMediaTypeOrderedCollectionPageIntermediateType(in []interface{}) (t []*mediaTypeOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &mediaTypeOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializemediaTypeOrderedCollectionPageIntermediateType will accept a mediaTypeOrderedCollectionPageIntermediateType to create a map
func serializeMediaTypeOrderedCollectionPageIntermediateType(t *mediaTypeOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicemediaTypeOrderedCollectionPageIntermediateType will accept a slice of mediaTypeOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceMediaTypeOrderedCollectionPageIntermediateType(s []*mediaTypeOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializedurationOrderedCollectionPageIntermediateType will accept a map to create a durationOrderedCollectionPageIntermediateType
func deserializeDurationOrderedCollectionPageIntermediateType(in interface{}) (t *durationOrderedCollectionPageIntermediateType, err error) {
	tmp := &durationOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice durationOrderedCollectionPageIntermediateType will accept a slice to create a slice of durationOrderedCollectionPageIntermediateType
func deserializeSliceDurationOrderedCollectionPageIntermediateType(in []interface{}) (t []*durationOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &durationOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializedurationOrderedCollectionPageIntermediateType will accept a durationOrderedCollectionPageIntermediateType to create a map
func serializeDurationOrderedCollectionPageIntermediateType(t *durationOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicedurationOrderedCollectionPageIntermediateType will accept a slice of durationOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceDurationOrderedCollectionPageIntermediateType(s []*durationOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializesourceOrderedCollectionPageIntermediateType will accept a map to create a sourceOrderedCollectionPageIntermediateType
func deserializeSourceOrderedCollectionPageIntermediateType(in interface{}) (t *sourceOrderedCollectionPageIntermediateType, err error) {
	tmp := &sourceOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice sourceOrderedCollectionPageIntermediateType will accept a slice to create a slice of sourceOrderedCollectionPageIntermediateType
func deserializeSliceSourceOrderedCollectionPageIntermediateType(in []interface{}) (t []*sourceOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &sourceOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializesourceOrderedCollectionPageIntermediateType will accept a sourceOrderedCollectionPageIntermediateType to create a map
func serializeSourceOrderedCollectionPageIntermediateType(t *sourceOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicesourceOrderedCollectionPageIntermediateType will accept a slice of sourceOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceSourceOrderedCollectionPageIntermediateType(s []*sourceOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeinboxOrderedCollectionPageIntermediateType will accept a map to create a inboxOrderedCollectionPageIntermediateType
func deserializeInboxOrderedCollectionPageIntermediateType(in interface{}) (t *inboxOrderedCollectionPageIntermediateType, err error) {
	tmp := &inboxOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice inboxOrderedCollectionPageIntermediateType will accept a slice to create a slice of inboxOrderedCollectionPageIntermediateType
func deserializeSliceInboxOrderedCollectionPageIntermediateType(in []interface{}) (t []*inboxOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &inboxOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeinboxOrderedCollectionPageIntermediateType will accept a inboxOrderedCollectionPageIntermediateType to create a map
func serializeInboxOrderedCollectionPageIntermediateType(t *inboxOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceinboxOrderedCollectionPageIntermediateType will accept a slice of inboxOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceInboxOrderedCollectionPageIntermediateType(s []*inboxOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeoutboxOrderedCollectionPageIntermediateType will accept a map to create a outboxOrderedCollectionPageIntermediateType
func deserializeOutboxOrderedCollectionPageIntermediateType(in interface{}) (t *outboxOrderedCollectionPageIntermediateType, err error) {
	tmp := &outboxOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice outboxOrderedCollectionPageIntermediateType will accept a slice to create a slice of outboxOrderedCollectionPageIntermediateType
func deserializeSliceOutboxOrderedCollectionPageIntermediateType(in []interface{}) (t []*outboxOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &outboxOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeoutboxOrderedCollectionPageIntermediateType will accept a outboxOrderedCollectionPageIntermediateType to create a map
func serializeOutboxOrderedCollectionPageIntermediateType(t *outboxOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceoutboxOrderedCollectionPageIntermediateType will accept a slice of outboxOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceOutboxOrderedCollectionPageIntermediateType(s []*outboxOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializefollowingOrderedCollectionPageIntermediateType will accept a map to create a followingOrderedCollectionPageIntermediateType
func deserializeFollowingOrderedCollectionPageIntermediateType(in interface{}) (t *followingOrderedCollectionPageIntermediateType, err error) {
	tmp := &followingOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice followingOrderedCollectionPageIntermediateType will accept a slice to create a slice of followingOrderedCollectionPageIntermediateType
func deserializeSliceFollowingOrderedCollectionPageIntermediateType(in []interface{}) (t []*followingOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &followingOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializefollowingOrderedCollectionPageIntermediateType will accept a followingOrderedCollectionPageIntermediateType to create a map
func serializeFollowingOrderedCollectionPageIntermediateType(t *followingOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicefollowingOrderedCollectionPageIntermediateType will accept a slice of followingOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceFollowingOrderedCollectionPageIntermediateType(s []*followingOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializefollowersOrderedCollectionPageIntermediateType will accept a map to create a followersOrderedCollectionPageIntermediateType
func deserializeFollowersOrderedCollectionPageIntermediateType(in interface{}) (t *followersOrderedCollectionPageIntermediateType, err error) {
	tmp := &followersOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice followersOrderedCollectionPageIntermediateType will accept a slice to create a slice of followersOrderedCollectionPageIntermediateType
func deserializeSliceFollowersOrderedCollectionPageIntermediateType(in []interface{}) (t []*followersOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &followersOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializefollowersOrderedCollectionPageIntermediateType will accept a followersOrderedCollectionPageIntermediateType to create a map
func serializeFollowersOrderedCollectionPageIntermediateType(t *followersOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicefollowersOrderedCollectionPageIntermediateType will accept a slice of followersOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceFollowersOrderedCollectionPageIntermediateType(s []*followersOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializelikedOrderedCollectionPageIntermediateType will accept a map to create a likedOrderedCollectionPageIntermediateType
func deserializeLikedOrderedCollectionPageIntermediateType(in interface{}) (t *likedOrderedCollectionPageIntermediateType, err error) {
	tmp := &likedOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice likedOrderedCollectionPageIntermediateType will accept a slice to create a slice of likedOrderedCollectionPageIntermediateType
func deserializeSliceLikedOrderedCollectionPageIntermediateType(in []interface{}) (t []*likedOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &likedOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializelikedOrderedCollectionPageIntermediateType will accept a likedOrderedCollectionPageIntermediateType to create a map
func serializeLikedOrderedCollectionPageIntermediateType(t *likedOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicelikedOrderedCollectionPageIntermediateType will accept a slice of likedOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceLikedOrderedCollectionPageIntermediateType(s []*likedOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializelikesOrderedCollectionPageIntermediateType will accept a map to create a likesOrderedCollectionPageIntermediateType
func deserializeLikesOrderedCollectionPageIntermediateType(in interface{}) (t *likesOrderedCollectionPageIntermediateType, err error) {
	tmp := &likesOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice likesOrderedCollectionPageIntermediateType will accept a slice to create a slice of likesOrderedCollectionPageIntermediateType
func deserializeSliceLikesOrderedCollectionPageIntermediateType(in []interface{}) (t []*likesOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &likesOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializelikesOrderedCollectionPageIntermediateType will accept a likesOrderedCollectionPageIntermediateType to create a map
func serializeLikesOrderedCollectionPageIntermediateType(t *likesOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicelikesOrderedCollectionPageIntermediateType will accept a slice of likesOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceLikesOrderedCollectionPageIntermediateType(s []*likesOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializepreferredUsernameOrderedCollectionPageIntermediateType will accept a map to create a preferredUsernameOrderedCollectionPageIntermediateType
func deserializePreferredUsernameOrderedCollectionPageIntermediateType(in interface{}) (t *preferredUsernameOrderedCollectionPageIntermediateType, err error) {
	tmp := &preferredUsernameOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice preferredUsernameOrderedCollectionPageIntermediateType will accept a slice to create a slice of preferredUsernameOrderedCollectionPageIntermediateType
func deserializeSlicePreferredUsernameOrderedCollectionPageIntermediateType(in []interface{}) (t []*preferredUsernameOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &preferredUsernameOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializepreferredUsernameOrderedCollectionPageIntermediateType will accept a preferredUsernameOrderedCollectionPageIntermediateType to create a map
func serializePreferredUsernameOrderedCollectionPageIntermediateType(t *preferredUsernameOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicepreferredUsernameOrderedCollectionPageIntermediateType will accept a slice of preferredUsernameOrderedCollectionPageIntermediateType to create a slice result
func serializeSlicePreferredUsernameOrderedCollectionPageIntermediateType(s []*preferredUsernameOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializeendpointsOrderedCollectionPageIntermediateType will accept a map to create a endpointsOrderedCollectionPageIntermediateType
func deserializeEndpointsOrderedCollectionPageIntermediateType(in interface{}) (t *endpointsOrderedCollectionPageIntermediateType, err error) {
	tmp := &endpointsOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice endpointsOrderedCollectionPageIntermediateType will accept a slice to create a slice of endpointsOrderedCollectionPageIntermediateType
func deserializeSliceEndpointsOrderedCollectionPageIntermediateType(in []interface{}) (t []*endpointsOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &endpointsOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializeendpointsOrderedCollectionPageIntermediateType will accept a endpointsOrderedCollectionPageIntermediateType to create a map
func serializeEndpointsOrderedCollectionPageIntermediateType(t *endpointsOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSliceendpointsOrderedCollectionPageIntermediateType will accept a slice of endpointsOrderedCollectionPageIntermediateType to create a slice result
func serializeSliceEndpointsOrderedCollectionPageIntermediateType(s []*endpointsOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

// deserializepartOfOrderedCollectionPageIntermediateType will accept a map to create a partOfOrderedCollectionPageIntermediateType
func deserializePartOfOrderedCollectionPageIntermediateType(in interface{}) (t *partOfOrderedCollectionPageIntermediateType, err error) {
	tmp := &partOfOrderedCollectionPageIntermediateType{}
	err = tmp.Deserialize(in)
	return tmp, err

}

// deserializeSlice partOfOrderedCollectionPageIntermediateType will accept a slice to create a slice of partOfOrderedCollectionPageIntermediateType
func deserializeSlicePartOfOrderedCollectionPageIntermediateType(in []interface{}) (t []*partOfOrderedCollectionPageIntermediateType, err error) {
	for _, i := range in {
		tmp := &partOfOrderedCollectionPageIntermediateType{}
		err = tmp.Deserialize(i)
		if err != nil {
			return
		}
		t = append(t, tmp)
	}
	return

}

// serializepartOfOrderedCollectionPageIntermediateType will accept a partOfOrderedCollectionPageIntermediateType to create a map
func serializePartOfOrderedCollectionPageIntermediateType(t *partOfOrderedCollectionPageIntermediateType) (i interface{}, err error) {
	i, err = t.Serialize()
	return

}

// serializeSlicepartOfOrderedCollectionPageIntermediateType will accept a slice of partOfOrderedCollectionPageIntermediateType to create a slice result
func serializeSlicePartOfOrderedCollectionPageIntermediateType(s []*partOfOrderedCollectionPageIntermediateType) (out []interface{}, err error) {
	for _, t := range s {
		v, err := t.Serialize()
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return

}

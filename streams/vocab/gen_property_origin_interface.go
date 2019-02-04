package vocab

import "net/url"

// OriginPropertyIteratorInterface represents a single value for the "origin"
// property.
type OriginPropertyIteratorInterface interface {
	// GetAccept returns the value of this property. When IsAccept returns
	// false, GetAccept will return an arbitrary value.
	GetAccept() AcceptInterface
	// GetActivity returns the value of this property. When IsActivity returns
	// false, GetActivity will return an arbitrary value.
	GetActivity() ActivityInterface
	// GetAdd returns the value of this property. When IsAdd returns false,
	// GetAdd will return an arbitrary value.
	GetAdd() AddInterface
	// GetAnnounce returns the value of this property. When IsAnnounce returns
	// false, GetAnnounce will return an arbitrary value.
	GetAnnounce() AnnounceInterface
	// GetApplication returns the value of this property. When IsApplication
	// returns false, GetApplication will return an arbitrary value.
	GetApplication() ApplicationInterface
	// GetArrive returns the value of this property. When IsArrive returns
	// false, GetArrive will return an arbitrary value.
	GetArrive() ArriveInterface
	// GetArticle returns the value of this property. When IsArticle returns
	// false, GetArticle will return an arbitrary value.
	GetArticle() ArticleInterface
	// GetAudio returns the value of this property. When IsAudio returns
	// false, GetAudio will return an arbitrary value.
	GetAudio() AudioInterface
	// GetBlock returns the value of this property. When IsBlock returns
	// false, GetBlock will return an arbitrary value.
	GetBlock() BlockInterface
	// GetCollection returns the value of this property. When IsCollection
	// returns false, GetCollection will return an arbitrary value.
	GetCollection() CollectionInterface
	// GetCollectionPage returns the value of this property. When
	// IsCollectionPage returns false, GetCollectionPage will return an
	// arbitrary value.
	GetCollectionPage() CollectionPageInterface
	// GetCreate returns the value of this property. When IsCreate returns
	// false, GetCreate will return an arbitrary value.
	GetCreate() CreateInterface
	// GetDelete returns the value of this property. When IsDelete returns
	// false, GetDelete will return an arbitrary value.
	GetDelete() DeleteInterface
	// GetDislike returns the value of this property. When IsDislike returns
	// false, GetDislike will return an arbitrary value.
	GetDislike() DislikeInterface
	// GetDocument returns the value of this property. When IsDocument returns
	// false, GetDocument will return an arbitrary value.
	GetDocument() DocumentInterface
	// GetEvent returns the value of this property. When IsEvent returns
	// false, GetEvent will return an arbitrary value.
	GetEvent() EventInterface
	// GetFlag returns the value of this property. When IsFlag returns false,
	// GetFlag will return an arbitrary value.
	GetFlag() FlagInterface
	// GetFollow returns the value of this property. When IsFollow returns
	// false, GetFollow will return an arbitrary value.
	GetFollow() FollowInterface
	// GetGroup returns the value of this property. When IsGroup returns
	// false, GetGroup will return an arbitrary value.
	GetGroup() GroupInterface
	// GetIRI returns the IRI of this property. When IsIRI returns false,
	// GetIRI will return an arbitrary value.
	GetIRI() *url.URL
	// GetIgnore returns the value of this property. When IsIgnore returns
	// false, GetIgnore will return an arbitrary value.
	GetIgnore() IgnoreInterface
	// GetImage returns the value of this property. When IsImage returns
	// false, GetImage will return an arbitrary value.
	GetImage() ImageInterface
	// GetIntransitiveActivity returns the value of this property. When
	// IsIntransitiveActivity returns false, GetIntransitiveActivity will
	// return an arbitrary value.
	GetIntransitiveActivity() IntransitiveActivityInterface
	// GetInvite returns the value of this property. When IsInvite returns
	// false, GetInvite will return an arbitrary value.
	GetInvite() InviteInterface
	// GetJoin returns the value of this property. When IsJoin returns false,
	// GetJoin will return an arbitrary value.
	GetJoin() JoinInterface
	// GetLeave returns the value of this property. When IsLeave returns
	// false, GetLeave will return an arbitrary value.
	GetLeave() LeaveInterface
	// GetLike returns the value of this property. When IsLike returns false,
	// GetLike will return an arbitrary value.
	GetLike() LikeInterface
	// GetLink returns the value of this property. When IsLink returns false,
	// GetLink will return an arbitrary value.
	GetLink() LinkInterface
	// GetListen returns the value of this property. When IsListen returns
	// false, GetListen will return an arbitrary value.
	GetListen() ListenInterface
	// GetMention returns the value of this property. When IsMention returns
	// false, GetMention will return an arbitrary value.
	GetMention() MentionInterface
	// GetMove returns the value of this property. When IsMove returns false,
	// GetMove will return an arbitrary value.
	GetMove() MoveInterface
	// GetNote returns the value of this property. When IsNote returns false,
	// GetNote will return an arbitrary value.
	GetNote() NoteInterface
	// GetObject returns the value of this property. When IsObject returns
	// false, GetObject will return an arbitrary value.
	GetObject() ObjectInterface
	// GetOffer returns the value of this property. When IsOffer returns
	// false, GetOffer will return an arbitrary value.
	GetOffer() OfferInterface
	// GetOrderedCollection returns the value of this property. When
	// IsOrderedCollection returns false, GetOrderedCollection will return
	// an arbitrary value.
	GetOrderedCollection() OrderedCollectionInterface
	// GetOrderedCollectionPage returns the value of this property. When
	// IsOrderedCollectionPage returns false, GetOrderedCollectionPage
	// will return an arbitrary value.
	GetOrderedCollectionPage() OrderedCollectionPageInterface
	// GetOrganization returns the value of this property. When IsOrganization
	// returns false, GetOrganization will return an arbitrary value.
	GetOrganization() OrganizationInterface
	// GetPage returns the value of this property. When IsPage returns false,
	// GetPage will return an arbitrary value.
	GetPage() PageInterface
	// GetPerson returns the value of this property. When IsPerson returns
	// false, GetPerson will return an arbitrary value.
	GetPerson() PersonInterface
	// GetPlace returns the value of this property. When IsPlace returns
	// false, GetPlace will return an arbitrary value.
	GetPlace() PlaceInterface
	// GetProfile returns the value of this property. When IsProfile returns
	// false, GetProfile will return an arbitrary value.
	GetProfile() ProfileInterface
	// GetQuestion returns the value of this property. When IsQuestion returns
	// false, GetQuestion will return an arbitrary value.
	GetQuestion() QuestionInterface
	// GetRead returns the value of this property. When IsRead returns false,
	// GetRead will return an arbitrary value.
	GetRead() ReadInterface
	// GetReject returns the value of this property. When IsReject returns
	// false, GetReject will return an arbitrary value.
	GetReject() RejectInterface
	// GetRelationship returns the value of this property. When IsRelationship
	// returns false, GetRelationship will return an arbitrary value.
	GetRelationship() RelationshipInterface
	// GetRemove returns the value of this property. When IsRemove returns
	// false, GetRemove will return an arbitrary value.
	GetRemove() RemoveInterface
	// GetService returns the value of this property. When IsService returns
	// false, GetService will return an arbitrary value.
	GetService() ServiceInterface
	// GetTentativeAccept returns the value of this property. When
	// IsTentativeAccept returns false, GetTentativeAccept will return an
	// arbitrary value.
	GetTentativeAccept() TentativeAcceptInterface
	// GetTentativeReject returns the value of this property. When
	// IsTentativeReject returns false, GetTentativeReject will return an
	// arbitrary value.
	GetTentativeReject() TentativeRejectInterface
	// GetTombstone returns the value of this property. When IsTombstone
	// returns false, GetTombstone will return an arbitrary value.
	GetTombstone() TombstoneInterface
	// GetTravel returns the value of this property. When IsTravel returns
	// false, GetTravel will return an arbitrary value.
	GetTravel() TravelInterface
	// GetUndo returns the value of this property. When IsUndo returns false,
	// GetUndo will return an arbitrary value.
	GetUndo() UndoInterface
	// GetUpdate returns the value of this property. When IsUpdate returns
	// false, GetUpdate will return an arbitrary value.
	GetUpdate() UpdateInterface
	// GetVideo returns the value of this property. When IsVideo returns
	// false, GetVideo will return an arbitrary value.
	GetVideo() VideoInterface
	// GetView returns the value of this property. When IsView returns false,
	// GetView will return an arbitrary value.
	GetView() ViewInterface
	// HasAny returns true if any of the different values is set.
	HasAny() bool
	// IsAccept returns true if this property has a type of "Accept". When
	// true, use the GetAccept and SetAccept methods to access and set
	// this property.
	IsAccept() bool
	// IsActivity returns true if this property has a type of "Activity". When
	// true, use the GetActivity and SetActivity methods to access and set
	// this property.
	IsActivity() bool
	// IsAdd returns true if this property has a type of "Add". When true, use
	// the GetAdd and SetAdd methods to access and set this property.
	IsAdd() bool
	// IsAnnounce returns true if this property has a type of "Announce". When
	// true, use the GetAnnounce and SetAnnounce methods to access and set
	// this property.
	IsAnnounce() bool
	// IsApplication returns true if this property has a type of
	// "Application". When true, use the GetApplication and SetApplication
	// methods to access and set this property.
	IsApplication() bool
	// IsArrive returns true if this property has a type of "Arrive". When
	// true, use the GetArrive and SetArrive methods to access and set
	// this property.
	IsArrive() bool
	// IsArticle returns true if this property has a type of "Article". When
	// true, use the GetArticle and SetArticle methods to access and set
	// this property.
	IsArticle() bool
	// IsAudio returns true if this property has a type of "Audio". When true,
	// use the GetAudio and SetAudio methods to access and set this
	// property.
	IsAudio() bool
	// IsBlock returns true if this property has a type of "Block". When true,
	// use the GetBlock and SetBlock methods to access and set this
	// property.
	IsBlock() bool
	// IsCollection returns true if this property has a type of "Collection".
	// When true, use the GetCollection and SetCollection methods to
	// access and set this property.
	IsCollection() bool
	// IsCollectionPage returns true if this property has a type of
	// "CollectionPage". When true, use the GetCollectionPage and
	// SetCollectionPage methods to access and set this property.
	IsCollectionPage() bool
	// IsCreate returns true if this property has a type of "Create". When
	// true, use the GetCreate and SetCreate methods to access and set
	// this property.
	IsCreate() bool
	// IsDelete returns true if this property has a type of "Delete". When
	// true, use the GetDelete and SetDelete methods to access and set
	// this property.
	IsDelete() bool
	// IsDislike returns true if this property has a type of "Dislike". When
	// true, use the GetDislike and SetDislike methods to access and set
	// this property.
	IsDislike() bool
	// IsDocument returns true if this property has a type of "Document". When
	// true, use the GetDocument and SetDocument methods to access and set
	// this property.
	IsDocument() bool
	// IsEvent returns true if this property has a type of "Event". When true,
	// use the GetEvent and SetEvent methods to access and set this
	// property.
	IsEvent() bool
	// IsFlag returns true if this property has a type of "Flag". When true,
	// use the GetFlag and SetFlag methods to access and set this property.
	IsFlag() bool
	// IsFollow returns true if this property has a type of "Follow". When
	// true, use the GetFollow and SetFollow methods to access and set
	// this property.
	IsFollow() bool
	// IsGroup returns true if this property has a type of "Group". When true,
	// use the GetGroup and SetGroup methods to access and set this
	// property.
	IsGroup() bool
	// IsIRI returns true if this property is an IRI. When true, use GetIRI
	// and SetIRI to access and set this property
	IsIRI() bool
	// IsIgnore returns true if this property has a type of "Ignore". When
	// true, use the GetIgnore and SetIgnore methods to access and set
	// this property.
	IsIgnore() bool
	// IsImage returns true if this property has a type of "Image". When true,
	// use the GetImage and SetImage methods to access and set this
	// property.
	IsImage() bool
	// IsIntransitiveActivity returns true if this property has a type of
	// "IntransitiveActivity". When true, use the GetIntransitiveActivity
	// and SetIntransitiveActivity methods to access and set this property.
	IsIntransitiveActivity() bool
	// IsInvite returns true if this property has a type of "Invite". When
	// true, use the GetInvite and SetInvite methods to access and set
	// this property.
	IsInvite() bool
	// IsJoin returns true if this property has a type of "Join". When true,
	// use the GetJoin and SetJoin methods to access and set this property.
	IsJoin() bool
	// IsLeave returns true if this property has a type of "Leave". When true,
	// use the GetLeave and SetLeave methods to access and set this
	// property.
	IsLeave() bool
	// IsLike returns true if this property has a type of "Like". When true,
	// use the GetLike and SetLike methods to access and set this property.
	IsLike() bool
	// IsLink returns true if this property has a type of "Link". When true,
	// use the GetLink and SetLink methods to access and set this property.
	IsLink() bool
	// IsListen returns true if this property has a type of "Listen". When
	// true, use the GetListen and SetListen methods to access and set
	// this property.
	IsListen() bool
	// IsMention returns true if this property has a type of "Mention". When
	// true, use the GetMention and SetMention methods to access and set
	// this property.
	IsMention() bool
	// IsMove returns true if this property has a type of "Move". When true,
	// use the GetMove and SetMove methods to access and set this property.
	IsMove() bool
	// IsNote returns true if this property has a type of "Note". When true,
	// use the GetNote and SetNote methods to access and set this property.
	IsNote() bool
	// IsObject returns true if this property has a type of "Object". When
	// true, use the GetObject and SetObject methods to access and set
	// this property.
	IsObject() bool
	// IsOffer returns true if this property has a type of "Offer". When true,
	// use the GetOffer and SetOffer methods to access and set this
	// property.
	IsOffer() bool
	// IsOrderedCollection returns true if this property has a type of
	// "OrderedCollection". When true, use the GetOrderedCollection and
	// SetOrderedCollection methods to access and set this property.
	IsOrderedCollection() bool
	// IsOrderedCollectionPage returns true if this property has a type of
	// "OrderedCollectionPage". When true, use the
	// GetOrderedCollectionPage and SetOrderedCollectionPage methods to
	// access and set this property.
	IsOrderedCollectionPage() bool
	// IsOrganization returns true if this property has a type of
	// "Organization". When true, use the GetOrganization and
	// SetOrganization methods to access and set this property.
	IsOrganization() bool
	// IsPage returns true if this property has a type of "Page". When true,
	// use the GetPage and SetPage methods to access and set this property.
	IsPage() bool
	// IsPerson returns true if this property has a type of "Person". When
	// true, use the GetPerson and SetPerson methods to access and set
	// this property.
	IsPerson() bool
	// IsPlace returns true if this property has a type of "Place". When true,
	// use the GetPlace and SetPlace methods to access and set this
	// property.
	IsPlace() bool
	// IsProfile returns true if this property has a type of "Profile". When
	// true, use the GetProfile and SetProfile methods to access and set
	// this property.
	IsProfile() bool
	// IsQuestion returns true if this property has a type of "Question". When
	// true, use the GetQuestion and SetQuestion methods to access and set
	// this property.
	IsQuestion() bool
	// IsRead returns true if this property has a type of "Read". When true,
	// use the GetRead and SetRead methods to access and set this property.
	IsRead() bool
	// IsReject returns true if this property has a type of "Reject". When
	// true, use the GetReject and SetReject methods to access and set
	// this property.
	IsReject() bool
	// IsRelationship returns true if this property has a type of
	// "Relationship". When true, use the GetRelationship and
	// SetRelationship methods to access and set this property.
	IsRelationship() bool
	// IsRemove returns true if this property has a type of "Remove". When
	// true, use the GetRemove and SetRemove methods to access and set
	// this property.
	IsRemove() bool
	// IsService returns true if this property has a type of "Service". When
	// true, use the GetService and SetService methods to access and set
	// this property.
	IsService() bool
	// IsTentativeAccept returns true if this property has a type of
	// "TentativeAccept". When true, use the GetTentativeAccept and
	// SetTentativeAccept methods to access and set this property.
	IsTentativeAccept() bool
	// IsTentativeReject returns true if this property has a type of
	// "TentativeReject". When true, use the GetTentativeReject and
	// SetTentativeReject methods to access and set this property.
	IsTentativeReject() bool
	// IsTombstone returns true if this property has a type of "Tombstone".
	// When true, use the GetTombstone and SetTombstone methods to access
	// and set this property.
	IsTombstone() bool
	// IsTravel returns true if this property has a type of "Travel". When
	// true, use the GetTravel and SetTravel methods to access and set
	// this property.
	IsTravel() bool
	// IsUndo returns true if this property has a type of "Undo". When true,
	// use the GetUndo and SetUndo methods to access and set this property.
	IsUndo() bool
	// IsUpdate returns true if this property has a type of "Update". When
	// true, use the GetUpdate and SetUpdate methods to access and set
	// this property.
	IsUpdate() bool
	// IsVideo returns true if this property has a type of "Video". When true,
	// use the GetVideo and SetVideo methods to access and set this
	// property.
	IsVideo() bool
	// IsView returns true if this property has a type of "View". When true,
	// use the GetView and SetView methods to access and set this property.
	IsView() bool
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this property and the specific values that are set. The value
	// in the map is the alias used to import the property's value or
	// values.
	JSONLDContext() map[string]string
	// KindIndex computes an arbitrary value for indexing this kind of value.
	// This is a leaky API detail only for folks looking to replace the
	// go-fed implementation. Applications should not use this method.
	KindIndex() int
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o OriginPropertyIteratorInterface) bool
	// Name returns the name of this property: "origin".
	Name() string
	// Next returns the next iterator, or nil if there is no next iterator.
	Next() OriginPropertyIteratorInterface
	// Prev returns the previous iterator, or nil if there is no previous
	// iterator.
	Prev() OriginPropertyIteratorInterface
	// SetAccept sets the value of this property. Calling IsAccept afterwards
	// returns true.
	SetAccept(v AcceptInterface)
	// SetActivity sets the value of this property. Calling IsActivity
	// afterwards returns true.
	SetActivity(v ActivityInterface)
	// SetAdd sets the value of this property. Calling IsAdd afterwards
	// returns true.
	SetAdd(v AddInterface)
	// SetAnnounce sets the value of this property. Calling IsAnnounce
	// afterwards returns true.
	SetAnnounce(v AnnounceInterface)
	// SetApplication sets the value of this property. Calling IsApplication
	// afterwards returns true.
	SetApplication(v ApplicationInterface)
	// SetArrive sets the value of this property. Calling IsArrive afterwards
	// returns true.
	SetArrive(v ArriveInterface)
	// SetArticle sets the value of this property. Calling IsArticle
	// afterwards returns true.
	SetArticle(v ArticleInterface)
	// SetAudio sets the value of this property. Calling IsAudio afterwards
	// returns true.
	SetAudio(v AudioInterface)
	// SetBlock sets the value of this property. Calling IsBlock afterwards
	// returns true.
	SetBlock(v BlockInterface)
	// SetCollection sets the value of this property. Calling IsCollection
	// afterwards returns true.
	SetCollection(v CollectionInterface)
	// SetCollectionPage sets the value of this property. Calling
	// IsCollectionPage afterwards returns true.
	SetCollectionPage(v CollectionPageInterface)
	// SetCreate sets the value of this property. Calling IsCreate afterwards
	// returns true.
	SetCreate(v CreateInterface)
	// SetDelete sets the value of this property. Calling IsDelete afterwards
	// returns true.
	SetDelete(v DeleteInterface)
	// SetDislike sets the value of this property. Calling IsDislike
	// afterwards returns true.
	SetDislike(v DislikeInterface)
	// SetDocument sets the value of this property. Calling IsDocument
	// afterwards returns true.
	SetDocument(v DocumentInterface)
	// SetEvent sets the value of this property. Calling IsEvent afterwards
	// returns true.
	SetEvent(v EventInterface)
	// SetFlag sets the value of this property. Calling IsFlag afterwards
	// returns true.
	SetFlag(v FlagInterface)
	// SetFollow sets the value of this property. Calling IsFollow afterwards
	// returns true.
	SetFollow(v FollowInterface)
	// SetGroup sets the value of this property. Calling IsGroup afterwards
	// returns true.
	SetGroup(v GroupInterface)
	// SetIRI sets the value of this property. Calling IsIRI afterwards
	// returns true.
	SetIRI(v *url.URL)
	// SetIgnore sets the value of this property. Calling IsIgnore afterwards
	// returns true.
	SetIgnore(v IgnoreInterface)
	// SetImage sets the value of this property. Calling IsImage afterwards
	// returns true.
	SetImage(v ImageInterface)
	// SetIntransitiveActivity sets the value of this property. Calling
	// IsIntransitiveActivity afterwards returns true.
	SetIntransitiveActivity(v IntransitiveActivityInterface)
	// SetInvite sets the value of this property. Calling IsInvite afterwards
	// returns true.
	SetInvite(v InviteInterface)
	// SetJoin sets the value of this property. Calling IsJoin afterwards
	// returns true.
	SetJoin(v JoinInterface)
	// SetLeave sets the value of this property. Calling IsLeave afterwards
	// returns true.
	SetLeave(v LeaveInterface)
	// SetLike sets the value of this property. Calling IsLike afterwards
	// returns true.
	SetLike(v LikeInterface)
	// SetLink sets the value of this property. Calling IsLink afterwards
	// returns true.
	SetLink(v LinkInterface)
	// SetListen sets the value of this property. Calling IsListen afterwards
	// returns true.
	SetListen(v ListenInterface)
	// SetMention sets the value of this property. Calling IsMention
	// afterwards returns true.
	SetMention(v MentionInterface)
	// SetMove sets the value of this property. Calling IsMove afterwards
	// returns true.
	SetMove(v MoveInterface)
	// SetNote sets the value of this property. Calling IsNote afterwards
	// returns true.
	SetNote(v NoteInterface)
	// SetObject sets the value of this property. Calling IsObject afterwards
	// returns true.
	SetObject(v ObjectInterface)
	// SetOffer sets the value of this property. Calling IsOffer afterwards
	// returns true.
	SetOffer(v OfferInterface)
	// SetOrderedCollection sets the value of this property. Calling
	// IsOrderedCollection afterwards returns true.
	SetOrderedCollection(v OrderedCollectionInterface)
	// SetOrderedCollectionPage sets the value of this property. Calling
	// IsOrderedCollectionPage afterwards returns true.
	SetOrderedCollectionPage(v OrderedCollectionPageInterface)
	// SetOrganization sets the value of this property. Calling IsOrganization
	// afterwards returns true.
	SetOrganization(v OrganizationInterface)
	// SetPage sets the value of this property. Calling IsPage afterwards
	// returns true.
	SetPage(v PageInterface)
	// SetPerson sets the value of this property. Calling IsPerson afterwards
	// returns true.
	SetPerson(v PersonInterface)
	// SetPlace sets the value of this property. Calling IsPlace afterwards
	// returns true.
	SetPlace(v PlaceInterface)
	// SetProfile sets the value of this property. Calling IsProfile
	// afterwards returns true.
	SetProfile(v ProfileInterface)
	// SetQuestion sets the value of this property. Calling IsQuestion
	// afterwards returns true.
	SetQuestion(v QuestionInterface)
	// SetRead sets the value of this property. Calling IsRead afterwards
	// returns true.
	SetRead(v ReadInterface)
	// SetReject sets the value of this property. Calling IsReject afterwards
	// returns true.
	SetReject(v RejectInterface)
	// SetRelationship sets the value of this property. Calling IsRelationship
	// afterwards returns true.
	SetRelationship(v RelationshipInterface)
	// SetRemove sets the value of this property. Calling IsRemove afterwards
	// returns true.
	SetRemove(v RemoveInterface)
	// SetService sets the value of this property. Calling IsService
	// afterwards returns true.
	SetService(v ServiceInterface)
	// SetTentativeAccept sets the value of this property. Calling
	// IsTentativeAccept afterwards returns true.
	SetTentativeAccept(v TentativeAcceptInterface)
	// SetTentativeReject sets the value of this property. Calling
	// IsTentativeReject afterwards returns true.
	SetTentativeReject(v TentativeRejectInterface)
	// SetTombstone sets the value of this property. Calling IsTombstone
	// afterwards returns true.
	SetTombstone(v TombstoneInterface)
	// SetTravel sets the value of this property. Calling IsTravel afterwards
	// returns true.
	SetTravel(v TravelInterface)
	// SetUndo sets the value of this property. Calling IsUndo afterwards
	// returns true.
	SetUndo(v UndoInterface)
	// SetUpdate sets the value of this property. Calling IsUpdate afterwards
	// returns true.
	SetUpdate(v UpdateInterface)
	// SetVideo sets the value of this property. Calling IsVideo afterwards
	// returns true.
	SetVideo(v VideoInterface)
	// SetView sets the value of this property. Calling IsView afterwards
	// returns true.
	SetView(v ViewInterface)
}

// Describes an indirect object of the activity from which the activity is
// directed. The precise meaning of the origin is the object of the English
// preposition "from". For instance, in the activity "John moved an item to
// List B from List A", the origin of the activity is "List A".
//
// Example 94 (https://www.w3.org/TR/activitystreams-vocabulary/#ex166-jsonld):
//   {
//     "actor": "http://sally.example.org",
//     "object": "http://example.org/posts/1",
//     "origin": {
//       "name": "List A",
//       "type": "Collection"
//     },
//     "summary": "Sally moved a post from List A to List B",
//     "target": {
//       "name": "List B",
//       "type": "Collection"
//     },
//     "type": "Move"
//   }
type OriginPropertyInterface interface {
	// AppendAccept appends a Accept value to the back of a list of the
	// property "origin". Invalidates iterators that are traversing using
	// Prev.
	AppendAccept(v AcceptInterface)
	// AppendActivity appends a Activity value to the back of a list of the
	// property "origin". Invalidates iterators that are traversing using
	// Prev.
	AppendActivity(v ActivityInterface)
	// AppendAdd appends a Add value to the back of a list of the property
	// "origin". Invalidates iterators that are traversing using Prev.
	AppendAdd(v AddInterface)
	// AppendAnnounce appends a Announce value to the back of a list of the
	// property "origin". Invalidates iterators that are traversing using
	// Prev.
	AppendAnnounce(v AnnounceInterface)
	// AppendApplication appends a Application value to the back of a list of
	// the property "origin". Invalidates iterators that are traversing
	// using Prev.
	AppendApplication(v ApplicationInterface)
	// AppendArrive appends a Arrive value to the back of a list of the
	// property "origin". Invalidates iterators that are traversing using
	// Prev.
	AppendArrive(v ArriveInterface)
	// AppendArticle appends a Article value to the back of a list of the
	// property "origin". Invalidates iterators that are traversing using
	// Prev.
	AppendArticle(v ArticleInterface)
	// AppendAudio appends a Audio value to the back of a list of the property
	// "origin". Invalidates iterators that are traversing using Prev.
	AppendAudio(v AudioInterface)
	// AppendBlock appends a Block value to the back of a list of the property
	// "origin". Invalidates iterators that are traversing using Prev.
	AppendBlock(v BlockInterface)
	// AppendCollection appends a Collection value to the back of a list of
	// the property "origin". Invalidates iterators that are traversing
	// using Prev.
	AppendCollection(v CollectionInterface)
	// AppendCollectionPage appends a CollectionPage value to the back of a
	// list of the property "origin". Invalidates iterators that are
	// traversing using Prev.
	AppendCollectionPage(v CollectionPageInterface)
	// AppendCreate appends a Create value to the back of a list of the
	// property "origin". Invalidates iterators that are traversing using
	// Prev.
	AppendCreate(v CreateInterface)
	// AppendDelete appends a Delete value to the back of a list of the
	// property "origin". Invalidates iterators that are traversing using
	// Prev.
	AppendDelete(v DeleteInterface)
	// AppendDislike appends a Dislike value to the back of a list of the
	// property "origin". Invalidates iterators that are traversing using
	// Prev.
	AppendDislike(v DislikeInterface)
	// AppendDocument appends a Document value to the back of a list of the
	// property "origin". Invalidates iterators that are traversing using
	// Prev.
	AppendDocument(v DocumentInterface)
	// AppendEvent appends a Event value to the back of a list of the property
	// "origin". Invalidates iterators that are traversing using Prev.
	AppendEvent(v EventInterface)
	// AppendFlag appends a Flag value to the back of a list of the property
	// "origin". Invalidates iterators that are traversing using Prev.
	AppendFlag(v FlagInterface)
	// AppendFollow appends a Follow value to the back of a list of the
	// property "origin". Invalidates iterators that are traversing using
	// Prev.
	AppendFollow(v FollowInterface)
	// AppendGroup appends a Group value to the back of a list of the property
	// "origin". Invalidates iterators that are traversing using Prev.
	AppendGroup(v GroupInterface)
	// AppendIRI appends an IRI value to the back of a list of the property
	// "origin"
	AppendIRI(v *url.URL)
	// AppendIgnore appends a Ignore value to the back of a list of the
	// property "origin". Invalidates iterators that are traversing using
	// Prev.
	AppendIgnore(v IgnoreInterface)
	// AppendImage appends a Image value to the back of a list of the property
	// "origin". Invalidates iterators that are traversing using Prev.
	AppendImage(v ImageInterface)
	// AppendIntransitiveActivity appends a IntransitiveActivity value to the
	// back of a list of the property "origin". Invalidates iterators that
	// are traversing using Prev.
	AppendIntransitiveActivity(v IntransitiveActivityInterface)
	// AppendInvite appends a Invite value to the back of a list of the
	// property "origin". Invalidates iterators that are traversing using
	// Prev.
	AppendInvite(v InviteInterface)
	// AppendJoin appends a Join value to the back of a list of the property
	// "origin". Invalidates iterators that are traversing using Prev.
	AppendJoin(v JoinInterface)
	// AppendLeave appends a Leave value to the back of a list of the property
	// "origin". Invalidates iterators that are traversing using Prev.
	AppendLeave(v LeaveInterface)
	// AppendLike appends a Like value to the back of a list of the property
	// "origin". Invalidates iterators that are traversing using Prev.
	AppendLike(v LikeInterface)
	// AppendLink appends a Link value to the back of a list of the property
	// "origin". Invalidates iterators that are traversing using Prev.
	AppendLink(v LinkInterface)
	// AppendListen appends a Listen value to the back of a list of the
	// property "origin". Invalidates iterators that are traversing using
	// Prev.
	AppendListen(v ListenInterface)
	// AppendMention appends a Mention value to the back of a list of the
	// property "origin". Invalidates iterators that are traversing using
	// Prev.
	AppendMention(v MentionInterface)
	// AppendMove appends a Move value to the back of a list of the property
	// "origin". Invalidates iterators that are traversing using Prev.
	AppendMove(v MoveInterface)
	// AppendNote appends a Note value to the back of a list of the property
	// "origin". Invalidates iterators that are traversing using Prev.
	AppendNote(v NoteInterface)
	// AppendObject appends a Object value to the back of a list of the
	// property "origin". Invalidates iterators that are traversing using
	// Prev.
	AppendObject(v ObjectInterface)
	// AppendOffer appends a Offer value to the back of a list of the property
	// "origin". Invalidates iterators that are traversing using Prev.
	AppendOffer(v OfferInterface)
	// AppendOrderedCollection appends a OrderedCollection value to the back
	// of a list of the property "origin". Invalidates iterators that are
	// traversing using Prev.
	AppendOrderedCollection(v OrderedCollectionInterface)
	// AppendOrderedCollectionPage appends a OrderedCollectionPage value to
	// the back of a list of the property "origin". Invalidates iterators
	// that are traversing using Prev.
	AppendOrderedCollectionPage(v OrderedCollectionPageInterface)
	// AppendOrganization appends a Organization value to the back of a list
	// of the property "origin". Invalidates iterators that are traversing
	// using Prev.
	AppendOrganization(v OrganizationInterface)
	// AppendPage appends a Page value to the back of a list of the property
	// "origin". Invalidates iterators that are traversing using Prev.
	AppendPage(v PageInterface)
	// AppendPerson appends a Person value to the back of a list of the
	// property "origin". Invalidates iterators that are traversing using
	// Prev.
	AppendPerson(v PersonInterface)
	// AppendPlace appends a Place value to the back of a list of the property
	// "origin". Invalidates iterators that are traversing using Prev.
	AppendPlace(v PlaceInterface)
	// AppendProfile appends a Profile value to the back of a list of the
	// property "origin". Invalidates iterators that are traversing using
	// Prev.
	AppendProfile(v ProfileInterface)
	// AppendQuestion appends a Question value to the back of a list of the
	// property "origin". Invalidates iterators that are traversing using
	// Prev.
	AppendQuestion(v QuestionInterface)
	// AppendRead appends a Read value to the back of a list of the property
	// "origin". Invalidates iterators that are traversing using Prev.
	AppendRead(v ReadInterface)
	// AppendReject appends a Reject value to the back of a list of the
	// property "origin". Invalidates iterators that are traversing using
	// Prev.
	AppendReject(v RejectInterface)
	// AppendRelationship appends a Relationship value to the back of a list
	// of the property "origin". Invalidates iterators that are traversing
	// using Prev.
	AppendRelationship(v RelationshipInterface)
	// AppendRemove appends a Remove value to the back of a list of the
	// property "origin". Invalidates iterators that are traversing using
	// Prev.
	AppendRemove(v RemoveInterface)
	// AppendService appends a Service value to the back of a list of the
	// property "origin". Invalidates iterators that are traversing using
	// Prev.
	AppendService(v ServiceInterface)
	// AppendTentativeAccept appends a TentativeAccept value to the back of a
	// list of the property "origin". Invalidates iterators that are
	// traversing using Prev.
	AppendTentativeAccept(v TentativeAcceptInterface)
	// AppendTentativeReject appends a TentativeReject value to the back of a
	// list of the property "origin". Invalidates iterators that are
	// traversing using Prev.
	AppendTentativeReject(v TentativeRejectInterface)
	// AppendTombstone appends a Tombstone value to the back of a list of the
	// property "origin". Invalidates iterators that are traversing using
	// Prev.
	AppendTombstone(v TombstoneInterface)
	// AppendTravel appends a Travel value to the back of a list of the
	// property "origin". Invalidates iterators that are traversing using
	// Prev.
	AppendTravel(v TravelInterface)
	// AppendUndo appends a Undo value to the back of a list of the property
	// "origin". Invalidates iterators that are traversing using Prev.
	AppendUndo(v UndoInterface)
	// AppendUpdate appends a Update value to the back of a list of the
	// property "origin". Invalidates iterators that are traversing using
	// Prev.
	AppendUpdate(v UpdateInterface)
	// AppendVideo appends a Video value to the back of a list of the property
	// "origin". Invalidates iterators that are traversing using Prev.
	AppendVideo(v VideoInterface)
	// AppendView appends a View value to the back of a list of the property
	// "origin". Invalidates iterators that are traversing using Prev.
	AppendView(v ViewInterface)
	// At returns the property value for the specified index. Panics if the
	// index is out of bounds.
	At(index int) OriginPropertyIteratorInterface
	// Begin returns the first iterator, or nil if empty. Can be used with the
	// iterator's Next method and this property's End method to iterate
	// from front to back through all values.
	Begin() OriginPropertyIteratorInterface
	// Empty returns returns true if there are no elements.
	Empty() bool
	// End returns beyond-the-last iterator, which is nil. Can be used with
	// the iterator's Next method and this property's Begin method to
	// iterate from front to back through all values.
	End() OriginPropertyIteratorInterface
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this property and the specific values that are set. The value
	// in the map is the alias used to import the property's value or
	// values.
	JSONLDContext() map[string]string
	// KindIndex computes an arbitrary value for indexing this kind of value.
	// This is a leaky API method specifically needed only for alternate
	// implementations for go-fed. Applications should not use this
	// method. Panics if the index is out of bounds.
	KindIndex(idx int) int
	// Len returns the number of values that exist for the "origin" property.
	Len() (length int)
	// Less computes whether another property is less than this one. Mixing
	// types results in a consistent but arbitrary ordering
	Less(i, j int) bool
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o OriginPropertyInterface) bool
	// Name returns the name of this property: "origin".
	Name() string
	// PrependAccept prepends a Accept value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependAccept(v AcceptInterface)
	// PrependActivity prepends a Activity value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependActivity(v ActivityInterface)
	// PrependAdd prepends a Add value to the front of a list of the property
	// "origin". Invalidates all iterators.
	PrependAdd(v AddInterface)
	// PrependAnnounce prepends a Announce value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependAnnounce(v AnnounceInterface)
	// PrependApplication prepends a Application value to the front of a list
	// of the property "origin". Invalidates all iterators.
	PrependApplication(v ApplicationInterface)
	// PrependArrive prepends a Arrive value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependArrive(v ArriveInterface)
	// PrependArticle prepends a Article value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependArticle(v ArticleInterface)
	// PrependAudio prepends a Audio value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependAudio(v AudioInterface)
	// PrependBlock prepends a Block value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependBlock(v BlockInterface)
	// PrependCollection prepends a Collection value to the front of a list of
	// the property "origin". Invalidates all iterators.
	PrependCollection(v CollectionInterface)
	// PrependCollectionPage prepends a CollectionPage value to the front of a
	// list of the property "origin". Invalidates all iterators.
	PrependCollectionPage(v CollectionPageInterface)
	// PrependCreate prepends a Create value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependCreate(v CreateInterface)
	// PrependDelete prepends a Delete value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependDelete(v DeleteInterface)
	// PrependDislike prepends a Dislike value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependDislike(v DislikeInterface)
	// PrependDocument prepends a Document value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependDocument(v DocumentInterface)
	// PrependEvent prepends a Event value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependEvent(v EventInterface)
	// PrependFlag prepends a Flag value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependFlag(v FlagInterface)
	// PrependFollow prepends a Follow value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependFollow(v FollowInterface)
	// PrependGroup prepends a Group value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependGroup(v GroupInterface)
	// PrependIRI prepends an IRI value to the front of a list of the property
	// "origin".
	PrependIRI(v *url.URL)
	// PrependIgnore prepends a Ignore value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependIgnore(v IgnoreInterface)
	// PrependImage prepends a Image value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependImage(v ImageInterface)
	// PrependIntransitiveActivity prepends a IntransitiveActivity value to
	// the front of a list of the property "origin". Invalidates all
	// iterators.
	PrependIntransitiveActivity(v IntransitiveActivityInterface)
	// PrependInvite prepends a Invite value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependInvite(v InviteInterface)
	// PrependJoin prepends a Join value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependJoin(v JoinInterface)
	// PrependLeave prepends a Leave value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependLeave(v LeaveInterface)
	// PrependLike prepends a Like value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependLike(v LikeInterface)
	// PrependLink prepends a Link value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependLink(v LinkInterface)
	// PrependListen prepends a Listen value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependListen(v ListenInterface)
	// PrependMention prepends a Mention value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependMention(v MentionInterface)
	// PrependMove prepends a Move value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependMove(v MoveInterface)
	// PrependNote prepends a Note value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependNote(v NoteInterface)
	// PrependObject prepends a Object value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependObject(v ObjectInterface)
	// PrependOffer prepends a Offer value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependOffer(v OfferInterface)
	// PrependOrderedCollection prepends a OrderedCollection value to the
	// front of a list of the property "origin". Invalidates all iterators.
	PrependOrderedCollection(v OrderedCollectionInterface)
	// PrependOrderedCollectionPage prepends a OrderedCollectionPage value to
	// the front of a list of the property "origin". Invalidates all
	// iterators.
	PrependOrderedCollectionPage(v OrderedCollectionPageInterface)
	// PrependOrganization prepends a Organization value to the front of a
	// list of the property "origin". Invalidates all iterators.
	PrependOrganization(v OrganizationInterface)
	// PrependPage prepends a Page value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependPage(v PageInterface)
	// PrependPerson prepends a Person value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependPerson(v PersonInterface)
	// PrependPlace prepends a Place value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependPlace(v PlaceInterface)
	// PrependProfile prepends a Profile value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependProfile(v ProfileInterface)
	// PrependQuestion prepends a Question value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependQuestion(v QuestionInterface)
	// PrependRead prepends a Read value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependRead(v ReadInterface)
	// PrependReject prepends a Reject value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependReject(v RejectInterface)
	// PrependRelationship prepends a Relationship value to the front of a
	// list of the property "origin". Invalidates all iterators.
	PrependRelationship(v RelationshipInterface)
	// PrependRemove prepends a Remove value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependRemove(v RemoveInterface)
	// PrependService prepends a Service value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependService(v ServiceInterface)
	// PrependTentativeAccept prepends a TentativeAccept value to the front of
	// a list of the property "origin". Invalidates all iterators.
	PrependTentativeAccept(v TentativeAcceptInterface)
	// PrependTentativeReject prepends a TentativeReject value to the front of
	// a list of the property "origin". Invalidates all iterators.
	PrependTentativeReject(v TentativeRejectInterface)
	// PrependTombstone prepends a Tombstone value to the front of a list of
	// the property "origin". Invalidates all iterators.
	PrependTombstone(v TombstoneInterface)
	// PrependTravel prepends a Travel value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependTravel(v TravelInterface)
	// PrependUndo prepends a Undo value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependUndo(v UndoInterface)
	// PrependUpdate prepends a Update value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependUpdate(v UpdateInterface)
	// PrependVideo prepends a Video value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependVideo(v VideoInterface)
	// PrependView prepends a View value to the front of a list of the
	// property "origin". Invalidates all iterators.
	PrependView(v ViewInterface)
	// Remove deletes an element at the specified index from a list of the
	// property "origin", regardless of its type. Panics if the index is
	// out of bounds. Invalidates all iterators.
	Remove(idx int)
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format. Applications should not
	// need this function as most typical use cases serialize types
	// instead of individual properties. It is exposed for alternatives to
	// go-fed implementations to use.
	Serialize() (interface{}, error)
	// SetAccept sets a Accept value to be at the specified index for the
	// property "origin". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetAccept(idx int, v AcceptInterface)
	// SetActivity sets a Activity value to be at the specified index for the
	// property "origin". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetActivity(idx int, v ActivityInterface)
	// SetAdd sets a Add value to be at the specified index for the property
	// "origin". Panics if the index is out of bounds. Invalidates all
	// iterators.
	SetAdd(idx int, v AddInterface)
	// SetAnnounce sets a Announce value to be at the specified index for the
	// property "origin". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetAnnounce(idx int, v AnnounceInterface)
	// SetApplication sets a Application value to be at the specified index
	// for the property "origin". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetApplication(idx int, v ApplicationInterface)
	// SetArrive sets a Arrive value to be at the specified index for the
	// property "origin". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetArrive(idx int, v ArriveInterface)
	// SetArticle sets a Article value to be at the specified index for the
	// property "origin". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetArticle(idx int, v ArticleInterface)
	// SetAudio sets a Audio value to be at the specified index for the
	// property "origin". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetAudio(idx int, v AudioInterface)
	// SetBlock sets a Block value to be at the specified index for the
	// property "origin". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetBlock(idx int, v BlockInterface)
	// SetCollection sets a Collection value to be at the specified index for
	// the property "origin". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetCollection(idx int, v CollectionInterface)
	// SetCollectionPage sets a CollectionPage value to be at the specified
	// index for the property "origin". Panics if the index is out of
	// bounds. Invalidates all iterators.
	SetCollectionPage(idx int, v CollectionPageInterface)
	// SetCreate sets a Create value to be at the specified index for the
	// property "origin". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetCreate(idx int, v CreateInterface)
	// SetDelete sets a Delete value to be at the specified index for the
	// property "origin". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetDelete(idx int, v DeleteInterface)
	// SetDislike sets a Dislike value to be at the specified index for the
	// property "origin". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetDislike(idx int, v DislikeInterface)
	// SetDocument sets a Document value to be at the specified index for the
	// property "origin". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetDocument(idx int, v DocumentInterface)
	// SetEvent sets a Event value to be at the specified index for the
	// property "origin". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetEvent(idx int, v EventInterface)
	// SetFlag sets a Flag value to be at the specified index for the property
	// "origin". Panics if the index is out of bounds. Invalidates all
	// iterators.
	SetFlag(idx int, v FlagInterface)
	// SetFollow sets a Follow value to be at the specified index for the
	// property "origin". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetFollow(idx int, v FollowInterface)
	// SetGroup sets a Group value to be at the specified index for the
	// property "origin". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetGroup(idx int, v GroupInterface)
	// SetIRI sets an IRI value to be at the specified index for the property
	// "origin". Panics if the index is out of bounds.
	SetIRI(idx int, v *url.URL)
	// SetIgnore sets a Ignore value to be at the specified index for the
	// property "origin". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetIgnore(idx int, v IgnoreInterface)
	// SetImage sets a Image value to be at the specified index for the
	// property "origin". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetImage(idx int, v ImageInterface)
	// SetIntransitiveActivity sets a IntransitiveActivity value to be at the
	// specified index for the property "origin". Panics if the index is
	// out of bounds. Invalidates all iterators.
	SetIntransitiveActivity(idx int, v IntransitiveActivityInterface)
	// SetInvite sets a Invite value to be at the specified index for the
	// property "origin". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetInvite(idx int, v InviteInterface)
	// SetJoin sets a Join value to be at the specified index for the property
	// "origin". Panics if the index is out of bounds. Invalidates all
	// iterators.
	SetJoin(idx int, v JoinInterface)
	// SetLeave sets a Leave value to be at the specified index for the
	// property "origin". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetLeave(idx int, v LeaveInterface)
	// SetLike sets a Like value to be at the specified index for the property
	// "origin". Panics if the index is out of bounds. Invalidates all
	// iterators.
	SetLike(idx int, v LikeInterface)
	// SetLink sets a Link value to be at the specified index for the property
	// "origin". Panics if the index is out of bounds. Invalidates all
	// iterators.
	SetLink(idx int, v LinkInterface)
	// SetListen sets a Listen value to be at the specified index for the
	// property "origin". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetListen(idx int, v ListenInterface)
	// SetMention sets a Mention value to be at the specified index for the
	// property "origin". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetMention(idx int, v MentionInterface)
	// SetMove sets a Move value to be at the specified index for the property
	// "origin". Panics if the index is out of bounds. Invalidates all
	// iterators.
	SetMove(idx int, v MoveInterface)
	// SetNote sets a Note value to be at the specified index for the property
	// "origin". Panics if the index is out of bounds. Invalidates all
	// iterators.
	SetNote(idx int, v NoteInterface)
	// SetObject sets a Object value to be at the specified index for the
	// property "origin". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetObject(idx int, v ObjectInterface)
	// SetOffer sets a Offer value to be at the specified index for the
	// property "origin". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetOffer(idx int, v OfferInterface)
	// SetOrderedCollection sets a OrderedCollection value to be at the
	// specified index for the property "origin". Panics if the index is
	// out of bounds. Invalidates all iterators.
	SetOrderedCollection(idx int, v OrderedCollectionInterface)
	// SetOrderedCollectionPage sets a OrderedCollectionPage value to be at
	// the specified index for the property "origin". Panics if the index
	// is out of bounds. Invalidates all iterators.
	SetOrderedCollectionPage(idx int, v OrderedCollectionPageInterface)
	// SetOrganization sets a Organization value to be at the specified index
	// for the property "origin". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetOrganization(idx int, v OrganizationInterface)
	// SetPage sets a Page value to be at the specified index for the property
	// "origin". Panics if the index is out of bounds. Invalidates all
	// iterators.
	SetPage(idx int, v PageInterface)
	// SetPerson sets a Person value to be at the specified index for the
	// property "origin". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetPerson(idx int, v PersonInterface)
	// SetPlace sets a Place value to be at the specified index for the
	// property "origin". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetPlace(idx int, v PlaceInterface)
	// SetProfile sets a Profile value to be at the specified index for the
	// property "origin". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetProfile(idx int, v ProfileInterface)
	// SetQuestion sets a Question value to be at the specified index for the
	// property "origin". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetQuestion(idx int, v QuestionInterface)
	// SetRead sets a Read value to be at the specified index for the property
	// "origin". Panics if the index is out of bounds. Invalidates all
	// iterators.
	SetRead(idx int, v ReadInterface)
	// SetReject sets a Reject value to be at the specified index for the
	// property "origin". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetReject(idx int, v RejectInterface)
	// SetRelationship sets a Relationship value to be at the specified index
	// for the property "origin". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetRelationship(idx int, v RelationshipInterface)
	// SetRemove sets a Remove value to be at the specified index for the
	// property "origin". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetRemove(idx int, v RemoveInterface)
	// SetService sets a Service value to be at the specified index for the
	// property "origin". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetService(idx int, v ServiceInterface)
	// SetTentativeAccept sets a TentativeAccept value to be at the specified
	// index for the property "origin". Panics if the index is out of
	// bounds. Invalidates all iterators.
	SetTentativeAccept(idx int, v TentativeAcceptInterface)
	// SetTentativeReject sets a TentativeReject value to be at the specified
	// index for the property "origin". Panics if the index is out of
	// bounds. Invalidates all iterators.
	SetTentativeReject(idx int, v TentativeRejectInterface)
	// SetTombstone sets a Tombstone value to be at the specified index for
	// the property "origin". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetTombstone(idx int, v TombstoneInterface)
	// SetTravel sets a Travel value to be at the specified index for the
	// property "origin". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetTravel(idx int, v TravelInterface)
	// SetUndo sets a Undo value to be at the specified index for the property
	// "origin". Panics if the index is out of bounds. Invalidates all
	// iterators.
	SetUndo(idx int, v UndoInterface)
	// SetUpdate sets a Update value to be at the specified index for the
	// property "origin". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetUpdate(idx int, v UpdateInterface)
	// SetVideo sets a Video value to be at the specified index for the
	// property "origin". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetVideo(idx int, v VideoInterface)
	// SetView sets a View value to be at the specified index for the property
	// "origin". Panics if the index is out of bounds. Invalidates all
	// iterators.
	SetView(idx int, v ViewInterface)
	// Swap swaps the location of values at two indices for the "origin"
	// property.
	Swap(i, j int)
}

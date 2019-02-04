package vocab

import "net/url"

// On a Profile object, the describes property identifies the object described by
// the Profile.
//
// Example 141 (https://www.w3.org/TR/activitystreams-vocabulary/#ex185-jsonld):
//   {
//     "describes": {
//       "name": "Sally",
//       "type": "Person"
//     },
//     "summary": "Sally's profile",
//     "type": "Profile",
//     "url": "http://sally.example.org"
//   }
type DescribesPropertyInterface interface {
	// Clear ensures no value of this property is set. Calling HasAny or any
	// of the 'Is' methods afterwards will return false.
	Clear()
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
	// GetListen returns the value of this property. When IsListen returns
	// false, GetListen will return an arbitrary value.
	GetListen() ListenInterface
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
	// IsListen returns true if this property has a type of "Listen". When
	// true, use the GetListen and SetListen methods to access and set
	// this property.
	IsListen() bool
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
	LessThan(o DescribesPropertyInterface) bool
	// Name returns the name of this property: "describes".
	Name() string
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format. Applications should not
	// need this function as most typical use cases serialize types
	// instead of individual properties. It is exposed for alternatives to
	// go-fed implementations to use.
	Serialize() (interface{}, error)
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
	// SetListen sets the value of this property. Calling IsListen afterwards
	// returns true.
	SetListen(v ListenInterface)
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

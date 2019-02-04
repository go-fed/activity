package streams

import (
	typeaccept "github.com/go-fed/activity/streams/impl/activitystreams/type_accept"
	typeactivity "github.com/go-fed/activity/streams/impl/activitystreams/type_activity"
	typeadd "github.com/go-fed/activity/streams/impl/activitystreams/type_add"
	typeannounce "github.com/go-fed/activity/streams/impl/activitystreams/type_announce"
	typeapplication "github.com/go-fed/activity/streams/impl/activitystreams/type_application"
	typearrive "github.com/go-fed/activity/streams/impl/activitystreams/type_arrive"
	typearticle "github.com/go-fed/activity/streams/impl/activitystreams/type_article"
	typeaudio "github.com/go-fed/activity/streams/impl/activitystreams/type_audio"
	typeblock "github.com/go-fed/activity/streams/impl/activitystreams/type_block"
	typecollection "github.com/go-fed/activity/streams/impl/activitystreams/type_collection"
	typecollectionpage "github.com/go-fed/activity/streams/impl/activitystreams/type_collectionpage"
	typecreate "github.com/go-fed/activity/streams/impl/activitystreams/type_create"
	typedelete "github.com/go-fed/activity/streams/impl/activitystreams/type_delete"
	typedislike "github.com/go-fed/activity/streams/impl/activitystreams/type_dislike"
	typedocument "github.com/go-fed/activity/streams/impl/activitystreams/type_document"
	typeevent "github.com/go-fed/activity/streams/impl/activitystreams/type_event"
	typeflag "github.com/go-fed/activity/streams/impl/activitystreams/type_flag"
	typefollow "github.com/go-fed/activity/streams/impl/activitystreams/type_follow"
	typegroup "github.com/go-fed/activity/streams/impl/activitystreams/type_group"
	typeignore "github.com/go-fed/activity/streams/impl/activitystreams/type_ignore"
	typeimage "github.com/go-fed/activity/streams/impl/activitystreams/type_image"
	typeintransitiveactivity "github.com/go-fed/activity/streams/impl/activitystreams/type_intransitiveactivity"
	typeinvite "github.com/go-fed/activity/streams/impl/activitystreams/type_invite"
	typejoin "github.com/go-fed/activity/streams/impl/activitystreams/type_join"
	typeleave "github.com/go-fed/activity/streams/impl/activitystreams/type_leave"
	typelike "github.com/go-fed/activity/streams/impl/activitystreams/type_like"
	typelink "github.com/go-fed/activity/streams/impl/activitystreams/type_link"
	typelisten "github.com/go-fed/activity/streams/impl/activitystreams/type_listen"
	typemention "github.com/go-fed/activity/streams/impl/activitystreams/type_mention"
	typemove "github.com/go-fed/activity/streams/impl/activitystreams/type_move"
	typenote "github.com/go-fed/activity/streams/impl/activitystreams/type_note"
	typeobject "github.com/go-fed/activity/streams/impl/activitystreams/type_object"
	typeoffer "github.com/go-fed/activity/streams/impl/activitystreams/type_offer"
	typeorderedcollection "github.com/go-fed/activity/streams/impl/activitystreams/type_orderedcollection"
	typeorderedcollectionpage "github.com/go-fed/activity/streams/impl/activitystreams/type_orderedcollectionpage"
	typeorganization "github.com/go-fed/activity/streams/impl/activitystreams/type_organization"
	typepage "github.com/go-fed/activity/streams/impl/activitystreams/type_page"
	typeperson "github.com/go-fed/activity/streams/impl/activitystreams/type_person"
	typeplace "github.com/go-fed/activity/streams/impl/activitystreams/type_place"
	typeprofile "github.com/go-fed/activity/streams/impl/activitystreams/type_profile"
	typequestion "github.com/go-fed/activity/streams/impl/activitystreams/type_question"
	typeread "github.com/go-fed/activity/streams/impl/activitystreams/type_read"
	typereject "github.com/go-fed/activity/streams/impl/activitystreams/type_reject"
	typerelationship "github.com/go-fed/activity/streams/impl/activitystreams/type_relationship"
	typeremove "github.com/go-fed/activity/streams/impl/activitystreams/type_remove"
	typeservice "github.com/go-fed/activity/streams/impl/activitystreams/type_service"
	typetentativeaccept "github.com/go-fed/activity/streams/impl/activitystreams/type_tentativeaccept"
	typetentativereject "github.com/go-fed/activity/streams/impl/activitystreams/type_tentativereject"
	typetombstone "github.com/go-fed/activity/streams/impl/activitystreams/type_tombstone"
	typetravel "github.com/go-fed/activity/streams/impl/activitystreams/type_travel"
	typeundo "github.com/go-fed/activity/streams/impl/activitystreams/type_undo"
	typeupdate "github.com/go-fed/activity/streams/impl/activitystreams/type_update"
	typevideo "github.com/go-fed/activity/streams/impl/activitystreams/type_video"
	typeview "github.com/go-fed/activity/streams/impl/activitystreams/type_view"
	vocab "github.com/go-fed/activity/streams/vocab"
)

// NewActivityStreamsAccept creates a new AcceptInterface
func NewActivityStreamsAccept() vocab.AcceptInterface {
	return typeaccept.NewAccept()
}

// NewActivityStreamsActivity creates a new ActivityInterface
func NewActivityStreamsActivity() vocab.ActivityInterface {
	return typeactivity.NewActivity()
}

// NewActivityStreamsAdd creates a new AddInterface
func NewActivityStreamsAdd() vocab.AddInterface {
	return typeadd.NewAdd()
}

// NewActivityStreamsAnnounce creates a new AnnounceInterface
func NewActivityStreamsAnnounce() vocab.AnnounceInterface {
	return typeannounce.NewAnnounce()
}

// NewActivityStreamsApplication creates a new ApplicationInterface
func NewActivityStreamsApplication() vocab.ApplicationInterface {
	return typeapplication.NewApplication()
}

// NewActivityStreamsArrive creates a new ArriveInterface
func NewActivityStreamsArrive() vocab.ArriveInterface {
	return typearrive.NewArrive()
}

// NewActivityStreamsArticle creates a new ArticleInterface
func NewActivityStreamsArticle() vocab.ArticleInterface {
	return typearticle.NewArticle()
}

// NewActivityStreamsAudio creates a new AudioInterface
func NewActivityStreamsAudio() vocab.AudioInterface {
	return typeaudio.NewAudio()
}

// NewActivityStreamsBlock creates a new BlockInterface
func NewActivityStreamsBlock() vocab.BlockInterface {
	return typeblock.NewBlock()
}

// NewActivityStreamsCollection creates a new CollectionInterface
func NewActivityStreamsCollection() vocab.CollectionInterface {
	return typecollection.NewCollection()
}

// NewActivityStreamsCollectionPage creates a new CollectionPageInterface
func NewActivityStreamsCollectionPage() vocab.CollectionPageInterface {
	return typecollectionpage.NewCollectionPage()
}

// NewActivityStreamsCreate creates a new CreateInterface
func NewActivityStreamsCreate() vocab.CreateInterface {
	return typecreate.NewCreate()
}

// NewActivityStreamsDelete creates a new DeleteInterface
func NewActivityStreamsDelete() vocab.DeleteInterface {
	return typedelete.NewDelete()
}

// NewActivityStreamsDislike creates a new DislikeInterface
func NewActivityStreamsDislike() vocab.DislikeInterface {
	return typedislike.NewDislike()
}

// NewActivityStreamsDocument creates a new DocumentInterface
func NewActivityStreamsDocument() vocab.DocumentInterface {
	return typedocument.NewDocument()
}

// NewActivityStreamsEvent creates a new EventInterface
func NewActivityStreamsEvent() vocab.EventInterface {
	return typeevent.NewEvent()
}

// NewActivityStreamsFlag creates a new FlagInterface
func NewActivityStreamsFlag() vocab.FlagInterface {
	return typeflag.NewFlag()
}

// NewActivityStreamsFollow creates a new FollowInterface
func NewActivityStreamsFollow() vocab.FollowInterface {
	return typefollow.NewFollow()
}

// NewActivityStreamsGroup creates a new GroupInterface
func NewActivityStreamsGroup() vocab.GroupInterface {
	return typegroup.NewGroup()
}

// NewActivityStreamsIgnore creates a new IgnoreInterface
func NewActivityStreamsIgnore() vocab.IgnoreInterface {
	return typeignore.NewIgnore()
}

// NewActivityStreamsImage creates a new ImageInterface
func NewActivityStreamsImage() vocab.ImageInterface {
	return typeimage.NewImage()
}

// NewActivityStreamsIntransitiveActivity creates a new
// IntransitiveActivityInterface
func NewActivityStreamsIntransitiveActivity() vocab.IntransitiveActivityInterface {
	return typeintransitiveactivity.NewIntransitiveActivity()
}

// NewActivityStreamsInvite creates a new InviteInterface
func NewActivityStreamsInvite() vocab.InviteInterface {
	return typeinvite.NewInvite()
}

// NewActivityStreamsJoin creates a new JoinInterface
func NewActivityStreamsJoin() vocab.JoinInterface {
	return typejoin.NewJoin()
}

// NewActivityStreamsLeave creates a new LeaveInterface
func NewActivityStreamsLeave() vocab.LeaveInterface {
	return typeleave.NewLeave()
}

// NewActivityStreamsLike creates a new LikeInterface
func NewActivityStreamsLike() vocab.LikeInterface {
	return typelike.NewLike()
}

// NewActivityStreamsLink creates a new LinkInterface
func NewActivityStreamsLink() vocab.LinkInterface {
	return typelink.NewLink()
}

// NewActivityStreamsListen creates a new ListenInterface
func NewActivityStreamsListen() vocab.ListenInterface {
	return typelisten.NewListen()
}

// NewActivityStreamsMention creates a new MentionInterface
func NewActivityStreamsMention() vocab.MentionInterface {
	return typemention.NewMention()
}

// NewActivityStreamsMove creates a new MoveInterface
func NewActivityStreamsMove() vocab.MoveInterface {
	return typemove.NewMove()
}

// NewActivityStreamsNote creates a new NoteInterface
func NewActivityStreamsNote() vocab.NoteInterface {
	return typenote.NewNote()
}

// NewActivityStreamsObject creates a new ObjectInterface
func NewActivityStreamsObject() vocab.ObjectInterface {
	return typeobject.NewObject()
}

// NewActivityStreamsOffer creates a new OfferInterface
func NewActivityStreamsOffer() vocab.OfferInterface {
	return typeoffer.NewOffer()
}

// NewActivityStreamsOrderedCollection creates a new OrderedCollectionInterface
func NewActivityStreamsOrderedCollection() vocab.OrderedCollectionInterface {
	return typeorderedcollection.NewOrderedCollection()
}

// NewActivityStreamsOrderedCollectionPage creates a new
// OrderedCollectionPageInterface
func NewActivityStreamsOrderedCollectionPage() vocab.OrderedCollectionPageInterface {
	return typeorderedcollectionpage.NewOrderedCollectionPage()
}

// NewActivityStreamsOrganization creates a new OrganizationInterface
func NewActivityStreamsOrganization() vocab.OrganizationInterface {
	return typeorganization.NewOrganization()
}

// NewActivityStreamsPage creates a new PageInterface
func NewActivityStreamsPage() vocab.PageInterface {
	return typepage.NewPage()
}

// NewActivityStreamsPerson creates a new PersonInterface
func NewActivityStreamsPerson() vocab.PersonInterface {
	return typeperson.NewPerson()
}

// NewActivityStreamsPlace creates a new PlaceInterface
func NewActivityStreamsPlace() vocab.PlaceInterface {
	return typeplace.NewPlace()
}

// NewActivityStreamsProfile creates a new ProfileInterface
func NewActivityStreamsProfile() vocab.ProfileInterface {
	return typeprofile.NewProfile()
}

// NewActivityStreamsQuestion creates a new QuestionInterface
func NewActivityStreamsQuestion() vocab.QuestionInterface {
	return typequestion.NewQuestion()
}

// NewActivityStreamsRead creates a new ReadInterface
func NewActivityStreamsRead() vocab.ReadInterface {
	return typeread.NewRead()
}

// NewActivityStreamsReject creates a new RejectInterface
func NewActivityStreamsReject() vocab.RejectInterface {
	return typereject.NewReject()
}

// NewActivityStreamsRelationship creates a new RelationshipInterface
func NewActivityStreamsRelationship() vocab.RelationshipInterface {
	return typerelationship.NewRelationship()
}

// NewActivityStreamsRemove creates a new RemoveInterface
func NewActivityStreamsRemove() vocab.RemoveInterface {
	return typeremove.NewRemove()
}

// NewActivityStreamsService creates a new ServiceInterface
func NewActivityStreamsService() vocab.ServiceInterface {
	return typeservice.NewService()
}

// NewActivityStreamsTentativeAccept creates a new TentativeAcceptInterface
func NewActivityStreamsTentativeAccept() vocab.TentativeAcceptInterface {
	return typetentativeaccept.NewTentativeAccept()
}

// NewActivityStreamsTentativeReject creates a new TentativeRejectInterface
func NewActivityStreamsTentativeReject() vocab.TentativeRejectInterface {
	return typetentativereject.NewTentativeReject()
}

// NewActivityStreamsTombstone creates a new TombstoneInterface
func NewActivityStreamsTombstone() vocab.TombstoneInterface {
	return typetombstone.NewTombstone()
}

// NewActivityStreamsTravel creates a new TravelInterface
func NewActivityStreamsTravel() vocab.TravelInterface {
	return typetravel.NewTravel()
}

// NewActivityStreamsUndo creates a new UndoInterface
func NewActivityStreamsUndo() vocab.UndoInterface {
	return typeundo.NewUndo()
}

// NewActivityStreamsUpdate creates a new UpdateInterface
func NewActivityStreamsUpdate() vocab.UpdateInterface {
	return typeupdate.NewUpdate()
}

// NewActivityStreamsVideo creates a new VideoInterface
func NewActivityStreamsVideo() vocab.VideoInterface {
	return typevideo.NewVideo()
}

// NewActivityStreamsView creates a new ViewInterface
func NewActivityStreamsView() vocab.ViewInterface {
	return typeview.NewView()
}

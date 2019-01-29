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

// NewActivityStreamsLink creates a new LinkInterface
func NewActivityStreamsLink() vocab.LinkInterface {
	return typelink.NewLink()
}

// NewActivityStreamsObject creates a new ObjectInterface
func NewActivityStreamsObject() vocab.ObjectInterface {
	return typeobject.NewObject()
}

// NewActivityStreamsView creates a new ViewInterface
func NewActivityStreamsView() vocab.ViewInterface {
	return typeview.NewView()
}

// NewActivityStreamsOrderedCollectionPage creates a new
// OrderedCollectionPageInterface
func NewActivityStreamsOrderedCollectionPage() vocab.OrderedCollectionPageInterface {
	return typeorderedcollectionpage.NewOrderedCollectionPage()
}

// NewActivityStreamsPlace creates a new PlaceInterface
func NewActivityStreamsPlace() vocab.PlaceInterface {
	return typeplace.NewPlace()
}

// NewActivityStreamsGroup creates a new GroupInterface
func NewActivityStreamsGroup() vocab.GroupInterface {
	return typegroup.NewGroup()
}

// NewActivityStreamsFlag creates a new FlagInterface
func NewActivityStreamsFlag() vocab.FlagInterface {
	return typeflag.NewFlag()
}

// NewActivityStreamsAudio creates a new AudioInterface
func NewActivityStreamsAudio() vocab.AudioInterface {
	return typeaudio.NewAudio()
}

// NewActivityStreamsJoin creates a new JoinInterface
func NewActivityStreamsJoin() vocab.JoinInterface {
	return typejoin.NewJoin()
}

// NewActivityStreamsIgnore creates a new IgnoreInterface
func NewActivityStreamsIgnore() vocab.IgnoreInterface {
	return typeignore.NewIgnore()
}

// NewActivityStreamsVideo creates a new VideoInterface
func NewActivityStreamsVideo() vocab.VideoInterface {
	return typevideo.NewVideo()
}

// NewActivityStreamsCollection creates a new CollectionInterface
func NewActivityStreamsCollection() vocab.CollectionInterface {
	return typecollection.NewCollection()
}

// NewActivityStreamsCollectionPage creates a new CollectionPageInterface
func NewActivityStreamsCollectionPage() vocab.CollectionPageInterface {
	return typecollectionpage.NewCollectionPage()
}

// NewActivityStreamsApplication creates a new ApplicationInterface
func NewActivityStreamsApplication() vocab.ApplicationInterface {
	return typeapplication.NewApplication()
}

// NewActivityStreamsEvent creates a new EventInterface
func NewActivityStreamsEvent() vocab.EventInterface {
	return typeevent.NewEvent()
}

// NewActivityStreamsUndo creates a new UndoInterface
func NewActivityStreamsUndo() vocab.UndoInterface {
	return typeundo.NewUndo()
}

// NewActivityStreamsDislike creates a new DislikeInterface
func NewActivityStreamsDislike() vocab.DislikeInterface {
	return typedislike.NewDislike()
}

// NewActivityStreamsPerson creates a new PersonInterface
func NewActivityStreamsPerson() vocab.PersonInterface {
	return typeperson.NewPerson()
}

// NewActivityStreamsOrderedCollection creates a new OrderedCollectionInterface
func NewActivityStreamsOrderedCollection() vocab.OrderedCollectionInterface {
	return typeorderedcollection.NewOrderedCollection()
}

// NewActivityStreamsRead creates a new ReadInterface
func NewActivityStreamsRead() vocab.ReadInterface {
	return typeread.NewRead()
}

// NewActivityStreamsDelete creates a new DeleteInterface
func NewActivityStreamsDelete() vocab.DeleteInterface {
	return typedelete.NewDelete()
}

// NewActivityStreamsImage creates a new ImageInterface
func NewActivityStreamsImage() vocab.ImageInterface {
	return typeimage.NewImage()
}

// NewActivityStreamsFollow creates a new FollowInterface
func NewActivityStreamsFollow() vocab.FollowInterface {
	return typefollow.NewFollow()
}

// NewActivityStreamsTentativeReject creates a new TentativeRejectInterface
func NewActivityStreamsTentativeReject() vocab.TentativeRejectInterface {
	return typetentativereject.NewTentativeReject()
}

// NewActivityStreamsDocument creates a new DocumentInterface
func NewActivityStreamsDocument() vocab.DocumentInterface {
	return typedocument.NewDocument()
}

// NewActivityStreamsRelationship creates a new RelationshipInterface
func NewActivityStreamsRelationship() vocab.RelationshipInterface {
	return typerelationship.NewRelationship()
}

// NewActivityStreamsPage creates a new PageInterface
func NewActivityStreamsPage() vocab.PageInterface {
	return typepage.NewPage()
}

// NewActivityStreamsReject creates a new RejectInterface
func NewActivityStreamsReject() vocab.RejectInterface {
	return typereject.NewReject()
}

// NewActivityStreamsAccept creates a new AcceptInterface
func NewActivityStreamsAccept() vocab.AcceptInterface {
	return typeaccept.NewAccept()
}

// NewActivityStreamsListen creates a new ListenInterface
func NewActivityStreamsListen() vocab.ListenInterface {
	return typelisten.NewListen()
}

// NewActivityStreamsLike creates a new LikeInterface
func NewActivityStreamsLike() vocab.LikeInterface {
	return typelike.NewLike()
}

// NewActivityStreamsTombstone creates a new TombstoneInterface
func NewActivityStreamsTombstone() vocab.TombstoneInterface {
	return typetombstone.NewTombstone()
}

// NewActivityStreamsUpdate creates a new UpdateInterface
func NewActivityStreamsUpdate() vocab.UpdateInterface {
	return typeupdate.NewUpdate()
}

// NewActivityStreamsOrganization creates a new OrganizationInterface
func NewActivityStreamsOrganization() vocab.OrganizationInterface {
	return typeorganization.NewOrganization()
}

// NewActivityStreamsProfile creates a new ProfileInterface
func NewActivityStreamsProfile() vocab.ProfileInterface {
	return typeprofile.NewProfile()
}

// NewActivityStreamsMove creates a new MoveInterface
func NewActivityStreamsMove() vocab.MoveInterface {
	return typemove.NewMove()
}

// NewActivityStreamsQuestion creates a new QuestionInterface
func NewActivityStreamsQuestion() vocab.QuestionInterface {
	return typequestion.NewQuestion()
}

// NewActivityStreamsService creates a new ServiceInterface
func NewActivityStreamsService() vocab.ServiceInterface {
	return typeservice.NewService()
}

// NewActivityStreamsRemove creates a new RemoveInterface
func NewActivityStreamsRemove() vocab.RemoveInterface {
	return typeremove.NewRemove()
}

// NewActivityStreamsIntransitiveActivity creates a new
// IntransitiveActivityInterface
func NewActivityStreamsIntransitiveActivity() vocab.IntransitiveActivityInterface {
	return typeintransitiveactivity.NewIntransitiveActivity()
}

// NewActivityStreamsTravel creates a new TravelInterface
func NewActivityStreamsTravel() vocab.TravelInterface {
	return typetravel.NewTravel()
}

// NewActivityStreamsAnnounce creates a new AnnounceInterface
func NewActivityStreamsAnnounce() vocab.AnnounceInterface {
	return typeannounce.NewAnnounce()
}

// NewActivityStreamsNote creates a new NoteInterface
func NewActivityStreamsNote() vocab.NoteInterface {
	return typenote.NewNote()
}

// NewActivityStreamsInvite creates a new InviteInterface
func NewActivityStreamsInvite() vocab.InviteInterface {
	return typeinvite.NewInvite()
}

// NewActivityStreamsBlock creates a new BlockInterface
func NewActivityStreamsBlock() vocab.BlockInterface {
	return typeblock.NewBlock()
}

// NewActivityStreamsActivity creates a new ActivityInterface
func NewActivityStreamsActivity() vocab.ActivityInterface {
	return typeactivity.NewActivity()
}

// NewActivityStreamsArticle creates a new ArticleInterface
func NewActivityStreamsArticle() vocab.ArticleInterface {
	return typearticle.NewArticle()
}

// NewActivityStreamsArrive creates a new ArriveInterface
func NewActivityStreamsArrive() vocab.ArriveInterface {
	return typearrive.NewArrive()
}

// NewActivityStreamsMention creates a new MentionInterface
func NewActivityStreamsMention() vocab.MentionInterface {
	return typemention.NewMention()
}

// NewActivityStreamsTentativeAccept creates a new TentativeAcceptInterface
func NewActivityStreamsTentativeAccept() vocab.TentativeAcceptInterface {
	return typetentativeaccept.NewTentativeAccept()
}

// NewActivityStreamsOffer creates a new OfferInterface
func NewActivityStreamsOffer() vocab.OfferInterface {
	return typeoffer.NewOffer()
}

// NewActivityStreamsCreate creates a new CreateInterface
func NewActivityStreamsCreate() vocab.CreateInterface {
	return typecreate.NewCreate()
}

// NewActivityStreamsAdd creates a new AddInterface
func NewActivityStreamsAdd() vocab.AddInterface {
	return typeadd.NewAdd()
}

// NewActivityStreamsLeave creates a new LeaveInterface
func NewActivityStreamsLeave() vocab.LeaveInterface {
	return typeleave.NewLeave()
}

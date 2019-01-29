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

// ActivityStreamsLinkExtends returns true if Link extends from the other's type.
func ActivityStreamsLinkExtends(other vocab.Type) bool {
	return typelink.LinkExtends(other)
}

// ActivityStreamsObjectExtends returns true if Object extends from the other's
// type.
func ActivityStreamsObjectExtends(other vocab.Type) bool {
	return typeobject.ObjectExtends(other)
}

// ActivityStreamsViewExtends returns true if View extends from the other's type.
func ActivityStreamsViewExtends(other vocab.Type) bool {
	return typeview.ViewExtends(other)
}

// ActivityStreamsOrderedCollectionPageExtends returns true if
// OrderedCollectionPage extends from the other's type.
func ActivityStreamsOrderedCollectionPageExtends(other vocab.Type) bool {
	return typeorderedcollectionpage.OrderedCollectionPageExtends(other)
}

// ActivityStreamsPlaceExtends returns true if Place extends from the other's type.
func ActivityStreamsPlaceExtends(other vocab.Type) bool {
	return typeplace.PlaceExtends(other)
}

// ActivityStreamsGroupExtends returns true if Group extends from the other's type.
func ActivityStreamsGroupExtends(other vocab.Type) bool {
	return typegroup.GroupExtends(other)
}

// ActivityStreamsFlagExtends returns true if Flag extends from the other's type.
func ActivityStreamsFlagExtends(other vocab.Type) bool {
	return typeflag.FlagExtends(other)
}

// ActivityStreamsAudioExtends returns true if Audio extends from the other's type.
func ActivityStreamsAudioExtends(other vocab.Type) bool {
	return typeaudio.AudioExtends(other)
}

// ActivityStreamsJoinExtends returns true if Join extends from the other's type.
func ActivityStreamsJoinExtends(other vocab.Type) bool {
	return typejoin.JoinExtends(other)
}

// ActivityStreamsIgnoreExtends returns true if Ignore extends from the other's
// type.
func ActivityStreamsIgnoreExtends(other vocab.Type) bool {
	return typeignore.IgnoreExtends(other)
}

// ActivityStreamsVideoExtends returns true if Video extends from the other's type.
func ActivityStreamsVideoExtends(other vocab.Type) bool {
	return typevideo.VideoExtends(other)
}

// ActivityStreamsCollectionExtends returns true if Collection extends from the
// other's type.
func ActivityStreamsCollectionExtends(other vocab.Type) bool {
	return typecollection.CollectionExtends(other)
}

// ActivityStreamsCollectionPageExtends returns true if CollectionPage extends
// from the other's type.
func ActivityStreamsCollectionPageExtends(other vocab.Type) bool {
	return typecollectionpage.CollectionPageExtends(other)
}

// ActivityStreamsApplicationExtends returns true if Application extends from the
// other's type.
func ActivityStreamsApplicationExtends(other vocab.Type) bool {
	return typeapplication.ApplicationExtends(other)
}

// ActivityStreamsEventExtends returns true if Event extends from the other's type.
func ActivityStreamsEventExtends(other vocab.Type) bool {
	return typeevent.EventExtends(other)
}

// ActivityStreamsUndoExtends returns true if Undo extends from the other's type.
func ActivityStreamsUndoExtends(other vocab.Type) bool {
	return typeundo.UndoExtends(other)
}

// ActivityStreamsDislikeExtends returns true if Dislike extends from the other's
// type.
func ActivityStreamsDislikeExtends(other vocab.Type) bool {
	return typedislike.DislikeExtends(other)
}

// ActivityStreamsPersonExtends returns true if Person extends from the other's
// type.
func ActivityStreamsPersonExtends(other vocab.Type) bool {
	return typeperson.PersonExtends(other)
}

// ActivityStreamsOrderedCollectionExtends returns true if OrderedCollection
// extends from the other's type.
func ActivityStreamsOrderedCollectionExtends(other vocab.Type) bool {
	return typeorderedcollection.OrderedCollectionExtends(other)
}

// ActivityStreamsReadExtends returns true if Read extends from the other's type.
func ActivityStreamsReadExtends(other vocab.Type) bool {
	return typeread.ReadExtends(other)
}

// ActivityStreamsDeleteExtends returns true if Delete extends from the other's
// type.
func ActivityStreamsDeleteExtends(other vocab.Type) bool {
	return typedelete.DeleteExtends(other)
}

// ActivityStreamsImageExtends returns true if Image extends from the other's type.
func ActivityStreamsImageExtends(other vocab.Type) bool {
	return typeimage.ImageExtends(other)
}

// ActivityStreamsFollowExtends returns true if Follow extends from the other's
// type.
func ActivityStreamsFollowExtends(other vocab.Type) bool {
	return typefollow.FollowExtends(other)
}

// ActivityStreamsTentativeRejectExtends returns true if TentativeReject extends
// from the other's type.
func ActivityStreamsTentativeRejectExtends(other vocab.Type) bool {
	return typetentativereject.TentativeRejectExtends(other)
}

// ActivityStreamsDocumentExtends returns true if Document extends from the
// other's type.
func ActivityStreamsDocumentExtends(other vocab.Type) bool {
	return typedocument.DocumentExtends(other)
}

// ActivityStreamsRelationshipExtends returns true if Relationship extends from
// the other's type.
func ActivityStreamsRelationshipExtends(other vocab.Type) bool {
	return typerelationship.RelationshipExtends(other)
}

// ActivityStreamsPageExtends returns true if Page extends from the other's type.
func ActivityStreamsPageExtends(other vocab.Type) bool {
	return typepage.PageExtends(other)
}

// ActivityStreamsRejectExtends returns true if Reject extends from the other's
// type.
func ActivityStreamsRejectExtends(other vocab.Type) bool {
	return typereject.RejectExtends(other)
}

// ActivityStreamsAcceptExtends returns true if Accept extends from the other's
// type.
func ActivityStreamsAcceptExtends(other vocab.Type) bool {
	return typeaccept.AcceptExtends(other)
}

// ActivityStreamsListenExtends returns true if Listen extends from the other's
// type.
func ActivityStreamsListenExtends(other vocab.Type) bool {
	return typelisten.ListenExtends(other)
}

// ActivityStreamsLikeExtends returns true if Like extends from the other's type.
func ActivityStreamsLikeExtends(other vocab.Type) bool {
	return typelike.LikeExtends(other)
}

// ActivityStreamsTombstoneExtends returns true if Tombstone extends from the
// other's type.
func ActivityStreamsTombstoneExtends(other vocab.Type) bool {
	return typetombstone.TombstoneExtends(other)
}

// ActivityStreamsUpdateExtends returns true if Update extends from the other's
// type.
func ActivityStreamsUpdateExtends(other vocab.Type) bool {
	return typeupdate.UpdateExtends(other)
}

// ActivityStreamsOrganizationExtends returns true if Organization extends from
// the other's type.
func ActivityStreamsOrganizationExtends(other vocab.Type) bool {
	return typeorganization.OrganizationExtends(other)
}

// ActivityStreamsProfileExtends returns true if Profile extends from the other's
// type.
func ActivityStreamsProfileExtends(other vocab.Type) bool {
	return typeprofile.ProfileExtends(other)
}

// ActivityStreamsMoveExtends returns true if Move extends from the other's type.
func ActivityStreamsMoveExtends(other vocab.Type) bool {
	return typemove.MoveExtends(other)
}

// ActivityStreamsQuestionExtends returns true if Question extends from the
// other's type.
func ActivityStreamsQuestionExtends(other vocab.Type) bool {
	return typequestion.QuestionExtends(other)
}

// ActivityStreamsServiceExtends returns true if Service extends from the other's
// type.
func ActivityStreamsServiceExtends(other vocab.Type) bool {
	return typeservice.ServiceExtends(other)
}

// ActivityStreamsRemoveExtends returns true if Remove extends from the other's
// type.
func ActivityStreamsRemoveExtends(other vocab.Type) bool {
	return typeremove.RemoveExtends(other)
}

// ActivityStreamsIntransitiveActivityExtends returns true if IntransitiveActivity
// extends from the other's type.
func ActivityStreamsIntransitiveActivityExtends(other vocab.Type) bool {
	return typeintransitiveactivity.IntransitiveActivityExtends(other)
}

// ActivityStreamsTravelExtends returns true if Travel extends from the other's
// type.
func ActivityStreamsTravelExtends(other vocab.Type) bool {
	return typetravel.TravelExtends(other)
}

// ActivityStreamsAnnounceExtends returns true if Announce extends from the
// other's type.
func ActivityStreamsAnnounceExtends(other vocab.Type) bool {
	return typeannounce.AnnounceExtends(other)
}

// ActivityStreamsNoteExtends returns true if Note extends from the other's type.
func ActivityStreamsNoteExtends(other vocab.Type) bool {
	return typenote.NoteExtends(other)
}

// ActivityStreamsInviteExtends returns true if Invite extends from the other's
// type.
func ActivityStreamsInviteExtends(other vocab.Type) bool {
	return typeinvite.InviteExtends(other)
}

// ActivityStreamsBlockExtends returns true if Block extends from the other's type.
func ActivityStreamsBlockExtends(other vocab.Type) bool {
	return typeblock.BlockExtends(other)
}

// ActivityStreamsActivityExtends returns true if Activity extends from the
// other's type.
func ActivityStreamsActivityExtends(other vocab.Type) bool {
	return typeactivity.ActivityExtends(other)
}

// ActivityStreamsArticleExtends returns true if Article extends from the other's
// type.
func ActivityStreamsArticleExtends(other vocab.Type) bool {
	return typearticle.ArticleExtends(other)
}

// ActivityStreamsArriveExtends returns true if Arrive extends from the other's
// type.
func ActivityStreamsArriveExtends(other vocab.Type) bool {
	return typearrive.ArriveExtends(other)
}

// ActivityStreamsMentionExtends returns true if Mention extends from the other's
// type.
func ActivityStreamsMentionExtends(other vocab.Type) bool {
	return typemention.MentionExtends(other)
}

// ActivityStreamsTentativeAcceptExtends returns true if TentativeAccept extends
// from the other's type.
func ActivityStreamsTentativeAcceptExtends(other vocab.Type) bool {
	return typetentativeaccept.TentativeAcceptExtends(other)
}

// ActivityStreamsOfferExtends returns true if Offer extends from the other's type.
func ActivityStreamsOfferExtends(other vocab.Type) bool {
	return typeoffer.OfferExtends(other)
}

// ActivityStreamsCreateExtends returns true if Create extends from the other's
// type.
func ActivityStreamsCreateExtends(other vocab.Type) bool {
	return typecreate.CreateExtends(other)
}

// ActivityStreamsAddExtends returns true if Add extends from the other's type.
func ActivityStreamsAddExtends(other vocab.Type) bool {
	return typeadd.AddExtends(other)
}

// ActivityStreamsLeaveExtends returns true if Leave extends from the other's type.
func ActivityStreamsLeaveExtends(other vocab.Type) bool {
	return typeleave.LeaveExtends(other)
}

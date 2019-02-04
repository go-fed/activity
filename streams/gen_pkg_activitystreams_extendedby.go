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

// ActivityStreamsAcceptIsExtendedBy returns true if the other's type extends from
// Accept.
func ActivityStreamsAcceptIsExtendedBy(other vocab.Type) bool {
	return typeaccept.AcceptIsExtendedBy(other)
}

// ActivityStreamsActivityIsExtendedBy returns true if the other's type extends
// from Activity.
func ActivityStreamsActivityIsExtendedBy(other vocab.Type) bool {
	return typeactivity.ActivityIsExtendedBy(other)
}

// ActivityStreamsAddIsExtendedBy returns true if the other's type extends from
// Add.
func ActivityStreamsAddIsExtendedBy(other vocab.Type) bool {
	return typeadd.AddIsExtendedBy(other)
}

// ActivityStreamsAnnounceIsExtendedBy returns true if the other's type extends
// from Announce.
func ActivityStreamsAnnounceIsExtendedBy(other vocab.Type) bool {
	return typeannounce.AnnounceIsExtendedBy(other)
}

// ActivityStreamsApplicationIsExtendedBy returns true if the other's type extends
// from Application.
func ActivityStreamsApplicationIsExtendedBy(other vocab.Type) bool {
	return typeapplication.ApplicationIsExtendedBy(other)
}

// ActivityStreamsArriveIsExtendedBy returns true if the other's type extends from
// Arrive.
func ActivityStreamsArriveIsExtendedBy(other vocab.Type) bool {
	return typearrive.ArriveIsExtendedBy(other)
}

// ActivityStreamsArticleIsExtendedBy returns true if the other's type extends
// from Article.
func ActivityStreamsArticleIsExtendedBy(other vocab.Type) bool {
	return typearticle.ArticleIsExtendedBy(other)
}

// ActivityStreamsAudioIsExtendedBy returns true if the other's type extends from
// Audio.
func ActivityStreamsAudioIsExtendedBy(other vocab.Type) bool {
	return typeaudio.AudioIsExtendedBy(other)
}

// ActivityStreamsBlockIsExtendedBy returns true if the other's type extends from
// Block.
func ActivityStreamsBlockIsExtendedBy(other vocab.Type) bool {
	return typeblock.BlockIsExtendedBy(other)
}

// ActivityStreamsCollectionIsExtendedBy returns true if the other's type extends
// from Collection.
func ActivityStreamsCollectionIsExtendedBy(other vocab.Type) bool {
	return typecollection.CollectionIsExtendedBy(other)
}

// ActivityStreamsCollectionPageIsExtendedBy returns true if the other's type
// extends from CollectionPage.
func ActivityStreamsCollectionPageIsExtendedBy(other vocab.Type) bool {
	return typecollectionpage.CollectionPageIsExtendedBy(other)
}

// ActivityStreamsCreateIsExtendedBy returns true if the other's type extends from
// Create.
func ActivityStreamsCreateIsExtendedBy(other vocab.Type) bool {
	return typecreate.CreateIsExtendedBy(other)
}

// ActivityStreamsDeleteIsExtendedBy returns true if the other's type extends from
// Delete.
func ActivityStreamsDeleteIsExtendedBy(other vocab.Type) bool {
	return typedelete.DeleteIsExtendedBy(other)
}

// ActivityStreamsDislikeIsExtendedBy returns true if the other's type extends
// from Dislike.
func ActivityStreamsDislikeIsExtendedBy(other vocab.Type) bool {
	return typedislike.DislikeIsExtendedBy(other)
}

// ActivityStreamsDocumentIsExtendedBy returns true if the other's type extends
// from Document.
func ActivityStreamsDocumentIsExtendedBy(other vocab.Type) bool {
	return typedocument.DocumentIsExtendedBy(other)
}

// ActivityStreamsEventIsExtendedBy returns true if the other's type extends from
// Event.
func ActivityStreamsEventIsExtendedBy(other vocab.Type) bool {
	return typeevent.EventIsExtendedBy(other)
}

// ActivityStreamsFlagIsExtendedBy returns true if the other's type extends from
// Flag.
func ActivityStreamsFlagIsExtendedBy(other vocab.Type) bool {
	return typeflag.FlagIsExtendedBy(other)
}

// ActivityStreamsFollowIsExtendedBy returns true if the other's type extends from
// Follow.
func ActivityStreamsFollowIsExtendedBy(other vocab.Type) bool {
	return typefollow.FollowIsExtendedBy(other)
}

// ActivityStreamsGroupIsExtendedBy returns true if the other's type extends from
// Group.
func ActivityStreamsGroupIsExtendedBy(other vocab.Type) bool {
	return typegroup.GroupIsExtendedBy(other)
}

// ActivityStreamsIgnoreIsExtendedBy returns true if the other's type extends from
// Ignore.
func ActivityStreamsIgnoreIsExtendedBy(other vocab.Type) bool {
	return typeignore.IgnoreIsExtendedBy(other)
}

// ActivityStreamsImageIsExtendedBy returns true if the other's type extends from
// Image.
func ActivityStreamsImageIsExtendedBy(other vocab.Type) bool {
	return typeimage.ImageIsExtendedBy(other)
}

// ActivityStreamsIntransitiveActivityIsExtendedBy returns true if the other's
// type extends from IntransitiveActivity.
func ActivityStreamsIntransitiveActivityIsExtendedBy(other vocab.Type) bool {
	return typeintransitiveactivity.IntransitiveActivityIsExtendedBy(other)
}

// ActivityStreamsInviteIsExtendedBy returns true if the other's type extends from
// Invite.
func ActivityStreamsInviteIsExtendedBy(other vocab.Type) bool {
	return typeinvite.InviteIsExtendedBy(other)
}

// ActivityStreamsJoinIsExtendedBy returns true if the other's type extends from
// Join.
func ActivityStreamsJoinIsExtendedBy(other vocab.Type) bool {
	return typejoin.JoinIsExtendedBy(other)
}

// ActivityStreamsLeaveIsExtendedBy returns true if the other's type extends from
// Leave.
func ActivityStreamsLeaveIsExtendedBy(other vocab.Type) bool {
	return typeleave.LeaveIsExtendedBy(other)
}

// ActivityStreamsLikeIsExtendedBy returns true if the other's type extends from
// Like.
func ActivityStreamsLikeIsExtendedBy(other vocab.Type) bool {
	return typelike.LikeIsExtendedBy(other)
}

// ActivityStreamsLinkIsExtendedBy returns true if the other's type extends from
// Link.
func ActivityStreamsLinkIsExtendedBy(other vocab.Type) bool {
	return typelink.LinkIsExtendedBy(other)
}

// ActivityStreamsListenIsExtendedBy returns true if the other's type extends from
// Listen.
func ActivityStreamsListenIsExtendedBy(other vocab.Type) bool {
	return typelisten.ListenIsExtendedBy(other)
}

// ActivityStreamsMentionIsExtendedBy returns true if the other's type extends
// from Mention.
func ActivityStreamsMentionIsExtendedBy(other vocab.Type) bool {
	return typemention.MentionIsExtendedBy(other)
}

// ActivityStreamsMoveIsExtendedBy returns true if the other's type extends from
// Move.
func ActivityStreamsMoveIsExtendedBy(other vocab.Type) bool {
	return typemove.MoveIsExtendedBy(other)
}

// ActivityStreamsNoteIsExtendedBy returns true if the other's type extends from
// Note.
func ActivityStreamsNoteIsExtendedBy(other vocab.Type) bool {
	return typenote.NoteIsExtendedBy(other)
}

// ActivityStreamsObjectIsExtendedBy returns true if the other's type extends from
// Object.
func ActivityStreamsObjectIsExtendedBy(other vocab.Type) bool {
	return typeobject.ObjectIsExtendedBy(other)
}

// ActivityStreamsOfferIsExtendedBy returns true if the other's type extends from
// Offer.
func ActivityStreamsOfferIsExtendedBy(other vocab.Type) bool {
	return typeoffer.OfferIsExtendedBy(other)
}

// ActivityStreamsOrderedCollectionIsExtendedBy returns true if the other's type
// extends from OrderedCollection.
func ActivityStreamsOrderedCollectionIsExtendedBy(other vocab.Type) bool {
	return typeorderedcollection.OrderedCollectionIsExtendedBy(other)
}

// ActivityStreamsOrderedCollectionPageIsExtendedBy returns true if the other's
// type extends from OrderedCollectionPage.
func ActivityStreamsOrderedCollectionPageIsExtendedBy(other vocab.Type) bool {
	return typeorderedcollectionpage.OrderedCollectionPageIsExtendedBy(other)
}

// ActivityStreamsOrganizationIsExtendedBy returns true if the other's type
// extends from Organization.
func ActivityStreamsOrganizationIsExtendedBy(other vocab.Type) bool {
	return typeorganization.OrganizationIsExtendedBy(other)
}

// ActivityStreamsPageIsExtendedBy returns true if the other's type extends from
// Page.
func ActivityStreamsPageIsExtendedBy(other vocab.Type) bool {
	return typepage.PageIsExtendedBy(other)
}

// ActivityStreamsPersonIsExtendedBy returns true if the other's type extends from
// Person.
func ActivityStreamsPersonIsExtendedBy(other vocab.Type) bool {
	return typeperson.PersonIsExtendedBy(other)
}

// ActivityStreamsPlaceIsExtendedBy returns true if the other's type extends from
// Place.
func ActivityStreamsPlaceIsExtendedBy(other vocab.Type) bool {
	return typeplace.PlaceIsExtendedBy(other)
}

// ActivityStreamsProfileIsExtendedBy returns true if the other's type extends
// from Profile.
func ActivityStreamsProfileIsExtendedBy(other vocab.Type) bool {
	return typeprofile.ProfileIsExtendedBy(other)
}

// ActivityStreamsQuestionIsExtendedBy returns true if the other's type extends
// from Question.
func ActivityStreamsQuestionIsExtendedBy(other vocab.Type) bool {
	return typequestion.QuestionIsExtendedBy(other)
}

// ActivityStreamsReadIsExtendedBy returns true if the other's type extends from
// Read.
func ActivityStreamsReadIsExtendedBy(other vocab.Type) bool {
	return typeread.ReadIsExtendedBy(other)
}

// ActivityStreamsRejectIsExtendedBy returns true if the other's type extends from
// Reject.
func ActivityStreamsRejectIsExtendedBy(other vocab.Type) bool {
	return typereject.RejectIsExtendedBy(other)
}

// ActivityStreamsRelationshipIsExtendedBy returns true if the other's type
// extends from Relationship.
func ActivityStreamsRelationshipIsExtendedBy(other vocab.Type) bool {
	return typerelationship.RelationshipIsExtendedBy(other)
}

// ActivityStreamsRemoveIsExtendedBy returns true if the other's type extends from
// Remove.
func ActivityStreamsRemoveIsExtendedBy(other vocab.Type) bool {
	return typeremove.RemoveIsExtendedBy(other)
}

// ActivityStreamsServiceIsExtendedBy returns true if the other's type extends
// from Service.
func ActivityStreamsServiceIsExtendedBy(other vocab.Type) bool {
	return typeservice.ServiceIsExtendedBy(other)
}

// ActivityStreamsTentativeAcceptIsExtendedBy returns true if the other's type
// extends from TentativeAccept.
func ActivityStreamsTentativeAcceptIsExtendedBy(other vocab.Type) bool {
	return typetentativeaccept.TentativeAcceptIsExtendedBy(other)
}

// ActivityStreamsTentativeRejectIsExtendedBy returns true if the other's type
// extends from TentativeReject.
func ActivityStreamsTentativeRejectIsExtendedBy(other vocab.Type) bool {
	return typetentativereject.TentativeRejectIsExtendedBy(other)
}

// ActivityStreamsTombstoneIsExtendedBy returns true if the other's type extends
// from Tombstone.
func ActivityStreamsTombstoneIsExtendedBy(other vocab.Type) bool {
	return typetombstone.TombstoneIsExtendedBy(other)
}

// ActivityStreamsTravelIsExtendedBy returns true if the other's type extends from
// Travel.
func ActivityStreamsTravelIsExtendedBy(other vocab.Type) bool {
	return typetravel.TravelIsExtendedBy(other)
}

// ActivityStreamsUndoIsExtendedBy returns true if the other's type extends from
// Undo.
func ActivityStreamsUndoIsExtendedBy(other vocab.Type) bool {
	return typeundo.UndoIsExtendedBy(other)
}

// ActivityStreamsUpdateIsExtendedBy returns true if the other's type extends from
// Update.
func ActivityStreamsUpdateIsExtendedBy(other vocab.Type) bool {
	return typeupdate.UpdateIsExtendedBy(other)
}

// ActivityStreamsVideoIsExtendedBy returns true if the other's type extends from
// Video.
func ActivityStreamsVideoIsExtendedBy(other vocab.Type) bool {
	return typevideo.VideoIsExtendedBy(other)
}

// ActivityStreamsViewIsExtendedBy returns true if the other's type extends from
// View.
func ActivityStreamsViewIsExtendedBy(other vocab.Type) bool {
	return typeview.ViewIsExtendedBy(other)
}

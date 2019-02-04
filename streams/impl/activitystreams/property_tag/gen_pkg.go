package propertytag

import vocab "github.com/go-fed/activity/streams/vocab"

var mgr privateManager

// privateManager abstracts the code-generated manager that provides access to
// concrete implementations.
type privateManager interface {
	// DeserializeAcceptActivityStreams returns the deserialization method for
	// the "AcceptInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeAcceptActivityStreams() func(map[string]interface{}, map[string]string) (vocab.AcceptInterface, error)
	// DeserializeActivityActivityStreams returns the deserialization method
	// for the "ActivityInterface" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeActivityActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityInterface, error)
	// DeserializeAddActivityStreams returns the deserialization method for
	// the "AddInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeAddActivityStreams() func(map[string]interface{}, map[string]string) (vocab.AddInterface, error)
	// DeserializeAnnounceActivityStreams returns the deserialization method
	// for the "AnnounceInterface" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeAnnounceActivityStreams() func(map[string]interface{}, map[string]string) (vocab.AnnounceInterface, error)
	// DeserializeApplicationActivityStreams returns the deserialization
	// method for the "ApplicationInterface" non-functional property in
	// the vocabulary "ActivityStreams"
	DeserializeApplicationActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ApplicationInterface, error)
	// DeserializeArriveActivityStreams returns the deserialization method for
	// the "ArriveInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeArriveActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ArriveInterface, error)
	// DeserializeArticleActivityStreams returns the deserialization method
	// for the "ArticleInterface" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeArticleActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ArticleInterface, error)
	// DeserializeAudioActivityStreams returns the deserialization method for
	// the "AudioInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeAudioActivityStreams() func(map[string]interface{}, map[string]string) (vocab.AudioInterface, error)
	// DeserializeBlockActivityStreams returns the deserialization method for
	// the "BlockInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeBlockActivityStreams() func(map[string]interface{}, map[string]string) (vocab.BlockInterface, error)
	// DeserializeCollectionActivityStreams returns the deserialization method
	// for the "CollectionInterface" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeCollectionActivityStreams() func(map[string]interface{}, map[string]string) (vocab.CollectionInterface, error)
	// DeserializeCollectionPageActivityStreams returns the deserialization
	// method for the "CollectionPageInterface" non-functional property in
	// the vocabulary "ActivityStreams"
	DeserializeCollectionPageActivityStreams() func(map[string]interface{}, map[string]string) (vocab.CollectionPageInterface, error)
	// DeserializeCreateActivityStreams returns the deserialization method for
	// the "CreateInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeCreateActivityStreams() func(map[string]interface{}, map[string]string) (vocab.CreateInterface, error)
	// DeserializeDeleteActivityStreams returns the deserialization method for
	// the "DeleteInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeDeleteActivityStreams() func(map[string]interface{}, map[string]string) (vocab.DeleteInterface, error)
	// DeserializeDislikeActivityStreams returns the deserialization method
	// for the "DislikeInterface" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeDislikeActivityStreams() func(map[string]interface{}, map[string]string) (vocab.DislikeInterface, error)
	// DeserializeDocumentActivityStreams returns the deserialization method
	// for the "DocumentInterface" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeDocumentActivityStreams() func(map[string]interface{}, map[string]string) (vocab.DocumentInterface, error)
	// DeserializeEventActivityStreams returns the deserialization method for
	// the "EventInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeEventActivityStreams() func(map[string]interface{}, map[string]string) (vocab.EventInterface, error)
	// DeserializeFlagActivityStreams returns the deserialization method for
	// the "FlagInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeFlagActivityStreams() func(map[string]interface{}, map[string]string) (vocab.FlagInterface, error)
	// DeserializeFollowActivityStreams returns the deserialization method for
	// the "FollowInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeFollowActivityStreams() func(map[string]interface{}, map[string]string) (vocab.FollowInterface, error)
	// DeserializeGroupActivityStreams returns the deserialization method for
	// the "GroupInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeGroupActivityStreams() func(map[string]interface{}, map[string]string) (vocab.GroupInterface, error)
	// DeserializeIgnoreActivityStreams returns the deserialization method for
	// the "IgnoreInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeIgnoreActivityStreams() func(map[string]interface{}, map[string]string) (vocab.IgnoreInterface, error)
	// DeserializeImageActivityStreams returns the deserialization method for
	// the "ImageInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeImageActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ImageInterface, error)
	// DeserializeIntransitiveActivityActivityStreams returns the
	// deserialization method for the "IntransitiveActivityInterface"
	// non-functional property in the vocabulary "ActivityStreams"
	DeserializeIntransitiveActivityActivityStreams() func(map[string]interface{}, map[string]string) (vocab.IntransitiveActivityInterface, error)
	// DeserializeInviteActivityStreams returns the deserialization method for
	// the "InviteInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeInviteActivityStreams() func(map[string]interface{}, map[string]string) (vocab.InviteInterface, error)
	// DeserializeJoinActivityStreams returns the deserialization method for
	// the "JoinInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeJoinActivityStreams() func(map[string]interface{}, map[string]string) (vocab.JoinInterface, error)
	// DeserializeLeaveActivityStreams returns the deserialization method for
	// the "LeaveInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeLeaveActivityStreams() func(map[string]interface{}, map[string]string) (vocab.LeaveInterface, error)
	// DeserializeLikeActivityStreams returns the deserialization method for
	// the "LikeInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeLikeActivityStreams() func(map[string]interface{}, map[string]string) (vocab.LikeInterface, error)
	// DeserializeLinkActivityStreams returns the deserialization method for
	// the "LinkInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeLinkActivityStreams() func(map[string]interface{}, map[string]string) (vocab.LinkInterface, error)
	// DeserializeListenActivityStreams returns the deserialization method for
	// the "ListenInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeListenActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ListenInterface, error)
	// DeserializeMentionActivityStreams returns the deserialization method
	// for the "MentionInterface" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeMentionActivityStreams() func(map[string]interface{}, map[string]string) (vocab.MentionInterface, error)
	// DeserializeMoveActivityStreams returns the deserialization method for
	// the "MoveInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeMoveActivityStreams() func(map[string]interface{}, map[string]string) (vocab.MoveInterface, error)
	// DeserializeNoteActivityStreams returns the deserialization method for
	// the "NoteInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeNoteActivityStreams() func(map[string]interface{}, map[string]string) (vocab.NoteInterface, error)
	// DeserializeObjectActivityStreams returns the deserialization method for
	// the "ObjectInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeObjectActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ObjectInterface, error)
	// DeserializeOfferActivityStreams returns the deserialization method for
	// the "OfferInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeOfferActivityStreams() func(map[string]interface{}, map[string]string) (vocab.OfferInterface, error)
	// DeserializeOrderedCollectionActivityStreams returns the deserialization
	// method for the "OrderedCollectionInterface" non-functional property
	// in the vocabulary "ActivityStreams"
	DeserializeOrderedCollectionActivityStreams() func(map[string]interface{}, map[string]string) (vocab.OrderedCollectionInterface, error)
	// DeserializeOrderedCollectionPageActivityStreams returns the
	// deserialization method for the "OrderedCollectionPageInterface"
	// non-functional property in the vocabulary "ActivityStreams"
	DeserializeOrderedCollectionPageActivityStreams() func(map[string]interface{}, map[string]string) (vocab.OrderedCollectionPageInterface, error)
	// DeserializeOrganizationActivityStreams returns the deserialization
	// method for the "OrganizationInterface" non-functional property in
	// the vocabulary "ActivityStreams"
	DeserializeOrganizationActivityStreams() func(map[string]interface{}, map[string]string) (vocab.OrganizationInterface, error)
	// DeserializePageActivityStreams returns the deserialization method for
	// the "PageInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializePageActivityStreams() func(map[string]interface{}, map[string]string) (vocab.PageInterface, error)
	// DeserializePersonActivityStreams returns the deserialization method for
	// the "PersonInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializePersonActivityStreams() func(map[string]interface{}, map[string]string) (vocab.PersonInterface, error)
	// DeserializePlaceActivityStreams returns the deserialization method for
	// the "PlaceInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializePlaceActivityStreams() func(map[string]interface{}, map[string]string) (vocab.PlaceInterface, error)
	// DeserializeProfileActivityStreams returns the deserialization method
	// for the "ProfileInterface" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeProfileActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ProfileInterface, error)
	// DeserializeQuestionActivityStreams returns the deserialization method
	// for the "QuestionInterface" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeQuestionActivityStreams() func(map[string]interface{}, map[string]string) (vocab.QuestionInterface, error)
	// DeserializeReadActivityStreams returns the deserialization method for
	// the "ReadInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeReadActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ReadInterface, error)
	// DeserializeRejectActivityStreams returns the deserialization method for
	// the "RejectInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeRejectActivityStreams() func(map[string]interface{}, map[string]string) (vocab.RejectInterface, error)
	// DeserializeRelationshipActivityStreams returns the deserialization
	// method for the "RelationshipInterface" non-functional property in
	// the vocabulary "ActivityStreams"
	DeserializeRelationshipActivityStreams() func(map[string]interface{}, map[string]string) (vocab.RelationshipInterface, error)
	// DeserializeRemoveActivityStreams returns the deserialization method for
	// the "RemoveInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeRemoveActivityStreams() func(map[string]interface{}, map[string]string) (vocab.RemoveInterface, error)
	// DeserializeServiceActivityStreams returns the deserialization method
	// for the "ServiceInterface" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeServiceActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ServiceInterface, error)
	// DeserializeTentativeAcceptActivityStreams returns the deserialization
	// method for the "TentativeAcceptInterface" non-functional property
	// in the vocabulary "ActivityStreams"
	DeserializeTentativeAcceptActivityStreams() func(map[string]interface{}, map[string]string) (vocab.TentativeAcceptInterface, error)
	// DeserializeTentativeRejectActivityStreams returns the deserialization
	// method for the "TentativeRejectInterface" non-functional property
	// in the vocabulary "ActivityStreams"
	DeserializeTentativeRejectActivityStreams() func(map[string]interface{}, map[string]string) (vocab.TentativeRejectInterface, error)
	// DeserializeTombstoneActivityStreams returns the deserialization method
	// for the "TombstoneInterface" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeTombstoneActivityStreams() func(map[string]interface{}, map[string]string) (vocab.TombstoneInterface, error)
	// DeserializeTravelActivityStreams returns the deserialization method for
	// the "TravelInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeTravelActivityStreams() func(map[string]interface{}, map[string]string) (vocab.TravelInterface, error)
	// DeserializeUndoActivityStreams returns the deserialization method for
	// the "UndoInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeUndoActivityStreams() func(map[string]interface{}, map[string]string) (vocab.UndoInterface, error)
	// DeserializeUpdateActivityStreams returns the deserialization method for
	// the "UpdateInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeUpdateActivityStreams() func(map[string]interface{}, map[string]string) (vocab.UpdateInterface, error)
	// DeserializeVideoActivityStreams returns the deserialization method for
	// the "VideoInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeVideoActivityStreams() func(map[string]interface{}, map[string]string) (vocab.VideoInterface, error)
	// DeserializeViewActivityStreams returns the deserialization method for
	// the "ViewInterface" non-functional property in the vocabulary
	// "ActivityStreams"
	DeserializeViewActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ViewInterface, error)
}

// SetManager sets the manager package-global variable. For internal use only, do
// not use as part of Application behavior. Must be called at golang init time.
func SetManager(m privateManager) {
	mgr = m
}

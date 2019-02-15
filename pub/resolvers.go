package pub

import (
	"context"
	"github.com/go-fed/activity/streams"
	"github.com/go-fed/activity/streams/vocab"
)

// IsAnActivityType returns true if the ActivityStreams value is an Activity or
// extends from the Activity type.
func IsAnActivityType(value vocab.Type) bool {
	// TODO: Generate these strings as exported constants.
	return value.GetTypeName() == "Activity" || streams.ActivityStreamsActivityIsExtendedBy(value)
}

// toAsType converts a generic map[string]interface{} into a known Type.
//
// Returns errors under the same conditions as streams.JSONResolver does.
func toType(c context.Context, m map[string]interface{}) (a vocab.Type, e error) {
	var r *streams.JSONResolver
	// Every time new types are added, need to update this list. It looks
	// painful, but in practice VIM macros make it easier to manage.
	//
	// TODO: Somehow generate this more easily.
	r, e = streams.NewJSONResolver(
		func(ctx context.Context, i vocab.ActivityStreamsAccept) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsActivity) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsAdd) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsAnnounce) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsApplication) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsArrive) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsArticle) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsAudio) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsBlock) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsCollection) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsCollectionPage) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsCreate) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsDelete) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsDislike) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsDocument) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsEvent) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsFlag) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsFollow) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsGroup) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsIgnore) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsImage) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsIntransitiveActivity) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsInvite) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsJoin) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsLeave) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsLike) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsLink) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsListen) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsMention) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsMove) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsNote) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsObject) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsOffer) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsOrderedCollection) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsOrderedCollectionPage) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsOrganization) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsPage) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsPerson) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsPlace) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsProfile) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsQuestion) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsRead) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsReject) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsRelationship) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsRemove) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsService) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsTentativeAccept) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsTentativeReject) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsTombstone) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsTravel) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsUndo) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsUpdate) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsVideo) error {
			a = i
			return nil
		},
		func(ctx context.Context, i vocab.ActivityStreamsView) error {
			a = i
			return nil
		},
	)
	if e != nil {
		return
	}
	e = r.Resolve(c, m)
	return
}

// addToCreate adds the object to the Create activity.
func addToCreate(ctx context.Context, c vocab.ActivityStreamsCreate, o vocab.Type) error {
	obj := c.GetActivityStreamsObject()
	if obj == nil {
		obj = streams.NewActivityStreamsObjectProperty()
	}
	// Every time new types are added, need to update this list. It looks
	// painful, but in practice VIM macros make it easier to manage.
	//
	// TODO: Somehow generate this more easily.
	r, e := streams.NewTypeResolver(
		func(ctx context.Context, v vocab.ActivityStreamsAccept) error {
			obj.AppendActivityStreamsAccept(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsActivity) error {
			obj.AppendActivityStreamsActivity(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsAdd) error {
			obj.AppendActivityStreamsAdd(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsAnnounce) error {
			obj.AppendActivityStreamsAnnounce(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsApplication) error {
			obj.AppendActivityStreamsApplication(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsArrive) error {
			obj.AppendActivityStreamsArrive(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsArticle) error {
			obj.AppendActivityStreamsArticle(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsAudio) error {
			obj.AppendActivityStreamsAudio(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsBlock) error {
			obj.AppendActivityStreamsBlock(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsCollection) error {
			obj.AppendActivityStreamsCollection(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsCollectionPage) error {
			obj.AppendActivityStreamsCollectionPage(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsCreate) error {
			obj.AppendActivityStreamsCreate(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsDelete) error {
			obj.AppendActivityStreamsDelete(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsDislike) error {
			obj.AppendActivityStreamsDislike(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsDocument) error {
			obj.AppendActivityStreamsDocument(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsEvent) error {
			obj.AppendActivityStreamsEvent(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsFlag) error {
			obj.AppendActivityStreamsFlag(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsFollow) error {
			obj.AppendActivityStreamsFollow(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsGroup) error {
			obj.AppendActivityStreamsGroup(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsIgnore) error {
			obj.AppendActivityStreamsIgnore(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsImage) error {
			obj.AppendActivityStreamsImage(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsIntransitiveActivity) error {
			obj.AppendActivityStreamsIntransitiveActivity(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsInvite) error {
			obj.AppendActivityStreamsInvite(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsJoin) error {
			obj.AppendActivityStreamsJoin(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsLeave) error {
			obj.AppendActivityStreamsLeave(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsLike) error {
			obj.AppendActivityStreamsLike(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsLink) error {
			obj.AppendActivityStreamsLink(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsListen) error {
			obj.AppendActivityStreamsListen(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsMention) error {
			obj.AppendActivityStreamsMention(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsMove) error {
			obj.AppendActivityStreamsMove(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsNote) error {
			obj.AppendActivityStreamsNote(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsObject) error {
			obj.AppendActivityStreamsObject(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsOffer) error {
			obj.AppendActivityStreamsOffer(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsOrderedCollection) error {
			obj.AppendActivityStreamsOrderedCollection(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsOrderedCollectionPage) error {
			obj.AppendActivityStreamsOrderedCollectionPage(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsOrganization) error {
			obj.AppendActivityStreamsOrganization(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsPage) error {
			obj.AppendActivityStreamsPage(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsPerson) error {
			obj.AppendActivityStreamsPerson(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsPlace) error {
			obj.AppendActivityStreamsPlace(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsProfile) error {
			obj.AppendActivityStreamsProfile(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsQuestion) error {
			obj.AppendActivityStreamsQuestion(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsRead) error {
			obj.AppendActivityStreamsRead(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsReject) error {
			obj.AppendActivityStreamsReject(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsRelationship) error {
			obj.AppendActivityStreamsRelationship(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsRemove) error {
			obj.AppendActivityStreamsRemove(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsService) error {
			obj.AppendActivityStreamsService(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsTentativeAccept) error {
			obj.AppendActivityStreamsTentativeAccept(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsTentativeReject) error {
			obj.AppendActivityStreamsTentativeReject(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsTombstone) error {
			obj.AppendActivityStreamsTombstone(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsTravel) error {
			obj.AppendActivityStreamsTravel(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsUndo) error {
			obj.AppendActivityStreamsUndo(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsUpdate) error {
			obj.AppendActivityStreamsUpdate(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsVideo) error {
			obj.AppendActivityStreamsVideo(v)
			return nil
		},
		func(ctx context.Context, v vocab.ActivityStreamsView) error {
			obj.AppendActivityStreamsView(v)
			return nil
		},
	)
	if e != nil {
		return e
	}
	return r.Resolve(ctx, o)
}

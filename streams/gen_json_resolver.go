package streams

import (
	"context"
	"errors"
	"fmt"
	vocab "github.com/go-fed/activity/streams/vocab"
	"strings"
)

// JSONResolver resolves a JSON-deserialized map into its concrete ActivityStreams
// type
type JSONResolver struct {
	callbacks []interface{}
}

// NewJSONResolver creates a new Resolver that takes a JSON-deserialized generic
// map and determines the correct concrete Go type. The callback function is
// guaranteed to receive a value whose underlying ActivityStreams type matches
// the concrete interface name in its signature. The callback functions must
// be of the form:
//
//   func(context.Context, <TypeInterface>) error
//
// where TypeInterface is the code-generated interface for an ActivityStream
// type. An error is returned if a callback function does not match this
// signature.
func NewJSONResolver(callbacks ...interface{}) (*JSONResolver, error) {
	for _, cb := range callbacks {
		// Each callback function must satisfy one known function signature, or else we will generate a runtime error instead of silently fail.
		switch cb.(type) {
		case func(context.Context, vocab.AcceptInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.AddInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.AnnounceInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ApplicationInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ArriveInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ArticleInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.AudioInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.BlockInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.CollectionInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.CollectionPageInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.CreateInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.DeleteInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.DislikeInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.DocumentInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.EventInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.FlagInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.FollowInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.GroupInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.IgnoreInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ImageInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.IntransitiveActivityInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.InviteInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.JoinInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.LeaveInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.LikeInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.LinkInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ListenInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.MentionInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.MoveInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.NoteInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ObjectInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.OfferInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.OrderedCollectionInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.OrderedCollectionPageInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.OrganizationInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.PageInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.PersonInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.PlaceInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ProfileInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.QuestionInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ReadInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.RejectInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.RelationshipInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.RemoveInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ServiceInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.TentativeAcceptInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.TentativeRejectInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.TombstoneInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.TravelInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.UndoInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.UpdateInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.VideoInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ViewInterface) error:
			// Do nothing, this callback has a correct signature.
		default:
			return nil, errors.New("a callback function is of the wrong signature and would never be called")
		}
	}
	return &JSONResolver{callbacks: callbacks}, nil
}

// toAliasMap converts a JSONLD context into a map of vocabulary name to alias.
func toAliasMap(i interface{}) (m map[string]string) {
	m = make(map[string]string)
	toHttpHttpsFn := func(s string) (ok bool, http, https string) {
		if strings.HasPrefix(s, "http://") {
			ok = true
			http = s
			https = "https" + strings.TrimPrefix(s, "http")
		} else if strings.HasPrefix(s, "https://") {
			ok = true
			https = s
			http = "http" + strings.TrimPrefix(s, "https")
		}
		return
	}
	switch v := i.(type) {
	case string:
		// Single entry, no alias.

		if ok, http, https := toHttpHttpsFn(v); ok {
			m[http] = ""
			m[https] = ""
		} else {
			m[v] = ""
		}
	case []interface{}:
		// Recursively apply.

		for _, elem := range v {
			r := toAliasMap(elem)
			for k, val := range r {
				m[k] = val
			}
		}
	case map[string]interface{}:
		// Map any aliases.

		for k, val := range v {
			// Only handle string aliases.
			switch conc := val.(type) {
			case string:
				m[k] = conc
			}
		}
	}
	return
}

// Resolve determines the ActivityStreams type of the payload, then applies the
// first callback function whose signature accepts the ActivityStreams value's
// type. This strictly assures that the callback function will only be passed
// ActivityStream objects whose type matches its interface. Returns an error
// if the ActivityStreams type does not match callbackers or is not a type
// handled by the generated code. If multiple types are present, it will check
// each one in order and apply only the first one. It returns an unhandled
// error for a multi-typed object if none of the types were able to be handled.
func (this JSONResolver) Resolve(ctx context.Context, m map[string]interface{}) error {
	typeValue, ok := m["type"]
	if !ok {
		return fmt.Errorf("cannot determine ActivityStreams type: 'type' property is missing")
	}
	rawContext, ok := m["@context"]
	if !ok {
		return fmt.Errorf("cannot determine ActivityStreams type: '@context' is missing")
	}
	aliasMap := toAliasMap(rawContext)
	// Begin: Private lambda to handle a single string "type" value. Makes code generation easier.
	handleFn := func(typeString string) error {
		switch typeString {
		case "Accept":
			v, err := mgr.DeserializeAcceptActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.AcceptInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Activity":
			v, err := mgr.DeserializeActivityActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Add":
			v, err := mgr.DeserializeAddActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.AddInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Announce":
			v, err := mgr.DeserializeAnnounceActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.AnnounceInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Application":
			v, err := mgr.DeserializeApplicationActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ApplicationInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Arrive":
			v, err := mgr.DeserializeArriveActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ArriveInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Article":
			v, err := mgr.DeserializeArticleActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ArticleInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Audio":
			v, err := mgr.DeserializeAudioActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.AudioInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Block":
			v, err := mgr.DeserializeBlockActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.BlockInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Collection":
			v, err := mgr.DeserializeCollectionActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.CollectionInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "CollectionPage":
			v, err := mgr.DeserializeCollectionPageActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.CollectionPageInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Create":
			v, err := mgr.DeserializeCreateActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.CreateInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Delete":
			v, err := mgr.DeserializeDeleteActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.DeleteInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Dislike":
			v, err := mgr.DeserializeDislikeActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.DislikeInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Document":
			v, err := mgr.DeserializeDocumentActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.DocumentInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Event":
			v, err := mgr.DeserializeEventActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.EventInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Flag":
			v, err := mgr.DeserializeFlagActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.FlagInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Follow":
			v, err := mgr.DeserializeFollowActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.FollowInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Group":
			v, err := mgr.DeserializeGroupActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.GroupInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Ignore":
			v, err := mgr.DeserializeIgnoreActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.IgnoreInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Image":
			v, err := mgr.DeserializeImageActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ImageInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "IntransitiveActivity":
			v, err := mgr.DeserializeIntransitiveActivityActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.IntransitiveActivityInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Invite":
			v, err := mgr.DeserializeInviteActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.InviteInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Join":
			v, err := mgr.DeserializeJoinActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.JoinInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Leave":
			v, err := mgr.DeserializeLeaveActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.LeaveInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Like":
			v, err := mgr.DeserializeLikeActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.LikeInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Link":
			v, err := mgr.DeserializeLinkActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.LinkInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Listen":
			v, err := mgr.DeserializeListenActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ListenInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Mention":
			v, err := mgr.DeserializeMentionActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.MentionInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Move":
			v, err := mgr.DeserializeMoveActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.MoveInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Note":
			v, err := mgr.DeserializeNoteActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.NoteInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Object":
			v, err := mgr.DeserializeObjectActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ObjectInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Offer":
			v, err := mgr.DeserializeOfferActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.OfferInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "OrderedCollection":
			v, err := mgr.DeserializeOrderedCollectionActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.OrderedCollectionInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "OrderedCollectionPage":
			v, err := mgr.DeserializeOrderedCollectionPageActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.OrderedCollectionPageInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Organization":
			v, err := mgr.DeserializeOrganizationActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.OrganizationInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Page":
			v, err := mgr.DeserializePageActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.PageInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Person":
			v, err := mgr.DeserializePersonActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.PersonInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Place":
			v, err := mgr.DeserializePlaceActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.PlaceInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Profile":
			v, err := mgr.DeserializeProfileActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ProfileInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Question":
			v, err := mgr.DeserializeQuestionActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.QuestionInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Read":
			v, err := mgr.DeserializeReadActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ReadInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Reject":
			v, err := mgr.DeserializeRejectActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.RejectInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Relationship":
			v, err := mgr.DeserializeRelationshipActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.RelationshipInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Remove":
			v, err := mgr.DeserializeRemoveActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.RemoveInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Service":
			v, err := mgr.DeserializeServiceActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ServiceInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "TentativeAccept":
			v, err := mgr.DeserializeTentativeAcceptActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.TentativeAcceptInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "TentativeReject":
			v, err := mgr.DeserializeTentativeRejectActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.TentativeRejectInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Tombstone":
			v, err := mgr.DeserializeTombstoneActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.TombstoneInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Travel":
			v, err := mgr.DeserializeTravelActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.TravelInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Undo":
			v, err := mgr.DeserializeUndoActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.UndoInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Update":
			v, err := mgr.DeserializeUpdateActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.UpdateInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "Video":
			v, err := mgr.DeserializeVideoActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.VideoInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		case "View":
			v, err := mgr.DeserializeViewActivityStreams()(m, aliasMap)

			if err != nil {
				return err
			}

			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ViewInterface) error); ok {
					return fn(ctx, v)
				}
			}

			return ErrNoCallbackMatch
		default:
			return ErrUnhandledType
		}
	}
	// End: Private lambda
	if typeStr, ok := typeValue.(string); ok {
		return handleFn(typeStr)
	} else if typeIArr, ok := typeValue.([]interface{}); ok {
		for _, typeI := range typeIArr {
			if typeStr, ok := typeI.(string); ok {
				if err := handleFn(typeStr); err == nil {
					return nil
				} else if err == ErrUnhandledType {
					// Keep trying other types: only if all fail do we return this error.
					continue
				} else {
					return err
				}
			}
		}
		return ErrUnhandledType
	} else {
		return ErrUnhandledType
	}
}

package pub

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-fed/activity/streams"
	"github.com/go-fed/activity/streams/vocab"
	"net/url"
)

// OnFollowBehavior enumerates the different default actions that the go-fed
// library can provide when receiving a Follow Activity from a peer.
type OnFollowBehavior int

const (
	// OnFollowDoNothing does not take any action when a Follow Activity
	// is received.
	OnFollowDoNothing OnFollowBehavior = iota
	// OnFollowAutomaticallyAccept triggers the side effect of sending an
	// Accept of this Follow request in response.
	OnFollowAutomaticallyAccept
	// OnFollowAutomaticallyAccept triggers the side effect of sending a
	// Reject of this Follow request in response.
	OnFollowAutomaticallyReject
)

// FederatingWrappedCallbacks lists the callback functions that already have
// some side effect behavior provided by the pub library.
//
// These functions are wrapped for the Federating Protocol.
type FederatingWrappedCallbacks struct {
	// Create handles additional side effects for the Create ActivityStreams
	// type, specific to the application using go-fed.
	//
	// The wrapping callback for the Federating Protocol ensures the
	// 'object' property is created in the database.
	//
	// Create calls Create for each object in the federated Activity.
	Create func(context.Context, vocab.ActivityStreamsCreate) error
	// Update handles additional side effects for the Update ActivityStreams
	// type, specific to the application using go-fed.
	//
	// The wrapping callback for the Federating Protocol ensures the
	// 'object' property is updated in the database.
	//
	// Update calls Update on the federated entry from the database, with a
	// new value.
	Update func(context.Context, vocab.ActivityStreamsUpdate) error
	// Delete handles additional side effects for the Delete ActivityStreams
	// type, specific to the application using go-fed.
	//
	// Delete removes the federated entry from the database.
	Delete func(context.Context, vocab.ActivityStreamsDelete) error
	// Follow handles additional side effects for the Follow ActivityStreams
	// type, specific to the application using go-fed.
	//
	// The wrapping function can have one of several default behaviors,
	// depending on the value of the OnFollow setting.
	Follow func(context.Context, vocab.ActivityStreamsFollow) error
	// OnFollow determines what action to take for this particular callback
	// if a Follow Activity is handled.
	OnFollow OnFollowBehavior
	// Accept handles additional side effects for the Accept ActivityStreams
	// type, specific to the application using go-fed.
	//
	// The wrapping function determines if this 'Accept' is in response to a
	// 'Follow'. If so, then the 'actor' is added to the original 'actor's
	// 'following' collection.
	Accept func(context.Context, vocab.ActivityStreamsAccept) error
	// Reject handles additional side effects for the Reject ActivityStreams
	// type, specific to the application using go-fed.
	//
	// The wrapping function does nothing. However, if this 'Reject' is in
	// response to a 'Follow' then the client MUST NOT go forward with
	// adding the 'actor' to the original 'actor's 'following' collection by
	// the client application.
	Reject func(context.Context, vocab.ActivityStreamsReject) error
	// Add handles additional side effects for the Add ActivityStreams
	// type, specific to the application using go-fed.
	//
	// The wrapping function will add the 'object' IRIs to a specific
	// 'target' collection if the 'target' collection(s) live on this
	// server.
	Add func(context.Context, vocab.ActivityStreamsAdd) error
	// Remove handles additional side effects for the Remove ActivityStreams
	// type, specific to the application using go-fed.
	//
	// The wrapping function will remove all 'object' IRIs from a specific
	// 'target' collection if the 'target' collection(s) live on this
	// server.
	Remove func(context.Context, vocab.ActivityStreamsRemove) error
	// Like handles additional side effects for the Like ActivityStreams
	// type, specific to the application using go-fed.
	//
	// The wrapping function will add the activity to the "likes" collection
	// on all 'object' targets owned by this server.
	Like func(context.Context, vocab.ActivityStreamsLike) error
	// Announce handles additional side effects for the Announce
	// ActivityStreams type, specific to the application using go-fed.
	//
	// The wrapping function will add the activity to the "shares"
	// collection on all 'object' targets owned by this server.
	Announce func(context.Context, vocab.ActivityStreamsAnnounce) error
	// Undo handles additional side effects for the Undo ActivityStreams
	// type, specific to the application using go-fed.
	//
	// The wrapping function ensures the 'actor' on the 'Undo'
	// is be the same as the 'actor' on all Activities being undone.
	// It enforces that the actors on the Undo must correspond to all of the
	// 'object' actors in some manner.
	Undo func(context.Context, vocab.ActivityStreamsUndo) error
	// Block handles additional side effects for the Block ActivityStreams
	// type, specific to the application using go-fed.
	//
	// The wrapping function does nothing. It simply calls this wrapped
	// function. However, note that Blocks should not be received from a
	// federated peer, as delivering Blocks explicitly deviates from the
	// original ActivityPub specification.
	Block func(context.Context, vocab.ActivityStreamsBlock) error

	// Sidechannel data -- this is set at request handling time. These must
	// be set before the callbacks are used.

	// db is the Database the FederatingWrappedCallbacks should use.
	db Database
	// inboxIRI is the inboxIRI that is handling this callback.
	// TODO: Populate
	inboxIRI *url.URL
	// newTransport creates a new Transport.
	// TODO: Populate
	newTransport func(c context.Context, actorBoxIRI *url.URL, gofedAgent string) (t Transport, err error)
}

// disjoint ensures that the functions given do not share a type signature with
// the functions being wrapped in FederatingWrappedCallbacks.
func (w FederatingWrappedCallbacks) disjoint(fns []interface{}) error {
	// TODO: Instead, if provided in "other" it should override this behavior.
	var s string
	for _, fn := range fns {
		switch fn.(type) {
		default:
			// OK, no collision
			continue
		case func(context.Context, vocab.ActivityStreamsCreate) error:
			s = "Create"
		case func(context.Context, vocab.ActivityStreamsUpdate) error:
			s = "Update"
		case func(context.Context, vocab.ActivityStreamsDelete) error:
			s = "Delete"
		case func(context.Context, vocab.ActivityStreamsFollow) error:
			s = "Follow"
		case func(context.Context, vocab.ActivityStreamsAccept) error:
			s = "Accept"
		case func(context.Context, vocab.ActivityStreamsReject) error:
			s = "Reject"
		case func(context.Context, vocab.ActivityStreamsAdd) error:
			s = "Add"
		case func(context.Context, vocab.ActivityStreamsRemove) error:
			s = "Remove"
		case func(context.Context, vocab.ActivityStreamsLike) error:
			s = "Like"
		case func(context.Context, vocab.ActivityStreamsAnnounce) error:
			s = "Announce"
		case func(context.Context, vocab.ActivityStreamsUndo) error:
			s = "Undo"
		case func(context.Context, vocab.ActivityStreamsBlock) error:
			s = "Block"
		}
		return fmt.Errorf("callback function handling type %q conflicts with FederatingWrappedCallbacks", s)
	}
	return nil
}

// callbacks returns the WrappedCallbacks members into a single interface slice
// for use in streams.Resolver callbacks.
func (w FederatingWrappedCallbacks) callbacks() []interface{} {
	return []interface{}{
		w.create,
		w.update,
		w.deleteFn,
		w.follow,
		w.accept,
		w.reject,
		w.add,
		w.remove,
		w.like,
		w.announce,
		w.undo,
		w.block,
	}
}

// create implements the federating Create activity side effects.
func (w FederatingWrappedCallbacks) create(c context.Context, a vocab.ActivityStreamsCreate) error {
	op := a.GetActivityStreamsObject()
	if op == nil || op.Len() == 0 {
		return ErrObjectRequired
	}
	for iter := op.Begin(); iter != op.End(); iter = iter.Next() {
		t := iter.GetType()
		if t == nil {
			// TODO: Dereference the IRI and store it.
			continue
		}
		id, err := GetId(t)
		if err != nil {
			return err
		}
		err = w.db.Lock(c, id)
		if err != nil {
			return err
		}
		// WARNING: Unlock is not deferred.
		if err := w.db.Create(c, t); err != nil {
			w.db.Unlock(c, id)
			return err
		}
		w.db.Unlock(c, id)
		// At this point, Unlock should be called and in every above
		// branch.
	}
	if w.Create != nil {
		return w.Create(c, a)
	}
	return nil
}

// update implements the federating Update activity side effects.
func (w FederatingWrappedCallbacks) update(c context.Context, a vocab.ActivityStreamsUpdate) error {
	op := a.GetActivityStreamsObject()
	if op == nil || op.Len() == 0 {
		return ErrObjectRequired
	}
	if err := mustHaveActivityOriginMatchObjects(a); err != nil {
		return err
	}
	for iter := op.Begin(); iter != op.End(); iter = iter.Next() {
		t := iter.GetType()
		if t == nil {
			return fmt.Errorf("update requires an object to be wholly provided")
		}
		id, err := GetId(t)
		if err != nil {
			return err
		}
		err = w.db.Lock(c, id)
		if err != nil {
			return err
		}
		// WARNING: Unlock is not deferred.
		if err := w.db.Update(c, t); err != nil {
			w.db.Unlock(c, id)
			return err
		}
		w.db.Unlock(c, id)
		// At this point, Unlock should be called and in every above
		// branch.
	}
	if w.Update != nil {
		return w.Update(c, a)
	}
	return nil
}

// deleteFn implements the federating Delete activity side effects.
func (w FederatingWrappedCallbacks) deleteFn(c context.Context, a vocab.ActivityStreamsDelete) error {
	op := a.GetActivityStreamsObject()
	if op == nil || op.Len() == 0 {
		return ErrObjectRequired
	}
	if err := mustHaveActivityOriginMatchObjects(a); err != nil {
		return err
	}
	for iter := op.Begin(); iter != op.End(); iter = iter.Next() {
		id, err := ToId(iter)
		if err != nil {
			return err
		}
		err = w.db.Lock(c, id)
		if err != nil {
			return err
		}
		// WARNING: Unlock is not deferred.
		if err := w.db.Delete(c, id); err != nil {
			w.db.Unlock(c, id)
			return err
		}
		w.db.Unlock(c, id)
		// At this point, Unlock should be called and in every above
		// branch.
	}
	if w.Delete != nil {
		return w.Delete(c, a)
	}
	return nil
}

// follow implements the federating Follow activity side effects.
func (w FederatingWrappedCallbacks) follow(c context.Context, a vocab.ActivityStreamsFollow) error {
	op := a.GetActivityStreamsObject()
	if op == nil || op.Len() == 0 {
		return ErrObjectRequired
	}
	// Check that we own at least one of the 'object' properties, and ensure
	// it is to the actor that owns this inbox.
	//
	// If not then don't send a response. It was federated to us as an FYI,
	// by mistake, or some other reason.
	if err := w.db.Lock(c, w.inboxIRI); err != nil {
		return err
	}
	// WARNING: Unlock not deferred.
	actorIRI, err := w.db.ActorForInbox(c, w.inboxIRI)
	if err != nil {
		w.db.Unlock(c, w.inboxIRI)
		return err
	}
	w.db.Unlock(c, w.inboxIRI)
	// Unlock must be called by now and every branch above.
	isMe := false
	if w.OnFollow != OnFollowDoNothing {
		for iter := op.Begin(); iter != op.End(); iter = iter.Next() {
			id, err := ToId(iter)
			if err != nil {
				return err
			}
			if id.String() == actorIRI.String() {
				isMe = true
				break
			}
		}
	}
	if isMe {
		// Prepare the response.
		var response Activity
		if w.OnFollow == OnFollowAutomaticallyAccept {
			response = streams.NewActivityStreamsAccept()
		} else if w.OnFollow == OnFollowAutomaticallyReject {
			response = streams.NewActivityStreamsReject()
		} else {
			return fmt.Errorf("unknown OnFollowBehavior: %d", w.OnFollow)
		}
		// Set us as the 'actor'.
		me := streams.NewActivityStreamsActorProperty()
		response.SetActivityStreamsActor(me)
		me.AppendIRI(actorIRI)
		// Set the Follow as the 'object' property.
		op := streams.NewActivityStreamsObjectProperty()
		response.SetActivityStreamsObject(op)
		op.AppendActivityStreamsFollow(a)
		// Add all actors on the original Follow to the 'to' property.
		recipients := make([]*url.URL, 0)
		to := streams.NewActivityStreamsToProperty()
		response.SetActivityStreamsTo(to)
		followActors := a.GetActivityStreamsActor()
		for iter := followActors.Begin(); iter != followActors.End(); iter = iter.Next() {
			id, err := ToId(iter)
			if err != nil {
				return err
			}
			to.AppendIRI(id)
			recipients = append(recipients, id)
		}
		if w.OnFollow == OnFollowAutomaticallyAccept {
			// If automatically accepting, then also update our
			// followers collection with the new actors.
			//
			// If automatically rejecting, do not update the
			// followers collection.
			if err := w.db.Lock(c, actorIRI); err != nil {
				return err
			}
			// WARNING: Unlock not deferred.
			followers, err := w.db.Followers(c, actorIRI)
			if err != nil {
				w.db.Unlock(c, actorIRI)
				return err
			}
			items := followers.GetActivityStreamsItems()
			for _, elem := range recipients {
				items.PrependIRI(elem)
			}
			if err = w.db.Update(c, followers); err != nil {
				w.db.Unlock(c, actorIRI)
				return err
			}
			w.db.Unlock(c, actorIRI)
			// Unlock must be called by now and every branch above.
		}
		m, err := serialize(response)
		if err != nil {
			return err
		}
		b, err := json.Marshal(m)
		if err != nil {
			return err
		}
		t, err := w.newTransport(c, w.inboxIRI, goFedUserAgent())
		if err != nil {
			return err
		}
		if err := t.BatchDeliver(c, b, recipients); err != nil {
			return err
		}
	}
	if w.Follow != nil {
		return w.Follow(c, a)
	}
	return nil
}

// accept implements the federating Accept activity side effects.
func (w FederatingWrappedCallbacks) accept(c context.Context, a vocab.ActivityStreamsAccept) error {
	op := a.GetActivityStreamsObject()
	if op != nil && op.Len() > 0 {
		// Get this actor's id.
		if err := w.db.Lock(c, w.inboxIRI); err != nil {
			return err
		}
		// WARNING: Unlock not deferred.
		actorIRI, err := w.db.ActorForInbox(c, w.inboxIRI)
		if err != nil {
			w.db.Unlock(c, w.inboxIRI)
			return err
		}
		w.db.Unlock(c, w.inboxIRI)
		// Unlock must be called by now and every branch above.
		//
		// Determine if we are in a follow on the 'object' property.
		isMe := false
		for iter := op.Begin(); iter != op.End(); iter = iter.Next() {
			t := iter.GetType()
			if t == nil {
				// TODO: Fetch by IRI
				continue
			}
			// Ensure it is a Follow.
			if !streams.ActivityStreamsFollowIsExtendedBy(t) {
				continue
			}
			follow, ok := t.(Activity)
			if !ok {
				return fmt.Errorf("a Follow in an Accept does not satisfy the Activity interface")
			}
			// Ensure that we are one of the actors on the Follow.
			actors := follow.GetActivityStreamsActor()
			for iter := actors.Begin(); iter != actors.End(); iter = iter.Next() {
				id, err := ToId(iter)
				if err != nil {
					return err
				}
				if id.String() == actorIRI.String() {
					isMe = true
					break
				}
			}
			// TODO: Double check and verify it exists.
			// Continue breaking if we found ourselves
			if isMe {
				break
			}
		}
		// If we received an Accept whose 'object' is a Follow with an
		// Accept that we sent, add to the following collection.
		if isMe {
			actors := a.GetActivityStreamsActor()
			if actors == nil || actors.Len() == 0 {
				return fmt.Errorf("an Accept with a Follow has no actors")
			}
			if err := w.db.Lock(c, actorIRI); err != nil {
				return err
			}
			// WARNING: Unlock not deferred.
			following, err := w.db.Following(c, actorIRI)
			if err != nil {
				w.db.Unlock(c, actorIRI)
				return err
			}
			items := following.GetActivityStreamsItems()
			for iter := actors.Begin(); iter != actors.End(); iter = iter.Next() {
				id, err := ToId(iter)
				if err != nil {
					w.db.Unlock(c, actorIRI)
					return err
				}
				items.PrependIRI(id)
			}
			if err = w.db.Update(c, following); err != nil {
				w.db.Unlock(c, actorIRI)
				return err
			}
			w.db.Unlock(c, actorIRI)
			// Unlock must be called by now and every branch above.
		}
	}
	if w.Accept != nil {
		return w.Accept(c, a)
	}
	return nil
}

// reject implements the federating Reject activity side effects.
func (w FederatingWrappedCallbacks) reject(c context.Context, a vocab.ActivityStreamsReject) error {
	if w.Reject != nil {
		return w.Reject(c, a)
	}
	return nil
}

// add implements the federating Add activity side effects.
func (w FederatingWrappedCallbacks) add(c context.Context, a vocab.ActivityStreamsAdd) error {
	op := a.GetActivityStreamsObject()
	if op == nil || op.Len() == 0 {
		return ErrObjectRequired
	}
	target := a.GetActivityStreamsTarget()
	if target == nil || target.Len() == 0 {
		return ErrTargetRequired
	}
	opIds := make([]*url.URL, 0, op.Len())
	for iter := op.Begin(); iter != op.End(); iter = iter.Next() {
		id, err := ToId(iter)
		if err != nil {
			return err
		}
		opIds = append(opIds, id)
	}
	targetIds := make([]*url.URL, 0, op.Len())
	for iter := target.Begin(); iter != target.End(); iter = iter.Next() {
		id, err := ToId(iter)
		if err != nil {
			return err
		}
		targetIds = append(targetIds, id)
	}
	for _, t := range targetIds {
		if err := w.db.Lock(c, t); err != nil {
			return err
		}
		// WARNING: Unlock not deferred.
		if owns, err := w.db.Owns(c, t); err != nil {
			w.db.Unlock(c, t)
			return err
		} else if !owns {
			w.db.Unlock(c, t)
			continue
		}
		tp, err := w.db.Get(c, t)
		if err != nil {
			w.db.Unlock(c, t)
			return err
		}
		if streams.ActivityStreamsOrderedCollectionIsExtendedBy(tp) {
			oi, ok := tp.(orderedItemser)
			if !ok {
				w.db.Unlock(c, t)
				return fmt.Errorf("type extending from OrderedCollection cannot convert to orderedItemser interface")
			}
			oiProp := oi.GetActivityStreamsOrderedItems()
			if oiProp == nil {
				oiProp = streams.NewActivityStreamsOrderedItemsProperty()
				oi.SetActivityStreamsOrderedItems(oiProp)
			}
			for _, objId := range opIds {
				oiProp.AppendIRI(objId)
			}
		} else if streams.ActivityStreamsCollectionIsExtendedBy(tp) {
			i, ok := tp.(itemser)
			if !ok {
				w.db.Unlock(c, t)
				return fmt.Errorf("type extending from Collection cannot convert to itemser interface")
			}
			iProp := i.GetActivityStreamsItems()
			if iProp == nil {
				iProp = streams.NewActivityStreamsItemsProperty()
				i.SetActivityStreamsItems(iProp)
			}
			for _, objId := range opIds {
				iProp.AppendIRI(objId)
			}
		}
		err = w.db.Update(c, tp)
		if err != nil {
			w.db.Unlock(c, t)
			return err
		}
		w.db.Unlock(c, t)
		// Unlock must be called by now and every branch above.
	}
	if w.Add != nil {
		return w.Add(c, a)
	}
	return nil
}

// remove implements the federating Remove activity side effects.
func (w FederatingWrappedCallbacks) remove(c context.Context, a vocab.ActivityStreamsRemove) error {
	op := a.GetActivityStreamsObject()
	if op == nil || op.Len() == 0 {
		return ErrObjectRequired
	}
	target := a.GetActivityStreamsTarget()
	if target == nil || target.Len() == 0 {
		return ErrTargetRequired
	}
	opIds := make(map[string]bool, op.Len())
	for iter := op.Begin(); iter != op.End(); iter = iter.Next() {
		id, err := ToId(iter)
		if err != nil {
			return err
		}
		opIds[id.String()] = true
	}
	targetIds := make([]*url.URL, 0, op.Len())
	for iter := target.Begin(); iter != target.End(); iter = iter.Next() {
		id, err := ToId(iter)
		if err != nil {
			return err
		}
		targetIds = append(targetIds, id)
	}
	for _, t := range targetIds {
		if err := w.db.Lock(c, t); err != nil {
			return err
		}
		// WARNING: Unlock not deferred.
		if owns, err := w.db.Owns(c, t); err != nil {
			w.db.Unlock(c, t)
			return err
		} else if !owns {
			w.db.Unlock(c, t)
			continue
		}
		tp, err := w.db.Get(c, t)
		if err != nil {
			w.db.Unlock(c, t)
			return err
		}
		if streams.ActivityStreamsOrderedCollectionIsExtendedBy(tp) {
			oi, ok := tp.(orderedItemser)
			if !ok {
				w.db.Unlock(c, t)
				return fmt.Errorf("type extending from OrderedCollection cannot convert to orderedItemser interface")
			}
			oiProp := oi.GetActivityStreamsOrderedItems()
			if oiProp != nil {
				for i := 0; i < oiProp.Len(); /*Conditional*/ {
					id, err := ToId(oiProp.At(i))
					if err != nil {
						w.db.Unlock(c, t)
						return err
					}
					if opIds[id.String()] {
						oiProp.Remove(i)
					} else {
						i++
					}
				}
			}
		} else if streams.ActivityStreamsCollectionIsExtendedBy(tp) {
			i, ok := tp.(itemser)
			if !ok {
				w.db.Unlock(c, t)
				return fmt.Errorf("type extending from Collection cannot convert to itemser interface")
			}
			iProp := i.GetActivityStreamsItems()
			if iProp != nil {
				for i := 0; i < iProp.Len(); /*Conditional*/ {
					id, err := ToId(iProp.At(i))
					if err != nil {
						w.db.Unlock(c, t)
						return err
					}
					if opIds[id.String()] {
						iProp.Remove(i)
					} else {
						i++
					}
				}
			}
		}
		err = w.db.Update(c, tp)
		if err != nil {
			w.db.Unlock(c, t)
			return err
		}
		w.db.Unlock(c, t)
		// Unlock must be called by now and every branch above.
	}
	if w.Remove != nil {
		return w.Remove(c, a)
	}
	return nil
}

// like implements the federating Like activity side effects.
func (w FederatingWrappedCallbacks) like(c context.Context, a vocab.ActivityStreamsLike) error {
	op := a.GetActivityStreamsObject()
	if op == nil || op.Len() == 0 {
		return ErrObjectRequired
	}
	id, err := GetId(a)
	if err != nil {
		return err
	}
	for iter := op.Begin(); iter != op.End(); iter = iter.Next() {
		objId, err := ToId(iter)
		if err != nil {
			return err
		}
		if err := w.db.Lock(c, objId); err != nil {
			return err
		}
		// WARNING: Unlock not deferred.
		if owns, err := w.db.Owns(c, objId); err != nil {
			w.db.Unlock(c, objId)
			return err
		} else if !owns {
			w.db.Unlock(c, objId)
			continue
		}
		t, err := w.db.Get(c, objId)
		if err != nil {
			w.db.Unlock(c, objId)
			return err
		}
		l, ok := t.(likeser)
		if !ok {
			w.db.Unlock(c, objId)
			return fmt.Errorf("cannot add Like to likes collection for type %T", t)
		}
		// Get 'likes' property on the object, creating default if
		// necessary.
		likes := l.GetActivityStreamsLikes()
		if likes == nil {
			likes = streams.NewActivityStreamsLikesProperty()
			l.SetActivityStreamsLikes(likes)
		}
		// Get 'likes' value, defaulting to a collection.
		likesT := likes.GetType()
		if likesT == nil {
			col := streams.NewActivityStreamsCollection()
			likesT = col
			likes.SetActivityStreamsCollection(col)
		}
		// Prepend the activity's 'id' on the 'likes' Collection or
		// OrderedCollection.
		if col, ok := likesT.(itemser); ok {
			items := col.GetActivityStreamsItems()
			if items == nil {
				items = streams.NewActivityStreamsItemsProperty()
				col.SetActivityStreamsItems(items)
			}
			items.PrependIRI(id)
		} else if oCol, ok := likesT.(orderedItemser); ok {
			oItems := oCol.GetActivityStreamsOrderedItems()
			if oItems == nil {
				oItems = streams.NewActivityStreamsOrderedItemsProperty()
				oCol.SetActivityStreamsOrderedItems(oItems)
			}
			oItems.PrependIRI(id)
		} else {
			w.db.Unlock(c, objId)
			return fmt.Errorf("likes type is neither a Collection nor an OrderedCollection: %T", likesT)
		}
		err = w.db.Update(c, t)
		if err != nil {
			w.db.Unlock(c, objId)
			return err
		}
		w.db.Unlock(c, objId)
		// Unlock must be called by now and every branch above.
	}
	if w.Like != nil {
		return w.Like(c, a)
	}
	return nil
}

// announce implements the federating Announce activity side effects.
func (w FederatingWrappedCallbacks) announce(c context.Context, a vocab.ActivityStreamsAnnounce) error {
	id, err := GetId(a)
	if err != nil {
		return err
	}
	op := a.GetActivityStreamsObject()
	for iter := op.Begin(); iter != op.End(); iter = iter.Next() {
		objId, err := ToId(iter)
		if err != nil {
			return err
		}
		if err := w.db.Lock(c, objId); err != nil {
			return err
		}
		// WARNING: Unlock not deferred.
		if owns, err := w.db.Owns(c, objId); err != nil {
			w.db.Unlock(c, objId)
			return err
		} else if !owns {
			w.db.Unlock(c, objId)
			continue
		}
		t, err := w.db.Get(c, objId)
		if err != nil {
			w.db.Unlock(c, objId)
			return err
		}
		s, ok := t.(shareser)
		if !ok {
			w.db.Unlock(c, objId)
			return fmt.Errorf("cannot add Announce to Shares collection for type %T", t)
		}
		// Get 'shares' property on the object, creating default if
		// necessary.
		shares := s.GetActivityStreamsShares()
		if shares == nil {
			shares = streams.NewActivityStreamsSharesProperty()
			s.SetActivityStreamsShares(shares)
		}
		// Get 'shares' value, defaulting to a collection.
		sharesT := shares.GetType()
		if sharesT== nil {
			col := streams.NewActivityStreamsCollection()
			sharesT = col
			shares.SetActivityStreamsCollection(col)
		}
		// Prepend the activity's 'id' on the 'shares' Collection or
		// OrderedCollection.
		if col, ok := sharesT.(itemser); ok {
			items := col.GetActivityStreamsItems()
			if items == nil {
				items = streams.NewActivityStreamsItemsProperty()
				col.SetActivityStreamsItems(items)
			}
			items.PrependIRI(id)
		} else if oCol, ok := sharesT.(orderedItemser); ok {
			oItems := oCol.GetActivityStreamsOrderedItems()
			if oItems == nil {
				oItems = streams.NewActivityStreamsOrderedItemsProperty()
				oCol.SetActivityStreamsOrderedItems(oItems)
			}
			oItems.PrependIRI(id)
		} else {
			w.db.Unlock(c, objId)
			return fmt.Errorf("shares type is neither a Collection nor an OrderedCollection: %T", sharesT)
		}
		err = w.db.Update(c, t)
		if err != nil {
			w.db.Unlock(c, objId)
			return err
		}
		w.db.Unlock(c, objId)
		// Unlock must be called by now and every branch above.
	}
	if w.Announce != nil {
		return w.Announce(c, a)
	}
	return nil
}

// undo implements the federating Undo activity side effects.
func (w FederatingWrappedCallbacks) undo(c context.Context, a vocab.ActivityStreamsUndo) error {
	op := a.GetActivityStreamsObject()
	if op == nil || op.Len() == 0 {
		return ErrObjectRequired
	}
	actors := a.GetActivityStreamsActor()
	activityActorMap := make(map[string]bool, actors.Len())
	for iter := actors.Begin(); iter != actors.End(); iter = iter.Next() {
		id, err := ToId(iter)
		if err != nil {
			return err
		}
		activityActorMap[id.String()] = true
	}
	for iter := op.Begin(); iter != op.End(); iter = iter.Next() {
		t := iter.GetType()
		if t == nil {
			// TODO: Fetch the IRI
			continue
		}
		ac, ok := t.(actorer)
		if !ok {
			return fmt.Errorf("cannot undo an object with no 'actor' property")
		}
		objActors := ac.GetActivityStreamsActor()
		for iter := objActors.Begin(); iter != objActors.End(); iter = iter.Next() {
			id, err := ToId(iter)
			if err != nil {
				return err
			}
			if !activityActorMap[id.String()] {
				return fmt.Errorf("activity Undoing another does not have all actors on original activities")
			}
		}
	}
	if w.Undo != nil {
		return w.Undo(c, a)
	}
	return nil
}

// block implements the federating Block activity side effects.
func (w FederatingWrappedCallbacks) block(c context.Context, a vocab.ActivityStreamsBlock) error {
	op := a.GetActivityStreamsObject()
	if op == nil || op.Len() == 0 {
		return ErrObjectRequired
	}
	if w.Block != nil {
		return w.Block(c, a)
	}
	return nil
}

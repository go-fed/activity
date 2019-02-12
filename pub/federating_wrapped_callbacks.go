package pub

import (
	"context"
	"fmt"
	"github.com/go-fed/activity/streams/vocab"
)

// OnFollowBehavior enumerates the different default actions that the go-fed
// library can provide when receiving a Follow Activity from a peer.
type OnFollowBehavior int

const (
	// OnFollowDoNothing does not take any action when a Follow Activity
	// is received.
	OnFollowDoNothing OnFollowBehavior = iota
	// OnFollowAutomaticallyAccept will trigger the side effect of sending
	// an Accept of this Follow request in response.
	OnFollowAutomaticallyAccept
	// OnFollowAutomaticallyAccept will trigger the side effect of sending
	// a Reject of this Follow request in response.
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
	// The wrapping callback for the Federating Protocol will ensure the
	// 'object' property is created in the database.
	//
	// TODO: Toggle this behavior locally
	Create func(context.Context, vocab.ActivityStreamsCreate) error
	// Update handles additional side effects for the Update ActivityStreams
	// type, specific to the application using go-fed.
	//
	// The wrapping callback for the Federating Protocol will ensure the
	// 'object' property is updated in the database.
	//
	// TODO: Toggle this behavior locally
	Update func(context.Context, vocab.ActivityStreamsUpdate) error
	// Delete handles additional side effects for the Delete ActivityStreams
	// type, specific to the application using go-fed.
	//
	// TODO: Describe
	Delete func(context.Context, vocab.ActivityStreamsDelete) error
	// Follow handles additional side effects for the Follow ActivityStreams
	// type, specific to the application using go-fed.
	//
	// TODO: Describe
	Follow func(context.Context, vocab.ActivityStreamsFollow) error
	// OnFollow determines what action to take for this particular callback
	// if a Follow Activity is handled.
	OnFollow OnFollowBehavior
	// Accept handles additional side effects for the Accept ActivityStreams
	// type, specific to the application using go-fed.
	//
	// TODO: Describe
	Accept func(context.Context, vocab.ActivityStreamsAccept) error
	// Reject handles additional side effects for the Reject ActivityStreams
	// type, specific to the application using go-fed.
	//
	// TODO: Describe
	Reject func(context.Context, vocab.ActivityStreamsReject) error
	// Add handles additional side effects for the Add ActivityStreams
	// type, specific to the application using go-fed.
	//
	// TODO: Describe
	Add func(context.Context, vocab.ActivityStreamsAdd) error
	// Remove handles additional side effects for the Remove ActivityStreams
	// type, specific to the application using go-fed.
	//
	// TODO: Describe
	Remove func(context.Context, vocab.ActivityStreamsRemove) error
	// Like handles additional side effects for the Like ActivityStreams
	// type, specific to the application using go-fed.
	//
	// TODO: Describe
	Like func(context.Context, vocab.ActivityStreamsLike) error
	// Announce handles additional side effects for the Announce
	// ActivityStreams type, specific to the application using go-fed.
	//
	// TODO: Describe
	Announce func(context.Context, vocab.ActivityStreamsAnnounce) error
	// Undo handles additional side effects for the Undo ActivityStreams
	// type, specific to the application using go-fed.
	//
	// TODO: Describe
	Undo func(context.Context, vocab.ActivityStreamsUndo) error
	// Block handles additional side effects for the Block ActivityStreams
	// type, specific to the application using go-fed.
	//
	// TODO: Describe
	Block func(context.Context, vocab.ActivityStreamsBlock) error

	// Sidechannel data -- this is set at request handling time. These must
	// be set before the callbacks are used.

	// db is the Database the FederatingWrappedCallbacks should use.
	db Database
}

// disjoint ensures that the functions given do not share a type signature with
// the functions being wrapped in FederatingWrappedCallbacks.
func (w FederatingWrappedCallbacks) disjoint(fns []interface{}) error {
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
	respond := false
	if w.OnFollow != OnFollowDoNothing {
		// Check that we own at least one of the 'object' properties. If
		// not then don't send a response. It was federated to us as an
		// FYI.
		// TODO: Check that this actor is the object.
		for iter := op.Begin(); iter != op.End(); iter = iter.Next() {
			id, err := ToId(iter)
			if err != nil {
				return err
			}
			if w.db.Owns(c, id) {
				respond = true
				break
			}
		}
	}
	if respond {
		// Prepare the response.
		var response Activity
		if w.OnFollow == OnFollowAutomaticallyAccept {
			response = streams.NewActivityStreamsAccept()
		} else if w.OnFollow == AutomaticReject {
			response = streams.NewActivityStreamsReject()
		} else {
			return fmt.Errorf("unknown OnFollowBehavior: %d", w.OnFollow)
		}
		// Set the Follow as the 'object' property.
		op := streams.NewActivityStreamsObjectProperty()
		response.SetActivityStreamsObject(op)
		op.AppendActivityStreamsFollow(a)
		// Add all actors on the original Follow to the 'to' property.
		to := streams.NewActivityStreamsToProperty()
		response.SetActivityStreamsTo(to)
		followActors := a.GetActivityStreamsActor()
		for iter := followActors.Begin(); iter != followActors.End(); iter = iter.Next() {
			id, err := ToId(iter)
			if err != nil {
				return err
			}
			to.AppendIRI(id)
		}
		ownsAny := false
		if todo == AutomaticAccept {
			// If automatically accepting, then also update our
			// followers collection.
			getter := func(object vocab.ObjectType, lc *vocab.CollectionType, loc *vocab.OrderedCollectionType) (bool, error) {
				if object.IsFollowersAnyURI() {
					pObj, err := f.App.Get(c, object.GetFollowersAnyURI(), ReadWrite)
					if err != nil {
						return true, err
					}
					ok := false
					if *lc, ok = pObj.(vocab.CollectionType); !ok {
						if *loc, ok = pObj.(vocab.OrderedCollectionType); !ok {
							return true, fmt.Errorf("object followers collection not CollectionType nor OrderedCollectionType")
						}
					}
					return true, nil
				} else if object.IsFollowersCollection() {
					*lc = object.GetFollowersCollection()
					return false, nil
				} else if object.IsFollowersOrderedCollection() {
					*loc = object.GetFollowersOrderedCollection()
					return false, nil
				}
				*loc = &vocab.OrderedCollection{}
				object.SetFollowersOrderedCollection(*loc)
				return false, nil
			}
			var err error
			if ownsAny, err = f.addAllActorsToObjectCollection(c, getter, raw, true); err != nil {
				return err
			}
		} else if todo == AutomaticReject {
			// If automatically rejecting, do not update the
			// followers collection.
			var err error
			ownsAny, err = f.ownsAnyObjects(c, raw)
			if err != nil {
				return err
			}
		}
		if ownsAny {
			if err := f.deliver(activity, inboxURL); err != nil {
				return err
			}
		}
	}
	return f.ServerCallbacker.Follow(c, s)
}

// accept implements the federating Accept activity side effects.
func (w FederatingWrappedCallbacks) accept(c context.Context, a vocab.ActivityStreamsAccept) error {
	// Accept can be client application specific. However, if this 'Accept'
	// is in response to a 'Follow' then the 'actor' should be added to the
	// original 'actor's 'following' collection by the client application.
	raw := s.Raw()
	for i := 0; i < raw.ObjectLen(); i++ {
		if raw.IsObject(i) {
			obj := raw.GetObject(i)
			follow, ok := obj.(vocab.FollowType) // TODO: Random audit: Make sure this is correct
			if !ok {
				continue
			}
			getter := func(actor vocab.ObjectType, lc *vocab.CollectionType, loc *vocab.OrderedCollectionType) (bool, error) {
				if actor.IsFollowingAnyURI() {
					pObj, err := f.App.Get(c, actor.GetFollowingAnyURI(), ReadWrite)
					if err != nil {
						return true, err
					}
					ok := false
					if *lc, ok = pObj.(vocab.CollectionType); !ok {
						if *loc, ok = pObj.(vocab.OrderedCollectionType); !ok {
							return true, fmt.Errorf("actors following collection not CollectionType nor OrderedCollectionType")
						}
					}
					return true, nil
				} else if actor.IsFollowingCollection() {
					*lc = actor.GetFollowingCollection()
					return false, nil
				} else if actor.IsFollowingOrderedCollection() {
					*loc = actor.GetFollowingOrderedCollection()
					return false, nil
				}
				*loc = &vocab.OrderedCollection{}
				actor.SetFollowingOrderedCollection(*loc)
				return false, nil
			}
			if err := f.addAllObjectsToActorCollection(c, getter, follow, true); err != nil {
				return err
			}
		}
	}
	return f.ServerCallbacker.Accept(c, s)
}

// reject implements the federating Reject activity side effects.
func (w FederatingWrappedCallbacks) reject(c context.Context, a vocab.ActivityStreamsReject) error {
	// Reject can be client application specific. However, if this 'Reject'
	// is in response to a 'Follow' then the client MUST NOT go forward with
	// adding the 'actor' to the original 'actor's 'following' collection
	// by the client application.
	return f.ServerCallbacker.Reject(c, s)
}

// add implements the federating Add activity side effects.
func (w FederatingWrappedCallbacks) add(c context.Context, a vocab.ActivityStreamsAdd) error {
	// Add is client application specific, generally involving adding an
	// 'object' to a specific 'target' collection.
	if s.LenObject() == 0 {
		return errObjectRequired
	} else if s.LenTarget() == 0 {
		return errTargetRequired
	}
	raw := s.Raw()
	ids, err := getTargetIds(raw)
	if err != nil {
		return err
	} else if len(ids) == 0 {
		return fmt.Errorf("add target has no ids: %v", s)
	}
	objIds, err := getObjectIds(s.Raw())
	if err != nil {
		return err
	} else if len(objIds) == 0 {
		return fmt.Errorf("add object has no ids: %v", s)
	}
	var targets []vocab.ObjectType
	for _, id := range ids {
		if !f.App.Owns(c, id) {
			continue
		}
		target, err := f.App.Get(c, id, ReadWrite)
		if err != nil {
			return err
		}
		ct, okCollection := target.(vocab.CollectionType)
		oct, okOrdered := target.(vocab.OrderedCollectionType)
		if okCollection {
			targets = append(targets, ct)
		} else if okOrdered {
			targets = append(targets, oct)
		}
		// else ignore non-collection, let Application handle.
	}
	for i := 0; i < raw.ObjectLen(); i++ {
		var obj vocab.ObjectType
		var objId *url.URL
		if raw.IsObject(i) {
			obj = raw.GetObject(i)
			if !obj.HasId() {
				return fmt.Errorf("add object missing id")
			}
			objId = obj.GetId()
		} else if raw.IsObjectIRI(i) {
			// TODO: Fetch IRI
			return fmt.Errorf("add object must not be IRI")
		}
		for _, target := range targets {
			if !f.App.CanAdd(c, obj, target) {
				continue
			}
			if ct, ok := target.(vocab.CollectionType); ok {
				ct.AppendItemsIRI(objId)
			} else if oct, ok := target.(vocab.OrderedCollectionType); ok {
				oct.AppendOrderedItemsIRI(objId)
			}
			if err := f.App.Set(c, target); err != nil {
				return err
			}
		}
	}
	return f.ServerCallbacker.Add(c, s)
}

// remove implements the federating Remove activity side effects.
func (w FederatingWrappedCallbacks) remove(c context.Context, a vocab.ActivityStreamsRemove) error {
	// Remove is client application specific, generally involving removing
	// an 'object' from a specific 'target' collection.
	if s.LenObject() == 0 {
		return errObjectRequired
	} else if s.LenTarget() == 0 {
		return errTargetRequired
	}
	raw := s.Raw()
	ids, err := getTargetIds(raw)
	if err != nil {
		return err
	} else if len(ids) == 0 {
		return fmt.Errorf("remove target has no ids: %v", s)
	}
	objIds, err := getObjectIds(s.Raw())
	if err != nil {
		return err
	} else if len(objIds) == 0 {
		return fmt.Errorf("remove object has no ids: %v", s)
	}
	var targets []vocab.ObjectType
	for _, id := range ids {
		if !f.App.Owns(c, id) {
			continue
		}
		target, err := f.App.Get(c, id, ReadWrite)
		if err != nil {
			return err
		}
		ct, okCollection := target.(vocab.CollectionType)
		oct, okOrdered := target.(vocab.OrderedCollectionType)
		if okCollection {
			targets = append(targets, ct)
		} else if okOrdered {
			targets = append(targets, oct)
		}
		// else ignore non-collection, let Application handle.
	}
	for i := 0; i < raw.ObjectLen(); i++ {
		if !raw.IsObject(i) {
			// TODO: Fetch IRIs as well
			return fmt.Errorf("remove object must be object type: %v", raw)
		}
		obj := raw.GetObject(i)
		for _, target := range targets {
			if !f.App.CanRemove(c, obj, target) {
				continue
			}
			if ct, ok := target.(vocab.CollectionType); ok {
				removeCollectionItemWithId(ct, obj.GetId())
			} else if oct, ok := target.(vocab.OrderedCollectionType); ok {
				removeOrderedCollectionItemWithId(oct, obj.GetId())
			}
			if err := f.App.Set(c, target); err != nil {
				return err
			}
		}
	}
	return f.ServerCallbacker.Remove(c, s)
}

// like implements the federating Like activity side effects.
func (w FederatingWrappedCallbacks) like(c context.Context, a vocab.ActivityStreamsLike) error {
	if s.LenObject() == 0 {
		return errObjectRequired
	}
	getter := func(object vocab.ObjectType, lc *vocab.CollectionType, loc *vocab.OrderedCollectionType) (bool, error) {
		if object.IsLikesAnyURI() {
			pObj, err := f.App.Get(c, object.GetLikesAnyURI(), ReadWrite)
			if err != nil {
				return true, err
			}
			ok := false
			if *lc, ok = pObj.(vocab.CollectionType); !ok {
				if *loc, ok = pObj.(vocab.OrderedCollectionType); !ok {
					return true, fmt.Errorf("object likes collection not CollectionType nor OrderedCollectionType")
				}
			}
			return true, nil
		} else if object.IsLikesCollection() {
			*lc = object.GetLikesCollection()
			return false, nil
		} else if object.IsLikesOrderedCollection() {
			*loc = object.GetLikesOrderedCollection()
			return false, nil
		}
		*loc = &vocab.OrderedCollection{}
		object.SetLikesOrderedCollection(*loc)
		return false, nil
	}
	if _, err := f.addActivityToObjectCollection(c, getter, s.Raw(), true); err != nil {
		return err
	}
	return f.ServerCallbacker.Like(c, s)
}

// announce implements the federating Announce activity side effects.
func (w FederatingWrappedCallbacks) announce(c context.Context, a vocab.ActivityStreamsAnnounce) error {
	getter := func(object vocab.ObjectType, lc *vocab.CollectionType, loc *vocab.OrderedCollectionType) (bool, error) {
		if object.IsSharesAnyURI() {
			pObj, err := f.App.Get(c, object.GetSharesAnyURI(), ReadWrite)
			if err != nil {
				return true, err
			}
			ok := false
			if *lc, ok = pObj.(vocab.CollectionType); !ok {
				if *loc, ok = pObj.(vocab.OrderedCollectionType); !ok {
					return true, fmt.Errorf("object shares collection not CollectionType nor OrderedCollectionType")
				}
			}
			return true, nil
		} else if object.IsSharesCollection() {
			*lc = object.GetSharesCollection()
			return false, nil
		} else if object.IsSharesOrderedCollection() {
			*loc = object.GetSharesOrderedCollection()
			return false, nil
		}
		*loc = &vocab.OrderedCollection{}
		object.SetSharesOrderedCollection(*loc)
		return false, nil
	}
	if _, err := f.addActivityToObjectCollection(c, getter, s.Raw(), true); err != nil {
		return err
	}
	if t, ok := f.ServerCallbacker.(callbackerAnnounce); ok {
		return t.Announce(c, s)
	}
	return nil
}

// undo implements the federating Undo activity side effects.
func (w FederatingWrappedCallbacks) undo(c context.Context, a vocab.ActivityStreamsUndo) error {
	// Undo negates a previous action. The 'actor' on the 'Undo'
	// MUST be the same as the 'actor' on the Activity being undone.
	// Here we enforce that the actors on the Undo must correspond
	// to all objects' original actors in some manner.
	if s.LenObject() == 0 {
		return errObjectRequired
	}
	raw := s.Raw()
	if err := f.ensureActivityActorsMatchObjectActors(raw); err != nil {
		return err
	}
	return f.ServerCallbacker.Undo(c, s)
}

// block implements the federating Block activity side effects.
func (w FederatingWrappedCallbacks) block(c context.Context, a vocab.ActivityStreamsBlock) error {
	// Servers SHOULD NOT deliver Block Activities to their object. So in
	// this case we will explicitly ignore it, but validate it as if we
	// were to accept it.
	if s.LenObject() == 0 {
		return errObjectRequired
	}
	return nil
}

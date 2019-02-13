package pub

import (
	"context"
	"fmt"
	"github.com/go-fed/activity/streams"
	"github.com/go-fed/activity/streams/vocab"
	"net/url"
)

// SocialWrappedCallbacks lists the callback functions that already have some
// side effect behavior provided by the pub library.
//
// These functions are wrapped for the Social Protocol.
type SocialWrappedCallbacks struct {
	// Create handles additional side effects for the Create ActivityStreams
	// type.
	//
	// The wrapping callback for the Social Protocol copies the actor(s) to
	// the 'attributedTo' property, copying recipients between the Create
	// activity and all objects. It then saves the entry in the database.
	Create func(context.Context, vocab.ActivityStreamsCreate) error
	// Update handles additional side effects for the Update ActivityStreams
	// type.
	//
	// TODO: Describe
	Update func(context.Context, vocab.ActivityStreamsUpdate) error
	// Delete handles additional side effects for the Delete ActivityStreams
	// type.
	//
	// TODO: Describe
	Delete func(context.Context, vocab.ActivityStreamsDelete) error
	// Follow handles additional side effects for the Follow ActivityStreams
	// type.
	//
	// TODO: Describe
	Follow func(context.Context, vocab.ActivityStreamsFollow) error
	// Add handles additional side effects for the Add ActivityStreams
	// type.
	//
	// TODO: Describe
	Add func(context.Context, vocab.ActivityStreamsAdd) error
	// Remove handles additional side effects for the Remove ActivityStreams
	// type.
	//
	// TODO: Describe
	Remove func(context.Context, vocab.ActivityStreamsRemove) error
	// Like handles additional side effects for the Like ActivityStreams
	// type.
	//
	// TODO: Describe
	Like func(context.Context, vocab.ActivityStreamsLike) error
	// Undo handles additional side effects for the Undo ActivityStreams
	// type.
	//
	// TODO: Describe
	Undo func(context.Context, vocab.ActivityStreamsUndo) error
	// Block handles additional side effects for the Block ActivityStreams
	// type.
	//
	// TODO: Describe
	Block func(context.Context, vocab.ActivityStreamsBlock) error

	// Sidechannel data -- this is set at request handling time. These must
	// be set before the callbacks are used.

	// db is the Database the SocialWrappedCallbacks should use. It must be
	// set before calling the callbacks.
	db Database
	// deliverable is a sidechannel out, indicating if the handled activity
	// should be delivered to a peer.
	deliverable bool
}

// disjoint ensures that the functions given do not share a type signature with
// the functions being wrapped in SocialWrappedCallbacks.
func (w SocialWrappedCallbacks) disjoint(fns []interface{}) error {
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
		case func(context.Context, vocab.ActivityStreamsAdd) error:
			s = "Add"
		case func(context.Context, vocab.ActivityStreamsRemove) error:
			s = "Remove"
		case func(context.Context, vocab.ActivityStreamsLike) error:
			s = "Like"
		case func(context.Context, vocab.ActivityStreamsUndo) error:
			s = "Undo"
		case func(context.Context, vocab.ActivityStreamsBlock) error:
			s = "Block"
		}
		return fmt.Errorf("callback function handling type %q conflicts with SocialWrappedCallbacks", s)
	}
	return nil
}

// callbacks returns the WrappedCallbacks members into a single interface slice
// for use in streams.Resolver callbacks.
func (w SocialWrappedCallbacks) callbacks() []interface{} {
	return []interface{}{
		w.create,
		w.update,
		w.deleteFn,
		w.follow,
		w.add,
		w.remove,
		w.like,
		w.undo,
		w.block,
	}
}

// create implements the social Create activity side effects.
func (w SocialWrappedCallbacks) create(c context.Context, a vocab.ActivityStreamsCreate) error {
	w.deliverable = true
	op := a.GetActivityStreamsObject()
	if op == nil || op.Len() == 0 {
		return ErrObjectRequired
	}
	// Obtain all actor IRIs.
	actors := a.GetActivityStreamsActor()
	createActorIds := make(map[string]*url.URL, actors.Len())
	for iter := actors.Begin(); iter != actors.End(); iter = iter.Next() {
		id, err := ToId(iter)
		if err != nil {
			return err
		}
		createActorIds[id.String()] = id
	}
	// Obtain each object's 'attributedTo' IRIs.
	objectAttributedToIds := make([]map[string]*url.URL, op.Len())
	for i := range objectAttributedToIds {
		objectAttributedToIds[i] = make(map[string]*url.URL)
	}
	for i := 0; i < op.Len(); i++ {
		t := op.At(i).GetType()
		attrToer, ok := t.(attributedToer)
		if !ok {
			continue
		}
		attr := attrToer.GetActivityStreamsAttributedTo()
		if attr == nil {
			attr = streams.NewActivityStreamsAttributedToProperty()
			attrToer.SetActivityStreamsAttributedTo(attr)
		}
		for iter := attr.Begin(); iter != attr.End(); iter = iter.Next() {
			id, err := ToId(iter)
			if err != nil {
				return err
			}
			objectAttributedToIds[i][id.String()] = id
		}
	}
	// Put all missing actor IRIs onto all object attributedTo properties.
	for k, v := range createActorIds {
		for i, attributedToMap := range objectAttributedToIds {
			if _, ok := attributedToMap[k]; !ok {
				t := op.At(i).GetType()
				attrToer, ok := t.(attributedToer)
				if !ok {
					continue
				}
				attr := attrToer.GetActivityStreamsAttributedTo()
				attr.AppendIRI(v)
			}
		}
	}
	// Put all missing object attributedTo IRIs onto the actor property.
	for _, attributedToMap := range objectAttributedToIds {
		for k, v := range attributedToMap {
			if _, ok := createActorIds[k]; !ok {
				actors.AppendIRI(v)
			}
		}
	}
	// Copy over the 'to', 'bto', 'cc', 'bcc', and 'audience' recipients
	// between the activity and all child objects and vice versa.
	if err := normalizeRecipients(a); err != nil {
		return err
	}
	// Persist all objects we've created, which will include sensitive
	// recipients such as 'bcc' and 'bto'.
	for i := 0; i < op.Len(); i++ {
		obj := op.At(i).GetType()
		// TODO: Lock
		if err := w.db.Create(c, obj); err != nil {
			return err
		}
	}
	if w.Create != nil {
		return w.Create(c, a)
	}
	return nil
}

// update implements the social Update activity side effects.
func (w SocialWrappedCallbacks) update(c context.Context, a vocab.ActivityStreamsUpdate) error {
	*deliverable = true
	if s.LenObject() == 0 {
		return errObjectRequired
	}
	// Update should partially replace the 'object' with only the
	// changed top-level fields.
	ids, err := getObjectIds(s.Raw())
	if err != nil {
		return err
	} else if len(ids) == 0 {
		return fmt.Errorf("update has no id: %v", s)
	}
	for idx, id := range ids {
		pObj, err := f.App.Get(c, id, ReadWrite)
		if err != nil {
			return err
		}
		m, err := pObj.Serialize()
		if err != nil {
			return err
		}
		if !s.Raw().IsObject(idx) {
			return fmt.Errorf("update requires object to be wholly provided at index %d", idx)
		}
		updated, err := s.Raw().GetObject(idx).Serialize()
		if err != nil {
			return err
		}
		for k, v := range updated {
			m[k] = v
		}
		if rawUpdatedObject := getRawObject(rawJson, id.String()); rawUpdatedObject != nil {
			recursivelyApplyDeletes(m, rawUpdatedObject)
		}
		p, err := ToPubObject(m)
		if err != nil {
			return err
		}
		for _, elem := range p {
			if err := f.App.Set(c, elem); err != nil {
				return err
			}
		}
	}
	return f.ClientCallbacker.Update(c, s)
}

// deleteFn implements the social Delete activity side effects.
func (w SocialWrappedCallbacks) deleteFn(c context.Context, a vocab.ActivityStreamsDelete) error {
	*deliverable = true
	if s.LenObject() == 0 {
		return errObjectRequired
	}
	ids, err := getObjectIds(s.Raw())
	if err != nil {
		return err
	} else if len(ids) == 0 {
		return fmt.Errorf("delete has no id: %v", s)
	}
	for _, id := range ids {
		pObj, err := f.App.Get(c, id, ReadWrite)
		if err != nil {
			return err
		}
		obj, ok := pObj.(vocab.ObjectType)
		if !ok {
			return fmt.Errorf("cannot delete non-ObjectType: %T", pObj)
		}
		tomb := toTombstone(obj, id, f.Clock.Now())
		if err := f.App.Set(c, tomb); err != nil {
			return err
		}
	}
	return f.ClientCallbacker.Delete(c, s)
}

// follow implements the social Follow activity side effects.
func (w SocialWrappedCallbacks) follow(c context.Context, a vocab.ActivityStreamsFollow) error {
	*deliverable = true
	if s.LenObject() == 0 {
		return errObjectRequired
	}
	return f.ClientCallbacker.Follow(c, s)
}

// add implements the social Add activity side effects.
func (w SocialWrappedCallbacks) add(c context.Context, a vocab.ActivityStreamsAdd) error {
	// TODO: Dedupe with FederatingWrappedCallbacks
	*deliverable = true
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
		if !okCollection && !okOrdered {
			return fmt.Errorf("cannot add to type that is not Collection and not OrderedCollection: %v", target)
		} else if okCollection {
			targets = append(targets, ct)
		} else {
			targets = append(targets, oct)
		}
	}
	for i := 0; i < raw.ObjectLen(); i++ {
		var obj vocab.ObjectType
		var objId *url.URL
		if raw.IsObjectIRI(i) {
			objId = raw.GetObjectIRI(i)
			if f.App.Owns(c, objId) {
				pObj, err := f.App.Get(c, objId, Read)
				var ok bool
				if obj, ok = pObj.(vocab.ObjectType); !ok {
					return fmt.Errorf("add object must be an activitypub object: %v", raw)
				}
				if err != nil {
					return err
				}
			} else {
				obj, err = f.dereferenceAsUser(outboxURL, objId)
				if err != nil {
					return err
				}
			}
		} else if raw.IsObject(i) {
			obj = raw.GetObject(i)
			if !obj.HasId() {
				return fmt.Errorf("add object missing iri")
			}
			objId = obj.GetId()
		} else {
			return fmt.Errorf("add object must be of object or iri type: %v", raw)
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
	return f.ClientCallbacker.Add(c, s)
}

// remove implements the social Remove activity side effects.
func (w SocialWrappedCallbacks) remove(c context.Context, a vocab.ActivityStreamsRemove) error {
	// TODO: Dedupe with FederatingWrappedCallbacks
	*deliverable = true
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
		if !okCollection && !okOrdered {
			return fmt.Errorf("cannot remove from type that is not Collection and not OrderedCollection: %v", target)
		} else if okCollection {
			targets = append(targets, ct)
		} else {
			targets = append(targets, oct)
		}
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
	return f.ClientCallbacker.Remove(c, s)
}

// like implements the social Like activity side effects.
func (w SocialWrappedCallbacks) like(c context.Context, a vocab.ActivityStreamsLike) error {
	*deliverable = true
	if s.LenObject() == 0 {
		return errObjectRequired
	}
	getter := func(actor vocab.ObjectType, lc *vocab.CollectionType, loc *vocab.OrderedCollectionType) (bool, error) {
		if actor.IsLikedAnyURI() {
			pObj, err := f.App.Get(ctx, actor.GetLikedAnyURI(), ReadWrite)
			if err != nil {
				return true, err
			}
			ok := false
			if *lc, ok = pObj.(vocab.CollectionType); !ok {
				if *loc, ok = pObj.(vocab.OrderedCollectionType); !ok {
					return true, fmt.Errorf("actors liked collection not CollectionType nor OrderedCollectionType")
				}
			}
			return true, nil
		} else if actor.IsLikedCollection() {
			*lc = actor.GetLikedCollection()
			return false, nil
		} else if actor.IsLikedOrderedCollection() {
			*loc = actor.GetLikedOrderedCollection()
			return false, nil
		}
		*loc = &vocab.OrderedCollection{}
		actor.SetLikedOrderedCollection(*loc)
		return false, nil
	}
	if err := f.addAllObjectsToActorCollection(ctx, getter, s.Raw(), true); err != nil {
		return err
	}
	return f.ClientCallbacker.Like(ctx, s)
}

// undo implements the social Undo activity side effects.
func (w SocialWrappedCallbacks) undo(c context.Context, a vocab.ActivityStreamsUndo) error {
	*deliverable = true
	if s.LenObject() == 0 {
		return errObjectRequired
	}
	raw := s.Raw()
	if err := f.ensureActivityActorsMatchObjectActors(raw); err != nil {
		return err
	}
	return f.ClientCallbacker.Undo(c, s)
}

// block implements the social Block activity side effects.
func (w SocialWrappedCallbacks) block(c context.Context, a vocab.ActivityStreamsBlock) error {
	*deliverable = false
	if s.LenObject() == 0 {
		return errObjectRequired
	}
	return f.ClientCallbacker.Block(c, s)
}

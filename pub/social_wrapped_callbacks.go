package pub

import (
	"context"
	"github.com/go-fed/activity/streams/vocab"
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
	// activity and all objects, save the entry in the database, and adds it
	// to the outbox.
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
	// db is the Database the SocialWrappedCallbacks should use. It must be
	// set before calling the callbacks.
	db Database
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
		w.accept,
		w.reject,
		w.add,
		w.remove,
		w.like,
		w.undo,
		w.block,
	}
}

// create implements the social Create activity side effects.
func (w SocialWrappedCallbacks) create(c context.Context, a vocab.ActivityStreamsCreate) error {
	*deliverable = true
	if s.LenObject() == 0 {
		return errObjectRequired
	}
	c := s.Raw()
	// When a Create activity is posted, the actor of the activity
	// SHOULD be copied onto the object's attributedTo field.
	// Presumably only if it doesn't already exist, to prevent
	// duplicate deliveries.
	createActorIds := make(map[string]interface{})
	for i := 0; i < c.ActorLen(); i++ {
		if c.IsActorObject(i) {
			obj := c.GetActorObject(i)
			id := obj.GetId()
			createActorIds[id.String()] = obj
		} else if c.IsActorLink(i) {
			l := c.GetActorLink(i)
			href := l.GetHref()
			createActorIds[href.String()] = l
		} else if c.IsActorIRI(i) {
			iri := c.GetActorIRI(i)
			createActorIds[iri.String()] = iri
		}
	}
	var obj []vocab.ObjectType
	for i := 0; i < c.ObjectLen(); i++ {
		if !c.IsObject(i) {
			return fmt.Errorf("unsupported: Create Activity with 'object' that is only an IRI")
		}
		obj = append(obj, c.GetObject(i))
	}
	objectAttributedToIds := make([]map[string]interface{}, len(obj))
	for i := range objectAttributedToIds {
		objectAttributedToIds[i] = make(map[string]interface{})
	}
	for k, o := range obj {
		for i := 0; i < o.AttributedToLen(); i++ {
			if o.IsAttributedToObject(i) {
				at := o.GetAttributedToObject(i)
				id := o.GetId()
				objectAttributedToIds[k][id.String()] = at
			} else if o.IsAttributedToLink(i) {
				at := o.GetAttributedToLink(i)
				href := at.GetHref()
				objectAttributedToIds[k][href.String()] = at
			} else if o.IsAttributedToIRI(i) {
				iri := o.GetAttributedToIRI(i)
				objectAttributedToIds[k][iri.String()] = iri
			}
		}
	}
	for k, v := range createActorIds {
		for i, attributedToMap := range objectAttributedToIds {
			if _, ok := attributedToMap[k]; !ok {
				var iri *url.URL
				if vObj, ok := v.(vocab.ObjectType); ok {
					if !vObj.HasId() {
						return fmt.Errorf("create actor object missing id")
					}
					iri = vObj.GetId()
				} else if vLink, ok := v.(vocab.LinkType); ok {
					if !vLink.HasHref() {
						return fmt.Errorf("create actor link missing href")
					}
					iri = vLink.GetHref()
				} else if vIRI, ok := v.(*url.URL); ok {
					iri = vIRI
				}
				obj[i].AppendAttributedToIRI(iri)
			}
		}
	}
	for _, attributedToMap := range objectAttributedToIds {
		for k, v := range attributedToMap {
			if _, ok := createActorIds[k]; !ok {
				var iri *url.URL
				if vObj, ok := v.(vocab.ObjectType); ok {
					if !vObj.HasId() {
						return fmt.Errorf("attributedTo object missing id")
					}
					iri = vObj.GetId()
				} else if vLink, ok := v.(vocab.LinkType); ok {
					if !vLink.HasHref() {
						return fmt.Errorf("attributedTo link missing href")
					}
					iri = vLink.GetHref()
				} else if vIRI, ok := v.(*url.URL); ok {
					iri = vIRI
				}
				c.AppendActorIRI(iri)
			}
		}
	}
	// As such, a server SHOULD copy any recipients of the Create activity to its
	// object upon initial distribution, and likewise with copying recipients from
	// the object to the wrapping Create activity.
	if err := f.sameRecipients(c); err != nil {
		return err
	}
	// Create requires the client application to persist the 'object' that
	// was created.
	for _, o := range obj {
		if err := f.App.Set(ctx, o); err != nil {
			return err
		}
	}
	// Persist the above changes in the outbox
	var err error
	*toAddToOutbox = make(map[string]interface{})
	*toAddToOutbox, err = c.Serialize()
	if err != nil {
		return err
	}
	return f.ClientCallbacker.Create(ctx, s)
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

package pub

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-fed/activity/streams"
	"github.com/go-fed/activity/vocab"
	"github.com/go-fed/httpsig"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	// errObjectRequired means the activity needs its object property set.
	errObjectRequired = errors.New("object property required")
	// errTargetRequired means the activity needs its target property set.
	errTargetRequired = errors.New("target property required")
)

// TODO: Helper for sending arbitrary ActivityPub objects.

// Pubber provides methods for interacting with ActivityPub clients and
// ActivityPub federating servers.
type Pubber interface {
	// PostInbox returns true if the request was handled as an ActivityPub
	// POST to an actor's inbox. If false, the request was not an
	// ActivityPub request.
	//
	// If the error is nil, then the ResponseWriter's headers and response
	// has already been written. If a non-nil error is returned, then no
	// response has been written.
	PostInbox(c context.Context, w http.ResponseWriter, r *http.Request) (bool, error)
	// GetInbox returns true if the request was handled as an ActivityPub
	// GET to an actor's inbox. If false, the request was not an ActivityPub
	// request.
	//
	// If the error is nil, then the ResponseWriter's headers and response
	// has already been written. If a non-nil error is returned, then no
	// response has been written.
	GetInbox(c context.Context, w http.ResponseWriter, r *http.Request) (bool, error)
	// PostOutbox returns true if the request was handled as an ActivityPub
	// POST to an actor's outbox. If false, the request was not an
	// ActivityPub request.
	//
	// If the error is nil, then the ResponseWriter's headers and response
	// has already been written. If a non-nil error is returned, then no
	// response has been written.
	PostOutbox(c context.Context, w http.ResponseWriter, r *http.Request) (bool, error)
	// GetOutbox returns true if the request was handled as an ActivityPub
	// GET to an actor's outbox. If false, the request was not an
	// ActivityPub request.
	//
	// If the error is nil, then the ResponseWriter's headers and response
	// has already been written. If a non-nil error is returned, then no
	// response has been written.
	GetOutbox(c context.Context, w http.ResponseWriter, r *http.Request) (bool, error)
}

// NewSocialPubber provides a Pubber that implements only the Social API in
// ActivityPub.
func NewSocialPubber(clock Clock, app SocialApplication, cb Callbacker) Pubber {
	return &federator{
		Clock:            clock,
		App:              app,
		SocialAPI:        app,
		ClientCallbacker: cb,
		EnableClient:     true,
	}
}

// NewFederatingPubber provides a Pubber that implements only the Federating API
// in ActivityPub.
func NewFederatingPubber(clock Clock, app FederateApplication, cb Callbacker, d Deliverer, client HttpClient, userAgent string, maxDeliveryDepth, maxForwardingDepth int) Pubber {
	return &federator{
		Clock:                   clock,
		App:                     app,
		FederateAPI:             app,
		ServerCallbacker:        cb,
		Client:                  client,
		Agent:                   userAgent,
		MaxDeliveryDepth:        maxDeliveryDepth,
		MaxInboxForwardingDepth: maxForwardingDepth,
		EnableServer:            true,
		deliverer:               d,
	}
}

// NewPubber provides a Pubber that implements both the Social API and the
// Federating API in ActivityPub.
func NewPubber(clock Clock, app SocialFederateApplication, client, server Callbacker, d Deliverer, httpClient HttpClient, userAgent string, maxDeliveryDepth, maxForwardingDepth int) Pubber {
	return &federator{
		Clock:                   clock,
		App:                     app,
		SocialAPI:               app,
		FederateAPI:             app,
		Client:                  httpClient,
		Agent:                   userAgent,
		MaxDeliveryDepth:        maxDeliveryDepth,
		MaxInboxForwardingDepth: maxForwardingDepth,
		ServerCallbacker:        server,
		ClientCallbacker:        client,
		EnableClient:            true,
		EnableServer:            true,
		deliverer:               d,
	}
}

type federator struct {
	// EnableClient enables or disables the Social API, the client to
	// server part of ActivityPub. Useful if permitting remote clients to
	// act on behalf of the users of the client application.
	EnableClient bool
	// EnableServer enables or disables the Federated Protocol, or the
	// server to server part of ActivityPub. Useful to permit integrating
	// with the rest of the federative web.
	EnableServer bool
	// Clock determines the time of this federator.
	Clock Clock
	// App is the client application that is ActivityPub aware.
	//
	// It is always required.
	App Application
	// FederateAPI provides utility when handling incoming messages received
	// via the Federated Protocol, or server-to-server communications.
	//
	// It is only required if EnableServer is true.
	FederateAPI FederateAPI
	// SocialAPI provides utility when handling incoming messages
	// received via the Social API, or client-to-server communications.
	//
	// It is only required if EnableClient is true.
	SocialAPI SocialAPI
	// ClientCallbacker for handing Social API messages.
	ClientCallbacker Callbacker
	// ServerCallbacker for handling Federated API messages.
	ServerCallbacker Callbacker
	// Client is used to federate with other ActivityPub servers.
	//
	// It is only required if EnableServer is true.
	Client HttpClient
	// Agent is the User-Agent string to use in HTTP headers when
	// federating with another server. It will automatically be appended
	// with '(go-fed ActivityPub)'.
	//
	// It is only required if EnableServer is true.
	Agent string
	// MaxDeliveryDepth is how deep collections of recipients will be
	// expanded for delivery. It must be at least 1 to be compliant with the
	// ActivityPub spec.
	//
	// It is only required if EnableServer is true.
	MaxDeliveryDepth int
	// MaxInboxForwardingDepth is how deep the values are examined for
	// determining ownership of whether to forward an Activity to
	// collections or followers. Once this maximum is exceeded, the ghost
	// replies issue may become a problem, but users may not mind.
	//
	// It is only required if EnableServer is true.
	MaxInboxForwardingDepth int
	// deliverer handles deliveries to other federated servers.
	//
	// It is only required if EnableServer is true.
	deliverer Deliverer
}

func (f *federator) PostInbox(c context.Context, w http.ResponseWriter, r *http.Request) (bool, error) {
	if !isActivityPubPost(r) {
		return false, nil
	}
	if !f.EnableServer {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return true, nil
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return true, err
	}
	var m map[string]interface{}
	if err = json.Unmarshal(b, &m); err != nil {
		return true, err
	}
	ao, err := getActorObject(m)
	if err != nil {
		return true, err
	}
	var iris []*url.URL
	for i := 0; i < ao.ActorLen(); i++ {
		if ao.IsActorObject(i) {
			obj := ao.GetActorObject(i)
			if obj.HasId() {
				iris = append(iris, obj.GetId())
			}
		} else if ao.IsActorLink(i) {
			l := ao.GetActorLink(i)
			if l.HasHref() {
				iris = append(iris, l.GetHref())
			}
		} else if ao.IsActorIRI(i) {
			iris = append(iris, ao.GetActorIRI(i))
		}
	}
	if err = f.FederateAPI.Unblocked(c, iris); err != nil {
		return true, err
	}
	if err = f.getPostInboxResolver(c, r.URL).Deserialize(m); err != nil {
		if err == errObjectRequired || err == errTargetRequired {
			w.WriteHeader(http.StatusBadRequest)
			return true, nil
		}
		return true, err
	}
	if err := f.addToInbox(c, r, m); err != nil {
		return true, err
	}
	if err := f.inboxForwarding(c, m); err != nil {
		return true, err
	}
	w.WriteHeader(http.StatusOK)
	return true, nil
}

func (f *federator) GetInbox(c context.Context, w http.ResponseWriter, r *http.Request) (bool, error) {
	if !isActivityPubGet(r) {
		return false, nil
	}
	oc, err := f.App.GetInbox(c, r, Read)
	if err != nil {
		return true, err
	}
	oc, err = f.dedupeOrderedItems(oc)
	if err != nil {
		return true, err
	}
	m, err := oc.Serialize()
	if err != nil {
		return true, err
	}
	addJSONLDContext(m)
	b, err := json.Marshal(m)
	if err != nil {
		return true, err
	}
	addResponseHeaders(w.Header(), f.Clock, b)
	w.WriteHeader(http.StatusOK)
	n, err := w.Write(b)
	if err != nil {
		return true, err
	} else if n != len(b) {
		return true, fmt.Errorf("ResponseWriter.Write wrote %d of %d bytes", n, len(b))
	}
	return true, nil
}

func (f *federator) PostOutbox(c context.Context, w http.ResponseWriter, r *http.Request) (bool, error) {
	if !isActivityPubPost(r) {
		return false, nil
	}
	if !f.EnableClient {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return true, nil
	}
	// By default, enforce HTTP Signatures.
	authenticated := false
	authorized := true
	if verifier := f.SocialAPI.GetSocialAPIVerifier(c); verifier != nil {
		// Use custom Social API method to authenticate and authorize.
		authenticated, authorized, err := verifier.VerifyForOutbox(r, r.URL)
		if err != nil {
			return true, err
		} else if authenticated && !authorized {
			w.WriteHeader(http.StatusForbidden)
			return true, nil
		} else if !authenticated && !authorized {
			w.WriteHeader(http.StatusBadRequest)
			return true, nil
		}
	}
	if !authenticated && authorized {
		// Use HTTP Signatures to authenticate and authorize.
		v, err := httpsig.NewVerifier(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return true, nil
		}
		pk, algo, err := f.SocialAPI.GetPublicKeyForOutbox(c, v.KeyId(), r.URL)
		if err != nil {
			return true, err
		}
		err = v.Verify(pk, algo)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return true, nil
		}
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return true, err
	}
	var m map[string]interface{}
	if err = json.Unmarshal(b, &m); err != nil {
		return true, err
	}
	typer, err := toTypeIder(m)
	if err != nil {
		return true, err
	}
	if !isActivityType(typer) {
		actorIri, err := f.SocialAPI.ActorIRI(c, r)
		if err != nil {
			return true, err
		}
		obj, ok := typer.(vocab.ObjectType)
		if !ok {
			return true, fmt.Errorf("wrap in create: cannot convert to vocab.ObjectType: %T", typer)
		}
		typer = f.wrapInCreate(obj, actorIri)
	}
	newId := f.App.NewId(c, typer)
	typer.SetId(newId)
	if m, err = typer.Serialize(); err != nil {
		return true, err
	}
	deliverable := false
	if err = f.getPostOutboxResolver(c, m, &deliverable, &m).Deserialize(m); err != nil {
		if err == errObjectRequired || err == errTargetRequired {
			w.WriteHeader(http.StatusBadRequest)
			return true, nil
		}
		return true, err
	}
	if err := f.addToOutbox(c, r, m); err != nil {
		return true, err
	}
	if f.EnableServer && deliverable {
		obj, err := toAnyActivity(m)
		if err != nil {
			return true, err
		}
		if err := f.deliver(obj, r.URL); err != nil {
			return true, err
		}
	}
	w.Header().Set("Location", newId.String())
	w.WriteHeader(http.StatusCreated)
	return true, nil
}

func (f *federator) GetOutbox(c context.Context, w http.ResponseWriter, r *http.Request) (bool, error) {
	if !isActivityPubGet(r) {
		return false, nil
	}
	oc, err := f.App.GetOutbox(c, r, Read)
	if err != nil {
		return true, err
	}
	m, err := oc.Serialize()
	if err != nil {
		return true, err
	}
	addJSONLDContext(m)
	b, err := json.Marshal(m)
	if err != nil {
		return true, err
	}
	addResponseHeaders(w.Header(), f.Clock, b)
	w.WriteHeader(http.StatusOK)
	n, err := w.Write(b)
	if err != nil {
		return true, err
	} else if n != len(b) {
		return true, fmt.Errorf("ResponseWriter.Write wrote %d of %d bytes", n, len(b))
	}
	return true, nil
}

func (f *federator) getPostOutboxResolver(c context.Context, rawJson map[string]interface{}, deliverable *bool, toAddToOutbox *map[string]interface{}) *streams.Resolver {
	return &streams.Resolver{
		CreateCallback: f.handleClientCreate(c, deliverable, toAddToOutbox),
		UpdateCallback: f.handleClientUpdate(c, rawJson, deliverable),
		DeleteCallback: f.handleClientDelete(c, deliverable),
		FollowCallback: f.handleClientFollow(c, deliverable),
		AcceptCallback: f.handleClientAccept(c, deliverable),
		RejectCallback: f.handleClientReject(c, deliverable),
		AddCallback:    f.handleClientAdd(c, deliverable),
		RemoveCallback: f.handleClientRemove(c, deliverable),
		LikeCallback:   f.handleClientLike(c, deliverable),
		UndoCallback:   f.handleClientUndo(c, deliverable),
		BlockCallback:  f.handleClientBlock(c, deliverable),
		// TODO: Extended activity types, such as Announce, Arrive, etc.
	}
}

func (f *federator) handleClientCreate(ctx context.Context, deliverable *bool, toAddToOutbox *map[string]interface{}) func(s *streams.Create) error {
	return func(s *streams.Create) error {
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
					if vObj, ok := v.(vocab.ObjectType); ok {
						obj[i].AppendAttributedToObject(vObj)
					} else if vLink, ok := v.(vocab.LinkType); ok {
						obj[i].AppendAttributedToLink(vLink)
					} else if vIRI, ok := v.(*url.URL); ok {
						obj[i].AppendAttributedToIRI(vIRI)
					}
				}
			}
		}
		for _, attributedToMap := range objectAttributedToIds {
			for k, v := range attributedToMap {
				if _, ok := createActorIds[k]; !ok {
					if vObj, ok := v.(vocab.ObjectType); ok {
						c.AppendActorObject(vObj)
					} else if vLink, ok := v.(vocab.LinkType); ok {
						c.AppendActorLink(vLink)
					} else if vIRI, ok := v.(*url.URL); ok {
						c.AppendActorIRI(vIRI)
					}
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
}

func (f *federator) handleClientUpdate(c context.Context, rawJson map[string]interface{}, deliverable *bool) func(s *streams.Update) error {
	return func(s *streams.Update) error {
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
}

func (f *federator) handleClientDelete(c context.Context, deliverable *bool) func(s *streams.Delete) error {
	return func(s *streams.Delete) error {
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
}

func (f *federator) handleClientFollow(c context.Context, deliverable *bool) func(s *streams.Follow) error {
	return func(s *streams.Follow) error {
		*deliverable = true
		if s.LenObject() == 0 {
			return errObjectRequired
		}
		return f.ClientCallbacker.Follow(c, s)
	}
}

func (f *federator) handleClientAccept(c context.Context, deliverable *bool) func(s *streams.Accept) error {
	return func(s *streams.Accept) error {
		*deliverable = true
		return f.ClientCallbacker.Accept(c, s)
	}
}

func (f *federator) handleClientReject(c context.Context, deliverable *bool) func(s *streams.Reject) error {
	return func(s *streams.Reject) error {
		*deliverable = true
		return f.ClientCallbacker.Reject(c, s)
	}
}

func (f *federator) handleClientAdd(c context.Context, deliverable *bool) func(s *streams.Add) error {
	return func(s *streams.Add) error {
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
			if !raw.IsObject(i) {
				// TODO: Fetch IRIs as well
				return fmt.Errorf("add object must be object type: %v", raw)
			}
			obj := raw.GetObject(i)
			for _, target := range targets {
				if !f.App.CanAdd(c, obj, target) {
					continue
				}
				if ct, ok := target.(vocab.CollectionType); ok {
					ct.AppendItemsObject(obj)
				} else if oct, ok := target.(vocab.OrderedCollectionType); ok {
					oct.AppendOrderedItemsObject(obj)
				}
				if err := f.App.Set(c, target); err != nil {
					return err
				}
			}
		}
		return f.ClientCallbacker.Add(c, s)
	}
}

func (f *federator) handleClientRemove(c context.Context, deliverable *bool) func(s *streams.Remove) error {
	return func(s *streams.Remove) error {
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
}

func (f *federator) handleClientLike(ctx context.Context, deliverable *bool) func(s *streams.Like) error {
	return func(s *streams.Like) error {
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
			return false, fmt.Errorf("cannot determine type of actor liked")
		}
		if err := f.addAllObjectsToActorCollection(ctx, getter, s.Raw(), true); err != nil {
			return err
		}
		return f.ClientCallbacker.Like(ctx, s)
	}
}

func (f *federator) handleClientUndo(c context.Context, deliverable *bool) func(s *streams.Undo) error {
	return func(s *streams.Undo) error {
		*deliverable = true
		if s.LenObject() == 0 {
			return errObjectRequired
		}
		// TODO: Determine if we can support common forms of undo natively.
		return f.ClientCallbacker.Undo(c, s)
	}
}

func (f *federator) handleClientBlock(c context.Context, deliverable *bool) func(s *streams.Block) error {
	return func(s *streams.Block) error {
		*deliverable = false
		if s.LenObject() == 0 {
			return errObjectRequired
		}
		return f.ClientCallbacker.Block(c, s)
	}
}

func (f *federator) getPostInboxResolver(c context.Context, inboxURL *url.URL) *streams.Resolver {
	return &streams.Resolver{
		CreateCallback: f.handleCreate(c),
		UpdateCallback: f.handleUpdate(c),
		DeleteCallback: f.handleDelete(c),
		FollowCallback: f.handleFollow(c, inboxURL),
		AcceptCallback: f.handleAccept(c),
		RejectCallback: f.handleReject(c),
		AddCallback:    f.handleAdd(c),
		RemoveCallback: f.handleRemove(c),
		LikeCallback:   f.handleLike(c),
		UndoCallback:   f.handleUndo(c),
		BlockCallback:  f.handleBlock(c),
		// TODO: Extended activity types, such as Announce, Arrive, etc.
	}
}

func (f *federator) handleCreate(c context.Context) func(s *streams.Create) error {
	return func(s *streams.Create) error {
		// Create requires the client application to persist the 'object' that
		// was created.
		if s.LenObject() == 0 {
			return errObjectRequired
		}
		raw := s.Raw()
		for i := 0; i < raw.ObjectLen(); i++ {
			if !raw.IsObject(i) {
				// TODO: Fetch IRIs as well
				return fmt.Errorf("create requires object to be wholly provided at index %d", i)
			}
			obj := raw.GetObject(i)
			if err := f.App.Set(c, obj); err != nil {
				return err
			}
		}
		return f.ServerCallbacker.Create(c, s)
	}
}

func (f *federator) handleUpdate(c context.Context) func(s *streams.Update) error {
	return func(s *streams.Update) error {
		if s.LenObject() == 0 {
			return errObjectRequired
		}
		raw := s.Raw()
		if err := f.ensureActivityOriginMatchesObjects(raw); err != nil {
			return err
		}
		for i := 0; i < raw.ObjectLen(); i++ {
			if !raw.IsObject(i) {
				// TODO: Fetch IRIs as well
				return fmt.Errorf("update requires object to be wholly provided at index %d", i)
			}
			obj := raw.GetObject(i)
			if err := f.App.Set(c, obj); err != nil {
				return err
			}
		}
		return f.ServerCallbacker.Update(c, s)
	}
}

func (f *federator) handleDelete(c context.Context) func(s *streams.Delete) error {
	return func(s *streams.Delete) error {
		if s.LenObject() == 0 {
			return errObjectRequired
		}
		raw := s.Raw()
		if err := f.ensureActivityOriginMatchesObjects(raw); err != nil {
			return err
		}
		ids, err := getObjectIds(raw)
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
		return f.ServerCallbacker.Delete(c, s)
	}
}

func (f *federator) handleFollow(c context.Context, inboxURL *url.URL) func(s *streams.Follow) error {
	return func(s *streams.Follow) error {
		// Permit either human-triggered or automatically triggering
		// 'Accept'/'Reject'.
		if s.LenObject() == 0 {
			return errObjectRequired
		}
		todo := f.FederateAPI.OnFollow(c, s)
		if todo != DoNothing {
			var activity vocab.ActivityType
			if todo == AutomaticAccept {
				activity = &vocab.Accept{}
			} else if todo == AutomaticReject {
				activity = &vocab.Reject{}
			}
			raw := s.Raw()
			activity.AppendObject(raw)
			for i := 0; i < raw.ActorLen(); i++ {
				if raw.IsActorObject(i) {
					activity.AppendToObject(raw.GetActorObject(i))
				} else if raw.IsActorLink(i) {
					activity.AppendToLink(raw.GetActorLink(i))
				} else if raw.IsActorIRI(i) {
					activity.AppendToIRI(raw.GetActorIRI(i))
				}
			}
			ownsAny := false
			if todo == AutomaticAccept {
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
					return false, fmt.Errorf("cannot determine type of object followers")
				}
				var err error
				if ownsAny, err = f.addAllActorsToObjectCollection(c, getter, raw, true); err != nil {
					return err
				}
			} else if todo == AutomaticReject {
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
}

func (f *federator) handleAccept(c context.Context) func(s *streams.Accept) error {
	return func(s *streams.Accept) error {
		// Accept can be client application specific. However, if this 'Accept'
		// is in response to a 'Follow' then the 'actor' should be added to the
		// original 'actor's 'following' collection by the client application.
		raw := s.Raw()
		for i := 0; i < raw.ObjectLen(); i++ {
			if raw.IsObject(i) {
				obj := raw.GetObject(i)
				follow, ok := obj.(vocab.FollowType)
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
					return false, fmt.Errorf("cannot determine type of actor following")
				}
				if err := f.addAllObjectsToActorCollection(c, getter, follow, true); err != nil {
					return err
				}
			}
		}
		return f.ServerCallbacker.Accept(c, s)
	}
}

func (f *federator) handleReject(c context.Context) func(s *streams.Reject) error {
	return func(s *streams.Reject) error {
		// Reject can be client application specific. However, if this 'Reject'
		// is in response to a 'Follow' then the client MUST NOT go forward with
		// adding the 'actor' to the original 'actor's 'following' collection
		// by the client application.
		return f.ServerCallbacker.Reject(c, s)
	}
}

func (f *federator) handleAdd(c context.Context) func(s *streams.Add) error {
	return func(s *streams.Add) error {
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
			if !raw.IsObject(i) {
				// TODO: Fetch IRIs as well
				return fmt.Errorf("add object must be object type: %v", raw)
			}
			obj := raw.GetObject(i)
			for _, target := range targets {
				if !f.App.CanAdd(c, obj, target) {
					continue
				}
				if ct, ok := target.(vocab.CollectionType); ok {
					ct.AppendItemsObject(obj)
				} else if oct, ok := target.(vocab.OrderedCollectionType); ok {
					oct.AppendOrderedItemsObject(obj)
				}
				if err := f.App.Set(c, target); err != nil {
					return err
				}
			}
		}
		return f.ServerCallbacker.Add(c, s)
	}
}

func (f *federator) handleRemove(c context.Context) func(s *streams.Remove) error {
	return func(s *streams.Remove) error {
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
}

func (f *federator) handleLike(c context.Context) func(s *streams.Like) error {
	return func(s *streams.Like) error {
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
			return false, fmt.Errorf("cannot determine type of object likes")
		}
		if _, err := f.addAllActorsToObjectCollection(c, getter, s.Raw(), true); err != nil {
			return err
		}
		return f.ServerCallbacker.Like(c, s)
	}
}

func (f *federator) handleUndo(c context.Context) func(s *streams.Undo) error {
	return func(s *streams.Undo) error {
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
}

func (f *federator) handleBlock(c context.Context) func(s *streams.Block) error {
	// Servers SHOULD NOT deliver Block Activities to their object. So in
	// this case we will explicitly ignore it, but validate it as if we
	// were to accept it.
	return func(s *streams.Block) error {
		if s.LenObject() == 0 {
			return errObjectRequired
		}
		return nil
	}
}

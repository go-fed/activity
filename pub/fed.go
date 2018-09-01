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
	if err := f.addToInboxIfNew(c, r, m, func() error {
		if err = f.getPostInboxResolver(c, r.URL).Deserialize(m); err != nil {
			return err
		}
		return nil
	}); err != nil {
		if err == errObjectRequired || err == errTargetRequired {
			w.WriteHeader(http.StatusBadRequest)
			return true, nil
		} else {
			return true, err
		}
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
		var err error
		// Use custom Social API method to authenticate and authorize.
		authenticated, authorized, err = verifier.VerifyForOutbox(r, r.URL)
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
	if !vocab.IsActivityType(typer) {
		actorIri, err := f.SocialAPI.ActorIRI(c, r)
		if err != nil {
			return true, err
		}
		obj, ok := typer.(vocab.ObjectType)
		if !ok {
			return true, fmt.Errorf("wrap in create: cannot convert to vocab.ObjectType: %T", typer)
		}
		typer, err = f.wrapInCreate(obj, actorIri)
		if err != nil {
			return true, err
		}
	}
	var activity vocab.ActivityType
	var iActivity vocab.IntransitiveActivityType
	var ok bool
	activity, ok = typer.(vocab.ActivityType)
	if !ok {
		iActivity, ok = typer.(vocab.IntransitiveActivityType)
		if !ok {
			return true, fmt.Errorf("assigning new ids: cannot convert to vocab.ActivityType nor vocab.IntransitiveActivityType: %T", typer)
		} else {
			f.addNewIdsIntransitive(c, iActivity)
		}
	} else {
		f.addNewIds(c, activity)
	}
	if m, err = typer.Serialize(); err != nil {
		return true, err
	}
	deliverable := false
	if err = f.getPostOutboxResolver(c, m, &deliverable, &m, r.URL).Deserialize(m); err != nil {
		if err == errObjectRequired || err == errTargetRequired {
			w.WriteHeader(http.StatusBadRequest)
			return true, nil
		}
		return true, err
	}
	if err = f.addToOutbox(c, r, m); err != nil {
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
	w.Header().Set("Location", activity.GetId().String())
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

func (f *federator) getPostOutboxResolver(c context.Context, rawJson map[string]interface{}, deliverable *bool, toAddToOutbox *map[string]interface{}, outboxURL *url.URL) *streams.Resolver {
	return &streams.Resolver{
		CreateCallback: f.handleClientCreate(c, deliverable, toAddToOutbox),
		UpdateCallback: f.handleClientUpdate(c, rawJson, deliverable),
		DeleteCallback: f.handleClientDelete(c, deliverable),
		FollowCallback: f.handleClientFollow(c, deliverable),
		AcceptCallback: f.handleClientAccept(c, deliverable),
		RejectCallback: f.handleClientReject(c, deliverable),
		AddCallback:    f.handleClientAdd(c, deliverable, outboxURL),
		RemoveCallback: f.handleClientRemove(c, deliverable),
		LikeCallback:   f.handleClientLike(c, deliverable),
		UndoCallback:   f.handleClientUndo(c, deliverable),
		BlockCallback:  f.handleClientBlock(c, deliverable),
		// Other activities whose behaviors are not examined by the pub
		// package, and are passed through to extensions of the
		// Callbacker interface.
		AnnounceCallback:        f.handleClientAnnounce(c, deliverable),
		ArriveCallback:          f.handleClientArrive(c, deliverable),
		DislikeCallback:         f.handleClientDislike(c, deliverable),
		FlagCallback:            f.handleClientFlag(c, deliverable),
		IgnoreCallback:          f.handleClientIgnore(c, deliverable),
		InviteCallback:          f.handleClientInvite(c, deliverable),
		JoinCallback:            f.handleClientJoin(c, deliverable),
		LeaveCallback:           f.handleClientLeave(c, deliverable),
		ListenCallback:          f.handleClientListen(c, deliverable),
		MoveCallback:            f.handleClientMove(c, deliverable),
		OfferCallback:           f.handleClientOffer(c, deliverable),
		QuestionCallback:        f.handleClientQuestion(c, deliverable),
		ReadCallback:            f.handleClientRead(c, deliverable),
		TentativeAcceptCallback: f.handleClientTentativeAccept(c, deliverable),
		TentativeRejectCallback: f.handleClientTentativeReject(c, deliverable),
		TravelCallback:          f.handleClientTravel(c, deliverable),
		ViewCallback:            f.handleClientView(c, deliverable),
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

func (f *federator) handleClientAdd(c context.Context, deliverable *bool, outboxURL *url.URL) func(s *streams.Add) error {
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
			*loc = &vocab.OrderedCollection{}
			actor.SetLikedOrderedCollection(*loc)
			return false, nil
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
		raw := s.Raw()
		if err := f.ensureActivityActorsMatchObjectActors(raw); err != nil {
			return err
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

func (f *federator) handleClientAnnounce(c context.Context, deliverable *bool) func(s *streams.Announce) error {
	return func(s *streams.Announce) error {
		*deliverable = true
		if t, ok := f.ClientCallbacker.(callbackerAnnounce); ok {
			return t.Announce(c, s)
		}
		return nil
	}
}

func (f *federator) handleClientArrive(c context.Context, deliverable *bool) func(s *streams.Arrive) error {
	return func(s *streams.Arrive) error {
		*deliverable = true
		if t, ok := f.ClientCallbacker.(callbackerArrive); ok {
			return t.Arrive(c, s)
		}
		return nil
	}
}

func (f *federator) handleClientDislike(c context.Context, deliverable *bool) func(s *streams.Dislike) error {
	return func(s *streams.Dislike) error {
		*deliverable = true
		if t, ok := f.ClientCallbacker.(callbackerDislike); ok {
			return t.Dislike(c, s)
		}
		return nil
	}
}

func (f *federator) handleClientFlag(c context.Context, deliverable *bool) func(s *streams.Flag) error {
	return func(s *streams.Flag) error {
		*deliverable = true
		if t, ok := f.ClientCallbacker.(callbackerFlag); ok {
			return t.Flag(c, s)
		}
		return nil
	}
}

func (f *federator) handleClientIgnore(c context.Context, deliverable *bool) func(s *streams.Ignore) error {
	return func(s *streams.Ignore) error {
		*deliverable = true
		if t, ok := f.ClientCallbacker.(callbackerIgnore); ok {
			return t.Ignore(c, s)
		}
		return nil
	}
}

func (f *federator) handleClientInvite(c context.Context, deliverable *bool) func(s *streams.Invite) error {
	return func(s *streams.Invite) error {
		*deliverable = true
		if t, ok := f.ClientCallbacker.(callbackerInvite); ok {
			return t.Invite(c, s)
		}
		return nil
	}
}

func (f *federator) handleClientJoin(c context.Context, deliverable *bool) func(s *streams.Join) error {
	return func(s *streams.Join) error {
		*deliverable = true
		if t, ok := f.ClientCallbacker.(callbackerJoin); ok {
			return t.Join(c, s)
		}
		return nil
	}
}

func (f *federator) handleClientLeave(c context.Context, deliverable *bool) func(s *streams.Leave) error {
	return func(s *streams.Leave) error {
		*deliverable = true
		if t, ok := f.ClientCallbacker.(callbackerLeave); ok {
			return t.Leave(c, s)
		}
		return nil
	}
}

func (f *federator) handleClientListen(c context.Context, deliverable *bool) func(s *streams.Listen) error {
	return func(s *streams.Listen) error {
		*deliverable = true
		if t, ok := f.ClientCallbacker.(callbackerListen); ok {
			return t.Listen(c, s)
		}
		return nil
	}
}

func (f *federator) handleClientMove(c context.Context, deliverable *bool) func(s *streams.Move) error {
	return func(s *streams.Move) error {
		*deliverable = true
		if t, ok := f.ClientCallbacker.(callbackerMove); ok {
			return t.Move(c, s)
		}
		return nil
	}
}

func (f *federator) handleClientOffer(c context.Context, deliverable *bool) func(s *streams.Offer) error {
	return func(s *streams.Offer) error {
		*deliverable = true
		if t, ok := f.ClientCallbacker.(callbackerOffer); ok {
			return t.Offer(c, s)
		}
		return nil
	}
}

func (f *federator) handleClientQuestion(c context.Context, deliverable *bool) func(s *streams.Question) error {
	return func(s *streams.Question) error {
		*deliverable = true
		if t, ok := f.ClientCallbacker.(callbackerQuestion); ok {
			return t.Question(c, s)
		}
		return nil
	}
}

func (f *federator) handleClientRead(c context.Context, deliverable *bool) func(s *streams.Read) error {
	return func(s *streams.Read) error {
		*deliverable = true
		if t, ok := f.ClientCallbacker.(callbackerRead); ok {
			return t.Read(c, s)
		}
		return nil
	}
}

func (f *federator) handleClientTentativeAccept(c context.Context, deliverable *bool) func(s *streams.TentativeAccept) error {
	return func(s *streams.TentativeAccept) error {
		*deliverable = true
		if t, ok := f.ClientCallbacker.(callbackerTentativeAccept); ok {
			return t.TentativeAccept(c, s)
		}
		return nil
	}
}

func (f *federator) handleClientTentativeReject(c context.Context, deliverable *bool) func(s *streams.TentativeReject) error {
	return func(s *streams.TentativeReject) error {
		*deliverable = true
		if t, ok := f.ClientCallbacker.(callbackerTentativeReject); ok {
			return t.TentativeReject(c, s)
		}
		return nil
	}
}

func (f *federator) handleClientTravel(c context.Context, deliverable *bool) func(s *streams.Travel) error {
	return func(s *streams.Travel) error {
		*deliverable = true
		if t, ok := f.ClientCallbacker.(callbackerTravel); ok {
			return t.Travel(c, s)
		}
		return nil
	}
}

func (f *federator) handleClientView(c context.Context, deliverable *bool) func(s *streams.View) error {
	return func(s *streams.View) error {
		*deliverable = true
		if t, ok := f.ClientCallbacker.(callbackerView); ok {
			return t.View(c, s)
		}
		return nil
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
		// Other activities whose behaviors are not examined by the pub
		// package (except for Announce), and are passed through to
		// extensions of the Callbacker interface.
		AnnounceCallback:        f.handleAnnounce(c),
		ArriveCallback:          f.handleArrive(c),
		DislikeCallback:         f.handleDislike(c),
		FlagCallback:            f.handleFlag(c),
		IgnoreCallback:          f.handleIgnore(c),
		InviteCallback:          f.handleInvite(c),
		JoinCallback:            f.handleJoin(c),
		LeaveCallback:           f.handleLeave(c),
		ListenCallback:          f.handleListen(c),
		MoveCallback:            f.handleMove(c),
		OfferCallback:           f.handleOffer(c),
		QuestionCallback:        f.handleQuestion(c),
		ReadCallback:            f.handleRead(c),
		TentativeAcceptCallback: f.handleTentativeAccept(c),
		TentativeRejectCallback: f.handleTentativeReject(c),
		TravelCallback:          f.handleTravel(c),
		ViewCallback:            f.handleView(c),
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
				var to *url.URL
				if raw.IsActorObject(i) {
					obj := raw.GetActorObject(i)
					if !obj.HasId() {
						return fmt.Errorf("to object missing id")
					}
					to = obj.GetId()
				} else if raw.IsActorLink(i) {
					href := raw.GetActorLink(i)
					if !href.HasHref() {
						return fmt.Errorf("to link missing href")
					}
					to = href.GetHref()
				} else if raw.IsActorIRI(i) {
					to = raw.GetActorIRI(i)
				}
				activity.AppendToIRI(to)
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
					*loc = &vocab.OrderedCollection{}
					object.SetFollowersOrderedCollection(*loc)
					return false, nil
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
			*loc = &vocab.OrderedCollection{}
			object.SetLikesOrderedCollection(*loc)
			return false, nil
		}
		if _, err := f.addActivityToObjectCollection(c, getter, s.Raw(), true); err != nil {
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

func (f *federator) handleAnnounce(c context.Context) func(s *streams.Announce) error {
	return func(s *streams.Announce) error {
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
}

func (f *federator) handleArrive(c context.Context) func(s *streams.Arrive) error {
	return func(s *streams.Arrive) error {
		if t, ok := f.ServerCallbacker.(callbackerArrive); ok {
			return t.Arrive(c, s)
		}
		return nil
	}
}

func (f *federator) handleDislike(c context.Context) func(s *streams.Dislike) error {
	return func(s *streams.Dislike) error {
		if t, ok := f.ServerCallbacker.(callbackerDislike); ok {
			return t.Dislike(c, s)
		}
		return nil
	}
}

func (f *federator) handleFlag(c context.Context) func(s *streams.Flag) error {
	return func(s *streams.Flag) error {
		if t, ok := f.ServerCallbacker.(callbackerFlag); ok {
			return t.Flag(c, s)
		}
		return nil
	}
}

func (f *federator) handleIgnore(c context.Context) func(s *streams.Ignore) error {
	return func(s *streams.Ignore) error {
		if t, ok := f.ServerCallbacker.(callbackerIgnore); ok {
			return t.Ignore(c, s)
		}
		return nil
	}
}

func (f *federator) handleInvite(c context.Context) func(s *streams.Invite) error {
	return func(s *streams.Invite) error {
		if t, ok := f.ServerCallbacker.(callbackerInvite); ok {
			return t.Invite(c, s)
		}
		return nil
	}
}

func (f *federator) handleJoin(c context.Context) func(s *streams.Join) error {
	return func(s *streams.Join) error {
		if t, ok := f.ServerCallbacker.(callbackerJoin); ok {
			return t.Join(c, s)
		}
		return nil
	}
}

func (f *federator) handleLeave(c context.Context) func(s *streams.Leave) error {
	return func(s *streams.Leave) error {
		if t, ok := f.ServerCallbacker.(callbackerLeave); ok {
			return t.Leave(c, s)
		}
		return nil
	}
}

func (f *federator) handleListen(c context.Context) func(s *streams.Listen) error {
	return func(s *streams.Listen) error {
		if t, ok := f.ServerCallbacker.(callbackerListen); ok {
			return t.Listen(c, s)
		}
		return nil
	}
}

func (f *federator) handleMove(c context.Context) func(s *streams.Move) error {
	return func(s *streams.Move) error {
		if t, ok := f.ServerCallbacker.(callbackerMove); ok {
			return t.Move(c, s)
		}
		return nil
	}
}

func (f *federator) handleOffer(c context.Context) func(s *streams.Offer) error {
	return func(s *streams.Offer) error {
		if t, ok := f.ServerCallbacker.(callbackerOffer); ok {
			return t.Offer(c, s)
		}
		return nil
	}
}

func (f *federator) handleQuestion(c context.Context) func(s *streams.Question) error {
	return func(s *streams.Question) error {
		if t, ok := f.ServerCallbacker.(callbackerQuestion); ok {
			return t.Question(c, s)
		}
		return nil
	}
}

func (f *federator) handleRead(c context.Context) func(s *streams.Read) error {
	return func(s *streams.Read) error {
		if t, ok := f.ServerCallbacker.(callbackerRead); ok {
			return t.Read(c, s)
		}
		return nil
	}
}

func (f *federator) handleTentativeAccept(c context.Context) func(s *streams.TentativeAccept) error {
	return func(s *streams.TentativeAccept) error {
		if t, ok := f.ServerCallbacker.(callbackerTentativeAccept); ok {
			return t.TentativeAccept(c, s)
		}
		return nil
	}
}

func (f *federator) handleTentativeReject(c context.Context) func(s *streams.TentativeReject) error {
	return func(s *streams.TentativeReject) error {
		if t, ok := f.ServerCallbacker.(callbackerTentativeReject); ok {
			return t.TentativeReject(c, s)
		}
		return nil
	}
}

func (f *federator) handleTravel(c context.Context) func(s *streams.Travel) error {
	return func(s *streams.Travel) error {
		if t, ok := f.ServerCallbacker.(callbackerTravel); ok {
			return t.Travel(c, s)
		}
		return nil
	}
}

func (f *federator) handleView(c context.Context) func(s *streams.View) error {
	return func(s *streams.View) error {
		if t, ok := f.ServerCallbacker.(callbackerView); ok {
			return t.View(c, s)
		}
		return nil
	}
}

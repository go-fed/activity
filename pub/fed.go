package pub

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-fed/activity/streams"
	"github.com/go-fed/activity/vocab"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	// ErrObjectRequired means the activity needs its object property set.
	ErrObjectRequired = errors.New("object property required")
	// ErrTypeRequired means the activity needs its type property set.
	ErrTypeRequired = errors.New("type property required")
)

// Object is the interface for ActivityPub compliant ActivityStream types.
//
// ActivityPub obects must have 'id' and 'type' properties to be spec compliant.
// This library enforces the "MUST" requirement and purposefully does NOT
// support transient objects.
//
// Furthermore, the spec extends the Object ActivityStream type to also contain
// the 'source' property, whose value appears to be an Object with only
// 'content' and 'mediaType' properties being semantically useful.
//
// TODO: Possibly delete this interface
type Object interface {
	GetId() (streams.Resolution, url.URL)
	SetId(url.URL)
	HasId() streams.Presence
	LenType() int
	GetType(int) (streams.Resolution, string)
	AddType(interface{})
	RemoveType(int)
	ResolveSource(*streams.Resolver) (streams.Resolution, error)
	HasSource() streams.Presence
	SetSource(vocab.ObjectType)
	Serialize() (m map[string]interface{}, err error)
}

var _ Object = &streams.Object{}

// Actor is the interface for ActivityPub compliant ActivityStream types.
//
// ActivityPub actors must have a valid unique IRI as the 'id' property in
// addition to the 'type' property. Furthermore, actors must have the 'inbox'
// and 'outbox' properties to be considered actors. There are more suggested
// properties by the spec, which we include here.
//
// TODO: Possibly delete this interface
type Actor interface {
	GetId() (streams.Resolution, url.URL)
	SetId(url.URL)
	LenType() int
	GetType(int) (streams.Resolution, string)
	AddType(interface{})
	RemoveType(int)
	GetInbox() (streams.Resolution, url.URL)
	HasInbox() streams.Presence
	SetInbox(url.URL)
	GetOutbox() (streams.Resolution, url.URL)
	HasOutbox() streams.Presence
	SetOutbox(url.URL)
	GetFollowing() (streams.Resolution, url.URL)
	HasFollowing() streams.Presence
	SetFollowing(url.URL)
	GetFollowers() (streams.Resolution, url.URL)
	HasFollowers() streams.Presence
	SetFollowers(url.URL)
	GetLiked() (streams.Resolution, url.URL)
	HasLiked() streams.Presence
	SetLiked(url.URL)
	LenStreams() int
	GetStreams(int) (streams.Resolution, url.URL)
	AddStreams(vocab.CollectionType)
	RemoveStreams(int)
	GetPreferredUsername() (streams.Resolution, string)
	HasPreferredUsername() streams.Presence
	SetPreferredUsername(string)
	PreferredUsernameLanguages() []string
	GetPreferredUsernameForLanguage(string) string
	SetPreferredUsernameForLanguage(string, string)
	ResolveEndpoints(streams.Resolver) (streams.Resolution, error)
	HasEndpoints() streams.Presence
	SetEndpoints(vocab.ObjectType)
	Serialize() (m map[string]interface{}, err error)
}

var _ Object = &streams.Object{}

// Endpoint is a logical grouping of specific properties within the ActivityPub
// specification.
//
// TODO: Possibly delete this interface
type Endpoint interface {
	GetProxyUrl() (streams.Resolution, url.URL)
	HasProxyUrl() streams.Presence
	SetProxyUrl(url.URL)
	GetOauthAuthorizationEndpoint() (streams.Resolution, url.URL)
	HasOauthAuthorizationEndpoint() streams.Presence
	SetOauthAuthorizationEndpoint(url.URL)
	GetOauthTokenEndpoint() (streams.Resolution, url.URL)
	HasOauthTokenEndpoint() streams.Presence
	SetOauthTokenEndpoint(url.URL)
	GetProvideClientKey() (streams.Resolution, url.URL)
	HasProvideClientKey() streams.Presence
	SetProvideClientKey(url.URL)
	GetSignClientKey() (streams.Resolution, url.URL)
	HasSignClientKey() streams.Presence
	SetSignClientKey(url.URL)
	GetSharedInbox() (streams.Resolution, url.URL)
	HasSharedInbox() streams.Presence
	SetSharedInbox(url.URL)
}

var _ Endpoint = &streams.Object{}

// TODO: Helper http Handler for serving ActivityStream objects
// TODO: Helper http Handler for serving Tombstone objects
// TODO: Helper http Handler for serving deleted objects

// Typer is an object that has a type.
type Typer interface {
	TypeLen() (l int)
	GetType(index int) (v interface{})
}

type typeIder interface {
	Typer
	SetId(v url.URL)
	Serialize() (m map[string]interface{}, e error)
}

// Application is provided by users of this library in order to
type Application interface {
	// GetInbox returns the OrderedCollection inbox of the actor with the
	// provided ID. It is up to the implementation to provide the correct
	// collection for the kind of authorization given in the request.
	GetInbox(id string, r *http.Request) (vocab.OrderedCollectionType, error)
	// GetOutbox returns the OrderedCollection inbox of the actor with the
	// provided ID. It is up to the implementation to provide the correct
	// collection for the kind of authorization given in the request.
	GetOutbox(id string, r *http.Request) (vocab.OrderedCollectionType, error)
	// PostOutboxAuthorized determines whether the request is able to post
	// an Activity to the outbox for the specified id.
	PostOutboxAuthorized(id string, r *http.Request) (bool, error)
	// NewId takes in a client id token and returns an ActivityStreams IRI
	// id for a new Activity posted to the outbox. The object is provided
	// as a Typer so clients can use it to decide how to generate the IRI.
	NewId(id string, t Typer) url.URL
	// AddToOutboxResolver returns the client's Resolver which must store
	// the provided object in the outbox of the user represented by the
	// client id token.
	AddToOutboxResolver(id string) (*streams.Resolver, error)
	// ActorIRI returns the actor's IRI associated with the given client ID
	// token.
	ActorIRI(id string) (url.URL, error)
}

// Receiver is provided by users of this library and designed to handle
// receiving federated messages through the Federated Protocol.
//
// Note that although Receiver and ClientReceiver have similar methods, many of
// them have different requirements and thus implementations are not
// interchangeable between Receiver and ClientReceiver.
type Receiver interface {
	// Create requires the client application to persist the 'object' that
	// was created.
	Create(id string, s *streams.Create) error
	// Update should completely replace the 'object' with the same 'id'.
	Update(id string, s *streams.Update) error
	// Delete SHOULD completely remove the 'object' with its 'id', or have
	// the 'object' be replaced by a 'Tombstone' ActivityStream type.
	Delete(id string, s *streams.Delete) error
	// Follow means the client application SHOULD reply with an 'Accept' or
	// 'Reject' ActivityStream with the 'Follow' as the 'object' and deliver
	// it to the 'actor' of the 'Follow'. This can be human-triggered or
	// automatically triggered.
	Follow(id string, s *streams.Follow) error
	// Accept can be client application specific. However, if this 'Accept'
	// is in response to a 'Follow' then the 'actor' should be added to the
	// original 'actor's 'following' collection by the client application.
	Accept(id string, s *streams.Accept) error
	// Reject can be client application specific. However, if this 'Reject'
	// is in response to a 'Follow' then the client MUST NOT go forward with
	// adding the 'actor' to the original 'actor's 'following' collection
	// by the client application.
	Reject(id string, s *streams.Reject) error
	// Add is client application specific, generally involving adding an
	// 'object' to a specific 'target' collection.
	Add(id string, s *streams.Add) error
	// Remove is client application specific, generally involving removing
	// an 'object' from a specific 'target' collection.
	Remove(id string, s *streams.Remove) error
	// Like triggers adding the like to an object's `like` collection.
	Like(id string, s *streams.Like) error
	// Undo negates a previous action. The 'actor' on the 'Undo' MUST be the
	// same as the 'actor' on the Activity being undone, and the client
	// application is responsible for enforcing this. Note that 'Undo'-ing
	// is not a deletion of a previous Activity, but the addition of its
	// opposite.
	Undo(id string, s *streams.Undo) error
}

// ClientReceiver is provided by users of this library and designed to handle
// receiving messaged from ActivityPub clients through the Social API.
//
// Note that although ClientReceiver and Receiver have similar methods, many of
// them have different requirements and thus implementations are not
// interchangeable between ClientReceiver and Receiver.
type ClientReceiver interface {
	// Create requires the client application to persist the 'object' that
	// was created.
	Create(id string, s *streams.Create) error
	// ClientUpdate should partially replace the 'object' with only the
	// changed top-level fields.
	//
	// TODO: Support deletions.
	Update(id string, s *streams.Update) error
	// Delete SHOULD completely remove the 'object' with its 'id', or have
	// the 'object' be replaced by a 'Tombstone' ActivityStream type.
	Delete(id string, s *streams.Delete) error
	// Add is client application specific, generally involving adding an
	// 'object' to a specific 'target' collection. The application may at
	// its discretion determine whether this is permissible, by determining
	// if it owns the 'target' collection and/or by other application
	// specific criteria.
	Add(id string, s *streams.Add) error
	// Remove is client application specific, generally involving removing
	// an 'object' from a specific 'target' collection. The application may
	// at its discretion determine whether this is permissible, by
	// determining if it owns the 'target' collection and/or by other
	// application specific criteria.
	Remove(id string, s *streams.Remove) error
	// Like triggers adding the 'object' to the 'actor's `like` collection.
	Like(id string, s *streams.Like) error
	// Block means that the server should not let the 'object' actor
	// interact with the 'actor'.
	Block(id string, s *streams.Block) error
	// Undo negates a previous action. The 'actor' on the 'Undo' MUST be the
	// same as the 'actor' on the Activity being undone, and the client
	// application is responsible for enforcing this. Note that 'Undo'-ing
	// is not a deletion of a previous Activity, but the addition of its
	// opposite.
	Undo(id string, s *streams.Undo) error
	// Accept can be client application specific. However, if this 'Accept'
	// is in response to a 'Follow' then the follower should be added to
	// the 'actor's 'followers' collection.
	Accept(id string, s *streams.Accept) error
	// Reject can be client application specific. However, if this 'Reject'
	// is in response to a 'Follow' then the client MUST NOT go forward with
	// adding the follower to the 'actor's 'followers' collection.
	Reject(id string, s *streams.Reject) error
}

type federator struct {
	// App is the client application that is ActivityPub aware.
	//
	// It is always required.
	App Application
	// Receiver provides callbacks when handling incoming messages received
	// via the Federated Protocol, or server-to-server communications.
	//
	// It is only required if EnableServer is true.
	Receiver Receiver
	// ClientReceiver provides callbacks when handling incoming messages
	// received via the Social API, or client-to-server communications.
	//
	// It is only required if EnableClient is true.
	ClientReceiver ClientReceiver
	// Client is used to federate with other ActivityPub servers.
	//
	// It is only required if EnableServer is true.
	Client *http.Client
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
	// EnableClient enables or disables the Social API, the client to
	// server part of ActivityPub. Useful if permitting remote clients to
	// act on behalf of the users of the client application.
	EnableClient bool
	// EnableServer enables or disables the Federated Protocol, or the
	// server to server part of ActivityPub. Useful to permit integrating
	// with the rest of the federative web.
	EnableServer bool
	// pool properly handles rate-limiting and retrying deliveries to other
	// federated servers.
	//
	// It is only required if EnableServer is true.
	pool *delivererPool
}

func (f *federator) Stop() {
	if f.pool != nil {
		f.pool.Stop()
	}
}

func (f *federator) Errors() <-chan error {
	if f.pool != nil {
		return f.pool.Errors()
	}
	return nil
}

func (f *federator) PostInbox(id string) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) (bool, error) {
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
		if err = f.getPostInboxResolver(id).Deserialize(m); err != nil {
			return true, err
		}
		// TODO: 7.1.2 Inbox forwarding
		w.WriteHeader(http.StatusOK)
		return true, nil
	}
}

func (f *federator) GetInbox(id string) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) (bool, error) {
		if !isActivityPubGet(r) {
			return false, nil
		}
		oc, err := f.App.GetInbox(id, r)
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
		b, err := json.Marshal(m)
		if err != nil {
			return true, err
		}
		w.Header().Set(contentTypeHeader, responseContentTypeHeader)
		w.WriteHeader(http.StatusOK)
		n, err := w.Write(b)
		if err != nil {
			return true, err
		} else if n != len(b) {
			return true, fmt.Errorf("ResponseWriter.Write wrote %d of %d bytes", n, len(b))
		}
		return true, nil
	}
}

// PostOutpox provides a HTTP handler for ActivityPub requests for the given id
// token. The client ID token is passed forwards to other interfaces for
// application specific behavior. The handler will return true if it handled
// the request as an ActivityPub request. If it returns an error, it is up to
// the client to determine how to respond via HTTP.
//
// Note that the error could be ErrObjectRequired or ErrTypeRequired.
func (f *federator) PostOutbox(id string) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) (bool, error) {
		if !isActivityPubPost(r) {
			return false, nil
		}
		if !f.EnableClient {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return true, nil
		}
		ok, err := f.App.PostOutboxAuthorized(id, r)
		if err != nil {
			return true, err
		}
		if !ok {
			w.WriteHeader(http.StatusForbidden)
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
		typer, err := toTypeIder(m)
		if err != nil {
			return true, err
		}
		if !isActivityType(typer) {
			actorIri, err := f.App.ActorIRI(id)
			if err != nil {
				return true, err
			}
			obj, ok := typer.(vocab.ObjectType)
			if !ok {
				return true, fmt.Errorf("wrap in create: cannot convert to vocab.ObjectType: %T", typer)
			}
			typer = f.wrapInCreate(obj, actorIri)
		}
		newId := f.App.NewId(id, typer)
		typer.SetId(newId)
		if m, err = typer.Serialize(); err != nil {
			return true, err
		}
		outboxAdder, err := f.App.AddToOutboxResolver(id)
		if err != nil {
			return true, err
		}
		if err = outboxAdder.Deserialize(m); err != nil {
			return true, err
		}
		if err = f.getPostOutboxResolver(id).Deserialize(m); err != nil {
			return true, err
		}
		if f.EnableServer {
			obj, err := toAnyActivity(m)
			if err != nil {
				return true, err
			}
			recipients, err := f.prepare(obj)
			if err != nil {
				return true, err
			}
			m, err := obj.Serialize()
			if err != nil {
				return true, err
			}
			m["@context"] = "https://www.w3.org/ns/activitystreams"
			b, err := json.Marshal(m)
			if err != nil {
				return true, err
			}
			for _, to := range recipients {
				f.pool.Do(b, to, func(b []byte, u url.URL) error {
					return postToOutbox(f.Client, b, u, f.Agent)
				})
			}
		}
		w.Header().Set("Location", newId.String())
		w.WriteHeader(http.StatusCreated)
		return true, nil
	}
}

func (f *federator) GetOutbox(id string) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) (bool, error) {
		if !isActivityPubGet(r) {
			return false, nil
		}
		oc, err := f.App.GetOutbox(id, r)
		if err != nil {
			return true, err
		}
		m, err := oc.Serialize()
		if err != nil {
			return true, err
		}
		b, err := json.Marshal(m)
		if err != nil {
			return true, err
		}
		w.Header().Set(contentTypeHeader, responseContentTypeHeader)
		w.WriteHeader(http.StatusOK)
		n, err := w.Write(b)
		if err != nil {
			return true, err
		} else if n != len(b) {
			return true, fmt.Errorf("ResponseWriter.Write wrote %d of %d bytes", n, len(b))
		}
		return true, nil
	}
}

func (f *federator) getPostOutboxResolver(id string) *streams.Resolver {
	return &streams.Resolver{
		CreateCallback: f.handleClientCreate(id),
		UpdateCallback: f.handleClientUpdate(id),
		DeleteCallback: f.handleClientDelete(id),
		FollowCallback: f.handleClientFollow(id),
		AcceptCallback: f.handleClientAccept(id),
		RejectCallback: f.handleClientReject(id),
		AddCallback:    f.handleClientAdd(id),
		RemoveCallback: f.handleClientRemove(id),
		LikeCallback:   f.handleClientLike(id),
		UndoCallback:   f.handleClientUndo(id),
		BlockCallback:  f.handleClientBlock(id),
		// TODO: Extended activity types, such as Announce, Arrive, etc.
	}
}

func (f *federator) handleClientCreate(id string) func(s *streams.Create) error {
	return func(s *streams.Create) error {
		if s.LenObject() == 0 {
			return ErrObjectRequired
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
				createActorIds[(&id).String()] = obj
			} else if c.IsActorLink(i) {
				l := c.GetActorLink(i)
				href := l.GetHref()
				createActorIds[(&href).String()] = l
			} else if c.IsActorIRI(i) {
				iri := c.GetActorIRI(i)
				createActorIds[(&iri).String()] = iri
			}
		}
		var obj []vocab.ObjectType
		for i := 0; i < c.ObjectLen(); i++ {
			if c.IsObject(i) {
				obj = append(obj, c.GetObject(i))
			} else if c.IsObjectIRI(i) {
				return fmt.Errorf("unsupported: Create Activity with 'object' that is only an IRI")
			}
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
					objectAttributedToIds[k][(&id).String()] = at
				} else if o.IsAttributedToLink(i) {
					at := o.GetAttributedToLink(i)
					href := at.GetHref()
					objectAttributedToIds[k][(&href).String()] = at
				} else if o.IsAttributedToIRI(i) {
					iri := o.GetAttributedToIRI(i)
					objectAttributedToIds[k][(&iri).String()] = iri
				}
			}
		}
		for k, v := range createActorIds {
			for i, attributedToMap := range objectAttributedToIds {
				if _, ok := attributedToMap[k]; !ok {
					if vObj, ok := v.(vocab.ObjectType); ok {
						obj[i].AddAttributedToObject(vObj)
					} else if vLink, ok := v.(vocab.LinkType); ok {
						obj[i].AddAttributedToLink(vLink)
					} else if vIRI, ok := v.(url.URL); ok {
						obj[i].AddAttributedToIRI(vIRI)
					}
				}
			}
		}
		// As such, a server SHOULD copy any recipients of the Create activity to its
		// object upon initial distribution, and likewise with copying recipients from
		// the object to the wrapping Create activity.
		// Again, presumably if it does not already exist.
		for _, attributedToMap := range objectAttributedToIds {
			for k, v := range attributedToMap {
				if _, ok := createActorIds[k]; !ok {
					if vObj, ok := v.(vocab.ObjectType); ok {
						c.AddActorObject(vObj)
					} else if vLink, ok := v.(vocab.LinkType); ok {
						c.AddActorLink(vLink)
					} else if vIRI, ok := v.(url.URL); ok {
						c.AddActorIRI(vIRI)
					}
				}
			}
		}
		return f.ClientReceiver.Create(id, s)
	}
}

func (f *federator) handleClientUpdate(id string) func(s *streams.Update) error {
	return func(s *streams.Update) error {
		if s.LenObject() == 0 {
			return ErrObjectRequired
		}
		return f.ClientReceiver.Update(id, s)
	}
}

func (f *federator) handleClientDelete(id string) func(s *streams.Delete) error {
	return func(s *streams.Delete) error {
		if s.LenObject() == 0 {
			return ErrObjectRequired
		}
		return f.ClientReceiver.Delete(id, s)
	}
}

func (f *federator) handleClientFollow(id string) func(s *streams.Follow) error {
	return func(s *streams.Follow) error {
		if s.LenObject() == 0 {
			return ErrObjectRequired
		}
		// Nothing extra to do.
		return nil
	}
}

func (f *federator) handleClientAccept(id string) func(s *streams.Accept) error {
	return func(s *streams.Accept) error {
		return f.ClientReceiver.Accept(id, s)
	}
}

func (f *federator) handleClientReject(id string) func(s *streams.Reject) error {
	return func(s *streams.Reject) error {
		return f.ClientReceiver.Reject(id, s)
	}
}

func (f *federator) handleClientAdd(id string) func(s *streams.Add) error {
	return func(s *streams.Add) error {
		if s.LenObject() == 0 {
			return ErrObjectRequired
		} else if s.LenType() == 0 {
			return ErrTypeRequired
		}
		return f.ClientReceiver.Add(id, s)
	}
}

func (f *federator) handleClientRemove(id string) func(s *streams.Remove) error {
	return func(s *streams.Remove) error {
		if s.LenObject() == 0 {
			return ErrObjectRequired
		} else if s.LenType() == 0 {
			return ErrTypeRequired
		}
		return f.ClientReceiver.Remove(id, s)
	}
}

func (f *federator) handleClientLike(id string) func(s *streams.Like) error {
	return func(s *streams.Like) error {
		if s.LenObject() == 0 {
			return ErrObjectRequired
		}
		return f.ClientReceiver.Like(id, s)
	}
}

func (f *federator) handleClientUndo(id string) func(s *streams.Undo) error {
	return func(s *streams.Undo) error {
		if s.LenObject() == 0 {
			return ErrObjectRequired
		}
		return f.ClientReceiver.Undo(id, s)
	}
}

func (f *federator) handleClientBlock(id string) func(s *streams.Block) error {
	return func(s *streams.Block) error {
		if s.LenObject() == 0 {
			return ErrObjectRequired
		}
		return f.ClientReceiver.Block(id, s)
	}
}

func (f *federator) getPostInboxResolver(id string) *streams.Resolver {
	return &streams.Resolver{
		CreateCallback: f.handleCreate(id),
		UpdateCallback: f.handleUpdate(id),
		DeleteCallback: f.handleDelete(id),
		FollowCallback: f.handleFollow(id),
		AcceptCallback: f.handleAccept(id),
		RejectCallback: f.handleReject(id),
		AddCallback:    f.handleAdd(id),
		RemoveCallback: f.handleRemove(id),
		LikeCallback:   f.handleLike(id),
		UndoCallback:   f.handleUndo(id),
		// TODO: Extended activity types, such as Announce, Arrive, etc.
	}
}

func (f *federator) handleCreate(id string) func(s *streams.Create) error {
	return func(s *streams.Create) error {
		return f.Receiver.Create(id, s)
	}
}

func (f *federator) handleUpdate(id string) func(s *streams.Update) error {
	return func(s *streams.Update) error {
		// TODO: The receiving server MUST take care to be sure that the Update
		// is authorized to modify its object.
		return f.Receiver.Update(id, s)
	}
}

func (f *federator) handleDelete(id string) func(s *streams.Delete) error {
	return func(s *streams.Delete) error {
		// TODO: Verify ownership. I think the spec unintentionally suggests to
		// just assume it is owned, so we will actually verify.
		return f.Receiver.Delete(id, s)
	}
}

func (f *federator) handleFollow(id string) func(s *streams.Follow) error {
	return func(s *streams.Follow) error {
		return f.Receiver.Follow(id, s)
	}
}

func (f *federator) handleAccept(id string) func(s *streams.Accept) error {
	return func(s *streams.Accept) error {
		return f.Receiver.Accept(id, s)
	}
}

func (f *federator) handleReject(id string) func(s *streams.Reject) error {
	return func(s *streams.Reject) error {
		return f.Receiver.Reject(id, s)
	}
}

func (f *federator) handleAdd(id string) func(s *streams.Add) error {
	return func(s *streams.Add) error {
		return f.Receiver.Add(id, s)
	}
}

func (f *federator) handleRemove(id string) func(s *streams.Remove) error {
	return func(s *streams.Remove) error {
		return f.Receiver.Remove(id, s)
	}
}

func (f *federator) handleLike(id string) func(s *streams.Like) error {
	return func(s *streams.Like) error {
		return f.Receiver.Like(id, s)
	}
}

func (f *federator) handleUndo(id string) func(s *streams.Undo) error {
	return func(s *streams.Undo) error {
		return f.Receiver.Undo(id, s)
	}
}

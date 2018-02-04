package pub

import (
	"encoding/json"
	"fmt"
	"github.com/go-fed/activity/streams"
	"github.com/go-fed/activity/vocab"
	"io/ioutil"
	"net/http"
	"net/url"
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
}

// Receiver is provided by users of this library and designed to handle
// receiving federated messages through the Federated Protocol.
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
	// is in response to a 'Follow' then the follower should be added by the
	// client application.
	Accept(id string, s *streams.Accept) error
	// Reject can be client application specific. However, if this 'Reject'
	// is in response to a 'Follow' then the client MUST NOT go forward with
	// adding the follower.
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
	// application is responsible for enforcing this.
	Undo(id string, s *streams.Undo) error
}

type federator struct {
	// App is the client application that is ActivityPub aware.
	//
	// It is always required.
	App Application
	// Receiver provides callbacks when handling incoming messages received
	// via the Federated Protocol.
	//
	// It is only required if EnableServer is true.
	Receiver Receiver
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
		newId := f.App.NewId(id, typer)
		if !isActivityType(typer) {
			// TODO: Wrap in Create Activity
		}
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
				f.pool.Do(b, to, f.postToOutbox)
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
		// TODO: Enforce object
		// TODO: Implement
		return nil
	}
}

func (f *federator) handleClientUpdate(id string) func(s *streams.Update) error {
	return func(s *streams.Update) error {
		// TODO: Enforce object
		// TODO: Implement
		return nil
	}
}

func (f *federator) handleClientDelete(id string) func(s *streams.Delete) error {
	return func(s *streams.Delete) error {
		// TODO: Enforce object
		// TODO: Implement
		return nil
	}
}

func (f *federator) handleClientFollow(id string) func(s *streams.Follow) error {
	return func(s *streams.Follow) error {
		// TODO: Enforce object
		// TODO: Implement
		return nil
	}
}

func (f *federator) handleClientAccept(id string) func(s *streams.Accept) error {
	return func(s *streams.Accept) error {
		// TODO: Implement
		return nil
	}
}

func (f *federator) handleClientReject(id string) func(s *streams.Reject) error {
	return func(s *streams.Reject) error {
		// TODO: Implement
		return nil
	}
}

func (f *federator) handleClientAdd(id string) func(s *streams.Add) error {
	return func(s *streams.Add) error {
		// TODO: Enforce object & target
		// TODO: Implement
		return nil
	}
}

func (f *federator) handleClientRemove(id string) func(s *streams.Remove) error {
	return func(s *streams.Remove) error {
		// TODO: Enforce object & target
		// TODO: Implement
		return nil
	}
}

func (f *federator) handleClientLike(id string) func(s *streams.Like) error {
	return func(s *streams.Like) error {
		// TODO: Enforce object
		// TODO: Implement
		return nil
	}
}

func (f *federator) handleClientUndo(id string) func(s *streams.Undo) error {
	return func(s *streams.Undo) error {
		// TODO: Enforce object
		// TODO: Implement
		return nil
	}
}

func (f *federator) handleClientBlock(id string) func(s *streams.Block) error {
	return func(s *streams.Block) error {
		// TODO: Enforce object
		// TODO: Implement
		return nil
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

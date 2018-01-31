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

type clienter struct {
}

func (c *clienter) PostInbox() (http.HandlerFunc, error) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}, nil
}

func (c *clienter) handleCreate() error {
	// TODO: Enforce object
	// TODO: Implement
	return nil
}

func (c *clienter) handleUpdate() error {
	// TODO: Enforce object
	// TODO: Implement
	return nil
}

func (c *clienter) handleDelete() error {
	// TODO: Enforce object
	// TODO: Implement
	return nil
}

func (c *clienter) handleFollow() error {
	// TODO: Enforce object
	// TODO: Implement
	return nil
}

func (c *clienter) handleAccept() error {
	// TODO: Implement
	return nil
}

func (c *clienter) handleReject() error {
	// TODO: Implement
	return nil
}

func (c *clienter) handleAdd() error {
	// TODO: Enforce object & target
	// TODO: Implement
	return nil
}

func (c *clienter) handleRemove() error {
	// TODO: Enforce object & target
	// TODO: Implement
	return nil
}

func (c *clienter) handleLike() error {
	// TODO: Enforce object
	// TODO: Implement
	return nil
}

func (c *clienter) handleUndo() error {
	// TODO: Enforce object
	// TODO: Implement
	return nil
}

// FederatorStorer is a long term storage solution provided by clients so that
// data can be saved and retrieved by the ActivityPub federated server.
type FederatorStorer interface {
	// GetInbox returns the OrderedCollection inbox of the actor with the
	// provided ID.
	GetInbox(id string, r *http.Request) (vocab.OrderedCollectionType, error)
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
	// Storer is the permanent storage solution provided by the client
	// application.
	Storer FederatorStorer
	// Client is used to federate with other ActivityPub servers.
	Client *http.Client
	// Agent is the User-Agent string to use in HTTP headers when
	// federating with another server. It will automatically be appended
	// with '(go-fed ActivityPub)'.
	Agent string
	// MaxDepth is how deep collections of recipients will be expanded for
	// delivery. It must be at least 1 to be compliant with the ActivityPub
	// spec.
	MaxDepth int
	// EnableClient enables or disables the Social API, the client to
	// server part of ActivityPub. Useful if permitting remote clients to
	// act on behalf of the users of the client application.
	EnableClient bool
	// EnableServer enables or disables the Federated Protocol, or the
	// server to server part of ActivityPub. Useful to permit integrating
	// with the rest of the federative web.
	EnableServer bool
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
		oc, err := f.Storer.GetInbox(id, r)
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
		// TODO: Implement
		if f.EnableServer {
			// TODO: Hook in delivery
		}
		return true, nil
	}
}

func (f *federator) GetOutbox(id string) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) (bool, error) {
		if !isActivityPubGet(r) {
			return false, nil
		}
		// TODO: Implement
		return true, nil
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
		return f.Storer.Create(id, s)
	}
}

func (f *federator) handleUpdate(id string) func(s *streams.Update) error {
	return func(s *streams.Update) error {
		// TODO: The receiving server MUST take care to be sure that the Update
		// is authorized to modify its object.
		return f.Storer.Update(id, s)
	}
}

func (f *federator) handleDelete(id string) func(s *streams.Delete) error {
	return func(s *streams.Delete) error {
		// TODO: Verify ownership. I think the spec unintentionally suggests to
		// just assume it is owned, so we will actually verify.
		return f.Storer.Delete(id, s)
	}
}

func (f *federator) handleFollow(id string) func(s *streams.Follow) error {
	return func(s *streams.Follow) error {
		return f.Storer.Follow(id, s)
	}
}

func (f *federator) handleAccept(id string) func(s *streams.Accept) error {
	return func(s *streams.Accept) error {
		return f.Storer.Accept(id, s)
	}
}

func (f *federator) handleReject(id string) func(s *streams.Reject) error {
	return func(s *streams.Reject) error {
		return f.Storer.Reject(id, s)
	}
}

func (f *federator) handleAdd(id string) func(s *streams.Add) error {
	return func(s *streams.Add) error {
		return f.Storer.Add(id, s)
	}
}

func (f *federator) handleRemove(id string) func(s *streams.Remove) error {
	return func(s *streams.Remove) error {
		return f.Storer.Remove(id, s)
	}
}

func (f *federator) handleLike(id string) func(s *streams.Like) error {
	return func(s *streams.Like) error {
		return f.Storer.Like(id, s)
	}
}

func (f *federator) handleUndo(id string) func(s *streams.Undo) error {
	return func(s *streams.Undo) error {
		return f.Storer.Undo(id, s)
	}
}

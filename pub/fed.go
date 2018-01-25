package pub

import (
	"github.com/go-fed/activity/streams"
	"github.com/go-fed/activity/vocab"
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

// Storer is a long term storage solution provided by clients so that data can
// be saved and retrieved by the ActivityPub federated server.
type Storer interface {
	// TODO
}

// Server implements receiving the federated portion of the ActivityPub
// specification.
//
// It implements a single 'sharedinbox' to trade off numerous messages over the
// network for increased internal processing. Additionally, it will keep track
// of externally-known 'sharedInbox' in order to send public messages
// efficiently, as permitted by the spec.
//
// Fields that are able to be 'nil' are marked as such; otherwise assume that
// all fields are required.
type Server struct {
	// S is used for long-term storage and retrieval of ActivityPub
	// messages.
	S Storer
}

// OnInbox proides a handler function for when an Actor receives an Activity.
func (s *Server) OnInbox() (http.HandlerFunc, error) {
	// TODO: Implement
	return nil, nil
}

func (s *Server) OnOutbox() (http.HandlerFunc, error) {
	// TODO: Implement
	return nil, nil
}

func (s *Server) OnSharedInbox() (http.HandlerFunc, error) {
	// TODO: Implement
	return nil, nil
}

func (s *Server) OnFollowers() (http.HandlerFunc, error) {
	// TODO: Implement
	return nil, nil
}

func (s *Server) OnFollowing() (http.HandlerFunc, error) {
	// TODO: Implement
	return nil, nil
}

func (s *Server) OnLiked() (http.HandlerFunc, error) {
	// TODO: Implement
	return nil, nil
}

func (s *Server) OnLikes() (http.HandlerFunc, error) {
	// TODO: Implement
	return nil, nil
}

func (s *Server) OnShares() (http.HandlerFunc, error) {
	// TODO: Implement
	return nil, nil
}

// Client implements sending the federated portion of the ActivityPub
// specification.
type Client struct {
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
}

func (c *Client) Create() error {
	// TODO: Enforce object
	// TODO: Implement
	return nil
}

func (c *Client) Update() error {
	// TODO: Enforce object
	// TODO: Implement
	return nil
}

func (c *Client) Delete() error {
	// TODO: Enforce object
	// TODO: Implement
	return nil
}

func (c *Client) Follow() error {
	// TODO: Enforce object
	// TODO: Implement
	return nil
}

func (c *Client) Accept() error {
	// TODO: Implement
	return nil
}

func (c *Client) Reject() error {
	// TODO: Implement
	return nil
}

func (c *Client) Add() error {
	// TODO: Enforce object & target
	// TODO: Implement
	return nil
}

func (c *Client) Remove() error {
	// TODO: Enforce object & target
	// TODO: Implement
	return nil
}

func (c *Client) Like() error {
	// TODO: Enforce object
	// TODO: Implement
	return nil
}

func (c *Client) Undo() error {
	// TODO: Enforce object
	// TODO: Implement
	return nil
}

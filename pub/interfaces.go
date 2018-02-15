package pub

import (
	"context"
	"github.com/go-fed/activity/vocab"
	"net/http"
	"net/url"
	"time"
)

// HandlerFunc returns true if it was able to handle the request as an
// ActivityPub request. If it handled the request then the error should be
// checked. The response will have already been written to when handled and
// there was no error. Client applications can freely choose how to handle the
// request if this function does not handle it.
//
// Note that if the handler attempted to handle the request but returned an
// error, it is up to the client application to determine what headers and
// response to send to the requester.
type HandlerFunc func(context.Context, http.ResponseWriter, *http.Request) (bool, error)

// Clock determines the time.
type Clock interface {
	Now() time.Time
}

// Application is provided by users of this library in order to implement a
// social-federative-web application.
//
// The contexts provided in these calls are passed through this library without
// modification, allowing implementations to pass-through request-scoped data in
// order to properly handle the request.
type Application interface {
// Get fetches the ActivityStream representation of the given id.
	Get(c context.Context, id url.URL) (PubObject, error)
	// Set should write or overwrite the value of the provided object for
	// its 'id'.
	Set(c context.Context, o PubObject) error
	// GetInbox returns the OrderedCollection inbox of the actor with the
	// provided ID. It is up to the implementation to provide the correct
	// collection for the kind of authorization given in the request.
	GetInbox(c context.Context, r *http.Request) (vocab.OrderedCollectionType, error)
	// GetOutbox returns the OrderedCollection inbox of the actor with the
	// provided ID. It is up to the implementation to provide the correct
	// collection for the kind of authorization given in the request.
	GetOutbox(c context.Context, r *http.Request) (vocab.OrderedCollectionType, error)
	// PostOutboxAuthorized determines whether the request is able to post
	// an Activity to the outbox.
	PostOutboxAuthorized(c context.Context, r *http.Request) (bool, error)
	// NewId takes in a client id token and returns an ActivityStreams IRI
	// id for a new Activity posted to the outbox. The object is provided
	// as a Typer so clients can use it to decide how to generate the IRI.
	NewId(c context.Context, t Typer) url.URL
	// AddToOutboxResolver(c context.Context) (*streams.Resolver, error)
	// ActorIRI returns the actor's IRI associated with the given request.
	ActorIRI(c context.Context, r *http.Request) (url.URL, error)
}

// SocialApp is provided by users of this library and designed to handle
// receiving messages from ActivityPub clients through the Social API.
type SocialApp interface {
	// Owns returns true if the provided id is owned by this server.
	Owns(c context.Context, id url.URL) bool
	// CanAdd returns true if the provided object is allowed to be added to
	// the given target collection.
	CanAdd(c context.Context, o vocab.ObjectType, t vocab.ObjectType) bool
	// CanRemove returns true if the provided object is allowed to be
	// removed from the given target collection.
	CanRemove(c context.Context, o vocab.ObjectType, t vocab.ObjectType) bool
}

// PubObject is an ActivityPub Object.
type PubObject interface {
	GetId() url.URL
	SetId(url.URL)
	HasId() bool
	TypeLen() int
	GetType(int) interface{}
	AddType(interface{})
	RemoveType(int)
}

// Typer is an object that has a type.
type Typer interface {
	TypeLen() (l int)
	GetType(index int) (v interface{})
}

// typeIder is a Typer with additional generic capabilities.
type typeIder interface {
	Typer
	SetId(v url.URL)
	Serialize() (m map[string]interface{}, e error)
}

// actorObject is an object that has "actor" or "attributedTo" properties,
// representing the author or originator of the object.
type actorObject interface {
	HasInbox() (ok bool)
	GetInbox() (v url.URL)
	AttributedToLen() (l int)
	IsAttributedToObject(index int) (ok bool)
	GetAttributedToObject(index int) (v vocab.ObjectType)
	IsAttributedToLink(index int) (ok bool)
	GetAttributedToLink(index int) (v vocab.LinkType)
	IsAttributedToIRI(index int) (ok bool)
	GetAttributedToIRI(index int) (v url.URL)
	ActorLen() (l int)
	IsActorObject(index int) (ok bool)
	GetActorObject(index int) (v vocab.ObjectType)
	IsActorLink(index int) (ok bool)
	GetActorLink(index int) (v vocab.LinkType)
	IsActorIRI(index int) (ok bool)
	GetActorIRI(index int) (v url.URL)
}

// deliverableObject is an object that is able to be sent to recipients via the
// "to", "bto", "cc", "bcc", and "audience" objects and/or links and/or IRIs.
type deliverableObject interface {
	actorObject
	ToLen() (l int)
	IsToObject(index int) (ok bool)
	GetToObject(index int) (v vocab.ObjectType)
	IsToLink(index int) (ok bool)
	GetToLink(index int) (v vocab.LinkType)
	IsToIRI(index int) (ok bool)
	GetToIRI(index int) (v url.URL)
	BtoLen() (l int)
	IsBtoObject(index int) (ok bool)
	GetBtoObject(index int) (v vocab.ObjectType)
	RemoveBtoObject(index int)
	IsBtoLink(index int) (ok bool)
	GetBtoLink(index int) (v vocab.LinkType)
	RemoveBtoLink(index int)
	IsBtoIRI(index int) (ok bool)
	GetBtoIRI(index int) (v url.URL)
	RemoveBtoIRI(index int)
	CcLen() (l int)
	IsCcObject(index int) (ok bool)
	GetCcObject(index int) (v vocab.ObjectType)
	IsCcLink(index int) (ok bool)
	GetCcLink(index int) (v vocab.LinkType)
	IsCcIRI(index int) (ok bool)
	GetCcIRI(index int) (v url.URL)
	BccLen() (l int)
	IsBccObject(index int) (ok bool)
	GetBccObject(index int) (v vocab.ObjectType)
	RemoveBccObject(index int)
	IsBccLink(index int) (ok bool)
	GetBccLink(index int) (v vocab.LinkType)
	RemoveBccLink(index int)
	IsBccIRI(index int) (ok bool)
	GetBccIRI(index int) (v url.URL)
	RemoveBccIRI(index int)
	AudienceLen() (l int)
	IsAudienceObject(index int) (ok bool)
	GetAudienceObject(index int) (v vocab.ObjectType)
	IsAudienceLink(index int) (ok bool)
	GetAudienceLink(index int) (v vocab.LinkType)
	IsAudienceIRI(index int) (ok bool)
	GetAudienceIRI(index int) (v url.URL)
}

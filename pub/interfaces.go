package pub

import (
	"context"
	"crypto"
	"github.com/go-fed/activity/streams"
	"github.com/go-fed/activity/vocab"
	"github.com/go-fed/httpsig"
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

// HttpClient sends http requests.
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// SocialAPIVerifier will verify incoming requests from clients and is meant to
// encapsulate authentication functionality by standards such as OAuth (RFC
// 6749).
type SocialAPIVerifier interface {
	// Verify will determine the authenticated user for the given request,
	// returning false if verification fails. If the request is entirely
	// missing the required fields in order to authenticate, this function
	// must return nil and false for all values to permit attempting
	// validation by HTTP Signatures. If there was an internal error
	// determining the authentication of the request, it is returned.
	//
	// Return values are interpreted as follows:
	//     (userFoo, true,  true,  <nil>) => userFoo passed authentication and is authorized
	//     (<any>,   true,  false, <nil>) => a user passed authentication but failed authorization (Permission denied)
	//     (<any>,   false, false, <nil>) => authentication failed: deny access (Bad request)
	//     (<nil>,   false, true,  <nil>) => authentication failed: must pass HTTP Signature verification or will be Permission Denied
	//     (<nil>,   true,  true,  <nil>) => "I don't care, try to validate using HTTP Signatures. If that still doesn't work, permit raw requests access anyway."
	//     (<any>,   <any>, <any>, error) => an internal error occurred during validation
	//
	// Be very careful that the 'authenticatedUser' value is non-nil when
	// returning 'authn' and 'authz' values of true, or else the library
	// will use the most permissive logic instead of the most restrictive as
	// outlined above.
	Verify(r *http.Request) (authenticatedUser *url.URL, authn, authz bool, err error)
	// VerifyForOutbox is the same as Verify, except that the request must
	// authenticate the owner of the provided outbox IRI.
	//
	// Return values are interpreted as follows:
	//     (true,  true,   <nil>) => user for the outbox passed authentication and is authorized
	//     (true,  false,  <nil>) => a user passed authentication but failed authorization for this outbox (Permission denied)
	//     (false, true,   <nil>) => authentication failed: must pass HTTP Signature verification or will be Permission Denied
	//     (false, false,  <nil>) => authentication failed: deny access (Bad request)
	//     (<any>, <any>,  error) => an internal error occurred during validation
	VerifyForOutbox(r *http.Request, outbox *url.URL) (authn, authz bool, err error)
}

// Application is provided by users of this library in order to implement a
// social-federative-web application.
//
// The contexts provided in these calls are passed through this library without
// modification, allowing implementations to pass-through request-scoped data in
// order to properly handle the request.
type Application interface {
	// Owns returns true if the provided id is owned by this server.
	Owns(c context.Context, id *url.URL) bool
	// Get fetches the ActivityStream representation of the given id.
	Get(c context.Context, id *url.URL, rw RWType) (PubObject, error)
	// GetAsVerifiedUser fetches the ActivityStream representation of the
	// given id with the provided IRI representing the authenticated user
	// making the request.
	GetAsVerifiedUser(c context.Context, id, authdUser *url.URL, rw RWType) (PubObject, error)
	// Has determines if the server already knows about the object or
	// Activity specified by the given id.
	Has(c context.Context, id *url.URL) (bool, error)
	// Set should write or overwrite the value of the provided object for
	// its 'id'.
	Set(c context.Context, o PubObject) error
	// GetInbox returns the OrderedCollection inbox of the actor for this
	// context. It is up to the implementation to provide the correct
	// collection for the kind of authorization given in the request.
	GetInbox(c context.Context, r *http.Request, rw RWType) (vocab.OrderedCollectionType, error)
	// GetOutbox returns the OrderedCollection inbox of the actor for this
	// context. It is up to the implementation to provide the correct
	// collection for the kind of authorization given in the request.
	GetOutbox(c context.Context, r *http.Request, rw RWType) (vocab.OrderedCollectionType, error)
	// NewId takes in a client id token and returns an ActivityStreams IRI
	// id for a new Activity posted to the outbox. The object is provided
	// as a Typer so clients can use it to decide how to generate the IRI.
	NewId(c context.Context, t Typer) *url.URL
	// GetPublicKey fetches the public key for a user based on the public
	// key id. It also determines which algorithm to use to verify the
	// signature.
	GetPublicKey(c context.Context, publicKeyId string) (pubKey crypto.PublicKey, algo httpsig.Algorithm, user *url.URL, err error)
	// CanAdd returns true if the provided object is allowed to be added to
	// the given target collection. Applicable to either or both of the
	// SocialAPI and FederateAPI.
	CanAdd(c context.Context, o vocab.ObjectType, t vocab.ObjectType) bool
	// CanRemove returns true if the provided object is allowed to be
	// removed from the given target collection. Applicable to either or
	// both of the SocialAPI and FederateAPI.
	CanRemove(c context.Context, o vocab.ObjectType, t vocab.ObjectType) bool
}

// RWType indicates the kind of reading being done.
type RWType bool

const (
	// Read indicates the object is only being read.
	Read RWType = false
	// ReadWrite indicates the object is being mutated as well.
	ReadWrite = true
)

// SocialAPI is provided by users of this library and designed to handle
// receiving messages from ActivityPub clients through the Social API.
type SocialAPI interface {
	// ActorIRI returns the actor's IRI associated with the given request.
	ActorIRI(c context.Context, r *http.Request) (*url.URL, error)
	// GetSocialAPIVerifier returns the authentication mechanism used for
	// incoming ActivityPub client requests. It is optional and allowed to
	// return null.
	//
	// Note that regardless of what this implementation returns, HTTP
	// Signatures is supported natively as a fallback.
	GetSocialAPIVerifier(c context.Context) SocialAPIVerifier
	// GetPublicKeyForOutbox fetches the public key for a user based on the
	// public key id. It also determines which algorithm to use to verify
	// the signature.
	//
	// Note that a key difference from Application's GetPublicKey is that
	// this function must make sure that the actor whose boxIRI is passed in
	// matches the public key id that is requested, or return an error.
	GetPublicKeyForOutbox(c context.Context, publicKeyId string, boxIRI *url.URL) (crypto.PublicKey, httpsig.Algorithm, error)
}

// FederateAPI is provided by users of this library and designed to handle
// receiving messages from ActivityPub servers through the Federative API.
type FederateAPI interface {
	// OnFollow determines whether to take any automatic reactions in
	// response to this follow. Note that if this application does not own
	// an object on the activity, then the 'AutomaticAccept' and
	// 'AutomaticReject' results will behave as if they were 'DoNothing'.
	OnFollow(c context.Context, s *streams.Follow) FollowResponse
	// Unblocked should return an error if the provided actor ids are not
	// able to interact with this particular end user due to being blocked
	// or other application-specific logic. This error is passed
	// transparently back to the request thread via PostInbox.
	//
	// If nil error is returned, then the received activity is processed as
	// a normal unblocked interaction.
	Unblocked(c context.Context, actorIRIs []*url.URL) error
	// FilterForwarding is invoked when a received activity needs to be
	// forwarded to specific inboxes owned by this server in order to avoid
	// the ghost reply problem. The IRIs provided are collections owned by
	// this server that the federate peer requested inbox forwarding to.
	//
	// Implementors must apply some sort of filtering to the provided IRI
	// collections. Without any filtering, the resulting application is
	// vulnerable to becoming a spam bot for a malicious federate peer.
	// Typical implementations will filter the iris down to be only the
	// follower collections owned by the actors targeted in the activity.
	FilterForwarding(c context.Context, activity vocab.ActivityType, iris []*url.URL) ([]*url.URL, error)
	// NewSigner returns a new httpsig.Signer for which deliveries can be
	// signed by the actor delivering the Activity. Let me take this moment
	// to really level with you, dear anonymous reader-of-documentation. You
	// want to use httpsig.RSA_SHA256 as the algorithm. Otherwise, your app
	// will not federate correctly and peers will reject the signatures. All
	// other known implementations using HTTP Signatures use RSA_SHA256,
	// hardcoded just like your implementation will be.
	//
	// Some people might think it funny to split the federation and use
	// their own algorithm. And while I give you the power to build the
	// largest foot-gun possible to blow away your limbs because I respect
	// your freedom, you as a developer have the responsibility to also be
	// cognizant of the wider community you are building for. Don't be a
	// dick.
	//
	// The headers available for inclusion in the signature are:
	//     Date
	//     User-Agent
	NewSigner() (httpsig.Signer, error)
	// PrivateKey fetches the private key and its associated public key ID.
	// The given URL is the inbox or outbox for the actor whose key is
	// needed.
	PrivateKey(boxIRI *url.URL) (privKey crypto.PrivateKey, pubKeyId string, err error)
}

// SocialApp is an implementation only for the Social API part of the
// ActivityPub specification.
type SocialApplication interface {
	Application
	SocialAPI
}

// FederateApp is an implementation only for the Federating API part of the
// ActivityPub specification.
type FederateApplication interface {
	Application
	FederateAPI
}

// SocialFederateApplication is an implementation for both the Social API and
// the Federating API parts of the ActivityPub specification.
type SocialFederateApplication interface {
	Application
	SocialAPI
	FederateAPI
}

// FollowResponse instructs how to proceed upon immediately receiving a request
// to follow.
type FollowResponse uint8

const (
	AutomaticAccept FollowResponse = iota
	AutomaticReject
	DoNothing
)

// Callbacker provides an Application hooks into the lifecycle of the
// ActivityPub processes for both client-to-server and server-to-server
// interactions. These callbacks are called after their spec-compliant actions
// are completed, but before inbox forwarding and before delivery.
//
// Note that at minimum, for inbox forwarding to work correctly, these
// Activities must be stored in the client application as a system of record.
//
// Note that modifying the ActivityStream objects in a callback may cause
// unintentionally non-standard behavior if modifying core attributes, but
// otherwise affords clients powerful flexibility. Use responsibly.
type Callbacker interface {
	// Create Activity callback.
	Create(c context.Context, s *streams.Create) error
	// Update Activity callback.
	Update(c context.Context, s *streams.Update) error
	// Delete Activity callback.
	Delete(c context.Context, s *streams.Delete) error
	// Add Activity callback.
	Add(c context.Context, s *streams.Add) error
	// Remove Activity callback.
	Remove(c context.Context, s *streams.Remove) error
	// Like Activity callback.
	Like(c context.Context, s *streams.Like) error
	// Block Activity callback. By default, this implmentation does not
	// dictate how blocking should be implemented, so it is up to the
	// application to enforce this by implementing the FederateApp
	// interface.
	Block(c context.Context, s *streams.Block) error
	// Follow Activity callback. In the special case of server-to-server
	// delivery of a Follow activity, this implementation supports the
	// option of automatically replying with an 'Accept', 'Reject', or
	// waiting for human interaction as provided in the FederateApp
	// interface.
	//
	// In the special case that the FederateApp returned AutomaticAccept,
	// this library automatically handles adding the 'actor' to the
	// 'followers' collection of the 'object'.
	Follow(c context.Context, s *streams.Follow) error
	// Undo Activity callback. It is up to the client to provide support
	// for all 'Undo' operations; this implementation does not attempt to
	// provide a generic implementation.
	Undo(c context.Context, s *streams.Undo) error
	// Accept Activity callback. In the special case that this 'Accept'
	// activity has an 'object' of 'Follow' type, then the library will
	// handle adding the 'actor' to the 'following' collection of the
	// original 'actor' who requested the 'Follow'.
	Accept(c context.Context, s *streams.Accept) error
	// Reject Activity callback. Note that in the special case that this
	// 'Reject' activity has an 'object' of 'Follow' type, then the client
	// MUST NOT add the 'actor' to the 'following' collection of the
	// original 'actor' who requested the 'Follow'.
	Reject(c context.Context, s *streams.Reject) error
}

// Deliverer schedules federated ActivityPub messages for delivery, possibly
// asynchronously.
type Deliverer interface {
	// Do schedules a message to be sent to a specific URL endpoint by
	// using toDo.
	Do(b []byte, to *url.URL, toDo func(b []byte, u *url.URL) error)
}

// PubObject is an ActivityPub Object.
type PubObject interface {
	vocab.Serializer
	Typer
	GetId() *url.URL
	SetId(*url.URL)
	HasId() bool
	AppendType(interface{})
	RemoveType(int)
}

// Typer is an object that has a type.
type Typer interface {
	vocab.Typer
}

// typeIder is a Typer with additional generic capabilities.
type typeIder interface {
	Typer
	SetId(v *url.URL)
	Serialize() (m map[string]interface{}, e error)
}

// actor is an object that is an ActivityPub Actor. The specification is more
// strict than what we include here, only for our internal use.
type actor interface {
	IsInboxAnyURI() (ok bool)
	GetInboxAnyURI() (v *url.URL)
	IsInboxOrderedCollection() (ok bool)
	GetInboxOrderedCollection() (v vocab.OrderedCollectionType)
}

var _ actor = &vocab.Object{}

// actorObject is an object that has "actor" or "attributedTo" properties,
// representing the author or originator of the object.
type actorObject interface {
	IsInboxAnyURI() (ok bool)
	GetInboxAnyURI() (v *url.URL)
	IsInboxOrderedCollection() (ok bool)
	GetInboxOrderedCollection() (v vocab.OrderedCollectionType)
	AttributedToLen() (l int)
	IsAttributedToObject(index int) (ok bool)
	GetAttributedToObject(index int) (v vocab.ObjectType)
	IsAttributedToLink(index int) (ok bool)
	GetAttributedToLink(index int) (v vocab.LinkType)
	IsAttributedToIRI(index int) (ok bool)
	GetAttributedToIRI(index int) (v *url.URL)
	ActorLen() (l int)
	IsActorObject(index int) (ok bool)
	GetActorObject(index int) (v vocab.ObjectType)
	IsActorLink(index int) (ok bool)
	GetActorLink(index int) (v vocab.LinkType)
	IsActorIRI(index int) (ok bool)
	GetActorIRI(index int) (v *url.URL)
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
	GetToIRI(index int) (v *url.URL)
	BtoLen() (l int)
	IsBtoObject(index int) (ok bool)
	GetBtoObject(index int) (v vocab.ObjectType)
	RemoveBtoObject(index int)
	IsBtoLink(index int) (ok bool)
	GetBtoLink(index int) (v vocab.LinkType)
	RemoveBtoLink(index int)
	IsBtoIRI(index int) (ok bool)
	GetBtoIRI(index int) (v *url.URL)
	RemoveBtoIRI(index int)
	CcLen() (l int)
	IsCcObject(index int) (ok bool)
	GetCcObject(index int) (v vocab.ObjectType)
	IsCcLink(index int) (ok bool)
	GetCcLink(index int) (v vocab.LinkType)
	IsCcIRI(index int) (ok bool)
	GetCcIRI(index int) (v *url.URL)
	BccLen() (l int)
	IsBccObject(index int) (ok bool)
	GetBccObject(index int) (v vocab.ObjectType)
	RemoveBccObject(index int)
	IsBccLink(index int) (ok bool)
	GetBccLink(index int) (v vocab.LinkType)
	RemoveBccLink(index int)
	IsBccIRI(index int) (ok bool)
	GetBccIRI(index int) (v *url.URL)
	RemoveBccIRI(index int)
	AudienceLen() (l int)
	IsAudienceObject(index int) (ok bool)
	GetAudienceObject(index int) (v vocab.ObjectType)
	IsAudienceLink(index int) (ok bool)
	GetAudienceLink(index int) (v vocab.LinkType)
	IsAudienceIRI(index int) (ok bool)
	GetAudienceIRI(index int) (v *url.URL)
}

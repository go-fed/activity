package pub

import (
	"context"
	"github.com/go-fed/activity/streams/vocab"
	"net/http"
	"net/url"
)

// FederatingProtocol contains behaviors an application needs to satisfy for the
// full ActivityPub S2S implementation to be supported by this library.
//
// It is only required if the client application wants to support the server-to-
// server, or federating, protocol.
//
// It is passed to the library as a dependency injection from the client
// application.
type FederatingProtocol interface {
	// AuthenticatePostInbox delegates the authentication of a POST to an
	// inbox.
	//
	// If an error is returned, it is passed back to the caller of
	// PostInbox. In this case, the implementation must not write a
	// response to the ResponseWriter as is expected that the client will
	// do so when handling the error. The 'shouldReturn' is ignored.
	//
	// If no error is returned, but authentication or authorization fails,
	// then shouldReturn must be true and error nil. It is expected that
	// the implementation handles writing to the ResponseWriter in this
	// case.
	//
	// Finally, if the authentication and authorization succeeds, then
	// shouldReturn must be false and error nil. The request will continue
	// to be processed.
	AuthenticatePostInbox(c context.Context, w http.ResponseWriter, r *http.Request) (shouldReturn bool, err error)
	// Blocked should determine whether to permit a set of actors given by
	// their ids are able to interact with this particular end user due to
	// being blocked or other application-specific logic.
	//
	// If an error is returned, it is passed back to the caller of
	// PostInbox.
	//
	// If no error is returned, but authentication or authorization fails,
	// then shouldReturn must be true and error nil. An http.StatusForbidden
	// will be written in the wresponse.
	//
	// Finally, if the authentication and authorization succeeds, then
	// shouldReturn must be false and error nil. The request will continue
	// to be processed.
	Blocked(c context.Context, actorIRIs []*url.URL) (blocked bool, err error)
	// Callbacks returns the application logic that handles ActivityStreams
	// received from federating peers.
	//
	// Note that certain types of callbacks will be 'wrapped' with default
	// behaviors supported natively by the library. Other callbacks
	// compatible with streams.TypeResolver can be specified by 'other'.
	//
	// For example, setting the 'Create' field in the
	// FederatingWrappedCallbacks lets an application dependency inject
	// additional behaviors they want to take place, including the default
	// behavior supplied by this library. This is guaranteed to be compliant
	// with the ActivityPub Social protocol.
	//
	// To override the default behavior, instead supply the function in
	// 'other', which does not guarantee the application will be compliant
	// with the ActivityPub Social Protocol.
	Callbacks(c context.Context) (wrapped FederatingWrappedCallbacks, other []interface{})
	// MaxInboxForwardingRecursionDepth determines how deep to search within
	// an activity to determine if inbox forwarding needs to occur.
	//
	// Zero or negative numbers indicate infinite recursion.
	MaxInboxForwardingRecursionDepth(c context.Context) int
	// MaxDeliveryRecursionDepth determines how deep to search within
	// collections owned by peers when they are targeted to receive a
	// delivery.
	//
	// Zero or negative numbers indicate infinite recursion.
	MaxDeliveryRecursionDepth(c context.Context) int
	// FilterForwarding allows the implementation to apply business logic
	// such as blocks, spam filtering, and so on to a list of potential
	// Collections and OrderedCollections of recipients when inbox
	// forwarding has been triggered.
	//
	// The activity is provided as a reference for more intelligent
	// logic to be used, but the implementation must not modify it.
	FilterForwarding(c context.Context, potentialRecipients []*url.URL, a Activity) (filteredRecipients []*url.URL, err error)
	// NewTransport returns a new Transport on behalf of a specific actor.
	//
	// The actorBoxIRI will be either the inbox or outbox of an actor who is
	// attempting to do the dereferencing or delivery. Any authentication
	// scheme applied on the request must be based on this actor. The
	// request must contain some sort of credential of the user, such as a
	// HTTP Signature.
	//
	// The gofedAgent passed in should be used by the Transport
	// implementation in the User-Agent, as well as the application-specific
	// user agent string. The gofedAgent will indicate this library's use as
	// well as the library's version number.
	//
	// Any server-wide rate-limiting that needs to occur should happen in a
	// Transport implementation. This factory function allows this to be
	// created, so peer servers are not DOS'd.
	//
	// Any retry logic should also be handled by the Transport
	// implementation.
	//
	// Note that the library will not maintain a long-lived pointer to the
	// returned Transport so that any private credentials are able to be
	// garbage collected.
	NewTransport(c context.Context, actorBoxIRI *url.URL, gofedAgent string) (t Transport, err error)
	// GetInbox returns the OrderedCollection inbox of the actor for this
	// context. It is up to the implementation to provide the correct
	// collection for the kind of authorization given in the request.
	//
	// AuthenticateGetInbox will be called prior to this.
	//
	// Always called, regardless whether the Federated Protocol or Social
	// API is enabled.
	GetInbox(c context.Context, r *http.Request) (vocab.ActivityStreamsOrderedCollectionPage, error)
}

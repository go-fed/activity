package pub

import (
	"context"
	"github.com/go-fed/activity/streams/vocab"
	"net/http"
	"net/url"
)

// SocialProtocol contains behaviors an application needs to satisfy for the
// full ActivityPub C2S implementation to be supported by this library.
//
// It is only required if the client application wants to support the client-to-
// server, or social, protocol.
type SocialProtocol interface {
	// AuthenticatePostOutbox delegates the authentication of a POST to an
	// outbox.
	//
	// Only called if the Social API is enabled.
	//
	// If an error is returned, it is passed back to the caller of
	// PostOutbox. In this case, the implementation must not write a
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
	AuthenticatePostOutbox(c context.Context, w http.ResponseWriter, r *http.Request) (shouldReturn bool, err error)
	// Callbacks returns the application logic that handles ActivityStreams
	// received from C2S clients. Note that certain types of callbacks
	// will be 'wrapped' with default behaviors supported natively by the
	// library. Other callbacks compatible with streams.TypeResolver can
	// be specified by 'other'.
	//
	// Note that the functions in 'wrapped' cannot be provided in 'other'.
	Callbacks(c context.Context) (wrapped SocialWrappedCallbacks, other []interface{})
	// ActorIRI fetches the outbox's actor's IRI.
	ActorIRI(c context.Context, outboxIRI *url.URL) (actorIRI *url.URL, err error)
	// GetOutbox returns the OrderedCollection inbox of the actor for this
	// context. It is up to the implementation to provide the correct
	// collection for the kind of authorization given in the request.
	//
	// AuthenticateGetOutbox will be called prior to this.
	//
	// Always called, regardless whether the Federated Protocol or Social
	// API is enabled.
	GetOutbox(c context.Context, r *http.Request) (vocab.ActivityStreamsOrderedCollectionPage, error)
}

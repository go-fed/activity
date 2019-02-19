package pub

import (
	"context"
	"github.com/go-fed/activity/streams/vocab"
	"net/http"
)

// SocialProtocol contains behaviors an application needs to satisfy for the
// full ActivityPub C2S implementation to be supported by this library.
//
// It is only required if the client application wants to support the client-to-
// server, or social, protocol.
//
// It is passed to the library as a dependency injection from the client
// application.
type SocialProtocol interface {
	// AuthenticatePostOutbox delegates the authentication of a POST to an
	// outbox.
	//
	// Only called if the Social API is enabled.
	//
	// If an error is returned, it is passed back to the caller of
	// PostOutbox. In this case, the implementation must not write a
	// response to the ResponseWriter as is expected that the client will
	// do so when handling the error. The 'authenticated' is ignored.
	//
	// If no error is returned, but authentication or authorization fails,
	// then authenticated must be false and error nil. It is expected that
	// the implementation handles writing to the ResponseWriter in this
	// case.
	//
	// Finally, if the authentication and authorization succeeds, then
	// authenticated must be true and error nil. The request will continue
	// to be processed.
	AuthenticatePostOutbox(c context.Context, w http.ResponseWriter, r *http.Request) (authenticated bool, err error)
	// Callbacks returns the application logic that handles ActivityStreams
	// received from C2S clients.
	//
	// Note that certain types of callbacks will be 'wrapped' with default
	// behaviors supported natively by the library. Other callbacks
	// compatible with streams.TypeResolver can be specified by 'other'.
	//
	// For example, setting the 'Create' field in the SocialWrappedCallbacks
	// lets an application dependency inject additional behaviors they want
	// to take place, including the default behavior supplied by this
	// library. This is guaranteed to be compliant with the ActivityPub
	// Social protocol.
	//
	// To override the default behavior, instead supply the function in
	// 'other', which does not guarantee the application will be compliant
	// with the ActivityPub Social Protocol.
	Callbacks(c context.Context) (wrapped SocialWrappedCallbacks, other []interface{})
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

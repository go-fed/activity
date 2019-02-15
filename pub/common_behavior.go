package pub

import (
	"context"
	"net/http"
)

// Common contains functions required for both the Social API and Federating
// Protocol.
//
// It is passed to the library as a dependency injection from the client
// application.
type CommonBehavior interface {
	// AuthenticateGetInbox delegates the authentication of a GET to an
	// inbox.
	//
	// Always called, regardless whether the Federated Protocol or Social
	// API is enabled.
	//
	// If an error is returned, it is passed back to the caller of
	// GetInbox. In this case, the implementation must not write a
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
	AuthenticateGetInbox(c context.Context, w http.ResponseWriter, r *http.Request) (shouldReturn bool, err error)
	// AuthenticateGetOutbox delegates the authentication of a GET to an
	// outbox.
	//
	// Always called, regardless whether the Federated Protocol or Social
	// API is enabled.
	//
	// If an error is returned, it is passed back to the caller of
	// GetOutbox. In this case, the implementation must not write a
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
	AuthenticateGetOutbox(c context.Context, w http.ResponseWriter, r *http.Request) (shouldReturn bool, err error)
}

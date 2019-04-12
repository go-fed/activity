package pub

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-fed/activity/streams"
	"net/http"
)

// HandlerFunc determines whether an incoming HTTP request is an ActivityStreams
// GET request, and if so attempts to serve ActivityStreams data.
//
// If an error is returned, then the calling function is responsible for writing
// to the ResponseWriter as part of error handling.
//
// If 'isASRequest' is false and there is no error, then the calling function
// may continue processing the request, and the HandlerFunc will not have
// written anything to the ResponseWriter. For example, a webpage may be served
// instead.
//
// If 'isASRequest' is true and there is no error, then the HandlerFunc
// successfully served the request and wrote to the ResponseWriter.
type HandlerFunc func(c context.Context, w http.ResponseWriter, r *http.Request) (isASRequest bool, err error)

// AuthenticateFunc is responsible for authenticating and authorizing a GET
// ActivityStreams request.
//
// If an error is returned, 'shouldReturn' is ignored. It is expected that the
// calling function will write to the ResponseWriter while handling the error.
//
// If 'shouldReturn' is true and no error is returned, then this function
// immediately returns to the caller. This function is responsible for writing
// the authentication or authorization failure on the ResponseWriter.
//
// If 'shouldReturn' is false and no error is returned, then processing of the
// request will continue.
type AuthenticateFunc func(c context.Context, w http.ResponseWriter, r *http.Request) (shouldReturn bool, err error)

// NewActivityStreamsHandler creates a HandlerFunc to serve ActivityStreams
// requests which are coming from other clients or servers that wish to obtain
// an ActivityStreams representation of data.
//
// Strips retrieved ActivityStreams values of sensitive fields ('bto' and 'bcc')
// before responding with them. Sets the appropriate HTTP status code for
// Tombstone Activities as well.
func NewActivityStreamsHandler(authFn AuthenticateFunc, db Database, clock Clock) HandlerFunc {
	return func(c context.Context, w http.ResponseWriter, r *http.Request) (isASRequest bool, err error) {
		// Do nothing if it is not an ActivityPub GET request
		if !isActivityPubGet(r) {
			return
		}
		isASRequest = true
		// Authenticate the request
		var shouldReturn bool
		if shouldReturn, err = authFn(c, w, r); err != nil {
			return
		} else if shouldReturn {
			return
		}
		id := requestId(r)
		// Lock and obtain a copy of the requested ActivityStreams value
		err = db.Lock(c, id)
		if err != nil {
			return
		}
		// WARNING: Unlock not deferred
		t, err := db.Get(c, id)
		if err != nil {
			db.Unlock(c, id)
			return
		}
		db.Unlock(c, id)
		// Unlock must have been called by this point and in every
		// branch above
		//
		// Remove sensitive fields.
		clearSensitiveFields(t)
		// Serialize the fetched value.
		m, err := serialize(t)
		if err != nil {
			return
		}
		raw, err := json.Marshal(m)
		if err != nil {
			return
		}
		// Construct the response.
		addResponseHeaders(w.Header(), clock, raw)
		// Write the response.
		if streams.IsOrExtendsActivityStreamsTombstone(t) {
			w.WriteHeader(http.StatusGone)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		n, err := w.Write(raw)
		if err != nil {
			return
		} else if n != len(raw) {
			err = fmt.Errorf("only wrote %d of %d bytes", n, len(raw))
			return
		}
		return
	}
}

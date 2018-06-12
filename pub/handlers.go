package pub

import (
	"context"
	"crypto"
	"encoding/json"
	"fmt"
	"github.com/go-fed/activity/vocab"
	"github.com/go-fed/httpsig"
	"net/http"
	"net/url"
)

// ServeActivityPubObject will serve the ActivityPub object with the given IRI
// in the request. Note that requests may be signed with HTTP signatures or be
// permitted without any authentication scheme. To change this default behavior,
// use ServeActivityPubObjectWithVerificationMethod instead.
func ServeActivityPubObject(a Application, clock Clock) HandlerFunc {
	return func(c context.Context, w http.ResponseWriter, r *http.Request) (bool, error) {
		return serveActivityPubObject(c, a, clock, w, r, nil)
	}
}

// ServeActivityPubObjectWithVerificationMethod will serve the ActivityPub
// object with the given IRI in the request. The rules for accessing the data
// are governed by the SocialAPIVerifier's behavior and may permit accessing
// data without having any credentials in the request.
func ServeActivityPubObjectWithVerificationMethod(a Application, clock Clock, verifierFn func(context.Context) SocialAPIVerifier) HandlerFunc {
	return func(c context.Context, w http.ResponseWriter, r *http.Request) (bool, error) {
		if verifierFn != nil {
			verifier := verifierFn(c)
			return serveActivityPubObject(c, a, clock, w, r, verifier)
		} else {
			return serveActivityPubObject(c, a, clock, w, r, nil)
		}
	}
}

func serveActivityPubObject(c context.Context, a Application, clock Clock, w http.ResponseWriter, r *http.Request, verifier SocialAPIVerifier) (handled bool, err error) {
	handled = isActivityPubGet(r)
	if !handled {
		return
	}
	id := r.URL
	if !a.Owns(c, id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var verifiedUser *url.URL
	// By default, permit unsigned access to resource. however, if there is
	// an HTTP Signature present, it must pass validation.
	authenticated := false
	authorized := false
	if verifier != nil {
		verifiedUser, authenticated, authorized, err = verifier.Verify(r)
		if err != nil {
			return
		} else if authenticated && !authorized {
			w.WriteHeader(http.StatusForbidden)
			return
		} else if !authenticated && !authorized {
			w.WriteHeader(http.StatusBadRequest)
			return
		} else if !authenticated && authorized {
			// Protect against bad implementations: There is no
			// recognized reason for an implementation to pass back
			// a non-nil verifiedUser that is authorized but not
			// authenticated.
			//
			// Force HTTP Signature validation to trigger by
			// ensuring the verifiedUser is nil.
			if verifiedUser != nil {
				verifiedUser = nil
			}
		}
	}
	if verifiedUser == nil {
		var v httpsig.Verifier
		v, err = httpsig.NewVerifier(r)
		if err != nil { // Unsigned request
			if !authenticated && authorized { // Must pass HTTP Signature verification
				w.WriteHeader(http.StatusBadRequest)
				err = nil
				return
			} // Else permit unsigned requests access
		} else { // Signed request
			var publicKey crypto.PublicKey
			var algo httpsig.Algorithm
			var user *url.URL
			publicKey, algo, user, err = a.GetPublicKey(c, v.KeyId())
			if err != nil {
				return
			}
			err = v.Verify(publicKey, algo)
			if err != nil && !authenticated { // Failed and must pass HTTP Signature verification
				w.WriteHeader(http.StatusForbidden)
				err = nil
				return
			} else if err == nil {
				verifiedUser = user
			} // Else failed HTTP Signature verification but we still allow access.
		}
	}
	var pObj PubObject
	if verifiedUser != nil {
		pObj, err = a.GetAsVerifiedUser(c, r.URL, verifiedUser, Read)
	} else {
		pObj, err = a.Get(c, r.URL, Read)
	}
	if err != nil {
		return
	}
	if obj, ok := pObj.(vocab.ObjectType); ok {
		clearSensitiveFields(obj)
	}
	var m map[string]interface{}
	m, err = pObj.Serialize()
	if err != nil {
		return
	}
	addJSONLDContext(m)
	var b []byte
	b, err = json.Marshal(m)
	if err != nil {
		return
	}
	addResponseHeaders(w.Header(), clock, b)
	if vocab.HasTypeTombstone(pObj) {
		w.WriteHeader(http.StatusGone)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	n, err := w.Write(b)
	if err != nil {
		return
	} else if n != len(b) {
		err = fmt.Errorf("ResponseWriter.Write wrote %d of %d bytes", n, len(b))
		return
	}
	return
}

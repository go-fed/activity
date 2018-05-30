package pub

import (
	"context"
	"crypto"
	"github.com/go-fed/activity/vocab"
	"github.com/go-fed/httpsig"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestServeActivityPubObject(t *testing.T) {
	tests := []struct {
		name          string
		app           *MockApplication
		clock         *MockClock
		input         *http.Request
		expectedCode  int
		expectedObjFn func() vocab.Serializer
		expectHandled bool
	}{
		{
			name: "unsigned request",
			app: &MockApplication{
				t: t,
				get: func(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
					if rw != Read {
						t.Fatalf("expected RWType of %d, got %d", Read, rw)
					} else if s := id.String(); s != noteURIString {
						t.Fatalf("expected %s, got %s", noteURIString, s)
					}
					testNote = &vocab.Note{}
					testNote.SetId(noteIRI)
					testNote.AppendNameString(noteName)
					testNote.AppendContentString("This is a simple note")
					return testNote, nil
				},
				owns: func(c context.Context, id *url.URL) bool {
					if s := id.String(); s != noteURIString {
						t.Fatalf("expected %s, got %s", noteURIString, s)
					}
					return true
				},
			},
			clock:        &MockClock{now},
			input:        ActivityPubRequest(httptest.NewRequest("GET", noteURIString, nil)),
			expectedCode: http.StatusOK,
			expectedObjFn: func() vocab.Serializer {
				testNote = &vocab.Note{}
				testNote.SetId(noteIRI)
				testNote.AppendNameString(noteName)
				testNote.AppendContentString("This is a simple note")
				return testNote
			},
			expectHandled: true,
		},
		{
			name:  "http signature request",
			input: Sign(ActivityPubRequest(httptest.NewRequest("GET", noteURIString, nil))),
			app: &MockApplication{
				t: t,
				getPublicKey: func(c context.Context, publicKeyId string) (crypto.PublicKey, httpsig.Algorithm, *url.URL, error) {
					if publicKeyId != testPublicKeyId {
						t.Fatalf("expected %s, got %s", testPublicKeyId, publicKeyId)
					}
					return testPrivateKey.Public(), httpsig.RSA_SHA256, samIRI, nil
				},
				getAsVerifiedUser: func(c context.Context, id, user *url.URL, rw RWType) (PubObject, error) {
					if rw != Read {
						t.Fatalf("expected RWType of %d, got %d", Read, rw)
					} else if s := id.String(); s != noteURIString {
						t.Fatalf("expected %s, got %s", noteURIString, s)
					} else if u := user.String(); u != samIRIString {
						t.Fatalf("expected %s, got %s", samIRIString, u)
					}
					testNote = &vocab.Note{}
					testNote.SetId(noteIRI)
					testNote.AppendNameString(noteName)
					testNote.AppendContentString("This is a simple note")
					return testNote, nil
				},
				owns: func(c context.Context, id *url.URL) bool {
					if s := id.String(); s != noteURIString {
						t.Fatalf("expected %s, got %s", noteURIString, s)
					}
					return true
				},
			},
			clock:        &MockClock{now},
			expectedCode: http.StatusOK,
			expectedObjFn: func() vocab.Serializer {
				testNote = &vocab.Note{}
				testNote.SetId(noteIRI)
				testNote.AppendNameString(noteName)
				testNote.AppendContentString("This is a simple note")
				return testNote
			},
			expectHandled: true,
		},
		{
			name:  "not owned",
			input: ActivityPubRequest(httptest.NewRequest("GET", noteURIString, nil)),
			app: &MockApplication{
				t: t,
				owns: func(c context.Context, id *url.URL) bool {
					if s := id.String(); s != noteURIString {
						t.Fatalf("expected %s, got %s", noteURIString, s)
					}
					return false
				},
			},
			expectedCode:  http.StatusNotFound,
			expectHandled: true,
		},
		{
			name:          "not activitypub get",
			input:         httptest.NewRequest("GET", noteURIString, nil),
			expectHandled: false,
		},
		{
			name:  "bad http signature",
			input: BadSignature(ActivityPubRequest(httptest.NewRequest("GET", noteURIString, nil))),
			app: &MockApplication{
				t: t,
				getPublicKey: func(c context.Context, publicKeyId string) (crypto.PublicKey, httpsig.Algorithm, *url.URL, error) {
					if publicKeyId != testPublicKeyId {
						t.Fatalf("expected %s, got %s", testPublicKeyId, publicKeyId)
					}
					return testPrivateKey.Public(), httpsig.RSA_SHA256, samIRI, nil
				},
				owns: func(c context.Context, id *url.URL) bool {
					if s := id.String(); s != noteURIString {
						t.Fatalf("expected %s, got %s", noteURIString, s)
					}
					return true
				},
			},
			expectedCode:  http.StatusForbidden,
			expectHandled: true,
		},
	}
	for _, test := range tests {
		t.Logf("Running table test case %q", test.name)
		resp := httptest.NewRecorder()
		fnUnderTest := ServeActivityPubObject(test.app, test.clock)
		handled, err := fnUnderTest(context.Background(), resp, test.input)
		if err != nil {
			t.Fatalf("(%q) %s", test.name, err)
		} else if handled != test.expectHandled {
			t.Fatalf("(%q) expected %v, got %v", test.name, test.expectHandled, handled)
		} else if test.expectedCode != 0 {
			if resp.Code != test.expectedCode {
				t.Fatalf("(%q) expected %d, got %d", test.name, test.expectedCode, resp.Code)
			}
		} else if test.expectedObjFn != nil {
			if err := VocabEquals(resp.Body, test.expectedObjFn()); err != nil {
				t.Fatalf("(%q) unexpected object: %s", test.name, err)
			}
		}
	}
}

func TestServeActivityPubObjectWithVerificationMethod(t *testing.T) {
	tests := []struct {
		name          string
		app           *MockApplication
		clock         *MockClock
		verifier      *MockSocialAPIVerifier
		input         *http.Request
		expectedCode  int
		expectedObjFn func() vocab.Serializer
		expectHandled bool
	}{
		{
			name: "unsigned request",
			app: &MockApplication{
				t: t,
				get: func(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
					if rw != Read {
						t.Fatalf("expected RWType of %d, got %d", Read, rw)
					} else if s := id.String(); s != noteURIString {
						t.Fatalf("expected %s, got %s", noteURIString, s)
					}
					testNote = &vocab.Note{}
					testNote.SetId(noteIRI)
					testNote.AppendNameString(noteName)
					testNote.AppendContentString("This is a simple note")
					return testNote, nil
				},
				owns: func(c context.Context, id *url.URL) bool {
					if s := id.String(); s != noteURIString {
						t.Fatalf("expected %s, got %s", noteURIString, s)
					}
					return true
				},
			},
			clock:        &MockClock{now},
			input:        ActivityPubRequest(httptest.NewRequest("GET", noteURIString, nil)),
			expectedCode: http.StatusOK,
			expectedObjFn: func() vocab.Serializer {
				testNote = &vocab.Note{}
				testNote.SetId(noteIRI)
				testNote.AppendNameString(noteName)
				testNote.AppendContentString("This is a simple note")
				return testNote
			},
			expectHandled: true,
		},
		{
			name:  "http signature request",
			input: Sign(ActivityPubRequest(httptest.NewRequest("GET", noteURIString, nil))),
			app: &MockApplication{
				t: t,
				getPublicKey: func(c context.Context, publicKeyId string) (crypto.PublicKey, httpsig.Algorithm, *url.URL, error) {
					if publicKeyId != testPublicKeyId {
						t.Fatalf("expected %s, got %s", testPublicKeyId, publicKeyId)
					}
					return testPrivateKey.Public(), httpsig.RSA_SHA256, samIRI, nil
				},
				getAsVerifiedUser: func(c context.Context, id, user *url.URL, rw RWType) (PubObject, error) {
					if rw != Read {
						t.Fatalf("expected RWType of %d, got %d", Read, rw)
					} else if s := id.String(); s != noteURIString {
						t.Fatalf("expected %s, got %s", noteURIString, s)
					} else if u := user.String(); u != samIRIString {
						t.Fatalf("expected %s, got %s", samIRIString, u)
					}
					testNote = &vocab.Note{}
					testNote.SetId(noteIRI)
					testNote.AppendNameString(noteName)
					testNote.AppendContentString("This is a simple note")
					return testNote, nil
				},
				owns: func(c context.Context, id *url.URL) bool {
					if s := id.String(); s != noteURIString {
						t.Fatalf("expected %s, got %s", noteURIString, s)
					}
					return true
				},
			},
			clock:        &MockClock{now},
			expectedCode: http.StatusOK,
			expectedObjFn: func() vocab.Serializer {
				testNote = &vocab.Note{}
				testNote.SetId(noteIRI)
				testNote.AppendNameString(noteName)
				testNote.AppendContentString("This is a simple note")
				return testNote
			},
			expectHandled: true,
		},
		{
			name:  "not owned",
			input: ActivityPubRequest(httptest.NewRequest("GET", noteURIString, nil)),
			app: &MockApplication{
				t: t,
				owns: func(c context.Context, id *url.URL) bool {
					if s := id.String(); s != noteURIString {
						t.Fatalf("expected %s, got %s", noteURIString, s)
					}
					return false
				},
			},
			expectedCode:  http.StatusNotFound,
			expectHandled: true,
		},
		{
			name:          "not activitypub get",
			input:         httptest.NewRequest("GET", noteURIString, nil),
			expectHandled: false,
		},
		{
			name:  "bad http signature",
			input: BadSignature(ActivityPubRequest(httptest.NewRequest("GET", noteURIString, nil))),
			app: &MockApplication{
				t: t,
				getPublicKey: func(c context.Context, publicKeyId string) (crypto.PublicKey, httpsig.Algorithm, *url.URL, error) {
					if publicKeyId != testPublicKeyId {
						t.Fatalf("expected %s, got %s", testPublicKeyId, publicKeyId)
					}
					return testPrivateKey.Public(), httpsig.RSA_SHA256, samIRI, nil
				},
				owns: func(c context.Context, id *url.URL) bool {
					if s := id.String(); s != noteURIString {
						t.Fatalf("expected %s, got %s", noteURIString, s)
					}
					return true
				},
			},
			expectedCode:  http.StatusForbidden,
			expectHandled: true,
		},
		{
			name: "unsigned request passes verifier",
			app: &MockApplication{
				t: t,
				getAsVerifiedUser: func(c context.Context, id, user *url.URL, rw RWType) (PubObject, error) {
					if rw != Read {
						t.Fatalf("expected RWType of %d, got %d", Read, rw)
					} else if s := id.String(); s != noteURIString {
						t.Fatalf("expected %s, got %s", noteURIString, s)
					} else if u := user.String(); u != samIRIString {
						t.Fatalf("expected %s, got %s", samIRIString, u)
					}
					testNote = &vocab.Note{}
					testNote.SetId(noteIRI)
					testNote.AppendNameString(noteName)
					testNote.AppendContentString("This is a simple note")
					return testNote, nil
				},
				owns: func(c context.Context, id *url.URL) bool {
					if s := id.String(); s != noteURIString {
						t.Fatalf("expected %s, got %s", noteURIString, s)
					}
					return true
				},
			},
			clock: &MockClock{now},
			verifier: &MockSocialAPIVerifier{
				t: t,
				verify: func(r *http.Request) (*url.URL, bool, bool, error) {
					return samIRI, true, true, nil
				},
			},
			input:        ActivityPubRequest(httptest.NewRequest("GET", noteURIString, nil)),
			expectedCode: http.StatusOK,
			expectedObjFn: func() vocab.Serializer {
				testNote = &vocab.Note{}
				testNote.SetId(noteIRI)
				testNote.AppendNameString(noteName)
				testNote.AppendContentString("This is a simple note")
				return testNote
			},
			expectHandled: true,
		},
		{
			name:  "http signature request passes verifier",
			input: Sign(ActivityPubRequest(httptest.NewRequest("GET", noteURIString, nil))),
			app: &MockApplication{
				t: t,
				getAsVerifiedUser: func(c context.Context, id, user *url.URL, rw RWType) (PubObject, error) {
					if rw != Read {
						t.Fatalf("expected RWType of %d, got %d", Read, rw)
					} else if s := id.String(); s != noteURIString {
						t.Fatalf("expected %s, got %s", noteURIString, s)
					} else if u := user.String(); u != samIRIString {
						t.Fatalf("expected %s, got %s", samIRIString, u)
					}
					testNote = &vocab.Note{}
					testNote.SetId(noteIRI)
					testNote.AppendNameString(noteName)
					testNote.AppendContentString("This is a simple note")
					return testNote, nil
				},
				owns: func(c context.Context, id *url.URL) bool {
					if s := id.String(); s != noteURIString {
						t.Fatalf("expected %s, got %s", noteURIString, s)
					}
					return true
				},
			},
			clock: &MockClock{now},
			verifier: &MockSocialAPIVerifier{
				t: t,
				verify: func(r *http.Request) (*url.URL, bool, bool, error) {
					return samIRI, true, true, nil
				},
			},
			expectedCode: http.StatusOK,
			expectedObjFn: func() vocab.Serializer {
				testNote = &vocab.Note{}
				testNote.SetId(noteIRI)
				testNote.AppendNameString(noteName)
				testNote.AppendContentString("This is a simple note")
				return testNote
			},
			expectHandled: true,
		},
		{
			name: "verifier authed unauthz",
			app: &MockApplication{
				t: t,
				owns: func(c context.Context, id *url.URL) bool {
					if s := id.String(); s != noteURIString {
						t.Fatalf("expected %s, got %s", noteURIString, s)
					}
					return true
				},
			},
			clock: &MockClock{now},
			verifier: &MockSocialAPIVerifier{
				t: t,
				verify: func(r *http.Request) (*url.URL, bool, bool, error) {
					return samIRI, true, false, nil
				},
			},
			input:         ActivityPubRequest(httptest.NewRequest("GET", noteURIString, nil)),
			expectedCode:  http.StatusForbidden,
			expectHandled: true,
		},
		{
			name: "verifier unauthed unauthz",
			app: &MockApplication{
				t: t,
				owns: func(c context.Context, id *url.URL) bool {
					if s := id.String(); s != noteURIString {
						t.Fatalf("expected %s, got %s", noteURIString, s)
					}
					return true
				},
			},
			clock: &MockClock{now},
			verifier: &MockSocialAPIVerifier{
				t: t,
				verify: func(r *http.Request) (*url.URL, bool, bool, error) {
					return nil, false, false, nil
				},
			},
			input:         ActivityPubRequest(httptest.NewRequest("GET", noteURIString, nil)),
			expectedCode:  http.StatusBadRequest,
			expectHandled: true,
		},
		{
			name: "verifier unauthed authz unsigned fails",
			app: &MockApplication{
				t: t,
				owns: func(c context.Context, id *url.URL) bool {
					if s := id.String(); s != noteURIString {
						t.Fatalf("expected %s, got %s", noteURIString, s)
					}
					return true
				},
			},
			clock: &MockClock{now},
			verifier: &MockSocialAPIVerifier{
				t: t,
				verify: func(r *http.Request) (*url.URL, bool, bool, error) {
					return nil, false, true, nil
				},
			},
			input:         ActivityPubRequest(httptest.NewRequest("GET", noteURIString, nil)),
			expectedCode:  http.StatusBadRequest,
			expectHandled: true,
		},
		{
			name:  "verifier unauthed authz signed success",
			input: Sign(ActivityPubRequest(httptest.NewRequest("GET", noteURIString, nil))),
			app: &MockApplication{
				t: t,
				getPublicKey: func(c context.Context, publicKeyId string) (crypto.PublicKey, httpsig.Algorithm, *url.URL, error) {
					if publicKeyId != testPublicKeyId {
						t.Fatalf("expected %s, got %s", testPublicKeyId, publicKeyId)
					}
					return testPrivateKey.Public(), httpsig.RSA_SHA256, samIRI, nil
				},
				getAsVerifiedUser: func(c context.Context, id, user *url.URL, rw RWType) (PubObject, error) {
					if rw != Read {
						t.Fatalf("expected RWType of %d, got %d", Read, rw)
					} else if s := id.String(); s != noteURIString {
						t.Fatalf("expected %s, got %s", noteURIString, s)
					} else if u := user.String(); u != samIRIString {
						t.Fatalf("expected %s, got %s", samIRIString, u)
					}
					testNote = &vocab.Note{}
					testNote.SetId(noteIRI)
					testNote.AppendNameString(noteName)
					testNote.AppendContentString("This is a simple note")
					return testNote, nil
				},
				owns: func(c context.Context, id *url.URL) bool {
					if s := id.String(); s != noteURIString {
						t.Fatalf("expected %s, got %s", noteURIString, s)
					}
					return true
				},
			},
			clock: &MockClock{now},
			verifier: &MockSocialAPIVerifier{
				t: t,
				verify: func(r *http.Request) (*url.URL, bool, bool, error) {
					return nil, false, true, nil
				},
			},
			expectedCode: http.StatusOK,
			expectedObjFn: func() vocab.Serializer {
				testNote = &vocab.Note{}
				testNote.SetId(noteIRI)
				testNote.AppendNameString(noteName)
				testNote.AppendContentString("This is a simple note")
				return testNote
			},
			expectHandled: true,
		},
		{
			name: "verifier unauthed authz unsigned fails with bad impl returning user",
			app: &MockApplication{
				t: t,
				owns: func(c context.Context, id *url.URL) bool {
					if s := id.String(); s != noteURIString {
						t.Fatalf("expected %s, got %s", noteURIString, s)
					}
					return true
				},
			},
			clock: &MockClock{now},
			verifier: &MockSocialAPIVerifier{
				t: t,
				verify: func(r *http.Request) (*url.URL, bool, bool, error) {
					return samIRI, false, true, nil
				},
			},
			input:         ActivityPubRequest(httptest.NewRequest("GET", noteURIString, nil)),
			expectedCode:  http.StatusBadRequest,
			expectHandled: true,
		},
	}
	for _, test := range tests {
		t.Logf("Running table test case %q", test.name)
		resp := httptest.NewRecorder()
		var fnUnderTest HandlerFunc
		if test.verifier != nil {
			verifierFn := func(c context.Context) SocialAPIVerifier {
				return test.verifier
			}
			fnUnderTest = ServeActivityPubObjectWithVerificationMethod(test.app, test.clock, verifierFn)
		} else {
			fnUnderTest = ServeActivityPubObjectWithVerificationMethod(test.app, test.clock, nil)
		}
		handled, err := fnUnderTest(context.Background(), resp, test.input)
		if err != nil {
			t.Fatalf("(%q) %s", test.name, err)
		} else if handled != test.expectHandled {
			t.Fatalf("(%q) expected %v, got %v", test.name, test.expectHandled, handled)
		} else if test.expectedCode != 0 {
			if resp.Code != test.expectedCode {
				t.Fatalf("(%q) expected %d, got %d", test.name, test.expectedCode, resp.Code)
			}
		} else if test.expectedObjFn != nil {
			if err := VocabEquals(resp.Body, test.expectedObjFn()); err != nil {
				t.Fatalf("(%q) unexpected object: %s", test.name, err)
			}
		}
	}
}

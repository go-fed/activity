package pub

import (
	"bytes"
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"github.com/go-fed/activity/streams"
	"github.com/go-fed/activity/vocab"
	"github.com/go-fed/httpsig"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

const (
	iriString                 = "https://example.com/something"
	noteURIString             = "https://example.com/note/123"
	updateURIString           = "https://example.com/note/update/123"
	noteActivityURIString     = "https://example.com/activity/987"
	testAgent                 = "test agent string"
	testInboxURI              = "https://example.com/sally/inbox"
	testOutboxURI             = "https://example.com/sally/outbox"
	testNewIRIString          = "https://example.com/test/new/iri/1"
	testNewIRIString2         = "https://example.com/test/new/iri/2"
	sallyIRIString            = "https://example.com/sally"
	sallyFollowingIRIString   = "https://example.com/sally/following"
	samIRIString              = "https://example.com/sam"
	samIRIInboxString         = "https://example.com/sam/inbox"
	samIRIFollowersString     = "https://example.com/sam/followers"
	sallyIRIInboxString       = "https://example.com/sally/inbox"
	noteName                  = "A Note"
	otherOriginIRIString      = "https://foo.net/activity/112358"
	otherOriginActorIRIString = "https://foo.net/peyton"

	testPublicKeyId = "publicKeyId"
)

var (
	iri                              *url.URL
	noteIRI                          *url.URL
	noteActivityIRI                  *url.URL
	updateActivityIRI                *url.URL
	testNewIRI                       *url.URL
	testNewIRI2                      *url.URL
	sallyIRI                         *url.URL
	sallyIRIInbox                    *url.URL
	sallyFollowingIRI                *url.URL
	sallyActor                       *vocab.Person
	sallyActorJSON                   []byte
	samIRI                           *url.URL
	samIRIInbox                      *url.URL
	samIRIFollowers                  *url.URL
	otherOriginIRI                   *url.URL
	otherOriginActorIRI              *url.URL
	samActor                         *vocab.Person
	samActorJSON                     []byte
	testNote                         *vocab.Note
	testSingleOrderedCollection      *vocab.OrderedCollection
	testCreateNote                   *vocab.Create
	testUpdateNote                   *vocab.Update
	testDeleteNote                   *vocab.Delete
	testTombstoneNote                *vocab.Tombstone
	testFollow                       *vocab.Follow
	testAcceptNote                   *vocab.Accept
	testAcceptFollow                 *vocab.Accept
	testRejectFollow                 *vocab.Reject
	testAddNote                      *vocab.Add
	testRemoveNote                   *vocab.Remove
	testLikeNote                     *vocab.Like
	testUndoLike                     *vocab.Undo
	testBlock                        *vocab.Block
	testClientExpectedNote           *vocab.Note
	testClientExpectedCreateNote     *vocab.Create
	testDeleteSubFields              string
	testDeleteFields                 string
	testDeleteFieldsDifferentObjects string
	testClientUpdateNote             *vocab.Update
	testClientExpectedUpdateNote     *vocab.Update
	testClientExpectedDeleteNote     *vocab.Delete
	testClientExpectedFollow         *vocab.Follow
	testClientExpectedAcceptFollow   *vocab.Accept
	testClientExpectedRejectFollow   *vocab.Reject
	testClientExpectedAdd            *vocab.Add
	testClientExpectedRemove         *vocab.Remove
	testClientExpectedLike           *vocab.Like
	testClientExpectedUndo           *vocab.Undo
	testClientExpectedBlock          *vocab.Block

	testPrivateKey      *rsa.PrivateKey
	testOtherPrivateKey *rsa.PrivateKey
)

func init() {
	var err error
	iri, err = url.Parse(iriString)
	if err != nil {
		panic(err)
	}
	noteIRI, err = url.Parse(noteURIString)
	if err != nil {
		panic(err)
	}
	noteActivityIRI, err = url.Parse(noteActivityURIString)
	if err != nil {
		panic(err)
	}
	updateActivityIRI, err = url.Parse(updateURIString)
	if err != nil {
		panic(err)
	}
	testNewIRI, err = url.Parse(testNewIRIString)
	if err != nil {
		panic(err)
	}
	testNewIRI2, err = url.Parse(testNewIRIString2)
	if err != nil {
		panic(err)
	}
	sallyIRI, err = url.Parse(sallyIRIString)
	if err != nil {
		panic(err)
	}
	sallyIRIInbox, err = url.Parse(sallyIRIInboxString)
	if err != nil {
		panic(err)
	}
	sallyFollowingIRI, err = url.Parse(sallyFollowingIRIString)
	if err != nil {
		panic(err)
	}
	samIRI, err = url.Parse(samIRIString)
	if err != nil {
		panic(err)
	}
	samIRIInbox, err = url.Parse(samIRIInboxString)
	if err != nil {
		panic(err)
	}
	samIRIFollowers, err = url.Parse(samIRIFollowersString)
	if err != nil {
		panic(err)
	}
	otherOriginIRI, err = url.Parse(otherOriginIRIString)
	if err != nil {
		panic(err)
	}
	otherOriginActorIRI, err = url.Parse(otherOriginActorIRIString)
	if err != nil {
		panic(err)
	}
	samActor = &vocab.Person{}
	samActor.SetInboxAnyURI(samIRIInbox)
	samActor.SetId(samIRI)
	samActor.AppendNameString("Sam")
	m, err := samActor.Serialize()
	if err != nil {
		panic(err)
	}
	samActorJSON, err = json.Marshal(m)
	if err != nil {
		panic(err)
	}
	sallyInbox, err := url.Parse(testInboxURI)
	if err != nil {
		panic(err)
	}
	sallyOutbox, err := url.Parse(testOutboxURI)
	if err != nil {
		panic(err)
	}
	sallyActor = &vocab.Person{}
	sallyActor.AppendNameString("Sally")
	sallyActor.SetId(sallyIRI)
	sallyActor.SetInboxAnyURI(sallyInbox)
	sallyActor.SetOutboxAnyURI(sallyOutbox)
	m, err = sallyActor.Serialize()
	if err != nil {
		panic(err)
	}
	sallyActorJSON, err = json.Marshal(m)
	if err != nil {
		panic(err)
	}
	testNote = &vocab.Note{}
	testNote.SetId(noteIRI)
	testNote.AppendNameString(noteName)
	testNote.AppendContentString("This is a simple note")
	testSingleOrderedCollection = &vocab.OrderedCollection{}
	testSingleOrderedCollection.AppendItemsObject(testNote)
	testCreateNote = &vocab.Create{}
	testCreateNote.SetId(noteActivityIRI)
	testCreateNote.AppendSummaryString("Sally created a note")
	testCreateNote.AppendActorObject(sallyActor)
	testCreateNote.AppendObject(testNote)
	testCreateNote.AppendToObject(samActor)
	testUpdateNote = &vocab.Update{}
	testUpdateNote.SetId(updateActivityIRI)
	testUpdateNote.AppendSummaryString("Sally updated a note")
	testUpdateNote.AppendActorObject(sallyActor)
	testUpdateNote.AppendObject(testNote)
	testUpdateNote.AppendToObject(samActor)
	testDeleteNote = &vocab.Delete{}
	testDeleteNote.SetId(noteActivityIRI)
	testDeleteNote.AppendActorObject(sallyActor)
	testDeleteNote.AppendObject(testNote)
	testDeleteNote.AppendToObject(samActor)
	testTombstoneNote = &vocab.Tombstone{}
	testTombstoneNote.SetId(noteIRI)
	testTombstoneNote.AppendFormerTypeString("Note")
	testTombstoneNote.SetDeleted(now)
	testFollow = &vocab.Follow{}
	testFollow.SetId(noteActivityIRI)
	testFollow.AppendActorObject(sallyActor)
	testFollow.AppendObject(samActor)
	testFollow.AppendToObject(samActor)
	testAcceptNote = &vocab.Accept{}
	testAcceptNote.SetId(noteActivityIRI)
	testAcceptNote.AppendActorObject(sallyActor)
	testAcceptNote.AppendObject(&vocab.Offer{})
	testAcceptNote.AppendToObject(samActor)
	testAcceptFollow = &vocab.Accept{}
	testAcceptFollow.SetId(noteActivityIRI)
	testAcceptFollow.AppendActorObject(samActor)
	testAcceptFollow.AppendObject(testFollow)
	testAcceptFollow.AppendToObject(sallyActor)
	testRejectFollow = &vocab.Reject{}
	testRejectFollow.SetId(noteActivityIRI)
	testRejectFollow.AppendActorObject(samActor)
	testRejectFollow.AppendObject(testFollow)
	testRejectFollow.AppendToObject(sallyActor)
	testAddNote = &vocab.Add{}
	testAddNote.SetId(noteActivityIRI)
	testAddNote.AppendActorObject(sallyActor)
	testAddNote.AppendObject(testNote)
	testAddNote.AppendTargetIRI(iri)
	testAddNote.AppendToObject(samActor)
	testRemoveNote = &vocab.Remove{}
	testRemoveNote.SetId(noteActivityIRI)
	testRemoveNote.AppendActorObject(sallyActor)
	testRemoveNote.AppendObject(testNote)
	testRemoveNote.AppendTargetIRI(iri)
	testRemoveNote.AppendToObject(samActor)
	testLikeNote = &vocab.Like{}
	testLikeNote.SetId(noteActivityIRI)
	testLikeNote.AppendActorObject(sallyActor)
	testLikeNote.AppendObject(testNote)
	testLikeNote.AppendToObject(samActor)
	testUndoLike = &vocab.Undo{}
	testUndoLike.SetId(noteActivityIRI)
	testUndoLike.AppendActorObject(sallyActor)
	testUndoLike.AppendObject(testLikeNote)
	testUndoLike.AppendToObject(samActor)
	testBlock = &vocab.Block{}
	testBlock.SetId(noteActivityIRI)
	testBlock.AppendActorObject(sallyActor)
	testBlock.AppendObject(samActor)

	testClientExpectedNote = &vocab.Note{}
	testClientExpectedNote.SetId(testNewIRI2)
	testClientExpectedNote.AppendNameString(noteName)
	testClientExpectedNote.AppendContentString("This is a simple note")
	testClientExpectedNote.AppendAttributedToIRI(sallyIRI)
	testClientExpectedNote.AppendToIRI(samIRI)
	testClientExpectedCreateNote = &vocab.Create{}
	testClientExpectedCreateNote.SetId(testNewIRI)
	testClientExpectedCreateNote.AppendSummaryString("Sally created a note")
	testClientExpectedCreateNote.AppendActorObject(sallyActor)
	testClientExpectedCreateNote.AppendObject(testClientExpectedNote)
	testClientExpectedCreateNote.AppendToObject(samActor)
	testDeleteSubFields = `
        {
          "@context": "https://www.w3.org/ns/activitystreams",
          "summary": "Sally updated her note",
          "type": "Update",
          "actor": "https://example.com/sally",
	  "id": "https://example.com/test/new/iri",
          "object": {
            "id": "https://example.com/note/123",
	    "type": "Note",
            "to": {
              "id": "https://example.com/sam",
              "inbox": "https://example.com/sam/inbox",
	      "type": "Person",
	      "name": null
            }
	  }
        }
	`
	testDeleteFields = `
        {
          "@context": "https://www.w3.org/ns/activitystreams",
          "summary": "Sally updated her note",
          "type": "Update",
          "actor": "https://example.com/sally",
	  "id": "https://example.com/test/new/iri",
          "object": {
            "id": "https://example.com/note/123",
	    "type": "Note",
            "to": null
	  }
        }
	`
	testDeleteFieldsDifferentObjects = `
	{
          "@context": "https://www.w3.org/ns/activitystreams",
          "summary": "Sally updated her notes",
          "type": "Update",
          "actor": "https://example.com/sally",
	  "id": "https://example.com/test/new/iri",
          "object": [
            {
              "id": "https://example.com/note/123",
	      "type": "Note",
              "to": {
                "id": "https://example.com/sam",
                "inbox": "https://example.com/sam/inbox",
	        "type": "Person",
	        "name": null
              }
            },
	    {
              "id": "https://example.com/note/update/123",
	      "type": "Note",
              "to": {
                "id": "https://example.com/sam",
                "inbox": "https://example.com/sam/inbox",
	        "type": "Person",
	        "name": null
              }
            }
	  ]
	}
	`
	sammActor := &vocab.Person{}
	sammActor.SetInboxAnyURI(samIRIInbox)
	sammActor.SetId(samIRI)
	sammActor.AppendNameString("Samm")
	testUpdateNote := &vocab.Note{}
	testUpdateNote.SetId(noteIRI)
	testUpdateNote.AppendNameString(noteName)
	testUpdateNote.AppendContentString("This is a simple note")
	testUpdateNote.AppendToObject(sammActor)
	testClientUpdateNote = &vocab.Update{}
	testClientUpdateNote.SetId(updateActivityIRI)
	testClientUpdateNote.AppendSummaryString("Sally updated a note")
	testClientUpdateNote.AppendActorObject(sallyActor)
	testClientUpdateNote.AppendObject(testUpdateNote)
	testClientUpdateNote.AppendToObject(samActor)
	testClientExpectedUpdateNote = &vocab.Update{}
	testClientExpectedUpdateNote.SetId(testNewIRI)
	testClientExpectedUpdateNote.AppendSummaryString("Sally updated a note")
	testClientExpectedUpdateNote.AppendActorObject(sallyActor)
	testClientExpectedUpdateNote.AppendObject(testNote)
	testClientExpectedUpdateNote.AppendToObject(samActor)
	testClientExpectedDeleteNote = &vocab.Delete{}
	testClientExpectedDeleteNote.SetId(testNewIRI)
	testClientExpectedDeleteNote.AppendActorObject(sallyActor)
	testClientExpectedDeleteNote.AppendObject(testNote)
	testClientExpectedDeleteNote.AppendToObject(samActor)
	testClientExpectedFollow = &vocab.Follow{}
	testClientExpectedFollow.SetId(testNewIRI)
	testClientExpectedFollow.AppendActorObject(sallyActor)
	testClientExpectedFollow.AppendObject(samActor)
	testClientExpectedFollow.AppendToObject(samActor)
	testClientExpectedAcceptFollow = &vocab.Accept{}
	testClientExpectedAcceptFollow.SetId(testNewIRI)
	testClientExpectedAcceptFollow.AppendActorObject(samActor)
	testClientExpectedAcceptFollow.AppendObject(testFollow)
	testClientExpectedAcceptFollow.AppendToObject(sallyActor)
	testClientExpectedRejectFollow = &vocab.Reject{}
	testClientExpectedRejectFollow.SetId(testNewIRI)
	testClientExpectedRejectFollow.AppendActorObject(samActor)
	testClientExpectedRejectFollow.AppendObject(testFollow)
	testClientExpectedRejectFollow.AppendToObject(sallyActor)
	testClientExpectedAdd = &vocab.Add{}
	testClientExpectedAdd.SetId(testNewIRI)
	testClientExpectedAdd.AppendActorObject(sallyActor)
	testClientExpectedAdd.AppendObject(testNote)
	testClientExpectedAdd.AppendTargetIRI(iri)
	testClientExpectedAdd.AppendToObject(samActor)
	testClientExpectedRemove = &vocab.Remove{}
	testClientExpectedRemove.SetId(testNewIRI)
	testClientExpectedRemove.AppendActorObject(sallyActor)
	testClientExpectedRemove.AppendObject(testNote)
	testClientExpectedRemove.AppendTargetIRI(iri)
	testClientExpectedRemove.AppendToObject(samActor)
	testClientExpectedLike = &vocab.Like{}
	testClientExpectedLike.SetId(testNewIRI)
	testClientExpectedLike.AppendActorObject(sallyActor)
	testClientExpectedLike.AppendObject(testNote)
	testClientExpectedLike.AppendToObject(samActor)
	testClientExpectedUndo = &vocab.Undo{}
	testClientExpectedUndo.SetId(testNewIRI)
	testClientExpectedUndo.AppendActorObject(sallyActor)
	testClientExpectedUndo.AppendObject(testLikeNote)
	testClientExpectedUndo.AppendToObject(samActor)
	testClientExpectedBlock = &vocab.Block{}
	testClientExpectedBlock.SetId(testNewIRI)
	testClientExpectedBlock.AppendActorObject(sallyActor)
	testClientExpectedBlock.AppendObject(samActor)

	testPrivateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	testOtherPrivateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
}

func Must(l *time.Location, e error) *time.Location {
	if e != nil {
		panic(e)
	}
	return l
}

func MustSerialize(s vocab.Serializer) []byte {
	m, err := s.Serialize()
	if err != nil {
		panic(err)
	}
	addJSONLDContext(m)
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return b
}

func BadSignature(r *http.Request) *http.Request {
	s, _, err := httpsig.NewSigner([]httpsig.Algorithm{httpsig.RSA_SHA256}, nil, httpsig.Signature)
	if err != nil {
		panic(err)
	}
	err = s.SignRequest(testOtherPrivateKey, testPublicKeyId, r)
	if err != nil {
		panic(err)
	}
	return r
}

func Sign(r *http.Request) *http.Request {
	s, _, err := httpsig.NewSigner([]httpsig.Algorithm{httpsig.RSA_SHA256}, nil, httpsig.Signature)
	if err != nil {
		panic(err)
	}
	err = s.SignRequest(testPrivateKey, testPublicKeyId, r)
	if err != nil {
		panic(err)
	}
	return r
}

func ActivityPubRequest(r *http.Request) *http.Request {
	if r.Method == "POST" {
		existing, ok := r.Header["Content-Type"]
		if ok {
			r.Header["Content-Type"] = append(existing, "application/ld+json; profile=\"https://www.w3.org/ns/activitystreams\"")
		} else {
			r.Header["Content-Type"] = []string{"application/ld+json; profile=\"https://www.w3.org/ns/activitystreams\""}
		}
	} else {
		existing, ok := r.Header["Accept"]
		if ok {
			r.Header["Accept"] = append(existing, "application/ld+json; profile=\"https://www.w3.org/ns/activitystreams\"")
		} else {
			r.Header["Accept"] = []string{"application/ld+json; profile=\"https://www.w3.org/ns/activitystreams\""}
		}
	}
	r.Header["Date"] = []string{now.UTC().Format("Mon, 02 Jan 2006 15:04:05") + " GMT"}
	return r
}

func PubObjectEquals(p PubObject, s vocab.Serializer) error {
	ps, ok := p.(vocab.Serializer)
	if !ok {
		return fmt.Errorf("PubObject is not Serializer")
	}
	m, err := ps.Serialize()
	if err != nil {
		return err
	}
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return VocabEqualsContext(bytes.NewBuffer(b), s, false)
}

func VocabSerializerEquals(i, j vocab.Serializer) error {
	m, err := i.Serialize()
	if err != nil {
		return err
	}
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return VocabEqualsContext(bytes.NewBuffer(b), j, false)
}

func VocabEquals(b *bytes.Buffer, s vocab.Serializer) error {
	return VocabEqualsContext(b, s, true)
}

func VocabEqualsContext(b *bytes.Buffer, s vocab.Serializer, requireContext bool) error {
	m, err := s.Serialize()
	if err != nil {
		return err
	}
	if requireContext {
		m["@context"] = "https://www.w3.org/ns/activitystreams"
	}
	expected, err := json.Marshal(m)
	if err != nil {
		return err
	}
	actual := b.Bytes()
	if len(actual) != len(expected) {
		return fmt.Errorf("expected len %d, actual len %d:\nexpected value %s\nactual value %s", len(expected), len(actual), expected, actual)
	}
	var diffs []string
	for i := range actual {
		if actual[i] != expected[i] {
			diffs = append(diffs, fmt.Sprintf("at %d expected %d but got %d", i, expected[i], actual[i]))
		}
	}
	if len(diffs) == 0 {
		return nil
	}
	return fmt.Errorf(strings.Join(diffs, "; "))
}

var (
	now               = time.Date(2000, 2, 3, 4, 5, 6, 7, Must(time.LoadLocation("America/New_York")))
	testPublishedTime = time.Date(2001, 2, 3, 4, 5, 6, 7, Must(time.LoadLocation("America/New_York")))
	testUpdateTime    = time.Date(2002, 2, 3, 4, 5, 6, 7, Must(time.LoadLocation("America/New_York")))
)

var _ Clock = &MockClock{}

type MockClock struct {
	now time.Time
}

func (m *MockClock) Now() time.Time {
	return m.now
}

var _ Application = &MockApplication{}

type MockApplication struct {
	t                 *testing.T
	owns              func(c context.Context, id *url.URL) bool
	get               func(c context.Context, id *url.URL, rw RWType) (PubObject, error)
	getAsVerifiedUser func(c context.Context, id, authdUser *url.URL, rw RWType) (PubObject, error)
	has               func(c context.Context, id *url.URL) (bool, error)
	set               func(c context.Context, o PubObject) error
	getInbox          func(c context.Context, r *http.Request, rw RWType) (vocab.OrderedCollectionType, error)
	getOutbox         func(c context.Context, r *http.Request, rw RWType) (vocab.OrderedCollectionType, error)
	newId             func(c context.Context, t Typer) *url.URL
	getPublicKey      func(c context.Context, publicKeyId string) (crypto.PublicKey, httpsig.Algorithm, *url.URL, error)
	canAdd            func(c context.Context, o vocab.ObjectType, t vocab.ObjectType) bool
	canRemove         func(c context.Context, o vocab.ObjectType, t vocab.ObjectType) bool
}

func (m *MockApplication) Owns(c context.Context, id *url.URL) bool {
	if m.owns == nil {
		m.t.Fatal("unexpected call to MockApplication Owns")
	}
	return m.owns(c, id)
}

func (m *MockApplication) Get(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
	if m.get == nil {
		m.t.Fatal("unexpected call to MockApplication Get")
	}
	return m.get(c, id, rw)
}

func (m *MockApplication) GetAsVerifiedUser(c context.Context, id, authdUser *url.URL, rw RWType) (PubObject, error) {
	if m.getAsVerifiedUser == nil {
		m.t.Fatal("unexpected call to MockApplication GetAsVerifiedUser")
	}
	return m.getAsVerifiedUser(c, id, authdUser, rw)
}

func (m *MockApplication) Has(c context.Context, id *url.URL) (bool, error) {
	if m.has == nil {
		m.t.Fatal("unexpected call to MockApplication Has")
	}
	return m.has(c, id)
}

func (m *MockApplication) Set(c context.Context, o PubObject) error {
	if m.set == nil {
		m.t.Fatal("unexpected call to MockApplication Set")
	}
	return m.set(c, o)
}

func (m *MockApplication) GetInbox(c context.Context, r *http.Request, rw RWType) (vocab.OrderedCollectionType, error) {
	if m.getInbox == nil {
		m.t.Fatal("unexpected call to MockApplication GetInbox")
	}
	return m.getInbox(c, r, rw)
}

func (m *MockApplication) GetOutbox(c context.Context, r *http.Request, rw RWType) (vocab.OrderedCollectionType, error) {
	if m.getOutbox == nil {
		m.t.Fatal("unexpected call to MockApplication GetOutbox")
	}
	return m.getOutbox(c, r, rw)
}

func (m *MockApplication) NewId(c context.Context, t Typer) *url.URL {
	if m.newId == nil {
		m.t.Fatal("unexpected call to MockApplication NewId")
	}
	return m.newId(c, t)
}

func (m *MockApplication) GetPublicKey(c context.Context, publicKeyId string) (crypto.PublicKey, httpsig.Algorithm, *url.URL, error) {
	if m.getPublicKey == nil {
		m.t.Fatal("unexpected call to MockApplication GetPublicKey")
	}
	return m.getPublicKey(c, publicKeyId)
}

func (m *MockApplication) CanAdd(c context.Context, o vocab.ObjectType, t vocab.ObjectType) bool {
	if m.canAdd == nil {
		m.t.Fatal("unexpected call to MockApplication CanAdd")
	}
	return m.canAdd(c, o, t)
}

func (m *MockApplication) CanRemove(c context.Context, o vocab.ObjectType, t vocab.ObjectType) bool {
	if m.canRemove == nil {
		m.t.Fatal("unexpected call to MockApplication CanRemove")
	}
	return m.canRemove(c, o, t)
}

var _ SocialApplication = &MockSocialApp{}

type MockSocialApp struct {
	*MockApplication
	t                     *testing.T
	actorIRI              func(c context.Context, r *http.Request) (*url.URL, error)
	getPublicKeyForOutbox func(c context.Context, publicKeyId string, boxIRI *url.URL) (crypto.PublicKey, httpsig.Algorithm, error)
	getSocialAPIVerifier  func(c context.Context) SocialAPIVerifier
}

func (m *MockSocialApp) ActorIRI(c context.Context, r *http.Request) (*url.URL, error) {
	if m.actorIRI == nil {
		m.t.Fatal("unexpected call to MockSocialApp ActorIRI")
	}
	return m.actorIRI(c, r)
}

func (m *MockSocialApp) GetPublicKeyForOutbox(c context.Context, publicKeyId string, boxIRI *url.URL) (crypto.PublicKey, httpsig.Algorithm, error) {
	if m.getPublicKeyForOutbox == nil {
		m.t.Fatal("unexpected call to MockSocialApp GetPublicKeyForOutbox")
	}
	return m.getPublicKeyForOutbox(c, publicKeyId, boxIRI)
}

func (m *MockSocialApp) GetSocialAPIVerifier(c context.Context) SocialAPIVerifier {
	if m.getSocialAPIVerifier == nil {
		m.t.Fatal("unexpected call to MockSocialApp GetSocialAPIVerifier")
	}
	return m.getSocialAPIVerifier(c)
}

var _ Callbacker = &MockCallbacker{}

type MockCallbacker struct {
	t      *testing.T
	create func(c context.Context, s *streams.Create) error
	update func(c context.Context, s *streams.Update) error
	delete func(c context.Context, s *streams.Delete) error
	add    func(c context.Context, s *streams.Add) error
	remove func(c context.Context, s *streams.Remove) error
	like   func(c context.Context, s *streams.Like) error
	block  func(c context.Context, s *streams.Block) error
	follow func(c context.Context, s *streams.Follow) error
	undo   func(c context.Context, s *streams.Undo) error
	accept func(c context.Context, s *streams.Accept) error
	reject func(c context.Context, s *streams.Reject) error
}

func (m *MockCallbacker) Create(c context.Context, s *streams.Create) error {
	if m.create == nil {
		m.t.Logf("unimplemented MockCallbacker Create called: %v %v", c, s)
		return nil
	} else {
		return m.create(c, s)
	}
}

func (m *MockCallbacker) Update(c context.Context, s *streams.Update) error {
	if m.update == nil {
		m.t.Logf("unimplemented MockCallbacker Update called: %v %v", c, s)
		return nil
	} else {
		return m.update(c, s)
	}
}

func (m *MockCallbacker) Delete(c context.Context, s *streams.Delete) error {
	if m.delete == nil {
		m.t.Logf("unimplemented MockCallbacker Delete called: %v %v", c, s)
		return nil
	} else {
		return m.delete(c, s)
	}
}

func (m *MockCallbacker) Add(c context.Context, s *streams.Add) error {
	if m.add == nil {
		m.t.Logf("unimplemented MockCallbacker Add called: %v %v", c, s)
		return nil
	} else {
		return m.add(c, s)
	}
}

func (m *MockCallbacker) Remove(c context.Context, s *streams.Remove) error {
	if m.remove == nil {
		m.t.Logf("unimplemented MockCallbacker Remove called: %v %v", c, s)
		return nil
	} else {
		return m.remove(c, s)
	}
}

func (m *MockCallbacker) Like(c context.Context, s *streams.Like) error {
	if m.like == nil {
		m.t.Logf("unimplemented MockCallbacker Like called: %v %v", c, s)
		return nil
	} else {
		return m.like(c, s)
	}
}

func (m *MockCallbacker) Block(c context.Context, s *streams.Block) error {
	if m.block == nil {
		m.t.Logf("unimplemented MockCallbacker Block called: %v %v", c, s)
		return nil
	} else {
		return m.block(c, s)
	}
}

func (m *MockCallbacker) Follow(c context.Context, s *streams.Follow) error {
	if m.follow == nil {
		m.t.Logf("unimplemented MockCallbacker Follow called: %v %v", c, s)
		return nil
	} else {
		return m.follow(c, s)
	}
}

func (m *MockCallbacker) Undo(c context.Context, s *streams.Undo) error {
	if m.undo == nil {
		m.t.Logf("unimplemented MockCallbacker Undo called: %v %v", c, s)
		return nil
	} else {
		return m.undo(c, s)
	}
}

func (m *MockCallbacker) Accept(c context.Context, s *streams.Accept) error {
	if m.accept == nil {
		m.t.Logf("unimplemented MockCallbacker Accept called: %v %v", c, s)
		return nil
	} else {
		return m.accept(c, s)
	}
}

func (m *MockCallbacker) Reject(c context.Context, s *streams.Reject) error {
	if m.reject == nil {
		m.t.Logf("unimplemented MockCallbacker Reject called: %v %v", c, s)
		return nil
	} else {
		return m.reject(c, s)
	}
}

var _ FederateApplication = &MockFederateApp{}

type MockFederateApp struct {
	*MockApplication
	t                *testing.T
	onFollow         func(c context.Context, s *streams.Follow) FollowResponse
	unblocked        func(c context.Context, actorIRIs []*url.URL) error
	filterForwarding func(c context.Context, activity vocab.ActivityType, iris []*url.URL) ([]*url.URL, error)
	newSigner        func() (httpsig.Signer, error)
	privateKey       func(boxIRI *url.URL) (crypto.PrivateKey, string, error)
}

func (m *MockFederateApp) OnFollow(c context.Context, s *streams.Follow) FollowResponse {
	if m.onFollow == nil {
		m.t.Fatal("unexpected call to MockFederateApp OnFollow")
	}
	return m.onFollow(c, s)
}

func (m *MockFederateApp) Unblocked(c context.Context, actorIRIs []*url.URL) error {
	if m.unblocked == nil {
		m.t.Fatal("unexpected call to MockFederateApp Unblocked")
	}
	return m.unblocked(c, actorIRIs)
}

func (m *MockFederateApp) FilterForwarding(c context.Context, activity vocab.ActivityType, iris []*url.URL) ([]*url.URL, error) {
	if m.filterForwarding == nil {
		m.t.Fatal("unexpected call to MockFederateApp FilterForwarding")
	}
	return m.filterForwarding(c, activity, iris)
}

func (m *MockFederateApp) NewSigner() (httpsig.Signer, error) {
	if m.newSigner == nil {
		m.t.Fatal("unexpected call to MockFederateApp NewSigner")
	}
	return m.newSigner()
}

func (m *MockFederateApp) PrivateKey(boxIRI *url.URL) (privKey crypto.PrivateKey, pubKeyId string, err error) {
	if m.privateKey == nil {
		m.t.Fatal("unexpected call to MockFederateApp PrivateKey")
	}
	return m.privateKey(boxIRI)
}

var _ SocialFederateApplication = &MockSocialFederateApp{}

type MockSocialFederateApp struct {
	*MockFederateApp
	*MockSocialApp
}

func (m *MockSocialFederateApp) Owns(c context.Context, id *url.URL) bool {
	return m.MockFederateApp.Owns(c, id)
}
func (m *MockSocialFederateApp) Get(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
	return m.MockFederateApp.Get(c, id, rw)
}
func (m *MockSocialFederateApp) GetAsVerifiedUser(c context.Context, id, authdUser *url.URL, rw RWType) (PubObject, error) {
	return m.MockFederateApp.GetAsVerifiedUser(c, id, authdUser, rw)
}
func (m *MockSocialFederateApp) Has(c context.Context, id *url.URL) (bool, error) {
	return m.MockFederateApp.Has(c, id)
}
func (m *MockSocialFederateApp) Set(c context.Context, o PubObject) error {
	return m.MockFederateApp.Set(c, o)
}
func (m *MockSocialFederateApp) GetInbox(c context.Context, r *http.Request, rw RWType) (vocab.OrderedCollectionType, error) {
	return m.MockFederateApp.GetInbox(c, r, rw)
}
func (m *MockSocialFederateApp) GetOutbox(c context.Context, r *http.Request, rw RWType) (vocab.OrderedCollectionType, error) {
	return m.MockFederateApp.GetOutbox(c, r, rw)
}
func (m *MockSocialFederateApp) NewId(c context.Context, t Typer) *url.URL {
	return m.MockFederateApp.NewId(c, t)
}
func (m *MockSocialFederateApp) GetPublicKey(c context.Context, publicKeyId string) (crypto.PublicKey, httpsig.Algorithm, *url.URL, error) {
	return m.MockFederateApp.GetPublicKey(c, publicKeyId)
}
func (m *MockSocialFederateApp) CanAdd(c context.Context, o vocab.ObjectType, t vocab.ObjectType) bool {
	return m.MockFederateApp.CanAdd(c, o, t)
}
func (m *MockSocialFederateApp) CanRemove(c context.Context, o vocab.ObjectType, t vocab.ObjectType) bool {
	return m.MockFederateApp.CanRemove(c, o, t)
}

var _ Deliverer = &MockDeliverer{}

type MockDeliverer struct {
	t  *testing.T
	do func(b []byte, to *url.URL, toDo func(b []byte, u *url.URL) error)
}

func (m *MockDeliverer) Do(b []byte, to *url.URL, toDo func(b []byte, u *url.URL) error) {
	if m.do == nil {
		m.t.Fatal("unexpected call to MockDeliverer Do")
	}
	m.do(b, to, toDo)
}

var _ HttpClient = &MockHttpClient{}

type MockHttpClient struct {
	t  *testing.T
	do func(req *http.Request) (*http.Response, error)
}

func (m *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	if m.do == nil {
		m.t.Fatal("unexpected call to MockHttpClient Do")
	}
	return m.do(req)
}

var _ SocialAPIVerifier = &MockSocialAPIVerifier{}

type MockSocialAPIVerifier struct {
	t               *testing.T
	verify          func(r *http.Request) (authenticatedUser *url.URL, authn, authz bool, err error)
	verifyForOutbox func(r *http.Request, outbox *url.URL) (authn, authz bool, err error)
}

func (m *MockSocialAPIVerifier) Verify(r *http.Request) (authenticatedUser *url.URL, authn, authz bool, err error) {
	if m.verify == nil {
		m.t.Fatal("unexpected call to MockSocialAPIVerifier Verify")
	}
	return m.verify(r)
}

func (m *MockSocialAPIVerifier) VerifyForOutbox(r *http.Request, outbox *url.URL) (authn, authz bool, err error) {
	if m.verifyForOutbox == nil {
		m.t.Fatal("unexpected call to MockSocialAPIVerifier VerifyForOutbox")
	}
	return m.verifyForOutbox(r, outbox)
}

func NewSocialPubberTest(t *testing.T) (app *MockApplication, socialApp *MockSocialApp, cb *MockCallbacker, p Pubber) {
	clock := &MockClock{now}
	app = &MockApplication{t: t}
	socialApp = &MockSocialApp{MockApplication: app, t: t}
	cb = &MockCallbacker{t: t}
	p = NewSocialPubber(clock, socialApp, cb)
	return
}

func NewFederatingPubberTest(t *testing.T) (app *MockApplication, fedApp *MockFederateApp, cb *MockCallbacker, d *MockDeliverer, h *MockHttpClient, p Pubber) {
	clock := &MockClock{now}
	app = &MockApplication{t: t}
	fedApp = &MockFederateApp{MockApplication: app, t: t}
	cb = &MockCallbacker{t: t}
	d = &MockDeliverer{t: t}
	h = &MockHttpClient{t: t}
	p = NewFederatingPubber(clock, fedApp, cb, d, h, testAgent, 1, 1)
	return
}

func NewPubberTest(t *testing.T) (app *MockSocialFederateApp, socialApp *MockSocialApp, fedApp *MockFederateApp, socialCb, fedCb *MockCallbacker, d *MockDeliverer, h *MockHttpClient, p Pubber) {
	clock := &MockClock{now}
	appl := &MockApplication{t: t}
	socialApp = &MockSocialApp{t: t}
	fedApp = &MockFederateApp{MockApplication: appl, t: t}
	app = &MockSocialFederateApp{MockSocialApp: socialApp, MockFederateApp: fedApp}
	socialCb = &MockCallbacker{t: t}
	fedCb = &MockCallbacker{t: t}
	d = &MockDeliverer{t: t}
	h = &MockHttpClient{t: t}
	p = NewPubber(clock, app, socialCb, fedCb, d, h, testAgent, 1, 1)
	return
}

func PreparePubberPostInboxTest(t *testing.T, app *MockSocialFederateApp, socialApp *MockSocialApp, fedApp *MockFederateApp, socialCb, fedCb *MockCallbacker, d *MockDeliverer, h *MockHttpClient, p Pubber) {
	PreparePostInboxTest(t, app.MockFederateApp.MockApplication, socialApp, fedApp, socialCb, fedCb, d, h, p)
}

func PreparePostInboxTest(t *testing.T, app *MockApplication, socialApp *MockSocialApp, fedApp *MockFederateApp, socialCb, fedCb *MockCallbacker, d *MockDeliverer, h *MockHttpClient, p Pubber) {
	fedApp.unblocked = func(c context.Context, actorIRIs []*url.URL) error {
		return nil
	}
	app.getInbox = func(c context.Context, r *http.Request, rw RWType) (vocab.OrderedCollectionType, error) {
		if rw != ReadWrite {
			t.Fatalf("expected RWType of %d, got %d", ReadWrite, rw)
		}
		oc := &vocab.OrderedCollection{}
		oc.AppendType("OrderedCollection")
		return oc, nil
	}
	app.set = func(c context.Context, o PubObject) error {
		return nil
	}
	app.has = func(c context.Context, id *url.URL) (bool, error) {
		return false, nil
	}
	return
}

func PreparePubberPostOutboxTest(t *testing.T, app *MockSocialFederateApp, socialApp *MockSocialApp, fedApp *MockFederateApp, socialCb, fedCb *MockCallbacker, d *MockDeliverer, h *MockHttpClient, p Pubber) {
	PreparePostOutboxTest(t, app.MockFederateApp.MockApplication, socialApp, fedApp, socialCb, fedCb, d, h, p)
}

func PreparePostOutboxTest(t *testing.T, app *MockApplication, socialApp *MockSocialApp, fedApp *MockFederateApp, socialCb, fedCb *MockCallbacker, d *MockDeliverer, h *MockHttpClient, p Pubber) {
	socialApp.getPublicKeyForOutbox = func(c context.Context, publicKeyId string, boxIRI *url.URL) (crypto.PublicKey, httpsig.Algorithm, error) {
		return testPrivateKey.Public(), httpsig.RSA_SHA256, nil
	}
	socialApp.getSocialAPIVerifier = func(c context.Context) SocialAPIVerifier {
		return nil
	}
	fedApp.newSigner = func() (httpsig.Signer, error) {
		s, _, err := httpsig.NewSigner([]httpsig.Algorithm{httpsig.RSA_SHA256}, nil, httpsig.Signature)
		if err != nil {
			t.Fatal(err)
		}
		return s, err
	}
	fedApp.privateKey = func(boxIRI *url.URL) (crypto.PrivateKey, string, error) {
		return testPrivateKey, testPublicKeyId, nil
	}
	gotNewId := 0
	app.newId = func(c context.Context, t Typer) *url.URL {
		gotNewId++
		if gotNewId == 1 {
			return testNewIRI
		} else {
			return testNewIRI2
		}
	}
	app.getOutbox = func(c context.Context, r *http.Request, rw RWType) (vocab.OrderedCollectionType, error) {
		if rw != ReadWrite {
			t.Fatalf("expected RWType of %d, got %d", ReadWrite, rw)
		}
		oc := &vocab.OrderedCollection{}
		oc.AppendType("OrderedCollection")
		return oc, nil
	}
	app.set = func(c context.Context, o PubObject) error {
		return nil
	}
	gotHttpDo := 0
	h.do = func(req *http.Request) (*http.Response, error) {
		gotHttpDo++
		if gotHttpDo == 1 {
			actorResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer(samActorJSON)),
			}
			return actorResp, nil
		} else if gotHttpDo == 2 {
			actorResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer(sallyActorJSON)),
			}
			return actorResp, nil
		} else if gotHttpDo == 3 {
			okResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte{})),
			}
			return okResp, nil
		}
		return nil, nil
	}
	d.do = func(b []byte, u *url.URL, toDo func(b []byte, u *url.URL) error) {
		if err := toDo(b, u); err != nil {
			t.Fatalf("Unexpected error in MockDeliverer.Do: %s", err)
		}
	}
	return
}

func TestSocialPubber_RejectPostInbox(t *testing.T) {
	_, _, _, p := NewSocialPubberTest(t)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, nil))
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if resp.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected %d, got %d", http.StatusMethodNotAllowed, resp.Code)
	}
}

func TestSocialPubber_GetInbox(t *testing.T) {
	app, _, _, p := NewSocialPubberTest(t)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("GET", testInboxURI, nil))
	gotInbox := 0
	app.getInbox = func(c context.Context, r *http.Request, rw RWType) (vocab.OrderedCollectionType, error) {
		if rw != Read {
			t.Fatalf("expected RWType of %d, got %d", Read, rw)
		}
		gotInbox++
		return testSingleOrderedCollection, nil
	}
	handled, err := p.GetInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotInbox != 1 {
		t.Fatalf("expected %d, got %d", 1, gotInbox)
	} else if l := len(resp.HeaderMap["Content-Type"]); l != 1 {
		t.Fatalf("expected %d, got %d", 1, l)
	} else if h := resp.HeaderMap["Content-Type"][0]; h != responseContentTypeHeader {
		t.Fatalf("expected %s, got %s", responseContentTypeHeader, h)
	} else if resp.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, resp.Code)
	} else if e := VocabEquals(resp.Body, testSingleOrderedCollection); e != nil {
		t.Fatal(e)
	}
}

func TestSocialPubber_PostOutbox(t *testing.T) {
	app, socialApp, cb, p := NewSocialPubberTest(t)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testCreateNote)))))
	gotPublicKeyForOutbox := 0
	var gotPublicKeyId string
	var gotBoxIRI *url.URL
	socialApp.getPublicKeyForOutbox = func(c context.Context, publicKeyId string, boxIRI *url.URL) (crypto.PublicKey, httpsig.Algorithm, error) {
		gotPublicKeyForOutbox++
		gotPublicKeyId = publicKeyId
		gotBoxIRI = boxIRI
		return testPrivateKey.Public(), httpsig.RSA_SHA256, nil
	}
	gotSocialAPIVerifier := 0
	socialApp.getSocialAPIVerifier = func(c context.Context) SocialAPIVerifier {
		gotSocialAPIVerifier++
		return nil
	}
	gotNewId := 0
	app.newId = func(c context.Context, t Typer) *url.URL {
		gotNewId++
		if gotNewId == 1 {
			return testNewIRI
		}
		return testNewIRI2
	}
	gotOutbox := 0
	app.getOutbox = func(c context.Context, r *http.Request, rw RWType) (vocab.OrderedCollectionType, error) {
		if rw != ReadWrite {
			t.Fatalf("expected RWType of %d, got %d", ReadWrite, rw)
		}
		gotOutbox++
		oc := &vocab.OrderedCollection{}
		oc.AppendType("OrderedCollection")
		return oc, nil
	}
	gotSet := 0
	var gotSetOutbox PubObject
	var gotSetActivity PubObject
	var gotSetCreateObject PubObject
	app.set = func(c context.Context, o PubObject) error {
		gotSet++
		if gotSet == 1 {
			gotSetCreateObject = o
		} else if gotSet == 2 {
			gotSetActivity = o
		} else if gotSet == 3 {
			gotSetOutbox = o
		}
		return nil
	}
	gotCreate := 0
	var gotCreateCallback *streams.Create
	cb.create = func(c context.Context, s *streams.Create) error {
		gotCreate++
		gotCreateCallback = s
		return nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if resp.Code != http.StatusCreated {
		t.Fatalf("expected %d, got %d", http.StatusCreated, resp.Code)
	} else if gotPublicKeyForOutbox != 1 {
		t.Fatalf("expected %d, got %d", 1, gotPublicKeyForOutbox)
	} else if gotPublicKeyId != testPublicKeyId {
		t.Fatalf("expected %s, got %s", testPublicKeyId, gotPublicKeyId)
	} else if s := gotBoxIRI.String(); s != testOutboxURI {
		t.Fatalf("expected %s, got %s", testOutboxURI, s)
	} else if gotSocialAPIVerifier != 1 {
		t.Fatalf("expected %d, got %d", 1, gotSocialAPIVerifier)
	} else if gotNewId != 2 {
		t.Fatalf("expected %d, got %d", 2, gotNewId)
	} else if gotOutbox != 1 {
		t.Fatalf("expected %d, got %d", 1, gotOutbox)
	} else if gotSet != 3 {
		t.Fatalf("expected %d, got %d", 3, gotSet)
	} else if l := gotSetOutbox.GetType(0).(string); l != "OrderedCollection" {
		t.Fatalf("expected %s, got %s", "OrderedCollection", l)
	} else if l := gotSetCreateObject.GetType(0).(string); l != "Note" {
		t.Fatalf("expected %s, got %s", "Note", l)
	} else if l := gotSetActivity.GetType(0).(string); l != "Create" {
		t.Fatalf("expected %s, got %s", "Create", l)
	} else if gotCreate != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCreate)
	} else if iri := gotCreateCallback.Raw().GetActorObject(0).GetId(); *iri != *sallyIRI {
		t.Fatalf("expected %s, got %s", sallyIRIString, iri.String())
	} else if l := len(resp.HeaderMap["Location"]); l != 1 {
		t.Fatalf("expected %d, got %d", 1, l)
	} else if h := resp.HeaderMap["Location"][0]; h != testNewIRIString {
		t.Fatalf("expected %s, got %s", testNewIRI, h)
	} else if resp.Code != http.StatusCreated {
		t.Fatalf("expected %d, got %d", http.StatusCreated, resp.Code)
	}
}

func TestSocialPubber_PostOutbox_SocialAPIVerified(t *testing.T) {
	app, socialApp, cb, p := NewSocialPubberTest(t)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testCreateNote)))))
	gotVerifyForOutbox := 0
	var gotVerifiedOutbox *url.URL
	socialApp.getSocialAPIVerifier = func(c context.Context) SocialAPIVerifier {
		mockV := &MockSocialAPIVerifier{
			verifyForOutbox: func(r *http.Request, outbox *url.URL) (bool, bool, error) {
				gotVerifyForOutbox++
				gotVerifiedOutbox = outbox
				return true, true, nil
			},
		}
		return mockV
	}
	gotNewId := 0
	app.newId = func(c context.Context, t Typer) *url.URL {
		gotNewId++
		if gotNewId == 1 {
			return testNewIRI
		} else {
			return testNewIRI2
		}
	}
	gotOutbox := 0
	app.getOutbox = func(c context.Context, r *http.Request, rw RWType) (vocab.OrderedCollectionType, error) {
		if rw != ReadWrite {
			t.Fatalf("expected RWType of %d, got %d", ReadWrite, rw)
		}
		gotOutbox++
		oc := &vocab.OrderedCollection{}
		oc.AppendType("OrderedCollection")
		return oc, nil
	}
	gotSet := 0
	var gotSetOutbox PubObject
	var gotSetActivity PubObject
	var gotSetCreateObject PubObject
	app.set = func(c context.Context, o PubObject) error {
		gotSet++
		if gotSet == 1 {
			gotSetCreateObject = o
		} else if gotSet == 2 {
			gotSetActivity = o
		} else if gotSet == 3 {
			gotSetOutbox = o
		}
		return nil
	}
	gotCreate := 0
	var gotCreateCallback *streams.Create
	cb.create = func(c context.Context, s *streams.Create) error {
		gotCreate++
		gotCreateCallback = s
		return nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if resp.Code != http.StatusCreated {
		t.Fatalf("expected %d, got %d", http.StatusCreated, resp.Code)
	} else if gotVerifyForOutbox != 1 {
		t.Fatalf("expected %d, got %d", 1, gotVerifyForOutbox)
	} else if o := gotVerifiedOutbox.String(); o != testOutboxURI {
		t.Fatalf("expected %s, got %s", testOutboxURI, o)
	} else if gotNewId != 2 {
		t.Fatalf("expected %d, got %d", 2, gotNewId)
	} else if gotOutbox != 1 {
		t.Fatalf("expected %d, got %d", 1, gotOutbox)
	} else if gotSet != 3 {
		t.Fatalf("expected %d, got %d", 3, gotSet)
	} else if l := gotSetOutbox.GetType(0).(string); l != "OrderedCollection" {
		t.Fatalf("expected %s, got %s", "OrderedCollection", l)
	} else if l := gotSetCreateObject.GetType(0).(string); l != "Note" {
		t.Fatalf("expected %s, got %s", "Note", l)
	} else if l := gotSetActivity.GetType(0).(string); l != "Create" {
		t.Fatalf("expected %s, got %s", "Create", l)
	} else if gotCreate != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCreate)
	} else if iri := gotCreateCallback.Raw().GetActorObject(0).GetId(); *iri != *sallyIRI {
		t.Fatalf("expected %s, got %s", sallyIRIString, iri.String())
	} else if l := len(resp.HeaderMap["Location"]); l != 1 {
		t.Fatalf("expected %d, got %d", 1, l)
	} else if h := resp.HeaderMap["Location"][0]; h != testNewIRIString {
		t.Fatalf("expected %s, got %s", testNewIRI, h)
	} else if resp.Code != http.StatusCreated {
		t.Fatalf("expected %d, got %d", http.StatusCreated, resp.Code)
	}
}

func TestSocialPubber_GetOutbox(t *testing.T) {
	app, _, _, p := NewSocialPubberTest(t)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("GET", testOutboxURI, nil))
	gotOutbox := 0
	app.getOutbox = func(c context.Context, r *http.Request, rw RWType) (vocab.OrderedCollectionType, error) {
		if rw != Read {
			t.Fatalf("expected RWType of %d, got %d", Read, rw)
		}
		gotOutbox++
		return testSingleOrderedCollection, nil
	}
	handled, err := p.GetOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotOutbox != 1 {
		t.Fatalf("expected %d, got %d", 1, gotOutbox)
	} else if l := len(resp.HeaderMap["Content-Type"]); l != 1 {
		t.Fatalf("expected %d, got %d", 1, l)
	} else if h := resp.HeaderMap["Content-Type"][0]; h != responseContentTypeHeader {
		t.Fatalf("expected %s, got %s", responseContentTypeHeader, h)
	} else if resp.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, resp.Code)
	} else if e := VocabEquals(resp.Body, testSingleOrderedCollection); e != nil {
		t.Fatal(e)
	}
}

func TestFederatingPubber_PostInbox(t *testing.T) {
	app, fedApp, cb, _, _, p := NewFederatingPubberTest(t)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testCreateNote))))
	gotUnblocked := 0
	var iri *url.URL
	fedApp.unblocked = func(c context.Context, actorIRIs []*url.URL) error {
		gotUnblocked++
		iri = actorIRIs[0]
		return nil
	}
	gotInbox := 0
	app.getInbox = func(c context.Context, r *http.Request, rw RWType) (vocab.OrderedCollectionType, error) {
		if rw != ReadWrite {
			t.Fatalf("expected RWType of %d, got %d", ReadWrite, rw)
		}
		gotInbox++
		oc := &vocab.OrderedCollection{}
		oc.AppendType("OrderedCollection")
		return oc, nil
	}
	gotSet := 0
	var setObject PubObject
	var inboxObject PubObject
	app.set = func(c context.Context, o PubObject) error {
		gotSet++
		if gotSet == 1 {
			setObject = o
		} else if gotSet == 2 {
			inboxObject = o
		}
		return nil
	}
	gotHas := 0
	var hasIriActivity *url.URL
	var hasIriTo *url.URL
	app.has = func(c context.Context, id *url.URL) (bool, error) {
		gotHas++
		if gotHas == 1 {
			hasIriActivity = id
			return false, nil
		} else {
			hasIriTo = id
			return true, nil
		}
	}
	gotGet := 0
	var gotIri *url.URL
	app.get = func(c context.Context, iri *url.URL, rw RWType) (PubObject, error) {
		if rw != Read {
			t.Fatalf("expected RWType of %d, got %d", Read, rw)
		}
		gotGet++
		gotIri = iri
		return samActor, nil
	}
	gotCreate := 0
	var gotCreateCallback *streams.Create
	cb.create = func(c context.Context, s *streams.Create) error {
		gotCreate++
		gotCreateCallback = s
		return nil
	}
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotUnblocked != 1 {
		t.Fatalf("expected %d, got %d", 1, gotUnblocked)
	} else if iri.String() != sallyIRIString {
		t.Fatalf("expected %s, got %s", sallyIRIString, iri.String())
	} else if gotInbox != 1 {
		t.Fatalf("expected %d, got %d", 1, gotInbox)
	} else if gotSet != 2 {
		t.Fatalf("expected %d, got %d", 2, gotSet)
	} else if l := inboxObject.GetType(0).(string); l != "OrderedCollection" {
		t.Fatalf("expected %s, got %s", "OrderedCollection", l)
	} else if l := setObject.GetType(0).(string); l != "Note" {
		t.Fatalf("expected %s, got %s", "Note", l)
	} else if gotHas != 2 {
		t.Fatalf("expected %d, got %d", 2, gotHas)
	} else if hasIriActivityString := hasIriActivity.String(); hasIriActivityString != noteActivityURIString {
		t.Fatalf("expected %s, got %s", noteActivityURIString, hasIriActivityString)
	} else if hasIriToString := hasIriTo.String(); hasIriToString != samIRIString {
		t.Fatalf("expected %s, got %s", samIRIString, hasIriToString)
	} else if gotGet != 1 {
		t.Fatalf("expected %d, got %d", 1, gotGet)
	} else if gotIriString := gotIri.String(); gotIriString != samIRIString {
		t.Fatalf("expected %s, got %s", samIRIString, gotIriString)
	} else if gotCreate != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCreate)
	} else if s := gotCreateCallback.Raw().GetActorObject(0).GetId(); s.String() != sallyIRIString {
		t.Fatalf("expected %s, got %s", sallyIRIString, s)
	} else if resp.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, resp.Code)
	}
}

func TestFederatingPubber_GetInbox(t *testing.T) {
	app, _, _, _, _, p := NewFederatingPubberTest(t)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("GET", testInboxURI, nil))
	gotInbox := 0
	app.getInbox = func(c context.Context, r *http.Request, rw RWType) (vocab.OrderedCollectionType, error) {
		if rw != Read {
			t.Fatalf("expected RWType of %d, got %d", Read, rw)
		}
		gotInbox++
		return testSingleOrderedCollection, nil
	}
	handled, err := p.GetInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotInbox != 1 {
		t.Fatalf("expected %d, got %d", 1, gotInbox)
	} else if l := len(resp.HeaderMap["Content-Type"]); l != 1 {
		t.Fatalf("expected %d, got %d", 1, l)
	} else if h := resp.HeaderMap["Content-Type"][0]; h != responseContentTypeHeader {
		t.Fatalf("expected %s, got %s", responseContentTypeHeader, h)
	} else if resp.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, resp.Code)
	} else if e := VocabEquals(resp.Body, testSingleOrderedCollection); e != nil {
		t.Fatal(e)
	}
}

func TestFederatingPubber_RejectPostOutbox(t *testing.T) {
	_, _, _, _, _, p := NewFederatingPubberTest(t)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, nil))
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if resp.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected %d, got %d", http.StatusMethodNotAllowed, resp.Code)
	}
}

func TestFederatingPubber_GetOutbox(t *testing.T) {
	app, _, _, _, _, p := NewFederatingPubberTest(t)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("GET", testOutboxURI, nil))
	gotOutbox := 0
	app.getOutbox = func(c context.Context, r *http.Request, rw RWType) (vocab.OrderedCollectionType, error) {
		if rw != Read {
			t.Fatalf("expected RWType of %d, got %d", Read, rw)
		}
		gotOutbox++
		return testSingleOrderedCollection, nil
	}
	handled, err := p.GetOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotOutbox != 1 {
		t.Fatalf("expected %d, got %d", 1, gotOutbox)
	} else if l := len(resp.HeaderMap["Content-Type"]); l != 1 {
		t.Fatalf("expected %d, got %d", 1, l)
	} else if h := resp.HeaderMap["Content-Type"][0]; h != responseContentTypeHeader {
		t.Fatalf("expected %s, got %s", responseContentTypeHeader, h)
	} else if resp.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, resp.Code)
	} else if e := VocabEquals(resp.Body, testSingleOrderedCollection); e != nil {
		t.Fatal(e)
	}
}

func TestPubber_PostInbox(t *testing.T) {
	app, _, fedApp, _, fedCb, _, _, p := NewPubberTest(t)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testCreateNote))))
	gotUnblocked := 0
	var iri *url.URL
	fedApp.unblocked = func(c context.Context, actorIRIs []*url.URL) error {
		gotUnblocked++
		iri = actorIRIs[0]
		return nil
	}
	gotInbox := 0
	app.MockFederateApp.getInbox = func(c context.Context, r *http.Request, rw RWType) (vocab.OrderedCollectionType, error) {
		if rw != ReadWrite {
			t.Fatalf("expected RWType of %d, got %d", ReadWrite, rw)
		}
		gotInbox++
		oc := &vocab.OrderedCollection{}
		oc.AppendType("OrderedCollection")
		return oc, nil
	}
	gotSet := 0
	var setObject PubObject
	var inboxObject PubObject
	app.MockFederateApp.set = func(c context.Context, o PubObject) error {
		gotSet++
		if gotSet == 1 {
			setObject = o
		} else if gotSet == 2 {
			inboxObject = o
		}
		return nil
	}
	gotHas := 0
	var hasIriActivity *url.URL
	var hasIriTo *url.URL
	app.MockFederateApp.has = func(c context.Context, id *url.URL) (bool, error) {
		gotHas++
		if gotHas == 1 {
			hasIriActivity = id
			return false, nil
		} else {
			hasIriTo = id
			return true, nil
		}
	}
	gotGet := 0
	var gotIri *url.URL
	app.MockFederateApp.get = func(c context.Context, iri *url.URL, rw RWType) (PubObject, error) {
		if rw != Read {
			t.Fatalf("expected RWType of %d, got %d", Read, rw)
		}
		gotGet++
		gotIri = iri
		return samActor, nil
	}
	gotCreate := 0
	var gotCreateCallback *streams.Create
	fedCb.create = func(c context.Context, s *streams.Create) error {
		gotCreate++
		gotCreateCallback = s
		return nil
	}
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotUnblocked != 1 {
		t.Fatalf("expected %d, got %d", 1, gotUnblocked)
	} else if iri.String() != sallyIRIString {
		t.Fatalf("expected %s, got %s", sallyIRIString, iri.String())
	} else if gotInbox != 1 {
		t.Fatalf("expected %d, got %d", 1, gotInbox)
	} else if gotSet != 2 {
		t.Fatalf("expected %d, got %d", 2, gotSet)
	} else if l := inboxObject.GetType(0).(string); l != "OrderedCollection" {
		t.Fatalf("expected %s, got %s", "OrderedCollection", l)
	} else if l := setObject.GetType(0).(string); l != "Note" {
		t.Fatalf("expected %s, got %s", "Note", l)
	} else if gotHas != 2 {
		t.Fatalf("expected %d, got %d", 2, gotHas)
	} else if hasIriActivityString := hasIriActivity.String(); hasIriActivityString != noteActivityURIString {
		t.Fatalf("expected %s, got %s", noteActivityURIString, hasIriActivityString)
	} else if hasIriToString := hasIriTo.String(); hasIriToString != samIRIString {
		t.Fatalf("expected %s, got %s", samIRIString, hasIriToString)
	} else if gotGet != 1 {
		t.Fatalf("expected %d, got %d", 1, gotGet)
	} else if gotIriString := gotIri.String(); gotIriString != samIRIString {
		t.Fatalf("expected %s, got %s", samIRIString, gotIriString)
	} else if gotCreate != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCreate)
	} else if s := gotCreateCallback.Raw().GetActorObject(0).GetId(); s.String() != sallyIRIString {
		t.Fatalf("expected %s, got %s", sallyIRIString, s)
	} else if resp.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, resp.Code)
	}
}

func TestPubber_GetInbox(t *testing.T) {
	app, _, _, _, _, _, _, p := NewPubberTest(t)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("GET", testInboxURI, nil))
	gotInbox := 0
	app.MockFederateApp.getInbox = func(c context.Context, r *http.Request, rw RWType) (vocab.OrderedCollectionType, error) {
		if rw != Read {
			t.Fatalf("expected RWType of %d, got %d", Read, rw)
		}
		gotInbox++
		return testSingleOrderedCollection, nil
	}
	handled, err := p.GetInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotInbox != 1 {
		t.Fatalf("expected %d, got %d", 1, gotInbox)
	} else if l := len(resp.HeaderMap["Content-Type"]); l != 1 {
		t.Fatalf("expected %d, got %d", 1, l)
	} else if h := resp.HeaderMap["Content-Type"][0]; h != responseContentTypeHeader {
		t.Fatalf("expected %s, got %s", responseContentTypeHeader, h)
	} else if resp.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, resp.Code)
	} else if e := VocabEquals(resp.Body, testSingleOrderedCollection); e != nil {
		t.Fatal(e)
	}
}

func TestPubber_PostOutbox(t *testing.T) {
	app, socialApp, fedApp, socialCb, _, d, httpClient, p := NewPubberTest(t)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testCreateNote)))))
	gotPublicKey := 0
	var gotPublicKeyId string
	var gotBoxIRI *url.URL
	socialApp.getPublicKeyForOutbox = func(c context.Context, publicKeyId string, boxIRI *url.URL) (crypto.PublicKey, httpsig.Algorithm, error) {
		gotPublicKey++
		gotPublicKeyId = publicKeyId
		gotBoxIRI = boxIRI
		return testPrivateKey.Public(), httpsig.RSA_SHA256, nil
	}
	gotSocialAPIVerifier := 0
	socialApp.getSocialAPIVerifier = func(c context.Context) SocialAPIVerifier {
		gotSocialAPIVerifier++
		return nil
	}
	gotNewSigner := 0
	fedApp.newSigner = func() (httpsig.Signer, error) {
		gotNewSigner++
		s, _, err := httpsig.NewSigner([]httpsig.Algorithm{httpsig.RSA_SHA256}, nil, httpsig.Signature)
		if err != nil {
			t.Fatal(err)
		}
		return s, err
	}
	gotPrivateKey := 0
	var gotPrivateKeyIRI *url.URL
	fedApp.privateKey = func(boxIRI *url.URL) (crypto.PrivateKey, string, error) {
		gotPrivateKey++
		gotPrivateKeyIRI = boxIRI
		return testPrivateKey, testPublicKeyId, nil
	}
	gotNewId := 0
	app.MockFederateApp.newId = func(c context.Context, t Typer) *url.URL {
		gotNewId++
		if gotNewId == 1 {
			return testNewIRI
		} else {
			return testNewIRI2
		}
	}
	gotOutbox := 0
	app.MockFederateApp.getOutbox = func(c context.Context, r *http.Request, rw RWType) (vocab.OrderedCollectionType, error) {
		if rw != ReadWrite {
			t.Fatalf("expected RWType of %d, got %d", ReadWrite, rw)
		}
		gotOutbox++
		oc := &vocab.OrderedCollection{}
		oc.AppendType("OrderedCollection")
		return oc, nil
	}
	gotSet := 0
	var gotSetOutbox PubObject
	var gotSetActivity PubObject
	var gotSetCreateObject PubObject
	app.MockFederateApp.set = func(c context.Context, o PubObject) error {
		gotSet++
		if gotSet == 1 {
			gotSetCreateObject = o
		} else if gotSet == 2 {
			gotSetActivity = o
		} else if gotSet == 3 {
			gotSetOutbox = o
		}
		return nil
	}
	gotCreate := 0
	var gotCreateCallback *streams.Create
	socialCb.create = func(c context.Context, s *streams.Create) error {
		gotCreate++
		gotCreateCallback = s
		return nil
	}
	gotHttpDo := 0
	var httpActorRequest *http.Request
	var httpSenderRequest *http.Request
	var httpDeliveryRequest *http.Request
	httpClient.do = func(req *http.Request) (*http.Response, error) {
		gotHttpDo++
		if gotHttpDo == 1 {
			httpActorRequest = req
			actorResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer(samActorJSON)),
			}
			return actorResp, nil
		} else if gotHttpDo == 2 {
			httpSenderRequest = req
			actorResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer(sallyActorJSON)),
			}
			return actorResp, nil
		} else if gotHttpDo == 3 {
			httpDeliveryRequest = req
			okResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte{})),
			}
			return okResp, nil
		}
		return nil, nil
	}
	gotDoDelivery := 0
	var doDeliveryURL *url.URL
	d.do = func(b []byte, u *url.URL, toDo func(b []byte, u *url.URL) error) {
		gotDoDelivery++
		doDeliveryURL = u
		if err := toDo(b, u); err != nil {
			t.Fatalf("Unexpected error in MockDeliverer.Do: %s", err)
		}
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotPublicKey != 1 {
		t.Fatalf("expected %d, got %d", 1, gotPublicKey)
	} else if gotPublicKeyId != testPublicKeyId {
		t.Fatalf("expected %s, got %s", testPublicKeyId, gotPublicKeyId)
	} else if s := gotBoxIRI.String(); s != testOutboxURI {
		t.Fatalf("expected %s, got %s", testOutboxURI, s)
	} else if gotSocialAPIVerifier != 1 {
		t.Fatalf("expected %d, got %d", 1, gotSocialAPIVerifier)
	} else if gotNewId != 2 {
		t.Fatalf("expected %d, got %d", 2, gotNewId)
	} else if gotNewSigner != 1 {
		t.Fatalf("expected %d, got %d", 1, gotNewSigner)
	} else if gotPrivateKey != 1 {
		t.Fatalf("expected %d, got %d", 1, gotPrivateKey)
	} else if s := gotPrivateKeyIRI.String(); s != testOutboxURI {
		t.Fatalf("expected %s, got %s", testOutboxURI, s)
	} else if gotOutbox != 1 {
		t.Fatalf("expected %d, got %d", 1, gotOutbox)
	} else if gotSet != 3 {
		t.Fatalf("expected %d, got %d", 3, gotSet)
	} else if l := gotSetOutbox.GetType(0).(string); l != "OrderedCollection" {
		t.Fatalf("expected %s, got %s", "OrderedCollection", l)
	} else if l := gotSetCreateObject.GetType(0).(string); l != "Note" {
		t.Fatalf("expected %s, got %s", "Note", l)
	} else if l := gotSetActivity.GetType(0).(string); l != "Create" {
		t.Fatalf("expected %s, got %s", "Create", l)
	} else if gotCreate != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCreate)
	} else if iri := gotCreateCallback.Raw().GetActorObject(0).GetId(); iri.String() != sallyIRIString {
		t.Fatalf("expected %s, got %s", sallyIRIString, iri.String())
	} else if gotDoDelivery != 1 {
		t.Fatalf("expected %d, got %d", 1, gotDoDelivery)
	} else if doDeliveryURL.String() != samIRIInboxString {
		t.Fatalf("expected %s, got %s", samIRIInboxString, doDeliveryURL.String())
	} else if gotHttpDo != 3 {
		t.Fatalf("expected %d, got %d", 3, gotHttpDo)
	} else if h := httpActorRequest.Header.Get("Accept"); h != getAcceptHeader {
		t.Fatalf("expected %s, got %s", getAcceptHeader, h)
	} else if h := httpSenderRequest.Header.Get("Accept"); h != getAcceptHeader {
		t.Fatalf("expected %s, got %s", getAcceptHeader, h)
	} else if h := httpDeliveryRequest.Header.Get("Content-Type"); h != postContentTypeHeader {
		t.Fatalf("expected %s, got %s", postContentTypeHeader, h)
	} else if httpActorRequest.Method != "GET" {
		t.Fatalf("expected %s, got %s", "GET", httpActorRequest.Method)
	} else if httpSenderRequest.Method != "GET" {
		t.Fatalf("expected %s, got %s", "GET", httpSenderRequest.Method)
	} else if httpDeliveryRequest.Method != "POST" {
		t.Fatalf("expected %s, got %s", "POST", httpDeliveryRequest.Method)
	} else if s := httpActorRequest.URL.String(); s != samIRIString {
		t.Fatalf("expected %s, got %s", samIRIString, s)
	} else if s := httpSenderRequest.URL.String(); s != sallyIRIString {
		t.Fatalf("expected %s, got %s", sallyIRIString, s)
	} else if s := httpDeliveryRequest.URL.String(); s != samIRIInboxString {
		t.Fatalf("expected %s, got %s", samIRIInboxString, s)
	} else if l := len(resp.HeaderMap["Location"]); l != 1 {
		t.Fatalf("expected %d, got %d", 1, l)
	} else if h := resp.HeaderMap["Location"][0]; h != testNewIRIString {
		t.Fatalf("expected %s, got %s", testNewIRI, h)
	} else if resp.Code != http.StatusCreated {
		t.Fatalf("expected %d, got %d", http.StatusCreated, resp.Code)
	}
	verif, err := httpsig.NewVerifier(httpDeliveryRequest)
	if err != nil {
		t.Fatal(err)
	} else if verif.KeyId() != testPublicKeyId {
		t.Fatalf("expected %s, got %s", testPublicKeyId, verif.KeyId())
	} else if err := verif.Verify(testPrivateKey.Public(), httpsig.RSA_SHA256); err != nil {
		t.Fatal(err)
	}
}

func TestPubber_GetOutbox(t *testing.T) {
	app, _, _, _, _, _, _, p := NewPubberTest(t)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("GET", testOutboxURI, nil))
	gotOutbox := 0
	app.MockFederateApp.getOutbox = func(c context.Context, r *http.Request, rw RWType) (vocab.OrderedCollectionType, error) {
		if rw != Read {
			t.Fatalf("expected RWType of %d, got %d", Read, rw)
		}
		gotOutbox++
		return testSingleOrderedCollection, nil
	}
	handled, err := p.GetOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotOutbox != 1 {
		t.Fatalf("expected %d, got %d", 1, gotOutbox)
	} else if l := len(resp.HeaderMap["Content-Type"]); l != 1 {
		t.Fatalf("expected %d, got %d", 1, l)
	} else if h := resp.HeaderMap["Content-Type"][0]; h != responseContentTypeHeader {
		t.Fatalf("expected %s, got %s", responseContentTypeHeader, h)
	} else if resp.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, resp.Code)
	} else if e := VocabEquals(resp.Body, testSingleOrderedCollection); e != nil {
		t.Fatal(e)
	}
}

func TestPostInbox_RejectNonActivityPub(t *testing.T) {
	_, _, _, _, _, _, _, p := NewPubberTest(t)
	resp := httptest.NewRecorder()
	req := httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testCreateNote)))
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if handled {
		t.Fatalf("expected !handled, got handled")
	}
}

func TestPostInbox_HandlesBlocked(t *testing.T) {
	_, _, fedApp, _, _, _, _, p := NewPubberTest(t)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testCreateNote))))
	blockedErr := fmt.Errorf("blocked")
	gotBlocked := 0
	var iri *url.URL
	fedApp.unblocked = func(c context.Context, actorIRIs []*url.URL) error {
		gotBlocked++
		iri = actorIRIs[0]
		return blockedErr
	}
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != blockedErr {
		t.Fatalf("expected %s, got %s", blockedErr, err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if *iri != *sallyIRI {
		t.Fatalf("expected %s, got %s", sallyIRIString, iri.String())
	}
}

func TestPostInbox_RequiresObject(t *testing.T) {
	tests := []struct {
		name  string
		input func() vocab.Serializer
	}{
		{
			name: "create",
			input: func() vocab.Serializer {
				v := &vocab.Create{}
				v.SetId(noteActivityIRI)
				v.AppendSummaryString("Sally created a note")
				v.AppendActorObject(sallyActor)
				v.AppendToObject(samActor)
				return v
			},
		},
		{
			name: "update",
			input: func() vocab.Serializer {
				v := &vocab.Update{}
				v.SetId(noteActivityIRI)
				v.AppendSummaryString("Sally updated a note")
				v.AppendActorObject(sallyActor)
				v.AppendToObject(samActor)
				return v
			},
		},
		{
			name: "delete",
			input: func() vocab.Serializer {
				v := &vocab.Delete{}
				v.SetId(noteActivityIRI)
				v.AppendActorObject(sallyActor)
				v.AppendToObject(samActor)
				return v
			},
		},
		{
			name: "follow",
			input: func() vocab.Serializer {
				v := &vocab.Follow{}
				v.SetId(noteActivityIRI)
				v.AppendActorObject(sallyActor)
				v.AppendToObject(samActor)
				return v
			},
		},
		{
			name: "add",
			input: func() vocab.Serializer {
				v := &vocab.Add{}
				v.SetId(noteActivityIRI)
				v.AppendActorObject(sallyActor)
				v.AppendToObject(samActor)
				v.AppendTargetObject(testNote)
				return v
			},
		},
		{
			name: "remove",
			input: func() vocab.Serializer {
				v := &vocab.Remove{}
				v.SetId(noteActivityIRI)
				v.AppendActorObject(sallyActor)
				v.AppendToObject(samActor)
				v.AppendTargetObject(testNote)
				return v
			},
		},
		{
			name: "like",
			input: func() vocab.Serializer {
				v := &vocab.Like{}
				v.SetId(noteActivityIRI)
				v.AppendActorObject(sallyActor)
				v.AppendToObject(samActor)
				return v
			},
		},
		{
			name: "block",
			input: func() vocab.Serializer {
				v := &vocab.Block{}
				v.SetId(noteActivityIRI)
				v.AppendActorObject(sallyActor)
				v.AppendToObject(samActor)
				return v
			},
		},
		{
			name: "undo",
			input: func() vocab.Serializer {
				v := &vocab.Undo{}
				v.SetId(noteActivityIRI)
				v.AppendActorObject(sallyActor)
				v.AppendToObject(samActor)
				return v
			},
		},
	}
	app, _, fedApp, _, _, _, _, p := NewPubberTest(t)
	fedApp.unblocked = func(c context.Context, actorIRIs []*url.URL) error {
		return nil
	}
	app.MockFederateApp.getInbox = func(c context.Context, r *http.Request, rw RWType) (vocab.OrderedCollectionType, error) {
		inbox := &vocab.OrderedCollection{}
		return inbox, nil
	}
	for _, test := range tests {
		t.Logf("Running table test case %q", test.name)
		resp := httptest.NewRecorder()
		req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(test.input()))))
		handled, err := p.PostInbox(context.Background(), resp, req)
		if resp.Code != http.StatusBadRequest {
			t.Fatalf("(%s) expected %d, got %d", test.name, http.StatusBadRequest, resp.Code)
		} else if err != nil {
			t.Fatal(err)
		} else if !handled {
			t.Fatalf("(%s) expected handled, got !handled", test.name)
		}
	}
}

func TestPostInbox_RequiresTarget(t *testing.T) {
	tests := []struct {
		name  string
		input func() vocab.Serializer
	}{
		{
			name: "add",
			input: func() vocab.Serializer {
				v := &vocab.Add{}
				v.SetId(noteActivityIRI)
				v.AppendActorObject(sallyActor)
				v.AppendToObject(samActor)
				v.AppendObject(testNote)
				return v
			},
		},
		{
			name: "remove",
			input: func() vocab.Serializer {
				v := &vocab.Remove{}
				v.SetId(noteActivityIRI)
				v.AppendActorObject(sallyActor)
				v.AppendToObject(samActor)
				v.AppendObject(testNote)
				return v
			},
		},
	}
	app, _, fedApp, _, _, _, _, p := NewPubberTest(t)
	fedApp.unblocked = func(c context.Context, actorIRIs []*url.URL) error {
		return nil
	}
	app.MockFederateApp.getInbox = func(c context.Context, r *http.Request, rw RWType) (vocab.OrderedCollectionType, error) {
		inbox := &vocab.OrderedCollection{}
		return inbox, nil
	}
	for _, test := range tests {
		t.Logf("Running table test case %q", test.name)
		resp := httptest.NewRecorder()
		req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(test.input()))))
		handled, err := p.PostInbox(context.Background(), resp, req)
		if resp.Code != http.StatusBadRequest {
			t.Fatalf("(%s) expected %d, got %d", test.name, http.StatusBadRequest, resp.Code)
		} else if err != nil {
			t.Fatal(err)
		} else if !handled {
			t.Fatalf("(%s) expected handled, got !handled", test.name)
		}
	}
}

func TestPostInbox_DoesNotAddToInboxIfDuplicate(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testCreateNote))))
	gotSet := 0
	app.MockFederateApp.set = func(c context.Context, o PubObject) error {
		gotSet++
		return nil
	}
	app.MockFederateApp.getInbox = func(c context.Context, r *http.Request, rw RWType) (vocab.OrderedCollectionType, error) {
		inbox := &vocab.OrderedCollection{}
		inbox.AppendOrderedItemsIRI(noteActivityIRI)
		return inbox, nil
	}
	fedCb.create = func(c context.Context, s *streams.Create) error {
		return nil
	}
	handled, err := p.PostInbox(context.Background(), resp, req)
	expectedInbox := &vocab.OrderedCollection{}
	expectedInbox.AppendOrderedItemsIRI(noteActivityIRI)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotSet != 0 {
		t.Fatalf("expected %d, got %d", 0, gotSet)
	}
}

func TestPostInbox_OriginMustMatch(t *testing.T) {
	tests := []struct {
		name  string
		input func() vocab.ActivityType
	}{
		{
			name: "update",
			input: func() vocab.ActivityType {
				a := &vocab.Update{}
				a.SetId(otherOriginIRI)
				a.AppendActorIRI(otherOriginActorIRI)
				a.AppendObject(testCreateNote)
				return a
			},
		},
		{
			name: "delete",
			input: func() vocab.ActivityType {
				a := &vocab.Delete{}
				a.SetId(otherOriginIRI)
				a.AppendActorIRI(otherOriginActorIRI)
				a.AppendObject(testCreateNote)
				return a
			},
		},
		{
			name: "update",
			input: func() vocab.ActivityType {
				a := &vocab.Update{}
				a.SetId(otherOriginIRI)
				a.AppendActorIRI(otherOriginActorIRI)
				a.AppendObjectIRI(noteIRI)
				return a
			},
		},
		{
			name: "delete",
			input: func() vocab.ActivityType {
				a := &vocab.Delete{}
				a.SetId(otherOriginIRI)
				a.AppendActorIRI(otherOriginActorIRI)
				a.AppendObjectIRI(noteIRI)
				return a
			},
		},
	}
	for _, test := range tests {
		t.Logf("Running table test case %q", test.name)
		app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
		PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
		resp := httptest.NewRecorder()
		req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(test.input()))))
		handled, err := p.PostInbox(context.Background(), resp, req)
		if !handled {
			t.Fatalf("%q: expected handled, got !handled", test.name)
		} else if err == nil {
			t.Fatalf("%q: expected same origin error", test.name)
		}
	}
}

func TestPostInbox_ActivityActorsMustCoverObjectActors(t *testing.T) {
	tests := []struct {
		name  string
		input func() vocab.ActivityType
	}{
		{
			name: "undo",
			input: func() vocab.ActivityType {
				a := &vocab.Undo{}
				a.SetId(noteActivityIRI)
				a.AppendActorObject(samActor)
				a.AppendObject(testLikeNote)
				return a
			},
		},
	}
	for _, test := range tests {
		t.Logf("Running table test case %q", test.name)
		app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
		PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
		resp := httptest.NewRecorder()
		req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(test.input()))))
		handled, err := p.PostInbox(context.Background(), resp, req)
		if !handled {
			t.Fatalf("%q: expected handled, got !handled", test.name)
		} else if err == nil {
			t.Fatalf("%q: expected actor mismatch error", test.name)
		}
	}
}

func TestPostInbox_Create_SetsObject(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testCreateNote))))
	gotSet := 0
	var setObject PubObject
	var gotSetInbox PubObject
	app.MockFederateApp.set = func(c context.Context, o PubObject) error {
		gotSet++
		if gotSet == 1 {
			setObject = o
		} else if gotSet == 2 {
			gotSetInbox = o
		}
		return nil
	}
	fedCb.create = func(c context.Context, s *streams.Create) error {
		return nil
	}
	handled, err := p.PostInbox(context.Background(), resp, req)
	expectedInbox := &vocab.OrderedCollection{}
	expectedInbox.AppendType("OrderedCollection")
	expectedInbox.AppendOrderedItemsIRI(testCreateNote.GetId())
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotSet != 2 {
		t.Fatalf("expected %d, got %d", 2, gotSet)
	} else if err := PubObjectEquals(setObject, testNote); err != nil {
		t.Fatalf("unexpected set object: %s", err)
	} else if err := PubObjectEquals(gotSetInbox, expectedInbox); err != nil {
		t.Fatalf("unexpected callback object: %s", err)
	}
}

func TestPostInbox_Create_CallsCallback(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testCreateNote))))
	gotCreate := 0
	var gotCreateCallback *streams.Create
	fedCb.create = func(c context.Context, s *streams.Create) error {
		gotCreate++
		gotCreateCallback = s
		return nil
	}
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotCreate != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCreate)
	} else if err := PubObjectEquals(gotCreateCallback.Raw(), testCreateNote); err != nil {
		t.Fatalf("unexpected create callback: %s", err)
	}
}

func TestPostInbox_Update_SetsObject(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testUpdateNote))))
	gotSet := 0
	var setObject PubObject
	app.MockFederateApp.set = func(c context.Context, o PubObject) error {
		gotSet++
		if gotSet == 1 {
			setObject = o
		}
		return nil
	}
	fedCb.update = func(c context.Context, s *streams.Update) error {
		return nil
	}
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotSet != 2 {
		t.Fatalf("expected %d, got %d", 2, gotSet)
	} else if err := PubObjectEquals(setObject, testNote); err != nil {
		t.Fatalf("unexpected set object: %s", err)
	}
}

func TestPostInbox_Update_CallsCallback(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testUpdateNote))))
	gotCallback := 0
	var gotStreamCallback *streams.Update
	fedCb.update = func(c context.Context, s *streams.Update) error {
		gotCallback++
		gotStreamCallback = s
		return nil
	}
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotCallback != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCallback)
	} else if err := PubObjectEquals(gotStreamCallback.Raw(), testUpdateNote); err != nil {
		t.Fatalf("unexpected callback object: %s", err)
	}
}

func TestPostInbox_Delete_FetchesObject(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testDeleteNote))))
	app.MockFederateApp.get = func(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
		if rw != ReadWrite {
			t.Fatalf("expected RWType of %d, got %d", ReadWrite, rw)
		} else if *id != *noteIRI {
			t.Fatalf("expected %s, got %s", noteIRI, id)
		}
		return testNote, nil
	}
	fedCb.delete = func(c context.Context, s *streams.Delete) error {
		return nil
	}
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	}
}

func TestPostInbox_Delete_SetsTombstone(t *testing.T) {
	// Table test
	tests := []struct {
		name     string
		input    func() PubObject
		expected func() vocab.Serializer
	}{
		{
			name: "delete note",
			input: func() PubObject {
				return testNote
			},
			expected: func() vocab.Serializer {
				return testTombstoneNote
			},
		},
		{
			name: "forward published time",
			input: func() PubObject {
				testNote := &vocab.Note{}
				testNote.SetId(noteIRI)
				testNote.AppendType("Note")
				testNote.AppendNameString(noteName)
				testNote.AppendContentString("This is a simple note")
				testNote.SetPublished(now)
				return testNote
			},
			expected: func() vocab.Serializer {
				testTombstoneNote := &vocab.Tombstone{}
				testTombstoneNote.SetId(noteIRI)
				testTombstoneNote.AppendFormerTypeString("Note")
				testTombstoneNote.SetDeleted(now)
				testTombstoneNote.SetPublished(now)
				return testTombstoneNote
			},
		},
		{
			name: "forward published iri",
			input: func() PubObject {
				testNote := &vocab.Note{}
				testNote.SetId(noteIRI)
				testNote.AppendType("Note")
				testNote.AppendNameString(noteName)
				testNote.AppendContentString("This is a simple note")
				testNote.SetPublishedIRI(iri)
				return testNote
			},
			expected: func() vocab.Serializer {
				testTombstoneNote := &vocab.Tombstone{}
				testTombstoneNote.SetId(noteIRI)
				testTombstoneNote.AppendFormerTypeString("Note")
				testTombstoneNote.SetDeleted(now)
				testTombstoneNote.SetPublishedIRI(iri)
				return testTombstoneNote
			},
		},
		{
			name: "forward updated time",
			input: func() PubObject {
				testNote := &vocab.Note{}
				testNote.SetId(noteIRI)
				testNote.AppendType("Note")
				testNote.AppendNameString(noteName)
				testNote.AppendContentString("This is a simple note")
				testNote.SetUpdated(now)
				return testNote
			},
			expected: func() vocab.Serializer {
				testTombstoneNote := &vocab.Tombstone{}
				testTombstoneNote.SetId(noteIRI)
				testTombstoneNote.AppendFormerTypeString("Note")
				testTombstoneNote.SetDeleted(now)
				testTombstoneNote.SetUpdated(now)
				return testTombstoneNote
			},
		},
		{
			name: "forward updated iri",
			input: func() PubObject {
				testNote := &vocab.Note{}
				testNote.SetId(noteIRI)
				testNote.AppendType("Note")
				testNote.AppendNameString(noteName)
				testNote.AppendContentString("This is a simple note")
				testNote.SetUpdatedIRI(iri)
				return testNote
			},
			expected: func() vocab.Serializer {
				testTombstoneNote := &vocab.Tombstone{}
				testTombstoneNote.SetId(noteIRI)
				testTombstoneNote.AppendFormerTypeString("Note")
				testTombstoneNote.SetDeleted(now)
				testTombstoneNote.SetUpdatedIRI(iri)
				return testTombstoneNote
			},
		},
	}
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	gotSet := 0
	var gotSetObject PubObject
	app.MockFederateApp.set = func(c context.Context, p PubObject) error {
		gotSet++
		if gotSet == 1 {
			gotSetObject = p
		}
		return nil
	}
	fedCb.delete = func(c context.Context, s *streams.Delete) error {
		return nil
	}
	for _, test := range tests {
		t.Logf("Running table test case %q", test.name)
		app.MockFederateApp.get = func(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
			return test.input(), nil
		}
		gotSet = 0
		resp := httptest.NewRecorder()
		req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testDeleteNote))))
		handled, err := p.PostInbox(context.Background(), resp, req)
		if err != nil {
			t.Fatalf("(%q) %s", test.name, err)
		} else if !handled {
			t.Fatalf("(%q) expected handled, got !handled", test.name)
		} else if gotSet != 2 {
			t.Fatalf("(%q) expected %d, got %d", test.name, 2, gotSet)
		} else if err := PubObjectEquals(gotSetObject, test.expected()); err != nil {
			t.Fatalf("(%q) unexpected tombstone object: %s", test.name, err)
		}
	}
}

func TestPostInbox_Delete_CallsCallback(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testDeleteNote))))
	app.MockFederateApp.get = func(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
		return testNote, nil
	}
	gotCallback := 0
	var gotStreamCallback *streams.Delete
	fedCb.delete = func(c context.Context, s *streams.Delete) error {
		gotCallback++
		gotStreamCallback = s
		return nil
	}
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotCallback != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCallback)
	} else if err := PubObjectEquals(gotStreamCallback.Raw(), testDeleteNote); err != nil {
		t.Fatalf("unexpected callback object: %s", err)
	}
}

func TestPostInbox_Follow_DoNothing(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testFollow))))
	gotOnFollow := 0
	fedApp.onFollow = func(c context.Context, s *streams.Follow) FollowResponse {
		gotOnFollow++
		return DoNothing
	}
	fedCb.follow = func(c context.Context, s *streams.Follow) error {
		return nil
	}
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotOnFollow != 1 {
		t.Fatalf("expected %d, got %d", 1, gotOnFollow)
	}
}

func TestPostInbox_Follow_AutoReject(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testFollow))))
	gotOnFollow := 0
	fedApp.onFollow = func(c context.Context, s *streams.Follow) FollowResponse {
		gotOnFollow++
		return AutomaticReject
	}
	gotNewSigner := 0
	fedApp.newSigner = func() (httpsig.Signer, error) {
		gotNewSigner++
		s, _, err := httpsig.NewSigner([]httpsig.Algorithm{httpsig.RSA_SHA256}, nil, httpsig.Signature)
		if err != nil {
			t.Fatal(err)
		}
		return s, err
	}
	gotPrivateKey := 0
	var gotPrivateKeyIRI *url.URL
	fedApp.privateKey = func(boxIRI *url.URL) (crypto.PrivateKey, string, error) {
		gotPrivateKey++
		gotPrivateKeyIRI = boxIRI
		return testPrivateKey, testPublicKeyId, nil
	}
	fedCb.follow = func(c context.Context, s *streams.Follow) error {
		return nil
	}
	gotOwns := 0
	var ownsIRI *url.URL
	app.MockFederateApp.owns = func(c context.Context, id *url.URL) bool {
		gotOwns++
		ownsIRI = id
		return true
	}
	gotHttpDo := 0
	var httpActorRequest *http.Request
	var httpDeliveryRequest *http.Request
	httpClient.do = func(req *http.Request) (*http.Response, error) {
		gotHttpDo++
		if gotHttpDo == 1 {
			httpActorRequest = req
			actorResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer(sallyActorJSON)),
			}
			return actorResp, nil
		} else if gotHttpDo == 2 {
			httpDeliveryRequest = req
			okResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte{})),
			}
			return okResp, nil
		}
		return nil, nil
	}
	gotDoDelivery := 0
	var doDeliveryURL *url.URL
	var bytesToSend []byte
	d.do = func(b []byte, u *url.URL, toDo func(b []byte, u *url.URL) error) {
		gotDoDelivery++
		doDeliveryURL = u
		bytesToSend = b
		if err := toDo(b, u); err != nil {
			t.Fatalf("Unexpected error in MockDeliverer.Do: %s", err)
		}
	}
	expected := &vocab.Reject{}
	expected.AppendObject(testFollow)
	expected.AppendToIRI(sallyIRI)
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotNewSigner != 1 {
		t.Fatalf("expected %d, got %d", 1, gotNewSigner)
	} else if gotPrivateKey != 1 {
		t.Fatalf("expected %d, got %d", 1, gotPrivateKey)
	} else if s := gotPrivateKeyIRI.String(); s != testInboxURI {
		t.Fatalf("expected %s, got %s", testInboxURI, s)
	} else if gotOnFollow != 1 {
		t.Fatalf("expected %d, got %d", 1, gotOnFollow)
	} else if gotHttpDo != 2 {
		t.Fatalf("expected %d, got %d", 2, gotHttpDo)
	} else if s := httpActorRequest.URL.String(); s != sallyIRIString {
		t.Fatalf("expected %s, got %s", sallyIRIString, s)
	} else if s := httpDeliveryRequest.URL.String(); s != sallyIRIInboxString {
		t.Fatalf("expected %s, got %s", sallyIRIInboxString, s)
	} else if gotOwns != 1 {
		t.Fatalf("expected %d, got %d", 1, gotOwns)
	} else if ownsIRI.String() != samIRIString {
		t.Fatalf("expected %s, got %s", samIRIString, ownsIRI.String())
	} else if gotDoDelivery != 1 {
		t.Fatalf("expected %d, got %d", 1, gotDoDelivery)
	} else if doDeliveryURL.String() != sallyIRIInboxString {
		t.Fatalf("expected %s, got %s", sallyIRIInboxString, doDeliveryURL.String())
	} else if err := VocabEquals(bytes.NewBuffer(bytesToSend), expected); err != nil {
		t.Fatal(err)
	}
}

func TestPostInbox_Follow_AutoAccept(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testFollow))))
	gotOnFollow := 0
	fedApp.onFollow = func(c context.Context, s *streams.Follow) FollowResponse {
		gotOnFollow++
		return AutomaticAccept
	}
	gotNewSigner := 0
	fedApp.newSigner = func() (httpsig.Signer, error) {
		gotNewSigner++
		s, _, err := httpsig.NewSigner([]httpsig.Algorithm{httpsig.RSA_SHA256}, nil, httpsig.Signature)
		if err != nil {
			t.Fatal(err)
		}
		return s, err
	}
	gotPrivateKey := 0
	var gotPrivateKeyIRI *url.URL
	fedApp.privateKey = func(boxIRI *url.URL) (crypto.PrivateKey, string, error) {
		gotPrivateKey++
		gotPrivateKeyIRI = boxIRI
		return testPrivateKey, testPublicKeyId, nil
	}
	fedCb.follow = func(c context.Context, s *streams.Follow) error {
		return nil
	}
	gotHttpDo := 0
	var httpActorRequest *http.Request
	var httpDeliveryRequest *http.Request
	httpClient.do = func(req *http.Request) (*http.Response, error) {
		gotHttpDo++
		if gotHttpDo == 1 {
			httpActorRequest = req
			actorResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer(sallyActorJSON)),
			}
			return actorResp, nil
		} else if gotHttpDo == 2 {
			httpDeliveryRequest = req
			okResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte{})),
			}
			return okResp, nil
		}
		return nil, nil
	}
	gotDoDelivery := 0
	var doDeliveryURL *url.URL
	var bytesToSend []byte
	d.do = func(b []byte, u *url.URL, toDo func(b []byte, u *url.URL) error) {
		gotDoDelivery++
		doDeliveryURL = u
		bytesToSend = b
		if err := toDo(b, u); err != nil {
			t.Fatalf("Unexpected error in MockDeliverer.Do: %s", err)
		}
	}
	gotOwns := 0
	var ownsIRI *url.URL
	app.MockFederateApp.owns = func(c context.Context, id *url.URL) bool {
		gotOwns++
		ownsIRI = id
		return true
	}
	gotGet := 0
	var getIRI *url.URL
	app.MockFederateApp.get = func(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
		if rw != ReadWrite {
			t.Fatalf("expected RWType of %d, got %d", ReadWrite, rw)
		}
		gotGet++
		getIRI = id
		samActor := &vocab.Person{}
		samActor.SetInboxAnyURI(samIRIInbox)
		samActor.SetId(samIRI)
		samActor.SetFollowersCollection(&vocab.Collection{})
		return samActor, nil
	}
	gotSet := 0
	var setObject PubObject
	app.MockFederateApp.set = func(c context.Context, o PubObject) error {
		gotSet++
		if gotSet == 1 {
			setObject = o
		}
		return nil
	}
	expected := &vocab.Accept{}
	expected.AppendObject(testFollow)
	expected.AppendToIRI(sallyIRI)
	expectedFollowers := &vocab.Collection{}
	expectedFollowers.AppendItemsIRI(sallyIRI)
	expectedActor := &vocab.Person{}
	expectedActor.SetInboxAnyURI(samIRIInbox)
	expectedActor.SetId(samIRI)
	expectedActor.SetFollowersCollection(expectedFollowers)
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotOnFollow != 1 {
		t.Fatalf("expected %d, got %d", 1, gotOnFollow)
	} else if gotNewSigner != 1 {
		t.Fatalf("expected %d, got %d", 1, gotNewSigner)
	} else if gotPrivateKey != 1 {
		t.Fatalf("expected %d, got %d", 1, gotPrivateKey)
	} else if s := gotPrivateKeyIRI.String(); s != testInboxURI {
		t.Fatalf("expected %s, got %s", testInboxURI, s)
	} else if gotHttpDo != 2 {
		t.Fatalf("expected %d, got %d", 2, gotHttpDo)
	} else if s := httpActorRequest.URL.String(); s != sallyIRIString {
		t.Fatalf("expected %s, got %s", sallyIRIString, s)
	} else if s := httpDeliveryRequest.URL.String(); s != sallyIRIInboxString {
		t.Fatalf("expected %s, got %s", sallyIRIInboxString, s)
	} else if gotDoDelivery != 1 {
		t.Fatalf("expected %d, got %d", 1, gotDoDelivery)
	} else if doDeliveryURL.String() != sallyIRIInboxString {
		t.Fatalf("expected %s, got %s", sallyIRIInboxString, doDeliveryURL.String())
	} else if err := VocabEquals(bytes.NewBuffer(bytesToSend), expected); err != nil {
		t.Fatal(err)
	} else if gotOwns != 1 {
		t.Fatalf("expected %d, got %d", 1, gotOwns)
	} else if ownsIRI.String() != samIRIString {
		t.Fatalf("expected %s, got %s", samIRIString, ownsIRI.String())
	} else if gotGet != 1 {
		t.Fatalf("expected %d, got %d", 1, gotGet)
	} else if getIRI.String() != samIRIString {
		t.Fatalf("expected %s, got %s", samIRIString, getIRI.String())
	} else if gotSet != 2 {
		t.Fatalf("expected %d, got %d", 2, gotSet)
	} else if err := PubObjectEquals(setObject, expectedActor); err != nil {
		t.Fatal(err)
	}
}

func TestPostInbox_Follow_AutoAcceptDefaultFollowersOrderedCollection(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testFollow))))
	gotOnFollow := 0
	fedApp.onFollow = func(c context.Context, s *streams.Follow) FollowResponse {
		gotOnFollow++
		return AutomaticAccept
	}
	gotNewSigner := 0
	fedApp.newSigner = func() (httpsig.Signer, error) {
		gotNewSigner++
		s, _, err := httpsig.NewSigner([]httpsig.Algorithm{httpsig.RSA_SHA256}, nil, httpsig.Signature)
		if err != nil {
			t.Fatal(err)
		}
		return s, err
	}
	gotPrivateKey := 0
	var gotPrivateKeyIRI *url.URL
	fedApp.privateKey = func(boxIRI *url.URL) (crypto.PrivateKey, string, error) {
		gotPrivateKey++
		gotPrivateKeyIRI = boxIRI
		return testPrivateKey, testPublicKeyId, nil
	}
	fedCb.follow = func(c context.Context, s *streams.Follow) error {
		return nil
	}
	gotHttpDo := 0
	var httpActorRequest *http.Request
	var httpDeliveryRequest *http.Request
	httpClient.do = func(req *http.Request) (*http.Response, error) {
		gotHttpDo++
		if gotHttpDo == 1 {
			httpActorRequest = req
			actorResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer(sallyActorJSON)),
			}
			return actorResp, nil
		} else if gotHttpDo == 2 {
			httpDeliveryRequest = req
			okResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte{})),
			}
			return okResp, nil
		}
		return nil, nil
	}
	gotDoDelivery := 0
	var doDeliveryURL *url.URL
	var bytesToSend []byte
	d.do = func(b []byte, u *url.URL, toDo func(b []byte, u *url.URL) error) {
		gotDoDelivery++
		doDeliveryURL = u
		bytesToSend = b
		if err := toDo(b, u); err != nil {
			t.Fatalf("Unexpected error in MockDeliverer.Do: %s", err)
		}
	}
	gotOwns := 0
	var ownsIRI *url.URL
	app.MockFederateApp.owns = func(c context.Context, id *url.URL) bool {
		gotOwns++
		ownsIRI = id
		return true
	}
	gotGet := 0
	var getIRI *url.URL
	app.MockFederateApp.get = func(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
		if rw != ReadWrite {
			t.Fatalf("expected RWType of %d, got %d", ReadWrite, rw)
		}
		gotGet++
		getIRI = id
		samActor := &vocab.Person{}
		samActor.SetInboxAnyURI(samIRIInbox)
		samActor.SetId(samIRI)
		return samActor, nil
	}
	gotSet := 0
	var setObject PubObject
	app.MockFederateApp.set = func(c context.Context, o PubObject) error {
		gotSet++
		if gotSet == 1 {
			setObject = o
		}
		return nil
	}
	expected := &vocab.Accept{}
	expected.AppendObject(testFollow)
	expected.AppendToIRI(sallyIRI)
	expectedFollowers := &vocab.OrderedCollection{}
	expectedFollowers.AppendOrderedItemsIRI(sallyIRI)
	expectedActor := &vocab.Person{}
	expectedActor.SetInboxAnyURI(samIRIInbox)
	expectedActor.SetId(samIRI)
	expectedActor.SetFollowersOrderedCollection(expectedFollowers)
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotOnFollow != 1 {
		t.Fatalf("expected %d, got %d", 1, gotOnFollow)
	} else if gotNewSigner != 1 {
		t.Fatalf("expected %d, got %d", 1, gotNewSigner)
	} else if gotPrivateKey != 1 {
		t.Fatalf("expected %d, got %d", 1, gotPrivateKey)
	} else if s := gotPrivateKeyIRI.String(); s != testInboxURI {
		t.Fatalf("expected %s, got %s", testInboxURI, s)
	} else if gotHttpDo != 2 {
		t.Fatalf("expected %d, got %d", 2, gotHttpDo)
	} else if s := httpActorRequest.URL.String(); s != sallyIRIString {
		t.Fatalf("expected %s, got %s", sallyIRIString, s)
	} else if s := httpDeliveryRequest.URL.String(); s != sallyIRIInboxString {
		t.Fatalf("expected %s, got %s", sallyIRIInboxString, s)
	} else if gotDoDelivery != 1 {
		t.Fatalf("expected %d, got %d", 1, gotDoDelivery)
	} else if doDeliveryURL.String() != sallyIRIInboxString {
		t.Fatalf("expected %s, got %s", sallyIRIInboxString, doDeliveryURL.String())
	} else if err := VocabEquals(bytes.NewBuffer(bytesToSend), expected); err != nil {
		t.Fatal(err)
	} else if gotOwns != 1 {
		t.Fatalf("expected %d, got %d", 1, gotOwns)
	} else if ownsIRI.String() != samIRIString {
		t.Fatalf("expected %s, got %s", samIRIString, ownsIRI.String())
	} else if gotGet != 1 {
		t.Fatalf("expected %d, got %d", 1, gotGet)
	} else if getIRI.String() != samIRIString {
		t.Fatalf("expected %s, got %s", samIRIString, getIRI.String())
	} else if gotSet != 2 {
		t.Fatalf("expected %d, got %d", 2, gotSet)
	} else if err := PubObjectEquals(setObject, expectedActor); err != nil {
		t.Fatal(err)
	}
}

func TestPostInbox_Follow_DoesNotAddForAutoAcceptIfAlreadyPresent(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testFollow))))
	fedApp.onFollow = func(c context.Context, s *streams.Follow) FollowResponse {
		return AutomaticAccept
	}
	fedApp.newSigner = func() (httpsig.Signer, error) {
		s, _, err := httpsig.NewSigner([]httpsig.Algorithm{httpsig.RSA_SHA256}, nil, httpsig.Signature)
		if err != nil {
			t.Fatal(err)
		}
		return s, err
	}
	fedApp.privateKey = func(boxIRI *url.URL) (crypto.PrivateKey, string, error) {
		return testPrivateKey, testPublicKeyId, nil
	}
	fedCb.follow = func(c context.Context, s *streams.Follow) error {
		return nil
	}
	gotHttpDo := 0
	httpClient.do = func(req *http.Request) (*http.Response, error) {
		gotHttpDo++
		if gotHttpDo == 1 {
			actorResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer(sallyActorJSON)),
			}
			return actorResp, nil
		} else if gotHttpDo == 2 {
			okResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte{})),
			}
			return okResp, nil
		}
		return nil, nil
	}
	d.do = func(b []byte, u *url.URL, toDo func(b []byte, u *url.URL) error) {
		if err := toDo(b, u); err != nil {
			t.Fatalf("Unexpected error in MockDeliverer.Do: %s", err)
		}
	}
	app.MockFederateApp.owns = func(c context.Context, id *url.URL) bool {
		return true
	}
	app.MockFederateApp.get = func(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
		if rw != ReadWrite {
			t.Fatalf("expected RWType of %d, got %d", ReadWrite, rw)
		}
		followers := &vocab.Collection{}
		followers.AppendItemsIRI(sallyIRI)
		samActor := &vocab.Person{}
		samActor.SetInboxAnyURI(samIRIInbox)
		samActor.SetId(samIRI)
		samActor.SetFollowersCollection(followers)
		return samActor, nil
	}
	gotSet := 0
	var setObject PubObject
	app.MockFederateApp.set = func(c context.Context, o PubObject) error {
		gotSet++
		if gotSet == 1 {
			setObject = o
		}
		return nil
	}
	expectedFollowers := &vocab.Collection{}
	expectedFollowers.AppendItemsIRI(sallyIRI)
	expectedActor := &vocab.Person{}
	expectedActor.SetInboxAnyURI(samIRIInbox)
	expectedActor.SetId(samIRI)
	expectedActor.SetFollowersCollection(expectedFollowers)
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotSet != 2 {
		t.Fatalf("expected %d, got %d", 2, gotSet)
	} else if err := PubObjectEquals(setObject, expectedActor); err != nil {
		t.Fatal(err)
	}
}

func TestPostInbox_Follow_AutoAcceptFollowersIsOrderedCollection(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testFollow))))
	fedApp.onFollow = func(c context.Context, s *streams.Follow) FollowResponse {
		return AutomaticAccept
	}
	fedApp.newSigner = func() (httpsig.Signer, error) {
		s, _, err := httpsig.NewSigner([]httpsig.Algorithm{httpsig.RSA_SHA256}, nil, httpsig.Signature)
		if err != nil {
			t.Fatal(err)
		}
		return s, err
	}
	fedApp.privateKey = func(boxIRI *url.URL) (crypto.PrivateKey, string, error) {
		return testPrivateKey, testPublicKeyId, nil
	}
	fedCb.follow = func(c context.Context, s *streams.Follow) error {
		return nil
	}
	gotHttpDo := 0
	httpClient.do = func(req *http.Request) (*http.Response, error) {
		gotHttpDo++
		if gotHttpDo == 1 {
			actorResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer(sallyActorJSON)),
			}
			return actorResp, nil
		} else if gotHttpDo == 2 {
			okResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte{})),
			}
			return okResp, nil
		}
		return nil, nil
	}
	d.do = func(b []byte, u *url.URL, toDo func(b []byte, u *url.URL) error) {
		if err := toDo(b, u); err != nil {
			t.Fatalf("Unexpected error in MockDeliverer.Do: %s", err)
		}
	}
	app.MockFederateApp.owns = func(c context.Context, id *url.URL) bool {
		return true
	}
	app.MockFederateApp.get = func(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
		if rw != ReadWrite {
			t.Fatalf("expected RWType of %d, got %d", ReadWrite, rw)
		}
		samActor := &vocab.Person{}
		samActor.SetInboxAnyURI(samIRIInbox)
		samActor.SetId(samIRI)
		samActor.SetFollowersOrderedCollection(&vocab.OrderedCollection{})
		return samActor, nil
	}
	gotSet := 0
	var setObject PubObject
	app.MockFederateApp.set = func(c context.Context, o PubObject) error {
		gotSet++
		if gotSet == 1 {
			setObject = o
		}
		return nil
	}
	handled, err := p.PostInbox(context.Background(), resp, req)
	expectedFollowers := &vocab.OrderedCollection{}
	expectedFollowers.AppendOrderedItemsIRI(sallyIRI)
	expectedActor := &vocab.Person{}
	expectedActor.SetInboxAnyURI(samIRIInbox)
	expectedActor.SetId(samIRI)
	expectedActor.SetFollowersOrderedCollection(expectedFollowers)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if err := PubObjectEquals(setObject, expectedActor); err != nil {
		t.Fatal(err)
	}
}

func TestPostInbox_Follow_AutoAcceptFollowersIsIRI(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testFollow))))
	fedApp.onFollow = func(c context.Context, s *streams.Follow) FollowResponse {
		return AutomaticAccept
	}
	fedApp.newSigner = func() (httpsig.Signer, error) {
		s, _, err := httpsig.NewSigner([]httpsig.Algorithm{httpsig.RSA_SHA256}, nil, httpsig.Signature)
		if err != nil {
			t.Fatal(err)
		}
		return s, err
	}
	fedApp.privateKey = func(boxIRI *url.URL) (crypto.PrivateKey, string, error) {
		return testPrivateKey, testPublicKeyId, nil
	}
	fedCb.follow = func(c context.Context, s *streams.Follow) error {
		return nil
	}
	gotHttpDo := 0
	httpClient.do = func(req *http.Request) (*http.Response, error) {
		gotHttpDo++
		if gotHttpDo == 1 {
			actorResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer(sallyActorJSON)),
			}
			return actorResp, nil
		} else if gotHttpDo == 2 {
			okResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte{})),
			}
			return okResp, nil
		}
		return nil, nil
	}
	d.do = func(b []byte, u *url.URL, toDo func(b []byte, u *url.URL) error) {
		if err := toDo(b, u); err != nil {
			t.Fatalf("Unexpected error in MockDeliverer.Do: %s", err)
		}
	}
	app.MockFederateApp.owns = func(c context.Context, id *url.URL) bool {
		return true
	}
	app.MockFederateApp.get = func(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
		if rw != ReadWrite {
			t.Fatalf("expected RWType of %d, got %d", ReadWrite, rw)
		}
		if *id == *samIRI {
			samActor := &vocab.Person{}
			samActor.SetInboxAnyURI(samIRIInbox)
			samActor.SetId(samIRI)
			samActor.SetFollowersAnyURI(testNewIRI)
			return samActor, nil
		} else if *id == *testNewIRI {
			return &vocab.Collection{}, nil
		}
		t.Fatalf("unexpected get(%s)", id)
		return nil, nil
	}
	gotSet := 0
	var setObject PubObject
	app.MockFederateApp.set = func(c context.Context, o PubObject) error {
		gotSet++
		if gotSet == 1 {
			setObject = o
		}
		return nil
	}
	handled, err := p.PostInbox(context.Background(), resp, req)
	expectedFollowers := &vocab.Collection{}
	expectedFollowers.AppendItemsIRI(sallyIRI)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if err := PubObjectEquals(setObject, expectedFollowers); err != nil {
		t.Fatal(err)
	}
}

func TestPostInbox_Follow_DoesNotAutoAcceptIfNotOwned(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testFollow))))
	fedApp.onFollow = func(c context.Context, s *streams.Follow) FollowResponse {
		return AutomaticAccept
	}
	fedCb.follow = func(c context.Context, s *streams.Follow) error {
		return nil
	}
	app.MockFederateApp.owns = func(c context.Context, id *url.URL) bool {
		return false
	}
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	}
}

func TestPostInbox_Follow_DoesNotAutoRejectIfNotOwned(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testFollow))))
	fedApp.onFollow = func(c context.Context, s *streams.Follow) FollowResponse {
		return AutomaticReject
	}
	fedCb.follow = func(c context.Context, s *streams.Follow) error {
		return nil
	}
	app.MockFederateApp.owns = func(c context.Context, id *url.URL) bool {
		return false
	}
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	}
}

func TestPostInbox_Follow_CallsCallback(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testFollow))))
	fedApp.onFollow = func(c context.Context, s *streams.Follow) FollowResponse {
		return DoNothing
	}
	gotCallback := 0
	var gotCallbackObject *streams.Follow
	fedCb.follow = func(c context.Context, s *streams.Follow) error {
		gotCallback++
		gotCallbackObject = s
		return nil
	}
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotCallback != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCallback)
	} else if err := PubObjectEquals(gotCallbackObject.Raw(), testFollow); err != nil {
		t.Fatalf("unexpected callback object: %s", err)
	}
}

func TestPostInbox_Accept_DoesNothingIfNotAcceptingFollow(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testAcceptNote))))
	fedCb.accept = func(c context.Context, s *streams.Accept) error {
		return nil
	}
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	}
}

func TestPostInbox_Accept_AcceptFollowAddsToFollowersIfOwned(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testAcceptFollow))))
	gotOwns := 0
	var ownsIRI *url.URL
	app.MockFederateApp.owns = func(c context.Context, id *url.URL) bool {
		gotOwns++
		ownsIRI = id
		return true
	}
	gotGet := 0
	var getIRI *url.URL
	app.MockFederateApp.get = func(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
		if rw != ReadWrite {
			t.Fatalf("expected RWType of %d, got %d", ReadWrite, rw)
		}
		gotGet++
		getIRI = id
		sallyActor := &vocab.Person{}
		sallyActor.SetInboxAnyURI(sallyIRIInbox)
		sallyActor.SetId(sallyIRI)
		sallyActor.SetFollowingCollection(&vocab.Collection{})
		return sallyActor, nil
	}
	gotSet := 0
	var setObject PubObject
	app.MockFederateApp.set = func(c context.Context, o PubObject) error {
		gotSet++
		if gotSet == 1 {
			setObject = o
		}
		return nil
	}
	fedCb.accept = func(c context.Context, s *streams.Accept) error {
		return nil
	}
	expectedFollowing := &vocab.Collection{}
	expectedFollowing.AppendItemsIRI(samIRI)
	expectedActor := &vocab.Person{}
	expectedActor.SetInboxAnyURI(sallyIRIInbox)
	expectedActor.SetId(sallyIRI)
	expectedActor.SetFollowingCollection(expectedFollowing)
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotOwns != 1 {
		t.Fatalf("expected %d, got %d", 1, gotOwns)
	} else if ownsIRI.String() != sallyIRIString {
		t.Fatalf("expected %s, got %s", sallyIRIString, ownsIRI.String())
	} else if gotGet != 1 {
		t.Fatalf("expected %d, got %d", 1, gotGet)
	} else if getIRI.String() != sallyIRIString {
		t.Fatalf("expected %s, got %s", sallyIRIString, getIRI.String())
	} else if gotSet != 2 {
		t.Fatalf("expected %d, got %d", 2, gotSet)
	} else if err := PubObjectEquals(setObject, expectedActor); err != nil {
		t.Fatal(err)
	}
}

func TestPostInbox_Accept_AcceptFollowAddsToDefaultFollowersOrderedCollection(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testAcceptFollow))))
	gotOwns := 0
	var ownsIRI *url.URL
	app.MockFederateApp.owns = func(c context.Context, id *url.URL) bool {
		gotOwns++
		ownsIRI = id
		return true
	}
	gotGet := 0
	var getIRI *url.URL
	app.MockFederateApp.get = func(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
		if rw != ReadWrite {
			t.Fatalf("expected RWType of %d, got %d", ReadWrite, rw)
		}
		gotGet++
		getIRI = id
		sallyActor := &vocab.Person{}
		sallyActor.SetInboxAnyURI(sallyIRIInbox)
		sallyActor.SetId(sallyIRI)
		return sallyActor, nil
	}
	gotSet := 0
	var setObject PubObject
	app.MockFederateApp.set = func(c context.Context, o PubObject) error {
		gotSet++
		if gotSet == 1 {
			setObject = o
		}
		return nil
	}
	fedCb.accept = func(c context.Context, s *streams.Accept) error {
		return nil
	}
	expectedFollowing := &vocab.OrderedCollection{}
	expectedFollowing.AppendOrderedItemsIRI(samIRI)
	expectedActor := &vocab.Person{}
	expectedActor.SetInboxAnyURI(sallyIRIInbox)
	expectedActor.SetId(sallyIRI)
	expectedActor.SetFollowingOrderedCollection(expectedFollowing)
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotOwns != 1 {
		t.Fatalf("expected %d, got %d", 1, gotOwns)
	} else if ownsIRI.String() != sallyIRIString {
		t.Fatalf("expected %s, got %s", sallyIRIString, ownsIRI.String())
	} else if gotGet != 1 {
		t.Fatalf("expected %d, got %d", 1, gotGet)
	} else if getIRI.String() != sallyIRIString {
		t.Fatalf("expected %s, got %s", sallyIRIString, getIRI.String())
	} else if gotSet != 2 {
		t.Fatalf("expected %d, got %d", 2, gotSet)
	} else if err := PubObjectEquals(setObject, expectedActor); err != nil {
		t.Fatal(err)
	}
}

func TestPostInbox_Accept_AcceptFollowDoesNotAddIfAlreadyInCollection(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testAcceptFollow))))
	app.MockFederateApp.owns = func(c context.Context, id *url.URL) bool {
		return true
	}
	app.MockFederateApp.get = func(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
		following := &vocab.Collection{}
		following.AppendItemsIRI(samIRI)
		sallyActor := &vocab.Person{}
		sallyActor.SetInboxAnyURI(sallyIRIInbox)
		sallyActor.SetId(sallyIRI)
		sallyActor.SetFollowingCollection(following)
		return sallyActor, nil
	}
	gotSet := 0
	var setObject PubObject
	app.MockFederateApp.set = func(c context.Context, o PubObject) error {
		gotSet++
		if gotSet == 1 {
			setObject = o
		}
		return nil
	}
	fedCb.accept = func(c context.Context, s *streams.Accept) error {
		return nil
	}
	expectedFollowing := &vocab.Collection{}
	expectedFollowing.AppendItemsIRI(samIRI)
	expectedActor := &vocab.Person{}
	expectedActor.SetInboxAnyURI(sallyIRIInbox)
	expectedActor.SetId(sallyIRI)
	expectedActor.SetFollowingCollection(expectedFollowing)
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotSet != 2 {
		t.Fatalf("expected %d, got %d", 2, gotSet)
	} else if err := PubObjectEquals(setObject, expectedActor); err != nil {
		t.Fatal(err)
	}
}

func TestPostInbox_Accept_AcceptFollowAddsToFollowersOrderedCollection(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testAcceptFollow))))
	app.MockFederateApp.owns = func(c context.Context, id *url.URL) bool {
		return true
	}
	app.MockFederateApp.get = func(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
		if rw != ReadWrite {
			t.Fatalf("expected RWType of %d, got %d", ReadWrite, rw)
		}
		sallyActor := &vocab.Person{}
		sallyActor.SetInboxAnyURI(sallyIRIInbox)
		sallyActor.SetId(sallyIRI)
		sallyActor.SetFollowingOrderedCollection(&vocab.OrderedCollection{})
		return sallyActor, nil
	}
	gotSet := 0
	var setObject PubObject
	app.MockFederateApp.set = func(c context.Context, o PubObject) error {
		gotSet++
		if gotSet == 1 {
			setObject = o
		}
		return nil
	}
	fedCb.accept = func(c context.Context, s *streams.Accept) error {
		return nil
	}
	expectedFollowing := &vocab.OrderedCollection{}
	expectedFollowing.AppendOrderedItemsIRI(samIRI)
	expectedActor := &vocab.Person{}
	expectedActor.SetInboxAnyURI(sallyIRIInbox)
	expectedActor.SetId(sallyIRI)
	expectedActor.SetFollowingOrderedCollection(expectedFollowing)
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if err := PubObjectEquals(setObject, expectedActor); err != nil {
		t.Fatal(err)
	}
}

func TestPostInbox_Accept_AcceptFollowAddsToFollowersIRI(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testAcceptFollow))))
	app.MockFederateApp.owns = func(c context.Context, id *url.URL) bool {
		return true
	}
	app.MockFederateApp.get = func(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
		if rw != ReadWrite {
			t.Fatalf("expected RWType of %d, got %d", ReadWrite, rw)
		}
		if *id == *sallyIRI {
			sallyActor := &vocab.Person{}
			sallyActor.SetInboxAnyURI(sallyIRIInbox)
			sallyActor.SetId(sallyIRI)
			sallyActor.SetFollowingAnyURI(sallyFollowingIRI)
			return sallyActor, nil
		} else if *id == *sallyFollowingIRI {
			return &vocab.OrderedCollection{}, nil
		}
		t.Fatalf("unexpected get(%s)", id.String())
		return nil, nil
	}
	gotSet := 0
	var setObject PubObject
	app.MockFederateApp.set = func(c context.Context, o PubObject) error {
		gotSet++
		if gotSet == 1 {
			setObject = o
		}
		return nil
	}
	fedCb.accept = func(c context.Context, s *streams.Accept) error {
		return nil
	}
	expectedFollowing := &vocab.OrderedCollection{}
	expectedFollowing.AppendOrderedItemsIRI(samIRI)
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if err := PubObjectEquals(setObject, expectedFollowing); err != nil {
		t.Fatal(err)
	}
}

func TestPostInbox_Accept_DoesNothingIfNotOwned(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testAcceptFollow))))
	gotOwns := 0
	var ownsIRI *url.URL
	app.MockFederateApp.owns = func(c context.Context, id *url.URL) bool {
		gotOwns++
		ownsIRI = id
		return false
	}
	fedCb.accept = func(c context.Context, s *streams.Accept) error {
		return nil
	}
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotOwns != 1 {
		t.Fatalf("expected %d, got %d", 1, gotOwns)
	} else if ownsIRI.String() != sallyIRIString {
		t.Fatalf("expected %s, got %s", sallyIRIString, ownsIRI.String())
	}
}

func TestPostInbox_Accept_CallsCallback(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testAcceptFollow))))
	app.MockFederateApp.owns = func(c context.Context, id *url.URL) bool {
		return true
	}
	app.MockFederateApp.get = func(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
		sallyActor := &vocab.Person{}
		sallyActor.SetInboxAnyURI(sallyIRIInbox)
		sallyActor.SetId(sallyIRI)
		sallyActor.SetFollowingCollection(&vocab.Collection{})
		return sallyActor, nil
	}
	gotCallback := 0
	var gotCallbackObject *streams.Accept
	fedCb.accept = func(c context.Context, s *streams.Accept) error {
		gotCallback++
		gotCallbackObject = s
		return nil
	}
	expectedFollowing := &vocab.Collection{}
	expectedFollowing.AppendItemsObject(samActor)
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotCallback != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCallback)
	} else if err := PubObjectEquals(gotCallbackObject.Raw(), testAcceptFollow); err != nil {
		t.Fatalf("unexpected callback object: %s", err)
	}
}

func TestPostInbox_Reject_CallsCallback(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testRejectFollow))))
	gotCallback := 0
	var gotCallbackObject *streams.Reject
	fedCb.reject = func(c context.Context, s *streams.Reject) error {
		gotCallback++
		gotCallbackObject = s
		return nil
	}
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotCallback != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCallback)
	} else if err := PubObjectEquals(gotCallbackObject.Raw(), testRejectFollow); err != nil {
		t.Fatalf("unexpected callback object: %s", err)
	}
}

func TestPostInbox_Add_DoesNotAddIfTargetNotOwned(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testAddNote))))
	gotOwns := 0
	var gotOwnsId *url.URL
	app.MockFederateApp.owns = func(c context.Context, id *url.URL) bool {
		gotOwns++
		gotOwnsId = id
		return false
	}
	fedCb.add = func(c context.Context, s *streams.Add) error {
		return nil
	}
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotOwns != 1 {
		t.Fatalf("expected %d, got %d", 1, gotOwns)
	} else if gotOwnsId.String() != iriString {
		t.Fatalf("expected %s, got %s", iriString, gotOwnsId.String())
	}
}

func TestPostInbox_Add_AddIfTargetOwnedAndAppCanAdd(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testAddNote))))
	gotOwns := 0
	var gotOwnsId *url.URL
	app.MockFederateApp.owns = func(c context.Context, id *url.URL) bool {
		gotOwns++
		gotOwnsId = id
		return true
	}
	gotGet := 0
	var gotGetId *url.URL
	app.MockFederateApp.get = func(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
		if rw != ReadWrite {
			t.Fatalf("expected RWType of %d, got %d", ReadWrite, rw)
		}
		gotGet++
		gotGetId = id
		v := &vocab.Collection{}
		return v, nil
	}
	gotCanAdd := 0
	var gotCanAddObject vocab.ObjectType
	var gotCanAddTarget vocab.ObjectType
	app.MockFederateApp.canAdd = func(c context.Context, o vocab.ObjectType, t vocab.ObjectType) bool {
		gotCanAdd++
		gotCanAddObject = o
		gotCanAddTarget = t
		return true
	}
	gotSet := 0
	var gotSetTarget PubObject
	app.MockFederateApp.set = func(c context.Context, target PubObject) error {
		gotSet++
		if gotSet == 1 {
			gotSetTarget = target
		}
		return nil
	}
	fedCb.add = func(c context.Context, s *streams.Add) error {
		return nil
	}
	handled, err := p.PostInbox(context.Background(), resp, req)
	expected := &vocab.Collection{}
	expected.AppendItemsIRI(noteIRI)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotOwns != 1 {
		t.Fatalf("expected %d, got %d", 1, gotOwns)
	} else if gotOwnsId.String() != iriString {
		t.Fatalf("expected %s, got %s", iriString, gotOwnsId.String())
	} else if gotGet != 1 {
		t.Fatalf("expected %d, got %d", 1, gotGet)
	} else if gotGetId.String() != iriString {
		t.Fatalf("expected %s, got %s", iriString, gotGetId.String())
	} else if gotCanAdd != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCanAdd)
	} else if err := VocabSerializerEquals(gotCanAddObject, testNote); err != nil {
		t.Fatal(err)
	} else if err := VocabSerializerEquals(gotCanAddTarget, expected); err != nil {
		t.Fatal(err)
	} else if gotSet != 2 {
		t.Fatalf("expected %d, got %d", 2, gotSet)
	} else if err := PubObjectEquals(gotSetTarget, expected); err != nil {
		t.Fatal(err)
	}
}

func TestPostInbox_Add_DoesNotAddIfAppCannotAdd(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testAddNote))))
	gotOwns := 0
	var gotOwnsId *url.URL
	app.MockFederateApp.owns = func(c context.Context, id *url.URL) bool {
		gotOwns++
		gotOwnsId = id
		return true
	}
	gotGet := 0
	var gotGetId *url.URL
	app.MockFederateApp.get = func(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
		gotGet++
		gotGetId = id
		v := &vocab.Collection{}
		return v, nil
	}
	gotCanAdd := 0
	var gotCanAddObject vocab.ObjectType
	var gotCanAddTarget vocab.ObjectType
	app.MockFederateApp.canAdd = func(c context.Context, o vocab.ObjectType, t vocab.ObjectType) bool {
		gotCanAdd++
		gotCanAddObject = o
		gotCanAddTarget = t
		return false
	}
	fedCb.add = func(c context.Context, s *streams.Add) error {
		return nil
	}
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotOwns != 1 {
		t.Fatalf("expected %d, got %d", 1, gotOwns)
	} else if gotOwnsId.String() != iriString {
		t.Fatalf("expected %s, got %s", iriString, gotOwnsId.String())
	} else if gotGet != 1 {
		t.Fatalf("expected %d, got %d", 1, gotGet)
	} else if gotGetId.String() != iriString {
		t.Fatalf("expected %s, got %s", iriString, gotGetId.String())
	} else if gotCanAdd != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCanAdd)
	} else if err := VocabSerializerEquals(gotCanAddObject, testNote); err != nil {
		t.Fatal(err)
	} else if err := VocabSerializerEquals(gotCanAddTarget, &vocab.Collection{}); err != nil {
		t.Fatal(err)
	}
}

func TestPostInbox_Add_CallsCallback(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testAddNote))))
	app.MockFederateApp.owns = func(c context.Context, id *url.URL) bool {
		return true
	}
	app.MockFederateApp.get = func(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
		v := &vocab.Collection{}
		return v, nil
	}
	app.MockFederateApp.canAdd = func(c context.Context, o vocab.ObjectType, t vocab.ObjectType) bool {
		return true
	}
	gotCallback := 0
	var gotCallbackObject *streams.Add
	fedCb.add = func(c context.Context, s *streams.Add) error {
		gotCallback++
		gotCallbackObject = s
		return nil
	}
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotCallback != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCallback)
	} else if err := PubObjectEquals(gotCallbackObject.Raw(), testAddNote); err != nil {
		t.Fatalf("unexpected callback object: %s", err)
	}
}

func TestPostInbox_Remove_DoesNotRemoveIfTargetNotOwned(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testRemoveNote))))
	gotOwns := 0
	var gotOwnsId *url.URL
	app.MockFederateApp.owns = func(c context.Context, id *url.URL) bool {
		gotOwns++
		gotOwnsId = id
		return false
	}
	fedCb.remove = func(c context.Context, s *streams.Remove) error {
		return nil
	}
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotOwns != 1 {
		t.Fatalf("expected %d, got %d", 1, gotOwns)
	} else if gotOwnsId.String() != iriString {
		t.Fatalf("expected %s, got %s", iriString, gotOwnsId.String())
	}
}

func TestPostInbox_Remove_RemoveIfTargetOwnedAndCanRemove(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testRemoveNote))))
	gotOwns := 0
	var gotOwnsId *url.URL
	app.MockFederateApp.owns = func(c context.Context, id *url.URL) bool {
		gotOwns++
		gotOwnsId = id
		return true
	}
	gotGet := 0
	var gotGetId *url.URL
	app.MockFederateApp.get = func(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
		if rw != ReadWrite {
			t.Fatalf("expected RWType of %d, got %d", ReadWrite, rw)
		}
		gotGet++
		gotGetId = id
		v := &vocab.Collection{}
		v.AppendItemsObject(testNote)
		return v, nil
	}
	gotCanRemove := 0
	var gotCanRemoveObject vocab.ObjectType
	var gotCanRemoveTarget vocab.ObjectType
	app.MockFederateApp.canRemove = func(c context.Context, o vocab.ObjectType, t vocab.ObjectType) bool {
		gotCanRemove++
		gotCanRemoveObject = o
		gotCanRemoveTarget = t
		return true
	}
	gotSet := 0
	var gotSetTarget PubObject
	app.MockFederateApp.set = func(c context.Context, target PubObject) error {
		gotSet++
		if gotSet == 1 {
			gotSetTarget = target
		}
		return nil
	}
	fedCb.remove = func(c context.Context, s *streams.Remove) error {
		return nil
	}
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotOwns != 1 {
		t.Fatalf("expected %d, got %d", 1, gotOwns)
	} else if gotOwnsId.String() != iriString {
		t.Fatalf("expected %s, got %s", iriString, gotOwnsId.String())
	} else if gotGet != 1 {
		t.Fatalf("expected %d, got %d", 1, gotGet)
	} else if gotGetId.String() != iriString {
		t.Fatalf("expected %s, got %s", iriString, gotGetId.String())
	} else if gotCanRemove != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCanRemove)
	} else if err := VocabSerializerEquals(gotCanRemoveObject, testNote); err != nil {
		t.Fatal(err)
	} else if err := VocabSerializerEquals(gotCanRemoveTarget, &vocab.Collection{}); err != nil {
		t.Fatal(err)
	} else if gotSet != 2 {
		t.Fatalf("expected %d, got %d", 2, gotSet)
	} else if err := PubObjectEquals(gotSetTarget, &vocab.Collection{}); err != nil {
		t.Fatal(err)
	}
}

func TestPostInbox_Remove_DoesNotRemoveIfAppCannotRemove(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testRemoveNote))))
	gotOwns := 0
	var gotOwnsId *url.URL
	app.MockFederateApp.owns = func(c context.Context, id *url.URL) bool {
		gotOwns++
		gotOwnsId = id
		return true
	}
	gotGet := 0
	var gotGetId *url.URL
	app.MockFederateApp.get = func(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
		gotGet++
		gotGetId = id
		v := &vocab.Collection{}
		v.AppendItemsObject(testNote)
		return v, nil
	}
	gotCanRemove := 0
	var gotCanRemoveObject vocab.ObjectType
	var gotCanRemoveTarget vocab.ObjectType
	app.MockFederateApp.canRemove = func(c context.Context, o vocab.ObjectType, t vocab.ObjectType) bool {
		gotCanRemove++
		gotCanRemoveObject = o
		gotCanRemoveTarget = t
		return false
	}
	fedCb.remove = func(c context.Context, s *streams.Remove) error {
		return nil
	}
	handled, err := p.PostInbox(context.Background(), resp, req)
	expected := &vocab.Collection{}
	expected.AppendItemsObject(testNote)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotOwns != 1 {
		t.Fatalf("expected %d, got %d", 1, gotOwns)
	} else if gotOwnsId.String() != iriString {
		t.Fatalf("expected %s, got %s", iriString, gotOwnsId.String())
	} else if gotGet != 1 {
		t.Fatalf("expected %d, got %d", 1, gotGet)
	} else if gotGetId.String() != iriString {
		t.Fatalf("expected %s, got %s", iriString, gotGetId.String())
	} else if gotCanRemove != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCanRemove)
	} else if err := VocabSerializerEquals(gotCanRemoveObject, testNote); err != nil {
		t.Fatal(err)
	} else if err := VocabSerializerEquals(gotCanRemoveTarget, expected); err != nil {
		t.Fatal(err)
	}
}

func TestPostInbox_Remove_CallsCallback(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testRemoveNote))))
	app.MockFederateApp.owns = func(c context.Context, id *url.URL) bool {
		return true
	}
	app.MockFederateApp.get = func(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
		v := &vocab.Collection{}
		return v, nil
	}
	app.MockFederateApp.canRemove = func(c context.Context, o vocab.ObjectType, t vocab.ObjectType) bool {
		return true
	}
	gotCallback := 0
	var gotCallbackObject *streams.Remove
	fedCb.remove = func(c context.Context, s *streams.Remove) error {
		gotCallback++
		gotCallbackObject = s
		return nil
	}
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotCallback != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCallback)
	} else if err := PubObjectEquals(gotCallbackObject.Raw(), testRemoveNote); err != nil {
		t.Fatalf("unexpected callback object: %s", err)
	}
}

func TestPostInbox_Like_AddsToLikeCollection(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testLikeNote))))
	gotOwns := 0
	var gotOwnsId *url.URL
	app.MockFederateApp.owns = func(c context.Context, id *url.URL) bool {
		gotOwns++
		gotOwnsId = id
		return true
	}
	gotGet := 0
	var gotGetId *url.URL
	app.MockFederateApp.get = func(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
		if rw != ReadWrite {
			t.Fatalf("expected RWType of %d, got %d", ReadWrite, rw)
		}
		gotGet++
		gotGetId = id
		v := &vocab.Note{}
		v.SetId(noteIRI)
		v.AppendNameString(noteName)
		v.AppendContentString("This is a simple note")
		v.SetLikesCollection(&vocab.Collection{})
		return v, nil
	}
	gotSet := 0
	var gotSetObject PubObject
	app.MockFederateApp.set = func(c context.Context, target PubObject) error {
		gotSet++
		if gotSet == 1 {
			gotSetObject = target
		}
		return nil
	}
	fedCb.like = func(c context.Context, s *streams.Like) error {
		return nil
	}
	handled, err := p.PostInbox(context.Background(), resp, req)
	expected := &vocab.Collection{}
	expected.AppendItemsIRI(noteActivityIRI)
	expectedNote := &vocab.Note{}
	expectedNote.SetId(noteIRI)
	expectedNote.AppendNameString(noteName)
	expectedNote.AppendContentString("This is a simple note")
	expectedNote.SetLikesCollection(expected)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotOwns != 1 {
		t.Fatalf("expected %d, got %d", 1, gotOwns)
	} else if gotOwnsId.String() != noteURIString {
		t.Fatalf("expected %s, got %s", noteURIString, gotOwnsId.String())
	} else if gotGet != 1 {
		t.Fatalf("expected %d, got %d", 1, gotGet)
	} else if gotGetId.String() != noteURIString {
		t.Fatalf("expected %s, got %s", noteURIString, gotGetId.String())
	} else if gotSet != 2 {
		t.Fatalf("expected %d, got %d", 2, gotSet)
	} else if err := PubObjectEquals(gotSetObject, expectedNote); err != nil {
		t.Fatalf("unexpected callback object: %s", err)
	}
}

func TestPostInbox_Like_AddsToDefaultOrderedCollection(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testLikeNote))))
	gotOwns := 0
	var gotOwnsId *url.URL
	app.MockFederateApp.owns = func(c context.Context, id *url.URL) bool {
		gotOwns++
		gotOwnsId = id
		return true
	}
	gotGet := 0
	var gotGetId *url.URL
	app.MockFederateApp.get = func(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
		if rw != ReadWrite {
			t.Fatalf("expected RWType of %d, got %d", ReadWrite, rw)
		}
		gotGet++
		gotGetId = id
		v := &vocab.Note{}
		v.SetId(noteIRI)
		v.AppendNameString(noteName)
		v.AppendContentString("This is a simple note")
		return v, nil
	}
	gotSet := 0
	var gotSetObject PubObject
	app.MockFederateApp.set = func(c context.Context, target PubObject) error {
		gotSet++
		if gotSet == 1 {
			gotSetObject = target
		}
		return nil
	}
	fedCb.like = func(c context.Context, s *streams.Like) error {
		return nil
	}
	handled, err := p.PostInbox(context.Background(), resp, req)
	expected := &vocab.OrderedCollection{}
	expected.AppendOrderedItemsIRI(noteActivityIRI)
	expectedNote := &vocab.Note{}
	expectedNote.SetId(noteIRI)
	expectedNote.AppendNameString(noteName)
	expectedNote.AppendContentString("This is a simple note")
	expectedNote.SetLikesOrderedCollection(expected)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotOwns != 1 {
		t.Fatalf("expected %d, got %d", 1, gotOwns)
	} else if gotOwnsId.String() != noteURIString {
		t.Fatalf("expected %s, got %s", noteURIString, gotOwnsId.String())
	} else if gotGet != 1 {
		t.Fatalf("expected %d, got %d", 1, gotGet)
	} else if gotGetId.String() != noteURIString {
		t.Fatalf("expected %s, got %s", noteURIString, gotGetId.String())
	} else if gotSet != 2 {
		t.Fatalf("expected %d, got %d", 2, gotSet)
	} else if err := PubObjectEquals(gotSetObject, expectedNote); err != nil {
		t.Fatalf("unexpected callback object: %s", err)
	}
}

func TestPostInbox_Like_DoesNotAddLikeToCollectionIfAlreadyPresent(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testLikeNote))))
	app.MockFederateApp.owns = func(c context.Context, id *url.URL) bool {
		return true
	}
	app.MockFederateApp.get = func(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
		likes := &vocab.Collection{}
		likes.AppendItemsIRI(noteActivityIRI)
		v := &vocab.Note{}
		v.SetId(noteIRI)
		v.AppendNameString(noteName)
		v.AppendContentString("This is a simple note")
		v.SetLikesCollection(likes)
		return v, nil
	}
	gotSet := 0
	var gotSetObject PubObject
	app.MockFederateApp.set = func(c context.Context, target PubObject) error {
		gotSet++
		if gotSet == 1 {
			gotSetObject = target
		}
		return nil
	}
	fedCb.like = func(c context.Context, s *streams.Like) error {
		return nil
	}
	handled, err := p.PostInbox(context.Background(), resp, req)
	expected := &vocab.OrderedCollection{}
	expected.AppendOrderedItemsIRI(noteActivityIRI)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotSet != 1 {
		t.Fatalf("expected %d, got %d", 1, gotSet)
	} else if err := PubObjectEquals(gotSetObject, expected); err != nil {
		t.Fatalf("unexpected callback object: %s", err)
	}
}

func TestPostInbox_Like_AddsToLikeOrderedCollection(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testLikeNote))))
	app.MockFederateApp.owns = func(c context.Context, id *url.URL) bool {
		return true
	}
	app.MockFederateApp.get = func(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
		v := &vocab.Note{}
		v.SetId(noteIRI)
		v.AppendNameString(noteName)
		v.AppendContentString("This is a simple note")
		v.SetLikesOrderedCollection(&vocab.OrderedCollection{})
		return v, nil
	}
	gotSet := 0
	var gotSetObject PubObject
	app.MockFederateApp.set = func(c context.Context, target PubObject) error {
		gotSet++
		if gotSet == 1 {
			gotSetObject = target
		}
		return nil
	}
	fedCb.like = func(c context.Context, s *streams.Like) error {
		return nil
	}
	handled, err := p.PostInbox(context.Background(), resp, req)
	expected := &vocab.OrderedCollection{}
	expected.AppendOrderedItemsIRI(noteActivityIRI)
	expectedNote := &vocab.Note{}
	expectedNote.SetId(noteIRI)
	expectedNote.AppendNameString(noteName)
	expectedNote.AppendContentString("This is a simple note")
	expectedNote.SetLikesOrderedCollection(expected)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if err := PubObjectEquals(gotSetObject, expectedNote); err != nil {
		t.Fatalf("unexpected callback object: %s", err)
	}
}

func TestPostInbox_Like_AddsToLikeIRI(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testLikeNote))))
	app.MockFederateApp.owns = func(c context.Context, id *url.URL) bool {
		return true
	}
	app.MockFederateApp.get = func(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
		if *id == *noteIRI {
			v := &vocab.Note{}
			v.SetId(noteIRI)
			v.AppendNameString(noteName)
			v.AppendContentString("This is a simple note")
			v.SetLikesAnyURI(testNewIRI)
			return v, nil
		} else if *id == *testNewIRI {
			return &vocab.OrderedCollection{}, nil
		}
		t.Fatalf("unexpected get(%s)", id)
		return nil, nil
	}
	gotSet := 0
	var gotSetObject PubObject
	app.MockFederateApp.set = func(c context.Context, target PubObject) error {
		gotSet++
		if gotSet == 1 {
			gotSetObject = target
		}
		return nil
	}
	fedCb.like = func(c context.Context, s *streams.Like) error {
		return nil
	}
	handled, err := p.PostInbox(context.Background(), resp, req)
	expected := &vocab.OrderedCollection{}
	expected.AppendOrderedItemsIRI(noteActivityIRI)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if err := PubObjectEquals(gotSetObject, expected); err != nil {
		t.Fatalf("unexpected callback object: %s", err)
	}
}

func TestPostInbox_Like_CallsCallback(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testLikeNote))))
	app.MockFederateApp.owns = func(c context.Context, id *url.URL) bool {
		return true
	}
	app.MockFederateApp.get = func(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
		v := &vocab.Note{}
		v.SetId(noteIRI)
		v.AppendNameString(noteName)
		v.AppendContentString("This is a simple note")
		v.SetLikesCollection(&vocab.Collection{})
		return v, nil
	}
	gotCallback := 0
	var gotCallbackObject *streams.Like
	fedCb.like = func(c context.Context, s *streams.Like) error {
		gotCallback++
		gotCallbackObject = s
		return nil
	}
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotCallback != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCallback)
	} else if err := PubObjectEquals(gotCallbackObject.Raw(), testLikeNote); err != nil {
		t.Fatalf("unexpected callback object: %s", err)
	}
}

func TestPostInbox_Undo_CallsCallback(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostInboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testUndoLike))))
	gotCallback := 0
	var gotCallbackObject *streams.Undo
	fedCb.undo = func(c context.Context, s *streams.Undo) error {
		gotCallback++
		gotCallbackObject = s
		return nil
	}
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotCallback != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCallback)
	} else if err := PubObjectEquals(gotCallbackObject.Raw(), testUndoLike); err != nil {
		t.Fatalf("unexpected callback object: %s", err)
	}
}

func TestGetInbox_RejectNonActivityPub(t *testing.T) {
	app, _, _, _, _, _, _, p := NewPubberTest(t)
	resp := httptest.NewRecorder()
	req := httptest.NewRequest("GET", testInboxURI, nil)
	app.MockFederateApp.getInbox = func(c context.Context, r *http.Request, rw RWType) (vocab.OrderedCollectionType, error) {
		return &vocab.OrderedCollection{}, nil
	}
	handled, err := p.GetInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if handled {
		t.Fatalf("expected !handled, got handled")
	}
}

func TestGetInbox_SetsContentTypeHeader(t *testing.T) {
	app, _, _, _, _, _, _, p := NewPubberTest(t)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("GET", testInboxURI, nil))
	app.MockFederateApp.getInbox = func(c context.Context, r *http.Request, rw RWType) (vocab.OrderedCollectionType, error) {
		return &vocab.OrderedCollection{}, nil
	}
	handled, err := p.GetInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if l := len(resp.HeaderMap["Content-Type"]); l != 1 {
		t.Fatalf("expected %d, got %d", 1, l)
	} else if h := resp.HeaderMap["Content-Type"][0]; h != responseContentTypeHeader {
		t.Fatalf("expected %s, got %s", responseContentTypeHeader, h)
	}
}

func TestGetInbox_DeduplicateInboxItems(t *testing.T) {
	app, _, _, _, _, _, _, p := NewPubberTest(t)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("GET", testInboxURI, nil))
	app.MockFederateApp.getInbox = func(c context.Context, r *http.Request, rw RWType) (vocab.OrderedCollectionType, error) {
		v := &vocab.OrderedCollection{}
		v.AppendOrderedItemsObject(testCreateNote)
		v.AppendOrderedItemsObject(testCreateNote)
		v.AppendOrderedItemsObject(testUpdateNote)
		return v, nil
	}
	handled, err := p.GetInbox(context.Background(), resp, req)
	expected := &vocab.OrderedCollection{}
	expected.AppendOrderedItemsObject(testCreateNote)
	expected.AppendOrderedItemsObject(testUpdateNote)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if err := VocabEquals(resp.Body, expected); err != nil {
		t.Fatalf("unexpected callback object: %s", err)
	}
}

func TestPostOutbox_RejectNonActivityPub(t *testing.T) {
	_, _, _, _, _, _, _, p := NewPubberTest(t)
	resp := httptest.NewRecorder()
	req := httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testCreateNote)))
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if handled {
		t.Fatalf("expected !handled, got handled")
	}
}

func TestPostOutbox_RejectUnauthorized_NotSigned(t *testing.T) {
	_, socialApp, _, _, _, _, _, p := NewPubberTest(t)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testCreateNote))))
	socialApp.getSocialAPIVerifier = func(c context.Context) SocialAPIVerifier {
		return nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if resp.Code != http.StatusBadRequest {
		t.Fatalf("expected %d, got %d", http.StatusBadRequest, resp.Code)
	}
}

func TestPostOutbox_RejectUnauthenticatedUnauthorized(t *testing.T) {
	_, socialApp, _, _, _, _, _, p := NewPubberTest(t)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testCreateNote))))
	gotVerifyForOutbox := 0
	socialApp.getSocialAPIVerifier = func(c context.Context) SocialAPIVerifier {
		mockV := &MockSocialAPIVerifier{
			verifyForOutbox: func(r *http.Request, outbox *url.URL) (bool, bool, error) {
				gotVerifyForOutbox++
				return false, false, nil
			},
		}
		return mockV
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotVerifyForOutbox != 1 {
		t.Fatalf("expected %d, got %d", 1, gotVerifyForOutbox)
	} else if resp.Code != http.StatusBadRequest {
		t.Fatalf("expected %d, got %d", http.StatusBadRequest, resp.Code)
	}
}

func TestPostOutbox_RejectAuthenticatedUnauthorized(t *testing.T) {
	_, socialApp, _, _, _, _, _, p := NewPubberTest(t)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testCreateNote))))
	gotVerifyForOutbox := 0
	socialApp.getSocialAPIVerifier = func(c context.Context) SocialAPIVerifier {
		mockV := &MockSocialAPIVerifier{
			verifyForOutbox: func(r *http.Request, outbox *url.URL) (bool, bool, error) {
				gotVerifyForOutbox++
				return true, false, nil
			},
		}
		return mockV
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotVerifyForOutbox != 1 {
		t.Fatalf("expected %d, got %d", 1, gotVerifyForOutbox)
	} else if resp.Code != http.StatusForbidden {
		t.Fatalf("expected %d, got %d", http.StatusForbidden, resp.Code)
	}
}

func TestPostOutbox_RejectFallback_Unauthorized_NotSigned(t *testing.T) {
	_, socialApp, _, _, _, _, _, p := NewPubberTest(t)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testCreateNote))))
	gotVerifyForOutbox := 0
	socialApp.getSocialAPIVerifier = func(c context.Context) SocialAPIVerifier {
		mockV := &MockSocialAPIVerifier{
			verifyForOutbox: func(r *http.Request, outbox *url.URL) (bool, bool, error) {
				gotVerifyForOutbox++
				return false, true, nil
			},
		}
		return mockV
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotVerifyForOutbox != 1 {
		t.Fatalf("expected %d, got %d", 1, gotVerifyForOutbox)
	} else if resp.Code != http.StatusBadRequest {
		t.Fatalf("expected %d, got %d", http.StatusBadRequest, resp.Code)
	}
}

func TestPostOutbox_RejectFallback_Unauthorized_WrongSignature(t *testing.T) {
	_, socialApp, _, _, _, _, _, p := NewPubberTest(t)
	resp := httptest.NewRecorder()
	req := BadSignature(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testCreateNote)))))
	gotVerifyForOutbox := 0
	socialApp.getSocialAPIVerifier = func(c context.Context) SocialAPIVerifier {
		mockV := &MockSocialAPIVerifier{
			verifyForOutbox: func(r *http.Request, outbox *url.URL) (bool, bool, error) {
				gotVerifyForOutbox++
				return false, true, nil
			},
		}
		return mockV
	}
	gotPublicKey := 0
	var gotPublicKeyId string
	var gotBoxIRI *url.URL
	socialApp.getPublicKeyForOutbox = func(c context.Context, publicKeyId string, boxIRI *url.URL) (crypto.PublicKey, httpsig.Algorithm, error) {
		gotPublicKey++
		gotPublicKeyId = publicKeyId
		gotBoxIRI = boxIRI
		return testPrivateKey.Public(), httpsig.RSA_SHA256, nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotVerifyForOutbox != 1 {
		t.Fatalf("expected %d, got %d", 1, gotVerifyForOutbox)
	} else if gotPublicKey != 1 {
		t.Fatalf("expected %d, got %d", 1, gotPublicKey)
	} else if gotPublicKeyId != testPublicKeyId {
		t.Fatalf("expected %s, got %s", testPublicKeyId, gotPublicKeyId)
	} else if s := gotBoxIRI.String(); s != testOutboxURI {
		t.Fatalf("expected %s, got %s", testOutboxURI, s)
	} else if resp.Code != http.StatusForbidden {
		t.Fatalf("expected %d, got %d", http.StatusForbidden, resp.Code)
	}
}

func TestPostOutbox_RejectUnauthorized_WrongSignature(t *testing.T) {
	_, socialApp, _, _, _, _, _, p := NewPubberTest(t)
	resp := httptest.NewRecorder()
	req := BadSignature(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testCreateNote)))))
	gotSocialAPIVerifier := 0
	socialApp.getSocialAPIVerifier = func(c context.Context) SocialAPIVerifier {
		gotSocialAPIVerifier++
		return nil
	}
	gotPublicKey := 0
	var gotPublicKeyId string
	var gotBoxIRI *url.URL
	socialApp.getPublicKeyForOutbox = func(c context.Context, publicKeyId string, boxIRI *url.URL) (crypto.PublicKey, httpsig.Algorithm, error) {
		gotPublicKey++
		gotPublicKeyId = publicKeyId
		gotBoxIRI = boxIRI
		return testPrivateKey.Public(), httpsig.RSA_SHA256, nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotSocialAPIVerifier != 1 {
		t.Fatalf("expected %d, got %d", 1, gotSocialAPIVerifier)
	} else if gotPublicKey != 1 {
		t.Fatalf("expected %d, got %d", 1, gotPublicKey)
	} else if gotPublicKeyId != testPublicKeyId {
		t.Fatalf("expected %s, got %s", testPublicKeyId, gotPublicKeyId)
	} else if s := gotBoxIRI.String(); s != testOutboxURI {
		t.Fatalf("expected %s, got %s", testOutboxURI, s)
	} else if resp.Code != http.StatusForbidden {
		t.Fatalf("expected %d, got %d", http.StatusForbidden, resp.Code)
	}
}

func TestPostOutbox_WrapInCreateActivity(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	// Raw Note
	rawNote := &vocab.Note{}
	rawNote.SetId(noteIRI)
	rawNote.AppendNameString(noteName)
	rawNote.AppendContentString("This is a simple note")
	rawNote.AppendToObject(samActor)
	// Expected result
	expectedNote := &vocab.Note{}
	expectedNote.SetId(testNewIRI2)
	expectedNote.AppendNameString(noteName)
	expectedNote.AppendContentString("This is a simple note")
	expectedNote.AppendToObject(samActor)
	expectedNote.AppendAttributedToIRI(sallyIRI)
	expectedCreate := &vocab.Create{}
	expectedCreate.SetId(testNewIRI)
	expectedCreate.AppendActorIRI(sallyIRI)
	expectedCreate.AppendObject(expectedNote)
	expectedCreate.AppendToIRI(samIRI)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(rawNote)))))
	gotActorIRI := 0
	socialApp.actorIRI = func(c context.Context, r *http.Request) (*url.URL, error) {
		gotActorIRI++
		return sallyIRI, nil
	}
	gotCallback := 0
	var gotCallbackObject *streams.Create
	socialCb.create = func(c context.Context, s *streams.Create) error {
		gotCallback++
		gotCallbackObject = s
		return nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotActorIRI != 1 {
		t.Fatalf("expected %d, got %d", 1, gotActorIRI)
	} else if gotCallback != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCallback)
	} else if err := PubObjectEquals(gotCallbackObject.Raw(), expectedCreate); err != nil {
		t.Fatalf("unexpected callback object: %s", err)
	}
}

func TestPostOutbox_RequiresObject(t *testing.T) {
	tests := []struct {
		name  string
		input func() vocab.Serializer
	}{
		{
			name: "create",
			input: func() vocab.Serializer {
				v := &vocab.Create{}
				v.SetId(noteActivityIRI)
				v.AppendSummaryString("Sally created a note")
				v.AppendActorObject(sallyActor)
				v.AppendToObject(samActor)
				return v
			},
		},
		{
			name: "update",
			input: func() vocab.Serializer {
				v := &vocab.Update{}
				v.SetId(noteActivityIRI)
				v.AppendSummaryString("Sally updated a note")
				v.AppendActorObject(sallyActor)
				v.AppendToObject(samActor)
				return v
			},
		},
		{
			name: "delete",
			input: func() vocab.Serializer {
				v := &vocab.Delete{}
				v.SetId(noteActivityIRI)
				v.AppendActorObject(sallyActor)
				v.AppendToObject(samActor)
				return v
			},
		},
		{
			name: "follow",
			input: func() vocab.Serializer {
				v := &vocab.Follow{}
				v.SetId(noteActivityIRI)
				v.AppendActorObject(sallyActor)
				v.AppendToObject(samActor)
				return v
			},
		},
		{
			name: "add",
			input: func() vocab.Serializer {
				v := &vocab.Add{}
				v.SetId(noteActivityIRI)
				v.AppendActorObject(sallyActor)
				v.AppendToObject(samActor)
				v.AppendTargetObject(testNote)
				return v
			},
		},
		{
			name: "remove",
			input: func() vocab.Serializer {
				v := &vocab.Remove{}
				v.SetId(noteActivityIRI)
				v.AppendActorObject(sallyActor)
				v.AppendToObject(samActor)
				v.AppendTargetObject(testNote)
				return v
			},
		},
		{
			name: "like",
			input: func() vocab.Serializer {
				v := &vocab.Like{}
				v.SetId(noteActivityIRI)
				v.AppendActorObject(sallyActor)
				v.AppendToObject(samActor)
				return v
			},
		},
		{
			name: "block",
			input: func() vocab.Serializer {
				v := &vocab.Block{}
				v.SetId(noteActivityIRI)
				v.AppendActorObject(sallyActor)
				v.AppendToObject(samActor)
				return v
			},
		},
		{
			name: "undo",
			input: func() vocab.Serializer {
				v := &vocab.Undo{}
				v.SetId(noteActivityIRI)
				v.AppendActorObject(sallyActor)
				v.AppendToObject(samActor)
				return v
			},
		},
	}
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	for _, test := range tests {
		t.Logf("Running table test case %q", test.name)
		resp := httptest.NewRecorder()
		req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(test.input())))))
		handled, err := p.PostOutbox(context.Background(), resp, req)
		if resp.Code != http.StatusBadRequest {
			t.Fatalf("(%s) expected %d, got %d", test.name, http.StatusBadRequest, resp.Code)
		} else if err != nil {
			t.Fatal(err)
		} else if !handled {
			t.Fatalf("(%s) expected handled, got !handled", test.name)
		}
	}
}

func TestPostOutbox_RequiresTarget(t *testing.T) {
	tests := []struct {
		name  string
		input func() vocab.Serializer
	}{
		{
			name: "add",
			input: func() vocab.Serializer {
				v := &vocab.Add{}
				v.SetId(noteActivityIRI)
				v.AppendActorObject(sallyActor)
				v.AppendToObject(samActor)
				v.AppendObject(testNote)
				return v
			},
		},
		{
			name: "remove",
			input: func() vocab.Serializer {
				v := &vocab.Remove{}
				v.SetId(noteActivityIRI)
				v.AppendActorObject(sallyActor)
				v.AppendToObject(samActor)
				v.AppendObject(testNote)
				return v
			},
		},
	}
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	for _, test := range tests {
		t.Logf("Running table test case %q", test.name)
		resp := httptest.NewRecorder()
		req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(test.input())))))
		handled, err := p.PostOutbox(context.Background(), resp, req)
		if resp.Code != http.StatusBadRequest {
			t.Fatalf("(%s) expected %d, got %d", test.name, http.StatusBadRequest, resp.Code)
		} else if err != nil {
			t.Fatal(err)
		} else if !handled {
			t.Fatalf("(%s) expected handled, got !handled", test.name)
		}
	}
}

func TestPostOutbox_Create_CopyToAttributedTo(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testCreateNote)))))
	gotCallback := 0
	var gotCallbackObject *streams.Create
	socialCb.create = func(c context.Context, s *streams.Create) error {
		gotCallback++
		gotCallbackObject = s
		return nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotCallback != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCallback)
	} else if e := PubObjectEquals(gotCallbackObject.Raw(), testClientExpectedCreateNote); e != nil {
		t.Fatal(e)
	}
}

func TestPostOutbox_Create_SetCreatedObject(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testCreateNote)))))
	socialCb.create = func(c context.Context, s *streams.Create) error {
		return nil
	}
	gotSet := 0
	var gotSetOutbox PubObject
	var gotSetActivity PubObject
	var gotSetCreate PubObject
	app.MockFederateApp.set = func(c context.Context, o PubObject) error {
		gotSet++
		if gotSet == 1 {
			gotSetCreate = o
		} else if gotSet == 2 {
			gotSetActivity = o
		} else if gotSet == 3 {
			gotSetOutbox = o
		}
		return nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	expectedOutbox := &vocab.OrderedCollection{}
	expectedOutbox.AppendType("OrderedCollection")
	expectedOutbox.AppendOrderedItemsIRI(testClientExpectedCreateNote.GetId())
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotSet != 3 {
		t.Fatalf("expected %d, got %d", 3, gotSet)
	} else if err := PubObjectEquals(gotSetCreate, testClientExpectedNote); err != nil {
		t.Fatalf("unexpected callback object: %s", err)
	} else if err := PubObjectEquals(gotSetActivity, testClientExpectedCreateNote); err != nil {
		t.Fatalf("unexpected callback object: %s", err)
	} else if err := PubObjectEquals(gotSetOutbox, expectedOutbox); err != nil {
		t.Fatalf("unexpected callback object: %s", err)
	}
}

func TestPostOutbox_Create_CallsCallback(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testCreateNote)))))
	gotCallback := 0
	var gotCallbackObject *streams.Create
	socialCb.create = func(c context.Context, s *streams.Create) error {
		gotCallback++
		gotCallbackObject = s
		return nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotCallback != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCallback)
	} else if err := PubObjectEquals(gotCallbackObject.Raw(), testClientExpectedCreateNote); err != nil {
		t.Fatalf("unexpected callback object: %s", err)
	}
}

func TestPostOutbox_Create_IsDelivered(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testCreateNote)))))
	socialCb.create = func(c context.Context, s *streams.Create) error {
		return nil
	}
	gotHttpDo := 0
	var httpDeliveryRequest *http.Request
	httpClient.do = func(req *http.Request) (*http.Response, error) {
		gotHttpDo++
		if gotHttpDo == 1 {
			actorResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer(samActorJSON)),
			}
			return actorResp, nil
		} else if gotHttpDo == 2 {
			actorResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer(sallyActorJSON)),
			}
			return actorResp, nil
		} else if gotHttpDo == 3 {
			httpDeliveryRequest = req
			okResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte{})),
			}
			return okResp, nil
		}
		return nil, nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotHttpDo != 3 {
		t.Fatalf("expected %d, got %d", 3, gotHttpDo)
	} else if httpDeliveryRequest.Method != "POST" {
		t.Fatalf("expected %s, got %s", "POST", httpDeliveryRequest.Method)
	} else if s := httpDeliveryRequest.URL.String(); s != samIRIInboxString {
		t.Fatalf("expected %s, got %s", samIRIInboxString, s)
	}
}

func TestPostOutbox_Update_DeleteSubFields(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer([]byte(testDeleteSubFields)))))
	socialCb.update = func(c context.Context, s *streams.Update) error {
		return nil
	}
	gotGet := 0
	app.MockFederateApp.get = func(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
		gotGet++
		if *id != *noteIRI {
			t.Fatalf("expected %s, got %s", noteIRI, id)
		}
		samActor := &vocab.Person{}
		samActor.SetInboxAnyURI(samIRIInbox)
		samActor.SetId(samIRI)
		samActor.AppendNameString("Sam")
		v := &vocab.Note{}
		v.SetId(noteIRI)
		v.AppendNameString(noteName)
		v.AppendContentString("This is a simple note")
		v.AppendToObject(samActor)
		return v, nil
	}
	gotSet := 0
	var gotSetObject PubObject
	app.MockFederateApp.set = func(c context.Context, p PubObject) error {
		gotSet++
		if gotSet == 1 {
			gotSetObject = p
		}
		return nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	expectedSamActor := &vocab.Person{}
	expectedSamActor.SetInboxAnyURI(samIRIInbox)
	expectedSamActor.SetId(samIRI)
	expectedUpdatedNote := &vocab.Note{}
	expectedUpdatedNote.SetId(noteIRI)
	expectedUpdatedNote.AppendNameString(noteName)
	expectedUpdatedNote.AppendContentString("This is a simple note")
	expectedUpdatedNote.AppendToObject(expectedSamActor)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotGet != 1 {
		t.Fatalf("expected %d, got %d", 1, gotGet)
	} else if gotSet != 3 {
		t.Fatalf("expected %d, got %d", 3, gotSet)
	} else if err := PubObjectEquals(gotSetObject, expectedUpdatedNote); err != nil {
		t.Fatalf("unexpected set object: %s", err)
	}
}

func TestPostOutbox_Update_DeleteFields(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer([]byte(testDeleteFields)))))
	socialCb.update = func(c context.Context, s *streams.Update) error {
		return nil
	}
	gotGet := 0
	app.MockFederateApp.get = func(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
		if rw != ReadWrite {
			t.Fatalf("expected RWType of %d, got %d", ReadWrite, rw)
		}
		gotGet++
		if *id != *noteIRI {
			t.Fatalf("expected %s, got %s", noteIRI, id)
		}
		samActor := &vocab.Person{}
		samActor.SetInboxAnyURI(samIRIInbox)
		samActor.SetId(samIRI)
		samActor.AppendNameString("Sam")
		v := &vocab.Note{}
		v.SetId(noteIRI)
		v.AppendNameString(noteName)
		v.AppendContentString("This is a simple note")
		v.AppendToObject(samActor)
		return v, nil
	}
	gotSet := 0
	var gotSetObject PubObject
	app.MockFederateApp.set = func(c context.Context, p PubObject) error {
		gotSet++
		if gotSet == 1 {
			gotSetObject = p
		}
		return nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	expectedUpdatedNote := &vocab.Note{}
	expectedUpdatedNote.SetId(noteIRI)
	expectedUpdatedNote.AppendNameString(noteName)
	expectedUpdatedNote.AppendContentString("This is a simple note")
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotGet != 1 {
		t.Fatalf("expected %d, got %d", 1, gotGet)
	} else if gotSet != 3 {
		t.Fatalf("expected %d, got %d", 3, gotSet)
	} else if err := PubObjectEquals(gotSetObject, expectedUpdatedNote); err != nil {
		t.Fatalf("unexpected set object: %s", err)
	}
}

func TestPostOutbox_Update_DeleteSubFieldsMultipleObjects(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer([]byte(testDeleteFieldsDifferentObjects)))))
	socialCb.update = func(c context.Context, s *streams.Update) error {
		return nil
	}
	gotGet := 0
	app.MockFederateApp.get = func(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
		gotGet++
		var v *vocab.Note
		if *id == *noteIRI {
			samActor := &vocab.Person{}
			samActor.SetInboxAnyURI(samIRIInbox)
			samActor.SetId(samIRI)
			samActor.AppendNameString("Sam")
			v = &vocab.Note{}
			v.SetId(noteIRI)
			v.AppendNameString(noteName)
			v.AppendContentString("This is a simple note")
			v.AppendToObject(samActor)
		} else if *id == *updateActivityIRI {
			samActor := &vocab.Person{}
			samActor.SetInboxAnyURI(samIRIInbox)
			samActor.SetId(samIRI)
			samActor.AppendNameString("Sam")
			v = &vocab.Note{}
			v.SetId(updateActivityIRI)
			v.AppendNameString(noteName)
			v.AppendContentString("This is a simple note")
			v.AppendToObject(samActor)
		} else {
			t.Fatalf("unexpected app.MockFederateApp.Get id: %s", id)
		}
		return v, nil
	}
	gotSet := 0
	var gotSetObject PubObject
	var gotSetObject2 PubObject
	app.MockFederateApp.set = func(c context.Context, p PubObject) error {
		gotSet++
		if gotSet == 1 {
			gotSetObject = p
		} else if gotSet == 2 {
			gotSetObject2 = p
		}
		return nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	expectedSamActor := &vocab.Person{}
	expectedSamActor.SetInboxAnyURI(samIRIInbox)
	expectedSamActor.SetId(samIRI)
	expectedUpdatedNote := &vocab.Note{}
	expectedUpdatedNote.SetId(noteIRI)
	expectedUpdatedNote.AppendNameString(noteName)
	expectedUpdatedNote.AppendContentString("This is a simple note")
	expectedUpdatedNote.AppendToObject(expectedSamActor)
	expectedUpdatedNote2 := &vocab.Note{}
	expectedUpdatedNote2.SetId(updateActivityIRI)
	expectedUpdatedNote2.AppendNameString(noteName)
	expectedUpdatedNote2.AppendContentString("This is a simple note")
	expectedUpdatedNote2.AppendToObject(expectedSamActor)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotGet != 2 {
		t.Fatalf("expected %d, got %d", 2, gotGet)
	} else if gotSet != 4 {
		t.Fatalf("expected %d, got %d", 4, gotSet)
	} else if err := PubObjectEquals(gotSetObject, expectedUpdatedNote); err != nil {
		t.Fatalf("unexpected set object: %s", err)
	} else if err := PubObjectEquals(gotSetObject2, expectedUpdatedNote2); err != nil {
		t.Fatalf("unexpected set object: %s", err)
	}
}

func TestPostOutbox_Update_OverwriteUpdatedFields(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testClientUpdateNote)))))
	socialCb.update = func(c context.Context, s *streams.Update) error {
		return nil
	}
	gotGet := 0
	app.MockFederateApp.get = func(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
		gotGet++
		if *id != *noteIRI {
			t.Fatalf("expected %s, got %s", noteIRI, id)
		}
		samActor := &vocab.Person{}
		samActor.SetInboxAnyURI(samIRIInbox)
		samActor.SetId(samIRI)
		samActor.AppendNameString("Sam")
		v := &vocab.Note{}
		v.SetId(noteIRI)
		v.AppendNameString(noteName)
		v.AppendContentString("This is a simple note")
		v.AppendToObject(samActor)
		return v, nil
	}
	gotSet := 0
	var gotSetObject PubObject
	app.MockFederateApp.set = func(c context.Context, p PubObject) error {
		gotSet++
		if gotSet == 1 {
			gotSetObject = p
		}
		return nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	samActor := &vocab.Person{}
	samActor.SetInboxAnyURI(samIRIInbox)
	samActor.SetId(samIRI)
	samActor.AppendNameString("Samm")
	expectedUpdatedNote := &vocab.Note{}
	expectedUpdatedNote.SetId(noteIRI)
	expectedUpdatedNote.AppendNameString(noteName)
	expectedUpdatedNote.AppendContentString("This is a simple note")
	expectedUpdatedNote.AppendToObject(samActor)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotGet != 1 {
		t.Fatalf("expected %d, got %d", 1, gotGet)
	} else if gotSet != 3 {
		t.Fatalf("expected %d, got %d", 3, gotSet)
	} else if err := PubObjectEquals(gotSetObject, expectedUpdatedNote); err != nil {
		t.Fatalf("unexpected set object: %s", err)
	}
}

func TestPostOutbox_Update_CallsCallback(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testUpdateNote)))))
	gotCallback := 0
	var gotCallbackObject *streams.Update
	socialCb.update = func(c context.Context, s *streams.Update) error {
		gotCallback++
		gotCallbackObject = s
		return nil
	}
	app.MockFederateApp.get = func(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
		samActor := &vocab.Person{}
		samActor.SetInboxAnyURI(samIRIInbox)
		samActor.SetId(samIRI)
		samActor.AppendNameString("Sam")
		v := &vocab.Note{}
		v.SetId(noteIRI)
		v.AppendNameString(noteName)
		v.AppendContentString("This is a simple note")
		v.AppendToObject(samActor)
		return v, nil
	}
	app.MockFederateApp.set = func(c context.Context, p PubObject) error {
		return nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotCallback != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCallback)
	} else if err := PubObjectEquals(gotCallbackObject.Raw(), testClientExpectedUpdateNote); err != nil {
		t.Fatalf("unexpected callback object: %s", err)
	}
}

func TestPostOutbox_Update_IsDelivered(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testUpdateNote)))))
	socialCb.update = func(c context.Context, s *streams.Update) error {
		return nil
	}
	gotHttpDo := 0
	var httpDeliveryRequest *http.Request
	httpClient.do = func(req *http.Request) (*http.Response, error) {
		gotHttpDo++
		if gotHttpDo == 1 {
			actorResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer(samActorJSON)),
			}
			return actorResp, nil
		} else if gotHttpDo == 2 {
			actorResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer(sallyActorJSON)),
			}
			return actorResp, nil
		} else if gotHttpDo == 3 {
			httpDeliveryRequest = req
			okResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte{})),
			}
			return okResp, nil
		}
		return nil, nil
	}
	app.MockFederateApp.get = func(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
		samActor := &vocab.Person{}
		samActor.SetInboxAnyURI(samIRIInbox)
		samActor.SetId(samIRI)
		samActor.AppendNameString("Sam")
		v := &vocab.Note{}
		v.SetId(noteIRI)
		v.AppendNameString(noteName)
		v.AppendContentString("This is a simple note")
		v.AppendToObject(samActor)
		return v, nil
	}
	app.MockFederateApp.set = func(c context.Context, p PubObject) error {
		return nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotHttpDo != 3 {
		t.Fatalf("expected %d, got %d", 3, gotHttpDo)
	} else if httpDeliveryRequest.Method != "POST" {
		t.Fatalf("expected %s, got %s", "POST", httpDeliveryRequest.Method)
	} else if s := httpDeliveryRequest.URL.String(); s != samIRIInboxString {
		t.Fatalf("expected %s, got %s", samIRIInboxString, s)
	}
}

func TestPostOutbox_Delete_SetsTombstone(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testDeleteNote)))))
	socialCb.delete = func(c context.Context, s *streams.Delete) error {
		return nil
	}
	gotGet := 0
	app.MockFederateApp.get = func(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
		if rw != ReadWrite {
			t.Fatalf("expected RWType of %d, got %d", ReadWrite, rw)
		}
		gotGet++
		if *id != *noteIRI {
			t.Fatalf("expected %s, got %s", noteIRI, id)
		}
		v := &vocab.Note{}
		v.AppendType("Note")
		v.SetId(noteIRI)
		v.SetPublished(testPublishedTime)
		v.SetUpdated(testUpdateTime)
		return v, nil
	}
	gotSet := 0
	var gotSetObject PubObject
	app.MockFederateApp.set = func(c context.Context, p PubObject) error {
		gotSet++
		if gotSet == 1 {
			gotSetObject = p
		}
		return nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	expectedTombstone := &vocab.Tombstone{}
	expectedTombstone.SetId(noteIRI)
	expectedTombstone.SetPublished(testPublishedTime)
	expectedTombstone.SetUpdated(testUpdateTime)
	expectedTombstone.SetDeleted(now)
	expectedTombstone.AppendFormerTypeString("Note")
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotGet != 1 {
		t.Fatalf("expected %d, got %d", 1, gotGet)
	} else if gotSet != 3 {
		t.Fatalf("expected %d, got %d", 3, gotSet)
	} else if err := PubObjectEquals(gotSetObject, expectedTombstone); err != nil {
		t.Fatalf("unexpected set object: %s", err)
	}
}

func TestPostOutbox_Delete_CallsCallback(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testDeleteNote)))))
	gotCallback := 0
	var gotCallbackObject *streams.Delete
	socialCb.delete = func(c context.Context, s *streams.Delete) error {
		gotCallback++
		gotCallbackObject = s
		return nil
	}
	app.MockFederateApp.get = func(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
		return testNote, nil
	}
	app.MockFederateApp.set = func(c context.Context, p PubObject) error {
		return nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotCallback != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCallback)
	} else if err := PubObjectEquals(gotCallbackObject.Raw(), testClientExpectedDeleteNote); err != nil {
		t.Fatalf("unexpected callback object: %s", err)
	}
}

func TestPostOutbox_Delete_IsDelivered(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testDeleteNote)))))
	socialCb.delete = func(c context.Context, s *streams.Delete) error {
		return nil
	}
	gotHttpDo := 0
	var httpDeliveryRequest *http.Request
	httpClient.do = func(req *http.Request) (*http.Response, error) {
		gotHttpDo++
		if gotHttpDo == 1 {
			actorResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer(samActorJSON)),
			}
			return actorResp, nil
		} else if gotHttpDo == 2 {
			actorResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer(sallyActorJSON)),
			}
			return actorResp, nil
		} else if gotHttpDo == 3 {
			httpDeliveryRequest = req
			okResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte{})),
			}
			return okResp, nil
		}
		return nil, nil
	}
	app.MockFederateApp.get = func(c context.Context, id *url.URL, rw RWType) (PubObject, error) {
		return testNote, nil
	}
	app.MockFederateApp.set = func(c context.Context, p PubObject) error {
		return nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotHttpDo != 3 {
		t.Fatalf("expected %d, got %d", 3, gotHttpDo)
	} else if httpDeliveryRequest.Method != "POST" {
		t.Fatalf("expected %s, got %s", "POST", httpDeliveryRequest.Method)
	} else if s := httpDeliveryRequest.URL.String(); s != samIRIInboxString {
		t.Fatalf("expected %s, got %s", samIRIInboxString, s)
	}
}

func TestPostOutbox_Follow_CallsCallback(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testFollow)))))
	gotCallback := 0
	var gotCallbackObject *streams.Follow
	socialCb.follow = func(c context.Context, s *streams.Follow) error {
		gotCallback++
		gotCallbackObject = s
		return nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotCallback != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCallback)
	} else if err := PubObjectEquals(gotCallbackObject.Raw(), testClientExpectedFollow); err != nil {
		t.Fatalf("unexpected callback object: %s", err)
	}
}

func TestPostOutbox_Follow_IsDelivered(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testFollow)))))
	socialCb.follow = func(c context.Context, s *streams.Follow) error {
		return nil
	}
	gotHttpDo := 0
	var httpDeliveryRequest *http.Request
	httpClient.do = func(req *http.Request) (*http.Response, error) {
		gotHttpDo++
		if gotHttpDo == 1 {
			actorResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer(samActorJSON)),
			}
			return actorResp, nil
		} else if gotHttpDo == 2 {
			actorResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer(sallyActorJSON)),
			}
			return actorResp, nil
		} else if gotHttpDo == 3 {
			httpDeliveryRequest = req
			okResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte{})),
			}
			return okResp, nil
		}
		return nil, nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotHttpDo != 3 {
		t.Fatalf("expected %d, got %d", 3, gotHttpDo)
	} else if httpDeliveryRequest.Method != "POST" {
		t.Fatalf("expected %s, got %s", "POST", httpDeliveryRequest.Method)
	} else if s := httpDeliveryRequest.URL.String(); s != samIRIInboxString {
		t.Fatalf("expected %s, got %s", samIRIInboxString, s)
	}
}

func TestPostOutbox_Accept_CallsCallback(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testAcceptFollow)))))
	gotCallback := 0
	var gotCallbackObject *streams.Accept
	socialCb.accept = func(c context.Context, s *streams.Accept) error {
		gotCallback++
		gotCallbackObject = s
		return nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotCallback != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCallback)
	} else if err := PubObjectEquals(gotCallbackObject.Raw(), testClientExpectedAcceptFollow); err != nil {
		t.Fatalf("unexpected callback object: %s", err)
	}
}

func TestPostOutbox_Accept_IsDelivered(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testAcceptFollow)))))
	socialCb.accept = func(c context.Context, s *streams.Accept) error {
		return nil
	}
	gotHttpDo := 0
	var httpDeliveryRequest *http.Request
	httpClient.do = func(req *http.Request) (*http.Response, error) {
		gotHttpDo++
		if gotHttpDo == 1 {
			actorResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer(samActorJSON)),
			}
			return actorResp, nil
		} else if gotHttpDo == 2 {
			actorResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer(sallyActorJSON)),
			}
			return actorResp, nil
		} else if gotHttpDo == 3 {
			httpDeliveryRequest = req
			okResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte{})),
			}
			return okResp, nil
		}
		return nil, nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotHttpDo != 3 {
		t.Fatalf("expected %d, got %d", 3, gotHttpDo)
	} else if httpDeliveryRequest.Method != "POST" {
		t.Fatalf("expected %s, got %s", "POST", httpDeliveryRequest.Method)
	} else if s := httpDeliveryRequest.URL.String(); s != samIRIInboxString {
		t.Fatalf("expected %s, got %s", samIRIInboxString, s)
	}
}

func TestPostOutbox_Reject_CallsCallback(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testRejectFollow)))))
	gotCallback := 0
	var gotCallbackObject *streams.Reject
	socialCb.reject = func(c context.Context, s *streams.Reject) error {
		gotCallback++
		gotCallbackObject = s
		return nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotCallback != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCallback)
	} else if err := PubObjectEquals(gotCallbackObject.Raw(), testClientExpectedRejectFollow); err != nil {
		t.Fatalf("unexpected callback object: %s", err)
	}
}

func TestPostOutbox_Reject_IsDelivered(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testRejectFollow)))))
	socialCb.reject = func(c context.Context, s *streams.Reject) error {
		return nil
	}
	gotHttpDo := 0
	var httpDeliveryRequest *http.Request
	httpClient.do = func(req *http.Request) (*http.Response, error) {
		gotHttpDo++
		if gotHttpDo == 1 {
			actorResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer(samActorJSON)),
			}
			return actorResp, nil
		} else if gotHttpDo == 2 {
			actorResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer(sallyActorJSON)),
			}
			return actorResp, nil
		} else if gotHttpDo == 3 {
			httpDeliveryRequest = req
			okResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte{})),
			}
			return okResp, nil
		}
		return nil, nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotHttpDo != 3 {
		t.Fatalf("expected %d, got %d", 3, gotHttpDo)
	} else if httpDeliveryRequest.Method != "POST" {
		t.Fatalf("expected %s, got %s", "POST", httpDeliveryRequest.Method)
	} else if s := httpDeliveryRequest.URL.String(); s != samIRIInboxString {
		t.Fatalf("expected %s, got %s", samIRIInboxString, s)
	}
}

func TestPostOutbox_Add_DoesNotAddIfTargetNotOwned(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testAddNote)))))
	gotOwns := 0
	var gotOwnsIri *url.URL
	app.MockFederateApp.owns = func(c context.Context, iri *url.URL) bool {
		gotOwns++
		gotOwnsIri = iri
		return false
	}
	socialCb.add = func(c context.Context, s *streams.Add) error {
		return nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotOwns != 1 {
		t.Fatalf("expected %d, got %d", 1, gotOwns)
	} else if ownsIri := gotOwnsIri.String(); ownsIri != iriString {
		t.Fatalf("expected %s, got %s", iriString, ownsIri)
	}
}

func TestPostOutbox_Add_AddsIfTargetOwnedAndAppCanAdd(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testAddNote)))))
	app.MockFederateApp.owns = func(c context.Context, iri *url.URL) bool {
		return true
	}
	gotGet := 0
	var gotGetIri *url.URL
	app.MockFederateApp.get = func(c context.Context, iri *url.URL, rw RWType) (PubObject, error) {
		if rw != ReadWrite {
			t.Fatalf("expected RWType of %d, got %d", ReadWrite, rw)
		}
		gotGet++
		gotGetIri = iri
		col := &vocab.Collection{}
		col.AppendType("Collection")
		return col, nil
	}
	gotCanAdd := 0
	var canAddObj vocab.ObjectType
	var canAddTarget vocab.ObjectType
	app.MockFederateApp.canAdd = func(c context.Context, obj vocab.ObjectType, t vocab.ObjectType) bool {
		gotCanAdd++
		canAddObj = obj
		canAddTarget = t
		return true
	}
	gotSet := 0
	var gotSetObj PubObject
	app.MockFederateApp.set = func(c context.Context, t PubObject) error {
		gotSet++
		if gotSet == 1 {
			gotSetObj = t
		}
		return nil
	}
	socialCb.add = func(c context.Context, s *streams.Add) error {
		return nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	expectedTarget := &vocab.Collection{}
	expectedTarget.AppendType("Collection")
	expectedTarget.AppendItemsIRI(noteIRI)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotGet != 1 {
		t.Fatalf("expected %d, got %d", 1, gotGet)
	} else if getIri := gotGetIri.String(); getIri != iriString {
		t.Fatalf("expected %s, got %s", iriString, getIri)
	} else if gotCanAdd != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCanAdd)
	} else if err := VocabSerializerEquals(canAddTarget, expectedTarget); err != nil {
		t.Fatal(err)
	} else if err := VocabSerializerEquals(canAddObj, testNote); err != nil {
		t.Fatal(err)
	} else if gotSet != 3 {
		t.Fatalf("expected %d, got %d", 3, gotSet)
	} else if err := PubObjectEquals(gotSetObj, expectedTarget); err != nil {
		t.Fatalf("unexpected set object: %s", err)
	}
}

func TestPostOutbox_Add_DoesNotAddIfAppCannotAdd(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testAddNote)))))
	app.MockFederateApp.owns = func(c context.Context, iri *url.URL) bool {
		return true
	}
	gotGet := 0
	var gotGetIri *url.URL
	app.MockFederateApp.get = func(c context.Context, iri *url.URL, rw RWType) (PubObject, error) {
		gotGet++
		gotGetIri = iri
		col := &vocab.Collection{}
		col.AppendType("Collection")
		return col, nil
	}
	gotCanAdd := 0
	var canAddObj vocab.ObjectType
	var canAddTarget vocab.ObjectType
	app.MockFederateApp.canAdd = func(c context.Context, obj vocab.ObjectType, t vocab.ObjectType) bool {
		gotCanAdd++
		canAddObj = obj
		canAddTarget = t
		return false
	}
	socialCb.add = func(c context.Context, s *streams.Add) error {
		return nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	expectedTarget := &vocab.Collection{}
	expectedTarget.AppendType("Collection")
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotGet != 1 {
		t.Fatalf("expected %d, got %d", 1, gotGet)
	} else if getIri := gotGetIri.String(); getIri != iriString {
		t.Fatalf("expected %s, got %s", iriString, getIri)
	} else if gotCanAdd != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCanAdd)
	} else if err := VocabSerializerEquals(canAddTarget, expectedTarget); err != nil {
		t.Fatal(err)
	} else if err := VocabSerializerEquals(canAddObj, testNote); err != nil {
		t.Fatal(err)
	}
}

func TestPostOutbox_Add_CallsCallback(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testAddNote)))))
	app.MockFederateApp.owns = func(c context.Context, iri *url.URL) bool {
		return false
	}
	gotCallback := 0
	var gotCallbackObject *streams.Add
	socialCb.add = func(c context.Context, s *streams.Add) error {
		gotCallback++
		gotCallbackObject = s
		return nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotCallback != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCallback)
	} else if err := PubObjectEquals(gotCallbackObject.Raw(), testClientExpectedAdd); err != nil {
		t.Fatalf("unexpected callback object: %s", err)
	}
}

func TestPostOutbox_Add_IsDelivered(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testAddNote)))))
	app.MockFederateApp.owns = func(c context.Context, iri *url.URL) bool {
		return false
	}
	socialCb.add = func(c context.Context, s *streams.Add) error {
		return nil
	}
	gotHttpDo := 0
	var httpDeliveryRequest *http.Request
	httpClient.do = func(req *http.Request) (*http.Response, error) {
		gotHttpDo++
		if gotHttpDo == 1 {
			actorResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer(samActorJSON)),
			}
			return actorResp, nil
		} else if gotHttpDo == 2 {
			actorResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer(sallyActorJSON)),
			}
			return actorResp, nil
		} else if gotHttpDo == 3 {
			httpDeliveryRequest = req
			okResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte{})),
			}
			return okResp, nil
		}
		return nil, nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotHttpDo != 3 {
		t.Fatalf("expected %d, got %d", 3, gotHttpDo)
	} else if httpDeliveryRequest.Method != "POST" {
		t.Fatalf("expected %s, got %s", "POST", httpDeliveryRequest.Method)
	} else if s := httpDeliveryRequest.URL.String(); s != samIRIInboxString {
		t.Fatalf("expected %s, got %s", samIRIInboxString, s)
	}
}

func TestPostOutbox_Remove_DoesNotRemoveIfTargetNotOwned(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testRemoveNote)))))
	gotOwns := 0
	var gotOwnsIri *url.URL
	app.MockFederateApp.owns = func(c context.Context, iri *url.URL) bool {
		gotOwns++
		gotOwnsIri = iri
		return false
	}
	socialCb.remove = func(c context.Context, s *streams.Remove) error {
		return nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotOwns != 1 {
		t.Fatalf("expected %d, got %d", 1, gotOwns)
	} else if ownsIri := gotOwnsIri.String(); ownsIri != iriString {
		t.Fatalf("expected %s, got %s", iriString, ownsIri)
	}
}

func TestPostOutbox_Remove_RemoveIfTargetOwnedAndCanRemove(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testRemoveNote)))))
	app.MockFederateApp.owns = func(c context.Context, iri *url.URL) bool {
		return true
	}
	gotGet := 0
	var gotGetIri *url.URL
	app.MockFederateApp.get = func(c context.Context, iri *url.URL, rw RWType) (PubObject, error) {
		gotGet++
		gotGetIri = iri
		col := &vocab.Collection{}
		col.AppendType("Collection")
		col.AppendItemsObject(testNote)
		return col, nil
	}
	gotCanRemove := 0
	var canRemoveObj vocab.ObjectType
	var canRemoveTarget vocab.ObjectType
	app.MockFederateApp.canRemove = func(c context.Context, obj vocab.ObjectType, t vocab.ObjectType) bool {
		gotCanRemove++
		canRemoveObj = obj
		canRemoveTarget = t
		return true
	}
	gotSet := 0
	var gotSetObj PubObject
	app.MockFederateApp.set = func(c context.Context, t PubObject) error {
		gotSet++
		if gotSet == 1 {
			gotSetObj = t
		}
		return nil
	}
	socialCb.remove = func(c context.Context, s *streams.Remove) error {
		return nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	expectedTarget := &vocab.Collection{}
	expectedTarget.AppendType("Collection")
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotGet != 1 {
		t.Fatalf("expected %d, got %d", 1, gotGet)
	} else if getIri := gotGetIri.String(); getIri != iriString {
		t.Fatalf("expected %s, got %s", iriString, getIri)
	} else if gotCanRemove != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCanRemove)
	} else if err := VocabSerializerEquals(canRemoveTarget, expectedTarget); err != nil {
		t.Fatal(err)
	} else if err := VocabSerializerEquals(canRemoveObj, testNote); err != nil {
		t.Fatal(err)
	} else if gotSet != 3 {
		t.Fatalf("expected %d, got %d", 3, gotSet)
	} else if err := PubObjectEquals(gotSetObj, expectedTarget); err != nil {
		t.Fatalf("unexpected set object: %s", err)
	}
}

func TestPostOutbox_Remove_DoesNotRemoveIfAppCannotRemove(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testRemoveNote)))))
	app.MockFederateApp.owns = func(c context.Context, iri *url.URL) bool {
		return true
	}
	gotGet := 0
	var gotGetIri *url.URL
	app.MockFederateApp.get = func(c context.Context, iri *url.URL, rw RWType) (PubObject, error) {
		gotGet++
		gotGetIri = iri
		col := &vocab.Collection{}
		col.AppendType("Collection")
		col.AppendItemsObject(testNote)
		return col, nil
	}
	gotCanRemove := 0
	var canRemoveObj vocab.ObjectType
	var canRemoveTarget vocab.ObjectType
	app.MockFederateApp.canRemove = func(c context.Context, obj vocab.ObjectType, t vocab.ObjectType) bool {
		gotCanRemove++
		canRemoveObj = obj
		canRemoveTarget = t
		return false
	}
	socialCb.remove = func(c context.Context, s *streams.Remove) error {
		return nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	expectedTarget := &vocab.Collection{}
	expectedTarget.AppendType("Collection")
	expectedTarget.AppendItemsObject(testNote)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotGet != 1 {
		t.Fatalf("expected %d, got %d", 1, gotGet)
	} else if getIri := gotGetIri.String(); getIri != iriString {
		t.Fatalf("expected %s, got %s", iriString, getIri)
	} else if gotCanRemove != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCanRemove)
	} else if err := VocabSerializerEquals(canRemoveTarget, expectedTarget); err != nil {
		t.Fatal(err)
	} else if err := VocabSerializerEquals(canRemoveObj, testNote); err != nil {
		t.Fatal(err)
	}
}

func TestPostOutbox_Remove_CallsCallback(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testRemoveNote)))))
	app.MockFederateApp.owns = func(c context.Context, iri *url.URL) bool {
		return false
	}
	gotCallback := 0
	var gotCallbackObject *streams.Remove
	socialCb.remove = func(c context.Context, s *streams.Remove) error {
		gotCallback++
		gotCallbackObject = s
		return nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotCallback != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCallback)
	} else if err := PubObjectEquals(gotCallbackObject.Raw(), testClientExpectedRemove); err != nil {
		t.Fatalf("unexpected callback object: %s", err)
	}
}

func TestPostOutbox_Remove_IsDelivered(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testRemoveNote)))))
	app.MockFederateApp.owns = func(c context.Context, iri *url.URL) bool {
		return false
	}
	socialCb.remove = func(c context.Context, s *streams.Remove) error {
		return nil
	}
	gotHttpDo := 0
	var httpDeliveryRequest *http.Request
	httpClient.do = func(req *http.Request) (*http.Response, error) {
		gotHttpDo++
		if gotHttpDo == 1 {
			actorResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer(samActorJSON)),
			}
			return actorResp, nil
		} else if gotHttpDo == 2 {
			actorResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer(sallyActorJSON)),
			}
			return actorResp, nil
		} else if gotHttpDo == 3 {
			httpDeliveryRequest = req
			okResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte{})),
			}
			return okResp, nil
		}
		return nil, nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotHttpDo != 3 {
		t.Fatalf("expected %d, got %d", 3, gotHttpDo)
	} else if httpDeliveryRequest.Method != "POST" {
		t.Fatalf("expected %s, got %s", "POST", httpDeliveryRequest.Method)
	} else if s := httpDeliveryRequest.URL.String(); s != samIRIInboxString {
		t.Fatalf("expected %s, got %s", samIRIInboxString, s)
	}
}

func TestPostOutbox_Like_AddsToLikedCollection(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testLikeNote)))))
	gotOwns := 0
	var gotOwnsIri *url.URL
	app.MockFederateApp.owns = func(c context.Context, iri *url.URL) bool {
		gotOwns++
		gotOwnsIri = iri
		return true
	}
	gotGet := 0
	var gotGetIri *url.URL
	app.MockFederateApp.get = func(c context.Context, iri *url.URL, rw RWType) (PubObject, error) {
		if rw != ReadWrite {
			t.Fatalf("expected RWType of %d, got %d", ReadWrite, rw)
		}
		gotGet++
		gotGetIri = iri
		v := &vocab.Person{}
		v.AppendNameString("Sally")
		v.SetId(sallyIRI)
		v.SetLikedCollection(&vocab.Collection{})
		return v, nil
	}
	gotSet := 0
	var gotSetObj PubObject
	app.MockFederateApp.set = func(c context.Context, o PubObject) error {
		gotSet++
		if gotSet == 1 {
			gotSetObj = o
		}
		return nil
	}
	socialCb.like = func(c context.Context, s *streams.Like) error {
		return nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	expectedLikes := &vocab.Collection{}
	expectedLikes.AppendItemsIRI(noteIRI)
	expectedActor := &vocab.Person{}
	expectedActor.AppendNameString("Sally")
	expectedActor.SetId(sallyIRI)
	expectedActor.SetLikedCollection(expectedLikes)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotOwns != 1 {
		t.Fatalf("expected %d, got %d", 1, gotOwns)
	} else if gotOwnsString := gotOwnsIri.String(); gotOwnsString != sallyIRIString {
		t.Fatalf("expected %s, got %s", noteURIString, sallyIRIString)
	} else if gotGet != 1 {
		t.Fatalf("expected %d, got %d", 1, gotGet)
	} else if gotGetString := gotGetIri.String(); gotGetString != sallyIRIString {
		t.Fatalf("expected %s, got %s", noteURIString, sallyIRIString)
	} else if gotSet != 3 {
		t.Fatalf("expected %d, got %d", 3, gotSet)
	} else if err := PubObjectEquals(gotSetObj, expectedActor); err != nil {
		t.Fatalf("set obj: %s", err)
	}
}

func TestPostOutbox_Like_AddsToDefaultOrderedCollection(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testLikeNote)))))
	gotOwns := 0
	var gotOwnsIri *url.URL
	app.MockFederateApp.owns = func(c context.Context, iri *url.URL) bool {
		gotOwns++
		gotOwnsIri = iri
		return true
	}
	gotGet := 0
	var gotGetIri *url.URL
	app.MockFederateApp.get = func(c context.Context, iri *url.URL, rw RWType) (PubObject, error) {
		if rw != ReadWrite {
			t.Fatalf("expected RWType of %d, got %d", ReadWrite, rw)
		}
		gotGet++
		gotGetIri = iri
		v := &vocab.Person{}
		v.AppendNameString("Sally")
		v.SetId(sallyIRI)
		return v, nil
	}
	gotSet := 0
	var gotSetObj PubObject
	app.MockFederateApp.set = func(c context.Context, o PubObject) error {
		gotSet++
		if gotSet == 1 {
			gotSetObj = o
		}
		return nil
	}
	socialCb.like = func(c context.Context, s *streams.Like) error {
		return nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	expectedLikes := &vocab.OrderedCollection{}
	expectedLikes.AppendOrderedItemsIRI(noteIRI)
	expectedActor := &vocab.Person{}
	expectedActor.AppendNameString("Sally")
	expectedActor.SetId(sallyIRI)
	expectedActor.SetLikedOrderedCollection(expectedLikes)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotOwns != 1 {
		t.Fatalf("expected %d, got %d", 1, gotOwns)
	} else if gotOwnsString := gotOwnsIri.String(); gotOwnsString != sallyIRIString {
		t.Fatalf("expected %s, got %s", noteURIString, sallyIRIString)
	} else if gotGet != 1 {
		t.Fatalf("expected %d, got %d", 1, gotGet)
	} else if gotGetString := gotGetIri.String(); gotGetString != sallyIRIString {
		t.Fatalf("expected %s, got %s", noteURIString, sallyIRIString)
	} else if gotSet != 3 {
		t.Fatalf("expected %d, got %d", 3, gotSet)
	} else if err := PubObjectEquals(gotSetObj, expectedActor); err != nil {
		t.Fatalf("set obj: %s", err)
	}
}

func TestPostOutbox_Like_DoesNotAddIfAlreadyLiked(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testLikeNote)))))
	app.MockFederateApp.owns = func(c context.Context, iri *url.URL) bool {
		return true
	}
	app.MockFederateApp.get = func(c context.Context, iri *url.URL, rw RWType) (PubObject, error) {
		liked := &vocab.Collection{}
		liked.AppendItemsIRI(noteIRI)
		v := &vocab.Person{}
		v.AppendNameString("Sally")
		v.SetId(sallyIRI)
		v.SetLikedCollection(liked)
		return v, nil
	}
	gotSet := 0
	var gotSetObj PubObject
	app.MockFederateApp.set = func(c context.Context, o PubObject) error {
		gotSet++
		if gotSet == 1 {
			gotSetObj = o
		}
		return nil
	}
	socialCb.like = func(c context.Context, s *streams.Like) error {
		return nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	expectedLikes := &vocab.Collection{}
	expectedLikes.AppendItemsIRI(noteIRI)
	expectedActor := &vocab.Person{}
	expectedActor.AppendNameString("Sally")
	expectedActor.SetId(sallyIRI)
	expectedActor.SetLikedCollection(expectedLikes)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotSet != 3 {
		t.Fatalf("expected %d, got %d", 3, gotSet)
	} else if err := PubObjectEquals(gotSetObj, expectedActor); err != nil {
		t.Fatalf("set obj: %s", err)
	}
}

func TestPostOutbox_Like_AddsToLikedOrderedCollection(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testLikeNote)))))
	app.MockFederateApp.owns = func(c context.Context, iri *url.URL) bool {
		return true
	}
	app.MockFederateApp.get = func(c context.Context, iri *url.URL, rw RWType) (PubObject, error) {
		v := &vocab.Person{}
		v.AppendNameString("Sally")
		v.SetId(sallyIRI)
		v.SetLikedOrderedCollection(&vocab.OrderedCollection{})
		return v, nil
	}
	gotSet := 0
	var gotSetObj PubObject
	app.MockFederateApp.set = func(c context.Context, o PubObject) error {
		gotSet++
		if gotSet == 1 {
			gotSetObj = o
		}
		return nil
	}
	socialCb.like = func(c context.Context, s *streams.Like) error {
		return nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	expectedLikes := &vocab.OrderedCollection{}
	expectedLikes.AppendOrderedItemsIRI(noteIRI)
	expectedActor := &vocab.Person{}
	expectedActor.AppendNameString("Sally")
	expectedActor.SetId(sallyIRI)
	expectedActor.SetLikedOrderedCollection(expectedLikes)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if err := PubObjectEquals(gotSetObj, expectedActor); err != nil {
		t.Fatalf("set obj: %s", err)
	}
}

func TestPostOutbox_Like_DoesNotAddIfCollectionNotOwned(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testLikeNote)))))
	gotOwns := 0
	var gotOwnsIri *url.URL
	app.MockFederateApp.owns = func(c context.Context, iri *url.URL) bool {
		gotOwns++
		gotOwnsIri = iri
		return false
	}
	socialCb.like = func(c context.Context, s *streams.Like) error {
		return nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotOwns != 1 {
		t.Fatalf("expected %d, got %d", 1, gotOwns)
	} else if gotOwnsString := gotOwnsIri.String(); gotOwnsString != sallyIRIString {
		t.Fatalf("expected %s, got %s", noteURIString, sallyIRIString)
	}
}

func TestPostOutbox_Like_CallsCallback(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testLikeNote)))))
	app.MockFederateApp.owns = func(c context.Context, iri *url.URL) bool {
		return true
	}
	app.MockFederateApp.get = func(c context.Context, iri *url.URL, rw RWType) (PubObject, error) {
		v := &vocab.Person{}
		v.AppendNameString("Sally")
		v.SetId(sallyIRI)
		v.SetLikedOrderedCollection(&vocab.OrderedCollection{})
		return v, nil
	}
	app.MockFederateApp.set = func(c context.Context, o PubObject) error {
		return nil
	}
	gotCallback := 0
	var gotCallbackObject *streams.Like
	socialCb.like = func(c context.Context, s *streams.Like) error {
		gotCallback++
		gotCallbackObject = s
		return nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotCallback != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCallback)
	} else if err := PubObjectEquals(gotCallbackObject.Raw(), testClientExpectedLike); err != nil {
		t.Fatalf("unexpected callback object: %s", err)
	}
}

func TestPostOutbox_Like_IsDelivered(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testLikeNote)))))
	app.MockFederateApp.owns = func(c context.Context, iri *url.URL) bool {
		return true
	}
	app.MockFederateApp.get = func(c context.Context, iri *url.URL, rw RWType) (PubObject, error) {
		v := &vocab.Person{}
		v.AppendNameString("Sally")
		v.SetId(sallyIRI)
		v.SetLikedOrderedCollection(&vocab.OrderedCollection{})
		return v, nil
	}
	app.MockFederateApp.set = func(c context.Context, o PubObject) error {
		return nil
	}
	socialCb.like = func(c context.Context, s *streams.Like) error {
		return nil
	}
	gotHttpDo := 0
	var httpDeliveryRequest *http.Request
	httpClient.do = func(req *http.Request) (*http.Response, error) {
		gotHttpDo++
		if gotHttpDo == 1 {
			actorResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer(samActorJSON)),
			}
			return actorResp, nil
		} else if gotHttpDo == 2 {
			actorResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer(sallyActorJSON)),
			}
			return actorResp, nil
		} else if gotHttpDo == 3 {
			httpDeliveryRequest = req
			okResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte{})),
			}
			return okResp, nil
		}
		return nil, nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotHttpDo != 3 {
		t.Fatalf("expected %d, got %d", 3, gotHttpDo)
	} else if httpDeliveryRequest.Method != "POST" {
		t.Fatalf("expected %s, got %s", "POST", httpDeliveryRequest.Method)
	} else if s := httpDeliveryRequest.URL.String(); s != samIRIInboxString {
		t.Fatalf("expected %s, got %s", samIRIInboxString, s)
	}
}

func TestPostOutbox_Undo_CallsCallback(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testUndoLike)))))
	gotCallback := 0
	var gotCallbackObject *streams.Undo
	socialCb.undo = func(c context.Context, s *streams.Undo) error {
		gotCallback++
		gotCallbackObject = s
		return nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotCallback != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCallback)
	} else if err := PubObjectEquals(gotCallbackObject.Raw(), testClientExpectedUndo); err != nil {
		t.Fatalf("unexpected callback object: %s", err)
	}
}

func TestPostOutbox_Undo_IsDelivered(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testUndoLike)))))
	socialCb.undo = func(c context.Context, s *streams.Undo) error {
		return nil
	}
	gotHttpDo := 0
	var httpDeliveryRequest *http.Request
	httpClient.do = func(req *http.Request) (*http.Response, error) {
		gotHttpDo++
		if gotHttpDo == 1 {
			actorResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer(samActorJSON)),
			}
			return actorResp, nil
		} else if gotHttpDo == 2 {
			actorResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer(sallyActorJSON)),
			}
			return actorResp, nil
		} else if gotHttpDo == 3 {
			httpDeliveryRequest = req
			okResp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte{})),
			}
			return okResp, nil
		}
		return nil, nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotHttpDo != 3 {
		t.Fatalf("expected %d, got %d", 3, gotHttpDo)
	} else if httpDeliveryRequest.Method != "POST" {
		t.Fatalf("expected %s, got %s", "POST", httpDeliveryRequest.Method)
	} else if s := httpDeliveryRequest.URL.String(); s != samIRIInboxString {
		t.Fatalf("expected %s, got %s", samIRIInboxString, s)
	}
}

func TestPostOutbox_Block_CallsCallback(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testBlock)))))
	gotCallback := 0
	var gotCallbackObject *streams.Block
	socialCb.block = func(c context.Context, s *streams.Block) error {
		gotCallback++
		gotCallbackObject = s
		return nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotCallback != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCallback)
	} else if err := PubObjectEquals(gotCallbackObject.Raw(), testClientExpectedBlock); err != nil {
		t.Fatalf("unexpected callback object: %s", err)
	}
}

func TestPostOutbox_Block_IsNotDelivered(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testBlock)))))
	socialCb.block = func(c context.Context, s *streams.Block) error {
		return nil
	}
	httpClient.do = func(req *http.Request) (*http.Response, error) {
		t.Fatalf("expected no calls to httpClient.Do")
		return nil, nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	}
}

func TestPostOutbox_SetsLocationHeader(t *testing.T) {
	app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p := NewPubberTest(t)
	PreparePubberPostOutboxTest(t, app, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()
	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testCreateNote)))))
	socialCb.create = func(c context.Context, s *streams.Create) error {
		return nil
	}
	handled, err := p.PostOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if loc, ok := resp.HeaderMap["Location"]; !ok {
		t.Fatalf("expected Location header, got none")
	} else if len(loc) != 1 {
		t.Fatalf("expected Location header to have length 1, got %d", len(loc))
	} else if loc[0] != testNewIRIString {
		t.Fatalf("expected %s, got %s", testNewIRIString, loc[0])
	}
}

func TestGetOutbox_RejectNonActivityPub(t *testing.T) {
	app, _, _, _, _, _, _, p := NewPubberTest(t)
	resp := httptest.NewRecorder()
	req := httptest.NewRequest("GET", testOutboxURI, nil)
	app.MockFederateApp.getOutbox = func(c context.Context, r *http.Request, rw RWType) (vocab.OrderedCollectionType, error) {
		return &vocab.OrderedCollection{}, nil
	}
	handled, err := p.GetOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if handled {
		t.Fatalf("expected !handled, got handled")
	}
}

func TestGetOutbox_SetsContentTypeHeader(t *testing.T) {
	app, _, _, _, _, _, _, p := NewPubberTest(t)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("GET", testOutboxURI, nil))
	app.MockFederateApp.getOutbox = func(c context.Context, r *http.Request, rw RWType) (vocab.OrderedCollectionType, error) {
		return &vocab.OrderedCollection{}, nil
	}
	handled, err := p.GetOutbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if l := len(resp.HeaderMap["Content-Type"]); l != 1 {
		t.Fatalf("expected %d, got %d", 1, l)
	} else if h := resp.HeaderMap["Content-Type"][0]; h != responseContentTypeHeader {
		t.Fatalf("expected %s, got %s", responseContentTypeHeader, h)
	}
}

func TestIssue75(t *testing.T) {
	// Root cause: Setting a maxDeliveryDepth to zero means no C2S messages
	// will be sent to federate peers.
	clock := &MockClock{now}
	app := &MockApplication{t: t}
	socialApp := &MockSocialApp{t: t}
	fedApp := &MockFederateApp{MockApplication: app, t: t}
	appl := &MockSocialFederateApp{MockSocialApp: socialApp, MockFederateApp: fedApp}
	socialCb := &MockCallbacker{t: t}
	fedCb := &MockCallbacker{t: t}
	d := &MockDeliverer{t: t}
	httpClient := &MockHttpClient{t: t}
	p := NewPubber(clock, appl, socialCb, fedCb, d, httpClient, testAgent /*maxDeliveryDepth=*/, 0, 1)

	PreparePubberPostOutboxTest(t, appl, socialApp, fedApp, socialCb, fedCb, d, httpClient, p)
	resp := httptest.NewRecorder()

	// Data as described in the issue
	jamesIRI, err := url.Parse("http://example.com/activity/person/james")
	if err != nil {
		t.Fatal(err)
	}
	jamesIRIOutbox, err := url.Parse("http://example.com/activity/person/james/outbox")
	if err != nil {
		t.Fatal(err)
	}
	jessIRI, err := url.Parse("http://example.com/activity/person/jess")
	if err != nil {
		t.Fatal(err)
	}
	publicIRI, err := url.Parse("https://www.w3.org/ns/activitystreams#Public")
	if err != nil {
		t.Fatal(err)
	}
	docIRI, err := url.Parse("https://www.w3.org/TR/activitypub/")
	if err != nil {
		t.Fatal(err)
	}
	payload := &vocab.Document{}
	payload.AppendCcIRI(jessIRI)
	payload.AppendToIRI(publicIRI)
	payload.AppendNameString("ActivityPub")
	payload.AppendUrlAnyURI(docIRI)

	req := Sign(ActivityPubRequest(httptest.NewRequest("POST", jamesIRIOutbox.String(), bytes.NewBuffer(MustSerialize(payload)))))
	gotActorIRI := 0
	socialApp.actorIRI = func(c context.Context, r *http.Request) (*url.URL, error) {
		gotActorIRI++
		return jamesIRI, nil
	}
	gotCallback := 0
	var gotCallbackObject *streams.Create
	socialCb.create = func(c context.Context, s *streams.Create) error {
		gotCallback++
		gotCallbackObject = s
		return nil
	}
	httpClient.do = nil

	handled, err := p.PostOutbox(context.Background(), resp, req)
	expectedPayload := &vocab.Document{}
	expectedPayload.AppendCcIRI(jessIRI)
	expectedPayload.AppendToIRI(publicIRI)
	expectedPayload.AppendNameString("ActivityPub")
	expectedPayload.AppendUrlAnyURI(docIRI)
	expectedPayload.AppendAttributedToIRI(jamesIRI)
	expectedPayload.SetId(testNewIRI2)
	expectedCreate := &vocab.Create{}
	expectedCreate.AppendActorIRI(jamesIRI)
	expectedCreate.AppendObject(expectedPayload)
	expectedCreate.AppendType("Create")
	expectedCreate.SetId(testNewIRI)
	expectedCreate.AppendCcIRI(jessIRI)
	expectedCreate.AppendToIRI(publicIRI)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotActorIRI != 1 {
		t.Fatalf("expected %d, got %d", 1, gotActorIRI)
	} else if gotCallback != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCallback)
	} else if err := PubObjectEquals(gotCallbackObject.Raw(), expectedCreate); err != nil {
		t.Fatalf("unexpected callback object: %s", err)
	}
}

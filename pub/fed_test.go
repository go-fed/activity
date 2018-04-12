package pub

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-fed/activity/streams"
	"github.com/go-fed/activity/vocab"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

const (
	iriString             = "https://example.com/something"
	noteURIString         = "https://example.com/note/123"
	noteActivityURIString = "https://example.com/activity/987"
	testAgent             = "test agent string"
	testInboxURI          = "https://example.com/sally/inbox"
	testOutboxURI         = "https://example.com/sally/outbox"
	testNewIRIString      = "https://example.com/test/new/iri"
	sallyIRIString        = "https://example.com/sally"
	samIRIString          = "https://example.com/sam"
	samIRIInboxString     = "https://example.com/sam/inbox"
	samIRIFollowersString = "https://example.com/sam/followers"
	sallyIRIInboxString   = "https://example.com/sally/inbox"
	noteName              = "A Note"
)

var (
	iri                         *url.URL
	noteIRI                     *url.URL
	noteActivityIRI             *url.URL
	testNewIRI                  *url.URL
	sallyIRI                    *url.URL
	sallyIRIInbox               *url.URL
	sallyActor                  *vocab.Person
	sallyActorJSON              []byte
	samIRI                      *url.URL
	samIRIInbox                 *url.URL
	samIRIFollowers             *url.URL
	samActor                    *vocab.Person
	samActorJSON                []byte
	testNote                    *vocab.Note
	testSingleOrderedCollection *vocab.OrderedCollection
	testCreateNote              *vocab.Create
	testUpdateNote              *vocab.Update
	testDeleteNote              *vocab.Delete
	testTombstoneNote           *vocab.Tombstone
	testFollow                  *vocab.Follow
	testAcceptNote              *vocab.Accept
	testAcceptFollow            *vocab.Accept
	testRejectFollow            *vocab.Reject
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
	testNewIRI, err = url.Parse(testNewIRIString)
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
	samActor = &vocab.Person{}
	samActor.SetInboxAnyURI(*samIRIInbox)
	samActor.SetId(*samIRI)
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
	sallyActor.AddNameString("Sally")
	sallyActor.SetId(*sallyIRI)
	sallyActor.SetInboxAnyURI(*sallyInbox)
	sallyActor.SetOutboxAnyURI(*sallyOutbox)
	m, err = sallyActor.Serialize()
	if err != nil {
		panic(err)
	}
	sallyActorJSON, err = json.Marshal(m)
	if err != nil {
		panic(err)
	}
	testNote = &vocab.Note{}
	testNote.SetId(*noteIRI)
	testNote.AddNameString(noteName)
	testNote.AddContentString("This is a simple note")
	testSingleOrderedCollection = &vocab.OrderedCollection{}
	testSingleOrderedCollection.AddItemsObject(testNote)
	testCreateNote = &vocab.Create{}
	testCreateNote.SetId(*noteActivityIRI)
	testCreateNote.AddSummaryString("Sally created a note")
	testCreateNote.AddActorObject(sallyActor)
	testCreateNote.AddObject(testNote)
	testCreateNote.AddToObject(samActor)
	testUpdateNote = &vocab.Update{}
	testUpdateNote.SetId(*noteActivityIRI)
	testUpdateNote.AddSummaryString("Sally updated a note")
	testUpdateNote.AddActorObject(sallyActor)
	testUpdateNote.AddObject(testNote)
	testUpdateNote.AddToObject(samActor)
	testDeleteNote = &vocab.Delete{}
	testDeleteNote.SetId(*noteActivityIRI)
	testDeleteNote.AddActorObject(sallyActor)
	testDeleteNote.AddObject(testNote)
	testDeleteNote.AddToObject(samActor)
	testTombstoneNote = &vocab.Tombstone{}
	testTombstoneNote.SetId(*noteIRI)
	testTombstoneNote.AddFormerTypeString("Note")
	testTombstoneNote.SetDeleted(now)
	testFollow = &vocab.Follow{}
	testFollow.SetId(*noteActivityIRI)
	testFollow.AddActorObject(sallyActor)
	testFollow.AddObject(samActor)
	testFollow.AddToObject(samActor)
	testAcceptNote = &vocab.Accept{}
	testAcceptNote.SetId(*noteActivityIRI)
	testAcceptNote.AddActorObject(sallyActor)
	testAcceptNote.AddObject(&vocab.Offer{})
	testAcceptNote.AddToObject(samActor)
	testAcceptFollow = &vocab.Accept{}
	testAcceptFollow.SetId(*noteActivityIRI)
	testAcceptFollow.AddActorObject(samActor)
	testAcceptFollow.AddObject(testFollow)
	testAcceptFollow.AddToObject(sallyActor)
	testRejectFollow = &vocab.Reject{}
	testRejectFollow.SetId(*noteActivityIRI)
	testRejectFollow.AddActorObject(samActor)
	testRejectFollow.AddObject(testFollow)
	testRejectFollow.AddToObject(sallyActor)
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
	now = time.Date(2000, 2, 3, 4, 5, 6, 7, Must(time.LoadLocation("America/New_York")))
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
	t                    *testing.T
	owns                 func(c context.Context, id url.URL) bool
	get                  func(c context.Context, id url.URL) (PubObject, error)
	set                  func(c context.Context, o PubObject) error
	getInbox             func(c context.Context, r *http.Request) (vocab.OrderedCollectionType, error)
	getOutbox            func(c context.Context, r *http.Request) (vocab.OrderedCollectionType, error)
	postOutboxAuthorized func(c context.Context, r *http.Request) (bool, error)
	newId                func(c context.Context, t Typer) url.URL
	actorIRI             func(c context.Context, r *http.Request) (url.URL, error)
}

func (m *MockApplication) Owns(c context.Context, id url.URL) bool {
	if m.owns == nil {
		m.t.Fatal("unexpected call to MockApplication Owns")
	}
	return m.owns(c, id)
}

func (m *MockApplication) Get(c context.Context, id url.URL) (PubObject, error) {
	if m.get == nil {
		m.t.Fatal("unexpected call to MockApplication Get")
	}
	return m.get(c, id)
}

func (m *MockApplication) Set(c context.Context, o PubObject) error {
	if m.set == nil {
		m.t.Fatal("unexpected call to MockApplication Set")
	}
	return m.set(c, o)
}

func (m *MockApplication) GetInbox(c context.Context, r *http.Request) (vocab.OrderedCollectionType, error) {
	if m.getInbox == nil {
		m.t.Fatal("unexpected call to MockApplication GetInbox")
	}
	return m.getInbox(c, r)
}

func (m *MockApplication) GetOutbox(c context.Context, r *http.Request) (vocab.OrderedCollectionType, error) {
	if m.getOutbox == nil {
		m.t.Fatal("unexpected call to MockApplication GetOutbox")
	}
	return m.getOutbox(c, r)
}

func (m *MockApplication) PostOutboxAuthorized(c context.Context, r *http.Request) (bool, error) {
	if m.postOutboxAuthorized == nil {
		m.t.Fatal("unexpected call to MockApplication PostOutboxAuthorized")
	}
	return m.postOutboxAuthorized(c, r)
}

func (m *MockApplication) NewId(c context.Context, t Typer) url.URL {
	if m.newId == nil {
		m.t.Fatal("unexpected call to MockApplication NewId")
	}
	return m.newId(c, t)
}

func (m *MockApplication) ActorIRI(c context.Context, r *http.Request) (url.URL, error) {
	if m.actorIRI == nil {
		m.t.Fatal("unexpected call to MockApplication ActorIRI")
	}
	return m.actorIRI(c, r)
}

var _ SocialApp = &MockSocialApp{}

type MockSocialApp struct {
	t         *testing.T
	canAdd    func(c context.Context, o vocab.ObjectType, t vocab.ObjectType) bool
	canRemove func(c context.Context, o vocab.ObjectType, t vocab.ObjectType) bool
}

func (m *MockSocialApp) CanAdd(c context.Context, o vocab.ObjectType, t vocab.ObjectType) bool {
	if m.canAdd == nil {
		m.t.Fatal("unexpected call to MockSocialApp CanAdd")
	}
	return m.canAdd(c, o, t)
}

func (m *MockSocialApp) CanRemove(c context.Context, o vocab.ObjectType, t vocab.ObjectType) bool {
	if m.canRemove == nil {
		m.t.Fatal("unexpected call to MockSocialApp CanRemove")
	}
	return m.canRemove(c, o, t)
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
	if m.Remove == nil {
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

var _ FederateApp = &MockFederateApp{}

type MockFederateApp struct {
	t            *testing.T
	canAdd       func(c context.Context, obj vocab.ObjectType, target vocab.ObjectType) bool
	canRemove    func(c context.Context, obj vocab.ObjectType, target vocab.ObjectType) bool
	onFollow     func(c context.Context, s *streams.Follow) FollowResponse
	unblocked    func(c context.Context, actorIRIs []url.URL) error
	getFollowing func(c context.Context, actor url.URL) (vocab.CollectionType, error)
}

func (m *MockFederateApp) CanAdd(c context.Context, obj vocab.ObjectType, target vocab.ObjectType) bool {
	if m.canAdd == nil {
		m.t.Fatal("unexpected call to MockFederateApp CanAdd")
	}
	return m.canAdd(c, obj, target)
}

func (m *MockFederateApp) CanRemove(c context.Context, obj vocab.ObjectType, target vocab.ObjectType) bool {
	if m.canRemove == nil {
		m.t.Fatal("unexpected call to MockFederateApp CanRemove")
	}
	return m.canRemove(c, obj, target)
}

func (m *MockFederateApp) OnFollow(c context.Context, s *streams.Follow) FollowResponse {
	if m.onFollow == nil {
		m.t.Fatal("unexpected call to MockFederateApp OnFollow")
	}
	return m.onFollow(c, s)
}

func (m *MockFederateApp) Unblocked(c context.Context, actorIRIs []url.URL) error {
	if m.unblocked == nil {
		m.t.Fatal("unexpected call to MockFederateApp Unblocked")
	}
	return m.unblocked(c, actorIRIs)
}

func (m *MockFederateApp) GetFollowing(c context.Context, actor url.URL) (vocab.CollectionType, error) {
	if m.getFollowing == nil {
		m.t.Fatal("unexpected call to MockFederateApp GetFollowing")
	}
	return m.getFollowing(c, actor)
}

var _ Deliverer = &MockDeliverer{}

type MockDeliverer struct {
	t  *testing.T
	do func(b []byte, to url.URL, toDo func(b []byte, u url.URL) error)
}

func (m *MockDeliverer) Do(b []byte, to url.URL, toDo func(b []byte, u url.URL) error) {
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

func NewSocialPubberTest(t *testing.T) (app *MockApplication, socialApp *MockSocialApp, cb *MockCallbacker, p Pubber) {
	clock := &MockClock{now}
	app = &MockApplication{t: t}
	socialApp = &MockSocialApp{t: t}
	cb = &MockCallbacker{t: t}
	p = NewSocialPubber(clock, app, socialApp, cb)
	return
}

func NewFederatingPubberTest(t *testing.T) (app *MockApplication, fedApp *MockFederateApp, cb *MockCallbacker, d *MockDeliverer, h *MockHttpClient, p Pubber) {
	clock := &MockClock{now}
	app = &MockApplication{t: t}
	fedApp = &MockFederateApp{t: t}
	cb = &MockCallbacker{t: t}
	d = &MockDeliverer{t: t}
	h = &MockHttpClient{t: t}
	p = NewFederatingPubber(clock, app, fedApp, cb, d, h, testAgent, 1)
	return
}

func NewPubberTest(t *testing.T) (app *MockApplication, socialApp *MockSocialApp, fedApp *MockFederateApp, socialCb, fedCb *MockCallbacker, d *MockDeliverer, h *MockHttpClient, p Pubber) {
	clock := &MockClock{now}
	app = &MockApplication{t: t}
	socialApp = &MockSocialApp{t: t}
	fedApp = &MockFederateApp{t: t}
	socialCb = &MockCallbacker{t: t}
	fedCb = &MockCallbacker{t: t}
	d = &MockDeliverer{t: t}
	h = &MockHttpClient{t: t}
	p = NewPubber(clock, app, socialApp, fedApp, socialCb, fedCb, d, h, testAgent, 1)
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
	app.getInbox = func(c context.Context, r *http.Request) (vocab.OrderedCollectionType, error) {
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
	app, _, cb, p := NewSocialPubberTest(t)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testCreateNote))))
	gotPostOutboxAuthorized := 0
	app.postOutboxAuthorized = func(c context.Context, r *http.Request) (bool, error) {
		gotPostOutboxAuthorized++
		return true, nil
	}
	gotNewId := 0
	app.newId = func(c context.Context, t Typer) url.URL {
		gotNewId++
		return *testNewIRI
	}
	gotOutbox := 0
	app.getOutbox = func(c context.Context, r *http.Request) (vocab.OrderedCollectionType, error) {
		gotOutbox++
		oc := &vocab.OrderedCollection{}
		oc.AddType("OrderedCollection")
		return oc, nil
	}
	gotSet := 0
	var gotSetOutbox PubObject
	var gotSetCreateObject PubObject
	app.set = func(c context.Context, o PubObject) error {
		gotSet++
		if gotSet == 1 {
			gotSetOutbox = o
		} else if gotSet == 2 {
			gotSetCreateObject = o
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
	} else if gotPostOutboxAuthorized != 1 {
		t.Fatalf("expected %d, got %d", 1, gotPostOutboxAuthorized)
	} else if gotNewId != 1 {
		t.Fatalf("expected %d, got %d", 1, gotNewId)
	} else if gotOutbox != 1 {
		t.Fatalf("expected %d, got %d", 1, gotOutbox)
	} else if gotSet != 2 {
		t.Fatalf("expected %d, got %d", 2, gotSet)
	} else if l := gotSetOutbox.GetType(0).(string); l != "OrderedCollection" {
		t.Fatalf("expected %s, got %s", "OrderedCollection", l)
	} else if l := gotSetCreateObject.GetType(0).(string); l != "Note" {
		t.Fatalf("expected %s, got %s", "Note", l)
	} else if gotCreate != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCreate)
	} else if iri := gotCreateCallback.Raw().GetActorObject(0).GetId(); iri.String() != sallyIRIString {
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
	app.getOutbox = func(c context.Context, r *http.Request) (vocab.OrderedCollectionType, error) {
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
	var iri url.URL
	fedApp.unblocked = func(c context.Context, actorIRIs []url.URL) error {
		gotUnblocked++
		iri = actorIRIs[0]
		return nil
	}
	gotSet := 0
	var setObject PubObject
	app.set = func(c context.Context, o PubObject) error {
		gotSet++
		setObject = o
		return nil
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
	} else if gotSet != 1 {
		t.Fatalf("expected %d, got %d", 1, gotSet)
	} else if l := setObject.GetType(0).(string); l != "Note" {
		t.Fatalf("expected %s, got %s", "Note", l)
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
	app.getInbox = func(c context.Context, r *http.Request) (vocab.OrderedCollectionType, error) {
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
	app.getOutbox = func(c context.Context, r *http.Request) (vocab.OrderedCollectionType, error) {
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
	var iri url.URL
	fedApp.unblocked = func(c context.Context, actorIRIs []url.URL) error {
		gotUnblocked++
		iri = actorIRIs[0]
		return nil
	}
	gotSet := 0
	var setObject PubObject
	app.set = func(c context.Context, o PubObject) error {
		gotSet++
		setObject = o
		return nil
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
	} else if gotSet != 1 {
		t.Fatalf("expected %d, got %d", 1, gotSet)
	} else if l := setObject.GetType(0).(string); l != "Note" {
		t.Fatalf("expected %s, got %s", "Note", l)
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
	app.getInbox = func(c context.Context, r *http.Request) (vocab.OrderedCollectionType, error) {
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
	app, _, _, socialCb, _, d, httpClient, p := NewPubberTest(t)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testOutboxURI, bytes.NewBuffer(MustSerialize(testCreateNote))))
	gotPostOutboxAuthorized := 0
	app.postOutboxAuthorized = func(c context.Context, r *http.Request) (bool, error) {
		gotPostOutboxAuthorized++
		return true, nil
	}
	gotNewId := 0
	app.newId = func(c context.Context, t Typer) url.URL {
		gotNewId++
		return *testNewIRI
	}
	gotOutbox := 0
	app.getOutbox = func(c context.Context, r *http.Request) (vocab.OrderedCollectionType, error) {
		gotOutbox++
		oc := &vocab.OrderedCollection{}
		oc.AddType("OrderedCollection")
		return oc, nil
	}
	gotSet := 0
	var gotSetOutbox PubObject
	var gotSetCreateObject PubObject
	app.set = func(c context.Context, o PubObject) error {
		gotSet++
		if gotSet == 1 {
			gotSetOutbox = o
		} else if gotSet == 2 {
			gotSetCreateObject = o
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
	var doDeliveryURL url.URL
	d.do = func(b []byte, u url.URL, toDo func(b []byte, u url.URL) error) {
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
	} else if gotPostOutboxAuthorized != 1 {
		t.Fatalf("expected %d, got %d", 1, gotPostOutboxAuthorized)
	} else if gotNewId != 1 {
		t.Fatalf("expected %d, got %d", 1, gotNewId)
	} else if gotOutbox != 1 {
		t.Fatalf("expected %d, got %d", 1, gotOutbox)
	} else if gotSet != 2 {
		t.Fatalf("expected %d, got %d", 2, gotSet)
	} else if l := gotSetOutbox.GetType(0).(string); l != "OrderedCollection" {
		t.Fatalf("expected %s, got %s", "OrderedCollection", l)
	} else if l := gotSetCreateObject.GetType(0).(string); l != "Note" {
		t.Fatalf("expected %s, got %s", "Note", l)
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
}

func TestPubber_GetOutbox(t *testing.T) {
	app, _, _, _, _, _, _, p := NewPubberTest(t)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("GET", testOutboxURI, nil))
	gotOutbox := 0
	app.getOutbox = func(c context.Context, r *http.Request) (vocab.OrderedCollectionType, error) {
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
	var iri url.URL
	fedApp.unblocked = func(c context.Context, actorIRIs []url.URL) error {
		gotBlocked++
		iri = actorIRIs[0]
		return blockedErr
	}
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != blockedErr {
		t.Fatalf("expected %s, got %s", blockedErr, err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
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
				v.SetId(*noteActivityIRI)
				v.AddSummaryString("Sally created a note")
				v.AddActorObject(sallyActor)
				v.AddToObject(samActor)
				return v
			},
		},
		{
			name: "update",
			input: func() vocab.Serializer {
				v := &vocab.Update{}
				v.SetId(*noteActivityIRI)
				v.AddSummaryString("Sally updated a note")
				v.AddActorObject(sallyActor)
				v.AddToObject(samActor)
				return v
			},
		},
		{
			name: "delete",
			input: func() vocab.Serializer {
				v := &vocab.Delete{}
				v.SetId(*noteActivityIRI)
				v.AddActorObject(sallyActor)
				v.AddToObject(samActor)
				return v
			},
		},
		{
			name: "follow",
			input: func() vocab.Serializer {
				v := &vocab.Follow{}
				v.SetId(*noteActivityIRI)
				v.AddActorObject(sallyActor)
				v.AddToObject(samActor)
				return v
			},
		},
		{
			name: "add",
			input: func() vocab.Serializer {
				v := &vocab.Add{}
				v.SetId(*noteActivityIRI)
				v.AddActorObject(sallyActor)
				v.AddToObject(samActor)
				v.AddTargetObject(testNote)
				return v
			},
		},
		{
			name: "remove",
			input: func() vocab.Serializer {
				v := &vocab.Remove{}
				v.SetId(*noteActivityIRI)
				v.AddActorObject(sallyActor)
				v.AddToObject(samActor)
				v.AddTargetObject(testNote)
				return v
			},
		},
		{
			name: "like",
			input: func() vocab.Serializer {
				v := &vocab.Like{}
				v.SetId(*noteActivityIRI)
				v.AddActorObject(sallyActor)
				v.AddToObject(samActor)
				return v
			},
		},
		{
			name: "block",
			input: func() vocab.Serializer {
				v := &vocab.Block{}
				v.SetId(*noteActivityIRI)
				v.AddActorObject(sallyActor)
				v.AddToObject(samActor)
				return v
			},
		},
		{
			name: "undo",
			input: func() vocab.Serializer {
				v := &vocab.Undo{}
				v.SetId(*noteActivityIRI)
				v.AddActorObject(sallyActor)
				v.AddToObject(samActor)
				return v
			},
		},
	}
	_, _, fedApp, _, _, _, _, p := NewPubberTest(t)
	fedApp.unblocked = func(c context.Context, actorIRIs []url.URL) error {
		return nil
	}
	for _, test := range tests {
		resp := httptest.NewRecorder()
		req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(test.input()))))
		handled, err := p.PostInbox(context.Background(), resp, req)
		if err != ErrObjectRequired {
			t.Fatalf("(%s) expected %s, got %s", test.name, ErrObjectRequired, err)
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
				v.SetId(*noteActivityIRI)
				v.AddActorObject(sallyActor)
				v.AddToObject(samActor)
				v.AddObject(testNote)
				return v
			},
		},
		{
			name: "remove",
			input: func() vocab.Serializer {
				v := &vocab.Remove{}
				v.SetId(*noteActivityIRI)
				v.AddActorObject(sallyActor)
				v.AddToObject(samActor)
				v.AddObject(testNote)
				return v
			},
		},
	}
	_, _, fedApp, _, _, _, _, p := NewPubberTest(t)
	fedApp.unblocked = func(c context.Context, actorIRIs []url.URL) error {
		return nil
	}
	for _, test := range tests {
		resp := httptest.NewRecorder()
		req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(test.input()))))
		handled, err := p.PostInbox(context.Background(), resp, req)
		if err != ErrTargetRequired {
			t.Fatalf("(%s) expected %s, got %s", test.name, ErrTargetRequired, err)
		} else if !handled {
			t.Fatalf("(%s) expected handled, got !handled", test.name)
		}
	}
}

func TestPostInbox_Create_SetsObject(t *testing.T) {
	app, _, fedApp, _, fedCb, _, _, p := NewPubberTest(t)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testCreateNote))))
	fedApp.unblocked = func(c context.Context, actorIRIs []url.URL) error {
		return nil
	}
	gotSet := 0
	var setObject PubObject
	app.set = func(c context.Context, o PubObject) error {
		gotSet++
		setObject = o
		return nil
	}
	fedCb.create = func(c context.Context, s *streams.Create) error {
		return nil
	}
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotSet != 1 {
		t.Fatalf("expected %d, got %d", 1, gotSet)
	} else if err := PubObjectEquals(setObject, testNote); err != nil {
		t.Fatalf("unexpected set object: %s", err)
	}
}

func TestPostInbox_Create_CallsCallback(t *testing.T) {
	app, _, fedApp, _, fedCb, _, _, p := NewPubberTest(t)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testCreateNote))))
	fedApp.unblocked = func(c context.Context, actorIRIs []url.URL) error {
		return nil
	}
	app.set = func(c context.Context, o PubObject) error {
		return nil
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
	} else if gotCreate != 1 {
		t.Fatalf("expected %d, got %d", 1, gotCreate)
	} else if err := PubObjectEquals(gotCreateCallback.Raw(), testCreateNote); err != nil {
		t.Fatalf("unexpected create callback: %s", err)
	}
}

func TestPostInbox_Update_SetsObject(t *testing.T) {
	app, _, fedApp, _, fedCb, _, _, p := NewPubberTest(t)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testUpdateNote))))
	fedApp.unblocked = func(c context.Context, actorIRIs []url.URL) error {
		return nil
	}
	gotSet := 0
	var setObject PubObject
	app.set = func(c context.Context, o PubObject) error {
		gotSet++
		setObject = o
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
	} else if gotSet != 1 {
		t.Fatalf("expected %d, got %d", 1, gotSet)
	} else if err := PubObjectEquals(setObject, testNote); err != nil {
		t.Fatalf("unexpected set object: %s", err)
	}
}

func TestPostInbox_Update_CallsCallback(t *testing.T) {
	app, _, fedApp, _, fedCb, _, _, p := NewPubberTest(t)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testUpdateNote))))
	fedApp.unblocked = func(c context.Context, actorIRIs []url.URL) error {
		return nil
	}
	app.set = func(c context.Context, o PubObject) error {
		return nil
	}
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
	app, _, fedApp, _, fedCb, _, _, p := NewPubberTest(t)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testDeleteNote))))
	fedApp.unblocked = func(c context.Context, actorIRIs []url.URL) error {
		return nil
	}
	app.get = func(c context.Context, id url.URL) (PubObject, error) {
		if id != *noteIRI {
			t.Fatalf("expected %s, got %s", noteIRI, id)
		}
		return testNote, nil
	}
	app.set = func(c context.Context, p PubObject) error {
		return nil
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
				testNote.SetId(*noteIRI)
				testNote.AddType("Note")
				testNote.AddNameString(noteName)
				testNote.AddContentString("This is a simple note")
				testNote.SetPublished(now)
				return testNote
			},
			expected: func() vocab.Serializer {
				testTombstoneNote := &vocab.Tombstone{}
				testTombstoneNote.SetId(*noteIRI)
				testTombstoneNote.AddFormerTypeString("Note")
				testTombstoneNote.SetDeleted(now)
				testTombstoneNote.SetPublished(now)
				return testTombstoneNote
			},
		},
		{
			name: "forward published iri",
			input: func() PubObject {
				testNote := &vocab.Note{}
				testNote.SetId(*noteIRI)
				testNote.AddType("Note")
				testNote.AddNameString(noteName)
				testNote.AddContentString("This is a simple note")
				testNote.SetPublishedIRI(*iri)
				return testNote
			},
			expected: func() vocab.Serializer {
				testTombstoneNote := &vocab.Tombstone{}
				testTombstoneNote.SetId(*noteIRI)
				testTombstoneNote.AddFormerTypeString("Note")
				testTombstoneNote.SetDeleted(now)
				testTombstoneNote.SetPublishedIRI(*iri)
				return testTombstoneNote
			},
		},
		{
			name: "forward updated time",
			input: func() PubObject {
				testNote := &vocab.Note{}
				testNote.SetId(*noteIRI)
				testNote.AddType("Note")
				testNote.AddNameString(noteName)
				testNote.AddContentString("This is a simple note")
				testNote.SetUpdated(now)
				return testNote
			},
			expected: func() vocab.Serializer {
				testTombstoneNote := &vocab.Tombstone{}
				testTombstoneNote.SetId(*noteIRI)
				testTombstoneNote.AddFormerTypeString("Note")
				testTombstoneNote.SetDeleted(now)
				testTombstoneNote.SetUpdated(now)
				return testTombstoneNote
			},
		},
		{
			name: "forward updated iri",
			input: func() PubObject {
				testNote := &vocab.Note{}
				testNote.SetId(*noteIRI)
				testNote.AddType("Note")
				testNote.AddNameString(noteName)
				testNote.AddContentString("This is a simple note")
				testNote.SetUpdatedIRI(*iri)
				return testNote
			},
			expected: func() vocab.Serializer {
				testTombstoneNote := &vocab.Tombstone{}
				testTombstoneNote.SetId(*noteIRI)
				testTombstoneNote.AddFormerTypeString("Note")
				testTombstoneNote.SetDeleted(now)
				testTombstoneNote.SetUpdatedIRI(*iri)
				return testTombstoneNote
			},
		},
	}
	app, _, fedApp, _, fedCb, _, _, p := NewPubberTest(t)
	fedApp.unblocked = func(c context.Context, actorIRIs []url.URL) error {
		return nil
	}
	gotSet := 0
	var gotSetObject PubObject
	app.set = func(c context.Context, p PubObject) error {
		gotSet++
		gotSetObject = p
		return nil
	}
	fedCb.delete = func(c context.Context, s *streams.Delete) error {
		return nil
	}
	for _, test := range tests {
		app.get = func(c context.Context, id url.URL) (PubObject, error) {
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
		} else if gotSet != 1 {
			t.Fatalf("(%q) expected %d, got %d", 1, test.name, gotSet)
		} else if err := PubObjectEquals(gotSetObject, test.expected()); err != nil {
			t.Fatalf("(%q) unexpected tombstone object: %s", test.name, err)
		}
	}
}

func TestPostInbox_Delete_CallsCallback(t *testing.T) {
	app, _, fedApp, _, fedCb, _, _, p := NewPubberTest(t)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testDeleteNote))))
	fedApp.unblocked = func(c context.Context, actorIRIs []url.URL) error {
		return nil
	}
	app.get = func(c context.Context, id url.URL) (PubObject, error) {
		return testNote, nil
	}
	app.set = func(c context.Context, p PubObject) error {
		return nil
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
	_, _, fedApp, _, fedCb, _, _, p := NewPubberTest(t)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testFollow))))
	fedApp.unblocked = func(c context.Context, actorIRIs []url.URL) error {
		return nil
	}
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
	_, _, fedApp, _, fedCb, d, httpClient, p := NewPubberTest(t)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testFollow))))
	fedApp.unblocked = func(c context.Context, actorIRIs []url.URL) error {
		return nil
	}
	gotOnFollow := 0
	fedApp.onFollow = func(c context.Context, s *streams.Follow) FollowResponse {
		gotOnFollow++
		return AutomaticReject
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
	var doDeliveryURL url.URL
	var bytesToSend []byte
	d.do = func(b []byte, u url.URL, toDo func(b []byte, u url.URL) error) {
		gotDoDelivery++
		doDeliveryURL = u
		bytesToSend = b
		if err := toDo(b, u); err != nil {
			t.Fatalf("Unexpected error in MockDeliverer.Do: %s", err)
		}
	}
	expected := &vocab.Reject{}
	expected.AddObject(testFollow)
	expected.AddToObject(sallyActor)
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotOnFollow != 1 {
		t.Fatalf("expected %d, got %d", 1, gotOnFollow)
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
	}
}

// TODO: Test follower OrderedCollection & IRI.
// TODO: Test does not own one of the objects.
func TestPostInbox_Follow_AutoAccept(t *testing.T) {
	app, _, fedApp, _, fedCb, d, httpClient, p := NewPubberTest(t)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testFollow))))
	fedApp.unblocked = func(c context.Context, actorIRIs []url.URL) error {
		return nil
	}
	gotOnFollow := 0
	fedApp.onFollow = func(c context.Context, s *streams.Follow) FollowResponse {
		gotOnFollow++
		return AutomaticAccept
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
	var doDeliveryURL url.URL
	var bytesToSend []byte
	d.do = func(b []byte, u url.URL, toDo func(b []byte, u url.URL) error) {
		gotDoDelivery++
		doDeliveryURL = u
		bytesToSend = b
		if err := toDo(b, u); err != nil {
			t.Fatalf("Unexpected error in MockDeliverer.Do: %s", err)
		}
	}
	gotOwns := 0
	var ownsIRI url.URL
	app.owns = func(c context.Context, id url.URL) bool {
		gotOwns++
		ownsIRI = id
		return true
	}
	gotGet := 0
	var getIRI url.URL
	app.get = func(c context.Context, id url.URL) (PubObject, error) {
		gotGet++
		getIRI = id
		samActor := &vocab.Person{}
		samActor.SetInboxAnyURI(*samIRIInbox)
		samActor.SetId(*samIRI)
		samActor.SetFollowersCollection(&vocab.Collection{})
		return samActor, nil
	}
	gotSet := 0
	var setObject PubObject
	app.set = func(c context.Context, o PubObject) error {
		gotSet++
		setObject = o
		return nil
	}
	expected := &vocab.Accept{}
	expected.AddObject(testFollow)
	expected.AddToObject(sallyActor)
	expectedFollowers := &vocab.Collection{}
	expectedFollowers.AddItemsObject(sallyActor)
	handled, err := p.PostInbox(context.Background(), resp, req)
	if err != nil {
		t.Fatal(err)
	} else if !handled {
		t.Fatalf("expected handled, got !handled")
	} else if gotOnFollow != 1 {
		t.Fatalf("expected %d, got %d", 1, gotOnFollow)
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
	} else if gotSet != 1 {
		t.Fatalf("expected %d, got %d", 1, gotSet)
	} else if err := PubObjectEquals(setObject, expectedFollowers); err != nil {
		t.Fatal(err)
	}
}

func TestPostInbox_Follow_CallsCallback(t *testing.T) {
	_, _, fedApp, _, fedCb, _, _, p := NewPubberTest(t)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testFollow))))
	fedApp.unblocked = func(c context.Context, actorIRIs []url.URL) error {
		return nil
	}
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
	_, _, fedApp, _, fedCb, _, _, p := NewPubberTest(t)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testAcceptNote))))
	fedApp.unblocked = func(c context.Context, actorIRIs []url.URL) error {
		return nil
	}
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

// TODO: Test follower OrderedCollection & IRI.
func TestPostInbox_Accept_AcceptFollowAddsToFollowersIfOwned(t *testing.T) {
	app, _, fedApp, _, fedCb, _, _, p := NewPubberTest(t)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testAcceptFollow))))
	fedApp.unblocked = func(c context.Context, actorIRIs []url.URL) error {
		return nil
	}
	gotOwns := 0
	var ownsIRI url.URL
	app.owns = func(c context.Context, id url.URL) bool {
		gotOwns++
		ownsIRI = id
		return true
	}
	gotGet := 0
	var getIRI url.URL
	app.get = func(c context.Context, id url.URL) (PubObject, error) {
		gotGet++
		getIRI = id
		sallyActor := &vocab.Person{}
		sallyActor.SetInboxAnyURI(*sallyIRIInbox)
		sallyActor.SetId(*sallyIRI)
		sallyActor.SetFollowingCollection(&vocab.Collection{})
		return sallyActor, nil
	}
	gotSet := 0
	var setObject PubObject
	app.set = func(c context.Context, o PubObject) error {
		gotSet++
		setObject = o
		return nil
	}
	fedCb.accept = func(c context.Context, s *streams.Accept) error {
		return nil
	}
	expectedFollowing := &vocab.Collection{}
	expectedFollowing.AddItemsObject(samActor)
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
	} else if gotSet != 1 {
		t.Fatalf("expected %d, got %d", 1, gotSet)
	} else if err := PubObjectEquals(setObject, expectedFollowing); err != nil {
		t.Fatal(err)
	}
}

func TestPostInbox_Accept_DoesNothingIfNotOwned(t *testing.T) {
	app, _, fedApp, _, fedCb, _, _, p := NewPubberTest(t)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testAcceptFollow))))
	fedApp.unblocked = func(c context.Context, actorIRIs []url.URL) error {
		return nil
	}
	gotOwns := 0
	var ownsIRI url.URL
	app.owns = func(c context.Context, id url.URL) bool {
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
	app, _, fedApp, _, fedCb, _, _, p := NewPubberTest(t)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testAcceptFollow))))
	fedApp.unblocked = func(c context.Context, actorIRIs []url.URL) error {
		return nil
	}
	app.owns = func(c context.Context, id url.URL) bool {
		return true
	}
	app.get = func(c context.Context, id url.URL) (PubObject, error) {
		sallyActor := &vocab.Person{}
		sallyActor.SetInboxAnyURI(*sallyIRIInbox)
		sallyActor.SetId(*sallyIRI)
		sallyActor.SetFollowingCollection(&vocab.Collection{})
		return sallyActor, nil
	}
	app.set = func(c context.Context, o PubObject) error {
		return nil
	}
	gotCallback := 0
	var gotCallbackObject *streams.Accept
	fedCb.accept = func(c context.Context, s *streams.Accept) error {
		gotCallback++
		gotCallbackObject = s
		return nil
	}
	expectedFollowing := &vocab.Collection{}
	expectedFollowing.AddItemsObject(samActor)
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
	_, _, fedApp, _, fedCb, _, _, p := NewPubberTest(t)
	resp := httptest.NewRecorder()
	req := ActivityPubRequest(httptest.NewRequest("POST", testInboxURI, bytes.NewBuffer(MustSerialize(testRejectFollow))))
	fedApp.unblocked = func(c context.Context, actorIRIs []url.URL) error {
		return nil
	}
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
	// TODO: Implement
}

func TestPostInbox_Add_AddIfTargetOwned(t *testing.T) {
	// TODO: Implement
}

func TestPostInbox_Add_DoesNotAddIfAppCannotAdd(t *testing.T) {
	// TODO: Implement
}

func TestPostInbox_Add_AddIfAppCanAdd(t *testing.T) {
	// TODO: Implement
}

func TestPostInbox_Add_SetsTarget(t *testing.T) {
	// TODO: Implement
}

func TestPostInbox_Add_CallsCallback(t *testing.T) {
	// TODO: Implement
}

func TestPostInbox_Remove_DoesNotAddIfTargetNotOwned(t *testing.T) {
	// TODO: Implement
}

func TestPostInbox_Remove_AddIfTargetOwned(t *testing.T) {
	// TODO: Implement
}

func TestPostInbox_Remove_DoesNotAddIfAppCannotAdd(t *testing.T) {
	// TODO: Implement
}

func TestPostInbox_Remove_AddIfAppCanAdd(t *testing.T) {
	// TODO: Implement
}

func TestPostInbox_Remove_SetsTarget(t *testing.T) {
	// TODO: Implement
}

func TestPostInbox_Remove_CallsCallback(t *testing.T) {
	// TODO: Implement
}

func TestPostInbox_Like_AddsToLikeCollection(t *testing.T) {
	// TODO: Implement
}

func TestPostInbox_Like_CallsCallback(t *testing.T) {
	// TODO: Implement
}

func TestPostInbox_Undo_CallsCallback(t *testing.T) {
	// TODO: Implement
}

func TestGetInbox_RejectNonActivityPub(t *testing.T) {
	// TODO: Implement
}

func TestGetInbox_SetsContentTypeHeader(t *testing.T) {
	// TODO: Implement
}

func TestGetInbox_DeduplicateInboxItems(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_RejectNonActivityPub(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_RejectUnauthorized(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_WrapInCreateActivity(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Create_RequireObject(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Create_CopyToAttributedTo(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Create_CopyRecipientsToObject(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Create_SetCreatedObject(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Create_CallsCallback(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Create_IsDelivered(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Update_RequireObject(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Update_OverwriteUpdatedFields(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Update_SetUpdatedObject(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Update_CallsCallback(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Update_IsDelivered(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Delete_RequireObject(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Delete_SetsTombstone(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Delete_CallsCallback(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Delete_IsDelivered(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Follow_RequireObject(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Follow_CallsCallback(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Follow_IsDelivered(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Accept_CallsCallback(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Accept_IsDelivered(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Reject_CallsCallback(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Reject_IsDelivered(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Add_RequireObject(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Add_RequireTarget(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Add_DoesNotAddIfTargetNotOwned(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Add_AddsIfTargetOwned(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Add_DoesNotAddIfAppCannotAdd(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Add_AddIfAppCanAdd(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Add_SetsTarget(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Add_CallsCallback(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Add_IsDelivered(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Remove_RequireObject(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Remove_RequireTarget(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Remove_DoesNotRemoveIfTargetNotOwned(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Remove_RemoveIfTargetOwned(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Remove_DoesNotRemoveIfAppCannotRemove(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Remove_RemoveIfAppCanRemove(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Remove_SetsTarget(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Remove_CallsCallback(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Remove_IsDelivered(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Like_RequireObject(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Like_RequireTarget(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Like_AddsToLikedCollection(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Like_AddsToLikedOrderedCollection(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Like_CallsCallback(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Like_IsDelivered(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Undo_RequireObject(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Undo_CallsCallback(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Undo_IsDelivered(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Block_RequireObject(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Block_CallsCallback(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_Block_IsNotDelivered(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_DoesNotDeliverNondeliverable(t *testing.T) {
	// TODO: Implement
}

func TestPostOutbox_SetsLocationHeader(t *testing.T) {
	// TODO: Implement
}

func TestGetOutbox_RejectNonActivityPub(t *testing.T) {
	// TODO: Implement
}

func TestGetOutbox_SetsContentTypeHeader(t *testing.T) {
	// TODO: Implement
}

package pub

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-fed/activity/streams"
	"github.com/go-fed/activity/streams/vocab"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

const (
	testMyInboxIRI            = "https://example.com/addison/inbox"
	testMyOutboxIRI           = "https://example.com/addison/outbox"
	testFederatedActivityIRI  = "https://other.example.com/activity/1"
	testFederatedActivityIRI2 = "https://other.example.com/activity/2"
	testFederatedActorIRI     = "https://other.example.com/dakota"
	testFederatedActorIRI2    = "https://other.example.com/addison"
	testNoteId1               = "https://example.com/note/1"
	testNoteId2               = "https://example.com/note/2"
	testNewActivityIRI        = "https://example.com/new/1"
	testToIRI = "https://maybe.example.com/to/1"
	testToIRI2 = "https://maybe.example.com/to/2"
	testCcIRI = "https://maybe.example.com/cc/1"
	testCcIRI2 = "https://maybe.example.com/cc/2"
	testAudienceIRI = "https://maybe.example.com/audience/1"
	testAudienceIRI2 = "https://maybe.example.com/audience/2"
	testPersonIRI = "https://maybe.example.com/person"
	testServiceIRI = "https://maybe.example.com/service"
)

// mustParse parses a URL or panics.
func mustParse(s string) *url.URL {
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	return u
}

// assertEqual ensures two values are equal.
func assertEqual(t *testing.T, a, b interface{}) {
	if a != b {
		t.Errorf("expected equal: %v != %v", a, b)
	}
}

// assertByteEqual ensures two byte slices are equal.
func assertByteEqual(t *testing.T, a, b []byte) {
	if string(a) != string(b) {
		t.Errorf("expected equal:\n%s\n\n%s", a, b)
	}
}

// assertNotEqual ensures two values are not equal.
func assertNotEqual(t *testing.T, a, b interface{}) {
	if a == b {
		t.Errorf("expected not equal: %v != %v", a, b)
	}
}

var (
	// testErr is a test error.
	testErr = errors.New("test error")
	// testNote is a test Note from a federated peer.
	testFederatedNote vocab.ActivityStreamsNote
	// testNote is a test Note owned by this server.
	testMyNote vocab.ActivityStreamsNote
	// testCreate is a test Create Activity.
	testCreate vocab.ActivityStreamsCreate
	// testCreate2 is a test Create Activity with two actors.
	testCreate2 vocab.ActivityStreamsCreate
	// testCreateNoId is a test Create Activity without an 'id' set.
	testCreateNoId vocab.ActivityStreamsCreate
	// testOrderedCollectionUniqueElems is a collection with only unique
	// ids.
	testOrderedCollectionUniqueElems vocab.ActivityStreamsOrderedCollectionPage
	// testOrderedCollectionUniqueElemsString is the JSON-LD version of the
	// testOrderedCollectionUniqueElems value
	testOrderedCollectionUniqueElemsString string
	// testOrderedCollectionDupedElems is a collection with duplicated ids.
	testOrderedCollectionDupedElems vocab.ActivityStreamsOrderedCollectionPage
	// testOrderedCollectionDedupedElemsString is the JSON-LD version of the
	// testOrderedCollectionDedupedElems value with duplicates removed
	testOrderedCollectionDedupedElemsString string
	// testEmptyOrderedCollection is an empty OrderedCollectionPage.
	testEmptyOrderedCollection vocab.ActivityStreamsOrderedCollectionPage
	// testOrderedCollectionWithFederatedId has the federated Activity id.
	testOrderedCollectionWithFederatedId vocab.ActivityStreamsOrderedCollectionPage
	// testListen is a test Listen Activity.
	testListen vocab.ActivityStreamsListen
	// testOrderedCollectionWithFederatedId2 has the second federated
	// Activity id.
	testOrderedCollectionWithFederatedId2 vocab.ActivityStreamsOrderedCollectionPage
	// testOrderedCollectionWithBothFederatedIds has both federated Activity id.
	testOrderedCollectionWithBothFederatedIds vocab.ActivityStreamsOrderedCollectionPage
	// testPerson is a Person.
	testPerson vocab.ActivityStreamsPerson
	// testService is a Service.
	testService vocab.ActivityStreamsService
)

// The test data cannot be created at init time since that is when the hooks of
// the `streams` package are set up. So initialize the data in this call instead
// of at init time.
func setupData() {
	// testFederatedNote
	func() {
		testFederatedNote = streams.NewActivityStreamsNote()
		name := streams.NewActivityStreamsNameProperty()
		name.AppendXMLSchemaString("A Federated Note")
		testFederatedNote.SetActivityStreamsName(name)
		content := streams.NewActivityStreamsContentProperty()
		content.AppendXMLSchemaString("This is a simple note being federated.")
		testFederatedNote.SetActivityStreamsContent(content)
		id := streams.NewActivityStreamsIdProperty()
		id.Set(mustParse(testNoteId1))
		testFederatedNote.SetActivityStreamsId(id)
	}()
	// testMyNote
	func() {
		testMyNote = streams.NewActivityStreamsNote()
		name := streams.NewActivityStreamsNameProperty()
		name.AppendXMLSchemaString("My Note")
		testMyNote.SetActivityStreamsName(name)
		content := streams.NewActivityStreamsContentProperty()
		content.AppendXMLSchemaString("This is a simple note of mine.")
		testMyNote.SetActivityStreamsContent(content)
	}()
	// testCreate
	func() {
		testCreate = streams.NewActivityStreamsCreate()
		id := streams.NewActivityStreamsIdProperty()
		id.Set(mustParse(testFederatedActivityIRI))
		testCreate.SetActivityStreamsId(id)
		actor := streams.NewActivityStreamsActorProperty()
		actor.AppendIRI(mustParse(testFederatedActorIRI))
		testCreate.SetActivityStreamsActor(actor)
		op := streams.NewActivityStreamsObjectProperty()
		op.AppendActivityStreamsNote(testFederatedNote)
		testCreate.SetActivityStreamsObject(op)
	}()
	// testCreate2
	func() {
		testCreate2 = streams.NewActivityStreamsCreate()
		id := streams.NewActivityStreamsIdProperty()
		id.Set(mustParse(testFederatedActivityIRI))
		testCreate2.SetActivityStreamsId(id)
		actor := streams.NewActivityStreamsActorProperty()
		actor.AppendIRI(mustParse(testFederatedActorIRI))
		actor.AppendIRI(mustParse(testFederatedActorIRI2))
		testCreate2.SetActivityStreamsActor(actor)
		op := streams.NewActivityStreamsObjectProperty()
		op.AppendActivityStreamsNote(testFederatedNote)
		testCreate2.SetActivityStreamsObject(op)
	}()
	// testCreateNoId
	func() {
		testCreateNoId = streams.NewActivityStreamsCreate()
		actor := streams.NewActivityStreamsActorProperty()
		actor.AppendIRI(mustParse(testFederatedActorIRI))
		testCreateNoId.SetActivityStreamsActor(actor)
		op := streams.NewActivityStreamsObjectProperty()
		op.AppendActivityStreamsNote(testFederatedNote)
		testCreateNoId.SetActivityStreamsObject(op)
	}()
	// testOrderedCollectionUniqueElems and
	// testOrderedCollectionUniqueElemsString
	func() {
		testOrderedCollectionUniqueElems = streams.NewActivityStreamsOrderedCollectionPage()
		oi := streams.NewActivityStreamsOrderedItemsProperty()
		oi.AppendIRI(mustParse(testNoteId1))
		oi.AppendIRI(mustParse(testNoteId2))
		testOrderedCollectionUniqueElems.SetActivityStreamsOrderedItems(oi)
		testOrderedCollectionUniqueElemsString = `{"@context":"https://www.w3.org/TR/activitystreams-vocabulary","orderedItems":["https://example.com/note/1","https://example.com/note/2"],"type":"OrderedCollectionPage"}`
	}()
	// testOrderedCollectionDupedElems and
	// testOrderedCollectionDedupedElemsString
	func() {
		testOrderedCollectionDupedElems = streams.NewActivityStreamsOrderedCollectionPage()
		oi := streams.NewActivityStreamsOrderedItemsProperty()
		oi.AppendIRI(mustParse(testNoteId1))
		oi.AppendIRI(mustParse(testNoteId1))
		testOrderedCollectionDupedElems.SetActivityStreamsOrderedItems(oi)
		testOrderedCollectionDedupedElemsString = `{"@context":"https://www.w3.org/TR/activitystreams-vocabulary","orderedItems":"https://example.com/note/1","type":"OrderedCollectionPage"}`
	}()
	// testEmptyOrderedCollection
	func() {
		testEmptyOrderedCollection = streams.NewActivityStreamsOrderedCollectionPage()
	}()
	// testOrderedCollectionWithFederatedId
	func() {
		testOrderedCollectionWithFederatedId = streams.NewActivityStreamsOrderedCollectionPage()
		oi := streams.NewActivityStreamsOrderedItemsProperty()
		oi.AppendIRI(mustParse(testFederatedActivityIRI))
		testOrderedCollectionWithFederatedId.SetActivityStreamsOrderedItems(oi)
	}()
	// testListen
	func() {
		testListen = streams.NewActivityStreamsListen()
		id := streams.NewActivityStreamsIdProperty()
		id.Set(mustParse(testFederatedActivityIRI))
		testListen.SetActivityStreamsId(id)
		actor := streams.NewActivityStreamsActorProperty()
		actor.AppendIRI(mustParse(testFederatedActorIRI))
		testListen.SetActivityStreamsActor(actor)
		op := streams.NewActivityStreamsObjectProperty()
		op.AppendActivityStreamsNote(testFederatedNote)
		testListen.SetActivityStreamsObject(op)
	}()
	// testOrderedCollectionWithFederatedId2
	func() {
		testOrderedCollectionWithFederatedId2 = streams.NewActivityStreamsOrderedCollectionPage()
		oi := streams.NewActivityStreamsOrderedItemsProperty()
		oi.AppendIRI(mustParse(testFederatedActivityIRI2))
		testOrderedCollectionWithFederatedId2.SetActivityStreamsOrderedItems(oi)
	}()
	// testOrderedCollectionWithBothFederatedIds
	func() {
		testOrderedCollectionWithBothFederatedIds = streams.NewActivityStreamsOrderedCollectionPage()
		oi := streams.NewActivityStreamsOrderedItemsProperty()
		oi.AppendIRI(mustParse(testFederatedActivityIRI))
		oi.AppendIRI(mustParse(testFederatedActivityIRI2))
		testOrderedCollectionWithBothFederatedIds.SetActivityStreamsOrderedItems(oi)
	}()
	// testPerson
	func () {
		testPerson = streams.NewActivityStreamsPerson()
		id := streams.NewActivityStreamsIdProperty()
		id.Set(mustParse(testPersonIRI))
		testPerson.SetActivityStreamsId(id)
	}()
	// testService
	func () {
		testService = streams.NewActivityStreamsService()
		id := streams.NewActivityStreamsIdProperty()
		id.Set(mustParse(testServiceIRI))
		testService.SetActivityStreamsId(id)
	} ()
}

// wrappedInCreate returns a Create activity wrapping the given type.
func wrappedInCreate(t vocab.Type) vocab.ActivityStreamsCreate {
	create := streams.NewActivityStreamsCreate()
	op := streams.NewActivityStreamsObjectProperty()
	op.AppendType(t)
	create.SetActivityStreamsObject(op)
	return create
}

// mustSerialize serializes a type or panics.
func mustSerialize(t vocab.Type) map[string]interface{} {
	m, err := serialize(t)
	if err != nil {
		panic(err)
	}
	return m
}

// toDeserializedForm serializes and deserializes a type so that it works with
// mock expectations.
func toDeserializedForm(t vocab.Type) vocab.Type {
	m := mustSerialize(t)
	asValue, err := streams.ToType(context.Background(), m)
	if err != nil {
		panic(err)
	}
	return asValue
}

// withNewId sets a new id property on the activity
func withNewId(t vocab.Type) Activity {
	a, ok := t.(Activity)
	if !ok {
		panic("activity streams value is not an Activity")
	}
	id := streams.NewActivityStreamsIdProperty()
	id.Set(mustParse(testNewActivityIRI))
	a.SetActivityStreamsId(id)
	return a
}

// now returns the "current" time for tests.
func now() time.Time {
	l, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}
	return time.Date(2000, 2, 3, 4, 5, 6, 7, l)
}

// nowDateHeader returns the "current" time formatted in a form expected by the
// Date header in HTTP responses.
func nowDateHeader() string {
	return now().UTC().Format("Mon, 02 Jan 2006 15:04:05") + " GMT"
}

// toAPRequests adds the appropriate Content-Type or Accept headers to indicate
// that the HTTP request is an ActivityPub one. Also sets the Date header with
// the "current" test time.
func toAPRequest(r *http.Request) *http.Request {
	if r.Method == "POST" {
		existing, ok := r.Header[contentTypeHeader]
		if ok {
			r.Header[contentTypeHeader] = append(existing, activityStreamsMediaTypes[0])
		} else {
			r.Header[contentTypeHeader] = []string{activityStreamsMediaTypes[0]}
		}
	} else if r.Method == "GET" {
		existing, ok := r.Header[acceptHeader]
		if ok {
			r.Header[acceptHeader] = append(existing, activityStreamsMediaTypes[0])
		} else {
			r.Header[acceptHeader] = []string{activityStreamsMediaTypes[0]}
		}
	} else {
		panic("cannot toAPRequest with method " + r.Method)
	}
	r.Header[dateHeader] = []string{now().UTC().Format("Mon, 02 Jan 2006 15:04:05") + " GMT"}
	return r
}

// toPostInboxRequest creates a new POST HTTP request with the given type as
// the payload.
func toPostInboxRequest(t vocab.Type) *http.Request {
	m, err := serialize(t)
	if err != nil {
		panic(err)
	}
	b, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(b)
	return httptest.NewRequest("POST", testMyInboxIRI, buf)
}

// toPostOutboxRequest creates a new POST HTTP request with the given type as
// the payload.
func toPostOutboxRequest(t vocab.Type) *http.Request {
	m, err := serialize(t)
	if err != nil {
		panic(err)
	}
	b, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(b)
	return httptest.NewRequest("POST", testMyOutboxIRI, buf)
}

// toPostOutboxUnknownRequest creates a new POST HTTP request with an unknown
// type in the payload.
func toPostOutboxUnknownRequest() *http.Request {
	s := `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "http://www.types.example/ProductOffer",
  "id": "http://www.example.com/spam"
}`
	b := []byte(s)
	buf := bytes.NewBuffer(b)
	return httptest.NewRequest("POST", testMyOutboxIRI, buf)
}

// toPostInboxUnknownRequest creates a new POST HTTP request with an unknown
// type in the payload.
func toPostInboxUnknownRequest() *http.Request {
	s := `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "http://www.types.example/ProductOffer",
  "id": "http://www.example.com/spam"
}`
	b := []byte(s)
	buf := bytes.NewBuffer(b)
	return httptest.NewRequest("POST", testMyInboxIRI, buf)
}

// toGetInboxRequest creates a new GET HTTP request.
func toGetInboxRequest() *http.Request {
	return httptest.NewRequest("GET", testMyInboxIRI, nil)
}

// toGetOutboxRequest creates a new GET HTTP request.
func toGetOutboxRequest() *http.Request {
	return httptest.NewRequest("GET", testMyOutboxIRI, nil)
}

// addToIds adds two IRIs to the 'to' property
func addToIds(t Activity) Activity {
	to := streams.NewActivityStreamsToProperty()
	to.AppendIRI(mustParse(testToIRI))
	to.AppendIRI(mustParse(testToIRI2))
	t.SetActivityStreamsTo(to)
	return t
}

// mustAddCcIds adds two IRIs to the 'cc' property
func mustAddCcIds(t Activity) Activity {
	if ccer, ok := t.(ccer); ok {
		cc := streams.NewActivityStreamsCcProperty()
		cc.AppendIRI(mustParse(testCcIRI))
		cc.AppendIRI(mustParse(testCcIRI2))
		ccer.SetActivityStreamsCc(cc)
	} else {
		panic(fmt.Sprintf("activity is not ccer: %T", t))
	}
	return t
}

// mustAddAudienceIds adds two IRIs to the 'audience' property
func mustAddAudienceIds(t Activity) Activity {
	if audiencer, ok := t.(audiencer); ok {
		aud := streams.NewActivityStreamsAudienceProperty()
		aud.AppendIRI(mustParse(testAudienceIRI))
		aud.AppendIRI(mustParse(testAudienceIRI2))
		audiencer.SetActivityStreamsAudience(aud)
	} else {
		panic(fmt.Sprintf("activity is not audiencer: %T", t))
	}
	return t
}

package streams

import (
	"context"
	"encoding/json"
	"github.com/go-fed/activity/streams/vocab"
	"github.com/go-test/deep"
	"net/url"
	"testing"
)

type serializer interface {
	Serialize() (map[string]interface{}, error)
}

// IsKnownResolverError returns true if it is known that an example from
// GetTestTable will trigger a JSONResolver error.
func IsKnownResolverError(t TestTable) (b bool, reason string) {
	if t.name == "Example 61" {
		b = true
		reason = "no \"type\" property is on the root object"
	} else if t.name == "Example 62" {
		b = true
		reason = "an unknown \"type\" property is on the root object"
	} else if t.name == "Example 153" {
		b = true
		reason = "no \"type\" property is on the root object"
	}
	return
}

func TestJSONResolver(t *testing.T) {
	for _, example := range GetTestTable() {
		if skip, reason := IsKnownResolverError(example); skip {
			t.Logf("Skipping table test case %q as it's known an error will be returned because %s", example.name, reason)
			continue
		}
		name := example.name
		t.Logf("Testing table test case %q", name)
		ex := example.expectedJSON
		resFn := func(s serializer) error {
			m, err := s.Serialize()
			if err != nil {
				return err
			}
			m["@context"] = "https://www.w3.org/ns/activitystreams"
			actual, err := json.Marshal(m)
			if diff, err := GetJSONDiff(actual, []byte(ex)); err == nil && diff != nil {
				t.Errorf("%s: Serialize JSON equality is false:\n%s", name, diff)
			} else if err != nil {
				t.Errorf("%s: GetJSONDiff returned error: %s", name, err)
			}
			return nil
		}
		r, err := NewJSONResolver(
			func(c context.Context, x vocab.ActivityStreamsAccept) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsActivity) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsAdd) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsAnnounce) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsApplication) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsArrive) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsArticle) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsAudio) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsBlock) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsCollection) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsCollectionPage) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsCreate) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsDelete) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsDislike) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsDocument) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsEvent) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsFlag) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsFollow) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsGroup) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsIgnore) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsImage) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsIntransitiveActivity) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsInvite) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsJoin) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsLeave) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsLike) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsLink) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsListen) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsMention) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsMove) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsNote) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsObject) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsOffer) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsOrderedCollection) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsOrderedCollectionPage) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsOrganization) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsPage) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsPerson) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsPlace) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsProfile) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsQuestion) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsRead) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsReject) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsRelationship) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsRemove) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsService) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsTentativeAccept) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsTentativeReject) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsTombstone) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsTravel) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsUndo) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsUpdate) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsVideo) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsView) error {
				return resFn(x)
			},
		)
		if err != nil {
			t.Errorf("%s: Cannot create JSONResolver: %s", name, err)
			continue
		}
		m := make(map[string]interface{})
		err = json.Unmarshal([]byte(ex), &m)
		if err != nil {
			t.Errorf("%s: Cannot json.Unmarshal: %s", name, err)
			continue
		}
		err = r.Resolve(context.Background(), m)
		if err != nil {
			t.Errorf("%s: Cannot JSONResolver.Deserialize: %s", name, err)
			continue
		}
	}
}

func TestJSONResolverErrors(t *testing.T) {
	for _, example := range GetTestTable() {
		isError, reason := IsKnownResolverError(example)
		if !isError {
			continue
		}
		name := example.name
		t.Logf("Testing table test case %q", name)
		ex := example.expectedJSON
		resFn := func(s serializer) error { return nil }
		r, err := NewJSONResolver(
			func(c context.Context, x vocab.ActivityStreamsAccept) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsActivity) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsAdd) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsAnnounce) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsApplication) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsArrive) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsArticle) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsAudio) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsBlock) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsCollection) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsCollectionPage) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsCreate) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsDelete) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsDislike) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsDocument) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsEvent) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsFlag) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsFollow) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsGroup) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsIgnore) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsImage) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsIntransitiveActivity) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsInvite) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsJoin) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsLeave) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsLike) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsLink) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsListen) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsMention) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsMove) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsNote) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsObject) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsOffer) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsOrderedCollection) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsOrderedCollectionPage) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsOrganization) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsPage) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsPerson) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsPlace) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsProfile) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsQuestion) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsRead) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsReject) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsRelationship) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsRemove) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsService) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsTentativeAccept) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsTentativeReject) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsTombstone) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsTravel) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsUndo) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsUpdate) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsVideo) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityStreamsView) error {
				return resFn(x)
			},
		)
		if err != nil {
			t.Errorf("%s: Cannot create JSONResolver: %s", name, err)
			continue
		}
		m := make(map[string]interface{})
		err = json.Unmarshal([]byte(ex), &m)
		if err != nil {
			t.Errorf("%s: Cannot json.Unmarshal: %s", name, err)
			continue
		}
		err = r.Resolve(context.Background(), m)
		if err == nil {
			t.Errorf("%s: Expected error because %s", name, reason)
		}
	}
}

func TestNulls(t *testing.T) {
	const (
		samIRIInboxString = "https://example.com/sam/inbox"
		samIRIString      = "https://example.com/sam"
		noteURIString     = "https://example.com/note/123"
		sallyIRIString    = "https://example.com/sally"
		activityIRIString = "https://example.com/test/new/iri"
	)
	samIRIInbox, err := url.Parse(samIRIInboxString)
	if err != nil {
		t.Fatal(err)
	}
	samIRI, err := url.Parse(samIRIString)
	if err != nil {
		t.Fatal(err)
	}
	noteIRI, err := url.Parse(noteURIString)
	if err != nil {
		t.Fatal(err)
	}
	sallyIRI, err := url.Parse(sallyIRIString)
	if err != nil {
		t.Fatal(err)
	}
	activityIRI, err := url.Parse(activityIRIString)
	if err != nil {
		t.Fatal(err)
	}
	noteIdProperty := NewActivityStreamsIdProperty()
	noteIdProperty.SetIRI(noteIRI)
	expectedNote := NewActivityStreamsNote()
	expectedNote.SetActivityStreamsId(noteIdProperty)
	noteNameProperty := NewActivityStreamsNameProperty()
	noteNameProperty.AppendXMLSchemaString("A Note")
	expectedNote.SetActivityStreamsName(noteNameProperty)
	noteContentProperty := NewActivityStreamsContentProperty()
	noteContentProperty.AppendXMLSchemaString("This is a simple note")
	expectedNote.SetActivityStreamsContent(noteContentProperty)
	noteToProperty := NewActivityStreamsToProperty()
	expectedSamActor := NewActivityStreamsPerson()
	samInboxProperty := NewActivityStreamsInboxProperty()
	samInboxProperty.SetIRI(samIRIInbox)
	expectedSamActor.SetActivityStreamsInbox(samInboxProperty)
	samIdProperty := NewActivityStreamsIdProperty()
	samIdProperty.SetIRI(samIRI)
	expectedSamActor.SetActivityStreamsId(samIdProperty)
	noteToProperty.AppendActivityStreamsPerson(expectedSamActor)
	expectedNote.SetActivityStreamsTo(noteToProperty)
	expectedUpdate := NewActivityStreamsUpdate()
	sallyIdProperty := NewActivityStreamsIdProperty()
	sallyIdProperty.SetIRI(sallyIRI)
	sallyPerson := NewActivityStreamsPerson()
	sallyPerson.SetActivityStreamsId(sallyIdProperty)
	sallyActor := NewActivityStreamsActorProperty()
	sallyActor.AppendActivityStreamsPerson(sallyPerson)
	expectedUpdate.SetActivityStreamsActor(sallyActor)
	summaryProperty := NewActivityStreamsSummaryProperty()
	summaryProperty.AppendXMLSchemaString("Sally updated her note")
	expectedUpdate.SetActivityStreamsSummary(summaryProperty)
	updateIdProperty := NewActivityStreamsIdProperty()
	updateIdProperty.SetIRI(activityIRI)
	expectedUpdate.SetActivityStreamsId(updateIdProperty)
	objectNote := NewActivityStreamsObjectProperty()
	objectNote.AppendActivityStreamsNote(expectedNote)
	expectedUpdate.SetActivityStreamsObject(objectNote)

	// Variable to aid in deserialization in tests
	var actual serializer
	tables := []struct {
		name              string
		expected          serializer
		callback          interface{}
		input             string
		inputWithoutNulls string
	}{
		{
			name:     "JSON nulls are not preserved",
			expected: expectedUpdate,
			callback: func(c context.Context, v vocab.ActivityStreamsUpdate) error {
				actual = v
				return nil
			},
			input: `
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
			`,
			inputWithoutNulls: `
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
	                      "type": "Person"
                            }
	                  }
                        }
			`,
		},
	}
	for _, r := range tables {
		t.Logf("Testing table test case %q", r.name)
		res, err := NewJSONResolver(r.callback)
		if err != nil {
			t.Errorf("%s: cannot create resolver: %s", r.name, err)
			continue
		}
		m := make(map[string]interface{})
		err = json.Unmarshal([]byte(r.input), &m)
		if err != nil {
			t.Errorf("%s: Cannot json.Unmarshal: %s", r.name, err)
			continue
		}
		err = res.Resolve(context.Background(), m)
		if err != nil {
			t.Errorf("%s: Cannot Deserialize: %s", r.name, err)
			continue
		}
		if diff := deep.Equal(actual, r.expected); diff != nil {
			t.Errorf("%s: Deserialize deep equal is false: %s", r.name, diff)
		}
		m, err = actual.Serialize()
		if err != nil {
			t.Errorf("%s: Cannot Serialize: %s", r.name, err)
			continue
		}
		m["@context"] = "https://www.w3.org/ns/activitystreams"
		reser, err := json.Marshal(m)
		if err != nil {
			t.Errorf("%s: Cannot json.Marshal: %s", r.name, err)
			continue
		}
		if diff, err := GetJSONDiff(reser, []byte(r.inputWithoutNulls)); err == nil && diff != nil {
			t.Errorf("%s: Serialize JSON equality is false:\n%s", r.name, diff)
		} else if err != nil {
			t.Errorf("%s: GetJSONDiff returned error: %s", r.name, err)
		}
	}
}

func GetJSONDiff(str1, str2 []byte) ([]string, error) {
	var i1 interface{}
	var i2 interface{}

	err := json.Unmarshal(str1, &i1)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(str2, &i2)
	if err != nil {
		return nil, err
	}
	return deep.Equal(i1, i2), nil
}

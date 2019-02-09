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

func TestJSONResolver(t *testing.T) {
	for _, example := range allRepoExamples {
		name := example.name
		t.Logf("%s: Testing table test case", name)
		ex := example.example
		resFn := func(s serializer) error {
			m, err := s.Serialize()
			if err != nil {
				return err
			}
			m["@context"] = "http://www.w3.org/ns/activitystreams"
			actual, err := json.Marshal(m)
			if diff, err := GetJSONDiff(actual, []byte(ex)); err == nil && diff != nil {
				t.Errorf("%s: Serialize JSON equality is false:\n%s", name, diff)
			} else if err != nil {
				t.Errorf("%s: GetJSONDiff returned error: %s", name, err)
			}
			return nil
		}
		r, err := NewJSONResolver(
			func(c context.Context, x vocab.AcceptInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ActivityInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.AddInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.AnnounceInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ApplicationInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ArriveInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ArticleInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.AudioInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.BlockInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.CollectionInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.CollectionPageInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.CreateInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.DeleteInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.DislikeInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.DocumentInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.EventInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.FlagInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.FollowInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.GroupInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.IgnoreInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ImageInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.IntransitiveActivityInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.InviteInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.JoinInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.LeaveInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.LikeInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.LinkInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ListenInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.MentionInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.MoveInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.NoteInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ObjectInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.OfferInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.OrderedCollectionInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.OrderedCollectionPageInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.OrganizationInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.PageInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.PersonInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.PlaceInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ProfileInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.QuestionInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ReadInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.RejectInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.RelationshipInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.RemoveInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ServiceInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.TentativeAcceptInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.TentativeRejectInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.TombstoneInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.TravelInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.UndoInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.UpdateInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.VideoInterface) error {
				return resFn(x)
			},
			func(c context.Context, x vocab.ViewInterface) error {
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
		// Examples that needed adjustment:
		// Example 64: Array of one attachment - deserializes OK, but re-serialization does not match
		// Example 68: Array of one bcc - deserializes OK, but re-serialization does not match
		// Example 70: Array of one cc - deserializes OK, but re-serialization does not match
		// Example 112: Array of one item - deserializes OK, but re-serialization does not match
		// Example 118: Array of one tag - deserializes OK, but re-serialization does not match
		// Example 123: Array of one to - deserializes OK, but re-serialization does not match
		// Example 184: missing @context
		// Example 196: '\n' characters were literal newlines in source
		err = r.Resolve(context.Background(), m)
		if err != nil {
			t.Errorf("%s: Cannot JSONResolver.Deserialize: %s", name, err)
			continue
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
	expectedNote.SetId(noteIdProperty)
	noteNameProperty := NewActivityStreamsNameProperty()
	noteNameProperty.AppendString("A Note")
	expectedNote.SetName(noteNameProperty)
	noteContentProperty := NewActivityStreamsContentProperty()
	noteContentProperty.AppendString("This is a simple note")
	expectedNote.SetContent(noteContentProperty)
	noteToProperty := NewActivityStreamsToProperty()
	expectedSamActor := NewActivityStreamsPerson()
	samInboxProperty := NewActivityStreamsInboxProperty()
	samInboxProperty.SetIRI(samIRIInbox)
	expectedSamActor.SetInbox(samInboxProperty)
	samIdProperty := NewActivityStreamsIdProperty()
	samIdProperty.SetIRI(samIRI)
	expectedSamActor.SetId(samIdProperty)
	noteToProperty.AppendPerson(expectedSamActor)
	expectedNote.SetTo(noteToProperty)
	expectedUpdate := NewActivityStreamsUpdate()
	sallyIdProperty := NewActivityStreamsIdProperty()
	sallyIdProperty.SetIRI(sallyIRI)
	sallyPerson := NewActivityStreamsPerson()
	sallyPerson.SetId(sallyIdProperty)
	sallyActor := NewActivityStreamsActorProperty()
	sallyActor.AppendPerson(sallyPerson)
	expectedUpdate.SetActor(sallyActor)
	summaryProperty := NewActivityStreamsSummaryProperty()
	summaryProperty.AppendString("Sally updated her note")
	expectedUpdate.SetSummary(summaryProperty)
	updateIdProperty := NewActivityStreamsIdProperty()
	updateIdProperty.SetIRI(activityIRI)
	expectedUpdate.SetId(updateIdProperty)
	objectNote := NewActivityStreamsObjectProperty()
	objectNote.AppendNote(expectedNote)
	expectedUpdate.SetObject(objectNote)

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
			callback: func(c context.Context, v vocab.UpdateInterface) error {
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
		t.Logf("%s: Testing table test case", r.name)
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

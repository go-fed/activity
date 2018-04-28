//go:generate go install github.com/go-fed/activity/tools/streams
//go:generate streams
package streams

import (
	"net/url"
	"encoding/json"
	"github.com/go-fed/activity/vocab"
	"github.com/go-test/deep"
	"testing"
)

func TestRepoExamples(t *testing.T) {
	for name, ex := range allRepoExamples {
		resFn := func(s vocab.Serializer) error {
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
		r := &Resolver{
			ObjectCallback: func(x *Object) error {
				return resFn(x)
			},
			LinkCallback: func(x *Link) error {
				return resFn(x)
			},
			ActivityCallback: func(x *Activity) error {
				return resFn(x)
			},
			IntransitiveActivityCallback: func(x *IntransitiveActivity) error {
				return resFn(x)
			},
			CollectionCallback: func(x *Collection) error {
				return resFn(x)
			},
			OrderedCollectionCallback: func(x *OrderedCollection) error {
				return resFn(x)
			},
			CollectionPageCallback: func(x *CollectionPage) error {
				return resFn(x)
			},
			OrderedCollectionPageCallback: func(x *OrderedCollectionPage) error {
				return resFn(x)
			},
			AcceptCallback: func(x *Accept) error {
				return resFn(x)
			},
			TentativeAcceptCallback: func(x *TentativeAccept) error {
				return resFn(x)
			},
			AddCallback: func(x *Add) error {
				return resFn(x)
			},
			ArriveCallback: func(x *Arrive) error {
				return resFn(x)
			},
			CreateCallback: func(x *Create) error {
				return resFn(x)
			},
			DeleteCallback: func(x *Delete) error {
				return resFn(x)
			},
			FollowCallback: func(x *Follow) error {
				return resFn(x)
			},
			IgnoreCallback: func(x *Ignore) error {
				return resFn(x)
			},
			JoinCallback: func(x *Join) error {
				return resFn(x)
			},
			LeaveCallback: func(x *Leave) error {
				return resFn(x)
			},
			LikeCallback: func(x *Like) error {
				return resFn(x)
			},
			OfferCallback: func(x *Offer) error {
				return resFn(x)
			},
			InviteCallback: func(x *Invite) error {
				return resFn(x)
			},
			RejectCallback: func(x *Reject) error {
				return resFn(x)
			},
			TentativeRejectCallback: func(x *TentativeReject) error {
				return resFn(x)
			},
			RemoveCallback: func(x *Remove) error {
				return resFn(x)
			},
			UndoCallback: func(x *Undo) error {
				return resFn(x)
			},
			UpdateCallback: func(x *Update) error {
				return resFn(x)
			},
			ViewCallback: func(x *View) error {
				return resFn(x)
			},
			ListenCallback: func(x *Listen) error {
				return resFn(x)
			},
			ReadCallback: func(x *Read) error {
				return resFn(x)
			},
			MoveCallback: func(x *Move) error {
				return resFn(x)
			},
			TravelCallback: func(x *Travel) error {
				return resFn(x)
			},
			AnnounceCallback: func(x *Announce) error {
				return resFn(x)
			},
			BlockCallback: func(x *Block) error {
				return resFn(x)
			},
			FlagCallback: func(x *Flag) error {
				return resFn(x)
			},
			DislikeCallback: func(x *Dislike) error {
				return resFn(x)
			},
			QuestionCallback: func(x *Question) error {
				return resFn(x)
			},
			ApplicationCallback: func(x *Application) error {
				return resFn(x)
			},
			GroupCallback: func(x *Group) error {
				return resFn(x)
			},
			OrganizationCallback: func(x *Organization) error {
				return resFn(x)
			},
			PersonCallback: func(x *Person) error {
				return resFn(x)
			},
			ServiceCallback: func(x *Service) error {
				return resFn(x)
			},
			RelationshipCallback: func(x *Relationship) error {
				return resFn(x)
			},
			ArticleCallback: func(x *Article) error {
				return resFn(x)
			},
			DocumentCallback: func(x *Document) error {
				return resFn(x)
			},
			AudioCallback: func(x *Audio) error {
				return resFn(x)
			},
			ImageCallback: func(x *Image) error {
				return resFn(x)
			},
			VideoCallback: func(x *Video) error {
				return resFn(x)
			},
			NoteCallback: func(x *Note) error {
				return resFn(x)
			},
			PageCallback: func(x *Page) error {
				return resFn(x)
			},
			EventCallback: func(x *Event) error {
				return resFn(x)
			},
			PlaceCallback: func(x *Place) error {
				return resFn(x)
			},
			ProfileCallback: func(x *Profile) error {
				return resFn(x)
			},
			TombstoneCallback: func(x *Tombstone) error {
				return resFn(x)
			},
			MentionCallback: func(x *Mention) error {
				return resFn(x)
			},
		}
		m := make(map[string]interface{})
		err := json.Unmarshal([]byte(ex), &m)
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
		err = r.Deserialize(m)
		if err != nil {
			t.Errorf("%s: Cannot Resolver.Deserialize: %s", name, err)
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
	expectedSamActor := &vocab.Person{}
	expectedSamActor.SetInboxAnyURI(*samIRIInbox)
	expectedSamActor.SetId(*samIRI)
	expectedNote := &vocab.Note{}
	expectedNote.SetId(*noteIRI)
	expectedNote.AddNameString("A Note")
	expectedNote.AddContentString("This is a simple note")
	expectedNote.AddToObject(expectedSamActor)
	expectedUpdate := &vocab.Update{}
	expectedUpdate.AddActorIRI(*sallyIRI)
	expectedUpdate.AddSummaryString("Sally updated her note")
	expectedUpdate.SetId(*activityIRI)
	expectedUpdate.AddObject(expectedNote)
	tables := []struct {
		name         string
		expected     vocab.Serializer
		input        string
	}{
		{
			name:         "JSON with null",
			expected:     expectedUpdate,
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
		},
	}
	for _, r := range tables {
		var actual vocab.ActivityType
		res := &Resolver{
			AnyActivityCallback: func(v vocab.ActivityType) error {
				actual = v
				return nil
			},
		}
		m := make(map[string]interface{})
		err := json.Unmarshal([]byte(r.input), &m)
		if err != nil {
			t.Errorf("%s: Cannot json.Unmarshal: %s", r.name, err)
			continue
		}
		err = res.Deserialize(m)
		if err != nil {
			t.Errorf("%s: Cannot Deserialize: %s", r.name, err)
			continue
		}
		if diff := deep.Equal(actual, r.expected); diff != nil {
			t.Errorf("%s: Deserialize deep equal is false: %s", r.name, diff)
		}
		// Only test ability to reserialize
		serializer, ok := actual.(vocab.Serializer)
		if !ok {
			t.Errorf("Deserializer is not also Serializer")
		}
		m, err = serializer.Serialize()
		if err != nil {
			t.Errorf("%s: Cannot Serialize: %s", r.name, err)
			continue
		}
		m["@context"] = "https://www.w3.org/ns/activitystreams"
		_, err = json.Marshal(m)
		if err != nil {
			t.Errorf("%s: Cannot json.Marshal: %s", r.name, err)
			continue
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

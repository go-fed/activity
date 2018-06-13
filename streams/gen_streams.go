// Package streams is a convenience wrapper around the raw ActivityStream vocabulary. This package is code-generated to permit more powerful expressions and manipulations of the ActivityStreams Vocabulary types. This package also does not permit use of 'unknown' properties, or those that are outside of the ActivityStream Vocabulary specification. However, it still correctly propagates them when repeatedly re-and-de-serialized. Custom extensions of the vocabulary are supported by modifying the data definitions in the generation tool and rerunning it. Do not modify this package directly.
package streams

import (
	"fmt"
	"github.com/go-fed/activity/vocab"
)

type Resolution int

const (
	Resolved Resolution = iota
	RawResolutionNeeded
	Unresolved
)

type Presence int

const (
	NoPresence Presence = iota
	ConvenientPresence
	RawPresence
)

// Resolver contains callback functions to execute when it Deserializes a raw map[string]interface{} into a concrete type. Clients can set only the callbacks they care about and handle the resulting concrete type.
type Resolver struct {
	// Callback function for the Object type
	ObjectCallback func(*Object) error
	// Callback function for the Link type
	LinkCallback func(*Link) error
	// Callback function for the Activity type
	ActivityCallback func(*Activity) error
	// Callback function for the IntransitiveActivity type
	IntransitiveActivityCallback func(*IntransitiveActivity) error
	// Callback function for the Collection type
	CollectionCallback func(*Collection) error
	// Callback function for the OrderedCollection type
	OrderedCollectionCallback func(*OrderedCollection) error
	// Callback function for the CollectionPage type
	CollectionPageCallback func(*CollectionPage) error
	// Callback function for the OrderedCollectionPage type
	OrderedCollectionPageCallback func(*OrderedCollectionPage) error
	// Callback function for the Accept type
	AcceptCallback func(*Accept) error
	// Callback function for the TentativeAccept type
	TentativeAcceptCallback func(*TentativeAccept) error
	// Callback function for the Add type
	AddCallback func(*Add) error
	// Callback function for the Arrive type
	ArriveCallback func(*Arrive) error
	// Callback function for the Create type
	CreateCallback func(*Create) error
	// Callback function for the Delete type
	DeleteCallback func(*Delete) error
	// Callback function for the Follow type
	FollowCallback func(*Follow) error
	// Callback function for the Ignore type
	IgnoreCallback func(*Ignore) error
	// Callback function for the Join type
	JoinCallback func(*Join) error
	// Callback function for the Leave type
	LeaveCallback func(*Leave) error
	// Callback function for the Like type
	LikeCallback func(*Like) error
	// Callback function for the Offer type
	OfferCallback func(*Offer) error
	// Callback function for the Invite type
	InviteCallback func(*Invite) error
	// Callback function for the Reject type
	RejectCallback func(*Reject) error
	// Callback function for the TentativeReject type
	TentativeRejectCallback func(*TentativeReject) error
	// Callback function for the Remove type
	RemoveCallback func(*Remove) error
	// Callback function for the Undo type
	UndoCallback func(*Undo) error
	// Callback function for the Update type
	UpdateCallback func(*Update) error
	// Callback function for the View type
	ViewCallback func(*View) error
	// Callback function for the Listen type
	ListenCallback func(*Listen) error
	// Callback function for the Read type
	ReadCallback func(*Read) error
	// Callback function for the Move type
	MoveCallback func(*Move) error
	// Callback function for the Travel type
	TravelCallback func(*Travel) error
	// Callback function for the Announce type
	AnnounceCallback func(*Announce) error
	// Callback function for the Block type
	BlockCallback func(*Block) error
	// Callback function for the Flag type
	FlagCallback func(*Flag) error
	// Callback function for the Dislike type
	DislikeCallback func(*Dislike) error
	// Callback function for the Question type
	QuestionCallback func(*Question) error
	// Callback function for the Application type
	ApplicationCallback func(*Application) error
	// Callback function for the Group type
	GroupCallback func(*Group) error
	// Callback function for the Organization type
	OrganizationCallback func(*Organization) error
	// Callback function for the Person type
	PersonCallback func(*Person) error
	// Callback function for the Service type
	ServiceCallback func(*Service) error
	// Callback function for the Relationship type
	RelationshipCallback func(*Relationship) error
	// Callback function for the Article type
	ArticleCallback func(*Article) error
	// Callback function for the Document type
	DocumentCallback func(*Document) error
	// Callback function for the Audio type
	AudioCallback func(*Audio) error
	// Callback function for the Image type
	ImageCallback func(*Image) error
	// Callback function for the Video type
	VideoCallback func(*Video) error
	// Callback function for the Note type
	NoteCallback func(*Note) error
	// Callback function for the Page type
	PageCallback func(*Page) error
	// Callback function for the Event type
	EventCallback func(*Event) error
	// Callback function for the Place type
	PlaceCallback func(*Place) error
	// Callback function for the Profile type
	ProfileCallback func(*Profile) error
	// Callback function for the Tombstone type
	TombstoneCallback func(*Tombstone) error
	// Callback function for the Mention type
	MentionCallback func(*Mention) error
	// Callback function for any type that satisfies the vocab.ObjectType interface. Note that this will be called in addition to the specific type callbacks.
	AnyObjectCallback func(vocab.ObjectType) error
	// Callback function for any type that satisfies the vocab.LinkType interface. Note that this will be called in addition to the specific type callbacks.
	AnyLinkCallback func(vocab.LinkType) error
	// Callback function for any type that satisfies the vocab.ActivityType interface. Note that this will be called in addition to the specific type callbacks.
	AnyActivityCallback func(vocab.ActivityType) error
}

// dispatch routes the given type to the appropriate Resolver callback.
func (t *Resolver) dispatch(i interface{}) (handled bool, err error) {
	// Begin generateResolver for type 'Object'
	if rawV, ok := i.(*vocab.Object); ok {
		if t.ObjectCallback != nil {
			v := &Object{raw: rawV}
			return true, t.ObjectCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Object'
	// Begin generateResolver for type 'Link'
	if rawV, ok := i.(*vocab.Link); ok {
		if t.LinkCallback != nil {
			v := &Link{raw: rawV}
			return true, t.LinkCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Link'
	// Begin generateResolver for type 'Activity'
	if rawV, ok := i.(*vocab.Activity); ok {
		if t.ActivityCallback != nil {
			v := &Activity{raw: rawV}
			return true, t.ActivityCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Activity'
	// Begin generateResolver for type 'IntransitiveActivity'
	if rawV, ok := i.(*vocab.IntransitiveActivity); ok {
		if t.IntransitiveActivityCallback != nil {
			v := &IntransitiveActivity{raw: rawV}
			return true, t.IntransitiveActivityCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'IntransitiveActivity'
	// Begin generateResolver for type 'Collection'
	if rawV, ok := i.(*vocab.Collection); ok {
		if t.CollectionCallback != nil {
			v := &Collection{raw: rawV}
			return true, t.CollectionCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Collection'
	// Begin generateResolver for type 'OrderedCollection'
	if rawV, ok := i.(*vocab.OrderedCollection); ok {
		if t.OrderedCollectionCallback != nil {
			v := &OrderedCollection{raw: rawV}
			return true, t.OrderedCollectionCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'OrderedCollection'
	// Begin generateResolver for type 'CollectionPage'
	if rawV, ok := i.(*vocab.CollectionPage); ok {
		if t.CollectionPageCallback != nil {
			v := &CollectionPage{raw: rawV}
			return true, t.CollectionPageCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'CollectionPage'
	// Begin generateResolver for type 'OrderedCollectionPage'
	if rawV, ok := i.(*vocab.OrderedCollectionPage); ok {
		if t.OrderedCollectionPageCallback != nil {
			v := &OrderedCollectionPage{raw: rawV}
			return true, t.OrderedCollectionPageCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'OrderedCollectionPage'
	// Begin generateResolver for type 'Accept'
	if rawV, ok := i.(*vocab.Accept); ok {
		if t.AcceptCallback != nil {
			v := &Accept{raw: rawV}
			return true, t.AcceptCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Accept'
	// Begin generateResolver for type 'TentativeAccept'
	if rawV, ok := i.(*vocab.TentativeAccept); ok {
		if t.TentativeAcceptCallback != nil {
			v := &TentativeAccept{raw: rawV}
			return true, t.TentativeAcceptCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'TentativeAccept'
	// Begin generateResolver for type 'Add'
	if rawV, ok := i.(*vocab.Add); ok {
		if t.AddCallback != nil {
			v := &Add{raw: rawV}
			return true, t.AddCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Add'
	// Begin generateResolver for type 'Arrive'
	if rawV, ok := i.(*vocab.Arrive); ok {
		if t.ArriveCallback != nil {
			v := &Arrive{raw: rawV}
			return true, t.ArriveCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Arrive'
	// Begin generateResolver for type 'Create'
	if rawV, ok := i.(*vocab.Create); ok {
		if t.CreateCallback != nil {
			v := &Create{raw: rawV}
			return true, t.CreateCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Create'
	// Begin generateResolver for type 'Delete'
	if rawV, ok := i.(*vocab.Delete); ok {
		if t.DeleteCallback != nil {
			v := &Delete{raw: rawV}
			return true, t.DeleteCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Delete'
	// Begin generateResolver for type 'Follow'
	if rawV, ok := i.(*vocab.Follow); ok {
		if t.FollowCallback != nil {
			v := &Follow{raw: rawV}
			return true, t.FollowCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Follow'
	// Begin generateResolver for type 'Ignore'
	if rawV, ok := i.(*vocab.Ignore); ok {
		if t.IgnoreCallback != nil {
			v := &Ignore{raw: rawV}
			return true, t.IgnoreCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Ignore'
	// Begin generateResolver for type 'Join'
	if rawV, ok := i.(*vocab.Join); ok {
		if t.JoinCallback != nil {
			v := &Join{raw: rawV}
			return true, t.JoinCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Join'
	// Begin generateResolver for type 'Leave'
	if rawV, ok := i.(*vocab.Leave); ok {
		if t.LeaveCallback != nil {
			v := &Leave{raw: rawV}
			return true, t.LeaveCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Leave'
	// Begin generateResolver for type 'Like'
	if rawV, ok := i.(*vocab.Like); ok {
		if t.LikeCallback != nil {
			v := &Like{raw: rawV}
			return true, t.LikeCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Like'
	// Begin generateResolver for type 'Offer'
	if rawV, ok := i.(*vocab.Offer); ok {
		if t.OfferCallback != nil {
			v := &Offer{raw: rawV}
			return true, t.OfferCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Offer'
	// Begin generateResolver for type 'Invite'
	if rawV, ok := i.(*vocab.Invite); ok {
		if t.InviteCallback != nil {
			v := &Invite{raw: rawV}
			return true, t.InviteCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Invite'
	// Begin generateResolver for type 'Reject'
	if rawV, ok := i.(*vocab.Reject); ok {
		if t.RejectCallback != nil {
			v := &Reject{raw: rawV}
			return true, t.RejectCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Reject'
	// Begin generateResolver for type 'TentativeReject'
	if rawV, ok := i.(*vocab.TentativeReject); ok {
		if t.TentativeRejectCallback != nil {
			v := &TentativeReject{raw: rawV}
			return true, t.TentativeRejectCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'TentativeReject'
	// Begin generateResolver for type 'Remove'
	if rawV, ok := i.(*vocab.Remove); ok {
		if t.RemoveCallback != nil {
			v := &Remove{raw: rawV}
			return true, t.RemoveCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Remove'
	// Begin generateResolver for type 'Undo'
	if rawV, ok := i.(*vocab.Undo); ok {
		if t.UndoCallback != nil {
			v := &Undo{raw: rawV}
			return true, t.UndoCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Undo'
	// Begin generateResolver for type 'Update'
	if rawV, ok := i.(*vocab.Update); ok {
		if t.UpdateCallback != nil {
			v := &Update{raw: rawV}
			return true, t.UpdateCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Update'
	// Begin generateResolver for type 'View'
	if rawV, ok := i.(*vocab.View); ok {
		if t.ViewCallback != nil {
			v := &View{raw: rawV}
			return true, t.ViewCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'View'
	// Begin generateResolver for type 'Listen'
	if rawV, ok := i.(*vocab.Listen); ok {
		if t.ListenCallback != nil {
			v := &Listen{raw: rawV}
			return true, t.ListenCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Listen'
	// Begin generateResolver for type 'Read'
	if rawV, ok := i.(*vocab.Read); ok {
		if t.ReadCallback != nil {
			v := &Read{raw: rawV}
			return true, t.ReadCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Read'
	// Begin generateResolver for type 'Move'
	if rawV, ok := i.(*vocab.Move); ok {
		if t.MoveCallback != nil {
			v := &Move{raw: rawV}
			return true, t.MoveCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Move'
	// Begin generateResolver for type 'Travel'
	if rawV, ok := i.(*vocab.Travel); ok {
		if t.TravelCallback != nil {
			v := &Travel{raw: rawV}
			return true, t.TravelCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Travel'
	// Begin generateResolver for type 'Announce'
	if rawV, ok := i.(*vocab.Announce); ok {
		if t.AnnounceCallback != nil {
			v := &Announce{raw: rawV}
			return true, t.AnnounceCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Announce'
	// Begin generateResolver for type 'Block'
	if rawV, ok := i.(*vocab.Block); ok {
		if t.BlockCallback != nil {
			v := &Block{raw: rawV}
			return true, t.BlockCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Block'
	// Begin generateResolver for type 'Flag'
	if rawV, ok := i.(*vocab.Flag); ok {
		if t.FlagCallback != nil {
			v := &Flag{raw: rawV}
			return true, t.FlagCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Flag'
	// Begin generateResolver for type 'Dislike'
	if rawV, ok := i.(*vocab.Dislike); ok {
		if t.DislikeCallback != nil {
			v := &Dislike{raw: rawV}
			return true, t.DislikeCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Dislike'
	// Begin generateResolver for type 'Question'
	if rawV, ok := i.(*vocab.Question); ok {
		if t.QuestionCallback != nil {
			v := &Question{raw: rawV}
			return true, t.QuestionCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Question'
	// Begin generateResolver for type 'Application'
	if rawV, ok := i.(*vocab.Application); ok {
		if t.ApplicationCallback != nil {
			v := &Application{raw: rawV}
			return true, t.ApplicationCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Application'
	// Begin generateResolver for type 'Group'
	if rawV, ok := i.(*vocab.Group); ok {
		if t.GroupCallback != nil {
			v := &Group{raw: rawV}
			return true, t.GroupCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Group'
	// Begin generateResolver for type 'Organization'
	if rawV, ok := i.(*vocab.Organization); ok {
		if t.OrganizationCallback != nil {
			v := &Organization{raw: rawV}
			return true, t.OrganizationCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Organization'
	// Begin generateResolver for type 'Person'
	if rawV, ok := i.(*vocab.Person); ok {
		if t.PersonCallback != nil {
			v := &Person{raw: rawV}
			return true, t.PersonCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Person'
	// Begin generateResolver for type 'Service'
	if rawV, ok := i.(*vocab.Service); ok {
		if t.ServiceCallback != nil {
			v := &Service{raw: rawV}
			return true, t.ServiceCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Service'
	// Begin generateResolver for type 'Relationship'
	if rawV, ok := i.(*vocab.Relationship); ok {
		if t.RelationshipCallback != nil {
			v := &Relationship{raw: rawV}
			return true, t.RelationshipCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Relationship'
	// Begin generateResolver for type 'Article'
	if rawV, ok := i.(*vocab.Article); ok {
		if t.ArticleCallback != nil {
			v := &Article{raw: rawV}
			return true, t.ArticleCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Article'
	// Begin generateResolver for type 'Document'
	if rawV, ok := i.(*vocab.Document); ok {
		if t.DocumentCallback != nil {
			v := &Document{raw: rawV}
			return true, t.DocumentCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Document'
	// Begin generateResolver for type 'Audio'
	if rawV, ok := i.(*vocab.Audio); ok {
		if t.AudioCallback != nil {
			v := &Audio{raw: rawV}
			return true, t.AudioCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Audio'
	// Begin generateResolver for type 'Image'
	if rawV, ok := i.(*vocab.Image); ok {
		if t.ImageCallback != nil {
			v := &Image{raw: rawV}
			return true, t.ImageCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Image'
	// Begin generateResolver for type 'Video'
	if rawV, ok := i.(*vocab.Video); ok {
		if t.VideoCallback != nil {
			v := &Video{raw: rawV}
			return true, t.VideoCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Video'
	// Begin generateResolver for type 'Note'
	if rawV, ok := i.(*vocab.Note); ok {
		if t.NoteCallback != nil {
			v := &Note{raw: rawV}
			return true, t.NoteCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Note'
	// Begin generateResolver for type 'Page'
	if rawV, ok := i.(*vocab.Page); ok {
		if t.PageCallback != nil {
			v := &Page{raw: rawV}
			return true, t.PageCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Page'
	// Begin generateResolver for type 'Event'
	if rawV, ok := i.(*vocab.Event); ok {
		if t.EventCallback != nil {
			v := &Event{raw: rawV}
			return true, t.EventCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Event'
	// Begin generateResolver for type 'Place'
	if rawV, ok := i.(*vocab.Place); ok {
		if t.PlaceCallback != nil {
			v := &Place{raw: rawV}
			return true, t.PlaceCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Place'
	// Begin generateResolver for type 'Profile'
	if rawV, ok := i.(*vocab.Profile); ok {
		if t.ProfileCallback != nil {
			v := &Profile{raw: rawV}
			return true, t.ProfileCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Profile'
	// Begin generateResolver for type 'Tombstone'
	if rawV, ok := i.(*vocab.Tombstone); ok {
		if t.TombstoneCallback != nil {
			v := &Tombstone{raw: rawV}
			return true, t.TombstoneCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Tombstone'
	// Begin generateResolver for type 'Mention'
	if rawV, ok := i.(*vocab.Mention); ok {
		if t.MentionCallback != nil {
			v := &Mention{raw: rawV}
			return true, t.MentionCallback(v)
		} else {
			return false, nil
		}
	}
	// End generateResolver for type 'Mention'
	if obj, ok := i.(vocab.ObjectType); ok {
		if t.AnyObjectCallback != nil {
			return true, t.AnyObjectCallback(obj)
		}
	}
	if link, ok := i.(vocab.LinkType); ok {
		if t.AnyLinkCallback != nil {
			return true, t.AnyLinkCallback(link)
		}
	}
	if activity, ok := i.(vocab.ActivityType); ok {
		if t.AnyActivityCallback != nil {
			return true, t.AnyActivityCallback(activity)
		}
	}
	return false, fmt.Errorf("The interface did not match any known types: %T", i)

}

// Determines which concrete type to deserialize this json-unmarshalled item into, returning an error if it cannot determine which type to deserialize into. The appropriate callback, if present, will then be invoked with the concrete deserialized type. If the callback function returns an error, it is passed back through Deserialize.
func (t *Resolver) Deserialize(m map[string]interface{}) (err error) {
	var typeStringVals []string
	typeInterface, ok := m["type"]
	if !ok {
		return fmt.Errorf("Cannot determine type: missing 'type' property")
	}
	if typeStr, ok := typeInterface.(string); ok {
		typeStringVals = append(typeStringVals, typeStr)
	} else if typeSlice, ok := typeInterface.([]interface{}); ok {
		for _, elem := range typeSlice {
			if typeStr, ok := elem.(string); ok {
				typeStringVals = append(typeStringVals, typeStr)
			}
		}
		if len(typeStringVals) == 0 {
			return fmt.Errorf("Cannot determine type: 'type' property is []interface{} with no string elements: %+v", typeInterface)
		}
	} else {
		return fmt.Errorf("Cannot determine type: 'type' property is not string nor []interface{}: %T", typeInterface)
	}
	// Begin generateResolver for type 'Object'
	for _, typeName := range typeStringVals {
		if typeName == "Object" {
			if t.ObjectCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Object{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Object{v}
				if t.ObjectCallback != nil {
					if err := t.ObjectCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Object'
	// Begin generateResolver for type 'Link'
	for _, typeName := range typeStringVals {
		if typeName == "Link" {
			if t.LinkCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Link{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Link{v}
				if t.LinkCallback != nil {
					if err := t.LinkCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Link'
	// Begin generateResolver for type 'Activity'
	for _, typeName := range typeStringVals {
		if typeName == "Activity" {
			if t.ActivityCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Activity{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Activity{v}
				if t.ActivityCallback != nil {
					if err := t.ActivityCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Activity'
	// Begin generateResolver for type 'IntransitiveActivity'
	for _, typeName := range typeStringVals {
		if typeName == "IntransitiveActivity" {
			if t.IntransitiveActivityCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.IntransitiveActivity{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &IntransitiveActivity{v}
				if t.IntransitiveActivityCallback != nil {
					if err := t.IntransitiveActivityCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'IntransitiveActivity'
	// Begin generateResolver for type 'Collection'
	for _, typeName := range typeStringVals {
		if typeName == "Collection" {
			if t.CollectionCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Collection{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Collection{v}
				if t.CollectionCallback != nil {
					if err := t.CollectionCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Collection'
	// Begin generateResolver for type 'OrderedCollection'
	for _, typeName := range typeStringVals {
		if typeName == "OrderedCollection" {
			if t.OrderedCollectionCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.OrderedCollection{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &OrderedCollection{v}
				if t.OrderedCollectionCallback != nil {
					if err := t.OrderedCollectionCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'OrderedCollection'
	// Begin generateResolver for type 'CollectionPage'
	for _, typeName := range typeStringVals {
		if typeName == "CollectionPage" {
			if t.CollectionPageCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.CollectionPage{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &CollectionPage{v}
				if t.CollectionPageCallback != nil {
					if err := t.CollectionPageCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'CollectionPage'
	// Begin generateResolver for type 'OrderedCollectionPage'
	for _, typeName := range typeStringVals {
		if typeName == "OrderedCollectionPage" {
			if t.OrderedCollectionPageCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.OrderedCollectionPage{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &OrderedCollectionPage{v}
				if t.OrderedCollectionPageCallback != nil {
					if err := t.OrderedCollectionPageCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'OrderedCollectionPage'
	// Begin generateResolver for type 'Accept'
	for _, typeName := range typeStringVals {
		if typeName == "Accept" {
			if t.AcceptCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Accept{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Accept{v}
				if t.AcceptCallback != nil {
					if err := t.AcceptCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Accept'
	// Begin generateResolver for type 'TentativeAccept'
	for _, typeName := range typeStringVals {
		if typeName == "TentativeAccept" {
			if t.TentativeAcceptCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.TentativeAccept{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &TentativeAccept{v}
				if t.TentativeAcceptCallback != nil {
					if err := t.TentativeAcceptCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'TentativeAccept'
	// Begin generateResolver for type 'Add'
	for _, typeName := range typeStringVals {
		if typeName == "Add" {
			if t.AddCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Add{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Add{v}
				if t.AddCallback != nil {
					if err := t.AddCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Add'
	// Begin generateResolver for type 'Arrive'
	for _, typeName := range typeStringVals {
		if typeName == "Arrive" {
			if t.ArriveCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Arrive{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Arrive{v}
				if t.ArriveCallback != nil {
					if err := t.ArriveCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Arrive'
	// Begin generateResolver for type 'Create'
	for _, typeName := range typeStringVals {
		if typeName == "Create" {
			if t.CreateCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Create{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Create{v}
				if t.CreateCallback != nil {
					if err := t.CreateCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Create'
	// Begin generateResolver for type 'Delete'
	for _, typeName := range typeStringVals {
		if typeName == "Delete" {
			if t.DeleteCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Delete{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Delete{v}
				if t.DeleteCallback != nil {
					if err := t.DeleteCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Delete'
	// Begin generateResolver for type 'Follow'
	for _, typeName := range typeStringVals {
		if typeName == "Follow" {
			if t.FollowCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Follow{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Follow{v}
				if t.FollowCallback != nil {
					if err := t.FollowCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Follow'
	// Begin generateResolver for type 'Ignore'
	for _, typeName := range typeStringVals {
		if typeName == "Ignore" {
			if t.IgnoreCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Ignore{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Ignore{v}
				if t.IgnoreCallback != nil {
					if err := t.IgnoreCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Ignore'
	// Begin generateResolver for type 'Join'
	for _, typeName := range typeStringVals {
		if typeName == "Join" {
			if t.JoinCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Join{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Join{v}
				if t.JoinCallback != nil {
					if err := t.JoinCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Join'
	// Begin generateResolver for type 'Leave'
	for _, typeName := range typeStringVals {
		if typeName == "Leave" {
			if t.LeaveCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Leave{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Leave{v}
				if t.LeaveCallback != nil {
					if err := t.LeaveCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Leave'
	// Begin generateResolver for type 'Like'
	for _, typeName := range typeStringVals {
		if typeName == "Like" {
			if t.LikeCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Like{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Like{v}
				if t.LikeCallback != nil {
					if err := t.LikeCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Like'
	// Begin generateResolver for type 'Offer'
	for _, typeName := range typeStringVals {
		if typeName == "Offer" {
			if t.OfferCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Offer{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Offer{v}
				if t.OfferCallback != nil {
					if err := t.OfferCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Offer'
	// Begin generateResolver for type 'Invite'
	for _, typeName := range typeStringVals {
		if typeName == "Invite" {
			if t.InviteCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Invite{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Invite{v}
				if t.InviteCallback != nil {
					if err := t.InviteCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Invite'
	// Begin generateResolver for type 'Reject'
	for _, typeName := range typeStringVals {
		if typeName == "Reject" {
			if t.RejectCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Reject{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Reject{v}
				if t.RejectCallback != nil {
					if err := t.RejectCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Reject'
	// Begin generateResolver for type 'TentativeReject'
	for _, typeName := range typeStringVals {
		if typeName == "TentativeReject" {
			if t.TentativeRejectCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.TentativeReject{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &TentativeReject{v}
				if t.TentativeRejectCallback != nil {
					if err := t.TentativeRejectCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'TentativeReject'
	// Begin generateResolver for type 'Remove'
	for _, typeName := range typeStringVals {
		if typeName == "Remove" {
			if t.RemoveCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Remove{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Remove{v}
				if t.RemoveCallback != nil {
					if err := t.RemoveCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Remove'
	// Begin generateResolver for type 'Undo'
	for _, typeName := range typeStringVals {
		if typeName == "Undo" {
			if t.UndoCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Undo{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Undo{v}
				if t.UndoCallback != nil {
					if err := t.UndoCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Undo'
	// Begin generateResolver for type 'Update'
	for _, typeName := range typeStringVals {
		if typeName == "Update" {
			if t.UpdateCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Update{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Update{v}
				if t.UpdateCallback != nil {
					if err := t.UpdateCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Update'
	// Begin generateResolver for type 'View'
	for _, typeName := range typeStringVals {
		if typeName == "View" {
			if t.ViewCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.View{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &View{v}
				if t.ViewCallback != nil {
					if err := t.ViewCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'View'
	// Begin generateResolver for type 'Listen'
	for _, typeName := range typeStringVals {
		if typeName == "Listen" {
			if t.ListenCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Listen{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Listen{v}
				if t.ListenCallback != nil {
					if err := t.ListenCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Listen'
	// Begin generateResolver for type 'Read'
	for _, typeName := range typeStringVals {
		if typeName == "Read" {
			if t.ReadCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Read{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Read{v}
				if t.ReadCallback != nil {
					if err := t.ReadCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Read'
	// Begin generateResolver for type 'Move'
	for _, typeName := range typeStringVals {
		if typeName == "Move" {
			if t.MoveCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Move{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Move{v}
				if t.MoveCallback != nil {
					if err := t.MoveCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Move'
	// Begin generateResolver for type 'Travel'
	for _, typeName := range typeStringVals {
		if typeName == "Travel" {
			if t.TravelCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Travel{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Travel{v}
				if t.TravelCallback != nil {
					if err := t.TravelCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Travel'
	// Begin generateResolver for type 'Announce'
	for _, typeName := range typeStringVals {
		if typeName == "Announce" {
			if t.AnnounceCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Announce{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Announce{v}
				if t.AnnounceCallback != nil {
					if err := t.AnnounceCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Announce'
	// Begin generateResolver for type 'Block'
	for _, typeName := range typeStringVals {
		if typeName == "Block" {
			if t.BlockCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Block{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Block{v}
				if t.BlockCallback != nil {
					if err := t.BlockCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Block'
	// Begin generateResolver for type 'Flag'
	for _, typeName := range typeStringVals {
		if typeName == "Flag" {
			if t.FlagCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Flag{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Flag{v}
				if t.FlagCallback != nil {
					if err := t.FlagCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Flag'
	// Begin generateResolver for type 'Dislike'
	for _, typeName := range typeStringVals {
		if typeName == "Dislike" {
			if t.DislikeCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Dislike{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Dislike{v}
				if t.DislikeCallback != nil {
					if err := t.DislikeCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Dislike'
	// Begin generateResolver for type 'Question'
	for _, typeName := range typeStringVals {
		if typeName == "Question" {
			if t.QuestionCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Question{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Question{v}
				if t.QuestionCallback != nil {
					if err := t.QuestionCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Question'
	// Begin generateResolver for type 'Application'
	for _, typeName := range typeStringVals {
		if typeName == "Application" {
			if t.ApplicationCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Application{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Application{v}
				if t.ApplicationCallback != nil {
					if err := t.ApplicationCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Application'
	// Begin generateResolver for type 'Group'
	for _, typeName := range typeStringVals {
		if typeName == "Group" {
			if t.GroupCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Group{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Group{v}
				if t.GroupCallback != nil {
					if err := t.GroupCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Group'
	// Begin generateResolver for type 'Organization'
	for _, typeName := range typeStringVals {
		if typeName == "Organization" {
			if t.OrganizationCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Organization{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Organization{v}
				if t.OrganizationCallback != nil {
					if err := t.OrganizationCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Organization'
	// Begin generateResolver for type 'Person'
	for _, typeName := range typeStringVals {
		if typeName == "Person" {
			if t.PersonCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Person{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Person{v}
				if t.PersonCallback != nil {
					if err := t.PersonCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Person'
	// Begin generateResolver for type 'Service'
	for _, typeName := range typeStringVals {
		if typeName == "Service" {
			if t.ServiceCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Service{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Service{v}
				if t.ServiceCallback != nil {
					if err := t.ServiceCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Service'
	// Begin generateResolver for type 'Relationship'
	for _, typeName := range typeStringVals {
		if typeName == "Relationship" {
			if t.RelationshipCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Relationship{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Relationship{v}
				if t.RelationshipCallback != nil {
					if err := t.RelationshipCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Relationship'
	// Begin generateResolver for type 'Article'
	for _, typeName := range typeStringVals {
		if typeName == "Article" {
			if t.ArticleCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Article{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Article{v}
				if t.ArticleCallback != nil {
					if err := t.ArticleCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Article'
	// Begin generateResolver for type 'Document'
	for _, typeName := range typeStringVals {
		if typeName == "Document" {
			if t.DocumentCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Document{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Document{v}
				if t.DocumentCallback != nil {
					if err := t.DocumentCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Document'
	// Begin generateResolver for type 'Audio'
	for _, typeName := range typeStringVals {
		if typeName == "Audio" {
			if t.AudioCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Audio{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Audio{v}
				if t.AudioCallback != nil {
					if err := t.AudioCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Audio'
	// Begin generateResolver for type 'Image'
	for _, typeName := range typeStringVals {
		if typeName == "Image" {
			if t.ImageCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Image{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Image{v}
				if t.ImageCallback != nil {
					if err := t.ImageCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Image'
	// Begin generateResolver for type 'Video'
	for _, typeName := range typeStringVals {
		if typeName == "Video" {
			if t.VideoCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Video{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Video{v}
				if t.VideoCallback != nil {
					if err := t.VideoCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Video'
	// Begin generateResolver for type 'Note'
	for _, typeName := range typeStringVals {
		if typeName == "Note" {
			if t.NoteCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Note{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Note{v}
				if t.NoteCallback != nil {
					if err := t.NoteCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Note'
	// Begin generateResolver for type 'Page'
	for _, typeName := range typeStringVals {
		if typeName == "Page" {
			if t.PageCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Page{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Page{v}
				if t.PageCallback != nil {
					if err := t.PageCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Page'
	// Begin generateResolver for type 'Event'
	for _, typeName := range typeStringVals {
		if typeName == "Event" {
			if t.EventCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Event{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Event{v}
				if t.EventCallback != nil {
					if err := t.EventCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Event'
	// Begin generateResolver for type 'Place'
	for _, typeName := range typeStringVals {
		if typeName == "Place" {
			if t.PlaceCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Place{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Place{v}
				if t.PlaceCallback != nil {
					if err := t.PlaceCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Place'
	// Begin generateResolver for type 'Profile'
	for _, typeName := range typeStringVals {
		if typeName == "Profile" {
			if t.ProfileCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Profile{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Profile{v}
				if t.ProfileCallback != nil {
					if err := t.ProfileCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Profile'
	// Begin generateResolver for type 'Tombstone'
	for _, typeName := range typeStringVals {
		if typeName == "Tombstone" {
			if t.TombstoneCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Tombstone{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Tombstone{v}
				if t.TombstoneCallback != nil {
					if err := t.TombstoneCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Tombstone'
	// Begin generateResolver for type 'Mention'
	for _, typeName := range typeStringVals {
		if typeName == "Mention" {
			if t.MentionCallback != nil || t.AnyObjectCallback != nil || t.AnyLinkCallback != nil || t.AnyActivityCallback != nil {
				v := &vocab.Mention{}
				if err := v.Deserialize(m); err != nil {
					return err
				}
				as := &Mention{v}
				if t.MentionCallback != nil {
					if err := t.MentionCallback(as); err != nil {
						return err
					}
				}
				var i interface{} = v
				if obj, ok := i.(vocab.ObjectType); ok {
					if t.AnyObjectCallback != nil {
						if err := t.AnyObjectCallback(obj); err != nil {
							return err
						}
					}
				}
				if link, ok := i.(vocab.LinkType); ok {
					if t.AnyLinkCallback != nil {
						if err := t.AnyLinkCallback(link); err != nil {
							return err
						}
					}
				}
				if activity, ok := i.(vocab.ActivityType); ok {
					if t.AnyActivityCallback != nil {
						if err := t.AnyActivityCallback(activity); err != nil {
							return err
						}
					}
				}
				return nil
			} else {
				return nil
			}
		}
	}
	// End generateResolver for type 'Mention'
	return fmt.Errorf("The 'type' property did not match any known types: %+v", typeStringVals)

}

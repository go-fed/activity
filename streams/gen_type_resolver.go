package streams

import (
	"context"
	"errors"
	vocab "github.com/go-fed/activity/streams/vocab"
)

// TypeResolver resolves ActivityStreams values based on their type name.
type TypeResolver struct {
	callbacks []interface{}
}

// NewTypeResolver creates a new Resolver that examines the type of an
// ActivityStream value to determine what callback function to pass the
// concretely typed value. The callback is guaranteed to receive a value whose
// underlying ActivityStreams type matches the concrete interface name in its
// signature. The callback functions must be of the form:
//
//   func(context.Context, <TypeInterface>) error
//
// where TypeInterface is the code-generated interface for an ActivityStream
// type. An error is returned if a callback function does not match this
// signature.
func NewTypeResolver(callbacks []interface{}) (*TypeResolver, error) {
	for _, cb := range callbacks {
		// Each callback function must satisfy one known function signature, or else we will generate a runtime error instead of silently fail.
		switch cb.(type) {
		case func(context.Context, vocab.AcceptInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.AddInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.AnnounceInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ApplicationInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ArriveInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ArticleInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.AudioInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.BlockInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.CollectionInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.CollectionPageInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.CreateInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.DeleteInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.DislikeInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.DocumentInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.EventInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.FlagInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.FollowInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.GroupInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.IgnoreInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ImageInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.IntransitiveActivityInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.InviteInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.JoinInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.LeaveInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.LikeInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.LinkInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ListenInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.MentionInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.MoveInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.NoteInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ObjectInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.OfferInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.OrderedCollectionInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.OrderedCollectionPageInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.OrganizationInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.PageInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.PersonInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.PlaceInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ProfileInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.QuestionInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ReadInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.RejectInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.RelationshipInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.RemoveInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ServiceInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.TentativeAcceptInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.TentativeRejectInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.TombstoneInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.TravelInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.UndoInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.UpdateInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.VideoInterface) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ViewInterface) error:
			// Do nothing, this callback has a correct signature.
		default:
			return nil, errors.New("a callback function is of the wrong signature and would never be called")
		}
	}
	return &TypeResolver{callbacks: callbacks}, nil
}

// Resolve applies the first callback function whose signature accepts the
// ActivityStreams value's type. This strictly assures that the callback
// function will only be passed ActivityStream objects whose type matches its
// interface. Returns an error if the ActivityStreams type does not match
// callbackers, is not a type handled by the generated code, or the value
// passed in is not go-fed compatible.
func (this TypeResolver) Resolve(ctx context.Context, o ActivityStreamsInterface) error {
	for _, i := range this.callbacks {
		switch o.GetName() {
		case "Accept":
			if fn, ok := i.(func(context.Context, vocab.AcceptInterface) error); ok {
				if v, ok := o.(vocab.AcceptInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Activity":
			if fn, ok := i.(func(context.Context, vocab.ActivityInterface) error); ok {
				if v, ok := o.(vocab.ActivityInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Add":
			if fn, ok := i.(func(context.Context, vocab.AddInterface) error); ok {
				if v, ok := o.(vocab.AddInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Announce":
			if fn, ok := i.(func(context.Context, vocab.AnnounceInterface) error); ok {
				if v, ok := o.(vocab.AnnounceInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Application":
			if fn, ok := i.(func(context.Context, vocab.ApplicationInterface) error); ok {
				if v, ok := o.(vocab.ApplicationInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Arrive":
			if fn, ok := i.(func(context.Context, vocab.ArriveInterface) error); ok {
				if v, ok := o.(vocab.ArriveInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Article":
			if fn, ok := i.(func(context.Context, vocab.ArticleInterface) error); ok {
				if v, ok := o.(vocab.ArticleInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Audio":
			if fn, ok := i.(func(context.Context, vocab.AudioInterface) error); ok {
				if v, ok := o.(vocab.AudioInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Block":
			if fn, ok := i.(func(context.Context, vocab.BlockInterface) error); ok {
				if v, ok := o.(vocab.BlockInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Collection":
			if fn, ok := i.(func(context.Context, vocab.CollectionInterface) error); ok {
				if v, ok := o.(vocab.CollectionInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "CollectionPage":
			if fn, ok := i.(func(context.Context, vocab.CollectionPageInterface) error); ok {
				if v, ok := o.(vocab.CollectionPageInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Create":
			if fn, ok := i.(func(context.Context, vocab.CreateInterface) error); ok {
				if v, ok := o.(vocab.CreateInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Delete":
			if fn, ok := i.(func(context.Context, vocab.DeleteInterface) error); ok {
				if v, ok := o.(vocab.DeleteInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Dislike":
			if fn, ok := i.(func(context.Context, vocab.DislikeInterface) error); ok {
				if v, ok := o.(vocab.DislikeInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Document":
			if fn, ok := i.(func(context.Context, vocab.DocumentInterface) error); ok {
				if v, ok := o.(vocab.DocumentInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Event":
			if fn, ok := i.(func(context.Context, vocab.EventInterface) error); ok {
				if v, ok := o.(vocab.EventInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Flag":
			if fn, ok := i.(func(context.Context, vocab.FlagInterface) error); ok {
				if v, ok := o.(vocab.FlagInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Follow":
			if fn, ok := i.(func(context.Context, vocab.FollowInterface) error); ok {
				if v, ok := o.(vocab.FollowInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Group":
			if fn, ok := i.(func(context.Context, vocab.GroupInterface) error); ok {
				if v, ok := o.(vocab.GroupInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Ignore":
			if fn, ok := i.(func(context.Context, vocab.IgnoreInterface) error); ok {
				if v, ok := o.(vocab.IgnoreInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Image":
			if fn, ok := i.(func(context.Context, vocab.ImageInterface) error); ok {
				if v, ok := o.(vocab.ImageInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "IntransitiveActivity":
			if fn, ok := i.(func(context.Context, vocab.IntransitiveActivityInterface) error); ok {
				if v, ok := o.(vocab.IntransitiveActivityInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Invite":
			if fn, ok := i.(func(context.Context, vocab.InviteInterface) error); ok {
				if v, ok := o.(vocab.InviteInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Join":
			if fn, ok := i.(func(context.Context, vocab.JoinInterface) error); ok {
				if v, ok := o.(vocab.JoinInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Leave":
			if fn, ok := i.(func(context.Context, vocab.LeaveInterface) error); ok {
				if v, ok := o.(vocab.LeaveInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Like":
			if fn, ok := i.(func(context.Context, vocab.LikeInterface) error); ok {
				if v, ok := o.(vocab.LikeInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Link":
			if fn, ok := i.(func(context.Context, vocab.LinkInterface) error); ok {
				if v, ok := o.(vocab.LinkInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Listen":
			if fn, ok := i.(func(context.Context, vocab.ListenInterface) error); ok {
				if v, ok := o.(vocab.ListenInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Mention":
			if fn, ok := i.(func(context.Context, vocab.MentionInterface) error); ok {
				if v, ok := o.(vocab.MentionInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Move":
			if fn, ok := i.(func(context.Context, vocab.MoveInterface) error); ok {
				if v, ok := o.(vocab.MoveInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Note":
			if fn, ok := i.(func(context.Context, vocab.NoteInterface) error); ok {
				if v, ok := o.(vocab.NoteInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Object":
			if fn, ok := i.(func(context.Context, vocab.ObjectInterface) error); ok {
				if v, ok := o.(vocab.ObjectInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Offer":
			if fn, ok := i.(func(context.Context, vocab.OfferInterface) error); ok {
				if v, ok := o.(vocab.OfferInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "OrderedCollection":
			if fn, ok := i.(func(context.Context, vocab.OrderedCollectionInterface) error); ok {
				if v, ok := o.(vocab.OrderedCollectionInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "OrderedCollectionPage":
			if fn, ok := i.(func(context.Context, vocab.OrderedCollectionPageInterface) error); ok {
				if v, ok := o.(vocab.OrderedCollectionPageInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Organization":
			if fn, ok := i.(func(context.Context, vocab.OrganizationInterface) error); ok {
				if v, ok := o.(vocab.OrganizationInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Page":
			if fn, ok := i.(func(context.Context, vocab.PageInterface) error); ok {
				if v, ok := o.(vocab.PageInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Person":
			if fn, ok := i.(func(context.Context, vocab.PersonInterface) error); ok {
				if v, ok := o.(vocab.PersonInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Place":
			if fn, ok := i.(func(context.Context, vocab.PlaceInterface) error); ok {
				if v, ok := o.(vocab.PlaceInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Profile":
			if fn, ok := i.(func(context.Context, vocab.ProfileInterface) error); ok {
				if v, ok := o.(vocab.ProfileInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Question":
			if fn, ok := i.(func(context.Context, vocab.QuestionInterface) error); ok {
				if v, ok := o.(vocab.QuestionInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Read":
			if fn, ok := i.(func(context.Context, vocab.ReadInterface) error); ok {
				if v, ok := o.(vocab.ReadInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Reject":
			if fn, ok := i.(func(context.Context, vocab.RejectInterface) error); ok {
				if v, ok := o.(vocab.RejectInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Relationship":
			if fn, ok := i.(func(context.Context, vocab.RelationshipInterface) error); ok {
				if v, ok := o.(vocab.RelationshipInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Remove":
			if fn, ok := i.(func(context.Context, vocab.RemoveInterface) error); ok {
				if v, ok := o.(vocab.RemoveInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Service":
			if fn, ok := i.(func(context.Context, vocab.ServiceInterface) error); ok {
				if v, ok := o.(vocab.ServiceInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "TentativeAccept":
			if fn, ok := i.(func(context.Context, vocab.TentativeAcceptInterface) error); ok {
				if v, ok := o.(vocab.TentativeAcceptInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "TentativeReject":
			if fn, ok := i.(func(context.Context, vocab.TentativeRejectInterface) error); ok {
				if v, ok := o.(vocab.TentativeRejectInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Tombstone":
			if fn, ok := i.(func(context.Context, vocab.TombstoneInterface) error); ok {
				if v, ok := o.(vocab.TombstoneInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Travel":
			if fn, ok := i.(func(context.Context, vocab.TravelInterface) error); ok {
				if v, ok := o.(vocab.TravelInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Undo":
			if fn, ok := i.(func(context.Context, vocab.UndoInterface) error); ok {
				if v, ok := o.(vocab.UndoInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Update":
			if fn, ok := i.(func(context.Context, vocab.UpdateInterface) error); ok {
				if v, ok := o.(vocab.UpdateInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "Video":
			if fn, ok := i.(func(context.Context, vocab.VideoInterface) error); ok {
				if v, ok := o.(vocab.VideoInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		case "View":
			if fn, ok := i.(func(context.Context, vocab.ViewInterface) error); ok {
				if v, ok := o.(vocab.ViewInterface); ok {
					return fn(ctx, v)
				} else {
					// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
					return errCannotTypeAssertType
				}
			}
		default:
			return ErrUnhandledType
		}
	}
	return ErrNoCallbackMatch
}

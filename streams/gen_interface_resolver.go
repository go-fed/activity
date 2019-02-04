package streams

import (
	"context"
	"errors"
	vocab "github.com/go-fed/activity/streams/vocab"
)

// InterfaceResolver resolves ActivityStreams values based on the interface or
// interfaces they satisfy.
type InterfaceResolver struct {
	callbacks []interface{}
}

// NewInterfaceResolver creates a new Resolver that examines the interface
// assertions on an ActivityStreams value to determine what callback function
// to pass a concretely typed value. The callback may receive a value whose
// underlying ActivityStreams type does not match the concrete interface name
// in its signature. The callback functions must be of the form:
//
//   func(context.Context, <TypeInterface>) error
//
// where TypeInterface is the code-generated interface for an ActivityStream
// type. An error is returned if a callback function does not match this
// signature.
func NewInterfaceResolver(callbacks []interface{}) (*InterfaceResolver, error) {
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
	return &InterfaceResolver{callbacks: callbacks}, nil
}

// Resolve applies the first callback function whose signature accepts an
// interface interpretation of the ActivityStreams value. Note that the Go
// interface rules mean that this can result in multiple unrelated
// ActivityStreams types to be passed into a single callback function and
// potentially causing unintended behaviors. It is best to assume nothing
// about the ActivityStreams' type in the callback function, and instead only
// reason about an ActivityStreams' properties. Returns an error if the
// ActivityStreams interface does not match callbackers or the value passed in
// is not go-fed compatible.
func (this InterfaceResolver) Resolve(ctx context.Context, o ActivityStreamsInterface) error {
	for _, i := range this.callbacks {
		if v, ok := o.(vocab.AcceptInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.AcceptInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.ActivityInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.ActivityInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.AddInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.AddInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.AnnounceInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.AnnounceInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.ApplicationInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.ApplicationInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.ArriveInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.ArriveInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.ArticleInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.ArticleInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.AudioInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.AudioInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.BlockInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.BlockInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.CollectionInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.CollectionInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.CollectionPageInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.CollectionPageInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.CreateInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.CreateInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.DeleteInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.DeleteInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.DislikeInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.DislikeInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.DocumentInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.DocumentInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.EventInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.EventInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.FlagInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.FlagInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.FollowInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.FollowInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.GroupInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.GroupInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.IgnoreInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.IgnoreInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.ImageInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.ImageInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.IntransitiveActivityInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.IntransitiveActivityInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.InviteInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.InviteInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.JoinInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.JoinInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.LeaveInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.LeaveInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.LikeInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.LikeInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.LinkInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.LinkInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.ListenInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.ListenInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.MentionInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.MentionInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.MoveInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.MoveInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.NoteInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.NoteInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.ObjectInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.ObjectInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.OfferInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.OfferInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.OrderedCollectionInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.OrderedCollectionInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.OrderedCollectionPageInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.OrderedCollectionPageInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.OrganizationInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.OrganizationInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.PageInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.PageInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.PersonInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.PersonInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.PlaceInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.PlaceInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.ProfileInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.ProfileInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.QuestionInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.QuestionInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.ReadInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.ReadInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.RejectInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.RejectInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.RelationshipInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.RelationshipInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.RemoveInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.RemoveInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.ServiceInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.ServiceInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.TentativeAcceptInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.TentativeAcceptInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.TentativeRejectInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.TentativeRejectInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.TombstoneInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.TombstoneInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.TravelInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.TravelInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.UndoInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.UndoInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.UpdateInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.UpdateInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.VideoInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.VideoInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}
		if v, ok := o.(vocab.ViewInterface); ok {
			if fn, ok := i.(func(context.Context, vocab.ViewInterface) error); ok {
				return fn(ctx, v)
			}
			// Else: this callback function doesn't support this duck-typed interface.
		}

	}
	return ErrNoCallbackMatch
}

package streams

import (
	"context"
	"errors"
	vocab "github.com/go-fed/activity/streams/vocab"
)

// InterfacePredicatedResolver resolves ActivityStreams values if the value
// satisfies a predicate condition based on the interface or interfaces they
// satisfy.
type InterfacePredicatedResolver struct {
	delegate  Resolver
	predicate interface{}
}

// NewInterfacePredicatedResolver creates a new Resolver that applies a predicate
// to an ActivityStreams value to determine whether to Resolve or not. The
// ActivityStreams value's interface assertions are examined to determine if
// the predicate can apply itself to the value. The predicate will will
// receive a concrete value whose underlying ActivityStreams type may not
// match the concrete interface name, and could be completely unrelated
// ActivityStreams types.The predicate function must be of the form:
//
//   func(context.Context, <TypeInterface>) (bool, error)
//
// where TypeInterface is the code-generated interface for an ActivityStreams
// type. An error is returned if the predicate does not match this signature.
func NewInterfacePredicatedResolver(delegate Resolver, predicate interface{}) (*InterfacePredicatedResolver, error) {
	// The predicate must satisfy one known predicate function signature, or else we will generate a runtime error instead of silently fail.
	switch predicate.(type) {
	case func(context.Context, vocab.InviteInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.BlockInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ArticleInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ArriveInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.OfferInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.CreateInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.AddInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.MentionInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.TentativeAcceptInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.LeaveInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.LinkInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ObjectInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ViewInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.OrderedCollectionPageInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.PlaceInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.AudioInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.JoinInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.IgnoreInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.VideoInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.GroupInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.FlagInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.UndoInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.DislikeInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.CollectionInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.CollectionPageInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ApplicationInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.EventInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ImageInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.PersonInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.OrderedCollectionInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ReadInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.DeleteInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.FollowInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.TentativeRejectInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.DocumentInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.RelationshipInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.PageInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ListenInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.RejectInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.AcceptInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.LikeInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.OrganizationInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.TombstoneInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.UpdateInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ProfileInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.MoveInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.QuestionInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ServiceInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.RemoveInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.IntransitiveActivityInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.TravelInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.AnnounceInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.NoteInterface) (bool, error):
		// Do nothing, this predicate has a correct signature.
	default:
		return nil, errors.New("the predicate function is of the wrong signature and would never be called")
	}
	return &InterfacePredicatedResolver{
		delegate:  delegate,
		predicate: predicate,
	}, nil
}

// Apply uses a predicate to determine whether to resolve the ActivityStreams
// value. The predicate function is applied if its signature accepts an
// interface interpretation of the ActivityStreams value. Note that the Go
// interface rules mean that this can result in multiple unrelated
// ActivityStreams types to be passed into the predicate and potentially cause
// unintended behaviors. It is best to assume nothing about the
// ActivityStreams' type in the predicate function, and instead only reason
// about an ActivityStreams' properties. Returns an error if the
// ActivityStreams interface does not match the predicate or the resolver
// returns an error.
func (this InterfacePredicatedResolver) Apply(ctx context.Context, o ActivityStreamsInterface) (bool, error) {
	var predicatePasses bool
	var err error
	switch fn := this.predicate.(type) {
	case func(context.Context, vocab.InviteInterface) (bool, error):
		if v, ok := o.(vocab.InviteInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.BlockInterface) (bool, error):
		if v, ok := o.(vocab.BlockInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.ActivityInterface) (bool, error):
		if v, ok := o.(vocab.ActivityInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.ArticleInterface) (bool, error):
		if v, ok := o.(vocab.ArticleInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.ArriveInterface) (bool, error):
		if v, ok := o.(vocab.ArriveInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.OfferInterface) (bool, error):
		if v, ok := o.(vocab.OfferInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.CreateInterface) (bool, error):
		if v, ok := o.(vocab.CreateInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.AddInterface) (bool, error):
		if v, ok := o.(vocab.AddInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.MentionInterface) (bool, error):
		if v, ok := o.(vocab.MentionInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.TentativeAcceptInterface) (bool, error):
		if v, ok := o.(vocab.TentativeAcceptInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.LeaveInterface) (bool, error):
		if v, ok := o.(vocab.LeaveInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.LinkInterface) (bool, error):
		if v, ok := o.(vocab.LinkInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.ObjectInterface) (bool, error):
		if v, ok := o.(vocab.ObjectInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.ViewInterface) (bool, error):
		if v, ok := o.(vocab.ViewInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.OrderedCollectionPageInterface) (bool, error):
		if v, ok := o.(vocab.OrderedCollectionPageInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.PlaceInterface) (bool, error):
		if v, ok := o.(vocab.PlaceInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.AudioInterface) (bool, error):
		if v, ok := o.(vocab.AudioInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.JoinInterface) (bool, error):
		if v, ok := o.(vocab.JoinInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.IgnoreInterface) (bool, error):
		if v, ok := o.(vocab.IgnoreInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.VideoInterface) (bool, error):
		if v, ok := o.(vocab.VideoInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.GroupInterface) (bool, error):
		if v, ok := o.(vocab.GroupInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.FlagInterface) (bool, error):
		if v, ok := o.(vocab.FlagInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.UndoInterface) (bool, error):
		if v, ok := o.(vocab.UndoInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.DislikeInterface) (bool, error):
		if v, ok := o.(vocab.DislikeInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.CollectionInterface) (bool, error):
		if v, ok := o.(vocab.CollectionInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.CollectionPageInterface) (bool, error):
		if v, ok := o.(vocab.CollectionPageInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.ApplicationInterface) (bool, error):
		if v, ok := o.(vocab.ApplicationInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.EventInterface) (bool, error):
		if v, ok := o.(vocab.EventInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.ImageInterface) (bool, error):
		if v, ok := o.(vocab.ImageInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.PersonInterface) (bool, error):
		if v, ok := o.(vocab.PersonInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.OrderedCollectionInterface) (bool, error):
		if v, ok := o.(vocab.OrderedCollectionInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.ReadInterface) (bool, error):
		if v, ok := o.(vocab.ReadInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.DeleteInterface) (bool, error):
		if v, ok := o.(vocab.DeleteInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.FollowInterface) (bool, error):
		if v, ok := o.(vocab.FollowInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.TentativeRejectInterface) (bool, error):
		if v, ok := o.(vocab.TentativeRejectInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.DocumentInterface) (bool, error):
		if v, ok := o.(vocab.DocumentInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.RelationshipInterface) (bool, error):
		if v, ok := o.(vocab.RelationshipInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.PageInterface) (bool, error):
		if v, ok := o.(vocab.PageInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.ListenInterface) (bool, error):
		if v, ok := o.(vocab.ListenInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.RejectInterface) (bool, error):
		if v, ok := o.(vocab.RejectInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.AcceptInterface) (bool, error):
		if v, ok := o.(vocab.AcceptInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.LikeInterface) (bool, error):
		if v, ok := o.(vocab.LikeInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.OrganizationInterface) (bool, error):
		if v, ok := o.(vocab.OrganizationInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.TombstoneInterface) (bool, error):
		if v, ok := o.(vocab.TombstoneInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.UpdateInterface) (bool, error):
		if v, ok := o.(vocab.UpdateInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.ProfileInterface) (bool, error):
		if v, ok := o.(vocab.ProfileInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.MoveInterface) (bool, error):
		if v, ok := o.(vocab.MoveInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.QuestionInterface) (bool, error):
		if v, ok := o.(vocab.QuestionInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.ServiceInterface) (bool, error):
		if v, ok := o.(vocab.ServiceInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.RemoveInterface) (bool, error):
		if v, ok := o.(vocab.RemoveInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.IntransitiveActivityInterface) (bool, error):
		if v, ok := o.(vocab.IntransitiveActivityInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.TravelInterface) (bool, error):
		if v, ok := o.(vocab.TravelInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.AnnounceInterface) (bool, error):
		if v, ok := o.(vocab.AnnounceInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	case func(context.Context, vocab.NoteInterface) (bool, error):
		if v, ok := o.(vocab.NoteInterface); ok {
			predicatePasses, err = fn(ctx, v)
		} else {
			return false, ErrPredicateUnmatched
		}
	default:
		// The constructor should guard against this error. If it is encountered, then there is a bug in the code generator.

		return false, errCannotTypeAssertPredicate
	}
	if err != nil {
		return predicatePasses, err
	}
	if predicatePasses {
		return true, this.delegate.Resolve(ctx, o)
	} else {
		return false, nil
	}
}

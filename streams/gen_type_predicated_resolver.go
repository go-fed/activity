package streams

import (
	"context"
	"errors"
	vocab "github.com/go-fed/activity/streams/vocab"
)

// TypePredicatedResolver resolves ActivityStreams values if the value satisfies a
// predicate condition based on its type.
type TypePredicatedResolver struct {
	delegate  Resolver
	predicate interface{}
}

// NewTypePredicatedResolver creates a new Resolver that applies a predicate to an
// ActivityStreams value to determine whether to Resolve or not. The
// ActivityStreams value's type is examined to determine if the predicate can
// apply itself to the value. This guarantees the predicate will receive a
// concrete value whose underlying ActivityStreams type matches the concrete
// interface name. The predicate function must be of the form:
//
//   func(context.Context, <TypeInterface>) (bool, error)
//
// where TypeInterface is the code-generated interface for an ActivityStreams
// type. An error is returned if the predicate does not match this signature.
func NewTypePredicatedResolver(delegate Resolver, predicate interface{}) (*TypePredicatedResolver, error) {
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
	return &TypePredicatedResolver{
		delegate:  delegate,
		predicate: predicate,
	}, nil
}

// Apply uses a predicate to determine whether to resolve the ActivityStreams
// value. The predicate's signature is matched with the ActivityStreams
// value's type. This strictly assures that the predicate will only be passed
// ActivityStream objects whose type matches its interface. Returns an error
// if the ActivityStreams type does not match the predicate, is not a type
// handled by the generated code, or the resolver returns an error. Returns
// true if the predicate returned true.
func (this TypePredicatedResolver) Apply(ctx context.Context, o ActivityStreamsInterface) (bool, error) {
	var predicatePasses bool
	var err error
	switch o.GetName() {
	case "Invite":
		if fn, ok := this.predicate.(func(context.Context, vocab.InviteInterface) (bool, error)); ok {
			if v, ok := o.(vocab.InviteInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Block":
		if fn, ok := this.predicate.(func(context.Context, vocab.BlockInterface) (bool, error)); ok {
			if v, ok := o.(vocab.BlockInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Activity":
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityInterface) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Article":
		if fn, ok := this.predicate.(func(context.Context, vocab.ArticleInterface) (bool, error)); ok {
			if v, ok := o.(vocab.ArticleInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Arrive":
		if fn, ok := this.predicate.(func(context.Context, vocab.ArriveInterface) (bool, error)); ok {
			if v, ok := o.(vocab.ArriveInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Offer":
		if fn, ok := this.predicate.(func(context.Context, vocab.OfferInterface) (bool, error)); ok {
			if v, ok := o.(vocab.OfferInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Create":
		if fn, ok := this.predicate.(func(context.Context, vocab.CreateInterface) (bool, error)); ok {
			if v, ok := o.(vocab.CreateInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Add":
		if fn, ok := this.predicate.(func(context.Context, vocab.AddInterface) (bool, error)); ok {
			if v, ok := o.(vocab.AddInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Mention":
		if fn, ok := this.predicate.(func(context.Context, vocab.MentionInterface) (bool, error)); ok {
			if v, ok := o.(vocab.MentionInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "TentativeAccept":
		if fn, ok := this.predicate.(func(context.Context, vocab.TentativeAcceptInterface) (bool, error)); ok {
			if v, ok := o.(vocab.TentativeAcceptInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Leave":
		if fn, ok := this.predicate.(func(context.Context, vocab.LeaveInterface) (bool, error)); ok {
			if v, ok := o.(vocab.LeaveInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Link":
		if fn, ok := this.predicate.(func(context.Context, vocab.LinkInterface) (bool, error)); ok {
			if v, ok := o.(vocab.LinkInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Object":
		if fn, ok := this.predicate.(func(context.Context, vocab.ObjectInterface) (bool, error)); ok {
			if v, ok := o.(vocab.ObjectInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "View":
		if fn, ok := this.predicate.(func(context.Context, vocab.ViewInterface) (bool, error)); ok {
			if v, ok := o.(vocab.ViewInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "OrderedCollectionPage":
		if fn, ok := this.predicate.(func(context.Context, vocab.OrderedCollectionPageInterface) (bool, error)); ok {
			if v, ok := o.(vocab.OrderedCollectionPageInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Place":
		if fn, ok := this.predicate.(func(context.Context, vocab.PlaceInterface) (bool, error)); ok {
			if v, ok := o.(vocab.PlaceInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Audio":
		if fn, ok := this.predicate.(func(context.Context, vocab.AudioInterface) (bool, error)); ok {
			if v, ok := o.(vocab.AudioInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Join":
		if fn, ok := this.predicate.(func(context.Context, vocab.JoinInterface) (bool, error)); ok {
			if v, ok := o.(vocab.JoinInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Ignore":
		if fn, ok := this.predicate.(func(context.Context, vocab.IgnoreInterface) (bool, error)); ok {
			if v, ok := o.(vocab.IgnoreInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Video":
		if fn, ok := this.predicate.(func(context.Context, vocab.VideoInterface) (bool, error)); ok {
			if v, ok := o.(vocab.VideoInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Group":
		if fn, ok := this.predicate.(func(context.Context, vocab.GroupInterface) (bool, error)); ok {
			if v, ok := o.(vocab.GroupInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Flag":
		if fn, ok := this.predicate.(func(context.Context, vocab.FlagInterface) (bool, error)); ok {
			if v, ok := o.(vocab.FlagInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Undo":
		if fn, ok := this.predicate.(func(context.Context, vocab.UndoInterface) (bool, error)); ok {
			if v, ok := o.(vocab.UndoInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Dislike":
		if fn, ok := this.predicate.(func(context.Context, vocab.DislikeInterface) (bool, error)); ok {
			if v, ok := o.(vocab.DislikeInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Collection":
		if fn, ok := this.predicate.(func(context.Context, vocab.CollectionInterface) (bool, error)); ok {
			if v, ok := o.(vocab.CollectionInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "CollectionPage":
		if fn, ok := this.predicate.(func(context.Context, vocab.CollectionPageInterface) (bool, error)); ok {
			if v, ok := o.(vocab.CollectionPageInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Application":
		if fn, ok := this.predicate.(func(context.Context, vocab.ApplicationInterface) (bool, error)); ok {
			if v, ok := o.(vocab.ApplicationInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Event":
		if fn, ok := this.predicate.(func(context.Context, vocab.EventInterface) (bool, error)); ok {
			if v, ok := o.(vocab.EventInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Image":
		if fn, ok := this.predicate.(func(context.Context, vocab.ImageInterface) (bool, error)); ok {
			if v, ok := o.(vocab.ImageInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Person":
		if fn, ok := this.predicate.(func(context.Context, vocab.PersonInterface) (bool, error)); ok {
			if v, ok := o.(vocab.PersonInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "OrderedCollection":
		if fn, ok := this.predicate.(func(context.Context, vocab.OrderedCollectionInterface) (bool, error)); ok {
			if v, ok := o.(vocab.OrderedCollectionInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Read":
		if fn, ok := this.predicate.(func(context.Context, vocab.ReadInterface) (bool, error)); ok {
			if v, ok := o.(vocab.ReadInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Delete":
		if fn, ok := this.predicate.(func(context.Context, vocab.DeleteInterface) (bool, error)); ok {
			if v, ok := o.(vocab.DeleteInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Follow":
		if fn, ok := this.predicate.(func(context.Context, vocab.FollowInterface) (bool, error)); ok {
			if v, ok := o.(vocab.FollowInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "TentativeReject":
		if fn, ok := this.predicate.(func(context.Context, vocab.TentativeRejectInterface) (bool, error)); ok {
			if v, ok := o.(vocab.TentativeRejectInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Document":
		if fn, ok := this.predicate.(func(context.Context, vocab.DocumentInterface) (bool, error)); ok {
			if v, ok := o.(vocab.DocumentInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Relationship":
		if fn, ok := this.predicate.(func(context.Context, vocab.RelationshipInterface) (bool, error)); ok {
			if v, ok := o.(vocab.RelationshipInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Page":
		if fn, ok := this.predicate.(func(context.Context, vocab.PageInterface) (bool, error)); ok {
			if v, ok := o.(vocab.PageInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Listen":
		if fn, ok := this.predicate.(func(context.Context, vocab.ListenInterface) (bool, error)); ok {
			if v, ok := o.(vocab.ListenInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Reject":
		if fn, ok := this.predicate.(func(context.Context, vocab.RejectInterface) (bool, error)); ok {
			if v, ok := o.(vocab.RejectInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Accept":
		if fn, ok := this.predicate.(func(context.Context, vocab.AcceptInterface) (bool, error)); ok {
			if v, ok := o.(vocab.AcceptInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Like":
		if fn, ok := this.predicate.(func(context.Context, vocab.LikeInterface) (bool, error)); ok {
			if v, ok := o.(vocab.LikeInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Organization":
		if fn, ok := this.predicate.(func(context.Context, vocab.OrganizationInterface) (bool, error)); ok {
			if v, ok := o.(vocab.OrganizationInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Tombstone":
		if fn, ok := this.predicate.(func(context.Context, vocab.TombstoneInterface) (bool, error)); ok {
			if v, ok := o.(vocab.TombstoneInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Update":
		if fn, ok := this.predicate.(func(context.Context, vocab.UpdateInterface) (bool, error)); ok {
			if v, ok := o.(vocab.UpdateInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Profile":
		if fn, ok := this.predicate.(func(context.Context, vocab.ProfileInterface) (bool, error)); ok {
			if v, ok := o.(vocab.ProfileInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Move":
		if fn, ok := this.predicate.(func(context.Context, vocab.MoveInterface) (bool, error)); ok {
			if v, ok := o.(vocab.MoveInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Question":
		if fn, ok := this.predicate.(func(context.Context, vocab.QuestionInterface) (bool, error)); ok {
			if v, ok := o.(vocab.QuestionInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Service":
		if fn, ok := this.predicate.(func(context.Context, vocab.ServiceInterface) (bool, error)); ok {
			if v, ok := o.(vocab.ServiceInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Remove":
		if fn, ok := this.predicate.(func(context.Context, vocab.RemoveInterface) (bool, error)); ok {
			if v, ok := o.(vocab.RemoveInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "IntransitiveActivity":
		if fn, ok := this.predicate.(func(context.Context, vocab.IntransitiveActivityInterface) (bool, error)); ok {
			if v, ok := o.(vocab.IntransitiveActivityInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Travel":
		if fn, ok := this.predicate.(func(context.Context, vocab.TravelInterface) (bool, error)); ok {
			if v, ok := o.(vocab.TravelInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Announce":
		if fn, ok := this.predicate.(func(context.Context, vocab.AnnounceInterface) (bool, error)); ok {
			if v, ok := o.(vocab.AnnounceInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	case "Note":
		if fn, ok := this.predicate.(func(context.Context, vocab.NoteInterface) (bool, error)); ok {
			if v, ok := o.(vocab.NoteInterface); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	default:
		return false, ErrUnhandledType
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

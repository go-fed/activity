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
	case func(context.Context, vocab.ActivityStreamsAccept) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsActivity) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsAdd) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsAnnounce) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsApplication) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsArrive) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsArticle) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsAudio) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsBlock) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsCollection) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsCollectionPage) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsCreate) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsDelete) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsDislike) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsDocument) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsEvent) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsFlag) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsFollow) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsGroup) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsIgnore) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsImage) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsIntransitiveActivity) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsInvite) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsJoin) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsLeave) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsLike) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsLink) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsListen) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsMention) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsMove) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsNote) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsObject) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsOffer) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsOrderedCollection) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsOrderedCollectionPage) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsOrganization) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsPage) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsPerson) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsPlace) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsProfile) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsQuestion) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsRead) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsReject) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsRelationship) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsRemove) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsService) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsTentativeAccept) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsTentativeReject) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsTombstone) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsTravel) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsUndo) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsUpdate) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsVideo) (bool, error):
		// Do nothing, this predicate has a correct signature.
	case func(context.Context, vocab.ActivityStreamsView) (bool, error):
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
	if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Accept" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsAccept) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsAccept); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Activity" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsActivity) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsActivity); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Add" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsAdd) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsAdd); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Announce" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsAnnounce) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsAnnounce); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Application" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsApplication) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsApplication); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Arrive" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsArrive) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsArrive); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Article" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsArticle) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsArticle); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Audio" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsAudio) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsAudio); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Block" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsBlock) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsBlock); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Collection" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsCollection) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsCollection); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "CollectionPage" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsCollectionPage) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsCollectionPage); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Create" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsCreate) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsCreate); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Delete" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsDelete) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsDelete); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Dislike" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsDislike) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsDislike); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Document" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsDocument) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsDocument); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Event" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsEvent) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsEvent); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Flag" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsFlag) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsFlag); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Follow" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsFollow) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsFollow); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Group" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsGroup) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsGroup); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Ignore" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsIgnore) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsIgnore); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Image" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsImage) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsImage); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "IntransitiveActivity" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsIntransitiveActivity) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsIntransitiveActivity); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Invite" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsInvite) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsInvite); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Join" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsJoin) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsJoin); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Leave" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsLeave) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsLeave); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Like" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsLike) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsLike); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Link" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsLink) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsLink); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Listen" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsListen) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsListen); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Mention" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsMention) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsMention); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Move" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsMove) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsMove); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Note" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsNote) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsNote); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Object" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsObject) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsObject); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Offer" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsOffer) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsOffer); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "OrderedCollection" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsOrderedCollection) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsOrderedCollection); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "OrderedCollectionPage" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsOrderedCollectionPage) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsOrderedCollectionPage); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Organization" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsOrganization) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsOrganization); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Page" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsPage) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsPage); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Person" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsPerson) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsPerson); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Place" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsPlace) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsPlace); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Profile" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsProfile) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsProfile); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Question" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsQuestion) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsQuestion); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Read" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsRead) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsRead); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Reject" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsReject) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsReject); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Relationship" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsRelationship) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsRelationship); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Remove" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsRemove) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsRemove); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Service" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsService) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsService); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "TentativeAccept" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsTentativeAccept) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsTentativeAccept); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "TentativeReject" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsTentativeReject) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsTentativeReject); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Tombstone" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsTombstone) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsTombstone); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Travel" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsTravel) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsTravel); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Undo" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsUndo) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsUndo); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Update" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsUpdate) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsUpdate); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "Video" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsVideo) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsVideo); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else if o.VocabularyURI() == "https://www.w3.org/TR/activitystreams-vocabulary" && o.GetName() == "View" {
		if fn, ok := this.predicate.(func(context.Context, vocab.ActivityStreamsView) (bool, error)); ok {
			if v, ok := o.(vocab.ActivityStreamsView); ok {
				predicatePasses, err = fn(ctx, v)
			} else {
				// This occurs when the value is either not a go-fed type and is improperly satisfying various interfaces, or there is a bug in the go-fed generated code.
				return false, errCannotTypeAssertType
			}
		} else {
			return false, ErrPredicateUnmatched
		}
	} else {
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

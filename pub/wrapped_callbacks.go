package pub

import (
	"context"
	"github.com/go-fed/activity/streams/vocab"
)

// WrappedCallbacks lists the callback functions that already have some side
// effect behavior provided by the pub library.
//
// These functions may be wrapped for either the Federating Protocol or the
// Social API. However, the side effects in these callbacks should not be the
// same for both of these use cases.
//
// These are not used when using a DelegateActor directly. The wrapping
// behaviors defined below would need to be handled by another implementation
// of that interface.
type WrappedCallbacks struct {
	// Create handles additional side effects for the Create ActivityStreams
	// type.
	//
	// The wrapping callback for the Federating Protocol will ensure the
	// 'object' property exists, create an entry in the database, and add it
	// to the recipient(s) inbox if not yet already in the inbox.
	//
	// The wrapping callback for the Social API copies the actor(s) to the
	// 'attributedTo' property, copying recipients between the Create
	// activity and all objects, save the entry in the database, and adds it
	// to the outbox.
	Create func(context.Context, vocab.ActivityStreamsCreate) error
	// Update handles additional side effects for the Update ActivityStreams
	// type.
	//
	// TODO: Describe
	Update func(context.Context, vocab.ActivityStreamsUpdate) error
	// Delete handles additional side effects for the Delete ActivityStreams
	// type.
	//
	// TODO: Describe
	Delete func(context.Context, vocab.ActivityStreamsDelete) error
	// Follow handles additional side effects for the Follow ActivityStreams
	// type.
	//
	// TODO: Describe
	Follow func(context.Context, vocab.ActivityStreamsFollow) error
	// Accept handles additional side effects for the Accept ActivityStreams
	// type.
	//
	// TODO: Describe
	Accept func(context.Context, vocab.ActivityStreamsAccept) error
	// Reject handles additional side effects for the Reject ActivityStreams
	// type.
	//
	// TODO: Describe
	Reject func(context.Context, vocab.ActivityStreamsReject) error
	// Add handles additional side effects for the Add ActivityStreams
	// type.
	//
	// TODO: Describe
	Add func(context.Context, vocab.ActivityStreamsAdd) error
	// Remove handles additional side effects for the Remove ActivityStreams
	// type.
	//
	// TODO: Describe
	Remove func(context.Context, vocab.ActivityStreamsRemove) error
	// Like handles additional side effects for the Like ActivityStreams
	// type.
	//
	// TODO: Describe
	Like func(context.Context, vocab.ActivityStreamsLike) error
	// Undo handles additional side effects for the Undo ActivityStreams
	// type.
	//
	// TODO: Describe
	Undo func(context.Context, vocab.ActivityStreamsUndo) error
	// Block handles additional side effects for the Block ActivityStreams
	// type.
	//
	// TODO: Describe
	Block func(context.Context, vocab.ActivityStreamsBlock) error
}

// callbacks returns the WrappedCallbacks members into a single interface slice
// for use in streams.Resolver callbacks.
func (w WrappedCallbacks) callbacks() []interface{} {
	return []interface{}{
		w.Create,
		w.Update,
		w.Delete,
		w.Follow,
		w.Accept,
		w.Reject,
		w.Add,
		w.Remove,
		w.Like,
		w.Undo,
		w.Block,
	}
}

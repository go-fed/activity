package pub

import (
	"context"
	"github.com/go-fed/activity/streams/vocab"
	"net/url"
)

type Database interface {
	// Lock takes a lock for the object at the specified id. If an error
	// is returned, the lock must not have been taken.
	//
	// The lock must be able to succeed for an id that does not exist in
	// the database. This means acquiring the lock does not guarantee the
	// entry exists in the database.
	//
	// Locks are encouraged to be lightweight and in the Go layer, as some
	// processes require tight loops acquiring and releasing locks.
	//
	// Used to ensure race conditions in multiple requests do not occur.
	Lock(c context.Context, id *url.URL) error
	// Unlock makes the lock for the object at the specified id available.
	// If an error is returned, the lock must have still been freed.
	//
	// Used to ensure race conditions in multiple requests do not occur.
	Unlock(c context.Context, id *url.URL) error
	// InboxContains returns true if the OrderedCollection at 'inbox'
	// contains the specified 'id'.
	//
	// The library makes this call only after acquiring a lock first.
	InboxContains(c context.Context, inbox, id *url.URL) (contains bool, err error)
	// GetInbox returns the first ordered collection page of the outbox at
	// the specified IRI, for prepending new items.
	//
	// The library makes this call only after acquiring a lock first.
	GetInbox(c context.Context, inboxIRI *url.URL) (inbox vocab.ActivityStreamsOrderedCollectionPage, err error)
	// SetInbox saves the inbox value given from GetInbox, with new items
	// prepended. Note that the new items must not be added as independent
	// database entries. Separate calls to Create will do that.
	//
	// The library makes this call only after acquiring a lock first.
	SetInbox(c context.Context, inbox vocab.ActivityStreamsOrderedCollectionPage) error
	// Owns returns true if the database has an entry for the IRI and it
	// exists in the database.
	//
	// Owns is called even without acquiring a lock.
	Owns(c context.Context, id *url.URL) (owns bool, err error)
	// Exists returns true if the database has an entry for the specified
	// id. It may not be owned by this application instance.
	//
	// The library makes this call only after acquiring a lock first.
	Exists(c context.Context, id *url.URL) (exists bool, err error)
	// Get returns the database entry for the specified id.
	//
	// The library makes this call only after acquiring a lock first.
	Get(c context.Context, id *url.URL) (value vocab.Type, err error)
	// Create adds a new entry to the database which must be able to be
	// keyed by its id.
	//
	// Note that Activity values received from federated peers may also be
	// created in the database this way if the Federating Protocol is
	// enabled.
	//
	// The library makes this call only after acquiring a lock first.
	Create(c context.Context, asType vocab.Type) error
	// GetOutbox returns the first ordered collection page of the outbox
	// at the specified IRI, for prepending new items.
	//
	// The library makes this call only after acquiring a lock first.
	GetOutbox(c context.Context, inboxIRI *url.URL) (inbox vocab.ActivityStreamsOrderedCollectionPage, err error)
	// SetOutbox saves the outbox value given from GetOutbox, with new items
	// prepended. Note that the new items must not be added as independent
	// database entries. Separate calls to Create will do that.
	//
	// The library makes this call only after acquiring a lock first.
	SetOutbox(c context.Context, inbox vocab.ActivityStreamsOrderedCollectionPage) error
	// NewId creates a new IRI id for the provided activity or object. The
	// implementation does not need to set the 'id' property and simply
	// needs to determine the value.
	//
	// The go-fed library will handle setting the 'id' property on the
	// activity or object provided with the value returned.
	NewId(c context.Context, t vocab.Type) (id *url.URL, err error)
}

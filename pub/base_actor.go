package pub

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// TODO: Rename GetType and GetName

// baseActor must satisfy the Actor interface.
var _ Actor = &baseActor{}

// baseActor is an application-independent ActivityPub implementation. It does
// not implement the entire protocol, and relies on a delegate to do so. It
// only implements the part of the protocol that is side-effect-free, allowing
// an existing application to write a DelegateActor that glues their application
// into the ActivityPub world.
//
// It is preferred to use a DelegateActor provided by this library, so that the
// application does not need to worry about the ActivityPub implementation.
type baseActor struct {
	// delegate contains application-specific delegation logic.
	delegate DelegateActor
	// enableSocialProtocol enables or disables the Social API, the client to
	// server part of ActivityPub. Useful if permitting remote clients to
	// act on behalf of the users of the client application.
	enableSocialProtocol bool
	// enableFederatedProtocol enables or disables the Federated Protocol, or the
	// server to server part of ActivityPub. Useful to permit integrating
	// with the rest of the federative web.
	enableFederatedProtocol bool
	// clock simply tracks the current time.
	clock Clock
}

// NewSocialActor builds a new Actor concept that handles only the Social
// Protocol part of ActivityPub.
//
// This Actor can be created once in an application and reused to handle
// multiple requests concurrently and for different endpoints.
//
// It leverages as much of go-fed as possible to ensure the implementation is
// compliant with the ActivityPub specification, while providing enough freedom
// to be productive without shooting one's self in the foot.
//
// Do not try to use NewSocialActor and NewFederatingActor together to cover
// both the Social and Federating parts of the protocol. Instead, use NewActor.
func NewSocialActor(c CommonBehavior,
	c2s SocialProtocol,
	db Database,
	clock Clock) Actor {
	return &baseActor{
		delegate: &sideEffectActor{
			common: c,
			c2s:    c2s,
			db:     db,
			clock:  clock,
		},
		enableSocialProtocol: true,
		clock:                clock,
	}
}

// NewFederatingActor builds a new Actor concept that handles only the Federating
// Protocol part of ActivityPub.
//
// This Actor can be created once in an application and reused to handle
// multiple requests concurrently and for different endpoints.
//
// It leverages as much of go-fed as possible to ensure the implementation is
// compliant with the ActivityPub specification, while providing enough freedom
// to be productive without shooting one's self in the foot.
//
// Do not try to use NewSocialActor and NewFederatingActor together to cover
// both the Social and Federating parts of the protocol. Instead, use NewActor.
func NewFederatingActor(c CommonBehavior,
	s2s FederatingProtocol,
	db Database,
	clock Clock) Actor {
	return &baseActor{
		delegate: &sideEffectActor{
			common: c,
			s2s:    s2s,
			db:     db,
			clock:  clock,
		},
		enableFederatedProtocol: true,
		clock:                   clock,
	}
}

// NewActor builds a new Actor concept that handles both the Social and
// Federating Protocol parts of ActivityPub.
//
// This Actor can be created once in an application and reused to handle
// multiple requests concurrently and for different endpoints.
//
// It leverages as much of go-fed as possible to ensure the implementation is
// compliant with the ActivityPub specification, while providing enough freedom
// to be productive without shooting one's self in the foot.
func NewActor(c CommonBehavior,
	c2s SocialProtocol,
	s2s FederatingProtocol,
	db Database,
	clock Clock) Actor {
	return &baseActor{
		delegate: &sideEffectActor{
			common: c,
			c2s:    c2s,
			s2s:    s2s,
			db:     db,
			clock:  clock,
		},
		enableSocialProtocol:    true,
		enableFederatedProtocol: true,
		clock:                   clock,
	}
}

// NewCustomActor allows clients to create a custom ActivityPub implementation
// for the Social Protocol, Federating Protocol, or both.
//
// It still uses the library as a high-level scaffold, which has the benefit of
// allowing applications to grow into a custom solution without having to
// refactor the code that passes HTTP requests into the Actor.
//
// It is possible to create a DelegateActor that is not ActivityPub compliant.
// Use with care.
func NewCustomActor(delegate DelegateActor,
	enableSocialProtocol, enableFederatedProtocol bool,
	clock Clock) Actor {
	return &baseActor{
		delegate:                delegate,
		enableSocialProtocol:    enableSocialProtocol,
		enableFederatedProtocol: enableFederatedProtocol,
		clock:                   clock,
	}
}

// PostInbox implements the generic algorithm for handling a POST request to an
// actor's inbox independent on an application. It relies on a delegate to
// implement application specific functionality.
func (b *baseActor) PostInbox(c context.Context, w http.ResponseWriter, r *http.Request) (bool, error) {
	// Do nothing if it is not an ActivityPub POST request.
	if !isActivityPubPost(r) {
		return false, nil
	}
	// If the Federated Protocol is not enabled, then this endpoint is not
	// enabled.
	if !b.enableFederatedProtocol {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return true, nil
	}
	// Check the peer request is authentic.
	shouldReturn, err := b.delegate.AuthenticatePostInbox(c, w, r)
	if err != nil {
		return true, err
	} else if shouldReturn {
		return true, nil
	}
	// Begin processing the request, but have not yet applied
	// authorization (ex: blocks). Obtain the activity reject unknown
	// activities.
	raw, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return true, err
	}
	var m map[string]interface{}
	if err = json.Unmarshal(raw, &m); err != nil {
		return true, err
	}
	// TODO: No longer reject unknown activities.
	asValue, err := toType(c, m)
	if err != nil {
		return true, err
	}
	activity, ok := asValue.(Activity)
	if !ok {
		return true, fmt.Errorf("activity streams value is not an Activity: %T", asValue)
	}
	if activity.GetActivityStreamsId() == nil {
		w.WriteHeader(http.StatusBadRequest)
		return true, nil
	}
	// Check authorization of the activity.
	shouldReturn, err = b.delegate.AuthorizePostInbox(c, w, activity)
	if err != nil {
		return true, err
	} else if shouldReturn {
		return true, nil
	}
	// Post the activity to the actor's inbox and trigger side effects for
	// that particular Activity type. It is up to the delegate to resolve
	// the given map.
	err = b.delegate.PostInbox(c, r.URL, activity)
	if err != nil {
		// Special case: We know it is a bad request if the object or
		// target properties needed to be populated, but weren't.
		//
		// Send the rejection to the peer.
		if err == ErrObjectRequired || err == ErrTargetRequired {
			w.WriteHeader(http.StatusBadRequest)
			return true, nil
		}
		return true, err
	}
	// Our side effects are complete, now delegate determining whether to
	// do inbox forwarding, as well as the action to do it.
	if err := b.delegate.InboxForwarding(c, r.URL, activity); err != nil {
		return true, err
	}
	// Request has been processed. Begin responding to the request.
	//
	// Simply respond with an OK status to the peer.
	w.WriteHeader(http.StatusOK)
	return true, nil
}

// GetInbox implements the generic algorithm for handling a GET request to an
// actor's inbox independent on an application. It relies on a delegate to
// implement application specific functionality.
func (b *baseActor) GetInbox(c context.Context, w http.ResponseWriter, r *http.Request) (bool, error) {
	// Do nothing if it is not an ActivityPub GET request.
	if !isActivityPubGet(r) {
		return false, nil
	}
	// Delegate authenticating and authorizing the request.
	shouldReturn, err := b.delegate.AuthenticateGetInbox(c, w, r)
	if err != nil {
		return true, err
	} else if shouldReturn {
		return true, nil
	}
	// Everything is good to begin processing the request.
	oc, err := b.delegate.GetInbox(c, r)
	if err != nil {
		return true, err
	}
	// Deduplicate the 'orderedItems' property by ID.
	err = dedupeOrderedItems(oc)
	if err != nil {
		return true, err
	}
	// Request has been processed. Begin responding to the request.
	//
	// Serialize the OrderedCollection.
	m, err := serialize(oc)
	if err != nil {
		return true, err
	}
	raw, err := json.Marshal(m)
	if err != nil {
		return true, err
	}
	// Write the response.
	addResponseHeaders(w.Header(), b.clock, raw)
	w.WriteHeader(http.StatusOK)
	n, err := w.Write(raw)
	if err != nil {
		return true, err
	} else if n != len(raw) {
		return true, fmt.Errorf("ResponseWriter.Write wrote %d of %d bytes", n, len(raw))
	}
	return true, nil
}

// PostOutbox implements the generic algorithm for handling a POST request to an
// actor's outbox independent on an application. It relies on a delegate to
// implement application specific functionality.
func (b *baseActor) PostOutbox(c context.Context, w http.ResponseWriter, r *http.Request) (bool, error) {
	// Do nothing if it is not an ActivityPub POST request.
	if !isActivityPubPost(r) {
		return false, nil
	}
	// If the Social API is not enabled, then this endpoint is not enabled.
	if !b.enableSocialProtocol {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return true, nil
	}
	// Delegate authenticating and authorizing the request.
	shouldReturn, err := b.delegate.AuthenticatePostOutbox(c, w, r)
	if err != nil {
		return true, err
	} else if shouldReturn {
		return true, nil
	}
	// Everything is good to begin processing the request.
	raw, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return true, err
	}
	var m map[string]interface{}
	if err = json.Unmarshal(raw, &m); err != nil {
		return true, err
	}
	// Note that converting to a Type will NOT successfully convert types
	// not known to go-fed. This prevents accidentally wrapping an Activity
	// type unknown to go-fed in a Create below. Instead,
	// streams.ErrUnhandledType will be returned here.
	//
	// TODO: No longer reject unknown activities.
	asValue, err := toType(c, m)
	if err != nil {
		return true, err
	}
	// If the value is not an Activity or type extending from Activity, then
	// we need to wrap it in a Create Activity.
	if !IsAnActivityType(asValue) {
		asValue, err = b.delegate.WrapInCreate(c, asValue, r.URL)
		if err != nil {
			return true, err
		}
	}
	// At this point, this should be a safe conversion. If this error is
	// triggered, then there is either a bug in the delegation of
	// WrapInCreate, behavior is not lining up in the generated ExtendedBy
	// code, or something else is incorrect with the type system.
	activity, ok := asValue.(Activity)
	if !ok {
		return true, fmt.Errorf("activity streams value is not an Activity: %T", asValue)
	}
	// Delegate generating new IDs for the activity and all new objects.
	if err = b.delegate.AddNewIds(c, activity); err != nil {
		return true, err
	}
	// Post the activity to the actor's outbox and trigger side effects for
	// that particular Activity type.
	deliverable, err := b.delegate.PostOutbox(c, activity, r.URL, m)
	if err != nil {
		// Special case: We know it is a bad request if the object or
		// target properties needed to be populated, but weren't.
		//
		// Send the rejection to the peer.
		if err == ErrObjectRequired || err == ErrTargetRequired {
			w.WriteHeader(http.StatusBadRequest)
			return true, nil
		}
		return true, err
	}
	// Request has been processed and all side effects internal to this
	// application server have finished. Begin side effects affecting other
	// servers and/or the client who sent this request.
	//
	// If we are federating and the type is a deliverable one, then deliver
	// the activity to federating peers.
	if b.enableFederatedProtocol && deliverable {
		if err := b.delegate.Deliver(c, r.URL, activity); err != nil {
			return true, err
		}
	}
	// Respond to the request with the new Activity's IRI location.
	w.Header().Set("Location", activity.GetActivityStreamsId().Get().String())
	w.WriteHeader(http.StatusCreated)
	return true, nil
}

// GetOutbox implements the generic algorithm for handling a Get request to an
// actor's outbox independent on an application. It relies on a delegate to
// implement application specific functionality.
func (b *baseActor) GetOutbox(c context.Context, w http.ResponseWriter, r *http.Request) (bool, error) {
	// Do nothing if it is not an ActivityPub GET request.
	if !isActivityPubGet(r) {
		return false, nil
	}
	// Delegate authenticating and authorizing the request.
	shouldReturn, err := b.delegate.AuthenticateGetOutbox(c, w, r)
	if err != nil {
		return true, err
	} else if shouldReturn {
		return true, nil
	}
	// Everything is good to begin processing the request.
	oc, err := b.delegate.GetOutbox(c, r)
	if err != nil {
		return true, err
	}
	// Request has been processed. Begin responding to the request.
	//
	// Serialize the OrderedCollection.
	m, err := serialize(oc)
	if err != nil {
		return true, err
	}
	raw, err := json.Marshal(m)
	if err != nil {
		return true, err
	}
	// Write the response.
	addResponseHeaders(w.Header(), b.clock, raw)
	w.WriteHeader(http.StatusOK)
	n, err := w.Write(raw)
	if err != nil {
		return true, err
	} else if n != len(raw) {
		return true, fmt.Errorf("ResponseWriter.Write wrote %d of %d bytes", n, len(raw))
	}
	return true, nil
}

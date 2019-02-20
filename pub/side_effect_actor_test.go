package pub

import (
	"context"
	"github.com/go-fed/activity/streams/vocab"
	"github.com/golang/mock/gomock"
	"net/http/httptest"
	"net/url"
	"testing"
)

// TestPassThroughMethods tests the methods that pass-through to other
// dependency-injected types.
func TestPassThroughMethods(t *testing.T) {
	ctx := context.Background()
	resp := httptest.NewRecorder()
	setupFn := func(ctl *gomock.Controller) (c *MockCommonBehavior, fp *MockFederatingProtocol, sp *MockSocialProtocol, db *MockDatabase, cl *MockClock, a DelegateActor) {
		setupData()
		c = NewMockCommonBehavior(ctl)
		fp = NewMockFederatingProtocol(ctl)
		sp = NewMockSocialProtocol(ctl)
		db = NewMockDatabase(ctl)
		cl = NewMockClock(ctl)
		a = &sideEffectActor{
			common: c,
			s2s:    fp,
			c2s:    sp,
			db:     db,
			clock:  cl,
		}
		return
	}
	// Run tests
	t.Run("AuthenticatePostInbox", func(t *testing.T) {
		// Setup
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		_, fp, _, _, _, a := setupFn(ctl)
		req := toAPRequest(toPostInboxRequest(testCreate))
		fp.EXPECT().AuthenticatePostInbox(ctx, resp, req).Return(true, testErr)
		// Run
		b, err := a.AuthenticatePostInbox(ctx, resp, req)
		// Verify
		assertEqual(t, b, true)
		assertEqual(t, err, testErr)
	})
	t.Run("AuthenticateGetInbox", func(t *testing.T) {
		// Setup
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		c, _, _, _, _, a := setupFn(ctl)
		req := toAPRequest(toGetInboxRequest())
		c.EXPECT().AuthenticateGetInbox(ctx, resp, req).Return(true, testErr)
		// Run
		b, err := a.AuthenticateGetInbox(ctx, resp, req)
		// Verify
		assertEqual(t, b, true)
		assertEqual(t, err, testErr)
	})
	t.Run("AuthenticatePostOutbox", func(t *testing.T) {
		// Setup
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		_, _, sp, _, _, a := setupFn(ctl)
		req := toAPRequest(toPostOutboxRequest(testCreate))
		sp.EXPECT().AuthenticatePostOutbox(ctx, resp, req).Return(true, testErr)
		// Run
		b, err := a.AuthenticatePostOutbox(ctx, resp, req)
		// Verify
		assertEqual(t, b, true)
		assertEqual(t, err, testErr)
	})
	t.Run("AuthenticateGetOutbox", func(t *testing.T) {
		// Setup
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		c, _, _, _, _, a := setupFn(ctl)
		req := toAPRequest(toGetOutboxRequest())
		c.EXPECT().AuthenticateGetOutbox(ctx, resp, req).Return(true, testErr)
		// Run
		b, err := a.AuthenticateGetOutbox(ctx, resp, req)
		// Verify
		assertEqual(t, b, true)
		assertEqual(t, err, testErr)
	})
	t.Run("GetOutbox", func(t *testing.T) {
		// Setup
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		_, _, sp, _, _, a := setupFn(ctl)
		req := toAPRequest(toGetOutboxRequest())
		sp.EXPECT().GetOutbox(ctx, req).Return(testOrderedCollectionUniqueElems, testErr)
		// Run
		p, err := a.GetOutbox(ctx, req)
		// Verify
		assertEqual(t, p, testOrderedCollectionUniqueElems)
		assertEqual(t, err, testErr)
	})
	t.Run("GetInbox", func(t *testing.T) {
		// Setup
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		_, fp, _, _, _, a := setupFn(ctl)
		req := toAPRequest(toGetInboxRequest())
		fp.EXPECT().GetInbox(ctx, req).Return(testOrderedCollectionUniqueElems, testErr)
		// Run
		p, err := a.GetInbox(ctx, req)
		// Verify
		assertEqual(t, p, testOrderedCollectionUniqueElems)
		assertEqual(t, err, testErr)
	})
}

// TestAuthorizePostInbox tests the Authorization for a federated message, which
// is only based on blocks.
func TestAuthorizePostInbox(t *testing.T) {
	ctx := context.Background()
	resp := httptest.NewRecorder()
	setupFn := func(ctl *gomock.Controller) (c *MockCommonBehavior, fp *MockFederatingProtocol, sp *MockSocialProtocol, db *MockDatabase, cl *MockClock, a DelegateActor) {
		setupData()
		c = NewMockCommonBehavior(ctl)
		fp = NewMockFederatingProtocol(ctl)
		sp = NewMockSocialProtocol(ctl)
		db = NewMockDatabase(ctl)
		cl = NewMockClock(ctl)
		a = &sideEffectActor{
			common: c,
			s2s:    fp,
			c2s:    sp,
			db:     db,
			clock:  cl,
		}
		return
	}
	// Run tests
	t.Run("ActorAuthorized", func(t *testing.T) {
		// Setup
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		_, fp, _, _, _, a := setupFn(ctl)
		fp.EXPECT().Blocked(ctx, []*url.URL{mustParse(testFederatedActorIRI)}).Return(false, nil)
		// Run
		b, err := a.AuthorizePostInbox(ctx, resp, testCreate)
		// Verify
		assertEqual(t, b, true)
		assertEqual(t, err, nil)
	})
	t.Run("ActorNotAuthorized", func(t *testing.T) {
		// Setup
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		_, fp, _, _, _, a := setupFn(ctl)
		fp.EXPECT().Blocked(ctx, []*url.URL{mustParse(testFederatedActorIRI)}).Return(true, nil)
		// Run
		b, err := a.AuthorizePostInbox(ctx, resp, testCreate)
		// Verify
		assertEqual(t, b, false)
		assertEqual(t, err, nil)
	})
	t.Run("AllActorsAuthorized", func(t *testing.T) {
		// Setup
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		_, fp, _, _, _, a := setupFn(ctl)
		fp.EXPECT().Blocked(ctx, []*url.URL{mustParse(testFederatedActorIRI), mustParse(testFederatedActorIRI2)}).Return(false, nil)
		// Run
		b, err := a.AuthorizePostInbox(ctx, resp, testCreate2)
		// Verify
		assertEqual(t, b, true)
		assertEqual(t, err, nil)
	})
	t.Run("OneActorNotAuthorized", func(t *testing.T) {
		// Setup
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		_, fp, _, _, _, a := setupFn(ctl)
		fp.EXPECT().Blocked(ctx, []*url.URL{mustParse(testFederatedActorIRI), mustParse(testFederatedActorIRI2)}).Return(true, nil)
		// Run
		b, err := a.AuthorizePostInbox(ctx, resp, testCreate2)
		// Verify
		assertEqual(t, b, false)
		assertEqual(t, err, nil)
	})
}

// TestPostInbox ensures that the main application side effects of receiving a
// federated message occur.
func TestPostInbox(t *testing.T) {
	ctx := context.Background()
	setupFn := func(ctl *gomock.Controller) (c *MockCommonBehavior, fp *MockFederatingProtocol, sp *MockSocialProtocol, db *MockDatabase, cl *MockClock, a DelegateActor) {
		setupData()
		c = NewMockCommonBehavior(ctl)
		fp = NewMockFederatingProtocol(ctl)
		sp = NewMockSocialProtocol(ctl)
		db = NewMockDatabase(ctl)
		cl = NewMockClock(ctl)
		a = &sideEffectActor{
			common: c,
			s2s:    fp,
			c2s:    sp,
			db:     db,
			clock:  cl,
		}
		return
	}
	// Run tests
	t.Run("AddsToEmptyInbox", func(t *testing.T) {
		// Setup
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		_, fp, _, db, _, a := setupFn(ctl)
		inboxIRI := mustParse(testMyInboxIRI)
		gomock.InOrder(
			db.EXPECT().Lock(ctx, inboxIRI),
			db.EXPECT().InboxContains(ctx, inboxIRI, mustParse(testFederatedActivityIRI)).Return(false, nil),
			db.EXPECT().GetInbox(ctx, inboxIRI).Return(testEmptyOrderedCollection, nil),
			db.EXPECT().SetInbox(ctx, testOrderedCollectionWithFederatedId).Return(nil),
			db.EXPECT().Unlock(ctx, inboxIRI),
		)
		fp.EXPECT().Callbacks(ctx).Return(FederatingWrappedCallbacks{}, nil)
		fp.EXPECT().DefaultCallback(ctx, testListen).Return(nil)
		// Run
		err := a.PostInbox(ctx, inboxIRI, testListen)
		// Verify
		assertEqual(t, err, nil)
	})
	t.Run("DoesNotAddToInboxNorDoSideEffectsIfDuplicate", func(t *testing.T) {
		// Setup
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		_, _, _, db, _, a := setupFn(ctl)
		inboxIRI := mustParse(testMyInboxIRI)
		gomock.InOrder(
			db.EXPECT().Lock(ctx, inboxIRI),
			db.EXPECT().InboxContains(ctx, inboxIRI, mustParse(testFederatedActivityIRI)).Return(true, nil),
			db.EXPECT().Unlock(ctx, inboxIRI),
		)
		// Run
		err := a.PostInbox(ctx, inboxIRI, testListen)
		// Verify
		assertEqual(t, err, nil)
	})
	t.Run("AddsToInbox", func(t *testing.T) {
		// Setup
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		_, fp, _, db, _, a := setupFn(ctl)
		inboxIRI := mustParse(testMyInboxIRI)
		gomock.InOrder(
			db.EXPECT().Lock(ctx, inboxIRI),
			db.EXPECT().InboxContains(ctx, inboxIRI, mustParse(testFederatedActivityIRI)).Return(false, nil),
			db.EXPECT().GetInbox(ctx, inboxIRI).Return(testOrderedCollectionWithFederatedId2, nil),
			db.EXPECT().SetInbox(ctx, testOrderedCollectionWithBothFederatedIds).Return(nil),
			db.EXPECT().Unlock(ctx, inboxIRI),
		)
		fp.EXPECT().Callbacks(ctx).Return(FederatingWrappedCallbacks{}, nil)
		fp.EXPECT().DefaultCallback(ctx, testListen).Return(nil)
		// Run
		err := a.PostInbox(ctx, inboxIRI, testListen)
		// Verify
		assertEqual(t, err, nil)
	})
	t.Run("ResolvesToCustomFunction", func(t *testing.T) {
		// Setup
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		_, fp, _, db, _, a := setupFn(ctl)
		inboxIRI := mustParse(testMyInboxIRI)
		gomock.InOrder(
			db.EXPECT().Lock(ctx, inboxIRI),
			db.EXPECT().InboxContains(ctx, inboxIRI, mustParse(testFederatedActivityIRI)).Return(false, nil),
			db.EXPECT().GetInbox(ctx, inboxIRI).Return(testEmptyOrderedCollection, nil),
			db.EXPECT().SetInbox(ctx, testOrderedCollectionWithFederatedId).Return(nil),
			db.EXPECT().Unlock(ctx, inboxIRI),
		)
		pass := false
		fp.EXPECT().Callbacks(ctx).Return(FederatingWrappedCallbacks{}, []interface{}{
			func(c context.Context, a vocab.ActivityStreamsListen) error {
				pass = true
				return nil
			},
		})
		// Run
		err := a.PostInbox(ctx, inboxIRI, testListen)
		// Verify
		assertEqual(t, err, nil)
		assertEqual(t, pass, true)
	})
	t.Run("ResolvesToOverriddenFunction", func(t *testing.T) {
		// Setup
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		_, fp, _, db, _, a := setupFn(ctl)
		inboxIRI := mustParse(testMyInboxIRI)
		gomock.InOrder(
			db.EXPECT().Lock(ctx, inboxIRI),
			db.EXPECT().InboxContains(ctx, inboxIRI, mustParse(testFederatedActivityIRI)).Return(false, nil),
			db.EXPECT().GetInbox(ctx, inboxIRI).Return(testEmptyOrderedCollection, nil),
			db.EXPECT().SetInbox(ctx, testOrderedCollectionWithFederatedId).Return(nil),
			db.EXPECT().Unlock(ctx, inboxIRI),
		)
		pass := false
		fp.EXPECT().Callbacks(ctx).Return(FederatingWrappedCallbacks{}, []interface{}{
			func(c context.Context, a vocab.ActivityStreamsCreate) error {
				pass = true
				return nil
			},
		})
		// Run
		err := a.PostInbox(ctx, inboxIRI, testCreate)
		// Verify
		assertEqual(t, err, nil)
		assertEqual(t, pass, true)
	})
	t.Run("ResolvesToDefaultFunction", func(t *testing.T) {
		// Setup
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		_, fp, _, db, _, a := setupFn(ctl)
		inboxIRI := mustParse(testMyInboxIRI)
		gomock.InOrder(
			db.EXPECT().Lock(ctx, inboxIRI),
			db.EXPECT().InboxContains(ctx, inboxIRI, mustParse(testFederatedActivityIRI)).Return(false, nil),
			db.EXPECT().GetInbox(ctx, inboxIRI).Return(testEmptyOrderedCollection, nil),
			db.EXPECT().SetInbox(ctx, testOrderedCollectionWithFederatedId).Return(nil),
			db.EXPECT().Unlock(ctx, inboxIRI),
		)
		pass := false
		fp.EXPECT().Callbacks(ctx).Return(FederatingWrappedCallbacks{
			Create: func(c context.Context, a vocab.ActivityStreamsCreate) error {
				pass = true
				return nil
			},
		}, nil)
		db.EXPECT().Lock(ctx, mustParse(testNoteId1))
		db.EXPECT().Create(ctx, testFederatedNote)
		db.EXPECT().Unlock(ctx, mustParse(testNoteId1))
		// Run
		err := a.PostInbox(ctx, inboxIRI, testCreate)
		// Verify
		assertEqual(t, err, nil)
		assertEqual(t, pass, true)
	})
}

// TestInboxForwarding ensures that the inbox forwarding logic is correct.
func TestInboxForwarding(t *testing.T) {
	t.Run("DoesNotForwardIfAlreadyExists", func(t *testing.T) {
		t.Fail()
	})
	t.Run("DoesNotForwardIfToCollectionNotOwned", func(t *testing.T) {
		t.Fail()
	})
	t.Run("DoesNotForwardIfCcCollectionNotOwned", func(t *testing.T) {
		t.Fail()
	})
	t.Run("DoesNotForwardIfAudienceCollectionNotOwned", func(t *testing.T) {
		t.Fail()
	})
	t.Run("DoesNotForwardIfToOwnedButNotCollection", func(t *testing.T) {
		t.Fail()
	})
	t.Run("DoesNotForwardIfCcOwnedButNotCollection", func(t *testing.T) {
		t.Fail()
	})
	t.Run("DoesNotForwardIfAudienceOwnedButNotCollection", func(t *testing.T) {
		t.Fail()
	})
	t.Run("DoesNotForwardIfNoChainIsOwned", func(t *testing.T) {
		t.Fail()
	})
	t.Run("ForwardsToRecipients", func(t *testing.T) {
		t.Fail()
	})
	t.Run("ForwardsToRecipientsIfChainIsNested", func(t *testing.T) {
		t.Fail()
	})
	t.Run("DoesNotForwardIfChainIsNestedTooDeep", func(t *testing.T) {
		t.Fail()
	})
	t.Run("ForwardsToRecipientsIfChainNeedsDereferencing", func(t *testing.T) {
		t.Fail()
	})
}

// TestPostOutbox ensures that the main application side effects of receiving a
// social protocol message occur.
func TestPostOutbox(t *testing.T) {
	t.Run("AddsToEmptyOutbox", func(t *testing.T) {
		t.Fail()
	})
	t.Run("AddsToOutbox", func(t *testing.T) {
		t.Fail()
	})
	t.Run("ResolvesToCustomFunction", func(t *testing.T) {
		t.Fail()
	})
	t.Run("ResolvesToOverriddenFunction", func(t *testing.T) {
		t.Fail()
	})
	t.Run("ResolvesToDefaultFunction", func(t *testing.T) {
		t.Fail()
	})
}

// TestAddNewIds ensures that new 'id' properties are set on an activity and all
// of its 'object' property values if it is a Create activity.
func TestAddNewIds(t *testing.T) {
	t.Run("AddsIdToActivity", func(t *testing.T) {
		t.Fail()
	})
	t.Run("AddsIdsToObjectsIfCreateActivity", func(t *testing.T) {
		t.Fail()
	})
	t.Run("DoesNotAddIdsToObjectsIfNotCreateActivity", func(t *testing.T) {
		t.Fail()
	})
}

// TestDeliver ensures federated delivery of an activity happens correctly to
// the ActivityPub specification.
func TestDeliver(t *testing.T) {
	t.Run("SendToRecipientsInTo", func(t *testing.T) {
		t.Fail()
	})
	t.Run("SendToRecipientsInBto", func(t *testing.T) {
		t.Fail()
	})
	t.Run("SendToRecipientsInCc", func(t *testing.T) {
		t.Fail()
	})
	t.Run("SendToRecipientsInBcc", func(t *testing.T) {
		t.Fail()
	})
	t.Run("SendToRecipientsInAudience", func(t *testing.T) {
		t.Fail()
	})
	t.Run("DoesNotSendToPublicIRI", func(t *testing.T) {
		t.Fail()
	})
	t.Run("RecursivelyResolveCollectionActors", func(t *testing.T) {
		t.Fail()
	})
	t.Run("RecursivelyResolveOrderedCollectionActors", func(t *testing.T) {
		t.Fail()
	})
	t.Run("DoesNotRecursivelyResolveCollectionActorsIfExceedingMaxDepth", func(t *testing.T) {
		t.Fail()
	})
	t.Run("DoesNotSendIfMoreThanOneAttributedTo", func(t *testing.T) {
		t.Fail()
	})
	t.Run("DoesNotSendIfThisActorIsNotInAttributedTo", func(t *testing.T) {
		t.Fail()
	})
	t.Run("DedupesRecipients", func(t *testing.T) {
		t.Fail()
	})
	t.Run("StripsBto", func(t *testing.T) {
		t.Fail()
	})
	t.Run("StripsBcc", func(t *testing.T) {
		t.Fail()
	})
	t.Run("ReturnsErrorIfAnyTransportRequestsFail", func(t *testing.T) {
		t.Fail()
	})
}

// TestWrapInCreate ensures an object received by the Social Protocol is
// properly wrapped in a Create Activity.
func TestWrapInCreate(t *testing.T) {
	t.Run("CreateHasTo", func(t *testing.T) {
		t.Fail()
	})
	t.Run("CreateHasCc", func(t *testing.T) {
		t.Fail()
	})
	t.Run("CreateHasBto", func(t *testing.T) {
		t.Fail()
	})
	t.Run("CreateHasBcc", func(t *testing.T) {
		t.Fail()
	})
	t.Run("CreateHasAudience", func(t *testing.T) {
		t.Fail()
	})
	t.Run("CreateHasPublished", func(t *testing.T) {
		t.Fail()
	})
	t.Run("CreateHasActor", func(t *testing.T) {
		t.Fail()
	})
	t.Run("CreateHasObject", func(t *testing.T) {
		t.Fail()
	})
}

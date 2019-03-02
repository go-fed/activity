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
	t.Run("DoesNotForwardIfAlreadyExists", func(t *testing.T) {
		// Setup
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		_, _, _, db, _, a := setupFn(ctl)
		gomock.InOrder(
			db.EXPECT().Lock(ctx, mustParse(testFederatedActivityIRI)),
			db.EXPECT().Exists(ctx, mustParse(testFederatedActivityIRI)).Return(true, nil),
			db.EXPECT().Unlock(ctx, mustParse(testFederatedActivityIRI)),
		)
		// Run
		err := a.InboxForwarding(ctx, mustParse(testMyInboxIRI), testListen)
		// Verify
		assertEqual(t, err, nil)
	})
	t.Run("DoesNotForwardIfToCollectionNotOwned", func(t *testing.T) {
		// Setup
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		_, _, _, db, _, a := setupFn(ctl)
		input := addToIds(testListen)
		gomock.InOrder(
			db.EXPECT().Lock(ctx, mustParse(testFederatedActivityIRI)),
			db.EXPECT().Exists(ctx, mustParse(testFederatedActivityIRI)).Return(false, nil),
			db.EXPECT().Create(ctx, input).Return(nil),
			db.EXPECT().Unlock(ctx, mustParse(testFederatedActivityIRI)),
			db.EXPECT().Lock(ctx, mustParse(testToIRI)),
			db.EXPECT().Owns(ctx, mustParse(testToIRI)).Return(false, nil),
			db.EXPECT().Unlock(ctx, mustParse(testToIRI)),
			db.EXPECT().Lock(ctx, mustParse(testToIRI2)),
			db.EXPECT().Owns(ctx, mustParse(testToIRI2)).Return(false, nil),
			db.EXPECT().Unlock(ctx, mustParse(testToIRI2)),
		)
		// Run
		err := a.InboxForwarding(ctx, mustParse(testMyInboxIRI), input)
		// Verify
		assertEqual(t, err, nil)
	})
	t.Run("DoesNotForwardIfCcCollectionNotOwned", func(t *testing.T) {
		// Setup
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		_, _, _, db, _, a := setupFn(ctl)
		input := mustAddCcIds(testListen)
		gomock.InOrder(
			db.EXPECT().Lock(ctx, mustParse(testFederatedActivityIRI)),
			db.EXPECT().Exists(ctx, mustParse(testFederatedActivityIRI)).Return(false, nil),
			db.EXPECT().Create(ctx, input).Return(nil),
			db.EXPECT().Unlock(ctx, mustParse(testFederatedActivityIRI)),
			db.EXPECT().Lock(ctx, mustParse(testCcIRI)),
			db.EXPECT().Owns(ctx, mustParse(testCcIRI)).Return(false, nil),
			db.EXPECT().Unlock(ctx, mustParse(testCcIRI)),
			db.EXPECT().Lock(ctx, mustParse(testCcIRI2)),
			db.EXPECT().Owns(ctx, mustParse(testCcIRI2)).Return(false, nil),
			db.EXPECT().Unlock(ctx, mustParse(testCcIRI2)),
		)
		// Run
		err := a.InboxForwarding(ctx, mustParse(testMyInboxIRI), input)
		// Verify
		assertEqual(t, err, nil)
	})
	t.Run("DoesNotForwardIfAudienceCollectionNotOwned", func(t *testing.T) {
		// Setup
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		_, _, _, db, _, a := setupFn(ctl)
		input := mustAddAudienceIds(testListen)
		gomock.InOrder(
			db.EXPECT().Lock(ctx, mustParse(testFederatedActivityIRI)),
			db.EXPECT().Exists(ctx, mustParse(testFederatedActivityIRI)).Return(false, nil),
			db.EXPECT().Create(ctx, input).Return(nil),
			db.EXPECT().Unlock(ctx, mustParse(testFederatedActivityIRI)),
			db.EXPECT().Lock(ctx, mustParse(testAudienceIRI)),
			db.EXPECT().Owns(ctx, mustParse(testAudienceIRI)).Return(false, nil),
			db.EXPECT().Unlock(ctx, mustParse(testAudienceIRI)),
			db.EXPECT().Lock(ctx, mustParse(testAudienceIRI2)),
			db.EXPECT().Owns(ctx, mustParse(testAudienceIRI2)).Return(false, nil),
			db.EXPECT().Unlock(ctx, mustParse(testAudienceIRI2)),
		)
		// Run
		err := a.InboxForwarding(ctx, mustParse(testMyInboxIRI), input)
		// Verify
		assertEqual(t, err, nil)
	})
	t.Run("DoesNotForwardIfToOwnedButNotCollection", func(t *testing.T) {
		// Setup
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		_, _, _, db, _, a := setupFn(ctl)
		input := addToIds(testListen)
		gomock.InOrder(
			db.EXPECT().Lock(ctx, mustParse(testFederatedActivityIRI)),
			db.EXPECT().Exists(ctx, mustParse(testFederatedActivityIRI)).Return(false, nil),
			db.EXPECT().Create(ctx, input).Return(nil),
			db.EXPECT().Unlock(ctx, mustParse(testFederatedActivityIRI)),
			db.EXPECT().Lock(ctx, mustParse(testToIRI)),
			db.EXPECT().Owns(ctx, mustParse(testToIRI)).Return(true, nil),
			db.EXPECT().Unlock(ctx, mustParse(testToIRI)),
			db.EXPECT().Lock(ctx, mustParse(testToIRI2)),
			db.EXPECT().Owns(ctx, mustParse(testToIRI2)).Return(true, nil),
			db.EXPECT().Unlock(ctx, mustParse(testToIRI2)),
			db.EXPECT().Lock(ctx, mustParse(testToIRI)),
			db.EXPECT().Get(ctx, mustParse(testToIRI)).Return(testPerson, nil),
			db.EXPECT().Unlock(ctx, mustParse(testToIRI)),
			db.EXPECT().Lock(ctx, mustParse(testToIRI2)),
			db.EXPECT().Get(ctx, mustParse(testToIRI2)).Return(testService, nil),
			db.EXPECT().Unlock(ctx, mustParse(testToIRI2)),
			// Deferred
			db.EXPECT().Unlock(ctx, mustParse(testToIRI2)),
			db.EXPECT().Unlock(ctx, mustParse(testToIRI)),
		)
		// Run
		err := a.InboxForwarding(ctx, mustParse(testMyInboxIRI), input)
		// Verify
		assertEqual(t, err, nil)
	})
	t.Run("DoesNotForwardIfCcOwnedButNotCollection", func(t *testing.T) {
		// Setup
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		_, _, _, db, _, a := setupFn(ctl)
		input := mustAddCcIds(testListen)
		gomock.InOrder(
			db.EXPECT().Lock(ctx, mustParse(testFederatedActivityIRI)),
			db.EXPECT().Exists(ctx, mustParse(testFederatedActivityIRI)).Return(false, nil),
			db.EXPECT().Create(ctx, input).Return(nil),
			db.EXPECT().Unlock(ctx, mustParse(testFederatedActivityIRI)),
			db.EXPECT().Lock(ctx, mustParse(testCcIRI)),
			db.EXPECT().Owns(ctx, mustParse(testCcIRI)).Return(true, nil),
			db.EXPECT().Unlock(ctx, mustParse(testCcIRI)),
			db.EXPECT().Lock(ctx, mustParse(testCcIRI2)),
			db.EXPECT().Owns(ctx, mustParse(testCcIRI2)).Return(true, nil),
			db.EXPECT().Unlock(ctx, mustParse(testCcIRI2)),
			db.EXPECT().Lock(ctx, mustParse(testCcIRI)),
			db.EXPECT().Get(ctx, mustParse(testCcIRI)).Return(testPerson, nil),
			db.EXPECT().Unlock(ctx, mustParse(testCcIRI)),
			db.EXPECT().Lock(ctx, mustParse(testCcIRI2)),
			db.EXPECT().Get(ctx, mustParse(testCcIRI2)).Return(testService, nil),
			db.EXPECT().Unlock(ctx, mustParse(testCcIRI2)),
			// Deferred
			db.EXPECT().Unlock(ctx, mustParse(testCcIRI2)),
			db.EXPECT().Unlock(ctx, mustParse(testCcIRI)),
		)
		// Run
		err := a.InboxForwarding(ctx, mustParse(testMyInboxIRI), input)
		// Verify
		assertEqual(t, err, nil)
	})
	t.Run("DoesNotForwardIfAudienceOwnedButNotCollection", func(t *testing.T) {
		// Setup
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		_, _, _, db, _, a := setupFn(ctl)
		input := mustAddAudienceIds(testListen)
		gomock.InOrder(
			db.EXPECT().Lock(ctx, mustParse(testFederatedActivityIRI)),
			db.EXPECT().Exists(ctx, mustParse(testFederatedActivityIRI)).Return(false, nil),
			db.EXPECT().Create(ctx, input).Return(nil),
			db.EXPECT().Unlock(ctx, mustParse(testFederatedActivityIRI)),
			db.EXPECT().Lock(ctx, mustParse(testAudienceIRI)),
			db.EXPECT().Owns(ctx, mustParse(testAudienceIRI)).Return(true, nil),
			db.EXPECT().Unlock(ctx, mustParse(testAudienceIRI)),
			db.EXPECT().Lock(ctx, mustParse(testAudienceIRI2)),
			db.EXPECT().Owns(ctx, mustParse(testAudienceIRI2)).Return(true, nil),
			db.EXPECT().Unlock(ctx, mustParse(testAudienceIRI2)),
			db.EXPECT().Lock(ctx, mustParse(testAudienceIRI)),
			db.EXPECT().Get(ctx, mustParse(testAudienceIRI)).Return(testPerson, nil),
			db.EXPECT().Unlock(ctx, mustParse(testAudienceIRI)),
			db.EXPECT().Lock(ctx, mustParse(testAudienceIRI2)),
			db.EXPECT().Get(ctx, mustParse(testAudienceIRI2)).Return(testService, nil),
			db.EXPECT().Unlock(ctx, mustParse(testAudienceIRI2)),
			// Deferred
			db.EXPECT().Unlock(ctx, mustParse(testAudienceIRI2)),
			db.EXPECT().Unlock(ctx, mustParse(testAudienceIRI)),
		)
		// Run
		err := a.InboxForwarding(ctx, mustParse(testMyInboxIRI), input)
		// Verify
		assertEqual(t, err, nil)
	})
	t.Run("DoesNotForwardIfNoChainIsOwned", func(t *testing.T) {
		// Setup
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		cm, fp, _, db, _, a := setupFn(ctl)
		input := mustAddTagIds(
			mustAddAudienceIds(testListen))
		mockTPortTag := NewMockTransport(ctl)
		mockTPortTag2 := NewMockTransport(ctl)
		gomock.InOrder(
			db.EXPECT().Lock(ctx, mustParse(testFederatedActivityIRI)),
			db.EXPECT().Exists(ctx, mustParse(testFederatedActivityIRI)).Return(false, nil),
			db.EXPECT().Create(ctx, input).Return(nil),
			db.EXPECT().Unlock(ctx, mustParse(testFederatedActivityIRI)),
			db.EXPECT().Lock(ctx, mustParse(testAudienceIRI)),
			db.EXPECT().Owns(ctx, mustParse(testAudienceIRI)).Return(true, nil),
			db.EXPECT().Unlock(ctx, mustParse(testAudienceIRI)),
			db.EXPECT().Lock(ctx, mustParse(testAudienceIRI2)),
			db.EXPECT().Owns(ctx, mustParse(testAudienceIRI2)).Return(true, nil),
			db.EXPECT().Unlock(ctx, mustParse(testAudienceIRI2)),
			db.EXPECT().Lock(ctx, mustParse(testAudienceIRI)),
			db.EXPECT().Get(ctx, mustParse(testAudienceIRI)).Return(testOrderedCollectionOfActors, nil),
			db.EXPECT().Lock(ctx, mustParse(testAudienceIRI2)),
			db.EXPECT().Get(ctx, mustParse(testAudienceIRI2)).Return(testCollectionOfActors, nil),
			fp.EXPECT().MaxInboxForwardingRecursionDepth(ctx).Return(0),
			// hasInboxForwardingValues
			db.EXPECT().Lock(ctx, mustParse(testTagIRI)),
			db.EXPECT().Owns(ctx, mustParse(testTagIRI)).Return(false, nil),
			db.EXPECT().Unlock(ctx, mustParse(testTagIRI)),
			db.EXPECT().Lock(ctx, mustParse(testTagIRI2)),
			db.EXPECT().Owns(ctx, mustParse(testTagIRI2)).Return(false, nil),
			db.EXPECT().Unlock(ctx, mustParse(testTagIRI2)),
			db.EXPECT().Lock(ctx, mustParse(testNoteId1)),
			db.EXPECT().Owns(ctx, mustParse(testNoteId1)).Return(false, nil),
			db.EXPECT().Unlock(ctx, mustParse(testNoteId1)),
			cm.EXPECT().NewTransport(ctx, mustParse(testMyInboxIRI), goFedUserAgent()).Return(mockTPortTag, nil),
			mockTPortTag.EXPECT().Dereference(ctx, mustParse(testTagIRI)).Return(mustSerializeToBytes(newObjectWithId(testTagIRI)), nil),
			cm.EXPECT().NewTransport(ctx, mustParse(testMyInboxIRI), goFedUserAgent()).Return(mockTPortTag2, nil),
			mockTPortTag2.EXPECT().Dereference(ctx, mustParse(testTagIRI2)).Return(mustSerializeToBytes(newObjectWithId(testTagIRI2)), nil),
			// Deferred
			db.EXPECT().Unlock(ctx, mustParse(testAudienceIRI2)),
			db.EXPECT().Unlock(ctx, mustParse(testAudienceIRI)),
		)
		// Run
		err := a.InboxForwarding(ctx, mustParse(testMyInboxIRI), input)
		// Verify
		assertEqual(t, err, nil)
	})
	t.Run("ForwardsToRecipients", func(t *testing.T) {
		// Setup
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		cm, fp, _, db, _, a := setupFn(ctl)
		input := mustAddTagIds(
			mustAddAudienceIds(testListen))
		tPort := NewMockTransport(ctl)
		gomock.InOrder(
			db.EXPECT().Lock(ctx, mustParse(testFederatedActivityIRI)),
			db.EXPECT().Exists(ctx, mustParse(testFederatedActivityIRI)).Return(false, nil),
			db.EXPECT().Create(ctx, input).Return(nil),
			db.EXPECT().Unlock(ctx, mustParse(testFederatedActivityIRI)),
			db.EXPECT().Lock(ctx, mustParse(testAudienceIRI)),
			db.EXPECT().Owns(ctx, mustParse(testAudienceIRI)).Return(true, nil),
			db.EXPECT().Unlock(ctx, mustParse(testAudienceIRI)),
			db.EXPECT().Lock(ctx, mustParse(testAudienceIRI2)),
			db.EXPECT().Owns(ctx, mustParse(testAudienceIRI2)).Return(true, nil),
			db.EXPECT().Unlock(ctx, mustParse(testAudienceIRI2)),
			db.EXPECT().Lock(ctx, mustParse(testAudienceIRI)),
			db.EXPECT().Get(ctx, mustParse(testAudienceIRI)).Return(testOrderedCollectionOfActors, nil),
			db.EXPECT().Lock(ctx, mustParse(testAudienceIRI2)),
			db.EXPECT().Get(ctx, mustParse(testAudienceIRI2)).Return(testCollectionOfActors, nil),
			fp.EXPECT().MaxInboxForwardingRecursionDepth(ctx).Return(0),
			// hasInboxForwardingValues
			db.EXPECT().Lock(ctx, mustParse(testTagIRI)),
			db.EXPECT().Owns(ctx, mustParse(testTagIRI)).Return(true, nil),
			db.EXPECT().Unlock(ctx, mustParse(testTagIRI)),
			// after hasInboxForwardingValues
			fp.EXPECT().FilterForwarding(
				ctx,
				[]*url.URL{
					mustParse(testAudienceIRI),
					mustParse(testAudienceIRI2),
				},
				input,
			).Return(
				[]*url.URL{
					mustParse(testAudienceIRI),
				},
				nil,
			),
			// deliverToRecipients
			cm.EXPECT().NewTransport(ctx, mustParse(testMyInboxIRI), goFedUserAgent()).Return(tPort, nil),
			tPort.EXPECT().BatchDeliver(
				ctx,
				mustSerializeToBytes(input),
				[]*url.URL{
					mustParse(testFederatedActorIRI3),
					mustParse(testFederatedActorIRI4),
				},
			),
			// Deferred
			db.EXPECT().Unlock(ctx, mustParse(testAudienceIRI2)),
			db.EXPECT().Unlock(ctx, mustParse(testAudienceIRI)),
		)
		// Run
		err := a.InboxForwarding(ctx, mustParse(testMyInboxIRI), input)
		// Verify
		assertEqual(t, err, nil)
	})
	t.Run("ForwardsToRecipientsIfChainIsNested", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("ForwardsToRecipientsAfterDereferencing", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("DoesNotForwardIfChainIsNestedTooDeep", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("ForwardsToRecipientsIfChainNeedsDereferencing", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
}

// TestPostOutbox ensures that the main application side effects of receiving a
// social protocol message occur.
func TestPostOutbox(t *testing.T) {
	t.Run("AddsToEmptyOutbox", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("AddsToOutbox", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("ResolvesToCustomFunction", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("ResolvesToOverriddenFunction", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("ResolvesToDefaultFunction", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
}

// TestAddNewIds ensures that new 'id' properties are set on an activity and all
// of its 'object' property values if it is a Create activity.
func TestAddNewIds(t *testing.T) {
	t.Run("AddsIdToActivity", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("AddsIdsToObjectsIfCreateActivity", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("DoesNotAddIdsToObjectsIfNotCreateActivity", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
}

// TestDeliver ensures federated delivery of an activity happens correctly to
// the ActivityPub specification.
func TestDeliver(t *testing.T) {
	t.Run("SendToRecipientsInTo", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("SendToRecipientsInBto", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("SendToRecipientsInCc", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("SendToRecipientsInBcc", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("SendToRecipientsInAudience", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("DoesNotSendToPublicIRI", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("RecursivelyResolveCollectionActors", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("RecursivelyResolveOrderedCollectionActors", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("DoesNotRecursivelyResolveCollectionActorsIfExceedingMaxDepth", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("DoesNotSendIfMoreThanOneAttributedTo", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("DoesNotSendIfThisActorIsNotInAttributedTo", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("DedupesRecipients", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("StripsBto", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("StripsBcc", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("ReturnsErrorIfAnyTransportRequestsFail", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
}

// TestWrapInCreate ensures an object received by the Social Protocol is
// properly wrapped in a Create Activity.
func TestWrapInCreate(t *testing.T) {
	t.Run("CreateHasTo", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("CreateHasCc", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("CreateHasBto", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("CreateHasBcc", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("CreateHasAudience", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("CreateHasPublished", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("CreateHasActor", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("CreateHasObject", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
}

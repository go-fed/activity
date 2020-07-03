package pub

import (
	"context"
	"net/url"
	"testing"

	"github.com/go-fed/activity/streams"
	"github.com/go-fed/activity/streams/vocab"
	"github.com/golang/mock/gomock"
)

// TestFederatedCallbacks tests the overriding functionality.
func TestFederatedCallbacks(t *testing.T) {
	t.Run("ReturnsOtherCallback", func(t *testing.T) {
		ok := false
		o := func(context.Context, vocab.ActivityStreamsListen) error {
			ok = true
			return nil
		}
		var w FederatingWrappedCallbacks
		for _, f := range w.callbacks([]interface{}{o}) {
			if fn, ok := f.(func(context.Context, vocab.ActivityStreamsListen) error); ok {
				fn(nil, nil)
			}
		}
		if !ok {
			t.Fatalf("could not find extra function")
		}

	})
	t.Run("OverridesCreate", func(t *testing.T) {
		ok := false
		o := func(context.Context, vocab.ActivityStreamsCreate) error {
			ok = true
			return nil
		}
		var w FederatingWrappedCallbacks
		for _, f := range w.callbacks([]interface{}{o}) {
			if fn, ok := f.(func(context.Context, vocab.ActivityStreamsCreate) error); ok {
				fn(nil, nil)
			}
		}
		if !ok {
			t.Fatalf("could not find overridden function")
		}
	})
	t.Run("OverridesUpdate", func(t *testing.T) {
		ok := false
		o := func(context.Context, vocab.ActivityStreamsUpdate) error {
			ok = true
			return nil
		}
		var w FederatingWrappedCallbacks
		for _, f := range w.callbacks([]interface{}{o}) {
			if fn, ok := f.(func(context.Context, vocab.ActivityStreamsUpdate) error); ok {
				fn(nil, nil)
			}
		}
		if !ok {
			t.Fatalf("could not find overridden function")
		}
	})
	t.Run("OverridesDelete", func(t *testing.T) {
		ok := false
		o := func(context.Context, vocab.ActivityStreamsDelete) error {
			ok = true
			return nil
		}
		var w FederatingWrappedCallbacks
		for _, f := range w.callbacks([]interface{}{o}) {
			if fn, ok := f.(func(context.Context, vocab.ActivityStreamsDelete) error); ok {
				fn(nil, nil)
			}
		}
		if !ok {
			t.Fatalf("could not find overridden function")
		}
	})
	t.Run("OverridesFollow", func(t *testing.T) {
		ok := false
		o := func(context.Context, vocab.ActivityStreamsFollow) error {
			ok = true
			return nil
		}
		var w FederatingWrappedCallbacks
		for _, f := range w.callbacks([]interface{}{o}) {
			if fn, ok := f.(func(context.Context, vocab.ActivityStreamsFollow) error); ok {
				fn(nil, nil)
			}
		}
		if !ok {
			t.Fatalf("could not find overridden function")
		}
	})
	t.Run("OverridesAccept", func(t *testing.T) {
		ok := false
		o := func(context.Context, vocab.ActivityStreamsAccept) error {
			ok = true
			return nil
		}
		var w FederatingWrappedCallbacks
		for _, f := range w.callbacks([]interface{}{o}) {
			if fn, ok := f.(func(context.Context, vocab.ActivityStreamsAccept) error); ok {
				fn(nil, nil)
			}
		}
		if !ok {
			t.Fatalf("could not find overridden function")
		}
	})
	t.Run("OverridesReject", func(t *testing.T) {
		ok := false
		o := func(context.Context, vocab.ActivityStreamsReject) error {
			ok = true
			return nil
		}
		var w FederatingWrappedCallbacks
		for _, f := range w.callbacks([]interface{}{o}) {
			if fn, ok := f.(func(context.Context, vocab.ActivityStreamsReject) error); ok {
				fn(nil, nil)
			}
		}
		if !ok {
			t.Fatalf("could not find overridden function")
		}
	})
	t.Run("OverridesAdd", func(t *testing.T) {
		ok := false
		o := func(context.Context, vocab.ActivityStreamsAdd) error {
			ok = true
			return nil
		}
		var w FederatingWrappedCallbacks
		for _, f := range w.callbacks([]interface{}{o}) {
			if fn, ok := f.(func(context.Context, vocab.ActivityStreamsAdd) error); ok {
				fn(nil, nil)
			}
		}
		if !ok {
			t.Fatalf("could not find overridden function")
		}
	})
	t.Run("OverridesRemove", func(t *testing.T) {
		ok := false
		o := func(context.Context, vocab.ActivityStreamsRemove) error {
			ok = true
			return nil
		}
		var w FederatingWrappedCallbacks
		for _, f := range w.callbacks([]interface{}{o}) {
			if fn, ok := f.(func(context.Context, vocab.ActivityStreamsRemove) error); ok {
				fn(nil, nil)
			}
		}
		if !ok {
			t.Fatalf("could not find overridden function")
		}
	})
	t.Run("OverridesLike", func(t *testing.T) {
		ok := false
		o := func(context.Context, vocab.ActivityStreamsLike) error {
			ok = true
			return nil
		}
		var w FederatingWrappedCallbacks
		for _, f := range w.callbacks([]interface{}{o}) {
			if fn, ok := f.(func(context.Context, vocab.ActivityStreamsLike) error); ok {
				fn(nil, nil)
			}
		}
		if !ok {
			t.Fatalf("could not find overridden function")
		}
	})
	t.Run("OverridesAnnounce", func(t *testing.T) {
		ok := false
		o := func(context.Context, vocab.ActivityStreamsAnnounce) error {
			ok = true
			return nil
		}
		var w FederatingWrappedCallbacks
		for _, f := range w.callbacks([]interface{}{o}) {
			if fn, ok := f.(func(context.Context, vocab.ActivityStreamsAnnounce) error); ok {
				fn(nil, nil)
			}
		}
		if !ok {
			t.Fatalf("could not find overridden function")
		}
	})
	t.Run("OverridesUndo", func(t *testing.T) {
		ok := false
		o := func(context.Context, vocab.ActivityStreamsUndo) error {
			ok = true
			return nil
		}
		var w FederatingWrappedCallbacks
		for _, f := range w.callbacks([]interface{}{o}) {
			if fn, ok := f.(func(context.Context, vocab.ActivityStreamsUndo) error); ok {
				fn(nil, nil)
			}
		}
		if !ok {
			t.Fatalf("could not find overridden function")
		}
	})
	t.Run("OverridesBlock", func(t *testing.T) {
		ok := false
		o := func(context.Context, vocab.ActivityStreamsBlock) error {
			ok = true
			return nil
		}
		var w FederatingWrappedCallbacks
		for _, f := range w.callbacks([]interface{}{o}) {
			if fn, ok := f.(func(context.Context, vocab.ActivityStreamsBlock) error); ok {
				fn(nil, nil)
			}
		}
		if !ok {
			t.Fatalf("could not find overridden function")
		}
	})
}

func TestFederatedCreate(t *testing.T) {
	newCreateFn := func() vocab.ActivityStreamsCreate {
		c := streams.NewActivityStreamsCreate()
		id := streams.NewJSONLDIdProperty()
		id.Set(mustParse(testFederatedActivityIRI))
		c.SetJSONLDId(id)
		actor := streams.NewActivityStreamsActorProperty()
		actor.AppendIRI(mustParse(testFederatedActorIRI))
		c.SetActivityStreamsActor(actor)
		op := streams.NewActivityStreamsObjectProperty()
		op.AppendActivityStreamsNote(testFederatedNote)
		c.SetActivityStreamsObject(op)
		return c
	}
	ctx := context.Background()
	setupFn := func(ctl *gomock.Controller) (w FederatingWrappedCallbacks, mockDB *MockDatabase, mockTp *MockTransport) {
		mockDB = NewMockDatabase(ctl)
		mockTp = NewMockTransport(ctl)
		w.db = mockDB
		w.newTransport = func(c context.Context, a *url.URL, s string) (Transport, error) {
			return mockTp, nil
		}
		return
	}
	t.Run("ErrorIfNoObject", func(t *testing.T) {
		c := newCreateFn()
		c.SetActivityStreamsObject(nil)
		var w FederatingWrappedCallbacks
		err := w.create(ctx, c)
		if err == nil {
			t.Fatalf("expected error, got none")
		}
	})
	t.Run("ErrorIfObjectLengthZero", func(t *testing.T) {
		c := newCreateFn()
		c.GetActivityStreamsObject().Remove(0)
		var w FederatingWrappedCallbacks
		err := w.create(ctx, c)
		if err == nil {
			t.Fatalf("expected error, got none")
		}
	})
	t.Run("CreatesFederatedObject", func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		w, mockDB, _ := setupFn(ctl)
		mockDB.EXPECT().Lock(ctx, mustParse(testNoteId1))
		mockDB.EXPECT().Create(ctx, testFederatedNote)
		mockDB.EXPECT().Unlock(ctx, mustParse(testNoteId1))
		c := newCreateFn()
		err := w.create(ctx, c)
		if err != nil {
			t.Fatalf("got error %s", err)
		}
	})
	t.Run("CreatesAllFederatedObjects", func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		w, mockDB, _ := setupFn(ctl)
		mockDB.EXPECT().Lock(ctx, mustParse(testNoteId1))
		mockDB.EXPECT().Create(ctx, testFederatedNote)
		mockDB.EXPECT().Unlock(ctx, mustParse(testNoteId1))
		mockDB.EXPECT().Lock(ctx, mustParse(testNoteId2))
		mockDB.EXPECT().Create(ctx, testFederatedNote2)
		mockDB.EXPECT().Unlock(ctx, mustParse(testNoteId2))
		c := newCreateFn()
		c.GetActivityStreamsObject().AppendActivityStreamsNote(testFederatedNote2)
		err := w.create(ctx, c)
		if err != nil {
			t.Fatalf("got error %s", err)
		}
	})
	t.Run("DereferencesIRIObject", func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		w, mockDB, mockTp := setupFn(ctl)
		mockDB.EXPECT().Lock(ctx, mustParse(testNoteId1))
		mockDB.EXPECT().Create(ctx, toDeserializedForm(testFederatedNote))
		mockDB.EXPECT().Unlock(ctx, mustParse(testNoteId1))
		mockTp.EXPECT().Dereference(ctx, mustParse(testNoteId1)).Return(
			mustSerializeToBytes(testFederatedNote), nil)
		c := newCreateFn()
		op := streams.NewActivityStreamsObjectProperty()
		op.AppendIRI(mustParse(testNoteId1))
		c.SetActivityStreamsObject(op)
		err := w.create(ctx, c)
		if err != nil {
			t.Fatalf("got error %s", err)
		}
	})
	t.Run("CallsCustomCallback", func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		w, mockDB, _ := setupFn(ctl)
		mockDB.EXPECT().Lock(ctx, mustParse(testNoteId1))
		mockDB.EXPECT().Create(ctx, testFederatedNote)
		mockDB.EXPECT().Unlock(ctx, mustParse(testNoteId1))
		c := newCreateFn()
		var gotc context.Context
		var got vocab.ActivityStreamsCreate
		w.Create = func(ctx context.Context, v vocab.ActivityStreamsCreate) error {
			gotc = ctx
			got = v
			return nil
		}
		err := w.create(ctx, c)
		if err != nil {
			t.Fatalf("got error %s", err)
		}
		assertEqual(t, ctx, gotc)
		assertEqual(t, c, got)
	})
}

func TestFederatedUpdate(t *testing.T) {
	newUpdateFn := func() vocab.ActivityStreamsUpdate {
		u := streams.NewActivityStreamsUpdate()
		id := streams.NewJSONLDIdProperty()
		id.Set(mustParse(testNewActivityIRI))
		u.SetJSONLDId(id)
		actor := streams.NewActivityStreamsActorProperty()
		actor.AppendIRI(mustParse(testFederatedActorIRI))
		u.SetActivityStreamsActor(actor)
		op := streams.NewActivityStreamsObjectProperty()
		op.AppendActivityStreamsNote(testFederatedNote)
		u.SetActivityStreamsObject(op)
		return u
	}
	ctx := context.Background()
	setupFn := func(ctl *gomock.Controller) (w FederatingWrappedCallbacks, mockDB *MockDatabase) {
		mockDB = NewMockDatabase(ctl)
		w.db = mockDB
		return
	}
	t.Run("ErrorIfNoObject", func(t *testing.T) {
		u := newUpdateFn()
		u.SetActivityStreamsObject(nil)
		var w FederatingWrappedCallbacks
		err := w.update(ctx, u)
		if err == nil {
			t.Fatalf("expected error, got none")
		}
	})
	t.Run("ErrorIfObjectLengthZero", func(t *testing.T) {
		u := newUpdateFn()
		u.GetActivityStreamsObject().Remove(0)
		var w FederatingWrappedCallbacks
		err := w.update(ctx, u)
		if err == nil {
			t.Fatalf("expected error, got none")
		}
	})
	t.Run("ErrorIfOriginMismatchesObject", func(t *testing.T) {
		u := newUpdateFn()
		id := streams.NewJSONLDIdProperty()
		id.Set(mustParse(testFederatedActivityIRI))
		u.SetJSONLDId(id)
		var w FederatingWrappedCallbacks
		err := w.update(ctx, u)
		if err == nil {
			t.Fatalf("expected error, got none")
		}
	})
	t.Run("UpdatesFederatedObject", func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		w, mockDB := setupFn(ctl)
		mockDB.EXPECT().Lock(ctx, mustParse(testNoteId1))
		mockDB.EXPECT().Update(ctx, testFederatedNote)
		mockDB.EXPECT().Unlock(ctx, mustParse(testNoteId1))
		u := newUpdateFn()
		err := w.update(ctx, u)
		if err != nil {
			t.Fatalf("got error %s", err)
		}
	})
	t.Run("UpdatesAllFederatedObjects", func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		w, mockDB := setupFn(ctl)
		mockDB.EXPECT().Lock(ctx, mustParse(testNoteId1))
		mockDB.EXPECT().Update(ctx, testFederatedNote)
		mockDB.EXPECT().Unlock(ctx, mustParse(testNoteId1))
		mockDB.EXPECT().Lock(ctx, mustParse(testNoteId2))
		mockDB.EXPECT().Update(ctx, testFederatedNote2)
		mockDB.EXPECT().Unlock(ctx, mustParse(testNoteId2))
		u := newUpdateFn()
		u.GetActivityStreamsObject().AppendActivityStreamsNote(testFederatedNote2)
		err := w.update(ctx, u)
		if err != nil {
			t.Fatalf("got error %s", err)
		}
	})
	t.Run("ErrorIfObjectIsIRI", func(t *testing.T) {
		u := newUpdateFn()
		op := streams.NewActivityStreamsObjectProperty()
		op.AppendIRI(mustParse(testNoteId1))
		u.SetActivityStreamsObject(op)
		var w FederatingWrappedCallbacks
		err := w.update(ctx, u)
		if err == nil {
			t.Fatalf("expected error, got none")
		}
	})
	t.Run("CallsCustomCallback", func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		w, mockDB := setupFn(ctl)
		mockDB.EXPECT().Lock(ctx, mustParse(testNoteId1))
		mockDB.EXPECT().Update(ctx, testFederatedNote)
		mockDB.EXPECT().Unlock(ctx, mustParse(testNoteId1))
		u := newUpdateFn()
		var gotc context.Context
		var got vocab.ActivityStreamsUpdate
		w.Update = func(ctx context.Context, v vocab.ActivityStreamsUpdate) error {
			gotc = ctx
			got = v
			return nil
		}
		err := w.update(ctx, u)
		if err != nil {
			t.Fatalf("got error %s", err)
		}
		assertEqual(t, ctx, gotc)
		assertEqual(t, u, got)
	})
}

func TestFederatedDelete(t *testing.T) {
	t.Run("ErrorIfNoObject", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("ErrorIfObjectLengthZero", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("ErrorIfOriginMismatchesObject", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("DeletesFederatedObject", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("DeletesAllFederatedObjects", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("CallsCustomCallback", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
}

func TestFederatedFollow(t *testing.T) {
	t.Run("ErrorIfNoObject", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("ErrorIfObjectLengthZero", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("OnFollowNothingDoesNothing", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("OnFollowAutomaticallyAcceptUpdatesFollowers", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("OnFollowAutomaticallyAcceptDelivers", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("OnFollowAutomaticallyRejectDelivers", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("OnFollowNothingCallsCustomCallback", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("OnFollowAutomaticallyAcceptCallsCustomCallback", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("OnFollowAutomaticallyRejectCallsCustomCallback", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
}

func TestFederatedAccept(t *testing.T) {
	newAcceptFn := func() vocab.ActivityStreamsAccept {
		c := streams.NewActivityStreamsAccept()
		id := streams.NewJSONLDIdProperty()
		id.Set(mustParse(testFederatedActivityIRI2))
		c.SetJSONLDId(id)
		actor := streams.NewActivityStreamsActorProperty()
		actor.AppendIRI(mustParse(testFederatedActorIRI))
		c.SetActivityStreamsActor(actor)
		op := streams.NewActivityStreamsObjectProperty()
		op.AppendActivityStreamsFollow(testFollow)
		c.SetActivityStreamsObject(op)
		return c
	}
	ctx := context.Background()
	setupFn := func(ctl *gomock.Controller) (w FederatingWrappedCallbacks, mockDB *MockDatabase, mockTp *MockTransport) {
		mockDB = NewMockDatabase(ctl)
		mockTp = NewMockTransport(ctl)
		w.inboxIRI = mustParse(testMyInboxIRI)
		w.db = mockDB
		w.newTransport = func(c context.Context, a *url.URL, s string) (Transport, error) {
			return mockTp, nil
		}
		return
	}
	t.Run("DoesNothingIfNoObjects", func(t *testing.T) {
		a := newAcceptFn()
		a.SetActivityStreamsObject(nil)
		var w FederatingWrappedCallbacks
		err := w.accept(ctx, a)
		if err != nil {
			t.Fatalf("got error %s", err)
		}
	})
	t.Run("DereferencesObjectIRI", func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		w, mockDB, mockTp := setupFn(ctl)
		followers := streams.NewActivityStreamsCollection()
		expectFollowers := streams.NewActivityStreamsCollection()
		expectItems := streams.NewActivityStreamsItemsProperty()
		expectItems.AppendIRI(mustParse(testFederatedActorIRI))
		expectFollowers.SetActivityStreamsItems(expectItems)
		mockDB.EXPECT().Lock(ctx, mustParse(testMyInboxIRI))
		mockDB.EXPECT().ActorForInbox(ctx, mustParse(testMyInboxIRI)).Return(
			mustParse(testFederatedActorIRI2), nil)
		mockDB.EXPECT().Unlock(ctx, mustParse(testMyInboxIRI))
		mockTp.EXPECT().Dereference(ctx, mustParse(testFederatedActivityIRI)).Return(
			mustSerializeToBytes(testFollow), nil)
		mockDB.EXPECT().Lock(ctx, mustParse(testFederatedActivityIRI))
		mockDB.EXPECT().Get(ctx, mustParse(testFederatedActivityIRI)).Return(
			testFollow, nil)
		mockDB.EXPECT().Unlock(ctx, mustParse(testFederatedActivityIRI))
		mockDB.EXPECT().Lock(ctx, mustParse(testFederatedActorIRI2))
		mockDB.EXPECT().Following(ctx, mustParse(testFederatedActorIRI2)).Return(
			followers, nil)
		mockDB.EXPECT().Update(ctx, expectFollowers)
		mockDB.EXPECT().Unlock(ctx, mustParse(testFederatedActorIRI2))
		a := newAcceptFn()
		op := streams.NewActivityStreamsObjectProperty()
		op.AppendIRI(mustParse(testFederatedActivityIRI))
		a.SetActivityStreamsObject(op)
		err := w.accept(ctx, a)
		if err != nil {
			t.Fatalf("got error %s", err)
		}
	})
	t.Run("IgnoresNonFollowObjects", func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		w, mockDB, _ := setupFn(ctl)
		mockDB.EXPECT().Lock(ctx, mustParse(testMyInboxIRI))
		mockDB.EXPECT().ActorForInbox(ctx, mustParse(testMyInboxIRI)).Return(
			mustParse(testFederatedActorIRI2), nil)
		mockDB.EXPECT().Unlock(ctx, mustParse(testMyInboxIRI))
		a := newAcceptFn()
		op := streams.NewActivityStreamsObjectProperty()
		op.AppendActivityStreamsListen(testListen)
		a.SetActivityStreamsObject(op)
		err := w.accept(ctx, a)
		if err != nil {
			t.Fatalf("got error %s", err)
		}
	})
	t.Run("IgnoresFollowObjectsNotContainingMe", func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		w, mockDB, _ := setupFn(ctl)
		mockDB.EXPECT().Lock(ctx, mustParse(testMyInboxIRI))
		mockDB.EXPECT().ActorForInbox(ctx, mustParse(testMyInboxIRI)).Return(
			mustParse(testFederatedActorIRI3), nil)
		mockDB.EXPECT().Unlock(ctx, mustParse(testMyInboxIRI))
		a := newAcceptFn()
		err := w.accept(ctx, a)
		if err != nil {
			t.Fatalf("got error %s", err)
		}
	})
	t.Run("ErrorIfPeerLiedAboutOurFollowId", func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		w, mockDB, _ := setupFn(ctl)
		mockDB.EXPECT().Lock(ctx, mustParse(testMyInboxIRI))
		mockDB.EXPECT().ActorForInbox(ctx, mustParse(testMyInboxIRI)).Return(
			mustParse(testFederatedActorIRI2), nil)
		mockDB.EXPECT().Unlock(ctx, mustParse(testMyInboxIRI))
		mockDB.EXPECT().Lock(ctx, mustParse(testFederatedActivityIRI))
		mockDB.EXPECT().Get(ctx, mustParse(testFederatedActivityIRI)).Return(
			testListen, nil)
		mockDB.EXPECT().Unlock(ctx, mustParse(testFederatedActivityIRI))
		a := newAcceptFn()
		err := w.accept(ctx, a)
		if err == nil {
			t.Fatalf("expected error, got none")
		}
	})
	t.Run("UpdatesFollowingCollection", func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		w, mockDB, _ := setupFn(ctl)
		followers := streams.NewActivityStreamsCollection()
		expectFollowers := streams.NewActivityStreamsCollection()
		expectItems := streams.NewActivityStreamsItemsProperty()
		expectItems.AppendIRI(mustParse(testFederatedActorIRI))
		expectFollowers.SetActivityStreamsItems(expectItems)
		mockDB.EXPECT().Lock(ctx, mustParse(testMyInboxIRI))
		mockDB.EXPECT().ActorForInbox(ctx, mustParse(testMyInboxIRI)).Return(
			mustParse(testFederatedActorIRI2), nil)
		mockDB.EXPECT().Unlock(ctx, mustParse(testMyInboxIRI))
		mockDB.EXPECT().Lock(ctx, mustParse(testFederatedActivityIRI))
		mockDB.EXPECT().Get(ctx, mustParse(testFederatedActivityIRI)).Return(
			testFollow, nil)
		mockDB.EXPECT().Unlock(ctx, mustParse(testFederatedActivityIRI))
		mockDB.EXPECT().Lock(ctx, mustParse(testFederatedActorIRI2))
		mockDB.EXPECT().Following(ctx, mustParse(testFederatedActorIRI2)).Return(
			followers, nil)
		mockDB.EXPECT().Update(ctx, expectFollowers)
		mockDB.EXPECT().Unlock(ctx, mustParse(testFederatedActorIRI2))
		a := newAcceptFn()
		err := w.accept(ctx, a)
		if err != nil {
			t.Fatalf("got error %s", err)
		}
	})
	t.Run("CallsCustomCallback", func(t *testing.T) {
		a := newAcceptFn()
		a.SetActivityStreamsObject(nil)
		var w FederatingWrappedCallbacks
		var gotc context.Context
		var got vocab.ActivityStreamsAccept
		w.Accept = func(ctx context.Context, v vocab.ActivityStreamsAccept) error {
			gotc = ctx
			got = v
			return nil
		}
		err := w.accept(ctx, a)
		if err != nil {
			t.Fatalf("got error %s", err)
		}
		assertEqual(t, ctx, gotc)
		assertEqual(t, a, got)
	})
}

func TestFederatedReject(t *testing.T) {
	ctx := context.Background()
	t.Run("CallsCustomCallback", func(t *testing.T) {
		r := streams.NewActivityStreamsReject()
		var w FederatingWrappedCallbacks
		var gotc context.Context
		var got vocab.ActivityStreamsReject
		w.Reject = func(ctx context.Context, v vocab.ActivityStreamsReject) error {
			gotc = ctx
			got = v
			return nil
		}
		err := w.reject(ctx, r)
		if err != nil {
			t.Fatalf("got error %s", err)
		}
		assertEqual(t, ctx, gotc)
		assertEqual(t, r, got)
	})
}

func TestFederatedAdd(t *testing.T) {
	t.Run("ErrorIfNoObject", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("ErrorIfObjectLengthZero", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("ErrorIfNoTarget", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("ErrorIfTargetLengthZero", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("AddsAllObjectIdsToCollection", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("AddsAllObjectIdsToOrderedCollection", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("ReturnsErrorIfTargetIsNotCollection", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("CallsCustomCallback", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
}

func TestFederatedRemove(t *testing.T) {
	t.Run("ErrorIfNoObject", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("ErrorIfObjectLengthZero", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("ErrorIfNoTarget", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("ErrorIfTargetLengthZero", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("RemovesAllObjectIdsFromCollection", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("RemovesAllObjectIdsFromOrderedCollection", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("ReturnsErrorIfTargetIsNotCollection", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("CallsCustomCallback", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
}

func TestFederatedLike(t *testing.T) {
	t.Run("ErrorIfNoObject", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("ErrorIfObjectLengthZero", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("SkipsUnownedObjects", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("AddsToNewLikesCollection", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("AddsToExistingLikesCollection", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("AddsToExistingLikesOrderedCollection", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("CallsCustomCallback", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
}

func TestFederatedAnnounce(t *testing.T) {
	t.Run("SkipsUnownedObjects", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("AddsToNewSharesCollection", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("AddsToExistingSharesCollection", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("AddsToExistingSharesOrderedCollection", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("CallsCustomCallback", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
}

func TestFederatedUndo(t *testing.T) {
	t.Run("ErrorIfNoObject", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("ErrorIfObjectLengthZero", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("ErrorIfActorMismatch", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("CallsCustomCallback", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
}

func TestFederatedBlock(t *testing.T) {
	t.Run("ErrorIfNoObject", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("ErrorIfObjectLengthZero", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("CallsCustomCallback", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
}

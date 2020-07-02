package pub

import (
	"context"
	"net/url"
	"testing"

	"github.com/go-fed/activity/streams/vocab"
	"github.com/go-fed/activity/streams"
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
		w.Create =func(ctx context.Context, v vocab.ActivityStreamsCreate) error {
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
	t.Run("ErrorIfNoObject", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("ErrorIfObjectLengthZero", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("ErrorIfOriginMismatchesObject", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("UpdatesFederatedObject", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("UpdatesAllFederatedObjects", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("ErrorIfObjectIsIRI", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("CallsCustomCallback", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
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
	t.Run("DoesNothingIfNoObjects", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("DereferencesObjectIRI", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("IgnoresNonFollowObjects", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("IgnoresFollowObjectsNotContainingMe", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("VerifiesFollowExistsAndIsWellFormatted", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("UpdatesFollowingCollection", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("CallsCustomCallback", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
}

func TestFederatedReject(t *testing.T) {
	t.Run("CallsCustomCallback", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
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

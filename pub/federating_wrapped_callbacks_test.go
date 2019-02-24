package pub

import (
	"testing"
)

// TestFederatedCallbacks tests the overriding functionality.
func TestFederatedCallbacks(t *testing.T) {
	t.Run("ReturnsOtherCallback", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("OverridesCreate", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("OverridesUpdate", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("OverridesDelete", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("OverridesFollow", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("OverridesAccept", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("OverridesReject", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("OverridesAdd", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("OverridesRemove", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("OverridesLike", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("OverridesAnnounce", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("OverridesUndo", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("OverridesBlock", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
}

func TestFederatedCreate(t *testing.T) {
	t.Run("ErrorIfNoObject", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("ErrorIfObjectLengthZero", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("CreatesFederatedObject", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("CreatesAllFederatedObjects", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("DereferencesIRIObject", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
	})
	t.Run("CallsCustomCallback", func(t *testing.T) {
		t.Errorf("Not yet implemented.")
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

package parser

import (
	"sort"
	"testing"

	"github.com/Kallrish/quakeLogParser/shared"
)

func TestCheckIfStringSliceNeedSort(t *testing.T) {
	slice1 := []string{shared.TEST_NAME_JOHN, shared.TEST_NAME_ARTHUR, shared.TEST_NAME_MARY}
	isListSortedBefore := sort.StringsAreSorted(slice1)

	checkIfStringSliceNeedSort(slice1)

	isListSortedAfter := sort.StringsAreSorted(slice1)

	if isListSortedBefore == isListSortedAfter {
		t.Errorf(shared.TEST_EXPECTED_SORTED_SLICE)
	}
}

func TestGetGameListKey(t *testing.T) {
	if numberResult := getGameListKey(shared.ONE); numberResult != shared.ZERO {
		t.Errorf(shared.TEST_EXPECTED_GOT_INT, shared.ZERO, numberResult)
	}
}

func TestHasPlayerInSlice(t *testing.T) {
	emptyPlayersList := make([]string, shared.ZERO)
	oneNamePlayerList := []string{shared.TEST_NAME_MARY}
	threeNamePlayerList := []string{shared.TEST_NAME_JOHN, shared.TEST_NAME_ARTHUR, shared.TEST_NAME_MARY}

	if hasPlayerInSlice(shared.TEST_NAME_JOHN, emptyPlayersList) != false {
		t.Errorf(shared.TEXT_EXPECTED_EMPTY_SLICE_BUT, shared.TEST_NAME_JOHN)
	}
	if hasPlayerInSlice(shared.TEST_NAME_MARY, oneNamePlayerList) != true {
		t.Errorf(shared.TEST_EXPECTED_PLAYER, shared.TEST_NAME_MARY)
	}
	if hasPlayerInSlice(shared.TEST_NAME_ARTHUR, threeNamePlayerList) != true {
		t.Errorf(shared.TEST_EXPECTED_PLAYER, shared.TEST_NAME_ARTHUR)
	}
}

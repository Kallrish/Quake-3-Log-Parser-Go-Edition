package parser

import (
	"sort"

	"github.com/Kallrish/quakeLogParser/shared"
)

func checkIfStringSliceNeedSort(list []string) {
	isListSorted := sort.StringsAreSorted(list)
	if isListSorted {
		return
	} else {
		sort.Strings(list)
	}
}

func getGameListKey(gameNumber int) int {
	return gameNumber - shared.ONE
}

func hasPlayerInSlice(player string, list []string) bool {
	for _, value := range list {
		if value == player {
			return true
		}
	}
	return false
}

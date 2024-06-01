package parser

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/Kallrish/quakeLogParser/model"
	"github.com/Kallrish/quakeLogParser/shared"
	"github.com/Kallrish/quakeLogParser/util"
)

func ParseLogFile(logFile string) {

	gameNumberCount := shared.ZERO
	gameList := make([]model.GameInfo, shared.ZERO)
	logLines := strings.Split(logFile, shared.BREAK_LINE)

	for _, line := range logLines {
		trimmedLine := strings.TrimSpace(line)
		infoList := strings.Split(trimmedLine, shared.SPACE)

		if len(infoList) > shared.TWO {
			switch infoList[shared.ONE] {
			case shared.INIT_GAME:
				gameNumberCount++
				gameList = append(gameList, createNewGame())
			case shared.CLIENT_INFO:
				addPlayer(gameList, gameNumberCount, infoList)
			case shared.KILL:
				updateTotalKillInfo(gameList, gameNumberCount)
				addKillDeathInfo(gameList, gameNumberCount, infoList)
			}
		}
	}
	generateFinalGameListMap(gameList, shared.FinalGameList)
}

func createNewGame() model.GameInfo {
	game := model.GameInfo{
		TotalKills: shared.ZERO,
		Players:    make([]string, shared.ZERO),
		Kills:      make(map[string]int),
		Deaths:     make(map[string]int),
	}
	game.Deaths = generateDeathTypesMap()
	return game
}

func generateDeathTypesMap() map[string]int {
	return map[string]int{
		"MOD_UNKNOWN":        shared.ZERO,
		"MOD_SHOTGUN":        shared.ZERO,
		"MOD_GAUNTLET":       shared.ZERO,
		"MOD_MACHINEGUN":     shared.ZERO,
		"MOD_GRENADE":        shared.ZERO,
		"MOD_GRENADE_SPLASH": shared.ZERO,
		"MOD_ROCKET":         shared.ZERO,
		"MOD_ROCKET_SPLASH":  shared.ZERO,
		"MOD_PLASMA":         shared.ZERO,
		"MOD_PLASMA_SPLASH":  shared.ZERO,
		"MOD_RAILGUN":        shared.ZERO,
		"MOD_LIGHTNING":      shared.ZERO,
		"MOD_BFG":            shared.ZERO,
		"MOD_BFG_SPLASH":     shared.ZERO,
		"MOD_WATER":          shared.ZERO,
		"MOD_SLIME":          shared.ZERO,
		"MOD_LAVA":           shared.ZERO,
		"MOD_CRUSH":          shared.ZERO,
		"MOD_TELEFRAG":       shared.ZERO,
		"MOD_FALLING":        shared.ZERO,
		"MOD_SUICIDE":        shared.ZERO,
		"MOD_TARGET_LASER":   shared.ZERO,
		"MOD_TRIGGER_HURT":   shared.ZERO,
		"MOD_NAIL":           shared.ZERO,
		"MOD_CHAINGUN":       shared.ZERO,
		"MOD_PROXIMITY_MINE": shared.ZERO,
		"MOD_KAMIKAZE":       shared.ZERO,
		"MOD_JUICED":         shared.ZERO,
		"MOD_GRAPPL":         shared.ZERO,
	}
}

func addPlayer(list []model.GameInfo, gameNumber int, infoList []string) {
	key := getGameListKey(gameNumber)
	regex := util.ReturnRegex(shared.PLAYER_CONECT_INFO)
	player := regex.FindString(strings.Join(infoList[shared.THREE:], shared.SPACE))

	if len(player) > 1 {
		hasPlayer := hasPlayerInSlice(player, list[key].Players)
		if hasPlayer {
			return
		} else {
			list[key].Players = append(list[key].Players, player)
			checkIfStringSliceNeedSort(list, key)
		}
	} else {
		fmt.Println(shared.NO_PLAYER_FOUND)
	}
}

func checkIfStringSliceNeedSort(list []model.GameInfo, listKey int) {
	isListSorted := sort.StringsAreSorted(list[listKey].Players)
	if isListSorted {
		return
	} else {
		sort.Strings(list[listKey].Players)
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

func updateTotalKillInfo(list []model.GameInfo, gameNumber int) {
	key := getGameListKey(gameNumber)
	list[key].TotalKills++
}

func addKillDeathInfo(list []model.GameInfo, gameNumber int, infoList []string) {
	key := getGameListKey(gameNumber)
	killer, killed, death := getKillerKilledDeathInfo(infoList)
	if killer == shared.WORLD {
		list[key].Kills[killed]--
	} else if killer == killed {
		return
	} else {
		list[key].Kills[killer]++
	}
	list[key].Deaths[death]++
}

func getKillerKilledDeathInfo(infoList []string) (string, string, string) {
	var killer, killed, death string
	listIndex := shared.FIVE
	for infoList[listIndex] != shared.KILLED {
		if infoList[listIndex+shared.ONE] == shared.KILLED {
			killer += infoList[listIndex]
		} else {
			killer += infoList[listIndex] + shared.SPACE
		}
		listIndex++
	}
	listIndex++
	for infoList[listIndex] != shared.BY {
		if infoList[listIndex+shared.ONE] == shared.BY {
			killed += infoList[listIndex]
		} else {
			killed += infoList[listIndex] + shared.SPACE
		}
		listIndex++
	}
	listIndex++
	death = infoList[listIndex]
	return killer, killed, death
}

func generateFinalGameListMap(list []model.GameInfo, finalList map[string]model.GameInfo) {
	for key, value := range list {
		var gameName string
		if key < 9 {
			gameName = shared.GAME_NUMBER_PREFIX + shared.ZERO_SINGLE + strconv.Itoa(key+shared.ONE)
		} else {
			gameName = shared.GAME_NUMBER_PREFIX + strconv.Itoa(key+shared.ONE)
		}
		finalList[gameName] = value
	}
}

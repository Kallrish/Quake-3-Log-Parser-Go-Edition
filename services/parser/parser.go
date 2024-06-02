package parser

import (
	"fmt"
	"strings"

	"github.com/Kallrish/quakeLogParser/model"
	"github.com/Kallrish/quakeLogParser/shared"
	"github.com/Kallrish/quakeLogParser/util"
)

func ParseLogFile(logFile string) map[string]model.GameInfo {

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
				gameList = append(gameList, generateEmptyGameInfo())
			case shared.CLIENT_INFO:
				addPlayerToPlayerList(gameList, gameNumberCount, infoList)
			case shared.KILL:
				updateTotalKillGameInfo(gameList, gameNumberCount)
				addKillDeathEventData(gameList, gameNumberCount, infoList)
			}
		}
	}
	return generateFinalGameListMap(gameList)
}

func addPlayerToPlayerList(list []model.GameInfo, gameNumber int, infoList []string) {
	key := getGameListKey(gameNumber)
	regex := util.ReturnRegex(shared.PLAYER_CONECT_INFO)
	player := regex.FindString(strings.Join(infoList[shared.THREE:], shared.SPACE))

	if len(player) > 1 {
		hasPlayer := hasPlayerInSlice(player, list[key].Players)
		if hasPlayer {
			return
		} else {
			list[key].Players = append(list[key].Players, player)
			checkIfStringSliceNeedSort(list[key].Players)
		}
	} else {
		fmt.Println(shared.NO_PLAYER_FOUND)
	}
}

func updateTotalKillGameInfo(list []model.GameInfo, gameNumber int) {
	key := getGameListKey(gameNumber)
	list[key].TotalKills++
}

func addKillDeathEventData(list []model.GameInfo, gameNumber int, infoList []string) {
	key := getGameListKey(gameNumber)
	killer, killed, death := getKillerKilledDeathEventData(infoList)
	if killer == shared.WORLD {
		list[key].Kills[killed]--
	} else if killer == killed {
		return
	} else {
		list[key].Kills[killer]++
	}
	list[key].Deaths[death]++
}

func getKillerKilledDeathEventData(infoList []string) (string, string, string) {
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

package parser

import (
	"strconv"

	"github.com/Kallrish/quakeLogParser/model"
	"github.com/Kallrish/quakeLogParser/shared"
)

func generateEmptyGameInfo() model.GameInfo {
	game := model.GameInfo{
		TotalKills: shared.ZERO,
		Players:    make([]string, shared.ZERO),
		Kills:      make(map[string]int),
		Deaths:     make(map[string]int),
	}
	game.Deaths = generateEmptyDeathTypesMap()
	return game
}

func generateEmptyDeathTypesMap() map[string]int {
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

func generateFinalGameListMap(list []model.GameInfo) map[string]model.GameInfo {
	gameList := make(map[string]model.GameInfo)
	for key, value := range list {
		var gameName string
		if key < shared.NINE {
			gameName = shared.GAME_NUMBER_PREFIX + shared.ZERO_SINGLE + strconv.Itoa(key+shared.ONE)
		} else {
			gameName = shared.GAME_NUMBER_PREFIX + strconv.Itoa(key+shared.ONE)
		}
		gameList[gameName] = value
	}
	return gameList
}

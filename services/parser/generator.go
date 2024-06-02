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
		shared.MOD_UNKNOWN:        shared.ZERO,
		shared.MOD_SHOTGUN:        shared.ZERO,
		shared.MOD_GAUNTLET:       shared.ZERO,
		shared.MOD_MACHINEGUN:     shared.ZERO,
		shared.MOD_GRENADE:        shared.ZERO,
		shared.MOD_GRENADE_SPLASH: shared.ZERO,
		shared.MOD_ROCKET:         shared.ZERO,
		shared.MOD_ROCKET_SPLASH:  shared.ZERO,
		shared.MOD_PLASMA:         shared.ZERO,
		shared.MOD_PLASMA_SPLASH:  shared.ZERO,
		shared.MOD_RAILGUN:        shared.ZERO,
		shared.MOD_LIGHTNING:      shared.ZERO,
		shared.MOD_BFG:            shared.ZERO,
		shared.MOD_BFG_SPLASH:     shared.ZERO,
		shared.MOD_WATER:          shared.ZERO,
		shared.MOD_SLIME:          shared.ZERO,
		shared.MOD_LAVA:           shared.ZERO,
		shared.MOD_CRUSH:          shared.ZERO,
		shared.MOD_TELEFRAG:       shared.ZERO,
		shared.MOD_FALLING:        shared.ZERO,
		shared.MOD_SUICIDE:        shared.ZERO,
		shared.MOD_TARGET_LASER:   shared.ZERO,
		shared.MOD_TRIGGER_HURT:   shared.ZERO,
		shared.MOD_NAIL:           shared.ZERO,
		shared.MOD_CHAINGUN:       shared.ZERO,
		shared.MOD_PROXIMITY_MINE: shared.ZERO,
		shared.MOD_KAMIKAZE:       shared.ZERO,
		shared.MOD_JUICED:         shared.ZERO,
		shared.MOD_GRAPPL:         shared.ZERO,
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

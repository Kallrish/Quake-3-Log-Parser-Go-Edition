package parser

import (
	"testing"

	"github.com/Kallrish/quakeLogParser/model"
	"github.com/Kallrish/quakeLogParser/shared"
)

var list = []string{
	shared.HYPHEN,
	shared.HYPHEN,
	shared.HYPHEN,
	shared.HYPHEN,
	shared.HYPHEN,
	shared.WORLD,
	shared.KILLED,
	shared.TEST_NAME_JOHN,
	shared.BY,
	shared.MOD_TRIGGER_HURT,
}
var list2 = []string{
	shared.HYPHEN,
	shared.HYPHEN,
	shared.HYPHEN,
	shared.HYPHEN,
	shared.HYPHEN,
	shared.TEST_NAME_MARY,
	shared.KILLED,
	shared.TEST_NAME_JOHN,
	shared.BY,
	shared.MOD_SHOTGUN,
}
var list3 = []string{
	shared.HYPHEN,
	shared.HYPHEN,
	shared.HYPHEN,
	shared.HYPHEN,
	shared.HYPHEN,
	shared.TEST_NAME_JOHN,
	shared.KILLED,
	shared.TEST_NAME_MARY,
	shared.BY,
	shared.MOD_SHOTGUN,
}
var game1 = generateEmptyGameInfo()
var game2 = generateEmptyGameInfo()
var gameNumber = 1
var gameKey = getGameListKey(gameNumber)

func TestGetKillerKilledDeathEventData(t *testing.T) {
	tKiller, tKilled, tDeath := getKillerKilledDeathEventData(list)

	if tKiller != shared.WORLD {
		t.Errorf(shared.TEST_EXPECTED_GOT_STRING, shared.WORLD, tKiller)
	}
	if tKilled != shared.TEST_NAME_JOHN {
		t.Errorf(shared.TEST_EXPECTED_GOT_STRING, shared.TEST_NAME_JOHN, tKilled)
	}
	if tDeath != shared.MOD_TRIGGER_HURT {
		t.Errorf(shared.TEST_EXPECTED_GOT_STRING, shared.MOD_TRIGGER_HURT, tDeath)
	}
}

func TestUpdateTotalKillGameInfo(t *testing.T) {
	gameList := []model.GameInfo{game1, game2}

	updateTotalKillGameInfo(gameList, gameNumber)

	totalKills := gameList[gameKey].TotalKills

	if totalKills != shared.ONE {
		t.Errorf(shared.TEST_EXPECTED_GOT_INT, shared.ONE, totalKills)
	}

}

func TestAddKillDeathEventData(t *testing.T) {
	gameList := []model.GameInfo{game1, game2}

	addKillDeathEventData(gameList, gameNumber, list2)

	maryKills := gameList[gameKey].Kills[shared.TEST_NAME_MARY]
	deathCount := gameList[gameKey].Deaths[shared.MOD_SHOTGUN]

	if maryKills != shared.ONE {
		t.Errorf(shared.TEST_EXPECTED_GOT_INT, shared.ONE, maryKills)
	}
	if deathCount != shared.ONE {
		t.Errorf(shared.TEST_EXPECTED_GOT_INT, shared.ONE, deathCount)
	}

}

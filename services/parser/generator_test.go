package parser

import (
	"testing"

	"github.com/Kallrish/quakeLogParser/model"
	"github.com/Kallrish/quakeLogParser/shared"
)

var emptyDeathTypeMap = map[string]int{
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

func TestGenerateEmptyGameInfo(t *testing.T) {
	correctGame := model.GameInfo{
		TotalKills: shared.ZERO,
		Players:    make([]string, shared.ZERO),
		Kills:      make(map[string]int),
		Deaths:     emptyDeathTypeMap,
	}
	testGame := generateEmptyGameInfo()

	if testGame.TotalKills != correctGame.TotalKills {
		t.Errorf(shared.EXPECTED_TOTAL_KILLS, correctGame.TotalKills, testGame.TotalKills)
	}

	if len(testGame.Players) != len(correctGame.Players) {
		t.Errorf(shared.EXPECTED_PLAYERS, len(correctGame.Players), len(testGame.Players))
	}
	if len(testGame.Kills) != len(correctGame.Kills) {
		t.Errorf(shared.EXPECTED_KILLS, len(correctGame.Kills), len(testGame.Kills))
	}
	if len(testGame.Deaths) != len(correctGame.Deaths) {
		t.Errorf(shared.EXPECTED_DEATHS, len(correctGame.Deaths), len(testGame.Deaths))
	}
}

func TestGenerateEmptyDeathTypesMap(t *testing.T) {
	deathTestMap := generateEmptyDeathTypesMap()
	deathTestMapLen := len(deathTestMap)
	correctSize := len(emptyDeathTypeMap)

	if deathTestMapLen != len(emptyDeathTypeMap) {
		t.Errorf(shared.EXPECTED_MAP_LENGTH, correctSize, deathTestMapLen)
	}

	zeroCount := shared.ZERO
	for _, value := range deathTestMap {
		if value != shared.ZERO {
			zeroCount++
		}
	}
	if zeroCount != shared.ZERO {
		t.Errorf(shared.EXPECTED_MAP_POSITIONS, correctSize, shared.ZERO, zeroCount)
	}
}

func TestGenerateFinalGameListMap(t *testing.T) {
	testDeathTypeMap1 := map[string]int{
		"MOD_UNKNOWN":        shared.ZERO,
		"MOD_SHOTGUN":        shared.ZERO,
		"MOD_GAUNTLET":       shared.ZERO,
		"MOD_MACHINEGUN":     shared.ZERO,
		"MOD_GRENADE":        shared.ZERO,
		"MOD_GRENADE_SPLASH": shared.ONE,
		"MOD_ROCKET":         shared.ZERO,
		"MOD_ROCKET_SPLASH":  shared.ZERO,
		"MOD_PLASMA":         shared.ZERO,
		"MOD_PLASMA_SPLASH":  shared.ZERO,
		"MOD_RAILGUN":        shared.ZERO,
		"MOD_LIGHTNING":      shared.ZERO,
		"MOD_BFG":            shared.FOUR,
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

	testDeathTypeMap2 := map[string]int{
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
		"MOD_BFG":            shared.FOUR,
		"MOD_BFG_SPLASH":     shared.ZERO,
		"MOD_WATER":          shared.ZERO,
		"MOD_SLIME":          shared.ZERO,
		"MOD_LAVA":           shared.ZERO,
		"MOD_CRUSH":          shared.ZERO,
		"MOD_TELEFRAG":       shared.TWO,
		"MOD_FALLING":        shared.ZERO,
		"MOD_SUICIDE":        shared.ZERO,
		"MOD_TARGET_LASER":   shared.ZERO,
		"MOD_TRIGGER_HURT":   shared.ZERO,
		"MOD_NAIL":           shared.ZERO,
		"MOD_CHAINGUN":       shared.ZERO,
		"MOD_PROXIMITY_MINE": shared.ZERO,
		"MOD_KAMIKAZE":       shared.ZERO,
		"MOD_JUICED":         shared.ZERO,
		"MOD_GRAPPL":         shared.THREE,
	}

	testKillMap1 := map[string]int{
		shared.TEST_NAME_JOHN: shared.ONE,
		shared.TEST_NAME_MARY: shared.FOUR,
	}

	testKillMap2 := map[string]int{
		shared.TEST_NAME_ARTHUR: shared.FOUR,
		shared.TEST_NAME_JOHN:   shared.TWO,
		shared.TEST_NAME_MARY:   shared.THREE,
	}

	testPlayerSlice1 := []string{
		shared.TEST_NAME_JOHN,
		shared.TEST_NAME_MARY,
	}

	testPlayerSlice2 := []string{
		shared.TEST_NAME_ARTHUR,
		shared.TEST_NAME_JOHN,
		shared.TEST_NAME_MARY,
	}

	testTotalKills1 := shared.FIVE
	testTotalKills2 := shared.NINE

	testGameKey1 := shared.TEST_GAME_ONE
	testGameKey2 := shared.TEST_GAME_TEN

	testgGameInfo1 := model.GameInfo{
		TotalKills: testTotalKills1,
		Players:    testPlayerSlice1,
		Kills:      testKillMap1,
		Deaths:     testDeathTypeMap1,
	}
	testgGameInfo2 := model.GameInfo{
		TotalKills: testTotalKills2,
		Players:    testPlayerSlice2,
		Kills:      testKillMap2,
		Deaths:     testDeathTypeMap2,
	}
	testgGameInfo3 := model.GameInfo{
		TotalKills: testTotalKills1,
		Players:    testPlayerSlice1,
		Kills:      testKillMap1,
		Deaths:     testDeathTypeMap1,
	}
	testGameSlice := []model.GameInfo{
		testgGameInfo1,
		testgGameInfo2,
		testgGameInfo3,
		testgGameInfo3,
		testgGameInfo3,
		testgGameInfo3,
		testgGameInfo3,
		testgGameInfo3,
		testgGameInfo3,
		testgGameInfo3,
	}
	testGameList := generateFinalGameListMap(testGameSlice)

	_, ok := testGameList[testGameKey1]
	_, ok2 := testGameList[testGameKey2]
	if ok != true {
		t.Errorf(shared.EXPECTED_MAP_KEY_1, testGameKey1, shared.ZERO)
	}
	if ok2 != true {
		t.Errorf(shared.EXPECTED_MAP_KEY_2, testGameKey2, shared.NINE)
	}

}

package parser

import (
	"encoding/json"
	"testing"

	"github.com/Kallrish/quakeLogParser/model"
	"github.com/Kallrish/quakeLogParser/shared"
)

func TestGenerateJsonString(t *testing.T) {
	testDeathTypeMap1 := map[string]int{
		shared.MOD_UNKNOWN:        shared.ZERO,
		shared.MOD_SHOTGUN:        shared.ZERO,
		shared.MOD_GAUNTLET:       shared.ZERO,
		shared.MOD_MACHINEGUN:     shared.ZERO,
		shared.MOD_GRENADE:        shared.ZERO,
		shared.MOD_GRENADE_SPLASH: shared.ONE,
		shared.MOD_ROCKET:         shared.ZERO,
		shared.MOD_ROCKET_SPLASH:  shared.ZERO,
		shared.MOD_PLASMA:         shared.ZERO,
		shared.MOD_PLASMA_SPLASH:  shared.ZERO,
		shared.MOD_RAILGUN:        shared.ZERO,
		shared.MOD_LIGHTNING:      shared.ZERO,
		shared.MOD_BFG:            shared.FOUR,
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

	testDeathTypeMap2 := map[string]int{
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
		shared.MOD_BFG:            shared.FOUR,
		shared.MOD_BFG_SPLASH:     shared.ZERO,
		shared.MOD_WATER:          shared.ZERO,
		shared.MOD_SLIME:          shared.ZERO,
		shared.MOD_LAVA:           shared.ZERO,
		shared.MOD_CRUSH:          shared.ZERO,
		shared.MOD_TELEFRAG:       shared.TWO,
		shared.MOD_FALLING:        shared.ZERO,
		shared.MOD_SUICIDE:        shared.ZERO,
		shared.MOD_TARGET_LASER:   shared.ZERO,
		shared.MOD_TRIGGER_HURT:   shared.ZERO,
		shared.MOD_NAIL:           shared.ZERO,
		shared.MOD_CHAINGUN:       shared.ZERO,
		shared.MOD_PROXIMITY_MINE: shared.ZERO,
		shared.MOD_KAMIKAZE:       shared.ZERO,
		shared.MOD_JUICED:         shared.ZERO,
		shared.MOD_GRAPPL:         shared.THREE,
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

	testMap := GenerateJsonString(testGameList)

	testMapByte := []byte(testMap)

	var resultData map[string]interface{}
	marshalError := json.Unmarshal(testMapByte, &resultData)
	if marshalError != nil {
		t.Errorf(shared.TEST_FAILED_UNMARSHAL, marshalError)
	}
	if len(resultData) != len(testGameList) {
		t.Errorf(shared.TEST_UNEXPECTED_NUMBER, len(resultData), len(testGameList))
	}

}

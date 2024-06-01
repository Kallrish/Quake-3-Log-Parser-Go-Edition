package parser

import (
	"encoding/json"
	"fmt"

	"github.com/Kallrish/quakeLogParser/model"
	"github.com/Kallrish/quakeLogParser/shared"
)

func GenerateJsonString(finalGameList map[string]model.GameInfo) string {
	jsonFile, err := json.MarshalIndent(finalGameList, shared.EMPTY_TEXT, shared.SPACE)
	if err != nil {
		fmt.Println(err)
	}

	return string(jsonFile)
}

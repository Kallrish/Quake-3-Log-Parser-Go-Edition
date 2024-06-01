package report

import (
	"log"
	"os"

	"github.com/Kallrish/quakeLogParser/shared"
)

func ExportJsonFileToPath(logString string) {
	filePath := shared.LogFilePath + shared.LogReportName
	byteConvertedString := []byte(logString)

	if _, pathError := os.Stat(shared.LogFilePath); os.IsNotExist(pathError) {
		os.MkdirAll(shared.LogFilePath, shared.PATH_CREATION_PERMISSION)
	}

	if _, fileError := os.Stat(filePath); os.IsExist(fileError) {
		os.Remove(filePath)
	}

	err := os.WriteFile(filePath, byteConvertedString, shared.REPORT_FILE_PERMISSION)
	if err != nil {
		log.Fatal(err)
	}

}

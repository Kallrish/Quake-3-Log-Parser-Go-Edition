package report

import (
	"github.com/Kallrish/quakeLogParser/shared"
)

func ExportJsonFileToPath(logString string) {
	filePath := shared.LogFilePath + shared.LogReportName
	byteConvertedString := []byte(logString)

	checkIfPathAndDirectoryStillExists(shared.LogFilePath)
	checkIfFileExists(filePath)
	createJsonFileInTargetPath(filePath, byteConvertedString, shared.REPORT_FILE_PERMISSION)
}

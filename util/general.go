package util

import (
	"regexp"
	"strings"

	"github.com/Kallrish/quakeLogParser/shared"
)

func SetLogFileName(filePath string) {
	shared.LogFileName = extractLogNameFromFilePath(filePath)
}

func SetLogFilePath(filePath string) {
	shared.LogFilePath = extractLogPathFromFullFilePath(filePath, shared.LogFileName)
}

func SetLogReportName() {
	shared.LogReportName = createLogReportName(shared.LogFileName, shared.LOG_SUFFIX)
}

func extractLogNameFromFilePath(filePath string) string {
	splitList := strings.Split(filePath, shared.SLASH)
	return splitList[len(splitList)-1]
}

func extractLogPathFromFullFilePath(filePath string, fileName string) string {
	return strings.TrimSuffix(filePath, fileName)
}

func createLogReportName(filename string, logSuffix string) string {
	return strings.TrimSuffix(filename, logSuffix) + shared.LOG_REPORT_NAME
}

func ReturnRegex(regex string) *regexp.Regexp {
	return regexp.MustCompile(regex)
}

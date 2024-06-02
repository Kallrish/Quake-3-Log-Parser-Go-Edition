package util

import (
	"strings"
	"testing"

	"github.com/Kallrish/quakeLogParser/shared"
)

func TestExtractLogNameFromFilePath(t *testing.T) {
	testLogName := extractLogNameFromFilePath(shared.TEST_LOG_FULL_PATH)

	if testLogName != shared.TEST_LOG_NAME {
		t.Errorf(shared.TEST_EXPECTED_GOT_STRING, shared.TEST_LOG_NAME, testLogName)
	}
}

func TestExtractLogPathFromFullFilePath(t *testing.T) {
	testPath := extractLogPathFromFullFilePath(shared.TEST_LOG_FULL_PATH, shared.TEST_LOG_NAME)

	if testPath != shared.TEST_LOG_PATH {
		t.Errorf(shared.TEST_EXPECTED_GOT_STRING, shared.TEST_LOG_PATH, testPath)
	}
}

func TestCreateLogReportName(t *testing.T) {
	logName := shared.TEST_LOG_NAME
	logReport := strings.TrimSuffix(logName, shared.LOG_SUFFIX) + shared.LOG_REPORT_NAME
	testReport := createLogReportName(shared.TEST_LOG_NAME, shared.LOG_SUFFIX)

	if testReport != logReport {
		t.Errorf(shared.TEST_EXPECTED_GOT_STRING, logReport, testReport)
	}
}

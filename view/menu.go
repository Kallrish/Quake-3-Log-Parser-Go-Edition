package view

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Kallrish/quakeLogParser/services/parser"
	"github.com/Kallrish/quakeLogParser/services/report"
	"github.com/Kallrish/quakeLogParser/shared"
	"github.com/Kallrish/quakeLogParser/util"
)

func PrintMainMenuWelcomeText() {
	fmt.Println()
	fmt.Println(shared.DIVIDER_DOUBLE)
	fmt.Println(shared.MENU_WELCOME)
	fmt.Println(shared.DIVIDER_DOUBLE)
	fmt.Println()
}

func GetUserLogFilePathInput() {
	var typedLogFilePath string
	fmt.Println(shared.DIVIDER_SINGLE)
	fmt.Println(shared.MENU_TYPE_LOG_PATH_AND_NAME)
	fmt.Println(shared.MENU_TYPE_LOG_PATH_AND_NAME_EXAMPLE)
	fmt.Println(shared.DIVIDER_SINGLE)
	fmt.Print(shared.PATH_TYPED)
	fmt.Scanln(&typedLogFilePath)
	trimmedTypedLogFilePath := strings.TrimSpace(typedLogFilePath)
	fmt.Println(trimmedTypedLogFilePath)

	userLog, err := os.ReadFile(trimmedTypedLogFilePath)
	if err != nil {
		fmt.Println(shared.DIVIDER_SINGLE)
		getUserRetryInput()
	}

	util.SetLogFileName(trimmedTypedLogFilePath)
	util.SetLogFilePath(trimmedTypedLogFilePath)
	util.SetLogReportName()
	showLogReportProccess(userLog)
}

func showLogReportProccess(userLog []byte) {
	fmt.Println(shared.DIVIDER_SINGLE)
	fmt.Println(shared.PARSING_LOG)
	finalGameList := parser.ParseLogFile(string(userLog))
	fmt.Println(shared.DIVIDER_SINGLE)
	fmt.Println(shared.CONVERTING_JSON)
	readyToExportGamesList := parser.GenerateJsonString(finalGameList)
	fmt.Println(shared.DIVIDER_SINGLE)
	fmt.Println(shared.EXPORTING_FILE)
	report.ExportJsonFileToPath(readyToExportGamesList)
	fmt.Println(shared.DIVIDER_SINGLE)
	fmt.Println(shared.EXPORT_FILE_SUCCESS + shared.LogFilePath + shared.LogReportName)
	fmt.Println(shared.DIVIDER_DOUBLE)
	fmt.Println(shared.THANKS)
	fmt.Println(shared.CLOSING_APPLICATION)
	fmt.Println(shared.DIVIDER_DOUBLE)
	fmt.Println()
	os.Exit(shared.ZERO)
}

func getUserRetryInput() {
	var userTypedOption string
	fmt.Println(shared.ERROR_LOCATE_FILE_PATH)
	fmt.Println(shared.DIVIDER_SINGLE)
	fmt.Println(shared.AVAILABLE_OPTIONS)
	fmt.Println(shared.TRY_TYPE_LOG_FILE_PATH_AGAIN)
	fmt.Println(shared.LEAVE_APPLICATION)
	fmt.Println(shared.TYPE_OPTION)
	fmt.Println(shared.DIVIDER_SINGLE)
	fmt.Print(shared.SELECTED_OPTION)
	fmt.Scanln(&userTypedOption)
	trimmedUserTypedOption := strings.TrimSpace(userTypedOption)
	showSelectedOption(trimmedUserTypedOption)
}

func showSelectedOption(selectedOption string) {
	if selectedOption == strconv.Itoa(shared.ONE) {
		fmt.Println(shared.DIVIDER_SINGLE)
		fmt.Println(shared.SELECTED_OPTION + shared.TRY_TYPE_LOG_FILE_PATH_AGAIN)
		GetUserLogFilePathInput()
	} else {
		fmt.Println(shared.DIVIDER_SINGLE)
		fmt.Println(shared.SELECTED_OPTION + shared.LEAVE_APPLICATION)
		fmt.Println(shared.DIVIDER_SINGLE)
		os.Exit(0)
	}
}

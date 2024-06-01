package shared

const (
	ZERO  = 0
	ONE   = 1
	TWO   = 2
	THREE = 3
	FOUR  = 4
	FIVE  = 5

	BACK_SLASH = "\\"
	BREAK_LINE = "\n"
	COLON      = ":"
	COMMA      = ","
	DOT        = "."
	EMPTY_TEXT = ""
	HYPHEN     = "-"
	PIPE       = "|"
	SEMI_COLON = ";"
	SLASH      = "/"
	SPACE      = " "
	UNDERSCORE = "_"

	BY                 = "by"
	CLIENT_INFO        = "ClientUserinfoChanged:"
	GAME_NUMBER_PREFIX = "game_"
	INIT_GAME          = "InitGame:"
	KILL               = "Kill:"
	KILLED             = "killed"
	WORLD              = "<world>"

	AVAILABLE_OPTIONS                   = "Available options:"
	CLOSING_APPLICATION                 = "Closing application..."
	CONVERTING_JSON                     = "Converting to JSON format..."
	DIVIDER_SINGLE                      = "---------------------------------------------------"
	DIVIDER_DOUBLE                      = "==================================================="
	EXPORTING_FILE                      = "Exporting file to path..."
	EXPORT_FILE_SUCCESS                 = "File successfuly exported to: "
	ERROR_LOCATE_FILE_PATH              = "Error. Could not locate a file with the typed path"
	LEAVE_APPLICATION                   = "Type any other -> Close application"
	LOG_REPORT_NAME                     = "LogParsedReport.json"
	LOG_SUFFIX                          = ".log"
	MENU_WELCOME                        = "Welcome to Quake 3 Log Parser - Go Edition"
	MENU_TYPE_LOG_PATH_AND_NAME         = "Type the full path to the log file and log name."
	MENU_TYPE_LOG_PATH_AND_NAME_EXAMPLE = "Example: C:/John/Downloads/quake.log"
	NO_PLAYER_FOUND                     = "No player was found!"
	PARSING_LOG                         = "Parsing log file..."
	PATH_TYPED                          = "Path typed: "
	SELECTED_OPTION                     = "Selected option: "
	THANKS                              = "Thank you for your preference!"
	TRY_TYPE_LOG_FILE_PATH_AGAIN        = "Type the number 1 -> Try type log file path again "
	TYPE_OPTION                         = "Type the desired option:"
	TYPE_OPTION_ERROR                   = "You need to type a number of the given options!"
	ZERO_SINGLE                         = "0"

	REPORT_FILE_PERMISSION   = 0644
	PATH_CREATION_PERMISSION = 0700

	PLAYER_CONECT_INFO = `[^\\n](\w*|\w* )*`
)

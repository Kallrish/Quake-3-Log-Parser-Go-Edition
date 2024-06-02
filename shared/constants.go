package shared

const (
	ZERO  = 0
	ONE   = 1
	TWO   = 2
	THREE = 3
	FOUR  = 4
	FIVE  = 5
	SIX   = 6
	SEVEN = 7
	EIGHT = 8
	NINE  = 9
	TEN   = 10

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

	MOD_UNKNOWN        = "MOD_UNKNOWN"
	MOD_SHOTGUN        = "MOD_SHOTGUN"
	MOD_GAUNTLET       = "MOD_GAUNTLET"
	MOD_MACHINEGUN     = "MOD_MACHINEGUN"
	MOD_GRENADE        = "MOD_GRENADE"
	MOD_GRENADE_SPLASH = "MOD_GRENADE_SPLASH"
	MOD_ROCKET         = "MOD_ROCKET"
	MOD_ROCKET_SPLASH  = "MOD_ROCKET_SPLASH"
	MOD_PLASMA         = "MOD_PLASMA"
	MOD_PLASMA_SPLASH  = "MOD_PLASMA_SPLASH"
	MOD_RAILGUN        = "MOD_RAILGUN"
	MOD_LIGHTNING      = "MOD_LIGHTNING"
	MOD_BFG            = "MOD_BFG"
	MOD_BFG_SPLASH     = "MOD_BFG_SPLASH"
	MOD_WATER          = "MOD_WATER"
	MOD_SLIME          = "MOD_SLIME"
	MOD_LAVA           = "MOD_LAVA"
	MOD_CRUSH          = "MOD_CRUSH"
	MOD_TELEFRAG       = "MOD_TELEFRAG"
	MOD_FALLING        = "MOD_FALLING"
	MOD_SUICIDE        = "MOD_SUICIDE"
	MOD_TARGET_LASER   = "MOD_TARGET_LASER"
	MOD_TRIGGER_HURT   = "MOD_TRIGGER_HURT"
	MOD_NAIL           = "MOD_NAIL"
	MOD_CHAINGUN       = "MOD_CHAINGUN"
	MOD_PROXIMITY_MINE = "MOD_PROXIMITY_MINE"
	MOD_KAMIKAZE       = "MOD_KAMIKAZE"
	MOD_JUICED         = "MOD_JUICED"
	MOD_GRAPPL         = "MOD_GRAPPL"

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

	TEST_NAME_JOHN                = "John"
	TEST_NAME_MARY                = "Mary"
	TEST_NAME_ARTHUR              = "Arthur Seagull"
	TEST_GAME_ONE                 = "game_01"
	TEST_GAME_FIVE                = "game_05"
	TEST_GAME_TEN                 = "game_10"
	TEST_LOG_NAME                 = "quake.log"
	TEST_LOG_PATH                 = "C:/Downloads/QuakeLogParser/"
	TEST_LOG_FULL_PATH            = "C:/Downloads/QuakeLogParser/quake.log"
	TEST_EXPECTED_GOT_STRING      = "Expected %s, got %s."
	TEST_EXPECTED_GOT_INT         = "Expected %d, got %d."
	TEST_EXPECTED_PLAYER          = "Expected player %s, found player."
	TEST_EXPECTED_SORTED_SLICE    = "Expected sorted slice, got unsorted slice."
	TEXT_EXPECTED_EMPTY_SLICE_BUT = "Expected empty slice, but got %s."
	EXPECTED_TOTAL_KILLS          = "Expected %d total kills, got %d."
	EXPECTED_PLAYERS              = "Expected %d players slice length, got %d."
	EXPECTED_KILLS                = "Expected %d kills map length, got %d."
	EXPECTED_DEATHS               = "Expected %d deaths map length, got %d."
	EXPECTED_MAP_LENGTH           = "Expected %d map length, got %d."
	EXPECTED_MAP_POSITIONS        = "Expected %d map positions with %d value, got %d different."
	EXPECTED_MAP_KEY_1            = "Expected map key to be %s as game slice first position is %d."
	EXPECTED_MAP_KEY_2            = "Expected map key to be %s as game slice position is %d."
	TEST_EXPECTED_PATH_EXIST      = "Expected path %s to exist!"
	TEST_FAILED_UNMARSHAL         = "Failed to unmarshal the JSON content: %v"
	TEST_UNEXPECTED_NUMBER        = "Unexpected number of elements in the result data. Got %d, want %d"

	REPORT_FILE_PERMISSION   = 0644
	PATH_CREATION_PERMISSION = 0700

	PLAYER_CONECT_INFO = `[^\\n](\w*|\w* )*`
)

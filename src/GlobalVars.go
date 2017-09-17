package src

// Global Variables
var (
    DBPath string = "/./yumdb";
    DatabaseInterface = make(map[string] []interface{});
)

// commands/flags
var (
    Targetdb string
    Action string
    Set string
    Run string

    Initialize bool

    // format string
    // yumFilePath string

    // time bool
    // script bool
)

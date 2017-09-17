package main

import (
    "flag"
    "fmt"
    "os"
)

import (
    globalVars "./src"
    // globalStructs "./src"

    "./src/sys"

    "./src/yum"
    yumMethod "./src/yum/methods"
    // "src/yum/readDatabase"
    // "src/yum/setDatabase"
    // "src/yum/updateDatabase"
    // "src/yum/deleteDatabase"

    // "src/execute"
    // "src/execute/YumFunctions"
)

func main() {

    // flags
    flag.StringVar(&globalVars.Action, "action", "", "This is comamnds to send to yumdb.")
    flag.StringVar(&globalVars.Targetdb, "t", "", "Targets a dtatabase to use/inspect/create.")
    flag.StringVar(&globalVars.Set, "set", "", "Set command is used for setting premission for you database, datatypes for data or other things.")

    flag.StringVar(&globalVars.Run, "run", "", "This runs a certain program or mode of developement; repl, prod, dev.")
    flag.BoolVar(&globalVars.Initialize, "init", false, "Initializies the yumdb store at the root folder.")

    flag.Parse()

    if sys.FileExist(globalVars.DBPath) == false {
        if globalVars.Initialize {
           yum.CreateYumStore()
        } else {
            fmt.Println("yum: yumdb was initialized, please run command `yum --init`.\nTo see other yum commands run `yum --help`.")
        }
    } else {

        if len( os.Args ) != 0 {
            dispathAction(globalVars.Action)
        } else {
            fmt.Println("yum: No command was given run command `yum --help` to see yum commands.")
        }
    }
}

func dispathAction(action string) {

    switch action {
    case "CREATE":
        yumMethod.Create_db()
        break;
    // case "DROP":
    // case"ERASE":
    // case "MERGE":
    // case "UPDATE":
    // case "SERVE":
    default:
        fmt.Printf("yum: command '%s' is unknown command in yum;\nRun command `yum --help` to see other commands.")
    }
}

/*
package main

import (
    "encoding/json"
    "io/ioutil"
    "reflect"
    "fmt"
)


func main() {

    // A map of the data being pulled out of that Compressed JSON
    var databaseInterface = make(map[string] []interface{})

    jsonFileBytes, err := getFileBytes("config.json")
    if err != nil {
        fmt.Println( err )
    } else {

        // Unmarshal the JSON data to the 'databaseInterface' map
        json.Unmarshal(jsonFileBytes, &databaseInterface)

        // iterates over each key in the map in its vlaues
        for k, v := range databaseInterface {

            if k != "row_names" {

                  iterates over all the items
                  in the array slices of each key vlaue

                for _, i := range v {
                    typeOf := reflect.TypeOf( i ).String()
                    if ( typeOf == "string" ) {
                        fmt.Println( i.(string) )
                    }
                    // databaseInterface[k] = assignType()
                }
            }
        }
    }
}

func getFileBytes( file_name string ) ([]byte, error) {

    byte_stream, err := ioutil.ReadFile( file_name )
    if err != nil {
       return []byte(""), err
    }
    return byte_stream, nil
}

/*
package main

import (
	"encoding/json"
	"reflect"
	"fmt"
)

func main() {

	var test = make(map[string] interface{})
	test["x"] = "fsdf"
	fmt.Println( reflect.TypeOf( test["x"] ), test["x"] )
	test["x"] = 45654
	fmt.Println( reflect.TypeOf( test["x"] ), test["x"] )
	x := test["x"].(int)
	if ( x > 122 ) {
		fmt.Println( true )
	}

	// turn it back unto JSON data
	JSON, err := json.Marshal(test)
	if err != nil {
		fmt.Println( err )
	} else {
		fmt.Println( string( JSON ) )
	}
	// test["x"] = test["x"].(int) + 1
	// x := isOfType(true, "float64")
	// fmt.Println( x )
}

// func isOfType(data interface{}, givenType string ) bool {

// 	if reflect.TypeOf( data ).String() == givenType {
// 		return true
// 	}
// 	return false
// }

    row := raw["age"]
    fmt.Println( row.(int) )
    // v := reflect.New(reflect.TypeOf(raw))
    // fmt.Println( v.Elem().Set(reflect.ValueOf(raw)) )
    // for _, i := range raw {
        // fmt.Printf("%T\n", i)
    // }
    // jsonParser := json.NewDecoder(jsonFile)
    // if err := jsonParser.Decode(&rowsStruct); err != nil {
    //     fmt.Println( err )
    // } else {
    //     fmt.Println( rowsStruct )
    // }
}

package main

import (
    // "encoding/json"
    "io/ioutil"
    "fmt"
)

func main() {
    data, err := ioutil.ReadFile("config.json")
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Print(string(data))
    // err = json.Unmarshal(data, interface{})
    // if err != nil {
    //     fmt.Println(err)
    //     return
    // }
    // fmt.Printf()
}

package main

import (
    // "encoding/json"
    "io/ioutil"
    "fmt"
)

func main() {
    data, err := ioutil.ReadFile("config.json")
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Print(string(data))
    // err = json.Unmarshal(data, interface{})
    // if err != nil {
    //     fmt.Println(err)
    //     return
    // }
    // fmt.Printf()
}
*/

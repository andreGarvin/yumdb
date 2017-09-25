/*
  Yumdb is a fast, light weight, very configurable
  and costimezable database mangement software
*/
package main

import (
    "flag"
    "fmt"
    "os"
)

import (
    globalVar "./src"
    yum "./src/yum"
    sys "./src/sys"
)

func main() {

    // flags
    flag.StringVar(&globalVar.Action, "action", "", "This is comamnds to send to yumdb.")
    flag.StringVar(&globalVar.TargetTable, "t", "", "Targets a dtatabase to use/inspect/create.")
    flag.StringVar(&globalVar.Set, "set", "", "Set command is used for setting premission for you database, datatypes for data or other things.")

    flag.StringVar(&globalVar.Run, "run", "", "This runs a certain program or mode of developement; repl, prod, dev.")
    flag.BoolVar(&globalVar.Initialize, "init", false, "Initializies the yumdb store at the root folder.")

    flag.Parse()

    if sys.FileExist(globalVar.YumPath) == false {
        if globalVar.Initialize {
            yum.CreateYumStore()
        } else {
            fmt.Println("yum: yumdb was initialized, please run command `yum --init`.\nTo see other yum commands run `yum --help`.")
        }
    } else {

        if len( os.Args ) != 0 {
            dispathAction(globalVar.Action)
        } else {
            fmt.Println("yum: No command was given run command `yum --help` to see yum commands.")
        }
    }
}

func dispathAction(action string) {

    switch action {
    case "CREATE":
        table, err := yum.Create_Table(globalVar.TargetTable, globalVar.Set)
        if err != nil {
            fmt.Println(err)
        } else {
            fmt.Println(table)
        }
        break;
    // case "SET":
    // case "DROP":
    // case"ERASE":
    // case "MERGE":
    // case "UPDATE":
    // case "SERVE":
    default:
        fmt.Printf("yum: command '%s' is unknown command in yum;\nRun command `yum --help` to see other commands.", action)
    }
}

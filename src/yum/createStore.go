package yum

// library
import (
    "path/filepath"
    "fmt"
)

// my library
import (
    globalVars "../."
    "../sys"
)


func CreateYumStore() {

    if err := sys.Mkdir(globalVars.DBPath); err != nil {
        fmt.Println(err)
    } else {

        dbFiles := []string{ "db-logs.yum", "yum.json" }

        for _, i := range dbFiles {
            path := filepath.Join( globalVars.DBPath, i )

            if err := sys.CreateFile(path, 0777); err != nil {
                fmt.Printf("Could not create file %s\n", path)
            } else {
                fmt.Printf("Created %s\n", path)
            }
        }
        fmt.Println("yum: yumdb was initialized and is ready to be used.")
    }
}

package yum

// library
import (
    "fmt"
)

// my library
import "../sys"


func CreateYumStore() {

    if err := sys.Mkdir("/./yumdb"); err != nil {
        fmt.Println(err)
    } else {

        dbFiles := []string{ "db-logs.yum", "yum.json" }

        for _, i := range dbFiles {

            if err := sys.CreateFile(i, 0777); err != nil {
                fmt.Printf("Could not create file %s\n", i)
            } else {
                fmt.Printf("Created %s\n", i)
            }
        }
    }
}

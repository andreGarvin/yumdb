package yum

import (
    "encoding/json"
    "path/filepath"
    "errors"
    "fmt"

    globalStruct "./src"
    globalVar "./src"
    "./src/sys"
)

type yumdb struct {
    TableName string
    Table *globalStruct.Table
}

func ( yum *yumdb ) ReadTableColumn( colName string ) *globalStruct.Payload {
      respPayload := &globalStruct.Payload{
          ColName: colName,
          Err: nil,
      }

      tablePath := filepath.Clean( filepath.Join(globalVar.YumPath, yum.TableName) )
      if sys.FileExist(tablePath) {
          tableFilePath := filepath.Clean( filepath.Join(tablePath, yum.TableName + "_table.json") )

          stream, err := sys.ReadFile(tableFilePath)
          if err == nil {

              err := json.Unmarshal([]byte(stream), &yum.Table)
              if err == nil {
                  cols := yum.Table.Cols
                  for col, values := range cols {
                      if col == colName {
                          respPayload.Values = values
                          respPayload.Coldef = yum.Table.Config.ColsStruct[col]
                          return respPayload
                      }
                  }
                  respPayload.Err = errors.New("yum: Column <" + colName + "> does not exist.")
                  return respPayload
              }
              respPayload.Err = err
              return respPayload
          }
          respPayload.Err = err
          return respPayload
      } else {
          respPayload.Err = errors.New("yum: Table <" + yum.TableName  + "> does not exist.")
      }
      return respPayload
}


func main() {
    Table := &yumdb{ TableName: "users" }

    TableCol := Table.ReadTableColumn("name")
    if TableCol.Err != nil {
        fmt.Println(TableCol.Err)
    } else {
        fmt.Println(TableCol)
    }
}

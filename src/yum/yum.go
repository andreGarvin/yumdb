package yum

import (
    "encoding/json"
    "path/filepath"
    // "reflect"
    "strings"
    "errors"
    // "fmt"

    globalStruct ".."
    globalVar ".."
    "../sys"
)

// The yumdb struct holding the table data fro the table file
type Yumdb struct {
    // the table name
    TableName string

    // the table data which is a struct of type Table
    Table *globalStruct.Table
}

func included( arr []string, item string ) bool {
     if len(arr) != 0 {
        for _, i := range arr {
            if i == item {
               return true
            }
        }
        return false
     }
     return false
}

// creates and constructs the the table file and returns back a struct of yumdb
func Create_Table( tableName, set string ) ( *Yumdb, error ) {
      tableName = strings.Trim(tableName, " ")
      set = strings.Trim(set, " ")

      if len(tableName) != 0 {
          tablePath := filepath.Clean(filepath.Join(globalVar.YumPath, tableName))

          if err := sys.Mkdir(tablePath); err != nil {
              return nil, err
          }

          if len(set) != 0 {
              var values []interface {}

              colsDefMap := make(map[string]globalStruct.ColumnsDef)
              cols := make(map[string] []interface{})

              NewTable := &globalStruct.Table{
                  Config: &globalStruct.ConfigStruct{
                      TableName: tableName,
                      Size: 0,
                      ColsStruct: colsDefMap,
                  },
                  Cols: cols,
              }

              for _, i := range strings.Split(set, " ") {
                  splitString := strings.Split(i, ":")

                  if strings.Contains(splitString[1], "prime") {
                      NewTable.Config.PrimaryKey = splitString[0]
                      splitString[1] = strings.Split(splitString[1], ">")[1]
                  }
                  prop := globalStruct.ColumnsDef{
                      Type: splitString[1],
                      Limit: 0,
                  }

                  NewTable.Config.ColsStruct[splitString[0]] = prop
                  NewTable.Cols[splitString[0]] = values
              }

              tableFileName := tableName + "_Table.json"
              TableFilePath := filepath.Clean(filepath.Join(tablePath, tableFileName))
              if err := sys.WriteJSONToFile(TableFilePath, NewTable); err != nil {
                  return nil, err
              }

              return &Yumdb{
                  TableName: tableName,
                  Table: NewTable,
              }, nil
          }
          return nil, errors.New("yum: There was not 'SET' data given to configure the Table;\n type `yum --help` see all commands")
      }
      return nil, errors.New("yum: Database name was not given; exit.")
}

// retruns table struct from yumdb and a error
func ( yum *Yumdb ) GetTable() ( *globalStruct.Table, error ) {

      if yum.Table != nil {
          return yum.Table, nil
      }

      yumTable, err := onLoadTable(yum.TableName)
      if err != nil {
          return nil, err
      } else {
          return yumTable.Table, nil
      }
}


// retunrs back the column of the given column name in a struct of type Payload
func ( yum *Yumdb ) ReadTableColumn( colName string ) *globalStruct.Payload {
      colName = strings.Trim(colName, " ")
      respPayload := &globalStruct.Payload{
          ColName: colName,
          Err: nil,
      }

      if len(colName) != 0 {

          if yum.Table != nil {

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


          yumTable, err := onLoadTable(yum.TableName)
          if err != nil {
              respPayload.Err = err
              return respPayload
          } else {
              return yumTable.ReadTableColumn(colName)
          }
    }

    respPayload.Err = errors.New("yum: No column name was.")
    return respPayload
}

// writes to the given column name and returns a error
func ( yum *Yumdb ) WriteToTable( input map[string] interface{} ) error {

      if yum.Table != nil {

          getMapKeys := func( Map map[string] interface {} ) ( keys []string ) {
                for k, _ := range Map {
                    keys = append(keys, k)
                }
                return keys
          }

          colNames, err := yum.GetColNames()
          if err != nil {
              return err
          }

          if len(input) != 0 {
             for e := 0; e < len(colNames); e++ {
                 col := yum.Table.Cols[colNames[e]]

                 if included(getMapKeys(input), colNames[e]) {
                    col = append(col, input[colNames[e]])
                 } else {
                    col = append(col, nil)
                }
                yum.Table.Cols[colNames[e]] = col
             }

             tableFilePath := filepath.Clean( filepath.Join(
                    // gets the path to the directory holding the table
                    filepath.Clean( filepath.Join(globalVar.YumPath, yum.TableName) ),
             yum.TableName + "_table.json") )

             if err := sys.WriteJSONToFile(tableFilePath, yum.Table); err != nil {
                return err
             }
             return nil
          }
          return errors.New("yum: no data was given")
      }

      yumTable, err := onLoadTable(yum.TableName)
      if err != nil {
          return err
      } else {
          return yumTable.WriteToTable(input)
      }
}

// returns back a string slice array of the column names in the file table
func ( yum *Yumdb ) GetColNames() ( []string, error ) {
      colNames := []string {}

      if yum.Table != nil {
          cols := yum.Table.Config.ColsStruct
          for col, _ := range cols {
              colNames = append(colNames, col)
          }
          return colNames, nil
      }

      yumTable, err := onLoadTable(yum.TableName)
      if err != nil {
          return colNames, err
      } else {
          return yumTable.GetColNames()
      }
}


// loads, parses, adn retruns back a struct if type yumdb
func onLoadTable(tableName string) ( *Yumdb, error ) {
      tableName = strings.Trim(tableName, " ")

      if len(tableName) != 0 {
            yum := &Yumdb{
                TableName: tableName,
            }

            tableDirPath := filepath.Clean( filepath.Join(globalVar.YumPath, yum.TableName) )
            if sys.FileExist(tableDirPath) {
                tableFilePath := filepath.Clean( filepath.Join(tableDirPath, yum.TableName + "_table.json") )

                stream, err := sys.ReadFile(tableFilePath)
                if err == nil {

                   err := json.Unmarshal([]byte(stream), &yum.Table)
                   if err == nil {
                      return yum, nil
                   }
                   return nil, err
                }
                return nil, err
            }
            return nil, errors.New("yum: Table <" + yum.TableName  + "> does not exist.")
      }
      return nil, errors.New("yum: No table name was given.")
}

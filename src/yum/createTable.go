package yum

import (
    "path/filepath"
    "strings"
    "errors"
)

import (
    globalStruct "../"
    globalVar "../"
    sys "../sys"
)

func Create_Table( tableName, set string ) (interface{}, error) {

      if tableName != "" {
          tablePath := filepath.Clean(filepath.Join(globalVar.YumPath, tableName))

          if err := sys.Mkdir(tablePath); err != nil {
              return nil, err
          } else {

              if len(set) != 0 {
                  var values []interface {}

                  colsDefMap := make(map[string]globalStruct.ColumnsDef)
                  cols := make(map[string] []interface{})

                  NewTable := &globalStruct.Table{
                      Config: &globalStruct.ConfigStruct{
                          TableName: "users",
                          Size: 0,
                          ColsStruct: colsDefMap,
                      },
                      Cols: cols,
                  }

                  for _, i := range strings.Split(set, " ") {
                      splitString := strings.Split(i, ":")
                      prop := globalStruct.ColumnsDef{
                          Type: splitString[1],
                          Limit: 0,
                      }

                      NewTable.Config.ColsStruct[splitString[0]] = prop
                      NewTable.Cols[splitString[0]] = values
                  }

                  tableName = tableName + "_Table.json"
                  TableFilePath := filepath.Clean(filepath.Join(tablePath, tableName))
                  if err := sys.WriteJSONToFile(TableFilePath, NewTable); err != nil {
                      return nil, err
                  }
                  return NewTable, nil
              }
              return nil, errors.New("yum: There was not 'SET' data given to configure the Table;\n type `yum --help` see all commands")
          }
      } else {
          return nil, errors.New("yum: Database name was not given; exit.")
      }

}

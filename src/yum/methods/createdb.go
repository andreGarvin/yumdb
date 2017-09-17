package methods

import "fmt"

import (
    // sys "../../sys"
    globalVars "../../"
)

type err interface {
    Error() string
}

func Create_db()  error {
      var db string = globalVars.Targetdb
      var ymError err

      dbConfig := getConfig( globalVars.Set )
      if db != "" {
          return nil
      }
      return yumError
}

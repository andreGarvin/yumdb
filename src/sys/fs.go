package sys

import (
    "path/filepath"
    "encoding/json"
    "io/ioutil"
    "os"
)


// checks if a file exist
func FileExist( path string ) bool {

    if _, err := os.Stat(path); !os.IsNotExist(err) {
        return true
    }
    return false
}

// returns the conents of a given existing file name
func ReadFile( file_name string ) ( string, error ) {

    byte_stream, err := ioutil.ReadFile( file_name )
    if err != nil {
        return "", err
    }
    return string( byte_stream ), nil
}

// writes to the file with the given data
func WriteToFile(filename string, input string) error {

     err := ioutil.WriteFile(filename, []byte( input ), 0744)
     if err != nil {
        return err
     }
     return nil
}

// takes a interface of some value, transforms that into JSON
// and writes to a certain file path
func WriteJSONToFile(path string, input interface{}) error {
    bytes, err := json.MarshalIndent(input, "", "  ")
    if err != nil {
        return err
    }

    if err := WriteToFile(path, string(bytes)); err != nil {
        return err
    }
    return nil
}

// makes a directory with the given dirname
func Mkdir( dirname string ) error {
    if FileExist(dirname) {
        return nil
    }

    if err := os.Mkdir(dirname, 0744); err != nil {
        return err
    }
    return nil
}

// creates  file with the given filename, and permisson
func CreateFile( filename string, mode os.FileMode) error {

    _, err := os.Create(filename)
    if err != nil {
        return err
    }

    if err := os.Chmod(filename, mode); err != nil {
        return err
    }
    return nil
}

// return a slice array with files of a certain extension
func FilesExt( sufix string, dirPath string ) ( []string, error ) {
      yumFiles := []string {}

      fs, err := ioutil.ReadDir(dirPath)
      if err != nil {
          return nil, err
      }

      for _, f := range fs {
          f := f.Name()

          ext := filepath.Ext(f)
          if ext != "" && ext == sufix {
              yumFiles = append(yumFiles, f)
          }
      }

      return yumFiles, nil
}

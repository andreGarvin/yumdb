package sys

import (
    "path/filepath"
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

func FilesWithSufix( sufix string, dirPath string ) ( []string, error ) {
      yumFiles := []string {}

      fs, err := ioutil.ReadDir(dirPath)
      if err != nil {
          return yumFiles, nil
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

package utils

import (
   "testing"
   "fmt"
)


func TestCreateDir(t *testing.T) {
    savePath, err := Mkdir("upload")
    if err != nil {
        fmt.Println(err.Error())
    }
    fmt.Println(savePath)
}

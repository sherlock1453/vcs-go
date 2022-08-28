package internal

import "os"

func PathExists(path string) bool{
    _, fileExistsErr := os.Stat(path)
    if fileExistsErr == nil{return true}
    if os.IsNotExist(fileExistsErr) {return false}
    return false
}

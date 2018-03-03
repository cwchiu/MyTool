package main

import (
    "fmt"
	"commands"
    _ "statik" 
)


    
    // go build -o tool.exe
func main() {
	defer func() {
		if err := recover(); err != nil {
        fmt.Println(err.(string))
		}
	}()
    commands.Execute()
}

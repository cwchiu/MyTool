package main

import (
	"fmt"

	"github.com/cwchiu/MyTool/commands"
	_ "github.com/cwchiu/MyTool/statik"
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

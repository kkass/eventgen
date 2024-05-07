package main

import (
	"fmt"
)

type FileCommand struct {
	FileDir string `short:"d" long:"dir" description:"Directory Name" default:"/tmp"`

	FileSleep int `short:s long:sleep description:"Desired Sleep Time" default:0`

	FileClose bool `long:close description:"Close File Before Exit" default:true`
}

var fileCommand FileCommand

func (x *FileCommand) Execute(args []string) error {
	fmt.Printf("Removing (force=%v): %#v\n", x.Force, args)
	return nil
}

func init() {
	parser.AddCommand("file",
		"Create a file",
		"",
		&fileCommand)
}

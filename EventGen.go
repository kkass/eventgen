package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jessevdk/go-flags"
)

var options struct {
	// Slice of bool will append 'true' each time the option
	// is encountered (can be set multiple times, like -vvv)
	Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information"`

	// Example of option group
	// File FileCommand `command:"file" shortDescription:"Create a file" longDescription:"" data:&fileCommand`
}

var parser = flags.NewParser(&options, flags.Default)

func main() {
	//fork, thread, fork-parent, daemon, exec, exec2, exec-bad
	//exec-raw, fexec, run, daemon2, crash

	// fileCmd := flag.NewFlagSet("file", flag.ExitOnError)
	// fileDir := fileCmd.String("dir", "/tmp", "Directory Name")
	// fileSleep := fileCmd.Int("sleep", 3, "Desired Sleep time")
	// fileClose := fileCmd.Bool("close", true, "Close file before exit")

	// switch os.Args[1] {
	// case "file":
	// 	fileCmd.Parse(os.Args[2:])
	// 	FileOp(*fileDir, *fileSleep, *fileClose)
	// default:
	// 	fmt.Println("invalid command:", os.Args[1])
	// 	os.Exit(1)
	// }

	if _, err := parser.Parse(); err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			fmt.Println("invalid command:1")
			os.Exit(1)
		default:
			os.Exit(1)
		}
	}
}

func FileOp(dir string, sleep int, shouldClose bool) {
	// fmt.Println("subcommand: file")
	// fmt.Println("  name:", name)
	// fmt.Println("  sleep:", sleep)
	// fmt.Println("  close:", shouldClose)

	if sleep > 0 {
		fmt.Println("Sleeping for ", sleep)
		time.Sleep(time.Duration(sleep) * time.Second)
	}

	file, err := os.CreateTemp(dir, "eventgen")
	if err != nil {
		fmt.Println("Error creating file ", file.Name())
		return
	}

	fmt.Println("Writing file ", file.Name())

	if shouldClose {
		defer file.Close()
		defer os.Remove(file.Name())
	} else {
		fmt.Println("Leaking file ", file.Name())
	}

	file.Write([]byte("foo\n"))
}

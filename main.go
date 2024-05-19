package main

import (
	"VCS/JSN"
	"fmt"

	"os"
)

func main() {

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Printf("JSN: no args provided\n")
		os.Exit(1)
	}

	command := args[0]
	switch command {
	case "init":
		{
			JSN.Init()
		}
	case "commit":
		JSN.Commit()
	default:
		fmt.Printf("JSN: %s is not a valid command\n", command)
		os.Exit(1)
	}
}

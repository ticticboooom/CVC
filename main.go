package main

import (
	"cvc/commands"
	"flag"
	"os"
)

func main() {
	command := os.Args[1]
	if command == "init" {
		flags := flag.NewFlagSet("init", flag.ExitOnError)
		commands.ParseRunInit(flags)
	}
	if command == "add" {
		flags := flag.NewFlagSet("add", flag.ExitOnError)
		commands.ParseRunAdd(flags)
	}

}

package main

import (
	"fmt"
	"os"
	"flag"
	"github.com/git-duet/git-duet/internal/cmd"
	"github.com/git-duet/git-duet/internal/cmdrunner"
)

func main() {
	var flags = flag.NewFlagSet("git duet-commit", flag.ContinueOnError)
	var amendFlag = flags.Bool("amend", false, "amends commit")

	var command cmd.Command
	if *amendFlag {
		command = cmd.New("commit")
	} else {
		command = cmd.NewWithSignoff("commit")
	}

	err := cmdrunner.Execute(command)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

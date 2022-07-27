package main

import (
	"fmt"
	"github.com/valbeat/whatday/command"
	"os"

	"github.com/urfave/cli"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{
	{
		Name:   "random",
		Usage:  "",
		Action: command.CmdRandom,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "list",
		Usage:  "",
		Action: command.CmdList,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}

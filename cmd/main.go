package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "Takuma Kajikawa"
	app.Email = ""
	app.Usage = ""

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	err := app.Run(os.Args)
	if err != nil {
		return
	}
}

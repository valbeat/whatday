package command

import (
	"flag"
	"testing"

	"github.com/urfave/cli"
)

func TestCmdList(t *testing.T) {
	app := cli.NewApp()
	set := flag.NewFlagSet("", 0)
	c := cli.NewContext(app, set, nil)

	command := cli.Command{
		Action: CmdList,
	}
	err := command.Run(c)
	if err != nil {
		t.Error(err)
	}
}

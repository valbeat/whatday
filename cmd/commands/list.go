package commands

import (
	"fmt"
	"github.com/valbeat/whatday/whatday"
	"time"

	"github.com/urfave/cli"
)

// CmdList prints article
func CmdList(c *cli.Context) {
	today := time.Now()
	articles, err := whatday.WhatDay(today)
	if err != nil {
		fmt.Println(err)
	}

	for _, article := range articles {
		fmt.Printf("%s\n%s\n", article.Title, article.Text)
	}
}

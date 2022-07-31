package commands

import (
	"fmt"
	"github.com/valbeat/whatday/whatday"
	"math/rand"
	"time"

	"github.com/urfave/cli"
)

// CmdRandom prints random article
func CmdRandom(c *cli.Context) {
	today := time.Now()
	articles, err := whatday.WhatDay(today)
	if err != nil {
		fmt.Println(err)
	}

	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(articles))
	article := articles[i]
	fmt.Printf("%s\n%s\n", article.Title, article.Text)
}

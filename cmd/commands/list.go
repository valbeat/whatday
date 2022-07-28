package commands

import (
	"fmt"
	"github.com/valbeat/whatday/whatday"
	"math/rand"
	"time"

	"github.com/urfave/cli"
)

// CmdList prints article
func CmdList(c *cli.Context) {
	t := time.Now()
	articles, err := whatday.New().GetArticles(t)
	if err != nil {
		fmt.Println(err)
	}

	rand.Seed(time.Now().UnixNano())
	for _, article := range articles {
		fmt.Printf("%s\n%s\n", article.Title, article.Text)
	}
}

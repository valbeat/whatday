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
	t := time.Now()
	articles, err := whatday.New(whatday.NewClient()).GetArticles(t)
	if err != nil {
		fmt.Println(err)
	}

	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(articles))
	article := articles[i]
	fmt.Printf("%s\n%s\n", article.Title, article.Text)
}

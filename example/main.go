package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/valbeat/whatday"
)

func main() {
	t := time.Now()
	articles, err := whatday.NewArticles(t)
	if err != nil {
		fmt.Println(err)
	}

	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(articles))
	article := articles[i]
	fmt.Printf("## %s\n> %s", article.Title, article.Text)
}

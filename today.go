package whatday

import (
	"fmt"
	"math/rand"
	"time"
)

// Print today article by default format
func Print() {
	t := time.Now()
	articles, err := NewArticles(t)
	if err != nil {
		fmt.Println(err)
	}

	i := rand.Intn(articles.Length())
	article := articles.Articles[i]

	println(article.Title)
	println(article.Text)
}

package whatday

import (
	"time"
)

func Print() {
	article := NewArticle(time.Now())
	println(article.Title)
	println(article.Text)
}

package whatday

import (
	"time"
)

// Print is
func Print() {
	article := NewArticle(time.Now())
	println(article.Title)
	println(article.Text)
}

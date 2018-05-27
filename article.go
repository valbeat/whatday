package whatday

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Article is
type Article struct {
	Title string
	Text  string
}

func NewArticle(t time.Time) Article {
	b, _ := GetListBody(t)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(b)))
	if err != nil {
		fmt.Println(err)
	}

	nodeList := doc.Find(".today_kinenbilist .winDetail")

	node := nodeList.Eq(rand.Intn(nodeList.Length()))
	href, _ := node.Attr("href")
	article := getArticle(href)
	return article
}

func getArticle(href string) Article {
	detail, _ := GetDetailBody(href)
	_doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(detail)))

	article := Article{}
	article.Title = strings.TrimSpace(_doc.Find("tr").First().Text())
	article.Text = strings.TrimSpace(_doc.Find("tr").Last().Text())
	return article
}

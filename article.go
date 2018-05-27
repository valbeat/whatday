package whatday

import (
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Articles is a struct
type Articles struct {
	Date     time.Time
	Articles []Article
}

// Length returns the number of articles.
func (a *Articles) Length() int {
	return len(a.Articles)
}

// Get retrieves the article at the specified index.
func (a *Articles) Get(i int) Article {
	return a.Articles[i]
}

// Article is a struct
type Article struct {
	Title string
	Text  string
}

// NewArticles return Articles
func NewArticles(t time.Time) (*Articles, error) {
	cli, err := NewClient("http://www.kinenbi.gr.jp/")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	res, err := cli.GetList(ctx, t)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	articles := decodeArticles(res.Body)

	return &Articles{
		Date:     t,
		Articles: articles,
	}, nil
}

func newArticle(spath string) (*Article, error) {
	cli, err := NewClient("http://www.kinenbi.gr.jp/" + spath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	res, err := cli.GetDetail(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()
	article := decodeArticle(res.Body)
	return article, nil
}

func decodeArticles(body io.Reader) []Article {

	doc, _ := goquery.NewDocumentFromReader(body)

	var articles []Article
	nodeList := doc.Find(".today_kinenbilist .winDetail")
	nodeList.Each(func(i int, node *goquery.Selection) {
		href, _ := node.Attr("href")
		article, _ := newArticle(href)
		articles = append(articles, *article)
	})
	return articles
}

func decodeArticle(body io.Reader) *Article {
	doc, _ := goquery.NewDocumentFromReader(body)

	return &Article{
		Title: strings.TrimSpace(doc.Find("tr").First().Text()),
		Text:  strings.TrimSpace(doc.Find("tr").Last().Text()),
	}
}

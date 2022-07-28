package whatday

import (
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Article is a struct
type Article struct {
	Title string
	Text  string
}

const EndPoint = "http://www.kinenbi.gr.jp/"

// NewArticles return Articles
func NewArticles(t time.Time) ([]Article, error) {
	cli, err := NewClient(EndPoint)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := cli.GetList(ctx, t)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	var articles []Article
	err = encodeArticles(res.Body, &articles)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func newArticle(spath string) (*Article, error) {
	cli, err := NewClient(EndPoint + spath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := cli.GetDetail(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)
	var article Article
	err = decodeArticle(res.Body, &article)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &article, nil
}

func encodeArticles(body io.Reader, articles *[]Article) error {

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return err
	}

	res := *articles

	nodeList := doc.Find(".today_kinenbilist .winDetail")
	nodeList.Each(func(i int, node *goquery.Selection) {
		href, _ := node.Attr("href")
		article, err := newArticle(href)
		if err != nil {
			fmt.Println(err)
			return
		}
		if article == nil {
			// article is nil but not error
			return
		}
		res = append(res, *article)
	})

	*articles = res
	return nil
}

func decodeArticle(body io.Reader, article *Article) error {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return err
	}
	article.Title = strings.TrimSpace(doc.Find("tr").First().Text())
	article.Text = strings.TrimSpace(doc.Find("tr").Last().Text())
	return nil
}

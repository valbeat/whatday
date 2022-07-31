package whatday

import (
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Whatday struct {
	client Client
}

func New(client Client) Whatday {
	return Whatday{
		client: client,
	}
}

// GetArticles return Articles
func (w Whatday) GetArticles(day time.Time) ([]Article, error) {
	pathList, err := w.getPaths(day)
	if err != nil {
		return nil, err
	}
	var articles []Article
	for _, path := range pathList {
		article, err := w.getArticle(path)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		articles = append(articles, *article)
	}
	return articles, nil
}

func (w Whatday) getPaths(day time.Time) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := w.client.GetList(ctx, day)
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

	var pathList []string
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return pathList, err
	}
	nodeList := doc.Find(".today_kinenbilist .winDetail")
	nodeList.Each(func(i int, node *goquery.Selection) {
		path, _ := node.Attr("href")
		pathList = append(pathList, path)
	})
	return pathList, nil
}

func (w Whatday) getArticle(path string) (*Article, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := w.client.GetDetail(ctx, path)
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

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}
	var article Article
	article.Title = strings.TrimSpace(doc.Find("tr").First().Text())
	article.Text = strings.TrimSpace(doc.Find("tr").Last().Text())
	return &article, nil
}

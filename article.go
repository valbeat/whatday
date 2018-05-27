package whatday

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/rand"
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

// Article is
type Article struct {
	Title string
	Text  string
}

func NewArticles(t time.Time) *Articles {

	b, _ := GetList(t)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(b)))
	if err != nil {
		fmt.Println(err)
	}

	var articles []Article

	nodeList := doc.Find(".today_kinenbilist .winDetail")
	nodeList.Each(func(i int, node *goquery.Selection) {
		href, _ := node.Attr("href")
		article := getArticle(href)
		articles = append(articles, article)
	})

	return &Articles{
		Date:     t,
		Articles: articles,
	}
}

func NewArticle(t time.Time) Article {
	articles := NewArticles(t)
	i := rand.Intn(articles.Length())
	return articles.Articles[i]
}

func GetList(now time.Time) ([]byte, error) {
	cli, err := NewClient("http://www.kinenbi.gr.jp/")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	res, err := cli.GetList(ctx, now)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return body, nil
}

func GetDetail(spath string) ([]byte, error) {
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

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return body, nil
}

func getArticle(href string) Article {
	detail, _ := GetDetail(href)
	_doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(detail)))

	article := Article{}
	article.Title = strings.TrimSpace(_doc.Find("tr").First().Text())
	article.Text = strings.TrimSpace(_doc.Find("tr").Last().Text())
	return article
}

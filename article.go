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

// Article is
type Article struct {
	Title string
	Text  string
}

func NewArticle(t time.Time) Article {
	b, _ := GetList(t)

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

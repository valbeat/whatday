package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const cacheDir string = "./cache/"

func GetListBody(now time.Time) ([]byte, error) {
	m := int(now.Month())
	d := int(now.Day())
	cacheName := cacheDir + "list_" + strconv.Itoa(m) + "-" + strconv.Itoa(d) + ".html"
	f, _ := os.Open(cacheName)
	defer f.Close()
	if f != nil {
		var b []byte
		buf := make([]byte, 10)
		for {
			n, err := f.Read(buf)
			if n == 0 {
				break
			}
			if err != nil {
				break
			}
			b = append(b, buf[:n]...)
		}
		return b, nil
	}
	endpoint := "http://www.kinenbi.gr.jp/"

	values := url.Values{}
	values.Add("M", strconv.Itoa(m))
	values.Add("D", strconv.Itoa(d))

	fmt.Println(values.Encode())
	res, err := http.PostForm(endpoint, values)
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

	f, _ = os.Create(cacheName)
	defer f.Close()
	f.Write(body)
	return body, nil
}

func GetDetailBody(path string) ([]byte, error) {
	cacheName := cacheDir + "detail_" + path + ".html"
	f, _ := os.Open(cacheName)
	defer f.Close()
	if f != nil {
		var b []byte
		buf := make([]byte, 10)
		for {
			n, err := f.Read(buf)
			if n == 0 {
				break
			}
			if err != nil {
				break
			}
			b = append(b, buf[:n]...)
		}
		return b, nil
	}

	base := "http://www.kinenbi.gr.jp/"
	req := base + path
	res, err := http.Get(req)
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

	f, _ = os.Create(cacheName)
	defer f.Close()
	f.Write(body)
	return body, nil
}

// Article is
type Article struct {
	Title string
	Text  string
}

func GetArticle(s *goquery.Selection) Article {
	href, _ := s.Attr("href")
	detail, _ := GetDetailBody(href)
	_doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(detail)))

	article := Article{}
	article.Title = strings.TrimSpace(_doc.Find("tr").First().Text())
	article.Text = strings.TrimSpace(_doc.Find("tr").Last().Text())
	return article
}

func NewArticle(t time.Time) Article {
	b, _ := GetListBody(t)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(b)))
	if err != nil {
		fmt.Println(err)
	}

	nodeList := doc.Find(".today_kinenbilist .winDetail")

	node := nodeList.Eq(rand.Intn(nodeList.Length()))
	article := GetArticle(node)
	return article
}

func Print() {
	article := NewArticle(time.Now())
	println(article.Title)
	println(article.Text)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	Print()
}

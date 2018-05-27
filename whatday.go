package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
)

const cacheDir string = "./cache/"

type Client struct {
	URL        *url.URL
	HTTPClient *http.Client

	Logger *log.Logger
}

func NewClient(urlStr string, logger *log.Logger) (*Client, error) {
	if len(urlStr) == 0 {
		return nil, errors.New("urlStr is empty")
	}
	if logger == nil {
		return nil, errors.New("logger is empty")
	}

	parsedURL, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return nil, errors.New("faild to pars url: {}")
	}

	httpClient := http.DefaultClient

	return &Client{
		URL:        parsedURL,
		HTTPClient: httpClient,
		Logger:     logger,
	}, nil
}

func (c *Client) newRequest(ctx context.Context, method string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, c.URL.String(), body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return req, nil
}

func (c *Client) GetList(ctx context.Context, now time.Time) (*http.Response, error) {
	m := int(now.Month())
	d := int(now.Day())

	values := url.Values{}
	values.Add("M", strconv.Itoa(m))
	values.Add("D", strconv.Itoa(d))

	body := strings.NewReader(values.Encode())

	req, err := c.newRequest(ctx, "POST", body)
	if err != nil {
		return nil, err
	}

	return c.HTTPClient.Do(req)
}

func (c *Client) GetDetail(ctx context.Context) (*http.Response, error) {
	req, err := c.newRequest(ctx, "GET", nil)
	if err != nil {
		return nil, err
	}
	return c.HTTPClient.Do(req)
}

func GetListBody(now time.Time) ([]byte, error) {
	logger := log.New()
	cli, err := NewClient("http://www.kinenbi.gr.jp/", logger)
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

func GetDetailBody(spath string) ([]byte, error) {
	logger := log.New()
	cli, err := NewClient("http://www.kinenbi.gr.jp/"+spath, logger)
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

func GetListBodyByCache(now time.Time) ([]byte, error) {
	m := int(now.Month())
	d := int(now.Day())
	cacheName := cacheDir + "list_" + strconv.Itoa(m) + "-" + strconv.Itoa(d) + ".html"
	return GetCacheBody(cacheName)
}

func GetDetailByCache(path string) ([]byte, error) {
	cacheName := cacheDir + "detail_" + path + ".html"
	return GetCacheBody(cacheName)
}

func GetCacheBody(cacheName string) ([]byte, error) {
	f, _ := os.Open(cacheName)
	defer f.Close()
	if f == nil {
		return nil, nil
	}
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

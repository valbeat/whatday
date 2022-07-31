package whatday

//go:generate mockgen -source client.go -destination client_mock.go -package whatday

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"net/url"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const version = "1.0.0"
const EndPoint = "https://www.kinenbi.gr.jp/"

var userAgent = fmt.Sprintf("whatday GoClient/%s (%v)", version, runtime.Version())

type Client interface {
	ListPath(ctx context.Context, now time.Time) ([]string, error)
	GetArticle(ctx context.Context, path string) (*Article, error)
}

// Client is a struct
type client struct {
	URL        *url.URL
	HTTPClient *http.Client

	Logger *log.Logger
}

// NewClient create client
func NewClient() Client {
	parsedURL, _ := url.ParseRequestURI(EndPoint)
	httpClient := http.DefaultClient
	return &client{
		URL:        parsedURL,
		HTTPClient: httpClient,
	}
}

func (c *client) ListPath(ctx context.Context, now time.Time) ([]string, error) {
	m := int(now.Month())
	d := now.Day()

	values := url.Values{}
	values.Add("M", strconv.Itoa(m))
	values.Add("D", strconv.Itoa(d))

	body := strings.NewReader(values.Encode())

	req, err := c.newRequest(ctx, "POST", c.URL.String(), body)
	if err != nil {
		return nil, err
	}

	response, err := c.HTTPClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}
	var pathList []string
	nodeList := doc.Find(".today_kinenbilist .winDetail")
	nodeList.Each(func(i int, node *goquery.Selection) {
		path, _ := node.Attr("href")
		pathList = append(pathList, path)
	})
	return pathList, nil
}

func (c *client) GetArticle(ctx context.Context, path string) (*Article, error) {
	url := c.URL.String()+path
	req, err := c.newRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	response, err := c.HTTPClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}
	var article Article
	article.Title = strings.TrimSpace(doc.Find("tr").First().Text())
	article.Text = strings.TrimSpace(doc.Find("tr").Last().Text())
	article.Url = url
	return &article, nil
}

// newRequest returns a new Request given a method, URL, and optional body.
func (c *client) newRequest(ctx context.Context, method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", userAgent)
	return req, nil
}

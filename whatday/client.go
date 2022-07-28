package whatday

//go:generate mockgen -source client.go -destination client_mock.go -package whatday

import (
	"context"
	"errors"
	"fmt"
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

var userAgent = fmt.Sprintf("whatday GoClient/%s (%v)", version, runtime.Version())

type Client interface {
	GetList(ctx context.Context, now time.Time) (*http.Response, error)
	GetDetail(ctx context.Context, path string) (*http.Response, error)
}

// Client is a struct
type client struct {
	URL        *url.URL
	HTTPClient *http.Client

	Logger *log.Logger
}

// NewClient create client
func NewClient(urlStr string) (Client, error) {
	if len(urlStr) == 0 {
		return nil, errors.New("urlStr is empty")
	}
	parsedURL, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return nil, errors.New("failed to parse url: {}")
	}

	httpClient := http.DefaultClient
	return &client{
		URL:        parsedURL,
		HTTPClient: httpClient,
	}, nil
}

// GetList return list
func (c *client) GetList(ctx context.Context, now time.Time) (*http.Response, error) {
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

	return c.HTTPClient.Do(req)
}

// GetDetail return detail
func (c *client) GetDetail(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.newRequest(ctx, "GET", c.URL.String() + path, nil)
	if err != nil {
		return nil, err
	}
	return c.HTTPClient.Do(req)
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

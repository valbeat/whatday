package whatday

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

// Client is a struct
type Client struct {
	URL        *url.URL
	HTTPClient *http.Client

	Logger *log.Logger
}

// NewClient create client
func NewClient(urlStr string) (*Client, error) {
	if len(urlStr) == 0 {
		return nil, errors.New("urlStr is empty")
	}
	parsedURL, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return nil, errors.New("faild to pars url: {}")
	}

	httpClient := http.DefaultClient

	return &Client{
		URL:        parsedURL,
		HTTPClient: httpClient,
	}, nil
}

// NewRequest returns a new Request given a method, URL, and optional body.
func (c *Client) NewRequest(ctx context.Context, method string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, c.URL.String(), body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", userAgent)
	return req, nil
}

// GetList return list
func (c *Client) GetList(ctx context.Context, now time.Time) (*http.Response, error) {
	m := int(now.Month())
	d := int(now.Day())

	values := url.Values{}
	values.Add("M", strconv.Itoa(m))
	values.Add("D", strconv.Itoa(d))

	body := strings.NewReader(values.Encode())

	req, err := c.NewRequest(ctx, "POST", body)
	if err != nil {
		return nil, err
	}

	return c.HTTPClient.Do(req)
}

// GetDetail return detail
func (c *Client) GetDetail(ctx context.Context) (*http.Response, error) {
	req, err := c.NewRequest(ctx, "GET", nil)
	if err != nil {
		return nil, err
	}
	return c.HTTPClient.Do(req)
}

package whatday

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

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

func (c *Client) newRequest(ctx context.Context, method string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, c.URL.String(), body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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

	req, err := c.newRequest(ctx, "POST", body)
	if err != nil {
		return nil, err
	}

	return c.HTTPClient.Do(req)
}

// GetDetail return detail
func (c *Client) GetDetail(ctx context.Context) (*http.Response, error) {
	req, err := c.newRequest(ctx, "GET", nil)
	if err != nil {
		return nil, err
	}
	return c.HTTPClient.Do(req)
}

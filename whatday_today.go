package whatday

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

const cacheDir string = "./cache/"

func GetListBody(now time.Time) ([]byte, error) {
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

func GetDetailBody(spath string) ([]byte, error) {
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

func Print() {
	article := NewArticle(time.Now())
	println(article.Title)
	println(article.Text)
}

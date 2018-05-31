package whatday

import (
	"io"
	"os"
	"strconv"
	"time"
)

const cacheDir string = "./cache/"

// ReadList returns a list reader
func ReadList(now time.Time) (io.Reader, error) {
	m := int(now.Month())
	d := int(now.Day())
	cacheName := cacheDir + "list_" + strconv.Itoa(m) + "-" + strconv.Itoa(d) + ".html"
	return read(cacheName)
}

// ReadArticle returns an article reader
func ReadArticle(path string) (io.Reader, error) {
	cacheName := cacheDir + "article_" + path + ".html"
	return read(cacheName)
}

func read(cacheName string) (io.Reader, error) {
	f, _ := os.Open(cacheName)
	defer f.Close()
	if f == nil {
		return nil, nil
	}
	var r io.Reader
	io.Copy(f, r)
	return r, nil
}

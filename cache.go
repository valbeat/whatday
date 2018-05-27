package whatday

import (
	"os"
	"strconv"
	"time"
)

const cacheDir string = "./cache/"

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

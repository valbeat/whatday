package main

import (
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func TestGetListBody(t *testing.T) {
	res, _ := GetListBody(time.Now())
	if res == nil {
		t.Errorf("GetListBody() return empty ")
	}
}

func TestGetArticle(t *testing.T) {
	file, _ := os.Open(`./list_5-3.html`)
	doc, _ := goquery.NewDocumentFromReader(file)

	nodeList := doc.Find(".today_kinenbilist .winDetail")

	// rand.Seed(time.Now().UnixNano)
	node := nodeList.Eq(rand.Intn(nodeList.Length()))
	res := GetArticle(node)

	if res.Title == "" {
		t.Errorf("title is empty ")
	}
	if res.Text == "" {
		t.Errorf("text is empty")
	}

}

func TestNewArticle(t *testing.T) {
	got := NewArticle()
	want := Article{Title: "そうじの日", Text: "神奈川県横浜市に本部を置き、掃除技術についての研究や普及活動などを行っている一般財団法人日本そうじ協会が制定。日付は５と３の語呂合わせの「ゴミ」と「護美」からで、「ゴミを減らすこと」と「環境の美しさを護ること」が目的。この日には全国一斉に「おそうじしましょう！」と呼びかける。日本そうじ協会では環境整備の技術力を高め、良い習慣を身につける「掃除道」の普及促進も行っている。"}
	if got != want {
		t.Errorf("NewArticle is invalid %v", got)
	}
}

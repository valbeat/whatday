package whatday

import (
	"testing"
	"time"
)

func TestGetListBody(t *testing.T) {
	res, _ := GetListBody(time.Now())
	if res == nil {
		t.Errorf("GetListBody() return empty ")
	}
}

func TestNewArticle(t *testing.T) {
	date := time.Date(2018, 5, 9, 0, 0, 0, 0, time.Local)
	got := NewArticle(date)
	// ref: http://www.kinenbi.gr.jp/
	want := Article{Title: "合格の日", Text: "福岡県福岡市に本社を置き、全国、海外に店舗を展開する天然とんこつラーメン専門店の株式会社一蘭が制定。同社では福岡県太宰府市の太宰府参道店で「合格ラーメン」を提供していることから、入学や資格試験などを受ける受験生を応援するのが目的。日付は５と９で「合（５）格（９）」と読む語呂合わせから。「合格ラーメン」は五角形の器に長さ５９センチの麺が入っているなど、合格（ごうかく）にこだわった内容が人気。"}
	if got != want {
		t.Errorf("NewArticle is invalid %v", got)
	}
}

package whatday

import (
	"math/rand"
	"testing"
	"time"
)

func TestNewArticles(t *testing.T) {
	date := time.Date(2018, 5, 9, 0, 0, 0, 0, time.Local)
	articles, err := New().GetArticles(date)
	if err != nil {
		t.Error(err)
	}

	i := rand.Intn(len(articles))
	got := articles[i]
	if got.Title == "" || got.Text == "" {
		t.Errorf("NewArticle is invalid %v", got)
	}
}

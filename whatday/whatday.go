package whatday

import (
	"context"
	"fmt"
	"time"
)

type Whatday struct {
	client Client
}

func New(client Client) Whatday {
	return Whatday{
		client: client,
	}
}

// GetArticles return Articles
func (w Whatday) GetArticles(day time.Time) ([]Article, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pathList, err := w.client.ListPath(ctx, day)
	if err != nil {
		return nil, err
	}
	var articles []Article
	for _, path := range pathList {
		article, err := w.client.GetArticle(ctx, path)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		articles = append(articles, *article)
	}
	return articles, nil
}

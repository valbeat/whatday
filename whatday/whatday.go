package whatday

import (
	"context"
	"fmt"
	"time"
)

type whatday struct {
	client Client
}

func WhatDay(day time.Time) ([]Article, error) {
	return whatday{NewClient()}.getArticles(day)
}

// getArticles return Articles
func (w whatday) getArticles(day time.Time) ([]Article, error) {
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

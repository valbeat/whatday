package whatday

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	_ = NewClient()
}

func TestClientGetList(t *testing.T) {
	t.Skip("Only development")
	cli := NewClient()
	ctx := context.Background()
	paths, err := cli.ListPath(ctx, time.Now())
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%s", paths)
}

func TestClientGetDetail(t *testing.T) {
	t.Skip("Only development")
	cli := NewClient()
	ctx := context.Background()
	article, err := cli.GetArticle(ctx, "yurai_other.php?MD=4&NM=10088")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%s", article)
}

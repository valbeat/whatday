package whatday

import (
	"context"
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
	_, err := cli.GetList(ctx, time.Now())
	if err != nil {
		t.Error(err)
	}
}

func TestClientGetDetail(t *testing.T) {
	t.Skip("Only development")
	cli := NewClient()
	ctx := context.Background()
	_, err := cli.GetDetail(ctx, "")
	if err != nil {
		t.Error(err)
	}
}

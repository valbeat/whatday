package whatday

import (
	"context"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	_, err := NewClient("http://www.kinenbi.gr.jp/")
	if err != nil {
		t.Fatalf("Failed to construct client: %s", err)
	}
}

func TestClientNewRequest(t *testing.T) {
	cli, err := NewClient("http://www.kinenbi.gr.jp/")
	if err != nil {
		t.Fatalf("Failed to construct client: %s", err)
	}
	ctx := context.Background()
	_, err = cli.NewRequest(ctx, "GET", nil)
	if err != nil {
		t.Error(err)
	}
}

func TestClientGetList(t *testing.T) {
	cli, err := NewClient("http://www.kinenbi.gr.jp/")
	if err != nil {
		t.Fatalf("Failed to construct client: %s", err)
	}
	ctx := context.Background()
	_, err = cli.GetList(ctx, time.Now())
	if err != nil {
		t.Error(err)
	}
}

func TestClientGetDetail(t *testing.T) {
	cli, err := NewClient("http://www.kinenbi.gr.jp/")
	if err != nil {
		t.Fatalf("Failed to construct client: %s", err)
	}
	ctx := context.Background()
	_, err = cli.GetDetail(ctx)
	if err != nil {
		t.Error(err)
	}
}

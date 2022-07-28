package whatday

import (
	"context"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	_, err := NewClient(EndPoint)
	if err != nil {
		t.Fatalf("Failed to construct client: %s", err)
	}
}

func TestClientGetList(t *testing.T) {
	t.Skip("Only development")
	cli, err := NewClient(EndPoint)
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
	t.Skip("Only development")
	cli, err := NewClient(EndPoint)
	if err != nil {
		t.Fatalf("Failed to construct client: %s", err)
	}
	ctx := context.Background()
	_, err = cli.GetDetail(ctx, "")
	if err != nil {
		t.Error(err)
	}
}

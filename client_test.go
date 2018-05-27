package whatday

import "testing"

func TestNewClient(t *testing.T) {
	_, err := NewClient("http://www.kinenbi.gr.jp/")
	if err != nil {
		t.Fatalf("Failed to construct client: %s", err)
	}
}

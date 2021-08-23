package rc

import (
	"context"
	"testing"
	"time"
)

func TestClient_Account(t *testing.T) {
	c, _ := NewFromKeySecret("2827d81f-6ae0-4842-b92f-6576afe36863", "ad6a84a9edf84d42ac1e1915b194529f", "a7f46d5fad1f43ca9bcbdc0afc22aad0")

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	account, err := c.ReadMe(ctx)
	t.Log(account)
	if err != nil {
		t.Error(err)
	}
}

func TestClient_Register(t *testing.T) {
	c, _ := NewFromKeySecret("2827d81f-6ae0-4842-b92f-6576afe36863", "ad6a84a9edf84d42ac1e1915b194529f", "a7f46d5fad1f43ca9bcbdc0afc22aad0")

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	account, err := c.Register(ctx)
	t.Log(account)
	if err != nil {
		t.Error(err)
	}
}

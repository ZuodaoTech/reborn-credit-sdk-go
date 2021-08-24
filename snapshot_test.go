package rc

import (
	"context"
	"testing"
	"time"
)

func TestClient_ReadSnapshots(t *testing.T) {
	c, _ := NewFromKeySecret("2827d81f-6ae0-4842-b92f-6576afe36863", "ad6a84a9edf84d42ac1e1915b194529f", "a7f46d5fad1f43ca9bcbdc0afc22aad0")

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	result, err := c.ReadSnapshots(ctx, "rbc", time.Now(), "DESC", 100)
	t.Log(result)
	if err != nil {
		t.Error(err)
	}
}

func TestClient_ReadSnapshot(t *testing.T) {
	c, _ := NewFromKeySecret("2827d81f-6ae0-4842-b92f-6576afe36863", "ad6a84a9edf84d42ac1e1915b194529f", "a7f46d5fad1f43ca9bcbdc0afc22aad0")

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	result, err := c.ReadSnapshot(ctx, "11bf2d27-0f3b-494e-8a3d-44673b464ca9")
	t.Log(result)
	if err != nil {
		t.Error(err)
	}
}

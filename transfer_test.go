package rc

import (
	"context"
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

func TestClient_Transfer(t *testing.T) {
	c, _ := NewFromKeySecret("2827d81f-6ae0-4842-b92f-6576afe36863", "ad6a84a9edf84d42ac1e1915b194529f", "a7f46d5fad1f43ca9bcbdc0afc22aad0")

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var input = &TransferInput{
		AssetID:    "rbc",
		OpponentID: "a30bf8dd-e87f-4ff7-9da8-4abad6f3234e",
		Amount:     decimal.NewFromFloat(2),
		Memo:       "unit test",
	}

	result, err := c.Transfer(ctx, input)
	t.Log(result)
	if err != nil {
		t.Error(err)
	}
}

func TestClientGenerateTransfer(t *testing.T) {
	c, _ := NewFromKeySecret("2827d81f-6ae0-4842-b92f-6576afe36863", "ad6a84a9edf84d42ac1e1915b194529f", "a7f46d5fad1f43ca9bcbdc0afc22aad0")

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var input = &GenerateTransferInput{
		AssetID: "rbc",
		UserID:  "a30bf8dd-e87f-4ff7-9da8-4abad6f3234e",
		Amount:  decimal.NewFromFloat(2),
		Memo:    "unit test",
	}

	result, err := c.GenerateTransfer(ctx, input)
	t.Log(result)
	if err != nil {
		t.Error(err)
	}
}

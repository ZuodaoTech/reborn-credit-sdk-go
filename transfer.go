package rc

import (
	"context"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

// Transfer Transfer
type Transfer struct {
	ID         int64           `json:"id,omitempty"`
	SnapshotID string          `json:"snapshot_id,omitempty"`
	TraceID    string          `json:"trace_id,omitempty"`
	OpponentID string          `json:"opponent_id,omitempty"`
	UserID     string          `json:"user_id,omitempty"`
	AssetID    string          `json:"asset_id,omitempty"`
	Amount     decimal.Decimal `json:"amount,omitempty"`
	Memo       string          `json:"memo,omitempty"`
	Type       string          `json:"type,omitempty"`
	Status     string          `json:"status,omitempty"`
	CreatedAt  *time.Time      `json:"created_at,omitempty"`
}

// TransferInput TransferInput
type TransferInput struct {
	AssetID    string          `json:"asset_id,omitempty" binding:"required"`
	OpponentID string          `json:"opponent_id,omitempty" binding:"required"`
	Amount     decimal.Decimal `json:"amount,omitempty" binding:"required"`
	TraceID    string          `json:"trace_id,omitempty"`
	Memo       string          `json:"memo,omitempty"`
}

// GenerateTransferInput GenerateTransferInput
type GenerateTransferInput struct {
	AssetID string          `json:"asset_id,omitempty" binding:"required"`
	UserID  string          `json:"user_id,omitempty" binding:"required"`
	Amount  decimal.Decimal `json:"amount,omitempty" binding:"required"`
	TraceID string          `json:"trace_id,omitempty"`
	Memo    string          `json:"memo,omitempty"`
}

// Transfer Transfer
func (c *Client) Transfer(ctx context.Context, input *TransferInput) (*Transfer, error) {
	var transfer Transfer
	var body struct {
		TransferInput
		UserID string `json:"user_id,omitempty"`
	}
	body.TransferInput = *input
	body.UserID = c.ClientID

	if err := c.Post(ctx, "/v1/app/transfers", body, &transfer); err != nil {
		return nil, err
	}

	return &transfer, nil
}

// GenerateTransfer GenerateTransfer
func (c *Client) GenerateTransfer(ctx context.Context, input *GenerateTransferInput) (*Transfer, error) {
	var transfer Transfer
	var body struct {
		GenerateTransferInput
		OpponentID string `json:"opponent_id,omitempty"`
	}
	body.GenerateTransferInput = *input
	body.OpponentID = c.ClientID

	if err := c.Post(ctx, "/v1/app/transfers/generate", body, &transfer); err != nil {
		return nil, err
	}

	return &transfer, nil
}

func (c *Client) ReadTransfer(ctx context.Context, traceID string) (*Transfer, error) {
	uri := fmt.Sprintf("/v1/app/transfers/trace/%s", traceID)

	var transfer Transfer
	if err := c.Get(ctx, uri, nil, &transfer); err != nil {
		return nil, err
	}

	return &transfer, nil
}

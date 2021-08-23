package rc

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
)

// Snapshot Snapshot
type Snapshot struct {
	ID         string          `json:"snapshot_id,omitempty"`
	TraceID    string          `json:"trace_id,omitempty"`
	AssetID    string          `json:"asset_id,omitempty"`
	Amount     decimal.Decimal `json:"amount,omitempty"`
	UserID     string          `json:"user_id,omitempty"`
	OpponentID string          `json:"opponent_id,omitempty"`
	Memo       string          `json:"memo,omitempty"`
	Type       string          `json:"type,omitempty"`
	CreatedAt  *time.Time      `json:"created_at,omitempty"`
	Asset      *Asset          `json:"asset,omitempty"`
}

func (c *Client) ReadSnapshots(ctx context.Context, assetID string, offset time.Time, order string, limit int) ([]*Snapshot, error) {
	var snapshots []*Snapshot
	params := buildReadSnapshotsParams(assetID, offset, order, limit)
	if err := c.Get(ctx, fmt.Sprintf("/v1/app/accounts/%s/snapshots", c.ClientID), params, &snapshots); err != nil {
		return nil, err
	}

	return snapshots, nil
}

func (c *Client) ReadSnapshot(ctx context.Context, snapshotID string) (*Snapshot, error) {
	var snapshot Snapshot
	if err := c.Get(ctx, fmt.Sprintf("/v1/app/accounts/%s/snapshots/%s", c.ClientID, snapshotID), nil, &snapshot); err != nil {
		return nil, err
	}

	return &snapshot, nil
}

func (c *Client) ReadPublicSnapshots(ctx context.Context, assetID string, offset time.Time, order string, limit int) ([]*Snapshot, error) {
	var snapshots []*Snapshot
	params := buildReadSnapshotsParams(assetID, offset, order, limit)
	if err := c.Get(ctx, "/v1/public/snapshots", params, &snapshots); err != nil {
		return nil, err
	}

	return snapshots, nil
}

func (c *Client) ReadPublicSnapshot(ctx context.Context, snapshotID string) (*Snapshot, error) {
	uri := fmt.Sprintf("/v1/public/snapshots/%s", snapshotID)

	var snapshot Snapshot
	if err := c.Get(ctx, uri, nil, &snapshot); err != nil {
		return nil, err
	}

	return &snapshot, nil
}

func buildReadSnapshotsParams(assetID string, offset time.Time, order string, limit int) map[string]string {
	params := make(map[string]string)

	if assetID != "" {
		params["asset"] = assetID
	}

	if !offset.IsZero() {
		params["offset"] = offset.UTC().Format(time.RFC3339Nano)
	}

	switch order {
	case "ASC", "DESC":
	default:
		order = "DESC"
	}

	params["order"] = order

	if limit > 0 {
		params["limit"] = strconv.Itoa(limit)
	}

	return params
}

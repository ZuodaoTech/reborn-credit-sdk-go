package rc

import (
	"context"
	"fmt"
	"time"
)

const (
	UserAssetStatusNormal = "normal"
	UserAssetStatusLocked = "locked"
)

// Asset Asset
type Asset struct {
	AssetID        string     `json:"asset_id,omitempty"`
	Name           string     `json:"name,omitempty"`
	Symbol         string     `json:"symbol,omitempty"`
	IconUrl        string     `json:"icon_url,omitempty"`
	Capitalization string     `json:"capitalization,omitempty"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	UpdatedAt      *time.Time `json:"upated_at,omitempty"`
	Status         string     `json:"status,omitempty"`
	Balance        string     `json:"balance,omitempty"`
}

func (c *Client) ReadAsset(ctx context.Context, assetID string) (*Asset, error) {
	uri := fmt.Sprintf("/v1/app/accounts/%s/assets/%s", c.ClientID, assetID)

	var asset Asset
	if err := c.Get(ctx, uri, nil, &asset); err != nil {
		return nil, err
	}

	return &asset, nil
}

func (c *Client) ReadAssets(ctx context.Context) ([]*Asset, error) {
	uri := fmt.Sprintf("/v1/app/accounts/%s/assets", c.ClientID)

	var assets []*Asset
	if err := c.Get(ctx, uri, nil, &assets); err != nil {
		return nil, err
	}

	return assets, nil
}

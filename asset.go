package rbc

import "time"

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

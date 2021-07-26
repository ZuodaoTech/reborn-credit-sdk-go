package rbc

import (
	"context"
	"fmt"
	"time"
)

const (
	AccountStatusNormal  = "normal"
	AccountStatusBlocked = "blocked"
)

const (
	AccountTypeUser = "user"
	AccountTypeApp  = "app"
)

// Account Account
type Account struct {
	UserID    string     `json:"user_id,omitempty"`
	Type      string     `json:"type,omitempty"`
	Status    string     `json:"status,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Assets    []*Asset   `json:"assets,omitempty"`
}

// Register Register
func (c *Client) Register(ctx context.Context) (*Account, error) {
	var account Account
	if err := c.Post(ctx, fmt.Sprintf("/v1/app/accounts/%s/register", c.ClientID), nil, &account); err != nil {
		return nil, err
	}

	return &account, nil
}

// ReadUserAccount ReadUserAccount
func (c *Client) ReadUserAccount(ctx context.Context, userID string) (*Account, error) {
	var account Account
	if err := c.Get(ctx, fmt.Sprintf("/v1/app/accounts/%s", userID), nil, &account); err != nil {
		return nil, err
	}

	return &account, nil
}

// ReadMe ReadMe
func (c *Client) ReadMe(ctx context.Context) (*Account, error) {
	var account Account
	if err := c.Get(ctx, fmt.Sprintf("/v1/app/accounts/%s", c.ClientID), nil, &account); err != nil {
		return nil, err
	}
	return &account, nil
}

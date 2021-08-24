package rc

import (
	"context"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	Auth
	ClientID string
}

func NewFromKeySecret(clientID, appKey, appSecret string) (*Client, error) {
	return &Client{
		ClientID: clientID,
		Auth: &BasicAuth{
			Key:    appKey,
			Secret: appSecret,
		},
	}, nil
}

func (c *Client) Request(ctx context.Context) *resty.Request {
	ctx = WithAuth(ctx, c.Auth)
	return Request(ctx)
}

func (c *Client) Get(ctx context.Context, uri string, params map[string]string, resp interface{}) error {
	r, err := c.Request(ctx).SetQueryParams(params).Get(uri)
	if err != nil {
		if requestID := extractRequestID(r); requestID != "" {
			return WrapErrWithRequestID(err, requestID)
		}

		return err
	}

	return UnmarshalResponse(r, resp)
}

func (c *Client) Post(ctx context.Context, uri string, body interface{}, resp interface{}) error {
	r, err := c.Request(ctx).SetBody(body).Post(uri)
	if err != nil {
		if requestID := extractRequestID(r); requestID != "" {
			return WrapErrWithRequestID(err, requestID)
		}

		return err
	}

	return UnmarshalResponse(r, resp)
}

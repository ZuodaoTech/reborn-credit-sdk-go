package rc

import (
	"context"
)

type contextKey int

const (
	_ contextKey = iota
	authKey
	verifierKey
	requestIdKey
)

func WithVerifier(ctx context.Context, v Verifier) context.Context {
	return context.WithValue(ctx, verifierKey, v)
}

func WithAuth(ctx context.Context, basicAuth Auth) context.Context {
	return context.WithValue(ctx, authKey, basicAuth)
}

// WithRequestID bind request id to context
// request id must be uuid
func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIdKey, requestID)
}

var newRequestID = newUUID

func RequestIdFromContext(ctx context.Context) string {
	if v, ok := ctx.Value(requestIdKey).(string); ok {
		return v
	}

	return newRequestID()
}

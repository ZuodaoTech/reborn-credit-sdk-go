package rc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

var (
	xRequestID           = http.CanonicalHeaderKey("x-request-id")
	xIntegrityToken      = http.CanonicalHeaderKey("x-integrity-token")
	xForceAuthentication = http.CanonicalHeaderKey("x-force-authentication")

	ErrResponseVerifyFailed = errors.New("response verify failed")
)

var httpClient = resty.New().
	SetHeader("Content-Type", "application/json").
	SetHostURL(DefaultApiHost).
	SetTimeout(10 * time.Second).
	SetPreRequestHook(func(c *resty.Client, r *http.Request) error {
		ctx := r.Context()
		requestID := r.Header.Get(xRequestID)
		if requestID == "" {
			requestID = RequestIdFromContext(ctx)
			r.Header.Set(xRequestID, requestID)
		}

		if s, ok := ctx.Value(authKey).(Auth); ok {
			token := s.GenToken()
			r.Header.Set("Authorization", "Basic "+token)
			r.Header.Set(xForceAuthentication, "true")
		}

		return nil
	}).
	OnAfterResponse(func(c *resty.Client, r *resty.Response) error {
		if r.IsError() {
			return nil
		}

		if err := checkResponseRequestID(r); err != nil {
			return err
		}

		if v, ok := r.Request.Context().Value(verifierKey).(Verifier); ok {
			if err := v.Verify(r); err != nil {
				return ErrResponseVerifyFailed
			}
		}

		return nil
	})

func GetClient() *http.Client {
	return httpClient.GetClient()
}

func GetRestyClient() *resty.Client {
	return httpClient
}

func checkResponseRequestID(r *resty.Response) error {
	expect := r.Request.Header.Get(xRequestID)
	got := r.Header().Get(xRequestID)
	if expect != got {
		return fmt.Errorf("%s mismatch, expect %q but got %q", xRequestID, expect, got)
	}

	return nil
}

func Request(ctx context.Context) *resty.Request {
	return httpClient.R().SetContext(ctx)
}

func UnmarshalResponse(resp *resty.Response, v interface{}) (err error) {
	if requestID := extractRequestID(resp); requestID != "" {
		defer bindRequestID(&err, requestID)
	}

	if resp.IsSuccess() {
		if err := json.Unmarshal([]byte(resp.Body()), v); err != nil {
			var appErr Error
			if err := json.Unmarshal([]byte(resp.Body()), &appErr); err == nil {
				appErr.Status = resp.StatusCode()
				return &appErr
			}
			return createError(resp.StatusCode(), resp.StatusCode(), err.Error())
		}
		return nil
	}

	if resp.IsError() {
		var appErr Error
		if err := json.Unmarshal([]byte(resp.Body()), &appErr); err == nil {
			appErr.Status = resp.StatusCode()
			return &appErr
		}
		return createError(resp.StatusCode(), resp.StatusCode(), resp.Status())
	}

	return nil
}

func extractRequestID(r *resty.Response) string {
	if r != nil {
		return r.Request.Header.Get(xRequestID)
	}

	return ""
}

func bindRequestID(errp *error, id string) {
	if err := *errp; err != nil {
		*errp = WrapErrWithRequestID(err, id)
	}
}

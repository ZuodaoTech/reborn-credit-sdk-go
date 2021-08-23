package rc

import (
	"errors"
	"fmt"
)

type Error struct {
	Status    int                    `json:"status"`
	Code      int                    `json:"code"`
	Message   string                 `json:"message"`
	Extra     map[string]interface{} `json:"extra,omitempty"`
	RequestID string                 `json:"request_id,omitempty"`
}

func (e *Error) Error() string {
	s := fmt.Sprintf("[%d/%d] %s", e.Status, e.Code, e.Message)
	for k, v := range e.Extra {
		s += fmt.Sprintf(" %v=%v", k, v)
	}

	if e.RequestID != "" {
		s += fmt.Sprintf(" id=%s", e.RequestID)
	}

	return s
}

func IsErrorCodes(err error, codes ...int) bool {
	var e *Error
	if errors.As(err, &e) {
		for _, code := range codes {
			if e.Code == code {
				return true
			}
		}
	}

	return false
}

func createError(status, code int, message string) error {
	return &Error{
		Status:  status,
		Code:    code,
		Message: message,
	}
}

// errWithRequestID wrap err with request id
type errWithRequestID struct {
	err       error
	requestID string
}

func (e *errWithRequestID) Unwrap() error {
	return e.err
}

func (e *errWithRequestID) Error() string {
	return fmt.Sprintf("%v id=%s", e.err, e.requestID)
}

func WrapErrWithRequestID(err error, id string) error {
	if e, ok := err.(*Error); ok {
		e.RequestID = id
		return e
	}

	return &errWithRequestID{err: err, requestID: id}
}

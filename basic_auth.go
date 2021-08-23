package rc

import (
	"encoding/base64"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type BasicAuth struct {
	Key    string
	Secret string
}

func (auth *BasicAuth) GenToken() string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", auth.Key, auth.Secret)))
}

func (auth *BasicAuth) Verify(resp *resty.Response) error {
	return nil
}

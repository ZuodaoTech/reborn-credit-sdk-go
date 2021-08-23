package rc

import "github.com/go-resty/resty/v2"

const ()

type Auth interface {
	GenToken() string
}

type Verifier interface {
	Verify(resp *resty.Response) error
}

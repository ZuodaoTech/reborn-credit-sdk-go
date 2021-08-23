package rc

import (
	"os"
)

const (
	DefaultApiHost = "https://reborn-credits-api.firesbox.com"
)

func UseApiHost(host string) {
	httpClient.HostURL = host
}

func init() {

	if host, ok := os.LookupEnv("REBORN_CREDIT_SDK_API_HOST"); ok && host != "" {
		UseApiHost(host)
	}

}

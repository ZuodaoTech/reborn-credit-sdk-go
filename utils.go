package rc

import (
	"github.com/gofrs/uuid"
)

func newUUID() string {
	return uuid.Must(uuid.NewV4()).String()
}

func RandomTraceID() string {
	return newUUID()
}

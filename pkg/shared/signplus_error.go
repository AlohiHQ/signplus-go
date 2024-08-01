package shared

import (
	"github.com/alohihq/signplus-go/internal/clients/rest/httptransport"
)

type SignplusError struct {
	Err      error
	Body     []byte
	Metadata SignplusErrorMetadata
}

type SignplusErrorMetadata struct {
	Headers    map[string]string
	StatusCode int
}

func NewSignplusError[T any](transportError *httptransport.ErrorResponse[T]) *SignplusError {
	return &SignplusError{
		Err:  transportError.GetError(),
		Body: transportError.GetBody(),
		Metadata: SignplusErrorMetadata{
			StatusCode: transportError.GetStatusCode(),
			Headers:    transportError.GetHeaders(),
		},
	}
}

func (e *SignplusError) Error() string {
	return e.Err.Error()
}

package sharedmodels

import (
	"github.com/alohihq/signplus-go/internal/clients/rest/httptransport"
	"net/http"
)

// SignplusError wraps API errors with detailed metadata including status code, headers, and raw response.
// It implements the error interface and provides structured access to error information.
type SignplusError[T any] struct {
	Err      error
	Data     *T
	Body     []byte
	Raw      *http.Response
	Metadata SignplusErrorMetadata
}

// SignplusErrorMetadata contains HTTP metadata associated with an error response.
type SignplusErrorMetadata struct {
	Headers    map[string]string
	StatusCode int
}

// NewSignplusError creates a new SignplusError from an internal transport error.
// It extracts error details, body, status code, and headers into a user-facing error structure.
func NewSignplusError[T any](transportError *httptransport.ErrorResponse[T]) *SignplusError[T] {
	return &SignplusError[T]{
		Err:  transportError.GetError(),
		Data: transportError.Data,
		Body: transportError.GetBody(),
		Raw:  transportError.Raw,
		Metadata: SignplusErrorMetadata{
			StatusCode: transportError.GetStatusCode(),
			Headers:    transportError.GetHeaders(),
		},
	}
}

// Error implements the error interface, returning the error message string.
func (e *SignplusError[T]) Error() string {
	if e == nil || e.Err == nil {
		return ""
	}
	return e.Err.Error()
}

// Unwrap returns the underlying error, enabling errors.Is and errors.As to traverse the chain.
func (e *SignplusError[T]) Unwrap() error {
	return e.Err
}

// GetData returns the deserialized error response data.
// Returns nil if unmarshaling failed or the response body was empty.
func (e *SignplusError[T]) GetData() *T {
	return e.Data
}

// GetBody returns the raw response body bytes from the error response.
// Returns nil if no response body was received.
func (e *SignplusError[T]) GetBody() []byte {
	return e.Body
}

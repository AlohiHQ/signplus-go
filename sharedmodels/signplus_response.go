package sharedmodels

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/clients/rest/httptransport"
	"net/http"
)

// SignplusResponse is the user-facing wrapper for API responses.
// It contains the deserialized data, raw HTTP response, and metadata like headers and status code.
type SignplusResponse[T any] struct {
	Data     T
	Raw      *http.Response
	Metadata SignplusResponseMetadata
}

// SignplusResponseMetadata contains HTTP metadata from the API response.
// Includes status code and headers for inspection and debugging.
type SignplusResponseMetadata struct {
	Headers    map[string]string
	StatusCode int
}

// NewSignplusResponse creates a new response wrapper from an internal transport response.
// Extracts data and metadata into a user-facing structure.
func NewSignplusResponse[T any](resp *httptransport.Response[T]) *SignplusResponse[T] {
	return &SignplusResponse[T]{
		Data: resp.Data,
		Raw:  resp.Raw,
		Metadata: SignplusResponseMetadata{
			StatusCode: resp.StatusCode,
			Headers:    resp.Headers,
		},
	}
}

// GetData returns the deserialized response data.
func (r *SignplusResponse[T]) GetData() T {
	return r.Data
}

// String returns a JSON representation of the response for debugging.
// Returns an error message if JSON marshaling fails.
func (r SignplusResponse[T]) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: SignplusResponse to string"
	}
	return string(jsonData)
}

package shared

import (
	"github.com/alohihq/signplus-go/internal/clients/rest/httptransport"
)

type SignplusResponse[T any] struct {
	Data     T
	Metadata SignplusResponseMetadata
}

type SignplusResponseMetadata struct {
	Headers    map[string]string
	StatusCode int
}

func NewSignplusResponse[T any](resp *httptransport.Response[T]) *SignplusResponse[T] {
	return &SignplusResponse[T]{
		Data: resp.Data,
		Metadata: SignplusResponseMetadata{
			StatusCode: resp.StatusCode,
			Headers:    resp.Headers,
		},
	}
}

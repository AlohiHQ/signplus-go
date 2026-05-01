package documents

import (
	"github.com/alohihq/signplus-go/param"
)

// GetEnvelopeDocumentsRequestParams holds the optional parameters for the API request.
type GetEnvelopeDocumentsRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

package envelopeid

import (
	"github.com/alohihq/signplus-go/param"
)

// GetEnvelopeRequestParams holds the optional parameters for the API request.
type GetEnvelopeRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

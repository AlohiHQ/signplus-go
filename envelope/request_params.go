package envelope

import (
	"github.com/alohihq/signplus-go/param"
)

// CreateEnvelopeRequestParams holds the optional parameters for the API request.
type CreateEnvelopeRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

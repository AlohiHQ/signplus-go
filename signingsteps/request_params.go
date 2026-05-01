package signingsteps

import (
	"github.com/alohihq/signplus-go/param"
)

// AddEnvelopeSigningStepsRequestParams holds the optional parameters for the API request.
type AddEnvelopeSigningStepsRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

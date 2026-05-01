package void

import (
	"github.com/alohihq/signplus-go/param"
)

// VoidEnvelopeRequestParams holds the optional parameters for the API request.
type VoidEnvelopeRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

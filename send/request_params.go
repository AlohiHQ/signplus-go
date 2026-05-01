package send

import (
	"github.com/alohihq/signplus-go/param"
)

// SendEnvelopeRequestParams holds the optional parameters for the API request.
type SendEnvelopeRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

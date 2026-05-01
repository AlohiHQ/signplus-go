package duplicate

import (
	"github.com/alohihq/signplus-go/param"
)

// DuplicateEnvelopeRequestParams holds the optional parameters for the API request.
type DuplicateEnvelopeRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

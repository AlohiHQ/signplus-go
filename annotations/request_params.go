package annotations

import (
	"github.com/alohihq/signplus-go/param"
)

// GetEnvelopeAnnotationsRequestParams holds the optional parameters for the API request.
type GetEnvelopeAnnotationsRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

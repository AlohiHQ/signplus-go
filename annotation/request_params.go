package annotation

import (
	"github.com/alohihq/signplus-go/param"
)

// AddEnvelopeAnnotationRequestParams holds the optional parameters for the API request.
type AddEnvelopeAnnotationRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

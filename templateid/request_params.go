package templateid

import (
	"github.com/alohihq/signplus-go/param"
)

// CreateEnvelopeFromTemplateRequestParams holds the optional parameters for the API request.
type CreateEnvelopeFromTemplateRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

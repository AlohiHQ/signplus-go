package setexpirationdate

import (
	"github.com/alohihq/signplus-go/param"
)

// SetEnvelopeExpirationDateRequestParams holds the optional parameters for the API request.
type SetEnvelopeExpirationDateRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

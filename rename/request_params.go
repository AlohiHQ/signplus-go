package rename

import (
	"github.com/alohihq/signplus-go/param"
)

// RenameEnvelopeRequestParams holds the optional parameters for the API request.
type RenameEnvelopeRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

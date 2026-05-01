package document

import (
	"github.com/alohihq/signplus-go/param"
)

// AddEnvelopeDocumentRequestParams holds the optional parameters for the API request.
type AddEnvelopeDocumentRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

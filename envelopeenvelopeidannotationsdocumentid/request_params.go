package envelopeenvelopeidannotationsdocumentid

import (
	"github.com/alohihq/signplus-go/param"
)

// GetEnvelopeDocumentAnnotationsRequestParams holds the optional parameters for the API request.
type GetEnvelopeDocumentAnnotationsRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

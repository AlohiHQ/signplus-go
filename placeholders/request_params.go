package placeholders

import (
	"github.com/alohihq/signplus-go/param"
)

// SetEnvelopeAttachmentsPlaceholdersRequestParams holds the optional parameters for the API request.
type SetEnvelopeAttachmentsPlaceholdersRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

package settings

import (
	"github.com/alohihq/signplus-go/param"
)

// SetEnvelopeAttachmentsSettingsRequestParams holds the optional parameters for the API request.
type SetEnvelopeAttachmentsSettingsRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

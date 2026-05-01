package templatetemplateidattachmentssettings

import (
	"github.com/alohihq/signplus-go/param"
)

// SetTemplateAttachmentsSettingsRequestParams holds the optional parameters for the API request.
type SetTemplateAttachmentsSettingsRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

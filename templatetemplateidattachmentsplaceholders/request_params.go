package templatetemplateidattachmentsplaceholders

import (
	"github.com/alohihq/signplus-go/param"
)

// SetTemplateAttachmentsPlaceholdersRequestParams holds the optional parameters for the API request.
type SetTemplateAttachmentsPlaceholdersRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

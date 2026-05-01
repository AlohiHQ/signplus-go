package templatetemplateid

import (
	"github.com/alohihq/signplus-go/param"
)

// GetTemplateRequestParams holds the optional parameters for the API request.
type GetTemplateRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

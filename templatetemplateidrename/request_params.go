package templatetemplateidrename

import (
	"github.com/alohihq/signplus-go/param"
)

// RenameTemplateRequestParams holds the optional parameters for the API request.
type RenameTemplateRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

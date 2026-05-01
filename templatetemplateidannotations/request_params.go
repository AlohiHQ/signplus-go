package templatetemplateidannotations

import (
	"github.com/alohihq/signplus-go/param"
)

// GetTemplateAnnotationsRequestParams holds the optional parameters for the API request.
type GetTemplateAnnotationsRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

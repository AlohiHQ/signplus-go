package templatetemplateidannotationsdocumentid

import (
	"github.com/alohihq/signplus-go/param"
)

// GetDocumentTemplateAnnotationsRequestParams holds the optional parameters for the API request.
type GetDocumentTemplateAnnotationsRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

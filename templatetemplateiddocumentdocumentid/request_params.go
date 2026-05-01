package templatetemplateiddocumentdocumentid

import (
	"github.com/alohihq/signplus-go/param"
)

// GetTemplateDocumentRequestParams holds the optional parameters for the API request.
type GetTemplateDocumentRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

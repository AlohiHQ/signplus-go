package templatetemplateiddocuments

import (
	"github.com/alohihq/signplus-go/param"
)

// GetTemplateDocumentsRequestParams holds the optional parameters for the API request.
type GetTemplateDocumentsRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

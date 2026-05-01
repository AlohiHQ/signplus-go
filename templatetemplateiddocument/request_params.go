package templatetemplateiddocument

import (
	"github.com/alohihq/signplus-go/param"
)

// AddTemplateDocumentRequestParams holds the optional parameters for the API request.
type AddTemplateDocumentRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

package templatetemplateidduplicate

import (
	"github.com/alohihq/signplus-go/param"
)

// DuplicateTemplateRequestParams holds the optional parameters for the API request.
type DuplicateTemplateRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

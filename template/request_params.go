package template

import (
	"github.com/alohihq/signplus-go/param"
)

// CreateTemplateRequestParams holds the optional parameters for the API request.
type CreateTemplateRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

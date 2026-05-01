package templates

import (
	"github.com/alohihq/signplus-go/param"
)

// ListTemplatesRequestParams holds the optional parameters for the API request.
type ListTemplatesRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

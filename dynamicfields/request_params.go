package dynamicfields

import (
	"github.com/alohihq/signplus-go/param"
)

// SetEnvelopeDynamicFieldsRequestParams holds the optional parameters for the API request.
type SetEnvelopeDynamicFieldsRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

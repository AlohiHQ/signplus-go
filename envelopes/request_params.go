package envelopes

import (
	"github.com/alohihq/signplus-go/param"
)

// ListEnvelopesRequestParams holds the optional parameters for the API request.
type ListEnvelopesRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

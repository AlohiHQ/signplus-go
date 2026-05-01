package setcomment

import (
	"github.com/alohihq/signplus-go/param"
)

// SetEnvelopeCommentRequestParams holds the optional parameters for the API request.
type SetEnvelopeCommentRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

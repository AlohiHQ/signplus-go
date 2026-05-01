package templatetemplateidsetcomment

import (
	"github.com/alohihq/signplus-go/param"
)

// SetTemplateCommentRequestParams holds the optional parameters for the API request.
type SetTemplateCommentRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

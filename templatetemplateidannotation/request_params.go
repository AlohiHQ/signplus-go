package templatetemplateidannotation

import (
	"github.com/alohihq/signplus-go/param"
)

// AddTemplateAnnotationRequestParams holds the optional parameters for the API request.
type AddTemplateAnnotationRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

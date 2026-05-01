package templatetemplateidsigningsteps

import (
	"github.com/alohihq/signplus-go/param"
)

// AddTemplateSigningStepsRequestParams holds the optional parameters for the API request.
type AddTemplateSigningStepsRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

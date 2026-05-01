package webhook

import (
	"github.com/alohihq/signplus-go/param"
)

// CreateWebhookRequestParams holds the optional parameters for the API request.
type CreateWebhookRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

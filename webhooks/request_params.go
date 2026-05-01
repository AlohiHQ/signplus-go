package webhooks

import (
	"github.com/alohihq/signplus-go/param"
)

// ListWebhooksRequestParams holds the optional parameters for the API request.
type ListWebhooksRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

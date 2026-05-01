package setnotification

import (
	"github.com/alohihq/signplus-go/param"
)

// SetEnvelopeNotificationRequestParams holds the optional parameters for the API request.
type SetEnvelopeNotificationRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

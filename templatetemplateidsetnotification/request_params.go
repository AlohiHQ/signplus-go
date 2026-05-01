package templatetemplateidsetnotification

import (
	"github.com/alohihq/signplus-go/param"
)

// SetTemplateNotificationRequestParams holds the optional parameters for the API request.
type SetTemplateNotificationRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

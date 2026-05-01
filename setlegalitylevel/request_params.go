package setlegalitylevel

import (
	"github.com/alohihq/signplus-go/param"
)

// SetEnvelopeLegalityLevelRequestParams holds the optional parameters for the API request.
type SetEnvelopeLegalityLevelRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

package fileid

import (
	"github.com/alohihq/signplus-go/param"
)

// GetAttachmentFileRequestParams holds the optional parameters for the API request.
type GetAttachmentFileRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

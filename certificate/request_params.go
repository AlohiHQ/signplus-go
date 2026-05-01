package certificate

import (
	"github.com/alohihq/signplus-go/param"
)

// DownloadEnvelopeCertificateRequestParams holds the optional parameters for the API request.
type DownloadEnvelopeCertificateRequestParams struct {
	Accept *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

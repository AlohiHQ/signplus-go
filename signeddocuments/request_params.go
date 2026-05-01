package signeddocuments

import (
	"github.com/alohihq/signplus-go/param"
)

// DownloadEnvelopeSignedDocumentsRequestParams holds the optional parameters for the API request.
type DownloadEnvelopeSignedDocumentsRequestParams struct {
	CertificateOfCompletion *param.Nullable[string] `explode:"true" serializationStyle:"form" queryParam:"certificate_of_completion"`
	Accept                  *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

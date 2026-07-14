package signplus1

// DownloadEnvelopeSignedDocumentsRequestParams holds the optional parameters for the API request.
type DownloadEnvelopeSignedDocumentsRequestParams struct {
	CertificateOfCompletion *bool `explode:"true" serializationStyle:"form" queryParam:"certificate_of_completion"`
}

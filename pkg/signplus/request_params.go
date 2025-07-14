package signplus

type DownloadEnvelopeSignedDocumentsRequestParams struct {
	CertificateOfCompletion *bool `explode:"true" serializationStyle:"form" queryParam:"certificate_of_completion"`
}

func (params *DownloadEnvelopeSignedDocumentsRequestParams) SetCertificateOfCompletion(certificateOfCompletion bool) {
	params.CertificateOfCompletion = &certificateOfCompletion
}

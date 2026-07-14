package signplus1

// Status of the envelope
type EnvelopeStatus string

const (
	EnvelopeStatusDraft      EnvelopeStatus = "DRAFT"
	EnvelopeStatusInProgress EnvelopeStatus = "IN_PROGRESS"
	EnvelopeStatusCompleted  EnvelopeStatus = "COMPLETED"
	EnvelopeStatusExpired    EnvelopeStatus = "EXPIRED"
	EnvelopeStatusDeclined   EnvelopeStatus = "DECLINED"
	EnvelopeStatusVoided     EnvelopeStatus = "VOIDED"
	EnvelopeStatusPending    EnvelopeStatus = "PENDING"
)

package signplus

// Status of the envelope
type EnvelopeStatus string

const (
	ENVELOPE_STATUS_DRAFT       EnvelopeStatus = "DRAFT"
	ENVELOPE_STATUS_IN_PROGRESS EnvelopeStatus = "IN_PROGRESS"
	ENVELOPE_STATUS_COMPLETED   EnvelopeStatus = "COMPLETED"
	ENVELOPE_STATUS_EXPIRED     EnvelopeStatus = "EXPIRED"
	ENVELOPE_STATUS_DECLINED    EnvelopeStatus = "DECLINED"
	ENVELOPE_STATUS_VOIDED      EnvelopeStatus = "VOIDED"
	ENVELOPE_STATUS_PENDING     EnvelopeStatus = "PENDING"
)

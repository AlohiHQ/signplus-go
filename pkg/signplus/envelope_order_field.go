package signplus

// Field to order envelopes by
type EnvelopeOrderField string

const (
	ENVELOPE_ORDER_FIELD_CREATION_DATE        EnvelopeOrderField = "CREATION_DATE"
	ENVELOPE_ORDER_FIELD_MODIFICATION_DATE    EnvelopeOrderField = "MODIFICATION_DATE"
	ENVELOPE_ORDER_FIELD_NAME                 EnvelopeOrderField = "NAME"
	ENVELOPE_ORDER_FIELD_STATUS               EnvelopeOrderField = "STATUS"
	ENVELOPE_ORDER_FIELD_LAST_DOCUMENT_CHANGE EnvelopeOrderField = "LAST_DOCUMENT_CHANGE"
)

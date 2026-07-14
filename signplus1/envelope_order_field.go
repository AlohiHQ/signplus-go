package signplus1

// Field to order envelopes by
type EnvelopeOrderField string

const (
	EnvelopeOrderFieldCreationDate       EnvelopeOrderField = "CREATION_DATE"
	EnvelopeOrderFieldModificationDate   EnvelopeOrderField = "MODIFICATION_DATE"
	EnvelopeOrderFieldName               EnvelopeOrderField = "NAME"
	EnvelopeOrderFieldStatus             EnvelopeOrderField = "STATUS"
	EnvelopeOrderFieldLastDocumentChange EnvelopeOrderField = "LAST_DOCUMENT_CHANGE"
)

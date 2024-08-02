package signplus

// Legal level of the envelope (SES is Simple Electronic Signature, QES_EIDAS is Qualified Electronic Signature, QES_ZERTES is Qualified Electronic Signature with Zertes)
type EnvelopeLegalityLevel string

const (
	ENVELOPE_LEGALITY_LEVEL_SES        EnvelopeLegalityLevel = "SES"
	ENVELOPE_LEGALITY_LEVEL_QES_EIDAS  EnvelopeLegalityLevel = "QES_EIDAS"
	ENVELOPE_LEGALITY_LEVEL_QES_ZERTES EnvelopeLegalityLevel = "QES_ZERTES"
)

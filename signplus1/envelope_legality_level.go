package signplus1

// Legal level of the envelope (SES is Simple Electronic Signature, QES_EIDAS is Qualified Electronic Signature, QES_ZERTES is Qualified Electronic Signature with Zertes)
type EnvelopeLegalityLevel string

const (
	EnvelopeLegalityLevelSes       EnvelopeLegalityLevel = "SES"
	EnvelopeLegalityLevelQesEidas  EnvelopeLegalityLevel = "QES_EIDAS"
	EnvelopeLegalityLevelQesZertes EnvelopeLegalityLevel = "QES_ZERTES"
)

package signplus

// Type of verification the recipient must complete before accessing the envelope.
//
// - `PASSCODE`: requires a code to be entered.
// - `SMS`: sends a code via SMS.
// - `ID_VERIFICATION`: prompts the recipient to complete an automated ID and selfie check.
type RecipientVerificationType string

const (
	RECIPIENT_VERIFICATION_TYPE_SMS             RecipientVerificationType = "SMS"
	RECIPIENT_VERIFICATION_TYPE_PASSCODE        RecipientVerificationType = "PASSCODE"
	RECIPIENT_VERIFICATION_TYPE_ID_VERIFICATION RecipientVerificationType = "ID_VERIFICATION"
)

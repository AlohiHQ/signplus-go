package signplus1

// Type of verification the recipient must complete before accessing the envelope.
//
// - `PASSCODE`: requires a code to be entered.
// - `SMS`: sends a code via SMS.
// - `ID_VERIFICATION`: prompts the recipient to complete an automated ID and selfie check.
type RecipientVerificationType string

const (
	RecipientVerificationTypeSms            RecipientVerificationType = "SMS"
	RecipientVerificationTypePasscode       RecipientVerificationType = "PASSCODE"
	RecipientVerificationTypeIdVerification RecipientVerificationType = "ID_VERIFICATION"
)

package signplus

// Type of signature verification (SMS sends a code via SMS, PASSCODE requires a code to be entered)
type RecipientVerificationType string

const (
	RECIPIENT_VERIFICATION_TYPE_SMS      RecipientVerificationType = "SMS"
	RECIPIENT_VERIFICATION_TYPE_PASSCODE RecipientVerificationType = "PASSCODE"
)

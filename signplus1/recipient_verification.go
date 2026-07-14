package signplus1

import "encoding/json"

type RecipientVerification struct {
	// Type of verification the recipient must complete before accessing the envelope.
	//
	// - `PASSCODE`: requires a code to be entered.
	// - `SMS`: sends a code via SMS.
	// - `ID_VERIFICATION`: prompts the recipient to complete an automated ID and selfie check.
	Type *RecipientVerificationType `json:"type,omitempty" xml:"type,omitempty"`
	// Required for `PASSCODE` and `SMS` verification.
	//
	// - `PASSCODE`: code required by the recipient to sign the document.
	// - `SMS`: recipient's phone number.
	// - `ID_VERIFICATION`: leave empty.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (r RecipientVerification) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RecipientVerification to string"
	}
	return string(jsonData)
}

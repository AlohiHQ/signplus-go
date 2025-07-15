package signplus

import "encoding/json"

type RecipientVerification struct {
	// Type of verification the recipient must complete before accessing the envelope.
	//
	// - `PASSCODE`: requires a code to be entered.
	// - `SMS`: sends a code via SMS.
	// - `ID_VERIFICATION`: prompts the recipient to complete an automated ID and selfie check.
	Type_ *RecipientVerificationType `json:"type,omitempty"`
	// Required for `PASSCODE` and `SMS` verification.
	//
	// - `PASSCODE`: code required by the recipient to sign the document.
	// - `SMS`: recipient's phone number.
	// - `ID_VERIFICATION`: leave empty.
	Value *string `json:"value,omitempty"`
}

func (r *RecipientVerification) GetType_() *RecipientVerificationType {
	if r == nil {
		return nil
	}
	return r.Type_
}

func (r *RecipientVerification) SetType_(type_ RecipientVerificationType) {
	r.Type_ = &type_
}

func (r *RecipientVerification) GetValue() *string {
	if r == nil {
		return nil
	}
	return r.Value
}

func (r *RecipientVerification) SetValue(value string) {
	r.Value = &value
}

func (r RecipientVerification) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RecipientVerification to string"
	}
	return string(jsonData)
}

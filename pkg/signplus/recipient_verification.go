package signplus

type RecipientVerification struct {
	// Type of signature verification (SMS sends a code via SMS, PASSCODE requires a code to be entered)
	Type_ *RecipientVerificationType `json:"type,omitempty"`
	Value *string                    `json:"value,omitempty"`
}

func (r *RecipientVerification) SetType_(type_ RecipientVerificationType) {
	r.Type_ = &type_
}

func (r *RecipientVerification) GetType_() *RecipientVerificationType {
	if r == nil {
		return nil
	}
	return r.Type_
}

func (r *RecipientVerification) SetValue(value string) {
	r.Value = &value
}

func (r *RecipientVerification) GetValue() *string {
	if r == nil {
		return nil
	}
	return r.Value
}

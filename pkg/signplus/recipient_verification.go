package signplus

import (
	"encoding/json"
)

type RecipientVerification struct {
	// Type of signature verification (SMS sends a code via SMS, PASSCODE requires a code to be entered)
	Type_   *RecipientVerificationType `json:"type,omitempty"`
	Value   *string                    `json:"value,omitempty"`
	touched map[string]bool
}

func (r *RecipientVerification) GetType_() *RecipientVerificationType {
	if r == nil {
		return nil
	}
	return r.Type_
}

func (r *RecipientVerification) SetType_(type_ RecipientVerificationType) {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Type_"] = true
	r.Type_ = &type_
}

func (r *RecipientVerification) SetType_Nil() {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Type_"] = true
	r.Type_ = nil
}

func (r *RecipientVerification) GetValue() *string {
	if r == nil {
		return nil
	}
	return r.Value
}

func (r *RecipientVerification) SetValue(value string) {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Value"] = true
	r.Value = &value
}

func (r *RecipientVerification) SetValueNil() {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Value"] = true
	r.Value = nil
}
func (r RecipientVerification) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if r.touched["Type_"] && r.Type_ == nil {
		data["type"] = nil
	} else if r.Type_ != nil {
		data["type"] = r.Type_
	}

	if r.touched["Value"] && r.Value == nil {
		data["value"] = nil
	} else if r.Value != nil {
		data["value"] = r.Value
	}

	return json.Marshal(data)
}

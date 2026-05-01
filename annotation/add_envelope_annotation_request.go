package annotation

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type AddEnvelopeAnnotationRequest struct {
	DocumentID  *param.Nullable[string]                                `json:"document_id,omitempty" xml:"document_id,omitempty"`
	Page        *param.Nullable[string]                                `json:"page,omitempty" xml:"page,omitempty"`
	X           *param.Nullable[string]                                `json:"x,omitempty" xml:"x,omitempty"`
	Y           *param.Nullable[string]                                `json:"y,omitempty" xml:"y,omitempty"`
	Width       *param.Nullable[string]                                `json:"width,omitempty" xml:"width,omitempty"`
	Height      *param.Nullable[string]                                `json:"height,omitempty" xml:"height,omitempty"`
	Type        *param.Nullable[string]                                `json:"type,omitempty" xml:"type,omitempty"`
	RecipientID *param.Nullable[string]                                `json:"recipient_id,omitempty" xml:"recipient_id,omitempty"`
	Required    *param.Nullable[string]                                `json:"required,omitempty" xml:"required,omitempty"`
	Signature   *param.Nullable[AddEnvelopeAnnotationRequestSignature] `json:"signature,omitempty" xml:"signature,omitempty"`
	Initials    *param.Nullable[AddEnvelopeAnnotationRequestInitials]  `json:"initials,omitempty" xml:"initials,omitempty"`
	Text        *param.Nullable[AddEnvelopeAnnotationRequestText]      `json:"text,omitempty" xml:"text,omitempty"`
	Datetime    *param.Nullable[AddEnvelopeAnnotationRequestDatetime]  `json:"datetime,omitempty" xml:"datetime,omitempty"`
	Checkbox    *param.Nullable[AddEnvelopeAnnotationRequestCheckbox]  `json:"checkbox,omitempty" xml:"checkbox,omitempty"`
}

func (a AddEnvelopeAnnotationRequest) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddEnvelopeAnnotationRequest to string"
	}
	return string(jsonData)
}

func (a *AddEnvelopeAnnotationRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, a)
}

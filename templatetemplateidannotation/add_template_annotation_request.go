package templatetemplateidannotation

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type AddTemplateAnnotationRequest struct {
	DocumentID  *param.Nullable[string]                                `json:"document_id,omitempty" xml:"document_id,omitempty"`
	Page        *param.Nullable[float64]                               `json:"page,omitempty" xml:"page,omitempty"`
	X           *param.Nullable[float64]                               `json:"x,omitempty" xml:"x,omitempty"`
	Y           *param.Nullable[float64]                               `json:"y,omitempty" xml:"y,omitempty"`
	Width       *param.Nullable[float64]                               `json:"width,omitempty" xml:"width,omitempty"`
	Height      *param.Nullable[float64]                               `json:"height,omitempty" xml:"height,omitempty"`
	Type        *param.Nullable[string]                                `json:"type,omitempty" xml:"type,omitempty"`
	RecipientID *param.Nullable[string]                                `json:"recipient_id,omitempty" xml:"recipient_id,omitempty"`
	Required    *param.Nullable[bool]                                  `json:"required,omitempty" xml:"required,omitempty"`
	Signature   *param.Nullable[AddTemplateAnnotationRequestSignature] `json:"signature,omitempty" xml:"signature,omitempty"`
	Initials    *param.Nullable[AddTemplateAnnotationRequestInitials]  `json:"initials,omitempty" xml:"initials,omitempty"`
	Text        *param.Nullable[AddTemplateAnnotationRequestText]      `json:"text,omitempty" xml:"text,omitempty"`
	Datetime    *param.Nullable[AddTemplateAnnotationRequestDatetime]  `json:"datetime,omitempty" xml:"datetime,omitempty"`
	Checkbox    *param.Nullable[AddTemplateAnnotationRequestCheckbox]  `json:"checkbox,omitempty" xml:"checkbox,omitempty"`
}

func (a AddTemplateAnnotationRequest) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddTemplateAnnotationRequest to string"
	}
	return string(jsonData)
}

func (a *AddTemplateAnnotationRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, a)
}

package signplus1

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
)

type AddAnnotationRequest struct {
	// ID of the recipient
	RecipientID *string `json:"recipient_id,omitempty" xml:"recipient_id,omitempty"`
	// ID of the document
	DocumentID string `json:"document_id" xml:"document_id" required:"true"`
	// Page number where the annotation is placed
	Page int64 `json:"page" xml:"page" required:"true"`
	// X coordinate of the annotation (in % of the page width from 0 to 100) from the top left corner
	X float64 `json:"x" xml:"x" required:"true"`
	// Y coordinate of the annotation (in % of the page height from 0 to 100) from the top left corner
	Y float64 `json:"y" xml:"y" required:"true"`
	// Width of the annotation (in % of the page width from 0 to 100)
	Width float64 `json:"width" xml:"width" required:"true"`
	// Height of the annotation (in % of the page height from 0 to 100)
	Height   float64 `json:"height" xml:"height" required:"true"`
	Required *bool   `json:"required,omitempty" xml:"required,omitempty"`
	// Type of the annotation
	Type AnnotationType `json:"type" xml:"type" required:"true"`
	// Signature annotation (null if annotation is not a signature)
	Signature *AnnotationSignature `json:"signature,omitempty" xml:"signature,omitempty"`
	// Initials annotation (null if annotation is not initials)
	Initials *AnnotationInitials `json:"initials,omitempty" xml:"initials,omitempty"`
	// Text annotation (null if annotation is not a text)
	Text *AnnotationText `json:"text,omitempty" xml:"text,omitempty"`
	// Date annotation (null if annotation is not a date)
	Datetime *AnnotationDateTime `json:"datetime,omitempty" xml:"datetime,omitempty"`
	// Checkbox annotation (null if annotation is not a checkbox)
	Checkbox *AnnotationCheckbox `json:"checkbox,omitempty" xml:"checkbox,omitempty"`
}

func (a AddAnnotationRequest) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddAnnotationRequest to string"
	}
	return string(jsonData)
}

func (a *AddAnnotationRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AddAnnotationRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AddAnnotationRequest(tmp)
	return nil
}

package signplus1

import "encoding/json"

type Annotation struct {
	// Unique identifier of the annotation
	ID *string `json:"id,omitempty" xml:"id,omitempty"`
	// ID of the recipient
	RecipientID *string `json:"recipient_id,omitempty" xml:"recipient_id,omitempty"`
	// ID of the document
	DocumentID *string `json:"document_id,omitempty" xml:"document_id,omitempty"`
	// Page number where the annotation is placed
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// X coordinate of the annotation (in % of the page width from 0 to 100) from the top left corner
	X *float64 `json:"x,omitempty" xml:"x,omitempty"`
	// Y coordinate of the annotation (in % of the page height from 0 to 100) from the top left corner
	Y *float64 `json:"y,omitempty" xml:"y,omitempty"`
	// Width of the annotation (in % of the page width from 0 to 100)
	Width *float64 `json:"width,omitempty" xml:"width,omitempty"`
	// Height of the annotation (in % of the page height from 0 to 100)
	Height *float64 `json:"height,omitempty" xml:"height,omitempty"`
	// Whether the annotation is required
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// Type of the annotation
	Type *AnnotationType `json:"type,omitempty" xml:"type,omitempty"`
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

func (a Annotation) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: Annotation to string"
	}
	return string(jsonData)
}

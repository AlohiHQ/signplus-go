package signplus

import (
	"encoding/json"
)

type Annotation struct {
	// Unique identifier of the annotation
	Id *string `json:"id,omitempty"`
	// ID of the recipient
	RecipientId *string `json:"recipient_id,omitempty"`
	// ID of the document
	DocumentId *string `json:"document_id,omitempty"`
	// Page number where the annotation is placed
	Page *int64 `json:"page,omitempty"`
	// X coordinate of the annotation (in % of the page width from 0 to 100) from the top left corner
	X *float64 `json:"x,omitempty"`
	// Y coordinate of the annotation (in % of the page height from 0 to 100) from the top left corner
	Y *float64 `json:"y,omitempty"`
	// Width of the annotation (in % of the page width from 0 to 100)
	Width *float64 `json:"width,omitempty"`
	// Height of the annotation (in % of the page height from 0 to 100)
	Height *float64 `json:"height,omitempty"`
	// Whether the annotation is required
	Required *bool `json:"required,omitempty"`
	// Type of the annotation
	Type_ *AnnotationType `json:"type,omitempty"`
	// Signature annotation (null if annotation is not a signature)
	Signature *AnnotationSignature `json:"signature,omitempty"`
	// Initials annotation (null if annotation is not initials)
	Initials *AnnotationInitials `json:"initials,omitempty"`
	// Text annotation (null if annotation is not a text)
	Text *AnnotationText `json:"text,omitempty"`
	// Date annotation (null if annotation is not a date)
	Datetime *AnnotationDateTime `json:"datetime,omitempty"`
	// Checkbox annotation (null if annotation is not a checkbox)
	Checkbox *AnnotationCheckbox `json:"checkbox,omitempty"`
	touched  map[string]bool
}

func (a *Annotation) GetId() *string {
	if a == nil {
		return nil
	}
	return a.Id
}

func (a *Annotation) SetId(id string) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Id"] = true
	a.Id = &id
}

func (a *Annotation) SetIdNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Id"] = true
	a.Id = nil
}

func (a *Annotation) GetRecipientId() *string {
	if a == nil {
		return nil
	}
	return a.RecipientId
}

func (a *Annotation) SetRecipientId(recipientId string) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["RecipientId"] = true
	a.RecipientId = &recipientId
}

func (a *Annotation) SetRecipientIdNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["RecipientId"] = true
	a.RecipientId = nil
}

func (a *Annotation) GetDocumentId() *string {
	if a == nil {
		return nil
	}
	return a.DocumentId
}

func (a *Annotation) SetDocumentId(documentId string) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["DocumentId"] = true
	a.DocumentId = &documentId
}

func (a *Annotation) SetDocumentIdNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["DocumentId"] = true
	a.DocumentId = nil
}

func (a *Annotation) GetPage() *int64 {
	if a == nil {
		return nil
	}
	return a.Page
}

func (a *Annotation) SetPage(page int64) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Page"] = true
	a.Page = &page
}

func (a *Annotation) SetPageNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Page"] = true
	a.Page = nil
}

func (a *Annotation) GetX() *float64 {
	if a == nil {
		return nil
	}
	return a.X
}

func (a *Annotation) SetX(x float64) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["X"] = true
	a.X = &x
}

func (a *Annotation) SetXNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["X"] = true
	a.X = nil
}

func (a *Annotation) GetY() *float64 {
	if a == nil {
		return nil
	}
	return a.Y
}

func (a *Annotation) SetY(y float64) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Y"] = true
	a.Y = &y
}

func (a *Annotation) SetYNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Y"] = true
	a.Y = nil
}

func (a *Annotation) GetWidth() *float64 {
	if a == nil {
		return nil
	}
	return a.Width
}

func (a *Annotation) SetWidth(width float64) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Width"] = true
	a.Width = &width
}

func (a *Annotation) SetWidthNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Width"] = true
	a.Width = nil
}

func (a *Annotation) GetHeight() *float64 {
	if a == nil {
		return nil
	}
	return a.Height
}

func (a *Annotation) SetHeight(height float64) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Height"] = true
	a.Height = &height
}

func (a *Annotation) SetHeightNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Height"] = true
	a.Height = nil
}

func (a *Annotation) GetRequired() *bool {
	if a == nil {
		return nil
	}
	return a.Required
}

func (a *Annotation) SetRequired(required bool) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Required"] = true
	a.Required = &required
}

func (a *Annotation) SetRequiredNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Required"] = true
	a.Required = nil
}

func (a *Annotation) GetType_() *AnnotationType {
	if a == nil {
		return nil
	}
	return a.Type_
}

func (a *Annotation) SetType_(type_ AnnotationType) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Type_"] = true
	a.Type_ = &type_
}

func (a *Annotation) SetType_Nil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Type_"] = true
	a.Type_ = nil
}

func (a *Annotation) GetSignature() *AnnotationSignature {
	if a == nil {
		return nil
	}
	return a.Signature
}

func (a *Annotation) SetSignature(signature AnnotationSignature) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Signature"] = true
	a.Signature = &signature
}

func (a *Annotation) SetSignatureNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Signature"] = true
	a.Signature = nil
}

func (a *Annotation) GetInitials() *AnnotationInitials {
	if a == nil {
		return nil
	}
	return a.Initials
}

func (a *Annotation) SetInitials(initials AnnotationInitials) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Initials"] = true
	a.Initials = &initials
}

func (a *Annotation) SetInitialsNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Initials"] = true
	a.Initials = nil
}

func (a *Annotation) GetText() *AnnotationText {
	if a == nil {
		return nil
	}
	return a.Text
}

func (a *Annotation) SetText(text AnnotationText) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Text"] = true
	a.Text = &text
}

func (a *Annotation) SetTextNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Text"] = true
	a.Text = nil
}

func (a *Annotation) GetDatetime() *AnnotationDateTime {
	if a == nil {
		return nil
	}
	return a.Datetime
}

func (a *Annotation) SetDatetime(datetime AnnotationDateTime) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Datetime"] = true
	a.Datetime = &datetime
}

func (a *Annotation) SetDatetimeNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Datetime"] = true
	a.Datetime = nil
}

func (a *Annotation) GetCheckbox() *AnnotationCheckbox {
	if a == nil {
		return nil
	}
	return a.Checkbox
}

func (a *Annotation) SetCheckbox(checkbox AnnotationCheckbox) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Checkbox"] = true
	a.Checkbox = &checkbox
}

func (a *Annotation) SetCheckboxNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Checkbox"] = true
	a.Checkbox = nil
}

func (a Annotation) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if a.touched["Id"] && a.Id == nil {
		data["id"] = nil
	} else if a.Id != nil {
		data["id"] = a.Id
	}

	if a.touched["RecipientId"] && a.RecipientId == nil {
		data["recipient_id"] = nil
	} else if a.RecipientId != nil {
		data["recipient_id"] = a.RecipientId
	}

	if a.touched["DocumentId"] && a.DocumentId == nil {
		data["document_id"] = nil
	} else if a.DocumentId != nil {
		data["document_id"] = a.DocumentId
	}

	if a.touched["Page"] && a.Page == nil {
		data["page"] = nil
	} else if a.Page != nil {
		data["page"] = a.Page
	}

	if a.touched["X"] && a.X == nil {
		data["x"] = nil
	} else if a.X != nil {
		data["x"] = a.X
	}

	if a.touched["Y"] && a.Y == nil {
		data["y"] = nil
	} else if a.Y != nil {
		data["y"] = a.Y
	}

	if a.touched["Width"] && a.Width == nil {
		data["width"] = nil
	} else if a.Width != nil {
		data["width"] = a.Width
	}

	if a.touched["Height"] && a.Height == nil {
		data["height"] = nil
	} else if a.Height != nil {
		data["height"] = a.Height
	}

	if a.touched["Required"] && a.Required == nil {
		data["required"] = nil
	} else if a.Required != nil {
		data["required"] = a.Required
	}

	if a.touched["Type_"] && a.Type_ == nil {
		data["type"] = nil
	} else if a.Type_ != nil {
		data["type"] = a.Type_
	}

	if a.touched["Signature"] && a.Signature == nil {
		data["signature"] = nil
	} else if a.Signature != nil {
		data["signature"] = a.Signature
	}

	if a.touched["Initials"] && a.Initials == nil {
		data["initials"] = nil
	} else if a.Initials != nil {
		data["initials"] = a.Initials
	}

	if a.touched["Text"] && a.Text == nil {
		data["text"] = nil
	} else if a.Text != nil {
		data["text"] = a.Text
	}

	if a.touched["Datetime"] && a.Datetime == nil {
		data["datetime"] = nil
	} else if a.Datetime != nil {
		data["datetime"] = a.Datetime
	}

	if a.touched["Checkbox"] && a.Checkbox == nil {
		data["checkbox"] = nil
	} else if a.Checkbox != nil {
		data["checkbox"] = a.Checkbox
	}

	return json.Marshal(data)
}

func (a Annotation) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: Annotation to string"
	}
	return string(jsonData)
}

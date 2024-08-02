package signplus

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
}

func (a *Annotation) SetId(id string) {
	a.Id = &id
}

func (a *Annotation) GetId() *string {
	if a == nil {
		return nil
	}
	return a.Id
}

func (a *Annotation) SetRecipientId(recipientId string) {
	a.RecipientId = &recipientId
}

func (a *Annotation) GetRecipientId() *string {
	if a == nil {
		return nil
	}
	return a.RecipientId
}

func (a *Annotation) SetDocumentId(documentId string) {
	a.DocumentId = &documentId
}

func (a *Annotation) GetDocumentId() *string {
	if a == nil {
		return nil
	}
	return a.DocumentId
}

func (a *Annotation) SetPage(page int64) {
	a.Page = &page
}

func (a *Annotation) GetPage() *int64 {
	if a == nil {
		return nil
	}
	return a.Page
}

func (a *Annotation) SetX(x float64) {
	a.X = &x
}

func (a *Annotation) GetX() *float64 {
	if a == nil {
		return nil
	}
	return a.X
}

func (a *Annotation) SetY(y float64) {
	a.Y = &y
}

func (a *Annotation) GetY() *float64 {
	if a == nil {
		return nil
	}
	return a.Y
}

func (a *Annotation) SetWidth(width float64) {
	a.Width = &width
}

func (a *Annotation) GetWidth() *float64 {
	if a == nil {
		return nil
	}
	return a.Width
}

func (a *Annotation) SetHeight(height float64) {
	a.Height = &height
}

func (a *Annotation) GetHeight() *float64 {
	if a == nil {
		return nil
	}
	return a.Height
}

func (a *Annotation) SetRequired(required bool) {
	a.Required = &required
}

func (a *Annotation) GetRequired() *bool {
	if a == nil {
		return nil
	}
	return a.Required
}

func (a *Annotation) SetType_(type_ AnnotationType) {
	a.Type_ = &type_
}

func (a *Annotation) GetType_() *AnnotationType {
	if a == nil {
		return nil
	}
	return a.Type_
}

func (a *Annotation) SetSignature(signature AnnotationSignature) {
	a.Signature = &signature
}

func (a *Annotation) GetSignature() *AnnotationSignature {
	if a == nil {
		return nil
	}
	return a.Signature
}

func (a *Annotation) SetInitials(initials AnnotationInitials) {
	a.Initials = &initials
}

func (a *Annotation) GetInitials() *AnnotationInitials {
	if a == nil {
		return nil
	}
	return a.Initials
}

func (a *Annotation) SetText(text AnnotationText) {
	a.Text = &text
}

func (a *Annotation) GetText() *AnnotationText {
	if a == nil {
		return nil
	}
	return a.Text
}

func (a *Annotation) SetDatetime(datetime AnnotationDateTime) {
	a.Datetime = &datetime
}

func (a *Annotation) GetDatetime() *AnnotationDateTime {
	if a == nil {
		return nil
	}
	return a.Datetime
}

func (a *Annotation) SetCheckbox(checkbox AnnotationCheckbox) {
	a.Checkbox = &checkbox
}

func (a *Annotation) GetCheckbox() *AnnotationCheckbox {
	if a == nil {
		return nil
	}
	return a.Checkbox
}

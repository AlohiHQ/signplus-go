package signplus

type AddAnnotationRequest struct {
	// ID of the recipient
	RecipientId *string `json:"recipient_id,omitempty"`
	// ID of the document
	DocumentId *string `json:"document_id,omitempty" required:"true"`
	// Page number where the annotation is placed
	Page *int64 `json:"page,omitempty" required:"true"`
	// X coordinate of the annotation (in % of the page width from 0 to 100) from the top left corner
	X *float64 `json:"x,omitempty" required:"true"`
	// Y coordinate of the annotation (in % of the page height from 0 to 100) from the top left corner
	Y *float64 `json:"y,omitempty" required:"true"`
	// Width of the annotation (in % of the page width from 0 to 100)
	Width *float64 `json:"width,omitempty" required:"true"`
	// Height of the annotation (in % of the page height from 0 to 100)
	Height   *float64 `json:"height,omitempty" required:"true"`
	Required *bool    `json:"required,omitempty"`
	// Type of the annotation
	Type_ *AnnotationType `json:"type,omitempty" required:"true"`
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

func (a *AddAnnotationRequest) SetRecipientId(recipientId string) {
	a.RecipientId = &recipientId
}

func (a *AddAnnotationRequest) GetRecipientId() *string {
	if a == nil {
		return nil
	}
	return a.RecipientId
}

func (a *AddAnnotationRequest) SetDocumentId(documentId string) {
	a.DocumentId = &documentId
}

func (a *AddAnnotationRequest) GetDocumentId() *string {
	if a == nil {
		return nil
	}
	return a.DocumentId
}

func (a *AddAnnotationRequest) SetPage(page int64) {
	a.Page = &page
}

func (a *AddAnnotationRequest) GetPage() *int64 {
	if a == nil {
		return nil
	}
	return a.Page
}

func (a *AddAnnotationRequest) SetX(x float64) {
	a.X = &x
}

func (a *AddAnnotationRequest) GetX() *float64 {
	if a == nil {
		return nil
	}
	return a.X
}

func (a *AddAnnotationRequest) SetY(y float64) {
	a.Y = &y
}

func (a *AddAnnotationRequest) GetY() *float64 {
	if a == nil {
		return nil
	}
	return a.Y
}

func (a *AddAnnotationRequest) SetWidth(width float64) {
	a.Width = &width
}

func (a *AddAnnotationRequest) GetWidth() *float64 {
	if a == nil {
		return nil
	}
	return a.Width
}

func (a *AddAnnotationRequest) SetHeight(height float64) {
	a.Height = &height
}

func (a *AddAnnotationRequest) GetHeight() *float64 {
	if a == nil {
		return nil
	}
	return a.Height
}

func (a *AddAnnotationRequest) SetRequired(required bool) {
	a.Required = &required
}

func (a *AddAnnotationRequest) GetRequired() *bool {
	if a == nil {
		return nil
	}
	return a.Required
}

func (a *AddAnnotationRequest) SetType_(type_ AnnotationType) {
	a.Type_ = &type_
}

func (a *AddAnnotationRequest) GetType_() *AnnotationType {
	if a == nil {
		return nil
	}
	return a.Type_
}

func (a *AddAnnotationRequest) SetSignature(signature AnnotationSignature) {
	a.Signature = &signature
}

func (a *AddAnnotationRequest) GetSignature() *AnnotationSignature {
	if a == nil {
		return nil
	}
	return a.Signature
}

func (a *AddAnnotationRequest) SetInitials(initials AnnotationInitials) {
	a.Initials = &initials
}

func (a *AddAnnotationRequest) GetInitials() *AnnotationInitials {
	if a == nil {
		return nil
	}
	return a.Initials
}

func (a *AddAnnotationRequest) SetText(text AnnotationText) {
	a.Text = &text
}

func (a *AddAnnotationRequest) GetText() *AnnotationText {
	if a == nil {
		return nil
	}
	return a.Text
}

func (a *AddAnnotationRequest) SetDatetime(datetime AnnotationDateTime) {
	a.Datetime = &datetime
}

func (a *AddAnnotationRequest) GetDatetime() *AnnotationDateTime {
	if a == nil {
		return nil
	}
	return a.Datetime
}

func (a *AddAnnotationRequest) SetCheckbox(checkbox AnnotationCheckbox) {
	a.Checkbox = &checkbox
}

func (a *AddAnnotationRequest) GetCheckbox() *AnnotationCheckbox {
	if a == nil {
		return nil
	}
	return a.Checkbox
}

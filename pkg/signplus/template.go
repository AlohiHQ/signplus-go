package signplus

type Template struct {
	// Unique identifier of the template
	Id *string `json:"id,omitempty"`
	// Name of the template
	Name *string `json:"name,omitempty"`
	// Comment for the template
	Comment *string `json:"comment,omitempty"`
	// Total number of pages in the template
	Pages *int64 `json:"pages,omitempty"`
	// Legal level of the envelope (SES is Simple Electronic Signature, QES_EIDAS is Qualified Electronic Signature, QES_ZERTES is Qualified Electronic Signature with Zertes)
	LegalityLevel *EnvelopeLegalityLevel `json:"legality_level,omitempty"`
	// Unix timestamp of the creation date
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Unix timestamp of the last modification date
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// Expiration delay added to the current time when an envelope is created from this template
	ExpirationDelay *int64 `json:"expiration_delay,omitempty"`
	// Number of recipients in the envelope
	NumRecipients *int64                `json:"num_recipients,omitempty"`
	SigningSteps  []TemplateSigningStep `json:"signing_steps,omitempty"`
	Documents     []Document            `json:"documents,omitempty"`
	Notification  *EnvelopeNotification `json:"notification,omitempty"`
	// List of dynamic fields
	DynamicFields []string `json:"dynamic_fields,omitempty"`
}

func (t *Template) SetId(id string) {
	t.Id = &id
}

func (t *Template) GetId() *string {
	if t == nil {
		return nil
	}
	return t.Id
}

func (t *Template) SetName(name string) {
	t.Name = &name
}

func (t *Template) GetName() *string {
	if t == nil {
		return nil
	}
	return t.Name
}

func (t *Template) SetComment(comment string) {
	t.Comment = &comment
}

func (t *Template) GetComment() *string {
	if t == nil {
		return nil
	}
	return t.Comment
}

func (t *Template) SetPages(pages int64) {
	t.Pages = &pages
}

func (t *Template) GetPages() *int64 {
	if t == nil {
		return nil
	}
	return t.Pages
}

func (t *Template) SetLegalityLevel(legalityLevel EnvelopeLegalityLevel) {
	t.LegalityLevel = &legalityLevel
}

func (t *Template) GetLegalityLevel() *EnvelopeLegalityLevel {
	if t == nil {
		return nil
	}
	return t.LegalityLevel
}

func (t *Template) SetCreatedAt(createdAt int64) {
	t.CreatedAt = &createdAt
}

func (t *Template) GetCreatedAt() *int64 {
	if t == nil {
		return nil
	}
	return t.CreatedAt
}

func (t *Template) SetUpdatedAt(updatedAt int64) {
	t.UpdatedAt = &updatedAt
}

func (t *Template) GetUpdatedAt() *int64 {
	if t == nil {
		return nil
	}
	return t.UpdatedAt
}

func (t *Template) SetExpirationDelay(expirationDelay int64) {
	t.ExpirationDelay = &expirationDelay
}

func (t *Template) GetExpirationDelay() *int64 {
	if t == nil {
		return nil
	}
	return t.ExpirationDelay
}

func (t *Template) SetNumRecipients(numRecipients int64) {
	t.NumRecipients = &numRecipients
}

func (t *Template) GetNumRecipients() *int64 {
	if t == nil {
		return nil
	}
	return t.NumRecipients
}

func (t *Template) SetSigningSteps(signingSteps []TemplateSigningStep) {
	t.SigningSteps = signingSteps
}

func (t *Template) GetSigningSteps() []TemplateSigningStep {
	if t == nil {
		return nil
	}
	return t.SigningSteps
}

func (t *Template) SetDocuments(documents []Document) {
	t.Documents = documents
}

func (t *Template) GetDocuments() []Document {
	if t == nil {
		return nil
	}
	return t.Documents
}

func (t *Template) SetNotification(notification EnvelopeNotification) {
	t.Notification = &notification
}

func (t *Template) GetNotification() *EnvelopeNotification {
	if t == nil {
		return nil
	}
	return t.Notification
}

func (t *Template) SetDynamicFields(dynamicFields []string) {
	t.DynamicFields = dynamicFields
}

func (t *Template) GetDynamicFields() []string {
	if t == nil {
		return nil
	}
	return t.DynamicFields
}

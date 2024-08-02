package signplus

type Envelope struct {
	// Unique identifier of the envelope
	Id *string `json:"id,omitempty"`
	// Name of the envelope
	Name *string `json:"name,omitempty"`
	// Comment for the envelope
	Comment *string `json:"comment,omitempty"`
	// Total number of pages in the envelope
	Pages *int64 `json:"pages,omitempty"`
	// Flow type of the envelope (REQUEST_SIGNATURE is a request for signature, SIGN_MYSELF is a self-signing flow)
	FlowType *EnvelopeFlowType `json:"flow_type,omitempty"`
	// Legal level of the envelope (SES is Simple Electronic Signature, QES_EIDAS is Qualified Electronic Signature, QES_ZERTES is Qualified Electronic Signature with Zertes)
	LegalityLevel *EnvelopeLegalityLevel `json:"legality_level,omitempty"`
	// Status of the envelope
	Status *EnvelopeStatus `json:"status,omitempty"`
	// Unix timestamp of the creation date
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Unix timestamp of the last modification date
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// Unix timestamp of the expiration date
	ExpiresAt *int64 `json:"expires_at,omitempty"`
	// Number of recipients in the envelope
	NumRecipients *int64 `json:"num_recipients,omitempty"`
	// Whether the envelope can be duplicated
	IsDuplicable *bool                 `json:"is_duplicable,omitempty"`
	SigningSteps []SigningStep         `json:"signing_steps,omitempty"`
	Documents    []Document            `json:"documents,omitempty"`
	Notification *EnvelopeNotification `json:"notification,omitempty"`
}

func (e *Envelope) SetId(id string) {
	e.Id = &id
}

func (e *Envelope) GetId() *string {
	if e == nil {
		return nil
	}
	return e.Id
}

func (e *Envelope) SetName(name string) {
	e.Name = &name
}

func (e *Envelope) GetName() *string {
	if e == nil {
		return nil
	}
	return e.Name
}

func (e *Envelope) SetComment(comment string) {
	e.Comment = &comment
}

func (e *Envelope) GetComment() *string {
	if e == nil {
		return nil
	}
	return e.Comment
}

func (e *Envelope) SetPages(pages int64) {
	e.Pages = &pages
}

func (e *Envelope) GetPages() *int64 {
	if e == nil {
		return nil
	}
	return e.Pages
}

func (e *Envelope) SetFlowType(flowType EnvelopeFlowType) {
	e.FlowType = &flowType
}

func (e *Envelope) GetFlowType() *EnvelopeFlowType {
	if e == nil {
		return nil
	}
	return e.FlowType
}

func (e *Envelope) SetLegalityLevel(legalityLevel EnvelopeLegalityLevel) {
	e.LegalityLevel = &legalityLevel
}

func (e *Envelope) GetLegalityLevel() *EnvelopeLegalityLevel {
	if e == nil {
		return nil
	}
	return e.LegalityLevel
}

func (e *Envelope) SetStatus(status EnvelopeStatus) {
	e.Status = &status
}

func (e *Envelope) GetStatus() *EnvelopeStatus {
	if e == nil {
		return nil
	}
	return e.Status
}

func (e *Envelope) SetCreatedAt(createdAt int64) {
	e.CreatedAt = &createdAt
}

func (e *Envelope) GetCreatedAt() *int64 {
	if e == nil {
		return nil
	}
	return e.CreatedAt
}

func (e *Envelope) SetUpdatedAt(updatedAt int64) {
	e.UpdatedAt = &updatedAt
}

func (e *Envelope) GetUpdatedAt() *int64 {
	if e == nil {
		return nil
	}
	return e.UpdatedAt
}

func (e *Envelope) SetExpiresAt(expiresAt int64) {
	e.ExpiresAt = &expiresAt
}

func (e *Envelope) GetExpiresAt() *int64 {
	if e == nil {
		return nil
	}
	return e.ExpiresAt
}

func (e *Envelope) SetNumRecipients(numRecipients int64) {
	e.NumRecipients = &numRecipients
}

func (e *Envelope) GetNumRecipients() *int64 {
	if e == nil {
		return nil
	}
	return e.NumRecipients
}

func (e *Envelope) SetIsDuplicable(isDuplicable bool) {
	e.IsDuplicable = &isDuplicable
}

func (e *Envelope) GetIsDuplicable() *bool {
	if e == nil {
		return nil
	}
	return e.IsDuplicable
}

func (e *Envelope) SetSigningSteps(signingSteps []SigningStep) {
	e.SigningSteps = signingSteps
}

func (e *Envelope) GetSigningSteps() []SigningStep {
	if e == nil {
		return nil
	}
	return e.SigningSteps
}

func (e *Envelope) SetDocuments(documents []Document) {
	e.Documents = documents
}

func (e *Envelope) GetDocuments() []Document {
	if e == nil {
		return nil
	}
	return e.Documents
}

func (e *Envelope) SetNotification(notification EnvelopeNotification) {
	e.Notification = &notification
}

func (e *Envelope) GetNotification() *EnvelopeNotification {
	if e == nil {
		return nil
	}
	return e.Notification
}

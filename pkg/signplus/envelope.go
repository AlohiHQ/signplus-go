package signplus

import (
	"encoding/json"
)

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
	touched      map[string]bool
}

func (e *Envelope) GetId() *string {
	if e == nil {
		return nil
	}
	return e.Id
}

func (e *Envelope) SetId(id string) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Id"] = true
	e.Id = &id
}

func (e *Envelope) SetIdNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Id"] = true
	e.Id = nil
}

func (e *Envelope) GetName() *string {
	if e == nil {
		return nil
	}
	return e.Name
}

func (e *Envelope) SetName(name string) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Name"] = true
	e.Name = &name
}

func (e *Envelope) SetNameNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Name"] = true
	e.Name = nil
}

func (e *Envelope) GetComment() *string {
	if e == nil {
		return nil
	}
	return e.Comment
}

func (e *Envelope) SetComment(comment string) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Comment"] = true
	e.Comment = &comment
}

func (e *Envelope) SetCommentNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Comment"] = true
	e.Comment = nil
}

func (e *Envelope) GetPages() *int64 {
	if e == nil {
		return nil
	}
	return e.Pages
}

func (e *Envelope) SetPages(pages int64) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Pages"] = true
	e.Pages = &pages
}

func (e *Envelope) SetPagesNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Pages"] = true
	e.Pages = nil
}

func (e *Envelope) GetFlowType() *EnvelopeFlowType {
	if e == nil {
		return nil
	}
	return e.FlowType
}

func (e *Envelope) SetFlowType(flowType EnvelopeFlowType) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["FlowType"] = true
	e.FlowType = &flowType
}

func (e *Envelope) SetFlowTypeNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["FlowType"] = true
	e.FlowType = nil
}

func (e *Envelope) GetLegalityLevel() *EnvelopeLegalityLevel {
	if e == nil {
		return nil
	}
	return e.LegalityLevel
}

func (e *Envelope) SetLegalityLevel(legalityLevel EnvelopeLegalityLevel) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["LegalityLevel"] = true
	e.LegalityLevel = &legalityLevel
}

func (e *Envelope) SetLegalityLevelNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["LegalityLevel"] = true
	e.LegalityLevel = nil
}

func (e *Envelope) GetStatus() *EnvelopeStatus {
	if e == nil {
		return nil
	}
	return e.Status
}

func (e *Envelope) SetStatus(status EnvelopeStatus) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Status"] = true
	e.Status = &status
}

func (e *Envelope) SetStatusNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Status"] = true
	e.Status = nil
}

func (e *Envelope) GetCreatedAt() *int64 {
	if e == nil {
		return nil
	}
	return e.CreatedAt
}

func (e *Envelope) SetCreatedAt(createdAt int64) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["CreatedAt"] = true
	e.CreatedAt = &createdAt
}

func (e *Envelope) SetCreatedAtNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["CreatedAt"] = true
	e.CreatedAt = nil
}

func (e *Envelope) GetUpdatedAt() *int64 {
	if e == nil {
		return nil
	}
	return e.UpdatedAt
}

func (e *Envelope) SetUpdatedAt(updatedAt int64) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["UpdatedAt"] = true
	e.UpdatedAt = &updatedAt
}

func (e *Envelope) SetUpdatedAtNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["UpdatedAt"] = true
	e.UpdatedAt = nil
}

func (e *Envelope) GetExpiresAt() *int64 {
	if e == nil {
		return nil
	}
	return e.ExpiresAt
}

func (e *Envelope) SetExpiresAt(expiresAt int64) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["ExpiresAt"] = true
	e.ExpiresAt = &expiresAt
}

func (e *Envelope) SetExpiresAtNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["ExpiresAt"] = true
	e.ExpiresAt = nil
}

func (e *Envelope) GetNumRecipients() *int64 {
	if e == nil {
		return nil
	}
	return e.NumRecipients
}

func (e *Envelope) SetNumRecipients(numRecipients int64) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["NumRecipients"] = true
	e.NumRecipients = &numRecipients
}

func (e *Envelope) SetNumRecipientsNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["NumRecipients"] = true
	e.NumRecipients = nil
}

func (e *Envelope) GetIsDuplicable() *bool {
	if e == nil {
		return nil
	}
	return e.IsDuplicable
}

func (e *Envelope) SetIsDuplicable(isDuplicable bool) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["IsDuplicable"] = true
	e.IsDuplicable = &isDuplicable
}

func (e *Envelope) SetIsDuplicableNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["IsDuplicable"] = true
	e.IsDuplicable = nil
}

func (e *Envelope) GetSigningSteps() []SigningStep {
	if e == nil {
		return nil
	}
	return e.SigningSteps
}

func (e *Envelope) SetSigningSteps(signingSteps []SigningStep) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["SigningSteps"] = true
	e.SigningSteps = signingSteps
}

func (e *Envelope) SetSigningStepsNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["SigningSteps"] = true
	e.SigningSteps = nil
}

func (e *Envelope) GetDocuments() []Document {
	if e == nil {
		return nil
	}
	return e.Documents
}

func (e *Envelope) SetDocuments(documents []Document) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Documents"] = true
	e.Documents = documents
}

func (e *Envelope) SetDocumentsNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Documents"] = true
	e.Documents = nil
}

func (e *Envelope) GetNotification() *EnvelopeNotification {
	if e == nil {
		return nil
	}
	return e.Notification
}

func (e *Envelope) SetNotification(notification EnvelopeNotification) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Notification"] = true
	e.Notification = &notification
}

func (e *Envelope) SetNotificationNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Notification"] = true
	e.Notification = nil
}
func (e Envelope) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if e.touched["Id"] && e.Id == nil {
		data["id"] = nil
	} else if e.Id != nil {
		data["id"] = e.Id
	}

	if e.touched["Name"] && e.Name == nil {
		data["name"] = nil
	} else if e.Name != nil {
		data["name"] = e.Name
	}

	if e.touched["Comment"] && e.Comment == nil {
		data["comment"] = nil
	} else if e.Comment != nil {
		data["comment"] = e.Comment
	}

	if e.touched["Pages"] && e.Pages == nil {
		data["pages"] = nil
	} else if e.Pages != nil {
		data["pages"] = e.Pages
	}

	if e.touched["FlowType"] && e.FlowType == nil {
		data["flow_type"] = nil
	} else if e.FlowType != nil {
		data["flow_type"] = e.FlowType
	}

	if e.touched["LegalityLevel"] && e.LegalityLevel == nil {
		data["legality_level"] = nil
	} else if e.LegalityLevel != nil {
		data["legality_level"] = e.LegalityLevel
	}

	if e.touched["Status"] && e.Status == nil {
		data["status"] = nil
	} else if e.Status != nil {
		data["status"] = e.Status
	}

	if e.touched["CreatedAt"] && e.CreatedAt == nil {
		data["created_at"] = nil
	} else if e.CreatedAt != nil {
		data["created_at"] = e.CreatedAt
	}

	if e.touched["UpdatedAt"] && e.UpdatedAt == nil {
		data["updated_at"] = nil
	} else if e.UpdatedAt != nil {
		data["updated_at"] = e.UpdatedAt
	}

	if e.touched["ExpiresAt"] && e.ExpiresAt == nil {
		data["expires_at"] = nil
	} else if e.ExpiresAt != nil {
		data["expires_at"] = e.ExpiresAt
	}

	if e.touched["NumRecipients"] && e.NumRecipients == nil {
		data["num_recipients"] = nil
	} else if e.NumRecipients != nil {
		data["num_recipients"] = e.NumRecipients
	}

	if e.touched["IsDuplicable"] && e.IsDuplicable == nil {
		data["is_duplicable"] = nil
	} else if e.IsDuplicable != nil {
		data["is_duplicable"] = e.IsDuplicable
	}

	if e.touched["SigningSteps"] && e.SigningSteps == nil {
		data["signing_steps"] = nil
	} else if e.SigningSteps != nil {
		data["signing_steps"] = e.SigningSteps
	}

	if e.touched["Documents"] && e.Documents == nil {
		data["documents"] = nil
	} else if e.Documents != nil {
		data["documents"] = e.Documents
	}

	if e.touched["Notification"] && e.Notification == nil {
		data["notification"] = nil
	} else if e.Notification != nil {
		data["notification"] = e.Notification
	}

	return json.Marshal(data)
}

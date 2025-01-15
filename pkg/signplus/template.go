package signplus

import (
	"encoding/json"
)

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
	touched       map[string]bool
}

func (t *Template) GetId() *string {
	if t == nil {
		return nil
	}
	return t.Id
}

func (t *Template) SetId(id string) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Id"] = true
	t.Id = &id
}

func (t *Template) SetIdNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Id"] = true
	t.Id = nil
}

func (t *Template) GetName() *string {
	if t == nil {
		return nil
	}
	return t.Name
}

func (t *Template) SetName(name string) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Name"] = true
	t.Name = &name
}

func (t *Template) SetNameNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Name"] = true
	t.Name = nil
}

func (t *Template) GetComment() *string {
	if t == nil {
		return nil
	}
	return t.Comment
}

func (t *Template) SetComment(comment string) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Comment"] = true
	t.Comment = &comment
}

func (t *Template) SetCommentNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Comment"] = true
	t.Comment = nil
}

func (t *Template) GetPages() *int64 {
	if t == nil {
		return nil
	}
	return t.Pages
}

func (t *Template) SetPages(pages int64) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Pages"] = true
	t.Pages = &pages
}

func (t *Template) SetPagesNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Pages"] = true
	t.Pages = nil
}

func (t *Template) GetLegalityLevel() *EnvelopeLegalityLevel {
	if t == nil {
		return nil
	}
	return t.LegalityLevel
}

func (t *Template) SetLegalityLevel(legalityLevel EnvelopeLegalityLevel) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["LegalityLevel"] = true
	t.LegalityLevel = &legalityLevel
}

func (t *Template) SetLegalityLevelNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["LegalityLevel"] = true
	t.LegalityLevel = nil
}

func (t *Template) GetCreatedAt() *int64 {
	if t == nil {
		return nil
	}
	return t.CreatedAt
}

func (t *Template) SetCreatedAt(createdAt int64) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["CreatedAt"] = true
	t.CreatedAt = &createdAt
}

func (t *Template) SetCreatedAtNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["CreatedAt"] = true
	t.CreatedAt = nil
}

func (t *Template) GetUpdatedAt() *int64 {
	if t == nil {
		return nil
	}
	return t.UpdatedAt
}

func (t *Template) SetUpdatedAt(updatedAt int64) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["UpdatedAt"] = true
	t.UpdatedAt = &updatedAt
}

func (t *Template) SetUpdatedAtNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["UpdatedAt"] = true
	t.UpdatedAt = nil
}

func (t *Template) GetExpirationDelay() *int64 {
	if t == nil {
		return nil
	}
	return t.ExpirationDelay
}

func (t *Template) SetExpirationDelay(expirationDelay int64) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["ExpirationDelay"] = true
	t.ExpirationDelay = &expirationDelay
}

func (t *Template) SetExpirationDelayNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["ExpirationDelay"] = true
	t.ExpirationDelay = nil
}

func (t *Template) GetNumRecipients() *int64 {
	if t == nil {
		return nil
	}
	return t.NumRecipients
}

func (t *Template) SetNumRecipients(numRecipients int64) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["NumRecipients"] = true
	t.NumRecipients = &numRecipients
}

func (t *Template) SetNumRecipientsNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["NumRecipients"] = true
	t.NumRecipients = nil
}

func (t *Template) GetSigningSteps() []TemplateSigningStep {
	if t == nil {
		return nil
	}
	return t.SigningSteps
}

func (t *Template) SetSigningSteps(signingSteps []TemplateSigningStep) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["SigningSteps"] = true
	t.SigningSteps = signingSteps
}

func (t *Template) SetSigningStepsNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["SigningSteps"] = true
	t.SigningSteps = nil
}

func (t *Template) GetDocuments() []Document {
	if t == nil {
		return nil
	}
	return t.Documents
}

func (t *Template) SetDocuments(documents []Document) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Documents"] = true
	t.Documents = documents
}

func (t *Template) SetDocumentsNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Documents"] = true
	t.Documents = nil
}

func (t *Template) GetNotification() *EnvelopeNotification {
	if t == nil {
		return nil
	}
	return t.Notification
}

func (t *Template) SetNotification(notification EnvelopeNotification) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Notification"] = true
	t.Notification = &notification
}

func (t *Template) SetNotificationNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Notification"] = true
	t.Notification = nil
}

func (t *Template) GetDynamicFields() []string {
	if t == nil {
		return nil
	}
	return t.DynamicFields
}

func (t *Template) SetDynamicFields(dynamicFields []string) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["DynamicFields"] = true
	t.DynamicFields = dynamicFields
}

func (t *Template) SetDynamicFieldsNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["DynamicFields"] = true
	t.DynamicFields = nil
}
func (t Template) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if t.touched["Id"] && t.Id == nil {
		data["id"] = nil
	} else if t.Id != nil {
		data["id"] = t.Id
	}

	if t.touched["Name"] && t.Name == nil {
		data["name"] = nil
	} else if t.Name != nil {
		data["name"] = t.Name
	}

	if t.touched["Comment"] && t.Comment == nil {
		data["comment"] = nil
	} else if t.Comment != nil {
		data["comment"] = t.Comment
	}

	if t.touched["Pages"] && t.Pages == nil {
		data["pages"] = nil
	} else if t.Pages != nil {
		data["pages"] = t.Pages
	}

	if t.touched["LegalityLevel"] && t.LegalityLevel == nil {
		data["legality_level"] = nil
	} else if t.LegalityLevel != nil {
		data["legality_level"] = t.LegalityLevel
	}

	if t.touched["CreatedAt"] && t.CreatedAt == nil {
		data["created_at"] = nil
	} else if t.CreatedAt != nil {
		data["created_at"] = t.CreatedAt
	}

	if t.touched["UpdatedAt"] && t.UpdatedAt == nil {
		data["updated_at"] = nil
	} else if t.UpdatedAt != nil {
		data["updated_at"] = t.UpdatedAt
	}

	if t.touched["ExpirationDelay"] && t.ExpirationDelay == nil {
		data["expiration_delay"] = nil
	} else if t.ExpirationDelay != nil {
		data["expiration_delay"] = t.ExpirationDelay
	}

	if t.touched["NumRecipients"] && t.NumRecipients == nil {
		data["num_recipients"] = nil
	} else if t.NumRecipients != nil {
		data["num_recipients"] = t.NumRecipients
	}

	if t.touched["SigningSteps"] && t.SigningSteps == nil {
		data["signing_steps"] = nil
	} else if t.SigningSteps != nil {
		data["signing_steps"] = t.SigningSteps
	}

	if t.touched["Documents"] && t.Documents == nil {
		data["documents"] = nil
	} else if t.Documents != nil {
		data["documents"] = t.Documents
	}

	if t.touched["Notification"] && t.Notification == nil {
		data["notification"] = nil
	} else if t.Notification != nil {
		data["notification"] = t.Notification
	}

	if t.touched["DynamicFields"] && t.DynamicFields == nil {
		data["dynamic_fields"] = nil
	} else if t.DynamicFields != nil {
		data["dynamic_fields"] = t.DynamicFields
	}

	return json.Marshal(data)
}

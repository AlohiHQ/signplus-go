package signplus1

import "encoding/json"

type Template struct {
	// Unique identifier of the template
	ID *string `json:"id,omitempty" xml:"id,omitempty"`
	// Name of the template
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Comment for the template
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// Total number of pages in the template
	Pages *int64 `json:"pages,omitempty" xml:"pages,omitempty"`
	// Legal level of the envelope (SES is Simple Electronic Signature, QES_EIDAS is Qualified Electronic Signature, QES_ZERTES is Qualified Electronic Signature with Zertes)
	LegalityLevel *EnvelopeLegalityLevel `json:"legality_level,omitempty" xml:"legality_level,omitempty"`
	// Unix timestamp of the creation date
	CreatedAt *int64 `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// Unix timestamp of the last modification date
	UpdatedAt *int64 `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// Expiration delay added to the current time when an envelope is created from this template
	ExpirationDelay *int64 `json:"expiration_delay,omitempty" xml:"expiration_delay,omitempty"`
	// Number of recipients in the envelope
	NumRecipients *int64                `json:"num_recipients,omitempty" xml:"num_recipients,omitempty"`
	SigningSteps  []TemplateSigningStep `json:"signing_steps,omitempty" xml:"signing_steps,omitempty"`
	Documents     []Document            `json:"documents,omitempty" xml:"documents,omitempty"`
	Notification  *EnvelopeNotification `json:"notification,omitempty" xml:"notification,omitempty"`
	// List of dynamic fields
	DynamicFields []string             `json:"dynamic_fields,omitempty" xml:"dynamic_fields,omitempty"`
	Attachments   *EnvelopeAttachments `json:"attachments,omitempty" xml:"attachments,omitempty"`
}

func (t Template) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: Template to string"
	}
	return string(jsonData)
}

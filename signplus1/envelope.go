package signplus1

import "encoding/json"

type Envelope struct {
	// Unique identifier of the envelope
	ID *string `json:"id,omitempty" xml:"id,omitempty"`
	// Name of the envelope
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Comment for the envelope
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// Total number of pages in the envelope
	Pages *int64 `json:"pages,omitempty" xml:"pages,omitempty"`
	// Flow type of the envelope (REQUEST_SIGNATURE is a request for signature, SIGN_MYSELF is a self-signing flow)
	FlowType *EnvelopeFlowType `json:"flow_type,omitempty" xml:"flow_type,omitempty"`
	// Legal level of the envelope (SES is Simple Electronic Signature, QES_EIDAS is Qualified Electronic Signature, QES_ZERTES is Qualified Electronic Signature with Zertes)
	LegalityLevel *EnvelopeLegalityLevel `json:"legality_level,omitempty" xml:"legality_level,omitempty"`
	// Status of the envelope
	Status *EnvelopeStatus `json:"status,omitempty" xml:"status,omitempty"`
	// Unix timestamp of the creation date
	CreatedAt *int64 `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// Unix timestamp of the last modification date
	UpdatedAt *int64 `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// Unix timestamp of the expiration date
	ExpiresAt *int64 `json:"expires_at,omitempty" xml:"expires_at,omitempty"`
	// Number of recipients in the envelope
	NumRecipients *int64 `json:"num_recipients,omitempty" xml:"num_recipients,omitempty"`
	// Whether the envelope can be duplicated
	IsDuplicable *bool                 `json:"is_duplicable,omitempty" xml:"is_duplicable,omitempty"`
	SigningSteps []SigningStep         `json:"signing_steps,omitempty" xml:"signing_steps,omitempty"`
	Documents    []Document            `json:"documents,omitempty" xml:"documents,omitempty"`
	Notification *EnvelopeNotification `json:"notification,omitempty" xml:"notification,omitempty"`
	Attachments  *EnvelopeAttachments  `json:"attachments,omitempty" xml:"attachments,omitempty"`
}

func (e Envelope) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: Envelope to string"
	}
	return string(jsonData)
}

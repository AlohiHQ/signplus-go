package signplus1

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
)

type AttachmentPlaceholderRequest struct {
	// ID of the recipient
	RecipientID string `json:"recipient_id" xml:"recipient_id" required:"true"`
	// ID of the attachment placeholder
	ID   *string `json:"id,omitempty" xml:"id,omitempty"`
	Name string  `json:"name" xml:"name" required:"true"`
	// Hint of the attachment placeholder
	Hint *string `json:"hint,omitempty" xml:"hint,omitempty"`
	// Whether the attachment placeholder is required
	Required bool `json:"required" xml:"required" required:"true"`
	Multiple bool `json:"multiple" xml:"multiple" required:"true"`
}

func (a AttachmentPlaceholderRequest) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AttachmentPlaceholderRequest to string"
	}
	return string(jsonData)
}

func (a *AttachmentPlaceholderRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AttachmentPlaceholderRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AttachmentPlaceholderRequest(tmp)
	return nil
}

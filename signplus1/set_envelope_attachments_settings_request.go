package signplus1

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
)

type SetEnvelopeAttachmentsSettingsRequest struct {
	Settings AttachmentSettings `json:"settings" xml:"settings" required:"true"`
}

func (s SetEnvelopeAttachmentsSettingsRequest) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SetEnvelopeAttachmentsSettingsRequest to string"
	}
	return string(jsonData)
}

func (s *SetEnvelopeAttachmentsSettingsRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SetEnvelopeAttachmentsSettingsRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SetEnvelopeAttachmentsSettingsRequest(tmp)
	return nil
}

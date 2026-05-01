package settings

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type SetEnvelopeAttachmentsSettingsRequest struct {
	Settings *param.Nullable[SetEnvelopeAttachmentsSettingsRequestSettings] `json:"settings,omitempty" xml:"settings,omitempty"`
}

func (s SetEnvelopeAttachmentsSettingsRequest) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SetEnvelopeAttachmentsSettingsRequest to string"
	}
	return string(jsonData)
}

func (s *SetEnvelopeAttachmentsSettingsRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, s)
}

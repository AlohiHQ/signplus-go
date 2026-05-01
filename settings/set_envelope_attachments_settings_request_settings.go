package settings

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type SetEnvelopeAttachmentsSettingsRequestSettings struct {
	VisibleToRecipients *param.Nullable[bool] `json:"visible_to_recipients,omitempty" xml:"visible_to_recipients,omitempty"`
}

func (s SetEnvelopeAttachmentsSettingsRequestSettings) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SetEnvelopeAttachmentsSettingsRequestSettings to string"
	}
	return string(jsonData)
}

func (s *SetEnvelopeAttachmentsSettingsRequestSettings) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, s)
}

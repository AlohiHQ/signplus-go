package templatetemplateidattachmentssettings

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type SetTemplateAttachmentsSettingsRequest struct {
	Settings *param.Nullable[SetTemplateAttachmentsSettingsRequestSettings] `json:"settings,omitempty" xml:"settings,omitempty"`
}

func (s SetTemplateAttachmentsSettingsRequest) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SetTemplateAttachmentsSettingsRequest to string"
	}
	return string(jsonData)
}

func (s *SetTemplateAttachmentsSettingsRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, s)
}

package templatetemplateidattachmentssettings

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type SetTemplateAttachmentsSettingsRequestSettings struct {
	VisibleToRecipients *param.Nullable[string] `json:"visible_to_recipients,omitempty" xml:"visible_to_recipients,omitempty"`
}

func (s SetTemplateAttachmentsSettingsRequestSettings) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SetTemplateAttachmentsSettingsRequestSettings to string"
	}
	return string(jsonData)
}

func (s *SetTemplateAttachmentsSettingsRequestSettings) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, s)
}

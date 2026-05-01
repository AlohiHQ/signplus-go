package templatetemplateidsetnotification

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type SetTemplateNotificationRequest struct {
	Subject          *param.Nullable[string]  `json:"subject,omitempty" xml:"subject,omitempty"`
	Message          *param.Nullable[string]  `json:"message,omitempty" xml:"message,omitempty"`
	ReminderInterval *param.Nullable[float64] `json:"reminder_interval,omitempty" xml:"reminder_interval,omitempty"`
}

func (s SetTemplateNotificationRequest) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SetTemplateNotificationRequest to string"
	}
	return string(jsonData)
}

func (s *SetTemplateNotificationRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, s)
}

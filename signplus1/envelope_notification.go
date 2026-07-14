package signplus1

import "encoding/json"

type EnvelopeNotification struct {
	// Subject of the notification
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
	// Message of the notification
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Interval in days to send reminder
	ReminderInterval *int64 `json:"reminder_interval,omitempty" xml:"reminder_interval,omitempty"`
}

func (e EnvelopeNotification) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EnvelopeNotification to string"
	}
	return string(jsonData)
}

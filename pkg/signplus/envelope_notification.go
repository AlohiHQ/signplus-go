package signplus

import (
	"encoding/json"
)

type EnvelopeNotification struct {
	// Subject of the notification
	Subject *string `json:"subject,omitempty"`
	// Message of the notification
	Message *string `json:"message,omitempty"`
	// Interval in days to send reminder
	ReminderInterval *int64 `json:"reminder_interval,omitempty"`
	touched          map[string]bool
}

func (e *EnvelopeNotification) GetSubject() *string {
	if e == nil {
		return nil
	}
	return e.Subject
}

func (e *EnvelopeNotification) SetSubject(subject string) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Subject"] = true
	e.Subject = &subject
}

func (e *EnvelopeNotification) SetSubjectNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Subject"] = true
	e.Subject = nil
}

func (e *EnvelopeNotification) GetMessage() *string {
	if e == nil {
		return nil
	}
	return e.Message
}

func (e *EnvelopeNotification) SetMessage(message string) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Message"] = true
	e.Message = &message
}

func (e *EnvelopeNotification) SetMessageNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Message"] = true
	e.Message = nil
}

func (e *EnvelopeNotification) GetReminderInterval() *int64 {
	if e == nil {
		return nil
	}
	return e.ReminderInterval
}

func (e *EnvelopeNotification) SetReminderInterval(reminderInterval int64) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["ReminderInterval"] = true
	e.ReminderInterval = &reminderInterval
}

func (e *EnvelopeNotification) SetReminderIntervalNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["ReminderInterval"] = true
	e.ReminderInterval = nil
}
func (e EnvelopeNotification) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if e.touched["Subject"] && e.Subject == nil {
		data["subject"] = nil
	} else if e.Subject != nil {
		data["subject"] = e.Subject
	}

	if e.touched["Message"] && e.Message == nil {
		data["message"] = nil
	} else if e.Message != nil {
		data["message"] = e.Message
	}

	if e.touched["ReminderInterval"] && e.ReminderInterval == nil {
		data["reminder_interval"] = nil
	} else if e.ReminderInterval != nil {
		data["reminder_interval"] = e.ReminderInterval
	}

	return json.Marshal(data)
}

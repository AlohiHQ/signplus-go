package signplus

type EnvelopeNotification struct {
	// Subject of the notification
	Subject *string `json:"subject,omitempty"`
	// Message of the notification
	Message *string `json:"message,omitempty"`
	// Interval in days to send reminder
	ReminderInterval *int64 `json:"reminder_interval,omitempty"`
}

func (e *EnvelopeNotification) SetSubject(subject string) {
	e.Subject = &subject
}

func (e *EnvelopeNotification) GetSubject() *string {
	if e == nil {
		return nil
	}
	return e.Subject
}

func (e *EnvelopeNotification) SetMessage(message string) {
	e.Message = &message
}

func (e *EnvelopeNotification) GetMessage() *string {
	if e == nil {
		return nil
	}
	return e.Message
}

func (e *EnvelopeNotification) SetReminderInterval(reminderInterval int64) {
	e.ReminderInterval = &reminderInterval
}

func (e *EnvelopeNotification) GetReminderInterval() *int64 {
	if e == nil {
		return nil
	}
	return e.ReminderInterval
}

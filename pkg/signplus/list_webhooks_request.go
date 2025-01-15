package signplus

import (
	"encoding/json"
)

type ListWebhooksRequest struct {
	// ID of the webhook
	WebhookId *string `json:"webhook_id,omitempty"`
	// Event of the webhook
	Event   *WebhookEvent `json:"event,omitempty"`
	touched map[string]bool
}

func (l *ListWebhooksRequest) GetWebhookId() *string {
	if l == nil {
		return nil
	}
	return l.WebhookId
}

func (l *ListWebhooksRequest) SetWebhookId(webhookId string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["WebhookId"] = true
	l.WebhookId = &webhookId
}

func (l *ListWebhooksRequest) SetWebhookIdNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["WebhookId"] = true
	l.WebhookId = nil
}

func (l *ListWebhooksRequest) GetEvent() *WebhookEvent {
	if l == nil {
		return nil
	}
	return l.Event
}

func (l *ListWebhooksRequest) SetEvent(event WebhookEvent) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Event"] = true
	l.Event = &event
}

func (l *ListWebhooksRequest) SetEventNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Event"] = true
	l.Event = nil
}
func (l ListWebhooksRequest) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if l.touched["WebhookId"] && l.WebhookId == nil {
		data["webhook_id"] = nil
	} else if l.WebhookId != nil {
		data["webhook_id"] = l.WebhookId
	}

	if l.touched["Event"] && l.Event == nil {
		data["event"] = nil
	} else if l.Event != nil {
		data["event"] = l.Event
	}

	return json.Marshal(data)
}

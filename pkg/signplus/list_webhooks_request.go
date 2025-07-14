package signplus

import "encoding/json"

type ListWebhooksRequest struct {
	// ID of the webhook
	WebhookId *string `json:"webhook_id,omitempty"`
	// Event of the webhook
	Event *WebhookEvent `json:"event,omitempty"`
}

func (l *ListWebhooksRequest) GetWebhookId() *string {
	if l == nil {
		return nil
	}
	return l.WebhookId
}

func (l *ListWebhooksRequest) SetWebhookId(webhookId string) {
	l.WebhookId = &webhookId
}

func (l *ListWebhooksRequest) GetEvent() *WebhookEvent {
	if l == nil {
		return nil
	}
	return l.Event
}

func (l *ListWebhooksRequest) SetEvent(event WebhookEvent) {
	l.Event = &event
}

func (l ListWebhooksRequest) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListWebhooksRequest to string"
	}
	return string(jsonData)
}

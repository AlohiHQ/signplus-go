package signplus

type ListWebhooksRequest struct {
	// ID of the webhook
	WebhookId *string `json:"webhook_id,omitempty"`
	// Event of the webhook
	Event *WebhookEvent `json:"event,omitempty"`
}

func (l *ListWebhooksRequest) SetWebhookId(webhookId string) {
	l.WebhookId = &webhookId
}

func (l *ListWebhooksRequest) GetWebhookId() *string {
	if l == nil {
		return nil
	}
	return l.WebhookId
}

func (l *ListWebhooksRequest) SetEvent(event WebhookEvent) {
	l.Event = &event
}

func (l *ListWebhooksRequest) GetEvent() *WebhookEvent {
	if l == nil {
		return nil
	}
	return l.Event
}

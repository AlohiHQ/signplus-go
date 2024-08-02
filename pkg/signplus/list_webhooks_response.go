package signplus

type ListWebhooksResponse struct {
	Webhooks []Webhook `json:"webhooks,omitempty"`
}

func (l *ListWebhooksResponse) SetWebhooks(webhooks []Webhook) {
	l.Webhooks = webhooks
}

func (l *ListWebhooksResponse) GetWebhooks() []Webhook {
	if l == nil {
		return nil
	}
	return l.Webhooks
}

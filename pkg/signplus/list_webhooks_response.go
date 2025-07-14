package signplus

import "encoding/json"

type ListWebhooksResponse struct {
	Webhooks []Webhook `json:"webhooks,omitempty"`
}

func (l *ListWebhooksResponse) GetWebhooks() []Webhook {
	if l == nil {
		return nil
	}
	return l.Webhooks
}

func (l *ListWebhooksResponse) SetWebhooks(webhooks []Webhook) {
	l.Webhooks = webhooks
}

func (l ListWebhooksResponse) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListWebhooksResponse to string"
	}
	return string(jsonData)
}

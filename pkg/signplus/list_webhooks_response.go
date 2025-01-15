package signplus

import (
	"encoding/json"
)

type ListWebhooksResponse struct {
	Webhooks []Webhook `json:"webhooks,omitempty"`
	touched  map[string]bool
}

func (l *ListWebhooksResponse) GetWebhooks() []Webhook {
	if l == nil {
		return nil
	}
	return l.Webhooks
}

func (l *ListWebhooksResponse) SetWebhooks(webhooks []Webhook) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Webhooks"] = true
	l.Webhooks = webhooks
}

func (l *ListWebhooksResponse) SetWebhooksNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Webhooks"] = true
	l.Webhooks = nil
}
func (l ListWebhooksResponse) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if l.touched["Webhooks"] && l.Webhooks == nil {
		data["webhooks"] = nil
	} else if l.Webhooks != nil {
		data["webhooks"] = l.Webhooks
	}

	return json.Marshal(data)
}

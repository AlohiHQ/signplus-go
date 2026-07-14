package signplus1

import "encoding/json"

type ListWebhooksRequest struct {
	// ID of the webhook
	WebhookID *string `json:"webhook_id,omitempty" xml:"webhook_id,omitempty"`
	// Event of the webhook
	Event *WebhookEvent `json:"event,omitempty" xml:"event,omitempty"`
}

func (l ListWebhooksRequest) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListWebhooksRequest to string"
	}
	return string(jsonData)
}

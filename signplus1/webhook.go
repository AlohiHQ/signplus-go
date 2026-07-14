package signplus1

import "encoding/json"

type Webhook struct {
	// Unique identifier of the webhook
	ID *string `json:"id,omitempty" xml:"id,omitempty"`
	// Event of the webhook
	Event *WebhookEvent `json:"event,omitempty" xml:"event,omitempty"`
	// Target URL of the webhook
	Target *string `json:"target,omitempty" xml:"target,omitempty"`
}

func (w Webhook) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: Webhook to string"
	}
	return string(jsonData)
}

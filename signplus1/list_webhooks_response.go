package signplus1

import "encoding/json"

type ListWebhooksResponse struct {
	Webhooks []Webhook `json:"webhooks,omitempty" xml:"webhooks,omitempty"`
}

func (l ListWebhooksResponse) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListWebhooksResponse to string"
	}
	return string(jsonData)
}

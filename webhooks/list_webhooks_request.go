package webhooks

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type ListWebhooksRequest struct {
	WebhookID *param.Nullable[string] `json:"webhook_id,omitempty" xml:"webhook_id,omitempty"`
	Event     *param.Nullable[string] `json:"event,omitempty" xml:"event,omitempty"`
}

func (l ListWebhooksRequest) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListWebhooksRequest to string"
	}
	return string(jsonData)
}

func (l *ListWebhooksRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, l)
}

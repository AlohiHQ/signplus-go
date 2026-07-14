package signplus1

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
)

type CreateWebhookRequest struct {
	// Event of the webhook
	Event WebhookEvent `json:"event" xml:"event" required:"true"`
	// URL of the webhook target
	Target string `json:"target" xml:"target" required:"true"`
}

func (c CreateWebhookRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateWebhookRequest to string"
	}
	return string(jsonData)
}

func (c *CreateWebhookRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CreateWebhookRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CreateWebhookRequest(tmp)
	return nil
}

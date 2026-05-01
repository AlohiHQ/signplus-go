package webhook

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type CreateWebhookRequest struct {
	Event  *param.Nullable[string] `json:"event,omitempty" xml:"event,omitempty"`
	Target *param.Nullable[string] `json:"target,omitempty" xml:"target,omitempty"`
}

func (c CreateWebhookRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateWebhookRequest to string"
	}
	return string(jsonData)
}

func (c *CreateWebhookRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, c)
}

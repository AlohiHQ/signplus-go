package signplus

import (
	"encoding/json"
)

type CreateWebhookRequest struct {
	// Event of the webhook
	Event *WebhookEvent `json:"event,omitempty" required:"true"`
	// URL of the webhook target
	Target  *string `json:"target,omitempty" required:"true"`
	touched map[string]bool
}

func (c *CreateWebhookRequest) GetEvent() *WebhookEvent {
	if c == nil {
		return nil
	}
	return c.Event
}

func (c *CreateWebhookRequest) SetEvent(event WebhookEvent) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Event"] = true
	c.Event = &event
}

func (c *CreateWebhookRequest) SetEventNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Event"] = true
	c.Event = nil
}

func (c *CreateWebhookRequest) GetTarget() *string {
	if c == nil {
		return nil
	}
	return c.Target
}

func (c *CreateWebhookRequest) SetTarget(target string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Target"] = true
	c.Target = &target
}

func (c *CreateWebhookRequest) SetTargetNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Target"] = true
	c.Target = nil
}
func (c CreateWebhookRequest) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["Event"] && c.Event == nil {
		data["event"] = nil
	} else if c.Event != nil {
		data["event"] = c.Event
	}

	if c.touched["Target"] && c.Target == nil {
		data["target"] = nil
	} else if c.Target != nil {
		data["target"] = c.Target
	}

	return json.Marshal(data)
}

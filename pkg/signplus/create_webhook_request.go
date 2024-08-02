package signplus

type CreateWebhookRequest struct {
	// Event of the webhook
	Event *WebhookEvent `json:"event,omitempty" required:"true"`
	// URL of the webhook target
	Target *string `json:"target,omitempty" required:"true"`
}

func (c *CreateWebhookRequest) SetEvent(event WebhookEvent) {
	c.Event = &event
}

func (c *CreateWebhookRequest) GetEvent() *WebhookEvent {
	if c == nil {
		return nil
	}
	return c.Event
}

func (c *CreateWebhookRequest) SetTarget(target string) {
	c.Target = &target
}

func (c *CreateWebhookRequest) GetTarget() *string {
	if c == nil {
		return nil
	}
	return c.Target
}

package signplus

type Webhook struct {
	// Unique identifier of the webhook
	Id *string `json:"id,omitempty"`
	// Event of the webhook
	Event *WebhookEvent `json:"event,omitempty"`
	// Target URL of the webhook
	Target *string `json:"target,omitempty"`
}

func (w *Webhook) SetId(id string) {
	w.Id = &id
}

func (w *Webhook) GetId() *string {
	if w == nil {
		return nil
	}
	return w.Id
}

func (w *Webhook) SetEvent(event WebhookEvent) {
	w.Event = &event
}

func (w *Webhook) GetEvent() *WebhookEvent {
	if w == nil {
		return nil
	}
	return w.Event
}

func (w *Webhook) SetTarget(target string) {
	w.Target = &target
}

func (w *Webhook) GetTarget() *string {
	if w == nil {
		return nil
	}
	return w.Target
}

package signplus

import (
	"encoding/json"
)

type Webhook struct {
	// Unique identifier of the webhook
	Id *string `json:"id,omitempty"`
	// Event of the webhook
	Event *WebhookEvent `json:"event,omitempty"`
	// Target URL of the webhook
	Target  *string `json:"target,omitempty"`
	touched map[string]bool
}

func (w *Webhook) GetId() *string {
	if w == nil {
		return nil
	}
	return w.Id
}

func (w *Webhook) SetId(id string) {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["Id"] = true
	w.Id = &id
}

func (w *Webhook) SetIdNil() {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["Id"] = true
	w.Id = nil
}

func (w *Webhook) GetEvent() *WebhookEvent {
	if w == nil {
		return nil
	}
	return w.Event
}

func (w *Webhook) SetEvent(event WebhookEvent) {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["Event"] = true
	w.Event = &event
}

func (w *Webhook) SetEventNil() {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["Event"] = true
	w.Event = nil
}

func (w *Webhook) GetTarget() *string {
	if w == nil {
		return nil
	}
	return w.Target
}

func (w *Webhook) SetTarget(target string) {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["Target"] = true
	w.Target = &target
}

func (w *Webhook) SetTargetNil() {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["Target"] = true
	w.Target = nil
}

func (w Webhook) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if w.touched["Id"] && w.Id == nil {
		data["id"] = nil
	} else if w.Id != nil {
		data["id"] = w.Id
	}

	if w.touched["Event"] && w.Event == nil {
		data["event"] = nil
	} else if w.Event != nil {
		data["event"] = w.Event
	}

	if w.touched["Target"] && w.Target == nil {
		data["target"] = nil
	} else if w.Target != nil {
		data["target"] = w.Target
	}

	return json.Marshal(data)
}

func (w Webhook) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: Webhook to string"
	}
	return string(jsonData)
}

package signplus

import (
	"encoding/json"
)

type RenameEnvelopeRequest struct {
	// Name of the envelope
	Name    *string `json:"name,omitempty"`
	touched map[string]bool
}

func (r *RenameEnvelopeRequest) GetName() *string {
	if r == nil {
		return nil
	}
	return r.Name
}

func (r *RenameEnvelopeRequest) SetName(name string) {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Name"] = true
	r.Name = &name
}

func (r *RenameEnvelopeRequest) SetNameNil() {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Name"] = true
	r.Name = nil
}
func (r RenameEnvelopeRequest) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if r.touched["Name"] && r.Name == nil {
		data["name"] = nil
	} else if r.Name != nil {
		data["name"] = r.Name
	}

	return json.Marshal(data)
}

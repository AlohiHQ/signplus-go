package signplus

import "encoding/json"

type RenameEnvelopeRequest struct {
	// Name of the envelope
	Name *string `json:"name,omitempty"`
}

func (r *RenameEnvelopeRequest) GetName() *string {
	if r == nil {
		return nil
	}
	return r.Name
}

func (r *RenameEnvelopeRequest) SetName(name string) {
	r.Name = &name
}

func (r RenameEnvelopeRequest) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RenameEnvelopeRequest to string"
	}
	return string(jsonData)
}

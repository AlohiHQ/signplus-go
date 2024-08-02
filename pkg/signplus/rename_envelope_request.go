package signplus

type RenameEnvelopeRequest struct {
	// Name of the envelope
	Name *string `json:"name,omitempty"`
}

func (r *RenameEnvelopeRequest) SetName(name string) {
	r.Name = &name
}

func (r *RenameEnvelopeRequest) GetName() *string {
	if r == nil {
		return nil
	}
	return r.Name
}

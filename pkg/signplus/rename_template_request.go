package signplus

type RenameTemplateRequest struct {
	// Name of the template
	Name *string `json:"name,omitempty" required:"true"`
}

func (r *RenameTemplateRequest) SetName(name string) {
	r.Name = &name
}

func (r *RenameTemplateRequest) GetName() *string {
	if r == nil {
		return nil
	}
	return r.Name
}

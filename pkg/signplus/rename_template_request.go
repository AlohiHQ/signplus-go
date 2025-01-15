package signplus

import (
	"encoding/json"
)

type RenameTemplateRequest struct {
	// Name of the template
	Name    *string `json:"name,omitempty" required:"true"`
	touched map[string]bool
}

func (r *RenameTemplateRequest) GetName() *string {
	if r == nil {
		return nil
	}
	return r.Name
}

func (r *RenameTemplateRequest) SetName(name string) {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Name"] = true
	r.Name = &name
}

func (r *RenameTemplateRequest) SetNameNil() {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Name"] = true
	r.Name = nil
}
func (r RenameTemplateRequest) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if r.touched["Name"] && r.Name == nil {
		data["name"] = nil
	} else if r.Name != nil {
		data["name"] = r.Name
	}

	return json.Marshal(data)
}

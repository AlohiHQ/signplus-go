package signplus

import "encoding/json"

type RenameTemplateRequest struct {
	// Name of the template
	Name *string `json:"name,omitempty" required:"true"`
}

func (r *RenameTemplateRequest) GetName() *string {
	if r == nil {
		return nil
	}
	return r.Name
}

func (r *RenameTemplateRequest) SetName(name string) {
	r.Name = &name
}

func (r RenameTemplateRequest) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RenameTemplateRequest to string"
	}
	return string(jsonData)
}

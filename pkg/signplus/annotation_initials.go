package signplus

import (
	"encoding/json"
)

// Initials annotation (null if annotation is not initials)
type AnnotationInitials struct {
	// Unique identifier of the annotation initials
	Id      *string `json:"id,omitempty"`
	touched map[string]bool
}

func (a *AnnotationInitials) GetId() *string {
	if a == nil {
		return nil
	}
	return a.Id
}

func (a *AnnotationInitials) SetId(id string) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Id"] = true
	a.Id = &id
}

func (a *AnnotationInitials) SetIdNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Id"] = true
	a.Id = nil
}

func (a AnnotationInitials) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if a.touched["Id"] && a.Id == nil {
		data["id"] = nil
	} else if a.Id != nil {
		data["id"] = a.Id
	}

	return json.Marshal(data)
}

func (a AnnotationInitials) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AnnotationInitials to string"
	}
	return string(jsonData)
}

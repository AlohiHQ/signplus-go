package signplus

import (
	"encoding/json"
)

// Signature annotation (null if annotation is not a signature)
type AnnotationSignature struct {
	// Unique identifier of the annotation signature
	Id      *string `json:"id,omitempty"`
	touched map[string]bool
}

func (a *AnnotationSignature) GetId() *string {
	if a == nil {
		return nil
	}
	return a.Id
}

func (a *AnnotationSignature) SetId(id string) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Id"] = true
	a.Id = &id
}

func (a *AnnotationSignature) SetIdNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Id"] = true
	a.Id = nil
}

func (a AnnotationSignature) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if a.touched["Id"] && a.Id == nil {
		data["id"] = nil
	} else if a.Id != nil {
		data["id"] = a.Id
	}

	return json.Marshal(data)
}

func (a AnnotationSignature) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AnnotationSignature to string"
	}
	return string(jsonData)
}

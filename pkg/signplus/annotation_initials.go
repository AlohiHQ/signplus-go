package signplus

import "encoding/json"

// Initials annotation (null if annotation is not initials)
type AnnotationInitials struct {
	// Unique identifier of the annotation initials
	Id *string `json:"id,omitempty"`
}

func (a *AnnotationInitials) GetId() *string {
	if a == nil {
		return nil
	}
	return a.Id
}

func (a *AnnotationInitials) SetId(id string) {
	a.Id = &id
}

func (a AnnotationInitials) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AnnotationInitials to string"
	}
	return string(jsonData)
}

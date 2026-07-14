package signplus1

import "encoding/json"

// Initials annotation (null if annotation is not initials)
type AnnotationInitials struct {
	// Unique identifier of the annotation initials
	ID *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (a AnnotationInitials) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AnnotationInitials to string"
	}
	return string(jsonData)
}

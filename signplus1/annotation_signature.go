package signplus1

import "encoding/json"

// Signature annotation (null if annotation is not a signature)
type AnnotationSignature struct {
	// Unique identifier of the annotation signature
	ID *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (a AnnotationSignature) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AnnotationSignature to string"
	}
	return string(jsonData)
}

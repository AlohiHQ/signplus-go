package signplus

import "encoding/json"

// Signature annotation (null if annotation is not a signature)
type AnnotationSignature struct {
	// Unique identifier of the annotation signature
	Id *string `json:"id,omitempty"`
}

func (a *AnnotationSignature) GetId() *string {
	if a == nil {
		return nil
	}
	return a.Id
}

func (a *AnnotationSignature) SetId(id string) {
	a.Id = &id
}

func (a AnnotationSignature) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AnnotationSignature to string"
	}
	return string(jsonData)
}

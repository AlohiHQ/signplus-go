package signplus

// Signature annotation (null if annotation is not a signature)
type AnnotationSignature struct {
	// Unique identifier of the annotation signature
	Id *string `json:"id,omitempty"`
}

func (a *AnnotationSignature) SetId(id string) {
	a.Id = &id
}

func (a *AnnotationSignature) GetId() *string {
	if a == nil {
		return nil
	}
	return a.Id
}

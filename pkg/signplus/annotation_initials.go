package signplus

// Initials annotation (null if annotation is not initials)
type AnnotationInitials struct {
	// Unique identifier of the annotation initials
	Id *string `json:"id,omitempty"`
}

func (a *AnnotationInitials) SetId(id string) {
	a.Id = &id
}

func (a *AnnotationInitials) GetId() *string {
	if a == nil {
		return nil
	}
	return a.Id
}

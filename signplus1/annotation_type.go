package signplus1

// Type of the annotation
type AnnotationType string

const (
	AnnotationTypeText      AnnotationType = "TEXT"
	AnnotationTypeSignature AnnotationType = "SIGNATURE"
	AnnotationTypeInitials  AnnotationType = "INITIALS"
	AnnotationTypeCheckbox  AnnotationType = "CHECKBOX"
	AnnotationTypeDate      AnnotationType = "DATE"
)

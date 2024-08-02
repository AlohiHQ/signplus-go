package signplus

// Type of the annotation
type AnnotationType string

const (
	ANNOTATION_TYPE_TEXT      AnnotationType = "TEXT"
	ANNOTATION_TYPE_SIGNATURE AnnotationType = "SIGNATURE"
	ANNOTATION_TYPE_INITIALS  AnnotationType = "INITIALS"
	ANNOTATION_TYPE_CHECKBOX  AnnotationType = "CHECKBOX"
	ANNOTATION_TYPE_DATE      AnnotationType = "DATE"
)

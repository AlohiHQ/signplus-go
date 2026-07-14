package signplus1

import "encoding/json"

// Text annotation (null if annotation is not a text)
type AnnotationText struct {
	// Font size of the text in pt
	Size *float64 `json:"size,omitempty" xml:"size,omitempty"`
	// Text color in 32bit representation
	Color *float64 `json:"color,omitempty" xml:"color,omitempty"`
	// Text content of the annotation
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// Tooltip of the annotation
	Tooltip *string `json:"tooltip,omitempty" xml:"tooltip,omitempty"`
	// Name of the dynamic field
	DynamicFieldName *string         `json:"dynamic_field_name,omitempty" xml:"dynamic_field_name,omitempty"`
	Font             *AnnotationFont `json:"font,omitempty" xml:"font,omitempty"`
}

func (a AnnotationText) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AnnotationText to string"
	}
	return string(jsonData)
}

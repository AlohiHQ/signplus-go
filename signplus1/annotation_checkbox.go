package signplus1

import "encoding/json"

// Checkbox annotation (null if annotation is not a checkbox)
type AnnotationCheckbox struct {
	// Whether the checkbox is checked
	Checked *bool `json:"checked,omitempty" xml:"checked,omitempty"`
	// Style of the checkbox
	Style *AnnotationCheckboxStyle `json:"style,omitempty" xml:"style,omitempty"`
}

func (a AnnotationCheckbox) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AnnotationCheckbox to string"
	}
	return string(jsonData)
}

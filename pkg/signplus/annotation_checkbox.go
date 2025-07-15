package signplus

import "encoding/json"

// Checkbox annotation (null if annotation is not a checkbox)
type AnnotationCheckbox struct {
	// Whether the checkbox is checked
	Checked *bool `json:"checked,omitempty"`
	// Style of the checkbox
	Style *AnnotationCheckboxStyle `json:"style,omitempty"`
}

func (a *AnnotationCheckbox) GetChecked() *bool {
	if a == nil {
		return nil
	}
	return a.Checked
}

func (a *AnnotationCheckbox) SetChecked(checked bool) {
	a.Checked = &checked
}

func (a *AnnotationCheckbox) GetStyle() *AnnotationCheckboxStyle {
	if a == nil {
		return nil
	}
	return a.Style
}

func (a *AnnotationCheckbox) SetStyle(style AnnotationCheckboxStyle) {
	a.Style = &style
}

func (a AnnotationCheckbox) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AnnotationCheckbox to string"
	}
	return string(jsonData)
}

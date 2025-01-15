package signplus

import (
	"encoding/json"
)

// Checkbox annotation (null if annotation is not a checkbox)
type AnnotationCheckbox struct {
	// Whether the checkbox is checked
	Checked *bool `json:"checked,omitempty"`
	// Style of the checkbox
	Style   *AnnotationCheckboxStyle `json:"style,omitempty"`
	touched map[string]bool
}

func (a *AnnotationCheckbox) GetChecked() *bool {
	if a == nil {
		return nil
	}
	return a.Checked
}

func (a *AnnotationCheckbox) SetChecked(checked bool) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Checked"] = true
	a.Checked = &checked
}

func (a *AnnotationCheckbox) SetCheckedNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Checked"] = true
	a.Checked = nil
}

func (a *AnnotationCheckbox) GetStyle() *AnnotationCheckboxStyle {
	if a == nil {
		return nil
	}
	return a.Style
}

func (a *AnnotationCheckbox) SetStyle(style AnnotationCheckboxStyle) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Style"] = true
	a.Style = &style
}

func (a *AnnotationCheckbox) SetStyleNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Style"] = true
	a.Style = nil
}
func (a AnnotationCheckbox) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if a.touched["Checked"] && a.Checked == nil {
		data["checked"] = nil
	} else if a.Checked != nil {
		data["checked"] = a.Checked
	}

	if a.touched["Style"] && a.Style == nil {
		data["style"] = nil
	} else if a.Style != nil {
		data["style"] = a.Style
	}

	return json.Marshal(data)
}

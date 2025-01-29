package signplus

import (
	"encoding/json"
)

// Text annotation (null if annotation is not a text)
type AnnotationText struct {
	// Font size of the text in pt
	Size *float64 `json:"size,omitempty"`
	// Text color in 32bit representation
	Color *float64 `json:"color,omitempty"`
	// Text content of the annotation
	Value *string `json:"value,omitempty"`
	// Tooltip of the annotation
	Tooltip *string `json:"tooltip,omitempty"`
	// Name of the dynamic field
	DynamicFieldName *string         `json:"dynamic_field_name,omitempty"`
	Font             *AnnotationFont `json:"font,omitempty"`
	touched          map[string]bool
}

func (a *AnnotationText) GetSize() *float64 {
	if a == nil {
		return nil
	}
	return a.Size
}

func (a *AnnotationText) SetSize(size float64) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Size"] = true
	a.Size = &size
}

func (a *AnnotationText) SetSizeNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Size"] = true
	a.Size = nil
}

func (a *AnnotationText) GetColor() *float64 {
	if a == nil {
		return nil
	}
	return a.Color
}

func (a *AnnotationText) SetColor(color float64) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Color"] = true
	a.Color = &color
}

func (a *AnnotationText) SetColorNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Color"] = true
	a.Color = nil
}

func (a *AnnotationText) GetValue() *string {
	if a == nil {
		return nil
	}
	return a.Value
}

func (a *AnnotationText) SetValue(value string) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Value"] = true
	a.Value = &value
}

func (a *AnnotationText) SetValueNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Value"] = true
	a.Value = nil
}

func (a *AnnotationText) GetTooltip() *string {
	if a == nil {
		return nil
	}
	return a.Tooltip
}

func (a *AnnotationText) SetTooltip(tooltip string) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Tooltip"] = true
	a.Tooltip = &tooltip
}

func (a *AnnotationText) SetTooltipNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Tooltip"] = true
	a.Tooltip = nil
}

func (a *AnnotationText) GetDynamicFieldName() *string {
	if a == nil {
		return nil
	}
	return a.DynamicFieldName
}

func (a *AnnotationText) SetDynamicFieldName(dynamicFieldName string) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["DynamicFieldName"] = true
	a.DynamicFieldName = &dynamicFieldName
}

func (a *AnnotationText) SetDynamicFieldNameNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["DynamicFieldName"] = true
	a.DynamicFieldName = nil
}

func (a *AnnotationText) GetFont() *AnnotationFont {
	if a == nil {
		return nil
	}
	return a.Font
}

func (a *AnnotationText) SetFont(font AnnotationFont) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Font"] = true
	a.Font = &font
}

func (a *AnnotationText) SetFontNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Font"] = true
	a.Font = nil
}

func (a AnnotationText) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if a.touched["Size"] && a.Size == nil {
		data["size"] = nil
	} else if a.Size != nil {
		data["size"] = a.Size
	}

	if a.touched["Color"] && a.Color == nil {
		data["color"] = nil
	} else if a.Color != nil {
		data["color"] = a.Color
	}

	if a.touched["Value"] && a.Value == nil {
		data["value"] = nil
	} else if a.Value != nil {
		data["value"] = a.Value
	}

	if a.touched["Tooltip"] && a.Tooltip == nil {
		data["tooltip"] = nil
	} else if a.Tooltip != nil {
		data["tooltip"] = a.Tooltip
	}

	if a.touched["DynamicFieldName"] && a.DynamicFieldName == nil {
		data["dynamic_field_name"] = nil
	} else if a.DynamicFieldName != nil {
		data["dynamic_field_name"] = a.DynamicFieldName
	}

	if a.touched["Font"] && a.Font == nil {
		data["font"] = nil
	} else if a.Font != nil {
		data["font"] = a.Font
	}

	return json.Marshal(data)
}

func (a AnnotationText) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AnnotationText to string"
	}
	return string(jsonData)
}

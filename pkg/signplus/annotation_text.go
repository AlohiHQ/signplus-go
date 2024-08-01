package signplus

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
}

func (a *AnnotationText) SetSize(size float64) {
	a.Size = &size
}

func (a *AnnotationText) GetSize() *float64 {
	if a == nil {
		return nil
	}
	return a.Size
}

func (a *AnnotationText) SetColor(color float64) {
	a.Color = &color
}

func (a *AnnotationText) GetColor() *float64 {
	if a == nil {
		return nil
	}
	return a.Color
}

func (a *AnnotationText) SetValue(value string) {
	a.Value = &value
}

func (a *AnnotationText) GetValue() *string {
	if a == nil {
		return nil
	}
	return a.Value
}

func (a *AnnotationText) SetTooltip(tooltip string) {
	a.Tooltip = &tooltip
}

func (a *AnnotationText) GetTooltip() *string {
	if a == nil {
		return nil
	}
	return a.Tooltip
}

func (a *AnnotationText) SetDynamicFieldName(dynamicFieldName string) {
	a.DynamicFieldName = &dynamicFieldName
}

func (a *AnnotationText) GetDynamicFieldName() *string {
	if a == nil {
		return nil
	}
	return a.DynamicFieldName
}

func (a *AnnotationText) SetFont(font AnnotationFont) {
	a.Font = &font
}

func (a *AnnotationText) GetFont() *AnnotationFont {
	if a == nil {
		return nil
	}
	return a.Font
}

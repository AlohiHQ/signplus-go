package signplus

import "encoding/json"

type AnnotationFont struct {
	// Font family of the text
	Family *AnnotationFontFamily `json:"family,omitempty"`
	// Whether the text is italic
	Italic *bool `json:"italic,omitempty"`
	// Whether the text is bold
	Bold *bool `json:"bold,omitempty"`
}

func (a *AnnotationFont) GetFamily() *AnnotationFontFamily {
	if a == nil {
		return nil
	}
	return a.Family
}

func (a *AnnotationFont) SetFamily(family AnnotationFontFamily) {
	a.Family = &family
}

func (a *AnnotationFont) GetItalic() *bool {
	if a == nil {
		return nil
	}
	return a.Italic
}

func (a *AnnotationFont) SetItalic(italic bool) {
	a.Italic = &italic
}

func (a *AnnotationFont) GetBold() *bool {
	if a == nil {
		return nil
	}
	return a.Bold
}

func (a *AnnotationFont) SetBold(bold bool) {
	a.Bold = &bold
}

func (a AnnotationFont) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AnnotationFont to string"
	}
	return string(jsonData)
}

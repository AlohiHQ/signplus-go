package signplus

import (
	"encoding/json"
)

type AnnotationFont struct {
	// Font family of the text
	Family *AnnotationFontFamily `json:"family,omitempty"`
	// Whether the text is italic
	Italic *bool `json:"italic,omitempty"`
	// Whether the text is bold
	Bold    *bool `json:"bold,omitempty"`
	touched map[string]bool
}

func (a *AnnotationFont) GetFamily() *AnnotationFontFamily {
	if a == nil {
		return nil
	}
	return a.Family
}

func (a *AnnotationFont) SetFamily(family AnnotationFontFamily) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Family"] = true
	a.Family = &family
}

func (a *AnnotationFont) SetFamilyNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Family"] = true
	a.Family = nil
}

func (a *AnnotationFont) GetItalic() *bool {
	if a == nil {
		return nil
	}
	return a.Italic
}

func (a *AnnotationFont) SetItalic(italic bool) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Italic"] = true
	a.Italic = &italic
}

func (a *AnnotationFont) SetItalicNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Italic"] = true
	a.Italic = nil
}

func (a *AnnotationFont) GetBold() *bool {
	if a == nil {
		return nil
	}
	return a.Bold
}

func (a *AnnotationFont) SetBold(bold bool) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Bold"] = true
	a.Bold = &bold
}

func (a *AnnotationFont) SetBoldNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Bold"] = true
	a.Bold = nil
}
func (a AnnotationFont) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if a.touched["Family"] && a.Family == nil {
		data["family"] = nil
	} else if a.Family != nil {
		data["family"] = a.Family
	}

	if a.touched["Italic"] && a.Italic == nil {
		data["italic"] = nil
	} else if a.Italic != nil {
		data["italic"] = a.Italic
	}

	if a.touched["Bold"] && a.Bold == nil {
		data["bold"] = nil
	} else if a.Bold != nil {
		data["bold"] = a.Bold
	}

	return json.Marshal(data)
}

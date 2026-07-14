package signplus1

import "encoding/json"

type AnnotationFont struct {
	// Font family of the text
	Family *AnnotationFontFamily `json:"family,omitempty" xml:"family,omitempty"`
	// Whether the text is italic
	Italic *bool `json:"italic,omitempty" xml:"italic,omitempty"`
	// Whether the text is bold
	Bold *bool `json:"bold,omitempty" xml:"bold,omitempty"`
}

func (a AnnotationFont) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AnnotationFont to string"
	}
	return string(jsonData)
}

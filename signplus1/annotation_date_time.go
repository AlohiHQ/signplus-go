package signplus1

import "encoding/json"

// Date annotation (null if annotation is not a date)
type AnnotationDateTime struct {
	// Font size of the text in pt
	Size *float64        `json:"size,omitempty" xml:"size,omitempty"`
	Font *AnnotationFont `json:"font,omitempty" xml:"font,omitempty"`
	// Color of the text in hex format
	Color *string `json:"color,omitempty" xml:"color,omitempty"`
	// Whether the date should be automatically filled
	AutoFill *bool `json:"auto_fill,omitempty" xml:"auto_fill,omitempty"`
	// Timezone of the date
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// Unix timestamp of the date
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// Format of the date time (DMY_NUMERIC_SLASH is day/month/year with slashes, MDY_NUMERIC_SLASH is month/day/year with slashes, YMD_NUMERIC_SLASH is year/month/day with slashes, DMY_NUMERIC_DASH_SHORT is day/month/year with dashes, DMY_NUMERIC_DASH is day/month/year with dashes, YMD_NUMERIC_DASH is year/month/day with dashes, MDY_TEXT_DASH_SHORT is month/day/year with dashes, MDY_TEXT_SPACE_SHORT is month/day/year with spaces, MDY_TEXT_SPACE is month/day/year with spaces)
	Format *AnnotationDateTimeFormat `json:"format,omitempty" xml:"format,omitempty"`
}

func (a AnnotationDateTime) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AnnotationDateTime to string"
	}
	return string(jsonData)
}

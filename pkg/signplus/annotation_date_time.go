package signplus

import (
	"encoding/json"
)

// Date annotation (null if annotation is not a date)
type AnnotationDateTime struct {
	// Font size of the text in pt
	Size *float64        `json:"size,omitempty"`
	Font *AnnotationFont `json:"font,omitempty"`
	// Color of the text in hex format
	Color *string `json:"color,omitempty"`
	// Whether the date should be automatically filled
	AutoFill *bool `json:"auto_fill,omitempty"`
	// Timezone of the date
	Timezone *string `json:"timezone,omitempty"`
	// Unix timestamp of the date
	Timestamp *int64 `json:"timestamp,omitempty"`
	// Format of the date time (DMY_NUMERIC_SLASH is day/month/year with slashes, MDY_NUMERIC_SLASH is month/day/year with slashes, YMD_NUMERIC_SLASH is year/month/day with slashes, DMY_NUMERIC_DASH_SHORT is day/month/year with dashes, DMY_NUMERIC_DASH is day/month/year with dashes, YMD_NUMERIC_DASH is year/month/day with dashes, MDY_TEXT_DASH_SHORT is month/day/year with dashes, MDY_TEXT_SPACE_SHORT is month/day/year with spaces, MDY_TEXT_SPACE is month/day/year with spaces)
	Format  *AnnotationDateTimeFormat `json:"format,omitempty"`
	touched map[string]bool
}

func (a *AnnotationDateTime) GetSize() *float64 {
	if a == nil {
		return nil
	}
	return a.Size
}

func (a *AnnotationDateTime) SetSize(size float64) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Size"] = true
	a.Size = &size
}

func (a *AnnotationDateTime) SetSizeNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Size"] = true
	a.Size = nil
}

func (a *AnnotationDateTime) GetFont() *AnnotationFont {
	if a == nil {
		return nil
	}
	return a.Font
}

func (a *AnnotationDateTime) SetFont(font AnnotationFont) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Font"] = true
	a.Font = &font
}

func (a *AnnotationDateTime) SetFontNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Font"] = true
	a.Font = nil
}

func (a *AnnotationDateTime) GetColor() *string {
	if a == nil {
		return nil
	}
	return a.Color
}

func (a *AnnotationDateTime) SetColor(color string) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Color"] = true
	a.Color = &color
}

func (a *AnnotationDateTime) SetColorNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Color"] = true
	a.Color = nil
}

func (a *AnnotationDateTime) GetAutoFill() *bool {
	if a == nil {
		return nil
	}
	return a.AutoFill
}

func (a *AnnotationDateTime) SetAutoFill(autoFill bool) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["AutoFill"] = true
	a.AutoFill = &autoFill
}

func (a *AnnotationDateTime) SetAutoFillNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["AutoFill"] = true
	a.AutoFill = nil
}

func (a *AnnotationDateTime) GetTimezone() *string {
	if a == nil {
		return nil
	}
	return a.Timezone
}

func (a *AnnotationDateTime) SetTimezone(timezone string) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Timezone"] = true
	a.Timezone = &timezone
}

func (a *AnnotationDateTime) SetTimezoneNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Timezone"] = true
	a.Timezone = nil
}

func (a *AnnotationDateTime) GetTimestamp() *int64 {
	if a == nil {
		return nil
	}
	return a.Timestamp
}

func (a *AnnotationDateTime) SetTimestamp(timestamp int64) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Timestamp"] = true
	a.Timestamp = &timestamp
}

func (a *AnnotationDateTime) SetTimestampNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Timestamp"] = true
	a.Timestamp = nil
}

func (a *AnnotationDateTime) GetFormat() *AnnotationDateTimeFormat {
	if a == nil {
		return nil
	}
	return a.Format
}

func (a *AnnotationDateTime) SetFormat(format AnnotationDateTimeFormat) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Format"] = true
	a.Format = &format
}

func (a *AnnotationDateTime) SetFormatNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Format"] = true
	a.Format = nil
}

func (a AnnotationDateTime) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if a.touched["Size"] && a.Size == nil {
		data["size"] = nil
	} else if a.Size != nil {
		data["size"] = a.Size
	}

	if a.touched["Font"] && a.Font == nil {
		data["font"] = nil
	} else if a.Font != nil {
		data["font"] = a.Font
	}

	if a.touched["Color"] && a.Color == nil {
		data["color"] = nil
	} else if a.Color != nil {
		data["color"] = a.Color
	}

	if a.touched["AutoFill"] && a.AutoFill == nil {
		data["auto_fill"] = nil
	} else if a.AutoFill != nil {
		data["auto_fill"] = a.AutoFill
	}

	if a.touched["Timezone"] && a.Timezone == nil {
		data["timezone"] = nil
	} else if a.Timezone != nil {
		data["timezone"] = a.Timezone
	}

	if a.touched["Timestamp"] && a.Timestamp == nil {
		data["timestamp"] = nil
	} else if a.Timestamp != nil {
		data["timestamp"] = a.Timestamp
	}

	if a.touched["Format"] && a.Format == nil {
		data["format"] = nil
	} else if a.Format != nil {
		data["format"] = a.Format
	}

	return json.Marshal(data)
}

func (a AnnotationDateTime) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AnnotationDateTime to string"
	}
	return string(jsonData)
}

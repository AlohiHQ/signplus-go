package signplus

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
	Format *AnnotationDateTimeFormat `json:"format,omitempty"`
}

func (a *AnnotationDateTime) SetSize(size float64) {
	a.Size = &size
}

func (a *AnnotationDateTime) GetSize() *float64 {
	if a == nil {
		return nil
	}
	return a.Size
}

func (a *AnnotationDateTime) SetFont(font AnnotationFont) {
	a.Font = &font
}

func (a *AnnotationDateTime) GetFont() *AnnotationFont {
	if a == nil {
		return nil
	}
	return a.Font
}

func (a *AnnotationDateTime) SetColor(color string) {
	a.Color = &color
}

func (a *AnnotationDateTime) GetColor() *string {
	if a == nil {
		return nil
	}
	return a.Color
}

func (a *AnnotationDateTime) SetAutoFill(autoFill bool) {
	a.AutoFill = &autoFill
}

func (a *AnnotationDateTime) GetAutoFill() *bool {
	if a == nil {
		return nil
	}
	return a.AutoFill
}

func (a *AnnotationDateTime) SetTimezone(timezone string) {
	a.Timezone = &timezone
}

func (a *AnnotationDateTime) GetTimezone() *string {
	if a == nil {
		return nil
	}
	return a.Timezone
}

func (a *AnnotationDateTime) SetTimestamp(timestamp int64) {
	a.Timestamp = &timestamp
}

func (a *AnnotationDateTime) GetTimestamp() *int64 {
	if a == nil {
		return nil
	}
	return a.Timestamp
}

func (a *AnnotationDateTime) SetFormat(format AnnotationDateTimeFormat) {
	a.Format = &format
}

func (a *AnnotationDateTime) GetFormat() *AnnotationDateTimeFormat {
	if a == nil {
		return nil
	}
	return a.Format
}

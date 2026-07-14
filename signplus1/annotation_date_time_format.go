package signplus1

// Format of the date time (DMY_NUMERIC_SLASH is day/month/year with slashes, MDY_NUMERIC_SLASH is month/day/year with slashes, YMD_NUMERIC_SLASH is year/month/day with slashes, DMY_NUMERIC_DASH_SHORT is day/month/year with dashes, DMY_NUMERIC_DASH is day/month/year with dashes, YMD_NUMERIC_DASH is year/month/day with dashes, MDY_TEXT_DASH_SHORT is month/day/year with dashes, MDY_TEXT_SPACE_SHORT is month/day/year with spaces, MDY_TEXT_SPACE is month/day/year with spaces)
type AnnotationDateTimeFormat string

const (
	AnnotationDateTimeFormatDmyNumericSlash     AnnotationDateTimeFormat = "DMY_NUMERIC_SLASH"
	AnnotationDateTimeFormatMdyNumericSlash     AnnotationDateTimeFormat = "MDY_NUMERIC_SLASH"
	AnnotationDateTimeFormatYmdNumericSlash     AnnotationDateTimeFormat = "YMD_NUMERIC_SLASH"
	AnnotationDateTimeFormatDmyNumericDashShort AnnotationDateTimeFormat = "DMY_NUMERIC_DASH_SHORT"
	AnnotationDateTimeFormatDmyNumericDash      AnnotationDateTimeFormat = "DMY_NUMERIC_DASH"
	AnnotationDateTimeFormatYmdNumericDash      AnnotationDateTimeFormat = "YMD_NUMERIC_DASH"
	AnnotationDateTimeFormatMdyTextDashShort    AnnotationDateTimeFormat = "MDY_TEXT_DASH_SHORT"
	AnnotationDateTimeFormatMdyTextSpaceShort   AnnotationDateTimeFormat = "MDY_TEXT_SPACE_SHORT"
	AnnotationDateTimeFormatMdyTextSpace        AnnotationDateTimeFormat = "MDY_TEXT_SPACE"
)

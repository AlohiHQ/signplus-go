package signplus

// Format of the date time (DMY_NUMERIC_SLASH is day/month/year with slashes, MDY_NUMERIC_SLASH is month/day/year with slashes, YMD_NUMERIC_SLASH is year/month/day with slashes, DMY_NUMERIC_DASH_SHORT is day/month/year with dashes, DMY_NUMERIC_DASH is day/month/year with dashes, YMD_NUMERIC_DASH is year/month/day with dashes, MDY_TEXT_DASH_SHORT is month/day/year with dashes, MDY_TEXT_SPACE_SHORT is month/day/year with spaces, MDY_TEXT_SPACE is month/day/year with spaces)
type AnnotationDateTimeFormat string

const (
	ANNOTATION_DATE_TIME_FORMAT_DMY_NUMERIC_SLASH      AnnotationDateTimeFormat = "DMY_NUMERIC_SLASH"
	ANNOTATION_DATE_TIME_FORMAT_MDY_NUMERIC_SLASH      AnnotationDateTimeFormat = "MDY_NUMERIC_SLASH"
	ANNOTATION_DATE_TIME_FORMAT_YMD_NUMERIC_SLASH      AnnotationDateTimeFormat = "YMD_NUMERIC_SLASH"
	ANNOTATION_DATE_TIME_FORMAT_DMY_NUMERIC_DASH_SHORT AnnotationDateTimeFormat = "DMY_NUMERIC_DASH_SHORT"
	ANNOTATION_DATE_TIME_FORMAT_DMY_NUMERIC_DASH       AnnotationDateTimeFormat = "DMY_NUMERIC_DASH"
	ANNOTATION_DATE_TIME_FORMAT_YMD_NUMERIC_DASH       AnnotationDateTimeFormat = "YMD_NUMERIC_DASH"
	ANNOTATION_DATE_TIME_FORMAT_MDY_TEXT_DASH_SHORT    AnnotationDateTimeFormat = "MDY_TEXT_DASH_SHORT"
	ANNOTATION_DATE_TIME_FORMAT_MDY_TEXT_SPACE_SHORT   AnnotationDateTimeFormat = "MDY_TEXT_SPACE_SHORT"
	ANNOTATION_DATE_TIME_FORMAT_MDY_TEXT_SPACE         AnnotationDateTimeFormat = "MDY_TEXT_SPACE"
)

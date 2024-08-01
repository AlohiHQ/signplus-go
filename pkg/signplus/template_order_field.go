package signplus

// Field to order templates by
type TemplateOrderField string

const (
	TEMPLATE_ORDER_FIELD_TEMPLATE_ID                TemplateOrderField = "TEMPLATE_ID"
	TEMPLATE_ORDER_FIELD_TEMPLATE_CREATION_DATE     TemplateOrderField = "TEMPLATE_CREATION_DATE"
	TEMPLATE_ORDER_FIELD_TEMPLATE_MODIFICATION_DATE TemplateOrderField = "TEMPLATE_MODIFICATION_DATE"
	TEMPLATE_ORDER_FIELD_TEMPLATE_NAME              TemplateOrderField = "TEMPLATE_NAME"
)

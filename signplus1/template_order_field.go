package signplus1

// Field to order templates by
type TemplateOrderField string

const (
	TemplateOrderFieldTemplateId               TemplateOrderField = "TEMPLATE_ID"
	TemplateOrderFieldTemplateCreationDate     TemplateOrderField = "TEMPLATE_CREATION_DATE"
	TemplateOrderFieldTemplateModificationDate TemplateOrderField = "TEMPLATE_MODIFICATION_DATE"
	TemplateOrderFieldTemplateName             TemplateOrderField = "TEMPLATE_NAME"
)

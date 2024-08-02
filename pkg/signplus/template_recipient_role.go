package signplus

// Role of the recipient (SIGNER signs the document, RECEIVES_COPY receives a copy of the document, IN_PERSON_SIGNER signs the document in person, SENDER sends the document)
type TemplateRecipientRole string

const (
	TEMPLATE_RECIPIENT_ROLE_SIGNER           TemplateRecipientRole = "SIGNER"
	TEMPLATE_RECIPIENT_ROLE_RECEIVES_COPY    TemplateRecipientRole = "RECEIVES_COPY"
	TEMPLATE_RECIPIENT_ROLE_IN_PERSON_SIGNER TemplateRecipientRole = "IN_PERSON_SIGNER"
)

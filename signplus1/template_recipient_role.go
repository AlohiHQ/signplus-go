package signplus1

// Role of the recipient (SIGNER signs the document, RECEIVES_COPY receives a copy of the document, IN_PERSON_SIGNER signs the document in person, SENDER sends the document)
type TemplateRecipientRole string

const (
	TemplateRecipientRoleSigner         TemplateRecipientRole = "SIGNER"
	TemplateRecipientRoleReceivesCopy   TemplateRecipientRole = "RECEIVES_COPY"
	TemplateRecipientRoleInPersonSigner TemplateRecipientRole = "IN_PERSON_SIGNER"
	TemplateRecipientRoleSender         TemplateRecipientRole = "SENDER"
)

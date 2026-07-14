package signplus1

// Role of the recipient (SIGNER signs the document, RECEIVES_COPY receives a copy of the document, IN_PERSON_SIGNER signs the document in person, SENDER sends the document)
type RecipientRole string

const (
	RecipientRoleSigner         RecipientRole = "SIGNER"
	RecipientRoleReceivesCopy   RecipientRole = "RECEIVES_COPY"
	RecipientRoleInPersonSigner RecipientRole = "IN_PERSON_SIGNER"
)

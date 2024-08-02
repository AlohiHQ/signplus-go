package signplus

// Role of the recipient (SIGNER signs the document, RECEIVES_COPY receives a copy of the document, IN_PERSON_SIGNER signs the document in person, SENDER sends the document)
type RecipientRole string

const (
	RECIPIENT_ROLE_SIGNER           RecipientRole = "SIGNER"
	RECIPIENT_ROLE_RECEIVES_COPY    RecipientRole = "RECEIVES_COPY"
	RECIPIENT_ROLE_IN_PERSON_SIGNER RecipientRole = "IN_PERSON_SIGNER"
)

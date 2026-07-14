# Recipient

**Properties**

| Name         | Type                                                         | Required | Description                                                                                                                                                                |
| :----------- | :----------------------------------------------------------- | :------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Name         | string                                                       | ✅       | Name of the recipient                                                                                                                                                      |
| Email        | string                                                       | ✅       | Email of the recipient                                                                                                                                                     |
| Role         | [signplus1.RecipientRole](recipient_role.md)                 | ✅       | Role of the recipient (SIGNER signs the document, RECEIVES_COPY receives a copy of the document, IN_PERSON_SIGNER signs the document in person, SENDER sends the document) |
| ID           | string                                                       | ❌       | Unique identifier of the recipient                                                                                                                                         |
| UID          | string                                                       | ❌       | Unique identifier of the user associated with the recipient                                                                                                                |
| Verification | [signplus1.RecipientVerification](recipient_verification.md) | ❌       |                                                                                                                                                                            |

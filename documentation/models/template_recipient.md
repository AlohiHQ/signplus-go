# TemplateRecipient

**Properties**

| Name  | Type                                                          | Required | Description                                                                                                                                                                |
| :---- | :------------------------------------------------------------ | :------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ID    | string                                                        | ❌       | Unique identifier of the recipient                                                                                                                                         |
| UID   | string                                                        | ❌       | Unique identifier of the user associated with the recipient                                                                                                                |
| Name  | string                                                        | ❌       | Name of the recipient                                                                                                                                                      |
| Email | string                                                        | ❌       | Email of the recipient                                                                                                                                                     |
| Role  | [signplus1.TemplateRecipientRole](template_recipient_role.md) | ❌       | Role of the recipient (SIGNER signs the document, RECEIVES_COPY receives a copy of the document, IN_PERSON_SIGNER signs the document in person, SENDER sends the document) |

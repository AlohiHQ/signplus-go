# Template

**Properties**

| Name            | Type                           | Required | Description                                                                                                                                                             |
| :-------------- | :----------------------------- | :------- | :---------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Id              | string                         | ❌       | Unique identifier of the template                                                                                                                                       |
| Name            | string                         | ❌       | Name of the template                                                                                                                                                    |
| Comment         | string                         | ❌       | Comment for the template                                                                                                                                                |
| Pages           | int64                          | ❌       | Total number of pages in the template                                                                                                                                   |
| LegalityLevel   | signplus.EnvelopeLegalityLevel | ❌       | Legal level of the envelope (SES is Simple Electronic Signature, QES_EIDAS is Qualified Electronic Signature, QES_ZERTES is Qualified Electronic Signature with Zertes) |
| CreatedAt       | int64                          | ❌       | Unix timestamp of the creation date                                                                                                                                     |
| UpdatedAt       | int64                          | ❌       | Unix timestamp of the last modification date                                                                                                                            |
| ExpirationDelay | int64                          | ❌       | Expiration delay added to the current time when an envelope is created from this template                                                                               |
| NumRecipients   | int64                          | ❌       | Number of recipients in the envelope                                                                                                                                    |
| SigningSteps    | []signplus.TemplateSigningStep | ❌       |                                                                                                                                                                         |
| Documents       | []signplus.Document            | ❌       |                                                                                                                                                                         |
| Notification    | signplus.EnvelopeNotification  | ❌       |                                                                                                                                                                         |
| DynamicFields   | []string                       | ❌       | List of dynamic fields                                                                                                                                                  |

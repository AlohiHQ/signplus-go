# Envelope

**Properties**

| Name          | Type                           | Required | Description                                                                                                                                                             |
| :------------ | :----------------------------- | :------- | :---------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Id            | string                         | ❌       | Unique identifier of the envelope                                                                                                                                       |
| Name          | string                         | ❌       | Name of the envelope                                                                                                                                                    |
| Comment       | string                         | ❌       | Comment for the envelope                                                                                                                                                |
| Pages         | int64                          | ❌       | Total number of pages in the envelope                                                                                                                                   |
| FlowType      | signplus.EnvelopeFlowType      | ❌       | Flow type of the envelope (REQUEST_SIGNATURE is a request for signature)                                                                                                |
| LegalityLevel | signplus.EnvelopeLegalityLevel | ❌       | Legal level of the envelope (SES is Simple Electronic Signature, QES_EIDAS is Qualified Electronic Signature, QES_ZERTES is Qualified Electronic Signature with Zertes) |
| Status        | signplus.EnvelopeStatus        | ❌       | Status of the envelope                                                                                                                                                  |
| CreatedAt     | int64                          | ❌       | Unix timestamp of the creation date                                                                                                                                     |
| UpdatedAt     | int64                          | ❌       | Unix timestamp of the last modification date                                                                                                                            |
| ExpiresAt     | int64                          | ❌       | Unix timestamp of the expiration date                                                                                                                                   |
| NumRecipients | int64                          | ❌       | Number of recipients in the envelope                                                                                                                                    |
| IsDuplicable  | bool                           | ❌       | Whether the envelope can be duplicated                                                                                                                                  |
| SigningSteps  | []signplus.SigningStep         | ❌       |                                                                                                                                                                         |
| Documents     | []signplus.Document            | ❌       |                                                                                                                                                                         |
| Notification  | signplus.EnvelopeNotification  | ❌       |                                                                                                                                                                         |

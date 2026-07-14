# Signplus1

A list of all methods in the `Signplus1` service. Click on the method name to view detailed information about that method.

| Methods                                                                   | Description                                                     |
| :------------------------------------------------------------------------ | :-------------------------------------------------------------- |
| [CreateEnvelope](#createenvelope)                                         | Create new envelope                                             |
| [CreateEnvelopeFromTemplate](#createenvelopefromtemplate)                 | Create new envelope from template                               |
| [ListEnvelopes](#listenvelopes)                                           | List envelopes                                                  |
| [GetEnvelope](#getenvelope)                                               | Get envelope                                                    |
| [DeleteEnvelope](#deleteenvelope)                                         | Delete envelope                                                 |
| [DownloadEnvelopeSignedDocuments](#downloadenvelopesigneddocuments)       | Download signed documents for an envelope                       |
| [DownloadEnvelopeCertificate](#downloadenvelopecertificate)               | Download certificate of completion for an envelope              |
| [GetEnvelopeDocument](#getenvelopedocument)                               | Get envelope document                                           |
| [GetEnvelopeDocuments](#getenvelopedocuments)                             | Get envelope documents                                          |
| [AddEnvelopeDocument](#addenvelopedocument)                               | Add envelope document                                           |
| [SetEnvelopeDynamicFields](#setenvelopedynamicfields)                     | Set envelope dynamic fields                                     |
| [AddEnvelopeSigningSteps](#addenvelopesigningsteps)                       | Add envelope signing steps                                      |
| [SetEnvelopeAttachmentsSettings](#setenvelopeattachmentssettings)         | Set envelope attachment settings                                |
| [SetEnvelopeAttachmentsPlaceholders](#setenvelopeattachmentsplaceholders) | Placeholders to be set, completely replacing the existing ones. |
| [GetAttachmentFile](#getattachmentfile)                                   | Get envelope attachment file                                    |
| [SendEnvelope](#sendenvelope)                                             | Send envelope for signature                                     |
| [DuplicateEnvelope](#duplicateenvelope)                                   | Duplicate envelope                                              |
| [VoidEnvelope](#voidenvelope)                                             | Void envelope                                                   |
| [RenameEnvelope](#renameenvelope)                                         | Rename envelope                                                 |
| [SetEnvelopeComment](#setenvelopecomment)                                 | Set envelope comment                                            |
| [SetEnvelopeNotification](#setenvelopenotification)                       | Set envelope notification                                       |
| [SetEnvelopeExpirationDate](#setenvelopeexpirationdate)                   | Set envelope expiration date                                    |
| [SetEnvelopeLegalityLevel](#setenvelopelegalitylevel)                     | Set envelope legality level                                     |
| [GetEnvelopeAnnotations](#getenvelopeannotations)                         | Get envelope annotations                                        |
| [GetEnvelopeDocumentAnnotations](#getenvelopedocumentannotations)         | Get envelope document annotations                               |
| [AddEnvelopeAnnotation](#addenvelopeannotation)                           | Add envelope annotation                                         |
| [DeleteEnvelopeAnnotation](#deleteenvelopeannotation)                     | Delete envelope annotation                                      |
| [CreateTemplate](#createtemplate)                                         | Create new template                                             |
| [ListTemplates](#listtemplates)                                           | List templates                                                  |
| [GetTemplate](#gettemplate)                                               | Get template                                                    |
| [DeleteTemplate](#deletetemplate)                                         | Delete template                                                 |
| [DuplicateTemplate](#duplicatetemplate)                                   | Duplicate template                                              |
| [AddTemplateDocument](#addtemplatedocument)                               | Add template document                                           |
| [GetTemplateDocument](#gettemplatedocument)                               | Get template document                                           |
| [GetTemplateDocuments](#gettemplatedocuments)                             | Get template documents                                          |
| [AddTemplateSigningSteps](#addtemplatesigningsteps)                       | Add template signing steps                                      |
| [RenameTemplate](#renametemplate)                                         | Rename template                                                 |
| [SetTemplateComment](#settemplatecomment)                                 | Set template comment                                            |
| [SetTemplateNotification](#settemplatenotification)                       | Set template notification                                       |
| [GetTemplateAnnotations](#gettemplateannotations)                         | Get template annotations                                        |
| [GetDocumentTemplateAnnotations](#getdocumenttemplateannotations)         | Get document template annotations                               |
| [AddTemplateAnnotation](#addtemplateannotation)                           | Add template annotation                                         |
| [DeleteTemplateAnnotation](#deletetemplateannotation)                     | Delete template annotation                                      |
| [SetTemplateAttachmentsSettings](#settemplateattachmentssettings)         | Set template attachment settings                                |
| [SetTemplateAttachmentsPlaceholders](#settemplateattachmentsplaceholders) | Placeholders to be set, completely replacing the existing ones. |
| [CreateWebhook](#createwebhook)                                           | Create webhook                                                  |
| [ListWebhooks](#listwebhooks)                                             | List webhooks                                                   |
| [DeleteWebhook](#deletewebhook)                                           | Delete webhook                                                  |

## CreateEnvelope

Create new envelope

- HTTP Method: `POST`
- Endpoint: `/envelope`

**Parameters**

| Name                  | Type                  | Required | Description                 |
| :-------------------- | :-------------------- | :------- | :-------------------------- |
| ctx                   | Context               | ✅       | Default go language context |
| createEnvelopeRequest | CreateEnvelopeRequest | ✅       |                             |

**Return Type**

`Envelope`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/signplus1"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)

envelopeLegalityLevel := signplus1.EnvelopeLegalityLevelSes

request := signplus1.CreateEnvelopeRequest{
  Name: "name",
  LegalityLevel: envelopeLegalityLevel,
  ExpiresAt: signplus.Ptr(int64(6)),
  Comment: signplus.Ptr("comment"),
  Sandbox: signplus.Ptr(true),
}

response, err := client.Signplus1.CreateEnvelope(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## CreateEnvelopeFromTemplate

Create new envelope from template

- HTTP Method: `POST`
- Endpoint: `/envelope/from_template/{template_id}`

**Parameters**

| Name                              | Type                              | Required | Description                 |
| :-------------------------------- | :-------------------------------- | :------- | :-------------------------- |
| ctx                               | Context                           | ✅       | Default go language context |
| templateID                        | string                            | ✅       |                             |
| createEnvelopeFromTemplateRequest | CreateEnvelopeFromTemplateRequest | ✅       |                             |

**Return Type**

`Envelope`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/signplus1"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


request := signplus1.CreateEnvelopeFromTemplateRequest{
  Name: "name",
  Comment: signplus.Ptr("comment"),
  Sandbox: signplus.Ptr(true),
}

response, err := client.Signplus1.CreateEnvelopeFromTemplate(context.Background(), "template_id", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListEnvelopes

List envelopes

- HTTP Method: `POST`
- Endpoint: `/envelopes`

**Parameters**

| Name                 | Type                 | Required | Description                 |
| :------------------- | :------------------- | :------- | :-------------------------- |
| ctx                  | Context              | ✅       | Default go language context |
| listEnvelopesRequest | ListEnvelopesRequest | ✅       |                             |

**Return Type**

`ListEnvelopesResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/signplus1"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)

envelopeStatus := signplus1.EnvelopeStatusDraft

envelopeOrderField := signplus1.EnvelopeOrderFieldCreationDate

request := signplus1.ListEnvelopesRequest{
  Name: signplus.Ptr("name"),
  Tags: []string{},
  Comment: signplus.Ptr("comment"),
  Ids: []string{},
  Statuses: []signplus1.EnvelopeStatus{envelopeStatus},
  FolderIds: []string{},
  OnlyRootFolder: signplus.Ptr(true),
  DateFrom: signplus.Ptr(int64(10)),
  DateTo: signplus.Ptr(int64(9)),
  UID: signplus.Ptr("uid"),
  First: signplus.Ptr(int64(2)),
  Last: signplus.Ptr(int64(6)),
  After: signplus.Ptr("after"),
  Before: signplus.Ptr("before"),
  OrderField: &envelopeOrderField,
  Ascending: signplus.Ptr(true),
  IncludeTrash: signplus.Ptr(true),
}

response, err := client.Signplus1.ListEnvelopes(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetEnvelope

Get envelope

- HTTP Method: `GET`
- Endpoint: `/envelope/{envelope_id}`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| envelopeID | string  | ✅       |                             |

**Return Type**

`Envelope`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)

response, err := client.Signplus1.GetEnvelope(context.Background(), "envelope_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteEnvelope

Delete envelope

- HTTP Method: `DELETE`
- Endpoint: `/envelope/{envelope_id}`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| envelopeID | string  | ✅       |                             |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)

response, err := client.Signplus1.DeleteEnvelope(context.Background(), "envelope_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DownloadEnvelopeSignedDocuments

Download signed documents for an envelope

- HTTP Method: `GET`
- Endpoint: `/envelope/{envelope_id}/signed_documents`

**Parameters**

| Name       | Type                                         | Required | Description                   |
| :--------- | :------------------------------------------- | :------- | :---------------------------- |
| ctx        | Context                                      | ✅       | Default go language context   |
| envelopeID | string                                       | ✅       | ID of the envelope            |
| params     | DownloadEnvelopeSignedDocumentsRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/signplus1"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := signplus1.DownloadEnvelopeSignedDocumentsRequestParams{
  CertificateOfCompletion: signplus.Ptr(true),
}

response, err := client.Signplus1.DownloadEnvelopeSignedDocuments(context.Background(), "envelope_id", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DownloadEnvelopeCertificate

Download certificate of completion for an envelope

- HTTP Method: `GET`
- Endpoint: `/envelope/{envelope_id}/certificate`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| envelopeID | string  | ✅       | ID of the envelope          |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)

response, err := client.Signplus1.DownloadEnvelopeCertificate(context.Background(), "envelope_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetEnvelopeDocument

Get envelope document

- HTTP Method: `GET`
- Endpoint: `/envelope/{envelope_id}/document/{document_id}`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| envelopeID | string  | ✅       |                             |
| documentID | string  | ✅       |                             |

**Return Type**

`Document`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)

response, err := client.Signplus1.GetEnvelopeDocument(context.Background(), "envelope_id", "document_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetEnvelopeDocuments

Get envelope documents

- HTTP Method: `GET`
- Endpoint: `/envelope/{envelope_id}/documents`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| envelopeID | string  | ✅       |                             |

**Return Type**

`ListEnvelopeDocumentsResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)

response, err := client.Signplus1.GetEnvelopeDocuments(context.Background(), "envelope_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## AddEnvelopeDocument

Add envelope document

- HTTP Method: `POST`
- Endpoint: `/envelope/{envelope_id}/document`

**Parameters**

| Name                       | Type                       | Required | Description                 |
| :------------------------- | :------------------------- | :------- | :-------------------------- |
| ctx                        | Context                    | ✅       | Default go language context |
| envelopeID                 | string                     | ✅       |                             |
| addEnvelopeDocumentRequest | AddEnvelopeDocumentRequest | ✅       |                             |

**Return Type**

`Document`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/signplus1"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


request := signplus1.AddEnvelopeDocumentRequest{
  File: []byte{},
}

response, err := client.Signplus1.AddEnvelopeDocument(context.Background(), "envelope_id", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SetEnvelopeDynamicFields

Set envelope dynamic fields

- HTTP Method: `PUT`
- Endpoint: `/envelope/{envelope_id}/dynamic_fields`

**Parameters**

| Name                            | Type                            | Required | Description                 |
| :------------------------------ | :------------------------------ | :------- | :-------------------------- |
| ctx                             | Context                         | ✅       | Default go language context |
| envelopeID                      | string                          | ✅       |                             |
| setEnvelopeDynamicFieldsRequest | SetEnvelopeDynamicFieldsRequest | ✅       |                             |

**Return Type**

`Envelope`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/signplus1"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


dynamicField := signplus1.DynamicField{
  Name: signplus.Ptr("name"),
  Value: signplus.Ptr("value"),
}

request := signplus1.SetEnvelopeDynamicFieldsRequest{
  DynamicFields: []signplus1.DynamicField{dynamicField},
}

response, err := client.Signplus1.SetEnvelopeDynamicFields(context.Background(), "envelope_id", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## AddEnvelopeSigningSteps

Add envelope signing steps

- HTTP Method: `POST`
- Endpoint: `/envelope/{envelope_id}/signing_steps`

**Parameters**

| Name                           | Type                           | Required | Description                 |
| :----------------------------- | :----------------------------- | :------- | :-------------------------- |
| ctx                            | Context                        | ✅       | Default go language context |
| envelopeID                     | string                         | ✅       |                             |
| addEnvelopeSigningStepsRequest | AddEnvelopeSigningStepsRequest | ✅       |                             |

**Return Type**

`Envelope`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/signplus1"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)

recipientRole := signplus1.RecipientRoleSigner

recipientVerificationType := signplus1.RecipientVerificationTypeSms

recipientVerification := signplus1.RecipientVerification{
  Type: &recipientVerificationType,
  Value: signplus.Ptr("value"),
}

recipient := signplus1.Recipient{
  ID: signplus.Ptr("id"),
  UID: signplus.Ptr("uid"),
  Name: "name",
  Email: "email",
  Role: recipientRole,
  Verification: &recipientVerification,
}

signingStep := signplus1.SigningStep{
  Recipients: []signplus1.Recipient{recipient},
}

request := signplus1.AddEnvelopeSigningStepsRequest{
  SigningSteps: []signplus1.SigningStep{signingStep},
}

response, err := client.Signplus1.AddEnvelopeSigningSteps(context.Background(), "envelope_id", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SetEnvelopeAttachmentsSettings

Set envelope attachment settings

- HTTP Method: `PUT`
- Endpoint: `/envelope/{envelope_id}/attachments/settings`

**Parameters**

| Name                                  | Type                                  | Required | Description                 |
| :------------------------------------ | :------------------------------------ | :------- | :-------------------------- |
| ctx                                   | Context                               | ✅       | Default go language context |
| envelopeID                            | string                                | ✅       |                             |
| setEnvelopeAttachmentsSettingsRequest | SetEnvelopeAttachmentsSettingsRequest | ✅       |                             |

**Return Type**

`EnvelopeAttachments`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/signplus1"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


attachmentSettings := signplus1.AttachmentSettings{
  VisibleToRecipients: signplus.Ptr(true),
}

request := signplus1.SetEnvelopeAttachmentsSettingsRequest{
  Settings: attachmentSettings,
}

response, err := client.Signplus1.SetEnvelopeAttachmentsSettings(context.Background(), "envelope_id", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SetEnvelopeAttachmentsPlaceholders

Placeholders to be set, completely replacing the existing ones.

- HTTP Method: `PUT`
- Endpoint: `/envelope/{envelope_id}/attachments/placeholders`

**Parameters**

| Name                                      | Type                                      | Required | Description                 |
| :---------------------------------------- | :---------------------------------------- | :------- | :-------------------------- |
| ctx                                       | Context                                   | ✅       | Default go language context |
| envelopeID                                | string                                    | ✅       |                             |
| setEnvelopeAttachmentsPlaceholdersRequest | SetEnvelopeAttachmentsPlaceholdersRequest | ✅       |                             |

**Return Type**

`EnvelopeAttachments`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/signplus1"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


attachmentPlaceholderRequest := signplus1.AttachmentPlaceholderRequest{
  RecipientID: "recipient_id",
  ID: signplus.Ptr("id"),
  Name: "name",
  Hint: signplus.Ptr("hint"),
  Required: true,
  Multiple: true,
}

request := signplus1.SetEnvelopeAttachmentsPlaceholdersRequest{
  Placeholders: []signplus1.AttachmentPlaceholderRequest{attachmentPlaceholderRequest},
}

response, err := client.Signplus1.SetEnvelopeAttachmentsPlaceholders(context.Background(), "envelope_id", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetAttachmentFile

Get envelope attachment file

- HTTP Method: `GET`
- Endpoint: `/envelope/{envelope_id}/attachments/{file_id}`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| envelopeID | string  | ✅       |                             |
| fileID     | string  | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)

response, err := client.Signplus1.GetAttachmentFile(context.Background(), "envelope_id", "file_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SendEnvelope

Send envelope for signature

- HTTP Method: `POST`
- Endpoint: `/envelope/{envelope_id}/send`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| envelopeID | string  | ✅       |                             |

**Return Type**

`Envelope`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)

response, err := client.Signplus1.SendEnvelope(context.Background(), "envelope_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DuplicateEnvelope

Duplicate envelope

- HTTP Method: `POST`
- Endpoint: `/envelope/{envelope_id}/duplicate`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| envelopeID | string  | ✅       |                             |

**Return Type**

`Envelope`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)

response, err := client.Signplus1.DuplicateEnvelope(context.Background(), "envelope_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## VoidEnvelope

Void envelope

- HTTP Method: `PUT`
- Endpoint: `/envelope/{envelope_id}/void`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| envelopeID | string  | ✅       |                             |

**Return Type**

`Envelope`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)

response, err := client.Signplus1.VoidEnvelope(context.Background(), "envelope_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## RenameEnvelope

Rename envelope

- HTTP Method: `PUT`
- Endpoint: `/envelope/{envelope_id}/rename`

**Parameters**

| Name                  | Type                  | Required | Description                 |
| :-------------------- | :-------------------- | :------- | :-------------------------- |
| ctx                   | Context               | ✅       | Default go language context |
| envelopeID            | string                | ✅       |                             |
| renameEnvelopeRequest | RenameEnvelopeRequest | ✅       |                             |

**Return Type**

`Envelope`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/signplus1"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


request := signplus1.RenameEnvelopeRequest{
  Name: signplus.Ptr("name"),
}

response, err := client.Signplus1.RenameEnvelope(context.Background(), "envelope_id", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SetEnvelopeComment

Set envelope comment

- HTTP Method: `PUT`
- Endpoint: `/envelope/{envelope_id}/set_comment`

**Parameters**

| Name                      | Type                      | Required | Description                 |
| :------------------------ | :------------------------ | :------- | :-------------------------- |
| ctx                       | Context                   | ✅       | Default go language context |
| envelopeID                | string                    | ✅       |                             |
| setEnvelopeCommentRequest | SetEnvelopeCommentRequest | ✅       |                             |

**Return Type**

`Envelope`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/signplus1"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


request := signplus1.SetEnvelopeCommentRequest{
  Comment: "comment",
}

response, err := client.Signplus1.SetEnvelopeComment(context.Background(), "envelope_id", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SetEnvelopeNotification

Set envelope notification

- HTTP Method: `PUT`
- Endpoint: `/envelope/{envelope_id}/set_notification`

**Parameters**

| Name                 | Type                 | Required | Description                 |
| :------------------- | :------------------- | :------- | :-------------------------- |
| ctx                  | Context              | ✅       | Default go language context |
| envelopeID           | string               | ✅       |                             |
| envelopeNotification | EnvelopeNotification | ✅       |                             |

**Return Type**

`Envelope`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/signplus1"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


request := signplus1.EnvelopeNotification{
  Subject: signplus.Ptr("subject"),
  Message: signplus.Ptr("message"),
  ReminderInterval: signplus.Ptr(int64(2)),
}

response, err := client.Signplus1.SetEnvelopeNotification(context.Background(), "envelope_id", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SetEnvelopeExpirationDate

Set envelope expiration date

- HTTP Method: `PUT`
- Endpoint: `/envelope/{envelope_id}/set_expiration_date`

**Parameters**

| Name                         | Type                         | Required | Description                 |
| :--------------------------- | :--------------------------- | :------- | :-------------------------- |
| ctx                          | Context                      | ✅       | Default go language context |
| envelopeID                   | string                       | ✅       |                             |
| setEnvelopeExpirationRequest | SetEnvelopeExpirationRequest | ✅       |                             |

**Return Type**

`Envelope`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/signplus1"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


request := signplus1.SetEnvelopeExpirationRequest{
  ExpiresAt: int64(8),
}

response, err := client.Signplus1.SetEnvelopeExpirationDate(context.Background(), "envelope_id", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SetEnvelopeLegalityLevel

Set envelope legality level

- HTTP Method: `PUT`
- Endpoint: `/envelope/{envelope_id}/set_legality_level`

**Parameters**

| Name                            | Type                            | Required | Description                 |
| :------------------------------ | :------------------------------ | :------- | :-------------------------- |
| ctx                             | Context                         | ✅       | Default go language context |
| envelopeID                      | string                          | ✅       |                             |
| setEnvelopeLegalityLevelRequest | SetEnvelopeLegalityLevelRequest | ✅       |                             |

**Return Type**

`Envelope`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/signplus1"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)

envelopeLegalityLevel := signplus1.EnvelopeLegalityLevelSes

request := signplus1.SetEnvelopeLegalityLevelRequest{
  LegalityLevel: &envelopeLegalityLevel,
}

response, err := client.Signplus1.SetEnvelopeLegalityLevel(context.Background(), "envelope_id", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetEnvelopeAnnotations

Get envelope annotations

- HTTP Method: `GET`
- Endpoint: `/envelope/{envelope_id}/annotations`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| envelopeID | string  | ✅       | ID of the envelope          |

**Return Type**

`[]Annotation`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)

response, err := client.Signplus1.GetEnvelopeAnnotations(context.Background(), "envelope_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetEnvelopeDocumentAnnotations

Get envelope document annotations

- HTTP Method: `GET`
- Endpoint: `/envelope/{envelope_id}/annotations/{document_id}`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| envelopeID | string  | ✅       | ID of the envelope          |
| documentID | string  | ✅       | ID of document              |

**Return Type**

`ListEnvelopeDocumentAnnotationsResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)

response, err := client.Signplus1.GetEnvelopeDocumentAnnotations(context.Background(), "envelope_id", "document_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## AddEnvelopeAnnotation

Add envelope annotation

- HTTP Method: `POST`
- Endpoint: `/envelope/{envelope_id}/annotation`

**Parameters**

| Name                 | Type                 | Required | Description                 |
| :------------------- | :------------------- | :------- | :-------------------------- |
| ctx                  | Context              | ✅       | Default go language context |
| envelopeID           | string               | ✅       | ID of the envelope          |
| addAnnotationRequest | AddAnnotationRequest | ✅       |                             |

**Return Type**

`Annotation`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/signplus1"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)

annotationType := signplus1.AnnotationTypeText


annotationSignature := signplus1.AnnotationSignature{
  ID: signplus.Ptr("id"),
}


annotationInitials := signplus1.AnnotationInitials{
  ID: signplus.Ptr("id"),
}

annotationFontFamily := signplus1.AnnotationFontFamilyUnknown

annotationFont := signplus1.AnnotationFont{
  Family: &annotationFontFamily,
  Italic: signplus.Ptr(true),
  Bold: signplus.Ptr(true),
}

annotationText := signplus1.AnnotationText{
  Size: signplus.Ptr(float64(4.85)),
  Color: signplus.Ptr(float64(0.58)),
  Value: signplus.Ptr("value"),
  Tooltip: signplus.Ptr("tooltip"),
  DynamicFieldName: signplus.Ptr("dynamic_field_name"),
  Font: &annotationFont,
}

annotationFontFamily := signplus1.AnnotationFontFamilyUnknown

annotationFont := signplus1.AnnotationFont{
  Family: &annotationFontFamily,
  Italic: signplus.Ptr(true),
  Bold: signplus.Ptr(true),
}

annotationDateTimeFormat := signplus1.AnnotationDateTimeFormatDmyNumericSlash

annotationDateTime := signplus1.AnnotationDateTime{
  Size: signplus.Ptr(float64(4.53)),
  Font: &annotationFont,
  Color: signplus.Ptr("color"),
  AutoFill: signplus.Ptr(true),
  Timezone: signplus.Ptr("timezone"),
  Timestamp: signplus.Ptr(int64(3)),
  Format: &annotationDateTimeFormat,
}

annotationCheckboxStyle := signplus1.AnnotationCheckboxStyleCircleCheck

annotationCheckbox := signplus1.AnnotationCheckbox{
  Checked: signplus.Ptr(true),
  Style: &annotationCheckboxStyle,
}

request := signplus1.AddAnnotationRequest{
  RecipientID: signplus.Ptr("recipient_id"),
  DocumentID: "document_id",
  Page: int64(1),
  X: float64(1.09),
  Y: float64(1.5),
  Width: float64(9.87),
  Height: float64(1.48),
  Required: signplus.Ptr(true),
  Type: annotationType,
  Signature: &annotationSignature,
  Initials: &annotationInitials,
  Text: &annotationText,
  Datetime: &annotationDateTime,
  Checkbox: &annotationCheckbox,
}

response, err := client.Signplus1.AddEnvelopeAnnotation(context.Background(), "envelope_id", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteEnvelopeAnnotation

Delete envelope annotation

- HTTP Method: `DELETE`
- Endpoint: `/envelope/{envelope_id}/annotation/{annotation_id}`

**Parameters**

| Name         | Type    | Required | Description                    |
| :----------- | :------ | :------- | :----------------------------- |
| ctx          | Context | ✅       | Default go language context    |
| envelopeID   | string  | ✅       | ID of the envelope             |
| annotationID | string  | ✅       | ID of the annotation to delete |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)

response, err := client.Signplus1.DeleteEnvelopeAnnotation(context.Background(), "envelope_id", "annotation_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## CreateTemplate

Create new template

- HTTP Method: `POST`
- Endpoint: `/template`

**Parameters**

| Name                  | Type                  | Required | Description                 |
| :-------------------- | :-------------------- | :------- | :-------------------------- |
| ctx                   | Context               | ✅       | Default go language context |
| createTemplateRequest | CreateTemplateRequest | ✅       |                             |

**Return Type**

`Template`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/signplus1"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


request := signplus1.CreateTemplateRequest{
  Name: "name",
}

response, err := client.Signplus1.CreateTemplate(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListTemplates

List templates

- HTTP Method: `POST`
- Endpoint: `/templates`

**Parameters**

| Name                 | Type                 | Required | Description                 |
| :------------------- | :------------------- | :------- | :-------------------------- |
| ctx                  | Context              | ✅       | Default go language context |
| listTemplatesRequest | ListTemplatesRequest | ✅       |                             |

**Return Type**

`ListTemplatesResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/signplus1"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)

templateOrderField := signplus1.TemplateOrderFieldTemplateId

request := signplus1.ListTemplatesRequest{
  Name: signplus.Ptr("name"),
  Tags: []string{},
  Ids: []string{},
  First: signplus.Ptr(int64(7)),
  Last: signplus.Ptr(int64(7)),
  After: signplus.Ptr("after"),
  Before: signplus.Ptr("before"),
  OrderField: &templateOrderField,
  Ascending: signplus.Ptr(true),
}

response, err := client.Signplus1.ListTemplates(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetTemplate

Get template

- HTTP Method: `GET`
- Endpoint: `/template/{template_id}`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| templateID | string  | ✅       |                             |

**Return Type**

`Template`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)

response, err := client.Signplus1.GetTemplate(context.Background(), "template_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteTemplate

Delete template

- HTTP Method: `DELETE`
- Endpoint: `/template/{template_id}`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| templateID | string  | ✅       |                             |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)

response, err := client.Signplus1.DeleteTemplate(context.Background(), "template_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DuplicateTemplate

Duplicate template

- HTTP Method: `POST`
- Endpoint: `/template/{template_id}/duplicate`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| templateID | string  | ✅       |                             |

**Return Type**

`Template`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)

response, err := client.Signplus1.DuplicateTemplate(context.Background(), "template_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## AddTemplateDocument

Add template document

- HTTP Method: `POST`
- Endpoint: `/template/{template_id}/document`

**Parameters**

| Name                       | Type                       | Required | Description                 |
| :------------------------- | :------------------------- | :------- | :-------------------------- |
| ctx                        | Context                    | ✅       | Default go language context |
| templateID                 | string                     | ✅       |                             |
| addTemplateDocumentRequest | AddTemplateDocumentRequest | ✅       |                             |

**Return Type**

`Document`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/signplus1"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


request := signplus1.AddTemplateDocumentRequest{
  File: []byte{},
}

response, err := client.Signplus1.AddTemplateDocument(context.Background(), "template_id", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetTemplateDocument

Get template document

- HTTP Method: `GET`
- Endpoint: `/template/{template_id}/document/{document_id}`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| templateID | string  | ✅       |                             |
| documentID | string  | ✅       |                             |

**Return Type**

`Document`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)

response, err := client.Signplus1.GetTemplateDocument(context.Background(), "template_id", "document_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetTemplateDocuments

Get template documents

- HTTP Method: `GET`
- Endpoint: `/template/{template_id}/documents`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| templateID | string  | ✅       |                             |

**Return Type**

`ListTemplateDocumentsResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)

response, err := client.Signplus1.GetTemplateDocuments(context.Background(), "template_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## AddTemplateSigningSteps

Add template signing steps

- HTTP Method: `POST`
- Endpoint: `/template/{template_id}/signing_steps`

**Parameters**

| Name                           | Type                           | Required | Description                 |
| :----------------------------- | :----------------------------- | :------- | :-------------------------- |
| ctx                            | Context                        | ✅       | Default go language context |
| templateID                     | string                         | ✅       |                             |
| addTemplateSigningStepsRequest | AddTemplateSigningStepsRequest | ✅       |                             |

**Return Type**

`Template`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/signplus1"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)

templateRecipientRole := signplus1.TemplateRecipientRoleSigner

templateRecipient := signplus1.TemplateRecipient{
  ID: signplus.Ptr("id"),
  UID: signplus.Ptr("uid"),
  Name: signplus.Ptr("name"),
  Email: signplus.Ptr("email"),
  Role: &templateRecipientRole,
}

templateSigningStep := signplus1.TemplateSigningStep{
  Recipients: []signplus1.TemplateRecipient{templateRecipient},
}

request := signplus1.AddTemplateSigningStepsRequest{
  SigningSteps: []signplus1.TemplateSigningStep{templateSigningStep},
}

response, err := client.Signplus1.AddTemplateSigningSteps(context.Background(), "template_id", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## RenameTemplate

Rename template

- HTTP Method: `PUT`
- Endpoint: `/template/{template_id}/rename`

**Parameters**

| Name                  | Type                  | Required | Description                 |
| :-------------------- | :-------------------- | :------- | :-------------------------- |
| ctx                   | Context               | ✅       | Default go language context |
| templateID            | string                | ✅       |                             |
| renameTemplateRequest | RenameTemplateRequest | ✅       |                             |

**Return Type**

`Template`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/signplus1"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


request := signplus1.RenameTemplateRequest{
  Name: "name",
}

response, err := client.Signplus1.RenameTemplate(context.Background(), "template_id", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SetTemplateComment

Set template comment

- HTTP Method: `PUT`
- Endpoint: `/template/{template_id}/set_comment`

**Parameters**

| Name                      | Type                      | Required | Description                 |
| :------------------------ | :------------------------ | :------- | :-------------------------- |
| ctx                       | Context                   | ✅       | Default go language context |
| templateID                | string                    | ✅       |                             |
| setTemplateCommentRequest | SetTemplateCommentRequest | ✅       |                             |

**Return Type**

`Template`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/signplus1"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


request := signplus1.SetTemplateCommentRequest{
  Comment: "comment",
}

response, err := client.Signplus1.SetTemplateComment(context.Background(), "template_id", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SetTemplateNotification

Set template notification

- HTTP Method: `PUT`
- Endpoint: `/template/{template_id}/set_notification`

**Parameters**

| Name                 | Type                 | Required | Description                 |
| :------------------- | :------------------- | :------- | :-------------------------- |
| ctx                  | Context              | ✅       | Default go language context |
| templateID           | string               | ✅       |                             |
| envelopeNotification | EnvelopeNotification | ✅       |                             |

**Return Type**

`Template`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/signplus1"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


request := signplus1.EnvelopeNotification{
  Subject: signplus.Ptr("subject"),
  Message: signplus.Ptr("message"),
  ReminderInterval: signplus.Ptr(int64(2)),
}

response, err := client.Signplus1.SetTemplateNotification(context.Background(), "template_id", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetTemplateAnnotations

Get template annotations

- HTTP Method: `GET`
- Endpoint: `/template/{template_id}/annotations`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| templateID | string  | ✅       | ID of the template          |

**Return Type**

`ListTemplateAnnotationsResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)

response, err := client.Signplus1.GetTemplateAnnotations(context.Background(), "template_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetDocumentTemplateAnnotations

Get document template annotations

- HTTP Method: `GET`
- Endpoint: `/template/{template_id}/annotations/{document_id}`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| templateID | string  | ✅       | ID of the template          |
| documentID | string  | ✅       | ID of document              |

**Return Type**

`ListTemplateDocumentAnnotationsResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)

response, err := client.Signplus1.GetDocumentTemplateAnnotations(context.Background(), "template_id", "document_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## AddTemplateAnnotation

Add template annotation

- HTTP Method: `POST`
- Endpoint: `/template/{template_id}/annotation`

**Parameters**

| Name                 | Type                 | Required | Description                 |
| :------------------- | :------------------- | :------- | :-------------------------- |
| ctx                  | Context              | ✅       | Default go language context |
| templateID           | string               | ✅       | ID of the template          |
| addAnnotationRequest | AddAnnotationRequest | ✅       |                             |

**Return Type**

`Annotation`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/signplus1"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)

annotationType := signplus1.AnnotationTypeText


annotationSignature := signplus1.AnnotationSignature{
  ID: signplus.Ptr("id"),
}


annotationInitials := signplus1.AnnotationInitials{
  ID: signplus.Ptr("id"),
}

annotationFontFamily := signplus1.AnnotationFontFamilyUnknown

annotationFont := signplus1.AnnotationFont{
  Family: &annotationFontFamily,
  Italic: signplus.Ptr(true),
  Bold: signplus.Ptr(true),
}

annotationText := signplus1.AnnotationText{
  Size: signplus.Ptr(float64(4.85)),
  Color: signplus.Ptr(float64(0.58)),
  Value: signplus.Ptr("value"),
  Tooltip: signplus.Ptr("tooltip"),
  DynamicFieldName: signplus.Ptr("dynamic_field_name"),
  Font: &annotationFont,
}

annotationFontFamily := signplus1.AnnotationFontFamilyUnknown

annotationFont := signplus1.AnnotationFont{
  Family: &annotationFontFamily,
  Italic: signplus.Ptr(true),
  Bold: signplus.Ptr(true),
}

annotationDateTimeFormat := signplus1.AnnotationDateTimeFormatDmyNumericSlash

annotationDateTime := signplus1.AnnotationDateTime{
  Size: signplus.Ptr(float64(4.53)),
  Font: &annotationFont,
  Color: signplus.Ptr("color"),
  AutoFill: signplus.Ptr(true),
  Timezone: signplus.Ptr("timezone"),
  Timestamp: signplus.Ptr(int64(3)),
  Format: &annotationDateTimeFormat,
}

annotationCheckboxStyle := signplus1.AnnotationCheckboxStyleCircleCheck

annotationCheckbox := signplus1.AnnotationCheckbox{
  Checked: signplus.Ptr(true),
  Style: &annotationCheckboxStyle,
}

request := signplus1.AddAnnotationRequest{
  RecipientID: signplus.Ptr("recipient_id"),
  DocumentID: "document_id",
  Page: int64(1),
  X: float64(1.09),
  Y: float64(1.5),
  Width: float64(9.87),
  Height: float64(1.48),
  Required: signplus.Ptr(true),
  Type: annotationType,
  Signature: &annotationSignature,
  Initials: &annotationInitials,
  Text: &annotationText,
  Datetime: &annotationDateTime,
  Checkbox: &annotationCheckbox,
}

response, err := client.Signplus1.AddTemplateAnnotation(context.Background(), "template_id", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteTemplateAnnotation

Delete template annotation

- HTTP Method: `DELETE`
- Endpoint: `/template/{template_id}/annotation/{annotation_id}`

**Parameters**

| Name         | Type    | Required | Description                    |
| :----------- | :------ | :------- | :----------------------------- |
| ctx          | Context | ✅       | Default go language context    |
| templateID   | string  | ✅       | ID of the template             |
| annotationID | string  | ✅       | ID of the annotation to delete |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)

response, err := client.Signplus1.DeleteTemplateAnnotation(context.Background(), "template_id", "annotation_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SetTemplateAttachmentsSettings

Set template attachment settings

- HTTP Method: `PUT`
- Endpoint: `/template/{template_id}/attachments/settings`

**Parameters**

| Name                                  | Type                                  | Required | Description                 |
| :------------------------------------ | :------------------------------------ | :------- | :-------------------------- |
| ctx                                   | Context                               | ✅       | Default go language context |
| templateID                            | string                                | ✅       |                             |
| setEnvelopeAttachmentsSettingsRequest | SetEnvelopeAttachmentsSettingsRequest | ✅       |                             |

**Return Type**

`EnvelopeAttachments`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/signplus1"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


attachmentSettings := signplus1.AttachmentSettings{
  VisibleToRecipients: signplus.Ptr(true),
}

request := signplus1.SetEnvelopeAttachmentsSettingsRequest{
  Settings: attachmentSettings,
}

response, err := client.Signplus1.SetTemplateAttachmentsSettings(context.Background(), "template_id", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SetTemplateAttachmentsPlaceholders

Placeholders to be set, completely replacing the existing ones.

- HTTP Method: `PUT`
- Endpoint: `/template/{template_id}/attachments/placeholders`

**Parameters**

| Name                                      | Type                                      | Required | Description                 |
| :---------------------------------------- | :---------------------------------------- | :------- | :-------------------------- |
| ctx                                       | Context                                   | ✅       | Default go language context |
| templateID                                | string                                    | ✅       |                             |
| setEnvelopeAttachmentsPlaceholdersRequest | SetEnvelopeAttachmentsPlaceholdersRequest | ✅       |                             |

**Return Type**

`EnvelopeAttachments`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/signplus1"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


attachmentPlaceholderRequest := signplus1.AttachmentPlaceholderRequest{
  RecipientID: "recipient_id",
  ID: signplus.Ptr("id"),
  Name: "name",
  Hint: signplus.Ptr("hint"),
  Required: true,
  Multiple: true,
}

request := signplus1.SetEnvelopeAttachmentsPlaceholdersRequest{
  Placeholders: []signplus1.AttachmentPlaceholderRequest{attachmentPlaceholderRequest},
}

response, err := client.Signplus1.SetTemplateAttachmentsPlaceholders(context.Background(), "template_id", request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## CreateWebhook

Create webhook

- HTTP Method: `POST`
- Endpoint: `/webhook`

**Parameters**

| Name                 | Type                 | Required | Description                 |
| :------------------- | :------------------- | :------- | :-------------------------- |
| ctx                  | Context              | ✅       | Default go language context |
| createWebhookRequest | CreateWebhookRequest | ✅       |                             |

**Return Type**

`Webhook`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/signplus1"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)

webhookEvent := signplus1.WebhookEventEnvelopeExpired

request := signplus1.CreateWebhookRequest{
  Event: webhookEvent,
  Target: "target",
}

response, err := client.Signplus1.CreateWebhook(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListWebhooks

List webhooks

- HTTP Method: `POST`
- Endpoint: `/webhooks`

**Parameters**

| Name                | Type                | Required | Description                 |
| :------------------ | :------------------ | :------- | :-------------------------- |
| ctx                 | Context             | ✅       | Default go language context |
| listWebhooksRequest | ListWebhooksRequest | ✅       |                             |

**Return Type**

`ListWebhooksResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/signplus1"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)

webhookEvent := signplus1.WebhookEventEnvelopeExpired

request := signplus1.ListWebhooksRequest{
  WebhookID: signplus.Ptr("webhook_id"),
  Event: &webhookEvent,
}

response, err := client.Signplus1.ListWebhooks(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteWebhook

Delete webhook

- HTTP Method: `DELETE`
- Endpoint: `/webhook/{webhook_id}`

**Parameters**

| Name      | Type    | Required | Description                 |
| :-------- | :------ | :------- | :-------------------------- |
| ctx       | Context | ✅       | Default go language context |
| webhookID | string  | ✅       |                             |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)

response, err := client.Signplus1.DeleteWebhook(context.Background(), "webhook_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

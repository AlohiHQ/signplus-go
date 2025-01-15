# SignplusService

A list of all methods in the `SignplusService` service. Click on the method name to view detailed information about that method.

| Methods                                                           | Description                       |
| :---------------------------------------------------------------- | :-------------------------------- |
| [CreateEnvelope](#createenvelope)                                 | Create new envelope               |
| [CreateEnvelopeFromTemplate](#createenvelopefromtemplate)         | Create new envelope from template |
| [ListEnvelopes](#listenvelopes)                                   | List envelopes                    |
| [GetEnvelope](#getenvelope)                                       | Get envelope                      |
| [DeleteEnvelope](#deleteenvelope)                                 | Delete envelope                   |
| [GetEnvelopeDocument](#getenvelopedocument)                       | Get envelope document             |
| [GetEnvelopeDocuments](#getenvelopedocuments)                     | Get envelope documents            |
| [AddEnvelopeDocument](#addenvelopedocument)                       | Add envelope document             |
| [SetEnvelopeDynamicFields](#setenvelopedynamicfields)             | Set envelope dynamic fields       |
| [AddEnvelopeSigningSteps](#addenvelopesigningsteps)               | Add envelope signing steps        |
| [SendEnvelope](#sendenvelope)                                     | Send envelope for signature       |
| [DuplicateEnvelope](#duplicateenvelope)                           | Duplicate envelope                |
| [VoidEnvelope](#voidenvelope)                                     | Void envelope                     |
| [RenameEnvelope](#renameenvelope)                                 | Rename envelope                   |
| [SetEnvelopeComment](#setenvelopecomment)                         | Set envelope comment              |
| [SetEnvelopeNotification](#setenvelopenotification)               | Set envelope notification         |
| [SetEnvelopeExpirationDate](#setenvelopeexpirationdate)           | Set envelope expiration date      |
| [SetEnvelopeLegalityLevel](#setenvelopelegalitylevel)             | Set envelope legality level       |
| [GetEnvelopeAnnotations](#getenvelopeannotations)                 | Get envelope annotations          |
| [GetEnvelopeDocumentAnnotations](#getenvelopedocumentannotations) | Get envelope document annotations |
| [AddEnvelopeAnnotation](#addenvelopeannotation)                   | Add envelope annotation           |
| [DeleteEnvelopeAnnotation](#deleteenvelopeannotation)             | Delete envelope annotation        |
| [CreateTemplate](#createtemplate)                                 | Create new template               |
| [ListTemplates](#listtemplates)                                   | List templates                    |
| [GetTemplate](#gettemplate)                                       | Get template                      |
| [DeleteTemplate](#deletetemplate)                                 | Delete template                   |
| [DuplicateTemplate](#duplicatetemplate)                           | Duplicate template                |
| [AddTemplateDocument](#addtemplatedocument)                       | Add template document             |
| [GetTemplateDocument](#gettemplatedocument)                       | Get template document             |
| [GetTemplateDocuments](#gettemplatedocuments)                     | Get template documents            |
| [AddTemplateSigningSteps](#addtemplatesigningsteps)               | Add template signing steps        |
| [RenameTemplate](#renametemplate)                                 | Rename template                   |
| [SetTemplateComment](#settemplatecomment)                         | Set template comment              |
| [SetTemplateNotification](#settemplatenotification)               | Set template notification         |
| [GetTemplateAnnotations](#gettemplateannotations)                 | Get template annotations          |
| [GetDocumentTemplateAnnotations](#getdocumenttemplateannotations) | Get document template annotations |
| [AddTemplateAnnotation](#addtemplateannotation)                   | Add template annotation           |
| [DeleteTemplateAnnotation](#deletetemplateannotation)             | Delete template annotation        |
| [CreateWebhook](#createwebhook)                                   | Create webhook                    |
| [ListWebhooks](#listwebhooks)                                     | List webhooks                     |
| [DeleteWebhook](#deletewebhook)                                   | Delete webhook                    |

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
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)

envelopeLegalityLevel := signplus.ENVELOPE_LEGALITY_LEVEL_SES

request := signplus.CreateEnvelopeRequest{}
request.SetName("Name")
request.SetLegalityLevel(envelopeLegalityLevel)
request.SetExpiresAt(int64(123))
request.SetComment("Comment")
request.SetSandbox(true)

response, err := client.Signplus.CreateEnvelope(context.Background(), request)
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
| templateId                        | string                            | ✅       |                             |
| createEnvelopeFromTemplateRequest | CreateEnvelopeFromTemplateRequest | ✅       |                             |

**Return Type**

`Envelope`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)


request := signplus.CreateEnvelopeFromTemplateRequest{}
request.SetName("Name")
request.SetComment("Comment")
request.SetSandbox(true)

response, err := client.Signplus.CreateEnvelopeFromTemplate(context.Background(), "templateId", request)
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
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)

envelopeStatus := signplus.ENVELOPE_STATUS_DRAFT

envelopeOrderField := signplus.ENVELOPE_ORDER_FIELD_CREATION_DATE

request := signplus.ListEnvelopesRequest{}
request.SetName("Name")
request.SetTags([]string{})
request.SetComment("Comment")
request.SetIds([]string{})
request.SetStatuses([]signplus.EnvelopeStatus{envelopeStatus})
request.SetFolderIds([]string{})
request.SetOnlyRootFolder(true)
request.SetDateFrom(int64(123))
request.SetDateTo(int64(123))
request.SetUid("Uid")
request.SetFirst(int64(123))
request.SetLast(int64(123))
request.SetAfter("After")
request.SetBefore("Before")
request.SetOrderField(envelopeOrderField)
request.SetAscending(true)
request.SetIncludeTrash(true)

response, err := client.Signplus.ListEnvelopes(context.Background(), request)
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
| envelopeId | string  | ✅       |                             |

**Return Type**

`Envelope`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)

response, err := client.Signplus.GetEnvelope(context.Background(), "envelopeId")
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
| envelopeId | string  | ✅       |                             |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)

response, err := client.Signplus.DeleteEnvelope(context.Background(), "envelopeId")
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
| envelopeId | string  | ✅       |                             |
| documentId | string  | ✅       |                             |

**Return Type**

`Document`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)

response, err := client.Signplus.GetEnvelopeDocument(context.Background(), "envelopeId", "documentId")
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
| envelopeId | string  | ✅       |                             |

**Return Type**

`ListEnvelopeDocumentsResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)

response, err := client.Signplus.GetEnvelopeDocuments(context.Background(), "envelopeId")
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
| envelopeId                 | string                     | ✅       |                             |
| addEnvelopeDocumentRequest | AddEnvelopeDocumentRequest | ✅       |                             |

**Return Type**

`Document`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)


request := signplus.AddEnvelopeDocumentRequest{}
request.SetFile("")

response, err := client.Signplus.AddEnvelopeDocument(context.Background(), "envelopeId", request)
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
| envelopeId                      | string                          | ✅       |                             |
| setEnvelopeDynamicFieldsRequest | SetEnvelopeDynamicFieldsRequest | ✅       |                             |

**Return Type**

`Envelope`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)


dynamicField := signplus.DynamicField{}
dynamicField.SetName("Name")
dynamicField.SetValue("Value")

request := signplus.SetEnvelopeDynamicFieldsRequest{}
request.SetDynamicFields([]signplus.DynamicField{dynamicField})

response, err := client.Signplus.SetEnvelopeDynamicFields(context.Background(), "envelopeId", request)
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
| envelopeId                     | string                         | ✅       |                             |
| addEnvelopeSigningStepsRequest | AddEnvelopeSigningStepsRequest | ✅       |                             |

**Return Type**

`Envelope`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)

recipientRole := signplus.RECIPIENT_ROLE_SIGNER

recipientVerificationType := signplus.RECIPIENT_VERIFICATION_TYPE_SMS

recipientVerification := signplus.RecipientVerification{}
recipientVerification.SetType_(recipientVerificationType)
recipientVerification.SetValue("Value")

recipient := signplus.Recipient{}
recipient.SetId("Id")
recipient.SetUid("Uid")
recipient.SetName("Name")
recipient.SetEmail("Email")
recipient.SetRole(recipientRole)
recipient.SetVerification(recipientVerification)

signingStep := signplus.SigningStep{}
signingStep.SetRecipients([]signplus.Recipient{recipient})

request := signplus.AddEnvelopeSigningStepsRequest{}
request.SetSigningSteps([]signplus.SigningStep{signingStep})

response, err := client.Signplus.AddEnvelopeSigningSteps(context.Background(), "envelopeId", request)
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
| envelopeId | string  | ✅       |                             |

**Return Type**

`Envelope`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)

response, err := client.Signplus.SendEnvelope(context.Background(), "envelopeId")
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
| envelopeId | string  | ✅       |                             |

**Return Type**

`Envelope`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)

response, err := client.Signplus.DuplicateEnvelope(context.Background(), "envelopeId")
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
| envelopeId | string  | ✅       |                             |

**Return Type**

`Envelope`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)

response, err := client.Signplus.VoidEnvelope(context.Background(), "envelopeId")
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
| envelopeId            | string                | ✅       |                             |
| renameEnvelopeRequest | RenameEnvelopeRequest | ✅       |                             |

**Return Type**

`Envelope`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)


request := signplus.RenameEnvelopeRequest{}
request.SetName("Name")

response, err := client.Signplus.RenameEnvelope(context.Background(), "envelopeId", request)
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
| envelopeId                | string                    | ✅       |                             |
| setEnvelopeCommentRequest | SetEnvelopeCommentRequest | ✅       |                             |

**Return Type**

`Envelope`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)


request := signplus.SetEnvelopeCommentRequest{}
request.SetComment("Comment")

response, err := client.Signplus.SetEnvelopeComment(context.Background(), "envelopeId", request)
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
| envelopeId           | string               | ✅       |                             |
| envelopeNotification | EnvelopeNotification | ✅       |                             |

**Return Type**

`Envelope`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)


request := signplus.EnvelopeNotification{}
request.SetSubject("Subject")
request.SetMessage("Message")
request.SetReminderInterval(int64(123))

response, err := client.Signplus.SetEnvelopeNotification(context.Background(), "envelopeId", request)
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
| envelopeId                   | string                       | ✅       |                             |
| setEnvelopeExpirationRequest | SetEnvelopeExpirationRequest | ✅       |                             |

**Return Type**

`Envelope`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)


request := signplus.SetEnvelopeExpirationRequest{}
request.SetExpiresAt(int64(123))

response, err := client.Signplus.SetEnvelopeExpirationDate(context.Background(), "envelopeId", request)
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
| envelopeId                      | string                          | ✅       |                             |
| setEnvelopeLegalityLevelRequest | SetEnvelopeLegalityLevelRequest | ✅       |                             |

**Return Type**

`Envelope`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)

envelopeLegalityLevel := signplus.ENVELOPE_LEGALITY_LEVEL_SES

request := signplus.SetEnvelopeLegalityLevelRequest{}
request.SetLegalityLevel(envelopeLegalityLevel)

response, err := client.Signplus.SetEnvelopeLegalityLevel(context.Background(), "envelopeId", request)
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
| envelopeId | string  | ✅       | ID of the envelope          |

**Return Type**

`[]Annotation`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)

response, err := client.Signplus.GetEnvelopeAnnotations(context.Background(), "envelopeId")
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
| envelopeId | string  | ✅       | ID of the envelope          |
| documentId | string  | ✅       | ID of document              |

**Return Type**

`ListEnvelopeDocumentAnnotationsResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)

response, err := client.Signplus.GetEnvelopeDocumentAnnotations(context.Background(), "envelopeId", "documentId")
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
| envelopeId           | string               | ✅       | ID of the envelope          |
| addAnnotationRequest | AddAnnotationRequest | ✅       |                             |

**Return Type**

`Annotation`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)

annotationType := signplus.ANNOTATION_TYPE_TEXT


annotationSignature := signplus.AnnotationSignature{}
annotationSignature.SetId("Id")


annotationInitials := signplus.AnnotationInitials{}
annotationInitials.SetId("Id")

annotationFontFamily := signplus.ANNOTATION_FONT_FAMILY_UNKNOWN

annotationFont := signplus.AnnotationFont{}
annotationFont.SetFamily(annotationFontFamily)
annotationFont.SetItalic(true)
annotationFont.SetBold(true)

annotationText := signplus.AnnotationText{}
annotationText.SetSize(float64(123))
annotationText.SetColor(float64(123))
annotationText.SetValue("Value")
annotationText.SetTooltip("Tooltip")
annotationText.SetDynamicFieldName("DynamicFieldName")
annotationText.SetFont(annotationFont)

annotationFontFamily := signplus.ANNOTATION_FONT_FAMILY_UNKNOWN

annotationFont := signplus.AnnotationFont{}
annotationFont.SetFamily(annotationFontFamily)
annotationFont.SetItalic(true)
annotationFont.SetBold(true)

annotationDateTimeFormat := signplus.ANNOTATION_DATE_TIME_FORMAT_DMY_NUMERIC_SLASH

annotationDateTime := signplus.AnnotationDateTime{}
annotationDateTime.SetSize(float64(123))
annotationDateTime.SetFont(annotationFont)
annotationDateTime.SetColor("Color")
annotationDateTime.SetAutoFill(true)
annotationDateTime.SetTimezone("Timezone")
annotationDateTime.SetTimestamp(int64(123))
annotationDateTime.SetFormat(annotationDateTimeFormat)

annotationCheckboxStyle := signplus.ANNOTATION_CHECKBOX_STYLE_CIRCLE_CHECK

annotationCheckbox := signplus.AnnotationCheckbox{}
annotationCheckbox.SetChecked(true)
annotationCheckbox.SetStyle(annotationCheckboxStyle)

request := signplus.AddAnnotationRequest{}
request.SetRecipientId("RecipientId")
request.SetDocumentId("DocumentId")
request.SetPage(int64(123))
request.SetX(float64(123))
request.SetY(float64(123))
request.SetWidth(float64(123))
request.SetHeight(float64(123))
request.SetRequired(true)
request.SetType_(annotationType)
request.SetSignature(annotationSignature)
request.SetInitials(annotationInitials)
request.SetText(annotationText)
request.SetDatetime(annotationDateTime)
request.SetCheckbox(annotationCheckbox)

response, err := client.Signplus.AddEnvelopeAnnotation(context.Background(), "envelopeId", request)
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
| envelopeId   | string  | ✅       | ID of the envelope             |
| annotationId | string  | ✅       | ID of the annotation to delete |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)

response, err := client.Signplus.DeleteEnvelopeAnnotation(context.Background(), "envelopeId", "annotationId")
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
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)


request := signplus.CreateTemplateRequest{}
request.SetName("Name")

response, err := client.Signplus.CreateTemplate(context.Background(), request)
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
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)

templateOrderField := signplus.TEMPLATE_ORDER_FIELD_TEMPLATE_ID

request := signplus.ListTemplatesRequest{}
request.SetName("Name")
request.SetTags([]string{})
request.SetIds([]string{})
request.SetFirst(int64(123))
request.SetLast(int64(123))
request.SetAfter("After")
request.SetBefore("Before")
request.SetOrderField(templateOrderField)
request.SetAscending(true)

response, err := client.Signplus.ListTemplates(context.Background(), request)
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
| templateId | string  | ✅       |                             |

**Return Type**

`Template`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)

response, err := client.Signplus.GetTemplate(context.Background(), "templateId")
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
| templateId | string  | ✅       |                             |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)

response, err := client.Signplus.DeleteTemplate(context.Background(), "templateId")
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
| templateId | string  | ✅       |                             |

**Return Type**

`Template`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)

response, err := client.Signplus.DuplicateTemplate(context.Background(), "templateId")
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
| templateId                 | string                     | ✅       |                             |
| addTemplateDocumentRequest | AddTemplateDocumentRequest | ✅       |                             |

**Return Type**

`Document`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)


request := signplus.AddTemplateDocumentRequest{}
request.SetFile("")

response, err := client.Signplus.AddTemplateDocument(context.Background(), "templateId", request)
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
| templateId | string  | ✅       |                             |
| documentId | string  | ✅       |                             |

**Return Type**

`Document`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)

response, err := client.Signplus.GetTemplateDocument(context.Background(), "templateId", "documentId")
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
| templateId | string  | ✅       |                             |

**Return Type**

`ListTemplateDocumentsResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)

response, err := client.Signplus.GetTemplateDocuments(context.Background(), "templateId")
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
| templateId                     | string                         | ✅       |                             |
| addTemplateSigningStepsRequest | AddTemplateSigningStepsRequest | ✅       |                             |

**Return Type**

`Template`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)

templateRecipientRole := signplus.TEMPLATE_RECIPIENT_ROLE_SIGNER

templateRecipient := signplus.TemplateRecipient{}
templateRecipient.SetId("Id")
templateRecipient.SetUid("Uid")
templateRecipient.SetName("Name")
templateRecipient.SetEmail("Email")
templateRecipient.SetRole(templateRecipientRole)

templateSigningStep := signplus.TemplateSigningStep{}
templateSigningStep.SetRecipients([]signplus.TemplateRecipient{templateRecipient})

request := signplus.AddTemplateSigningStepsRequest{}
request.SetSigningSteps([]signplus.TemplateSigningStep{templateSigningStep})

response, err := client.Signplus.AddTemplateSigningSteps(context.Background(), "templateId", request)
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
| templateId            | string                | ✅       |                             |
| renameTemplateRequest | RenameTemplateRequest | ✅       |                             |

**Return Type**

`Template`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)


request := signplus.RenameTemplateRequest{}
request.SetName("Name")

response, err := client.Signplus.RenameTemplate(context.Background(), "templateId", request)
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
| templateId                | string                    | ✅       |                             |
| setTemplateCommentRequest | SetTemplateCommentRequest | ✅       |                             |

**Return Type**

`Template`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)


request := signplus.SetTemplateCommentRequest{}
request.SetComment("Comment")

response, err := client.Signplus.SetTemplateComment(context.Background(), "templateId", request)
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
| templateId           | string               | ✅       |                             |
| envelopeNotification | EnvelopeNotification | ✅       |                             |

**Return Type**

`Template`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)


request := signplus.EnvelopeNotification{}
request.SetSubject("Subject")
request.SetMessage("Message")
request.SetReminderInterval(int64(123))

response, err := client.Signplus.SetTemplateNotification(context.Background(), "templateId", request)
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
| templateId | string  | ✅       | ID of the template          |

**Return Type**

`ListTemplateAnnotationsResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)

response, err := client.Signplus.GetTemplateAnnotations(context.Background(), "templateId")
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
| templateId | string  | ✅       | ID of the template          |
| documentId | string  | ✅       | ID of document              |

**Return Type**

`ListTemplateDocumentAnnotationsResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)

response, err := client.Signplus.GetDocumentTemplateAnnotations(context.Background(), "templateId", "documentId")
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
| templateId           | string               | ✅       | ID of the template          |
| addAnnotationRequest | AddAnnotationRequest | ✅       |                             |

**Return Type**

`Annotation`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)

annotationType := signplus.ANNOTATION_TYPE_TEXT


annotationSignature := signplus.AnnotationSignature{}
annotationSignature.SetId("Id")


annotationInitials := signplus.AnnotationInitials{}
annotationInitials.SetId("Id")

annotationFontFamily := signplus.ANNOTATION_FONT_FAMILY_UNKNOWN

annotationFont := signplus.AnnotationFont{}
annotationFont.SetFamily(annotationFontFamily)
annotationFont.SetItalic(true)
annotationFont.SetBold(true)

annotationText := signplus.AnnotationText{}
annotationText.SetSize(float64(123))
annotationText.SetColor(float64(123))
annotationText.SetValue("Value")
annotationText.SetTooltip("Tooltip")
annotationText.SetDynamicFieldName("DynamicFieldName")
annotationText.SetFont(annotationFont)

annotationFontFamily := signplus.ANNOTATION_FONT_FAMILY_UNKNOWN

annotationFont := signplus.AnnotationFont{}
annotationFont.SetFamily(annotationFontFamily)
annotationFont.SetItalic(true)
annotationFont.SetBold(true)

annotationDateTimeFormat := signplus.ANNOTATION_DATE_TIME_FORMAT_DMY_NUMERIC_SLASH

annotationDateTime := signplus.AnnotationDateTime{}
annotationDateTime.SetSize(float64(123))
annotationDateTime.SetFont(annotationFont)
annotationDateTime.SetColor("Color")
annotationDateTime.SetAutoFill(true)
annotationDateTime.SetTimezone("Timezone")
annotationDateTime.SetTimestamp(int64(123))
annotationDateTime.SetFormat(annotationDateTimeFormat)

annotationCheckboxStyle := signplus.ANNOTATION_CHECKBOX_STYLE_CIRCLE_CHECK

annotationCheckbox := signplus.AnnotationCheckbox{}
annotationCheckbox.SetChecked(true)
annotationCheckbox.SetStyle(annotationCheckboxStyle)

request := signplus.AddAnnotationRequest{}
request.SetRecipientId("RecipientId")
request.SetDocumentId("DocumentId")
request.SetPage(int64(123))
request.SetX(float64(123))
request.SetY(float64(123))
request.SetWidth(float64(123))
request.SetHeight(float64(123))
request.SetRequired(true)
request.SetType_(annotationType)
request.SetSignature(annotationSignature)
request.SetInitials(annotationInitials)
request.SetText(annotationText)
request.SetDatetime(annotationDateTime)
request.SetCheckbox(annotationCheckbox)

response, err := client.Signplus.AddTemplateAnnotation(context.Background(), "templateId", request)
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
| templateId   | string  | ✅       | ID of the template             |
| annotationId | string  | ✅       | ID of the annotation to delete |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)

response, err := client.Signplus.DeleteTemplateAnnotation(context.Background(), "templateId", "annotationId")
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
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)

webhookEvent := signplus.WEBHOOK_EVENT_ENVELOPE_EXPIRED

request := signplus.CreateWebhookRequest{}
request.SetEvent(webhookEvent)
request.SetTarget("Target")

response, err := client.Signplus.CreateWebhook(context.Background(), request)
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
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)

webhookEvent := signplus.WEBHOOK_EVENT_ENVELOPE_EXPIRED

request := signplus.ListWebhooksRequest{}
request.SetWebhookId("WebhookId")
request.SetEvent(webhookEvent)

response, err := client.Signplus.ListWebhooks(context.Background(), request)
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
| webhookId | string  | ✅       |                             |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/alohihq/signplus-go/pkg/signplusconfig"
  "github.com/alohihq/signplus-go/pkg/signplus"
)

config := signplusconfig.NewConfig()
client := signplus.NewSignplus(config)

response, err := client.Signplus.DeleteWebhook(context.Background(), "webhookId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

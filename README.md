# Signplus Go SDK 3.0.0

Welcome to the Signplus SDK documentation. This guide will help you get started with integrating and using the Signplus SDK in your project.

## Versions

- API version: `2.5.0`
- SDK version: `3.0.0`

## About the API

Integrate legally-binding electronic signature to your workflow

Contact Support:
Name: Sign.Plus
Email: support@alohi.com

## Table of Contents

- [Setup & Configuration](#setup--configuration)
  - [Supported Language Versions](#supported-language-versions)
- [Authentication](#authentication)
  - [Access Token Authentication](#access-token-authentication)
- [Setting a Custom Timeout](#setting-a-custom-timeout)
- [Sample Usage](#sample-usage)
- [Services](#services)
  - [Response Wrappers](#response-wrappers)
- [Models](#models)
- [License](#license)

# Setup & Configuration

## Supported Language Versions

This SDK is compatible with the following versions: `Go >= 1.19.0`

## Authentication

### Access Token Authentication

The signplus API uses an Access Token for authentication.

This token must be provided to authenticate your requests to the API.

#### Setting the Access Token

When you initialize the SDK, you can set the access token as follows:

```go
import (
    "github.com/alohihq/signplus-go"
  )

config := signplus.NewConfig()
config.SetAccessToken("YOUR-TOKEN")

sdk := signplus.NewSignplus(config)
```

If you need to set or update the access token after initializing the SDK, you can use:

```go
import (
    "github.com/alohihq/signplus-go"
  )

config := signplus.NewConfig()

sdk := signplus.NewSignplus(config)
sdk.SetAccessToken("YOUR-TOKEN")
```

## Setting a Custom Timeout

You can set a custom timeout for the SDK's HTTP requests as follows:

```go
import "time"

config := signplus.NewConfig()

sdk := signplus.NewSignplus(config)

sdk.SetTimeout(10 * time.Second)
```

# Sample Usage

Below is a comprehensive example demonstrating how to authenticate and call a simple endpoint:

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

response, err := client.EnvelopeID.DeleteEnvelope(context.Background(), "envelope_id")
if err != nil {
  panic(err)
}

fmt.Println(response)

```

## Services

The SDK provides various services to interact with the API.

<details>
<summary>Below is a list of all available services with links to their detailed documentation:</summary>

| Name                                                                                                                 |
| :------------------------------------------------------------------------------------------------------------------- |
| [TemplateID](documentation/services/template_id.md)                                                                  |
| [SignedDocuments](documentation/services/signed_documents.md)                                                        |
| [Certificate](documentation/services/certificate.md)                                                                 |
| [DocumentID](documentation/services/document_id.md)                                                                  |
| [Document](documentation/services/document.md)                                                                       |
| [Documents](documentation/services/documents.md)                                                                     |
| [DynamicFields](documentation/services/dynamic_fields.md)                                                            |
| [SigningSteps](documentation/services/signing_steps.md)                                                              |
| [Settings](documentation/services/settings.md)                                                                       |
| [Placeholders](documentation/services/placeholders.md)                                                               |
| [FileID](documentation/services/file_id.md)                                                                          |
| [Send](documentation/services/send.md)                                                                               |
| [Duplicate](documentation/services/duplicate.md)                                                                     |
| [Void](documentation/services/void.md)                                                                               |
| [Rename](documentation/services/rename.md)                                                                           |
| [SetComment](documentation/services/set_comment.md)                                                                  |
| [SetNotification](documentation/services/set_notification.md)                                                        |
| [SetExpirationDate](documentation/services/set_expiration_date.md)                                                   |
| [SetLegalityLevel](documentation/services/set_legality_level.md)                                                     |
| [EnvelopeEnvelopeIDAnnotationsDocumentID](documentation/services/envelope_envelope_id_annotations_document_id.md)    |
| [Annotations](documentation/services/annotations.md)                                                                 |
| [AnnotationID](documentation/services/annotation_id.md)                                                              |
| [Annotation](documentation/services/annotation.md)                                                                   |
| [EnvelopeID](documentation/services/envelope_id.md)                                                                  |
| [Envelope](documentation/services/envelope.md)                                                                       |
| [Envelopes](documentation/services/envelopes.md)                                                                     |
| [TemplateTemplateIDDuplicate](documentation/services/template_template_id_duplicate.md)                              |
| [TemplateTemplateIDDocumentDocumentID](documentation/services/template_template_id_document_document_id.md)          |
| [TemplateTemplateIDDocument](documentation/services/template_template_id_document.md)                                |
| [TemplateTemplateIDDocuments](documentation/services/template_template_id_documents.md)                              |
| [TemplateTemplateIDSigningSteps](documentation/services/template_template_id_signing_steps.md)                       |
| [TemplateTemplateIDRename](documentation/services/template_template_id_rename.md)                                    |
| [TemplateTemplateIDSetComment](documentation/services/template_template_id_set_comment.md)                           |
| [TemplateTemplateIDSetNotification](documentation/services/template_template_id_set_notification.md)                 |
| [TemplateTemplateIDAnnotationsDocumentID](documentation/services/template_template_id_annotations_document_id.md)    |
| [TemplateTemplateIDAnnotations](documentation/services/template_template_id_annotations.md)                          |
| [TemplateTemplateIDAnnotationAnnotationID](documentation/services/template_template_id_annotation_annotation_id.md)  |
| [TemplateTemplateIDAnnotation](documentation/services/template_template_id_annotation.md)                            |
| [TemplateTemplateIDAttachmentsSettings](documentation/services/template_template_id_attachments_settings.md)         |
| [TemplateTemplateIDAttachmentsPlaceholders](documentation/services/template_template_id_attachments_placeholders.md) |
| [TemplateTemplateID](documentation/services/template_template_id.md)                                                 |
| [Template](documentation/services/template.md)                                                                       |
| [Templates](documentation/services/templates.md)                                                                     |
| [WebhookID](documentation/services/webhook_id.md)                                                                    |
| [Webhook](documentation/services/webhook.md)                                                                         |
| [Webhooks](documentation/services/webhooks.md)                                                                       |

</details>

### Response Wrappers

All services use response wrappers to provide a consistent interface to return the responses from the API.

The response wrapper itself is a generic struct that contains the response data and metadata.

<details>
<summary>Below are the response wrappers used in the SDK:</summary>

#### `SignplusResponse[T]`

This response wrapper is used to return the response data from the API. It contains the following fields:

| Name     | Type                       | Description                                 |
| :------- | :------------------------- | :------------------------------------------ |
| Data     | `T`                        | The body of the API response                |
| Metadata | `SignplusResponseMetadata` | Status code and headers returned by the API |

#### `SignplusError[T]`

This response wrapper is used to return an error. It contains the following fields:

| Name     | Type                    | Description                                                       |
| :------- | :---------------------- | :---------------------------------------------------------------- |
| Err      | `error`                 | The error that occurred                                           |
| Data     | `*T`                    | The deserialized error response data (nil if unmarshaling failed) |
| Body     | `[]byte`                | The raw body of the API response                                  |
| Metadata | `SignplusErrorMetadata` | Status code and headers returned by the API                       |

#### `SignplusResponseMetadata`

This struct is shared by both response wrappers and contains the following fields:

| Name       | Type                | Description                                      |
| :--------- | :------------------ | :----------------------------------------------- |
| Headers    | `map[string]string` | A map containing the headers returned by the API |
| StatusCode | `int`               | The status code returned by the API              |

</details>

## Models

The SDK includes several models that represent the data structures used in API requests and responses. These models help in organizing and managing the data efficiently.

<details>
<summary>Below is a list of all available models with links to their detailed documentation:</summary>

| Name                                                                                                                                        | Description |
| :------------------------------------------------------------------------------------------------------------------------------------------ | :---------- |
| [CreateEnvelopeFromTemplateRequest](documentation/models/create_envelope_from_template_request.md)                                          |             |
| [AddEnvelopeDocumentRequest](documentation/models/add_envelope_document_request.md)                                                         |             |
| [SetEnvelopeDynamicFieldsRequest](documentation/models/set_envelope_dynamic_fields_request.md)                                              |             |
| [DynamicFields](documentation/models/dynamic_fields.md)                                                                                     |             |
| [AddEnvelopeSigningStepsRequest](documentation/models/add_envelope_signing_steps_request.md)                                                |             |
| [AddEnvelopeSigningStepsRequestSigningSteps](documentation/models/add_envelope_signing_steps_request_signing_steps.md)                      |             |
| [SigningStepsRecipients1](documentation/models/signing_steps_recipients_1.md)                                                               |             |
| [Verification](documentation/models/verification.md)                                                                                        |             |
| [SetEnvelopeAttachmentsSettingsRequest](documentation/models/set_envelope_attachments_settings_request.md)                                  |             |
| [SetEnvelopeAttachmentsSettingsRequestSettings](documentation/models/set_envelope_attachments_settings_request_settings.md)                 |             |
| [SetEnvelopeAttachmentsPlaceholdersRequest](documentation/models/set_envelope_attachments_placeholders_request.md)                          |             |
| [SetEnvelopeAttachmentsPlaceholdersRequestPlaceholders](documentation/models/set_envelope_attachments_placeholders_request_placeholders.md) |             |
| [RenameEnvelopeRequest](documentation/models/rename_envelope_request.md)                                                                    |             |
| [SetEnvelopeCommentRequest](documentation/models/set_envelope_comment_request.md)                                                           |             |
| [SetEnvelopeNotificationRequest](documentation/models/set_envelope_notification_request.md)                                                 |             |
| [SetEnvelopeExpirationDateRequest](documentation/models/set_envelope_expiration_date_request.md)                                            |             |
| [SetEnvelopeLegalityLevelRequest](documentation/models/set_envelope_legality_level_request.md)                                              |             |
| [AddEnvelopeAnnotationRequest](documentation/models/add_envelope_annotation_request.md)                                                     |             |
| [AddEnvelopeAnnotationRequestSignature](documentation/models/add_envelope_annotation_request_signature.md)                                  |             |
| [AddEnvelopeAnnotationRequestInitials](documentation/models/add_envelope_annotation_request_initials.md)                                    |             |
| [AddEnvelopeAnnotationRequestText](documentation/models/add_envelope_annotation_request_text.md)                                            |             |
| [TextFont1](documentation/models/text_font_1.md)                                                                                            |             |
| [AddEnvelopeAnnotationRequestDatetime](documentation/models/add_envelope_annotation_request_datetime.md)                                    |             |
| [DatetimeFont1](documentation/models/datetime_font_1.md)                                                                                    |             |
| [AddEnvelopeAnnotationRequestCheckbox](documentation/models/add_envelope_annotation_request_checkbox.md)                                    |             |
| [CreateEnvelopeRequest](documentation/models/create_envelope_request.md)                                                                    |             |
| [ListEnvelopesRequest](documentation/models/list_envelopes_request.md)                                                                      |             |
| [AddTemplateDocumentRequest](documentation/models/add_template_document_request.md)                                                         |             |
| [AddTemplateSigningStepsRequest](documentation/models/add_template_signing_steps_request.md)                                                |             |
| [AddTemplateSigningStepsRequestSigningSteps](documentation/models/add_template_signing_steps_request_signing_steps.md)                      |             |
| [SigningStepsRecipients2](documentation/models/signing_steps_recipients_2.md)                                                               |             |
| [RenameTemplateRequest](documentation/models/rename_template_request.md)                                                                    |             |
| [SetTemplateCommentRequest](documentation/models/set_template_comment_request.md)                                                           |             |
| [SetTemplateNotificationRequest](documentation/models/set_template_notification_request.md)                                                 |             |
| [AddTemplateAnnotationRequest](documentation/models/add_template_annotation_request.md)                                                     |             |
| [AddTemplateAnnotationRequestSignature](documentation/models/add_template_annotation_request_signature.md)                                  |             |
| [AddTemplateAnnotationRequestInitials](documentation/models/add_template_annotation_request_initials.md)                                    |             |
| [AddTemplateAnnotationRequestText](documentation/models/add_template_annotation_request_text.md)                                            |             |
| [TextFont2](documentation/models/text_font_2.md)                                                                                            |             |
| [AddTemplateAnnotationRequestDatetime](documentation/models/add_template_annotation_request_datetime.md)                                    |             |
| [DatetimeFont2](documentation/models/datetime_font_2.md)                                                                                    |             |
| [AddTemplateAnnotationRequestCheckbox](documentation/models/add_template_annotation_request_checkbox.md)                                    |             |
| [SetTemplateAttachmentsSettingsRequest](documentation/models/set_template_attachments_settings_request.md)                                  |             |
| [SetTemplateAttachmentsSettingsRequestSettings](documentation/models/set_template_attachments_settings_request_settings.md)                 |             |
| [SetTemplateAttachmentsPlaceholdersRequest](documentation/models/set_template_attachments_placeholders_request.md)                          |             |
| [SetTemplateAttachmentsPlaceholdersRequestPlaceholders](documentation/models/set_template_attachments_placeholders_request_placeholders.md) |             |
| [CreateTemplateRequest](documentation/models/create_template_request.md)                                                                    |             |
| [ListTemplatesRequest](documentation/models/list_templates_request.md)                                                                      |             |
| [CreateWebhookRequest](documentation/models/create_webhook_request.md)                                                                      |             |
| [ListWebhooksRequest](documentation/models/list_webhooks_request.md)                                                                        |             |

</details>

## License

This SDK is licensed under the MIT License.

See the [LICENSE](LICENSE) file for more details.

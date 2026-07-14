# Signplus Go SDK 1.0.0

Welcome to the Signplus SDK documentation. This guide will help you get started with integrating and using the Signplus SDK in your project.

## Versions

- API version: `2.5.0`
- SDK version: `1.0.0`

## About the API

Integrate legally-binding electronic signature to your workflow

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

response, err := client.Signplus1.GetEnvelope(context.Background(), "envelope_id")
if err != nil {
  panic(err)
}

fmt.Println(response)

```

## Services

The SDK provides various services to interact with the API.

<details>
<summary>Below is a list of all available services with links to their detailed documentation:</summary>

| Name                                            |
| :---------------------------------------------- |
| [Signplus1](documentation/services/signplus.md) |

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

| Name                                                                                                               | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| :----------------------------------------------------------------------------------------------------------------- | :---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [Envelope](documentation/models/envelope.md)                                                                       |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [EnvelopeFlowType](documentation/models/envelope_flow_type.md)                                                     | Flow type of the envelope (REQUEST_SIGNATURE is a request for signature, SIGN_MYSELF is a self-signing flow)                                                                                                                                                                                                                                                                                                                                                                              |
| [EnvelopeLegalityLevel](documentation/models/envelope_legality_level.md)                                           | Legal level of the envelope (SES is Simple Electronic Signature, QES_EIDAS is Qualified Electronic Signature, QES_ZERTES is Qualified Electronic Signature with Zertes)                                                                                                                                                                                                                                                                                                                   |
| [EnvelopeStatus](documentation/models/envelope_status.md)                                                          | Status of the envelope                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| [SigningStep](documentation/models/signing_step.md)                                                                |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [Recipient](documentation/models/recipient.md)                                                                     |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [RecipientRole](documentation/models/recipient_role.md)                                                            | Role of the recipient (SIGNER signs the document, RECEIVES_COPY receives a copy of the document, IN_PERSON_SIGNER signs the document in person, SENDER sends the document)                                                                                                                                                                                                                                                                                                                |
| [RecipientVerification](documentation/models/recipient_verification.md)                                            |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [RecipientVerificationType](documentation/models/recipient_verification_type.md)                                   | Type of verification the recipient must complete before accessing the envelope. - `PASSCODE`: requires a code to be entered. - `SMS`: sends a code via SMS. - `ID_VERIFICATION`: prompts the recipient to complete an automated ID and selfie check.                                                                                                                                                                                                                                      |
| [Document](documentation/models/document.md)                                                                       |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [Page](documentation/models/page.md)                                                                               |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [EnvelopeNotification](documentation/models/envelope_notification.md)                                              |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [EnvelopeAttachments](documentation/models/envelope_attachments.md)                                                |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [AttachmentSettings](documentation/models/attachment_settings.md)                                                  |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [AttachmentPlaceholdersPerRecipient](documentation/models/attachment_placeholders_per_recipient.md)                |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [AttachmentPlaceholder](documentation/models/attachment_placeholder.md)                                            |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [AttachmentPlaceholderFile](documentation/models/attachment_placeholder_file.md)                                   |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [CreateEnvelopeRequest](documentation/models/create_envelope_request.md)                                           |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [CreateEnvelopeFromTemplateRequest](documentation/models/create_envelope_from_template_request.md)                 |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [ListEnvelopesResponse](documentation/models/list_envelopes_response.md)                                           |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [ListEnvelopesRequest](documentation/models/list_envelopes_request.md)                                             |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [EnvelopeOrderField](documentation/models/envelope_order_field.md)                                                 | Field to order envelopes by                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| [ListEnvelopeDocumentsResponse](documentation/models/list_envelope_documents_response.md)                          |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [AddEnvelopeDocumentRequest](documentation/models/add_envelope_document_request.md)                                |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [SetEnvelopeDynamicFieldsRequest](documentation/models/set_envelope_dynamic_fields_request.md)                     |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [DynamicField](documentation/models/dynamic_field.md)                                                              |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [AddEnvelopeSigningStepsRequest](documentation/models/add_envelope_signing_steps_request.md)                       |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [SetEnvelopeAttachmentsSettingsRequest](documentation/models/set_envelope_attachments_settings_request.md)         |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [SetEnvelopeAttachmentsPlaceholdersRequest](documentation/models/set_envelope_attachments_placeholders_request.md) |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [AttachmentPlaceholderRequest](documentation/models/attachment_placeholder_request.md)                             |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [RenameEnvelopeRequest](documentation/models/rename_envelope_request.md)                                           |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [SetEnvelopeCommentRequest](documentation/models/set_envelope_comment_request.md)                                  |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [SetEnvelopeExpirationRequest](documentation/models/set_envelope_expiration_request.md)                            |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [SetEnvelopeLegalityLevelRequest](documentation/models/set_envelope_legality_level_request.md)                     |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [Annotation](documentation/models/annotation.md)                                                                   |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [AnnotationType](documentation/models/annotation_type.md)                                                          | Type of the annotation                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| [AnnotationSignature](documentation/models/annotation_signature.md)                                                | Signature annotation (null if annotation is not a signature)                                                                                                                                                                                                                                                                                                                                                                                                                              |
| [AnnotationInitials](documentation/models/annotation_initials.md)                                                  | Initials annotation (null if annotation is not initials)                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| [AnnotationText](documentation/models/annotation_text.md)                                                          | Text annotation (null if annotation is not a text)                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| [AnnotationFont](documentation/models/annotation_font.md)                                                          |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [AnnotationFontFamily](documentation/models/annotation_font_family.md)                                             | Font family of the text                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| [AnnotationDateTime](documentation/models/annotation_date_time.md)                                                 | Date annotation (null if annotation is not a date)                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| [AnnotationDateTimeFormat](documentation/models/annotation_date_time_format.md)                                    | Format of the date time (DMY_NUMERIC_SLASH is day/month/year with slashes, MDY_NUMERIC_SLASH is month/day/year with slashes, YMD_NUMERIC_SLASH is year/month/day with slashes, DMY_NUMERIC_DASH_SHORT is day/month/year with dashes, DMY_NUMERIC_DASH is day/month/year with dashes, YMD_NUMERIC_DASH is year/month/day with dashes, MDY_TEXT_DASH_SHORT is month/day/year with dashes, MDY_TEXT_SPACE_SHORT is month/day/year with spaces, MDY_TEXT_SPACE is month/day/year with spaces) |
| [AnnotationCheckbox](documentation/models/annotation_checkbox.md)                                                  | Checkbox annotation (null if annotation is not a checkbox)                                                                                                                                                                                                                                                                                                                                                                                                                                |
| [AnnotationCheckboxStyle](documentation/models/annotation_checkbox_style.md)                                       | Style of the checkbox                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| [ListEnvelopeDocumentAnnotationsResponse](documentation/models/list_envelope_document_annotations_response.md)     |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [AddAnnotationRequest](documentation/models/add_annotation_request.md)                                             |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [Template](documentation/models/template.md)                                                                       |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [TemplateSigningStep](documentation/models/template_signing_step.md)                                               |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [TemplateRecipient](documentation/models/template_recipient.md)                                                    |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [TemplateRecipientRole](documentation/models/template_recipient_role.md)                                           | Role of the recipient (SIGNER signs the document, RECEIVES_COPY receives a copy of the document, IN_PERSON_SIGNER signs the document in person, SENDER sends the document)                                                                                                                                                                                                                                                                                                                |
| [CreateTemplateRequest](documentation/models/create_template_request.md)                                           |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [ListTemplatesResponse](documentation/models/list_templates_response.md)                                           |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [ListTemplatesRequest](documentation/models/list_templates_request.md)                                             |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [TemplateOrderField](documentation/models/template_order_field.md)                                                 | Field to order templates by                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| [AddTemplateDocumentRequest](documentation/models/add_template_document_request.md)                                |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [ListTemplateDocumentsResponse](documentation/models/list_template_documents_response.md)                          |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [AddTemplateSigningStepsRequest](documentation/models/add_template_signing_steps_request.md)                       |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [RenameTemplateRequest](documentation/models/rename_template_request.md)                                           |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [SetTemplateCommentRequest](documentation/models/set_template_comment_request.md)                                  |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [ListTemplateAnnotationsResponse](documentation/models/list_template_annotations_response.md)                      |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [ListTemplateDocumentAnnotationsResponse](documentation/models/list_template_document_annotations_response.md)     |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [Webhook](documentation/models/webhook.md)                                                                         |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [WebhookEvent](documentation/models/webhook_event.md)                                                              | Event of the webhook                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| [CreateWebhookRequest](documentation/models/create_webhook_request.md)                                             |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [ListWebhooksResponse](documentation/models/list_webhooks_response.md)                                             |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [ListWebhooksRequest](documentation/models/list_webhooks_request.md)                                               |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |

</details>

## License

This SDK is licensed under the MIT License.

See the [LICENSE](LICENSE) file for more details.

# TemplateTemplateIDAttachmentsSettings

A list of all methods in the `TemplateTemplateIDAttachmentsSettings` service. Click on the method name to view detailed information about that method.

| Methods                                                           | Description                      |
| :---------------------------------------------------------------- | :------------------------------- |
| [SetTemplateAttachmentsSettings](#settemplateattachmentssettings) | Set template attachment settings |

## SetTemplateAttachmentsSettings

Set template attachment settings

- HTTP Method: `PUT`
- Endpoint: `/template/{template_id}/attachments/settings`

**Parameters**

| Name                                  | Type                                        | Required | Description                   |
| :------------------------------------ | :------------------------------------------ | :------- | :---------------------------- |
| ctx                                   | Context                                     | ✅       | Default go language context   |
| templateID                            | string                                      | ✅       |                               |
| setTemplateAttachmentsSettingsRequest | SetTemplateAttachmentsSettingsRequest       | ✅       |                               |
| params                                | SetTemplateAttachmentsSettingsRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/templatetemplateidattachmentssettings"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := templatetemplateidattachmentssettings.SetTemplateAttachmentsSettingsRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}


setTemplateAttachmentsSettingsRequestSettings := templatetemplateidattachmentssettings.SetTemplateAttachmentsSettingsRequestSettings{
  VisibleToRecipients: signplus.Nullable[string]("<boolean>"),
}

request := templatetemplateidattachmentssettings.SetTemplateAttachmentsSettingsRequest{
  Settings: signplus.Nullable[templatetemplateidattachmentssettings.SetTemplateAttachmentsSettingsRequestSettings](setTemplateAttachmentsSettingsRequestSettings),
}

response, err := client.TemplateTemplateIDAttachmentsSettings.SetTemplateAttachmentsSettings(context.Background(), "template_id", request, params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

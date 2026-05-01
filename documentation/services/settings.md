# Settings

A list of all methods in the `Settings` service. Click on the method name to view detailed information about that method.

| Methods                                                           | Description                      |
| :---------------------------------------------------------------- | :------------------------------- |
| [SetEnvelopeAttachmentsSettings](#setenvelopeattachmentssettings) | Set envelope attachment settings |

## SetEnvelopeAttachmentsSettings

Set envelope attachment settings

- HTTP Method: `PUT`
- Endpoint: `/envelope/{envelope_id}/attachments/settings`

**Parameters**

| Name                                  | Type                                        | Required | Description                   |
| :------------------------------------ | :------------------------------------------ | :------- | :---------------------------- |
| ctx                                   | Context                                     | ✅       | Default go language context   |
| envelopeID                            | string                                      | ✅       |                               |
| setEnvelopeAttachmentsSettingsRequest | SetEnvelopeAttachmentsSettingsRequest       | ✅       |                               |
| params                                | SetEnvelopeAttachmentsSettingsRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/settings"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := settings.SetEnvelopeAttachmentsSettingsRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}


setEnvelopeAttachmentsSettingsRequestSettings := settings.SetEnvelopeAttachmentsSettingsRequestSettings{
  VisibleToRecipients: signplus.Nullable[string]("<boolean>"),
}

request := settings.SetEnvelopeAttachmentsSettingsRequest{
  Settings: signplus.Nullable[settings.SetEnvelopeAttachmentsSettingsRequestSettings](setEnvelopeAttachmentsSettingsRequestSettings),
}

response, err := client.Settings.SetEnvelopeAttachmentsSettings(context.Background(), "envelope_id", request, params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

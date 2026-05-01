# Webhooks

A list of all methods in the `Webhooks` service. Click on the method name to view detailed information about that method.

| Methods                       | Description   |
| :---------------------------- | :------------ |
| [ListWebhooks](#listwebhooks) | List webhooks |

## ListWebhooks

List webhooks

- HTTP Method: `POST`
- Endpoint: `/webhooks`

**Parameters**

| Name                | Type                      | Required | Description                   |
| :------------------ | :------------------------ | :------- | :---------------------------- |
| ctx                 | Context                   | ✅       | Default go language context   |
| listWebhooksRequest | ListWebhooksRequest       | ✅       |                               |
| params              | ListWebhooksRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/webhooks"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := webhooks.ListWebhooksRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}


request := webhooks.ListWebhooksRequest{
  WebhookID: signplus.Nullable[string]("<string>"),
  Event: signplus.Nullable[string]("ENVELOPE_COMPLETED"),
}

response, err := client.Webhooks.ListWebhooks(context.Background(), request, params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

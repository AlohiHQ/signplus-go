# Webhook

A list of all methods in the `Webhook` service. Click on the method name to view detailed information about that method.

| Methods                         | Description    |
| :------------------------------ | :------------- |
| [CreateWebhook](#createwebhook) | Create webhook |

## CreateWebhook

Create webhook

- HTTP Method: `POST`
- Endpoint: `/webhook`

**Parameters**

| Name                 | Type                       | Required | Description                   |
| :------------------- | :------------------------- | :------- | :---------------------------- |
| ctx                  | Context                    | ✅       | Default go language context   |
| createWebhookRequest | CreateWebhookRequest       | ✅       |                               |
| params               | CreateWebhookRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/webhook"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := webhook.CreateWebhookRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}


request := webhook.CreateWebhookRequest{
  Event: signplus.Nullable[string]("ENVELOPE_AUDIT_TRAIL"),
  Target: signplus.Nullable[string]("string"),
}

response, err := client.Webhook.CreateWebhook(context.Background(), request, params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

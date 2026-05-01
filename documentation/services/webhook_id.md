# WebhookID

A list of all methods in the `WebhookID` service. Click on the method name to view detailed information about that method.

| Methods                         | Description    |
| :------------------------------ | :------------- |
| [DeleteWebhook](#deletewebhook) | Delete webhook |

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

response, err := client.WebhookID.DeleteWebhook(context.Background(), "webhook_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

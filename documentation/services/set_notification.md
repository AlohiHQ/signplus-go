# SetNotification

A list of all methods in the `SetNotification` service. Click on the method name to view detailed information about that method.

| Methods                                             | Description               |
| :-------------------------------------------------- | :------------------------ |
| [SetEnvelopeNotification](#setenvelopenotification) | Set envelope notification |

## SetEnvelopeNotification

Set envelope notification

- HTTP Method: `PUT`
- Endpoint: `/envelope/{envelope_id}/set_notification`

**Parameters**

| Name                           | Type                                 | Required | Description                   |
| :----------------------------- | :----------------------------------- | :------- | :---------------------------- |
| ctx                            | Context                              | ✅       | Default go language context   |
| envelopeID                     | string                               | ✅       |                               |
| setEnvelopeNotificationRequest | SetEnvelopeNotificationRequest       | ✅       |                               |
| params                         | SetEnvelopeNotificationRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/setnotification"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := setnotification.SetEnvelopeNotificationRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}


request := setnotification.SetEnvelopeNotificationRequest{
  Subject: signplus.Nullable[string]("string"),
  Message: signplus.Nullable[string]("string"),
  ReminderInterval: signplus.Nullable[float64](float64(4732)),
}

response, err := client.SetNotification.SetEnvelopeNotification(context.Background(), "envelope_id", request, params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

# TemplateTemplateIDSetNotification

A list of all methods in the `TemplateTemplateIDSetNotification` service. Click on the method name to view detailed information about that method.

| Methods                                             | Description               |
| :-------------------------------------------------- | :------------------------ |
| [SetTemplateNotification](#settemplatenotification) | Set template notification |

## SetTemplateNotification

Set template notification

- HTTP Method: `PUT`
- Endpoint: `/template/{template_id}/set_notification`

**Parameters**

| Name                           | Type                                 | Required | Description                   |
| :----------------------------- | :----------------------------------- | :------- | :---------------------------- |
| ctx                            | Context                              | ✅       | Default go language context   |
| templateID                     | string                               | ✅       |                               |
| setTemplateNotificationRequest | SetTemplateNotificationRequest       | ✅       |                               |
| params                         | SetTemplateNotificationRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/templatetemplateidsetnotification"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := templatetemplateidsetnotification.SetTemplateNotificationRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}


request := templatetemplateidsetnotification.SetTemplateNotificationRequest{
  Subject: signplus.Nullable[string]("string"),
  Message: signplus.Nullable[string]("string"),
  ReminderInterval: signplus.Nullable[float64](float64(4732)),
}

response, err := client.TemplateTemplateIDSetNotification.SetTemplateNotification(context.Background(), "template_id", request, params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

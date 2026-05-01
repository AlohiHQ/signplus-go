# TemplateTemplateIDRename

A list of all methods in the `TemplateTemplateIDRename` service. Click on the method name to view detailed information about that method.

| Methods                           | Description     |
| :-------------------------------- | :-------------- |
| [RenameTemplate](#renametemplate) | Rename template |

## RenameTemplate

Rename template

- HTTP Method: `PUT`
- Endpoint: `/template/{template_id}/rename`

**Parameters**

| Name                  | Type                        | Required | Description                   |
| :-------------------- | :-------------------------- | :------- | :---------------------------- |
| ctx                   | Context                     | ✅       | Default go language context   |
| templateID            | string                      | ✅       |                               |
| renameTemplateRequest | RenameTemplateRequest       | ✅       |                               |
| params                | RenameTemplateRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/templatetemplateidrename"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := templatetemplateidrename.RenameTemplateRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}


request := templatetemplateidrename.RenameTemplateRequest{
  Name: signplus.Nullable[string]("string"),
}

response, err := client.TemplateTemplateIDRename.RenameTemplate(context.Background(), "template_id", request, params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

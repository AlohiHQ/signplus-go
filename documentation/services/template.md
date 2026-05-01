# Template

A list of all methods in the `Template` service. Click on the method name to view detailed information about that method.

| Methods                           | Description         |
| :-------------------------------- | :------------------ |
| [CreateTemplate](#createtemplate) | Create new template |

## CreateTemplate

Create new template

- HTTP Method: `POST`
- Endpoint: `/template`

**Parameters**

| Name                  | Type                        | Required | Description                   |
| :-------------------- | :-------------------------- | :------- | :---------------------------- |
| ctx                   | Context                     | ✅       | Default go language context   |
| createTemplateRequest | CreateTemplateRequest       | ✅       |                               |
| params                | CreateTemplateRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/template"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := template.CreateTemplateRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}


request := template.CreateTemplateRequest{
  Name: signplus.Nullable[string]("rEATXel"),
}

response, err := client.Template.CreateTemplate(context.Background(), request, params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

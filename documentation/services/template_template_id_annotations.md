# TemplateTemplateIDAnnotations

A list of all methods in the `TemplateTemplateIDAnnotations` service. Click on the method name to view detailed information about that method.

| Methods                                           | Description              |
| :------------------------------------------------ | :----------------------- |
| [GetTemplateAnnotations](#gettemplateannotations) | Get template annotations |

## GetTemplateAnnotations

Get template annotations

- HTTP Method: `GET`
- Endpoint: `/template/{template_id}/annotations`

**Parameters**

| Name       | Type                                | Required | Description                   |
| :--------- | :---------------------------------- | :------- | :---------------------------- |
| ctx        | Context                             | ✅       | Default go language context   |
| templateID | string                              | ✅       |                               |
| params     | GetTemplateAnnotationsRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/templatetemplateidannotations"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := templatetemplateidannotations.GetTemplateAnnotationsRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}

response, err := client.TemplateTemplateIDAnnotations.GetTemplateAnnotations(context.Background(), "template_id", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

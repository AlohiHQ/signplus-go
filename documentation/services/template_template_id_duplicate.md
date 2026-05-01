# TemplateTemplateIDDuplicate

A list of all methods in the `TemplateTemplateIDDuplicate` service. Click on the method name to view detailed information about that method.

| Methods                                 | Description        |
| :-------------------------------------- | :----------------- |
| [DuplicateTemplate](#duplicatetemplate) | Duplicate template |

## DuplicateTemplate

Duplicate template

- HTTP Method: `POST`
- Endpoint: `/template/{template_id}/duplicate`

**Parameters**

| Name       | Type                           | Required | Description                   |
| :--------- | :----------------------------- | :------- | :---------------------------- |
| ctx        | Context                        | ✅       | Default go language context   |
| templateID | string                         | ✅       |                               |
| params     | DuplicateTemplateRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/templatetemplateidduplicate"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := templatetemplateidduplicate.DuplicateTemplateRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}

response, err := client.TemplateTemplateIDDuplicate.DuplicateTemplate(context.Background(), "template_id", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

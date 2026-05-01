# TemplateTemplateIDDocuments

A list of all methods in the `TemplateTemplateIDDocuments` service. Click on the method name to view detailed information about that method.

| Methods                                       | Description            |
| :-------------------------------------------- | :--------------------- |
| [GetTemplateDocuments](#gettemplatedocuments) | Get template documents |

## GetTemplateDocuments

Get template documents

- HTTP Method: `GET`
- Endpoint: `/template/{template_id}/documents`

**Parameters**

| Name       | Type                              | Required | Description                   |
| :--------- | :-------------------------------- | :------- | :---------------------------- |
| ctx        | Context                           | ✅       | Default go language context   |
| templateID | string                            | ✅       |                               |
| params     | GetTemplateDocumentsRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/templatetemplateiddocuments"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := templatetemplateiddocuments.GetTemplateDocumentsRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}

response, err := client.TemplateTemplateIDDocuments.GetTemplateDocuments(context.Background(), "template_id", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

# TemplateTemplateIDAnnotationsDocumentID

A list of all methods in the `TemplateTemplateIDAnnotationsDocumentID` service. Click on the method name to view detailed information about that method.

| Methods                                                           | Description                       |
| :---------------------------------------------------------------- | :-------------------------------- |
| [GetDocumentTemplateAnnotations](#getdocumenttemplateannotations) | Get document template annotations |

## GetDocumentTemplateAnnotations

Get document template annotations

- HTTP Method: `GET`
- Endpoint: `/template/{template_id}/annotations/{document_id}`

**Parameters**

| Name       | Type                                        | Required | Description                   |
| :--------- | :------------------------------------------ | :------- | :---------------------------- |
| ctx        | Context                                     | ✅       | Default go language context   |
| templateID | string                                      | ✅       |                               |
| documentID | string                                      | ✅       |                               |
| params     | GetDocumentTemplateAnnotationsRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/templatetemplateidannotationsdocumentid"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := templatetemplateidannotationsdocumentid.GetDocumentTemplateAnnotationsRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}

response, err := client.TemplateTemplateIDAnnotationsDocumentID.GetDocumentTemplateAnnotations(context.Background(), "template_id", "document_id", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

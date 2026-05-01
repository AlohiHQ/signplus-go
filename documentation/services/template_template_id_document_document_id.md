# TemplateTemplateIDDocumentDocumentID

A list of all methods in the `TemplateTemplateIDDocumentDocumentID` service. Click on the method name to view detailed information about that method.

| Methods                                     | Description           |
| :------------------------------------------ | :-------------------- |
| [GetTemplateDocument](#gettemplatedocument) | Get template document |

## GetTemplateDocument

Get template document

- HTTP Method: `GET`
- Endpoint: `/template/{template_id}/document/{document_id}`

**Parameters**

| Name       | Type                             | Required | Description                   |
| :--------- | :------------------------------- | :------- | :---------------------------- |
| ctx        | Context                          | ✅       | Default go language context   |
| templateID | string                           | ✅       |                               |
| documentID | string                           | ✅       |                               |
| params     | GetTemplateDocumentRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/templatetemplateiddocumentdocumentid"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := templatetemplateiddocumentdocumentid.GetTemplateDocumentRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}

response, err := client.TemplateTemplateIDDocumentDocumentID.GetTemplateDocument(context.Background(), "template_id", "document_id", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

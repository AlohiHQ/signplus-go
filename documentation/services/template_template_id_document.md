# TemplateTemplateIDDocument

A list of all methods in the `TemplateTemplateIDDocument` service. Click on the method name to view detailed information about that method.

| Methods                                     | Description           |
| :------------------------------------------ | :-------------------- |
| [AddTemplateDocument](#addtemplatedocument) | Add template document |

## AddTemplateDocument

Add template document

- HTTP Method: `POST`
- Endpoint: `/template/{template_id}/document`

**Parameters**

| Name                       | Type                             | Required | Description                   |
| :------------------------- | :------------------------------- | :------- | :---------------------------- |
| ctx                        | Context                          | ✅       | Default go language context   |
| templateID                 | string                           | ✅       |                               |
| addTemplateDocumentRequest | AddTemplateDocumentRequest       | ✅       |                               |
| params                     | AddTemplateDocumentRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/templatetemplateiddocument"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := templatetemplateiddocument.AddTemplateDocumentRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}


request := templatetemplateiddocument.AddTemplateDocumentRequest{
  File: signplus.Nullable[[]byte]([]byte{}),
}

response, err := client.TemplateTemplateIDDocument.AddTemplateDocument(context.Background(), "template_id", request, params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

# TemplateTemplateIDSetComment

A list of all methods in the `TemplateTemplateIDSetComment` service. Click on the method name to view detailed information about that method.

| Methods                                   | Description          |
| :---------------------------------------- | :------------------- |
| [SetTemplateComment](#settemplatecomment) | Set template comment |

## SetTemplateComment

Set template comment

- HTTP Method: `PUT`
- Endpoint: `/template/{template_id}/set_comment`

**Parameters**

| Name                      | Type                            | Required | Description                   |
| :------------------------ | :------------------------------ | :------- | :---------------------------- |
| ctx                       | Context                         | ✅       | Default go language context   |
| templateID                | string                          | ✅       |                               |
| setTemplateCommentRequest | SetTemplateCommentRequest       | ✅       |                               |
| params                    | SetTemplateCommentRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/templatetemplateidsetcomment"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := templatetemplateidsetcomment.SetTemplateCommentRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}


request := templatetemplateidsetcomment.SetTemplateCommentRequest{
  Comment: signplus.Nullable[string]("string"),
}

response, err := client.TemplateTemplateIDSetComment.SetTemplateComment(context.Background(), "template_id", request, params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

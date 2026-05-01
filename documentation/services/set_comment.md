# SetComment

A list of all methods in the `SetComment` service. Click on the method name to view detailed information about that method.

| Methods                                   | Description          |
| :---------------------------------------- | :------------------- |
| [SetEnvelopeComment](#setenvelopecomment) | Set envelope comment |

## SetEnvelopeComment

Set envelope comment

- HTTP Method: `PUT`
- Endpoint: `/envelope/{envelope_id}/set_comment`

**Parameters**

| Name                      | Type                            | Required | Description                   |
| :------------------------ | :------------------------------ | :------- | :---------------------------- |
| ctx                       | Context                         | ✅       | Default go language context   |
| envelopeID                | string                          | ✅       |                               |
| setEnvelopeCommentRequest | SetEnvelopeCommentRequest       | ✅       |                               |
| params                    | SetEnvelopeCommentRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/setcomment"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := setcomment.SetEnvelopeCommentRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}


request := setcomment.SetEnvelopeCommentRequest{
  Comment: signplus.Nullable[string]("<string>"),
}

response, err := client.SetComment.SetEnvelopeComment(context.Background(), "envelope_id", request, params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

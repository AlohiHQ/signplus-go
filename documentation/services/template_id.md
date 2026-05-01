# TemplateID

A list of all methods in the `TemplateID` service. Click on the method name to view detailed information about that method.

| Methods                                                   | Description                       |
| :-------------------------------------------------------- | :-------------------------------- |
| [CreateEnvelopeFromTemplate](#createenvelopefromtemplate) | Create new envelope from template |

## CreateEnvelopeFromTemplate

Create new envelope from template

- HTTP Method: `POST`
- Endpoint: `/envelope/from_template/{template_id}`

**Parameters**

| Name                              | Type                                    | Required | Description                   |
| :-------------------------------- | :-------------------------------------- | :------- | :---------------------------- |
| ctx                               | Context                                 | ✅       | Default go language context   |
| templateID                        | string                                  | ✅       |                               |
| createEnvelopeFromTemplateRequest | CreateEnvelopeFromTemplateRequest       | ✅       |                               |
| params                            | CreateEnvelopeFromTemplateRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/templateid"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := templateid.CreateEnvelopeFromTemplateRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}


request := templateid.CreateEnvelopeFromTemplateRequest{
  Name: signplus.Nullable[string]("fND"),
  Comment: signplus.Nullable[string]("<string>"),
  Sandbox: signplus.Nullable[bool](true),
}

response, err := client.TemplateID.CreateEnvelopeFromTemplate(context.Background(), "template_id", request, params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

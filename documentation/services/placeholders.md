# Placeholders

A list of all methods in the `Placeholders` service. Click on the method name to view detailed information about that method.

| Methods                                                                   | Description                                                     |
| :------------------------------------------------------------------------ | :-------------------------------------------------------------- |
| [SetEnvelopeAttachmentsPlaceholders](#setenvelopeattachmentsplaceholders) | Placeholders to be set, completely replacing the existing ones. |

## SetEnvelopeAttachmentsPlaceholders

Placeholders to be set, completely replacing the existing ones.

- HTTP Method: `PUT`
- Endpoint: `/envelope/{envelope_id}/attachments/placeholders`

**Parameters**

| Name                                      | Type                                            | Required | Description                   |
| :---------------------------------------- | :---------------------------------------------- | :------- | :---------------------------- |
| ctx                                       | Context                                         | ✅       | Default go language context   |
| envelopeID                                | string                                          | ✅       |                               |
| setEnvelopeAttachmentsPlaceholdersRequest | SetEnvelopeAttachmentsPlaceholdersRequest       | ✅       |                               |
| params                                    | SetEnvelopeAttachmentsPlaceholdersRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/placeholders"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := placeholders.SetEnvelopeAttachmentsPlaceholdersRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}


setEnvelopeAttachmentsPlaceholdersRequestPlaceholders := placeholders.SetEnvelopeAttachmentsPlaceholdersRequestPlaceholders{
  RecipientID: signplus.Nullable[string]("string"),
  Name: signplus.Nullable[string]("string"),
  Required: signplus.Nullable[bool](true),
  Multiple: signplus.Nullable[bool](true),
  ID: signplus.Nullable[string]("string"),
  Hint: signplus.Nullable[string]("string"),
}

request := placeholders.SetEnvelopeAttachmentsPlaceholdersRequest{
  Placeholders: signplus.Nullable[[]placeholders.SetEnvelopeAttachmentsPlaceholdersRequestPlaceholders]([]placeholders.SetEnvelopeAttachmentsPlaceholdersRequestPlaceholders{setEnvelopeAttachmentsPlaceholdersRequestPlaceholders}),
}

response, err := client.Placeholders.SetEnvelopeAttachmentsPlaceholders(context.Background(), "envelope_id", request, params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

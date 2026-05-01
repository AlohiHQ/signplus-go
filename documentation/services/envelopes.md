# Envelopes

A list of all methods in the `Envelopes` service. Click on the method name to view detailed information about that method.

| Methods                         | Description    |
| :------------------------------ | :------------- |
| [ListEnvelopes](#listenvelopes) | List envelopes |

## ListEnvelopes

List envelopes

- HTTP Method: `POST`
- Endpoint: `/envelopes`

**Parameters**

| Name                 | Type                       | Required | Description                   |
| :------------------- | :------------------------- | :------- | :---------------------------- |
| ctx                  | Context                    | ✅       | Default go language context   |
| listEnvelopesRequest | ListEnvelopesRequest       | ✅       |                               |
| params               | ListEnvelopesRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/envelopes"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := envelopes.ListEnvelopesRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}


request := envelopes.ListEnvelopesRequest{
  Name: signplus.Nullable[string]("string"),
  Tags: signplus.Nullable[[]string]([]string{}),
  Comment: signplus.Nullable[string]("string"),
  Ids: signplus.Nullable[[]string]([]string{}),
  Statuses: signplus.Nullable[[]string]([]string{}),
  FolderIds: signplus.Nullable[[]string]([]string{}),
  OnlyRootFolder: signplus.Nullable[bool](true),
  DateFrom: signplus.Nullable[float64](float64(6508)),
  DateTo: signplus.Nullable[float64](float64(690)),
  UID: signplus.Nullable[string]("string"),
  First: signplus.Nullable[float64](float64(6875)),
  Last: signplus.Nullable[float64](float64(4384)),
  After: signplus.Nullable[string]("string"),
  Before: signplus.Nullable[string]("string"),
  OrderField: signplus.Nullable[string]("LAST_DOCUMENT_CHANGE"),
  Ascending: signplus.Nullable[bool](true),
  IncludeTrash: signplus.Nullable[bool](true),
}

response, err := client.Envelopes.ListEnvelopes(context.Background(), request, params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

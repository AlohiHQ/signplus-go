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
  Name: signplus.Nullable[string]("<string>"),
  Tags: signplus.Nullable[[]string]([]string{}),
  Comment: signplus.Nullable[string]("<string>"),
  Ids: signplus.Nullable[[]string]([]string{}),
  Statuses: signplus.Nullable[[]string]([]string{}),
  FolderIds: signplus.Nullable[[]string]([]string{}),
  OnlyRootFolder: signplus.Nullable[string]("<boolean>"),
  DateFrom: signplus.Nullable[string]("<integer>"),
  DateTo: signplus.Nullable[string]("<integer>"),
  UID: signplus.Nullable[string]("<string>"),
  First: signplus.Nullable[string]("<integer>"),
  Last: signplus.Nullable[string]("<integer>"),
  After: signplus.Nullable[string]("<string>"),
  Before: signplus.Nullable[string]("<string>"),
  OrderField: signplus.Nullable[string]("LAST_DOCUMENT_CHANGE"),
  Ascending: signplus.Nullable[string]("<boolean>"),
  IncludeTrash: signplus.Nullable[string]("<boolean>"),
}

response, err := client.Envelopes.ListEnvelopes(context.Background(), request, params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

# FileID

A list of all methods in the `FileID` service. Click on the method name to view detailed information about that method.

| Methods                                 | Description                  |
| :-------------------------------------- | :--------------------------- |
| [GetAttachmentFile](#getattachmentfile) | Get envelope attachment file |

## GetAttachmentFile

Get envelope attachment file

- HTTP Method: `GET`
- Endpoint: `/envelope/{envelope_id}/attachments/{file_id}`

**Parameters**

| Name       | Type                           | Required | Description                   |
| :--------- | :----------------------------- | :------- | :---------------------------- |
| ctx        | Context                        | ✅       | Default go language context   |
| envelopeID | string                         | ✅       |                               |
| fileID     | string                         | ✅       |                               |
| params     | GetAttachmentFileRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/fileid"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := fileid.GetAttachmentFileRequestParams{
  Accept: signplus.Nullable[string]("application/octet-stream"),
}

response, err := client.FileID.GetAttachmentFile(context.Background(), "envelope_id", "file_id", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

# SignedDocuments

A list of all methods in the `SignedDocuments` service. Click on the method name to view detailed information about that method.

| Methods                                                             | Description                               |
| :------------------------------------------------------------------ | :---------------------------------------- |
| [DownloadEnvelopeSignedDocuments](#downloadenvelopesigneddocuments) | Download signed documents for an envelope |

## DownloadEnvelopeSignedDocuments

Download signed documents for an envelope

- HTTP Method: `GET`
- Endpoint: `/envelope/{envelope_id}/signed_documents`

**Parameters**

| Name       | Type                                         | Required | Description                   |
| :--------- | :------------------------------------------- | :------- | :---------------------------- |
| ctx        | Context                                      | ✅       | Default go language context   |
| envelopeID | string                                       | ✅       |                               |
| params     | DownloadEnvelopeSignedDocumentsRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/signeddocuments"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := signeddocuments.DownloadEnvelopeSignedDocumentsRequestParams{
  CertificateOfCompletion: signplus.Nullable[string]("true"),
  Accept: signplus.Nullable[string]("application/pdf"),
}

response, err := client.SignedDocuments.DownloadEnvelopeSignedDocuments(context.Background(), "envelope_id", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

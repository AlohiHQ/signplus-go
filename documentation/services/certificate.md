# Certificate

A list of all methods in the `Certificate` service. Click on the method name to view detailed information about that method.

| Methods                                                     | Description                                        |
| :---------------------------------------------------------- | :------------------------------------------------- |
| [DownloadEnvelopeCertificate](#downloadenvelopecertificate) | Download certificate of completion for an envelope |

## DownloadEnvelopeCertificate

Download certificate of completion for an envelope

- HTTP Method: `GET`
- Endpoint: `/envelope/{envelope_id}/certificate`

**Parameters**

| Name       | Type                                     | Required | Description                   |
| :--------- | :--------------------------------------- | :------- | :---------------------------- |
| ctx        | Context                                  | ✅       | Default go language context   |
| envelopeID | string                                   | ✅       |                               |
| params     | DownloadEnvelopeCertificateRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/certificate"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := certificate.DownloadEnvelopeCertificateRequestParams{
  Accept: signplus.Nullable[string]("application/pdf"),
}

response, err := client.Certificate.DownloadEnvelopeCertificate(context.Background(), "envelope_id", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

# SigningSteps

A list of all methods in the `SigningSteps` service. Click on the method name to view detailed information about that method.

| Methods                                             | Description                |
| :-------------------------------------------------- | :------------------------- |
| [AddEnvelopeSigningSteps](#addenvelopesigningsteps) | Add envelope signing steps |

## AddEnvelopeSigningSteps

Add envelope signing steps

- HTTP Method: `POST`
- Endpoint: `/envelope/{envelope_id}/signing_steps`

**Parameters**

| Name                           | Type                                 | Required | Description                   |
| :----------------------------- | :----------------------------------- | :------- | :---------------------------- |
| ctx                            | Context                              | ✅       | Default go language context   |
| envelopeID                     | string                               | ✅       |                               |
| addEnvelopeSigningStepsRequest | AddEnvelopeSigningStepsRequest       | ✅       |                               |
| params                         | AddEnvelopeSigningStepsRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/signingsteps"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := signingsteps.AddEnvelopeSigningStepsRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}


verification := signingsteps.Verification{
  Type: signplus.Nullable[string]("SMS"),
  Value: signplus.Nullable[string]("string"),
}

signingStepsRecipients1 := signingsteps.SigningStepsRecipients1{
  Name: signplus.Nullable[string]("string"),
  Email: signplus.Nullable[string]("string"),
  Role: signplus.Nullable[string]("IN_PERSON_SIGNER"),
  ID: signplus.Nullable[string]("string"),
  UID: signplus.Nullable[string]("string"),
  Verification: signplus.Nullable[signingsteps.Verification](verification),
}

addEnvelopeSigningStepsRequestSigningSteps := signingsteps.AddEnvelopeSigningStepsRequestSigningSteps{
  Recipients: signplus.Nullable[[]signingsteps.SigningStepsRecipients1]([]signingsteps.SigningStepsRecipients1{signingStepsRecipients1}),
}

request := signingsteps.AddEnvelopeSigningStepsRequest{
  SigningSteps: signplus.Nullable[[]signingsteps.AddEnvelopeSigningStepsRequestSigningSteps]([]signingsteps.AddEnvelopeSigningStepsRequestSigningSteps{addEnvelopeSigningStepsRequestSigningSteps}),
}

response, err := client.SigningSteps.AddEnvelopeSigningSteps(context.Background(), "envelope_id", request, params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

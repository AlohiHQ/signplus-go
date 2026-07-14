package signplus1

// Flow type of the envelope (REQUEST_SIGNATURE is a request for signature, SIGN_MYSELF is a self-signing flow)
type EnvelopeFlowType string

const (
	EnvelopeFlowTypeRequestSignature EnvelopeFlowType = "REQUEST_SIGNATURE"
	EnvelopeFlowTypeSignMyself       EnvelopeFlowType = "SIGN_MYSELF"
)

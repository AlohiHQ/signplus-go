package signplus

// Flow type of the envelope (REQUEST_SIGNATURE is a request for signature, SIGN_MYSELF is a self-signing flow)
type EnvelopeFlowType string

const (
	ENVELOPE_FLOW_TYPE_REQUEST_SIGNATURE EnvelopeFlowType = "REQUEST_SIGNATURE"
	ENVELOPE_FLOW_TYPE_SIGN_MYSELF       EnvelopeFlowType = "SIGN_MYSELF"
)

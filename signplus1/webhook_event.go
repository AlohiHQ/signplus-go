package signplus1

// Event of the webhook
type WebhookEvent string

const (
	WebhookEventEnvelopeExpired    WebhookEvent = "ENVELOPE_EXPIRED"
	WebhookEventEnvelopeDeclined   WebhookEvent = "ENVELOPE_DECLINED"
	WebhookEventEnvelopeVoided     WebhookEvent = "ENVELOPE_VOIDED"
	WebhookEventEnvelopeCompleted  WebhookEvent = "ENVELOPE_COMPLETED"
	WebhookEventEnvelopeAuditTrail WebhookEvent = "ENVELOPE_AUDIT_TRAIL"
)

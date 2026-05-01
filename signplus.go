package signplus

import (
	"github.com/alohihq/signplus-go/annotation"
	"github.com/alohihq/signplus-go/annotationid"
	"github.com/alohihq/signplus-go/annotations"
	"github.com/alohihq/signplus-go/certificate"
	"github.com/alohihq/signplus-go/document"
	"github.com/alohihq/signplus-go/documentid"
	"github.com/alohihq/signplus-go/documents"
	"github.com/alohihq/signplus-go/duplicate"
	"github.com/alohihq/signplus-go/dynamicfields"
	"github.com/alohihq/signplus-go/envelope"
	"github.com/alohihq/signplus-go/envelopeenvelopeidannotationsdocumentid"
	"github.com/alohihq/signplus-go/envelopeid"
	"github.com/alohihq/signplus-go/envelopes"
	"github.com/alohihq/signplus-go/fileid"
	"github.com/alohihq/signplus-go/internal/clients/rest/hooks"
	"github.com/alohihq/signplus-go/internal/configmanager"
	"github.com/alohihq/signplus-go/placeholders"
	"github.com/alohihq/signplus-go/rename"
	"github.com/alohihq/signplus-go/send"
	"github.com/alohihq/signplus-go/setcomment"
	"github.com/alohihq/signplus-go/setexpirationdate"
	"github.com/alohihq/signplus-go/setlegalitylevel"
	"github.com/alohihq/signplus-go/setnotification"
	"github.com/alohihq/signplus-go/settings"
	"github.com/alohihq/signplus-go/signeddocuments"
	"github.com/alohihq/signplus-go/signingsteps"
	"github.com/alohihq/signplus-go/template"
	"github.com/alohihq/signplus-go/templateid"
	"github.com/alohihq/signplus-go/templates"
	"github.com/alohihq/signplus-go/templatetemplateid"
	"github.com/alohihq/signplus-go/templatetemplateidannotation"
	"github.com/alohihq/signplus-go/templatetemplateidannotationannotationid"
	"github.com/alohihq/signplus-go/templatetemplateidannotations"
	"github.com/alohihq/signplus-go/templatetemplateidannotationsdocumentid"
	"github.com/alohihq/signplus-go/templatetemplateidattachmentsplaceholders"
	"github.com/alohihq/signplus-go/templatetemplateidattachmentssettings"
	"github.com/alohihq/signplus-go/templatetemplateiddocument"
	"github.com/alohihq/signplus-go/templatetemplateiddocumentdocumentid"
	"github.com/alohihq/signplus-go/templatetemplateiddocuments"
	"github.com/alohihq/signplus-go/templatetemplateidduplicate"
	"github.com/alohihq/signplus-go/templatetemplateidrename"
	"github.com/alohihq/signplus-go/templatetemplateidsetcomment"
	"github.com/alohihq/signplus-go/templatetemplateidsetnotification"
	"github.com/alohihq/signplus-go/templatetemplateidsigningsteps"
	"github.com/alohihq/signplus-go/void"
	"github.com/alohihq/signplus-go/webhook"
	"github.com/alohihq/signplus-go/webhookid"
	"github.com/alohihq/signplus-go/webhooks"
	"time"
)

// Signplus is the main SDK client that provides access to all service endpoints.
// It manages configuration, authentication, and service instances with centralized settings.
type Signplus struct {
	TemplateID                                *templateid.Service
	SignedDocuments                           *signeddocuments.Service
	Certificate                               *certificate.Service
	DocumentID                                *documentid.Service
	Document                                  *document.Service
	Documents                                 *documents.Service
	DynamicFields                             *dynamicfields.Service
	SigningSteps                              *signingsteps.Service
	Settings                                  *settings.Service
	Placeholders                              *placeholders.Service
	FileID                                    *fileid.Service
	Send                                      *send.Service
	Duplicate                                 *duplicate.Service
	Void                                      *void.Service
	Rename                                    *rename.Service
	SetComment                                *setcomment.Service
	SetNotification                           *setnotification.Service
	SetExpirationDate                         *setexpirationdate.Service
	SetLegalityLevel                          *setlegalitylevel.Service
	EnvelopeEnvelopeIDAnnotationsDocumentID   *envelopeenvelopeidannotationsdocumentid.Service
	Annotations                               *annotations.Service
	AnnotationID                              *annotationid.Service
	Annotation                                *annotation.Service
	EnvelopeID                                *envelopeid.Service
	Envelope                                  *envelope.Service
	Envelopes                                 *envelopes.Service
	TemplateTemplateIDDuplicate               *templatetemplateidduplicate.Service
	TemplateTemplateIDDocumentDocumentID      *templatetemplateiddocumentdocumentid.Service
	TemplateTemplateIDDocument                *templatetemplateiddocument.Service
	TemplateTemplateIDDocuments               *templatetemplateiddocuments.Service
	TemplateTemplateIDSigningSteps            *templatetemplateidsigningsteps.Service
	TemplateTemplateIDRename                  *templatetemplateidrename.Service
	TemplateTemplateIDSetComment              *templatetemplateidsetcomment.Service
	TemplateTemplateIDSetNotification         *templatetemplateidsetnotification.Service
	TemplateTemplateIDAnnotationsDocumentID   *templatetemplateidannotationsdocumentid.Service
	TemplateTemplateIDAnnotations             *templatetemplateidannotations.Service
	TemplateTemplateIDAnnotationAnnotationID  *templatetemplateidannotationannotationid.Service
	TemplateTemplateIDAnnotation              *templatetemplateidannotation.Service
	TemplateTemplateIDAttachmentsSettings     *templatetemplateidattachmentssettings.Service
	TemplateTemplateIDAttachmentsPlaceholders *templatetemplateidattachmentsplaceholders.Service
	TemplateTemplateID                        *templatetemplateid.Service
	Template                                  *template.Service
	Templates                                 *templates.Service
	WebhookID                                 *webhookid.Service
	Webhook                                   *webhook.Service
	Webhooks                                  *webhooks.Service
	manager                                   *configmanager.ConfigManager
}

func NewSignplus(config Config) *Signplus {
	templateID := templateid.NewService()
	signedDocuments := signeddocuments.NewService()
	certificate := certificate.NewService()
	documentID := documentid.NewService()
	document := document.NewService()
	documents := documents.NewService()
	dynamicFields := dynamicfields.NewService()
	signingSteps := signingsteps.NewService()
	settings := settings.NewService()
	placeholders := placeholders.NewService()
	fileID := fileid.NewService()
	send := send.NewService()
	duplicate := duplicate.NewService()
	void := void.NewService()
	rename := rename.NewService()
	setComment := setcomment.NewService()
	setNotification := setnotification.NewService()
	setExpirationDate := setexpirationdate.NewService()
	setLegalityLevel := setlegalitylevel.NewService()
	envelopeEnvelopeIDAnnotationsDocumentID := envelopeenvelopeidannotationsdocumentid.NewService()
	annotations := annotations.NewService()
	annotationID := annotationid.NewService()
	annotation := annotation.NewService()
	envelopeID := envelopeid.NewService()
	envelope := envelope.NewService()
	envelopes := envelopes.NewService()
	templateTemplateIDDuplicate := templatetemplateidduplicate.NewService()
	templateTemplateIDDocumentDocumentID := templatetemplateiddocumentdocumentid.NewService()
	templateTemplateIDDocument := templatetemplateiddocument.NewService()
	templateTemplateIDDocuments := templatetemplateiddocuments.NewService()
	templateTemplateIDSigningSteps := templatetemplateidsigningsteps.NewService()
	templateTemplateIDRename := templatetemplateidrename.NewService()
	templateTemplateIDSetComment := templatetemplateidsetcomment.NewService()
	templateTemplateIDSetNotification := templatetemplateidsetnotification.NewService()
	templateTemplateIDAnnotationsDocumentID := templatetemplateidannotationsdocumentid.NewService()
	templateTemplateIDAnnotations := templatetemplateidannotations.NewService()
	templateTemplateIDAnnotationAnnotationID := templatetemplateidannotationannotationid.NewService()
	templateTemplateIDAnnotation := templatetemplateidannotation.NewService()
	templateTemplateIDAttachmentsSettings := templatetemplateidattachmentssettings.NewService()
	templateTemplateIDAttachmentsPlaceholders := templatetemplateidattachmentsplaceholders.NewService()
	templateTemplateID := templatetemplateid.NewService()
	template := template.NewService()
	templates := templates.NewService()
	webhookID := webhookid.NewService()
	webhook := webhook.NewService()
	webhooks := webhooks.NewService()

	manager := configmanager.NewConfigManager(config)
	hook := hooks.NewDefaultHook()
	templateID.WithConfigManager(manager)
	signedDocuments.WithConfigManager(manager)
	certificate.WithConfigManager(manager)
	documentID.WithConfigManager(manager)
	document.WithConfigManager(manager)
	documents.WithConfigManager(manager)
	dynamicFields.WithConfigManager(manager)
	signingSteps.WithConfigManager(manager)
	settings.WithConfigManager(manager)
	placeholders.WithConfigManager(manager)
	fileID.WithConfigManager(manager)
	send.WithConfigManager(manager)
	duplicate.WithConfigManager(manager)
	void.WithConfigManager(manager)
	rename.WithConfigManager(manager)
	setComment.WithConfigManager(manager)
	setNotification.WithConfigManager(manager)
	setExpirationDate.WithConfigManager(manager)
	setLegalityLevel.WithConfigManager(manager)
	envelopeEnvelopeIDAnnotationsDocumentID.WithConfigManager(manager)
	annotations.WithConfigManager(manager)
	annotationID.WithConfigManager(manager)
	annotation.WithConfigManager(manager)
	envelopeID.WithConfigManager(manager)
	envelope.WithConfigManager(manager)
	envelopes.WithConfigManager(manager)
	templateTemplateIDDuplicate.WithConfigManager(manager)
	templateTemplateIDDocumentDocumentID.WithConfigManager(manager)
	templateTemplateIDDocument.WithConfigManager(manager)
	templateTemplateIDDocuments.WithConfigManager(manager)
	templateTemplateIDSigningSteps.WithConfigManager(manager)
	templateTemplateIDRename.WithConfigManager(manager)
	templateTemplateIDSetComment.WithConfigManager(manager)
	templateTemplateIDSetNotification.WithConfigManager(manager)
	templateTemplateIDAnnotationsDocumentID.WithConfigManager(manager)
	templateTemplateIDAnnotations.WithConfigManager(manager)
	templateTemplateIDAnnotationAnnotationID.WithConfigManager(manager)
	templateTemplateIDAnnotation.WithConfigManager(manager)
	templateTemplateIDAttachmentsSettings.WithConfigManager(manager)
	templateTemplateIDAttachmentsPlaceholders.WithConfigManager(manager)
	templateTemplateID.WithConfigManager(manager)
	template.WithConfigManager(manager)
	templates.WithConfigManager(manager)
	webhookID.WithConfigManager(manager)
	webhook.WithConfigManager(manager)
	webhooks.WithConfigManager(manager)
	templateID.WithHook(hook)
	signedDocuments.WithHook(hook)
	certificate.WithHook(hook)
	documentID.WithHook(hook)
	document.WithHook(hook)
	documents.WithHook(hook)
	dynamicFields.WithHook(hook)
	signingSteps.WithHook(hook)
	settings.WithHook(hook)
	placeholders.WithHook(hook)
	fileID.WithHook(hook)
	send.WithHook(hook)
	duplicate.WithHook(hook)
	void.WithHook(hook)
	rename.WithHook(hook)
	setComment.WithHook(hook)
	setNotification.WithHook(hook)
	setExpirationDate.WithHook(hook)
	setLegalityLevel.WithHook(hook)
	envelopeEnvelopeIDAnnotationsDocumentID.WithHook(hook)
	annotations.WithHook(hook)
	annotationID.WithHook(hook)
	annotation.WithHook(hook)
	envelopeID.WithHook(hook)
	envelope.WithHook(hook)
	envelopes.WithHook(hook)
	templateTemplateIDDuplicate.WithHook(hook)
	templateTemplateIDDocumentDocumentID.WithHook(hook)
	templateTemplateIDDocument.WithHook(hook)
	templateTemplateIDDocuments.WithHook(hook)
	templateTemplateIDSigningSteps.WithHook(hook)
	templateTemplateIDRename.WithHook(hook)
	templateTemplateIDSetComment.WithHook(hook)
	templateTemplateIDSetNotification.WithHook(hook)
	templateTemplateIDAnnotationsDocumentID.WithHook(hook)
	templateTemplateIDAnnotations.WithHook(hook)
	templateTemplateIDAnnotationAnnotationID.WithHook(hook)
	templateTemplateIDAnnotation.WithHook(hook)
	templateTemplateIDAttachmentsSettings.WithHook(hook)
	templateTemplateIDAttachmentsPlaceholders.WithHook(hook)
	templateTemplateID.WithHook(hook)
	template.WithHook(hook)
	templates.WithHook(hook)
	webhookID.WithHook(hook)
	webhook.WithHook(hook)
	webhooks.WithHook(hook)

	return &Signplus{
		TemplateID:                               templateID,
		SignedDocuments:                          signedDocuments,
		Certificate:                              certificate,
		DocumentID:                               documentID,
		Document:                                 document,
		Documents:                                documents,
		DynamicFields:                            dynamicFields,
		SigningSteps:                             signingSteps,
		Settings:                                 settings,
		Placeholders:                             placeholders,
		FileID:                                   fileID,
		Send:                                     send,
		Duplicate:                                duplicate,
		Void:                                     void,
		Rename:                                   rename,
		SetComment:                               setComment,
		SetNotification:                          setNotification,
		SetExpirationDate:                        setExpirationDate,
		SetLegalityLevel:                         setLegalityLevel,
		EnvelopeEnvelopeIDAnnotationsDocumentID:  envelopeEnvelopeIDAnnotationsDocumentID,
		Annotations:                              annotations,
		AnnotationID:                             annotationID,
		Annotation:                               annotation,
		EnvelopeID:                               envelopeID,
		Envelope:                                 envelope,
		Envelopes:                                envelopes,
		TemplateTemplateIDDuplicate:              templateTemplateIDDuplicate,
		TemplateTemplateIDDocumentDocumentID:     templateTemplateIDDocumentDocumentID,
		TemplateTemplateIDDocument:               templateTemplateIDDocument,
		TemplateTemplateIDDocuments:              templateTemplateIDDocuments,
		TemplateTemplateIDSigningSteps:           templateTemplateIDSigningSteps,
		TemplateTemplateIDRename:                 templateTemplateIDRename,
		TemplateTemplateIDSetComment:             templateTemplateIDSetComment,
		TemplateTemplateIDSetNotification:        templateTemplateIDSetNotification,
		TemplateTemplateIDAnnotationsDocumentID:  templateTemplateIDAnnotationsDocumentID,
		TemplateTemplateIDAnnotations:            templateTemplateIDAnnotations,
		TemplateTemplateIDAnnotationAnnotationID: templateTemplateIDAnnotationAnnotationID,
		TemplateTemplateIDAnnotation:             templateTemplateIDAnnotation,
		TemplateTemplateIDAttachmentsSettings:    templateTemplateIDAttachmentsSettings,
		TemplateTemplateIDAttachmentsPlaceholders: templateTemplateIDAttachmentsPlaceholders,
		TemplateTemplateID:                        templateTemplateID,
		Template:                                  template,
		Templates:                                 templates,
		WebhookID:                                 webhookID,
		Webhook:                                   webhook,
		Webhooks:                                  webhooks,
		manager:                                   manager,
	}
}

func (s *Signplus) SetBaseURL(baseURL string) {
	s.manager.SetBaseURL(baseURL)
}

func (s *Signplus) SetTimeout(timeout time.Duration) {
	s.manager.SetTimeout(timeout)
}

func (s *Signplus) SetAccessToken(accessToken string) {
	s.manager.SetAccessToken(accessToken)
}

// SetEnvironment configures the SDK to use the specified environment's base URL.
func (s *Signplus) SetEnvironment(environment Environment) {
	s.manager.SetBaseURL(string(environment))
}

// c029837e0e474b76bc487506e8799df5e3335891efe4fb02bda7a1441840310c

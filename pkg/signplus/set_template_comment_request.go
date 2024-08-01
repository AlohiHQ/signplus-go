package signplus

type SetTemplateCommentRequest struct {
	// Comment for the template
	Comment *string `json:"comment,omitempty"`
}

func (s *SetTemplateCommentRequest) SetComment(comment string) {
	s.Comment = &comment
}

func (s *SetTemplateCommentRequest) GetComment() *string {
	if s == nil {
		return nil
	}
	return s.Comment
}

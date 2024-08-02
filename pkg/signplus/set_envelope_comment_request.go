package signplus

type SetEnvelopeCommentRequest struct {
	// Comment for the envelope
	Comment *string `json:"comment,omitempty" required:"true"`
}

func (s *SetEnvelopeCommentRequest) SetComment(comment string) {
	s.Comment = &comment
}

func (s *SetEnvelopeCommentRequest) GetComment() *string {
	if s == nil {
		return nil
	}
	return s.Comment
}

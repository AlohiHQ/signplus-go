package signplus

import "encoding/json"

type SetEnvelopeCommentRequest struct {
	// Comment for the envelope
	Comment *string `json:"comment,omitempty" required:"true"`
}

func (s *SetEnvelopeCommentRequest) GetComment() *string {
	if s == nil {
		return nil
	}
	return s.Comment
}

func (s *SetEnvelopeCommentRequest) SetComment(comment string) {
	s.Comment = &comment
}

func (s SetEnvelopeCommentRequest) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SetEnvelopeCommentRequest to string"
	}
	return string(jsonData)
}

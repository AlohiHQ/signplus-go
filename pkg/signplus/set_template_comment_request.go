package signplus

import "encoding/json"

type SetTemplateCommentRequest struct {
	// Comment for the template
	Comment *string `json:"comment,omitempty" required:"true"`
}

func (s *SetTemplateCommentRequest) GetComment() *string {
	if s == nil {
		return nil
	}
	return s.Comment
}

func (s *SetTemplateCommentRequest) SetComment(comment string) {
	s.Comment = &comment
}

func (s SetTemplateCommentRequest) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SetTemplateCommentRequest to string"
	}
	return string(jsonData)
}

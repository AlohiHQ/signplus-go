package signplus

import (
	"encoding/json"
)

type SetTemplateCommentRequest struct {
	// Comment for the template
	Comment *string `json:"comment,omitempty" required:"true"`
	touched map[string]bool
}

func (s *SetTemplateCommentRequest) GetComment() *string {
	if s == nil {
		return nil
	}
	return s.Comment
}

func (s *SetTemplateCommentRequest) SetComment(comment string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Comment"] = true
	s.Comment = &comment
}

func (s *SetTemplateCommentRequest) SetCommentNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Comment"] = true
	s.Comment = nil
}
func (s SetTemplateCommentRequest) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if s.touched["Comment"] && s.Comment == nil {
		data["comment"] = nil
	} else if s.Comment != nil {
		data["comment"] = s.Comment
	}

	return json.Marshal(data)
}

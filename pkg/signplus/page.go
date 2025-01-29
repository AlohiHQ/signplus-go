package signplus

import (
	"encoding/json"
)

type Page struct {
	// Width of the page in pixels
	Width *int64 `json:"width,omitempty"`
	// Height of the page in pixels
	Height  *int64 `json:"height,omitempty"`
	touched map[string]bool
}

func (p *Page) GetWidth() *int64 {
	if p == nil {
		return nil
	}
	return p.Width
}

func (p *Page) SetWidth(width int64) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Width"] = true
	p.Width = &width
}

func (p *Page) SetWidthNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Width"] = true
	p.Width = nil
}

func (p *Page) GetHeight() *int64 {
	if p == nil {
		return nil
	}
	return p.Height
}

func (p *Page) SetHeight(height int64) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Height"] = true
	p.Height = &height
}

func (p *Page) SetHeightNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Height"] = true
	p.Height = nil
}

func (p Page) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if p.touched["Width"] && p.Width == nil {
		data["width"] = nil
	} else if p.Width != nil {
		data["width"] = p.Width
	}

	if p.touched["Height"] && p.Height == nil {
		data["height"] = nil
	} else if p.Height != nil {
		data["height"] = p.Height
	}

	return json.Marshal(data)
}

func (p Page) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: Page to string"
	}
	return string(jsonData)
}

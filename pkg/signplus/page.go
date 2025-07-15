package signplus

import "encoding/json"

type Page struct {
	// Width of the page in pixels
	Width *int64 `json:"width,omitempty"`
	// Height of the page in pixels
	Height *int64 `json:"height,omitempty"`
}

func (p *Page) GetWidth() *int64 {
	if p == nil {
		return nil
	}
	return p.Width
}

func (p *Page) SetWidth(width int64) {
	p.Width = &width
}

func (p *Page) GetHeight() *int64 {
	if p == nil {
		return nil
	}
	return p.Height
}

func (p *Page) SetHeight(height int64) {
	p.Height = &height
}

func (p Page) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: Page to string"
	}
	return string(jsonData)
}

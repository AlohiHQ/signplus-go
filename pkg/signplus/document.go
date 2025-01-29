package signplus

import (
	"encoding/json"
)

type Document struct {
	// Unique identifier of the document
	Id *string `json:"id,omitempty"`
	// Name of the document
	Name *string `json:"name,omitempty"`
	// Filename of the document
	Filename *string `json:"filename,omitempty"`
	// Number of pages in the document
	PageCount *int64 `json:"page_count,omitempty"`
	// List of pages in the document
	Pages   []Page `json:"pages,omitempty"`
	touched map[string]bool
}

func (d *Document) GetId() *string {
	if d == nil {
		return nil
	}
	return d.Id
}

func (d *Document) SetId(id string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Id"] = true
	d.Id = &id
}

func (d *Document) SetIdNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Id"] = true
	d.Id = nil
}

func (d *Document) GetName() *string {
	if d == nil {
		return nil
	}
	return d.Name
}

func (d *Document) SetName(name string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Name"] = true
	d.Name = &name
}

func (d *Document) SetNameNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Name"] = true
	d.Name = nil
}

func (d *Document) GetFilename() *string {
	if d == nil {
		return nil
	}
	return d.Filename
}

func (d *Document) SetFilename(filename string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Filename"] = true
	d.Filename = &filename
}

func (d *Document) SetFilenameNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Filename"] = true
	d.Filename = nil
}

func (d *Document) GetPageCount() *int64 {
	if d == nil {
		return nil
	}
	return d.PageCount
}

func (d *Document) SetPageCount(pageCount int64) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["PageCount"] = true
	d.PageCount = &pageCount
}

func (d *Document) SetPageCountNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["PageCount"] = true
	d.PageCount = nil
}

func (d *Document) GetPages() []Page {
	if d == nil {
		return nil
	}
	return d.Pages
}

func (d *Document) SetPages(pages []Page) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Pages"] = true
	d.Pages = pages
}

func (d *Document) SetPagesNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Pages"] = true
	d.Pages = nil
}

func (d Document) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if d.touched["Id"] && d.Id == nil {
		data["id"] = nil
	} else if d.Id != nil {
		data["id"] = d.Id
	}

	if d.touched["Name"] && d.Name == nil {
		data["name"] = nil
	} else if d.Name != nil {
		data["name"] = d.Name
	}

	if d.touched["Filename"] && d.Filename == nil {
		data["filename"] = nil
	} else if d.Filename != nil {
		data["filename"] = d.Filename
	}

	if d.touched["PageCount"] && d.PageCount == nil {
		data["page_count"] = nil
	} else if d.PageCount != nil {
		data["page_count"] = d.PageCount
	}

	if d.touched["Pages"] && d.Pages == nil {
		data["pages"] = nil
	} else if d.Pages != nil {
		data["pages"] = d.Pages
	}

	return json.Marshal(data)
}

func (d Document) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: Document to string"
	}
	return string(jsonData)
}

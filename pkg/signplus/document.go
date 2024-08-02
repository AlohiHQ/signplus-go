package signplus

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
	Pages []Page `json:"pages,omitempty"`
}

func (d *Document) SetId(id string) {
	d.Id = &id
}

func (d *Document) GetId() *string {
	if d == nil {
		return nil
	}
	return d.Id
}

func (d *Document) SetName(name string) {
	d.Name = &name
}

func (d *Document) GetName() *string {
	if d == nil {
		return nil
	}
	return d.Name
}

func (d *Document) SetFilename(filename string) {
	d.Filename = &filename
}

func (d *Document) GetFilename() *string {
	if d == nil {
		return nil
	}
	return d.Filename
}

func (d *Document) SetPageCount(pageCount int64) {
	d.PageCount = &pageCount
}

func (d *Document) GetPageCount() *int64 {
	if d == nil {
		return nil
	}
	return d.PageCount
}

func (d *Document) SetPages(pages []Page) {
	d.Pages = pages
}

func (d *Document) GetPages() []Page {
	if d == nil {
		return nil
	}
	return d.Pages
}

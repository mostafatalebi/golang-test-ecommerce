package models

type Category struct {
	Title string `json:"title"`
	Slug  string `json:"slug"`
}

func (c *Category) Clone() *Category {
	return &Category{
		Title: c.Title,
		Slug:  c.Slug,
	}
}

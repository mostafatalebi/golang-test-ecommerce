package models

import "sync"

type ProductModel struct {
	ID          uint64  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Rating      *Rate   `json:"rating"`

	CreatedAt int64 `json:"createdAt"`
	UpdatedAt int64 `json:"updatedAt"`
}

func NewProductModel() *ProductModel {
	pp := &ProductModel{}
	pp.Rating = &Rate{}

	return pp
}

// reference to the products objects' are bound to
// change at any moment during runtime (it depends on
// implementer too, though) but
// for consistency, make sure you use Clone()
// since this allows you to get a lasting object
// for your scoped usage.
func (p *ProductModel) Clone() *ProductModel {
	return &ProductModel{
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		Price:       p.Price,
		Image:       p.Image,
		Category:    p.Category,
		Rating:      p.Rating.Clone(),
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

// such a project doesn't need object reusing
// but I am putting this here for demonstation
// purposes. Normally for a production level
// project, pooling must absolutely be accompanied by
// reflection-based test cased to avoid accidental
// fields missing (as this is hugely critical)
var productPool *sync.Pool

func init() {
	productPool = &sync.Pool{
		New: func() any {
			return NewProductModel()
		},
	}
}

func AcquireProductModel() *ProductModel {
	return productPool.Get().(*ProductModel)
}

// this function MUST have reflection based
// testing for fields matching to ensure
// all product model's fields are put in this
// function
func ReleaseProductModel(p *ProductModel) {
	p.ID = 0
	p.Title = ""
	p.Description = ""
	p.Image = ""
	p.Category = ""
	p.Rating = nil
	p.CreatedAt = 0
	p.UpdatedAt = 0
	productPool.Put(p)
}

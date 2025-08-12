package models

import (
	"errors"
	"net/url"
)

var ErrProductTitle       = errors.New("product title is required and must be between 2 to 128 characters")
var ErrProductDescription = errors.New("product title is required and must be between 32 to 2048 characters")
var ErrProductImage       = errors.New("product image is required and must be a valid URL")
var ErrProductCategory    = errors.New("product category is required")
var ErrProductPrice       = errors.New("product price is required")

func ValidateProduct(p *ProductModel) error {
	if length := len(p.Title); length < 2 || length > 32 {
		return ErrProductTitle
	}

	if length := len(p.Description); length < 2 || length > 32 {
		return ErrProductDescription
	}

	if p.Image == "" {
		return ErrProductImage
	} else if _, err := url.Parse(p.Image); err != nil {
		return ErrProductImage
	}

	if p.Category == "" {
		return ErrProductCategory
	}

	if p.Price == 0.0 {
		return ErrProductPrice
	}

	return nil
}

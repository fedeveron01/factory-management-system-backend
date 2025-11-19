package gateway_entities

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name             string
	Description      string
	Color            string
	SizeType         string
	Size             float64
	ImageUrl         string
	Price            float64
	Stock            float64
	MaterialProduct  []MaterialProduct
	ProductVariation []ProductVariation
}

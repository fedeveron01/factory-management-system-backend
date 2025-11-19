package entities

import (
	"github.com/fedeveron01/golang-base/cmd/core"
	"github.com/fedeveron01/golang-base/cmd/core/enums"
)

type Product struct {
	core.EntitiesBase
	Name             string
	Description      string
	Color            string
	SizeType         enums.Enum
	Size             float64
	ImageUrl         string
	Price            float64
	Stock            float64
	MaterialProduct  []MaterialProduct
	ProductVariation []ProductVariation
}

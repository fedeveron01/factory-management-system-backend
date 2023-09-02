package entities

import "github.com/fedeveron01/golang-base/cmd/core"

type Product struct {
	core.EntitiesBase
	Name            string
	Description     string
	Price           float64
	Stock           int
	MaterialProduct []MaterialProduct
}
package gateways

import (
	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
	"github.com/fedeveron01/golang-base/cmd/core"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	_ "github.com/fedeveron01/golang-base/cmd/core/entities"
	"gorm.io/gorm"
)

type ProductVariationGateway interface {
	FindById(id uint) *entities.ProductVariation
}

type ProductVariationRepository interface {
	CreateProductVariation(productVariation gateway_entities.ProductVariation) (gateway_entities.ProductVariation, error)
	FindById(id uint) *gateway_entities.ProductVariation
	FindByProductIDAndNumber(productID uint, number float64) *gateway_entities.ProductVariation
}

type ProductVariationGatewayImpl struct {
	productVariationRepository ProductVariationRepository
}

func NewProductVariationGateway(productVariationRepository ProductVariationRepository) *ProductVariationGatewayImpl {
	return &ProductVariationGatewayImpl{
		productVariationRepository: productVariationRepository,
	}
}

func (i *ProductVariationGatewayImpl) Create(productVariation entities.ProductVariation, productID uint) (entities.ProductVariation, error) {
	productVariationDB := i.ToServiceEntity(productVariation, productID)
	productVariationDB, err := i.productVariationRepository.CreateProductVariation(productVariationDB)
	if err != nil {
		return entities.ProductVariation{}, err
	}
	productVariation = i.ToBusinessEntity(productVariationDB)
	return productVariation, nil
}

func (i *ProductVariationGatewayImpl) FindById(id uint) *entities.ProductVariation {
	productVariationDB := i.productVariationRepository.FindById(id)
	if productVariationDB == nil {
		return nil
	}
	productVariation := i.ToBusinessEntity(*productVariationDB)
	return &productVariation
}

func (i *ProductVariationGatewayImpl) FindByProductIDAndNumber(productID uint, number float64) *entities.ProductVariation {
	productVariationDB := i.productVariationRepository.FindByProductIDAndNumber(productID, number)
	if productVariationDB == nil {
		return nil
	}
	productVariation := i.ToBusinessEntity(*productVariationDB)
	return &productVariation
}

func (i *ProductVariationGatewayImpl) ToBusinessEntity(productVariationDB gateway_entities.ProductVariation) entities.ProductVariation {
	return entities.ProductVariation{
		EntitiesBase: core.EntitiesBase{
			ID: productVariationDB.ID,
		},
		Number: productVariationDB.Number,
		Stock:  productVariationDB.Stock,
		Product: entities.Product{
			EntitiesBase: core.EntitiesBase{
				ID: productVariationDB.ProductID,
			},
		},
	}
}

func (i *ProductVariationGatewayImpl) ToServiceEntity(productVariation entities.ProductVariation, productID uint) gateway_entities.ProductVariation {
	return gateway_entities.ProductVariation{
		Model: gorm.Model{
			ID: productVariation.ID,
		},
		Number:    productVariation.Number,
		Stock:     productVariation.Stock,
		ProductID: productID,
	}
}
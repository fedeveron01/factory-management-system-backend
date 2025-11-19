package gateways

import (
	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
	"github.com/fedeveron01/golang-base/cmd/core"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	"github.com/fedeveron01/golang-base/cmd/core/enums"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll() ([]gateway_entities.Product, error)
	FindAllWithVariations() ([]gateway_entities.Product, error)
	FindById(id uint) *gateway_entities.Product
	FindByName(name string) []*gateway_entities.Product
	FindByNameAndColor(name string, color string) *gateway_entities.Product
	CreateProduct(product gateway_entities.Product) (gateway_entities.Product, error)
	UpdateProduct(product gateway_entities.Product) (gateway_entities.Product, error)
	DeleteProduct(id uint) error
	UpdateMaterialProducts(productId int64, materialsProduct []gateway_entities.MaterialProduct) ([]gateway_entities.MaterialProduct, error)
}

type ProductGatewayImpl struct {
	productRepository ProductRepository
}

func NewProductGateway(productRepository ProductRepository) *ProductGatewayImpl {
	return &ProductGatewayImpl{
		productRepository: productRepository,
	}
}

func (e *ProductGatewayImpl) FindAll() ([]entities.Product, error) {
	productsDB, err := e.productRepository.FindAll()
	if err != nil {
		return nil, err
	}
	products := make([]entities.Product, len(productsDB))
	for i, productDB := range productsDB {
		products[i] = e.ToBusinessEntity(productDB)
	}
	return products, err
}

func (e *ProductGatewayImpl) FindAllWithVariations() ([]entities.Product, error) {
	productsDB, err := e.productRepository.FindAllWithVariations()
	if err != nil {
		return nil, err
	}
	products := make([]entities.Product, len(productsDB))
	for i, productDB := range productsDB {
		products[i] = e.ToBusinessEntity(productDB)
	}
	return products, err
}

func (e *ProductGatewayImpl) FindById(id uint) *entities.Product {
	productDB := e.productRepository.FindById(id)
	if productDB == nil {
		return nil
	}
	product := e.ToBusinessEntity(*productDB)
	product.MaterialProduct = e.MaterialsProductToBusinessEntity(productDB.MaterialProduct)
	return &product
}

func (e *ProductGatewayImpl) FindByName(name string) []entities.Product {
	productsDB := e.productRepository.FindByName(name)
	if len(productsDB) == 0 {
		return nil
	}
	products := make([]entities.Product, len(productsDB))
	for i, productDB := range productsDB {
		products[i] = e.ToBusinessEntity(*productDB)
	}
	return products
}

func (e *ProductGatewayImpl) FindByNameAndColor(name string, color string) *entities.Product {
	productDB := e.productRepository.FindByNameAndColor(name, color)
	if productDB == nil {
		return nil
	}
	product := e.ToBusinessEntity(*productDB)
	return &product
}

func (e *ProductGatewayImpl) CreateProduct(product entities.Product) (entities.Product, error) {
	productDB := e.ToServiceEntity(product)
	created, err := e.productRepository.CreateProduct(productDB)
	if err != nil {
		return entities.Product{}, err
	}
	product = e.ToBusinessEntity(created)
	return product, nil
}

func (e *ProductGatewayImpl) UpdateProduct(product entities.Product) (entities.Product, error) {

	productDB := e.ToServiceEntity(product)
	var err error
	productDB, err = e.productRepository.UpdateProduct(productDB)
	if err != nil {
		return entities.Product{}, err
	}
	product = e.ToBusinessEntity(productDB)
	return product, nil

}

func (e *ProductGatewayImpl) DeleteProduct(id uint) error {
	return e.productRepository.DeleteProduct(id)
}

func (e *ProductGatewayImpl) UpdateMaterialProducts(productId uint, materialsProduct []entities.MaterialProduct) ([]entities.MaterialProduct, error) {
	materialsProductDB := make([]gateway_entities.MaterialProduct, len(materialsProduct))
	for i, materialProduct := range materialsProduct {
		materialsProductDB[i] = gateway_entities.MaterialProduct{
			ProductId:  productId,
			MaterialId: materialProduct.Material.ID,
			Quantity:   materialProduct.Quantity,
		}
	}
	materialsProductDB, err := e.productRepository.UpdateMaterialProducts(int64(productId), materialsProductDB)
	if err != nil {
		return nil, err
	}
	materialsProduct = e.MaterialsProductToBusinessEntity(materialsProductDB)
	return materialsProduct, nil
}

func (e *ProductGatewayImpl) ToBusinessEntity(product gateway_entities.Product) entities.Product {
	productBusiness := entities.Product{
		EntitiesBase: core.EntitiesBase{
			ID: product.ID,
		},
		Name:             product.Name,
		Description:      product.Description,
		Color:            product.Color,
		SizeType:         enums.Enum(product.SizeType),
		Size:             product.Size,
		ImageUrl:         product.ImageUrl,
		Price:            product.Price,
		ProductVariation: ToProductVariationBusinessEntity(product.ProductVariation),
		Stock:            getStock(product.ProductVariation),
	}
	return productBusiness
}

func getStock(productVariations []gateway_entities.ProductVariation) float64 {
	stock := 0.0
	for _, productVariation := range productVariations {
		stock += productVariation.Stock
	}
	return stock
}

func ToProductVariationBusinessEntity(productVariations []gateway_entities.ProductVariation) []entities.ProductVariation {
	productVariationsBusiness := make([]entities.ProductVariation, len(productVariations))
	for i, productVariation := range productVariations {
		productVariationsBusiness[i] = entities.ProductVariation{
			EntitiesBase: core.EntitiesBase{
				ID: productVariation.ID,
			},
			Number: productVariation.Number,
			Stock:  productVariation.Stock,
		}
	}
	return productVariationsBusiness

}

func (e *ProductGatewayImpl) ToServiceEntity(product entities.Product) gateway_entities.Product {
	productDB := gateway_entities.Product{
		Model: gorm.Model{
			ID: product.ID,
		},
		Name:        product.Name,
		Description: product.Description,
		Color:       product.Color,
		SizeType:    string(product.SizeType),
		Size:        product.Size,
		ImageUrl:    product.ImageUrl,
		Price:       product.Price,
		Stock:       product.Stock,
	}
	return productDB
}

func (e *ProductGatewayImpl) MaterialsProductToBusinessEntity(materialProducts []gateway_entities.MaterialProduct) []entities.MaterialProduct {
	materialProductsBusiness := make([]entities.MaterialProduct, len(materialProducts))
	for i, materialProduct := range materialProducts {
		materialProductsBusiness[i] = entities.MaterialProduct{
			EntitiesBase: core.EntitiesBase{
				ID: materialProduct.ID,
			},
			Material: entities.Material{
				EntitiesBase: core.EntitiesBase{
					ID: materialProduct.Material.ID,
				},
				Name: materialProduct.Material.Name,
				MaterialType: entities.MaterialType{
					EntitiesBase: core.EntitiesBase{
						ID: materialProduct.Material.MaterialType.ID,
					},
					Name:              materialProduct.Material.MaterialType.Name,
					UnitOfMeasurement: enums.StringToUnitOfMeasurementEnum(materialProduct.Material.MaterialType.UnitOfMeasurement),
				},
			},

			Quantity: materialProduct.Quantity,
		}
	}
	return materialProductsBusiness
}

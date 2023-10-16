package repositories

import (
	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
	core_errors "github.com/fedeveron01/golang-base/cmd/core/errors"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(database *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: database,
	}
}

func (r *ProductRepository) CreateProduct(product gateway_entities.Product) (gateway_entities.Product, error) {
	id := r.db.Create(&product)
	if id.Error != nil {
		return gateway_entities.Product{}, id.Error
	}
	return product, nil
}

func (r *ProductRepository) FindAll() ([]gateway_entities.Product, error) {
	var products []gateway_entities.Product
	r.db.Find(&products)
	return products, nil
}

func (r *ProductRepository) FindByName(name string) *gateway_entities.Product {
	var product gateway_entities.Product
	r.db.Where("name = ?", name).First(&product)
	if product.ID == 0 {
		return nil
	}
	return &product
}

func (r *ProductRepository) FindById(id uint) *gateway_entities.Product {
	var product gateway_entities.Product
	res := r.db.Find(&product, id).First(&product)
	if res.Error != nil {
		return nil
	}
	if product.ID == 0 {
		return nil
	}
	var materialProducts []gateway_entities.MaterialProduct
	res = r.db.InnerJoins("Material").Preload("Material.MaterialType").Find(&materialProducts, "product_id = ?", id)
	if res.Error != nil {
		return &product
	}
	product.MaterialProduct = materialProducts

	return &product
}

func (r *ProductRepository) UpdateProduct(product gateway_entities.Product) (gateway_entities.Product, error) {
	res := r.db.Save(&product)
	if res.Error != nil {
		return gateway_entities.Product{}, res.Error
	}
	if res.RowsAffected == 0 {
		return gateway_entities.Product{}, core_errors.NewInternalServerError("product update failed")
	}
	return product, nil
}

func (r *ProductRepository) DeleteProduct(id uint) error {
	r.db.Delete(&gateway_entities.Product{}, id)
	return nil
}
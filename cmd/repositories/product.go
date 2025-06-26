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

func (r *ProductRepository) FindAllWithVariations() ([]gateway_entities.Product, error) {
	var products []gateway_entities.Product
	r.db.Preload("ProductVariation").Find(&products)
	return products, nil
}

func (r *ProductRepository) FindByName(name string) []*gateway_entities.Product {
	var products []gateway_entities.Product
	r.db.Where("name = ?", name).Find(&products)
	if len(products) == 0 {
		return nil
	}
	var result []*gateway_entities.Product
	for _, product := range products {
		var productVariations []gateway_entities.ProductVariation
		r.db.Find(&productVariations, "product_id = ?", product.ID)
		p := product
		p.ProductVariation = productVariations
		result = append(result, &p)
	}
	return result
}

func (r *ProductRepository) FindByNameAndColor(name string, color string) *gateway_entities.Product {
	var product gateway_entities.Product
	r.db.Where("name = ? AND color = ?", name, color).First(&product)
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
	var productVariations []gateway_entities.ProductVariation
	res = r.db.Find(&productVariations, "product_id = ?", id)
	product.ProductVariation = productVariations
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

func (r *ProductRepository) UpdateMaterialProducts(productId int64, materialProduct []gateway_entities.MaterialProduct) ([]gateway_entities.MaterialProduct, error) {
	r.db.Unscoped().Where("product_id = ?", productId).Delete(&gateway_entities.MaterialProduct{})
	if len(materialProduct) == 0 {
		return materialProduct, nil
	}
	res := r.db.Create(&materialProduct)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, core_errors.NewInternalServerError("material product update failed")
	}

	r.db.InnerJoins("Material").Preload("Material.MaterialType").Find(&materialProduct, "product_id = ?", productId)

	return materialProduct, nil
}

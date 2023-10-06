package repositories

import (
	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
	core_errors "github.com/fedeveron01/golang-base/cmd/core/errors"
	"gorm.io/gorm"
)

type MaterialRepository struct {
	db *gorm.DB
}

func NewMaterialRepository(database *gorm.DB) *MaterialRepository {
	return &MaterialRepository{
		db: database,
	}
}

func (r *MaterialRepository) CreateMaterial(material gateway_entities.Material) error {
	id := r.db.Create(&material)
	if id.Error != nil {
		return id.Error
	}
	return nil
}

func (r *MaterialRepository) FindAll() ([]gateway_entities.Material, error) {
	var materials []gateway_entities.Material
	r.db.InnerJoins("MaterialType").Find(&materials)
	return materials, nil
}

func (r *MaterialRepository) UpdateMaterial(material gateway_entities.Material) error {
	r.db.Save(&material)
	return nil
}

func (r *MaterialRepository) DeleteMaterial(id string) error {
	result := r.db.Delete(&gateway_entities.Material{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return core_errors.NewInternalServerError("Material Id not found")
	}
	return nil
}

package repositories

import (
	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
	core_errors "github.com/fedeveron01/golang-base/cmd/core/errors"
	"gorm.io/gorm"
)

type MovementRepository struct {
	db *gorm.DB
}

func NewMovementRepository(database *gorm.DB) *MovementRepository {
	return &MovementRepository{
		db: database,
	}
}

func (r *MovementRepository) CreateMovement(movement gateway_entities.Movement) (gateway_entities.Movement, error) {
	id := r.db.Create(&movement)
	if id.Error != nil {
		return gateway_entities.Movement{}, id.Error
	}
	return movement, nil
}

func (r *MovementRepository) FindAll() ([]gateway_entities.Movement, error) {
	var movements []gateway_entities.Movement
	result := r.db.Find(&movements)
	if result.Error != nil {
		return nil, result.Error
	}
	return movements, nil
}

func (r *MovementRepository) FindAllByType(typeValue string) ([]gateway_entities.Movement, error) {
	var movements []gateway_entities.Movement
	result := r.db.Find(&movements, "type = ?", typeValue)
	if result.Error != nil {
		return nil, result.Error
	}
	return movements, nil
}

func (r *MovementRepository) FindById(id uint) (movement gateway_entities.Movement, err error) {
	res := r.db.Find(&movement, id).First(&movement)
	if res.Error != nil {
		return movement, res.Error
	}
	if movement.ID == 0 {
		return movement, core_errors.NewNotFoundError("movement not found")
	}
	var movementDetails []gateway_entities.MovementDetail
	res = r.db.Preload("Material.MaterialType").Preload("ProductVariation.Product").Find(&movementDetails, "movement_id = ?", id)
	if res.Error != nil {
		return movement, res.Error
	}
	movement.MovementDetail = movementDetails

	return movement, nil
}

package gateways

import (
	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
	"github.com/fedeveron01/golang-base/cmd/core"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	_ "github.com/fedeveron01/golang-base/cmd/core/entities"
	"github.com/fedeveron01/golang-base/cmd/core/enums"
)

type MovementRepository interface {
	CreateMovement(movement gateway_entities.Movement) (gateway_entities.Movement, error)
	FindAll() ([]gateway_entities.Movement, error)
	FindAllByType(typeValue string) ([]gateway_entities.Movement, error)
	FindById(id uint) (gateway_entities.Movement, error)
}

type MovementGateway interface {
	Create(movement entities.Movement) (entities.Movement, error)
	FindAll() ([]entities.Movement, error)
	FindAllByType(typeValue string) ([]entities.Movement, error)
	FindById(id uint) (entities.Movement, error)
}

type MovementGatewayImpl struct {
	movementRepository MovementRepository
}

func NewMovementGateway(movementRepository MovementRepository) *MovementGatewayImpl {
	return &MovementGatewayImpl{
		movementRepository: movementRepository,
	}
}

func (i *MovementGatewayImpl) FindAll() ([]entities.Movement, error) {
	movementsDB, err := i.movementRepository.FindAll()
	if err != nil {
		return nil, err
	}
	movements := make([]entities.Movement, len(movementsDB))
	for index, movementDB := range movementsDB {
		movements[index] = i.ToBusinessEntity(movementDB)
	}
	return movements, nil
}

func (i *MovementGatewayImpl) FindAllByType(typeValue string) ([]entities.Movement, error) {
	movementsDB, err := i.movementRepository.FindAllByType(typeValue)
	if err != nil {
		return nil, err
	}
	movements := make([]entities.Movement, len(movementsDB))
	for index, movementDB := range movementsDB {
		movements[index] = i.ToBusinessEntity(movementDB)
	}
	return movements, nil
}

func (i *MovementGatewayImpl) FindById(id uint) (entities.Movement, error) {
	movementDB, err := i.movementRepository.FindById(id)
	if err != nil {
		return entities.Movement{}, err
	}
	movement := i.ToBusinessEntity(movementDB)
	return movement, nil
}

func (i *MovementGatewayImpl) Create(movement entities.Movement, employeeID uint) (entities.Movement, error) {
	movementDB := i.ToServiceEntity(movement, movement.ID)
	movementDB.EmployeeId = employeeID
	movementDB.MovementDetail = nil
	movementDBCreated, err := i.movementRepository.CreateMovement(movementDB)
	if err != nil {
		return entities.Movement{}, err
	}
	movementCreated := i.ToBusinessEntity(movementDBCreated)
	return movementCreated, nil
}

func (i *MovementGatewayImpl) ToServiceEntity(movement entities.Movement, movementID uint) gateway_entities.Movement {
	return gateway_entities.Movement{
		Number:         float64(movement.Number),
		Type:           movement.Type,
		Total:          movement.Total,
		DateTime:       movement.DateTime,
		Description:    movement.Description,
		MovementDetail: toServiceMovementDetail(movement.MovementDetail),
	}
}

func toServiceMovementDetail(movementDetail []entities.MovementDetail) []gateway_entities.MovementDetail {
	var movementDetailDB []gateway_entities.MovementDetail
	for _, v := range movementDetail {
		movementDetailDB = append(movementDetailDB, gateway_entities.MovementDetail{
			Quantity:           v.Quantity,
			Price:              v.Price,
			MaterialId:         &v.Material.ID,
			ProductVariationId: &v.ProductVariation.ID,
		})
	}
	return movementDetailDB
}

func (i *MovementGatewayImpl) ToBusinessEntity(movement gateway_entities.Movement) entities.Movement {
	return entities.Movement{
		EntitiesBase: core.EntitiesBase{
			ID: movement.ID,
		},
		Number:             int(movement.Number),
		Type:               movement.Type,
		Total:              movement.Total,
		DateTime:           movement.DateTime,
		IsMaterialMovement: movement.IsMaterialMovement,
		Description:        movement.Description,
		MovementDetail:     i.toBusinessMovementDetails(movement.MovementDetail),
	}
}

func (i *MovementGatewayImpl) toServiceMovementDetails(movementDetails []entities.MovementDetail, movementID uint) []gateway_entities.MovementDetail {
	movementDetailsDB := make([]gateway_entities.MovementDetail, len(movementDetails))
	for index, movementDetail := range movementDetails {
		movementDetailsDB[index] = i.toServiceMovementDetail(movementDetail, movementID)
	}
	return movementDetailsDB
}

func (i *MovementGatewayImpl) toServiceMovementDetail(movementDetail entities.MovementDetail, movementID uint) gateway_entities.MovementDetail {
	var materialID uint
	if movementDetail.Material != nil {
		materialID = movementDetail.Material.ID
	}
	var productVariationID uint
	if movementDetail.ProductVariation != nil {
		productVariationID = movementDetail.ProductVariation.ID
	}
	return gateway_entities.MovementDetail{
		MaterialId:         &materialID,
		ProductVariationId: &productVariationID,
		Quantity:           movementDetail.Quantity,
		MovementId:         movementID,
	}
}

func (i *MovementGatewayImpl) toBusinessMovementDetails(movementDetails []gateway_entities.MovementDetail) []entities.MovementDetail {
	movementDetailsBusiness := make([]entities.MovementDetail, len(movementDetails))
	for index, movementDetail := range movementDetails {
		movementDetailsBusiness[index] = i.toBusinessMovementDetail(movementDetail)
	}
	return movementDetailsBusiness
}

func (i *MovementGatewayImpl) toBusinessMovementDetail(movementDetail gateway_entities.MovementDetail) entities.MovementDetail {
	var material *entities.Material
	var productVariation *entities.ProductVariation
	if movementDetail.Material != nil && movementDetail.Material.ID != 0 {
		material = &entities.Material{
			EntitiesBase: core.EntitiesBase{
				ID: movementDetail.Material.ID,
			},
			Name:            movementDetail.Material.Name,
			Description:     movementDetail.Material.Description,
			Price:           movementDetail.Material.Price,
			Stock:           movementDetail.Material.Stock,
			RepositionPoint: movementDetail.Material.RepositionPoint,
			MaterialType: entities.MaterialType{
				EntitiesBase: core.EntitiesBase{
					ID: movementDetail.Material.MaterialType.ID,
				},
				Name:              movementDetail.Material.MaterialType.Name,
				Description:       movementDetail.Material.MaterialType.Description,
				UnitOfMeasurement: enums.StringToUnitOfMeasurementEnum(movementDetail.Material.MaterialType.UnitOfMeasurement),
			},
		}
	}
	if movementDetail.ProductVariation != nil && movementDetail.ProductVariation.ID != 0 {
		productVariation = &entities.ProductVariation{
			EntitiesBase: core.EntitiesBase{
				ID: movementDetail.ProductVariation.ID,
			},
			Number: movementDetail.ProductVariation.Number,
			Stock:  movementDetail.ProductVariation.Stock,
			Product: entities.Product{
				EntitiesBase: core.EntitiesBase{
					ID: movementDetail.ProductVariation.Product.ID,
				},
				Name:             movementDetail.ProductVariation.Product.Name,
				Description:      movementDetail.ProductVariation.Product.Description,
				Color:            movementDetail.ProductVariation.Product.Color,
				Size:             movementDetail.ProductVariation.Product.Size,
				ImageUrl:         movementDetail.ProductVariation.Product.ImageUrl,
				Price:            movementDetail.ProductVariation.Product.Price,
				Stock:            movementDetail.ProductVariation.Product.Stock,
			},
		}
	}

	return entities.MovementDetail{
		EntitiesBase: core.EntitiesBase{
			ID: movementDetail.ID,
		},
		Material:         material,
		ProductVariation: productVariation,
		Quantity:         movementDetail.Quantity,
		Price:            movementDetail.Price,
	}
}

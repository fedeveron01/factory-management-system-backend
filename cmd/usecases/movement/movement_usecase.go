package movement_usecase

import (
	"errors"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	"strconv"
)

type MovementUseCase interface {
	Create(movement entities.Movement, employeeId uint) (entities.Movement, error)
	FindAllByType(typeValue string) ([]entities.Movement, error)
	FindAll() ([]entities.Movement, error)
	FindById(id uint) (entities.Movement, error)
}

type MovementGateway interface {
	Create(movement entities.Movement, employeeID uint) (entities.Movement, error)
	FindAllByType(typeValue string) ([]entities.Movement, error)
	FindAll() ([]entities.Movement, error)
	FindById(id uint) (entities.Movement, error)
}

type ProductVariationGateway interface {
	Create(productVariation entities.ProductVariation, productID uint) (entities.ProductVariation, error)
	FindByProductIDAndNumber(productID uint, number float64) *entities.ProductVariation
}

type MovementDetailGateway interface {
	CreateMovementDetailsTransaction(movementDetails []entities.MovementDetail, movement entities.Movement, employeeID uint) ([]entities.MovementDetail, entities.Movement, error)
}
type MaterialGateway interface {
	FindMaterialById(id uint) *entities.Material
}

type MovementUseCaseImpl struct {
	movementGateway         MovementGateway
	materialGateway         MaterialGateway
	movementDetailGateway   MovementDetailGateway
	productVariationGateway ProductVariationGateway
}

func NewMovementUseCase(movementGateway MovementGateway, movementDetailGateway MovementDetailGateway, materialGateway MaterialGateway, productVariationGateway ProductVariationGateway) *MovementUseCaseImpl {
	return &MovementUseCaseImpl{
		movementGateway:         movementGateway,
		materialGateway:         materialGateway,
		movementDetailGateway:   movementDetailGateway,
		productVariationGateway: productVariationGateway,
	}
}

func (i *MovementUseCaseImpl) FindAllByType(typeValue string) ([]entities.Movement, error) {
	movements, err := i.movementGateway.FindAllByType(typeValue)
	if err != nil {
		return nil, err
	}
	return movements, nil
}

func (i *MovementUseCaseImpl) FindAll() ([]entities.Movement, error) {
	movements, err := i.movementGateway.FindAll()
	if err != nil {
		return nil, err
	}
	return movements, nil
}

func (i *MovementUseCaseImpl) FindById(id uint) (entities.Movement, error) {
	movement, err := i.movementGateway.FindById(id)
	if err != nil {
		return movement, errors.New("movement not found")
	}
	return movement, nil
}

func (i *MovementUseCaseImpl) updateMaterial(movementDetail *entities.MovementDetail, input bool) error {
	material := i.materialGateway.FindMaterialById(movementDetail.Material.ID)
	if material == nil {
		return errors.New("material not found")
	}

	if input {
		material.Stock += movementDetail.Quantity
	} else {
		if material.Stock-movementDetail.Quantity < 0 {
			return errors.New("insufficient stock in material " + material.Name)
		}
		material.Stock -= movementDetail.Quantity
	}

	movementDetail.Material = material

	return nil
}

func (i *MovementUseCaseImpl) updateProductVariation(movementDetail *entities.MovementDetail, input bool) error {
	productVariation := i.productVariationGateway.FindByProductIDAndNumber(movementDetail.ProductVariation.Product.ID, movementDetail.ProductVariation.Number)
	if productVariation == nil {
		productVariationCreated, err := i.productVariationGateway.Create(*movementDetail.ProductVariation, movementDetail.ProductVariation.Product.ID)
		productVariation = &productVariationCreated
		if err != nil {
			return err
		}
	}
	if input {
		productVariation.Stock += movementDetail.Quantity
	} else {
		if productVariation.Stock-movementDetail.Quantity < 0 {
			return errors.New("insufficient stock in product variation " + strconv.Itoa(int(productVariation.ID)))
		}
		productVariation.Stock -= movementDetail.Quantity
	}
	movementDetail.ProductVariation = productVariation

	return nil
}

func (i *MovementUseCaseImpl) Create(movement entities.Movement, employeeID uint) (entities.Movement, error) {
	if movement.Type == "" {
		return entities.Movement{}, errors.New("type is required")
	}
	if employeeID <= 0 {
		return entities.Movement{}, errors.New("employee is required")
	}
	isInput := movement.Type == "input"

	for index, movementDetail := range movement.MovementDetail {
		if movementDetail.Material.ID == 0 && movementDetail.ProductVariation.Product.ID == 0 {
			return entities.Movement{}, errors.New("material or product variation required")
		}
		if movementDetail.Material.ID != 0 {
			if !movement.IsMaterialMovement {
				return entities.Movement{}, errors.New("movement is not material movement")
			}

			err := i.updateMaterial(&movementDetail, isInput)
			if err != nil {
				return entities.Movement{}, err
			}
			movement.MovementDetail[index] = movementDetail

		}
		if movementDetail.ProductVariation.Product.ID != 0 {
			if movement.IsMaterialMovement {
				return entities.Movement{}, errors.New("movement is not product movement")
			}
			err := i.updateProductVariation(&movementDetail, isInput)
			if err != nil {
				return entities.Movement{}, err
			}
			movement.MovementDetail[index] = movementDetail
		}

	}
	movementDetails := movement.MovementDetail

	movementDetailsCreated, movementCreated, err := i.movementDetailGateway.CreateMovementDetailsTransaction(movementDetails, movement, employeeID)
	if err != nil {
		return entities.Movement{}, err
	}
	movementCreated.MovementDetail = movementDetailsCreated
	return movementCreated, nil

}

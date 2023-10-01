package material_type_usecase

import (
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	core_errors "github.com/fedeveron01/golang-base/cmd/core/errors"
)

type MaterialTypeUseCase interface {
	FindAll() ([]entities.MaterialType, error)
	CreateMaterialType(materialType entities.MaterialType) (entities.MaterialType, error)
}

type MaterialTypeTypeGateway interface {
	FindAll() ([]entities.MaterialType, error)
	FindByName(name string) *entities.MaterialType
	CreateMaterialType(materialTypeType entities.MaterialType) (entities.MaterialType, error)
}

type Implementation struct {
	materialTypeGateway MaterialTypeTypeGateway
}

func NewMaterialTypeUsecase(materialTypeTypeGateway MaterialTypeTypeGateway) *Implementation {
	return &Implementation{
		materialTypeGateway: materialTypeTypeGateway,
	}
}

func (i *Implementation) FindAll() ([]entities.MaterialType, error) {
	materials, err := i.materialTypeGateway.FindAll()
	if err != nil {
		return nil, err
	}
	return materials, nil
}

func (i *Implementation) CreateMaterialType(materialType entities.MaterialType) (entities.MaterialType, error) {
	if materialType.Name == "" {
		return entities.MaterialType{}, core_errors.NewBadRequestError("materialType name is required")
	}
	if len(materialType.Name) < 2 {
		return entities.MaterialType{}, core_errors.NewBadRequestError("materialType name must be at least 2 characters")
	}
	if len(materialType.Name) > 20 {
		return entities.MaterialType{}, core_errors.NewBadRequestError("materialType name must be at most 20 characters")
	}
	repeated := i.materialTypeGateway.FindByName(materialType.Name)
	if repeated != nil {
		return entities.MaterialType{}, core_errors.NewInternalServerError("materialType already exists")
	}
	var err error
	materialType, err = i.materialTypeGateway.CreateMaterialType(materialType)
	if err != nil {
		return entities.MaterialType{}, err
	}
	return materialType, nil
}

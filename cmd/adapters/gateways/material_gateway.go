package gateways

import (
	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
	"github.com/fedeveron01/golang-base/cmd/core"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	_ "github.com/fedeveron01/golang-base/cmd/core/entities"
	"github.com/fedeveron01/golang-base/cmd/repositories"
)

type MaterialGateway interface {
	CreateMaterial(material entities.Material) error
	FindAll() ([]entities.Material, error)
	UpdateMaterial(material entities.Material) error
	DeleteMaterial(id string) error
}

type MaterialGatewayImpl struct {
	materialRepository repositories.MaterialRepository
}

func NewMaterialGateway(materialRepository repositories.MaterialRepository) *MaterialGatewayImpl {
	return &MaterialGatewayImpl{
		materialRepository: materialRepository,
	}
}

func (i *MaterialGatewayImpl) CreateMaterial(material entities.Material) (entities.Material, error) {

	materialDB := i.ToServiceEntity(material)
	materialDBCreated, err := i.materialRepository.CreateMaterial(materialDB)
	if err != nil {
		return entities.Material{}, err
	}
	materialCreated := i.ToBusinessEntity(materialDBCreated)

	return materialCreated, nil
}

func (i *MaterialGatewayImpl) FindAll() ([]entities.Material, error) {
	materialsDB, err := i.materialRepository.FindAll()
	if err != nil {
		return nil, err
	}
	materials := make([]entities.Material, len(materialsDB))
	for index, materialDB := range materialsDB {
		materials[index] = i.ToBusinessEntity(materialDB)

	}
	return materials, err
}

func (i *MaterialGatewayImpl) FindByName(name string) *entities.Material {
	materialDB := i.materialRepository.FindByName(name)
	if materialDB == nil {
		return nil
	}
	material := i.ToBusinessEntity(*materialDB)
	return &material
}

func (i *MaterialGatewayImpl) UpdateMaterial(material entities.Material) error {
	materialDB := i.ToServiceEntity(material)
	return i.materialRepository.UpdateMaterial(materialDB)
}

func (i *MaterialGatewayImpl) DeleteMaterial(id string) error {
	return i.materialRepository.DeleteMaterial(id)
}

func (i *MaterialGatewayImpl) ToBusinessEntity(materialDB gateway_entities.Material) entities.Material {

	material := entities.Material{
		EntitiesBase: core.EntitiesBase{
			ID: materialDB.ID,
		},
		Name:            materialDB.Name,
		Description:     materialDB.Description,
		Price:           materialDB.Price,
		RepositionPoint: materialDB.RepositionPoint,
		Stock:           materialDB.Stock,
		MaterialType: entities.MaterialType{
			EntitiesBase: core.EntitiesBase{
				ID: materialDB.MaterialType.ID,
			},
			Name: materialDB.MaterialType.Name},
	}
	return material
}

func (i *MaterialGatewayImpl) ToServiceEntity(material entities.Material) gateway_entities.Material {
	materialDB := gateway_entities.Material{
		Name:            material.Name,
		Description:     material.Description,
		Price:           material.Price,
		RepositionPoint: material.RepositionPoint,
		Stock:           material.Stock,
		MaterialTypeId:  material.MaterialType.ID,
	}
	return materialDB
}
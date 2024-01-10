package gateways

import (
	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
	"github.com/fedeveron01/golang-base/cmd/core"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	_ "github.com/fedeveron01/golang-base/cmd/core/entities"
	"github.com/fedeveron01/golang-base/cmd/core/enums"
	"gorm.io/gorm"
)

type MaterialGateway interface {
	CreateMaterial(material entities.Material) error
	FindAll() ([]entities.Material, error)
	FindById(id uint) *entities.Material
	UpdateMaterial(material entities.Material) (entities.Material, error)
	DeleteMaterial(id string) error
}

type MaterialRepository interface {
	CreateMaterial(material gateway_entities.Material) (gateway_entities.Material, error)
	FindAll() ([]gateway_entities.Material, error)
	FindById(id uint) *gateway_entities.Material
	FindMaterialById(id uint) *gateway_entities.Material
	FindByName(name string) *gateway_entities.Material
	UpdateMaterial(material gateway_entities.Material) (gateway_entities.Material, error)
	DeleteMaterial(id string) error
}

type MaterialGatewayImpl struct {
	materialRepository MaterialRepository
}

func NewMaterialGateway(materialRepository MaterialRepository) *MaterialGatewayImpl {
	return &MaterialGatewayImpl{
		materialRepository: materialRepository,
	}
}
func (i *MaterialGatewayImpl) FindMaterialById(id uint) *entities.Material {
	materialDB := i.materialRepository.FindMaterialById(id)
	if materialDB == nil {
		return nil
	}
	material := i.ToBusinessEntity(*materialDB)
	return &material
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

func (i *MaterialGatewayImpl) FindById(id uint) *entities.Material {
	materialDB := i.materialRepository.FindById(id)
	if materialDB == nil {
		return nil
	}
	material := i.ToBusinessEntity(*materialDB)
	return &material
}

func (i *MaterialGatewayImpl) FindByName(name string) *entities.Material {
	materialDB := i.materialRepository.FindByName(name)
	if materialDB == nil {
		return nil
	}
	material := i.ToBusinessEntity(*materialDB)
	return &material
}

func (i *MaterialGatewayImpl) UpdateMaterial(material entities.Material) (entities.Material, error) {
	materialDB := i.ToServiceEntity(material)
	var err error
	materialDB, err = i.materialRepository.UpdateMaterial(materialDB)
	if err != nil {
		return entities.Material{}, err
	}
	material = i.ToBusinessEntity(materialDB)
	return material, nil
}

func (i *MaterialGatewayImpl) DeleteMaterial(id string) error {
	return i.materialRepository.DeleteMaterial(id)
}

func (i *MaterialGatewayImpl) ToBusinessEntity(materialDB gateway_entities.Material) entities.Material {

	material := entities.Material{
		EntitiesBase: core.EntitiesBase{
			ID: materialDB.ID,
		},
		Name:        materialDB.Name,
		Number:      materialDB.Number,
		Description: materialDB.Description,
		Price:       materialDB.Price,
		Stock:       materialDB.Stock,
		MaterialType: entities.MaterialType{
			EntitiesBase: core.EntitiesBase{
				ID: materialDB.MaterialType.ID,
			},
			Name:              materialDB.MaterialType.Name,
			UnitOfMeasurement: enums.StringToUnitOfMeasurementEnum(materialDB.MaterialType.UnitOfMeasurement),
		},
		RepositionPoint: materialDB.RepositionPoint,
	}
	return material
}

func (i *MaterialGatewayImpl) ToServiceEntity(material entities.Material) gateway_entities.Material {
	materialDB := gateway_entities.Material{
		Model: gorm.Model{
			ID: material.ID,
		},
		Name:            material.Name,
		Number:          material.Number,
		Description:     material.Description,
		Price:           material.Price,
		Stock:           material.Stock,
		MaterialTypeId:  material.MaterialType.ID,
		RepositionPoint: material.RepositionPoint,
	}
	return materialDB
}

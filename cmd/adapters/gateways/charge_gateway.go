package gateways

import (
	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
	"github.com/fedeveron01/golang-base/cmd/core"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	"gorm.io/gorm"
)

type ChargeRepository interface {
	FindAll() ([]gateway_entities.Charge, error)
	FindByName(name string) (gateway_entities.Charge, error)
	FindById(id uint) (gateway_entities.Charge, error)
	CreateCharge(charge gateway_entities.Charge) (gateway_entities.Charge, error)
}

type ChargeGatewayImpl struct {
	chargeRepository ChargeRepository
}

func NewChargeGateway(chargeRepository ChargeRepository) *ChargeGatewayImpl {
	return &ChargeGatewayImpl{
		chargeRepository: chargeRepository,
	}
}
func (c ChargeGatewayImpl) FindAll() ([]entities.Charge, error) {
	chargesDB, err := c.chargeRepository.FindAll()
	if err != nil {
		return nil, err
	}
	charges := make([]entities.Charge, len(chargesDB))
	for i, chargesDB := range chargesDB {
		charges[i] = c.ToBusinessEntity(chargesDB)
	}
	return charges, err
}

func (c ChargeGatewayImpl) FindByName(name string) (uint, error) {
	charge, err := c.chargeRepository.FindByName(name)
	if err != nil {
		return 0, err
	}
	return charge.ID, nil
}

func (c ChargeGatewayImpl) FindById(id uint) (entities.Charge, error) {
	charge, err := c.chargeRepository.FindById(id)
	if err != nil {
		return entities.Charge{}, err
	}
	return entities.Charge{
		EntitiesBase: core.EntitiesBase{
			ID: charge.ID,
		},
		Name: charge.Name,
	}, nil
}

func (c ChargeGatewayImpl) CreateCharge(charge entities.Charge) (entities.Charge, error) {
	chargeDB := c.ToServiceEntity(charge)
	created, err := c.chargeRepository.CreateCharge(chargeDB)
	if err != nil {
		return entities.Charge{}, err
	}
	charge = c.ToBusinessEntity(created)
	return charge, nil
}

func (c ChargeGatewayImpl) ToBusinessEntity(chargeDB gateway_entities.Charge) entities.Charge {
	return entities.Charge{
		EntitiesBase: core.EntitiesBase{
			ID: chargeDB.ID,
		},
		Name: chargeDB.Name,
	}
}

func (c ChargeGatewayImpl) ToServiceEntity(charge entities.Charge) gateway_entities.Charge {
	return gateway_entities.Charge{
		Model: gorm.Model{
			ID: charge.ID,
		},
		Name: charge.Name,
	}
}

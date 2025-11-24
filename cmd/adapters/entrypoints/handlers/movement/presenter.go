package movement

import (
	"time"

	"github.com/fedeveron01/golang-base/cmd/core"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	"github.com/fedeveron01/golang-base/cmd/core/enums"
	//. "github.com/fedeveron01/golang-base/cmd/adapters/entrypoints/handlers/material"
)

func ToMovement(movementRequest MovementRequest) entities.Movement {
	return entities.Movement{
		Type:               movementRequest.Type,
		DateTime:           time.Now(),
		Description:        movementRequest.Description,
		IsMaterialMovement: movementRequest.IsMaterialMovement,
		MovementDetail:     ToMovementDetails(movementRequest.Details),
	}
}

func ToMovementDetails(movementDetailsRequest []MovementDetailRequest) []entities.MovementDetail {
	var movementDetails []entities.MovementDetail
	for _, movementDetailRequest := range movementDetailsRequest {
		movementDetails = append(movementDetails, ToMovementDetail(movementDetailRequest))
	}
	return movementDetails
}

func ToMovementDetail(movementDetailRequest MovementDetailRequest) entities.MovementDetail {
	return entities.MovementDetail{
		Quantity: movementDetailRequest.Quantity,
		Price:    movementDetailRequest.Price,
		Material: &entities.Material{
			EntitiesBase: core.EntitiesBase{
				ID: movementDetailRequest.MaterialID,
			},
		},
		ProductVariation: &entities.ProductVariation{
			Number: movementDetailRequest.Number,
			Product: entities.Product{
				EntitiesBase: core.EntitiesBase{
					ID: movementDetailRequest.ProductID,
				},
			},
		},
	}
}
func ToMovementResponse(movement entities.Movement) MovementResponse {
	return MovementResponse{
		ID:                 movement.ID,
		Number:             movement.Number,
		Type:               movement.Type,
		Total:              movement.Total,
		IsMaterialMovement: movement.IsMaterialMovement,
		DateTime:           movement.DateTime.Format("2006-01-02 15:04:05"),
		Description:        movement.Description,
		MovementDetails:    ToMovementDetailsResponse(movement.MovementDetail),
	}
}

func ToMovementDetailsResponse(movementDetails []entities.MovementDetail) []MovementDetailResponse {
	var movementDetailsResponse []MovementDetailResponse
	for _, movementDetail := range movementDetails {
		movementDetailsResponse = append(movementDetailsResponse, ToMovementDetailResponse(movementDetail))
	}
	return movementDetailsResponse
}

func ToMovementDetailResponse(movementDetail entities.MovementDetail) MovementDetailResponse {
	var materialID uint
	if movementDetail.Material != nil && movementDetail.Material.ID != 0 {
		materialID = movementDetail.Material.ID
		return MovementDetailResponse{
			ID:         movementDetail.ID,
			MaterialID: &materialID,
			Quantity:   movementDetail.Quantity,
			Price:      movementDetail.Price,
			Material:   ToMaterialResponse(*movementDetail.Material),
		}
	}
	if movementDetail.ProductVariation != nil && movementDetail.ProductVariation.ID != 0 {
		return MovementDetailResponse{
			ID:       movementDetail.ID,
			Quantity: movementDetail.Quantity,
			Price:    movementDetail.Price,
			ProductVariation: &ProductVariationResponse{
				ID:       movementDetail.ProductVariation.ID,
				Number:   movementDetail.ProductVariation.Number,
				Quantity: movementDetail.Quantity,
				Product: &ProductInfoResponse{
					ID:    movementDetail.ProductVariation.Product.ID,
					Name:  movementDetail.ProductVariation.Product.Name,
					Color: movementDetail.ProductVariation.Product.Color,
				},
			},
		}
	}
	return MovementDetailResponse{}
}

func ToMovementsResponse(movements []entities.Movement) []MovementResponse {
	var movementsResponse []MovementResponse
	for _, movement := range movements {
		movementsResponse = append(movementsResponse, ToMovementResponse(movement))
	}
	return movementsResponse
}

func ToMaterialResponse(material entities.Material) *MaterialResponse {
	return &MaterialResponse{
		ID:                      material.ID,
		Name:                    material.Name,
		Description:             material.Description,
		MaterialType:            material.MaterialType.Name,
		Stock:                   material.Stock,
		UnitOfMeasurement:       enums.EnumToUnitOfMeasurementStringInSpanish(material.MaterialType.UnitOfMeasurement),
		UnitOfMeasurementSymbol: enums.GetSymbolByUnitOfMeasurementEnum(material.MaterialType.UnitOfMeasurement),
	}
}

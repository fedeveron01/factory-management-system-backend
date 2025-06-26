package product_handler

import (
	"github.com/fedeveron01/golang-base/cmd/core"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	"github.com/fedeveron01/golang-base/cmd/core/enums"
)

func ToProductsResponse(products []entities.Product) []ProductResponse {
	var productResponses []ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}
	return productResponses
}

func ToProductResponse(product entities.Product) ProductResponse {
	return ProductResponse{
		Id:          float64(product.ID),
		Name:        product.Name,
		Description: product.Description,
		Color:       product.Color,
		ImageUrl:    product.ImageUrl,
		Price:       product.Price,
	}
}

func ToProductWithStockResponse(product entities.Product) ProductWithStockResponse {
	return ProductWithStockResponse{
		Id:          float64(product.ID),
		Name:        product.Name,
		Description: product.Description,
		Color:       product.Color,
		ImageUrl:    product.ImageUrl,
		Price:       product.Price,
		Stock:       product.Stock,
	}
}

func ToProductsByNameResponse(products [][]entities.Product) [][]ProductResponse {
	var productResponses [][]ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ToProductsResponse(product))
	}
	return productResponses
}

func ToProductsByNameMapResponse(products map[string][]entities.Product) map[string][]ProductResponse {
	productResponses := make(map[string][]ProductResponse)
	for key, value := range products {
		productResponses[key] = ToProductsResponse(value)
	}
	return productResponses
}

func ToProductsByNameMapWithStockResponse(products map[string][]entities.Product) map[string][]ProductWithStockResponse {
	productResponses := make(map[string][]ProductWithStockResponse)
	for key, value := range products {
		var productWithStockResponses []ProductWithStockResponse
		for _, product := range value {
			productWithStockResponses = append(productWithStockResponses, ToProductWithStockResponse(product))
		}
		productResponses[key] = productWithStockResponses
	}
	return productResponses
}

func ToProductsWhitVariationsResponse(products []entities.Product) []ProductWithVariationsResponse {
	var productResponses []ProductWithVariationsResponse
	for _, product := range products {
		productResponses = append(productResponses, ToProductWithVariationsResponse(product))
	}
	return productResponses
}

func ToProductWithVariationsResponse(product entities.Product) ProductWithVariationsResponse {
	return ProductWithVariationsResponse{
		Id:               float64(product.ID),
		Name:             product.Name,
		Description:      product.Description,
		Color:            product.Color,
		ImageUrl:         product.ImageUrl,
		Price:            product.Price,
		ProductVariation: ToProductVariationResponse(product.ProductVariation),
		Stock:            product.Stock,
	}
}

func ToProductVariationResponse(productVariation []entities.ProductVariation) []ProductVariationResponse {
	var productVariationResponses []ProductVariationResponse
	for _, productVariation := range productVariation {
		productVariationResponses = append(productVariationResponses, toProductVariationResponse(productVariation))
	}
	return productVariationResponses
}

func toProductVariationResponse(productVariation entities.ProductVariation) ProductVariationResponse {
	return ProductVariationResponse{
		Id:     float64(productVariation.ID),
		Number: productVariation.Number,
		Stock:  productVariation.Stock,
	}
}

func ToProductWithAssignationsResponse(product entities.Product, language string) ProductWithAssignationsResponse {
	return ProductWithAssignationsResponse{
		Id:               float64(product.ID),
		Name:             product.Name,
		Description:      product.Description,
		Color:            product.Color,
		Size:             product.Size,
		ImageUrl:         product.ImageUrl,
		Price:            product.Price,
		Assignations:     ToAssignationsResponse(product.MaterialProduct, language),
		ProductVariation: ToProductVariationResponse(product.ProductVariation),
		Stock:            product.Stock,
	}
}

func ToAssignationsResponse(materialProducts []entities.MaterialProduct, language string) []AssignationResponse {
	var assignationsResponse []AssignationResponse
	for _, materialProduct := range materialProducts {
		assignationsResponse = append(assignationsResponse, ToAssignationResponse(materialProduct, language))
	}
	return assignationsResponse
}

func ToAssignationResponse(materialProduct entities.MaterialProduct, language string) AssignationResponse {
	return AssignationResponse{
		Quantity: materialProduct.Quantity,
		Material: ToMaterialResponse(materialProduct.Material, language),
	}
}

func ToMaterialResponse(material entities.Material, language string) MaterialResponse {
	return MaterialResponse{
		Id:           float64(material.ID),
		Name:         material.Name,
		MaterialType: ToMaterialTypeResponse(material.MaterialType, language),
	}
}

func ToMaterialTypeResponse(materialType entities.MaterialType, language string) MaterialTypeResponse {
	return MaterialTypeResponse{
		Id:                      float64(materialType.ID),
		Name:                    materialType.Name,
		UnitOfMeasurement:       materialType.UnitOfMeasurement.String(language),
		UnitOfMeasurementSymbol: enums.GetSymbolByUnitOfMeasurementEnum(materialType.UnitOfMeasurement),
	}
}

func ToMaterialsProductEntity(assignationsRequest []AssignationRequest) []entities.MaterialProduct {
	var materialsProduct []entities.MaterialProduct
	for _, assignationRequest := range assignationsRequest {
		materialsProduct = append(materialsProduct, ToMaterialProductEntity(assignationRequest))
	}
	return materialsProduct
}

func ToMaterialProductEntity(assignationRequest AssignationRequest) entities.MaterialProduct {
	return entities.MaterialProduct{
		Material: entities.Material{
			EntitiesBase: core.EntitiesBase{
				ID: uint(assignationRequest.MaterialId),
			},
		},
		Quantity: assignationRequest.Quantity,
	}
}

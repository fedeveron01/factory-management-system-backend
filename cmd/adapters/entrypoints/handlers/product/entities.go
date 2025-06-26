package product_handler

type MaterialsProductAssignationRequest struct {
	ProductId    float64              `json:"productId"`
	Assignations []AssignationRequest `json:"assignations"`
}

type AssignationRequest struct {
	MaterialId float64 `json:"materialId"`
	Quantity   float64 `json:"quantity"`
}

type ProductResponse struct {
	Id          float64 `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Color       string  `json:"color"`
	ImageUrl    string  `json:"imageUrl"`
	Price       float64 `json:"price"`
}

type ProductWithStockResponse struct {
	Id          float64 `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Color       string  `json:"color"`
	ImageUrl    string  `json:"imageUrl"`
	Price       float64 `json:"price"`
	Stock       float64 `json:"stock"`
}

type ProductVariationResponse struct {
	Id     float64 `json:"id"`
	Number float64 `json:"number"`
	Stock  float64 `json:"stock"`
}

type ProductWithAssignationsResponse struct {
	Id               float64                    `json:"id"`
	Name             string                     `json:"name"`
	Description      string                     `json:"description"`
	Color            string                     `json:"color"`
	Size             float64                    `json:"size"`
	ImageUrl         string                     `json:"imageUrl"`
	Price            float64                    `json:"price"`
	Assignations     []AssignationResponse      `json:"assignations"`
	ProductVariation []ProductVariationResponse `json:"productVariation"`
	Stock            float64                    `json:"stock"`
}

type ProductWithVariationsResponse struct {
	Id               float64                    `json:"id"`
	Name             string                     `json:"name"`
	Description      string                     `json:"description"`
	Color            string                     `json:"color"`
	ImageUrl         string                     `json:"imageUrl"`
	Price            float64                    `json:"price"`
	ProductVariation []ProductVariationResponse `json:"productVariation"`
	Stock            float64                    `json:"stock"`
}

type AssignationResponse struct {
	Quantity float64          `json:"quantity"`
	Material MaterialResponse `json:"material"`
}

type MaterialResponse struct {
	Id           float64              `json:"id"`
	Name         string               `json:"name"`
	MaterialType MaterialTypeResponse `json:"materialType"`
}

type MaterialTypeResponse struct {
	Id                      float64 `json:"id"`
	Name                    string  `json:"name"`
	UnitOfMeasurement       string  `json:"unitOfMeasurement"`
	UnitOfMeasurementSymbol string  `json:"unitOfMeasurementSymbol"`
}

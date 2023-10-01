package material_type_handler

type MaterialTypeResponse struct {
	Id                float64 `json:"id"`
	Name              string  `json:"name"`
	Description       string  `json:"description"`
	UnitOfMeasurement string  `json:"unitOfMeasurement"`
}

type UnitOfMeasurementResponse struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

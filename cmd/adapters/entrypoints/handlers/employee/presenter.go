package employee_handler

import "github.com/fedeveron01/golang-base/cmd/core/entities"

func ToEmployeeResponse(employee entities.Employee) EmployeeResponse {
	return EmployeeResponse{
		Id:       employee.ID,
		Name:     employee.Name,
		LastName: employee.LastName,
		DNI:      employee.DNI,
		Charge: ChargeResponse{
			Id:   employee.Charge.ID,
			Name: employee.Charge.Name,
		},
		User: UserResponse{
			Id:       employee.User.ID,
			UserName: employee.User.UserName,
			Inactive: employee.User.Inactive,
		},
	}
}

func ToEmployeeResponses(employees []entities.Employee) []EmployeeResponse {
	var employeeResponses []EmployeeResponse
	for _, employee := range employees {
		employeeResponses = append(employeeResponses, ToEmployeeResponse(employee))
	}
	return employeeResponses
}
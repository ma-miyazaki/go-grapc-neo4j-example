package handler

import (
	"context"

	"github.com/ma-miyazaki/go-grpc-neo4j-example/pb/employee"
)

type mockEmployeeHandler struct {
	employee.EmployeeServiceServer
}

func NewEmployeeHandler() employee.EmployeeServiceServer {
	return &mockEmployeeHandler{}
}

func (mockEmployeeHandler) AddEmployee(context.Context, *employee.AddEmployeeRequest) (*employee.Employee, error) {
	employee := &employee.Employee{
		Id:        "9141a45b-de24-4bd4-8f84-bb0e5398e02b",
		Email:     "sample@example.com",
		LastName:  "山田",
		FirstName: "太郎",
	}
	return employee, nil
}

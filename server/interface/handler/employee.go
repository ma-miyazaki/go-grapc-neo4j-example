package handler

import (
	"context"
	"log"

	"github.com/ma-miyazaki/go-grpc-neo4j-example/pb/employee"
	"github.com/ma-miyazaki/go-grpc-neo4j-example/server/usecase"
)

type mockEmployeeHandler struct {
	employee.EmployeeServiceServer
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

type employeeHandler struct {
	employee.EmployeeServiceServer
	usecase usecase.EmployeeUseCase
}

func (handler employeeHandler) AddEmployee(_ context.Context, in *employee.AddEmployeeRequest) (*employee.Employee, error) {
	log.Printf("サーバの受け取り [%v]", in)
	em, err := handler.usecase.AddEmployee(in.Email, in.LastName, in.FirstName)
	if err != nil {
		return nil, err
	}
	return &employee.Employee{
		Id:        em.ID.String(),
		Email:     em.Email,
		LastName:  em.LastName,
		FirstName: em.FirstName,
	}, nil
}

func NewEmployeeHandler(usecase usecase.EmployeeUseCase) employee.EmployeeServiceServer {
	// return &mockEmployeeHandler{}
	return &employeeHandler{usecase: usecase}
}

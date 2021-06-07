package handler

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/ma-miyazaki/go-grpc-neo4j-example/pb/employee"
	"github.com/ma-miyazaki/go-grpc-neo4j-example/server/usecase"
	"github.com/rs/zerolog/log"
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
	log.Info().Msgf("Recieved a create emplopyee request [%v]", in)
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

func (handler employeeHandler) ListEmployees(context.Context, *empty.Empty) (*employee.ListEmployeesReply, error) {
	log.Info().Msg("Recieved a list emplopyees request")
	employees, err := handler.usecase.ListEmployees()
	if err != nil {
		return nil, err
	}

	list := make([]*employee.Employee, len(employees))
	for _, em := range employees {
		list = append(list, &employee.Employee{
			Id:        em.ID.String(),
			Email:     em.Email,
			LastName:  em.LastName,
			FirstName: em.FirstName,
		})
	}
	return &employee.ListEmployeesReply{Employees: list}, nil
}

func NewEmployeeHandler(usecase usecase.EmployeeUseCase) employee.EmployeeServiceServer {
	// return mockEmployeeHandler{}
	return employeeHandler{usecase: usecase}
}

package usecase

import (
	"github.com/ma-miyazaki/go-grpc-neo4j-example/server/domain/model"
	"github.com/ma-miyazaki/go-grpc-neo4j-example/server/domain/repository"
)

type EmployeeUseCase interface {
	AddEmployee(email string, lastName string, firstName string) (*model.Employee, error)
}

type employeeUseCase struct {
	repository repository.EmployeeRepository
}

func NewEmployeeUseCase(repository repository.EmployeeRepository) EmployeeUseCase {
	return employeeUseCase{repository}
}

func (employeeUseCase) AddEmployee(email string, lastName string, firstName string) (*model.Employee, error) {
	employee, err := model.NewEmployee(email, lastName, firstName)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

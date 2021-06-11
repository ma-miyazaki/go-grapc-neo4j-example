package usecase

import (
	"fmt"

	"github.com/ma-miyazaki/go-grpc-neo4j-example/server/domain/model"
	"github.com/ma-miyazaki/go-grpc-neo4j-example/server/domain/repository"
	"github.com/ma-miyazaki/go-grpc-neo4j-example/server/domain/service"
	"github.com/rs/zerolog/log"
)

type EmployeeUseCase interface {
	AddEmployee(email string, lastName string, firstName string) (*model.Employee, error)
	ListEmployees() ([]*model.Employee, error)
}

type employeeUseCase struct {
	repository repository.EmployeeRepository
	service    service.EmployeeService
}

func NewEmployeeUseCase(repository repository.EmployeeRepository, service service.EmployeeService) EmployeeUseCase {
	return employeeUseCase{repository: repository, service: service}
}

func (uc employeeUseCase) AddEmployee(email string, lastName string, firstName string) (employee *model.Employee, err error) {
	err = uc.repository.DoInTransaction(func() error {
		employee, err = model.NewEmployee(email, lastName, firstName)
		if err != nil {
			log.Fatal().Stack().Err(err).Msg("failed to create employee model")
			return err
		}

		if uc.service.IsDuplicated(*employee) {
			return fmt.Errorf("Employee is duplicated. %v", employee)
		}

		if err = uc.repository.Create(employee); err != nil {
			log.Fatal().Stack().Err(err).Msg("failed to store employee")
			return err
		}

		log.Info().Msgf("Employee created. [%v]", employee)
		return nil
	})
	return
}

func (uc employeeUseCase) ListEmployees() ([]*model.Employee, error) {
	return uc.repository.List()
}

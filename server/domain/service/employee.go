package service

import (
	"github.com/ma-miyazaki/go-grpc-neo4j-example/server/domain/model"
	"github.com/ma-miyazaki/go-grpc-neo4j-example/server/domain/repository"
	"github.com/rs/zerolog/log"
)

type EmployeeService interface {
	IsDuplicated(employee model.Employee) bool
}

type employeeService struct {
	repository repository.EmployeeRepository
}

func NewEmployeeService(repository repository.EmployeeRepository) EmployeeService {
	return &employeeService{repository: repository}
}

func (srv *employeeService) IsDuplicated(employee model.Employee) bool {
	result, err := srv.repository.FindByEmail(employee.Email)
	log.Info().Msgf("FindByEmail. [%v][%v]", result, err)
	return result != nil
}

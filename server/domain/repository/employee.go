package repository

import (
	"github.com/ma-miyazaki/go-grpc-neo4j-example/server/domain/model"
)

type EmployeeRepository interface {
	Repository
	Create(employee *model.Employee) error
	List() ([]*model.Employee, error)
}

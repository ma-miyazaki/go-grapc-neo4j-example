package repository

import (
	"github.com/ma-miyazaki/go-grpc-neo4j-example/server/domain/model"
)

type EmployeeRepository interface {
	Create(employee *model.Employee) error
}

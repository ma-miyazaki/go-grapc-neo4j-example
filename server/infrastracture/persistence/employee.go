package persistence

import (
	"github.com/ma-miyazaki/go-grpc-neo4j-example/server/domain/model"
	"github.com/ma-miyazaki/go-grpc-neo4j-example/server/domain/repository"
)

type neo4jEmployeeRepository struct {
}

func NewEmployeeRepository() repository.EmployeeRepository {
	return &neo4jEmployeeRepository{}
}

const createEmployeeQuery = "CREATE (:Person {uuid: $uuid, email: $email, lastName: $lastName, firstName: $firstName})"

func createEmployeeParams(employee *model.Employee) map[string]interface{} {
	return map[string]interface{}{
		"uuid":      employee.ID.String(),
		"email":     employee.Email,
		"lastName":  employee.LastName,
		"firstName": employee.FirstName,
	}
}

func (repository neo4jEmployeeRepository) Create(employee *model.Employee) error {
	session, err := NewNeo4jSession()
	if err != nil {
		return err
	}

	result, err := session.Run(createEmployeeQuery, createEmployeeParams(employee))
	if err != nil {
		return err
	}
	return result.Err()
}

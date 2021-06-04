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

func (repository neo4jEmployeeRepository) Create(employee *model.Employee) error {
	var query = ""
	var params = map[string]interface{}{}

	session, err := NewNeo4jSession()
	if err != nil {
		return err
	}
	result, err := session.Run(query, params)
	if err != nil {
		return err
	}
	return result.Err()
}

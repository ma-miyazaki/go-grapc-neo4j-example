package persistence

import (
	"github.com/google/uuid"
	"github.com/ma-miyazaki/go-grpc-neo4j-example/server/domain/model"
	"github.com/ma-miyazaki/go-grpc-neo4j-example/server/domain/repository"
)

type neo4jEmployeeRepository struct {
}

func NewEmployeeRepository() repository.EmployeeRepository {
	return neo4jEmployeeRepository{}
}

const createEmployeeQuery = "CREATE (:Person {uuid: $uuid, email: $email, lastName: $lastName, firstName: $firstName})"
const listEmployeeQuery = "MATCH (p:Person) RETURN p.uuid, p.email, p.lastName, p.firstName"

func createEmployeeParams(employee *model.Employee) map[string]interface{} {
	return map[string]interface{}{
		"uuid":      employee.ID.String(),
		"email":     employee.Email,
		"lastName":  employee.LastName,
		"firstName": employee.FirstName,
	}
}

func (repository neo4jEmployeeRepository) Create(employee *model.Employee) error {
	result, err := NewNeo4jSession().Run(createEmployeeQuery, createEmployeeParams(employee))
	if err != nil {
		return err
	}
	return result.Err()
}

func (repository neo4jEmployeeRepository) List() ([]*model.Employee, error) {
	result, err := NewNeo4jSession().Run(listEmployeeQuery, nil)
	if err != nil {
		return nil, err
	}
	if err := result.Err(); err != nil {
		return nil, err
	}

	var employees []*model.Employee
	for result.Next() {
		record := result.Record()
		id := uuid.MustParse(record.Values[0].(string))
		employee := &model.Employee{
			ID:        model.EmployeeID{id},
			Email:     record.Values[1].(string),
			LastName:  record.Values[2].(string),
			FirstName: record.Values[3].(string),
		}
		employees = append(employees, employee)
	}

	return employees, nil
}

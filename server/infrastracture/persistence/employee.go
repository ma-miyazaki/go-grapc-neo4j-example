package persistence

import (
	"github.com/google/uuid"
	"github.com/ma-miyazaki/go-grpc-neo4j-example/server/domain/model"
	"github.com/ma-miyazaki/go-grpc-neo4j-example/server/domain/repository"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type neo4jEmployeeRepository struct {
	neo4jRepository
}

func NewEmployeeRepository() repository.EmployeeRepository {
	return &neo4jEmployeeRepository{}
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

func (repository *neo4jEmployeeRepository) Create(employee *model.Employee) error {
	_, err := repository.run(createEmployeeQuery, createEmployeeParams(employee), func(result neo4j.Result) (interface{}, error) { return nil, nil })
	return err
}

func (repository *neo4jEmployeeRepository) List() ([]*model.Employee, error) {
	result, err := repository.run(listEmployeeQuery, nil, func(result neo4j.Result) (interface{}, error) {
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
	})
	if err != nil {
		return nil, err
	}
	return result.([]*model.Employee), nil
}

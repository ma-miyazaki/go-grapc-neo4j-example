package model

import (
	"github.com/google/uuid"
)

type Employee struct {
	ID        EmployeeID
	Email     string
	LastName  string
	FirstName string
}

type EmployeeID struct {
	uuid.UUID
}

func NewEmployee(email string, lastName string, firstName string) (*Employee, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	employee := &Employee{
		ID:        EmployeeID{id},
		Email:     email,
		LastName:  lastName,
		FirstName: firstName,
	}
	return employee, nil
}

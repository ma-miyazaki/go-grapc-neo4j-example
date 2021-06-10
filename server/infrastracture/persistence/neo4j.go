package persistence

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/rs/zerolog/log"
)

var driver = newNeo4jDriver()

func newNeo4jDriver() neo4j.Driver {
	driver, err := neo4j.NewDriver("bolt://neo4j:7687", neo4j.NoAuth())
	if err != nil {
		log.Fatal().Stack().Err(err).Msg("failed to open neo4j connection")
		panic(err)
	}
	return driver
}

func CloseNeo4jDriver() {
	driver.Close()
}

func NewNeo4jSession() neo4j.Session {
	return driver.NewSession(neo4j.SessionConfig{})
}

type neo4jRepository struct {
	transaction neo4j.Transaction
}

func (repository *neo4jRepository) DoInTransaction(fx func() error) (err error) {
	session := NewNeo4jSession()
	defer session.Close()
	repository.transaction, err = session.BeginTransaction()
	if err != nil {
		return
	}

	err = fx()
	if err != nil {
		repository.transaction.Rollback()
		return
	}

	repository.transaction.Commit()
	return
}

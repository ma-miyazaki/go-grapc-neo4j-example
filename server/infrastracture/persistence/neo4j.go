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
	neo4j.Transaction
	session neo4j.Session
}

func (repository *neo4jRepository) Begin() (err error) {
	repository.session = NewNeo4jSession()
	repository.Transaction, err = repository.session.BeginTransaction()
	log.Info().Msgf("transaction: %p", &repository.Transaction)
	return err
}

func (repository *neo4jRepository) Close() {
	if repository.Transaction != nil {
		repository.Transaction.Close()
	}
	if repository.session != nil {
		repository.session.Close()
	}
}

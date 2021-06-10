package persistence

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/rs/zerolog/log"
)

var neo4jDriver = newNeo4jDriver()

func newNeo4jDriver() neo4j.Driver {
	driver, err := neo4j.NewDriver("bolt://neo4j:7687", neo4j.NoAuth())
	if err != nil {
		log.Fatal().Stack().Err(err).Msg("failed to open neo4j connection")
		panic(err)
	}
	return driver
}

func CloseNeo4jDriver() {
	neo4jDriver.Close()
}

type neo4jRepository struct {
	transaction neo4j.Transaction
}

func (repository *neo4jRepository) DoInTransaction(fx func() error) error {
	session := neo4jDriver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		repository.transaction = tx
		err := fx()
		return nil, err
	})
	return err
}

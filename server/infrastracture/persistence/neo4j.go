package persistence

import (
	"log"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

var driver = newNeo4jDriver()

func newNeo4jDriver() neo4j.Driver {
	driver, err := neo4j.NewDriver("bolt://neo4j:7687", neo4j.NoAuth(), func(config *neo4j.Config) {
		config.Encrypted = false
	})
	if err != nil {
		log.Fatalf("%v", err)
		panic(err)
	}
	return driver
}

func CloseNeo4jDriver() {
	driver.Close()
}

func NewNeo4jSession() (neo4j.Session, error) {
	return driver.Session(neo4j.AccessModeWrite)
}

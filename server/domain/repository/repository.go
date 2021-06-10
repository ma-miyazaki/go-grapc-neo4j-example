package repository

// import "github.com/neo4j/neo4j-go-driver/v4/neo4j"

type Repository interface {
	Begin() error
	Commit() error
	Close()
	// Run(query string, params map[string]interface{}) (neo4j.Result, error)
}

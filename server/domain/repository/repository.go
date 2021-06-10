package repository

type Repository interface {
	DoInTransaction(func() error) error
}

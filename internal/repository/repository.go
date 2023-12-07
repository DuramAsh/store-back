package repository

import (
	"store-back/internal/domain/product"
	"store-back/internal/repository/postgres"
	"store-back/pkg/store"
)

// Configuration is an alias for a function that will take in a pointer to a Repository and modify it
type Configuration func(r *Repository) error

// Repository is an implementation of the Repository
type Repository struct {
	Postgres store.SQLX

	Product product.Repository
}

// New takes a variable amount of Configuration functions and returns a new Repository
// Each Configuration will be called in the order they are passed in
func New(configs ...Configuration) (s *Repository, err error) {
	// Create the repository
	s = &Repository{}

	// Apply all Configurations passed in
	for _, cfg := range configs {
		// Pass the repository into the configuration function
		if err = cfg(s); err != nil {
			return
		}
	}

	return
}

// Close closes the repository and prevents new queries from starting.
// Close then waits for all queries that have started processing on the server to finish.
func (r *Repository) Close() {
	if r.Postgres.Client != nil {
		r.Postgres.Client.Close()
	}
}

// WithPostgresStore applies a postgres store to the Repository
func WithPostgresStore(dataSourceName string) Configuration {
	return func(s *Repository) (err error) {
		// Create the postgres store, if we needed parameters, such as connection strings they could be inputted here
		s.Postgres, err = store.NewSQL(dataSourceName)
		if err != nil {
			return
		}

		if err = store.Migrate(dataSourceName); err != nil {
			return
		}

		s.Product = postgres.NewProductRepository(s.Postgres.Client)

		return
	}
}

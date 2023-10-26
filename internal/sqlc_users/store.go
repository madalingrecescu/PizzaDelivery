package db

import "database/sql"

// Store provides all functions to execute db queries and transactions
type Store struct {
	db *sql.DB
	*Queries
}

// NewStore creates a bew store

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

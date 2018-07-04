package storage

import "fmt"

// Creatable are the functions necessary to be able to call Storage.Create
type Creatable interface {
	Table() string
	Fields() []string
	Values() []interface{}
}

// QueryBuilder is what each implementation must implement to build queries in their language
type QueryBuilder interface {
	Create(Creatable) string
}

// Storage is the struct with the storing methods
type Storage struct {
	Database
	QueryBuilder
}

// New returns a struct with the methods to interact with a db
func New(db Database, qb QueryBuilder) *Storage {
	return &Storage{
		Database:     db,
		QueryBuilder: qb,
	}
}

// Create will save the object without making sure it exists returning the last inserted id and an error.
func (s *Storage) Create(creatable Creatable) (int64, error) {
	query := s.QueryBuilder.Create(creatable)
	r, err := s.Exec(query, creatable.Values()...)
	if err != nil {
		return int64(-1), fmt.Errorf("failed to run query %q: %v", query, err)
	}
	return r.LastInsertId()
}

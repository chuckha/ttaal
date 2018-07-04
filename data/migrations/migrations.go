package migrations

import (
	"database/sql"
)

// Database are the functions required to interact with a database
type Database interface {
	Exec(string, ...interface{}) (sql.Result, error)
}

type MigratorDatabase interface {
	Database
	QueryRow(string, ...interface{}) *sql.Row
}

// Migrator is a thing to help you with db migrations
type Migrator struct {
	Database         MigratorDatabase
	MetaQueryBuilder MetaQueryBuilder
}

// MetaQueryBuilder are the required methods for generating meta queries (questions about the database/tables themselves)
type MetaQueryBuilder interface {
	TableExistsQuery(string, string) string
	CreateTableQuery(string, interface{}) string
}

// New returns a migrator
func New(db MigratorDatabase, mqb MetaQueryBuilder) *Migrator {
	return &Migrator{
		Database:         db,
		MetaQueryBuilder: mqb,
	}
}

// TableExists tells us if a table exists or not
func (m *Migrator) TableExists(db, tableName string) bool {
	query := m.MetaQueryBuilder.TableExistsQuery(db, tableName)
	row := m.Database.QueryRow(query)
	var out int
	return row.Scan(&out) == nil
}

// CreateTable creates a table
func (m *Migrator) CreateTable(tableName string, instance interface{}) error {
	query := m.MetaQueryBuilder.CreateTableQuery(tableName, instance)
	_, err := m.Database.Exec(query)
	return err
}

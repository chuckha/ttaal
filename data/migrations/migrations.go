package migrations

import (
	"database/sql"

	"github.com/chuckha/ttaal/data/storage"
)

// MigratorDatabase extends the typical storage.Database interface.
// Another way to think about it is that it's the same as a storage.Database but with an extended set of functions.
// I expect storage.Database will have all of these methods and this interface can collapse.
type MigratorDatabase interface {
	storage.Database
	QueryRow(string, ...interface{}) *sql.Row
}

// Migrator ties together the connection and the query builder.
type Migrator struct {
	Database         MigratorDatabase
	MetaQueryBuilder MetaQueryBuilder
}

// MetaQueryBuilder describes the behavior required for generating metaqueries (questions about the database/tables themselves).
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

// TableExists returns true if the table exists and false if it does not.
func (m *Migrator) TableExists(db, tableName string) bool {
	query := m.MetaQueryBuilder.TableExistsQuery(db, tableName)
	row := m.Database.QueryRow(query)
	var out int
	return row.Scan(&out) == nil
}

// CreateTable generates a create table query and executes it.
func (m *Migrator) CreateTable(tableName string, instance interface{}) error {
	query := m.MetaQueryBuilder.CreateTableQuery(tableName, instance)
	_, err := m.Database.Exec(query)
	return err
}

package mysql

import (
	"database/sql"
	"fmt"

	"github.com/chuckha/ttaal/data/migrations"
	"github.com/chuckha/ttaal/data/storage"

	// import the mysql driver
	_ "github.com/go-sql-driver/mysql"
)

const (
	metadatabaseName = "information_schema"
	databaseName     = "ttaal"
	databaseUser     = "root"
	databasePassword = ""
	databaseHost     = "127.0.0.1"
	databasePort     = 3306

	queryPlaceholder = "?"
	// table, fields, field values
	createQuery = `INSERT INTO %v (%v) VALUES (%v)`
)

// mysql DSN [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
func dsn(database string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", databaseUser, databasePassword, databaseHost, databasePort, database)
}

func conn(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return db
}

// Storage returns a storage implementation with mysql as the underlying store
func Storage(log logger) *storage.Storage {
	details := dsn(databaseName)
	log.Debugln("Connecting to:", details)
	db := conn(details)
	return storage.New(db, &queryBuilder{log})
}

// Migrations returns a migrator for mysql
// Migrations require different permissions from above since this queries mysql tables directly
func Migrations(log logger) *migrations.Migrator {
	mqb := &metaQueryBuilder{log}
	return migrations.New(conn(dsn(metadatabaseName)), mqb)
}

type logger interface {
	Debugf(string, ...interface{})
	Debugln(...interface{})
}

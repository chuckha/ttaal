package storage

import "database/sql"

// Database are the functions required to interact with a database
type Database interface {
	Exec(string, ...interface{}) (sql.Result, error)
}

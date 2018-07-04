package mysql

import (
	"database/sql"
	"fmt"

	"github.com/chuckha/ttaal/data/repr"
)

// Representer converts various data into a *Table representation.
type Representer struct{}

// RepresentationFromRows converts sql rows into a Table.
/*
	+-------------+-------------------+-------------+-----------+
	| COLUMN_NAME | COLUMN_DEFAULT    | IS_NULLABLE | DATA_TYPE |
	+-------------+-------------------+-------------+-----------+
	| id          | 0                 | YES         | int       |
	| user_id     | NULL              | YES         | text      |
	| ended       | 0                 | YES         | tinyint   |
	| preamble    | NULL              | YES         | text      |
	| created     | CURRENT_TIMESTAMP | NO          | timestamp |
	| updated     | CURRENT_TIMESTAMP | NO          | timestamp |
	+-------------+-------------------+-------------+-----------+
*/
func (r *Representer) RepresentationFromRows(rows *sql.Rows) (*repr.Table, error) {
	t := &repr.Table{
		Fields: make([]*repr.Field, 0),
	}
	defer rows.Close()
	for rows.Next() {
		f := &repr.Field{}
		var nullable string
		err := rows.Scan(&f.Name, &f.Default, &nullable, &f.DataType)
		if err != nil {
			return nil, fmt.Errorf("could not scan row: %v", err)
		}
		f.Nullable = nullable == "YES"
		t.Fields = append(t.Fields, f)
	}
	return t, rows.Err()
}

// RepresentationFromInterface converts a model into a Table.
func (r *Representer) RepresentationFromInterface(model interface{}) *repr.Table {
	return nil
}

package repr

import "database/sql"

// Table is how we represent a storage table in go.
type Table struct {
	Name   string
	Fields []*Field
}

// Same compares two tables and returns true if they are the same and false if they are not.
func (t *Table) Same(t2 *Table) bool {
	return false
}

// Field is one field (of a struct/table) represented in go.
type Field struct {
	Name     string
	Default  string
	Nullable bool
	DataType string
}

// Representer converts various data into a *Table representation.
type Representer struct{}

// RepresentationFromRows converts sql rows into a Table.
func (r *Representer) RepresentationFromRows(rows *sql.Rows) *Table {
	return nil
}

// RepresentationFromInterface converts a model into a Table.
func (r *Representer) RepresentationFromInterface(model interface{}) *Table {
	return nil
}

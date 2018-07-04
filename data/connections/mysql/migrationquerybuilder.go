package mysql

import (
	"fmt"
	"reflect"
	"strings"
)

type metaQueryBuilder struct {
	logger
}

const (
	existsQueryFormat = "select 1 from tables where table_schema = '%s' and table_name = '%s' limit 1;"
)

func (m *metaQueryBuilder) TableExistsQuery(db, table string) string {
	q := fmt.Sprintf(existsQueryFormat, db, table)
	m.logger.Debugln("Exists query:", q)
	return q
}

// CreateTableQuery returns the create table query for the instance type
func (m *metaQueryBuilder) CreateTableQuery(tableName string, model interface{}) string {
	tab := &table{
		name: tableName,
		cols: make([]*col, 0),
	}
	t := reflect.TypeOf(model).Elem()
	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		// Ignore unexported fields
		if sf.Name[0] == strings.ToLower(sf.Name)[0] {
			continue
		}
		sqlTagValue := sf.Tag.Get("sql")
		// Ignore unmarked fields
		if sqlTagValue == "" {
			continue
		}
		col := newCol(sf.Type)
		for _, tag := range strings.Split(sqlTagValue, ",") {
			col.Add(tag)
		}
		tab.cols = append(tab.cols, col)
	}
	createQuery := tab.createQuery()
	m.logger.Debugln(createQuery)
	return createQuery
}

// table is the metadata information of a table.
// Things that relate to no particular field go here.
// Possible additions include engine, constraints, and encoding.
type table struct {
	name string
	cols []*col
}

// createQuery will generate a create table statement.
// All necessary information to build the create table statement is available to this function.
func (t *table) createQuery() string {
	primaryKeys := make([]string, 0)
	createStmts := make([]string, 0)
	out := fmt.Sprintf("CREATE TABLE %s.%s (\n", databaseName, t.name)
	for _, col := range t.cols {
		if col.primary {
			primaryKeys = append(primaryKeys, fmt.Sprintf("`%v`", col.name))
		}
		createStmts = append(createStmts, col.String())
	}
	createStmts = append(createStmts, fmt.Sprintf("PRIMARY KEY(%v)", strings.Join(primaryKeys, ",")))
	return out + strings.Join(createStmts, ",\n") + "\n)"
}

// col is one column in a table.
// It is kind of like the glue between go struct field and mysql field.
type col struct {
	name         string
	mysqlType    string
	defaultValue string
	primary      bool
	options      map[string]string
}

// current idea is to pass the type and tags into this function then i can initialize and get
// defaults and stuff all in one go here.
func newCol(sf reflect.Type) *col {
	return &col{
		// start with defaults, override in add
		mysqlType: reflectTypeToMySQLType(sf),
		options: map[string]string{
			"default": reflectTypeToDefaultValue(sf),
		},
	}
}

func (c *col) Add(tag string) {
	switch tag {
	case "primary":
		c.primary = true
	case "autoinc":
		c.options["autoinc"] = "AUTO_INCREMENT"
	case "null":
		c.options["nullable"] = "NULL"
	default: // assume it's the name
		c.name = tag
	}
}

func (c *col) String() string {
	out := make([]string, 0)
	out = append(out, fmt.Sprintf("`%v`", c.name))
	out = append(out, c.mysqlType)
	//  Delete the default if this is a primary key
	if c.primary {
		delete(c.options, "default")
		delete(c.options, "nullable")
	}
	for k, v := range c.options {
		if v == "" {
			continue
		}
		if k == "default" {
			out = append(out, fmt.Sprintf("DEFAULT %s", v))
			continue
		}
		out = append(out, v)
	}
	return strings.Join(out, " ")
}

func reflectTypeToDefaultValue(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Int64, reflect.Bool:
		return "0"
	case reflect.String:
		return ""
	default:
		switch t.PkgPath() {
		case "time":
			switch t.Name() {
			case "Time":
				return "CURRENT_TIMESTAMP"
			default:
				fmt.Println("unknown struct from 'package time'")
				return ""
			}
		default:
			fmt.Println("[unknown package " + t.PkgPath() + "]")
			return ""
		}
	}
}

func reflectTypeToMySQLType(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Int64:
		return "INT"
	case reflect.Bool:
		return "TINYINT"
	case reflect.String:
		return "TEXT"
	default:
		switch t.PkgPath() {
		case "time":
			switch t.Name() {
			case "Time":
				return "TIMESTAMP"
			default:
				fmt.Println("unknown struct from 'package time'")
				return ""
			}
		default:
			fmt.Println("[unknown package " + t.PkgPath() + "]")
			return ""
		}
	}
}

// 	_, err := db.Exec(`CREATE TABLE statements (
// 		id INT PRIMARY KEY AUTO_INCREMENT,
// 		poll_id INT,
// 		is_a_lie TINYINT,
// 		statement TEXT,
// 		created TIMESTAMP,
// 		updated TIMESTAMP
// 		)`)

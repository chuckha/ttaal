package models

import (
	"reflect"
	"strings"
)

// Data implements Creatable
type Data struct {
	TableName   string
	FieldNames  []string
	FieldValues []interface{}
}

// NewData returns a Creatable version of your model
func NewData(s interface{}, table string) *Data {
	t := reflect.TypeOf(s).Elem()
	v := reflect.ValueOf(s).Elem()
	fields := make([]string, 0)
	values := make([]interface{}, 0)
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
		fields = append(fields, sf.Tag.Get("sql"))
		values = append(values, v.Field(i).Interface())
	}

	return &Data{
		TableName:   table,
		FieldNames:  fields,
		FieldValues: values,
	}
}

// Table returns the name of the table to store the data in
func (d *Data) Table() string { return d.TableName }

// Fields returns the list of fields in the table
func (d *Data) Fields() []string { return d.FieldNames }

// Values returns the list of values corresponding to the table
func (d *Data) Values() []interface{} { return d.FieldValues }

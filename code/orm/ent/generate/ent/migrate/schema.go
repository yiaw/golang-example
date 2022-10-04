// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// PersonColumns holds the columns for the "person" table.
	PersonColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(32) primary key"}},
		{Name: "first_name", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(32) not null"}},
		{Name: "second_name", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(32) not null"}},
	}
	// PersonTable holds the schema information for the "person" table.
	PersonTable = &schema.Table{
		Name:       "person",
		Columns:    PersonColumns,
		PrimaryKey: []*schema.Column{PersonColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		PersonTable,
	}
)

func init() {
	PersonTable.Annotation = &entsql.Annotation{
		Table: "person",
	}
}

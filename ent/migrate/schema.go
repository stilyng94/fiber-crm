// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// LeadsColumns holds the columns for the "leads" table.
	LeadsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "company", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString},
	}
	// LeadsTable holds the schema information for the "leads" table.
	LeadsTable = &schema.Table{
		Name:       "leads",
		Columns:    LeadsColumns,
		PrimaryKey: []*schema.Column{LeadsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		LeadsTable,
	}
)

func init() {
}

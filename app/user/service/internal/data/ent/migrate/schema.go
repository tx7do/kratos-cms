// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// UserColumns holds the columns for the "user" table.
	UserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true, SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
		{Name: "create_time", Type: field.TypeInt64, Nullable: true},
		{Name: "update_time", Type: field.TypeInt64, Nullable: true},
		{Name: "username", Type: field.TypeString, Unique: true, Nullable: true, Size: 50},
		{Name: "nickname", Type: field.TypeString, Nullable: true, Size: 128},
		{Name: "email", Type: field.TypeString, Nullable: true, Size: 127},
		{Name: "avatar", Type: field.TypeString, Nullable: true, Size: 1023},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 1023},
		{Name: "password", Type: field.TypeString, Unique: true, Nullable: true, Size: 255},
	}
	// UserTable holds the schema information for the "user" table.
	UserTable = &schema.Table{
		Name:       "user",
		Columns:    UserColumns,
		PrimaryKey: []*schema.Column{UserColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_id",
				Unique:  false,
				Columns: []*schema.Column{UserColumns[0]},
			},
			{
				Name:    "user_id_username",
				Unique:  true,
				Columns: []*schema.Column{UserColumns[0], UserColumns[3]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		UserTable,
	}
)

func init() {
	UserTable.Annotation = &entsql.Annotation{
		Table:     "user",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}
}

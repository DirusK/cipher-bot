// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// RequestsColumns holds the columns for the "requests" table.
	RequestsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"active", "done", "expired"}, Default: "active"},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"encryption", "decryption"}},
		{Name: "algorithm", Type: field.TypeEnum, Nullable: true, Enums: []string{"AES", "RC4"}},
		{Name: "key_mode", Type: field.TypeEnum, Nullable: true, Enums: []string{"auto", "manual"}},
		{Name: "manual_key_validation", Type: field.TypeBool, Nullable: true},
		{Name: "user_id", Type: field.TypeInt},
	}
	// RequestsTable holds the schema information for the "requests" table.
	RequestsTable = &schema.Table{
		Name:       "requests",
		Columns:    RequestsColumns,
		PrimaryKey: []*schema.Column{RequestsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "requests_users_requests",
				Columns:    []*schema.Column{RequestsColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "telegram_id", Type: field.TypeInt64, Unique: true},
		{Name: "first_name", Type: field.TypeString},
		{Name: "last_name", Type: field.TypeString},
		{Name: "language", Type: field.TypeString, Nullable: true, Default: "ru"},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		RequestsTable,
		UsersTable,
	}
)

func init() {
	RequestsTable.ForeignKeys[0].RefTable = UsersTable
}

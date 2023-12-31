// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AdminUsersColumns holds the columns for the "admin_users" table.
	AdminUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "wallet_address", Type: field.TypeString},
	}
	// AdminUsersTable holds the schema information for the "admin_users" table.
	AdminUsersTable = &schema.Table{
		Name:       "admin_users",
		Columns:    AdminUsersColumns,
		PrimaryKey: []*schema.Column{AdminUsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "adminuser_wallet_address",
				Unique:  false,
				Columns: []*schema.Column{AdminUsersColumns[3]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "wallet_address", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_wallet_address",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[3]},
			},
		},
	}
	// UserActivitiesColumns holds the columns for the "user_activities" table.
	UserActivitiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "type", Type: field.TypeInt},
		{Name: "user_user_activity", Type: field.TypeInt, Nullable: true},
	}
	// UserActivitiesTable holds the schema information for the "user_activities" table.
	UserActivitiesTable = &schema.Table{
		Name:       "user_activities",
		Columns:    UserActivitiesColumns,
		PrimaryKey: []*schema.Column{UserActivitiesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_activities_users_user_activity",
				Columns:    []*schema.Column{UserActivitiesColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AdminUsersTable,
		UsersTable,
		UserActivitiesTable,
	}
)

func init() {
	UserActivitiesTable.ForeignKeys[0].RefTable = UsersTable
}

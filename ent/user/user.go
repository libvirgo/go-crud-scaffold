// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldWalletAddress holds the string denoting the wallet_address field in the database.
	FieldWalletAddress = "wallet_address"
	// EdgeUserActivity holds the string denoting the user_activity edge name in mutations.
	EdgeUserActivity = "user_activity"
	// Table holds the table name of the user in the database.
	Table = "users"
	// UserActivityTable is the table that holds the user_activity relation/edge.
	UserActivityTable = "user_activities"
	// UserActivityInverseTable is the table name for the UserActivity entity.
	// It exists in this package in order to avoid circular dependency with the "useractivity" package.
	UserActivityInverseTable = "user_activities"
	// UserActivityColumn is the table column denoting the user_activity relation/edge.
	UserActivityColumn = "user_user_activity"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldWalletAddress,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// WalletAddressValidator is a validator for the "wallet_address" field. It is called by the builders before save.
	WalletAddressValidator func(string) error
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByWalletAddress orders the results by the wallet_address field.
func ByWalletAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWalletAddress, opts...).ToFunc()
}

// ByUserActivityCount orders the results by user_activity count.
func ByUserActivityCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserActivityStep(), opts...)
	}
}

// ByUserActivity orders the results by user_activity terms.
func ByUserActivity(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserActivityStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUserActivityStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserActivityInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, UserActivityTable, UserActivityColumn),
	)
}

// Code generated by ent, DO NOT EDIT.

package authtoken

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the authtoken type in the database.
	Label = "auth_token"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldClientID holds the string denoting the client_id field in the database.
	FieldClientID = "client_id"
	// FieldAccessToken holds the string denoting the access_token field in the database.
	FieldAccessToken = "access_token"
	// FieldRefreshToken holds the string denoting the refresh_token field in the database.
	FieldRefreshToken = "refresh_token"
	// FieldExpiresIn holds the string denoting the expires_in field in the database.
	FieldExpiresIn = "expires_in"
	// Table holds the table name of the authtoken in the database.
	Table = "auth_token"
)

// Columns holds all SQL columns for authtoken fields.
var Columns = []string{
	FieldID,
	FieldUUID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldClientID,
	FieldAccessToken,
	FieldRefreshToken,
	FieldExpiresIn,
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
	// DefaultUUID holds the default value on creation for the "uuid" field.
	DefaultUUID func() string
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultClientID holds the default value on creation for the "client_id" field.
	DefaultClientID string
	// DefaultAccessToken holds the default value on creation for the "access_token" field.
	DefaultAccessToken string
	// DefaultRefreshToken holds the default value on creation for the "refresh_token" field.
	DefaultRefreshToken string
)

// OrderOption defines the ordering options for the AuthToken queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUUID orders the results by the uuid field.
func ByUUID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUUID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByClientID orders the results by the client_id field.
func ByClientID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClientID, opts...).ToFunc()
}

// ByAccessToken orders the results by the access_token field.
func ByAccessToken(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccessToken, opts...).ToFunc()
}

// ByRefreshToken orders the results by the refresh_token field.
func ByRefreshToken(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRefreshToken, opts...).ToFunc()
}

// ByExpiresIn orders the results by the expires_in field.
func ByExpiresIn(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExpiresIn, opts...).ToFunc()
}

// Code generated by ent, DO NOT EDIT.

package room

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the room type in the database.
	Label = "room"
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
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSid holds the string denoting the sid field in the database.
	FieldSid = "sid"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldStartTime holds the string denoting the start_time field in the database.
	FieldStartTime = "start_time"
	// FieldLeftTime holds the string denoting the left_time field in the database.
	FieldLeftTime = "left_time"
	// FieldEndTime holds the string denoting the end_time field in the database.
	FieldEndTime = "end_time"
	// FieldVodStatus holds the string denoting the vod_status field in the database.
	FieldVodStatus = "vod_status"
	// FieldMetadata holds the string denoting the metadata field in the database.
	FieldMetadata = "metadata"
	// Table holds the table name of the room in the database.
	Table = "room"
)

// Columns holds all SQL columns for room fields.
var Columns = []string{
	FieldID,
	FieldUUID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldName,
	FieldSid,
	FieldStatus,
	FieldStartTime,
	FieldLeftTime,
	FieldEndTime,
	FieldVodStatus,
	FieldMetadata,
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
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// DefaultSid holds the default value on creation for the "sid" field.
	DefaultSid string
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus uint8
	// DefaultVodStatus holds the default value on creation for the "vod_status" field.
	DefaultVodStatus uint8
)

// OrderOption defines the ordering options for the Room queries.
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

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// BySid orders the results by the sid field.
func BySid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSid, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByStartTime orders the results by the start_time field.
func ByStartTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartTime, opts...).ToFunc()
}

// ByLeftTime orders the results by the left_time field.
func ByLeftTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLeftTime, opts...).ToFunc()
}

// ByEndTime orders the results by the end_time field.
func ByEndTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndTime, opts...).ToFunc()
}

// ByVodStatus orders the results by the vod_status field.
func ByVodStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVodStatus, opts...).ToFunc()
}
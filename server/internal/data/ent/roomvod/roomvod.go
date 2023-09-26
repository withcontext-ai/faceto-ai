// Code generated by ent, DO NOT EDIT.

package roomvod

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the roomvod type in the database.
	Label = "room_vod"
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
	// FieldEgressID holds the string denoting the egress_id field in the database.
	FieldEgressID = "egress_id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldPlatform holds the string denoting the platform field in the database.
	FieldPlatform = "platform"
	// FieldVodType holds the string denoting the vod_type field in the database.
	FieldVodType = "vod_type"
	// FieldVodPath holds the string denoting the vod_path field in the database.
	FieldVodPath = "vod_path"
	// FieldVodURL holds the string denoting the vod_url field in the database.
	FieldVodURL = "vod_url"
	// FieldStartTime holds the string denoting the start_time field in the database.
	FieldStartTime = "start_time"
	// FieldCompleteTime holds the string denoting the complete_time field in the database.
	FieldCompleteTime = "complete_time"
	// FieldDuration holds the string denoting the duration field in the database.
	FieldDuration = "duration"
	// Table holds the table name of the roomvod in the database.
	Table = "room_vod"
)

// Columns holds all SQL columns for roomvod fields.
var Columns = []string{
	FieldID,
	FieldUUID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldName,
	FieldSid,
	FieldEgressID,
	FieldStatus,
	FieldPlatform,
	FieldVodType,
	FieldVodPath,
	FieldVodURL,
	FieldStartTime,
	FieldCompleteTime,
	FieldDuration,
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
	// DefaultEgressID holds the default value on creation for the "egress_id" field.
	DefaultEgressID string
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus uint8
	// DefaultPlatform holds the default value on creation for the "platform" field.
	DefaultPlatform uint8
	// DefaultVodType holds the default value on creation for the "vod_type" field.
	DefaultVodType uint8
	// DefaultVodPath holds the default value on creation for the "vod_path" field.
	DefaultVodPath string
	// DefaultVodURL holds the default value on creation for the "vod_url" field.
	DefaultVodURL string
	// DefaultDuration holds the default value on creation for the "duration" field.
	DefaultDuration uint64
)

// OrderOption defines the ordering options for the RoomVod queries.
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

// ByEgressID orders the results by the egress_id field.
func ByEgressID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEgressID, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByPlatform orders the results by the platform field.
func ByPlatform(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPlatform, opts...).ToFunc()
}

// ByVodType orders the results by the vod_type field.
func ByVodType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVodType, opts...).ToFunc()
}

// ByVodPath orders the results by the vod_path field.
func ByVodPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVodPath, opts...).ToFunc()
}

// ByVodURL orders the results by the vod_url field.
func ByVodURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVodURL, opts...).ToFunc()
}

// ByStartTime orders the results by the start_time field.
func ByStartTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartTime, opts...).ToFunc()
}

// ByCompleteTime orders the results by the complete_time field.
func ByCompleteTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCompleteTime, opts...).ToFunc()
}

// ByDuration orders the results by the duration field.
func ByDuration(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDuration, opts...).ToFunc()
}

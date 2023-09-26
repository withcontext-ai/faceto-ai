// Code generated by ent, DO NOT EDIT.

package roommessage

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the roommessage type in the database.
	Label = "room_message"
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
	// FieldSid holds the string denoting the sid field in the database.
	FieldSid = "sid"
	// FieldParticipantSid holds the string denoting the participant_sid field in the database.
	FieldParticipantSid = "participant_sid"
	// FieldParticipantName holds the string denoting the participant_name field in the database.
	FieldParticipantName = "participant_name"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldIsReply holds the string denoting the is_reply field in the database.
	FieldIsReply = "is_reply"
	// FieldReplyID holds the string denoting the reply_id field in the database.
	FieldReplyID = "reply_id"
	// FieldEventTime holds the string denoting the event_time field in the database.
	FieldEventTime = "event_time"
	// FieldText holds the string denoting the text field in the database.
	FieldText = "text"
	// Table holds the table name of the roommessage in the database.
	Table = "room_message"
)

// Columns holds all SQL columns for roommessage fields.
var Columns = []string{
	FieldID,
	FieldUUID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldSid,
	FieldParticipantSid,
	FieldParticipantName,
	FieldType,
	FieldIsReply,
	FieldReplyID,
	FieldEventTime,
	FieldText,
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
	// DefaultSid holds the default value on creation for the "sid" field.
	DefaultSid string
	// DefaultParticipantSid holds the default value on creation for the "participant_sid" field.
	DefaultParticipantSid string
	// DefaultParticipantName holds the default value on creation for the "participant_name" field.
	DefaultParticipantName string
	// DefaultType holds the default value on creation for the "type" field.
	DefaultType uint32
	// DefaultIsReply holds the default value on creation for the "is_reply" field.
	DefaultIsReply uint32
	// DefaultReplyID holds the default value on creation for the "reply_id" field.
	DefaultReplyID uint64
	// DefaultText holds the default value on creation for the "text" field.
	DefaultText string
)

// OrderOption defines the ordering options for the RoomMessage queries.
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

// BySid orders the results by the sid field.
func BySid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSid, opts...).ToFunc()
}

// ByParticipantSid orders the results by the participant_sid field.
func ByParticipantSid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldParticipantSid, opts...).ToFunc()
}

// ByParticipantName orders the results by the participant_name field.
func ByParticipantName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldParticipantName, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByIsReply orders the results by the is_reply field.
func ByIsReply(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsReply, opts...).ToFunc()
}

// ByReplyID orders the results by the reply_id field.
func ByReplyID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReplyID, opts...).ToFunc()
}

// ByEventTime orders the results by the event_time field.
func ByEventTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEventTime, opts...).ToFunc()
}

// ByText orders the results by the text field.
func ByText(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldText, opts...).ToFunc()
}
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"faceto-ai/internal/data/ent/roommessage"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// RoomMessage is the model entity for the RoomMessage schema.
type RoomMessage struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// UUID holds the value of the "uuid" field.
	UUID string `json:"uuid,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// Sid holds the value of the "sid" field.
	Sid string `json:"sid,omitempty"`
	// ParticipantSid holds the value of the "participant_sid" field.
	ParticipantSid string `json:"participant_sid,omitempty"`
	// ParticipantName holds the value of the "participant_name" field.
	ParticipantName string `json:"participant_name,omitempty"`
	// Type holds the value of the "type" field.
	Type uint32 `json:"type,omitempty"`
	// IsReply holds the value of the "is_reply" field.
	IsReply uint32 `json:"is_reply,omitempty"`
	// ReplyID holds the value of the "reply_id" field.
	ReplyID uint64 `json:"reply_id,omitempty"`
	// EventTime holds the value of the "event_time" field.
	EventTime time.Time `json:"event_time,omitempty"`
	// Text holds the value of the "text" field.
	Text         string `json:"text,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RoomMessage) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case roommessage.FieldID, roommessage.FieldType, roommessage.FieldIsReply, roommessage.FieldReplyID:
			values[i] = new(sql.NullInt64)
		case roommessage.FieldUUID, roommessage.FieldSid, roommessage.FieldParticipantSid, roommessage.FieldParticipantName, roommessage.FieldText:
			values[i] = new(sql.NullString)
		case roommessage.FieldCreatedAt, roommessage.FieldUpdatedAt, roommessage.FieldDeletedAt, roommessage.FieldEventTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RoomMessage fields.
func (rm *RoomMessage) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case roommessage.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			rm.ID = uint64(value.Int64)
		case roommessage.FieldUUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value.Valid {
				rm.UUID = value.String
			}
		case roommessage.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				rm.CreatedAt = value.Time
			}
		case roommessage.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				rm.UpdatedAt = value.Time
			}
		case roommessage.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				rm.DeletedAt = value.Time
			}
		case roommessage.FieldSid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sid", values[i])
			} else if value.Valid {
				rm.Sid = value.String
			}
		case roommessage.FieldParticipantSid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field participant_sid", values[i])
			} else if value.Valid {
				rm.ParticipantSid = value.String
			}
		case roommessage.FieldParticipantName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field participant_name", values[i])
			} else if value.Valid {
				rm.ParticipantName = value.String
			}
		case roommessage.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				rm.Type = uint32(value.Int64)
			}
		case roommessage.FieldIsReply:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_reply", values[i])
			} else if value.Valid {
				rm.IsReply = uint32(value.Int64)
			}
		case roommessage.FieldReplyID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field reply_id", values[i])
			} else if value.Valid {
				rm.ReplyID = uint64(value.Int64)
			}
		case roommessage.FieldEventTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field event_time", values[i])
			} else if value.Valid {
				rm.EventTime = value.Time
			}
		case roommessage.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				rm.Text = value.String
			}
		default:
			rm.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the RoomMessage.
// This includes values selected through modifiers, order, etc.
func (rm *RoomMessage) Value(name string) (ent.Value, error) {
	return rm.selectValues.Get(name)
}

// Update returns a builder for updating this RoomMessage.
// Note that you need to call RoomMessage.Unwrap() before calling this method if this RoomMessage
// was returned from a transaction, and the transaction was committed or rolled back.
func (rm *RoomMessage) Update() *RoomMessageUpdateOne {
	return NewRoomMessageClient(rm.config).UpdateOne(rm)
}

// Unwrap unwraps the RoomMessage entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rm *RoomMessage) Unwrap() *RoomMessage {
	_tx, ok := rm.config.driver.(*txDriver)
	if !ok {
		panic("ent: RoomMessage is not a transactional entity")
	}
	rm.config.driver = _tx.drv
	return rm
}

// String implements the fmt.Stringer.
func (rm *RoomMessage) String() string {
	var builder strings.Builder
	builder.WriteString("RoomMessage(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rm.ID))
	builder.WriteString("uuid=")
	builder.WriteString(rm.UUID)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(rm.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(rm.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(rm.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("sid=")
	builder.WriteString(rm.Sid)
	builder.WriteString(", ")
	builder.WriteString("participant_sid=")
	builder.WriteString(rm.ParticipantSid)
	builder.WriteString(", ")
	builder.WriteString("participant_name=")
	builder.WriteString(rm.ParticipantName)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", rm.Type))
	builder.WriteString(", ")
	builder.WriteString("is_reply=")
	builder.WriteString(fmt.Sprintf("%v", rm.IsReply))
	builder.WriteString(", ")
	builder.WriteString("reply_id=")
	builder.WriteString(fmt.Sprintf("%v", rm.ReplyID))
	builder.WriteString(", ")
	builder.WriteString("event_time=")
	builder.WriteString(rm.EventTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("text=")
	builder.WriteString(rm.Text)
	builder.WriteByte(')')
	return builder.String()
}

// RoomMessages is a parsable slice of RoomMessage.
type RoomMessages []*RoomMessage

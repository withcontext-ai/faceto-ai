// Code generated by ent, DO NOT EDIT.

package ent

import (
	"faceto-ai/internal/data/ent/roomvod"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// RoomVod is the model entity for the RoomVod schema.
type RoomVod struct {
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
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Sid holds the value of the "sid" field.
	Sid string `json:"sid,omitempty"`
	// EgressID holds the value of the "egress_id" field.
	EgressID string `json:"egress_id,omitempty"`
	// Status holds the value of the "status" field.
	Status uint8 `json:"status,omitempty"`
	// Platform holds the value of the "platform" field.
	Platform uint8 `json:"platform,omitempty"`
	// VodType holds the value of the "vod_type" field.
	VodType uint8 `json:"vod_type,omitempty"`
	// VodPath holds the value of the "vod_path" field.
	VodPath string `json:"vod_path,omitempty"`
	// VodURL holds the value of the "vod_url" field.
	VodURL string `json:"vod_url,omitempty"`
	// StartTime holds the value of the "start_time" field.
	StartTime time.Time `json:"start_time,omitempty"`
	// CompleteTime holds the value of the "complete_time" field.
	CompleteTime time.Time `json:"complete_time,omitempty"`
	// Duration holds the value of the "duration" field.
	Duration     uint64 `json:"duration,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RoomVod) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case roomvod.FieldID, roomvod.FieldStatus, roomvod.FieldPlatform, roomvod.FieldVodType, roomvod.FieldDuration:
			values[i] = new(sql.NullInt64)
		case roomvod.FieldUUID, roomvod.FieldName, roomvod.FieldSid, roomvod.FieldEgressID, roomvod.FieldVodPath, roomvod.FieldVodURL:
			values[i] = new(sql.NullString)
		case roomvod.FieldCreatedAt, roomvod.FieldUpdatedAt, roomvod.FieldDeletedAt, roomvod.FieldStartTime, roomvod.FieldCompleteTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RoomVod fields.
func (rv *RoomVod) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case roomvod.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			rv.ID = uint64(value.Int64)
		case roomvod.FieldUUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value.Valid {
				rv.UUID = value.String
			}
		case roomvod.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				rv.CreatedAt = value.Time
			}
		case roomvod.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				rv.UpdatedAt = value.Time
			}
		case roomvod.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				rv.DeletedAt = value.Time
			}
		case roomvod.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				rv.Name = value.String
			}
		case roomvod.FieldSid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sid", values[i])
			} else if value.Valid {
				rv.Sid = value.String
			}
		case roomvod.FieldEgressID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field egress_id", values[i])
			} else if value.Valid {
				rv.EgressID = value.String
			}
		case roomvod.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				rv.Status = uint8(value.Int64)
			}
		case roomvod.FieldPlatform:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field platform", values[i])
			} else if value.Valid {
				rv.Platform = uint8(value.Int64)
			}
		case roomvod.FieldVodType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field vod_type", values[i])
			} else if value.Valid {
				rv.VodType = uint8(value.Int64)
			}
		case roomvod.FieldVodPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field vod_path", values[i])
			} else if value.Valid {
				rv.VodPath = value.String
			}
		case roomvod.FieldVodURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field vod_url", values[i])
			} else if value.Valid {
				rv.VodURL = value.String
			}
		case roomvod.FieldStartTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_time", values[i])
			} else if value.Valid {
				rv.StartTime = value.Time
			}
		case roomvod.FieldCompleteTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field complete_time", values[i])
			} else if value.Valid {
				rv.CompleteTime = value.Time
			}
		case roomvod.FieldDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration", values[i])
			} else if value.Valid {
				rv.Duration = uint64(value.Int64)
			}
		default:
			rv.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the RoomVod.
// This includes values selected through modifiers, order, etc.
func (rv *RoomVod) Value(name string) (ent.Value, error) {
	return rv.selectValues.Get(name)
}

// Update returns a builder for updating this RoomVod.
// Note that you need to call RoomVod.Unwrap() before calling this method if this RoomVod
// was returned from a transaction, and the transaction was committed or rolled back.
func (rv *RoomVod) Update() *RoomVodUpdateOne {
	return NewRoomVodClient(rv.config).UpdateOne(rv)
}

// Unwrap unwraps the RoomVod entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rv *RoomVod) Unwrap() *RoomVod {
	_tx, ok := rv.config.driver.(*txDriver)
	if !ok {
		panic("ent: RoomVod is not a transactional entity")
	}
	rv.config.driver = _tx.drv
	return rv
}

// String implements the fmt.Stringer.
func (rv *RoomVod) String() string {
	var builder strings.Builder
	builder.WriteString("RoomVod(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rv.ID))
	builder.WriteString("uuid=")
	builder.WriteString(rv.UUID)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(rv.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(rv.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(rv.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(rv.Name)
	builder.WriteString(", ")
	builder.WriteString("sid=")
	builder.WriteString(rv.Sid)
	builder.WriteString(", ")
	builder.WriteString("egress_id=")
	builder.WriteString(rv.EgressID)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", rv.Status))
	builder.WriteString(", ")
	builder.WriteString("platform=")
	builder.WriteString(fmt.Sprintf("%v", rv.Platform))
	builder.WriteString(", ")
	builder.WriteString("vod_type=")
	builder.WriteString(fmt.Sprintf("%v", rv.VodType))
	builder.WriteString(", ")
	builder.WriteString("vod_path=")
	builder.WriteString(rv.VodPath)
	builder.WriteString(", ")
	builder.WriteString("vod_url=")
	builder.WriteString(rv.VodURL)
	builder.WriteString(", ")
	builder.WriteString("start_time=")
	builder.WriteString(rv.StartTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("complete_time=")
	builder.WriteString(rv.CompleteTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("duration=")
	builder.WriteString(fmt.Sprintf("%v", rv.Duration))
	builder.WriteByte(')')
	return builder.String()
}

// RoomVods is a parsable slice of RoomVod.
type RoomVods []*RoomVod

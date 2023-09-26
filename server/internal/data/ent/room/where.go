// Code generated by ent, DO NOT EDIT.

package room

import (
	"faceto-ai/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.Room {
	return predicate.Room(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.Room {
	return predicate.Room(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.Room {
	return predicate.Room(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.Room {
	return predicate.Room(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.Room {
	return predicate.Room(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.Room {
	return predicate.Room(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.Room {
	return predicate.Room(sql.FieldLTE(FieldID, id))
}

// UUID applies equality check predicate on the "uuid" field. It's identical to UUIDEQ.
func UUID(v string) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldUUID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldDeletedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldName, v))
}

// Sid applies equality check predicate on the "sid" field. It's identical to SidEQ.
func Sid(v string) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldSid, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v uint8) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldStatus, v))
}

// StartTime applies equality check predicate on the "start_time" field. It's identical to StartTimeEQ.
func StartTime(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldStartTime, v))
}

// LeftTime applies equality check predicate on the "left_time" field. It's identical to LeftTimeEQ.
func LeftTime(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldLeftTime, v))
}

// EndTime applies equality check predicate on the "end_time" field. It's identical to EndTimeEQ.
func EndTime(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldEndTime, v))
}

// VodStatus applies equality check predicate on the "vod_status" field. It's identical to VodStatusEQ.
func VodStatus(v uint8) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldVodStatus, v))
}

// UUIDEQ applies the EQ predicate on the "uuid" field.
func UUIDEQ(v string) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldUUID, v))
}

// UUIDNEQ applies the NEQ predicate on the "uuid" field.
func UUIDNEQ(v string) predicate.Room {
	return predicate.Room(sql.FieldNEQ(FieldUUID, v))
}

// UUIDIn applies the In predicate on the "uuid" field.
func UUIDIn(vs ...string) predicate.Room {
	return predicate.Room(sql.FieldIn(FieldUUID, vs...))
}

// UUIDNotIn applies the NotIn predicate on the "uuid" field.
func UUIDNotIn(vs ...string) predicate.Room {
	return predicate.Room(sql.FieldNotIn(FieldUUID, vs...))
}

// UUIDGT applies the GT predicate on the "uuid" field.
func UUIDGT(v string) predicate.Room {
	return predicate.Room(sql.FieldGT(FieldUUID, v))
}

// UUIDGTE applies the GTE predicate on the "uuid" field.
func UUIDGTE(v string) predicate.Room {
	return predicate.Room(sql.FieldGTE(FieldUUID, v))
}

// UUIDLT applies the LT predicate on the "uuid" field.
func UUIDLT(v string) predicate.Room {
	return predicate.Room(sql.FieldLT(FieldUUID, v))
}

// UUIDLTE applies the LTE predicate on the "uuid" field.
func UUIDLTE(v string) predicate.Room {
	return predicate.Room(sql.FieldLTE(FieldUUID, v))
}

// UUIDContains applies the Contains predicate on the "uuid" field.
func UUIDContains(v string) predicate.Room {
	return predicate.Room(sql.FieldContains(FieldUUID, v))
}

// UUIDHasPrefix applies the HasPrefix predicate on the "uuid" field.
func UUIDHasPrefix(v string) predicate.Room {
	return predicate.Room(sql.FieldHasPrefix(FieldUUID, v))
}

// UUIDHasSuffix applies the HasSuffix predicate on the "uuid" field.
func UUIDHasSuffix(v string) predicate.Room {
	return predicate.Room(sql.FieldHasSuffix(FieldUUID, v))
}

// UUIDEqualFold applies the EqualFold predicate on the "uuid" field.
func UUIDEqualFold(v string) predicate.Room {
	return predicate.Room(sql.FieldEqualFold(FieldUUID, v))
}

// UUIDContainsFold applies the ContainsFold predicate on the "uuid" field.
func UUIDContainsFold(v string) predicate.Room {
	return predicate.Room(sql.FieldContainsFold(FieldUUID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Room {
	return predicate.Room(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Room {
	return predicate.Room(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Room {
	return predicate.Room(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Room {
	return predicate.Room(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Room {
	return predicate.Room(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Room {
	return predicate.Room(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Room {
	return predicate.Room(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Room {
	return predicate.Room(sql.FieldNotNull(FieldDeletedAt))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Room {
	return predicate.Room(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Room {
	return predicate.Room(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Room {
	return predicate.Room(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Room {
	return predicate.Room(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Room {
	return predicate.Room(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Room {
	return predicate.Room(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Room {
	return predicate.Room(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Room {
	return predicate.Room(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Room {
	return predicate.Room(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Room {
	return predicate.Room(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Room {
	return predicate.Room(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Room {
	return predicate.Room(sql.FieldContainsFold(FieldName, v))
}

// SidEQ applies the EQ predicate on the "sid" field.
func SidEQ(v string) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldSid, v))
}

// SidNEQ applies the NEQ predicate on the "sid" field.
func SidNEQ(v string) predicate.Room {
	return predicate.Room(sql.FieldNEQ(FieldSid, v))
}

// SidIn applies the In predicate on the "sid" field.
func SidIn(vs ...string) predicate.Room {
	return predicate.Room(sql.FieldIn(FieldSid, vs...))
}

// SidNotIn applies the NotIn predicate on the "sid" field.
func SidNotIn(vs ...string) predicate.Room {
	return predicate.Room(sql.FieldNotIn(FieldSid, vs...))
}

// SidGT applies the GT predicate on the "sid" field.
func SidGT(v string) predicate.Room {
	return predicate.Room(sql.FieldGT(FieldSid, v))
}

// SidGTE applies the GTE predicate on the "sid" field.
func SidGTE(v string) predicate.Room {
	return predicate.Room(sql.FieldGTE(FieldSid, v))
}

// SidLT applies the LT predicate on the "sid" field.
func SidLT(v string) predicate.Room {
	return predicate.Room(sql.FieldLT(FieldSid, v))
}

// SidLTE applies the LTE predicate on the "sid" field.
func SidLTE(v string) predicate.Room {
	return predicate.Room(sql.FieldLTE(FieldSid, v))
}

// SidContains applies the Contains predicate on the "sid" field.
func SidContains(v string) predicate.Room {
	return predicate.Room(sql.FieldContains(FieldSid, v))
}

// SidHasPrefix applies the HasPrefix predicate on the "sid" field.
func SidHasPrefix(v string) predicate.Room {
	return predicate.Room(sql.FieldHasPrefix(FieldSid, v))
}

// SidHasSuffix applies the HasSuffix predicate on the "sid" field.
func SidHasSuffix(v string) predicate.Room {
	return predicate.Room(sql.FieldHasSuffix(FieldSid, v))
}

// SidEqualFold applies the EqualFold predicate on the "sid" field.
func SidEqualFold(v string) predicate.Room {
	return predicate.Room(sql.FieldEqualFold(FieldSid, v))
}

// SidContainsFold applies the ContainsFold predicate on the "sid" field.
func SidContainsFold(v string) predicate.Room {
	return predicate.Room(sql.FieldContainsFold(FieldSid, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v uint8) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v uint8) predicate.Room {
	return predicate.Room(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...uint8) predicate.Room {
	return predicate.Room(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...uint8) predicate.Room {
	return predicate.Room(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v uint8) predicate.Room {
	return predicate.Room(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v uint8) predicate.Room {
	return predicate.Room(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v uint8) predicate.Room {
	return predicate.Room(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v uint8) predicate.Room {
	return predicate.Room(sql.FieldLTE(FieldStatus, v))
}

// StartTimeEQ applies the EQ predicate on the "start_time" field.
func StartTimeEQ(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldStartTime, v))
}

// StartTimeNEQ applies the NEQ predicate on the "start_time" field.
func StartTimeNEQ(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldNEQ(FieldStartTime, v))
}

// StartTimeIn applies the In predicate on the "start_time" field.
func StartTimeIn(vs ...time.Time) predicate.Room {
	return predicate.Room(sql.FieldIn(FieldStartTime, vs...))
}

// StartTimeNotIn applies the NotIn predicate on the "start_time" field.
func StartTimeNotIn(vs ...time.Time) predicate.Room {
	return predicate.Room(sql.FieldNotIn(FieldStartTime, vs...))
}

// StartTimeGT applies the GT predicate on the "start_time" field.
func StartTimeGT(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldGT(FieldStartTime, v))
}

// StartTimeGTE applies the GTE predicate on the "start_time" field.
func StartTimeGTE(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldGTE(FieldStartTime, v))
}

// StartTimeLT applies the LT predicate on the "start_time" field.
func StartTimeLT(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldLT(FieldStartTime, v))
}

// StartTimeLTE applies the LTE predicate on the "start_time" field.
func StartTimeLTE(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldLTE(FieldStartTime, v))
}

// StartTimeIsNil applies the IsNil predicate on the "start_time" field.
func StartTimeIsNil() predicate.Room {
	return predicate.Room(sql.FieldIsNull(FieldStartTime))
}

// StartTimeNotNil applies the NotNil predicate on the "start_time" field.
func StartTimeNotNil() predicate.Room {
	return predicate.Room(sql.FieldNotNull(FieldStartTime))
}

// LeftTimeEQ applies the EQ predicate on the "left_time" field.
func LeftTimeEQ(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldLeftTime, v))
}

// LeftTimeNEQ applies the NEQ predicate on the "left_time" field.
func LeftTimeNEQ(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldNEQ(FieldLeftTime, v))
}

// LeftTimeIn applies the In predicate on the "left_time" field.
func LeftTimeIn(vs ...time.Time) predicate.Room {
	return predicate.Room(sql.FieldIn(FieldLeftTime, vs...))
}

// LeftTimeNotIn applies the NotIn predicate on the "left_time" field.
func LeftTimeNotIn(vs ...time.Time) predicate.Room {
	return predicate.Room(sql.FieldNotIn(FieldLeftTime, vs...))
}

// LeftTimeGT applies the GT predicate on the "left_time" field.
func LeftTimeGT(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldGT(FieldLeftTime, v))
}

// LeftTimeGTE applies the GTE predicate on the "left_time" field.
func LeftTimeGTE(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldGTE(FieldLeftTime, v))
}

// LeftTimeLT applies the LT predicate on the "left_time" field.
func LeftTimeLT(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldLT(FieldLeftTime, v))
}

// LeftTimeLTE applies the LTE predicate on the "left_time" field.
func LeftTimeLTE(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldLTE(FieldLeftTime, v))
}

// LeftTimeIsNil applies the IsNil predicate on the "left_time" field.
func LeftTimeIsNil() predicate.Room {
	return predicate.Room(sql.FieldIsNull(FieldLeftTime))
}

// LeftTimeNotNil applies the NotNil predicate on the "left_time" field.
func LeftTimeNotNil() predicate.Room {
	return predicate.Room(sql.FieldNotNull(FieldLeftTime))
}

// EndTimeEQ applies the EQ predicate on the "end_time" field.
func EndTimeEQ(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldEndTime, v))
}

// EndTimeNEQ applies the NEQ predicate on the "end_time" field.
func EndTimeNEQ(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldNEQ(FieldEndTime, v))
}

// EndTimeIn applies the In predicate on the "end_time" field.
func EndTimeIn(vs ...time.Time) predicate.Room {
	return predicate.Room(sql.FieldIn(FieldEndTime, vs...))
}

// EndTimeNotIn applies the NotIn predicate on the "end_time" field.
func EndTimeNotIn(vs ...time.Time) predicate.Room {
	return predicate.Room(sql.FieldNotIn(FieldEndTime, vs...))
}

// EndTimeGT applies the GT predicate on the "end_time" field.
func EndTimeGT(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldGT(FieldEndTime, v))
}

// EndTimeGTE applies the GTE predicate on the "end_time" field.
func EndTimeGTE(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldGTE(FieldEndTime, v))
}

// EndTimeLT applies the LT predicate on the "end_time" field.
func EndTimeLT(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldLT(FieldEndTime, v))
}

// EndTimeLTE applies the LTE predicate on the "end_time" field.
func EndTimeLTE(v time.Time) predicate.Room {
	return predicate.Room(sql.FieldLTE(FieldEndTime, v))
}

// EndTimeIsNil applies the IsNil predicate on the "end_time" field.
func EndTimeIsNil() predicate.Room {
	return predicate.Room(sql.FieldIsNull(FieldEndTime))
}

// EndTimeNotNil applies the NotNil predicate on the "end_time" field.
func EndTimeNotNil() predicate.Room {
	return predicate.Room(sql.FieldNotNull(FieldEndTime))
}

// VodStatusEQ applies the EQ predicate on the "vod_status" field.
func VodStatusEQ(v uint8) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldVodStatus, v))
}

// VodStatusNEQ applies the NEQ predicate on the "vod_status" field.
func VodStatusNEQ(v uint8) predicate.Room {
	return predicate.Room(sql.FieldNEQ(FieldVodStatus, v))
}

// VodStatusIn applies the In predicate on the "vod_status" field.
func VodStatusIn(vs ...uint8) predicate.Room {
	return predicate.Room(sql.FieldIn(FieldVodStatus, vs...))
}

// VodStatusNotIn applies the NotIn predicate on the "vod_status" field.
func VodStatusNotIn(vs ...uint8) predicate.Room {
	return predicate.Room(sql.FieldNotIn(FieldVodStatus, vs...))
}

// VodStatusGT applies the GT predicate on the "vod_status" field.
func VodStatusGT(v uint8) predicate.Room {
	return predicate.Room(sql.FieldGT(FieldVodStatus, v))
}

// VodStatusGTE applies the GTE predicate on the "vod_status" field.
func VodStatusGTE(v uint8) predicate.Room {
	return predicate.Room(sql.FieldGTE(FieldVodStatus, v))
}

// VodStatusLT applies the LT predicate on the "vod_status" field.
func VodStatusLT(v uint8) predicate.Room {
	return predicate.Room(sql.FieldLT(FieldVodStatus, v))
}

// VodStatusLTE applies the LTE predicate on the "vod_status" field.
func VodStatusLTE(v uint8) predicate.Room {
	return predicate.Room(sql.FieldLTE(FieldVodStatus, v))
}

// VodStatusIsNil applies the IsNil predicate on the "vod_status" field.
func VodStatusIsNil() predicate.Room {
	return predicate.Room(sql.FieldIsNull(FieldVodStatus))
}

// VodStatusNotNil applies the NotNil predicate on the "vod_status" field.
func VodStatusNotNil() predicate.Room {
	return predicate.Room(sql.FieldNotNull(FieldVodStatus))
}

// MetadataIsNil applies the IsNil predicate on the "metadata" field.
func MetadataIsNil() predicate.Room {
	return predicate.Room(sql.FieldIsNull(FieldMetadata))
}

// MetadataNotNil applies the NotNil predicate on the "metadata" field.
func MetadataNotNil() predicate.Room {
	return predicate.Room(sql.FieldNotNull(FieldMetadata))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Room) predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Room) predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Room) predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		p(s.Not())
	})
}
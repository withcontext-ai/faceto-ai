// Code generated by ent, DO NOT EDIT.

package roommessage

import (
	"faceto-ai/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldLTE(FieldID, id))
}

// UUID applies equality check predicate on the "uuid" field. It's identical to UUIDEQ.
func UUID(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldEQ(FieldUUID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldEQ(FieldDeletedAt, v))
}

// Sid applies equality check predicate on the "sid" field. It's identical to SidEQ.
func Sid(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldEQ(FieldSid, v))
}

// ParticipantSid applies equality check predicate on the "participant_sid" field. It's identical to ParticipantSidEQ.
func ParticipantSid(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldEQ(FieldParticipantSid, v))
}

// ParticipantName applies equality check predicate on the "participant_name" field. It's identical to ParticipantNameEQ.
func ParticipantName(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldEQ(FieldParticipantName, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v uint32) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldEQ(FieldType, v))
}

// IsReply applies equality check predicate on the "is_reply" field. It's identical to IsReplyEQ.
func IsReply(v uint32) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldEQ(FieldIsReply, v))
}

// ReplyID applies equality check predicate on the "reply_id" field. It's identical to ReplyIDEQ.
func ReplyID(v uint64) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldEQ(FieldReplyID, v))
}

// EventTime applies equality check predicate on the "event_time" field. It's identical to EventTimeEQ.
func EventTime(v time.Time) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldEQ(FieldEventTime, v))
}

// Text applies equality check predicate on the "text" field. It's identical to TextEQ.
func Text(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldEQ(FieldText, v))
}

// UUIDEQ applies the EQ predicate on the "uuid" field.
func UUIDEQ(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldEQ(FieldUUID, v))
}

// UUIDNEQ applies the NEQ predicate on the "uuid" field.
func UUIDNEQ(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldNEQ(FieldUUID, v))
}

// UUIDIn applies the In predicate on the "uuid" field.
func UUIDIn(vs ...string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldIn(FieldUUID, vs...))
}

// UUIDNotIn applies the NotIn predicate on the "uuid" field.
func UUIDNotIn(vs ...string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldNotIn(FieldUUID, vs...))
}

// UUIDGT applies the GT predicate on the "uuid" field.
func UUIDGT(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldGT(FieldUUID, v))
}

// UUIDGTE applies the GTE predicate on the "uuid" field.
func UUIDGTE(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldGTE(FieldUUID, v))
}

// UUIDLT applies the LT predicate on the "uuid" field.
func UUIDLT(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldLT(FieldUUID, v))
}

// UUIDLTE applies the LTE predicate on the "uuid" field.
func UUIDLTE(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldLTE(FieldUUID, v))
}

// UUIDContains applies the Contains predicate on the "uuid" field.
func UUIDContains(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldContains(FieldUUID, v))
}

// UUIDHasPrefix applies the HasPrefix predicate on the "uuid" field.
func UUIDHasPrefix(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldHasPrefix(FieldUUID, v))
}

// UUIDHasSuffix applies the HasSuffix predicate on the "uuid" field.
func UUIDHasSuffix(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldHasSuffix(FieldUUID, v))
}

// UUIDEqualFold applies the EqualFold predicate on the "uuid" field.
func UUIDEqualFold(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldEqualFold(FieldUUID, v))
}

// UUIDContainsFold applies the ContainsFold predicate on the "uuid" field.
func UUIDContainsFold(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldContainsFold(FieldUUID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldNotNull(FieldDeletedAt))
}

// SidEQ applies the EQ predicate on the "sid" field.
func SidEQ(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldEQ(FieldSid, v))
}

// SidNEQ applies the NEQ predicate on the "sid" field.
func SidNEQ(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldNEQ(FieldSid, v))
}

// SidIn applies the In predicate on the "sid" field.
func SidIn(vs ...string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldIn(FieldSid, vs...))
}

// SidNotIn applies the NotIn predicate on the "sid" field.
func SidNotIn(vs ...string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldNotIn(FieldSid, vs...))
}

// SidGT applies the GT predicate on the "sid" field.
func SidGT(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldGT(FieldSid, v))
}

// SidGTE applies the GTE predicate on the "sid" field.
func SidGTE(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldGTE(FieldSid, v))
}

// SidLT applies the LT predicate on the "sid" field.
func SidLT(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldLT(FieldSid, v))
}

// SidLTE applies the LTE predicate on the "sid" field.
func SidLTE(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldLTE(FieldSid, v))
}

// SidContains applies the Contains predicate on the "sid" field.
func SidContains(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldContains(FieldSid, v))
}

// SidHasPrefix applies the HasPrefix predicate on the "sid" field.
func SidHasPrefix(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldHasPrefix(FieldSid, v))
}

// SidHasSuffix applies the HasSuffix predicate on the "sid" field.
func SidHasSuffix(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldHasSuffix(FieldSid, v))
}

// SidEqualFold applies the EqualFold predicate on the "sid" field.
func SidEqualFold(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldEqualFold(FieldSid, v))
}

// SidContainsFold applies the ContainsFold predicate on the "sid" field.
func SidContainsFold(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldContainsFold(FieldSid, v))
}

// ParticipantSidEQ applies the EQ predicate on the "participant_sid" field.
func ParticipantSidEQ(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldEQ(FieldParticipantSid, v))
}

// ParticipantSidNEQ applies the NEQ predicate on the "participant_sid" field.
func ParticipantSidNEQ(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldNEQ(FieldParticipantSid, v))
}

// ParticipantSidIn applies the In predicate on the "participant_sid" field.
func ParticipantSidIn(vs ...string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldIn(FieldParticipantSid, vs...))
}

// ParticipantSidNotIn applies the NotIn predicate on the "participant_sid" field.
func ParticipantSidNotIn(vs ...string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldNotIn(FieldParticipantSid, vs...))
}

// ParticipantSidGT applies the GT predicate on the "participant_sid" field.
func ParticipantSidGT(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldGT(FieldParticipantSid, v))
}

// ParticipantSidGTE applies the GTE predicate on the "participant_sid" field.
func ParticipantSidGTE(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldGTE(FieldParticipantSid, v))
}

// ParticipantSidLT applies the LT predicate on the "participant_sid" field.
func ParticipantSidLT(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldLT(FieldParticipantSid, v))
}

// ParticipantSidLTE applies the LTE predicate on the "participant_sid" field.
func ParticipantSidLTE(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldLTE(FieldParticipantSid, v))
}

// ParticipantSidContains applies the Contains predicate on the "participant_sid" field.
func ParticipantSidContains(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldContains(FieldParticipantSid, v))
}

// ParticipantSidHasPrefix applies the HasPrefix predicate on the "participant_sid" field.
func ParticipantSidHasPrefix(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldHasPrefix(FieldParticipantSid, v))
}

// ParticipantSidHasSuffix applies the HasSuffix predicate on the "participant_sid" field.
func ParticipantSidHasSuffix(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldHasSuffix(FieldParticipantSid, v))
}

// ParticipantSidEqualFold applies the EqualFold predicate on the "participant_sid" field.
func ParticipantSidEqualFold(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldEqualFold(FieldParticipantSid, v))
}

// ParticipantSidContainsFold applies the ContainsFold predicate on the "participant_sid" field.
func ParticipantSidContainsFold(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldContainsFold(FieldParticipantSid, v))
}

// ParticipantNameEQ applies the EQ predicate on the "participant_name" field.
func ParticipantNameEQ(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldEQ(FieldParticipantName, v))
}

// ParticipantNameNEQ applies the NEQ predicate on the "participant_name" field.
func ParticipantNameNEQ(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldNEQ(FieldParticipantName, v))
}

// ParticipantNameIn applies the In predicate on the "participant_name" field.
func ParticipantNameIn(vs ...string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldIn(FieldParticipantName, vs...))
}

// ParticipantNameNotIn applies the NotIn predicate on the "participant_name" field.
func ParticipantNameNotIn(vs ...string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldNotIn(FieldParticipantName, vs...))
}

// ParticipantNameGT applies the GT predicate on the "participant_name" field.
func ParticipantNameGT(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldGT(FieldParticipantName, v))
}

// ParticipantNameGTE applies the GTE predicate on the "participant_name" field.
func ParticipantNameGTE(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldGTE(FieldParticipantName, v))
}

// ParticipantNameLT applies the LT predicate on the "participant_name" field.
func ParticipantNameLT(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldLT(FieldParticipantName, v))
}

// ParticipantNameLTE applies the LTE predicate on the "participant_name" field.
func ParticipantNameLTE(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldLTE(FieldParticipantName, v))
}

// ParticipantNameContains applies the Contains predicate on the "participant_name" field.
func ParticipantNameContains(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldContains(FieldParticipantName, v))
}

// ParticipantNameHasPrefix applies the HasPrefix predicate on the "participant_name" field.
func ParticipantNameHasPrefix(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldHasPrefix(FieldParticipantName, v))
}

// ParticipantNameHasSuffix applies the HasSuffix predicate on the "participant_name" field.
func ParticipantNameHasSuffix(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldHasSuffix(FieldParticipantName, v))
}

// ParticipantNameEqualFold applies the EqualFold predicate on the "participant_name" field.
func ParticipantNameEqualFold(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldEqualFold(FieldParticipantName, v))
}

// ParticipantNameContainsFold applies the ContainsFold predicate on the "participant_name" field.
func ParticipantNameContainsFold(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldContainsFold(FieldParticipantName, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v uint32) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v uint32) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...uint32) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...uint32) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v uint32) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v uint32) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v uint32) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v uint32) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldLTE(FieldType, v))
}

// TypeIsNil applies the IsNil predicate on the "type" field.
func TypeIsNil() predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldIsNull(FieldType))
}

// TypeNotNil applies the NotNil predicate on the "type" field.
func TypeNotNil() predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldNotNull(FieldType))
}

// IsReplyEQ applies the EQ predicate on the "is_reply" field.
func IsReplyEQ(v uint32) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldEQ(FieldIsReply, v))
}

// IsReplyNEQ applies the NEQ predicate on the "is_reply" field.
func IsReplyNEQ(v uint32) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldNEQ(FieldIsReply, v))
}

// IsReplyIn applies the In predicate on the "is_reply" field.
func IsReplyIn(vs ...uint32) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldIn(FieldIsReply, vs...))
}

// IsReplyNotIn applies the NotIn predicate on the "is_reply" field.
func IsReplyNotIn(vs ...uint32) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldNotIn(FieldIsReply, vs...))
}

// IsReplyGT applies the GT predicate on the "is_reply" field.
func IsReplyGT(v uint32) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldGT(FieldIsReply, v))
}

// IsReplyGTE applies the GTE predicate on the "is_reply" field.
func IsReplyGTE(v uint32) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldGTE(FieldIsReply, v))
}

// IsReplyLT applies the LT predicate on the "is_reply" field.
func IsReplyLT(v uint32) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldLT(FieldIsReply, v))
}

// IsReplyLTE applies the LTE predicate on the "is_reply" field.
func IsReplyLTE(v uint32) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldLTE(FieldIsReply, v))
}

// ReplyIDEQ applies the EQ predicate on the "reply_id" field.
func ReplyIDEQ(v uint64) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldEQ(FieldReplyID, v))
}

// ReplyIDNEQ applies the NEQ predicate on the "reply_id" field.
func ReplyIDNEQ(v uint64) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldNEQ(FieldReplyID, v))
}

// ReplyIDIn applies the In predicate on the "reply_id" field.
func ReplyIDIn(vs ...uint64) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldIn(FieldReplyID, vs...))
}

// ReplyIDNotIn applies the NotIn predicate on the "reply_id" field.
func ReplyIDNotIn(vs ...uint64) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldNotIn(FieldReplyID, vs...))
}

// ReplyIDGT applies the GT predicate on the "reply_id" field.
func ReplyIDGT(v uint64) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldGT(FieldReplyID, v))
}

// ReplyIDGTE applies the GTE predicate on the "reply_id" field.
func ReplyIDGTE(v uint64) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldGTE(FieldReplyID, v))
}

// ReplyIDLT applies the LT predicate on the "reply_id" field.
func ReplyIDLT(v uint64) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldLT(FieldReplyID, v))
}

// ReplyIDLTE applies the LTE predicate on the "reply_id" field.
func ReplyIDLTE(v uint64) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldLTE(FieldReplyID, v))
}

// ReplyIDIsNil applies the IsNil predicate on the "reply_id" field.
func ReplyIDIsNil() predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldIsNull(FieldReplyID))
}

// ReplyIDNotNil applies the NotNil predicate on the "reply_id" field.
func ReplyIDNotNil() predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldNotNull(FieldReplyID))
}

// EventTimeEQ applies the EQ predicate on the "event_time" field.
func EventTimeEQ(v time.Time) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldEQ(FieldEventTime, v))
}

// EventTimeNEQ applies the NEQ predicate on the "event_time" field.
func EventTimeNEQ(v time.Time) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldNEQ(FieldEventTime, v))
}

// EventTimeIn applies the In predicate on the "event_time" field.
func EventTimeIn(vs ...time.Time) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldIn(FieldEventTime, vs...))
}

// EventTimeNotIn applies the NotIn predicate on the "event_time" field.
func EventTimeNotIn(vs ...time.Time) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldNotIn(FieldEventTime, vs...))
}

// EventTimeGT applies the GT predicate on the "event_time" field.
func EventTimeGT(v time.Time) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldGT(FieldEventTime, v))
}

// EventTimeGTE applies the GTE predicate on the "event_time" field.
func EventTimeGTE(v time.Time) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldGTE(FieldEventTime, v))
}

// EventTimeLT applies the LT predicate on the "event_time" field.
func EventTimeLT(v time.Time) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldLT(FieldEventTime, v))
}

// EventTimeLTE applies the LTE predicate on the "event_time" field.
func EventTimeLTE(v time.Time) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldLTE(FieldEventTime, v))
}

// EventTimeIsNil applies the IsNil predicate on the "event_time" field.
func EventTimeIsNil() predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldIsNull(FieldEventTime))
}

// EventTimeNotNil applies the NotNil predicate on the "event_time" field.
func EventTimeNotNil() predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldNotNull(FieldEventTime))
}

// TextEQ applies the EQ predicate on the "text" field.
func TextEQ(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldEQ(FieldText, v))
}

// TextNEQ applies the NEQ predicate on the "text" field.
func TextNEQ(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldNEQ(FieldText, v))
}

// TextIn applies the In predicate on the "text" field.
func TextIn(vs ...string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldIn(FieldText, vs...))
}

// TextNotIn applies the NotIn predicate on the "text" field.
func TextNotIn(vs ...string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldNotIn(FieldText, vs...))
}

// TextGT applies the GT predicate on the "text" field.
func TextGT(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldGT(FieldText, v))
}

// TextGTE applies the GTE predicate on the "text" field.
func TextGTE(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldGTE(FieldText, v))
}

// TextLT applies the LT predicate on the "text" field.
func TextLT(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldLT(FieldText, v))
}

// TextLTE applies the LTE predicate on the "text" field.
func TextLTE(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldLTE(FieldText, v))
}

// TextContains applies the Contains predicate on the "text" field.
func TextContains(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldContains(FieldText, v))
}

// TextHasPrefix applies the HasPrefix predicate on the "text" field.
func TextHasPrefix(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldHasPrefix(FieldText, v))
}

// TextHasSuffix applies the HasSuffix predicate on the "text" field.
func TextHasSuffix(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldHasSuffix(FieldText, v))
}

// TextEqualFold applies the EqualFold predicate on the "text" field.
func TextEqualFold(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldEqualFold(FieldText, v))
}

// TextContainsFold applies the ContainsFold predicate on the "text" field.
func TextContainsFold(v string) predicate.RoomMessage {
	return predicate.RoomMessage(sql.FieldContainsFold(FieldText, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.RoomMessage) predicate.RoomMessage {
	return predicate.RoomMessage(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.RoomMessage) predicate.RoomMessage {
	return predicate.RoomMessage(func(s *sql.Selector) {
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
func Not(p predicate.RoomMessage) predicate.RoomMessage {
	return predicate.RoomMessage(func(s *sql.Selector) {
		p(s.Not())
	})
}

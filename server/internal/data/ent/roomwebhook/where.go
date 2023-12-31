// Code generated by ent, DO NOT EDIT.

package roomwebhook

import (
	"faceto-ai/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldLTE(FieldID, id))
}

// UUID applies equality check predicate on the "uuid" field. It's identical to UUIDEQ.
func UUID(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldEQ(FieldUUID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldEQ(FieldDeletedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldEQ(FieldName, v))
}

// Sid applies equality check predicate on the "sid" field. It's identical to SidEQ.
func Sid(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldEQ(FieldSid, v))
}

// Event applies equality check predicate on the "event" field. It's identical to EventEQ.
func Event(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldEQ(FieldEvent, v))
}

// EventTime applies equality check predicate on the "event_time" field. It's identical to EventTimeEQ.
func EventTime(v time.Time) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldEQ(FieldEventTime, v))
}

// UUIDEQ applies the EQ predicate on the "uuid" field.
func UUIDEQ(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldEQ(FieldUUID, v))
}

// UUIDNEQ applies the NEQ predicate on the "uuid" field.
func UUIDNEQ(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldNEQ(FieldUUID, v))
}

// UUIDIn applies the In predicate on the "uuid" field.
func UUIDIn(vs ...string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldIn(FieldUUID, vs...))
}

// UUIDNotIn applies the NotIn predicate on the "uuid" field.
func UUIDNotIn(vs ...string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldNotIn(FieldUUID, vs...))
}

// UUIDGT applies the GT predicate on the "uuid" field.
func UUIDGT(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldGT(FieldUUID, v))
}

// UUIDGTE applies the GTE predicate on the "uuid" field.
func UUIDGTE(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldGTE(FieldUUID, v))
}

// UUIDLT applies the LT predicate on the "uuid" field.
func UUIDLT(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldLT(FieldUUID, v))
}

// UUIDLTE applies the LTE predicate on the "uuid" field.
func UUIDLTE(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldLTE(FieldUUID, v))
}

// UUIDContains applies the Contains predicate on the "uuid" field.
func UUIDContains(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldContains(FieldUUID, v))
}

// UUIDHasPrefix applies the HasPrefix predicate on the "uuid" field.
func UUIDHasPrefix(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldHasPrefix(FieldUUID, v))
}

// UUIDHasSuffix applies the HasSuffix predicate on the "uuid" field.
func UUIDHasSuffix(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldHasSuffix(FieldUUID, v))
}

// UUIDEqualFold applies the EqualFold predicate on the "uuid" field.
func UUIDEqualFold(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldEqualFold(FieldUUID, v))
}

// UUIDContainsFold applies the ContainsFold predicate on the "uuid" field.
func UUIDContainsFold(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldContainsFold(FieldUUID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldNotNull(FieldDeletedAt))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldContainsFold(FieldName, v))
}

// SidEQ applies the EQ predicate on the "sid" field.
func SidEQ(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldEQ(FieldSid, v))
}

// SidNEQ applies the NEQ predicate on the "sid" field.
func SidNEQ(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldNEQ(FieldSid, v))
}

// SidIn applies the In predicate on the "sid" field.
func SidIn(vs ...string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldIn(FieldSid, vs...))
}

// SidNotIn applies the NotIn predicate on the "sid" field.
func SidNotIn(vs ...string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldNotIn(FieldSid, vs...))
}

// SidGT applies the GT predicate on the "sid" field.
func SidGT(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldGT(FieldSid, v))
}

// SidGTE applies the GTE predicate on the "sid" field.
func SidGTE(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldGTE(FieldSid, v))
}

// SidLT applies the LT predicate on the "sid" field.
func SidLT(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldLT(FieldSid, v))
}

// SidLTE applies the LTE predicate on the "sid" field.
func SidLTE(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldLTE(FieldSid, v))
}

// SidContains applies the Contains predicate on the "sid" field.
func SidContains(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldContains(FieldSid, v))
}

// SidHasPrefix applies the HasPrefix predicate on the "sid" field.
func SidHasPrefix(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldHasPrefix(FieldSid, v))
}

// SidHasSuffix applies the HasSuffix predicate on the "sid" field.
func SidHasSuffix(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldHasSuffix(FieldSid, v))
}

// SidEqualFold applies the EqualFold predicate on the "sid" field.
func SidEqualFold(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldEqualFold(FieldSid, v))
}

// SidContainsFold applies the ContainsFold predicate on the "sid" field.
func SidContainsFold(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldContainsFold(FieldSid, v))
}

// EventEQ applies the EQ predicate on the "event" field.
func EventEQ(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldEQ(FieldEvent, v))
}

// EventNEQ applies the NEQ predicate on the "event" field.
func EventNEQ(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldNEQ(FieldEvent, v))
}

// EventIn applies the In predicate on the "event" field.
func EventIn(vs ...string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldIn(FieldEvent, vs...))
}

// EventNotIn applies the NotIn predicate on the "event" field.
func EventNotIn(vs ...string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldNotIn(FieldEvent, vs...))
}

// EventGT applies the GT predicate on the "event" field.
func EventGT(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldGT(FieldEvent, v))
}

// EventGTE applies the GTE predicate on the "event" field.
func EventGTE(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldGTE(FieldEvent, v))
}

// EventLT applies the LT predicate on the "event" field.
func EventLT(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldLT(FieldEvent, v))
}

// EventLTE applies the LTE predicate on the "event" field.
func EventLTE(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldLTE(FieldEvent, v))
}

// EventContains applies the Contains predicate on the "event" field.
func EventContains(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldContains(FieldEvent, v))
}

// EventHasPrefix applies the HasPrefix predicate on the "event" field.
func EventHasPrefix(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldHasPrefix(FieldEvent, v))
}

// EventHasSuffix applies the HasSuffix predicate on the "event" field.
func EventHasSuffix(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldHasSuffix(FieldEvent, v))
}

// EventEqualFold applies the EqualFold predicate on the "event" field.
func EventEqualFold(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldEqualFold(FieldEvent, v))
}

// EventContainsFold applies the ContainsFold predicate on the "event" field.
func EventContainsFold(v string) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldContainsFold(FieldEvent, v))
}

// EventTimeEQ applies the EQ predicate on the "event_time" field.
func EventTimeEQ(v time.Time) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldEQ(FieldEventTime, v))
}

// EventTimeNEQ applies the NEQ predicate on the "event_time" field.
func EventTimeNEQ(v time.Time) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldNEQ(FieldEventTime, v))
}

// EventTimeIn applies the In predicate on the "event_time" field.
func EventTimeIn(vs ...time.Time) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldIn(FieldEventTime, vs...))
}

// EventTimeNotIn applies the NotIn predicate on the "event_time" field.
func EventTimeNotIn(vs ...time.Time) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldNotIn(FieldEventTime, vs...))
}

// EventTimeGT applies the GT predicate on the "event_time" field.
func EventTimeGT(v time.Time) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldGT(FieldEventTime, v))
}

// EventTimeGTE applies the GTE predicate on the "event_time" field.
func EventTimeGTE(v time.Time) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldGTE(FieldEventTime, v))
}

// EventTimeLT applies the LT predicate on the "event_time" field.
func EventTimeLT(v time.Time) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldLT(FieldEventTime, v))
}

// EventTimeLTE applies the LTE predicate on the "event_time" field.
func EventTimeLTE(v time.Time) predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldLTE(FieldEventTime, v))
}

// EventTimeIsNil applies the IsNil predicate on the "event_time" field.
func EventTimeIsNil() predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldIsNull(FieldEventTime))
}

// EventTimeNotNil applies the NotNil predicate on the "event_time" field.
func EventTimeNotNil() predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldNotNull(FieldEventTime))
}

// RoomIsNil applies the IsNil predicate on the "room" field.
func RoomIsNil() predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldIsNull(FieldRoom))
}

// RoomNotNil applies the NotNil predicate on the "room" field.
func RoomNotNil() predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldNotNull(FieldRoom))
}

// ParticipantIsNil applies the IsNil predicate on the "participant" field.
func ParticipantIsNil() predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldIsNull(FieldParticipant))
}

// ParticipantNotNil applies the NotNil predicate on the "participant" field.
func ParticipantNotNil() predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldNotNull(FieldParticipant))
}

// TrackIsNil applies the IsNil predicate on the "track" field.
func TrackIsNil() predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldIsNull(FieldTrack))
}

// TrackNotNil applies the NotNil predicate on the "track" field.
func TrackNotNil() predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldNotNull(FieldTrack))
}

// EgressInfoIsNil applies the IsNil predicate on the "egressInfo" field.
func EgressInfoIsNil() predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldIsNull(FieldEgressInfo))
}

// EgressInfoNotNil applies the NotNil predicate on the "egressInfo" field.
func EgressInfoNotNil() predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldNotNull(FieldEgressInfo))
}

// IngressInfoIsNil applies the IsNil predicate on the "ingressInfo" field.
func IngressInfoIsNil() predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldIsNull(FieldIngressInfo))
}

// IngressInfoNotNil applies the NotNil predicate on the "ingressInfo" field.
func IngressInfoNotNil() predicate.RoomWebhook {
	return predicate.RoomWebhook(sql.FieldNotNull(FieldIngressInfo))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.RoomWebhook) predicate.RoomWebhook {
	return predicate.RoomWebhook(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.RoomWebhook) predicate.RoomWebhook {
	return predicate.RoomWebhook(func(s *sql.Selector) {
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
func Not(p predicate.RoomWebhook) predicate.RoomWebhook {
	return predicate.RoomWebhook(func(s *sql.Selector) {
		p(s.Not())
	})
}

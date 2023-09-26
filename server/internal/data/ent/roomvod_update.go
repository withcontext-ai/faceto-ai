// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"faceto-ai/internal/data/ent/predicate"
	"faceto-ai/internal/data/ent/roomvod"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RoomVodUpdate is the builder for updating RoomVod entities.
type RoomVodUpdate struct {
	config
	hooks    []Hook
	mutation *RoomVodMutation
}

// Where appends a list predicates to the RoomVodUpdate builder.
func (rvu *RoomVodUpdate) Where(ps ...predicate.RoomVod) *RoomVodUpdate {
	rvu.mutation.Where(ps...)
	return rvu
}

// SetUUID sets the "uuid" field.
func (rvu *RoomVodUpdate) SetUUID(s string) *RoomVodUpdate {
	rvu.mutation.SetUUID(s)
	return rvu
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (rvu *RoomVodUpdate) SetNillableUUID(s *string) *RoomVodUpdate {
	if s != nil {
		rvu.SetUUID(*s)
	}
	return rvu
}

// SetCreatedAt sets the "created_at" field.
func (rvu *RoomVodUpdate) SetCreatedAt(t time.Time) *RoomVodUpdate {
	rvu.mutation.SetCreatedAt(t)
	return rvu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rvu *RoomVodUpdate) SetNillableCreatedAt(t *time.Time) *RoomVodUpdate {
	if t != nil {
		rvu.SetCreatedAt(*t)
	}
	return rvu
}

// SetUpdatedAt sets the "updated_at" field.
func (rvu *RoomVodUpdate) SetUpdatedAt(t time.Time) *RoomVodUpdate {
	rvu.mutation.SetUpdatedAt(t)
	return rvu
}

// SetDeletedAt sets the "deleted_at" field.
func (rvu *RoomVodUpdate) SetDeletedAt(t time.Time) *RoomVodUpdate {
	rvu.mutation.SetDeletedAt(t)
	return rvu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (rvu *RoomVodUpdate) SetNillableDeletedAt(t *time.Time) *RoomVodUpdate {
	if t != nil {
		rvu.SetDeletedAt(*t)
	}
	return rvu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (rvu *RoomVodUpdate) ClearDeletedAt() *RoomVodUpdate {
	rvu.mutation.ClearDeletedAt()
	return rvu
}

// SetName sets the "name" field.
func (rvu *RoomVodUpdate) SetName(s string) *RoomVodUpdate {
	rvu.mutation.SetName(s)
	return rvu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (rvu *RoomVodUpdate) SetNillableName(s *string) *RoomVodUpdate {
	if s != nil {
		rvu.SetName(*s)
	}
	return rvu
}

// SetSid sets the "sid" field.
func (rvu *RoomVodUpdate) SetSid(s string) *RoomVodUpdate {
	rvu.mutation.SetSid(s)
	return rvu
}

// SetNillableSid sets the "sid" field if the given value is not nil.
func (rvu *RoomVodUpdate) SetNillableSid(s *string) *RoomVodUpdate {
	if s != nil {
		rvu.SetSid(*s)
	}
	return rvu
}

// SetEgressID sets the "egress_id" field.
func (rvu *RoomVodUpdate) SetEgressID(s string) *RoomVodUpdate {
	rvu.mutation.SetEgressID(s)
	return rvu
}

// SetNillableEgressID sets the "egress_id" field if the given value is not nil.
func (rvu *RoomVodUpdate) SetNillableEgressID(s *string) *RoomVodUpdate {
	if s != nil {
		rvu.SetEgressID(*s)
	}
	return rvu
}

// SetStatus sets the "status" field.
func (rvu *RoomVodUpdate) SetStatus(u uint8) *RoomVodUpdate {
	rvu.mutation.ResetStatus()
	rvu.mutation.SetStatus(u)
	return rvu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (rvu *RoomVodUpdate) SetNillableStatus(u *uint8) *RoomVodUpdate {
	if u != nil {
		rvu.SetStatus(*u)
	}
	return rvu
}

// AddStatus adds u to the "status" field.
func (rvu *RoomVodUpdate) AddStatus(u int8) *RoomVodUpdate {
	rvu.mutation.AddStatus(u)
	return rvu
}

// ClearStatus clears the value of the "status" field.
func (rvu *RoomVodUpdate) ClearStatus() *RoomVodUpdate {
	rvu.mutation.ClearStatus()
	return rvu
}

// SetPlatform sets the "platform" field.
func (rvu *RoomVodUpdate) SetPlatform(u uint8) *RoomVodUpdate {
	rvu.mutation.ResetPlatform()
	rvu.mutation.SetPlatform(u)
	return rvu
}

// SetNillablePlatform sets the "platform" field if the given value is not nil.
func (rvu *RoomVodUpdate) SetNillablePlatform(u *uint8) *RoomVodUpdate {
	if u != nil {
		rvu.SetPlatform(*u)
	}
	return rvu
}

// AddPlatform adds u to the "platform" field.
func (rvu *RoomVodUpdate) AddPlatform(u int8) *RoomVodUpdate {
	rvu.mutation.AddPlatform(u)
	return rvu
}

// ClearPlatform clears the value of the "platform" field.
func (rvu *RoomVodUpdate) ClearPlatform() *RoomVodUpdate {
	rvu.mutation.ClearPlatform()
	return rvu
}

// SetVodType sets the "vod_type" field.
func (rvu *RoomVodUpdate) SetVodType(u uint8) *RoomVodUpdate {
	rvu.mutation.ResetVodType()
	rvu.mutation.SetVodType(u)
	return rvu
}

// SetNillableVodType sets the "vod_type" field if the given value is not nil.
func (rvu *RoomVodUpdate) SetNillableVodType(u *uint8) *RoomVodUpdate {
	if u != nil {
		rvu.SetVodType(*u)
	}
	return rvu
}

// AddVodType adds u to the "vod_type" field.
func (rvu *RoomVodUpdate) AddVodType(u int8) *RoomVodUpdate {
	rvu.mutation.AddVodType(u)
	return rvu
}

// ClearVodType clears the value of the "vod_type" field.
func (rvu *RoomVodUpdate) ClearVodType() *RoomVodUpdate {
	rvu.mutation.ClearVodType()
	return rvu
}

// SetVodPath sets the "vod_path" field.
func (rvu *RoomVodUpdate) SetVodPath(s string) *RoomVodUpdate {
	rvu.mutation.SetVodPath(s)
	return rvu
}

// SetNillableVodPath sets the "vod_path" field if the given value is not nil.
func (rvu *RoomVodUpdate) SetNillableVodPath(s *string) *RoomVodUpdate {
	if s != nil {
		rvu.SetVodPath(*s)
	}
	return rvu
}

// ClearVodPath clears the value of the "vod_path" field.
func (rvu *RoomVodUpdate) ClearVodPath() *RoomVodUpdate {
	rvu.mutation.ClearVodPath()
	return rvu
}

// SetVodURL sets the "vod_url" field.
func (rvu *RoomVodUpdate) SetVodURL(s string) *RoomVodUpdate {
	rvu.mutation.SetVodURL(s)
	return rvu
}

// SetNillableVodURL sets the "vod_url" field if the given value is not nil.
func (rvu *RoomVodUpdate) SetNillableVodURL(s *string) *RoomVodUpdate {
	if s != nil {
		rvu.SetVodURL(*s)
	}
	return rvu
}

// ClearVodURL clears the value of the "vod_url" field.
func (rvu *RoomVodUpdate) ClearVodURL() *RoomVodUpdate {
	rvu.mutation.ClearVodURL()
	return rvu
}

// SetStartTime sets the "start_time" field.
func (rvu *RoomVodUpdate) SetStartTime(t time.Time) *RoomVodUpdate {
	rvu.mutation.SetStartTime(t)
	return rvu
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (rvu *RoomVodUpdate) SetNillableStartTime(t *time.Time) *RoomVodUpdate {
	if t != nil {
		rvu.SetStartTime(*t)
	}
	return rvu
}

// ClearStartTime clears the value of the "start_time" field.
func (rvu *RoomVodUpdate) ClearStartTime() *RoomVodUpdate {
	rvu.mutation.ClearStartTime()
	return rvu
}

// SetCompleteTime sets the "complete_time" field.
func (rvu *RoomVodUpdate) SetCompleteTime(t time.Time) *RoomVodUpdate {
	rvu.mutation.SetCompleteTime(t)
	return rvu
}

// SetNillableCompleteTime sets the "complete_time" field if the given value is not nil.
func (rvu *RoomVodUpdate) SetNillableCompleteTime(t *time.Time) *RoomVodUpdate {
	if t != nil {
		rvu.SetCompleteTime(*t)
	}
	return rvu
}

// ClearCompleteTime clears the value of the "complete_time" field.
func (rvu *RoomVodUpdate) ClearCompleteTime() *RoomVodUpdate {
	rvu.mutation.ClearCompleteTime()
	return rvu
}

// SetDuration sets the "duration" field.
func (rvu *RoomVodUpdate) SetDuration(u uint64) *RoomVodUpdate {
	rvu.mutation.ResetDuration()
	rvu.mutation.SetDuration(u)
	return rvu
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (rvu *RoomVodUpdate) SetNillableDuration(u *uint64) *RoomVodUpdate {
	if u != nil {
		rvu.SetDuration(*u)
	}
	return rvu
}

// AddDuration adds u to the "duration" field.
func (rvu *RoomVodUpdate) AddDuration(u int64) *RoomVodUpdate {
	rvu.mutation.AddDuration(u)
	return rvu
}

// ClearDuration clears the value of the "duration" field.
func (rvu *RoomVodUpdate) ClearDuration() *RoomVodUpdate {
	rvu.mutation.ClearDuration()
	return rvu
}

// Mutation returns the RoomVodMutation object of the builder.
func (rvu *RoomVodUpdate) Mutation() *RoomVodMutation {
	return rvu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rvu *RoomVodUpdate) Save(ctx context.Context) (int, error) {
	rvu.defaults()
	return withHooks(ctx, rvu.sqlSave, rvu.mutation, rvu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rvu *RoomVodUpdate) SaveX(ctx context.Context) int {
	affected, err := rvu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rvu *RoomVodUpdate) Exec(ctx context.Context) error {
	_, err := rvu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rvu *RoomVodUpdate) ExecX(ctx context.Context) {
	if err := rvu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rvu *RoomVodUpdate) defaults() {
	if _, ok := rvu.mutation.UpdatedAt(); !ok {
		v := roomvod.UpdateDefaultUpdatedAt()
		rvu.mutation.SetUpdatedAt(v)
	}
}

func (rvu *RoomVodUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(roomvod.Table, roomvod.Columns, sqlgraph.NewFieldSpec(roomvod.FieldID, field.TypeUint64))
	if ps := rvu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rvu.mutation.UUID(); ok {
		_spec.SetField(roomvod.FieldUUID, field.TypeString, value)
	}
	if value, ok := rvu.mutation.CreatedAt(); ok {
		_spec.SetField(roomvod.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := rvu.mutation.UpdatedAt(); ok {
		_spec.SetField(roomvod.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := rvu.mutation.DeletedAt(); ok {
		_spec.SetField(roomvod.FieldDeletedAt, field.TypeTime, value)
	}
	if rvu.mutation.DeletedAtCleared() {
		_spec.ClearField(roomvod.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := rvu.mutation.Name(); ok {
		_spec.SetField(roomvod.FieldName, field.TypeString, value)
	}
	if value, ok := rvu.mutation.Sid(); ok {
		_spec.SetField(roomvod.FieldSid, field.TypeString, value)
	}
	if value, ok := rvu.mutation.EgressID(); ok {
		_spec.SetField(roomvod.FieldEgressID, field.TypeString, value)
	}
	if value, ok := rvu.mutation.Status(); ok {
		_spec.SetField(roomvod.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := rvu.mutation.AddedStatus(); ok {
		_spec.AddField(roomvod.FieldStatus, field.TypeUint8, value)
	}
	if rvu.mutation.StatusCleared() {
		_spec.ClearField(roomvod.FieldStatus, field.TypeUint8)
	}
	if value, ok := rvu.mutation.Platform(); ok {
		_spec.SetField(roomvod.FieldPlatform, field.TypeUint8, value)
	}
	if value, ok := rvu.mutation.AddedPlatform(); ok {
		_spec.AddField(roomvod.FieldPlatform, field.TypeUint8, value)
	}
	if rvu.mutation.PlatformCleared() {
		_spec.ClearField(roomvod.FieldPlatform, field.TypeUint8)
	}
	if value, ok := rvu.mutation.VodType(); ok {
		_spec.SetField(roomvod.FieldVodType, field.TypeUint8, value)
	}
	if value, ok := rvu.mutation.AddedVodType(); ok {
		_spec.AddField(roomvod.FieldVodType, field.TypeUint8, value)
	}
	if rvu.mutation.VodTypeCleared() {
		_spec.ClearField(roomvod.FieldVodType, field.TypeUint8)
	}
	if value, ok := rvu.mutation.VodPath(); ok {
		_spec.SetField(roomvod.FieldVodPath, field.TypeString, value)
	}
	if rvu.mutation.VodPathCleared() {
		_spec.ClearField(roomvod.FieldVodPath, field.TypeString)
	}
	if value, ok := rvu.mutation.VodURL(); ok {
		_spec.SetField(roomvod.FieldVodURL, field.TypeString, value)
	}
	if rvu.mutation.VodURLCleared() {
		_spec.ClearField(roomvod.FieldVodURL, field.TypeString)
	}
	if value, ok := rvu.mutation.StartTime(); ok {
		_spec.SetField(roomvod.FieldStartTime, field.TypeTime, value)
	}
	if rvu.mutation.StartTimeCleared() {
		_spec.ClearField(roomvod.FieldStartTime, field.TypeTime)
	}
	if value, ok := rvu.mutation.CompleteTime(); ok {
		_spec.SetField(roomvod.FieldCompleteTime, field.TypeTime, value)
	}
	if rvu.mutation.CompleteTimeCleared() {
		_spec.ClearField(roomvod.FieldCompleteTime, field.TypeTime)
	}
	if value, ok := rvu.mutation.Duration(); ok {
		_spec.SetField(roomvod.FieldDuration, field.TypeUint64, value)
	}
	if value, ok := rvu.mutation.AddedDuration(); ok {
		_spec.AddField(roomvod.FieldDuration, field.TypeUint64, value)
	}
	if rvu.mutation.DurationCleared() {
		_spec.ClearField(roomvod.FieldDuration, field.TypeUint64)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rvu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{roomvod.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rvu.mutation.done = true
	return n, nil
}

// RoomVodUpdateOne is the builder for updating a single RoomVod entity.
type RoomVodUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RoomVodMutation
}

// SetUUID sets the "uuid" field.
func (rvuo *RoomVodUpdateOne) SetUUID(s string) *RoomVodUpdateOne {
	rvuo.mutation.SetUUID(s)
	return rvuo
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (rvuo *RoomVodUpdateOne) SetNillableUUID(s *string) *RoomVodUpdateOne {
	if s != nil {
		rvuo.SetUUID(*s)
	}
	return rvuo
}

// SetCreatedAt sets the "created_at" field.
func (rvuo *RoomVodUpdateOne) SetCreatedAt(t time.Time) *RoomVodUpdateOne {
	rvuo.mutation.SetCreatedAt(t)
	return rvuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rvuo *RoomVodUpdateOne) SetNillableCreatedAt(t *time.Time) *RoomVodUpdateOne {
	if t != nil {
		rvuo.SetCreatedAt(*t)
	}
	return rvuo
}

// SetUpdatedAt sets the "updated_at" field.
func (rvuo *RoomVodUpdateOne) SetUpdatedAt(t time.Time) *RoomVodUpdateOne {
	rvuo.mutation.SetUpdatedAt(t)
	return rvuo
}

// SetDeletedAt sets the "deleted_at" field.
func (rvuo *RoomVodUpdateOne) SetDeletedAt(t time.Time) *RoomVodUpdateOne {
	rvuo.mutation.SetDeletedAt(t)
	return rvuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (rvuo *RoomVodUpdateOne) SetNillableDeletedAt(t *time.Time) *RoomVodUpdateOne {
	if t != nil {
		rvuo.SetDeletedAt(*t)
	}
	return rvuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (rvuo *RoomVodUpdateOne) ClearDeletedAt() *RoomVodUpdateOne {
	rvuo.mutation.ClearDeletedAt()
	return rvuo
}

// SetName sets the "name" field.
func (rvuo *RoomVodUpdateOne) SetName(s string) *RoomVodUpdateOne {
	rvuo.mutation.SetName(s)
	return rvuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (rvuo *RoomVodUpdateOne) SetNillableName(s *string) *RoomVodUpdateOne {
	if s != nil {
		rvuo.SetName(*s)
	}
	return rvuo
}

// SetSid sets the "sid" field.
func (rvuo *RoomVodUpdateOne) SetSid(s string) *RoomVodUpdateOne {
	rvuo.mutation.SetSid(s)
	return rvuo
}

// SetNillableSid sets the "sid" field if the given value is not nil.
func (rvuo *RoomVodUpdateOne) SetNillableSid(s *string) *RoomVodUpdateOne {
	if s != nil {
		rvuo.SetSid(*s)
	}
	return rvuo
}

// SetEgressID sets the "egress_id" field.
func (rvuo *RoomVodUpdateOne) SetEgressID(s string) *RoomVodUpdateOne {
	rvuo.mutation.SetEgressID(s)
	return rvuo
}

// SetNillableEgressID sets the "egress_id" field if the given value is not nil.
func (rvuo *RoomVodUpdateOne) SetNillableEgressID(s *string) *RoomVodUpdateOne {
	if s != nil {
		rvuo.SetEgressID(*s)
	}
	return rvuo
}

// SetStatus sets the "status" field.
func (rvuo *RoomVodUpdateOne) SetStatus(u uint8) *RoomVodUpdateOne {
	rvuo.mutation.ResetStatus()
	rvuo.mutation.SetStatus(u)
	return rvuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (rvuo *RoomVodUpdateOne) SetNillableStatus(u *uint8) *RoomVodUpdateOne {
	if u != nil {
		rvuo.SetStatus(*u)
	}
	return rvuo
}

// AddStatus adds u to the "status" field.
func (rvuo *RoomVodUpdateOne) AddStatus(u int8) *RoomVodUpdateOne {
	rvuo.mutation.AddStatus(u)
	return rvuo
}

// ClearStatus clears the value of the "status" field.
func (rvuo *RoomVodUpdateOne) ClearStatus() *RoomVodUpdateOne {
	rvuo.mutation.ClearStatus()
	return rvuo
}

// SetPlatform sets the "platform" field.
func (rvuo *RoomVodUpdateOne) SetPlatform(u uint8) *RoomVodUpdateOne {
	rvuo.mutation.ResetPlatform()
	rvuo.mutation.SetPlatform(u)
	return rvuo
}

// SetNillablePlatform sets the "platform" field if the given value is not nil.
func (rvuo *RoomVodUpdateOne) SetNillablePlatform(u *uint8) *RoomVodUpdateOne {
	if u != nil {
		rvuo.SetPlatform(*u)
	}
	return rvuo
}

// AddPlatform adds u to the "platform" field.
func (rvuo *RoomVodUpdateOne) AddPlatform(u int8) *RoomVodUpdateOne {
	rvuo.mutation.AddPlatform(u)
	return rvuo
}

// ClearPlatform clears the value of the "platform" field.
func (rvuo *RoomVodUpdateOne) ClearPlatform() *RoomVodUpdateOne {
	rvuo.mutation.ClearPlatform()
	return rvuo
}

// SetVodType sets the "vod_type" field.
func (rvuo *RoomVodUpdateOne) SetVodType(u uint8) *RoomVodUpdateOne {
	rvuo.mutation.ResetVodType()
	rvuo.mutation.SetVodType(u)
	return rvuo
}

// SetNillableVodType sets the "vod_type" field if the given value is not nil.
func (rvuo *RoomVodUpdateOne) SetNillableVodType(u *uint8) *RoomVodUpdateOne {
	if u != nil {
		rvuo.SetVodType(*u)
	}
	return rvuo
}

// AddVodType adds u to the "vod_type" field.
func (rvuo *RoomVodUpdateOne) AddVodType(u int8) *RoomVodUpdateOne {
	rvuo.mutation.AddVodType(u)
	return rvuo
}

// ClearVodType clears the value of the "vod_type" field.
func (rvuo *RoomVodUpdateOne) ClearVodType() *RoomVodUpdateOne {
	rvuo.mutation.ClearVodType()
	return rvuo
}

// SetVodPath sets the "vod_path" field.
func (rvuo *RoomVodUpdateOne) SetVodPath(s string) *RoomVodUpdateOne {
	rvuo.mutation.SetVodPath(s)
	return rvuo
}

// SetNillableVodPath sets the "vod_path" field if the given value is not nil.
func (rvuo *RoomVodUpdateOne) SetNillableVodPath(s *string) *RoomVodUpdateOne {
	if s != nil {
		rvuo.SetVodPath(*s)
	}
	return rvuo
}

// ClearVodPath clears the value of the "vod_path" field.
func (rvuo *RoomVodUpdateOne) ClearVodPath() *RoomVodUpdateOne {
	rvuo.mutation.ClearVodPath()
	return rvuo
}

// SetVodURL sets the "vod_url" field.
func (rvuo *RoomVodUpdateOne) SetVodURL(s string) *RoomVodUpdateOne {
	rvuo.mutation.SetVodURL(s)
	return rvuo
}

// SetNillableVodURL sets the "vod_url" field if the given value is not nil.
func (rvuo *RoomVodUpdateOne) SetNillableVodURL(s *string) *RoomVodUpdateOne {
	if s != nil {
		rvuo.SetVodURL(*s)
	}
	return rvuo
}

// ClearVodURL clears the value of the "vod_url" field.
func (rvuo *RoomVodUpdateOne) ClearVodURL() *RoomVodUpdateOne {
	rvuo.mutation.ClearVodURL()
	return rvuo
}

// SetStartTime sets the "start_time" field.
func (rvuo *RoomVodUpdateOne) SetStartTime(t time.Time) *RoomVodUpdateOne {
	rvuo.mutation.SetStartTime(t)
	return rvuo
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (rvuo *RoomVodUpdateOne) SetNillableStartTime(t *time.Time) *RoomVodUpdateOne {
	if t != nil {
		rvuo.SetStartTime(*t)
	}
	return rvuo
}

// ClearStartTime clears the value of the "start_time" field.
func (rvuo *RoomVodUpdateOne) ClearStartTime() *RoomVodUpdateOne {
	rvuo.mutation.ClearStartTime()
	return rvuo
}

// SetCompleteTime sets the "complete_time" field.
func (rvuo *RoomVodUpdateOne) SetCompleteTime(t time.Time) *RoomVodUpdateOne {
	rvuo.mutation.SetCompleteTime(t)
	return rvuo
}

// SetNillableCompleteTime sets the "complete_time" field if the given value is not nil.
func (rvuo *RoomVodUpdateOne) SetNillableCompleteTime(t *time.Time) *RoomVodUpdateOne {
	if t != nil {
		rvuo.SetCompleteTime(*t)
	}
	return rvuo
}

// ClearCompleteTime clears the value of the "complete_time" field.
func (rvuo *RoomVodUpdateOne) ClearCompleteTime() *RoomVodUpdateOne {
	rvuo.mutation.ClearCompleteTime()
	return rvuo
}

// SetDuration sets the "duration" field.
func (rvuo *RoomVodUpdateOne) SetDuration(u uint64) *RoomVodUpdateOne {
	rvuo.mutation.ResetDuration()
	rvuo.mutation.SetDuration(u)
	return rvuo
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (rvuo *RoomVodUpdateOne) SetNillableDuration(u *uint64) *RoomVodUpdateOne {
	if u != nil {
		rvuo.SetDuration(*u)
	}
	return rvuo
}

// AddDuration adds u to the "duration" field.
func (rvuo *RoomVodUpdateOne) AddDuration(u int64) *RoomVodUpdateOne {
	rvuo.mutation.AddDuration(u)
	return rvuo
}

// ClearDuration clears the value of the "duration" field.
func (rvuo *RoomVodUpdateOne) ClearDuration() *RoomVodUpdateOne {
	rvuo.mutation.ClearDuration()
	return rvuo
}

// Mutation returns the RoomVodMutation object of the builder.
func (rvuo *RoomVodUpdateOne) Mutation() *RoomVodMutation {
	return rvuo.mutation
}

// Where appends a list predicates to the RoomVodUpdate builder.
func (rvuo *RoomVodUpdateOne) Where(ps ...predicate.RoomVod) *RoomVodUpdateOne {
	rvuo.mutation.Where(ps...)
	return rvuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rvuo *RoomVodUpdateOne) Select(field string, fields ...string) *RoomVodUpdateOne {
	rvuo.fields = append([]string{field}, fields...)
	return rvuo
}

// Save executes the query and returns the updated RoomVod entity.
func (rvuo *RoomVodUpdateOne) Save(ctx context.Context) (*RoomVod, error) {
	rvuo.defaults()
	return withHooks(ctx, rvuo.sqlSave, rvuo.mutation, rvuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rvuo *RoomVodUpdateOne) SaveX(ctx context.Context) *RoomVod {
	node, err := rvuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rvuo *RoomVodUpdateOne) Exec(ctx context.Context) error {
	_, err := rvuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rvuo *RoomVodUpdateOne) ExecX(ctx context.Context) {
	if err := rvuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rvuo *RoomVodUpdateOne) defaults() {
	if _, ok := rvuo.mutation.UpdatedAt(); !ok {
		v := roomvod.UpdateDefaultUpdatedAt()
		rvuo.mutation.SetUpdatedAt(v)
	}
}

func (rvuo *RoomVodUpdateOne) sqlSave(ctx context.Context) (_node *RoomVod, err error) {
	_spec := sqlgraph.NewUpdateSpec(roomvod.Table, roomvod.Columns, sqlgraph.NewFieldSpec(roomvod.FieldID, field.TypeUint64))
	id, ok := rvuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RoomVod.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rvuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, roomvod.FieldID)
		for _, f := range fields {
			if !roomvod.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != roomvod.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rvuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rvuo.mutation.UUID(); ok {
		_spec.SetField(roomvod.FieldUUID, field.TypeString, value)
	}
	if value, ok := rvuo.mutation.CreatedAt(); ok {
		_spec.SetField(roomvod.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := rvuo.mutation.UpdatedAt(); ok {
		_spec.SetField(roomvod.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := rvuo.mutation.DeletedAt(); ok {
		_spec.SetField(roomvod.FieldDeletedAt, field.TypeTime, value)
	}
	if rvuo.mutation.DeletedAtCleared() {
		_spec.ClearField(roomvod.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := rvuo.mutation.Name(); ok {
		_spec.SetField(roomvod.FieldName, field.TypeString, value)
	}
	if value, ok := rvuo.mutation.Sid(); ok {
		_spec.SetField(roomvod.FieldSid, field.TypeString, value)
	}
	if value, ok := rvuo.mutation.EgressID(); ok {
		_spec.SetField(roomvod.FieldEgressID, field.TypeString, value)
	}
	if value, ok := rvuo.mutation.Status(); ok {
		_spec.SetField(roomvod.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := rvuo.mutation.AddedStatus(); ok {
		_spec.AddField(roomvod.FieldStatus, field.TypeUint8, value)
	}
	if rvuo.mutation.StatusCleared() {
		_spec.ClearField(roomvod.FieldStatus, field.TypeUint8)
	}
	if value, ok := rvuo.mutation.Platform(); ok {
		_spec.SetField(roomvod.FieldPlatform, field.TypeUint8, value)
	}
	if value, ok := rvuo.mutation.AddedPlatform(); ok {
		_spec.AddField(roomvod.FieldPlatform, field.TypeUint8, value)
	}
	if rvuo.mutation.PlatformCleared() {
		_spec.ClearField(roomvod.FieldPlatform, field.TypeUint8)
	}
	if value, ok := rvuo.mutation.VodType(); ok {
		_spec.SetField(roomvod.FieldVodType, field.TypeUint8, value)
	}
	if value, ok := rvuo.mutation.AddedVodType(); ok {
		_spec.AddField(roomvod.FieldVodType, field.TypeUint8, value)
	}
	if rvuo.mutation.VodTypeCleared() {
		_spec.ClearField(roomvod.FieldVodType, field.TypeUint8)
	}
	if value, ok := rvuo.mutation.VodPath(); ok {
		_spec.SetField(roomvod.FieldVodPath, field.TypeString, value)
	}
	if rvuo.mutation.VodPathCleared() {
		_spec.ClearField(roomvod.FieldVodPath, field.TypeString)
	}
	if value, ok := rvuo.mutation.VodURL(); ok {
		_spec.SetField(roomvod.FieldVodURL, field.TypeString, value)
	}
	if rvuo.mutation.VodURLCleared() {
		_spec.ClearField(roomvod.FieldVodURL, field.TypeString)
	}
	if value, ok := rvuo.mutation.StartTime(); ok {
		_spec.SetField(roomvod.FieldStartTime, field.TypeTime, value)
	}
	if rvuo.mutation.StartTimeCleared() {
		_spec.ClearField(roomvod.FieldStartTime, field.TypeTime)
	}
	if value, ok := rvuo.mutation.CompleteTime(); ok {
		_spec.SetField(roomvod.FieldCompleteTime, field.TypeTime, value)
	}
	if rvuo.mutation.CompleteTimeCleared() {
		_spec.ClearField(roomvod.FieldCompleteTime, field.TypeTime)
	}
	if value, ok := rvuo.mutation.Duration(); ok {
		_spec.SetField(roomvod.FieldDuration, field.TypeUint64, value)
	}
	if value, ok := rvuo.mutation.AddedDuration(); ok {
		_spec.AddField(roomvod.FieldDuration, field.TypeUint64, value)
	}
	if rvuo.mutation.DurationCleared() {
		_spec.ClearField(roomvod.FieldDuration, field.TypeUint64)
	}
	_node = &RoomVod{config: rvuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rvuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{roomvod.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rvuo.mutation.done = true
	return _node, nil
}
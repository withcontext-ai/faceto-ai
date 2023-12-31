// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"faceto-ai/internal/data/ent/room"
	"faceto-ai/internal/data/schema"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RoomCreate is the builder for creating a Room entity.
type RoomCreate struct {
	config
	mutation *RoomMutation
	hooks    []Hook
}

// SetUUID sets the "uuid" field.
func (rc *RoomCreate) SetUUID(s string) *RoomCreate {
	rc.mutation.SetUUID(s)
	return rc
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (rc *RoomCreate) SetNillableUUID(s *string) *RoomCreate {
	if s != nil {
		rc.SetUUID(*s)
	}
	return rc
}

// SetCreatedAt sets the "created_at" field.
func (rc *RoomCreate) SetCreatedAt(t time.Time) *RoomCreate {
	rc.mutation.SetCreatedAt(t)
	return rc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rc *RoomCreate) SetNillableCreatedAt(t *time.Time) *RoomCreate {
	if t != nil {
		rc.SetCreatedAt(*t)
	}
	return rc
}

// SetUpdatedAt sets the "updated_at" field.
func (rc *RoomCreate) SetUpdatedAt(t time.Time) *RoomCreate {
	rc.mutation.SetUpdatedAt(t)
	return rc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rc *RoomCreate) SetNillableUpdatedAt(t *time.Time) *RoomCreate {
	if t != nil {
		rc.SetUpdatedAt(*t)
	}
	return rc
}

// SetDeletedAt sets the "deleted_at" field.
func (rc *RoomCreate) SetDeletedAt(t time.Time) *RoomCreate {
	rc.mutation.SetDeletedAt(t)
	return rc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (rc *RoomCreate) SetNillableDeletedAt(t *time.Time) *RoomCreate {
	if t != nil {
		rc.SetDeletedAt(*t)
	}
	return rc
}

// SetName sets the "name" field.
func (rc *RoomCreate) SetName(s string) *RoomCreate {
	rc.mutation.SetName(s)
	return rc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (rc *RoomCreate) SetNillableName(s *string) *RoomCreate {
	if s != nil {
		rc.SetName(*s)
	}
	return rc
}

// SetSid sets the "sid" field.
func (rc *RoomCreate) SetSid(s string) *RoomCreate {
	rc.mutation.SetSid(s)
	return rc
}

// SetNillableSid sets the "sid" field if the given value is not nil.
func (rc *RoomCreate) SetNillableSid(s *string) *RoomCreate {
	if s != nil {
		rc.SetSid(*s)
	}
	return rc
}

// SetStatus sets the "status" field.
func (rc *RoomCreate) SetStatus(u uint8) *RoomCreate {
	rc.mutation.SetStatus(u)
	return rc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (rc *RoomCreate) SetNillableStatus(u *uint8) *RoomCreate {
	if u != nil {
		rc.SetStatus(*u)
	}
	return rc
}

// SetStartTime sets the "start_time" field.
func (rc *RoomCreate) SetStartTime(t time.Time) *RoomCreate {
	rc.mutation.SetStartTime(t)
	return rc
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (rc *RoomCreate) SetNillableStartTime(t *time.Time) *RoomCreate {
	if t != nil {
		rc.SetStartTime(*t)
	}
	return rc
}

// SetLeftTime sets the "left_time" field.
func (rc *RoomCreate) SetLeftTime(t time.Time) *RoomCreate {
	rc.mutation.SetLeftTime(t)
	return rc
}

// SetNillableLeftTime sets the "left_time" field if the given value is not nil.
func (rc *RoomCreate) SetNillableLeftTime(t *time.Time) *RoomCreate {
	if t != nil {
		rc.SetLeftTime(*t)
	}
	return rc
}

// SetEndTime sets the "end_time" field.
func (rc *RoomCreate) SetEndTime(t time.Time) *RoomCreate {
	rc.mutation.SetEndTime(t)
	return rc
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (rc *RoomCreate) SetNillableEndTime(t *time.Time) *RoomCreate {
	if t != nil {
		rc.SetEndTime(*t)
	}
	return rc
}

// SetVodStatus sets the "vod_status" field.
func (rc *RoomCreate) SetVodStatus(u uint8) *RoomCreate {
	rc.mutation.SetVodStatus(u)
	return rc
}

// SetNillableVodStatus sets the "vod_status" field if the given value is not nil.
func (rc *RoomCreate) SetNillableVodStatus(u *uint8) *RoomCreate {
	if u != nil {
		rc.SetVodStatus(*u)
	}
	return rc
}

// SetMetadata sets the "metadata" field.
func (rc *RoomCreate) SetMetadata(s *schema.Metadata) *RoomCreate {
	rc.mutation.SetMetadata(s)
	return rc
}

// SetID sets the "id" field.
func (rc *RoomCreate) SetID(u uint64) *RoomCreate {
	rc.mutation.SetID(u)
	return rc
}

// Mutation returns the RoomMutation object of the builder.
func (rc *RoomCreate) Mutation() *RoomMutation {
	return rc.mutation
}

// Save creates the Room in the database.
func (rc *RoomCreate) Save(ctx context.Context) (*Room, error) {
	rc.defaults()
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RoomCreate) SaveX(ctx context.Context) *Room {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RoomCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RoomCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RoomCreate) defaults() {
	if _, ok := rc.mutation.UUID(); !ok {
		v := room.DefaultUUID()
		rc.mutation.SetUUID(v)
	}
	if _, ok := rc.mutation.CreatedAt(); !ok {
		v := room.DefaultCreatedAt()
		rc.mutation.SetCreatedAt(v)
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		v := room.DefaultUpdatedAt()
		rc.mutation.SetUpdatedAt(v)
	}
	if _, ok := rc.mutation.Name(); !ok {
		v := room.DefaultName
		rc.mutation.SetName(v)
	}
	if _, ok := rc.mutation.Sid(); !ok {
		v := room.DefaultSid
		rc.mutation.SetSid(v)
	}
	if _, ok := rc.mutation.Status(); !ok {
		v := room.DefaultStatus
		rc.mutation.SetStatus(v)
	}
	if _, ok := rc.mutation.VodStatus(); !ok {
		v := room.DefaultVodStatus
		rc.mutation.SetVodStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RoomCreate) check() error {
	if _, ok := rc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New(`ent: missing required field "Room.uuid"`)}
	}
	if _, ok := rc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Room.created_at"`)}
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Room.updated_at"`)}
	}
	if _, ok := rc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Room.name"`)}
	}
	if _, ok := rc.mutation.Sid(); !ok {
		return &ValidationError{Name: "sid", err: errors.New(`ent: missing required field "Room.sid"`)}
	}
	if _, ok := rc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Room.status"`)}
	}
	return nil
}

func (rc *RoomCreate) sqlSave(ctx context.Context) (*Room, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *RoomCreate) createSpec() (*Room, *sqlgraph.CreateSpec) {
	var (
		_node = &Room{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(room.Table, sqlgraph.NewFieldSpec(room.FieldID, field.TypeUint64))
	)
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rc.mutation.UUID(); ok {
		_spec.SetField(room.FieldUUID, field.TypeString, value)
		_node.UUID = value
	}
	if value, ok := rc.mutation.CreatedAt(); ok {
		_spec.SetField(room.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := rc.mutation.UpdatedAt(); ok {
		_spec.SetField(room.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := rc.mutation.DeletedAt(); ok {
		_spec.SetField(room.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := rc.mutation.Name(); ok {
		_spec.SetField(room.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := rc.mutation.Sid(); ok {
		_spec.SetField(room.FieldSid, field.TypeString, value)
		_node.Sid = value
	}
	if value, ok := rc.mutation.Status(); ok {
		_spec.SetField(room.FieldStatus, field.TypeUint8, value)
		_node.Status = value
	}
	if value, ok := rc.mutation.StartTime(); ok {
		_spec.SetField(room.FieldStartTime, field.TypeTime, value)
		_node.StartTime = value
	}
	if value, ok := rc.mutation.LeftTime(); ok {
		_spec.SetField(room.FieldLeftTime, field.TypeTime, value)
		_node.LeftTime = value
	}
	if value, ok := rc.mutation.EndTime(); ok {
		_spec.SetField(room.FieldEndTime, field.TypeTime, value)
		_node.EndTime = value
	}
	if value, ok := rc.mutation.VodStatus(); ok {
		_spec.SetField(room.FieldVodStatus, field.TypeUint8, value)
		_node.VodStatus = value
	}
	if value, ok := rc.mutation.Metadata(); ok {
		_spec.SetField(room.FieldMetadata, field.TypeJSON, value)
		_node.Metadata = value
	}
	return _node, _spec
}

// RoomCreateBulk is the builder for creating many Room entities in bulk.
type RoomCreateBulk struct {
	config
	builders []*RoomCreate
}

// Save creates the Room entities in the database.
func (rcb *RoomCreateBulk) Save(ctx context.Context) ([]*Room, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Room, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RoomMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RoomCreateBulk) SaveX(ctx context.Context) []*Room {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RoomCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RoomCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}

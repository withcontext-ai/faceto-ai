// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"faceto-ai/internal/data/ent/link"
	"faceto-ai/internal/data/schema"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LinkCreate is the builder for creating a Link entity.
type LinkCreate struct {
	config
	mutation *LinkMutation
	hooks    []Hook
}

// SetUUID sets the "uuid" field.
func (lc *LinkCreate) SetUUID(s string) *LinkCreate {
	lc.mutation.SetUUID(s)
	return lc
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (lc *LinkCreate) SetNillableUUID(s *string) *LinkCreate {
	if s != nil {
		lc.SetUUID(*s)
	}
	return lc
}

// SetCreatedAt sets the "created_at" field.
func (lc *LinkCreate) SetCreatedAt(t time.Time) *LinkCreate {
	lc.mutation.SetCreatedAt(t)
	return lc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lc *LinkCreate) SetNillableCreatedAt(t *time.Time) *LinkCreate {
	if t != nil {
		lc.SetCreatedAt(*t)
	}
	return lc
}

// SetUpdatedAt sets the "updated_at" field.
func (lc *LinkCreate) SetUpdatedAt(t time.Time) *LinkCreate {
	lc.mutation.SetUpdatedAt(t)
	return lc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (lc *LinkCreate) SetNillableUpdatedAt(t *time.Time) *LinkCreate {
	if t != nil {
		lc.SetUpdatedAt(*t)
	}
	return lc
}

// SetDeletedAt sets the "deleted_at" field.
func (lc *LinkCreate) SetDeletedAt(t time.Time) *LinkCreate {
	lc.mutation.SetDeletedAt(t)
	return lc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (lc *LinkCreate) SetNillableDeletedAt(t *time.Time) *LinkCreate {
	if t != nil {
		lc.SetDeletedAt(*t)
	}
	return lc
}

// SetRoomName sets the "room_name" field.
func (lc *LinkCreate) SetRoomName(s string) *LinkCreate {
	lc.mutation.SetRoomName(s)
	return lc
}

// SetNillableRoomName sets the "room_name" field if the given value is not nil.
func (lc *LinkCreate) SetNillableRoomName(s *string) *LinkCreate {
	if s != nil {
		lc.SetRoomName(*s)
	}
	return lc
}

// SetLink sets the "link" field.
func (lc *LinkCreate) SetLink(s string) *LinkCreate {
	lc.mutation.SetLink(s)
	return lc
}

// SetNillableLink sets the "link" field if the given value is not nil.
func (lc *LinkCreate) SetNillableLink(s *string) *LinkCreate {
	if s != nil {
		lc.SetLink(*s)
	}
	return lc
}

// SetChatAPI sets the "chat_api" field.
func (lc *LinkCreate) SetChatAPI(s string) *LinkCreate {
	lc.mutation.SetChatAPI(s)
	return lc
}

// SetNillableChatAPI sets the "chat_api" field if the given value is not nil.
func (lc *LinkCreate) SetNillableChatAPI(s *string) *LinkCreate {
	if s != nil {
		lc.SetChatAPI(*s)
	}
	return lc
}

// SetChatAPIKey sets the "chat_api_key" field.
func (lc *LinkCreate) SetChatAPIKey(s string) *LinkCreate {
	lc.mutation.SetChatAPIKey(s)
	return lc
}

// SetNillableChatAPIKey sets the "chat_api_key" field if the given value is not nil.
func (lc *LinkCreate) SetNillableChatAPIKey(s *string) *LinkCreate {
	if s != nil {
		lc.SetChatAPIKey(*s)
	}
	return lc
}

// SetToken sets the "token" field.
func (lc *LinkCreate) SetToken(s string) *LinkCreate {
	lc.mutation.SetToken(s)
	return lc
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (lc *LinkCreate) SetNillableToken(s *string) *LinkCreate {
	if s != nil {
		lc.SetToken(*s)
	}
	return lc
}

// SetConfig sets the "config" field.
func (lc *LinkCreate) SetConfig(sc *schema.RoomConfig) *LinkCreate {
	lc.mutation.SetConfig(sc)
	return lc
}

// SetWebhook sets the "webhook" field.
func (lc *LinkCreate) SetWebhook(s *schema.Webhook) *LinkCreate {
	lc.mutation.SetWebhook(s)
	return lc
}

// SetPrompt sets the "prompt" field.
func (lc *LinkCreate) SetPrompt(s *schema.Prompt) *LinkCreate {
	lc.mutation.SetPrompt(s)
	return lc
}

// SetID sets the "id" field.
func (lc *LinkCreate) SetID(u uint64) *LinkCreate {
	lc.mutation.SetID(u)
	return lc
}

// Mutation returns the LinkMutation object of the builder.
func (lc *LinkCreate) Mutation() *LinkMutation {
	return lc.mutation
}

// Save creates the Link in the database.
func (lc *LinkCreate) Save(ctx context.Context) (*Link, error) {
	lc.defaults()
	return withHooks(ctx, lc.sqlSave, lc.mutation, lc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LinkCreate) SaveX(ctx context.Context) *Link {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LinkCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LinkCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LinkCreate) defaults() {
	if _, ok := lc.mutation.UUID(); !ok {
		v := link.DefaultUUID()
		lc.mutation.SetUUID(v)
	}
	if _, ok := lc.mutation.CreatedAt(); !ok {
		v := link.DefaultCreatedAt()
		lc.mutation.SetCreatedAt(v)
	}
	if _, ok := lc.mutation.UpdatedAt(); !ok {
		v := link.DefaultUpdatedAt()
		lc.mutation.SetUpdatedAt(v)
	}
	if _, ok := lc.mutation.RoomName(); !ok {
		v := link.DefaultRoomName
		lc.mutation.SetRoomName(v)
	}
	if _, ok := lc.mutation.Link(); !ok {
		v := link.DefaultLink
		lc.mutation.SetLink(v)
	}
	if _, ok := lc.mutation.ChatAPI(); !ok {
		v := link.DefaultChatAPI
		lc.mutation.SetChatAPI(v)
	}
	if _, ok := lc.mutation.ChatAPIKey(); !ok {
		v := link.DefaultChatAPIKey
		lc.mutation.SetChatAPIKey(v)
	}
	if _, ok := lc.mutation.Token(); !ok {
		v := link.DefaultToken
		lc.mutation.SetToken(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LinkCreate) check() error {
	if _, ok := lc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New(`ent: missing required field "Link.uuid"`)}
	}
	if _, ok := lc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Link.created_at"`)}
	}
	if _, ok := lc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Link.updated_at"`)}
	}
	if _, ok := lc.mutation.RoomName(); !ok {
		return &ValidationError{Name: "room_name", err: errors.New(`ent: missing required field "Link.room_name"`)}
	}
	return nil
}

func (lc *LinkCreate) sqlSave(ctx context.Context) (*Link, error) {
	if err := lc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	lc.mutation.id = &_node.ID
	lc.mutation.done = true
	return _node, nil
}

func (lc *LinkCreate) createSpec() (*Link, *sqlgraph.CreateSpec) {
	var (
		_node = &Link{config: lc.config}
		_spec = sqlgraph.NewCreateSpec(link.Table, sqlgraph.NewFieldSpec(link.FieldID, field.TypeUint64))
	)
	if id, ok := lc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := lc.mutation.UUID(); ok {
		_spec.SetField(link.FieldUUID, field.TypeString, value)
		_node.UUID = value
	}
	if value, ok := lc.mutation.CreatedAt(); ok {
		_spec.SetField(link.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := lc.mutation.UpdatedAt(); ok {
		_spec.SetField(link.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := lc.mutation.DeletedAt(); ok {
		_spec.SetField(link.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := lc.mutation.RoomName(); ok {
		_spec.SetField(link.FieldRoomName, field.TypeString, value)
		_node.RoomName = value
	}
	if value, ok := lc.mutation.Link(); ok {
		_spec.SetField(link.FieldLink, field.TypeString, value)
		_node.Link = value
	}
	if value, ok := lc.mutation.ChatAPI(); ok {
		_spec.SetField(link.FieldChatAPI, field.TypeString, value)
		_node.ChatAPI = value
	}
	if value, ok := lc.mutation.ChatAPIKey(); ok {
		_spec.SetField(link.FieldChatAPIKey, field.TypeString, value)
		_node.ChatAPIKey = value
	}
	if value, ok := lc.mutation.Token(); ok {
		_spec.SetField(link.FieldToken, field.TypeString, value)
		_node.Token = value
	}
	if value, ok := lc.mutation.Config(); ok {
		_spec.SetField(link.FieldConfig, field.TypeJSON, value)
		_node.Config = value
	}
	if value, ok := lc.mutation.Webhook(); ok {
		_spec.SetField(link.FieldWebhook, field.TypeJSON, value)
		_node.Webhook = value
	}
	if value, ok := lc.mutation.Prompt(); ok {
		_spec.SetField(link.FieldPrompt, field.TypeJSON, value)
		_node.Prompt = value
	}
	return _node, _spec
}

// LinkCreateBulk is the builder for creating many Link entities in bulk.
type LinkCreateBulk struct {
	config
	builders []*LinkCreate
}

// Save creates the Link entities in the database.
func (lcb *LinkCreateBulk) Save(ctx context.Context) ([]*Link, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Link, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LinkMutation)
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
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LinkCreateBulk) SaveX(ctx context.Context) []*Link {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LinkCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LinkCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}

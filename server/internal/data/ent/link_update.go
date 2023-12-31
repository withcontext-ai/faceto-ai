// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"faceto-ai/internal/data/ent/link"
	"faceto-ai/internal/data/ent/predicate"
	"faceto-ai/internal/data/schema"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LinkUpdate is the builder for updating Link entities.
type LinkUpdate struct {
	config
	hooks    []Hook
	mutation *LinkMutation
}

// Where appends a list predicates to the LinkUpdate builder.
func (lu *LinkUpdate) Where(ps ...predicate.Link) *LinkUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetUUID sets the "uuid" field.
func (lu *LinkUpdate) SetUUID(s string) *LinkUpdate {
	lu.mutation.SetUUID(s)
	return lu
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (lu *LinkUpdate) SetNillableUUID(s *string) *LinkUpdate {
	if s != nil {
		lu.SetUUID(*s)
	}
	return lu
}

// SetCreatedAt sets the "created_at" field.
func (lu *LinkUpdate) SetCreatedAt(t time.Time) *LinkUpdate {
	lu.mutation.SetCreatedAt(t)
	return lu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lu *LinkUpdate) SetNillableCreatedAt(t *time.Time) *LinkUpdate {
	if t != nil {
		lu.SetCreatedAt(*t)
	}
	return lu
}

// SetUpdatedAt sets the "updated_at" field.
func (lu *LinkUpdate) SetUpdatedAt(t time.Time) *LinkUpdate {
	lu.mutation.SetUpdatedAt(t)
	return lu
}

// SetDeletedAt sets the "deleted_at" field.
func (lu *LinkUpdate) SetDeletedAt(t time.Time) *LinkUpdate {
	lu.mutation.SetDeletedAt(t)
	return lu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (lu *LinkUpdate) SetNillableDeletedAt(t *time.Time) *LinkUpdate {
	if t != nil {
		lu.SetDeletedAt(*t)
	}
	return lu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (lu *LinkUpdate) ClearDeletedAt() *LinkUpdate {
	lu.mutation.ClearDeletedAt()
	return lu
}

// SetRoomName sets the "room_name" field.
func (lu *LinkUpdate) SetRoomName(s string) *LinkUpdate {
	lu.mutation.SetRoomName(s)
	return lu
}

// SetNillableRoomName sets the "room_name" field if the given value is not nil.
func (lu *LinkUpdate) SetNillableRoomName(s *string) *LinkUpdate {
	if s != nil {
		lu.SetRoomName(*s)
	}
	return lu
}

// SetLink sets the "link" field.
func (lu *LinkUpdate) SetLink(s string) *LinkUpdate {
	lu.mutation.SetLink(s)
	return lu
}

// SetNillableLink sets the "link" field if the given value is not nil.
func (lu *LinkUpdate) SetNillableLink(s *string) *LinkUpdate {
	if s != nil {
		lu.SetLink(*s)
	}
	return lu
}

// ClearLink clears the value of the "link" field.
func (lu *LinkUpdate) ClearLink() *LinkUpdate {
	lu.mutation.ClearLink()
	return lu
}

// SetChatAPI sets the "chat_api" field.
func (lu *LinkUpdate) SetChatAPI(s string) *LinkUpdate {
	lu.mutation.SetChatAPI(s)
	return lu
}

// SetNillableChatAPI sets the "chat_api" field if the given value is not nil.
func (lu *LinkUpdate) SetNillableChatAPI(s *string) *LinkUpdate {
	if s != nil {
		lu.SetChatAPI(*s)
	}
	return lu
}

// ClearChatAPI clears the value of the "chat_api" field.
func (lu *LinkUpdate) ClearChatAPI() *LinkUpdate {
	lu.mutation.ClearChatAPI()
	return lu
}

// SetChatAPIKey sets the "chat_api_key" field.
func (lu *LinkUpdate) SetChatAPIKey(s string) *LinkUpdate {
	lu.mutation.SetChatAPIKey(s)
	return lu
}

// SetNillableChatAPIKey sets the "chat_api_key" field if the given value is not nil.
func (lu *LinkUpdate) SetNillableChatAPIKey(s *string) *LinkUpdate {
	if s != nil {
		lu.SetChatAPIKey(*s)
	}
	return lu
}

// ClearChatAPIKey clears the value of the "chat_api_key" field.
func (lu *LinkUpdate) ClearChatAPIKey() *LinkUpdate {
	lu.mutation.ClearChatAPIKey()
	return lu
}

// SetToken sets the "token" field.
func (lu *LinkUpdate) SetToken(s string) *LinkUpdate {
	lu.mutation.SetToken(s)
	return lu
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (lu *LinkUpdate) SetNillableToken(s *string) *LinkUpdate {
	if s != nil {
		lu.SetToken(*s)
	}
	return lu
}

// ClearToken clears the value of the "token" field.
func (lu *LinkUpdate) ClearToken() *LinkUpdate {
	lu.mutation.ClearToken()
	return lu
}

// SetConfig sets the "config" field.
func (lu *LinkUpdate) SetConfig(sc *schema.RoomConfig) *LinkUpdate {
	lu.mutation.SetConfig(sc)
	return lu
}

// ClearConfig clears the value of the "config" field.
func (lu *LinkUpdate) ClearConfig() *LinkUpdate {
	lu.mutation.ClearConfig()
	return lu
}

// SetWebhook sets the "webhook" field.
func (lu *LinkUpdate) SetWebhook(s *schema.Webhook) *LinkUpdate {
	lu.mutation.SetWebhook(s)
	return lu
}

// ClearWebhook clears the value of the "webhook" field.
func (lu *LinkUpdate) ClearWebhook() *LinkUpdate {
	lu.mutation.ClearWebhook()
	return lu
}

// SetPrompt sets the "prompt" field.
func (lu *LinkUpdate) SetPrompt(s *schema.Prompt) *LinkUpdate {
	lu.mutation.SetPrompt(s)
	return lu
}

// ClearPrompt clears the value of the "prompt" field.
func (lu *LinkUpdate) ClearPrompt() *LinkUpdate {
	lu.mutation.ClearPrompt()
	return lu
}

// Mutation returns the LinkMutation object of the builder.
func (lu *LinkUpdate) Mutation() *LinkMutation {
	return lu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LinkUpdate) Save(ctx context.Context) (int, error) {
	lu.defaults()
	return withHooks(ctx, lu.sqlSave, lu.mutation, lu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LinkUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LinkUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LinkUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lu *LinkUpdate) defaults() {
	if _, ok := lu.mutation.UpdatedAt(); !ok {
		v := link.UpdateDefaultUpdatedAt()
		lu.mutation.SetUpdatedAt(v)
	}
}

func (lu *LinkUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(link.Table, link.Columns, sqlgraph.NewFieldSpec(link.FieldID, field.TypeUint64))
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.UUID(); ok {
		_spec.SetField(link.FieldUUID, field.TypeString, value)
	}
	if value, ok := lu.mutation.CreatedAt(); ok {
		_spec.SetField(link.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := lu.mutation.UpdatedAt(); ok {
		_spec.SetField(link.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := lu.mutation.DeletedAt(); ok {
		_spec.SetField(link.FieldDeletedAt, field.TypeTime, value)
	}
	if lu.mutation.DeletedAtCleared() {
		_spec.ClearField(link.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := lu.mutation.RoomName(); ok {
		_spec.SetField(link.FieldRoomName, field.TypeString, value)
	}
	if value, ok := lu.mutation.Link(); ok {
		_spec.SetField(link.FieldLink, field.TypeString, value)
	}
	if lu.mutation.LinkCleared() {
		_spec.ClearField(link.FieldLink, field.TypeString)
	}
	if value, ok := lu.mutation.ChatAPI(); ok {
		_spec.SetField(link.FieldChatAPI, field.TypeString, value)
	}
	if lu.mutation.ChatAPICleared() {
		_spec.ClearField(link.FieldChatAPI, field.TypeString)
	}
	if value, ok := lu.mutation.ChatAPIKey(); ok {
		_spec.SetField(link.FieldChatAPIKey, field.TypeString, value)
	}
	if lu.mutation.ChatAPIKeyCleared() {
		_spec.ClearField(link.FieldChatAPIKey, field.TypeString)
	}
	if value, ok := lu.mutation.Token(); ok {
		_spec.SetField(link.FieldToken, field.TypeString, value)
	}
	if lu.mutation.TokenCleared() {
		_spec.ClearField(link.FieldToken, field.TypeString)
	}
	if value, ok := lu.mutation.Config(); ok {
		_spec.SetField(link.FieldConfig, field.TypeJSON, value)
	}
	if lu.mutation.ConfigCleared() {
		_spec.ClearField(link.FieldConfig, field.TypeJSON)
	}
	if value, ok := lu.mutation.Webhook(); ok {
		_spec.SetField(link.FieldWebhook, field.TypeJSON, value)
	}
	if lu.mutation.WebhookCleared() {
		_spec.ClearField(link.FieldWebhook, field.TypeJSON)
	}
	if value, ok := lu.mutation.Prompt(); ok {
		_spec.SetField(link.FieldPrompt, field.TypeJSON, value)
	}
	if lu.mutation.PromptCleared() {
		_spec.ClearField(link.FieldPrompt, field.TypeJSON)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{link.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lu.mutation.done = true
	return n, nil
}

// LinkUpdateOne is the builder for updating a single Link entity.
type LinkUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LinkMutation
}

// SetUUID sets the "uuid" field.
func (luo *LinkUpdateOne) SetUUID(s string) *LinkUpdateOne {
	luo.mutation.SetUUID(s)
	return luo
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (luo *LinkUpdateOne) SetNillableUUID(s *string) *LinkUpdateOne {
	if s != nil {
		luo.SetUUID(*s)
	}
	return luo
}

// SetCreatedAt sets the "created_at" field.
func (luo *LinkUpdateOne) SetCreatedAt(t time.Time) *LinkUpdateOne {
	luo.mutation.SetCreatedAt(t)
	return luo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (luo *LinkUpdateOne) SetNillableCreatedAt(t *time.Time) *LinkUpdateOne {
	if t != nil {
		luo.SetCreatedAt(*t)
	}
	return luo
}

// SetUpdatedAt sets the "updated_at" field.
func (luo *LinkUpdateOne) SetUpdatedAt(t time.Time) *LinkUpdateOne {
	luo.mutation.SetUpdatedAt(t)
	return luo
}

// SetDeletedAt sets the "deleted_at" field.
func (luo *LinkUpdateOne) SetDeletedAt(t time.Time) *LinkUpdateOne {
	luo.mutation.SetDeletedAt(t)
	return luo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (luo *LinkUpdateOne) SetNillableDeletedAt(t *time.Time) *LinkUpdateOne {
	if t != nil {
		luo.SetDeletedAt(*t)
	}
	return luo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (luo *LinkUpdateOne) ClearDeletedAt() *LinkUpdateOne {
	luo.mutation.ClearDeletedAt()
	return luo
}

// SetRoomName sets the "room_name" field.
func (luo *LinkUpdateOne) SetRoomName(s string) *LinkUpdateOne {
	luo.mutation.SetRoomName(s)
	return luo
}

// SetNillableRoomName sets the "room_name" field if the given value is not nil.
func (luo *LinkUpdateOne) SetNillableRoomName(s *string) *LinkUpdateOne {
	if s != nil {
		luo.SetRoomName(*s)
	}
	return luo
}

// SetLink sets the "link" field.
func (luo *LinkUpdateOne) SetLink(s string) *LinkUpdateOne {
	luo.mutation.SetLink(s)
	return luo
}

// SetNillableLink sets the "link" field if the given value is not nil.
func (luo *LinkUpdateOne) SetNillableLink(s *string) *LinkUpdateOne {
	if s != nil {
		luo.SetLink(*s)
	}
	return luo
}

// ClearLink clears the value of the "link" field.
func (luo *LinkUpdateOne) ClearLink() *LinkUpdateOne {
	luo.mutation.ClearLink()
	return luo
}

// SetChatAPI sets the "chat_api" field.
func (luo *LinkUpdateOne) SetChatAPI(s string) *LinkUpdateOne {
	luo.mutation.SetChatAPI(s)
	return luo
}

// SetNillableChatAPI sets the "chat_api" field if the given value is not nil.
func (luo *LinkUpdateOne) SetNillableChatAPI(s *string) *LinkUpdateOne {
	if s != nil {
		luo.SetChatAPI(*s)
	}
	return luo
}

// ClearChatAPI clears the value of the "chat_api" field.
func (luo *LinkUpdateOne) ClearChatAPI() *LinkUpdateOne {
	luo.mutation.ClearChatAPI()
	return luo
}

// SetChatAPIKey sets the "chat_api_key" field.
func (luo *LinkUpdateOne) SetChatAPIKey(s string) *LinkUpdateOne {
	luo.mutation.SetChatAPIKey(s)
	return luo
}

// SetNillableChatAPIKey sets the "chat_api_key" field if the given value is not nil.
func (luo *LinkUpdateOne) SetNillableChatAPIKey(s *string) *LinkUpdateOne {
	if s != nil {
		luo.SetChatAPIKey(*s)
	}
	return luo
}

// ClearChatAPIKey clears the value of the "chat_api_key" field.
func (luo *LinkUpdateOne) ClearChatAPIKey() *LinkUpdateOne {
	luo.mutation.ClearChatAPIKey()
	return luo
}

// SetToken sets the "token" field.
func (luo *LinkUpdateOne) SetToken(s string) *LinkUpdateOne {
	luo.mutation.SetToken(s)
	return luo
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (luo *LinkUpdateOne) SetNillableToken(s *string) *LinkUpdateOne {
	if s != nil {
		luo.SetToken(*s)
	}
	return luo
}

// ClearToken clears the value of the "token" field.
func (luo *LinkUpdateOne) ClearToken() *LinkUpdateOne {
	luo.mutation.ClearToken()
	return luo
}

// SetConfig sets the "config" field.
func (luo *LinkUpdateOne) SetConfig(sc *schema.RoomConfig) *LinkUpdateOne {
	luo.mutation.SetConfig(sc)
	return luo
}

// ClearConfig clears the value of the "config" field.
func (luo *LinkUpdateOne) ClearConfig() *LinkUpdateOne {
	luo.mutation.ClearConfig()
	return luo
}

// SetWebhook sets the "webhook" field.
func (luo *LinkUpdateOne) SetWebhook(s *schema.Webhook) *LinkUpdateOne {
	luo.mutation.SetWebhook(s)
	return luo
}

// ClearWebhook clears the value of the "webhook" field.
func (luo *LinkUpdateOne) ClearWebhook() *LinkUpdateOne {
	luo.mutation.ClearWebhook()
	return luo
}

// SetPrompt sets the "prompt" field.
func (luo *LinkUpdateOne) SetPrompt(s *schema.Prompt) *LinkUpdateOne {
	luo.mutation.SetPrompt(s)
	return luo
}

// ClearPrompt clears the value of the "prompt" field.
func (luo *LinkUpdateOne) ClearPrompt() *LinkUpdateOne {
	luo.mutation.ClearPrompt()
	return luo
}

// Mutation returns the LinkMutation object of the builder.
func (luo *LinkUpdateOne) Mutation() *LinkMutation {
	return luo.mutation
}

// Where appends a list predicates to the LinkUpdate builder.
func (luo *LinkUpdateOne) Where(ps ...predicate.Link) *LinkUpdateOne {
	luo.mutation.Where(ps...)
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LinkUpdateOne) Select(field string, fields ...string) *LinkUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Link entity.
func (luo *LinkUpdateOne) Save(ctx context.Context) (*Link, error) {
	luo.defaults()
	return withHooks(ctx, luo.sqlSave, luo.mutation, luo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LinkUpdateOne) SaveX(ctx context.Context) *Link {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LinkUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LinkUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (luo *LinkUpdateOne) defaults() {
	if _, ok := luo.mutation.UpdatedAt(); !ok {
		v := link.UpdateDefaultUpdatedAt()
		luo.mutation.SetUpdatedAt(v)
	}
}

func (luo *LinkUpdateOne) sqlSave(ctx context.Context) (_node *Link, err error) {
	_spec := sqlgraph.NewUpdateSpec(link.Table, link.Columns, sqlgraph.NewFieldSpec(link.FieldID, field.TypeUint64))
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Link.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, link.FieldID)
		for _, f := range fields {
			if !link.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != link.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.UUID(); ok {
		_spec.SetField(link.FieldUUID, field.TypeString, value)
	}
	if value, ok := luo.mutation.CreatedAt(); ok {
		_spec.SetField(link.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := luo.mutation.UpdatedAt(); ok {
		_spec.SetField(link.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := luo.mutation.DeletedAt(); ok {
		_spec.SetField(link.FieldDeletedAt, field.TypeTime, value)
	}
	if luo.mutation.DeletedAtCleared() {
		_spec.ClearField(link.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := luo.mutation.RoomName(); ok {
		_spec.SetField(link.FieldRoomName, field.TypeString, value)
	}
	if value, ok := luo.mutation.Link(); ok {
		_spec.SetField(link.FieldLink, field.TypeString, value)
	}
	if luo.mutation.LinkCleared() {
		_spec.ClearField(link.FieldLink, field.TypeString)
	}
	if value, ok := luo.mutation.ChatAPI(); ok {
		_spec.SetField(link.FieldChatAPI, field.TypeString, value)
	}
	if luo.mutation.ChatAPICleared() {
		_spec.ClearField(link.FieldChatAPI, field.TypeString)
	}
	if value, ok := luo.mutation.ChatAPIKey(); ok {
		_spec.SetField(link.FieldChatAPIKey, field.TypeString, value)
	}
	if luo.mutation.ChatAPIKeyCleared() {
		_spec.ClearField(link.FieldChatAPIKey, field.TypeString)
	}
	if value, ok := luo.mutation.Token(); ok {
		_spec.SetField(link.FieldToken, field.TypeString, value)
	}
	if luo.mutation.TokenCleared() {
		_spec.ClearField(link.FieldToken, field.TypeString)
	}
	if value, ok := luo.mutation.Config(); ok {
		_spec.SetField(link.FieldConfig, field.TypeJSON, value)
	}
	if luo.mutation.ConfigCleared() {
		_spec.ClearField(link.FieldConfig, field.TypeJSON)
	}
	if value, ok := luo.mutation.Webhook(); ok {
		_spec.SetField(link.FieldWebhook, field.TypeJSON, value)
	}
	if luo.mutation.WebhookCleared() {
		_spec.ClearField(link.FieldWebhook, field.TypeJSON)
	}
	if value, ok := luo.mutation.Prompt(); ok {
		_spec.SetField(link.FieldPrompt, field.TypeJSON, value)
	}
	if luo.mutation.PromptCleared() {
		_spec.ClearField(link.FieldPrompt, field.TypeJSON)
	}
	_node = &Link{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{link.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	luo.mutation.done = true
	return _node, nil
}

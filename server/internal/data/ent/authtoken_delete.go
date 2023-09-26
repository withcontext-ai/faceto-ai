// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"faceto-ai/internal/data/ent/authtoken"
	"faceto-ai/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AuthTokenDelete is the builder for deleting a AuthToken entity.
type AuthTokenDelete struct {
	config
	hooks    []Hook
	mutation *AuthTokenMutation
}

// Where appends a list predicates to the AuthTokenDelete builder.
func (atd *AuthTokenDelete) Where(ps ...predicate.AuthToken) *AuthTokenDelete {
	atd.mutation.Where(ps...)
	return atd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (atd *AuthTokenDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, atd.sqlExec, atd.mutation, atd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (atd *AuthTokenDelete) ExecX(ctx context.Context) int {
	n, err := atd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (atd *AuthTokenDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(authtoken.Table, sqlgraph.NewFieldSpec(authtoken.FieldID, field.TypeUint64))
	if ps := atd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, atd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	atd.mutation.done = true
	return affected, err
}

// AuthTokenDeleteOne is the builder for deleting a single AuthToken entity.
type AuthTokenDeleteOne struct {
	atd *AuthTokenDelete
}

// Where appends a list predicates to the AuthTokenDelete builder.
func (atdo *AuthTokenDeleteOne) Where(ps ...predicate.AuthToken) *AuthTokenDeleteOne {
	atdo.atd.mutation.Where(ps...)
	return atdo
}

// Exec executes the deletion query.
func (atdo *AuthTokenDeleteOne) Exec(ctx context.Context) error {
	n, err := atdo.atd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{authtoken.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (atdo *AuthTokenDeleteOne) ExecX(ctx context.Context) {
	if err := atdo.Exec(ctx); err != nil {
		panic(err)
	}
}
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"faceto-ai/internal/data/ent/predicate"
	"faceto-ai/internal/data/ent/roomwebhook"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RoomWebhookQuery is the builder for querying RoomWebhook entities.
type RoomWebhookQuery struct {
	config
	ctx        *QueryContext
	order      []roomwebhook.OrderOption
	inters     []Interceptor
	predicates []predicate.RoomWebhook
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RoomWebhookQuery builder.
func (rwq *RoomWebhookQuery) Where(ps ...predicate.RoomWebhook) *RoomWebhookQuery {
	rwq.predicates = append(rwq.predicates, ps...)
	return rwq
}

// Limit the number of records to be returned by this query.
func (rwq *RoomWebhookQuery) Limit(limit int) *RoomWebhookQuery {
	rwq.ctx.Limit = &limit
	return rwq
}

// Offset to start from.
func (rwq *RoomWebhookQuery) Offset(offset int) *RoomWebhookQuery {
	rwq.ctx.Offset = &offset
	return rwq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rwq *RoomWebhookQuery) Unique(unique bool) *RoomWebhookQuery {
	rwq.ctx.Unique = &unique
	return rwq
}

// Order specifies how the records should be ordered.
func (rwq *RoomWebhookQuery) Order(o ...roomwebhook.OrderOption) *RoomWebhookQuery {
	rwq.order = append(rwq.order, o...)
	return rwq
}

// First returns the first RoomWebhook entity from the query.
// Returns a *NotFoundError when no RoomWebhook was found.
func (rwq *RoomWebhookQuery) First(ctx context.Context) (*RoomWebhook, error) {
	nodes, err := rwq.Limit(1).All(setContextOp(ctx, rwq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{roomwebhook.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rwq *RoomWebhookQuery) FirstX(ctx context.Context) *RoomWebhook {
	node, err := rwq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RoomWebhook ID from the query.
// Returns a *NotFoundError when no RoomWebhook ID was found.
func (rwq *RoomWebhookQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = rwq.Limit(1).IDs(setContextOp(ctx, rwq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{roomwebhook.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rwq *RoomWebhookQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := rwq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RoomWebhook entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one RoomWebhook entity is found.
// Returns a *NotFoundError when no RoomWebhook entities are found.
func (rwq *RoomWebhookQuery) Only(ctx context.Context) (*RoomWebhook, error) {
	nodes, err := rwq.Limit(2).All(setContextOp(ctx, rwq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{roomwebhook.Label}
	default:
		return nil, &NotSingularError{roomwebhook.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rwq *RoomWebhookQuery) OnlyX(ctx context.Context) *RoomWebhook {
	node, err := rwq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RoomWebhook ID in the query.
// Returns a *NotSingularError when more than one RoomWebhook ID is found.
// Returns a *NotFoundError when no entities are found.
func (rwq *RoomWebhookQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = rwq.Limit(2).IDs(setContextOp(ctx, rwq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{roomwebhook.Label}
	default:
		err = &NotSingularError{roomwebhook.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rwq *RoomWebhookQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := rwq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RoomWebhooks.
func (rwq *RoomWebhookQuery) All(ctx context.Context) ([]*RoomWebhook, error) {
	ctx = setContextOp(ctx, rwq.ctx, "All")
	if err := rwq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*RoomWebhook, *RoomWebhookQuery]()
	return withInterceptors[[]*RoomWebhook](ctx, rwq, qr, rwq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rwq *RoomWebhookQuery) AllX(ctx context.Context) []*RoomWebhook {
	nodes, err := rwq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RoomWebhook IDs.
func (rwq *RoomWebhookQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if rwq.ctx.Unique == nil && rwq.path != nil {
		rwq.Unique(true)
	}
	ctx = setContextOp(ctx, rwq.ctx, "IDs")
	if err = rwq.Select(roomwebhook.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rwq *RoomWebhookQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := rwq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rwq *RoomWebhookQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rwq.ctx, "Count")
	if err := rwq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rwq, querierCount[*RoomWebhookQuery](), rwq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rwq *RoomWebhookQuery) CountX(ctx context.Context) int {
	count, err := rwq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rwq *RoomWebhookQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rwq.ctx, "Exist")
	switch _, err := rwq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rwq *RoomWebhookQuery) ExistX(ctx context.Context) bool {
	exist, err := rwq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RoomWebhookQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rwq *RoomWebhookQuery) Clone() *RoomWebhookQuery {
	if rwq == nil {
		return nil
	}
	return &RoomWebhookQuery{
		config:     rwq.config,
		ctx:        rwq.ctx.Clone(),
		order:      append([]roomwebhook.OrderOption{}, rwq.order...),
		inters:     append([]Interceptor{}, rwq.inters...),
		predicates: append([]predicate.RoomWebhook{}, rwq.predicates...),
		// clone intermediate query.
		sql:  rwq.sql.Clone(),
		path: rwq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UUID string `json:"uuid,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.RoomWebhook.Query().
//		GroupBy(roomwebhook.FieldUUID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rwq *RoomWebhookQuery) GroupBy(field string, fields ...string) *RoomWebhookGroupBy {
	rwq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RoomWebhookGroupBy{build: rwq}
	grbuild.flds = &rwq.ctx.Fields
	grbuild.label = roomwebhook.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UUID string `json:"uuid,omitempty"`
//	}
//
//	client.RoomWebhook.Query().
//		Select(roomwebhook.FieldUUID).
//		Scan(ctx, &v)
func (rwq *RoomWebhookQuery) Select(fields ...string) *RoomWebhookSelect {
	rwq.ctx.Fields = append(rwq.ctx.Fields, fields...)
	sbuild := &RoomWebhookSelect{RoomWebhookQuery: rwq}
	sbuild.label = roomwebhook.Label
	sbuild.flds, sbuild.scan = &rwq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RoomWebhookSelect configured with the given aggregations.
func (rwq *RoomWebhookQuery) Aggregate(fns ...AggregateFunc) *RoomWebhookSelect {
	return rwq.Select().Aggregate(fns...)
}

func (rwq *RoomWebhookQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rwq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rwq); err != nil {
				return err
			}
		}
	}
	for _, f := range rwq.ctx.Fields {
		if !roomwebhook.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rwq.path != nil {
		prev, err := rwq.path(ctx)
		if err != nil {
			return err
		}
		rwq.sql = prev
	}
	return nil
}

func (rwq *RoomWebhookQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*RoomWebhook, error) {
	var (
		nodes = []*RoomWebhook{}
		_spec = rwq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*RoomWebhook).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &RoomWebhook{config: rwq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rwq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (rwq *RoomWebhookQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rwq.querySpec()
	_spec.Node.Columns = rwq.ctx.Fields
	if len(rwq.ctx.Fields) > 0 {
		_spec.Unique = rwq.ctx.Unique != nil && *rwq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rwq.driver, _spec)
}

func (rwq *RoomWebhookQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(roomwebhook.Table, roomwebhook.Columns, sqlgraph.NewFieldSpec(roomwebhook.FieldID, field.TypeUint64))
	_spec.From = rwq.sql
	if unique := rwq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rwq.path != nil {
		_spec.Unique = true
	}
	if fields := rwq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, roomwebhook.FieldID)
		for i := range fields {
			if fields[i] != roomwebhook.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rwq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rwq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rwq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rwq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rwq *RoomWebhookQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rwq.driver.Dialect())
	t1 := builder.Table(roomwebhook.Table)
	columns := rwq.ctx.Fields
	if len(columns) == 0 {
		columns = roomwebhook.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rwq.sql != nil {
		selector = rwq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rwq.ctx.Unique != nil && *rwq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rwq.predicates {
		p(selector)
	}
	for _, p := range rwq.order {
		p(selector)
	}
	if offset := rwq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rwq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RoomWebhookGroupBy is the group-by builder for RoomWebhook entities.
type RoomWebhookGroupBy struct {
	selector
	build *RoomWebhookQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rwgb *RoomWebhookGroupBy) Aggregate(fns ...AggregateFunc) *RoomWebhookGroupBy {
	rwgb.fns = append(rwgb.fns, fns...)
	return rwgb
}

// Scan applies the selector query and scans the result into the given value.
func (rwgb *RoomWebhookGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rwgb.build.ctx, "GroupBy")
	if err := rwgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RoomWebhookQuery, *RoomWebhookGroupBy](ctx, rwgb.build, rwgb, rwgb.build.inters, v)
}

func (rwgb *RoomWebhookGroupBy) sqlScan(ctx context.Context, root *RoomWebhookQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rwgb.fns))
	for _, fn := range rwgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rwgb.flds)+len(rwgb.fns))
		for _, f := range *rwgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rwgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rwgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RoomWebhookSelect is the builder for selecting fields of RoomWebhook entities.
type RoomWebhookSelect struct {
	*RoomWebhookQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rws *RoomWebhookSelect) Aggregate(fns ...AggregateFunc) *RoomWebhookSelect {
	rws.fns = append(rws.fns, fns...)
	return rws
}

// Scan applies the selector query and scans the result into the given value.
func (rws *RoomWebhookSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rws.ctx, "Select")
	if err := rws.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RoomWebhookQuery, *RoomWebhookSelect](ctx, rws.RoomWebhookQuery, rws, rws.inters, v)
}

func (rws *RoomWebhookSelect) sqlScan(ctx context.Context, root *RoomWebhookQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rws.fns))
	for _, fn := range rws.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rws.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rws.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

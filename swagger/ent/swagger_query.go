// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"swagger/ent/predicate"
	"swagger/ent/swagger"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SwaggerQuery is the builder for querying Swagger entities.
type SwaggerQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.Swagger
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SwaggerQuery builder.
func (sq *SwaggerQuery) Where(ps ...predicate.Swagger) *SwaggerQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit the number of records to be returned by this query.
func (sq *SwaggerQuery) Limit(limit int) *SwaggerQuery {
	sq.ctx.Limit = &limit
	return sq
}

// Offset to start from.
func (sq *SwaggerQuery) Offset(offset int) *SwaggerQuery {
	sq.ctx.Offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *SwaggerQuery) Unique(unique bool) *SwaggerQuery {
	sq.ctx.Unique = &unique
	return sq
}

// Order specifies how the records should be ordered.
func (sq *SwaggerQuery) Order(o ...OrderFunc) *SwaggerQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// First returns the first Swagger entity from the query.
// Returns a *NotFoundError when no Swagger was found.
func (sq *SwaggerQuery) First(ctx context.Context) (*Swagger, error) {
	nodes, err := sq.Limit(1).All(setContextOp(ctx, sq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{swagger.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *SwaggerQuery) FirstX(ctx context.Context) *Swagger {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Swagger ID from the query.
// Returns a *NotFoundError when no Swagger ID was found.
func (sq *SwaggerQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(1).IDs(setContextOp(ctx, sq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{swagger.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *SwaggerQuery) FirstIDX(ctx context.Context) int {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Swagger entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Swagger entity is found.
// Returns a *NotFoundError when no Swagger entities are found.
func (sq *SwaggerQuery) Only(ctx context.Context) (*Swagger, error) {
	nodes, err := sq.Limit(2).All(setContextOp(ctx, sq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{swagger.Label}
	default:
		return nil, &NotSingularError{swagger.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *SwaggerQuery) OnlyX(ctx context.Context) *Swagger {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Swagger ID in the query.
// Returns a *NotSingularError when more than one Swagger ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *SwaggerQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(2).IDs(setContextOp(ctx, sq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{swagger.Label}
	default:
		err = &NotSingularError{swagger.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *SwaggerQuery) OnlyIDX(ctx context.Context) int {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Swaggers.
func (sq *SwaggerQuery) All(ctx context.Context) ([]*Swagger, error) {
	ctx = setContextOp(ctx, sq.ctx, "All")
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Swagger, *SwaggerQuery]()
	return withInterceptors[[]*Swagger](ctx, sq, qr, sq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sq *SwaggerQuery) AllX(ctx context.Context) []*Swagger {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Swagger IDs.
func (sq *SwaggerQuery) IDs(ctx context.Context) (ids []int, err error) {
	if sq.ctx.Unique == nil && sq.path != nil {
		sq.Unique(true)
	}
	ctx = setContextOp(ctx, sq.ctx, "IDs")
	if err = sq.Select(swagger.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *SwaggerQuery) IDsX(ctx context.Context) []int {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *SwaggerQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sq.ctx, "Count")
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sq, querierCount[*SwaggerQuery](), sq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sq *SwaggerQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *SwaggerQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sq.ctx, "Exist")
	switch _, err := sq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *SwaggerQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SwaggerQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *SwaggerQuery) Clone() *SwaggerQuery {
	if sq == nil {
		return nil
	}
	return &SwaggerQuery{
		config:     sq.config,
		ctx:        sq.ctx.Clone(),
		order:      append([]OrderFunc{}, sq.order...),
		inters:     append([]Interceptor{}, sq.inters...),
		predicates: append([]predicate.Swagger{}, sq.predicates...),
		// clone intermediate query.
		sql:  sq.sql.Clone(),
		path: sq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (sq *SwaggerQuery) GroupBy(field string, fields ...string) *SwaggerGroupBy {
	sq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SwaggerGroupBy{build: sq}
	grbuild.flds = &sq.ctx.Fields
	grbuild.label = swagger.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (sq *SwaggerQuery) Select(fields ...string) *SwaggerSelect {
	sq.ctx.Fields = append(sq.ctx.Fields, fields...)
	sbuild := &SwaggerSelect{SwaggerQuery: sq}
	sbuild.label = swagger.Label
	sbuild.flds, sbuild.scan = &sq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SwaggerSelect configured with the given aggregations.
func (sq *SwaggerQuery) Aggregate(fns ...AggregateFunc) *SwaggerSelect {
	return sq.Select().Aggregate(fns...)
}

func (sq *SwaggerQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sq); err != nil {
				return err
			}
		}
	}
	for _, f := range sq.ctx.Fields {
		if !swagger.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *SwaggerQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Swagger, error) {
	var (
		nodes = []*Swagger{}
		_spec = sq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Swagger).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Swagger{config: sq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (sq *SwaggerQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	_spec.Node.Columns = sq.ctx.Fields
	if len(sq.ctx.Fields) > 0 {
		_spec.Unique = sq.ctx.Unique != nil && *sq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *SwaggerQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(swagger.Table, swagger.Columns, sqlgraph.NewFieldSpec(swagger.FieldID, field.TypeInt))
	_spec.From = sq.sql
	if unique := sq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sq.path != nil {
		_spec.Unique = true
	}
	if fields := sq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, swagger.FieldID)
		for i := range fields {
			if fields[i] != swagger.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *SwaggerQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(swagger.Table)
	columns := sq.ctx.Fields
	if len(columns) == 0 {
		columns = swagger.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.ctx.Unique != nil && *sq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SwaggerGroupBy is the group-by builder for Swagger entities.
type SwaggerGroupBy struct {
	selector
	build *SwaggerQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *SwaggerGroupBy) Aggregate(fns ...AggregateFunc) *SwaggerGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the selector query and scans the result into the given value.
func (sgb *SwaggerGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sgb.build.ctx, "GroupBy")
	if err := sgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SwaggerQuery, *SwaggerGroupBy](ctx, sgb.build, sgb, sgb.build.inters, v)
}

func (sgb *SwaggerGroupBy) sqlScan(ctx context.Context, root *SwaggerQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sgb.flds)+len(sgb.fns))
		for _, f := range *sgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SwaggerSelect is the builder for selecting fields of Swagger entities.
type SwaggerSelect struct {
	*SwaggerQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ss *SwaggerSelect) Aggregate(fns ...AggregateFunc) *SwaggerSelect {
	ss.fns = append(ss.fns, fns...)
	return ss
}

// Scan applies the selector query and scans the result into the given value.
func (ss *SwaggerSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ss.ctx, "Select")
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SwaggerQuery, *SwaggerSelect](ctx, ss.SwaggerQuery, ss, ss.inters, v)
}

func (ss *SwaggerSelect) sqlScan(ctx context.Context, root *SwaggerQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ss.fns))
	for _, fn := range ss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

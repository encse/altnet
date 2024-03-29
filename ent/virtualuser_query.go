// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/encse/altnet/ent/predicate"
	"github.com/encse/altnet/ent/virtualuser"
)

// VirtualUserQuery is the builder for querying VirtualUser entities.
type VirtualUserQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.VirtualUser
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VirtualUserQuery builder.
func (vuq *VirtualUserQuery) Where(ps ...predicate.VirtualUser) *VirtualUserQuery {
	vuq.predicates = append(vuq.predicates, ps...)
	return vuq
}

// Limit the number of records to be returned by this query.
func (vuq *VirtualUserQuery) Limit(limit int) *VirtualUserQuery {
	vuq.ctx.Limit = &limit
	return vuq
}

// Offset to start from.
func (vuq *VirtualUserQuery) Offset(offset int) *VirtualUserQuery {
	vuq.ctx.Offset = &offset
	return vuq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vuq *VirtualUserQuery) Unique(unique bool) *VirtualUserQuery {
	vuq.ctx.Unique = &unique
	return vuq
}

// Order specifies how the records should be ordered.
func (vuq *VirtualUserQuery) Order(o ...OrderFunc) *VirtualUserQuery {
	vuq.order = append(vuq.order, o...)
	return vuq
}

// First returns the first VirtualUser entity from the query.
// Returns a *NotFoundError when no VirtualUser was found.
func (vuq *VirtualUserQuery) First(ctx context.Context) (*VirtualUser, error) {
	nodes, err := vuq.Limit(1).All(setContextOp(ctx, vuq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{virtualuser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vuq *VirtualUserQuery) FirstX(ctx context.Context) *VirtualUser {
	node, err := vuq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first VirtualUser ID from the query.
// Returns a *NotFoundError when no VirtualUser ID was found.
func (vuq *VirtualUserQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = vuq.Limit(1).IDs(setContextOp(ctx, vuq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{virtualuser.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vuq *VirtualUserQuery) FirstIDX(ctx context.Context) int {
	id, err := vuq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single VirtualUser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one VirtualUser entity is found.
// Returns a *NotFoundError when no VirtualUser entities are found.
func (vuq *VirtualUserQuery) Only(ctx context.Context) (*VirtualUser, error) {
	nodes, err := vuq.Limit(2).All(setContextOp(ctx, vuq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{virtualuser.Label}
	default:
		return nil, &NotSingularError{virtualuser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vuq *VirtualUserQuery) OnlyX(ctx context.Context) *VirtualUser {
	node, err := vuq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only VirtualUser ID in the query.
// Returns a *NotSingularError when more than one VirtualUser ID is found.
// Returns a *NotFoundError when no entities are found.
func (vuq *VirtualUserQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = vuq.Limit(2).IDs(setContextOp(ctx, vuq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{virtualuser.Label}
	default:
		err = &NotSingularError{virtualuser.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vuq *VirtualUserQuery) OnlyIDX(ctx context.Context) int {
	id, err := vuq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of VirtualUsers.
func (vuq *VirtualUserQuery) All(ctx context.Context) ([]*VirtualUser, error) {
	ctx = setContextOp(ctx, vuq.ctx, "All")
	if err := vuq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*VirtualUser, *VirtualUserQuery]()
	return withInterceptors[[]*VirtualUser](ctx, vuq, qr, vuq.inters)
}

// AllX is like All, but panics if an error occurs.
func (vuq *VirtualUserQuery) AllX(ctx context.Context) []*VirtualUser {
	nodes, err := vuq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of VirtualUser IDs.
func (vuq *VirtualUserQuery) IDs(ctx context.Context) (ids []int, err error) {
	if vuq.ctx.Unique == nil && vuq.path != nil {
		vuq.Unique(true)
	}
	ctx = setContextOp(ctx, vuq.ctx, "IDs")
	if err = vuq.Select(virtualuser.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vuq *VirtualUserQuery) IDsX(ctx context.Context) []int {
	ids, err := vuq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vuq *VirtualUserQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, vuq.ctx, "Count")
	if err := vuq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, vuq, querierCount[*VirtualUserQuery](), vuq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (vuq *VirtualUserQuery) CountX(ctx context.Context) int {
	count, err := vuq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vuq *VirtualUserQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, vuq.ctx, "Exist")
	switch _, err := vuq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (vuq *VirtualUserQuery) ExistX(ctx context.Context) bool {
	exist, err := vuq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VirtualUserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vuq *VirtualUserQuery) Clone() *VirtualUserQuery {
	if vuq == nil {
		return nil
	}
	return &VirtualUserQuery{
		config:     vuq.config,
		ctx:        vuq.ctx.Clone(),
		order:      append([]OrderFunc{}, vuq.order...),
		inters:     append([]Interceptor{}, vuq.inters...),
		predicates: append([]predicate.VirtualUser{}, vuq.predicates...),
		// clone intermediate query.
		sql:  vuq.sql.Clone(),
		path: vuq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		User schema.Uname `json:"user,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.VirtualUser.Query().
//		GroupBy(virtualuser.FieldUser).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (vuq *VirtualUserQuery) GroupBy(field string, fields ...string) *VirtualUserGroupBy {
	vuq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &VirtualUserGroupBy{build: vuq}
	grbuild.flds = &vuq.ctx.Fields
	grbuild.label = virtualuser.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		User schema.Uname `json:"user,omitempty"`
//	}
//
//	client.VirtualUser.Query().
//		Select(virtualuser.FieldUser).
//		Scan(ctx, &v)
func (vuq *VirtualUserQuery) Select(fields ...string) *VirtualUserSelect {
	vuq.ctx.Fields = append(vuq.ctx.Fields, fields...)
	sbuild := &VirtualUserSelect{VirtualUserQuery: vuq}
	sbuild.label = virtualuser.Label
	sbuild.flds, sbuild.scan = &vuq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a VirtualUserSelect configured with the given aggregations.
func (vuq *VirtualUserQuery) Aggregate(fns ...AggregateFunc) *VirtualUserSelect {
	return vuq.Select().Aggregate(fns...)
}

func (vuq *VirtualUserQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range vuq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, vuq); err != nil {
				return err
			}
		}
	}
	for _, f := range vuq.ctx.Fields {
		if !virtualuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vuq.path != nil {
		prev, err := vuq.path(ctx)
		if err != nil {
			return err
		}
		vuq.sql = prev
	}
	return nil
}

func (vuq *VirtualUserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*VirtualUser, error) {
	var (
		nodes   = []*VirtualUser{}
		withFKs = vuq.withFKs
		_spec   = vuq.querySpec()
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, virtualuser.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*VirtualUser).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &VirtualUser{config: vuq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vuq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (vuq *VirtualUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vuq.querySpec()
	_spec.Node.Columns = vuq.ctx.Fields
	if len(vuq.ctx.Fields) > 0 {
		_spec.Unique = vuq.ctx.Unique != nil && *vuq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, vuq.driver, _spec)
}

func (vuq *VirtualUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(virtualuser.Table, virtualuser.Columns, sqlgraph.NewFieldSpec(virtualuser.FieldID, field.TypeInt))
	_spec.From = vuq.sql
	if unique := vuq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if vuq.path != nil {
		_spec.Unique = true
	}
	if fields := vuq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, virtualuser.FieldID)
		for i := range fields {
			if fields[i] != virtualuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vuq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vuq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vuq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vuq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vuq *VirtualUserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vuq.driver.Dialect())
	t1 := builder.Table(virtualuser.Table)
	columns := vuq.ctx.Fields
	if len(columns) == 0 {
		columns = virtualuser.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vuq.sql != nil {
		selector = vuq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vuq.ctx.Unique != nil && *vuq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range vuq.predicates {
		p(selector)
	}
	for _, p := range vuq.order {
		p(selector)
	}
	if offset := vuq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vuq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// VirtualUserGroupBy is the group-by builder for VirtualUser entities.
type VirtualUserGroupBy struct {
	selector
	build *VirtualUserQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vugb *VirtualUserGroupBy) Aggregate(fns ...AggregateFunc) *VirtualUserGroupBy {
	vugb.fns = append(vugb.fns, fns...)
	return vugb
}

// Scan applies the selector query and scans the result into the given value.
func (vugb *VirtualUserGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vugb.build.ctx, "GroupBy")
	if err := vugb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VirtualUserQuery, *VirtualUserGroupBy](ctx, vugb.build, vugb, vugb.build.inters, v)
}

func (vugb *VirtualUserGroupBy) sqlScan(ctx context.Context, root *VirtualUserQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(vugb.fns))
	for _, fn := range vugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*vugb.flds)+len(vugb.fns))
		for _, f := range *vugb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*vugb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vugb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// VirtualUserSelect is the builder for selecting fields of VirtualUser entities.
type VirtualUserSelect struct {
	*VirtualUserQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (vus *VirtualUserSelect) Aggregate(fns ...AggregateFunc) *VirtualUserSelect {
	vus.fns = append(vus.fns, fns...)
	return vus
}

// Scan applies the selector query and scans the result into the given value.
func (vus *VirtualUserSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vus.ctx, "Select")
	if err := vus.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VirtualUserQuery, *VirtualUserSelect](ctx, vus.VirtualUserQuery, vus, vus.inters, v)
}

func (vus *VirtualUserSelect) sqlScan(ctx context.Context, root *VirtualUserQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(vus.fns))
	for _, fn := range vus.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*vus.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

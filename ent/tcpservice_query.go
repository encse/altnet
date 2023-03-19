// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/encse/altnet/ent/host"
	"github.com/encse/altnet/ent/predicate"
	"github.com/encse/altnet/ent/tcpservice"
)

// TcpServiceQuery is the builder for querying TcpService entities.
type TcpServiceQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.TcpService
	withHosts  *HostQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TcpServiceQuery builder.
func (tsq *TcpServiceQuery) Where(ps ...predicate.TcpService) *TcpServiceQuery {
	tsq.predicates = append(tsq.predicates, ps...)
	return tsq
}

// Limit the number of records to be returned by this query.
func (tsq *TcpServiceQuery) Limit(limit int) *TcpServiceQuery {
	tsq.ctx.Limit = &limit
	return tsq
}

// Offset to start from.
func (tsq *TcpServiceQuery) Offset(offset int) *TcpServiceQuery {
	tsq.ctx.Offset = &offset
	return tsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tsq *TcpServiceQuery) Unique(unique bool) *TcpServiceQuery {
	tsq.ctx.Unique = &unique
	return tsq
}

// Order specifies how the records should be ordered.
func (tsq *TcpServiceQuery) Order(o ...OrderFunc) *TcpServiceQuery {
	tsq.order = append(tsq.order, o...)
	return tsq
}

// QueryHosts chains the current query on the "hosts" edge.
func (tsq *TcpServiceQuery) QueryHosts() *HostQuery {
	query := (&HostClient{config: tsq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tcpservice.Table, tcpservice.FieldID, selector),
			sqlgraph.To(host.Table, host.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, tcpservice.HostsTable, tcpservice.HostsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(tsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TcpService entity from the query.
// Returns a *NotFoundError when no TcpService was found.
func (tsq *TcpServiceQuery) First(ctx context.Context) (*TcpService, error) {
	nodes, err := tsq.Limit(1).All(setContextOp(ctx, tsq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tcpservice.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tsq *TcpServiceQuery) FirstX(ctx context.Context) *TcpService {
	node, err := tsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TcpService ID from the query.
// Returns a *NotFoundError when no TcpService ID was found.
func (tsq *TcpServiceQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tsq.Limit(1).IDs(setContextOp(ctx, tsq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tcpservice.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tsq *TcpServiceQuery) FirstIDX(ctx context.Context) int {
	id, err := tsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TcpService entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TcpService entity is found.
// Returns a *NotFoundError when no TcpService entities are found.
func (tsq *TcpServiceQuery) Only(ctx context.Context) (*TcpService, error) {
	nodes, err := tsq.Limit(2).All(setContextOp(ctx, tsq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tcpservice.Label}
	default:
		return nil, &NotSingularError{tcpservice.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tsq *TcpServiceQuery) OnlyX(ctx context.Context) *TcpService {
	node, err := tsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TcpService ID in the query.
// Returns a *NotSingularError when more than one TcpService ID is found.
// Returns a *NotFoundError when no entities are found.
func (tsq *TcpServiceQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tsq.Limit(2).IDs(setContextOp(ctx, tsq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tcpservice.Label}
	default:
		err = &NotSingularError{tcpservice.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tsq *TcpServiceQuery) OnlyIDX(ctx context.Context) int {
	id, err := tsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TcpServices.
func (tsq *TcpServiceQuery) All(ctx context.Context) ([]*TcpService, error) {
	ctx = setContextOp(ctx, tsq.ctx, "All")
	if err := tsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TcpService, *TcpServiceQuery]()
	return withInterceptors[[]*TcpService](ctx, tsq, qr, tsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tsq *TcpServiceQuery) AllX(ctx context.Context) []*TcpService {
	nodes, err := tsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TcpService IDs.
func (tsq *TcpServiceQuery) IDs(ctx context.Context) (ids []int, err error) {
	if tsq.ctx.Unique == nil && tsq.path != nil {
		tsq.Unique(true)
	}
	ctx = setContextOp(ctx, tsq.ctx, "IDs")
	if err = tsq.Select(tcpservice.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tsq *TcpServiceQuery) IDsX(ctx context.Context) []int {
	ids, err := tsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tsq *TcpServiceQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tsq.ctx, "Count")
	if err := tsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tsq, querierCount[*TcpServiceQuery](), tsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tsq *TcpServiceQuery) CountX(ctx context.Context) int {
	count, err := tsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tsq *TcpServiceQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tsq.ctx, "Exist")
	switch _, err := tsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tsq *TcpServiceQuery) ExistX(ctx context.Context) bool {
	exist, err := tsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TcpServiceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tsq *TcpServiceQuery) Clone() *TcpServiceQuery {
	if tsq == nil {
		return nil
	}
	return &TcpServiceQuery{
		config:     tsq.config,
		ctx:        tsq.ctx.Clone(),
		order:      append([]OrderFunc{}, tsq.order...),
		inters:     append([]Interceptor{}, tsq.inters...),
		predicates: append([]predicate.TcpService{}, tsq.predicates...),
		withHosts:  tsq.withHosts.Clone(),
		// clone intermediate query.
		sql:  tsq.sql.Clone(),
		path: tsq.path,
	}
}

// WithHosts tells the query-builder to eager-load the nodes that are connected to
// the "hosts" edge. The optional arguments are used to configure the query builder of the edge.
func (tsq *TcpServiceQuery) WithHosts(opts ...func(*HostQuery)) *TcpServiceQuery {
	query := (&HostClient{config: tsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tsq.withHosts = query
	return tsq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TcpService.Query().
//		GroupBy(tcpservice.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tsq *TcpServiceQuery) GroupBy(field string, fields ...string) *TcpServiceGroupBy {
	tsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TcpServiceGroupBy{build: tsq}
	grbuild.flds = &tsq.ctx.Fields
	grbuild.label = tcpservice.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.TcpService.Query().
//		Select(tcpservice.FieldName).
//		Scan(ctx, &v)
func (tsq *TcpServiceQuery) Select(fields ...string) *TcpServiceSelect {
	tsq.ctx.Fields = append(tsq.ctx.Fields, fields...)
	sbuild := &TcpServiceSelect{TcpServiceQuery: tsq}
	sbuild.label = tcpservice.Label
	sbuild.flds, sbuild.scan = &tsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TcpServiceSelect configured with the given aggregations.
func (tsq *TcpServiceQuery) Aggregate(fns ...AggregateFunc) *TcpServiceSelect {
	return tsq.Select().Aggregate(fns...)
}

func (tsq *TcpServiceQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tsq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tsq); err != nil {
				return err
			}
		}
	}
	for _, f := range tsq.ctx.Fields {
		if !tcpservice.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tsq.path != nil {
		prev, err := tsq.path(ctx)
		if err != nil {
			return err
		}
		tsq.sql = prev
	}
	return nil
}

func (tsq *TcpServiceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TcpService, error) {
	var (
		nodes       = []*TcpService{}
		_spec       = tsq.querySpec()
		loadedTypes = [1]bool{
			tsq.withHosts != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TcpService).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TcpService{config: tsq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tsq.withHosts; query != nil {
		if err := tsq.loadHosts(ctx, query, nodes,
			func(n *TcpService) { n.Edges.Hosts = []*Host{} },
			func(n *TcpService, e *Host) { n.Edges.Hosts = append(n.Edges.Hosts, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tsq *TcpServiceQuery) loadHosts(ctx context.Context, query *HostQuery, nodes []*TcpService, init func(*TcpService), assign func(*TcpService, *Host)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*TcpService)
	nids := make(map[int]map[*TcpService]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(tcpservice.HostsTable)
		s.Join(joinT).On(s.C(host.FieldID), joinT.C(tcpservice.HostsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(tcpservice.HostsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(tcpservice.HostsPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*TcpService]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Host](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "hosts" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (tsq *TcpServiceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tsq.querySpec()
	_spec.Node.Columns = tsq.ctx.Fields
	if len(tsq.ctx.Fields) > 0 {
		_spec.Unique = tsq.ctx.Unique != nil && *tsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tsq.driver, _spec)
}

func (tsq *TcpServiceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(tcpservice.Table, tcpservice.Columns, sqlgraph.NewFieldSpec(tcpservice.FieldID, field.TypeInt))
	_spec.From = tsq.sql
	if unique := tsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tsq.path != nil {
		_spec.Unique = true
	}
	if fields := tsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tcpservice.FieldID)
		for i := range fields {
			if fields[i] != tcpservice.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tsq *TcpServiceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tsq.driver.Dialect())
	t1 := builder.Table(tcpservice.Table)
	columns := tsq.ctx.Fields
	if len(columns) == 0 {
		columns = tcpservice.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tsq.sql != nil {
		selector = tsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tsq.ctx.Unique != nil && *tsq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tsq.predicates {
		p(selector)
	}
	for _, p := range tsq.order {
		p(selector)
	}
	if offset := tsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TcpServiceGroupBy is the group-by builder for TcpService entities.
type TcpServiceGroupBy struct {
	selector
	build *TcpServiceQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tsgb *TcpServiceGroupBy) Aggregate(fns ...AggregateFunc) *TcpServiceGroupBy {
	tsgb.fns = append(tsgb.fns, fns...)
	return tsgb
}

// Scan applies the selector query and scans the result into the given value.
func (tsgb *TcpServiceGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tsgb.build.ctx, "GroupBy")
	if err := tsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TcpServiceQuery, *TcpServiceGroupBy](ctx, tsgb.build, tsgb, tsgb.build.inters, v)
}

func (tsgb *TcpServiceGroupBy) sqlScan(ctx context.Context, root *TcpServiceQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tsgb.fns))
	for _, fn := range tsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tsgb.flds)+len(tsgb.fns))
		for _, f := range *tsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TcpServiceSelect is the builder for selecting fields of TcpService entities.
type TcpServiceSelect struct {
	*TcpServiceQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tss *TcpServiceSelect) Aggregate(fns ...AggregateFunc) *TcpServiceSelect {
	tss.fns = append(tss.fns, fns...)
	return tss
}

// Scan applies the selector query and scans the result into the given value.
func (tss *TcpServiceSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tss.ctx, "Select")
	if err := tss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TcpServiceQuery, *TcpServiceSelect](ctx, tss.TcpServiceQuery, tss, tss.inters, v)
}

func (tss *TcpServiceSelect) sqlScan(ctx context.Context, root *TcpServiceQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tss.fns))
	for _, fn := range tss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

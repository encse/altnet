// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/encse/altnet/ent/host"
	"github.com/encse/altnet/ent/predicate"
	"github.com/encse/altnet/ent/tcpservice"
)

// TcpServiceUpdate is the builder for updating TcpService entities.
type TcpServiceUpdate struct {
	config
	hooks    []Hook
	mutation *TcpServiceMutation
}

// Where appends a list predicates to the TcpServiceUpdate builder.
func (tsu *TcpServiceUpdate) Where(ps ...predicate.TcpService) *TcpServiceUpdate {
	tsu.mutation.Where(ps...)
	return tsu
}

// AddHostIDs adds the "hosts" edge to the Host entity by IDs.
func (tsu *TcpServiceUpdate) AddHostIDs(ids ...int) *TcpServiceUpdate {
	tsu.mutation.AddHostIDs(ids...)
	return tsu
}

// AddHosts adds the "hosts" edges to the Host entity.
func (tsu *TcpServiceUpdate) AddHosts(h ...*Host) *TcpServiceUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return tsu.AddHostIDs(ids...)
}

// Mutation returns the TcpServiceMutation object of the builder.
func (tsu *TcpServiceUpdate) Mutation() *TcpServiceMutation {
	return tsu.mutation
}

// ClearHosts clears all "hosts" edges to the Host entity.
func (tsu *TcpServiceUpdate) ClearHosts() *TcpServiceUpdate {
	tsu.mutation.ClearHosts()
	return tsu
}

// RemoveHostIDs removes the "hosts" edge to Host entities by IDs.
func (tsu *TcpServiceUpdate) RemoveHostIDs(ids ...int) *TcpServiceUpdate {
	tsu.mutation.RemoveHostIDs(ids...)
	return tsu
}

// RemoveHosts removes "hosts" edges to Host entities.
func (tsu *TcpServiceUpdate) RemoveHosts(h ...*Host) *TcpServiceUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return tsu.RemoveHostIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tsu *TcpServiceUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, TcpServiceMutation](ctx, tsu.sqlSave, tsu.mutation, tsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tsu *TcpServiceUpdate) SaveX(ctx context.Context) int {
	affected, err := tsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tsu *TcpServiceUpdate) Exec(ctx context.Context) error {
	_, err := tsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tsu *TcpServiceUpdate) ExecX(ctx context.Context) {
	if err := tsu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tsu *TcpServiceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(tcpservice.Table, tcpservice.Columns, sqlgraph.NewFieldSpec(tcpservice.FieldID, field.TypeInt))
	if ps := tsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if tsu.mutation.HostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tcpservice.HostsTable,
			Columns: tcpservice.HostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: host.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tsu.mutation.RemovedHostsIDs(); len(nodes) > 0 && !tsu.mutation.HostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tcpservice.HostsTable,
			Columns: tcpservice.HostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: host.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tsu.mutation.HostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tcpservice.HostsTable,
			Columns: tcpservice.HostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: host.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tcpservice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tsu.mutation.done = true
	return n, nil
}

// TcpServiceUpdateOne is the builder for updating a single TcpService entity.
type TcpServiceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TcpServiceMutation
}

// AddHostIDs adds the "hosts" edge to the Host entity by IDs.
func (tsuo *TcpServiceUpdateOne) AddHostIDs(ids ...int) *TcpServiceUpdateOne {
	tsuo.mutation.AddHostIDs(ids...)
	return tsuo
}

// AddHosts adds the "hosts" edges to the Host entity.
func (tsuo *TcpServiceUpdateOne) AddHosts(h ...*Host) *TcpServiceUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return tsuo.AddHostIDs(ids...)
}

// Mutation returns the TcpServiceMutation object of the builder.
func (tsuo *TcpServiceUpdateOne) Mutation() *TcpServiceMutation {
	return tsuo.mutation
}

// ClearHosts clears all "hosts" edges to the Host entity.
func (tsuo *TcpServiceUpdateOne) ClearHosts() *TcpServiceUpdateOne {
	tsuo.mutation.ClearHosts()
	return tsuo
}

// RemoveHostIDs removes the "hosts" edge to Host entities by IDs.
func (tsuo *TcpServiceUpdateOne) RemoveHostIDs(ids ...int) *TcpServiceUpdateOne {
	tsuo.mutation.RemoveHostIDs(ids...)
	return tsuo
}

// RemoveHosts removes "hosts" edges to Host entities.
func (tsuo *TcpServiceUpdateOne) RemoveHosts(h ...*Host) *TcpServiceUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return tsuo.RemoveHostIDs(ids...)
}

// Where appends a list predicates to the TcpServiceUpdate builder.
func (tsuo *TcpServiceUpdateOne) Where(ps ...predicate.TcpService) *TcpServiceUpdateOne {
	tsuo.mutation.Where(ps...)
	return tsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tsuo *TcpServiceUpdateOne) Select(field string, fields ...string) *TcpServiceUpdateOne {
	tsuo.fields = append([]string{field}, fields...)
	return tsuo
}

// Save executes the query and returns the updated TcpService entity.
func (tsuo *TcpServiceUpdateOne) Save(ctx context.Context) (*TcpService, error) {
	return withHooks[*TcpService, TcpServiceMutation](ctx, tsuo.sqlSave, tsuo.mutation, tsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tsuo *TcpServiceUpdateOne) SaveX(ctx context.Context) *TcpService {
	node, err := tsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tsuo *TcpServiceUpdateOne) Exec(ctx context.Context) error {
	_, err := tsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tsuo *TcpServiceUpdateOne) ExecX(ctx context.Context) {
	if err := tsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tsuo *TcpServiceUpdateOne) sqlSave(ctx context.Context) (_node *TcpService, err error) {
	_spec := sqlgraph.NewUpdateSpec(tcpservice.Table, tcpservice.Columns, sqlgraph.NewFieldSpec(tcpservice.FieldID, field.TypeInt))
	id, ok := tsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TcpService.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tcpservice.FieldID)
		for _, f := range fields {
			if !tcpservice.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tcpservice.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if tsuo.mutation.HostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tcpservice.HostsTable,
			Columns: tcpservice.HostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: host.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tsuo.mutation.RemovedHostsIDs(); len(nodes) > 0 && !tsuo.mutation.HostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tcpservice.HostsTable,
			Columns: tcpservice.HostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: host.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tsuo.mutation.HostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tcpservice.HostsTable,
			Columns: tcpservice.HostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: host.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &TcpService{config: tsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tcpservice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tsuo.mutation.done = true
	return _node, nil
}
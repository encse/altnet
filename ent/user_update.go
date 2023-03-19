// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/encse/altnet/ent/host"
	"github.com/encse/altnet/ent/predicate"
	"github.com/encse/altnet/ent/user"
	"github.com/encse/altnet/schema"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUser sets the "user" field.
func (uu *UserUpdate) SetUser(s schema.Uname) *UserUpdate {
	uu.mutation.SetUser(s)
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(sh schema.PasswordHash) *UserUpdate {
	uu.mutation.SetPassword(sh)
	return uu
}

// SetLastLogin sets the "last_login" field.
func (uu *UserUpdate) SetLastLogin(t time.Time) *UserUpdate {
	uu.mutation.SetLastLogin(t)
	return uu
}

// SetLastLoginAttempt sets the "last_login_attempt" field.
func (uu *UserUpdate) SetLastLoginAttempt(t time.Time) *UserUpdate {
	uu.mutation.SetLastLoginAttempt(t)
	return uu
}

// AddHostIDs adds the "hosts" edge to the Host entity by IDs.
func (uu *UserUpdate) AddHostIDs(ids ...int) *UserUpdate {
	uu.mutation.AddHostIDs(ids...)
	return uu
}

// AddHosts adds the "hosts" edges to the Host entity.
func (uu *UserUpdate) AddHosts(h ...*Host) *UserUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return uu.AddHostIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearHosts clears all "hosts" edges to the Host entity.
func (uu *UserUpdate) ClearHosts() *UserUpdate {
	uu.mutation.ClearHosts()
	return uu
}

// RemoveHostIDs removes the "hosts" edge to Host entities by IDs.
func (uu *UserUpdate) RemoveHostIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveHostIDs(ids...)
	return uu
}

// RemoveHosts removes "hosts" edges to Host entities.
func (uu *UserUpdate) RemoveHosts(h ...*Host) *UserUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return uu.RemoveHostIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, UserMutation](ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.User(); ok {
		_spec.SetField(user.FieldUser, field.TypeString, value)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uu.mutation.LastLogin(); ok {
		_spec.SetField(user.FieldLastLogin, field.TypeTime, value)
	}
	if value, ok := uu.mutation.LastLoginAttempt(); ok {
		_spec.SetField(user.FieldLastLoginAttempt, field.TypeTime, value)
	}
	if uu.mutation.HostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.HostsTable,
			Columns: user.HostsPrimaryKey,
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
	if nodes := uu.mutation.RemovedHostsIDs(); len(nodes) > 0 && !uu.mutation.HostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.HostsTable,
			Columns: user.HostsPrimaryKey,
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
	if nodes := uu.mutation.HostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.HostsTable,
			Columns: user.HostsPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUser sets the "user" field.
func (uuo *UserUpdateOne) SetUser(s schema.Uname) *UserUpdateOne {
	uuo.mutation.SetUser(s)
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(sh schema.PasswordHash) *UserUpdateOne {
	uuo.mutation.SetPassword(sh)
	return uuo
}

// SetLastLogin sets the "last_login" field.
func (uuo *UserUpdateOne) SetLastLogin(t time.Time) *UserUpdateOne {
	uuo.mutation.SetLastLogin(t)
	return uuo
}

// SetLastLoginAttempt sets the "last_login_attempt" field.
func (uuo *UserUpdateOne) SetLastLoginAttempt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetLastLoginAttempt(t)
	return uuo
}

// AddHostIDs adds the "hosts" edge to the Host entity by IDs.
func (uuo *UserUpdateOne) AddHostIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddHostIDs(ids...)
	return uuo
}

// AddHosts adds the "hosts" edges to the Host entity.
func (uuo *UserUpdateOne) AddHosts(h ...*Host) *UserUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return uuo.AddHostIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearHosts clears all "hosts" edges to the Host entity.
func (uuo *UserUpdateOne) ClearHosts() *UserUpdateOne {
	uuo.mutation.ClearHosts()
	return uuo
}

// RemoveHostIDs removes the "hosts" edge to Host entities by IDs.
func (uuo *UserUpdateOne) RemoveHostIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveHostIDs(ids...)
	return uuo
}

// RemoveHosts removes "hosts" edges to Host entities.
func (uuo *UserUpdateOne) RemoveHosts(h ...*Host) *UserUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return uuo.RemoveHostIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks[*User, UserMutation](ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.User(); ok {
		_spec.SetField(user.FieldUser, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uuo.mutation.LastLogin(); ok {
		_spec.SetField(user.FieldLastLogin, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.LastLoginAttempt(); ok {
		_spec.SetField(user.FieldLastLoginAttempt, field.TypeTime, value)
	}
	if uuo.mutation.HostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.HostsTable,
			Columns: user.HostsPrimaryKey,
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
	if nodes := uuo.mutation.RemovedHostsIDs(); len(nodes) > 0 && !uuo.mutation.HostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.HostsTable,
			Columns: user.HostsPrimaryKey,
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
	if nodes := uuo.mutation.HostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.HostsTable,
			Columns: user.HostsPrimaryKey,
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
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}

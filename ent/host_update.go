// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/encse/altnet/ent/host"
	"github.com/encse/altnet/ent/predicate"
	"github.com/encse/altnet/ent/schema"
)

// HostUpdate is the builder for updating Host entities.
type HostUpdate struct {
	config
	hooks    []Hook
	mutation *HostMutation
}

// Where appends a list predicates to the HostUpdate builder.
func (hu *HostUpdate) Where(ps ...predicate.Host) *HostUpdate {
	hu.mutation.Where(ps...)
	return hu
}

// SetEntry sets the "entry" field.
func (hu *HostUpdate) SetEntry(s string) *HostUpdate {
	hu.mutation.SetEntry(s)
	return hu
}

// SetNillableEntry sets the "entry" field if the given value is not nil.
func (hu *HostUpdate) SetNillableEntry(s *string) *HostUpdate {
	if s != nil {
		hu.SetEntry(*s)
	}
	return hu
}

// SetMachineType sets the "machine_type" field.
func (hu *HostUpdate) SetMachineType(s string) *HostUpdate {
	hu.mutation.SetMachineType(s)
	return hu
}

// SetNillableMachineType sets the "machine_type" field if the given value is not nil.
func (hu *HostUpdate) SetNillableMachineType(s *string) *HostUpdate {
	if s != nil {
		hu.SetMachineType(*s)
	}
	return hu
}

// SetOrganization sets the "organization" field.
func (hu *HostUpdate) SetOrganization(s string) *HostUpdate {
	hu.mutation.SetOrganization(s)
	return hu
}

// SetNillableOrganization sets the "organization" field if the given value is not nil.
func (hu *HostUpdate) SetNillableOrganization(s *string) *HostUpdate {
	if s != nil {
		hu.SetOrganization(*s)
	}
	return hu
}

// SetContact sets the "contact" field.
func (hu *HostUpdate) SetContact(s string) *HostUpdate {
	hu.mutation.SetContact(s)
	return hu
}

// SetNillableContact sets the "contact" field if the given value is not nil.
func (hu *HostUpdate) SetNillableContact(s *string) *HostUpdate {
	if s != nil {
		hu.SetContact(*s)
	}
	return hu
}

// SetContactAddress sets the "contact_address" field.
func (hu *HostUpdate) SetContactAddress(s string) *HostUpdate {
	hu.mutation.SetContactAddress(s)
	return hu
}

// SetNillableContactAddress sets the "contact_address" field if the given value is not nil.
func (hu *HostUpdate) SetNillableContactAddress(s *string) *HostUpdate {
	if s != nil {
		hu.SetContactAddress(*s)
	}
	return hu
}

// SetCountry sets the "country" field.
func (hu *HostUpdate) SetCountry(s string) *HostUpdate {
	hu.mutation.SetCountry(s)
	return hu
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (hu *HostUpdate) SetNillableCountry(s *string) *HostUpdate {
	if s != nil {
		hu.SetCountry(*s)
	}
	return hu
}

// SetLocation sets the "location" field.
func (hu *HostUpdate) SetLocation(s string) *HostUpdate {
	hu.mutation.SetLocation(s)
	return hu
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (hu *HostUpdate) SetNillableLocation(s *string) *HostUpdate {
	if s != nil {
		hu.SetLocation(*s)
	}
	return hu
}

// SetGeoLocation sets the "geo_location" field.
func (hu *HostUpdate) SetGeoLocation(s string) *HostUpdate {
	hu.mutation.SetGeoLocation(s)
	return hu
}

// SetNillableGeoLocation sets the "geo_location" field if the given value is not nil.
func (hu *HostUpdate) SetNillableGeoLocation(s *string) *HostUpdate {
	if s != nil {
		hu.SetGeoLocation(*s)
	}
	return hu
}

// SetPhone sets the "phone" field.
func (hu *HostUpdate) SetPhone(sn []schema.PhoneNumber) *HostUpdate {
	hu.mutation.SetPhone(sn)
	return hu
}

// AppendPhone appends sn to the "phone" field.
func (hu *HostUpdate) AppendPhone(sn []schema.PhoneNumber) *HostUpdate {
	hu.mutation.AppendPhone(sn)
	return hu
}

// ClearPhone clears the value of the "phone" field.
func (hu *HostUpdate) ClearPhone() *HostUpdate {
	hu.mutation.ClearPhone()
	return hu
}

// SetNeighbours sets the "neighbours" field.
func (hu *HostUpdate) SetNeighbours(sn []schema.HostName) *HostUpdate {
	hu.mutation.SetNeighbours(sn)
	return hu
}

// AppendNeighbours appends sn to the "neighbours" field.
func (hu *HostUpdate) AppendNeighbours(sn []schema.HostName) *HostUpdate {
	hu.mutation.AppendNeighbours(sn)
	return hu
}

// ClearNeighbours clears the value of the "neighbours" field.
func (hu *HostUpdate) ClearNeighbours() *HostUpdate {
	hu.mutation.ClearNeighbours()
	return hu
}

// Mutation returns the HostMutation object of the builder.
func (hu *HostUpdate) Mutation() *HostMutation {
	return hu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (hu *HostUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, HostMutation](ctx, hu.sqlSave, hu.mutation, hu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (hu *HostUpdate) SaveX(ctx context.Context) int {
	affected, err := hu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hu *HostUpdate) Exec(ctx context.Context) error {
	_, err := hu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hu *HostUpdate) ExecX(ctx context.Context) {
	if err := hu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (hu *HostUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(host.Table, host.Columns, sqlgraph.NewFieldSpec(host.FieldID, field.TypeInt))
	if ps := hu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hu.mutation.Entry(); ok {
		_spec.SetField(host.FieldEntry, field.TypeString, value)
	}
	if value, ok := hu.mutation.MachineType(); ok {
		_spec.SetField(host.FieldMachineType, field.TypeString, value)
	}
	if value, ok := hu.mutation.Organization(); ok {
		_spec.SetField(host.FieldOrganization, field.TypeString, value)
	}
	if value, ok := hu.mutation.Contact(); ok {
		_spec.SetField(host.FieldContact, field.TypeString, value)
	}
	if value, ok := hu.mutation.ContactAddress(); ok {
		_spec.SetField(host.FieldContactAddress, field.TypeString, value)
	}
	if value, ok := hu.mutation.Country(); ok {
		_spec.SetField(host.FieldCountry, field.TypeString, value)
	}
	if value, ok := hu.mutation.Location(); ok {
		_spec.SetField(host.FieldLocation, field.TypeString, value)
	}
	if value, ok := hu.mutation.GeoLocation(); ok {
		_spec.SetField(host.FieldGeoLocation, field.TypeString, value)
	}
	if value, ok := hu.mutation.Phone(); ok {
		_spec.SetField(host.FieldPhone, field.TypeJSON, value)
	}
	if value, ok := hu.mutation.AppendedPhone(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, host.FieldPhone, value)
		})
	}
	if hu.mutation.PhoneCleared() {
		_spec.ClearField(host.FieldPhone, field.TypeJSON)
	}
	if value, ok := hu.mutation.Neighbours(); ok {
		_spec.SetField(host.FieldNeighbours, field.TypeJSON, value)
	}
	if value, ok := hu.mutation.AppendedNeighbours(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, host.FieldNeighbours, value)
		})
	}
	if hu.mutation.NeighboursCleared() {
		_spec.ClearField(host.FieldNeighbours, field.TypeJSON)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, hu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{host.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	hu.mutation.done = true
	return n, nil
}

// HostUpdateOne is the builder for updating a single Host entity.
type HostUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *HostMutation
}

// SetEntry sets the "entry" field.
func (huo *HostUpdateOne) SetEntry(s string) *HostUpdateOne {
	huo.mutation.SetEntry(s)
	return huo
}

// SetNillableEntry sets the "entry" field if the given value is not nil.
func (huo *HostUpdateOne) SetNillableEntry(s *string) *HostUpdateOne {
	if s != nil {
		huo.SetEntry(*s)
	}
	return huo
}

// SetMachineType sets the "machine_type" field.
func (huo *HostUpdateOne) SetMachineType(s string) *HostUpdateOne {
	huo.mutation.SetMachineType(s)
	return huo
}

// SetNillableMachineType sets the "machine_type" field if the given value is not nil.
func (huo *HostUpdateOne) SetNillableMachineType(s *string) *HostUpdateOne {
	if s != nil {
		huo.SetMachineType(*s)
	}
	return huo
}

// SetOrganization sets the "organization" field.
func (huo *HostUpdateOne) SetOrganization(s string) *HostUpdateOne {
	huo.mutation.SetOrganization(s)
	return huo
}

// SetNillableOrganization sets the "organization" field if the given value is not nil.
func (huo *HostUpdateOne) SetNillableOrganization(s *string) *HostUpdateOne {
	if s != nil {
		huo.SetOrganization(*s)
	}
	return huo
}

// SetContact sets the "contact" field.
func (huo *HostUpdateOne) SetContact(s string) *HostUpdateOne {
	huo.mutation.SetContact(s)
	return huo
}

// SetNillableContact sets the "contact" field if the given value is not nil.
func (huo *HostUpdateOne) SetNillableContact(s *string) *HostUpdateOne {
	if s != nil {
		huo.SetContact(*s)
	}
	return huo
}

// SetContactAddress sets the "contact_address" field.
func (huo *HostUpdateOne) SetContactAddress(s string) *HostUpdateOne {
	huo.mutation.SetContactAddress(s)
	return huo
}

// SetNillableContactAddress sets the "contact_address" field if the given value is not nil.
func (huo *HostUpdateOne) SetNillableContactAddress(s *string) *HostUpdateOne {
	if s != nil {
		huo.SetContactAddress(*s)
	}
	return huo
}

// SetCountry sets the "country" field.
func (huo *HostUpdateOne) SetCountry(s string) *HostUpdateOne {
	huo.mutation.SetCountry(s)
	return huo
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (huo *HostUpdateOne) SetNillableCountry(s *string) *HostUpdateOne {
	if s != nil {
		huo.SetCountry(*s)
	}
	return huo
}

// SetLocation sets the "location" field.
func (huo *HostUpdateOne) SetLocation(s string) *HostUpdateOne {
	huo.mutation.SetLocation(s)
	return huo
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (huo *HostUpdateOne) SetNillableLocation(s *string) *HostUpdateOne {
	if s != nil {
		huo.SetLocation(*s)
	}
	return huo
}

// SetGeoLocation sets the "geo_location" field.
func (huo *HostUpdateOne) SetGeoLocation(s string) *HostUpdateOne {
	huo.mutation.SetGeoLocation(s)
	return huo
}

// SetNillableGeoLocation sets the "geo_location" field if the given value is not nil.
func (huo *HostUpdateOne) SetNillableGeoLocation(s *string) *HostUpdateOne {
	if s != nil {
		huo.SetGeoLocation(*s)
	}
	return huo
}

// SetPhone sets the "phone" field.
func (huo *HostUpdateOne) SetPhone(sn []schema.PhoneNumber) *HostUpdateOne {
	huo.mutation.SetPhone(sn)
	return huo
}

// AppendPhone appends sn to the "phone" field.
func (huo *HostUpdateOne) AppendPhone(sn []schema.PhoneNumber) *HostUpdateOne {
	huo.mutation.AppendPhone(sn)
	return huo
}

// ClearPhone clears the value of the "phone" field.
func (huo *HostUpdateOne) ClearPhone() *HostUpdateOne {
	huo.mutation.ClearPhone()
	return huo
}

// SetNeighbours sets the "neighbours" field.
func (huo *HostUpdateOne) SetNeighbours(sn []schema.HostName) *HostUpdateOne {
	huo.mutation.SetNeighbours(sn)
	return huo
}

// AppendNeighbours appends sn to the "neighbours" field.
func (huo *HostUpdateOne) AppendNeighbours(sn []schema.HostName) *HostUpdateOne {
	huo.mutation.AppendNeighbours(sn)
	return huo
}

// ClearNeighbours clears the value of the "neighbours" field.
func (huo *HostUpdateOne) ClearNeighbours() *HostUpdateOne {
	huo.mutation.ClearNeighbours()
	return huo
}

// Mutation returns the HostMutation object of the builder.
func (huo *HostUpdateOne) Mutation() *HostMutation {
	return huo.mutation
}

// Where appends a list predicates to the HostUpdate builder.
func (huo *HostUpdateOne) Where(ps ...predicate.Host) *HostUpdateOne {
	huo.mutation.Where(ps...)
	return huo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (huo *HostUpdateOne) Select(field string, fields ...string) *HostUpdateOne {
	huo.fields = append([]string{field}, fields...)
	return huo
}

// Save executes the query and returns the updated Host entity.
func (huo *HostUpdateOne) Save(ctx context.Context) (*Host, error) {
	return withHooks[*Host, HostMutation](ctx, huo.sqlSave, huo.mutation, huo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (huo *HostUpdateOne) SaveX(ctx context.Context) *Host {
	node, err := huo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (huo *HostUpdateOne) Exec(ctx context.Context) error {
	_, err := huo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (huo *HostUpdateOne) ExecX(ctx context.Context) {
	if err := huo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (huo *HostUpdateOne) sqlSave(ctx context.Context) (_node *Host, err error) {
	_spec := sqlgraph.NewUpdateSpec(host.Table, host.Columns, sqlgraph.NewFieldSpec(host.FieldID, field.TypeInt))
	id, ok := huo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Host.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := huo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, host.FieldID)
		for _, f := range fields {
			if !host.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != host.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := huo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := huo.mutation.Entry(); ok {
		_spec.SetField(host.FieldEntry, field.TypeString, value)
	}
	if value, ok := huo.mutation.MachineType(); ok {
		_spec.SetField(host.FieldMachineType, field.TypeString, value)
	}
	if value, ok := huo.mutation.Organization(); ok {
		_spec.SetField(host.FieldOrganization, field.TypeString, value)
	}
	if value, ok := huo.mutation.Contact(); ok {
		_spec.SetField(host.FieldContact, field.TypeString, value)
	}
	if value, ok := huo.mutation.ContactAddress(); ok {
		_spec.SetField(host.FieldContactAddress, field.TypeString, value)
	}
	if value, ok := huo.mutation.Country(); ok {
		_spec.SetField(host.FieldCountry, field.TypeString, value)
	}
	if value, ok := huo.mutation.Location(); ok {
		_spec.SetField(host.FieldLocation, field.TypeString, value)
	}
	if value, ok := huo.mutation.GeoLocation(); ok {
		_spec.SetField(host.FieldGeoLocation, field.TypeString, value)
	}
	if value, ok := huo.mutation.Phone(); ok {
		_spec.SetField(host.FieldPhone, field.TypeJSON, value)
	}
	if value, ok := huo.mutation.AppendedPhone(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, host.FieldPhone, value)
		})
	}
	if huo.mutation.PhoneCleared() {
		_spec.ClearField(host.FieldPhone, field.TypeJSON)
	}
	if value, ok := huo.mutation.Neighbours(); ok {
		_spec.SetField(host.FieldNeighbours, field.TypeJSON, value)
	}
	if value, ok := huo.mutation.AppendedNeighbours(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, host.FieldNeighbours, value)
		})
	}
	if huo.mutation.NeighboursCleared() {
		_spec.ClearField(host.FieldNeighbours, field.TypeJSON)
	}
	_node = &Host{config: huo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, huo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{host.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	huo.mutation.done = true
	return _node, nil
}

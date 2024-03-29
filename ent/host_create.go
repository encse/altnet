// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/encse/altnet/ent/host"
	"github.com/encse/altnet/ent/tcpservice"
	"github.com/encse/altnet/ent/user"
	"github.com/encse/altnet/ent/virtualuser"
	"github.com/encse/altnet/schema"
)

// HostCreate is the builder for creating a Host entity.
type HostCreate struct {
	config
	mutation *HostMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (hc *HostCreate) SetName(sn schema.HostName) *HostCreate {
	hc.mutation.SetName(sn)
	return hc
}

// SetType sets the "type" field.
func (hc *HostCreate) SetType(h host.Type) *HostCreate {
	hc.mutation.SetType(h)
	return hc
}

// SetEntry sets the "entry" field.
func (hc *HostCreate) SetEntry(s string) *HostCreate {
	hc.mutation.SetEntry(s)
	return hc
}

// SetNillableEntry sets the "entry" field if the given value is not nil.
func (hc *HostCreate) SetNillableEntry(s *string) *HostCreate {
	if s != nil {
		hc.SetEntry(*s)
	}
	return hc
}

// SetMachineType sets the "machine_type" field.
func (hc *HostCreate) SetMachineType(s string) *HostCreate {
	hc.mutation.SetMachineType(s)
	return hc
}

// SetNillableMachineType sets the "machine_type" field if the given value is not nil.
func (hc *HostCreate) SetNillableMachineType(s *string) *HostCreate {
	if s != nil {
		hc.SetMachineType(*s)
	}
	return hc
}

// SetOrganization sets the "organization" field.
func (hc *HostCreate) SetOrganization(s string) *HostCreate {
	hc.mutation.SetOrganization(s)
	return hc
}

// SetNillableOrganization sets the "organization" field if the given value is not nil.
func (hc *HostCreate) SetNillableOrganization(s *string) *HostCreate {
	if s != nil {
		hc.SetOrganization(*s)
	}
	return hc
}

// SetContact sets the "contact" field.
func (hc *HostCreate) SetContact(s string) *HostCreate {
	hc.mutation.SetContact(s)
	return hc
}

// SetNillableContact sets the "contact" field if the given value is not nil.
func (hc *HostCreate) SetNillableContact(s *string) *HostCreate {
	if s != nil {
		hc.SetContact(*s)
	}
	return hc
}

// SetContactAddress sets the "contact_address" field.
func (hc *HostCreate) SetContactAddress(s string) *HostCreate {
	hc.mutation.SetContactAddress(s)
	return hc
}

// SetNillableContactAddress sets the "contact_address" field if the given value is not nil.
func (hc *HostCreate) SetNillableContactAddress(s *string) *HostCreate {
	if s != nil {
		hc.SetContactAddress(*s)
	}
	return hc
}

// SetCountry sets the "country" field.
func (hc *HostCreate) SetCountry(s string) *HostCreate {
	hc.mutation.SetCountry(s)
	return hc
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (hc *HostCreate) SetNillableCountry(s *string) *HostCreate {
	if s != nil {
		hc.SetCountry(*s)
	}
	return hc
}

// SetLocation sets the "location" field.
func (hc *HostCreate) SetLocation(s string) *HostCreate {
	hc.mutation.SetLocation(s)
	return hc
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (hc *HostCreate) SetNillableLocation(s *string) *HostCreate {
	if s != nil {
		hc.SetLocation(*s)
	}
	return hc
}

// SetGeoLocation sets the "geo_location" field.
func (hc *HostCreate) SetGeoLocation(s string) *HostCreate {
	hc.mutation.SetGeoLocation(s)
	return hc
}

// SetNillableGeoLocation sets the "geo_location" field if the given value is not nil.
func (hc *HostCreate) SetNillableGeoLocation(s *string) *HostCreate {
	if s != nil {
		hc.SetGeoLocation(*s)
	}
	return hc
}

// SetPhone sets the "phone" field.
func (hc *HostCreate) SetPhone(sn []schema.PhoneNumber) *HostCreate {
	hc.mutation.SetPhone(sn)
	return hc
}

// SetNeighbours sets the "neighbours" field.
func (hc *HostCreate) SetNeighbours(sn []schema.HostName) *HostCreate {
	hc.mutation.SetNeighbours(sn)
	return hc
}

// AddServiceIDs adds the "services" edge to the TcpService entity by IDs.
func (hc *HostCreate) AddServiceIDs(ids ...int) *HostCreate {
	hc.mutation.AddServiceIDs(ids...)
	return hc
}

// AddServices adds the "services" edges to the TcpService entity.
func (hc *HostCreate) AddServices(t ...*TcpService) *HostCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return hc.AddServiceIDs(ids...)
}

// AddVirtualuserIDs adds the "virtualusers" edge to the VirtualUser entity by IDs.
func (hc *HostCreate) AddVirtualuserIDs(ids ...int) *HostCreate {
	hc.mutation.AddVirtualuserIDs(ids...)
	return hc
}

// AddVirtualusers adds the "virtualusers" edges to the VirtualUser entity.
func (hc *HostCreate) AddVirtualusers(v ...*VirtualUser) *HostCreate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return hc.AddVirtualuserIDs(ids...)
}

// AddHackerIDs adds the "hackers" edge to the User entity by IDs.
func (hc *HostCreate) AddHackerIDs(ids ...int) *HostCreate {
	hc.mutation.AddHackerIDs(ids...)
	return hc
}

// AddHackers adds the "hackers" edges to the User entity.
func (hc *HostCreate) AddHackers(u ...*User) *HostCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return hc.AddHackerIDs(ids...)
}

// Mutation returns the HostMutation object of the builder.
func (hc *HostCreate) Mutation() *HostMutation {
	return hc.mutation
}

// Save creates the Host in the database.
func (hc *HostCreate) Save(ctx context.Context) (*Host, error) {
	hc.defaults()
	return withHooks[*Host, HostMutation](ctx, hc.sqlSave, hc.mutation, hc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (hc *HostCreate) SaveX(ctx context.Context) *Host {
	v, err := hc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hc *HostCreate) Exec(ctx context.Context) error {
	_, err := hc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hc *HostCreate) ExecX(ctx context.Context) {
	if err := hc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hc *HostCreate) defaults() {
	if _, ok := hc.mutation.Entry(); !ok {
		v := host.DefaultEntry
		hc.mutation.SetEntry(v)
	}
	if _, ok := hc.mutation.MachineType(); !ok {
		v := host.DefaultMachineType
		hc.mutation.SetMachineType(v)
	}
	if _, ok := hc.mutation.Organization(); !ok {
		v := host.DefaultOrganization
		hc.mutation.SetOrganization(v)
	}
	if _, ok := hc.mutation.Contact(); !ok {
		v := host.DefaultContact
		hc.mutation.SetContact(v)
	}
	if _, ok := hc.mutation.ContactAddress(); !ok {
		v := host.DefaultContactAddress
		hc.mutation.SetContactAddress(v)
	}
	if _, ok := hc.mutation.Country(); !ok {
		v := host.DefaultCountry
		hc.mutation.SetCountry(v)
	}
	if _, ok := hc.mutation.Location(); !ok {
		v := host.DefaultLocation
		hc.mutation.SetLocation(v)
	}
	if _, ok := hc.mutation.GeoLocation(); !ok {
		v := host.DefaultGeoLocation
		hc.mutation.SetGeoLocation(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hc *HostCreate) check() error {
	if _, ok := hc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Host.name"`)}
	}
	if _, ok := hc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Host.type"`)}
	}
	if v, ok := hc.mutation.GetType(); ok {
		if err := host.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Host.type": %w`, err)}
		}
	}
	if _, ok := hc.mutation.Entry(); !ok {
		return &ValidationError{Name: "entry", err: errors.New(`ent: missing required field "Host.entry"`)}
	}
	if _, ok := hc.mutation.MachineType(); !ok {
		return &ValidationError{Name: "machine_type", err: errors.New(`ent: missing required field "Host.machine_type"`)}
	}
	if _, ok := hc.mutation.Organization(); !ok {
		return &ValidationError{Name: "organization", err: errors.New(`ent: missing required field "Host.organization"`)}
	}
	if _, ok := hc.mutation.Contact(); !ok {
		return &ValidationError{Name: "contact", err: errors.New(`ent: missing required field "Host.contact"`)}
	}
	if _, ok := hc.mutation.ContactAddress(); !ok {
		return &ValidationError{Name: "contact_address", err: errors.New(`ent: missing required field "Host.contact_address"`)}
	}
	if _, ok := hc.mutation.Country(); !ok {
		return &ValidationError{Name: "country", err: errors.New(`ent: missing required field "Host.country"`)}
	}
	if _, ok := hc.mutation.Location(); !ok {
		return &ValidationError{Name: "location", err: errors.New(`ent: missing required field "Host.location"`)}
	}
	if _, ok := hc.mutation.GeoLocation(); !ok {
		return &ValidationError{Name: "geo_location", err: errors.New(`ent: missing required field "Host.geo_location"`)}
	}
	return nil
}

func (hc *HostCreate) sqlSave(ctx context.Context) (*Host, error) {
	if err := hc.check(); err != nil {
		return nil, err
	}
	_node, _spec := hc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	hc.mutation.id = &_node.ID
	hc.mutation.done = true
	return _node, nil
}

func (hc *HostCreate) createSpec() (*Host, *sqlgraph.CreateSpec) {
	var (
		_node = &Host{config: hc.config}
		_spec = sqlgraph.NewCreateSpec(host.Table, sqlgraph.NewFieldSpec(host.FieldID, field.TypeInt))
	)
	if value, ok := hc.mutation.Name(); ok {
		_spec.SetField(host.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := hc.mutation.GetType(); ok {
		_spec.SetField(host.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := hc.mutation.Entry(); ok {
		_spec.SetField(host.FieldEntry, field.TypeString, value)
		_node.Entry = value
	}
	if value, ok := hc.mutation.MachineType(); ok {
		_spec.SetField(host.FieldMachineType, field.TypeString, value)
		_node.MachineType = value
	}
	if value, ok := hc.mutation.Organization(); ok {
		_spec.SetField(host.FieldOrganization, field.TypeString, value)
		_node.Organization = value
	}
	if value, ok := hc.mutation.Contact(); ok {
		_spec.SetField(host.FieldContact, field.TypeString, value)
		_node.Contact = value
	}
	if value, ok := hc.mutation.ContactAddress(); ok {
		_spec.SetField(host.FieldContactAddress, field.TypeString, value)
		_node.ContactAddress = value
	}
	if value, ok := hc.mutation.Country(); ok {
		_spec.SetField(host.FieldCountry, field.TypeString, value)
		_node.Country = value
	}
	if value, ok := hc.mutation.Location(); ok {
		_spec.SetField(host.FieldLocation, field.TypeString, value)
		_node.Location = value
	}
	if value, ok := hc.mutation.GeoLocation(); ok {
		_spec.SetField(host.FieldGeoLocation, field.TypeString, value)
		_node.GeoLocation = value
	}
	if value, ok := hc.mutation.Phone(); ok {
		_spec.SetField(host.FieldPhone, field.TypeJSON, value)
		_node.Phone = value
	}
	if value, ok := hc.mutation.Neighbours(); ok {
		_spec.SetField(host.FieldNeighbours, field.TypeJSON, value)
		_node.Neighbours = value
	}
	if nodes := hc.mutation.ServicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   host.ServicesTable,
			Columns: host.ServicesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tcpservice.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hc.mutation.VirtualusersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   host.VirtualusersTable,
			Columns: []string{host.VirtualusersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: virtualuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hc.mutation.HackersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   host.HackersTable,
			Columns: host.HackersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// HostCreateBulk is the builder for creating many Host entities in bulk.
type HostCreateBulk struct {
	config
	builders []*HostCreate
}

// Save creates the Host entities in the database.
func (hcb *HostCreateBulk) Save(ctx context.Context) ([]*Host, error) {
	specs := make([]*sqlgraph.CreateSpec, len(hcb.builders))
	nodes := make([]*Host, len(hcb.builders))
	mutators := make([]Mutator, len(hcb.builders))
	for i := range hcb.builders {
		func(i int, root context.Context) {
			builder := hcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HostMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, hcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, hcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hcb *HostCreateBulk) SaveX(ctx context.Context) []*Host {
	v, err := hcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hcb *HostCreateBulk) Exec(ctx context.Context) error {
	_, err := hcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hcb *HostCreateBulk) ExecX(ctx context.Context) {
	if err := hcb.Exec(ctx); err != nil {
		panic(err)
	}
}

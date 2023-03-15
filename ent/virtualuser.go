// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/encse/altnet/ent/schema"
	"github.com/encse/altnet/ent/virtualuser"
)

// VirtualUser is the model entity for the VirtualUser schema.
type VirtualUser struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// User holds the value of the "user" field.
	User schema.Uname `json:"user,omitempty"`
	// Password holds the value of the "password" field.
	Password          schema.Password `json:"password,omitempty"`
	host_virtualusers *int
}

// scanValues returns the types for scanning values from sql.Rows.
func (*VirtualUser) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case virtualuser.FieldID:
			values[i] = new(sql.NullInt64)
		case virtualuser.FieldUser, virtualuser.FieldPassword:
			values[i] = new(sql.NullString)
		case virtualuser.ForeignKeys[0]: // host_virtualusers
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type VirtualUser", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the VirtualUser fields.
func (vu *VirtualUser) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case virtualuser.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			vu.ID = int(value.Int64)
		case virtualuser.FieldUser:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user", values[i])
			} else if value.Valid {
				vu.User = schema.Uname(value.String)
			}
		case virtualuser.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				vu.Password = schema.Password(value.String)
			}
		case virtualuser.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field host_virtualusers", value)
			} else if value.Valid {
				vu.host_virtualusers = new(int)
				*vu.host_virtualusers = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this VirtualUser.
// Note that you need to call VirtualUser.Unwrap() before calling this method if this VirtualUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (vu *VirtualUser) Update() *VirtualUserUpdateOne {
	return NewVirtualUserClient(vu.config).UpdateOne(vu)
}

// Unwrap unwraps the VirtualUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (vu *VirtualUser) Unwrap() *VirtualUser {
	_tx, ok := vu.config.driver.(*txDriver)
	if !ok {
		panic("ent: VirtualUser is not a transactional entity")
	}
	vu.config.driver = _tx.drv
	return vu
}

// String implements the fmt.Stringer.
func (vu *VirtualUser) String() string {
	var builder strings.Builder
	builder.WriteString("VirtualUser(")
	builder.WriteString(fmt.Sprintf("id=%v, ", vu.ID))
	builder.WriteString("user=")
	builder.WriteString(fmt.Sprintf("%v", vu.User))
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(fmt.Sprintf("%v", vu.Password))
	builder.WriteByte(')')
	return builder.String()
}

// VirtualUsers is a parsable slice of VirtualUser.
type VirtualUsers []*VirtualUser
